"""Fit the netperfbench cost model from repeated, load-qualified runs.

The headline for an idle run is the median. Bootstrap intervals show sampling
uncertainty from the repetitions; they do not repair a loaded machine.

Usage: python3 analyze.py raw.jsonl
"""

import argparse
import json
import math
import random
import sys
from collections import defaultdict


BOOTSTRAPS = 10_000
SMALL_LIMIT = 64 << 10


def load(path):
    """Read concatenated, pretty-printed JSON objects."""
    results = []
    decoder = json.JSONDecoder()
    with open(path) as f:
        text = f.read()
    index = 0
    while index < len(text):
        while index < len(text) and text[index].isspace():
            index += 1
        if index >= len(text):
            break
        result, index = decoder.raw_decode(text, index)
        results.append(result)
    return results


def median(values):
    values = sorted(values)
    mid = len(values) // 2
    if len(values) % 2:
        return values[mid]
    return (values[mid - 1] + values[mid]) / 2


def interval(values):
    values = sorted(values)
    if not values:
        return (math.nan, math.nan)
    return (values[int(0.025 * (len(values) - 1))],
            values[int(0.975 * (len(values) - 1))])


def fmt(value, digits=1):
    if math.isnan(value):
        return "-"
    return f"{value:.{digits}f}"


def fit(points):
    """Least-squares fit of y = a + b*x."""
    n = len(points)
    sx = sum(x for x, _ in points)
    sy = sum(y for _, y in points)
    sxx = sum(x * x for x, _ in points)
    sxy = sum(x * y for x, y in points)
    denom = n * sxx - sx * sx
    if denom == 0:
        return sy / n, 0
    slope = (n * sxy - sx * sy) / denom
    return (sy - slope * sx) / n, slope


def observations(result):
    """Return per-repetition p50 latency and throughput observations."""
    runs = result.get("runs") or [result]
    values = []
    for run in runs:
        latency = run.get("p50_us_per_message")
        elapsed = run.get("elapsed_sec")
        if latency is None or elapsed is None:
            continue
        throughput = run.get("throughput_mbps")
        if throughput is None:
            messages = result["round_trips"] * result["inflight"]
            throughput = messages * result["payload_bytes"] * 2 / elapsed / 1048576
        values.append((latency, throughput))
    return values


def statistic(values, stat):
    if stat == "minimum":
        return min(values)
    return median(values)


def boot_cell(values, rng, stat):
    return statistic([values[rng.randrange(len(values))] for _ in values], stat)


def term_samples(cells, impl, deepest, rng, stat):
    """Bootstrap fixed, bandwidth, amortized, and serialized terms."""
    small_sizes = sorted({
        size for name, depth, size in cells
        if name == impl and depth == 1 and size <= SMALL_LIMIT
    })
    amortized = cells.get((impl, deepest, 4096))
    if len(small_sizes) < 2 or not amortized:
        return []

    values = []
    for _ in range(BOOTSTRAPS):
        points = [
            (size, boot_cell(cells[(impl, 1, size)], rng, stat))
            for size in small_sizes
        ]
        fixed, slope = fit(points)
        if slope <= 0:
            continue
        bandwidth = 1e6 / slope / (1 << 20)
        overlap = boot_cell(amortized, rng, stat)
        values.append((fixed, bandwidth, overlap, fixed - overlap))
    return values


def summary(samples, column, digits=1):
    values = [sample[column] for sample in samples]
    lo, hi = interval(values)
    return f"{fmt(median(values), digits)} [{fmt(lo, digits)}, {fmt(hi, digits)}]"


def delta(samples, left, right, column):
    count = min(len(left), len(right))
    return [left[i][column] - right[i][column] for i in range(count)]


def verdict(values):
    lo, hi = interval(values)
    if lo > 0 or hi < 0:
        return "different from zero (95% bootstrap interval excludes zero)"
    return "not resolved (95% bootstrap interval includes zero)"


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("path", help="concatenated JSON results")
    parser.add_argument("--stat", choices=("median", "minimum"), default="median",
                        help="headline statistic (default: median)")
    args = parser.parse_args()
    results = load(args.path)
    cells = defaultdict(list)
    loads = set()
    for result in results:
        key = (result["label"], result["inflight"], result["payload_bytes"])
        cells[key].extend(latency for latency, _ in observations(result))
        env_load = result.get("env", {}).get("load_average")
        if env_load:
            loads.add(env_load)

    if loads:
        print("recorded load averages:")
        for recorded_load in sorted(loads):
            print(f"  {recorded_load}")
        print()
    else:
        print("recorded load averages: missing; do not use this output as evidence")
        print()

    impls = sorted({name for name, _, _ in cells},
                   key=lambda name: {"std": 0, "swift": 1, "nw": 2}.get(name, 9))
    depths = sorted({depth for _, depth, _ in cells})
    deepest = max(depths)

    print(f"4 KiB p50 latency per message, microseconds ({args.stat} [95% bootstrap interval])")
    print()
    for impl in impls:
        for depth in depths:
            values = cells.get((impl, depth, 4096))
            if values:
                lo, hi = interval([
                    boot_cell(values, random.Random(seed), args.stat)
                    for seed in range(BOOTSTRAPS)
                ])
                print(f"  {impl:5} depth {depth:2}: {fmt(statistic(values, args.stat))} [{fmt(lo)}, {fmt(hi)}]")
    print()

    rng = random.Random(1)
    terms = {impl: term_samples(cells, impl, deepest, rng, args.stat) for impl in impls}
    print("cost model from depth-1 sizes <= 64 KiB")
    print("all values are microseconds/message except bandwidth")
    print()
    print(f"{'impl':6} {'fixed':25} {'bandwidth MB/s':25} {'amortized':25} {'serialization':25}")
    print("-" * 112)
    for impl in impls:
        samples = terms[impl]
        print(f"{impl:6} {summary(samples, 0):25} {summary(samples, 1, 0):25} "
              f"{summary(samples, 2):25} {summary(samples, 3):25}")
    print()

    if all(terms.get(impl) for impl in ("std", "swift", "nw")):
        claims = [
            ("framework amortized work (swift - std)", delta(terms, terms["swift"], terms["std"], 2)),
            ("binding amortized work (nw - swift)", delta(terms, terms["nw"], terms["swift"], 2)),
            ("binding serialization (nw - swift)", delta(terms, terms["nw"], terms["swift"], 3)),
        ]
        print("claim deltas")
        for name, values in claims:
            lo, hi = interval(values)
            print(f"  {name}: {fmt(median(values))} [{fmt(lo)}, {fmt(hi)}] us/msg; {verdict(values)}")


if __name__ == "__main__":
    main()

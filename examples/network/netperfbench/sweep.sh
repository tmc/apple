#!/bin/bash
# sweep.sh measures every transport across payload sizes and in-flight
# depths, then fits a cost model to the result.
#
# The model it is trying to separate:
#
#	latency(size) = fixed + size / bandwidth
#
# A sweep over payload size at depth 1 gives both terms: the intercept is
# what one message costs regardless of how big it is (syscalls, blocks,
# thread handoffs, framework bookkeeping), and the slope is what a byte
# costs (copies, the wire). Re-running at depth N overlaps those fixed
# costs; whatever the per-message time falls to is the part that was
# serialization latency rather than work.
#
# Usage: ./sweep.sh [output_dir]

set -euo pipefail

dir=$(cd "$(dirname "$0")" && pwd)
out=${1:-$(mktemp -d "${TMPDIR:-$HOME/tmp}/netperfsweep.XXXXXX")}
repeat=${NETPERFBENCH_REPEAT:-5}
stat=${NETPERFBENCH_STAT:-median}
mkdir -p "$out"

sizes="64 256 1024 4096 16384 65536 262144 1048576"
depths="1 8 64"
impls="std nw swift"

# Fewer batches at the large sizes and depths: they move the same bytes.
batches() {
	local size=$1 depth=$2
	local bytes=$((size * depth))
	if [ "$bytes" -ge 1048576 ]; then echo 300
	elif [ "$bytes" -ge 65536 ]; then echo 1000
	else echo 3000
	fi
}

echo "building into $out" >&2
go build -o "$out/netperfbench" "$dir"
swiftc -O "$dir/swift/netperfbench.swift" -o "$out/netperfbench-swift"

server_pid=""
cleanup() {
	[ -n "$server_pid" ] && kill "$server_pid" 2>/dev/null
	return 0
}
trap cleanup EXIT

port=52000
for impl in $impls; do
	for depth in $depths; do
		for size in $sizes; do
			port=$((port + 1))
			n=$(batches "$size" "$depth")
			case $impl in
			swift) "$out/netperfbench-swift" -role server -port $port 2>/dev/null & ;;
			*) "$out/netperfbench" -role server -impl "$impl" -port $port 2>/dev/null & ;;
			esac
			server_pid=$!
			sleep 0.4
			case $impl in
			swift) "$out/netperfbench-swift" -role client -addr "127.0.0.1:$port" \
				-size "$size" -inflight "$depth" -n "$n" -repeat "$repeat" -label "$impl" -json ;;
			*) "$out/netperfbench" -role client -impl "$impl" -addr "127.0.0.1:$port" \
				-size "$size" -inflight "$depth" -n "$n" -repeat "$repeat" -label "$impl" -json ;;
			esac >>"$out/raw.jsonl" 2>/dev/null || echo "FAILED $impl size=$size depth=$depth" >&2
			kill "$server_pid" 2>/dev/null || true
			wait "$server_pid" 2>/dev/null || true
			server_pid=""
			echo -n "." >&2
		done
	done
done
echo >&2

# The JSON objects are multi-line, so glue them back together before
# analysis rather than reading the file line by line.
python3 "$dir/analyze.py" --stat "$stat" "$out/raw.jsonl" | tee "$out/report.txt"
echo "raw results in $out/raw.jsonl" >&2

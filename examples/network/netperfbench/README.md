# netperfbench

A TCP echo benchmark written four ways — Network.framework through the Go
bindings, its raw C API, Swift's `NWConnection`, and the Go standard library
`net` package — so the cost of the bindings can be told apart from the cost
of the framework.

All four speak the same wire protocol: the client sends a fixed-size payload
and reads the same number of bytes back, one round trip at a time on one
connection. Because the protocol is shared, any client can be pointed at any
server, and the crossed pairings are what isolate each side:

| pairing | what it measures |
| --- | --- |
| `std` / `std` | plain sockets — the floor |
| `c` / `c` | raw Network.framework C API |
| `nw` / `nw` | Network.framework end to end, through Go |
| `swift` / `swift` | Network.framework end to end, natively |
| `swift` server / `nw` client | the bindings' send and receive path |
| `nw` server / `swift` client | the bindings' listener and echo path |

## Running

The whole matrix, built and tabulated:

```
./run.sh                # 4096-byte payload, 20000 round trips
./run.sh 65536 5000     # larger payload, fewer trips
```

One pairing at a time, in two shells or on two hosts:

```
go run . -role server -impl nw -port 51000
go run . -role client -impl std -addr 127.0.0.1:51000 -size 4096 -n 20000
```

Both roles in one process:

```
go run . -role both -impl nw -n 20000
```

The Swift program takes the same flags:

```
swiftc -O swift/netperfbench.swift -o netperfbench-swift
./netperfbench-swift -role both -n 20000
```

The C program takes the same benchmark flags:

```
clang -O2 -fblocks c/netperfbench.c -framework Network -framework CoreFoundation -o netperfbench-c
./netperfbench-c -role both -n 20000
```

`-json` makes the client print machine-readable results instead of a table;
that is what `run.sh` consumes. Set `NETPERFBENCH_OUT` to keep every
pairing's JSON:

```
NETPERFBENCH_OUT=~/bench/2026-08-06 ./run.sh
```

## As a Go benchmark

For "did this change make round trips slower" rather than "how do the three
implementations compare", `bench_test.go` runs both roles in one process:

```
go test -run '^$' -bench BenchmarkRoundTrip -benchmem -benchtime=2000x -count=6 | tee new.txt
benchstat old.txt new.txt
```

A fixed `-benchtime=Nx` keeps every arm on the same iteration count, so `b.N`
autoscaling is not itself a source of variance between arms. Compare two
trees, not two lines of one run.

`-benchmem` is the reason to run this even when wall time is noisy:
allocations per round trip are load-independent. On a 4 KiB payload the std
transport allocates nothing per round trip and the nw transport allocates
about 2 KiB across ~83 allocations — the blocks, contexts, and channel
traffic each send and receive sets up. That is a real and separable cost even
though the end-to-end matrix shows the framework, not the bindings,
dominating wall time.

The related cost on the dispatch side — getting a Go slice into and out of a
`dispatch_data_t`, which every `nw_connection_send` pays and a socket write
does not — is benchmarked in the `dispatch` package:

```
go test github.com/tmc/apple/dispatch -run '^$' -bench BenchmarkData -benchmem -benchtime=100x -count=6
```

## Trusting a number

A few flags exist because the most expensive benchmark is one that measured
something other than what you thought.

- `-repeat N` runs the measurement N times and reports the **median** run,
  which resists the one repetition that caught a scheduler hiccup in a way a
  mean of means does not. Every run is kept in the JSON.
- `-require-interface en0` and `-forbid-loopback` turn the resolved network
  path into an assertion: the run exits non-zero if Network.framework did not
  use the link you meant to measure. The path is recorded in the JSON either
  way, so a result says which interface it *used*, not which one was asked
  for. Only the `nw` transport can report this.
- The client warns on stderr when the machine's load average is more than
  half its CPU count, and the load average, git commit, dirty-file count,
  OS and kernel version, `GOMAXPROCS`, and hostname all land in the JSON.
  A number found in a log six months later can still be placed.
- `-cpuprofile` and `-memprofile` profile the run you are actually
  measuring, rather than a reconstruction of it under `go test`.
- `-recv-batch` is an experimental receive mode. It asks a framework client
  for one whole in-flight batch, reducing callback re-arming. It changes the
  one-receive-per-echo workload; JSON records `receive_batch` so its results
  cannot be mixed with default-mode results.

The payload is non-constant, so nothing along the path can shortcut it, and
warmup round trips run outside the timer.

## Current cost-model status

The original three-arm subtraction was not sufficient: Swift's `NWConnection`
is a wrapper over the raw `nw_*` entry points that the Go bindings call. The C
client is the reference for those entry points.

In an independent loaded-machine experiment at 4 KiB and depth 64 (load
10.49--18.31 on 16 CPUs), elapsed-time means with unpaired bootstrap intervals
found these terms in microseconds/message:

| term | result |
| --- | ---: |
| framework floor, C - std | +23.26 [21.66, 25.19] |
| Swift wrapper, Swift - C | +3.87 [-0.28, 9.81] |
| Go bindings, nw - C | +16.59 [13.79, 19.24] |

The framework and Go-binding terms are positive. The Swift-wrapper term is
not yet distinguishable from zero. CPU deltas are load-sensitive and must not
be used as fixed model coefficients: paired ordering did not expose a large
latency order effect at load 10--18, but it does not cancel multiplicative CPU
inflation under heavy load.

`-recv-batch` is a deliberately different workload that moves receive
re-arming into Network.framework. In a separately randomized loaded run
(load 14.39--24.42), it reduced nw by 35.82 us/message [33.55, 38.52]; C and
Swift also improved by about 25--27 us/message. Batch-mode nw and C were
indistinguishable. It is the leading optimization candidate, but it must be
adopted only where whole-batch receive semantics are acceptable.

In a five-round four-arm size sweep at depth 64 (load 8.66--22.02), the C -
std framework term was +26.23 us [22.93, 29.62] at 16 KiB and +0.76 us
[-7.97, 10.12] at 64 KiB. The fixed-cost crossover is therefore between 16
and 64 KiB on this loopback system, not a single precisely fitted size.

## Historical example

Latency is round-trip: send plus read-back. Throughput counts both
directions, so one round trip moves `2 * size` bytes. A representative run on
an M-series Mac, loopback, 4 KiB payload:

```
pairing                        p50 us     p90 us     p99 us    trips/sec       MB/s
std server / std client          19.9       31.7       57.0        43673      341.2
nw server / nw client            76.8      100.4      150.8        12284       96.0
swift server / swift client      72.5      102.1      193.8        12394       96.8
swift server / nw client         81.7      105.8      173.5        10852       84.8
nw server / swift client         57.1       78.3      115.7        16457      128.6
```

This table predates the raw-C client and is an example of the output format,
not the current cost model. Re-run the four-arm experiment before drawing a
conclusion from it.

Numbers move with payload size, machine, and load. Re-run the matrix rather
than trusting the table above, and prefer p50 over the mean — the max is
almost always a scheduler artifact. That table was taken on an idle machine;
the same matrix under a load average of 28 on 16 CPUs reported std/std at
69 µs p50 rather than 20 µs, which is the machine being measured, not the
transport.

## Notes

- Every implementation disables Nagle, so the comparison is not measuring
  someone's delayed-ACK behaviour.
- The client's receive asks for a minimum of the full payload length, so a
  completion fires once per round trip rather than once per segment.
- The Go `nw` transport runs its callbacks on a dispatch queue and hands
  results back over channels, keeping the measured loop synchronous — the
  same structure the Swift program gets from `DispatchSemaphore`.
- Loopback exaggerates the framework's relative cost, since there is no wire
  time to hide it behind. Over a real link the gap narrows.

# threeaccel — one matmul, three engines, one Go binary

The same C = A×B float32 matmul dispatched to all three of Apple silicon's
matrix engines from a single Go process:

- **AMX** — Accelerate's `cblas_sgemm`, which Apple routes to the AMX
  coprocessor blocks for this shape.
- **GPU, three effort levels** — a deliberately naive hand-written Metal
  kernel (one thread per output element), a simdgroup-matrix tiled kernel
  (each simdgroup accumulates an 8×8 tile through the hardware matrix
  units), and `MPSMatrixMultiplication` — Apple's tuned GEMM — in fp32 and
  fp16.
- **ANE** — `x/ane/dynamicmatmul`, two ways: the dynamic path that re-sends
  weights on every eval, and the resident path (`PrimeWeightsIO` +
  `EvalCFIOInto`) where weights stay put and only activations move.

Every arm is checksum-validated against a float64 CPU reference before it
is timed. The ANE arm is presence-checked (it rides private frameworks) and
reports "unavailable" rather than failing the demo. After the table, all
three engines run the largest size **concurrently** from three goroutines,
each on its fastest path (sgemm / MPS / ANE-resident).

```
go run ./examples/compute/threeaccel
go run ./examples/compute/threeaccel -sizes 256,512,1024,2048 -window 2s
go run ./examples/compute/threeaccel -sizes 1024,2048 -power   # no sudo needed
```

## Timing methodology

- Every measurement is calibrated to a **200ms floor** (doubling iteration
  counts until the loop runs that long), never a best-of-N single call.
- **GPU columns are kernel-only time** from the command buffer's
  `GPUStartTime`/`GPUEndTime`; wall-clock (which includes submission
  latency and the `WaitUntilCompleted` round trip) is in the notes.
- ANE compile cost is excluded and reported separately, matching how the
  other arms exclude pipeline setup.
- Matrices come from a seeded RNG (`-seed`), so runs are reproducible.

## Measured 2026-08-11 (M-series, macOS 26.x, go1.26.3)

Full sequential run (`-sizes 256,512,1024,2048`):

```
Size          AMX  GPU-naive   GPU-simd    GPU-MPS        ANE    ANE-res
------  ---------  ---------  ---------  ---------  ---------  ---------
256         701.8      178.4      560.4      730.3       46.8      161.5
512         516.1      477.4      730.5     1829.4      117.0      510.0
1024        440.1      675.4     1504.6     2685.9      154.6     1050.9
2048        958.2      446.0      975.0     3440.7      181.5     1311.0

  n=2048 GPU-MPS: ...; fp16 5446.3 GFLOPS kernel-only
  n=2048 ANE-res: hardware-reported eval 3.089ms (5561.6 GFLOPS on-device)

Concurrent: 3 engines at once, n=2048, 2s window (AMX sgemm + GPU MPS + ANE resident)
  AMX      512.2 GFLOPS  (solo 958.2, -47%)
  GPU     2386.8 GFLOPS  (solo 3440.7, -31%)
  ANE     1186.8 GFLOPS  (solo 1311.0, -9%)
  aggregate 4085.8 GFLOPS
```

A fresh single-size run (`-sizes 2048`, cold machine) is faster across the
board — GPU-MPS 4605 fp32 / 5503 fp16, ANE-res 1760 wall (5618 on-device),
AMX 1062 — see "variance" below.

## The max-perf ladder

Same silicon, same math, four effort levels on the GPU alone (n=2048,
kernel-only):

| path | GFLOPS | note |
|---|---|---|
| naive MSL | ~450 | one thread per element |
| simdgroup-tiled MSL | ~1000–1235 | ~40 lines of MSL, 2–3x |
| MPS fp32 | ~3400–4600 | Apple's tuned GEMM, via generated bindings |
| MPS fp16 | ~5450–5500 | the precision the hardware actually wants |

And the ANE's hardware-reported execution at n=2048 is **~5.6 TFLOPS
on-device** — within noise of the GPU's fp16 GEMM, at a fraction of the
power. Peak sustained aggregate with all three engines running
concurrently: **~4.1 TFLOPS** (window-averaged, wall-clock, fp32).

## Reading the numbers honestly

- **Calibrated timing changed every number** vs the earlier best-of-5
  build (AMX n=1024 was reported 2468; sustained it's 440–1260 depending
  on thermal state). Best-of-N catches one boosted-clock call; the 200ms
  floor measures what the engine sustains.
- **Run-to-run variance is real and thermal/ordering-driven.** The AMX
  column swings ~2x between a cold single-size run and the middle of a
  six-arm sweep; GPU-MPS at n=2048 spans 3440–4605. Treat any single cell
  as ±30%; the *ordering* of the ladder is stable even when the values
  wobble.
- **Resident weights are worth ~7x on the ANE** (n=1024: 1051 vs 155
  wall-clock GFLOPS). Most of the previously prose-only wall-vs-on-device
  gap was weight staging; what remains (1311 wall vs 5562 on-device at
  n=2048) is activation IO.
- **The engines contend, unevenly.** Concurrent at n=2048: AMX keeps 53%
  of solo, GPU 69%, ANE 91% — the ANE barely notices, consistent with it
  having its own local memory for resident weights while AMX and GPU fight
  over the shared fabric.
- **fp16 verification passes at 2e-2 relative tolerance** (same as the
  ANE's), versus 1e-4 for the fp32 arms: reduced-precision accumulation
  stated instead of hidden.
- **MPSMatrixMultiplication was supposed to be the "once generated" future
  row** in the design doc — the `metalperformanceshaders` bindings turned
  out to already be generated, so the naive-vs-tuned argument is now a
  measurement, not a plan: 10x at n=2048.
- With `-power`, a dedicated phase runs each engine *alone* for the window
  under `x/powersample` (unprivileged IOReport energy counters — no sudo)
  and charges each arm only its own rail, which is what makes per-engine
  attribution honest: the GPU and ANE rails draw ~0 at idle, so their
  numbers need no baseline. The AMX row rides the CPU rail, which also
  carries every other process — its GFLOPS/W is only meaningful on a
  quiet machine, and the output says so.

## Perf per watt, measured 2026-08-15 (same host, n=1024, 3s window, shared machine)

| engine | GFLOPS | W on its rail | GFLOPS/W |
|---|---|---|---|
| AMX | 1531 | 7.79 (CPU rail, shared-host caveat) | ~197 |
| GPU (MPS) | 2672 | 5.16 | 518 |
| ANE (resident) | 1584 | 1.07 | **1476** |

The ANE is not the fast engine — it is the efficient one: ~3x the GPU's
GFLOPS/W and ~7.5x AMX's at n=1024, at about one watt. This is the
column that decides which engine a battery-powered caller should pick,
and the reason "perf-per-watt is the ANE's entire reason to exist" was
worth wiring.

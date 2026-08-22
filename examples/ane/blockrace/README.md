# blockrace

Runs the same feed forward network on the Neural Engine and on the GPU, one at a
time and then both at once, to measure whether the two engines contend when a
single process drives them together.

The network is a transformer block's feed forward half, `out = W2*relu(W1*x)²`.
The Neural Engine runs it through the private `e5rt` route from a compiled MIL
program; the GPU runs the same arithmetic as an MPSGraph. Both are checked
against the same float64 CPU reference before any timing is reported, and again
after the concurrent phase, so an arm cannot win by being fast and wrong.

```sh
go run ./examples/ane/blockrace
go run ./examples/ane/blockrace -dim 256 -hidden 1024 -seq 64 -window 3s
```

```
feed forward: dim 128, hidden 512, seq 32 — out = W2*relu(W1*x)^2
  backend: the compiler emitted [ane] for the neural engine arm
  neural engine  matches the float64 reference: worst |diff| 0.00066, tolerance 0.0239, reference peak |0.6968|
  gpu            matches the float64 reference: worst |diff| 0.00000, tolerance 0.0239, reference peak |0.6968|

warming up for 1.5s per arm
measuring 5 rounds of solo and concurrent, 1.5s per arm per round

solo, median of 5 rounds
  neural engine   10170.4 runs/s      85.3 GFLOPS  (spread 2% across rounds)
  gpu              3804.9 runs/s      31.9 GFLOPS  (spread 12% across rounds)

concurrent, both engines at once, median of 5 rounds
  neural engine    9939.7 runs/s    -2.3% against its own solo rate
  gpu              3298.5 runs/s   -13.3% against its own solo rate

  together 13238.2 runs/s against 10170.4 for the faster engine alone: 1.30x
  the engines add capacity rather than divide it
```

Two independent runs both landed on **1.30x**, with the Neural Engine giving up
1–2% and the GPU 13–33%.

## What is being compared, and what is not

The engines run at different precisions — the Neural Engine is fp16 throughout
and the MPSGraph arm is fp32 — so their **absolute rates are not comparable**
and this program does not present them as if they were. What is comparable is
what each arm does to itself: each engine's concurrent rate against its own solo
rate is a ratio of like to like, and that is what answers the contention
question.

The aggregate line is the sum of the two concurrent rates against the faster
solo rate. Above 1.0 means the two engines together did more work per second
than the better of them alone.

For a rigorous cross-engine *speed* comparison, including the AMX blocks and
several GPU paths at matched precision, see `examples/compute/threeaccel`.

## Getting the measurement right was most of the work

**The first version reported both engines running faster while contending** —
the Neural Engine +48%, the GPU +113%. That is not a contention result. It
measured both solo rates first and both concurrent rates second, so the solo
phase ran on cold clocks and the concurrent phase on hot ones; the "speedup" was
the ramp.

A warm-up alone did not fix it. A later run still showed +20% and +6.8%. Three
things together make the number mean something:

- **A warm-up window per arm**, discarded.
- **Alternating rounds** — solo, solo, concurrent, repeated — so that any drift
  during the run lands on both phases alike instead of only on the first.
- **The median of several rounds**, with a stability guard.

The guard is the important part. If the solo rates vary by more than a quarter
across rounds, the program **refuses to report a contention figure at all** and
says so, rather than dividing two numbers that were measured under different
conditions. On a busy machine it fires regularly; that is the intended
behavior, and a longer `-window` is the fix.

Spread is measured across the interquartile range rather than min-to-max,
because with a handful of rounds a single round disturbed by unrelated system
activity is the outlier and would otherwise decide the answer. A genuinely
unsettled machine spreads every round, not one.

## Controls

**Placement.** The device mask handed to the compiler is a permission, not a
placement, so the compiled bundle is read for which backend the compiler
actually chose. An arm that quietly fell back to the CPU would return the same
correct answer and would contend with everything.

**Correctness before and after.** Both arms are verified against the float64
reference before timing, and verified again after the concurrent phase — rates
of producing garbage are not rates.

**A tolerance that can fail.** The comparison has an absolute floor, so weights
that left the output near zero would put the whole check inside its own
tolerance, where a badly wrong answer still passes. An earlier version had
exactly that: a reference peaking at 0.0435 against a tolerance of 0.0109. The
weights are now scaled so the result lands near unit magnitude, and the program
refuses to run if the reference peaks below 0.25.

**A fixed window rather than a fixed count**, so that both arms are under load
for the same span during the concurrent phase. A count-based loop would let the
faster arm finish early and leave the slower one running alone, which measures
something else.

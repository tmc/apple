# e5rtblock

Runs a transformer block on the Neural Engine through the private `e5rt`
direct-dispatch route, and checks the result against a float64 CPU reference of
the whole block.

`e5rtdispatch` proves the route runs, using a 1x1 convolution. This runs work
with the shape of a real model: RMSNorm, a QKV projection, multi-head attention
with a causal mask, an output projection and a residual, then a second RMSNorm
and a ReLU-squared feed forward network.

```sh
go run ./examples/ane/e5rtblock
go run ./examples/ane/e5rtblock -dim 256 -heads 8 -seq 64 -hidden 512
```

```
block: dim 64, heads 4, seq 16, hidden 128
  compiled both programs in 337ms
  backend for attention: the compiler emitted [ane]
  backend for feed forward: the compiler emitted [ane]
  encoded both operations into one stream, sharing the intermediate buffer
  attention score spread 1.549, block contributes 1.16x the input magnitude
  dispatch 663µs
  block output matches the float64 reference over 1024 values: worst |diff| 0.0008, tolerance 0.0285, reference peak |0.9254|
  control: perturbing one input element moved 1021 of 1024 outputs, and the new output still matches its reference
```

## The two programs share a buffer

Attention and the feed forward network are separate compiled programs, chained
without a host round trip: the buffer object bound to the attention program's
output port is bound again to the feed forward program's input port, so the
second reads the first's result in place. Both operations go into one execution
stream and one `execute_sync` runs them. Nothing copies the intermediate
activation and the host never sees it.

## What this found

**`GenFFNForwardRMSReLU2` has no residual add.** It computes
`W2(relu(W1(rms_norm(x)))²)` and stops. A caller who chains the two generators
and takes the second program's output as the block output gets a transformer
block missing one of its two residual connections — silently, with
plausible-looking numbers. This program does that add on the host and says so.

**Neither generator had ever been checked for what it computes.** Before this,
`GenSDPAForward` and `GenFFNForwardRMSReLU2` appeared only in
`x/ane/mil/mil_compile_test.go`, which checks that they compile. This is the
first numerical verification of either.

## Why the weights are what they are

The first working version of this program passed, and verified almost nothing.
Its weights came from a smooth sine, and correlated weights cancel under a dot
product over hundreds of terms, so the attention scores came out with a spread
of **0.004** across the whole sequence. A softmax that uniform makes attention
an average over positions — and then a wrong head split, a wrong query-key
transpose and a missing causal mask all produce the same numbers. The block was
also dominated by its own residual, so the output was mostly just the input.

It matched the reference to within 0.0004 and meant nothing.

The weights are now decorrelated, and scaled per stage so the query and key
projections land near unit magnitude — a projection sums over its input width,
so a weight of scale `s` gives roughly `s*sqrt(width)`. `checkDiscriminating`
refuses to report a pass unless the score spread is at least 1 and the block
moves its input by at least a quarter of the input's own magnitude, so this
cannot silently regress.

## Controls

The device mask is a permission, not a placement, so the compiled bundle is read
for which backend the compiler actually chose; a block that fell back to the CPU
would return the same correct answer. One input element is then perturbed and
the block re-dispatched **without re-encoding**, which must move the output and
must still match a reference recomputed from the perturbed input — a program
bound to a stale buffer, or returning a constant, passes the first check and
fails this one.

Four deliberate mutations were each run against the real engine output and each
was caught, with the tolerance at roughly 0.03:

| mutation | worst \|diff\| |
| --- | --- |
| wrong head split (`c = d*heads + h`) | 0.0889 |
| swapped query/key indices | 0.0797 |
| reference not causal | 0.4900 |
| second residual dropped | 0.9048 |

The residual mutation had to be written to still compile: the obvious edit left
a variable unused, and a mutant the compiler rejects is caught by the compiler
rather than by the test, which proves nothing about the test.

## Precision

Everything on the engine is fp16 and the block accumulates over four matrix
products and a softmax, so the comparison is to an fp16-scale tolerance. The
program prints the worst difference next to the peak magnitude of the reference,
so a small difference against a small reference is not read as agreement. In
practice the agreement is far tighter than the tolerance: worst `|diff|` around
0.001 against a reference peak near 1.0, holding from dim 64 through dim 256.

# modelchain

Runs a transformer block on the Neural Engine through the **public** `x/ane`
API, chaining two compiled models by sharing an IOSurface, and checks the result
against a float64 CPU reference.

The other ANE examples all drive the private `e5rt` route. This one uses what an
outside caller of the package has: `ane.Probe`, `ane.Open`, `Client.Compile`,
`ane.ShareSurface`, and `Model.Eval`.

```sh
go run ./examples/ane/modelchain
go run ./examples/ane/modelchain -dim 128 -heads 8 -seq 32 -hidden 256
```

```
macOS, h16g, 16 ANE core(s)

block: dim 64, heads 4, seq 16, hidden 128
  attention: 1 input(s) 1 output(s), output "out@output" is 64 channels over 16 positions
  feed forward: 1 input(s) 1 output(s)
  attention score spread 1.549, block contributes 1.16x the input magnitude
  unchained: worst |diff| 0.3833 against a tolerance of 0.0285, as it must be
  chained: the feed forward input is now the attention output surface
  block output matches the float64 reference over 1024 values: worst |diff| 0.0008, tolerance 0.0285, reference peak |0.9254|
  perturbed block output matches the float64 reference over 1024 values: worst |diff| 0.0008, tolerance 0.0290, reference peak |0.9505|
  control: perturbing one input element moved 1021 of 1024 outputs, and the new output still matches its reference
```

## What this found: ShareSurface did not reach Eval

Writing this example turned up a bug in the public API, now fixed.

`Compile` gives every model its own IOSurfaces and builds an `ANERequest` from
`ANEIOSurfaceObject` wrappers **at compile time**. `Eval` runs that request.
`ShareSurface` assigned `dst.inputs[dstInput]`, which is the Go slice consulted
by `WriteInput` and `ReadOutput` — but not by the engine. So the second model
kept evaluating against the surface it was compiled with, and the documented
zero-copy handoff never happened.

Nothing structural would have caught it. After the call `InputSurface(0)` really
did report the shared surface, so a check that the two models name one surface
passed while the chain did not work. What caught it was running the block
**before** chaining and requiring it to fail: the chained and unchained results
came back byte-identical, which they cannot be if sharing does anything.

The fix rebuilds the destination model's request from its current surfaces.
`x/ane`'s `TestShareSurfaceReachesEval` covers it, with a mutation control
confirming the test fails against the old behavior (it reads back all zeros).

## Controls

The negative control described above is the load-bearing one, and it stays in
the program: the block is run once before chaining, where it must **not** match
the reference, and once after, where it must. Asserting only that the surfaces
are now the same object would show that the call assigned something, not that
the assignment is what carries the activation.

The weights are decorrelated so that the softmax selects rather than averages.
Attention whose scores are all equal is an average over positions, and a wrong
head split, a wrong query-key transpose and a missing causal mask all reproduce
that exactly — which is how an earlier version of the `e5rtblock` example passed
while verifying nothing. `checkDiscriminating` refuses to report a pass unless
the score spread is at least 1 and the block moves its input by at least a
quarter of the input's magnitude.

One input element is then perturbed and the block re-evaluated, which must move
the output and must still match a reference recomputed from the perturbed input.

## The residual the generator does not add

`GenFFNForwardRMSReLU2` computes `W2(relu(W1(rms_norm(x)))²)` and stops — no
residual add. The block output is the attention result plus the feed forward
result, and this program does that last add on the host. The `e5rt` examples
close it with a third program on the device; through the public API there are
two models and two `Eval` calls, so the host is already in the loop between
them.

## Precision

Everything on the engine is fp16 and the block accumulates over four matrix
products and a softmax, so the comparison is to an fp16-scale tolerance. The
worst difference is printed next to the peak magnitude of the reference, so a
small difference against a small reference is not read as agreement.

# pipelinelogits

Runs a language model's logits tail — final RMSNorm, classifier projection,
softmax — as one linked `e5rt.Pipeline` on the Neural Engine, and checks every
probability against a float64 CPU reference.

The other `e5rt` examples encode a stream by hand. This one drives the package's
multi-stage interface: three independently compiled MIL programs, two links, one
`Execute`. A linked pair of ports is bound to one E5RT buffer object, so this
program stages no host copy between the stages. That is what is measured — not
device-internal movement, and not a claim that the platform performs no hidden
copy of its own.

```sh
go run ./examples/ane/pipelinelogits
go run ./examples/ane/pipelinelogits -dim 128 -vocab 256 -seq 32
```

```
logits tail: dim=64 vocab=32 seq=8

  backend for norm: the compiler emitted [ane]
  backend for cls: the compiler emitted [ane]
  backend for smax: the compiler emitted [ane]

  linked pipeline vs float64 CPU reference: max diff 5.75e-05 (tolerance 0.005)
  stage 0 "out" -> stage 1 "x": one Go buffer = true (by construction, not a test)
  stage 1 "out" -> stage 2 "x": one Go buffer = true (by construction, not a test)
  link 0 omitted: max diff 0.0299, well outside the 0.005 tolerance — the link is load-bearing
  link 1 omitted: max diff 0.0299, well outside the 0.005 tolerance — the link is load-bearing
  host-staged arm vs linked pipeline: max diff 0
  mutation control: perturbing one input moved the output by 0.00397 (must exceed 0.000575, ten times the measured disagreement)

  refused, as it must be: link between ports of different sizes
    e5rt: link from stage 0 port "out" (1024 bytes) to stage 2 port "x" (512 bytes)
  refused, as it must be: link that runs backwards
    e5rt: link from stage 2 port "out" (512 bytes) to stage 1 port "x" (1024 bytes)

OK
```

## What the checks are for

A demo that only reported "it ran" would pass on a pipeline that ignored its
input, on links that silently copied, and on a validator that accepted anything.
Each check below exists to close one of those.

| check | what it would catch |
| --- | --- |
| every probability against a float64 CPU evaluation of the whole tail | a stage running the wrong arithmetic |
| the same three programs run separately with host copies between them | links that did not carry what the host copies carried |
| **each link removed in turn, requiring the answer to change** | a link that carries nothing |
| perturbing one input element, and requiring the moved output to match its *own* recomputed reference | a stale buffer, and separately, a pipeline that reacts to the input incorrectly |
| a wrong-size link and a backward link | a validator that checks nothing |

### The check that was removed, and why

An earlier version wrote a marker byte through stage 0's output and read it back
through stage 1's input, and called that an aliasing control. **It was a
tautology.** `CompilePipeline` hands the consumer port the producer's own
`*Buffer` — `data: source.data` at `pipeline.go:418` — so both handles are the
same Go object by construction, whatever the native bind did with them. The
probe could not produce a negative on any implementation, including a broken
one.

Port identity is still printed, labelled as API state rather than as a test. The
control that replaced it removes each link and requires the result to move
outside tolerance, which exercises the native binding rather than a Go handle.
Both links are now covered; the byte probe only ever looked at the first.

Two independent reviews caught this, one of them by reading `bindExisting`.

The two arms agree exactly (`max diff 0`) because they are the same fp16
programs on the same hardware; only the float64 reference is a different
computation, and it is the one that disagrees, by 5.75e-05.

The comparator rejects NaN and unequal lengths. An earlier version did neither,
and an all-NaN result reported a maximum difference of **zero** — `NaN > max` is
false, so the running maximum never moved. Infinity fails loudly; NaN was the
silent one, and NaN is what an fp16 overflow here would produce.

The mutation bar is ten times the *measured* disagreement, not the agreement
tolerance. Those are different quantities. Perturbing one element moves one
channel at one position, which the softmax over the vocabulary then dilutes, so
the move lands under a tolerance sized for the fp16 error of the whole chain — a
control keyed to the tolerance reports that dilution as a defect. It did, on the
first run of this program.

## Why a logits tail

It is the part of a decoder that a host round trip hurts most, and no other
example covers it. It is also the smallest interesting chain whose stages do not
commute, so a program that ran them out of order or dropped one could not match
the reference by luck.

The three programs come from `x/ane/mil`'s `GenFinalRMSNorm`,
`GenClassifierForward` and `GenSoftmaxVocab`. None of the three had a caller
anywhere in the module before this example — they had never been compiled, let
alone run.

## Limits

The device mask is a permission, not a placement, so the emitted
`main_<backend>` directory is the only local evidence the work reached the
engine. Both refusals are predicates in the same `validatePipelineOptions` call, so they
are two validation controls rather than two independent entry points.

`MaxPipelineFunctions` is 15 distinct compiled functions per stream, and
the pool is shared across processes, so a pipeline that fits may still fail
while another process holds slots.

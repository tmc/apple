# e5rtdispatch

`e5rtdispatch` compiles a MIL program and runs it on the Apple Neural Engine
through Espresso's private `e5rt_*` direct dispatch route, then checks the
answer against a CPU reference.

The route does not go through Core ML. A model directory holding `model.mil`
and its weight files is handed to Espresso's compiler, which lowers it to a
signed hardware program inside the `aned` daemon; the result is bound to buffer
objects and submitted on an execution stream.

```sh
go run ./examples/ane/e5rtdispatch
go run ./examples/ane/e5rtdispatch -in 64 -out 64 -spatial 32
go run ./examples/ane/e5rtdispatch -device cpu
go run ./examples/ane/e5rtdispatch -cache ~/tmp/bundles   # keep compiled bundles
```

## What it runs

Two programs, both 1×1 convolutions, whose results are cheap to predict exactly.

1. **One weight file.** The layout every MIL template in this repository emits:
   a single weight tensor in its own file, referenced at offset 64.
2. **Two weights, one file.** Two chained convolutions whose weight tensors are
   packed into a single blob at two different offsets — the layout real models
   use, and the one `x/ane/mil.BlobWriter` produces. This repository had never
   submitted a multi-entry blob to the compiler before this example.

   Its default of 12 channels is deliberate. Blob Storage v2 rounds each entry
   up to 64 bytes, and a naive packed layout would put the next descriptor
   immediately after the previous payload; the two rules agree whenever every
   payload is a multiple of 64, which most tensors are. At 12 channels the first
   tensor is 288 bytes and the rules disagree — 448 against 416 — so the file
   actually exercises the rounding. At the 16 channels used elsewhere it would
   not.

   What that establishes is that the compiler *accepts* the rounded-up
   placement and reads the right tensor back from it. It does not establish that
   a packed file would be *rejected*, and in fact it is not:
   `TestBlobPackedVersusAlignedLayout` in `x/ane/mil` compiles this same shape
   from a deliberately packed file with matching packed offsets, and it runs and
   matches its reference too. The compiler follows whatever offsets the MIL text
   gives as long as the file agrees with them.

   Alignment still matters when *reading* a file this repository did not write.
   Every Apple-produced blob file examined here is aligned, so a parser using
   packed arithmetic reads them wrong.

## What a run establishes

Three things, each with a control, because none of them is safe to assume from
a call that merely returned zero.

**The program ran.** The output is compared against a float64 CPU reference.
A route that returned zeros, or ran a different program, fails the comparison.
The reference's peak magnitude is printed alongside the worst difference, so a
reported difference of zero can be read against a reference that is not itself
zero.

**The engine read our input.** One input element is changed and the program is
dispatched again. The output has to move, and the moved output has to match its
own reference. A constant result, a cached result, or a stale buffer passes the
first check and fails this one.

**The work was placed on the engine.** The device mask is a permission, not a
placement: setting the Neural Engine bit does not stop the planner falling back.
The compiled bundle names the choice, in a directory per selected backend:

| `-device` | directory | payload |
| --- | --- | --- |
| `ane` | `main_ane` | `model.anehash` |
| `cpu` | `main_bnns` | `bnns_program.bnnsir` |
| `gpu` | `main_mps_graph` | `main_mps_graph.mpsgraphpackage` |

All three were produced by compiling this same program under each flag, so the
reading is anchored at both ends rather than assumed. Where no bundle is found
the tool prints `UNKNOWN` rather than guessing.

ANEForge documents a different oracle for this — a `SelectedBackend` entry in
the bundle's `analytics.mil`. No `analytics.mil` appears in the bundles produced
here, so that form is either older or written only under a compiler option this
program does not set.

## Observed

On macOS 26.x, an M-series host, at the default sizes:

```
program 1: 1x1 conv, 16->16 channels, 8 spatial, one weight file
  backend: the compiler emitted [ane] for "main"
  compile 56ms, bind 161µs, first dispatch 273µs, 10 further dispatches 121µs each
  output matches the CPU reference over 128 elements: worst |diff| 0, tolerance 0.02875, reference peak |2.875|
  control: perturbing one input element moved the output by 3 and the new output still matches its reference
```

Compile dominates and happens once; dispatch is roughly 120 µs. The same
program under `-device gpu` takes about 57 ms on its first dispatch, which is a
second, independent signal that the backend really did change.

## Status of the binding

The `e5rt_*` argument lists are not documented by Apple. They come from
ANEForge, the source paper's author's own working implementation, and are cited
per wrapper in `x/ane/e5rt`. A wrong signature faults in C, where the fault is
fatal and `recover` cannot see it, so a crash here would be a real result.

This example is what took the package from "every symbol resolves" to "the route
runs". The calls it does not make are covered by `TestUnverifiedCalls` in
`x/ane/e5rt`, which runs each in its own process because two of them abort.
Every wrapper in the package has now been called from Go.

Four disagreements with ANEForge came out of that, each recorded at its
wrapper. `ExecutionStreamReset` is said to reject a never-executed stream and
does not. `PrepareOpForEncode` is said to reject a never-encoded operation and
returns zero instead — its error really exists, but it surfaces as an uncaught
C++ exception when the stream is later released, which no Go code can catch.
And `ProgramFunctionLoadForExecution`, whose place in the sequence was an open
question, turns out to have been removed from the runtime altogether.

The fourth came from the event family, which is wrapped now that
`SubmitAsync` exists. Building the dependency graph works — a completion event
binds, a second operation takes it as a dependency, the pair dispatches, and a
cycle is refused — but nothing here makes an event's value move, on either
submit path. ANEForge reports that events advance under `submit_async`; that is
not reproduced. Nor is the opposite claimed, because `AsyncEventSyncWait`
returns immediately on an event whose work was deliberately never submitted, so
the reader is inert and a zero from it measures nothing.

For symbol resolution reporting without calling anything, see
[`e5rtprobe`](../e5rtprobe).

// Package engineattest turns "this ran on the ANE" or "this ran on the
// GPU" from a comment into a checked runtime assertion.
//
// The failure mode it guards against is hollow code: a function that
// claims to dispatch work to an accelerator, benchmarks well, and never
// touches the silicon. Timing cannot catch that impostor — a no-op is
// fast — so the assertion reads hardware-owned evidence instead:
//
//   - [ANE] attaches a zeroed performance-stats object to the model's
//     request, runs the attested function, and fails unless the driver
//     reported nonzero hardware execution time. Before trusting a zero
//     it proves the counter can move at all by running one extra
//     evaluation of the model (the sensitivity canary).
//
//   - [Queue.GPU] tracks the command buffers a wrapped
//     [metal.MTLCommandQueue] vends during the attested function and
//     fails unless at least one completed on the GPU with a nonzero
//     scheduling window (GPUEndTime > GPUStartTime > 0).
//
// Basic use:
//
//	if err := engineattest.ANE(k.Model(), func() error { return k.Eval() }); err != nil {
//		log.Fatalf("ANE claim did not survive: %v", err)
//	}
//
// Limits, stated rather than implied: the ANE probe needs a model that
// reports hardware execution time. In-memory MIL models expose no
// hardware counters and fail the canary with [ErrUnattestable];
// package-backed models can report it, but on some OS/firmware versions
// (observed on macOS 26.6) the driver leaves every perf counter zero
// for them too, and the canary likewise returns [ErrUnattestable]
// rather than a false verdict. The canary runs one real
// evaluation, which overwrites the model's output surfaces from its
// current inputs. The GPU probe attests that work went through the
// wrapped queue — code that submits real work through a different queue
// fails the assertion, and code that commits empty command buffers on
// purpose can fool it; the threat model is hollow code, not adversarial
// code.
package engineattest

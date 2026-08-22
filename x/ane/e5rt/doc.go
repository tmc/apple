// Package e5rt binds the e5rt_* direct dispatch route exported by the private
// Espresso framework.
//
// The route reaches the Neural Engine without Core ML: build a network
// description, compile it to the engine's program format, load the compiled
// program, bind a buffer object to each named I/O port, then encode the loaded
// program into an execution stream and submit it. Compile and load run once;
// bind and dispatch run in the hot loop.
// [Compile] packages that established single-function route for callers that
// need host-visible input and output buffers.
//
// # Status of this binding
//
// Every symbol named here has been observed to resolve with dlsym against
// /System/Library/PrivateFrameworks/Espresso.framework/Espresso on macOS 26.x,
// with a negative control proving the probe can report a miss.
//
// The whole route has now been driven from this package on macOS 26.x: a MIL
// program compiled, its function retained, buffer objects bound to its ports,
// and the operation encoded and submitted, with the output matching a float64
// CPU reference exactly over 128 elements. Two controls back that up. Perturbing
// one input element moved the output and the new output matched its own
// reference, so the engine is reading the bound buffer rather than returning a
// constant or a stale result; and the compiled bundle records which backend was
// chosen, which reads ane under a Neural Engine mask and bnns or mps_graph under
// the other two, so the work is placed on the engine and not silently fallen
// back. See the e5rtdispatch example.
//
// That covers only the wrappers that route uses. The rest have since been driven
// too, by TestUnverifiedCalls, which runs each in its own process because two of
// them abort:
//
//   - [Lib.ProgramLibraryCreate] works. A compiled bundle can be reopened
//     directly, with no compiler.
//   - [Lib.ProgramFunctionLoadForExecution] is gone, returning status 2 and
//     saying so.
//   - [Lib.PrepareOpForEncode] returns zero everywhere and leaves the stream
//     unusable and unreleasable. Do not call it.
//   - [Lib.ExecutionStreamReset] works on a fresh stream, contradicting ANEForge.
//   - [Lib.SubmitAsync] works. It needs a real Objective-C block, which
//     [objc.NewBlock] supplies, and the completion block runs.
//
// The event family has since been wrapped and driven too. Building the graph
// works: a completion event binds to an
// operation, that event binds as a second operation's dependency, and the pair
// encodes and dispatches correctly, while an operation given its own event is
// refused. [Lib.AsyncEventSignal] moves an event to a caller-supplied value;
// its one-argument predecessor was wrong. [Lib.AsyncEventSyncWait] still
// returns immediately on an unreached event and is not a barrier. Completion
// events under submit_async are probed separately; use [Lib.SubmitAsync]'s
// completion function for blocking completion.
//
// A zero status means only that an E5RT entry point accepted the call. It is
// not, by itself, evidence that the operation had its apparent effect:
// [Lib.PrepareOpForEncode] and [Lib.AsyncEventSetActiveFutureValue] both return
// zero while failing to do so. Treat an effect as established only where the
// wrapper documentation names a control that observed it.
//
// Every name in [Symbols] now has a wrapper that has been called, save
// e5rt_e5_compiler_is_new_compile_required, which is listed so the probe reports
// on it but has no argument list anywhere consulted. Being called is not the
// same as being understood: [Lib.CompilerOptionsSetCustomANECompilerOptions]
// runs and returns zero without any observable effect, and the event family is
// as described above. A wrapper's doc comment names the evidence behind it and
// says where that evidence runs out; where it cites only ANEForge, that is
// outside evidence and nothing more.
//
// Every argument list here came from one of two outside sources, and each doc
// comment says which:
//
//   - The paper cited below.
//   - ANEForge (github.com/sbryngelson/ANEForge), the paper author's own working
//     reference implementation, whose aneforge/_lib/e5rt_api.h carries recovered
//     signatures and aneforge/_lib/ane_e5rt_dispatch.mm drives the whole route.
//     A call site there outranks a declaration, and both outrank the paper.
//
// Where a citation survives it records provenance, which is not the same as
// evidence. A signature this package has since driven is confirmed by that run,
// and the doc comment says so; the citation stays only to record where the
// argument list came from. A signature that has never been exercised, and every
// claim about behavior that no test here reproduces, has nothing behind it but
// the citation, and those comments say that too. Read a bare ANEForge reference
// as the second case unless the comment states an observation.
//
// Symbols this package will not guess at have no typed wrapper; reach them
// through [Lib.Sym] or [Lib.Lookup] and supply your own convention.
//
// # Stability
//
// These interfaces are private and undocumented. They require no entitlement,
// but that means only that the path is reachable from ordinary user space, not
// that it is supported. Expect it to change across operating-system updates and
// do not ship it in App Store software.
//
// # Performance
//
// The source paper reports (section 6.4) that this unentitled direct path is
// performance-complete for dispatch: a submission driving several on-engine
// steps runs at the same rate as the same steps issued one host call at a time,
// so it reaches the same throughput as any more-privileged path. That claim is
// the paper's and has not been reproduced here.
//
// Sources: Bryngelson, arXiv:2606.22283, chapter 6; and ANEForge, at
// github.com/sbryngelson/ANEForge.
package e5rt

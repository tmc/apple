// Package e5rt binds the e5rt_* direct dispatch route exported by the private
// Espresso framework.
//
// The route reaches the Neural Engine without Core ML: build a network
// description, compile it to the engine's program format, load the compiled
// program, bind a buffer object to each named I/O port, then encode the loaded
// program into an execution stream and submit it. Compile and load run once;
// bind and dispatch run in the hot loop.
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
// That covers only the wrappers that route uses. Everything the example does not
// reach — [Lib.ProgramLibraryCreate], [Lib.ProgramFunctionLoadForExecution],
// [Lib.PrepareOpForEncode], [Lib.ExecutionStreamReset] and the asynchronous
// submission and event calls — still rests on ANEForge alone and has never been
// called from Go. A wrapper's doc comment names the evidence behind it; where it
// cites only ANEForge, that is outside evidence and nothing more.
//
// Every argument list here comes from one of two outside sources, and each doc
// comment says which:
//
//   - The paper cited below.
//   - ANEForge (github.com/sbryngelson/ANEForge), the paper author's own working
//     reference implementation, whose aneforge/_lib/e5rt_api.h carries recovered
//     signatures and aneforge/_lib/ane_e5rt_dispatch.mm drives the whole route.
//     A call site there outranks a declaration, and both outrank the paper.
//
// A signature confirmed by ANEForge is strong outside evidence, not local
// verification. Symbols this package will not guess at have no typed wrapper;
// reach them through [Lib.Sym] or [Lib.Lookup] and supply your own convention.
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

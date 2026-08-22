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
// with a negative control proving the probe can report a miss. Resolution is
// all that has been verified. No function in this package has been called, so
// the argument lists come from the paper cited below and are marked UNVERIFIED
// in the doc comment of each wrapper that carries one. Symbols whose argument
// list the paper does not give, or gives inconsistently, have no typed wrapper
// at all; reach them through [Lib.Sym] and supply your own convention.
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
// Source: Bryngelson, arXiv:2606.22283, chapter 6.
package e5rt

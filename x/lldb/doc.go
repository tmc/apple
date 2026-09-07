// Package lldb provides a minimal, cgo-free binding to the LLDB SB (scripting
// bridge) API. It loads LLDB.framework at runtime with purego and calls the
// public C++ SB API directly.
//
// LLDB has no stable public C API; only the C++ SB API is exported. This package
// binds the small subset needed to attach to a process, set symbol breakpoints,
// inspect stopped threads and registers, evaluate expressions in the target, and
// patch register values. It is not a complete binding.
//
// SB wrapper zero values are not usable. Start with Create and obtain other
// values through its debugger, targets, and processes.
//
// # Platform
//
// macOS only (arm64 and amd64). Every SB value is a small C++ object returned by
// value; under the platform C++ ABI that return goes through an indirect-result
// buffer, so the Go binding models each SB value as a fixed 32-byte cell to route
// the return correctly.
//
// # Lifetime
//
// Each returned SB wrapper owns a native object and must be closed, including
// invalid results and Error values. Copies of a wrapper share ownership: closing
// any copy invalidates them all. Close is safe on zero values and repeated calls.
// Other methods require an open handle; Close must not race with those methods.
// There is no finalizer: Go garbage collection does not release native objects.
//
// Close transient handles at the end of each operation. Detach or kill processes
// before closing their handles, and close the debugger last. Debugger.Close also
// removes LLDB's global debugger registration. Closing a target or breakpoint
// handle does not remove the target or breakpoint from its debugger.
//
// # Usage
//
//	d, err := lldb.Create()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer d.Close()
//	t := d.CreateTargetWithArch("/bin/sleep", "")
//	defer t.Close()
//	bp := t.BreakpointCreateByName("nanosleep")
//	defer bp.Close()
//	fmt.Println(bp.IsValid())
package lldb

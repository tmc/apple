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
// SB values wrap reference-counted smart pointers. This binding moves them by raw
// copy and never runs their C++ destructors, so native memory usage can grow
// with each API call. Go garbage collection does not release these native
// objects. Long-running programs must account for this limitation.
//
// # Usage
//
//	d, err := lldb.Create()
//	if err != nil {
//		log.Fatal(err)
//	}
//	d.SetAsync(false)
//	t := d.CreateTargetWithArch("/bin/sleep", "")
//	t.BreakpointCreateByName("nanosleep")
//	proc := t.LaunchSimple([]string{"30"}, nil, "")
//	for i := uint32(0); i < proc.NumThreads(); i++ {
//		th := proc.ThreadAtIndex(i)
//		if th.StopReason() == lldb.StopReasonBreakpoint {
//			fmt.Println(th.FrameAtIndex(0).FunctionName())
//		}
//	}
package lldb

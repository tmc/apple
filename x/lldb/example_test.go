//go:build darwin

package lldb_test

import (
	"fmt"

	"github.com/tmc/apple/x/lldb"
)

func Example() {
	d, err := lldb.Create()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer d.Close()
	target := d.CreateTargetWithArch("/bin/sleep", "")
	defer target.Close()
	bp := target.BreakpointCreateByName("nanosleep")
	defer bp.Close()
	fmt.Println(bp.IsValid())
	// Output: true
}

func ExampleState_String() {
	fmt.Println(lldb.StateStopped)
	// Output: stopped
}

func ExampleDebugger_Close() {
	d, err := lldb.Create()
	if err != nil {
		fmt.Println(err)
		return
	}
	alias := d
	d.Close()
	alias.Close() // Copies share the same lifetime.
	fmt.Println("closed")
	// Output: closed
}

// Close is safe even when a handle has not been initialized.
func ExampleTarget_Close() {
	var target lldb.Target
	target.Close()
	// Output:
}

func ExampleListener_Close() {
	var listener lldb.Listener
	listener.Close()
	// Output:
}

func ExampleBreakpoint_Close() {
	var breakpoint lldb.Breakpoint
	breakpoint.Close()
	// Output:
}

func ExampleProcess_Close() {
	var process lldb.Process
	process.Close()
	// Output:
}

func ExampleThread_Close() {
	var thread lldb.Thread
	thread.Close()
	// Output:
}

func ExampleFrame_Close() {
	var frame lldb.Frame
	frame.Close()
	// Output:
}

func ExampleValue_Close() {
	var value lldb.Value
	value.Close()
	// Output:
}

func ExampleError_Close() {
	var err lldb.Error
	err.Close()
	// Output:
}

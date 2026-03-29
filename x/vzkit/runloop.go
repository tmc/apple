package vzkit

import vmruntime "github.com/tmc/apple/x/vzkit/vm"

// RunRunLoopOnce runs the main CFRunLoop briefly to process pending callbacks.
func RunRunLoopOnce() { vmruntime.RunLoopOnce() }

// RunRunLoopUntilDone runs the main run loop until done returns true.
func RunRunLoopUntilDone(done func() bool, progress func()) {
	vmruntime.RunLoopUntilDone(done, progress)
}

// RunRunLoopAggressively pumps both CFRunLoop and NSRunLoop multiple times.
func RunRunLoopAggressively() { vmruntime.RunLoopAggressively() }

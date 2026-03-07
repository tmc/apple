// Exception-handler demonstrates Objective-C exception handling from Go.
//
// The objc package automatically installs a default exception handler that
// prints diagnostic information (exception name, reason, call stacks) and
// exits cleanly instead of triggering the ObjC runtime's abort().
//
// This example shows how to register a custom handler and preprocessor.
package main

import (
	"fmt"
	"os"

	"github.com/tmc/apple/objc"
)

func main() {
	// Show that the default handler is already active.
	fmt.Println("Default exception handler is active (installed by objc.init).")
	fmt.Println("Set OBJC_EXCEPTION_HANDLER=off to disable it.")
	fmt.Println()

	// Register a custom uncaught exception handler.
	// This replaces the default colorized output with our own format.
	objc.SetUncaughtExceptionHandler(func(name, reason string, callStack []string) {
		fmt.Fprintf(os.Stderr, "Exception: %s — %s\n", name, reason)
		for i, frame := range callStack {
			fmt.Fprintf(os.Stderr, "  [%d] %s\n", i, frame)
		}
		os.Exit(1)
	})
	fmt.Println("Custom uncaught exception handler registered.")

	// Register a preprocessor that fires before any exception is thrown.
	// This is the earliest interception point.
	objc.SetExceptionPreprocessor(func(exc objc.ID) objc.ID {
		info := objc.GetExceptionInfo(exc)
		if info != nil {
			fmt.Fprintf(os.Stderr, "Preprocessor: about to throw %s: %s\n", info.Name, info.Reason)
		}
		return exc
	})
	fmt.Println("Exception preprocessor registered.")
	fmt.Println()

	// Trigger a real ObjC exception by sending an unrecognized selector.
	// This will invoke the preprocessor, then the uncaught handler, then exit.
	fmt.Println("Sending unrecognized selector to trigger an exception...")
	obj := objc.Send[objc.ID](objc.ID(objc.GetClass("NSObject")), objc.Sel("new"))
	objc.Send[objc.ID](obj, objc.Sel("doesNotExistMethod"))

	// Not reached — the exception handler calls os.Exit(1).
	fmt.Println("This line is never printed.")
}

// Exception-handler demonstrates observing Objective-C exceptions from Go.
//
// An Objective-C exception raised through a Go-initiated message send is
// always fatal. On 64-bit Apple platforms the runtime unwinds using DWARF
// tables that Go frames do not carry, so the search for a handler stops at the
// first Go frame and the process aborts. recover cannot intercept the result.
// The hooks below therefore observe an exception on its way out -- reading its
// name, reason and call stack, writing a report -- and cannot resume
// execution.
//
// Validate before calling instead: check RespondsToSelector, bounds and nil.
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
	fmt.Println("It reports uncaught exceptions and hides the native frames it")
	fmt.Println("cannot explain. Two environment variables widen that:")
	fmt.Println("  OBJC_EXCEPTION_STACK=full     every frame, plus the goroutine dump")
	fmt.Println("  OBJC_EXCEPTION_HANDLER=off    no reporting; the runtime's own output")
	fmt.Println()

	// Register a custom uncaught exception handler, replacing the default
	// report with our own format. The process exits once it returns, so there
	// is no need to exit from inside it.
	objc.SetUncaughtExceptionHandler(func(name, reason string, callStack []string) {
		fmt.Fprintf(os.Stderr, "Exception: %s — %s\n", name, reason)
		for i, frame := range callStack {
			fmt.Fprintf(os.Stderr, "  [%d] %s\n", i, frame)
		}
	})
	fmt.Println("Custom uncaught exception handler registered.")

	// Register a preprocessor. This fires earlier than the handler above --
	// before the throw, and for every exception, including ones that
	// Objective-C code goes on to catch.
	//
	// It is not installed by default, and is worth avoiding unless you need
	// it. CoreFoundation occupies this slot, and that is what records an
	// exception's call stack; displacing it loses the frames that identify the
	// failing call.
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

	// Not reached: the process exits once the uncaught handler returns.
	fmt.Println("This line is never printed.")
}

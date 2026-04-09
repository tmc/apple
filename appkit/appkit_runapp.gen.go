// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"runtime"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// RunApp initializes the shared NSApplication, creates a delegate via the
// delegate builder, and enters the main event loop. setupFn is called once
// the application finishes launching.
func RunApp(setupFn func(app NSApplication, delegate NSApplicationDelegateObject)) {
	runtime.LockOSThread()

	objc.SetupExceptionHandler(objc.ExceptionHandlerConfig{
		LogExceptions: true,
		OnException: func(exc *objc.ObjCException) {
			fmt.Printf("ObjC exception in RunApp: %s — %s\n", exc.Name, exc.Reason)
		},
	})

	app := GetNSApplicationClass().SharedApplication()
	app.SetActivationPolicy(NSApplicationActivationPolicyRegular)

	var delegateObj NSApplicationDelegateObject
	delegateObj = NewNSApplicationDelegate(NSApplicationDelegateConfig{
		DidFinishLaunching: func(_ foundation.NSNotification) {
			setupFn(app, delegateObj)
		},
	})
	app.SetDelegate(delegateObj)
	app.Run()
}

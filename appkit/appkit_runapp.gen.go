// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"runtime"
	"github.com/tmc/apple/foundation"
)

// RunApp initializes the shared NSApplication, creates a delegate via the
// delegate builder, and enters the main event loop. setupFn is called once
// the application finishes launching.
func RunApp(setupFn func(app NSApplication, delegate NSApplicationDelegateObject)) {
	runtime.LockOSThread()

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


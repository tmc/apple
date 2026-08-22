// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"runtime"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// AppKit is main-thread-only: NSUpdateCycleInitialize asserts with a bare
// brk #0x1 when it runs anywhere else, producing no ObjC exception and no
// crash report. Locking in RunApp is too late — it pins whichever thread the
// main goroutine has already migrated to, and a program that does blocking
// work before opening a window does migrate. Measured under load, an unpinned
// binary starts off the main thread 30% of the time.
//
// Pinning here means importing appkit pins the main OS thread, even for a
// program that never calls RunApp. That is a deliberate trade: every correct
// use of this package requires a pinned main thread anyway.
func init() {
	runtime.LockOSThread()
}

func isMainThread() bool {
	return objc.Send[bool](objc.ID(objc.GetClass("NSThread")), objc.Sel("isMainThread"))
}

// RunApp initializes the shared NSApplication, creates a delegate via the
// delegate builder, and enters the main event loop. setupFn is called once
// the application finishes launching.
//
// RunApp panics unless it is called from the main OS thread. The package init
// above covers the ordinary case; the check remains because init cannot cover
// "go appkit.RunApp(…)" or a dependency that called runtime.UnlockOSThread.
func RunApp(setupFn func(app NSApplication, delegate NSApplicationDelegateObject)) {
	if !isMainThread() {
		panic("appkit.RunApp must be called from the main OS thread; " +
			"add func init() { runtime.LockOSThread() } to package main, " +
			"and do not call RunApp from a goroutine")
	}
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

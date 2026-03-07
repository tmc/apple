package vzkit

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
)

// CFRunLoop functions loaded via purego.
var (
	cfRunLoopRunInMode    func(mode uintptr, seconds float64, returnAfterSourceHandled bool) int32
	cfRunLoopGetMain      func() uintptr
	cfRunLoopGetCurrent   func() uintptr
	cfRunLoopRun          func()
	cfRunLoopWakeUp       func(rl uintptr)
	kCFRunLoopDefaultMode uintptr
)

// CFRunLoop return values.
const (
	RunLoopRunFinished      = 1
	RunLoopRunStopped       = 2
	RunLoopRunTimedOut      = 3
	RunLoopRunHandledSource = 4
)

func initRunLoop() {
	cfLib, err := purego.Dlopen("/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation", purego.RTLD_LAZY)
	if err != nil {
		panic(err)
	}
	purego.RegisterLibFunc(&cfRunLoopRunInMode, cfLib, "CFRunLoopRunInMode")
	purego.RegisterLibFunc(&cfRunLoopGetMain, cfLib, "CFRunLoopGetMain")
	purego.RegisterLibFunc(&cfRunLoopGetCurrent, cfLib, "CFRunLoopGetCurrent")
	purego.RegisterLibFunc(&cfRunLoopRun, cfLib, "CFRunLoopRun")
	purego.RegisterLibFunc(&cfRunLoopWakeUp, cfLib, "CFRunLoopWakeUp")

	kCFRunLoopDefaultMode, err = purego.Dlsym(cfLib, "kCFRunLoopDefaultMode")
	if err != nil {
		panic(err)
	}
	// kCFRunLoopDefaultMode is a pointer to CFStringRef, dereference it.
	kCFRunLoopDefaultMode = *(*uintptr)(unsafe.Pointer(kCFRunLoopDefaultMode))
}

// RunRunLoopOnce runs the main CFRunLoop briefly to process pending callbacks.
func RunRunLoopOnce() {
	cfRunLoopRunInMode(kCFRunLoopDefaultMode, 0.01, false)
}

// RunRunLoopUntilDone runs the main run loop until done returns true.
// It pumps both CFRunLoop and NSRunLoop to ensure async callbacks are delivered.
// The optional progress callback is invoked each iteration.
func RunRunLoopUntilDone(done func() bool, progress func()) {
	for !done() {
		cfRunLoopRunInMode(kCFRunLoopDefaultMode, 0.1, false)

		mainRunLoop := foundation.GetRunLoopClass().MainRunLoop()
		if mainRunLoop.ID != 0 {
			futureDate := foundation.GetNSDateClass().DateWithTimeIntervalSinceNow(0.05)
			if futureDate.ID != 0 {
				mainRunLoop.RunModeBeforeDate(foundation.RunLoopDefaultMode, &futureDate)
			}
		}

		if progress != nil {
			progress()
		}
	}
}

// RunRunLoopAggressively pumps both CFRunLoop and NSRunLoop multiple times.
// Use this for long-running async operations that need thorough event processing.
func RunRunLoopAggressively() {
	mainRL := cfRunLoopGetMain()
	if mainRL != 0 {
		cfRunLoopWakeUp(mainRL)
	}

	for i := 0; i < 5; i++ {
		cfRunLoopRunInMode(kCFRunLoopDefaultMode, 0.1, false)
	}

	mainRunLoop := foundation.GetRunLoopClass().MainRunLoop()
	currentRunLoop := foundation.GetRunLoopClass().CurrentRunLoop()

	futureDate := foundation.GetNSDateClass().DateWithTimeIntervalSinceNow(0.1)
	if futureDate.ID != 0 {
		if mainRunLoop.ID != 0 {
			mainRunLoop.RunModeBeforeDate(foundation.RunLoopDefaultMode, &futureDate)
		}
		if currentRunLoop.ID != 0 && currentRunLoop.ID != mainRunLoop.ID {
			futureDate2 := foundation.GetNSDateClass().DateWithTimeIntervalSinceNow(0.1)
			currentRunLoop.RunModeBeforeDate(foundation.RunLoopDefaultMode, &futureDate2)
		}
	}
}

package vm

import (
	"fmt"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
)

var (
	cfRunLoopRunInMode    func(mode uintptr, seconds float64, returnAfterSourceHandled bool) int32
	cfRunLoopGetMain      func() uintptr
	cfRunLoopWakeUp       func(rl uintptr)
	kCFRunLoopDefaultMode uintptr
)

var (
	runLoopOnce sync.Once
	runLoopErr  error
)

func ensureRunLoop() error {
	runLoopOnce.Do(func() {
		runLoopErr = doInitRunLoop()
	})
	return runLoopErr
}

func doInitRunLoop() error {
	cfLib, err := purego.Dlopen("/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation", purego.RTLD_LAZY)
	if err != nil {
		return fmt.Errorf("vzkit/vm: load CoreFoundation: %w", err)
	}
	purego.RegisterLibFunc(&cfRunLoopRunInMode, cfLib, "CFRunLoopRunInMode")
	purego.RegisterLibFunc(&cfRunLoopGetMain, cfLib, "CFRunLoopGetMain")
	purego.RegisterLibFunc(&cfRunLoopWakeUp, cfLib, "CFRunLoopWakeUp")

	sym, err := purego.Dlsym(cfLib, "kCFRunLoopDefaultMode")
	if err != nil {
		return fmt.Errorf("vzkit/vm: resolve kCFRunLoopDefaultMode: %w", err)
	}
	kCFRunLoopDefaultMode = *(*uintptr)(unsafe.Pointer(sym))
	return nil
}

// RunLoopOnce runs the main CFRunLoop briefly to process pending callbacks.
func RunLoopOnce() {
	if err := ensureRunLoop(); err != nil {
		return
	}
	cfRunLoopRunInMode(kCFRunLoopDefaultMode, 0.01, false)
}

// RunLoopUntilDone runs the main run loop until done returns true.
func RunLoopUntilDone(done func() bool, progress func()) {
	if err := ensureRunLoop(); err != nil {
		return
	}
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

// RunLoopAggressively pumps both CFRunLoop and NSRunLoop multiple times.
func RunLoopAggressively() {
	if err := ensureRunLoop(); err != nil {
		return
	}
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

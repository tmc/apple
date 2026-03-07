// Code generated from Apple documentation by applegen. DO NOT EDIT.

//go:build darwin

package objc

import (
	"fmt"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
	purego_objc "github.com/ebitengine/purego/objc"
)

// ExceptionHandlerConfig configures how exceptions are handled.
type ExceptionHandlerConfig struct {
	// LogExceptions enables automatic logging of all exceptions
	LogExceptions bool

	// UseNSExceptionHandler enables the NSExceptionHandler delegate approach
	// This provides more control but requires the ExceptionHandling framework
	UseNSExceptionHandler bool

	// ExceptionHandlingMask sets which exception types to handle (if UseNSExceptionHandler)
	// Use constants: NSLogUncaughtExceptionMask, NSHandleUncaughtExceptionMask, etc.
	ExceptionHandlingMask uint

	// OnException is called when an exception is detected (before crash)
	OnException func(exc *ObjCException)
}

var (
	exHandlerInitOnce sync.Once

	// Low-level objc exception functions
	objcAddExceptionHandler    func(fn uintptr, ctx uintptr) uintptr
	objcRemoveExceptionHandler func(token uintptr)

	// NSExceptionHandler class (loaded on demand)
	nsExceptionHandlerClass Class
	nsExceptionHandlerOnce  sync.Once

	// Current exception handler delegate
	delegateClass       Class
	currentDelegateID   ID
	currentConfig       ExceptionHandlerConfig
	exHandlerMu         sync.Mutex
	exceptionHandlerSet bool

	// Exception handler masks (from ExceptionHandling framework)
	// These are defined here so we don't need the generated bindings
	NSLogUncaughtExceptionMask         uint = 1 << 0
	NSHandleUncaughtExceptionMask      uint = 1 << 1
	NSLogUncaughtSystemExceptionMask   uint = 1 << 2
	NSHandleUncaughtSystemExceptionMask uint = 1 << 3
	NSLogUncaughtRuntimeErrorMask      uint = 1 << 4
	NSHandleUncaughtRuntimeErrorMask   uint = 1 << 5
	NSLogTopLevelExceptionMask         uint = 1 << 6
	NSHandleTopLevelExceptionMask      uint = 1 << 7
	NSLogOtherExceptionMask            uint = 1 << 8
	NSHandleOtherExceptionMask         uint = 1 << 9
)

// initExceptionHandlerFuncs loads objc exception handler functions.
func initExceptionHandlerFuncs() {
	exHandlerInitOnce.Do(func() {
		objcLib, err := purego.Dlopen("/usr/lib/libobjc.A.dylib", purego.RTLD_LAZY)
		if err != nil {
			return
		}

		// These are macOS-only functions
		purego.RegisterLibFunc(&objcAddExceptionHandler, objcLib, "objc_addExceptionHandler")
		purego.RegisterLibFunc(&objcRemoveExceptionHandler, objcLib, "objc_removeExceptionHandler")
	})
}

// getNSExceptionHandlerClass returns the NSExceptionHandler class, loading it on demand.
func getNSExceptionHandlerClass() Class {
	nsExceptionHandlerOnce.Do(func() {
		// Load the ExceptionHandling framework
		_, err := purego.Dlopen("/System/Library/Frameworks/ExceptionHandling.framework/ExceptionHandling", purego.RTLD_LAZY)
		if err != nil {
			return
		}
		nsExceptionHandlerClass = GetClass("NSExceptionHandler")
	})
	return nsExceptionHandlerClass
}

// SetupExceptionHandler configures exception handling using the most effective
// available mechanism. This uses NSExceptionHandler with a delegate to intercept
// exceptions via purego.
//
// Example:
//
//	objc.SetupExceptionHandler(objc.ExceptionHandlerConfig{
//	    LogExceptions: true,
//	    OnException: func(exc *objc.ObjCException) {
//	        fmt.Printf("Exception: %s\n", exc.Name)
//	    },
//	})
func SetupExceptionHandler(config ExceptionHandlerConfig) error {
	exHandlerMu.Lock()
	defer exHandlerMu.Unlock()

	currentConfig = config

	// Always set up the preprocessor for exception detection
	initExceptionPreprocessor()

	if config.LogExceptions {
		SetExceptionPreprocessor(func(exc ID) ID {
			info := GetExceptionInfo(exc)
			if info != nil {
				fmt.Printf("\n=== Objective-C Exception ===\n")
				fmt.Printf("Name: %s\n", info.Name)
				fmt.Printf("Reason: %s\n", info.Reason)
				if len(info.CallStack) > 0 {
					fmt.Printf("Call Stack:\n")
					for i, frame := range info.CallStack {
						fmt.Printf("  %d: %s\n", i, frame)
					}
				}
				fmt.Printf("=============================\n\n")

				if config.OnException != nil {
					config.OnException(info)
				}
			}
			return exc
		})
	} else if config.OnException != nil {
		SetExceptionPreprocessor(func(exc ID) ID {
			info := GetExceptionInfo(exc)
			if info != nil {
				config.OnException(info)
			}
			return exc
		})
	}

	// Set up NSExceptionHandler with delegate if requested
	if config.UseNSExceptionHandler {
		return setupNSExceptionHandlerDelegate(config)
	}

	return nil
}

// setupNSExceptionHandlerDelegate creates and registers an NSExceptionHandler delegate.
func setupNSExceptionHandlerDelegate(config ExceptionHandlerConfig) error {
	handlerClass := getNSExceptionHandlerClass()
	if handlerClass == 0 {
		return fmt.Errorf("NSExceptionHandler class not available (ExceptionHandling framework not loaded)")
	}

	// Create delegate class if needed
	if delegateClass == 0 {
		var err error
		delegateClass, err = createExceptionHandlerDelegate()
		if err != nil {
			return fmt.Errorf("failed to create delegate class: %w", err)
		}
	}

	// Create delegate instance
	if currentDelegateID != 0 {
		// Release old delegate
		Send[ID](currentDelegateID, Sel("release"))
	}
	currentDelegateID = Send[ID](ID(delegateClass), Sel("new"))

	// Get default exception handler
	handler := Send[ID](ID(handlerClass), Sel("defaultExceptionHandler"))

	// Set our delegate
	Send[ID](handler, Sel("setDelegate:"), currentDelegateID)

	// Set handling mask
	mask := config.ExceptionHandlingMask
	if mask == 0 {
		// Default to logging and handling uncaught exceptions
		mask = NSLogUncaughtExceptionMask | NSHandleUncaughtExceptionMask |
			NSLogUncaughtSystemExceptionMask | NSLogUncaughtRuntimeErrorMask
	}
	Send[ID](handler, Sel("setExceptionHandlingMask:"), mask)

	exceptionHandlerSet = true
	return nil
}

// createExceptionHandlerDelegate creates an Objective-C class that implements
// the NSExceptionHandler delegate protocol.
func createExceptionHandlerDelegate() (Class, error) {
	nsObjectClass := GetClass("NSObject")

	// Define the delegate methods
	methods := []purego_objc.MethodDef{
		{
			// - (BOOL)exceptionHandler:(NSExceptionHandler *)sender shouldLogException:(NSException *)exception mask:(NSUInteger)aMask
			Cmd: Sel("exceptionHandler:shouldLogException:mask:"),
			Fn: func(self ID, cmd SEL, sender ID, exception ID, mask uint) bool {
				// Get exception info and call user callback
				exHandlerMu.Lock()
				cfg := currentConfig
				exHandlerMu.Unlock()

				if cfg.OnException != nil {
					info := GetExceptionInfo(exception)
					if info != nil {
						cfg.OnException(info)
					}
				}

				// Return true to allow logging
				return true
			},
		},
		{
			// - (BOOL)exceptionHandler:(NSExceptionHandler *)sender shouldHandleException:(NSException *)exception mask:(NSUInteger)aMask
			Cmd: Sel("exceptionHandler:shouldHandleException:mask:"),
			Fn: func(self ID, cmd SEL, sender ID, exception ID, mask uint) bool {
				// Get exception info and call user callback
				exHandlerMu.Lock()
				cfg := currentConfig
				exHandlerMu.Unlock()

				if cfg.OnException != nil {
					info := GetExceptionInfo(exception)
					if info != nil {
						cfg.OnException(info)
					}
				}

				// Return true to allow handling (which terminates the thread)
				// Return false to try to continue (but this usually doesn't work)
				return true
			},
		},
	}

	return purego_objc.RegisterClass(
		"GoExceptionHandlerDelegate",
		nsObjectClass,
		nil,    // protocols
		nil,    // ivars
		methods,
	)
}

// AddExceptionHandler adds a low-level exception handler using objc_addExceptionHandler.
// This is called during exception unwinding and cannot prevent the crash, but can
// capture exception information.
//
// Returns a token that can be passed to RemoveExceptionHandler.
// Note: This is macOS-only (not available on iOS).
func AddExceptionHandler(fn func(exception ID, context unsafe.Pointer)) uintptr {
	initExceptionHandlerFuncs()
	if objcAddExceptionHandler == nil {
		return 0
	}

	callback := purego.NewCallback(func(exception uintptr, ctx uintptr) {
		fn(ID(exception), unsafe.Pointer(ctx))
	})

	return objcAddExceptionHandler(callback, 0)
}

// RemoveExceptionHandler removes an exception handler added by AddExceptionHandler.
func RemoveExceptionHandler(token uintptr) {
	initExceptionHandlerFuncs()
	if objcRemoveExceptionHandler == nil || token == 0 {
		return
	}
	objcRemoveExceptionHandler(token)
}

// CleanupExceptionHandler removes any configured exception handlers.
func CleanupExceptionHandler() {
	exHandlerMu.Lock()
	defer exHandlerMu.Unlock()

	if exceptionHandlerSet && nsExceptionHandlerClass != 0 {
		handler := Send[ID](ID(nsExceptionHandlerClass), Sel("defaultExceptionHandler"))
		Send[ID](handler, Sel("setDelegate:"), ID(0))
	}

	if currentDelegateID != 0 {
		Send[ID](currentDelegateID, Sel("release"))
		currentDelegateID = 0
	}

	exceptionHandlerSet = false
}

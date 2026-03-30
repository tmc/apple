// Code generated from Apple documentation by applegen. DO NOT EDIT.

//go:build darwin

// Package objc provides cached Objective-C runtime helpers.
package objc

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/ebitengine/purego"
)

// ANSI color codes for optional colorized output.
const (
	colorReset  = "\033[0m"
	colorBold   = "\033[1m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
)

// colorEnabled reports whether stderr is a terminal that supports colors.
func colorEnabled() bool {
	fi, err := os.Stderr.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

// ObjCException represents an Objective-C exception.
type ObjCException struct {
	Name      string   // Exception name (e.g., "NSInvalidArgumentException")
	Reason    string   // Exception reason/message
	CallStack []string // Objective-C call stack symbols
	exception ID       // The underlying NSException object
}

// Error implements the error interface.
func (e *ObjCException) Error() string {
	return fmt.Sprintf("%s: %s", e.Name, e.Reason)
}

// Exception returns the underlying NSException ID.
func (e *ObjCException) Exception() ID {
	return e.exception
}

// ExceptionHandler is called when an uncaught Objective-C exception occurs.
// The handler receives the exception name, reason, and call stack.
// Note: The program will terminate after this handler returns - recovery is not possible.
type ExceptionHandler func(name, reason string, callStack []string)

// ExceptionPreprocessorFunc is called before any Objective-C exception is thrown.
// This allows logging/debugging but cannot prevent the exception in pure Go.
// For actual exception catching, use the cgoexception subpackage.
type ExceptionPreprocessorFunc func(exception ID) ID

var (
	// Preprocessor state
	catchInitOnce              sync.Once
	objcSetExceptionPreproc    func(uintptr) uintptr
	preprocessorCallbackHandle uintptr
	catchMu                    sync.Mutex
	lastSeenException          ID
	exceptionPreprocess        ExceptionPreprocessorFunc

	// NSSetUncaughtExceptionHandler state
	exceptionHandlerOnce    sync.Once
	customExceptionHandler  ExceptionHandler
	exceptionHandlerMu      sync.Mutex
	foundationLib           uintptr
	nsSetExceptionHandler   func(uintptr)
	exceptionCallbackHandle uintptr
)

func init() {
	if os.Getenv("OBJC_EXCEPTION_HANDLER") == "off" {
		return
	}
	EnableDefaultExceptionHandler()
}

// initExceptionPreprocessor sets up the exception preprocessor via
// objc_setExceptionPreprocessor. This fires BEFORE objc_exception_throw,
// making it the only interception point that fires before C++ abort().
func initExceptionPreprocessor() {
	catchInitOnce.Do(func() {
		objcLib, err := purego.Dlopen("/usr/lib/libobjc.A.dylib", purego.RTLD_LAZY)
		if err != nil {
			return
		}

		purego.RegisterLibFunc(&objcSetExceptionPreproc, objcLib, "objc_setExceptionPreprocessor")

		preprocessorCallbackHandle = purego.NewCallback(func(exception uintptr) uintptr {
			exc := ID(exception)

			catchMu.Lock()
			lastSeenException = exc
			fn := exceptionPreprocess
			catchMu.Unlock()

			if fn != nil {
				return uintptr(fn(exc))
			}
			// No custom preprocessor — handle and exit before C++ abort().
			handleException(exc)
			os.Exit(1)
			return exception
		})

		objcSetExceptionPreproc(preprocessorCallbackHandle)
	})
}

// SetExceptionPreprocessor sets a custom function that is called before
// any Objective-C exception is thrown. This can be used for logging and debugging.
//
// Passing nil restores the default behavior (handleException + os.Exit(1)).
//
// NOTE: The preprocessor cannot prevent the exception in pure Go - the program
// will still crash. For actual exception catching, use the cgoexception subpackage.
//
// Example:
//
//	objc.SetExceptionPreprocessor(func(exc objc.ID) objc.ID {
//	    info := objc.GetExceptionInfo(exc)
//	    fmt.Printf("Exception about to throw: %s\n", info.Name)
//	    return exc
//	})
func SetExceptionPreprocessor(fn ExceptionPreprocessorFunc) {
	catchMu.Lock()
	exceptionPreprocess = fn
	catchMu.Unlock()
	initExceptionPreprocessor()
}

// GetExceptionInfo extracts information from an NSException object.
func GetExceptionInfo(exception ID) *ObjCException {
	if exception == 0 {
		return nil
	}

	exc := &ObjCException{exception: exception}

	nameID := Send[ID](exception, Sel("name"))
	if nameID != 0 {
		cstr := Send[*byte](nameID, Sel("UTF8String"))
		exc.Name = GoString(cstr)
	}

	reasonID := Send[ID](exception, Sel("reason"))
	if reasonID != 0 {
		cstr := Send[*byte](reasonID, Sel("UTF8String"))
		exc.Reason = GoString(cstr)
	}

	stackID := Send[ID](exception, Sel("callStackSymbols"))
	if stackID != 0 {
		count := Send[uint](stackID, Sel("count"))
		for i := uint(0); i < count; i++ {
			symbolID := Send[ID](stackID, Sel("objectAtIndex:"), i)
			if symbolID != 0 {
				cstr := Send[*byte](symbolID, Sel("UTF8String"))
				exc.CallStack = append(exc.CallStack, GoString(cstr))
			}
		}
	}

	return exc
}

// GetLastException returns the last exception seen by the preprocessor.
func GetLastException() ID {
	catchMu.Lock()
	defer catchMu.Unlock()
	return lastSeenException
}

// EnableExceptionLogging enables automatic logging of all Objective-C exceptions
// before they are thrown. Useful for debugging.
func EnableExceptionLogging() {
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
		}
		return exc
	})
}

// SetUncaughtExceptionHandler registers a custom handler for uncaught Objective-C exceptions.
//
// IMPORTANT: Objective-C exceptions cannot be recovered from in Go. This handler is called
// just before the program terminates, giving you a chance to:
//   - Log diagnostic information
//   - Write crash reports
//   - Clean up resources
//
// The handler receives:
//   - name: The exception name (e.g., "NSInvalidArgumentException")
//   - reason: The exception reason (e.g., "-[NSObject fakeMethod]: unrecognized selector")
//   - callStack: The Objective-C call stack symbols
//
// Example:
//
//	objc.SetUncaughtExceptionHandler(func(name, reason string, callStack []string) {
//	    fmt.Fprintf(os.Stderr, "Objective-C Exception: %s\n", name)
//	    fmt.Fprintf(os.Stderr, "Reason: %s\n", reason)
//	    for _, frame := range callStack {
//	        fmt.Fprintf(os.Stderr, "  %s\n", frame)
//	    }
//	})
func SetUncaughtExceptionHandler(handler ExceptionHandler) {
	exceptionHandlerMu.Lock()
	customExceptionHandler = handler
	exceptionHandlerMu.Unlock()

	exceptionHandlerOnce.Do(func() {
		initExceptionHandler()
	})
}

// initExceptionHandler sets up the Objective-C exception handler using purego.
func initExceptionHandler() {
	var err error
	foundationLib, err = purego.Dlopen("/System/Library/Frameworks/Foundation.framework/Foundation", purego.RTLD_LAZY)
	if err != nil {
		fmt.Fprintf(os.Stderr, "objc: failed to load Foundation framework: %v\n", err)
		return
	}

	// Register NSSetUncaughtExceptionHandler
	purego.RegisterLibFunc(&nsSetExceptionHandler, foundationLib, "NSSetUncaughtExceptionHandler")

	// Create callback for the exception handler
	// The callback receives an NSException* (objc.ID)
	exceptionCallbackHandle = purego.NewCallback(func(exception uintptr) uintptr {
		handleException(ID(exception))
		return 0
	})

	// Set the exception handler
	nsSetExceptionHandler(exceptionCallbackHandle)
}

// handleException is called by the Objective-C runtime when an uncaught exception occurs.
func handleException(exception ID) {
	if exception == 0 {
		return
	}

	// Extract exception information
	name := ""
	reason := ""
	var callStack []string

	// Get exception name
	nameID := Send[ID](exception, Sel("name"))
	if nameID != 0 {
		cstr := Send[*byte](nameID, Sel("UTF8String"))
		name = GoString(cstr)
	}

	// Get exception reason
	reasonID := Send[ID](exception, Sel("reason"))
	if reasonID != 0 {
		cstr := Send[*byte](reasonID, Sel("UTF8String"))
		reason = GoString(cstr)
	}

	// Get call stack symbols
	stackID := Send[ID](exception, Sel("callStackSymbols"))
	if stackID != 0 {
		count := Send[uint](stackID, Sel("count"))
		for i := uint(0); i < count; i++ {
			symbolID := Send[ID](stackID, Sel("objectAtIndex:"), i)
			if symbolID != 0 {
				cstr := Send[*byte](symbolID, Sel("UTF8String"))
				callStack = append(callStack, GoString(cstr))
			}
		}
	}

	// Call custom handler if set
	exceptionHandlerMu.Lock()
	handler := customExceptionHandler
	exceptionHandlerMu.Unlock()

	if handler != nil {
		handler(name, reason, callStack)
	} else {
		// Default behavior: print to stderr
		defaultExceptionHandler(name, reason, callStack)
	}
}

// defaultExceptionHandler provides a default handler that prints diagnostic information
// and exits cleanly. Without os.Exit, the ObjC runtime calls abort() which triggers
// the Go runtime's crash handler (all goroutines + register dump).
func defaultExceptionHandler(name, reason string, callStack []string) {
	color := colorEnabled()

	// Print exception info — the preprocessor exits before the ObjC runtime
	// gets a chance to print this, so we must do it ourselves.
	if color {
		fmt.Fprintf(os.Stderr, "\n%s%sObjC exception:%s %s%s%s\n", colorBold, colorRed, colorReset, colorBold, name, colorReset)
		fmt.Fprintf(os.Stderr, "%s%s%s\n", colorYellow, reason, colorReset)
	} else {
		fmt.Fprintf(os.Stderr, "\nObjC exception: %s\n", name)
		fmt.Fprintf(os.Stderr, "%s\n", reason)
	}

	if len(callStack) > 0 {
		fmt.Fprintf(os.Stderr, "\nObjC call stack:\n")
		for _, frame := range callStack {
			fmt.Fprintf(os.Stderr, "  %s\n", frame)
		}
	}

	fmt.Fprintf(os.Stderr, "\ngoroutine stack:\n")
	buf := make([]byte, 8192)
	n := runtime.Stack(buf, false)
	printFilteredStack(os.Stderr, string(buf[:n]), color)
	if color {
		fmt.Fprintf(os.Stderr, "\n%s(set OBJC_EXCEPTION_HANDLER=off for full crash)%s\n", colorCyan, colorReset)
	} else {
		fmt.Fprintf(os.Stderr, "\n(set OBJC_EXCEPTION_HANDLER=off for full crash)\n")
	}
	os.Exit(1)
}

// printFilteredStack parses runtime.Stack output and prints only relevant frames,
// skipping internal handler frames (this package's handleException, reflect, purego).
func printFilteredStack(w *os.File, stack string, color bool) {
	lines := strings.Split(stack, "\n")

	// Stack format: pairs of lines (function, then file:line),
	// preceded by a "goroutine N [status]:" header.
	type frame struct {
		funcLine string // e.g. "github.com/.../objc.Send[...](0x...)"
		fileLine string // e.g. "\t/path/to/objc.go:133 +0x18c"
	}

	var frames []frame
	for i := 1; i < len(lines)-1; i += 2 {
		if i+1 < len(lines) {
			frames = append(frames, frame{lines[i], lines[i+1]})
		}
	}

	// Find the first objc.Send frame — that's where user code begins.
	// Show from there downward, skipping reflect/purego internals.
	startIdx := -1
	for i, f := range frames {
		if strings.Contains(f.funcLine, "/objc.Send[") || strings.Contains(f.funcLine, "/objc.Send(") {
			startIdx = i
			break
		}
	}
	if startIdx < 0 {
		// No Send frame found, print everything
		fmt.Fprintf(w, "%s\n", stack)
		return
	}

	for _, f := range frames[startIdx:] {
		// Skip reflect and purego internals
		if strings.Contains(f.funcLine, "reflect.Value.") ||
			strings.Contains(f.funcLine, "ebitengine/purego.") {
			continue
		}
		// Extract short package name and function from full path
		pkg, fn := shortFrame(f.funcLine)
		// Extract file:line from the file line
		fileLine := strings.TrimSpace(f.fileLine)
		if idx := strings.LastIndex(fileLine, "/"); idx >= 0 {
			fileLine = fileLine[idx+1:]
		}
		// Remove +0x... offset
		if idx := strings.Index(fileLine, " +0x"); idx >= 0 {
			fileLine = fileLine[:idx]
		}
		if color {
			c := colorCyan
			if pkg == "main" {
				c = colorYellow
			}
			fmt.Fprintf(w, "  %s%-8s%s %-28s %s%s%s\n", colorBold, pkg, colorReset, fileLine, c, fn, colorReset)
		} else {
			fmt.Fprintf(w, "  %-8s %-28s %s\n", pkg, fileLine, fn)
		}
	}
}

// shortFrame extracts the short package name and function name from a full Go stack frame.
// "github.com/tmc/appledocs/generated/objc.Send[...](...)" -> "objc", "Send[...](...)"
func shortFrame(funcLine string) (pkg, fn string) {
	s := strings.TrimSpace(funcLine)
	// Split off args at first '(' that isn't inside '[]' (generics).
	// For "objc.Send[go.shape.uintptr](0x...)", we want to split at the '(' after ']'.
	bracketDepth := 0
	parenIdx := -1
	for i, c := range s {
		switch c {
		case '[':
			bracketDepth++
		case ']':
			bracketDepth--
		case '(':
			if bracketDepth == 0 {
				parenIdx = i
			}
		}
		if parenIdx >= 0 {
			break
		}
	}
	qualName := s
	args := ""
	if parenIdx >= 0 {
		qualName = s[:parenIdx]
		args = s[parenIdx:]
	}
	// qualName is like "github.com/tmc/appledocs/generated/objc.Send[...]"
	// Strip generic type params for cleaner display
	if bracketIdx := strings.Index(qualName, "["); bracketIdx >= 0 {
		qualName = qualName[:bracketIdx] + "[...]"
	}
	// Find package.Function boundary — the first '.' after the last '/'
	if slashIdx := strings.LastIndex(qualName, "/"); slashIdx >= 0 {
		rest := qualName[slashIdx+1:] // "objc.Send[...]"
		if dotIdx := strings.Index(rest, "."); dotIdx >= 0 {
			pkg = rest[:dotIdx]
			fn = rest[dotIdx+1:]
		} else {
			pkg = rest
			fn = ""
		}
	} else if dotIdx := strings.Index(qualName, "."); dotIdx >= 0 {
		pkg = qualName[:dotIdx]
		fn = qualName[dotIdx+1:]
	} else {
		pkg = qualName
	}
	// Truncate long args
	if len(args) > 40 {
		args = args[:37] + "...)"
	}
	return pkg, fn + args
}

// EnableDefaultExceptionHandler enables the default exception handler that prints
// diagnostic information to stderr and exits cleanly. This is called automatically
// by init(). The preprocessor is also installed so that exceptions are intercepted
// before the C++ runtime calls abort().
//
// To replace the default behavior, call [SetExceptionPreprocessor] or
// [SetUncaughtExceptionHandler] with a non-nil handler. To disable exception
// handling entirely, call [DisableDefaultExceptionHandler].
func EnableDefaultExceptionHandler() {
	SetUncaughtExceptionHandler(nil)
	initExceptionPreprocessor()
}

// DisableDefaultExceptionHandler removes both the exception preprocessor and the
// uncaught exception handler, reverting to the default ObjC runtime behavior
// (abort on uncaught exceptions). This is useful if you want to handle exceptions
// entirely through your own mechanism or prefer raw SIGABRT crashes for a debugger.
func DisableDefaultExceptionHandler() {
	// Clear custom handlers so the preprocessor falls through to the ObjC runtime.
	catchMu.Lock()
	exceptionPreprocess = func(exc ID) ID { return exc }
	catchMu.Unlock()

	exceptionHandlerMu.Lock()
	customExceptionHandler = nil
	exceptionHandlerMu.Unlock()
}

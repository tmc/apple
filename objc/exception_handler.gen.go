// Code generated from internal/generator/templates/runtime/objc.txtar by applegen. DO NOT EDIT.

//go:build darwin

// Package objc provides cached Objective-C runtime helpers.
package objc

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ebitengine/purego"
)

// reportWatchdog is how long the uncaught-exception report may take before the
// process is killed anyway. It exists so a report that deadlocks degrades to a
// nonzero exit rather than a hang.
const reportWatchdog = 2 * time.Second

// armReportWatchdog kills the process if the report has not finished in
// [reportWatchdog]. The report itself calls os.Exit on success, so this fires
// only when reporting is stuck.
func armReportWatchdog() {
	go func() {
		time.Sleep(reportWatchdog)
		os.Stderr.WriteString("\nobjc: exception report timed out; exiting\n")
		os.Exit(1)
	}()
}

// ANSI color codes for optional colorized output.
const (
	colorReset  = "\033[0m"
	colorBold   = "\033[1m"
	colorDim    = "\033[2m"
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

// ExceptionHandler is called when an Objective-C exception reaches the top of
// the stack with nothing to catch it. The handler receives the exception name,
// reason, and call stack.
//
// The process terminates once the handler returns: an Objective-C exception
// cannot be recovered from in Go, and recover cannot see one. See the package
// documentation for why.
type ExceptionHandler func(name, reason string, callStack []string)

// ExceptionPreprocessorFunc is called before any Objective-C exception is
// thrown, whether or not something later catches it.
//
// It can observe but not prevent: pure Go cannot catch an Objective-C
// exception. Any method documented to raise aborts the process, so validate
// arguments before the call rather than relying on recovery.
//
// Installing a preprocessor has a cost. CoreFoundation occupies this slot with
// __exceptionPreprocess, which is what records an exception's call stack;
// displacing it loses the frames that identify the failing call. This package
// therefore installs no preprocessor unless you ask for one.
type ExceptionPreprocessorFunc func(exception ID) ID

var (
	// Preprocessor state. Nothing here is initialized unless the caller
	// explicitly installs a preprocessor.
	catchInitOnce              sync.Once
	objcSetExceptionPreproc    func(uintptr) uintptr
	preprocessorCallbackHandle uintptr
	prevExceptionPreprocHandle uintptr
	prevExceptionPreproc       func(uintptr) uintptr
	catchMu                    sync.Mutex
	lastSeenException          ID
	exceptionPreprocess        ExceptionPreprocessorFunc

	// Uncaught-handler state.
	uncaughtOnce           sync.Once
	uncaughtHandlerEntered atomic.Bool
	objcSetUncaughtHandler func(uintptr) uintptr
	uncaughtCallbackHandle uintptr
	prevUncaughtHandle     uintptr
	exceptionHandlerMu     sync.Mutex
	customExceptionHandler ExceptionHandler
)

func init() {
	if os.Getenv("OBJC_EXCEPTION_HANDLER") == "off" {
		return
	}
	EnableDefaultExceptionHandler()
}

// initUncaughtExceptionHandler installs this package's reporter via libobjc's
// objc_setUncaughtExceptionHandler.
//
// That hook, rather than Foundation's NSSetUncaughtExceptionHandler, because it
// runs earlier: it fires before CoreFoundation logs its own "*** Terminating
// app due to uncaught exception" block. Exiting from here therefore replaces
// that output instead of appending to it. Foundation's hook runs after the log
// and can only add a second copy.
//
// The hook is only reached when nothing caught the exception, so installing it
// costs nothing on paths where an Objective-C @catch would have handled the
// throw.
func initUncaughtExceptionHandler() {
	uncaughtOnce.Do(func() {
		objcLib, err := purego.Dlopen("/usr/lib/libobjc.A.dylib", purego.RTLD_LAZY)
		if err != nil {
			return
		}
		sym, err := purego.Dlsym(objcLib, "objc_setUncaughtExceptionHandler")
		if err != nil {
			return
		}
		purego.RegisterFunc(&objcSetUncaughtHandler, sym)

		// objc-exception.h:
		//
		//	typedef void (*objc_uncaught_exception_handler)(id exception);
		//
		// One argument, no return value. The two-argument shape that also takes
		// a context pointer is objc_exception_handler, used by
		// objc_addExceptionHandler — a different hook with a different slot.
		uncaughtCallbackHandle = purego.NewCallback(func(exception uintptr) {
			// Reporting is not reentrant, and it is not obviously safe to run
			// twice. Building the report sends Objective-C messages and
			// allocates -- callStackSymbols allocates objects -- so the report
			// can itself throw, and that second exception is also uncaught, so
			// libobjc calls this handler again on this same thread from inside
			// itself. Without this guard that is unbounded recursion or two
			// interleaved reports.
			if !uncaughtHandlerEntered.CompareAndSwap(false, true) {
				os.Stderr.WriteString("\nobjc: exception raised while reporting an exception; exiting\n")
				os.Exit(1)
			}

			// Formatting runs with the process in an arbitrary state. If the
			// throw came from inside an allocator path, symbolication can
			// deadlock against a lock the throwing thread already holds, and a
			// hang is worse than a crash: no output, no exit status, and a
			// caller that waits forever. Bound it.
			armReportWatchdog()

			exc := ID(exception)

			exceptionHandlerMu.Lock()
			handler := customExceptionHandler
			exceptionHandlerMu.Unlock()

			catchMu.Lock()
			setLastSeenExceptionLocked(exc)
			catchMu.Unlock()

			if handler != nil {
				info := GetExceptionInfo(exc)
				if info != nil {
					handler(info.Name, info.Reason, info.CallStack)
				}
			} else {
				reportException(exc)
			}
			os.Exit(1)
		})

		// The setter returns the handler it displaced, which is what makes
		// DisableDefaultExceptionHandler able to put the runtime back exactly
		// as it found it.
		prevUncaughtHandle = objcSetUncaughtHandler(uncaughtCallbackHandle)
	})
}

// initExceptionPreprocessor installs a preprocessor via
// objc_setExceptionPreprocessor, chaining to whatever occupied the slot before.
// It is called only from [SetExceptionPreprocessor]; nothing installs a
// preprocessor by default.
func initExceptionPreprocessor() {
	catchInitOnce.Do(func() {
		objcLib, err := purego.Dlopen("/usr/lib/libobjc.A.dylib", purego.RTLD_LAZY)
		if err != nil {
			return
		}

		purego.RegisterLibFunc(&objcSetExceptionPreproc, objcLib, "objc_setExceptionPreprocessor")

		preprocessorCallbackHandle = purego.NewCallback(func(exception uintptr) uintptr {
			// Chain first. CoreFoundation registers __exceptionPreprocess here,
			// and that is what populates callStackReturnAddresses; skipping it
			// leaves every NSException without a call stack.
			exc := callPreviousPreprocessor(ID(exception))

			catchMu.Lock()
			setLastSeenExceptionLocked(exc)
			fn := exceptionPreprocess
			catchMu.Unlock()

			if fn != nil {
				return uintptr(fn(exc))
			}
			return uintptr(exc)
		})

		prevExceptionPreprocHandle = objcSetExceptionPreproc(preprocessorCallbackHandle)
		if prevExceptionPreprocHandle != 0 {
			purego.RegisterFunc(&prevExceptionPreproc, prevExceptionPreprocHandle)
		}
	})
}

// callPreviousPreprocessor forwards an exception to the preprocessor that was
// installed before this package's, returning the (possibly replaced) exception.
// It returns exc unchanged when there was no previous preprocessor.
func callPreviousPreprocessor(exc ID) ID {
	catchMu.Lock()
	fn := prevExceptionPreproc
	catchMu.Unlock()
	if fn == nil {
		return exc
	}
	if replaced := ID(fn(uintptr(exc))); replaced != 0 {
		return replaced
	}
	return exc
}

// setLastSeenExceptionLocked stores exc as the last seen exception, retaining it
// and releasing the value it replaces. Callers must hold catchMu.
//
// The retain is required: the exception is released once the runtime finishes
// unwinding, so an unretained copy of the pointer dangles as soon as an
// autorelease pool drains, and any later message send targets whatever object
// has since reused the address.
func setLastSeenExceptionLocked(exc ID) {
	if exc == lastSeenException {
		return
	}
	if exc != 0 {
		Send[ID](exc, Sel("retain"))
	}
	if lastSeenException != 0 {
		Send[ID](lastSeenException, Sel("release"))
	}
	lastSeenException = exc
}

// SetExceptionPreprocessor sets a function called before any Objective-C
// exception is thrown, for logging and debugging. Passing nil clears it.
//
// The preprocessor cannot prevent the exception; see [ExceptionPreprocessorFunc]
// for the cost of occupying this slot.
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

// GetLastException returns the last exception this package observed, either in
// an installed preprocessor or in the uncaught-exception handler.
func GetLastException() ID {
	catchMu.Lock()
	defer catchMu.Unlock()
	return lastSeenException
}

// EnableExceptionLogging logs every Objective-C exception as it is thrown,
// including ones that are subsequently caught. Useful for debugging; it
// installs a preprocessor, so see [ExceptionPreprocessorFunc] for the cost.
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

// SetUncaughtExceptionHandler replaces the report printed when an Objective-C
// exception reaches the top of the stack uncaught. Passing nil restores the
// default report.
//
// The process exits once the handler returns. The exception cannot be
// recovered from: recover does not see it, and there is no point at which
// execution can resume. Use the handler to log or to write a crash report.
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
	initUncaughtExceptionHandler()

	// Reinstall if DisableDefaultExceptionHandler put the previous handler
	// back; initUncaughtExceptionHandler only runs its body once.
	exceptionHandlerMu.Lock()
	if objcSetUncaughtHandler != nil && uncaughtCallbackHandle != 0 {
		if displaced := objcSetUncaughtHandler(uncaughtCallbackHandle); displaced != uncaughtCallbackHandle {
			prevUncaughtHandle = displaced
		}
	}
	exceptionHandlerMu.Unlock()
}

// EnableDefaultExceptionHandler installs the default report for uncaught
// Objective-C exceptions. init calls it unless OBJC_EXCEPTION_HANDLER=off.
//
// It installs no exception preprocessor: that slot belongs to CoreFoundation,
// and taking it costs the call-stack frames that identify the failing method.
//
// To replace the report, call [SetUncaughtExceptionHandler]. To restore the
// runtime's own behavior, call [DisableDefaultExceptionHandler].
func EnableDefaultExceptionHandler() {
	SetUncaughtExceptionHandler(nil)
}

// DisableDefaultExceptionHandler stops this package from reporting uncaught
// exceptions and uninstalls any preprocessor it installed, restoring the one it
// displaced. Afterwards an uncaught exception produces the runtime's own output.
func DisableDefaultExceptionHandler() {
	catchMu.Lock()
	if objcSetExceptionPreproc != nil {
		// Put back what we displaced rather than leaving a Go pass-through in
		// the slot; a pass-through keeps the callback wired and keeps
		// CoreFoundation's preprocessor out of the chain.
		objcSetExceptionPreproc(prevExceptionPreprocHandle)
	}
	exceptionPreprocess = nil
	setLastSeenExceptionLocked(0)
	catchMu.Unlock()

	exceptionHandlerMu.Lock()
	customExceptionHandler = nil
	if objcSetUncaughtHandler != nil && prevUncaughtHandle != 0 {
		objcSetUncaughtHandler(prevUncaughtHandle)
	}
	exceptionHandlerMu.Unlock()
}

// frameNotes annotates native stack frames whose presence explains something
// about the failure. It adds a note to a frame; it does not decide whether the
// frame is shown.
//
// Every native frame is printed, with its image and its offset into that image.
// Symbol names cannot be trusted on their own: symbolication resolves a return
// address to the nearest preceding exported symbol, so frames from static or
// inlined code are attributed to whatever happened to be linked before them.
// "-[NSTaggedPointerString hash]" and "CFArrayApply" have both appeared this
// way in stacks that touched neither.
//
// An earlier version of this code hid every frame it could not annotate, on the
// grounds that an unexplainable frame is not evidence. That reasoning does not
// hold: misattribution lands on annotated names exactly as easily, so hiding
// the rest filtered which misattributions were shown while lending them more
// authority. Offset does not rescue the name either, and rescues it backwards
// if you assume small offsets are trustworthy -- in a real unrecognized-selector
// stack the genuine frame is "___forwarding___ + 1480" and the misattributed one
// is "-[NSObject(NSObject) __retain_OA] + 0". Offset 0 is the signature of
// misattribution: it means the return address landed exactly on a symbol
// boundary, which is when the nearest preceding symbol is most likely the wrong
// one.
//
// So: print the raw material. "image + offset" is correct whether or not the
// name is, and it can be resolved later with atos from a crash report. A
// confidently misattributed name cannot be un-believed.
var frameNotes = map[string]string{
	"___forwarding___":             "class has no such method",
	"objc_exception_rethrow":       "rethrown from a @catch",
	"-[NSException raise]":         "raised explicitly",
	"+[NSException raise:format:]": "raised explicitly",
	"CFRunLoopRunSpecific":         "thrown from inside a run loop",
	"-[NSApplication run]":         "thrown from the AppKit event loop",
}

// reportException prints the default uncaught-exception report: the exception,
// then one stack running innermost-first across the Go/Objective-C boundary.
func reportException(exc ID) {
	full := os.Getenv("OBJC_EXCEPTION_STACK") == "full"
	color := colorEnabled()
	c := func(s, code string) string {
		if !color {
			return s
		}
		return code + s + colorReset
	}

	info := GetExceptionInfo(exc)
	if info == nil {
		fmt.Fprintf(os.Stderr, "\n%s\n", c("ObjC exception: <nil>", colorBold+colorRed))
		return
	}

	fmt.Fprintf(os.Stderr, "\n%s %s\n", c("ObjC exception:", colorBold+colorRed), c(info.Name, colorBold))
	fmt.Fprintf(os.Stderr, "%s\n\n", c(info.Reason, colorYellow))

	// Native frames, innermost first, stopping at the Go boundary. All of them:
	// see [frameNotes] for why none are hidden.
	native := 0
	for _, raw := range info.CallStack {
		f := splitFrame(raw)
		if isGoBoundaryFrame(f.symbol) {
			break
		}
		native++
		loc := f.symbol
		if f.offset != "" {
			loc += " +" + f.offset
		}
		if note, explained := lookupFrameNote(f.symbol); explained {
			fmt.Fprintf(os.Stderr, "  %-16s %-34s %s\n", f.image, loc, c(note, colorDim))
		} else {
			fmt.Fprintf(os.Stderr, "  %-16s %s\n", f.image, loc)
		}
	}

	// An exception raised through objc_exception_throw directly, rather than
	// through -[NSException raise], can arrive with no recorded stack. Say so,
	// rather than printing an empty section that looks like a bug in this code.
	if native == 0 {
		fmt.Fprintf(os.Stderr, "  %s\n", c("(no native stack recorded for this exception)", colorDim))
	}

	// This marker stands in for the dispatch frames on either side of the
	// boundary. They appear on every stack and never vary, so naming them adds
	// a line per frame and no information.
	fmt.Fprintf(os.Stderr, "  %s\n", c("-- Go/ObjC boundary --", colorDim))

	buf := make([]byte, 1<<16)
	n := runtime.Stack(buf, false)
	for _, f := range goStackFrames(string(buf[:n]), full) {
		fmt.Fprintf(os.Stderr, "  %-16s %s\n", c(f.loc, colorBold), c(f.fn, colorCyan))
	}

	// What the reader will reach for next is recover, and it does not work
	// here: the exception unwinds through C++ machinery that Go's panic
	// mechanism never sees.
	fmt.Fprintf(os.Stderr, "\n%s\n",
		c("fatal: this cannot be recovered. Validate before the call instead", colorCyan))
	fmt.Fprintf(os.Stderr, "%s\n",
		c("(RespondsToSelector, bounds checks, nil checks).", colorCyan))
	// Native frames are all shown; what full adds is the dispatch machinery on
	// the Go side and the whole goroutine dump.
	if !full {
		fmt.Fprintf(os.Stderr, "%s\n",
			c("full Go stack and goroutine dump: OBJC_EXCEPTION_STACK=full", colorDim))
	}
}

// NoteDelegatePanic reports that a panic is unwinding out of a delegate method
// that Objective-C called into. Generated delegate bodies call it from a
// deferred function when the body did not run to completion.
//
// It does not recover. Recovering here would let the Objective-C caller carry
// on as though the delegate had succeeded, which is worse than the crash: a
// nil-map write in a delegate was measured exiting 0 with "delegate succeeded"
// printed after it. The panic continues, and Go's traceback follows this line.
//
// The line exists because that traceback alone does not say which delegate ran.
// The frames between the panic and the Objective-C caller are runtime and
// purego dispatch, so the class and selector that dispatched appear nowhere.
func NoteDelegatePanic(protocolName, selector string) {
	color := colorEnabled()
	msg := fmt.Sprintf("panic in delegate %s %s", protocolName, selector)
	if color {
		fmt.Fprintf(os.Stderr, "\n%s%s%s\n", colorBold+colorRed, msg, colorReset)
	} else {
		fmt.Fprintf(os.Stderr, "\n%s\n", msg)
	}
}

// nativeFrame is one parsed line of callStackSymbols.
type nativeFrame struct {
	image  string // "CoreFoundation"
	symbol string // "___forwarding___"
	offset string // "1480"; empty when the line carried none
}

// splitFrame parses one line of callStackSymbols, such as
//
//	3   CoreFoundation   0x00000001878c5f38 ___forwarding___ + 1480
//
// The offset is kept: it is the part that stays true when the symbol name is a
// misattribution, and image+offset is what atos needs later.
func splitFrame(raw string) nativeFrame {
	fields := strings.Fields(raw)
	if len(fields) < 4 {
		return nativeFrame{symbol: raw}
	}
	f := nativeFrame{image: fields[1]}
	rest := fields[3:]
	if len(rest) >= 2 && rest[len(rest)-2] == "+" {
		f.offset = rest[len(rest)-1]
		rest = rest[:len(rest)-2]
	}
	f.symbol = strings.Join(rest, " ")
	return f
}

// lookupFrameNote returns the explanation for a native frame, if there is one.
// Matching is by substring because symbols arrive decorated, as in
// "-[NSObject(NSObject) __retain_OA]".
func lookupFrameNote(symbol string) (note string, ok bool) {
	if n, exact := frameNotes[symbol]; exact {
		return n, true
	}
	for k, n := range frameNotes {
		if strings.Contains(symbol, k) {
			return n, true
		}
	}
	return "", false
}

// isGoBoundaryFrame reports whether a native frame is the point where the stack
// crosses back into Go. Everything below it is printed from the Go traceback,
// which names files and lines.
func isGoBoundaryFrame(symbol string) bool {
	return strings.HasPrefix(symbol, "syscallX") ||
		strings.HasPrefix(symbol, "runtime.asmcgocall")
}

type goStackFrame struct{ loc, fn string }

// goStackFrames extracts caller frames from a runtime.Stack dump, dropping the
// dispatch machinery between the message send and the calling code.
func goStackFrames(dump string, full bool) []goStackFrame {
	skip := []string{
		"github.com/ebitengine/purego",
		"reflect.",
		"runtime.",
	}
	lines := strings.Split(dump, "\n")
	var out []goStackFrame
	for i := 0; i+1 < len(lines); i++ {
		fn := strings.TrimSpace(lines[i])
		loc := strings.TrimSpace(lines[i+1])
		if !strings.HasPrefix(loc, "/") || !strings.Contains(fn, "(") {
			continue
		}
		if !full {
			drop := false
			for _, p := range skip {
				if strings.HasPrefix(fn, p) {
					drop = true
					break
				}
			}
			// This package's own reporting frames, and the send that got here.
			if !drop && strings.Contains(fn, "/objc.") {
				switch {
				case strings.Contains(fn, "/objc.Send"),
					strings.Contains(fn, "/objc.fastSend"),
					strings.Contains(fn, "/objc.slowSend"),
					strings.Contains(fn, "/objc.reportException"),
					strings.Contains(fn, "/objc.GetExceptionInfo"),
					strings.Contains(fn, "/objc.initUncaughtExceptionHandler"):
					drop = true
				}
			}
			if drop {
				continue
			}
		}
		// "/path/file.go:63 +0x90" -> "file.go:63"
		loc = filepath.Base(strings.Fields(loc)[0])
		// Drop the package path, keeping "main.crashSelector()".
		if idx := strings.LastIndex(fn, "/"); idx >= 0 {
			fn = fn[idx+1:]
		}
		if !full {
			// Drop the argument values: "Send[...](0x102de9cb0, ...)" -> "Send[...]()".
			if p := strings.Index(fn, "("); p >= 0 {
				fn = fn[:p] + "()"
			}
		}
		out = append(out, goStackFrame{loc: loc, fn: fn})
		i++ // consume the location line
	}
	return out
}

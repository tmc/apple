package axuiautomation

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// AX error codes
const (
	axErrorSuccess                  = 0
	axErrorFailure                  = -25200
	axErrorIllegalArgument          = -25201
	axErrorInvalidUIElement         = -25202
	axErrorInvalidUIElementObserver = -25203
	axErrorCannotComplete           = -25204
	axErrorAttributeUnsupported     = -25205
	axErrorActionUnsupported        = -25206
	axErrorNotificationUnsupported  = -25207
	axErrorNotImplemented           = -25208
	axErrorNotificationAlreadyReg   = -25209
	axErrorNotificationNotReg       = -25210
	axErrorAPIDisabled              = -25211
	axErrorNoValue                  = -25212
	axErrorParameterizedAttrUnsup   = -25213
	axErrorNotEnoughPrecision       = -25214
)

// AXValue type constants
const (
	axValueTypeCGPoint = 1
	axValueTypeCGSize  = 2
	axValueTypeCGRect  = 3
	axValueTypeCFRange = 4
)

// CFString encoding
var cfStringEncodingUTF8 = uint32(corefoundation.KCFStringEncodingUTF8)

// CGEvent constants
const (
	cgEventLeftMouseDown   = 1
	cgEventLeftMouseUp     = 2
	cgMouseEventClickState = 1
	cgHIDEventTap          = 0
)

// Error wraps accessibility error codes
type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

// Common errors
var (
	ErrNotRunning        = errors.New("application not running")
	ErrElementNotFound   = errors.New("element not found")
	ErrTimeout           = errors.New("operation timed out")
	ErrAPIDisabled       = errors.New("accessibility API disabled")
	ErrActionUnsupported = errors.New("action not supported")
	ErrInvalidElement    = errors.New("invalid UI element")
)

// axErrorToGo converts an AXError to a Go error
func axErrorToGo(err AXError) error {
	code := int(err)
	if code == axErrorSuccess {
		return nil
	}

	var msg string
	switch code {
	case axErrorFailure:
		msg = "ax: general failure"
	case axErrorIllegalArgument:
		msg = "ax: illegal argument"
	case axErrorInvalidUIElement:
		return ErrInvalidElement
	case axErrorInvalidUIElementObserver:
		msg = "ax: invalid UI element observer"
	case axErrorCannotComplete:
		msg = "ax: cannot complete"
	case axErrorAttributeUnsupported:
		msg = "ax: attribute unsupported"
	case axErrorActionUnsupported:
		return ErrActionUnsupported
	case axErrorNotificationUnsupported:
		msg = "ax: notification unsupported"
	case axErrorNotImplemented:
		msg = "ax: not implemented"
	case axErrorNotificationAlreadyReg:
		msg = "ax: notification already registered"
	case axErrorNotificationNotReg:
		msg = "ax: notification not registered"
	case axErrorAPIDisabled:
		return ErrAPIDisabled
	case axErrorNoValue:
		msg = "ax: no value"
	case axErrorParameterizedAttrUnsup:
		msg = "ax: parameterized attribute unsupported"
	case axErrorNotEnoughPrecision:
		msg = "ax: not enough precision"
	default:
		msg = fmt.Sprintf("ax: unknown error %d", code)
	}

	return &Error{Code: code, Message: msg}
}

// stringCache provides thread-safe caching of CFString references
type stringCache struct {
	mu    sync.RWMutex
	cache map[string]uintptr
}

func newStringCache() *stringCache {
	return &stringCache{
		cache: make(map[string]uintptr),
	}
}

func (c *stringCache) get(s string) uintptr {
	c.mu.RLock()
	if ref, ok := c.cache[s]; ok {
		c.mu.RUnlock()
		return ref
	}
	c.mu.RUnlock()

	c.mu.Lock()
	defer c.mu.Unlock()

	// Double-check
	if ref, ok := c.cache[s]; ok {
		return ref
	}

	ref := corefoundation.CFStringCreateWithCString(0, s, cfStringEncodingUTF8)
	if ref != 0 {
		c.cache[s] = uintptr(ref)
	}
	return uintptr(ref)
}

func (c *stringCache) release() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, ref := range c.cache {
		if ref != 0 {
			corefoundation.CFRelease(corefoundation.CFTypeRef(ref))
		}
	}
	c.cache = make(map[string]uintptr)
}

var globalStringCache = newStringCache()

func axAttr(name string) uintptr {
	return globalStringCache.get(name)
}

// Point represents a point in 2D space
type Point struct {
	X, Y float64
}

// Size represents a size in 2D space
type Size struct {
	Width, Height float64
}

// Rect represents a rectangle in 2D space
type Rect struct {
	Origin Point
	Size   Size
}

// CGEvent functions
var (
	cgEventCreateMouseEvent     func(source uintptr, mouseType int32, x, y float64, button int32) uintptr
	cgEventCreateScrollWheelEvt func(source uintptr, units int32, wheelCount uint32, wheel1, wheel2, wheel3 int32) uintptr
	cgEventPost                 func(tap int32, event uintptr)
	cgEventSetIntegerValueField func(event uintptr, field uint32, value int64)
	cgEventCreate               func(source uintptr) uintptr
	cgEventGetDoubleValueField  func(event uintptr, field uint32) float64
	cgWarpMouseCursorPosition   func(x, y float64) int32

	cgEventsInitialized bool
	cgEventsInitOnce    sync.Once
)

func initCGEvents() {
	cgEventsInitOnce.Do(func() {
		libCG, err := purego.Dlopen("/System/Library/Frameworks/CoreGraphics.framework/CoreGraphics", purego.RTLD_GLOBAL)
		if err != nil {
			return
		}

		purego.RegisterLibFunc(&cgEventCreateMouseEvent, libCG, "CGEventCreateMouseEvent")
		purego.RegisterLibFunc(&cgEventCreateScrollWheelEvt, libCG, "CGEventCreateScrollWheelEvent")
		purego.RegisterLibFunc(&cgEventPost, libCG, "CGEventPost")
		purego.RegisterLibFunc(&cgEventSetIntegerValueField, libCG, "CGEventSetIntegerValueField")
		purego.RegisterLibFunc(&cgEventCreate, libCG, "CGEventCreate")
		purego.RegisterLibFunc(&cgEventGetDoubleValueField, libCG, "CGEventGetDoubleValueField")
		purego.RegisterLibFunc(&cgWarpMouseCursorPosition, libCG, "CGWarpMouseCursorPosition")

		cgEventsInitialized = true
	})
}

func getCurrentMousePosition() (x, y float64) {
	initCGEvents()
	if !cgEventsInitialized || cgEventCreate == nil {
		return 0, 0
	}

	event := cgEventCreate(0)
	if event == 0 {
		return 0, 0
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(event))

	x = cgEventGetDoubleValueField(event, 56) // kCGMouseEventX
	y = cgEventGetDoubleValueField(event, 57) // kCGMouseEventY
	return x, y
}

func cgEventClick(x, y int) error {
	initCGEvents()
	if !cgEventsInitialized {
		return errors.New("CGEvents not initialized")
	}

	oldX, oldY := getCurrentMousePosition()
	cgWarpMouseCursorPosition(float64(x), float64(y))
	time.Sleep(10 * time.Millisecond)

	mouseDown := cgEventCreateMouseEvent(0, cgEventLeftMouseDown, float64(x), float64(y), 0)
	if mouseDown == 0 {
		return errors.New("failed to create mouse down event")
	}
	cgEventPost(cgHIDEventTap, mouseDown)
	corefoundation.CFRelease(corefoundation.CFTypeRef(mouseDown))

	time.Sleep(50 * time.Millisecond)

	mouseUp := cgEventCreateMouseEvent(0, cgEventLeftMouseUp, float64(x), float64(y), 0)
	if mouseUp == 0 {
		return errors.New("failed to create mouse up event")
	}
	cgEventPost(cgHIDEventTap, mouseUp)
	corefoundation.CFRelease(corefoundation.CFTypeRef(mouseUp))

	time.Sleep(10 * time.Millisecond)
	cgWarpMouseCursorPosition(oldX, oldY)

	return nil
}

func cgEventDoubleClick(x, y int) error {
	initCGEvents()
	if !cgEventsInitialized {
		return errors.New("CGEvents not initialized")
	}

	oldX, oldY := getCurrentMousePosition()
	cgWarpMouseCursorPosition(float64(x), float64(y))
	time.Sleep(10 * time.Millisecond)

	// First click
	mouseDown1 := cgEventCreateMouseEvent(0, cgEventLeftMouseDown, float64(x), float64(y), 0)
	if mouseDown1 == 0 {
		return errors.New("failed to create mouse down event")
	}
	cgEventSetIntegerValueField(mouseDown1, cgMouseEventClickState, 1)
	cgEventPost(cgHIDEventTap, mouseDown1)
	corefoundation.CFRelease(corefoundation.CFTypeRef(mouseDown1))

	time.Sleep(10 * time.Millisecond)

	mouseUp1 := cgEventCreateMouseEvent(0, cgEventLeftMouseUp, float64(x), float64(y), 0)
	if mouseUp1 == 0 {
		return errors.New("failed to create mouse up event")
	}
	cgEventSetIntegerValueField(mouseUp1, cgMouseEventClickState, 1)
	cgEventPost(cgHIDEventTap, mouseUp1)
	corefoundation.CFRelease(corefoundation.CFTypeRef(mouseUp1))

	time.Sleep(50 * time.Millisecond)

	// Second click
	mouseDown2 := cgEventCreateMouseEvent(0, cgEventLeftMouseDown, float64(x), float64(y), 0)
	if mouseDown2 == 0 {
		return errors.New("failed to create mouse down event")
	}
	cgEventSetIntegerValueField(mouseDown2, cgMouseEventClickState, 2)
	cgEventPost(cgHIDEventTap, mouseDown2)
	corefoundation.CFRelease(corefoundation.CFTypeRef(mouseDown2))

	time.Sleep(10 * time.Millisecond)

	mouseUp2 := cgEventCreateMouseEvent(0, cgEventLeftMouseUp, float64(x), float64(y), 0)
	if mouseUp2 == 0 {
		return errors.New("failed to create mouse up event")
	}
	cgEventSetIntegerValueField(mouseUp2, cgMouseEventClickState, 2)
	cgEventPost(cgHIDEventTap, mouseUp2)
	corefoundation.CFRelease(corefoundation.CFTypeRef(mouseUp2))

	time.Sleep(10 * time.Millisecond)
	cgWarpMouseCursorPosition(oldX, oldY)

	return nil
}

// CFRunLoop functions
var (
	cfRunLoopGetMain      func() uintptr
	cfRunLoopGetCurrent   func() uintptr
	cfRunLoopAddSource    func(rl uintptr, source uintptr, mode uintptr)
	cfRunLoopRemoveSource func(rl uintptr, source uintptr, mode uintptr)
	cfRunLoopRunInMode    func(mode uintptr, seconds float64, returnAfterSourceHandled bool) int32
	cfRunLoopStop         func(rl uintptr)

	kCFRunLoopDefaultMode uintptr
	kCFRunLoopCommonModes uintptr

	cfRunLoopInitialized bool
	cfRunLoopInitOnce    sync.Once
)

func initCFRunLoop() {
	cfRunLoopInitOnce.Do(func() {
		libCF, err := purego.Dlopen("/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation", purego.RTLD_GLOBAL)
		if err != nil {
			return
		}

		purego.RegisterLibFunc(&cfRunLoopGetMain, libCF, "CFRunLoopGetMain")
		purego.RegisterLibFunc(&cfRunLoopGetCurrent, libCF, "CFRunLoopGetCurrent")
		purego.RegisterLibFunc(&cfRunLoopAddSource, libCF, "CFRunLoopAddSource")
		purego.RegisterLibFunc(&cfRunLoopRemoveSource, libCF, "CFRunLoopRemoveSource")
		purego.RegisterLibFunc(&cfRunLoopRunInMode, libCF, "CFRunLoopRunInMode")
		purego.RegisterLibFunc(&cfRunLoopStop, libCF, "CFRunLoopStop")

		if sym, err := purego.Dlsym(libCF, "kCFRunLoopDefaultMode"); err == nil {
			kCFRunLoopDefaultMode = *(*uintptr)(unsafe.Pointer(sym))
		}
		if sym, err := purego.Dlsym(libCF, "kCFRunLoopCommonModes"); err == nil {
			kCFRunLoopCommonModes = *(*uintptr)(unsafe.Pointer(sym))
		}

		cfRunLoopInitialized = true
	})
}

type runLoopHelper struct {
	mu       sync.Mutex
	runLoop  uintptr
	sources  []uintptr
	running  bool
	stopChan chan struct{}
}

func newRunLoopHelper() *runLoopHelper {
	initCFRunLoop()
	return &runLoopHelper{
		sources:  make([]uintptr, 0),
		stopChan: make(chan struct{}),
	}
}

func (h *runLoopHelper) addSource(source uintptr) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if !cfRunLoopInitialized {
		return
	}

	if h.runLoop == 0 {
		h.runLoop = cfRunLoopGetMain()
	}

	cfRunLoopAddSource(h.runLoop, source, kCFRunLoopDefaultMode)
	h.sources = append(h.sources, source)
}

func (h *runLoopHelper) removeSource(source uintptr) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if !cfRunLoopInitialized || h.runLoop == 0 {
		return
	}

	cfRunLoopRemoveSource(h.runLoop, source, kCFRunLoopDefaultMode)

	for i, s := range h.sources {
		if s == source {
			h.sources = append(h.sources[:i], h.sources[i+1:]...)
			break
		}
	}
}

func (h *runLoopHelper) runOnce(timeout time.Duration) {
	if !cfRunLoopInitialized {
		return
	}
	cfRunLoopRunInMode(kCFRunLoopDefaultMode, timeout.Seconds(), true)
}

func (h *runLoopHelper) stop() {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.running {
		close(h.stopChan)
		h.running = false
	}

	if cfRunLoopInitialized && h.runLoop != 0 {
		cfRunLoopStop(h.runLoop)
	}
}

func (h *runLoopHelper) cleanup() {
	h.mu.Lock()
	defer h.mu.Unlock()

	if !cfRunLoopInitialized || h.runLoop == 0 {
		return
	}

	for _, source := range h.sources {
		cfRunLoopRemoveSource(h.runLoop, source, kCFRunLoopDefaultMode)
	}
	h.sources = nil
}

// CFBoolean helper
var cfBooleanGetValue func(uintptr) bool
var cfBooleanInitOnce sync.Once

func initCFBoolean() {
	cfBooleanInitOnce.Do(func() {
		libCF, err := purego.Dlopen("/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation", purego.RTLD_GLOBAL)
		if err != nil {
			return
		}
		purego.RegisterLibFunc(&cfBooleanGetValue, libCF, "CFBooleanGetValue")
	})
}

// SpinRunLoop briefly spins the current thread's CFRunLoop to allow pending
// accessibility IPC replies to be delivered. Many AX attributes (notably
// AXMenuBar) require at least one run-loop iteration to return a value in
// processes that do not otherwise run a persistent run loop (e.g. CLI tools).
func SpinRunLoop(d time.Duration) {
	initCFRunLoop()
	if !cfRunLoopInitialized {
		return
	}
	cfRunLoopRunInMode(kCFRunLoopDefaultMode, d.Seconds(), false)
}

// SpinMainRunLoop briefly runs the NSRunLoop on the main thread for the given
// duration. Unlike SpinRunLoop (which only pumps CFRunLoop), this also drains
// the GCD main queue, so DispatchMainSafe callbacks and AppKit UI events are
// processed. Must be called from the main OS thread.
func SpinMainRunLoop(d time.Duration) {
	deadline := foundation.GetNSDateClass().Alloc().InitWithTimeIntervalSinceNow(d.Seconds())
	foundation.GetRunLoopClass().MainRunLoop().RunUntilDate(deadline)
}

// cfStringTypeID holds the CFTypeID for CFString, used to guard cfStringToGo.
var (
	cfStringTypeID    uintptr
	cfGetTypeID       func(uintptr) uintptr
	cfStringGetTypeID func() uintptr
	cfTypeIDInitOnce  sync.Once
)

func initCFTypeID() {
	cfTypeIDInitOnce.Do(func() {
		lib, err := purego.Dlopen("/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation", purego.RTLD_GLOBAL)
		if err != nil {
			return
		}
		purego.RegisterLibFunc(&cfGetTypeID, lib, "CFGetTypeID")
		purego.RegisterLibFunc(&cfStringGetTypeID, lib, "CFStringGetTypeID")
		cfStringTypeID = cfStringGetTypeID()
	})
}

// String conversion helpers
func cfStringToGo(cfStr uintptr) string {
	if cfStr == 0 {
		return ""
	}

	// Guard against non-string CF types (e.g. CFNumber) that would crash.
	initCFTypeID()
	if cfGetTypeID != nil && cfStringGetTypeID != nil {
		if cfGetTypeID(cfStr) != cfStringTypeID {
			return ""
		}
	}

	ref := corefoundation.CFStringRef(cfStr)
	length := corefoundation.CFStringGetLength(ref)
	if length == 0 {
		return ""
	}

	bufSize := length*4 + 1
	buf := make([]byte, bufSize)

	result := corefoundation.CFStringGetCString(ref, &buf[0], int(bufSize), cfStringEncodingUTF8)
	if !result {
		return ""
	}

	for i, b := range buf {
		if b == 0 {
			return string(buf[:i])
		}
	}
	return string(buf)
}

func cfArrayToSlice(cfArray uintptr) []uintptr {
	if cfArray == 0 {
		return nil
	}

	count := corefoundation.CFArrayGetCount(corefoundation.CFArrayRef(cfArray))
	if count == 0 {
		return nil
	}

	result := make([]uintptr, count)
	for i := 0; i < count; i++ {
		ptr := corefoundation.CFArrayGetValueAtIndex(corefoundation.CFArrayRef(cfArray), i)
		result[i] = uintptr(ptr)
	}
	return result
}

// AX attribute helpers
func getAXAttributeString(element AXUIElementRef, attrName string) string {
	var value uintptr
	err := AXUIElementCopyAttributeValue(element, axAttr(attrName), &value)
	if int(err) != axErrorSuccess || value == 0 {
		return ""
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(unsafe.Pointer(value)))

	return cfStringToGo(value)
}

func getAXAttributeBool(element AXUIElementRef, attrName string) bool {
	initCFBoolean()

	var value uintptr
	err := AXUIElementCopyAttributeValue(element, axAttr(attrName), &value)
	if int(err) != axErrorSuccess || value == 0 {
		return false
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(unsafe.Pointer(value)))

	if cfBooleanGetValue == nil {
		return false
	}
	return cfBooleanGetValue(value)
}

func getAXAttributePoint(element AXUIElementRef, attrName string) (Point, bool) {
	var value uintptr
	err := AXUIElementCopyAttributeValue(element, axAttr(attrName), &value)
	if int(err) != axErrorSuccess || value == 0 {
		return Point{}, false
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(unsafe.Pointer(value)))

	var point Point
	if AXValueGetValue(AXValueRef(value), AXValueType(axValueTypeCGPoint), unsafe.Pointer(&point)) {
		return point, true
	}
	return Point{}, false
}

func getAXAttributeSize(element AXUIElementRef, attrName string) (Size, bool) {
	var value uintptr
	err := AXUIElementCopyAttributeValue(element, axAttr(attrName), &value)
	if int(err) != axErrorSuccess || value == 0 {
		return Size{}, false
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(unsafe.Pointer(value)))

	var size Size
	if AXValueGetValue(AXValueRef(value), AXValueType(axValueTypeCGSize), unsafe.Pointer(&size)) {
		return size, true
	}
	return Size{}, false
}

func getAXAttributeElements(element AXUIElementRef, attrName string) []AXUIElementRef {
	var value uintptr
	err := AXUIElementCopyAttributeValue(element, axAttr(attrName), &value)
	if int(err) != axErrorSuccess || value == 0 {
		return nil
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(unsafe.Pointer(value)))

	refs := cfArrayToSlice(value)
	result := make([]AXUIElementRef, len(refs))
	for i, ref := range refs {
		corefoundation.CFRetain(corefoundation.CFTypeRef(unsafe.Pointer(ref)))
		result[i] = AXUIElementRef(ref)
	}
	return result
}

// Keyboard event functions
var (
	cgEventCreateKeyboardEvent      func(source uintptr, keycode uint16, keyDown bool) uintptr
	cgEventKeyboardSetUnicodeString func(event uintptr, length uint32, unicodeString *uint16)
	cgEventSetFlags                 func(event uintptr, flags uint64)
	cgEventGetFlags                 func(event uintptr) uint64

	keyboardEventsInitialized bool
	keyboardEventsInitOnce    sync.Once
)

// CGEvent flag constants
const (
	cgEventFlagMaskShift     = 0x00020000
	cgEventFlagMaskControl   = 0x00040000
	cgEventFlagMaskAlternate = 0x00080000
	cgEventFlagMaskCommand   = 0x00100000
)

// Key codes
const (
	keyCodeEscape     = 0x35
	keyCodeReturn     = 0x24
	keyCodeTab        = 0x30
	keyCodeDelete     = 0x33
	keyCodeArrowUp    = 0x7E
	keyCodeArrowDown  = 0x7D
	keyCodeArrowLeft  = 0x7B
	keyCodeArrowRight = 0x7C
	keyCodeN          = 0x2D
	keyCodeG          = 0x05
)

func initKeyboardEvents() {
	keyboardEventsInitOnce.Do(func() {
		initCGEvents() // Ensure CGEvents are initialized first

		libCG, err := purego.Dlopen("/System/Library/Frameworks/CoreGraphics.framework/CoreGraphics", purego.RTLD_GLOBAL)
		if err != nil {
			return
		}

		purego.RegisterLibFunc(&cgEventCreateKeyboardEvent, libCG, "CGEventCreateKeyboardEvent")
		purego.RegisterLibFunc(&cgEventKeyboardSetUnicodeString, libCG, "CGEventKeyboardSetUnicodeString")
		purego.RegisterLibFunc(&cgEventSetFlags, libCG, "CGEventSetFlags")
		purego.RegisterLibFunc(&cgEventGetFlags, libCG, "CGEventGetFlags")

		keyboardEventsInitialized = true
	})
}

// SendEscape sends an escape key press.
func SendEscape() error {
	initKeyboardEvents()
	if !keyboardEventsInitialized {
		return errors.New("keyboard events not initialized")
	}

	keyDown := cgEventCreateKeyboardEvent(0, keyCodeEscape, true)
	if keyDown == 0 {
		return errors.New("failed to create key down event")
	}
	cgEventPost(cgHIDEventTap, keyDown)
	corefoundation.CFRelease(corefoundation.CFTypeRef(keyDown))

	time.Sleep(10 * time.Millisecond)

	keyUp := cgEventCreateKeyboardEvent(0, keyCodeEscape, false)
	if keyUp == 0 {
		return errors.New("failed to create key up event")
	}
	cgEventPost(cgHIDEventTap, keyUp)
	corefoundation.CFRelease(corefoundation.CFTypeRef(keyUp))

	return nil
}

// SendReturn sends a return key press.
func SendReturn() error {
	initKeyboardEvents()
	if !keyboardEventsInitialized {
		return errors.New("keyboard events not initialized")
	}

	keyDown := cgEventCreateKeyboardEvent(0, keyCodeReturn, true)
	if keyDown == 0 {
		return errors.New("failed to create key down event")
	}
	cgEventPost(cgHIDEventTap, keyDown)
	corefoundation.CFRelease(corefoundation.CFTypeRef(keyDown))

	time.Sleep(10 * time.Millisecond)

	keyUp := cgEventCreateKeyboardEvent(0, keyCodeReturn, false)
	if keyUp == 0 {
		return errors.New("failed to create key up event")
	}
	cgEventPost(cgHIDEventTap, keyUp)
	corefoundation.CFRelease(corefoundation.CFTypeRef(keyUp))

	return nil
}

// SendKeyCombo sends a key combination with modifiers.
// Modifiers: shift, control, option, command
func SendKeyCombo(keyCode uint16, shift, control, option, command bool) error {
	initKeyboardEvents()
	if !keyboardEventsInitialized {
		return errors.New("keyboard events not initialized")
	}

	var flags uint64
	if shift {
		flags |= cgEventFlagMaskShift
	}
	if control {
		flags |= cgEventFlagMaskControl
	}
	if option {
		flags |= cgEventFlagMaskAlternate
	}
	if command {
		flags |= cgEventFlagMaskCommand
	}

	keyDown := cgEventCreateKeyboardEvent(0, keyCode, true)
	if keyDown == 0 {
		return errors.New("failed to create key down event")
	}
	cgEventSetFlags(keyDown, flags)
	cgEventPost(cgHIDEventTap, keyDown)
	corefoundation.CFRelease(corefoundation.CFTypeRef(keyDown))

	time.Sleep(10 * time.Millisecond)

	keyUp := cgEventCreateKeyboardEvent(0, keyCode, false)
	if keyUp == 0 {
		return errors.New("failed to create key up event")
	}
	cgEventSetFlags(keyUp, flags)
	cgEventPost(cgHIDEventTap, keyUp)
	corefoundation.CFRelease(corefoundation.CFTypeRef(keyUp))

	return nil
}

// SendCmdShiftG sends Command+Shift+G (Go to Folder in save dialogs).
func SendCmdShiftG() error {
	return SendKeyCombo(keyCodeG, true, false, false, true)
}

// SendCmdN sends Command+N (New file/item).
func SendCmdN() error {
	return SendKeyCombo(keyCodeN, false, false, false, true)
}

// SendTab sends a Tab key press.
func SendTab() error {
	return SendKeyCombo(keyCodeTab, false, false, false, false)
}

// SendShiftTab sends Shift+Tab.
func SendShiftTab() error {
	return SendKeyCombo(keyCodeTab, true, false, false, false)
}

// SendDelete sends a Delete (backspace) key press.
func SendDelete() error {
	return SendKeyCombo(keyCodeDelete, false, false, false, false)
}

// SendArrowUp sends an Up arrow key press.
func SendArrowUp() error {
	return SendKeyCombo(keyCodeArrowUp, false, false, false, false)
}

// SendArrowDown sends a Down arrow key press.
func SendArrowDown() error {
	return SendKeyCombo(keyCodeArrowDown, false, false, false, false)
}

// SendArrowLeft sends a Left arrow key press.
func SendArrowLeft() error {
	return SendKeyCombo(keyCodeArrowLeft, false, false, false, false)
}

// SendArrowRight sends a Right arrow key press.
func SendArrowRight() error {
	return SendKeyCombo(keyCodeArrowRight, false, false, false, false)
}

// TypeText types text into the currently focused element by sending keyboard events.
// It supports printable ASCII characters. Use Element.TypeText for element-targeted input.
func TypeText(text string) error {
	initKeyboardEvents()
	if !keyboardEventsInitialized {
		return errors.New("keyboard events not initialized")
	}

	for _, ch := range text {
		if err := typeChar(ch); err != nil {
			return fmt.Errorf("typing %q: %w", ch, err)
		}
		time.Sleep(5 * time.Millisecond)
	}
	return nil
}

// typeChar sends a single character as keyboard events using CGEventKeyboardSetUnicodeString.
func typeChar(ch rune) error {
	initKeyboardEvents()

	// Create a key-down event with key code 0; we'll set unicode string directly.
	keyDown := cgEventCreateKeyboardEvent(0, 0, true)
	if keyDown == 0 {
		return errors.New("failed to create key down event")
	}

	// Set the unicode character directly.
	u16 := uint16(ch)
	cgEventKeyboardSetUnicodeString(keyDown, 1, &u16)
	cgEventPost(cgHIDEventTap, keyDown)
	corefoundation.CFRelease(corefoundation.CFTypeRef(keyDown))

	time.Sleep(5 * time.Millisecond)

	keyUp := cgEventCreateKeyboardEvent(0, 0, false)
	if keyUp == 0 {
		return errors.New("failed to create key up event")
	}
	cgEventKeyboardSetUnicodeString(keyUp, 1, &u16)
	cgEventPost(cgHIDEventTap, keyUp)
	corefoundation.CFRelease(corefoundation.CFTypeRef(keyUp))

	return nil
}

// ScrollDirection specifies the direction to scroll.
type ScrollDirection int

const (
	ScrollUp ScrollDirection = iota
	ScrollDown
	ScrollLeft
	ScrollRight
)

// cgScrollWheel posts a CGEvent scroll wheel event at the given position.
// units: 0 = pixel, 1 = line. amount: positive scrolls up/left, negative down/right.
func cgScrollWheel(x, y int, direction ScrollDirection, amount int) error {
	initCGEvents()
	if !cgEventsInitialized || cgEventCreateScrollWheelEvt == nil {
		return errors.New("CGEvents not initialized")
	}

	// CGEventCreateScrollWheelEvent: wheel1=vertical (positive=up), wheel2=horizontal (positive=left)
	var wheel1, wheel2 int32
	switch direction {
	case ScrollUp:
		wheel1 = int32(amount)
	case ScrollDown:
		wheel1 = -int32(amount)
	case ScrollLeft:
		wheel2 = int32(amount)
	case ScrollRight:
		wheel2 = -int32(amount)
	}

	// Move mouse to target position first so scroll hits the right element
	oldX, oldY := getCurrentMousePosition()
	cgWarpMouseCursorPosition(float64(x), float64(y))
	time.Sleep(10 * time.Millisecond)

	// kCGScrollEventUnitLine = 1
	evt := cgEventCreateScrollWheelEvt(0, 1, 2, wheel1, wheel2, 0)
	if evt == 0 {
		cgWarpMouseCursorPosition(oldX, oldY)
		return errors.New("failed to create scroll wheel event")
	}
	cgEventPost(cgHIDEventTap, evt)
	corefoundation.CFRelease(corefoundation.CFTypeRef(evt))

	time.Sleep(10 * time.Millisecond)
	cgWarpMouseCursorPosition(oldX, oldY)
	return nil
}

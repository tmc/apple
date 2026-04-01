// Code generated from Apple documentation for CoreVideo. DO NOT EDIT.

package corevideo

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

type unavailableSymbolError struct {
	symbol     string
	introduced string
	cause      error
}

func (e *unavailableSymbolError) Error() string {
	if e == nil {
		return ""
	}
	if e.introduced != "" {
		return fmt.Sprintf("CoreVideo: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("CoreVideo: symbol %s unavailable on this system", e.symbol)
}

func (e *unavailableSymbolError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	if err != nil {
		return err
	}
	if frameworkHandle == 0 {
		return fmt.Errorf("CoreVideo: symbol %s unavailable because the framework could not be loaded", name)
	}
	return missingSymbolError(name, introduced, nil)
}

// registerFunc resolves a framework symbol and registers it as a Go function.
func registerFunc(fptr any, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			*errDst = fmt.Errorf("CoreVideo: register symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	*errDst = nil
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	*dst = sym
	*errDst = nil
}

var _cVBufferCopyAttachment func(buffer CVBufferRef, key corefoundation.CFStringRef, attachmentMode *CVAttachmentMode) corefoundation.CFTypeRef
var _cVBufferCopyAttachmentErr error

func tryCVBufferCopyAttachment(buffer CVBufferRef, key corefoundation.CFStringRef, attachmentMode *CVAttachmentMode) (corefoundation.CFTypeRef, error) {
	if _cVBufferCopyAttachment == nil {
		return 0, symbolCallError("CVBufferCopyAttachment", "12.0", _cVBufferCopyAttachmentErr)
	}
	return _cVBufferCopyAttachment(buffer, key, attachmentMode), nil
}

// CVBufferCopyAttachment returns a copy of an attachment from a Core Video buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferCopyAttachment(_:_:_:)
func CVBufferCopyAttachment(buffer CVBufferRef, key corefoundation.CFStringRef, attachmentMode *CVAttachmentMode) corefoundation.CFTypeRef {
	result, callErr := tryCVBufferCopyAttachment(buffer, key, attachmentMode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVBufferCopyAttachments func(buffer CVBufferRef, attachmentMode CVAttachmentMode) corefoundation.CFDictionaryRef
var _cVBufferCopyAttachmentsErr error

func tryCVBufferCopyAttachments(buffer CVBufferRef, attachmentMode CVAttachmentMode) (corefoundation.CFDictionaryRef, error) {
	if _cVBufferCopyAttachments == nil {
		return 0, symbolCallError("CVBufferCopyAttachments", "12.0", _cVBufferCopyAttachmentsErr)
	}
	return _cVBufferCopyAttachments(buffer, attachmentMode), nil
}

// CVBufferCopyAttachments returns a copy of all attachments from a Core Video buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferCopyAttachments(_:_:)
func CVBufferCopyAttachments(buffer CVBufferRef, attachmentMode CVAttachmentMode) corefoundation.CFDictionaryRef {
	result, callErr := tryCVBufferCopyAttachments(buffer, attachmentMode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVBufferHasAttachment func(buffer CVBufferRef, key corefoundation.CFStringRef) bool
var _cVBufferHasAttachmentErr error

func tryCVBufferHasAttachment(buffer CVBufferRef, key corefoundation.CFStringRef) (bool, error) {
	if _cVBufferHasAttachment == nil {
		return false, symbolCallError("CVBufferHasAttachment", "12.0", _cVBufferHasAttachmentErr)
	}
	return _cVBufferHasAttachment(buffer, key), nil
}

// CVBufferHasAttachment returns a Boolean value that indicates whether a Core Video buffer contains a specified attachment.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferHasAttachment(_:_:)
func CVBufferHasAttachment(buffer CVBufferRef, key corefoundation.CFStringRef) bool {
	result, callErr := tryCVBufferHasAttachment(buffer, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVBufferPropagateAttachments func(sourceBuffer CVBufferRef, destinationBuffer CVBufferRef)
var _cVBufferPropagateAttachmentsErr error

func tryCVBufferPropagateAttachments(sourceBuffer CVBufferRef, destinationBuffer CVBufferRef) error {
	if _cVBufferPropagateAttachments == nil {
		return symbolCallError("CVBufferPropagateAttachments", "10.4", _cVBufferPropagateAttachmentsErr)
	}
	_cVBufferPropagateAttachments(sourceBuffer, destinationBuffer)
	return nil
}

// CVBufferPropagateAttachments copies all attachments that Core Video can propagate from one buffer to another.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferPropagateAttachments(_:_:)
func CVBufferPropagateAttachments(sourceBuffer CVBufferRef, destinationBuffer CVBufferRef) {
	if callErr := tryCVBufferPropagateAttachments(sourceBuffer, destinationBuffer); callErr != nil {
		panic(callErr)
	}
}

var _cVBufferRelease func(buffer CVBufferRef)
var _cVBufferReleaseErr error

func tryCVBufferRelease(buffer CVBufferRef) error {
	if _cVBufferRelease == nil {
		return symbolCallError("CVBufferRelease", "10.4", _cVBufferReleaseErr)
	}
	_cVBufferRelease(buffer)
	return nil
}

// CVBufferRelease releases a Core Video buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferRelease
func CVBufferRelease(buffer CVBufferRef) {
	if callErr := tryCVBufferRelease(buffer); callErr != nil {
		panic(callErr)
	}
}

var _cVBufferRemoveAllAttachments func(buffer CVBufferRef)
var _cVBufferRemoveAllAttachmentsErr error

func tryCVBufferRemoveAllAttachments(buffer CVBufferRef) error {
	if _cVBufferRemoveAllAttachments == nil {
		return symbolCallError("CVBufferRemoveAllAttachments", "10.4", _cVBufferRemoveAllAttachmentsErr)
	}
	_cVBufferRemoveAllAttachments(buffer)
	return nil
}

// CVBufferRemoveAllAttachments removes all attachments from a Core Video buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferRemoveAllAttachments(_:)
func CVBufferRemoveAllAttachments(buffer CVBufferRef) {
	if callErr := tryCVBufferRemoveAllAttachments(buffer); callErr != nil {
		panic(callErr)
	}
}

var _cVBufferRemoveAttachment func(buffer CVBufferRef, key corefoundation.CFStringRef)
var _cVBufferRemoveAttachmentErr error

func tryCVBufferRemoveAttachment(buffer CVBufferRef, key corefoundation.CFStringRef) error {
	if _cVBufferRemoveAttachment == nil {
		return symbolCallError("CVBufferRemoveAttachment", "10.4", _cVBufferRemoveAttachmentErr)
	}
	_cVBufferRemoveAttachment(buffer, key)
	return nil
}

// CVBufferRemoveAttachment removes the attachment you specify from a Core Video buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferRemoveAttachment(_:_:)
func CVBufferRemoveAttachment(buffer CVBufferRef, key corefoundation.CFStringRef) {
	if callErr := tryCVBufferRemoveAttachment(buffer, key); callErr != nil {
		panic(callErr)
	}
}

var _cVBufferRetain func(buffer CVBufferRef) CVBufferRef
var _cVBufferRetainErr error

func tryCVBufferRetain(buffer CVBufferRef) (CVBufferRef, error) {
	if _cVBufferRetain == nil {
		return 0, symbolCallError("CVBufferRetain", "10.4", _cVBufferRetainErr)
	}
	return _cVBufferRetain(buffer), nil
}

// CVBufferRetain retains a Core Video buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferRetain
func CVBufferRetain(buffer CVBufferRef) CVBufferRef {
	result, callErr := tryCVBufferRetain(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVBufferSetAttachment func(buffer CVBufferRef, key corefoundation.CFStringRef, value corefoundation.CFTypeRef, attachmentMode CVAttachmentMode)
var _cVBufferSetAttachmentErr error

func tryCVBufferSetAttachment(buffer CVBufferRef, key corefoundation.CFStringRef, value corefoundation.CFTypeRef, attachmentMode CVAttachmentMode) error {
	if _cVBufferSetAttachment == nil {
		return symbolCallError("CVBufferSetAttachment", "10.4", _cVBufferSetAttachmentErr)
	}
	_cVBufferSetAttachment(buffer, key, value, attachmentMode)
	return nil
}

// CVBufferSetAttachment sets or adds an attachment to a Core Video buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferSetAttachment(_:_:_:_:)
func CVBufferSetAttachment(buffer CVBufferRef, key corefoundation.CFStringRef, value corefoundation.CFTypeRef, attachmentMode CVAttachmentMode) {
	if callErr := tryCVBufferSetAttachment(buffer, key, value, attachmentMode); callErr != nil {
		panic(callErr)
	}
}

var _cVBufferSetAttachments func(buffer CVBufferRef, theAttachments corefoundation.CFDictionaryRef, attachmentMode CVAttachmentMode)
var _cVBufferSetAttachmentsErr error

func tryCVBufferSetAttachments(buffer CVBufferRef, theAttachments corefoundation.CFDictionaryRef, attachmentMode CVAttachmentMode) error {
	if _cVBufferSetAttachments == nil {
		return symbolCallError("CVBufferSetAttachments", "10.4", _cVBufferSetAttachmentsErr)
	}
	_cVBufferSetAttachments(buffer, theAttachments, attachmentMode)
	return nil
}

// CVBufferSetAttachments sets a dictionary of attachments on a Core Video buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferSetAttachments(_:_:_:)
func CVBufferSetAttachments(buffer CVBufferRef, theAttachments corefoundation.CFDictionaryRef, attachmentMode CVAttachmentMode) {
	if callErr := tryCVBufferSetAttachments(buffer, theAttachments, attachmentMode); callErr != nil {
		panic(callErr)
	}
}

var _cVColorPrimariesGetIntegerCodePointForString func(colorPrimariesString corefoundation.CFStringRef) int
var _cVColorPrimariesGetIntegerCodePointForStringErr error

func tryCVColorPrimariesGetIntegerCodePointForString(colorPrimariesString corefoundation.CFStringRef) (int, error) {
	if _cVColorPrimariesGetIntegerCodePointForString == nil {
		return 0, symbolCallError("CVColorPrimariesGetIntegerCodePointForString", "10.13", _cVColorPrimariesGetIntegerCodePointForStringErr)
	}
	return _cVColorPrimariesGetIntegerCodePointForString(colorPrimariesString), nil
}

// CVColorPrimariesGetIntegerCodePointForString returns the standard integer code point corresponding to the Core Video color primaries constant string that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVColorPrimariesGetIntegerCodePointForString(_:)
func CVColorPrimariesGetIntegerCodePointForString(colorPrimariesString corefoundation.CFStringRef) int {
	result, callErr := tryCVColorPrimariesGetIntegerCodePointForString(colorPrimariesString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVColorPrimariesGetStringForIntegerCodePoint func(colorPrimariesCodePoint int) corefoundation.CFStringRef
var _cVColorPrimariesGetStringForIntegerCodePointErr error

func tryCVColorPrimariesGetStringForIntegerCodePoint(colorPrimariesCodePoint int) (corefoundation.CFStringRef, error) {
	if _cVColorPrimariesGetStringForIntegerCodePoint == nil {
		return 0, symbolCallError("CVColorPrimariesGetStringForIntegerCodePoint", "10.13", _cVColorPrimariesGetStringForIntegerCodePointErr)
	}
	return _cVColorPrimariesGetStringForIntegerCodePoint(colorPrimariesCodePoint), nil
}

// CVColorPrimariesGetStringForIntegerCodePoint returns the Core Video color primaries string corresponding to the standard integer code point that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVColorPrimariesGetStringForIntegerCodePoint(_:)
func CVColorPrimariesGetStringForIntegerCodePoint(colorPrimariesCodePoint int) corefoundation.CFStringRef {
	result, callErr := tryCVColorPrimariesGetStringForIntegerCodePoint(colorPrimariesCodePoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVDisplayLinkCreateWithActiveCGDisplays func(displayLinkOut *CVDisplayLinkRef) CVReturn
var _cVDisplayLinkCreateWithActiveCGDisplaysErr error

func tryCVDisplayLinkCreateWithActiveCGDisplays(displayLinkOut *CVDisplayLinkRef) (CVReturn, error) {
	if _cVDisplayLinkCreateWithActiveCGDisplays == nil {
		return *new(CVReturn), symbolCallError("CVDisplayLinkCreateWithActiveCGDisplays", "10.4", _cVDisplayLinkCreateWithActiveCGDisplaysErr)
	}
	return _cVDisplayLinkCreateWithActiveCGDisplays(displayLinkOut), nil
}

// CVDisplayLinkCreateWithActiveCGDisplays creates a display link capable of being used with all active displays.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkCreateWithActiveCGDisplays(_:)
func CVDisplayLinkCreateWithActiveCGDisplays(displayLinkOut *CVDisplayLinkRef) CVReturn {
	result, callErr := tryCVDisplayLinkCreateWithActiveCGDisplays(displayLinkOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVDisplayLinkCreateWithCGDisplay func(displayID uint32, displayLinkOut *CVDisplayLinkRef) CVReturn
var _cVDisplayLinkCreateWithCGDisplayErr error

func tryCVDisplayLinkCreateWithCGDisplay(displayID uint32, displayLinkOut *CVDisplayLinkRef) (CVReturn, error) {
	if _cVDisplayLinkCreateWithCGDisplay == nil {
		return *new(CVReturn), symbolCallError("CVDisplayLinkCreateWithCGDisplay", "10.4", _cVDisplayLinkCreateWithCGDisplayErr)
	}
	return _cVDisplayLinkCreateWithCGDisplay(displayID, displayLinkOut), nil
}

// CVDisplayLinkCreateWithCGDisplay creates a display link for a single display.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkCreateWithCGDisplay(_:_:)
func CVDisplayLinkCreateWithCGDisplay(displayID uint32, displayLinkOut *CVDisplayLinkRef) CVReturn {
	result, callErr := tryCVDisplayLinkCreateWithCGDisplay(displayID, displayLinkOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVDisplayLinkCreateWithCGDisplays func(displayArray *uint32, count int, displayLinkOut *CVDisplayLinkRef) CVReturn
var _cVDisplayLinkCreateWithCGDisplaysErr error

func tryCVDisplayLinkCreateWithCGDisplays(displayArray *uint32, count int, displayLinkOut *CVDisplayLinkRef) (CVReturn, error) {
	if _cVDisplayLinkCreateWithCGDisplays == nil {
		return *new(CVReturn), symbolCallError("CVDisplayLinkCreateWithCGDisplays", "10.4", _cVDisplayLinkCreateWithCGDisplaysErr)
	}
	return _cVDisplayLinkCreateWithCGDisplays(displayArray, count, displayLinkOut), nil
}

// CVDisplayLinkCreateWithCGDisplays creates a display link for an array of displays.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkCreateWithCGDisplays(_:_:_:)
func CVDisplayLinkCreateWithCGDisplays(displayArray *uint32, count int, displayLinkOut *CVDisplayLinkRef) CVReturn {
	result, callErr := tryCVDisplayLinkCreateWithCGDisplays(displayArray, count, displayLinkOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVDisplayLinkCreateWithOpenGLDisplayMask func(mask coregraphics.CGOpenGLDisplayMask, displayLinkOut *CVDisplayLinkRef) CVReturn
var _cVDisplayLinkCreateWithOpenGLDisplayMaskErr error

func tryCVDisplayLinkCreateWithOpenGLDisplayMask(mask coregraphics.CGOpenGLDisplayMask, displayLinkOut *CVDisplayLinkRef) (CVReturn, error) {
	if _cVDisplayLinkCreateWithOpenGLDisplayMask == nil {
		return *new(CVReturn), symbolCallError("CVDisplayLinkCreateWithOpenGLDisplayMask", "10.4", _cVDisplayLinkCreateWithOpenGLDisplayMaskErr)
	}
	return _cVDisplayLinkCreateWithOpenGLDisplayMask(mask, displayLinkOut), nil
}

// CVDisplayLinkCreateWithOpenGLDisplayMask creates a display link from an OpenGL display mask.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkCreateWithOpenGLDisplayMask(_:_:)
func CVDisplayLinkCreateWithOpenGLDisplayMask(mask coregraphics.CGOpenGLDisplayMask, displayLinkOut *CVDisplayLinkRef) CVReturn {
	result, callErr := tryCVDisplayLinkCreateWithOpenGLDisplayMask(mask, displayLinkOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVDisplayLinkGetActualOutputVideoRefreshPeriod func(displayLink CVDisplayLinkRef) float64
var _cVDisplayLinkGetActualOutputVideoRefreshPeriodErr error

func tryCVDisplayLinkGetActualOutputVideoRefreshPeriod(displayLink CVDisplayLinkRef) (float64, error) {
	if _cVDisplayLinkGetActualOutputVideoRefreshPeriod == nil {
		return 0.0, symbolCallError("CVDisplayLinkGetActualOutputVideoRefreshPeriod", "10.4", _cVDisplayLinkGetActualOutputVideoRefreshPeriodErr)
	}
	return _cVDisplayLinkGetActualOutputVideoRefreshPeriod(displayLink), nil
}

// CVDisplayLinkGetActualOutputVideoRefreshPeriod retrieves the actual output refresh period of a display as measured by the system time.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetActualOutputVideoRefreshPeriod(_:)
func CVDisplayLinkGetActualOutputVideoRefreshPeriod(displayLink CVDisplayLinkRef) float64 {
	result, callErr := tryCVDisplayLinkGetActualOutputVideoRefreshPeriod(displayLink)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVDisplayLinkGetCurrentCGDisplay func(displayLink CVDisplayLinkRef) uint32
var _cVDisplayLinkGetCurrentCGDisplayErr error

func tryCVDisplayLinkGetCurrentCGDisplay(displayLink CVDisplayLinkRef) (uint32, error) {
	if _cVDisplayLinkGetCurrentCGDisplay == nil {
		return 0, symbolCallError("CVDisplayLinkGetCurrentCGDisplay", "10.4", _cVDisplayLinkGetCurrentCGDisplayErr)
	}
	return _cVDisplayLinkGetCurrentCGDisplay(displayLink), nil
}

// CVDisplayLinkGetCurrentCGDisplay gets the current display associated with a display link.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetCurrentCGDisplay(_:)
func CVDisplayLinkGetCurrentCGDisplay(displayLink CVDisplayLinkRef) uint32 {
	result, callErr := tryCVDisplayLinkGetCurrentCGDisplay(displayLink)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVDisplayLinkGetCurrentTime func(displayLink CVDisplayLinkRef, outTime *CVTimeStamp) CVReturn
var _cVDisplayLinkGetCurrentTimeErr error

func tryCVDisplayLinkGetCurrentTime(displayLink CVDisplayLinkRef, outTime *CVTimeStamp) (CVReturn, error) {
	if _cVDisplayLinkGetCurrentTime == nil {
		return *new(CVReturn), symbolCallError("CVDisplayLinkGetCurrentTime", "10.4", _cVDisplayLinkGetCurrentTimeErr)
	}
	return _cVDisplayLinkGetCurrentTime(displayLink, outTime), nil
}

// CVDisplayLinkGetCurrentTime retrieves the current (“now”) time of a given display link.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetCurrentTime(_:_:)
func CVDisplayLinkGetCurrentTime(displayLink CVDisplayLinkRef, outTime *CVTimeStamp) CVReturn {
	result, callErr := tryCVDisplayLinkGetCurrentTime(displayLink, outTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVDisplayLinkGetNominalOutputVideoRefreshPeriod func(displayLink CVDisplayLinkRef) CVTime
var _cVDisplayLinkGetNominalOutputVideoRefreshPeriodErr error

func tryCVDisplayLinkGetNominalOutputVideoRefreshPeriod(displayLink CVDisplayLinkRef) (CVTime, error) {
	if _cVDisplayLinkGetNominalOutputVideoRefreshPeriod == nil {
		return CVTime{}, symbolCallError("CVDisplayLinkGetNominalOutputVideoRefreshPeriod", "10.4", _cVDisplayLinkGetNominalOutputVideoRefreshPeriodErr)
	}
	return _cVDisplayLinkGetNominalOutputVideoRefreshPeriod(displayLink), nil
}

// CVDisplayLinkGetNominalOutputVideoRefreshPeriod retrieves the nominal refresh period of a display link.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetNominalOutputVideoRefreshPeriod(_:)
func CVDisplayLinkGetNominalOutputVideoRefreshPeriod(displayLink CVDisplayLinkRef) CVTime {
	result, callErr := tryCVDisplayLinkGetNominalOutputVideoRefreshPeriod(displayLink)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVDisplayLinkGetOutputVideoLatency func(displayLink CVDisplayLinkRef) CVTime
var _cVDisplayLinkGetOutputVideoLatencyErr error

func tryCVDisplayLinkGetOutputVideoLatency(displayLink CVDisplayLinkRef) (CVTime, error) {
	if _cVDisplayLinkGetOutputVideoLatency == nil {
		return CVTime{}, symbolCallError("CVDisplayLinkGetOutputVideoLatency", "10.4", _cVDisplayLinkGetOutputVideoLatencyErr)
	}
	return _cVDisplayLinkGetOutputVideoLatency(displayLink), nil
}

// CVDisplayLinkGetOutputVideoLatency retrieves the nominal latency of a display link.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetOutputVideoLatency(_:)
func CVDisplayLinkGetOutputVideoLatency(displayLink CVDisplayLinkRef) CVTime {
	result, callErr := tryCVDisplayLinkGetOutputVideoLatency(displayLink)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVDisplayLinkGetTypeID func() uint
var _cVDisplayLinkGetTypeIDErr error

func tryCVDisplayLinkGetTypeID() (uint, error) {
	if _cVDisplayLinkGetTypeID == nil {
		return 0, symbolCallError("CVDisplayLinkGetTypeID", "10.4", _cVDisplayLinkGetTypeIDErr)
	}
	return _cVDisplayLinkGetTypeID(), nil
}

// CVDisplayLinkGetTypeID obtains the Core Foundation ID for the display link data type.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetTypeID()
func CVDisplayLinkGetTypeID() uint {
	result, callErr := tryCVDisplayLinkGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVDisplayLinkIsRunning func(displayLink CVDisplayLinkRef) bool
var _cVDisplayLinkIsRunningErr error

func tryCVDisplayLinkIsRunning(displayLink CVDisplayLinkRef) (bool, error) {
	if _cVDisplayLinkIsRunning == nil {
		return false, symbolCallError("CVDisplayLinkIsRunning", "10.4", _cVDisplayLinkIsRunningErr)
	}
	return _cVDisplayLinkIsRunning(displayLink), nil
}

// CVDisplayLinkIsRunning indicates whether a given display link is running.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkIsRunning(_:)
func CVDisplayLinkIsRunning(displayLink CVDisplayLinkRef) bool {
	result, callErr := tryCVDisplayLinkIsRunning(displayLink)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVDisplayLinkRelease func(displayLink CVDisplayLinkRef)
var _cVDisplayLinkReleaseErr error

func tryCVDisplayLinkRelease(displayLink CVDisplayLinkRef) error {
	if _cVDisplayLinkRelease == nil {
		return symbolCallError("CVDisplayLinkRelease", "10.4", _cVDisplayLinkReleaseErr)
	}
	_cVDisplayLinkRelease(displayLink)
	return nil
}

// CVDisplayLinkRelease releases a display link.
//
// Deprecated: Deprecated since macOS 15.0.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkRelease
func CVDisplayLinkRelease(displayLink CVDisplayLinkRef) {
	if callErr := tryCVDisplayLinkRelease(displayLink); callErr != nil {
		panic(callErr)
	}
}

var _cVDisplayLinkRetain func(displayLink CVDisplayLinkRef) CVDisplayLinkRef
var _cVDisplayLinkRetainErr error

func tryCVDisplayLinkRetain(displayLink CVDisplayLinkRef) (CVDisplayLinkRef, error) {
	if _cVDisplayLinkRetain == nil {
		return 0, symbolCallError("CVDisplayLinkRetain", "10.4", _cVDisplayLinkRetainErr)
	}
	return _cVDisplayLinkRetain(displayLink), nil
}

// CVDisplayLinkRetain retains a display link.
//
// Deprecated: Deprecated since macOS 15.0.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkRetain
func CVDisplayLinkRetain(displayLink CVDisplayLinkRef) CVDisplayLinkRef {
	result, callErr := tryCVDisplayLinkRetain(displayLink)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVDisplayLinkSetCurrentCGDisplay func(displayLink CVDisplayLinkRef, displayID uint32) CVReturn
var _cVDisplayLinkSetCurrentCGDisplayErr error

func tryCVDisplayLinkSetCurrentCGDisplay(displayLink CVDisplayLinkRef, displayID uint32) (CVReturn, error) {
	if _cVDisplayLinkSetCurrentCGDisplay == nil {
		return *new(CVReturn), symbolCallError("CVDisplayLinkSetCurrentCGDisplay", "10.4", _cVDisplayLinkSetCurrentCGDisplayErr)
	}
	return _cVDisplayLinkSetCurrentCGDisplay(displayLink, displayID), nil
}

// CVDisplayLinkSetCurrentCGDisplay sets the current display of a display link.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkSetCurrentCGDisplay(_:_:)
func CVDisplayLinkSetCurrentCGDisplay(displayLink CVDisplayLinkRef, displayID uint32) CVReturn {
	result, callErr := tryCVDisplayLinkSetCurrentCGDisplay(displayLink, displayID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVDisplayLinkSetCurrentCGDisplayFromOpenGLContext func(displayLink CVDisplayLinkRef, cglContext unsafe.Pointer, cglPixelFormat unsafe.Pointer) CVReturn
var _cVDisplayLinkSetCurrentCGDisplayFromOpenGLContextErr error

func tryCVDisplayLinkSetCurrentCGDisplayFromOpenGLContext(displayLink CVDisplayLinkRef, cglContext unsafe.Pointer, cglPixelFormat unsafe.Pointer) (CVReturn, error) {
	if _cVDisplayLinkSetCurrentCGDisplayFromOpenGLContext == nil {
		return *new(CVReturn), symbolCallError("CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext", "10.4", _cVDisplayLinkSetCurrentCGDisplayFromOpenGLContextErr)
	}
	return _cVDisplayLinkSetCurrentCGDisplayFromOpenGLContext(displayLink, cglContext, cglPixelFormat), nil
}

// CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext selects the display link most optimal for the current renderer of an OpenGL context.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext(_:_:_:)
func CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext(displayLink CVDisplayLinkRef, cglContext unsafe.Pointer, cglPixelFormat unsafe.Pointer) CVReturn {
	result, callErr := tryCVDisplayLinkSetCurrentCGDisplayFromOpenGLContext(displayLink, cglContext, cglPixelFormat)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVDisplayLinkSetOutputCallback func(displayLink CVDisplayLinkRef, callback CVDisplayLinkOutputCallback, userInfo unsafe.Pointer) CVReturn
var _cVDisplayLinkSetOutputCallbackErr error

func tryCVDisplayLinkSetOutputCallback(displayLink CVDisplayLinkRef, callback CVDisplayLinkOutputCallback, userInfo unsafe.Pointer) (CVReturn, error) {
	if _cVDisplayLinkSetOutputCallback == nil {
		return *new(CVReturn), symbolCallError("CVDisplayLinkSetOutputCallback", "10.4", _cVDisplayLinkSetOutputCallbackErr)
	}
	return _cVDisplayLinkSetOutputCallback(displayLink, callback, userInfo), nil
}

// CVDisplayLinkSetOutputCallback sets the renderer output callback function.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkSetOutputCallback(_:_:_:)
func CVDisplayLinkSetOutputCallback(displayLink CVDisplayLinkRef, callback CVDisplayLinkOutputCallback, userInfo unsafe.Pointer) CVReturn {
	result, callErr := tryCVDisplayLinkSetOutputCallback(displayLink, callback, userInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVDisplayLinkSetOutputHandler func(displayLink CVDisplayLinkRef, handler unsafe.Pointer) CVReturn
var _cVDisplayLinkSetOutputHandlerErr error

func tryCVDisplayLinkSetOutputHandler(displayLink CVDisplayLinkRef, handler CVDisplayLinkOutputHandler) (CVReturn, error) {
	if _cVDisplayLinkSetOutputHandler == nil {
		return *new(CVReturn), symbolCallError("CVDisplayLinkSetOutputHandler", "10.4", _cVDisplayLinkSetOutputHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.IObject, arg1 *CVTimeStamp, arg2 *CVTimeStamp, arg3 uint64, arg4 *uint64) int {
		return handler(arg0, arg1, arg2, arg3, arg4)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _cVDisplayLinkSetOutputHandler(displayLink, _block0), nil
}

// CVDisplayLinkSetOutputHandler.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkSetOutputHandler(_:_:)
func CVDisplayLinkSetOutputHandler(displayLink CVDisplayLinkRef, handler CVDisplayLinkOutputHandler) CVReturn {
	result, callErr := tryCVDisplayLinkSetOutputHandler(displayLink, handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVDisplayLinkStart func(displayLink CVDisplayLinkRef) CVReturn
var _cVDisplayLinkStartErr error

func tryCVDisplayLinkStart(displayLink CVDisplayLinkRef) (CVReturn, error) {
	if _cVDisplayLinkStart == nil {
		return *new(CVReturn), symbolCallError("CVDisplayLinkStart", "10.4", _cVDisplayLinkStartErr)
	}
	return _cVDisplayLinkStart(displayLink), nil
}

// CVDisplayLinkStart activates a display link.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkStart(_:)
func CVDisplayLinkStart(displayLink CVDisplayLinkRef) CVReturn {
	result, callErr := tryCVDisplayLinkStart(displayLink)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVDisplayLinkStop func(displayLink CVDisplayLinkRef) CVReturn
var _cVDisplayLinkStopErr error

func tryCVDisplayLinkStop(displayLink CVDisplayLinkRef) (CVReturn, error) {
	if _cVDisplayLinkStop == nil {
		return *new(CVReturn), symbolCallError("CVDisplayLinkStop", "10.4", _cVDisplayLinkStopErr)
	}
	return _cVDisplayLinkStop(displayLink), nil
}

// CVDisplayLinkStop stops a display link.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkStop(_:)
func CVDisplayLinkStop(displayLink CVDisplayLinkRef) CVReturn {
	result, callErr := tryCVDisplayLinkStop(displayLink)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVDisplayLinkTranslateTime func(displayLink CVDisplayLinkRef, inTime *CVTimeStamp, outTime *CVTimeStamp) CVReturn
var _cVDisplayLinkTranslateTimeErr error

func tryCVDisplayLinkTranslateTime(displayLink CVDisplayLinkRef, inTime *CVTimeStamp, outTime *CVTimeStamp) (CVReturn, error) {
	if _cVDisplayLinkTranslateTime == nil {
		return *new(CVReturn), symbolCallError("CVDisplayLinkTranslateTime", "10.4", _cVDisplayLinkTranslateTimeErr)
	}
	return _cVDisplayLinkTranslateTime(displayLink, inTime, outTime), nil
}

// CVDisplayLinkTranslateTime translates the time in the display link’s time base from one representation to another.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkTranslateTime(_:_:_:)
func CVDisplayLinkTranslateTime(displayLink CVDisplayLinkRef, inTime *CVTimeStamp, outTime *CVTimeStamp) CVReturn {
	result, callErr := tryCVDisplayLinkTranslateTime(displayLink, inTime, outTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVGetCurrentHostTime func() uint64
var _cVGetCurrentHostTimeErr error

func tryCVGetCurrentHostTime() (uint64, error) {
	if _cVGetCurrentHostTime == nil {
		return 0, symbolCallError("CVGetCurrentHostTime", "10.4", _cVGetCurrentHostTimeErr)
	}
	return _cVGetCurrentHostTime(), nil
}

// CVGetCurrentHostTime returns the current system time.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVGetCurrentHostTime()
func CVGetCurrentHostTime() uint64 {
	result, callErr := tryCVGetCurrentHostTime()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVGetHostClockFrequency func() float64
var _cVGetHostClockFrequencyErr error

func tryCVGetHostClockFrequency() (float64, error) {
	if _cVGetHostClockFrequency == nil {
		return 0.0, symbolCallError("CVGetHostClockFrequency", "10.4", _cVGetHostClockFrequencyErr)
	}
	return _cVGetHostClockFrequency(), nil
}

// CVGetHostClockFrequency returns the frequency of updates to the system time.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVGetHostClockFrequency()
func CVGetHostClockFrequency() float64 {
	result, callErr := tryCVGetHostClockFrequency()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVGetHostClockMinimumTimeDelta func() uint32
var _cVGetHostClockMinimumTimeDeltaErr error

func tryCVGetHostClockMinimumTimeDelta() (uint32, error) {
	if _cVGetHostClockMinimumTimeDelta == nil {
		return 0, symbolCallError("CVGetHostClockMinimumTimeDelta", "10.4", _cVGetHostClockMinimumTimeDeltaErr)
	}
	return _cVGetHostClockMinimumTimeDelta(), nil
}

// CVGetHostClockMinimumTimeDelta returns the smallest possible increment in the system time.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVGetHostClockMinimumTimeDelta()
func CVGetHostClockMinimumTimeDelta() uint32 {
	result, callErr := tryCVGetHostClockMinimumTimeDelta()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVImageBufferCreateColorSpaceFromAttachments func(attachments corefoundation.CFDictionaryRef) coregraphics.CGColorSpaceRef
var _cVImageBufferCreateColorSpaceFromAttachmentsErr error

func tryCVImageBufferCreateColorSpaceFromAttachments(attachments corefoundation.CFDictionaryRef) (coregraphics.CGColorSpaceRef, error) {
	if _cVImageBufferCreateColorSpaceFromAttachments == nil {
		return 0, symbolCallError("CVImageBufferCreateColorSpaceFromAttachments", "10.8", _cVImageBufferCreateColorSpaceFromAttachmentsErr)
	}
	return _cVImageBufferCreateColorSpaceFromAttachments(attachments), nil
}

// CVImageBufferCreateColorSpaceFromAttachments attempts to create a Core Graphics color space from the image buffer’s attachments that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVImageBufferCreateColorSpaceFromAttachments(_:)
func CVImageBufferCreateColorSpaceFromAttachments(attachments corefoundation.CFDictionaryRef) coregraphics.CGColorSpaceRef {
	result, callErr := tryCVImageBufferCreateColorSpaceFromAttachments(attachments)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVImageBufferGetCleanRect func(imageBuffer CVImageBufferRef) corefoundation.CGRect
var _cVImageBufferGetCleanRectErr error

func tryCVImageBufferGetCleanRect(imageBuffer CVImageBufferRef) (corefoundation.CGRect, error) {
	if _cVImageBufferGetCleanRect == nil {
		return corefoundation.CGRect{}, symbolCallError("CVImageBufferGetCleanRect", "10.4", _cVImageBufferGetCleanRectErr)
	}
	return _cVImageBufferGetCleanRect(imageBuffer), nil
}

// CVImageBufferGetCleanRect returns the source rectangle of a Core Video image buffer that represents the clean aperture of the buffer in encoded pixels.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVImageBufferGetCleanRect(_:)
func CVImageBufferGetCleanRect(imageBuffer CVImageBufferRef) corefoundation.CGRect {
	result, callErr := tryCVImageBufferGetCleanRect(imageBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVImageBufferGetColorSpace func(imageBuffer CVImageBufferRef) coregraphics.CGColorSpaceRef
var _cVImageBufferGetColorSpaceErr error

func tryCVImageBufferGetColorSpace(imageBuffer CVImageBufferRef) (coregraphics.CGColorSpaceRef, error) {
	if _cVImageBufferGetColorSpace == nil {
		return 0, symbolCallError("CVImageBufferGetColorSpace", "10.4", _cVImageBufferGetColorSpaceErr)
	}
	return _cVImageBufferGetColorSpace(imageBuffer), nil
}

// CVImageBufferGetColorSpace returns the color space of a Core Video image buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVImageBufferGetColorSpace(_:)
func CVImageBufferGetColorSpace(imageBuffer CVImageBufferRef) coregraphics.CGColorSpaceRef {
	result, callErr := tryCVImageBufferGetColorSpace(imageBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVImageBufferGetDisplaySize func(imageBuffer CVImageBufferRef) corefoundation.CGSize
var _cVImageBufferGetDisplaySizeErr error

func tryCVImageBufferGetDisplaySize(imageBuffer CVImageBufferRef) (corefoundation.CGSize, error) {
	if _cVImageBufferGetDisplaySize == nil {
		return corefoundation.CGSize{}, symbolCallError("CVImageBufferGetDisplaySize", "10.4", _cVImageBufferGetDisplaySizeErr)
	}
	return _cVImageBufferGetDisplaySize(imageBuffer), nil
}

// CVImageBufferGetDisplaySize returns the nominal output display size, in square pixels, of a Core Video image buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVImageBufferGetDisplaySize(_:)
func CVImageBufferGetDisplaySize(imageBuffer CVImageBufferRef) corefoundation.CGSize {
	result, callErr := tryCVImageBufferGetDisplaySize(imageBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVImageBufferGetEncodedSize func(imageBuffer CVImageBufferRef) corefoundation.CGSize
var _cVImageBufferGetEncodedSizeErr error

func tryCVImageBufferGetEncodedSize(imageBuffer CVImageBufferRef) (corefoundation.CGSize, error) {
	if _cVImageBufferGetEncodedSize == nil {
		return corefoundation.CGSize{}, symbolCallError("CVImageBufferGetEncodedSize", "10.4", _cVImageBufferGetEncodedSizeErr)
	}
	return _cVImageBufferGetEncodedSize(imageBuffer), nil
}

// CVImageBufferGetEncodedSize returns the full encoded dimensions of a Core Video image buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVImageBufferGetEncodedSize(_:)
func CVImageBufferGetEncodedSize(imageBuffer CVImageBufferRef) corefoundation.CGSize {
	result, callErr := tryCVImageBufferGetEncodedSize(imageBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVImageBufferIsFlipped func(imageBuffer CVImageBufferRef) bool
var _cVImageBufferIsFlippedErr error

func tryCVImageBufferIsFlipped(imageBuffer CVImageBufferRef) (bool, error) {
	if _cVImageBufferIsFlipped == nil {
		return false, symbolCallError("CVImageBufferIsFlipped", "10.4", _cVImageBufferIsFlippedErr)
	}
	return _cVImageBufferIsFlipped(imageBuffer), nil
}

// CVImageBufferIsFlipped returns a Boolean value indicating whether the image is vertically flipped.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVImageBufferIsFlipped(_:)
func CVImageBufferIsFlipped(imageBuffer CVImageBufferRef) bool {
	result, callErr := tryCVImageBufferIsFlipped(imageBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVIsCompressedPixelFormatAvailable func(pixelFormatType uint32) bool
var _cVIsCompressedPixelFormatAvailableErr error

func tryCVIsCompressedPixelFormatAvailable(pixelFormatType uint32) (bool, error) {
	if _cVIsCompressedPixelFormatAvailable == nil {
		return false, symbolCallError("CVIsCompressedPixelFormatAvailable", "12.0", _cVIsCompressedPixelFormatAvailableErr)
	}
	return _cVIsCompressedPixelFormatAvailable(pixelFormatType), nil
}

// CVIsCompressedPixelFormatAvailable.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVIsCompressedPixelFormatAvailable(_:)
func CVIsCompressedPixelFormatAvailable(pixelFormatType uint32) bool {
	result, callErr := tryCVIsCompressedPixelFormatAvailable(pixelFormatType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVMetalBufferCacheCreate func(allocator corefoundation.CFAllocatorRef, cacheAttributes corefoundation.CFDictionaryRef, metalDevice metal.MTLDeviceObject, cacheOut *CVMetalBufferCacheRef) CVReturn
var _cVMetalBufferCacheCreateErr error

func tryCVMetalBufferCacheCreate(allocator corefoundation.CFAllocatorRef, cacheAttributes corefoundation.CFDictionaryRef, metalDevice metal.MTLDeviceObject, cacheOut *CVMetalBufferCacheRef) (CVReturn, error) {
	if _cVMetalBufferCacheCreate == nil {
		return *new(CVReturn), symbolCallError("CVMetalBufferCacheCreate", "15.0", _cVMetalBufferCacheCreateErr)
	}
	return _cVMetalBufferCacheCreate(allocator, cacheAttributes, metalDevice, cacheOut), nil
}

// CVMetalBufferCacheCreate.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferCacheCreate(_:_:_:_:)
func CVMetalBufferCacheCreate(allocator corefoundation.CFAllocatorRef, cacheAttributes corefoundation.CFDictionaryRef, metalDevice metal.MTLDeviceObject, cacheOut *CVMetalBufferCacheRef) CVReturn {
	result, callErr := tryCVMetalBufferCacheCreate(allocator, cacheAttributes, metalDevice, cacheOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVMetalBufferCacheCreateBufferFromImage func(allocator corefoundation.CFAllocatorRef, bufferCache CVMetalBufferCacheRef, imageBuffer CVImageBufferRef, bufferOut *CVMetalBufferRef) CVReturn
var _cVMetalBufferCacheCreateBufferFromImageErr error

func tryCVMetalBufferCacheCreateBufferFromImage(allocator corefoundation.CFAllocatorRef, bufferCache CVMetalBufferCacheRef, imageBuffer CVImageBufferRef, bufferOut *CVMetalBufferRef) (CVReturn, error) {
	if _cVMetalBufferCacheCreateBufferFromImage == nil {
		return *new(CVReturn), symbolCallError("CVMetalBufferCacheCreateBufferFromImage", "15.0", _cVMetalBufferCacheCreateBufferFromImageErr)
	}
	return _cVMetalBufferCacheCreateBufferFromImage(allocator, bufferCache, imageBuffer, bufferOut), nil
}

// CVMetalBufferCacheCreateBufferFromImage.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferCacheCreateBufferFromImage(_:_:_:_:)
func CVMetalBufferCacheCreateBufferFromImage(allocator corefoundation.CFAllocatorRef, bufferCache CVMetalBufferCacheRef, imageBuffer CVImageBufferRef, bufferOut *CVMetalBufferRef) CVReturn {
	result, callErr := tryCVMetalBufferCacheCreateBufferFromImage(allocator, bufferCache, imageBuffer, bufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVMetalBufferCacheFlush func(bufferCache CVMetalBufferCacheRef, options CVOptionFlags)
var _cVMetalBufferCacheFlushErr error

func tryCVMetalBufferCacheFlush(bufferCache CVMetalBufferCacheRef, options CVOptionFlags) error {
	if _cVMetalBufferCacheFlush == nil {
		return symbolCallError("CVMetalBufferCacheFlush", "15.0", _cVMetalBufferCacheFlushErr)
	}
	_cVMetalBufferCacheFlush(bufferCache, options)
	return nil
}

// CVMetalBufferCacheFlush.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferCacheFlush(_:_:)
func CVMetalBufferCacheFlush(bufferCache CVMetalBufferCacheRef, options CVOptionFlags) {
	if callErr := tryCVMetalBufferCacheFlush(bufferCache, options); callErr != nil {
		panic(callErr)
	}
}

var _cVMetalBufferCacheGetTypeID func() uint
var _cVMetalBufferCacheGetTypeIDErr error

func tryCVMetalBufferCacheGetTypeID() (uint, error) {
	if _cVMetalBufferCacheGetTypeID == nil {
		return 0, symbolCallError("CVMetalBufferCacheGetTypeID", "15.0", _cVMetalBufferCacheGetTypeIDErr)
	}
	return _cVMetalBufferCacheGetTypeID(), nil
}

// CVMetalBufferCacheGetTypeID.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferCacheGetTypeID()
func CVMetalBufferCacheGetTypeID() uint {
	result, callErr := tryCVMetalBufferCacheGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVMetalBufferGetBuffer func(buffer CVMetalBufferRef) unsafe.Pointer
var _cVMetalBufferGetBufferErr error

func tryCVMetalBufferGetBuffer(buffer CVMetalBufferRef) (metal.MTLBufferObject, error) {
	if _cVMetalBufferGetBuffer == nil {
		return metal.MTLBufferObject{}, symbolCallError("CVMetalBufferGetBuffer", "15.0", _cVMetalBufferGetBufferErr)
	}
	rv := _cVMetalBufferGetBuffer(buffer)
	return metal.MTLBufferObjectFromID(objc.IDFrom(rv)), nil
}

// CVMetalBufferGetBuffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferGetBuffer(_:)
func CVMetalBufferGetBuffer(buffer CVMetalBufferRef) metal.MTLBufferObject {
	result, callErr := tryCVMetalBufferGetBuffer(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVMetalBufferGetTypeID func() uint
var _cVMetalBufferGetTypeIDErr error

func tryCVMetalBufferGetTypeID() (uint, error) {
	if _cVMetalBufferGetTypeID == nil {
		return 0, symbolCallError("CVMetalBufferGetTypeID", "15.0", _cVMetalBufferGetTypeIDErr)
	}
	return _cVMetalBufferGetTypeID(), nil
}

// CVMetalBufferGetTypeID.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferGetTypeID()
func CVMetalBufferGetTypeID() uint {
	result, callErr := tryCVMetalBufferGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVMetalTextureCacheCreate func(allocator corefoundation.CFAllocatorRef, cacheAttributes corefoundation.CFDictionaryRef, metalDevice metal.MTLDeviceObject, textureAttributes corefoundation.CFDictionaryRef, cacheOut *CVMetalTextureCacheRef) CVReturn
var _cVMetalTextureCacheCreateErr error

func tryCVMetalTextureCacheCreate(allocator corefoundation.CFAllocatorRef, cacheAttributes corefoundation.CFDictionaryRef, metalDevice metal.MTLDeviceObject, textureAttributes corefoundation.CFDictionaryRef, cacheOut *CVMetalTextureCacheRef) (CVReturn, error) {
	if _cVMetalTextureCacheCreate == nil {
		return *new(CVReturn), symbolCallError("CVMetalTextureCacheCreate", "10.11", _cVMetalTextureCacheCreateErr)
	}
	return _cVMetalTextureCacheCreate(allocator, cacheAttributes, metalDevice, textureAttributes, cacheOut), nil
}

// CVMetalTextureCacheCreate creates a new texture cache.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureCacheCreate(_:_:_:_:_:)
func CVMetalTextureCacheCreate(allocator corefoundation.CFAllocatorRef, cacheAttributes corefoundation.CFDictionaryRef, metalDevice metal.MTLDeviceObject, textureAttributes corefoundation.CFDictionaryRef, cacheOut *CVMetalTextureCacheRef) CVReturn {
	result, callErr := tryCVMetalTextureCacheCreate(allocator, cacheAttributes, metalDevice, textureAttributes, cacheOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVMetalTextureCacheCreateTextureFromImage func(allocator corefoundation.CFAllocatorRef, textureCache CVMetalTextureCacheRef, sourceImage CVImageBufferRef, textureAttributes corefoundation.CFDictionaryRef, pixelFormat metal.MTLPixelFormat, width uintptr, height uintptr, planeIndex uintptr, textureOut *CVMetalTextureRef) CVReturn
var _cVMetalTextureCacheCreateTextureFromImageErr error

func tryCVMetalTextureCacheCreateTextureFromImage(allocator corefoundation.CFAllocatorRef, textureCache CVMetalTextureCacheRef, sourceImage CVImageBufferRef, textureAttributes corefoundation.CFDictionaryRef, pixelFormat metal.MTLPixelFormat, width uintptr, height uintptr, planeIndex uintptr, textureOut *CVMetalTextureRef) (CVReturn, error) {
	if _cVMetalTextureCacheCreateTextureFromImage == nil {
		return *new(CVReturn), symbolCallError("CVMetalTextureCacheCreateTextureFromImage", "10.11", _cVMetalTextureCacheCreateTextureFromImageErr)
	}
	return _cVMetalTextureCacheCreateTextureFromImage(allocator, textureCache, sourceImage, textureAttributes, pixelFormat, width, height, planeIndex, textureOut), nil
}

// CVMetalTextureCacheCreateTextureFromImage creates a Core Video Metal texture buffer from an existing image buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureCacheCreateTextureFromImage(_:_:_:_:_:_:_:_:_:)
func CVMetalTextureCacheCreateTextureFromImage(allocator corefoundation.CFAllocatorRef, textureCache CVMetalTextureCacheRef, sourceImage CVImageBufferRef, textureAttributes corefoundation.CFDictionaryRef, pixelFormat metal.MTLPixelFormat, width uintptr, height uintptr, planeIndex uintptr, textureOut *CVMetalTextureRef) CVReturn {
	result, callErr := tryCVMetalTextureCacheCreateTextureFromImage(allocator, textureCache, sourceImage, textureAttributes, pixelFormat, width, height, planeIndex, textureOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVMetalTextureCacheFlush func(textureCache CVMetalTextureCacheRef, options CVOptionFlags)
var _cVMetalTextureCacheFlushErr error

func tryCVMetalTextureCacheFlush(textureCache CVMetalTextureCacheRef, options CVOptionFlags) error {
	if _cVMetalTextureCacheFlush == nil {
		return symbolCallError("CVMetalTextureCacheFlush", "10.11", _cVMetalTextureCacheFlushErr)
	}
	_cVMetalTextureCacheFlush(textureCache, options)
	return nil
}

// CVMetalTextureCacheFlush manually flushes the contents of the provided texture cache.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureCacheFlush(_:_:)
func CVMetalTextureCacheFlush(textureCache CVMetalTextureCacheRef, options CVOptionFlags) {
	if callErr := tryCVMetalTextureCacheFlush(textureCache, options); callErr != nil {
		panic(callErr)
	}
}

var _cVMetalTextureCacheGetTypeID func() uint
var _cVMetalTextureCacheGetTypeIDErr error

func tryCVMetalTextureCacheGetTypeID() (uint, error) {
	if _cVMetalTextureCacheGetTypeID == nil {
		return 0, symbolCallError("CVMetalTextureCacheGetTypeID", "10.11", _cVMetalTextureCacheGetTypeIDErr)
	}
	return _cVMetalTextureCacheGetTypeID(), nil
}

// CVMetalTextureCacheGetTypeID returns the Core Foundation type identifier for a Core Video Metal texture cache.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureCacheGetTypeID()
func CVMetalTextureCacheGetTypeID() uint {
	result, callErr := tryCVMetalTextureCacheGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVMetalTextureGetCleanTexCoords func(image CVMetalTextureRef, lowerLeft unsafe.Pointer, lowerRight unsafe.Pointer, upperRight unsafe.Pointer, upperLeft unsafe.Pointer)
var _cVMetalTextureGetCleanTexCoordsErr error

func tryCVMetalTextureGetCleanTexCoords(image CVMetalTextureRef, lowerLeft unsafe.Pointer, lowerRight unsafe.Pointer, upperRight unsafe.Pointer, upperLeft unsafe.Pointer) error {
	if _cVMetalTextureGetCleanTexCoords == nil {
		return symbolCallError("CVMetalTextureGetCleanTexCoords", "10.11", _cVMetalTextureGetCleanTexCoordsErr)
	}
	_cVMetalTextureGetCleanTexCoords(image, lowerLeft, lowerRight, upperRight, upperLeft)
	return nil
}

// CVMetalTextureGetCleanTexCoords returns convenient normalized texture coordinates for the part of the image that should be displayed.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureGetCleanTexCoords(_:_:_:_:_:)
func CVMetalTextureGetCleanTexCoords(image CVMetalTextureRef, lowerLeft unsafe.Pointer, lowerRight unsafe.Pointer, upperRight unsafe.Pointer, upperLeft unsafe.Pointer) {
	if callErr := tryCVMetalTextureGetCleanTexCoords(image, lowerLeft, lowerRight, upperRight, upperLeft); callErr != nil {
		panic(callErr)
	}
}

var _cVMetalTextureGetTexture func(image CVMetalTextureRef) unsafe.Pointer
var _cVMetalTextureGetTextureErr error

func tryCVMetalTextureGetTexture(image CVMetalTextureRef) (metal.MTLTextureObject, error) {
	if _cVMetalTextureGetTexture == nil {
		return metal.MTLTextureObject{}, symbolCallError("CVMetalTextureGetTexture", "10.11", _cVMetalTextureGetTextureErr)
	}
	rv := _cVMetalTextureGetTexture(image)
	return metal.MTLTextureObjectFromID(objc.IDFrom(rv)), nil
}

// CVMetalTextureGetTexture returns the Metal texture object for the image buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureGetTexture(_:)
func CVMetalTextureGetTexture(image CVMetalTextureRef) metal.MTLTextureObject {
	result, callErr := tryCVMetalTextureGetTexture(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVMetalTextureGetTypeID func() uint
var _cVMetalTextureGetTypeIDErr error

func tryCVMetalTextureGetTypeID() (uint, error) {
	if _cVMetalTextureGetTypeID == nil {
		return 0, symbolCallError("CVMetalTextureGetTypeID", "10.11", _cVMetalTextureGetTypeIDErr)
	}
	return _cVMetalTextureGetTypeID(), nil
}

// CVMetalTextureGetTypeID returns the Core Foundation type identifier for a CoreVideo Metal texture-based image buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureGetTypeID()
func CVMetalTextureGetTypeID() uint {
	result, callErr := tryCVMetalTextureGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVMetalTextureIsFlipped func(image CVMetalTextureRef) bool
var _cVMetalTextureIsFlippedErr error

func tryCVMetalTextureIsFlipped(image CVMetalTextureRef) (bool, error) {
	if _cVMetalTextureIsFlipped == nil {
		return false, symbolCallError("CVMetalTextureIsFlipped", "10.11", _cVMetalTextureIsFlippedErr)
	}
	return _cVMetalTextureIsFlipped(image), nil
}

// CVMetalTextureIsFlipped returns a Boolean value indicating whether the texture image is vertically flipped.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureIsFlipped(_:)
func CVMetalTextureIsFlipped(image CVMetalTextureRef) bool {
	result, callErr := tryCVMetalTextureIsFlipped(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVOpenGLBufferPoolRelease func(openGLBufferPool CVOpenGLBufferPoolRef)
var _cVOpenGLBufferPoolReleaseErr error

func tryCVOpenGLBufferPoolRelease(openGLBufferPool CVOpenGLBufferPoolRef) error {
	if _cVOpenGLBufferPoolRelease == nil {
		return symbolCallError("CVOpenGLBufferPoolRelease", "10.4", _cVOpenGLBufferPoolReleaseErr)
	}
	_cVOpenGLBufferPoolRelease(openGLBufferPool)
	return nil
}

// CVOpenGLBufferPoolRelease releases an OpenGL buffer pool.
//
// Deprecated: Deprecated since macOS 10.14.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferPoolRelease
func CVOpenGLBufferPoolRelease(openGLBufferPool CVOpenGLBufferPoolRef) {
	if callErr := tryCVOpenGLBufferPoolRelease(openGLBufferPool); callErr != nil {
		panic(callErr)
	}
}

var _cVOpenGLBufferPoolRetain func(openGLBufferPool CVOpenGLBufferPoolRef) CVOpenGLBufferPoolRef
var _cVOpenGLBufferPoolRetainErr error

func tryCVOpenGLBufferPoolRetain(openGLBufferPool CVOpenGLBufferPoolRef) (CVOpenGLBufferPoolRef, error) {
	if _cVOpenGLBufferPoolRetain == nil {
		return 0, symbolCallError("CVOpenGLBufferPoolRetain", "10.4", _cVOpenGLBufferPoolRetainErr)
	}
	return _cVOpenGLBufferPoolRetain(openGLBufferPool), nil
}

// CVOpenGLBufferPoolRetain retains an OpenGL buffer pool.
//
// Deprecated: Deprecated since macOS 10.14.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferPoolRetain
func CVOpenGLBufferPoolRetain(openGLBufferPool CVOpenGLBufferPoolRef) CVOpenGLBufferPoolRef {
	result, callErr := tryCVOpenGLBufferPoolRetain(openGLBufferPool)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVOpenGLBufferRelease func(buffer CVOpenGLBufferRef)
var _cVOpenGLBufferReleaseErr error

func tryCVOpenGLBufferRelease(buffer CVOpenGLBufferRef) error {
	if _cVOpenGLBufferRelease == nil {
		return symbolCallError("CVOpenGLBufferRelease", "10.4", _cVOpenGLBufferReleaseErr)
	}
	_cVOpenGLBufferRelease(buffer)
	return nil
}

// CVOpenGLBufferRelease releases a Core Video OpenGL buffer.
//
// Deprecated: Deprecated since macOS 10.14.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferRelease
func CVOpenGLBufferRelease(buffer CVOpenGLBufferRef) {
	if callErr := tryCVOpenGLBufferRelease(buffer); callErr != nil {
		panic(callErr)
	}
}

var _cVOpenGLBufferRetain func(buffer CVOpenGLBufferRef) CVOpenGLBufferRef
var _cVOpenGLBufferRetainErr error

func tryCVOpenGLBufferRetain(buffer CVOpenGLBufferRef) (CVOpenGLBufferRef, error) {
	if _cVOpenGLBufferRetain == nil {
		return 0, symbolCallError("CVOpenGLBufferRetain", "10.4", _cVOpenGLBufferRetainErr)
	}
	return _cVOpenGLBufferRetain(buffer), nil
}

// CVOpenGLBufferRetain retains a Core Video OpenGL buffer.
//
// Deprecated: Deprecated since macOS 10.14.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferRetain
func CVOpenGLBufferRetain(buffer CVOpenGLBufferRef) CVOpenGLBufferRef {
	result, callErr := tryCVOpenGLBufferRetain(buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVOpenGLTextureCacheRelease func(textureCache CVOpenGLTextureCacheRef)
var _cVOpenGLTextureCacheReleaseErr error

func tryCVOpenGLTextureCacheRelease(textureCache CVOpenGLTextureCacheRef) error {
	if _cVOpenGLTextureCacheRelease == nil {
		return symbolCallError("CVOpenGLTextureCacheRelease", "10.4", _cVOpenGLTextureCacheReleaseErr)
	}
	_cVOpenGLTextureCacheRelease(textureCache)
	return nil
}

// CVOpenGLTextureCacheRelease releases a texture cache object.
//
// Deprecated: Deprecated since macOS 10.14.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureCacheRelease
func CVOpenGLTextureCacheRelease(textureCache CVOpenGLTextureCacheRef) {
	if callErr := tryCVOpenGLTextureCacheRelease(textureCache); callErr != nil {
		panic(callErr)
	}
}

var _cVOpenGLTextureCacheRetain func(textureCache CVOpenGLTextureCacheRef) CVOpenGLTextureCacheRef
var _cVOpenGLTextureCacheRetainErr error

func tryCVOpenGLTextureCacheRetain(textureCache CVOpenGLTextureCacheRef) (CVOpenGLTextureCacheRef, error) {
	if _cVOpenGLTextureCacheRetain == nil {
		return 0, symbolCallError("CVOpenGLTextureCacheRetain", "10.4", _cVOpenGLTextureCacheRetainErr)
	}
	return _cVOpenGLTextureCacheRetain(textureCache), nil
}

// CVOpenGLTextureCacheRetain retains a texture cache object.
//
// Deprecated: Deprecated since macOS 10.14.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureCacheRetain
func CVOpenGLTextureCacheRetain(textureCache CVOpenGLTextureCacheRef) CVOpenGLTextureCacheRef {
	result, callErr := tryCVOpenGLTextureCacheRetain(textureCache)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVOpenGLTextureRelease func(texture CVOpenGLTextureRef)
var _cVOpenGLTextureReleaseErr error

func tryCVOpenGLTextureRelease(texture CVOpenGLTextureRef) error {
	if _cVOpenGLTextureRelease == nil {
		return symbolCallError("CVOpenGLTextureRelease", "10.4", _cVOpenGLTextureReleaseErr)
	}
	_cVOpenGLTextureRelease(texture)
	return nil
}

// CVOpenGLTextureRelease releases a Core Video OpenGL texture.
//
// Deprecated: Deprecated since macOS 10.14.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureRelease
func CVOpenGLTextureRelease(texture CVOpenGLTextureRef) {
	if callErr := tryCVOpenGLTextureRelease(texture); callErr != nil {
		panic(callErr)
	}
}

var _cVOpenGLTextureRetain func(texture CVOpenGLTextureRef) CVOpenGLTextureRef
var _cVOpenGLTextureRetainErr error

func tryCVOpenGLTextureRetain(texture CVOpenGLTextureRef) (CVOpenGLTextureRef, error) {
	if _cVOpenGLTextureRetain == nil {
		return 0, symbolCallError("CVOpenGLTextureRetain", "10.4", _cVOpenGLTextureRetainErr)
	}
	return _cVOpenGLTextureRetain(texture), nil
}

// CVOpenGLTextureRetain retains a Core Video OpenGL texture.
//
// Deprecated: Deprecated since macOS 10.14.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureRetain
func CVOpenGLTextureRetain(texture CVOpenGLTextureRef) CVOpenGLTextureRef {
	result, callErr := tryCVOpenGLTextureRetain(texture)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferCopyCreationAttributes func(pixelBuffer CVPixelBufferRef) corefoundation.CFDictionaryRef
var _cVPixelBufferCopyCreationAttributesErr error

func tryCVPixelBufferCopyCreationAttributes(pixelBuffer CVPixelBufferRef) (corefoundation.CFDictionaryRef, error) {
	if _cVPixelBufferCopyCreationAttributes == nil {
		return 0, symbolCallError("CVPixelBufferCopyCreationAttributes", "12.0", _cVPixelBufferCopyCreationAttributesErr)
	}
	return _cVPixelBufferCopyCreationAttributes(pixelBuffer), nil
}

// CVPixelBufferCopyCreationAttributes.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCopyCreationAttributes(_:)
func CVPixelBufferCopyCreationAttributes(pixelBuffer CVPixelBufferRef) corefoundation.CFDictionaryRef {
	result, callErr := tryCVPixelBufferCopyCreationAttributes(pixelBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferCreate func(allocator corefoundation.CFAllocatorRef, width uintptr, height uintptr, pixelFormatType uint32, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn
var _cVPixelBufferCreateErr error

func tryCVPixelBufferCreate(allocator corefoundation.CFAllocatorRef, width uintptr, height uintptr, pixelFormatType uint32, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) (CVReturn, error) {
	if _cVPixelBufferCreate == nil {
		return *new(CVReturn), symbolCallError("CVPixelBufferCreate", "10.4", _cVPixelBufferCreateErr)
	}
	return _cVPixelBufferCreate(allocator, width, height, pixelFormatType, pixelBufferAttributes, pixelBufferOut), nil
}

// CVPixelBufferCreate creates a single pixel buffer for a given size and pixel format.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreate(_:_:_:_:_:_:)
func CVPixelBufferCreate(allocator corefoundation.CFAllocatorRef, width uintptr, height uintptr, pixelFormatType uint32, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn {
	result, callErr := tryCVPixelBufferCreate(allocator, width, height, pixelFormatType, pixelBufferAttributes, pixelBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferCreateResolvedAttributesDictionary func(allocator corefoundation.CFAllocatorRef, attributes corefoundation.CFArrayRef, resolvedDictionaryOut *corefoundation.CFDictionaryRef) CVReturn
var _cVPixelBufferCreateResolvedAttributesDictionaryErr error

func tryCVPixelBufferCreateResolvedAttributesDictionary(allocator corefoundation.CFAllocatorRef, attributes corefoundation.CFArrayRef, resolvedDictionaryOut *corefoundation.CFDictionaryRef) (CVReturn, error) {
	if _cVPixelBufferCreateResolvedAttributesDictionary == nil {
		return *new(CVReturn), symbolCallError("CVPixelBufferCreateResolvedAttributesDictionary", "10.4", _cVPixelBufferCreateResolvedAttributesDictionaryErr)
	}
	return _cVPixelBufferCreateResolvedAttributesDictionary(allocator, attributes, resolvedDictionaryOut), nil
}

// CVPixelBufferCreateResolvedAttributesDictionary resolves an array of [CFDictionary] objects describing various pixel buffer attributes into a single dictionary.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreateResolvedAttributesDictionary(_:_:_:)
func CVPixelBufferCreateResolvedAttributesDictionary(allocator corefoundation.CFAllocatorRef, attributes corefoundation.CFArrayRef, resolvedDictionaryOut *corefoundation.CFDictionaryRef) CVReturn {
	result, callErr := tryCVPixelBufferCreateResolvedAttributesDictionary(allocator, attributes, resolvedDictionaryOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferCreateWithBytes func(allocator corefoundation.CFAllocatorRef, width uintptr, height uintptr, pixelFormatType uint32, baseAddress unsafe.Pointer, bytesPerRow uintptr, releaseCallback CVPixelBufferReleaseBytesCallback, releaseRefCon unsafe.Pointer, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn
var _cVPixelBufferCreateWithBytesErr error

func tryCVPixelBufferCreateWithBytes(allocator corefoundation.CFAllocatorRef, width uintptr, height uintptr, pixelFormatType uint32, baseAddress unsafe.Pointer, bytesPerRow uintptr, releaseCallback CVPixelBufferReleaseBytesCallback, releaseRefCon unsafe.Pointer, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) (CVReturn, error) {
	if _cVPixelBufferCreateWithBytes == nil {
		return *new(CVReturn), symbolCallError("CVPixelBufferCreateWithBytes", "10.4", _cVPixelBufferCreateWithBytesErr)
	}
	return _cVPixelBufferCreateWithBytes(allocator, width, height, pixelFormatType, baseAddress, bytesPerRow, releaseCallback, releaseRefCon, pixelBufferAttributes, pixelBufferOut), nil
}

// CVPixelBufferCreateWithBytes creates a pixel buffer for a given size and pixel format containing data specified by a memory location.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreateWithBytes(_:_:_:_:_:_:_:_:_:_:)
func CVPixelBufferCreateWithBytes(allocator corefoundation.CFAllocatorRef, width uintptr, height uintptr, pixelFormatType uint32, baseAddress unsafe.Pointer, bytesPerRow uintptr, releaseCallback CVPixelBufferReleaseBytesCallback, releaseRefCon unsafe.Pointer, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn {
	result, callErr := tryCVPixelBufferCreateWithBytes(allocator, width, height, pixelFormatType, baseAddress, bytesPerRow, releaseCallback, releaseRefCon, pixelBufferAttributes, pixelBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferCreateWithIOSurface func(allocator corefoundation.CFAllocatorRef, surface unsafe.Pointer, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn
var _cVPixelBufferCreateWithIOSurfaceErr error

func tryCVPixelBufferCreateWithIOSurface(allocator corefoundation.CFAllocatorRef, surface unsafe.Pointer, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) (CVReturn, error) {
	if _cVPixelBufferCreateWithIOSurface == nil {
		return *new(CVReturn), symbolCallError("CVPixelBufferCreateWithIOSurface", "10.6", _cVPixelBufferCreateWithIOSurfaceErr)
	}
	return _cVPixelBufferCreateWithIOSurface(allocator, surface, pixelBufferAttributes, pixelBufferOut), nil
}

// CVPixelBufferCreateWithIOSurface creates a single pixel buffer for the IO surface that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreateWithIOSurface(_:_:_:_:)
func CVPixelBufferCreateWithIOSurface(allocator corefoundation.CFAllocatorRef, surface unsafe.Pointer, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn {
	result, callErr := tryCVPixelBufferCreateWithIOSurface(allocator, surface, pixelBufferAttributes, pixelBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferCreateWithPlanarBytes func(allocator corefoundation.CFAllocatorRef, width uintptr, height uintptr, pixelFormatType uint32, dataPtr unsafe.Pointer, dataSize uintptr, numberOfPlanes uintptr, planeBaseAddress unsafe.Pointer, planeWidth uintptr, planeHeight uintptr, planeBytesPerRow uintptr, releaseCallback CVPixelBufferReleasePlanarBytesCallback, releaseRefCon unsafe.Pointer, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn
var _cVPixelBufferCreateWithPlanarBytesErr error

func tryCVPixelBufferCreateWithPlanarBytes(allocator corefoundation.CFAllocatorRef, width uintptr, height uintptr, pixelFormatType uint32, dataPtr unsafe.Pointer, dataSize uintptr, numberOfPlanes uintptr, planeBaseAddress unsafe.Pointer, planeWidth uintptr, planeHeight uintptr, planeBytesPerRow uintptr, releaseCallback CVPixelBufferReleasePlanarBytesCallback, releaseRefCon unsafe.Pointer, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) (CVReturn, error) {
	if _cVPixelBufferCreateWithPlanarBytes == nil {
		return *new(CVReturn), symbolCallError("CVPixelBufferCreateWithPlanarBytes", "10.4", _cVPixelBufferCreateWithPlanarBytesErr)
	}
	return _cVPixelBufferCreateWithPlanarBytes(allocator, width, height, pixelFormatType, dataPtr, dataSize, numberOfPlanes, planeBaseAddress, planeWidth, planeHeight, planeBytesPerRow, releaseCallback, releaseRefCon, pixelBufferAttributes, pixelBufferOut), nil
}

// CVPixelBufferCreateWithPlanarBytes creates a single pixel buffer in planar format for a given size and pixel format containing data specified by a memory location.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreateWithPlanarBytes(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func CVPixelBufferCreateWithPlanarBytes(allocator corefoundation.CFAllocatorRef, width uintptr, height uintptr, pixelFormatType uint32, dataPtr unsafe.Pointer, dataSize uintptr, numberOfPlanes uintptr, planeBaseAddress unsafe.Pointer, planeWidth uintptr, planeHeight uintptr, planeBytesPerRow uintptr, releaseCallback CVPixelBufferReleasePlanarBytesCallback, releaseRefCon unsafe.Pointer, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn {
	result, callErr := tryCVPixelBufferCreateWithPlanarBytes(allocator, width, height, pixelFormatType, dataPtr, dataSize, numberOfPlanes, planeBaseAddress, planeWidth, planeHeight, planeBytesPerRow, releaseCallback, releaseRefCon, pixelBufferAttributes, pixelBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferFillExtendedPixels func(pixelBuffer CVPixelBufferRef) CVReturn
var _cVPixelBufferFillExtendedPixelsErr error

func tryCVPixelBufferFillExtendedPixels(pixelBuffer CVPixelBufferRef) (CVReturn, error) {
	if _cVPixelBufferFillExtendedPixels == nil {
		return *new(CVReturn), symbolCallError("CVPixelBufferFillExtendedPixels", "10.4", _cVPixelBufferFillExtendedPixelsErr)
	}
	return _cVPixelBufferFillExtendedPixels(pixelBuffer), nil
}

// CVPixelBufferFillExtendedPixels fills the extended pixels of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferFillExtendedPixels(_:)
func CVPixelBufferFillExtendedPixels(pixelBuffer CVPixelBufferRef) CVReturn {
	result, callErr := tryCVPixelBufferFillExtendedPixels(pixelBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferGetBaseAddress func(pixelBuffer CVPixelBufferRef) unsafe.Pointer
var _cVPixelBufferGetBaseAddressErr error

func tryCVPixelBufferGetBaseAddress(pixelBuffer CVPixelBufferRef) (unsafe.Pointer, error) {
	if _cVPixelBufferGetBaseAddress == nil {
		return nil, symbolCallError("CVPixelBufferGetBaseAddress", "10.4", _cVPixelBufferGetBaseAddressErr)
	}
	return _cVPixelBufferGetBaseAddress(pixelBuffer), nil
}

// CVPixelBufferGetBaseAddress returns the base address of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetBaseAddress(_:)
func CVPixelBufferGetBaseAddress(pixelBuffer CVPixelBufferRef) unsafe.Pointer {
	result, callErr := tryCVPixelBufferGetBaseAddress(pixelBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferGetBaseAddressOfPlane func(pixelBuffer CVPixelBufferRef, planeIndex uintptr) unsafe.Pointer
var _cVPixelBufferGetBaseAddressOfPlaneErr error

func tryCVPixelBufferGetBaseAddressOfPlane(pixelBuffer CVPixelBufferRef, planeIndex uintptr) (unsafe.Pointer, error) {
	if _cVPixelBufferGetBaseAddressOfPlane == nil {
		return nil, symbolCallError("CVPixelBufferGetBaseAddressOfPlane", "10.4", _cVPixelBufferGetBaseAddressOfPlaneErr)
	}
	return _cVPixelBufferGetBaseAddressOfPlane(pixelBuffer, planeIndex), nil
}

// CVPixelBufferGetBaseAddressOfPlane returns the base address of the plane at the specified plane index.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetBaseAddressOfPlane(_:_:)
func CVPixelBufferGetBaseAddressOfPlane(pixelBuffer CVPixelBufferRef, planeIndex uintptr) unsafe.Pointer {
	result, callErr := tryCVPixelBufferGetBaseAddressOfPlane(pixelBuffer, planeIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferGetBytesPerRow func(pixelBuffer CVPixelBufferRef) uintptr
var _cVPixelBufferGetBytesPerRowErr error

func tryCVPixelBufferGetBytesPerRow(pixelBuffer CVPixelBufferRef) (uintptr, error) {
	if _cVPixelBufferGetBytesPerRow == nil {
		return 0, symbolCallError("CVPixelBufferGetBytesPerRow", "10.4", _cVPixelBufferGetBytesPerRowErr)
	}
	return _cVPixelBufferGetBytesPerRow(pixelBuffer), nil
}

// CVPixelBufferGetBytesPerRow returns the number of bytes per row of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetBytesPerRow(_:)
func CVPixelBufferGetBytesPerRow(pixelBuffer CVPixelBufferRef) uintptr {
	result, callErr := tryCVPixelBufferGetBytesPerRow(pixelBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferGetBytesPerRowOfPlane func(pixelBuffer CVPixelBufferRef, planeIndex uintptr) uintptr
var _cVPixelBufferGetBytesPerRowOfPlaneErr error

func tryCVPixelBufferGetBytesPerRowOfPlane(pixelBuffer CVPixelBufferRef, planeIndex uintptr) (uintptr, error) {
	if _cVPixelBufferGetBytesPerRowOfPlane == nil {
		return 0, symbolCallError("CVPixelBufferGetBytesPerRowOfPlane", "10.4", _cVPixelBufferGetBytesPerRowOfPlaneErr)
	}
	return _cVPixelBufferGetBytesPerRowOfPlane(pixelBuffer, planeIndex), nil
}

// CVPixelBufferGetBytesPerRowOfPlane returns the number of bytes per row for a plane at the specified index in the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetBytesPerRowOfPlane(_:_:)
func CVPixelBufferGetBytesPerRowOfPlane(pixelBuffer CVPixelBufferRef, planeIndex uintptr) uintptr {
	result, callErr := tryCVPixelBufferGetBytesPerRowOfPlane(pixelBuffer, planeIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferGetDataSize func(pixelBuffer CVPixelBufferRef) uintptr
var _cVPixelBufferGetDataSizeErr error

func tryCVPixelBufferGetDataSize(pixelBuffer CVPixelBufferRef) (uintptr, error) {
	if _cVPixelBufferGetDataSize == nil {
		return 0, symbolCallError("CVPixelBufferGetDataSize", "10.4", _cVPixelBufferGetDataSizeErr)
	}
	return _cVPixelBufferGetDataSize(pixelBuffer), nil
}

// CVPixelBufferGetDataSize returns the data size for contiguous planes of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetDataSize(_:)
func CVPixelBufferGetDataSize(pixelBuffer CVPixelBufferRef) uintptr {
	result, callErr := tryCVPixelBufferGetDataSize(pixelBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferGetExtendedPixels func(pixelBuffer CVPixelBufferRef, extraColumnsOnLeft *uintptr, extraColumnsOnRight *uintptr, extraRowsOnTop *uintptr, extraRowsOnBottom *uintptr)
var _cVPixelBufferGetExtendedPixelsErr error

func tryCVPixelBufferGetExtendedPixels(pixelBuffer CVPixelBufferRef, extraColumnsOnLeft *uintptr, extraColumnsOnRight *uintptr, extraRowsOnTop *uintptr, extraRowsOnBottom *uintptr) error {
	if _cVPixelBufferGetExtendedPixels == nil {
		return symbolCallError("CVPixelBufferGetExtendedPixels", "10.4", _cVPixelBufferGetExtendedPixelsErr)
	}
	_cVPixelBufferGetExtendedPixels(pixelBuffer, extraColumnsOnLeft, extraColumnsOnRight, extraRowsOnTop, extraRowsOnBottom)
	return nil
}

// CVPixelBufferGetExtendedPixels returns the amount of extended pixel padding in the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetExtendedPixels(_:_:_:_:_:)
func CVPixelBufferGetExtendedPixels(pixelBuffer CVPixelBufferRef, extraColumnsOnLeft *uintptr, extraColumnsOnRight *uintptr, extraRowsOnTop *uintptr, extraRowsOnBottom *uintptr) {
	if callErr := tryCVPixelBufferGetExtendedPixels(pixelBuffer, extraColumnsOnLeft, extraColumnsOnRight, extraRowsOnTop, extraRowsOnBottom); callErr != nil {
		panic(callErr)
	}
}

var _cVPixelBufferGetHeight func(pixelBuffer CVPixelBufferRef) uintptr
var _cVPixelBufferGetHeightErr error

func tryCVPixelBufferGetHeight(pixelBuffer CVPixelBufferRef) (uintptr, error) {
	if _cVPixelBufferGetHeight == nil {
		return 0, symbolCallError("CVPixelBufferGetHeight", "10.4", _cVPixelBufferGetHeightErr)
	}
	return _cVPixelBufferGetHeight(pixelBuffer), nil
}

// CVPixelBufferGetHeight returns the height of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetHeight(_:)
func CVPixelBufferGetHeight(pixelBuffer CVPixelBufferRef) uintptr {
	result, callErr := tryCVPixelBufferGetHeight(pixelBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferGetHeightOfPlane func(pixelBuffer CVPixelBufferRef, planeIndex uintptr) uintptr
var _cVPixelBufferGetHeightOfPlaneErr error

func tryCVPixelBufferGetHeightOfPlane(pixelBuffer CVPixelBufferRef, planeIndex uintptr) (uintptr, error) {
	if _cVPixelBufferGetHeightOfPlane == nil {
		return 0, symbolCallError("CVPixelBufferGetHeightOfPlane", "10.4", _cVPixelBufferGetHeightOfPlaneErr)
	}
	return _cVPixelBufferGetHeightOfPlane(pixelBuffer, planeIndex), nil
}

// CVPixelBufferGetHeightOfPlane returns the height of the plane at planeIndex in the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetHeightOfPlane(_:_:)
func CVPixelBufferGetHeightOfPlane(pixelBuffer CVPixelBufferRef, planeIndex uintptr) uintptr {
	result, callErr := tryCVPixelBufferGetHeightOfPlane(pixelBuffer, planeIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferGetIOSurface func(pixelBuffer CVPixelBufferRef) unsafe.Pointer
var _cVPixelBufferGetIOSurfaceErr error

func tryCVPixelBufferGetIOSurface(pixelBuffer CVPixelBufferRef) (unsafe.Pointer, error) {
	if _cVPixelBufferGetIOSurface == nil {
		return nil, symbolCallError("CVPixelBufferGetIOSurface", "10.6", _cVPixelBufferGetIOSurfaceErr)
	}
	return _cVPixelBufferGetIOSurface(pixelBuffer), nil
}

// CVPixelBufferGetIOSurface returns the IOSurface backing the pixel buffer, or [NULL] if it is not backed by an IOSurface.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetIOSurface(_:)
func CVPixelBufferGetIOSurface(pixelBuffer CVPixelBufferRef) unsafe.Pointer {
	result, callErr := tryCVPixelBufferGetIOSurface(pixelBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferGetPixelFormatType func(pixelBuffer CVPixelBufferRef) uint32
var _cVPixelBufferGetPixelFormatTypeErr error

func tryCVPixelBufferGetPixelFormatType(pixelBuffer CVPixelBufferRef) (uint32, error) {
	if _cVPixelBufferGetPixelFormatType == nil {
		return 0, symbolCallError("CVPixelBufferGetPixelFormatType", "10.4", _cVPixelBufferGetPixelFormatTypeErr)
	}
	return _cVPixelBufferGetPixelFormatType(pixelBuffer), nil
}

// CVPixelBufferGetPixelFormatType returns the pixel format type of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetPixelFormatType(_:)
func CVPixelBufferGetPixelFormatType(pixelBuffer CVPixelBufferRef) uint32 {
	result, callErr := tryCVPixelBufferGetPixelFormatType(pixelBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferGetPlaneCount func(pixelBuffer CVPixelBufferRef) uintptr
var _cVPixelBufferGetPlaneCountErr error

func tryCVPixelBufferGetPlaneCount(pixelBuffer CVPixelBufferRef) (uintptr, error) {
	if _cVPixelBufferGetPlaneCount == nil {
		return 0, symbolCallError("CVPixelBufferGetPlaneCount", "10.4", _cVPixelBufferGetPlaneCountErr)
	}
	return _cVPixelBufferGetPlaneCount(pixelBuffer), nil
}

// CVPixelBufferGetPlaneCount returns number of planes of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetPlaneCount(_:)
func CVPixelBufferGetPlaneCount(pixelBuffer CVPixelBufferRef) uintptr {
	result, callErr := tryCVPixelBufferGetPlaneCount(pixelBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferGetTypeID func() uint
var _cVPixelBufferGetTypeIDErr error

func tryCVPixelBufferGetTypeID() (uint, error) {
	if _cVPixelBufferGetTypeID == nil {
		return 0, symbolCallError("CVPixelBufferGetTypeID", "10.4", _cVPixelBufferGetTypeIDErr)
	}
	return _cVPixelBufferGetTypeID(), nil
}

// CVPixelBufferGetTypeID returns the Core Foundation type identifier of the pixel buffer type.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetTypeID()
func CVPixelBufferGetTypeID() uint {
	result, callErr := tryCVPixelBufferGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferGetWidth func(pixelBuffer CVPixelBufferRef) uintptr
var _cVPixelBufferGetWidthErr error

func tryCVPixelBufferGetWidth(pixelBuffer CVPixelBufferRef) (uintptr, error) {
	if _cVPixelBufferGetWidth == nil {
		return 0, symbolCallError("CVPixelBufferGetWidth", "10.4", _cVPixelBufferGetWidthErr)
	}
	return _cVPixelBufferGetWidth(pixelBuffer), nil
}

// CVPixelBufferGetWidth returns the width of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetWidth(_:)
func CVPixelBufferGetWidth(pixelBuffer CVPixelBufferRef) uintptr {
	result, callErr := tryCVPixelBufferGetWidth(pixelBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferGetWidthOfPlane func(pixelBuffer CVPixelBufferRef, planeIndex uintptr) uintptr
var _cVPixelBufferGetWidthOfPlaneErr error

func tryCVPixelBufferGetWidthOfPlane(pixelBuffer CVPixelBufferRef, planeIndex uintptr) (uintptr, error) {
	if _cVPixelBufferGetWidthOfPlane == nil {
		return 0, symbolCallError("CVPixelBufferGetWidthOfPlane", "10.4", _cVPixelBufferGetWidthOfPlaneErr)
	}
	return _cVPixelBufferGetWidthOfPlane(pixelBuffer, planeIndex), nil
}

// CVPixelBufferGetWidthOfPlane returns the width of the plane at a given index in the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetWidthOfPlane(_:_:)
func CVPixelBufferGetWidthOfPlane(pixelBuffer CVPixelBufferRef, planeIndex uintptr) uintptr {
	result, callErr := tryCVPixelBufferGetWidthOfPlane(pixelBuffer, planeIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferIsCompatibleWithAttributes func(pixelBuffer CVPixelBufferRef, attributes corefoundation.CFDictionaryRef) bool
var _cVPixelBufferIsCompatibleWithAttributesErr error

func tryCVPixelBufferIsCompatibleWithAttributes(pixelBuffer CVPixelBufferRef, attributes corefoundation.CFDictionaryRef) (bool, error) {
	if _cVPixelBufferIsCompatibleWithAttributes == nil {
		return false, symbolCallError("CVPixelBufferIsCompatibleWithAttributes", "10.4", _cVPixelBufferIsCompatibleWithAttributesErr)
	}
	return _cVPixelBufferIsCompatibleWithAttributes(pixelBuffer, attributes), nil
}

// CVPixelBufferIsCompatibleWithAttributes.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferIsCompatibleWithAttributes(_:_:)
func CVPixelBufferIsCompatibleWithAttributes(pixelBuffer CVPixelBufferRef, attributes corefoundation.CFDictionaryRef) bool {
	result, callErr := tryCVPixelBufferIsCompatibleWithAttributes(pixelBuffer, attributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferIsPlanar func(pixelBuffer CVPixelBufferRef) bool
var _cVPixelBufferIsPlanarErr error

func tryCVPixelBufferIsPlanar(pixelBuffer CVPixelBufferRef) (bool, error) {
	if _cVPixelBufferIsPlanar == nil {
		return false, symbolCallError("CVPixelBufferIsPlanar", "10.4", _cVPixelBufferIsPlanarErr)
	}
	return _cVPixelBufferIsPlanar(pixelBuffer), nil
}

// CVPixelBufferIsPlanar determines whether the pixel buffer is planar.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferIsPlanar(_:)
func CVPixelBufferIsPlanar(pixelBuffer CVPixelBufferRef) bool {
	result, callErr := tryCVPixelBufferIsPlanar(pixelBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferLockBaseAddress func(pixelBuffer CVPixelBufferRef, lockFlags CVPixelBufferLockFlags) CVReturn
var _cVPixelBufferLockBaseAddressErr error

func tryCVPixelBufferLockBaseAddress(pixelBuffer CVPixelBufferRef, lockFlags CVPixelBufferLockFlags) (CVReturn, error) {
	if _cVPixelBufferLockBaseAddress == nil {
		return *new(CVReturn), symbolCallError("CVPixelBufferLockBaseAddress", "10.4", _cVPixelBufferLockBaseAddressErr)
	}
	return _cVPixelBufferLockBaseAddress(pixelBuffer, lockFlags), nil
}

// CVPixelBufferLockBaseAddress locks the base address of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferLockBaseAddress(_:_:)
func CVPixelBufferLockBaseAddress(pixelBuffer CVPixelBufferRef, lockFlags CVPixelBufferLockFlags) CVReturn {
	result, callErr := tryCVPixelBufferLockBaseAddress(pixelBuffer, lockFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferPoolCreate func(allocator corefoundation.CFAllocatorRef, poolAttributes corefoundation.CFDictionaryRef, pixelBufferAttributes corefoundation.CFDictionaryRef, poolOut *CVPixelBufferPoolRef) CVReturn
var _cVPixelBufferPoolCreateErr error

func tryCVPixelBufferPoolCreate(allocator corefoundation.CFAllocatorRef, poolAttributes corefoundation.CFDictionaryRef, pixelBufferAttributes corefoundation.CFDictionaryRef, poolOut *CVPixelBufferPoolRef) (CVReturn, error) {
	if _cVPixelBufferPoolCreate == nil {
		return *new(CVReturn), symbolCallError("CVPixelBufferPoolCreate", "10.4", _cVPixelBufferPoolCreateErr)
	}
	return _cVPixelBufferPoolCreate(allocator, poolAttributes, pixelBufferAttributes, poolOut), nil
}

// CVPixelBufferPoolCreate creates a pixel buffer pool using the allocator and attributes that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolCreate(_:_:_:_:)
func CVPixelBufferPoolCreate(allocator corefoundation.CFAllocatorRef, poolAttributes corefoundation.CFDictionaryRef, pixelBufferAttributes corefoundation.CFDictionaryRef, poolOut *CVPixelBufferPoolRef) CVReturn {
	result, callErr := tryCVPixelBufferPoolCreate(allocator, poolAttributes, pixelBufferAttributes, poolOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferPoolCreatePixelBuffer func(allocator corefoundation.CFAllocatorRef, pixelBufferPool CVPixelBufferPoolRef, pixelBufferOut *CVPixelBufferRef) CVReturn
var _cVPixelBufferPoolCreatePixelBufferErr error

func tryCVPixelBufferPoolCreatePixelBuffer(allocator corefoundation.CFAllocatorRef, pixelBufferPool CVPixelBufferPoolRef, pixelBufferOut *CVPixelBufferRef) (CVReturn, error) {
	if _cVPixelBufferPoolCreatePixelBuffer == nil {
		return *new(CVReturn), symbolCallError("CVPixelBufferPoolCreatePixelBuffer", "10.4", _cVPixelBufferPoolCreatePixelBufferErr)
	}
	return _cVPixelBufferPoolCreatePixelBuffer(allocator, pixelBufferPool, pixelBufferOut), nil
}

// CVPixelBufferPoolCreatePixelBuffer creates a pixel buffer from a pixel buffer pool, using the allocator that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolCreatePixelBuffer(_:_:_:)
func CVPixelBufferPoolCreatePixelBuffer(allocator corefoundation.CFAllocatorRef, pixelBufferPool CVPixelBufferPoolRef, pixelBufferOut *CVPixelBufferRef) CVReturn {
	result, callErr := tryCVPixelBufferPoolCreatePixelBuffer(allocator, pixelBufferPool, pixelBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferPoolCreatePixelBufferWithAuxAttributes func(allocator corefoundation.CFAllocatorRef, pixelBufferPool CVPixelBufferPoolRef, auxAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn
var _cVPixelBufferPoolCreatePixelBufferWithAuxAttributesErr error

func tryCVPixelBufferPoolCreatePixelBufferWithAuxAttributes(allocator corefoundation.CFAllocatorRef, pixelBufferPool CVPixelBufferPoolRef, auxAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) (CVReturn, error) {
	if _cVPixelBufferPoolCreatePixelBufferWithAuxAttributes == nil {
		return *new(CVReturn), symbolCallError("CVPixelBufferPoolCreatePixelBufferWithAuxAttributes", "10.7", _cVPixelBufferPoolCreatePixelBufferWithAuxAttributesErr)
	}
	return _cVPixelBufferPoolCreatePixelBufferWithAuxAttributes(allocator, pixelBufferPool, auxAttributes, pixelBufferOut), nil
}

// CVPixelBufferPoolCreatePixelBufferWithAuxAttributes creates a new pixel buffer with auxiliary attributes from the pool.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolCreatePixelBufferWithAuxAttributes(_:_:_:_:)
func CVPixelBufferPoolCreatePixelBufferWithAuxAttributes(allocator corefoundation.CFAllocatorRef, pixelBufferPool CVPixelBufferPoolRef, auxAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn {
	result, callErr := tryCVPixelBufferPoolCreatePixelBufferWithAuxAttributes(allocator, pixelBufferPool, auxAttributes, pixelBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferPoolFlush func(pool CVPixelBufferPoolRef, options CVPixelBufferPoolFlushFlags)
var _cVPixelBufferPoolFlushErr error

func tryCVPixelBufferPoolFlush(pool CVPixelBufferPoolRef, options CVPixelBufferPoolFlushFlags) error {
	if _cVPixelBufferPoolFlush == nil {
		return symbolCallError("CVPixelBufferPoolFlush", "10.4", _cVPixelBufferPoolFlushErr)
	}
	_cVPixelBufferPoolFlush(pool, options)
	return nil
}

// CVPixelBufferPoolFlush frees pixel buffers from the pool based on the options that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolFlush(_:_:)
func CVPixelBufferPoolFlush(pool CVPixelBufferPoolRef, options CVPixelBufferPoolFlushFlags) {
	if callErr := tryCVPixelBufferPoolFlush(pool, options); callErr != nil {
		panic(callErr)
	}
}

var _cVPixelBufferPoolGetAttributes func(pool CVPixelBufferPoolRef) corefoundation.CFDictionaryRef
var _cVPixelBufferPoolGetAttributesErr error

func tryCVPixelBufferPoolGetAttributes(pool CVPixelBufferPoolRef) (corefoundation.CFDictionaryRef, error) {
	if _cVPixelBufferPoolGetAttributes == nil {
		return 0, symbolCallError("CVPixelBufferPoolGetAttributes", "10.4", _cVPixelBufferPoolGetAttributesErr)
	}
	return _cVPixelBufferPoolGetAttributes(pool), nil
}

// CVPixelBufferPoolGetAttributes the pool attributes dictionary for a pixel buffer pool.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolGetAttributes(_:)
func CVPixelBufferPoolGetAttributes(pool CVPixelBufferPoolRef) corefoundation.CFDictionaryRef {
	result, callErr := tryCVPixelBufferPoolGetAttributes(pool)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferPoolGetPixelBufferAttributes func(pool CVPixelBufferPoolRef) corefoundation.CFDictionaryRef
var _cVPixelBufferPoolGetPixelBufferAttributesErr error

func tryCVPixelBufferPoolGetPixelBufferAttributes(pool CVPixelBufferPoolRef) (corefoundation.CFDictionaryRef, error) {
	if _cVPixelBufferPoolGetPixelBufferAttributes == nil {
		return 0, symbolCallError("CVPixelBufferPoolGetPixelBufferAttributes", "10.4", _cVPixelBufferPoolGetPixelBufferAttributesErr)
	}
	return _cVPixelBufferPoolGetPixelBufferAttributes(pool), nil
}

// CVPixelBufferPoolGetPixelBufferAttributes the attributes of pixel buffers which the system creates using the pool you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolGetPixelBufferAttributes(_:)
func CVPixelBufferPoolGetPixelBufferAttributes(pool CVPixelBufferPoolRef) corefoundation.CFDictionaryRef {
	result, callErr := tryCVPixelBufferPoolGetPixelBufferAttributes(pool)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferPoolGetTypeID func() uint
var _cVPixelBufferPoolGetTypeIDErr error

func tryCVPixelBufferPoolGetTypeID() (uint, error) {
	if _cVPixelBufferPoolGetTypeID == nil {
		return 0, symbolCallError("CVPixelBufferPoolGetTypeID", "10.4", _cVPixelBufferPoolGetTypeIDErr)
	}
	return _cVPixelBufferPoolGetTypeID(), nil
}

// CVPixelBufferPoolGetTypeID returns the Core Foundation type identifier of the pixel buffer pool type.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolGetTypeID()
func CVPixelBufferPoolGetTypeID() uint {
	result, callErr := tryCVPixelBufferPoolGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferPoolRelease func(pixelBufferPool CVPixelBufferPoolRef)
var _cVPixelBufferPoolReleaseErr error

func tryCVPixelBufferPoolRelease(pixelBufferPool CVPixelBufferPoolRef) error {
	if _cVPixelBufferPoolRelease == nil {
		return symbolCallError("CVPixelBufferPoolRelease", "10.4", _cVPixelBufferPoolReleaseErr)
	}
	_cVPixelBufferPoolRelease(pixelBufferPool)
	return nil
}

// CVPixelBufferPoolRelease releases a pixel buffer pool.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolRelease
func CVPixelBufferPoolRelease(pixelBufferPool CVPixelBufferPoolRef) {
	if callErr := tryCVPixelBufferPoolRelease(pixelBufferPool); callErr != nil {
		panic(callErr)
	}
}

var _cVPixelBufferPoolRetain func(pixelBufferPool CVPixelBufferPoolRef) CVPixelBufferPoolRef
var _cVPixelBufferPoolRetainErr error

func tryCVPixelBufferPoolRetain(pixelBufferPool CVPixelBufferPoolRef) (CVPixelBufferPoolRef, error) {
	if _cVPixelBufferPoolRetain == nil {
		return 0, symbolCallError("CVPixelBufferPoolRetain", "10.4", _cVPixelBufferPoolRetainErr)
	}
	return _cVPixelBufferPoolRetain(pixelBufferPool), nil
}

// CVPixelBufferPoolRetain retains the pixel buffer pool that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolRetain
func CVPixelBufferPoolRetain(pixelBufferPool CVPixelBufferPoolRef) CVPixelBufferPoolRef {
	result, callErr := tryCVPixelBufferPoolRetain(pixelBufferPool)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferRelease func(texture CVPixelBufferRef)
var _cVPixelBufferReleaseErr error

func tryCVPixelBufferRelease(texture CVPixelBufferRef) error {
	if _cVPixelBufferRelease == nil {
		return symbolCallError("CVPixelBufferRelease", "10.4", _cVPixelBufferReleaseErr)
	}
	_cVPixelBufferRelease(texture)
	return nil
}

// CVPixelBufferRelease releases a pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferRelease
func CVPixelBufferRelease(texture CVPixelBufferRef) {
	if callErr := tryCVPixelBufferRelease(texture); callErr != nil {
		panic(callErr)
	}
}

var _cVPixelBufferRetain func(texture CVPixelBufferRef) CVPixelBufferRef
var _cVPixelBufferRetainErr error

func tryCVPixelBufferRetain(texture CVPixelBufferRef) (CVPixelBufferRef, error) {
	if _cVPixelBufferRetain == nil {
		return 0, symbolCallError("CVPixelBufferRetain", "10.4", _cVPixelBufferRetainErr)
	}
	return _cVPixelBufferRetain(texture), nil
}

// CVPixelBufferRetain retains a pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferRetain
func CVPixelBufferRetain(texture CVPixelBufferRef) CVPixelBufferRef {
	result, callErr := tryCVPixelBufferRetain(texture)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelBufferUnlockBaseAddress func(pixelBuffer CVPixelBufferRef, unlockFlags CVPixelBufferLockFlags) CVReturn
var _cVPixelBufferUnlockBaseAddressErr error

func tryCVPixelBufferUnlockBaseAddress(pixelBuffer CVPixelBufferRef, unlockFlags CVPixelBufferLockFlags) (CVReturn, error) {
	if _cVPixelBufferUnlockBaseAddress == nil {
		return *new(CVReturn), symbolCallError("CVPixelBufferUnlockBaseAddress", "10.4", _cVPixelBufferUnlockBaseAddressErr)
	}
	return _cVPixelBufferUnlockBaseAddress(pixelBuffer, unlockFlags), nil
}

// CVPixelBufferUnlockBaseAddress unlocks the base address of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferUnlockBaseAddress(_:_:)
func CVPixelBufferUnlockBaseAddress(pixelBuffer CVPixelBufferRef, unlockFlags CVPixelBufferLockFlags) CVReturn {
	result, callErr := tryCVPixelBufferUnlockBaseAddress(pixelBuffer, unlockFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes func(allocator corefoundation.CFAllocatorRef) corefoundation.CFArrayRef
var _cVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypesErr error

func tryCVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes(allocator corefoundation.CFAllocatorRef) (corefoundation.CFArrayRef, error) {
	if _cVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes == nil {
		return 0, symbolCallError("CVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes", "10.4", _cVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypesErr)
	}
	return _cVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes(allocator), nil
}

// CVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes returns all the pixel format descriptions known to Core Video.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes(_:)
func CVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes(allocator corefoundation.CFAllocatorRef) corefoundation.CFArrayRef {
	result, callErr := tryCVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes(allocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelFormatDescriptionCreateWithPixelFormatType func(allocator corefoundation.CFAllocatorRef, pixelFormat uint32) corefoundation.CFDictionaryRef
var _cVPixelFormatDescriptionCreateWithPixelFormatTypeErr error

func tryCVPixelFormatDescriptionCreateWithPixelFormatType(allocator corefoundation.CFAllocatorRef, pixelFormat uint32) (corefoundation.CFDictionaryRef, error) {
	if _cVPixelFormatDescriptionCreateWithPixelFormatType == nil {
		return 0, symbolCallError("CVPixelFormatDescriptionCreateWithPixelFormatType", "10.4", _cVPixelFormatDescriptionCreateWithPixelFormatTypeErr)
	}
	return _cVPixelFormatDescriptionCreateWithPixelFormatType(allocator, pixelFormat), nil
}

// CVPixelFormatDescriptionCreateWithPixelFormatType creates a pixel format description from a given [OSType] identifier.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelFormatDescriptionCreateWithPixelFormatType(_:_:)
func CVPixelFormatDescriptionCreateWithPixelFormatType(allocator corefoundation.CFAllocatorRef, pixelFormat uint32) corefoundation.CFDictionaryRef {
	result, callErr := tryCVPixelFormatDescriptionCreateWithPixelFormatType(allocator, pixelFormat)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType func(description corefoundation.CFDictionaryRef, pixelFormat uint32)
var _cVPixelFormatDescriptionRegisterDescriptionWithPixelFormatTypeErr error

func tryCVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType(description corefoundation.CFDictionaryRef, pixelFormat uint32) error {
	if _cVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType == nil {
		return symbolCallError("CVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType", "10.4", _cVPixelFormatDescriptionRegisterDescriptionWithPixelFormatTypeErr)
	}
	_cVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType(description, pixelFormat)
	return nil
}

// CVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType registers a pixel format description with Core Video.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType(_:_:)
func CVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType(description corefoundation.CFDictionaryRef, pixelFormat uint32) {
	if callErr := tryCVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType(description, pixelFormat); callErr != nil {
		panic(callErr)
	}
}

var _cVPixelFormatTypeCopyFourCharCodeString func(pixelFormat uint32) corefoundation.CFStringRef
var _cVPixelFormatTypeCopyFourCharCodeStringErr error

func tryCVPixelFormatTypeCopyFourCharCodeString(pixelFormat uint32) (corefoundation.CFStringRef, error) {
	if _cVPixelFormatTypeCopyFourCharCodeString == nil {
		return 0, symbolCallError("CVPixelFormatTypeCopyFourCharCodeString", "26.0", _cVPixelFormatTypeCopyFourCharCodeStringErr)
	}
	return _cVPixelFormatTypeCopyFourCharCodeString(pixelFormat), nil
}

// CVPixelFormatTypeCopyFourCharCodeString.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelFormatTypeCopyFourCharCodeString(_:)
func CVPixelFormatTypeCopyFourCharCodeString(pixelFormat uint32) corefoundation.CFStringRef {
	result, callErr := tryCVPixelFormatTypeCopyFourCharCodeString(pixelFormat)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVTransferFunctionGetIntegerCodePointForString func(transferFunctionString corefoundation.CFStringRef) int
var _cVTransferFunctionGetIntegerCodePointForStringErr error

func tryCVTransferFunctionGetIntegerCodePointForString(transferFunctionString corefoundation.CFStringRef) (int, error) {
	if _cVTransferFunctionGetIntegerCodePointForString == nil {
		return 0, symbolCallError("CVTransferFunctionGetIntegerCodePointForString", "10.13", _cVTransferFunctionGetIntegerCodePointForStringErr)
	}
	return _cVTransferFunctionGetIntegerCodePointForString(transferFunctionString), nil
}

// CVTransferFunctionGetIntegerCodePointForString returns the standard integer code point corresponding to the Core Video transfer function string that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVTransferFunctionGetIntegerCodePointForString(_:)
func CVTransferFunctionGetIntegerCodePointForString(transferFunctionString corefoundation.CFStringRef) int {
	result, callErr := tryCVTransferFunctionGetIntegerCodePointForString(transferFunctionString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVTransferFunctionGetStringForIntegerCodePoint func(transferFunctionCodePoint int) corefoundation.CFStringRef
var _cVTransferFunctionGetStringForIntegerCodePointErr error

func tryCVTransferFunctionGetStringForIntegerCodePoint(transferFunctionCodePoint int) (corefoundation.CFStringRef, error) {
	if _cVTransferFunctionGetStringForIntegerCodePoint == nil {
		return 0, symbolCallError("CVTransferFunctionGetStringForIntegerCodePoint", "10.13", _cVTransferFunctionGetStringForIntegerCodePointErr)
	}
	return _cVTransferFunctionGetStringForIntegerCodePoint(transferFunctionCodePoint), nil
}

// CVTransferFunctionGetStringForIntegerCodePoint returns the Core Video transfer function string corresponding to the standard integer code point that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVTransferFunctionGetStringForIntegerCodePoint(_:)
func CVTransferFunctionGetStringForIntegerCodePoint(transferFunctionCodePoint int) corefoundation.CFStringRef {
	result, callErr := tryCVTransferFunctionGetStringForIntegerCodePoint(transferFunctionCodePoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVYCbCrMatrixGetIntegerCodePointForString func(yCbCrMatrixString corefoundation.CFStringRef) int
var _cVYCbCrMatrixGetIntegerCodePointForStringErr error

func tryCVYCbCrMatrixGetIntegerCodePointForString(yCbCrMatrixString corefoundation.CFStringRef) (int, error) {
	if _cVYCbCrMatrixGetIntegerCodePointForString == nil {
		return 0, symbolCallError("CVYCbCrMatrixGetIntegerCodePointForString", "10.13", _cVYCbCrMatrixGetIntegerCodePointForStringErr)
	}
	return _cVYCbCrMatrixGetIntegerCodePointForString(yCbCrMatrixString), nil
}

// CVYCbCrMatrixGetIntegerCodePointForString returns the standard integer code point corresponding to the Core Video YCbCr matrix string that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVYCbCrMatrixGetIntegerCodePointForString(_:)
func CVYCbCrMatrixGetIntegerCodePointForString(yCbCrMatrixString corefoundation.CFStringRef) int {
	result, callErr := tryCVYCbCrMatrixGetIntegerCodePointForString(yCbCrMatrixString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cVYCbCrMatrixGetStringForIntegerCodePoint func(yCbCrMatrixCodePoint int) corefoundation.CFStringRef
var _cVYCbCrMatrixGetStringForIntegerCodePointErr error

func tryCVYCbCrMatrixGetStringForIntegerCodePoint(yCbCrMatrixCodePoint int) (corefoundation.CFStringRef, error) {
	if _cVYCbCrMatrixGetStringForIntegerCodePoint == nil {
		return 0, symbolCallError("CVYCbCrMatrixGetStringForIntegerCodePoint", "10.13", _cVYCbCrMatrixGetStringForIntegerCodePointErr)
	}
	return _cVYCbCrMatrixGetStringForIntegerCodePoint(yCbCrMatrixCodePoint), nil
}

// CVYCbCrMatrixGetStringForIntegerCodePoint returns the Core Video YCbCr matrix string corresponding to the standard integer code point that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVYCbCrMatrixGetStringForIntegerCodePoint(_:)
func CVYCbCrMatrixGetStringForIntegerCodePoint(yCbCrMatrixCodePoint int) corefoundation.CFStringRef {
	result, callErr := tryCVYCbCrMatrixGetStringForIntegerCodePoint(yCbCrMatrixCodePoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_cVBufferCopyAttachment, &_cVBufferCopyAttachmentErr, frameworkHandle, "CVBufferCopyAttachment", "12.0")
	registerFunc(&_cVBufferCopyAttachments, &_cVBufferCopyAttachmentsErr, frameworkHandle, "CVBufferCopyAttachments", "12.0")
	registerFunc(&_cVBufferHasAttachment, &_cVBufferHasAttachmentErr, frameworkHandle, "CVBufferHasAttachment", "12.0")
	registerFunc(&_cVBufferPropagateAttachments, &_cVBufferPropagateAttachmentsErr, frameworkHandle, "CVBufferPropagateAttachments", "10.4")
	registerFunc(&_cVBufferRelease, &_cVBufferReleaseErr, frameworkHandle, "CVBufferRelease", "10.4")
	registerFunc(&_cVBufferRemoveAllAttachments, &_cVBufferRemoveAllAttachmentsErr, frameworkHandle, "CVBufferRemoveAllAttachments", "10.4")
	registerFunc(&_cVBufferRemoveAttachment, &_cVBufferRemoveAttachmentErr, frameworkHandle, "CVBufferRemoveAttachment", "10.4")
	registerFunc(&_cVBufferRetain, &_cVBufferRetainErr, frameworkHandle, "CVBufferRetain", "10.4")
	registerFunc(&_cVBufferSetAttachment, &_cVBufferSetAttachmentErr, frameworkHandle, "CVBufferSetAttachment", "10.4")
	registerFunc(&_cVBufferSetAttachments, &_cVBufferSetAttachmentsErr, frameworkHandle, "CVBufferSetAttachments", "10.4")
	registerFunc(&_cVColorPrimariesGetIntegerCodePointForString, &_cVColorPrimariesGetIntegerCodePointForStringErr, frameworkHandle, "CVColorPrimariesGetIntegerCodePointForString", "10.13")
	registerFunc(&_cVColorPrimariesGetStringForIntegerCodePoint, &_cVColorPrimariesGetStringForIntegerCodePointErr, frameworkHandle, "CVColorPrimariesGetStringForIntegerCodePoint", "10.13")
	registerFunc(&_cVDisplayLinkCreateWithActiveCGDisplays, &_cVDisplayLinkCreateWithActiveCGDisplaysErr, frameworkHandle, "CVDisplayLinkCreateWithActiveCGDisplays", "10.4")
	registerFunc(&_cVDisplayLinkCreateWithCGDisplay, &_cVDisplayLinkCreateWithCGDisplayErr, frameworkHandle, "CVDisplayLinkCreateWithCGDisplay", "10.4")
	registerFunc(&_cVDisplayLinkCreateWithCGDisplays, &_cVDisplayLinkCreateWithCGDisplaysErr, frameworkHandle, "CVDisplayLinkCreateWithCGDisplays", "10.4")
	registerFunc(&_cVDisplayLinkCreateWithOpenGLDisplayMask, &_cVDisplayLinkCreateWithOpenGLDisplayMaskErr, frameworkHandle, "CVDisplayLinkCreateWithOpenGLDisplayMask", "10.4")
	registerFunc(&_cVDisplayLinkGetActualOutputVideoRefreshPeriod, &_cVDisplayLinkGetActualOutputVideoRefreshPeriodErr, frameworkHandle, "CVDisplayLinkGetActualOutputVideoRefreshPeriod", "10.4")
	registerFunc(&_cVDisplayLinkGetCurrentCGDisplay, &_cVDisplayLinkGetCurrentCGDisplayErr, frameworkHandle, "CVDisplayLinkGetCurrentCGDisplay", "10.4")
	registerFunc(&_cVDisplayLinkGetCurrentTime, &_cVDisplayLinkGetCurrentTimeErr, frameworkHandle, "CVDisplayLinkGetCurrentTime", "10.4")
	registerFunc(&_cVDisplayLinkGetNominalOutputVideoRefreshPeriod, &_cVDisplayLinkGetNominalOutputVideoRefreshPeriodErr, frameworkHandle, "CVDisplayLinkGetNominalOutputVideoRefreshPeriod", "10.4")
	registerFunc(&_cVDisplayLinkGetOutputVideoLatency, &_cVDisplayLinkGetOutputVideoLatencyErr, frameworkHandle, "CVDisplayLinkGetOutputVideoLatency", "10.4")
	registerFunc(&_cVDisplayLinkGetTypeID, &_cVDisplayLinkGetTypeIDErr, frameworkHandle, "CVDisplayLinkGetTypeID", "10.4")
	registerFunc(&_cVDisplayLinkIsRunning, &_cVDisplayLinkIsRunningErr, frameworkHandle, "CVDisplayLinkIsRunning", "10.4")
	registerFunc(&_cVDisplayLinkRelease, &_cVDisplayLinkReleaseErr, frameworkHandle, "CVDisplayLinkRelease", "10.4")
	registerFunc(&_cVDisplayLinkRetain, &_cVDisplayLinkRetainErr, frameworkHandle, "CVDisplayLinkRetain", "10.4")
	registerFunc(&_cVDisplayLinkSetCurrentCGDisplay, &_cVDisplayLinkSetCurrentCGDisplayErr, frameworkHandle, "CVDisplayLinkSetCurrentCGDisplay", "10.4")
	registerFunc(&_cVDisplayLinkSetCurrentCGDisplayFromOpenGLContext, &_cVDisplayLinkSetCurrentCGDisplayFromOpenGLContextErr, frameworkHandle, "CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext", "10.4")
	registerFunc(&_cVDisplayLinkSetOutputCallback, &_cVDisplayLinkSetOutputCallbackErr, frameworkHandle, "CVDisplayLinkSetOutputCallback", "10.4")
	registerFunc(&_cVDisplayLinkSetOutputHandler, &_cVDisplayLinkSetOutputHandlerErr, frameworkHandle, "CVDisplayLinkSetOutputHandler", "10.4")
	registerFunc(&_cVDisplayLinkStart, &_cVDisplayLinkStartErr, frameworkHandle, "CVDisplayLinkStart", "10.4")
	registerFunc(&_cVDisplayLinkStop, &_cVDisplayLinkStopErr, frameworkHandle, "CVDisplayLinkStop", "10.4")
	registerFunc(&_cVDisplayLinkTranslateTime, &_cVDisplayLinkTranslateTimeErr, frameworkHandle, "CVDisplayLinkTranslateTime", "10.4")
	registerFunc(&_cVGetCurrentHostTime, &_cVGetCurrentHostTimeErr, frameworkHandle, "CVGetCurrentHostTime", "10.4")
	registerFunc(&_cVGetHostClockFrequency, &_cVGetHostClockFrequencyErr, frameworkHandle, "CVGetHostClockFrequency", "10.4")
	registerFunc(&_cVGetHostClockMinimumTimeDelta, &_cVGetHostClockMinimumTimeDeltaErr, frameworkHandle, "CVGetHostClockMinimumTimeDelta", "10.4")
	registerFunc(&_cVImageBufferCreateColorSpaceFromAttachments, &_cVImageBufferCreateColorSpaceFromAttachmentsErr, frameworkHandle, "CVImageBufferCreateColorSpaceFromAttachments", "10.8")
	registerFunc(&_cVImageBufferGetCleanRect, &_cVImageBufferGetCleanRectErr, frameworkHandle, "CVImageBufferGetCleanRect", "10.4")
	registerFunc(&_cVImageBufferGetColorSpace, &_cVImageBufferGetColorSpaceErr, frameworkHandle, "CVImageBufferGetColorSpace", "10.4")
	registerFunc(&_cVImageBufferGetDisplaySize, &_cVImageBufferGetDisplaySizeErr, frameworkHandle, "CVImageBufferGetDisplaySize", "10.4")
	registerFunc(&_cVImageBufferGetEncodedSize, &_cVImageBufferGetEncodedSizeErr, frameworkHandle, "CVImageBufferGetEncodedSize", "10.4")
	registerFunc(&_cVImageBufferIsFlipped, &_cVImageBufferIsFlippedErr, frameworkHandle, "CVImageBufferIsFlipped", "10.4")
	registerFunc(&_cVIsCompressedPixelFormatAvailable, &_cVIsCompressedPixelFormatAvailableErr, frameworkHandle, "CVIsCompressedPixelFormatAvailable", "12.0")
	registerFunc(&_cVMetalBufferCacheCreate, &_cVMetalBufferCacheCreateErr, frameworkHandle, "CVMetalBufferCacheCreate", "15.0")
	registerFunc(&_cVMetalBufferCacheCreateBufferFromImage, &_cVMetalBufferCacheCreateBufferFromImageErr, frameworkHandle, "CVMetalBufferCacheCreateBufferFromImage", "15.0")
	registerFunc(&_cVMetalBufferCacheFlush, &_cVMetalBufferCacheFlushErr, frameworkHandle, "CVMetalBufferCacheFlush", "15.0")
	registerFunc(&_cVMetalBufferCacheGetTypeID, &_cVMetalBufferCacheGetTypeIDErr, frameworkHandle, "CVMetalBufferCacheGetTypeID", "15.0")
	registerFunc(&_cVMetalBufferGetBuffer, &_cVMetalBufferGetBufferErr, frameworkHandle, "CVMetalBufferGetBuffer", "15.0")
	registerFunc(&_cVMetalBufferGetTypeID, &_cVMetalBufferGetTypeIDErr, frameworkHandle, "CVMetalBufferGetTypeID", "15.0")
	registerFunc(&_cVMetalTextureCacheCreate, &_cVMetalTextureCacheCreateErr, frameworkHandle, "CVMetalTextureCacheCreate", "10.11")
	registerFunc(&_cVMetalTextureCacheCreateTextureFromImage, &_cVMetalTextureCacheCreateTextureFromImageErr, frameworkHandle, "CVMetalTextureCacheCreateTextureFromImage", "10.11")
	registerFunc(&_cVMetalTextureCacheFlush, &_cVMetalTextureCacheFlushErr, frameworkHandle, "CVMetalTextureCacheFlush", "10.11")
	registerFunc(&_cVMetalTextureCacheGetTypeID, &_cVMetalTextureCacheGetTypeIDErr, frameworkHandle, "CVMetalTextureCacheGetTypeID", "10.11")
	registerFunc(&_cVMetalTextureGetCleanTexCoords, &_cVMetalTextureGetCleanTexCoordsErr, frameworkHandle, "CVMetalTextureGetCleanTexCoords", "10.11")
	registerFunc(&_cVMetalTextureGetTexture, &_cVMetalTextureGetTextureErr, frameworkHandle, "CVMetalTextureGetTexture", "10.11")
	registerFunc(&_cVMetalTextureGetTypeID, &_cVMetalTextureGetTypeIDErr, frameworkHandle, "CVMetalTextureGetTypeID", "10.11")
	registerFunc(&_cVMetalTextureIsFlipped, &_cVMetalTextureIsFlippedErr, frameworkHandle, "CVMetalTextureIsFlipped", "10.11")
	registerFunc(&_cVOpenGLBufferPoolRelease, &_cVOpenGLBufferPoolReleaseErr, frameworkHandle, "CVOpenGLBufferPoolRelease", "10.4")
	registerFunc(&_cVOpenGLBufferPoolRetain, &_cVOpenGLBufferPoolRetainErr, frameworkHandle, "CVOpenGLBufferPoolRetain", "10.4")
	registerFunc(&_cVOpenGLBufferRelease, &_cVOpenGLBufferReleaseErr, frameworkHandle, "CVOpenGLBufferRelease", "10.4")
	registerFunc(&_cVOpenGLBufferRetain, &_cVOpenGLBufferRetainErr, frameworkHandle, "CVOpenGLBufferRetain", "10.4")
	registerFunc(&_cVOpenGLTextureCacheRelease, &_cVOpenGLTextureCacheReleaseErr, frameworkHandle, "CVOpenGLTextureCacheRelease", "10.4")
	registerFunc(&_cVOpenGLTextureCacheRetain, &_cVOpenGLTextureCacheRetainErr, frameworkHandle, "CVOpenGLTextureCacheRetain", "10.4")
	registerFunc(&_cVOpenGLTextureRelease, &_cVOpenGLTextureReleaseErr, frameworkHandle, "CVOpenGLTextureRelease", "10.4")
	registerFunc(&_cVOpenGLTextureRetain, &_cVOpenGLTextureRetainErr, frameworkHandle, "CVOpenGLTextureRetain", "10.4")
	registerFunc(&_cVPixelBufferCopyCreationAttributes, &_cVPixelBufferCopyCreationAttributesErr, frameworkHandle, "CVPixelBufferCopyCreationAttributes", "12.0")
	registerFunc(&_cVPixelBufferCreate, &_cVPixelBufferCreateErr, frameworkHandle, "CVPixelBufferCreate", "10.4")
	registerFunc(&_cVPixelBufferCreateResolvedAttributesDictionary, &_cVPixelBufferCreateResolvedAttributesDictionaryErr, frameworkHandle, "CVPixelBufferCreateResolvedAttributesDictionary", "10.4")
	registerFunc(&_cVPixelBufferCreateWithBytes, &_cVPixelBufferCreateWithBytesErr, frameworkHandle, "CVPixelBufferCreateWithBytes", "10.4")
	registerFunc(&_cVPixelBufferCreateWithIOSurface, &_cVPixelBufferCreateWithIOSurfaceErr, frameworkHandle, "CVPixelBufferCreateWithIOSurface", "10.6")
	registerFunc(&_cVPixelBufferCreateWithPlanarBytes, &_cVPixelBufferCreateWithPlanarBytesErr, frameworkHandle, "CVPixelBufferCreateWithPlanarBytes", "10.4")
	registerFunc(&_cVPixelBufferFillExtendedPixels, &_cVPixelBufferFillExtendedPixelsErr, frameworkHandle, "CVPixelBufferFillExtendedPixels", "10.4")
	registerFunc(&_cVPixelBufferGetBaseAddress, &_cVPixelBufferGetBaseAddressErr, frameworkHandle, "CVPixelBufferGetBaseAddress", "10.4")
	registerFunc(&_cVPixelBufferGetBaseAddressOfPlane, &_cVPixelBufferGetBaseAddressOfPlaneErr, frameworkHandle, "CVPixelBufferGetBaseAddressOfPlane", "10.4")
	registerFunc(&_cVPixelBufferGetBytesPerRow, &_cVPixelBufferGetBytesPerRowErr, frameworkHandle, "CVPixelBufferGetBytesPerRow", "10.4")
	registerFunc(&_cVPixelBufferGetBytesPerRowOfPlane, &_cVPixelBufferGetBytesPerRowOfPlaneErr, frameworkHandle, "CVPixelBufferGetBytesPerRowOfPlane", "10.4")
	registerFunc(&_cVPixelBufferGetDataSize, &_cVPixelBufferGetDataSizeErr, frameworkHandle, "CVPixelBufferGetDataSize", "10.4")
	registerFunc(&_cVPixelBufferGetExtendedPixels, &_cVPixelBufferGetExtendedPixelsErr, frameworkHandle, "CVPixelBufferGetExtendedPixels", "10.4")
	registerFunc(&_cVPixelBufferGetHeight, &_cVPixelBufferGetHeightErr, frameworkHandle, "CVPixelBufferGetHeight", "10.4")
	registerFunc(&_cVPixelBufferGetHeightOfPlane, &_cVPixelBufferGetHeightOfPlaneErr, frameworkHandle, "CVPixelBufferGetHeightOfPlane", "10.4")
	registerFunc(&_cVPixelBufferGetIOSurface, &_cVPixelBufferGetIOSurfaceErr, frameworkHandle, "CVPixelBufferGetIOSurface", "10.6")
	registerFunc(&_cVPixelBufferGetPixelFormatType, &_cVPixelBufferGetPixelFormatTypeErr, frameworkHandle, "CVPixelBufferGetPixelFormatType", "10.4")
	registerFunc(&_cVPixelBufferGetPlaneCount, &_cVPixelBufferGetPlaneCountErr, frameworkHandle, "CVPixelBufferGetPlaneCount", "10.4")
	registerFunc(&_cVPixelBufferGetTypeID, &_cVPixelBufferGetTypeIDErr, frameworkHandle, "CVPixelBufferGetTypeID", "10.4")
	registerFunc(&_cVPixelBufferGetWidth, &_cVPixelBufferGetWidthErr, frameworkHandle, "CVPixelBufferGetWidth", "10.4")
	registerFunc(&_cVPixelBufferGetWidthOfPlane, &_cVPixelBufferGetWidthOfPlaneErr, frameworkHandle, "CVPixelBufferGetWidthOfPlane", "10.4")
	registerFunc(&_cVPixelBufferIsCompatibleWithAttributes, &_cVPixelBufferIsCompatibleWithAttributesErr, frameworkHandle, "CVPixelBufferIsCompatibleWithAttributes", "10.4")
	registerFunc(&_cVPixelBufferIsPlanar, &_cVPixelBufferIsPlanarErr, frameworkHandle, "CVPixelBufferIsPlanar", "10.4")
	registerFunc(&_cVPixelBufferLockBaseAddress, &_cVPixelBufferLockBaseAddressErr, frameworkHandle, "CVPixelBufferLockBaseAddress", "10.4")
	registerFunc(&_cVPixelBufferPoolCreate, &_cVPixelBufferPoolCreateErr, frameworkHandle, "CVPixelBufferPoolCreate", "10.4")
	registerFunc(&_cVPixelBufferPoolCreatePixelBuffer, &_cVPixelBufferPoolCreatePixelBufferErr, frameworkHandle, "CVPixelBufferPoolCreatePixelBuffer", "10.4")
	registerFunc(&_cVPixelBufferPoolCreatePixelBufferWithAuxAttributes, &_cVPixelBufferPoolCreatePixelBufferWithAuxAttributesErr, frameworkHandle, "CVPixelBufferPoolCreatePixelBufferWithAuxAttributes", "10.7")
	registerFunc(&_cVPixelBufferPoolFlush, &_cVPixelBufferPoolFlushErr, frameworkHandle, "CVPixelBufferPoolFlush", "10.4")
	registerFunc(&_cVPixelBufferPoolGetAttributes, &_cVPixelBufferPoolGetAttributesErr, frameworkHandle, "CVPixelBufferPoolGetAttributes", "10.4")
	registerFunc(&_cVPixelBufferPoolGetPixelBufferAttributes, &_cVPixelBufferPoolGetPixelBufferAttributesErr, frameworkHandle, "CVPixelBufferPoolGetPixelBufferAttributes", "10.4")
	registerFunc(&_cVPixelBufferPoolGetTypeID, &_cVPixelBufferPoolGetTypeIDErr, frameworkHandle, "CVPixelBufferPoolGetTypeID", "10.4")
	registerFunc(&_cVPixelBufferPoolRelease, &_cVPixelBufferPoolReleaseErr, frameworkHandle, "CVPixelBufferPoolRelease", "10.4")
	registerFunc(&_cVPixelBufferPoolRetain, &_cVPixelBufferPoolRetainErr, frameworkHandle, "CVPixelBufferPoolRetain", "10.4")
	registerFunc(&_cVPixelBufferRelease, &_cVPixelBufferReleaseErr, frameworkHandle, "CVPixelBufferRelease", "10.4")
	registerFunc(&_cVPixelBufferRetain, &_cVPixelBufferRetainErr, frameworkHandle, "CVPixelBufferRetain", "10.4")
	registerFunc(&_cVPixelBufferUnlockBaseAddress, &_cVPixelBufferUnlockBaseAddressErr, frameworkHandle, "CVPixelBufferUnlockBaseAddress", "10.4")
	registerFunc(&_cVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes, &_cVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypesErr, frameworkHandle, "CVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes", "10.4")
	registerFunc(&_cVPixelFormatDescriptionCreateWithPixelFormatType, &_cVPixelFormatDescriptionCreateWithPixelFormatTypeErr, frameworkHandle, "CVPixelFormatDescriptionCreateWithPixelFormatType", "10.4")
	registerFunc(&_cVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType, &_cVPixelFormatDescriptionRegisterDescriptionWithPixelFormatTypeErr, frameworkHandle, "CVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType", "10.4")
	registerFunc(&_cVPixelFormatTypeCopyFourCharCodeString, &_cVPixelFormatTypeCopyFourCharCodeStringErr, frameworkHandle, "CVPixelFormatTypeCopyFourCharCodeString", "26.0")
	registerFunc(&_cVTransferFunctionGetIntegerCodePointForString, &_cVTransferFunctionGetIntegerCodePointForStringErr, frameworkHandle, "CVTransferFunctionGetIntegerCodePointForString", "10.13")
	registerFunc(&_cVTransferFunctionGetStringForIntegerCodePoint, &_cVTransferFunctionGetStringForIntegerCodePointErr, frameworkHandle, "CVTransferFunctionGetStringForIntegerCodePoint", "10.13")
	registerFunc(&_cVYCbCrMatrixGetIntegerCodePointForString, &_cVYCbCrMatrixGetIntegerCodePointForStringErr, frameworkHandle, "CVYCbCrMatrixGetIntegerCodePointForString", "10.13")
	registerFunc(&_cVYCbCrMatrixGetStringForIntegerCodePoint, &_cVYCbCrMatrixGetStringForIntegerCodePointErr, frameworkHandle, "CVYCbCrMatrixGetStringForIntegerCodePoint", "10.13")
}

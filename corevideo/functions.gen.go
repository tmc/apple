// Code generated from Apple documentation for CoreVideo. DO NOT EDIT.

package corevideo

import (
	"fmt"
	"os"
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/metal"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: CoreVideo: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: CoreVideo: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: CoreVideo: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _cVBufferCopyAttachment func(buffer CVBufferRef, key corefoundation.CFStringRef, attachmentMode *CVAttachmentMode) corefoundation.CFTypeRef

// CVBufferCopyAttachment returns a copy of an attachment from a Core Video buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferCopyAttachment(_:_:_:)
func CVBufferCopyAttachment(buffer CVBufferRef, key corefoundation.CFStringRef, attachmentMode *CVAttachmentMode) corefoundation.CFTypeRef {
	if _cVBufferCopyAttachment == nil {
		panic("CoreVideo: symbol CVBufferCopyAttachment not loaded")
	}
	return _cVBufferCopyAttachment(buffer, key, attachmentMode)
}

var _cVBufferCopyAttachments func(buffer CVBufferRef, attachmentMode CVAttachmentMode) corefoundation.CFDictionaryRef

// CVBufferCopyAttachments returns a copy of all attachments from a Core Video buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferCopyAttachments(_:_:)
func CVBufferCopyAttachments(buffer CVBufferRef, attachmentMode CVAttachmentMode) corefoundation.CFDictionaryRef {
	if _cVBufferCopyAttachments == nil {
		panic("CoreVideo: symbol CVBufferCopyAttachments not loaded")
	}
	return _cVBufferCopyAttachments(buffer, attachmentMode)
}

var _cVBufferHasAttachment func(buffer CVBufferRef, key corefoundation.CFStringRef) bool

// CVBufferHasAttachment returns a Boolean value that indicates whether a Core Video buffer contains a specified attachment.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferHasAttachment(_:_:)
func CVBufferHasAttachment(buffer CVBufferRef, key corefoundation.CFStringRef) bool {
	if _cVBufferHasAttachment == nil {
		panic("CoreVideo: symbol CVBufferHasAttachment not loaded")
	}
	return _cVBufferHasAttachment(buffer, key)
}

var _cVBufferPropagateAttachments func(sourceBuffer CVBufferRef, destinationBuffer CVBufferRef)

// CVBufferPropagateAttachments copies all attachments that Core Video can propagate from one buffer to another.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferPropagateAttachments(_:_:)
func CVBufferPropagateAttachments(sourceBuffer CVBufferRef, destinationBuffer CVBufferRef) {
	if _cVBufferPropagateAttachments == nil {
		panic("CoreVideo: symbol CVBufferPropagateAttachments not loaded")
	}
	_cVBufferPropagateAttachments(sourceBuffer, destinationBuffer)
}

var _cVBufferRelease func(buffer CVBufferRef)

// CVBufferRelease releases a Core Video buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferRelease
func CVBufferRelease(buffer CVBufferRef) {
	if _cVBufferRelease == nil {
		panic("CoreVideo: symbol CVBufferRelease not loaded")
	}
	_cVBufferRelease(buffer)
}

var _cVBufferRemoveAllAttachments func(buffer CVBufferRef)

// CVBufferRemoveAllAttachments removes all attachments from a Core Video buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferRemoveAllAttachments(_:)
func CVBufferRemoveAllAttachments(buffer CVBufferRef) {
	if _cVBufferRemoveAllAttachments == nil {
		panic("CoreVideo: symbol CVBufferRemoveAllAttachments not loaded")
	}
	_cVBufferRemoveAllAttachments(buffer)
}

var _cVBufferRemoveAttachment func(buffer CVBufferRef, key corefoundation.CFStringRef)

// CVBufferRemoveAttachment removes the attachment you specify from a Core Video buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferRemoveAttachment(_:_:)
func CVBufferRemoveAttachment(buffer CVBufferRef, key corefoundation.CFStringRef) {
	if _cVBufferRemoveAttachment == nil {
		panic("CoreVideo: symbol CVBufferRemoveAttachment not loaded")
	}
	_cVBufferRemoveAttachment(buffer, key)
}

var _cVBufferRetain func(buffer CVBufferRef) CVBufferRef

// CVBufferRetain retains a Core Video buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferRetain
func CVBufferRetain(buffer CVBufferRef) CVBufferRef {
	if _cVBufferRetain == nil {
		panic("CoreVideo: symbol CVBufferRetain not loaded")
	}
	return _cVBufferRetain(buffer)
}

var _cVBufferSetAttachment func(buffer CVBufferRef, key corefoundation.CFStringRef, value corefoundation.CFTypeRef, attachmentMode CVAttachmentMode)

// CVBufferSetAttachment sets or adds an attachment to a Core Video buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferSetAttachment(_:_:_:_:)
func CVBufferSetAttachment(buffer CVBufferRef, key corefoundation.CFStringRef, value corefoundation.CFTypeRef, attachmentMode CVAttachmentMode) {
	if _cVBufferSetAttachment == nil {
		panic("CoreVideo: symbol CVBufferSetAttachment not loaded")
	}
	_cVBufferSetAttachment(buffer, key, value, attachmentMode)
}

var _cVBufferSetAttachments func(buffer CVBufferRef, theAttachments corefoundation.CFDictionaryRef, attachmentMode CVAttachmentMode)

// CVBufferSetAttachments sets a dictionary of attachments on a Core Video buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVBufferSetAttachments(_:_:_:)
func CVBufferSetAttachments(buffer CVBufferRef, theAttachments corefoundation.CFDictionaryRef, attachmentMode CVAttachmentMode) {
	if _cVBufferSetAttachments == nil {
		panic("CoreVideo: symbol CVBufferSetAttachments not loaded")
	}
	_cVBufferSetAttachments(buffer, theAttachments, attachmentMode)
}

var _cVColorPrimariesGetIntegerCodePointForString func(colorPrimariesString corefoundation.CFStringRef) int

// CVColorPrimariesGetIntegerCodePointForString returns the standard integer code point corresponding to the Core Video color primaries constant string that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVColorPrimariesGetIntegerCodePointForString(_:)
func CVColorPrimariesGetIntegerCodePointForString(colorPrimariesString corefoundation.CFStringRef) int {
	if _cVColorPrimariesGetIntegerCodePointForString == nil {
		panic("CoreVideo: symbol CVColorPrimariesGetIntegerCodePointForString not loaded")
	}
	return _cVColorPrimariesGetIntegerCodePointForString(colorPrimariesString)
}

var _cVColorPrimariesGetStringForIntegerCodePoint func(colorPrimariesCodePoint int) corefoundation.CFStringRef

// CVColorPrimariesGetStringForIntegerCodePoint returns the Core Video color primaries string corresponding to the standard integer code point that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVColorPrimariesGetStringForIntegerCodePoint(_:)
func CVColorPrimariesGetStringForIntegerCodePoint(colorPrimariesCodePoint int) corefoundation.CFStringRef {
	if _cVColorPrimariesGetStringForIntegerCodePoint == nil {
		panic("CoreVideo: symbol CVColorPrimariesGetStringForIntegerCodePoint not loaded")
	}
	return _cVColorPrimariesGetStringForIntegerCodePoint(colorPrimariesCodePoint)
}

var _cVDisplayLinkCreateWithActiveCGDisplays func(displayLinkOut *CVDisplayLinkRef) CVReturn

// CVDisplayLinkCreateWithActiveCGDisplays creates a display link capable of being used with all active displays.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkCreateWithActiveCGDisplays(_:)
func CVDisplayLinkCreateWithActiveCGDisplays(displayLinkOut *CVDisplayLinkRef) CVReturn {
	if _cVDisplayLinkCreateWithActiveCGDisplays == nil {
		panic("CoreVideo: symbol CVDisplayLinkCreateWithActiveCGDisplays not loaded")
	}
	return _cVDisplayLinkCreateWithActiveCGDisplays(displayLinkOut)
}

var _cVDisplayLinkCreateWithCGDisplay func(displayID uint32, displayLinkOut *CVDisplayLinkRef) CVReturn

// CVDisplayLinkCreateWithCGDisplay creates a display link for a single display.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkCreateWithCGDisplay(_:_:)
func CVDisplayLinkCreateWithCGDisplay(displayID uint32, displayLinkOut *CVDisplayLinkRef) CVReturn {
	if _cVDisplayLinkCreateWithCGDisplay == nil {
		panic("CoreVideo: symbol CVDisplayLinkCreateWithCGDisplay not loaded")
	}
	return _cVDisplayLinkCreateWithCGDisplay(displayID, displayLinkOut)
}

var _cVDisplayLinkCreateWithCGDisplays func(displayArray *uint32, count int, displayLinkOut *CVDisplayLinkRef) CVReturn

// CVDisplayLinkCreateWithCGDisplays creates a display link for an array of displays.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkCreateWithCGDisplays(_:_:_:)
func CVDisplayLinkCreateWithCGDisplays(displayArray *uint32, count int, displayLinkOut *CVDisplayLinkRef) CVReturn {
	if _cVDisplayLinkCreateWithCGDisplays == nil {
		panic("CoreVideo: symbol CVDisplayLinkCreateWithCGDisplays not loaded")
	}
	return _cVDisplayLinkCreateWithCGDisplays(displayArray, count, displayLinkOut)
}

var _cVDisplayLinkCreateWithOpenGLDisplayMask func(mask coregraphics.CGOpenGLDisplayMask, displayLinkOut *CVDisplayLinkRef) CVReturn

// CVDisplayLinkCreateWithOpenGLDisplayMask creates a display link from an OpenGL display mask.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkCreateWithOpenGLDisplayMask(_:_:)
func CVDisplayLinkCreateWithOpenGLDisplayMask(mask coregraphics.CGOpenGLDisplayMask, displayLinkOut *CVDisplayLinkRef) CVReturn {
	if _cVDisplayLinkCreateWithOpenGLDisplayMask == nil {
		panic("CoreVideo: symbol CVDisplayLinkCreateWithOpenGLDisplayMask not loaded")
	}
	return _cVDisplayLinkCreateWithOpenGLDisplayMask(mask, displayLinkOut)
}

var _cVDisplayLinkGetActualOutputVideoRefreshPeriod func(displayLink CVDisplayLinkRef) float64

// CVDisplayLinkGetActualOutputVideoRefreshPeriod retrieves the actual output refresh period of a display as measured by the system time.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetActualOutputVideoRefreshPeriod(_:)
func CVDisplayLinkGetActualOutputVideoRefreshPeriod(displayLink CVDisplayLinkRef) float64 {
	if _cVDisplayLinkGetActualOutputVideoRefreshPeriod == nil {
		panic("CoreVideo: symbol CVDisplayLinkGetActualOutputVideoRefreshPeriod not loaded")
	}
	return _cVDisplayLinkGetActualOutputVideoRefreshPeriod(displayLink)
}

var _cVDisplayLinkGetCurrentCGDisplay func(displayLink CVDisplayLinkRef) uint32

// CVDisplayLinkGetCurrentCGDisplay gets the current display associated with a display link.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetCurrentCGDisplay(_:)
func CVDisplayLinkGetCurrentCGDisplay(displayLink CVDisplayLinkRef) uint32 {
	if _cVDisplayLinkGetCurrentCGDisplay == nil {
		panic("CoreVideo: symbol CVDisplayLinkGetCurrentCGDisplay not loaded")
	}
	return _cVDisplayLinkGetCurrentCGDisplay(displayLink)
}

var _cVDisplayLinkGetCurrentTime func(displayLink CVDisplayLinkRef, outTime *CVTimeStamp) CVReturn

// CVDisplayLinkGetCurrentTime retrieves the current (“now”) time of a given display link.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetCurrentTime(_:_:)
func CVDisplayLinkGetCurrentTime(displayLink CVDisplayLinkRef, outTime *CVTimeStamp) CVReturn {
	if _cVDisplayLinkGetCurrentTime == nil {
		panic("CoreVideo: symbol CVDisplayLinkGetCurrentTime not loaded")
	}
	return _cVDisplayLinkGetCurrentTime(displayLink, outTime)
}

var _cVDisplayLinkGetNominalOutputVideoRefreshPeriod func(displayLink CVDisplayLinkRef) CVTime

// CVDisplayLinkGetNominalOutputVideoRefreshPeriod retrieves the nominal refresh period of a display link.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetNominalOutputVideoRefreshPeriod(_:)
func CVDisplayLinkGetNominalOutputVideoRefreshPeriod(displayLink CVDisplayLinkRef) CVTime {
	if _cVDisplayLinkGetNominalOutputVideoRefreshPeriod == nil {
		panic("CoreVideo: symbol CVDisplayLinkGetNominalOutputVideoRefreshPeriod not loaded")
	}
	return _cVDisplayLinkGetNominalOutputVideoRefreshPeriod(displayLink)
}

var _cVDisplayLinkGetOutputVideoLatency func(displayLink CVDisplayLinkRef) CVTime

// CVDisplayLinkGetOutputVideoLatency retrieves the nominal latency of a display link.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetOutputVideoLatency(_:)
func CVDisplayLinkGetOutputVideoLatency(displayLink CVDisplayLinkRef) CVTime {
	if _cVDisplayLinkGetOutputVideoLatency == nil {
		panic("CoreVideo: symbol CVDisplayLinkGetOutputVideoLatency not loaded")
	}
	return _cVDisplayLinkGetOutputVideoLatency(displayLink)
}

var _cVDisplayLinkGetTypeID func() uint

// CVDisplayLinkGetTypeID obtains the Core Foundation ID for the display link data type.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkGetTypeID()
func CVDisplayLinkGetTypeID() uint {
	if _cVDisplayLinkGetTypeID == nil {
		panic("CoreVideo: symbol CVDisplayLinkGetTypeID not loaded")
	}
	return _cVDisplayLinkGetTypeID()
}

var _cVDisplayLinkIsRunning func(displayLink CVDisplayLinkRef) bool

// CVDisplayLinkIsRunning indicates whether a given display link is running.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkIsRunning(_:)
func CVDisplayLinkIsRunning(displayLink CVDisplayLinkRef) bool {
	if _cVDisplayLinkIsRunning == nil {
		panic("CoreVideo: symbol CVDisplayLinkIsRunning not loaded")
	}
	return _cVDisplayLinkIsRunning(displayLink)
}

var _cVDisplayLinkRelease func(displayLink CVDisplayLinkRef)

// CVDisplayLinkRelease releases a display link.
//
// Deprecated: Deprecated since macOS 15.0.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkRelease
func CVDisplayLinkRelease(displayLink CVDisplayLinkRef) {
	if _cVDisplayLinkRelease == nil {
		panic("CoreVideo: symbol CVDisplayLinkRelease not loaded")
	}
	_cVDisplayLinkRelease(displayLink)
}

var _cVDisplayLinkRetain func(displayLink CVDisplayLinkRef) CVDisplayLinkRef

// CVDisplayLinkRetain retains a display link.
//
// Deprecated: Deprecated since macOS 15.0.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkRetain
func CVDisplayLinkRetain(displayLink CVDisplayLinkRef) CVDisplayLinkRef {
	if _cVDisplayLinkRetain == nil {
		panic("CoreVideo: symbol CVDisplayLinkRetain not loaded")
	}
	return _cVDisplayLinkRetain(displayLink)
}

var _cVDisplayLinkSetCurrentCGDisplay func(displayLink CVDisplayLinkRef, displayID uint32) CVReturn

// CVDisplayLinkSetCurrentCGDisplay sets the current display of a display link.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkSetCurrentCGDisplay(_:_:)
func CVDisplayLinkSetCurrentCGDisplay(displayLink CVDisplayLinkRef, displayID uint32) CVReturn {
	if _cVDisplayLinkSetCurrentCGDisplay == nil {
		panic("CoreVideo: symbol CVDisplayLinkSetCurrentCGDisplay not loaded")
	}
	return _cVDisplayLinkSetCurrentCGDisplay(displayLink, displayID)
}

var _cVDisplayLinkSetCurrentCGDisplayFromOpenGLContext func(displayLink CVDisplayLinkRef, cglContext unsafe.Pointer, cglPixelFormat unsafe.Pointer) CVReturn

// CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext selects the display link most optimal for the current renderer of an OpenGL context.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext(_:_:_:)
func CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext(displayLink CVDisplayLinkRef, cglContext unsafe.Pointer, cglPixelFormat unsafe.Pointer) CVReturn {
	if _cVDisplayLinkSetCurrentCGDisplayFromOpenGLContext == nil {
		panic("CoreVideo: symbol CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext not loaded")
	}
	return _cVDisplayLinkSetCurrentCGDisplayFromOpenGLContext(displayLink, cglContext, cglPixelFormat)
}

var _cVDisplayLinkSetOutputCallback func(displayLink CVDisplayLinkRef, callback CVDisplayLinkOutputCallback, userInfo unsafe.Pointer) CVReturn

// CVDisplayLinkSetOutputCallback sets the renderer output callback function.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkSetOutputCallback(_:_:_:)
func CVDisplayLinkSetOutputCallback(displayLink CVDisplayLinkRef, callback CVDisplayLinkOutputCallback, userInfo unsafe.Pointer) CVReturn {
	if _cVDisplayLinkSetOutputCallback == nil {
		panic("CoreVideo: symbol CVDisplayLinkSetOutputCallback not loaded")
	}
	return _cVDisplayLinkSetOutputCallback(displayLink, callback, userInfo)
}

var _cVDisplayLinkSetOutputHandler func(displayLink CVDisplayLinkRef, handler CVDisplayLinkOutputHandler) CVReturn

// CVDisplayLinkSetOutputHandler.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkSetOutputHandler(_:_:)
func CVDisplayLinkSetOutputHandler(displayLink CVDisplayLinkRef, handler CVDisplayLinkOutputHandler) CVReturn {
	if _cVDisplayLinkSetOutputHandler == nil {
		panic("CoreVideo: symbol CVDisplayLinkSetOutputHandler not loaded")
	}
	return _cVDisplayLinkSetOutputHandler(displayLink, handler)
}

var _cVDisplayLinkStart func(displayLink CVDisplayLinkRef) CVReturn

// CVDisplayLinkStart activates a display link.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkStart(_:)
func CVDisplayLinkStart(displayLink CVDisplayLinkRef) CVReturn {
	if _cVDisplayLinkStart == nil {
		panic("CoreVideo: symbol CVDisplayLinkStart not loaded")
	}
	return _cVDisplayLinkStart(displayLink)
}

var _cVDisplayLinkStop func(displayLink CVDisplayLinkRef) CVReturn

// CVDisplayLinkStop stops a display link.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkStop(_:)
func CVDisplayLinkStop(displayLink CVDisplayLinkRef) CVReturn {
	if _cVDisplayLinkStop == nil {
		panic("CoreVideo: symbol CVDisplayLinkStop not loaded")
	}
	return _cVDisplayLinkStop(displayLink)
}

var _cVDisplayLinkTranslateTime func(displayLink CVDisplayLinkRef, inTime *CVTimeStamp, outTime *CVTimeStamp) CVReturn

// CVDisplayLinkTranslateTime translates the time in the display link’s time base from one representation to another.
//
// Deprecated: Deprecated since macOS 15.0. use NSView.displayLink(target:selector:), NSWindow.displayLink(target:selector:), or NSScreen.displayLink(target:selector:)
//
// See: https://developer.apple.com/documentation/CoreVideo/CVDisplayLinkTranslateTime(_:_:_:)
func CVDisplayLinkTranslateTime(displayLink CVDisplayLinkRef, inTime *CVTimeStamp, outTime *CVTimeStamp) CVReturn {
	if _cVDisplayLinkTranslateTime == nil {
		panic("CoreVideo: symbol CVDisplayLinkTranslateTime not loaded")
	}
	return _cVDisplayLinkTranslateTime(displayLink, inTime, outTime)
}

var _cVGetCurrentHostTime func() uint64

// CVGetCurrentHostTime returns the current system time.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVGetCurrentHostTime()
func CVGetCurrentHostTime() uint64 {
	if _cVGetCurrentHostTime == nil {
		panic("CoreVideo: symbol CVGetCurrentHostTime not loaded")
	}
	return _cVGetCurrentHostTime()
}

var _cVGetHostClockFrequency func() float64

// CVGetHostClockFrequency returns the frequency of updates to the system time.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVGetHostClockFrequency()
func CVGetHostClockFrequency() float64 {
	if _cVGetHostClockFrequency == nil {
		panic("CoreVideo: symbol CVGetHostClockFrequency not loaded")
	}
	return _cVGetHostClockFrequency()
}

var _cVGetHostClockMinimumTimeDelta func() uint32

// CVGetHostClockMinimumTimeDelta returns the smallest possible increment in the system time.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVGetHostClockMinimumTimeDelta()
func CVGetHostClockMinimumTimeDelta() uint32 {
	if _cVGetHostClockMinimumTimeDelta == nil {
		panic("CoreVideo: symbol CVGetHostClockMinimumTimeDelta not loaded")
	}
	return _cVGetHostClockMinimumTimeDelta()
}

var _cVImageBufferCreateColorSpaceFromAttachments func(attachments corefoundation.CFDictionaryRef) coregraphics.CGColorSpaceRef

// CVImageBufferCreateColorSpaceFromAttachments attempts to create a Core Graphics color space from the image buffer’s attachments that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVImageBufferCreateColorSpaceFromAttachments(_:)
func CVImageBufferCreateColorSpaceFromAttachments(attachments corefoundation.CFDictionaryRef) coregraphics.CGColorSpaceRef {
	if _cVImageBufferCreateColorSpaceFromAttachments == nil {
		panic("CoreVideo: symbol CVImageBufferCreateColorSpaceFromAttachments not loaded")
	}
	return _cVImageBufferCreateColorSpaceFromAttachments(attachments)
}

var _cVImageBufferGetCleanRect func(imageBuffer CVImageBufferRef) corefoundation.CGRect

// CVImageBufferGetCleanRect returns the source rectangle of a Core Video image buffer that represents the clean aperture of the buffer in encoded pixels.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVImageBufferGetCleanRect(_:)
func CVImageBufferGetCleanRect(imageBuffer CVImageBufferRef) corefoundation.CGRect {
	if _cVImageBufferGetCleanRect == nil {
		panic("CoreVideo: symbol CVImageBufferGetCleanRect not loaded")
	}
	return _cVImageBufferGetCleanRect(imageBuffer)
}

var _cVImageBufferGetColorSpace func(imageBuffer CVImageBufferRef) coregraphics.CGColorSpaceRef

// CVImageBufferGetColorSpace returns the color space of a Core Video image buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVImageBufferGetColorSpace(_:)
func CVImageBufferGetColorSpace(imageBuffer CVImageBufferRef) coregraphics.CGColorSpaceRef {
	if _cVImageBufferGetColorSpace == nil {
		panic("CoreVideo: symbol CVImageBufferGetColorSpace not loaded")
	}
	return _cVImageBufferGetColorSpace(imageBuffer)
}

var _cVImageBufferGetDisplaySize func(imageBuffer CVImageBufferRef) corefoundation.CGSize

// CVImageBufferGetDisplaySize returns the nominal output display size, in square pixels, of a Core Video image buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVImageBufferGetDisplaySize(_:)
func CVImageBufferGetDisplaySize(imageBuffer CVImageBufferRef) corefoundation.CGSize {
	if _cVImageBufferGetDisplaySize == nil {
		panic("CoreVideo: symbol CVImageBufferGetDisplaySize not loaded")
	}
	return _cVImageBufferGetDisplaySize(imageBuffer)
}

var _cVImageBufferGetEncodedSize func(imageBuffer CVImageBufferRef) corefoundation.CGSize

// CVImageBufferGetEncodedSize returns the full encoded dimensions of a Core Video image buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVImageBufferGetEncodedSize(_:)
func CVImageBufferGetEncodedSize(imageBuffer CVImageBufferRef) corefoundation.CGSize {
	if _cVImageBufferGetEncodedSize == nil {
		panic("CoreVideo: symbol CVImageBufferGetEncodedSize not loaded")
	}
	return _cVImageBufferGetEncodedSize(imageBuffer)
}

var _cVImageBufferIsFlipped func(imageBuffer CVImageBufferRef) bool

// CVImageBufferIsFlipped returns a Boolean value indicating whether the image is vertically flipped.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVImageBufferIsFlipped(_:)
func CVImageBufferIsFlipped(imageBuffer CVImageBufferRef) bool {
	if _cVImageBufferIsFlipped == nil {
		panic("CoreVideo: symbol CVImageBufferIsFlipped not loaded")
	}
	return _cVImageBufferIsFlipped(imageBuffer)
}

var _cVIsCompressedPixelFormatAvailable func(pixelFormatType uint32) bool

// CVIsCompressedPixelFormatAvailable.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVIsCompressedPixelFormatAvailable(_:)
func CVIsCompressedPixelFormatAvailable(pixelFormatType uint32) bool {
	if _cVIsCompressedPixelFormatAvailable == nil {
		panic("CoreVideo: symbol CVIsCompressedPixelFormatAvailable not loaded")
	}
	return _cVIsCompressedPixelFormatAvailable(pixelFormatType)
}

var _cVMetalBufferCacheCreate func(allocator corefoundation.CFAllocatorRef, cacheAttributes corefoundation.CFDictionaryRef, metalDevice metal.MTLDeviceObject, cacheOut *CVMetalBufferCacheRef) CVReturn

// CVMetalBufferCacheCreate.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferCacheCreate(_:_:_:_:)
func CVMetalBufferCacheCreate(allocator corefoundation.CFAllocatorRef, cacheAttributes corefoundation.CFDictionaryRef, metalDevice metal.MTLDeviceObject, cacheOut *CVMetalBufferCacheRef) CVReturn {
	if _cVMetalBufferCacheCreate == nil {
		panic("CoreVideo: symbol CVMetalBufferCacheCreate not loaded")
	}
	return _cVMetalBufferCacheCreate(allocator, cacheAttributes, metalDevice, cacheOut)
}

var _cVMetalBufferCacheCreateBufferFromImage func(allocator corefoundation.CFAllocatorRef, bufferCache CVMetalBufferCacheRef, imageBuffer CVImageBufferRef, bufferOut *CVMetalBufferRef) CVReturn

// CVMetalBufferCacheCreateBufferFromImage.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferCacheCreateBufferFromImage(_:_:_:_:)
func CVMetalBufferCacheCreateBufferFromImage(allocator corefoundation.CFAllocatorRef, bufferCache CVMetalBufferCacheRef, imageBuffer CVImageBufferRef, bufferOut *CVMetalBufferRef) CVReturn {
	if _cVMetalBufferCacheCreateBufferFromImage == nil {
		panic("CoreVideo: symbol CVMetalBufferCacheCreateBufferFromImage not loaded")
	}
	return _cVMetalBufferCacheCreateBufferFromImage(allocator, bufferCache, imageBuffer, bufferOut)
}

var _cVMetalBufferCacheFlush func(bufferCache CVMetalBufferCacheRef, options CVOptionFlags)

// CVMetalBufferCacheFlush.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferCacheFlush(_:_:)
func CVMetalBufferCacheFlush(bufferCache CVMetalBufferCacheRef, options CVOptionFlags) {
	if _cVMetalBufferCacheFlush == nil {
		panic("CoreVideo: symbol CVMetalBufferCacheFlush not loaded")
	}
	_cVMetalBufferCacheFlush(bufferCache, options)
}

var _cVMetalBufferCacheGetTypeID func() uint

// CVMetalBufferCacheGetTypeID.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferCacheGetTypeID()
func CVMetalBufferCacheGetTypeID() uint {
	if _cVMetalBufferCacheGetTypeID == nil {
		panic("CoreVideo: symbol CVMetalBufferCacheGetTypeID not loaded")
	}
	return _cVMetalBufferCacheGetTypeID()
}

var _cVMetalBufferGetBuffer func(buffer CVMetalBufferRef) unsafe.Pointer

// CVMetalBufferGetBuffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferGetBuffer(_:)
func CVMetalBufferGetBuffer(buffer CVMetalBufferRef) metal.MTLBufferObject {
	if _cVMetalBufferGetBuffer == nil {
		panic("CoreVideo: symbol CVMetalBufferGetBuffer not loaded")
	}
	rv := _cVMetalBufferGetBuffer(buffer)
	return metal.MTLBufferObjectFromID(objc.IDFrom(rv))
}

var _cVMetalBufferGetTypeID func() uint

// CVMetalBufferGetTypeID.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalBufferGetTypeID()
func CVMetalBufferGetTypeID() uint {
	if _cVMetalBufferGetTypeID == nil {
		panic("CoreVideo: symbol CVMetalBufferGetTypeID not loaded")
	}
	return _cVMetalBufferGetTypeID()
}

var _cVMetalTextureCacheCreate func(allocator corefoundation.CFAllocatorRef, cacheAttributes corefoundation.CFDictionaryRef, metalDevice metal.MTLDeviceObject, textureAttributes corefoundation.CFDictionaryRef, cacheOut *CVMetalTextureCacheRef) CVReturn

// CVMetalTextureCacheCreate creates a new texture cache.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureCacheCreate(_:_:_:_:_:)
func CVMetalTextureCacheCreate(allocator corefoundation.CFAllocatorRef, cacheAttributes corefoundation.CFDictionaryRef, metalDevice metal.MTLDeviceObject, textureAttributes corefoundation.CFDictionaryRef, cacheOut *CVMetalTextureCacheRef) CVReturn {
	if _cVMetalTextureCacheCreate == nil {
		panic("CoreVideo: symbol CVMetalTextureCacheCreate not loaded")
	}
	return _cVMetalTextureCacheCreate(allocator, cacheAttributes, metalDevice, textureAttributes, cacheOut)
}

var _cVMetalTextureCacheCreateTextureFromImage func(allocator corefoundation.CFAllocatorRef, textureCache CVMetalTextureCacheRef, sourceImage CVImageBufferRef, textureAttributes corefoundation.CFDictionaryRef, pixelFormat metal.MTLPixelFormat, width uintptr, height uintptr, planeIndex uintptr, textureOut *CVMetalTextureRef) CVReturn

// CVMetalTextureCacheCreateTextureFromImage creates a Core Video Metal texture buffer from an existing image buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureCacheCreateTextureFromImage(_:_:_:_:_:_:_:_:_:)
func CVMetalTextureCacheCreateTextureFromImage(allocator corefoundation.CFAllocatorRef, textureCache CVMetalTextureCacheRef, sourceImage CVImageBufferRef, textureAttributes corefoundation.CFDictionaryRef, pixelFormat metal.MTLPixelFormat, width uintptr, height uintptr, planeIndex uintptr, textureOut *CVMetalTextureRef) CVReturn {
	if _cVMetalTextureCacheCreateTextureFromImage == nil {
		panic("CoreVideo: symbol CVMetalTextureCacheCreateTextureFromImage not loaded")
	}
	return _cVMetalTextureCacheCreateTextureFromImage(allocator, textureCache, sourceImage, textureAttributes, pixelFormat, width, height, planeIndex, textureOut)
}

var _cVMetalTextureCacheFlush func(textureCache CVMetalTextureCacheRef, options CVOptionFlags)

// CVMetalTextureCacheFlush manually flushes the contents of the provided texture cache.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureCacheFlush(_:_:)
func CVMetalTextureCacheFlush(textureCache CVMetalTextureCacheRef, options CVOptionFlags) {
	if _cVMetalTextureCacheFlush == nil {
		panic("CoreVideo: symbol CVMetalTextureCacheFlush not loaded")
	}
	_cVMetalTextureCacheFlush(textureCache, options)
}

var _cVMetalTextureCacheGetTypeID func() uint

// CVMetalTextureCacheGetTypeID returns the Core Foundation type identifier for a Core Video Metal texture cache.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureCacheGetTypeID()
func CVMetalTextureCacheGetTypeID() uint {
	if _cVMetalTextureCacheGetTypeID == nil {
		panic("CoreVideo: symbol CVMetalTextureCacheGetTypeID not loaded")
	}
	return _cVMetalTextureCacheGetTypeID()
}

var _cVMetalTextureGetCleanTexCoords func(image CVMetalTextureRef, lowerLeft unsafe.Pointer, lowerRight unsafe.Pointer, upperRight unsafe.Pointer, upperLeft unsafe.Pointer)

// CVMetalTextureGetCleanTexCoords returns convenient normalized texture coordinates for the part of the image that should be displayed.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureGetCleanTexCoords(_:_:_:_:_:)
func CVMetalTextureGetCleanTexCoords(image CVMetalTextureRef, lowerLeft unsafe.Pointer, lowerRight unsafe.Pointer, upperRight unsafe.Pointer, upperLeft unsafe.Pointer) {
	if _cVMetalTextureGetCleanTexCoords == nil {
		panic("CoreVideo: symbol CVMetalTextureGetCleanTexCoords not loaded")
	}
	_cVMetalTextureGetCleanTexCoords(image, lowerLeft, lowerRight, upperRight, upperLeft)
}

var _cVMetalTextureGetTexture func(image CVMetalTextureRef) unsafe.Pointer

// CVMetalTextureGetTexture returns the Metal texture object for the image buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureGetTexture(_:)
func CVMetalTextureGetTexture(image CVMetalTextureRef) metal.MTLTextureObject {
	if _cVMetalTextureGetTexture == nil {
		panic("CoreVideo: symbol CVMetalTextureGetTexture not loaded")
	}
	rv := _cVMetalTextureGetTexture(image)
	return metal.MTLTextureObjectFromID(objc.IDFrom(rv))
}

var _cVMetalTextureGetTypeID func() uint

// CVMetalTextureGetTypeID returns the Core Foundation type identifier for a CoreVideo Metal texture-based image buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureGetTypeID()
func CVMetalTextureGetTypeID() uint {
	if _cVMetalTextureGetTypeID == nil {
		panic("CoreVideo: symbol CVMetalTextureGetTypeID not loaded")
	}
	return _cVMetalTextureGetTypeID()
}

var _cVMetalTextureIsFlipped func(image CVMetalTextureRef) bool

// CVMetalTextureIsFlipped returns a Boolean value indicating whether the texture image is vertically flipped.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVMetalTextureIsFlipped(_:)
func CVMetalTextureIsFlipped(image CVMetalTextureRef) bool {
	if _cVMetalTextureIsFlipped == nil {
		panic("CoreVideo: symbol CVMetalTextureIsFlipped not loaded")
	}
	return _cVMetalTextureIsFlipped(image)
}

var _cVOpenGLBufferPoolRelease func(openGLBufferPool CVOpenGLBufferPoolRef)

// CVOpenGLBufferPoolRelease releases an OpenGL buffer pool.
//
// Deprecated: Deprecated since macOS 10.14.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferPoolRelease
func CVOpenGLBufferPoolRelease(openGLBufferPool CVOpenGLBufferPoolRef) {
	if _cVOpenGLBufferPoolRelease == nil {
		panic("CoreVideo: symbol CVOpenGLBufferPoolRelease not loaded")
	}
	_cVOpenGLBufferPoolRelease(openGLBufferPool)
}

var _cVOpenGLBufferPoolRetain func(openGLBufferPool CVOpenGLBufferPoolRef) CVOpenGLBufferPoolRef

// CVOpenGLBufferPoolRetain retains an OpenGL buffer pool.
//
// Deprecated: Deprecated since macOS 10.14.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferPoolRetain
func CVOpenGLBufferPoolRetain(openGLBufferPool CVOpenGLBufferPoolRef) CVOpenGLBufferPoolRef {
	if _cVOpenGLBufferPoolRetain == nil {
		panic("CoreVideo: symbol CVOpenGLBufferPoolRetain not loaded")
	}
	return _cVOpenGLBufferPoolRetain(openGLBufferPool)
}

var _cVOpenGLBufferRelease func(buffer CVOpenGLBufferRef)

// CVOpenGLBufferRelease releases a Core Video OpenGL buffer.
//
// Deprecated: Deprecated since macOS 10.14.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferRelease
func CVOpenGLBufferRelease(buffer CVOpenGLBufferRef) {
	if _cVOpenGLBufferRelease == nil {
		panic("CoreVideo: symbol CVOpenGLBufferRelease not loaded")
	}
	_cVOpenGLBufferRelease(buffer)
}

var _cVOpenGLBufferRetain func(buffer CVOpenGLBufferRef) CVOpenGLBufferRef

// CVOpenGLBufferRetain retains a Core Video OpenGL buffer.
//
// Deprecated: Deprecated since macOS 10.14.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLBufferRetain
func CVOpenGLBufferRetain(buffer CVOpenGLBufferRef) CVOpenGLBufferRef {
	if _cVOpenGLBufferRetain == nil {
		panic("CoreVideo: symbol CVOpenGLBufferRetain not loaded")
	}
	return _cVOpenGLBufferRetain(buffer)
}

var _cVOpenGLTextureCacheRelease func(textureCache CVOpenGLTextureCacheRef)

// CVOpenGLTextureCacheRelease releases a texture cache object.
//
// Deprecated: Deprecated since macOS 10.14.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureCacheRelease
func CVOpenGLTextureCacheRelease(textureCache CVOpenGLTextureCacheRef) {
	if _cVOpenGLTextureCacheRelease == nil {
		panic("CoreVideo: symbol CVOpenGLTextureCacheRelease not loaded")
	}
	_cVOpenGLTextureCacheRelease(textureCache)
}

var _cVOpenGLTextureCacheRetain func(textureCache CVOpenGLTextureCacheRef) CVOpenGLTextureCacheRef

// CVOpenGLTextureCacheRetain retains a texture cache object.
//
// Deprecated: Deprecated since macOS 10.14.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureCacheRetain
func CVOpenGLTextureCacheRetain(textureCache CVOpenGLTextureCacheRef) CVOpenGLTextureCacheRef {
	if _cVOpenGLTextureCacheRetain == nil {
		panic("CoreVideo: symbol CVOpenGLTextureCacheRetain not loaded")
	}
	return _cVOpenGLTextureCacheRetain(textureCache)
}

var _cVOpenGLTextureRelease func(texture CVOpenGLTextureRef)

// CVOpenGLTextureRelease releases a Core Video OpenGL texture.
//
// Deprecated: Deprecated since macOS 10.14.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureRelease
func CVOpenGLTextureRelease(texture CVOpenGLTextureRef) {
	if _cVOpenGLTextureRelease == nil {
		panic("CoreVideo: symbol CVOpenGLTextureRelease not loaded")
	}
	_cVOpenGLTextureRelease(texture)
}

var _cVOpenGLTextureRetain func(texture CVOpenGLTextureRef) CVOpenGLTextureRef

// CVOpenGLTextureRetain retains a Core Video OpenGL texture.
//
// Deprecated: Deprecated since macOS 10.14.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVOpenGLTextureRetain
func CVOpenGLTextureRetain(texture CVOpenGLTextureRef) CVOpenGLTextureRef {
	if _cVOpenGLTextureRetain == nil {
		panic("CoreVideo: symbol CVOpenGLTextureRetain not loaded")
	}
	return _cVOpenGLTextureRetain(texture)
}

var _cVPixelBufferCopyCreationAttributes func(pixelBuffer CVPixelBufferRef) corefoundation.CFDictionaryRef

// CVPixelBufferCopyCreationAttributes.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCopyCreationAttributes(_:)
func CVPixelBufferCopyCreationAttributes(pixelBuffer CVPixelBufferRef) corefoundation.CFDictionaryRef {
	if _cVPixelBufferCopyCreationAttributes == nil {
		panic("CoreVideo: symbol CVPixelBufferCopyCreationAttributes not loaded")
	}
	return _cVPixelBufferCopyCreationAttributes(pixelBuffer)
}

var _cVPixelBufferCreate func(allocator corefoundation.CFAllocatorRef, width uintptr, height uintptr, pixelFormatType uint32, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn

// CVPixelBufferCreate creates a single pixel buffer for a given size and pixel format.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreate(_:_:_:_:_:_:)
func CVPixelBufferCreate(allocator corefoundation.CFAllocatorRef, width uintptr, height uintptr, pixelFormatType uint32, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn {
	if _cVPixelBufferCreate == nil {
		panic("CoreVideo: symbol CVPixelBufferCreate not loaded")
	}
	return _cVPixelBufferCreate(allocator, width, height, pixelFormatType, pixelBufferAttributes, pixelBufferOut)
}

var _cVPixelBufferCreateResolvedAttributesDictionary func(allocator corefoundation.CFAllocatorRef, attributes corefoundation.CFArrayRef, resolvedDictionaryOut *corefoundation.CFDictionaryRef) CVReturn

// CVPixelBufferCreateResolvedAttributesDictionary resolves an array of [CFDictionary] objects describing various pixel buffer attributes into a single dictionary.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreateResolvedAttributesDictionary(_:_:_:)
func CVPixelBufferCreateResolvedAttributesDictionary(allocator corefoundation.CFAllocatorRef, attributes corefoundation.CFArrayRef, resolvedDictionaryOut *corefoundation.CFDictionaryRef) CVReturn {
	if _cVPixelBufferCreateResolvedAttributesDictionary == nil {
		panic("CoreVideo: symbol CVPixelBufferCreateResolvedAttributesDictionary not loaded")
	}
	return _cVPixelBufferCreateResolvedAttributesDictionary(allocator, attributes, resolvedDictionaryOut)
}

var _cVPixelBufferCreateWithBytes func(allocator corefoundation.CFAllocatorRef, width uintptr, height uintptr, pixelFormatType uint32, baseAddress unsafe.Pointer, bytesPerRow uintptr, releaseCallback CVPixelBufferReleaseBytesCallback, releaseRefCon unsafe.Pointer, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn

// CVPixelBufferCreateWithBytes creates a pixel buffer for a given size and pixel format containing data specified by a memory location.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreateWithBytes(_:_:_:_:_:_:_:_:_:_:)
func CVPixelBufferCreateWithBytes(allocator corefoundation.CFAllocatorRef, width uintptr, height uintptr, pixelFormatType uint32, baseAddress unsafe.Pointer, bytesPerRow uintptr, releaseCallback CVPixelBufferReleaseBytesCallback, releaseRefCon unsafe.Pointer, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn {
	if _cVPixelBufferCreateWithBytes == nil {
		panic("CoreVideo: symbol CVPixelBufferCreateWithBytes not loaded")
	}
	return _cVPixelBufferCreateWithBytes(allocator, width, height, pixelFormatType, baseAddress, bytesPerRow, releaseCallback, releaseRefCon, pixelBufferAttributes, pixelBufferOut)
}

var _cVPixelBufferCreateWithIOSurface func(allocator corefoundation.CFAllocatorRef, surface unsafe.Pointer, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn

// CVPixelBufferCreateWithIOSurface creates a single pixel buffer for the IO surface that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreateWithIOSurface(_:_:_:_:)
func CVPixelBufferCreateWithIOSurface(allocator corefoundation.CFAllocatorRef, surface unsafe.Pointer, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn {
	if _cVPixelBufferCreateWithIOSurface == nil {
		panic("CoreVideo: symbol CVPixelBufferCreateWithIOSurface not loaded")
	}
	return _cVPixelBufferCreateWithIOSurface(allocator, surface, pixelBufferAttributes, pixelBufferOut)
}

var _cVPixelBufferCreateWithPlanarBytes func(allocator corefoundation.CFAllocatorRef, width uintptr, height uintptr, pixelFormatType uint32, dataPtr unsafe.Pointer, dataSize uintptr, numberOfPlanes uintptr, planeBaseAddress unsafe.Pointer, planeWidth uintptr, planeHeight uintptr, planeBytesPerRow uintptr, releaseCallback CVPixelBufferReleasePlanarBytesCallback, releaseRefCon unsafe.Pointer, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn

// CVPixelBufferCreateWithPlanarBytes creates a single pixel buffer in planar format for a given size and pixel format containing data specified by a memory location.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferCreateWithPlanarBytes(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func CVPixelBufferCreateWithPlanarBytes(allocator corefoundation.CFAllocatorRef, width uintptr, height uintptr, pixelFormatType uint32, dataPtr unsafe.Pointer, dataSize uintptr, numberOfPlanes uintptr, planeBaseAddress unsafe.Pointer, planeWidth uintptr, planeHeight uintptr, planeBytesPerRow uintptr, releaseCallback CVPixelBufferReleasePlanarBytesCallback, releaseRefCon unsafe.Pointer, pixelBufferAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn {
	if _cVPixelBufferCreateWithPlanarBytes == nil {
		panic("CoreVideo: symbol CVPixelBufferCreateWithPlanarBytes not loaded")
	}
	return _cVPixelBufferCreateWithPlanarBytes(allocator, width, height, pixelFormatType, dataPtr, dataSize, numberOfPlanes, planeBaseAddress, planeWidth, planeHeight, planeBytesPerRow, releaseCallback, releaseRefCon, pixelBufferAttributes, pixelBufferOut)
}

var _cVPixelBufferFillExtendedPixels func(pixelBuffer CVPixelBufferRef) CVReturn

// CVPixelBufferFillExtendedPixels fills the extended pixels of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferFillExtendedPixels(_:)
func CVPixelBufferFillExtendedPixels(pixelBuffer CVPixelBufferRef) CVReturn {
	if _cVPixelBufferFillExtendedPixels == nil {
		panic("CoreVideo: symbol CVPixelBufferFillExtendedPixels not loaded")
	}
	return _cVPixelBufferFillExtendedPixels(pixelBuffer)
}

var _cVPixelBufferGetBaseAddress func(pixelBuffer CVPixelBufferRef) unsafe.Pointer

// CVPixelBufferGetBaseAddress returns the base address of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetBaseAddress(_:)
func CVPixelBufferGetBaseAddress(pixelBuffer CVPixelBufferRef) unsafe.Pointer {
	if _cVPixelBufferGetBaseAddress == nil {
		panic("CoreVideo: symbol CVPixelBufferGetBaseAddress not loaded")
	}
	return _cVPixelBufferGetBaseAddress(pixelBuffer)
}

var _cVPixelBufferGetBaseAddressOfPlane func(pixelBuffer CVPixelBufferRef, planeIndex uintptr) unsafe.Pointer

// CVPixelBufferGetBaseAddressOfPlane returns the base address of the plane at the specified plane index.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetBaseAddressOfPlane(_:_:)
func CVPixelBufferGetBaseAddressOfPlane(pixelBuffer CVPixelBufferRef, planeIndex uintptr) unsafe.Pointer {
	if _cVPixelBufferGetBaseAddressOfPlane == nil {
		panic("CoreVideo: symbol CVPixelBufferGetBaseAddressOfPlane not loaded")
	}
	return _cVPixelBufferGetBaseAddressOfPlane(pixelBuffer, planeIndex)
}

var _cVPixelBufferGetBytesPerRow func(pixelBuffer CVPixelBufferRef) uintptr

// CVPixelBufferGetBytesPerRow returns the number of bytes per row of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetBytesPerRow(_:)
func CVPixelBufferGetBytesPerRow(pixelBuffer CVPixelBufferRef) uintptr {
	if _cVPixelBufferGetBytesPerRow == nil {
		panic("CoreVideo: symbol CVPixelBufferGetBytesPerRow not loaded")
	}
	return _cVPixelBufferGetBytesPerRow(pixelBuffer)
}

var _cVPixelBufferGetBytesPerRowOfPlane func(pixelBuffer CVPixelBufferRef, planeIndex uintptr) uintptr

// CVPixelBufferGetBytesPerRowOfPlane returns the number of bytes per row for a plane at the specified index in the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetBytesPerRowOfPlane(_:_:)
func CVPixelBufferGetBytesPerRowOfPlane(pixelBuffer CVPixelBufferRef, planeIndex uintptr) uintptr {
	if _cVPixelBufferGetBytesPerRowOfPlane == nil {
		panic("CoreVideo: symbol CVPixelBufferGetBytesPerRowOfPlane not loaded")
	}
	return _cVPixelBufferGetBytesPerRowOfPlane(pixelBuffer, planeIndex)
}

var _cVPixelBufferGetDataSize func(pixelBuffer CVPixelBufferRef) uintptr

// CVPixelBufferGetDataSize returns the data size for contiguous planes of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetDataSize(_:)
func CVPixelBufferGetDataSize(pixelBuffer CVPixelBufferRef) uintptr {
	if _cVPixelBufferGetDataSize == nil {
		panic("CoreVideo: symbol CVPixelBufferGetDataSize not loaded")
	}
	return _cVPixelBufferGetDataSize(pixelBuffer)
}

var _cVPixelBufferGetExtendedPixels func(pixelBuffer CVPixelBufferRef, extraColumnsOnLeft *uintptr, extraColumnsOnRight *uintptr, extraRowsOnTop *uintptr, extraRowsOnBottom *uintptr)

// CVPixelBufferGetExtendedPixels returns the amount of extended pixel padding in the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetExtendedPixels(_:_:_:_:_:)
func CVPixelBufferGetExtendedPixels(pixelBuffer CVPixelBufferRef, extraColumnsOnLeft *uintptr, extraColumnsOnRight *uintptr, extraRowsOnTop *uintptr, extraRowsOnBottom *uintptr) {
	if _cVPixelBufferGetExtendedPixels == nil {
		panic("CoreVideo: symbol CVPixelBufferGetExtendedPixels not loaded")
	}
	_cVPixelBufferGetExtendedPixels(pixelBuffer, extraColumnsOnLeft, extraColumnsOnRight, extraRowsOnTop, extraRowsOnBottom)
}

var _cVPixelBufferGetHeight func(pixelBuffer CVPixelBufferRef) uintptr

// CVPixelBufferGetHeight returns the height of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetHeight(_:)
func CVPixelBufferGetHeight(pixelBuffer CVPixelBufferRef) uintptr {
	if _cVPixelBufferGetHeight == nil {
		panic("CoreVideo: symbol CVPixelBufferGetHeight not loaded")
	}
	return _cVPixelBufferGetHeight(pixelBuffer)
}

var _cVPixelBufferGetHeightOfPlane func(pixelBuffer CVPixelBufferRef, planeIndex uintptr) uintptr

// CVPixelBufferGetHeightOfPlane returns the height of the plane at planeIndex in the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetHeightOfPlane(_:_:)
func CVPixelBufferGetHeightOfPlane(pixelBuffer CVPixelBufferRef, planeIndex uintptr) uintptr {
	if _cVPixelBufferGetHeightOfPlane == nil {
		panic("CoreVideo: symbol CVPixelBufferGetHeightOfPlane not loaded")
	}
	return _cVPixelBufferGetHeightOfPlane(pixelBuffer, planeIndex)
}

var _cVPixelBufferGetIOSurface func(pixelBuffer CVPixelBufferRef) unsafe.Pointer

// CVPixelBufferGetIOSurface returns the IOSurface backing the pixel buffer, or [NULL] if it is not backed by an IOSurface.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetIOSurface(_:)
func CVPixelBufferGetIOSurface(pixelBuffer CVPixelBufferRef) unsafe.Pointer {
	if _cVPixelBufferGetIOSurface == nil {
		panic("CoreVideo: symbol CVPixelBufferGetIOSurface not loaded")
	}
	return _cVPixelBufferGetIOSurface(pixelBuffer)
}

var _cVPixelBufferGetPixelFormatType func(pixelBuffer CVPixelBufferRef) uint32

// CVPixelBufferGetPixelFormatType returns the pixel format type of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetPixelFormatType(_:)
func CVPixelBufferGetPixelFormatType(pixelBuffer CVPixelBufferRef) uint32 {
	if _cVPixelBufferGetPixelFormatType == nil {
		panic("CoreVideo: symbol CVPixelBufferGetPixelFormatType not loaded")
	}
	return _cVPixelBufferGetPixelFormatType(pixelBuffer)
}

var _cVPixelBufferGetPlaneCount func(pixelBuffer CVPixelBufferRef) uintptr

// CVPixelBufferGetPlaneCount returns number of planes of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetPlaneCount(_:)
func CVPixelBufferGetPlaneCount(pixelBuffer CVPixelBufferRef) uintptr {
	if _cVPixelBufferGetPlaneCount == nil {
		panic("CoreVideo: symbol CVPixelBufferGetPlaneCount not loaded")
	}
	return _cVPixelBufferGetPlaneCount(pixelBuffer)
}

var _cVPixelBufferGetTypeID func() uint

// CVPixelBufferGetTypeID returns the Core Foundation type identifier of the pixel buffer type.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetTypeID()
func CVPixelBufferGetTypeID() uint {
	if _cVPixelBufferGetTypeID == nil {
		panic("CoreVideo: symbol CVPixelBufferGetTypeID not loaded")
	}
	return _cVPixelBufferGetTypeID()
}

var _cVPixelBufferGetWidth func(pixelBuffer CVPixelBufferRef) uintptr

// CVPixelBufferGetWidth returns the width of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetWidth(_:)
func CVPixelBufferGetWidth(pixelBuffer CVPixelBufferRef) uintptr {
	if _cVPixelBufferGetWidth == nil {
		panic("CoreVideo: symbol CVPixelBufferGetWidth not loaded")
	}
	return _cVPixelBufferGetWidth(pixelBuffer)
}

var _cVPixelBufferGetWidthOfPlane func(pixelBuffer CVPixelBufferRef, planeIndex uintptr) uintptr

// CVPixelBufferGetWidthOfPlane returns the width of the plane at a given index in the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferGetWidthOfPlane(_:_:)
func CVPixelBufferGetWidthOfPlane(pixelBuffer CVPixelBufferRef, planeIndex uintptr) uintptr {
	if _cVPixelBufferGetWidthOfPlane == nil {
		panic("CoreVideo: symbol CVPixelBufferGetWidthOfPlane not loaded")
	}
	return _cVPixelBufferGetWidthOfPlane(pixelBuffer, planeIndex)
}

var _cVPixelBufferIsCompatibleWithAttributes func(pixelBuffer CVPixelBufferRef, attributes corefoundation.CFDictionaryRef) bool

// CVPixelBufferIsCompatibleWithAttributes.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferIsCompatibleWithAttributes(_:_:)
func CVPixelBufferIsCompatibleWithAttributes(pixelBuffer CVPixelBufferRef, attributes corefoundation.CFDictionaryRef) bool {
	if _cVPixelBufferIsCompatibleWithAttributes == nil {
		panic("CoreVideo: symbol CVPixelBufferIsCompatibleWithAttributes not loaded")
	}
	return _cVPixelBufferIsCompatibleWithAttributes(pixelBuffer, attributes)
}

var _cVPixelBufferIsPlanar func(pixelBuffer CVPixelBufferRef) bool

// CVPixelBufferIsPlanar determines whether the pixel buffer is planar.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferIsPlanar(_:)
func CVPixelBufferIsPlanar(pixelBuffer CVPixelBufferRef) bool {
	if _cVPixelBufferIsPlanar == nil {
		panic("CoreVideo: symbol CVPixelBufferIsPlanar not loaded")
	}
	return _cVPixelBufferIsPlanar(pixelBuffer)
}

var _cVPixelBufferLockBaseAddress func(pixelBuffer CVPixelBufferRef, lockFlags CVPixelBufferLockFlags) CVReturn

// CVPixelBufferLockBaseAddress locks the base address of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferLockBaseAddress(_:_:)
func CVPixelBufferLockBaseAddress(pixelBuffer CVPixelBufferRef, lockFlags CVPixelBufferLockFlags) CVReturn {
	if _cVPixelBufferLockBaseAddress == nil {
		panic("CoreVideo: symbol CVPixelBufferLockBaseAddress not loaded")
	}
	return _cVPixelBufferLockBaseAddress(pixelBuffer, lockFlags)
}

var _cVPixelBufferPoolCreate func(allocator corefoundation.CFAllocatorRef, poolAttributes corefoundation.CFDictionaryRef, pixelBufferAttributes corefoundation.CFDictionaryRef, poolOut *CVPixelBufferPoolRef) CVReturn

// CVPixelBufferPoolCreate creates a pixel buffer pool using the allocator and attributes that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolCreate(_:_:_:_:)
func CVPixelBufferPoolCreate(allocator corefoundation.CFAllocatorRef, poolAttributes corefoundation.CFDictionaryRef, pixelBufferAttributes corefoundation.CFDictionaryRef, poolOut *CVPixelBufferPoolRef) CVReturn {
	if _cVPixelBufferPoolCreate == nil {
		panic("CoreVideo: symbol CVPixelBufferPoolCreate not loaded")
	}
	return _cVPixelBufferPoolCreate(allocator, poolAttributes, pixelBufferAttributes, poolOut)
}

var _cVPixelBufferPoolCreatePixelBuffer func(allocator corefoundation.CFAllocatorRef, pixelBufferPool CVPixelBufferPoolRef, pixelBufferOut *CVPixelBufferRef) CVReturn

// CVPixelBufferPoolCreatePixelBuffer creates a pixel buffer from a pixel buffer pool, using the allocator that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolCreatePixelBuffer(_:_:_:)
func CVPixelBufferPoolCreatePixelBuffer(allocator corefoundation.CFAllocatorRef, pixelBufferPool CVPixelBufferPoolRef, pixelBufferOut *CVPixelBufferRef) CVReturn {
	if _cVPixelBufferPoolCreatePixelBuffer == nil {
		panic("CoreVideo: symbol CVPixelBufferPoolCreatePixelBuffer not loaded")
	}
	return _cVPixelBufferPoolCreatePixelBuffer(allocator, pixelBufferPool, pixelBufferOut)
}

var _cVPixelBufferPoolCreatePixelBufferWithAuxAttributes func(allocator corefoundation.CFAllocatorRef, pixelBufferPool CVPixelBufferPoolRef, auxAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn

// CVPixelBufferPoolCreatePixelBufferWithAuxAttributes creates a new pixel buffer with auxiliary attributes from the pool.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolCreatePixelBufferWithAuxAttributes(_:_:_:_:)
func CVPixelBufferPoolCreatePixelBufferWithAuxAttributes(allocator corefoundation.CFAllocatorRef, pixelBufferPool CVPixelBufferPoolRef, auxAttributes corefoundation.CFDictionaryRef, pixelBufferOut *CVPixelBufferRef) CVReturn {
	if _cVPixelBufferPoolCreatePixelBufferWithAuxAttributes == nil {
		panic("CoreVideo: symbol CVPixelBufferPoolCreatePixelBufferWithAuxAttributes not loaded")
	}
	return _cVPixelBufferPoolCreatePixelBufferWithAuxAttributes(allocator, pixelBufferPool, auxAttributes, pixelBufferOut)
}

var _cVPixelBufferPoolFlush func(pool CVPixelBufferPoolRef, options CVPixelBufferPoolFlushFlags)

// CVPixelBufferPoolFlush frees pixel buffers from the pool based on the options that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolFlush(_:_:)
func CVPixelBufferPoolFlush(pool CVPixelBufferPoolRef, options CVPixelBufferPoolFlushFlags) {
	if _cVPixelBufferPoolFlush == nil {
		panic("CoreVideo: symbol CVPixelBufferPoolFlush not loaded")
	}
	_cVPixelBufferPoolFlush(pool, options)
}

var _cVPixelBufferPoolGetAttributes func(pool CVPixelBufferPoolRef) corefoundation.CFDictionaryRef

// CVPixelBufferPoolGetAttributes the pool attributes dictionary for a pixel buffer pool.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolGetAttributes(_:)
func CVPixelBufferPoolGetAttributes(pool CVPixelBufferPoolRef) corefoundation.CFDictionaryRef {
	if _cVPixelBufferPoolGetAttributes == nil {
		panic("CoreVideo: symbol CVPixelBufferPoolGetAttributes not loaded")
	}
	return _cVPixelBufferPoolGetAttributes(pool)
}

var _cVPixelBufferPoolGetPixelBufferAttributes func(pool CVPixelBufferPoolRef) corefoundation.CFDictionaryRef

// CVPixelBufferPoolGetPixelBufferAttributes the attributes of pixel buffers which the system creates using the pool you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolGetPixelBufferAttributes(_:)
func CVPixelBufferPoolGetPixelBufferAttributes(pool CVPixelBufferPoolRef) corefoundation.CFDictionaryRef {
	if _cVPixelBufferPoolGetPixelBufferAttributes == nil {
		panic("CoreVideo: symbol CVPixelBufferPoolGetPixelBufferAttributes not loaded")
	}
	return _cVPixelBufferPoolGetPixelBufferAttributes(pool)
}

var _cVPixelBufferPoolGetTypeID func() uint

// CVPixelBufferPoolGetTypeID returns the Core Foundation type identifier of the pixel buffer pool type.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolGetTypeID()
func CVPixelBufferPoolGetTypeID() uint {
	if _cVPixelBufferPoolGetTypeID == nil {
		panic("CoreVideo: symbol CVPixelBufferPoolGetTypeID not loaded")
	}
	return _cVPixelBufferPoolGetTypeID()
}

var _cVPixelBufferPoolRelease func(pixelBufferPool CVPixelBufferPoolRef)

// CVPixelBufferPoolRelease releases a pixel buffer pool.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolRelease
func CVPixelBufferPoolRelease(pixelBufferPool CVPixelBufferPoolRef) {
	if _cVPixelBufferPoolRelease == nil {
		panic("CoreVideo: symbol CVPixelBufferPoolRelease not loaded")
	}
	_cVPixelBufferPoolRelease(pixelBufferPool)
}

var _cVPixelBufferPoolRetain func(pixelBufferPool CVPixelBufferPoolRef) CVPixelBufferPoolRef

// CVPixelBufferPoolRetain retains the pixel buffer pool that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferPoolRetain
func CVPixelBufferPoolRetain(pixelBufferPool CVPixelBufferPoolRef) CVPixelBufferPoolRef {
	if _cVPixelBufferPoolRetain == nil {
		panic("CoreVideo: symbol CVPixelBufferPoolRetain not loaded")
	}
	return _cVPixelBufferPoolRetain(pixelBufferPool)
}

var _cVPixelBufferRelease func(texture CVPixelBufferRef)

// CVPixelBufferRelease releases a pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferRelease
func CVPixelBufferRelease(texture CVPixelBufferRef) {
	if _cVPixelBufferRelease == nil {
		panic("CoreVideo: symbol CVPixelBufferRelease not loaded")
	}
	_cVPixelBufferRelease(texture)
}

var _cVPixelBufferRetain func(texture CVPixelBufferRef) CVPixelBufferRef

// CVPixelBufferRetain retains a pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferRetain
func CVPixelBufferRetain(texture CVPixelBufferRef) CVPixelBufferRef {
	if _cVPixelBufferRetain == nil {
		panic("CoreVideo: symbol CVPixelBufferRetain not loaded")
	}
	return _cVPixelBufferRetain(texture)
}

var _cVPixelBufferUnlockBaseAddress func(pixelBuffer CVPixelBufferRef, unlockFlags CVPixelBufferLockFlags) CVReturn

// CVPixelBufferUnlockBaseAddress unlocks the base address of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelBufferUnlockBaseAddress(_:_:)
func CVPixelBufferUnlockBaseAddress(pixelBuffer CVPixelBufferRef, unlockFlags CVPixelBufferLockFlags) CVReturn {
	if _cVPixelBufferUnlockBaseAddress == nil {
		panic("CoreVideo: symbol CVPixelBufferUnlockBaseAddress not loaded")
	}
	return _cVPixelBufferUnlockBaseAddress(pixelBuffer, unlockFlags)
}

var _cVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes func(allocator corefoundation.CFAllocatorRef) corefoundation.CFArrayRef

// CVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes returns all the pixel format descriptions known to Core Video.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes(_:)
func CVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes(allocator corefoundation.CFAllocatorRef) corefoundation.CFArrayRef {
	if _cVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes == nil {
		panic("CoreVideo: symbol CVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes not loaded")
	}
	return _cVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes(allocator)
}

var _cVPixelFormatDescriptionCreateWithPixelFormatType func(allocator corefoundation.CFAllocatorRef, pixelFormat uint32) corefoundation.CFDictionaryRef

// CVPixelFormatDescriptionCreateWithPixelFormatType creates a pixel format description from a given [OSType] identifier.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelFormatDescriptionCreateWithPixelFormatType(_:_:)
func CVPixelFormatDescriptionCreateWithPixelFormatType(allocator corefoundation.CFAllocatorRef, pixelFormat uint32) corefoundation.CFDictionaryRef {
	if _cVPixelFormatDescriptionCreateWithPixelFormatType == nil {
		panic("CoreVideo: symbol CVPixelFormatDescriptionCreateWithPixelFormatType not loaded")
	}
	return _cVPixelFormatDescriptionCreateWithPixelFormatType(allocator, pixelFormat)
}

var _cVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType func(description corefoundation.CFDictionaryRef, pixelFormat uint32)

// CVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType registers a pixel format description with Core Video.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType(_:_:)
func CVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType(description corefoundation.CFDictionaryRef, pixelFormat uint32) {
	if _cVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType == nil {
		panic("CoreVideo: symbol CVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType not loaded")
	}
	_cVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType(description, pixelFormat)
}

var _cVPixelFormatTypeCopyFourCharCodeString func(pixelFormat uint32) corefoundation.CFStringRef

// CVPixelFormatTypeCopyFourCharCodeString.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVPixelFormatTypeCopyFourCharCodeString(_:)
func CVPixelFormatTypeCopyFourCharCodeString(pixelFormat uint32) corefoundation.CFStringRef {
	if _cVPixelFormatTypeCopyFourCharCodeString == nil {
		panic("CoreVideo: symbol CVPixelFormatTypeCopyFourCharCodeString not loaded")
	}
	return _cVPixelFormatTypeCopyFourCharCodeString(pixelFormat)
}

var _cVTransferFunctionGetIntegerCodePointForString func(transferFunctionString corefoundation.CFStringRef) int

// CVTransferFunctionGetIntegerCodePointForString returns the standard integer code point corresponding to the Core Video transfer function string that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVTransferFunctionGetIntegerCodePointForString(_:)
func CVTransferFunctionGetIntegerCodePointForString(transferFunctionString corefoundation.CFStringRef) int {
	if _cVTransferFunctionGetIntegerCodePointForString == nil {
		panic("CoreVideo: symbol CVTransferFunctionGetIntegerCodePointForString not loaded")
	}
	return _cVTransferFunctionGetIntegerCodePointForString(transferFunctionString)
}

var _cVTransferFunctionGetStringForIntegerCodePoint func(transferFunctionCodePoint int) corefoundation.CFStringRef

// CVTransferFunctionGetStringForIntegerCodePoint returns the Core Video transfer function string corresponding to the standard integer code point that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVTransferFunctionGetStringForIntegerCodePoint(_:)
func CVTransferFunctionGetStringForIntegerCodePoint(transferFunctionCodePoint int) corefoundation.CFStringRef {
	if _cVTransferFunctionGetStringForIntegerCodePoint == nil {
		panic("CoreVideo: symbol CVTransferFunctionGetStringForIntegerCodePoint not loaded")
	}
	return _cVTransferFunctionGetStringForIntegerCodePoint(transferFunctionCodePoint)
}

var _cVYCbCrMatrixGetIntegerCodePointForString func(yCbCrMatrixString corefoundation.CFStringRef) int

// CVYCbCrMatrixGetIntegerCodePointForString returns the standard integer code point corresponding to the Core Video YCbCr matrix string that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVYCbCrMatrixGetIntegerCodePointForString(_:)
func CVYCbCrMatrixGetIntegerCodePointForString(yCbCrMatrixString corefoundation.CFStringRef) int {
	if _cVYCbCrMatrixGetIntegerCodePointForString == nil {
		panic("CoreVideo: symbol CVYCbCrMatrixGetIntegerCodePointForString not loaded")
	}
	return _cVYCbCrMatrixGetIntegerCodePointForString(yCbCrMatrixString)
}

var _cVYCbCrMatrixGetStringForIntegerCodePoint func(yCbCrMatrixCodePoint int) corefoundation.CFStringRef

// CVYCbCrMatrixGetStringForIntegerCodePoint returns the Core Video YCbCr matrix string corresponding to the standard integer code point that you specify.
//
// See: https://developer.apple.com/documentation/CoreVideo/CVYCbCrMatrixGetStringForIntegerCodePoint(_:)
func CVYCbCrMatrixGetStringForIntegerCodePoint(yCbCrMatrixCodePoint int) corefoundation.CFStringRef {
	if _cVYCbCrMatrixGetStringForIntegerCodePoint == nil {
		panic("CoreVideo: symbol CVYCbCrMatrixGetStringForIntegerCodePoint not loaded")
	}
	return _cVYCbCrMatrixGetStringForIntegerCodePoint(yCbCrMatrixCodePoint)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
		registerFunc(&_cVBufferCopyAttachment, frameworkHandle, "CVBufferCopyAttachment")
		registerFunc(&_cVBufferCopyAttachments, frameworkHandle, "CVBufferCopyAttachments")
		registerFunc(&_cVBufferHasAttachment, frameworkHandle, "CVBufferHasAttachment")
		registerFunc(&_cVBufferPropagateAttachments, frameworkHandle, "CVBufferPropagateAttachments")
		registerFunc(&_cVBufferRelease, frameworkHandle, "CVBufferRelease")
		registerFunc(&_cVBufferRemoveAllAttachments, frameworkHandle, "CVBufferRemoveAllAttachments")
		registerFunc(&_cVBufferRemoveAttachment, frameworkHandle, "CVBufferRemoveAttachment")
		registerFunc(&_cVBufferRetain, frameworkHandle, "CVBufferRetain")
		registerFunc(&_cVBufferSetAttachment, frameworkHandle, "CVBufferSetAttachment")
		registerFunc(&_cVBufferSetAttachments, frameworkHandle, "CVBufferSetAttachments")
		registerFunc(&_cVColorPrimariesGetIntegerCodePointForString, frameworkHandle, "CVColorPrimariesGetIntegerCodePointForString")
		registerFunc(&_cVColorPrimariesGetStringForIntegerCodePoint, frameworkHandle, "CVColorPrimariesGetStringForIntegerCodePoint")
		registerFunc(&_cVDisplayLinkCreateWithActiveCGDisplays, frameworkHandle, "CVDisplayLinkCreateWithActiveCGDisplays")
		registerFunc(&_cVDisplayLinkCreateWithCGDisplay, frameworkHandle, "CVDisplayLinkCreateWithCGDisplay")
		registerFunc(&_cVDisplayLinkCreateWithCGDisplays, frameworkHandle, "CVDisplayLinkCreateWithCGDisplays")
		registerFunc(&_cVDisplayLinkCreateWithOpenGLDisplayMask, frameworkHandle, "CVDisplayLinkCreateWithOpenGLDisplayMask")
		registerFunc(&_cVDisplayLinkGetActualOutputVideoRefreshPeriod, frameworkHandle, "CVDisplayLinkGetActualOutputVideoRefreshPeriod")
		registerFunc(&_cVDisplayLinkGetCurrentCGDisplay, frameworkHandle, "CVDisplayLinkGetCurrentCGDisplay")
		registerFunc(&_cVDisplayLinkGetCurrentTime, frameworkHandle, "CVDisplayLinkGetCurrentTime")
		registerFunc(&_cVDisplayLinkGetNominalOutputVideoRefreshPeriod, frameworkHandle, "CVDisplayLinkGetNominalOutputVideoRefreshPeriod")
		registerFunc(&_cVDisplayLinkGetOutputVideoLatency, frameworkHandle, "CVDisplayLinkGetOutputVideoLatency")
		registerFunc(&_cVDisplayLinkGetTypeID, frameworkHandle, "CVDisplayLinkGetTypeID")
		registerFunc(&_cVDisplayLinkIsRunning, frameworkHandle, "CVDisplayLinkIsRunning")
		registerFunc(&_cVDisplayLinkRelease, frameworkHandle, "CVDisplayLinkRelease")
		registerFunc(&_cVDisplayLinkRetain, frameworkHandle, "CVDisplayLinkRetain")
		registerFunc(&_cVDisplayLinkSetCurrentCGDisplay, frameworkHandle, "CVDisplayLinkSetCurrentCGDisplay")
		registerFunc(&_cVDisplayLinkSetCurrentCGDisplayFromOpenGLContext, frameworkHandle, "CVDisplayLinkSetCurrentCGDisplayFromOpenGLContext")
		registerFunc(&_cVDisplayLinkSetOutputCallback, frameworkHandle, "CVDisplayLinkSetOutputCallback")
		registerFunc(&_cVDisplayLinkSetOutputHandler, frameworkHandle, "CVDisplayLinkSetOutputHandler")
		registerFunc(&_cVDisplayLinkStart, frameworkHandle, "CVDisplayLinkStart")
		registerFunc(&_cVDisplayLinkStop, frameworkHandle, "CVDisplayLinkStop")
		registerFunc(&_cVDisplayLinkTranslateTime, frameworkHandle, "CVDisplayLinkTranslateTime")
		registerFunc(&_cVGetCurrentHostTime, frameworkHandle, "CVGetCurrentHostTime")
		registerFunc(&_cVGetHostClockFrequency, frameworkHandle, "CVGetHostClockFrequency")
		registerFunc(&_cVGetHostClockMinimumTimeDelta, frameworkHandle, "CVGetHostClockMinimumTimeDelta")
		registerFunc(&_cVImageBufferCreateColorSpaceFromAttachments, frameworkHandle, "CVImageBufferCreateColorSpaceFromAttachments")
		registerFunc(&_cVImageBufferGetCleanRect, frameworkHandle, "CVImageBufferGetCleanRect")
		registerFunc(&_cVImageBufferGetColorSpace, frameworkHandle, "CVImageBufferGetColorSpace")
		registerFunc(&_cVImageBufferGetDisplaySize, frameworkHandle, "CVImageBufferGetDisplaySize")
		registerFunc(&_cVImageBufferGetEncodedSize, frameworkHandle, "CVImageBufferGetEncodedSize")
		registerFunc(&_cVImageBufferIsFlipped, frameworkHandle, "CVImageBufferIsFlipped")
		registerFunc(&_cVIsCompressedPixelFormatAvailable, frameworkHandle, "CVIsCompressedPixelFormatAvailable")
		registerFunc(&_cVMetalBufferCacheCreate, frameworkHandle, "CVMetalBufferCacheCreate")
		registerFunc(&_cVMetalBufferCacheCreateBufferFromImage, frameworkHandle, "CVMetalBufferCacheCreateBufferFromImage")
		registerFunc(&_cVMetalBufferCacheFlush, frameworkHandle, "CVMetalBufferCacheFlush")
		registerFunc(&_cVMetalBufferCacheGetTypeID, frameworkHandle, "CVMetalBufferCacheGetTypeID")
		registerFunc(&_cVMetalBufferGetBuffer, frameworkHandle, "CVMetalBufferGetBuffer")
		registerFunc(&_cVMetalBufferGetTypeID, frameworkHandle, "CVMetalBufferGetTypeID")
		registerFunc(&_cVMetalTextureCacheCreate, frameworkHandle, "CVMetalTextureCacheCreate")
		registerFunc(&_cVMetalTextureCacheCreateTextureFromImage, frameworkHandle, "CVMetalTextureCacheCreateTextureFromImage")
		registerFunc(&_cVMetalTextureCacheFlush, frameworkHandle, "CVMetalTextureCacheFlush")
		registerFunc(&_cVMetalTextureCacheGetTypeID, frameworkHandle, "CVMetalTextureCacheGetTypeID")
		registerFunc(&_cVMetalTextureGetCleanTexCoords, frameworkHandle, "CVMetalTextureGetCleanTexCoords")
		registerFunc(&_cVMetalTextureGetTexture, frameworkHandle, "CVMetalTextureGetTexture")
		registerFunc(&_cVMetalTextureGetTypeID, frameworkHandle, "CVMetalTextureGetTypeID")
		registerFunc(&_cVMetalTextureIsFlipped, frameworkHandle, "CVMetalTextureIsFlipped")
		registerFunc(&_cVOpenGLBufferPoolRelease, frameworkHandle, "CVOpenGLBufferPoolRelease")
		registerFunc(&_cVOpenGLBufferPoolRetain, frameworkHandle, "CVOpenGLBufferPoolRetain")
		registerFunc(&_cVOpenGLBufferRelease, frameworkHandle, "CVOpenGLBufferRelease")
		registerFunc(&_cVOpenGLBufferRetain, frameworkHandle, "CVOpenGLBufferRetain")
		registerFunc(&_cVOpenGLTextureCacheRelease, frameworkHandle, "CVOpenGLTextureCacheRelease")
		registerFunc(&_cVOpenGLTextureCacheRetain, frameworkHandle, "CVOpenGLTextureCacheRetain")
		registerFunc(&_cVOpenGLTextureRelease, frameworkHandle, "CVOpenGLTextureRelease")
		registerFunc(&_cVOpenGLTextureRetain, frameworkHandle, "CVOpenGLTextureRetain")
		registerFunc(&_cVPixelBufferCopyCreationAttributes, frameworkHandle, "CVPixelBufferCopyCreationAttributes")
		registerFunc(&_cVPixelBufferCreate, frameworkHandle, "CVPixelBufferCreate")
		registerFunc(&_cVPixelBufferCreateResolvedAttributesDictionary, frameworkHandle, "CVPixelBufferCreateResolvedAttributesDictionary")
		registerFunc(&_cVPixelBufferCreateWithBytes, frameworkHandle, "CVPixelBufferCreateWithBytes")
		registerFunc(&_cVPixelBufferCreateWithIOSurface, frameworkHandle, "CVPixelBufferCreateWithIOSurface")
		registerFunc(&_cVPixelBufferCreateWithPlanarBytes, frameworkHandle, "CVPixelBufferCreateWithPlanarBytes")
		registerFunc(&_cVPixelBufferFillExtendedPixels, frameworkHandle, "CVPixelBufferFillExtendedPixels")
		registerFunc(&_cVPixelBufferGetBaseAddress, frameworkHandle, "CVPixelBufferGetBaseAddress")
		registerFunc(&_cVPixelBufferGetBaseAddressOfPlane, frameworkHandle, "CVPixelBufferGetBaseAddressOfPlane")
		registerFunc(&_cVPixelBufferGetBytesPerRow, frameworkHandle, "CVPixelBufferGetBytesPerRow")
		registerFunc(&_cVPixelBufferGetBytesPerRowOfPlane, frameworkHandle, "CVPixelBufferGetBytesPerRowOfPlane")
		registerFunc(&_cVPixelBufferGetDataSize, frameworkHandle, "CVPixelBufferGetDataSize")
		registerFunc(&_cVPixelBufferGetExtendedPixels, frameworkHandle, "CVPixelBufferGetExtendedPixels")
		registerFunc(&_cVPixelBufferGetHeight, frameworkHandle, "CVPixelBufferGetHeight")
		registerFunc(&_cVPixelBufferGetHeightOfPlane, frameworkHandle, "CVPixelBufferGetHeightOfPlane")
		registerFunc(&_cVPixelBufferGetIOSurface, frameworkHandle, "CVPixelBufferGetIOSurface")
		registerFunc(&_cVPixelBufferGetPixelFormatType, frameworkHandle, "CVPixelBufferGetPixelFormatType")
		registerFunc(&_cVPixelBufferGetPlaneCount, frameworkHandle, "CVPixelBufferGetPlaneCount")
		registerFunc(&_cVPixelBufferGetTypeID, frameworkHandle, "CVPixelBufferGetTypeID")
		registerFunc(&_cVPixelBufferGetWidth, frameworkHandle, "CVPixelBufferGetWidth")
		registerFunc(&_cVPixelBufferGetWidthOfPlane, frameworkHandle, "CVPixelBufferGetWidthOfPlane")
		registerFunc(&_cVPixelBufferIsCompatibleWithAttributes, frameworkHandle, "CVPixelBufferIsCompatibleWithAttributes")
		registerFunc(&_cVPixelBufferIsPlanar, frameworkHandle, "CVPixelBufferIsPlanar")
		registerFunc(&_cVPixelBufferLockBaseAddress, frameworkHandle, "CVPixelBufferLockBaseAddress")
		registerFunc(&_cVPixelBufferPoolCreate, frameworkHandle, "CVPixelBufferPoolCreate")
		registerFunc(&_cVPixelBufferPoolCreatePixelBuffer, frameworkHandle, "CVPixelBufferPoolCreatePixelBuffer")
		registerFunc(&_cVPixelBufferPoolCreatePixelBufferWithAuxAttributes, frameworkHandle, "CVPixelBufferPoolCreatePixelBufferWithAuxAttributes")
		registerFunc(&_cVPixelBufferPoolFlush, frameworkHandle, "CVPixelBufferPoolFlush")
		registerFunc(&_cVPixelBufferPoolGetAttributes, frameworkHandle, "CVPixelBufferPoolGetAttributes")
		registerFunc(&_cVPixelBufferPoolGetPixelBufferAttributes, frameworkHandle, "CVPixelBufferPoolGetPixelBufferAttributes")
		registerFunc(&_cVPixelBufferPoolGetTypeID, frameworkHandle, "CVPixelBufferPoolGetTypeID")
		registerFunc(&_cVPixelBufferPoolRelease, frameworkHandle, "CVPixelBufferPoolRelease")
		registerFunc(&_cVPixelBufferPoolRetain, frameworkHandle, "CVPixelBufferPoolRetain")
		registerFunc(&_cVPixelBufferRelease, frameworkHandle, "CVPixelBufferRelease")
		registerFunc(&_cVPixelBufferRetain, frameworkHandle, "CVPixelBufferRetain")
		registerFunc(&_cVPixelBufferUnlockBaseAddress, frameworkHandle, "CVPixelBufferUnlockBaseAddress")
		registerFunc(&_cVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes, frameworkHandle, "CVPixelFormatDescriptionArrayCreateWithAllPixelFormatTypes")
		registerFunc(&_cVPixelFormatDescriptionCreateWithPixelFormatType, frameworkHandle, "CVPixelFormatDescriptionCreateWithPixelFormatType")
		registerFunc(&_cVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType, frameworkHandle, "CVPixelFormatDescriptionRegisterDescriptionWithPixelFormatType")
		registerFunc(&_cVPixelFormatTypeCopyFourCharCodeString, frameworkHandle, "CVPixelFormatTypeCopyFourCharCodeString")
		registerFunc(&_cVTransferFunctionGetIntegerCodePointForString, frameworkHandle, "CVTransferFunctionGetIntegerCodePointForString")
		registerFunc(&_cVTransferFunctionGetStringForIntegerCodePoint, frameworkHandle, "CVTransferFunctionGetStringForIntegerCodePoint")
		registerFunc(&_cVYCbCrMatrixGetIntegerCodePointForString, frameworkHandle, "CVYCbCrMatrixGetIntegerCodePointForString")
		registerFunc(&_cVYCbCrMatrixGetStringForIntegerCodePoint, frameworkHandle, "CVYCbCrMatrixGetStringForIntegerCodePoint")
	}


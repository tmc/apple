// Code generated from Apple documentation. DO NOT EDIT.

package coregraphics

import (
	"unsafe"
	"github.com/tmc/apple/corefoundation"
)

type CGBitmapContextReleaseDataCallback = func(unsafe.Pointer, unsafe.Pointer)

type CGButtonCount = uint32

type CGCharCode = uint16

type CGColorConversionInfoRef uintptr

type CGColorRef uintptr

type CGColorSpaceRef uintptr

type CGContextRef uintptr

type CGDataConsumerPutBytesCallback = func(unsafe.Pointer, unsafe.Pointer, uint) uint

type CGDataConsumerRef uintptr

type CGDataConsumerReleaseInfoCallback = func(unsafe.Pointer)

type CGDataProviderGetBytePointerCallback = func(unsafe.Pointer) unsafe.Pointer

type CGDataProviderGetBytesAtPositionCallback = func(unsafe.Pointer, unsafe.Pointer, int64, uint) uint

type CGDataProviderGetBytesCallback = func(unsafe.Pointer, unsafe.Pointer, uint) uint

type CGDataProviderRef uintptr

type CGDataProviderReleaseBytePointerCallback = func(unsafe.Pointer, unsafe.Pointer)

type CGDataProviderReleaseDataCallback = func(unsafe.Pointer, unsafe.Pointer, uint)

type CGDataProviderReleaseInfoCallback = func(unsafe.Pointer)

type CGDataProviderRewindCallback = func(unsafe.Pointer)

type CGDataProviderSkipForwardCallback = func(unsafe.Pointer, int64) int64

type CGDirectDisplayID = uint32

type CGDisplayBlendFraction = float32

type CGDisplayConfigRef uintptr

type CGDisplayFadeInterval = float32

type CGDisplayFadeReservationToken = uint32

type CGDisplayModeRef uintptr

type CGDisplayReconfigurationCallBack = func(uint, uint, unsafe.Pointer)

type CGDisplayReservationInterval = float32

type CGDisplayStreamFrameAvailableHandler = func(CGDisplayStreamFrameStatus, uint64, IOSurfaceRef, *CGDisplayStreamUpdateRef)

type CGDisplayStreamRef uintptr

type CGDisplayStreamUpdateRef uintptr

type CGErrorCallback = func()

type CGEventErr = unsafe.Pointer

type CGEventMask = uint64

type CGEventRef uintptr

type CGEventSourceKeyboardType = uint32

type CGEventSourceRef uintptr

type CGEventTapCallBack = func(uintptr, CGEventType, uintptr, unsafe.Pointer) uintptr

type CGEventTapInformation = unsafe.Pointer

type CGEventTapProxy = uintptr

type CGEventTimestamp = uint64

type CGFontIndex = uint16

type CGFontRef uintptr

type CGFunctionEvaluateCallback = func(unsafe.Pointer, []float64, []float64)

type CGFunctionRef uintptr

type CGFunctionReleaseInfoCallback = func(unsafe.Pointer)

type CGGammaValue = float32

type CGGlyph = unsafe.Pointer

type CGGradientRef uintptr

type CGImageRef uintptr

type CGKeyCode = uint16

type CGLayerRef uintptr

type CGMutablePathRef uintptr

type CGOpenGLDisplayMask = uint32

type CGPDFArrayApplierBlock = func(uint32, *CGPDFObjectRef, unsafe.Pointer) bool

type CGPDFArrayRef uintptr

type CGPDFBoolean = uint8

type CGPDFContentStreamRef uintptr

type CGPDFDictionaryApplierBlock = func([]byte, *CGPDFObjectRef, unsafe.Pointer) bool

type CGPDFDictionaryApplierFunction = func(*byte, uintptr, unsafe.Pointer)

type CGPDFDictionaryRef uintptr

type CGPDFDocumentRef uintptr

type CGPDFInteger = int

type CGPDFObjectRef uintptr

type CGPDFOperatorCallback = func(uintptr, unsafe.Pointer)

type CGPDFOperatorTableRef uintptr

type CGPDFPageRef uintptr

type CGPDFReal = float64

type CGPDFScannerRef uintptr

type CGPDFStreamRef uintptr

type CGPDFStringRef uintptr

type CGPDFTagProperty = corefoundation.CFStringRef

type CGPSConverterBeginDocumentCallback = func(unsafe.Pointer)

type CGPSConverterBeginPageCallback = func(unsafe.Pointer, uint, uintptr)

type CGPSConverterEndDocumentCallback = func(unsafe.Pointer, bool)

type CGPSConverterEndPageCallback = func(unsafe.Pointer, uint, uintptr)

type CGPSConverterMessageCallback = func(unsafe.Pointer, uintptr)

type CGPSConverterProgressCallback = func(unsafe.Pointer)

type CGPSConverterRef uintptr

type CGPSConverterReleaseInfoCallback = func(unsafe.Pointer)

type CGPathApplierFunction = func(unsafe.Pointer, uintptr)

type CGPathApplyBlock = func(*CGPathElement)

type CGPathRef uintptr

type CGPatternDrawPatternCallback = func(unsafe.Pointer, uintptr)

type CGPatternRef uintptr

type CGPatternReleaseInfoCallback = func(unsafe.Pointer)

type CGRectCount = uint32

type CGRefreshRate = float64

type CGRenderingBufferProviderRef uintptr

type CGScreenRefreshCallback = func(uint, uintptr, unsafe.Pointer)

type CGScreenUpdateMoveCallback = func(CGScreenUpdateMoveDelta, uint, uintptr, unsafe.Pointer)

type CGShadingRef uintptr

type CGWheelCount = uint32

type CGWindowID = uint32

type CGWindowLevel = int32

type ColorSyncProfileRef uintptr

type IOSurfaceRef uintptr

// CGEvent is an opaque Core Graphics event reference type.

type CGEvent = CGEventRef

// CGFont is a bare alias for CGFontRef.

type CGFont = CGFontRef


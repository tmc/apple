// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

package coregraphics

import (
	"fmt"
	"os"
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: CoreGraphics: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: CoreGraphics: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: CoreGraphics: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _cGAcquireDisplayFadeReservation func(seconds float32, token *CGDisplayFadeReservationToken) CGError

// CGAcquireDisplayFadeReservation reserves the fade hardware for a specified time interval.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAcquireDisplayFadeReservation(_:_:)
func CGAcquireDisplayFadeReservation(seconds float32, token *CGDisplayFadeReservationToken) CGError {
	if _cGAcquireDisplayFadeReservation == nil {
		panic("CoreGraphics: symbol CGAcquireDisplayFadeReservation not loaded")
	}
	return _cGAcquireDisplayFadeReservation(seconds, token)
}

var _cGAffineTransformConcat func(t1 corefoundation.CGAffineTransform, t2 corefoundation.CGAffineTransform) corefoundation.CGAffineTransform

// CGAffineTransformConcat returns an affine transformation matrix constructed by combining two existing affine transforms.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformConcat(_:_:)
func CGAffineTransformConcat(t1 corefoundation.CGAffineTransform, t2 corefoundation.CGAffineTransform) corefoundation.CGAffineTransform {
	if _cGAffineTransformConcat == nil {
		panic("CoreGraphics: symbol CGAffineTransformConcat not loaded")
	}
	return _cGAffineTransformConcat(t1, t2)
}

var _cGAffineTransformDecompose func(transform corefoundation.CGAffineTransform) uintptr

// CGAffineTransformDecompose.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformDecompose
func CGAffineTransformDecompose(transform corefoundation.CGAffineTransform) uintptr {
	if _cGAffineTransformDecompose == nil {
		panic("CoreGraphics: symbol CGAffineTransformDecompose not loaded")
	}
	return _cGAffineTransformDecompose(transform)
}

var _cGAffineTransformEqualToTransform func(t1 corefoundation.CGAffineTransform, t2 corefoundation.CGAffineTransform) bool

// CGAffineTransformEqualToTransform checks whether two affine transforms are equal.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformEqualToTransform(_:_:)
func CGAffineTransformEqualToTransform(t1 corefoundation.CGAffineTransform, t2 corefoundation.CGAffineTransform) bool {
	if _cGAffineTransformEqualToTransform == nil {
		panic("CoreGraphics: symbol CGAffineTransformEqualToTransform not loaded")
	}
	return _cGAffineTransformEqualToTransform(t1, t2)
}

var _cGAffineTransformInvert func(t corefoundation.CGAffineTransform) corefoundation.CGAffineTransform

// CGAffineTransformInvert returns an affine transformation matrix constructed by inverting an existing affine transform.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformInvert(_:)
func CGAffineTransformInvert(t corefoundation.CGAffineTransform) corefoundation.CGAffineTransform {
	if _cGAffineTransformInvert == nil {
		panic("CoreGraphics: symbol CGAffineTransformInvert not loaded")
	}
	return _cGAffineTransformInvert(t)
}

var _cGAffineTransformIsIdentity func(t corefoundation.CGAffineTransform) bool

// CGAffineTransformIsIdentity checks whether an affine transform is the identity transform.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformIsIdentity(_:)
func CGAffineTransformIsIdentity(t corefoundation.CGAffineTransform) bool {
	if _cGAffineTransformIsIdentity == nil {
		panic("CoreGraphics: symbol CGAffineTransformIsIdentity not loaded")
	}
	return _cGAffineTransformIsIdentity(t)
}

var _cGAffineTransformMake func(a float64, b float64, c float64, d float64, tx float64, ty float64) corefoundation.CGAffineTransform

// CGAffineTransformMake returns an affine transformation matrix constructed from values you provide.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformMake(_:_:_:_:_:_:)
func CGAffineTransformMake(a float64, b float64, c float64, d float64, tx float64, ty float64) corefoundation.CGAffineTransform {
	if _cGAffineTransformMake == nil {
		panic("CoreGraphics: symbol CGAffineTransformMake not loaded")
	}
	return _cGAffineTransformMake(a, b, c, d, tx, ty)
}

var _cGAffineTransformMakeRotation func(angle float64) corefoundation.CGAffineTransform

// CGAffineTransformMakeRotation returns an affine transformation matrix constructed from a rotation value you provide.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformMakeRotation(_:)
func CGAffineTransformMakeRotation(angle float64) corefoundation.CGAffineTransform {
	if _cGAffineTransformMakeRotation == nil {
		panic("CoreGraphics: symbol CGAffineTransformMakeRotation not loaded")
	}
	return _cGAffineTransformMakeRotation(angle)
}

var _cGAffineTransformMakeScale func(sx float64, sy float64) corefoundation.CGAffineTransform

// CGAffineTransformMakeScale returns an affine transformation matrix constructed from scaling values you provide.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformMakeScale(_:_:)
func CGAffineTransformMakeScale(sx float64, sy float64) corefoundation.CGAffineTransform {
	if _cGAffineTransformMakeScale == nil {
		panic("CoreGraphics: symbol CGAffineTransformMakeScale not loaded")
	}
	return _cGAffineTransformMakeScale(sx, sy)
}

var _cGAffineTransformMakeTranslation func(tx float64, ty float64) corefoundation.CGAffineTransform

// CGAffineTransformMakeTranslation returns an affine transformation matrix constructed from translation values you provide.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformMakeTranslation(_:_:)
func CGAffineTransformMakeTranslation(tx float64, ty float64) corefoundation.CGAffineTransform {
	if _cGAffineTransformMakeTranslation == nil {
		panic("CoreGraphics: symbol CGAffineTransformMakeTranslation not loaded")
	}
	return _cGAffineTransformMakeTranslation(tx, ty)
}

var _cGAffineTransformMakeWithComponents func(components uintptr) corefoundation.CGAffineTransform

// CGAffineTransformMakeWithComponents.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformMakeWithComponents
func CGAffineTransformMakeWithComponents(components uintptr) corefoundation.CGAffineTransform {
	if _cGAffineTransformMakeWithComponents == nil {
		panic("CoreGraphics: symbol CGAffineTransformMakeWithComponents not loaded")
	}
	return _cGAffineTransformMakeWithComponents(components)
}

var _cGAffineTransformRotate func(t corefoundation.CGAffineTransform, angle float64) corefoundation.CGAffineTransform

// CGAffineTransformRotate returns an affine transformation matrix constructed by rotating an existing affine transform.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformRotate(_:_:)
func CGAffineTransformRotate(t corefoundation.CGAffineTransform, angle float64) corefoundation.CGAffineTransform {
	if _cGAffineTransformRotate == nil {
		panic("CoreGraphics: symbol CGAffineTransformRotate not loaded")
	}
	return _cGAffineTransformRotate(t, angle)
}

var _cGAffineTransformScale func(t corefoundation.CGAffineTransform, sx float64, sy float64) corefoundation.CGAffineTransform

// CGAffineTransformScale returns an affine transformation matrix constructed by scaling an existing affine transform.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformScale(_:_:_:)
func CGAffineTransformScale(t corefoundation.CGAffineTransform, sx float64, sy float64) corefoundation.CGAffineTransform {
	if _cGAffineTransformScale == nil {
		panic("CoreGraphics: symbol CGAffineTransformScale not loaded")
	}
	return _cGAffineTransformScale(t, sx, sy)
}

var _cGAffineTransformTranslate func(t corefoundation.CGAffineTransform, tx float64, ty float64) corefoundation.CGAffineTransform

// CGAffineTransformTranslate returns an affine transformation matrix constructed by translating an existing affine transform.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformTranslate(_:_:_:)
func CGAffineTransformTranslate(t corefoundation.CGAffineTransform, tx float64, ty float64) corefoundation.CGAffineTransform {
	if _cGAffineTransformTranslate == nil {
		panic("CoreGraphics: symbol CGAffineTransformTranslate not loaded")
	}
	return _cGAffineTransformTranslate(t, tx, ty)
}

var _cGAssociateMouseAndMouseCursorPosition func(connected uintptr) CGError

// CGAssociateMouseAndMouseCursorPosition connects or disconnects the mouse and cursor while an application is in the foreground.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAssociateMouseAndMouseCursorPosition(_:)
func CGAssociateMouseAndMouseCursorPosition(connected uintptr) CGError {
	if _cGAssociateMouseAndMouseCursorPosition == nil {
		panic("CoreGraphics: symbol CGAssociateMouseAndMouseCursorPosition not loaded")
	}
	return _cGAssociateMouseAndMouseCursorPosition(connected)
}

var _cGBeginDisplayConfiguration func(config *CGDisplayConfigRef) CGError

// CGBeginDisplayConfiguration begins a new set of display configuration changes.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGBeginDisplayConfiguration(_:)
func CGBeginDisplayConfiguration(config *CGDisplayConfigRef) CGError {
	if _cGBeginDisplayConfiguration == nil {
		panic("CoreGraphics: symbol CGBeginDisplayConfiguration not loaded")
	}
	return _cGBeginDisplayConfiguration(config)
}

var _cGBitmapContextCreate func(data unsafe.Pointer, width uintptr, height uintptr, bitsPerComponent uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo) CGContextRef

// CGBitmapContextCreate.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/init(data:width:height:bitsPerComponent:bytesPerRow:space:bitmapInfo:)-10b3i
func CGBitmapContextCreate(data unsafe.Pointer, width uintptr, height uintptr, bitsPerComponent uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo) CGContextRef {
	if _cGBitmapContextCreate == nil {
		panic("CoreGraphics: symbol CGBitmapContextCreate not loaded")
	}
	return _cGBitmapContextCreate(data, width, height, bitsPerComponent, bytesPerRow, space, bitmapInfo)
}

var _cGBitmapContextCreateAdaptive func(width uintptr, height uintptr, auxiliaryInfo corefoundation.CFDictionaryRef, onResolve bool) CGContextRef

// CGBitmapContextCreateAdaptive.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGBitmapContextCreateAdaptive
func CGBitmapContextCreateAdaptive(width uintptr, height uintptr, auxiliaryInfo corefoundation.CFDictionaryRef, onResolve bool) CGContextRef {
	if _cGBitmapContextCreateAdaptive == nil {
		panic("CoreGraphics: symbol CGBitmapContextCreateAdaptive not loaded")
	}
	return _cGBitmapContextCreateAdaptive(width, height, auxiliaryInfo, onResolve)
}

var _cGBitmapContextCreateImage func(context CGContextRef) CGImageRef

// CGBitmapContextCreateImage creates and returns a CGImage from the pixel data in a bitmap graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/makeImage()
func CGBitmapContextCreateImage(context CGContextRef) CGImageRef {
	if _cGBitmapContextCreateImage == nil {
		panic("CoreGraphics: symbol CGBitmapContextCreateImage not loaded")
	}
	return _cGBitmapContextCreateImage(context)
}

var _cGBitmapContextCreateWithData func(data unsafe.Pointer, width uintptr, height uintptr, bitsPerComponent uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo, releaseCallback CGBitmapContextReleaseDataCallback, releaseInfo unsafe.Pointer) CGContextRef

// CGBitmapContextCreateWithData.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/init(data:width:height:bitsPerComponent:bytesPerRow:space:bitmapInfo:releaseCallback:releaseInfo:)-4yzt5
func CGBitmapContextCreateWithData(data unsafe.Pointer, width uintptr, height uintptr, bitsPerComponent uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo, releaseCallback CGBitmapContextReleaseDataCallback, releaseInfo unsafe.Pointer) CGContextRef {
	if _cGBitmapContextCreateWithData == nil {
		panic("CoreGraphics: symbol CGBitmapContextCreateWithData not loaded")
	}
	return _cGBitmapContextCreateWithData(data, width, height, bitsPerComponent, bytesPerRow, space, bitmapInfo, releaseCallback, releaseInfo)
}

var _cGBitmapContextGetAlphaInfo func(context CGContextRef) CGImageAlphaInfo

// CGBitmapContextGetAlphaInfo returns the alpha information associated with the context, which indicates how a bitmap context handles the alpha component.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/alphaInfo
func CGBitmapContextGetAlphaInfo(context CGContextRef) CGImageAlphaInfo {
	if _cGBitmapContextGetAlphaInfo == nil {
		panic("CoreGraphics: symbol CGBitmapContextGetAlphaInfo not loaded")
	}
	return _cGBitmapContextGetAlphaInfo(context)
}

var _cGBitmapContextGetBitmapInfo func(context CGContextRef) CGBitmapInfo

// CGBitmapContextGetBitmapInfo obtains the bitmap information associated with a bitmap graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/bitmapInfo
func CGBitmapContextGetBitmapInfo(context CGContextRef) CGBitmapInfo {
	if _cGBitmapContextGetBitmapInfo == nil {
		panic("CoreGraphics: symbol CGBitmapContextGetBitmapInfo not loaded")
	}
	return _cGBitmapContextGetBitmapInfo(context)
}

var _cGBitmapContextGetBitsPerComponent func(context CGContextRef) uintptr

// CGBitmapContextGetBitsPerComponent returns the bits per component of a bitmap context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/bitsPerComponent
func CGBitmapContextGetBitsPerComponent(context CGContextRef) uintptr {
	if _cGBitmapContextGetBitsPerComponent == nil {
		panic("CoreGraphics: symbol CGBitmapContextGetBitsPerComponent not loaded")
	}
	return _cGBitmapContextGetBitsPerComponent(context)
}

var _cGBitmapContextGetBitsPerPixel func(context CGContextRef) uintptr

// CGBitmapContextGetBitsPerPixel returns the bits per pixel of a bitmap context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/bitsPerPixel
func CGBitmapContextGetBitsPerPixel(context CGContextRef) uintptr {
	if _cGBitmapContextGetBitsPerPixel == nil {
		panic("CoreGraphics: symbol CGBitmapContextGetBitsPerPixel not loaded")
	}
	return _cGBitmapContextGetBitsPerPixel(context)
}

var _cGBitmapContextGetBytesPerRow func(context CGContextRef) uintptr

// CGBitmapContextGetBytesPerRow returns the bytes per row of a bitmap context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/bytesPerRow
func CGBitmapContextGetBytesPerRow(context CGContextRef) uintptr {
	if _cGBitmapContextGetBytesPerRow == nil {
		panic("CoreGraphics: symbol CGBitmapContextGetBytesPerRow not loaded")
	}
	return _cGBitmapContextGetBytesPerRow(context)
}

var _cGBitmapContextGetColorSpace func(context CGContextRef) CGColorSpaceRef

// CGBitmapContextGetColorSpace returns the color space of a bitmap context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/colorSpace
func CGBitmapContextGetColorSpace(context CGContextRef) CGColorSpaceRef {
	if _cGBitmapContextGetColorSpace == nil {
		panic("CoreGraphics: symbol CGBitmapContextGetColorSpace not loaded")
	}
	return _cGBitmapContextGetColorSpace(context)
}

var _cGBitmapContextGetData func(context CGContextRef) unsafe.Pointer

// CGBitmapContextGetData returns a pointer to the image data associated with a bitmap context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/data
func CGBitmapContextGetData(context CGContextRef) unsafe.Pointer {
	if _cGBitmapContextGetData == nil {
		panic("CoreGraphics: symbol CGBitmapContextGetData not loaded")
	}
	return _cGBitmapContextGetData(context)
}

var _cGBitmapContextGetHeight func(context CGContextRef) uintptr

// CGBitmapContextGetHeight returns the height in pixels of a bitmap context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/height
func CGBitmapContextGetHeight(context CGContextRef) uintptr {
	if _cGBitmapContextGetHeight == nil {
		panic("CoreGraphics: symbol CGBitmapContextGetHeight not loaded")
	}
	return _cGBitmapContextGetHeight(context)
}

var _cGBitmapContextGetWidth func(context CGContextRef) uintptr

// CGBitmapContextGetWidth returns the width in pixels of a bitmap context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/width
func CGBitmapContextGetWidth(context CGContextRef) uintptr {
	if _cGBitmapContextGetWidth == nil {
		panic("CoreGraphics: symbol CGBitmapContextGetWidth not loaded")
	}
	return _cGBitmapContextGetWidth(context)
}

var _cGCancelDisplayConfiguration func(config CGDisplayConfigRef) CGError

// CGCancelDisplayConfiguration cancels a set of display configuration changes.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGCancelDisplayConfiguration(_:)
func CGCancelDisplayConfiguration(config CGDisplayConfigRef) CGError {
	if _cGCancelDisplayConfiguration == nil {
		panic("CoreGraphics: symbol CGCancelDisplayConfiguration not loaded")
	}
	return _cGCancelDisplayConfiguration(config)
}

var _cGCaptureAllDisplays func() CGError

// CGCaptureAllDisplays obtains exclusive use of all active displays, preventing other applications and system services from using the display or changing its configuration.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGCaptureAllDisplays()
func CGCaptureAllDisplays() CGError {
	if _cGCaptureAllDisplays == nil {
		panic("CoreGraphics: symbol CGCaptureAllDisplays not loaded")
	}
	return _cGCaptureAllDisplays()
}

var _cGCaptureAllDisplaysWithOptions func(options CGCaptureOptions) CGError

// CGCaptureAllDisplaysWithOptions captures all attached displays, using the specified options.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGCaptureAllDisplaysWithOptions(_:)
func CGCaptureAllDisplaysWithOptions(options CGCaptureOptions) CGError {
	if _cGCaptureAllDisplaysWithOptions == nil {
		panic("CoreGraphics: symbol CGCaptureAllDisplaysWithOptions not loaded")
	}
	return _cGCaptureAllDisplaysWithOptions(options)
}

var _cGColorConversionInfoConvertData func(info CGColorConversionInfoRef, width uintptr, height uintptr, dst_data unsafe.Pointer, dst_format CGColorBufferFormat, src_data unsafe.Pointer, src_format CGColorBufferFormat, options corefoundation.CFDictionaryRef) bool

// CGColorConversionInfoConvertData.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfo/convert(width:height:to:format:from:format:options:)
func CGColorConversionInfoConvertData(info CGColorConversionInfoRef, width uintptr, height uintptr, dst_data unsafe.Pointer, dst_format CGColorBufferFormat, src_data unsafe.Pointer, src_format CGColorBufferFormat, options corefoundation.CFDictionaryRef) bool {
	if _cGColorConversionInfoConvertData == nil {
		panic("CoreGraphics: symbol CGColorConversionInfoConvertData not loaded")
	}
	return _cGColorConversionInfoConvertData(info, width, height, dst_data, dst_format, src_data, src_format, options)
}

var _cGColorConversionInfoCreate func(src CGColorSpaceRef, dst CGColorSpaceRef) CGColorConversionInfoRef

// CGColorConversionInfoCreate creates a conversion between two specified color spaces.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfo/init(src:dst:)
func CGColorConversionInfoCreate(src CGColorSpaceRef, dst CGColorSpaceRef) CGColorConversionInfoRef {
	if _cGColorConversionInfoCreate == nil {
		panic("CoreGraphics: symbol CGColorConversionInfoCreate not loaded")
	}
	return _cGColorConversionInfoCreate(src, dst)
}

var _cGColorConversionInfoCreateForToneMapping func(from CGColorSpaceRef, source_headroom float32, to CGColorSpaceRef, target_headroom float32, method CGToneMapping, options corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) CGColorConversionInfoRef

// CGColorConversionInfoCreateForToneMapping.
//
// Deprecated: declared Swift name 'init(src:srcHeadroom:dst:dstHeadroom:toneMapping:options:)' was adjusted to 'init(src:srcHeadroom:dst:dstHeadroom:toneMapping:options:_:)' because it does not have the correct number of parameters (6 vs. 7); please report this to its maintainer
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfo/init(src:srcHeadroom:dst:dstHeadroom:toneMapping:options:_:)
func CGColorConversionInfoCreateForToneMapping(from CGColorSpaceRef, source_headroom float32, to CGColorSpaceRef, target_headroom float32, method CGToneMapping, options corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) CGColorConversionInfoRef {
	if _cGColorConversionInfoCreateForToneMapping == nil {
		panic("CoreGraphics: symbol CGColorConversionInfoCreateForToneMapping not loaded")
	}
	return _cGColorConversionInfoCreateForToneMapping(from, source_headroom, to, target_headroom, method, options, err)
}

var _cGColorConversionInfoCreateFromList func(options corefoundation.CFDictionaryRef, arg1 CGColorSpaceRef, arg2 CGColorConversionInfoTransformType, arg3 CGColorRenderingIntent) CGColorConversionInfoRef

// CGColorConversionInfoCreateFromList creates a conversion between an arbitrary number of specified color spaces.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfoCreateFromList
func CGColorConversionInfoCreateFromList(options corefoundation.CFDictionaryRef, arg1 CGColorSpaceRef, arg2 CGColorConversionInfoTransformType, arg3 CGColorRenderingIntent) CGColorConversionInfoRef {
	if _cGColorConversionInfoCreateFromList == nil {
		panic("CoreGraphics: symbol CGColorConversionInfoCreateFromList not loaded")
	}
	return _cGColorConversionInfoCreateFromList(options, arg1, arg2, arg3)
}

var _cGColorConversionInfoCreateFromListWithArguments func(options corefoundation.CFDictionaryRef, arg1 CGColorSpaceRef, arg2 CGColorConversionInfoTransformType, arg3 CGColorRenderingIntent, arg4 uintptr) CGColorConversionInfoRef

// CGColorConversionInfoCreateFromListWithArguments.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfoCreateFromListWithArguments
func CGColorConversionInfoCreateFromListWithArguments(options corefoundation.CFDictionaryRef, arg1 CGColorSpaceRef, arg2 CGColorConversionInfoTransformType, arg3 CGColorRenderingIntent, arg4 uintptr) CGColorConversionInfoRef {
	if _cGColorConversionInfoCreateFromListWithArguments == nil {
		panic("CoreGraphics: symbol CGColorConversionInfoCreateFromListWithArguments not loaded")
	}
	return _cGColorConversionInfoCreateFromListWithArguments(options, arg1, arg2, arg3, arg4)
}

var _cGColorConversionInfoCreateWithOptions func(src CGColorSpaceRef, dst CGColorSpaceRef, options corefoundation.CFDictionaryRef) CGColorConversionInfoRef

// CGColorConversionInfoCreateWithOptions.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfo/init(optionsSrc:dst:options:)
func CGColorConversionInfoCreateWithOptions(src CGColorSpaceRef, dst CGColorSpaceRef, options corefoundation.CFDictionaryRef) CGColorConversionInfoRef {
	if _cGColorConversionInfoCreateWithOptions == nil {
		panic("CoreGraphics: symbol CGColorConversionInfoCreateWithOptions not loaded")
	}
	return _cGColorConversionInfoCreateWithOptions(src, dst, options)
}

var _cGColorConversionInfoGetTypeID func() uint

// CGColorConversionInfoGetTypeID returns the Core Foundation type identifier for a color conversion info data type.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfo/typeID
func CGColorConversionInfoGetTypeID() uint {
	if _cGColorConversionInfoGetTypeID == nil {
		panic("CoreGraphics: symbol CGColorConversionInfoGetTypeID not loaded")
	}
	return _cGColorConversionInfoGetTypeID()
}

var _cGColorCreate func(space CGColorSpaceRef, components *float64) CGColorRef

// CGColorCreate creates a color using a list of intensity values (including alpha) and an associated color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(colorSpace:components:)
func CGColorCreate(space CGColorSpaceRef, components *float64) CGColorRef {
	if _cGColorCreate == nil {
		panic("CoreGraphics: symbol CGColorCreate not loaded")
	}
	return _cGColorCreate(space, components)
}

var _cGColorCreateCopy func(color CGColorRef) CGColorRef

// CGColorCreateCopy creates a copy of an existing color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/copy()
func CGColorCreateCopy(color CGColorRef) CGColorRef {
	if _cGColorCreateCopy == nil {
		panic("CoreGraphics: symbol CGColorCreateCopy not loaded")
	}
	return _cGColorCreateCopy(color)
}

var _cGColorCreateCopyByMatchingToColorSpace func(arg0 CGColorSpaceRef, intent CGColorRenderingIntent, color CGColorRef, options corefoundation.CFDictionaryRef) CGColorRef

// CGColorCreateCopyByMatchingToColorSpace creates a new color in a different color space that matches the provided color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/converted(to:intent:options:)
func CGColorCreateCopyByMatchingToColorSpace(arg0 CGColorSpaceRef, intent CGColorRenderingIntent, color CGColorRef, options corefoundation.CFDictionaryRef) CGColorRef {
	if _cGColorCreateCopyByMatchingToColorSpace == nil {
		panic("CoreGraphics: symbol CGColorCreateCopyByMatchingToColorSpace not loaded")
	}
	return _cGColorCreateCopyByMatchingToColorSpace(arg0, intent, color, options)
}

var _cGColorCreateCopyWithAlpha func(color CGColorRef, alpha float64) CGColorRef

// CGColorCreateCopyWithAlpha creates a copy of an existing color, substituting a new alpha value.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/copy(alpha:)
func CGColorCreateCopyWithAlpha(color CGColorRef, alpha float64) CGColorRef {
	if _cGColorCreateCopyWithAlpha == nil {
		panic("CoreGraphics: symbol CGColorCreateCopyWithAlpha not loaded")
	}
	return _cGColorCreateCopyWithAlpha(color, alpha)
}

var _cGColorCreateGenericCMYK func(cyan float64, magenta float64, yellow float64, black float64, alpha float64) CGColorRef

// CGColorCreateGenericCMYK creates a color in the Generic CMYK color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(genericCMYKCyan:magenta:yellow:black:alpha:)
func CGColorCreateGenericCMYK(cyan float64, magenta float64, yellow float64, black float64, alpha float64) CGColorRef {
	if _cGColorCreateGenericCMYK == nil {
		panic("CoreGraphics: symbol CGColorCreateGenericCMYK not loaded")
	}
	return _cGColorCreateGenericCMYK(cyan, magenta, yellow, black, alpha)
}

var _cGColorCreateGenericGray func(gray float64, alpha float64) CGColorRef

// CGColorCreateGenericGray creates a color in the Generic gray color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(gray:alpha:)
func CGColorCreateGenericGray(gray float64, alpha float64) CGColorRef {
	if _cGColorCreateGenericGray == nil {
		panic("CoreGraphics: symbol CGColorCreateGenericGray not loaded")
	}
	return _cGColorCreateGenericGray(gray, alpha)
}

var _cGColorCreateGenericGrayGamma2_2 func(gray float64, alpha float64) CGColorRef

// CGColorCreateGenericGrayGamma2_2 creates a color in the Generic gray color space with a gamma ramp of 2.2.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(genericGrayGamma2_2Gray:alpha:)
func CGColorCreateGenericGrayGamma2_2(gray float64, alpha float64) CGColorRef {
	if _cGColorCreateGenericGrayGamma2_2 == nil {
		panic("CoreGraphics: symbol CGColorCreateGenericGrayGamma2_2 not loaded")
	}
	return _cGColorCreateGenericGrayGamma2_2(gray, alpha)
}

var _cGColorCreateGenericRGB func(red float64, green float64, blue float64, alpha float64) CGColorRef

// CGColorCreateGenericRGB creates a color in the Generic RGB color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(red:green:blue:alpha:)
func CGColorCreateGenericRGB(red float64, green float64, blue float64, alpha float64) CGColorRef {
	if _cGColorCreateGenericRGB == nil {
		panic("CoreGraphics: symbol CGColorCreateGenericRGB not loaded")
	}
	return _cGColorCreateGenericRGB(red, green, blue, alpha)
}

var _cGColorCreateSRGB func(red float64, green float64, blue float64, alpha float64) CGColorRef

// CGColorCreateSRGB creates a color in the sRGB color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(srgbRed:green:blue:alpha:)
func CGColorCreateSRGB(red float64, green float64, blue float64, alpha float64) CGColorRef {
	if _cGColorCreateSRGB == nil {
		panic("CoreGraphics: symbol CGColorCreateSRGB not loaded")
	}
	return _cGColorCreateSRGB(red, green, blue, alpha)
}

var _cGColorCreateWithContentHeadroom func(headroom float32, space CGColorSpaceRef, red float64, green float64, blue float64, alpha float64) CGColorRef

// CGColorCreateWithContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(headroom:colorSpace:red:green:blue:alpha:)
func CGColorCreateWithContentHeadroom(headroom float32, space CGColorSpaceRef, red float64, green float64, blue float64, alpha float64) CGColorRef {
	if _cGColorCreateWithContentHeadroom == nil {
		panic("CoreGraphics: symbol CGColorCreateWithContentHeadroom not loaded")
	}
	return _cGColorCreateWithContentHeadroom(headroom, space, red, green, blue, alpha)
}

var _cGColorCreateWithPattern func(space CGColorSpaceRef, pattern CGPatternRef, components *float64) CGColorRef

// CGColorCreateWithPattern creates a color using a list of intensity values (including alpha), a pattern color space, and a pattern.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(patternSpace:pattern:components:)
func CGColorCreateWithPattern(space CGColorSpaceRef, pattern CGPatternRef, components *float64) CGColorRef {
	if _cGColorCreateWithPattern == nil {
		panic("CoreGraphics: symbol CGColorCreateWithPattern not loaded")
	}
	return _cGColorCreateWithPattern(space, pattern, components)
}

var _cGColorEqualToColor func(color1 CGColorRef, color2 CGColorRef) bool

// CGColorEqualToColor indicates whether two colors are equal.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorEqualToColor
func CGColorEqualToColor(color1 CGColorRef, color2 CGColorRef) bool {
	if _cGColorEqualToColor == nil {
		panic("CoreGraphics: symbol CGColorEqualToColor not loaded")
	}
	return _cGColorEqualToColor(color1, color2)
}

var _cGColorGetAlpha func(color CGColorRef) float64

// CGColorGetAlpha returns the value of the alpha component associated with a color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/alpha
func CGColorGetAlpha(color CGColorRef) float64 {
	if _cGColorGetAlpha == nil {
		panic("CoreGraphics: symbol CGColorGetAlpha not loaded")
	}
	return _cGColorGetAlpha(color)
}

var _cGColorGetColorSpace func(color CGColorRef) CGColorSpaceRef

// CGColorGetColorSpace returns the color space associated with a color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/colorSpace
func CGColorGetColorSpace(color CGColorRef) CGColorSpaceRef {
	if _cGColorGetColorSpace == nil {
		panic("CoreGraphics: symbol CGColorGetColorSpace not loaded")
	}
	return _cGColorGetColorSpace(color)
}

var _cGColorGetComponents func(color CGColorRef) *float64

// CGColorGetComponents returns the values of the color components (including alpha) associated with a color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorGetComponents
func CGColorGetComponents(color CGColorRef) *float64 {
	if _cGColorGetComponents == nil {
		panic("CoreGraphics: symbol CGColorGetComponents not loaded")
	}
	return _cGColorGetComponents(color)
}

var _cGColorGetConstantColor func(colorName corefoundation.CFStringRef) CGColorRef

// CGColorGetConstantColor returns a color object that represents a constant color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorGetConstantColor
func CGColorGetConstantColor(colorName corefoundation.CFStringRef) CGColorRef {
	if _cGColorGetConstantColor == nil {
		panic("CoreGraphics: symbol CGColorGetConstantColor not loaded")
	}
	return _cGColorGetConstantColor(colorName)
}

var _cGColorGetContentHeadroom func(color CGColorRef) float32

// CGColorGetContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/contentHeadroom
func CGColorGetContentHeadroom(color CGColorRef) float32 {
	if _cGColorGetContentHeadroom == nil {
		panic("CoreGraphics: symbol CGColorGetContentHeadroom not loaded")
	}
	return _cGColorGetContentHeadroom(color)
}

var _cGColorGetNumberOfComponents func(color CGColorRef) uintptr

// CGColorGetNumberOfComponents returns the number of color components (including alpha) associated with a color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/numberOfComponents
func CGColorGetNumberOfComponents(color CGColorRef) uintptr {
	if _cGColorGetNumberOfComponents == nil {
		panic("CoreGraphics: symbol CGColorGetNumberOfComponents not loaded")
	}
	return _cGColorGetNumberOfComponents(color)
}

var _cGColorGetPattern func(color CGColorRef) CGPatternRef

// CGColorGetPattern returns the pattern associated with a color in a pattern color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/pattern
func CGColorGetPattern(color CGColorRef) CGPatternRef {
	if _cGColorGetPattern == nil {
		panic("CoreGraphics: symbol CGColorGetPattern not loaded")
	}
	return _cGColorGetPattern(color)
}

var _cGColorGetTypeID func() uint

// CGColorGetTypeID returns the Core Foundation type identifier for a color data type.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/typeID
func CGColorGetTypeID() uint {
	if _cGColorGetTypeID == nil {
		panic("CoreGraphics: symbol CGColorGetTypeID not loaded")
	}
	return _cGColorGetTypeID()
}

var _cGColorRelease func(color CGColorRef)

// CGColorRelease decrements the retain count of a color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorRelease
func CGColorRelease(color CGColorRef) {
	if _cGColorRelease == nil {
		panic("CoreGraphics: symbol CGColorRelease not loaded")
	}
	_cGColorRelease(color)
}

var _cGColorRetain func(color CGColorRef) CGColorRef

// CGColorRetain increments the retain count of a color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorRetain
func CGColorRetain(color CGColorRef) CGColorRef {
	if _cGColorRetain == nil {
		panic("CoreGraphics: symbol CGColorRetain not loaded")
	}
	return _cGColorRetain(color)
}

var _cGColorSpaceCopyBaseColorSpace func(space CGColorSpaceRef) CGColorSpaceRef

// CGColorSpaceCopyBaseColorSpace.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCopyBaseColorSpace(_:)
func CGColorSpaceCopyBaseColorSpace(space CGColorSpaceRef) CGColorSpaceRef {
	if _cGColorSpaceCopyBaseColorSpace == nil {
		panic("CoreGraphics: symbol CGColorSpaceCopyBaseColorSpace not loaded")
	}
	return _cGColorSpaceCopyBaseColorSpace(space)
}

var _cGColorSpaceCopyICCData func(space CGColorSpaceRef) corefoundation.CFDataRef

// CGColorSpaceCopyICCData returns a copy of the ICC profile data of the provided color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/copyICCData()
func CGColorSpaceCopyICCData(space CGColorSpaceRef) corefoundation.CFDataRef {
	if _cGColorSpaceCopyICCData == nil {
		panic("CoreGraphics: symbol CGColorSpaceCopyICCData not loaded")
	}
	return _cGColorSpaceCopyICCData(space)
}

var _cGColorSpaceCopyName func(space CGColorSpaceRef) corefoundation.CFStringRef

// CGColorSpaceCopyName returns the name used to create the specified color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/name
func CGColorSpaceCopyName(space CGColorSpaceRef) corefoundation.CFStringRef {
	if _cGColorSpaceCopyName == nil {
		panic("CoreGraphics: symbol CGColorSpaceCopyName not loaded")
	}
	return _cGColorSpaceCopyName(space)
}

var _cGColorSpaceCopyPropertyList func(space CGColorSpaceRef) corefoundation.CFPropertyListRef

// CGColorSpaceCopyPropertyList returns a copy of the color space’s properties.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/copyPropertyList()
func CGColorSpaceCopyPropertyList(space CGColorSpaceRef) corefoundation.CFPropertyListRef {
	if _cGColorSpaceCopyPropertyList == nil {
		panic("CoreGraphics: symbol CGColorSpaceCopyPropertyList not loaded")
	}
	return _cGColorSpaceCopyPropertyList(space)
}

var _cGColorSpaceCreateCalibratedGray func(whitePoint float64, blackPoint float64, gamma float64) CGColorSpaceRef

// CGColorSpaceCreateCalibratedGray creates a calibrated grayscale color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(calibratedGrayWhitePoint:blackPoint:gamma:)
func CGColorSpaceCreateCalibratedGray(whitePoint float64, blackPoint float64, gamma float64) CGColorSpaceRef {
	if _cGColorSpaceCreateCalibratedGray == nil {
		panic("CoreGraphics: symbol CGColorSpaceCreateCalibratedGray not loaded")
	}
	return _cGColorSpaceCreateCalibratedGray(whitePoint, blackPoint, gamma)
}

var _cGColorSpaceCreateCalibratedRGB func(whitePoint float64, blackPoint float64, gamma float64, matrix float64) CGColorSpaceRef

// CGColorSpaceCreateCalibratedRGB creates a calibrated RGB color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(calibratedRGBWhitePoint:blackPoint:gamma:matrix:)
func CGColorSpaceCreateCalibratedRGB(whitePoint float64, blackPoint float64, gamma float64, matrix float64) CGColorSpaceRef {
	if _cGColorSpaceCreateCalibratedRGB == nil {
		panic("CoreGraphics: symbol CGColorSpaceCreateCalibratedRGB not loaded")
	}
	return _cGColorSpaceCreateCalibratedRGB(whitePoint, blackPoint, gamma, matrix)
}

var _cGColorSpaceCreateCopyWithStandardRange func(space CGColorSpaceRef) CGColorSpaceRef

// CGColorSpaceCreateCopyWithStandardRange.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateCopyWithStandardRange(_:)
func CGColorSpaceCreateCopyWithStandardRange(space CGColorSpaceRef) CGColorSpaceRef {
	if _cGColorSpaceCreateCopyWithStandardRange == nil {
		panic("CoreGraphics: symbol CGColorSpaceCreateCopyWithStandardRange not loaded")
	}
	return _cGColorSpaceCreateCopyWithStandardRange(space)
}

var _cGColorSpaceCreateDeviceCMYK func() CGColorSpaceRef

// CGColorSpaceCreateDeviceCMYK creates a device-dependent CMYK color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateDeviceCMYK()
func CGColorSpaceCreateDeviceCMYK() CGColorSpaceRef {
	if _cGColorSpaceCreateDeviceCMYK == nil {
		panic("CoreGraphics: symbol CGColorSpaceCreateDeviceCMYK not loaded")
	}
	return _cGColorSpaceCreateDeviceCMYK()
}

var _cGColorSpaceCreateDeviceGray func() CGColorSpaceRef

// CGColorSpaceCreateDeviceGray creates a device-dependent grayscale color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateDeviceGray()
func CGColorSpaceCreateDeviceGray() CGColorSpaceRef {
	if _cGColorSpaceCreateDeviceGray == nil {
		panic("CoreGraphics: symbol CGColorSpaceCreateDeviceGray not loaded")
	}
	return _cGColorSpaceCreateDeviceGray()
}

var _cGColorSpaceCreateDeviceRGB func() CGColorSpaceRef

// CGColorSpaceCreateDeviceRGB creates a device-dependent RGB color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateDeviceRGB()
func CGColorSpaceCreateDeviceRGB() CGColorSpaceRef {
	if _cGColorSpaceCreateDeviceRGB == nil {
		panic("CoreGraphics: symbol CGColorSpaceCreateDeviceRGB not loaded")
	}
	return _cGColorSpaceCreateDeviceRGB()
}

var _cGColorSpaceCreateExtended func(space CGColorSpaceRef) CGColorSpaceRef

// CGColorSpaceCreateExtended.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateExtended(_:)
func CGColorSpaceCreateExtended(space CGColorSpaceRef) CGColorSpaceRef {
	if _cGColorSpaceCreateExtended == nil {
		panic("CoreGraphics: symbol CGColorSpaceCreateExtended not loaded")
	}
	return _cGColorSpaceCreateExtended(space)
}

var _cGColorSpaceCreateExtendedLinearized func(space CGColorSpaceRef) CGColorSpaceRef

// CGColorSpaceCreateExtendedLinearized.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateExtendedLinearized(_:)
func CGColorSpaceCreateExtendedLinearized(space CGColorSpaceRef) CGColorSpaceRef {
	if _cGColorSpaceCreateExtendedLinearized == nil {
		panic("CoreGraphics: symbol CGColorSpaceCreateExtendedLinearized not loaded")
	}
	return _cGColorSpaceCreateExtendedLinearized(space)
}

var _cGColorSpaceCreateICCBased func(nComponents uintptr, range_ *float64, profile CGDataProviderRef, alternate CGColorSpaceRef) CGColorSpaceRef

// CGColorSpaceCreateICCBased creates a device-independent color space that is defined according to the ICC color profile specification.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(iccBasedNComponents:range:profile:alternate:)
func CGColorSpaceCreateICCBased(nComponents uintptr, range_ *float64, profile CGDataProviderRef, alternate CGColorSpaceRef) CGColorSpaceRef {
	if _cGColorSpaceCreateICCBased == nil {
		panic("CoreGraphics: symbol CGColorSpaceCreateICCBased not loaded")
	}
	return _cGColorSpaceCreateICCBased(nComponents, range_, profile, alternate)
}

var _cGColorSpaceCreateIndexed func(baseSpace CGColorSpaceRef, lastIndex uintptr, colorTable string) CGColorSpaceRef

// CGColorSpaceCreateIndexed creates an indexed color space, consisting of colors specified by a color lookup table.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(indexedBaseSpace:last:colorTable:)
func CGColorSpaceCreateIndexed(baseSpace CGColorSpaceRef, lastIndex uintptr, colorTable string) CGColorSpaceRef {
	if _cGColorSpaceCreateIndexed == nil {
		panic("CoreGraphics: symbol CGColorSpaceCreateIndexed not loaded")
	}
	return _cGColorSpaceCreateIndexed(baseSpace, lastIndex, colorTable)
}

var _cGColorSpaceCreateLab func(whitePoint float64, blackPoint float64, range_ float64) CGColorSpaceRef

// CGColorSpaceCreateLab creates a device-independent color space that is relative to human color perception, according to the CIE L*a*b* standard.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(labWhitePoint:blackPoint:range:)
func CGColorSpaceCreateLab(whitePoint float64, blackPoint float64, range_ float64) CGColorSpaceRef {
	if _cGColorSpaceCreateLab == nil {
		panic("CoreGraphics: symbol CGColorSpaceCreateLab not loaded")
	}
	return _cGColorSpaceCreateLab(whitePoint, blackPoint, range_)
}

var _cGColorSpaceCreateLinearized func(space CGColorSpaceRef) CGColorSpaceRef

// CGColorSpaceCreateLinearized.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateLinearized(_:)
func CGColorSpaceCreateLinearized(space CGColorSpaceRef) CGColorSpaceRef {
	if _cGColorSpaceCreateLinearized == nil {
		panic("CoreGraphics: symbol CGColorSpaceCreateLinearized not loaded")
	}
	return _cGColorSpaceCreateLinearized(space)
}

var _cGColorSpaceCreatePattern func(baseSpace CGColorSpaceRef) CGColorSpaceRef

// CGColorSpaceCreatePattern creates a pattern color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(patternBaseSpace:)
func CGColorSpaceCreatePattern(baseSpace CGColorSpaceRef) CGColorSpaceRef {
	if _cGColorSpaceCreatePattern == nil {
		panic("CoreGraphics: symbol CGColorSpaceCreatePattern not loaded")
	}
	return _cGColorSpaceCreatePattern(baseSpace)
}

var _cGColorSpaceCreateWithColorSyncProfile func(arg0 ColorSyncProfileRef, options corefoundation.CFDictionaryRef) CGColorSpaceRef

// CGColorSpaceCreateWithColorSyncProfile.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateWithColorSyncProfile(_:_:)
func CGColorSpaceCreateWithColorSyncProfile(arg0 ColorSyncProfileRef, options corefoundation.CFDictionaryRef) CGColorSpaceRef {
	if _cGColorSpaceCreateWithColorSyncProfile == nil {
		panic("CoreGraphics: symbol CGColorSpaceCreateWithColorSyncProfile not loaded")
	}
	return _cGColorSpaceCreateWithColorSyncProfile(arg0, options)
}

var _cGColorSpaceCreateWithICCData func(data corefoundation.CFTypeRef) CGColorSpaceRef

// CGColorSpaceCreateWithICCData creates an ICC-based color space using the ICC profile contained in the specified data.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(iccData:)
func CGColorSpaceCreateWithICCData(data corefoundation.CFTypeRef) CGColorSpaceRef {
	if _cGColorSpaceCreateWithICCData == nil {
		panic("CoreGraphics: symbol CGColorSpaceCreateWithICCData not loaded")
	}
	return _cGColorSpaceCreateWithICCData(data)
}

var _cGColorSpaceCreateWithName func(name corefoundation.CFStringRef) CGColorSpaceRef

// CGColorSpaceCreateWithName creates a specified type of Quartz color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(name:)
func CGColorSpaceCreateWithName(name corefoundation.CFStringRef) CGColorSpaceRef {
	if _cGColorSpaceCreateWithName == nil {
		panic("CoreGraphics: symbol CGColorSpaceCreateWithName not loaded")
	}
	return _cGColorSpaceCreateWithName(name)
}

var _cGColorSpaceCreateWithPropertyList func(plist corefoundation.CFPropertyListRef) CGColorSpaceRef

// CGColorSpaceCreateWithPropertyList creates a color space from a property list.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(propertyListPlist:)
func CGColorSpaceCreateWithPropertyList(plist corefoundation.CFPropertyListRef) CGColorSpaceRef {
	if _cGColorSpaceCreateWithPropertyList == nil {
		panic("CoreGraphics: symbol CGColorSpaceCreateWithPropertyList not loaded")
	}
	return _cGColorSpaceCreateWithPropertyList(plist)
}

var _cGColorSpaceGetBaseColorSpace func(space CGColorSpaceRef) CGColorSpaceRef

// CGColorSpaceGetBaseColorSpace returns the base color space of a pattern or indexed color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/baseColorSpace
func CGColorSpaceGetBaseColorSpace(space CGColorSpaceRef) CGColorSpaceRef {
	if _cGColorSpaceGetBaseColorSpace == nil {
		panic("CoreGraphics: symbol CGColorSpaceGetBaseColorSpace not loaded")
	}
	return _cGColorSpaceGetBaseColorSpace(space)
}

var _cGColorSpaceGetColorTable func(space CGColorSpaceRef, table *uint8)

// CGColorSpaceGetColorTable copies the entries in the color table of an indexed color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceGetColorTable
func CGColorSpaceGetColorTable(space CGColorSpaceRef, table *uint8) {
	if _cGColorSpaceGetColorTable == nil {
		panic("CoreGraphics: symbol CGColorSpaceGetColorTable not loaded")
	}
	_cGColorSpaceGetColorTable(space, table)
}

var _cGColorSpaceGetColorTableCount func(space CGColorSpaceRef) uintptr

// CGColorSpaceGetColorTableCount returns the number of entries in the color table of an indexed color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceGetColorTableCount
func CGColorSpaceGetColorTableCount(space CGColorSpaceRef) uintptr {
	if _cGColorSpaceGetColorTableCount == nil {
		panic("CoreGraphics: symbol CGColorSpaceGetColorTableCount not loaded")
	}
	return _cGColorSpaceGetColorTableCount(space)
}

var _cGColorSpaceGetModel func(space CGColorSpaceRef) CGColorSpaceModel

// CGColorSpaceGetModel returns the color space model of the provided color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/model
func CGColorSpaceGetModel(space CGColorSpaceRef) CGColorSpaceModel {
	if _cGColorSpaceGetModel == nil {
		panic("CoreGraphics: symbol CGColorSpaceGetModel not loaded")
	}
	return _cGColorSpaceGetModel(space)
}

var _cGColorSpaceGetName func(space CGColorSpaceRef) corefoundation.CFStringRef

// CGColorSpaceGetName.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceGetName
func CGColorSpaceGetName(space CGColorSpaceRef) corefoundation.CFStringRef {
	if _cGColorSpaceGetName == nil {
		panic("CoreGraphics: symbol CGColorSpaceGetName not loaded")
	}
	return _cGColorSpaceGetName(space)
}

var _cGColorSpaceGetNumberOfComponents func(space CGColorSpaceRef) uintptr

// CGColorSpaceGetNumberOfComponents returns the number of color components in a color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/numberOfComponents
func CGColorSpaceGetNumberOfComponents(space CGColorSpaceRef) uintptr {
	if _cGColorSpaceGetNumberOfComponents == nil {
		panic("CoreGraphics: symbol CGColorSpaceGetNumberOfComponents not loaded")
	}
	return _cGColorSpaceGetNumberOfComponents(space)
}

var _cGColorSpaceGetTypeID func() uint

// CGColorSpaceGetTypeID returns the Core Foundation type identifier for Quartz color spaces.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/typeID
func CGColorSpaceGetTypeID() uint {
	if _cGColorSpaceGetTypeID == nil {
		panic("CoreGraphics: symbol CGColorSpaceGetTypeID not loaded")
	}
	return _cGColorSpaceGetTypeID()
}

var _cGColorSpaceIsHDR func(arg0 CGColorSpaceRef) bool

// CGColorSpaceIsHDR.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/isHDR()
func CGColorSpaceIsHDR(arg0 CGColorSpaceRef) bool {
	if _cGColorSpaceIsHDR == nil {
		panic("CoreGraphics: symbol CGColorSpaceIsHDR not loaded")
	}
	return _cGColorSpaceIsHDR(arg0)
}

var _cGColorSpaceIsHLGBased func(s CGColorSpaceRef) bool

// CGColorSpaceIsHLGBased.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceIsHLGBased(_:)
func CGColorSpaceIsHLGBased(s CGColorSpaceRef) bool {
	if _cGColorSpaceIsHLGBased == nil {
		panic("CoreGraphics: symbol CGColorSpaceIsHLGBased not loaded")
	}
	return _cGColorSpaceIsHLGBased(s)
}

var _cGColorSpaceIsPQBased func(s CGColorSpaceRef) bool

// CGColorSpaceIsPQBased.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceIsPQBased(_:)
func CGColorSpaceIsPQBased(s CGColorSpaceRef) bool {
	if _cGColorSpaceIsPQBased == nil {
		panic("CoreGraphics: symbol CGColorSpaceIsPQBased not loaded")
	}
	return _cGColorSpaceIsPQBased(s)
}

var _cGColorSpaceIsWideGamutRGB func(arg0 CGColorSpaceRef) bool

// CGColorSpaceIsWideGamutRGB returns whether the RGB color space covers a significant portion of the NTSC color gamut.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/isWideGamutRGB
func CGColorSpaceIsWideGamutRGB(arg0 CGColorSpaceRef) bool {
	if _cGColorSpaceIsWideGamutRGB == nil {
		panic("CoreGraphics: symbol CGColorSpaceIsWideGamutRGB not loaded")
	}
	return _cGColorSpaceIsWideGamutRGB(arg0)
}

var _cGColorSpaceRelease func(space CGColorSpaceRef)

// CGColorSpaceRelease decrements the retain count of a color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceRelease
func CGColorSpaceRelease(space CGColorSpaceRef) {
	if _cGColorSpaceRelease == nil {
		panic("CoreGraphics: symbol CGColorSpaceRelease not loaded")
	}
	_cGColorSpaceRelease(space)
}

var _cGColorSpaceRetain func(space CGColorSpaceRef) CGColorSpaceRef

// CGColorSpaceRetain increments the retain count of a color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceRetain
func CGColorSpaceRetain(space CGColorSpaceRef) CGColorSpaceRef {
	if _cGColorSpaceRetain == nil {
		panic("CoreGraphics: symbol CGColorSpaceRetain not loaded")
	}
	return _cGColorSpaceRetain(space)
}

var _cGColorSpaceSupportsOutput func(space CGColorSpaceRef) bool

// CGColorSpaceSupportsOutput returns a Boolean indicating whether the color space can be used as a destination color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/supportsOutput
func CGColorSpaceSupportsOutput(space CGColorSpaceRef) bool {
	if _cGColorSpaceSupportsOutput == nil {
		panic("CoreGraphics: symbol CGColorSpaceSupportsOutput not loaded")
	}
	return _cGColorSpaceSupportsOutput(space)
}

var _cGColorSpaceUsesExtendedRange func(space CGColorSpaceRef) bool

// CGColorSpaceUsesExtendedRange.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceUsesExtendedRange(_:)
func CGColorSpaceUsesExtendedRange(space CGColorSpaceRef) bool {
	if _cGColorSpaceUsesExtendedRange == nil {
		panic("CoreGraphics: symbol CGColorSpaceUsesExtendedRange not loaded")
	}
	return _cGColorSpaceUsesExtendedRange(space)
}

var _cGColorSpaceUsesITUR_2100TF func(arg0 CGColorSpaceRef) bool

// CGColorSpaceUsesITUR_2100TF.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceUsesITUR_2100TF(_:)
func CGColorSpaceUsesITUR_2100TF(arg0 CGColorSpaceRef) bool {
	if _cGColorSpaceUsesITUR_2100TF == nil {
		panic("CoreGraphics: symbol CGColorSpaceUsesITUR_2100TF not loaded")
	}
	return _cGColorSpaceUsesITUR_2100TF(arg0)
}

var _cGCompleteDisplayConfiguration func(config CGDisplayConfigRef, option CGConfigureOption) CGError

// CGCompleteDisplayConfiguration completes a set of display configuration changes.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGCompleteDisplayConfiguration(_:_:)
func CGCompleteDisplayConfiguration(config CGDisplayConfigRef, option CGConfigureOption) CGError {
	if _cGCompleteDisplayConfiguration == nil {
		panic("CoreGraphics: symbol CGCompleteDisplayConfiguration not loaded")
	}
	return _cGCompleteDisplayConfiguration(config, option)
}

var _cGConfigureDisplayFadeEffect func(config CGDisplayConfigRef, fadeOutSeconds float32, fadeInSeconds float32, fadeRed float32, fadeGreen float32, fadeBlue float32) CGError

// CGConfigureDisplayFadeEffect modifies the settings of the built-in fade effect that occurs during a display configuration.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGConfigureDisplayFadeEffect(_:_:_:_:_:_:)
func CGConfigureDisplayFadeEffect(config CGDisplayConfigRef, fadeOutSeconds float32, fadeInSeconds float32, fadeRed float32, fadeGreen float32, fadeBlue float32) CGError {
	if _cGConfigureDisplayFadeEffect == nil {
		panic("CoreGraphics: symbol CGConfigureDisplayFadeEffect not loaded")
	}
	return _cGConfigureDisplayFadeEffect(config, fadeOutSeconds, fadeInSeconds, fadeRed, fadeGreen, fadeBlue)
}

var _cGConfigureDisplayMirrorOfDisplay func(config CGDisplayConfigRef, display uint32, master uint32) CGError

// CGConfigureDisplayMirrorOfDisplay changes the configuration of a mirroring set.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGConfigureDisplayMirrorOfDisplay(_:_:_:)
func CGConfigureDisplayMirrorOfDisplay(config CGDisplayConfigRef, display uint32, master uint32) CGError {
	if _cGConfigureDisplayMirrorOfDisplay == nil {
		panic("CoreGraphics: symbol CGConfigureDisplayMirrorOfDisplay not loaded")
	}
	return _cGConfigureDisplayMirrorOfDisplay(config, display, master)
}

var _cGConfigureDisplayOrigin func(config CGDisplayConfigRef, display uint32, x int32, y int32) CGError

// CGConfigureDisplayOrigin configures the origin of a display relative to the global display coordinate space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGConfigureDisplayOrigin(_:_:_:_:)
func CGConfigureDisplayOrigin(config CGDisplayConfigRef, display uint32, x int32, y int32) CGError {
	if _cGConfigureDisplayOrigin == nil {
		panic("CoreGraphics: symbol CGConfigureDisplayOrigin not loaded")
	}
	return _cGConfigureDisplayOrigin(config, display, x, y)
}

var _cGConfigureDisplayStereoOperation func(config CGDisplayConfigRef, display uint32, stereo uintptr, forceBlueLine uintptr) CGError

// CGConfigureDisplayStereoOperation enables or disables stereo operation for a display, as part of a display configuration.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGConfigureDisplayStereoOperation(_:_:_:_:)
func CGConfigureDisplayStereoOperation(config CGDisplayConfigRef, display uint32, stereo uintptr, forceBlueLine uintptr) CGError {
	if _cGConfigureDisplayStereoOperation == nil {
		panic("CoreGraphics: symbol CGConfigureDisplayStereoOperation not loaded")
	}
	return _cGConfigureDisplayStereoOperation(config, display, stereo, forceBlueLine)
}

var _cGConfigureDisplayWithDisplayMode func(config CGDisplayConfigRef, display uint32, mode CGDisplayModeRef, options corefoundation.CFDictionaryRef) CGError

// CGConfigureDisplayWithDisplayMode configures the display mode of a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGConfigureDisplayWithDisplayMode(_:_:_:_:)
func CGConfigureDisplayWithDisplayMode(config CGDisplayConfigRef, display uint32, mode CGDisplayModeRef, options corefoundation.CFDictionaryRef) CGError {
	if _cGConfigureDisplayWithDisplayMode == nil {
		panic("CoreGraphics: symbol CGConfigureDisplayWithDisplayMode not loaded")
	}
	return _cGConfigureDisplayWithDisplayMode(config, display, mode, options)
}

var _cGContextAddArc func(c CGContextRef, x float64, y float64, radius float64, startAngle float64, endAngle float64, clockwise int)

// CGContextAddArc adds an arc of a circle to the current path, possibly preceded by a straight line segment
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextAddArc
func CGContextAddArc(c CGContextRef, x float64, y float64, radius float64, startAngle float64, endAngle float64, clockwise int) {
	if _cGContextAddArc == nil {
		panic("CoreGraphics: symbol CGContextAddArc not loaded")
	}
	_cGContextAddArc(c, x, y, radius, startAngle, endAngle, clockwise)
}

var _cGContextAddArcToPoint func(c CGContextRef, x1 float64, y1 float64, x2 float64, y2 float64, radius float64)

// CGContextAddArcToPoint adds an arc of a circle to the current path, using a radius and tangent points.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextAddArcToPoint
func CGContextAddArcToPoint(c CGContextRef, x1 float64, y1 float64, x2 float64, y2 float64, radius float64) {
	if _cGContextAddArcToPoint == nil {
		panic("CoreGraphics: symbol CGContextAddArcToPoint not loaded")
	}
	_cGContextAddArcToPoint(c, x1, y1, x2, y2, radius)
}

var _cGContextAddCurveToPoint func(c CGContextRef, cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64)

// CGContextAddCurveToPoint appends a cubic Bézier curve from the current point, using the provided control points and end point .
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextAddCurveToPoint
func CGContextAddCurveToPoint(c CGContextRef, cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) {
	if _cGContextAddCurveToPoint == nil {
		panic("CoreGraphics: symbol CGContextAddCurveToPoint not loaded")
	}
	_cGContextAddCurveToPoint(c, cp1x, cp1y, cp2x, cp2y, x, y)
}

var _cGContextAddEllipseInRect func(c CGContextRef, rect corefoundation.CGRect)

// CGContextAddEllipseInRect adds an ellipse that fits inside the specified rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/addEllipse(in:)
func CGContextAddEllipseInRect(c CGContextRef, rect corefoundation.CGRect) {
	if _cGContextAddEllipseInRect == nil {
		panic("CoreGraphics: symbol CGContextAddEllipseInRect not loaded")
	}
	_cGContextAddEllipseInRect(c, rect)
}

var _cGContextAddLineToPoint func(c CGContextRef, x float64, y float64)

// CGContextAddLineToPoint appends a straight line segment from the current point to the provided point .
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextAddLineToPoint
func CGContextAddLineToPoint(c CGContextRef, x float64, y float64) {
	if _cGContextAddLineToPoint == nil {
		panic("CoreGraphics: symbol CGContextAddLineToPoint not loaded")
	}
	_cGContextAddLineToPoint(c, x, y)
}

var _cGContextAddLines func(c CGContextRef, points *corefoundation.CGPoint, count uintptr)

// CGContextAddLines adds a sequence of connected straight-line segments to the current path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextAddLines
func CGContextAddLines(c CGContextRef, points *corefoundation.CGPoint, count uintptr) {
	if _cGContextAddLines == nil {
		panic("CoreGraphics: symbol CGContextAddLines not loaded")
	}
	_cGContextAddLines(c, points, count)
}

var _cGContextAddPath func(c CGContextRef, path CGPathRef)

// CGContextAddPath adds a previously created path object to the current path in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/addPath(_:)
func CGContextAddPath(c CGContextRef, path CGPathRef) {
	if _cGContextAddPath == nil {
		panic("CoreGraphics: symbol CGContextAddPath not loaded")
	}
	_cGContextAddPath(c, path)
}

var _cGContextAddQuadCurveToPoint func(c CGContextRef, cpx float64, cpy float64, x float64, y float64)

// CGContextAddQuadCurveToPoint appends a quadratic Bézier curve from the current point, using a control point and an end point you specify.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextAddQuadCurveToPoint
func CGContextAddQuadCurveToPoint(c CGContextRef, cpx float64, cpy float64, x float64, y float64) {
	if _cGContextAddQuadCurveToPoint == nil {
		panic("CoreGraphics: symbol CGContextAddQuadCurveToPoint not loaded")
	}
	_cGContextAddQuadCurveToPoint(c, cpx, cpy, x, y)
}

var _cGContextAddRect func(c CGContextRef, rect corefoundation.CGRect)

// CGContextAddRect adds a rectangular path to the current path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/addRect(_:)
func CGContextAddRect(c CGContextRef, rect corefoundation.CGRect) {
	if _cGContextAddRect == nil {
		panic("CoreGraphics: symbol CGContextAddRect not loaded")
	}
	_cGContextAddRect(c, rect)
}

var _cGContextAddRects func(c CGContextRef, rects *corefoundation.CGRect, count uintptr)

// CGContextAddRects adds a set of rectangular paths to the current path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextAddRects
func CGContextAddRects(c CGContextRef, rects *corefoundation.CGRect, count uintptr) {
	if _cGContextAddRects == nil {
		panic("CoreGraphics: symbol CGContextAddRects not loaded")
	}
	_cGContextAddRects(c, rects, count)
}

var _cGContextBeginPage func(c CGContextRef, mediaBox *corefoundation.CGRect)

// CGContextBeginPage starts a new page in a page-based graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/beginPage(mediaBox:)
func CGContextBeginPage(c CGContextRef, mediaBox *corefoundation.CGRect) {
	if _cGContextBeginPage == nil {
		panic("CoreGraphics: symbol CGContextBeginPage not loaded")
	}
	_cGContextBeginPage(c, mediaBox)
}

var _cGContextBeginPath func(c CGContextRef)

// CGContextBeginPath creates a new empty path in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/beginPath()
func CGContextBeginPath(c CGContextRef) {
	if _cGContextBeginPath == nil {
		panic("CoreGraphics: symbol CGContextBeginPath not loaded")
	}
	_cGContextBeginPath(c)
}

var _cGContextBeginTransparencyLayer func(c CGContextRef, auxiliaryInfo corefoundation.CFDictionaryRef)

// CGContextBeginTransparencyLayer begins a transparency layer.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/beginTransparencyLayer(auxiliaryInfo:)
func CGContextBeginTransparencyLayer(c CGContextRef, auxiliaryInfo corefoundation.CFDictionaryRef) {
	if _cGContextBeginTransparencyLayer == nil {
		panic("CoreGraphics: symbol CGContextBeginTransparencyLayer not loaded")
	}
	_cGContextBeginTransparencyLayer(c, auxiliaryInfo)
}

var _cGContextBeginTransparencyLayerWithRect func(c CGContextRef, rect corefoundation.CGRect, auxInfo corefoundation.CFDictionaryRef)

// CGContextBeginTransparencyLayerWithRect begins a transparency layer whose contents are bounded by the specified rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/beginTransparencyLayer(in:auxiliaryInfo:)
func CGContextBeginTransparencyLayerWithRect(c CGContextRef, rect corefoundation.CGRect, auxInfo corefoundation.CFDictionaryRef) {
	if _cGContextBeginTransparencyLayerWithRect == nil {
		panic("CoreGraphics: symbol CGContextBeginTransparencyLayerWithRect not loaded")
	}
	_cGContextBeginTransparencyLayerWithRect(c, rect, auxInfo)
}

var _cGContextClearRect func(c CGContextRef, rect corefoundation.CGRect)

// CGContextClearRect paints a transparent rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/clear(_:)
func CGContextClearRect(c CGContextRef, rect corefoundation.CGRect) {
	if _cGContextClearRect == nil {
		panic("CoreGraphics: symbol CGContextClearRect not loaded")
	}
	_cGContextClearRect(c, rect)
}

var _cGContextClip func(c CGContextRef)

// CGContextClip modifies the current clipping path, using the nonzero winding number rule.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextClip
func CGContextClip(c CGContextRef) {
	if _cGContextClip == nil {
		panic("CoreGraphics: symbol CGContextClip not loaded")
	}
	_cGContextClip(c)
}

var _cGContextClipToMask func(c CGContextRef, rect corefoundation.CGRect, mask CGImageRef)

// CGContextClipToMask maps a mask into the specified rectangle and intersects it with the current clipping area of the graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/clip(to:mask:)
func CGContextClipToMask(c CGContextRef, rect corefoundation.CGRect, mask CGImageRef) {
	if _cGContextClipToMask == nil {
		panic("CoreGraphics: symbol CGContextClipToMask not loaded")
	}
	_cGContextClipToMask(c, rect, mask)
}

var _cGContextClipToRect func(c CGContextRef, rect corefoundation.CGRect)

// CGContextClipToRect sets the clipping path to the intersection of the current clipping path with the area defined by the specified rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/clip(to:)-7cbwq
func CGContextClipToRect(c CGContextRef, rect corefoundation.CGRect) {
	if _cGContextClipToRect == nil {
		panic("CoreGraphics: symbol CGContextClipToRect not loaded")
	}
	_cGContextClipToRect(c, rect)
}

var _cGContextClipToRects func(c CGContextRef, rects *corefoundation.CGRect, count uintptr)

// CGContextClipToRects sets the clipping path to the intersection of the current clipping path with the region defined by an array of rectangles.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextClipToRects
func CGContextClipToRects(c CGContextRef, rects *corefoundation.CGRect, count uintptr) {
	if _cGContextClipToRects == nil {
		panic("CoreGraphics: symbol CGContextClipToRects not loaded")
	}
	_cGContextClipToRects(c, rects, count)
}

var _cGContextClosePath func(c CGContextRef)

// CGContextClosePath closes and terminates the current path’s subpath.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/closePath()
func CGContextClosePath(c CGContextRef) {
	if _cGContextClosePath == nil {
		panic("CoreGraphics: symbol CGContextClosePath not loaded")
	}
	_cGContextClosePath(c)
}

var _cGContextConcatCTM func(c CGContextRef, transform corefoundation.CGAffineTransform)

// CGContextConcatCTM transforms the user coordinate system in a context using a specified matrix.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/concatenate(_:)
func CGContextConcatCTM(c CGContextRef, transform corefoundation.CGAffineTransform) {
	if _cGContextConcatCTM == nil {
		panic("CoreGraphics: symbol CGContextConcatCTM not loaded")
	}
	_cGContextConcatCTM(c, transform)
}

var _cGContextConvertPointToDeviceSpace func(c CGContextRef, point corefoundation.CGPoint) corefoundation.CGPoint

// CGContextConvertPointToDeviceSpace returns a point that is transformed from user space coordinates to device space coordinates.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/convertToDeviceSpace(_:)-53m7u
func CGContextConvertPointToDeviceSpace(c CGContextRef, point corefoundation.CGPoint) corefoundation.CGPoint {
	if _cGContextConvertPointToDeviceSpace == nil {
		panic("CoreGraphics: symbol CGContextConvertPointToDeviceSpace not loaded")
	}
	return _cGContextConvertPointToDeviceSpace(c, point)
}

var _cGContextConvertPointToUserSpace func(c CGContextRef, point corefoundation.CGPoint) corefoundation.CGPoint

// CGContextConvertPointToUserSpace returns a point that is transformed from device space coordinates to user space coordinates.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/convertToUserSpace(_:)-3mtg3
func CGContextConvertPointToUserSpace(c CGContextRef, point corefoundation.CGPoint) corefoundation.CGPoint {
	if _cGContextConvertPointToUserSpace == nil {
		panic("CoreGraphics: symbol CGContextConvertPointToUserSpace not loaded")
	}
	return _cGContextConvertPointToUserSpace(c, point)
}

var _cGContextConvertRectToDeviceSpace func(c CGContextRef, rect corefoundation.CGRect) corefoundation.CGRect

// CGContextConvertRectToDeviceSpace returns a rectangle that is transformed from user space coordinate to device space coordinates.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/convertToDeviceSpace(_:)-91x5g
func CGContextConvertRectToDeviceSpace(c CGContextRef, rect corefoundation.CGRect) corefoundation.CGRect {
	if _cGContextConvertRectToDeviceSpace == nil {
		panic("CoreGraphics: symbol CGContextConvertRectToDeviceSpace not loaded")
	}
	return _cGContextConvertRectToDeviceSpace(c, rect)
}

var _cGContextConvertRectToUserSpace func(c CGContextRef, rect corefoundation.CGRect) corefoundation.CGRect

// CGContextConvertRectToUserSpace returns a rectangle that is transformed from device space coordinate to user space coordinates.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/convertToUserSpace(_:)-1hk5r
func CGContextConvertRectToUserSpace(c CGContextRef, rect corefoundation.CGRect) corefoundation.CGRect {
	if _cGContextConvertRectToUserSpace == nil {
		panic("CoreGraphics: symbol CGContextConvertRectToUserSpace not loaded")
	}
	return _cGContextConvertRectToUserSpace(c, rect)
}

var _cGContextConvertSizeToDeviceSpace func(c CGContextRef, size corefoundation.CGSize) corefoundation.CGSize

// CGContextConvertSizeToDeviceSpace returns a size that is transformed from user space coordinates to device space coordinates.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/convertToDeviceSpace(_:)-224h2
func CGContextConvertSizeToDeviceSpace(c CGContextRef, size corefoundation.CGSize) corefoundation.CGSize {
	if _cGContextConvertSizeToDeviceSpace == nil {
		panic("CoreGraphics: symbol CGContextConvertSizeToDeviceSpace not loaded")
	}
	return _cGContextConvertSizeToDeviceSpace(c, size)
}

var _cGContextConvertSizeToUserSpace func(c CGContextRef, size corefoundation.CGSize) corefoundation.CGSize

// CGContextConvertSizeToUserSpace returns a size that is transformed from device space coordinates to user space coordinates.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/convertToUserSpace(_:)-693ur
func CGContextConvertSizeToUserSpace(c CGContextRef, size corefoundation.CGSize) corefoundation.CGSize {
	if _cGContextConvertSizeToUserSpace == nil {
		panic("CoreGraphics: symbol CGContextConvertSizeToUserSpace not loaded")
	}
	return _cGContextConvertSizeToUserSpace(c, size)
}

var _cGContextCopyPath func(c CGContextRef) CGPathRef

// CGContextCopyPath returns a path object built from the current path information in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/path
func CGContextCopyPath(c CGContextRef) CGPathRef {
	if _cGContextCopyPath == nil {
		panic("CoreGraphics: symbol CGContextCopyPath not loaded")
	}
	return _cGContextCopyPath(c)
}

var _cGContextDrawConicGradient func(c CGContextRef, gradient CGGradientRef, center corefoundation.CGPoint, angle float64)

// CGContextDrawConicGradient.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawConicGradient(_:_:_:_:)
func CGContextDrawConicGradient(c CGContextRef, gradient CGGradientRef, center corefoundation.CGPoint, angle float64) {
	if _cGContextDrawConicGradient == nil {
		panic("CoreGraphics: symbol CGContextDrawConicGradient not loaded")
	}
	_cGContextDrawConicGradient(c, gradient, center, angle)
}

var _cGContextDrawImage func(c CGContextRef, rect corefoundation.CGRect, image CGImageRef)

// CGContextDrawImage draws an image into a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawImage
func CGContextDrawImage(c CGContextRef, rect corefoundation.CGRect, image CGImageRef) {
	if _cGContextDrawImage == nil {
		panic("CoreGraphics: symbol CGContextDrawImage not loaded")
	}
	_cGContextDrawImage(c, rect, image)
}

var _cGContextDrawImageApplyingToneMapping func(c CGContextRef, r corefoundation.CGRect, image CGImageRef, method CGToneMapping, options corefoundation.CFDictionaryRef) bool

// CGContextDrawImageApplyingToneMapping.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawImageApplyingToneMapping
func CGContextDrawImageApplyingToneMapping(c CGContextRef, r corefoundation.CGRect, image CGImageRef, method CGToneMapping, options corefoundation.CFDictionaryRef) bool {
	if _cGContextDrawImageApplyingToneMapping == nil {
		panic("CoreGraphics: symbol CGContextDrawImageApplyingToneMapping not loaded")
	}
	return _cGContextDrawImageApplyingToneMapping(c, r, image, method, options)
}

var _cGContextDrawLayerAtPoint func(context CGContextRef, point corefoundation.CGPoint, layer uintptr)

// CGContextDrawLayerAtPoint draws the contents of a CGLayer object at the specified point.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawLayerAtPoint
func CGContextDrawLayerAtPoint(context CGContextRef, point corefoundation.CGPoint, layer uintptr) {
	if _cGContextDrawLayerAtPoint == nil {
		panic("CoreGraphics: symbol CGContextDrawLayerAtPoint not loaded")
	}
	_cGContextDrawLayerAtPoint(context, point, layer)
}

var _cGContextDrawLayerInRect func(context CGContextRef, rect corefoundation.CGRect, layer uintptr)

// CGContextDrawLayerInRect draws the contents of a layer object into the specified rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawLayerInRect
func CGContextDrawLayerInRect(context CGContextRef, rect corefoundation.CGRect, layer uintptr) {
	if _cGContextDrawLayerInRect == nil {
		panic("CoreGraphics: symbol CGContextDrawLayerInRect not loaded")
	}
	_cGContextDrawLayerInRect(context, rect, layer)
}

var _cGContextDrawLinearGradient func(c CGContextRef, gradient CGGradientRef, startPoint corefoundation.CGPoint, endPoint corefoundation.CGPoint, options CGGradientDrawingOptions)

// CGContextDrawLinearGradient paints a gradient fill that varies along the line defined by the provided starting and ending points.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/drawLinearGradient(_:start:end:options:)
func CGContextDrawLinearGradient(c CGContextRef, gradient CGGradientRef, startPoint corefoundation.CGPoint, endPoint corefoundation.CGPoint, options CGGradientDrawingOptions) {
	if _cGContextDrawLinearGradient == nil {
		panic("CoreGraphics: symbol CGContextDrawLinearGradient not loaded")
	}
	_cGContextDrawLinearGradient(c, gradient, startPoint, endPoint, options)
}

var _cGContextDrawPDFDocument func(c CGContextRef, rect corefoundation.CGRect, document CGPDFDocumentRef, page int)

// CGContextDrawPDFDocument.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawPDFDocument
func CGContextDrawPDFDocument(c CGContextRef, rect corefoundation.CGRect, document CGPDFDocumentRef, page int) {
	if _cGContextDrawPDFDocument == nil {
		panic("CoreGraphics: symbol CGContextDrawPDFDocument not loaded")
	}
	_cGContextDrawPDFDocument(c, rect, document, page)
}

var _cGContextDrawPDFPage func(c CGContextRef, page CGPDFPageRef)

// CGContextDrawPDFPage draws the content of a PDF page into the current graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/drawPDFPage(_:)
func CGContextDrawPDFPage(c CGContextRef, page CGPDFPageRef) {
	if _cGContextDrawPDFPage == nil {
		panic("CoreGraphics: symbol CGContextDrawPDFPage not loaded")
	}
	_cGContextDrawPDFPage(c, page)
}

var _cGContextDrawPath func(c CGContextRef, mode CGPathDrawingMode)

// CGContextDrawPath draws the current path using the provided drawing mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/drawPath(using:)
func CGContextDrawPath(c CGContextRef, mode CGPathDrawingMode) {
	if _cGContextDrawPath == nil {
		panic("CoreGraphics: symbol CGContextDrawPath not loaded")
	}
	_cGContextDrawPath(c, mode)
}

var _cGContextDrawRadialGradient func(c CGContextRef, gradient CGGradientRef, startCenter corefoundation.CGPoint, startRadius float64, endCenter corefoundation.CGPoint, endRadius float64, options CGGradientDrawingOptions)

// CGContextDrawRadialGradient paints a gradient fill that varies along the area defined by the provided starting and ending circles.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/drawRadialGradient(_:startCenter:startRadius:endCenter:endRadius:options:)
func CGContextDrawRadialGradient(c CGContextRef, gradient CGGradientRef, startCenter corefoundation.CGPoint, startRadius float64, endCenter corefoundation.CGPoint, endRadius float64, options CGGradientDrawingOptions) {
	if _cGContextDrawRadialGradient == nil {
		panic("CoreGraphics: symbol CGContextDrawRadialGradient not loaded")
	}
	_cGContextDrawRadialGradient(c, gradient, startCenter, startRadius, endCenter, endRadius, options)
}

var _cGContextDrawShading func(c CGContextRef, shading CGShadingRef)

// CGContextDrawShading fills the clipping path of a context with the specified shading.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/drawShading(_:)
func CGContextDrawShading(c CGContextRef, shading CGShadingRef) {
	if _cGContextDrawShading == nil {
		panic("CoreGraphics: symbol CGContextDrawShading not loaded")
	}
	_cGContextDrawShading(c, shading)
}

var _cGContextDrawTiledImage func(c CGContextRef, rect corefoundation.CGRect, image CGImageRef)

// CGContextDrawTiledImage repeatedly draws an image, scaled to the provided rectangle, to fill the current clip region.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawTiledImage
func CGContextDrawTiledImage(c CGContextRef, rect corefoundation.CGRect, image CGImageRef) {
	if _cGContextDrawTiledImage == nil {
		panic("CoreGraphics: symbol CGContextDrawTiledImage not loaded")
	}
	_cGContextDrawTiledImage(c, rect, image)
}

var _cGContextEOClip func(c CGContextRef)

// CGContextEOClip modifies the current clipping path, using the even-odd rule.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextEOClip
func CGContextEOClip(c CGContextRef) {
	if _cGContextEOClip == nil {
		panic("CoreGraphics: symbol CGContextEOClip not loaded")
	}
	_cGContextEOClip(c)
}

var _cGContextEOFillPath func(c CGContextRef)

// CGContextEOFillPath paints the area within the current path, using the even-odd fill rule.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextEOFillPath
func CGContextEOFillPath(c CGContextRef) {
	if _cGContextEOFillPath == nil {
		panic("CoreGraphics: symbol CGContextEOFillPath not loaded")
	}
	_cGContextEOFillPath(c)
}

var _cGContextEndPage func(c CGContextRef)

// CGContextEndPage ends the current page in a page-based graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/endPage()
func CGContextEndPage(c CGContextRef) {
	if _cGContextEndPage == nil {
		panic("CoreGraphics: symbol CGContextEndPage not loaded")
	}
	_cGContextEndPage(c)
}

var _cGContextEndTransparencyLayer func(c CGContextRef)

// CGContextEndTransparencyLayer ends a transparency layer.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/endTransparencyLayer()
func CGContextEndTransparencyLayer(c CGContextRef) {
	if _cGContextEndTransparencyLayer == nil {
		panic("CoreGraphics: symbol CGContextEndTransparencyLayer not loaded")
	}
	_cGContextEndTransparencyLayer(c)
}

var _cGContextFillEllipseInRect func(c CGContextRef, rect corefoundation.CGRect)

// CGContextFillEllipseInRect paints the area of the ellipse that fits inside the provided rectangle, using the fill color in the current graphics state.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/fillEllipse(in:)
func CGContextFillEllipseInRect(c CGContextRef, rect corefoundation.CGRect) {
	if _cGContextFillEllipseInRect == nil {
		panic("CoreGraphics: symbol CGContextFillEllipseInRect not loaded")
	}
	_cGContextFillEllipseInRect(c, rect)
}

var _cGContextFillPath func(c CGContextRef)

// CGContextFillPath paints the area within the current path, using the nonzero winding number rule.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextFillPath
func CGContextFillPath(c CGContextRef) {
	if _cGContextFillPath == nil {
		panic("CoreGraphics: symbol CGContextFillPath not loaded")
	}
	_cGContextFillPath(c)
}

var _cGContextFillRect func(c CGContextRef, rect corefoundation.CGRect)

// CGContextFillRect paints the area contained within the provided rectangle, using the fill color in the current graphics state.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/fill(_:)-7a0rk
func CGContextFillRect(c CGContextRef, rect corefoundation.CGRect) {
	if _cGContextFillRect == nil {
		panic("CoreGraphics: symbol CGContextFillRect not loaded")
	}
	_cGContextFillRect(c, rect)
}

var _cGContextFillRects func(c CGContextRef, rects *corefoundation.CGRect, count uintptr)

// CGContextFillRects paints the areas contained within the provided rectangles, using the fill color in the current graphics state.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextFillRects
func CGContextFillRects(c CGContextRef, rects *corefoundation.CGRect, count uintptr) {
	if _cGContextFillRects == nil {
		panic("CoreGraphics: symbol CGContextFillRects not loaded")
	}
	_cGContextFillRects(c, rects, count)
}

var _cGContextFlush func(c CGContextRef)

// CGContextFlush forces all pending drawing operations in a window context to be rendered immediately to the destination device.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/flush()
func CGContextFlush(c CGContextRef) {
	if _cGContextFlush == nil {
		panic("CoreGraphics: symbol CGContextFlush not loaded")
	}
	_cGContextFlush(c)
}

var _cGContextGetCTM func(c CGContextRef) corefoundation.CGAffineTransform

// CGContextGetCTM returns the current transformation matrix.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/ctm
func CGContextGetCTM(c CGContextRef) corefoundation.CGAffineTransform {
	if _cGContextGetCTM == nil {
		panic("CoreGraphics: symbol CGContextGetCTM not loaded")
	}
	return _cGContextGetCTM(c)
}

var _cGContextGetClipBoundingBox func(c CGContextRef) corefoundation.CGRect

// CGContextGetClipBoundingBox returns the bounding box of a clipping path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/boundingBoxOfClipPath
func CGContextGetClipBoundingBox(c CGContextRef) corefoundation.CGRect {
	if _cGContextGetClipBoundingBox == nil {
		panic("CoreGraphics: symbol CGContextGetClipBoundingBox not loaded")
	}
	return _cGContextGetClipBoundingBox(c)
}

var _cGContextGetContentToneMappingInfo func(c CGContextRef) CGContentToneMappingInfo

// CGContextGetContentToneMappingInfo.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextGetContentToneMappingInfo
func CGContextGetContentToneMappingInfo(c CGContextRef) CGContentToneMappingInfo {
	if _cGContextGetContentToneMappingInfo == nil {
		panic("CoreGraphics: symbol CGContextGetContentToneMappingInfo not loaded")
	}
	return _cGContextGetContentToneMappingInfo(c)
}

var _cGContextGetEDRTargetHeadroom func(c CGContextRef) float32

// CGContextGetEDRTargetHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextGetEDRTargetHeadroom(_:)
func CGContextGetEDRTargetHeadroom(c CGContextRef) float32 {
	if _cGContextGetEDRTargetHeadroom == nil {
		panic("CoreGraphics: symbol CGContextGetEDRTargetHeadroom not loaded")
	}
	return _cGContextGetEDRTargetHeadroom(c)
}

var _cGContextGetInterpolationQuality func(c CGContextRef) CGInterpolationQuality

// CGContextGetInterpolationQuality returns the current level of interpolation quality for a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/interpolationQuality
func CGContextGetInterpolationQuality(c CGContextRef) CGInterpolationQuality {
	if _cGContextGetInterpolationQuality == nil {
		panic("CoreGraphics: symbol CGContextGetInterpolationQuality not loaded")
	}
	return _cGContextGetInterpolationQuality(c)
}

var _cGContextGetPathBoundingBox func(c CGContextRef) corefoundation.CGRect

// CGContextGetPathBoundingBox returns the smallest rectangle that contains the current path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/boundingBoxOfPath
func CGContextGetPathBoundingBox(c CGContextRef) corefoundation.CGRect {
	if _cGContextGetPathBoundingBox == nil {
		panic("CoreGraphics: symbol CGContextGetPathBoundingBox not loaded")
	}
	return _cGContextGetPathBoundingBox(c)
}

var _cGContextGetPathCurrentPoint func(c CGContextRef) corefoundation.CGPoint

// CGContextGetPathCurrentPoint returns the current point in a non-empty path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/currentPointOfPath
func CGContextGetPathCurrentPoint(c CGContextRef) corefoundation.CGPoint {
	if _cGContextGetPathCurrentPoint == nil {
		panic("CoreGraphics: symbol CGContextGetPathCurrentPoint not loaded")
	}
	return _cGContextGetPathCurrentPoint(c)
}

var _cGContextGetTextMatrix func(c CGContextRef) corefoundation.CGAffineTransform

// CGContextGetTextMatrix returns the current text matrix.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/textMatrix
func CGContextGetTextMatrix(c CGContextRef) corefoundation.CGAffineTransform {
	if _cGContextGetTextMatrix == nil {
		panic("CoreGraphics: symbol CGContextGetTextMatrix not loaded")
	}
	return _cGContextGetTextMatrix(c)
}

var _cGContextGetTextPosition func(c CGContextRef) corefoundation.CGPoint

// CGContextGetTextPosition.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextGetTextPosition
func CGContextGetTextPosition(c CGContextRef) corefoundation.CGPoint {
	if _cGContextGetTextPosition == nil {
		panic("CoreGraphics: symbol CGContextGetTextPosition not loaded")
	}
	return _cGContextGetTextPosition(c)
}

var _cGContextGetTypeID func() uint

// CGContextGetTypeID returns the type identifier for a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/typeID
func CGContextGetTypeID() uint {
	if _cGContextGetTypeID == nil {
		panic("CoreGraphics: symbol CGContextGetTypeID not loaded")
	}
	return _cGContextGetTypeID()
}

var _cGContextGetUserSpaceToDeviceSpaceTransform func(c CGContextRef) corefoundation.CGAffineTransform

// CGContextGetUserSpaceToDeviceSpaceTransform returns an affine transform that maps user space coordinates to device space coordinates.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/userSpaceToDeviceSpaceTransform
func CGContextGetUserSpaceToDeviceSpaceTransform(c CGContextRef) corefoundation.CGAffineTransform {
	if _cGContextGetUserSpaceToDeviceSpaceTransform == nil {
		panic("CoreGraphics: symbol CGContextGetUserSpaceToDeviceSpaceTransform not loaded")
	}
	return _cGContextGetUserSpaceToDeviceSpaceTransform(c)
}

var _cGContextIsPathEmpty func(c CGContextRef) bool

// CGContextIsPathEmpty indicates whether the current path contains any subpaths.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/isPathEmpty
func CGContextIsPathEmpty(c CGContextRef) bool {
	if _cGContextIsPathEmpty == nil {
		panic("CoreGraphics: symbol CGContextIsPathEmpty not loaded")
	}
	return _cGContextIsPathEmpty(c)
}

var _cGContextMoveToPoint func(c CGContextRef, x float64, y float64)

// CGContextMoveToPoint begins a new subpath at the point you specify.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextMoveToPoint
func CGContextMoveToPoint(c CGContextRef, x float64, y float64) {
	if _cGContextMoveToPoint == nil {
		panic("CoreGraphics: symbol CGContextMoveToPoint not loaded")
	}
	_cGContextMoveToPoint(c, x, y)
}

var _cGContextPathContainsPoint func(c CGContextRef, point corefoundation.CGPoint, mode CGPathDrawingMode) bool

// CGContextPathContainsPoint checks to see whether the specified point is contained in the current path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/pathContains(_:mode:)
func CGContextPathContainsPoint(c CGContextRef, point corefoundation.CGPoint, mode CGPathDrawingMode) bool {
	if _cGContextPathContainsPoint == nil {
		panic("CoreGraphics: symbol CGContextPathContainsPoint not loaded")
	}
	return _cGContextPathContainsPoint(c, point, mode)
}

var _cGContextRelease func(c CGContextRef)

// CGContextRelease decrements the retain count of a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextRelease
func CGContextRelease(c CGContextRef) {
	if _cGContextRelease == nil {
		panic("CoreGraphics: symbol CGContextRelease not loaded")
	}
	_cGContextRelease(c)
}

var _cGContextReplacePathWithStrokedPath func(c CGContextRef)

// CGContextReplacePathWithStrokedPath replaces the path in the graphics context with the stroked version of the path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/replacePathWithStrokedPath()
func CGContextReplacePathWithStrokedPath(c CGContextRef) {
	if _cGContextReplacePathWithStrokedPath == nil {
		panic("CoreGraphics: symbol CGContextReplacePathWithStrokedPath not loaded")
	}
	_cGContextReplacePathWithStrokedPath(c)
}

var _cGContextResetClip func(c CGContextRef)

// CGContextResetClip.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/resetClip()
func CGContextResetClip(c CGContextRef) {
	if _cGContextResetClip == nil {
		panic("CoreGraphics: symbol CGContextResetClip not loaded")
	}
	_cGContextResetClip(c)
}

var _cGContextRestoreGState func(c CGContextRef)

// CGContextRestoreGState sets the current graphics state to the state most recently saved.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/restoreGState()
func CGContextRestoreGState(c CGContextRef) {
	if _cGContextRestoreGState == nil {
		panic("CoreGraphics: symbol CGContextRestoreGState not loaded")
	}
	_cGContextRestoreGState(c)
}

var _cGContextRetain func(c CGContextRef) CGContextRef

// CGContextRetain increments the retain count of a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextRetain
func CGContextRetain(c CGContextRef) CGContextRef {
	if _cGContextRetain == nil {
		panic("CoreGraphics: symbol CGContextRetain not loaded")
	}
	return _cGContextRetain(c)
}

var _cGContextRotateCTM func(c CGContextRef, angle float64)

// CGContextRotateCTM rotates the user coordinate system in a context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/rotate(by:)
func CGContextRotateCTM(c CGContextRef, angle float64) {
	if _cGContextRotateCTM == nil {
		panic("CoreGraphics: symbol CGContextRotateCTM not loaded")
	}
	_cGContextRotateCTM(c, angle)
}

var _cGContextSaveGState func(c CGContextRef)

// CGContextSaveGState pushes a copy of the current graphics state onto the graphics state stack for the context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/saveGState()
func CGContextSaveGState(c CGContextRef) {
	if _cGContextSaveGState == nil {
		panic("CoreGraphics: symbol CGContextSaveGState not loaded")
	}
	_cGContextSaveGState(c)
}

var _cGContextScaleCTM func(c CGContextRef, sx float64, sy float64)

// CGContextScaleCTM changes the scale of the user coordinate system in a context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/scaleBy(x:y:)
func CGContextScaleCTM(c CGContextRef, sx float64, sy float64) {
	if _cGContextScaleCTM == nil {
		panic("CoreGraphics: symbol CGContextScaleCTM not loaded")
	}
	_cGContextScaleCTM(c, sx, sy)
}

var _cGContextSetAllowsAntialiasing func(c CGContextRef, allowsAntialiasing bool)

// CGContextSetAllowsAntialiasing sets whether or not to allow antialiasing for a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setAllowsAntialiasing(_:)
func CGContextSetAllowsAntialiasing(c CGContextRef, allowsAntialiasing bool) {
	if _cGContextSetAllowsAntialiasing == nil {
		panic("CoreGraphics: symbol CGContextSetAllowsAntialiasing not loaded")
	}
	_cGContextSetAllowsAntialiasing(c, allowsAntialiasing)
}

var _cGContextSetAllowsFontSmoothing func(c CGContextRef, allowsFontSmoothing bool)

// CGContextSetAllowsFontSmoothing sets whether or not to allow font smoothing for a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setAllowsFontSmoothing(_:)
func CGContextSetAllowsFontSmoothing(c CGContextRef, allowsFontSmoothing bool) {
	if _cGContextSetAllowsFontSmoothing == nil {
		panic("CoreGraphics: symbol CGContextSetAllowsFontSmoothing not loaded")
	}
	_cGContextSetAllowsFontSmoothing(c, allowsFontSmoothing)
}

var _cGContextSetAllowsFontSubpixelPositioning func(c CGContextRef, allowsFontSubpixelPositioning bool)

// CGContextSetAllowsFontSubpixelPositioning sets whether or not to allow subpixel positioning for a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setAllowsFontSubpixelPositioning(_:)
func CGContextSetAllowsFontSubpixelPositioning(c CGContextRef, allowsFontSubpixelPositioning bool) {
	if _cGContextSetAllowsFontSubpixelPositioning == nil {
		panic("CoreGraphics: symbol CGContextSetAllowsFontSubpixelPositioning not loaded")
	}
	_cGContextSetAllowsFontSubpixelPositioning(c, allowsFontSubpixelPositioning)
}

var _cGContextSetAllowsFontSubpixelQuantization func(c CGContextRef, allowsFontSubpixelQuantization bool)

// CGContextSetAllowsFontSubpixelQuantization sets whether or not to allow subpixel quantization for a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setAllowsFontSubpixelQuantization(_:)
func CGContextSetAllowsFontSubpixelQuantization(c CGContextRef, allowsFontSubpixelQuantization bool) {
	if _cGContextSetAllowsFontSubpixelQuantization == nil {
		panic("CoreGraphics: symbol CGContextSetAllowsFontSubpixelQuantization not loaded")
	}
	_cGContextSetAllowsFontSubpixelQuantization(c, allowsFontSubpixelQuantization)
}

var _cGContextSetAlpha func(c CGContextRef, alpha float64)

// CGContextSetAlpha sets the opacity level for objects drawn in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setAlpha(_:)
func CGContextSetAlpha(c CGContextRef, alpha float64) {
	if _cGContextSetAlpha == nil {
		panic("CoreGraphics: symbol CGContextSetAlpha not loaded")
	}
	_cGContextSetAlpha(c, alpha)
}

var _cGContextSetBlendMode func(c CGContextRef, mode uintptr)

// CGContextSetBlendMode sets how sample values are composited by a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setBlendMode(_:)
func CGContextSetBlendMode(c CGContextRef, mode uintptr) {
	if _cGContextSetBlendMode == nil {
		panic("CoreGraphics: symbol CGContextSetBlendMode not loaded")
	}
	_cGContextSetBlendMode(c, mode)
}

var _cGContextSetCMYKFillColor func(c CGContextRef, cyan float64, magenta float64, yellow float64, black float64, alpha float64)

// CGContextSetCMYKFillColor sets the current fill color to a value in the DeviceCMYK color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillColor(cyan:magenta:yellow:black:alpha:)
func CGContextSetCMYKFillColor(c CGContextRef, cyan float64, magenta float64, yellow float64, black float64, alpha float64) {
	if _cGContextSetCMYKFillColor == nil {
		panic("CoreGraphics: symbol CGContextSetCMYKFillColor not loaded")
	}
	_cGContextSetCMYKFillColor(c, cyan, magenta, yellow, black, alpha)
}

var _cGContextSetCMYKStrokeColor func(c CGContextRef, cyan float64, magenta float64, yellow float64, black float64, alpha float64)

// CGContextSetCMYKStrokeColor sets the current stroke color to a value in the DeviceCMYK color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokeColor(cyan:magenta:yellow:black:alpha:)
func CGContextSetCMYKStrokeColor(c CGContextRef, cyan float64, magenta float64, yellow float64, black float64, alpha float64) {
	if _cGContextSetCMYKStrokeColor == nil {
		panic("CoreGraphics: symbol CGContextSetCMYKStrokeColor not loaded")
	}
	_cGContextSetCMYKStrokeColor(c, cyan, magenta, yellow, black, alpha)
}

var _cGContextSetCharacterSpacing func(c CGContextRef, spacing float64)

// CGContextSetCharacterSpacing sets the current character spacing.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setCharacterSpacing(_:)
func CGContextSetCharacterSpacing(c CGContextRef, spacing float64) {
	if _cGContextSetCharacterSpacing == nil {
		panic("CoreGraphics: symbol CGContextSetCharacterSpacing not loaded")
	}
	_cGContextSetCharacterSpacing(c, spacing)
}

var _cGContextSetContentToneMappingInfo func(c CGContextRef, info CGContentToneMappingInfo)

// CGContextSetContentToneMappingInfo.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextSetContentToneMappingInfo
func CGContextSetContentToneMappingInfo(c CGContextRef, info CGContentToneMappingInfo) {
	if _cGContextSetContentToneMappingInfo == nil {
		panic("CoreGraphics: symbol CGContextSetContentToneMappingInfo not loaded")
	}
	_cGContextSetContentToneMappingInfo(c, info)
}

var _cGContextSetEDRTargetHeadroom func(c CGContextRef, headroom float32) bool

// CGContextSetEDRTargetHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setEDRTargetHeadroom(_:)
func CGContextSetEDRTargetHeadroom(c CGContextRef, headroom float32) bool {
	if _cGContextSetEDRTargetHeadroom == nil {
		panic("CoreGraphics: symbol CGContextSetEDRTargetHeadroom not loaded")
	}
	return _cGContextSetEDRTargetHeadroom(c, headroom)
}

var _cGContextSetFillColor func(c CGContextRef, components *float64)

// CGContextSetFillColor sets the current fill color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillColor(_:)-756dy
func CGContextSetFillColor(c CGContextRef, components *float64) {
	if _cGContextSetFillColor == nil {
		panic("CoreGraphics: symbol CGContextSetFillColor not loaded")
	}
	_cGContextSetFillColor(c, components)
}

var _cGContextSetFillColorSpace func(c CGContextRef, space CGColorSpaceRef)

// CGContextSetFillColorSpace sets the fill color space in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillColorSpace(_:)
func CGContextSetFillColorSpace(c CGContextRef, space CGColorSpaceRef) {
	if _cGContextSetFillColorSpace == nil {
		panic("CoreGraphics: symbol CGContextSetFillColorSpace not loaded")
	}
	_cGContextSetFillColorSpace(c, space)
}

var _cGContextSetFillColorWithColor func(c CGContextRef, color CGColorRef)

// CGContextSetFillColorWithColor sets the current fill color in a graphics context, using a CGColor.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillColor(_:)-8lhn8
func CGContextSetFillColorWithColor(c CGContextRef, color CGColorRef) {
	if _cGContextSetFillColorWithColor == nil {
		panic("CoreGraphics: symbol CGContextSetFillColorWithColor not loaded")
	}
	_cGContextSetFillColorWithColor(c, color)
}

var _cGContextSetFillPattern func(c CGContextRef, pattern CGPatternRef, components *float64)

// CGContextSetFillPattern sets the fill pattern in the specified graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillPattern(_:colorComponents:)
func CGContextSetFillPattern(c CGContextRef, pattern CGPatternRef, components *float64) {
	if _cGContextSetFillPattern == nil {
		panic("CoreGraphics: symbol CGContextSetFillPattern not loaded")
	}
	_cGContextSetFillPattern(c, pattern, components)
}

var _cGContextSetFlatness func(c CGContextRef, flatness float64)

// CGContextSetFlatness sets the accuracy of curved paths in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFlatness(_:)
func CGContextSetFlatness(c CGContextRef, flatness float64) {
	if _cGContextSetFlatness == nil {
		panic("CoreGraphics: symbol CGContextSetFlatness not loaded")
	}
	_cGContextSetFlatness(c, flatness)
}

var _cGContextSetFont func(c CGContextRef, font CGFontRef)

// CGContextSetFont sets the platform font in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFont(_:)
func CGContextSetFont(c CGContextRef, font CGFontRef) {
	if _cGContextSetFont == nil {
		panic("CoreGraphics: symbol CGContextSetFont not loaded")
	}
	_cGContextSetFont(c, font)
}

var _cGContextSetFontSize func(c CGContextRef, size float64)

// CGContextSetFontSize sets the current font size.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFontSize(_:)
func CGContextSetFontSize(c CGContextRef, size float64) {
	if _cGContextSetFontSize == nil {
		panic("CoreGraphics: symbol CGContextSetFontSize not loaded")
	}
	_cGContextSetFontSize(c, size)
}

var _cGContextSetGrayFillColor func(c CGContextRef, gray float64, alpha float64)

// CGContextSetGrayFillColor sets the current fill color to a value in the DeviceGray color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillColor(gray:alpha:)
func CGContextSetGrayFillColor(c CGContextRef, gray float64, alpha float64) {
	if _cGContextSetGrayFillColor == nil {
		panic("CoreGraphics: symbol CGContextSetGrayFillColor not loaded")
	}
	_cGContextSetGrayFillColor(c, gray, alpha)
}

var _cGContextSetGrayStrokeColor func(c CGContextRef, gray float64, alpha float64)

// CGContextSetGrayStrokeColor sets the current stroke color to a value in the DeviceGray color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokeColor(gray:alpha:)
func CGContextSetGrayStrokeColor(c CGContextRef, gray float64, alpha float64) {
	if _cGContextSetGrayStrokeColor == nil {
		panic("CoreGraphics: symbol CGContextSetGrayStrokeColor not loaded")
	}
	_cGContextSetGrayStrokeColor(c, gray, alpha)
}

var _cGContextSetInterpolationQuality func(c CGContextRef, quality CGInterpolationQuality)

// CGContextSetInterpolationQuality sets the level of interpolation quality for a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextSetInterpolationQuality
func CGContextSetInterpolationQuality(c CGContextRef, quality CGInterpolationQuality) {
	if _cGContextSetInterpolationQuality == nil {
		panic("CoreGraphics: symbol CGContextSetInterpolationQuality not loaded")
	}
	_cGContextSetInterpolationQuality(c, quality)
}

var _cGContextSetLineCap func(c CGContextRef, cap_ uintptr)

// CGContextSetLineCap sets the style for the endpoints of lines drawn in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setLineCap(_:)
func CGContextSetLineCap(c CGContextRef, cap_ uintptr) {
	if _cGContextSetLineCap == nil {
		panic("CoreGraphics: symbol CGContextSetLineCap not loaded")
	}
	_cGContextSetLineCap(c, cap_)
}

var _cGContextSetLineDash func(c CGContextRef, phase float64, lengths *float64, count uintptr)

// CGContextSetLineDash sets the pattern for dashed lines in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextSetLineDash
func CGContextSetLineDash(c CGContextRef, phase float64, lengths *float64, count uintptr) {
	if _cGContextSetLineDash == nil {
		panic("CoreGraphics: symbol CGContextSetLineDash not loaded")
	}
	_cGContextSetLineDash(c, phase, lengths, count)
}

var _cGContextSetLineJoin func(c CGContextRef, join uintptr)

// CGContextSetLineJoin sets the style for the joins of connected lines in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setLineJoin(_:)
func CGContextSetLineJoin(c CGContextRef, join uintptr) {
	if _cGContextSetLineJoin == nil {
		panic("CoreGraphics: symbol CGContextSetLineJoin not loaded")
	}
	_cGContextSetLineJoin(c, join)
}

var _cGContextSetLineWidth func(c CGContextRef, width float64)

// CGContextSetLineWidth sets the line width for a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setLineWidth(_:)
func CGContextSetLineWidth(c CGContextRef, width float64) {
	if _cGContextSetLineWidth == nil {
		panic("CoreGraphics: symbol CGContextSetLineWidth not loaded")
	}
	_cGContextSetLineWidth(c, width)
}

var _cGContextSetMiterLimit func(c CGContextRef, limit float64)

// CGContextSetMiterLimit sets the miter limit for the joins of connected lines in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setMiterLimit(_:)
func CGContextSetMiterLimit(c CGContextRef, limit float64) {
	if _cGContextSetMiterLimit == nil {
		panic("CoreGraphics: symbol CGContextSetMiterLimit not loaded")
	}
	_cGContextSetMiterLimit(c, limit)
}

var _cGContextSetPatternPhase func(c CGContextRef, phase corefoundation.CGSize)

// CGContextSetPatternPhase sets the pattern phase of a context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setPatternPhase(_:)
func CGContextSetPatternPhase(c CGContextRef, phase corefoundation.CGSize) {
	if _cGContextSetPatternPhase == nil {
		panic("CoreGraphics: symbol CGContextSetPatternPhase not loaded")
	}
	_cGContextSetPatternPhase(c, phase)
}

var _cGContextSetRGBFillColor func(c CGContextRef, red float64, green float64, blue float64, alpha float64)

// CGContextSetRGBFillColor sets the current fill color to a value in the DeviceRGB color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillColor(red:green:blue:alpha:)
func CGContextSetRGBFillColor(c CGContextRef, red float64, green float64, blue float64, alpha float64) {
	if _cGContextSetRGBFillColor == nil {
		panic("CoreGraphics: symbol CGContextSetRGBFillColor not loaded")
	}
	_cGContextSetRGBFillColor(c, red, green, blue, alpha)
}

var _cGContextSetRGBStrokeColor func(c CGContextRef, red float64, green float64, blue float64, alpha float64)

// CGContextSetRGBStrokeColor sets the current stroke color to a value in the DeviceRGB color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokeColor(red:green:blue:alpha:)
func CGContextSetRGBStrokeColor(c CGContextRef, red float64, green float64, blue float64, alpha float64) {
	if _cGContextSetRGBStrokeColor == nil {
		panic("CoreGraphics: symbol CGContextSetRGBStrokeColor not loaded")
	}
	_cGContextSetRGBStrokeColor(c, red, green, blue, alpha)
}

var _cGContextSetRenderingIntent func(c CGContextRef, intent CGColorRenderingIntent)

// CGContextSetRenderingIntent sets the rendering intent in the current graphics state.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setRenderingIntent(_:)
func CGContextSetRenderingIntent(c CGContextRef, intent CGColorRenderingIntent) {
	if _cGContextSetRenderingIntent == nil {
		panic("CoreGraphics: symbol CGContextSetRenderingIntent not loaded")
	}
	_cGContextSetRenderingIntent(c, intent)
}

var _cGContextSetShadow func(c CGContextRef, offset corefoundation.CGSize, blur float64)

// CGContextSetShadow enables shadowing in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setShadow(offset:blur:)
func CGContextSetShadow(c CGContextRef, offset corefoundation.CGSize, blur float64) {
	if _cGContextSetShadow == nil {
		panic("CoreGraphics: symbol CGContextSetShadow not loaded")
	}
	_cGContextSetShadow(c, offset, blur)
}

var _cGContextSetShadowWithColor func(c CGContextRef, offset corefoundation.CGSize, blur float64, color CGColorRef)

// CGContextSetShadowWithColor enables shadowing with color a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setShadow(offset:blur:color:)
func CGContextSetShadowWithColor(c CGContextRef, offset corefoundation.CGSize, blur float64, color CGColorRef) {
	if _cGContextSetShadowWithColor == nil {
		panic("CoreGraphics: symbol CGContextSetShadowWithColor not loaded")
	}
	_cGContextSetShadowWithColor(c, offset, blur, color)
}

var _cGContextSetShouldAntialias func(c CGContextRef, shouldAntialias bool)

// CGContextSetShouldAntialias sets antialiasing on or off for a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setShouldAntialias(_:)
func CGContextSetShouldAntialias(c CGContextRef, shouldAntialias bool) {
	if _cGContextSetShouldAntialias == nil {
		panic("CoreGraphics: symbol CGContextSetShouldAntialias not loaded")
	}
	_cGContextSetShouldAntialias(c, shouldAntialias)
}

var _cGContextSetShouldSmoothFonts func(c CGContextRef, shouldSmoothFonts bool)

// CGContextSetShouldSmoothFonts enables or disables font smoothing in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setShouldSmoothFonts(_:)
func CGContextSetShouldSmoothFonts(c CGContextRef, shouldSmoothFonts bool) {
	if _cGContextSetShouldSmoothFonts == nil {
		panic("CoreGraphics: symbol CGContextSetShouldSmoothFonts not loaded")
	}
	_cGContextSetShouldSmoothFonts(c, shouldSmoothFonts)
}

var _cGContextSetShouldSubpixelPositionFonts func(c CGContextRef, shouldSubpixelPositionFonts bool)

// CGContextSetShouldSubpixelPositionFonts enables or disables subpixel positioning in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setShouldSubpixelPositionFonts(_:)
func CGContextSetShouldSubpixelPositionFonts(c CGContextRef, shouldSubpixelPositionFonts bool) {
	if _cGContextSetShouldSubpixelPositionFonts == nil {
		panic("CoreGraphics: symbol CGContextSetShouldSubpixelPositionFonts not loaded")
	}
	_cGContextSetShouldSubpixelPositionFonts(c, shouldSubpixelPositionFonts)
}

var _cGContextSetShouldSubpixelQuantizeFonts func(c CGContextRef, shouldSubpixelQuantizeFonts bool)

// CGContextSetShouldSubpixelQuantizeFonts enables or disables subpixel quantization in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setShouldSubpixelQuantizeFonts(_:)
func CGContextSetShouldSubpixelQuantizeFonts(c CGContextRef, shouldSubpixelQuantizeFonts bool) {
	if _cGContextSetShouldSubpixelQuantizeFonts == nil {
		panic("CoreGraphics: symbol CGContextSetShouldSubpixelQuantizeFonts not loaded")
	}
	_cGContextSetShouldSubpixelQuantizeFonts(c, shouldSubpixelQuantizeFonts)
}

var _cGContextSetStrokeColor func(c CGContextRef, components *float64)

// CGContextSetStrokeColor sets the current stroke color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokeColor(_:)-4pd8p
func CGContextSetStrokeColor(c CGContextRef, components *float64) {
	if _cGContextSetStrokeColor == nil {
		panic("CoreGraphics: symbol CGContextSetStrokeColor not loaded")
	}
	_cGContextSetStrokeColor(c, components)
}

var _cGContextSetStrokeColorSpace func(c CGContextRef, space CGColorSpaceRef)

// CGContextSetStrokeColorSpace sets the stroke color space in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokeColorSpace(_:)
func CGContextSetStrokeColorSpace(c CGContextRef, space CGColorSpaceRef) {
	if _cGContextSetStrokeColorSpace == nil {
		panic("CoreGraphics: symbol CGContextSetStrokeColorSpace not loaded")
	}
	_cGContextSetStrokeColorSpace(c, space)
}

var _cGContextSetStrokeColorWithColor func(c CGContextRef, color CGColorRef)

// CGContextSetStrokeColorWithColor sets the current stroke color in a context, using a CGColor.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokeColor(_:)-1sskg
func CGContextSetStrokeColorWithColor(c CGContextRef, color CGColorRef) {
	if _cGContextSetStrokeColorWithColor == nil {
		panic("CoreGraphics: symbol CGContextSetStrokeColorWithColor not loaded")
	}
	_cGContextSetStrokeColorWithColor(c, color)
}

var _cGContextSetStrokePattern func(c CGContextRef, pattern CGPatternRef, components *float64)

// CGContextSetStrokePattern sets the stroke pattern in the specified graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokePattern(_:colorComponents:)
func CGContextSetStrokePattern(c CGContextRef, pattern CGPatternRef, components *float64) {
	if _cGContextSetStrokePattern == nil {
		panic("CoreGraphics: symbol CGContextSetStrokePattern not loaded")
	}
	_cGContextSetStrokePattern(c, pattern, components)
}

var _cGContextSetTextDrawingMode func(c CGContextRef, mode CGTextDrawingMode)

// CGContextSetTextDrawingMode sets the current text drawing mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setTextDrawingMode(_:)
func CGContextSetTextDrawingMode(c CGContextRef, mode CGTextDrawingMode) {
	if _cGContextSetTextDrawingMode == nil {
		panic("CoreGraphics: symbol CGContextSetTextDrawingMode not loaded")
	}
	_cGContextSetTextDrawingMode(c, mode)
}

var _cGContextSetTextMatrix func(c CGContextRef, t corefoundation.CGAffineTransform)

// CGContextSetTextMatrix sets the current text matrix.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextSetTextMatrix
func CGContextSetTextMatrix(c CGContextRef, t corefoundation.CGAffineTransform) {
	if _cGContextSetTextMatrix == nil {
		panic("CoreGraphics: symbol CGContextSetTextMatrix not loaded")
	}
	_cGContextSetTextMatrix(c, t)
}

var _cGContextSetTextPosition func(c CGContextRef, x float64, y float64)

// CGContextSetTextPosition sets the location at which text is drawn.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextSetTextPosition
func CGContextSetTextPosition(c CGContextRef, x float64, y float64) {
	if _cGContextSetTextPosition == nil {
		panic("CoreGraphics: symbol CGContextSetTextPosition not loaded")
	}
	_cGContextSetTextPosition(c, x, y)
}

var _cGContextShowGlyphsAtPositions func(c CGContextRef, glyphs uintptr, Lpositions *corefoundation.CGPoint, count uintptr)

// CGContextShowGlyphsAtPositions draws glyphs at the provided position.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextShowGlyphsAtPositions
func CGContextShowGlyphsAtPositions(c CGContextRef, glyphs uintptr, Lpositions *corefoundation.CGPoint, count uintptr) {
	if _cGContextShowGlyphsAtPositions == nil {
		panic("CoreGraphics: symbol CGContextShowGlyphsAtPositions not loaded")
	}
	_cGContextShowGlyphsAtPositions(c, glyphs, Lpositions, count)
}

var _cGContextStrokeEllipseInRect func(c CGContextRef, rect corefoundation.CGRect)

// CGContextStrokeEllipseInRect strokes an ellipse that fits inside the specified rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/strokeEllipse(in:)
func CGContextStrokeEllipseInRect(c CGContextRef, rect corefoundation.CGRect) {
	if _cGContextStrokeEllipseInRect == nil {
		panic("CoreGraphics: symbol CGContextStrokeEllipseInRect not loaded")
	}
	_cGContextStrokeEllipseInRect(c, rect)
}

var _cGContextStrokeLineSegments func(c CGContextRef, points *corefoundation.CGPoint, count uintptr)

// CGContextStrokeLineSegments strokes a sequence of line segments.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextStrokeLineSegments
func CGContextStrokeLineSegments(c CGContextRef, points *corefoundation.CGPoint, count uintptr) {
	if _cGContextStrokeLineSegments == nil {
		panic("CoreGraphics: symbol CGContextStrokeLineSegments not loaded")
	}
	_cGContextStrokeLineSegments(c, points, count)
}

var _cGContextStrokePath func(c CGContextRef)

// CGContextStrokePath paints a line along the current path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/strokePath()
func CGContextStrokePath(c CGContextRef) {
	if _cGContextStrokePath == nil {
		panic("CoreGraphics: symbol CGContextStrokePath not loaded")
	}
	_cGContextStrokePath(c)
}

var _cGContextStrokeRect func(c CGContextRef, rect corefoundation.CGRect)

// CGContextStrokeRect paints a rectangular path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/stroke(_:)
func CGContextStrokeRect(c CGContextRef, rect corefoundation.CGRect) {
	if _cGContextStrokeRect == nil {
		panic("CoreGraphics: symbol CGContextStrokeRect not loaded")
	}
	_cGContextStrokeRect(c, rect)
}

var _cGContextStrokeRectWithWidth func(c CGContextRef, rect corefoundation.CGRect, width float64)

// CGContextStrokeRectWithWidth paints a rectangular path, using the specified line width.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/stroke(_:width:)
func CGContextStrokeRectWithWidth(c CGContextRef, rect corefoundation.CGRect, width float64) {
	if _cGContextStrokeRectWithWidth == nil {
		panic("CoreGraphics: symbol CGContextStrokeRectWithWidth not loaded")
	}
	_cGContextStrokeRectWithWidth(c, rect, width)
}

var _cGContextSynchronize func(c CGContextRef)

// CGContextSynchronize marks a window context for update.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/synchronize()
func CGContextSynchronize(c CGContextRef) {
	if _cGContextSynchronize == nil {
		panic("CoreGraphics: symbol CGContextSynchronize not loaded")
	}
	_cGContextSynchronize(c)
}

var _cGContextSynchronizeAttributes func(c CGContextRef)

// CGContextSynchronizeAttributes.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/synchronizeAttributes()
func CGContextSynchronizeAttributes(c CGContextRef) {
	if _cGContextSynchronizeAttributes == nil {
		panic("CoreGraphics: symbol CGContextSynchronizeAttributes not loaded")
	}
	_cGContextSynchronizeAttributes(c)
}

var _cGContextTranslateCTM func(c CGContextRef, tx float64, ty float64)

// CGContextTranslateCTM changes the origin of the user coordinate system in a context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/translateBy(x:y:)
func CGContextTranslateCTM(c CGContextRef, tx float64, ty float64) {
	if _cGContextTranslateCTM == nil {
		panic("CoreGraphics: symbol CGContextTranslateCTM not loaded")
	}
	_cGContextTranslateCTM(c, tx, ty)
}

var _cGConvertColorDataWithFormat func(width uintptr, height uintptr, dst_data unsafe.Pointer, dst_format CGColorDataFormat, src_data unsafe.Pointer, src_format CGColorDataFormat, options corefoundation.CFDictionaryRef) bool

// CGConvertColorDataWithFormat.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGConvertColorDataWithFormat(_:_:_:_:_:_:_:)
func CGConvertColorDataWithFormat(width uintptr, height uintptr, dst_data unsafe.Pointer, dst_format CGColorDataFormat, src_data unsafe.Pointer, src_format CGColorDataFormat, options corefoundation.CFDictionaryRef) bool {
	if _cGConvertColorDataWithFormat == nil {
		panic("CoreGraphics: symbol CGConvertColorDataWithFormat not loaded")
	}
	return _cGConvertColorDataWithFormat(width, height, dst_data, dst_format, src_data, src_format, options)
}

var _cGDataConsumerCreate func(info unsafe.Pointer, cbks *CGDataConsumerCallbacks) CGDataConsumerRef

// CGDataConsumerCreate creates a data consumer that uses callback functions to write data.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumer/init(info:cbks:)
func CGDataConsumerCreate(info unsafe.Pointer, cbks *CGDataConsumerCallbacks) CGDataConsumerRef {
	if _cGDataConsumerCreate == nil {
		panic("CoreGraphics: symbol CGDataConsumerCreate not loaded")
	}
	return _cGDataConsumerCreate(info, cbks)
}

var _cGDataConsumerCreateWithCFData func(data corefoundation.CFMutableDataRef) CGDataConsumerRef

// CGDataConsumerCreateWithCFData creates a data consumer that writes to a CFData object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumer/init(data:)
func CGDataConsumerCreateWithCFData(data corefoundation.CFMutableDataRef) CGDataConsumerRef {
	if _cGDataConsumerCreateWithCFData == nil {
		panic("CoreGraphics: symbol CGDataConsumerCreateWithCFData not loaded")
	}
	return _cGDataConsumerCreateWithCFData(data)
}

var _cGDataConsumerCreateWithURL func(url corefoundation.CFURLRef) CGDataConsumerRef

// CGDataConsumerCreateWithURL creates a data consumer that writes data to a location specified by a URL.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumer/init(url:)
func CGDataConsumerCreateWithURL(url corefoundation.CFURLRef) CGDataConsumerRef {
	if _cGDataConsumerCreateWithURL == nil {
		panic("CoreGraphics: symbol CGDataConsumerCreateWithURL not loaded")
	}
	return _cGDataConsumerCreateWithURL(url)
}

var _cGDataConsumerGetTypeID func() uint

// CGDataConsumerGetTypeID returns the Core Foundation type identifier for Core Graphics data consumers.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumer/typeID
func CGDataConsumerGetTypeID() uint {
	if _cGDataConsumerGetTypeID == nil {
		panic("CoreGraphics: symbol CGDataConsumerGetTypeID not loaded")
	}
	return _cGDataConsumerGetTypeID()
}

var _cGDataConsumerRelease func(consumer CGDataConsumerRef)

// CGDataConsumerRelease decrements the retain count of a data consumer.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumerRelease
func CGDataConsumerRelease(consumer CGDataConsumerRef) {
	if _cGDataConsumerRelease == nil {
		panic("CoreGraphics: symbol CGDataConsumerRelease not loaded")
	}
	_cGDataConsumerRelease(consumer)
}

var _cGDataConsumerRetain func(consumer CGDataConsumerRef) CGDataConsumerRef

// CGDataConsumerRetain increments the retain count of a data consumer.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumerRetain
func CGDataConsumerRetain(consumer CGDataConsumerRef) CGDataConsumerRef {
	if _cGDataConsumerRetain == nil {
		panic("CoreGraphics: symbol CGDataConsumerRetain not loaded")
	}
	return _cGDataConsumerRetain(consumer)
}

var _cGDataProviderCopyData func(provider CGDataProviderRef) corefoundation.CFDataRef

// CGDataProviderCopyData returns a copy of the provider’s data.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/data
func CGDataProviderCopyData(provider CGDataProviderRef) corefoundation.CFDataRef {
	if _cGDataProviderCopyData == nil {
		panic("CoreGraphics: symbol CGDataProviderCopyData not loaded")
	}
	return _cGDataProviderCopyData(provider)
}

var _cGDataProviderCreateDirect func(info unsafe.Pointer, size int64, callbacks *CGDataProviderDirectCallbacks) CGDataProviderRef

// CGDataProviderCreateDirect creates a direct-access data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/init(directInfo:size:callbacks:)
func CGDataProviderCreateDirect(info unsafe.Pointer, size int64, callbacks *CGDataProviderDirectCallbacks) CGDataProviderRef {
	if _cGDataProviderCreateDirect == nil {
		panic("CoreGraphics: symbol CGDataProviderCreateDirect not loaded")
	}
	return _cGDataProviderCreateDirect(info, size, callbacks)
}

var _cGDataProviderCreateSequential func(info unsafe.Pointer, callbacks *CGDataProviderSequentialCallbacks) CGDataProviderRef

// CGDataProviderCreateSequential creates a sequential-access data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/init(sequentialInfo:callbacks:)
func CGDataProviderCreateSequential(info unsafe.Pointer, callbacks *CGDataProviderSequentialCallbacks) CGDataProviderRef {
	if _cGDataProviderCreateSequential == nil {
		panic("CoreGraphics: symbol CGDataProviderCreateSequential not loaded")
	}
	return _cGDataProviderCreateSequential(info, callbacks)
}

var _cGDataProviderCreateWithCFData func(data corefoundation.CFDataRef) CGDataProviderRef

// CGDataProviderCreateWithCFData creates a data provider that reads from a CFData object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/init(data:)
func CGDataProviderCreateWithCFData(data corefoundation.CFDataRef) CGDataProviderRef {
	if _cGDataProviderCreateWithCFData == nil {
		panic("CoreGraphics: symbol CGDataProviderCreateWithCFData not loaded")
	}
	return _cGDataProviderCreateWithCFData(data)
}

var _cGDataProviderCreateWithData func(info unsafe.Pointer, data unsafe.Pointer, size uintptr, releaseData CGDataProviderReleaseDataCallback) CGDataProviderRef

// CGDataProviderCreateWithData creates a direct-access data provider that uses data your program supplies.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/init(dataInfo:data:size:releaseData:)
func CGDataProviderCreateWithData(info unsafe.Pointer, data unsafe.Pointer, size uintptr, releaseData CGDataProviderReleaseDataCallback) CGDataProviderRef {
	if _cGDataProviderCreateWithData == nil {
		panic("CoreGraphics: symbol CGDataProviderCreateWithData not loaded")
	}
	return _cGDataProviderCreateWithData(info, data, size, releaseData)
}

var _cGDataProviderCreateWithFilename func(filename string) CGDataProviderRef

// CGDataProviderCreateWithFilename creates a direct-access data provider that uses a file to supply data.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/init(filename:)
func CGDataProviderCreateWithFilename(filename string) CGDataProviderRef {
	if _cGDataProviderCreateWithFilename == nil {
		panic("CoreGraphics: symbol CGDataProviderCreateWithFilename not loaded")
	}
	return _cGDataProviderCreateWithFilename(filename)
}

var _cGDataProviderCreateWithURL func(url corefoundation.CFURLRef) CGDataProviderRef

// CGDataProviderCreateWithURL creates a direct-access data provider that uses a URL to supply data.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/init(url:)
func CGDataProviderCreateWithURL(url corefoundation.CFURLRef) CGDataProviderRef {
	if _cGDataProviderCreateWithURL == nil {
		panic("CoreGraphics: symbol CGDataProviderCreateWithURL not loaded")
	}
	return _cGDataProviderCreateWithURL(url)
}

var _cGDataProviderGetInfo func(provider CGDataProviderRef) unsafe.Pointer

// CGDataProviderGetInfo.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/info
func CGDataProviderGetInfo(provider CGDataProviderRef) unsafe.Pointer {
	if _cGDataProviderGetInfo == nil {
		panic("CoreGraphics: symbol CGDataProviderGetInfo not loaded")
	}
	return _cGDataProviderGetInfo(provider)
}

var _cGDataProviderGetTypeID func() uint

// CGDataProviderGetTypeID returns the Core Foundation type identifier for data providers.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/typeID
func CGDataProviderGetTypeID() uint {
	if _cGDataProviderGetTypeID == nil {
		panic("CoreGraphics: symbol CGDataProviderGetTypeID not loaded")
	}
	return _cGDataProviderGetTypeID()
}

var _cGDataProviderRelease func(provider CGDataProviderRef)

// CGDataProviderRelease decrements the retain count of a data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProviderRelease
func CGDataProviderRelease(provider CGDataProviderRef) {
	if _cGDataProviderRelease == nil {
		panic("CoreGraphics: symbol CGDataProviderRelease not loaded")
	}
	_cGDataProviderRelease(provider)
}

var _cGDataProviderRetain func(provider CGDataProviderRef) CGDataProviderRef

// CGDataProviderRetain increments the retain count of a data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProviderRetain
func CGDataProviderRetain(provider CGDataProviderRef) CGDataProviderRef {
	if _cGDataProviderRetain == nil {
		panic("CoreGraphics: symbol CGDataProviderRetain not loaded")
	}
	return _cGDataProviderRetain(provider)
}

var _cGDirectDisplayCopyCurrentMetalDevice func(display uint32) unsafe.Pointer

// CGDirectDisplayCopyCurrentMetalDevice returns the GPU device instance that’s currently driving a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDirectDisplayCopyCurrentMetalDevice(_:)
func CGDirectDisplayCopyCurrentMetalDevice(display uint32) objectivec.IObject {
	if _cGDirectDisplayCopyCurrentMetalDevice == nil {
		panic("CoreGraphics: symbol CGDirectDisplayCopyCurrentMetalDevice not loaded")
	}
	rv := _cGDirectDisplayCopyCurrentMetalDevice(display)
	return objectivec.ObjectFromID(objc.IDFrom(rv))
}

var _cGDisplayBounds func(display uint32) corefoundation.CGRect

// CGDisplayBounds returns the bounds of a display in the global display coordinate space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayBounds(_:)
func CGDisplayBounds(display uint32) corefoundation.CGRect {
	if _cGDisplayBounds == nil {
		panic("CoreGraphics: symbol CGDisplayBounds not loaded")
	}
	return _cGDisplayBounds(display)
}

var _cGDisplayCapture func(display uint32) CGError

// CGDisplayCapture obtains exclusive use of a display, preventing other applications and system services from using the display or changing its configuration.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCapture(_:)
func CGDisplayCapture(display uint32) CGError {
	if _cGDisplayCapture == nil {
		panic("CoreGraphics: symbol CGDisplayCapture not loaded")
	}
	return _cGDisplayCapture(display)
}

var _cGDisplayCaptureWithOptions func(display uint32, options CGCaptureOptions) CGError

// CGDisplayCaptureWithOptions obtains exclusive use of a display for an application using the options you specify.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCaptureWithOptions(_:_:)
func CGDisplayCaptureWithOptions(display uint32, options CGCaptureOptions) CGError {
	if _cGDisplayCaptureWithOptions == nil {
		panic("CoreGraphics: symbol CGDisplayCaptureWithOptions not loaded")
	}
	return _cGDisplayCaptureWithOptions(display, options)
}

var _cGDisplayCopyAllDisplayModes func(display uint32, options corefoundation.CFDictionaryRef) corefoundation.CFArrayRef

// CGDisplayCopyAllDisplayModes returns information about the currently available display modes.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCopyAllDisplayModes(_:_:)
func CGDisplayCopyAllDisplayModes(display uint32, options corefoundation.CFDictionaryRef) corefoundation.CFArrayRef {
	if _cGDisplayCopyAllDisplayModes == nil {
		panic("CoreGraphics: symbol CGDisplayCopyAllDisplayModes not loaded")
	}
	return _cGDisplayCopyAllDisplayModes(display, options)
}

var _cGDisplayCopyColorSpace func(display uint32) CGColorSpaceRef

// CGDisplayCopyColorSpace returns the color space for a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCopyColorSpace(_:)
func CGDisplayCopyColorSpace(display uint32) CGColorSpaceRef {
	if _cGDisplayCopyColorSpace == nil {
		panic("CoreGraphics: symbol CGDisplayCopyColorSpace not loaded")
	}
	return _cGDisplayCopyColorSpace(display)
}

var _cGDisplayCopyDisplayMode func(display uint32) CGDisplayModeRef

// CGDisplayCopyDisplayMode returns information about a display’s current configuration.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCopyDisplayMode(_:)
func CGDisplayCopyDisplayMode(display uint32) CGDisplayModeRef {
	if _cGDisplayCopyDisplayMode == nil {
		panic("CoreGraphics: symbol CGDisplayCopyDisplayMode not loaded")
	}
	return _cGDisplayCopyDisplayMode(display)
}

var _cGDisplayCreateImage func(displayID uint32) CGImageRef

// CGDisplayCreateImage returns an image containing the contents of the specified display.
//
// Deprecated: Please use ScreenCaptureKit instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCreateImage(_:)
func CGDisplayCreateImage(displayID uint32) CGImageRef {
	if _cGDisplayCreateImage == nil {
		panic("CoreGraphics: symbol CGDisplayCreateImage not loaded")
	}
	return _cGDisplayCreateImage(displayID)
}

var _cGDisplayCreateImageForRect func(display uint32, rect corefoundation.CGRect) CGImageRef

// CGDisplayCreateImageForRect returns an image containing the contents of a portion of the specified display.
//
// Deprecated: Please use ScreenCaptureKit instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCreateImage(_:rect:)
func CGDisplayCreateImageForRect(display uint32, rect corefoundation.CGRect) CGImageRef {
	if _cGDisplayCreateImageForRect == nil {
		panic("CoreGraphics: symbol CGDisplayCreateImageForRect not loaded")
	}
	return _cGDisplayCreateImageForRect(display, rect)
}

var _cGDisplayFade func(token CGDisplayFadeReservationToken, duration float32, startBlend CGDisplayBlendFraction, endBlend CGDisplayBlendFraction, redBlend float32, greenBlend float32, blueBlend float32, synchronous uintptr) CGError

// CGDisplayFade performs a single fade operation.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayFade(_:_:_:_:_:_:_:_:)
func CGDisplayFade(token CGDisplayFadeReservationToken, duration float32, startBlend CGDisplayBlendFraction, endBlend CGDisplayBlendFraction, redBlend float32, greenBlend float32, blueBlend float32, synchronous uintptr) CGError {
	if _cGDisplayFade == nil {
		panic("CoreGraphics: symbol CGDisplayFade not loaded")
	}
	return _cGDisplayFade(token, duration, startBlend, endBlend, redBlend, greenBlend, blueBlend, synchronous)
}

var _cGDisplayGammaTableCapacity func(display uint32) uint32

// CGDisplayGammaTableCapacity returns the capacity, or number of entries, in the gamma table for a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayGammaTableCapacity(_:)
func CGDisplayGammaTableCapacity(display uint32) uint32 {
	if _cGDisplayGammaTableCapacity == nil {
		panic("CoreGraphics: symbol CGDisplayGammaTableCapacity not loaded")
	}
	return _cGDisplayGammaTableCapacity(display)
}

var _cGDisplayGetDrawingContext func(display uint32) CGContextRef

// CGDisplayGetDrawingContext returns a graphics context suitable for drawing to a captured display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayGetDrawingContext(_:)
func CGDisplayGetDrawingContext(display uint32) CGContextRef {
	if _cGDisplayGetDrawingContext == nil {
		panic("CoreGraphics: symbol CGDisplayGetDrawingContext not loaded")
	}
	return _cGDisplayGetDrawingContext(display)
}

var _cGDisplayHideCursor func(display uint32) CGError

// CGDisplayHideCursor hides the mouse cursor, and increments the hide cursor count.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayHideCursor(_:)
func CGDisplayHideCursor(display uint32) CGError {
	if _cGDisplayHideCursor == nil {
		panic("CoreGraphics: symbol CGDisplayHideCursor not loaded")
	}
	return _cGDisplayHideCursor(display)
}

var _cGDisplayIDToOpenGLDisplayMask func(display uint32) CGOpenGLDisplayMask

// CGDisplayIDToOpenGLDisplayMask maps a display ID to an OpenGL display mask.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIDToOpenGLDisplayMask(_:)
func CGDisplayIDToOpenGLDisplayMask(display uint32) CGOpenGLDisplayMask {
	if _cGDisplayIDToOpenGLDisplayMask == nil {
		panic("CoreGraphics: symbol CGDisplayIDToOpenGLDisplayMask not loaded")
	}
	return _cGDisplayIDToOpenGLDisplayMask(display)
}

var _cGDisplayIsActive func(display uint32) objectivec.IObject

// CGDisplayIsActive returns a Boolean value indicating whether a display is active.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsActive(_:)
func CGDisplayIsActive(display uint32) objectivec.IObject {
	if _cGDisplayIsActive == nil {
		panic("CoreGraphics: symbol CGDisplayIsActive not loaded")
	}
	return _cGDisplayIsActive(display)
}

var _cGDisplayIsAlwaysInMirrorSet func(display uint32) objectivec.IObject

// CGDisplayIsAlwaysInMirrorSet returns a Boolean value indicating whether a display is always in a mirroring set.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsAlwaysInMirrorSet(_:)
func CGDisplayIsAlwaysInMirrorSet(display uint32) objectivec.IObject {
	if _cGDisplayIsAlwaysInMirrorSet == nil {
		panic("CoreGraphics: symbol CGDisplayIsAlwaysInMirrorSet not loaded")
	}
	return _cGDisplayIsAlwaysInMirrorSet(display)
}

var _cGDisplayIsAsleep func(display uint32) objectivec.IObject

// CGDisplayIsAsleep returns a Boolean value indicating whether a display is sleeping (and is therefore not drawable).
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsAsleep(_:)
func CGDisplayIsAsleep(display uint32) objectivec.IObject {
	if _cGDisplayIsAsleep == nil {
		panic("CoreGraphics: symbol CGDisplayIsAsleep not loaded")
	}
	return _cGDisplayIsAsleep(display)
}

var _cGDisplayIsBuiltin func(display uint32) objectivec.IObject

// CGDisplayIsBuiltin returns a Boolean value indicating whether a display is built-in, such as the internal display in portable systems.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsBuiltin(_:)
func CGDisplayIsBuiltin(display uint32) objectivec.IObject {
	if _cGDisplayIsBuiltin == nil {
		panic("CoreGraphics: symbol CGDisplayIsBuiltin not loaded")
	}
	return _cGDisplayIsBuiltin(display)
}

var _cGDisplayIsInHWMirrorSet func(display uint32) objectivec.IObject

// CGDisplayIsInHWMirrorSet returns a Boolean value indicating whether a display is in a hardware mirroring set.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsInHWMirrorSet(_:)
func CGDisplayIsInHWMirrorSet(display uint32) objectivec.IObject {
	if _cGDisplayIsInHWMirrorSet == nil {
		panic("CoreGraphics: symbol CGDisplayIsInHWMirrorSet not loaded")
	}
	return _cGDisplayIsInHWMirrorSet(display)
}

var _cGDisplayIsInMirrorSet func(display uint32) objectivec.IObject

// CGDisplayIsInMirrorSet returns a Boolean value indicating whether a display is in a mirroring set.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsInMirrorSet(_:)
func CGDisplayIsInMirrorSet(display uint32) objectivec.IObject {
	if _cGDisplayIsInMirrorSet == nil {
		panic("CoreGraphics: symbol CGDisplayIsInMirrorSet not loaded")
	}
	return _cGDisplayIsInMirrorSet(display)
}

var _cGDisplayIsMain func(display uint32) objectivec.IObject

// CGDisplayIsMain returns a Boolean value indicating whether a display is the main display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsMain(_:)
func CGDisplayIsMain(display uint32) objectivec.IObject {
	if _cGDisplayIsMain == nil {
		panic("CoreGraphics: symbol CGDisplayIsMain not loaded")
	}
	return _cGDisplayIsMain(display)
}

var _cGDisplayIsOnline func(display uint32) objectivec.IObject

// CGDisplayIsOnline returns a Boolean value indicating whether a display is connected or online.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsOnline(_:)
func CGDisplayIsOnline(display uint32) objectivec.IObject {
	if _cGDisplayIsOnline == nil {
		panic("CoreGraphics: symbol CGDisplayIsOnline not loaded")
	}
	return _cGDisplayIsOnline(display)
}

var _cGDisplayIsStereo func(display uint32) objectivec.IObject

// CGDisplayIsStereo returns a Boolean value indicating whether a display is running in a stereo graphics mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsStereo(_:)
func CGDisplayIsStereo(display uint32) objectivec.IObject {
	if _cGDisplayIsStereo == nil {
		panic("CoreGraphics: symbol CGDisplayIsStereo not loaded")
	}
	return _cGDisplayIsStereo(display)
}

var _cGDisplayMirrorsDisplay func(display uint32) uint32

// CGDisplayMirrorsDisplay for a secondary display in a mirroring set, returns the primary display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMirrorsDisplay(_:)
func CGDisplayMirrorsDisplay(display uint32) uint32 {
	if _cGDisplayMirrorsDisplay == nil {
		panic("CoreGraphics: symbol CGDisplayMirrorsDisplay not loaded")
	}
	return _cGDisplayMirrorsDisplay(display)
}

var _cGDisplayModeGetHeight func(mode CGDisplayModeRef) uintptr

// CGDisplayModeGetHeight returns the height of the specified display mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/height
func CGDisplayModeGetHeight(mode CGDisplayModeRef) uintptr {
	if _cGDisplayModeGetHeight == nil {
		panic("CoreGraphics: symbol CGDisplayModeGetHeight not loaded")
	}
	return _cGDisplayModeGetHeight(mode)
}

var _cGDisplayModeGetIODisplayModeID func(mode CGDisplayModeRef) int32

// CGDisplayModeGetIODisplayModeID returns the I/O Kit display mode ID of the specified display mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/ioDisplayModeID
func CGDisplayModeGetIODisplayModeID(mode CGDisplayModeRef) int32 {
	if _cGDisplayModeGetIODisplayModeID == nil {
		panic("CoreGraphics: symbol CGDisplayModeGetIODisplayModeID not loaded")
	}
	return _cGDisplayModeGetIODisplayModeID(mode)
}

var _cGDisplayModeGetIOFlags func(mode CGDisplayModeRef) uint32

// CGDisplayModeGetIOFlags returns the I/O Kit flags of the specified display mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/ioFlags
func CGDisplayModeGetIOFlags(mode CGDisplayModeRef) uint32 {
	if _cGDisplayModeGetIOFlags == nil {
		panic("CoreGraphics: symbol CGDisplayModeGetIOFlags not loaded")
	}
	return _cGDisplayModeGetIOFlags(mode)
}

var _cGDisplayModeGetPixelHeight func(mode CGDisplayModeRef) uintptr

// CGDisplayModeGetPixelHeight.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/pixelHeight
func CGDisplayModeGetPixelHeight(mode CGDisplayModeRef) uintptr {
	if _cGDisplayModeGetPixelHeight == nil {
		panic("CoreGraphics: symbol CGDisplayModeGetPixelHeight not loaded")
	}
	return _cGDisplayModeGetPixelHeight(mode)
}

var _cGDisplayModeGetPixelWidth func(mode CGDisplayModeRef) uintptr

// CGDisplayModeGetPixelWidth.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/pixelWidth
func CGDisplayModeGetPixelWidth(mode CGDisplayModeRef) uintptr {
	if _cGDisplayModeGetPixelWidth == nil {
		panic("CoreGraphics: symbol CGDisplayModeGetPixelWidth not loaded")
	}
	return _cGDisplayModeGetPixelWidth(mode)
}

var _cGDisplayModeGetRefreshRate func(mode CGDisplayModeRef) float64

// CGDisplayModeGetRefreshRate returns the refresh rate of the specified display mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/refreshRate
func CGDisplayModeGetRefreshRate(mode CGDisplayModeRef) float64 {
	if _cGDisplayModeGetRefreshRate == nil {
		panic("CoreGraphics: symbol CGDisplayModeGetRefreshRate not loaded")
	}
	return _cGDisplayModeGetRefreshRate(mode)
}

var _cGDisplayModeGetTypeID func() uint

// CGDisplayModeGetTypeID returns the type identifier of Quartz display modes.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/typeID
func CGDisplayModeGetTypeID() uint {
	if _cGDisplayModeGetTypeID == nil {
		panic("CoreGraphics: symbol CGDisplayModeGetTypeID not loaded")
	}
	return _cGDisplayModeGetTypeID()
}

var _cGDisplayModeGetWidth func(mode CGDisplayModeRef) uintptr

// CGDisplayModeGetWidth returns the width of the specified display mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/width
func CGDisplayModeGetWidth(mode CGDisplayModeRef) uintptr {
	if _cGDisplayModeGetWidth == nil {
		panic("CoreGraphics: symbol CGDisplayModeGetWidth not loaded")
	}
	return _cGDisplayModeGetWidth(mode)
}

var _cGDisplayModeIsUsableForDesktopGUI func(mode CGDisplayModeRef) bool

// CGDisplayModeIsUsableForDesktopGUI returns a Boolean value indicating whether the specified display mode is usable for a desktop graphical user interface.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/isUsableForDesktopGUI()
func CGDisplayModeIsUsableForDesktopGUI(mode CGDisplayModeRef) bool {
	if _cGDisplayModeIsUsableForDesktopGUI == nil {
		panic("CoreGraphics: symbol CGDisplayModeIsUsableForDesktopGUI not loaded")
	}
	return _cGDisplayModeIsUsableForDesktopGUI(mode)
}

var _cGDisplayModeRelease func(mode CGDisplayModeRef)

// CGDisplayModeRelease releases a Core Graphics display mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayModeRelease
func CGDisplayModeRelease(mode CGDisplayModeRef) {
	if _cGDisplayModeRelease == nil {
		panic("CoreGraphics: symbol CGDisplayModeRelease not loaded")
	}
	_cGDisplayModeRelease(mode)
}

var _cGDisplayModeRetain func(mode CGDisplayModeRef) CGDisplayModeRef

// CGDisplayModeRetain retains a Core Graphics display mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayModeRetain
func CGDisplayModeRetain(mode CGDisplayModeRef) CGDisplayModeRef {
	if _cGDisplayModeRetain == nil {
		panic("CoreGraphics: symbol CGDisplayModeRetain not loaded")
	}
	return _cGDisplayModeRetain(mode)
}

var _cGDisplayModelNumber func(display uint32) uint32

// CGDisplayModelNumber returns the model number of a display monitor.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayModelNumber(_:)
func CGDisplayModelNumber(display uint32) uint32 {
	if _cGDisplayModelNumber == nil {
		panic("CoreGraphics: symbol CGDisplayModelNumber not loaded")
	}
	return _cGDisplayModelNumber(display)
}

var _cGDisplayMoveCursorToPoint func(display uint32, point corefoundation.CGPoint) CGError

// CGDisplayMoveCursorToPoint moves the mouse cursor to a specified point relative to the upper-left corner of the display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMoveCursorToPoint(_:_:)
func CGDisplayMoveCursorToPoint(display uint32, point corefoundation.CGPoint) CGError {
	if _cGDisplayMoveCursorToPoint == nil {
		panic("CoreGraphics: symbol CGDisplayMoveCursorToPoint not loaded")
	}
	return _cGDisplayMoveCursorToPoint(display, point)
}

var _cGDisplayPixelsHigh func(display uint32) uintptr

// CGDisplayPixelsHigh returns the display height in pixel units.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayPixelsHigh(_:)
func CGDisplayPixelsHigh(display uint32) uintptr {
	if _cGDisplayPixelsHigh == nil {
		panic("CoreGraphics: symbol CGDisplayPixelsHigh not loaded")
	}
	return _cGDisplayPixelsHigh(display)
}

var _cGDisplayPixelsWide func(display uint32) uintptr

// CGDisplayPixelsWide returns the display width in pixel units.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayPixelsWide(_:)
func CGDisplayPixelsWide(display uint32) uintptr {
	if _cGDisplayPixelsWide == nil {
		panic("CoreGraphics: symbol CGDisplayPixelsWide not loaded")
	}
	return _cGDisplayPixelsWide(display)
}

var _cGDisplayPrimaryDisplay func(display uint32) uint32

// CGDisplayPrimaryDisplay returns the primary display in a hardware mirroring set.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayPrimaryDisplay(_:)
func CGDisplayPrimaryDisplay(display uint32) uint32 {
	if _cGDisplayPrimaryDisplay == nil {
		panic("CoreGraphics: symbol CGDisplayPrimaryDisplay not loaded")
	}
	return _cGDisplayPrimaryDisplay(display)
}

var _cGDisplayRegisterReconfigurationCallback func(callback CGDisplayReconfigurationCallBack, userInfo unsafe.Pointer) CGError

// CGDisplayRegisterReconfigurationCallback registers a callback function to be invoked whenever a local display is reconfigured.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayRegisterReconfigurationCallback(_:_:)
func CGDisplayRegisterReconfigurationCallback(callback CGDisplayReconfigurationCallBack, userInfo unsafe.Pointer) CGError {
	if _cGDisplayRegisterReconfigurationCallback == nil {
		panic("CoreGraphics: symbol CGDisplayRegisterReconfigurationCallback not loaded")
	}
	return _cGDisplayRegisterReconfigurationCallback(callback, userInfo)
}

var _cGDisplayRelease func(display uint32) CGError

// CGDisplayRelease releases a captured display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayRelease(_:)
func CGDisplayRelease(display uint32) CGError {
	if _cGDisplayRelease == nil {
		panic("CoreGraphics: symbol CGDisplayRelease not loaded")
	}
	return _cGDisplayRelease(display)
}

var _cGDisplayRemoveReconfigurationCallback func(callback CGDisplayReconfigurationCallBack, userInfo unsafe.Pointer) CGError

// CGDisplayRemoveReconfigurationCallback removes the registration of a callback function that’s invoked whenever a local display is reconfigured.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayRemoveReconfigurationCallback(_:_:)
func CGDisplayRemoveReconfigurationCallback(callback CGDisplayReconfigurationCallBack, userInfo unsafe.Pointer) CGError {
	if _cGDisplayRemoveReconfigurationCallback == nil {
		panic("CoreGraphics: symbol CGDisplayRemoveReconfigurationCallback not loaded")
	}
	return _cGDisplayRemoveReconfigurationCallback(callback, userInfo)
}

var _cGDisplayRestoreColorSyncSettings func()

// CGDisplayRestoreColorSyncSettings restores the gamma tables to the values in the user’s ColorSync display profile.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayRestoreColorSyncSettings()
func CGDisplayRestoreColorSyncSettings() {
	if _cGDisplayRestoreColorSyncSettings == nil {
		panic("CoreGraphics: symbol CGDisplayRestoreColorSyncSettings not loaded")
	}
	_cGDisplayRestoreColorSyncSettings()
}

var _cGDisplayRotation func(display uint32) float64

// CGDisplayRotation returns the rotation angle of a display in degrees.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayRotation(_:)
func CGDisplayRotation(display uint32) float64 {
	if _cGDisplayRotation == nil {
		panic("CoreGraphics: symbol CGDisplayRotation not loaded")
	}
	return _cGDisplayRotation(display)
}

var _cGDisplayScreenSize func(display uint32) corefoundation.CGSize

// CGDisplayScreenSize returns the width and height of a display in millimeters.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayScreenSize(_:)
func CGDisplayScreenSize(display uint32) corefoundation.CGSize {
	if _cGDisplayScreenSize == nil {
		panic("CoreGraphics: symbol CGDisplayScreenSize not loaded")
	}
	return _cGDisplayScreenSize(display)
}

var _cGDisplaySerialNumber func(display uint32) uint32

// CGDisplaySerialNumber returns the serial number of a display monitor.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplaySerialNumber(_:)
func CGDisplaySerialNumber(display uint32) uint32 {
	if _cGDisplaySerialNumber == nil {
		panic("CoreGraphics: symbol CGDisplaySerialNumber not loaded")
	}
	return _cGDisplaySerialNumber(display)
}

var _cGDisplaySetDisplayMode func(display uint32, mode CGDisplayModeRef, options corefoundation.CFDictionaryRef) CGError

// CGDisplaySetDisplayMode switches a display to a different mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplaySetDisplayMode(_:_:_:)
func CGDisplaySetDisplayMode(display uint32, mode CGDisplayModeRef, options corefoundation.CFDictionaryRef) CGError {
	if _cGDisplaySetDisplayMode == nil {
		panic("CoreGraphics: symbol CGDisplaySetDisplayMode not loaded")
	}
	return _cGDisplaySetDisplayMode(display, mode, options)
}

var _cGDisplaySetStereoOperation func(display uint32, stereo uintptr, forceBlueLine uintptr, option CGConfigureOption) CGError

// CGDisplaySetStereoOperation immediately enables or disables stereo operation for a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplaySetStereoOperation(_:_:_:_:)
func CGDisplaySetStereoOperation(display uint32, stereo uintptr, forceBlueLine uintptr, option CGConfigureOption) CGError {
	if _cGDisplaySetStereoOperation == nil {
		panic("CoreGraphics: symbol CGDisplaySetStereoOperation not loaded")
	}
	return _cGDisplaySetStereoOperation(display, stereo, forceBlueLine, option)
}

var _cGDisplayShowCursor func(display uint32) CGError

// CGDisplayShowCursor decrements the hide cursor count, and shows the mouse cursor if the count is `0`.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayShowCursor(_:)
func CGDisplayShowCursor(display uint32) CGError {
	if _cGDisplayShowCursor == nil {
		panic("CoreGraphics: symbol CGDisplayShowCursor not loaded")
	}
	return _cGDisplayShowCursor(display)
}

var _cGDisplayUnitNumber func(display uint32) uint32

// CGDisplayUnitNumber returns the logical unit number of a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayUnitNumber(_:)
func CGDisplayUnitNumber(display uint32) uint32 {
	if _cGDisplayUnitNumber == nil {
		panic("CoreGraphics: symbol CGDisplayUnitNumber not loaded")
	}
	return _cGDisplayUnitNumber(display)
}

var _cGDisplayUsesOpenGLAcceleration func(display uint32) objectivec.IObject

// CGDisplayUsesOpenGLAcceleration returns a Boolean value indicating whether Quartz is using OpenGL-based window acceleration (Quartz Extreme) to render in a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayUsesOpenGLAcceleration(_:)
func CGDisplayUsesOpenGLAcceleration(display uint32) objectivec.IObject {
	if _cGDisplayUsesOpenGLAcceleration == nil {
		panic("CoreGraphics: symbol CGDisplayUsesOpenGLAcceleration not loaded")
	}
	return _cGDisplayUsesOpenGLAcceleration(display)
}

var _cGDisplayVendorNumber func(display uint32) uint32

// CGDisplayVendorNumber returns the vendor number of the specified display’s monitor.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayVendorNumber(_:)
func CGDisplayVendorNumber(display uint32) uint32 {
	if _cGDisplayVendorNumber == nil {
		panic("CoreGraphics: symbol CGDisplayVendorNumber not loaded")
	}
	return _cGDisplayVendorNumber(display)
}

var _cGEXRToneMappingGammaGetDefaultOptions func() corefoundation.CFDictionaryRef

// CGEXRToneMappingGammaGetDefaultOptions.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEXRToneMappingGammaGetDefaultOptions
func CGEXRToneMappingGammaGetDefaultOptions() corefoundation.CFDictionaryRef {
	if _cGEXRToneMappingGammaGetDefaultOptions == nil {
		panic("CoreGraphics: symbol CGEXRToneMappingGammaGetDefaultOptions not loaded")
	}
	return _cGEXRToneMappingGammaGetDefaultOptions()
}

var _cGErrorSetCallback func(callback CGErrorCallback)

// CGErrorSetCallback.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGErrorSetCallback(_:)
func CGErrorSetCallback(callback CGErrorCallback) {
	if _cGErrorSetCallback == nil {
		panic("CoreGraphics: symbol CGErrorSetCallback not loaded")
	}
	_cGErrorSetCallback(callback)
}

var _cGEventCreate func(source CGEventSourceRef) CGEventRef

// CGEventCreate returns a new Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/init(source:)
func CGEventCreate(source CGEventSourceRef) CGEventRef {
	if _cGEventCreate == nil {
		panic("CoreGraphics: symbol CGEventCreate not loaded")
	}
	return _cGEventCreate(source)
}

var _cGEventCreateCopy func(event CGEventRef) CGEventRef

// CGEventCreateCopy returns a copy of an existing Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/copy()
func CGEventCreateCopy(event CGEventRef) CGEventRef {
	if _cGEventCreateCopy == nil {
		panic("CoreGraphics: symbol CGEventCreateCopy not loaded")
	}
	return _cGEventCreateCopy(event)
}

var _cGEventCreateData func(allocator corefoundation.CFAllocatorRef, event CGEventRef) corefoundation.CFDataRef

// CGEventCreateData returns a flattened data representation of a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventCreateData
func CGEventCreateData(allocator corefoundation.CFAllocatorRef, event CGEventRef) corefoundation.CFDataRef {
	if _cGEventCreateData == nil {
		panic("CoreGraphics: symbol CGEventCreateData not loaded")
	}
	return _cGEventCreateData(allocator, event)
}

var _cGEventCreateFromData func(allocator corefoundation.CFAllocatorRef, data corefoundation.CFDataRef) CGEventRef

// CGEventCreateFromData returns a Quartz event created from a flattened data representation of the event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/init(withDataAllocator:data:)
func CGEventCreateFromData(allocator corefoundation.CFAllocatorRef, data corefoundation.CFDataRef) CGEventRef {
	if _cGEventCreateFromData == nil {
		panic("CoreGraphics: symbol CGEventCreateFromData not loaded")
	}
	return _cGEventCreateFromData(allocator, data)
}

var _cGEventCreateKeyboardEvent func(source CGEventSourceRef, virtualKey uint16, keyDown bool) CGEventRef

// CGEventCreateKeyboardEvent returns a new Quartz keyboard event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/init(keyboardEventSource:virtualKey:keyDown:)
func CGEventCreateKeyboardEvent(source CGEventSourceRef, virtualKey uint16, keyDown bool) CGEventRef {
	if _cGEventCreateKeyboardEvent == nil {
		panic("CoreGraphics: symbol CGEventCreateKeyboardEvent not loaded")
	}
	return _cGEventCreateKeyboardEvent(source, virtualKey, keyDown)
}

var _cGEventCreateMouseEvent func(source CGEventSourceRef, mouseType CGEventType, mouseCursorPosition corefoundation.CGPoint, mouseButton CGMouseButton) CGEventRef

// CGEventCreateMouseEvent returns a new Quartz mouse event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/init(mouseEventSource:mouseType:mouseCursorPosition:mouseButton:)
func CGEventCreateMouseEvent(source CGEventSourceRef, mouseType CGEventType, mouseCursorPosition corefoundation.CGPoint, mouseButton CGMouseButton) CGEventRef {
	if _cGEventCreateMouseEvent == nil {
		panic("CoreGraphics: symbol CGEventCreateMouseEvent not loaded")
	}
	return _cGEventCreateMouseEvent(source, mouseType, mouseCursorPosition, mouseButton)
}

var _cGEventCreateScrollWheelEvent func(source CGEventSourceRef, units CGScrollEventUnit, wheelCount uint32, wheel1 int32) CGEventRef

// CGEventCreateScrollWheelEvent returns a new Quartz scrolling event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventCreateScrollWheelEvent
func CGEventCreateScrollWheelEvent(source CGEventSourceRef, units CGScrollEventUnit, wheelCount uint32, wheel1 int32) CGEventRef {
	if _cGEventCreateScrollWheelEvent == nil {
		panic("CoreGraphics: symbol CGEventCreateScrollWheelEvent not loaded")
	}
	return _cGEventCreateScrollWheelEvent(source, units, wheelCount, wheel1)
}

var _cGEventCreateScrollWheelEvent2 func(source CGEventSourceRef, units CGScrollEventUnit, wheelCount uint32, wheel1 int32, wheel2 int32, wheel3 int32) CGEventRef

// CGEventCreateScrollWheelEvent2.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/init(scrollWheelEvent2Source:units:wheelCount:wheel1:wheel2:wheel3:)
func CGEventCreateScrollWheelEvent2(source CGEventSourceRef, units CGScrollEventUnit, wheelCount uint32, wheel1 int32, wheel2 int32, wheel3 int32) CGEventRef {
	if _cGEventCreateScrollWheelEvent2 == nil {
		panic("CoreGraphics: symbol CGEventCreateScrollWheelEvent2 not loaded")
	}
	return _cGEventCreateScrollWheelEvent2(source, units, wheelCount, wheel1, wheel2, wheel3)
}

var _cGEventCreateSourceFromEvent func(event CGEventRef) CGEventSourceRef

// CGEventCreateSourceFromEvent returns a Quartz event source created from an existing Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/init(event:)
func CGEventCreateSourceFromEvent(event CGEventRef) CGEventSourceRef {
	if _cGEventCreateSourceFromEvent == nil {
		panic("CoreGraphics: symbol CGEventCreateSourceFromEvent not loaded")
	}
	return _cGEventCreateSourceFromEvent(event)
}

var _cGEventGetDoubleValueField func(event CGEventRef, field CGEventField) float64

// CGEventGetDoubleValueField returns the floating-point value of a field in a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/getDoubleValueField(_:)
func CGEventGetDoubleValueField(event CGEventRef, field CGEventField) float64 {
	if _cGEventGetDoubleValueField == nil {
		panic("CoreGraphics: symbol CGEventGetDoubleValueField not loaded")
	}
	return _cGEventGetDoubleValueField(event, field)
}

var _cGEventGetFlags func(event CGEventRef) CGEventFlags

// CGEventGetFlags returns the event flags of a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/flags
func CGEventGetFlags(event CGEventRef) CGEventFlags {
	if _cGEventGetFlags == nil {
		panic("CoreGraphics: symbol CGEventGetFlags not loaded")
	}
	return _cGEventGetFlags(event)
}

var _cGEventGetIntegerValueField func(event CGEventRef, field CGEventField) int64

// CGEventGetIntegerValueField returns the integer value of a field in a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/getIntegerValueField(_:)
func CGEventGetIntegerValueField(event CGEventRef, field CGEventField) int64 {
	if _cGEventGetIntegerValueField == nil {
		panic("CoreGraphics: symbol CGEventGetIntegerValueField not loaded")
	}
	return _cGEventGetIntegerValueField(event, field)
}

var _cGEventGetLocation func(event CGEventRef) corefoundation.CGPoint

// CGEventGetLocation returns the location of a Quartz mouse event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/location
func CGEventGetLocation(event CGEventRef) corefoundation.CGPoint {
	if _cGEventGetLocation == nil {
		panic("CoreGraphics: symbol CGEventGetLocation not loaded")
	}
	return _cGEventGetLocation(event)
}

var _cGEventGetTimestamp func(event CGEventRef) CGEventTimestamp

// CGEventGetTimestamp returns the timestamp of a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/timestamp
func CGEventGetTimestamp(event CGEventRef) CGEventTimestamp {
	if _cGEventGetTimestamp == nil {
		panic("CoreGraphics: symbol CGEventGetTimestamp not loaded")
	}
	return _cGEventGetTimestamp(event)
}

var _cGEventGetType func(event CGEventRef) CGEventType

// CGEventGetType returns the event type of a Quartz event (left mouse down, for example).
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/type
func CGEventGetType(event CGEventRef) CGEventType {
	if _cGEventGetType == nil {
		panic("CoreGraphics: symbol CGEventGetType not loaded")
	}
	return _cGEventGetType(event)
}

var _cGEventGetTypeID func() uint

// CGEventGetTypeID returns the type identifier for the opaque type [CGEventRef].
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/typeID
func CGEventGetTypeID() uint {
	if _cGEventGetTypeID == nil {
		panic("CoreGraphics: symbol CGEventGetTypeID not loaded")
	}
	return _cGEventGetTypeID()
}

var _cGEventGetUnflippedLocation func(event CGEventRef) corefoundation.CGPoint

// CGEventGetUnflippedLocation returns the location of a Quartz mouse event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/unflippedLocation
func CGEventGetUnflippedLocation(event CGEventRef) corefoundation.CGPoint {
	if _cGEventGetUnflippedLocation == nil {
		panic("CoreGraphics: symbol CGEventGetUnflippedLocation not loaded")
	}
	return _cGEventGetUnflippedLocation(event)
}

var _cGEventKeyboardGetUnicodeString func(event CGEventRef, maxStringLength uint, actualStringLength *uint, unicodeString *uint16)

// CGEventKeyboardGetUnicodeString returns the Unicode string associated with a Quartz keyboard event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/keyboardGetUnicodeString(maxStringLength:actualStringLength:unicodeString:)
func CGEventKeyboardGetUnicodeString(event CGEventRef, maxStringLength uint, actualStringLength *uint, unicodeString *uint16) {
	if _cGEventKeyboardGetUnicodeString == nil {
		panic("CoreGraphics: symbol CGEventKeyboardGetUnicodeString not loaded")
	}
	_cGEventKeyboardGetUnicodeString(event, maxStringLength, actualStringLength, unicodeString)
}

var _cGEventKeyboardSetUnicodeString func(event CGEventRef, stringLength uint, unicodeString *uint16)

// CGEventKeyboardSetUnicodeString sets the Unicode string associated with a Quartz keyboard event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/keyboardSetUnicodeString(stringLength:unicodeString:)
func CGEventKeyboardSetUnicodeString(event CGEventRef, stringLength uint, unicodeString *uint16) {
	if _cGEventKeyboardSetUnicodeString == nil {
		panic("CoreGraphics: symbol CGEventKeyboardSetUnicodeString not loaded")
	}
	_cGEventKeyboardSetUnicodeString(event, stringLength, unicodeString)
}

var _cGEventPost func(tap CGEventTapLocation, event CGEventRef)

// CGEventPost posts a Quartz event into the event stream at a specified location.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/post(tap:)
func CGEventPost(tap CGEventTapLocation, event CGEventRef) {
	if _cGEventPost == nil {
		panic("CoreGraphics: symbol CGEventPost not loaded")
	}
	_cGEventPost(tap, event)
}

var _cGEventPostToPSN func(processSerialNumber unsafe.Pointer, event CGEventRef)

// CGEventPostToPSN posts a Quartz event into the event stream for a specific application.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/postToPSN(processSerialNumber:)
func CGEventPostToPSN(processSerialNumber unsafe.Pointer, event CGEventRef) {
	if _cGEventPostToPSN == nil {
		panic("CoreGraphics: symbol CGEventPostToPSN not loaded")
	}
	_cGEventPostToPSN(processSerialNumber, event)
}

var _cGEventPostToPid func(pid int32, event CGEventRef)

// CGEventPostToPid.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/postToPid(_:)
func CGEventPostToPid(pid int32, event CGEventRef) {
	if _cGEventPostToPid == nil {
		panic("CoreGraphics: symbol CGEventPostToPid not loaded")
	}
	_cGEventPostToPid(pid, event)
}

var _cGEventSetDoubleValueField func(event CGEventRef, field CGEventField, value float64)

// CGEventSetDoubleValueField sets the floating-point value of a field in a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/setDoubleValueField(_:value:)
func CGEventSetDoubleValueField(event CGEventRef, field CGEventField, value float64) {
	if _cGEventSetDoubleValueField == nil {
		panic("CoreGraphics: symbol CGEventSetDoubleValueField not loaded")
	}
	_cGEventSetDoubleValueField(event, field, value)
}

var _cGEventSetFlags func(event CGEventRef, flags CGEventFlags)

// CGEventSetFlags sets the event flags of a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSetFlags
func CGEventSetFlags(event CGEventRef, flags CGEventFlags) {
	if _cGEventSetFlags == nil {
		panic("CoreGraphics: symbol CGEventSetFlags not loaded")
	}
	_cGEventSetFlags(event, flags)
}

var _cGEventSetIntegerValueField func(event CGEventRef, field CGEventField, value int64)

// CGEventSetIntegerValueField sets the integer value of a field in a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/setIntegerValueField(_:value:)
func CGEventSetIntegerValueField(event CGEventRef, field CGEventField, value int64) {
	if _cGEventSetIntegerValueField == nil {
		panic("CoreGraphics: symbol CGEventSetIntegerValueField not loaded")
	}
	_cGEventSetIntegerValueField(event, field, value)
}

var _cGEventSetLocation func(event CGEventRef, location corefoundation.CGPoint)

// CGEventSetLocation sets the location of a Quartz mouse event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSetLocation
func CGEventSetLocation(event CGEventRef, location corefoundation.CGPoint) {
	if _cGEventSetLocation == nil {
		panic("CoreGraphics: symbol CGEventSetLocation not loaded")
	}
	_cGEventSetLocation(event, location)
}

var _cGEventSetSource func(event CGEventRef, source CGEventSourceRef)

// CGEventSetSource sets the event source of a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/setSource(_:)
func CGEventSetSource(event CGEventRef, source CGEventSourceRef) {
	if _cGEventSetSource == nil {
		panic("CoreGraphics: symbol CGEventSetSource not loaded")
	}
	_cGEventSetSource(event, source)
}

var _cGEventSetTimestamp func(event CGEventRef, timestamp CGEventTimestamp)

// CGEventSetTimestamp sets the timestamp of a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSetTimestamp
func CGEventSetTimestamp(event CGEventRef, timestamp CGEventTimestamp) {
	if _cGEventSetTimestamp == nil {
		panic("CoreGraphics: symbol CGEventSetTimestamp not loaded")
	}
	_cGEventSetTimestamp(event, timestamp)
}

var _cGEventSetType func(event CGEventRef, type_ CGEventType)

// CGEventSetType sets the event type of a Quartz event (left mouse down, for example).
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSetType
func CGEventSetType(event CGEventRef, type_ CGEventType) {
	if _cGEventSetType == nil {
		panic("CoreGraphics: symbol CGEventSetType not loaded")
	}
	_cGEventSetType(event, type_)
}

var _cGEventSourceButtonState func(stateID CGEventSourceStateID, button CGMouseButton) bool

// CGEventSourceButtonState returns a Boolean value indicating the current button state of a Quartz event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/buttonState(_:button:)
func CGEventSourceButtonState(stateID CGEventSourceStateID, button CGMouseButton) bool {
	if _cGEventSourceButtonState == nil {
		panic("CoreGraphics: symbol CGEventSourceButtonState not loaded")
	}
	return _cGEventSourceButtonState(stateID, button)
}

var _cGEventSourceCounterForEventType func(stateID CGEventSourceStateID, eventType CGEventType) uint32

// CGEventSourceCounterForEventType returns a count of events of a given type seen since the window server started.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/counterForEventType(_:eventType:)
func CGEventSourceCounterForEventType(stateID CGEventSourceStateID, eventType CGEventType) uint32 {
	if _cGEventSourceCounterForEventType == nil {
		panic("CoreGraphics: symbol CGEventSourceCounterForEventType not loaded")
	}
	return _cGEventSourceCounterForEventType(stateID, eventType)
}

var _cGEventSourceCreate func(stateID CGEventSourceStateID) CGEventSourceRef

// CGEventSourceCreate returns a Quartz event source created with a specified source state.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/init(stateID:)
func CGEventSourceCreate(stateID CGEventSourceStateID) CGEventSourceRef {
	if _cGEventSourceCreate == nil {
		panic("CoreGraphics: symbol CGEventSourceCreate not loaded")
	}
	return _cGEventSourceCreate(stateID)
}

var _cGEventSourceFlagsState func(stateID CGEventSourceStateID) CGEventFlags

// CGEventSourceFlagsState returns the current flags of a Quartz event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/flagsState(_:)
func CGEventSourceFlagsState(stateID CGEventSourceStateID) CGEventFlags {
	if _cGEventSourceFlagsState == nil {
		panic("CoreGraphics: symbol CGEventSourceFlagsState not loaded")
	}
	return _cGEventSourceFlagsState(stateID)
}

var _cGEventSourceGetKeyboardType func(source CGEventSourceRef) CGEventSourceKeyboardType

// CGEventSourceGetKeyboardType returns the keyboard type to be used with a Quartz event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/keyboardType
func CGEventSourceGetKeyboardType(source CGEventSourceRef) CGEventSourceKeyboardType {
	if _cGEventSourceGetKeyboardType == nil {
		panic("CoreGraphics: symbol CGEventSourceGetKeyboardType not loaded")
	}
	return _cGEventSourceGetKeyboardType(source)
}

var _cGEventSourceGetLocalEventsFilterDuringSuppressionState func(source CGEventSourceRef, state CGEventSuppressionState) CGEventFilterMask

// CGEventSourceGetLocalEventsFilterDuringSuppressionState returns the mask that indicates which classes of local hardware events are enabled during event suppression.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/getLocalEventsFilterDuringSuppressionState(_:)
func CGEventSourceGetLocalEventsFilterDuringSuppressionState(source CGEventSourceRef, state CGEventSuppressionState) CGEventFilterMask {
	if _cGEventSourceGetLocalEventsFilterDuringSuppressionState == nil {
		panic("CoreGraphics: symbol CGEventSourceGetLocalEventsFilterDuringSuppressionState not loaded")
	}
	return _cGEventSourceGetLocalEventsFilterDuringSuppressionState(source, state)
}

var _cGEventSourceGetLocalEventsSuppressionInterval func(source CGEventSourceRef) float64

// CGEventSourceGetLocalEventsSuppressionInterval returns the interval that local hardware events may be suppressed following the posting of a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/localEventsSuppressionInterval
func CGEventSourceGetLocalEventsSuppressionInterval(source CGEventSourceRef) float64 {
	if _cGEventSourceGetLocalEventsSuppressionInterval == nil {
		panic("CoreGraphics: symbol CGEventSourceGetLocalEventsSuppressionInterval not loaded")
	}
	return _cGEventSourceGetLocalEventsSuppressionInterval(source)
}

var _cGEventSourceGetPixelsPerLine func(source CGEventSourceRef) float64

// CGEventSourceGetPixelsPerLine gets the scale of pixels per line in a scrolling event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/pixelsPerLine
func CGEventSourceGetPixelsPerLine(source CGEventSourceRef) float64 {
	if _cGEventSourceGetPixelsPerLine == nil {
		panic("CoreGraphics: symbol CGEventSourceGetPixelsPerLine not loaded")
	}
	return _cGEventSourceGetPixelsPerLine(source)
}

var _cGEventSourceGetSourceStateID func(source CGEventSourceRef) CGEventSourceStateID

// CGEventSourceGetSourceStateID returns the source state associated with a Quartz event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/sourceStateID
func CGEventSourceGetSourceStateID(source CGEventSourceRef) CGEventSourceStateID {
	if _cGEventSourceGetSourceStateID == nil {
		panic("CoreGraphics: symbol CGEventSourceGetSourceStateID not loaded")
	}
	return _cGEventSourceGetSourceStateID(source)
}

var _cGEventSourceGetTypeID func() uint

// CGEventSourceGetTypeID returns the type identifier for the opaque type [CGEventSourceRef].
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/typeID
func CGEventSourceGetTypeID() uint {
	if _cGEventSourceGetTypeID == nil {
		panic("CoreGraphics: symbol CGEventSourceGetTypeID not loaded")
	}
	return _cGEventSourceGetTypeID()
}

var _cGEventSourceGetUserData func(source CGEventSourceRef) int64

// CGEventSourceGetUserData returns the 64-bit user-specified data for a Quartz event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/userData
func CGEventSourceGetUserData(source CGEventSourceRef) int64 {
	if _cGEventSourceGetUserData == nil {
		panic("CoreGraphics: symbol CGEventSourceGetUserData not loaded")
	}
	return _cGEventSourceGetUserData(source)
}

var _cGEventSourceKeyState func(stateID CGEventSourceStateID, key uint16) bool

// CGEventSourceKeyState returns a Boolean value indicating the current keyboard state of a Quartz event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/keyState(_:key:)
func CGEventSourceKeyState(stateID CGEventSourceStateID, key uint16) bool {
	if _cGEventSourceKeyState == nil {
		panic("CoreGraphics: symbol CGEventSourceKeyState not loaded")
	}
	return _cGEventSourceKeyState(stateID, key)
}

var _cGEventSourceSecondsSinceLastEventType func(stateID CGEventSourceStateID, eventType CGEventType) float64

// CGEventSourceSecondsSinceLastEventType returns the elapsed time since the last event for a Quartz event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/secondsSinceLastEventType(_:eventType:)
func CGEventSourceSecondsSinceLastEventType(stateID CGEventSourceStateID, eventType CGEventType) float64 {
	if _cGEventSourceSecondsSinceLastEventType == nil {
		panic("CoreGraphics: symbol CGEventSourceSecondsSinceLastEventType not loaded")
	}
	return _cGEventSourceSecondsSinceLastEventType(stateID, eventType)
}

var _cGEventSourceSetKeyboardType func(source CGEventSourceRef, keyboardType CGEventSourceKeyboardType)

// CGEventSourceSetKeyboardType sets the keyboard type to be used with a Quartz event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceSetKeyboardType
func CGEventSourceSetKeyboardType(source CGEventSourceRef, keyboardType CGEventSourceKeyboardType) {
	if _cGEventSourceSetKeyboardType == nil {
		panic("CoreGraphics: symbol CGEventSourceSetKeyboardType not loaded")
	}
	_cGEventSourceSetKeyboardType(source, keyboardType)
}

var _cGEventSourceSetLocalEventsFilterDuringSuppressionState func(source CGEventSourceRef, filter CGEventFilterMask, state CGEventSuppressionState)

// CGEventSourceSetLocalEventsFilterDuringSuppressionState sets the mask that indicates which classes of local hardware events are enabled during event suppression.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/setLocalEventsFilterDuringSuppressionState(_:state:)
func CGEventSourceSetLocalEventsFilterDuringSuppressionState(source CGEventSourceRef, filter CGEventFilterMask, state CGEventSuppressionState) {
	if _cGEventSourceSetLocalEventsFilterDuringSuppressionState == nil {
		panic("CoreGraphics: symbol CGEventSourceSetLocalEventsFilterDuringSuppressionState not loaded")
	}
	_cGEventSourceSetLocalEventsFilterDuringSuppressionState(source, filter, state)
}

var _cGEventSourceSetLocalEventsSuppressionInterval func(source CGEventSourceRef, seconds float64)

// CGEventSourceSetLocalEventsSuppressionInterval sets the interval that local hardware events may be suppressed following the posting of a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceSetLocalEventsSuppressionInterval
func CGEventSourceSetLocalEventsSuppressionInterval(source CGEventSourceRef, seconds float64) {
	if _cGEventSourceSetLocalEventsSuppressionInterval == nil {
		panic("CoreGraphics: symbol CGEventSourceSetLocalEventsSuppressionInterval not loaded")
	}
	_cGEventSourceSetLocalEventsSuppressionInterval(source, seconds)
}

var _cGEventSourceSetPixelsPerLine func(source CGEventSourceRef, pixelsPerLine float64)

// CGEventSourceSetPixelsPerLine sets the scale of pixels per line in a scrolling event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceSetPixelsPerLine
func CGEventSourceSetPixelsPerLine(source CGEventSourceRef, pixelsPerLine float64) {
	if _cGEventSourceSetPixelsPerLine == nil {
		panic("CoreGraphics: symbol CGEventSourceSetPixelsPerLine not loaded")
	}
	_cGEventSourceSetPixelsPerLine(source, pixelsPerLine)
}

var _cGEventSourceSetUserData func(source CGEventSourceRef, userData int64)

// CGEventSourceSetUserData sets the 64-bit user-specified data for a Quartz event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceSetUserData
func CGEventSourceSetUserData(source CGEventSourceRef, userData int64) {
	if _cGEventSourceSetUserData == nil {
		panic("CoreGraphics: symbol CGEventSourceSetUserData not loaded")
	}
	_cGEventSourceSetUserData(source, userData)
}

var _cGEventTapCreate func(tap CGEventTapLocation, place CGEventTapPlacement, options CGEventTapOptions, eventsOfInterest CGEventMask, callback CGEventTapCallBack, userInfo unsafe.Pointer) corefoundation.CFMachPort

// CGEventTapCreate creates an event tap.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/tapCreate(tap:place:options:eventsOfInterest:callback:userInfo:)
func CGEventTapCreate(tap CGEventTapLocation, place CGEventTapPlacement, options CGEventTapOptions, eventsOfInterest CGEventMask, callback CGEventTapCallBack, userInfo unsafe.Pointer) corefoundation.CFMachPort {
	if _cGEventTapCreate == nil {
		panic("CoreGraphics: symbol CGEventTapCreate not loaded")
	}
	return _cGEventTapCreate(tap, place, options, eventsOfInterest, callback, userInfo)
}

var _cGEventTapCreateForPSN func(processSerialNumber unsafe.Pointer, place CGEventTapPlacement, options CGEventTapOptions, eventsOfInterest CGEventMask, callback CGEventTapCallBack, userInfo unsafe.Pointer) corefoundation.CFMachPort

// CGEventTapCreateForPSN creates an event tap for a specified process.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/tapCreateForPSN(processSerialNumber:place:options:eventsOfInterest:callback:userInfo:)
func CGEventTapCreateForPSN(processSerialNumber unsafe.Pointer, place CGEventTapPlacement, options CGEventTapOptions, eventsOfInterest CGEventMask, callback CGEventTapCallBack, userInfo unsafe.Pointer) corefoundation.CFMachPort {
	if _cGEventTapCreateForPSN == nil {
		panic("CoreGraphics: symbol CGEventTapCreateForPSN not loaded")
	}
	return _cGEventTapCreateForPSN(processSerialNumber, place, options, eventsOfInterest, callback, userInfo)
}

var _cGEventTapCreateForPid func(pid int32, place CGEventTapPlacement, options CGEventTapOptions, eventsOfInterest CGEventMask, callback CGEventTapCallBack, userInfo unsafe.Pointer) corefoundation.CFMachPort

// CGEventTapCreateForPid.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/tapCreateForPid(pid:place:options:eventsOfInterest:callback:userInfo:)
func CGEventTapCreateForPid(pid int32, place CGEventTapPlacement, options CGEventTapOptions, eventsOfInterest CGEventMask, callback CGEventTapCallBack, userInfo unsafe.Pointer) corefoundation.CFMachPort {
	if _cGEventTapCreateForPid == nil {
		panic("CoreGraphics: symbol CGEventTapCreateForPid not loaded")
	}
	return _cGEventTapCreateForPid(pid, place, options, eventsOfInterest, callback, userInfo)
}

var _cGEventTapEnable func(tap corefoundation.CFMachPort, enable bool)

// CGEventTapEnable enables or disables an event tap.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/tapEnable(tap:enable:)
func CGEventTapEnable(tap corefoundation.CFMachPort, enable bool) {
	if _cGEventTapEnable == nil {
		panic("CoreGraphics: symbol CGEventTapEnable not loaded")
	}
	_cGEventTapEnable(tap, enable)
}

var _cGEventTapIsEnabled func(tap corefoundation.CFMachPort) bool

// CGEventTapIsEnabled returns a Boolean value indicating whether an event tap is enabled.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/tapIsEnabled(tap:)
func CGEventTapIsEnabled(tap corefoundation.CFMachPort) bool {
	if _cGEventTapIsEnabled == nil {
		panic("CoreGraphics: symbol CGEventTapIsEnabled not loaded")
	}
	return _cGEventTapIsEnabled(tap)
}

var _cGEventTapPostEvent func(proxy CGEventTapProxy, event CGEventRef)

// CGEventTapPostEvent posts a Quartz event from an event tap into the event stream.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/tapPostEvent(_:)
func CGEventTapPostEvent(proxy CGEventTapProxy, event CGEventRef) {
	if _cGEventTapPostEvent == nil {
		panic("CoreGraphics: symbol CGEventTapPostEvent not loaded")
	}
	_cGEventTapPostEvent(proxy, event)
}

var _cGFontCanCreatePostScriptSubset func(font CGFontRef, format CGFontPostScriptFormat) bool

// CGFontCanCreatePostScriptSubset determines whether Core Graphics can create a subset of the font in PostScript format.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/canCreatePostScriptSubset(_:)
func CGFontCanCreatePostScriptSubset(font CGFontRef, format CGFontPostScriptFormat) bool {
	if _cGFontCanCreatePostScriptSubset == nil {
		panic("CoreGraphics: symbol CGFontCanCreatePostScriptSubset not loaded")
	}
	return _cGFontCanCreatePostScriptSubset(font, format)
}

var _cGFontCopyFullName func(font CGFontRef) corefoundation.CFStringRef

// CGFontCopyFullName returns the full name associated with a font object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/fullName
func CGFontCopyFullName(font CGFontRef) corefoundation.CFStringRef {
	if _cGFontCopyFullName == nil {
		panic("CoreGraphics: symbol CGFontCopyFullName not loaded")
	}
	return _cGFontCopyFullName(font)
}

var _cGFontCopyGlyphNameForGlyph func(font CGFontRef, glyph uintptr) corefoundation.CFStringRef

// CGFontCopyGlyphNameForGlyph returns the glyph name of the specified glyph in the specified font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/name(for:)
func CGFontCopyGlyphNameForGlyph(font CGFontRef, glyph uintptr) corefoundation.CFStringRef {
	if _cGFontCopyGlyphNameForGlyph == nil {
		panic("CoreGraphics: symbol CGFontCopyGlyphNameForGlyph not loaded")
	}
	return _cGFontCopyGlyphNameForGlyph(font, glyph)
}

var _cGFontCopyPostScriptName func(font CGFontRef) corefoundation.CFStringRef

// CGFontCopyPostScriptName obtains the PostScript name of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/postScriptName
func CGFontCopyPostScriptName(font CGFontRef) corefoundation.CFStringRef {
	if _cGFontCopyPostScriptName == nil {
		panic("CoreGraphics: symbol CGFontCopyPostScriptName not loaded")
	}
	return _cGFontCopyPostScriptName(font)
}

var _cGFontCopyTableForTag func(font CGFontRef, tag uint32) corefoundation.CFDataRef

// CGFontCopyTableForTag returns the font table that corresponds to the provided tag.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/table(for:)
func CGFontCopyTableForTag(font CGFontRef, tag uint32) corefoundation.CFDataRef {
	if _cGFontCopyTableForTag == nil {
		panic("CoreGraphics: symbol CGFontCopyTableForTag not loaded")
	}
	return _cGFontCopyTableForTag(font, tag)
}

var _cGFontCopyTableTags func(font CGFontRef) corefoundation.CFArrayRef

// CGFontCopyTableTags returns an array of tags that correspond to the font tables for a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/tableTags
func CGFontCopyTableTags(font CGFontRef) corefoundation.CFArrayRef {
	if _cGFontCopyTableTags == nil {
		panic("CoreGraphics: symbol CGFontCopyTableTags not loaded")
	}
	return _cGFontCopyTableTags(font)
}

var _cGFontCopyVariationAxes func(font CGFontRef) corefoundation.CFArrayRef

// CGFontCopyVariationAxes returns an array of the variation axis dictionaries for a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/variationAxes
func CGFontCopyVariationAxes(font CGFontRef) corefoundation.CFArrayRef {
	if _cGFontCopyVariationAxes == nil {
		panic("CoreGraphics: symbol CGFontCopyVariationAxes not loaded")
	}
	return _cGFontCopyVariationAxes(font)
}

var _cGFontCopyVariations func(font CGFontRef) corefoundation.CFDictionaryRef

// CGFontCopyVariations returns the variation specification dictionary for a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/variations
func CGFontCopyVariations(font CGFontRef) corefoundation.CFDictionaryRef {
	if _cGFontCopyVariations == nil {
		panic("CoreGraphics: symbol CGFontCopyVariations not loaded")
	}
	return _cGFontCopyVariations(font)
}

var _cGFontCreateCopyWithVariations func(font CGFontRef, variations corefoundation.CFDictionaryRef) CGFontRef

// CGFontCreateCopyWithVariations creates a copy of a font using a variation specification dictionary.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/copy(withVariations:)
func CGFontCreateCopyWithVariations(font CGFontRef, variations corefoundation.CFDictionaryRef) CGFontRef {
	if _cGFontCreateCopyWithVariations == nil {
		panic("CoreGraphics: symbol CGFontCreateCopyWithVariations not loaded")
	}
	return _cGFontCreateCopyWithVariations(font, variations)
}

var _cGFontCreatePostScriptEncoding func(font CGFontRef, encoding uintptr) corefoundation.CFDataRef

// CGFontCreatePostScriptEncoding creates a PostScript encoding of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/createPostScriptEncoding(encoding:)
func CGFontCreatePostScriptEncoding(font CGFontRef, encoding uintptr) corefoundation.CFDataRef {
	if _cGFontCreatePostScriptEncoding == nil {
		panic("CoreGraphics: symbol CGFontCreatePostScriptEncoding not loaded")
	}
	return _cGFontCreatePostScriptEncoding(font, encoding)
}

var _cGFontCreatePostScriptSubset func(font CGFontRef, subsetName corefoundation.CFStringRef, format CGFontPostScriptFormat, glyphs uintptr, count uintptr, encoding uintptr) corefoundation.CFDataRef

// CGFontCreatePostScriptSubset creates a subset of the font in the specified PostScript format.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/createPostScriptSubset(subsetName:format:glyphs:count:encoding:)
func CGFontCreatePostScriptSubset(font CGFontRef, subsetName corefoundation.CFStringRef, format CGFontPostScriptFormat, glyphs uintptr, count uintptr, encoding uintptr) corefoundation.CFDataRef {
	if _cGFontCreatePostScriptSubset == nil {
		panic("CoreGraphics: symbol CGFontCreatePostScriptSubset not loaded")
	}
	return _cGFontCreatePostScriptSubset(font, subsetName, format, glyphs, count, encoding)
}

var _cGFontCreateWithDataProvider func(provider CGDataProviderRef) CGFontRef

// CGFontCreateWithDataProvider creates a font object from data supplied from a data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/init(_:)-9aour
func CGFontCreateWithDataProvider(provider CGDataProviderRef) CGFontRef {
	if _cGFontCreateWithDataProvider == nil {
		panic("CoreGraphics: symbol CGFontCreateWithDataProvider not loaded")
	}
	return _cGFontCreateWithDataProvider(provider)
}

var _cGFontCreateWithFontName func(name corefoundation.CFStringRef) CGFontRef

// CGFontCreateWithFontName creates a font object corresponding to the font specified by a PostScript or full name.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/init(_:)-1p4b
func CGFontCreateWithFontName(name corefoundation.CFStringRef) CGFontRef {
	if _cGFontCreateWithFontName == nil {
		panic("CoreGraphics: symbol CGFontCreateWithFontName not loaded")
	}
	return _cGFontCreateWithFontName(name)
}

var _cGFontGetAscent func(font CGFontRef) int

// CGFontGetAscent returns the ascent of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/ascent
func CGFontGetAscent(font CGFontRef) int {
	if _cGFontGetAscent == nil {
		panic("CoreGraphics: symbol CGFontGetAscent not loaded")
	}
	return _cGFontGetAscent(font)
}

var _cGFontGetCapHeight func(font CGFontRef) int

// CGFontGetCapHeight returns the cap height of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/capHeight
func CGFontGetCapHeight(font CGFontRef) int {
	if _cGFontGetCapHeight == nil {
		panic("CoreGraphics: symbol CGFontGetCapHeight not loaded")
	}
	return _cGFontGetCapHeight(font)
}

var _cGFontGetDescent func(font CGFontRef) int

// CGFontGetDescent returns the descent of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/descent
func CGFontGetDescent(font CGFontRef) int {
	if _cGFontGetDescent == nil {
		panic("CoreGraphics: symbol CGFontGetDescent not loaded")
	}
	return _cGFontGetDescent(font)
}

var _cGFontGetFontBBox func(font CGFontRef) corefoundation.CGRect

// CGFontGetFontBBox returns the bounding box of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/fontBBox
func CGFontGetFontBBox(font CGFontRef) corefoundation.CGRect {
	if _cGFontGetFontBBox == nil {
		panic("CoreGraphics: symbol CGFontGetFontBBox not loaded")
	}
	return _cGFontGetFontBBox(font)
}

var _cGFontGetGlyphAdvances func(font CGFontRef, glyphs uintptr, count uintptr, advances *int) bool

// CGFontGetGlyphAdvances gets the advance width of each glyph in the provided array.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/getGlyphAdvances(glyphs:count:advances:)
func CGFontGetGlyphAdvances(font CGFontRef, glyphs uintptr, count uintptr, advances []int) bool {
	if _cGFontGetGlyphAdvances == nil {
		panic("CoreGraphics: symbol CGFontGetGlyphAdvances not loaded")
	}
	return _cGFontGetGlyphAdvances(font, glyphs, count, unsafe.SliceData(advances))
}

var _cGFontGetGlyphBBoxes func(font CGFontRef, glyphs uintptr, count uintptr, bboxes *corefoundation.CGRect) bool

// CGFontGetGlyphBBoxes get the bounding box of each glyph in an array.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/getGlyphBBoxes(glyphs:count:bboxes:)
func CGFontGetGlyphBBoxes(font CGFontRef, glyphs uintptr, count uintptr, bboxes *corefoundation.CGRect) bool {
	if _cGFontGetGlyphBBoxes == nil {
		panic("CoreGraphics: symbol CGFontGetGlyphBBoxes not loaded")
	}
	return _cGFontGetGlyphBBoxes(font, glyphs, count, bboxes)
}

var _cGFontGetGlyphWithGlyphName func(font CGFontRef, name corefoundation.CFStringRef) objectivec.IObject

// CGFontGetGlyphWithGlyphName returns the glyph for the glyph name associated with the specified font object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/getGlyphWithGlyphName(name:)
func CGFontGetGlyphWithGlyphName(font CGFontRef, name corefoundation.CFStringRef) objectivec.IObject {
	if _cGFontGetGlyphWithGlyphName == nil {
		panic("CoreGraphics: symbol CGFontGetGlyphWithGlyphName not loaded")
	}
	return _cGFontGetGlyphWithGlyphName(font, name)
}

var _cGFontGetItalicAngle func(font CGFontRef) float64

// CGFontGetItalicAngle returns the italic angle of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/italicAngle
func CGFontGetItalicAngle(font CGFontRef) float64 {
	if _cGFontGetItalicAngle == nil {
		panic("CoreGraphics: symbol CGFontGetItalicAngle not loaded")
	}
	return _cGFontGetItalicAngle(font)
}

var _cGFontGetLeading func(font CGFontRef) int

// CGFontGetLeading returns the leading of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/leading
func CGFontGetLeading(font CGFontRef) int {
	if _cGFontGetLeading == nil {
		panic("CoreGraphics: symbol CGFontGetLeading not loaded")
	}
	return _cGFontGetLeading(font)
}

var _cGFontGetNumberOfGlyphs func(font CGFontRef) uintptr

// CGFontGetNumberOfGlyphs returns the number of glyphs in a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/numberOfGlyphs
func CGFontGetNumberOfGlyphs(font CGFontRef) uintptr {
	if _cGFontGetNumberOfGlyphs == nil {
		panic("CoreGraphics: symbol CGFontGetNumberOfGlyphs not loaded")
	}
	return _cGFontGetNumberOfGlyphs(font)
}

var _cGFontGetStemV func(font CGFontRef) float64

// CGFontGetStemV returns the thickness of the dominant vertical stems of glyphs in a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/stemV
func CGFontGetStemV(font CGFontRef) float64 {
	if _cGFontGetStemV == nil {
		panic("CoreGraphics: symbol CGFontGetStemV not loaded")
	}
	return _cGFontGetStemV(font)
}

var _cGFontGetTypeID func() uint

// CGFontGetTypeID returns the Core Foundation type identifier for Core Graphics fonts.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/typeID
func CGFontGetTypeID() uint {
	if _cGFontGetTypeID == nil {
		panic("CoreGraphics: symbol CGFontGetTypeID not loaded")
	}
	return _cGFontGetTypeID()
}

var _cGFontGetUnitsPerEm func(font CGFontRef) int

// CGFontGetUnitsPerEm returns the number of glyph space units per em for the provided font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/unitsPerEm
func CGFontGetUnitsPerEm(font CGFontRef) int {
	if _cGFontGetUnitsPerEm == nil {
		panic("CoreGraphics: symbol CGFontGetUnitsPerEm not loaded")
	}
	return _cGFontGetUnitsPerEm(font)
}

var _cGFontGetXHeight func(font CGFontRef) int

// CGFontGetXHeight returns the x-height of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/xHeight
func CGFontGetXHeight(font CGFontRef) int {
	if _cGFontGetXHeight == nil {
		panic("CoreGraphics: symbol CGFontGetXHeight not loaded")
	}
	return _cGFontGetXHeight(font)
}

var _cGFontRelease func(font CGFontRef)

// CGFontRelease decrements the retain count of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFontRelease
func CGFontRelease(font CGFontRef) {
	if _cGFontRelease == nil {
		panic("CoreGraphics: symbol CGFontRelease not loaded")
	}
	_cGFontRelease(font)
}

var _cGFontRetain func(font CGFontRef) CGFontRef

// CGFontRetain increments the retain count of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFontRetain
func CGFontRetain(font CGFontRef) CGFontRef {
	if _cGFontRetain == nil {
		panic("CoreGraphics: symbol CGFontRetain not loaded")
	}
	return _cGFontRetain(font)
}

var _cGFunctionCreate func(info unsafe.Pointer, domainDimension uintptr, domain *float64, rangeDimension uintptr, range_ *float64, callbacks *CGFunctionCallbacks) CGFunctionRef

// CGFunctionCreate creates a Core Graphics function.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFunction/init(info:domainDimension:domain:rangeDimension:range:callbacks:)
func CGFunctionCreate(info unsafe.Pointer, domainDimension uintptr, domain *float64, rangeDimension uintptr, range_ *float64, callbacks *CGFunctionCallbacks) CGFunctionRef {
	if _cGFunctionCreate == nil {
		panic("CoreGraphics: symbol CGFunctionCreate not loaded")
	}
	return _cGFunctionCreate(info, domainDimension, domain, rangeDimension, range_, callbacks)
}

var _cGFunctionGetTypeID func() uint

// CGFunctionGetTypeID returns the type identifier for Core Graphics function objects.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFunction/typeID
func CGFunctionGetTypeID() uint {
	if _cGFunctionGetTypeID == nil {
		panic("CoreGraphics: symbol CGFunctionGetTypeID not loaded")
	}
	return _cGFunctionGetTypeID()
}

var _cGFunctionRelease func(function CGFunctionRef)

// CGFunctionRelease decrements the retain count of a function object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFunctionRelease
func CGFunctionRelease(function CGFunctionRef) {
	if _cGFunctionRelease == nil {
		panic("CoreGraphics: symbol CGFunctionRelease not loaded")
	}
	_cGFunctionRelease(function)
}

var _cGFunctionRetain func(function CGFunctionRef) CGFunctionRef

// CGFunctionRetain increments the retain count of a function object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFunctionRetain
func CGFunctionRetain(function CGFunctionRef) CGFunctionRef {
	if _cGFunctionRetain == nil {
		panic("CoreGraphics: symbol CGFunctionRetain not loaded")
	}
	return _cGFunctionRetain(function)
}

var _cGGetActiveDisplayList func(maxDisplays uint32, activeDisplays *uint32, displayCount *uint32) CGError

// CGGetActiveDisplayList provides a list of displays that are active for drawing.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGetActiveDisplayList(_:_:_:)
func CGGetActiveDisplayList(maxDisplays uint32, activeDisplays *uint32, displayCount *uint32) CGError {
	if _cGGetActiveDisplayList == nil {
		panic("CoreGraphics: symbol CGGetActiveDisplayList not loaded")
	}
	return _cGGetActiveDisplayList(maxDisplays, activeDisplays, displayCount)
}

var _cGGetDisplayTransferByFormula func(display uint32, redMin *CGGammaValue, redMax *CGGammaValue, redGamma *CGGammaValue, greenMin *CGGammaValue, greenMax *CGGammaValue, greenGamma *CGGammaValue, blueMin *CGGammaValue, blueMax *CGGammaValue, blueGamma *CGGammaValue) CGError

// CGGetDisplayTransferByFormula gets the coefficients of the gamma transfer formula for a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGetDisplayTransferByFormula(_:_:_:_:_:_:_:_:_:_:)
func CGGetDisplayTransferByFormula(display uint32, redMin *CGGammaValue, redMax *CGGammaValue, redGamma *CGGammaValue, greenMin *CGGammaValue, greenMax *CGGammaValue, greenGamma *CGGammaValue, blueMin *CGGammaValue, blueMax *CGGammaValue, blueGamma *CGGammaValue) CGError {
	if _cGGetDisplayTransferByFormula == nil {
		panic("CoreGraphics: symbol CGGetDisplayTransferByFormula not loaded")
	}
	return _cGGetDisplayTransferByFormula(display, redMin, redMax, redGamma, greenMin, greenMax, greenGamma, blueMin, blueMax, blueGamma)
}

var _cGGetDisplayTransferByTable func(display uint32, capacity uint32, redTable *CGGammaValue, greenTable *CGGammaValue, blueTable *CGGammaValue, sampleCount *uint32) CGError

// CGGetDisplayTransferByTable gets the values in the RGB gamma tables for a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGetDisplayTransferByTable(_:_:_:_:_:_:)
func CGGetDisplayTransferByTable(display uint32, capacity uint32, redTable *CGGammaValue, greenTable *CGGammaValue, blueTable *CGGammaValue, sampleCount *uint32) CGError {
	if _cGGetDisplayTransferByTable == nil {
		panic("CoreGraphics: symbol CGGetDisplayTransferByTable not loaded")
	}
	return _cGGetDisplayTransferByTable(display, capacity, redTable, greenTable, blueTable, sampleCount)
}

var _cGGetDisplaysWithOpenGLDisplayMask func(mask CGOpenGLDisplayMask, maxDisplays uint32, displays *uint32, matchingDisplayCount *uint32) CGError

// CGGetDisplaysWithOpenGLDisplayMask provides a list of displays that corresponds to the bits set in an OpenGL display mask.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGetDisplaysWithOpenGLDisplayMask(_:_:_:_:)
func CGGetDisplaysWithOpenGLDisplayMask(mask CGOpenGLDisplayMask, maxDisplays uint32, displays *uint32, matchingDisplayCount *uint32) CGError {
	if _cGGetDisplaysWithOpenGLDisplayMask == nil {
		panic("CoreGraphics: symbol CGGetDisplaysWithOpenGLDisplayMask not loaded")
	}
	return _cGGetDisplaysWithOpenGLDisplayMask(mask, maxDisplays, displays, matchingDisplayCount)
}

var _cGGetDisplaysWithPoint func(point corefoundation.CGPoint, maxDisplays uint32, displays *uint32, matchingDisplayCount *uint32) CGError

// CGGetDisplaysWithPoint provides a list of online displays with bounds that include the specified point.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGetDisplaysWithPoint(_:_:_:_:)
func CGGetDisplaysWithPoint(point corefoundation.CGPoint, maxDisplays uint32, displays *uint32, matchingDisplayCount *uint32) CGError {
	if _cGGetDisplaysWithPoint == nil {
		panic("CoreGraphics: symbol CGGetDisplaysWithPoint not loaded")
	}
	return _cGGetDisplaysWithPoint(point, maxDisplays, displays, matchingDisplayCount)
}

var _cGGetDisplaysWithRect func(rect corefoundation.CGRect, maxDisplays uint32, displays *uint32, matchingDisplayCount *uint32) CGError

// CGGetDisplaysWithRect gets a list of online displays with bounds that intersect the specified rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGetDisplaysWithRect(_:_:_:_:)
func CGGetDisplaysWithRect(rect corefoundation.CGRect, maxDisplays uint32, displays *uint32, matchingDisplayCount *uint32) CGError {
	if _cGGetDisplaysWithRect == nil {
		panic("CoreGraphics: symbol CGGetDisplaysWithRect not loaded")
	}
	return _cGGetDisplaysWithRect(rect, maxDisplays, displays, matchingDisplayCount)
}

var _cGGetEventTapList func(maxNumberOfTaps uint32, tapList *CGEventTapInformation, eventTapCount *uint32) CGError

// CGGetEventTapList gets a list of currently installed event taps.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGetEventTapList(_:_:_:)
func CGGetEventTapList(maxNumberOfTaps uint32, tapList *CGEventTapInformation, eventTapCount *uint32) CGError {
	if _cGGetEventTapList == nil {
		panic("CoreGraphics: symbol CGGetEventTapList not loaded")
	}
	return _cGGetEventTapList(maxNumberOfTaps, tapList, eventTapCount)
}

var _cGGetLastMouseDelta func(deltaX *int32, deltaY *int32)

// CGGetLastMouseDelta reports the change in mouse position since the last mouse movement event received by the application.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGetLastMouseDelta
func CGGetLastMouseDelta(deltaX *int32, deltaY *int32) {
	if _cGGetLastMouseDelta == nil {
		panic("CoreGraphics: symbol CGGetLastMouseDelta not loaded")
	}
	_cGGetLastMouseDelta(deltaX, deltaY)
}

var _cGGetOnlineDisplayList func(maxDisplays uint32, onlineDisplays *uint32, displayCount *uint32) CGError

// CGGetOnlineDisplayList provides a list of displays that are online (active, mirrored, or sleeping).
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGetOnlineDisplayList(_:_:_:)
func CGGetOnlineDisplayList(maxDisplays uint32, onlineDisplays *uint32, displayCount *uint32) CGError {
	if _cGGetOnlineDisplayList == nil {
		panic("CoreGraphics: symbol CGGetOnlineDisplayList not loaded")
	}
	return _cGGetOnlineDisplayList(maxDisplays, onlineDisplays, displayCount)
}

var _cGGradientCreateWithColorComponents func(space CGColorSpaceRef, components *float64, locations *float64, count uintptr) CGGradientRef

// CGGradientCreateWithColorComponents creates a CGGradient object from a color space and the provided color components and locations.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGradient/init(colorSpace:colorComponents:locations:count:)
func CGGradientCreateWithColorComponents(space CGColorSpaceRef, components *float64, locations *float64, count uintptr) CGGradientRef {
	if _cGGradientCreateWithColorComponents == nil {
		panic("CoreGraphics: symbol CGGradientCreateWithColorComponents not loaded")
	}
	return _cGGradientCreateWithColorComponents(space, components, locations, count)
}

var _cGGradientCreateWithColors func(space CGColorSpaceRef, colors corefoundation.CFArrayRef, locations *float64) CGGradientRef

// CGGradientCreateWithColors creates a gradient object from a color space and the provided color objects and locations.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGradient/init(colorsSpace:colors:locations:)
func CGGradientCreateWithColors(space CGColorSpaceRef, colors corefoundation.CFArrayRef, locations *float64) CGGradientRef {
	if _cGGradientCreateWithColors == nil {
		panic("CoreGraphics: symbol CGGradientCreateWithColors not loaded")
	}
	return _cGGradientCreateWithColors(space, colors, locations)
}

var _cGGradientCreateWithContentHeadroom func(headroom float32, space CGColorSpaceRef, components *float64, locations *float64, count uintptr) CGGradientRef

// CGGradientCreateWithContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGradient/init(headroom:colorSpace:colorComponents:locations:count:)
func CGGradientCreateWithContentHeadroom(headroom float32, space CGColorSpaceRef, components *float64, locations *float64, count uintptr) CGGradientRef {
	if _cGGradientCreateWithContentHeadroom == nil {
		panic("CoreGraphics: symbol CGGradientCreateWithContentHeadroom not loaded")
	}
	return _cGGradientCreateWithContentHeadroom(headroom, space, components, locations, count)
}

var _cGGradientGetContentHeadroom func(gradient CGGradientRef) float32

// CGGradientGetContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGradient/contentHeadroom
func CGGradientGetContentHeadroom(gradient CGGradientRef) float32 {
	if _cGGradientGetContentHeadroom == nil {
		panic("CoreGraphics: symbol CGGradientGetContentHeadroom not loaded")
	}
	return _cGGradientGetContentHeadroom(gradient)
}

var _cGGradientGetTypeID func() uint

// CGGradientGetTypeID returns the Core Foundation type identifier for CGGradient objects.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGradient/typeID
func CGGradientGetTypeID() uint {
	if _cGGradientGetTypeID == nil {
		panic("CoreGraphics: symbol CGGradientGetTypeID not loaded")
	}
	return _cGGradientGetTypeID()
}

var _cGGradientRelease func(gradient CGGradientRef)

// CGGradientRelease decrements the retain count of a CGGradient object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGradientRelease
func CGGradientRelease(gradient CGGradientRef) {
	if _cGGradientRelease == nil {
		panic("CoreGraphics: symbol CGGradientRelease not loaded")
	}
	_cGGradientRelease(gradient)
}

var _cGGradientRetain func(gradient CGGradientRef) CGGradientRef

// CGGradientRetain increments the retain count of a CGGradient object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGradientRetain
func CGGradientRetain(gradient CGGradientRef) CGGradientRef {
	if _cGGradientRetain == nil {
		panic("CoreGraphics: symbol CGGradientRetain not loaded")
	}
	return _cGGradientRetain(gradient)
}

var _cGImageCalculateContentAverageLightLevel func(image CGImageRef) float32

// CGImageCalculateContentAverageLightLevel.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/calculatedContentAverageLightLevel
func CGImageCalculateContentAverageLightLevel(image CGImageRef) float32 {
	if _cGImageCalculateContentAverageLightLevel == nil {
		panic("CoreGraphics: symbol CGImageCalculateContentAverageLightLevel not loaded")
	}
	return _cGImageCalculateContentAverageLightLevel(image)
}

var _cGImageCalculateContentHeadroom func(image CGImageRef) float32

// CGImageCalculateContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/calculatedContentHeadroom
func CGImageCalculateContentHeadroom(image CGImageRef) float32 {
	if _cGImageCalculateContentHeadroom == nil {
		panic("CoreGraphics: symbol CGImageCalculateContentHeadroom not loaded")
	}
	return _cGImageCalculateContentHeadroom(image)
}

var _cGImageContainsImageSpecificToneMappingMetadata func(image CGImageRef) bool

// CGImageContainsImageSpecificToneMappingMetadata.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/containsImageSpecificToneMappingMetadata
func CGImageContainsImageSpecificToneMappingMetadata(image CGImageRef) bool {
	if _cGImageContainsImageSpecificToneMappingMetadata == nil {
		panic("CoreGraphics: symbol CGImageContainsImageSpecificToneMappingMetadata not loaded")
	}
	return _cGImageContainsImageSpecificToneMappingMetadata(image)
}

var _cGImageCreate func(width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo, provider CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) CGImageRef

// CGImageCreate creates a bitmap image from data supplied by a data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/init(width:height:bitsPerComponent:bitsPerPixel:bytesPerRow:space:bitmapInfo:provider:decode:shouldInterpolate:intent:)
func CGImageCreate(width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo, provider CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) CGImageRef {
	if _cGImageCreate == nil {
		panic("CoreGraphics: symbol CGImageCreate not loaded")
	}
	return _cGImageCreate(width, height, bitsPerComponent, bitsPerPixel, bytesPerRow, space, bitmapInfo, provider, decode, shouldInterpolate, intent)
}

var _cGImageCreateCopy func(image CGImageRef) CGImageRef

// CGImageCreateCopy creates a copy of a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/copy()
func CGImageCreateCopy(image CGImageRef) CGImageRef {
	if _cGImageCreateCopy == nil {
		panic("CoreGraphics: symbol CGImageCreateCopy not loaded")
	}
	return _cGImageCreateCopy(image)
}

var _cGImageCreateCopyWithCalculatedHDRStats func(image CGImageRef) CGImageRef

// CGImageCreateCopyWithCalculatedHDRStats.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/copyWithCalculatedHDRStats()
func CGImageCreateCopyWithCalculatedHDRStats(image CGImageRef) CGImageRef {
	if _cGImageCreateCopyWithCalculatedHDRStats == nil {
		panic("CoreGraphics: symbol CGImageCreateCopyWithCalculatedHDRStats not loaded")
	}
	return _cGImageCreateCopyWithCalculatedHDRStats(image)
}

var _cGImageCreateCopyWithColorSpace func(image CGImageRef, space CGColorSpaceRef) CGImageRef

// CGImageCreateCopyWithColorSpace creates a copy of a bitmap image, replacing its colorspace.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/copy(colorSpace:)
func CGImageCreateCopyWithColorSpace(image CGImageRef, space CGColorSpaceRef) CGImageRef {
	if _cGImageCreateCopyWithColorSpace == nil {
		panic("CoreGraphics: symbol CGImageCreateCopyWithColorSpace not loaded")
	}
	return _cGImageCreateCopyWithColorSpace(image, space)
}

var _cGImageCreateCopyWithContentAverageLightLevel func(image CGImageRef, avll float32) CGImageRef

// CGImageCreateCopyWithContentAverageLightLevel.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/copy(contentAverageLightLevel:)
func CGImageCreateCopyWithContentAverageLightLevel(image CGImageRef, avll float32) CGImageRef {
	if _cGImageCreateCopyWithContentAverageLightLevel == nil {
		panic("CoreGraphics: symbol CGImageCreateCopyWithContentAverageLightLevel not loaded")
	}
	return _cGImageCreateCopyWithContentAverageLightLevel(image, avll)
}

var _cGImageCreateCopyWithContentHeadroom func(headroom float32, image CGImageRef) CGImageRef

// CGImageCreateCopyWithContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImageCreateCopyWithContentHeadroom(_:_:)
func CGImageCreateCopyWithContentHeadroom(headroom float32, image CGImageRef) CGImageRef {
	if _cGImageCreateCopyWithContentHeadroom == nil {
		panic("CoreGraphics: symbol CGImageCreateCopyWithContentHeadroom not loaded")
	}
	return _cGImageCreateCopyWithContentHeadroom(headroom, image)
}

var _cGImageCreateWithContentHeadroom func(headroom float32, width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo, provider CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) CGImageRef

// CGImageCreateWithContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/init(headroom:width:height:bitsPerComponent:bitsPerPixel:bytesPerRow:space:bitmapInfo:provider:decode:shouldInterpolate:intent:)
func CGImageCreateWithContentHeadroom(headroom float32, width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo, provider CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) CGImageRef {
	if _cGImageCreateWithContentHeadroom == nil {
		panic("CoreGraphics: symbol CGImageCreateWithContentHeadroom not loaded")
	}
	return _cGImageCreateWithContentHeadroom(headroom, width, height, bitsPerComponent, bitsPerPixel, bytesPerRow, space, bitmapInfo, provider, decode, shouldInterpolate, intent)
}

var _cGImageCreateWithImageInRect func(image CGImageRef, rect corefoundation.CGRect) CGImageRef

// CGImageCreateWithImageInRect creates a bitmap image using the data contained within a subregion of an existing bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/cropping(to:)
func CGImageCreateWithImageInRect(image CGImageRef, rect corefoundation.CGRect) CGImageRef {
	if _cGImageCreateWithImageInRect == nil {
		panic("CoreGraphics: symbol CGImageCreateWithImageInRect not loaded")
	}
	return _cGImageCreateWithImageInRect(image, rect)
}

var _cGImageCreateWithJPEGDataProvider func(source CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) CGImageRef

// CGImageCreateWithJPEGDataProvider creates a bitmap image using JPEG-encoded data supplied by a data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/init(jpegDataProviderSource:decode:shouldInterpolate:intent:)
func CGImageCreateWithJPEGDataProvider(source CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) CGImageRef {
	if _cGImageCreateWithJPEGDataProvider == nil {
		panic("CoreGraphics: symbol CGImageCreateWithJPEGDataProvider not loaded")
	}
	return _cGImageCreateWithJPEGDataProvider(source, decode, shouldInterpolate, intent)
}

var _cGImageCreateWithMask func(image CGImageRef, mask CGImageRef) CGImageRef

// CGImageCreateWithMask creates a bitmap image from an existing image and an image mask.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/masking(_:)
func CGImageCreateWithMask(image CGImageRef, mask CGImageRef) CGImageRef {
	if _cGImageCreateWithMask == nil {
		panic("CoreGraphics: symbol CGImageCreateWithMask not loaded")
	}
	return _cGImageCreateWithMask(image, mask)
}

var _cGImageCreateWithMaskingColors func(image CGImageRef, components *float64) CGImageRef

// CGImageCreateWithMaskingColors creates a bitmap image by masking an existing bitmap image with the provided color values.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImageCreateWithMaskingColors
func CGImageCreateWithMaskingColors(image CGImageRef, components *float64) CGImageRef {
	if _cGImageCreateWithMaskingColors == nil {
		panic("CoreGraphics: symbol CGImageCreateWithMaskingColors not loaded")
	}
	return _cGImageCreateWithMaskingColors(image, components)
}

var _cGImageCreateWithPNGDataProvider func(source CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) CGImageRef

// CGImageCreateWithPNGDataProvider creates a bitmap image using PNG-encoded data supplied by a data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/init(pngDataProviderSource:decode:shouldInterpolate:intent:)
func CGImageCreateWithPNGDataProvider(source CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) CGImageRef {
	if _cGImageCreateWithPNGDataProvider == nil {
		panic("CoreGraphics: symbol CGImageCreateWithPNGDataProvider not loaded")
	}
	return _cGImageCreateWithPNGDataProvider(source, decode, shouldInterpolate, intent)
}

var _cGImageGetAlphaInfo func(image CGImageRef) CGImageAlphaInfo

// CGImageGetAlphaInfo returns the alpha channel information for a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/alphaInfo
func CGImageGetAlphaInfo(image CGImageRef) CGImageAlphaInfo {
	if _cGImageGetAlphaInfo == nil {
		panic("CoreGraphics: symbol CGImageGetAlphaInfo not loaded")
	}
	return _cGImageGetAlphaInfo(image)
}

var _cGImageGetBitmapInfo func(image CGImageRef) CGBitmapInfo

// CGImageGetBitmapInfo returns the bitmap information for a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/bitmapInfo
func CGImageGetBitmapInfo(image CGImageRef) CGBitmapInfo {
	if _cGImageGetBitmapInfo == nil {
		panic("CoreGraphics: symbol CGImageGetBitmapInfo not loaded")
	}
	return _cGImageGetBitmapInfo(image)
}

var _cGImageGetBitsPerComponent func(image CGImageRef) uintptr

// CGImageGetBitsPerComponent returns the number of bits allocated for a single color component of a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/bitsPerComponent
func CGImageGetBitsPerComponent(image CGImageRef) uintptr {
	if _cGImageGetBitsPerComponent == nil {
		panic("CoreGraphics: symbol CGImageGetBitsPerComponent not loaded")
	}
	return _cGImageGetBitsPerComponent(image)
}

var _cGImageGetBitsPerPixel func(image CGImageRef) uintptr

// CGImageGetBitsPerPixel returns the number of bits allocated for a single pixel in a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/bitsPerPixel
func CGImageGetBitsPerPixel(image CGImageRef) uintptr {
	if _cGImageGetBitsPerPixel == nil {
		panic("CoreGraphics: symbol CGImageGetBitsPerPixel not loaded")
	}
	return _cGImageGetBitsPerPixel(image)
}

var _cGImageGetByteOrderInfo func(image CGImageRef) CGImageByteOrderInfo

// CGImageGetByteOrderInfo.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/byteOrderInfo
func CGImageGetByteOrderInfo(image CGImageRef) CGImageByteOrderInfo {
	if _cGImageGetByteOrderInfo == nil {
		panic("CoreGraphics: symbol CGImageGetByteOrderInfo not loaded")
	}
	return _cGImageGetByteOrderInfo(image)
}

var _cGImageGetBytesPerRow func(image CGImageRef) uintptr

// CGImageGetBytesPerRow returns the number of bytes allocated for a single row of a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/bytesPerRow
func CGImageGetBytesPerRow(image CGImageRef) uintptr {
	if _cGImageGetBytesPerRow == nil {
		panic("CoreGraphics: symbol CGImageGetBytesPerRow not loaded")
	}
	return _cGImageGetBytesPerRow(image)
}

var _cGImageGetColorSpace func(image CGImageRef) CGColorSpaceRef

// CGImageGetColorSpace return the color space for a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/colorSpace
func CGImageGetColorSpace(image CGImageRef) CGColorSpaceRef {
	if _cGImageGetColorSpace == nil {
		panic("CoreGraphics: symbol CGImageGetColorSpace not loaded")
	}
	return _cGImageGetColorSpace(image)
}

var _cGImageGetContentAverageLightLevel func(image CGImageRef) float32

// CGImageGetContentAverageLightLevel.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/contentAverageLightLevel
func CGImageGetContentAverageLightLevel(image CGImageRef) float32 {
	if _cGImageGetContentAverageLightLevel == nil {
		panic("CoreGraphics: symbol CGImageGetContentAverageLightLevel not loaded")
	}
	return _cGImageGetContentAverageLightLevel(image)
}

var _cGImageGetContentHeadroom func(image CGImageRef) float32

// CGImageGetContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/contentHeadroom
func CGImageGetContentHeadroom(image CGImageRef) float32 {
	if _cGImageGetContentHeadroom == nil {
		panic("CoreGraphics: symbol CGImageGetContentHeadroom not loaded")
	}
	return _cGImageGetContentHeadroom(image)
}

var _cGImageGetDataProvider func(image CGImageRef) CGDataProviderRef

// CGImageGetDataProvider returns the data provider for a bitmap image or image mask.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/dataProvider
func CGImageGetDataProvider(image CGImageRef) CGDataProviderRef {
	if _cGImageGetDataProvider == nil {
		panic("CoreGraphics: symbol CGImageGetDataProvider not loaded")
	}
	return _cGImageGetDataProvider(image)
}

var _cGImageGetDecode func(image CGImageRef) *float64

// CGImageGetDecode returns the decode array for a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/decode
func CGImageGetDecode(image CGImageRef) *float64 {
	if _cGImageGetDecode == nil {
		panic("CoreGraphics: symbol CGImageGetDecode not loaded")
	}
	return _cGImageGetDecode(image)
}

var _cGImageGetHeight func(image CGImageRef) uintptr

// CGImageGetHeight returns the height of a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/height
func CGImageGetHeight(image CGImageRef) uintptr {
	if _cGImageGetHeight == nil {
		panic("CoreGraphics: symbol CGImageGetHeight not loaded")
	}
	return _cGImageGetHeight(image)
}

var _cGImageGetPixelFormatInfo func(image CGImageRef) CGImagePixelFormatInfo

// CGImageGetPixelFormatInfo.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/pixelFormatInfo
func CGImageGetPixelFormatInfo(image CGImageRef) CGImagePixelFormatInfo {
	if _cGImageGetPixelFormatInfo == nil {
		panic("CoreGraphics: symbol CGImageGetPixelFormatInfo not loaded")
	}
	return _cGImageGetPixelFormatInfo(image)
}

var _cGImageGetRenderingIntent func(image CGImageRef) CGColorRenderingIntent

// CGImageGetRenderingIntent returns the rendering intent setting for a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/renderingIntent
func CGImageGetRenderingIntent(image CGImageRef) CGColorRenderingIntent {
	if _cGImageGetRenderingIntent == nil {
		panic("CoreGraphics: symbol CGImageGetRenderingIntent not loaded")
	}
	return _cGImageGetRenderingIntent(image)
}

var _cGImageGetShouldInterpolate func(image CGImageRef) bool

// CGImageGetShouldInterpolate returns the interpolation setting for a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/shouldInterpolate
func CGImageGetShouldInterpolate(image CGImageRef) bool {
	if _cGImageGetShouldInterpolate == nil {
		panic("CoreGraphics: symbol CGImageGetShouldInterpolate not loaded")
	}
	return _cGImageGetShouldInterpolate(image)
}

var _cGImageGetTypeID func() uint

// CGImageGetTypeID returns the type identifier for CGImage objects.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/typeID
func CGImageGetTypeID() uint {
	if _cGImageGetTypeID == nil {
		panic("CoreGraphics: symbol CGImageGetTypeID not loaded")
	}
	return _cGImageGetTypeID()
}

var _cGImageGetUTType func(image CGImageRef) corefoundation.CFStringRef

// CGImageGetUTType the Universal Type Identifier for the image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/utType
func CGImageGetUTType(image CGImageRef) corefoundation.CFStringRef {
	if _cGImageGetUTType == nil {
		panic("CoreGraphics: symbol CGImageGetUTType not loaded")
	}
	return _cGImageGetUTType(image)
}

var _cGImageGetWidth func(image CGImageRef) uintptr

// CGImageGetWidth returns the width of a bitmap image, in pixels.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/width
func CGImageGetWidth(image CGImageRef) uintptr {
	if _cGImageGetWidth == nil {
		panic("CoreGraphics: symbol CGImageGetWidth not loaded")
	}
	return _cGImageGetWidth(image)
}

var _cGImageIsMask func(image CGImageRef) bool

// CGImageIsMask returns whether a bitmap image is an image mask.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/isMask
func CGImageIsMask(image CGImageRef) bool {
	if _cGImageIsMask == nil {
		panic("CoreGraphics: symbol CGImageIsMask not loaded")
	}
	return _cGImageIsMask(image)
}

var _cGImageMaskCreate func(width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, provider CGDataProviderRef, decode *float64, shouldInterpolate bool) CGImageRef

// CGImageMaskCreate creates a bitmap image mask from data supplied by a data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/init(maskWidth:height:bitsPerComponent:bitsPerPixel:bytesPerRow:provider:decode:shouldInterpolate:)
func CGImageMaskCreate(width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, provider CGDataProviderRef, decode *float64, shouldInterpolate bool) CGImageRef {
	if _cGImageMaskCreate == nil {
		panic("CoreGraphics: symbol CGImageMaskCreate not loaded")
	}
	return _cGImageMaskCreate(width, height, bitsPerComponent, bitsPerPixel, bytesPerRow, provider, decode, shouldInterpolate)
}

var _cGImageRelease func(image CGImageRef)

// CGImageRelease decrements the retain count of a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImageRelease
func CGImageRelease(image CGImageRef) {
	if _cGImageRelease == nil {
		panic("CoreGraphics: symbol CGImageRelease not loaded")
	}
	_cGImageRelease(image)
}

var _cGImageRetain func(image CGImageRef) CGImageRef

// CGImageRetain increments the retain count of a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImageRetain
func CGImageRetain(image CGImageRef) CGImageRef {
	if _cGImageRetain == nil {
		panic("CoreGraphics: symbol CGImageRetain not loaded")
	}
	return _cGImageRetain(image)
}

var _cGImageShouldToneMap func(image CGImageRef) bool

// CGImageShouldToneMap.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/shouldToneMap
func CGImageShouldToneMap(image CGImageRef) bool {
	if _cGImageShouldToneMap == nil {
		panic("CoreGraphics: symbol CGImageShouldToneMap not loaded")
	}
	return _cGImageShouldToneMap(image)
}

var _cGLayerCreateWithContext func(context CGContextRef, size corefoundation.CGSize, auxiliaryInfo corefoundation.CFDictionaryRef) objectivec.IObject

// CGLayerCreateWithContext creates a layer object that is associated with a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGLayer/init(_:size:auxiliaryInfo:)
func CGLayerCreateWithContext(context CGContextRef, size corefoundation.CGSize, auxiliaryInfo corefoundation.CFDictionaryRef) objectivec.IObject {
	if _cGLayerCreateWithContext == nil {
		panic("CoreGraphics: symbol CGLayerCreateWithContext not loaded")
	}
	return _cGLayerCreateWithContext(context, size, auxiliaryInfo)
}

var _cGLayerGetContext func(layer uintptr) CGContextRef

// CGLayerGetContext returns the graphics context associated with a layer object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGLayer/context
func CGLayerGetContext(layer uintptr) CGContextRef {
	if _cGLayerGetContext == nil {
		panic("CoreGraphics: symbol CGLayerGetContext not loaded")
	}
	return _cGLayerGetContext(layer)
}

var _cGLayerGetSize func(layer uintptr) corefoundation.CGSize

// CGLayerGetSize returns the width and height of a layer object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGLayer/size
func CGLayerGetSize(layer uintptr) corefoundation.CGSize {
	if _cGLayerGetSize == nil {
		panic("CoreGraphics: symbol CGLayerGetSize not loaded")
	}
	return _cGLayerGetSize(layer)
}

var _cGLayerGetTypeID func() uint

// CGLayerGetTypeID returns the unique type identifier used for CGLayer objects.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGLayer/typeID
func CGLayerGetTypeID() uint {
	if _cGLayerGetTypeID == nil {
		panic("CoreGraphics: symbol CGLayerGetTypeID not loaded")
	}
	return _cGLayerGetTypeID()
}

var _cGLayerRelease func(layer uintptr)

// CGLayerRelease decrements the retain count of a layer object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGLayerRelease
func CGLayerRelease(layer uintptr) {
	if _cGLayerRelease == nil {
		panic("CoreGraphics: symbol CGLayerRelease not loaded")
	}
	_cGLayerRelease(layer)
}

var _cGLayerRetain func(layer uintptr) objectivec.IObject

// CGLayerRetain increments the retain count of a layer object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGLayerRetain
func CGLayerRetain(layer uintptr) objectivec.IObject {
	if _cGLayerRetain == nil {
		panic("CoreGraphics: symbol CGLayerRetain not loaded")
	}
	return _cGLayerRetain(layer)
}

var _cGMainDisplayID func() uint32

// CGMainDisplayID returns the display ID of the main display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGMainDisplayID()
func CGMainDisplayID() uint32 {
	if _cGMainDisplayID == nil {
		panic("CoreGraphics: symbol CGMainDisplayID not loaded")
	}
	return _cGMainDisplayID()
}

var _cGOpenGLDisplayMaskToDisplayID func(mask CGOpenGLDisplayMask) uint32

// CGOpenGLDisplayMaskToDisplayID maps an OpenGL display mask to a display ID.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGOpenGLDisplayMaskToDisplayID(_:)
func CGOpenGLDisplayMaskToDisplayID(mask CGOpenGLDisplayMask) uint32 {
	if _cGOpenGLDisplayMaskToDisplayID == nil {
		panic("CoreGraphics: symbol CGOpenGLDisplayMaskToDisplayID not loaded")
	}
	return _cGOpenGLDisplayMaskToDisplayID(mask)
}

var _cGPDFArrayApplyBlock func(array CGPDFArrayRef, block CGPDFArrayApplierBlock, info unsafe.Pointer)

// CGPDFArrayApplyBlock.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayApplyBlock(_:_:_:)
func CGPDFArrayApplyBlock(array CGPDFArrayRef, block CGPDFArrayApplierBlock, info unsafe.Pointer) {
	if _cGPDFArrayApplyBlock == nil {
		panic("CoreGraphics: symbol CGPDFArrayApplyBlock not loaded")
	}
	_cGPDFArrayApplyBlock(array, block, info)
}

var _cGPDFArrayGetArray func(array CGPDFArrayRef, index uintptr, value *CGPDFArrayRef) bool

// CGPDFArrayGetArray returns whether an object at a given index in a PDF array is another PDF array and, if so, retrieves that array.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetArray(_:_:_:)
func CGPDFArrayGetArray(array CGPDFArrayRef, index uintptr, value *CGPDFArrayRef) bool {
	if _cGPDFArrayGetArray == nil {
		panic("CoreGraphics: symbol CGPDFArrayGetArray not loaded")
	}
	return _cGPDFArrayGetArray(array, index, value)
}

var _cGPDFArrayGetBoolean func(array CGPDFArrayRef, index uintptr, value *CGPDFBoolean) bool

// CGPDFArrayGetBoolean returns whether an object at a given index in a PDF array is a PDF Boolean and, if so, retrieves that Boolean.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetBoolean(_:_:_:)
func CGPDFArrayGetBoolean(array CGPDFArrayRef, index uintptr, value *CGPDFBoolean) bool {
	if _cGPDFArrayGetBoolean == nil {
		panic("CoreGraphics: symbol CGPDFArrayGetBoolean not loaded")
	}
	return _cGPDFArrayGetBoolean(array, index, value)
}

var _cGPDFArrayGetCount func(array CGPDFArrayRef) uintptr

// CGPDFArrayGetCount returns the number of items in a PDF array.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetCount(_:)
func CGPDFArrayGetCount(array CGPDFArrayRef) uintptr {
	if _cGPDFArrayGetCount == nil {
		panic("CoreGraphics: symbol CGPDFArrayGetCount not loaded")
	}
	return _cGPDFArrayGetCount(array)
}

var _cGPDFArrayGetDictionary func(array CGPDFArrayRef, index uintptr, value *CGPDFDictionaryRef) bool

// CGPDFArrayGetDictionary returns whether an object at a given index in a PDF array is a PDF dictionary and, if so, retrieves that dictionary.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetDictionary(_:_:_:)
func CGPDFArrayGetDictionary(array CGPDFArrayRef, index uintptr, value *CGPDFDictionaryRef) bool {
	if _cGPDFArrayGetDictionary == nil {
		panic("CoreGraphics: symbol CGPDFArrayGetDictionary not loaded")
	}
	return _cGPDFArrayGetDictionary(array, index, value)
}

var _cGPDFArrayGetInteger func(array CGPDFArrayRef, index uintptr, value *CGPDFInteger) bool

// CGPDFArrayGetInteger returns whether an object at a given index in a PDF array is a PDF integer and, if so, retrieves that object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetInteger(_:_:_:)
func CGPDFArrayGetInteger(array CGPDFArrayRef, index uintptr, value *CGPDFInteger) bool {
	if _cGPDFArrayGetInteger == nil {
		panic("CoreGraphics: symbol CGPDFArrayGetInteger not loaded")
	}
	return _cGPDFArrayGetInteger(array, index, value)
}

var _cGPDFArrayGetName func(array CGPDFArrayRef, index uintptr, value string) bool

// CGPDFArrayGetName returns whether an object at a given index in a PDF array is a PDF name reference (represented as a constant C string) and, if so, retrieves that name.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetName(_:_:_:)
func CGPDFArrayGetName(array CGPDFArrayRef, index uintptr, value string) bool {
	if _cGPDFArrayGetName == nil {
		panic("CoreGraphics: symbol CGPDFArrayGetName not loaded")
	}
	return _cGPDFArrayGetName(array, index, value)
}

var _cGPDFArrayGetNull func(array CGPDFArrayRef, index uintptr) bool

// CGPDFArrayGetNull returns whether an object at a given index in a Quartz PDF array is a PDF null.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetNull(_:_:)
func CGPDFArrayGetNull(array CGPDFArrayRef, index uintptr) bool {
	if _cGPDFArrayGetNull == nil {
		panic("CoreGraphics: symbol CGPDFArrayGetNull not loaded")
	}
	return _cGPDFArrayGetNull(array, index)
}

var _cGPDFArrayGetNumber func(array CGPDFArrayRef, index uintptr, value *CGPDFReal) bool

// CGPDFArrayGetNumber returns whether an object at a given index in a PDF array is a PDF number and, if so, retrieves that object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetNumber(_:_:_:)
func CGPDFArrayGetNumber(array CGPDFArrayRef, index uintptr, value *CGPDFReal) bool {
	if _cGPDFArrayGetNumber == nil {
		panic("CoreGraphics: symbol CGPDFArrayGetNumber not loaded")
	}
	return _cGPDFArrayGetNumber(array, index, value)
}

var _cGPDFArrayGetObject func(array CGPDFArrayRef, index uintptr, value *CGPDFObjectRef) bool

// CGPDFArrayGetObject returns whether an object at a given index in a PDF array is a PDF object and, if so, retrieves that object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetObject(_:_:_:)
func CGPDFArrayGetObject(array CGPDFArrayRef, index uintptr, value *CGPDFObjectRef) bool {
	if _cGPDFArrayGetObject == nil {
		panic("CoreGraphics: symbol CGPDFArrayGetObject not loaded")
	}
	return _cGPDFArrayGetObject(array, index, value)
}

var _cGPDFArrayGetStream func(array CGPDFArrayRef, index uintptr, value *CGPDFStreamRef) bool

// CGPDFArrayGetStream returns whether an object at a given index in a PDF array is a PDF stream and, if so, retrieves that stream.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetStream(_:_:_:)
func CGPDFArrayGetStream(array CGPDFArrayRef, index uintptr, value *CGPDFStreamRef) bool {
	if _cGPDFArrayGetStream == nil {
		panic("CoreGraphics: symbol CGPDFArrayGetStream not loaded")
	}
	return _cGPDFArrayGetStream(array, index, value)
}

var _cGPDFArrayGetString func(array CGPDFArrayRef, index uintptr, value *CGPDFStringRef) bool

// CGPDFArrayGetString returns whether an object at a given index in a PDF array is a PDF string and, if so, retrieves that string.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetString(_:_:_:)
func CGPDFArrayGetString(array CGPDFArrayRef, index uintptr, value *CGPDFStringRef) bool {
	if _cGPDFArrayGetString == nil {
		panic("CoreGraphics: symbol CGPDFArrayGetString not loaded")
	}
	return _cGPDFArrayGetString(array, index, value)
}

var _cGPDFContentStreamCreateWithPage func(page CGPDFPageRef) CGPDFContentStreamRef

// CGPDFContentStreamCreateWithPage creates a content stream object from a PDF page object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContentStreamCreateWithPage(_:)
func CGPDFContentStreamCreateWithPage(page CGPDFPageRef) CGPDFContentStreamRef {
	if _cGPDFContentStreamCreateWithPage == nil {
		panic("CoreGraphics: symbol CGPDFContentStreamCreateWithPage not loaded")
	}
	return _cGPDFContentStreamCreateWithPage(page)
}

var _cGPDFContentStreamCreateWithStream func(stream CGPDFStreamRef, streamResources CGPDFDictionaryRef, parent CGPDFContentStreamRef) CGPDFContentStreamRef

// CGPDFContentStreamCreateWithStream creates a PDF content stream object from an existing PDF content stream object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContentStreamCreateWithStream(_:_:_:)
func CGPDFContentStreamCreateWithStream(stream CGPDFStreamRef, streamResources CGPDFDictionaryRef, parent CGPDFContentStreamRef) CGPDFContentStreamRef {
	if _cGPDFContentStreamCreateWithStream == nil {
		panic("CoreGraphics: symbol CGPDFContentStreamCreateWithStream not loaded")
	}
	return _cGPDFContentStreamCreateWithStream(stream, streamResources, parent)
}

var _cGPDFContentStreamGetResource func(cs CGPDFContentStreamRef, category string, name string) CGPDFObjectRef

// CGPDFContentStreamGetResource gets the specified resource from a PDF content stream object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContentStreamGetResource(_:_:_:)
func CGPDFContentStreamGetResource(cs CGPDFContentStreamRef, category string, name string) CGPDFObjectRef {
	if _cGPDFContentStreamGetResource == nil {
		panic("CoreGraphics: symbol CGPDFContentStreamGetResource not loaded")
	}
	return _cGPDFContentStreamGetResource(cs, category, name)
}

var _cGPDFContentStreamGetStreams func(cs CGPDFContentStreamRef) corefoundation.CFArrayRef

// CGPDFContentStreamGetStreams gets the array of PDF content streams contained in a PDF content stream object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContentStreamGetStreams(_:)
func CGPDFContentStreamGetStreams(cs CGPDFContentStreamRef) corefoundation.CFArrayRef {
	if _cGPDFContentStreamGetStreams == nil {
		panic("CoreGraphics: symbol CGPDFContentStreamGetStreams not loaded")
	}
	return _cGPDFContentStreamGetStreams(cs)
}

var _cGPDFContentStreamRelease func(cs CGPDFContentStreamRef)

// CGPDFContentStreamRelease decrements the retain count of a PDF content stream object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContentStreamRelease(_:)
func CGPDFContentStreamRelease(cs CGPDFContentStreamRef) {
	if _cGPDFContentStreamRelease == nil {
		panic("CoreGraphics: symbol CGPDFContentStreamRelease not loaded")
	}
	_cGPDFContentStreamRelease(cs)
}

var _cGPDFContentStreamRetain func(cs CGPDFContentStreamRef) CGPDFContentStreamRef

// CGPDFContentStreamRetain increments the retain count of a PDF content stream object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContentStreamRetain(_:)
func CGPDFContentStreamRetain(cs CGPDFContentStreamRef) CGPDFContentStreamRef {
	if _cGPDFContentStreamRetain == nil {
		panic("CoreGraphics: symbol CGPDFContentStreamRetain not loaded")
	}
	return _cGPDFContentStreamRetain(cs)
}

var _cGPDFContextAddDestinationAtPoint func(context CGContextRef, name corefoundation.CFStringRef, point corefoundation.CGPoint)

// CGPDFContextAddDestinationAtPoint sets a destination to jump to when a point in the current page of a PDF graphics context is clicked.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/addDestination(_:at:)
func CGPDFContextAddDestinationAtPoint(context CGContextRef, name corefoundation.CFStringRef, point corefoundation.CGPoint) {
	if _cGPDFContextAddDestinationAtPoint == nil {
		panic("CoreGraphics: symbol CGPDFContextAddDestinationAtPoint not loaded")
	}
	_cGPDFContextAddDestinationAtPoint(context, name, point)
}

var _cGPDFContextAddDocumentMetadata func(context CGContextRef, metadata corefoundation.CFDataRef)

// CGPDFContextAddDocumentMetadata associates custom metadata with the PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/addDocumentMetadata(_:)
func CGPDFContextAddDocumentMetadata(context CGContextRef, metadata corefoundation.CFDataRef) {
	if _cGPDFContextAddDocumentMetadata == nil {
		panic("CoreGraphics: symbol CGPDFContextAddDocumentMetadata not loaded")
	}
	_cGPDFContextAddDocumentMetadata(context, metadata)
}

var _cGPDFContextBeginPage func(context CGContextRef, pageInfo corefoundation.CFDictionaryRef)

// CGPDFContextBeginPage begins a new page in a PDF graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/beginPDFPage(_:)
func CGPDFContextBeginPage(context CGContextRef, pageInfo corefoundation.CFDictionaryRef) {
	if _cGPDFContextBeginPage == nil {
		panic("CoreGraphics: symbol CGPDFContextBeginPage not loaded")
	}
	_cGPDFContextBeginPage(context, pageInfo)
}

var _cGPDFContextBeginTag func(context CGContextRef, tagType CGPDFTagType, tagProperties corefoundation.CFDictionaryRef)

// CGPDFContextBeginTag.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContextBeginTag(_:_:_:)
func CGPDFContextBeginTag(context CGContextRef, tagType CGPDFTagType, tagProperties corefoundation.CFDictionaryRef) {
	if _cGPDFContextBeginTag == nil {
		panic("CoreGraphics: symbol CGPDFContextBeginTag not loaded")
	}
	_cGPDFContextBeginTag(context, tagType, tagProperties)
}

var _cGPDFContextClose func(context CGContextRef)

// CGPDFContextClose closes a PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/closePDF()
func CGPDFContextClose(context CGContextRef) {
	if _cGPDFContextClose == nil {
		panic("CoreGraphics: symbol CGPDFContextClose not loaded")
	}
	_cGPDFContextClose(context)
}

var _cGPDFContextCreate func(consumer CGDataConsumerRef, mediaBox *corefoundation.CGRect, auxiliaryInfo corefoundation.CFDictionaryRef) CGContextRef

// CGPDFContextCreate creates a PDF graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/init(consumer:mediaBox:_:)
func CGPDFContextCreate(consumer CGDataConsumerRef, mediaBox *corefoundation.CGRect, auxiliaryInfo corefoundation.CFDictionaryRef) CGContextRef {
	if _cGPDFContextCreate == nil {
		panic("CoreGraphics: symbol CGPDFContextCreate not loaded")
	}
	return _cGPDFContextCreate(consumer, mediaBox, auxiliaryInfo)
}

var _cGPDFContextCreateWithURL func(url corefoundation.CFURLRef, mediaBox *corefoundation.CGRect, auxiliaryInfo corefoundation.CFDictionaryRef) CGContextRef

// CGPDFContextCreateWithURL creates a URL-based PDF graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/init(_:mediaBox:_:)
func CGPDFContextCreateWithURL(url corefoundation.CFURLRef, mediaBox *corefoundation.CGRect, auxiliaryInfo corefoundation.CFDictionaryRef) CGContextRef {
	if _cGPDFContextCreateWithURL == nil {
		panic("CoreGraphics: symbol CGPDFContextCreateWithURL not loaded")
	}
	return _cGPDFContextCreateWithURL(url, mediaBox, auxiliaryInfo)
}

var _cGPDFContextEndPage func(context CGContextRef)

// CGPDFContextEndPage ends the current page in the PDF graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/endPDFPage()
func CGPDFContextEndPage(context CGContextRef) {
	if _cGPDFContextEndPage == nil {
		panic("CoreGraphics: symbol CGPDFContextEndPage not loaded")
	}
	_cGPDFContextEndPage(context)
}

var _cGPDFContextEndTag func(context CGContextRef)

// CGPDFContextEndTag.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContextEndTag(_:)
func CGPDFContextEndTag(context CGContextRef) {
	if _cGPDFContextEndTag == nil {
		panic("CoreGraphics: symbol CGPDFContextEndTag not loaded")
	}
	_cGPDFContextEndTag(context)
}

var _cGPDFContextSetDestinationForRect func(context CGContextRef, name corefoundation.CFStringRef, rect corefoundation.CGRect)

// CGPDFContextSetDestinationForRect sets a destination to jump to when a rectangle in the current PDF page is clicked.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setDestination(_:for:)
func CGPDFContextSetDestinationForRect(context CGContextRef, name corefoundation.CFStringRef, rect corefoundation.CGRect) {
	if _cGPDFContextSetDestinationForRect == nil {
		panic("CoreGraphics: symbol CGPDFContextSetDestinationForRect not loaded")
	}
	_cGPDFContextSetDestinationForRect(context, name, rect)
}

var _cGPDFContextSetIDTree func(context CGContextRef, IDTreeDictionary CGPDFDictionaryRef)

// CGPDFContextSetIDTree.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContextSetIDTree(_:_:)
func CGPDFContextSetIDTree(context CGContextRef, IDTreeDictionary CGPDFDictionaryRef) {
	if _cGPDFContextSetIDTree == nil {
		panic("CoreGraphics: symbol CGPDFContextSetIDTree not loaded")
	}
	_cGPDFContextSetIDTree(context, IDTreeDictionary)
}

var _cGPDFContextSetOutline func(context CGContextRef, outline corefoundation.CFDictionaryRef)

// CGPDFContextSetOutline.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContextSetOutline(_:_:)
func CGPDFContextSetOutline(context CGContextRef, outline corefoundation.CFDictionaryRef) {
	if _cGPDFContextSetOutline == nil {
		panic("CoreGraphics: symbol CGPDFContextSetOutline not loaded")
	}
	_cGPDFContextSetOutline(context, outline)
}

var _cGPDFContextSetPageTagStructureTree func(context CGContextRef, pageTagStructureTreeDictionary corefoundation.CFDictionaryRef)

// CGPDFContextSetPageTagStructureTree.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContextSetPageTagStructureTree(_:_:)
func CGPDFContextSetPageTagStructureTree(context CGContextRef, pageTagStructureTreeDictionary corefoundation.CFDictionaryRef) {
	if _cGPDFContextSetPageTagStructureTree == nil {
		panic("CoreGraphics: symbol CGPDFContextSetPageTagStructureTree not loaded")
	}
	_cGPDFContextSetPageTagStructureTree(context, pageTagStructureTreeDictionary)
}

var _cGPDFContextSetParentTree func(context CGContextRef, parentTreeDictionary CGPDFDictionaryRef)

// CGPDFContextSetParentTree.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContextSetParentTree(_:_:)
func CGPDFContextSetParentTree(context CGContextRef, parentTreeDictionary CGPDFDictionaryRef) {
	if _cGPDFContextSetParentTree == nil {
		panic("CoreGraphics: symbol CGPDFContextSetParentTree not loaded")
	}
	_cGPDFContextSetParentTree(context, parentTreeDictionary)
}

var _cGPDFContextSetURLForRect func(context CGContextRef, url corefoundation.CFURLRef, rect corefoundation.CGRect)

// CGPDFContextSetURLForRect sets the URL associated with a rectangle in a PDF graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setURL(_:for:)
func CGPDFContextSetURLForRect(context CGContextRef, url corefoundation.CFURLRef, rect corefoundation.CGRect) {
	if _cGPDFContextSetURLForRect == nil {
		panic("CoreGraphics: symbol CGPDFContextSetURLForRect not loaded")
	}
	_cGPDFContextSetURLForRect(context, url, rect)
}

var _cGPDFDictionaryApplyBlock func(dict CGPDFDictionaryRef, block CGPDFDictionaryApplierBlock, info unsafe.Pointer)

// CGPDFDictionaryApplyBlock.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryApplyBlock(_:_:_:)
func CGPDFDictionaryApplyBlock(dict CGPDFDictionaryRef, block CGPDFDictionaryApplierBlock, info unsafe.Pointer) {
	if _cGPDFDictionaryApplyBlock == nil {
		panic("CoreGraphics: symbol CGPDFDictionaryApplyBlock not loaded")
	}
	_cGPDFDictionaryApplyBlock(dict, block, info)
}

var _cGPDFDictionaryApplyFunction func(dict CGPDFDictionaryRef, function CGPDFDictionaryApplierFunction, info unsafe.Pointer)

// CGPDFDictionaryApplyFunction applies a function to each entry in a dictionary.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryApplyFunction(_:_:_:)
func CGPDFDictionaryApplyFunction(dict CGPDFDictionaryRef, function CGPDFDictionaryApplierFunction, info unsafe.Pointer) {
	if _cGPDFDictionaryApplyFunction == nil {
		panic("CoreGraphics: symbol CGPDFDictionaryApplyFunction not loaded")
	}
	_cGPDFDictionaryApplyFunction(dict, function, info)
}

var _cGPDFDictionaryGetArray func(dict CGPDFDictionaryRef, key string, value *CGPDFArrayRef) bool

// CGPDFDictionaryGetArray returns whether there is a PDF array associated with a specified key in a PDF dictionary and, if so, retrieves that array.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetArray(_:_:_:)
func CGPDFDictionaryGetArray(dict CGPDFDictionaryRef, key string, value *CGPDFArrayRef) bool {
	if _cGPDFDictionaryGetArray == nil {
		panic("CoreGraphics: symbol CGPDFDictionaryGetArray not loaded")
	}
	return _cGPDFDictionaryGetArray(dict, key, value)
}

var _cGPDFDictionaryGetBoolean func(dict CGPDFDictionaryRef, key string, value *CGPDFBoolean) bool

// CGPDFDictionaryGetBoolean returns whether there is a PDF Boolean value associated with a specified key in a PDF dictionary and, if so, retrieves the Boolean value.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetBoolean(_:_:_:)
func CGPDFDictionaryGetBoolean(dict CGPDFDictionaryRef, key string, value *CGPDFBoolean) bool {
	if _cGPDFDictionaryGetBoolean == nil {
		panic("CoreGraphics: symbol CGPDFDictionaryGetBoolean not loaded")
	}
	return _cGPDFDictionaryGetBoolean(dict, key, value)
}

var _cGPDFDictionaryGetCount func(dict CGPDFDictionaryRef) uintptr

// CGPDFDictionaryGetCount returns the number of entries in a PDF dictionary.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetCount(_:)
func CGPDFDictionaryGetCount(dict CGPDFDictionaryRef) uintptr {
	if _cGPDFDictionaryGetCount == nil {
		panic("CoreGraphics: symbol CGPDFDictionaryGetCount not loaded")
	}
	return _cGPDFDictionaryGetCount(dict)
}

var _cGPDFDictionaryGetDictionary func(dict CGPDFDictionaryRef, key string, value *CGPDFDictionaryRef) bool

// CGPDFDictionaryGetDictionary returns whether there is another PDF dictionary associated with a specified key in a PDF dictionary and, if so, retrieves that dictionary.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetDictionary(_:_:_:)
func CGPDFDictionaryGetDictionary(dict CGPDFDictionaryRef, key string, value *CGPDFDictionaryRef) bool {
	if _cGPDFDictionaryGetDictionary == nil {
		panic("CoreGraphics: symbol CGPDFDictionaryGetDictionary not loaded")
	}
	return _cGPDFDictionaryGetDictionary(dict, key, value)
}

var _cGPDFDictionaryGetInteger func(dict CGPDFDictionaryRef, key string, value *CGPDFInteger) bool

// CGPDFDictionaryGetInteger returns whether there is a PDF integer associated with a specified key in a PDF dictionary and, if so, retrieves that integer.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetInteger(_:_:_:)
func CGPDFDictionaryGetInteger(dict CGPDFDictionaryRef, key string, value *CGPDFInteger) bool {
	if _cGPDFDictionaryGetInteger == nil {
		panic("CoreGraphics: symbol CGPDFDictionaryGetInteger not loaded")
	}
	return _cGPDFDictionaryGetInteger(dict, key, value)
}

var _cGPDFDictionaryGetName func(dict CGPDFDictionaryRef, key string, value string) bool

// CGPDFDictionaryGetName returns whether an object with a specified key in a PDF dictionary is a PDF name reference (represented as a constant C string) and, if so, retrieves that name.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetName(_:_:_:)
func CGPDFDictionaryGetName(dict CGPDFDictionaryRef, key string, value string) bool {
	if _cGPDFDictionaryGetName == nil {
		panic("CoreGraphics: symbol CGPDFDictionaryGetName not loaded")
	}
	return _cGPDFDictionaryGetName(dict, key, value)
}

var _cGPDFDictionaryGetNumber func(dict CGPDFDictionaryRef, key string, value *CGPDFReal) bool

// CGPDFDictionaryGetNumber returns whether there is a PDF number associated with a specified key in a PDF dictionary and, if so, retrieves that number.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetNumber(_:_:_:)
func CGPDFDictionaryGetNumber(dict CGPDFDictionaryRef, key string, value *CGPDFReal) bool {
	if _cGPDFDictionaryGetNumber == nil {
		panic("CoreGraphics: symbol CGPDFDictionaryGetNumber not loaded")
	}
	return _cGPDFDictionaryGetNumber(dict, key, value)
}

var _cGPDFDictionaryGetObject func(dict CGPDFDictionaryRef, key string, value *CGPDFObjectRef) bool

// CGPDFDictionaryGetObject returns whether there is a PDF object associated with a specified key in a PDF dictionary and, if so, retrieves that object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetObject(_:_:_:)
func CGPDFDictionaryGetObject(dict CGPDFDictionaryRef, key string, value *CGPDFObjectRef) bool {
	if _cGPDFDictionaryGetObject == nil {
		panic("CoreGraphics: symbol CGPDFDictionaryGetObject not loaded")
	}
	return _cGPDFDictionaryGetObject(dict, key, value)
}

var _cGPDFDictionaryGetStream func(dict CGPDFDictionaryRef, key string, value *CGPDFStreamRef) bool

// CGPDFDictionaryGetStream returns whether there is a PDF stream associated with a specified key in a PDF dictionary and, if so, retrieves that stream.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetStream(_:_:_:)
func CGPDFDictionaryGetStream(dict CGPDFDictionaryRef, key string, value *CGPDFStreamRef) bool {
	if _cGPDFDictionaryGetStream == nil {
		panic("CoreGraphics: symbol CGPDFDictionaryGetStream not loaded")
	}
	return _cGPDFDictionaryGetStream(dict, key, value)
}

var _cGPDFDictionaryGetString func(dict CGPDFDictionaryRef, key string, value *CGPDFStringRef) bool

// CGPDFDictionaryGetString returns whether there is a PDF string associated with a specified key in a PDF dictionary and, if so, retrieves that string.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetString(_:_:_:)
func CGPDFDictionaryGetString(dict CGPDFDictionaryRef, key string, value *CGPDFStringRef) bool {
	if _cGPDFDictionaryGetString == nil {
		panic("CoreGraphics: symbol CGPDFDictionaryGetString not loaded")
	}
	return _cGPDFDictionaryGetString(dict, key, value)
}

var _cGPDFDocumentAllowsCopying func(document CGPDFDocumentRef) bool

// CGPDFDocumentAllowsCopying returns whether the specified PDF document allows copying.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/allowsCopying
func CGPDFDocumentAllowsCopying(document CGPDFDocumentRef) bool {
	if _cGPDFDocumentAllowsCopying == nil {
		panic("CoreGraphics: symbol CGPDFDocumentAllowsCopying not loaded")
	}
	return _cGPDFDocumentAllowsCopying(document)
}

var _cGPDFDocumentAllowsPrinting func(document CGPDFDocumentRef) bool

// CGPDFDocumentAllowsPrinting returns whether a PDF document allows printing.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/allowsPrinting
func CGPDFDocumentAllowsPrinting(document CGPDFDocumentRef) bool {
	if _cGPDFDocumentAllowsPrinting == nil {
		panic("CoreGraphics: symbol CGPDFDocumentAllowsPrinting not loaded")
	}
	return _cGPDFDocumentAllowsPrinting(document)
}

var _cGPDFDocumentCreateWithProvider func(provider CGDataProviderRef) CGPDFDocumentRef

// CGPDFDocumentCreateWithProvider creates a Core Graphics PDF document using a data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/init(_:)-gbq6
func CGPDFDocumentCreateWithProvider(provider CGDataProviderRef) CGPDFDocumentRef {
	if _cGPDFDocumentCreateWithProvider == nil {
		panic("CoreGraphics: symbol CGPDFDocumentCreateWithProvider not loaded")
	}
	return _cGPDFDocumentCreateWithProvider(provider)
}

var _cGPDFDocumentCreateWithURL func(url corefoundation.CFURLRef) CGPDFDocumentRef

// CGPDFDocumentCreateWithURL creates a Core Graphics PDF document using data specified by a URL.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/init(_:)-2gtsd
func CGPDFDocumentCreateWithURL(url corefoundation.CFURLRef) CGPDFDocumentRef {
	if _cGPDFDocumentCreateWithURL == nil {
		panic("CoreGraphics: symbol CGPDFDocumentCreateWithURL not loaded")
	}
	return _cGPDFDocumentCreateWithURL(url)
}

var _cGPDFDocumentGetAccessPermissions func(document CGPDFDocumentRef) CGPDFAccessPermissions

// CGPDFDocumentGetAccessPermissions.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/accessPermissions
func CGPDFDocumentGetAccessPermissions(document CGPDFDocumentRef) CGPDFAccessPermissions {
	if _cGPDFDocumentGetAccessPermissions == nil {
		panic("CoreGraphics: symbol CGPDFDocumentGetAccessPermissions not loaded")
	}
	return _cGPDFDocumentGetAccessPermissions(document)
}

var _cGPDFDocumentGetArtBox func(document CGPDFDocumentRef, page int) corefoundation.CGRect

// CGPDFDocumentGetArtBox returns the art box of a page in a PDF document.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentGetArtBox
func CGPDFDocumentGetArtBox(document CGPDFDocumentRef, page int) corefoundation.CGRect {
	if _cGPDFDocumentGetArtBox == nil {
		panic("CoreGraphics: symbol CGPDFDocumentGetArtBox not loaded")
	}
	return _cGPDFDocumentGetArtBox(document, page)
}

var _cGPDFDocumentGetBleedBox func(document CGPDFDocumentRef, page int) corefoundation.CGRect

// CGPDFDocumentGetBleedBox returns the bleed box of a page in a PDF document.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentGetBleedBox
func CGPDFDocumentGetBleedBox(document CGPDFDocumentRef, page int) corefoundation.CGRect {
	if _cGPDFDocumentGetBleedBox == nil {
		panic("CoreGraphics: symbol CGPDFDocumentGetBleedBox not loaded")
	}
	return _cGPDFDocumentGetBleedBox(document, page)
}

var _cGPDFDocumentGetCatalog func(document CGPDFDocumentRef) CGPDFDictionaryRef

// CGPDFDocumentGetCatalog returns the document catalog of a Core Graphics PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/catalog
func CGPDFDocumentGetCatalog(document CGPDFDocumentRef) CGPDFDictionaryRef {
	if _cGPDFDocumentGetCatalog == nil {
		panic("CoreGraphics: symbol CGPDFDocumentGetCatalog not loaded")
	}
	return _cGPDFDocumentGetCatalog(document)
}

var _cGPDFDocumentGetCropBox func(document CGPDFDocumentRef, page int) corefoundation.CGRect

// CGPDFDocumentGetCropBox returns the crop box of a page in a PDF document.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentGetCropBox
func CGPDFDocumentGetCropBox(document CGPDFDocumentRef, page int) corefoundation.CGRect {
	if _cGPDFDocumentGetCropBox == nil {
		panic("CoreGraphics: symbol CGPDFDocumentGetCropBox not loaded")
	}
	return _cGPDFDocumentGetCropBox(document, page)
}

var _cGPDFDocumentGetID func(document CGPDFDocumentRef) CGPDFArrayRef

// CGPDFDocumentGetID gets the file identifier for a PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/fileIdentifier
func CGPDFDocumentGetID(document CGPDFDocumentRef) CGPDFArrayRef {
	if _cGPDFDocumentGetID == nil {
		panic("CoreGraphics: symbol CGPDFDocumentGetID not loaded")
	}
	return _cGPDFDocumentGetID(document)
}

var _cGPDFDocumentGetInfo func(document CGPDFDocumentRef) CGPDFDictionaryRef

// CGPDFDocumentGetInfo gets the information dictionary for a PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/info
func CGPDFDocumentGetInfo(document CGPDFDocumentRef) CGPDFDictionaryRef {
	if _cGPDFDocumentGetInfo == nil {
		panic("CoreGraphics: symbol CGPDFDocumentGetInfo not loaded")
	}
	return _cGPDFDocumentGetInfo(document)
}

var _cGPDFDocumentGetMediaBox func(document CGPDFDocumentRef, page int) corefoundation.CGRect

// CGPDFDocumentGetMediaBox returns the media box of a page in a PDF document.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentGetMediaBox
func CGPDFDocumentGetMediaBox(document CGPDFDocumentRef, page int) corefoundation.CGRect {
	if _cGPDFDocumentGetMediaBox == nil {
		panic("CoreGraphics: symbol CGPDFDocumentGetMediaBox not loaded")
	}
	return _cGPDFDocumentGetMediaBox(document, page)
}

var _cGPDFDocumentGetNumberOfPages func(document CGPDFDocumentRef) uintptr

// CGPDFDocumentGetNumberOfPages returns the number of pages in a PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/numberOfPages
func CGPDFDocumentGetNumberOfPages(document CGPDFDocumentRef) uintptr {
	if _cGPDFDocumentGetNumberOfPages == nil {
		panic("CoreGraphics: symbol CGPDFDocumentGetNumberOfPages not loaded")
	}
	return _cGPDFDocumentGetNumberOfPages(document)
}

var _cGPDFDocumentGetOutline func(document CGPDFDocumentRef) corefoundation.CFDictionaryRef

// CGPDFDocumentGetOutline.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/outline
func CGPDFDocumentGetOutline(document CGPDFDocumentRef) corefoundation.CFDictionaryRef {
	if _cGPDFDocumentGetOutline == nil {
		panic("CoreGraphics: symbol CGPDFDocumentGetOutline not loaded")
	}
	return _cGPDFDocumentGetOutline(document)
}

var _cGPDFDocumentGetPage func(document CGPDFDocumentRef, pageNumber uintptr) CGPDFPageRef

// CGPDFDocumentGetPage returns a page from a Core Graphics PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/page(at:)
func CGPDFDocumentGetPage(document CGPDFDocumentRef, pageNumber uintptr) CGPDFPageRef {
	if _cGPDFDocumentGetPage == nil {
		panic("CoreGraphics: symbol CGPDFDocumentGetPage not loaded")
	}
	return _cGPDFDocumentGetPage(document, pageNumber)
}

var _cGPDFDocumentGetRotationAngle func(document CGPDFDocumentRef, page int) int

// CGPDFDocumentGetRotationAngle returns the rotation angle of a page in a PDF document.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentGetRotationAngle
func CGPDFDocumentGetRotationAngle(document CGPDFDocumentRef, page int) int {
	if _cGPDFDocumentGetRotationAngle == nil {
		panic("CoreGraphics: symbol CGPDFDocumentGetRotationAngle not loaded")
	}
	return _cGPDFDocumentGetRotationAngle(document, page)
}

var _cGPDFDocumentGetTrimBox func(document CGPDFDocumentRef, page int) corefoundation.CGRect

// CGPDFDocumentGetTrimBox returns the trim box of a page in a PDF document.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentGetTrimBox
func CGPDFDocumentGetTrimBox(document CGPDFDocumentRef, page int) corefoundation.CGRect {
	if _cGPDFDocumentGetTrimBox == nil {
		panic("CoreGraphics: symbol CGPDFDocumentGetTrimBox not loaded")
	}
	return _cGPDFDocumentGetTrimBox(document, page)
}

var _cGPDFDocumentGetTypeID func() uint

// CGPDFDocumentGetTypeID returns the type identifier for Core Graphics PDF documents.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/typeID
func CGPDFDocumentGetTypeID() uint {
	if _cGPDFDocumentGetTypeID == nil {
		panic("CoreGraphics: symbol CGPDFDocumentGetTypeID not loaded")
	}
	return _cGPDFDocumentGetTypeID()
}

var _cGPDFDocumentGetVersion func(document CGPDFDocumentRef, majorVersion *int, minorVersion *int)

// CGPDFDocumentGetVersion returns the major and minor version numbers of a Core Graphics PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/getVersion(majorVersion:minorVersion:)
func CGPDFDocumentGetVersion(document CGPDFDocumentRef, majorVersion []int, minorVersion []int) {
	if _cGPDFDocumentGetVersion == nil {
		panic("CoreGraphics: symbol CGPDFDocumentGetVersion not loaded")
	}
	_cGPDFDocumentGetVersion(document, unsafe.SliceData(majorVersion), unsafe.SliceData(minorVersion))
}

var _cGPDFDocumentIsEncrypted func(document CGPDFDocumentRef) bool

// CGPDFDocumentIsEncrypted returns whether the specified PDF file is encrypted.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/isEncrypted
func CGPDFDocumentIsEncrypted(document CGPDFDocumentRef) bool {
	if _cGPDFDocumentIsEncrypted == nil {
		panic("CoreGraphics: symbol CGPDFDocumentIsEncrypted not loaded")
	}
	return _cGPDFDocumentIsEncrypted(document)
}

var _cGPDFDocumentIsUnlocked func(document CGPDFDocumentRef) bool

// CGPDFDocumentIsUnlocked returns whether the specified PDF document is currently unlocked.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/isUnlocked
func CGPDFDocumentIsUnlocked(document CGPDFDocumentRef) bool {
	if _cGPDFDocumentIsUnlocked == nil {
		panic("CoreGraphics: symbol CGPDFDocumentIsUnlocked not loaded")
	}
	return _cGPDFDocumentIsUnlocked(document)
}

var _cGPDFDocumentRelease func(document CGPDFDocumentRef)

// CGPDFDocumentRelease decrements the retain count of a PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentRelease
func CGPDFDocumentRelease(document CGPDFDocumentRef) {
	if _cGPDFDocumentRelease == nil {
		panic("CoreGraphics: symbol CGPDFDocumentRelease not loaded")
	}
	_cGPDFDocumentRelease(document)
}

var _cGPDFDocumentRetain func(document CGPDFDocumentRef) CGPDFDocumentRef

// CGPDFDocumentRetain increments the retain count of a Core Graphics PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentRetain
func CGPDFDocumentRetain(document CGPDFDocumentRef) CGPDFDocumentRef {
	if _cGPDFDocumentRetain == nil {
		panic("CoreGraphics: symbol CGPDFDocumentRetain not loaded")
	}
	return _cGPDFDocumentRetain(document)
}

var _cGPDFDocumentUnlockWithPassword func(document CGPDFDocumentRef, password string) bool

// CGPDFDocumentUnlockWithPassword unlocks an encrypted PDF document when a valid password is supplied.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/unlockWithPassword(_:)
func CGPDFDocumentUnlockWithPassword(document CGPDFDocumentRef, password string) bool {
	if _cGPDFDocumentUnlockWithPassword == nil {
		panic("CoreGraphics: symbol CGPDFDocumentUnlockWithPassword not loaded")
	}
	return _cGPDFDocumentUnlockWithPassword(document, password)
}

var _cGPDFObjectGetType func(object CGPDFObjectRef) CGPDFObjectType

// CGPDFObjectGetType returns the PDF type identifier of an object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectGetType(_:)
func CGPDFObjectGetType(object CGPDFObjectRef) CGPDFObjectType {
	if _cGPDFObjectGetType == nil {
		panic("CoreGraphics: symbol CGPDFObjectGetType not loaded")
	}
	return _cGPDFObjectGetType(object)
}

var _cGPDFObjectGetValue func(object CGPDFObjectRef, type_ CGPDFObjectType, value unsafe.Pointer) bool

// CGPDFObjectGetValue returns whether an object is of a given type and if it is, retrieves its value.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectGetValue(_:_:_:)
func CGPDFObjectGetValue(object CGPDFObjectRef, type_ CGPDFObjectType, value unsafe.Pointer) bool {
	if _cGPDFObjectGetValue == nil {
		panic("CoreGraphics: symbol CGPDFObjectGetValue not loaded")
	}
	return _cGPDFObjectGetValue(object, type_, value)
}

var _cGPDFOperatorTableCreate func() CGPDFOperatorTableRef

// CGPDFOperatorTableCreate creates an empty PDF operator table.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFOperatorTableCreate()
func CGPDFOperatorTableCreate() CGPDFOperatorTableRef {
	if _cGPDFOperatorTableCreate == nil {
		panic("CoreGraphics: symbol CGPDFOperatorTableCreate not loaded")
	}
	return _cGPDFOperatorTableCreate()
}

var _cGPDFOperatorTableRelease func(table CGPDFOperatorTableRef)

// CGPDFOperatorTableRelease decrements the retain count of a CGPDFOperatorTable object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFOperatorTableRelease(_:)
func CGPDFOperatorTableRelease(table CGPDFOperatorTableRef) {
	if _cGPDFOperatorTableRelease == nil {
		panic("CoreGraphics: symbol CGPDFOperatorTableRelease not loaded")
	}
	_cGPDFOperatorTableRelease(table)
}

var _cGPDFOperatorTableRetain func(table CGPDFOperatorTableRef) CGPDFOperatorTableRef

// CGPDFOperatorTableRetain increments the retain count of a CGPDFOperatorTable object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFOperatorTableRetain(_:)
func CGPDFOperatorTableRetain(table CGPDFOperatorTableRef) CGPDFOperatorTableRef {
	if _cGPDFOperatorTableRetain == nil {
		panic("CoreGraphics: symbol CGPDFOperatorTableRetain not loaded")
	}
	return _cGPDFOperatorTableRetain(table)
}

var _cGPDFOperatorTableSetCallback func(table CGPDFOperatorTableRef, name string, callback CGPDFOperatorCallback)

// CGPDFOperatorTableSetCallback sets a callback function for a PDF operator.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFOperatorTableSetCallback(_:_:_:)
func CGPDFOperatorTableSetCallback(table CGPDFOperatorTableRef, name string, callback CGPDFOperatorCallback) {
	if _cGPDFOperatorTableSetCallback == nil {
		panic("CoreGraphics: symbol CGPDFOperatorTableSetCallback not loaded")
	}
	_cGPDFOperatorTableSetCallback(table, name, callback)
}

var _cGPDFPageGetBoxRect func(page CGPDFPageRef, box CGPDFBox) corefoundation.CGRect

// CGPDFPageGetBoxRect returns the rectangle that represents a type of box for a content region or page dimensions of a PDF page.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/getBoxRect(_:)
func CGPDFPageGetBoxRect(page CGPDFPageRef, box CGPDFBox) corefoundation.CGRect {
	if _cGPDFPageGetBoxRect == nil {
		panic("CoreGraphics: symbol CGPDFPageGetBoxRect not loaded")
	}
	return _cGPDFPageGetBoxRect(page, box)
}

var _cGPDFPageGetDictionary func(page CGPDFPageRef) CGPDFDictionaryRef

// CGPDFPageGetDictionary returns the dictionary of a PDF page.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/dictionary
func CGPDFPageGetDictionary(page CGPDFPageRef) CGPDFDictionaryRef {
	if _cGPDFPageGetDictionary == nil {
		panic("CoreGraphics: symbol CGPDFPageGetDictionary not loaded")
	}
	return _cGPDFPageGetDictionary(page)
}

var _cGPDFPageGetDocument func(page CGPDFPageRef) CGPDFDocumentRef

// CGPDFPageGetDocument returns the document for a page.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/document
func CGPDFPageGetDocument(page CGPDFPageRef) CGPDFDocumentRef {
	if _cGPDFPageGetDocument == nil {
		panic("CoreGraphics: symbol CGPDFPageGetDocument not loaded")
	}
	return _cGPDFPageGetDocument(page)
}

var _cGPDFPageGetDrawingTransform func(page CGPDFPageRef, box CGPDFBox, rect corefoundation.CGRect, rotate int, preserveAspectRatio bool) corefoundation.CGAffineTransform

// CGPDFPageGetDrawingTransform returns the affine transform that maps a box to a given rectangle on a PDF page.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/getDrawingTransform(_:rect:rotate:preserveAspectRatio:)
func CGPDFPageGetDrawingTransform(page CGPDFPageRef, box CGPDFBox, rect corefoundation.CGRect, rotate int, preserveAspectRatio bool) corefoundation.CGAffineTransform {
	if _cGPDFPageGetDrawingTransform == nil {
		panic("CoreGraphics: symbol CGPDFPageGetDrawingTransform not loaded")
	}
	return _cGPDFPageGetDrawingTransform(page, box, rect, rotate, preserveAspectRatio)
}

var _cGPDFPageGetPageNumber func(page CGPDFPageRef) uintptr

// CGPDFPageGetPageNumber returns the page number of the specified PDF page.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/pageNumber
func CGPDFPageGetPageNumber(page CGPDFPageRef) uintptr {
	if _cGPDFPageGetPageNumber == nil {
		panic("CoreGraphics: symbol CGPDFPageGetPageNumber not loaded")
	}
	return _cGPDFPageGetPageNumber(page)
}

var _cGPDFPageGetRotationAngle func(page CGPDFPageRef) int

// CGPDFPageGetRotationAngle returns the rotation angle of a PDF page, in degrees.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/rotationAngle
func CGPDFPageGetRotationAngle(page CGPDFPageRef) int {
	if _cGPDFPageGetRotationAngle == nil {
		panic("CoreGraphics: symbol CGPDFPageGetRotationAngle not loaded")
	}
	return _cGPDFPageGetRotationAngle(page)
}

var _cGPDFPageGetTypeID func() uint

// CGPDFPageGetTypeID returns the CFType ID for PDF page objects.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/typeID
func CGPDFPageGetTypeID() uint {
	if _cGPDFPageGetTypeID == nil {
		panic("CoreGraphics: symbol CGPDFPageGetTypeID not loaded")
	}
	return _cGPDFPageGetTypeID()
}

var _cGPDFPageRelease func(page CGPDFPageRef)

// CGPDFPageRelease decrements the retain count of a PDF page.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFPageRelease
func CGPDFPageRelease(page CGPDFPageRef) {
	if _cGPDFPageRelease == nil {
		panic("CoreGraphics: symbol CGPDFPageRelease not loaded")
	}
	_cGPDFPageRelease(page)
}

var _cGPDFPageRetain func(page CGPDFPageRef) CGPDFPageRef

// CGPDFPageRetain increments the retain count of a PDF page.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFPageRetain
func CGPDFPageRetain(page CGPDFPageRef) CGPDFPageRef {
	if _cGPDFPageRetain == nil {
		panic("CoreGraphics: symbol CGPDFPageRetain not loaded")
	}
	return _cGPDFPageRetain(page)
}

var _cGPDFScannerCreate func(cs CGPDFContentStreamRef, table CGPDFOperatorTableRef, info unsafe.Pointer) CGPDFScannerRef

// CGPDFScannerCreate creates a PDF scanner.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerCreate(_:_:_:)
func CGPDFScannerCreate(cs CGPDFContentStreamRef, table CGPDFOperatorTableRef, info unsafe.Pointer) CGPDFScannerRef {
	if _cGPDFScannerCreate == nil {
		panic("CoreGraphics: symbol CGPDFScannerCreate not loaded")
	}
	return _cGPDFScannerCreate(cs, table, info)
}

var _cGPDFScannerGetContentStream func(scanner CGPDFScannerRef) CGPDFContentStreamRef

// CGPDFScannerGetContentStream returns the content stream associated with a PDF scanner object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerGetContentStream(_:)
func CGPDFScannerGetContentStream(scanner CGPDFScannerRef) CGPDFContentStreamRef {
	if _cGPDFScannerGetContentStream == nil {
		panic("CoreGraphics: symbol CGPDFScannerGetContentStream not loaded")
	}
	return _cGPDFScannerGetContentStream(scanner)
}

var _cGPDFScannerPopArray func(scanner CGPDFScannerRef, value *CGPDFArrayRef) bool

// CGPDFScannerPopArray retrieves an array object from the scanner stack.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopArray(_:_:)
func CGPDFScannerPopArray(scanner CGPDFScannerRef, value *CGPDFArrayRef) bool {
	if _cGPDFScannerPopArray == nil {
		panic("CoreGraphics: symbol CGPDFScannerPopArray not loaded")
	}
	return _cGPDFScannerPopArray(scanner, value)
}

var _cGPDFScannerPopBoolean func(scanner CGPDFScannerRef, value *CGPDFBoolean) bool

// CGPDFScannerPopBoolean retrieves a Boolean object from the scanner stack.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopBoolean(_:_:)
func CGPDFScannerPopBoolean(scanner CGPDFScannerRef, value *CGPDFBoolean) bool {
	if _cGPDFScannerPopBoolean == nil {
		panic("CoreGraphics: symbol CGPDFScannerPopBoolean not loaded")
	}
	return _cGPDFScannerPopBoolean(scanner, value)
}

var _cGPDFScannerPopDictionary func(scanner CGPDFScannerRef, value *CGPDFDictionaryRef) bool

// CGPDFScannerPopDictionary retrieves a PDF dictionary object from the scanner stack.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopDictionary(_:_:)
func CGPDFScannerPopDictionary(scanner CGPDFScannerRef, value *CGPDFDictionaryRef) bool {
	if _cGPDFScannerPopDictionary == nil {
		panic("CoreGraphics: symbol CGPDFScannerPopDictionary not loaded")
	}
	return _cGPDFScannerPopDictionary(scanner, value)
}

var _cGPDFScannerPopInteger func(scanner CGPDFScannerRef, value *CGPDFInteger) bool

// CGPDFScannerPopInteger retrieves an integer object from the scanner stack.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopInteger(_:_:)
func CGPDFScannerPopInteger(scanner CGPDFScannerRef, value *CGPDFInteger) bool {
	if _cGPDFScannerPopInteger == nil {
		panic("CoreGraphics: symbol CGPDFScannerPopInteger not loaded")
	}
	return _cGPDFScannerPopInteger(scanner, value)
}

var _cGPDFScannerPopName func(scanner CGPDFScannerRef, value string) bool

// CGPDFScannerPopName retrieves a character string from the scanner stack.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopName(_:_:)
func CGPDFScannerPopName(scanner CGPDFScannerRef, value string) bool {
	if _cGPDFScannerPopName == nil {
		panic("CoreGraphics: symbol CGPDFScannerPopName not loaded")
	}
	return _cGPDFScannerPopName(scanner, value)
}

var _cGPDFScannerPopNumber func(scanner CGPDFScannerRef, value *CGPDFReal) bool

// CGPDFScannerPopNumber retrieves a real value object from the scanner stack.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopNumber(_:_:)
func CGPDFScannerPopNumber(scanner CGPDFScannerRef, value *CGPDFReal) bool {
	if _cGPDFScannerPopNumber == nil {
		panic("CoreGraphics: symbol CGPDFScannerPopNumber not loaded")
	}
	return _cGPDFScannerPopNumber(scanner, value)
}

var _cGPDFScannerPopObject func(scanner CGPDFScannerRef, value *CGPDFObjectRef) bool

// CGPDFScannerPopObject retrieves an object from the scanner stack.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopObject(_:_:)
func CGPDFScannerPopObject(scanner CGPDFScannerRef, value *CGPDFObjectRef) bool {
	if _cGPDFScannerPopObject == nil {
		panic("CoreGraphics: symbol CGPDFScannerPopObject not loaded")
	}
	return _cGPDFScannerPopObject(scanner, value)
}

var _cGPDFScannerPopStream func(scanner CGPDFScannerRef, value *CGPDFStreamRef) bool

// CGPDFScannerPopStream retrieves a PDF stream object from the scanner stack.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopStream(_:_:)
func CGPDFScannerPopStream(scanner CGPDFScannerRef, value *CGPDFStreamRef) bool {
	if _cGPDFScannerPopStream == nil {
		panic("CoreGraphics: symbol CGPDFScannerPopStream not loaded")
	}
	return _cGPDFScannerPopStream(scanner, value)
}

var _cGPDFScannerPopString func(scanner CGPDFScannerRef, value *CGPDFStringRef) bool

// CGPDFScannerPopString retrieves a string object from the scanner stack.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopString(_:_:)
func CGPDFScannerPopString(scanner CGPDFScannerRef, value *CGPDFStringRef) bool {
	if _cGPDFScannerPopString == nil {
		panic("CoreGraphics: symbol CGPDFScannerPopString not loaded")
	}
	return _cGPDFScannerPopString(scanner, value)
}

var _cGPDFScannerRelease func(scanner CGPDFScannerRef)

// CGPDFScannerRelease decrements the retain count of a scanner object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerRelease(_:)
func CGPDFScannerRelease(scanner CGPDFScannerRef) {
	if _cGPDFScannerRelease == nil {
		panic("CoreGraphics: symbol CGPDFScannerRelease not loaded")
	}
	_cGPDFScannerRelease(scanner)
}

var _cGPDFScannerRetain func(scanner CGPDFScannerRef) CGPDFScannerRef

// CGPDFScannerRetain increments the retain count of a scanner object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerRetain(_:)
func CGPDFScannerRetain(scanner CGPDFScannerRef) CGPDFScannerRef {
	if _cGPDFScannerRetain == nil {
		panic("CoreGraphics: symbol CGPDFScannerRetain not loaded")
	}
	return _cGPDFScannerRetain(scanner)
}

var _cGPDFScannerScan func(scanner CGPDFScannerRef) bool

// CGPDFScannerScan parses the content stream of a PDF scanner object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerScan(_:)
func CGPDFScannerScan(scanner CGPDFScannerRef) bool {
	if _cGPDFScannerScan == nil {
		panic("CoreGraphics: symbol CGPDFScannerScan not loaded")
	}
	return _cGPDFScannerScan(scanner)
}

var _cGPDFScannerStop func(s CGPDFScannerRef)

// CGPDFScannerStop.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerStop(_:)
func CGPDFScannerStop(s CGPDFScannerRef) {
	if _cGPDFScannerStop == nil {
		panic("CoreGraphics: symbol CGPDFScannerStop not loaded")
	}
	_cGPDFScannerStop(s)
}

var _cGPDFStreamCopyData func(stream CGPDFStreamRef, format *CGPDFDataFormat) corefoundation.CFDataRef

// CGPDFStreamCopyData returns the data associated with a PDF stream.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFStreamCopyData(_:_:)
func CGPDFStreamCopyData(stream CGPDFStreamRef, format *CGPDFDataFormat) corefoundation.CFDataRef {
	if _cGPDFStreamCopyData == nil {
		panic("CoreGraphics: symbol CGPDFStreamCopyData not loaded")
	}
	return _cGPDFStreamCopyData(stream, format)
}

var _cGPDFStreamGetDictionary func(stream CGPDFStreamRef) CGPDFDictionaryRef

// CGPDFStreamGetDictionary returns the dictionary associated with a PDF stream.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFStreamGetDictionary(_:)
func CGPDFStreamGetDictionary(stream CGPDFStreamRef) CGPDFDictionaryRef {
	if _cGPDFStreamGetDictionary == nil {
		panic("CoreGraphics: symbol CGPDFStreamGetDictionary not loaded")
	}
	return _cGPDFStreamGetDictionary(stream)
}

var _cGPDFStringCopyDate func(string_ CGPDFStringRef) corefoundation.CFDateRef

// CGPDFStringCopyDate converts a string to a date.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFStringCopyDate(_:)
func CGPDFStringCopyDate(string_ CGPDFStringRef) corefoundation.CFDateRef {
	if _cGPDFStringCopyDate == nil {
		panic("CoreGraphics: symbol CGPDFStringCopyDate not loaded")
	}
	return _cGPDFStringCopyDate(string_)
}

var _cGPDFStringCopyTextString func(string_ CGPDFStringRef) corefoundation.CFStringRef

// CGPDFStringCopyTextString returns a CFString object that represents a PDF string as a text string.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFStringCopyTextString(_:)
func CGPDFStringCopyTextString(string_ CGPDFStringRef) corefoundation.CFStringRef {
	if _cGPDFStringCopyTextString == nil {
		panic("CoreGraphics: symbol CGPDFStringCopyTextString not loaded")
	}
	return _cGPDFStringCopyTextString(string_)
}

var _cGPDFStringGetBytePtr func(string_ CGPDFStringRef) *byte

// CGPDFStringGetBytePtr returns a pointer to the bytes of a PDF string.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFStringGetBytePtr(_:)
func CGPDFStringGetBytePtr(string_ CGPDFStringRef) *byte {
	if _cGPDFStringGetBytePtr == nil {
		panic("CoreGraphics: symbol CGPDFStringGetBytePtr not loaded")
	}
	return _cGPDFStringGetBytePtr(string_)
}

var _cGPDFStringGetLength func(string_ CGPDFStringRef) uintptr

// CGPDFStringGetLength returns the number of bytes in a PDF string.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFStringGetLength(_:)
func CGPDFStringGetLength(string_ CGPDFStringRef) uintptr {
	if _cGPDFStringGetLength == nil {
		panic("CoreGraphics: symbol CGPDFStringGetLength not loaded")
	}
	return _cGPDFStringGetLength(string_)
}

var _cGPSConverterAbort func(converter CGPSConverterRef) bool

// CGPSConverterAbort tells a PostScript converter to abort a conversion at the next available opportunity.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPSConverter/abort()
func CGPSConverterAbort(converter CGPSConverterRef) bool {
	if _cGPSConverterAbort == nil {
		panic("CoreGraphics: symbol CGPSConverterAbort not loaded")
	}
	return _cGPSConverterAbort(converter)
}

var _cGPSConverterConvert func(converter CGPSConverterRef, provider CGDataProviderRef, consumer CGDataConsumerRef, options corefoundation.CFDictionaryRef) bool

// CGPSConverterConvert uses a PostScript converter to convert PostScript data to PDF data.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPSConverter/convert(_:consumer:options:)
func CGPSConverterConvert(converter CGPSConverterRef, provider CGDataProviderRef, consumer CGDataConsumerRef, options corefoundation.CFDictionaryRef) bool {
	if _cGPSConverterConvert == nil {
		panic("CoreGraphics: symbol CGPSConverterConvert not loaded")
	}
	return _cGPSConverterConvert(converter, provider, consumer, options)
}

var _cGPSConverterCreate func(info unsafe.Pointer, callbacks *CGPSConverterCallbacks, options corefoundation.CFDictionaryRef) CGPSConverterRef

// CGPSConverterCreate creates a new PostScript converter.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPSConverter/init(info:callbacks:options:)
func CGPSConverterCreate(info unsafe.Pointer, callbacks *CGPSConverterCallbacks, options corefoundation.CFDictionaryRef) CGPSConverterRef {
	if _cGPSConverterCreate == nil {
		panic("CoreGraphics: symbol CGPSConverterCreate not loaded")
	}
	return _cGPSConverterCreate(info, callbacks, options)
}

var _cGPSConverterGetTypeID func() uint

// CGPSConverterGetTypeID returns the Core Foundation type identifier for PostScript converters.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPSConverter/typeID
func CGPSConverterGetTypeID() uint {
	if _cGPSConverterGetTypeID == nil {
		panic("CoreGraphics: symbol CGPSConverterGetTypeID not loaded")
	}
	return _cGPSConverterGetTypeID()
}

var _cGPSConverterIsConverting func(converter CGPSConverterRef) bool

// CGPSConverterIsConverting checks whether the converter is currently converting data.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPSConverter/isConverting
func CGPSConverterIsConverting(converter CGPSConverterRef) bool {
	if _cGPSConverterIsConverting == nil {
		panic("CoreGraphics: symbol CGPSConverterIsConverting not loaded")
	}
	return _cGPSConverterIsConverting(converter)
}

var _cGPathApply func(path CGPathRef, info unsafe.Pointer, function CGPathApplierFunction)

// CGPathApply for each element in a graphics path, calls a custom applier function.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/apply(info:function:)
func CGPathApply(path CGPathRef, info unsafe.Pointer, function CGPathApplierFunction) {
	if _cGPathApply == nil {
		panic("CoreGraphics: symbol CGPathApply not loaded")
	}
	_cGPathApply(path, info, function)
}

var _cGPathApplyWithBlock func(path CGPathRef, block CGPathApplyBlock)

// CGPathApplyWithBlock.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/applyWithBlock(_:)
func CGPathApplyWithBlock(path CGPathRef, block CGPathApplyBlock) {
	if _cGPathApplyWithBlock == nil {
		panic("CoreGraphics: symbol CGPathApplyWithBlock not loaded")
	}
	_cGPathApplyWithBlock(path, block)
}

var _cGPathCloseSubpath func(path CGMutablePathRef)

// CGPathCloseSubpath closes and completes a subpath in a mutable graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGMutablePath/closeSubpath()
func CGPathCloseSubpath(path CGMutablePathRef) {
	if _cGPathCloseSubpath == nil {
		panic("CoreGraphics: symbol CGPathCloseSubpath not loaded")
	}
	_cGPathCloseSubpath(path)
}

var _cGPathContainsPoint func(path CGPathRef, m *corefoundation.CGAffineTransform, point corefoundation.CGPoint, eoFill bool) bool

// CGPathContainsPoint checks whether a point is contained in a graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathContainsPoint
func CGPathContainsPoint(path CGPathRef, m *corefoundation.CGAffineTransform, point corefoundation.CGPoint, eoFill bool) bool {
	if _cGPathContainsPoint == nil {
		panic("CoreGraphics: symbol CGPathContainsPoint not loaded")
	}
	return _cGPathContainsPoint(path, m, point, eoFill)
}

var _cGPathCreateCopy func(path CGPathRef) CGPathRef

// CGPathCreateCopy creates an immutable copy of a graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/copy()
func CGPathCreateCopy(path CGPathRef) CGPathRef {
	if _cGPathCreateCopy == nil {
		panic("CoreGraphics: symbol CGPathCreateCopy not loaded")
	}
	return _cGPathCreateCopy(path)
}

var _cGPathCreateCopyByDashingPath func(path CGPathRef, transform *corefoundation.CGAffineTransform, phase float64, lengths *float64, count uintptr) CGPathRef

// CGPathCreateCopyByDashingPath creates a dashed copy of another path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyByDashingPath
func CGPathCreateCopyByDashingPath(path CGPathRef, transform *corefoundation.CGAffineTransform, phase float64, lengths *float64, count uintptr) CGPathRef {
	if _cGPathCreateCopyByDashingPath == nil {
		panic("CoreGraphics: symbol CGPathCreateCopyByDashingPath not loaded")
	}
	return _cGPathCreateCopyByDashingPath(path, transform, phase, lengths, count)
}

var _cGPathCreateCopyByFlattening func(path CGPathRef, flatteningThreshold float64) CGPathRef

// CGPathCreateCopyByFlattening.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyByFlattening
func CGPathCreateCopyByFlattening(path CGPathRef, flatteningThreshold float64) CGPathRef {
	if _cGPathCreateCopyByFlattening == nil {
		panic("CoreGraphics: symbol CGPathCreateCopyByFlattening not loaded")
	}
	return _cGPathCreateCopyByFlattening(path, flatteningThreshold)
}

var _cGPathCreateCopyByIntersectingPath func(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef

// CGPathCreateCopyByIntersectingPath.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyByIntersectingPath
func CGPathCreateCopyByIntersectingPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef {
	if _cGPathCreateCopyByIntersectingPath == nil {
		panic("CoreGraphics: symbol CGPathCreateCopyByIntersectingPath not loaded")
	}
	return _cGPathCreateCopyByIntersectingPath(path, maskPath, evenOddFillRule)
}

var _cGPathCreateCopyByNormalizing func(path CGPathRef, evenOddFillRule bool) CGPathRef

// CGPathCreateCopyByNormalizing.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyByNormalizing
func CGPathCreateCopyByNormalizing(path CGPathRef, evenOddFillRule bool) CGPathRef {
	if _cGPathCreateCopyByNormalizing == nil {
		panic("CoreGraphics: symbol CGPathCreateCopyByNormalizing not loaded")
	}
	return _cGPathCreateCopyByNormalizing(path, evenOddFillRule)
}

var _cGPathCreateCopyByStrokingPath func(path CGPathRef, transform *corefoundation.CGAffineTransform, lineWidth float64, lineCap uintptr, lineJoin uintptr, miterLimit float64) CGPathRef

// CGPathCreateCopyByStrokingPath creates a stroked copy of another path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyByStrokingPath
func CGPathCreateCopyByStrokingPath(path CGPathRef, transform *corefoundation.CGAffineTransform, lineWidth float64, lineCap uintptr, lineJoin uintptr, miterLimit float64) CGPathRef {
	if _cGPathCreateCopyByStrokingPath == nil {
		panic("CoreGraphics: symbol CGPathCreateCopyByStrokingPath not loaded")
	}
	return _cGPathCreateCopyByStrokingPath(path, transform, lineWidth, lineCap, lineJoin, miterLimit)
}

var _cGPathCreateCopyBySubtractingPath func(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef

// CGPathCreateCopyBySubtractingPath.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyBySubtractingPath
func CGPathCreateCopyBySubtractingPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef {
	if _cGPathCreateCopyBySubtractingPath == nil {
		panic("CoreGraphics: symbol CGPathCreateCopyBySubtractingPath not loaded")
	}
	return _cGPathCreateCopyBySubtractingPath(path, maskPath, evenOddFillRule)
}

var _cGPathCreateCopyBySymmetricDifferenceOfPath func(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef

// CGPathCreateCopyBySymmetricDifferenceOfPath.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyBySymmetricDifferenceOfPath
func CGPathCreateCopyBySymmetricDifferenceOfPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef {
	if _cGPathCreateCopyBySymmetricDifferenceOfPath == nil {
		panic("CoreGraphics: symbol CGPathCreateCopyBySymmetricDifferenceOfPath not loaded")
	}
	return _cGPathCreateCopyBySymmetricDifferenceOfPath(path, maskPath, evenOddFillRule)
}

var _cGPathCreateCopyByTransformingPath func(path CGPathRef, transform *corefoundation.CGAffineTransform) CGPathRef

// CGPathCreateCopyByTransformingPath creates an immutable copy of a graphics path transformed by a transformation matrix.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/copy(using:)
func CGPathCreateCopyByTransformingPath(path CGPathRef, transform *corefoundation.CGAffineTransform) CGPathRef {
	if _cGPathCreateCopyByTransformingPath == nil {
		panic("CoreGraphics: symbol CGPathCreateCopyByTransformingPath not loaded")
	}
	return _cGPathCreateCopyByTransformingPath(path, transform)
}

var _cGPathCreateCopyByUnioningPath func(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef

// CGPathCreateCopyByUnioningPath.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyByUnioningPath
func CGPathCreateCopyByUnioningPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef {
	if _cGPathCreateCopyByUnioningPath == nil {
		panic("CoreGraphics: symbol CGPathCreateCopyByUnioningPath not loaded")
	}
	return _cGPathCreateCopyByUnioningPath(path, maskPath, evenOddFillRule)
}

var _cGPathCreateCopyOfLineByIntersectingPath func(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef

// CGPathCreateCopyOfLineByIntersectingPath.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyOfLineByIntersectingPath
func CGPathCreateCopyOfLineByIntersectingPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef {
	if _cGPathCreateCopyOfLineByIntersectingPath == nil {
		panic("CoreGraphics: symbol CGPathCreateCopyOfLineByIntersectingPath not loaded")
	}
	return _cGPathCreateCopyOfLineByIntersectingPath(path, maskPath, evenOddFillRule)
}

var _cGPathCreateCopyOfLineBySubtractingPath func(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef

// CGPathCreateCopyOfLineBySubtractingPath.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyOfLineBySubtractingPath
func CGPathCreateCopyOfLineBySubtractingPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef {
	if _cGPathCreateCopyOfLineBySubtractingPath == nil {
		panic("CoreGraphics: symbol CGPathCreateCopyOfLineBySubtractingPath not loaded")
	}
	return _cGPathCreateCopyOfLineBySubtractingPath(path, maskPath, evenOddFillRule)
}

var _cGPathCreateMutable func() CGMutablePathRef

// CGPathCreateMutable creates a mutable graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGMutablePath/init()
func CGPathCreateMutable() CGMutablePathRef {
	if _cGPathCreateMutable == nil {
		panic("CoreGraphics: symbol CGPathCreateMutable not loaded")
	}
	return _cGPathCreateMutable()
}

var _cGPathCreateMutableCopy func(path CGPathRef) CGMutablePathRef

// CGPathCreateMutableCopy creates a mutable copy of an existing graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/mutableCopy()
func CGPathCreateMutableCopy(path CGPathRef) CGMutablePathRef {
	if _cGPathCreateMutableCopy == nil {
		panic("CoreGraphics: symbol CGPathCreateMutableCopy not loaded")
	}
	return _cGPathCreateMutableCopy(path)
}

var _cGPathCreateMutableCopyByTransformingPath func(path CGPathRef, transform *corefoundation.CGAffineTransform) CGMutablePathRef

// CGPathCreateMutableCopyByTransformingPath creates a mutable copy of a graphics path transformed by a transformation matrix.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/mutableCopy(using:)
func CGPathCreateMutableCopyByTransformingPath(path CGPathRef, transform *corefoundation.CGAffineTransform) CGMutablePathRef {
	if _cGPathCreateMutableCopyByTransformingPath == nil {
		panic("CoreGraphics: symbol CGPathCreateMutableCopyByTransformingPath not loaded")
	}
	return _cGPathCreateMutableCopyByTransformingPath(path, transform)
}

var _cGPathCreateSeparateComponents func(path CGPathRef, evenOddFillRule bool) corefoundation.CFArrayRef

// CGPathCreateSeparateComponents.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateSeparateComponents
func CGPathCreateSeparateComponents(path CGPathRef, evenOddFillRule bool) corefoundation.CFArrayRef {
	if _cGPathCreateSeparateComponents == nil {
		panic("CoreGraphics: symbol CGPathCreateSeparateComponents not loaded")
	}
	return _cGPathCreateSeparateComponents(path, evenOddFillRule)
}

var _cGPathCreateWithEllipseInRect func(rect corefoundation.CGRect, transform *corefoundation.CGAffineTransform) CGPathRef

// CGPathCreateWithEllipseInRect create an immutable path of an ellipse.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/init(ellipseIn:transform:)
func CGPathCreateWithEllipseInRect(rect corefoundation.CGRect, transform *corefoundation.CGAffineTransform) CGPathRef {
	if _cGPathCreateWithEllipseInRect == nil {
		panic("CoreGraphics: symbol CGPathCreateWithEllipseInRect not loaded")
	}
	return _cGPathCreateWithEllipseInRect(rect, transform)
}

var _cGPathCreateWithRect func(rect corefoundation.CGRect, transform *corefoundation.CGAffineTransform) CGPathRef

// CGPathCreateWithRect create an immutable path of a rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/init(rect:transform:)
func CGPathCreateWithRect(rect corefoundation.CGRect, transform *corefoundation.CGAffineTransform) CGPathRef {
	if _cGPathCreateWithRect == nil {
		panic("CoreGraphics: symbol CGPathCreateWithRect not loaded")
	}
	return _cGPathCreateWithRect(rect, transform)
}

var _cGPathCreateWithRoundedRect func(rect corefoundation.CGRect, cornerWidth float64, cornerHeight float64, transform *corefoundation.CGAffineTransform) CGPathRef

// CGPathCreateWithRoundedRect create an immutable path of a rounded rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/init(roundedRect:cornerWidth:cornerHeight:transform:)
func CGPathCreateWithRoundedRect(rect corefoundation.CGRect, cornerWidth float64, cornerHeight float64, transform *corefoundation.CGAffineTransform) CGPathRef {
	if _cGPathCreateWithRoundedRect == nil {
		panic("CoreGraphics: symbol CGPathCreateWithRoundedRect not loaded")
	}
	return _cGPathCreateWithRoundedRect(rect, cornerWidth, cornerHeight, transform)
}

var _cGPathEqualToPath func(path1 CGPathRef, path2 CGPathRef) bool

// CGPathEqualToPath indicates whether two graphics paths are equivalent.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathEqualToPath
func CGPathEqualToPath(path1 CGPathRef, path2 CGPathRef) bool {
	if _cGPathEqualToPath == nil {
		panic("CoreGraphics: symbol CGPathEqualToPath not loaded")
	}
	return _cGPathEqualToPath(path1, path2)
}

var _cGPathGetBoundingBox func(path CGPathRef) corefoundation.CGRect

// CGPathGetBoundingBox returns the bounding box containing all points in a graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/boundingBox
func CGPathGetBoundingBox(path CGPathRef) corefoundation.CGRect {
	if _cGPathGetBoundingBox == nil {
		panic("CoreGraphics: symbol CGPathGetBoundingBox not loaded")
	}
	return _cGPathGetBoundingBox(path)
}

var _cGPathGetCurrentPoint func(path CGPathRef) corefoundation.CGPoint

// CGPathGetCurrentPoint returns the current point in a graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/currentPoint
func CGPathGetCurrentPoint(path CGPathRef) corefoundation.CGPoint {
	if _cGPathGetCurrentPoint == nil {
		panic("CoreGraphics: symbol CGPathGetCurrentPoint not loaded")
	}
	return _cGPathGetCurrentPoint(path)
}

var _cGPathGetPathBoundingBox func(path CGPathRef) corefoundation.CGRect

// CGPathGetPathBoundingBox returns the bounding box of a graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/boundingBoxOfPath
func CGPathGetPathBoundingBox(path CGPathRef) corefoundation.CGRect {
	if _cGPathGetPathBoundingBox == nil {
		panic("CoreGraphics: symbol CGPathGetPathBoundingBox not loaded")
	}
	return _cGPathGetPathBoundingBox(path)
}

var _cGPathGetTypeID func() uint

// CGPathGetTypeID returns the Core Foundation type identifier for Core Graphics paths.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/typeID
func CGPathGetTypeID() uint {
	if _cGPathGetTypeID == nil {
		panic("CoreGraphics: symbol CGPathGetTypeID not loaded")
	}
	return _cGPathGetTypeID()
}

var _cGPathIntersectsPath func(path1 CGPathRef, path2 CGPathRef, evenOddFillRule bool) bool

// CGPathIntersectsPath.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathIntersectsPath
func CGPathIntersectsPath(path1 CGPathRef, path2 CGPathRef, evenOddFillRule bool) bool {
	if _cGPathIntersectsPath == nil {
		panic("CoreGraphics: symbol CGPathIntersectsPath not loaded")
	}
	return _cGPathIntersectsPath(path1, path2, evenOddFillRule)
}

var _cGPathIsEmpty func(path CGPathRef) bool

// CGPathIsEmpty indicates whether or not a graphics path is empty.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/isEmpty
func CGPathIsEmpty(path CGPathRef) bool {
	if _cGPathIsEmpty == nil {
		panic("CoreGraphics: symbol CGPathIsEmpty not loaded")
	}
	return _cGPathIsEmpty(path)
}

var _cGPathIsRect func(path CGPathRef, rect *corefoundation.CGRect) bool

// CGPathIsRect indicates whether or not a graphics path represents a rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/isRect(_:)
func CGPathIsRect(path CGPathRef, rect *corefoundation.CGRect) bool {
	if _cGPathIsRect == nil {
		panic("CoreGraphics: symbol CGPathIsRect not loaded")
	}
	return _cGPathIsRect(path, rect)
}

var _cGPathRelease func(path CGPathRef)

// CGPathRelease decrements the retain count of a graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathRelease
func CGPathRelease(path CGPathRef) {
	if _cGPathRelease == nil {
		panic("CoreGraphics: symbol CGPathRelease not loaded")
	}
	_cGPathRelease(path)
}

var _cGPathRetain func(path CGPathRef) CGPathRef

// CGPathRetain increments the retain count of a graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathRetain
func CGPathRetain(path CGPathRef) CGPathRef {
	if _cGPathRetain == nil {
		panic("CoreGraphics: symbol CGPathRetain not loaded")
	}
	return _cGPathRetain(path)
}

var _cGPatternCreate func(info unsafe.Pointer, bounds corefoundation.CGRect, matrix corefoundation.CGAffineTransform, xStep float64, yStep float64, tiling CGPatternTiling, isColored bool, callbacks *CGPatternCallbacks) CGPatternRef

// CGPatternCreate creates a pattern object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPattern/init(info:bounds:matrix:xStep:yStep:tiling:isColored:callbacks:)
func CGPatternCreate(info unsafe.Pointer, bounds corefoundation.CGRect, matrix corefoundation.CGAffineTransform, xStep float64, yStep float64, tiling CGPatternTiling, isColored bool, callbacks *CGPatternCallbacks) CGPatternRef {
	if _cGPatternCreate == nil {
		panic("CoreGraphics: symbol CGPatternCreate not loaded")
	}
	return _cGPatternCreate(info, bounds, matrix, xStep, yStep, tiling, isColored, callbacks)
}

var _cGPatternGetTypeID func() uint

// CGPatternGetTypeID returns the type identifier for Core Graphics patterns.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPattern/typeID
func CGPatternGetTypeID() uint {
	if _cGPatternGetTypeID == nil {
		panic("CoreGraphics: symbol CGPatternGetTypeID not loaded")
	}
	return _cGPatternGetTypeID()
}

var _cGPatternRelease func(pattern CGPatternRef)

// CGPatternRelease decrements the retain count of a Core Graphics pattern.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPatternRelease
func CGPatternRelease(pattern CGPatternRef) {
	if _cGPatternRelease == nil {
		panic("CoreGraphics: symbol CGPatternRelease not loaded")
	}
	_cGPatternRelease(pattern)
}

var _cGPatternRetain func(pattern CGPatternRef) CGPatternRef

// CGPatternRetain increments the retain count of a Core Graphics pattern.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPatternRetain
func CGPatternRetain(pattern CGPatternRef) CGPatternRef {
	if _cGPatternRetain == nil {
		panic("CoreGraphics: symbol CGPatternRetain not loaded")
	}
	return _cGPatternRetain(pattern)
}

var _cGPointApplyAffineTransform func(point corefoundation.CGPoint, t corefoundation.CGAffineTransform) corefoundation.CGPoint

// CGPointApplyAffineTransform returns the point resulting from an affine transformation of an existing point.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPointApplyAffineTransform(_:_:)
func CGPointApplyAffineTransform(point corefoundation.CGPoint, t corefoundation.CGAffineTransform) corefoundation.CGPoint {
	if _cGPointApplyAffineTransform == nil {
		panic("CoreGraphics: symbol CGPointApplyAffineTransform not loaded")
	}
	return _cGPointApplyAffineTransform(point, t)
}

var _cGPointCreateDictionaryRepresentation func(point corefoundation.CGPoint) corefoundation.CFDictionaryRef

// CGPointCreateDictionaryRepresentation returns a dictionary representation of the specified point.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPointCreateDictionaryRepresentation(_:)
func CGPointCreateDictionaryRepresentation(point corefoundation.CGPoint) corefoundation.CFDictionaryRef {
	if _cGPointCreateDictionaryRepresentation == nil {
		panic("CoreGraphics: symbol CGPointCreateDictionaryRepresentation not loaded")
	}
	return _cGPointCreateDictionaryRepresentation(point)
}

var _cGPointEqualToPoint func(point1 corefoundation.CGPoint, point2 corefoundation.CGPoint) bool

// CGPointEqualToPoint returns whether two points are equal.
//
// Deprecated: The [CGPoint](<doc://com.apple.documentation/documentation/CoreFoundation/CGPoint>) type adopts the [Equatable] protocol; use the `==` operator instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPointEqualToPoint(_:_:)
func CGPointEqualToPoint(point1 corefoundation.CGPoint, point2 corefoundation.CGPoint) bool {
	if _cGPointEqualToPoint == nil {
		panic("CoreGraphics: symbol CGPointEqualToPoint not loaded")
	}
	return _cGPointEqualToPoint(point1, point2)
}

var _cGPointMakeWithDictionaryRepresentation func(dict corefoundation.CFDictionaryRef, point *corefoundation.CGPoint) bool

// CGPointMakeWithDictionaryRepresentation fills in a point using the contents of the specified dictionary.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPointMakeWithDictionaryRepresentation(_:_:)
func CGPointMakeWithDictionaryRepresentation(dict corefoundation.CFDictionaryRef, point *corefoundation.CGPoint) bool {
	if _cGPointMakeWithDictionaryRepresentation == nil {
		panic("CoreGraphics: symbol CGPointMakeWithDictionaryRepresentation not loaded")
	}
	return _cGPointMakeWithDictionaryRepresentation(dict, point)
}

var _cGPostMouseEvent func(mouseCursorPosition corefoundation.CGPoint, updateMouseCursorPosition uintptr, buttonCount CGButtonCount, mouseButtonDown uintptr) CGError

// CGPostMouseEvent synthesizes a low-level mouse-button event on the local machine.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPostMouseEvent
func CGPostMouseEvent(mouseCursorPosition corefoundation.CGPoint, updateMouseCursorPosition uintptr, buttonCount CGButtonCount, mouseButtonDown uintptr) CGError {
	if _cGPostMouseEvent == nil {
		panic("CoreGraphics: symbol CGPostMouseEvent not loaded")
	}
	return _cGPostMouseEvent(mouseCursorPosition, updateMouseCursorPosition, buttonCount, mouseButtonDown)
}

var _cGPostScrollWheelEvent func(wheelCount CGWheelCount, wheel1 int32) CGError

// CGPostScrollWheelEvent synthesizes a low-level scrolling event on the local machine.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPostScrollWheelEvent
func CGPostScrollWheelEvent(wheelCount CGWheelCount, wheel1 int32) CGError {
	if _cGPostScrollWheelEvent == nil {
		panic("CoreGraphics: symbol CGPostScrollWheelEvent not loaded")
	}
	return _cGPostScrollWheelEvent(wheelCount, wheel1)
}

var _cGPreflightListenEventAccess func() bool

// CGPreflightListenEventAccess.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPreflightListenEventAccess()
func CGPreflightListenEventAccess() bool {
	if _cGPreflightListenEventAccess == nil {
		panic("CoreGraphics: symbol CGPreflightListenEventAccess not loaded")
	}
	return _cGPreflightListenEventAccess()
}

var _cGPreflightPostEventAccess func() bool

// CGPreflightPostEventAccess.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPreflightPostEventAccess()
func CGPreflightPostEventAccess() bool {
	if _cGPreflightPostEventAccess == nil {
		panic("CoreGraphics: symbol CGPreflightPostEventAccess not loaded")
	}
	return _cGPreflightPostEventAccess()
}

var _cGPreflightScreenCaptureAccess func() bool

// CGPreflightScreenCaptureAccess.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPreflightScreenCaptureAccess()
func CGPreflightScreenCaptureAccess() bool {
	if _cGPreflightScreenCaptureAccess == nil {
		panic("CoreGraphics: symbol CGPreflightScreenCaptureAccess not loaded")
	}
	return _cGPreflightScreenCaptureAccess()
}

var _cGRectApplyAffineTransform func(rect corefoundation.CGRect, t corefoundation.CGAffineTransform) corefoundation.CGRect

// CGRectApplyAffineTransform applies an affine transform to a rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectApplyAffineTransform(_:_:)
func CGRectApplyAffineTransform(rect corefoundation.CGRect, t corefoundation.CGAffineTransform) corefoundation.CGRect {
	if _cGRectApplyAffineTransform == nil {
		panic("CoreGraphics: symbol CGRectApplyAffineTransform not loaded")
	}
	return _cGRectApplyAffineTransform(rect, t)
}

var _cGRectContainsPoint func(rect corefoundation.CGRect, point corefoundation.CGPoint) bool

// CGRectContainsPoint returns whether a rectangle contains a specified point.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectContainsPoint(_:_:)
func CGRectContainsPoint(rect corefoundation.CGRect, point corefoundation.CGPoint) bool {
	if _cGRectContainsPoint == nil {
		panic("CoreGraphics: symbol CGRectContainsPoint not loaded")
	}
	return _cGRectContainsPoint(rect, point)
}

var _cGRectContainsRect func(rect1 corefoundation.CGRect, rect2 corefoundation.CGRect) bool

// CGRectContainsRect returns whether the first rectangle contains the second rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectContainsRect(_:_:)
func CGRectContainsRect(rect1 corefoundation.CGRect, rect2 corefoundation.CGRect) bool {
	if _cGRectContainsRect == nil {
		panic("CoreGraphics: symbol CGRectContainsRect not loaded")
	}
	return _cGRectContainsRect(rect1, rect2)
}

var _cGRectCreateDictionaryRepresentation func(arg0 corefoundation.CGRect) corefoundation.CFDictionaryRef

// CGRectCreateDictionaryRepresentation returns a dictionary representation of the provided rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectCreateDictionaryRepresentation(_:)
func CGRectCreateDictionaryRepresentation(arg0 corefoundation.CGRect) corefoundation.CFDictionaryRef {
	if _cGRectCreateDictionaryRepresentation == nil {
		panic("CoreGraphics: symbol CGRectCreateDictionaryRepresentation not loaded")
	}
	return _cGRectCreateDictionaryRepresentation(arg0)
}

var _cGRectDivide func(rect corefoundation.CGRect, slice *corefoundation.CGRect, remainder *corefoundation.CGRect, amount float64, edge CGRectEdge)

// CGRectDivide divides a source rectangle into two component rectangles.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectDivide
func CGRectDivide(rect corefoundation.CGRect, slice *corefoundation.CGRect, remainder *corefoundation.CGRect, amount float64, edge CGRectEdge) {
	if _cGRectDivide == nil {
		panic("CoreGraphics: symbol CGRectDivide not loaded")
	}
	_cGRectDivide(rect, slice, remainder, amount, edge)
}

var _cGRectEqualToRect func(rect1 corefoundation.CGRect, rect2 corefoundation.CGRect) bool

// CGRectEqualToRect returns whether two rectangles are equal in size and position.
//
// Deprecated: The [CGRect](<doc://com.apple.documentation/documentation/CoreFoundation/CGRect>) type adopts the [Equatable] protocol; use the `==` operator instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectEqualToRect(_:_:)
func CGRectEqualToRect(rect1 corefoundation.CGRect, rect2 corefoundation.CGRect) bool {
	if _cGRectEqualToRect == nil {
		panic("CoreGraphics: symbol CGRectEqualToRect not loaded")
	}
	return _cGRectEqualToRect(rect1, rect2)
}

var _cGRectGetHeight func(rect corefoundation.CGRect) float64

// CGRectGetHeight returns the height of a rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectGetHeight(_:)
func CGRectGetHeight(rect corefoundation.CGRect) float64 {
	if _cGRectGetHeight == nil {
		panic("CoreGraphics: symbol CGRectGetHeight not loaded")
	}
	return _cGRectGetHeight(rect)
}

var _cGRectGetMaxX func(rect corefoundation.CGRect) float64

// CGRectGetMaxX returns the largest value of the x-coordinate for the rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectGetMaxX(_:)
func CGRectGetMaxX(rect corefoundation.CGRect) float64 {
	if _cGRectGetMaxX == nil {
		panic("CoreGraphics: symbol CGRectGetMaxX not loaded")
	}
	return _cGRectGetMaxX(rect)
}

var _cGRectGetMaxY func(rect corefoundation.CGRect) float64

// CGRectGetMaxY returns the largest value for the y-coordinate of the rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectGetMaxY(_:)
func CGRectGetMaxY(rect corefoundation.CGRect) float64 {
	if _cGRectGetMaxY == nil {
		panic("CoreGraphics: symbol CGRectGetMaxY not loaded")
	}
	return _cGRectGetMaxY(rect)
}

var _cGRectGetMidX func(rect corefoundation.CGRect) float64

// CGRectGetMidX returns the x- coordinate that establishes the center of a rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectGetMidX(_:)
func CGRectGetMidX(rect corefoundation.CGRect) float64 {
	if _cGRectGetMidX == nil {
		panic("CoreGraphics: symbol CGRectGetMidX not loaded")
	}
	return _cGRectGetMidX(rect)
}

var _cGRectGetMidY func(rect corefoundation.CGRect) float64

// CGRectGetMidY returns the y-coordinate that establishes the center of the rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectGetMidY(_:)
func CGRectGetMidY(rect corefoundation.CGRect) float64 {
	if _cGRectGetMidY == nil {
		panic("CoreGraphics: symbol CGRectGetMidY not loaded")
	}
	return _cGRectGetMidY(rect)
}

var _cGRectGetMinX func(rect corefoundation.CGRect) float64

// CGRectGetMinX returns the smallest value for the x-coordinate of the rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectGetMinX(_:)
func CGRectGetMinX(rect corefoundation.CGRect) float64 {
	if _cGRectGetMinX == nil {
		panic("CoreGraphics: symbol CGRectGetMinX not loaded")
	}
	return _cGRectGetMinX(rect)
}

var _cGRectGetMinY func(rect corefoundation.CGRect) float64

// CGRectGetMinY returns the smallest value for the y-coordinate of the rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectGetMinY(_:)
func CGRectGetMinY(rect corefoundation.CGRect) float64 {
	if _cGRectGetMinY == nil {
		panic("CoreGraphics: symbol CGRectGetMinY not loaded")
	}
	return _cGRectGetMinY(rect)
}

var _cGRectGetWidth func(rect corefoundation.CGRect) float64

// CGRectGetWidth returns the width of a rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectGetWidth(_:)
func CGRectGetWidth(rect corefoundation.CGRect) float64 {
	if _cGRectGetWidth == nil {
		panic("CoreGraphics: symbol CGRectGetWidth not loaded")
	}
	return _cGRectGetWidth(rect)
}

var _cGRectInset func(rect corefoundation.CGRect, dx float64, dy float64) corefoundation.CGRect

// CGRectInset returns a rectangle that is smaller or larger than the source rectangle, with the same center point.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectInset(_:_:_:)
func CGRectInset(rect corefoundation.CGRect, dx float64, dy float64) corefoundation.CGRect {
	if _cGRectInset == nil {
		panic("CoreGraphics: symbol CGRectInset not loaded")
	}
	return _cGRectInset(rect, dx, dy)
}

var _cGRectIntegral func(rect corefoundation.CGRect) corefoundation.CGRect

// CGRectIntegral returns the smallest rectangle that results from converting the source rectangle values to integers.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectIntegral(_:)
func CGRectIntegral(rect corefoundation.CGRect) corefoundation.CGRect {
	if _cGRectIntegral == nil {
		panic("CoreGraphics: symbol CGRectIntegral not loaded")
	}
	return _cGRectIntegral(rect)
}

var _cGRectIntersection func(r1 corefoundation.CGRect, r2 corefoundation.CGRect) corefoundation.CGRect

// CGRectIntersection returns the intersection of two rectangles.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectIntersection(_:_:)
func CGRectIntersection(r1 corefoundation.CGRect, r2 corefoundation.CGRect) corefoundation.CGRect {
	if _cGRectIntersection == nil {
		panic("CoreGraphics: symbol CGRectIntersection not loaded")
	}
	return _cGRectIntersection(r1, r2)
}

var _cGRectIntersectsRect func(rect1 corefoundation.CGRect, rect2 corefoundation.CGRect) bool

// CGRectIntersectsRect returns whether two rectangles intersect.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectIntersectsRect(_:_:)
func CGRectIntersectsRect(rect1 corefoundation.CGRect, rect2 corefoundation.CGRect) bool {
	if _cGRectIntersectsRect == nil {
		panic("CoreGraphics: symbol CGRectIntersectsRect not loaded")
	}
	return _cGRectIntersectsRect(rect1, rect2)
}

var _cGRectIsEmpty func(rect corefoundation.CGRect) bool

// CGRectIsEmpty returns whether a rectangle has zero width or height, or is a null rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectIsEmpty(_:)
func CGRectIsEmpty(rect corefoundation.CGRect) bool {
	if _cGRectIsEmpty == nil {
		panic("CoreGraphics: symbol CGRectIsEmpty not loaded")
	}
	return _cGRectIsEmpty(rect)
}

var _cGRectIsInfinite func(rect corefoundation.CGRect) bool

// CGRectIsInfinite returns whether a rectangle is infinite.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectIsInfinite(_:)
func CGRectIsInfinite(rect corefoundation.CGRect) bool {
	if _cGRectIsInfinite == nil {
		panic("CoreGraphics: symbol CGRectIsInfinite not loaded")
	}
	return _cGRectIsInfinite(rect)
}

var _cGRectIsNull func(rect corefoundation.CGRect) bool

// CGRectIsNull returns whether the rectangle is equal to the null rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectIsNull(_:)
func CGRectIsNull(rect corefoundation.CGRect) bool {
	if _cGRectIsNull == nil {
		panic("CoreGraphics: symbol CGRectIsNull not loaded")
	}
	return _cGRectIsNull(rect)
}

var _cGRectMakeWithDictionaryRepresentation func(dict corefoundation.CFDictionaryRef, rect *corefoundation.CGRect) bool

// CGRectMakeWithDictionaryRepresentation fills in a rectangle using the contents of the specified dictionary.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectMakeWithDictionaryRepresentation(_:_:)
func CGRectMakeWithDictionaryRepresentation(dict corefoundation.CFDictionaryRef, rect *corefoundation.CGRect) bool {
	if _cGRectMakeWithDictionaryRepresentation == nil {
		panic("CoreGraphics: symbol CGRectMakeWithDictionaryRepresentation not loaded")
	}
	return _cGRectMakeWithDictionaryRepresentation(dict, rect)
}

var _cGRectOffset func(rect corefoundation.CGRect, dx float64, dy float64) corefoundation.CGRect

// CGRectOffset returns a rectangle with an origin that is offset from that of the source rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectOffset(_:_:_:)
func CGRectOffset(rect corefoundation.CGRect, dx float64, dy float64) corefoundation.CGRect {
	if _cGRectOffset == nil {
		panic("CoreGraphics: symbol CGRectOffset not loaded")
	}
	return _cGRectOffset(rect, dx, dy)
}

var _cGRectStandardize func(rect corefoundation.CGRect) corefoundation.CGRect

// CGRectStandardize returns a rectangle with a positive width and height.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectStandardize(_:)
func CGRectStandardize(rect corefoundation.CGRect) corefoundation.CGRect {
	if _cGRectStandardize == nil {
		panic("CoreGraphics: symbol CGRectStandardize not loaded")
	}
	return _cGRectStandardize(rect)
}

var _cGRectUnion func(r1 corefoundation.CGRect, r2 corefoundation.CGRect) corefoundation.CGRect

// CGRectUnion returns the smallest rectangle that contains the two source rectangles.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectUnion(_:_:)
func CGRectUnion(r1 corefoundation.CGRect, r2 corefoundation.CGRect) corefoundation.CGRect {
	if _cGRectUnion == nil {
		panic("CoreGraphics: symbol CGRectUnion not loaded")
	}
	return _cGRectUnion(r1, r2)
}

var _cGReleaseAllDisplays func() CGError

// CGReleaseAllDisplays releases all captured displays.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGReleaseAllDisplays()
func CGReleaseAllDisplays() CGError {
	if _cGReleaseAllDisplays == nil {
		panic("CoreGraphics: symbol CGReleaseAllDisplays not loaded")
	}
	return _cGReleaseAllDisplays()
}

var _cGReleaseDisplayFadeReservation func(token CGDisplayFadeReservationToken) CGError

// CGReleaseDisplayFadeReservation releases a display fade reservation, and unfades the display if needed.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGReleaseDisplayFadeReservation(_:)
func CGReleaseDisplayFadeReservation(token CGDisplayFadeReservationToken) CGError {
	if _cGReleaseDisplayFadeReservation == nil {
		panic("CoreGraphics: symbol CGReleaseDisplayFadeReservation not loaded")
	}
	return _cGReleaseDisplayFadeReservation(token)
}

var _cGRenderingBufferLockBytePtr func(provider CGRenderingBufferProviderRef) unsafe.Pointer

// CGRenderingBufferLockBytePtr.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRenderingBufferLockBytePtr
func CGRenderingBufferLockBytePtr(provider CGRenderingBufferProviderRef) unsafe.Pointer {
	if _cGRenderingBufferLockBytePtr == nil {
		panic("CoreGraphics: symbol CGRenderingBufferLockBytePtr not loaded")
	}
	return _cGRenderingBufferLockBytePtr(provider)
}

var _cGRenderingBufferProviderCreate func(info unsafe.Pointer, size uintptr) CGRenderingBufferProviderRef

// CGRenderingBufferProviderCreate.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRenderingBufferProviderCreate
func CGRenderingBufferProviderCreate(info unsafe.Pointer, size uintptr) CGRenderingBufferProviderRef {
	if _cGRenderingBufferProviderCreate == nil {
		panic("CoreGraphics: symbol CGRenderingBufferProviderCreate not loaded")
	}
	return _cGRenderingBufferProviderCreate(info, size)
}

var _cGRenderingBufferProviderCreateWithCFData func(data corefoundation.CFMutableDataRef) CGRenderingBufferProviderRef

// CGRenderingBufferProviderCreateWithCFData.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRenderingBufferProviderCreateWithCFData
func CGRenderingBufferProviderCreateWithCFData(data corefoundation.CFMutableDataRef) CGRenderingBufferProviderRef {
	if _cGRenderingBufferProviderCreateWithCFData == nil {
		panic("CoreGraphics: symbol CGRenderingBufferProviderCreateWithCFData not loaded")
	}
	return _cGRenderingBufferProviderCreateWithCFData(data)
}

var _cGRenderingBufferProviderGetSize func(provider CGRenderingBufferProviderRef) uintptr

// CGRenderingBufferProviderGetSize.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRenderingBufferProviderGetSize
func CGRenderingBufferProviderGetSize(provider CGRenderingBufferProviderRef) uintptr {
	if _cGRenderingBufferProviderGetSize == nil {
		panic("CoreGraphics: symbol CGRenderingBufferProviderGetSize not loaded")
	}
	return _cGRenderingBufferProviderGetSize(provider)
}

var _cGRenderingBufferProviderGetTypeID func() uint

// CGRenderingBufferProviderGetTypeID.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRenderingBufferProviderGetTypeID
func CGRenderingBufferProviderGetTypeID() uint {
	if _cGRenderingBufferProviderGetTypeID == nil {
		panic("CoreGraphics: symbol CGRenderingBufferProviderGetTypeID not loaded")
	}
	return _cGRenderingBufferProviderGetTypeID()
}

var _cGRenderingBufferUnlockBytePtr func(provider CGRenderingBufferProviderRef)

// CGRenderingBufferUnlockBytePtr.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRenderingBufferUnlockBytePtr
func CGRenderingBufferUnlockBytePtr(provider CGRenderingBufferProviderRef) {
	if _cGRenderingBufferUnlockBytePtr == nil {
		panic("CoreGraphics: symbol CGRenderingBufferUnlockBytePtr not loaded")
	}
	_cGRenderingBufferUnlockBytePtr(provider)
}

var _cGRequestListenEventAccess func() bool

// CGRequestListenEventAccess.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRequestListenEventAccess()
func CGRequestListenEventAccess() bool {
	if _cGRequestListenEventAccess == nil {
		panic("CoreGraphics: symbol CGRequestListenEventAccess not loaded")
	}
	return _cGRequestListenEventAccess()
}

var _cGRequestPostEventAccess func() bool

// CGRequestPostEventAccess.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRequestPostEventAccess()
func CGRequestPostEventAccess() bool {
	if _cGRequestPostEventAccess == nil {
		panic("CoreGraphics: symbol CGRequestPostEventAccess not loaded")
	}
	return _cGRequestPostEventAccess()
}

var _cGRequestScreenCaptureAccess func() bool

// CGRequestScreenCaptureAccess.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRequestScreenCaptureAccess()
func CGRequestScreenCaptureAccess() bool {
	if _cGRequestScreenCaptureAccess == nil {
		panic("CoreGraphics: symbol CGRequestScreenCaptureAccess not loaded")
	}
	return _cGRequestScreenCaptureAccess()
}

var _cGRestorePermanentDisplayConfiguration func()

// CGRestorePermanentDisplayConfiguration restores the permanent display configuration settings for the current user.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRestorePermanentDisplayConfiguration()
func CGRestorePermanentDisplayConfiguration() {
	if _cGRestorePermanentDisplayConfiguration == nil {
		panic("CoreGraphics: symbol CGRestorePermanentDisplayConfiguration not loaded")
	}
	_cGRestorePermanentDisplayConfiguration()
}

var _cGSessionCopyCurrentDictionary func() corefoundation.CFDictionaryRef

// CGSessionCopyCurrentDictionary returns information about the caller’s window server session.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGSessionCopyCurrentDictionary()
func CGSessionCopyCurrentDictionary() corefoundation.CFDictionaryRef {
	if _cGSessionCopyCurrentDictionary == nil {
		panic("CoreGraphics: symbol CGSessionCopyCurrentDictionary not loaded")
	}
	return _cGSessionCopyCurrentDictionary()
}

var _cGSetDisplayTransferByByteTable func(display uint32, tableSize uint32, redTable *uint8, greenTable *uint8, blueTable *uint8) CGError

// CGSetDisplayTransferByByteTable sets the byte values in the 8-bit RGB gamma tables for a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGSetDisplayTransferByByteTable(_:_:_:_:_:)
func CGSetDisplayTransferByByteTable(display uint32, tableSize uint32, redTable *uint8, greenTable *uint8, blueTable *uint8) CGError {
	if _cGSetDisplayTransferByByteTable == nil {
		panic("CoreGraphics: symbol CGSetDisplayTransferByByteTable not loaded")
	}
	return _cGSetDisplayTransferByByteTable(display, tableSize, redTable, greenTable, blueTable)
}

var _cGSetDisplayTransferByFormula func(display uint32, redMin CGGammaValue, redMax CGGammaValue, redGamma CGGammaValue, greenMin CGGammaValue, greenMax CGGammaValue, greenGamma CGGammaValue, blueMin CGGammaValue, blueMax CGGammaValue, blueGamma CGGammaValue) CGError

// CGSetDisplayTransferByFormula sets the gamma function for a display by specifying the coefficients of the gamma transfer formula.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGSetDisplayTransferByFormula(_:_:_:_:_:_:_:_:_:_:)
func CGSetDisplayTransferByFormula(display uint32, redMin CGGammaValue, redMax CGGammaValue, redGamma CGGammaValue, greenMin CGGammaValue, greenMax CGGammaValue, greenGamma CGGammaValue, blueMin CGGammaValue, blueMax CGGammaValue, blueGamma CGGammaValue) CGError {
	if _cGSetDisplayTransferByFormula == nil {
		panic("CoreGraphics: symbol CGSetDisplayTransferByFormula not loaded")
	}
	return _cGSetDisplayTransferByFormula(display, redMin, redMax, redGamma, greenMin, greenMax, greenGamma, blueMin, blueMax, blueGamma)
}

var _cGSetDisplayTransferByTable func(display uint32, tableSize uint32, redTable *CGGammaValue, greenTable *CGGammaValue, blueTable *CGGammaValue) CGError

// CGSetDisplayTransferByTable sets the color gamma function for a display by specifying the values in the RGB gamma tables.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGSetDisplayTransferByTable(_:_:_:_:_:)
func CGSetDisplayTransferByTable(display uint32, tableSize uint32, redTable *CGGammaValue, greenTable *CGGammaValue, blueTable *CGGammaValue) CGError {
	if _cGSetDisplayTransferByTable == nil {
		panic("CoreGraphics: symbol CGSetDisplayTransferByTable not loaded")
	}
	return _cGSetDisplayTransferByTable(display, tableSize, redTable, greenTable, blueTable)
}

var _cGShadingCreateAxial func(space CGColorSpaceRef, start corefoundation.CGPoint, end corefoundation.CGPoint, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef

// CGShadingCreateAxial creates a shading object to use for axial shading.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShading/init(axialSpace:start:end:function:extendStart:extendEnd:)
func CGShadingCreateAxial(space CGColorSpaceRef, start corefoundation.CGPoint, end corefoundation.CGPoint, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef {
	if _cGShadingCreateAxial == nil {
		panic("CoreGraphics: symbol CGShadingCreateAxial not loaded")
	}
	return _cGShadingCreateAxial(space, start, end, function, extendStart, extendEnd)
}

var _cGShadingCreateAxialWithContentHeadroom func(headroom float32, space CGColorSpaceRef, start corefoundation.CGPoint, end corefoundation.CGPoint, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef

// CGShadingCreateAxialWithContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShading/init(axialHeadroom:space:start:end:function:extendStart:extendEnd:)
func CGShadingCreateAxialWithContentHeadroom(headroom float32, space CGColorSpaceRef, start corefoundation.CGPoint, end corefoundation.CGPoint, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef {
	if _cGShadingCreateAxialWithContentHeadroom == nil {
		panic("CoreGraphics: symbol CGShadingCreateAxialWithContentHeadroom not loaded")
	}
	return _cGShadingCreateAxialWithContentHeadroom(headroom, space, start, end, function, extendStart, extendEnd)
}

var _cGShadingCreateRadial func(space CGColorSpaceRef, start corefoundation.CGPoint, startRadius float64, end corefoundation.CGPoint, endRadius float64, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef

// CGShadingCreateRadial creates a shading object to use for radial shading.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShading/init(radialSpace:start:startRadius:end:endRadius:function:extendStart:extendEnd:)
func CGShadingCreateRadial(space CGColorSpaceRef, start corefoundation.CGPoint, startRadius float64, end corefoundation.CGPoint, endRadius float64, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef {
	if _cGShadingCreateRadial == nil {
		panic("CoreGraphics: symbol CGShadingCreateRadial not loaded")
	}
	return _cGShadingCreateRadial(space, start, startRadius, end, endRadius, function, extendStart, extendEnd)
}

var _cGShadingCreateRadialWithContentHeadroom func(headroom float32, space CGColorSpaceRef, start corefoundation.CGPoint, startRadius float64, end corefoundation.CGPoint, endRadius float64, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef

// CGShadingCreateRadialWithContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShading/init(radialHeadroom:space:start:startRadius:end:endRadius:function:extendStart:extendEnd:)
func CGShadingCreateRadialWithContentHeadroom(headroom float32, space CGColorSpaceRef, start corefoundation.CGPoint, startRadius float64, end corefoundation.CGPoint, endRadius float64, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef {
	if _cGShadingCreateRadialWithContentHeadroom == nil {
		panic("CoreGraphics: symbol CGShadingCreateRadialWithContentHeadroom not loaded")
	}
	return _cGShadingCreateRadialWithContentHeadroom(headroom, space, start, startRadius, end, endRadius, function, extendStart, extendEnd)
}

var _cGShadingGetContentHeadroom func(shading CGShadingRef) float32

// CGShadingGetContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShading/contentHeadroom
func CGShadingGetContentHeadroom(shading CGShadingRef) float32 {
	if _cGShadingGetContentHeadroom == nil {
		panic("CoreGraphics: symbol CGShadingGetContentHeadroom not loaded")
	}
	return _cGShadingGetContentHeadroom(shading)
}

var _cGShadingGetTypeID func() uint

// CGShadingGetTypeID returns the Core Foundation type identifier for Core Graphics shading objects.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShading/typeID
func CGShadingGetTypeID() uint {
	if _cGShadingGetTypeID == nil {
		panic("CoreGraphics: symbol CGShadingGetTypeID not loaded")
	}
	return _cGShadingGetTypeID()
}

var _cGShadingRelease func(shading CGShadingRef)

// CGShadingRelease decrements the retain count of a shading object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShadingRelease
func CGShadingRelease(shading CGShadingRef) {
	if _cGShadingRelease == nil {
		panic("CoreGraphics: symbol CGShadingRelease not loaded")
	}
	_cGShadingRelease(shading)
}

var _cGShadingRetain func(shading CGShadingRef) CGShadingRef

// CGShadingRetain increments the retain count of a shading object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShadingRetain
func CGShadingRetain(shading CGShadingRef) CGShadingRef {
	if _cGShadingRetain == nil {
		panic("CoreGraphics: symbol CGShadingRetain not loaded")
	}
	return _cGShadingRetain(shading)
}

var _cGShieldingWindowID func(display uint32) CGWindowID

// CGShieldingWindowID returns the window ID of the shield window for a captured display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShieldingWindowID(_:)
func CGShieldingWindowID(display uint32) CGWindowID {
	if _cGShieldingWindowID == nil {
		panic("CoreGraphics: symbol CGShieldingWindowID not loaded")
	}
	return _cGShieldingWindowID(display)
}

var _cGShieldingWindowLevel func() CGWindowLevel

// CGShieldingWindowLevel returns the window level of the shield window for a captured display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShieldingWindowLevel()
func CGShieldingWindowLevel() CGWindowLevel {
	if _cGShieldingWindowLevel == nil {
		panic("CoreGraphics: symbol CGShieldingWindowLevel not loaded")
	}
	return _cGShieldingWindowLevel()
}

var _cGSizeApplyAffineTransform func(size corefoundation.CGSize, t corefoundation.CGAffineTransform) corefoundation.CGSize

// CGSizeApplyAffineTransform returns the height and width resulting from a transformation of an existing height and width.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGSizeApplyAffineTransform(_:_:)
func CGSizeApplyAffineTransform(size corefoundation.CGSize, t corefoundation.CGAffineTransform) corefoundation.CGSize {
	if _cGSizeApplyAffineTransform == nil {
		panic("CoreGraphics: symbol CGSizeApplyAffineTransform not loaded")
	}
	return _cGSizeApplyAffineTransform(size, t)
}

var _cGSizeCreateDictionaryRepresentation func(size corefoundation.CGSize) corefoundation.CFDictionaryRef

// CGSizeCreateDictionaryRepresentation returns a dictionary representation of the specified size.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGSizeCreateDictionaryRepresentation(_:)
func CGSizeCreateDictionaryRepresentation(size corefoundation.CGSize) corefoundation.CFDictionaryRef {
	if _cGSizeCreateDictionaryRepresentation == nil {
		panic("CoreGraphics: symbol CGSizeCreateDictionaryRepresentation not loaded")
	}
	return _cGSizeCreateDictionaryRepresentation(size)
}

var _cGSizeEqualToSize func(size1 corefoundation.CGSize, size2 corefoundation.CGSize) bool

// CGSizeEqualToSize returns whether two sizes are equal.
//
// Deprecated: The [CGSize](<doc://com.apple.documentation/documentation/CoreFoundation/CGSize>) type adopts the [Equatable] protocol; use the `==` operator instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGSizeEqualToSize(_:_:)
func CGSizeEqualToSize(size1 corefoundation.CGSize, size2 corefoundation.CGSize) bool {
	if _cGSizeEqualToSize == nil {
		panic("CoreGraphics: symbol CGSizeEqualToSize not loaded")
	}
	return _cGSizeEqualToSize(size1, size2)
}

var _cGSizeMakeWithDictionaryRepresentation func(dict corefoundation.CFDictionaryRef, size *corefoundation.CGSize) bool

// CGSizeMakeWithDictionaryRepresentation fills in a size using the contents of the specified dictionary.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGSizeMakeWithDictionaryRepresentation(_:_:)
func CGSizeMakeWithDictionaryRepresentation(dict corefoundation.CFDictionaryRef, size *corefoundation.CGSize) bool {
	if _cGSizeMakeWithDictionaryRepresentation == nil {
		panic("CoreGraphics: symbol CGSizeMakeWithDictionaryRepresentation not loaded")
	}
	return _cGSizeMakeWithDictionaryRepresentation(dict, size)
}

var _cGWarpMouseCursorPosition func(newCursorPosition corefoundation.CGPoint) CGError

// CGWarpMouseCursorPosition moves the mouse cursor without generating events.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGWarpMouseCursorPosition(_:)
func CGWarpMouseCursorPosition(newCursorPosition corefoundation.CGPoint) CGError {
	if _cGWarpMouseCursorPosition == nil {
		panic("CoreGraphics: symbol CGWarpMouseCursorPosition not loaded")
	}
	return _cGWarpMouseCursorPosition(newCursorPosition)
}

var _cGWindowLevelForKey func(key CGWindowLevelKey) CGWindowLevel

// CGWindowLevelForKey returns the window level that corresponds to one of the standard window types.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelForKey(_:)
func CGWindowLevelForKey(key CGWindowLevelKey) CGWindowLevel {
	if _cGWindowLevelForKey == nil {
		panic("CoreGraphics: symbol CGWindowLevelForKey not loaded")
	}
	return _cGWindowLevelForKey(key)
}

var _cGWindowListCopyWindowInfo func(option CGWindowListOption, relativeToWindow CGWindowID) corefoundation.CFArrayRef

// CGWindowListCopyWindowInfo generates and returns information about the selected windows in the current user session.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGWindowListCopyWindowInfo(_:_:)
func CGWindowListCopyWindowInfo(option CGWindowListOption, relativeToWindow CGWindowID) corefoundation.CFArrayRef {
	if _cGWindowListCopyWindowInfo == nil {
		panic("CoreGraphics: symbol CGWindowListCopyWindowInfo not loaded")
	}
	return _cGWindowListCopyWindowInfo(option, relativeToWindow)
}

var _cGWindowListCreate func(option CGWindowListOption, relativeToWindow CGWindowID) corefoundation.CFArrayRef

// CGWindowListCreate returns the list of window IDs associated with the specified windows in the current user session.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGWindowListCreate
func CGWindowListCreate(option CGWindowListOption, relativeToWindow CGWindowID) corefoundation.CFArrayRef {
	if _cGWindowListCreate == nil {
		panic("CoreGraphics: symbol CGWindowListCreate not loaded")
	}
	return _cGWindowListCreate(option, relativeToWindow)
}

var _cGWindowListCreateDescriptionFromArray func(windowArray corefoundation.CFArrayRef) corefoundation.CFArrayRef

// CGWindowListCreateDescriptionFromArray generates and returns information about windows with the specified window IDs.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGWindowListCreateDescriptionFromArray(_:)
func CGWindowListCreateDescriptionFromArray(windowArray corefoundation.CFArrayRef) corefoundation.CFArrayRef {
	if _cGWindowListCreateDescriptionFromArray == nil {
		panic("CoreGraphics: symbol CGWindowListCreateDescriptionFromArray not loaded")
	}
	return _cGWindowListCreateDescriptionFromArray(windowArray)
}

var _cGWindowListCreateImage func(screenBounds corefoundation.CGRect, listOption CGWindowListOption, windowID CGWindowID, imageOption CGWindowImageOption) CGImageRef

// CGWindowListCreateImage returns a composite image based on a dynamically generated list of windows.
//
// Deprecated: Please use ScreenCaptureKit instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGWindowListCreateImage(_:_:_:_:)
func CGWindowListCreateImage(screenBounds corefoundation.CGRect, listOption CGWindowListOption, windowID CGWindowID, imageOption CGWindowImageOption) CGImageRef {
	if _cGWindowListCreateImage == nil {
		panic("CoreGraphics: symbol CGWindowListCreateImage not loaded")
	}
	return _cGWindowListCreateImage(screenBounds, listOption, windowID, imageOption)
}

var _cGWindowServerCreateServerPort func() corefoundation.CFMachPort

// CGWindowServerCreateServerPort.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGWindowServerCreateServerPort()
func CGWindowServerCreateServerPort() corefoundation.CFMachPort {
	if _cGWindowServerCreateServerPort == nil {
		panic("CoreGraphics: symbol CGWindowServerCreateServerPort not loaded")
	}
	return _cGWindowServerCreateServerPort()
}

func init() {
	if frameworkHandle == 0 {
		return
	}
		registerFunc(&_cGAcquireDisplayFadeReservation, frameworkHandle, "CGAcquireDisplayFadeReservation")
		registerFunc(&_cGAffineTransformConcat, frameworkHandle, "CGAffineTransformConcat")
		registerFunc(&_cGAffineTransformDecompose, frameworkHandle, "CGAffineTransformDecompose")
		registerFunc(&_cGAffineTransformEqualToTransform, frameworkHandle, "CGAffineTransformEqualToTransform")
		registerFunc(&_cGAffineTransformInvert, frameworkHandle, "CGAffineTransformInvert")
		registerFunc(&_cGAffineTransformIsIdentity, frameworkHandle, "CGAffineTransformIsIdentity")
		registerFunc(&_cGAffineTransformMake, frameworkHandle, "CGAffineTransformMake")
		registerFunc(&_cGAffineTransformMakeRotation, frameworkHandle, "CGAffineTransformMakeRotation")
		registerFunc(&_cGAffineTransformMakeScale, frameworkHandle, "CGAffineTransformMakeScale")
		registerFunc(&_cGAffineTransformMakeTranslation, frameworkHandle, "CGAffineTransformMakeTranslation")
		registerFunc(&_cGAffineTransformMakeWithComponents, frameworkHandle, "CGAffineTransformMakeWithComponents")
		registerFunc(&_cGAffineTransformRotate, frameworkHandle, "CGAffineTransformRotate")
		registerFunc(&_cGAffineTransformScale, frameworkHandle, "CGAffineTransformScale")
		registerFunc(&_cGAffineTransformTranslate, frameworkHandle, "CGAffineTransformTranslate")
		registerFunc(&_cGAssociateMouseAndMouseCursorPosition, frameworkHandle, "CGAssociateMouseAndMouseCursorPosition")
		registerFunc(&_cGBeginDisplayConfiguration, frameworkHandle, "CGBeginDisplayConfiguration")
		registerFunc(&_cGBitmapContextCreate, frameworkHandle, "CGBitmapContextCreate")
		registerFunc(&_cGBitmapContextCreateAdaptive, frameworkHandle, "CGBitmapContextCreateAdaptive")
		registerFunc(&_cGBitmapContextCreateImage, frameworkHandle, "CGBitmapContextCreateImage")
		registerFunc(&_cGBitmapContextCreateWithData, frameworkHandle, "CGBitmapContextCreateWithData")
		registerFunc(&_cGBitmapContextGetAlphaInfo, frameworkHandle, "CGBitmapContextGetAlphaInfo")
		registerFunc(&_cGBitmapContextGetBitmapInfo, frameworkHandle, "CGBitmapContextGetBitmapInfo")
		registerFunc(&_cGBitmapContextGetBitsPerComponent, frameworkHandle, "CGBitmapContextGetBitsPerComponent")
		registerFunc(&_cGBitmapContextGetBitsPerPixel, frameworkHandle, "CGBitmapContextGetBitsPerPixel")
		registerFunc(&_cGBitmapContextGetBytesPerRow, frameworkHandle, "CGBitmapContextGetBytesPerRow")
		registerFunc(&_cGBitmapContextGetColorSpace, frameworkHandle, "CGBitmapContextGetColorSpace")
		registerFunc(&_cGBitmapContextGetData, frameworkHandle, "CGBitmapContextGetData")
		registerFunc(&_cGBitmapContextGetHeight, frameworkHandle, "CGBitmapContextGetHeight")
		registerFunc(&_cGBitmapContextGetWidth, frameworkHandle, "CGBitmapContextGetWidth")
		registerFunc(&_cGCancelDisplayConfiguration, frameworkHandle, "CGCancelDisplayConfiguration")
		registerFunc(&_cGCaptureAllDisplays, frameworkHandle, "CGCaptureAllDisplays")
		registerFunc(&_cGCaptureAllDisplaysWithOptions, frameworkHandle, "CGCaptureAllDisplaysWithOptions")
		registerFunc(&_cGColorConversionInfoConvertData, frameworkHandle, "CGColorConversionInfoConvertData")
		registerFunc(&_cGColorConversionInfoCreate, frameworkHandle, "CGColorConversionInfoCreate")
		registerFunc(&_cGColorConversionInfoCreateForToneMapping, frameworkHandle, "CGColorConversionInfoCreateForToneMapping")
		registerFunc(&_cGColorConversionInfoCreateFromList, frameworkHandle, "CGColorConversionInfoCreateFromList")
		registerFunc(&_cGColorConversionInfoCreateFromListWithArguments, frameworkHandle, "CGColorConversionInfoCreateFromListWithArguments")
		registerFunc(&_cGColorConversionInfoCreateWithOptions, frameworkHandle, "CGColorConversionInfoCreateWithOptions")
		registerFunc(&_cGColorConversionInfoGetTypeID, frameworkHandle, "CGColorConversionInfoGetTypeID")
		registerFunc(&_cGColorCreate, frameworkHandle, "CGColorCreate")
		registerFunc(&_cGColorCreateCopy, frameworkHandle, "CGColorCreateCopy")
		registerFunc(&_cGColorCreateCopyByMatchingToColorSpace, frameworkHandle, "CGColorCreateCopyByMatchingToColorSpace")
		registerFunc(&_cGColorCreateCopyWithAlpha, frameworkHandle, "CGColorCreateCopyWithAlpha")
		registerFunc(&_cGColorCreateGenericCMYK, frameworkHandle, "CGColorCreateGenericCMYK")
		registerFunc(&_cGColorCreateGenericGray, frameworkHandle, "CGColorCreateGenericGray")
		registerFunc(&_cGColorCreateGenericGrayGamma2_2, frameworkHandle, "CGColorCreateGenericGrayGamma2_2")
		registerFunc(&_cGColorCreateGenericRGB, frameworkHandle, "CGColorCreateGenericRGB")
		registerFunc(&_cGColorCreateSRGB, frameworkHandle, "CGColorCreateSRGB")
		registerFunc(&_cGColorCreateWithContentHeadroom, frameworkHandle, "CGColorCreateWithContentHeadroom")
		registerFunc(&_cGColorCreateWithPattern, frameworkHandle, "CGColorCreateWithPattern")
		registerFunc(&_cGColorEqualToColor, frameworkHandle, "CGColorEqualToColor")
		registerFunc(&_cGColorGetAlpha, frameworkHandle, "CGColorGetAlpha")
		registerFunc(&_cGColorGetColorSpace, frameworkHandle, "CGColorGetColorSpace")
		registerFunc(&_cGColorGetComponents, frameworkHandle, "CGColorGetComponents")
		registerFunc(&_cGColorGetConstantColor, frameworkHandle, "CGColorGetConstantColor")
		registerFunc(&_cGColorGetContentHeadroom, frameworkHandle, "CGColorGetContentHeadroom")
		registerFunc(&_cGColorGetNumberOfComponents, frameworkHandle, "CGColorGetNumberOfComponents")
		registerFunc(&_cGColorGetPattern, frameworkHandle, "CGColorGetPattern")
		registerFunc(&_cGColorGetTypeID, frameworkHandle, "CGColorGetTypeID")
		registerFunc(&_cGColorRelease, frameworkHandle, "CGColorRelease")
		registerFunc(&_cGColorRetain, frameworkHandle, "CGColorRetain")
		registerFunc(&_cGColorSpaceCopyBaseColorSpace, frameworkHandle, "CGColorSpaceCopyBaseColorSpace")
		registerFunc(&_cGColorSpaceCopyICCData, frameworkHandle, "CGColorSpaceCopyICCData")
		registerFunc(&_cGColorSpaceCopyName, frameworkHandle, "CGColorSpaceCopyName")
		registerFunc(&_cGColorSpaceCopyPropertyList, frameworkHandle, "CGColorSpaceCopyPropertyList")
		registerFunc(&_cGColorSpaceCreateCalibratedGray, frameworkHandle, "CGColorSpaceCreateCalibratedGray")
		registerFunc(&_cGColorSpaceCreateCalibratedRGB, frameworkHandle, "CGColorSpaceCreateCalibratedRGB")
		registerFunc(&_cGColorSpaceCreateCopyWithStandardRange, frameworkHandle, "CGColorSpaceCreateCopyWithStandardRange")
		registerFunc(&_cGColorSpaceCreateDeviceCMYK, frameworkHandle, "CGColorSpaceCreateDeviceCMYK")
		registerFunc(&_cGColorSpaceCreateDeviceGray, frameworkHandle, "CGColorSpaceCreateDeviceGray")
		registerFunc(&_cGColorSpaceCreateDeviceRGB, frameworkHandle, "CGColorSpaceCreateDeviceRGB")
		registerFunc(&_cGColorSpaceCreateExtended, frameworkHandle, "CGColorSpaceCreateExtended")
		registerFunc(&_cGColorSpaceCreateExtendedLinearized, frameworkHandle, "CGColorSpaceCreateExtendedLinearized")
		registerFunc(&_cGColorSpaceCreateICCBased, frameworkHandle, "CGColorSpaceCreateICCBased")
		registerFunc(&_cGColorSpaceCreateIndexed, frameworkHandle, "CGColorSpaceCreateIndexed")
		registerFunc(&_cGColorSpaceCreateLab, frameworkHandle, "CGColorSpaceCreateLab")
		registerFunc(&_cGColorSpaceCreateLinearized, frameworkHandle, "CGColorSpaceCreateLinearized")
		registerFunc(&_cGColorSpaceCreatePattern, frameworkHandle, "CGColorSpaceCreatePattern")
		registerFunc(&_cGColorSpaceCreateWithColorSyncProfile, frameworkHandle, "CGColorSpaceCreateWithColorSyncProfile")
		registerFunc(&_cGColorSpaceCreateWithICCData, frameworkHandle, "CGColorSpaceCreateWithICCData")
		registerFunc(&_cGColorSpaceCreateWithName, frameworkHandle, "CGColorSpaceCreateWithName")
		registerFunc(&_cGColorSpaceCreateWithPropertyList, frameworkHandle, "CGColorSpaceCreateWithPropertyList")
		registerFunc(&_cGColorSpaceGetBaseColorSpace, frameworkHandle, "CGColorSpaceGetBaseColorSpace")
		registerFunc(&_cGColorSpaceGetColorTable, frameworkHandle, "CGColorSpaceGetColorTable")
		registerFunc(&_cGColorSpaceGetColorTableCount, frameworkHandle, "CGColorSpaceGetColorTableCount")
		registerFunc(&_cGColorSpaceGetModel, frameworkHandle, "CGColorSpaceGetModel")
		registerFunc(&_cGColorSpaceGetName, frameworkHandle, "CGColorSpaceGetName")
		registerFunc(&_cGColorSpaceGetNumberOfComponents, frameworkHandle, "CGColorSpaceGetNumberOfComponents")
		registerFunc(&_cGColorSpaceGetTypeID, frameworkHandle, "CGColorSpaceGetTypeID")
		registerFunc(&_cGColorSpaceIsHDR, frameworkHandle, "CGColorSpaceIsHDR")
		registerFunc(&_cGColorSpaceIsHLGBased, frameworkHandle, "CGColorSpaceIsHLGBased")
		registerFunc(&_cGColorSpaceIsPQBased, frameworkHandle, "CGColorSpaceIsPQBased")
		registerFunc(&_cGColorSpaceIsWideGamutRGB, frameworkHandle, "CGColorSpaceIsWideGamutRGB")
		registerFunc(&_cGColorSpaceRelease, frameworkHandle, "CGColorSpaceRelease")
		registerFunc(&_cGColorSpaceRetain, frameworkHandle, "CGColorSpaceRetain")
		registerFunc(&_cGColorSpaceSupportsOutput, frameworkHandle, "CGColorSpaceSupportsOutput")
		registerFunc(&_cGColorSpaceUsesExtendedRange, frameworkHandle, "CGColorSpaceUsesExtendedRange")
		registerFunc(&_cGColorSpaceUsesITUR_2100TF, frameworkHandle, "CGColorSpaceUsesITUR_2100TF")
		registerFunc(&_cGCompleteDisplayConfiguration, frameworkHandle, "CGCompleteDisplayConfiguration")
		registerFunc(&_cGConfigureDisplayFadeEffect, frameworkHandle, "CGConfigureDisplayFadeEffect")
		registerFunc(&_cGConfigureDisplayMirrorOfDisplay, frameworkHandle, "CGConfigureDisplayMirrorOfDisplay")
		registerFunc(&_cGConfigureDisplayOrigin, frameworkHandle, "CGConfigureDisplayOrigin")
		registerFunc(&_cGConfigureDisplayStereoOperation, frameworkHandle, "CGConfigureDisplayStereoOperation")
		registerFunc(&_cGConfigureDisplayWithDisplayMode, frameworkHandle, "CGConfigureDisplayWithDisplayMode")
		registerFunc(&_cGContextAddArc, frameworkHandle, "CGContextAddArc")
		registerFunc(&_cGContextAddArcToPoint, frameworkHandle, "CGContextAddArcToPoint")
		registerFunc(&_cGContextAddCurveToPoint, frameworkHandle, "CGContextAddCurveToPoint")
		registerFunc(&_cGContextAddEllipseInRect, frameworkHandle, "CGContextAddEllipseInRect")
		registerFunc(&_cGContextAddLineToPoint, frameworkHandle, "CGContextAddLineToPoint")
		registerFunc(&_cGContextAddLines, frameworkHandle, "CGContextAddLines")
		registerFunc(&_cGContextAddPath, frameworkHandle, "CGContextAddPath")
		registerFunc(&_cGContextAddQuadCurveToPoint, frameworkHandle, "CGContextAddQuadCurveToPoint")
		registerFunc(&_cGContextAddRect, frameworkHandle, "CGContextAddRect")
		registerFunc(&_cGContextAddRects, frameworkHandle, "CGContextAddRects")
		registerFunc(&_cGContextBeginPage, frameworkHandle, "CGContextBeginPage")
		registerFunc(&_cGContextBeginPath, frameworkHandle, "CGContextBeginPath")
		registerFunc(&_cGContextBeginTransparencyLayer, frameworkHandle, "CGContextBeginTransparencyLayer")
		registerFunc(&_cGContextBeginTransparencyLayerWithRect, frameworkHandle, "CGContextBeginTransparencyLayerWithRect")
		registerFunc(&_cGContextClearRect, frameworkHandle, "CGContextClearRect")
		registerFunc(&_cGContextClip, frameworkHandle, "CGContextClip")
		registerFunc(&_cGContextClipToMask, frameworkHandle, "CGContextClipToMask")
		registerFunc(&_cGContextClipToRect, frameworkHandle, "CGContextClipToRect")
		registerFunc(&_cGContextClipToRects, frameworkHandle, "CGContextClipToRects")
		registerFunc(&_cGContextClosePath, frameworkHandle, "CGContextClosePath")
		registerFunc(&_cGContextConcatCTM, frameworkHandle, "CGContextConcatCTM")
		registerFunc(&_cGContextConvertPointToDeviceSpace, frameworkHandle, "CGContextConvertPointToDeviceSpace")
		registerFunc(&_cGContextConvertPointToUserSpace, frameworkHandle, "CGContextConvertPointToUserSpace")
		registerFunc(&_cGContextConvertRectToDeviceSpace, frameworkHandle, "CGContextConvertRectToDeviceSpace")
		registerFunc(&_cGContextConvertRectToUserSpace, frameworkHandle, "CGContextConvertRectToUserSpace")
		registerFunc(&_cGContextConvertSizeToDeviceSpace, frameworkHandle, "CGContextConvertSizeToDeviceSpace")
		registerFunc(&_cGContextConvertSizeToUserSpace, frameworkHandle, "CGContextConvertSizeToUserSpace")
		registerFunc(&_cGContextCopyPath, frameworkHandle, "CGContextCopyPath")
		registerFunc(&_cGContextDrawConicGradient, frameworkHandle, "CGContextDrawConicGradient")
		registerFunc(&_cGContextDrawImage, frameworkHandle, "CGContextDrawImage")
		registerFunc(&_cGContextDrawImageApplyingToneMapping, frameworkHandle, "CGContextDrawImageApplyingToneMapping")
		registerFunc(&_cGContextDrawLayerAtPoint, frameworkHandle, "CGContextDrawLayerAtPoint")
		registerFunc(&_cGContextDrawLayerInRect, frameworkHandle, "CGContextDrawLayerInRect")
		registerFunc(&_cGContextDrawLinearGradient, frameworkHandle, "CGContextDrawLinearGradient")
		registerFunc(&_cGContextDrawPDFDocument, frameworkHandle, "CGContextDrawPDFDocument")
		registerFunc(&_cGContextDrawPDFPage, frameworkHandle, "CGContextDrawPDFPage")
		registerFunc(&_cGContextDrawPath, frameworkHandle, "CGContextDrawPath")
		registerFunc(&_cGContextDrawRadialGradient, frameworkHandle, "CGContextDrawRadialGradient")
		registerFunc(&_cGContextDrawShading, frameworkHandle, "CGContextDrawShading")
		registerFunc(&_cGContextDrawTiledImage, frameworkHandle, "CGContextDrawTiledImage")
		registerFunc(&_cGContextEOClip, frameworkHandle, "CGContextEOClip")
		registerFunc(&_cGContextEOFillPath, frameworkHandle, "CGContextEOFillPath")
		registerFunc(&_cGContextEndPage, frameworkHandle, "CGContextEndPage")
		registerFunc(&_cGContextEndTransparencyLayer, frameworkHandle, "CGContextEndTransparencyLayer")
		registerFunc(&_cGContextFillEllipseInRect, frameworkHandle, "CGContextFillEllipseInRect")
		registerFunc(&_cGContextFillPath, frameworkHandle, "CGContextFillPath")
		registerFunc(&_cGContextFillRect, frameworkHandle, "CGContextFillRect")
		registerFunc(&_cGContextFillRects, frameworkHandle, "CGContextFillRects")
		registerFunc(&_cGContextFlush, frameworkHandle, "CGContextFlush")
		registerFunc(&_cGContextGetCTM, frameworkHandle, "CGContextGetCTM")
		registerFunc(&_cGContextGetClipBoundingBox, frameworkHandle, "CGContextGetClipBoundingBox")
		registerFunc(&_cGContextGetContentToneMappingInfo, frameworkHandle, "CGContextGetContentToneMappingInfo")
		registerFunc(&_cGContextGetEDRTargetHeadroom, frameworkHandle, "CGContextGetEDRTargetHeadroom")
		registerFunc(&_cGContextGetInterpolationQuality, frameworkHandle, "CGContextGetInterpolationQuality")
		registerFunc(&_cGContextGetPathBoundingBox, frameworkHandle, "CGContextGetPathBoundingBox")
		registerFunc(&_cGContextGetPathCurrentPoint, frameworkHandle, "CGContextGetPathCurrentPoint")
		registerFunc(&_cGContextGetTextMatrix, frameworkHandle, "CGContextGetTextMatrix")
		registerFunc(&_cGContextGetTextPosition, frameworkHandle, "CGContextGetTextPosition")
		registerFunc(&_cGContextGetTypeID, frameworkHandle, "CGContextGetTypeID")
		registerFunc(&_cGContextGetUserSpaceToDeviceSpaceTransform, frameworkHandle, "CGContextGetUserSpaceToDeviceSpaceTransform")
		registerFunc(&_cGContextIsPathEmpty, frameworkHandle, "CGContextIsPathEmpty")
		registerFunc(&_cGContextMoveToPoint, frameworkHandle, "CGContextMoveToPoint")
		registerFunc(&_cGContextPathContainsPoint, frameworkHandle, "CGContextPathContainsPoint")
		registerFunc(&_cGContextRelease, frameworkHandle, "CGContextRelease")
		registerFunc(&_cGContextReplacePathWithStrokedPath, frameworkHandle, "CGContextReplacePathWithStrokedPath")
		registerFunc(&_cGContextResetClip, frameworkHandle, "CGContextResetClip")
		registerFunc(&_cGContextRestoreGState, frameworkHandle, "CGContextRestoreGState")
		registerFunc(&_cGContextRetain, frameworkHandle, "CGContextRetain")
		registerFunc(&_cGContextRotateCTM, frameworkHandle, "CGContextRotateCTM")
		registerFunc(&_cGContextSaveGState, frameworkHandle, "CGContextSaveGState")
		registerFunc(&_cGContextScaleCTM, frameworkHandle, "CGContextScaleCTM")
		registerFunc(&_cGContextSetAllowsAntialiasing, frameworkHandle, "CGContextSetAllowsAntialiasing")
		registerFunc(&_cGContextSetAllowsFontSmoothing, frameworkHandle, "CGContextSetAllowsFontSmoothing")
		registerFunc(&_cGContextSetAllowsFontSubpixelPositioning, frameworkHandle, "CGContextSetAllowsFontSubpixelPositioning")
		registerFunc(&_cGContextSetAllowsFontSubpixelQuantization, frameworkHandle, "CGContextSetAllowsFontSubpixelQuantization")
		registerFunc(&_cGContextSetAlpha, frameworkHandle, "CGContextSetAlpha")
		registerFunc(&_cGContextSetBlendMode, frameworkHandle, "CGContextSetBlendMode")
		registerFunc(&_cGContextSetCMYKFillColor, frameworkHandle, "CGContextSetCMYKFillColor")
		registerFunc(&_cGContextSetCMYKStrokeColor, frameworkHandle, "CGContextSetCMYKStrokeColor")
		registerFunc(&_cGContextSetCharacterSpacing, frameworkHandle, "CGContextSetCharacterSpacing")
		registerFunc(&_cGContextSetContentToneMappingInfo, frameworkHandle, "CGContextSetContentToneMappingInfo")
		registerFunc(&_cGContextSetEDRTargetHeadroom, frameworkHandle, "CGContextSetEDRTargetHeadroom")
		registerFunc(&_cGContextSetFillColor, frameworkHandle, "CGContextSetFillColor")
		registerFunc(&_cGContextSetFillColorSpace, frameworkHandle, "CGContextSetFillColorSpace")
		registerFunc(&_cGContextSetFillColorWithColor, frameworkHandle, "CGContextSetFillColorWithColor")
		registerFunc(&_cGContextSetFillPattern, frameworkHandle, "CGContextSetFillPattern")
		registerFunc(&_cGContextSetFlatness, frameworkHandle, "CGContextSetFlatness")
		registerFunc(&_cGContextSetFont, frameworkHandle, "CGContextSetFont")
		registerFunc(&_cGContextSetFontSize, frameworkHandle, "CGContextSetFontSize")
		registerFunc(&_cGContextSetGrayFillColor, frameworkHandle, "CGContextSetGrayFillColor")
		registerFunc(&_cGContextSetGrayStrokeColor, frameworkHandle, "CGContextSetGrayStrokeColor")
		registerFunc(&_cGContextSetInterpolationQuality, frameworkHandle, "CGContextSetInterpolationQuality")
		registerFunc(&_cGContextSetLineCap, frameworkHandle, "CGContextSetLineCap")
		registerFunc(&_cGContextSetLineDash, frameworkHandle, "CGContextSetLineDash")
		registerFunc(&_cGContextSetLineJoin, frameworkHandle, "CGContextSetLineJoin")
		registerFunc(&_cGContextSetLineWidth, frameworkHandle, "CGContextSetLineWidth")
		registerFunc(&_cGContextSetMiterLimit, frameworkHandle, "CGContextSetMiterLimit")
		registerFunc(&_cGContextSetPatternPhase, frameworkHandle, "CGContextSetPatternPhase")
		registerFunc(&_cGContextSetRGBFillColor, frameworkHandle, "CGContextSetRGBFillColor")
		registerFunc(&_cGContextSetRGBStrokeColor, frameworkHandle, "CGContextSetRGBStrokeColor")
		registerFunc(&_cGContextSetRenderingIntent, frameworkHandle, "CGContextSetRenderingIntent")
		registerFunc(&_cGContextSetShadow, frameworkHandle, "CGContextSetShadow")
		registerFunc(&_cGContextSetShadowWithColor, frameworkHandle, "CGContextSetShadowWithColor")
		registerFunc(&_cGContextSetShouldAntialias, frameworkHandle, "CGContextSetShouldAntialias")
		registerFunc(&_cGContextSetShouldSmoothFonts, frameworkHandle, "CGContextSetShouldSmoothFonts")
		registerFunc(&_cGContextSetShouldSubpixelPositionFonts, frameworkHandle, "CGContextSetShouldSubpixelPositionFonts")
		registerFunc(&_cGContextSetShouldSubpixelQuantizeFonts, frameworkHandle, "CGContextSetShouldSubpixelQuantizeFonts")
		registerFunc(&_cGContextSetStrokeColor, frameworkHandle, "CGContextSetStrokeColor")
		registerFunc(&_cGContextSetStrokeColorSpace, frameworkHandle, "CGContextSetStrokeColorSpace")
		registerFunc(&_cGContextSetStrokeColorWithColor, frameworkHandle, "CGContextSetStrokeColorWithColor")
		registerFunc(&_cGContextSetStrokePattern, frameworkHandle, "CGContextSetStrokePattern")
		registerFunc(&_cGContextSetTextDrawingMode, frameworkHandle, "CGContextSetTextDrawingMode")
		registerFunc(&_cGContextSetTextMatrix, frameworkHandle, "CGContextSetTextMatrix")
		registerFunc(&_cGContextSetTextPosition, frameworkHandle, "CGContextSetTextPosition")
		registerFunc(&_cGContextShowGlyphsAtPositions, frameworkHandle, "CGContextShowGlyphsAtPositions")
		registerFunc(&_cGContextStrokeEllipseInRect, frameworkHandle, "CGContextStrokeEllipseInRect")
		registerFunc(&_cGContextStrokeLineSegments, frameworkHandle, "CGContextStrokeLineSegments")
		registerFunc(&_cGContextStrokePath, frameworkHandle, "CGContextStrokePath")
		registerFunc(&_cGContextStrokeRect, frameworkHandle, "CGContextStrokeRect")
		registerFunc(&_cGContextStrokeRectWithWidth, frameworkHandle, "CGContextStrokeRectWithWidth")
		registerFunc(&_cGContextSynchronize, frameworkHandle, "CGContextSynchronize")
		registerFunc(&_cGContextSynchronizeAttributes, frameworkHandle, "CGContextSynchronizeAttributes")
		registerFunc(&_cGContextTranslateCTM, frameworkHandle, "CGContextTranslateCTM")
		registerFunc(&_cGConvertColorDataWithFormat, frameworkHandle, "CGConvertColorDataWithFormat")
		registerFunc(&_cGDataConsumerCreate, frameworkHandle, "CGDataConsumerCreate")
		registerFunc(&_cGDataConsumerCreateWithCFData, frameworkHandle, "CGDataConsumerCreateWithCFData")
		registerFunc(&_cGDataConsumerCreateWithURL, frameworkHandle, "CGDataConsumerCreateWithURL")
		registerFunc(&_cGDataConsumerGetTypeID, frameworkHandle, "CGDataConsumerGetTypeID")
		registerFunc(&_cGDataConsumerRelease, frameworkHandle, "CGDataConsumerRelease")
		registerFunc(&_cGDataConsumerRetain, frameworkHandle, "CGDataConsumerRetain")
		registerFunc(&_cGDataProviderCopyData, frameworkHandle, "CGDataProviderCopyData")
		registerFunc(&_cGDataProviderCreateDirect, frameworkHandle, "CGDataProviderCreateDirect")
		registerFunc(&_cGDataProviderCreateSequential, frameworkHandle, "CGDataProviderCreateSequential")
		registerFunc(&_cGDataProviderCreateWithCFData, frameworkHandle, "CGDataProviderCreateWithCFData")
		registerFunc(&_cGDataProviderCreateWithData, frameworkHandle, "CGDataProviderCreateWithData")
		registerFunc(&_cGDataProviderCreateWithFilename, frameworkHandle, "CGDataProviderCreateWithFilename")
		registerFunc(&_cGDataProviderCreateWithURL, frameworkHandle, "CGDataProviderCreateWithURL")
		registerFunc(&_cGDataProviderGetInfo, frameworkHandle, "CGDataProviderGetInfo")
		registerFunc(&_cGDataProviderGetTypeID, frameworkHandle, "CGDataProviderGetTypeID")
		registerFunc(&_cGDataProviderRelease, frameworkHandle, "CGDataProviderRelease")
		registerFunc(&_cGDataProviderRetain, frameworkHandle, "CGDataProviderRetain")
		registerFunc(&_cGDirectDisplayCopyCurrentMetalDevice, frameworkHandle, "CGDirectDisplayCopyCurrentMetalDevice")
		registerFunc(&_cGDisplayBounds, frameworkHandle, "CGDisplayBounds")
		registerFunc(&_cGDisplayCapture, frameworkHandle, "CGDisplayCapture")
		registerFunc(&_cGDisplayCaptureWithOptions, frameworkHandle, "CGDisplayCaptureWithOptions")
		registerFunc(&_cGDisplayCopyAllDisplayModes, frameworkHandle, "CGDisplayCopyAllDisplayModes")
		registerFunc(&_cGDisplayCopyColorSpace, frameworkHandle, "CGDisplayCopyColorSpace")
		registerFunc(&_cGDisplayCopyDisplayMode, frameworkHandle, "CGDisplayCopyDisplayMode")
		registerFunc(&_cGDisplayCreateImage, frameworkHandle, "CGDisplayCreateImage")
		registerFunc(&_cGDisplayCreateImageForRect, frameworkHandle, "CGDisplayCreateImageForRect")
		registerFunc(&_cGDisplayFade, frameworkHandle, "CGDisplayFade")
		registerFunc(&_cGDisplayGammaTableCapacity, frameworkHandle, "CGDisplayGammaTableCapacity")
		registerFunc(&_cGDisplayGetDrawingContext, frameworkHandle, "CGDisplayGetDrawingContext")
		registerFunc(&_cGDisplayHideCursor, frameworkHandle, "CGDisplayHideCursor")
		registerFunc(&_cGDisplayIDToOpenGLDisplayMask, frameworkHandle, "CGDisplayIDToOpenGLDisplayMask")
		registerFunc(&_cGDisplayIsActive, frameworkHandle, "CGDisplayIsActive")
		registerFunc(&_cGDisplayIsAlwaysInMirrorSet, frameworkHandle, "CGDisplayIsAlwaysInMirrorSet")
		registerFunc(&_cGDisplayIsAsleep, frameworkHandle, "CGDisplayIsAsleep")
		registerFunc(&_cGDisplayIsBuiltin, frameworkHandle, "CGDisplayIsBuiltin")
		registerFunc(&_cGDisplayIsInHWMirrorSet, frameworkHandle, "CGDisplayIsInHWMirrorSet")
		registerFunc(&_cGDisplayIsInMirrorSet, frameworkHandle, "CGDisplayIsInMirrorSet")
		registerFunc(&_cGDisplayIsMain, frameworkHandle, "CGDisplayIsMain")
		registerFunc(&_cGDisplayIsOnline, frameworkHandle, "CGDisplayIsOnline")
		registerFunc(&_cGDisplayIsStereo, frameworkHandle, "CGDisplayIsStereo")
		registerFunc(&_cGDisplayMirrorsDisplay, frameworkHandle, "CGDisplayMirrorsDisplay")
		registerFunc(&_cGDisplayModeGetHeight, frameworkHandle, "CGDisplayModeGetHeight")
		registerFunc(&_cGDisplayModeGetIODisplayModeID, frameworkHandle, "CGDisplayModeGetIODisplayModeID")
		registerFunc(&_cGDisplayModeGetIOFlags, frameworkHandle, "CGDisplayModeGetIOFlags")
		registerFunc(&_cGDisplayModeGetPixelHeight, frameworkHandle, "CGDisplayModeGetPixelHeight")
		registerFunc(&_cGDisplayModeGetPixelWidth, frameworkHandle, "CGDisplayModeGetPixelWidth")
		registerFunc(&_cGDisplayModeGetRefreshRate, frameworkHandle, "CGDisplayModeGetRefreshRate")
		registerFunc(&_cGDisplayModeGetTypeID, frameworkHandle, "CGDisplayModeGetTypeID")
		registerFunc(&_cGDisplayModeGetWidth, frameworkHandle, "CGDisplayModeGetWidth")
		registerFunc(&_cGDisplayModeIsUsableForDesktopGUI, frameworkHandle, "CGDisplayModeIsUsableForDesktopGUI")
		registerFunc(&_cGDisplayModeRelease, frameworkHandle, "CGDisplayModeRelease")
		registerFunc(&_cGDisplayModeRetain, frameworkHandle, "CGDisplayModeRetain")
		registerFunc(&_cGDisplayModelNumber, frameworkHandle, "CGDisplayModelNumber")
		registerFunc(&_cGDisplayMoveCursorToPoint, frameworkHandle, "CGDisplayMoveCursorToPoint")
		registerFunc(&_cGDisplayPixelsHigh, frameworkHandle, "CGDisplayPixelsHigh")
		registerFunc(&_cGDisplayPixelsWide, frameworkHandle, "CGDisplayPixelsWide")
		registerFunc(&_cGDisplayPrimaryDisplay, frameworkHandle, "CGDisplayPrimaryDisplay")
		registerFunc(&_cGDisplayRegisterReconfigurationCallback, frameworkHandle, "CGDisplayRegisterReconfigurationCallback")
		registerFunc(&_cGDisplayRelease, frameworkHandle, "CGDisplayRelease")
		registerFunc(&_cGDisplayRemoveReconfigurationCallback, frameworkHandle, "CGDisplayRemoveReconfigurationCallback")
		registerFunc(&_cGDisplayRestoreColorSyncSettings, frameworkHandle, "CGDisplayRestoreColorSyncSettings")
		registerFunc(&_cGDisplayRotation, frameworkHandle, "CGDisplayRotation")
		registerFunc(&_cGDisplayScreenSize, frameworkHandle, "CGDisplayScreenSize")
		registerFunc(&_cGDisplaySerialNumber, frameworkHandle, "CGDisplaySerialNumber")
		registerFunc(&_cGDisplaySetDisplayMode, frameworkHandle, "CGDisplaySetDisplayMode")
		registerFunc(&_cGDisplaySetStereoOperation, frameworkHandle, "CGDisplaySetStereoOperation")
		registerFunc(&_cGDisplayShowCursor, frameworkHandle, "CGDisplayShowCursor")
		registerFunc(&_cGDisplayUnitNumber, frameworkHandle, "CGDisplayUnitNumber")
		registerFunc(&_cGDisplayUsesOpenGLAcceleration, frameworkHandle, "CGDisplayUsesOpenGLAcceleration")
		registerFunc(&_cGDisplayVendorNumber, frameworkHandle, "CGDisplayVendorNumber")
		registerFunc(&_cGEXRToneMappingGammaGetDefaultOptions, frameworkHandle, "CGEXRToneMappingGammaGetDefaultOptions")
		registerFunc(&_cGErrorSetCallback, frameworkHandle, "CGErrorSetCallback")
		registerFunc(&_cGEventCreate, frameworkHandle, "CGEventCreate")
		registerFunc(&_cGEventCreateCopy, frameworkHandle, "CGEventCreateCopy")
		registerFunc(&_cGEventCreateData, frameworkHandle, "CGEventCreateData")
		registerFunc(&_cGEventCreateFromData, frameworkHandle, "CGEventCreateFromData")
		registerFunc(&_cGEventCreateKeyboardEvent, frameworkHandle, "CGEventCreateKeyboardEvent")
		registerFunc(&_cGEventCreateMouseEvent, frameworkHandle, "CGEventCreateMouseEvent")
		registerFunc(&_cGEventCreateScrollWheelEvent, frameworkHandle, "CGEventCreateScrollWheelEvent")
		registerFunc(&_cGEventCreateScrollWheelEvent2, frameworkHandle, "CGEventCreateScrollWheelEvent2")
		registerFunc(&_cGEventCreateSourceFromEvent, frameworkHandle, "CGEventCreateSourceFromEvent")
		registerFunc(&_cGEventGetDoubleValueField, frameworkHandle, "CGEventGetDoubleValueField")
		registerFunc(&_cGEventGetFlags, frameworkHandle, "CGEventGetFlags")
		registerFunc(&_cGEventGetIntegerValueField, frameworkHandle, "CGEventGetIntegerValueField")
		registerFunc(&_cGEventGetLocation, frameworkHandle, "CGEventGetLocation")
		registerFunc(&_cGEventGetTimestamp, frameworkHandle, "CGEventGetTimestamp")
		registerFunc(&_cGEventGetType, frameworkHandle, "CGEventGetType")
		registerFunc(&_cGEventGetTypeID, frameworkHandle, "CGEventGetTypeID")
		registerFunc(&_cGEventGetUnflippedLocation, frameworkHandle, "CGEventGetUnflippedLocation")
		registerFunc(&_cGEventKeyboardGetUnicodeString, frameworkHandle, "CGEventKeyboardGetUnicodeString")
		registerFunc(&_cGEventKeyboardSetUnicodeString, frameworkHandle, "CGEventKeyboardSetUnicodeString")
		registerFunc(&_cGEventPost, frameworkHandle, "CGEventPost")
		registerFunc(&_cGEventPostToPSN, frameworkHandle, "CGEventPostToPSN")
		registerFunc(&_cGEventPostToPid, frameworkHandle, "CGEventPostToPid")
		registerFunc(&_cGEventSetDoubleValueField, frameworkHandle, "CGEventSetDoubleValueField")
		registerFunc(&_cGEventSetFlags, frameworkHandle, "CGEventSetFlags")
		registerFunc(&_cGEventSetIntegerValueField, frameworkHandle, "CGEventSetIntegerValueField")
		registerFunc(&_cGEventSetLocation, frameworkHandle, "CGEventSetLocation")
		registerFunc(&_cGEventSetSource, frameworkHandle, "CGEventSetSource")
		registerFunc(&_cGEventSetTimestamp, frameworkHandle, "CGEventSetTimestamp")
		registerFunc(&_cGEventSetType, frameworkHandle, "CGEventSetType")
		registerFunc(&_cGEventSourceButtonState, frameworkHandle, "CGEventSourceButtonState")
		registerFunc(&_cGEventSourceCounterForEventType, frameworkHandle, "CGEventSourceCounterForEventType")
		registerFunc(&_cGEventSourceCreate, frameworkHandle, "CGEventSourceCreate")
		registerFunc(&_cGEventSourceFlagsState, frameworkHandle, "CGEventSourceFlagsState")
		registerFunc(&_cGEventSourceGetKeyboardType, frameworkHandle, "CGEventSourceGetKeyboardType")
		registerFunc(&_cGEventSourceGetLocalEventsFilterDuringSuppressionState, frameworkHandle, "CGEventSourceGetLocalEventsFilterDuringSuppressionState")
		registerFunc(&_cGEventSourceGetLocalEventsSuppressionInterval, frameworkHandle, "CGEventSourceGetLocalEventsSuppressionInterval")
		registerFunc(&_cGEventSourceGetPixelsPerLine, frameworkHandle, "CGEventSourceGetPixelsPerLine")
		registerFunc(&_cGEventSourceGetSourceStateID, frameworkHandle, "CGEventSourceGetSourceStateID")
		registerFunc(&_cGEventSourceGetTypeID, frameworkHandle, "CGEventSourceGetTypeID")
		registerFunc(&_cGEventSourceGetUserData, frameworkHandle, "CGEventSourceGetUserData")
		registerFunc(&_cGEventSourceKeyState, frameworkHandle, "CGEventSourceKeyState")
		registerFunc(&_cGEventSourceSecondsSinceLastEventType, frameworkHandle, "CGEventSourceSecondsSinceLastEventType")
		registerFunc(&_cGEventSourceSetKeyboardType, frameworkHandle, "CGEventSourceSetKeyboardType")
		registerFunc(&_cGEventSourceSetLocalEventsFilterDuringSuppressionState, frameworkHandle, "CGEventSourceSetLocalEventsFilterDuringSuppressionState")
		registerFunc(&_cGEventSourceSetLocalEventsSuppressionInterval, frameworkHandle, "CGEventSourceSetLocalEventsSuppressionInterval")
		registerFunc(&_cGEventSourceSetPixelsPerLine, frameworkHandle, "CGEventSourceSetPixelsPerLine")
		registerFunc(&_cGEventSourceSetUserData, frameworkHandle, "CGEventSourceSetUserData")
		registerFunc(&_cGEventTapCreate, frameworkHandle, "CGEventTapCreate")
		registerFunc(&_cGEventTapCreateForPSN, frameworkHandle, "CGEventTapCreateForPSN")
		registerFunc(&_cGEventTapCreateForPid, frameworkHandle, "CGEventTapCreateForPid")
		registerFunc(&_cGEventTapEnable, frameworkHandle, "CGEventTapEnable")
		registerFunc(&_cGEventTapIsEnabled, frameworkHandle, "CGEventTapIsEnabled")
		registerFunc(&_cGEventTapPostEvent, frameworkHandle, "CGEventTapPostEvent")
		registerFunc(&_cGFontCanCreatePostScriptSubset, frameworkHandle, "CGFontCanCreatePostScriptSubset")
		registerFunc(&_cGFontCopyFullName, frameworkHandle, "CGFontCopyFullName")
		registerFunc(&_cGFontCopyGlyphNameForGlyph, frameworkHandle, "CGFontCopyGlyphNameForGlyph")
		registerFunc(&_cGFontCopyPostScriptName, frameworkHandle, "CGFontCopyPostScriptName")
		registerFunc(&_cGFontCopyTableForTag, frameworkHandle, "CGFontCopyTableForTag")
		registerFunc(&_cGFontCopyTableTags, frameworkHandle, "CGFontCopyTableTags")
		registerFunc(&_cGFontCopyVariationAxes, frameworkHandle, "CGFontCopyVariationAxes")
		registerFunc(&_cGFontCopyVariations, frameworkHandle, "CGFontCopyVariations")
		registerFunc(&_cGFontCreateCopyWithVariations, frameworkHandle, "CGFontCreateCopyWithVariations")
		registerFunc(&_cGFontCreatePostScriptEncoding, frameworkHandle, "CGFontCreatePostScriptEncoding")
		registerFunc(&_cGFontCreatePostScriptSubset, frameworkHandle, "CGFontCreatePostScriptSubset")
		registerFunc(&_cGFontCreateWithDataProvider, frameworkHandle, "CGFontCreateWithDataProvider")
		registerFunc(&_cGFontCreateWithFontName, frameworkHandle, "CGFontCreateWithFontName")
		registerFunc(&_cGFontGetAscent, frameworkHandle, "CGFontGetAscent")
		registerFunc(&_cGFontGetCapHeight, frameworkHandle, "CGFontGetCapHeight")
		registerFunc(&_cGFontGetDescent, frameworkHandle, "CGFontGetDescent")
		registerFunc(&_cGFontGetFontBBox, frameworkHandle, "CGFontGetFontBBox")
		registerFunc(&_cGFontGetGlyphAdvances, frameworkHandle, "CGFontGetGlyphAdvances")
		registerFunc(&_cGFontGetGlyphBBoxes, frameworkHandle, "CGFontGetGlyphBBoxes")
		registerFunc(&_cGFontGetGlyphWithGlyphName, frameworkHandle, "CGFontGetGlyphWithGlyphName")
		registerFunc(&_cGFontGetItalicAngle, frameworkHandle, "CGFontGetItalicAngle")
		registerFunc(&_cGFontGetLeading, frameworkHandle, "CGFontGetLeading")
		registerFunc(&_cGFontGetNumberOfGlyphs, frameworkHandle, "CGFontGetNumberOfGlyphs")
		registerFunc(&_cGFontGetStemV, frameworkHandle, "CGFontGetStemV")
		registerFunc(&_cGFontGetTypeID, frameworkHandle, "CGFontGetTypeID")
		registerFunc(&_cGFontGetUnitsPerEm, frameworkHandle, "CGFontGetUnitsPerEm")
		registerFunc(&_cGFontGetXHeight, frameworkHandle, "CGFontGetXHeight")
		registerFunc(&_cGFontRelease, frameworkHandle, "CGFontRelease")
		registerFunc(&_cGFontRetain, frameworkHandle, "CGFontRetain")
		registerFunc(&_cGFunctionCreate, frameworkHandle, "CGFunctionCreate")
		registerFunc(&_cGFunctionGetTypeID, frameworkHandle, "CGFunctionGetTypeID")
		registerFunc(&_cGFunctionRelease, frameworkHandle, "CGFunctionRelease")
		registerFunc(&_cGFunctionRetain, frameworkHandle, "CGFunctionRetain")
		registerFunc(&_cGGetActiveDisplayList, frameworkHandle, "CGGetActiveDisplayList")
		registerFunc(&_cGGetDisplayTransferByFormula, frameworkHandle, "CGGetDisplayTransferByFormula")
		registerFunc(&_cGGetDisplayTransferByTable, frameworkHandle, "CGGetDisplayTransferByTable")
		registerFunc(&_cGGetDisplaysWithOpenGLDisplayMask, frameworkHandle, "CGGetDisplaysWithOpenGLDisplayMask")
		registerFunc(&_cGGetDisplaysWithPoint, frameworkHandle, "CGGetDisplaysWithPoint")
		registerFunc(&_cGGetDisplaysWithRect, frameworkHandle, "CGGetDisplaysWithRect")
		registerFunc(&_cGGetEventTapList, frameworkHandle, "CGGetEventTapList")
		registerFunc(&_cGGetLastMouseDelta, frameworkHandle, "CGGetLastMouseDelta")
		registerFunc(&_cGGetOnlineDisplayList, frameworkHandle, "CGGetOnlineDisplayList")
		registerFunc(&_cGGradientCreateWithColorComponents, frameworkHandle, "CGGradientCreateWithColorComponents")
		registerFunc(&_cGGradientCreateWithColors, frameworkHandle, "CGGradientCreateWithColors")
		registerFunc(&_cGGradientCreateWithContentHeadroom, frameworkHandle, "CGGradientCreateWithContentHeadroom")
		registerFunc(&_cGGradientGetContentHeadroom, frameworkHandle, "CGGradientGetContentHeadroom")
		registerFunc(&_cGGradientGetTypeID, frameworkHandle, "CGGradientGetTypeID")
		registerFunc(&_cGGradientRelease, frameworkHandle, "CGGradientRelease")
		registerFunc(&_cGGradientRetain, frameworkHandle, "CGGradientRetain")
		registerFunc(&_cGImageCalculateContentAverageLightLevel, frameworkHandle, "CGImageCalculateContentAverageLightLevel")
		registerFunc(&_cGImageCalculateContentHeadroom, frameworkHandle, "CGImageCalculateContentHeadroom")
		registerFunc(&_cGImageContainsImageSpecificToneMappingMetadata, frameworkHandle, "CGImageContainsImageSpecificToneMappingMetadata")
		registerFunc(&_cGImageCreate, frameworkHandle, "CGImageCreate")
		registerFunc(&_cGImageCreateCopy, frameworkHandle, "CGImageCreateCopy")
		registerFunc(&_cGImageCreateCopyWithCalculatedHDRStats, frameworkHandle, "CGImageCreateCopyWithCalculatedHDRStats")
		registerFunc(&_cGImageCreateCopyWithColorSpace, frameworkHandle, "CGImageCreateCopyWithColorSpace")
		registerFunc(&_cGImageCreateCopyWithContentAverageLightLevel, frameworkHandle, "CGImageCreateCopyWithContentAverageLightLevel")
		registerFunc(&_cGImageCreateCopyWithContentHeadroom, frameworkHandle, "CGImageCreateCopyWithContentHeadroom")
		registerFunc(&_cGImageCreateWithContentHeadroom, frameworkHandle, "CGImageCreateWithContentHeadroom")
		registerFunc(&_cGImageCreateWithImageInRect, frameworkHandle, "CGImageCreateWithImageInRect")
		registerFunc(&_cGImageCreateWithJPEGDataProvider, frameworkHandle, "CGImageCreateWithJPEGDataProvider")
		registerFunc(&_cGImageCreateWithMask, frameworkHandle, "CGImageCreateWithMask")
		registerFunc(&_cGImageCreateWithMaskingColors, frameworkHandle, "CGImageCreateWithMaskingColors")
		registerFunc(&_cGImageCreateWithPNGDataProvider, frameworkHandle, "CGImageCreateWithPNGDataProvider")
		registerFunc(&_cGImageGetAlphaInfo, frameworkHandle, "CGImageGetAlphaInfo")
		registerFunc(&_cGImageGetBitmapInfo, frameworkHandle, "CGImageGetBitmapInfo")
		registerFunc(&_cGImageGetBitsPerComponent, frameworkHandle, "CGImageGetBitsPerComponent")
		registerFunc(&_cGImageGetBitsPerPixel, frameworkHandle, "CGImageGetBitsPerPixel")
		registerFunc(&_cGImageGetByteOrderInfo, frameworkHandle, "CGImageGetByteOrderInfo")
		registerFunc(&_cGImageGetBytesPerRow, frameworkHandle, "CGImageGetBytesPerRow")
		registerFunc(&_cGImageGetColorSpace, frameworkHandle, "CGImageGetColorSpace")
		registerFunc(&_cGImageGetContentAverageLightLevel, frameworkHandle, "CGImageGetContentAverageLightLevel")
		registerFunc(&_cGImageGetContentHeadroom, frameworkHandle, "CGImageGetContentHeadroom")
		registerFunc(&_cGImageGetDataProvider, frameworkHandle, "CGImageGetDataProvider")
		registerFunc(&_cGImageGetDecode, frameworkHandle, "CGImageGetDecode")
		registerFunc(&_cGImageGetHeight, frameworkHandle, "CGImageGetHeight")
		registerFunc(&_cGImageGetPixelFormatInfo, frameworkHandle, "CGImageGetPixelFormatInfo")
		registerFunc(&_cGImageGetRenderingIntent, frameworkHandle, "CGImageGetRenderingIntent")
		registerFunc(&_cGImageGetShouldInterpolate, frameworkHandle, "CGImageGetShouldInterpolate")
		registerFunc(&_cGImageGetTypeID, frameworkHandle, "CGImageGetTypeID")
		registerFunc(&_cGImageGetUTType, frameworkHandle, "CGImageGetUTType")
		registerFunc(&_cGImageGetWidth, frameworkHandle, "CGImageGetWidth")
		registerFunc(&_cGImageIsMask, frameworkHandle, "CGImageIsMask")
		registerFunc(&_cGImageMaskCreate, frameworkHandle, "CGImageMaskCreate")
		registerFunc(&_cGImageRelease, frameworkHandle, "CGImageRelease")
		registerFunc(&_cGImageRetain, frameworkHandle, "CGImageRetain")
		registerFunc(&_cGImageShouldToneMap, frameworkHandle, "CGImageShouldToneMap")
		registerFunc(&_cGLayerCreateWithContext, frameworkHandle, "CGLayerCreateWithContext")
		registerFunc(&_cGLayerGetContext, frameworkHandle, "CGLayerGetContext")
		registerFunc(&_cGLayerGetSize, frameworkHandle, "CGLayerGetSize")
		registerFunc(&_cGLayerGetTypeID, frameworkHandle, "CGLayerGetTypeID")
		registerFunc(&_cGLayerRelease, frameworkHandle, "CGLayerRelease")
		registerFunc(&_cGLayerRetain, frameworkHandle, "CGLayerRetain")
		registerFunc(&_cGMainDisplayID, frameworkHandle, "CGMainDisplayID")
		registerFunc(&_cGOpenGLDisplayMaskToDisplayID, frameworkHandle, "CGOpenGLDisplayMaskToDisplayID")
		registerFunc(&_cGPDFArrayApplyBlock, frameworkHandle, "CGPDFArrayApplyBlock")
		registerFunc(&_cGPDFArrayGetArray, frameworkHandle, "CGPDFArrayGetArray")
		registerFunc(&_cGPDFArrayGetBoolean, frameworkHandle, "CGPDFArrayGetBoolean")
		registerFunc(&_cGPDFArrayGetCount, frameworkHandle, "CGPDFArrayGetCount")
		registerFunc(&_cGPDFArrayGetDictionary, frameworkHandle, "CGPDFArrayGetDictionary")
		registerFunc(&_cGPDFArrayGetInteger, frameworkHandle, "CGPDFArrayGetInteger")
		registerFunc(&_cGPDFArrayGetName, frameworkHandle, "CGPDFArrayGetName")
		registerFunc(&_cGPDFArrayGetNull, frameworkHandle, "CGPDFArrayGetNull")
		registerFunc(&_cGPDFArrayGetNumber, frameworkHandle, "CGPDFArrayGetNumber")
		registerFunc(&_cGPDFArrayGetObject, frameworkHandle, "CGPDFArrayGetObject")
		registerFunc(&_cGPDFArrayGetStream, frameworkHandle, "CGPDFArrayGetStream")
		registerFunc(&_cGPDFArrayGetString, frameworkHandle, "CGPDFArrayGetString")
		registerFunc(&_cGPDFContentStreamCreateWithPage, frameworkHandle, "CGPDFContentStreamCreateWithPage")
		registerFunc(&_cGPDFContentStreamCreateWithStream, frameworkHandle, "CGPDFContentStreamCreateWithStream")
		registerFunc(&_cGPDFContentStreamGetResource, frameworkHandle, "CGPDFContentStreamGetResource")
		registerFunc(&_cGPDFContentStreamGetStreams, frameworkHandle, "CGPDFContentStreamGetStreams")
		registerFunc(&_cGPDFContentStreamRelease, frameworkHandle, "CGPDFContentStreamRelease")
		registerFunc(&_cGPDFContentStreamRetain, frameworkHandle, "CGPDFContentStreamRetain")
		registerFunc(&_cGPDFContextAddDestinationAtPoint, frameworkHandle, "CGPDFContextAddDestinationAtPoint")
		registerFunc(&_cGPDFContextAddDocumentMetadata, frameworkHandle, "CGPDFContextAddDocumentMetadata")
		registerFunc(&_cGPDFContextBeginPage, frameworkHandle, "CGPDFContextBeginPage")
		registerFunc(&_cGPDFContextBeginTag, frameworkHandle, "CGPDFContextBeginTag")
		registerFunc(&_cGPDFContextClose, frameworkHandle, "CGPDFContextClose")
		registerFunc(&_cGPDFContextCreate, frameworkHandle, "CGPDFContextCreate")
		registerFunc(&_cGPDFContextCreateWithURL, frameworkHandle, "CGPDFContextCreateWithURL")
		registerFunc(&_cGPDFContextEndPage, frameworkHandle, "CGPDFContextEndPage")
		registerFunc(&_cGPDFContextEndTag, frameworkHandle, "CGPDFContextEndTag")
		registerFunc(&_cGPDFContextSetDestinationForRect, frameworkHandle, "CGPDFContextSetDestinationForRect")
		registerFunc(&_cGPDFContextSetIDTree, frameworkHandle, "CGPDFContextSetIDTree")
		registerFunc(&_cGPDFContextSetOutline, frameworkHandle, "CGPDFContextSetOutline")
		registerFunc(&_cGPDFContextSetPageTagStructureTree, frameworkHandle, "CGPDFContextSetPageTagStructureTree")
		registerFunc(&_cGPDFContextSetParentTree, frameworkHandle, "CGPDFContextSetParentTree")
		registerFunc(&_cGPDFContextSetURLForRect, frameworkHandle, "CGPDFContextSetURLForRect")
		registerFunc(&_cGPDFDictionaryApplyBlock, frameworkHandle, "CGPDFDictionaryApplyBlock")
		registerFunc(&_cGPDFDictionaryApplyFunction, frameworkHandle, "CGPDFDictionaryApplyFunction")
		registerFunc(&_cGPDFDictionaryGetArray, frameworkHandle, "CGPDFDictionaryGetArray")
		registerFunc(&_cGPDFDictionaryGetBoolean, frameworkHandle, "CGPDFDictionaryGetBoolean")
		registerFunc(&_cGPDFDictionaryGetCount, frameworkHandle, "CGPDFDictionaryGetCount")
		registerFunc(&_cGPDFDictionaryGetDictionary, frameworkHandle, "CGPDFDictionaryGetDictionary")
		registerFunc(&_cGPDFDictionaryGetInteger, frameworkHandle, "CGPDFDictionaryGetInteger")
		registerFunc(&_cGPDFDictionaryGetName, frameworkHandle, "CGPDFDictionaryGetName")
		registerFunc(&_cGPDFDictionaryGetNumber, frameworkHandle, "CGPDFDictionaryGetNumber")
		registerFunc(&_cGPDFDictionaryGetObject, frameworkHandle, "CGPDFDictionaryGetObject")
		registerFunc(&_cGPDFDictionaryGetStream, frameworkHandle, "CGPDFDictionaryGetStream")
		registerFunc(&_cGPDFDictionaryGetString, frameworkHandle, "CGPDFDictionaryGetString")
		registerFunc(&_cGPDFDocumentAllowsCopying, frameworkHandle, "CGPDFDocumentAllowsCopying")
		registerFunc(&_cGPDFDocumentAllowsPrinting, frameworkHandle, "CGPDFDocumentAllowsPrinting")
		registerFunc(&_cGPDFDocumentCreateWithProvider, frameworkHandle, "CGPDFDocumentCreateWithProvider")
		registerFunc(&_cGPDFDocumentCreateWithURL, frameworkHandle, "CGPDFDocumentCreateWithURL")
		registerFunc(&_cGPDFDocumentGetAccessPermissions, frameworkHandle, "CGPDFDocumentGetAccessPermissions")
		registerFunc(&_cGPDFDocumentGetArtBox, frameworkHandle, "CGPDFDocumentGetArtBox")
		registerFunc(&_cGPDFDocumentGetBleedBox, frameworkHandle, "CGPDFDocumentGetBleedBox")
		registerFunc(&_cGPDFDocumentGetCatalog, frameworkHandle, "CGPDFDocumentGetCatalog")
		registerFunc(&_cGPDFDocumentGetCropBox, frameworkHandle, "CGPDFDocumentGetCropBox")
		registerFunc(&_cGPDFDocumentGetID, frameworkHandle, "CGPDFDocumentGetID")
		registerFunc(&_cGPDFDocumentGetInfo, frameworkHandle, "CGPDFDocumentGetInfo")
		registerFunc(&_cGPDFDocumentGetMediaBox, frameworkHandle, "CGPDFDocumentGetMediaBox")
		registerFunc(&_cGPDFDocumentGetNumberOfPages, frameworkHandle, "CGPDFDocumentGetNumberOfPages")
		registerFunc(&_cGPDFDocumentGetOutline, frameworkHandle, "CGPDFDocumentGetOutline")
		registerFunc(&_cGPDFDocumentGetPage, frameworkHandle, "CGPDFDocumentGetPage")
		registerFunc(&_cGPDFDocumentGetRotationAngle, frameworkHandle, "CGPDFDocumentGetRotationAngle")
		registerFunc(&_cGPDFDocumentGetTrimBox, frameworkHandle, "CGPDFDocumentGetTrimBox")
		registerFunc(&_cGPDFDocumentGetTypeID, frameworkHandle, "CGPDFDocumentGetTypeID")
		registerFunc(&_cGPDFDocumentGetVersion, frameworkHandle, "CGPDFDocumentGetVersion")
		registerFunc(&_cGPDFDocumentIsEncrypted, frameworkHandle, "CGPDFDocumentIsEncrypted")
		registerFunc(&_cGPDFDocumentIsUnlocked, frameworkHandle, "CGPDFDocumentIsUnlocked")
		registerFunc(&_cGPDFDocumentRelease, frameworkHandle, "CGPDFDocumentRelease")
		registerFunc(&_cGPDFDocumentRetain, frameworkHandle, "CGPDFDocumentRetain")
		registerFunc(&_cGPDFDocumentUnlockWithPassword, frameworkHandle, "CGPDFDocumentUnlockWithPassword")
		registerFunc(&_cGPDFObjectGetType, frameworkHandle, "CGPDFObjectGetType")
		registerFunc(&_cGPDFObjectGetValue, frameworkHandle, "CGPDFObjectGetValue")
		registerFunc(&_cGPDFOperatorTableCreate, frameworkHandle, "CGPDFOperatorTableCreate")
		registerFunc(&_cGPDFOperatorTableRelease, frameworkHandle, "CGPDFOperatorTableRelease")
		registerFunc(&_cGPDFOperatorTableRetain, frameworkHandle, "CGPDFOperatorTableRetain")
		registerFunc(&_cGPDFOperatorTableSetCallback, frameworkHandle, "CGPDFOperatorTableSetCallback")
		registerFunc(&_cGPDFPageGetBoxRect, frameworkHandle, "CGPDFPageGetBoxRect")
		registerFunc(&_cGPDFPageGetDictionary, frameworkHandle, "CGPDFPageGetDictionary")
		registerFunc(&_cGPDFPageGetDocument, frameworkHandle, "CGPDFPageGetDocument")
		registerFunc(&_cGPDFPageGetDrawingTransform, frameworkHandle, "CGPDFPageGetDrawingTransform")
		registerFunc(&_cGPDFPageGetPageNumber, frameworkHandle, "CGPDFPageGetPageNumber")
		registerFunc(&_cGPDFPageGetRotationAngle, frameworkHandle, "CGPDFPageGetRotationAngle")
		registerFunc(&_cGPDFPageGetTypeID, frameworkHandle, "CGPDFPageGetTypeID")
		registerFunc(&_cGPDFPageRelease, frameworkHandle, "CGPDFPageRelease")
		registerFunc(&_cGPDFPageRetain, frameworkHandle, "CGPDFPageRetain")
		registerFunc(&_cGPDFScannerCreate, frameworkHandle, "CGPDFScannerCreate")
		registerFunc(&_cGPDFScannerGetContentStream, frameworkHandle, "CGPDFScannerGetContentStream")
		registerFunc(&_cGPDFScannerPopArray, frameworkHandle, "CGPDFScannerPopArray")
		registerFunc(&_cGPDFScannerPopBoolean, frameworkHandle, "CGPDFScannerPopBoolean")
		registerFunc(&_cGPDFScannerPopDictionary, frameworkHandle, "CGPDFScannerPopDictionary")
		registerFunc(&_cGPDFScannerPopInteger, frameworkHandle, "CGPDFScannerPopInteger")
		registerFunc(&_cGPDFScannerPopName, frameworkHandle, "CGPDFScannerPopName")
		registerFunc(&_cGPDFScannerPopNumber, frameworkHandle, "CGPDFScannerPopNumber")
		registerFunc(&_cGPDFScannerPopObject, frameworkHandle, "CGPDFScannerPopObject")
		registerFunc(&_cGPDFScannerPopStream, frameworkHandle, "CGPDFScannerPopStream")
		registerFunc(&_cGPDFScannerPopString, frameworkHandle, "CGPDFScannerPopString")
		registerFunc(&_cGPDFScannerRelease, frameworkHandle, "CGPDFScannerRelease")
		registerFunc(&_cGPDFScannerRetain, frameworkHandle, "CGPDFScannerRetain")
		registerFunc(&_cGPDFScannerScan, frameworkHandle, "CGPDFScannerScan")
		registerFunc(&_cGPDFScannerStop, frameworkHandle, "CGPDFScannerStop")
		registerFunc(&_cGPDFStreamCopyData, frameworkHandle, "CGPDFStreamCopyData")
		registerFunc(&_cGPDFStreamGetDictionary, frameworkHandle, "CGPDFStreamGetDictionary")
		registerFunc(&_cGPDFStringCopyDate, frameworkHandle, "CGPDFStringCopyDate")
		registerFunc(&_cGPDFStringCopyTextString, frameworkHandle, "CGPDFStringCopyTextString")
		registerFunc(&_cGPDFStringGetBytePtr, frameworkHandle, "CGPDFStringGetBytePtr")
		registerFunc(&_cGPDFStringGetLength, frameworkHandle, "CGPDFStringGetLength")
		registerFunc(&_cGPSConverterAbort, frameworkHandle, "CGPSConverterAbort")
		registerFunc(&_cGPSConverterConvert, frameworkHandle, "CGPSConverterConvert")
		registerFunc(&_cGPSConverterCreate, frameworkHandle, "CGPSConverterCreate")
		registerFunc(&_cGPSConverterGetTypeID, frameworkHandle, "CGPSConverterGetTypeID")
		registerFunc(&_cGPSConverterIsConverting, frameworkHandle, "CGPSConverterIsConverting")
		registerFunc(&_cGPathApply, frameworkHandle, "CGPathApply")
		registerFunc(&_cGPathApplyWithBlock, frameworkHandle, "CGPathApplyWithBlock")
		registerFunc(&_cGPathCloseSubpath, frameworkHandle, "CGPathCloseSubpath")
		registerFunc(&_cGPathContainsPoint, frameworkHandle, "CGPathContainsPoint")
		registerFunc(&_cGPathCreateCopy, frameworkHandle, "CGPathCreateCopy")
		registerFunc(&_cGPathCreateCopyByDashingPath, frameworkHandle, "CGPathCreateCopyByDashingPath")
		registerFunc(&_cGPathCreateCopyByFlattening, frameworkHandle, "CGPathCreateCopyByFlattening")
		registerFunc(&_cGPathCreateCopyByIntersectingPath, frameworkHandle, "CGPathCreateCopyByIntersectingPath")
		registerFunc(&_cGPathCreateCopyByNormalizing, frameworkHandle, "CGPathCreateCopyByNormalizing")
		registerFunc(&_cGPathCreateCopyByStrokingPath, frameworkHandle, "CGPathCreateCopyByStrokingPath")
		registerFunc(&_cGPathCreateCopyBySubtractingPath, frameworkHandle, "CGPathCreateCopyBySubtractingPath")
		registerFunc(&_cGPathCreateCopyBySymmetricDifferenceOfPath, frameworkHandle, "CGPathCreateCopyBySymmetricDifferenceOfPath")
		registerFunc(&_cGPathCreateCopyByTransformingPath, frameworkHandle, "CGPathCreateCopyByTransformingPath")
		registerFunc(&_cGPathCreateCopyByUnioningPath, frameworkHandle, "CGPathCreateCopyByUnioningPath")
		registerFunc(&_cGPathCreateCopyOfLineByIntersectingPath, frameworkHandle, "CGPathCreateCopyOfLineByIntersectingPath")
		registerFunc(&_cGPathCreateCopyOfLineBySubtractingPath, frameworkHandle, "CGPathCreateCopyOfLineBySubtractingPath")
		registerFunc(&_cGPathCreateMutable, frameworkHandle, "CGPathCreateMutable")
		registerFunc(&_cGPathCreateMutableCopy, frameworkHandle, "CGPathCreateMutableCopy")
		registerFunc(&_cGPathCreateMutableCopyByTransformingPath, frameworkHandle, "CGPathCreateMutableCopyByTransformingPath")
		registerFunc(&_cGPathCreateSeparateComponents, frameworkHandle, "CGPathCreateSeparateComponents")
		registerFunc(&_cGPathCreateWithEllipseInRect, frameworkHandle, "CGPathCreateWithEllipseInRect")
		registerFunc(&_cGPathCreateWithRect, frameworkHandle, "CGPathCreateWithRect")
		registerFunc(&_cGPathCreateWithRoundedRect, frameworkHandle, "CGPathCreateWithRoundedRect")
		registerFunc(&_cGPathEqualToPath, frameworkHandle, "CGPathEqualToPath")
		registerFunc(&_cGPathGetBoundingBox, frameworkHandle, "CGPathGetBoundingBox")
		registerFunc(&_cGPathGetCurrentPoint, frameworkHandle, "CGPathGetCurrentPoint")
		registerFunc(&_cGPathGetPathBoundingBox, frameworkHandle, "CGPathGetPathBoundingBox")
		registerFunc(&_cGPathGetTypeID, frameworkHandle, "CGPathGetTypeID")
		registerFunc(&_cGPathIntersectsPath, frameworkHandle, "CGPathIntersectsPath")
		registerFunc(&_cGPathIsEmpty, frameworkHandle, "CGPathIsEmpty")
		registerFunc(&_cGPathIsRect, frameworkHandle, "CGPathIsRect")
		registerFunc(&_cGPathRelease, frameworkHandle, "CGPathRelease")
		registerFunc(&_cGPathRetain, frameworkHandle, "CGPathRetain")
		registerFunc(&_cGPatternCreate, frameworkHandle, "CGPatternCreate")
		registerFunc(&_cGPatternGetTypeID, frameworkHandle, "CGPatternGetTypeID")
		registerFunc(&_cGPatternRelease, frameworkHandle, "CGPatternRelease")
		registerFunc(&_cGPatternRetain, frameworkHandle, "CGPatternRetain")
		registerFunc(&_cGPointApplyAffineTransform, frameworkHandle, "CGPointApplyAffineTransform")
		registerFunc(&_cGPointCreateDictionaryRepresentation, frameworkHandle, "CGPointCreateDictionaryRepresentation")
		registerFunc(&_cGPointEqualToPoint, frameworkHandle, "CGPointEqualToPoint")
		registerFunc(&_cGPointMakeWithDictionaryRepresentation, frameworkHandle, "CGPointMakeWithDictionaryRepresentation")
		registerFunc(&_cGPostMouseEvent, frameworkHandle, "CGPostMouseEvent")
		registerFunc(&_cGPostScrollWheelEvent, frameworkHandle, "CGPostScrollWheelEvent")
		registerFunc(&_cGPreflightListenEventAccess, frameworkHandle, "CGPreflightListenEventAccess")
		registerFunc(&_cGPreflightPostEventAccess, frameworkHandle, "CGPreflightPostEventAccess")
		registerFunc(&_cGPreflightScreenCaptureAccess, frameworkHandle, "CGPreflightScreenCaptureAccess")
		registerFunc(&_cGRectApplyAffineTransform, frameworkHandle, "CGRectApplyAffineTransform")
		registerFunc(&_cGRectContainsPoint, frameworkHandle, "CGRectContainsPoint")
		registerFunc(&_cGRectContainsRect, frameworkHandle, "CGRectContainsRect")
		registerFunc(&_cGRectCreateDictionaryRepresentation, frameworkHandle, "CGRectCreateDictionaryRepresentation")
		registerFunc(&_cGRectDivide, frameworkHandle, "CGRectDivide")
		registerFunc(&_cGRectEqualToRect, frameworkHandle, "CGRectEqualToRect")
		registerFunc(&_cGRectGetHeight, frameworkHandle, "CGRectGetHeight")
		registerFunc(&_cGRectGetMaxX, frameworkHandle, "CGRectGetMaxX")
		registerFunc(&_cGRectGetMaxY, frameworkHandle, "CGRectGetMaxY")
		registerFunc(&_cGRectGetMidX, frameworkHandle, "CGRectGetMidX")
		registerFunc(&_cGRectGetMidY, frameworkHandle, "CGRectGetMidY")
		registerFunc(&_cGRectGetMinX, frameworkHandle, "CGRectGetMinX")
		registerFunc(&_cGRectGetMinY, frameworkHandle, "CGRectGetMinY")
		registerFunc(&_cGRectGetWidth, frameworkHandle, "CGRectGetWidth")
		registerFunc(&_cGRectInset, frameworkHandle, "CGRectInset")
		registerFunc(&_cGRectIntegral, frameworkHandle, "CGRectIntegral")
		registerFunc(&_cGRectIntersection, frameworkHandle, "CGRectIntersection")
		registerFunc(&_cGRectIntersectsRect, frameworkHandle, "CGRectIntersectsRect")
		registerFunc(&_cGRectIsEmpty, frameworkHandle, "CGRectIsEmpty")
		registerFunc(&_cGRectIsInfinite, frameworkHandle, "CGRectIsInfinite")
		registerFunc(&_cGRectIsNull, frameworkHandle, "CGRectIsNull")
		registerFunc(&_cGRectMakeWithDictionaryRepresentation, frameworkHandle, "CGRectMakeWithDictionaryRepresentation")
		registerFunc(&_cGRectOffset, frameworkHandle, "CGRectOffset")
		registerFunc(&_cGRectStandardize, frameworkHandle, "CGRectStandardize")
		registerFunc(&_cGRectUnion, frameworkHandle, "CGRectUnion")
		registerFunc(&_cGReleaseAllDisplays, frameworkHandle, "CGReleaseAllDisplays")
		registerFunc(&_cGReleaseDisplayFadeReservation, frameworkHandle, "CGReleaseDisplayFadeReservation")
		registerFunc(&_cGRenderingBufferLockBytePtr, frameworkHandle, "CGRenderingBufferLockBytePtr")
		registerFunc(&_cGRenderingBufferProviderCreate, frameworkHandle, "CGRenderingBufferProviderCreate")
		registerFunc(&_cGRenderingBufferProviderCreateWithCFData, frameworkHandle, "CGRenderingBufferProviderCreateWithCFData")
		registerFunc(&_cGRenderingBufferProviderGetSize, frameworkHandle, "CGRenderingBufferProviderGetSize")
		registerFunc(&_cGRenderingBufferProviderGetTypeID, frameworkHandle, "CGRenderingBufferProviderGetTypeID")
		registerFunc(&_cGRenderingBufferUnlockBytePtr, frameworkHandle, "CGRenderingBufferUnlockBytePtr")
		registerFunc(&_cGRequestListenEventAccess, frameworkHandle, "CGRequestListenEventAccess")
		registerFunc(&_cGRequestPostEventAccess, frameworkHandle, "CGRequestPostEventAccess")
		registerFunc(&_cGRequestScreenCaptureAccess, frameworkHandle, "CGRequestScreenCaptureAccess")
		registerFunc(&_cGRestorePermanentDisplayConfiguration, frameworkHandle, "CGRestorePermanentDisplayConfiguration")
		registerFunc(&_cGSessionCopyCurrentDictionary, frameworkHandle, "CGSessionCopyCurrentDictionary")
		registerFunc(&_cGSetDisplayTransferByByteTable, frameworkHandle, "CGSetDisplayTransferByByteTable")
		registerFunc(&_cGSetDisplayTransferByFormula, frameworkHandle, "CGSetDisplayTransferByFormula")
		registerFunc(&_cGSetDisplayTransferByTable, frameworkHandle, "CGSetDisplayTransferByTable")
		registerFunc(&_cGShadingCreateAxial, frameworkHandle, "CGShadingCreateAxial")
		registerFunc(&_cGShadingCreateAxialWithContentHeadroom, frameworkHandle, "CGShadingCreateAxialWithContentHeadroom")
		registerFunc(&_cGShadingCreateRadial, frameworkHandle, "CGShadingCreateRadial")
		registerFunc(&_cGShadingCreateRadialWithContentHeadroom, frameworkHandle, "CGShadingCreateRadialWithContentHeadroom")
		registerFunc(&_cGShadingGetContentHeadroom, frameworkHandle, "CGShadingGetContentHeadroom")
		registerFunc(&_cGShadingGetTypeID, frameworkHandle, "CGShadingGetTypeID")
		registerFunc(&_cGShadingRelease, frameworkHandle, "CGShadingRelease")
		registerFunc(&_cGShadingRetain, frameworkHandle, "CGShadingRetain")
		registerFunc(&_cGShieldingWindowID, frameworkHandle, "CGShieldingWindowID")
		registerFunc(&_cGShieldingWindowLevel, frameworkHandle, "CGShieldingWindowLevel")
		registerFunc(&_cGSizeApplyAffineTransform, frameworkHandle, "CGSizeApplyAffineTransform")
		registerFunc(&_cGSizeCreateDictionaryRepresentation, frameworkHandle, "CGSizeCreateDictionaryRepresentation")
		registerFunc(&_cGSizeEqualToSize, frameworkHandle, "CGSizeEqualToSize")
		registerFunc(&_cGSizeMakeWithDictionaryRepresentation, frameworkHandle, "CGSizeMakeWithDictionaryRepresentation")
		registerFunc(&_cGWarpMouseCursorPosition, frameworkHandle, "CGWarpMouseCursorPosition")
		registerFunc(&_cGWindowLevelForKey, frameworkHandle, "CGWindowLevelForKey")
		registerFunc(&_cGWindowListCopyWindowInfo, frameworkHandle, "CGWindowListCopyWindowInfo")
		registerFunc(&_cGWindowListCreate, frameworkHandle, "CGWindowListCreate")
		registerFunc(&_cGWindowListCreateDescriptionFromArray, frameworkHandle, "CGWindowListCreateDescriptionFromArray")
		registerFunc(&_cGWindowListCreateImage, frameworkHandle, "CGWindowListCreateImage")
		registerFunc(&_cGWindowServerCreateServerPort, frameworkHandle, "CGWindowServerCreateServerPort")
	}


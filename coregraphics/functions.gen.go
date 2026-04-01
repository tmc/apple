// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

package coregraphics

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/kernel"
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
		return fmt.Sprintf("CoreGraphics: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("CoreGraphics: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("CoreGraphics: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("CoreGraphics: register symbol %s: %v", name, r)
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

var _cGAcquireDisplayFadeReservation func(seconds float32, token *CGDisplayFadeReservationToken) CGError
var _cGAcquireDisplayFadeReservationErr error

func tryCGAcquireDisplayFadeReservation(seconds float32, token *CGDisplayFadeReservationToken) (CGError, error) {
	if _cGAcquireDisplayFadeReservation == nil {
		return *new(CGError), symbolCallError("CGAcquireDisplayFadeReservation", "10.2", _cGAcquireDisplayFadeReservationErr)
	}
	return _cGAcquireDisplayFadeReservation(seconds, token), nil
}

// CGAcquireDisplayFadeReservation reserves the fade hardware for a specified time interval.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAcquireDisplayFadeReservation(_:_:)
func CGAcquireDisplayFadeReservation(seconds float32, token *CGDisplayFadeReservationToken) CGError {
	result, callErr := tryCGAcquireDisplayFadeReservation(seconds, token)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGAffineTransformConcat func(t1 corefoundation.CGAffineTransform, t2 corefoundation.CGAffineTransform) corefoundation.CGAffineTransform
var _cGAffineTransformConcatErr error

func tryCGAffineTransformConcat(t1 corefoundation.CGAffineTransform, t2 corefoundation.CGAffineTransform) (corefoundation.CGAffineTransform, error) {
	if _cGAffineTransformConcat == nil {
		return corefoundation.CGAffineTransform{}, symbolCallError("CGAffineTransformConcat", "10.0", _cGAffineTransformConcatErr)
	}
	return _cGAffineTransformConcat(t1, t2), nil
}

// CGAffineTransformConcat returns an affine transformation matrix constructed by combining two existing affine transforms.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformConcat(_:_:)
func CGAffineTransformConcat(t1 corefoundation.CGAffineTransform, t2 corefoundation.CGAffineTransform) corefoundation.CGAffineTransform {
	result, callErr := tryCGAffineTransformConcat(t1, t2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGAffineTransformDecompose func(transform corefoundation.CGAffineTransform) uintptr
var _cGAffineTransformDecomposeErr error

func tryCGAffineTransformDecompose(transform corefoundation.CGAffineTransform) (uintptr, error) {
	if _cGAffineTransformDecompose == nil {
		return 0, symbolCallError("CGAffineTransformDecompose", "13.0", _cGAffineTransformDecomposeErr)
	}
	return _cGAffineTransformDecompose(transform), nil
}

// CGAffineTransformDecompose.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformDecompose
func CGAffineTransformDecompose(transform corefoundation.CGAffineTransform) uintptr {
	result, callErr := tryCGAffineTransformDecompose(transform)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGAffineTransformEqualToTransform func(t1 corefoundation.CGAffineTransform, t2 corefoundation.CGAffineTransform) bool
var _cGAffineTransformEqualToTransformErr error

func tryCGAffineTransformEqualToTransform(t1 corefoundation.CGAffineTransform, t2 corefoundation.CGAffineTransform) (bool, error) {
	if _cGAffineTransformEqualToTransform == nil {
		return false, symbolCallError("CGAffineTransformEqualToTransform", "10.4", _cGAffineTransformEqualToTransformErr)
	}
	return _cGAffineTransformEqualToTransform(t1, t2), nil
}

// CGAffineTransformEqualToTransform checks whether two affine transforms are equal.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformEqualToTransform(_:_:)
func CGAffineTransformEqualToTransform(t1 corefoundation.CGAffineTransform, t2 corefoundation.CGAffineTransform) bool {
	result, callErr := tryCGAffineTransformEqualToTransform(t1, t2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGAffineTransformInvert func(t corefoundation.CGAffineTransform) corefoundation.CGAffineTransform
var _cGAffineTransformInvertErr error

func tryCGAffineTransformInvert(t corefoundation.CGAffineTransform) (corefoundation.CGAffineTransform, error) {
	if _cGAffineTransformInvert == nil {
		return corefoundation.CGAffineTransform{}, symbolCallError("CGAffineTransformInvert", "10.0", _cGAffineTransformInvertErr)
	}
	return _cGAffineTransformInvert(t), nil
}

// CGAffineTransformInvert returns an affine transformation matrix constructed by inverting an existing affine transform.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformInvert(_:)
func CGAffineTransformInvert(t corefoundation.CGAffineTransform) corefoundation.CGAffineTransform {
	result, callErr := tryCGAffineTransformInvert(t)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGAffineTransformIsIdentity func(t corefoundation.CGAffineTransform) bool
var _cGAffineTransformIsIdentityErr error

func tryCGAffineTransformIsIdentity(t corefoundation.CGAffineTransform) (bool, error) {
	if _cGAffineTransformIsIdentity == nil {
		return false, symbolCallError("CGAffineTransformIsIdentity", "10.4", _cGAffineTransformIsIdentityErr)
	}
	return _cGAffineTransformIsIdentity(t), nil
}

// CGAffineTransformIsIdentity checks whether an affine transform is the identity transform.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformIsIdentity(_:)
func CGAffineTransformIsIdentity(t corefoundation.CGAffineTransform) bool {
	result, callErr := tryCGAffineTransformIsIdentity(t)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGAffineTransformMake func(a float64, b float64, c float64, d float64, tx float64, ty float64) corefoundation.CGAffineTransform
var _cGAffineTransformMakeErr error

func tryCGAffineTransformMake(a float64, b float64, c float64, d float64, tx float64, ty float64) (corefoundation.CGAffineTransform, error) {
	if _cGAffineTransformMake == nil {
		return corefoundation.CGAffineTransform{}, symbolCallError("CGAffineTransformMake", "10.0", _cGAffineTransformMakeErr)
	}
	return _cGAffineTransformMake(a, b, c, d, tx, ty), nil
}

// CGAffineTransformMake returns an affine transformation matrix constructed from values you provide.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformMake(_:_:_:_:_:_:)
func CGAffineTransformMake(a float64, b float64, c float64, d float64, tx float64, ty float64) corefoundation.CGAffineTransform {
	result, callErr := tryCGAffineTransformMake(a, b, c, d, tx, ty)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGAffineTransformMakeRotation func(angle float64) corefoundation.CGAffineTransform
var _cGAffineTransformMakeRotationErr error

func tryCGAffineTransformMakeRotation(angle float64) (corefoundation.CGAffineTransform, error) {
	if _cGAffineTransformMakeRotation == nil {
		return corefoundation.CGAffineTransform{}, symbolCallError("CGAffineTransformMakeRotation", "10.0", _cGAffineTransformMakeRotationErr)
	}
	return _cGAffineTransformMakeRotation(angle), nil
}

// CGAffineTransformMakeRotation returns an affine transformation matrix constructed from a rotation value you provide.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformMakeRotation(_:)
func CGAffineTransformMakeRotation(angle float64) corefoundation.CGAffineTransform {
	result, callErr := tryCGAffineTransformMakeRotation(angle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGAffineTransformMakeScale func(sx float64, sy float64) corefoundation.CGAffineTransform
var _cGAffineTransformMakeScaleErr error

func tryCGAffineTransformMakeScale(sx float64, sy float64) (corefoundation.CGAffineTransform, error) {
	if _cGAffineTransformMakeScale == nil {
		return corefoundation.CGAffineTransform{}, symbolCallError("CGAffineTransformMakeScale", "10.0", _cGAffineTransformMakeScaleErr)
	}
	return _cGAffineTransformMakeScale(sx, sy), nil
}

// CGAffineTransformMakeScale returns an affine transformation matrix constructed from scaling values you provide.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformMakeScale(_:_:)
func CGAffineTransformMakeScale(sx float64, sy float64) corefoundation.CGAffineTransform {
	result, callErr := tryCGAffineTransformMakeScale(sx, sy)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGAffineTransformMakeTranslation func(tx float64, ty float64) corefoundation.CGAffineTransform
var _cGAffineTransformMakeTranslationErr error

func tryCGAffineTransformMakeTranslation(tx float64, ty float64) (corefoundation.CGAffineTransform, error) {
	if _cGAffineTransformMakeTranslation == nil {
		return corefoundation.CGAffineTransform{}, symbolCallError("CGAffineTransformMakeTranslation", "10.0", _cGAffineTransformMakeTranslationErr)
	}
	return _cGAffineTransformMakeTranslation(tx, ty), nil
}

// CGAffineTransformMakeTranslation returns an affine transformation matrix constructed from translation values you provide.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformMakeTranslation(_:_:)
func CGAffineTransformMakeTranslation(tx float64, ty float64) corefoundation.CGAffineTransform {
	result, callErr := tryCGAffineTransformMakeTranslation(tx, ty)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGAffineTransformMakeWithComponents func(components uintptr) corefoundation.CGAffineTransform
var _cGAffineTransformMakeWithComponentsErr error

func tryCGAffineTransformMakeWithComponents(components uintptr) (corefoundation.CGAffineTransform, error) {
	if _cGAffineTransformMakeWithComponents == nil {
		return corefoundation.CGAffineTransform{}, symbolCallError("CGAffineTransformMakeWithComponents", "13.0", _cGAffineTransformMakeWithComponentsErr)
	}
	return _cGAffineTransformMakeWithComponents(components), nil
}

// CGAffineTransformMakeWithComponents.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformMakeWithComponents
func CGAffineTransformMakeWithComponents(components uintptr) corefoundation.CGAffineTransform {
	result, callErr := tryCGAffineTransformMakeWithComponents(components)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGAffineTransformRotate func(t corefoundation.CGAffineTransform, angle float64) corefoundation.CGAffineTransform
var _cGAffineTransformRotateErr error

func tryCGAffineTransformRotate(t corefoundation.CGAffineTransform, angle float64) (corefoundation.CGAffineTransform, error) {
	if _cGAffineTransformRotate == nil {
		return corefoundation.CGAffineTransform{}, symbolCallError("CGAffineTransformRotate", "10.0", _cGAffineTransformRotateErr)
	}
	return _cGAffineTransformRotate(t, angle), nil
}

// CGAffineTransformRotate returns an affine transformation matrix constructed by rotating an existing affine transform.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformRotate(_:_:)
func CGAffineTransformRotate(t corefoundation.CGAffineTransform, angle float64) corefoundation.CGAffineTransform {
	result, callErr := tryCGAffineTransformRotate(t, angle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGAffineTransformScale func(t corefoundation.CGAffineTransform, sx float64, sy float64) corefoundation.CGAffineTransform
var _cGAffineTransformScaleErr error

func tryCGAffineTransformScale(t corefoundation.CGAffineTransform, sx float64, sy float64) (corefoundation.CGAffineTransform, error) {
	if _cGAffineTransformScale == nil {
		return corefoundation.CGAffineTransform{}, symbolCallError("CGAffineTransformScale", "10.0", _cGAffineTransformScaleErr)
	}
	return _cGAffineTransformScale(t, sx, sy), nil
}

// CGAffineTransformScale returns an affine transformation matrix constructed by scaling an existing affine transform.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformScale(_:_:_:)
func CGAffineTransformScale(t corefoundation.CGAffineTransform, sx float64, sy float64) corefoundation.CGAffineTransform {
	result, callErr := tryCGAffineTransformScale(t, sx, sy)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGAffineTransformTranslate func(t corefoundation.CGAffineTransform, tx float64, ty float64) corefoundation.CGAffineTransform
var _cGAffineTransformTranslateErr error

func tryCGAffineTransformTranslate(t corefoundation.CGAffineTransform, tx float64, ty float64) (corefoundation.CGAffineTransform, error) {
	if _cGAffineTransformTranslate == nil {
		return corefoundation.CGAffineTransform{}, symbolCallError("CGAffineTransformTranslate", "10.0", _cGAffineTransformTranslateErr)
	}
	return _cGAffineTransformTranslate(t, tx, ty), nil
}

// CGAffineTransformTranslate returns an affine transformation matrix constructed by translating an existing affine transform.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformTranslate(_:_:_:)
func CGAffineTransformTranslate(t corefoundation.CGAffineTransform, tx float64, ty float64) corefoundation.CGAffineTransform {
	result, callErr := tryCGAffineTransformTranslate(t, tx, ty)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGAssociateMouseAndMouseCursorPosition func(connected bool) CGError
var _cGAssociateMouseAndMouseCursorPositionErr error

func tryCGAssociateMouseAndMouseCursorPosition(connected bool) (CGError, error) {
	if _cGAssociateMouseAndMouseCursorPosition == nil {
		return *new(CGError), symbolCallError("CGAssociateMouseAndMouseCursorPosition", "10.0", _cGAssociateMouseAndMouseCursorPositionErr)
	}
	return _cGAssociateMouseAndMouseCursorPosition(connected), nil
}

// CGAssociateMouseAndMouseCursorPosition connects or disconnects the mouse and cursor while an application is in the foreground.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGAssociateMouseAndMouseCursorPosition(_:)
func CGAssociateMouseAndMouseCursorPosition(connected bool) CGError {
	result, callErr := tryCGAssociateMouseAndMouseCursorPosition(connected)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGBeginDisplayConfiguration func(config *CGDisplayConfigRef) CGError
var _cGBeginDisplayConfigurationErr error

func tryCGBeginDisplayConfiguration(config *CGDisplayConfigRef) (CGError, error) {
	if _cGBeginDisplayConfiguration == nil {
		return *new(CGError), symbolCallError("CGBeginDisplayConfiguration", "10.0", _cGBeginDisplayConfigurationErr)
	}
	return _cGBeginDisplayConfiguration(config), nil
}

// CGBeginDisplayConfiguration begins a new set of display configuration changes.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGBeginDisplayConfiguration(_:)
func CGBeginDisplayConfiguration(config *CGDisplayConfigRef) CGError {
	result, callErr := tryCGBeginDisplayConfiguration(config)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGBitmapContextCreate func(data unsafe.Pointer, width uintptr, height uintptr, bitsPerComponent uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo) CGContextRef
var _cGBitmapContextCreateErr error

func tryCGBitmapContextCreate(data unsafe.Pointer, width uintptr, height uintptr, bitsPerComponent uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo) (CGContextRef, error) {
	if _cGBitmapContextCreate == nil {
		return 0, symbolCallError("CGBitmapContextCreate", "10.0", _cGBitmapContextCreateErr)
	}
	return _cGBitmapContextCreate(data, width, height, bitsPerComponent, bytesPerRow, space, bitmapInfo), nil
}

// CGBitmapContextCreate.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/init(data:width:height:bitsPerComponent:bytesPerRow:space:bitmapInfo:)-10b3i
func CGBitmapContextCreate(data unsafe.Pointer, width uintptr, height uintptr, bitsPerComponent uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo) CGContextRef {
	result, callErr := tryCGBitmapContextCreate(data, width, height, bitsPerComponent, bytesPerRow, space, bitmapInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGBitmapContextCreateAdaptive func(width uintptr, height uintptr, auxiliaryInfo corefoundation.CFDictionaryRef, onResolve bool) CGContextRef
var _cGBitmapContextCreateAdaptiveErr error

func tryCGBitmapContextCreateAdaptive(width uintptr, height uintptr, auxiliaryInfo corefoundation.CFDictionaryRef, onResolve bool) (CGContextRef, error) {
	if _cGBitmapContextCreateAdaptive == nil {
		return 0, symbolCallError("CGBitmapContextCreateAdaptive", "26.0", _cGBitmapContextCreateAdaptiveErr)
	}
	return _cGBitmapContextCreateAdaptive(width, height, auxiliaryInfo, onResolve), nil
}

// CGBitmapContextCreateAdaptive.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGBitmapContextCreateAdaptive
func CGBitmapContextCreateAdaptive(width uintptr, height uintptr, auxiliaryInfo corefoundation.CFDictionaryRef, onResolve bool) CGContextRef {
	result, callErr := tryCGBitmapContextCreateAdaptive(width, height, auxiliaryInfo, onResolve)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGBitmapContextCreateImage func(context CGContextRef) CGImageRef
var _cGBitmapContextCreateImageErr error

func tryCGBitmapContextCreateImage(context CGContextRef) (CGImageRef, error) {
	if _cGBitmapContextCreateImage == nil {
		return 0, symbolCallError("CGBitmapContextCreateImage", "10.4", _cGBitmapContextCreateImageErr)
	}
	return _cGBitmapContextCreateImage(context), nil
}

// CGBitmapContextCreateImage creates and returns a CGImage from the pixel data in a bitmap graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/makeImage()
func CGBitmapContextCreateImage(context CGContextRef) CGImageRef {
	result, callErr := tryCGBitmapContextCreateImage(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGBitmapContextCreateWithData func(data unsafe.Pointer, width uintptr, height uintptr, bitsPerComponent uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo, releaseCallback CGBitmapContextReleaseDataCallback, releaseInfo unsafe.Pointer) CGContextRef
var _cGBitmapContextCreateWithDataErr error

func tryCGBitmapContextCreateWithData(data unsafe.Pointer, width uintptr, height uintptr, bitsPerComponent uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo, releaseCallback CGBitmapContextReleaseDataCallback, releaseInfo unsafe.Pointer) (CGContextRef, error) {
	if _cGBitmapContextCreateWithData == nil {
		return 0, symbolCallError("CGBitmapContextCreateWithData", "10.6", _cGBitmapContextCreateWithDataErr)
	}
	return _cGBitmapContextCreateWithData(data, width, height, bitsPerComponent, bytesPerRow, space, bitmapInfo, releaseCallback, releaseInfo), nil
}

// CGBitmapContextCreateWithData.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/init(data:width:height:bitsPerComponent:bytesPerRow:space:bitmapInfo:releaseCallback:releaseInfo:)-4yzt5
func CGBitmapContextCreateWithData(data unsafe.Pointer, width uintptr, height uintptr, bitsPerComponent uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo, releaseCallback CGBitmapContextReleaseDataCallback, releaseInfo unsafe.Pointer) CGContextRef {
	result, callErr := tryCGBitmapContextCreateWithData(data, width, height, bitsPerComponent, bytesPerRow, space, bitmapInfo, releaseCallback, releaseInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGBitmapContextGetAlphaInfo func(context CGContextRef) CGImageAlphaInfo
var _cGBitmapContextGetAlphaInfoErr error

func tryCGBitmapContextGetAlphaInfo(context CGContextRef) (CGImageAlphaInfo, error) {
	if _cGBitmapContextGetAlphaInfo == nil {
		return *new(CGImageAlphaInfo), symbolCallError("CGBitmapContextGetAlphaInfo", "10.2", _cGBitmapContextGetAlphaInfoErr)
	}
	return _cGBitmapContextGetAlphaInfo(context), nil
}

// CGBitmapContextGetAlphaInfo returns the alpha information associated with the context, which indicates how a bitmap context handles the alpha component.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/alphaInfo
func CGBitmapContextGetAlphaInfo(context CGContextRef) CGImageAlphaInfo {
	result, callErr := tryCGBitmapContextGetAlphaInfo(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGBitmapContextGetBitmapInfo func(context CGContextRef) CGBitmapInfo
var _cGBitmapContextGetBitmapInfoErr error

func tryCGBitmapContextGetBitmapInfo(context CGContextRef) (CGBitmapInfo, error) {
	if _cGBitmapContextGetBitmapInfo == nil {
		return *new(CGBitmapInfo), symbolCallError("CGBitmapContextGetBitmapInfo", "10.4", _cGBitmapContextGetBitmapInfoErr)
	}
	return _cGBitmapContextGetBitmapInfo(context), nil
}

// CGBitmapContextGetBitmapInfo obtains the bitmap information associated with a bitmap graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/bitmapInfo
func CGBitmapContextGetBitmapInfo(context CGContextRef) CGBitmapInfo {
	result, callErr := tryCGBitmapContextGetBitmapInfo(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGBitmapContextGetBitsPerComponent func(context CGContextRef) uintptr
var _cGBitmapContextGetBitsPerComponentErr error

func tryCGBitmapContextGetBitsPerComponent(context CGContextRef) (uintptr, error) {
	if _cGBitmapContextGetBitsPerComponent == nil {
		return 0, symbolCallError("CGBitmapContextGetBitsPerComponent", "10.2", _cGBitmapContextGetBitsPerComponentErr)
	}
	return _cGBitmapContextGetBitsPerComponent(context), nil
}

// CGBitmapContextGetBitsPerComponent returns the bits per component of a bitmap context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/bitsPerComponent
func CGBitmapContextGetBitsPerComponent(context CGContextRef) uintptr {
	result, callErr := tryCGBitmapContextGetBitsPerComponent(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGBitmapContextGetBitsPerPixel func(context CGContextRef) uintptr
var _cGBitmapContextGetBitsPerPixelErr error

func tryCGBitmapContextGetBitsPerPixel(context CGContextRef) (uintptr, error) {
	if _cGBitmapContextGetBitsPerPixel == nil {
		return 0, symbolCallError("CGBitmapContextGetBitsPerPixel", "10.2", _cGBitmapContextGetBitsPerPixelErr)
	}
	return _cGBitmapContextGetBitsPerPixel(context), nil
}

// CGBitmapContextGetBitsPerPixel returns the bits per pixel of a bitmap context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/bitsPerPixel
func CGBitmapContextGetBitsPerPixel(context CGContextRef) uintptr {
	result, callErr := tryCGBitmapContextGetBitsPerPixel(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGBitmapContextGetBytesPerRow func(context CGContextRef) uintptr
var _cGBitmapContextGetBytesPerRowErr error

func tryCGBitmapContextGetBytesPerRow(context CGContextRef) (uintptr, error) {
	if _cGBitmapContextGetBytesPerRow == nil {
		return 0, symbolCallError("CGBitmapContextGetBytesPerRow", "10.2", _cGBitmapContextGetBytesPerRowErr)
	}
	return _cGBitmapContextGetBytesPerRow(context), nil
}

// CGBitmapContextGetBytesPerRow returns the bytes per row of a bitmap context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/bytesPerRow
func CGBitmapContextGetBytesPerRow(context CGContextRef) uintptr {
	result, callErr := tryCGBitmapContextGetBytesPerRow(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGBitmapContextGetColorSpace func(context CGContextRef) CGColorSpaceRef
var _cGBitmapContextGetColorSpaceErr error

func tryCGBitmapContextGetColorSpace(context CGContextRef) (CGColorSpaceRef, error) {
	if _cGBitmapContextGetColorSpace == nil {
		return 0, symbolCallError("CGBitmapContextGetColorSpace", "10.2", _cGBitmapContextGetColorSpaceErr)
	}
	return _cGBitmapContextGetColorSpace(context), nil
}

// CGBitmapContextGetColorSpace returns the color space of a bitmap context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/colorSpace
func CGBitmapContextGetColorSpace(context CGContextRef) CGColorSpaceRef {
	result, callErr := tryCGBitmapContextGetColorSpace(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGBitmapContextGetData func(context CGContextRef) unsafe.Pointer
var _cGBitmapContextGetDataErr error

func tryCGBitmapContextGetData(context CGContextRef) (unsafe.Pointer, error) {
	if _cGBitmapContextGetData == nil {
		return nil, symbolCallError("CGBitmapContextGetData", "10.2", _cGBitmapContextGetDataErr)
	}
	return _cGBitmapContextGetData(context), nil
}

// CGBitmapContextGetData returns a pointer to the image data associated with a bitmap context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/data
func CGBitmapContextGetData(context CGContextRef) unsafe.Pointer {
	result, callErr := tryCGBitmapContextGetData(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGBitmapContextGetHeight func(context CGContextRef) uintptr
var _cGBitmapContextGetHeightErr error

func tryCGBitmapContextGetHeight(context CGContextRef) (uintptr, error) {
	if _cGBitmapContextGetHeight == nil {
		return 0, symbolCallError("CGBitmapContextGetHeight", "10.2", _cGBitmapContextGetHeightErr)
	}
	return _cGBitmapContextGetHeight(context), nil
}

// CGBitmapContextGetHeight returns the height in pixels of a bitmap context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/height
func CGBitmapContextGetHeight(context CGContextRef) uintptr {
	result, callErr := tryCGBitmapContextGetHeight(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGBitmapContextGetWidth func(context CGContextRef) uintptr
var _cGBitmapContextGetWidthErr error

func tryCGBitmapContextGetWidth(context CGContextRef) (uintptr, error) {
	if _cGBitmapContextGetWidth == nil {
		return 0, symbolCallError("CGBitmapContextGetWidth", "10.2", _cGBitmapContextGetWidthErr)
	}
	return _cGBitmapContextGetWidth(context), nil
}

// CGBitmapContextGetWidth returns the width in pixels of a bitmap context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/width
func CGBitmapContextGetWidth(context CGContextRef) uintptr {
	result, callErr := tryCGBitmapContextGetWidth(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGCancelDisplayConfiguration func(config CGDisplayConfigRef) CGError
var _cGCancelDisplayConfigurationErr error

func tryCGCancelDisplayConfiguration(config CGDisplayConfigRef) (CGError, error) {
	if _cGCancelDisplayConfiguration == nil {
		return *new(CGError), symbolCallError("CGCancelDisplayConfiguration", "10.0", _cGCancelDisplayConfigurationErr)
	}
	return _cGCancelDisplayConfiguration(config), nil
}

// CGCancelDisplayConfiguration cancels a set of display configuration changes.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGCancelDisplayConfiguration(_:)
func CGCancelDisplayConfiguration(config CGDisplayConfigRef) CGError {
	result, callErr := tryCGCancelDisplayConfiguration(config)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGCaptureAllDisplays func() CGError
var _cGCaptureAllDisplaysErr error

func tryCGCaptureAllDisplays() (CGError, error) {
	if _cGCaptureAllDisplays == nil {
		return *new(CGError), symbolCallError("CGCaptureAllDisplays", "10.0", _cGCaptureAllDisplaysErr)
	}
	return _cGCaptureAllDisplays(), nil
}

// CGCaptureAllDisplays obtains exclusive use of all active displays, preventing other applications and system services from using the display or changing its configuration.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGCaptureAllDisplays()
func CGCaptureAllDisplays() CGError {
	result, callErr := tryCGCaptureAllDisplays()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGCaptureAllDisplaysWithOptions func(options CGCaptureOptions) CGError
var _cGCaptureAllDisplaysWithOptionsErr error

func tryCGCaptureAllDisplaysWithOptions(options CGCaptureOptions) (CGError, error) {
	if _cGCaptureAllDisplaysWithOptions == nil {
		return *new(CGError), symbolCallError("CGCaptureAllDisplaysWithOptions", "10.3", _cGCaptureAllDisplaysWithOptionsErr)
	}
	return _cGCaptureAllDisplaysWithOptions(options), nil
}

// CGCaptureAllDisplaysWithOptions captures all attached displays, using the specified options.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGCaptureAllDisplaysWithOptions(_:)
func CGCaptureAllDisplaysWithOptions(options CGCaptureOptions) CGError {
	result, callErr := tryCGCaptureAllDisplaysWithOptions(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorConversionInfoConvertData func(info CGColorConversionInfoRef, width uintptr, height uintptr, dst_data unsafe.Pointer, dst_format CGColorBufferFormat, src_data unsafe.Pointer, src_format CGColorBufferFormat, options corefoundation.CFDictionaryRef) bool
var _cGColorConversionInfoConvertDataErr error

func tryCGColorConversionInfoConvertData(info CGColorConversionInfoRef, width uintptr, height uintptr, dst_data unsafe.Pointer, dst_format CGColorBufferFormat, src_data unsafe.Pointer, src_format CGColorBufferFormat, options corefoundation.CFDictionaryRef) (bool, error) {
	if _cGColorConversionInfoConvertData == nil {
		return false, symbolCallError("CGColorConversionInfoConvertData", "15.0", _cGColorConversionInfoConvertDataErr)
	}
	return _cGColorConversionInfoConvertData(info, width, height, dst_data, dst_format, src_data, src_format, options), nil
}

// CGColorConversionInfoConvertData.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfo/convert(width:height:to:format:from:format:options:)
func CGColorConversionInfoConvertData(info CGColorConversionInfoRef, width uintptr, height uintptr, dst_data unsafe.Pointer, dst_format CGColorBufferFormat, src_data unsafe.Pointer, src_format CGColorBufferFormat, options corefoundation.CFDictionaryRef) bool {
	result, callErr := tryCGColorConversionInfoConvertData(info, width, height, dst_data, dst_format, src_data, src_format, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorConversionInfoCreate func(src CGColorSpaceRef, dst CGColorSpaceRef) CGColorConversionInfoRef
var _cGColorConversionInfoCreateErr error

func tryCGColorConversionInfoCreate(src CGColorSpaceRef, dst CGColorSpaceRef) (CGColorConversionInfoRef, error) {
	if _cGColorConversionInfoCreate == nil {
		return 0, symbolCallError("CGColorConversionInfoCreate", "10.12", _cGColorConversionInfoCreateErr)
	}
	return _cGColorConversionInfoCreate(src, dst), nil
}

// CGColorConversionInfoCreate creates a conversion between two specified color spaces.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfo/init(src:dst:)
func CGColorConversionInfoCreate(src CGColorSpaceRef, dst CGColorSpaceRef) CGColorConversionInfoRef {
	result, callErr := tryCGColorConversionInfoCreate(src, dst)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorConversionInfoCreateForToneMapping func(from CGColorSpaceRef, source_headroom float32, to CGColorSpaceRef, target_headroom float32, method CGToneMapping, options corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) CGColorConversionInfoRef
var _cGColorConversionInfoCreateForToneMappingErr error

func tryCGColorConversionInfoCreateForToneMapping(from CGColorSpaceRef, source_headroom float32, to CGColorSpaceRef, target_headroom float32, method CGToneMapping, options corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) (CGColorConversionInfoRef, error) {
	if _cGColorConversionInfoCreateForToneMapping == nil {
		return 0, symbolCallError("CGColorConversionInfoCreateForToneMapping", "15.0", _cGColorConversionInfoCreateForToneMappingErr)
	}
	return _cGColorConversionInfoCreateForToneMapping(from, source_headroom, to, target_headroom, method, options, err), nil
}

// CGColorConversionInfoCreateForToneMapping.
//
// Deprecated: declared Swift name 'init(src:srcHeadroom:dst:dstHeadroom:toneMapping:options:)' was adjusted to 'init(src:srcHeadroom:dst:dstHeadroom:toneMapping:options:_:)' because it does not have the correct number of parameters (6 vs. 7); please report this to its maintainer
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfo/init(src:srcHeadroom:dst:dstHeadroom:toneMapping:options:_:)
func CGColorConversionInfoCreateForToneMapping(from CGColorSpaceRef, source_headroom float32, to CGColorSpaceRef, target_headroom float32, method CGToneMapping, options corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) CGColorConversionInfoRef {
	result, callErr := tryCGColorConversionInfoCreateForToneMapping(from, source_headroom, to, target_headroom, method, options, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorConversionInfoCreateFromList func(options corefoundation.CFDictionaryRef, arg1 CGColorSpaceRef, arg2 CGColorConversionInfoTransformType, arg3 CGColorRenderingIntent) CGColorConversionInfoRef
var _cGColorConversionInfoCreateFromListErr error

func tryCGColorConversionInfoCreateFromList(options corefoundation.CFDictionaryRef, arg1 CGColorSpaceRef, arg2 CGColorConversionInfoTransformType, arg3 CGColorRenderingIntent) (CGColorConversionInfoRef, error) {
	if _cGColorConversionInfoCreateFromList == nil {
		return 0, symbolCallError("CGColorConversionInfoCreateFromList", "10.12", _cGColorConversionInfoCreateFromListErr)
	}
	return _cGColorConversionInfoCreateFromList(options, arg1, arg2, arg3), nil
}

// CGColorConversionInfoCreateFromList creates a conversion between an arbitrary number of specified color spaces.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfoCreateFromList
func CGColorConversionInfoCreateFromList(options corefoundation.CFDictionaryRef, arg1 CGColorSpaceRef, arg2 CGColorConversionInfoTransformType, arg3 CGColorRenderingIntent) CGColorConversionInfoRef {
	result, callErr := tryCGColorConversionInfoCreateFromList(options, arg1, arg2, arg3)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorConversionInfoCreateFromListWithArguments func(options corefoundation.CFDictionaryRef, arg1 CGColorSpaceRef, arg2 CGColorConversionInfoTransformType, arg3 CGColorRenderingIntent, arg4 kernel.Va_list) CGColorConversionInfoRef
var _cGColorConversionInfoCreateFromListWithArgumentsErr error

func tryCGColorConversionInfoCreateFromListWithArguments(options corefoundation.CFDictionaryRef, arg1 CGColorSpaceRef, arg2 CGColorConversionInfoTransformType, arg3 CGColorRenderingIntent, arg4 kernel.Va_list) (CGColorConversionInfoRef, error) {
	if _cGColorConversionInfoCreateFromListWithArguments == nil {
		return 0, symbolCallError("CGColorConversionInfoCreateFromListWithArguments", "10.13", _cGColorConversionInfoCreateFromListWithArgumentsErr)
	}
	return _cGColorConversionInfoCreateFromListWithArguments(options, arg1, arg2, arg3, arg4), nil
}

// CGColorConversionInfoCreateFromListWithArguments.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfoCreateFromListWithArguments
func CGColorConversionInfoCreateFromListWithArguments(options corefoundation.CFDictionaryRef, arg1 CGColorSpaceRef, arg2 CGColorConversionInfoTransformType, arg3 CGColorRenderingIntent, arg4 kernel.Va_list) CGColorConversionInfoRef {
	result, callErr := tryCGColorConversionInfoCreateFromListWithArguments(options, arg1, arg2, arg3, arg4)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorConversionInfoCreateWithOptions func(src CGColorSpaceRef, dst CGColorSpaceRef, options corefoundation.CFDictionaryRef) CGColorConversionInfoRef
var _cGColorConversionInfoCreateWithOptionsErr error

func tryCGColorConversionInfoCreateWithOptions(src CGColorSpaceRef, dst CGColorSpaceRef, options corefoundation.CFDictionaryRef) (CGColorConversionInfoRef, error) {
	if _cGColorConversionInfoCreateWithOptions == nil {
		return 0, symbolCallError("CGColorConversionInfoCreateWithOptions", "10.14.6", _cGColorConversionInfoCreateWithOptionsErr)
	}
	return _cGColorConversionInfoCreateWithOptions(src, dst, options), nil
}

// CGColorConversionInfoCreateWithOptions.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfo/init(optionsSrc:dst:options:)
func CGColorConversionInfoCreateWithOptions(src CGColorSpaceRef, dst CGColorSpaceRef, options corefoundation.CFDictionaryRef) CGColorConversionInfoRef {
	result, callErr := tryCGColorConversionInfoCreateWithOptions(src, dst, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorConversionInfoGetTypeID func() uint
var _cGColorConversionInfoGetTypeIDErr error

func tryCGColorConversionInfoGetTypeID() (uint, error) {
	if _cGColorConversionInfoGetTypeID == nil {
		return 0, symbolCallError("CGColorConversionInfoGetTypeID", "", _cGColorConversionInfoGetTypeIDErr)
	}
	return _cGColorConversionInfoGetTypeID(), nil
}

// CGColorConversionInfoGetTypeID returns the Core Foundation type identifier for a color conversion info data type.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfo/typeID
func CGColorConversionInfoGetTypeID() uint {
	result, callErr := tryCGColorConversionInfoGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorCreate func(space CGColorSpaceRef, components *float64) CGColorRef
var _cGColorCreateErr error

func tryCGColorCreate(space CGColorSpaceRef, components *float64) (CGColorRef, error) {
	if _cGColorCreate == nil {
		return 0, symbolCallError("CGColorCreate", "10.3", _cGColorCreateErr)
	}
	return _cGColorCreate(space, components), nil
}

// CGColorCreate creates a color using a list of intensity values (including alpha) and an associated color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(colorSpace:components:)
func CGColorCreate(space CGColorSpaceRef, components *float64) CGColorRef {
	result, callErr := tryCGColorCreate(space, components)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorCreateCopy func(color CGColorRef) CGColorRef
var _cGColorCreateCopyErr error

func tryCGColorCreateCopy(color CGColorRef) (CGColorRef, error) {
	if _cGColorCreateCopy == nil {
		return 0, symbolCallError("CGColorCreateCopy", "10.3", _cGColorCreateCopyErr)
	}
	return _cGColorCreateCopy(color), nil
}

// CGColorCreateCopy creates a copy of an existing color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/copy()
func CGColorCreateCopy(color CGColorRef) CGColorRef {
	result, callErr := tryCGColorCreateCopy(color)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorCreateCopyByMatchingToColorSpace func(arg0 CGColorSpaceRef, intent CGColorRenderingIntent, color CGColorRef, options corefoundation.CFDictionaryRef) CGColorRef
var _cGColorCreateCopyByMatchingToColorSpaceErr error

func tryCGColorCreateCopyByMatchingToColorSpace(arg0 CGColorSpaceRef, intent CGColorRenderingIntent, color CGColorRef, options corefoundation.CFDictionaryRef) (CGColorRef, error) {
	if _cGColorCreateCopyByMatchingToColorSpace == nil {
		return 0, symbolCallError("CGColorCreateCopyByMatchingToColorSpace", "10.11", _cGColorCreateCopyByMatchingToColorSpaceErr)
	}
	return _cGColorCreateCopyByMatchingToColorSpace(arg0, intent, color, options), nil
}

// CGColorCreateCopyByMatchingToColorSpace creates a new color in a different color space that matches the provided color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/converted(to:intent:options:)
func CGColorCreateCopyByMatchingToColorSpace(arg0 CGColorSpaceRef, intent CGColorRenderingIntent, color CGColorRef, options corefoundation.CFDictionaryRef) CGColorRef {
	result, callErr := tryCGColorCreateCopyByMatchingToColorSpace(arg0, intent, color, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorCreateCopyWithAlpha func(color CGColorRef, alpha float64) CGColorRef
var _cGColorCreateCopyWithAlphaErr error

func tryCGColorCreateCopyWithAlpha(color CGColorRef, alpha float64) (CGColorRef, error) {
	if _cGColorCreateCopyWithAlpha == nil {
		return 0, symbolCallError("CGColorCreateCopyWithAlpha", "10.3", _cGColorCreateCopyWithAlphaErr)
	}
	return _cGColorCreateCopyWithAlpha(color, alpha), nil
}

// CGColorCreateCopyWithAlpha creates a copy of an existing color, substituting a new alpha value.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/copy(alpha:)
func CGColorCreateCopyWithAlpha(color CGColorRef, alpha float64) CGColorRef {
	result, callErr := tryCGColorCreateCopyWithAlpha(color, alpha)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorCreateGenericCMYK func(cyan float64, magenta float64, yellow float64, black float64, alpha float64) CGColorRef
var _cGColorCreateGenericCMYKErr error

func tryCGColorCreateGenericCMYK(cyan float64, magenta float64, yellow float64, black float64, alpha float64) (CGColorRef, error) {
	if _cGColorCreateGenericCMYK == nil {
		return 0, symbolCallError("CGColorCreateGenericCMYK", "10.5", _cGColorCreateGenericCMYKErr)
	}
	return _cGColorCreateGenericCMYK(cyan, magenta, yellow, black, alpha), nil
}

// CGColorCreateGenericCMYK creates a color in the Generic CMYK color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(genericCMYKCyan:magenta:yellow:black:alpha:)
func CGColorCreateGenericCMYK(cyan float64, magenta float64, yellow float64, black float64, alpha float64) CGColorRef {
	result, callErr := tryCGColorCreateGenericCMYK(cyan, magenta, yellow, black, alpha)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorCreateGenericGray func(gray float64, alpha float64) CGColorRef
var _cGColorCreateGenericGrayErr error

func tryCGColorCreateGenericGray(gray float64, alpha float64) (CGColorRef, error) {
	if _cGColorCreateGenericGray == nil {
		return 0, symbolCallError("CGColorCreateGenericGray", "10.5", _cGColorCreateGenericGrayErr)
	}
	return _cGColorCreateGenericGray(gray, alpha), nil
}

// CGColorCreateGenericGray creates a color in the Generic gray color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(gray:alpha:)
func CGColorCreateGenericGray(gray float64, alpha float64) CGColorRef {
	result, callErr := tryCGColorCreateGenericGray(gray, alpha)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorCreateGenericGrayGamma2_2 func(gray float64, alpha float64) CGColorRef
var _cGColorCreateGenericGrayGamma2_2Err error

func tryCGColorCreateGenericGrayGamma2_2(gray float64, alpha float64) (CGColorRef, error) {
	if _cGColorCreateGenericGrayGamma2_2 == nil {
		return 0, symbolCallError("CGColorCreateGenericGrayGamma2_2", "10.15", _cGColorCreateGenericGrayGamma2_2Err)
	}
	return _cGColorCreateGenericGrayGamma2_2(gray, alpha), nil
}

// CGColorCreateGenericGrayGamma2_2 creates a color in the Generic gray color space with a gamma ramp of 2.2.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(genericGrayGamma2_2Gray:alpha:)
func CGColorCreateGenericGrayGamma2_2(gray float64, alpha float64) CGColorRef {
	result, callErr := tryCGColorCreateGenericGrayGamma2_2(gray, alpha)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorCreateGenericRGB func(red float64, green float64, blue float64, alpha float64) CGColorRef
var _cGColorCreateGenericRGBErr error

func tryCGColorCreateGenericRGB(red float64, green float64, blue float64, alpha float64) (CGColorRef, error) {
	if _cGColorCreateGenericRGB == nil {
		return 0, symbolCallError("CGColorCreateGenericRGB", "10.5", _cGColorCreateGenericRGBErr)
	}
	return _cGColorCreateGenericRGB(red, green, blue, alpha), nil
}

// CGColorCreateGenericRGB creates a color in the Generic RGB color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(red:green:blue:alpha:)
func CGColorCreateGenericRGB(red float64, green float64, blue float64, alpha float64) CGColorRef {
	result, callErr := tryCGColorCreateGenericRGB(red, green, blue, alpha)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorCreateSRGB func(red float64, green float64, blue float64, alpha float64) CGColorRef
var _cGColorCreateSRGBErr error

func tryCGColorCreateSRGB(red float64, green float64, blue float64, alpha float64) (CGColorRef, error) {
	if _cGColorCreateSRGB == nil {
		return 0, symbolCallError("CGColorCreateSRGB", "10.15", _cGColorCreateSRGBErr)
	}
	return _cGColorCreateSRGB(red, green, blue, alpha), nil
}

// CGColorCreateSRGB creates a color in the sRGB color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(srgbRed:green:blue:alpha:)
func CGColorCreateSRGB(red float64, green float64, blue float64, alpha float64) CGColorRef {
	result, callErr := tryCGColorCreateSRGB(red, green, blue, alpha)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorCreateWithContentHeadroom func(headroom float32, space CGColorSpaceRef, red float64, green float64, blue float64, alpha float64) CGColorRef
var _cGColorCreateWithContentHeadroomErr error

func tryCGColorCreateWithContentHeadroom(headroom float32, space CGColorSpaceRef, red float64, green float64, blue float64, alpha float64) (CGColorRef, error) {
	if _cGColorCreateWithContentHeadroom == nil {
		return 0, symbolCallError("CGColorCreateWithContentHeadroom", "26.0", _cGColorCreateWithContentHeadroomErr)
	}
	return _cGColorCreateWithContentHeadroom(headroom, space, red, green, blue, alpha), nil
}

// CGColorCreateWithContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(headroom:colorSpace:red:green:blue:alpha:)
func CGColorCreateWithContentHeadroom(headroom float32, space CGColorSpaceRef, red float64, green float64, blue float64, alpha float64) CGColorRef {
	result, callErr := tryCGColorCreateWithContentHeadroom(headroom, space, red, green, blue, alpha)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorCreateWithPattern func(space CGColorSpaceRef, pattern CGPatternRef, components *float64) CGColorRef
var _cGColorCreateWithPatternErr error

func tryCGColorCreateWithPattern(space CGColorSpaceRef, pattern CGPatternRef, components *float64) (CGColorRef, error) {
	if _cGColorCreateWithPattern == nil {
		return 0, symbolCallError("CGColorCreateWithPattern", "10.3", _cGColorCreateWithPatternErr)
	}
	return _cGColorCreateWithPattern(space, pattern, components), nil
}

// CGColorCreateWithPattern creates a color using a list of intensity values (including alpha), a pattern color space, and a pattern.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/init(patternSpace:pattern:components:)
func CGColorCreateWithPattern(space CGColorSpaceRef, pattern CGPatternRef, components *float64) CGColorRef {
	result, callErr := tryCGColorCreateWithPattern(space, pattern, components)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorEqualToColor func(color1 CGColorRef, color2 CGColorRef) bool
var _cGColorEqualToColorErr error

func tryCGColorEqualToColor(color1 CGColorRef, color2 CGColorRef) (bool, error) {
	if _cGColorEqualToColor == nil {
		return false, symbolCallError("CGColorEqualToColor", "10.3", _cGColorEqualToColorErr)
	}
	return _cGColorEqualToColor(color1, color2), nil
}

// CGColorEqualToColor indicates whether two colors are equal.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorEqualToColor
func CGColorEqualToColor(color1 CGColorRef, color2 CGColorRef) bool {
	result, callErr := tryCGColorEqualToColor(color1, color2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorGetAlpha func(color CGColorRef) float64
var _cGColorGetAlphaErr error

func tryCGColorGetAlpha(color CGColorRef) (float64, error) {
	if _cGColorGetAlpha == nil {
		return 0.0, symbolCallError("CGColorGetAlpha", "10.3", _cGColorGetAlphaErr)
	}
	return _cGColorGetAlpha(color), nil
}

// CGColorGetAlpha returns the value of the alpha component associated with a color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/alpha
func CGColorGetAlpha(color CGColorRef) float64 {
	result, callErr := tryCGColorGetAlpha(color)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorGetColorSpace func(color CGColorRef) CGColorSpaceRef
var _cGColorGetColorSpaceErr error

func tryCGColorGetColorSpace(color CGColorRef) (CGColorSpaceRef, error) {
	if _cGColorGetColorSpace == nil {
		return 0, symbolCallError("CGColorGetColorSpace", "10.3", _cGColorGetColorSpaceErr)
	}
	return _cGColorGetColorSpace(color), nil
}

// CGColorGetColorSpace returns the color space associated with a color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/colorSpace
func CGColorGetColorSpace(color CGColorRef) CGColorSpaceRef {
	result, callErr := tryCGColorGetColorSpace(color)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorGetComponents func(color CGColorRef) *float64
var _cGColorGetComponentsErr error

func tryCGColorGetComponents(color CGColorRef) (*float64, error) {
	if _cGColorGetComponents == nil {
		return nil, symbolCallError("CGColorGetComponents", "10.3", _cGColorGetComponentsErr)
	}
	return _cGColorGetComponents(color), nil
}

// CGColorGetComponents returns the values of the color components (including alpha) associated with a color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorGetComponents
func CGColorGetComponents(color CGColorRef) *float64 {
	result, callErr := tryCGColorGetComponents(color)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorGetConstantColor func(colorName corefoundation.CFStringRef) CGColorRef
var _cGColorGetConstantColorErr error

func tryCGColorGetConstantColor(colorName corefoundation.CFStringRef) (CGColorRef, error) {
	if _cGColorGetConstantColor == nil {
		return 0, symbolCallError("CGColorGetConstantColor", "10.5", _cGColorGetConstantColorErr)
	}
	return _cGColorGetConstantColor(colorName), nil
}

// CGColorGetConstantColor returns a color object that represents a constant color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorGetConstantColor
func CGColorGetConstantColor(colorName corefoundation.CFStringRef) CGColorRef {
	result, callErr := tryCGColorGetConstantColor(colorName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorGetContentHeadroom func(color CGColorRef) float32
var _cGColorGetContentHeadroomErr error

func tryCGColorGetContentHeadroom(color CGColorRef) (float32, error) {
	if _cGColorGetContentHeadroom == nil {
		return 0.0, symbolCallError("CGColorGetContentHeadroom", "26.0", _cGColorGetContentHeadroomErr)
	}
	return _cGColorGetContentHeadroom(color), nil
}

// CGColorGetContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/contentHeadroom
func CGColorGetContentHeadroom(color CGColorRef) float32 {
	result, callErr := tryCGColorGetContentHeadroom(color)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorGetNumberOfComponents func(color CGColorRef) uintptr
var _cGColorGetNumberOfComponentsErr error

func tryCGColorGetNumberOfComponents(color CGColorRef) (uintptr, error) {
	if _cGColorGetNumberOfComponents == nil {
		return 0, symbolCallError("CGColorGetNumberOfComponents", "10.3", _cGColorGetNumberOfComponentsErr)
	}
	return _cGColorGetNumberOfComponents(color), nil
}

// CGColorGetNumberOfComponents returns the number of color components (including alpha) associated with a color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/numberOfComponents
func CGColorGetNumberOfComponents(color CGColorRef) uintptr {
	result, callErr := tryCGColorGetNumberOfComponents(color)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorGetPattern func(color CGColorRef) CGPatternRef
var _cGColorGetPatternErr error

func tryCGColorGetPattern(color CGColorRef) (CGPatternRef, error) {
	if _cGColorGetPattern == nil {
		return 0, symbolCallError("CGColorGetPattern", "10.3", _cGColorGetPatternErr)
	}
	return _cGColorGetPattern(color), nil
}

// CGColorGetPattern returns the pattern associated with a color in a pattern color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/pattern
func CGColorGetPattern(color CGColorRef) CGPatternRef {
	result, callErr := tryCGColorGetPattern(color)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorGetTypeID func() uint
var _cGColorGetTypeIDErr error

func tryCGColorGetTypeID() (uint, error) {
	if _cGColorGetTypeID == nil {
		return 0, symbolCallError("CGColorGetTypeID", "10.3", _cGColorGetTypeIDErr)
	}
	return _cGColorGetTypeID(), nil
}

// CGColorGetTypeID returns the Core Foundation type identifier for a color data type.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColor/typeID
func CGColorGetTypeID() uint {
	result, callErr := tryCGColorGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorRelease func(color CGColorRef)
var _cGColorReleaseErr error

func tryCGColorRelease(color CGColorRef) error {
	if _cGColorRelease == nil {
		return symbolCallError("CGColorRelease", "10.3", _cGColorReleaseErr)
	}
	_cGColorRelease(color)
	return nil
}

// CGColorRelease decrements the retain count of a color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorRelease
func CGColorRelease(color CGColorRef) {
	if callErr := tryCGColorRelease(color); callErr != nil {
		panic(callErr)
	}
}

var _cGColorRetain func(color CGColorRef) CGColorRef
var _cGColorRetainErr error

func tryCGColorRetain(color CGColorRef) (CGColorRef, error) {
	if _cGColorRetain == nil {
		return 0, symbolCallError("CGColorRetain", "10.3", _cGColorRetainErr)
	}
	return _cGColorRetain(color), nil
}

// CGColorRetain increments the retain count of a color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorRetain
func CGColorRetain(color CGColorRef) CGColorRef {
	result, callErr := tryCGColorRetain(color)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCopyBaseColorSpace func(space CGColorSpaceRef) CGColorSpaceRef
var _cGColorSpaceCopyBaseColorSpaceErr error

func tryCGColorSpaceCopyBaseColorSpace(space CGColorSpaceRef) (CGColorSpaceRef, error) {
	if _cGColorSpaceCopyBaseColorSpace == nil {
		return 0, symbolCallError("CGColorSpaceCopyBaseColorSpace", "15.0", _cGColorSpaceCopyBaseColorSpaceErr)
	}
	return _cGColorSpaceCopyBaseColorSpace(space), nil
}

// CGColorSpaceCopyBaseColorSpace.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCopyBaseColorSpace(_:)
func CGColorSpaceCopyBaseColorSpace(space CGColorSpaceRef) CGColorSpaceRef {
	result, callErr := tryCGColorSpaceCopyBaseColorSpace(space)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCopyICCData func(space CGColorSpaceRef) corefoundation.CFDataRef
var _cGColorSpaceCopyICCDataErr error

func tryCGColorSpaceCopyICCData(space CGColorSpaceRef) (corefoundation.CFDataRef, error) {
	if _cGColorSpaceCopyICCData == nil {
		return 0, symbolCallError("CGColorSpaceCopyICCData", "10.12", _cGColorSpaceCopyICCDataErr)
	}
	return _cGColorSpaceCopyICCData(space), nil
}

// CGColorSpaceCopyICCData returns a copy of the ICC profile data of the provided color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/copyICCData()
func CGColorSpaceCopyICCData(space CGColorSpaceRef) corefoundation.CFDataRef {
	result, callErr := tryCGColorSpaceCopyICCData(space)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCopyName func(space CGColorSpaceRef) corefoundation.CFStringRef
var _cGColorSpaceCopyNameErr error

func tryCGColorSpaceCopyName(space CGColorSpaceRef) (corefoundation.CFStringRef, error) {
	if _cGColorSpaceCopyName == nil {
		return 0, symbolCallError("CGColorSpaceCopyName", "10.6", _cGColorSpaceCopyNameErr)
	}
	return _cGColorSpaceCopyName(space), nil
}

// CGColorSpaceCopyName returns the name used to create the specified color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/name
func CGColorSpaceCopyName(space CGColorSpaceRef) corefoundation.CFStringRef {
	result, callErr := tryCGColorSpaceCopyName(space)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCopyPropertyList func(space CGColorSpaceRef) corefoundation.CFPropertyListRef
var _cGColorSpaceCopyPropertyListErr error

func tryCGColorSpaceCopyPropertyList(space CGColorSpaceRef) (corefoundation.CFPropertyListRef, error) {
	if _cGColorSpaceCopyPropertyList == nil {
		return 0, symbolCallError("CGColorSpaceCopyPropertyList", "10.12", _cGColorSpaceCopyPropertyListErr)
	}
	return _cGColorSpaceCopyPropertyList(space), nil
}

// CGColorSpaceCopyPropertyList returns a copy of the color space’s properties.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/copyPropertyList()
func CGColorSpaceCopyPropertyList(space CGColorSpaceRef) corefoundation.CFPropertyListRef {
	result, callErr := tryCGColorSpaceCopyPropertyList(space)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCreateCalibratedGray func(whitePoint float64, blackPoint float64, gamma float64) CGColorSpaceRef
var _cGColorSpaceCreateCalibratedGrayErr error

func tryCGColorSpaceCreateCalibratedGray(whitePoint float64, blackPoint float64, gamma float64) (CGColorSpaceRef, error) {
	if _cGColorSpaceCreateCalibratedGray == nil {
		return 0, symbolCallError("CGColorSpaceCreateCalibratedGray", "10.0", _cGColorSpaceCreateCalibratedGrayErr)
	}
	return _cGColorSpaceCreateCalibratedGray(whitePoint, blackPoint, gamma), nil
}

// CGColorSpaceCreateCalibratedGray creates a calibrated grayscale color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(calibratedGrayWhitePoint:blackPoint:gamma:)
func CGColorSpaceCreateCalibratedGray(whitePoint float64, blackPoint float64, gamma float64) CGColorSpaceRef {
	result, callErr := tryCGColorSpaceCreateCalibratedGray(whitePoint, blackPoint, gamma)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCreateCalibratedRGB func(whitePoint float64, blackPoint float64, gamma float64, matrix float64) CGColorSpaceRef
var _cGColorSpaceCreateCalibratedRGBErr error

func tryCGColorSpaceCreateCalibratedRGB(whitePoint float64, blackPoint float64, gamma float64, matrix float64) (CGColorSpaceRef, error) {
	if _cGColorSpaceCreateCalibratedRGB == nil {
		return 0, symbolCallError("CGColorSpaceCreateCalibratedRGB", "10.0", _cGColorSpaceCreateCalibratedRGBErr)
	}
	return _cGColorSpaceCreateCalibratedRGB(whitePoint, blackPoint, gamma, matrix), nil
}

// CGColorSpaceCreateCalibratedRGB creates a calibrated RGB color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(calibratedRGBWhitePoint:blackPoint:gamma:matrix:)
func CGColorSpaceCreateCalibratedRGB(whitePoint float64, blackPoint float64, gamma float64, matrix float64) CGColorSpaceRef {
	result, callErr := tryCGColorSpaceCreateCalibratedRGB(whitePoint, blackPoint, gamma, matrix)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCreateCopyWithStandardRange func(space CGColorSpaceRef) CGColorSpaceRef
var _cGColorSpaceCreateCopyWithStandardRangeErr error

func tryCGColorSpaceCreateCopyWithStandardRange(space CGColorSpaceRef) (CGColorSpaceRef, error) {
	if _cGColorSpaceCreateCopyWithStandardRange == nil {
		return 0, symbolCallError("CGColorSpaceCreateCopyWithStandardRange", "13.0", _cGColorSpaceCreateCopyWithStandardRangeErr)
	}
	return _cGColorSpaceCreateCopyWithStandardRange(space), nil
}

// CGColorSpaceCreateCopyWithStandardRange.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateCopyWithStandardRange(_:)
func CGColorSpaceCreateCopyWithStandardRange(space CGColorSpaceRef) CGColorSpaceRef {
	result, callErr := tryCGColorSpaceCreateCopyWithStandardRange(space)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCreateDeviceCMYK func() CGColorSpaceRef
var _cGColorSpaceCreateDeviceCMYKErr error

func tryCGColorSpaceCreateDeviceCMYK() (CGColorSpaceRef, error) {
	if _cGColorSpaceCreateDeviceCMYK == nil {
		return 0, symbolCallError("CGColorSpaceCreateDeviceCMYK", "10.0", _cGColorSpaceCreateDeviceCMYKErr)
	}
	return _cGColorSpaceCreateDeviceCMYK(), nil
}

// CGColorSpaceCreateDeviceCMYK creates a device-dependent CMYK color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateDeviceCMYK()
func CGColorSpaceCreateDeviceCMYK() CGColorSpaceRef {
	result, callErr := tryCGColorSpaceCreateDeviceCMYK()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCreateDeviceGray func() CGColorSpaceRef
var _cGColorSpaceCreateDeviceGrayErr error

func tryCGColorSpaceCreateDeviceGray() (CGColorSpaceRef, error) {
	if _cGColorSpaceCreateDeviceGray == nil {
		return 0, symbolCallError("CGColorSpaceCreateDeviceGray", "10.0", _cGColorSpaceCreateDeviceGrayErr)
	}
	return _cGColorSpaceCreateDeviceGray(), nil
}

// CGColorSpaceCreateDeviceGray creates a device-dependent grayscale color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateDeviceGray()
func CGColorSpaceCreateDeviceGray() CGColorSpaceRef {
	result, callErr := tryCGColorSpaceCreateDeviceGray()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCreateDeviceRGB func() CGColorSpaceRef
var _cGColorSpaceCreateDeviceRGBErr error

func tryCGColorSpaceCreateDeviceRGB() (CGColorSpaceRef, error) {
	if _cGColorSpaceCreateDeviceRGB == nil {
		return 0, symbolCallError("CGColorSpaceCreateDeviceRGB", "10.0", _cGColorSpaceCreateDeviceRGBErr)
	}
	return _cGColorSpaceCreateDeviceRGB(), nil
}

// CGColorSpaceCreateDeviceRGB creates a device-dependent RGB color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateDeviceRGB()
func CGColorSpaceCreateDeviceRGB() CGColorSpaceRef {
	result, callErr := tryCGColorSpaceCreateDeviceRGB()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCreateExtended func(space CGColorSpaceRef) CGColorSpaceRef
var _cGColorSpaceCreateExtendedErr error

func tryCGColorSpaceCreateExtended(space CGColorSpaceRef) (CGColorSpaceRef, error) {
	if _cGColorSpaceCreateExtended == nil {
		return 0, symbolCallError("CGColorSpaceCreateExtended", "11.0", _cGColorSpaceCreateExtendedErr)
	}
	return _cGColorSpaceCreateExtended(space), nil
}

// CGColorSpaceCreateExtended.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateExtended(_:)
func CGColorSpaceCreateExtended(space CGColorSpaceRef) CGColorSpaceRef {
	result, callErr := tryCGColorSpaceCreateExtended(space)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCreateExtendedLinearized func(space CGColorSpaceRef) CGColorSpaceRef
var _cGColorSpaceCreateExtendedLinearizedErr error

func tryCGColorSpaceCreateExtendedLinearized(space CGColorSpaceRef) (CGColorSpaceRef, error) {
	if _cGColorSpaceCreateExtendedLinearized == nil {
		return 0, symbolCallError("CGColorSpaceCreateExtendedLinearized", "11.0", _cGColorSpaceCreateExtendedLinearizedErr)
	}
	return _cGColorSpaceCreateExtendedLinearized(space), nil
}

// CGColorSpaceCreateExtendedLinearized.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateExtendedLinearized(_:)
func CGColorSpaceCreateExtendedLinearized(space CGColorSpaceRef) CGColorSpaceRef {
	result, callErr := tryCGColorSpaceCreateExtendedLinearized(space)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCreateICCBased func(nComponents uintptr, range_ *float64, profile CGDataProviderRef, alternate CGColorSpaceRef) CGColorSpaceRef
var _cGColorSpaceCreateICCBasedErr error

func tryCGColorSpaceCreateICCBased(nComponents uintptr, range_ *float64, profile CGDataProviderRef, alternate CGColorSpaceRef) (CGColorSpaceRef, error) {
	if _cGColorSpaceCreateICCBased == nil {
		return 0, symbolCallError("CGColorSpaceCreateICCBased", "10.0", _cGColorSpaceCreateICCBasedErr)
	}
	return _cGColorSpaceCreateICCBased(nComponents, range_, profile, alternate), nil
}

// CGColorSpaceCreateICCBased creates a device-independent color space that is defined according to the ICC color profile specification.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(iccBasedNComponents:range:profile:alternate:)
func CGColorSpaceCreateICCBased(nComponents uintptr, range_ *float64, profile CGDataProviderRef, alternate CGColorSpaceRef) CGColorSpaceRef {
	result, callErr := tryCGColorSpaceCreateICCBased(nComponents, range_, profile, alternate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCreateIndexed func(baseSpace CGColorSpaceRef, lastIndex uintptr, colorTable string) CGColorSpaceRef
var _cGColorSpaceCreateIndexedErr error

func tryCGColorSpaceCreateIndexed(baseSpace CGColorSpaceRef, lastIndex uintptr, colorTable string) (CGColorSpaceRef, error) {
	if _cGColorSpaceCreateIndexed == nil {
		return 0, symbolCallError("CGColorSpaceCreateIndexed", "10.0", _cGColorSpaceCreateIndexedErr)
	}
	return _cGColorSpaceCreateIndexed(baseSpace, lastIndex, colorTable), nil
}

// CGColorSpaceCreateIndexed creates an indexed color space, consisting of colors specified by a color lookup table.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(indexedBaseSpace:last:colorTable:)
func CGColorSpaceCreateIndexed(baseSpace CGColorSpaceRef, lastIndex uintptr, colorTable string) CGColorSpaceRef {
	result, callErr := tryCGColorSpaceCreateIndexed(baseSpace, lastIndex, colorTable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCreateLab func(whitePoint float64, blackPoint float64, range_ float64) CGColorSpaceRef
var _cGColorSpaceCreateLabErr error

func tryCGColorSpaceCreateLab(whitePoint float64, blackPoint float64, range_ float64) (CGColorSpaceRef, error) {
	if _cGColorSpaceCreateLab == nil {
		return 0, symbolCallError("CGColorSpaceCreateLab", "10.0", _cGColorSpaceCreateLabErr)
	}
	return _cGColorSpaceCreateLab(whitePoint, blackPoint, range_), nil
}

// CGColorSpaceCreateLab creates a device-independent color space that is relative to human color perception, according to the CIE L*a*b* standard.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(labWhitePoint:blackPoint:range:)
func CGColorSpaceCreateLab(whitePoint float64, blackPoint float64, range_ float64) CGColorSpaceRef {
	result, callErr := tryCGColorSpaceCreateLab(whitePoint, blackPoint, range_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCreateLinearized func(space CGColorSpaceRef) CGColorSpaceRef
var _cGColorSpaceCreateLinearizedErr error

func tryCGColorSpaceCreateLinearized(space CGColorSpaceRef) (CGColorSpaceRef, error) {
	if _cGColorSpaceCreateLinearized == nil {
		return 0, symbolCallError("CGColorSpaceCreateLinearized", "11.0", _cGColorSpaceCreateLinearizedErr)
	}
	return _cGColorSpaceCreateLinearized(space), nil
}

// CGColorSpaceCreateLinearized.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateLinearized(_:)
func CGColorSpaceCreateLinearized(space CGColorSpaceRef) CGColorSpaceRef {
	result, callErr := tryCGColorSpaceCreateLinearized(space)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCreatePattern func(baseSpace CGColorSpaceRef) CGColorSpaceRef
var _cGColorSpaceCreatePatternErr error

func tryCGColorSpaceCreatePattern(baseSpace CGColorSpaceRef) (CGColorSpaceRef, error) {
	if _cGColorSpaceCreatePattern == nil {
		return 0, symbolCallError("CGColorSpaceCreatePattern", "10.0", _cGColorSpaceCreatePatternErr)
	}
	return _cGColorSpaceCreatePattern(baseSpace), nil
}

// CGColorSpaceCreatePattern creates a pattern color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(patternBaseSpace:)
func CGColorSpaceCreatePattern(baseSpace CGColorSpaceRef) CGColorSpaceRef {
	result, callErr := tryCGColorSpaceCreatePattern(baseSpace)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCreateWithColorSyncProfile func(arg0 ColorSyncProfileRef, options corefoundation.CFDictionaryRef) CGColorSpaceRef
var _cGColorSpaceCreateWithColorSyncProfileErr error

func tryCGColorSpaceCreateWithColorSyncProfile(arg0 ColorSyncProfileRef, options corefoundation.CFDictionaryRef) (CGColorSpaceRef, error) {
	if _cGColorSpaceCreateWithColorSyncProfile == nil {
		return 0, symbolCallError("CGColorSpaceCreateWithColorSyncProfile", "12.0", _cGColorSpaceCreateWithColorSyncProfileErr)
	}
	return _cGColorSpaceCreateWithColorSyncProfile(arg0, options), nil
}

// CGColorSpaceCreateWithColorSyncProfile.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceCreateWithColorSyncProfile(_:_:)
func CGColorSpaceCreateWithColorSyncProfile(arg0 ColorSyncProfileRef, options corefoundation.CFDictionaryRef) CGColorSpaceRef {
	result, callErr := tryCGColorSpaceCreateWithColorSyncProfile(arg0, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCreateWithICCData func(data corefoundation.CFTypeRef) CGColorSpaceRef
var _cGColorSpaceCreateWithICCDataErr error

func tryCGColorSpaceCreateWithICCData(data corefoundation.CFTypeRef) (CGColorSpaceRef, error) {
	if _cGColorSpaceCreateWithICCData == nil {
		return 0, symbolCallError("CGColorSpaceCreateWithICCData", "10.12", _cGColorSpaceCreateWithICCDataErr)
	}
	return _cGColorSpaceCreateWithICCData(data), nil
}

// CGColorSpaceCreateWithICCData creates an ICC-based color space using the ICC profile contained in the specified data.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(iccData:)
func CGColorSpaceCreateWithICCData(data corefoundation.CFTypeRef) CGColorSpaceRef {
	result, callErr := tryCGColorSpaceCreateWithICCData(data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCreateWithName func(name corefoundation.CFStringRef) CGColorSpaceRef
var _cGColorSpaceCreateWithNameErr error

func tryCGColorSpaceCreateWithName(name corefoundation.CFStringRef) (CGColorSpaceRef, error) {
	if _cGColorSpaceCreateWithName == nil {
		return 0, symbolCallError("CGColorSpaceCreateWithName", "10.2", _cGColorSpaceCreateWithNameErr)
	}
	return _cGColorSpaceCreateWithName(name), nil
}

// CGColorSpaceCreateWithName creates a specified type of Quartz color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(name:)
func CGColorSpaceCreateWithName(name corefoundation.CFStringRef) CGColorSpaceRef {
	result, callErr := tryCGColorSpaceCreateWithName(name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceCreateWithPropertyList func(plist corefoundation.CFPropertyListRef) CGColorSpaceRef
var _cGColorSpaceCreateWithPropertyListErr error

func tryCGColorSpaceCreateWithPropertyList(plist corefoundation.CFPropertyListRef) (CGColorSpaceRef, error) {
	if _cGColorSpaceCreateWithPropertyList == nil {
		return 0, symbolCallError("CGColorSpaceCreateWithPropertyList", "10.12", _cGColorSpaceCreateWithPropertyListErr)
	}
	return _cGColorSpaceCreateWithPropertyList(plist), nil
}

// CGColorSpaceCreateWithPropertyList creates a color space from a property list.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/init(propertyListPlist:)
func CGColorSpaceCreateWithPropertyList(plist corefoundation.CFPropertyListRef) CGColorSpaceRef {
	result, callErr := tryCGColorSpaceCreateWithPropertyList(plist)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceGetBaseColorSpace func(space CGColorSpaceRef) CGColorSpaceRef
var _cGColorSpaceGetBaseColorSpaceErr error

func tryCGColorSpaceGetBaseColorSpace(space CGColorSpaceRef) (CGColorSpaceRef, error) {
	if _cGColorSpaceGetBaseColorSpace == nil {
		return 0, symbolCallError("CGColorSpaceGetBaseColorSpace", "10.5", _cGColorSpaceGetBaseColorSpaceErr)
	}
	return _cGColorSpaceGetBaseColorSpace(space), nil
}

// CGColorSpaceGetBaseColorSpace returns the base color space of a pattern or indexed color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/baseColorSpace
func CGColorSpaceGetBaseColorSpace(space CGColorSpaceRef) CGColorSpaceRef {
	result, callErr := tryCGColorSpaceGetBaseColorSpace(space)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceGetColorTable func(space CGColorSpaceRef, table *uint8)
var _cGColorSpaceGetColorTableErr error

func tryCGColorSpaceGetColorTable(space CGColorSpaceRef, table *uint8) error {
	if _cGColorSpaceGetColorTable == nil {
		return symbolCallError("CGColorSpaceGetColorTable", "10.5", _cGColorSpaceGetColorTableErr)
	}
	_cGColorSpaceGetColorTable(space, table)
	return nil
}

// CGColorSpaceGetColorTable copies the entries in the color table of an indexed color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceGetColorTable
func CGColorSpaceGetColorTable(space CGColorSpaceRef, table *uint8) {
	if callErr := tryCGColorSpaceGetColorTable(space, table); callErr != nil {
		panic(callErr)
	}
}

var _cGColorSpaceGetColorTableCount func(space CGColorSpaceRef) uintptr
var _cGColorSpaceGetColorTableCountErr error

func tryCGColorSpaceGetColorTableCount(space CGColorSpaceRef) (uintptr, error) {
	if _cGColorSpaceGetColorTableCount == nil {
		return 0, symbolCallError("CGColorSpaceGetColorTableCount", "10.5", _cGColorSpaceGetColorTableCountErr)
	}
	return _cGColorSpaceGetColorTableCount(space), nil
}

// CGColorSpaceGetColorTableCount returns the number of entries in the color table of an indexed color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceGetColorTableCount
func CGColorSpaceGetColorTableCount(space CGColorSpaceRef) uintptr {
	result, callErr := tryCGColorSpaceGetColorTableCount(space)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceGetModel func(space CGColorSpaceRef) CGColorSpaceModel
var _cGColorSpaceGetModelErr error

func tryCGColorSpaceGetModel(space CGColorSpaceRef) (CGColorSpaceModel, error) {
	if _cGColorSpaceGetModel == nil {
		return *new(CGColorSpaceModel), symbolCallError("CGColorSpaceGetModel", "10.5", _cGColorSpaceGetModelErr)
	}
	return _cGColorSpaceGetModel(space), nil
}

// CGColorSpaceGetModel returns the color space model of the provided color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/model
func CGColorSpaceGetModel(space CGColorSpaceRef) CGColorSpaceModel {
	result, callErr := tryCGColorSpaceGetModel(space)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceGetName func(space CGColorSpaceRef) corefoundation.CFStringRef
var _cGColorSpaceGetNameErr error

func tryCGColorSpaceGetName(space CGColorSpaceRef) (corefoundation.CFStringRef, error) {
	if _cGColorSpaceGetName == nil {
		return 0, symbolCallError("CGColorSpaceGetName", "10.13", _cGColorSpaceGetNameErr)
	}
	return _cGColorSpaceGetName(space), nil
}

// CGColorSpaceGetName.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceGetName
func CGColorSpaceGetName(space CGColorSpaceRef) corefoundation.CFStringRef {
	result, callErr := tryCGColorSpaceGetName(space)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceGetNumberOfComponents func(space CGColorSpaceRef) uintptr
var _cGColorSpaceGetNumberOfComponentsErr error

func tryCGColorSpaceGetNumberOfComponents(space CGColorSpaceRef) (uintptr, error) {
	if _cGColorSpaceGetNumberOfComponents == nil {
		return 0, symbolCallError("CGColorSpaceGetNumberOfComponents", "10.0", _cGColorSpaceGetNumberOfComponentsErr)
	}
	return _cGColorSpaceGetNumberOfComponents(space), nil
}

// CGColorSpaceGetNumberOfComponents returns the number of color components in a color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/numberOfComponents
func CGColorSpaceGetNumberOfComponents(space CGColorSpaceRef) uintptr {
	result, callErr := tryCGColorSpaceGetNumberOfComponents(space)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceGetTypeID func() uint
var _cGColorSpaceGetTypeIDErr error

func tryCGColorSpaceGetTypeID() (uint, error) {
	if _cGColorSpaceGetTypeID == nil {
		return 0, symbolCallError("CGColorSpaceGetTypeID", "10.2", _cGColorSpaceGetTypeIDErr)
	}
	return _cGColorSpaceGetTypeID(), nil
}

// CGColorSpaceGetTypeID returns the Core Foundation type identifier for Quartz color spaces.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/typeID
func CGColorSpaceGetTypeID() uint {
	result, callErr := tryCGColorSpaceGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceIsHDR func(arg0 CGColorSpaceRef) bool
var _cGColorSpaceIsHDRErr error

func tryCGColorSpaceIsHDR(arg0 CGColorSpaceRef) (bool, error) {
	if _cGColorSpaceIsHDR == nil {
		return false, symbolCallError("CGColorSpaceIsHDR", "10.15", _cGColorSpaceIsHDRErr)
	}
	return _cGColorSpaceIsHDR(arg0), nil
}

// CGColorSpaceIsHDR.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/isHDR()
func CGColorSpaceIsHDR(arg0 CGColorSpaceRef) bool {
	result, callErr := tryCGColorSpaceIsHDR(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceIsHLGBased func(s CGColorSpaceRef) bool
var _cGColorSpaceIsHLGBasedErr error

func tryCGColorSpaceIsHLGBased(s CGColorSpaceRef) (bool, error) {
	if _cGColorSpaceIsHLGBased == nil {
		return false, symbolCallError("CGColorSpaceIsHLGBased", "12.0", _cGColorSpaceIsHLGBasedErr)
	}
	return _cGColorSpaceIsHLGBased(s), nil
}

// CGColorSpaceIsHLGBased.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceIsHLGBased(_:)
func CGColorSpaceIsHLGBased(s CGColorSpaceRef) bool {
	result, callErr := tryCGColorSpaceIsHLGBased(s)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceIsPQBased func(s CGColorSpaceRef) bool
var _cGColorSpaceIsPQBasedErr error

func tryCGColorSpaceIsPQBased(s CGColorSpaceRef) (bool, error) {
	if _cGColorSpaceIsPQBased == nil {
		return false, symbolCallError("CGColorSpaceIsPQBased", "12.0", _cGColorSpaceIsPQBasedErr)
	}
	return _cGColorSpaceIsPQBased(s), nil
}

// CGColorSpaceIsPQBased.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceIsPQBased(_:)
func CGColorSpaceIsPQBased(s CGColorSpaceRef) bool {
	result, callErr := tryCGColorSpaceIsPQBased(s)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceIsWideGamutRGB func(arg0 CGColorSpaceRef) bool
var _cGColorSpaceIsWideGamutRGBErr error

func tryCGColorSpaceIsWideGamutRGB(arg0 CGColorSpaceRef) (bool, error) {
	if _cGColorSpaceIsWideGamutRGB == nil {
		return false, symbolCallError("CGColorSpaceIsWideGamutRGB", "10.12", _cGColorSpaceIsWideGamutRGBErr)
	}
	return _cGColorSpaceIsWideGamutRGB(arg0), nil
}

// CGColorSpaceIsWideGamutRGB returns whether the RGB color space covers a significant portion of the NTSC color gamut.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/isWideGamutRGB
func CGColorSpaceIsWideGamutRGB(arg0 CGColorSpaceRef) bool {
	result, callErr := tryCGColorSpaceIsWideGamutRGB(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceRelease func(space CGColorSpaceRef)
var _cGColorSpaceReleaseErr error

func tryCGColorSpaceRelease(space CGColorSpaceRef) error {
	if _cGColorSpaceRelease == nil {
		return symbolCallError("CGColorSpaceRelease", "10.0", _cGColorSpaceReleaseErr)
	}
	_cGColorSpaceRelease(space)
	return nil
}

// CGColorSpaceRelease decrements the retain count of a color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceRelease
func CGColorSpaceRelease(space CGColorSpaceRef) {
	if callErr := tryCGColorSpaceRelease(space); callErr != nil {
		panic(callErr)
	}
}

var _cGColorSpaceRetain func(space CGColorSpaceRef) CGColorSpaceRef
var _cGColorSpaceRetainErr error

func tryCGColorSpaceRetain(space CGColorSpaceRef) (CGColorSpaceRef, error) {
	if _cGColorSpaceRetain == nil {
		return 0, symbolCallError("CGColorSpaceRetain", "10.0", _cGColorSpaceRetainErr)
	}
	return _cGColorSpaceRetain(space), nil
}

// CGColorSpaceRetain increments the retain count of a color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceRetain
func CGColorSpaceRetain(space CGColorSpaceRef) CGColorSpaceRef {
	result, callErr := tryCGColorSpaceRetain(space)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceSupportsOutput func(space CGColorSpaceRef) bool
var _cGColorSpaceSupportsOutputErr error

func tryCGColorSpaceSupportsOutput(space CGColorSpaceRef) (bool, error) {
	if _cGColorSpaceSupportsOutput == nil {
		return false, symbolCallError("CGColorSpaceSupportsOutput", "10.12", _cGColorSpaceSupportsOutputErr)
	}
	return _cGColorSpaceSupportsOutput(space), nil
}

// CGColorSpaceSupportsOutput returns a Boolean indicating whether the color space can be used as a destination color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace/supportsOutput
func CGColorSpaceSupportsOutput(space CGColorSpaceRef) bool {
	result, callErr := tryCGColorSpaceSupportsOutput(space)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceUsesExtendedRange func(space CGColorSpaceRef) bool
var _cGColorSpaceUsesExtendedRangeErr error

func tryCGColorSpaceUsesExtendedRange(space CGColorSpaceRef) (bool, error) {
	if _cGColorSpaceUsesExtendedRange == nil {
		return false, symbolCallError("CGColorSpaceUsesExtendedRange", "10.12", _cGColorSpaceUsesExtendedRangeErr)
	}
	return _cGColorSpaceUsesExtendedRange(space), nil
}

// CGColorSpaceUsesExtendedRange.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceUsesExtendedRange(_:)
func CGColorSpaceUsesExtendedRange(space CGColorSpaceRef) bool {
	result, callErr := tryCGColorSpaceUsesExtendedRange(space)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGColorSpaceUsesITUR_2100TF func(arg0 CGColorSpaceRef) bool
var _cGColorSpaceUsesITUR_2100TFErr error

func tryCGColorSpaceUsesITUR_2100TF(arg0 CGColorSpaceRef) (bool, error) {
	if _cGColorSpaceUsesITUR_2100TF == nil {
		return false, symbolCallError("CGColorSpaceUsesITUR_2100TF", "11.0", _cGColorSpaceUsesITUR_2100TFErr)
	}
	return _cGColorSpaceUsesITUR_2100TF(arg0), nil
}

// CGColorSpaceUsesITUR_2100TF.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceUsesITUR_2100TF(_:)
func CGColorSpaceUsesITUR_2100TF(arg0 CGColorSpaceRef) bool {
	result, callErr := tryCGColorSpaceUsesITUR_2100TF(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGCompleteDisplayConfiguration func(config CGDisplayConfigRef, option CGConfigureOption) CGError
var _cGCompleteDisplayConfigurationErr error

func tryCGCompleteDisplayConfiguration(config CGDisplayConfigRef, option CGConfigureOption) (CGError, error) {
	if _cGCompleteDisplayConfiguration == nil {
		return *new(CGError), symbolCallError("CGCompleteDisplayConfiguration", "10.0", _cGCompleteDisplayConfigurationErr)
	}
	return _cGCompleteDisplayConfiguration(config, option), nil
}

// CGCompleteDisplayConfiguration completes a set of display configuration changes.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGCompleteDisplayConfiguration(_:_:)
func CGCompleteDisplayConfiguration(config CGDisplayConfigRef, option CGConfigureOption) CGError {
	result, callErr := tryCGCompleteDisplayConfiguration(config, option)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGConfigureDisplayFadeEffect func(config CGDisplayConfigRef, fadeOutSeconds float32, fadeInSeconds float32, fadeRed float32, fadeGreen float32, fadeBlue float32) CGError
var _cGConfigureDisplayFadeEffectErr error

func tryCGConfigureDisplayFadeEffect(config CGDisplayConfigRef, fadeOutSeconds float32, fadeInSeconds float32, fadeRed float32, fadeGreen float32, fadeBlue float32) (CGError, error) {
	if _cGConfigureDisplayFadeEffect == nil {
		return *new(CGError), symbolCallError("CGConfigureDisplayFadeEffect", "10.2", _cGConfigureDisplayFadeEffectErr)
	}
	return _cGConfigureDisplayFadeEffect(config, fadeOutSeconds, fadeInSeconds, fadeRed, fadeGreen, fadeBlue), nil
}

// CGConfigureDisplayFadeEffect modifies the settings of the built-in fade effect that occurs during a display configuration.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGConfigureDisplayFadeEffect(_:_:_:_:_:_:)
func CGConfigureDisplayFadeEffect(config CGDisplayConfigRef, fadeOutSeconds float32, fadeInSeconds float32, fadeRed float32, fadeGreen float32, fadeBlue float32) CGError {
	result, callErr := tryCGConfigureDisplayFadeEffect(config, fadeOutSeconds, fadeInSeconds, fadeRed, fadeGreen, fadeBlue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGConfigureDisplayMirrorOfDisplay func(config CGDisplayConfigRef, display uint32, master uint32) CGError
var _cGConfigureDisplayMirrorOfDisplayErr error

func tryCGConfigureDisplayMirrorOfDisplay(config CGDisplayConfigRef, display uint32, master uint32) (CGError, error) {
	if _cGConfigureDisplayMirrorOfDisplay == nil {
		return *new(CGError), symbolCallError("CGConfigureDisplayMirrorOfDisplay", "10.2", _cGConfigureDisplayMirrorOfDisplayErr)
	}
	return _cGConfigureDisplayMirrorOfDisplay(config, display, master), nil
}

// CGConfigureDisplayMirrorOfDisplay changes the configuration of a mirroring set.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGConfigureDisplayMirrorOfDisplay(_:_:_:)
func CGConfigureDisplayMirrorOfDisplay(config CGDisplayConfigRef, display uint32, master uint32) CGError {
	result, callErr := tryCGConfigureDisplayMirrorOfDisplay(config, display, master)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGConfigureDisplayOrigin func(config CGDisplayConfigRef, display uint32, x int32, y int32) CGError
var _cGConfigureDisplayOriginErr error

func tryCGConfigureDisplayOrigin(config CGDisplayConfigRef, display uint32, x int32, y int32) (CGError, error) {
	if _cGConfigureDisplayOrigin == nil {
		return *new(CGError), symbolCallError("CGConfigureDisplayOrigin", "10.0", _cGConfigureDisplayOriginErr)
	}
	return _cGConfigureDisplayOrigin(config, display, x, y), nil
}

// CGConfigureDisplayOrigin configures the origin of a display relative to the global display coordinate space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGConfigureDisplayOrigin(_:_:_:_:)
func CGConfigureDisplayOrigin(config CGDisplayConfigRef, display uint32, x int32, y int32) CGError {
	result, callErr := tryCGConfigureDisplayOrigin(config, display, x, y)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGConfigureDisplayStereoOperation func(config CGDisplayConfigRef, display uint32, stereo bool, forceBlueLine bool) CGError
var _cGConfigureDisplayStereoOperationErr error

func tryCGConfigureDisplayStereoOperation(config CGDisplayConfigRef, display uint32, stereo bool, forceBlueLine bool) (CGError, error) {
	if _cGConfigureDisplayStereoOperation == nil {
		return *new(CGError), symbolCallError("CGConfigureDisplayStereoOperation", "10.4", _cGConfigureDisplayStereoOperationErr)
	}
	return _cGConfigureDisplayStereoOperation(config, display, stereo, forceBlueLine), nil
}

// CGConfigureDisplayStereoOperation enables or disables stereo operation for a display, as part of a display configuration.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGConfigureDisplayStereoOperation(_:_:_:_:)
func CGConfigureDisplayStereoOperation(config CGDisplayConfigRef, display uint32, stereo bool, forceBlueLine bool) CGError {
	result, callErr := tryCGConfigureDisplayStereoOperation(config, display, stereo, forceBlueLine)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGConfigureDisplayWithDisplayMode func(config CGDisplayConfigRef, display uint32, mode CGDisplayModeRef, options corefoundation.CFDictionaryRef) CGError
var _cGConfigureDisplayWithDisplayModeErr error

func tryCGConfigureDisplayWithDisplayMode(config CGDisplayConfigRef, display uint32, mode CGDisplayModeRef, options corefoundation.CFDictionaryRef) (CGError, error) {
	if _cGConfigureDisplayWithDisplayMode == nil {
		return *new(CGError), symbolCallError("CGConfigureDisplayWithDisplayMode", "10.6", _cGConfigureDisplayWithDisplayModeErr)
	}
	return _cGConfigureDisplayWithDisplayMode(config, display, mode, options), nil
}

// CGConfigureDisplayWithDisplayMode configures the display mode of a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGConfigureDisplayWithDisplayMode(_:_:_:_:)
func CGConfigureDisplayWithDisplayMode(config CGDisplayConfigRef, display uint32, mode CGDisplayModeRef, options corefoundation.CFDictionaryRef) CGError {
	result, callErr := tryCGConfigureDisplayWithDisplayMode(config, display, mode, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextAddArc func(c CGContextRef, x float64, y float64, radius float64, startAngle float64, endAngle float64, clockwise int)
var _cGContextAddArcErr error

func tryCGContextAddArc(c CGContextRef, x float64, y float64, radius float64, startAngle float64, endAngle float64, clockwise int) error {
	if _cGContextAddArc == nil {
		return symbolCallError("CGContextAddArc", "10.0", _cGContextAddArcErr)
	}
	_cGContextAddArc(c, x, y, radius, startAngle, endAngle, clockwise)
	return nil
}

// CGContextAddArc adds an arc of a circle to the current path, possibly preceded by a straight line segment
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextAddArc
func CGContextAddArc(c CGContextRef, x float64, y float64, radius float64, startAngle float64, endAngle float64, clockwise int) {
	if callErr := tryCGContextAddArc(c, x, y, radius, startAngle, endAngle, clockwise); callErr != nil {
		panic(callErr)
	}
}

var _cGContextAddArcToPoint func(c CGContextRef, x1 float64, y1 float64, x2 float64, y2 float64, radius float64)
var _cGContextAddArcToPointErr error

func tryCGContextAddArcToPoint(c CGContextRef, x1 float64, y1 float64, x2 float64, y2 float64, radius float64) error {
	if _cGContextAddArcToPoint == nil {
		return symbolCallError("CGContextAddArcToPoint", "10.0", _cGContextAddArcToPointErr)
	}
	_cGContextAddArcToPoint(c, x1, y1, x2, y2, radius)
	return nil
}

// CGContextAddArcToPoint adds an arc of a circle to the current path, using a radius and tangent points.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextAddArcToPoint
func CGContextAddArcToPoint(c CGContextRef, x1 float64, y1 float64, x2 float64, y2 float64, radius float64) {
	if callErr := tryCGContextAddArcToPoint(c, x1, y1, x2, y2, radius); callErr != nil {
		panic(callErr)
	}
}

var _cGContextAddCurveToPoint func(c CGContextRef, cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64)
var _cGContextAddCurveToPointErr error

func tryCGContextAddCurveToPoint(c CGContextRef, cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) error {
	if _cGContextAddCurveToPoint == nil {
		return symbolCallError("CGContextAddCurveToPoint", "10.0", _cGContextAddCurveToPointErr)
	}
	_cGContextAddCurveToPoint(c, cp1x, cp1y, cp2x, cp2y, x, y)
	return nil
}

// CGContextAddCurveToPoint appends a cubic Bézier curve from the current point, using the provided control points and end point .
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextAddCurveToPoint
func CGContextAddCurveToPoint(c CGContextRef, cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) {
	if callErr := tryCGContextAddCurveToPoint(c, cp1x, cp1y, cp2x, cp2y, x, y); callErr != nil {
		panic(callErr)
	}
}

var _cGContextAddEllipseInRect func(c CGContextRef, rect corefoundation.CGRect)
var _cGContextAddEllipseInRectErr error

func tryCGContextAddEllipseInRect(c CGContextRef, rect corefoundation.CGRect) error {
	if _cGContextAddEllipseInRect == nil {
		return symbolCallError("CGContextAddEllipseInRect", "10.4", _cGContextAddEllipseInRectErr)
	}
	_cGContextAddEllipseInRect(c, rect)
	return nil
}

// CGContextAddEllipseInRect adds an ellipse that fits inside the specified rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/addEllipse(in:)
func CGContextAddEllipseInRect(c CGContextRef, rect corefoundation.CGRect) {
	if callErr := tryCGContextAddEllipseInRect(c, rect); callErr != nil {
		panic(callErr)
	}
}

var _cGContextAddLineToPoint func(c CGContextRef, x float64, y float64)
var _cGContextAddLineToPointErr error

func tryCGContextAddLineToPoint(c CGContextRef, x float64, y float64) error {
	if _cGContextAddLineToPoint == nil {
		return symbolCallError("CGContextAddLineToPoint", "10.0", _cGContextAddLineToPointErr)
	}
	_cGContextAddLineToPoint(c, x, y)
	return nil
}

// CGContextAddLineToPoint appends a straight line segment from the current point to the provided point .
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextAddLineToPoint
func CGContextAddLineToPoint(c CGContextRef, x float64, y float64) {
	if callErr := tryCGContextAddLineToPoint(c, x, y); callErr != nil {
		panic(callErr)
	}
}

var _cGContextAddLines func(c CGContextRef, points *corefoundation.CGPoint, count uintptr)
var _cGContextAddLinesErr error

func tryCGContextAddLines(c CGContextRef, points *corefoundation.CGPoint, count uintptr) error {
	if _cGContextAddLines == nil {
		return symbolCallError("CGContextAddLines", "10.0", _cGContextAddLinesErr)
	}
	_cGContextAddLines(c, points, count)
	return nil
}

// CGContextAddLines adds a sequence of connected straight-line segments to the current path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextAddLines
func CGContextAddLines(c CGContextRef, points *corefoundation.CGPoint, count uintptr) {
	if callErr := tryCGContextAddLines(c, points, count); callErr != nil {
		panic(callErr)
	}
}

var _cGContextAddPath func(c CGContextRef, path CGPathRef)
var _cGContextAddPathErr error

func tryCGContextAddPath(c CGContextRef, path CGPathRef) error {
	if _cGContextAddPath == nil {
		return symbolCallError("CGContextAddPath", "10.2", _cGContextAddPathErr)
	}
	_cGContextAddPath(c, path)
	return nil
}

// CGContextAddPath adds a previously created path object to the current path in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/addPath(_:)
func CGContextAddPath(c CGContextRef, path CGPathRef) {
	if callErr := tryCGContextAddPath(c, path); callErr != nil {
		panic(callErr)
	}
}

var _cGContextAddQuadCurveToPoint func(c CGContextRef, cpx float64, cpy float64, x float64, y float64)
var _cGContextAddQuadCurveToPointErr error

func tryCGContextAddQuadCurveToPoint(c CGContextRef, cpx float64, cpy float64, x float64, y float64) error {
	if _cGContextAddQuadCurveToPoint == nil {
		return symbolCallError("CGContextAddQuadCurveToPoint", "10.0", _cGContextAddQuadCurveToPointErr)
	}
	_cGContextAddQuadCurveToPoint(c, cpx, cpy, x, y)
	return nil
}

// CGContextAddQuadCurveToPoint appends a quadratic Bézier curve from the current point, using a control point and an end point you specify.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextAddQuadCurveToPoint
func CGContextAddQuadCurveToPoint(c CGContextRef, cpx float64, cpy float64, x float64, y float64) {
	if callErr := tryCGContextAddQuadCurveToPoint(c, cpx, cpy, x, y); callErr != nil {
		panic(callErr)
	}
}

var _cGContextAddRect func(c CGContextRef, rect corefoundation.CGRect)
var _cGContextAddRectErr error

func tryCGContextAddRect(c CGContextRef, rect corefoundation.CGRect) error {
	if _cGContextAddRect == nil {
		return symbolCallError("CGContextAddRect", "10.0", _cGContextAddRectErr)
	}
	_cGContextAddRect(c, rect)
	return nil
}

// CGContextAddRect adds a rectangular path to the current path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/addRect(_:)
func CGContextAddRect(c CGContextRef, rect corefoundation.CGRect) {
	if callErr := tryCGContextAddRect(c, rect); callErr != nil {
		panic(callErr)
	}
}

var _cGContextAddRects func(c CGContextRef, rects *corefoundation.CGRect, count uintptr)
var _cGContextAddRectsErr error

func tryCGContextAddRects(c CGContextRef, rects *corefoundation.CGRect, count uintptr) error {
	if _cGContextAddRects == nil {
		return symbolCallError("CGContextAddRects", "10.0", _cGContextAddRectsErr)
	}
	_cGContextAddRects(c, rects, count)
	return nil
}

// CGContextAddRects adds a set of rectangular paths to the current path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextAddRects
func CGContextAddRects(c CGContextRef, rects *corefoundation.CGRect, count uintptr) {
	if callErr := tryCGContextAddRects(c, rects, count); callErr != nil {
		panic(callErr)
	}
}

var _cGContextBeginPage func(c CGContextRef, mediaBox *corefoundation.CGRect)
var _cGContextBeginPageErr error

func tryCGContextBeginPage(c CGContextRef, mediaBox *corefoundation.CGRect) error {
	if _cGContextBeginPage == nil {
		return symbolCallError("CGContextBeginPage", "10.0", _cGContextBeginPageErr)
	}
	_cGContextBeginPage(c, mediaBox)
	return nil
}

// CGContextBeginPage starts a new page in a page-based graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/beginPage(mediaBox:)
func CGContextBeginPage(c CGContextRef, mediaBox *corefoundation.CGRect) {
	if callErr := tryCGContextBeginPage(c, mediaBox); callErr != nil {
		panic(callErr)
	}
}

var _cGContextBeginPath func(c CGContextRef)
var _cGContextBeginPathErr error

func tryCGContextBeginPath(c CGContextRef) error {
	if _cGContextBeginPath == nil {
		return symbolCallError("CGContextBeginPath", "10.0", _cGContextBeginPathErr)
	}
	_cGContextBeginPath(c)
	return nil
}

// CGContextBeginPath creates a new empty path in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/beginPath()
func CGContextBeginPath(c CGContextRef) {
	if callErr := tryCGContextBeginPath(c); callErr != nil {
		panic(callErr)
	}
}

var _cGContextBeginTransparencyLayer func(c CGContextRef, auxiliaryInfo corefoundation.CFDictionaryRef)
var _cGContextBeginTransparencyLayerErr error

func tryCGContextBeginTransparencyLayer(c CGContextRef, auxiliaryInfo corefoundation.CFDictionaryRef) error {
	if _cGContextBeginTransparencyLayer == nil {
		return symbolCallError("CGContextBeginTransparencyLayer", "10.3", _cGContextBeginTransparencyLayerErr)
	}
	_cGContextBeginTransparencyLayer(c, auxiliaryInfo)
	return nil
}

// CGContextBeginTransparencyLayer begins a transparency layer.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/beginTransparencyLayer(auxiliaryInfo:)
func CGContextBeginTransparencyLayer(c CGContextRef, auxiliaryInfo corefoundation.CFDictionaryRef) {
	if callErr := tryCGContextBeginTransparencyLayer(c, auxiliaryInfo); callErr != nil {
		panic(callErr)
	}
}

var _cGContextBeginTransparencyLayerWithRect func(c CGContextRef, rect corefoundation.CGRect, auxInfo corefoundation.CFDictionaryRef)
var _cGContextBeginTransparencyLayerWithRectErr error

func tryCGContextBeginTransparencyLayerWithRect(c CGContextRef, rect corefoundation.CGRect, auxInfo corefoundation.CFDictionaryRef) error {
	if _cGContextBeginTransparencyLayerWithRect == nil {
		return symbolCallError("CGContextBeginTransparencyLayerWithRect", "10.5", _cGContextBeginTransparencyLayerWithRectErr)
	}
	_cGContextBeginTransparencyLayerWithRect(c, rect, auxInfo)
	return nil
}

// CGContextBeginTransparencyLayerWithRect begins a transparency layer whose contents are bounded by the specified rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/beginTransparencyLayer(in:auxiliaryInfo:)
func CGContextBeginTransparencyLayerWithRect(c CGContextRef, rect corefoundation.CGRect, auxInfo corefoundation.CFDictionaryRef) {
	if callErr := tryCGContextBeginTransparencyLayerWithRect(c, rect, auxInfo); callErr != nil {
		panic(callErr)
	}
}

var _cGContextClearRect func(c CGContextRef, rect corefoundation.CGRect)
var _cGContextClearRectErr error

func tryCGContextClearRect(c CGContextRef, rect corefoundation.CGRect) error {
	if _cGContextClearRect == nil {
		return symbolCallError("CGContextClearRect", "10.0", _cGContextClearRectErr)
	}
	_cGContextClearRect(c, rect)
	return nil
}

// CGContextClearRect paints a transparent rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/clear(_:)
func CGContextClearRect(c CGContextRef, rect corefoundation.CGRect) {
	if callErr := tryCGContextClearRect(c, rect); callErr != nil {
		panic(callErr)
	}
}

var _cGContextClip func(c CGContextRef)
var _cGContextClipErr error

func tryCGContextClip(c CGContextRef) error {
	if _cGContextClip == nil {
		return symbolCallError("CGContextClip", "10.0", _cGContextClipErr)
	}
	_cGContextClip(c)
	return nil
}

// CGContextClip modifies the current clipping path, using the nonzero winding number rule.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextClip
func CGContextClip(c CGContextRef) {
	if callErr := tryCGContextClip(c); callErr != nil {
		panic(callErr)
	}
}

var _cGContextClipToMask func(c CGContextRef, rect corefoundation.CGRect, mask CGImageRef)
var _cGContextClipToMaskErr error

func tryCGContextClipToMask(c CGContextRef, rect corefoundation.CGRect, mask CGImageRef) error {
	if _cGContextClipToMask == nil {
		return symbolCallError("CGContextClipToMask", "10.4", _cGContextClipToMaskErr)
	}
	_cGContextClipToMask(c, rect, mask)
	return nil
}

// CGContextClipToMask maps a mask into the specified rectangle and intersects it with the current clipping area of the graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/clip(to:mask:)
func CGContextClipToMask(c CGContextRef, rect corefoundation.CGRect, mask CGImageRef) {
	if callErr := tryCGContextClipToMask(c, rect, mask); callErr != nil {
		panic(callErr)
	}
}

var _cGContextClipToRect func(c CGContextRef, rect corefoundation.CGRect)
var _cGContextClipToRectErr error

func tryCGContextClipToRect(c CGContextRef, rect corefoundation.CGRect) error {
	if _cGContextClipToRect == nil {
		return symbolCallError("CGContextClipToRect", "10.0", _cGContextClipToRectErr)
	}
	_cGContextClipToRect(c, rect)
	return nil
}

// CGContextClipToRect sets the clipping path to the intersection of the current clipping path with the area defined by the specified rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/clip(to:)-7cbwq
func CGContextClipToRect(c CGContextRef, rect corefoundation.CGRect) {
	if callErr := tryCGContextClipToRect(c, rect); callErr != nil {
		panic(callErr)
	}
}

var _cGContextClipToRects func(c CGContextRef, rects *corefoundation.CGRect, count uintptr)
var _cGContextClipToRectsErr error

func tryCGContextClipToRects(c CGContextRef, rects *corefoundation.CGRect, count uintptr) error {
	if _cGContextClipToRects == nil {
		return symbolCallError("CGContextClipToRects", "10.0", _cGContextClipToRectsErr)
	}
	_cGContextClipToRects(c, rects, count)
	return nil
}

// CGContextClipToRects sets the clipping path to the intersection of the current clipping path with the region defined by an array of rectangles.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextClipToRects
func CGContextClipToRects(c CGContextRef, rects *corefoundation.CGRect, count uintptr) {
	if callErr := tryCGContextClipToRects(c, rects, count); callErr != nil {
		panic(callErr)
	}
}

var _cGContextClosePath func(c CGContextRef)
var _cGContextClosePathErr error

func tryCGContextClosePath(c CGContextRef) error {
	if _cGContextClosePath == nil {
		return symbolCallError("CGContextClosePath", "10.0", _cGContextClosePathErr)
	}
	_cGContextClosePath(c)
	return nil
}

// CGContextClosePath closes and terminates the current path’s subpath.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/closePath()
func CGContextClosePath(c CGContextRef) {
	if callErr := tryCGContextClosePath(c); callErr != nil {
		panic(callErr)
	}
}

var _cGContextConcatCTM func(c CGContextRef, transform corefoundation.CGAffineTransform)
var _cGContextConcatCTMErr error

func tryCGContextConcatCTM(c CGContextRef, transform corefoundation.CGAffineTransform) error {
	if _cGContextConcatCTM == nil {
		return symbolCallError("CGContextConcatCTM", "10.0", _cGContextConcatCTMErr)
	}
	_cGContextConcatCTM(c, transform)
	return nil
}

// CGContextConcatCTM transforms the user coordinate system in a context using a specified matrix.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/concatenate(_:)
func CGContextConcatCTM(c CGContextRef, transform corefoundation.CGAffineTransform) {
	if callErr := tryCGContextConcatCTM(c, transform); callErr != nil {
		panic(callErr)
	}
}

var _cGContextConvertPointToDeviceSpace func(c CGContextRef, point corefoundation.CGPoint) corefoundation.CGPoint
var _cGContextConvertPointToDeviceSpaceErr error

func tryCGContextConvertPointToDeviceSpace(c CGContextRef, point corefoundation.CGPoint) (corefoundation.CGPoint, error) {
	if _cGContextConvertPointToDeviceSpace == nil {
		return corefoundation.CGPoint{}, symbolCallError("CGContextConvertPointToDeviceSpace", "10.4", _cGContextConvertPointToDeviceSpaceErr)
	}
	return _cGContextConvertPointToDeviceSpace(c, point), nil
}

// CGContextConvertPointToDeviceSpace returns a point that is transformed from user space coordinates to device space coordinates.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/convertToDeviceSpace(_:)-53m7u
func CGContextConvertPointToDeviceSpace(c CGContextRef, point corefoundation.CGPoint) corefoundation.CGPoint {
	result, callErr := tryCGContextConvertPointToDeviceSpace(c, point)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextConvertPointToUserSpace func(c CGContextRef, point corefoundation.CGPoint) corefoundation.CGPoint
var _cGContextConvertPointToUserSpaceErr error

func tryCGContextConvertPointToUserSpace(c CGContextRef, point corefoundation.CGPoint) (corefoundation.CGPoint, error) {
	if _cGContextConvertPointToUserSpace == nil {
		return corefoundation.CGPoint{}, symbolCallError("CGContextConvertPointToUserSpace", "10.4", _cGContextConvertPointToUserSpaceErr)
	}
	return _cGContextConvertPointToUserSpace(c, point), nil
}

// CGContextConvertPointToUserSpace returns a point that is transformed from device space coordinates to user space coordinates.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/convertToUserSpace(_:)-3mtg3
func CGContextConvertPointToUserSpace(c CGContextRef, point corefoundation.CGPoint) corefoundation.CGPoint {
	result, callErr := tryCGContextConvertPointToUserSpace(c, point)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextConvertRectToDeviceSpace func(c CGContextRef, rect corefoundation.CGRect) corefoundation.CGRect
var _cGContextConvertRectToDeviceSpaceErr error

func tryCGContextConvertRectToDeviceSpace(c CGContextRef, rect corefoundation.CGRect) (corefoundation.CGRect, error) {
	if _cGContextConvertRectToDeviceSpace == nil {
		return corefoundation.CGRect{}, symbolCallError("CGContextConvertRectToDeviceSpace", "10.4", _cGContextConvertRectToDeviceSpaceErr)
	}
	return _cGContextConvertRectToDeviceSpace(c, rect), nil
}

// CGContextConvertRectToDeviceSpace returns a rectangle that is transformed from user space coordinate to device space coordinates.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/convertToDeviceSpace(_:)-91x5g
func CGContextConvertRectToDeviceSpace(c CGContextRef, rect corefoundation.CGRect) corefoundation.CGRect {
	result, callErr := tryCGContextConvertRectToDeviceSpace(c, rect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextConvertRectToUserSpace func(c CGContextRef, rect corefoundation.CGRect) corefoundation.CGRect
var _cGContextConvertRectToUserSpaceErr error

func tryCGContextConvertRectToUserSpace(c CGContextRef, rect corefoundation.CGRect) (corefoundation.CGRect, error) {
	if _cGContextConvertRectToUserSpace == nil {
		return corefoundation.CGRect{}, symbolCallError("CGContextConvertRectToUserSpace", "10.4", _cGContextConvertRectToUserSpaceErr)
	}
	return _cGContextConvertRectToUserSpace(c, rect), nil
}

// CGContextConvertRectToUserSpace returns a rectangle that is transformed from device space coordinate to user space coordinates.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/convertToUserSpace(_:)-1hk5r
func CGContextConvertRectToUserSpace(c CGContextRef, rect corefoundation.CGRect) corefoundation.CGRect {
	result, callErr := tryCGContextConvertRectToUserSpace(c, rect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextConvertSizeToDeviceSpace func(c CGContextRef, size corefoundation.CGSize) corefoundation.CGSize
var _cGContextConvertSizeToDeviceSpaceErr error

func tryCGContextConvertSizeToDeviceSpace(c CGContextRef, size corefoundation.CGSize) (corefoundation.CGSize, error) {
	if _cGContextConvertSizeToDeviceSpace == nil {
		return corefoundation.CGSize{}, symbolCallError("CGContextConvertSizeToDeviceSpace", "10.4", _cGContextConvertSizeToDeviceSpaceErr)
	}
	return _cGContextConvertSizeToDeviceSpace(c, size), nil
}

// CGContextConvertSizeToDeviceSpace returns a size that is transformed from user space coordinates to device space coordinates.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/convertToDeviceSpace(_:)-224h2
func CGContextConvertSizeToDeviceSpace(c CGContextRef, size corefoundation.CGSize) corefoundation.CGSize {
	result, callErr := tryCGContextConvertSizeToDeviceSpace(c, size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextConvertSizeToUserSpace func(c CGContextRef, size corefoundation.CGSize) corefoundation.CGSize
var _cGContextConvertSizeToUserSpaceErr error

func tryCGContextConvertSizeToUserSpace(c CGContextRef, size corefoundation.CGSize) (corefoundation.CGSize, error) {
	if _cGContextConvertSizeToUserSpace == nil {
		return corefoundation.CGSize{}, symbolCallError("CGContextConvertSizeToUserSpace", "10.4", _cGContextConvertSizeToUserSpaceErr)
	}
	return _cGContextConvertSizeToUserSpace(c, size), nil
}

// CGContextConvertSizeToUserSpace returns a size that is transformed from device space coordinates to user space coordinates.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/convertToUserSpace(_:)-693ur
func CGContextConvertSizeToUserSpace(c CGContextRef, size corefoundation.CGSize) corefoundation.CGSize {
	result, callErr := tryCGContextConvertSizeToUserSpace(c, size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextCopyPath func(c CGContextRef) CGPathRef
var _cGContextCopyPathErr error

func tryCGContextCopyPath(c CGContextRef) (CGPathRef, error) {
	if _cGContextCopyPath == nil {
		return 0, symbolCallError("CGContextCopyPath", "10.2", _cGContextCopyPathErr)
	}
	return _cGContextCopyPath(c), nil
}

// CGContextCopyPath returns a path object built from the current path information in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/path
func CGContextCopyPath(c CGContextRef) CGPathRef {
	result, callErr := tryCGContextCopyPath(c)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextDrawConicGradient func(c CGContextRef, gradient CGGradientRef, center corefoundation.CGPoint, angle float64)
var _cGContextDrawConicGradientErr error

func tryCGContextDrawConicGradient(c CGContextRef, gradient CGGradientRef, center corefoundation.CGPoint, angle float64) error {
	if _cGContextDrawConicGradient == nil {
		return symbolCallError("CGContextDrawConicGradient", "14.0", _cGContextDrawConicGradientErr)
	}
	_cGContextDrawConicGradient(c, gradient, center, angle)
	return nil
}

// CGContextDrawConicGradient.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawConicGradient(_:_:_:_:)
func CGContextDrawConicGradient(c CGContextRef, gradient CGGradientRef, center corefoundation.CGPoint, angle float64) {
	if callErr := tryCGContextDrawConicGradient(c, gradient, center, angle); callErr != nil {
		panic(callErr)
	}
}

var _cGContextDrawImage func(c CGContextRef, rect corefoundation.CGRect, image CGImageRef)
var _cGContextDrawImageErr error

func tryCGContextDrawImage(c CGContextRef, rect corefoundation.CGRect, image CGImageRef) error {
	if _cGContextDrawImage == nil {
		return symbolCallError("CGContextDrawImage", "10.0", _cGContextDrawImageErr)
	}
	_cGContextDrawImage(c, rect, image)
	return nil
}

// CGContextDrawImage draws an image into a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawImage
func CGContextDrawImage(c CGContextRef, rect corefoundation.CGRect, image CGImageRef) {
	if callErr := tryCGContextDrawImage(c, rect, image); callErr != nil {
		panic(callErr)
	}
}

var _cGContextDrawImageApplyingToneMapping func(c CGContextRef, r corefoundation.CGRect, image CGImageRef, method CGToneMapping, options corefoundation.CFDictionaryRef) bool
var _cGContextDrawImageApplyingToneMappingErr error

func tryCGContextDrawImageApplyingToneMapping(c CGContextRef, r corefoundation.CGRect, image CGImageRef, method CGToneMapping, options corefoundation.CFDictionaryRef) (bool, error) {
	if _cGContextDrawImageApplyingToneMapping == nil {
		return false, symbolCallError("CGContextDrawImageApplyingToneMapping", "15.0", _cGContextDrawImageApplyingToneMappingErr)
	}
	return _cGContextDrawImageApplyingToneMapping(c, r, image, method, options), nil
}

// CGContextDrawImageApplyingToneMapping.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawImageApplyingToneMapping
func CGContextDrawImageApplyingToneMapping(c CGContextRef, r corefoundation.CGRect, image CGImageRef, method CGToneMapping, options corefoundation.CFDictionaryRef) bool {
	result, callErr := tryCGContextDrawImageApplyingToneMapping(c, r, image, method, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextDrawLayerAtPoint func(context CGContextRef, point corefoundation.CGPoint, layer unsafe.Pointer)
var _cGContextDrawLayerAtPointErr error

func tryCGContextDrawLayerAtPoint(context CGContextRef, point corefoundation.CGPoint, layer unsafe.Pointer) error {
	if _cGContextDrawLayerAtPoint == nil {
		return symbolCallError("CGContextDrawLayerAtPoint", "10.4", _cGContextDrawLayerAtPointErr)
	}
	_cGContextDrawLayerAtPoint(context, point, layer)
	return nil
}

// CGContextDrawLayerAtPoint draws the contents of a CGLayer object at the specified point.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawLayerAtPoint
func CGContextDrawLayerAtPoint(context CGContextRef, point corefoundation.CGPoint, layer unsafe.Pointer) {
	if callErr := tryCGContextDrawLayerAtPoint(context, point, layer); callErr != nil {
		panic(callErr)
	}
}

var _cGContextDrawLayerInRect func(context CGContextRef, rect corefoundation.CGRect, layer unsafe.Pointer)
var _cGContextDrawLayerInRectErr error

func tryCGContextDrawLayerInRect(context CGContextRef, rect corefoundation.CGRect, layer unsafe.Pointer) error {
	if _cGContextDrawLayerInRect == nil {
		return symbolCallError("CGContextDrawLayerInRect", "10.4", _cGContextDrawLayerInRectErr)
	}
	_cGContextDrawLayerInRect(context, rect, layer)
	return nil
}

// CGContextDrawLayerInRect draws the contents of a layer object into the specified rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawLayerInRect
func CGContextDrawLayerInRect(context CGContextRef, rect corefoundation.CGRect, layer unsafe.Pointer) {
	if callErr := tryCGContextDrawLayerInRect(context, rect, layer); callErr != nil {
		panic(callErr)
	}
}

var _cGContextDrawLinearGradient func(c CGContextRef, gradient CGGradientRef, startPoint corefoundation.CGPoint, endPoint corefoundation.CGPoint, options CGGradientDrawingOptions)
var _cGContextDrawLinearGradientErr error

func tryCGContextDrawLinearGradient(c CGContextRef, gradient CGGradientRef, startPoint corefoundation.CGPoint, endPoint corefoundation.CGPoint, options CGGradientDrawingOptions) error {
	if _cGContextDrawLinearGradient == nil {
		return symbolCallError("CGContextDrawLinearGradient", "10.5", _cGContextDrawLinearGradientErr)
	}
	_cGContextDrawLinearGradient(c, gradient, startPoint, endPoint, options)
	return nil
}

// CGContextDrawLinearGradient paints a gradient fill that varies along the line defined by the provided starting and ending points.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/drawLinearGradient(_:start:end:options:)
func CGContextDrawLinearGradient(c CGContextRef, gradient CGGradientRef, startPoint corefoundation.CGPoint, endPoint corefoundation.CGPoint, options CGGradientDrawingOptions) {
	if callErr := tryCGContextDrawLinearGradient(c, gradient, startPoint, endPoint, options); callErr != nil {
		panic(callErr)
	}
}

var _cGContextDrawPDFDocument func(c CGContextRef, rect corefoundation.CGRect, document CGPDFDocumentRef, page int)
var _cGContextDrawPDFDocumentErr error

func tryCGContextDrawPDFDocument(c CGContextRef, rect corefoundation.CGRect, document CGPDFDocumentRef, page int) error {
	if _cGContextDrawPDFDocument == nil {
		return symbolCallError("CGContextDrawPDFDocument", "10.0", _cGContextDrawPDFDocumentErr)
	}
	_cGContextDrawPDFDocument(c, rect, document, page)
	return nil
}

// CGContextDrawPDFDocument.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawPDFDocument
func CGContextDrawPDFDocument(c CGContextRef, rect corefoundation.CGRect, document CGPDFDocumentRef, page int) {
	if callErr := tryCGContextDrawPDFDocument(c, rect, document, page); callErr != nil {
		panic(callErr)
	}
}

var _cGContextDrawPDFPage func(c CGContextRef, page CGPDFPageRef)
var _cGContextDrawPDFPageErr error

func tryCGContextDrawPDFPage(c CGContextRef, page CGPDFPageRef) error {
	if _cGContextDrawPDFPage == nil {
		return symbolCallError("CGContextDrawPDFPage", "10.3", _cGContextDrawPDFPageErr)
	}
	_cGContextDrawPDFPage(c, page)
	return nil
}

// CGContextDrawPDFPage draws the content of a PDF page into the current graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/drawPDFPage(_:)
func CGContextDrawPDFPage(c CGContextRef, page CGPDFPageRef) {
	if callErr := tryCGContextDrawPDFPage(c, page); callErr != nil {
		panic(callErr)
	}
}

var _cGContextDrawPath func(c CGContextRef, mode CGPathDrawingMode)
var _cGContextDrawPathErr error

func tryCGContextDrawPath(c CGContextRef, mode CGPathDrawingMode) error {
	if _cGContextDrawPath == nil {
		return symbolCallError("CGContextDrawPath", "10.0", _cGContextDrawPathErr)
	}
	_cGContextDrawPath(c, mode)
	return nil
}

// CGContextDrawPath draws the current path using the provided drawing mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/drawPath(using:)
func CGContextDrawPath(c CGContextRef, mode CGPathDrawingMode) {
	if callErr := tryCGContextDrawPath(c, mode); callErr != nil {
		panic(callErr)
	}
}

var _cGContextDrawRadialGradient func(c CGContextRef, gradient CGGradientRef, startCenter corefoundation.CGPoint, startRadius float64, endCenter corefoundation.CGPoint, endRadius float64, options CGGradientDrawingOptions)
var _cGContextDrawRadialGradientErr error

func tryCGContextDrawRadialGradient(c CGContextRef, gradient CGGradientRef, startCenter corefoundation.CGPoint, startRadius float64, endCenter corefoundation.CGPoint, endRadius float64, options CGGradientDrawingOptions) error {
	if _cGContextDrawRadialGradient == nil {
		return symbolCallError("CGContextDrawRadialGradient", "10.5", _cGContextDrawRadialGradientErr)
	}
	_cGContextDrawRadialGradient(c, gradient, startCenter, startRadius, endCenter, endRadius, options)
	return nil
}

// CGContextDrawRadialGradient paints a gradient fill that varies along the area defined by the provided starting and ending circles.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/drawRadialGradient(_:startCenter:startRadius:endCenter:endRadius:options:)
func CGContextDrawRadialGradient(c CGContextRef, gradient CGGradientRef, startCenter corefoundation.CGPoint, startRadius float64, endCenter corefoundation.CGPoint, endRadius float64, options CGGradientDrawingOptions) {
	if callErr := tryCGContextDrawRadialGradient(c, gradient, startCenter, startRadius, endCenter, endRadius, options); callErr != nil {
		panic(callErr)
	}
}

var _cGContextDrawShading func(c CGContextRef, shading CGShadingRef)
var _cGContextDrawShadingErr error

func tryCGContextDrawShading(c CGContextRef, shading CGShadingRef) error {
	if _cGContextDrawShading == nil {
		return symbolCallError("CGContextDrawShading", "10.2", _cGContextDrawShadingErr)
	}
	_cGContextDrawShading(c, shading)
	return nil
}

// CGContextDrawShading fills the clipping path of a context with the specified shading.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/drawShading(_:)
func CGContextDrawShading(c CGContextRef, shading CGShadingRef) {
	if callErr := tryCGContextDrawShading(c, shading); callErr != nil {
		panic(callErr)
	}
}

var _cGContextDrawTiledImage func(c CGContextRef, rect corefoundation.CGRect, image CGImageRef)
var _cGContextDrawTiledImageErr error

func tryCGContextDrawTiledImage(c CGContextRef, rect corefoundation.CGRect, image CGImageRef) error {
	if _cGContextDrawTiledImage == nil {
		return symbolCallError("CGContextDrawTiledImage", "10.5", _cGContextDrawTiledImageErr)
	}
	_cGContextDrawTiledImage(c, rect, image)
	return nil
}

// CGContextDrawTiledImage repeatedly draws an image, scaled to the provided rectangle, to fill the current clip region.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextDrawTiledImage
func CGContextDrawTiledImage(c CGContextRef, rect corefoundation.CGRect, image CGImageRef) {
	if callErr := tryCGContextDrawTiledImage(c, rect, image); callErr != nil {
		panic(callErr)
	}
}

var _cGContextEOClip func(c CGContextRef)
var _cGContextEOClipErr error

func tryCGContextEOClip(c CGContextRef) error {
	if _cGContextEOClip == nil {
		return symbolCallError("CGContextEOClip", "10.0", _cGContextEOClipErr)
	}
	_cGContextEOClip(c)
	return nil
}

// CGContextEOClip modifies the current clipping path, using the even-odd rule.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextEOClip
func CGContextEOClip(c CGContextRef) {
	if callErr := tryCGContextEOClip(c); callErr != nil {
		panic(callErr)
	}
}

var _cGContextEOFillPath func(c CGContextRef)
var _cGContextEOFillPathErr error

func tryCGContextEOFillPath(c CGContextRef) error {
	if _cGContextEOFillPath == nil {
		return symbolCallError("CGContextEOFillPath", "10.0", _cGContextEOFillPathErr)
	}
	_cGContextEOFillPath(c)
	return nil
}

// CGContextEOFillPath paints the area within the current path, using the even-odd fill rule.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextEOFillPath
func CGContextEOFillPath(c CGContextRef) {
	if callErr := tryCGContextEOFillPath(c); callErr != nil {
		panic(callErr)
	}
}

var _cGContextEndPage func(c CGContextRef)
var _cGContextEndPageErr error

func tryCGContextEndPage(c CGContextRef) error {
	if _cGContextEndPage == nil {
		return symbolCallError("CGContextEndPage", "10.0", _cGContextEndPageErr)
	}
	_cGContextEndPage(c)
	return nil
}

// CGContextEndPage ends the current page in a page-based graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/endPage()
func CGContextEndPage(c CGContextRef) {
	if callErr := tryCGContextEndPage(c); callErr != nil {
		panic(callErr)
	}
}

var _cGContextEndTransparencyLayer func(c CGContextRef)
var _cGContextEndTransparencyLayerErr error

func tryCGContextEndTransparencyLayer(c CGContextRef) error {
	if _cGContextEndTransparencyLayer == nil {
		return symbolCallError("CGContextEndTransparencyLayer", "10.3", _cGContextEndTransparencyLayerErr)
	}
	_cGContextEndTransparencyLayer(c)
	return nil
}

// CGContextEndTransparencyLayer ends a transparency layer.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/endTransparencyLayer()
func CGContextEndTransparencyLayer(c CGContextRef) {
	if callErr := tryCGContextEndTransparencyLayer(c); callErr != nil {
		panic(callErr)
	}
}

var _cGContextFillEllipseInRect func(c CGContextRef, rect corefoundation.CGRect)
var _cGContextFillEllipseInRectErr error

func tryCGContextFillEllipseInRect(c CGContextRef, rect corefoundation.CGRect) error {
	if _cGContextFillEllipseInRect == nil {
		return symbolCallError("CGContextFillEllipseInRect", "10.4", _cGContextFillEllipseInRectErr)
	}
	_cGContextFillEllipseInRect(c, rect)
	return nil
}

// CGContextFillEllipseInRect paints the area of the ellipse that fits inside the provided rectangle, using the fill color in the current graphics state.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/fillEllipse(in:)
func CGContextFillEllipseInRect(c CGContextRef, rect corefoundation.CGRect) {
	if callErr := tryCGContextFillEllipseInRect(c, rect); callErr != nil {
		panic(callErr)
	}
}

var _cGContextFillPath func(c CGContextRef)
var _cGContextFillPathErr error

func tryCGContextFillPath(c CGContextRef) error {
	if _cGContextFillPath == nil {
		return symbolCallError("CGContextFillPath", "10.0", _cGContextFillPathErr)
	}
	_cGContextFillPath(c)
	return nil
}

// CGContextFillPath paints the area within the current path, using the nonzero winding number rule.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextFillPath
func CGContextFillPath(c CGContextRef) {
	if callErr := tryCGContextFillPath(c); callErr != nil {
		panic(callErr)
	}
}

var _cGContextFillRect func(c CGContextRef, rect corefoundation.CGRect)
var _cGContextFillRectErr error

func tryCGContextFillRect(c CGContextRef, rect corefoundation.CGRect) error {
	if _cGContextFillRect == nil {
		return symbolCallError("CGContextFillRect", "10.0", _cGContextFillRectErr)
	}
	_cGContextFillRect(c, rect)
	return nil
}

// CGContextFillRect paints the area contained within the provided rectangle, using the fill color in the current graphics state.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/fill(_:)-7a0rk
func CGContextFillRect(c CGContextRef, rect corefoundation.CGRect) {
	if callErr := tryCGContextFillRect(c, rect); callErr != nil {
		panic(callErr)
	}
}

var _cGContextFillRects func(c CGContextRef, rects *corefoundation.CGRect, count uintptr)
var _cGContextFillRectsErr error

func tryCGContextFillRects(c CGContextRef, rects *corefoundation.CGRect, count uintptr) error {
	if _cGContextFillRects == nil {
		return symbolCallError("CGContextFillRects", "10.0", _cGContextFillRectsErr)
	}
	_cGContextFillRects(c, rects, count)
	return nil
}

// CGContextFillRects paints the areas contained within the provided rectangles, using the fill color in the current graphics state.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextFillRects
func CGContextFillRects(c CGContextRef, rects *corefoundation.CGRect, count uintptr) {
	if callErr := tryCGContextFillRects(c, rects, count); callErr != nil {
		panic(callErr)
	}
}

var _cGContextFlush func(c CGContextRef)
var _cGContextFlushErr error

func tryCGContextFlush(c CGContextRef) error {
	if _cGContextFlush == nil {
		return symbolCallError("CGContextFlush", "10.0", _cGContextFlushErr)
	}
	_cGContextFlush(c)
	return nil
}

// CGContextFlush forces all pending drawing operations in a window context to be rendered immediately to the destination device.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/flush()
func CGContextFlush(c CGContextRef) {
	if callErr := tryCGContextFlush(c); callErr != nil {
		panic(callErr)
	}
}

var _cGContextGetCTM func(c CGContextRef) corefoundation.CGAffineTransform
var _cGContextGetCTMErr error

func tryCGContextGetCTM(c CGContextRef) (corefoundation.CGAffineTransform, error) {
	if _cGContextGetCTM == nil {
		return corefoundation.CGAffineTransform{}, symbolCallError("CGContextGetCTM", "10.0", _cGContextGetCTMErr)
	}
	return _cGContextGetCTM(c), nil
}

// CGContextGetCTM returns the current transformation matrix.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/ctm
func CGContextGetCTM(c CGContextRef) corefoundation.CGAffineTransform {
	result, callErr := tryCGContextGetCTM(c)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextGetClipBoundingBox func(c CGContextRef) corefoundation.CGRect
var _cGContextGetClipBoundingBoxErr error

func tryCGContextGetClipBoundingBox(c CGContextRef) (corefoundation.CGRect, error) {
	if _cGContextGetClipBoundingBox == nil {
		return corefoundation.CGRect{}, symbolCallError("CGContextGetClipBoundingBox", "10.3", _cGContextGetClipBoundingBoxErr)
	}
	return _cGContextGetClipBoundingBox(c), nil
}

// CGContextGetClipBoundingBox returns the bounding box of a clipping path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/boundingBoxOfClipPath
func CGContextGetClipBoundingBox(c CGContextRef) corefoundation.CGRect {
	result, callErr := tryCGContextGetClipBoundingBox(c)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextGetContentToneMappingInfo func(c CGContextRef) CGContentToneMappingInfo
var _cGContextGetContentToneMappingInfoErr error

func tryCGContextGetContentToneMappingInfo(c CGContextRef) (CGContentToneMappingInfo, error) {
	if _cGContextGetContentToneMappingInfo == nil {
		return CGContentToneMappingInfo{}, symbolCallError("CGContextGetContentToneMappingInfo", "26.0", _cGContextGetContentToneMappingInfoErr)
	}
	return _cGContextGetContentToneMappingInfo(c), nil
}

// CGContextGetContentToneMappingInfo.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextGetContentToneMappingInfo
func CGContextGetContentToneMappingInfo(c CGContextRef) CGContentToneMappingInfo {
	result, callErr := tryCGContextGetContentToneMappingInfo(c)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextGetEDRTargetHeadroom func(c CGContextRef) float32
var _cGContextGetEDRTargetHeadroomErr error

func tryCGContextGetEDRTargetHeadroom(c CGContextRef) (float32, error) {
	if _cGContextGetEDRTargetHeadroom == nil {
		return 0.0, symbolCallError("CGContextGetEDRTargetHeadroom", "15.0", _cGContextGetEDRTargetHeadroomErr)
	}
	return _cGContextGetEDRTargetHeadroom(c), nil
}

// CGContextGetEDRTargetHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextGetEDRTargetHeadroom(_:)
func CGContextGetEDRTargetHeadroom(c CGContextRef) float32 {
	result, callErr := tryCGContextGetEDRTargetHeadroom(c)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextGetInterpolationQuality func(c CGContextRef) CGInterpolationQuality
var _cGContextGetInterpolationQualityErr error

func tryCGContextGetInterpolationQuality(c CGContextRef) (CGInterpolationQuality, error) {
	if _cGContextGetInterpolationQuality == nil {
		return *new(CGInterpolationQuality), symbolCallError("CGContextGetInterpolationQuality", "10.0", _cGContextGetInterpolationQualityErr)
	}
	return _cGContextGetInterpolationQuality(c), nil
}

// CGContextGetInterpolationQuality returns the current level of interpolation quality for a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/interpolationQuality
func CGContextGetInterpolationQuality(c CGContextRef) CGInterpolationQuality {
	result, callErr := tryCGContextGetInterpolationQuality(c)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextGetPathBoundingBox func(c CGContextRef) corefoundation.CGRect
var _cGContextGetPathBoundingBoxErr error

func tryCGContextGetPathBoundingBox(c CGContextRef) (corefoundation.CGRect, error) {
	if _cGContextGetPathBoundingBox == nil {
		return corefoundation.CGRect{}, symbolCallError("CGContextGetPathBoundingBox", "10.0", _cGContextGetPathBoundingBoxErr)
	}
	return _cGContextGetPathBoundingBox(c), nil
}

// CGContextGetPathBoundingBox returns the smallest rectangle that contains the current path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/boundingBoxOfPath
func CGContextGetPathBoundingBox(c CGContextRef) corefoundation.CGRect {
	result, callErr := tryCGContextGetPathBoundingBox(c)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextGetPathCurrentPoint func(c CGContextRef) corefoundation.CGPoint
var _cGContextGetPathCurrentPointErr error

func tryCGContextGetPathCurrentPoint(c CGContextRef) (corefoundation.CGPoint, error) {
	if _cGContextGetPathCurrentPoint == nil {
		return corefoundation.CGPoint{}, symbolCallError("CGContextGetPathCurrentPoint", "10.0", _cGContextGetPathCurrentPointErr)
	}
	return _cGContextGetPathCurrentPoint(c), nil
}

// CGContextGetPathCurrentPoint returns the current point in a non-empty path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/currentPointOfPath
func CGContextGetPathCurrentPoint(c CGContextRef) corefoundation.CGPoint {
	result, callErr := tryCGContextGetPathCurrentPoint(c)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextGetTextMatrix func(c CGContextRef) corefoundation.CGAffineTransform
var _cGContextGetTextMatrixErr error

func tryCGContextGetTextMatrix(c CGContextRef) (corefoundation.CGAffineTransform, error) {
	if _cGContextGetTextMatrix == nil {
		return corefoundation.CGAffineTransform{}, symbolCallError("CGContextGetTextMatrix", "10.0", _cGContextGetTextMatrixErr)
	}
	return _cGContextGetTextMatrix(c), nil
}

// CGContextGetTextMatrix returns the current text matrix.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/textMatrix
func CGContextGetTextMatrix(c CGContextRef) corefoundation.CGAffineTransform {
	result, callErr := tryCGContextGetTextMatrix(c)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextGetTextPosition func(c CGContextRef) corefoundation.CGPoint
var _cGContextGetTextPositionErr error

func tryCGContextGetTextPosition(c CGContextRef) (corefoundation.CGPoint, error) {
	if _cGContextGetTextPosition == nil {
		return corefoundation.CGPoint{}, symbolCallError("CGContextGetTextPosition", "10.0", _cGContextGetTextPositionErr)
	}
	return _cGContextGetTextPosition(c), nil
}

// CGContextGetTextPosition.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextGetTextPosition
func CGContextGetTextPosition(c CGContextRef) corefoundation.CGPoint {
	result, callErr := tryCGContextGetTextPosition(c)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextGetTypeID func() uint
var _cGContextGetTypeIDErr error

func tryCGContextGetTypeID() (uint, error) {
	if _cGContextGetTypeID == nil {
		return 0, symbolCallError("CGContextGetTypeID", "10.2", _cGContextGetTypeIDErr)
	}
	return _cGContextGetTypeID(), nil
}

// CGContextGetTypeID returns the type identifier for a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/typeID
func CGContextGetTypeID() uint {
	result, callErr := tryCGContextGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextGetUserSpaceToDeviceSpaceTransform func(c CGContextRef) corefoundation.CGAffineTransform
var _cGContextGetUserSpaceToDeviceSpaceTransformErr error

func tryCGContextGetUserSpaceToDeviceSpaceTransform(c CGContextRef) (corefoundation.CGAffineTransform, error) {
	if _cGContextGetUserSpaceToDeviceSpaceTransform == nil {
		return corefoundation.CGAffineTransform{}, symbolCallError("CGContextGetUserSpaceToDeviceSpaceTransform", "10.4", _cGContextGetUserSpaceToDeviceSpaceTransformErr)
	}
	return _cGContextGetUserSpaceToDeviceSpaceTransform(c), nil
}

// CGContextGetUserSpaceToDeviceSpaceTransform returns an affine transform that maps user space coordinates to device space coordinates.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/userSpaceToDeviceSpaceTransform
func CGContextGetUserSpaceToDeviceSpaceTransform(c CGContextRef) corefoundation.CGAffineTransform {
	result, callErr := tryCGContextGetUserSpaceToDeviceSpaceTransform(c)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextIsPathEmpty func(c CGContextRef) bool
var _cGContextIsPathEmptyErr error

func tryCGContextIsPathEmpty(c CGContextRef) (bool, error) {
	if _cGContextIsPathEmpty == nil {
		return false, symbolCallError("CGContextIsPathEmpty", "10.0", _cGContextIsPathEmptyErr)
	}
	return _cGContextIsPathEmpty(c), nil
}

// CGContextIsPathEmpty indicates whether the current path contains any subpaths.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/isPathEmpty
func CGContextIsPathEmpty(c CGContextRef) bool {
	result, callErr := tryCGContextIsPathEmpty(c)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextMoveToPoint func(c CGContextRef, x float64, y float64)
var _cGContextMoveToPointErr error

func tryCGContextMoveToPoint(c CGContextRef, x float64, y float64) error {
	if _cGContextMoveToPoint == nil {
		return symbolCallError("CGContextMoveToPoint", "10.0", _cGContextMoveToPointErr)
	}
	_cGContextMoveToPoint(c, x, y)
	return nil
}

// CGContextMoveToPoint begins a new subpath at the point you specify.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextMoveToPoint
func CGContextMoveToPoint(c CGContextRef, x float64, y float64) {
	if callErr := tryCGContextMoveToPoint(c, x, y); callErr != nil {
		panic(callErr)
	}
}

var _cGContextPathContainsPoint func(c CGContextRef, point corefoundation.CGPoint, mode CGPathDrawingMode) bool
var _cGContextPathContainsPointErr error

func tryCGContextPathContainsPoint(c CGContextRef, point corefoundation.CGPoint, mode CGPathDrawingMode) (bool, error) {
	if _cGContextPathContainsPoint == nil {
		return false, symbolCallError("CGContextPathContainsPoint", "10.4", _cGContextPathContainsPointErr)
	}
	return _cGContextPathContainsPoint(c, point, mode), nil
}

// CGContextPathContainsPoint checks to see whether the specified point is contained in the current path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/pathContains(_:mode:)
func CGContextPathContainsPoint(c CGContextRef, point corefoundation.CGPoint, mode CGPathDrawingMode) bool {
	result, callErr := tryCGContextPathContainsPoint(c, point, mode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextRelease func(c CGContextRef)
var _cGContextReleaseErr error

func tryCGContextRelease(c CGContextRef) error {
	if _cGContextRelease == nil {
		return symbolCallError("CGContextRelease", "10.0", _cGContextReleaseErr)
	}
	_cGContextRelease(c)
	return nil
}

// CGContextRelease decrements the retain count of a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextRelease
func CGContextRelease(c CGContextRef) {
	if callErr := tryCGContextRelease(c); callErr != nil {
		panic(callErr)
	}
}

var _cGContextReplacePathWithStrokedPath func(c CGContextRef)
var _cGContextReplacePathWithStrokedPathErr error

func tryCGContextReplacePathWithStrokedPath(c CGContextRef) error {
	if _cGContextReplacePathWithStrokedPath == nil {
		return symbolCallError("CGContextReplacePathWithStrokedPath", "10.4", _cGContextReplacePathWithStrokedPathErr)
	}
	_cGContextReplacePathWithStrokedPath(c)
	return nil
}

// CGContextReplacePathWithStrokedPath replaces the path in the graphics context with the stroked version of the path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/replacePathWithStrokedPath()
func CGContextReplacePathWithStrokedPath(c CGContextRef) {
	if callErr := tryCGContextReplacePathWithStrokedPath(c); callErr != nil {
		panic(callErr)
	}
}

var _cGContextResetClip func(c CGContextRef)
var _cGContextResetClipErr error

func tryCGContextResetClip(c CGContextRef) error {
	if _cGContextResetClip == nil {
		return symbolCallError("CGContextResetClip", "", _cGContextResetClipErr)
	}
	_cGContextResetClip(c)
	return nil
}

// CGContextResetClip.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/resetClip()
func CGContextResetClip(c CGContextRef) {
	if callErr := tryCGContextResetClip(c); callErr != nil {
		panic(callErr)
	}
}

var _cGContextRestoreGState func(c CGContextRef)
var _cGContextRestoreGStateErr error

func tryCGContextRestoreGState(c CGContextRef) error {
	if _cGContextRestoreGState == nil {
		return symbolCallError("CGContextRestoreGState", "10.0", _cGContextRestoreGStateErr)
	}
	_cGContextRestoreGState(c)
	return nil
}

// CGContextRestoreGState sets the current graphics state to the state most recently saved.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/restoreGState()
func CGContextRestoreGState(c CGContextRef) {
	if callErr := tryCGContextRestoreGState(c); callErr != nil {
		panic(callErr)
	}
}

var _cGContextRetain func(c CGContextRef) CGContextRef
var _cGContextRetainErr error

func tryCGContextRetain(c CGContextRef) (CGContextRef, error) {
	if _cGContextRetain == nil {
		return 0, symbolCallError("CGContextRetain", "10.0", _cGContextRetainErr)
	}
	return _cGContextRetain(c), nil
}

// CGContextRetain increments the retain count of a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextRetain
func CGContextRetain(c CGContextRef) CGContextRef {
	result, callErr := tryCGContextRetain(c)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextRotateCTM func(c CGContextRef, angle float64)
var _cGContextRotateCTMErr error

func tryCGContextRotateCTM(c CGContextRef, angle float64) error {
	if _cGContextRotateCTM == nil {
		return symbolCallError("CGContextRotateCTM", "10.0", _cGContextRotateCTMErr)
	}
	_cGContextRotateCTM(c, angle)
	return nil
}

// CGContextRotateCTM rotates the user coordinate system in a context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/rotate(by:)
func CGContextRotateCTM(c CGContextRef, angle float64) {
	if callErr := tryCGContextRotateCTM(c, angle); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSaveGState func(c CGContextRef)
var _cGContextSaveGStateErr error

func tryCGContextSaveGState(c CGContextRef) error {
	if _cGContextSaveGState == nil {
		return symbolCallError("CGContextSaveGState", "10.0", _cGContextSaveGStateErr)
	}
	_cGContextSaveGState(c)
	return nil
}

// CGContextSaveGState pushes a copy of the current graphics state onto the graphics state stack for the context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/saveGState()
func CGContextSaveGState(c CGContextRef) {
	if callErr := tryCGContextSaveGState(c); callErr != nil {
		panic(callErr)
	}
}

var _cGContextScaleCTM func(c CGContextRef, sx float64, sy float64)
var _cGContextScaleCTMErr error

func tryCGContextScaleCTM(c CGContextRef, sx float64, sy float64) error {
	if _cGContextScaleCTM == nil {
		return symbolCallError("CGContextScaleCTM", "10.0", _cGContextScaleCTMErr)
	}
	_cGContextScaleCTM(c, sx, sy)
	return nil
}

// CGContextScaleCTM changes the scale of the user coordinate system in a context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/scaleBy(x:y:)
func CGContextScaleCTM(c CGContextRef, sx float64, sy float64) {
	if callErr := tryCGContextScaleCTM(c, sx, sy); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetAllowsAntialiasing func(c CGContextRef, allowsAntialiasing bool)
var _cGContextSetAllowsAntialiasingErr error

func tryCGContextSetAllowsAntialiasing(c CGContextRef, allowsAntialiasing bool) error {
	if _cGContextSetAllowsAntialiasing == nil {
		return symbolCallError("CGContextSetAllowsAntialiasing", "10.4", _cGContextSetAllowsAntialiasingErr)
	}
	_cGContextSetAllowsAntialiasing(c, allowsAntialiasing)
	return nil
}

// CGContextSetAllowsAntialiasing sets whether or not to allow antialiasing for a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setAllowsAntialiasing(_:)
func CGContextSetAllowsAntialiasing(c CGContextRef, allowsAntialiasing bool) {
	if callErr := tryCGContextSetAllowsAntialiasing(c, allowsAntialiasing); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetAllowsFontSmoothing func(c CGContextRef, allowsFontSmoothing bool)
var _cGContextSetAllowsFontSmoothingErr error

func tryCGContextSetAllowsFontSmoothing(c CGContextRef, allowsFontSmoothing bool) error {
	if _cGContextSetAllowsFontSmoothing == nil {
		return symbolCallError("CGContextSetAllowsFontSmoothing", "10.2", _cGContextSetAllowsFontSmoothingErr)
	}
	_cGContextSetAllowsFontSmoothing(c, allowsFontSmoothing)
	return nil
}

// CGContextSetAllowsFontSmoothing sets whether or not to allow font smoothing for a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setAllowsFontSmoothing(_:)
func CGContextSetAllowsFontSmoothing(c CGContextRef, allowsFontSmoothing bool) {
	if callErr := tryCGContextSetAllowsFontSmoothing(c, allowsFontSmoothing); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetAllowsFontSubpixelPositioning func(c CGContextRef, allowsFontSubpixelPositioning bool)
var _cGContextSetAllowsFontSubpixelPositioningErr error

func tryCGContextSetAllowsFontSubpixelPositioning(c CGContextRef, allowsFontSubpixelPositioning bool) error {
	if _cGContextSetAllowsFontSubpixelPositioning == nil {
		return symbolCallError("CGContextSetAllowsFontSubpixelPositioning", "10.5", _cGContextSetAllowsFontSubpixelPositioningErr)
	}
	_cGContextSetAllowsFontSubpixelPositioning(c, allowsFontSubpixelPositioning)
	return nil
}

// CGContextSetAllowsFontSubpixelPositioning sets whether or not to allow subpixel positioning for a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setAllowsFontSubpixelPositioning(_:)
func CGContextSetAllowsFontSubpixelPositioning(c CGContextRef, allowsFontSubpixelPositioning bool) {
	if callErr := tryCGContextSetAllowsFontSubpixelPositioning(c, allowsFontSubpixelPositioning); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetAllowsFontSubpixelQuantization func(c CGContextRef, allowsFontSubpixelQuantization bool)
var _cGContextSetAllowsFontSubpixelQuantizationErr error

func tryCGContextSetAllowsFontSubpixelQuantization(c CGContextRef, allowsFontSubpixelQuantization bool) error {
	if _cGContextSetAllowsFontSubpixelQuantization == nil {
		return symbolCallError("CGContextSetAllowsFontSubpixelQuantization", "10.5", _cGContextSetAllowsFontSubpixelQuantizationErr)
	}
	_cGContextSetAllowsFontSubpixelQuantization(c, allowsFontSubpixelQuantization)
	return nil
}

// CGContextSetAllowsFontSubpixelQuantization sets whether or not to allow subpixel quantization for a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setAllowsFontSubpixelQuantization(_:)
func CGContextSetAllowsFontSubpixelQuantization(c CGContextRef, allowsFontSubpixelQuantization bool) {
	if callErr := tryCGContextSetAllowsFontSubpixelQuantization(c, allowsFontSubpixelQuantization); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetAlpha func(c CGContextRef, alpha float64)
var _cGContextSetAlphaErr error

func tryCGContextSetAlpha(c CGContextRef, alpha float64) error {
	if _cGContextSetAlpha == nil {
		return symbolCallError("CGContextSetAlpha", "10.0", _cGContextSetAlphaErr)
	}
	_cGContextSetAlpha(c, alpha)
	return nil
}

// CGContextSetAlpha sets the opacity level for objects drawn in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setAlpha(_:)
func CGContextSetAlpha(c CGContextRef, alpha float64) {
	if callErr := tryCGContextSetAlpha(c, alpha); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetBlendMode func(c CGContextRef, mode uint)
var _cGContextSetBlendModeErr error

func tryCGContextSetBlendMode(c CGContextRef, mode uint) error {
	if _cGContextSetBlendMode == nil {
		return symbolCallError("CGContextSetBlendMode", "10.4", _cGContextSetBlendModeErr)
	}
	_cGContextSetBlendMode(c, mode)
	return nil
}

// CGContextSetBlendMode sets how sample values are composited by a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setBlendMode(_:)
func CGContextSetBlendMode(c CGContextRef, mode uint) {
	if callErr := tryCGContextSetBlendMode(c, mode); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetCMYKFillColor func(c CGContextRef, cyan float64, magenta float64, yellow float64, black float64, alpha float64)
var _cGContextSetCMYKFillColorErr error

func tryCGContextSetCMYKFillColor(c CGContextRef, cyan float64, magenta float64, yellow float64, black float64, alpha float64) error {
	if _cGContextSetCMYKFillColor == nil {
		return symbolCallError("CGContextSetCMYKFillColor", "10.0", _cGContextSetCMYKFillColorErr)
	}
	_cGContextSetCMYKFillColor(c, cyan, magenta, yellow, black, alpha)
	return nil
}

// CGContextSetCMYKFillColor sets the current fill color to a value in the DeviceCMYK color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillColor(cyan:magenta:yellow:black:alpha:)
func CGContextSetCMYKFillColor(c CGContextRef, cyan float64, magenta float64, yellow float64, black float64, alpha float64) {
	if callErr := tryCGContextSetCMYKFillColor(c, cyan, magenta, yellow, black, alpha); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetCMYKStrokeColor func(c CGContextRef, cyan float64, magenta float64, yellow float64, black float64, alpha float64)
var _cGContextSetCMYKStrokeColorErr error

func tryCGContextSetCMYKStrokeColor(c CGContextRef, cyan float64, magenta float64, yellow float64, black float64, alpha float64) error {
	if _cGContextSetCMYKStrokeColor == nil {
		return symbolCallError("CGContextSetCMYKStrokeColor", "10.0", _cGContextSetCMYKStrokeColorErr)
	}
	_cGContextSetCMYKStrokeColor(c, cyan, magenta, yellow, black, alpha)
	return nil
}

// CGContextSetCMYKStrokeColor sets the current stroke color to a value in the DeviceCMYK color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokeColor(cyan:magenta:yellow:black:alpha:)
func CGContextSetCMYKStrokeColor(c CGContextRef, cyan float64, magenta float64, yellow float64, black float64, alpha float64) {
	if callErr := tryCGContextSetCMYKStrokeColor(c, cyan, magenta, yellow, black, alpha); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetCharacterSpacing func(c CGContextRef, spacing float64)
var _cGContextSetCharacterSpacingErr error

func tryCGContextSetCharacterSpacing(c CGContextRef, spacing float64) error {
	if _cGContextSetCharacterSpacing == nil {
		return symbolCallError("CGContextSetCharacterSpacing", "10.0", _cGContextSetCharacterSpacingErr)
	}
	_cGContextSetCharacterSpacing(c, spacing)
	return nil
}

// CGContextSetCharacterSpacing sets the current character spacing.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setCharacterSpacing(_:)
func CGContextSetCharacterSpacing(c CGContextRef, spacing float64) {
	if callErr := tryCGContextSetCharacterSpacing(c, spacing); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetContentToneMappingInfo func(c CGContextRef, info CGContentToneMappingInfo)
var _cGContextSetContentToneMappingInfoErr error

func tryCGContextSetContentToneMappingInfo(c CGContextRef, info CGContentToneMappingInfo) error {
	if _cGContextSetContentToneMappingInfo == nil {
		return symbolCallError("CGContextSetContentToneMappingInfo", "26.0", _cGContextSetContentToneMappingInfoErr)
	}
	_cGContextSetContentToneMappingInfo(c, info)
	return nil
}

// CGContextSetContentToneMappingInfo.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextSetContentToneMappingInfo
func CGContextSetContentToneMappingInfo(c CGContextRef, info CGContentToneMappingInfo) {
	if callErr := tryCGContextSetContentToneMappingInfo(c, info); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetEDRTargetHeadroom func(c CGContextRef, headroom float32) bool
var _cGContextSetEDRTargetHeadroomErr error

func tryCGContextSetEDRTargetHeadroom(c CGContextRef, headroom float32) (bool, error) {
	if _cGContextSetEDRTargetHeadroom == nil {
		return false, symbolCallError("CGContextSetEDRTargetHeadroom", "15.0", _cGContextSetEDRTargetHeadroomErr)
	}
	return _cGContextSetEDRTargetHeadroom(c, headroom), nil
}

// CGContextSetEDRTargetHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setEDRTargetHeadroom(_:)
func CGContextSetEDRTargetHeadroom(c CGContextRef, headroom float32) bool {
	result, callErr := tryCGContextSetEDRTargetHeadroom(c, headroom)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGContextSetFillColor func(c CGContextRef, components *float64)
var _cGContextSetFillColorErr error

func tryCGContextSetFillColor(c CGContextRef, components *float64) error {
	if _cGContextSetFillColor == nil {
		return symbolCallError("CGContextSetFillColor", "10.0", _cGContextSetFillColorErr)
	}
	_cGContextSetFillColor(c, components)
	return nil
}

// CGContextSetFillColor sets the current fill color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillColor(_:)-756dy
func CGContextSetFillColor(c CGContextRef, components *float64) {
	if callErr := tryCGContextSetFillColor(c, components); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetFillColorSpace func(c CGContextRef, space CGColorSpaceRef)
var _cGContextSetFillColorSpaceErr error

func tryCGContextSetFillColorSpace(c CGContextRef, space CGColorSpaceRef) error {
	if _cGContextSetFillColorSpace == nil {
		return symbolCallError("CGContextSetFillColorSpace", "10.0", _cGContextSetFillColorSpaceErr)
	}
	_cGContextSetFillColorSpace(c, space)
	return nil
}

// CGContextSetFillColorSpace sets the fill color space in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillColorSpace(_:)
func CGContextSetFillColorSpace(c CGContextRef, space CGColorSpaceRef) {
	if callErr := tryCGContextSetFillColorSpace(c, space); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetFillColorWithColor func(c CGContextRef, color CGColorRef)
var _cGContextSetFillColorWithColorErr error

func tryCGContextSetFillColorWithColor(c CGContextRef, color CGColorRef) error {
	if _cGContextSetFillColorWithColor == nil {
		return symbolCallError("CGContextSetFillColorWithColor", "10.3", _cGContextSetFillColorWithColorErr)
	}
	_cGContextSetFillColorWithColor(c, color)
	return nil
}

// CGContextSetFillColorWithColor sets the current fill color in a graphics context, using a CGColor.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillColor(_:)-8lhn8
func CGContextSetFillColorWithColor(c CGContextRef, color CGColorRef) {
	if callErr := tryCGContextSetFillColorWithColor(c, color); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetFillPattern func(c CGContextRef, pattern CGPatternRef, components *float64)
var _cGContextSetFillPatternErr error

func tryCGContextSetFillPattern(c CGContextRef, pattern CGPatternRef, components *float64) error {
	if _cGContextSetFillPattern == nil {
		return symbolCallError("CGContextSetFillPattern", "10.0", _cGContextSetFillPatternErr)
	}
	_cGContextSetFillPattern(c, pattern, components)
	return nil
}

// CGContextSetFillPattern sets the fill pattern in the specified graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillPattern(_:colorComponents:)
func CGContextSetFillPattern(c CGContextRef, pattern CGPatternRef, components *float64) {
	if callErr := tryCGContextSetFillPattern(c, pattern, components); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetFlatness func(c CGContextRef, flatness float64)
var _cGContextSetFlatnessErr error

func tryCGContextSetFlatness(c CGContextRef, flatness float64) error {
	if _cGContextSetFlatness == nil {
		return symbolCallError("CGContextSetFlatness", "10.0", _cGContextSetFlatnessErr)
	}
	_cGContextSetFlatness(c, flatness)
	return nil
}

// CGContextSetFlatness sets the accuracy of curved paths in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFlatness(_:)
func CGContextSetFlatness(c CGContextRef, flatness float64) {
	if callErr := tryCGContextSetFlatness(c, flatness); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetFont func(c CGContextRef, font CGFontRef)
var _cGContextSetFontErr error

func tryCGContextSetFont(c CGContextRef, font CGFontRef) error {
	if _cGContextSetFont == nil {
		return symbolCallError("CGContextSetFont", "10.0", _cGContextSetFontErr)
	}
	_cGContextSetFont(c, font)
	return nil
}

// CGContextSetFont sets the platform font in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFont(_:)
func CGContextSetFont(c CGContextRef, font CGFontRef) {
	if callErr := tryCGContextSetFont(c, font); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetFontSize func(c CGContextRef, size float64)
var _cGContextSetFontSizeErr error

func tryCGContextSetFontSize(c CGContextRef, size float64) error {
	if _cGContextSetFontSize == nil {
		return symbolCallError("CGContextSetFontSize", "10.0", _cGContextSetFontSizeErr)
	}
	_cGContextSetFontSize(c, size)
	return nil
}

// CGContextSetFontSize sets the current font size.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFontSize(_:)
func CGContextSetFontSize(c CGContextRef, size float64) {
	if callErr := tryCGContextSetFontSize(c, size); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetGrayFillColor func(c CGContextRef, gray float64, alpha float64)
var _cGContextSetGrayFillColorErr error

func tryCGContextSetGrayFillColor(c CGContextRef, gray float64, alpha float64) error {
	if _cGContextSetGrayFillColor == nil {
		return symbolCallError("CGContextSetGrayFillColor", "10.0", _cGContextSetGrayFillColorErr)
	}
	_cGContextSetGrayFillColor(c, gray, alpha)
	return nil
}

// CGContextSetGrayFillColor sets the current fill color to a value in the DeviceGray color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillColor(gray:alpha:)
func CGContextSetGrayFillColor(c CGContextRef, gray float64, alpha float64) {
	if callErr := tryCGContextSetGrayFillColor(c, gray, alpha); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetGrayStrokeColor func(c CGContextRef, gray float64, alpha float64)
var _cGContextSetGrayStrokeColorErr error

func tryCGContextSetGrayStrokeColor(c CGContextRef, gray float64, alpha float64) error {
	if _cGContextSetGrayStrokeColor == nil {
		return symbolCallError("CGContextSetGrayStrokeColor", "10.0", _cGContextSetGrayStrokeColorErr)
	}
	_cGContextSetGrayStrokeColor(c, gray, alpha)
	return nil
}

// CGContextSetGrayStrokeColor sets the current stroke color to a value in the DeviceGray color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokeColor(gray:alpha:)
func CGContextSetGrayStrokeColor(c CGContextRef, gray float64, alpha float64) {
	if callErr := tryCGContextSetGrayStrokeColor(c, gray, alpha); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetInterpolationQuality func(c CGContextRef, quality CGInterpolationQuality)
var _cGContextSetInterpolationQualityErr error

func tryCGContextSetInterpolationQuality(c CGContextRef, quality CGInterpolationQuality) error {
	if _cGContextSetInterpolationQuality == nil {
		return symbolCallError("CGContextSetInterpolationQuality", "10.0", _cGContextSetInterpolationQualityErr)
	}
	_cGContextSetInterpolationQuality(c, quality)
	return nil
}

// CGContextSetInterpolationQuality sets the level of interpolation quality for a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextSetInterpolationQuality
func CGContextSetInterpolationQuality(c CGContextRef, quality CGInterpolationQuality) {
	if callErr := tryCGContextSetInterpolationQuality(c, quality); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetLineCap func(c CGContextRef, cap_ uint)
var _cGContextSetLineCapErr error

func tryCGContextSetLineCap(c CGContextRef, cap_ uint) error {
	if _cGContextSetLineCap == nil {
		return symbolCallError("CGContextSetLineCap", "10.0", _cGContextSetLineCapErr)
	}
	_cGContextSetLineCap(c, cap_)
	return nil
}

// CGContextSetLineCap sets the style for the endpoints of lines drawn in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setLineCap(_:)
func CGContextSetLineCap(c CGContextRef, cap_ uint) {
	if callErr := tryCGContextSetLineCap(c, cap_); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetLineDash func(c CGContextRef, phase float64, lengths *float64, count uintptr)
var _cGContextSetLineDashErr error

func tryCGContextSetLineDash(c CGContextRef, phase float64, lengths *float64, count uintptr) error {
	if _cGContextSetLineDash == nil {
		return symbolCallError("CGContextSetLineDash", "10.0", _cGContextSetLineDashErr)
	}
	_cGContextSetLineDash(c, phase, lengths, count)
	return nil
}

// CGContextSetLineDash sets the pattern for dashed lines in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextSetLineDash
func CGContextSetLineDash(c CGContextRef, phase float64, lengths *float64, count uintptr) {
	if callErr := tryCGContextSetLineDash(c, phase, lengths, count); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetLineJoin func(c CGContextRef, join uint)
var _cGContextSetLineJoinErr error

func tryCGContextSetLineJoin(c CGContextRef, join uint) error {
	if _cGContextSetLineJoin == nil {
		return symbolCallError("CGContextSetLineJoin", "10.0", _cGContextSetLineJoinErr)
	}
	_cGContextSetLineJoin(c, join)
	return nil
}

// CGContextSetLineJoin sets the style for the joins of connected lines in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setLineJoin(_:)
func CGContextSetLineJoin(c CGContextRef, join uint) {
	if callErr := tryCGContextSetLineJoin(c, join); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetLineWidth func(c CGContextRef, width float64)
var _cGContextSetLineWidthErr error

func tryCGContextSetLineWidth(c CGContextRef, width float64) error {
	if _cGContextSetLineWidth == nil {
		return symbolCallError("CGContextSetLineWidth", "10.0", _cGContextSetLineWidthErr)
	}
	_cGContextSetLineWidth(c, width)
	return nil
}

// CGContextSetLineWidth sets the line width for a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setLineWidth(_:)
func CGContextSetLineWidth(c CGContextRef, width float64) {
	if callErr := tryCGContextSetLineWidth(c, width); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetMiterLimit func(c CGContextRef, limit float64)
var _cGContextSetMiterLimitErr error

func tryCGContextSetMiterLimit(c CGContextRef, limit float64) error {
	if _cGContextSetMiterLimit == nil {
		return symbolCallError("CGContextSetMiterLimit", "10.0", _cGContextSetMiterLimitErr)
	}
	_cGContextSetMiterLimit(c, limit)
	return nil
}

// CGContextSetMiterLimit sets the miter limit for the joins of connected lines in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setMiterLimit(_:)
func CGContextSetMiterLimit(c CGContextRef, limit float64) {
	if callErr := tryCGContextSetMiterLimit(c, limit); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetPatternPhase func(c CGContextRef, phase corefoundation.CGSize)
var _cGContextSetPatternPhaseErr error

func tryCGContextSetPatternPhase(c CGContextRef, phase corefoundation.CGSize) error {
	if _cGContextSetPatternPhase == nil {
		return symbolCallError("CGContextSetPatternPhase", "10.0", _cGContextSetPatternPhaseErr)
	}
	_cGContextSetPatternPhase(c, phase)
	return nil
}

// CGContextSetPatternPhase sets the pattern phase of a context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setPatternPhase(_:)
func CGContextSetPatternPhase(c CGContextRef, phase corefoundation.CGSize) {
	if callErr := tryCGContextSetPatternPhase(c, phase); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetRGBFillColor func(c CGContextRef, red float64, green float64, blue float64, alpha float64)
var _cGContextSetRGBFillColorErr error

func tryCGContextSetRGBFillColor(c CGContextRef, red float64, green float64, blue float64, alpha float64) error {
	if _cGContextSetRGBFillColor == nil {
		return symbolCallError("CGContextSetRGBFillColor", "10.0", _cGContextSetRGBFillColorErr)
	}
	_cGContextSetRGBFillColor(c, red, green, blue, alpha)
	return nil
}

// CGContextSetRGBFillColor sets the current fill color to a value in the DeviceRGB color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setFillColor(red:green:blue:alpha:)
func CGContextSetRGBFillColor(c CGContextRef, red float64, green float64, blue float64, alpha float64) {
	if callErr := tryCGContextSetRGBFillColor(c, red, green, blue, alpha); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetRGBStrokeColor func(c CGContextRef, red float64, green float64, blue float64, alpha float64)
var _cGContextSetRGBStrokeColorErr error

func tryCGContextSetRGBStrokeColor(c CGContextRef, red float64, green float64, blue float64, alpha float64) error {
	if _cGContextSetRGBStrokeColor == nil {
		return symbolCallError("CGContextSetRGBStrokeColor", "10.0", _cGContextSetRGBStrokeColorErr)
	}
	_cGContextSetRGBStrokeColor(c, red, green, blue, alpha)
	return nil
}

// CGContextSetRGBStrokeColor sets the current stroke color to a value in the DeviceRGB color space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokeColor(red:green:blue:alpha:)
func CGContextSetRGBStrokeColor(c CGContextRef, red float64, green float64, blue float64, alpha float64) {
	if callErr := tryCGContextSetRGBStrokeColor(c, red, green, blue, alpha); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetRenderingIntent func(c CGContextRef, intent CGColorRenderingIntent)
var _cGContextSetRenderingIntentErr error

func tryCGContextSetRenderingIntent(c CGContextRef, intent CGColorRenderingIntent) error {
	if _cGContextSetRenderingIntent == nil {
		return symbolCallError("CGContextSetRenderingIntent", "10.0", _cGContextSetRenderingIntentErr)
	}
	_cGContextSetRenderingIntent(c, intent)
	return nil
}

// CGContextSetRenderingIntent sets the rendering intent in the current graphics state.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setRenderingIntent(_:)
func CGContextSetRenderingIntent(c CGContextRef, intent CGColorRenderingIntent) {
	if callErr := tryCGContextSetRenderingIntent(c, intent); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetShadow func(c CGContextRef, offset corefoundation.CGSize, blur float64)
var _cGContextSetShadowErr error

func tryCGContextSetShadow(c CGContextRef, offset corefoundation.CGSize, blur float64) error {
	if _cGContextSetShadow == nil {
		return symbolCallError("CGContextSetShadow", "10.3", _cGContextSetShadowErr)
	}
	_cGContextSetShadow(c, offset, blur)
	return nil
}

// CGContextSetShadow enables shadowing in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setShadow(offset:blur:)
func CGContextSetShadow(c CGContextRef, offset corefoundation.CGSize, blur float64) {
	if callErr := tryCGContextSetShadow(c, offset, blur); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetShadowWithColor func(c CGContextRef, offset corefoundation.CGSize, blur float64, color CGColorRef)
var _cGContextSetShadowWithColorErr error

func tryCGContextSetShadowWithColor(c CGContextRef, offset corefoundation.CGSize, blur float64, color CGColorRef) error {
	if _cGContextSetShadowWithColor == nil {
		return symbolCallError("CGContextSetShadowWithColor", "10.3", _cGContextSetShadowWithColorErr)
	}
	_cGContextSetShadowWithColor(c, offset, blur, color)
	return nil
}

// CGContextSetShadowWithColor enables shadowing with color a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setShadow(offset:blur:color:)
func CGContextSetShadowWithColor(c CGContextRef, offset corefoundation.CGSize, blur float64, color CGColorRef) {
	if callErr := tryCGContextSetShadowWithColor(c, offset, blur, color); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetShouldAntialias func(c CGContextRef, shouldAntialias bool)
var _cGContextSetShouldAntialiasErr error

func tryCGContextSetShouldAntialias(c CGContextRef, shouldAntialias bool) error {
	if _cGContextSetShouldAntialias == nil {
		return symbolCallError("CGContextSetShouldAntialias", "10.0", _cGContextSetShouldAntialiasErr)
	}
	_cGContextSetShouldAntialias(c, shouldAntialias)
	return nil
}

// CGContextSetShouldAntialias sets antialiasing on or off for a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setShouldAntialias(_:)
func CGContextSetShouldAntialias(c CGContextRef, shouldAntialias bool) {
	if callErr := tryCGContextSetShouldAntialias(c, shouldAntialias); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetShouldSmoothFonts func(c CGContextRef, shouldSmoothFonts bool)
var _cGContextSetShouldSmoothFontsErr error

func tryCGContextSetShouldSmoothFonts(c CGContextRef, shouldSmoothFonts bool) error {
	if _cGContextSetShouldSmoothFonts == nil {
		return symbolCallError("CGContextSetShouldSmoothFonts", "10.2", _cGContextSetShouldSmoothFontsErr)
	}
	_cGContextSetShouldSmoothFonts(c, shouldSmoothFonts)
	return nil
}

// CGContextSetShouldSmoothFonts enables or disables font smoothing in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setShouldSmoothFonts(_:)
func CGContextSetShouldSmoothFonts(c CGContextRef, shouldSmoothFonts bool) {
	if callErr := tryCGContextSetShouldSmoothFonts(c, shouldSmoothFonts); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetShouldSubpixelPositionFonts func(c CGContextRef, shouldSubpixelPositionFonts bool)
var _cGContextSetShouldSubpixelPositionFontsErr error

func tryCGContextSetShouldSubpixelPositionFonts(c CGContextRef, shouldSubpixelPositionFonts bool) error {
	if _cGContextSetShouldSubpixelPositionFonts == nil {
		return symbolCallError("CGContextSetShouldSubpixelPositionFonts", "10.5", _cGContextSetShouldSubpixelPositionFontsErr)
	}
	_cGContextSetShouldSubpixelPositionFonts(c, shouldSubpixelPositionFonts)
	return nil
}

// CGContextSetShouldSubpixelPositionFonts enables or disables subpixel positioning in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setShouldSubpixelPositionFonts(_:)
func CGContextSetShouldSubpixelPositionFonts(c CGContextRef, shouldSubpixelPositionFonts bool) {
	if callErr := tryCGContextSetShouldSubpixelPositionFonts(c, shouldSubpixelPositionFonts); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetShouldSubpixelQuantizeFonts func(c CGContextRef, shouldSubpixelQuantizeFonts bool)
var _cGContextSetShouldSubpixelQuantizeFontsErr error

func tryCGContextSetShouldSubpixelQuantizeFonts(c CGContextRef, shouldSubpixelQuantizeFonts bool) error {
	if _cGContextSetShouldSubpixelQuantizeFonts == nil {
		return symbolCallError("CGContextSetShouldSubpixelQuantizeFonts", "10.5", _cGContextSetShouldSubpixelQuantizeFontsErr)
	}
	_cGContextSetShouldSubpixelQuantizeFonts(c, shouldSubpixelQuantizeFonts)
	return nil
}

// CGContextSetShouldSubpixelQuantizeFonts enables or disables subpixel quantization in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setShouldSubpixelQuantizeFonts(_:)
func CGContextSetShouldSubpixelQuantizeFonts(c CGContextRef, shouldSubpixelQuantizeFonts bool) {
	if callErr := tryCGContextSetShouldSubpixelQuantizeFonts(c, shouldSubpixelQuantizeFonts); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetStrokeColor func(c CGContextRef, components *float64)
var _cGContextSetStrokeColorErr error

func tryCGContextSetStrokeColor(c CGContextRef, components *float64) error {
	if _cGContextSetStrokeColor == nil {
		return symbolCallError("CGContextSetStrokeColor", "10.0", _cGContextSetStrokeColorErr)
	}
	_cGContextSetStrokeColor(c, components)
	return nil
}

// CGContextSetStrokeColor sets the current stroke color.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokeColor(_:)-4pd8p
func CGContextSetStrokeColor(c CGContextRef, components *float64) {
	if callErr := tryCGContextSetStrokeColor(c, components); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetStrokeColorSpace func(c CGContextRef, space CGColorSpaceRef)
var _cGContextSetStrokeColorSpaceErr error

func tryCGContextSetStrokeColorSpace(c CGContextRef, space CGColorSpaceRef) error {
	if _cGContextSetStrokeColorSpace == nil {
		return symbolCallError("CGContextSetStrokeColorSpace", "10.0", _cGContextSetStrokeColorSpaceErr)
	}
	_cGContextSetStrokeColorSpace(c, space)
	return nil
}

// CGContextSetStrokeColorSpace sets the stroke color space in a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokeColorSpace(_:)
func CGContextSetStrokeColorSpace(c CGContextRef, space CGColorSpaceRef) {
	if callErr := tryCGContextSetStrokeColorSpace(c, space); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetStrokeColorWithColor func(c CGContextRef, color CGColorRef)
var _cGContextSetStrokeColorWithColorErr error

func tryCGContextSetStrokeColorWithColor(c CGContextRef, color CGColorRef) error {
	if _cGContextSetStrokeColorWithColor == nil {
		return symbolCallError("CGContextSetStrokeColorWithColor", "10.3", _cGContextSetStrokeColorWithColorErr)
	}
	_cGContextSetStrokeColorWithColor(c, color)
	return nil
}

// CGContextSetStrokeColorWithColor sets the current stroke color in a context, using a CGColor.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokeColor(_:)-1sskg
func CGContextSetStrokeColorWithColor(c CGContextRef, color CGColorRef) {
	if callErr := tryCGContextSetStrokeColorWithColor(c, color); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetStrokePattern func(c CGContextRef, pattern CGPatternRef, components *float64)
var _cGContextSetStrokePatternErr error

func tryCGContextSetStrokePattern(c CGContextRef, pattern CGPatternRef, components *float64) error {
	if _cGContextSetStrokePattern == nil {
		return symbolCallError("CGContextSetStrokePattern", "10.0", _cGContextSetStrokePatternErr)
	}
	_cGContextSetStrokePattern(c, pattern, components)
	return nil
}

// CGContextSetStrokePattern sets the stroke pattern in the specified graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setStrokePattern(_:colorComponents:)
func CGContextSetStrokePattern(c CGContextRef, pattern CGPatternRef, components *float64) {
	if callErr := tryCGContextSetStrokePattern(c, pattern, components); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetTextDrawingMode func(c CGContextRef, mode CGTextDrawingMode)
var _cGContextSetTextDrawingModeErr error

func tryCGContextSetTextDrawingMode(c CGContextRef, mode CGTextDrawingMode) error {
	if _cGContextSetTextDrawingMode == nil {
		return symbolCallError("CGContextSetTextDrawingMode", "10.0", _cGContextSetTextDrawingModeErr)
	}
	_cGContextSetTextDrawingMode(c, mode)
	return nil
}

// CGContextSetTextDrawingMode sets the current text drawing mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setTextDrawingMode(_:)
func CGContextSetTextDrawingMode(c CGContextRef, mode CGTextDrawingMode) {
	if callErr := tryCGContextSetTextDrawingMode(c, mode); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetTextMatrix func(c CGContextRef, t corefoundation.CGAffineTransform)
var _cGContextSetTextMatrixErr error

func tryCGContextSetTextMatrix(c CGContextRef, t corefoundation.CGAffineTransform) error {
	if _cGContextSetTextMatrix == nil {
		return symbolCallError("CGContextSetTextMatrix", "10.0", _cGContextSetTextMatrixErr)
	}
	_cGContextSetTextMatrix(c, t)
	return nil
}

// CGContextSetTextMatrix sets the current text matrix.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextSetTextMatrix
func CGContextSetTextMatrix(c CGContextRef, t corefoundation.CGAffineTransform) {
	if callErr := tryCGContextSetTextMatrix(c, t); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSetTextPosition func(c CGContextRef, x float64, y float64)
var _cGContextSetTextPositionErr error

func tryCGContextSetTextPosition(c CGContextRef, x float64, y float64) error {
	if _cGContextSetTextPosition == nil {
		return symbolCallError("CGContextSetTextPosition", "10.0", _cGContextSetTextPositionErr)
	}
	_cGContextSetTextPosition(c, x, y)
	return nil
}

// CGContextSetTextPosition sets the location at which text is drawn.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextSetTextPosition
func CGContextSetTextPosition(c CGContextRef, x float64, y float64) {
	if callErr := tryCGContextSetTextPosition(c, x, y); callErr != nil {
		panic(callErr)
	}
}

var _cGContextShowGlyphsAtPositions func(c CGContextRef, glyphs unsafe.Pointer, Lpositions *corefoundation.CGPoint, count uintptr)
var _cGContextShowGlyphsAtPositionsErr error

func tryCGContextShowGlyphsAtPositions(c CGContextRef, glyphs unsafe.Pointer, Lpositions *corefoundation.CGPoint, count uintptr) error {
	if _cGContextShowGlyphsAtPositions == nil {
		return symbolCallError("CGContextShowGlyphsAtPositions", "10.5", _cGContextShowGlyphsAtPositionsErr)
	}
	_cGContextShowGlyphsAtPositions(c, glyphs, Lpositions, count)
	return nil
}

// CGContextShowGlyphsAtPositions draws glyphs at the provided position.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextShowGlyphsAtPositions
func CGContextShowGlyphsAtPositions(c CGContextRef, glyphs unsafe.Pointer, Lpositions *corefoundation.CGPoint, count uintptr) {
	if callErr := tryCGContextShowGlyphsAtPositions(c, glyphs, Lpositions, count); callErr != nil {
		panic(callErr)
	}
}

var _cGContextStrokeEllipseInRect func(c CGContextRef, rect corefoundation.CGRect)
var _cGContextStrokeEllipseInRectErr error

func tryCGContextStrokeEllipseInRect(c CGContextRef, rect corefoundation.CGRect) error {
	if _cGContextStrokeEllipseInRect == nil {
		return symbolCallError("CGContextStrokeEllipseInRect", "10.4", _cGContextStrokeEllipseInRectErr)
	}
	_cGContextStrokeEllipseInRect(c, rect)
	return nil
}

// CGContextStrokeEllipseInRect strokes an ellipse that fits inside the specified rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/strokeEllipse(in:)
func CGContextStrokeEllipseInRect(c CGContextRef, rect corefoundation.CGRect) {
	if callErr := tryCGContextStrokeEllipseInRect(c, rect); callErr != nil {
		panic(callErr)
	}
}

var _cGContextStrokeLineSegments func(c CGContextRef, points *corefoundation.CGPoint, count uintptr)
var _cGContextStrokeLineSegmentsErr error

func tryCGContextStrokeLineSegments(c CGContextRef, points *corefoundation.CGPoint, count uintptr) error {
	if _cGContextStrokeLineSegments == nil {
		return symbolCallError("CGContextStrokeLineSegments", "10.4", _cGContextStrokeLineSegmentsErr)
	}
	_cGContextStrokeLineSegments(c, points, count)
	return nil
}

// CGContextStrokeLineSegments strokes a sequence of line segments.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContextStrokeLineSegments
func CGContextStrokeLineSegments(c CGContextRef, points *corefoundation.CGPoint, count uintptr) {
	if callErr := tryCGContextStrokeLineSegments(c, points, count); callErr != nil {
		panic(callErr)
	}
}

var _cGContextStrokePath func(c CGContextRef)
var _cGContextStrokePathErr error

func tryCGContextStrokePath(c CGContextRef) error {
	if _cGContextStrokePath == nil {
		return symbolCallError("CGContextStrokePath", "10.0", _cGContextStrokePathErr)
	}
	_cGContextStrokePath(c)
	return nil
}

// CGContextStrokePath paints a line along the current path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/strokePath()
func CGContextStrokePath(c CGContextRef) {
	if callErr := tryCGContextStrokePath(c); callErr != nil {
		panic(callErr)
	}
}

var _cGContextStrokeRect func(c CGContextRef, rect corefoundation.CGRect)
var _cGContextStrokeRectErr error

func tryCGContextStrokeRect(c CGContextRef, rect corefoundation.CGRect) error {
	if _cGContextStrokeRect == nil {
		return symbolCallError("CGContextStrokeRect", "10.0", _cGContextStrokeRectErr)
	}
	_cGContextStrokeRect(c, rect)
	return nil
}

// CGContextStrokeRect paints a rectangular path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/stroke(_:)
func CGContextStrokeRect(c CGContextRef, rect corefoundation.CGRect) {
	if callErr := tryCGContextStrokeRect(c, rect); callErr != nil {
		panic(callErr)
	}
}

var _cGContextStrokeRectWithWidth func(c CGContextRef, rect corefoundation.CGRect, width float64)
var _cGContextStrokeRectWithWidthErr error

func tryCGContextStrokeRectWithWidth(c CGContextRef, rect corefoundation.CGRect, width float64) error {
	if _cGContextStrokeRectWithWidth == nil {
		return symbolCallError("CGContextStrokeRectWithWidth", "10.0", _cGContextStrokeRectWithWidthErr)
	}
	_cGContextStrokeRectWithWidth(c, rect, width)
	return nil
}

// CGContextStrokeRectWithWidth paints a rectangular path, using the specified line width.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/stroke(_:width:)
func CGContextStrokeRectWithWidth(c CGContextRef, rect corefoundation.CGRect, width float64) {
	if callErr := tryCGContextStrokeRectWithWidth(c, rect, width); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSynchronize func(c CGContextRef)
var _cGContextSynchronizeErr error

func tryCGContextSynchronize(c CGContextRef) error {
	if _cGContextSynchronize == nil {
		return symbolCallError("CGContextSynchronize", "10.0", _cGContextSynchronizeErr)
	}
	_cGContextSynchronize(c)
	return nil
}

// CGContextSynchronize marks a window context for update.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/synchronize()
func CGContextSynchronize(c CGContextRef) {
	if callErr := tryCGContextSynchronize(c); callErr != nil {
		panic(callErr)
	}
}

var _cGContextSynchronizeAttributes func(c CGContextRef)
var _cGContextSynchronizeAttributesErr error

func tryCGContextSynchronizeAttributes(c CGContextRef) error {
	if _cGContextSynchronizeAttributes == nil {
		return symbolCallError("CGContextSynchronizeAttributes", "26.0", _cGContextSynchronizeAttributesErr)
	}
	_cGContextSynchronizeAttributes(c)
	return nil
}

// CGContextSynchronizeAttributes.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/synchronizeAttributes()
func CGContextSynchronizeAttributes(c CGContextRef) {
	if callErr := tryCGContextSynchronizeAttributes(c); callErr != nil {
		panic(callErr)
	}
}

var _cGContextTranslateCTM func(c CGContextRef, tx float64, ty float64)
var _cGContextTranslateCTMErr error

func tryCGContextTranslateCTM(c CGContextRef, tx float64, ty float64) error {
	if _cGContextTranslateCTM == nil {
		return symbolCallError("CGContextTranslateCTM", "10.0", _cGContextTranslateCTMErr)
	}
	_cGContextTranslateCTM(c, tx, ty)
	return nil
}

// CGContextTranslateCTM changes the origin of the user coordinate system in a context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/translateBy(x:y:)
func CGContextTranslateCTM(c CGContextRef, tx float64, ty float64) {
	if callErr := tryCGContextTranslateCTM(c, tx, ty); callErr != nil {
		panic(callErr)
	}
}

var _cGConvertColorDataWithFormat func(width uintptr, height uintptr, dst_data unsafe.Pointer, dst_format CGColorDataFormat, src_data unsafe.Pointer, src_format CGColorDataFormat, options corefoundation.CFDictionaryRef) bool
var _cGConvertColorDataWithFormatErr error

func tryCGConvertColorDataWithFormat(width uintptr, height uintptr, dst_data unsafe.Pointer, dst_format CGColorDataFormat, src_data unsafe.Pointer, src_format CGColorDataFormat, options corefoundation.CFDictionaryRef) (bool, error) {
	if _cGConvertColorDataWithFormat == nil {
		return false, symbolCallError("CGConvertColorDataWithFormat", "", _cGConvertColorDataWithFormatErr)
	}
	return _cGConvertColorDataWithFormat(width, height, dst_data, dst_format, src_data, src_format, options), nil
}

// CGConvertColorDataWithFormat.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGConvertColorDataWithFormat(_:_:_:_:_:_:_:)
func CGConvertColorDataWithFormat(width uintptr, height uintptr, dst_data unsafe.Pointer, dst_format CGColorDataFormat, src_data unsafe.Pointer, src_format CGColorDataFormat, options corefoundation.CFDictionaryRef) bool {
	result, callErr := tryCGConvertColorDataWithFormat(width, height, dst_data, dst_format, src_data, src_format, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDataConsumerCreate func(info unsafe.Pointer, cbks *CGDataConsumerCallbacks) CGDataConsumerRef
var _cGDataConsumerCreateErr error

func tryCGDataConsumerCreate(info unsafe.Pointer, cbks *CGDataConsumerCallbacks) (CGDataConsumerRef, error) {
	if _cGDataConsumerCreate == nil {
		return 0, symbolCallError("CGDataConsumerCreate", "10.0", _cGDataConsumerCreateErr)
	}
	return _cGDataConsumerCreate(info, cbks), nil
}

// CGDataConsumerCreate creates a data consumer that uses callback functions to write data.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumer/init(info:cbks:)
func CGDataConsumerCreate(info unsafe.Pointer, cbks *CGDataConsumerCallbacks) CGDataConsumerRef {
	result, callErr := tryCGDataConsumerCreate(info, cbks)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDataConsumerCreateWithCFData func(data corefoundation.CFMutableDataRef) CGDataConsumerRef
var _cGDataConsumerCreateWithCFDataErr error

func tryCGDataConsumerCreateWithCFData(data corefoundation.CFMutableDataRef) (CGDataConsumerRef, error) {
	if _cGDataConsumerCreateWithCFData == nil {
		return 0, symbolCallError("CGDataConsumerCreateWithCFData", "10.4", _cGDataConsumerCreateWithCFDataErr)
	}
	return _cGDataConsumerCreateWithCFData(data), nil
}

// CGDataConsumerCreateWithCFData creates a data consumer that writes to a CFData object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumer/init(data:)
func CGDataConsumerCreateWithCFData(data corefoundation.CFMutableDataRef) CGDataConsumerRef {
	result, callErr := tryCGDataConsumerCreateWithCFData(data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDataConsumerCreateWithURL func(url corefoundation.CFURLRef) CGDataConsumerRef
var _cGDataConsumerCreateWithURLErr error

func tryCGDataConsumerCreateWithURL(url corefoundation.CFURLRef) (CGDataConsumerRef, error) {
	if _cGDataConsumerCreateWithURL == nil {
		return 0, symbolCallError("CGDataConsumerCreateWithURL", "10.0", _cGDataConsumerCreateWithURLErr)
	}
	return _cGDataConsumerCreateWithURL(url), nil
}

// CGDataConsumerCreateWithURL creates a data consumer that writes data to a location specified by a URL.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumer/init(url:)
func CGDataConsumerCreateWithURL(url corefoundation.CFURLRef) CGDataConsumerRef {
	result, callErr := tryCGDataConsumerCreateWithURL(url)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDataConsumerGetTypeID func() uint
var _cGDataConsumerGetTypeIDErr error

func tryCGDataConsumerGetTypeID() (uint, error) {
	if _cGDataConsumerGetTypeID == nil {
		return 0, symbolCallError("CGDataConsumerGetTypeID", "10.2", _cGDataConsumerGetTypeIDErr)
	}
	return _cGDataConsumerGetTypeID(), nil
}

// CGDataConsumerGetTypeID returns the Core Foundation type identifier for Core Graphics data consumers.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumer/typeID
func CGDataConsumerGetTypeID() uint {
	result, callErr := tryCGDataConsumerGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDataConsumerRelease func(consumer CGDataConsumerRef)
var _cGDataConsumerReleaseErr error

func tryCGDataConsumerRelease(consumer CGDataConsumerRef) error {
	if _cGDataConsumerRelease == nil {
		return symbolCallError("CGDataConsumerRelease", "10.0", _cGDataConsumerReleaseErr)
	}
	_cGDataConsumerRelease(consumer)
	return nil
}

// CGDataConsumerRelease decrements the retain count of a data consumer.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumerRelease
func CGDataConsumerRelease(consumer CGDataConsumerRef) {
	if callErr := tryCGDataConsumerRelease(consumer); callErr != nil {
		panic(callErr)
	}
}

var _cGDataConsumerRetain func(consumer CGDataConsumerRef) CGDataConsumerRef
var _cGDataConsumerRetainErr error

func tryCGDataConsumerRetain(consumer CGDataConsumerRef) (CGDataConsumerRef, error) {
	if _cGDataConsumerRetain == nil {
		return 0, symbolCallError("CGDataConsumerRetain", "10.0", _cGDataConsumerRetainErr)
	}
	return _cGDataConsumerRetain(consumer), nil
}

// CGDataConsumerRetain increments the retain count of a data consumer.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataConsumerRetain
func CGDataConsumerRetain(consumer CGDataConsumerRef) CGDataConsumerRef {
	result, callErr := tryCGDataConsumerRetain(consumer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDataProviderCopyData func(provider CGDataProviderRef) corefoundation.CFDataRef
var _cGDataProviderCopyDataErr error

func tryCGDataProviderCopyData(provider CGDataProviderRef) (corefoundation.CFDataRef, error) {
	if _cGDataProviderCopyData == nil {
		return 0, symbolCallError("CGDataProviderCopyData", "10.3", _cGDataProviderCopyDataErr)
	}
	return _cGDataProviderCopyData(provider), nil
}

// CGDataProviderCopyData returns a copy of the provider’s data.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/data
func CGDataProviderCopyData(provider CGDataProviderRef) corefoundation.CFDataRef {
	result, callErr := tryCGDataProviderCopyData(provider)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDataProviderCreateDirect func(info unsafe.Pointer, size int64, callbacks *CGDataProviderDirectCallbacks) CGDataProviderRef
var _cGDataProviderCreateDirectErr error

func tryCGDataProviderCreateDirect(info unsafe.Pointer, size int64, callbacks *CGDataProviderDirectCallbacks) (CGDataProviderRef, error) {
	if _cGDataProviderCreateDirect == nil {
		return 0, symbolCallError("CGDataProviderCreateDirect", "10.5", _cGDataProviderCreateDirectErr)
	}
	return _cGDataProviderCreateDirect(info, size, callbacks), nil
}

// CGDataProviderCreateDirect creates a direct-access data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/init(directInfo:size:callbacks:)
func CGDataProviderCreateDirect(info unsafe.Pointer, size int64, callbacks *CGDataProviderDirectCallbacks) CGDataProviderRef {
	result, callErr := tryCGDataProviderCreateDirect(info, size, callbacks)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDataProviderCreateSequential func(info unsafe.Pointer, callbacks *CGDataProviderSequentialCallbacks) CGDataProviderRef
var _cGDataProviderCreateSequentialErr error

func tryCGDataProviderCreateSequential(info unsafe.Pointer, callbacks *CGDataProviderSequentialCallbacks) (CGDataProviderRef, error) {
	if _cGDataProviderCreateSequential == nil {
		return 0, symbolCallError("CGDataProviderCreateSequential", "10.5", _cGDataProviderCreateSequentialErr)
	}
	return _cGDataProviderCreateSequential(info, callbacks), nil
}

// CGDataProviderCreateSequential creates a sequential-access data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/init(sequentialInfo:callbacks:)
func CGDataProviderCreateSequential(info unsafe.Pointer, callbacks *CGDataProviderSequentialCallbacks) CGDataProviderRef {
	result, callErr := tryCGDataProviderCreateSequential(info, callbacks)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDataProviderCreateWithCFData func(data corefoundation.CFDataRef) CGDataProviderRef
var _cGDataProviderCreateWithCFDataErr error

func tryCGDataProviderCreateWithCFData(data corefoundation.CFDataRef) (CGDataProviderRef, error) {
	if _cGDataProviderCreateWithCFData == nil {
		return 0, symbolCallError("CGDataProviderCreateWithCFData", "10.4", _cGDataProviderCreateWithCFDataErr)
	}
	return _cGDataProviderCreateWithCFData(data), nil
}

// CGDataProviderCreateWithCFData creates a data provider that reads from a CFData object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/init(data:)
func CGDataProviderCreateWithCFData(data corefoundation.CFDataRef) CGDataProviderRef {
	result, callErr := tryCGDataProviderCreateWithCFData(data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDataProviderCreateWithData func(info unsafe.Pointer, data unsafe.Pointer, size uintptr, releaseData CGDataProviderReleaseDataCallback) CGDataProviderRef
var _cGDataProviderCreateWithDataErr error

func tryCGDataProviderCreateWithData(info unsafe.Pointer, data unsafe.Pointer, size uintptr, releaseData CGDataProviderReleaseDataCallback) (CGDataProviderRef, error) {
	if _cGDataProviderCreateWithData == nil {
		return 0, symbolCallError("CGDataProviderCreateWithData", "10.0", _cGDataProviderCreateWithDataErr)
	}
	return _cGDataProviderCreateWithData(info, data, size, releaseData), nil
}

// CGDataProviderCreateWithData creates a direct-access data provider that uses data your program supplies.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/init(dataInfo:data:size:releaseData:)
func CGDataProviderCreateWithData(info unsafe.Pointer, data unsafe.Pointer, size uintptr, releaseData CGDataProviderReleaseDataCallback) CGDataProviderRef {
	result, callErr := tryCGDataProviderCreateWithData(info, data, size, releaseData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDataProviderCreateWithFilename func(filename string) CGDataProviderRef
var _cGDataProviderCreateWithFilenameErr error

func tryCGDataProviderCreateWithFilename(filename string) (CGDataProviderRef, error) {
	if _cGDataProviderCreateWithFilename == nil {
		return 0, symbolCallError("CGDataProviderCreateWithFilename", "10.0", _cGDataProviderCreateWithFilenameErr)
	}
	return _cGDataProviderCreateWithFilename(filename), nil
}

// CGDataProviderCreateWithFilename creates a direct-access data provider that uses a file to supply data.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/init(filename:)
func CGDataProviderCreateWithFilename(filename string) CGDataProviderRef {
	result, callErr := tryCGDataProviderCreateWithFilename(filename)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDataProviderCreateWithURL func(url corefoundation.CFURLRef) CGDataProviderRef
var _cGDataProviderCreateWithURLErr error

func tryCGDataProviderCreateWithURL(url corefoundation.CFURLRef) (CGDataProviderRef, error) {
	if _cGDataProviderCreateWithURL == nil {
		return 0, symbolCallError("CGDataProviderCreateWithURL", "10.0", _cGDataProviderCreateWithURLErr)
	}
	return _cGDataProviderCreateWithURL(url), nil
}

// CGDataProviderCreateWithURL creates a direct-access data provider that uses a URL to supply data.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/init(url:)
func CGDataProviderCreateWithURL(url corefoundation.CFURLRef) CGDataProviderRef {
	result, callErr := tryCGDataProviderCreateWithURL(url)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDataProviderGetInfo func(provider CGDataProviderRef) unsafe.Pointer
var _cGDataProviderGetInfoErr error

func tryCGDataProviderGetInfo(provider CGDataProviderRef) (unsafe.Pointer, error) {
	if _cGDataProviderGetInfo == nil {
		return nil, symbolCallError("CGDataProviderGetInfo", "10.13", _cGDataProviderGetInfoErr)
	}
	return _cGDataProviderGetInfo(provider), nil
}

// CGDataProviderGetInfo.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/info
func CGDataProviderGetInfo(provider CGDataProviderRef) unsafe.Pointer {
	result, callErr := tryCGDataProviderGetInfo(provider)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDataProviderGetTypeID func() uint
var _cGDataProviderGetTypeIDErr error

func tryCGDataProviderGetTypeID() (uint, error) {
	if _cGDataProviderGetTypeID == nil {
		return 0, symbolCallError("CGDataProviderGetTypeID", "10.2", _cGDataProviderGetTypeIDErr)
	}
	return _cGDataProviderGetTypeID(), nil
}

// CGDataProviderGetTypeID returns the Core Foundation type identifier for data providers.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProvider/typeID
func CGDataProviderGetTypeID() uint {
	result, callErr := tryCGDataProviderGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDataProviderRelease func(provider CGDataProviderRef)
var _cGDataProviderReleaseErr error

func tryCGDataProviderRelease(provider CGDataProviderRef) error {
	if _cGDataProviderRelease == nil {
		return symbolCallError("CGDataProviderRelease", "10.0", _cGDataProviderReleaseErr)
	}
	_cGDataProviderRelease(provider)
	return nil
}

// CGDataProviderRelease decrements the retain count of a data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProviderRelease
func CGDataProviderRelease(provider CGDataProviderRef) {
	if callErr := tryCGDataProviderRelease(provider); callErr != nil {
		panic(callErr)
	}
}

var _cGDataProviderRetain func(provider CGDataProviderRef) CGDataProviderRef
var _cGDataProviderRetainErr error

func tryCGDataProviderRetain(provider CGDataProviderRef) (CGDataProviderRef, error) {
	if _cGDataProviderRetain == nil {
		return 0, symbolCallError("CGDataProviderRetain", "10.0", _cGDataProviderRetainErr)
	}
	return _cGDataProviderRetain(provider), nil
}

// CGDataProviderRetain increments the retain count of a data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDataProviderRetain
func CGDataProviderRetain(provider CGDataProviderRef) CGDataProviderRef {
	result, callErr := tryCGDataProviderRetain(provider)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDirectDisplayCopyCurrentMetalDevice func(display uint32) unsafe.Pointer
var _cGDirectDisplayCopyCurrentMetalDeviceErr error

func tryCGDirectDisplayCopyCurrentMetalDevice(display uint32) (objectivec.IObject, error) {
	if _cGDirectDisplayCopyCurrentMetalDevice == nil {
		return nil, symbolCallError("CGDirectDisplayCopyCurrentMetalDevice", "10.11", _cGDirectDisplayCopyCurrentMetalDeviceErr)
	}
	rv := _cGDirectDisplayCopyCurrentMetalDevice(display)
	return objectivec.ObjectFromID(objc.IDFrom(rv)), nil
}

// CGDirectDisplayCopyCurrentMetalDevice returns the GPU device instance that’s currently driving a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDirectDisplayCopyCurrentMetalDevice(_:)
func CGDirectDisplayCopyCurrentMetalDevice(display uint32) objectivec.IObject {
	result, callErr := tryCGDirectDisplayCopyCurrentMetalDevice(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayBounds func(display uint32) corefoundation.CGRect
var _cGDisplayBoundsErr error

func tryCGDisplayBounds(display uint32) (corefoundation.CGRect, error) {
	if _cGDisplayBounds == nil {
		return corefoundation.CGRect{}, symbolCallError("CGDisplayBounds", "10.0", _cGDisplayBoundsErr)
	}
	return _cGDisplayBounds(display), nil
}

// CGDisplayBounds returns the bounds of a display in the global display coordinate space.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayBounds(_:)
func CGDisplayBounds(display uint32) corefoundation.CGRect {
	result, callErr := tryCGDisplayBounds(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayCapture func(display uint32) CGError
var _cGDisplayCaptureErr error

func tryCGDisplayCapture(display uint32) (CGError, error) {
	if _cGDisplayCapture == nil {
		return *new(CGError), symbolCallError("CGDisplayCapture", "10.0", _cGDisplayCaptureErr)
	}
	return _cGDisplayCapture(display), nil
}

// CGDisplayCapture obtains exclusive use of a display, preventing other applications and system services from using the display or changing its configuration.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCapture(_:)
func CGDisplayCapture(display uint32) CGError {
	result, callErr := tryCGDisplayCapture(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayCaptureWithOptions func(display uint32, options CGCaptureOptions) CGError
var _cGDisplayCaptureWithOptionsErr error

func tryCGDisplayCaptureWithOptions(display uint32, options CGCaptureOptions) (CGError, error) {
	if _cGDisplayCaptureWithOptions == nil {
		return *new(CGError), symbolCallError("CGDisplayCaptureWithOptions", "10.3", _cGDisplayCaptureWithOptionsErr)
	}
	return _cGDisplayCaptureWithOptions(display, options), nil
}

// CGDisplayCaptureWithOptions obtains exclusive use of a display for an application using the options you specify.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCaptureWithOptions(_:_:)
func CGDisplayCaptureWithOptions(display uint32, options CGCaptureOptions) CGError {
	result, callErr := tryCGDisplayCaptureWithOptions(display, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayCopyAllDisplayModes func(display uint32, options corefoundation.CFDictionaryRef) corefoundation.CFArrayRef
var _cGDisplayCopyAllDisplayModesErr error

func tryCGDisplayCopyAllDisplayModes(display uint32, options corefoundation.CFDictionaryRef) (corefoundation.CFArrayRef, error) {
	if _cGDisplayCopyAllDisplayModes == nil {
		return 0, symbolCallError("CGDisplayCopyAllDisplayModes", "10.6", _cGDisplayCopyAllDisplayModesErr)
	}
	return _cGDisplayCopyAllDisplayModes(display, options), nil
}

// CGDisplayCopyAllDisplayModes returns information about the currently available display modes.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCopyAllDisplayModes(_:_:)
func CGDisplayCopyAllDisplayModes(display uint32, options corefoundation.CFDictionaryRef) corefoundation.CFArrayRef {
	result, callErr := tryCGDisplayCopyAllDisplayModes(display, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayCopyColorSpace func(display uint32) CGColorSpaceRef
var _cGDisplayCopyColorSpaceErr error

func tryCGDisplayCopyColorSpace(display uint32) (CGColorSpaceRef, error) {
	if _cGDisplayCopyColorSpace == nil {
		return 0, symbolCallError("CGDisplayCopyColorSpace", "10.5", _cGDisplayCopyColorSpaceErr)
	}
	return _cGDisplayCopyColorSpace(display), nil
}

// CGDisplayCopyColorSpace returns the color space for a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCopyColorSpace(_:)
func CGDisplayCopyColorSpace(display uint32) CGColorSpaceRef {
	result, callErr := tryCGDisplayCopyColorSpace(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayCopyDisplayMode func(display uint32) CGDisplayModeRef
var _cGDisplayCopyDisplayModeErr error

func tryCGDisplayCopyDisplayMode(display uint32) (CGDisplayModeRef, error) {
	if _cGDisplayCopyDisplayMode == nil {
		return 0, symbolCallError("CGDisplayCopyDisplayMode", "10.6", _cGDisplayCopyDisplayModeErr)
	}
	return _cGDisplayCopyDisplayMode(display), nil
}

// CGDisplayCopyDisplayMode returns information about a display’s current configuration.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCopyDisplayMode(_:)
func CGDisplayCopyDisplayMode(display uint32) CGDisplayModeRef {
	result, callErr := tryCGDisplayCopyDisplayMode(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayCreateImage func(displayID uint32) CGImageRef
var _cGDisplayCreateImageErr error

func tryCGDisplayCreateImage(displayID uint32) (CGImageRef, error) {
	if _cGDisplayCreateImage == nil {
		return 0, symbolCallError("CGDisplayCreateImage", "", _cGDisplayCreateImageErr)
	}
	return _cGDisplayCreateImage(displayID), nil
}

// CGDisplayCreateImage returns an image containing the contents of the specified display.
//
// Deprecated: Please use ScreenCaptureKit instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCreateImage(_:)
func CGDisplayCreateImage(displayID uint32) CGImageRef {
	result, callErr := tryCGDisplayCreateImage(displayID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayCreateImageForRect func(display uint32, rect corefoundation.CGRect) CGImageRef
var _cGDisplayCreateImageForRectErr error

func tryCGDisplayCreateImageForRect(display uint32, rect corefoundation.CGRect) (CGImageRef, error) {
	if _cGDisplayCreateImageForRect == nil {
		return 0, symbolCallError("CGDisplayCreateImageForRect", "", _cGDisplayCreateImageForRectErr)
	}
	return _cGDisplayCreateImageForRect(display, rect), nil
}

// CGDisplayCreateImageForRect returns an image containing the contents of a portion of the specified display.
//
// Deprecated: Please use ScreenCaptureKit instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayCreateImage(_:rect:)
func CGDisplayCreateImageForRect(display uint32, rect corefoundation.CGRect) CGImageRef {
	result, callErr := tryCGDisplayCreateImageForRect(display, rect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayFade func(token CGDisplayFadeReservationToken, duration float32, startBlend CGDisplayBlendFraction, endBlend CGDisplayBlendFraction, redBlend float32, greenBlend float32, blueBlend float32, synchronous bool) CGError
var _cGDisplayFadeErr error

func tryCGDisplayFade(token CGDisplayFadeReservationToken, duration float32, startBlend CGDisplayBlendFraction, endBlend CGDisplayBlendFraction, redBlend float32, greenBlend float32, blueBlend float32, synchronous bool) (CGError, error) {
	if _cGDisplayFade == nil {
		return *new(CGError), symbolCallError("CGDisplayFade", "10.2", _cGDisplayFadeErr)
	}
	return _cGDisplayFade(token, duration, startBlend, endBlend, redBlend, greenBlend, blueBlend, synchronous), nil
}

// CGDisplayFade performs a single fade operation.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayFade(_:_:_:_:_:_:_:_:)
func CGDisplayFade(token CGDisplayFadeReservationToken, duration float32, startBlend CGDisplayBlendFraction, endBlend CGDisplayBlendFraction, redBlend float32, greenBlend float32, blueBlend float32, synchronous bool) CGError {
	result, callErr := tryCGDisplayFade(token, duration, startBlend, endBlend, redBlend, greenBlend, blueBlend, synchronous)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayGammaTableCapacity func(display uint32) uint32
var _cGDisplayGammaTableCapacityErr error

func tryCGDisplayGammaTableCapacity(display uint32) (uint32, error) {
	if _cGDisplayGammaTableCapacity == nil {
		return 0, symbolCallError("CGDisplayGammaTableCapacity", "10.3", _cGDisplayGammaTableCapacityErr)
	}
	return _cGDisplayGammaTableCapacity(display), nil
}

// CGDisplayGammaTableCapacity returns the capacity, or number of entries, in the gamma table for a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayGammaTableCapacity(_:)
func CGDisplayGammaTableCapacity(display uint32) uint32 {
	result, callErr := tryCGDisplayGammaTableCapacity(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayGetDrawingContext func(display uint32) CGContextRef
var _cGDisplayGetDrawingContextErr error

func tryCGDisplayGetDrawingContext(display uint32) (CGContextRef, error) {
	if _cGDisplayGetDrawingContext == nil {
		return 0, symbolCallError("CGDisplayGetDrawingContext", "10.3", _cGDisplayGetDrawingContextErr)
	}
	return _cGDisplayGetDrawingContext(display), nil
}

// CGDisplayGetDrawingContext returns a graphics context suitable for drawing to a captured display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayGetDrawingContext(_:)
func CGDisplayGetDrawingContext(display uint32) CGContextRef {
	result, callErr := tryCGDisplayGetDrawingContext(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayHideCursor func(display uint32) CGError
var _cGDisplayHideCursorErr error

func tryCGDisplayHideCursor(display uint32) (CGError, error) {
	if _cGDisplayHideCursor == nil {
		return *new(CGError), symbolCallError("CGDisplayHideCursor", "10.0", _cGDisplayHideCursorErr)
	}
	return _cGDisplayHideCursor(display), nil
}

// CGDisplayHideCursor hides the mouse cursor, and increments the hide cursor count.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayHideCursor(_:)
func CGDisplayHideCursor(display uint32) CGError {
	result, callErr := tryCGDisplayHideCursor(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayIDToOpenGLDisplayMask func(display uint32) CGOpenGLDisplayMask
var _cGDisplayIDToOpenGLDisplayMaskErr error

func tryCGDisplayIDToOpenGLDisplayMask(display uint32) (CGOpenGLDisplayMask, error) {
	if _cGDisplayIDToOpenGLDisplayMask == nil {
		return *new(CGOpenGLDisplayMask), symbolCallError("CGDisplayIDToOpenGLDisplayMask", "10.0", _cGDisplayIDToOpenGLDisplayMaskErr)
	}
	return _cGDisplayIDToOpenGLDisplayMask(display), nil
}

// CGDisplayIDToOpenGLDisplayMask maps a display ID to an OpenGL display mask.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIDToOpenGLDisplayMask(_:)
func CGDisplayIDToOpenGLDisplayMask(display uint32) CGOpenGLDisplayMask {
	result, callErr := tryCGDisplayIDToOpenGLDisplayMask(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayIsActive func(display uint32) bool
var _cGDisplayIsActiveErr error

func tryCGDisplayIsActive(display uint32) (bool, error) {
	if _cGDisplayIsActive == nil {
		return false, symbolCallError("CGDisplayIsActive", "10.2", _cGDisplayIsActiveErr)
	}
	return _cGDisplayIsActive(display), nil
}

// CGDisplayIsActive returns a Boolean value indicating whether a display is active.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsActive(_:)
func CGDisplayIsActive(display uint32) bool {
	result, callErr := tryCGDisplayIsActive(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayIsAlwaysInMirrorSet func(display uint32) bool
var _cGDisplayIsAlwaysInMirrorSetErr error

func tryCGDisplayIsAlwaysInMirrorSet(display uint32) (bool, error) {
	if _cGDisplayIsAlwaysInMirrorSet == nil {
		return false, symbolCallError("CGDisplayIsAlwaysInMirrorSet", "10.2", _cGDisplayIsAlwaysInMirrorSetErr)
	}
	return _cGDisplayIsAlwaysInMirrorSet(display), nil
}

// CGDisplayIsAlwaysInMirrorSet returns a Boolean value indicating whether a display is always in a mirroring set.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsAlwaysInMirrorSet(_:)
func CGDisplayIsAlwaysInMirrorSet(display uint32) bool {
	result, callErr := tryCGDisplayIsAlwaysInMirrorSet(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayIsAsleep func(display uint32) bool
var _cGDisplayIsAsleepErr error

func tryCGDisplayIsAsleep(display uint32) (bool, error) {
	if _cGDisplayIsAsleep == nil {
		return false, symbolCallError("CGDisplayIsAsleep", "10.2", _cGDisplayIsAsleepErr)
	}
	return _cGDisplayIsAsleep(display), nil
}

// CGDisplayIsAsleep returns a Boolean value indicating whether a display is sleeping (and is therefore not drawable).
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsAsleep(_:)
func CGDisplayIsAsleep(display uint32) bool {
	result, callErr := tryCGDisplayIsAsleep(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayIsBuiltin func(display uint32) bool
var _cGDisplayIsBuiltinErr error

func tryCGDisplayIsBuiltin(display uint32) (bool, error) {
	if _cGDisplayIsBuiltin == nil {
		return false, symbolCallError("CGDisplayIsBuiltin", "10.2", _cGDisplayIsBuiltinErr)
	}
	return _cGDisplayIsBuiltin(display), nil
}

// CGDisplayIsBuiltin returns a Boolean value indicating whether a display is built-in, such as the internal display in portable systems.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsBuiltin(_:)
func CGDisplayIsBuiltin(display uint32) bool {
	result, callErr := tryCGDisplayIsBuiltin(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayIsInHWMirrorSet func(display uint32) bool
var _cGDisplayIsInHWMirrorSetErr error

func tryCGDisplayIsInHWMirrorSet(display uint32) (bool, error) {
	if _cGDisplayIsInHWMirrorSet == nil {
		return false, symbolCallError("CGDisplayIsInHWMirrorSet", "10.2", _cGDisplayIsInHWMirrorSetErr)
	}
	return _cGDisplayIsInHWMirrorSet(display), nil
}

// CGDisplayIsInHWMirrorSet returns a Boolean value indicating whether a display is in a hardware mirroring set.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsInHWMirrorSet(_:)
func CGDisplayIsInHWMirrorSet(display uint32) bool {
	result, callErr := tryCGDisplayIsInHWMirrorSet(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayIsInMirrorSet func(display uint32) bool
var _cGDisplayIsInMirrorSetErr error

func tryCGDisplayIsInMirrorSet(display uint32) (bool, error) {
	if _cGDisplayIsInMirrorSet == nil {
		return false, symbolCallError("CGDisplayIsInMirrorSet", "10.2", _cGDisplayIsInMirrorSetErr)
	}
	return _cGDisplayIsInMirrorSet(display), nil
}

// CGDisplayIsInMirrorSet returns a Boolean value indicating whether a display is in a mirroring set.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsInMirrorSet(_:)
func CGDisplayIsInMirrorSet(display uint32) bool {
	result, callErr := tryCGDisplayIsInMirrorSet(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayIsMain func(display uint32) bool
var _cGDisplayIsMainErr error

func tryCGDisplayIsMain(display uint32) (bool, error) {
	if _cGDisplayIsMain == nil {
		return false, symbolCallError("CGDisplayIsMain", "10.2", _cGDisplayIsMainErr)
	}
	return _cGDisplayIsMain(display), nil
}

// CGDisplayIsMain returns a Boolean value indicating whether a display is the main display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsMain(_:)
func CGDisplayIsMain(display uint32) bool {
	result, callErr := tryCGDisplayIsMain(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayIsOnline func(display uint32) bool
var _cGDisplayIsOnlineErr error

func tryCGDisplayIsOnline(display uint32) (bool, error) {
	if _cGDisplayIsOnline == nil {
		return false, symbolCallError("CGDisplayIsOnline", "10.2", _cGDisplayIsOnlineErr)
	}
	return _cGDisplayIsOnline(display), nil
}

// CGDisplayIsOnline returns a Boolean value indicating whether a display is connected or online.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsOnline(_:)
func CGDisplayIsOnline(display uint32) bool {
	result, callErr := tryCGDisplayIsOnline(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayIsStereo func(display uint32) bool
var _cGDisplayIsStereoErr error

func tryCGDisplayIsStereo(display uint32) (bool, error) {
	if _cGDisplayIsStereo == nil {
		return false, symbolCallError("CGDisplayIsStereo", "10.4", _cGDisplayIsStereoErr)
	}
	return _cGDisplayIsStereo(display), nil
}

// CGDisplayIsStereo returns a Boolean value indicating whether a display is running in a stereo graphics mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayIsStereo(_:)
func CGDisplayIsStereo(display uint32) bool {
	result, callErr := tryCGDisplayIsStereo(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayMirrorsDisplay func(display uint32) uint32
var _cGDisplayMirrorsDisplayErr error

func tryCGDisplayMirrorsDisplay(display uint32) (uint32, error) {
	if _cGDisplayMirrorsDisplay == nil {
		return 0, symbolCallError("CGDisplayMirrorsDisplay", "10.2", _cGDisplayMirrorsDisplayErr)
	}
	return _cGDisplayMirrorsDisplay(display), nil
}

// CGDisplayMirrorsDisplay for a secondary display in a mirroring set, returns the primary display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMirrorsDisplay(_:)
func CGDisplayMirrorsDisplay(display uint32) uint32 {
	result, callErr := tryCGDisplayMirrorsDisplay(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayModeGetHeight func(mode CGDisplayModeRef) uintptr
var _cGDisplayModeGetHeightErr error

func tryCGDisplayModeGetHeight(mode CGDisplayModeRef) (uintptr, error) {
	if _cGDisplayModeGetHeight == nil {
		return 0, symbolCallError("CGDisplayModeGetHeight", "10.6", _cGDisplayModeGetHeightErr)
	}
	return _cGDisplayModeGetHeight(mode), nil
}

// CGDisplayModeGetHeight returns the height of the specified display mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/height
func CGDisplayModeGetHeight(mode CGDisplayModeRef) uintptr {
	result, callErr := tryCGDisplayModeGetHeight(mode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayModeGetIODisplayModeID func(mode CGDisplayModeRef) int32
var _cGDisplayModeGetIODisplayModeIDErr error

func tryCGDisplayModeGetIODisplayModeID(mode CGDisplayModeRef) (int32, error) {
	if _cGDisplayModeGetIODisplayModeID == nil {
		return 0, symbolCallError("CGDisplayModeGetIODisplayModeID", "10.6", _cGDisplayModeGetIODisplayModeIDErr)
	}
	return _cGDisplayModeGetIODisplayModeID(mode), nil
}

// CGDisplayModeGetIODisplayModeID returns the I/O Kit display mode ID of the specified display mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/ioDisplayModeID
func CGDisplayModeGetIODisplayModeID(mode CGDisplayModeRef) int32 {
	result, callErr := tryCGDisplayModeGetIODisplayModeID(mode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayModeGetIOFlags func(mode CGDisplayModeRef) uint32
var _cGDisplayModeGetIOFlagsErr error

func tryCGDisplayModeGetIOFlags(mode CGDisplayModeRef) (uint32, error) {
	if _cGDisplayModeGetIOFlags == nil {
		return 0, symbolCallError("CGDisplayModeGetIOFlags", "10.6", _cGDisplayModeGetIOFlagsErr)
	}
	return _cGDisplayModeGetIOFlags(mode), nil
}

// CGDisplayModeGetIOFlags returns the I/O Kit flags of the specified display mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/ioFlags
func CGDisplayModeGetIOFlags(mode CGDisplayModeRef) uint32 {
	result, callErr := tryCGDisplayModeGetIOFlags(mode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayModeGetPixelHeight func(mode CGDisplayModeRef) uintptr
var _cGDisplayModeGetPixelHeightErr error

func tryCGDisplayModeGetPixelHeight(mode CGDisplayModeRef) (uintptr, error) {
	if _cGDisplayModeGetPixelHeight == nil {
		return 0, symbolCallError("CGDisplayModeGetPixelHeight", "10.8", _cGDisplayModeGetPixelHeightErr)
	}
	return _cGDisplayModeGetPixelHeight(mode), nil
}

// CGDisplayModeGetPixelHeight.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/pixelHeight
func CGDisplayModeGetPixelHeight(mode CGDisplayModeRef) uintptr {
	result, callErr := tryCGDisplayModeGetPixelHeight(mode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayModeGetPixelWidth func(mode CGDisplayModeRef) uintptr
var _cGDisplayModeGetPixelWidthErr error

func tryCGDisplayModeGetPixelWidth(mode CGDisplayModeRef) (uintptr, error) {
	if _cGDisplayModeGetPixelWidth == nil {
		return 0, symbolCallError("CGDisplayModeGetPixelWidth", "10.8", _cGDisplayModeGetPixelWidthErr)
	}
	return _cGDisplayModeGetPixelWidth(mode), nil
}

// CGDisplayModeGetPixelWidth.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/pixelWidth
func CGDisplayModeGetPixelWidth(mode CGDisplayModeRef) uintptr {
	result, callErr := tryCGDisplayModeGetPixelWidth(mode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayModeGetRefreshRate func(mode CGDisplayModeRef) float64
var _cGDisplayModeGetRefreshRateErr error

func tryCGDisplayModeGetRefreshRate(mode CGDisplayModeRef) (float64, error) {
	if _cGDisplayModeGetRefreshRate == nil {
		return 0.0, symbolCallError("CGDisplayModeGetRefreshRate", "10.6", _cGDisplayModeGetRefreshRateErr)
	}
	return _cGDisplayModeGetRefreshRate(mode), nil
}

// CGDisplayModeGetRefreshRate returns the refresh rate of the specified display mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/refreshRate
func CGDisplayModeGetRefreshRate(mode CGDisplayModeRef) float64 {
	result, callErr := tryCGDisplayModeGetRefreshRate(mode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayModeGetTypeID func() uint
var _cGDisplayModeGetTypeIDErr error

func tryCGDisplayModeGetTypeID() (uint, error) {
	if _cGDisplayModeGetTypeID == nil {
		return 0, symbolCallError("CGDisplayModeGetTypeID", "10.6", _cGDisplayModeGetTypeIDErr)
	}
	return _cGDisplayModeGetTypeID(), nil
}

// CGDisplayModeGetTypeID returns the type identifier of Quartz display modes.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/typeID
func CGDisplayModeGetTypeID() uint {
	result, callErr := tryCGDisplayModeGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayModeGetWidth func(mode CGDisplayModeRef) uintptr
var _cGDisplayModeGetWidthErr error

func tryCGDisplayModeGetWidth(mode CGDisplayModeRef) (uintptr, error) {
	if _cGDisplayModeGetWidth == nil {
		return 0, symbolCallError("CGDisplayModeGetWidth", "10.6", _cGDisplayModeGetWidthErr)
	}
	return _cGDisplayModeGetWidth(mode), nil
}

// CGDisplayModeGetWidth returns the width of the specified display mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/width
func CGDisplayModeGetWidth(mode CGDisplayModeRef) uintptr {
	result, callErr := tryCGDisplayModeGetWidth(mode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayModeIsUsableForDesktopGUI func(mode CGDisplayModeRef) bool
var _cGDisplayModeIsUsableForDesktopGUIErr error

func tryCGDisplayModeIsUsableForDesktopGUI(mode CGDisplayModeRef) (bool, error) {
	if _cGDisplayModeIsUsableForDesktopGUI == nil {
		return false, symbolCallError("CGDisplayModeIsUsableForDesktopGUI", "10.6", _cGDisplayModeIsUsableForDesktopGUIErr)
	}
	return _cGDisplayModeIsUsableForDesktopGUI(mode), nil
}

// CGDisplayModeIsUsableForDesktopGUI returns a Boolean value indicating whether the specified display mode is usable for a desktop graphical user interface.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMode/isUsableForDesktopGUI()
func CGDisplayModeIsUsableForDesktopGUI(mode CGDisplayModeRef) bool {
	result, callErr := tryCGDisplayModeIsUsableForDesktopGUI(mode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayModeRelease func(mode CGDisplayModeRef)
var _cGDisplayModeReleaseErr error

func tryCGDisplayModeRelease(mode CGDisplayModeRef) error {
	if _cGDisplayModeRelease == nil {
		return symbolCallError("CGDisplayModeRelease", "10.6", _cGDisplayModeReleaseErr)
	}
	_cGDisplayModeRelease(mode)
	return nil
}

// CGDisplayModeRelease releases a Core Graphics display mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayModeRelease
func CGDisplayModeRelease(mode CGDisplayModeRef) {
	if callErr := tryCGDisplayModeRelease(mode); callErr != nil {
		panic(callErr)
	}
}

var _cGDisplayModeRetain func(mode CGDisplayModeRef) CGDisplayModeRef
var _cGDisplayModeRetainErr error

func tryCGDisplayModeRetain(mode CGDisplayModeRef) (CGDisplayModeRef, error) {
	if _cGDisplayModeRetain == nil {
		return 0, symbolCallError("CGDisplayModeRetain", "10.6", _cGDisplayModeRetainErr)
	}
	return _cGDisplayModeRetain(mode), nil
}

// CGDisplayModeRetain retains a Core Graphics display mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayModeRetain
func CGDisplayModeRetain(mode CGDisplayModeRef) CGDisplayModeRef {
	result, callErr := tryCGDisplayModeRetain(mode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayModelNumber func(display uint32) uint32
var _cGDisplayModelNumberErr error

func tryCGDisplayModelNumber(display uint32) (uint32, error) {
	if _cGDisplayModelNumber == nil {
		return 0, symbolCallError("CGDisplayModelNumber", "10.2", _cGDisplayModelNumberErr)
	}
	return _cGDisplayModelNumber(display), nil
}

// CGDisplayModelNumber returns the model number of a display monitor.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayModelNumber(_:)
func CGDisplayModelNumber(display uint32) uint32 {
	result, callErr := tryCGDisplayModelNumber(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayMoveCursorToPoint func(display uint32, point corefoundation.CGPoint) CGError
var _cGDisplayMoveCursorToPointErr error

func tryCGDisplayMoveCursorToPoint(display uint32, point corefoundation.CGPoint) (CGError, error) {
	if _cGDisplayMoveCursorToPoint == nil {
		return *new(CGError), symbolCallError("CGDisplayMoveCursorToPoint", "10.0", _cGDisplayMoveCursorToPointErr)
	}
	return _cGDisplayMoveCursorToPoint(display, point), nil
}

// CGDisplayMoveCursorToPoint moves the mouse cursor to a specified point relative to the upper-left corner of the display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayMoveCursorToPoint(_:_:)
func CGDisplayMoveCursorToPoint(display uint32, point corefoundation.CGPoint) CGError {
	result, callErr := tryCGDisplayMoveCursorToPoint(display, point)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayPixelsHigh func(display uint32) uintptr
var _cGDisplayPixelsHighErr error

func tryCGDisplayPixelsHigh(display uint32) (uintptr, error) {
	if _cGDisplayPixelsHigh == nil {
		return 0, symbolCallError("CGDisplayPixelsHigh", "10.0", _cGDisplayPixelsHighErr)
	}
	return _cGDisplayPixelsHigh(display), nil
}

// CGDisplayPixelsHigh returns the display height in pixel units.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayPixelsHigh(_:)
func CGDisplayPixelsHigh(display uint32) uintptr {
	result, callErr := tryCGDisplayPixelsHigh(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayPixelsWide func(display uint32) uintptr
var _cGDisplayPixelsWideErr error

func tryCGDisplayPixelsWide(display uint32) (uintptr, error) {
	if _cGDisplayPixelsWide == nil {
		return 0, symbolCallError("CGDisplayPixelsWide", "10.0", _cGDisplayPixelsWideErr)
	}
	return _cGDisplayPixelsWide(display), nil
}

// CGDisplayPixelsWide returns the display width in pixel units.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayPixelsWide(_:)
func CGDisplayPixelsWide(display uint32) uintptr {
	result, callErr := tryCGDisplayPixelsWide(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayPrimaryDisplay func(display uint32) uint32
var _cGDisplayPrimaryDisplayErr error

func tryCGDisplayPrimaryDisplay(display uint32) (uint32, error) {
	if _cGDisplayPrimaryDisplay == nil {
		return 0, symbolCallError("CGDisplayPrimaryDisplay", "10.2", _cGDisplayPrimaryDisplayErr)
	}
	return _cGDisplayPrimaryDisplay(display), nil
}

// CGDisplayPrimaryDisplay returns the primary display in a hardware mirroring set.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayPrimaryDisplay(_:)
func CGDisplayPrimaryDisplay(display uint32) uint32 {
	result, callErr := tryCGDisplayPrimaryDisplay(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayRegisterReconfigurationCallback func(callback CGDisplayReconfigurationCallBack, userInfo unsafe.Pointer) CGError
var _cGDisplayRegisterReconfigurationCallbackErr error

func tryCGDisplayRegisterReconfigurationCallback(callback CGDisplayReconfigurationCallBack, userInfo unsafe.Pointer) (CGError, error) {
	if _cGDisplayRegisterReconfigurationCallback == nil {
		return *new(CGError), symbolCallError("CGDisplayRegisterReconfigurationCallback", "10.3", _cGDisplayRegisterReconfigurationCallbackErr)
	}
	return _cGDisplayRegisterReconfigurationCallback(callback, userInfo), nil
}

// CGDisplayRegisterReconfigurationCallback registers a callback function to be invoked whenever a local display is reconfigured.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayRegisterReconfigurationCallback(_:_:)
func CGDisplayRegisterReconfigurationCallback(callback CGDisplayReconfigurationCallBack, userInfo unsafe.Pointer) CGError {
	result, callErr := tryCGDisplayRegisterReconfigurationCallback(callback, userInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayRelease func(display uint32) CGError
var _cGDisplayReleaseErr error

func tryCGDisplayRelease(display uint32) (CGError, error) {
	if _cGDisplayRelease == nil {
		return *new(CGError), symbolCallError("CGDisplayRelease", "10.0", _cGDisplayReleaseErr)
	}
	return _cGDisplayRelease(display), nil
}

// CGDisplayRelease releases a captured display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayRelease(_:)
func CGDisplayRelease(display uint32) CGError {
	result, callErr := tryCGDisplayRelease(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayRemoveReconfigurationCallback func(callback CGDisplayReconfigurationCallBack, userInfo unsafe.Pointer) CGError
var _cGDisplayRemoveReconfigurationCallbackErr error

func tryCGDisplayRemoveReconfigurationCallback(callback CGDisplayReconfigurationCallBack, userInfo unsafe.Pointer) (CGError, error) {
	if _cGDisplayRemoveReconfigurationCallback == nil {
		return *new(CGError), symbolCallError("CGDisplayRemoveReconfigurationCallback", "10.3", _cGDisplayRemoveReconfigurationCallbackErr)
	}
	return _cGDisplayRemoveReconfigurationCallback(callback, userInfo), nil
}

// CGDisplayRemoveReconfigurationCallback removes the registration of a callback function that’s invoked whenever a local display is reconfigured.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayRemoveReconfigurationCallback(_:_:)
func CGDisplayRemoveReconfigurationCallback(callback CGDisplayReconfigurationCallBack, userInfo unsafe.Pointer) CGError {
	result, callErr := tryCGDisplayRemoveReconfigurationCallback(callback, userInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayRestoreColorSyncSettings func()
var _cGDisplayRestoreColorSyncSettingsErr error

func tryCGDisplayRestoreColorSyncSettings() error {
	if _cGDisplayRestoreColorSyncSettings == nil {
		return symbolCallError("CGDisplayRestoreColorSyncSettings", "10.0", _cGDisplayRestoreColorSyncSettingsErr)
	}
	_cGDisplayRestoreColorSyncSettings()
	return nil
}

// CGDisplayRestoreColorSyncSettings restores the gamma tables to the values in the user’s ColorSync display profile.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayRestoreColorSyncSettings()
func CGDisplayRestoreColorSyncSettings() {
	if callErr := tryCGDisplayRestoreColorSyncSettings(); callErr != nil {
		panic(callErr)
	}
}

var _cGDisplayRotation func(display uint32) float64
var _cGDisplayRotationErr error

func tryCGDisplayRotation(display uint32) (float64, error) {
	if _cGDisplayRotation == nil {
		return 0.0, symbolCallError("CGDisplayRotation", "10.5", _cGDisplayRotationErr)
	}
	return _cGDisplayRotation(display), nil
}

// CGDisplayRotation returns the rotation angle of a display in degrees.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayRotation(_:)
func CGDisplayRotation(display uint32) float64 {
	result, callErr := tryCGDisplayRotation(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayScreenSize func(display uint32) corefoundation.CGSize
var _cGDisplayScreenSizeErr error

func tryCGDisplayScreenSize(display uint32) (corefoundation.CGSize, error) {
	if _cGDisplayScreenSize == nil {
		return corefoundation.CGSize{}, symbolCallError("CGDisplayScreenSize", "10.3", _cGDisplayScreenSizeErr)
	}
	return _cGDisplayScreenSize(display), nil
}

// CGDisplayScreenSize returns the width and height of a display in millimeters.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayScreenSize(_:)
func CGDisplayScreenSize(display uint32) corefoundation.CGSize {
	result, callErr := tryCGDisplayScreenSize(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplaySerialNumber func(display uint32) uint32
var _cGDisplaySerialNumberErr error

func tryCGDisplaySerialNumber(display uint32) (uint32, error) {
	if _cGDisplaySerialNumber == nil {
		return 0, symbolCallError("CGDisplaySerialNumber", "10.2", _cGDisplaySerialNumberErr)
	}
	return _cGDisplaySerialNumber(display), nil
}

// CGDisplaySerialNumber returns the serial number of a display monitor.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplaySerialNumber(_:)
func CGDisplaySerialNumber(display uint32) uint32 {
	result, callErr := tryCGDisplaySerialNumber(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplaySetDisplayMode func(display uint32, mode CGDisplayModeRef, options corefoundation.CFDictionaryRef) CGError
var _cGDisplaySetDisplayModeErr error

func tryCGDisplaySetDisplayMode(display uint32, mode CGDisplayModeRef, options corefoundation.CFDictionaryRef) (CGError, error) {
	if _cGDisplaySetDisplayMode == nil {
		return *new(CGError), symbolCallError("CGDisplaySetDisplayMode", "10.6", _cGDisplaySetDisplayModeErr)
	}
	return _cGDisplaySetDisplayMode(display, mode, options), nil
}

// CGDisplaySetDisplayMode switches a display to a different mode.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplaySetDisplayMode(_:_:_:)
func CGDisplaySetDisplayMode(display uint32, mode CGDisplayModeRef, options corefoundation.CFDictionaryRef) CGError {
	result, callErr := tryCGDisplaySetDisplayMode(display, mode, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplaySetStereoOperation func(display uint32, stereo bool, forceBlueLine bool, option CGConfigureOption) CGError
var _cGDisplaySetStereoOperationErr error

func tryCGDisplaySetStereoOperation(display uint32, stereo bool, forceBlueLine bool, option CGConfigureOption) (CGError, error) {
	if _cGDisplaySetStereoOperation == nil {
		return *new(CGError), symbolCallError("CGDisplaySetStereoOperation", "10.4", _cGDisplaySetStereoOperationErr)
	}
	return _cGDisplaySetStereoOperation(display, stereo, forceBlueLine, option), nil
}

// CGDisplaySetStereoOperation immediately enables or disables stereo operation for a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplaySetStereoOperation(_:_:_:_:)
func CGDisplaySetStereoOperation(display uint32, stereo bool, forceBlueLine bool, option CGConfigureOption) CGError {
	result, callErr := tryCGDisplaySetStereoOperation(display, stereo, forceBlueLine, option)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayShowCursor func(display uint32) CGError
var _cGDisplayShowCursorErr error

func tryCGDisplayShowCursor(display uint32) (CGError, error) {
	if _cGDisplayShowCursor == nil {
		return *new(CGError), symbolCallError("CGDisplayShowCursor", "10.0", _cGDisplayShowCursorErr)
	}
	return _cGDisplayShowCursor(display), nil
}

// CGDisplayShowCursor decrements the hide cursor count, and shows the mouse cursor if the count is `0`.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayShowCursor(_:)
func CGDisplayShowCursor(display uint32) CGError {
	result, callErr := tryCGDisplayShowCursor(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayStreamCreate func(display uint32, outputWidth uintptr, outputHeight uintptr, pixelFormat int32, properties corefoundation.CFDictionaryRef, handler unsafe.Pointer) CGDisplayStreamRef
var _cGDisplayStreamCreateErr error

func tryCGDisplayStreamCreate(display uint32, outputWidth uintptr, outputHeight uintptr, pixelFormat int32, properties corefoundation.CFDictionaryRef, handler CGDisplayStreamFrameAvailableHandler) (CGDisplayStreamRef, error) {
	if _cGDisplayStreamCreate == nil {
		return 0, symbolCallError("CGDisplayStreamCreate", "", _cGDisplayStreamCreateErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 CGDisplayStreamFrameStatus, arg1 uint64, arg2 IOSurfaceRef, arg3 *CGDisplayStreamUpdateRef) {
		handler(arg0, arg1, arg2, arg3)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _cGDisplayStreamCreate(display, outputWidth, outputHeight, pixelFormat, properties, _block0), nil
}

// CGDisplayStreamCreate creates a new display stream to be used with a [CFRunloop].
//
// Deprecated: Please use ScreenCaptureKit instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStream/init(display:outputWidth:outputHeight:pixelFormat:properties:handler:)
func CGDisplayStreamCreate(display uint32, outputWidth uintptr, outputHeight uintptr, pixelFormat int32, properties corefoundation.CFDictionaryRef, handler CGDisplayStreamFrameAvailableHandler) CGDisplayStreamRef {
	result, callErr := tryCGDisplayStreamCreate(display, outputWidth, outputHeight, pixelFormat, properties, handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayStreamCreateWithDispatchQueue func(display uint32, outputWidth uintptr, outputHeight uintptr, pixelFormat int32, properties corefoundation.CFDictionaryRef, queue uintptr, handler unsafe.Pointer) CGDisplayStreamRef
var _cGDisplayStreamCreateWithDispatchQueueErr error

func tryCGDisplayStreamCreateWithDispatchQueue(display uint32, outputWidth uintptr, outputHeight uintptr, pixelFormat int32, properties corefoundation.CFDictionaryRef, queue dispatch.Queue, handler CGDisplayStreamFrameAvailableHandler) (CGDisplayStreamRef, error) {
	if _cGDisplayStreamCreateWithDispatchQueue == nil {
		return 0, symbolCallError("CGDisplayStreamCreateWithDispatchQueue", "", _cGDisplayStreamCreateWithDispatchQueueErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 CGDisplayStreamFrameStatus, arg1 uint64, arg2 IOSurfaceRef, arg3 *CGDisplayStreamUpdateRef) {
		handler(arg0, arg1, arg2, arg3)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _cGDisplayStreamCreateWithDispatchQueue(display, outputWidth, outputHeight, pixelFormat, properties, uintptr(queue.Handle()), _block0), nil
}

// CGDisplayStreamCreateWithDispatchQueue creates a new display stream whose updates are delivered to a dispatch queue.
//
// Deprecated: Please use ScreenCaptureKit instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStream/init(dispatchQueueDisplay:outputWidth:outputHeight:pixelFormat:properties:queue:handler:)
func CGDisplayStreamCreateWithDispatchQueue(display uint32, outputWidth uintptr, outputHeight uintptr, pixelFormat int32, properties corefoundation.CFDictionaryRef, queue dispatch.Queue, handler CGDisplayStreamFrameAvailableHandler) CGDisplayStreamRef {
	result, callErr := tryCGDisplayStreamCreateWithDispatchQueue(display, outputWidth, outputHeight, pixelFormat, properties, queue, handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayStreamGetRunLoopSource func(displayStream CGDisplayStreamRef) corefoundation.CFRunLoopSourceRef
var _cGDisplayStreamGetRunLoopSourceErr error

func tryCGDisplayStreamGetRunLoopSource(displayStream CGDisplayStreamRef) (corefoundation.CFRunLoopSourceRef, error) {
	if _cGDisplayStreamGetRunLoopSource == nil {
		return 0, symbolCallError("CGDisplayStreamGetRunLoopSource", "", _cGDisplayStreamGetRunLoopSourceErr)
	}
	return _cGDisplayStreamGetRunLoopSource(displayStream), nil
}

// CGDisplayStreamGetRunLoopSource gets the run loop source for a display stream.
//
// Deprecated: Please use ScreenCaptureKit instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStream/runLoopSource
func CGDisplayStreamGetRunLoopSource(displayStream CGDisplayStreamRef) corefoundation.CFRunLoopSourceRef {
	result, callErr := tryCGDisplayStreamGetRunLoopSource(displayStream)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayStreamGetTypeID func() uint
var _cGDisplayStreamGetTypeIDErr error

func tryCGDisplayStreamGetTypeID() (uint, error) {
	if _cGDisplayStreamGetTypeID == nil {
		return 0, symbolCallError("CGDisplayStreamGetTypeID", "", _cGDisplayStreamGetTypeIDErr)
	}
	return _cGDisplayStreamGetTypeID(), nil
}

// CGDisplayStreamGetTypeID returns the type identifier of a Quartz display stream.
//
// Deprecated: Please use ScreenCaptureKit instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStream/typeID
func CGDisplayStreamGetTypeID() uint {
	result, callErr := tryCGDisplayStreamGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayStreamStart func(displayStream CGDisplayStreamRef) CGError
var _cGDisplayStreamStartErr error

func tryCGDisplayStreamStart(displayStream CGDisplayStreamRef) (CGError, error) {
	if _cGDisplayStreamStart == nil {
		return *new(CGError), symbolCallError("CGDisplayStreamStart", "", _cGDisplayStreamStartErr)
	}
	return _cGDisplayStreamStart(displayStream), nil
}

// CGDisplayStreamStart tells a stream to start sending updates.
//
// Deprecated: Please use ScreenCaptureKit instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStream/start()
func CGDisplayStreamStart(displayStream CGDisplayStreamRef) CGError {
	result, callErr := tryCGDisplayStreamStart(displayStream)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayStreamStop func(displayStream CGDisplayStreamRef) CGError
var _cGDisplayStreamStopErr error

func tryCGDisplayStreamStop(displayStream CGDisplayStreamRef) (CGError, error) {
	if _cGDisplayStreamStop == nil {
		return *new(CGError), symbolCallError("CGDisplayStreamStop", "", _cGDisplayStreamStopErr)
	}
	return _cGDisplayStreamStop(displayStream), nil
}

// CGDisplayStreamStop tells a stream to stop sending updates.
//
// Deprecated: Please use ScreenCaptureKit instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStream/stop()
func CGDisplayStreamStop(displayStream CGDisplayStreamRef) CGError {
	result, callErr := tryCGDisplayStreamStop(displayStream)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayStreamUpdateCreateMergedUpdate func(firstUpdate CGDisplayStreamUpdateRef, secondUpdate CGDisplayStreamUpdateRef) CGDisplayStreamUpdateRef
var _cGDisplayStreamUpdateCreateMergedUpdateErr error

func tryCGDisplayStreamUpdateCreateMergedUpdate(firstUpdate CGDisplayStreamUpdateRef, secondUpdate CGDisplayStreamUpdateRef) (CGDisplayStreamUpdateRef, error) {
	if _cGDisplayStreamUpdateCreateMergedUpdate == nil {
		return 0, symbolCallError("CGDisplayStreamUpdateCreateMergedUpdate", "", _cGDisplayStreamUpdateCreateMergedUpdateErr)
	}
	return _cGDisplayStreamUpdateCreateMergedUpdate(firstUpdate, secondUpdate), nil
}

// CGDisplayStreamUpdateCreateMergedUpdate combines two updates into a new update that includes the metadata for both source updates.
//
// Deprecated: Please use ScreenCaptureKit instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdate/init(mergedUpdateFirstUpdate:secondUpdate:)
func CGDisplayStreamUpdateCreateMergedUpdate(firstUpdate CGDisplayStreamUpdateRef, secondUpdate CGDisplayStreamUpdateRef) CGDisplayStreamUpdateRef {
	result, callErr := tryCGDisplayStreamUpdateCreateMergedUpdate(firstUpdate, secondUpdate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayStreamUpdateGetDropCount func(updateRef CGDisplayStreamUpdateRef) uintptr
var _cGDisplayStreamUpdateGetDropCountErr error

func tryCGDisplayStreamUpdateGetDropCount(updateRef CGDisplayStreamUpdateRef) (uintptr, error) {
	if _cGDisplayStreamUpdateGetDropCount == nil {
		return 0, symbolCallError("CGDisplayStreamUpdateGetDropCount", "", _cGDisplayStreamUpdateGetDropCountErr)
	}
	return _cGDisplayStreamUpdateGetDropCount(updateRef), nil
}

// CGDisplayStreamUpdateGetDropCount returns the number of frames that have been dropped since the last call to your update handler.
//
// Deprecated: Please use ScreenCaptureKit instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdate/dropCount
func CGDisplayStreamUpdateGetDropCount(updateRef CGDisplayStreamUpdateRef) uintptr {
	result, callErr := tryCGDisplayStreamUpdateGetDropCount(updateRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayStreamUpdateGetMovedRectsDelta func(updateRef CGDisplayStreamUpdateRef, dx *float64, dy *float64)
var _cGDisplayStreamUpdateGetMovedRectsDeltaErr error

func tryCGDisplayStreamUpdateGetMovedRectsDelta(updateRef CGDisplayStreamUpdateRef, dx *float64, dy *float64) error {
	if _cGDisplayStreamUpdateGetMovedRectsDelta == nil {
		return symbolCallError("CGDisplayStreamUpdateGetMovedRectsDelta", "", _cGDisplayStreamUpdateGetMovedRectsDeltaErr)
	}
	_cGDisplayStreamUpdateGetMovedRectsDelta(updateRef, dx, dy)
	return nil
}

// CGDisplayStreamUpdateGetMovedRectsDelta return the movement delta values for a single update.
//
// Deprecated: Please use ScreenCaptureKit instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdate/getMovedRectsDelta(dx:dy:)
func CGDisplayStreamUpdateGetMovedRectsDelta(updateRef CGDisplayStreamUpdateRef, dx *float64, dy *float64) {
	if callErr := tryCGDisplayStreamUpdateGetMovedRectsDelta(updateRef, dx, dy); callErr != nil {
		panic(callErr)
	}
}

var _cGDisplayStreamUpdateGetRects func(updateRef CGDisplayStreamUpdateRef, rectType CGDisplayStreamUpdateRectType, rectCount *uintptr) *corefoundation.CGRect
var _cGDisplayStreamUpdateGetRectsErr error

func tryCGDisplayStreamUpdateGetRects(updateRef CGDisplayStreamUpdateRef, rectType CGDisplayStreamUpdateRectType, rectCount *uintptr) (*corefoundation.CGRect, error) {
	if _cGDisplayStreamUpdateGetRects == nil {
		return nil, symbolCallError("CGDisplayStreamUpdateGetRects", "", _cGDisplayStreamUpdateGetRectsErr)
	}
	return _cGDisplayStreamUpdateGetRects(updateRef, rectType, rectCount), nil
}

// CGDisplayStreamUpdateGetRects returns an array of rectangles that describe where the frame has changed since the previous frame.
//
// Deprecated: Please use ScreenCaptureKit instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdate/getRects(_:rectCount:)
func CGDisplayStreamUpdateGetRects(updateRef CGDisplayStreamUpdateRef, rectType CGDisplayStreamUpdateRectType, rectCount *uintptr) *corefoundation.CGRect {
	result, callErr := tryCGDisplayStreamUpdateGetRects(updateRef, rectType, rectCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayStreamUpdateGetTypeID func() uint
var _cGDisplayStreamUpdateGetTypeIDErr error

func tryCGDisplayStreamUpdateGetTypeID() (uint, error) {
	if _cGDisplayStreamUpdateGetTypeID == nil {
		return 0, symbolCallError("CGDisplayStreamUpdateGetTypeID", "", _cGDisplayStreamUpdateGetTypeIDErr)
	}
	return _cGDisplayStreamUpdateGetTypeID(), nil
}

// CGDisplayStreamUpdateGetTypeID returns the type identifier of a Quartz display stream update.
//
// Deprecated: Please use ScreenCaptureKit instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdate/typeID
func CGDisplayStreamUpdateGetTypeID() uint {
	result, callErr := tryCGDisplayStreamUpdateGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayUnitNumber func(display uint32) uint32
var _cGDisplayUnitNumberErr error

func tryCGDisplayUnitNumber(display uint32) (uint32, error) {
	if _cGDisplayUnitNumber == nil {
		return 0, symbolCallError("CGDisplayUnitNumber", "10.2", _cGDisplayUnitNumberErr)
	}
	return _cGDisplayUnitNumber(display), nil
}

// CGDisplayUnitNumber returns the logical unit number of a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayUnitNumber(_:)
func CGDisplayUnitNumber(display uint32) uint32 {
	result, callErr := tryCGDisplayUnitNumber(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayUsesOpenGLAcceleration func(display uint32) bool
var _cGDisplayUsesOpenGLAccelerationErr error

func tryCGDisplayUsesOpenGLAcceleration(display uint32) (bool, error) {
	if _cGDisplayUsesOpenGLAcceleration == nil {
		return false, symbolCallError("CGDisplayUsesOpenGLAcceleration", "10.2", _cGDisplayUsesOpenGLAccelerationErr)
	}
	return _cGDisplayUsesOpenGLAcceleration(display), nil
}

// CGDisplayUsesOpenGLAcceleration returns a Boolean value indicating whether Quartz is using OpenGL-based window acceleration (Quartz Extreme) to render in a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayUsesOpenGLAcceleration(_:)
func CGDisplayUsesOpenGLAcceleration(display uint32) bool {
	result, callErr := tryCGDisplayUsesOpenGLAcceleration(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGDisplayVendorNumber func(display uint32) uint32
var _cGDisplayVendorNumberErr error

func tryCGDisplayVendorNumber(display uint32) (uint32, error) {
	if _cGDisplayVendorNumber == nil {
		return 0, symbolCallError("CGDisplayVendorNumber", "10.2", _cGDisplayVendorNumberErr)
	}
	return _cGDisplayVendorNumber(display), nil
}

// CGDisplayVendorNumber returns the vendor number of the specified display’s monitor.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayVendorNumber(_:)
func CGDisplayVendorNumber(display uint32) uint32 {
	result, callErr := tryCGDisplayVendorNumber(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEXRToneMappingGammaGetDefaultOptions func() corefoundation.CFDictionaryRef
var _cGEXRToneMappingGammaGetDefaultOptionsErr error

func tryCGEXRToneMappingGammaGetDefaultOptions() (corefoundation.CFDictionaryRef, error) {
	if _cGEXRToneMappingGammaGetDefaultOptions == nil {
		return 0, symbolCallError("CGEXRToneMappingGammaGetDefaultOptions", "26.0", _cGEXRToneMappingGammaGetDefaultOptionsErr)
	}
	return _cGEXRToneMappingGammaGetDefaultOptions(), nil
}

// CGEXRToneMappingGammaGetDefaultOptions.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEXRToneMappingGammaGetDefaultOptions
func CGEXRToneMappingGammaGetDefaultOptions() corefoundation.CFDictionaryRef {
	result, callErr := tryCGEXRToneMappingGammaGetDefaultOptions()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGErrorSetCallback func(callback CGErrorCallback)
var _cGErrorSetCallbackErr error

func tryCGErrorSetCallback(callback CGErrorCallback) error {
	if _cGErrorSetCallback == nil {
		return symbolCallError("CGErrorSetCallback", "", _cGErrorSetCallbackErr)
	}
	_cGErrorSetCallback(callback)
	return nil
}

// CGErrorSetCallback.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGErrorSetCallback(_:)
func CGErrorSetCallback(callback CGErrorCallback) {
	if callErr := tryCGErrorSetCallback(callback); callErr != nil {
		panic(callErr)
	}
}

var _cGEventCreate func(source CGEventSourceRef) CGEventRef
var _cGEventCreateErr error

func tryCGEventCreate(source CGEventSourceRef) (CGEventRef, error) {
	if _cGEventCreate == nil {
		return 0, symbolCallError("CGEventCreate", "10.4", _cGEventCreateErr)
	}
	return _cGEventCreate(source), nil
}

// CGEventCreate returns a new Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/init(source:)
func CGEventCreate(source CGEventSourceRef) CGEventRef {
	result, callErr := tryCGEventCreate(source)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventCreateCopy func(event CGEventRef) CGEventRef
var _cGEventCreateCopyErr error

func tryCGEventCreateCopy(event CGEventRef) (CGEventRef, error) {
	if _cGEventCreateCopy == nil {
		return 0, symbolCallError("CGEventCreateCopy", "10.4", _cGEventCreateCopyErr)
	}
	return _cGEventCreateCopy(event), nil
}

// CGEventCreateCopy returns a copy of an existing Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/copy()
func CGEventCreateCopy(event CGEventRef) CGEventRef {
	result, callErr := tryCGEventCreateCopy(event)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventCreateData func(allocator corefoundation.CFAllocatorRef, event CGEventRef) corefoundation.CFDataRef
var _cGEventCreateDataErr error

func tryCGEventCreateData(allocator corefoundation.CFAllocatorRef, event CGEventRef) (corefoundation.CFDataRef, error) {
	if _cGEventCreateData == nil {
		return 0, symbolCallError("CGEventCreateData", "10.4", _cGEventCreateDataErr)
	}
	return _cGEventCreateData(allocator, event), nil
}

// CGEventCreateData returns a flattened data representation of a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventCreateData
func CGEventCreateData(allocator corefoundation.CFAllocatorRef, event CGEventRef) corefoundation.CFDataRef {
	result, callErr := tryCGEventCreateData(allocator, event)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventCreateFromData func(allocator corefoundation.CFAllocatorRef, data corefoundation.CFDataRef) CGEventRef
var _cGEventCreateFromDataErr error

func tryCGEventCreateFromData(allocator corefoundation.CFAllocatorRef, data corefoundation.CFDataRef) (CGEventRef, error) {
	if _cGEventCreateFromData == nil {
		return 0, symbolCallError("CGEventCreateFromData", "10.4", _cGEventCreateFromDataErr)
	}
	return _cGEventCreateFromData(allocator, data), nil
}

// CGEventCreateFromData returns a Quartz event created from a flattened data representation of the event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/init(withDataAllocator:data:)
func CGEventCreateFromData(allocator corefoundation.CFAllocatorRef, data corefoundation.CFDataRef) CGEventRef {
	result, callErr := tryCGEventCreateFromData(allocator, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventCreateKeyboardEvent func(source CGEventSourceRef, virtualKey uint16, keyDown bool) CGEventRef
var _cGEventCreateKeyboardEventErr error

func tryCGEventCreateKeyboardEvent(source CGEventSourceRef, virtualKey uint16, keyDown bool) (CGEventRef, error) {
	if _cGEventCreateKeyboardEvent == nil {
		return 0, symbolCallError("CGEventCreateKeyboardEvent", "10.4", _cGEventCreateKeyboardEventErr)
	}
	return _cGEventCreateKeyboardEvent(source, virtualKey, keyDown), nil
}

// CGEventCreateKeyboardEvent returns a new Quartz keyboard event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/init(keyboardEventSource:virtualKey:keyDown:)
func CGEventCreateKeyboardEvent(source CGEventSourceRef, virtualKey uint16, keyDown bool) CGEventRef {
	result, callErr := tryCGEventCreateKeyboardEvent(source, virtualKey, keyDown)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventCreateMouseEvent func(source CGEventSourceRef, mouseType CGEventType, mouseCursorPosition corefoundation.CGPoint, mouseButton CGMouseButton) CGEventRef
var _cGEventCreateMouseEventErr error

func tryCGEventCreateMouseEvent(source CGEventSourceRef, mouseType CGEventType, mouseCursorPosition corefoundation.CGPoint, mouseButton CGMouseButton) (CGEventRef, error) {
	if _cGEventCreateMouseEvent == nil {
		return 0, symbolCallError("CGEventCreateMouseEvent", "10.4", _cGEventCreateMouseEventErr)
	}
	return _cGEventCreateMouseEvent(source, mouseType, mouseCursorPosition, mouseButton), nil
}

// CGEventCreateMouseEvent returns a new Quartz mouse event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/init(mouseEventSource:mouseType:mouseCursorPosition:mouseButton:)
func CGEventCreateMouseEvent(source CGEventSourceRef, mouseType CGEventType, mouseCursorPosition corefoundation.CGPoint, mouseButton CGMouseButton) CGEventRef {
	result, callErr := tryCGEventCreateMouseEvent(source, mouseType, mouseCursorPosition, mouseButton)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventCreateScrollWheelEvent func(source CGEventSourceRef, units CGScrollEventUnit, wheelCount uint32, wheel1 int32) CGEventRef
var _cGEventCreateScrollWheelEventErr error

func tryCGEventCreateScrollWheelEvent(source CGEventSourceRef, units CGScrollEventUnit, wheelCount uint32, wheel1 int32) (CGEventRef, error) {
	if _cGEventCreateScrollWheelEvent == nil {
		return 0, symbolCallError("CGEventCreateScrollWheelEvent", "10.5", _cGEventCreateScrollWheelEventErr)
	}
	return _cGEventCreateScrollWheelEvent(source, units, wheelCount, wheel1), nil
}

// CGEventCreateScrollWheelEvent returns a new Quartz scrolling event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventCreateScrollWheelEvent
func CGEventCreateScrollWheelEvent(source CGEventSourceRef, units CGScrollEventUnit, wheelCount uint32, wheel1 int32) CGEventRef {
	result, callErr := tryCGEventCreateScrollWheelEvent(source, units, wheelCount, wheel1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventCreateScrollWheelEvent2 func(source CGEventSourceRef, units CGScrollEventUnit, wheelCount uint32, wheel1 int32, wheel2 int32, wheel3 int32) CGEventRef
var _cGEventCreateScrollWheelEvent2Err error

func tryCGEventCreateScrollWheelEvent2(source CGEventSourceRef, units CGScrollEventUnit, wheelCount uint32, wheel1 int32, wheel2 int32, wheel3 int32) (CGEventRef, error) {
	if _cGEventCreateScrollWheelEvent2 == nil {
		return 0, symbolCallError("CGEventCreateScrollWheelEvent2", "10.13", _cGEventCreateScrollWheelEvent2Err)
	}
	return _cGEventCreateScrollWheelEvent2(source, units, wheelCount, wheel1, wheel2, wheel3), nil
}

// CGEventCreateScrollWheelEvent2.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/init(scrollWheelEvent2Source:units:wheelCount:wheel1:wheel2:wheel3:)
func CGEventCreateScrollWheelEvent2(source CGEventSourceRef, units CGScrollEventUnit, wheelCount uint32, wheel1 int32, wheel2 int32, wheel3 int32) CGEventRef {
	result, callErr := tryCGEventCreateScrollWheelEvent2(source, units, wheelCount, wheel1, wheel2, wheel3)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventCreateSourceFromEvent func(event CGEventRef) CGEventSourceRef
var _cGEventCreateSourceFromEventErr error

func tryCGEventCreateSourceFromEvent(event CGEventRef) (CGEventSourceRef, error) {
	if _cGEventCreateSourceFromEvent == nil {
		return 0, symbolCallError("CGEventCreateSourceFromEvent", "10.4", _cGEventCreateSourceFromEventErr)
	}
	return _cGEventCreateSourceFromEvent(event), nil
}

// CGEventCreateSourceFromEvent returns a Quartz event source created from an existing Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/init(event:)
func CGEventCreateSourceFromEvent(event CGEventRef) CGEventSourceRef {
	result, callErr := tryCGEventCreateSourceFromEvent(event)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventGetDoubleValueField func(event CGEventRef, field CGEventField) float64
var _cGEventGetDoubleValueFieldErr error

func tryCGEventGetDoubleValueField(event CGEventRef, field CGEventField) (float64, error) {
	if _cGEventGetDoubleValueField == nil {
		return 0.0, symbolCallError("CGEventGetDoubleValueField", "10.4", _cGEventGetDoubleValueFieldErr)
	}
	return _cGEventGetDoubleValueField(event, field), nil
}

// CGEventGetDoubleValueField returns the floating-point value of a field in a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/getDoubleValueField(_:)
func CGEventGetDoubleValueField(event CGEventRef, field CGEventField) float64 {
	result, callErr := tryCGEventGetDoubleValueField(event, field)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventGetFlags func(event CGEventRef) CGEventFlags
var _cGEventGetFlagsErr error

func tryCGEventGetFlags(event CGEventRef) (CGEventFlags, error) {
	if _cGEventGetFlags == nil {
		return *new(CGEventFlags), symbolCallError("CGEventGetFlags", "10.4", _cGEventGetFlagsErr)
	}
	return _cGEventGetFlags(event), nil
}

// CGEventGetFlags returns the event flags of a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/flags
func CGEventGetFlags(event CGEventRef) CGEventFlags {
	result, callErr := tryCGEventGetFlags(event)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventGetIntegerValueField func(event CGEventRef, field CGEventField) int64
var _cGEventGetIntegerValueFieldErr error

func tryCGEventGetIntegerValueField(event CGEventRef, field CGEventField) (int64, error) {
	if _cGEventGetIntegerValueField == nil {
		return 0, symbolCallError("CGEventGetIntegerValueField", "10.4", _cGEventGetIntegerValueFieldErr)
	}
	return _cGEventGetIntegerValueField(event, field), nil
}

// CGEventGetIntegerValueField returns the integer value of a field in a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/getIntegerValueField(_:)
func CGEventGetIntegerValueField(event CGEventRef, field CGEventField) int64 {
	result, callErr := tryCGEventGetIntegerValueField(event, field)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventGetLocation func(event CGEventRef) corefoundation.CGPoint
var _cGEventGetLocationErr error

func tryCGEventGetLocation(event CGEventRef) (corefoundation.CGPoint, error) {
	if _cGEventGetLocation == nil {
		return corefoundation.CGPoint{}, symbolCallError("CGEventGetLocation", "10.4", _cGEventGetLocationErr)
	}
	return _cGEventGetLocation(event), nil
}

// CGEventGetLocation returns the location of a Quartz mouse event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/location
func CGEventGetLocation(event CGEventRef) corefoundation.CGPoint {
	result, callErr := tryCGEventGetLocation(event)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventGetTimestamp func(event CGEventRef) CGEventTimestamp
var _cGEventGetTimestampErr error

func tryCGEventGetTimestamp(event CGEventRef) (CGEventTimestamp, error) {
	if _cGEventGetTimestamp == nil {
		return *new(CGEventTimestamp), symbolCallError("CGEventGetTimestamp", "10.4", _cGEventGetTimestampErr)
	}
	return _cGEventGetTimestamp(event), nil
}

// CGEventGetTimestamp returns the timestamp of a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/timestamp
func CGEventGetTimestamp(event CGEventRef) CGEventTimestamp {
	result, callErr := tryCGEventGetTimestamp(event)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventGetType func(event CGEventRef) CGEventType
var _cGEventGetTypeErr error

func tryCGEventGetType(event CGEventRef) (CGEventType, error) {
	if _cGEventGetType == nil {
		return *new(CGEventType), symbolCallError("CGEventGetType", "10.4", _cGEventGetTypeErr)
	}
	return _cGEventGetType(event), nil
}

// CGEventGetType returns the event type of a Quartz event (left mouse down, for example).
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/type
func CGEventGetType(event CGEventRef) CGEventType {
	result, callErr := tryCGEventGetType(event)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventGetTypeID func() uint
var _cGEventGetTypeIDErr error

func tryCGEventGetTypeID() (uint, error) {
	if _cGEventGetTypeID == nil {
		return 0, symbolCallError("CGEventGetTypeID", "10.4", _cGEventGetTypeIDErr)
	}
	return _cGEventGetTypeID(), nil
}

// CGEventGetTypeID returns the type identifier for the opaque type [CGEventRef].
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/typeID
func CGEventGetTypeID() uint {
	result, callErr := tryCGEventGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventGetUnflippedLocation func(event CGEventRef) corefoundation.CGPoint
var _cGEventGetUnflippedLocationErr error

func tryCGEventGetUnflippedLocation(event CGEventRef) (corefoundation.CGPoint, error) {
	if _cGEventGetUnflippedLocation == nil {
		return corefoundation.CGPoint{}, symbolCallError("CGEventGetUnflippedLocation", "10.5", _cGEventGetUnflippedLocationErr)
	}
	return _cGEventGetUnflippedLocation(event), nil
}

// CGEventGetUnflippedLocation returns the location of a Quartz mouse event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/unflippedLocation
func CGEventGetUnflippedLocation(event CGEventRef) corefoundation.CGPoint {
	result, callErr := tryCGEventGetUnflippedLocation(event)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventKeyboardGetUnicodeString func(event CGEventRef, maxStringLength uint, actualStringLength *uint, unicodeString *uint16)
var _cGEventKeyboardGetUnicodeStringErr error

func tryCGEventKeyboardGetUnicodeString(event CGEventRef, maxStringLength uint, actualStringLength *uint, unicodeString *uint16) error {
	if _cGEventKeyboardGetUnicodeString == nil {
		return symbolCallError("CGEventKeyboardGetUnicodeString", "10.4", _cGEventKeyboardGetUnicodeStringErr)
	}
	_cGEventKeyboardGetUnicodeString(event, maxStringLength, actualStringLength, unicodeString)
	return nil
}

// CGEventKeyboardGetUnicodeString returns the Unicode string associated with a Quartz keyboard event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/keyboardGetUnicodeString(maxStringLength:actualStringLength:unicodeString:)
func CGEventKeyboardGetUnicodeString(event CGEventRef, maxStringLength uint, actualStringLength *uint, unicodeString *uint16) {
	if callErr := tryCGEventKeyboardGetUnicodeString(event, maxStringLength, actualStringLength, unicodeString); callErr != nil {
		panic(callErr)
	}
}

var _cGEventKeyboardSetUnicodeString func(event CGEventRef, stringLength uint, unicodeString *uint16)
var _cGEventKeyboardSetUnicodeStringErr error

func tryCGEventKeyboardSetUnicodeString(event CGEventRef, stringLength uint, unicodeString *uint16) error {
	if _cGEventKeyboardSetUnicodeString == nil {
		return symbolCallError("CGEventKeyboardSetUnicodeString", "10.4", _cGEventKeyboardSetUnicodeStringErr)
	}
	_cGEventKeyboardSetUnicodeString(event, stringLength, unicodeString)
	return nil
}

// CGEventKeyboardSetUnicodeString sets the Unicode string associated with a Quartz keyboard event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/keyboardSetUnicodeString(stringLength:unicodeString:)
func CGEventKeyboardSetUnicodeString(event CGEventRef, stringLength uint, unicodeString *uint16) {
	if callErr := tryCGEventKeyboardSetUnicodeString(event, stringLength, unicodeString); callErr != nil {
		panic(callErr)
	}
}

var _cGEventPost func(tap CGEventTapLocation, event CGEventRef)
var _cGEventPostErr error

func tryCGEventPost(tap CGEventTapLocation, event CGEventRef) error {
	if _cGEventPost == nil {
		return symbolCallError("CGEventPost", "10.4", _cGEventPostErr)
	}
	_cGEventPost(tap, event)
	return nil
}

// CGEventPost posts a Quartz event into the event stream at a specified location.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/post(tap:)
func CGEventPost(tap CGEventTapLocation, event CGEventRef) {
	if callErr := tryCGEventPost(tap, event); callErr != nil {
		panic(callErr)
	}
}

var _cGEventPostToPSN func(processSerialNumber unsafe.Pointer, event CGEventRef)
var _cGEventPostToPSNErr error

func tryCGEventPostToPSN(processSerialNumber unsafe.Pointer, event CGEventRef) error {
	if _cGEventPostToPSN == nil {
		return symbolCallError("CGEventPostToPSN", "10.4", _cGEventPostToPSNErr)
	}
	_cGEventPostToPSN(processSerialNumber, event)
	return nil
}

// CGEventPostToPSN posts a Quartz event into the event stream for a specific application.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/postToPSN(processSerialNumber:)
func CGEventPostToPSN(processSerialNumber unsafe.Pointer, event CGEventRef) {
	if callErr := tryCGEventPostToPSN(processSerialNumber, event); callErr != nil {
		panic(callErr)
	}
}

var _cGEventPostToPid func(pid int32, event CGEventRef)
var _cGEventPostToPidErr error

func tryCGEventPostToPid(pid int32, event CGEventRef) error {
	if _cGEventPostToPid == nil {
		return symbolCallError("CGEventPostToPid", "10.11", _cGEventPostToPidErr)
	}
	_cGEventPostToPid(pid, event)
	return nil
}

// CGEventPostToPid.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/postToPid(_:)
func CGEventPostToPid(pid int32, event CGEventRef) {
	if callErr := tryCGEventPostToPid(pid, event); callErr != nil {
		panic(callErr)
	}
}

var _cGEventSetDoubleValueField func(event CGEventRef, field CGEventField, value float64)
var _cGEventSetDoubleValueFieldErr error

func tryCGEventSetDoubleValueField(event CGEventRef, field CGEventField, value float64) error {
	if _cGEventSetDoubleValueField == nil {
		return symbolCallError("CGEventSetDoubleValueField", "10.4", _cGEventSetDoubleValueFieldErr)
	}
	_cGEventSetDoubleValueField(event, field, value)
	return nil
}

// CGEventSetDoubleValueField sets the floating-point value of a field in a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/setDoubleValueField(_:value:)
func CGEventSetDoubleValueField(event CGEventRef, field CGEventField, value float64) {
	if callErr := tryCGEventSetDoubleValueField(event, field, value); callErr != nil {
		panic(callErr)
	}
}

var _cGEventSetFlags func(event CGEventRef, flags CGEventFlags)
var _cGEventSetFlagsErr error

func tryCGEventSetFlags(event CGEventRef, flags CGEventFlags) error {
	if _cGEventSetFlags == nil {
		return symbolCallError("CGEventSetFlags", "10.4", _cGEventSetFlagsErr)
	}
	_cGEventSetFlags(event, flags)
	return nil
}

// CGEventSetFlags sets the event flags of a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSetFlags
func CGEventSetFlags(event CGEventRef, flags CGEventFlags) {
	if callErr := tryCGEventSetFlags(event, flags); callErr != nil {
		panic(callErr)
	}
}

var _cGEventSetIntegerValueField func(event CGEventRef, field CGEventField, value int64)
var _cGEventSetIntegerValueFieldErr error

func tryCGEventSetIntegerValueField(event CGEventRef, field CGEventField, value int64) error {
	if _cGEventSetIntegerValueField == nil {
		return symbolCallError("CGEventSetIntegerValueField", "10.4", _cGEventSetIntegerValueFieldErr)
	}
	_cGEventSetIntegerValueField(event, field, value)
	return nil
}

// CGEventSetIntegerValueField sets the integer value of a field in a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/setIntegerValueField(_:value:)
func CGEventSetIntegerValueField(event CGEventRef, field CGEventField, value int64) {
	if callErr := tryCGEventSetIntegerValueField(event, field, value); callErr != nil {
		panic(callErr)
	}
}

var _cGEventSetLocation func(event CGEventRef, location corefoundation.CGPoint)
var _cGEventSetLocationErr error

func tryCGEventSetLocation(event CGEventRef, location corefoundation.CGPoint) error {
	if _cGEventSetLocation == nil {
		return symbolCallError("CGEventSetLocation", "10.4", _cGEventSetLocationErr)
	}
	_cGEventSetLocation(event, location)
	return nil
}

// CGEventSetLocation sets the location of a Quartz mouse event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSetLocation
func CGEventSetLocation(event CGEventRef, location corefoundation.CGPoint) {
	if callErr := tryCGEventSetLocation(event, location); callErr != nil {
		panic(callErr)
	}
}

var _cGEventSetSource func(event CGEventRef, source CGEventSourceRef)
var _cGEventSetSourceErr error

func tryCGEventSetSource(event CGEventRef, source CGEventSourceRef) error {
	if _cGEventSetSource == nil {
		return symbolCallError("CGEventSetSource", "10.4", _cGEventSetSourceErr)
	}
	_cGEventSetSource(event, source)
	return nil
}

// CGEventSetSource sets the event source of a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/setSource(_:)
func CGEventSetSource(event CGEventRef, source CGEventSourceRef) {
	if callErr := tryCGEventSetSource(event, source); callErr != nil {
		panic(callErr)
	}
}

var _cGEventSetTimestamp func(event CGEventRef, timestamp CGEventTimestamp)
var _cGEventSetTimestampErr error

func tryCGEventSetTimestamp(event CGEventRef, timestamp CGEventTimestamp) error {
	if _cGEventSetTimestamp == nil {
		return symbolCallError("CGEventSetTimestamp", "10.4", _cGEventSetTimestampErr)
	}
	_cGEventSetTimestamp(event, timestamp)
	return nil
}

// CGEventSetTimestamp sets the timestamp of a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSetTimestamp
func CGEventSetTimestamp(event CGEventRef, timestamp CGEventTimestamp) {
	if callErr := tryCGEventSetTimestamp(event, timestamp); callErr != nil {
		panic(callErr)
	}
}

var _cGEventSetType func(event CGEventRef, type_ CGEventType)
var _cGEventSetTypeErr error

func tryCGEventSetType(event CGEventRef, type_ CGEventType) error {
	if _cGEventSetType == nil {
		return symbolCallError("CGEventSetType", "10.4", _cGEventSetTypeErr)
	}
	_cGEventSetType(event, type_)
	return nil
}

// CGEventSetType sets the event type of a Quartz event (left mouse down, for example).
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSetType
func CGEventSetType(event CGEventRef, type_ CGEventType) {
	if callErr := tryCGEventSetType(event, type_); callErr != nil {
		panic(callErr)
	}
}

var _cGEventSourceButtonState func(stateID CGEventSourceStateID, button CGMouseButton) bool
var _cGEventSourceButtonStateErr error

func tryCGEventSourceButtonState(stateID CGEventSourceStateID, button CGMouseButton) (bool, error) {
	if _cGEventSourceButtonState == nil {
		return false, symbolCallError("CGEventSourceButtonState", "10.4", _cGEventSourceButtonStateErr)
	}
	return _cGEventSourceButtonState(stateID, button), nil
}

// CGEventSourceButtonState returns a Boolean value indicating the current button state of a Quartz event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/buttonState(_:button:)
func CGEventSourceButtonState(stateID CGEventSourceStateID, button CGMouseButton) bool {
	result, callErr := tryCGEventSourceButtonState(stateID, button)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventSourceCounterForEventType func(stateID CGEventSourceStateID, eventType CGEventType) uint32
var _cGEventSourceCounterForEventTypeErr error

func tryCGEventSourceCounterForEventType(stateID CGEventSourceStateID, eventType CGEventType) (uint32, error) {
	if _cGEventSourceCounterForEventType == nil {
		return 0, symbolCallError("CGEventSourceCounterForEventType", "10.4", _cGEventSourceCounterForEventTypeErr)
	}
	return _cGEventSourceCounterForEventType(stateID, eventType), nil
}

// CGEventSourceCounterForEventType returns a count of events of a given type seen since the window server started.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/counterForEventType(_:eventType:)
func CGEventSourceCounterForEventType(stateID CGEventSourceStateID, eventType CGEventType) uint32 {
	result, callErr := tryCGEventSourceCounterForEventType(stateID, eventType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventSourceCreate func(stateID CGEventSourceStateID) CGEventSourceRef
var _cGEventSourceCreateErr error

func tryCGEventSourceCreate(stateID CGEventSourceStateID) (CGEventSourceRef, error) {
	if _cGEventSourceCreate == nil {
		return 0, symbolCallError("CGEventSourceCreate", "10.4", _cGEventSourceCreateErr)
	}
	return _cGEventSourceCreate(stateID), nil
}

// CGEventSourceCreate returns a Quartz event source created with a specified source state.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/init(stateID:)
func CGEventSourceCreate(stateID CGEventSourceStateID) CGEventSourceRef {
	result, callErr := tryCGEventSourceCreate(stateID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventSourceFlagsState func(stateID CGEventSourceStateID) CGEventFlags
var _cGEventSourceFlagsStateErr error

func tryCGEventSourceFlagsState(stateID CGEventSourceStateID) (CGEventFlags, error) {
	if _cGEventSourceFlagsState == nil {
		return *new(CGEventFlags), symbolCallError("CGEventSourceFlagsState", "10.4", _cGEventSourceFlagsStateErr)
	}
	return _cGEventSourceFlagsState(stateID), nil
}

// CGEventSourceFlagsState returns the current flags of a Quartz event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/flagsState(_:)
func CGEventSourceFlagsState(stateID CGEventSourceStateID) CGEventFlags {
	result, callErr := tryCGEventSourceFlagsState(stateID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventSourceGetKeyboardType func(source CGEventSourceRef) CGEventSourceKeyboardType
var _cGEventSourceGetKeyboardTypeErr error

func tryCGEventSourceGetKeyboardType(source CGEventSourceRef) (CGEventSourceKeyboardType, error) {
	if _cGEventSourceGetKeyboardType == nil {
		return *new(CGEventSourceKeyboardType), symbolCallError("CGEventSourceGetKeyboardType", "10.4", _cGEventSourceGetKeyboardTypeErr)
	}
	return _cGEventSourceGetKeyboardType(source), nil
}

// CGEventSourceGetKeyboardType returns the keyboard type to be used with a Quartz event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/keyboardType
func CGEventSourceGetKeyboardType(source CGEventSourceRef) CGEventSourceKeyboardType {
	result, callErr := tryCGEventSourceGetKeyboardType(source)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventSourceGetLocalEventsFilterDuringSuppressionState func(source CGEventSourceRef, state CGEventSuppressionState) CGEventFilterMask
var _cGEventSourceGetLocalEventsFilterDuringSuppressionStateErr error

func tryCGEventSourceGetLocalEventsFilterDuringSuppressionState(source CGEventSourceRef, state CGEventSuppressionState) (CGEventFilterMask, error) {
	if _cGEventSourceGetLocalEventsFilterDuringSuppressionState == nil {
		return *new(CGEventFilterMask), symbolCallError("CGEventSourceGetLocalEventsFilterDuringSuppressionState", "10.4", _cGEventSourceGetLocalEventsFilterDuringSuppressionStateErr)
	}
	return _cGEventSourceGetLocalEventsFilterDuringSuppressionState(source, state), nil
}

// CGEventSourceGetLocalEventsFilterDuringSuppressionState returns the mask that indicates which classes of local hardware events are enabled during event suppression.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/getLocalEventsFilterDuringSuppressionState(_:)
func CGEventSourceGetLocalEventsFilterDuringSuppressionState(source CGEventSourceRef, state CGEventSuppressionState) CGEventFilterMask {
	result, callErr := tryCGEventSourceGetLocalEventsFilterDuringSuppressionState(source, state)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventSourceGetLocalEventsSuppressionInterval func(source CGEventSourceRef) float64
var _cGEventSourceGetLocalEventsSuppressionIntervalErr error

func tryCGEventSourceGetLocalEventsSuppressionInterval(source CGEventSourceRef) (float64, error) {
	if _cGEventSourceGetLocalEventsSuppressionInterval == nil {
		return 0.0, symbolCallError("CGEventSourceGetLocalEventsSuppressionInterval", "10.4", _cGEventSourceGetLocalEventsSuppressionIntervalErr)
	}
	return _cGEventSourceGetLocalEventsSuppressionInterval(source), nil
}

// CGEventSourceGetLocalEventsSuppressionInterval returns the interval that local hardware events may be suppressed following the posting of a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/localEventsSuppressionInterval
func CGEventSourceGetLocalEventsSuppressionInterval(source CGEventSourceRef) float64 {
	result, callErr := tryCGEventSourceGetLocalEventsSuppressionInterval(source)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventSourceGetPixelsPerLine func(source CGEventSourceRef) float64
var _cGEventSourceGetPixelsPerLineErr error

func tryCGEventSourceGetPixelsPerLine(source CGEventSourceRef) (float64, error) {
	if _cGEventSourceGetPixelsPerLine == nil {
		return 0.0, symbolCallError("CGEventSourceGetPixelsPerLine", "10.5", _cGEventSourceGetPixelsPerLineErr)
	}
	return _cGEventSourceGetPixelsPerLine(source), nil
}

// CGEventSourceGetPixelsPerLine gets the scale of pixels per line in a scrolling event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/pixelsPerLine
func CGEventSourceGetPixelsPerLine(source CGEventSourceRef) float64 {
	result, callErr := tryCGEventSourceGetPixelsPerLine(source)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventSourceGetSourceStateID func(source CGEventSourceRef) CGEventSourceStateID
var _cGEventSourceGetSourceStateIDErr error

func tryCGEventSourceGetSourceStateID(source CGEventSourceRef) (CGEventSourceStateID, error) {
	if _cGEventSourceGetSourceStateID == nil {
		return *new(CGEventSourceStateID), symbolCallError("CGEventSourceGetSourceStateID", "10.4", _cGEventSourceGetSourceStateIDErr)
	}
	return _cGEventSourceGetSourceStateID(source), nil
}

// CGEventSourceGetSourceStateID returns the source state associated with a Quartz event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/sourceStateID
func CGEventSourceGetSourceStateID(source CGEventSourceRef) CGEventSourceStateID {
	result, callErr := tryCGEventSourceGetSourceStateID(source)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventSourceGetTypeID func() uint
var _cGEventSourceGetTypeIDErr error

func tryCGEventSourceGetTypeID() (uint, error) {
	if _cGEventSourceGetTypeID == nil {
		return 0, symbolCallError("CGEventSourceGetTypeID", "10.4", _cGEventSourceGetTypeIDErr)
	}
	return _cGEventSourceGetTypeID(), nil
}

// CGEventSourceGetTypeID returns the type identifier for the opaque type [CGEventSourceRef].
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/typeID
func CGEventSourceGetTypeID() uint {
	result, callErr := tryCGEventSourceGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventSourceGetUserData func(source CGEventSourceRef) int64
var _cGEventSourceGetUserDataErr error

func tryCGEventSourceGetUserData(source CGEventSourceRef) (int64, error) {
	if _cGEventSourceGetUserData == nil {
		return 0, symbolCallError("CGEventSourceGetUserData", "10.4", _cGEventSourceGetUserDataErr)
	}
	return _cGEventSourceGetUserData(source), nil
}

// CGEventSourceGetUserData returns the 64-bit user-specified data for a Quartz event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/userData
func CGEventSourceGetUserData(source CGEventSourceRef) int64 {
	result, callErr := tryCGEventSourceGetUserData(source)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventSourceKeyState func(stateID CGEventSourceStateID, key uint16) bool
var _cGEventSourceKeyStateErr error

func tryCGEventSourceKeyState(stateID CGEventSourceStateID, key uint16) (bool, error) {
	if _cGEventSourceKeyState == nil {
		return false, symbolCallError("CGEventSourceKeyState", "10.4", _cGEventSourceKeyStateErr)
	}
	return _cGEventSourceKeyState(stateID, key), nil
}

// CGEventSourceKeyState returns a Boolean value indicating the current keyboard state of a Quartz event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/keyState(_:key:)
func CGEventSourceKeyState(stateID CGEventSourceStateID, key uint16) bool {
	result, callErr := tryCGEventSourceKeyState(stateID, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventSourceSecondsSinceLastEventType func(stateID CGEventSourceStateID, eventType CGEventType) float64
var _cGEventSourceSecondsSinceLastEventTypeErr error

func tryCGEventSourceSecondsSinceLastEventType(stateID CGEventSourceStateID, eventType CGEventType) (float64, error) {
	if _cGEventSourceSecondsSinceLastEventType == nil {
		return 0.0, symbolCallError("CGEventSourceSecondsSinceLastEventType", "10.4", _cGEventSourceSecondsSinceLastEventTypeErr)
	}
	return _cGEventSourceSecondsSinceLastEventType(stateID, eventType), nil
}

// CGEventSourceSecondsSinceLastEventType returns the elapsed time since the last event for a Quartz event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/secondsSinceLastEventType(_:eventType:)
func CGEventSourceSecondsSinceLastEventType(stateID CGEventSourceStateID, eventType CGEventType) float64 {
	result, callErr := tryCGEventSourceSecondsSinceLastEventType(stateID, eventType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventSourceSetKeyboardType func(source CGEventSourceRef, keyboardType CGEventSourceKeyboardType)
var _cGEventSourceSetKeyboardTypeErr error

func tryCGEventSourceSetKeyboardType(source CGEventSourceRef, keyboardType CGEventSourceKeyboardType) error {
	if _cGEventSourceSetKeyboardType == nil {
		return symbolCallError("CGEventSourceSetKeyboardType", "10.4", _cGEventSourceSetKeyboardTypeErr)
	}
	_cGEventSourceSetKeyboardType(source, keyboardType)
	return nil
}

// CGEventSourceSetKeyboardType sets the keyboard type to be used with a Quartz event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceSetKeyboardType
func CGEventSourceSetKeyboardType(source CGEventSourceRef, keyboardType CGEventSourceKeyboardType) {
	if callErr := tryCGEventSourceSetKeyboardType(source, keyboardType); callErr != nil {
		panic(callErr)
	}
}

var _cGEventSourceSetLocalEventsFilterDuringSuppressionState func(source CGEventSourceRef, filter CGEventFilterMask, state CGEventSuppressionState)
var _cGEventSourceSetLocalEventsFilterDuringSuppressionStateErr error

func tryCGEventSourceSetLocalEventsFilterDuringSuppressionState(source CGEventSourceRef, filter CGEventFilterMask, state CGEventSuppressionState) error {
	if _cGEventSourceSetLocalEventsFilterDuringSuppressionState == nil {
		return symbolCallError("CGEventSourceSetLocalEventsFilterDuringSuppressionState", "10.4", _cGEventSourceSetLocalEventsFilterDuringSuppressionStateErr)
	}
	_cGEventSourceSetLocalEventsFilterDuringSuppressionState(source, filter, state)
	return nil
}

// CGEventSourceSetLocalEventsFilterDuringSuppressionState sets the mask that indicates which classes of local hardware events are enabled during event suppression.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSource/setLocalEventsFilterDuringSuppressionState(_:state:)
func CGEventSourceSetLocalEventsFilterDuringSuppressionState(source CGEventSourceRef, filter CGEventFilterMask, state CGEventSuppressionState) {
	if callErr := tryCGEventSourceSetLocalEventsFilterDuringSuppressionState(source, filter, state); callErr != nil {
		panic(callErr)
	}
}

var _cGEventSourceSetLocalEventsSuppressionInterval func(source CGEventSourceRef, seconds float64)
var _cGEventSourceSetLocalEventsSuppressionIntervalErr error

func tryCGEventSourceSetLocalEventsSuppressionInterval(source CGEventSourceRef, seconds float64) error {
	if _cGEventSourceSetLocalEventsSuppressionInterval == nil {
		return symbolCallError("CGEventSourceSetLocalEventsSuppressionInterval", "10.4", _cGEventSourceSetLocalEventsSuppressionIntervalErr)
	}
	_cGEventSourceSetLocalEventsSuppressionInterval(source, seconds)
	return nil
}

// CGEventSourceSetLocalEventsSuppressionInterval sets the interval that local hardware events may be suppressed following the posting of a Quartz event.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceSetLocalEventsSuppressionInterval
func CGEventSourceSetLocalEventsSuppressionInterval(source CGEventSourceRef, seconds float64) {
	if callErr := tryCGEventSourceSetLocalEventsSuppressionInterval(source, seconds); callErr != nil {
		panic(callErr)
	}
}

var _cGEventSourceSetPixelsPerLine func(source CGEventSourceRef, pixelsPerLine float64)
var _cGEventSourceSetPixelsPerLineErr error

func tryCGEventSourceSetPixelsPerLine(source CGEventSourceRef, pixelsPerLine float64) error {
	if _cGEventSourceSetPixelsPerLine == nil {
		return symbolCallError("CGEventSourceSetPixelsPerLine", "10.5", _cGEventSourceSetPixelsPerLineErr)
	}
	_cGEventSourceSetPixelsPerLine(source, pixelsPerLine)
	return nil
}

// CGEventSourceSetPixelsPerLine sets the scale of pixels per line in a scrolling event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceSetPixelsPerLine
func CGEventSourceSetPixelsPerLine(source CGEventSourceRef, pixelsPerLine float64) {
	if callErr := tryCGEventSourceSetPixelsPerLine(source, pixelsPerLine); callErr != nil {
		panic(callErr)
	}
}

var _cGEventSourceSetUserData func(source CGEventSourceRef, userData int64)
var _cGEventSourceSetUserDataErr error

func tryCGEventSourceSetUserData(source CGEventSourceRef, userData int64) error {
	if _cGEventSourceSetUserData == nil {
		return symbolCallError("CGEventSourceSetUserData", "10.4", _cGEventSourceSetUserDataErr)
	}
	_cGEventSourceSetUserData(source, userData)
	return nil
}

// CGEventSourceSetUserData sets the 64-bit user-specified data for a Quartz event source.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceSetUserData
func CGEventSourceSetUserData(source CGEventSourceRef, userData int64) {
	if callErr := tryCGEventSourceSetUserData(source, userData); callErr != nil {
		panic(callErr)
	}
}

var _cGEventTapCreate func(tap CGEventTapLocation, place CGEventTapPlacement, options CGEventTapOptions, eventsOfInterest CGEventMask, callback CGEventTapCallBack, userInfo unsafe.Pointer) corefoundation.CFMachPort
var _cGEventTapCreateErr error

func tryCGEventTapCreate(tap CGEventTapLocation, place CGEventTapPlacement, options CGEventTapOptions, eventsOfInterest CGEventMask, callback CGEventTapCallBack, userInfo unsafe.Pointer) (corefoundation.CFMachPort, error) {
	if _cGEventTapCreate == nil {
		return *new(corefoundation.CFMachPort), symbolCallError("CGEventTapCreate", "10.4", _cGEventTapCreateErr)
	}
	return _cGEventTapCreate(tap, place, options, eventsOfInterest, callback, userInfo), nil
}

// CGEventTapCreate creates an event tap.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/tapCreate(tap:place:options:eventsOfInterest:callback:userInfo:)
func CGEventTapCreate(tap CGEventTapLocation, place CGEventTapPlacement, options CGEventTapOptions, eventsOfInterest CGEventMask, callback CGEventTapCallBack, userInfo unsafe.Pointer) corefoundation.CFMachPort {
	result, callErr := tryCGEventTapCreate(tap, place, options, eventsOfInterest, callback, userInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventTapCreateForPSN func(processSerialNumber unsafe.Pointer, place CGEventTapPlacement, options CGEventTapOptions, eventsOfInterest CGEventMask, callback CGEventTapCallBack, userInfo unsafe.Pointer) corefoundation.CFMachPort
var _cGEventTapCreateForPSNErr error

func tryCGEventTapCreateForPSN(processSerialNumber unsafe.Pointer, place CGEventTapPlacement, options CGEventTapOptions, eventsOfInterest CGEventMask, callback CGEventTapCallBack, userInfo unsafe.Pointer) (corefoundation.CFMachPort, error) {
	if _cGEventTapCreateForPSN == nil {
		return *new(corefoundation.CFMachPort), symbolCallError("CGEventTapCreateForPSN", "10.4", _cGEventTapCreateForPSNErr)
	}
	return _cGEventTapCreateForPSN(processSerialNumber, place, options, eventsOfInterest, callback, userInfo), nil
}

// CGEventTapCreateForPSN creates an event tap for a specified process.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/tapCreateForPSN(processSerialNumber:place:options:eventsOfInterest:callback:userInfo:)
func CGEventTapCreateForPSN(processSerialNumber unsafe.Pointer, place CGEventTapPlacement, options CGEventTapOptions, eventsOfInterest CGEventMask, callback CGEventTapCallBack, userInfo unsafe.Pointer) corefoundation.CFMachPort {
	result, callErr := tryCGEventTapCreateForPSN(processSerialNumber, place, options, eventsOfInterest, callback, userInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventTapCreateForPid func(pid int32, place CGEventTapPlacement, options CGEventTapOptions, eventsOfInterest CGEventMask, callback CGEventTapCallBack, userInfo unsafe.Pointer) corefoundation.CFMachPort
var _cGEventTapCreateForPidErr error

func tryCGEventTapCreateForPid(pid int32, place CGEventTapPlacement, options CGEventTapOptions, eventsOfInterest CGEventMask, callback CGEventTapCallBack, userInfo unsafe.Pointer) (corefoundation.CFMachPort, error) {
	if _cGEventTapCreateForPid == nil {
		return *new(corefoundation.CFMachPort), symbolCallError("CGEventTapCreateForPid", "10.11", _cGEventTapCreateForPidErr)
	}
	return _cGEventTapCreateForPid(pid, place, options, eventsOfInterest, callback, userInfo), nil
}

// CGEventTapCreateForPid.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/tapCreateForPid(pid:place:options:eventsOfInterest:callback:userInfo:)
func CGEventTapCreateForPid(pid int32, place CGEventTapPlacement, options CGEventTapOptions, eventsOfInterest CGEventMask, callback CGEventTapCallBack, userInfo unsafe.Pointer) corefoundation.CFMachPort {
	result, callErr := tryCGEventTapCreateForPid(pid, place, options, eventsOfInterest, callback, userInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventTapEnable func(tap corefoundation.CFMachPort, enable bool)
var _cGEventTapEnableErr error

func tryCGEventTapEnable(tap corefoundation.CFMachPort, enable bool) error {
	if _cGEventTapEnable == nil {
		return symbolCallError("CGEventTapEnable", "10.4", _cGEventTapEnableErr)
	}
	_cGEventTapEnable(tap, enable)
	return nil
}

// CGEventTapEnable enables or disables an event tap.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/tapEnable(tap:enable:)
func CGEventTapEnable(tap corefoundation.CFMachPort, enable bool) {
	if callErr := tryCGEventTapEnable(tap, enable); callErr != nil {
		panic(callErr)
	}
}

var _cGEventTapIsEnabled func(tap corefoundation.CFMachPort) bool
var _cGEventTapIsEnabledErr error

func tryCGEventTapIsEnabled(tap corefoundation.CFMachPort) (bool, error) {
	if _cGEventTapIsEnabled == nil {
		return false, symbolCallError("CGEventTapIsEnabled", "10.4", _cGEventTapIsEnabledErr)
	}
	return _cGEventTapIsEnabled(tap), nil
}

// CGEventTapIsEnabled returns a Boolean value indicating whether an event tap is enabled.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/tapIsEnabled(tap:)
func CGEventTapIsEnabled(tap corefoundation.CFMachPort) bool {
	result, callErr := tryCGEventTapIsEnabled(tap)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGEventTapPostEvent func(proxy CGEventTapProxy, event CGEventRef)
var _cGEventTapPostEventErr error

func tryCGEventTapPostEvent(proxy CGEventTapProxy, event CGEventRef) error {
	if _cGEventTapPostEvent == nil {
		return symbolCallError("CGEventTapPostEvent", "10.4", _cGEventTapPostEventErr)
	}
	_cGEventTapPostEvent(proxy, event)
	return nil
}

// CGEventTapPostEvent posts a Quartz event from an event tap into the event stream.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGEvent/tapPostEvent(_:)
func CGEventTapPostEvent(proxy CGEventTapProxy, event CGEventRef) {
	if callErr := tryCGEventTapPostEvent(proxy, event); callErr != nil {
		panic(callErr)
	}
}

var _cGFontCanCreatePostScriptSubset func(font CGFontRef, format CGFontPostScriptFormat) bool
var _cGFontCanCreatePostScriptSubsetErr error

func tryCGFontCanCreatePostScriptSubset(font CGFontRef, format CGFontPostScriptFormat) (bool, error) {
	if _cGFontCanCreatePostScriptSubset == nil {
		return false, symbolCallError("CGFontCanCreatePostScriptSubset", "10.4", _cGFontCanCreatePostScriptSubsetErr)
	}
	return _cGFontCanCreatePostScriptSubset(font, format), nil
}

// CGFontCanCreatePostScriptSubset determines whether Core Graphics can create a subset of the font in PostScript format.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/canCreatePostScriptSubset(_:)
func CGFontCanCreatePostScriptSubset(font CGFontRef, format CGFontPostScriptFormat) bool {
	result, callErr := tryCGFontCanCreatePostScriptSubset(font, format)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontCopyFullName func(font CGFontRef) corefoundation.CFStringRef
var _cGFontCopyFullNameErr error

func tryCGFontCopyFullName(font CGFontRef) (corefoundation.CFStringRef, error) {
	if _cGFontCopyFullName == nil {
		return 0, symbolCallError("CGFontCopyFullName", "10.5", _cGFontCopyFullNameErr)
	}
	return _cGFontCopyFullName(font), nil
}

// CGFontCopyFullName returns the full name associated with a font object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/fullName
func CGFontCopyFullName(font CGFontRef) corefoundation.CFStringRef {
	result, callErr := tryCGFontCopyFullName(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontCopyGlyphNameForGlyph func(font CGFontRef, glyph unsafe.Pointer) corefoundation.CFStringRef
var _cGFontCopyGlyphNameForGlyphErr error

func tryCGFontCopyGlyphNameForGlyph(font CGFontRef, glyph unsafe.Pointer) (corefoundation.CFStringRef, error) {
	if _cGFontCopyGlyphNameForGlyph == nil {
		return 0, symbolCallError("CGFontCopyGlyphNameForGlyph", "10.5", _cGFontCopyGlyphNameForGlyphErr)
	}
	return _cGFontCopyGlyphNameForGlyph(font, glyph), nil
}

// CGFontCopyGlyphNameForGlyph returns the glyph name of the specified glyph in the specified font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/name(for:)
func CGFontCopyGlyphNameForGlyph(font CGFontRef, glyph unsafe.Pointer) corefoundation.CFStringRef {
	result, callErr := tryCGFontCopyGlyphNameForGlyph(font, glyph)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontCopyPostScriptName func(font CGFontRef) corefoundation.CFStringRef
var _cGFontCopyPostScriptNameErr error

func tryCGFontCopyPostScriptName(font CGFontRef) (corefoundation.CFStringRef, error) {
	if _cGFontCopyPostScriptName == nil {
		return 0, symbolCallError("CGFontCopyPostScriptName", "10.4", _cGFontCopyPostScriptNameErr)
	}
	return _cGFontCopyPostScriptName(font), nil
}

// CGFontCopyPostScriptName obtains the PostScript name of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/postScriptName
func CGFontCopyPostScriptName(font CGFontRef) corefoundation.CFStringRef {
	result, callErr := tryCGFontCopyPostScriptName(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontCopyTableForTag func(font CGFontRef, tag uint32) corefoundation.CFDataRef
var _cGFontCopyTableForTagErr error

func tryCGFontCopyTableForTag(font CGFontRef, tag uint32) (corefoundation.CFDataRef, error) {
	if _cGFontCopyTableForTag == nil {
		return 0, symbolCallError("CGFontCopyTableForTag", "10.5", _cGFontCopyTableForTagErr)
	}
	return _cGFontCopyTableForTag(font, tag), nil
}

// CGFontCopyTableForTag returns the font table that corresponds to the provided tag.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/table(for:)
func CGFontCopyTableForTag(font CGFontRef, tag uint32) corefoundation.CFDataRef {
	result, callErr := tryCGFontCopyTableForTag(font, tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontCopyTableTags func(font CGFontRef) corefoundation.CFArrayRef
var _cGFontCopyTableTagsErr error

func tryCGFontCopyTableTags(font CGFontRef) (corefoundation.CFArrayRef, error) {
	if _cGFontCopyTableTags == nil {
		return 0, symbolCallError("CGFontCopyTableTags", "10.5", _cGFontCopyTableTagsErr)
	}
	return _cGFontCopyTableTags(font), nil
}

// CGFontCopyTableTags returns an array of tags that correspond to the font tables for a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/tableTags
func CGFontCopyTableTags(font CGFontRef) corefoundation.CFArrayRef {
	result, callErr := tryCGFontCopyTableTags(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontCopyVariationAxes func(font CGFontRef) corefoundation.CFArrayRef
var _cGFontCopyVariationAxesErr error

func tryCGFontCopyVariationAxes(font CGFontRef) (corefoundation.CFArrayRef, error) {
	if _cGFontCopyVariationAxes == nil {
		return 0, symbolCallError("CGFontCopyVariationAxes", "10.4", _cGFontCopyVariationAxesErr)
	}
	return _cGFontCopyVariationAxes(font), nil
}

// CGFontCopyVariationAxes returns an array of the variation axis dictionaries for a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/variationAxes
func CGFontCopyVariationAxes(font CGFontRef) corefoundation.CFArrayRef {
	result, callErr := tryCGFontCopyVariationAxes(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontCopyVariations func(font CGFontRef) corefoundation.CFDictionaryRef
var _cGFontCopyVariationsErr error

func tryCGFontCopyVariations(font CGFontRef) (corefoundation.CFDictionaryRef, error) {
	if _cGFontCopyVariations == nil {
		return 0, symbolCallError("CGFontCopyVariations", "10.4", _cGFontCopyVariationsErr)
	}
	return _cGFontCopyVariations(font), nil
}

// CGFontCopyVariations returns the variation specification dictionary for a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/variations
func CGFontCopyVariations(font CGFontRef) corefoundation.CFDictionaryRef {
	result, callErr := tryCGFontCopyVariations(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontCreateCopyWithVariations func(font CGFontRef, variations corefoundation.CFDictionaryRef) CGFontRef
var _cGFontCreateCopyWithVariationsErr error

func tryCGFontCreateCopyWithVariations(font CGFontRef, variations corefoundation.CFDictionaryRef) (CGFontRef, error) {
	if _cGFontCreateCopyWithVariations == nil {
		return 0, symbolCallError("CGFontCreateCopyWithVariations", "10.4", _cGFontCreateCopyWithVariationsErr)
	}
	return _cGFontCreateCopyWithVariations(font, variations), nil
}

// CGFontCreateCopyWithVariations creates a copy of a font using a variation specification dictionary.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/copy(withVariations:)
func CGFontCreateCopyWithVariations(font CGFontRef, variations corefoundation.CFDictionaryRef) CGFontRef {
	result, callErr := tryCGFontCreateCopyWithVariations(font, variations)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontCreatePostScriptEncoding func(font CGFontRef, encoding unsafe.Pointer) corefoundation.CFDataRef
var _cGFontCreatePostScriptEncodingErr error

func tryCGFontCreatePostScriptEncoding(font CGFontRef, encoding unsafe.Pointer) (corefoundation.CFDataRef, error) {
	if _cGFontCreatePostScriptEncoding == nil {
		return 0, symbolCallError("CGFontCreatePostScriptEncoding", "10.4", _cGFontCreatePostScriptEncodingErr)
	}
	return _cGFontCreatePostScriptEncoding(font, encoding), nil
}

// CGFontCreatePostScriptEncoding creates a PostScript encoding of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/createPostScriptEncoding(encoding:)
func CGFontCreatePostScriptEncoding(font CGFontRef, encoding unsafe.Pointer) corefoundation.CFDataRef {
	result, callErr := tryCGFontCreatePostScriptEncoding(font, encoding)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontCreatePostScriptSubset func(font CGFontRef, subsetName corefoundation.CFStringRef, format CGFontPostScriptFormat, glyphs unsafe.Pointer, count uintptr, encoding unsafe.Pointer) corefoundation.CFDataRef
var _cGFontCreatePostScriptSubsetErr error

func tryCGFontCreatePostScriptSubset(font CGFontRef, subsetName corefoundation.CFStringRef, format CGFontPostScriptFormat, glyphs unsafe.Pointer, count uintptr, encoding unsafe.Pointer) (corefoundation.CFDataRef, error) {
	if _cGFontCreatePostScriptSubset == nil {
		return 0, symbolCallError("CGFontCreatePostScriptSubset", "10.4", _cGFontCreatePostScriptSubsetErr)
	}
	return _cGFontCreatePostScriptSubset(font, subsetName, format, glyphs, count, encoding), nil
}

// CGFontCreatePostScriptSubset creates a subset of the font in the specified PostScript format.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/createPostScriptSubset(subsetName:format:glyphs:count:encoding:)
func CGFontCreatePostScriptSubset(font CGFontRef, subsetName corefoundation.CFStringRef, format CGFontPostScriptFormat, glyphs unsafe.Pointer, count uintptr, encoding unsafe.Pointer) corefoundation.CFDataRef {
	result, callErr := tryCGFontCreatePostScriptSubset(font, subsetName, format, glyphs, count, encoding)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontCreateWithDataProvider func(provider CGDataProviderRef) CGFontRef
var _cGFontCreateWithDataProviderErr error

func tryCGFontCreateWithDataProvider(provider CGDataProviderRef) (CGFontRef, error) {
	if _cGFontCreateWithDataProvider == nil {
		return 0, symbolCallError("CGFontCreateWithDataProvider", "10.5", _cGFontCreateWithDataProviderErr)
	}
	return _cGFontCreateWithDataProvider(provider), nil
}

// CGFontCreateWithDataProvider creates a font object from data supplied from a data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/init(_:)-9aour
func CGFontCreateWithDataProvider(provider CGDataProviderRef) CGFontRef {
	result, callErr := tryCGFontCreateWithDataProvider(provider)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontCreateWithFontName func(name corefoundation.CFStringRef) CGFontRef
var _cGFontCreateWithFontNameErr error

func tryCGFontCreateWithFontName(name corefoundation.CFStringRef) (CGFontRef, error) {
	if _cGFontCreateWithFontName == nil {
		return 0, symbolCallError("CGFontCreateWithFontName", "10.5", _cGFontCreateWithFontNameErr)
	}
	return _cGFontCreateWithFontName(name), nil
}

// CGFontCreateWithFontName creates a font object corresponding to the font specified by a PostScript or full name.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/init(_:)-1p4b
func CGFontCreateWithFontName(name corefoundation.CFStringRef) CGFontRef {
	result, callErr := tryCGFontCreateWithFontName(name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontGetAscent func(font CGFontRef) int
var _cGFontGetAscentErr error

func tryCGFontGetAscent(font CGFontRef) (int, error) {
	if _cGFontGetAscent == nil {
		return 0, symbolCallError("CGFontGetAscent", "10.5", _cGFontGetAscentErr)
	}
	return _cGFontGetAscent(font), nil
}

// CGFontGetAscent returns the ascent of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/ascent
func CGFontGetAscent(font CGFontRef) int {
	result, callErr := tryCGFontGetAscent(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontGetCapHeight func(font CGFontRef) int
var _cGFontGetCapHeightErr error

func tryCGFontGetCapHeight(font CGFontRef) (int, error) {
	if _cGFontGetCapHeight == nil {
		return 0, symbolCallError("CGFontGetCapHeight", "10.5", _cGFontGetCapHeightErr)
	}
	return _cGFontGetCapHeight(font), nil
}

// CGFontGetCapHeight returns the cap height of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/capHeight
func CGFontGetCapHeight(font CGFontRef) int {
	result, callErr := tryCGFontGetCapHeight(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontGetDescent func(font CGFontRef) int
var _cGFontGetDescentErr error

func tryCGFontGetDescent(font CGFontRef) (int, error) {
	if _cGFontGetDescent == nil {
		return 0, symbolCallError("CGFontGetDescent", "10.5", _cGFontGetDescentErr)
	}
	return _cGFontGetDescent(font), nil
}

// CGFontGetDescent returns the descent of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/descent
func CGFontGetDescent(font CGFontRef) int {
	result, callErr := tryCGFontGetDescent(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontGetFontBBox func(font CGFontRef) corefoundation.CGRect
var _cGFontGetFontBBoxErr error

func tryCGFontGetFontBBox(font CGFontRef) (corefoundation.CGRect, error) {
	if _cGFontGetFontBBox == nil {
		return corefoundation.CGRect{}, symbolCallError("CGFontGetFontBBox", "10.5", _cGFontGetFontBBoxErr)
	}
	return _cGFontGetFontBBox(font), nil
}

// CGFontGetFontBBox returns the bounding box of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/fontBBox
func CGFontGetFontBBox(font CGFontRef) corefoundation.CGRect {
	result, callErr := tryCGFontGetFontBBox(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontGetGlyphAdvances func(font CGFontRef, glyphs unsafe.Pointer, count uintptr, advances *int) bool
var _cGFontGetGlyphAdvancesErr error

func tryCGFontGetGlyphAdvances(font CGFontRef, glyphs unsafe.Pointer, count uintptr, advances []int) (bool, error) {
	if _cGFontGetGlyphAdvances == nil {
		return false, symbolCallError("CGFontGetGlyphAdvances", "10.0", _cGFontGetGlyphAdvancesErr)
	}
	return _cGFontGetGlyphAdvances(font, glyphs, count, unsafe.SliceData(advances)), nil
}

// CGFontGetGlyphAdvances gets the advance width of each glyph in the provided array.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/getGlyphAdvances(glyphs:count:advances:)
func CGFontGetGlyphAdvances(font CGFontRef, glyphs unsafe.Pointer, count uintptr, advances []int) bool {
	result, callErr := tryCGFontGetGlyphAdvances(font, glyphs, count, advances)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontGetGlyphBBoxes func(font CGFontRef, glyphs unsafe.Pointer, count uintptr, bboxes *corefoundation.CGRect) bool
var _cGFontGetGlyphBBoxesErr error

func tryCGFontGetGlyphBBoxes(font CGFontRef, glyphs unsafe.Pointer, count uintptr, bboxes *corefoundation.CGRect) (bool, error) {
	if _cGFontGetGlyphBBoxes == nil {
		return false, symbolCallError("CGFontGetGlyphBBoxes", "10.5", _cGFontGetGlyphBBoxesErr)
	}
	return _cGFontGetGlyphBBoxes(font, glyphs, count, bboxes), nil
}

// CGFontGetGlyphBBoxes get the bounding box of each glyph in an array.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/getGlyphBBoxes(glyphs:count:bboxes:)
func CGFontGetGlyphBBoxes(font CGFontRef, glyphs unsafe.Pointer, count uintptr, bboxes *corefoundation.CGRect) bool {
	result, callErr := tryCGFontGetGlyphBBoxes(font, glyphs, count, bboxes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontGetGlyphWithGlyphName func(font CGFontRef, name corefoundation.CFStringRef) unsafe.Pointer
var _cGFontGetGlyphWithGlyphNameErr error

func tryCGFontGetGlyphWithGlyphName(font CGFontRef, name corefoundation.CFStringRef) (unsafe.Pointer, error) {
	if _cGFontGetGlyphWithGlyphName == nil {
		return nil, symbolCallError("CGFontGetGlyphWithGlyphName", "10.5", _cGFontGetGlyphWithGlyphNameErr)
	}
	return _cGFontGetGlyphWithGlyphName(font, name), nil
}

// CGFontGetGlyphWithGlyphName returns the glyph for the glyph name associated with the specified font object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/getGlyphWithGlyphName(name:)
func CGFontGetGlyphWithGlyphName(font CGFontRef, name corefoundation.CFStringRef) unsafe.Pointer {
	result, callErr := tryCGFontGetGlyphWithGlyphName(font, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontGetItalicAngle func(font CGFontRef) float64
var _cGFontGetItalicAngleErr error

func tryCGFontGetItalicAngle(font CGFontRef) (float64, error) {
	if _cGFontGetItalicAngle == nil {
		return 0.0, symbolCallError("CGFontGetItalicAngle", "10.5", _cGFontGetItalicAngleErr)
	}
	return _cGFontGetItalicAngle(font), nil
}

// CGFontGetItalicAngle returns the italic angle of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/italicAngle
func CGFontGetItalicAngle(font CGFontRef) float64 {
	result, callErr := tryCGFontGetItalicAngle(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontGetLeading func(font CGFontRef) int
var _cGFontGetLeadingErr error

func tryCGFontGetLeading(font CGFontRef) (int, error) {
	if _cGFontGetLeading == nil {
		return 0, symbolCallError("CGFontGetLeading", "10.5", _cGFontGetLeadingErr)
	}
	return _cGFontGetLeading(font), nil
}

// CGFontGetLeading returns the leading of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/leading
func CGFontGetLeading(font CGFontRef) int {
	result, callErr := tryCGFontGetLeading(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontGetNumberOfGlyphs func(font CGFontRef) uintptr
var _cGFontGetNumberOfGlyphsErr error

func tryCGFontGetNumberOfGlyphs(font CGFontRef) (uintptr, error) {
	if _cGFontGetNumberOfGlyphs == nil {
		return 0, symbolCallError("CGFontGetNumberOfGlyphs", "10.0", _cGFontGetNumberOfGlyphsErr)
	}
	return _cGFontGetNumberOfGlyphs(font), nil
}

// CGFontGetNumberOfGlyphs returns the number of glyphs in a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/numberOfGlyphs
func CGFontGetNumberOfGlyphs(font CGFontRef) uintptr {
	result, callErr := tryCGFontGetNumberOfGlyphs(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontGetStemV func(font CGFontRef) float64
var _cGFontGetStemVErr error

func tryCGFontGetStemV(font CGFontRef) (float64, error) {
	if _cGFontGetStemV == nil {
		return 0.0, symbolCallError("CGFontGetStemV", "10.5", _cGFontGetStemVErr)
	}
	return _cGFontGetStemV(font), nil
}

// CGFontGetStemV returns the thickness of the dominant vertical stems of glyphs in a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/stemV
func CGFontGetStemV(font CGFontRef) float64 {
	result, callErr := tryCGFontGetStemV(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontGetTypeID func() uint
var _cGFontGetTypeIDErr error

func tryCGFontGetTypeID() (uint, error) {
	if _cGFontGetTypeID == nil {
		return 0, symbolCallError("CGFontGetTypeID", "10.2", _cGFontGetTypeIDErr)
	}
	return _cGFontGetTypeID(), nil
}

// CGFontGetTypeID returns the Core Foundation type identifier for Core Graphics fonts.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/typeID
func CGFontGetTypeID() uint {
	result, callErr := tryCGFontGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontGetUnitsPerEm func(font CGFontRef) int
var _cGFontGetUnitsPerEmErr error

func tryCGFontGetUnitsPerEm(font CGFontRef) (int, error) {
	if _cGFontGetUnitsPerEm == nil {
		return 0, symbolCallError("CGFontGetUnitsPerEm", "10.0", _cGFontGetUnitsPerEmErr)
	}
	return _cGFontGetUnitsPerEm(font), nil
}

// CGFontGetUnitsPerEm returns the number of glyph space units per em for the provided font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/unitsPerEm
func CGFontGetUnitsPerEm(font CGFontRef) int {
	result, callErr := tryCGFontGetUnitsPerEm(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontGetXHeight func(font CGFontRef) int
var _cGFontGetXHeightErr error

func tryCGFontGetXHeight(font CGFontRef) (int, error) {
	if _cGFontGetXHeight == nil {
		return 0, symbolCallError("CGFontGetXHeight", "10.5", _cGFontGetXHeightErr)
	}
	return _cGFontGetXHeight(font), nil
}

// CGFontGetXHeight returns the x-height of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFont/xHeight
func CGFontGetXHeight(font CGFontRef) int {
	result, callErr := tryCGFontGetXHeight(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFontRelease func(font CGFontRef)
var _cGFontReleaseErr error

func tryCGFontRelease(font CGFontRef) error {
	if _cGFontRelease == nil {
		return symbolCallError("CGFontRelease", "10.0", _cGFontReleaseErr)
	}
	_cGFontRelease(font)
	return nil
}

// CGFontRelease decrements the retain count of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFontRelease
func CGFontRelease(font CGFontRef) {
	if callErr := tryCGFontRelease(font); callErr != nil {
		panic(callErr)
	}
}

var _cGFontRetain func(font CGFontRef) CGFontRef
var _cGFontRetainErr error

func tryCGFontRetain(font CGFontRef) (CGFontRef, error) {
	if _cGFontRetain == nil {
		return 0, symbolCallError("CGFontRetain", "10.0", _cGFontRetainErr)
	}
	return _cGFontRetain(font), nil
}

// CGFontRetain increments the retain count of a font.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFontRetain
func CGFontRetain(font CGFontRef) CGFontRef {
	result, callErr := tryCGFontRetain(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFunctionCreate func(info unsafe.Pointer, domainDimension uintptr, domain *float64, rangeDimension uintptr, range_ *float64, callbacks *CGFunctionCallbacks) CGFunctionRef
var _cGFunctionCreateErr error

func tryCGFunctionCreate(info unsafe.Pointer, domainDimension uintptr, domain *float64, rangeDimension uintptr, range_ *float64, callbacks *CGFunctionCallbacks) (CGFunctionRef, error) {
	if _cGFunctionCreate == nil {
		return 0, symbolCallError("CGFunctionCreate", "10.2", _cGFunctionCreateErr)
	}
	return _cGFunctionCreate(info, domainDimension, domain, rangeDimension, range_, callbacks), nil
}

// CGFunctionCreate creates a Core Graphics function.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFunction/init(info:domainDimension:domain:rangeDimension:range:callbacks:)
func CGFunctionCreate(info unsafe.Pointer, domainDimension uintptr, domain *float64, rangeDimension uintptr, range_ *float64, callbacks *CGFunctionCallbacks) CGFunctionRef {
	result, callErr := tryCGFunctionCreate(info, domainDimension, domain, rangeDimension, range_, callbacks)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFunctionGetTypeID func() uint
var _cGFunctionGetTypeIDErr error

func tryCGFunctionGetTypeID() (uint, error) {
	if _cGFunctionGetTypeID == nil {
		return 0, symbolCallError("CGFunctionGetTypeID", "10.2", _cGFunctionGetTypeIDErr)
	}
	return _cGFunctionGetTypeID(), nil
}

// CGFunctionGetTypeID returns the type identifier for Core Graphics function objects.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFunction/typeID
func CGFunctionGetTypeID() uint {
	result, callErr := tryCGFunctionGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGFunctionRelease func(function CGFunctionRef)
var _cGFunctionReleaseErr error

func tryCGFunctionRelease(function CGFunctionRef) error {
	if _cGFunctionRelease == nil {
		return symbolCallError("CGFunctionRelease", "10.2", _cGFunctionReleaseErr)
	}
	_cGFunctionRelease(function)
	return nil
}

// CGFunctionRelease decrements the retain count of a function object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFunctionRelease
func CGFunctionRelease(function CGFunctionRef) {
	if callErr := tryCGFunctionRelease(function); callErr != nil {
		panic(callErr)
	}
}

var _cGFunctionRetain func(function CGFunctionRef) CGFunctionRef
var _cGFunctionRetainErr error

func tryCGFunctionRetain(function CGFunctionRef) (CGFunctionRef, error) {
	if _cGFunctionRetain == nil {
		return 0, symbolCallError("CGFunctionRetain", "10.2", _cGFunctionRetainErr)
	}
	return _cGFunctionRetain(function), nil
}

// CGFunctionRetain increments the retain count of a function object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGFunctionRetain
func CGFunctionRetain(function CGFunctionRef) CGFunctionRef {
	result, callErr := tryCGFunctionRetain(function)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGGetActiveDisplayList func(maxDisplays uint32, activeDisplays *uint32, displayCount *uint32) CGError
var _cGGetActiveDisplayListErr error

func tryCGGetActiveDisplayList(maxDisplays uint32, activeDisplays *uint32, displayCount *uint32) (CGError, error) {
	if _cGGetActiveDisplayList == nil {
		return *new(CGError), symbolCallError("CGGetActiveDisplayList", "10.0", _cGGetActiveDisplayListErr)
	}
	return _cGGetActiveDisplayList(maxDisplays, activeDisplays, displayCount), nil
}

// CGGetActiveDisplayList provides a list of displays that are active for drawing.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGetActiveDisplayList(_:_:_:)
func CGGetActiveDisplayList(maxDisplays uint32, activeDisplays *uint32, displayCount *uint32) CGError {
	result, callErr := tryCGGetActiveDisplayList(maxDisplays, activeDisplays, displayCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGGetDisplayTransferByFormula func(display uint32, redMin *CGGammaValue, redMax *CGGammaValue, redGamma *CGGammaValue, greenMin *CGGammaValue, greenMax *CGGammaValue, greenGamma *CGGammaValue, blueMin *CGGammaValue, blueMax *CGGammaValue, blueGamma *CGGammaValue) CGError
var _cGGetDisplayTransferByFormulaErr error

func tryCGGetDisplayTransferByFormula(display uint32, redMin *CGGammaValue, redMax *CGGammaValue, redGamma *CGGammaValue, greenMin *CGGammaValue, greenMax *CGGammaValue, greenGamma *CGGammaValue, blueMin *CGGammaValue, blueMax *CGGammaValue, blueGamma *CGGammaValue) (CGError, error) {
	if _cGGetDisplayTransferByFormula == nil {
		return *new(CGError), symbolCallError("CGGetDisplayTransferByFormula", "10.0", _cGGetDisplayTransferByFormulaErr)
	}
	return _cGGetDisplayTransferByFormula(display, redMin, redMax, redGamma, greenMin, greenMax, greenGamma, blueMin, blueMax, blueGamma), nil
}

// CGGetDisplayTransferByFormula gets the coefficients of the gamma transfer formula for a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGetDisplayTransferByFormula(_:_:_:_:_:_:_:_:_:_:)
func CGGetDisplayTransferByFormula(display uint32, redMin *CGGammaValue, redMax *CGGammaValue, redGamma *CGGammaValue, greenMin *CGGammaValue, greenMax *CGGammaValue, greenGamma *CGGammaValue, blueMin *CGGammaValue, blueMax *CGGammaValue, blueGamma *CGGammaValue) CGError {
	result, callErr := tryCGGetDisplayTransferByFormula(display, redMin, redMax, redGamma, greenMin, greenMax, greenGamma, blueMin, blueMax, blueGamma)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGGetDisplayTransferByTable func(display uint32, capacity uint32, redTable *CGGammaValue, greenTable *CGGammaValue, blueTable *CGGammaValue, sampleCount *uint32) CGError
var _cGGetDisplayTransferByTableErr error

func tryCGGetDisplayTransferByTable(display uint32, capacity uint32, redTable *CGGammaValue, greenTable *CGGammaValue, blueTable *CGGammaValue, sampleCount *uint32) (CGError, error) {
	if _cGGetDisplayTransferByTable == nil {
		return *new(CGError), symbolCallError("CGGetDisplayTransferByTable", "10.0", _cGGetDisplayTransferByTableErr)
	}
	return _cGGetDisplayTransferByTable(display, capacity, redTable, greenTable, blueTable, sampleCount), nil
}

// CGGetDisplayTransferByTable gets the values in the RGB gamma tables for a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGetDisplayTransferByTable(_:_:_:_:_:_:)
func CGGetDisplayTransferByTable(display uint32, capacity uint32, redTable *CGGammaValue, greenTable *CGGammaValue, blueTable *CGGammaValue, sampleCount *uint32) CGError {
	result, callErr := tryCGGetDisplayTransferByTable(display, capacity, redTable, greenTable, blueTable, sampleCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGGetDisplaysWithOpenGLDisplayMask func(mask CGOpenGLDisplayMask, maxDisplays uint32, displays *uint32, matchingDisplayCount *uint32) CGError
var _cGGetDisplaysWithOpenGLDisplayMaskErr error

func tryCGGetDisplaysWithOpenGLDisplayMask(mask CGOpenGLDisplayMask, maxDisplays uint32, displays *uint32, matchingDisplayCount *uint32) (CGError, error) {
	if _cGGetDisplaysWithOpenGLDisplayMask == nil {
		return *new(CGError), symbolCallError("CGGetDisplaysWithOpenGLDisplayMask", "10.0", _cGGetDisplaysWithOpenGLDisplayMaskErr)
	}
	return _cGGetDisplaysWithOpenGLDisplayMask(mask, maxDisplays, displays, matchingDisplayCount), nil
}

// CGGetDisplaysWithOpenGLDisplayMask provides a list of displays that corresponds to the bits set in an OpenGL display mask.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGetDisplaysWithOpenGLDisplayMask(_:_:_:_:)
func CGGetDisplaysWithOpenGLDisplayMask(mask CGOpenGLDisplayMask, maxDisplays uint32, displays *uint32, matchingDisplayCount *uint32) CGError {
	result, callErr := tryCGGetDisplaysWithOpenGLDisplayMask(mask, maxDisplays, displays, matchingDisplayCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGGetDisplaysWithPoint func(point corefoundation.CGPoint, maxDisplays uint32, displays *uint32, matchingDisplayCount *uint32) CGError
var _cGGetDisplaysWithPointErr error

func tryCGGetDisplaysWithPoint(point corefoundation.CGPoint, maxDisplays uint32, displays *uint32, matchingDisplayCount *uint32) (CGError, error) {
	if _cGGetDisplaysWithPoint == nil {
		return *new(CGError), symbolCallError("CGGetDisplaysWithPoint", "10.0", _cGGetDisplaysWithPointErr)
	}
	return _cGGetDisplaysWithPoint(point, maxDisplays, displays, matchingDisplayCount), nil
}

// CGGetDisplaysWithPoint provides a list of online displays with bounds that include the specified point.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGetDisplaysWithPoint(_:_:_:_:)
func CGGetDisplaysWithPoint(point corefoundation.CGPoint, maxDisplays uint32, displays *uint32, matchingDisplayCount *uint32) CGError {
	result, callErr := tryCGGetDisplaysWithPoint(point, maxDisplays, displays, matchingDisplayCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGGetDisplaysWithRect func(rect corefoundation.CGRect, maxDisplays uint32, displays *uint32, matchingDisplayCount *uint32) CGError
var _cGGetDisplaysWithRectErr error

func tryCGGetDisplaysWithRect(rect corefoundation.CGRect, maxDisplays uint32, displays *uint32, matchingDisplayCount *uint32) (CGError, error) {
	if _cGGetDisplaysWithRect == nil {
		return *new(CGError), symbolCallError("CGGetDisplaysWithRect", "10.0", _cGGetDisplaysWithRectErr)
	}
	return _cGGetDisplaysWithRect(rect, maxDisplays, displays, matchingDisplayCount), nil
}

// CGGetDisplaysWithRect gets a list of online displays with bounds that intersect the specified rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGetDisplaysWithRect(_:_:_:_:)
func CGGetDisplaysWithRect(rect corefoundation.CGRect, maxDisplays uint32, displays *uint32, matchingDisplayCount *uint32) CGError {
	result, callErr := tryCGGetDisplaysWithRect(rect, maxDisplays, displays, matchingDisplayCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGGetEventTapList func(maxNumberOfTaps uint32, tapList *CGEventTapInformation, eventTapCount *uint32) CGError
var _cGGetEventTapListErr error

func tryCGGetEventTapList(maxNumberOfTaps uint32, tapList *CGEventTapInformation, eventTapCount *uint32) (CGError, error) {
	if _cGGetEventTapList == nil {
		return *new(CGError), symbolCallError("CGGetEventTapList", "10.4", _cGGetEventTapListErr)
	}
	return _cGGetEventTapList(maxNumberOfTaps, tapList, eventTapCount), nil
}

// CGGetEventTapList gets a list of currently installed event taps.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGetEventTapList(_:_:_:)
func CGGetEventTapList(maxNumberOfTaps uint32, tapList *CGEventTapInformation, eventTapCount *uint32) CGError {
	result, callErr := tryCGGetEventTapList(maxNumberOfTaps, tapList, eventTapCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGGetLastMouseDelta func(deltaX *int32, deltaY *int32)
var _cGGetLastMouseDeltaErr error

func tryCGGetLastMouseDelta(deltaX *int32, deltaY *int32) error {
	if _cGGetLastMouseDelta == nil {
		return symbolCallError("CGGetLastMouseDelta", "10.0", _cGGetLastMouseDeltaErr)
	}
	_cGGetLastMouseDelta(deltaX, deltaY)
	return nil
}

// CGGetLastMouseDelta reports the change in mouse position since the last mouse movement event received by the application.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGetLastMouseDelta
func CGGetLastMouseDelta(deltaX *int32, deltaY *int32) {
	if callErr := tryCGGetLastMouseDelta(deltaX, deltaY); callErr != nil {
		panic(callErr)
	}
}

var _cGGetOnlineDisplayList func(maxDisplays uint32, onlineDisplays *uint32, displayCount *uint32) CGError
var _cGGetOnlineDisplayListErr error

func tryCGGetOnlineDisplayList(maxDisplays uint32, onlineDisplays *uint32, displayCount *uint32) (CGError, error) {
	if _cGGetOnlineDisplayList == nil {
		return *new(CGError), symbolCallError("CGGetOnlineDisplayList", "10.2", _cGGetOnlineDisplayListErr)
	}
	return _cGGetOnlineDisplayList(maxDisplays, onlineDisplays, displayCount), nil
}

// CGGetOnlineDisplayList provides a list of displays that are online (active, mirrored, or sleeping).
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGetOnlineDisplayList(_:_:_:)
func CGGetOnlineDisplayList(maxDisplays uint32, onlineDisplays *uint32, displayCount *uint32) CGError {
	result, callErr := tryCGGetOnlineDisplayList(maxDisplays, onlineDisplays, displayCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGGradientCreateWithColorComponents func(space CGColorSpaceRef, components *float64, locations *float64, count uintptr) CGGradientRef
var _cGGradientCreateWithColorComponentsErr error

func tryCGGradientCreateWithColorComponents(space CGColorSpaceRef, components *float64, locations *float64, count uintptr) (CGGradientRef, error) {
	if _cGGradientCreateWithColorComponents == nil {
		return 0, symbolCallError("CGGradientCreateWithColorComponents", "10.5", _cGGradientCreateWithColorComponentsErr)
	}
	return _cGGradientCreateWithColorComponents(space, components, locations, count), nil
}

// CGGradientCreateWithColorComponents creates a CGGradient object from a color space and the provided color components and locations.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGradient/init(colorSpace:colorComponents:locations:count:)
func CGGradientCreateWithColorComponents(space CGColorSpaceRef, components *float64, locations *float64, count uintptr) CGGradientRef {
	result, callErr := tryCGGradientCreateWithColorComponents(space, components, locations, count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGGradientCreateWithColors func(space CGColorSpaceRef, colors corefoundation.CFArrayRef, locations *float64) CGGradientRef
var _cGGradientCreateWithColorsErr error

func tryCGGradientCreateWithColors(space CGColorSpaceRef, colors corefoundation.CFArrayRef, locations *float64) (CGGradientRef, error) {
	if _cGGradientCreateWithColors == nil {
		return 0, symbolCallError("CGGradientCreateWithColors", "10.5", _cGGradientCreateWithColorsErr)
	}
	return _cGGradientCreateWithColors(space, colors, locations), nil
}

// CGGradientCreateWithColors creates a gradient object from a color space and the provided color objects and locations.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGradient/init(colorsSpace:colors:locations:)
func CGGradientCreateWithColors(space CGColorSpaceRef, colors corefoundation.CFArrayRef, locations *float64) CGGradientRef {
	result, callErr := tryCGGradientCreateWithColors(space, colors, locations)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGGradientCreateWithContentHeadroom func(headroom float32, space CGColorSpaceRef, components *float64, locations *float64, count uintptr) CGGradientRef
var _cGGradientCreateWithContentHeadroomErr error

func tryCGGradientCreateWithContentHeadroom(headroom float32, space CGColorSpaceRef, components *float64, locations *float64, count uintptr) (CGGradientRef, error) {
	if _cGGradientCreateWithContentHeadroom == nil {
		return 0, symbolCallError("CGGradientCreateWithContentHeadroom", "26.0", _cGGradientCreateWithContentHeadroomErr)
	}
	return _cGGradientCreateWithContentHeadroom(headroom, space, components, locations, count), nil
}

// CGGradientCreateWithContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGradient/init(headroom:colorSpace:colorComponents:locations:count:)
func CGGradientCreateWithContentHeadroom(headroom float32, space CGColorSpaceRef, components *float64, locations *float64, count uintptr) CGGradientRef {
	result, callErr := tryCGGradientCreateWithContentHeadroom(headroom, space, components, locations, count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGGradientGetContentHeadroom func(gradient CGGradientRef) float32
var _cGGradientGetContentHeadroomErr error

func tryCGGradientGetContentHeadroom(gradient CGGradientRef) (float32, error) {
	if _cGGradientGetContentHeadroom == nil {
		return 0.0, symbolCallError("CGGradientGetContentHeadroom", "26.0", _cGGradientGetContentHeadroomErr)
	}
	return _cGGradientGetContentHeadroom(gradient), nil
}

// CGGradientGetContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGradient/contentHeadroom
func CGGradientGetContentHeadroom(gradient CGGradientRef) float32 {
	result, callErr := tryCGGradientGetContentHeadroom(gradient)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGGradientGetTypeID func() uint
var _cGGradientGetTypeIDErr error

func tryCGGradientGetTypeID() (uint, error) {
	if _cGGradientGetTypeID == nil {
		return 0, symbolCallError("CGGradientGetTypeID", "10.5", _cGGradientGetTypeIDErr)
	}
	return _cGGradientGetTypeID(), nil
}

// CGGradientGetTypeID returns the Core Foundation type identifier for CGGradient objects.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGradient/typeID
func CGGradientGetTypeID() uint {
	result, callErr := tryCGGradientGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGGradientRelease func(gradient CGGradientRef)
var _cGGradientReleaseErr error

func tryCGGradientRelease(gradient CGGradientRef) error {
	if _cGGradientRelease == nil {
		return symbolCallError("CGGradientRelease", "10.5", _cGGradientReleaseErr)
	}
	_cGGradientRelease(gradient)
	return nil
}

// CGGradientRelease decrements the retain count of a CGGradient object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGradientRelease
func CGGradientRelease(gradient CGGradientRef) {
	if callErr := tryCGGradientRelease(gradient); callErr != nil {
		panic(callErr)
	}
}

var _cGGradientRetain func(gradient CGGradientRef) CGGradientRef
var _cGGradientRetainErr error

func tryCGGradientRetain(gradient CGGradientRef) (CGGradientRef, error) {
	if _cGGradientRetain == nil {
		return 0, symbolCallError("CGGradientRetain", "10.5", _cGGradientRetainErr)
	}
	return _cGGradientRetain(gradient), nil
}

// CGGradientRetain increments the retain count of a CGGradient object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGGradientRetain
func CGGradientRetain(gradient CGGradientRef) CGGradientRef {
	result, callErr := tryCGGradientRetain(gradient)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageCalculateContentAverageLightLevel func(image CGImageRef) float32
var _cGImageCalculateContentAverageLightLevelErr error

func tryCGImageCalculateContentAverageLightLevel(image CGImageRef) (float32, error) {
	if _cGImageCalculateContentAverageLightLevel == nil {
		return 0.0, symbolCallError("CGImageCalculateContentAverageLightLevel", "26.0", _cGImageCalculateContentAverageLightLevelErr)
	}
	return _cGImageCalculateContentAverageLightLevel(image), nil
}

// CGImageCalculateContentAverageLightLevel.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/calculatedContentAverageLightLevel
func CGImageCalculateContentAverageLightLevel(image CGImageRef) float32 {
	result, callErr := tryCGImageCalculateContentAverageLightLevel(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageCalculateContentHeadroom func(image CGImageRef) float32
var _cGImageCalculateContentHeadroomErr error

func tryCGImageCalculateContentHeadroom(image CGImageRef) (float32, error) {
	if _cGImageCalculateContentHeadroom == nil {
		return 0.0, symbolCallError("CGImageCalculateContentHeadroom", "26.0", _cGImageCalculateContentHeadroomErr)
	}
	return _cGImageCalculateContentHeadroom(image), nil
}

// CGImageCalculateContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/calculatedContentHeadroom
func CGImageCalculateContentHeadroom(image CGImageRef) float32 {
	result, callErr := tryCGImageCalculateContentHeadroom(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageContainsImageSpecificToneMappingMetadata func(image CGImageRef) bool
var _cGImageContainsImageSpecificToneMappingMetadataErr error

func tryCGImageContainsImageSpecificToneMappingMetadata(image CGImageRef) (bool, error) {
	if _cGImageContainsImageSpecificToneMappingMetadata == nil {
		return false, symbolCallError("CGImageContainsImageSpecificToneMappingMetadata", "15.0", _cGImageContainsImageSpecificToneMappingMetadataErr)
	}
	return _cGImageContainsImageSpecificToneMappingMetadata(image), nil
}

// CGImageContainsImageSpecificToneMappingMetadata.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/containsImageSpecificToneMappingMetadata
func CGImageContainsImageSpecificToneMappingMetadata(image CGImageRef) bool {
	result, callErr := tryCGImageContainsImageSpecificToneMappingMetadata(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageCreate func(width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo, provider CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) CGImageRef
var _cGImageCreateErr error

func tryCGImageCreate(width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo, provider CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) (CGImageRef, error) {
	if _cGImageCreate == nil {
		return 0, symbolCallError("CGImageCreate", "10.0", _cGImageCreateErr)
	}
	return _cGImageCreate(width, height, bitsPerComponent, bitsPerPixel, bytesPerRow, space, bitmapInfo, provider, decode, shouldInterpolate, intent), nil
}

// CGImageCreate creates a bitmap image from data supplied by a data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/init(width:height:bitsPerComponent:bitsPerPixel:bytesPerRow:space:bitmapInfo:provider:decode:shouldInterpolate:intent:)
func CGImageCreate(width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo, provider CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) CGImageRef {
	result, callErr := tryCGImageCreate(width, height, bitsPerComponent, bitsPerPixel, bytesPerRow, space, bitmapInfo, provider, decode, shouldInterpolate, intent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageCreateCopy func(image CGImageRef) CGImageRef
var _cGImageCreateCopyErr error

func tryCGImageCreateCopy(image CGImageRef) (CGImageRef, error) {
	if _cGImageCreateCopy == nil {
		return 0, symbolCallError("CGImageCreateCopy", "10.4", _cGImageCreateCopyErr)
	}
	return _cGImageCreateCopy(image), nil
}

// CGImageCreateCopy creates a copy of a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/copy()
func CGImageCreateCopy(image CGImageRef) CGImageRef {
	result, callErr := tryCGImageCreateCopy(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageCreateCopyWithCalculatedHDRStats func(image CGImageRef) CGImageRef
var _cGImageCreateCopyWithCalculatedHDRStatsErr error

func tryCGImageCreateCopyWithCalculatedHDRStats(image CGImageRef) (CGImageRef, error) {
	if _cGImageCreateCopyWithCalculatedHDRStats == nil {
		return 0, symbolCallError("CGImageCreateCopyWithCalculatedHDRStats", "26.0", _cGImageCreateCopyWithCalculatedHDRStatsErr)
	}
	return _cGImageCreateCopyWithCalculatedHDRStats(image), nil
}

// CGImageCreateCopyWithCalculatedHDRStats.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/copyWithCalculatedHDRStats()
func CGImageCreateCopyWithCalculatedHDRStats(image CGImageRef) CGImageRef {
	result, callErr := tryCGImageCreateCopyWithCalculatedHDRStats(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageCreateCopyWithColorSpace func(image CGImageRef, space CGColorSpaceRef) CGImageRef
var _cGImageCreateCopyWithColorSpaceErr error

func tryCGImageCreateCopyWithColorSpace(image CGImageRef, space CGColorSpaceRef) (CGImageRef, error) {
	if _cGImageCreateCopyWithColorSpace == nil {
		return 0, symbolCallError("CGImageCreateCopyWithColorSpace", "10.3", _cGImageCreateCopyWithColorSpaceErr)
	}
	return _cGImageCreateCopyWithColorSpace(image, space), nil
}

// CGImageCreateCopyWithColorSpace creates a copy of a bitmap image, replacing its colorspace.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/copy(colorSpace:)
func CGImageCreateCopyWithColorSpace(image CGImageRef, space CGColorSpaceRef) CGImageRef {
	result, callErr := tryCGImageCreateCopyWithColorSpace(image, space)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageCreateCopyWithContentAverageLightLevel func(image CGImageRef, avll float32) CGImageRef
var _cGImageCreateCopyWithContentAverageLightLevelErr error

func tryCGImageCreateCopyWithContentAverageLightLevel(image CGImageRef, avll float32) (CGImageRef, error) {
	if _cGImageCreateCopyWithContentAverageLightLevel == nil {
		return 0, symbolCallError("CGImageCreateCopyWithContentAverageLightLevel", "26.0", _cGImageCreateCopyWithContentAverageLightLevelErr)
	}
	return _cGImageCreateCopyWithContentAverageLightLevel(image, avll), nil
}

// CGImageCreateCopyWithContentAverageLightLevel.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/copy(contentAverageLightLevel:)
func CGImageCreateCopyWithContentAverageLightLevel(image CGImageRef, avll float32) CGImageRef {
	result, callErr := tryCGImageCreateCopyWithContentAverageLightLevel(image, avll)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageCreateCopyWithContentHeadroom func(headroom float32, image CGImageRef) CGImageRef
var _cGImageCreateCopyWithContentHeadroomErr error

func tryCGImageCreateCopyWithContentHeadroom(headroom float32, image CGImageRef) (CGImageRef, error) {
	if _cGImageCreateCopyWithContentHeadroom == nil {
		return 0, symbolCallError("CGImageCreateCopyWithContentHeadroom", "15.0", _cGImageCreateCopyWithContentHeadroomErr)
	}
	return _cGImageCreateCopyWithContentHeadroom(headroom, image), nil
}

// CGImageCreateCopyWithContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImageCreateCopyWithContentHeadroom(_:_:)
func CGImageCreateCopyWithContentHeadroom(headroom float32, image CGImageRef) CGImageRef {
	result, callErr := tryCGImageCreateCopyWithContentHeadroom(headroom, image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageCreateWithContentHeadroom func(headroom float32, width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo, provider CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) CGImageRef
var _cGImageCreateWithContentHeadroomErr error

func tryCGImageCreateWithContentHeadroom(headroom float32, width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo, provider CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) (CGImageRef, error) {
	if _cGImageCreateWithContentHeadroom == nil {
		return 0, symbolCallError("CGImageCreateWithContentHeadroom", "15.0", _cGImageCreateWithContentHeadroomErr)
	}
	return _cGImageCreateWithContentHeadroom(headroom, width, height, bitsPerComponent, bitsPerPixel, bytesPerRow, space, bitmapInfo, provider, decode, shouldInterpolate, intent), nil
}

// CGImageCreateWithContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/init(headroom:width:height:bitsPerComponent:bitsPerPixel:bytesPerRow:space:bitmapInfo:provider:decode:shouldInterpolate:intent:)
func CGImageCreateWithContentHeadroom(headroom float32, width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, space CGColorSpaceRef, bitmapInfo CGBitmapInfo, provider CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) CGImageRef {
	result, callErr := tryCGImageCreateWithContentHeadroom(headroom, width, height, bitsPerComponent, bitsPerPixel, bytesPerRow, space, bitmapInfo, provider, decode, shouldInterpolate, intent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageCreateWithImageInRect func(image CGImageRef, rect corefoundation.CGRect) CGImageRef
var _cGImageCreateWithImageInRectErr error

func tryCGImageCreateWithImageInRect(image CGImageRef, rect corefoundation.CGRect) (CGImageRef, error) {
	if _cGImageCreateWithImageInRect == nil {
		return 0, symbolCallError("CGImageCreateWithImageInRect", "10.4", _cGImageCreateWithImageInRectErr)
	}
	return _cGImageCreateWithImageInRect(image, rect), nil
}

// CGImageCreateWithImageInRect creates a bitmap image using the data contained within a subregion of an existing bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/cropping(to:)
func CGImageCreateWithImageInRect(image CGImageRef, rect corefoundation.CGRect) CGImageRef {
	result, callErr := tryCGImageCreateWithImageInRect(image, rect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageCreateWithJPEGDataProvider func(source CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) CGImageRef
var _cGImageCreateWithJPEGDataProviderErr error

func tryCGImageCreateWithJPEGDataProvider(source CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) (CGImageRef, error) {
	if _cGImageCreateWithJPEGDataProvider == nil {
		return 0, symbolCallError("CGImageCreateWithJPEGDataProvider", "10.1", _cGImageCreateWithJPEGDataProviderErr)
	}
	return _cGImageCreateWithJPEGDataProvider(source, decode, shouldInterpolate, intent), nil
}

// CGImageCreateWithJPEGDataProvider creates a bitmap image using JPEG-encoded data supplied by a data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/init(jpegDataProviderSource:decode:shouldInterpolate:intent:)
func CGImageCreateWithJPEGDataProvider(source CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) CGImageRef {
	result, callErr := tryCGImageCreateWithJPEGDataProvider(source, decode, shouldInterpolate, intent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageCreateWithMask func(image CGImageRef, mask CGImageRef) CGImageRef
var _cGImageCreateWithMaskErr error

func tryCGImageCreateWithMask(image CGImageRef, mask CGImageRef) (CGImageRef, error) {
	if _cGImageCreateWithMask == nil {
		return 0, symbolCallError("CGImageCreateWithMask", "10.4", _cGImageCreateWithMaskErr)
	}
	return _cGImageCreateWithMask(image, mask), nil
}

// CGImageCreateWithMask creates a bitmap image from an existing image and an image mask.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/masking(_:)
func CGImageCreateWithMask(image CGImageRef, mask CGImageRef) CGImageRef {
	result, callErr := tryCGImageCreateWithMask(image, mask)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageCreateWithMaskingColors func(image CGImageRef, components *float64) CGImageRef
var _cGImageCreateWithMaskingColorsErr error

func tryCGImageCreateWithMaskingColors(image CGImageRef, components *float64) (CGImageRef, error) {
	if _cGImageCreateWithMaskingColors == nil {
		return 0, symbolCallError("CGImageCreateWithMaskingColors", "10.4", _cGImageCreateWithMaskingColorsErr)
	}
	return _cGImageCreateWithMaskingColors(image, components), nil
}

// CGImageCreateWithMaskingColors creates a bitmap image by masking an existing bitmap image with the provided color values.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImageCreateWithMaskingColors
func CGImageCreateWithMaskingColors(image CGImageRef, components *float64) CGImageRef {
	result, callErr := tryCGImageCreateWithMaskingColors(image, components)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageCreateWithPNGDataProvider func(source CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) CGImageRef
var _cGImageCreateWithPNGDataProviderErr error

func tryCGImageCreateWithPNGDataProvider(source CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) (CGImageRef, error) {
	if _cGImageCreateWithPNGDataProvider == nil {
		return 0, symbolCallError("CGImageCreateWithPNGDataProvider", "10.2", _cGImageCreateWithPNGDataProviderErr)
	}
	return _cGImageCreateWithPNGDataProvider(source, decode, shouldInterpolate, intent), nil
}

// CGImageCreateWithPNGDataProvider creates a bitmap image using PNG-encoded data supplied by a data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/init(pngDataProviderSource:decode:shouldInterpolate:intent:)
func CGImageCreateWithPNGDataProvider(source CGDataProviderRef, decode *float64, shouldInterpolate bool, intent CGColorRenderingIntent) CGImageRef {
	result, callErr := tryCGImageCreateWithPNGDataProvider(source, decode, shouldInterpolate, intent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageGetAlphaInfo func(image CGImageRef) CGImageAlphaInfo
var _cGImageGetAlphaInfoErr error

func tryCGImageGetAlphaInfo(image CGImageRef) (CGImageAlphaInfo, error) {
	if _cGImageGetAlphaInfo == nil {
		return *new(CGImageAlphaInfo), symbolCallError("CGImageGetAlphaInfo", "10.0", _cGImageGetAlphaInfoErr)
	}
	return _cGImageGetAlphaInfo(image), nil
}

// CGImageGetAlphaInfo returns the alpha channel information for a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/alphaInfo
func CGImageGetAlphaInfo(image CGImageRef) CGImageAlphaInfo {
	result, callErr := tryCGImageGetAlphaInfo(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageGetBitmapInfo func(image CGImageRef) CGBitmapInfo
var _cGImageGetBitmapInfoErr error

func tryCGImageGetBitmapInfo(image CGImageRef) (CGBitmapInfo, error) {
	if _cGImageGetBitmapInfo == nil {
		return *new(CGBitmapInfo), symbolCallError("CGImageGetBitmapInfo", "10.4", _cGImageGetBitmapInfoErr)
	}
	return _cGImageGetBitmapInfo(image), nil
}

// CGImageGetBitmapInfo returns the bitmap information for a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/bitmapInfo
func CGImageGetBitmapInfo(image CGImageRef) CGBitmapInfo {
	result, callErr := tryCGImageGetBitmapInfo(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageGetBitsPerComponent func(image CGImageRef) uintptr
var _cGImageGetBitsPerComponentErr error

func tryCGImageGetBitsPerComponent(image CGImageRef) (uintptr, error) {
	if _cGImageGetBitsPerComponent == nil {
		return 0, symbolCallError("CGImageGetBitsPerComponent", "10.0", _cGImageGetBitsPerComponentErr)
	}
	return _cGImageGetBitsPerComponent(image), nil
}

// CGImageGetBitsPerComponent returns the number of bits allocated for a single color component of a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/bitsPerComponent
func CGImageGetBitsPerComponent(image CGImageRef) uintptr {
	result, callErr := tryCGImageGetBitsPerComponent(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageGetBitsPerPixel func(image CGImageRef) uintptr
var _cGImageGetBitsPerPixelErr error

func tryCGImageGetBitsPerPixel(image CGImageRef) (uintptr, error) {
	if _cGImageGetBitsPerPixel == nil {
		return 0, symbolCallError("CGImageGetBitsPerPixel", "10.0", _cGImageGetBitsPerPixelErr)
	}
	return _cGImageGetBitsPerPixel(image), nil
}

// CGImageGetBitsPerPixel returns the number of bits allocated for a single pixel in a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/bitsPerPixel
func CGImageGetBitsPerPixel(image CGImageRef) uintptr {
	result, callErr := tryCGImageGetBitsPerPixel(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageGetByteOrderInfo func(image CGImageRef) CGImageByteOrderInfo
var _cGImageGetByteOrderInfoErr error

func tryCGImageGetByteOrderInfo(image CGImageRef) (CGImageByteOrderInfo, error) {
	if _cGImageGetByteOrderInfo == nil {
		return *new(CGImageByteOrderInfo), symbolCallError("CGImageGetByteOrderInfo", "10.14", _cGImageGetByteOrderInfoErr)
	}
	return _cGImageGetByteOrderInfo(image), nil
}

// CGImageGetByteOrderInfo.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/byteOrderInfo
func CGImageGetByteOrderInfo(image CGImageRef) CGImageByteOrderInfo {
	result, callErr := tryCGImageGetByteOrderInfo(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageGetBytesPerRow func(image CGImageRef) uintptr
var _cGImageGetBytesPerRowErr error

func tryCGImageGetBytesPerRow(image CGImageRef) (uintptr, error) {
	if _cGImageGetBytesPerRow == nil {
		return 0, symbolCallError("CGImageGetBytesPerRow", "10.0", _cGImageGetBytesPerRowErr)
	}
	return _cGImageGetBytesPerRow(image), nil
}

// CGImageGetBytesPerRow returns the number of bytes allocated for a single row of a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/bytesPerRow
func CGImageGetBytesPerRow(image CGImageRef) uintptr {
	result, callErr := tryCGImageGetBytesPerRow(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageGetColorSpace func(image CGImageRef) CGColorSpaceRef
var _cGImageGetColorSpaceErr error

func tryCGImageGetColorSpace(image CGImageRef) (CGColorSpaceRef, error) {
	if _cGImageGetColorSpace == nil {
		return 0, symbolCallError("CGImageGetColorSpace", "10.0", _cGImageGetColorSpaceErr)
	}
	return _cGImageGetColorSpace(image), nil
}

// CGImageGetColorSpace return the color space for a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/colorSpace
func CGImageGetColorSpace(image CGImageRef) CGColorSpaceRef {
	result, callErr := tryCGImageGetColorSpace(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageGetContentAverageLightLevel func(image CGImageRef) float32
var _cGImageGetContentAverageLightLevelErr error

func tryCGImageGetContentAverageLightLevel(image CGImageRef) (float32, error) {
	if _cGImageGetContentAverageLightLevel == nil {
		return 0.0, symbolCallError("CGImageGetContentAverageLightLevel", "26.0", _cGImageGetContentAverageLightLevelErr)
	}
	return _cGImageGetContentAverageLightLevel(image), nil
}

// CGImageGetContentAverageLightLevel.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/contentAverageLightLevel
func CGImageGetContentAverageLightLevel(image CGImageRef) float32 {
	result, callErr := tryCGImageGetContentAverageLightLevel(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageGetContentHeadroom func(image CGImageRef) float32
var _cGImageGetContentHeadroomErr error

func tryCGImageGetContentHeadroom(image CGImageRef) (float32, error) {
	if _cGImageGetContentHeadroom == nil {
		return 0.0, symbolCallError("CGImageGetContentHeadroom", "15.0", _cGImageGetContentHeadroomErr)
	}
	return _cGImageGetContentHeadroom(image), nil
}

// CGImageGetContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/contentHeadroom
func CGImageGetContentHeadroom(image CGImageRef) float32 {
	result, callErr := tryCGImageGetContentHeadroom(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageGetDataProvider func(image CGImageRef) CGDataProviderRef
var _cGImageGetDataProviderErr error

func tryCGImageGetDataProvider(image CGImageRef) (CGDataProviderRef, error) {
	if _cGImageGetDataProvider == nil {
		return 0, symbolCallError("CGImageGetDataProvider", "10.0", _cGImageGetDataProviderErr)
	}
	return _cGImageGetDataProvider(image), nil
}

// CGImageGetDataProvider returns the data provider for a bitmap image or image mask.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/dataProvider
func CGImageGetDataProvider(image CGImageRef) CGDataProviderRef {
	result, callErr := tryCGImageGetDataProvider(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageGetDecode func(image CGImageRef) *float64
var _cGImageGetDecodeErr error

func tryCGImageGetDecode(image CGImageRef) (*float64, error) {
	if _cGImageGetDecode == nil {
		return nil, symbolCallError("CGImageGetDecode", "10.0", _cGImageGetDecodeErr)
	}
	return _cGImageGetDecode(image), nil
}

// CGImageGetDecode returns the decode array for a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/decode
func CGImageGetDecode(image CGImageRef) *float64 {
	result, callErr := tryCGImageGetDecode(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageGetHeight func(image CGImageRef) uintptr
var _cGImageGetHeightErr error

func tryCGImageGetHeight(image CGImageRef) (uintptr, error) {
	if _cGImageGetHeight == nil {
		return 0, symbolCallError("CGImageGetHeight", "10.0", _cGImageGetHeightErr)
	}
	return _cGImageGetHeight(image), nil
}

// CGImageGetHeight returns the height of a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/height
func CGImageGetHeight(image CGImageRef) uintptr {
	result, callErr := tryCGImageGetHeight(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageGetPixelFormatInfo func(image CGImageRef) CGImagePixelFormatInfo
var _cGImageGetPixelFormatInfoErr error

func tryCGImageGetPixelFormatInfo(image CGImageRef) (CGImagePixelFormatInfo, error) {
	if _cGImageGetPixelFormatInfo == nil {
		return *new(CGImagePixelFormatInfo), symbolCallError("CGImageGetPixelFormatInfo", "10.14", _cGImageGetPixelFormatInfoErr)
	}
	return _cGImageGetPixelFormatInfo(image), nil
}

// CGImageGetPixelFormatInfo.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/pixelFormatInfo
func CGImageGetPixelFormatInfo(image CGImageRef) CGImagePixelFormatInfo {
	result, callErr := tryCGImageGetPixelFormatInfo(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageGetRenderingIntent func(image CGImageRef) CGColorRenderingIntent
var _cGImageGetRenderingIntentErr error

func tryCGImageGetRenderingIntent(image CGImageRef) (CGColorRenderingIntent, error) {
	if _cGImageGetRenderingIntent == nil {
		return *new(CGColorRenderingIntent), symbolCallError("CGImageGetRenderingIntent", "10.0", _cGImageGetRenderingIntentErr)
	}
	return _cGImageGetRenderingIntent(image), nil
}

// CGImageGetRenderingIntent returns the rendering intent setting for a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/renderingIntent
func CGImageGetRenderingIntent(image CGImageRef) CGColorRenderingIntent {
	result, callErr := tryCGImageGetRenderingIntent(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageGetShouldInterpolate func(image CGImageRef) bool
var _cGImageGetShouldInterpolateErr error

func tryCGImageGetShouldInterpolate(image CGImageRef) (bool, error) {
	if _cGImageGetShouldInterpolate == nil {
		return false, symbolCallError("CGImageGetShouldInterpolate", "10.0", _cGImageGetShouldInterpolateErr)
	}
	return _cGImageGetShouldInterpolate(image), nil
}

// CGImageGetShouldInterpolate returns the interpolation setting for a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/shouldInterpolate
func CGImageGetShouldInterpolate(image CGImageRef) bool {
	result, callErr := tryCGImageGetShouldInterpolate(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageGetTypeID func() uint
var _cGImageGetTypeIDErr error

func tryCGImageGetTypeID() (uint, error) {
	if _cGImageGetTypeID == nil {
		return 0, symbolCallError("CGImageGetTypeID", "10.2", _cGImageGetTypeIDErr)
	}
	return _cGImageGetTypeID(), nil
}

// CGImageGetTypeID returns the type identifier for CGImage objects.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/typeID
func CGImageGetTypeID() uint {
	result, callErr := tryCGImageGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageGetUTType func(image CGImageRef) corefoundation.CFStringRef
var _cGImageGetUTTypeErr error

func tryCGImageGetUTType(image CGImageRef) (corefoundation.CFStringRef, error) {
	if _cGImageGetUTType == nil {
		return 0, symbolCallError("CGImageGetUTType", "10.11", _cGImageGetUTTypeErr)
	}
	return _cGImageGetUTType(image), nil
}

// CGImageGetUTType the Universal Type Identifier for the image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/utType
func CGImageGetUTType(image CGImageRef) corefoundation.CFStringRef {
	result, callErr := tryCGImageGetUTType(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageGetWidth func(image CGImageRef) uintptr
var _cGImageGetWidthErr error

func tryCGImageGetWidth(image CGImageRef) (uintptr, error) {
	if _cGImageGetWidth == nil {
		return 0, symbolCallError("CGImageGetWidth", "10.0", _cGImageGetWidthErr)
	}
	return _cGImageGetWidth(image), nil
}

// CGImageGetWidth returns the width of a bitmap image, in pixels.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/width
func CGImageGetWidth(image CGImageRef) uintptr {
	result, callErr := tryCGImageGetWidth(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageIsMask func(image CGImageRef) bool
var _cGImageIsMaskErr error

func tryCGImageIsMask(image CGImageRef) (bool, error) {
	if _cGImageIsMask == nil {
		return false, symbolCallError("CGImageIsMask", "10.0", _cGImageIsMaskErr)
	}
	return _cGImageIsMask(image), nil
}

// CGImageIsMask returns whether a bitmap image is an image mask.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/isMask
func CGImageIsMask(image CGImageRef) bool {
	result, callErr := tryCGImageIsMask(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMaskCreate func(width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, provider CGDataProviderRef, decode *float64, shouldInterpolate bool) CGImageRef
var _cGImageMaskCreateErr error

func tryCGImageMaskCreate(width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, provider CGDataProviderRef, decode *float64, shouldInterpolate bool) (CGImageRef, error) {
	if _cGImageMaskCreate == nil {
		return 0, symbolCallError("CGImageMaskCreate", "10.0", _cGImageMaskCreateErr)
	}
	return _cGImageMaskCreate(width, height, bitsPerComponent, bitsPerPixel, bytesPerRow, provider, decode, shouldInterpolate), nil
}

// CGImageMaskCreate creates a bitmap image mask from data supplied by a data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/init(maskWidth:height:bitsPerComponent:bitsPerPixel:bytesPerRow:provider:decode:shouldInterpolate:)
func CGImageMaskCreate(width uintptr, height uintptr, bitsPerComponent uintptr, bitsPerPixel uintptr, bytesPerRow uintptr, provider CGDataProviderRef, decode *float64, shouldInterpolate bool) CGImageRef {
	result, callErr := tryCGImageMaskCreate(width, height, bitsPerComponent, bitsPerPixel, bytesPerRow, provider, decode, shouldInterpolate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageRelease func(image CGImageRef)
var _cGImageReleaseErr error

func tryCGImageRelease(image CGImageRef) error {
	if _cGImageRelease == nil {
		return symbolCallError("CGImageRelease", "10.0", _cGImageReleaseErr)
	}
	_cGImageRelease(image)
	return nil
}

// CGImageRelease decrements the retain count of a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImageRelease
func CGImageRelease(image CGImageRef) {
	if callErr := tryCGImageRelease(image); callErr != nil {
		panic(callErr)
	}
}

var _cGImageRetain func(image CGImageRef) CGImageRef
var _cGImageRetainErr error

func tryCGImageRetain(image CGImageRef) (CGImageRef, error) {
	if _cGImageRetain == nil {
		return 0, symbolCallError("CGImageRetain", "10.0", _cGImageRetainErr)
	}
	return _cGImageRetain(image), nil
}

// CGImageRetain increments the retain count of a bitmap image.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImageRetain
func CGImageRetain(image CGImageRef) CGImageRef {
	result, callErr := tryCGImageRetain(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageShouldToneMap func(image CGImageRef) bool
var _cGImageShouldToneMapErr error

func tryCGImageShouldToneMap(image CGImageRef) (bool, error) {
	if _cGImageShouldToneMap == nil {
		return false, symbolCallError("CGImageShouldToneMap", "15.0", _cGImageShouldToneMapErr)
	}
	return _cGImageShouldToneMap(image), nil
}

// CGImageShouldToneMap.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/shouldToneMap
func CGImageShouldToneMap(image CGImageRef) bool {
	result, callErr := tryCGImageShouldToneMap(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGLayerCreateWithContext func(context CGContextRef, size corefoundation.CGSize, auxiliaryInfo corefoundation.CFDictionaryRef) unsafe.Pointer
var _cGLayerCreateWithContextErr error

func tryCGLayerCreateWithContext(context CGContextRef, size corefoundation.CGSize, auxiliaryInfo corefoundation.CFDictionaryRef) (unsafe.Pointer, error) {
	if _cGLayerCreateWithContext == nil {
		return nil, symbolCallError("CGLayerCreateWithContext", "10.4", _cGLayerCreateWithContextErr)
	}
	return _cGLayerCreateWithContext(context, size, auxiliaryInfo), nil
}

// CGLayerCreateWithContext creates a layer object that is associated with a graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGLayer/init(_:size:auxiliaryInfo:)
func CGLayerCreateWithContext(context CGContextRef, size corefoundation.CGSize, auxiliaryInfo corefoundation.CFDictionaryRef) unsafe.Pointer {
	result, callErr := tryCGLayerCreateWithContext(context, size, auxiliaryInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGLayerGetContext func(layer unsafe.Pointer) CGContextRef
var _cGLayerGetContextErr error

func tryCGLayerGetContext(layer unsafe.Pointer) (CGContextRef, error) {
	if _cGLayerGetContext == nil {
		return 0, symbolCallError("CGLayerGetContext", "10.4", _cGLayerGetContextErr)
	}
	return _cGLayerGetContext(layer), nil
}

// CGLayerGetContext returns the graphics context associated with a layer object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGLayer/context
func CGLayerGetContext(layer unsafe.Pointer) CGContextRef {
	result, callErr := tryCGLayerGetContext(layer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGLayerGetSize func(layer unsafe.Pointer) corefoundation.CGSize
var _cGLayerGetSizeErr error

func tryCGLayerGetSize(layer unsafe.Pointer) (corefoundation.CGSize, error) {
	if _cGLayerGetSize == nil {
		return corefoundation.CGSize{}, symbolCallError("CGLayerGetSize", "10.4", _cGLayerGetSizeErr)
	}
	return _cGLayerGetSize(layer), nil
}

// CGLayerGetSize returns the width and height of a layer object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGLayer/size
func CGLayerGetSize(layer unsafe.Pointer) corefoundation.CGSize {
	result, callErr := tryCGLayerGetSize(layer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGLayerGetTypeID func() uint
var _cGLayerGetTypeIDErr error

func tryCGLayerGetTypeID() (uint, error) {
	if _cGLayerGetTypeID == nil {
		return 0, symbolCallError("CGLayerGetTypeID", "10.4", _cGLayerGetTypeIDErr)
	}
	return _cGLayerGetTypeID(), nil
}

// CGLayerGetTypeID returns the unique type identifier used for CGLayer objects.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGLayer/typeID
func CGLayerGetTypeID() uint {
	result, callErr := tryCGLayerGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGLayerRelease func(layer unsafe.Pointer)
var _cGLayerReleaseErr error

func tryCGLayerRelease(layer unsafe.Pointer) error {
	if _cGLayerRelease == nil {
		return symbolCallError("CGLayerRelease", "10.4", _cGLayerReleaseErr)
	}
	_cGLayerRelease(layer)
	return nil
}

// CGLayerRelease decrements the retain count of a layer object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGLayerRelease
func CGLayerRelease(layer unsafe.Pointer) {
	if callErr := tryCGLayerRelease(layer); callErr != nil {
		panic(callErr)
	}
}

var _cGLayerRetain func(layer unsafe.Pointer) unsafe.Pointer
var _cGLayerRetainErr error

func tryCGLayerRetain(layer unsafe.Pointer) (unsafe.Pointer, error) {
	if _cGLayerRetain == nil {
		return nil, symbolCallError("CGLayerRetain", "10.4", _cGLayerRetainErr)
	}
	return _cGLayerRetain(layer), nil
}

// CGLayerRetain increments the retain count of a layer object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGLayerRetain
func CGLayerRetain(layer unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryCGLayerRetain(layer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGMainDisplayID func() uint32
var _cGMainDisplayIDErr error

func tryCGMainDisplayID() (uint32, error) {
	if _cGMainDisplayID == nil {
		return 0, symbolCallError("CGMainDisplayID", "10.2", _cGMainDisplayIDErr)
	}
	return _cGMainDisplayID(), nil
}

// CGMainDisplayID returns the display ID of the main display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGMainDisplayID()
func CGMainDisplayID() uint32 {
	result, callErr := tryCGMainDisplayID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGOpenGLDisplayMaskToDisplayID func(mask CGOpenGLDisplayMask) uint32
var _cGOpenGLDisplayMaskToDisplayIDErr error

func tryCGOpenGLDisplayMaskToDisplayID(mask CGOpenGLDisplayMask) (uint32, error) {
	if _cGOpenGLDisplayMaskToDisplayID == nil {
		return 0, symbolCallError("CGOpenGLDisplayMaskToDisplayID", "10.2", _cGOpenGLDisplayMaskToDisplayIDErr)
	}
	return _cGOpenGLDisplayMaskToDisplayID(mask), nil
}

// CGOpenGLDisplayMaskToDisplayID maps an OpenGL display mask to a display ID.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGOpenGLDisplayMaskToDisplayID(_:)
func CGOpenGLDisplayMaskToDisplayID(mask CGOpenGLDisplayMask) uint32 {
	result, callErr := tryCGOpenGLDisplayMaskToDisplayID(mask)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFArrayApplyBlock func(array CGPDFArrayRef, block unsafe.Pointer, info unsafe.Pointer)
var _cGPDFArrayApplyBlockErr error

func tryCGPDFArrayApplyBlock(array CGPDFArrayRef, block CGPDFArrayApplierBlock, info unsafe.Pointer) error {
	if _cGPDFArrayApplyBlock == nil {
		return symbolCallError("CGPDFArrayApplyBlock", "10.14", _cGPDFArrayApplyBlockErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 uint32, arg1 *CGPDFObjectRef, arg2 unsafe.Pointer) bool {
		return block(arg0, arg1, arg2)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_cGPDFArrayApplyBlock(array, _block0, info)
	return nil
}

// CGPDFArrayApplyBlock.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayApplyBlock(_:_:_:)
func CGPDFArrayApplyBlock(array CGPDFArrayRef, block CGPDFArrayApplierBlock, info unsafe.Pointer) {
	if callErr := tryCGPDFArrayApplyBlock(array, block, info); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFArrayGetArray func(array CGPDFArrayRef, index uintptr, value *CGPDFArrayRef) bool
var _cGPDFArrayGetArrayErr error

func tryCGPDFArrayGetArray(array CGPDFArrayRef, index uintptr, value *CGPDFArrayRef) (bool, error) {
	if _cGPDFArrayGetArray == nil {
		return false, symbolCallError("CGPDFArrayGetArray", "10.3", _cGPDFArrayGetArrayErr)
	}
	return _cGPDFArrayGetArray(array, index, value), nil
}

// CGPDFArrayGetArray returns whether an object at a given index in a PDF array is another PDF array and, if so, retrieves that array.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetArray(_:_:_:)
func CGPDFArrayGetArray(array CGPDFArrayRef, index uintptr, value *CGPDFArrayRef) bool {
	result, callErr := tryCGPDFArrayGetArray(array, index, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFArrayGetBoolean func(array CGPDFArrayRef, index uintptr, value *CGPDFBoolean) bool
var _cGPDFArrayGetBooleanErr error

func tryCGPDFArrayGetBoolean(array CGPDFArrayRef, index uintptr, value *CGPDFBoolean) (bool, error) {
	if _cGPDFArrayGetBoolean == nil {
		return false, symbolCallError("CGPDFArrayGetBoolean", "10.3", _cGPDFArrayGetBooleanErr)
	}
	return _cGPDFArrayGetBoolean(array, index, value), nil
}

// CGPDFArrayGetBoolean returns whether an object at a given index in a PDF array is a PDF Boolean and, if so, retrieves that Boolean.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetBoolean(_:_:_:)
func CGPDFArrayGetBoolean(array CGPDFArrayRef, index uintptr, value *CGPDFBoolean) bool {
	result, callErr := tryCGPDFArrayGetBoolean(array, index, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFArrayGetCount func(array CGPDFArrayRef) uintptr
var _cGPDFArrayGetCountErr error

func tryCGPDFArrayGetCount(array CGPDFArrayRef) (uintptr, error) {
	if _cGPDFArrayGetCount == nil {
		return 0, symbolCallError("CGPDFArrayGetCount", "10.3", _cGPDFArrayGetCountErr)
	}
	return _cGPDFArrayGetCount(array), nil
}

// CGPDFArrayGetCount returns the number of items in a PDF array.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetCount(_:)
func CGPDFArrayGetCount(array CGPDFArrayRef) uintptr {
	result, callErr := tryCGPDFArrayGetCount(array)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFArrayGetDictionary func(array CGPDFArrayRef, index uintptr, value *CGPDFDictionaryRef) bool
var _cGPDFArrayGetDictionaryErr error

func tryCGPDFArrayGetDictionary(array CGPDFArrayRef, index uintptr, value *CGPDFDictionaryRef) (bool, error) {
	if _cGPDFArrayGetDictionary == nil {
		return false, symbolCallError("CGPDFArrayGetDictionary", "10.3", _cGPDFArrayGetDictionaryErr)
	}
	return _cGPDFArrayGetDictionary(array, index, value), nil
}

// CGPDFArrayGetDictionary returns whether an object at a given index in a PDF array is a PDF dictionary and, if so, retrieves that dictionary.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetDictionary(_:_:_:)
func CGPDFArrayGetDictionary(array CGPDFArrayRef, index uintptr, value *CGPDFDictionaryRef) bool {
	result, callErr := tryCGPDFArrayGetDictionary(array, index, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFArrayGetInteger func(array CGPDFArrayRef, index uintptr, value *CGPDFInteger) bool
var _cGPDFArrayGetIntegerErr error

func tryCGPDFArrayGetInteger(array CGPDFArrayRef, index uintptr, value *CGPDFInteger) (bool, error) {
	if _cGPDFArrayGetInteger == nil {
		return false, symbolCallError("CGPDFArrayGetInteger", "10.3", _cGPDFArrayGetIntegerErr)
	}
	return _cGPDFArrayGetInteger(array, index, value), nil
}

// CGPDFArrayGetInteger returns whether an object at a given index in a PDF array is a PDF integer and, if so, retrieves that object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetInteger(_:_:_:)
func CGPDFArrayGetInteger(array CGPDFArrayRef, index uintptr, value *CGPDFInteger) bool {
	result, callErr := tryCGPDFArrayGetInteger(array, index, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFArrayGetName func(array CGPDFArrayRef, index uintptr, value string) bool
var _cGPDFArrayGetNameErr error

func tryCGPDFArrayGetName(array CGPDFArrayRef, index uintptr, value string) (bool, error) {
	if _cGPDFArrayGetName == nil {
		return false, symbolCallError("CGPDFArrayGetName", "10.3", _cGPDFArrayGetNameErr)
	}
	return _cGPDFArrayGetName(array, index, value), nil
}

// CGPDFArrayGetName returns whether an object at a given index in a PDF array is a PDF name reference (represented as a constant C string) and, if so, retrieves that name.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetName(_:_:_:)
func CGPDFArrayGetName(array CGPDFArrayRef, index uintptr, value string) bool {
	result, callErr := tryCGPDFArrayGetName(array, index, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFArrayGetNull func(array CGPDFArrayRef, index uintptr) bool
var _cGPDFArrayGetNullErr error

func tryCGPDFArrayGetNull(array CGPDFArrayRef, index uintptr) (bool, error) {
	if _cGPDFArrayGetNull == nil {
		return false, symbolCallError("CGPDFArrayGetNull", "10.3", _cGPDFArrayGetNullErr)
	}
	return _cGPDFArrayGetNull(array, index), nil
}

// CGPDFArrayGetNull returns whether an object at a given index in a Quartz PDF array is a PDF null.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetNull(_:_:)
func CGPDFArrayGetNull(array CGPDFArrayRef, index uintptr) bool {
	result, callErr := tryCGPDFArrayGetNull(array, index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFArrayGetNumber func(array CGPDFArrayRef, index uintptr, value *CGPDFReal) bool
var _cGPDFArrayGetNumberErr error

func tryCGPDFArrayGetNumber(array CGPDFArrayRef, index uintptr, value *CGPDFReal) (bool, error) {
	if _cGPDFArrayGetNumber == nil {
		return false, symbolCallError("CGPDFArrayGetNumber", "10.3", _cGPDFArrayGetNumberErr)
	}
	return _cGPDFArrayGetNumber(array, index, value), nil
}

// CGPDFArrayGetNumber returns whether an object at a given index in a PDF array is a PDF number and, if so, retrieves that object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetNumber(_:_:_:)
func CGPDFArrayGetNumber(array CGPDFArrayRef, index uintptr, value *CGPDFReal) bool {
	result, callErr := tryCGPDFArrayGetNumber(array, index, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFArrayGetObject func(array CGPDFArrayRef, index uintptr, value *CGPDFObjectRef) bool
var _cGPDFArrayGetObjectErr error

func tryCGPDFArrayGetObject(array CGPDFArrayRef, index uintptr, value *CGPDFObjectRef) (bool, error) {
	if _cGPDFArrayGetObject == nil {
		return false, symbolCallError("CGPDFArrayGetObject", "10.3", _cGPDFArrayGetObjectErr)
	}
	return _cGPDFArrayGetObject(array, index, value), nil
}

// CGPDFArrayGetObject returns whether an object at a given index in a PDF array is a PDF object and, if so, retrieves that object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetObject(_:_:_:)
func CGPDFArrayGetObject(array CGPDFArrayRef, index uintptr, value *CGPDFObjectRef) bool {
	result, callErr := tryCGPDFArrayGetObject(array, index, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFArrayGetStream func(array CGPDFArrayRef, index uintptr, value *CGPDFStreamRef) bool
var _cGPDFArrayGetStreamErr error

func tryCGPDFArrayGetStream(array CGPDFArrayRef, index uintptr, value *CGPDFStreamRef) (bool, error) {
	if _cGPDFArrayGetStream == nil {
		return false, symbolCallError("CGPDFArrayGetStream", "10.3", _cGPDFArrayGetStreamErr)
	}
	return _cGPDFArrayGetStream(array, index, value), nil
}

// CGPDFArrayGetStream returns whether an object at a given index in a PDF array is a PDF stream and, if so, retrieves that stream.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetStream(_:_:_:)
func CGPDFArrayGetStream(array CGPDFArrayRef, index uintptr, value *CGPDFStreamRef) bool {
	result, callErr := tryCGPDFArrayGetStream(array, index, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFArrayGetString func(array CGPDFArrayRef, index uintptr, value *CGPDFStringRef) bool
var _cGPDFArrayGetStringErr error

func tryCGPDFArrayGetString(array CGPDFArrayRef, index uintptr, value *CGPDFStringRef) (bool, error) {
	if _cGPDFArrayGetString == nil {
		return false, symbolCallError("CGPDFArrayGetString", "10.3", _cGPDFArrayGetStringErr)
	}
	return _cGPDFArrayGetString(array, index, value), nil
}

// CGPDFArrayGetString returns whether an object at a given index in a PDF array is a PDF string and, if so, retrieves that string.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFArrayGetString(_:_:_:)
func CGPDFArrayGetString(array CGPDFArrayRef, index uintptr, value *CGPDFStringRef) bool {
	result, callErr := tryCGPDFArrayGetString(array, index, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFContentStreamCreateWithPage func(page CGPDFPageRef) CGPDFContentStreamRef
var _cGPDFContentStreamCreateWithPageErr error

func tryCGPDFContentStreamCreateWithPage(page CGPDFPageRef) (CGPDFContentStreamRef, error) {
	if _cGPDFContentStreamCreateWithPage == nil {
		return 0, symbolCallError("CGPDFContentStreamCreateWithPage", "10.4", _cGPDFContentStreamCreateWithPageErr)
	}
	return _cGPDFContentStreamCreateWithPage(page), nil
}

// CGPDFContentStreamCreateWithPage creates a content stream object from a PDF page object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContentStreamCreateWithPage(_:)
func CGPDFContentStreamCreateWithPage(page CGPDFPageRef) CGPDFContentStreamRef {
	result, callErr := tryCGPDFContentStreamCreateWithPage(page)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFContentStreamCreateWithStream func(stream CGPDFStreamRef, streamResources CGPDFDictionaryRef, parent CGPDFContentStreamRef) CGPDFContentStreamRef
var _cGPDFContentStreamCreateWithStreamErr error

func tryCGPDFContentStreamCreateWithStream(stream CGPDFStreamRef, streamResources CGPDFDictionaryRef, parent CGPDFContentStreamRef) (CGPDFContentStreamRef, error) {
	if _cGPDFContentStreamCreateWithStream == nil {
		return 0, symbolCallError("CGPDFContentStreamCreateWithStream", "10.4", _cGPDFContentStreamCreateWithStreamErr)
	}
	return _cGPDFContentStreamCreateWithStream(stream, streamResources, parent), nil
}

// CGPDFContentStreamCreateWithStream creates a PDF content stream object from an existing PDF content stream object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContentStreamCreateWithStream(_:_:_:)
func CGPDFContentStreamCreateWithStream(stream CGPDFStreamRef, streamResources CGPDFDictionaryRef, parent CGPDFContentStreamRef) CGPDFContentStreamRef {
	result, callErr := tryCGPDFContentStreamCreateWithStream(stream, streamResources, parent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFContentStreamGetResource func(cs CGPDFContentStreamRef, category string, name string) CGPDFObjectRef
var _cGPDFContentStreamGetResourceErr error

func tryCGPDFContentStreamGetResource(cs CGPDFContentStreamRef, category string, name string) (CGPDFObjectRef, error) {
	if _cGPDFContentStreamGetResource == nil {
		return 0, symbolCallError("CGPDFContentStreamGetResource", "10.4", _cGPDFContentStreamGetResourceErr)
	}
	return _cGPDFContentStreamGetResource(cs, category, name), nil
}

// CGPDFContentStreamGetResource gets the specified resource from a PDF content stream object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContentStreamGetResource(_:_:_:)
func CGPDFContentStreamGetResource(cs CGPDFContentStreamRef, category string, name string) CGPDFObjectRef {
	result, callErr := tryCGPDFContentStreamGetResource(cs, category, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFContentStreamGetStreams func(cs CGPDFContentStreamRef) corefoundation.CFArrayRef
var _cGPDFContentStreamGetStreamsErr error

func tryCGPDFContentStreamGetStreams(cs CGPDFContentStreamRef) (corefoundation.CFArrayRef, error) {
	if _cGPDFContentStreamGetStreams == nil {
		return 0, symbolCallError("CGPDFContentStreamGetStreams", "10.4", _cGPDFContentStreamGetStreamsErr)
	}
	return _cGPDFContentStreamGetStreams(cs), nil
}

// CGPDFContentStreamGetStreams gets the array of PDF content streams contained in a PDF content stream object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContentStreamGetStreams(_:)
func CGPDFContentStreamGetStreams(cs CGPDFContentStreamRef) corefoundation.CFArrayRef {
	result, callErr := tryCGPDFContentStreamGetStreams(cs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFContentStreamRelease func(cs CGPDFContentStreamRef)
var _cGPDFContentStreamReleaseErr error

func tryCGPDFContentStreamRelease(cs CGPDFContentStreamRef) error {
	if _cGPDFContentStreamRelease == nil {
		return symbolCallError("CGPDFContentStreamRelease", "10.4", _cGPDFContentStreamReleaseErr)
	}
	_cGPDFContentStreamRelease(cs)
	return nil
}

// CGPDFContentStreamRelease decrements the retain count of a PDF content stream object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContentStreamRelease(_:)
func CGPDFContentStreamRelease(cs CGPDFContentStreamRef) {
	if callErr := tryCGPDFContentStreamRelease(cs); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFContentStreamRetain func(cs CGPDFContentStreamRef) CGPDFContentStreamRef
var _cGPDFContentStreamRetainErr error

func tryCGPDFContentStreamRetain(cs CGPDFContentStreamRef) (CGPDFContentStreamRef, error) {
	if _cGPDFContentStreamRetain == nil {
		return 0, symbolCallError("CGPDFContentStreamRetain", "10.4", _cGPDFContentStreamRetainErr)
	}
	return _cGPDFContentStreamRetain(cs), nil
}

// CGPDFContentStreamRetain increments the retain count of a PDF content stream object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContentStreamRetain(_:)
func CGPDFContentStreamRetain(cs CGPDFContentStreamRef) CGPDFContentStreamRef {
	result, callErr := tryCGPDFContentStreamRetain(cs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFContextAddDestinationAtPoint func(context CGContextRef, name corefoundation.CFStringRef, point corefoundation.CGPoint)
var _cGPDFContextAddDestinationAtPointErr error

func tryCGPDFContextAddDestinationAtPoint(context CGContextRef, name corefoundation.CFStringRef, point corefoundation.CGPoint) error {
	if _cGPDFContextAddDestinationAtPoint == nil {
		return symbolCallError("CGPDFContextAddDestinationAtPoint", "10.4", _cGPDFContextAddDestinationAtPointErr)
	}
	_cGPDFContextAddDestinationAtPoint(context, name, point)
	return nil
}

// CGPDFContextAddDestinationAtPoint sets a destination to jump to when a point in the current page of a PDF graphics context is clicked.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/addDestination(_:at:)
func CGPDFContextAddDestinationAtPoint(context CGContextRef, name corefoundation.CFStringRef, point corefoundation.CGPoint) {
	if callErr := tryCGPDFContextAddDestinationAtPoint(context, name, point); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFContextAddDocumentMetadata func(context CGContextRef, metadata corefoundation.CFDataRef)
var _cGPDFContextAddDocumentMetadataErr error

func tryCGPDFContextAddDocumentMetadata(context CGContextRef, metadata corefoundation.CFDataRef) error {
	if _cGPDFContextAddDocumentMetadata == nil {
		return symbolCallError("CGPDFContextAddDocumentMetadata", "10.7", _cGPDFContextAddDocumentMetadataErr)
	}
	_cGPDFContextAddDocumentMetadata(context, metadata)
	return nil
}

// CGPDFContextAddDocumentMetadata associates custom metadata with the PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/addDocumentMetadata(_:)
func CGPDFContextAddDocumentMetadata(context CGContextRef, metadata corefoundation.CFDataRef) {
	if callErr := tryCGPDFContextAddDocumentMetadata(context, metadata); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFContextBeginPage func(context CGContextRef, pageInfo corefoundation.CFDictionaryRef)
var _cGPDFContextBeginPageErr error

func tryCGPDFContextBeginPage(context CGContextRef, pageInfo corefoundation.CFDictionaryRef) error {
	if _cGPDFContextBeginPage == nil {
		return symbolCallError("CGPDFContextBeginPage", "10.4", _cGPDFContextBeginPageErr)
	}
	_cGPDFContextBeginPage(context, pageInfo)
	return nil
}

// CGPDFContextBeginPage begins a new page in a PDF graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/beginPDFPage(_:)
func CGPDFContextBeginPage(context CGContextRef, pageInfo corefoundation.CFDictionaryRef) {
	if callErr := tryCGPDFContextBeginPage(context, pageInfo); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFContextBeginTag func(context CGContextRef, tagType CGPDFTagType, tagProperties corefoundation.CFDictionaryRef)
var _cGPDFContextBeginTagErr error

func tryCGPDFContextBeginTag(context CGContextRef, tagType CGPDFTagType, tagProperties corefoundation.CFDictionaryRef) error {
	if _cGPDFContextBeginTag == nil {
		return symbolCallError("CGPDFContextBeginTag", "10.15", _cGPDFContextBeginTagErr)
	}
	_cGPDFContextBeginTag(context, tagType, tagProperties)
	return nil
}

// CGPDFContextBeginTag.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContextBeginTag(_:_:_:)
func CGPDFContextBeginTag(context CGContextRef, tagType CGPDFTagType, tagProperties corefoundation.CFDictionaryRef) {
	if callErr := tryCGPDFContextBeginTag(context, tagType, tagProperties); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFContextClose func(context CGContextRef)
var _cGPDFContextCloseErr error

func tryCGPDFContextClose(context CGContextRef) error {
	if _cGPDFContextClose == nil {
		return symbolCallError("CGPDFContextClose", "10.5", _cGPDFContextCloseErr)
	}
	_cGPDFContextClose(context)
	return nil
}

// CGPDFContextClose closes a PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/closePDF()
func CGPDFContextClose(context CGContextRef) {
	if callErr := tryCGPDFContextClose(context); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFContextCreate func(consumer CGDataConsumerRef, mediaBox *corefoundation.CGRect, auxiliaryInfo corefoundation.CFDictionaryRef) CGContextRef
var _cGPDFContextCreateErr error

func tryCGPDFContextCreate(consumer CGDataConsumerRef, mediaBox *corefoundation.CGRect, auxiliaryInfo corefoundation.CFDictionaryRef) (CGContextRef, error) {
	if _cGPDFContextCreate == nil {
		return 0, symbolCallError("CGPDFContextCreate", "10.0", _cGPDFContextCreateErr)
	}
	return _cGPDFContextCreate(consumer, mediaBox, auxiliaryInfo), nil
}

// CGPDFContextCreate creates a PDF graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/init(consumer:mediaBox:_:)
func CGPDFContextCreate(consumer CGDataConsumerRef, mediaBox *corefoundation.CGRect, auxiliaryInfo corefoundation.CFDictionaryRef) CGContextRef {
	result, callErr := tryCGPDFContextCreate(consumer, mediaBox, auxiliaryInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFContextCreateWithURL func(url corefoundation.CFURLRef, mediaBox *corefoundation.CGRect, auxiliaryInfo corefoundation.CFDictionaryRef) CGContextRef
var _cGPDFContextCreateWithURLErr error

func tryCGPDFContextCreateWithURL(url corefoundation.CFURLRef, mediaBox *corefoundation.CGRect, auxiliaryInfo corefoundation.CFDictionaryRef) (CGContextRef, error) {
	if _cGPDFContextCreateWithURL == nil {
		return 0, symbolCallError("CGPDFContextCreateWithURL", "10.0", _cGPDFContextCreateWithURLErr)
	}
	return _cGPDFContextCreateWithURL(url, mediaBox, auxiliaryInfo), nil
}

// CGPDFContextCreateWithURL creates a URL-based PDF graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/init(_:mediaBox:_:)
func CGPDFContextCreateWithURL(url corefoundation.CFURLRef, mediaBox *corefoundation.CGRect, auxiliaryInfo corefoundation.CFDictionaryRef) CGContextRef {
	result, callErr := tryCGPDFContextCreateWithURL(url, mediaBox, auxiliaryInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFContextEndPage func(context CGContextRef)
var _cGPDFContextEndPageErr error

func tryCGPDFContextEndPage(context CGContextRef) error {
	if _cGPDFContextEndPage == nil {
		return symbolCallError("CGPDFContextEndPage", "10.4", _cGPDFContextEndPageErr)
	}
	_cGPDFContextEndPage(context)
	return nil
}

// CGPDFContextEndPage ends the current page in the PDF graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/endPDFPage()
func CGPDFContextEndPage(context CGContextRef) {
	if callErr := tryCGPDFContextEndPage(context); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFContextEndTag func(context CGContextRef)
var _cGPDFContextEndTagErr error

func tryCGPDFContextEndTag(context CGContextRef) error {
	if _cGPDFContextEndTag == nil {
		return symbolCallError("CGPDFContextEndTag", "10.15", _cGPDFContextEndTagErr)
	}
	_cGPDFContextEndTag(context)
	return nil
}

// CGPDFContextEndTag.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContextEndTag(_:)
func CGPDFContextEndTag(context CGContextRef) {
	if callErr := tryCGPDFContextEndTag(context); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFContextSetDestinationForRect func(context CGContextRef, name corefoundation.CFStringRef, rect corefoundation.CGRect)
var _cGPDFContextSetDestinationForRectErr error

func tryCGPDFContextSetDestinationForRect(context CGContextRef, name corefoundation.CFStringRef, rect corefoundation.CGRect) error {
	if _cGPDFContextSetDestinationForRect == nil {
		return symbolCallError("CGPDFContextSetDestinationForRect", "10.4", _cGPDFContextSetDestinationForRectErr)
	}
	_cGPDFContextSetDestinationForRect(context, name, rect)
	return nil
}

// CGPDFContextSetDestinationForRect sets a destination to jump to when a rectangle in the current PDF page is clicked.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setDestination(_:for:)
func CGPDFContextSetDestinationForRect(context CGContextRef, name corefoundation.CFStringRef, rect corefoundation.CGRect) {
	if callErr := tryCGPDFContextSetDestinationForRect(context, name, rect); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFContextSetIDTree func(context CGContextRef, IDTreeDictionary CGPDFDictionaryRef)
var _cGPDFContextSetIDTreeErr error

func tryCGPDFContextSetIDTree(context CGContextRef, IDTreeDictionary CGPDFDictionaryRef) error {
	if _cGPDFContextSetIDTree == nil {
		return symbolCallError("CGPDFContextSetIDTree", "", _cGPDFContextSetIDTreeErr)
	}
	_cGPDFContextSetIDTree(context, IDTreeDictionary)
	return nil
}

// CGPDFContextSetIDTree.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContextSetIDTree(_:_:)
func CGPDFContextSetIDTree(context CGContextRef, IDTreeDictionary CGPDFDictionaryRef) {
	if callErr := tryCGPDFContextSetIDTree(context, IDTreeDictionary); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFContextSetOutline func(context CGContextRef, outline corefoundation.CFDictionaryRef)
var _cGPDFContextSetOutlineErr error

func tryCGPDFContextSetOutline(context CGContextRef, outline corefoundation.CFDictionaryRef) error {
	if _cGPDFContextSetOutline == nil {
		return symbolCallError("CGPDFContextSetOutline", "10.13", _cGPDFContextSetOutlineErr)
	}
	_cGPDFContextSetOutline(context, outline)
	return nil
}

// CGPDFContextSetOutline.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContextSetOutline(_:_:)
func CGPDFContextSetOutline(context CGContextRef, outline corefoundation.CFDictionaryRef) {
	if callErr := tryCGPDFContextSetOutline(context, outline); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFContextSetPageTagStructureTree func(context CGContextRef, pageTagStructureTreeDictionary corefoundation.CFDictionaryRef)
var _cGPDFContextSetPageTagStructureTreeErr error

func tryCGPDFContextSetPageTagStructureTree(context CGContextRef, pageTagStructureTreeDictionary corefoundation.CFDictionaryRef) error {
	if _cGPDFContextSetPageTagStructureTree == nil {
		return symbolCallError("CGPDFContextSetPageTagStructureTree", "", _cGPDFContextSetPageTagStructureTreeErr)
	}
	_cGPDFContextSetPageTagStructureTree(context, pageTagStructureTreeDictionary)
	return nil
}

// CGPDFContextSetPageTagStructureTree.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContextSetPageTagStructureTree(_:_:)
func CGPDFContextSetPageTagStructureTree(context CGContextRef, pageTagStructureTreeDictionary corefoundation.CFDictionaryRef) {
	if callErr := tryCGPDFContextSetPageTagStructureTree(context, pageTagStructureTreeDictionary); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFContextSetParentTree func(context CGContextRef, parentTreeDictionary CGPDFDictionaryRef)
var _cGPDFContextSetParentTreeErr error

func tryCGPDFContextSetParentTree(context CGContextRef, parentTreeDictionary CGPDFDictionaryRef) error {
	if _cGPDFContextSetParentTree == nil {
		return symbolCallError("CGPDFContextSetParentTree", "", _cGPDFContextSetParentTreeErr)
	}
	_cGPDFContextSetParentTree(context, parentTreeDictionary)
	return nil
}

// CGPDFContextSetParentTree.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFContextSetParentTree(_:_:)
func CGPDFContextSetParentTree(context CGContextRef, parentTreeDictionary CGPDFDictionaryRef) {
	if callErr := tryCGPDFContextSetParentTree(context, parentTreeDictionary); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFContextSetURLForRect func(context CGContextRef, url corefoundation.CFURLRef, rect corefoundation.CGRect)
var _cGPDFContextSetURLForRectErr error

func tryCGPDFContextSetURLForRect(context CGContextRef, url corefoundation.CFURLRef, rect corefoundation.CGRect) error {
	if _cGPDFContextSetURLForRect == nil {
		return symbolCallError("CGPDFContextSetURLForRect", "10.4", _cGPDFContextSetURLForRectErr)
	}
	_cGPDFContextSetURLForRect(context, url, rect)
	return nil
}

// CGPDFContextSetURLForRect sets the URL associated with a rectangle in a PDF graphics context.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGContext/setURL(_:for:)
func CGPDFContextSetURLForRect(context CGContextRef, url corefoundation.CFURLRef, rect corefoundation.CGRect) {
	if callErr := tryCGPDFContextSetURLForRect(context, url, rect); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFDictionaryApplyBlock func(dict CGPDFDictionaryRef, block unsafe.Pointer, info unsafe.Pointer)
var _cGPDFDictionaryApplyBlockErr error

func tryCGPDFDictionaryApplyBlock(dict CGPDFDictionaryRef, block CGPDFDictionaryApplierBlock, info unsafe.Pointer) error {
	if _cGPDFDictionaryApplyBlock == nil {
		return symbolCallError("CGPDFDictionaryApplyBlock", "10.14", _cGPDFDictionaryApplyBlockErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 *byte, arg1 *CGPDFObjectRef, arg2 unsafe.Pointer) bool {
		return block(objc.GoString(arg0), arg1, arg2)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_cGPDFDictionaryApplyBlock(dict, _block0, info)
	return nil
}

// CGPDFDictionaryApplyBlock.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryApplyBlock(_:_:_:)
func CGPDFDictionaryApplyBlock(dict CGPDFDictionaryRef, block CGPDFDictionaryApplierBlock, info unsafe.Pointer) {
	if callErr := tryCGPDFDictionaryApplyBlock(dict, block, info); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFDictionaryApplyFunction func(dict CGPDFDictionaryRef, function CGPDFDictionaryApplierFunction, info unsafe.Pointer)
var _cGPDFDictionaryApplyFunctionErr error

func tryCGPDFDictionaryApplyFunction(dict CGPDFDictionaryRef, function CGPDFDictionaryApplierFunction, info unsafe.Pointer) error {
	if _cGPDFDictionaryApplyFunction == nil {
		return symbolCallError("CGPDFDictionaryApplyFunction", "10.3", _cGPDFDictionaryApplyFunctionErr)
	}
	_cGPDFDictionaryApplyFunction(dict, function, info)
	return nil
}

// CGPDFDictionaryApplyFunction applies a function to each entry in a dictionary.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryApplyFunction(_:_:_:)
func CGPDFDictionaryApplyFunction(dict CGPDFDictionaryRef, function CGPDFDictionaryApplierFunction, info unsafe.Pointer) {
	if callErr := tryCGPDFDictionaryApplyFunction(dict, function, info); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFDictionaryGetArray func(dict CGPDFDictionaryRef, key string, value *CGPDFArrayRef) bool
var _cGPDFDictionaryGetArrayErr error

func tryCGPDFDictionaryGetArray(dict CGPDFDictionaryRef, key string, value *CGPDFArrayRef) (bool, error) {
	if _cGPDFDictionaryGetArray == nil {
		return false, symbolCallError("CGPDFDictionaryGetArray", "10.3", _cGPDFDictionaryGetArrayErr)
	}
	return _cGPDFDictionaryGetArray(dict, key, value), nil
}

// CGPDFDictionaryGetArray returns whether there is a PDF array associated with a specified key in a PDF dictionary and, if so, retrieves that array.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetArray(_:_:_:)
func CGPDFDictionaryGetArray(dict CGPDFDictionaryRef, key string, value *CGPDFArrayRef) bool {
	result, callErr := tryCGPDFDictionaryGetArray(dict, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDictionaryGetBoolean func(dict CGPDFDictionaryRef, key string, value *CGPDFBoolean) bool
var _cGPDFDictionaryGetBooleanErr error

func tryCGPDFDictionaryGetBoolean(dict CGPDFDictionaryRef, key string, value *CGPDFBoolean) (bool, error) {
	if _cGPDFDictionaryGetBoolean == nil {
		return false, symbolCallError("CGPDFDictionaryGetBoolean", "10.3", _cGPDFDictionaryGetBooleanErr)
	}
	return _cGPDFDictionaryGetBoolean(dict, key, value), nil
}

// CGPDFDictionaryGetBoolean returns whether there is a PDF Boolean value associated with a specified key in a PDF dictionary and, if so, retrieves the Boolean value.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetBoolean(_:_:_:)
func CGPDFDictionaryGetBoolean(dict CGPDFDictionaryRef, key string, value *CGPDFBoolean) bool {
	result, callErr := tryCGPDFDictionaryGetBoolean(dict, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDictionaryGetCount func(dict CGPDFDictionaryRef) uintptr
var _cGPDFDictionaryGetCountErr error

func tryCGPDFDictionaryGetCount(dict CGPDFDictionaryRef) (uintptr, error) {
	if _cGPDFDictionaryGetCount == nil {
		return 0, symbolCallError("CGPDFDictionaryGetCount", "10.3", _cGPDFDictionaryGetCountErr)
	}
	return _cGPDFDictionaryGetCount(dict), nil
}

// CGPDFDictionaryGetCount returns the number of entries in a PDF dictionary.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetCount(_:)
func CGPDFDictionaryGetCount(dict CGPDFDictionaryRef) uintptr {
	result, callErr := tryCGPDFDictionaryGetCount(dict)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDictionaryGetDictionary func(dict CGPDFDictionaryRef, key string, value *CGPDFDictionaryRef) bool
var _cGPDFDictionaryGetDictionaryErr error

func tryCGPDFDictionaryGetDictionary(dict CGPDFDictionaryRef, key string, value *CGPDFDictionaryRef) (bool, error) {
	if _cGPDFDictionaryGetDictionary == nil {
		return false, symbolCallError("CGPDFDictionaryGetDictionary", "10.3", _cGPDFDictionaryGetDictionaryErr)
	}
	return _cGPDFDictionaryGetDictionary(dict, key, value), nil
}

// CGPDFDictionaryGetDictionary returns whether there is another PDF dictionary associated with a specified key in a PDF dictionary and, if so, retrieves that dictionary.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetDictionary(_:_:_:)
func CGPDFDictionaryGetDictionary(dict CGPDFDictionaryRef, key string, value *CGPDFDictionaryRef) bool {
	result, callErr := tryCGPDFDictionaryGetDictionary(dict, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDictionaryGetInteger func(dict CGPDFDictionaryRef, key string, value *CGPDFInteger) bool
var _cGPDFDictionaryGetIntegerErr error

func tryCGPDFDictionaryGetInteger(dict CGPDFDictionaryRef, key string, value *CGPDFInteger) (bool, error) {
	if _cGPDFDictionaryGetInteger == nil {
		return false, symbolCallError("CGPDFDictionaryGetInteger", "10.3", _cGPDFDictionaryGetIntegerErr)
	}
	return _cGPDFDictionaryGetInteger(dict, key, value), nil
}

// CGPDFDictionaryGetInteger returns whether there is a PDF integer associated with a specified key in a PDF dictionary and, if so, retrieves that integer.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetInteger(_:_:_:)
func CGPDFDictionaryGetInteger(dict CGPDFDictionaryRef, key string, value *CGPDFInteger) bool {
	result, callErr := tryCGPDFDictionaryGetInteger(dict, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDictionaryGetName func(dict CGPDFDictionaryRef, key string, value string) bool
var _cGPDFDictionaryGetNameErr error

func tryCGPDFDictionaryGetName(dict CGPDFDictionaryRef, key string, value string) (bool, error) {
	if _cGPDFDictionaryGetName == nil {
		return false, symbolCallError("CGPDFDictionaryGetName", "10.3", _cGPDFDictionaryGetNameErr)
	}
	return _cGPDFDictionaryGetName(dict, key, value), nil
}

// CGPDFDictionaryGetName returns whether an object with a specified key in a PDF dictionary is a PDF name reference (represented as a constant C string) and, if so, retrieves that name.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetName(_:_:_:)
func CGPDFDictionaryGetName(dict CGPDFDictionaryRef, key string, value string) bool {
	result, callErr := tryCGPDFDictionaryGetName(dict, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDictionaryGetNumber func(dict CGPDFDictionaryRef, key string, value *CGPDFReal) bool
var _cGPDFDictionaryGetNumberErr error

func tryCGPDFDictionaryGetNumber(dict CGPDFDictionaryRef, key string, value *CGPDFReal) (bool, error) {
	if _cGPDFDictionaryGetNumber == nil {
		return false, symbolCallError("CGPDFDictionaryGetNumber", "10.3", _cGPDFDictionaryGetNumberErr)
	}
	return _cGPDFDictionaryGetNumber(dict, key, value), nil
}

// CGPDFDictionaryGetNumber returns whether there is a PDF number associated with a specified key in a PDF dictionary and, if so, retrieves that number.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetNumber(_:_:_:)
func CGPDFDictionaryGetNumber(dict CGPDFDictionaryRef, key string, value *CGPDFReal) bool {
	result, callErr := tryCGPDFDictionaryGetNumber(dict, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDictionaryGetObject func(dict CGPDFDictionaryRef, key string, value *CGPDFObjectRef) bool
var _cGPDFDictionaryGetObjectErr error

func tryCGPDFDictionaryGetObject(dict CGPDFDictionaryRef, key string, value *CGPDFObjectRef) (bool, error) {
	if _cGPDFDictionaryGetObject == nil {
		return false, symbolCallError("CGPDFDictionaryGetObject", "10.3", _cGPDFDictionaryGetObjectErr)
	}
	return _cGPDFDictionaryGetObject(dict, key, value), nil
}

// CGPDFDictionaryGetObject returns whether there is a PDF object associated with a specified key in a PDF dictionary and, if so, retrieves that object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetObject(_:_:_:)
func CGPDFDictionaryGetObject(dict CGPDFDictionaryRef, key string, value *CGPDFObjectRef) bool {
	result, callErr := tryCGPDFDictionaryGetObject(dict, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDictionaryGetStream func(dict CGPDFDictionaryRef, key string, value *CGPDFStreamRef) bool
var _cGPDFDictionaryGetStreamErr error

func tryCGPDFDictionaryGetStream(dict CGPDFDictionaryRef, key string, value *CGPDFStreamRef) (bool, error) {
	if _cGPDFDictionaryGetStream == nil {
		return false, symbolCallError("CGPDFDictionaryGetStream", "10.3", _cGPDFDictionaryGetStreamErr)
	}
	return _cGPDFDictionaryGetStream(dict, key, value), nil
}

// CGPDFDictionaryGetStream returns whether there is a PDF stream associated with a specified key in a PDF dictionary and, if so, retrieves that stream.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetStream(_:_:_:)
func CGPDFDictionaryGetStream(dict CGPDFDictionaryRef, key string, value *CGPDFStreamRef) bool {
	result, callErr := tryCGPDFDictionaryGetStream(dict, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDictionaryGetString func(dict CGPDFDictionaryRef, key string, value *CGPDFStringRef) bool
var _cGPDFDictionaryGetStringErr error

func tryCGPDFDictionaryGetString(dict CGPDFDictionaryRef, key string, value *CGPDFStringRef) (bool, error) {
	if _cGPDFDictionaryGetString == nil {
		return false, symbolCallError("CGPDFDictionaryGetString", "10.3", _cGPDFDictionaryGetStringErr)
	}
	return _cGPDFDictionaryGetString(dict, key, value), nil
}

// CGPDFDictionaryGetString returns whether there is a PDF string associated with a specified key in a PDF dictionary and, if so, retrieves that string.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDictionaryGetString(_:_:_:)
func CGPDFDictionaryGetString(dict CGPDFDictionaryRef, key string, value *CGPDFStringRef) bool {
	result, callErr := tryCGPDFDictionaryGetString(dict, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentAllowsCopying func(document CGPDFDocumentRef) bool
var _cGPDFDocumentAllowsCopyingErr error

func tryCGPDFDocumentAllowsCopying(document CGPDFDocumentRef) (bool, error) {
	if _cGPDFDocumentAllowsCopying == nil {
		return false, symbolCallError("CGPDFDocumentAllowsCopying", "10.2", _cGPDFDocumentAllowsCopyingErr)
	}
	return _cGPDFDocumentAllowsCopying(document), nil
}

// CGPDFDocumentAllowsCopying returns whether the specified PDF document allows copying.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/allowsCopying
func CGPDFDocumentAllowsCopying(document CGPDFDocumentRef) bool {
	result, callErr := tryCGPDFDocumentAllowsCopying(document)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentAllowsPrinting func(document CGPDFDocumentRef) bool
var _cGPDFDocumentAllowsPrintingErr error

func tryCGPDFDocumentAllowsPrinting(document CGPDFDocumentRef) (bool, error) {
	if _cGPDFDocumentAllowsPrinting == nil {
		return false, symbolCallError("CGPDFDocumentAllowsPrinting", "10.2", _cGPDFDocumentAllowsPrintingErr)
	}
	return _cGPDFDocumentAllowsPrinting(document), nil
}

// CGPDFDocumentAllowsPrinting returns whether a PDF document allows printing.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/allowsPrinting
func CGPDFDocumentAllowsPrinting(document CGPDFDocumentRef) bool {
	result, callErr := tryCGPDFDocumentAllowsPrinting(document)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentCreateWithProvider func(provider CGDataProviderRef) CGPDFDocumentRef
var _cGPDFDocumentCreateWithProviderErr error

func tryCGPDFDocumentCreateWithProvider(provider CGDataProviderRef) (CGPDFDocumentRef, error) {
	if _cGPDFDocumentCreateWithProvider == nil {
		return 0, symbolCallError("CGPDFDocumentCreateWithProvider", "10.0", _cGPDFDocumentCreateWithProviderErr)
	}
	return _cGPDFDocumentCreateWithProvider(provider), nil
}

// CGPDFDocumentCreateWithProvider creates a Core Graphics PDF document using a data provider.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/init(_:)-gbq6
func CGPDFDocumentCreateWithProvider(provider CGDataProviderRef) CGPDFDocumentRef {
	result, callErr := tryCGPDFDocumentCreateWithProvider(provider)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentCreateWithURL func(url corefoundation.CFURLRef) CGPDFDocumentRef
var _cGPDFDocumentCreateWithURLErr error

func tryCGPDFDocumentCreateWithURL(url corefoundation.CFURLRef) (CGPDFDocumentRef, error) {
	if _cGPDFDocumentCreateWithURL == nil {
		return 0, symbolCallError("CGPDFDocumentCreateWithURL", "10.0", _cGPDFDocumentCreateWithURLErr)
	}
	return _cGPDFDocumentCreateWithURL(url), nil
}

// CGPDFDocumentCreateWithURL creates a Core Graphics PDF document using data specified by a URL.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/init(_:)-2gtsd
func CGPDFDocumentCreateWithURL(url corefoundation.CFURLRef) CGPDFDocumentRef {
	result, callErr := tryCGPDFDocumentCreateWithURL(url)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentGetAccessPermissions func(document CGPDFDocumentRef) CGPDFAccessPermissions
var _cGPDFDocumentGetAccessPermissionsErr error

func tryCGPDFDocumentGetAccessPermissions(document CGPDFDocumentRef) (CGPDFAccessPermissions, error) {
	if _cGPDFDocumentGetAccessPermissions == nil {
		return *new(CGPDFAccessPermissions), symbolCallError("CGPDFDocumentGetAccessPermissions", "10.13", _cGPDFDocumentGetAccessPermissionsErr)
	}
	return _cGPDFDocumentGetAccessPermissions(document), nil
}

// CGPDFDocumentGetAccessPermissions.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/accessPermissions
func CGPDFDocumentGetAccessPermissions(document CGPDFDocumentRef) CGPDFAccessPermissions {
	result, callErr := tryCGPDFDocumentGetAccessPermissions(document)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentGetArtBox func(document CGPDFDocumentRef, page int) corefoundation.CGRect
var _cGPDFDocumentGetArtBoxErr error

func tryCGPDFDocumentGetArtBox(document CGPDFDocumentRef, page int) (corefoundation.CGRect, error) {
	if _cGPDFDocumentGetArtBox == nil {
		return corefoundation.CGRect{}, symbolCallError("CGPDFDocumentGetArtBox", "10.0", _cGPDFDocumentGetArtBoxErr)
	}
	return _cGPDFDocumentGetArtBox(document, page), nil
}

// CGPDFDocumentGetArtBox returns the art box of a page in a PDF document.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentGetArtBox
func CGPDFDocumentGetArtBox(document CGPDFDocumentRef, page int) corefoundation.CGRect {
	result, callErr := tryCGPDFDocumentGetArtBox(document, page)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentGetBleedBox func(document CGPDFDocumentRef, page int) corefoundation.CGRect
var _cGPDFDocumentGetBleedBoxErr error

func tryCGPDFDocumentGetBleedBox(document CGPDFDocumentRef, page int) (corefoundation.CGRect, error) {
	if _cGPDFDocumentGetBleedBox == nil {
		return corefoundation.CGRect{}, symbolCallError("CGPDFDocumentGetBleedBox", "10.0", _cGPDFDocumentGetBleedBoxErr)
	}
	return _cGPDFDocumentGetBleedBox(document, page), nil
}

// CGPDFDocumentGetBleedBox returns the bleed box of a page in a PDF document.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentGetBleedBox
func CGPDFDocumentGetBleedBox(document CGPDFDocumentRef, page int) corefoundation.CGRect {
	result, callErr := tryCGPDFDocumentGetBleedBox(document, page)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentGetCatalog func(document CGPDFDocumentRef) CGPDFDictionaryRef
var _cGPDFDocumentGetCatalogErr error

func tryCGPDFDocumentGetCatalog(document CGPDFDocumentRef) (CGPDFDictionaryRef, error) {
	if _cGPDFDocumentGetCatalog == nil {
		return 0, symbolCallError("CGPDFDocumentGetCatalog", "10.3", _cGPDFDocumentGetCatalogErr)
	}
	return _cGPDFDocumentGetCatalog(document), nil
}

// CGPDFDocumentGetCatalog returns the document catalog of a Core Graphics PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/catalog
func CGPDFDocumentGetCatalog(document CGPDFDocumentRef) CGPDFDictionaryRef {
	result, callErr := tryCGPDFDocumentGetCatalog(document)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentGetCropBox func(document CGPDFDocumentRef, page int) corefoundation.CGRect
var _cGPDFDocumentGetCropBoxErr error

func tryCGPDFDocumentGetCropBox(document CGPDFDocumentRef, page int) (corefoundation.CGRect, error) {
	if _cGPDFDocumentGetCropBox == nil {
		return corefoundation.CGRect{}, symbolCallError("CGPDFDocumentGetCropBox", "10.0", _cGPDFDocumentGetCropBoxErr)
	}
	return _cGPDFDocumentGetCropBox(document, page), nil
}

// CGPDFDocumentGetCropBox returns the crop box of a page in a PDF document.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentGetCropBox
func CGPDFDocumentGetCropBox(document CGPDFDocumentRef, page int) corefoundation.CGRect {
	result, callErr := tryCGPDFDocumentGetCropBox(document, page)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentGetID func(document CGPDFDocumentRef) CGPDFArrayRef
var _cGPDFDocumentGetIDErr error

func tryCGPDFDocumentGetID(document CGPDFDocumentRef) (CGPDFArrayRef, error) {
	if _cGPDFDocumentGetID == nil {
		return 0, symbolCallError("CGPDFDocumentGetID", "10.4", _cGPDFDocumentGetIDErr)
	}
	return _cGPDFDocumentGetID(document), nil
}

// CGPDFDocumentGetID gets the file identifier for a PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/fileIdentifier
func CGPDFDocumentGetID(document CGPDFDocumentRef) CGPDFArrayRef {
	result, callErr := tryCGPDFDocumentGetID(document)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentGetInfo func(document CGPDFDocumentRef) CGPDFDictionaryRef
var _cGPDFDocumentGetInfoErr error

func tryCGPDFDocumentGetInfo(document CGPDFDocumentRef) (CGPDFDictionaryRef, error) {
	if _cGPDFDocumentGetInfo == nil {
		return 0, symbolCallError("CGPDFDocumentGetInfo", "10.4", _cGPDFDocumentGetInfoErr)
	}
	return _cGPDFDocumentGetInfo(document), nil
}

// CGPDFDocumentGetInfo gets the information dictionary for a PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/info
func CGPDFDocumentGetInfo(document CGPDFDocumentRef) CGPDFDictionaryRef {
	result, callErr := tryCGPDFDocumentGetInfo(document)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentGetMediaBox func(document CGPDFDocumentRef, page int) corefoundation.CGRect
var _cGPDFDocumentGetMediaBoxErr error

func tryCGPDFDocumentGetMediaBox(document CGPDFDocumentRef, page int) (corefoundation.CGRect, error) {
	if _cGPDFDocumentGetMediaBox == nil {
		return corefoundation.CGRect{}, symbolCallError("CGPDFDocumentGetMediaBox", "10.0", _cGPDFDocumentGetMediaBoxErr)
	}
	return _cGPDFDocumentGetMediaBox(document, page), nil
}

// CGPDFDocumentGetMediaBox returns the media box of a page in a PDF document.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentGetMediaBox
func CGPDFDocumentGetMediaBox(document CGPDFDocumentRef, page int) corefoundation.CGRect {
	result, callErr := tryCGPDFDocumentGetMediaBox(document, page)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentGetNumberOfPages func(document CGPDFDocumentRef) uintptr
var _cGPDFDocumentGetNumberOfPagesErr error

func tryCGPDFDocumentGetNumberOfPages(document CGPDFDocumentRef) (uintptr, error) {
	if _cGPDFDocumentGetNumberOfPages == nil {
		return 0, symbolCallError("CGPDFDocumentGetNumberOfPages", "10.0", _cGPDFDocumentGetNumberOfPagesErr)
	}
	return _cGPDFDocumentGetNumberOfPages(document), nil
}

// CGPDFDocumentGetNumberOfPages returns the number of pages in a PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/numberOfPages
func CGPDFDocumentGetNumberOfPages(document CGPDFDocumentRef) uintptr {
	result, callErr := tryCGPDFDocumentGetNumberOfPages(document)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentGetOutline func(document CGPDFDocumentRef) corefoundation.CFDictionaryRef
var _cGPDFDocumentGetOutlineErr error

func tryCGPDFDocumentGetOutline(document CGPDFDocumentRef) (corefoundation.CFDictionaryRef, error) {
	if _cGPDFDocumentGetOutline == nil {
		return 0, symbolCallError("CGPDFDocumentGetOutline", "10.13", _cGPDFDocumentGetOutlineErr)
	}
	return _cGPDFDocumentGetOutline(document), nil
}

// CGPDFDocumentGetOutline.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/outline
func CGPDFDocumentGetOutline(document CGPDFDocumentRef) corefoundation.CFDictionaryRef {
	result, callErr := tryCGPDFDocumentGetOutline(document)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentGetPage func(document CGPDFDocumentRef, pageNumber uintptr) CGPDFPageRef
var _cGPDFDocumentGetPageErr error

func tryCGPDFDocumentGetPage(document CGPDFDocumentRef, pageNumber uintptr) (CGPDFPageRef, error) {
	if _cGPDFDocumentGetPage == nil {
		return 0, symbolCallError("CGPDFDocumentGetPage", "10.3", _cGPDFDocumentGetPageErr)
	}
	return _cGPDFDocumentGetPage(document, pageNumber), nil
}

// CGPDFDocumentGetPage returns a page from a Core Graphics PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/page(at:)
func CGPDFDocumentGetPage(document CGPDFDocumentRef, pageNumber uintptr) CGPDFPageRef {
	result, callErr := tryCGPDFDocumentGetPage(document, pageNumber)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentGetRotationAngle func(document CGPDFDocumentRef, page int) int
var _cGPDFDocumentGetRotationAngleErr error

func tryCGPDFDocumentGetRotationAngle(document CGPDFDocumentRef, page int) (int, error) {
	if _cGPDFDocumentGetRotationAngle == nil {
		return 0, symbolCallError("CGPDFDocumentGetRotationAngle", "10.0", _cGPDFDocumentGetRotationAngleErr)
	}
	return _cGPDFDocumentGetRotationAngle(document, page), nil
}

// CGPDFDocumentGetRotationAngle returns the rotation angle of a page in a PDF document.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentGetRotationAngle
func CGPDFDocumentGetRotationAngle(document CGPDFDocumentRef, page int) int {
	result, callErr := tryCGPDFDocumentGetRotationAngle(document, page)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentGetTrimBox func(document CGPDFDocumentRef, page int) corefoundation.CGRect
var _cGPDFDocumentGetTrimBoxErr error

func tryCGPDFDocumentGetTrimBox(document CGPDFDocumentRef, page int) (corefoundation.CGRect, error) {
	if _cGPDFDocumentGetTrimBox == nil {
		return corefoundation.CGRect{}, symbolCallError("CGPDFDocumentGetTrimBox", "10.0", _cGPDFDocumentGetTrimBoxErr)
	}
	return _cGPDFDocumentGetTrimBox(document, page), nil
}

// CGPDFDocumentGetTrimBox returns the trim box of a page in a PDF document.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentGetTrimBox
func CGPDFDocumentGetTrimBox(document CGPDFDocumentRef, page int) corefoundation.CGRect {
	result, callErr := tryCGPDFDocumentGetTrimBox(document, page)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentGetTypeID func() uint
var _cGPDFDocumentGetTypeIDErr error

func tryCGPDFDocumentGetTypeID() (uint, error) {
	if _cGPDFDocumentGetTypeID == nil {
		return 0, symbolCallError("CGPDFDocumentGetTypeID", "10.2", _cGPDFDocumentGetTypeIDErr)
	}
	return _cGPDFDocumentGetTypeID(), nil
}

// CGPDFDocumentGetTypeID returns the type identifier for Core Graphics PDF documents.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/typeID
func CGPDFDocumentGetTypeID() uint {
	result, callErr := tryCGPDFDocumentGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentGetVersion func(document CGPDFDocumentRef, majorVersion *int, minorVersion *int)
var _cGPDFDocumentGetVersionErr error

func tryCGPDFDocumentGetVersion(document CGPDFDocumentRef, majorVersion []int, minorVersion []int) error {
	if _cGPDFDocumentGetVersion == nil {
		return symbolCallError("CGPDFDocumentGetVersion", "10.3", _cGPDFDocumentGetVersionErr)
	}
	_cGPDFDocumentGetVersion(document, unsafe.SliceData(majorVersion), unsafe.SliceData(minorVersion))
	return nil
}

// CGPDFDocumentGetVersion returns the major and minor version numbers of a Core Graphics PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/getVersion(majorVersion:minorVersion:)
func CGPDFDocumentGetVersion(document CGPDFDocumentRef, majorVersion []int, minorVersion []int) {
	if callErr := tryCGPDFDocumentGetVersion(document, majorVersion, minorVersion); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFDocumentIsEncrypted func(document CGPDFDocumentRef) bool
var _cGPDFDocumentIsEncryptedErr error

func tryCGPDFDocumentIsEncrypted(document CGPDFDocumentRef) (bool, error) {
	if _cGPDFDocumentIsEncrypted == nil {
		return false, symbolCallError("CGPDFDocumentIsEncrypted", "10.2", _cGPDFDocumentIsEncryptedErr)
	}
	return _cGPDFDocumentIsEncrypted(document), nil
}

// CGPDFDocumentIsEncrypted returns whether the specified PDF file is encrypted.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/isEncrypted
func CGPDFDocumentIsEncrypted(document CGPDFDocumentRef) bool {
	result, callErr := tryCGPDFDocumentIsEncrypted(document)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentIsUnlocked func(document CGPDFDocumentRef) bool
var _cGPDFDocumentIsUnlockedErr error

func tryCGPDFDocumentIsUnlocked(document CGPDFDocumentRef) (bool, error) {
	if _cGPDFDocumentIsUnlocked == nil {
		return false, symbolCallError("CGPDFDocumentIsUnlocked", "10.2", _cGPDFDocumentIsUnlockedErr)
	}
	return _cGPDFDocumentIsUnlocked(document), nil
}

// CGPDFDocumentIsUnlocked returns whether the specified PDF document is currently unlocked.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/isUnlocked
func CGPDFDocumentIsUnlocked(document CGPDFDocumentRef) bool {
	result, callErr := tryCGPDFDocumentIsUnlocked(document)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentRelease func(document CGPDFDocumentRef)
var _cGPDFDocumentReleaseErr error

func tryCGPDFDocumentRelease(document CGPDFDocumentRef) error {
	if _cGPDFDocumentRelease == nil {
		return symbolCallError("CGPDFDocumentRelease", "10.0", _cGPDFDocumentReleaseErr)
	}
	_cGPDFDocumentRelease(document)
	return nil
}

// CGPDFDocumentRelease decrements the retain count of a PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentRelease
func CGPDFDocumentRelease(document CGPDFDocumentRef) {
	if callErr := tryCGPDFDocumentRelease(document); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFDocumentRetain func(document CGPDFDocumentRef) CGPDFDocumentRef
var _cGPDFDocumentRetainErr error

func tryCGPDFDocumentRetain(document CGPDFDocumentRef) (CGPDFDocumentRef, error) {
	if _cGPDFDocumentRetain == nil {
		return 0, symbolCallError("CGPDFDocumentRetain", "10.0", _cGPDFDocumentRetainErr)
	}
	return _cGPDFDocumentRetain(document), nil
}

// CGPDFDocumentRetain increments the retain count of a Core Graphics PDF document.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocumentRetain
func CGPDFDocumentRetain(document CGPDFDocumentRef) CGPDFDocumentRef {
	result, callErr := tryCGPDFDocumentRetain(document)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFDocumentUnlockWithPassword func(document CGPDFDocumentRef, password string) bool
var _cGPDFDocumentUnlockWithPasswordErr error

func tryCGPDFDocumentUnlockWithPassword(document CGPDFDocumentRef, password string) (bool, error) {
	if _cGPDFDocumentUnlockWithPassword == nil {
		return false, symbolCallError("CGPDFDocumentUnlockWithPassword", "10.2", _cGPDFDocumentUnlockWithPasswordErr)
	}
	return _cGPDFDocumentUnlockWithPassword(document, password), nil
}

// CGPDFDocumentUnlockWithPassword unlocks an encrypted PDF document when a valid password is supplied.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDocument/unlockWithPassword(_:)
func CGPDFDocumentUnlockWithPassword(document CGPDFDocumentRef, password string) bool {
	result, callErr := tryCGPDFDocumentUnlockWithPassword(document, password)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFObjectGetType func(object CGPDFObjectRef) CGPDFObjectType
var _cGPDFObjectGetTypeErr error

func tryCGPDFObjectGetType(object CGPDFObjectRef) (CGPDFObjectType, error) {
	if _cGPDFObjectGetType == nil {
		return *new(CGPDFObjectType), symbolCallError("CGPDFObjectGetType", "10.3", _cGPDFObjectGetTypeErr)
	}
	return _cGPDFObjectGetType(object), nil
}

// CGPDFObjectGetType returns the PDF type identifier of an object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectGetType(_:)
func CGPDFObjectGetType(object CGPDFObjectRef) CGPDFObjectType {
	result, callErr := tryCGPDFObjectGetType(object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFObjectGetValue func(object CGPDFObjectRef, type_ CGPDFObjectType, value unsafe.Pointer) bool
var _cGPDFObjectGetValueErr error

func tryCGPDFObjectGetValue(object CGPDFObjectRef, type_ CGPDFObjectType, value unsafe.Pointer) (bool, error) {
	if _cGPDFObjectGetValue == nil {
		return false, symbolCallError("CGPDFObjectGetValue", "10.3", _cGPDFObjectGetValueErr)
	}
	return _cGPDFObjectGetValue(object, type_, value), nil
}

// CGPDFObjectGetValue returns whether an object is of a given type and if it is, retrieves its value.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectGetValue(_:_:_:)
func CGPDFObjectGetValue(object CGPDFObjectRef, type_ CGPDFObjectType, value unsafe.Pointer) bool {
	result, callErr := tryCGPDFObjectGetValue(object, type_, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFOperatorTableCreate func() CGPDFOperatorTableRef
var _cGPDFOperatorTableCreateErr error

func tryCGPDFOperatorTableCreate() (CGPDFOperatorTableRef, error) {
	if _cGPDFOperatorTableCreate == nil {
		return 0, symbolCallError("CGPDFOperatorTableCreate", "10.4", _cGPDFOperatorTableCreateErr)
	}
	return _cGPDFOperatorTableCreate(), nil
}

// CGPDFOperatorTableCreate creates an empty PDF operator table.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFOperatorTableCreate()
func CGPDFOperatorTableCreate() CGPDFOperatorTableRef {
	result, callErr := tryCGPDFOperatorTableCreate()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFOperatorTableRelease func(table CGPDFOperatorTableRef)
var _cGPDFOperatorTableReleaseErr error

func tryCGPDFOperatorTableRelease(table CGPDFOperatorTableRef) error {
	if _cGPDFOperatorTableRelease == nil {
		return symbolCallError("CGPDFOperatorTableRelease", "10.4", _cGPDFOperatorTableReleaseErr)
	}
	_cGPDFOperatorTableRelease(table)
	return nil
}

// CGPDFOperatorTableRelease decrements the retain count of a CGPDFOperatorTable object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFOperatorTableRelease(_:)
func CGPDFOperatorTableRelease(table CGPDFOperatorTableRef) {
	if callErr := tryCGPDFOperatorTableRelease(table); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFOperatorTableRetain func(table CGPDFOperatorTableRef) CGPDFOperatorTableRef
var _cGPDFOperatorTableRetainErr error

func tryCGPDFOperatorTableRetain(table CGPDFOperatorTableRef) (CGPDFOperatorTableRef, error) {
	if _cGPDFOperatorTableRetain == nil {
		return 0, symbolCallError("CGPDFOperatorTableRetain", "10.4", _cGPDFOperatorTableRetainErr)
	}
	return _cGPDFOperatorTableRetain(table), nil
}

// CGPDFOperatorTableRetain increments the retain count of a CGPDFOperatorTable object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFOperatorTableRetain(_:)
func CGPDFOperatorTableRetain(table CGPDFOperatorTableRef) CGPDFOperatorTableRef {
	result, callErr := tryCGPDFOperatorTableRetain(table)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFOperatorTableSetCallback func(table CGPDFOperatorTableRef, name string, callback CGPDFOperatorCallback)
var _cGPDFOperatorTableSetCallbackErr error

func tryCGPDFOperatorTableSetCallback(table CGPDFOperatorTableRef, name string, callback CGPDFOperatorCallback) error {
	if _cGPDFOperatorTableSetCallback == nil {
		return symbolCallError("CGPDFOperatorTableSetCallback", "10.4", _cGPDFOperatorTableSetCallbackErr)
	}
	_cGPDFOperatorTableSetCallback(table, name, callback)
	return nil
}

// CGPDFOperatorTableSetCallback sets a callback function for a PDF operator.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFOperatorTableSetCallback(_:_:_:)
func CGPDFOperatorTableSetCallback(table CGPDFOperatorTableRef, name string, callback CGPDFOperatorCallback) {
	if callErr := tryCGPDFOperatorTableSetCallback(table, name, callback); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFPageGetBoxRect func(page CGPDFPageRef, box CGPDFBox) corefoundation.CGRect
var _cGPDFPageGetBoxRectErr error

func tryCGPDFPageGetBoxRect(page CGPDFPageRef, box CGPDFBox) (corefoundation.CGRect, error) {
	if _cGPDFPageGetBoxRect == nil {
		return corefoundation.CGRect{}, symbolCallError("CGPDFPageGetBoxRect", "10.3", _cGPDFPageGetBoxRectErr)
	}
	return _cGPDFPageGetBoxRect(page, box), nil
}

// CGPDFPageGetBoxRect returns the rectangle that represents a type of box for a content region or page dimensions of a PDF page.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/getBoxRect(_:)
func CGPDFPageGetBoxRect(page CGPDFPageRef, box CGPDFBox) corefoundation.CGRect {
	result, callErr := tryCGPDFPageGetBoxRect(page, box)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFPageGetDictionary func(page CGPDFPageRef) CGPDFDictionaryRef
var _cGPDFPageGetDictionaryErr error

func tryCGPDFPageGetDictionary(page CGPDFPageRef) (CGPDFDictionaryRef, error) {
	if _cGPDFPageGetDictionary == nil {
		return 0, symbolCallError("CGPDFPageGetDictionary", "10.3", _cGPDFPageGetDictionaryErr)
	}
	return _cGPDFPageGetDictionary(page), nil
}

// CGPDFPageGetDictionary returns the dictionary of a PDF page.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/dictionary
func CGPDFPageGetDictionary(page CGPDFPageRef) CGPDFDictionaryRef {
	result, callErr := tryCGPDFPageGetDictionary(page)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFPageGetDocument func(page CGPDFPageRef) CGPDFDocumentRef
var _cGPDFPageGetDocumentErr error

func tryCGPDFPageGetDocument(page CGPDFPageRef) (CGPDFDocumentRef, error) {
	if _cGPDFPageGetDocument == nil {
		return 0, symbolCallError("CGPDFPageGetDocument", "10.3", _cGPDFPageGetDocumentErr)
	}
	return _cGPDFPageGetDocument(page), nil
}

// CGPDFPageGetDocument returns the document for a page.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/document
func CGPDFPageGetDocument(page CGPDFPageRef) CGPDFDocumentRef {
	result, callErr := tryCGPDFPageGetDocument(page)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFPageGetDrawingTransform func(page CGPDFPageRef, box CGPDFBox, rect corefoundation.CGRect, rotate int, preserveAspectRatio bool) corefoundation.CGAffineTransform
var _cGPDFPageGetDrawingTransformErr error

func tryCGPDFPageGetDrawingTransform(page CGPDFPageRef, box CGPDFBox, rect corefoundation.CGRect, rotate int, preserveAspectRatio bool) (corefoundation.CGAffineTransform, error) {
	if _cGPDFPageGetDrawingTransform == nil {
		return corefoundation.CGAffineTransform{}, symbolCallError("CGPDFPageGetDrawingTransform", "10.3", _cGPDFPageGetDrawingTransformErr)
	}
	return _cGPDFPageGetDrawingTransform(page, box, rect, rotate, preserveAspectRatio), nil
}

// CGPDFPageGetDrawingTransform returns the affine transform that maps a box to a given rectangle on a PDF page.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/getDrawingTransform(_:rect:rotate:preserveAspectRatio:)
func CGPDFPageGetDrawingTransform(page CGPDFPageRef, box CGPDFBox, rect corefoundation.CGRect, rotate int, preserveAspectRatio bool) corefoundation.CGAffineTransform {
	result, callErr := tryCGPDFPageGetDrawingTransform(page, box, rect, rotate, preserveAspectRatio)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFPageGetPageNumber func(page CGPDFPageRef) uintptr
var _cGPDFPageGetPageNumberErr error

func tryCGPDFPageGetPageNumber(page CGPDFPageRef) (uintptr, error) {
	if _cGPDFPageGetPageNumber == nil {
		return 0, symbolCallError("CGPDFPageGetPageNumber", "10.3", _cGPDFPageGetPageNumberErr)
	}
	return _cGPDFPageGetPageNumber(page), nil
}

// CGPDFPageGetPageNumber returns the page number of the specified PDF page.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/pageNumber
func CGPDFPageGetPageNumber(page CGPDFPageRef) uintptr {
	result, callErr := tryCGPDFPageGetPageNumber(page)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFPageGetRotationAngle func(page CGPDFPageRef) int
var _cGPDFPageGetRotationAngleErr error

func tryCGPDFPageGetRotationAngle(page CGPDFPageRef) (int, error) {
	if _cGPDFPageGetRotationAngle == nil {
		return 0, symbolCallError("CGPDFPageGetRotationAngle", "10.3", _cGPDFPageGetRotationAngleErr)
	}
	return _cGPDFPageGetRotationAngle(page), nil
}

// CGPDFPageGetRotationAngle returns the rotation angle of a PDF page, in degrees.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/rotationAngle
func CGPDFPageGetRotationAngle(page CGPDFPageRef) int {
	result, callErr := tryCGPDFPageGetRotationAngle(page)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFPageGetTypeID func() uint
var _cGPDFPageGetTypeIDErr error

func tryCGPDFPageGetTypeID() (uint, error) {
	if _cGPDFPageGetTypeID == nil {
		return 0, symbolCallError("CGPDFPageGetTypeID", "10.3", _cGPDFPageGetTypeIDErr)
	}
	return _cGPDFPageGetTypeID(), nil
}

// CGPDFPageGetTypeID returns the CFType ID for PDF page objects.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFPage/typeID
func CGPDFPageGetTypeID() uint {
	result, callErr := tryCGPDFPageGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFPageRelease func(page CGPDFPageRef)
var _cGPDFPageReleaseErr error

func tryCGPDFPageRelease(page CGPDFPageRef) error {
	if _cGPDFPageRelease == nil {
		return symbolCallError("CGPDFPageRelease", "10.3", _cGPDFPageReleaseErr)
	}
	_cGPDFPageRelease(page)
	return nil
}

// CGPDFPageRelease decrements the retain count of a PDF page.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFPageRelease
func CGPDFPageRelease(page CGPDFPageRef) {
	if callErr := tryCGPDFPageRelease(page); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFPageRetain func(page CGPDFPageRef) CGPDFPageRef
var _cGPDFPageRetainErr error

func tryCGPDFPageRetain(page CGPDFPageRef) (CGPDFPageRef, error) {
	if _cGPDFPageRetain == nil {
		return 0, symbolCallError("CGPDFPageRetain", "10.3", _cGPDFPageRetainErr)
	}
	return _cGPDFPageRetain(page), nil
}

// CGPDFPageRetain increments the retain count of a PDF page.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFPageRetain
func CGPDFPageRetain(page CGPDFPageRef) CGPDFPageRef {
	result, callErr := tryCGPDFPageRetain(page)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFScannerCreate func(cs CGPDFContentStreamRef, table CGPDFOperatorTableRef, info unsafe.Pointer) CGPDFScannerRef
var _cGPDFScannerCreateErr error

func tryCGPDFScannerCreate(cs CGPDFContentStreamRef, table CGPDFOperatorTableRef, info unsafe.Pointer) (CGPDFScannerRef, error) {
	if _cGPDFScannerCreate == nil {
		return 0, symbolCallError("CGPDFScannerCreate", "10.4", _cGPDFScannerCreateErr)
	}
	return _cGPDFScannerCreate(cs, table, info), nil
}

// CGPDFScannerCreate creates a PDF scanner.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerCreate(_:_:_:)
func CGPDFScannerCreate(cs CGPDFContentStreamRef, table CGPDFOperatorTableRef, info unsafe.Pointer) CGPDFScannerRef {
	result, callErr := tryCGPDFScannerCreate(cs, table, info)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFScannerGetContentStream func(scanner CGPDFScannerRef) CGPDFContentStreamRef
var _cGPDFScannerGetContentStreamErr error

func tryCGPDFScannerGetContentStream(scanner CGPDFScannerRef) (CGPDFContentStreamRef, error) {
	if _cGPDFScannerGetContentStream == nil {
		return 0, symbolCallError("CGPDFScannerGetContentStream", "10.4", _cGPDFScannerGetContentStreamErr)
	}
	return _cGPDFScannerGetContentStream(scanner), nil
}

// CGPDFScannerGetContentStream returns the content stream associated with a PDF scanner object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerGetContentStream(_:)
func CGPDFScannerGetContentStream(scanner CGPDFScannerRef) CGPDFContentStreamRef {
	result, callErr := tryCGPDFScannerGetContentStream(scanner)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFScannerPopArray func(scanner CGPDFScannerRef, value *CGPDFArrayRef) bool
var _cGPDFScannerPopArrayErr error

func tryCGPDFScannerPopArray(scanner CGPDFScannerRef, value *CGPDFArrayRef) (bool, error) {
	if _cGPDFScannerPopArray == nil {
		return false, symbolCallError("CGPDFScannerPopArray", "10.4", _cGPDFScannerPopArrayErr)
	}
	return _cGPDFScannerPopArray(scanner, value), nil
}

// CGPDFScannerPopArray retrieves an array object from the scanner stack.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopArray(_:_:)
func CGPDFScannerPopArray(scanner CGPDFScannerRef, value *CGPDFArrayRef) bool {
	result, callErr := tryCGPDFScannerPopArray(scanner, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFScannerPopBoolean func(scanner CGPDFScannerRef, value *CGPDFBoolean) bool
var _cGPDFScannerPopBooleanErr error

func tryCGPDFScannerPopBoolean(scanner CGPDFScannerRef, value *CGPDFBoolean) (bool, error) {
	if _cGPDFScannerPopBoolean == nil {
		return false, symbolCallError("CGPDFScannerPopBoolean", "10.4", _cGPDFScannerPopBooleanErr)
	}
	return _cGPDFScannerPopBoolean(scanner, value), nil
}

// CGPDFScannerPopBoolean retrieves a Boolean object from the scanner stack.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopBoolean(_:_:)
func CGPDFScannerPopBoolean(scanner CGPDFScannerRef, value *CGPDFBoolean) bool {
	result, callErr := tryCGPDFScannerPopBoolean(scanner, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFScannerPopDictionary func(scanner CGPDFScannerRef, value *CGPDFDictionaryRef) bool
var _cGPDFScannerPopDictionaryErr error

func tryCGPDFScannerPopDictionary(scanner CGPDFScannerRef, value *CGPDFDictionaryRef) (bool, error) {
	if _cGPDFScannerPopDictionary == nil {
		return false, symbolCallError("CGPDFScannerPopDictionary", "10.4", _cGPDFScannerPopDictionaryErr)
	}
	return _cGPDFScannerPopDictionary(scanner, value), nil
}

// CGPDFScannerPopDictionary retrieves a PDF dictionary object from the scanner stack.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopDictionary(_:_:)
func CGPDFScannerPopDictionary(scanner CGPDFScannerRef, value *CGPDFDictionaryRef) bool {
	result, callErr := tryCGPDFScannerPopDictionary(scanner, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFScannerPopInteger func(scanner CGPDFScannerRef, value *CGPDFInteger) bool
var _cGPDFScannerPopIntegerErr error

func tryCGPDFScannerPopInteger(scanner CGPDFScannerRef, value *CGPDFInteger) (bool, error) {
	if _cGPDFScannerPopInteger == nil {
		return false, symbolCallError("CGPDFScannerPopInteger", "10.4", _cGPDFScannerPopIntegerErr)
	}
	return _cGPDFScannerPopInteger(scanner, value), nil
}

// CGPDFScannerPopInteger retrieves an integer object from the scanner stack.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopInteger(_:_:)
func CGPDFScannerPopInteger(scanner CGPDFScannerRef, value *CGPDFInteger) bool {
	result, callErr := tryCGPDFScannerPopInteger(scanner, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFScannerPopName func(scanner CGPDFScannerRef, value string) bool
var _cGPDFScannerPopNameErr error

func tryCGPDFScannerPopName(scanner CGPDFScannerRef, value string) (bool, error) {
	if _cGPDFScannerPopName == nil {
		return false, symbolCallError("CGPDFScannerPopName", "10.4", _cGPDFScannerPopNameErr)
	}
	return _cGPDFScannerPopName(scanner, value), nil
}

// CGPDFScannerPopName retrieves a character string from the scanner stack.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopName(_:_:)
func CGPDFScannerPopName(scanner CGPDFScannerRef, value string) bool {
	result, callErr := tryCGPDFScannerPopName(scanner, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFScannerPopNumber func(scanner CGPDFScannerRef, value *CGPDFReal) bool
var _cGPDFScannerPopNumberErr error

func tryCGPDFScannerPopNumber(scanner CGPDFScannerRef, value *CGPDFReal) (bool, error) {
	if _cGPDFScannerPopNumber == nil {
		return false, symbolCallError("CGPDFScannerPopNumber", "10.4", _cGPDFScannerPopNumberErr)
	}
	return _cGPDFScannerPopNumber(scanner, value), nil
}

// CGPDFScannerPopNumber retrieves a real value object from the scanner stack.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopNumber(_:_:)
func CGPDFScannerPopNumber(scanner CGPDFScannerRef, value *CGPDFReal) bool {
	result, callErr := tryCGPDFScannerPopNumber(scanner, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFScannerPopObject func(scanner CGPDFScannerRef, value *CGPDFObjectRef) bool
var _cGPDFScannerPopObjectErr error

func tryCGPDFScannerPopObject(scanner CGPDFScannerRef, value *CGPDFObjectRef) (bool, error) {
	if _cGPDFScannerPopObject == nil {
		return false, symbolCallError("CGPDFScannerPopObject", "10.4", _cGPDFScannerPopObjectErr)
	}
	return _cGPDFScannerPopObject(scanner, value), nil
}

// CGPDFScannerPopObject retrieves an object from the scanner stack.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopObject(_:_:)
func CGPDFScannerPopObject(scanner CGPDFScannerRef, value *CGPDFObjectRef) bool {
	result, callErr := tryCGPDFScannerPopObject(scanner, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFScannerPopStream func(scanner CGPDFScannerRef, value *CGPDFStreamRef) bool
var _cGPDFScannerPopStreamErr error

func tryCGPDFScannerPopStream(scanner CGPDFScannerRef, value *CGPDFStreamRef) (bool, error) {
	if _cGPDFScannerPopStream == nil {
		return false, symbolCallError("CGPDFScannerPopStream", "10.4", _cGPDFScannerPopStreamErr)
	}
	return _cGPDFScannerPopStream(scanner, value), nil
}

// CGPDFScannerPopStream retrieves a PDF stream object from the scanner stack.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopStream(_:_:)
func CGPDFScannerPopStream(scanner CGPDFScannerRef, value *CGPDFStreamRef) bool {
	result, callErr := tryCGPDFScannerPopStream(scanner, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFScannerPopString func(scanner CGPDFScannerRef, value *CGPDFStringRef) bool
var _cGPDFScannerPopStringErr error

func tryCGPDFScannerPopString(scanner CGPDFScannerRef, value *CGPDFStringRef) (bool, error) {
	if _cGPDFScannerPopString == nil {
		return false, symbolCallError("CGPDFScannerPopString", "10.4", _cGPDFScannerPopStringErr)
	}
	return _cGPDFScannerPopString(scanner, value), nil
}

// CGPDFScannerPopString retrieves a string object from the scanner stack.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerPopString(_:_:)
func CGPDFScannerPopString(scanner CGPDFScannerRef, value *CGPDFStringRef) bool {
	result, callErr := tryCGPDFScannerPopString(scanner, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFScannerRelease func(scanner CGPDFScannerRef)
var _cGPDFScannerReleaseErr error

func tryCGPDFScannerRelease(scanner CGPDFScannerRef) error {
	if _cGPDFScannerRelease == nil {
		return symbolCallError("CGPDFScannerRelease", "10.4", _cGPDFScannerReleaseErr)
	}
	_cGPDFScannerRelease(scanner)
	return nil
}

// CGPDFScannerRelease decrements the retain count of a scanner object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerRelease(_:)
func CGPDFScannerRelease(scanner CGPDFScannerRef) {
	if callErr := tryCGPDFScannerRelease(scanner); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFScannerRetain func(scanner CGPDFScannerRef) CGPDFScannerRef
var _cGPDFScannerRetainErr error

func tryCGPDFScannerRetain(scanner CGPDFScannerRef) (CGPDFScannerRef, error) {
	if _cGPDFScannerRetain == nil {
		return 0, symbolCallError("CGPDFScannerRetain", "10.4", _cGPDFScannerRetainErr)
	}
	return _cGPDFScannerRetain(scanner), nil
}

// CGPDFScannerRetain increments the retain count of a scanner object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerRetain(_:)
func CGPDFScannerRetain(scanner CGPDFScannerRef) CGPDFScannerRef {
	result, callErr := tryCGPDFScannerRetain(scanner)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFScannerScan func(scanner CGPDFScannerRef) bool
var _cGPDFScannerScanErr error

func tryCGPDFScannerScan(scanner CGPDFScannerRef) (bool, error) {
	if _cGPDFScannerScan == nil {
		return false, symbolCallError("CGPDFScannerScan", "10.4", _cGPDFScannerScanErr)
	}
	return _cGPDFScannerScan(scanner), nil
}

// CGPDFScannerScan parses the content stream of a PDF scanner object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerScan(_:)
func CGPDFScannerScan(scanner CGPDFScannerRef) bool {
	result, callErr := tryCGPDFScannerScan(scanner)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFScannerStop func(s CGPDFScannerRef)
var _cGPDFScannerStopErr error

func tryCGPDFScannerStop(s CGPDFScannerRef) error {
	if _cGPDFScannerStop == nil {
		return symbolCallError("CGPDFScannerStop", "", _cGPDFScannerStopErr)
	}
	_cGPDFScannerStop(s)
	return nil
}

// CGPDFScannerStop.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFScannerStop(_:)
func CGPDFScannerStop(s CGPDFScannerRef) {
	if callErr := tryCGPDFScannerStop(s); callErr != nil {
		panic(callErr)
	}
}

var _cGPDFStreamCopyData func(stream CGPDFStreamRef, format *CGPDFDataFormat) corefoundation.CFDataRef
var _cGPDFStreamCopyDataErr error

func tryCGPDFStreamCopyData(stream CGPDFStreamRef, format *CGPDFDataFormat) (corefoundation.CFDataRef, error) {
	if _cGPDFStreamCopyData == nil {
		return 0, symbolCallError("CGPDFStreamCopyData", "10.3", _cGPDFStreamCopyDataErr)
	}
	return _cGPDFStreamCopyData(stream, format), nil
}

// CGPDFStreamCopyData returns the data associated with a PDF stream.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFStreamCopyData(_:_:)
func CGPDFStreamCopyData(stream CGPDFStreamRef, format *CGPDFDataFormat) corefoundation.CFDataRef {
	result, callErr := tryCGPDFStreamCopyData(stream, format)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFStreamGetDictionary func(stream CGPDFStreamRef) CGPDFDictionaryRef
var _cGPDFStreamGetDictionaryErr error

func tryCGPDFStreamGetDictionary(stream CGPDFStreamRef) (CGPDFDictionaryRef, error) {
	if _cGPDFStreamGetDictionary == nil {
		return 0, symbolCallError("CGPDFStreamGetDictionary", "10.3", _cGPDFStreamGetDictionaryErr)
	}
	return _cGPDFStreamGetDictionary(stream), nil
}

// CGPDFStreamGetDictionary returns the dictionary associated with a PDF stream.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFStreamGetDictionary(_:)
func CGPDFStreamGetDictionary(stream CGPDFStreamRef) CGPDFDictionaryRef {
	result, callErr := tryCGPDFStreamGetDictionary(stream)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFStringCopyDate func(string_ CGPDFStringRef) corefoundation.CFDateRef
var _cGPDFStringCopyDateErr error

func tryCGPDFStringCopyDate(string_ CGPDFStringRef) (corefoundation.CFDateRef, error) {
	if _cGPDFStringCopyDate == nil {
		return 0, symbolCallError("CGPDFStringCopyDate", "10.4", _cGPDFStringCopyDateErr)
	}
	return _cGPDFStringCopyDate(string_), nil
}

// CGPDFStringCopyDate converts a string to a date.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFStringCopyDate(_:)
func CGPDFStringCopyDate(string_ CGPDFStringRef) corefoundation.CFDateRef {
	result, callErr := tryCGPDFStringCopyDate(string_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFStringCopyTextString func(string_ CGPDFStringRef) corefoundation.CFStringRef
var _cGPDFStringCopyTextStringErr error

func tryCGPDFStringCopyTextString(string_ CGPDFStringRef) (corefoundation.CFStringRef, error) {
	if _cGPDFStringCopyTextString == nil {
		return 0, symbolCallError("CGPDFStringCopyTextString", "10.3", _cGPDFStringCopyTextStringErr)
	}
	return _cGPDFStringCopyTextString(string_), nil
}

// CGPDFStringCopyTextString returns a CFString object that represents a PDF string as a text string.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFStringCopyTextString(_:)
func CGPDFStringCopyTextString(string_ CGPDFStringRef) corefoundation.CFStringRef {
	result, callErr := tryCGPDFStringCopyTextString(string_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFStringGetBytePtr func(string_ CGPDFStringRef) *byte
var _cGPDFStringGetBytePtrErr error

func tryCGPDFStringGetBytePtr(string_ CGPDFStringRef) (*byte, error) {
	if _cGPDFStringGetBytePtr == nil {
		return nil, symbolCallError("CGPDFStringGetBytePtr", "10.3", _cGPDFStringGetBytePtrErr)
	}
	return _cGPDFStringGetBytePtr(string_), nil
}

// CGPDFStringGetBytePtr returns a pointer to the bytes of a PDF string.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFStringGetBytePtr(_:)
func CGPDFStringGetBytePtr(string_ CGPDFStringRef) *byte {
	result, callErr := tryCGPDFStringGetBytePtr(string_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFStringGetLength func(string_ CGPDFStringRef) uintptr
var _cGPDFStringGetLengthErr error

func tryCGPDFStringGetLength(string_ CGPDFStringRef) (uintptr, error) {
	if _cGPDFStringGetLength == nil {
		return 0, symbolCallError("CGPDFStringGetLength", "10.3", _cGPDFStringGetLengthErr)
	}
	return _cGPDFStringGetLength(string_), nil
}

// CGPDFStringGetLength returns the number of bytes in a PDF string.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFStringGetLength(_:)
func CGPDFStringGetLength(string_ CGPDFStringRef) uintptr {
	result, callErr := tryCGPDFStringGetLength(string_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPDFTagTypeGetName func(tagType CGPDFTagType) *byte
var _cGPDFTagTypeGetNameErr error

func tryCGPDFTagTypeGetName(tagType CGPDFTagType) (*byte, error) {
	if _cGPDFTagTypeGetName == nil {
		return nil, symbolCallError("CGPDFTagTypeGetName", "10.15", _cGPDFTagTypeGetNameErr)
	}
	return _cGPDFTagTypeGetName(tagType), nil
}

// CGPDFTagTypeGetName.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/name
func CGPDFTagTypeGetName(tagType CGPDFTagType) *byte {
	result, callErr := tryCGPDFTagTypeGetName(tagType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPSConverterAbort func(converter CGPSConverterRef) bool
var _cGPSConverterAbortErr error

func tryCGPSConverterAbort(converter CGPSConverterRef) (bool, error) {
	if _cGPSConverterAbort == nil {
		return false, symbolCallError("CGPSConverterAbort", "10.3", _cGPSConverterAbortErr)
	}
	return _cGPSConverterAbort(converter), nil
}

// CGPSConverterAbort tells a PostScript converter to abort a conversion at the next available opportunity.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPSConverter/abort()
func CGPSConverterAbort(converter CGPSConverterRef) bool {
	result, callErr := tryCGPSConverterAbort(converter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPSConverterConvert func(converter CGPSConverterRef, provider CGDataProviderRef, consumer CGDataConsumerRef, options corefoundation.CFDictionaryRef) bool
var _cGPSConverterConvertErr error

func tryCGPSConverterConvert(converter CGPSConverterRef, provider CGDataProviderRef, consumer CGDataConsumerRef, options corefoundation.CFDictionaryRef) (bool, error) {
	if _cGPSConverterConvert == nil {
		return false, symbolCallError("CGPSConverterConvert", "10.3", _cGPSConverterConvertErr)
	}
	return _cGPSConverterConvert(converter, provider, consumer, options), nil
}

// CGPSConverterConvert uses a PostScript converter to convert PostScript data to PDF data.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPSConverter/convert(_:consumer:options:)
func CGPSConverterConvert(converter CGPSConverterRef, provider CGDataProviderRef, consumer CGDataConsumerRef, options corefoundation.CFDictionaryRef) bool {
	result, callErr := tryCGPSConverterConvert(converter, provider, consumer, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPSConverterCreate func(info unsafe.Pointer, callbacks *CGPSConverterCallbacks, options corefoundation.CFDictionaryRef) CGPSConverterRef
var _cGPSConverterCreateErr error

func tryCGPSConverterCreate(info unsafe.Pointer, callbacks *CGPSConverterCallbacks, options corefoundation.CFDictionaryRef) (CGPSConverterRef, error) {
	if _cGPSConverterCreate == nil {
		return 0, symbolCallError("CGPSConverterCreate", "10.3", _cGPSConverterCreateErr)
	}
	return _cGPSConverterCreate(info, callbacks, options), nil
}

// CGPSConverterCreate creates a new PostScript converter.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPSConverter/init(info:callbacks:options:)
func CGPSConverterCreate(info unsafe.Pointer, callbacks *CGPSConverterCallbacks, options corefoundation.CFDictionaryRef) CGPSConverterRef {
	result, callErr := tryCGPSConverterCreate(info, callbacks, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPSConverterGetTypeID func() uint
var _cGPSConverterGetTypeIDErr error

func tryCGPSConverterGetTypeID() (uint, error) {
	if _cGPSConverterGetTypeID == nil {
		return 0, symbolCallError("CGPSConverterGetTypeID", "10.3", _cGPSConverterGetTypeIDErr)
	}
	return _cGPSConverterGetTypeID(), nil
}

// CGPSConverterGetTypeID returns the Core Foundation type identifier for PostScript converters.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPSConverter/typeID
func CGPSConverterGetTypeID() uint {
	result, callErr := tryCGPSConverterGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPSConverterIsConverting func(converter CGPSConverterRef) bool
var _cGPSConverterIsConvertingErr error

func tryCGPSConverterIsConverting(converter CGPSConverterRef) (bool, error) {
	if _cGPSConverterIsConverting == nil {
		return false, symbolCallError("CGPSConverterIsConverting", "10.3", _cGPSConverterIsConvertingErr)
	}
	return _cGPSConverterIsConverting(converter), nil
}

// CGPSConverterIsConverting checks whether the converter is currently converting data.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPSConverter/isConverting
func CGPSConverterIsConverting(converter CGPSConverterRef) bool {
	result, callErr := tryCGPSConverterIsConverting(converter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathAddArc func(path CGMutablePathRef, m *corefoundation.CGAffineTransform, x float64, y float64, radius float64, startAngle float64, endAngle float64, clockwise bool)
var _cGPathAddArcErr error

func tryCGPathAddArc(path CGMutablePathRef, m *corefoundation.CGAffineTransform, x float64, y float64, radius float64, startAngle float64, endAngle float64, clockwise bool) error {
	if _cGPathAddArc == nil {
		return symbolCallError("CGPathAddArc", "10.2", _cGPathAddArcErr)
	}
	_cGPathAddArc(path, m, x, y, radius, startAngle, endAngle, clockwise)
	return nil
}

// CGPathAddArc appends an arc to a mutable graphics path, possibly preceded by a straight line segment.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathAddArc
func CGPathAddArc(path CGMutablePathRef, m *corefoundation.CGAffineTransform, x float64, y float64, radius float64, startAngle float64, endAngle float64, clockwise bool) {
	if callErr := tryCGPathAddArc(path, m, x, y, radius, startAngle, endAngle, clockwise); callErr != nil {
		panic(callErr)
	}
}

var _cGPathAddArcToPoint func(path CGMutablePathRef, m *corefoundation.CGAffineTransform, x1 float64, y1 float64, x2 float64, y2 float64, radius float64)
var _cGPathAddArcToPointErr error

func tryCGPathAddArcToPoint(path CGMutablePathRef, m *corefoundation.CGAffineTransform, x1 float64, y1 float64, x2 float64, y2 float64, radius float64) error {
	if _cGPathAddArcToPoint == nil {
		return symbolCallError("CGPathAddArcToPoint", "10.2", _cGPathAddArcToPointErr)
	}
	_cGPathAddArcToPoint(path, m, x1, y1, x2, y2, radius)
	return nil
}

// CGPathAddArcToPoint appends an arc to a mutable graphics path, possibly preceded by a straight line segment.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathAddArcToPoint
func CGPathAddArcToPoint(path CGMutablePathRef, m *corefoundation.CGAffineTransform, x1 float64, y1 float64, x2 float64, y2 float64, radius float64) {
	if callErr := tryCGPathAddArcToPoint(path, m, x1, y1, x2, y2, radius); callErr != nil {
		panic(callErr)
	}
}

var _cGPathAddCurveToPoint func(path CGMutablePathRef, m *corefoundation.CGAffineTransform, cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64)
var _cGPathAddCurveToPointErr error

func tryCGPathAddCurveToPoint(path CGMutablePathRef, m *corefoundation.CGAffineTransform, cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) error {
	if _cGPathAddCurveToPoint == nil {
		return symbolCallError("CGPathAddCurveToPoint", "10.2", _cGPathAddCurveToPointErr)
	}
	_cGPathAddCurveToPoint(path, m, cp1x, cp1y, cp2x, cp2y, x, y)
	return nil
}

// CGPathAddCurveToPoint appends a cubic Bézier curve to a mutable graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathAddCurveToPoint
func CGPathAddCurveToPoint(path CGMutablePathRef, m *corefoundation.CGAffineTransform, cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) {
	if callErr := tryCGPathAddCurveToPoint(path, m, cp1x, cp1y, cp2x, cp2y, x, y); callErr != nil {
		panic(callErr)
	}
}

var _cGPathAddEllipseInRect func(path CGMutablePathRef, m *corefoundation.CGAffineTransform, rect corefoundation.CGRect)
var _cGPathAddEllipseInRectErr error

func tryCGPathAddEllipseInRect(path CGMutablePathRef, m *corefoundation.CGAffineTransform, rect corefoundation.CGRect) error {
	if _cGPathAddEllipseInRect == nil {
		return symbolCallError("CGPathAddEllipseInRect", "10.4", _cGPathAddEllipseInRectErr)
	}
	_cGPathAddEllipseInRect(path, m, rect)
	return nil
}

// CGPathAddEllipseInRect adds to a path an ellipse that fits inside a rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathAddEllipseInRect
func CGPathAddEllipseInRect(path CGMutablePathRef, m *corefoundation.CGAffineTransform, rect corefoundation.CGRect) {
	if callErr := tryCGPathAddEllipseInRect(path, m, rect); callErr != nil {
		panic(callErr)
	}
}

var _cGPathAddLineToPoint func(path CGMutablePathRef, m *corefoundation.CGAffineTransform, x float64, y float64)
var _cGPathAddLineToPointErr error

func tryCGPathAddLineToPoint(path CGMutablePathRef, m *corefoundation.CGAffineTransform, x float64, y float64) error {
	if _cGPathAddLineToPoint == nil {
		return symbolCallError("CGPathAddLineToPoint", "10.2", _cGPathAddLineToPointErr)
	}
	_cGPathAddLineToPoint(path, m, x, y)
	return nil
}

// CGPathAddLineToPoint appends a line segment to a mutable graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathAddLineToPoint
func CGPathAddLineToPoint(path CGMutablePathRef, m *corefoundation.CGAffineTransform, x float64, y float64) {
	if callErr := tryCGPathAddLineToPoint(path, m, x, y); callErr != nil {
		panic(callErr)
	}
}

var _cGPathAddLines func(path CGMutablePathRef, m *corefoundation.CGAffineTransform, points *corefoundation.CGPoint, count uintptr)
var _cGPathAddLinesErr error

func tryCGPathAddLines(path CGMutablePathRef, m *corefoundation.CGAffineTransform, points *corefoundation.CGPoint, count uintptr) error {
	if _cGPathAddLines == nil {
		return symbolCallError("CGPathAddLines", "10.2", _cGPathAddLinesErr)
	}
	_cGPathAddLines(path, m, points, count)
	return nil
}

// CGPathAddLines appends an array of new line segments to a mutable graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathAddLines
func CGPathAddLines(path CGMutablePathRef, m *corefoundation.CGAffineTransform, points *corefoundation.CGPoint, count uintptr) {
	if callErr := tryCGPathAddLines(path, m, points, count); callErr != nil {
		panic(callErr)
	}
}

var _cGPathAddPath func(path1 CGMutablePathRef, m *corefoundation.CGAffineTransform, path2 CGPathRef)
var _cGPathAddPathErr error

func tryCGPathAddPath(path1 CGMutablePathRef, m *corefoundation.CGAffineTransform, path2 CGPathRef) error {
	if _cGPathAddPath == nil {
		return symbolCallError("CGPathAddPath", "10.2", _cGPathAddPathErr)
	}
	_cGPathAddPath(path1, m, path2)
	return nil
}

// CGPathAddPath appends a path to onto a mutable graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathAddPath
func CGPathAddPath(path1 CGMutablePathRef, m *corefoundation.CGAffineTransform, path2 CGPathRef) {
	if callErr := tryCGPathAddPath(path1, m, path2); callErr != nil {
		panic(callErr)
	}
}

var _cGPathAddQuadCurveToPoint func(path CGMutablePathRef, m *corefoundation.CGAffineTransform, cpx float64, cpy float64, x float64, y float64)
var _cGPathAddQuadCurveToPointErr error

func tryCGPathAddQuadCurveToPoint(path CGMutablePathRef, m *corefoundation.CGAffineTransform, cpx float64, cpy float64, x float64, y float64) error {
	if _cGPathAddQuadCurveToPoint == nil {
		return symbolCallError("CGPathAddQuadCurveToPoint", "10.2", _cGPathAddQuadCurveToPointErr)
	}
	_cGPathAddQuadCurveToPoint(path, m, cpx, cpy, x, y)
	return nil
}

// CGPathAddQuadCurveToPoint appends a quadratic Bézier curve to a mutable graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathAddQuadCurveToPoint
func CGPathAddQuadCurveToPoint(path CGMutablePathRef, m *corefoundation.CGAffineTransform, cpx float64, cpy float64, x float64, y float64) {
	if callErr := tryCGPathAddQuadCurveToPoint(path, m, cpx, cpy, x, y); callErr != nil {
		panic(callErr)
	}
}

var _cGPathAddRect func(path CGMutablePathRef, m *corefoundation.CGAffineTransform, rect corefoundation.CGRect)
var _cGPathAddRectErr error

func tryCGPathAddRect(path CGMutablePathRef, m *corefoundation.CGAffineTransform, rect corefoundation.CGRect) error {
	if _cGPathAddRect == nil {
		return symbolCallError("CGPathAddRect", "10.2", _cGPathAddRectErr)
	}
	_cGPathAddRect(path, m, rect)
	return nil
}

// CGPathAddRect appends a rectangle to a mutable graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathAddRect
func CGPathAddRect(path CGMutablePathRef, m *corefoundation.CGAffineTransform, rect corefoundation.CGRect) {
	if callErr := tryCGPathAddRect(path, m, rect); callErr != nil {
		panic(callErr)
	}
}

var _cGPathAddRects func(path CGMutablePathRef, m *corefoundation.CGAffineTransform, rects *corefoundation.CGRect, count uintptr)
var _cGPathAddRectsErr error

func tryCGPathAddRects(path CGMutablePathRef, m *corefoundation.CGAffineTransform, rects *corefoundation.CGRect, count uintptr) error {
	if _cGPathAddRects == nil {
		return symbolCallError("CGPathAddRects", "10.2", _cGPathAddRectsErr)
	}
	_cGPathAddRects(path, m, rects, count)
	return nil
}

// CGPathAddRects appends an array of rectangles to a mutable graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathAddRects
func CGPathAddRects(path CGMutablePathRef, m *corefoundation.CGAffineTransform, rects *corefoundation.CGRect, count uintptr) {
	if callErr := tryCGPathAddRects(path, m, rects, count); callErr != nil {
		panic(callErr)
	}
}

var _cGPathAddRelativeArc func(path CGMutablePathRef, matrix *corefoundation.CGAffineTransform, x float64, y float64, radius float64, startAngle float64, delta float64)
var _cGPathAddRelativeArcErr error

func tryCGPathAddRelativeArc(path CGMutablePathRef, matrix *corefoundation.CGAffineTransform, x float64, y float64, radius float64, startAngle float64, delta float64) error {
	if _cGPathAddRelativeArc == nil {
		return symbolCallError("CGPathAddRelativeArc", "10.7", _cGPathAddRelativeArcErr)
	}
	_cGPathAddRelativeArc(path, matrix, x, y, radius, startAngle, delta)
	return nil
}

// CGPathAddRelativeArc appends an arc to a mutable graphics path, possibly preceded by a straight line segment.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathAddRelativeArc
func CGPathAddRelativeArc(path CGMutablePathRef, matrix *corefoundation.CGAffineTransform, x float64, y float64, radius float64, startAngle float64, delta float64) {
	if callErr := tryCGPathAddRelativeArc(path, matrix, x, y, radius, startAngle, delta); callErr != nil {
		panic(callErr)
	}
}

var _cGPathAddRoundedRect func(path CGMutablePathRef, transform *corefoundation.CGAffineTransform, rect corefoundation.CGRect, cornerWidth float64, cornerHeight float64)
var _cGPathAddRoundedRectErr error

func tryCGPathAddRoundedRect(path CGMutablePathRef, transform *corefoundation.CGAffineTransform, rect corefoundation.CGRect, cornerWidth float64, cornerHeight float64) error {
	if _cGPathAddRoundedRect == nil {
		return symbolCallError("CGPathAddRoundedRect", "10.9", _cGPathAddRoundedRectErr)
	}
	_cGPathAddRoundedRect(path, transform, rect, cornerWidth, cornerHeight)
	return nil
}

// CGPathAddRoundedRect appends a rounded rectangle to a mutable graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathAddRoundedRect
func CGPathAddRoundedRect(path CGMutablePathRef, transform *corefoundation.CGAffineTransform, rect corefoundation.CGRect, cornerWidth float64, cornerHeight float64) {
	if callErr := tryCGPathAddRoundedRect(path, transform, rect, cornerWidth, cornerHeight); callErr != nil {
		panic(callErr)
	}
}

var _cGPathApply func(path CGPathRef, info unsafe.Pointer, function CGPathApplierFunction)
var _cGPathApplyErr error

func tryCGPathApply(path CGPathRef, info unsafe.Pointer, function CGPathApplierFunction) error {
	if _cGPathApply == nil {
		return symbolCallError("CGPathApply", "10.2", _cGPathApplyErr)
	}
	_cGPathApply(path, info, function)
	return nil
}

// CGPathApply for each element in a graphics path, calls a custom applier function.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/apply(info:function:)
func CGPathApply(path CGPathRef, info unsafe.Pointer, function CGPathApplierFunction) {
	if callErr := tryCGPathApply(path, info, function); callErr != nil {
		panic(callErr)
	}
}

var _cGPathApplyWithBlock func(path CGPathRef, block unsafe.Pointer)
var _cGPathApplyWithBlockErr error

func tryCGPathApplyWithBlock(path CGPathRef, block CGPathApplyBlock) error {
	if _cGPathApplyWithBlock == nil {
		return symbolCallError("CGPathApplyWithBlock", "10.13", _cGPathApplyWithBlockErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 *CGPathElement) { block(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_cGPathApplyWithBlock(path, _block0)
	return nil
}

// CGPathApplyWithBlock.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/applyWithBlock(_:)
func CGPathApplyWithBlock(path CGPathRef, block CGPathApplyBlock) {
	if callErr := tryCGPathApplyWithBlock(path, block); callErr != nil {
		panic(callErr)
	}
}

var _cGPathCloseSubpath func(path CGMutablePathRef)
var _cGPathCloseSubpathErr error

func tryCGPathCloseSubpath(path CGMutablePathRef) error {
	if _cGPathCloseSubpath == nil {
		return symbolCallError("CGPathCloseSubpath", "10.2", _cGPathCloseSubpathErr)
	}
	_cGPathCloseSubpath(path)
	return nil
}

// CGPathCloseSubpath closes and completes a subpath in a mutable graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGMutablePath/closeSubpath()
func CGPathCloseSubpath(path CGMutablePathRef) {
	if callErr := tryCGPathCloseSubpath(path); callErr != nil {
		panic(callErr)
	}
}

var _cGPathContainsPoint func(path CGPathRef, m *corefoundation.CGAffineTransform, point corefoundation.CGPoint, eoFill bool) bool
var _cGPathContainsPointErr error

func tryCGPathContainsPoint(path CGPathRef, m *corefoundation.CGAffineTransform, point corefoundation.CGPoint, eoFill bool) (bool, error) {
	if _cGPathContainsPoint == nil {
		return false, symbolCallError("CGPathContainsPoint", "10.4", _cGPathContainsPointErr)
	}
	return _cGPathContainsPoint(path, m, point, eoFill), nil
}

// CGPathContainsPoint checks whether a point is contained in a graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathContainsPoint
func CGPathContainsPoint(path CGPathRef, m *corefoundation.CGAffineTransform, point corefoundation.CGPoint, eoFill bool) bool {
	result, callErr := tryCGPathContainsPoint(path, m, point, eoFill)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathCreateCopy func(path CGPathRef) CGPathRef
var _cGPathCreateCopyErr error

func tryCGPathCreateCopy(path CGPathRef) (CGPathRef, error) {
	if _cGPathCreateCopy == nil {
		return 0, symbolCallError("CGPathCreateCopy", "10.2", _cGPathCreateCopyErr)
	}
	return _cGPathCreateCopy(path), nil
}

// CGPathCreateCopy creates an immutable copy of a graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/copy()
func CGPathCreateCopy(path CGPathRef) CGPathRef {
	result, callErr := tryCGPathCreateCopy(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathCreateCopyByDashingPath func(path CGPathRef, transform *corefoundation.CGAffineTransform, phase float64, lengths *float64, count uintptr) CGPathRef
var _cGPathCreateCopyByDashingPathErr error

func tryCGPathCreateCopyByDashingPath(path CGPathRef, transform *corefoundation.CGAffineTransform, phase float64, lengths *float64, count uintptr) (CGPathRef, error) {
	if _cGPathCreateCopyByDashingPath == nil {
		return 0, symbolCallError("CGPathCreateCopyByDashingPath", "10.7", _cGPathCreateCopyByDashingPathErr)
	}
	return _cGPathCreateCopyByDashingPath(path, transform, phase, lengths, count), nil
}

// CGPathCreateCopyByDashingPath creates a dashed copy of another path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyByDashingPath
func CGPathCreateCopyByDashingPath(path CGPathRef, transform *corefoundation.CGAffineTransform, phase float64, lengths *float64, count uintptr) CGPathRef {
	result, callErr := tryCGPathCreateCopyByDashingPath(path, transform, phase, lengths, count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathCreateCopyByFlattening func(path CGPathRef, flatteningThreshold float64) CGPathRef
var _cGPathCreateCopyByFlatteningErr error

func tryCGPathCreateCopyByFlattening(path CGPathRef, flatteningThreshold float64) (CGPathRef, error) {
	if _cGPathCreateCopyByFlattening == nil {
		return 0, symbolCallError("CGPathCreateCopyByFlattening", "13.0", _cGPathCreateCopyByFlatteningErr)
	}
	return _cGPathCreateCopyByFlattening(path, flatteningThreshold), nil
}

// CGPathCreateCopyByFlattening.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyByFlattening
func CGPathCreateCopyByFlattening(path CGPathRef, flatteningThreshold float64) CGPathRef {
	result, callErr := tryCGPathCreateCopyByFlattening(path, flatteningThreshold)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathCreateCopyByIntersectingPath func(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef
var _cGPathCreateCopyByIntersectingPathErr error

func tryCGPathCreateCopyByIntersectingPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) (CGPathRef, error) {
	if _cGPathCreateCopyByIntersectingPath == nil {
		return 0, symbolCallError("CGPathCreateCopyByIntersectingPath", "13.0", _cGPathCreateCopyByIntersectingPathErr)
	}
	return _cGPathCreateCopyByIntersectingPath(path, maskPath, evenOddFillRule), nil
}

// CGPathCreateCopyByIntersectingPath.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyByIntersectingPath
func CGPathCreateCopyByIntersectingPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef {
	result, callErr := tryCGPathCreateCopyByIntersectingPath(path, maskPath, evenOddFillRule)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathCreateCopyByNormalizing func(path CGPathRef, evenOddFillRule bool) CGPathRef
var _cGPathCreateCopyByNormalizingErr error

func tryCGPathCreateCopyByNormalizing(path CGPathRef, evenOddFillRule bool) (CGPathRef, error) {
	if _cGPathCreateCopyByNormalizing == nil {
		return 0, symbolCallError("CGPathCreateCopyByNormalizing", "13.0", _cGPathCreateCopyByNormalizingErr)
	}
	return _cGPathCreateCopyByNormalizing(path, evenOddFillRule), nil
}

// CGPathCreateCopyByNormalizing.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyByNormalizing
func CGPathCreateCopyByNormalizing(path CGPathRef, evenOddFillRule bool) CGPathRef {
	result, callErr := tryCGPathCreateCopyByNormalizing(path, evenOddFillRule)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathCreateCopyByStrokingPath func(path CGPathRef, transform *corefoundation.CGAffineTransform, lineWidth float64, lineCap uint, lineJoin uint, miterLimit float64) CGPathRef
var _cGPathCreateCopyByStrokingPathErr error

func tryCGPathCreateCopyByStrokingPath(path CGPathRef, transform *corefoundation.CGAffineTransform, lineWidth float64, lineCap uint, lineJoin uint, miterLimit float64) (CGPathRef, error) {
	if _cGPathCreateCopyByStrokingPath == nil {
		return 0, symbolCallError("CGPathCreateCopyByStrokingPath", "10.7", _cGPathCreateCopyByStrokingPathErr)
	}
	return _cGPathCreateCopyByStrokingPath(path, transform, lineWidth, lineCap, lineJoin, miterLimit), nil
}

// CGPathCreateCopyByStrokingPath creates a stroked copy of another path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyByStrokingPath
func CGPathCreateCopyByStrokingPath(path CGPathRef, transform *corefoundation.CGAffineTransform, lineWidth float64, lineCap uint, lineJoin uint, miterLimit float64) CGPathRef {
	result, callErr := tryCGPathCreateCopyByStrokingPath(path, transform, lineWidth, lineCap, lineJoin, miterLimit)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathCreateCopyBySubtractingPath func(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef
var _cGPathCreateCopyBySubtractingPathErr error

func tryCGPathCreateCopyBySubtractingPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) (CGPathRef, error) {
	if _cGPathCreateCopyBySubtractingPath == nil {
		return 0, symbolCallError("CGPathCreateCopyBySubtractingPath", "13.0", _cGPathCreateCopyBySubtractingPathErr)
	}
	return _cGPathCreateCopyBySubtractingPath(path, maskPath, evenOddFillRule), nil
}

// CGPathCreateCopyBySubtractingPath.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyBySubtractingPath
func CGPathCreateCopyBySubtractingPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef {
	result, callErr := tryCGPathCreateCopyBySubtractingPath(path, maskPath, evenOddFillRule)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathCreateCopyBySymmetricDifferenceOfPath func(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef
var _cGPathCreateCopyBySymmetricDifferenceOfPathErr error

func tryCGPathCreateCopyBySymmetricDifferenceOfPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) (CGPathRef, error) {
	if _cGPathCreateCopyBySymmetricDifferenceOfPath == nil {
		return 0, symbolCallError("CGPathCreateCopyBySymmetricDifferenceOfPath", "13.0", _cGPathCreateCopyBySymmetricDifferenceOfPathErr)
	}
	return _cGPathCreateCopyBySymmetricDifferenceOfPath(path, maskPath, evenOddFillRule), nil
}

// CGPathCreateCopyBySymmetricDifferenceOfPath.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyBySymmetricDifferenceOfPath
func CGPathCreateCopyBySymmetricDifferenceOfPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef {
	result, callErr := tryCGPathCreateCopyBySymmetricDifferenceOfPath(path, maskPath, evenOddFillRule)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathCreateCopyByTransformingPath func(path CGPathRef, transform *corefoundation.CGAffineTransform) CGPathRef
var _cGPathCreateCopyByTransformingPathErr error

func tryCGPathCreateCopyByTransformingPath(path CGPathRef, transform *corefoundation.CGAffineTransform) (CGPathRef, error) {
	if _cGPathCreateCopyByTransformingPath == nil {
		return 0, symbolCallError("CGPathCreateCopyByTransformingPath", "10.7", _cGPathCreateCopyByTransformingPathErr)
	}
	return _cGPathCreateCopyByTransformingPath(path, transform), nil
}

// CGPathCreateCopyByTransformingPath creates an immutable copy of a graphics path transformed by a transformation matrix.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/copy(using:)
func CGPathCreateCopyByTransformingPath(path CGPathRef, transform *corefoundation.CGAffineTransform) CGPathRef {
	result, callErr := tryCGPathCreateCopyByTransformingPath(path, transform)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathCreateCopyByUnioningPath func(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef
var _cGPathCreateCopyByUnioningPathErr error

func tryCGPathCreateCopyByUnioningPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) (CGPathRef, error) {
	if _cGPathCreateCopyByUnioningPath == nil {
		return 0, symbolCallError("CGPathCreateCopyByUnioningPath", "13.0", _cGPathCreateCopyByUnioningPathErr)
	}
	return _cGPathCreateCopyByUnioningPath(path, maskPath, evenOddFillRule), nil
}

// CGPathCreateCopyByUnioningPath.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyByUnioningPath
func CGPathCreateCopyByUnioningPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef {
	result, callErr := tryCGPathCreateCopyByUnioningPath(path, maskPath, evenOddFillRule)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathCreateCopyOfLineByIntersectingPath func(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef
var _cGPathCreateCopyOfLineByIntersectingPathErr error

func tryCGPathCreateCopyOfLineByIntersectingPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) (CGPathRef, error) {
	if _cGPathCreateCopyOfLineByIntersectingPath == nil {
		return 0, symbolCallError("CGPathCreateCopyOfLineByIntersectingPath", "13.0", _cGPathCreateCopyOfLineByIntersectingPathErr)
	}
	return _cGPathCreateCopyOfLineByIntersectingPath(path, maskPath, evenOddFillRule), nil
}

// CGPathCreateCopyOfLineByIntersectingPath.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyOfLineByIntersectingPath
func CGPathCreateCopyOfLineByIntersectingPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef {
	result, callErr := tryCGPathCreateCopyOfLineByIntersectingPath(path, maskPath, evenOddFillRule)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathCreateCopyOfLineBySubtractingPath func(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef
var _cGPathCreateCopyOfLineBySubtractingPathErr error

func tryCGPathCreateCopyOfLineBySubtractingPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) (CGPathRef, error) {
	if _cGPathCreateCopyOfLineBySubtractingPath == nil {
		return 0, symbolCallError("CGPathCreateCopyOfLineBySubtractingPath", "13.0", _cGPathCreateCopyOfLineBySubtractingPathErr)
	}
	return _cGPathCreateCopyOfLineBySubtractingPath(path, maskPath, evenOddFillRule), nil
}

// CGPathCreateCopyOfLineBySubtractingPath.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateCopyOfLineBySubtractingPath
func CGPathCreateCopyOfLineBySubtractingPath(path CGPathRef, maskPath CGPathRef, evenOddFillRule bool) CGPathRef {
	result, callErr := tryCGPathCreateCopyOfLineBySubtractingPath(path, maskPath, evenOddFillRule)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathCreateMutable func() CGMutablePathRef
var _cGPathCreateMutableErr error

func tryCGPathCreateMutable() (CGMutablePathRef, error) {
	if _cGPathCreateMutable == nil {
		return 0, symbolCallError("CGPathCreateMutable", "10.2", _cGPathCreateMutableErr)
	}
	return _cGPathCreateMutable(), nil
}

// CGPathCreateMutable creates a mutable graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGMutablePath/init()
func CGPathCreateMutable() CGMutablePathRef {
	result, callErr := tryCGPathCreateMutable()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathCreateMutableCopy func(path CGPathRef) CGMutablePathRef
var _cGPathCreateMutableCopyErr error

func tryCGPathCreateMutableCopy(path CGPathRef) (CGMutablePathRef, error) {
	if _cGPathCreateMutableCopy == nil {
		return 0, symbolCallError("CGPathCreateMutableCopy", "10.2", _cGPathCreateMutableCopyErr)
	}
	return _cGPathCreateMutableCopy(path), nil
}

// CGPathCreateMutableCopy creates a mutable copy of an existing graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/mutableCopy()
func CGPathCreateMutableCopy(path CGPathRef) CGMutablePathRef {
	result, callErr := tryCGPathCreateMutableCopy(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathCreateMutableCopyByTransformingPath func(path CGPathRef, transform *corefoundation.CGAffineTransform) CGMutablePathRef
var _cGPathCreateMutableCopyByTransformingPathErr error

func tryCGPathCreateMutableCopyByTransformingPath(path CGPathRef, transform *corefoundation.CGAffineTransform) (CGMutablePathRef, error) {
	if _cGPathCreateMutableCopyByTransformingPath == nil {
		return 0, symbolCallError("CGPathCreateMutableCopyByTransformingPath", "10.7", _cGPathCreateMutableCopyByTransformingPathErr)
	}
	return _cGPathCreateMutableCopyByTransformingPath(path, transform), nil
}

// CGPathCreateMutableCopyByTransformingPath creates a mutable copy of a graphics path transformed by a transformation matrix.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/mutableCopy(using:)
func CGPathCreateMutableCopyByTransformingPath(path CGPathRef, transform *corefoundation.CGAffineTransform) CGMutablePathRef {
	result, callErr := tryCGPathCreateMutableCopyByTransformingPath(path, transform)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathCreateSeparateComponents func(path CGPathRef, evenOddFillRule bool) corefoundation.CFArrayRef
var _cGPathCreateSeparateComponentsErr error

func tryCGPathCreateSeparateComponents(path CGPathRef, evenOddFillRule bool) (corefoundation.CFArrayRef, error) {
	if _cGPathCreateSeparateComponents == nil {
		return 0, symbolCallError("CGPathCreateSeparateComponents", "13.0", _cGPathCreateSeparateComponentsErr)
	}
	return _cGPathCreateSeparateComponents(path, evenOddFillRule), nil
}

// CGPathCreateSeparateComponents.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathCreateSeparateComponents
func CGPathCreateSeparateComponents(path CGPathRef, evenOddFillRule bool) corefoundation.CFArrayRef {
	result, callErr := tryCGPathCreateSeparateComponents(path, evenOddFillRule)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathCreateWithEllipseInRect func(rect corefoundation.CGRect, transform *corefoundation.CGAffineTransform) CGPathRef
var _cGPathCreateWithEllipseInRectErr error

func tryCGPathCreateWithEllipseInRect(rect corefoundation.CGRect, transform *corefoundation.CGAffineTransform) (CGPathRef, error) {
	if _cGPathCreateWithEllipseInRect == nil {
		return 0, symbolCallError("CGPathCreateWithEllipseInRect", "10.7", _cGPathCreateWithEllipseInRectErr)
	}
	return _cGPathCreateWithEllipseInRect(rect, transform), nil
}

// CGPathCreateWithEllipseInRect create an immutable path of an ellipse.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/init(ellipseIn:transform:)
func CGPathCreateWithEllipseInRect(rect corefoundation.CGRect, transform *corefoundation.CGAffineTransform) CGPathRef {
	result, callErr := tryCGPathCreateWithEllipseInRect(rect, transform)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathCreateWithRect func(rect corefoundation.CGRect, transform *corefoundation.CGAffineTransform) CGPathRef
var _cGPathCreateWithRectErr error

func tryCGPathCreateWithRect(rect corefoundation.CGRect, transform *corefoundation.CGAffineTransform) (CGPathRef, error) {
	if _cGPathCreateWithRect == nil {
		return 0, symbolCallError("CGPathCreateWithRect", "10.5", _cGPathCreateWithRectErr)
	}
	return _cGPathCreateWithRect(rect, transform), nil
}

// CGPathCreateWithRect create an immutable path of a rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/init(rect:transform:)
func CGPathCreateWithRect(rect corefoundation.CGRect, transform *corefoundation.CGAffineTransform) CGPathRef {
	result, callErr := tryCGPathCreateWithRect(rect, transform)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathCreateWithRoundedRect func(rect corefoundation.CGRect, cornerWidth float64, cornerHeight float64, transform *corefoundation.CGAffineTransform) CGPathRef
var _cGPathCreateWithRoundedRectErr error

func tryCGPathCreateWithRoundedRect(rect corefoundation.CGRect, cornerWidth float64, cornerHeight float64, transform *corefoundation.CGAffineTransform) (CGPathRef, error) {
	if _cGPathCreateWithRoundedRect == nil {
		return 0, symbolCallError("CGPathCreateWithRoundedRect", "10.9", _cGPathCreateWithRoundedRectErr)
	}
	return _cGPathCreateWithRoundedRect(rect, cornerWidth, cornerHeight, transform), nil
}

// CGPathCreateWithRoundedRect create an immutable path of a rounded rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/init(roundedRect:cornerWidth:cornerHeight:transform:)
func CGPathCreateWithRoundedRect(rect corefoundation.CGRect, cornerWidth float64, cornerHeight float64, transform *corefoundation.CGAffineTransform) CGPathRef {
	result, callErr := tryCGPathCreateWithRoundedRect(rect, cornerWidth, cornerHeight, transform)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathEqualToPath func(path1 CGPathRef, path2 CGPathRef) bool
var _cGPathEqualToPathErr error

func tryCGPathEqualToPath(path1 CGPathRef, path2 CGPathRef) (bool, error) {
	if _cGPathEqualToPath == nil {
		return false, symbolCallError("CGPathEqualToPath", "10.2", _cGPathEqualToPathErr)
	}
	return _cGPathEqualToPath(path1, path2), nil
}

// CGPathEqualToPath indicates whether two graphics paths are equivalent.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathEqualToPath
func CGPathEqualToPath(path1 CGPathRef, path2 CGPathRef) bool {
	result, callErr := tryCGPathEqualToPath(path1, path2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathGetBoundingBox func(path CGPathRef) corefoundation.CGRect
var _cGPathGetBoundingBoxErr error

func tryCGPathGetBoundingBox(path CGPathRef) (corefoundation.CGRect, error) {
	if _cGPathGetBoundingBox == nil {
		return corefoundation.CGRect{}, symbolCallError("CGPathGetBoundingBox", "10.2", _cGPathGetBoundingBoxErr)
	}
	return _cGPathGetBoundingBox(path), nil
}

// CGPathGetBoundingBox returns the bounding box containing all points in a graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/boundingBox
func CGPathGetBoundingBox(path CGPathRef) corefoundation.CGRect {
	result, callErr := tryCGPathGetBoundingBox(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathGetCurrentPoint func(path CGPathRef) corefoundation.CGPoint
var _cGPathGetCurrentPointErr error

func tryCGPathGetCurrentPoint(path CGPathRef) (corefoundation.CGPoint, error) {
	if _cGPathGetCurrentPoint == nil {
		return corefoundation.CGPoint{}, symbolCallError("CGPathGetCurrentPoint", "10.2", _cGPathGetCurrentPointErr)
	}
	return _cGPathGetCurrentPoint(path), nil
}

// CGPathGetCurrentPoint returns the current point in a graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/currentPoint
func CGPathGetCurrentPoint(path CGPathRef) corefoundation.CGPoint {
	result, callErr := tryCGPathGetCurrentPoint(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathGetPathBoundingBox func(path CGPathRef) corefoundation.CGRect
var _cGPathGetPathBoundingBoxErr error

func tryCGPathGetPathBoundingBox(path CGPathRef) (corefoundation.CGRect, error) {
	if _cGPathGetPathBoundingBox == nil {
		return corefoundation.CGRect{}, symbolCallError("CGPathGetPathBoundingBox", "10.6", _cGPathGetPathBoundingBoxErr)
	}
	return _cGPathGetPathBoundingBox(path), nil
}

// CGPathGetPathBoundingBox returns the bounding box of a graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/boundingBoxOfPath
func CGPathGetPathBoundingBox(path CGPathRef) corefoundation.CGRect {
	result, callErr := tryCGPathGetPathBoundingBox(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathGetTypeID func() uint
var _cGPathGetTypeIDErr error

func tryCGPathGetTypeID() (uint, error) {
	if _cGPathGetTypeID == nil {
		return 0, symbolCallError("CGPathGetTypeID", "10.2", _cGPathGetTypeIDErr)
	}
	return _cGPathGetTypeID(), nil
}

// CGPathGetTypeID returns the Core Foundation type identifier for Core Graphics paths.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/typeID
func CGPathGetTypeID() uint {
	result, callErr := tryCGPathGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathIntersectsPath func(path1 CGPathRef, path2 CGPathRef, evenOddFillRule bool) bool
var _cGPathIntersectsPathErr error

func tryCGPathIntersectsPath(path1 CGPathRef, path2 CGPathRef, evenOddFillRule bool) (bool, error) {
	if _cGPathIntersectsPath == nil {
		return false, symbolCallError("CGPathIntersectsPath", "13.0", _cGPathIntersectsPathErr)
	}
	return _cGPathIntersectsPath(path1, path2, evenOddFillRule), nil
}

// CGPathIntersectsPath.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathIntersectsPath
func CGPathIntersectsPath(path1 CGPathRef, path2 CGPathRef, evenOddFillRule bool) bool {
	result, callErr := tryCGPathIntersectsPath(path1, path2, evenOddFillRule)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathIsEmpty func(path CGPathRef) bool
var _cGPathIsEmptyErr error

func tryCGPathIsEmpty(path CGPathRef) (bool, error) {
	if _cGPathIsEmpty == nil {
		return false, symbolCallError("CGPathIsEmpty", "10.2", _cGPathIsEmptyErr)
	}
	return _cGPathIsEmpty(path), nil
}

// CGPathIsEmpty indicates whether or not a graphics path is empty.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/isEmpty
func CGPathIsEmpty(path CGPathRef) bool {
	result, callErr := tryCGPathIsEmpty(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathIsRect func(path CGPathRef, rect *corefoundation.CGRect) bool
var _cGPathIsRectErr error

func tryCGPathIsRect(path CGPathRef, rect *corefoundation.CGRect) (bool, error) {
	if _cGPathIsRect == nil {
		return false, symbolCallError("CGPathIsRect", "10.2", _cGPathIsRectErr)
	}
	return _cGPathIsRect(path, rect), nil
}

// CGPathIsRect indicates whether or not a graphics path represents a rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPath/isRect(_:)
func CGPathIsRect(path CGPathRef, rect *corefoundation.CGRect) bool {
	result, callErr := tryCGPathIsRect(path, rect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPathMoveToPoint func(path CGMutablePathRef, m *corefoundation.CGAffineTransform, x float64, y float64)
var _cGPathMoveToPointErr error

func tryCGPathMoveToPoint(path CGMutablePathRef, m *corefoundation.CGAffineTransform, x float64, y float64) error {
	if _cGPathMoveToPoint == nil {
		return symbolCallError("CGPathMoveToPoint", "10.2", _cGPathMoveToPointErr)
	}
	_cGPathMoveToPoint(path, m, x, y)
	return nil
}

// CGPathMoveToPoint starts a new subpath at a specified location in a mutable graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathMoveToPoint
func CGPathMoveToPoint(path CGMutablePathRef, m *corefoundation.CGAffineTransform, x float64, y float64) {
	if callErr := tryCGPathMoveToPoint(path, m, x, y); callErr != nil {
		panic(callErr)
	}
}

var _cGPathRelease func(path CGPathRef)
var _cGPathReleaseErr error

func tryCGPathRelease(path CGPathRef) error {
	if _cGPathRelease == nil {
		return symbolCallError("CGPathRelease", "10.2", _cGPathReleaseErr)
	}
	_cGPathRelease(path)
	return nil
}

// CGPathRelease decrements the retain count of a graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathRelease
func CGPathRelease(path CGPathRef) {
	if callErr := tryCGPathRelease(path); callErr != nil {
		panic(callErr)
	}
}

var _cGPathRetain func(path CGPathRef) CGPathRef
var _cGPathRetainErr error

func tryCGPathRetain(path CGPathRef) (CGPathRef, error) {
	if _cGPathRetain == nil {
		return 0, symbolCallError("CGPathRetain", "10.2", _cGPathRetainErr)
	}
	return _cGPathRetain(path), nil
}

// CGPathRetain increments the retain count of a graphics path.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPathRetain
func CGPathRetain(path CGPathRef) CGPathRef {
	result, callErr := tryCGPathRetain(path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPatternCreate func(info unsafe.Pointer, bounds corefoundation.CGRect, matrix corefoundation.CGAffineTransform, xStep float64, yStep float64, tiling CGPatternTiling, isColored bool, callbacks *CGPatternCallbacks) CGPatternRef
var _cGPatternCreateErr error

func tryCGPatternCreate(info unsafe.Pointer, bounds corefoundation.CGRect, matrix corefoundation.CGAffineTransform, xStep float64, yStep float64, tiling CGPatternTiling, isColored bool, callbacks *CGPatternCallbacks) (CGPatternRef, error) {
	if _cGPatternCreate == nil {
		return 0, symbolCallError("CGPatternCreate", "10.0", _cGPatternCreateErr)
	}
	return _cGPatternCreate(info, bounds, matrix, xStep, yStep, tiling, isColored, callbacks), nil
}

// CGPatternCreate creates a pattern object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPattern/init(info:bounds:matrix:xStep:yStep:tiling:isColored:callbacks:)
func CGPatternCreate(info unsafe.Pointer, bounds corefoundation.CGRect, matrix corefoundation.CGAffineTransform, xStep float64, yStep float64, tiling CGPatternTiling, isColored bool, callbacks *CGPatternCallbacks) CGPatternRef {
	result, callErr := tryCGPatternCreate(info, bounds, matrix, xStep, yStep, tiling, isColored, callbacks)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPatternGetTypeID func() uint
var _cGPatternGetTypeIDErr error

func tryCGPatternGetTypeID() (uint, error) {
	if _cGPatternGetTypeID == nil {
		return 0, symbolCallError("CGPatternGetTypeID", "10.2", _cGPatternGetTypeIDErr)
	}
	return _cGPatternGetTypeID(), nil
}

// CGPatternGetTypeID returns the type identifier for Core Graphics patterns.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPattern/typeID
func CGPatternGetTypeID() uint {
	result, callErr := tryCGPatternGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPatternRelease func(pattern CGPatternRef)
var _cGPatternReleaseErr error

func tryCGPatternRelease(pattern CGPatternRef) error {
	if _cGPatternRelease == nil {
		return symbolCallError("CGPatternRelease", "10.0", _cGPatternReleaseErr)
	}
	_cGPatternRelease(pattern)
	return nil
}

// CGPatternRelease decrements the retain count of a Core Graphics pattern.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPatternRelease
func CGPatternRelease(pattern CGPatternRef) {
	if callErr := tryCGPatternRelease(pattern); callErr != nil {
		panic(callErr)
	}
}

var _cGPatternRetain func(pattern CGPatternRef) CGPatternRef
var _cGPatternRetainErr error

func tryCGPatternRetain(pattern CGPatternRef) (CGPatternRef, error) {
	if _cGPatternRetain == nil {
		return 0, symbolCallError("CGPatternRetain", "10.0", _cGPatternRetainErr)
	}
	return _cGPatternRetain(pattern), nil
}

// CGPatternRetain increments the retain count of a Core Graphics pattern.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPatternRetain
func CGPatternRetain(pattern CGPatternRef) CGPatternRef {
	result, callErr := tryCGPatternRetain(pattern)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPointApplyAffineTransform func(point corefoundation.CGPoint, t corefoundation.CGAffineTransform) corefoundation.CGPoint
var _cGPointApplyAffineTransformErr error

func tryCGPointApplyAffineTransform(point corefoundation.CGPoint, t corefoundation.CGAffineTransform) (corefoundation.CGPoint, error) {
	if _cGPointApplyAffineTransform == nil {
		return corefoundation.CGPoint{}, symbolCallError("CGPointApplyAffineTransform", "10.0", _cGPointApplyAffineTransformErr)
	}
	return _cGPointApplyAffineTransform(point, t), nil
}

// CGPointApplyAffineTransform returns the point resulting from an affine transformation of an existing point.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPointApplyAffineTransform(_:_:)
func CGPointApplyAffineTransform(point corefoundation.CGPoint, t corefoundation.CGAffineTransform) corefoundation.CGPoint {
	result, callErr := tryCGPointApplyAffineTransform(point, t)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPointCreateDictionaryRepresentation func(point corefoundation.CGPoint) corefoundation.CFDictionaryRef
var _cGPointCreateDictionaryRepresentationErr error

func tryCGPointCreateDictionaryRepresentation(point corefoundation.CGPoint) (corefoundation.CFDictionaryRef, error) {
	if _cGPointCreateDictionaryRepresentation == nil {
		return 0, symbolCallError("CGPointCreateDictionaryRepresentation", "10.5", _cGPointCreateDictionaryRepresentationErr)
	}
	return _cGPointCreateDictionaryRepresentation(point), nil
}

// CGPointCreateDictionaryRepresentation returns a dictionary representation of the specified point.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPointCreateDictionaryRepresentation(_:)
func CGPointCreateDictionaryRepresentation(point corefoundation.CGPoint) corefoundation.CFDictionaryRef {
	result, callErr := tryCGPointCreateDictionaryRepresentation(point)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPointEqualToPoint func(point1 corefoundation.CGPoint, point2 corefoundation.CGPoint) bool
var _cGPointEqualToPointErr error

func tryCGPointEqualToPoint(point1 corefoundation.CGPoint, point2 corefoundation.CGPoint) (bool, error) {
	if _cGPointEqualToPoint == nil {
		return false, symbolCallError("CGPointEqualToPoint", "10.0", _cGPointEqualToPointErr)
	}
	return _cGPointEqualToPoint(point1, point2), nil
}

// CGPointEqualToPoint returns whether two points are equal.
//
// Deprecated: The [CGPoint](<doc://com.apple.documentation/documentation/CoreFoundation/CGPoint>) type adopts the [Equatable] protocol; use the `==` operator instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPointEqualToPoint(_:_:)
func CGPointEqualToPoint(point1 corefoundation.CGPoint, point2 corefoundation.CGPoint) bool {
	result, callErr := tryCGPointEqualToPoint(point1, point2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPointMakeWithDictionaryRepresentation func(dict corefoundation.CFDictionaryRef, point *corefoundation.CGPoint) bool
var _cGPointMakeWithDictionaryRepresentationErr error

func tryCGPointMakeWithDictionaryRepresentation(dict corefoundation.CFDictionaryRef, point *corefoundation.CGPoint) (bool, error) {
	if _cGPointMakeWithDictionaryRepresentation == nil {
		return false, symbolCallError("CGPointMakeWithDictionaryRepresentation", "10.5", _cGPointMakeWithDictionaryRepresentationErr)
	}
	return _cGPointMakeWithDictionaryRepresentation(dict, point), nil
}

// CGPointMakeWithDictionaryRepresentation fills in a point using the contents of the specified dictionary.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPointMakeWithDictionaryRepresentation(_:_:)
func CGPointMakeWithDictionaryRepresentation(dict corefoundation.CFDictionaryRef, point *corefoundation.CGPoint) bool {
	result, callErr := tryCGPointMakeWithDictionaryRepresentation(dict, point)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPostMouseEvent func(mouseCursorPosition corefoundation.CGPoint, updateMouseCursorPosition bool, buttonCount CGButtonCount, mouseButtonDown bool) CGError
var _cGPostMouseEventErr error

func tryCGPostMouseEvent(mouseCursorPosition corefoundation.CGPoint, updateMouseCursorPosition bool, buttonCount CGButtonCount, mouseButtonDown bool) (CGError, error) {
	if _cGPostMouseEvent == nil {
		return *new(CGError), symbolCallError("CGPostMouseEvent", "10.0", _cGPostMouseEventErr)
	}
	return _cGPostMouseEvent(mouseCursorPosition, updateMouseCursorPosition, buttonCount, mouseButtonDown), nil
}

// CGPostMouseEvent synthesizes a low-level mouse-button event on the local machine.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPostMouseEvent
func CGPostMouseEvent(mouseCursorPosition corefoundation.CGPoint, updateMouseCursorPosition bool, buttonCount CGButtonCount, mouseButtonDown bool) CGError {
	result, callErr := tryCGPostMouseEvent(mouseCursorPosition, updateMouseCursorPosition, buttonCount, mouseButtonDown)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPostScrollWheelEvent func(wheelCount CGWheelCount, wheel1 int32) CGError
var _cGPostScrollWheelEventErr error

func tryCGPostScrollWheelEvent(wheelCount CGWheelCount, wheel1 int32) (CGError, error) {
	if _cGPostScrollWheelEvent == nil {
		return *new(CGError), symbolCallError("CGPostScrollWheelEvent", "10.0", _cGPostScrollWheelEventErr)
	}
	return _cGPostScrollWheelEvent(wheelCount, wheel1), nil
}

// CGPostScrollWheelEvent synthesizes a low-level scrolling event on the local machine.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPostScrollWheelEvent
func CGPostScrollWheelEvent(wheelCount CGWheelCount, wheel1 int32) CGError {
	result, callErr := tryCGPostScrollWheelEvent(wheelCount, wheel1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPreflightListenEventAccess func() bool
var _cGPreflightListenEventAccessErr error

func tryCGPreflightListenEventAccess() (bool, error) {
	if _cGPreflightListenEventAccess == nil {
		return false, symbolCallError("CGPreflightListenEventAccess", "10.15", _cGPreflightListenEventAccessErr)
	}
	return _cGPreflightListenEventAccess(), nil
}

// CGPreflightListenEventAccess.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPreflightListenEventAccess()
func CGPreflightListenEventAccess() bool {
	result, callErr := tryCGPreflightListenEventAccess()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPreflightPostEventAccess func() bool
var _cGPreflightPostEventAccessErr error

func tryCGPreflightPostEventAccess() (bool, error) {
	if _cGPreflightPostEventAccess == nil {
		return false, symbolCallError("CGPreflightPostEventAccess", "10.15", _cGPreflightPostEventAccessErr)
	}
	return _cGPreflightPostEventAccess(), nil
}

// CGPreflightPostEventAccess.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPreflightPostEventAccess()
func CGPreflightPostEventAccess() bool {
	result, callErr := tryCGPreflightPostEventAccess()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGPreflightScreenCaptureAccess func() bool
var _cGPreflightScreenCaptureAccessErr error

func tryCGPreflightScreenCaptureAccess() (bool, error) {
	if _cGPreflightScreenCaptureAccess == nil {
		return false, symbolCallError("CGPreflightScreenCaptureAccess", "10.15", _cGPreflightScreenCaptureAccessErr)
	}
	return _cGPreflightScreenCaptureAccess(), nil
}

// CGPreflightScreenCaptureAccess.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGPreflightScreenCaptureAccess()
func CGPreflightScreenCaptureAccess() bool {
	result, callErr := tryCGPreflightScreenCaptureAccess()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectApplyAffineTransform func(rect corefoundation.CGRect, t corefoundation.CGAffineTransform) corefoundation.CGRect
var _cGRectApplyAffineTransformErr error

func tryCGRectApplyAffineTransform(rect corefoundation.CGRect, t corefoundation.CGAffineTransform) (corefoundation.CGRect, error) {
	if _cGRectApplyAffineTransform == nil {
		return corefoundation.CGRect{}, symbolCallError("CGRectApplyAffineTransform", "10.4", _cGRectApplyAffineTransformErr)
	}
	return _cGRectApplyAffineTransform(rect, t), nil
}

// CGRectApplyAffineTransform applies an affine transform to a rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectApplyAffineTransform(_:_:)
func CGRectApplyAffineTransform(rect corefoundation.CGRect, t corefoundation.CGAffineTransform) corefoundation.CGRect {
	result, callErr := tryCGRectApplyAffineTransform(rect, t)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectContainsPoint func(rect corefoundation.CGRect, point corefoundation.CGPoint) bool
var _cGRectContainsPointErr error

func tryCGRectContainsPoint(rect corefoundation.CGRect, point corefoundation.CGPoint) (bool, error) {
	if _cGRectContainsPoint == nil {
		return false, symbolCallError("CGRectContainsPoint", "10.0", _cGRectContainsPointErr)
	}
	return _cGRectContainsPoint(rect, point), nil
}

// CGRectContainsPoint returns whether a rectangle contains a specified point.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectContainsPoint(_:_:)
func CGRectContainsPoint(rect corefoundation.CGRect, point corefoundation.CGPoint) bool {
	result, callErr := tryCGRectContainsPoint(rect, point)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectContainsRect func(rect1 corefoundation.CGRect, rect2 corefoundation.CGRect) bool
var _cGRectContainsRectErr error

func tryCGRectContainsRect(rect1 corefoundation.CGRect, rect2 corefoundation.CGRect) (bool, error) {
	if _cGRectContainsRect == nil {
		return false, symbolCallError("CGRectContainsRect", "10.0", _cGRectContainsRectErr)
	}
	return _cGRectContainsRect(rect1, rect2), nil
}

// CGRectContainsRect returns whether the first rectangle contains the second rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectContainsRect(_:_:)
func CGRectContainsRect(rect1 corefoundation.CGRect, rect2 corefoundation.CGRect) bool {
	result, callErr := tryCGRectContainsRect(rect1, rect2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectCreateDictionaryRepresentation func(arg0 corefoundation.CGRect) corefoundation.CFDictionaryRef
var _cGRectCreateDictionaryRepresentationErr error

func tryCGRectCreateDictionaryRepresentation(arg0 corefoundation.CGRect) (corefoundation.CFDictionaryRef, error) {
	if _cGRectCreateDictionaryRepresentation == nil {
		return 0, symbolCallError("CGRectCreateDictionaryRepresentation", "10.5", _cGRectCreateDictionaryRepresentationErr)
	}
	return _cGRectCreateDictionaryRepresentation(arg0), nil
}

// CGRectCreateDictionaryRepresentation returns a dictionary representation of the provided rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectCreateDictionaryRepresentation(_:)
func CGRectCreateDictionaryRepresentation(arg0 corefoundation.CGRect) corefoundation.CFDictionaryRef {
	result, callErr := tryCGRectCreateDictionaryRepresentation(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectDivide func(rect corefoundation.CGRect, slice *corefoundation.CGRect, remainder *corefoundation.CGRect, amount float64, edge CGRectEdge)
var _cGRectDivideErr error

func tryCGRectDivide(rect corefoundation.CGRect, slice *corefoundation.CGRect, remainder *corefoundation.CGRect, amount float64, edge CGRectEdge) error {
	if _cGRectDivide == nil {
		return symbolCallError("CGRectDivide", "10.0", _cGRectDivideErr)
	}
	_cGRectDivide(rect, slice, remainder, amount, edge)
	return nil
}

// CGRectDivide divides a source rectangle into two component rectangles.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectDivide
func CGRectDivide(rect corefoundation.CGRect, slice *corefoundation.CGRect, remainder *corefoundation.CGRect, amount float64, edge CGRectEdge) {
	if callErr := tryCGRectDivide(rect, slice, remainder, amount, edge); callErr != nil {
		panic(callErr)
	}
}

var _cGRectEqualToRect func(rect1 corefoundation.CGRect, rect2 corefoundation.CGRect) bool
var _cGRectEqualToRectErr error

func tryCGRectEqualToRect(rect1 corefoundation.CGRect, rect2 corefoundation.CGRect) (bool, error) {
	if _cGRectEqualToRect == nil {
		return false, symbolCallError("CGRectEqualToRect", "10.0", _cGRectEqualToRectErr)
	}
	return _cGRectEqualToRect(rect1, rect2), nil
}

// CGRectEqualToRect returns whether two rectangles are equal in size and position.
//
// Deprecated: The [CGRect](<doc://com.apple.documentation/documentation/CoreFoundation/CGRect>) type adopts the [Equatable] protocol; use the `==` operator instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectEqualToRect(_:_:)
func CGRectEqualToRect(rect1 corefoundation.CGRect, rect2 corefoundation.CGRect) bool {
	result, callErr := tryCGRectEqualToRect(rect1, rect2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectGetHeight func(rect corefoundation.CGRect) float64
var _cGRectGetHeightErr error

func tryCGRectGetHeight(rect corefoundation.CGRect) (float64, error) {
	if _cGRectGetHeight == nil {
		return 0.0, symbolCallError("CGRectGetHeight", "10.0", _cGRectGetHeightErr)
	}
	return _cGRectGetHeight(rect), nil
}

// CGRectGetHeight returns the height of a rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectGetHeight(_:)
func CGRectGetHeight(rect corefoundation.CGRect) float64 {
	result, callErr := tryCGRectGetHeight(rect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectGetMaxX func(rect corefoundation.CGRect) float64
var _cGRectGetMaxXErr error

func tryCGRectGetMaxX(rect corefoundation.CGRect) (float64, error) {
	if _cGRectGetMaxX == nil {
		return 0.0, symbolCallError("CGRectGetMaxX", "10.0", _cGRectGetMaxXErr)
	}
	return _cGRectGetMaxX(rect), nil
}

// CGRectGetMaxX returns the largest value of the x-coordinate for the rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectGetMaxX(_:)
func CGRectGetMaxX(rect corefoundation.CGRect) float64 {
	result, callErr := tryCGRectGetMaxX(rect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectGetMaxY func(rect corefoundation.CGRect) float64
var _cGRectGetMaxYErr error

func tryCGRectGetMaxY(rect corefoundation.CGRect) (float64, error) {
	if _cGRectGetMaxY == nil {
		return 0.0, symbolCallError("CGRectGetMaxY", "10.0", _cGRectGetMaxYErr)
	}
	return _cGRectGetMaxY(rect), nil
}

// CGRectGetMaxY returns the largest value for the y-coordinate of the rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectGetMaxY(_:)
func CGRectGetMaxY(rect corefoundation.CGRect) float64 {
	result, callErr := tryCGRectGetMaxY(rect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectGetMidX func(rect corefoundation.CGRect) float64
var _cGRectGetMidXErr error

func tryCGRectGetMidX(rect corefoundation.CGRect) (float64, error) {
	if _cGRectGetMidX == nil {
		return 0.0, symbolCallError("CGRectGetMidX", "10.0", _cGRectGetMidXErr)
	}
	return _cGRectGetMidX(rect), nil
}

// CGRectGetMidX returns the x- coordinate that establishes the center of a rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectGetMidX(_:)
func CGRectGetMidX(rect corefoundation.CGRect) float64 {
	result, callErr := tryCGRectGetMidX(rect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectGetMidY func(rect corefoundation.CGRect) float64
var _cGRectGetMidYErr error

func tryCGRectGetMidY(rect corefoundation.CGRect) (float64, error) {
	if _cGRectGetMidY == nil {
		return 0.0, symbolCallError("CGRectGetMidY", "10.0", _cGRectGetMidYErr)
	}
	return _cGRectGetMidY(rect), nil
}

// CGRectGetMidY returns the y-coordinate that establishes the center of the rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectGetMidY(_:)
func CGRectGetMidY(rect corefoundation.CGRect) float64 {
	result, callErr := tryCGRectGetMidY(rect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectGetMinX func(rect corefoundation.CGRect) float64
var _cGRectGetMinXErr error

func tryCGRectGetMinX(rect corefoundation.CGRect) (float64, error) {
	if _cGRectGetMinX == nil {
		return 0.0, symbolCallError("CGRectGetMinX", "10.0", _cGRectGetMinXErr)
	}
	return _cGRectGetMinX(rect), nil
}

// CGRectGetMinX returns the smallest value for the x-coordinate of the rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectGetMinX(_:)
func CGRectGetMinX(rect corefoundation.CGRect) float64 {
	result, callErr := tryCGRectGetMinX(rect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectGetMinY func(rect corefoundation.CGRect) float64
var _cGRectGetMinYErr error

func tryCGRectGetMinY(rect corefoundation.CGRect) (float64, error) {
	if _cGRectGetMinY == nil {
		return 0.0, symbolCallError("CGRectGetMinY", "10.0", _cGRectGetMinYErr)
	}
	return _cGRectGetMinY(rect), nil
}

// CGRectGetMinY returns the smallest value for the y-coordinate of the rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectGetMinY(_:)
func CGRectGetMinY(rect corefoundation.CGRect) float64 {
	result, callErr := tryCGRectGetMinY(rect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectGetWidth func(rect corefoundation.CGRect) float64
var _cGRectGetWidthErr error

func tryCGRectGetWidth(rect corefoundation.CGRect) (float64, error) {
	if _cGRectGetWidth == nil {
		return 0.0, symbolCallError("CGRectGetWidth", "10.0", _cGRectGetWidthErr)
	}
	return _cGRectGetWidth(rect), nil
}

// CGRectGetWidth returns the width of a rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectGetWidth(_:)
func CGRectGetWidth(rect corefoundation.CGRect) float64 {
	result, callErr := tryCGRectGetWidth(rect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectInset func(rect corefoundation.CGRect, dx float64, dy float64) corefoundation.CGRect
var _cGRectInsetErr error

func tryCGRectInset(rect corefoundation.CGRect, dx float64, dy float64) (corefoundation.CGRect, error) {
	if _cGRectInset == nil {
		return corefoundation.CGRect{}, symbolCallError("CGRectInset", "10.0", _cGRectInsetErr)
	}
	return _cGRectInset(rect, dx, dy), nil
}

// CGRectInset returns a rectangle that is smaller or larger than the source rectangle, with the same center point.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectInset(_:_:_:)
func CGRectInset(rect corefoundation.CGRect, dx float64, dy float64) corefoundation.CGRect {
	result, callErr := tryCGRectInset(rect, dx, dy)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectIntegral func(rect corefoundation.CGRect) corefoundation.CGRect
var _cGRectIntegralErr error

func tryCGRectIntegral(rect corefoundation.CGRect) (corefoundation.CGRect, error) {
	if _cGRectIntegral == nil {
		return corefoundation.CGRect{}, symbolCallError("CGRectIntegral", "10.0", _cGRectIntegralErr)
	}
	return _cGRectIntegral(rect), nil
}

// CGRectIntegral returns the smallest rectangle that results from converting the source rectangle values to integers.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectIntegral(_:)
func CGRectIntegral(rect corefoundation.CGRect) corefoundation.CGRect {
	result, callErr := tryCGRectIntegral(rect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectIntersection func(r1 corefoundation.CGRect, r2 corefoundation.CGRect) corefoundation.CGRect
var _cGRectIntersectionErr error

func tryCGRectIntersection(r1 corefoundation.CGRect, r2 corefoundation.CGRect) (corefoundation.CGRect, error) {
	if _cGRectIntersection == nil {
		return corefoundation.CGRect{}, symbolCallError("CGRectIntersection", "10.0", _cGRectIntersectionErr)
	}
	return _cGRectIntersection(r1, r2), nil
}

// CGRectIntersection returns the intersection of two rectangles.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectIntersection(_:_:)
func CGRectIntersection(r1 corefoundation.CGRect, r2 corefoundation.CGRect) corefoundation.CGRect {
	result, callErr := tryCGRectIntersection(r1, r2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectIntersectsRect func(rect1 corefoundation.CGRect, rect2 corefoundation.CGRect) bool
var _cGRectIntersectsRectErr error

func tryCGRectIntersectsRect(rect1 corefoundation.CGRect, rect2 corefoundation.CGRect) (bool, error) {
	if _cGRectIntersectsRect == nil {
		return false, symbolCallError("CGRectIntersectsRect", "10.0", _cGRectIntersectsRectErr)
	}
	return _cGRectIntersectsRect(rect1, rect2), nil
}

// CGRectIntersectsRect returns whether two rectangles intersect.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectIntersectsRect(_:_:)
func CGRectIntersectsRect(rect1 corefoundation.CGRect, rect2 corefoundation.CGRect) bool {
	result, callErr := tryCGRectIntersectsRect(rect1, rect2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectIsEmpty func(rect corefoundation.CGRect) bool
var _cGRectIsEmptyErr error

func tryCGRectIsEmpty(rect corefoundation.CGRect) (bool, error) {
	if _cGRectIsEmpty == nil {
		return false, symbolCallError("CGRectIsEmpty", "10.0", _cGRectIsEmptyErr)
	}
	return _cGRectIsEmpty(rect), nil
}

// CGRectIsEmpty returns whether a rectangle has zero width or height, or is a null rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectIsEmpty(_:)
func CGRectIsEmpty(rect corefoundation.CGRect) bool {
	result, callErr := tryCGRectIsEmpty(rect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectIsInfinite func(rect corefoundation.CGRect) bool
var _cGRectIsInfiniteErr error

func tryCGRectIsInfinite(rect corefoundation.CGRect) (bool, error) {
	if _cGRectIsInfinite == nil {
		return false, symbolCallError("CGRectIsInfinite", "10.4", _cGRectIsInfiniteErr)
	}
	return _cGRectIsInfinite(rect), nil
}

// CGRectIsInfinite returns whether a rectangle is infinite.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectIsInfinite(_:)
func CGRectIsInfinite(rect corefoundation.CGRect) bool {
	result, callErr := tryCGRectIsInfinite(rect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectIsNull func(rect corefoundation.CGRect) bool
var _cGRectIsNullErr error

func tryCGRectIsNull(rect corefoundation.CGRect) (bool, error) {
	if _cGRectIsNull == nil {
		return false, symbolCallError("CGRectIsNull", "10.0", _cGRectIsNullErr)
	}
	return _cGRectIsNull(rect), nil
}

// CGRectIsNull returns whether the rectangle is equal to the null rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectIsNull(_:)
func CGRectIsNull(rect corefoundation.CGRect) bool {
	result, callErr := tryCGRectIsNull(rect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectMakeWithDictionaryRepresentation func(dict corefoundation.CFDictionaryRef, rect *corefoundation.CGRect) bool
var _cGRectMakeWithDictionaryRepresentationErr error

func tryCGRectMakeWithDictionaryRepresentation(dict corefoundation.CFDictionaryRef, rect *corefoundation.CGRect) (bool, error) {
	if _cGRectMakeWithDictionaryRepresentation == nil {
		return false, symbolCallError("CGRectMakeWithDictionaryRepresentation", "10.5", _cGRectMakeWithDictionaryRepresentationErr)
	}
	return _cGRectMakeWithDictionaryRepresentation(dict, rect), nil
}

// CGRectMakeWithDictionaryRepresentation fills in a rectangle using the contents of the specified dictionary.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectMakeWithDictionaryRepresentation(_:_:)
func CGRectMakeWithDictionaryRepresentation(dict corefoundation.CFDictionaryRef, rect *corefoundation.CGRect) bool {
	result, callErr := tryCGRectMakeWithDictionaryRepresentation(dict, rect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectOffset func(rect corefoundation.CGRect, dx float64, dy float64) corefoundation.CGRect
var _cGRectOffsetErr error

func tryCGRectOffset(rect corefoundation.CGRect, dx float64, dy float64) (corefoundation.CGRect, error) {
	if _cGRectOffset == nil {
		return corefoundation.CGRect{}, symbolCallError("CGRectOffset", "10.0", _cGRectOffsetErr)
	}
	return _cGRectOffset(rect, dx, dy), nil
}

// CGRectOffset returns a rectangle with an origin that is offset from that of the source rectangle.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectOffset(_:_:_:)
func CGRectOffset(rect corefoundation.CGRect, dx float64, dy float64) corefoundation.CGRect {
	result, callErr := tryCGRectOffset(rect, dx, dy)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectStandardize func(rect corefoundation.CGRect) corefoundation.CGRect
var _cGRectStandardizeErr error

func tryCGRectStandardize(rect corefoundation.CGRect) (corefoundation.CGRect, error) {
	if _cGRectStandardize == nil {
		return corefoundation.CGRect{}, symbolCallError("CGRectStandardize", "10.0", _cGRectStandardizeErr)
	}
	return _cGRectStandardize(rect), nil
}

// CGRectStandardize returns a rectangle with a positive width and height.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectStandardize(_:)
func CGRectStandardize(rect corefoundation.CGRect) corefoundation.CGRect {
	result, callErr := tryCGRectStandardize(rect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRectUnion func(r1 corefoundation.CGRect, r2 corefoundation.CGRect) corefoundation.CGRect
var _cGRectUnionErr error

func tryCGRectUnion(r1 corefoundation.CGRect, r2 corefoundation.CGRect) (corefoundation.CGRect, error) {
	if _cGRectUnion == nil {
		return corefoundation.CGRect{}, symbolCallError("CGRectUnion", "10.0", _cGRectUnionErr)
	}
	return _cGRectUnion(r1, r2), nil
}

// CGRectUnion returns the smallest rectangle that contains the two source rectangles.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRectUnion(_:_:)
func CGRectUnion(r1 corefoundation.CGRect, r2 corefoundation.CGRect) corefoundation.CGRect {
	result, callErr := tryCGRectUnion(r1, r2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGReleaseAllDisplays func() CGError
var _cGReleaseAllDisplaysErr error

func tryCGReleaseAllDisplays() (CGError, error) {
	if _cGReleaseAllDisplays == nil {
		return *new(CGError), symbolCallError("CGReleaseAllDisplays", "10.0", _cGReleaseAllDisplaysErr)
	}
	return _cGReleaseAllDisplays(), nil
}

// CGReleaseAllDisplays releases all captured displays.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGReleaseAllDisplays()
func CGReleaseAllDisplays() CGError {
	result, callErr := tryCGReleaseAllDisplays()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGReleaseDisplayFadeReservation func(token CGDisplayFadeReservationToken) CGError
var _cGReleaseDisplayFadeReservationErr error

func tryCGReleaseDisplayFadeReservation(token CGDisplayFadeReservationToken) (CGError, error) {
	if _cGReleaseDisplayFadeReservation == nil {
		return *new(CGError), symbolCallError("CGReleaseDisplayFadeReservation", "10.2", _cGReleaseDisplayFadeReservationErr)
	}
	return _cGReleaseDisplayFadeReservation(token), nil
}

// CGReleaseDisplayFadeReservation releases a display fade reservation, and unfades the display if needed.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGReleaseDisplayFadeReservation(_:)
func CGReleaseDisplayFadeReservation(token CGDisplayFadeReservationToken) CGError {
	result, callErr := tryCGReleaseDisplayFadeReservation(token)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRenderingBufferLockBytePtr func(provider CGRenderingBufferProviderRef) unsafe.Pointer
var _cGRenderingBufferLockBytePtrErr error

func tryCGRenderingBufferLockBytePtr(provider CGRenderingBufferProviderRef) (unsafe.Pointer, error) {
	if _cGRenderingBufferLockBytePtr == nil {
		return nil, symbolCallError("CGRenderingBufferLockBytePtr", "26.0", _cGRenderingBufferLockBytePtrErr)
	}
	return _cGRenderingBufferLockBytePtr(provider), nil
}

// CGRenderingBufferLockBytePtr.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRenderingBufferLockBytePtr
func CGRenderingBufferLockBytePtr(provider CGRenderingBufferProviderRef) unsafe.Pointer {
	result, callErr := tryCGRenderingBufferLockBytePtr(provider)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRenderingBufferProviderCreate func(info unsafe.Pointer, size uintptr) CGRenderingBufferProviderRef
var _cGRenderingBufferProviderCreateErr error

func tryCGRenderingBufferProviderCreate(info unsafe.Pointer, size uintptr) (CGRenderingBufferProviderRef, error) {
	if _cGRenderingBufferProviderCreate == nil {
		return 0, symbolCallError("CGRenderingBufferProviderCreate", "26.0", _cGRenderingBufferProviderCreateErr)
	}
	return _cGRenderingBufferProviderCreate(info, size), nil
}

// CGRenderingBufferProviderCreate.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRenderingBufferProviderCreate
func CGRenderingBufferProviderCreate(info unsafe.Pointer, size uintptr) CGRenderingBufferProviderRef {
	result, callErr := tryCGRenderingBufferProviderCreate(info, size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRenderingBufferProviderCreateWithCFData func(data corefoundation.CFMutableDataRef) CGRenderingBufferProviderRef
var _cGRenderingBufferProviderCreateWithCFDataErr error

func tryCGRenderingBufferProviderCreateWithCFData(data corefoundation.CFMutableDataRef) (CGRenderingBufferProviderRef, error) {
	if _cGRenderingBufferProviderCreateWithCFData == nil {
		return 0, symbolCallError("CGRenderingBufferProviderCreateWithCFData", "26.0", _cGRenderingBufferProviderCreateWithCFDataErr)
	}
	return _cGRenderingBufferProviderCreateWithCFData(data), nil
}

// CGRenderingBufferProviderCreateWithCFData.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRenderingBufferProviderCreateWithCFData
func CGRenderingBufferProviderCreateWithCFData(data corefoundation.CFMutableDataRef) CGRenderingBufferProviderRef {
	result, callErr := tryCGRenderingBufferProviderCreateWithCFData(data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRenderingBufferProviderGetSize func(provider CGRenderingBufferProviderRef) uintptr
var _cGRenderingBufferProviderGetSizeErr error

func tryCGRenderingBufferProviderGetSize(provider CGRenderingBufferProviderRef) (uintptr, error) {
	if _cGRenderingBufferProviderGetSize == nil {
		return 0, symbolCallError("CGRenderingBufferProviderGetSize", "26.0", _cGRenderingBufferProviderGetSizeErr)
	}
	return _cGRenderingBufferProviderGetSize(provider), nil
}

// CGRenderingBufferProviderGetSize.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRenderingBufferProviderGetSize
func CGRenderingBufferProviderGetSize(provider CGRenderingBufferProviderRef) uintptr {
	result, callErr := tryCGRenderingBufferProviderGetSize(provider)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRenderingBufferProviderGetTypeID func() uint
var _cGRenderingBufferProviderGetTypeIDErr error

func tryCGRenderingBufferProviderGetTypeID() (uint, error) {
	if _cGRenderingBufferProviderGetTypeID == nil {
		return 0, symbolCallError("CGRenderingBufferProviderGetTypeID", "26.0", _cGRenderingBufferProviderGetTypeIDErr)
	}
	return _cGRenderingBufferProviderGetTypeID(), nil
}

// CGRenderingBufferProviderGetTypeID.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRenderingBufferProviderGetTypeID
func CGRenderingBufferProviderGetTypeID() uint {
	result, callErr := tryCGRenderingBufferProviderGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRenderingBufferUnlockBytePtr func(provider CGRenderingBufferProviderRef)
var _cGRenderingBufferUnlockBytePtrErr error

func tryCGRenderingBufferUnlockBytePtr(provider CGRenderingBufferProviderRef) error {
	if _cGRenderingBufferUnlockBytePtr == nil {
		return symbolCallError("CGRenderingBufferUnlockBytePtr", "26.0", _cGRenderingBufferUnlockBytePtrErr)
	}
	_cGRenderingBufferUnlockBytePtr(provider)
	return nil
}

// CGRenderingBufferUnlockBytePtr.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRenderingBufferUnlockBytePtr
func CGRenderingBufferUnlockBytePtr(provider CGRenderingBufferProviderRef) {
	if callErr := tryCGRenderingBufferUnlockBytePtr(provider); callErr != nil {
		panic(callErr)
	}
}

var _cGRequestListenEventAccess func() bool
var _cGRequestListenEventAccessErr error

func tryCGRequestListenEventAccess() (bool, error) {
	if _cGRequestListenEventAccess == nil {
		return false, symbolCallError("CGRequestListenEventAccess", "10.15", _cGRequestListenEventAccessErr)
	}
	return _cGRequestListenEventAccess(), nil
}

// CGRequestListenEventAccess.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRequestListenEventAccess()
func CGRequestListenEventAccess() bool {
	result, callErr := tryCGRequestListenEventAccess()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRequestPostEventAccess func() bool
var _cGRequestPostEventAccessErr error

func tryCGRequestPostEventAccess() (bool, error) {
	if _cGRequestPostEventAccess == nil {
		return false, symbolCallError("CGRequestPostEventAccess", "10.15", _cGRequestPostEventAccessErr)
	}
	return _cGRequestPostEventAccess(), nil
}

// CGRequestPostEventAccess.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRequestPostEventAccess()
func CGRequestPostEventAccess() bool {
	result, callErr := tryCGRequestPostEventAccess()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRequestScreenCaptureAccess func() bool
var _cGRequestScreenCaptureAccessErr error

func tryCGRequestScreenCaptureAccess() (bool, error) {
	if _cGRequestScreenCaptureAccess == nil {
		return false, symbolCallError("CGRequestScreenCaptureAccess", "10.15", _cGRequestScreenCaptureAccessErr)
	}
	return _cGRequestScreenCaptureAccess(), nil
}

// CGRequestScreenCaptureAccess.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRequestScreenCaptureAccess()
func CGRequestScreenCaptureAccess() bool {
	result, callErr := tryCGRequestScreenCaptureAccess()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGRestorePermanentDisplayConfiguration func()
var _cGRestorePermanentDisplayConfigurationErr error

func tryCGRestorePermanentDisplayConfiguration() error {
	if _cGRestorePermanentDisplayConfiguration == nil {
		return symbolCallError("CGRestorePermanentDisplayConfiguration", "10.2", _cGRestorePermanentDisplayConfigurationErr)
	}
	_cGRestorePermanentDisplayConfiguration()
	return nil
}

// CGRestorePermanentDisplayConfiguration restores the permanent display configuration settings for the current user.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGRestorePermanentDisplayConfiguration()
func CGRestorePermanentDisplayConfiguration() {
	if callErr := tryCGRestorePermanentDisplayConfiguration(); callErr != nil {
		panic(callErr)
	}
}

var _cGSessionCopyCurrentDictionary func() corefoundation.CFDictionaryRef
var _cGSessionCopyCurrentDictionaryErr error

func tryCGSessionCopyCurrentDictionary() (corefoundation.CFDictionaryRef, error) {
	if _cGSessionCopyCurrentDictionary == nil {
		return 0, symbolCallError("CGSessionCopyCurrentDictionary", "10.3", _cGSessionCopyCurrentDictionaryErr)
	}
	return _cGSessionCopyCurrentDictionary(), nil
}

// CGSessionCopyCurrentDictionary returns information about the caller’s window server session.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGSessionCopyCurrentDictionary()
func CGSessionCopyCurrentDictionary() corefoundation.CFDictionaryRef {
	result, callErr := tryCGSessionCopyCurrentDictionary()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGSetDisplayTransferByByteTable func(display uint32, tableSize uint32, redTable *uint8, greenTable *uint8, blueTable *uint8) CGError
var _cGSetDisplayTransferByByteTableErr error

func tryCGSetDisplayTransferByByteTable(display uint32, tableSize uint32, redTable *uint8, greenTable *uint8, blueTable *uint8) (CGError, error) {
	if _cGSetDisplayTransferByByteTable == nil {
		return *new(CGError), symbolCallError("CGSetDisplayTransferByByteTable", "10.0", _cGSetDisplayTransferByByteTableErr)
	}
	return _cGSetDisplayTransferByByteTable(display, tableSize, redTable, greenTable, blueTable), nil
}

// CGSetDisplayTransferByByteTable sets the byte values in the 8-bit RGB gamma tables for a display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGSetDisplayTransferByByteTable(_:_:_:_:_:)
func CGSetDisplayTransferByByteTable(display uint32, tableSize uint32, redTable *uint8, greenTable *uint8, blueTable *uint8) CGError {
	result, callErr := tryCGSetDisplayTransferByByteTable(display, tableSize, redTable, greenTable, blueTable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGSetDisplayTransferByFormula func(display uint32, redMin CGGammaValue, redMax CGGammaValue, redGamma CGGammaValue, greenMin CGGammaValue, greenMax CGGammaValue, greenGamma CGGammaValue, blueMin CGGammaValue, blueMax CGGammaValue, blueGamma CGGammaValue) CGError
var _cGSetDisplayTransferByFormulaErr error

func tryCGSetDisplayTransferByFormula(display uint32, redMin CGGammaValue, redMax CGGammaValue, redGamma CGGammaValue, greenMin CGGammaValue, greenMax CGGammaValue, greenGamma CGGammaValue, blueMin CGGammaValue, blueMax CGGammaValue, blueGamma CGGammaValue) (CGError, error) {
	if _cGSetDisplayTransferByFormula == nil {
		return *new(CGError), symbolCallError("CGSetDisplayTransferByFormula", "10.0", _cGSetDisplayTransferByFormulaErr)
	}
	return _cGSetDisplayTransferByFormula(display, redMin, redMax, redGamma, greenMin, greenMax, greenGamma, blueMin, blueMax, blueGamma), nil
}

// CGSetDisplayTransferByFormula sets the gamma function for a display by specifying the coefficients of the gamma transfer formula.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGSetDisplayTransferByFormula(_:_:_:_:_:_:_:_:_:_:)
func CGSetDisplayTransferByFormula(display uint32, redMin CGGammaValue, redMax CGGammaValue, redGamma CGGammaValue, greenMin CGGammaValue, greenMax CGGammaValue, greenGamma CGGammaValue, blueMin CGGammaValue, blueMax CGGammaValue, blueGamma CGGammaValue) CGError {
	result, callErr := tryCGSetDisplayTransferByFormula(display, redMin, redMax, redGamma, greenMin, greenMax, greenGamma, blueMin, blueMax, blueGamma)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGSetDisplayTransferByTable func(display uint32, tableSize uint32, redTable *CGGammaValue, greenTable *CGGammaValue, blueTable *CGGammaValue) CGError
var _cGSetDisplayTransferByTableErr error

func tryCGSetDisplayTransferByTable(display uint32, tableSize uint32, redTable *CGGammaValue, greenTable *CGGammaValue, blueTable *CGGammaValue) (CGError, error) {
	if _cGSetDisplayTransferByTable == nil {
		return *new(CGError), symbolCallError("CGSetDisplayTransferByTable", "10.0", _cGSetDisplayTransferByTableErr)
	}
	return _cGSetDisplayTransferByTable(display, tableSize, redTable, greenTable, blueTable), nil
}

// CGSetDisplayTransferByTable sets the color gamma function for a display by specifying the values in the RGB gamma tables.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGSetDisplayTransferByTable(_:_:_:_:_:)
func CGSetDisplayTransferByTable(display uint32, tableSize uint32, redTable *CGGammaValue, greenTable *CGGammaValue, blueTable *CGGammaValue) CGError {
	result, callErr := tryCGSetDisplayTransferByTable(display, tableSize, redTable, greenTable, blueTable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGShadingCreateAxial func(space CGColorSpaceRef, start corefoundation.CGPoint, end corefoundation.CGPoint, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef
var _cGShadingCreateAxialErr error

func tryCGShadingCreateAxial(space CGColorSpaceRef, start corefoundation.CGPoint, end corefoundation.CGPoint, function CGFunctionRef, extendStart bool, extendEnd bool) (CGShadingRef, error) {
	if _cGShadingCreateAxial == nil {
		return 0, symbolCallError("CGShadingCreateAxial", "10.2", _cGShadingCreateAxialErr)
	}
	return _cGShadingCreateAxial(space, start, end, function, extendStart, extendEnd), nil
}

// CGShadingCreateAxial creates a shading object to use for axial shading.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShading/init(axialSpace:start:end:function:extendStart:extendEnd:)
func CGShadingCreateAxial(space CGColorSpaceRef, start corefoundation.CGPoint, end corefoundation.CGPoint, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef {
	result, callErr := tryCGShadingCreateAxial(space, start, end, function, extendStart, extendEnd)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGShadingCreateAxialWithContentHeadroom func(headroom float32, space CGColorSpaceRef, start corefoundation.CGPoint, end corefoundation.CGPoint, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef
var _cGShadingCreateAxialWithContentHeadroomErr error

func tryCGShadingCreateAxialWithContentHeadroom(headroom float32, space CGColorSpaceRef, start corefoundation.CGPoint, end corefoundation.CGPoint, function CGFunctionRef, extendStart bool, extendEnd bool) (CGShadingRef, error) {
	if _cGShadingCreateAxialWithContentHeadroom == nil {
		return 0, symbolCallError("CGShadingCreateAxialWithContentHeadroom", "26.0", _cGShadingCreateAxialWithContentHeadroomErr)
	}
	return _cGShadingCreateAxialWithContentHeadroom(headroom, space, start, end, function, extendStart, extendEnd), nil
}

// CGShadingCreateAxialWithContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShading/init(axialHeadroom:space:start:end:function:extendStart:extendEnd:)
func CGShadingCreateAxialWithContentHeadroom(headroom float32, space CGColorSpaceRef, start corefoundation.CGPoint, end corefoundation.CGPoint, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef {
	result, callErr := tryCGShadingCreateAxialWithContentHeadroom(headroom, space, start, end, function, extendStart, extendEnd)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGShadingCreateRadial func(space CGColorSpaceRef, start corefoundation.CGPoint, startRadius float64, end corefoundation.CGPoint, endRadius float64, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef
var _cGShadingCreateRadialErr error

func tryCGShadingCreateRadial(space CGColorSpaceRef, start corefoundation.CGPoint, startRadius float64, end corefoundation.CGPoint, endRadius float64, function CGFunctionRef, extendStart bool, extendEnd bool) (CGShadingRef, error) {
	if _cGShadingCreateRadial == nil {
		return 0, symbolCallError("CGShadingCreateRadial", "10.2", _cGShadingCreateRadialErr)
	}
	return _cGShadingCreateRadial(space, start, startRadius, end, endRadius, function, extendStart, extendEnd), nil
}

// CGShadingCreateRadial creates a shading object to use for radial shading.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShading/init(radialSpace:start:startRadius:end:endRadius:function:extendStart:extendEnd:)
func CGShadingCreateRadial(space CGColorSpaceRef, start corefoundation.CGPoint, startRadius float64, end corefoundation.CGPoint, endRadius float64, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef {
	result, callErr := tryCGShadingCreateRadial(space, start, startRadius, end, endRadius, function, extendStart, extendEnd)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGShadingCreateRadialWithContentHeadroom func(headroom float32, space CGColorSpaceRef, start corefoundation.CGPoint, startRadius float64, end corefoundation.CGPoint, endRadius float64, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef
var _cGShadingCreateRadialWithContentHeadroomErr error

func tryCGShadingCreateRadialWithContentHeadroom(headroom float32, space CGColorSpaceRef, start corefoundation.CGPoint, startRadius float64, end corefoundation.CGPoint, endRadius float64, function CGFunctionRef, extendStart bool, extendEnd bool) (CGShadingRef, error) {
	if _cGShadingCreateRadialWithContentHeadroom == nil {
		return 0, symbolCallError("CGShadingCreateRadialWithContentHeadroom", "26.0", _cGShadingCreateRadialWithContentHeadroomErr)
	}
	return _cGShadingCreateRadialWithContentHeadroom(headroom, space, start, startRadius, end, endRadius, function, extendStart, extendEnd), nil
}

// CGShadingCreateRadialWithContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShading/init(radialHeadroom:space:start:startRadius:end:endRadius:function:extendStart:extendEnd:)
func CGShadingCreateRadialWithContentHeadroom(headroom float32, space CGColorSpaceRef, start corefoundation.CGPoint, startRadius float64, end corefoundation.CGPoint, endRadius float64, function CGFunctionRef, extendStart bool, extendEnd bool) CGShadingRef {
	result, callErr := tryCGShadingCreateRadialWithContentHeadroom(headroom, space, start, startRadius, end, endRadius, function, extendStart, extendEnd)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGShadingGetContentHeadroom func(shading CGShadingRef) float32
var _cGShadingGetContentHeadroomErr error

func tryCGShadingGetContentHeadroom(shading CGShadingRef) (float32, error) {
	if _cGShadingGetContentHeadroom == nil {
		return 0.0, symbolCallError("CGShadingGetContentHeadroom", "26.0", _cGShadingGetContentHeadroomErr)
	}
	return _cGShadingGetContentHeadroom(shading), nil
}

// CGShadingGetContentHeadroom.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShading/contentHeadroom
func CGShadingGetContentHeadroom(shading CGShadingRef) float32 {
	result, callErr := tryCGShadingGetContentHeadroom(shading)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGShadingGetTypeID func() uint
var _cGShadingGetTypeIDErr error

func tryCGShadingGetTypeID() (uint, error) {
	if _cGShadingGetTypeID == nil {
		return 0, symbolCallError("CGShadingGetTypeID", "10.2", _cGShadingGetTypeIDErr)
	}
	return _cGShadingGetTypeID(), nil
}

// CGShadingGetTypeID returns the Core Foundation type identifier for Core Graphics shading objects.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShading/typeID
func CGShadingGetTypeID() uint {
	result, callErr := tryCGShadingGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGShadingRelease func(shading CGShadingRef)
var _cGShadingReleaseErr error

func tryCGShadingRelease(shading CGShadingRef) error {
	if _cGShadingRelease == nil {
		return symbolCallError("CGShadingRelease", "10.2", _cGShadingReleaseErr)
	}
	_cGShadingRelease(shading)
	return nil
}

// CGShadingRelease decrements the retain count of a shading object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShadingRelease
func CGShadingRelease(shading CGShadingRef) {
	if callErr := tryCGShadingRelease(shading); callErr != nil {
		panic(callErr)
	}
}

var _cGShadingRetain func(shading CGShadingRef) CGShadingRef
var _cGShadingRetainErr error

func tryCGShadingRetain(shading CGShadingRef) (CGShadingRef, error) {
	if _cGShadingRetain == nil {
		return 0, symbolCallError("CGShadingRetain", "10.2", _cGShadingRetainErr)
	}
	return _cGShadingRetain(shading), nil
}

// CGShadingRetain increments the retain count of a shading object.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShadingRetain
func CGShadingRetain(shading CGShadingRef) CGShadingRef {
	result, callErr := tryCGShadingRetain(shading)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGShieldingWindowID func(display uint32) CGWindowID
var _cGShieldingWindowIDErr error

func tryCGShieldingWindowID(display uint32) (CGWindowID, error) {
	if _cGShieldingWindowID == nil {
		return *new(CGWindowID), symbolCallError("CGShieldingWindowID", "10.0", _cGShieldingWindowIDErr)
	}
	return _cGShieldingWindowID(display), nil
}

// CGShieldingWindowID returns the window ID of the shield window for a captured display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShieldingWindowID(_:)
func CGShieldingWindowID(display uint32) CGWindowID {
	result, callErr := tryCGShieldingWindowID(display)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGShieldingWindowLevel func() CGWindowLevel
var _cGShieldingWindowLevelErr error

func tryCGShieldingWindowLevel() (CGWindowLevel, error) {
	if _cGShieldingWindowLevel == nil {
		return *new(CGWindowLevel), symbolCallError("CGShieldingWindowLevel", "10.0", _cGShieldingWindowLevelErr)
	}
	return _cGShieldingWindowLevel(), nil
}

// CGShieldingWindowLevel returns the window level of the shield window for a captured display.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGShieldingWindowLevel()
func CGShieldingWindowLevel() CGWindowLevel {
	result, callErr := tryCGShieldingWindowLevel()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGSizeApplyAffineTransform func(size corefoundation.CGSize, t corefoundation.CGAffineTransform) corefoundation.CGSize
var _cGSizeApplyAffineTransformErr error

func tryCGSizeApplyAffineTransform(size corefoundation.CGSize, t corefoundation.CGAffineTransform) (corefoundation.CGSize, error) {
	if _cGSizeApplyAffineTransform == nil {
		return corefoundation.CGSize{}, symbolCallError("CGSizeApplyAffineTransform", "10.0", _cGSizeApplyAffineTransformErr)
	}
	return _cGSizeApplyAffineTransform(size, t), nil
}

// CGSizeApplyAffineTransform returns the height and width resulting from a transformation of an existing height and width.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGSizeApplyAffineTransform(_:_:)
func CGSizeApplyAffineTransform(size corefoundation.CGSize, t corefoundation.CGAffineTransform) corefoundation.CGSize {
	result, callErr := tryCGSizeApplyAffineTransform(size, t)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGSizeCreateDictionaryRepresentation func(size corefoundation.CGSize) corefoundation.CFDictionaryRef
var _cGSizeCreateDictionaryRepresentationErr error

func tryCGSizeCreateDictionaryRepresentation(size corefoundation.CGSize) (corefoundation.CFDictionaryRef, error) {
	if _cGSizeCreateDictionaryRepresentation == nil {
		return 0, symbolCallError("CGSizeCreateDictionaryRepresentation", "10.5", _cGSizeCreateDictionaryRepresentationErr)
	}
	return _cGSizeCreateDictionaryRepresentation(size), nil
}

// CGSizeCreateDictionaryRepresentation returns a dictionary representation of the specified size.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGSizeCreateDictionaryRepresentation(_:)
func CGSizeCreateDictionaryRepresentation(size corefoundation.CGSize) corefoundation.CFDictionaryRef {
	result, callErr := tryCGSizeCreateDictionaryRepresentation(size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGSizeEqualToSize func(size1 corefoundation.CGSize, size2 corefoundation.CGSize) bool
var _cGSizeEqualToSizeErr error

func tryCGSizeEqualToSize(size1 corefoundation.CGSize, size2 corefoundation.CGSize) (bool, error) {
	if _cGSizeEqualToSize == nil {
		return false, symbolCallError("CGSizeEqualToSize", "10.0", _cGSizeEqualToSizeErr)
	}
	return _cGSizeEqualToSize(size1, size2), nil
}

// CGSizeEqualToSize returns whether two sizes are equal.
//
// Deprecated: The [CGSize](<doc://com.apple.documentation/documentation/CoreFoundation/CGSize>) type adopts the [Equatable] protocol; use the `==` operator instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGSizeEqualToSize(_:_:)
func CGSizeEqualToSize(size1 corefoundation.CGSize, size2 corefoundation.CGSize) bool {
	result, callErr := tryCGSizeEqualToSize(size1, size2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGSizeMakeWithDictionaryRepresentation func(dict corefoundation.CFDictionaryRef, size *corefoundation.CGSize) bool
var _cGSizeMakeWithDictionaryRepresentationErr error

func tryCGSizeMakeWithDictionaryRepresentation(dict corefoundation.CFDictionaryRef, size *corefoundation.CGSize) (bool, error) {
	if _cGSizeMakeWithDictionaryRepresentation == nil {
		return false, symbolCallError("CGSizeMakeWithDictionaryRepresentation", "10.5", _cGSizeMakeWithDictionaryRepresentationErr)
	}
	return _cGSizeMakeWithDictionaryRepresentation(dict, size), nil
}

// CGSizeMakeWithDictionaryRepresentation fills in a size using the contents of the specified dictionary.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGSizeMakeWithDictionaryRepresentation(_:_:)
func CGSizeMakeWithDictionaryRepresentation(dict corefoundation.CFDictionaryRef, size *corefoundation.CGSize) bool {
	result, callErr := tryCGSizeMakeWithDictionaryRepresentation(dict, size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGWarpMouseCursorPosition func(newCursorPosition corefoundation.CGPoint) CGError
var _cGWarpMouseCursorPositionErr error

func tryCGWarpMouseCursorPosition(newCursorPosition corefoundation.CGPoint) (CGError, error) {
	if _cGWarpMouseCursorPosition == nil {
		return *new(CGError), symbolCallError("CGWarpMouseCursorPosition", "10.0", _cGWarpMouseCursorPositionErr)
	}
	return _cGWarpMouseCursorPosition(newCursorPosition), nil
}

// CGWarpMouseCursorPosition moves the mouse cursor without generating events.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGWarpMouseCursorPosition(_:)
func CGWarpMouseCursorPosition(newCursorPosition corefoundation.CGPoint) CGError {
	result, callErr := tryCGWarpMouseCursorPosition(newCursorPosition)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGWindowLevelForKey func(key CGWindowLevelKey) CGWindowLevel
var _cGWindowLevelForKeyErr error

func tryCGWindowLevelForKey(key CGWindowLevelKey) (CGWindowLevel, error) {
	if _cGWindowLevelForKey == nil {
		return *new(CGWindowLevel), symbolCallError("CGWindowLevelForKey", "10.0", _cGWindowLevelForKeyErr)
	}
	return _cGWindowLevelForKey(key), nil
}

// CGWindowLevelForKey returns the window level that corresponds to one of the standard window types.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelForKey(_:)
func CGWindowLevelForKey(key CGWindowLevelKey) CGWindowLevel {
	result, callErr := tryCGWindowLevelForKey(key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGWindowListCopyWindowInfo func(option CGWindowListOption, relativeToWindow CGWindowID) corefoundation.CFArrayRef
var _cGWindowListCopyWindowInfoErr error

func tryCGWindowListCopyWindowInfo(option CGWindowListOption, relativeToWindow CGWindowID) (corefoundation.CFArrayRef, error) {
	if _cGWindowListCopyWindowInfo == nil {
		return 0, symbolCallError("CGWindowListCopyWindowInfo", "10.5", _cGWindowListCopyWindowInfoErr)
	}
	return _cGWindowListCopyWindowInfo(option, relativeToWindow), nil
}

// CGWindowListCopyWindowInfo generates and returns information about the selected windows in the current user session.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGWindowListCopyWindowInfo(_:_:)
func CGWindowListCopyWindowInfo(option CGWindowListOption, relativeToWindow CGWindowID) corefoundation.CFArrayRef {
	result, callErr := tryCGWindowListCopyWindowInfo(option, relativeToWindow)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGWindowListCreate func(option CGWindowListOption, relativeToWindow CGWindowID) corefoundation.CFArrayRef
var _cGWindowListCreateErr error

func tryCGWindowListCreate(option CGWindowListOption, relativeToWindow CGWindowID) (corefoundation.CFArrayRef, error) {
	if _cGWindowListCreate == nil {
		return 0, symbolCallError("CGWindowListCreate", "10.5", _cGWindowListCreateErr)
	}
	return _cGWindowListCreate(option, relativeToWindow), nil
}

// CGWindowListCreate returns the list of window IDs associated with the specified windows in the current user session.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGWindowListCreate
func CGWindowListCreate(option CGWindowListOption, relativeToWindow CGWindowID) corefoundation.CFArrayRef {
	result, callErr := tryCGWindowListCreate(option, relativeToWindow)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGWindowListCreateDescriptionFromArray func(windowArray corefoundation.CFArrayRef) corefoundation.CFArrayRef
var _cGWindowListCreateDescriptionFromArrayErr error

func tryCGWindowListCreateDescriptionFromArray(windowArray corefoundation.CFArrayRef) (corefoundation.CFArrayRef, error) {
	if _cGWindowListCreateDescriptionFromArray == nil {
		return 0, symbolCallError("CGWindowListCreateDescriptionFromArray", "10.5", _cGWindowListCreateDescriptionFromArrayErr)
	}
	return _cGWindowListCreateDescriptionFromArray(windowArray), nil
}

// CGWindowListCreateDescriptionFromArray generates and returns information about windows with the specified window IDs.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGWindowListCreateDescriptionFromArray(_:)
func CGWindowListCreateDescriptionFromArray(windowArray corefoundation.CFArrayRef) corefoundation.CFArrayRef {
	result, callErr := tryCGWindowListCreateDescriptionFromArray(windowArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGWindowListCreateImage func(screenBounds corefoundation.CGRect, listOption CGWindowListOption, windowID CGWindowID, imageOption CGWindowImageOption) CGImageRef
var _cGWindowListCreateImageErr error

func tryCGWindowListCreateImage(screenBounds corefoundation.CGRect, listOption CGWindowListOption, windowID CGWindowID, imageOption CGWindowImageOption) (CGImageRef, error) {
	if _cGWindowListCreateImage == nil {
		return 0, symbolCallError("CGWindowListCreateImage", "", _cGWindowListCreateImageErr)
	}
	return _cGWindowListCreateImage(screenBounds, listOption, windowID, imageOption), nil
}

// CGWindowListCreateImage returns a composite image based on a dynamically generated list of windows.
//
// Deprecated: Please use ScreenCaptureKit instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGWindowListCreateImage(_:_:_:_:)
func CGWindowListCreateImage(screenBounds corefoundation.CGRect, listOption CGWindowListOption, windowID CGWindowID, imageOption CGWindowImageOption) CGImageRef {
	result, callErr := tryCGWindowListCreateImage(screenBounds, listOption, windowID, imageOption)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGWindowListCreateImageFromArray func(screenBounds corefoundation.CGRect, windowArray corefoundation.CFArrayRef, imageOption CGWindowImageOption) CGImageRef
var _cGWindowListCreateImageFromArrayErr error

func tryCGWindowListCreateImageFromArray(screenBounds corefoundation.CGRect, windowArray corefoundation.CFArrayRef, imageOption CGWindowImageOption) (CGImageRef, error) {
	if _cGWindowListCreateImageFromArray == nil {
		return 0, symbolCallError("CGWindowListCreateImageFromArray", "", _cGWindowListCreateImageFromArrayErr)
	}
	return _cGWindowListCreateImageFromArray(screenBounds, windowArray, imageOption), nil
}

// CGWindowListCreateImageFromArray returns a composite image of the specified windows.
//
// Deprecated: Please use ScreenCaptureKit instead.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGImage/init(windowListFromArrayScreenBounds:windowArray:imageOption:)
func CGWindowListCreateImageFromArray(screenBounds corefoundation.CGRect, windowArray corefoundation.CFArrayRef, imageOption CGWindowImageOption) CGImageRef {
	result, callErr := tryCGWindowListCreateImageFromArray(screenBounds, windowArray, imageOption)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGWindowServerCreateServerPort func() corefoundation.CFMachPort
var _cGWindowServerCreateServerPortErr error

func tryCGWindowServerCreateServerPort() (corefoundation.CFMachPort, error) {
	if _cGWindowServerCreateServerPort == nil {
		return *new(corefoundation.CFMachPort), symbolCallError("CGWindowServerCreateServerPort", "10.8", _cGWindowServerCreateServerPortErr)
	}
	return _cGWindowServerCreateServerPort(), nil
}

// CGWindowServerCreateServerPort.
//
// See: https://developer.apple.com/documentation/CoreGraphics/CGWindowServerCreateServerPort()
func CGWindowServerCreateServerPort() corefoundation.CFMachPort {
	result, callErr := tryCGWindowServerCreateServerPort()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_cGAcquireDisplayFadeReservation, &_cGAcquireDisplayFadeReservationErr, frameworkHandle, "CGAcquireDisplayFadeReservation", "10.2")
	registerFunc(&_cGAffineTransformConcat, &_cGAffineTransformConcatErr, frameworkHandle, "CGAffineTransformConcat", "10.0")
	registerFunc(&_cGAffineTransformDecompose, &_cGAffineTransformDecomposeErr, frameworkHandle, "CGAffineTransformDecompose", "13.0")
	registerFunc(&_cGAffineTransformEqualToTransform, &_cGAffineTransformEqualToTransformErr, frameworkHandle, "CGAffineTransformEqualToTransform", "10.4")
	registerFunc(&_cGAffineTransformInvert, &_cGAffineTransformInvertErr, frameworkHandle, "CGAffineTransformInvert", "10.0")
	registerFunc(&_cGAffineTransformIsIdentity, &_cGAffineTransformIsIdentityErr, frameworkHandle, "CGAffineTransformIsIdentity", "10.4")
	registerFunc(&_cGAffineTransformMake, &_cGAffineTransformMakeErr, frameworkHandle, "CGAffineTransformMake", "10.0")
	registerFunc(&_cGAffineTransformMakeRotation, &_cGAffineTransformMakeRotationErr, frameworkHandle, "CGAffineTransformMakeRotation", "10.0")
	registerFunc(&_cGAffineTransformMakeScale, &_cGAffineTransformMakeScaleErr, frameworkHandle, "CGAffineTransformMakeScale", "10.0")
	registerFunc(&_cGAffineTransformMakeTranslation, &_cGAffineTransformMakeTranslationErr, frameworkHandle, "CGAffineTransformMakeTranslation", "10.0")
	registerFunc(&_cGAffineTransformMakeWithComponents, &_cGAffineTransformMakeWithComponentsErr, frameworkHandle, "CGAffineTransformMakeWithComponents", "13.0")
	registerFunc(&_cGAffineTransformRotate, &_cGAffineTransformRotateErr, frameworkHandle, "CGAffineTransformRotate", "10.0")
	registerFunc(&_cGAffineTransformScale, &_cGAffineTransformScaleErr, frameworkHandle, "CGAffineTransformScale", "10.0")
	registerFunc(&_cGAffineTransformTranslate, &_cGAffineTransformTranslateErr, frameworkHandle, "CGAffineTransformTranslate", "10.0")
	registerFunc(&_cGAssociateMouseAndMouseCursorPosition, &_cGAssociateMouseAndMouseCursorPositionErr, frameworkHandle, "CGAssociateMouseAndMouseCursorPosition", "10.0")
	registerFunc(&_cGBeginDisplayConfiguration, &_cGBeginDisplayConfigurationErr, frameworkHandle, "CGBeginDisplayConfiguration", "10.0")
	registerFunc(&_cGBitmapContextCreate, &_cGBitmapContextCreateErr, frameworkHandle, "CGBitmapContextCreate", "10.0")
	registerFunc(&_cGBitmapContextCreateAdaptive, &_cGBitmapContextCreateAdaptiveErr, frameworkHandle, "CGBitmapContextCreateAdaptive", "26.0")
	registerFunc(&_cGBitmapContextCreateImage, &_cGBitmapContextCreateImageErr, frameworkHandle, "CGBitmapContextCreateImage", "10.4")
	registerFunc(&_cGBitmapContextCreateWithData, &_cGBitmapContextCreateWithDataErr, frameworkHandle, "CGBitmapContextCreateWithData", "10.6")
	registerFunc(&_cGBitmapContextGetAlphaInfo, &_cGBitmapContextGetAlphaInfoErr, frameworkHandle, "CGBitmapContextGetAlphaInfo", "10.2")
	registerFunc(&_cGBitmapContextGetBitmapInfo, &_cGBitmapContextGetBitmapInfoErr, frameworkHandle, "CGBitmapContextGetBitmapInfo", "10.4")
	registerFunc(&_cGBitmapContextGetBitsPerComponent, &_cGBitmapContextGetBitsPerComponentErr, frameworkHandle, "CGBitmapContextGetBitsPerComponent", "10.2")
	registerFunc(&_cGBitmapContextGetBitsPerPixel, &_cGBitmapContextGetBitsPerPixelErr, frameworkHandle, "CGBitmapContextGetBitsPerPixel", "10.2")
	registerFunc(&_cGBitmapContextGetBytesPerRow, &_cGBitmapContextGetBytesPerRowErr, frameworkHandle, "CGBitmapContextGetBytesPerRow", "10.2")
	registerFunc(&_cGBitmapContextGetColorSpace, &_cGBitmapContextGetColorSpaceErr, frameworkHandle, "CGBitmapContextGetColorSpace", "10.2")
	registerFunc(&_cGBitmapContextGetData, &_cGBitmapContextGetDataErr, frameworkHandle, "CGBitmapContextGetData", "10.2")
	registerFunc(&_cGBitmapContextGetHeight, &_cGBitmapContextGetHeightErr, frameworkHandle, "CGBitmapContextGetHeight", "10.2")
	registerFunc(&_cGBitmapContextGetWidth, &_cGBitmapContextGetWidthErr, frameworkHandle, "CGBitmapContextGetWidth", "10.2")
	registerFunc(&_cGCancelDisplayConfiguration, &_cGCancelDisplayConfigurationErr, frameworkHandle, "CGCancelDisplayConfiguration", "10.0")
	registerFunc(&_cGCaptureAllDisplays, &_cGCaptureAllDisplaysErr, frameworkHandle, "CGCaptureAllDisplays", "10.0")
	registerFunc(&_cGCaptureAllDisplaysWithOptions, &_cGCaptureAllDisplaysWithOptionsErr, frameworkHandle, "CGCaptureAllDisplaysWithOptions", "10.3")
	registerFunc(&_cGColorConversionInfoConvertData, &_cGColorConversionInfoConvertDataErr, frameworkHandle, "CGColorConversionInfoConvertData", "15.0")
	registerFunc(&_cGColorConversionInfoCreate, &_cGColorConversionInfoCreateErr, frameworkHandle, "CGColorConversionInfoCreate", "10.12")
	registerFunc(&_cGColorConversionInfoCreateForToneMapping, &_cGColorConversionInfoCreateForToneMappingErr, frameworkHandle, "CGColorConversionInfoCreateForToneMapping", "15.0")
	registerFunc(&_cGColorConversionInfoCreateFromList, &_cGColorConversionInfoCreateFromListErr, frameworkHandle, "CGColorConversionInfoCreateFromList", "10.12")
	registerFunc(&_cGColorConversionInfoCreateFromListWithArguments, &_cGColorConversionInfoCreateFromListWithArgumentsErr, frameworkHandle, "CGColorConversionInfoCreateFromListWithArguments", "10.13")
	registerFunc(&_cGColorConversionInfoCreateWithOptions, &_cGColorConversionInfoCreateWithOptionsErr, frameworkHandle, "CGColorConversionInfoCreateWithOptions", "10.14.6")
	registerFunc(&_cGColorConversionInfoGetTypeID, &_cGColorConversionInfoGetTypeIDErr, frameworkHandle, "CGColorConversionInfoGetTypeID", "")
	registerFunc(&_cGColorCreate, &_cGColorCreateErr, frameworkHandle, "CGColorCreate", "10.3")
	registerFunc(&_cGColorCreateCopy, &_cGColorCreateCopyErr, frameworkHandle, "CGColorCreateCopy", "10.3")
	registerFunc(&_cGColorCreateCopyByMatchingToColorSpace, &_cGColorCreateCopyByMatchingToColorSpaceErr, frameworkHandle, "CGColorCreateCopyByMatchingToColorSpace", "10.11")
	registerFunc(&_cGColorCreateCopyWithAlpha, &_cGColorCreateCopyWithAlphaErr, frameworkHandle, "CGColorCreateCopyWithAlpha", "10.3")
	registerFunc(&_cGColorCreateGenericCMYK, &_cGColorCreateGenericCMYKErr, frameworkHandle, "CGColorCreateGenericCMYK", "10.5")
	registerFunc(&_cGColorCreateGenericGray, &_cGColorCreateGenericGrayErr, frameworkHandle, "CGColorCreateGenericGray", "10.5")
	registerFunc(&_cGColorCreateGenericGrayGamma2_2, &_cGColorCreateGenericGrayGamma2_2Err, frameworkHandle, "CGColorCreateGenericGrayGamma2_2", "10.15")
	registerFunc(&_cGColorCreateGenericRGB, &_cGColorCreateGenericRGBErr, frameworkHandle, "CGColorCreateGenericRGB", "10.5")
	registerFunc(&_cGColorCreateSRGB, &_cGColorCreateSRGBErr, frameworkHandle, "CGColorCreateSRGB", "10.15")
	registerFunc(&_cGColorCreateWithContentHeadroom, &_cGColorCreateWithContentHeadroomErr, frameworkHandle, "CGColorCreateWithContentHeadroom", "26.0")
	registerFunc(&_cGColorCreateWithPattern, &_cGColorCreateWithPatternErr, frameworkHandle, "CGColorCreateWithPattern", "10.3")
	registerFunc(&_cGColorEqualToColor, &_cGColorEqualToColorErr, frameworkHandle, "CGColorEqualToColor", "10.3")
	registerFunc(&_cGColorGetAlpha, &_cGColorGetAlphaErr, frameworkHandle, "CGColorGetAlpha", "10.3")
	registerFunc(&_cGColorGetColorSpace, &_cGColorGetColorSpaceErr, frameworkHandle, "CGColorGetColorSpace", "10.3")
	registerFunc(&_cGColorGetComponents, &_cGColorGetComponentsErr, frameworkHandle, "CGColorGetComponents", "10.3")
	registerFunc(&_cGColorGetConstantColor, &_cGColorGetConstantColorErr, frameworkHandle, "CGColorGetConstantColor", "10.5")
	registerFunc(&_cGColorGetContentHeadroom, &_cGColorGetContentHeadroomErr, frameworkHandle, "CGColorGetContentHeadroom", "26.0")
	registerFunc(&_cGColorGetNumberOfComponents, &_cGColorGetNumberOfComponentsErr, frameworkHandle, "CGColorGetNumberOfComponents", "10.3")
	registerFunc(&_cGColorGetPattern, &_cGColorGetPatternErr, frameworkHandle, "CGColorGetPattern", "10.3")
	registerFunc(&_cGColorGetTypeID, &_cGColorGetTypeIDErr, frameworkHandle, "CGColorGetTypeID", "10.3")
	registerFunc(&_cGColorRelease, &_cGColorReleaseErr, frameworkHandle, "CGColorRelease", "10.3")
	registerFunc(&_cGColorRetain, &_cGColorRetainErr, frameworkHandle, "CGColorRetain", "10.3")
	registerFunc(&_cGColorSpaceCopyBaseColorSpace, &_cGColorSpaceCopyBaseColorSpaceErr, frameworkHandle, "CGColorSpaceCopyBaseColorSpace", "15.0")
	registerFunc(&_cGColorSpaceCopyICCData, &_cGColorSpaceCopyICCDataErr, frameworkHandle, "CGColorSpaceCopyICCData", "10.12")
	registerFunc(&_cGColorSpaceCopyName, &_cGColorSpaceCopyNameErr, frameworkHandle, "CGColorSpaceCopyName", "10.6")
	registerFunc(&_cGColorSpaceCopyPropertyList, &_cGColorSpaceCopyPropertyListErr, frameworkHandle, "CGColorSpaceCopyPropertyList", "10.12")
	registerFunc(&_cGColorSpaceCreateCalibratedGray, &_cGColorSpaceCreateCalibratedGrayErr, frameworkHandle, "CGColorSpaceCreateCalibratedGray", "10.0")
	registerFunc(&_cGColorSpaceCreateCalibratedRGB, &_cGColorSpaceCreateCalibratedRGBErr, frameworkHandle, "CGColorSpaceCreateCalibratedRGB", "10.0")
	registerFunc(&_cGColorSpaceCreateCopyWithStandardRange, &_cGColorSpaceCreateCopyWithStandardRangeErr, frameworkHandle, "CGColorSpaceCreateCopyWithStandardRange", "13.0")
	registerFunc(&_cGColorSpaceCreateDeviceCMYK, &_cGColorSpaceCreateDeviceCMYKErr, frameworkHandle, "CGColorSpaceCreateDeviceCMYK", "10.0")
	registerFunc(&_cGColorSpaceCreateDeviceGray, &_cGColorSpaceCreateDeviceGrayErr, frameworkHandle, "CGColorSpaceCreateDeviceGray", "10.0")
	registerFunc(&_cGColorSpaceCreateDeviceRGB, &_cGColorSpaceCreateDeviceRGBErr, frameworkHandle, "CGColorSpaceCreateDeviceRGB", "10.0")
	registerFunc(&_cGColorSpaceCreateExtended, &_cGColorSpaceCreateExtendedErr, frameworkHandle, "CGColorSpaceCreateExtended", "11.0")
	registerFunc(&_cGColorSpaceCreateExtendedLinearized, &_cGColorSpaceCreateExtendedLinearizedErr, frameworkHandle, "CGColorSpaceCreateExtendedLinearized", "11.0")
	registerFunc(&_cGColorSpaceCreateICCBased, &_cGColorSpaceCreateICCBasedErr, frameworkHandle, "CGColorSpaceCreateICCBased", "10.0")
	registerFunc(&_cGColorSpaceCreateIndexed, &_cGColorSpaceCreateIndexedErr, frameworkHandle, "CGColorSpaceCreateIndexed", "10.0")
	registerFunc(&_cGColorSpaceCreateLab, &_cGColorSpaceCreateLabErr, frameworkHandle, "CGColorSpaceCreateLab", "10.0")
	registerFunc(&_cGColorSpaceCreateLinearized, &_cGColorSpaceCreateLinearizedErr, frameworkHandle, "CGColorSpaceCreateLinearized", "11.0")
	registerFunc(&_cGColorSpaceCreatePattern, &_cGColorSpaceCreatePatternErr, frameworkHandle, "CGColorSpaceCreatePattern", "10.0")
	registerFunc(&_cGColorSpaceCreateWithColorSyncProfile, &_cGColorSpaceCreateWithColorSyncProfileErr, frameworkHandle, "CGColorSpaceCreateWithColorSyncProfile", "12.0")
	registerFunc(&_cGColorSpaceCreateWithICCData, &_cGColorSpaceCreateWithICCDataErr, frameworkHandle, "CGColorSpaceCreateWithICCData", "10.12")
	registerFunc(&_cGColorSpaceCreateWithName, &_cGColorSpaceCreateWithNameErr, frameworkHandle, "CGColorSpaceCreateWithName", "10.2")
	registerFunc(&_cGColorSpaceCreateWithPropertyList, &_cGColorSpaceCreateWithPropertyListErr, frameworkHandle, "CGColorSpaceCreateWithPropertyList", "10.12")
	registerFunc(&_cGColorSpaceGetBaseColorSpace, &_cGColorSpaceGetBaseColorSpaceErr, frameworkHandle, "CGColorSpaceGetBaseColorSpace", "10.5")
	registerFunc(&_cGColorSpaceGetColorTable, &_cGColorSpaceGetColorTableErr, frameworkHandle, "CGColorSpaceGetColorTable", "10.5")
	registerFunc(&_cGColorSpaceGetColorTableCount, &_cGColorSpaceGetColorTableCountErr, frameworkHandle, "CGColorSpaceGetColorTableCount", "10.5")
	registerFunc(&_cGColorSpaceGetModel, &_cGColorSpaceGetModelErr, frameworkHandle, "CGColorSpaceGetModel", "10.5")
	registerFunc(&_cGColorSpaceGetName, &_cGColorSpaceGetNameErr, frameworkHandle, "CGColorSpaceGetName", "10.13")
	registerFunc(&_cGColorSpaceGetNumberOfComponents, &_cGColorSpaceGetNumberOfComponentsErr, frameworkHandle, "CGColorSpaceGetNumberOfComponents", "10.0")
	registerFunc(&_cGColorSpaceGetTypeID, &_cGColorSpaceGetTypeIDErr, frameworkHandle, "CGColorSpaceGetTypeID", "10.2")
	registerFunc(&_cGColorSpaceIsHDR, &_cGColorSpaceIsHDRErr, frameworkHandle, "CGColorSpaceIsHDR", "10.15")
	registerFunc(&_cGColorSpaceIsHLGBased, &_cGColorSpaceIsHLGBasedErr, frameworkHandle, "CGColorSpaceIsHLGBased", "12.0")
	registerFunc(&_cGColorSpaceIsPQBased, &_cGColorSpaceIsPQBasedErr, frameworkHandle, "CGColorSpaceIsPQBased", "12.0")
	registerFunc(&_cGColorSpaceIsWideGamutRGB, &_cGColorSpaceIsWideGamutRGBErr, frameworkHandle, "CGColorSpaceIsWideGamutRGB", "10.12")
	registerFunc(&_cGColorSpaceRelease, &_cGColorSpaceReleaseErr, frameworkHandle, "CGColorSpaceRelease", "10.0")
	registerFunc(&_cGColorSpaceRetain, &_cGColorSpaceRetainErr, frameworkHandle, "CGColorSpaceRetain", "10.0")
	registerFunc(&_cGColorSpaceSupportsOutput, &_cGColorSpaceSupportsOutputErr, frameworkHandle, "CGColorSpaceSupportsOutput", "10.12")
	registerFunc(&_cGColorSpaceUsesExtendedRange, &_cGColorSpaceUsesExtendedRangeErr, frameworkHandle, "CGColorSpaceUsesExtendedRange", "10.12")
	registerFunc(&_cGColorSpaceUsesITUR_2100TF, &_cGColorSpaceUsesITUR_2100TFErr, frameworkHandle, "CGColorSpaceUsesITUR_2100TF", "11.0")
	registerFunc(&_cGCompleteDisplayConfiguration, &_cGCompleteDisplayConfigurationErr, frameworkHandle, "CGCompleteDisplayConfiguration", "10.0")
	registerFunc(&_cGConfigureDisplayFadeEffect, &_cGConfigureDisplayFadeEffectErr, frameworkHandle, "CGConfigureDisplayFadeEffect", "10.2")
	registerFunc(&_cGConfigureDisplayMirrorOfDisplay, &_cGConfigureDisplayMirrorOfDisplayErr, frameworkHandle, "CGConfigureDisplayMirrorOfDisplay", "10.2")
	registerFunc(&_cGConfigureDisplayOrigin, &_cGConfigureDisplayOriginErr, frameworkHandle, "CGConfigureDisplayOrigin", "10.0")
	registerFunc(&_cGConfigureDisplayStereoOperation, &_cGConfigureDisplayStereoOperationErr, frameworkHandle, "CGConfigureDisplayStereoOperation", "10.4")
	registerFunc(&_cGConfigureDisplayWithDisplayMode, &_cGConfigureDisplayWithDisplayModeErr, frameworkHandle, "CGConfigureDisplayWithDisplayMode", "10.6")
	registerFunc(&_cGContextAddArc, &_cGContextAddArcErr, frameworkHandle, "CGContextAddArc", "10.0")
	registerFunc(&_cGContextAddArcToPoint, &_cGContextAddArcToPointErr, frameworkHandle, "CGContextAddArcToPoint", "10.0")
	registerFunc(&_cGContextAddCurveToPoint, &_cGContextAddCurveToPointErr, frameworkHandle, "CGContextAddCurveToPoint", "10.0")
	registerFunc(&_cGContextAddEllipseInRect, &_cGContextAddEllipseInRectErr, frameworkHandle, "CGContextAddEllipseInRect", "10.4")
	registerFunc(&_cGContextAddLineToPoint, &_cGContextAddLineToPointErr, frameworkHandle, "CGContextAddLineToPoint", "10.0")
	registerFunc(&_cGContextAddLines, &_cGContextAddLinesErr, frameworkHandle, "CGContextAddLines", "10.0")
	registerFunc(&_cGContextAddPath, &_cGContextAddPathErr, frameworkHandle, "CGContextAddPath", "10.2")
	registerFunc(&_cGContextAddQuadCurveToPoint, &_cGContextAddQuadCurveToPointErr, frameworkHandle, "CGContextAddQuadCurveToPoint", "10.0")
	registerFunc(&_cGContextAddRect, &_cGContextAddRectErr, frameworkHandle, "CGContextAddRect", "10.0")
	registerFunc(&_cGContextAddRects, &_cGContextAddRectsErr, frameworkHandle, "CGContextAddRects", "10.0")
	registerFunc(&_cGContextBeginPage, &_cGContextBeginPageErr, frameworkHandle, "CGContextBeginPage", "10.0")
	registerFunc(&_cGContextBeginPath, &_cGContextBeginPathErr, frameworkHandle, "CGContextBeginPath", "10.0")
	registerFunc(&_cGContextBeginTransparencyLayer, &_cGContextBeginTransparencyLayerErr, frameworkHandle, "CGContextBeginTransparencyLayer", "10.3")
	registerFunc(&_cGContextBeginTransparencyLayerWithRect, &_cGContextBeginTransparencyLayerWithRectErr, frameworkHandle, "CGContextBeginTransparencyLayerWithRect", "10.5")
	registerFunc(&_cGContextClearRect, &_cGContextClearRectErr, frameworkHandle, "CGContextClearRect", "10.0")
	registerFunc(&_cGContextClip, &_cGContextClipErr, frameworkHandle, "CGContextClip", "10.0")
	registerFunc(&_cGContextClipToMask, &_cGContextClipToMaskErr, frameworkHandle, "CGContextClipToMask", "10.4")
	registerFunc(&_cGContextClipToRect, &_cGContextClipToRectErr, frameworkHandle, "CGContextClipToRect", "10.0")
	registerFunc(&_cGContextClipToRects, &_cGContextClipToRectsErr, frameworkHandle, "CGContextClipToRects", "10.0")
	registerFunc(&_cGContextClosePath, &_cGContextClosePathErr, frameworkHandle, "CGContextClosePath", "10.0")
	registerFunc(&_cGContextConcatCTM, &_cGContextConcatCTMErr, frameworkHandle, "CGContextConcatCTM", "10.0")
	registerFunc(&_cGContextConvertPointToDeviceSpace, &_cGContextConvertPointToDeviceSpaceErr, frameworkHandle, "CGContextConvertPointToDeviceSpace", "10.4")
	registerFunc(&_cGContextConvertPointToUserSpace, &_cGContextConvertPointToUserSpaceErr, frameworkHandle, "CGContextConvertPointToUserSpace", "10.4")
	registerFunc(&_cGContextConvertRectToDeviceSpace, &_cGContextConvertRectToDeviceSpaceErr, frameworkHandle, "CGContextConvertRectToDeviceSpace", "10.4")
	registerFunc(&_cGContextConvertRectToUserSpace, &_cGContextConvertRectToUserSpaceErr, frameworkHandle, "CGContextConvertRectToUserSpace", "10.4")
	registerFunc(&_cGContextConvertSizeToDeviceSpace, &_cGContextConvertSizeToDeviceSpaceErr, frameworkHandle, "CGContextConvertSizeToDeviceSpace", "10.4")
	registerFunc(&_cGContextConvertSizeToUserSpace, &_cGContextConvertSizeToUserSpaceErr, frameworkHandle, "CGContextConvertSizeToUserSpace", "10.4")
	registerFunc(&_cGContextCopyPath, &_cGContextCopyPathErr, frameworkHandle, "CGContextCopyPath", "10.2")
	registerFunc(&_cGContextDrawConicGradient, &_cGContextDrawConicGradientErr, frameworkHandle, "CGContextDrawConicGradient", "14.0")
	registerFunc(&_cGContextDrawImage, &_cGContextDrawImageErr, frameworkHandle, "CGContextDrawImage", "10.0")
	registerFunc(&_cGContextDrawImageApplyingToneMapping, &_cGContextDrawImageApplyingToneMappingErr, frameworkHandle, "CGContextDrawImageApplyingToneMapping", "15.0")
	registerFunc(&_cGContextDrawLayerAtPoint, &_cGContextDrawLayerAtPointErr, frameworkHandle, "CGContextDrawLayerAtPoint", "10.4")
	registerFunc(&_cGContextDrawLayerInRect, &_cGContextDrawLayerInRectErr, frameworkHandle, "CGContextDrawLayerInRect", "10.4")
	registerFunc(&_cGContextDrawLinearGradient, &_cGContextDrawLinearGradientErr, frameworkHandle, "CGContextDrawLinearGradient", "10.5")
	registerFunc(&_cGContextDrawPDFDocument, &_cGContextDrawPDFDocumentErr, frameworkHandle, "CGContextDrawPDFDocument", "10.0")
	registerFunc(&_cGContextDrawPDFPage, &_cGContextDrawPDFPageErr, frameworkHandle, "CGContextDrawPDFPage", "10.3")
	registerFunc(&_cGContextDrawPath, &_cGContextDrawPathErr, frameworkHandle, "CGContextDrawPath", "10.0")
	registerFunc(&_cGContextDrawRadialGradient, &_cGContextDrawRadialGradientErr, frameworkHandle, "CGContextDrawRadialGradient", "10.5")
	registerFunc(&_cGContextDrawShading, &_cGContextDrawShadingErr, frameworkHandle, "CGContextDrawShading", "10.2")
	registerFunc(&_cGContextDrawTiledImage, &_cGContextDrawTiledImageErr, frameworkHandle, "CGContextDrawTiledImage", "10.5")
	registerFunc(&_cGContextEOClip, &_cGContextEOClipErr, frameworkHandle, "CGContextEOClip", "10.0")
	registerFunc(&_cGContextEOFillPath, &_cGContextEOFillPathErr, frameworkHandle, "CGContextEOFillPath", "10.0")
	registerFunc(&_cGContextEndPage, &_cGContextEndPageErr, frameworkHandle, "CGContextEndPage", "10.0")
	registerFunc(&_cGContextEndTransparencyLayer, &_cGContextEndTransparencyLayerErr, frameworkHandle, "CGContextEndTransparencyLayer", "10.3")
	registerFunc(&_cGContextFillEllipseInRect, &_cGContextFillEllipseInRectErr, frameworkHandle, "CGContextFillEllipseInRect", "10.4")
	registerFunc(&_cGContextFillPath, &_cGContextFillPathErr, frameworkHandle, "CGContextFillPath", "10.0")
	registerFunc(&_cGContextFillRect, &_cGContextFillRectErr, frameworkHandle, "CGContextFillRect", "10.0")
	registerFunc(&_cGContextFillRects, &_cGContextFillRectsErr, frameworkHandle, "CGContextFillRects", "10.0")
	registerFunc(&_cGContextFlush, &_cGContextFlushErr, frameworkHandle, "CGContextFlush", "10.0")
	registerFunc(&_cGContextGetCTM, &_cGContextGetCTMErr, frameworkHandle, "CGContextGetCTM", "10.0")
	registerFunc(&_cGContextGetClipBoundingBox, &_cGContextGetClipBoundingBoxErr, frameworkHandle, "CGContextGetClipBoundingBox", "10.3")
	registerFunc(&_cGContextGetContentToneMappingInfo, &_cGContextGetContentToneMappingInfoErr, frameworkHandle, "CGContextGetContentToneMappingInfo", "26.0")
	registerFunc(&_cGContextGetEDRTargetHeadroom, &_cGContextGetEDRTargetHeadroomErr, frameworkHandle, "CGContextGetEDRTargetHeadroom", "15.0")
	registerFunc(&_cGContextGetInterpolationQuality, &_cGContextGetInterpolationQualityErr, frameworkHandle, "CGContextGetInterpolationQuality", "10.0")
	registerFunc(&_cGContextGetPathBoundingBox, &_cGContextGetPathBoundingBoxErr, frameworkHandle, "CGContextGetPathBoundingBox", "10.0")
	registerFunc(&_cGContextGetPathCurrentPoint, &_cGContextGetPathCurrentPointErr, frameworkHandle, "CGContextGetPathCurrentPoint", "10.0")
	registerFunc(&_cGContextGetTextMatrix, &_cGContextGetTextMatrixErr, frameworkHandle, "CGContextGetTextMatrix", "10.0")
	registerFunc(&_cGContextGetTextPosition, &_cGContextGetTextPositionErr, frameworkHandle, "CGContextGetTextPosition", "10.0")
	registerFunc(&_cGContextGetTypeID, &_cGContextGetTypeIDErr, frameworkHandle, "CGContextGetTypeID", "10.2")
	registerFunc(&_cGContextGetUserSpaceToDeviceSpaceTransform, &_cGContextGetUserSpaceToDeviceSpaceTransformErr, frameworkHandle, "CGContextGetUserSpaceToDeviceSpaceTransform", "10.4")
	registerFunc(&_cGContextIsPathEmpty, &_cGContextIsPathEmptyErr, frameworkHandle, "CGContextIsPathEmpty", "10.0")
	registerFunc(&_cGContextMoveToPoint, &_cGContextMoveToPointErr, frameworkHandle, "CGContextMoveToPoint", "10.0")
	registerFunc(&_cGContextPathContainsPoint, &_cGContextPathContainsPointErr, frameworkHandle, "CGContextPathContainsPoint", "10.4")
	registerFunc(&_cGContextRelease, &_cGContextReleaseErr, frameworkHandle, "CGContextRelease", "10.0")
	registerFunc(&_cGContextReplacePathWithStrokedPath, &_cGContextReplacePathWithStrokedPathErr, frameworkHandle, "CGContextReplacePathWithStrokedPath", "10.4")
	registerFunc(&_cGContextResetClip, &_cGContextResetClipErr, frameworkHandle, "CGContextResetClip", "")
	registerFunc(&_cGContextRestoreGState, &_cGContextRestoreGStateErr, frameworkHandle, "CGContextRestoreGState", "10.0")
	registerFunc(&_cGContextRetain, &_cGContextRetainErr, frameworkHandle, "CGContextRetain", "10.0")
	registerFunc(&_cGContextRotateCTM, &_cGContextRotateCTMErr, frameworkHandle, "CGContextRotateCTM", "10.0")
	registerFunc(&_cGContextSaveGState, &_cGContextSaveGStateErr, frameworkHandle, "CGContextSaveGState", "10.0")
	registerFunc(&_cGContextScaleCTM, &_cGContextScaleCTMErr, frameworkHandle, "CGContextScaleCTM", "10.0")
	registerFunc(&_cGContextSetAllowsAntialiasing, &_cGContextSetAllowsAntialiasingErr, frameworkHandle, "CGContextSetAllowsAntialiasing", "10.4")
	registerFunc(&_cGContextSetAllowsFontSmoothing, &_cGContextSetAllowsFontSmoothingErr, frameworkHandle, "CGContextSetAllowsFontSmoothing", "10.2")
	registerFunc(&_cGContextSetAllowsFontSubpixelPositioning, &_cGContextSetAllowsFontSubpixelPositioningErr, frameworkHandle, "CGContextSetAllowsFontSubpixelPositioning", "10.5")
	registerFunc(&_cGContextSetAllowsFontSubpixelQuantization, &_cGContextSetAllowsFontSubpixelQuantizationErr, frameworkHandle, "CGContextSetAllowsFontSubpixelQuantization", "10.5")
	registerFunc(&_cGContextSetAlpha, &_cGContextSetAlphaErr, frameworkHandle, "CGContextSetAlpha", "10.0")
	registerFunc(&_cGContextSetBlendMode, &_cGContextSetBlendModeErr, frameworkHandle, "CGContextSetBlendMode", "10.4")
	registerFunc(&_cGContextSetCMYKFillColor, &_cGContextSetCMYKFillColorErr, frameworkHandle, "CGContextSetCMYKFillColor", "10.0")
	registerFunc(&_cGContextSetCMYKStrokeColor, &_cGContextSetCMYKStrokeColorErr, frameworkHandle, "CGContextSetCMYKStrokeColor", "10.0")
	registerFunc(&_cGContextSetCharacterSpacing, &_cGContextSetCharacterSpacingErr, frameworkHandle, "CGContextSetCharacterSpacing", "10.0")
	registerFunc(&_cGContextSetContentToneMappingInfo, &_cGContextSetContentToneMappingInfoErr, frameworkHandle, "CGContextSetContentToneMappingInfo", "26.0")
	registerFunc(&_cGContextSetEDRTargetHeadroom, &_cGContextSetEDRTargetHeadroomErr, frameworkHandle, "CGContextSetEDRTargetHeadroom", "15.0")
	registerFunc(&_cGContextSetFillColor, &_cGContextSetFillColorErr, frameworkHandle, "CGContextSetFillColor", "10.0")
	registerFunc(&_cGContextSetFillColorSpace, &_cGContextSetFillColorSpaceErr, frameworkHandle, "CGContextSetFillColorSpace", "10.0")
	registerFunc(&_cGContextSetFillColorWithColor, &_cGContextSetFillColorWithColorErr, frameworkHandle, "CGContextSetFillColorWithColor", "10.3")
	registerFunc(&_cGContextSetFillPattern, &_cGContextSetFillPatternErr, frameworkHandle, "CGContextSetFillPattern", "10.0")
	registerFunc(&_cGContextSetFlatness, &_cGContextSetFlatnessErr, frameworkHandle, "CGContextSetFlatness", "10.0")
	registerFunc(&_cGContextSetFont, &_cGContextSetFontErr, frameworkHandle, "CGContextSetFont", "10.0")
	registerFunc(&_cGContextSetFontSize, &_cGContextSetFontSizeErr, frameworkHandle, "CGContextSetFontSize", "10.0")
	registerFunc(&_cGContextSetGrayFillColor, &_cGContextSetGrayFillColorErr, frameworkHandle, "CGContextSetGrayFillColor", "10.0")
	registerFunc(&_cGContextSetGrayStrokeColor, &_cGContextSetGrayStrokeColorErr, frameworkHandle, "CGContextSetGrayStrokeColor", "10.0")
	registerFunc(&_cGContextSetInterpolationQuality, &_cGContextSetInterpolationQualityErr, frameworkHandle, "CGContextSetInterpolationQuality", "10.0")
	registerFunc(&_cGContextSetLineCap, &_cGContextSetLineCapErr, frameworkHandle, "CGContextSetLineCap", "10.0")
	registerFunc(&_cGContextSetLineDash, &_cGContextSetLineDashErr, frameworkHandle, "CGContextSetLineDash", "10.0")
	registerFunc(&_cGContextSetLineJoin, &_cGContextSetLineJoinErr, frameworkHandle, "CGContextSetLineJoin", "10.0")
	registerFunc(&_cGContextSetLineWidth, &_cGContextSetLineWidthErr, frameworkHandle, "CGContextSetLineWidth", "10.0")
	registerFunc(&_cGContextSetMiterLimit, &_cGContextSetMiterLimitErr, frameworkHandle, "CGContextSetMiterLimit", "10.0")
	registerFunc(&_cGContextSetPatternPhase, &_cGContextSetPatternPhaseErr, frameworkHandle, "CGContextSetPatternPhase", "10.0")
	registerFunc(&_cGContextSetRGBFillColor, &_cGContextSetRGBFillColorErr, frameworkHandle, "CGContextSetRGBFillColor", "10.0")
	registerFunc(&_cGContextSetRGBStrokeColor, &_cGContextSetRGBStrokeColorErr, frameworkHandle, "CGContextSetRGBStrokeColor", "10.0")
	registerFunc(&_cGContextSetRenderingIntent, &_cGContextSetRenderingIntentErr, frameworkHandle, "CGContextSetRenderingIntent", "10.0")
	registerFunc(&_cGContextSetShadow, &_cGContextSetShadowErr, frameworkHandle, "CGContextSetShadow", "10.3")
	registerFunc(&_cGContextSetShadowWithColor, &_cGContextSetShadowWithColorErr, frameworkHandle, "CGContextSetShadowWithColor", "10.3")
	registerFunc(&_cGContextSetShouldAntialias, &_cGContextSetShouldAntialiasErr, frameworkHandle, "CGContextSetShouldAntialias", "10.0")
	registerFunc(&_cGContextSetShouldSmoothFonts, &_cGContextSetShouldSmoothFontsErr, frameworkHandle, "CGContextSetShouldSmoothFonts", "10.2")
	registerFunc(&_cGContextSetShouldSubpixelPositionFonts, &_cGContextSetShouldSubpixelPositionFontsErr, frameworkHandle, "CGContextSetShouldSubpixelPositionFonts", "10.5")
	registerFunc(&_cGContextSetShouldSubpixelQuantizeFonts, &_cGContextSetShouldSubpixelQuantizeFontsErr, frameworkHandle, "CGContextSetShouldSubpixelQuantizeFonts", "10.5")
	registerFunc(&_cGContextSetStrokeColor, &_cGContextSetStrokeColorErr, frameworkHandle, "CGContextSetStrokeColor", "10.0")
	registerFunc(&_cGContextSetStrokeColorSpace, &_cGContextSetStrokeColorSpaceErr, frameworkHandle, "CGContextSetStrokeColorSpace", "10.0")
	registerFunc(&_cGContextSetStrokeColorWithColor, &_cGContextSetStrokeColorWithColorErr, frameworkHandle, "CGContextSetStrokeColorWithColor", "10.3")
	registerFunc(&_cGContextSetStrokePattern, &_cGContextSetStrokePatternErr, frameworkHandle, "CGContextSetStrokePattern", "10.0")
	registerFunc(&_cGContextSetTextDrawingMode, &_cGContextSetTextDrawingModeErr, frameworkHandle, "CGContextSetTextDrawingMode", "10.0")
	registerFunc(&_cGContextSetTextMatrix, &_cGContextSetTextMatrixErr, frameworkHandle, "CGContextSetTextMatrix", "10.0")
	registerFunc(&_cGContextSetTextPosition, &_cGContextSetTextPositionErr, frameworkHandle, "CGContextSetTextPosition", "10.0")
	registerFunc(&_cGContextShowGlyphsAtPositions, &_cGContextShowGlyphsAtPositionsErr, frameworkHandle, "CGContextShowGlyphsAtPositions", "10.5")
	registerFunc(&_cGContextStrokeEllipseInRect, &_cGContextStrokeEllipseInRectErr, frameworkHandle, "CGContextStrokeEllipseInRect", "10.4")
	registerFunc(&_cGContextStrokeLineSegments, &_cGContextStrokeLineSegmentsErr, frameworkHandle, "CGContextStrokeLineSegments", "10.4")
	registerFunc(&_cGContextStrokePath, &_cGContextStrokePathErr, frameworkHandle, "CGContextStrokePath", "10.0")
	registerFunc(&_cGContextStrokeRect, &_cGContextStrokeRectErr, frameworkHandle, "CGContextStrokeRect", "10.0")
	registerFunc(&_cGContextStrokeRectWithWidth, &_cGContextStrokeRectWithWidthErr, frameworkHandle, "CGContextStrokeRectWithWidth", "10.0")
	registerFunc(&_cGContextSynchronize, &_cGContextSynchronizeErr, frameworkHandle, "CGContextSynchronize", "10.0")
	registerFunc(&_cGContextSynchronizeAttributes, &_cGContextSynchronizeAttributesErr, frameworkHandle, "CGContextSynchronizeAttributes", "26.0")
	registerFunc(&_cGContextTranslateCTM, &_cGContextTranslateCTMErr, frameworkHandle, "CGContextTranslateCTM", "10.0")
	registerFunc(&_cGConvertColorDataWithFormat, &_cGConvertColorDataWithFormatErr, frameworkHandle, "CGConvertColorDataWithFormat", "")
	registerFunc(&_cGDataConsumerCreate, &_cGDataConsumerCreateErr, frameworkHandle, "CGDataConsumerCreate", "10.0")
	registerFunc(&_cGDataConsumerCreateWithCFData, &_cGDataConsumerCreateWithCFDataErr, frameworkHandle, "CGDataConsumerCreateWithCFData", "10.4")
	registerFunc(&_cGDataConsumerCreateWithURL, &_cGDataConsumerCreateWithURLErr, frameworkHandle, "CGDataConsumerCreateWithURL", "10.0")
	registerFunc(&_cGDataConsumerGetTypeID, &_cGDataConsumerGetTypeIDErr, frameworkHandle, "CGDataConsumerGetTypeID", "10.2")
	registerFunc(&_cGDataConsumerRelease, &_cGDataConsumerReleaseErr, frameworkHandle, "CGDataConsumerRelease", "10.0")
	registerFunc(&_cGDataConsumerRetain, &_cGDataConsumerRetainErr, frameworkHandle, "CGDataConsumerRetain", "10.0")
	registerFunc(&_cGDataProviderCopyData, &_cGDataProviderCopyDataErr, frameworkHandle, "CGDataProviderCopyData", "10.3")
	registerFunc(&_cGDataProviderCreateDirect, &_cGDataProviderCreateDirectErr, frameworkHandle, "CGDataProviderCreateDirect", "10.5")
	registerFunc(&_cGDataProviderCreateSequential, &_cGDataProviderCreateSequentialErr, frameworkHandle, "CGDataProviderCreateSequential", "10.5")
	registerFunc(&_cGDataProviderCreateWithCFData, &_cGDataProviderCreateWithCFDataErr, frameworkHandle, "CGDataProviderCreateWithCFData", "10.4")
	registerFunc(&_cGDataProviderCreateWithData, &_cGDataProviderCreateWithDataErr, frameworkHandle, "CGDataProviderCreateWithData", "10.0")
	registerFunc(&_cGDataProviderCreateWithFilename, &_cGDataProviderCreateWithFilenameErr, frameworkHandle, "CGDataProviderCreateWithFilename", "10.0")
	registerFunc(&_cGDataProviderCreateWithURL, &_cGDataProviderCreateWithURLErr, frameworkHandle, "CGDataProviderCreateWithURL", "10.0")
	registerFunc(&_cGDataProviderGetInfo, &_cGDataProviderGetInfoErr, frameworkHandle, "CGDataProviderGetInfo", "10.13")
	registerFunc(&_cGDataProviderGetTypeID, &_cGDataProviderGetTypeIDErr, frameworkHandle, "CGDataProviderGetTypeID", "10.2")
	registerFunc(&_cGDataProviderRelease, &_cGDataProviderReleaseErr, frameworkHandle, "CGDataProviderRelease", "10.0")
	registerFunc(&_cGDataProviderRetain, &_cGDataProviderRetainErr, frameworkHandle, "CGDataProviderRetain", "10.0")
	registerFunc(&_cGDirectDisplayCopyCurrentMetalDevice, &_cGDirectDisplayCopyCurrentMetalDeviceErr, frameworkHandle, "CGDirectDisplayCopyCurrentMetalDevice", "10.11")
	registerFunc(&_cGDisplayBounds, &_cGDisplayBoundsErr, frameworkHandle, "CGDisplayBounds", "10.0")
	registerFunc(&_cGDisplayCapture, &_cGDisplayCaptureErr, frameworkHandle, "CGDisplayCapture", "10.0")
	registerFunc(&_cGDisplayCaptureWithOptions, &_cGDisplayCaptureWithOptionsErr, frameworkHandle, "CGDisplayCaptureWithOptions", "10.3")
	registerFunc(&_cGDisplayCopyAllDisplayModes, &_cGDisplayCopyAllDisplayModesErr, frameworkHandle, "CGDisplayCopyAllDisplayModes", "10.6")
	registerFunc(&_cGDisplayCopyColorSpace, &_cGDisplayCopyColorSpaceErr, frameworkHandle, "CGDisplayCopyColorSpace", "10.5")
	registerFunc(&_cGDisplayCopyDisplayMode, &_cGDisplayCopyDisplayModeErr, frameworkHandle, "CGDisplayCopyDisplayMode", "10.6")
	registerFunc(&_cGDisplayCreateImage, &_cGDisplayCreateImageErr, frameworkHandle, "CGDisplayCreateImage", "")
	registerFunc(&_cGDisplayCreateImageForRect, &_cGDisplayCreateImageForRectErr, frameworkHandle, "CGDisplayCreateImageForRect", "")
	registerFunc(&_cGDisplayFade, &_cGDisplayFadeErr, frameworkHandle, "CGDisplayFade", "10.2")
	registerFunc(&_cGDisplayGammaTableCapacity, &_cGDisplayGammaTableCapacityErr, frameworkHandle, "CGDisplayGammaTableCapacity", "10.3")
	registerFunc(&_cGDisplayGetDrawingContext, &_cGDisplayGetDrawingContextErr, frameworkHandle, "CGDisplayGetDrawingContext", "10.3")
	registerFunc(&_cGDisplayHideCursor, &_cGDisplayHideCursorErr, frameworkHandle, "CGDisplayHideCursor", "10.0")
	registerFunc(&_cGDisplayIDToOpenGLDisplayMask, &_cGDisplayIDToOpenGLDisplayMaskErr, frameworkHandle, "CGDisplayIDToOpenGLDisplayMask", "10.0")
	registerFunc(&_cGDisplayIsActive, &_cGDisplayIsActiveErr, frameworkHandle, "CGDisplayIsActive", "10.2")
	registerFunc(&_cGDisplayIsAlwaysInMirrorSet, &_cGDisplayIsAlwaysInMirrorSetErr, frameworkHandle, "CGDisplayIsAlwaysInMirrorSet", "10.2")
	registerFunc(&_cGDisplayIsAsleep, &_cGDisplayIsAsleepErr, frameworkHandle, "CGDisplayIsAsleep", "10.2")
	registerFunc(&_cGDisplayIsBuiltin, &_cGDisplayIsBuiltinErr, frameworkHandle, "CGDisplayIsBuiltin", "10.2")
	registerFunc(&_cGDisplayIsInHWMirrorSet, &_cGDisplayIsInHWMirrorSetErr, frameworkHandle, "CGDisplayIsInHWMirrorSet", "10.2")
	registerFunc(&_cGDisplayIsInMirrorSet, &_cGDisplayIsInMirrorSetErr, frameworkHandle, "CGDisplayIsInMirrorSet", "10.2")
	registerFunc(&_cGDisplayIsMain, &_cGDisplayIsMainErr, frameworkHandle, "CGDisplayIsMain", "10.2")
	registerFunc(&_cGDisplayIsOnline, &_cGDisplayIsOnlineErr, frameworkHandle, "CGDisplayIsOnline", "10.2")
	registerFunc(&_cGDisplayIsStereo, &_cGDisplayIsStereoErr, frameworkHandle, "CGDisplayIsStereo", "10.4")
	registerFunc(&_cGDisplayMirrorsDisplay, &_cGDisplayMirrorsDisplayErr, frameworkHandle, "CGDisplayMirrorsDisplay", "10.2")
	registerFunc(&_cGDisplayModeGetHeight, &_cGDisplayModeGetHeightErr, frameworkHandle, "CGDisplayModeGetHeight", "10.6")
	registerFunc(&_cGDisplayModeGetIODisplayModeID, &_cGDisplayModeGetIODisplayModeIDErr, frameworkHandle, "CGDisplayModeGetIODisplayModeID", "10.6")
	registerFunc(&_cGDisplayModeGetIOFlags, &_cGDisplayModeGetIOFlagsErr, frameworkHandle, "CGDisplayModeGetIOFlags", "10.6")
	registerFunc(&_cGDisplayModeGetPixelHeight, &_cGDisplayModeGetPixelHeightErr, frameworkHandle, "CGDisplayModeGetPixelHeight", "10.8")
	registerFunc(&_cGDisplayModeGetPixelWidth, &_cGDisplayModeGetPixelWidthErr, frameworkHandle, "CGDisplayModeGetPixelWidth", "10.8")
	registerFunc(&_cGDisplayModeGetRefreshRate, &_cGDisplayModeGetRefreshRateErr, frameworkHandle, "CGDisplayModeGetRefreshRate", "10.6")
	registerFunc(&_cGDisplayModeGetTypeID, &_cGDisplayModeGetTypeIDErr, frameworkHandle, "CGDisplayModeGetTypeID", "10.6")
	registerFunc(&_cGDisplayModeGetWidth, &_cGDisplayModeGetWidthErr, frameworkHandle, "CGDisplayModeGetWidth", "10.6")
	registerFunc(&_cGDisplayModeIsUsableForDesktopGUI, &_cGDisplayModeIsUsableForDesktopGUIErr, frameworkHandle, "CGDisplayModeIsUsableForDesktopGUI", "10.6")
	registerFunc(&_cGDisplayModeRelease, &_cGDisplayModeReleaseErr, frameworkHandle, "CGDisplayModeRelease", "10.6")
	registerFunc(&_cGDisplayModeRetain, &_cGDisplayModeRetainErr, frameworkHandle, "CGDisplayModeRetain", "10.6")
	registerFunc(&_cGDisplayModelNumber, &_cGDisplayModelNumberErr, frameworkHandle, "CGDisplayModelNumber", "10.2")
	registerFunc(&_cGDisplayMoveCursorToPoint, &_cGDisplayMoveCursorToPointErr, frameworkHandle, "CGDisplayMoveCursorToPoint", "10.0")
	registerFunc(&_cGDisplayPixelsHigh, &_cGDisplayPixelsHighErr, frameworkHandle, "CGDisplayPixelsHigh", "10.0")
	registerFunc(&_cGDisplayPixelsWide, &_cGDisplayPixelsWideErr, frameworkHandle, "CGDisplayPixelsWide", "10.0")
	registerFunc(&_cGDisplayPrimaryDisplay, &_cGDisplayPrimaryDisplayErr, frameworkHandle, "CGDisplayPrimaryDisplay", "10.2")
	registerFunc(&_cGDisplayRegisterReconfigurationCallback, &_cGDisplayRegisterReconfigurationCallbackErr, frameworkHandle, "CGDisplayRegisterReconfigurationCallback", "10.3")
	registerFunc(&_cGDisplayRelease, &_cGDisplayReleaseErr, frameworkHandle, "CGDisplayRelease", "10.0")
	registerFunc(&_cGDisplayRemoveReconfigurationCallback, &_cGDisplayRemoveReconfigurationCallbackErr, frameworkHandle, "CGDisplayRemoveReconfigurationCallback", "10.3")
	registerFunc(&_cGDisplayRestoreColorSyncSettings, &_cGDisplayRestoreColorSyncSettingsErr, frameworkHandle, "CGDisplayRestoreColorSyncSettings", "10.0")
	registerFunc(&_cGDisplayRotation, &_cGDisplayRotationErr, frameworkHandle, "CGDisplayRotation", "10.5")
	registerFunc(&_cGDisplayScreenSize, &_cGDisplayScreenSizeErr, frameworkHandle, "CGDisplayScreenSize", "10.3")
	registerFunc(&_cGDisplaySerialNumber, &_cGDisplaySerialNumberErr, frameworkHandle, "CGDisplaySerialNumber", "10.2")
	registerFunc(&_cGDisplaySetDisplayMode, &_cGDisplaySetDisplayModeErr, frameworkHandle, "CGDisplaySetDisplayMode", "10.6")
	registerFunc(&_cGDisplaySetStereoOperation, &_cGDisplaySetStereoOperationErr, frameworkHandle, "CGDisplaySetStereoOperation", "10.4")
	registerFunc(&_cGDisplayShowCursor, &_cGDisplayShowCursorErr, frameworkHandle, "CGDisplayShowCursor", "10.0")
	registerFunc(&_cGDisplayStreamCreate, &_cGDisplayStreamCreateErr, frameworkHandle, "CGDisplayStreamCreate", "")
	registerFunc(&_cGDisplayStreamCreateWithDispatchQueue, &_cGDisplayStreamCreateWithDispatchQueueErr, frameworkHandle, "CGDisplayStreamCreateWithDispatchQueue", "")
	registerFunc(&_cGDisplayStreamGetRunLoopSource, &_cGDisplayStreamGetRunLoopSourceErr, frameworkHandle, "CGDisplayStreamGetRunLoopSource", "")
	registerFunc(&_cGDisplayStreamGetTypeID, &_cGDisplayStreamGetTypeIDErr, frameworkHandle, "CGDisplayStreamGetTypeID", "")
	registerFunc(&_cGDisplayStreamStart, &_cGDisplayStreamStartErr, frameworkHandle, "CGDisplayStreamStart", "")
	registerFunc(&_cGDisplayStreamStop, &_cGDisplayStreamStopErr, frameworkHandle, "CGDisplayStreamStop", "")
	registerFunc(&_cGDisplayStreamUpdateCreateMergedUpdate, &_cGDisplayStreamUpdateCreateMergedUpdateErr, frameworkHandle, "CGDisplayStreamUpdateCreateMergedUpdate", "")
	registerFunc(&_cGDisplayStreamUpdateGetDropCount, &_cGDisplayStreamUpdateGetDropCountErr, frameworkHandle, "CGDisplayStreamUpdateGetDropCount", "")
	registerFunc(&_cGDisplayStreamUpdateGetMovedRectsDelta, &_cGDisplayStreamUpdateGetMovedRectsDeltaErr, frameworkHandle, "CGDisplayStreamUpdateGetMovedRectsDelta", "")
	registerFunc(&_cGDisplayStreamUpdateGetRects, &_cGDisplayStreamUpdateGetRectsErr, frameworkHandle, "CGDisplayStreamUpdateGetRects", "")
	registerFunc(&_cGDisplayStreamUpdateGetTypeID, &_cGDisplayStreamUpdateGetTypeIDErr, frameworkHandle, "CGDisplayStreamUpdateGetTypeID", "")
	registerFunc(&_cGDisplayUnitNumber, &_cGDisplayUnitNumberErr, frameworkHandle, "CGDisplayUnitNumber", "10.2")
	registerFunc(&_cGDisplayUsesOpenGLAcceleration, &_cGDisplayUsesOpenGLAccelerationErr, frameworkHandle, "CGDisplayUsesOpenGLAcceleration", "10.2")
	registerFunc(&_cGDisplayVendorNumber, &_cGDisplayVendorNumberErr, frameworkHandle, "CGDisplayVendorNumber", "10.2")
	registerFunc(&_cGEXRToneMappingGammaGetDefaultOptions, &_cGEXRToneMappingGammaGetDefaultOptionsErr, frameworkHandle, "CGEXRToneMappingGammaGetDefaultOptions", "26.0")
	registerFunc(&_cGErrorSetCallback, &_cGErrorSetCallbackErr, frameworkHandle, "CGErrorSetCallback", "")
	registerFunc(&_cGEventCreate, &_cGEventCreateErr, frameworkHandle, "CGEventCreate", "10.4")
	registerFunc(&_cGEventCreateCopy, &_cGEventCreateCopyErr, frameworkHandle, "CGEventCreateCopy", "10.4")
	registerFunc(&_cGEventCreateData, &_cGEventCreateDataErr, frameworkHandle, "CGEventCreateData", "10.4")
	registerFunc(&_cGEventCreateFromData, &_cGEventCreateFromDataErr, frameworkHandle, "CGEventCreateFromData", "10.4")
	registerFunc(&_cGEventCreateKeyboardEvent, &_cGEventCreateKeyboardEventErr, frameworkHandle, "CGEventCreateKeyboardEvent", "10.4")
	registerFunc(&_cGEventCreateMouseEvent, &_cGEventCreateMouseEventErr, frameworkHandle, "CGEventCreateMouseEvent", "10.4")
	registerFunc(&_cGEventCreateScrollWheelEvent, &_cGEventCreateScrollWheelEventErr, frameworkHandle, "CGEventCreateScrollWheelEvent", "10.5")
	registerFunc(&_cGEventCreateScrollWheelEvent2, &_cGEventCreateScrollWheelEvent2Err, frameworkHandle, "CGEventCreateScrollWheelEvent2", "10.13")
	registerFunc(&_cGEventCreateSourceFromEvent, &_cGEventCreateSourceFromEventErr, frameworkHandle, "CGEventCreateSourceFromEvent", "10.4")
	registerFunc(&_cGEventGetDoubleValueField, &_cGEventGetDoubleValueFieldErr, frameworkHandle, "CGEventGetDoubleValueField", "10.4")
	registerFunc(&_cGEventGetFlags, &_cGEventGetFlagsErr, frameworkHandle, "CGEventGetFlags", "10.4")
	registerFunc(&_cGEventGetIntegerValueField, &_cGEventGetIntegerValueFieldErr, frameworkHandle, "CGEventGetIntegerValueField", "10.4")
	registerFunc(&_cGEventGetLocation, &_cGEventGetLocationErr, frameworkHandle, "CGEventGetLocation", "10.4")
	registerFunc(&_cGEventGetTimestamp, &_cGEventGetTimestampErr, frameworkHandle, "CGEventGetTimestamp", "10.4")
	registerFunc(&_cGEventGetType, &_cGEventGetTypeErr, frameworkHandle, "CGEventGetType", "10.4")
	registerFunc(&_cGEventGetTypeID, &_cGEventGetTypeIDErr, frameworkHandle, "CGEventGetTypeID", "10.4")
	registerFunc(&_cGEventGetUnflippedLocation, &_cGEventGetUnflippedLocationErr, frameworkHandle, "CGEventGetUnflippedLocation", "10.5")
	registerFunc(&_cGEventKeyboardGetUnicodeString, &_cGEventKeyboardGetUnicodeStringErr, frameworkHandle, "CGEventKeyboardGetUnicodeString", "10.4")
	registerFunc(&_cGEventKeyboardSetUnicodeString, &_cGEventKeyboardSetUnicodeStringErr, frameworkHandle, "CGEventKeyboardSetUnicodeString", "10.4")
	registerFunc(&_cGEventPost, &_cGEventPostErr, frameworkHandle, "CGEventPost", "10.4")
	registerFunc(&_cGEventPostToPSN, &_cGEventPostToPSNErr, frameworkHandle, "CGEventPostToPSN", "10.4")
	registerFunc(&_cGEventPostToPid, &_cGEventPostToPidErr, frameworkHandle, "CGEventPostToPid", "10.11")
	registerFunc(&_cGEventSetDoubleValueField, &_cGEventSetDoubleValueFieldErr, frameworkHandle, "CGEventSetDoubleValueField", "10.4")
	registerFunc(&_cGEventSetFlags, &_cGEventSetFlagsErr, frameworkHandle, "CGEventSetFlags", "10.4")
	registerFunc(&_cGEventSetIntegerValueField, &_cGEventSetIntegerValueFieldErr, frameworkHandle, "CGEventSetIntegerValueField", "10.4")
	registerFunc(&_cGEventSetLocation, &_cGEventSetLocationErr, frameworkHandle, "CGEventSetLocation", "10.4")
	registerFunc(&_cGEventSetSource, &_cGEventSetSourceErr, frameworkHandle, "CGEventSetSource", "10.4")
	registerFunc(&_cGEventSetTimestamp, &_cGEventSetTimestampErr, frameworkHandle, "CGEventSetTimestamp", "10.4")
	registerFunc(&_cGEventSetType, &_cGEventSetTypeErr, frameworkHandle, "CGEventSetType", "10.4")
	registerFunc(&_cGEventSourceButtonState, &_cGEventSourceButtonStateErr, frameworkHandle, "CGEventSourceButtonState", "10.4")
	registerFunc(&_cGEventSourceCounterForEventType, &_cGEventSourceCounterForEventTypeErr, frameworkHandle, "CGEventSourceCounterForEventType", "10.4")
	registerFunc(&_cGEventSourceCreate, &_cGEventSourceCreateErr, frameworkHandle, "CGEventSourceCreate", "10.4")
	registerFunc(&_cGEventSourceFlagsState, &_cGEventSourceFlagsStateErr, frameworkHandle, "CGEventSourceFlagsState", "10.4")
	registerFunc(&_cGEventSourceGetKeyboardType, &_cGEventSourceGetKeyboardTypeErr, frameworkHandle, "CGEventSourceGetKeyboardType", "10.4")
	registerFunc(&_cGEventSourceGetLocalEventsFilterDuringSuppressionState, &_cGEventSourceGetLocalEventsFilterDuringSuppressionStateErr, frameworkHandle, "CGEventSourceGetLocalEventsFilterDuringSuppressionState", "10.4")
	registerFunc(&_cGEventSourceGetLocalEventsSuppressionInterval, &_cGEventSourceGetLocalEventsSuppressionIntervalErr, frameworkHandle, "CGEventSourceGetLocalEventsSuppressionInterval", "10.4")
	registerFunc(&_cGEventSourceGetPixelsPerLine, &_cGEventSourceGetPixelsPerLineErr, frameworkHandle, "CGEventSourceGetPixelsPerLine", "10.5")
	registerFunc(&_cGEventSourceGetSourceStateID, &_cGEventSourceGetSourceStateIDErr, frameworkHandle, "CGEventSourceGetSourceStateID", "10.4")
	registerFunc(&_cGEventSourceGetTypeID, &_cGEventSourceGetTypeIDErr, frameworkHandle, "CGEventSourceGetTypeID", "10.4")
	registerFunc(&_cGEventSourceGetUserData, &_cGEventSourceGetUserDataErr, frameworkHandle, "CGEventSourceGetUserData", "10.4")
	registerFunc(&_cGEventSourceKeyState, &_cGEventSourceKeyStateErr, frameworkHandle, "CGEventSourceKeyState", "10.4")
	registerFunc(&_cGEventSourceSecondsSinceLastEventType, &_cGEventSourceSecondsSinceLastEventTypeErr, frameworkHandle, "CGEventSourceSecondsSinceLastEventType", "10.4")
	registerFunc(&_cGEventSourceSetKeyboardType, &_cGEventSourceSetKeyboardTypeErr, frameworkHandle, "CGEventSourceSetKeyboardType", "10.4")
	registerFunc(&_cGEventSourceSetLocalEventsFilterDuringSuppressionState, &_cGEventSourceSetLocalEventsFilterDuringSuppressionStateErr, frameworkHandle, "CGEventSourceSetLocalEventsFilterDuringSuppressionState", "10.4")
	registerFunc(&_cGEventSourceSetLocalEventsSuppressionInterval, &_cGEventSourceSetLocalEventsSuppressionIntervalErr, frameworkHandle, "CGEventSourceSetLocalEventsSuppressionInterval", "10.4")
	registerFunc(&_cGEventSourceSetPixelsPerLine, &_cGEventSourceSetPixelsPerLineErr, frameworkHandle, "CGEventSourceSetPixelsPerLine", "10.5")
	registerFunc(&_cGEventSourceSetUserData, &_cGEventSourceSetUserDataErr, frameworkHandle, "CGEventSourceSetUserData", "10.4")
	registerFunc(&_cGEventTapCreate, &_cGEventTapCreateErr, frameworkHandle, "CGEventTapCreate", "10.4")
	registerFunc(&_cGEventTapCreateForPSN, &_cGEventTapCreateForPSNErr, frameworkHandle, "CGEventTapCreateForPSN", "10.4")
	registerFunc(&_cGEventTapCreateForPid, &_cGEventTapCreateForPidErr, frameworkHandle, "CGEventTapCreateForPid", "10.11")
	registerFunc(&_cGEventTapEnable, &_cGEventTapEnableErr, frameworkHandle, "CGEventTapEnable", "10.4")
	registerFunc(&_cGEventTapIsEnabled, &_cGEventTapIsEnabledErr, frameworkHandle, "CGEventTapIsEnabled", "10.4")
	registerFunc(&_cGEventTapPostEvent, &_cGEventTapPostEventErr, frameworkHandle, "CGEventTapPostEvent", "10.4")
	registerFunc(&_cGFontCanCreatePostScriptSubset, &_cGFontCanCreatePostScriptSubsetErr, frameworkHandle, "CGFontCanCreatePostScriptSubset", "10.4")
	registerFunc(&_cGFontCopyFullName, &_cGFontCopyFullNameErr, frameworkHandle, "CGFontCopyFullName", "10.5")
	registerFunc(&_cGFontCopyGlyphNameForGlyph, &_cGFontCopyGlyphNameForGlyphErr, frameworkHandle, "CGFontCopyGlyphNameForGlyph", "10.5")
	registerFunc(&_cGFontCopyPostScriptName, &_cGFontCopyPostScriptNameErr, frameworkHandle, "CGFontCopyPostScriptName", "10.4")
	registerFunc(&_cGFontCopyTableForTag, &_cGFontCopyTableForTagErr, frameworkHandle, "CGFontCopyTableForTag", "10.5")
	registerFunc(&_cGFontCopyTableTags, &_cGFontCopyTableTagsErr, frameworkHandle, "CGFontCopyTableTags", "10.5")
	registerFunc(&_cGFontCopyVariationAxes, &_cGFontCopyVariationAxesErr, frameworkHandle, "CGFontCopyVariationAxes", "10.4")
	registerFunc(&_cGFontCopyVariations, &_cGFontCopyVariationsErr, frameworkHandle, "CGFontCopyVariations", "10.4")
	registerFunc(&_cGFontCreateCopyWithVariations, &_cGFontCreateCopyWithVariationsErr, frameworkHandle, "CGFontCreateCopyWithVariations", "10.4")
	registerFunc(&_cGFontCreatePostScriptEncoding, &_cGFontCreatePostScriptEncodingErr, frameworkHandle, "CGFontCreatePostScriptEncoding", "10.4")
	registerFunc(&_cGFontCreatePostScriptSubset, &_cGFontCreatePostScriptSubsetErr, frameworkHandle, "CGFontCreatePostScriptSubset", "10.4")
	registerFunc(&_cGFontCreateWithDataProvider, &_cGFontCreateWithDataProviderErr, frameworkHandle, "CGFontCreateWithDataProvider", "10.5")
	registerFunc(&_cGFontCreateWithFontName, &_cGFontCreateWithFontNameErr, frameworkHandle, "CGFontCreateWithFontName", "10.5")
	registerFunc(&_cGFontGetAscent, &_cGFontGetAscentErr, frameworkHandle, "CGFontGetAscent", "10.5")
	registerFunc(&_cGFontGetCapHeight, &_cGFontGetCapHeightErr, frameworkHandle, "CGFontGetCapHeight", "10.5")
	registerFunc(&_cGFontGetDescent, &_cGFontGetDescentErr, frameworkHandle, "CGFontGetDescent", "10.5")
	registerFunc(&_cGFontGetFontBBox, &_cGFontGetFontBBoxErr, frameworkHandle, "CGFontGetFontBBox", "10.5")
	registerFunc(&_cGFontGetGlyphAdvances, &_cGFontGetGlyphAdvancesErr, frameworkHandle, "CGFontGetGlyphAdvances", "10.0")
	registerFunc(&_cGFontGetGlyphBBoxes, &_cGFontGetGlyphBBoxesErr, frameworkHandle, "CGFontGetGlyphBBoxes", "10.5")
	registerFunc(&_cGFontGetGlyphWithGlyphName, &_cGFontGetGlyphWithGlyphNameErr, frameworkHandle, "CGFontGetGlyphWithGlyphName", "10.5")
	registerFunc(&_cGFontGetItalicAngle, &_cGFontGetItalicAngleErr, frameworkHandle, "CGFontGetItalicAngle", "10.5")
	registerFunc(&_cGFontGetLeading, &_cGFontGetLeadingErr, frameworkHandle, "CGFontGetLeading", "10.5")
	registerFunc(&_cGFontGetNumberOfGlyphs, &_cGFontGetNumberOfGlyphsErr, frameworkHandle, "CGFontGetNumberOfGlyphs", "10.0")
	registerFunc(&_cGFontGetStemV, &_cGFontGetStemVErr, frameworkHandle, "CGFontGetStemV", "10.5")
	registerFunc(&_cGFontGetTypeID, &_cGFontGetTypeIDErr, frameworkHandle, "CGFontGetTypeID", "10.2")
	registerFunc(&_cGFontGetUnitsPerEm, &_cGFontGetUnitsPerEmErr, frameworkHandle, "CGFontGetUnitsPerEm", "10.0")
	registerFunc(&_cGFontGetXHeight, &_cGFontGetXHeightErr, frameworkHandle, "CGFontGetXHeight", "10.5")
	registerFunc(&_cGFontRelease, &_cGFontReleaseErr, frameworkHandle, "CGFontRelease", "10.0")
	registerFunc(&_cGFontRetain, &_cGFontRetainErr, frameworkHandle, "CGFontRetain", "10.0")
	registerFunc(&_cGFunctionCreate, &_cGFunctionCreateErr, frameworkHandle, "CGFunctionCreate", "10.2")
	registerFunc(&_cGFunctionGetTypeID, &_cGFunctionGetTypeIDErr, frameworkHandle, "CGFunctionGetTypeID", "10.2")
	registerFunc(&_cGFunctionRelease, &_cGFunctionReleaseErr, frameworkHandle, "CGFunctionRelease", "10.2")
	registerFunc(&_cGFunctionRetain, &_cGFunctionRetainErr, frameworkHandle, "CGFunctionRetain", "10.2")
	registerFunc(&_cGGetActiveDisplayList, &_cGGetActiveDisplayListErr, frameworkHandle, "CGGetActiveDisplayList", "10.0")
	registerFunc(&_cGGetDisplayTransferByFormula, &_cGGetDisplayTransferByFormulaErr, frameworkHandle, "CGGetDisplayTransferByFormula", "10.0")
	registerFunc(&_cGGetDisplayTransferByTable, &_cGGetDisplayTransferByTableErr, frameworkHandle, "CGGetDisplayTransferByTable", "10.0")
	registerFunc(&_cGGetDisplaysWithOpenGLDisplayMask, &_cGGetDisplaysWithOpenGLDisplayMaskErr, frameworkHandle, "CGGetDisplaysWithOpenGLDisplayMask", "10.0")
	registerFunc(&_cGGetDisplaysWithPoint, &_cGGetDisplaysWithPointErr, frameworkHandle, "CGGetDisplaysWithPoint", "10.0")
	registerFunc(&_cGGetDisplaysWithRect, &_cGGetDisplaysWithRectErr, frameworkHandle, "CGGetDisplaysWithRect", "10.0")
	registerFunc(&_cGGetEventTapList, &_cGGetEventTapListErr, frameworkHandle, "CGGetEventTapList", "10.4")
	registerFunc(&_cGGetLastMouseDelta, &_cGGetLastMouseDeltaErr, frameworkHandle, "CGGetLastMouseDelta", "10.0")
	registerFunc(&_cGGetOnlineDisplayList, &_cGGetOnlineDisplayListErr, frameworkHandle, "CGGetOnlineDisplayList", "10.2")
	registerFunc(&_cGGradientCreateWithColorComponents, &_cGGradientCreateWithColorComponentsErr, frameworkHandle, "CGGradientCreateWithColorComponents", "10.5")
	registerFunc(&_cGGradientCreateWithColors, &_cGGradientCreateWithColorsErr, frameworkHandle, "CGGradientCreateWithColors", "10.5")
	registerFunc(&_cGGradientCreateWithContentHeadroom, &_cGGradientCreateWithContentHeadroomErr, frameworkHandle, "CGGradientCreateWithContentHeadroom", "26.0")
	registerFunc(&_cGGradientGetContentHeadroom, &_cGGradientGetContentHeadroomErr, frameworkHandle, "CGGradientGetContentHeadroom", "26.0")
	registerFunc(&_cGGradientGetTypeID, &_cGGradientGetTypeIDErr, frameworkHandle, "CGGradientGetTypeID", "10.5")
	registerFunc(&_cGGradientRelease, &_cGGradientReleaseErr, frameworkHandle, "CGGradientRelease", "10.5")
	registerFunc(&_cGGradientRetain, &_cGGradientRetainErr, frameworkHandle, "CGGradientRetain", "10.5")
	registerFunc(&_cGImageCalculateContentAverageLightLevel, &_cGImageCalculateContentAverageLightLevelErr, frameworkHandle, "CGImageCalculateContentAverageLightLevel", "26.0")
	registerFunc(&_cGImageCalculateContentHeadroom, &_cGImageCalculateContentHeadroomErr, frameworkHandle, "CGImageCalculateContentHeadroom", "26.0")
	registerFunc(&_cGImageContainsImageSpecificToneMappingMetadata, &_cGImageContainsImageSpecificToneMappingMetadataErr, frameworkHandle, "CGImageContainsImageSpecificToneMappingMetadata", "15.0")
	registerFunc(&_cGImageCreate, &_cGImageCreateErr, frameworkHandle, "CGImageCreate", "10.0")
	registerFunc(&_cGImageCreateCopy, &_cGImageCreateCopyErr, frameworkHandle, "CGImageCreateCopy", "10.4")
	registerFunc(&_cGImageCreateCopyWithCalculatedHDRStats, &_cGImageCreateCopyWithCalculatedHDRStatsErr, frameworkHandle, "CGImageCreateCopyWithCalculatedHDRStats", "26.0")
	registerFunc(&_cGImageCreateCopyWithColorSpace, &_cGImageCreateCopyWithColorSpaceErr, frameworkHandle, "CGImageCreateCopyWithColorSpace", "10.3")
	registerFunc(&_cGImageCreateCopyWithContentAverageLightLevel, &_cGImageCreateCopyWithContentAverageLightLevelErr, frameworkHandle, "CGImageCreateCopyWithContentAverageLightLevel", "26.0")
	registerFunc(&_cGImageCreateCopyWithContentHeadroom, &_cGImageCreateCopyWithContentHeadroomErr, frameworkHandle, "CGImageCreateCopyWithContentHeadroom", "15.0")
	registerFunc(&_cGImageCreateWithContentHeadroom, &_cGImageCreateWithContentHeadroomErr, frameworkHandle, "CGImageCreateWithContentHeadroom", "15.0")
	registerFunc(&_cGImageCreateWithImageInRect, &_cGImageCreateWithImageInRectErr, frameworkHandle, "CGImageCreateWithImageInRect", "10.4")
	registerFunc(&_cGImageCreateWithJPEGDataProvider, &_cGImageCreateWithJPEGDataProviderErr, frameworkHandle, "CGImageCreateWithJPEGDataProvider", "10.1")
	registerFunc(&_cGImageCreateWithMask, &_cGImageCreateWithMaskErr, frameworkHandle, "CGImageCreateWithMask", "10.4")
	registerFunc(&_cGImageCreateWithMaskingColors, &_cGImageCreateWithMaskingColorsErr, frameworkHandle, "CGImageCreateWithMaskingColors", "10.4")
	registerFunc(&_cGImageCreateWithPNGDataProvider, &_cGImageCreateWithPNGDataProviderErr, frameworkHandle, "CGImageCreateWithPNGDataProvider", "10.2")
	registerFunc(&_cGImageGetAlphaInfo, &_cGImageGetAlphaInfoErr, frameworkHandle, "CGImageGetAlphaInfo", "10.0")
	registerFunc(&_cGImageGetBitmapInfo, &_cGImageGetBitmapInfoErr, frameworkHandle, "CGImageGetBitmapInfo", "10.4")
	registerFunc(&_cGImageGetBitsPerComponent, &_cGImageGetBitsPerComponentErr, frameworkHandle, "CGImageGetBitsPerComponent", "10.0")
	registerFunc(&_cGImageGetBitsPerPixel, &_cGImageGetBitsPerPixelErr, frameworkHandle, "CGImageGetBitsPerPixel", "10.0")
	registerFunc(&_cGImageGetByteOrderInfo, &_cGImageGetByteOrderInfoErr, frameworkHandle, "CGImageGetByteOrderInfo", "10.14")
	registerFunc(&_cGImageGetBytesPerRow, &_cGImageGetBytesPerRowErr, frameworkHandle, "CGImageGetBytesPerRow", "10.0")
	registerFunc(&_cGImageGetColorSpace, &_cGImageGetColorSpaceErr, frameworkHandle, "CGImageGetColorSpace", "10.0")
	registerFunc(&_cGImageGetContentAverageLightLevel, &_cGImageGetContentAverageLightLevelErr, frameworkHandle, "CGImageGetContentAverageLightLevel", "26.0")
	registerFunc(&_cGImageGetContentHeadroom, &_cGImageGetContentHeadroomErr, frameworkHandle, "CGImageGetContentHeadroom", "15.0")
	registerFunc(&_cGImageGetDataProvider, &_cGImageGetDataProviderErr, frameworkHandle, "CGImageGetDataProvider", "10.0")
	registerFunc(&_cGImageGetDecode, &_cGImageGetDecodeErr, frameworkHandle, "CGImageGetDecode", "10.0")
	registerFunc(&_cGImageGetHeight, &_cGImageGetHeightErr, frameworkHandle, "CGImageGetHeight", "10.0")
	registerFunc(&_cGImageGetPixelFormatInfo, &_cGImageGetPixelFormatInfoErr, frameworkHandle, "CGImageGetPixelFormatInfo", "10.14")
	registerFunc(&_cGImageGetRenderingIntent, &_cGImageGetRenderingIntentErr, frameworkHandle, "CGImageGetRenderingIntent", "10.0")
	registerFunc(&_cGImageGetShouldInterpolate, &_cGImageGetShouldInterpolateErr, frameworkHandle, "CGImageGetShouldInterpolate", "10.0")
	registerFunc(&_cGImageGetTypeID, &_cGImageGetTypeIDErr, frameworkHandle, "CGImageGetTypeID", "10.2")
	registerFunc(&_cGImageGetUTType, &_cGImageGetUTTypeErr, frameworkHandle, "CGImageGetUTType", "10.11")
	registerFunc(&_cGImageGetWidth, &_cGImageGetWidthErr, frameworkHandle, "CGImageGetWidth", "10.0")
	registerFunc(&_cGImageIsMask, &_cGImageIsMaskErr, frameworkHandle, "CGImageIsMask", "10.0")
	registerFunc(&_cGImageMaskCreate, &_cGImageMaskCreateErr, frameworkHandle, "CGImageMaskCreate", "10.0")
	registerFunc(&_cGImageRelease, &_cGImageReleaseErr, frameworkHandle, "CGImageRelease", "10.0")
	registerFunc(&_cGImageRetain, &_cGImageRetainErr, frameworkHandle, "CGImageRetain", "10.0")
	registerFunc(&_cGImageShouldToneMap, &_cGImageShouldToneMapErr, frameworkHandle, "CGImageShouldToneMap", "15.0")
	registerFunc(&_cGLayerCreateWithContext, &_cGLayerCreateWithContextErr, frameworkHandle, "CGLayerCreateWithContext", "10.4")
	registerFunc(&_cGLayerGetContext, &_cGLayerGetContextErr, frameworkHandle, "CGLayerGetContext", "10.4")
	registerFunc(&_cGLayerGetSize, &_cGLayerGetSizeErr, frameworkHandle, "CGLayerGetSize", "10.4")
	registerFunc(&_cGLayerGetTypeID, &_cGLayerGetTypeIDErr, frameworkHandle, "CGLayerGetTypeID", "10.4")
	registerFunc(&_cGLayerRelease, &_cGLayerReleaseErr, frameworkHandle, "CGLayerRelease", "10.4")
	registerFunc(&_cGLayerRetain, &_cGLayerRetainErr, frameworkHandle, "CGLayerRetain", "10.4")
	registerFunc(&_cGMainDisplayID, &_cGMainDisplayIDErr, frameworkHandle, "CGMainDisplayID", "10.2")
	registerFunc(&_cGOpenGLDisplayMaskToDisplayID, &_cGOpenGLDisplayMaskToDisplayIDErr, frameworkHandle, "CGOpenGLDisplayMaskToDisplayID", "10.2")
	registerFunc(&_cGPDFArrayApplyBlock, &_cGPDFArrayApplyBlockErr, frameworkHandle, "CGPDFArrayApplyBlock", "10.14")
	registerFunc(&_cGPDFArrayGetArray, &_cGPDFArrayGetArrayErr, frameworkHandle, "CGPDFArrayGetArray", "10.3")
	registerFunc(&_cGPDFArrayGetBoolean, &_cGPDFArrayGetBooleanErr, frameworkHandle, "CGPDFArrayGetBoolean", "10.3")
	registerFunc(&_cGPDFArrayGetCount, &_cGPDFArrayGetCountErr, frameworkHandle, "CGPDFArrayGetCount", "10.3")
	registerFunc(&_cGPDFArrayGetDictionary, &_cGPDFArrayGetDictionaryErr, frameworkHandle, "CGPDFArrayGetDictionary", "10.3")
	registerFunc(&_cGPDFArrayGetInteger, &_cGPDFArrayGetIntegerErr, frameworkHandle, "CGPDFArrayGetInteger", "10.3")
	registerFunc(&_cGPDFArrayGetName, &_cGPDFArrayGetNameErr, frameworkHandle, "CGPDFArrayGetName", "10.3")
	registerFunc(&_cGPDFArrayGetNull, &_cGPDFArrayGetNullErr, frameworkHandle, "CGPDFArrayGetNull", "10.3")
	registerFunc(&_cGPDFArrayGetNumber, &_cGPDFArrayGetNumberErr, frameworkHandle, "CGPDFArrayGetNumber", "10.3")
	registerFunc(&_cGPDFArrayGetObject, &_cGPDFArrayGetObjectErr, frameworkHandle, "CGPDFArrayGetObject", "10.3")
	registerFunc(&_cGPDFArrayGetStream, &_cGPDFArrayGetStreamErr, frameworkHandle, "CGPDFArrayGetStream", "10.3")
	registerFunc(&_cGPDFArrayGetString, &_cGPDFArrayGetStringErr, frameworkHandle, "CGPDFArrayGetString", "10.3")
	registerFunc(&_cGPDFContentStreamCreateWithPage, &_cGPDFContentStreamCreateWithPageErr, frameworkHandle, "CGPDFContentStreamCreateWithPage", "10.4")
	registerFunc(&_cGPDFContentStreamCreateWithStream, &_cGPDFContentStreamCreateWithStreamErr, frameworkHandle, "CGPDFContentStreamCreateWithStream", "10.4")
	registerFunc(&_cGPDFContentStreamGetResource, &_cGPDFContentStreamGetResourceErr, frameworkHandle, "CGPDFContentStreamGetResource", "10.4")
	registerFunc(&_cGPDFContentStreamGetStreams, &_cGPDFContentStreamGetStreamsErr, frameworkHandle, "CGPDFContentStreamGetStreams", "10.4")
	registerFunc(&_cGPDFContentStreamRelease, &_cGPDFContentStreamReleaseErr, frameworkHandle, "CGPDFContentStreamRelease", "10.4")
	registerFunc(&_cGPDFContentStreamRetain, &_cGPDFContentStreamRetainErr, frameworkHandle, "CGPDFContentStreamRetain", "10.4")
	registerFunc(&_cGPDFContextAddDestinationAtPoint, &_cGPDFContextAddDestinationAtPointErr, frameworkHandle, "CGPDFContextAddDestinationAtPoint", "10.4")
	registerFunc(&_cGPDFContextAddDocumentMetadata, &_cGPDFContextAddDocumentMetadataErr, frameworkHandle, "CGPDFContextAddDocumentMetadata", "10.7")
	registerFunc(&_cGPDFContextBeginPage, &_cGPDFContextBeginPageErr, frameworkHandle, "CGPDFContextBeginPage", "10.4")
	registerFunc(&_cGPDFContextBeginTag, &_cGPDFContextBeginTagErr, frameworkHandle, "CGPDFContextBeginTag", "10.15")
	registerFunc(&_cGPDFContextClose, &_cGPDFContextCloseErr, frameworkHandle, "CGPDFContextClose", "10.5")
	registerFunc(&_cGPDFContextCreate, &_cGPDFContextCreateErr, frameworkHandle, "CGPDFContextCreate", "10.0")
	registerFunc(&_cGPDFContextCreateWithURL, &_cGPDFContextCreateWithURLErr, frameworkHandle, "CGPDFContextCreateWithURL", "10.0")
	registerFunc(&_cGPDFContextEndPage, &_cGPDFContextEndPageErr, frameworkHandle, "CGPDFContextEndPage", "10.4")
	registerFunc(&_cGPDFContextEndTag, &_cGPDFContextEndTagErr, frameworkHandle, "CGPDFContextEndTag", "10.15")
	registerFunc(&_cGPDFContextSetDestinationForRect, &_cGPDFContextSetDestinationForRectErr, frameworkHandle, "CGPDFContextSetDestinationForRect", "10.4")
	registerFunc(&_cGPDFContextSetIDTree, &_cGPDFContextSetIDTreeErr, frameworkHandle, "CGPDFContextSetIDTree", "")
	registerFunc(&_cGPDFContextSetOutline, &_cGPDFContextSetOutlineErr, frameworkHandle, "CGPDFContextSetOutline", "10.13")
	registerFunc(&_cGPDFContextSetPageTagStructureTree, &_cGPDFContextSetPageTagStructureTreeErr, frameworkHandle, "CGPDFContextSetPageTagStructureTree", "")
	registerFunc(&_cGPDFContextSetParentTree, &_cGPDFContextSetParentTreeErr, frameworkHandle, "CGPDFContextSetParentTree", "")
	registerFunc(&_cGPDFContextSetURLForRect, &_cGPDFContextSetURLForRectErr, frameworkHandle, "CGPDFContextSetURLForRect", "10.4")
	registerFunc(&_cGPDFDictionaryApplyBlock, &_cGPDFDictionaryApplyBlockErr, frameworkHandle, "CGPDFDictionaryApplyBlock", "10.14")
	registerFunc(&_cGPDFDictionaryApplyFunction, &_cGPDFDictionaryApplyFunctionErr, frameworkHandle, "CGPDFDictionaryApplyFunction", "10.3")
	registerFunc(&_cGPDFDictionaryGetArray, &_cGPDFDictionaryGetArrayErr, frameworkHandle, "CGPDFDictionaryGetArray", "10.3")
	registerFunc(&_cGPDFDictionaryGetBoolean, &_cGPDFDictionaryGetBooleanErr, frameworkHandle, "CGPDFDictionaryGetBoolean", "10.3")
	registerFunc(&_cGPDFDictionaryGetCount, &_cGPDFDictionaryGetCountErr, frameworkHandle, "CGPDFDictionaryGetCount", "10.3")
	registerFunc(&_cGPDFDictionaryGetDictionary, &_cGPDFDictionaryGetDictionaryErr, frameworkHandle, "CGPDFDictionaryGetDictionary", "10.3")
	registerFunc(&_cGPDFDictionaryGetInteger, &_cGPDFDictionaryGetIntegerErr, frameworkHandle, "CGPDFDictionaryGetInteger", "10.3")
	registerFunc(&_cGPDFDictionaryGetName, &_cGPDFDictionaryGetNameErr, frameworkHandle, "CGPDFDictionaryGetName", "10.3")
	registerFunc(&_cGPDFDictionaryGetNumber, &_cGPDFDictionaryGetNumberErr, frameworkHandle, "CGPDFDictionaryGetNumber", "10.3")
	registerFunc(&_cGPDFDictionaryGetObject, &_cGPDFDictionaryGetObjectErr, frameworkHandle, "CGPDFDictionaryGetObject", "10.3")
	registerFunc(&_cGPDFDictionaryGetStream, &_cGPDFDictionaryGetStreamErr, frameworkHandle, "CGPDFDictionaryGetStream", "10.3")
	registerFunc(&_cGPDFDictionaryGetString, &_cGPDFDictionaryGetStringErr, frameworkHandle, "CGPDFDictionaryGetString", "10.3")
	registerFunc(&_cGPDFDocumentAllowsCopying, &_cGPDFDocumentAllowsCopyingErr, frameworkHandle, "CGPDFDocumentAllowsCopying", "10.2")
	registerFunc(&_cGPDFDocumentAllowsPrinting, &_cGPDFDocumentAllowsPrintingErr, frameworkHandle, "CGPDFDocumentAllowsPrinting", "10.2")
	registerFunc(&_cGPDFDocumentCreateWithProvider, &_cGPDFDocumentCreateWithProviderErr, frameworkHandle, "CGPDFDocumentCreateWithProvider", "10.0")
	registerFunc(&_cGPDFDocumentCreateWithURL, &_cGPDFDocumentCreateWithURLErr, frameworkHandle, "CGPDFDocumentCreateWithURL", "10.0")
	registerFunc(&_cGPDFDocumentGetAccessPermissions, &_cGPDFDocumentGetAccessPermissionsErr, frameworkHandle, "CGPDFDocumentGetAccessPermissions", "10.13")
	registerFunc(&_cGPDFDocumentGetArtBox, &_cGPDFDocumentGetArtBoxErr, frameworkHandle, "CGPDFDocumentGetArtBox", "10.0")
	registerFunc(&_cGPDFDocumentGetBleedBox, &_cGPDFDocumentGetBleedBoxErr, frameworkHandle, "CGPDFDocumentGetBleedBox", "10.0")
	registerFunc(&_cGPDFDocumentGetCatalog, &_cGPDFDocumentGetCatalogErr, frameworkHandle, "CGPDFDocumentGetCatalog", "10.3")
	registerFunc(&_cGPDFDocumentGetCropBox, &_cGPDFDocumentGetCropBoxErr, frameworkHandle, "CGPDFDocumentGetCropBox", "10.0")
	registerFunc(&_cGPDFDocumentGetID, &_cGPDFDocumentGetIDErr, frameworkHandle, "CGPDFDocumentGetID", "10.4")
	registerFunc(&_cGPDFDocumentGetInfo, &_cGPDFDocumentGetInfoErr, frameworkHandle, "CGPDFDocumentGetInfo", "10.4")
	registerFunc(&_cGPDFDocumentGetMediaBox, &_cGPDFDocumentGetMediaBoxErr, frameworkHandle, "CGPDFDocumentGetMediaBox", "10.0")
	registerFunc(&_cGPDFDocumentGetNumberOfPages, &_cGPDFDocumentGetNumberOfPagesErr, frameworkHandle, "CGPDFDocumentGetNumberOfPages", "10.0")
	registerFunc(&_cGPDFDocumentGetOutline, &_cGPDFDocumentGetOutlineErr, frameworkHandle, "CGPDFDocumentGetOutline", "10.13")
	registerFunc(&_cGPDFDocumentGetPage, &_cGPDFDocumentGetPageErr, frameworkHandle, "CGPDFDocumentGetPage", "10.3")
	registerFunc(&_cGPDFDocumentGetRotationAngle, &_cGPDFDocumentGetRotationAngleErr, frameworkHandle, "CGPDFDocumentGetRotationAngle", "10.0")
	registerFunc(&_cGPDFDocumentGetTrimBox, &_cGPDFDocumentGetTrimBoxErr, frameworkHandle, "CGPDFDocumentGetTrimBox", "10.0")
	registerFunc(&_cGPDFDocumentGetTypeID, &_cGPDFDocumentGetTypeIDErr, frameworkHandle, "CGPDFDocumentGetTypeID", "10.2")
	registerFunc(&_cGPDFDocumentGetVersion, &_cGPDFDocumentGetVersionErr, frameworkHandle, "CGPDFDocumentGetVersion", "10.3")
	registerFunc(&_cGPDFDocumentIsEncrypted, &_cGPDFDocumentIsEncryptedErr, frameworkHandle, "CGPDFDocumentIsEncrypted", "10.2")
	registerFunc(&_cGPDFDocumentIsUnlocked, &_cGPDFDocumentIsUnlockedErr, frameworkHandle, "CGPDFDocumentIsUnlocked", "10.2")
	registerFunc(&_cGPDFDocumentRelease, &_cGPDFDocumentReleaseErr, frameworkHandle, "CGPDFDocumentRelease", "10.0")
	registerFunc(&_cGPDFDocumentRetain, &_cGPDFDocumentRetainErr, frameworkHandle, "CGPDFDocumentRetain", "10.0")
	registerFunc(&_cGPDFDocumentUnlockWithPassword, &_cGPDFDocumentUnlockWithPasswordErr, frameworkHandle, "CGPDFDocumentUnlockWithPassword", "10.2")
	registerFunc(&_cGPDFObjectGetType, &_cGPDFObjectGetTypeErr, frameworkHandle, "CGPDFObjectGetType", "10.3")
	registerFunc(&_cGPDFObjectGetValue, &_cGPDFObjectGetValueErr, frameworkHandle, "CGPDFObjectGetValue", "10.3")
	registerFunc(&_cGPDFOperatorTableCreate, &_cGPDFOperatorTableCreateErr, frameworkHandle, "CGPDFOperatorTableCreate", "10.4")
	registerFunc(&_cGPDFOperatorTableRelease, &_cGPDFOperatorTableReleaseErr, frameworkHandle, "CGPDFOperatorTableRelease", "10.4")
	registerFunc(&_cGPDFOperatorTableRetain, &_cGPDFOperatorTableRetainErr, frameworkHandle, "CGPDFOperatorTableRetain", "10.4")
	registerFunc(&_cGPDFOperatorTableSetCallback, &_cGPDFOperatorTableSetCallbackErr, frameworkHandle, "CGPDFOperatorTableSetCallback", "10.4")
	registerFunc(&_cGPDFPageGetBoxRect, &_cGPDFPageGetBoxRectErr, frameworkHandle, "CGPDFPageGetBoxRect", "10.3")
	registerFunc(&_cGPDFPageGetDictionary, &_cGPDFPageGetDictionaryErr, frameworkHandle, "CGPDFPageGetDictionary", "10.3")
	registerFunc(&_cGPDFPageGetDocument, &_cGPDFPageGetDocumentErr, frameworkHandle, "CGPDFPageGetDocument", "10.3")
	registerFunc(&_cGPDFPageGetDrawingTransform, &_cGPDFPageGetDrawingTransformErr, frameworkHandle, "CGPDFPageGetDrawingTransform", "10.3")
	registerFunc(&_cGPDFPageGetPageNumber, &_cGPDFPageGetPageNumberErr, frameworkHandle, "CGPDFPageGetPageNumber", "10.3")
	registerFunc(&_cGPDFPageGetRotationAngle, &_cGPDFPageGetRotationAngleErr, frameworkHandle, "CGPDFPageGetRotationAngle", "10.3")
	registerFunc(&_cGPDFPageGetTypeID, &_cGPDFPageGetTypeIDErr, frameworkHandle, "CGPDFPageGetTypeID", "10.3")
	registerFunc(&_cGPDFPageRelease, &_cGPDFPageReleaseErr, frameworkHandle, "CGPDFPageRelease", "10.3")
	registerFunc(&_cGPDFPageRetain, &_cGPDFPageRetainErr, frameworkHandle, "CGPDFPageRetain", "10.3")
	registerFunc(&_cGPDFScannerCreate, &_cGPDFScannerCreateErr, frameworkHandle, "CGPDFScannerCreate", "10.4")
	registerFunc(&_cGPDFScannerGetContentStream, &_cGPDFScannerGetContentStreamErr, frameworkHandle, "CGPDFScannerGetContentStream", "10.4")
	registerFunc(&_cGPDFScannerPopArray, &_cGPDFScannerPopArrayErr, frameworkHandle, "CGPDFScannerPopArray", "10.4")
	registerFunc(&_cGPDFScannerPopBoolean, &_cGPDFScannerPopBooleanErr, frameworkHandle, "CGPDFScannerPopBoolean", "10.4")
	registerFunc(&_cGPDFScannerPopDictionary, &_cGPDFScannerPopDictionaryErr, frameworkHandle, "CGPDFScannerPopDictionary", "10.4")
	registerFunc(&_cGPDFScannerPopInteger, &_cGPDFScannerPopIntegerErr, frameworkHandle, "CGPDFScannerPopInteger", "10.4")
	registerFunc(&_cGPDFScannerPopName, &_cGPDFScannerPopNameErr, frameworkHandle, "CGPDFScannerPopName", "10.4")
	registerFunc(&_cGPDFScannerPopNumber, &_cGPDFScannerPopNumberErr, frameworkHandle, "CGPDFScannerPopNumber", "10.4")
	registerFunc(&_cGPDFScannerPopObject, &_cGPDFScannerPopObjectErr, frameworkHandle, "CGPDFScannerPopObject", "10.4")
	registerFunc(&_cGPDFScannerPopStream, &_cGPDFScannerPopStreamErr, frameworkHandle, "CGPDFScannerPopStream", "10.4")
	registerFunc(&_cGPDFScannerPopString, &_cGPDFScannerPopStringErr, frameworkHandle, "CGPDFScannerPopString", "10.4")
	registerFunc(&_cGPDFScannerRelease, &_cGPDFScannerReleaseErr, frameworkHandle, "CGPDFScannerRelease", "10.4")
	registerFunc(&_cGPDFScannerRetain, &_cGPDFScannerRetainErr, frameworkHandle, "CGPDFScannerRetain", "10.4")
	registerFunc(&_cGPDFScannerScan, &_cGPDFScannerScanErr, frameworkHandle, "CGPDFScannerScan", "10.4")
	registerFunc(&_cGPDFScannerStop, &_cGPDFScannerStopErr, frameworkHandle, "CGPDFScannerStop", "")
	registerFunc(&_cGPDFStreamCopyData, &_cGPDFStreamCopyDataErr, frameworkHandle, "CGPDFStreamCopyData", "10.3")
	registerFunc(&_cGPDFStreamGetDictionary, &_cGPDFStreamGetDictionaryErr, frameworkHandle, "CGPDFStreamGetDictionary", "10.3")
	registerFunc(&_cGPDFStringCopyDate, &_cGPDFStringCopyDateErr, frameworkHandle, "CGPDFStringCopyDate", "10.4")
	registerFunc(&_cGPDFStringCopyTextString, &_cGPDFStringCopyTextStringErr, frameworkHandle, "CGPDFStringCopyTextString", "10.3")
	registerFunc(&_cGPDFStringGetBytePtr, &_cGPDFStringGetBytePtrErr, frameworkHandle, "CGPDFStringGetBytePtr", "10.3")
	registerFunc(&_cGPDFStringGetLength, &_cGPDFStringGetLengthErr, frameworkHandle, "CGPDFStringGetLength", "10.3")
	registerFunc(&_cGPDFTagTypeGetName, &_cGPDFTagTypeGetNameErr, frameworkHandle, "CGPDFTagTypeGetName", "10.15")
	registerFunc(&_cGPSConverterAbort, &_cGPSConverterAbortErr, frameworkHandle, "CGPSConverterAbort", "10.3")
	registerFunc(&_cGPSConverterConvert, &_cGPSConverterConvertErr, frameworkHandle, "CGPSConverterConvert", "10.3")
	registerFunc(&_cGPSConverterCreate, &_cGPSConverterCreateErr, frameworkHandle, "CGPSConverterCreate", "10.3")
	registerFunc(&_cGPSConverterGetTypeID, &_cGPSConverterGetTypeIDErr, frameworkHandle, "CGPSConverterGetTypeID", "10.3")
	registerFunc(&_cGPSConverterIsConverting, &_cGPSConverterIsConvertingErr, frameworkHandle, "CGPSConverterIsConverting", "10.3")
	registerFunc(&_cGPathAddArc, &_cGPathAddArcErr, frameworkHandle, "CGPathAddArc", "10.2")
	registerFunc(&_cGPathAddArcToPoint, &_cGPathAddArcToPointErr, frameworkHandle, "CGPathAddArcToPoint", "10.2")
	registerFunc(&_cGPathAddCurveToPoint, &_cGPathAddCurveToPointErr, frameworkHandle, "CGPathAddCurveToPoint", "10.2")
	registerFunc(&_cGPathAddEllipseInRect, &_cGPathAddEllipseInRectErr, frameworkHandle, "CGPathAddEllipseInRect", "10.4")
	registerFunc(&_cGPathAddLineToPoint, &_cGPathAddLineToPointErr, frameworkHandle, "CGPathAddLineToPoint", "10.2")
	registerFunc(&_cGPathAddLines, &_cGPathAddLinesErr, frameworkHandle, "CGPathAddLines", "10.2")
	registerFunc(&_cGPathAddPath, &_cGPathAddPathErr, frameworkHandle, "CGPathAddPath", "10.2")
	registerFunc(&_cGPathAddQuadCurveToPoint, &_cGPathAddQuadCurveToPointErr, frameworkHandle, "CGPathAddQuadCurveToPoint", "10.2")
	registerFunc(&_cGPathAddRect, &_cGPathAddRectErr, frameworkHandle, "CGPathAddRect", "10.2")
	registerFunc(&_cGPathAddRects, &_cGPathAddRectsErr, frameworkHandle, "CGPathAddRects", "10.2")
	registerFunc(&_cGPathAddRelativeArc, &_cGPathAddRelativeArcErr, frameworkHandle, "CGPathAddRelativeArc", "10.7")
	registerFunc(&_cGPathAddRoundedRect, &_cGPathAddRoundedRectErr, frameworkHandle, "CGPathAddRoundedRect", "10.9")
	registerFunc(&_cGPathApply, &_cGPathApplyErr, frameworkHandle, "CGPathApply", "10.2")
	registerFunc(&_cGPathApplyWithBlock, &_cGPathApplyWithBlockErr, frameworkHandle, "CGPathApplyWithBlock", "10.13")
	registerFunc(&_cGPathCloseSubpath, &_cGPathCloseSubpathErr, frameworkHandle, "CGPathCloseSubpath", "10.2")
	registerFunc(&_cGPathContainsPoint, &_cGPathContainsPointErr, frameworkHandle, "CGPathContainsPoint", "10.4")
	registerFunc(&_cGPathCreateCopy, &_cGPathCreateCopyErr, frameworkHandle, "CGPathCreateCopy", "10.2")
	registerFunc(&_cGPathCreateCopyByDashingPath, &_cGPathCreateCopyByDashingPathErr, frameworkHandle, "CGPathCreateCopyByDashingPath", "10.7")
	registerFunc(&_cGPathCreateCopyByFlattening, &_cGPathCreateCopyByFlatteningErr, frameworkHandle, "CGPathCreateCopyByFlattening", "13.0")
	registerFunc(&_cGPathCreateCopyByIntersectingPath, &_cGPathCreateCopyByIntersectingPathErr, frameworkHandle, "CGPathCreateCopyByIntersectingPath", "13.0")
	registerFunc(&_cGPathCreateCopyByNormalizing, &_cGPathCreateCopyByNormalizingErr, frameworkHandle, "CGPathCreateCopyByNormalizing", "13.0")
	registerFunc(&_cGPathCreateCopyByStrokingPath, &_cGPathCreateCopyByStrokingPathErr, frameworkHandle, "CGPathCreateCopyByStrokingPath", "10.7")
	registerFunc(&_cGPathCreateCopyBySubtractingPath, &_cGPathCreateCopyBySubtractingPathErr, frameworkHandle, "CGPathCreateCopyBySubtractingPath", "13.0")
	registerFunc(&_cGPathCreateCopyBySymmetricDifferenceOfPath, &_cGPathCreateCopyBySymmetricDifferenceOfPathErr, frameworkHandle, "CGPathCreateCopyBySymmetricDifferenceOfPath", "13.0")
	registerFunc(&_cGPathCreateCopyByTransformingPath, &_cGPathCreateCopyByTransformingPathErr, frameworkHandle, "CGPathCreateCopyByTransformingPath", "10.7")
	registerFunc(&_cGPathCreateCopyByUnioningPath, &_cGPathCreateCopyByUnioningPathErr, frameworkHandle, "CGPathCreateCopyByUnioningPath", "13.0")
	registerFunc(&_cGPathCreateCopyOfLineByIntersectingPath, &_cGPathCreateCopyOfLineByIntersectingPathErr, frameworkHandle, "CGPathCreateCopyOfLineByIntersectingPath", "13.0")
	registerFunc(&_cGPathCreateCopyOfLineBySubtractingPath, &_cGPathCreateCopyOfLineBySubtractingPathErr, frameworkHandle, "CGPathCreateCopyOfLineBySubtractingPath", "13.0")
	registerFunc(&_cGPathCreateMutable, &_cGPathCreateMutableErr, frameworkHandle, "CGPathCreateMutable", "10.2")
	registerFunc(&_cGPathCreateMutableCopy, &_cGPathCreateMutableCopyErr, frameworkHandle, "CGPathCreateMutableCopy", "10.2")
	registerFunc(&_cGPathCreateMutableCopyByTransformingPath, &_cGPathCreateMutableCopyByTransformingPathErr, frameworkHandle, "CGPathCreateMutableCopyByTransformingPath", "10.7")
	registerFunc(&_cGPathCreateSeparateComponents, &_cGPathCreateSeparateComponentsErr, frameworkHandle, "CGPathCreateSeparateComponents", "13.0")
	registerFunc(&_cGPathCreateWithEllipseInRect, &_cGPathCreateWithEllipseInRectErr, frameworkHandle, "CGPathCreateWithEllipseInRect", "10.7")
	registerFunc(&_cGPathCreateWithRect, &_cGPathCreateWithRectErr, frameworkHandle, "CGPathCreateWithRect", "10.5")
	registerFunc(&_cGPathCreateWithRoundedRect, &_cGPathCreateWithRoundedRectErr, frameworkHandle, "CGPathCreateWithRoundedRect", "10.9")
	registerFunc(&_cGPathEqualToPath, &_cGPathEqualToPathErr, frameworkHandle, "CGPathEqualToPath", "10.2")
	registerFunc(&_cGPathGetBoundingBox, &_cGPathGetBoundingBoxErr, frameworkHandle, "CGPathGetBoundingBox", "10.2")
	registerFunc(&_cGPathGetCurrentPoint, &_cGPathGetCurrentPointErr, frameworkHandle, "CGPathGetCurrentPoint", "10.2")
	registerFunc(&_cGPathGetPathBoundingBox, &_cGPathGetPathBoundingBoxErr, frameworkHandle, "CGPathGetPathBoundingBox", "10.6")
	registerFunc(&_cGPathGetTypeID, &_cGPathGetTypeIDErr, frameworkHandle, "CGPathGetTypeID", "10.2")
	registerFunc(&_cGPathIntersectsPath, &_cGPathIntersectsPathErr, frameworkHandle, "CGPathIntersectsPath", "13.0")
	registerFunc(&_cGPathIsEmpty, &_cGPathIsEmptyErr, frameworkHandle, "CGPathIsEmpty", "10.2")
	registerFunc(&_cGPathIsRect, &_cGPathIsRectErr, frameworkHandle, "CGPathIsRect", "10.2")
	registerFunc(&_cGPathMoveToPoint, &_cGPathMoveToPointErr, frameworkHandle, "CGPathMoveToPoint", "10.2")
	registerFunc(&_cGPathRelease, &_cGPathReleaseErr, frameworkHandle, "CGPathRelease", "10.2")
	registerFunc(&_cGPathRetain, &_cGPathRetainErr, frameworkHandle, "CGPathRetain", "10.2")
	registerFunc(&_cGPatternCreate, &_cGPatternCreateErr, frameworkHandle, "CGPatternCreate", "10.0")
	registerFunc(&_cGPatternGetTypeID, &_cGPatternGetTypeIDErr, frameworkHandle, "CGPatternGetTypeID", "10.2")
	registerFunc(&_cGPatternRelease, &_cGPatternReleaseErr, frameworkHandle, "CGPatternRelease", "10.0")
	registerFunc(&_cGPatternRetain, &_cGPatternRetainErr, frameworkHandle, "CGPatternRetain", "10.0")
	registerFunc(&_cGPointApplyAffineTransform, &_cGPointApplyAffineTransformErr, frameworkHandle, "CGPointApplyAffineTransform", "10.0")
	registerFunc(&_cGPointCreateDictionaryRepresentation, &_cGPointCreateDictionaryRepresentationErr, frameworkHandle, "CGPointCreateDictionaryRepresentation", "10.5")
	registerFunc(&_cGPointEqualToPoint, &_cGPointEqualToPointErr, frameworkHandle, "CGPointEqualToPoint", "10.0")
	registerFunc(&_cGPointMakeWithDictionaryRepresentation, &_cGPointMakeWithDictionaryRepresentationErr, frameworkHandle, "CGPointMakeWithDictionaryRepresentation", "10.5")
	registerFunc(&_cGPostMouseEvent, &_cGPostMouseEventErr, frameworkHandle, "CGPostMouseEvent", "10.0")
	registerFunc(&_cGPostScrollWheelEvent, &_cGPostScrollWheelEventErr, frameworkHandle, "CGPostScrollWheelEvent", "10.0")
	registerFunc(&_cGPreflightListenEventAccess, &_cGPreflightListenEventAccessErr, frameworkHandle, "CGPreflightListenEventAccess", "10.15")
	registerFunc(&_cGPreflightPostEventAccess, &_cGPreflightPostEventAccessErr, frameworkHandle, "CGPreflightPostEventAccess", "10.15")
	registerFunc(&_cGPreflightScreenCaptureAccess, &_cGPreflightScreenCaptureAccessErr, frameworkHandle, "CGPreflightScreenCaptureAccess", "10.15")
	registerFunc(&_cGRectApplyAffineTransform, &_cGRectApplyAffineTransformErr, frameworkHandle, "CGRectApplyAffineTransform", "10.4")
	registerFunc(&_cGRectContainsPoint, &_cGRectContainsPointErr, frameworkHandle, "CGRectContainsPoint", "10.0")
	registerFunc(&_cGRectContainsRect, &_cGRectContainsRectErr, frameworkHandle, "CGRectContainsRect", "10.0")
	registerFunc(&_cGRectCreateDictionaryRepresentation, &_cGRectCreateDictionaryRepresentationErr, frameworkHandle, "CGRectCreateDictionaryRepresentation", "10.5")
	registerFunc(&_cGRectDivide, &_cGRectDivideErr, frameworkHandle, "CGRectDivide", "10.0")
	registerFunc(&_cGRectEqualToRect, &_cGRectEqualToRectErr, frameworkHandle, "CGRectEqualToRect", "10.0")
	registerFunc(&_cGRectGetHeight, &_cGRectGetHeightErr, frameworkHandle, "CGRectGetHeight", "10.0")
	registerFunc(&_cGRectGetMaxX, &_cGRectGetMaxXErr, frameworkHandle, "CGRectGetMaxX", "10.0")
	registerFunc(&_cGRectGetMaxY, &_cGRectGetMaxYErr, frameworkHandle, "CGRectGetMaxY", "10.0")
	registerFunc(&_cGRectGetMidX, &_cGRectGetMidXErr, frameworkHandle, "CGRectGetMidX", "10.0")
	registerFunc(&_cGRectGetMidY, &_cGRectGetMidYErr, frameworkHandle, "CGRectGetMidY", "10.0")
	registerFunc(&_cGRectGetMinX, &_cGRectGetMinXErr, frameworkHandle, "CGRectGetMinX", "10.0")
	registerFunc(&_cGRectGetMinY, &_cGRectGetMinYErr, frameworkHandle, "CGRectGetMinY", "10.0")
	registerFunc(&_cGRectGetWidth, &_cGRectGetWidthErr, frameworkHandle, "CGRectGetWidth", "10.0")
	registerFunc(&_cGRectInset, &_cGRectInsetErr, frameworkHandle, "CGRectInset", "10.0")
	registerFunc(&_cGRectIntegral, &_cGRectIntegralErr, frameworkHandle, "CGRectIntegral", "10.0")
	registerFunc(&_cGRectIntersection, &_cGRectIntersectionErr, frameworkHandle, "CGRectIntersection", "10.0")
	registerFunc(&_cGRectIntersectsRect, &_cGRectIntersectsRectErr, frameworkHandle, "CGRectIntersectsRect", "10.0")
	registerFunc(&_cGRectIsEmpty, &_cGRectIsEmptyErr, frameworkHandle, "CGRectIsEmpty", "10.0")
	registerFunc(&_cGRectIsInfinite, &_cGRectIsInfiniteErr, frameworkHandle, "CGRectIsInfinite", "10.4")
	registerFunc(&_cGRectIsNull, &_cGRectIsNullErr, frameworkHandle, "CGRectIsNull", "10.0")
	registerFunc(&_cGRectMakeWithDictionaryRepresentation, &_cGRectMakeWithDictionaryRepresentationErr, frameworkHandle, "CGRectMakeWithDictionaryRepresentation", "10.5")
	registerFunc(&_cGRectOffset, &_cGRectOffsetErr, frameworkHandle, "CGRectOffset", "10.0")
	registerFunc(&_cGRectStandardize, &_cGRectStandardizeErr, frameworkHandle, "CGRectStandardize", "10.0")
	registerFunc(&_cGRectUnion, &_cGRectUnionErr, frameworkHandle, "CGRectUnion", "10.0")
	registerFunc(&_cGReleaseAllDisplays, &_cGReleaseAllDisplaysErr, frameworkHandle, "CGReleaseAllDisplays", "10.0")
	registerFunc(&_cGReleaseDisplayFadeReservation, &_cGReleaseDisplayFadeReservationErr, frameworkHandle, "CGReleaseDisplayFadeReservation", "10.2")
	registerFunc(&_cGRenderingBufferLockBytePtr, &_cGRenderingBufferLockBytePtrErr, frameworkHandle, "CGRenderingBufferLockBytePtr", "26.0")
	registerFunc(&_cGRenderingBufferProviderCreate, &_cGRenderingBufferProviderCreateErr, frameworkHandle, "CGRenderingBufferProviderCreate", "26.0")
	registerFunc(&_cGRenderingBufferProviderCreateWithCFData, &_cGRenderingBufferProviderCreateWithCFDataErr, frameworkHandle, "CGRenderingBufferProviderCreateWithCFData", "26.0")
	registerFunc(&_cGRenderingBufferProviderGetSize, &_cGRenderingBufferProviderGetSizeErr, frameworkHandle, "CGRenderingBufferProviderGetSize", "26.0")
	registerFunc(&_cGRenderingBufferProviderGetTypeID, &_cGRenderingBufferProviderGetTypeIDErr, frameworkHandle, "CGRenderingBufferProviderGetTypeID", "26.0")
	registerFunc(&_cGRenderingBufferUnlockBytePtr, &_cGRenderingBufferUnlockBytePtrErr, frameworkHandle, "CGRenderingBufferUnlockBytePtr", "26.0")
	registerFunc(&_cGRequestListenEventAccess, &_cGRequestListenEventAccessErr, frameworkHandle, "CGRequestListenEventAccess", "10.15")
	registerFunc(&_cGRequestPostEventAccess, &_cGRequestPostEventAccessErr, frameworkHandle, "CGRequestPostEventAccess", "10.15")
	registerFunc(&_cGRequestScreenCaptureAccess, &_cGRequestScreenCaptureAccessErr, frameworkHandle, "CGRequestScreenCaptureAccess", "10.15")
	registerFunc(&_cGRestorePermanentDisplayConfiguration, &_cGRestorePermanentDisplayConfigurationErr, frameworkHandle, "CGRestorePermanentDisplayConfiguration", "10.2")
	registerFunc(&_cGSessionCopyCurrentDictionary, &_cGSessionCopyCurrentDictionaryErr, frameworkHandle, "CGSessionCopyCurrentDictionary", "10.3")
	registerFunc(&_cGSetDisplayTransferByByteTable, &_cGSetDisplayTransferByByteTableErr, frameworkHandle, "CGSetDisplayTransferByByteTable", "10.0")
	registerFunc(&_cGSetDisplayTransferByFormula, &_cGSetDisplayTransferByFormulaErr, frameworkHandle, "CGSetDisplayTransferByFormula", "10.0")
	registerFunc(&_cGSetDisplayTransferByTable, &_cGSetDisplayTransferByTableErr, frameworkHandle, "CGSetDisplayTransferByTable", "10.0")
	registerFunc(&_cGShadingCreateAxial, &_cGShadingCreateAxialErr, frameworkHandle, "CGShadingCreateAxial", "10.2")
	registerFunc(&_cGShadingCreateAxialWithContentHeadroom, &_cGShadingCreateAxialWithContentHeadroomErr, frameworkHandle, "CGShadingCreateAxialWithContentHeadroom", "26.0")
	registerFunc(&_cGShadingCreateRadial, &_cGShadingCreateRadialErr, frameworkHandle, "CGShadingCreateRadial", "10.2")
	registerFunc(&_cGShadingCreateRadialWithContentHeadroom, &_cGShadingCreateRadialWithContentHeadroomErr, frameworkHandle, "CGShadingCreateRadialWithContentHeadroom", "26.0")
	registerFunc(&_cGShadingGetContentHeadroom, &_cGShadingGetContentHeadroomErr, frameworkHandle, "CGShadingGetContentHeadroom", "26.0")
	registerFunc(&_cGShadingGetTypeID, &_cGShadingGetTypeIDErr, frameworkHandle, "CGShadingGetTypeID", "10.2")
	registerFunc(&_cGShadingRelease, &_cGShadingReleaseErr, frameworkHandle, "CGShadingRelease", "10.2")
	registerFunc(&_cGShadingRetain, &_cGShadingRetainErr, frameworkHandle, "CGShadingRetain", "10.2")
	registerFunc(&_cGShieldingWindowID, &_cGShieldingWindowIDErr, frameworkHandle, "CGShieldingWindowID", "10.0")
	registerFunc(&_cGShieldingWindowLevel, &_cGShieldingWindowLevelErr, frameworkHandle, "CGShieldingWindowLevel", "10.0")
	registerFunc(&_cGSizeApplyAffineTransform, &_cGSizeApplyAffineTransformErr, frameworkHandle, "CGSizeApplyAffineTransform", "10.0")
	registerFunc(&_cGSizeCreateDictionaryRepresentation, &_cGSizeCreateDictionaryRepresentationErr, frameworkHandle, "CGSizeCreateDictionaryRepresentation", "10.5")
	registerFunc(&_cGSizeEqualToSize, &_cGSizeEqualToSizeErr, frameworkHandle, "CGSizeEqualToSize", "10.0")
	registerFunc(&_cGSizeMakeWithDictionaryRepresentation, &_cGSizeMakeWithDictionaryRepresentationErr, frameworkHandle, "CGSizeMakeWithDictionaryRepresentation", "10.5")
	registerFunc(&_cGWarpMouseCursorPosition, &_cGWarpMouseCursorPositionErr, frameworkHandle, "CGWarpMouseCursorPosition", "10.0")
	registerFunc(&_cGWindowLevelForKey, &_cGWindowLevelForKeyErr, frameworkHandle, "CGWindowLevelForKey", "10.0")
	registerFunc(&_cGWindowListCopyWindowInfo, &_cGWindowListCopyWindowInfoErr, frameworkHandle, "CGWindowListCopyWindowInfo", "10.5")
	registerFunc(&_cGWindowListCreate, &_cGWindowListCreateErr, frameworkHandle, "CGWindowListCreate", "10.5")
	registerFunc(&_cGWindowListCreateDescriptionFromArray, &_cGWindowListCreateDescriptionFromArrayErr, frameworkHandle, "CGWindowListCreateDescriptionFromArray", "10.5")
	registerFunc(&_cGWindowListCreateImage, &_cGWindowListCreateImageErr, frameworkHandle, "CGWindowListCreateImage", "")
	registerFunc(&_cGWindowListCreateImageFromArray, &_cGWindowListCreateImageFromArrayErr, frameworkHandle, "CGWindowListCreateImageFromArray", "")
	registerFunc(&_cGWindowServerCreateServerPort, &_cGWindowServerCreateServerPortErr, frameworkHandle, "CGWindowServerCreateServerPort", "10.8")
}

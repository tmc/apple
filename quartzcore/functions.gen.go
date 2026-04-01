// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"fmt"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
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
		return fmt.Sprintf("QuartzCore: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("QuartzCore: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("QuartzCore: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("QuartzCore: register symbol %s: %v", name, r)
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

var _cACurrentMediaTime func() float64
var _cACurrentMediaTimeErr error

func tryCACurrentMediaTime() (float64, error) {
	if _cACurrentMediaTime == nil {
		return 0.0, symbolCallError("CACurrentMediaTime", "10.5", _cACurrentMediaTimeErr)
	}
	return _cACurrentMediaTime(), nil
}

// CACurrentMediaTime returns the current absolute time, in seconds.
//
// See: https://developer.apple.com/documentation/QuartzCore/CACurrentMediaTime()
func CACurrentMediaTime() float64 {
	result, callErr := tryCACurrentMediaTime()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAFrameRateRangeIsEqualToRange func(range_ CAFrameRateRange, other CAFrameRateRange) bool
var _cAFrameRateRangeIsEqualToRangeErr error

func tryCAFrameRateRangeIsEqualToRange(range_ CAFrameRateRange, other CAFrameRateRange) (bool, error) {
	if _cAFrameRateRangeIsEqualToRange == nil {
		return false, symbolCallError("CAFrameRateRangeIsEqualToRange", "12.0", _cAFrameRateRangeIsEqualToRangeErr)
	}
	return _cAFrameRateRangeIsEqualToRange(range_, other), nil
}

// CAFrameRateRangeIsEqualToRange.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAFrameRateRangeIsEqualToRange
func CAFrameRateRangeIsEqualToRange(range_ CAFrameRateRange, other CAFrameRateRange) bool {
	result, callErr := tryCAFrameRateRangeIsEqualToRange(range_, other)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAFrameRateRangeMake func(minimum float32, maximum float32, preferred float32) CAFrameRateRange
var _cAFrameRateRangeMakeErr error

func tryCAFrameRateRangeMake(minimum float32, maximum float32, preferred float32) (CAFrameRateRange, error) {
	if _cAFrameRateRangeMake == nil {
		return CAFrameRateRange{}, symbolCallError("CAFrameRateRangeMake", "12.0", _cAFrameRateRangeMakeErr)
	}
	return _cAFrameRateRangeMake(minimum, maximum, preferred), nil
}

// CAFrameRateRangeMake.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAFrameRateRangeMake
func CAFrameRateRangeMake(minimum float32, maximum float32, preferred float32) CAFrameRateRange {
	result, callErr := tryCAFrameRateRangeMake(minimum, maximum, preferred)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cATransform3DConcat func(a CATransform3D, b CATransform3D) CATransform3D
var _cATransform3DConcatErr error

func tryCATransform3DConcat(a CATransform3D, b CATransform3D) (CATransform3D, error) {
	if _cATransform3DConcat == nil {
		return CATransform3D{}, symbolCallError("CATransform3DConcat", "10.5", _cATransform3DConcatErr)
	}
	return _cATransform3DConcat(a, b), nil
}

// CATransform3DConcat concatenates `b` to `a` and returns the result: `t = a * b`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DConcat(_:_:)
func CATransform3DConcat(a CATransform3D, b CATransform3D) CATransform3D {
	result, callErr := tryCATransform3DConcat(a, b)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cATransform3DEqualToTransform func(a CATransform3D, b CATransform3D) bool
var _cATransform3DEqualToTransformErr error

func tryCATransform3DEqualToTransform(a CATransform3D, b CATransform3D) (bool, error) {
	if _cATransform3DEqualToTransform == nil {
		return false, symbolCallError("CATransform3DEqualToTransform", "10.5", _cATransform3DEqualToTransformErr)
	}
	return _cATransform3DEqualToTransform(a, b), nil
}

// CATransform3DEqualToTransform returns a Boolean value that indicates whether the two transforms are exactly equal.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DEqualToTransform(_:_:)
func CATransform3DEqualToTransform(a CATransform3D, b CATransform3D) bool {
	result, callErr := tryCATransform3DEqualToTransform(a, b)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cATransform3DGetAffineTransform func(t CATransform3D) corefoundation.CGAffineTransform
var _cATransform3DGetAffineTransformErr error

func tryCATransform3DGetAffineTransform(t CATransform3D) (corefoundation.CGAffineTransform, error) {
	if _cATransform3DGetAffineTransform == nil {
		return corefoundation.CGAffineTransform{}, symbolCallError("CATransform3DGetAffineTransform", "10.5", _cATransform3DGetAffineTransformErr)
	}
	return _cATransform3DGetAffineTransform(t), nil
}

// CATransform3DGetAffineTransform returns the affine transform represented by `t`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DGetAffineTransform(_:)
func CATransform3DGetAffineTransform(t CATransform3D) corefoundation.CGAffineTransform {
	result, callErr := tryCATransform3DGetAffineTransform(t)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cATransform3DInvert func(t CATransform3D) CATransform3D
var _cATransform3DInvertErr error

func tryCATransform3DInvert(t CATransform3D) (CATransform3D, error) {
	if _cATransform3DInvert == nil {
		return CATransform3D{}, symbolCallError("CATransform3DInvert", "10.5", _cATransform3DInvertErr)
	}
	return _cATransform3DInvert(t), nil
}

// CATransform3DInvert inverts `t` and returns the result.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DInvert(_:)
func CATransform3DInvert(t CATransform3D) CATransform3D {
	result, callErr := tryCATransform3DInvert(t)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cATransform3DIsAffine func(t CATransform3D) bool
var _cATransform3DIsAffineErr error

func tryCATransform3DIsAffine(t CATransform3D) (bool, error) {
	if _cATransform3DIsAffine == nil {
		return false, symbolCallError("CATransform3DIsAffine", "10.5", _cATransform3DIsAffineErr)
	}
	return _cATransform3DIsAffine(t), nil
}

// CATransform3DIsAffine returns a Boolean value that indicates whether a transform can be exactly represented by an affine transform.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DIsAffine(_:)
func CATransform3DIsAffine(t CATransform3D) bool {
	result, callErr := tryCATransform3DIsAffine(t)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cATransform3DIsIdentity func(t CATransform3D) bool
var _cATransform3DIsIdentityErr error

func tryCATransform3DIsIdentity(t CATransform3D) (bool, error) {
	if _cATransform3DIsIdentity == nil {
		return false, symbolCallError("CATransform3DIsIdentity", "10.5", _cATransform3DIsIdentityErr)
	}
	return _cATransform3DIsIdentity(t), nil
}

// CATransform3DIsIdentity returns a Boolean value that indicates whether the transform is the identity transform.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DIsIdentity(_:)
func CATransform3DIsIdentity(t CATransform3D) bool {
	result, callErr := tryCATransform3DIsIdentity(t)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cATransform3DMakeAffineTransform func(m corefoundation.CGAffineTransform) CATransform3D
var _cATransform3DMakeAffineTransformErr error

func tryCATransform3DMakeAffineTransform(m corefoundation.CGAffineTransform) (CATransform3D, error) {
	if _cATransform3DMakeAffineTransform == nil {
		return CATransform3D{}, symbolCallError("CATransform3DMakeAffineTransform", "10.5", _cATransform3DMakeAffineTransformErr)
	}
	return _cATransform3DMakeAffineTransform(m), nil
}

// CATransform3DMakeAffineTransform returns a transform with the same effect as affine transform `m`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DMakeAffineTransform(_:)
func CATransform3DMakeAffineTransform(m corefoundation.CGAffineTransform) CATransform3D {
	result, callErr := tryCATransform3DMakeAffineTransform(m)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cATransform3DMakeRotation func(angle float64, x float64, y float64, z float64) CATransform3D
var _cATransform3DMakeRotationErr error

func tryCATransform3DMakeRotation(angle float64, x float64, y float64, z float64) (CATransform3D, error) {
	if _cATransform3DMakeRotation == nil {
		return CATransform3D{}, symbolCallError("CATransform3DMakeRotation", "10.5", _cATransform3DMakeRotationErr)
	}
	return _cATransform3DMakeRotation(angle, x, y, z), nil
}

// CATransform3DMakeRotation returns a transform that rotates by `angle` radians about the vector `(x, y, z)`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DMakeRotation(_:_:_:_:)
func CATransform3DMakeRotation(angle float64, x float64, y float64, z float64) CATransform3D {
	result, callErr := tryCATransform3DMakeRotation(angle, x, y, z)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cATransform3DMakeScale func(sx float64, sy float64, sz float64) CATransform3D
var _cATransform3DMakeScaleErr error

func tryCATransform3DMakeScale(sx float64, sy float64, sz float64) (CATransform3D, error) {
	if _cATransform3DMakeScale == nil {
		return CATransform3D{}, symbolCallError("CATransform3DMakeScale", "10.5", _cATransform3DMakeScaleErr)
	}
	return _cATransform3DMakeScale(sx, sy, sz), nil
}

// CATransform3DMakeScale returns a transform that scales by `(sx, sy, sz)`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DMakeScale(_:_:_:)
func CATransform3DMakeScale(sx float64, sy float64, sz float64) CATransform3D {
	result, callErr := tryCATransform3DMakeScale(sx, sy, sz)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cATransform3DMakeTranslation func(tx float64, ty float64, tz float64) CATransform3D
var _cATransform3DMakeTranslationErr error

func tryCATransform3DMakeTranslation(tx float64, ty float64, tz float64) (CATransform3D, error) {
	if _cATransform3DMakeTranslation == nil {
		return CATransform3D{}, symbolCallError("CATransform3DMakeTranslation", "10.5", _cATransform3DMakeTranslationErr)
	}
	return _cATransform3DMakeTranslation(tx, ty, tz), nil
}

// CATransform3DMakeTranslation returns a transform that translates by `(tx, ty, tz)`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DMakeTranslation(_:_:_:)
func CATransform3DMakeTranslation(tx float64, ty float64, tz float64) CATransform3D {
	result, callErr := tryCATransform3DMakeTranslation(tx, ty, tz)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cATransform3DRotate func(t CATransform3D, angle float64, x float64, y float64, z float64) CATransform3D
var _cATransform3DRotateErr error

func tryCATransform3DRotate(t CATransform3D, angle float64, x float64, y float64, z float64) (CATransform3D, error) {
	if _cATransform3DRotate == nil {
		return CATransform3D{}, symbolCallError("CATransform3DRotate", "10.5", _cATransform3DRotateErr)
	}
	return _cATransform3DRotate(t, angle, x, y, z), nil
}

// CATransform3DRotate rotates `t` by `angle` radians about the vector `(x, y, z)` and returns the result.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DRotate(_:_:_:_:_:)
func CATransform3DRotate(t CATransform3D, angle float64, x float64, y float64, z float64) CATransform3D {
	result, callErr := tryCATransform3DRotate(t, angle, x, y, z)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cATransform3DScale func(t CATransform3D, sx float64, sy float64, sz float64) CATransform3D
var _cATransform3DScaleErr error

func tryCATransform3DScale(t CATransform3D, sx float64, sy float64, sz float64) (CATransform3D, error) {
	if _cATransform3DScale == nil {
		return CATransform3D{}, symbolCallError("CATransform3DScale", "10.5", _cATransform3DScaleErr)
	}
	return _cATransform3DScale(t, sx, sy, sz), nil
}

// CATransform3DScale scales `t` by `(sx, sy, sz)` and returns the result: `t = scale(sx, sy, sz) * t`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DScale(_:_:_:_:)
func CATransform3DScale(t CATransform3D, sx float64, sy float64, sz float64) CATransform3D {
	result, callErr := tryCATransform3DScale(t, sx, sy, sz)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cATransform3DTranslate func(t CATransform3D, tx float64, ty float64, tz float64) CATransform3D
var _cATransform3DTranslateErr error

func tryCATransform3DTranslate(t CATransform3D, tx float64, ty float64, tz float64) (CATransform3D, error) {
	if _cATransform3DTranslate == nil {
		return CATransform3D{}, symbolCallError("CATransform3DTranslate", "10.5", _cATransform3DTranslateErr)
	}
	return _cATransform3DTranslate(t, tx, ty, tz), nil
}

// CATransform3DTranslate translates `t` by `(tx, ty, tz)` and returns the result: `t` `= translate(tx, ty, tz) * t`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DTranslate(_:_:_:_:)
func CATransform3DTranslate(t CATransform3D, tx float64, ty float64, tz float64) CATransform3D {
	result, callErr := tryCATransform3DTranslate(t, tx, ty, tz)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_cACurrentMediaTime, &_cACurrentMediaTimeErr, frameworkHandle, "CACurrentMediaTime", "10.5")
	registerFunc(&_cAFrameRateRangeIsEqualToRange, &_cAFrameRateRangeIsEqualToRangeErr, frameworkHandle, "CAFrameRateRangeIsEqualToRange", "12.0")
	registerFunc(&_cAFrameRateRangeMake, &_cAFrameRateRangeMakeErr, frameworkHandle, "CAFrameRateRangeMake", "12.0")
	registerFunc(&_cATransform3DConcat, &_cATransform3DConcatErr, frameworkHandle, "CATransform3DConcat", "10.5")
	registerFunc(&_cATransform3DEqualToTransform, &_cATransform3DEqualToTransformErr, frameworkHandle, "CATransform3DEqualToTransform", "10.5")
	registerFunc(&_cATransform3DGetAffineTransform, &_cATransform3DGetAffineTransformErr, frameworkHandle, "CATransform3DGetAffineTransform", "10.5")
	registerFunc(&_cATransform3DInvert, &_cATransform3DInvertErr, frameworkHandle, "CATransform3DInvert", "10.5")
	registerFunc(&_cATransform3DIsAffine, &_cATransform3DIsAffineErr, frameworkHandle, "CATransform3DIsAffine", "10.5")
	registerFunc(&_cATransform3DIsIdentity, &_cATransform3DIsIdentityErr, frameworkHandle, "CATransform3DIsIdentity", "10.5")
	registerFunc(&_cATransform3DMakeAffineTransform, &_cATransform3DMakeAffineTransformErr, frameworkHandle, "CATransform3DMakeAffineTransform", "10.5")
	registerFunc(&_cATransform3DMakeRotation, &_cATransform3DMakeRotationErr, frameworkHandle, "CATransform3DMakeRotation", "10.5")
	registerFunc(&_cATransform3DMakeScale, &_cATransform3DMakeScaleErr, frameworkHandle, "CATransform3DMakeScale", "10.5")
	registerFunc(&_cATransform3DMakeTranslation, &_cATransform3DMakeTranslationErr, frameworkHandle, "CATransform3DMakeTranslation", "10.5")
	registerFunc(&_cATransform3DRotate, &_cATransform3DRotateErr, frameworkHandle, "CATransform3DRotate", "10.5")
	registerFunc(&_cATransform3DScale, &_cATransform3DScaleErr, frameworkHandle, "CATransform3DScale", "10.5")
	registerFunc(&_cATransform3DTranslate, &_cATransform3DTranslateErr, frameworkHandle, "CATransform3DTranslate", "10.5")
}

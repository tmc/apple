// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: QuartzCore: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: QuartzCore: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: QuartzCore: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _cACurrentMediaTime func() float64

// CACurrentMediaTime returns the current absolute time, in seconds.
//
// See: https://developer.apple.com/documentation/QuartzCore/CACurrentMediaTime()
func CACurrentMediaTime() float64 {
	if _cACurrentMediaTime == nil {
		panic("QuartzCore: symbol CACurrentMediaTime not loaded")
	}
	return _cACurrentMediaTime()
}

var _cAFrameRateRangeIsEqualToRange func(range_ CAFrameRateRange, other CAFrameRateRange) bool

// CAFrameRateRangeIsEqualToRange.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAFrameRateRangeIsEqualToRange
func CAFrameRateRangeIsEqualToRange(range_ CAFrameRateRange, other CAFrameRateRange) bool {
	if _cAFrameRateRangeIsEqualToRange == nil {
		panic("QuartzCore: symbol CAFrameRateRangeIsEqualToRange not loaded")
	}
	return _cAFrameRateRangeIsEqualToRange(range_, other)
}

var _cAFrameRateRangeMake func(minimum float32, maximum float32, preferred float32) CAFrameRateRange

// CAFrameRateRangeMake.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAFrameRateRangeMake
func CAFrameRateRangeMake(minimum float32, maximum float32, preferred float32) CAFrameRateRange {
	if _cAFrameRateRangeMake == nil {
		panic("QuartzCore: symbol CAFrameRateRangeMake not loaded")
	}
	return _cAFrameRateRangeMake(minimum, maximum, preferred)
}

var _cATransform3DConcat func(a CATransform3D, b CATransform3D) CATransform3D

// CATransform3DConcat concatenates `b` to `a` and returns the result: `t = a * b`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DConcat(_:_:)
func CATransform3DConcat(a CATransform3D, b CATransform3D) CATransform3D {
	if _cATransform3DConcat == nil {
		panic("QuartzCore: symbol CATransform3DConcat not loaded")
	}
	return _cATransform3DConcat(a, b)
}

var _cATransform3DEqualToTransform func(a CATransform3D, b CATransform3D) bool

// CATransform3DEqualToTransform returns a Boolean value that indicates whether the two transforms are exactly equal.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DEqualToTransform(_:_:)
func CATransform3DEqualToTransform(a CATransform3D, b CATransform3D) bool {
	if _cATransform3DEqualToTransform == nil {
		panic("QuartzCore: symbol CATransform3DEqualToTransform not loaded")
	}
	return _cATransform3DEqualToTransform(a, b)
}

var _cATransform3DGetAffineTransform func(t CATransform3D) corefoundation.CGAffineTransform

// CATransform3DGetAffineTransform returns the affine transform represented by `t`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DGetAffineTransform(_:)
func CATransform3DGetAffineTransform(t CATransform3D) corefoundation.CGAffineTransform {
	if _cATransform3DGetAffineTransform == nil {
		panic("QuartzCore: symbol CATransform3DGetAffineTransform not loaded")
	}
	return _cATransform3DGetAffineTransform(t)
}

var _cATransform3DInvert func(t CATransform3D) CATransform3D

// CATransform3DInvert inverts `t` and returns the result.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DInvert(_:)
func CATransform3DInvert(t CATransform3D) CATransform3D {
	if _cATransform3DInvert == nil {
		panic("QuartzCore: symbol CATransform3DInvert not loaded")
	}
	return _cATransform3DInvert(t)
}

var _cATransform3DIsAffine func(t CATransform3D) bool

// CATransform3DIsAffine returns a Boolean value that indicates whether a transform can be exactly represented by an affine transform.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DIsAffine(_:)
func CATransform3DIsAffine(t CATransform3D) bool {
	if _cATransform3DIsAffine == nil {
		panic("QuartzCore: symbol CATransform3DIsAffine not loaded")
	}
	return _cATransform3DIsAffine(t)
}

var _cATransform3DIsIdentity func(t CATransform3D) bool

// CATransform3DIsIdentity returns a Boolean value that indicates whether the transform is the identity transform.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DIsIdentity(_:)
func CATransform3DIsIdentity(t CATransform3D) bool {
	if _cATransform3DIsIdentity == nil {
		panic("QuartzCore: symbol CATransform3DIsIdentity not loaded")
	}
	return _cATransform3DIsIdentity(t)
}

var _cATransform3DMakeAffineTransform func(m corefoundation.CGAffineTransform) CATransform3D

// CATransform3DMakeAffineTransform returns a transform with the same effect as affine transform `m`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DMakeAffineTransform(_:)
func CATransform3DMakeAffineTransform(m corefoundation.CGAffineTransform) CATransform3D {
	if _cATransform3DMakeAffineTransform == nil {
		panic("QuartzCore: symbol CATransform3DMakeAffineTransform not loaded")
	}
	return _cATransform3DMakeAffineTransform(m)
}

var _cATransform3DMakeRotation func(angle float64, x float64, y float64, z float64) CATransform3D

// CATransform3DMakeRotation returns a transform that rotates by `angle` radians about the vector `(x, y, z)`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DMakeRotation(_:_:_:_:)
func CATransform3DMakeRotation(angle float64, x float64, y float64, z float64) CATransform3D {
	if _cATransform3DMakeRotation == nil {
		panic("QuartzCore: symbol CATransform3DMakeRotation not loaded")
	}
	return _cATransform3DMakeRotation(angle, x, y, z)
}

var _cATransform3DMakeScale func(sx float64, sy float64, sz float64) CATransform3D

// CATransform3DMakeScale returns a transform that scales by `(sx, sy, sz)`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DMakeScale(_:_:_:)
func CATransform3DMakeScale(sx float64, sy float64, sz float64) CATransform3D {
	if _cATransform3DMakeScale == nil {
		panic("QuartzCore: symbol CATransform3DMakeScale not loaded")
	}
	return _cATransform3DMakeScale(sx, sy, sz)
}

var _cATransform3DMakeTranslation func(tx float64, ty float64, tz float64) CATransform3D

// CATransform3DMakeTranslation returns a transform that translates by `(tx, ty, tz)`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DMakeTranslation(_:_:_:)
func CATransform3DMakeTranslation(tx float64, ty float64, tz float64) CATransform3D {
	if _cATransform3DMakeTranslation == nil {
		panic("QuartzCore: symbol CATransform3DMakeTranslation not loaded")
	}
	return _cATransform3DMakeTranslation(tx, ty, tz)
}

var _cATransform3DRotate func(t CATransform3D, angle float64, x float64, y float64, z float64) CATransform3D

// CATransform3DRotate rotates `t` by `angle` radians about the vector `(x, y, z)` and returns the result.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DRotate(_:_:_:_:_:)
func CATransform3DRotate(t CATransform3D, angle float64, x float64, y float64, z float64) CATransform3D {
	if _cATransform3DRotate == nil {
		panic("QuartzCore: symbol CATransform3DRotate not loaded")
	}
	return _cATransform3DRotate(t, angle, x, y, z)
}

var _cATransform3DScale func(t CATransform3D, sx float64, sy float64, sz float64) CATransform3D

// CATransform3DScale scales `t` by `(sx, sy, sz)` and returns the result: `t = scale(sx, sy, sz) * t`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DScale(_:_:_:_:)
func CATransform3DScale(t CATransform3D, sx float64, sy float64, sz float64) CATransform3D {
	if _cATransform3DScale == nil {
		panic("QuartzCore: symbol CATransform3DScale not loaded")
	}
	return _cATransform3DScale(t, sx, sy, sz)
}

var _cATransform3DTranslate func(t CATransform3D, tx float64, ty float64, tz float64) CATransform3D

// CATransform3DTranslate translates `t` by `(tx, ty, tz)` and returns the result: `t` `= translate(tx, ty, tz) * t`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DTranslate(_:_:_:_:)
func CATransform3DTranslate(t CATransform3D, tx float64, ty float64, tz float64) CATransform3D {
	if _cATransform3DTranslate == nil {
		panic("QuartzCore: symbol CATransform3DTranslate not loaded")
	}
	return _cATransform3DTranslate(t, tx, ty, tz)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
		registerFunc(&_cACurrentMediaTime, frameworkHandle, "CACurrentMediaTime")
		registerFunc(&_cAFrameRateRangeIsEqualToRange, frameworkHandle, "CAFrameRateRangeIsEqualToRange")
		registerFunc(&_cAFrameRateRangeMake, frameworkHandle, "CAFrameRateRangeMake")
		registerFunc(&_cATransform3DConcat, frameworkHandle, "CATransform3DConcat")
		registerFunc(&_cATransform3DEqualToTransform, frameworkHandle, "CATransform3DEqualToTransform")
		registerFunc(&_cATransform3DGetAffineTransform, frameworkHandle, "CATransform3DGetAffineTransform")
		registerFunc(&_cATransform3DInvert, frameworkHandle, "CATransform3DInvert")
		registerFunc(&_cATransform3DIsAffine, frameworkHandle, "CATransform3DIsAffine")
		registerFunc(&_cATransform3DIsIdentity, frameworkHandle, "CATransform3DIsIdentity")
		registerFunc(&_cATransform3DMakeAffineTransform, frameworkHandle, "CATransform3DMakeAffineTransform")
		registerFunc(&_cATransform3DMakeRotation, frameworkHandle, "CATransform3DMakeRotation")
		registerFunc(&_cATransform3DMakeScale, frameworkHandle, "CATransform3DMakeScale")
		registerFunc(&_cATransform3DMakeTranslation, frameworkHandle, "CATransform3DMakeTranslation")
		registerFunc(&_cATransform3DRotate, frameworkHandle, "CATransform3DRotate")
		registerFunc(&_cATransform3DScale, frameworkHandle, "CATransform3DScale")
		registerFunc(&_cATransform3DTranslate, frameworkHandle, "CATransform3DTranslate")
	}


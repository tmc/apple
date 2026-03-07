// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLImageSizeConstraint] class.
var (
	_MLImageSizeConstraintClass     MLImageSizeConstraintClass
	_MLImageSizeConstraintClassOnce sync.Once
)

func getMLImageSizeConstraintClass() MLImageSizeConstraintClass {
	_MLImageSizeConstraintClassOnce.Do(func() {
		_MLImageSizeConstraintClass = MLImageSizeConstraintClass{class: objc.GetClass("MLImageSizeConstraint")}
	})
	return _MLImageSizeConstraintClass
}

// GetMLImageSizeConstraintClass returns the class object for MLImageSizeConstraint.
func GetMLImageSizeConstraintClass() MLImageSizeConstraintClass {
	return getMLImageSizeConstraintClass()
}

type MLImageSizeConstraintClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLImageSizeConstraintClass) Alloc() MLImageSizeConstraint {
	rv := objc.Send[MLImageSizeConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A list or range of sizes that augment an image constraint’s default size.
//
// # Overview
// 
// You use an [MLImageSizeConstraint] to express what image sizes of an image
// feature a model will accept as input or produce as output.
// 
// Use [MLImageSizeConstraint.Type] to determine which properties describe what image sizes the
// model’s image feature expects as input or produces as output.
// 
// If `type` is:
// 
// - [MLImageSizeConstraintType.range], the image feature accepts any image
// that has a width in [MLImageSizeConstraint.PixelsWideRange] and a height in [MLImageSizeConstraint.PixelsHighRange]. -
// [MLImageSizeConstraintType.enumerated], the image feature accepts any image
// size listed in [MLImageSizeConstraint.EnumeratedImageSizes]. -
// [MLImageSizeConstraintType.unspecified], the [MLImageSizeConstraint]
// instance is not configured and should be ignored. Instead, use the image
// feature’s default image size constraint, defined by [MLImageSizeConstraint.PixelsWide] and
// [MLImageSizeConstraint.PixelsHigh].
// 
// [media-3027121]
//
// [MLImageSizeConstraintType.enumerated]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraintType/enumerated
// [MLImageSizeConstraintType.range]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraintType/range
// [MLImageSizeConstraintType.unspecified]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraintType/unspecified
//
// # Determining relevant constraints
//
//   - [MLImageSizeConstraint.Type]: Indicator of which properties to inspect for this image size constraint.
//
// # Accessing the image size ranges
//
//   - [MLImageSizeConstraint.PixelsWideRange]: The range of widths a model’s image feature accepts as input or produces as output.
//   - [MLImageSizeConstraint.PixelsHighRange]: The range of heights a model’s image feature accepts as input or produces as output.
//
// # Accessing the enumerated image sizes
//
//   - [MLImageSizeConstraint.EnumeratedImageSizes]: An array of image sizes a model’s image feature accepts as input or produces as output.
//
// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint
type MLImageSizeConstraint struct {
	objectivec.Object
}

// MLImageSizeConstraintFromID constructs a [MLImageSizeConstraint] from an objc.ID.
//
// A list or range of sizes that augment an image constraint’s default size.
func MLImageSizeConstraintFromID(id objc.ID) MLImageSizeConstraint {
	return MLImageSizeConstraint{objectivec.Object{id}}
}
// NOTE: MLImageSizeConstraint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MLImageSizeConstraint] class.
//
// # Determining relevant constraints
//
//   - [IMLImageSizeConstraint.Type]: Indicator of which properties to inspect for this image size constraint.
//
// # Accessing the image size ranges
//
//   - [IMLImageSizeConstraint.PixelsWideRange]: The range of widths a model’s image feature accepts as input or produces as output.
//   - [IMLImageSizeConstraint.PixelsHighRange]: The range of heights a model’s image feature accepts as input or produces as output.
//
// # Accessing the enumerated image sizes
//
//   - [IMLImageSizeConstraint.EnumeratedImageSizes]: An array of image sizes a model’s image feature accepts as input or produces as output.
//
// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint
type IMLImageSizeConstraint interface {
	objectivec.IObject

	// Topic: Determining relevant constraints

	// Indicator of which properties to inspect for this image size constraint.
	Type() MLImageSizeConstraintType

	// Topic: Accessing the image size ranges

	// The range of widths a model’s image feature accepts as input or produces as output.
	PixelsWideRange() foundation.NSRange
	// The range of heights a model’s image feature accepts as input or produces as output.
	PixelsHighRange() foundation.NSRange

	// Topic: Accessing the enumerated image sizes

	// An array of image sizes a model’s image feature accepts as input or produces as output.
	EnumeratedImageSizes() []MLImageSize

	// The model’s default height for an image feature.
	PixelsHigh() int
	SetPixelsHigh(value int)
	// The model’s default width for an image feature.
	PixelsWide() int
	SetPixelsWide(value int)
	// Additional sizes this image feature supports.
	SizeConstraint() IMLImageSizeConstraint
	SetSizeConstraint(value IMLImageSizeConstraint)
	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (i MLImageSizeConstraint) Init() MLImageSizeConstraint {
	rv := objc.Send[MLImageSizeConstraint](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MLImageSizeConstraint) Autorelease() MLImageSizeConstraint {
	rv := objc.Send[MLImageSizeConstraint](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLImageSizeConstraint creates a new MLImageSizeConstraint instance.
func NewMLImageSizeConstraint() MLImageSizeConstraint {
	class := getMLImageSizeConstraintClass()
	rv := objc.Send[MLImageSizeConstraint](objc.ID(class.class), objc.Sel("new"))
	return rv
}









func (i MLImageSizeConstraint) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeWithCoder:"), coder)
}












// Indicator of which properties to inspect for this image size constraint.
//
// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/type
func (i MLImageSizeConstraint) Type() MLImageSizeConstraintType {
	rv := objc.Send[MLImageSizeConstraintType](i.ID, objc.Sel("type"))
	return MLImageSizeConstraintType(rv)
}



// The range of widths a model’s image feature accepts as input or produces
// as output.
//
// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/pixelsWideRange
func (i MLImageSizeConstraint) PixelsWideRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](i.ID, objc.Sel("pixelsWideRange"))
	return foundation.NSRange(rv)
}



// The range of heights a model’s image feature accepts as input or produces
// as output.
//
// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/pixelsHighRange
func (i MLImageSizeConstraint) PixelsHighRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](i.ID, objc.Sel("pixelsHighRange"))
	return foundation.NSRange(rv)
}



// An array of image sizes a model’s image feature accepts as input or
// produces as output.
//
// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/enumeratedImageSizes
func (i MLImageSizeConstraint) EnumeratedImageSizes() []MLImageSize {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("enumeratedImageSizes"))
	return objc.ConvertSlice(rv, func(id objc.ID) MLImageSize {
		return MLImageSizeFromID(id)
	})
}



// The model’s default height for an image feature.
//
// See: https://developer.apple.com/documentation/coreml/mlimageconstraint/pixelshigh
func (i MLImageSizeConstraint) PixelsHigh() int {
	rv := objc.Send[int](i.ID, objc.Sel("pixelsHigh"))
	return rv
}
func (i MLImageSizeConstraint) SetPixelsHigh(value int) {
	objc.Send[struct{}](i.ID, objc.Sel("setPixelsHigh:"), value)
}



// The model’s default width for an image feature.
//
// See: https://developer.apple.com/documentation/coreml/mlimageconstraint/pixelswide
func (i MLImageSizeConstraint) PixelsWide() int {
	rv := objc.Send[int](i.ID, objc.Sel("pixelsWide"))
	return rv
}
func (i MLImageSizeConstraint) SetPixelsWide(value int) {
	objc.Send[struct{}](i.ID, objc.Sel("setPixelsWide:"), value)
}



// Additional sizes this image feature supports.
//
// See: https://developer.apple.com/documentation/coreml/mlimageconstraint/sizeconstraint
func (i MLImageSizeConstraint) SizeConstraint() IMLImageSizeConstraint {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("sizeConstraint"))
	return MLImageSizeConstraintFromID(objc.ID(rv))
}
func (i MLImageSizeConstraint) SetSizeConstraint(value IMLImageSizeConstraint) {
	objc.Send[struct{}](i.ID, objc.Sel("setSizeConstraint:"), value)
}


























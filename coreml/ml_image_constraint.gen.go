// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLImageConstraint] class.
var (
	_MLImageConstraintClass     MLImageConstraintClass
	_MLImageConstraintClassOnce sync.Once
)

func getMLImageConstraintClass() MLImageConstraintClass {
	_MLImageConstraintClassOnce.Do(func() {
		_MLImageConstraintClass = MLImageConstraintClass{class: objc.GetClass("MLImageConstraint")}
	})
	return _MLImageConstraintClass
}

// GetMLImageConstraintClass returns the class object for MLImageConstraint.
func GetMLImageConstraintClass() MLImageConstraintClass {
	return getMLImageConstraintClass()
}

type MLImageConstraintClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLImageConstraintClass) Alloc() MLImageConstraint {
	rv := objc.Send[MLImageConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// The width, height, and pixel format constraints of an image feature.
//
// # Overview
// 
// In CoreML, an is a collection of pixels represented by [CVPixelBuffer]
// (Swift) or [CVPixelBuffer] (Objective-C). An is a model input or output
// that accepts or produces, respectively, an image bundled in an
// [MLFeatureValue]. [MLImageConstraint] defines the image feature’s
// limitations for the images within an [MLFeatureValue].
// 
// If a model has an image feature for an input or output, the model author
// uses an by creating an [MLFeatureDescription]. The feature description for
// an image input or output has:
// 
// - Its [MLImageConstraint.Type] property set to [MLFeatureType.image] - Its [ImageConstraint]
// property set to an [MLImageConstraint] instance configured to the image
// feature’s size and format
// 
// Image features that support additional image sizes provide a range of
// sizes, or a list of discrete sizes, in their image constraint’s
// [MLImageConstraint.SizeConstraint] property.
//
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/CVPixelBuffer
// [MLFeatureType.image]: https://developer.apple.com/documentation/CoreML/MLFeatureType/image
//
// # Accessing the constraints
//
//   - [MLImageConstraint.PixelsWide]: The model’s default width for an image feature.
//   - [MLImageConstraint.PixelsHigh]: The model’s default height for an image feature.
//   - [MLImageConstraint.PixelFormatType]: The model’s pixel format for an image feature.
//
// # Inspecting acceptable sizes
//
//   - [MLImageConstraint.SizeConstraint]: Additional sizes this image feature supports.
//
// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint
type MLImageConstraint struct {
	objectivec.Object
}

// MLImageConstraintFromID constructs a [MLImageConstraint] from an objc.ID.
//
// The width, height, and pixel format constraints of an image feature.
func MLImageConstraintFromID(id objc.ID) MLImageConstraint {
	return MLImageConstraint{objectivec.Object{id}}
}
// NOTE: MLImageConstraint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MLImageConstraint] class.
//
// # Accessing the constraints
//
//   - [IMLImageConstraint.PixelsWide]: The model’s default width for an image feature.
//   - [IMLImageConstraint.PixelsHigh]: The model’s default height for an image feature.
//   - [IMLImageConstraint.PixelFormatType]: The model’s pixel format for an image feature.
//
// # Inspecting acceptable sizes
//
//   - [IMLImageConstraint.SizeConstraint]: Additional sizes this image feature supports.
//
// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint
type IMLImageConstraint interface {
	objectivec.IObject

	// Topic: Accessing the constraints

	// The model’s default width for an image feature.
	PixelsWide() int
	// The model’s default height for an image feature.
	PixelsHigh() int
	// The model’s pixel format for an image feature.
	PixelFormatType() uint32

	// Topic: Inspecting acceptable sizes

	// Additional sizes this image feature supports.
	SizeConstraint() IMLImageSizeConstraint

	// The size and format constraints for an image feature.
	ImageConstraint() IMLImageConstraint
	SetImageConstraint(value IMLImageConstraint)
	// The type of this feature.
	Type() MLFeatureType
	SetType(value MLFeatureType)
	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (i MLImageConstraint) Init() MLImageConstraint {
	rv := objc.Send[MLImageConstraint](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MLImageConstraint) Autorelease() MLImageConstraint {
	rv := objc.Send[MLImageConstraint](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLImageConstraint creates a new MLImageConstraint instance.
func NewMLImageConstraint() MLImageConstraint {
	class := getMLImageConstraintClass()
	rv := objc.Send[MLImageConstraint](objc.ID(class.class), objc.Sel("new"))
	return rv
}









func (i MLImageConstraint) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The model’s default width for an image feature.
//
// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/pixelsWide
func (i MLImageConstraint) PixelsWide() int {
	rv := objc.Send[int](i.ID, objc.Sel("pixelsWide"))
	return rv
}



// The model’s default height for an image feature.
//
// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/pixelsHigh
func (i MLImageConstraint) PixelsHigh() int {
	rv := objc.Send[int](i.ID, objc.Sel("pixelsHigh"))
	return rv
}



// The model’s pixel format for an image feature.
//
// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/pixelFormatType
func (i MLImageConstraint) PixelFormatType() uint32 {
	rv := objc.Send[uint32](i.ID, objc.Sel("pixelFormatType"))
	return rv
}



// Additional sizes this image feature supports.
//
// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/sizeConstraint
func (i MLImageConstraint) SizeConstraint() IMLImageSizeConstraint {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("sizeConstraint"))
	return MLImageSizeConstraintFromID(objc.ID(rv))
}



// The size and format constraints for an image feature.
//
// See: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (i MLImageConstraint) ImageConstraint() IMLImageConstraint {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageConstraint"))
	return MLImageConstraintFromID(objc.ID(rv))
}
func (i MLImageConstraint) SetImageConstraint(value IMLImageConstraint) {
	objc.Send[struct{}](i.ID, objc.Sel("setImageConstraint:"), value)
}



// The type of this feature.
//
// See: https://developer.apple.com/documentation/coreml/mlfeaturedescription/type
func (i MLImageConstraint) Type() MLFeatureType {
	rv := objc.Send[MLFeatureType](i.ID, objc.Sel("type"))
	return MLFeatureType(rv)
}
func (i MLImageConstraint) SetType(value MLFeatureType) {
	objc.Send[struct{}](i.ID, objc.Sel("setType:"), value)
}


























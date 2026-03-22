// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CIMaximumScaleTransform protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMaximumScaleTransform
type CIMaximumScaleTransform interface {
	objectivec.IObject
	CIFilterProtocol

	// AspectRatio protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMaximumScaleTransform/aspectRatio
	AspectRatio() float32

	// InputImage protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMaximumScaleTransform/inputImage
	InputImage() ICIImage

	// Scale protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMaximumScaleTransform/scale
	Scale() float32

	// aspectRatio protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMaximumScaleTransform/aspectRatio
	SetAspectRatio(value float32)

	// inputImage protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMaximumScaleTransform/inputImage
	SetInputImage(value ICIImage)

	// scale protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMaximumScaleTransform/scale
	SetScale(value float32)
}

// CIMaximumScaleTransformObject wraps an existing Objective-C object that conforms to the CIMaximumScaleTransform protocol.
type CIMaximumScaleTransformObject struct {
	objectivec.Object
}
func (o CIMaximumScaleTransformObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIMaximumScaleTransformObjectFromID constructs a [CIMaximumScaleTransformObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIMaximumScaleTransformObjectFromID(id objc.ID) CIMaximumScaleTransformObject {
	return CIMaximumScaleTransformObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIMaximumScaleTransform/aspectRatio
func (o CIMaximumScaleTransformObject) AspectRatio() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("aspectRatio"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIMaximumScaleTransform/inputImage
func (o CIMaximumScaleTransformObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CIMaximumScaleTransform/scale
func (o CIMaximumScaleTransformObject) Scale() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("scale"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIMaximumScaleTransformObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIMaximumScaleTransformObject) SetAspectRatio(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAspectRatio:"), value)
}

func (o CIMaximumScaleTransformObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIMaximumScaleTransformObject) SetScale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setScale:"), value)
}


// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a zoom blur filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIZoomBlur
type CIZoomBlur interface {
	objectivec.IObject
	CIFilterProtocol

	// The zoom-in amount.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIZoomBlur/amount
	Amount() float32

	// The center of the effect, as x and y coordinates.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIZoomBlur/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIZoomBlur/inputImage
	InputImage() ICIImage

	// The zoom-in amount.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIZoomBlur/amount
	SetAmount(value float32)

	// The center of the effect, as x and y coordinates.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIZoomBlur/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIZoomBlur/inputImage
	SetInputImage(value ICIImage)
}

// CIZoomBlurObject wraps an existing Objective-C object that conforms to the CIZoomBlur protocol.
type CIZoomBlurObject struct {
	objectivec.Object
}
func (o CIZoomBlurObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIZoomBlurObjectFromID constructs a [CIZoomBlurObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIZoomBlurObjectFromID(id objc.ID) CIZoomBlurObject {
	return CIZoomBlurObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The zoom-in amount.
//
// See: https://developer.apple.com/documentation/CoreImage/CIZoomBlur/amount
func (o CIZoomBlurObject) Amount() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("amount"))
	return rv
	}
// The center of the effect, as x and y coordinates.
//
// See: https://developer.apple.com/documentation/CoreImage/CIZoomBlur/center
func (o CIZoomBlurObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIZoomBlur/inputImage
func (o CIZoomBlurObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIZoomBlurObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIZoomBlurObject) SetAmount(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAmount:"), value)
}

func (o CIZoomBlurObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIZoomBlurObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}


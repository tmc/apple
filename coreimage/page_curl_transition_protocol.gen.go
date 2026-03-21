// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a page curl transition filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlTransition
type CIPageCurlTransition interface {
	objectivec.IObject
	CIFilterProtocol
	CITransitionFilter

	// The angle of the curling page.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlTransition/angle
	Angle() float32

	// The image that appears on the back of the source image as the page curls to reveal the target image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlTransition/backsideImage
	BacksideImage() ICIImage

	// The extent of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlTransition/extent
	Extent() corefoundation.CGRect

	// The radius of the curl.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlTransition/radius
	Radius() float32

	// An image that looks like a shaded sphere enclosed in a square.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlTransition/shadingImage
	ShadingImage() ICIImage

	// The angle of the curling page.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlTransition/angle
	SetAngle(value float32)

	// The image that appears on the back of the source image as the page curls to reveal the target image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlTransition/backsideImage
	SetBacksideImage(value ICIImage)

	// The extent of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlTransition/extent
	SetExtent(value corefoundation.CGRect)

	// The radius of the curl.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlTransition/radius
	SetRadius(value float32)

	// An image that looks like a shaded sphere enclosed in a square.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlTransition/shadingImage
	SetShadingImage(value ICIImage)
}

// CIPageCurlTransitionObject wraps an existing Objective-C object that conforms to the CIPageCurlTransition protocol.
type CIPageCurlTransitionObject struct {
	objectivec.Object
}
func (o CIPageCurlTransitionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIPageCurlTransitionObjectFromID constructs a [CIPageCurlTransitionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIPageCurlTransitionObjectFromID(id objc.ID) CIPageCurlTransitionObject {
	return CIPageCurlTransitionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The angle of the curling page.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlTransition/angle
func (o CIPageCurlTransitionObject) Angle() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// The image that appears on the back of the source image as the page curls to
// reveal the target image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlTransition/backsideImage
func (o CIPageCurlTransitionObject) BacksideImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("backsideImage"))
	return CIImageFromID(rv)
	}
// The extent of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlTransition/extent
func (o CIPageCurlTransitionObject) Extent() corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
	}
// The radius of the curl.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlTransition/radius
func (o CIPageCurlTransitionObject) Radius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// An image that looks like a shaded sphere enclosed in a square.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlTransition/shadingImage
func (o CIPageCurlTransitionObject) ShadingImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("shadingImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIPageCurlTransitionObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/inputImage
func (o CIPageCurlTransitionObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The target image for a transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/targetImage
func (o CIPageCurlTransitionObject) TargetImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("targetImage"))
	return CIImageFromID(rv)
	}
// The parametric time of the transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/time
func (o CIPageCurlTransitionObject) Time() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("time"))
	return rv
	}

func (o CIPageCurlTransitionObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CIPageCurlTransitionObject) SetBacksideImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setBacksideImage:"), value)
}

func (o CIPageCurlTransitionObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

func (o CIPageCurlTransitionObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

func (o CIPageCurlTransitionObject) SetShadingImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setShadingImage:"), value)
}

func (o CIPageCurlTransitionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIPageCurlTransitionObject) SetTargetImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setTargetImage:"), value)
}

func (o CIPageCurlTransitionObject) SetTime(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setTime:"), value)
}


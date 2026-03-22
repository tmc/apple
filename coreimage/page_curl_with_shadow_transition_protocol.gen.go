// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a page-curl-with-shadow transition filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition
type CIPageCurlWithShadowTransition interface {
	objectivec.IObject
	CIFilterProtocol
	CITransitionFilter

	// The angle of the curling page.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/angle
	Angle() float32

	// The image that appears on the back of the source image as the page curls to reveal the target image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/backsideImage
	BacksideImage() ICIImage

	// The extent of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/extent
	Extent() corefoundation.CGRect

	// The radius of the curl.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/radius
	Radius() float32

	// The strength of the shadow.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/shadowAmount
	ShadowAmount() float32

	// The rectagular portion of input image that casts a shadow.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/shadowExtent
	ShadowExtent() corefoundation.CGRect

	// The maximum size, in pixels, of the shadow.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/shadowSize
	ShadowSize() float32

	// The angle of the curling page.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/angle
	SetAngle(value float32)

	// The image that appears on the back of the source image as the page curls to reveal the target image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/backsideImage
	SetBacksideImage(value ICIImage)

	// The extent of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/extent
	SetExtent(value corefoundation.CGRect)

	// The radius of the curl.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/radius
	SetRadius(value float32)

	// The strength of the shadow.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/shadowAmount
	SetShadowAmount(value float32)

	// The rectagular portion of input image that casts a shadow.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/shadowExtent
	SetShadowExtent(value corefoundation.CGRect)

	// The maximum size, in pixels, of the shadow.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/shadowSize
	SetShadowSize(value float32)
}

// CIPageCurlWithShadowTransitionObject wraps an existing Objective-C object that conforms to the CIPageCurlWithShadowTransition protocol.
type CIPageCurlWithShadowTransitionObject struct {
	objectivec.Object
}
func (o CIPageCurlWithShadowTransitionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIPageCurlWithShadowTransitionObjectFromID constructs a [CIPageCurlWithShadowTransitionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIPageCurlWithShadowTransitionObjectFromID(id objc.ID) CIPageCurlWithShadowTransitionObject {
	return CIPageCurlWithShadowTransitionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The angle of the curling page.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/angle
func (o CIPageCurlWithShadowTransitionObject) Angle() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// The image that appears on the back of the source image as the page curls to
// reveal the target image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/backsideImage
func (o CIPageCurlWithShadowTransitionObject) BacksideImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("backsideImage"))
	return CIImageFromID(rv)
	}
// The extent of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/extent
func (o CIPageCurlWithShadowTransitionObject) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
	}
// The radius of the curl.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/radius
func (o CIPageCurlWithShadowTransitionObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// The strength of the shadow.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/shadowAmount
func (o CIPageCurlWithShadowTransitionObject) ShadowAmount() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("shadowAmount"))
	return rv
	}
// The rectagular portion of input image that casts a shadow.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/shadowExtent
func (o CIPageCurlWithShadowTransitionObject) ShadowExtent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("shadowExtent"))
	return rv
	}
// The maximum size, in pixels, of the shadow.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPageCurlWithShadowTransition/shadowSize
func (o CIPageCurlWithShadowTransitionObject) ShadowSize() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("shadowSize"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIPageCurlWithShadowTransitionObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/inputImage
func (o CIPageCurlWithShadowTransitionObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The target image for a transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/targetImage
func (o CIPageCurlWithShadowTransitionObject) TargetImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("targetImage"))
	return CIImageFromID(rv)
	}
// The parametric time of the transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/time
func (o CIPageCurlWithShadowTransitionObject) Time() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("time"))
	return rv
	}

func (o CIPageCurlWithShadowTransitionObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CIPageCurlWithShadowTransitionObject) SetBacksideImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setBacksideImage:"), value)
}

func (o CIPageCurlWithShadowTransitionObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

func (o CIPageCurlWithShadowTransitionObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

func (o CIPageCurlWithShadowTransitionObject) SetShadowAmount(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setShadowAmount:"), value)
}

func (o CIPageCurlWithShadowTransitionObject) SetShadowExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setShadowExtent:"), value)
}

func (o CIPageCurlWithShadowTransitionObject) SetShadowSize(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setShadowSize:"), value)
}

func (o CIPageCurlWithShadowTransitionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIPageCurlWithShadowTransitionObject) SetTargetImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setTargetImage:"), value)
}

func (o CIPageCurlWithShadowTransitionObject) SetTime(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setTime:"), value)
}


// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a ripple transition filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRippleTransition
type CIRippleTransition interface {
	objectivec.IObject
	CIFilterProtocol
	CITransitionFilter

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRippleTransition/center
	Center() corefoundation.CGPoint

	// A rectangle that defines the extent of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRippleTransition/extent
	Extent() corefoundation.CGRect

	// A value that determines whether the ripple starts as a bulge (a higher value) or a dimple (a lower value).
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRippleTransition/scale
	Scale() float32

	// An image that looks like a shaded sphere enclosed in a square.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRippleTransition/shadingImage
	ShadingImage() ICIImage

	// The width of the ripple.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRippleTransition/width
	Width() float32

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRippleTransition/center
	SetCenter(value corefoundation.CGPoint)

	// A rectangle that defines the extent of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRippleTransition/extent
	SetExtent(value corefoundation.CGRect)

	// A value that determines whether the ripple starts as a bulge (a higher value) or a dimple (a lower value).
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRippleTransition/scale
	SetScale(value float32)

	// An image that looks like a shaded sphere enclosed in a square.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRippleTransition/shadingImage
	SetShadingImage(value ICIImage)

	// The width of the ripple.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRippleTransition/width
	SetWidth(value float32)
}

// CIRippleTransitionObject wraps an existing Objective-C object that conforms to the CIRippleTransition protocol.
type CIRippleTransitionObject struct {
	objectivec.Object
}
func (o CIRippleTransitionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIRippleTransitionObjectFromID constructs a [CIRippleTransitionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIRippleTransitionObjectFromID(id objc.ID) CIRippleTransitionObject {
	return CIRippleTransitionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The x and y position to use as the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRippleTransition/center
func (o CIRippleTransitionObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// A rectangle that defines the extent of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRippleTransition/extent
func (o CIRippleTransitionObject) Extent() corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
	}
// A value that determines whether the ripple starts as a bulge (a higher
// value) or a dimple (a lower value).
//
// See: https://developer.apple.com/documentation/CoreImage/CIRippleTransition/scale
func (o CIRippleTransitionObject) Scale() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("scale"))
	return rv
	}
// An image that looks like a shaded sphere enclosed in a square.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRippleTransition/shadingImage
func (o CIRippleTransitionObject) ShadingImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("shadingImage"))
	return CIImageFromID(rv)
	}
// The width of the ripple.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRippleTransition/width
func (o CIRippleTransitionObject) Width() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIRippleTransitionObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/inputImage
func (o CIRippleTransitionObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The target image for a transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/targetImage
func (o CIRippleTransitionObject) TargetImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("targetImage"))
	return CIImageFromID(rv)
	}
// The parametric time of the transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/time
func (o CIRippleTransitionObject) Time() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("time"))
	return rv
	}

func (o CIRippleTransitionObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIRippleTransitionObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

func (o CIRippleTransitionObject) SetScale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setScale:"), value)
}

func (o CIRippleTransitionObject) SetShadingImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setShadingImage:"), value)
}

func (o CIRippleTransitionObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}

func (o CIRippleTransitionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIRippleTransitionObject) SetTargetImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setTargetImage:"), value)
}

func (o CIRippleTransitionObject) SetTime(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setTime:"), value)
}


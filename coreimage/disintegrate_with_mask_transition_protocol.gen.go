// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a disintegrate-with-mask transition filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDisintegrateWithMaskTransition
type CIDisintegrateWithMaskTransition interface {
	objectivec.IObject
	CIFilterProtocol
	CITransitionFilter

	// An image that defines the shape to use when disintegrating from the source to the target image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDisintegrateWithMaskTransition/maskImage
	MaskImage() ICIImage

	// The density of the shadow the mask creates.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDisintegrateWithMaskTransition/shadowDensity
	ShadowDensity() float32

	// The offset of the shadow the mask creates.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDisintegrateWithMaskTransition/shadowOffset
	ShadowOffset() corefoundation.CGPoint

	// The radius of the shadow the mask creates.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDisintegrateWithMaskTransition/shadowRadius
	ShadowRadius() float32

	// An image that defines the shape to use when disintegrating from the source to the target image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDisintegrateWithMaskTransition/maskImage
	SetMaskImage(value ICIImage)

	// The density of the shadow the mask creates.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDisintegrateWithMaskTransition/shadowDensity
	SetShadowDensity(value float32)

	// The offset of the shadow the mask creates.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDisintegrateWithMaskTransition/shadowOffset
	SetShadowOffset(value corefoundation.CGPoint)

	// The radius of the shadow the mask creates.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDisintegrateWithMaskTransition/shadowRadius
	SetShadowRadius(value float32)
}

// CIDisintegrateWithMaskTransitionObject wraps an existing Objective-C object that conforms to the CIDisintegrateWithMaskTransition protocol.
type CIDisintegrateWithMaskTransitionObject struct {
	objectivec.Object
}
func (o CIDisintegrateWithMaskTransitionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIDisintegrateWithMaskTransitionObjectFromID constructs a [CIDisintegrateWithMaskTransitionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIDisintegrateWithMaskTransitionObjectFromID(id objc.ID) CIDisintegrateWithMaskTransitionObject {
	return CIDisintegrateWithMaskTransitionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// An image that defines the shape to use when disintegrating from the source
// to the target image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDisintegrateWithMaskTransition/maskImage
func (o CIDisintegrateWithMaskTransitionObject) MaskImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("maskImage"))
	return CIImageFromID(rv)
	}
// The density of the shadow the mask creates.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDisintegrateWithMaskTransition/shadowDensity
func (o CIDisintegrateWithMaskTransitionObject) ShadowDensity() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("shadowDensity"))
	return rv
	}
// The offset of the shadow the mask creates.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDisintegrateWithMaskTransition/shadowOffset
func (o CIDisintegrateWithMaskTransitionObject) ShadowOffset() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("shadowOffset"))
	return rv
	}
// The radius of the shadow the mask creates.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDisintegrateWithMaskTransition/shadowRadius
func (o CIDisintegrateWithMaskTransitionObject) ShadowRadius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("shadowRadius"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIDisintegrateWithMaskTransitionObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/inputImage
func (o CIDisintegrateWithMaskTransitionObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The target image for a transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/targetImage
func (o CIDisintegrateWithMaskTransitionObject) TargetImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("targetImage"))
	return CIImageFromID(rv)
	}
// The parametric time of the transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/time
func (o CIDisintegrateWithMaskTransitionObject) Time() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("time"))
	return rv
	}

func (o CIDisintegrateWithMaskTransitionObject) SetMaskImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setMaskImage:"), value)
}

func (o CIDisintegrateWithMaskTransitionObject) SetShadowDensity(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setShadowDensity:"), value)
}

func (o CIDisintegrateWithMaskTransitionObject) SetShadowOffset(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setShadowOffset:"), value)
}

func (o CIDisintegrateWithMaskTransitionObject) SetShadowRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setShadowRadius:"), value)
}

func (o CIDisintegrateWithMaskTransitionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIDisintegrateWithMaskTransitionObject) SetTargetImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setTargetImage:"), value)
}

func (o CIDisintegrateWithMaskTransitionObject) SetTime(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setTime:"), value)
}


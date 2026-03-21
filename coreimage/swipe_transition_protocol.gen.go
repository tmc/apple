// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a swipe transition filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CISwipeTransition
type CISwipeTransition interface {
	objectivec.IObject
	CIFilterProtocol
	CITransitionFilter

	// The angle of the swipe.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISwipeTransition/angle
	Angle() float32

	// The color of the swipe.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISwipeTransition/color
	Color() ICIColor

	// The extent of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISwipeTransition/extent
	Extent() corefoundation.CGRect

	// The opacity of the swipe.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISwipeTransition/opacity
	Opacity() float32

	// The width of the swipe.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISwipeTransition/width
	Width() float32

	// The angle of the swipe.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISwipeTransition/angle
	SetAngle(value float32)

	// The color of the swipe.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISwipeTransition/color
	SetColor(value ICIColor)

	// The extent of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISwipeTransition/extent
	SetExtent(value corefoundation.CGRect)

	// The opacity of the swipe.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISwipeTransition/opacity
	SetOpacity(value float32)

	// The width of the swipe.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISwipeTransition/width
	SetWidth(value float32)
}

// CISwipeTransitionObject wraps an existing Objective-C object that conforms to the CISwipeTransition protocol.
type CISwipeTransitionObject struct {
	objectivec.Object
}
func (o CISwipeTransitionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CISwipeTransitionObjectFromID constructs a [CISwipeTransitionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CISwipeTransitionObjectFromID(id objc.ID) CISwipeTransitionObject {
	return CISwipeTransitionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The angle of the swipe.
//
// See: https://developer.apple.com/documentation/CoreImage/CISwipeTransition/angle
func (o CISwipeTransitionObject) Angle() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// The color of the swipe.
//
// See: https://developer.apple.com/documentation/CoreImage/CISwipeTransition/color
func (o CISwipeTransitionObject) Color() ICIColor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color"))
	return CIColorFromID(rv)
	}
// The extent of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CISwipeTransition/extent
func (o CISwipeTransitionObject) Extent() corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
	}
// The opacity of the swipe.
//
// See: https://developer.apple.com/documentation/CoreImage/CISwipeTransition/opacity
func (o CISwipeTransitionObject) Opacity() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("opacity"))
	return rv
	}
// The width of the swipe.
//
// See: https://developer.apple.com/documentation/CoreImage/CISwipeTransition/width
func (o CISwipeTransitionObject) Width() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CISwipeTransitionObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/inputImage
func (o CISwipeTransitionObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The target image for a transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/targetImage
func (o CISwipeTransitionObject) TargetImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("targetImage"))
	return CIImageFromID(rv)
	}
// The parametric time of the transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/time
func (o CISwipeTransitionObject) Time() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("time"))
	return rv
	}

func (o CISwipeTransitionObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CISwipeTransitionObject) SetColor(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor:"), value)
}

func (o CISwipeTransitionObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

func (o CISwipeTransitionObject) SetOpacity(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setOpacity:"), value)
}

func (o CISwipeTransitionObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}

func (o CISwipeTransitionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CISwipeTransitionObject) SetTargetImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setTargetImage:"), value)
}

func (o CISwipeTransitionObject) SetTime(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setTime:"), value)
}


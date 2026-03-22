// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a bars swipe transition filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBarsSwipeTransition
type CIBarsSwipeTransition interface {
	objectivec.IObject
	CIFilterProtocol
	CITransitionFilter

	// The angle, in radians, of the bars.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBarsSwipeTransition/angle
	Angle() float32

	// The offset of one bar with respect to another.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBarsSwipeTransition/barOffset
	BarOffset() float32

	// The width of each bar.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBarsSwipeTransition/width
	Width() float32

	// The angle, in radians, of the bars.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBarsSwipeTransition/angle
	SetAngle(value float32)

	// The offset of one bar with respect to another.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBarsSwipeTransition/barOffset
	SetBarOffset(value float32)

	// The width of each bar.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBarsSwipeTransition/width
	SetWidth(value float32)
}

// CIBarsSwipeTransitionObject wraps an existing Objective-C object that conforms to the CIBarsSwipeTransition protocol.
type CIBarsSwipeTransitionObject struct {
	objectivec.Object
}
func (o CIBarsSwipeTransitionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIBarsSwipeTransitionObjectFromID constructs a [CIBarsSwipeTransitionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIBarsSwipeTransitionObjectFromID(id objc.ID) CIBarsSwipeTransitionObject {
	return CIBarsSwipeTransitionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The angle, in radians, of the bars.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBarsSwipeTransition/angle
func (o CIBarsSwipeTransitionObject) Angle() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// The offset of one bar with respect to another.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBarsSwipeTransition/barOffset
func (o CIBarsSwipeTransitionObject) BarOffset() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("barOffset"))
	return rv
	}
// The width of each bar.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBarsSwipeTransition/width
func (o CIBarsSwipeTransitionObject) Width() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIBarsSwipeTransitionObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/inputImage
func (o CIBarsSwipeTransitionObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The target image for a transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/targetImage
func (o CIBarsSwipeTransitionObject) TargetImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("targetImage"))
	return CIImageFromID(rv)
	}
// The parametric time of the transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/time
func (o CIBarsSwipeTransitionObject) Time() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("time"))
	return rv
	}

func (o CIBarsSwipeTransitionObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CIBarsSwipeTransitionObject) SetBarOffset(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setBarOffset:"), value)
}

func (o CIBarsSwipeTransitionObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}

func (o CIBarsSwipeTransitionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIBarsSwipeTransitionObject) SetTargetImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setTargetImage:"), value)
}

func (o CIBarsSwipeTransitionObject) SetTime(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setTime:"), value)
}


// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a dissolve transition filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDissolveTransition
type CIDissolveTransition interface {
	objectivec.IObject
	CIFilterProtocol
	CITransitionFilter
}

// CIDissolveTransitionObject wraps an existing Objective-C object that conforms to the CIDissolveTransition protocol.
type CIDissolveTransitionObject struct {
	objectivec.Object
}
func (o CIDissolveTransitionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIDissolveTransitionObjectFromID constructs a [CIDissolveTransitionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIDissolveTransitionObjectFromID(id objc.ID) CIDissolveTransitionObject {
	return CIDissolveTransitionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIDissolveTransitionObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/inputImage
func (o CIDissolveTransitionObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The target image for a transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/targetImage
func (o CIDissolveTransitionObject) TargetImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("targetImage"))
	return CIImageFromID(rv)
	}
// The parametric time of the transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/time
func (o CIDissolveTransitionObject) Time() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("time"))
	return rv
	}

func (o CIDissolveTransitionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIDissolveTransitionObject) SetTargetImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setTargetImage:"), value)
}

func (o CIDissolveTransitionObject) SetTime(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setTime:"), value)
}


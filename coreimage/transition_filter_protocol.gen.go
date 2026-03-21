// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a transition filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter
type CITransitionFilter interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/inputImage
	InputImage() ICIImage

	// The target image for a transition.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/targetImage
	TargetImage() ICIImage

	// The parametric time of the transition.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/time
	Time() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/inputImage
	SetInputImage(value ICIImage)

	// The target image for a transition.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/targetImage
	SetTargetImage(value ICIImage)

	// The parametric time of the transition.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/time
	SetTime(value float32)
}

// CITransitionFilterObject wraps an existing Objective-C object that conforms to the CITransitionFilter protocol.
type CITransitionFilterObject struct {
	objectivec.Object
}
func (o CITransitionFilterObject) BaseObject() objectivec.Object {
	return o.Object
}

// CITransitionFilterObjectFromID constructs a [CITransitionFilterObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CITransitionFilterObjectFromID(id objc.ID) CITransitionFilterObject {
	return CITransitionFilterObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/inputImage
func (o CITransitionFilterObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The target image for a transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/targetImage
func (o CITransitionFilterObject) TargetImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("targetImage"))
	return CIImageFromID(rv)
	}
// The parametric time of the transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/time
func (o CITransitionFilterObject) Time() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("time"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CITransitionFilterObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CITransitionFilterObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CITransitionFilterObject) SetTargetImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setTargetImage:"), value)
}

func (o CITransitionFilterObject) SetTime(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setTime:"), value)
}


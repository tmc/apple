// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a crystalize filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CICrystallize
type CICrystallize interface {
	objectivec.IObject
	CIFilterProtocol

	// The center of the effect as x and y coordinates.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICrystallize/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICrystallize/inputImage
	InputImage() ICIImage

	// The radius, in pixels, of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICrystallize/radius
	Radius() float32

	// The center of the effect as x and y coordinates.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICrystallize/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICrystallize/inputImage
	SetInputImage(value ICIImage)

	// The radius, in pixels, of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICrystallize/radius
	SetRadius(value float32)
}

// CICrystallizeObject wraps an existing Objective-C object that conforms to the CICrystallize protocol.
type CICrystallizeObject struct {
	objectivec.Object
}
func (o CICrystallizeObject) BaseObject() objectivec.Object {
	return o.Object
}

// CICrystallizeObjectFromID constructs a [CICrystallizeObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CICrystallizeObjectFromID(id objc.ID) CICrystallizeObject {
	return CICrystallizeObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The center of the effect as x and y coordinates.
//
// See: https://developer.apple.com/documentation/CoreImage/CICrystallize/center
func (o CICrystallizeObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CICrystallize/inputImage
func (o CICrystallizeObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The radius, in pixels, of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CICrystallize/radius
func (o CICrystallizeObject) Radius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CICrystallizeObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CICrystallizeObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CICrystallizeObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CICrystallizeObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}


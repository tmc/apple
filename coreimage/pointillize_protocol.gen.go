// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a pointillize filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPointillize
type CIPointillize interface {
	objectivec.IObject
	CIFilterProtocol

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPointillize/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPointillize/inputImage
	InputImage() ICIImage

	// The radius of the circles in the resulting pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPointillize/radius
	Radius() float32

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPointillize/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPointillize/inputImage
	SetInputImage(value ICIImage)

	// The radius of the circles in the resulting pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPointillize/radius
	SetRadius(value float32)
}

// CIPointillizeObject wraps an existing Objective-C object that conforms to the CIPointillize protocol.
type CIPointillizeObject struct {
	objectivec.Object
}
func (o CIPointillizeObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIPointillizeObjectFromID constructs a [CIPointillizeObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIPointillizeObjectFromID(id objc.ID) CIPointillizeObject {
	return CIPointillizeObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The x and y position to use as the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPointillize/center
func (o CIPointillizeObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPointillize/inputImage
func (o CIPointillizeObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The radius of the circles in the resulting pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPointillize/radius
func (o CIPointillizeObject) Radius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIPointillizeObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIPointillizeObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIPointillizeObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIPointillizeObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}


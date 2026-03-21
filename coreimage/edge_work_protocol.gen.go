// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure an edge-work filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIEdgeWork
type CIEdgeWork interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEdgeWork/inputImage
	InputImage() ICIImage

	// The thickness of the edges.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEdgeWork/radius
	Radius() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEdgeWork/inputImage
	SetInputImage(value ICIImage)

	// The thickness of the edges.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEdgeWork/radius
	SetRadius(value float32)
}

// CIEdgeWorkObject wraps an existing Objective-C object that conforms to the CIEdgeWork protocol.
type CIEdgeWorkObject struct {
	objectivec.Object
}
func (o CIEdgeWorkObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIEdgeWorkObjectFromID constructs a [CIEdgeWorkObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIEdgeWorkObjectFromID(id objc.ID) CIEdgeWorkObject {
	return CIEdgeWorkObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIEdgeWork/inputImage
func (o CIEdgeWorkObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The thickness of the edges.
//
// See: https://developer.apple.com/documentation/CoreImage/CIEdgeWork/radius
func (o CIEdgeWorkObject) Radius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIEdgeWorkObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIEdgeWorkObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIEdgeWorkObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}


// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure an edges filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIEdges
type CIEdges interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEdges/inputImage
	InputImage() ICIImage

	// The intensity of the edges.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEdges/intensity
	Intensity() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEdges/inputImage
	SetInputImage(value ICIImage)

	// The intensity of the edges.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEdges/intensity
	SetIntensity(value float32)
}

// CIEdgesObject wraps an existing Objective-C object that conforms to the CIEdges protocol.
type CIEdgesObject struct {
	objectivec.Object
}
func (o CIEdgesObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIEdgesObjectFromID constructs a [CIEdgesObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIEdgesObjectFromID(id objc.ID) CIEdgesObject {
	return CIEdgesObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIEdges/inputImage
func (o CIEdgesObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The intensity of the edges.
//
// See: https://developer.apple.com/documentation/CoreImage/CIEdges/intensity
func (o CIEdgesObject) Intensity() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("intensity"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIEdgesObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIEdgesObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIEdgesObject) SetIntensity(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setIntensity:"), value)
}


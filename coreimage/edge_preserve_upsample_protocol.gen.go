// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure an edge preserve upsample filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIEdgePreserveUpsample
type CIEdgePreserveUpsample interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEdgePreserveUpsample/inputImage
	InputImage() ICIImage

	// A value that specifies the influence of the input image’s luma information on the upsampling operation.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEdgePreserveUpsample/lumaSigma
	LumaSigma() float32

	// The image that the filter upsamples.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEdgePreserveUpsample/smallImage
	SmallImage() ICIImage

	// A value that specifies the influence of the input image’s spatial information on the upsampling operation.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEdgePreserveUpsample/spatialSigma
	SpatialSigma() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEdgePreserveUpsample/inputImage
	SetInputImage(value ICIImage)

	// A value that specifies the influence of the input image’s luma information on the upsampling operation.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEdgePreserveUpsample/lumaSigma
	SetLumaSigma(value float32)

	// The image that the filter upsamples.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEdgePreserveUpsample/smallImage
	SetSmallImage(value ICIImage)

	// A value that specifies the influence of the input image’s spatial information on the upsampling operation.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEdgePreserveUpsample/spatialSigma
	SetSpatialSigma(value float32)
}

// CIEdgePreserveUpsampleObject wraps an existing Objective-C object that conforms to the CIEdgePreserveUpsample protocol.
type CIEdgePreserveUpsampleObject struct {
	objectivec.Object
}
func (o CIEdgePreserveUpsampleObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIEdgePreserveUpsampleObjectFromID constructs a [CIEdgePreserveUpsampleObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIEdgePreserveUpsampleObjectFromID(id objc.ID) CIEdgePreserveUpsampleObject {
	return CIEdgePreserveUpsampleObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIEdgePreserveUpsample/inputImage
func (o CIEdgePreserveUpsampleObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A value that specifies the influence of the input image’s luma
// information on the upsampling operation.
//
// See: https://developer.apple.com/documentation/CoreImage/CIEdgePreserveUpsample/lumaSigma
func (o CIEdgePreserveUpsampleObject) LumaSigma() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("lumaSigma"))
	return rv
	}
// The image that the filter upsamples.
//
// See: https://developer.apple.com/documentation/CoreImage/CIEdgePreserveUpsample/smallImage
func (o CIEdgePreserveUpsampleObject) SmallImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("smallImage"))
	return CIImageFromID(rv)
	}
// A value that specifies the influence of the input image’s spatial
// information on the upsampling operation.
//
// See: https://developer.apple.com/documentation/CoreImage/CIEdgePreserveUpsample/spatialSigma
func (o CIEdgePreserveUpsampleObject) SpatialSigma() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("spatialSigma"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIEdgePreserveUpsampleObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIEdgePreserveUpsampleObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIEdgePreserveUpsampleObject) SetLumaSigma(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setLumaSigma:"), value)
}

func (o CIEdgePreserveUpsampleObject) SetSmallImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setSmallImage:"), value)
}

func (o CIEdgePreserveUpsampleObject) SetSpatialSigma(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSpatialSigma:"), value)
}


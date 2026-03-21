// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The protocol for the Signed Distance Gradient From Red Mask filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CISignedDistanceGradientFromRedMask
type CISignedDistanceGradientFromRedMask interface {
	objectivec.IObject
	CIFilterProtocol

	// The input image whose red channel defines a mask. If the red channel pixel value is greater than 0.5 then the point is considered in the mask and output pixel will be a value between zero and negative one. Otherwise the output pixel will be a value between zero and one.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISignedDistanceGradientFromRedMask/inputImage
	InputImage() ICIImage

	// Determines the maximum distance to the mask that can be measured. Distances between zero and the maximum will be normalized to negative one and one.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISignedDistanceGradientFromRedMask/maximumDistance
	MaximumDistance() float32

	// The input image whose red channel defines a mask. If the red channel pixel value is greater than 0.5 then the point is considered in the mask and output pixel will be a value between zero and negative one. Otherwise the output pixel will be a value between zero and one.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISignedDistanceGradientFromRedMask/inputImage
	SetInputImage(value ICIImage)

	// Determines the maximum distance to the mask that can be measured. Distances between zero and the maximum will be normalized to negative one and one.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISignedDistanceGradientFromRedMask/maximumDistance
	SetMaximumDistance(value float32)
}

// CISignedDistanceGradientFromRedMaskObject wraps an existing Objective-C object that conforms to the CISignedDistanceGradientFromRedMask protocol.
type CISignedDistanceGradientFromRedMaskObject struct {
	objectivec.Object
}
func (o CISignedDistanceGradientFromRedMaskObject) BaseObject() objectivec.Object {
	return o.Object
}

// CISignedDistanceGradientFromRedMaskObjectFromID constructs a [CISignedDistanceGradientFromRedMaskObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CISignedDistanceGradientFromRedMaskObjectFromID(id objc.ID) CISignedDistanceGradientFromRedMaskObject {
	return CISignedDistanceGradientFromRedMaskObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The input image whose red channel defines a mask. If the red channel pixel
// value is greater than 0.5 then the point is considered in the mask and
// output pixel will be a value between zero and negative one. Otherwise the
// output pixel will be a value between zero and one.
//
// See: https://developer.apple.com/documentation/CoreImage/CISignedDistanceGradientFromRedMask/inputImage
func (o CISignedDistanceGradientFromRedMaskObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// Determines the maximum distance to the mask that can be measured. Distances
// between zero and the maximum will be normalized to negative one and one.
//
// See: https://developer.apple.com/documentation/CoreImage/CISignedDistanceGradientFromRedMask/maximumDistance
func (o CISignedDistanceGradientFromRedMaskObject) MaximumDistance() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("maximumDistance"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CISignedDistanceGradientFromRedMaskObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CISignedDistanceGradientFromRedMaskObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CISignedDistanceGradientFromRedMaskObject) SetMaximumDistance(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setMaximumDistance:"), value)
}


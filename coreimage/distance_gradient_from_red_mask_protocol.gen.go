// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The protocol for the Distance Gradient From Red Mask filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDistanceGradientFromRedMask
type CIDistanceGradientFromRedMask interface {
	objectivec.IObject
	CIFilterProtocol

	// The input image whose red channel defines a mask. If the red channel pixel value is greater than 0.5 then the point is considered in the mask and output pixel will be zero. Otherwise the output pixel will be a value between zero and one.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDistanceGradientFromRedMask/inputImage
	InputImage() ICIImage

	// Determines the maximum distance to the mask that can be measured. Distances between zero and the maximum will be normalized to zero and one.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDistanceGradientFromRedMask/maximumDistance
	MaximumDistance() float32

	// The input image whose red channel defines a mask. If the red channel pixel value is greater than 0.5 then the point is considered in the mask and output pixel will be zero. Otherwise the output pixel will be a value between zero and one.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDistanceGradientFromRedMask/inputImage
	SetInputImage(value ICIImage)

	// Determines the maximum distance to the mask that can be measured. Distances between zero and the maximum will be normalized to zero and one.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDistanceGradientFromRedMask/maximumDistance
	SetMaximumDistance(value float32)
}

// CIDistanceGradientFromRedMaskObject wraps an existing Objective-C object that conforms to the CIDistanceGradientFromRedMask protocol.
type CIDistanceGradientFromRedMaskObject struct {
	objectivec.Object
}

func (o CIDistanceGradientFromRedMaskObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIDistanceGradientFromRedMaskObjectFromID constructs a [CIDistanceGradientFromRedMaskObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIDistanceGradientFromRedMaskObjectFromID(id objc.ID) CIDistanceGradientFromRedMaskObject {
	return CIDistanceGradientFromRedMaskObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The input image whose red channel defines a mask. If the red channel pixel
// value is greater than 0.5 then the point is considered in the mask and
// output pixel will be zero. Otherwise the output pixel will be a value
// between zero and one.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDistanceGradientFromRedMask/inputImage
func (o CIDistanceGradientFromRedMaskObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// Determines the maximum distance to the mask that can be measured. Distances
// between zero and the maximum will be normalized to zero and one.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDistanceGradientFromRedMask/maximumDistance
func (o CIDistanceGradientFromRedMaskObject) MaximumDistance() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("maximumDistance"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIDistanceGradientFromRedMaskObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The input image whose red channel defines a mask. If the red channel pixel
// value is greater than 0.5 then the point is considered in the mask and
// output pixel will be zero. Otherwise the output pixel will be a value
// between zero and one.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDistanceGradientFromRedMask/inputImage
func (o CIDistanceGradientFromRedMaskObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// Determines the maximum distance to the mask that can be measured. Distances
// between zero and the maximum will be normalized to zero and one.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDistanceGradientFromRedMask/maximumDistance
func (o CIDistanceGradientFromRedMaskObject) SetMaximumDistance(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setMaximumDistance:"), value)
}

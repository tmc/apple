// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a color matrix filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorMatrix
type CIColorMatrix interface {
	objectivec.IObject
	CIFilterProtocol

	// The amount of alpha to multiply the source color values by.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorMatrix/aVector
	AVector() ICIVector
	SetAVector(value ICIVector)

	// The amount of blue to multiply the source color values by.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorMatrix/bVector
	BVector() ICIVector
	SetBVector(value ICIVector)

	// The amount of green to multiply the source color values by.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorMatrix/gVector
	GVector() ICIVector
	SetGVector(value ICIVector)

	// The amount of red to multiply the source color values by.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorMatrix/rVector
	RVector() ICIVector
	SetRVector(value ICIVector)

	// A vector that’s added to each color component.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorMatrix/biasVector
	BiasVector() ICIVector
	SetBiasVector(value ICIVector)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorMatrix/inputImage
	InputImage() ICIImage
	SetInputImage(value ICIImage)
}

// CIColorMatrixObject wraps an existing Objective-C object that conforms to the CIColorMatrix protocol.
type CIColorMatrixObject struct {
	objectivec.Object
}

func (o CIColorMatrixObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIColorMatrixObjectFromID constructs a [CIColorMatrixObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIColorMatrixObjectFromID(id objc.ID) CIColorMatrixObject {
	return CIColorMatrixObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIColorMatrixObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The amount of alpha to multiply the source color values by.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorMatrix/aVector
func (o CIColorMatrixObject) AVector() ICIVector {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("AVector"))
	return CIVectorFromID(rv)
}

func (o CIColorMatrixObject) SetAVector(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setAVector:"), value)
}

// The amount of blue to multiply the source color values by.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorMatrix/bVector
func (o CIColorMatrixObject) BVector() ICIVector {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("BVector"))
	return CIVectorFromID(rv)
}

func (o CIColorMatrixObject) SetBVector(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setBVector:"), value)
}

// The amount of green to multiply the source color values by.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorMatrix/gVector
func (o CIColorMatrixObject) GVector() ICIVector {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("GVector"))
	return CIVectorFromID(rv)
}

func (o CIColorMatrixObject) SetGVector(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setGVector:"), value)
}

// The amount of red to multiply the source color values by.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorMatrix/rVector
func (o CIColorMatrixObject) RVector() ICIVector {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("RVector"))
	return CIVectorFromID(rv)
}

func (o CIColorMatrixObject) SetRVector(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setRVector:"), value)
}

// A vector that’s added to each color component.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorMatrix/biasVector
func (o CIColorMatrixObject) BiasVector() ICIVector {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("biasVector"))
	return CIVectorFromID(rv)
}

func (o CIColorMatrixObject) SetBiasVector(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setBiasVector:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorMatrix/inputImage
func (o CIColorMatrixObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

func (o CIColorMatrixObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

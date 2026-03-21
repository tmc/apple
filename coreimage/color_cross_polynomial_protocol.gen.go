// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a color cross-polynomial filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCrossPolynomial
type CIColorCrossPolynomial interface {
	objectivec.IObject
	CIFilterProtocol

	// Polynomial coefficients for the blue channel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCrossPolynomial/blueCoefficients
	BlueCoefficients() ICIVector

	// Polynomial coefficients for the green channel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCrossPolynomial/greenCoefficients
	GreenCoefficients() ICIVector

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCrossPolynomial/inputImage
	InputImage() ICIImage

	// Polynomial coefficients for the red channel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCrossPolynomial/redCoefficients
	RedCoefficients() ICIVector

	// Polynomial coefficients for the blue channel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCrossPolynomial/blueCoefficients
	SetBlueCoefficients(value ICIVector)

	// Polynomial coefficients for the green channel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCrossPolynomial/greenCoefficients
	SetGreenCoefficients(value ICIVector)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCrossPolynomial/inputImage
	SetInputImage(value ICIImage)

	// Polynomial coefficients for the red channel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCrossPolynomial/redCoefficients
	SetRedCoefficients(value ICIVector)
}

// CIColorCrossPolynomialObject wraps an existing Objective-C object that conforms to the CIColorCrossPolynomial protocol.
type CIColorCrossPolynomialObject struct {
	objectivec.Object
}
func (o CIColorCrossPolynomialObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIColorCrossPolynomialObjectFromID constructs a [CIColorCrossPolynomialObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIColorCrossPolynomialObjectFromID(id objc.ID) CIColorCrossPolynomialObject {
	return CIColorCrossPolynomialObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Polynomial coefficients for the blue channel.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCrossPolynomial/blueCoefficients
func (o CIColorCrossPolynomialObject) BlueCoefficients() ICIVector {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("blueCoefficients"))
	return CIVectorFromID(rv)
	}
// Polynomial coefficients for the green channel.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCrossPolynomial/greenCoefficients
func (o CIColorCrossPolynomialObject) GreenCoefficients() ICIVector {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("greenCoefficients"))
	return CIVectorFromID(rv)
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCrossPolynomial/inputImage
func (o CIColorCrossPolynomialObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// Polynomial coefficients for the red channel.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCrossPolynomial/redCoefficients
func (o CIColorCrossPolynomialObject) RedCoefficients() ICIVector {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("redCoefficients"))
	return CIVectorFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIColorCrossPolynomialObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIColorCrossPolynomialObject) SetBlueCoefficients(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setBlueCoefficients:"), value)
}

func (o CIColorCrossPolynomialObject) SetGreenCoefficients(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setGreenCoefficients:"), value)
}

func (o CIColorCrossPolynomialObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIColorCrossPolynomialObject) SetRedCoefficients(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setRedCoefficients:"), value)
}


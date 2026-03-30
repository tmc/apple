// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a color polynomial filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial
type CIColorPolynomial interface {
	objectivec.IObject
	CIFilterProtocol

	// Polynomial coefficients for the alpha channel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/alphaCoefficients
	AlphaCoefficients() ICIVector

	// Polynomial coefficients for the blue channel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/blueCoefficients
	BlueCoefficients() ICIVector

	// Polynomial coefficients for the green channel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/greenCoefficients
	GreenCoefficients() ICIVector

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/inputImage
	InputImage() ICIImage

	// Polynomial coefficients for the red channel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/redCoefficients
	RedCoefficients() ICIVector

	// Polynomial coefficients for the alpha channel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/alphaCoefficients
	SetAlphaCoefficients(value ICIVector)

	// Polynomial coefficients for the blue channel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/blueCoefficients
	SetBlueCoefficients(value ICIVector)

	// Polynomial coefficients for the green channel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/greenCoefficients
	SetGreenCoefficients(value ICIVector)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/inputImage
	SetInputImage(value ICIImage)

	// Polynomial coefficients for the red channel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/redCoefficients
	SetRedCoefficients(value ICIVector)
}

// CIColorPolynomialObject wraps an existing Objective-C object that conforms to the CIColorPolynomial protocol.
type CIColorPolynomialObject struct {
	objectivec.Object
}

func (o CIColorPolynomialObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIColorPolynomialObjectFromID constructs a [CIColorPolynomialObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIColorPolynomialObjectFromID(id objc.ID) CIColorPolynomialObject {
	return CIColorPolynomialObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Polynomial coefficients for the alpha channel.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/alphaCoefficients
func (o CIColorPolynomialObject) AlphaCoefficients() ICIVector {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("alphaCoefficients"))
	return CIVectorFromID(rv)
}

// Polynomial coefficients for the blue channel.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/blueCoefficients
func (o CIColorPolynomialObject) BlueCoefficients() ICIVector {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("blueCoefficients"))
	return CIVectorFromID(rv)
}

// Polynomial coefficients for the green channel.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/greenCoefficients
func (o CIColorPolynomialObject) GreenCoefficients() ICIVector {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("greenCoefficients"))
	return CIVectorFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/inputImage
func (o CIColorPolynomialObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// Polynomial coefficients for the red channel.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/redCoefficients
func (o CIColorPolynomialObject) RedCoefficients() ICIVector {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("redCoefficients"))
	return CIVectorFromID(rv)
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIColorPolynomialObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// Polynomial coefficients for the alpha channel.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/alphaCoefficients
func (o CIColorPolynomialObject) SetAlphaCoefficients(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setAlphaCoefficients:"), value)
}

// Polynomial coefficients for the blue channel.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/blueCoefficients
func (o CIColorPolynomialObject) SetBlueCoefficients(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setBlueCoefficients:"), value)
}

// Polynomial coefficients for the green channel.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/greenCoefficients
func (o CIColorPolynomialObject) SetGreenCoefficients(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setGreenCoefficients:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/inputImage
func (o CIColorPolynomialObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// Polynomial coefficients for the red channel.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorPolynomial/redCoefficients
func (o CIColorPolynomialObject) SetRedCoefficients(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setRedCoefficients:"), value)
}

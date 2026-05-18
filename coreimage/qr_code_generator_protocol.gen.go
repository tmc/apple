// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a QR code generator filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeGenerator
type CIQRCodeGenerator interface {
	objectivec.IObject
	CIFilterProtocol

	// The QR code correction level: L, M, Q, or H.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeGenerator/correctionLevel
	CorrectionLevel() string
	SetCorrectionLevel(value string)

	// The message to encode in the QR code.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeGenerator/message
	Message() foundation.NSData
	SetMessage(value foundation.NSData)
}

// CIQRCodeGeneratorObject wraps an existing Objective-C object that conforms to the CIQRCodeGenerator protocol.
type CIQRCodeGeneratorObject struct {
	objectivec.Object
}

func (o CIQRCodeGeneratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIQRCodeGeneratorObjectFromID constructs a [CIQRCodeGeneratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIQRCodeGeneratorObjectFromID(id objc.ID) CIQRCodeGeneratorObject {
	return CIQRCodeGeneratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIQRCodeGeneratorObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The QR code correction level: L, M, Q, or H.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeGenerator/correctionLevel
func (o CIQRCodeGeneratorObject) CorrectionLevel() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("correctionLevel"))
	return foundation.NSStringFromID(rv).String()
}

func (o CIQRCodeGeneratorObject) SetCorrectionLevel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setCorrectionLevel:"), objc.String(value))
}

// The message to encode in the QR code.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeGenerator/message
func (o CIQRCodeGeneratorObject) Message() foundation.NSData {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("message"))
	return foundation.NSDataFromID(rv)
}

func (o CIQRCodeGeneratorObject) SetMessage(value foundation.NSData) {
	objc.Send[struct{}](o.ID, objc.Sel("setMessage:"), value)
}

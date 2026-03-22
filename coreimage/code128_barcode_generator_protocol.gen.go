// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a Code 128 barcode generator filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CICode128BarcodeGenerator
type CICode128BarcodeGenerator interface {
	objectivec.IObject
	CIFilterProtocol

	// The height, in pixels, of the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICode128BarcodeGenerator/barcodeHeight
	BarcodeHeight() float32

	// The message to encode in the Code 128 barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICode128BarcodeGenerator/message
	Message() foundation.INSData

	// The number of empty white pixels that should surround the barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICode128BarcodeGenerator/quietSpace
	QuietSpace() float32

	// The height, in pixels, of the generated barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICode128BarcodeGenerator/barcodeHeight
	SetBarcodeHeight(value float32)

	// The message to encode in the Code 128 barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICode128BarcodeGenerator/message
	SetMessage(value foundation.INSData)

	// The number of empty white pixels that should surround the barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICode128BarcodeGenerator/quietSpace
	SetQuietSpace(value float32)
}

// CICode128BarcodeGeneratorObject wraps an existing Objective-C object that conforms to the CICode128BarcodeGenerator protocol.
type CICode128BarcodeGeneratorObject struct {
	objectivec.Object
}
func (o CICode128BarcodeGeneratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CICode128BarcodeGeneratorObjectFromID constructs a [CICode128BarcodeGeneratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CICode128BarcodeGeneratorObjectFromID(id objc.ID) CICode128BarcodeGeneratorObject {
	return CICode128BarcodeGeneratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The height, in pixels, of the generated barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CICode128BarcodeGenerator/barcodeHeight
func (o CICode128BarcodeGeneratorObject) BarcodeHeight() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("barcodeHeight"))
	return rv
	}
// The message to encode in the Code 128 barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CICode128BarcodeGenerator/message
func (o CICode128BarcodeGeneratorObject) Message() foundation.INSData {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("message"))
	return foundation.NSDataFromID(rv)
	}
// The number of empty white pixels that should surround the barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CICode128BarcodeGenerator/quietSpace
func (o CICode128BarcodeGeneratorObject) QuietSpace() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("quietSpace"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CICode128BarcodeGeneratorObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CICode128BarcodeGeneratorObject) SetBarcodeHeight(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setBarcodeHeight:"), value)
}

func (o CICode128BarcodeGeneratorObject) SetMessage(value foundation.INSData) {
	objc.Send[struct{}](o.ID, objc.Sel("setMessage:"), value)
}

func (o CICode128BarcodeGeneratorObject) SetQuietSpace(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setQuietSpace:"), value)
}


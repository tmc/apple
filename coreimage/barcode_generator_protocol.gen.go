// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a barcode generator filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBarcodeGenerator
type CIBarcodeGenerator interface {
	objectivec.IObject
	CIFilterProtocol

	// The barcode descriptor to generate an image for.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBarcodeGenerator/barcodeDescriptor
	BarcodeDescriptor() ICIBarcodeDescriptor

	// The barcode descriptor to generate an image for.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBarcodeGenerator/barcodeDescriptor
	SetBarcodeDescriptor(value ICIBarcodeDescriptor)
}

// CIBarcodeGeneratorObject wraps an existing Objective-C object that conforms to the CIBarcodeGenerator protocol.
type CIBarcodeGeneratorObject struct {
	objectivec.Object
}
func (o CIBarcodeGeneratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIBarcodeGeneratorObjectFromID constructs a [CIBarcodeGeneratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIBarcodeGeneratorObjectFromID(id objc.ID) CIBarcodeGeneratorObject {
	return CIBarcodeGeneratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The barcode descriptor to generate an image for.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBarcodeGenerator/barcodeDescriptor
func (o CIBarcodeGeneratorObject) BarcodeDescriptor() ICIBarcodeDescriptor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("barcodeDescriptor"))
	return CIBarcodeDescriptorFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIBarcodeGeneratorObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIBarcodeGeneratorObject) SetBarcodeDescriptor(value ICIBarcodeDescriptor) {
	objc.Send[struct{}](o.ID, objc.Sel("setBarcodeDescriptor:"), value)
}


// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure an Aztec code generator filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeGenerator
type CIAztecCodeGenerator interface {
	objectivec.IObject
	CIFilterProtocol

	// A Boolean that specifies whether to force a compact style Aztec code.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeGenerator/compactStyle
	CompactStyle() float32

	// The Aztec error correction, a value from 5 to 95.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeGenerator/correctionLevel
	CorrectionLevel() float32

	// The number of Aztec layers, a value from 1 to 32.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeGenerator/layers
	Layers() float32

	// The message to encode in the Aztec barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeGenerator/message
	Message() foundation.INSData

	// A Boolean that specifies whether to force a compact style Aztec code.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeGenerator/compactStyle
	SetCompactStyle(value float32)

	// The Aztec error correction, a value from 5 to 95.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeGenerator/correctionLevel
	SetCorrectionLevel(value float32)

	// The number of Aztec layers, a value from 1 to 32.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeGenerator/layers
	SetLayers(value float32)

	// The message to encode in the Aztec barcode.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeGenerator/message
	SetMessage(value foundation.INSData)
}

// CIAztecCodeGeneratorObject wraps an existing Objective-C object that conforms to the CIAztecCodeGenerator protocol.
type CIAztecCodeGeneratorObject struct {
	objectivec.Object
}
func (o CIAztecCodeGeneratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIAztecCodeGeneratorObjectFromID constructs a [CIAztecCodeGeneratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIAztecCodeGeneratorObjectFromID(id objc.ID) CIAztecCodeGeneratorObject {
	return CIAztecCodeGeneratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A Boolean that specifies whether to force a compact style Aztec code.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeGenerator/compactStyle
func (o CIAztecCodeGeneratorObject) CompactStyle() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("compactStyle"))
	return rv
	}
// The Aztec error correction, a value from 5 to 95.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeGenerator/correctionLevel
func (o CIAztecCodeGeneratorObject) CorrectionLevel() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("correctionLevel"))
	return rv
	}
// The number of Aztec layers, a value from 1 to 32.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeGenerator/layers
func (o CIAztecCodeGeneratorObject) Layers() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("layers"))
	return rv
	}
// The message to encode in the Aztec barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeGenerator/message
func (o CIAztecCodeGeneratorObject) Message() foundation.INSData {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("message"))
	return foundation.NSDataFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIAztecCodeGeneratorObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIAztecCodeGeneratorObject) SetCompactStyle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setCompactStyle:"), value)
}

func (o CIAztecCodeGeneratorObject) SetCorrectionLevel(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setCorrectionLevel:"), value)
}

func (o CIAztecCodeGeneratorObject) SetLayers(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setLayers:"), value)
}

func (o CIAztecCodeGeneratorObject) SetMessage(value foundation.INSData) {
	objc.Send[struct{}](o.ID, objc.Sel("setMessage:"), value)
}


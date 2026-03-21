// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure an attributed-text image generator filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAttributedTextImageGenerator
type CIAttributedTextImageGenerator interface {
	objectivec.IObject
	CIFilterProtocol

	// The scale at which to render the text.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAttributedTextImageGenerator/scaleFactor
	ScaleFactor() float32

	// The text to render.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAttributedTextImageGenerator/text
	Text() foundation.NSAttributedString

	// Padding protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAttributedTextImageGenerator/padding
	Padding() float32

	// The scale at which to render the text.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAttributedTextImageGenerator/scaleFactor
	SetScaleFactor(value float32)

	// The text to render.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAttributedTextImageGenerator/text
	SetText(value foundation.NSAttributedString)

	// padding protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAttributedTextImageGenerator/padding
	SetPadding(value float32)
}

// CIAttributedTextImageGeneratorObject wraps an existing Objective-C object that conforms to the CIAttributedTextImageGenerator protocol.
type CIAttributedTextImageGeneratorObject struct {
	objectivec.Object
}
func (o CIAttributedTextImageGeneratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIAttributedTextImageGeneratorObjectFromID constructs a [CIAttributedTextImageGeneratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIAttributedTextImageGeneratorObjectFromID(id objc.ID) CIAttributedTextImageGeneratorObject {
	return CIAttributedTextImageGeneratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The scale at which to render the text.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAttributedTextImageGenerator/scaleFactor
func (o CIAttributedTextImageGeneratorObject) ScaleFactor() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("scaleFactor"))
	return rv
	}
// The text to render.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAttributedTextImageGenerator/text
func (o CIAttributedTextImageGeneratorObject) Text() foundation.NSAttributedString {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("text"))
	return foundation.NSAttributedStringFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CIAttributedTextImageGenerator/padding
func (o CIAttributedTextImageGeneratorObject) Padding() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("padding"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIAttributedTextImageGeneratorObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIAttributedTextImageGeneratorObject) SetScaleFactor(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setScaleFactor:"), value)
}

func (o CIAttributedTextImageGeneratorObject) SetText(value foundation.NSAttributedString) {
	objc.Send[struct{}](o.ID, objc.Sel("setText:"), value)
}

func (o CIAttributedTextImageGeneratorObject) SetPadding(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setPadding:"), value)
}


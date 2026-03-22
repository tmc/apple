// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a text image generator filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CITextImageGenerator
type CITextImageGenerator interface {
	objectivec.IObject
	CIFilterProtocol

	// The name of the font to use for the generated text.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITextImageGenerator/fontName
	FontName() string

	// The size of the font to use for the generated text.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITextImageGenerator/fontSize
	FontSize() float32

	// The scale of the font to use for the generated text.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITextImageGenerator/scaleFactor
	ScaleFactor() float32

	// The text to render.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITextImageGenerator/text
	Text() string

	// Padding protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITextImageGenerator/padding
	Padding() float32

	// The name of the font to use for the generated text.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITextImageGenerator/fontName
	SetFontName(value string)

	// The size of the font to use for the generated text.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITextImageGenerator/fontSize
	SetFontSize(value float32)

	// The scale of the font to use for the generated text.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITextImageGenerator/scaleFactor
	SetScaleFactor(value float32)

	// The text to render.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITextImageGenerator/text
	SetText(value string)

	// padding protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITextImageGenerator/padding
	SetPadding(value float32)
}

// CITextImageGeneratorObject wraps an existing Objective-C object that conforms to the CITextImageGenerator protocol.
type CITextImageGeneratorObject struct {
	objectivec.Object
}
func (o CITextImageGeneratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CITextImageGeneratorObjectFromID constructs a [CITextImageGeneratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CITextImageGeneratorObjectFromID(id objc.ID) CITextImageGeneratorObject {
	return CITextImageGeneratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The name of the font to use for the generated text.
//
// See: https://developer.apple.com/documentation/CoreImage/CITextImageGenerator/fontName
func (o CITextImageGeneratorObject) FontName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("fontName"))
	return foundation.NSStringFromID(rv).String()
	}
// The size of the font to use for the generated text.
//
// See: https://developer.apple.com/documentation/CoreImage/CITextImageGenerator/fontSize
func (o CITextImageGeneratorObject) FontSize() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("fontSize"))
	return rv
	}
// The scale of the font to use for the generated text.
//
// See: https://developer.apple.com/documentation/CoreImage/CITextImageGenerator/scaleFactor
func (o CITextImageGeneratorObject) ScaleFactor() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("scaleFactor"))
	return rv
	}
// The text to render.
//
// See: https://developer.apple.com/documentation/CoreImage/CITextImageGenerator/text
func (o CITextImageGeneratorObject) Text() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("text"))
	return foundation.NSStringFromID(rv).String()
	}
// See: https://developer.apple.com/documentation/CoreImage/CITextImageGenerator/padding
func (o CITextImageGeneratorObject) Padding() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("padding"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CITextImageGeneratorObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CITextImageGeneratorObject) SetFontName(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setFontName:"), objc.String(value))
}

func (o CITextImageGeneratorObject) SetFontSize(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setFontSize:"), value)
}

func (o CITextImageGeneratorObject) SetScaleFactor(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setScaleFactor:"), value)
}

func (o CITextImageGeneratorObject) SetText(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setText:"), objc.String(value))
}

func (o CITextImageGeneratorObject) SetPadding(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setPadding:"), value)
}


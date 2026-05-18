// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a palettize filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPalettize
type CIPalettize interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPalettize/inputImage
	InputImage() ICIImage
	SetInputImage(value ICIImage)

	// The input color palette, obtained by using a k-means clustering filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPalettize/paletteImage
	PaletteImage() ICIImage
	SetPaletteImage(value ICIImage)

	// A Boolean value that specifies whether the filter applies the color palette in a perceptual color space.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPalettize/perceptual
	Perceptual() bool
	SetPerceptual(value bool)
}

// CIPalettizeObject wraps an existing Objective-C object that conforms to the CIPalettize protocol.
type CIPalettizeObject struct {
	objectivec.Object
}

func (o CIPalettizeObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIPalettizeObjectFromID constructs a [CIPalettizeObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIPalettizeObjectFromID(id objc.ID) CIPalettizeObject {
	return CIPalettizeObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIPalettizeObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPalettize/inputImage
func (o CIPalettizeObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

func (o CIPalettizeObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The input color palette, obtained by using a k-means clustering filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPalettize/paletteImage
func (o CIPalettizeObject) PaletteImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("paletteImage"))
	return CIImageFromID(rv)
}

func (o CIPalettizeObject) SetPaletteImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setPaletteImage:"), value)
}

// A Boolean value that specifies whether the filter applies the color palette
// in a perceptual color space.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPalettize/perceptual
func (o CIPalettizeObject) Perceptual() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("perceptual"))
	return bool(rv)
}

func (o CIPalettizeObject) SetPerceptual(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setPerceptual:"), value)
}

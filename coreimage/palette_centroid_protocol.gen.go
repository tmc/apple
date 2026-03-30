// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a palette centroid filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPaletteCentroid
type CIPaletteCentroid interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPaletteCentroid/inputImage
	InputImage() ICIImage

	// The input color palette, obtained by using a k-means clustering filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPaletteCentroid/paletteImage
	PaletteImage() ICIImage

	// A Boolean value that specifies whether the filter applies the color palette in a perceptual color space.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPaletteCentroid/perceptual
	Perceptual() bool

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPaletteCentroid/inputImage
	SetInputImage(value ICIImage)

	// The input color palette, obtained by using a k-means clustering filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPaletteCentroid/paletteImage
	SetPaletteImage(value ICIImage)

	// A Boolean value that specifies whether the filter applies the color palette in a perceptual color space.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPaletteCentroid/perceptual
	SetPerceptual(value bool)
}

// CIPaletteCentroidObject wraps an existing Objective-C object that conforms to the CIPaletteCentroid protocol.
type CIPaletteCentroidObject struct {
	objectivec.Object
}

func (o CIPaletteCentroidObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIPaletteCentroidObjectFromID constructs a [CIPaletteCentroidObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIPaletteCentroidObjectFromID(id objc.ID) CIPaletteCentroidObject {
	return CIPaletteCentroidObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPaletteCentroid/inputImage
func (o CIPaletteCentroidObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// The input color palette, obtained by using a k-means clustering filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPaletteCentroid/paletteImage
func (o CIPaletteCentroidObject) PaletteImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("paletteImage"))
	return CIImageFromID(rv)
}

// A Boolean value that specifies whether the filter applies the color palette
// in a perceptual color space.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPaletteCentroid/perceptual
func (o CIPaletteCentroidObject) Perceptual() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("perceptual"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIPaletteCentroidObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPaletteCentroid/inputImage
func (o CIPaletteCentroidObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The input color palette, obtained by using a k-means clustering filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPaletteCentroid/paletteImage
func (o CIPaletteCentroidObject) SetPaletteImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setPaletteImage:"), value)
}

// A Boolean value that specifies whether the filter applies the color palette
// in a perceptual color space.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPaletteCentroid/perceptual
func (o CIPaletteCentroidObject) SetPerceptual(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setPerceptual:"), value)
}

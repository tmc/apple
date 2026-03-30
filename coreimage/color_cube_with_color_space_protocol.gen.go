// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a color cube with color space filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace
type CIColorCubeWithColorSpace interface {
	objectivec.IObject
	CIFilterProtocol

	// The working color space.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/colorSpace
	ColorSpace() coregraphics.CGColorSpaceRef

	// The cube texture data to use as a color lookup table.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/cubeData
	CubeData() foundation.INSData

	// The length, in texels, of each side of the cube texture.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/cubeDimension
	CubeDimension() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/inputImage
	InputImage() ICIImage

	// If true, then the filter extrapolates the color cube for any RGB component values outside the range 0.0 to 1.0.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/extrapolate
	Extrapolate() bool

	// The working color space.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/colorSpace
	SetColorSpace(value coregraphics.CGColorSpaceRef)

	// The cube texture data to use as a color lookup table.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/cubeData
	SetCubeData(value foundation.INSData)

	// The length, in texels, of each side of the cube texture.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/cubeDimension
	SetCubeDimension(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/inputImage
	SetInputImage(value ICIImage)

	// If true, then the filter extrapolates the color cube for any RGB component values outside the range 0.0 to 1.0.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/extrapolate
	SetExtrapolate(value bool)
}

// CIColorCubeWithColorSpaceObject wraps an existing Objective-C object that conforms to the CIColorCubeWithColorSpace protocol.
type CIColorCubeWithColorSpaceObject struct {
	objectivec.Object
}

func (o CIColorCubeWithColorSpaceObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIColorCubeWithColorSpaceObjectFromID constructs a [CIColorCubeWithColorSpaceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIColorCubeWithColorSpaceObjectFromID(id objc.ID) CIColorCubeWithColorSpaceObject {
	return CIColorCubeWithColorSpaceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The working color space.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/colorSpace
func (o CIColorCubeWithColorSpaceObject) ColorSpace() coregraphics.CGColorSpaceRef {
	rv := objc.Send[coregraphics.CGColorSpaceRef](o.ID, objc.Sel("colorSpace"))
	return rv
}

// The cube texture data to use as a color lookup table.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/cubeData
func (o CIColorCubeWithColorSpaceObject) CubeData() foundation.INSData {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("cubeData"))
	return foundation.NSDataFromID(rv)
}

// The length, in texels, of each side of the cube texture.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/cubeDimension
func (o CIColorCubeWithColorSpaceObject) CubeDimension() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("cubeDimension"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/inputImage
func (o CIColorCubeWithColorSpaceObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// If true, then the filter extrapolates the color cube for any RGB component
// values outside the range 0.0 to 1.0.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/extrapolate
func (o CIColorCubeWithColorSpaceObject) Extrapolate() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("extrapolate"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIColorCubeWithColorSpaceObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The working color space.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/colorSpace
func (o CIColorCubeWithColorSpaceObject) SetColorSpace(value coregraphics.CGColorSpaceRef) {
	objc.Send[struct{}](o.ID, objc.Sel("setColorSpace:"), value)
}

// The cube texture data to use as a color lookup table.
//
// # Discussion
//
// This filter maps color values in the input image to new color values using
// a three-dimensional color lookup table. For each RGBA pixel in the input
// image, the filter uses the R, G, and B component values as indices to
// identify a location in the table; the RGBA value at that location becomes
// the RGBA value of the output pixel.
//
// Use the [CubeData] parameter to provide data formatted for use as a color
// lookup table, and the inputCubeDimension parameter to specify the size of
// the table. This data should be an array of texel values in 32-bit
// floating-point RGBA linear premultiplied format. The inputCubeDimension
// parameter identifies the size of the cube by specifying the length of one
// side, so the size of the array should be [CubeDimension] cubed times the
// size of a single texel value. In the color table, the R component varies
// fastest, followed by G, then B.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/cubeData
func (o CIColorCubeWithColorSpaceObject) SetCubeData(value foundation.INSData) {
	objc.Send[struct{}](o.ID, objc.Sel("setCubeData:"), value)
}

// The length, in texels, of each side of the cube texture.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/cubeDimension
func (o CIColorCubeWithColorSpaceObject) SetCubeDimension(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setCubeDimension:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/inputImage
func (o CIColorCubeWithColorSpaceObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// If true, then the filter extrapolates the color cube for any RGB component
// values outside the range 0.0 to 1.0.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCubeWithColorSpace/extrapolate
func (o CIColorCubeWithColorSpaceObject) SetExtrapolate(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtrapolate:"), value)
}

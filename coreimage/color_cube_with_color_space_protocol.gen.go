// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
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

func (o CIColorCubeWithColorSpaceObject) SetColorSpace(value coregraphics.CGColorSpaceRef) {
	objc.Send[struct{}](o.ID, objc.Sel("setColorSpace:"), value)
}

func (o CIColorCubeWithColorSpaceObject) SetCubeData(value foundation.INSData) {
	objc.Send[struct{}](o.ID, objc.Sel("setCubeData:"), value)
}

func (o CIColorCubeWithColorSpaceObject) SetCubeDimension(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setCubeDimension:"), value)
}

func (o CIColorCubeWithColorSpaceObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIColorCubeWithColorSpaceObject) SetExtrapolate(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtrapolate:"), value)
}


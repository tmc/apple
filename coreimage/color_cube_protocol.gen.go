// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a color cube filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCube
type CIColorCube interface {
	objectivec.IObject
	CIFilterProtocol

	// The cube texture data to use as a color lookup table.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCube/cubeData
	CubeData() foundation.INSData

	// The length, in texels, of each side of the cube texture.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCube/cubeDimension
	CubeDimension() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCube/inputImage
	InputImage() ICIImage

	// If true, then the filter extrapolates the color cube for any RGB component values outside the range 0.0 to 1.0.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCube/extrapolate
	Extrapolate() bool

	// The cube texture data to use as a color lookup table.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCube/cubeData
	SetCubeData(value foundation.INSData)

	// The length, in texels, of each side of the cube texture.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCube/cubeDimension
	SetCubeDimension(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCube/inputImage
	SetInputImage(value ICIImage)

	// If true, then the filter extrapolates the color cube for any RGB component values outside the range 0.0 to 1.0.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCube/extrapolate
	SetExtrapolate(value bool)
}

// CIColorCubeObject wraps an existing Objective-C object that conforms to the CIColorCube protocol.
type CIColorCubeObject struct {
	objectivec.Object
}
func (o CIColorCubeObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIColorCubeObjectFromID constructs a [CIColorCubeObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIColorCubeObjectFromID(id objc.ID) CIColorCubeObject {
	return CIColorCubeObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The cube texture data to use as a color lookup table.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCube/cubeData
func (o CIColorCubeObject) CubeData() foundation.INSData {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("cubeData"))
	return foundation.NSDataFromID(rv)
	}
// The length, in texels, of each side of the cube texture.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCube/cubeDimension
func (o CIColorCubeObject) CubeDimension() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("cubeDimension"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCube/inputImage
func (o CIColorCubeObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// If true, then the filter extrapolates the color cube for any RGB component
// values outside the range 0.0 to 1.0.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCube/extrapolate
func (o CIColorCubeObject) Extrapolate() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("extrapolate"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIColorCubeObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIColorCubeObject) SetCubeData(value foundation.INSData) {
	objc.Send[struct{}](o.ID, objc.Sel("setCubeData:"), value)
}

func (o CIColorCubeObject) SetCubeDimension(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setCubeDimension:"), value)
}

func (o CIColorCubeObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIColorCubeObject) SetExtrapolate(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtrapolate:"), value)
}


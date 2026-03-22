// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a hexagonal pixellate filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHexagonalPixellate
type CIHexagonalPixellate interface {
	objectivec.IObject
	CIFilterProtocol

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHexagonalPixellate/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHexagonalPixellate/inputImage
	InputImage() ICIImage

	// The size of the hexagons.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHexagonalPixellate/scale
	Scale() float32

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHexagonalPixellate/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHexagonalPixellate/inputImage
	SetInputImage(value ICIImage)

	// The size of the hexagons.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHexagonalPixellate/scale
	SetScale(value float32)
}

// CIHexagonalPixellateObject wraps an existing Objective-C object that conforms to the CIHexagonalPixellate protocol.
type CIHexagonalPixellateObject struct {
	objectivec.Object
}
func (o CIHexagonalPixellateObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIHexagonalPixellateObjectFromID constructs a [CIHexagonalPixellateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIHexagonalPixellateObjectFromID(id objc.ID) CIHexagonalPixellateObject {
	return CIHexagonalPixellateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The x and y position to use as the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHexagonalPixellate/center
func (o CIHexagonalPixellateObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHexagonalPixellate/inputImage
func (o CIHexagonalPixellateObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The size of the hexagons.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHexagonalPixellate/scale
func (o CIHexagonalPixellateObject) Scale() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("scale"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIHexagonalPixellateObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIHexagonalPixellateObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIHexagonalPixellateObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIHexagonalPixellateObject) SetScale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setScale:"), value)
}


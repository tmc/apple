// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a mesh generator filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMeshGenerator
type CIMeshGenerator interface {
	objectivec.IObject
	CIFilterProtocol

	// The color of the rendered mesh.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMeshGenerator/color
	Color() ICIColor

	// An array that describes the mesh to render.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMeshGenerator/mesh
	Mesh() foundation.INSArray

	// The width of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMeshGenerator/width
	Width() float32

	// The color of the rendered mesh.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMeshGenerator/color
	SetColor(value ICIColor)

	// An array that describes the mesh to render.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMeshGenerator/mesh
	SetMesh(value foundation.INSArray)

	// The width of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMeshGenerator/width
	SetWidth(value float32)
}

// CIMeshGeneratorObject wraps an existing Objective-C object that conforms to the CIMeshGenerator protocol.
type CIMeshGeneratorObject struct {
	objectivec.Object
}
func (o CIMeshGeneratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIMeshGeneratorObjectFromID constructs a [CIMeshGeneratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIMeshGeneratorObjectFromID(id objc.ID) CIMeshGeneratorObject {
	return CIMeshGeneratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The color of the rendered mesh.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMeshGenerator/color
func (o CIMeshGeneratorObject) Color() ICIColor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color"))
	return CIColorFromID(rv)
	}
// An array that describes the mesh to render.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMeshGenerator/mesh
func (o CIMeshGeneratorObject) Mesh() foundation.INSArray {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("mesh"))
	return foundation.NSArrayFromID(rv)
	}
// The width of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMeshGenerator/width
func (o CIMeshGeneratorObject) Width() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIMeshGeneratorObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIMeshGeneratorObject) SetColor(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor:"), value)
}

func (o CIMeshGeneratorObject) SetMesh(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setMesh:"), value)
}

func (o CIMeshGeneratorObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}


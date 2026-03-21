// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CIDroste protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDroste
type CIDroste interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDroste/inputImage
	InputImage() ICIImage

	// InsetPoint0 protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDroste/insetPoint0
	InsetPoint0() corefoundation.CGPoint

	// InsetPoint1 protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDroste/insetPoint1
	InsetPoint1() corefoundation.CGPoint

	// Periodicity protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDroste/periodicity
	Periodicity() float32

	// Rotation protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDroste/rotation
	Rotation() float32

	// Strands protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDroste/strands
	Strands() float32

	// Zoom protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDroste/zoom
	Zoom() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDroste/inputImage
	SetInputImage(value ICIImage)

	// insetPoint0 protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDroste/insetPoint0
	SetInsetPoint0(value corefoundation.CGPoint)

	// insetPoint1 protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDroste/insetPoint1
	SetInsetPoint1(value corefoundation.CGPoint)

	// periodicity protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDroste/periodicity
	SetPeriodicity(value float32)

	// rotation protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDroste/rotation
	SetRotation(value float32)

	// strands protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDroste/strands
	SetStrands(value float32)

	// zoom protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDroste/zoom
	SetZoom(value float32)
}

// CIDrosteObject wraps an existing Objective-C object that conforms to the CIDroste protocol.
type CIDrosteObject struct {
	objectivec.Object
}
func (o CIDrosteObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIDrosteObjectFromID constructs a [CIDrosteObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIDrosteObjectFromID(id objc.ID) CIDrosteObject {
	return CIDrosteObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDroste/inputImage
func (o CIDrosteObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CIDroste/insetPoint0
func (o CIDrosteObject) InsetPoint0() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("insetPoint0"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIDroste/insetPoint1
func (o CIDrosteObject) InsetPoint1() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("insetPoint1"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIDroste/periodicity
func (o CIDrosteObject) Periodicity() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("periodicity"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIDroste/rotation
func (o CIDrosteObject) Rotation() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("rotation"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIDroste/strands
func (o CIDrosteObject) Strands() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("strands"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIDroste/zoom
func (o CIDrosteObject) Zoom() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("zoom"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIDrosteObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIDrosteObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIDrosteObject) SetInsetPoint0(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setInsetPoint0:"), value)
}

func (o CIDrosteObject) SetInsetPoint1(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setInsetPoint1:"), value)
}

func (o CIDrosteObject) SetPeriodicity(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setPeriodicity:"), value)
}

func (o CIDrosteObject) SetRotation(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRotation:"), value)
}

func (o CIDrosteObject) SetStrands(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setStrands:"), value)
}

func (o CIDrosteObject) SetZoom(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setZoom:"), value)
}


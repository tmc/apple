// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CINinePartStretched protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CINinePartStretched
type CINinePartStretched interface {
	objectivec.IObject
	CIFilterProtocol

	// Breakpoint0 protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINinePartStretched/breakpoint0
	Breakpoint0() corefoundation.CGPoint

	// Breakpoint1 protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINinePartStretched/breakpoint1
	Breakpoint1() corefoundation.CGPoint

	// GrowAmount protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINinePartStretched/growAmount
	GrowAmount() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINinePartStretched/inputImage
	InputImage() ICIImage

	// breakpoint0 protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINinePartStretched/breakpoint0
	SetBreakpoint0(value corefoundation.CGPoint)

	// breakpoint1 protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINinePartStretched/breakpoint1
	SetBreakpoint1(value corefoundation.CGPoint)

	// growAmount protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINinePartStretched/growAmount
	SetGrowAmount(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINinePartStretched/inputImage
	SetInputImage(value ICIImage)
}

// CINinePartStretchedObject wraps an existing Objective-C object that conforms to the CINinePartStretched protocol.
type CINinePartStretchedObject struct {
	objectivec.Object
}
func (o CINinePartStretchedObject) BaseObject() objectivec.Object {
	return o.Object
}

// CINinePartStretchedObjectFromID constructs a [CINinePartStretchedObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CINinePartStretchedObjectFromID(id objc.ID) CINinePartStretchedObject {
	return CINinePartStretchedObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CINinePartStretched/breakpoint0
func (o CINinePartStretchedObject) Breakpoint0() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("breakpoint0"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CINinePartStretched/breakpoint1
func (o CINinePartStretchedObject) Breakpoint1() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("breakpoint1"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CINinePartStretched/growAmount
func (o CINinePartStretchedObject) GrowAmount() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("growAmount"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CINinePartStretched/inputImage
func (o CINinePartStretchedObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CINinePartStretchedObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CINinePartStretchedObject) SetBreakpoint0(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBreakpoint0:"), value)
}

func (o CINinePartStretchedObject) SetBreakpoint1(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBreakpoint1:"), value)
}

func (o CINinePartStretchedObject) SetGrowAmount(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setGrowAmount:"), value)
}

func (o CINinePartStretchedObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}


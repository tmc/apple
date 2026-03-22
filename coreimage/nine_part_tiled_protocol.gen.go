// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CINinePartTiled protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CINinePartTiled
type CINinePartTiled interface {
	objectivec.IObject
	CIFilterProtocol

	// Breakpoint0 protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINinePartTiled/breakpoint0
	Breakpoint0() corefoundation.CGPoint

	// Breakpoint1 protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINinePartTiled/breakpoint1
	Breakpoint1() corefoundation.CGPoint

	// FlipYTiles protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINinePartTiled/flipYTiles
	FlipYTiles() bool

	// GrowAmount protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINinePartTiled/growAmount
	GrowAmount() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINinePartTiled/inputImage
	InputImage() ICIImage

	// breakpoint0 protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINinePartTiled/breakpoint0
	SetBreakpoint0(value corefoundation.CGPoint)

	// breakpoint1 protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINinePartTiled/breakpoint1
	SetBreakpoint1(value corefoundation.CGPoint)

	// flipYTiles protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINinePartTiled/flipYTiles
	SetFlipYTiles(value bool)

	// growAmount protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINinePartTiled/growAmount
	SetGrowAmount(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINinePartTiled/inputImage
	SetInputImage(value ICIImage)
}

// CINinePartTiledObject wraps an existing Objective-C object that conforms to the CINinePartTiled protocol.
type CINinePartTiledObject struct {
	objectivec.Object
}
func (o CINinePartTiledObject) BaseObject() objectivec.Object {
	return o.Object
}

// CINinePartTiledObjectFromID constructs a [CINinePartTiledObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CINinePartTiledObjectFromID(id objc.ID) CINinePartTiledObject {
	return CINinePartTiledObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CINinePartTiled/breakpoint0
func (o CINinePartTiledObject) Breakpoint0() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("breakpoint0"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CINinePartTiled/breakpoint1
func (o CINinePartTiledObject) Breakpoint1() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("breakpoint1"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CINinePartTiled/flipYTiles
func (o CINinePartTiledObject) FlipYTiles() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("flipYTiles"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CINinePartTiled/growAmount
func (o CINinePartTiledObject) GrowAmount() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("growAmount"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CINinePartTiled/inputImage
func (o CINinePartTiledObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CINinePartTiledObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CINinePartTiledObject) SetBreakpoint0(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBreakpoint0:"), value)
}

func (o CINinePartTiledObject) SetBreakpoint1(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBreakpoint1:"), value)
}

func (o CINinePartTiledObject) SetFlipYTiles(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setFlipYTiles:"), value)
}

func (o CINinePartTiledObject) SetGrowAmount(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setGrowAmount:"), value)
}

func (o CINinePartTiledObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}


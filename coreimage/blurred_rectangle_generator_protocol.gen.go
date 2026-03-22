// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CIBlurredRectangleGenerator protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRectangleGenerator
type CIBlurredRectangleGenerator interface {
	objectivec.IObject
	CIFilterProtocol

	// Color protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRectangleGenerator/color
	Color() ICIColor

	// Extent protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRectangleGenerator/extent
	Extent() corefoundation.CGRect

	// Sigma protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRectangleGenerator/sigma
	Sigma() float32

	// color protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRectangleGenerator/color
	SetColor(value ICIColor)

	// extent protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRectangleGenerator/extent
	SetExtent(value corefoundation.CGRect)

	// sigma protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRectangleGenerator/sigma
	SetSigma(value float32)
}

// CIBlurredRectangleGeneratorObject wraps an existing Objective-C object that conforms to the CIBlurredRectangleGenerator protocol.
type CIBlurredRectangleGeneratorObject struct {
	objectivec.Object
}
func (o CIBlurredRectangleGeneratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIBlurredRectangleGeneratorObjectFromID constructs a [CIBlurredRectangleGeneratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIBlurredRectangleGeneratorObjectFromID(id objc.ID) CIBlurredRectangleGeneratorObject {
	return CIBlurredRectangleGeneratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRectangleGenerator/color
func (o CIBlurredRectangleGeneratorObject) Color() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color"))
	return CIColorFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRectangleGenerator/extent
func (o CIBlurredRectangleGeneratorObject) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRectangleGenerator/sigma
func (o CIBlurredRectangleGeneratorObject) Sigma() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("sigma"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIBlurredRectangleGeneratorObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIBlurredRectangleGeneratorObject) SetColor(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor:"), value)
}

func (o CIBlurredRectangleGeneratorObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

func (o CIBlurredRectangleGeneratorObject) SetSigma(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSigma:"), value)
}


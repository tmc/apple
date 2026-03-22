// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CIAreaMinMaxRed protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaMinMaxRed
type CIAreaMinMaxRed interface {
	objectivec.IObject
	CIAreaReductionFilter
	CIFilterProtocol
}

// CIAreaMinMaxRedObject wraps an existing Objective-C object that conforms to the CIAreaMinMaxRed protocol.
type CIAreaMinMaxRedObject struct {
	objectivec.Object
}
func (o CIAreaMinMaxRedObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIAreaMinMaxRedObjectFromID constructs a [CIAreaMinMaxRedObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIAreaMinMaxRedObjectFromID(id objc.ID) CIAreaMinMaxRedObject {
	return CIAreaMinMaxRedObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/extent
func (o CIAreaMinMaxRedObject) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/inputImage
func (o CIAreaMinMaxRedObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIAreaMinMaxRedObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIAreaMinMaxRedObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

func (o CIAreaMinMaxRedObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}


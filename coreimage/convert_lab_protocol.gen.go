// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CIConvertLab protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIConvertLab
type CIConvertLab interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIConvertLab/inputImage
	InputImage() ICIImage

	// Normalize protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIConvertLab/normalize
	Normalize() bool

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIConvertLab/inputImage
	SetInputImage(value ICIImage)

	// normalize protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIConvertLab/normalize
	SetNormalize(value bool)
}

// CIConvertLabObject wraps an existing Objective-C object that conforms to the CIConvertLab protocol.
type CIConvertLabObject struct {
	objectivec.Object
}
func (o CIConvertLabObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIConvertLabObjectFromID constructs a [CIConvertLabObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIConvertLabObjectFromID(id objc.ID) CIConvertLabObject {
	return CIConvertLabObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIConvertLab/inputImage
func (o CIConvertLabObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CIConvertLab/normalize
func (o CIConvertLabObject) Normalize() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("normalize"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIConvertLabObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIConvertLabObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIConvertLabObject) SetNormalize(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setNormalize:"), value)
}


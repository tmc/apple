// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a color clamp filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorClamp
type CIColorClamp interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorClamp/inputImage
	InputImage() ICIImage

	// A vector containing the higher clamping values.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorClamp/maxComponents
	MaxComponents() ICIVector

	// A vector containing the lower clamping values.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorClamp/minComponents
	MinComponents() ICIVector

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorClamp/inputImage
	SetInputImage(value ICIImage)

	// A vector containing the higher clamping values.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorClamp/maxComponents
	SetMaxComponents(value ICIVector)

	// A vector containing the lower clamping values.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorClamp/minComponents
	SetMinComponents(value ICIVector)
}

// CIColorClampObject wraps an existing Objective-C object that conforms to the CIColorClamp protocol.
type CIColorClampObject struct {
	objectivec.Object
}
func (o CIColorClampObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIColorClampObjectFromID constructs a [CIColorClampObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIColorClampObjectFromID(id objc.ID) CIColorClampObject {
	return CIColorClampObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorClamp/inputImage
func (o CIColorClampObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A vector containing the higher clamping values.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorClamp/maxComponents
func (o CIColorClampObject) MaxComponents() ICIVector {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("maxComponents"))
	return CIVectorFromID(rv)
	}
// A vector containing the lower clamping values.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorClamp/minComponents
func (o CIColorClampObject) MinComponents() ICIVector {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("minComponents"))
	return CIVectorFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIColorClampObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIColorClampObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIColorClampObject) SetMaxComponents(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setMaxComponents:"), value)
}

func (o CIColorClampObject) SetMinComponents(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setMinComponents:"), value)
}


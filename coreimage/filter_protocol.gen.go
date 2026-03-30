// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a Core Image filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol
type CIFilterProtocol interface {
	objectivec.IObject

	// A [CIImage](<doc://com.apple.coreimage/documentation/CoreImage/CIImage>) object that encapsulates the operations configured in the filter.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
	OutputImage() ICIImage
}

// CIFilterProtocolObject wraps an existing Objective-C object that conforms to the CIFilterProtocol protocol.
type CIFilterProtocolObject struct {
	objectivec.Object
}

func (o CIFilterProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIFilterProtocolObjectFromID constructs a [CIFilterProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIFilterProtocolObjectFromID(id objc.ID) CIFilterProtocolObject {
	return CIFilterProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIFilterProtocolObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

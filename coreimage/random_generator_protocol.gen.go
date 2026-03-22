// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a random generator filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRandomGenerator
type CIRandomGenerator interface {
	objectivec.IObject
	CIFilterProtocol
}

// CIRandomGeneratorObject wraps an existing Objective-C object that conforms to the CIRandomGenerator protocol.
type CIRandomGeneratorObject struct {
	objectivec.Object
}
func (o CIRandomGeneratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIRandomGeneratorObjectFromID constructs a [CIRandomGeneratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIRandomGeneratorObjectFromID(id objc.ID) CIRandomGeneratorObject {
	return CIRandomGeneratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIRandomGeneratorObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}


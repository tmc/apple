// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol for specifying the revision number of Vision algorithms.
//
// See: https://developer.apple.com/documentation/Vision/VNRequestRevisionProviding
type VNRequestRevisionProviding interface {
	objectivec.IObject

	// The revision of the [VNRequest](<doc://Vision/documentation/Vision/VNRequest>) subclass used to generate the implementing object.
	//
	// See: https://developer.apple.com/documentation/Vision/VNRequestRevisionProviding/requestRevision
	RequestRevision() uint
}

// VNRequestRevisionProvidingObject wraps an existing Objective-C object that conforms to the VNRequestRevisionProviding protocol.
type VNRequestRevisionProvidingObject struct {
	objectivec.Object
}
func (o VNRequestRevisionProvidingObject) BaseObject() objectivec.Object {
	return o.Object
}

// VNRequestRevisionProvidingObjectFromID constructs a [VNRequestRevisionProvidingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VNRequestRevisionProvidingObjectFromID(id objc.ID) VNRequestRevisionProvidingObject {
	return VNRequestRevisionProvidingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The revision of the [VNRequest] subclass used to generate the implementing
// object.
//
// See: https://developer.apple.com/documentation/Vision/VNRequestRevisionProviding/requestRevision
func (o VNRequestRevisionProvidingObject) RequestRevision() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("requestRevision"))
	return rv
	}


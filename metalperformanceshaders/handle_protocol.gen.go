// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The protocol that provides resource identification.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSHandle
type MPSHandle interface {
	objectivec.IObject
	foundation.NSCoding
	foundation.NSSecureCoding

	// Label protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSHandle/label()
	Label() string
}

// MPSHandleObject wraps an existing Objective-C object that conforms to the MPSHandle protocol.
type MPSHandleObject struct {
	foundation.NSCodingObject
}

func (o MPSHandleObject) BaseObject() objectivec.Object {
	return o.NSCodingObject.BaseObject()
}

// MPSHandleObjectFromID constructs a [MPSHandleObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MPSHandleObjectFromID(id objc.ID) MPSHandleObject {
	return MPSHandleObject{
		NSCodingObject: foundation.NSCodingObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSHandle/label()
func (o MPSHandleObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

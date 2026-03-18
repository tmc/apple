// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol to identify types that customize how the Metal compiler stitches a function together.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingAttribute
type MTLFunctionStitchingAttribute interface {
	objectivec.IObject
}

// MTLFunctionStitchingAttributeObject wraps an existing Objective-C object that conforms to the MTLFunctionStitchingAttribute protocol.
type MTLFunctionStitchingAttributeObject struct {
	objectivec.Object
}
func (o MTLFunctionStitchingAttributeObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLFunctionStitchingAttributeObjectFromID constructs a [MTLFunctionStitchingAttributeObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLFunctionStitchingAttributeObjectFromID(id objc.ID) MTLFunctionStitchingAttributeObject {
	return MTLFunctionStitchingAttributeObject{
		Object: objectivec.ObjectFromID(id),
	}
}


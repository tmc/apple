// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A protocol to identify call graph nodes.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingNode
type MTLFunctionStitchingNode interface {
	objectivec.IObject
	foundation.NSCopying
}



// MTLFunctionStitchingNodeObject wraps an existing Objective-C object that conforms to the MTLFunctionStitchingNode protocol.
type MTLFunctionStitchingNodeObject struct {
	foundation.NSCopyingObject
}
func (o MTLFunctionStitchingNodeObject) BaseObject() objectivec.Object {
	return o.NSCopyingObject.BaseObject()
}



// MTLFunctionStitchingNodeObjectFromID constructs a [MTLFunctionStitchingNodeObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLFunctionStitchingNodeObjectFromID(id objc.ID) MTLFunctionStitchingNodeObject {
	return MTLFunctionStitchingNodeObject{
		NSCopyingObject: foundation.NSCopyingObjectFromID(id),
	}
}











// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The general interface for managing storage of vertex and index data used in loading, processing, and rendering meshes.
//
// See: https://developer.apple.com/documentation/ModelIO/MDLMeshBuffer
type MDLMeshBuffer interface {
	objectivec.IObject
	foundation.NSCopying
}

// MDLMeshBufferObject wraps an existing Objective-C object that conforms to the MDLMeshBuffer protocol.
type MDLMeshBufferObject struct {
	foundation.NSCopyingObject
}

func (o MDLMeshBufferObject) BaseObject() objectivec.Object {
	return o.NSCopyingObject.BaseObject()
}

// MDLMeshBufferObjectFromID constructs a [MDLMeshBufferObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MDLMeshBufferObjectFromID(id objc.ID) MDLMeshBufferObject {
	return MDLMeshBufferObject{
		NSCopyingObject: foundation.NSCopyingObjectFromID(id),
	}
}

// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The general interface for managing allocation of data buffers to be used in loading, processing, and rendering meshes.
//
// See: https://developer.apple.com/documentation/ModelIO/MDLMeshBufferAllocator
type MDLMeshBufferAllocator interface {
	objectivec.IObject
}

// MDLMeshBufferAllocatorObject wraps an existing Objective-C object that conforms to the MDLMeshBufferAllocator protocol.
type MDLMeshBufferAllocatorObject struct {
	objectivec.Object
}

func (o MDLMeshBufferAllocatorObject) BaseObject() objectivec.Object {
	return o.Object
}

// MDLMeshBufferAllocatorObjectFromID constructs a [MDLMeshBufferAllocatorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MDLMeshBufferAllocatorObjectFromID(id objc.ID) MDLMeshBufferAllocatorObject {
	return MDLMeshBufferAllocatorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

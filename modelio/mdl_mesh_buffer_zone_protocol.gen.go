// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The general interface for logical pools of memory used in allocation of related mesh data buffers.
//
// See: https://developer.apple.com/documentation/ModelIO/MDLMeshBufferZone
type MDLMeshBufferZone interface {
	objectivec.IObject
}

// MDLMeshBufferZoneObject wraps an existing Objective-C object that conforms to the MDLMeshBufferZone protocol.
type MDLMeshBufferZoneObject struct {
	objectivec.Object
}

func (o MDLMeshBufferZoneObject) BaseObject() objectivec.Object {
	return o.Object
}

// MDLMeshBufferZoneObjectFromID constructs a [MDLMeshBufferZoneObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MDLMeshBufferZoneObjectFromID(id objc.ID) MDLMeshBufferZoneObject {
	return MDLMeshBufferZoneObject{
		Object: objectivec.ObjectFromID(id),
	}
}

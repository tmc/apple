// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// BSDescriptionStreaming protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/BSDescriptionStreaming
type BSDescriptionStreaming interface {
	objectivec.IObject
}

// BSDescriptionStreamingObject wraps an existing Objective-C object that conforms to the BSDescriptionStreaming protocol.
type BSDescriptionStreamingObject struct {
	objectivec.Object
}

func (o BSDescriptionStreamingObject) BaseObject() objectivec.Object {
	return o.Object
}

// BSDescriptionStreamingObjectFromID constructs a [BSDescriptionStreamingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func BSDescriptionStreamingObjectFromID(id objc.ID) BSDescriptionStreamingObject {
	return BSDescriptionStreamingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/BSDescriptionStreaming/appendDescriptionToStream:
func (o BSDescriptionStreamingObject) AppendDescriptionToStream(stream objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("appendDescriptionToStream:"), stream)
}

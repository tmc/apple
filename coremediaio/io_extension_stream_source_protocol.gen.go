// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol for objects that act as stream sources.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionStreamSource
type CMIOExtensionStreamSource interface {
	objectivec.IObject
}

// CMIOExtensionStreamSourceObject wraps an existing Objective-C object that conforms to the CMIOExtensionStreamSource protocol.
type CMIOExtensionStreamSourceObject struct {
	objectivec.Object
}

func (o CMIOExtensionStreamSourceObject) BaseObject() objectivec.Object {
	return o.Object
}

// CMIOExtensionStreamSourceObjectFromID constructs a [CMIOExtensionStreamSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CMIOExtensionStreamSourceObjectFromID(id objc.ID) CMIOExtensionStreamSourceObject {
	return CMIOExtensionStreamSourceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

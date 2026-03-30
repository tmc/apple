// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol for objects that act as provider sources.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOExtensionProviderSource
type CMIOExtensionProviderSource interface {
	objectivec.IObject
}

// CMIOExtensionProviderSourceObject wraps an existing Objective-C object that conforms to the CMIOExtensionProviderSource protocol.
type CMIOExtensionProviderSourceObject struct {
	objectivec.Object
}

func (o CMIOExtensionProviderSourceObject) BaseObject() objectivec.Object {
	return o.Object
}

// CMIOExtensionProviderSourceObjectFromID constructs a [CMIOExtensionProviderSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CMIOExtensionProviderSourceObjectFromID(id objc.ID) CMIOExtensionProviderSourceObject {
	return CMIOExtensionProviderSourceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

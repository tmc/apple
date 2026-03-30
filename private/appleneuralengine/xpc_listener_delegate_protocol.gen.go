// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSXPCListenerDelegate protocol.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/NSXPCListenerDelegate
type NSXPCListenerDelegate interface {
	objectivec.IObject
}

// NSXPCListenerDelegateObject wraps an existing Objective-C object that conforms to the NSXPCListenerDelegate protocol.
type NSXPCListenerDelegateObject struct {
	objectivec.Object
}

func (o NSXPCListenerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSXPCListenerDelegateObjectFromID constructs a [NSXPCListenerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSXPCListenerDelegateObjectFromID(id objc.ID) NSXPCListenerDelegateObject {
	return NSXPCListenerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/NSXPCListenerDelegate/listener:shouldAcceptNewConnection:
func (o NSXPCListenerDelegateObject) ListenerShouldAcceptNewConnection(listener objectivec.IObject, connection objectivec.IObject) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("listener:shouldAcceptNewConnection:"), listener, connection)
	return rv
}

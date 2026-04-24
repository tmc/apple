// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// BSServiceConnectionListenerDelegate protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/BSServiceConnectionListenerDelegate
type BSServiceConnectionListenerDelegate interface {
	objectivec.IObject
}

// BSServiceConnectionListenerDelegateObject wraps an existing Objective-C object that conforms to the BSServiceConnectionListenerDelegate protocol.
type BSServiceConnectionListenerDelegateObject struct {
	objectivec.Object
}

func (o BSServiceConnectionListenerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// BSServiceConnectionListenerDelegateObjectFromID constructs a [BSServiceConnectionListenerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func BSServiceConnectionListenerDelegateObjectFromID(id objc.ID) BSServiceConnectionListenerDelegateObject {
	return BSServiceConnectionListenerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/BSServiceConnectionListenerDelegate/listener:didReceiveConnection:withContext:
func (o BSServiceConnectionListenerDelegateObject) ListenerDidReceiveConnectionWithContext(listener objectivec.IObject, connection objectivec.IObject, context objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("listener:didReceiveConnection:withContext:"), listener, connection, context)
}

// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// BSServiceListenerConnectionEventObserver protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/BSServiceListenerConnectionEventObserver
type BSServiceListenerConnectionEventObserver interface {
	objectivec.IObject
}

// BSServiceListenerConnectionEventObserverObject wraps an existing Objective-C object that conforms to the BSServiceListenerConnectionEventObserver protocol.
type BSServiceListenerConnectionEventObserverObject struct {
	objectivec.Object
}

func (o BSServiceListenerConnectionEventObserverObject) BaseObject() objectivec.Object {
	return o.Object
}

// BSServiceListenerConnectionEventObserverObjectFromID constructs a [BSServiceListenerConnectionEventObserverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func BSServiceListenerConnectionEventObserverObjectFromID(id objc.ID) BSServiceListenerConnectionEventObserverObject {
	return BSServiceListenerConnectionEventObserverObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/BSServiceListenerConnectionEventObserver/connection:revokedWithEvent:
func (o BSServiceListenerConnectionEventObserverObject) ConnectionRevokedWithEvent(connection objectivec.IObject, event objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("connection:revokedWithEvent:"), connection, event)
}

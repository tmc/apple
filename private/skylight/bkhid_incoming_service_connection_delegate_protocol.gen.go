// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// BKHIDIncomingServiceConnectionDelegate protocol.
type BKHIDIncomingServiceConnectionDelegate interface {
	objectivec.IObject

	// IncomingServiceConnectionDidRevoke protocol.
	IncomingServiceConnectionDidRevoke(revoke objectivec.IObject)
}

// BKHIDIncomingServiceConnectionDelegateObject wraps an existing Objective-C object that conforms to the BKHIDIncomingServiceConnectionDelegate protocol.
type BKHIDIncomingServiceConnectionDelegateObject struct {
	objectivec.Object
}

func (o BKHIDIncomingServiceConnectionDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// BKHIDIncomingServiceConnectionDelegateObjectFromID constructs a [BKHIDIncomingServiceConnectionDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func BKHIDIncomingServiceConnectionDelegateObjectFromID(id objc.ID) BKHIDIncomingServiceConnectionDelegateObject {
	return BKHIDIncomingServiceConnectionDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o BKHIDIncomingServiceConnectionDelegateObject) IncomingServiceConnectionDidRevoke(revoke objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("incomingServiceConnectionDidRevoke:"), revoke)
}

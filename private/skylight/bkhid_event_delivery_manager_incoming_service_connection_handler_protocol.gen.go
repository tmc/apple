// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// BKHIDEventDeliveryManagerIncomingServiceConnectionHandler protocol.
type BKHIDEventDeliveryManagerIncomingServiceConnectionHandler interface {
	objectivec.IObject

	// HandleIncomingDeliveryManagerConnection protocol.
	HandleIncomingDeliveryManagerConnection(connection objectivec.IObject)
}

// BKHIDEventDeliveryManagerIncomingServiceConnectionHandlerObject wraps an existing Objective-C object that conforms to the BKHIDEventDeliveryManagerIncomingServiceConnectionHandler protocol.
type BKHIDEventDeliveryManagerIncomingServiceConnectionHandlerObject struct {
	objectivec.Object
}

func (o BKHIDEventDeliveryManagerIncomingServiceConnectionHandlerObject) BaseObject() objectivec.Object {
	return o.Object
}

// BKHIDEventDeliveryManagerIncomingServiceConnectionHandlerObjectFromID constructs a [BKHIDEventDeliveryManagerIncomingServiceConnectionHandlerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func BKHIDEventDeliveryManagerIncomingServiceConnectionHandlerObjectFromID(id objc.ID) BKHIDEventDeliveryManagerIncomingServiceConnectionHandlerObject {
	return BKHIDEventDeliveryManagerIncomingServiceConnectionHandlerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o BKHIDEventDeliveryManagerIncomingServiceConnectionHandlerObject) HandleIncomingDeliveryManagerConnection(connection objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("handleIncomingDeliveryManagerConnection:"), connection)
}

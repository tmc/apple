// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// BKHIDEventDeliveryObserverIncomingServiceConnectionHandler protocol.
type BKHIDEventDeliveryObserverIncomingServiceConnectionHandler interface {
	objectivec.IObject

	// HandleIncomingDeliveryObserverConnection protocol.
	HandleIncomingDeliveryObserverConnection(connection objectivec.IObject)
}

// BKHIDEventDeliveryObserverIncomingServiceConnectionHandlerObject wraps an existing Objective-C object that conforms to the BKHIDEventDeliveryObserverIncomingServiceConnectionHandler protocol.
type BKHIDEventDeliveryObserverIncomingServiceConnectionHandlerObject struct {
	objectivec.Object
}

func (o BKHIDEventDeliveryObserverIncomingServiceConnectionHandlerObject) BaseObject() objectivec.Object {
	return o.Object
}

// BKHIDEventDeliveryObserverIncomingServiceConnectionHandlerObjectFromID constructs a [BKHIDEventDeliveryObserverIncomingServiceConnectionHandlerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func BKHIDEventDeliveryObserverIncomingServiceConnectionHandlerObjectFromID(id objc.ID) BKHIDEventDeliveryObserverIncomingServiceConnectionHandlerObject {
	return BKHIDEventDeliveryObserverIncomingServiceConnectionHandlerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o BKHIDEventDeliveryObserverIncomingServiceConnectionHandlerObject) HandleIncomingDeliveryObserverConnection(connection objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("handleIncomingDeliveryObserverConnection:"), connection)
}

// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// BKHIDEventDeliveryObserverServiceProvider protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/BKHIDEventDeliveryObserverServiceProvider
type BKHIDEventDeliveryObserverServiceProvider interface {
	objectivec.IObject
}

// BKHIDEventDeliveryObserverServiceProviderObject wraps an existing Objective-C object that conforms to the BKHIDEventDeliveryObserverServiceProvider protocol.
type BKHIDEventDeliveryObserverServiceProviderObject struct {
	objectivec.Object
}

func (o BKHIDEventDeliveryObserverServiceProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// BKHIDEventDeliveryObserverServiceProviderObjectFromID constructs a [BKHIDEventDeliveryObserverServiceProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func BKHIDEventDeliveryObserverServiceProviderObjectFromID(id objc.ID) BKHIDEventDeliveryObserverServiceProviderObject {
	return BKHIDEventDeliveryObserverServiceProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/BKHIDEventDeliveryObserverServiceProvider/deliveryObserverServiceForAuditToken:
func (o BKHIDEventDeliveryObserverServiceProviderObject) DeliveryObserverServiceForAuditToken(token objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("deliveryObserverServiceForAuditToken:"), token)
	return objectivec.Object{ID: rv}
}

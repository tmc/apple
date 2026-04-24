// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// BKHIDEventDeliveryManagerProvider protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/BKHIDEventDeliveryManagerProvider
type BKHIDEventDeliveryManagerProvider interface {
	objectivec.IObject
}

// BKHIDEventDeliveryManagerProviderObject wraps an existing Objective-C object that conforms to the BKHIDEventDeliveryManagerProvider protocol.
type BKHIDEventDeliveryManagerProviderObject struct {
	objectivec.Object
}

func (o BKHIDEventDeliveryManagerProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// BKHIDEventDeliveryManagerProviderObjectFromID constructs a [BKHIDEventDeliveryManagerProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func BKHIDEventDeliveryManagerProviderObjectFromID(id objc.ID) BKHIDEventDeliveryManagerProviderObject {
	return BKHIDEventDeliveryManagerProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/BKHIDEventDeliveryManagerProvider/deliveryManagerForAuditToken:
func (o BKHIDEventDeliveryManagerProviderObject) DeliveryManagerForAuditToken(token objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("deliveryManagerForAuditToken:"), token)
	return objectivec.Object{ID: rv}
}

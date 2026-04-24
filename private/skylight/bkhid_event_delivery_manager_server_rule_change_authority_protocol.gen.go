// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// BKHIDEventDeliveryManagerServerRuleChangeAuthority protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/BKHIDEventDeliveryManagerServerRuleChangeAuthority
type BKHIDEventDeliveryManagerServerRuleChangeAuthority interface {
	objectivec.IObject
}

// BKHIDEventDeliveryManagerServerRuleChangeAuthorityObject wraps an existing Objective-C object that conforms to the BKHIDEventDeliveryManagerServerRuleChangeAuthority protocol.
type BKHIDEventDeliveryManagerServerRuleChangeAuthorityObject struct {
	objectivec.Object
}

func (o BKHIDEventDeliveryManagerServerRuleChangeAuthorityObject) BaseObject() objectivec.Object {
	return o.Object
}

// BKHIDEventDeliveryManagerServerRuleChangeAuthorityObjectFromID constructs a [BKHIDEventDeliveryManagerServerRuleChangeAuthorityObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func BKHIDEventDeliveryManagerServerRuleChangeAuthorityObjectFromID(id objc.ID) BKHIDEventDeliveryManagerServerRuleChangeAuthorityObject {
	return BKHIDEventDeliveryManagerServerRuleChangeAuthorityObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/BKHIDEventDeliveryManagerServerRuleChangeAuthority/permittedRuleChangeMaskForAuditToken:
func (o BKHIDEventDeliveryManagerServerRuleChangeAuthorityObject) PermittedRuleChangeMaskForAuditToken(token objectivec.IObject) uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("permittedRuleChangeMaskForAuditToken:"), token)
	return rv
}

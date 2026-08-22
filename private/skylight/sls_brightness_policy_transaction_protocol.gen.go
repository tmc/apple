// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLSBrightnessPolicyTransaction protocol.
type SLSBrightnessPolicyTransaction interface {
	objectivec.IObject

	// CommitBrightnessPolicy protocol.
	CommitBrightnessPolicy(policy []objectivec.IObject) bool

	// SetDimMessagingPolicy protocol.
	SetDimMessagingPolicy(policy byte)

	// SetShieldingPolicy protocol.
	SetShieldingPolicy(policy byte)

	// SetSleepMessagingPolicy protocol.
	SetSleepMessagingPolicy(policy byte)
}

// SLSBrightnessPolicyTransactionObject wraps an existing Objective-C object that conforms to the SLSBrightnessPolicyTransaction protocol.
type SLSBrightnessPolicyTransactionObject struct {
	objectivec.Object
}

func (o SLSBrightnessPolicyTransactionObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLSBrightnessPolicyTransactionObjectFromID constructs a [SLSBrightnessPolicyTransactionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLSBrightnessPolicyTransactionObjectFromID(id objc.ID) SLSBrightnessPolicyTransactionObject {
	return SLSBrightnessPolicyTransactionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o SLSBrightnessPolicyTransactionObject) CommitBrightnessPolicy(policy []objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("commitBrightnessPolicy:"), objectivec.IObjectSliceToNSArray(policy))
	return rv
}
func (o SLSBrightnessPolicyTransactionObject) SetDimMessagingPolicy(policy byte) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setDimMessagingPolicy:"), policy)
}
func (o SLSBrightnessPolicyTransactionObject) SetShieldingPolicy(policy byte) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setShieldingPolicy:"), policy)
}
func (o SLSBrightnessPolicyTransactionObject) SetSleepMessagingPolicy(policy byte) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setSleepMessagingPolicy:"), policy)
}

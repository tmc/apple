// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLSBrightnessPolicyTransaction protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessPolicyTransaction
type SLSBrightnessPolicyTransaction interface {
	objectivec.IObject

	// CommitBrightnessPolicy protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessPolicyTransaction/commitBrightnessPolicy:
	CommitBrightnessPolicy(policy []objectivec.IObject) bool

	// SetDimMessagingPolicy protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessPolicyTransaction/setDimMessagingPolicy:
	SetDimMessagingPolicy(policy byte)

	// SetShieldingPolicy protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessPolicyTransaction/setShieldingPolicy:
	SetShieldingPolicy(policy byte)

	// SetSleepMessagingPolicy protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessPolicyTransaction/setSleepMessagingPolicy:
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

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessPolicyTransaction/commitBrightnessPolicy:
func (o SLSBrightnessPolicyTransactionObject) CommitBrightnessPolicy(policy []objectivec.IObject) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("commitBrightnessPolicy:"), objectivec.IObjectSliceToNSArray(policy))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessPolicyTransaction/setDimMessagingPolicy:
func (o SLSBrightnessPolicyTransactionObject) SetDimMessagingPolicy(policy byte) {
	objc.Send[struct{}](o.ID, objc.Sel("setDimMessagingPolicy:"), policy)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessPolicyTransaction/setShieldingPolicy:
func (o SLSBrightnessPolicyTransactionObject) SetShieldingPolicy(policy byte) {
	objc.Send[struct{}](o.ID, objc.Sel("setShieldingPolicy:"), policy)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessPolicyTransaction/setSleepMessagingPolicy:
func (o SLSBrightnessPolicyTransactionObject) SetSleepMessagingPolicy(policy byte) {
	objc.Send[struct{}](o.ID, objc.Sel("setSleepMessagingPolicy:"), policy)
}

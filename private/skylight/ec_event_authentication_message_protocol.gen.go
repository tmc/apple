// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ECEventAuthenticationMessage protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/ECEventAuthenticationMessage
type ECEventAuthenticationMessage interface {
	objectivec.IObject

	// Capabilities protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/ECEventAuthenticationMessage/capabilities
	Capabilities() uint64
}

// ECEventAuthenticationMessageObject wraps an existing Objective-C object that conforms to the ECEventAuthenticationMessage protocol.
type ECEventAuthenticationMessageObject struct {
	objectivec.Object
}

func (o ECEventAuthenticationMessageObject) BaseObject() objectivec.Object {
	return o.Object
}

// ECEventAuthenticationMessageObjectFromID constructs a [ECEventAuthenticationMessageObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ECEventAuthenticationMessageObjectFromID(id objc.ID) ECEventAuthenticationMessageObject {
	return ECEventAuthenticationMessageObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/ECEventAuthenticationMessage/capabilities
func (o ECEventAuthenticationMessageObject) Capabilities() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("capabilities"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECEventAuthenticationMessage/eventType
func (o ECEventAuthenticationMessageObject) EventType() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("eventType"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/ECEventAuthenticationMessage/proxyTargetProcess
func (o ECEventAuthenticationMessageObject) ProxyTargetProcess() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("proxyTargetProcess"))
	return objectivec.Object{ID: rv}
}

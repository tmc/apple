// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ECEventAuthenticationMessage protocol.
type ECEventAuthenticationMessage interface {
	objectivec.IObject

	// Capabilities protocol.
	Capabilities() uint64

	// EventType protocol.
	EventType() objectivec.IObject

	// ProxyTargetProcess protocol.
	ProxyTargetProcess() objectivec.IObject
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

func (o ECEventAuthenticationMessageObject) Capabilities() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("capabilities"))
	return rv
}
func (o ECEventAuthenticationMessageObject) EventType() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("eventType"))
	return objectivec.Object{ID: rv}
}
func (o ECEventAuthenticationMessageObject) ProxyTargetProcess() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("proxyTargetProcess"))
	return objectivec.Object{ID: rv}
}

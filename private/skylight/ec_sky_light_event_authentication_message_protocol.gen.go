// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ECSkyLightEventAuthenticationMessage protocol.
type ECSkyLightEventAuthenticationMessage interface {
	objectivec.IObject

	// Connection protocol.
	Connection() uint32

	// Flags protocol.
	Flags() uint32

	// Location protocol.
	Location() corefoundation.CGPoint

	// MatchesEvent protocol.
	MatchesEvent(event unsafe.Pointer) bool

	// Window protocol.
	Window() uint32
}

// ECSkyLightEventAuthenticationMessageObject wraps an existing Objective-C object that conforms to the ECSkyLightEventAuthenticationMessage protocol.
type ECSkyLightEventAuthenticationMessageObject struct {
	objectivec.Object
}

func (o ECSkyLightEventAuthenticationMessageObject) BaseObject() objectivec.Object {
	return o.Object
}

// ECSkyLightEventAuthenticationMessageObjectFromID constructs a [ECSkyLightEventAuthenticationMessageObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ECSkyLightEventAuthenticationMessageObjectFromID(id objc.ID) ECSkyLightEventAuthenticationMessageObject {
	return ECSkyLightEventAuthenticationMessageObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o ECSkyLightEventAuthenticationMessageObject) Attributes() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("attributes"))
	return objectivec.Object{ID: rv}
}
func (o ECSkyLightEventAuthenticationMessageObject) Connection() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("connection"))
	return rv
}
func (o ECSkyLightEventAuthenticationMessageObject) Flags() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("flags"))
	return rv
}
func (o ECSkyLightEventAuthenticationMessageObject) Gesture() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("gesture"))
	return objectivec.Object{ID: rv}
}
func (o ECSkyLightEventAuthenticationMessageObject) Key() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("key"))
	return objectivec.Object{ID: rv}
}
func (o ECSkyLightEventAuthenticationMessageObject) Location() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("location"))
	return rv
}
func (o ECSkyLightEventAuthenticationMessageObject) MatchesEvent(event unsafe.Pointer) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("matchesEvent:"), event)
	return rv
}
func (o ECSkyLightEventAuthenticationMessageObject) Mouse() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("mouse"))
	return objectivec.Object{ID: rv}
}
func (o ECSkyLightEventAuthenticationMessageObject) Window() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("window"))
	return rv
}

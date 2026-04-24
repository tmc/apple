// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ECSkyLightEventAuthenticationMessage protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/ECSkyLightEventAuthenticationMessage
type ECSkyLightEventAuthenticationMessage interface {
	objectivec.IObject

	// Connection protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/ECSkyLightEventAuthenticationMessage/connection
	Connection() uint32

	// Flags protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/ECSkyLightEventAuthenticationMessage/flags
	Flags() uint32

	// Location protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/ECSkyLightEventAuthenticationMessage/location
	Location() corefoundation.CGPoint

	// MatchesEvent protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/ECSkyLightEventAuthenticationMessage/matchesEvent:
	MatchesEvent(event objectivec.IObject) bool

	// Window protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/ECSkyLightEventAuthenticationMessage/window
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

// See: https://developer.apple.com/documentation/SkyLight/ECSkyLightEventAuthenticationMessage/attributes
func (o ECSkyLightEventAuthenticationMessageObject) Attributes() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("attributes"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/ECSkyLightEventAuthenticationMessage/connection
func (o ECSkyLightEventAuthenticationMessageObject) Connection() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("connection"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECSkyLightEventAuthenticationMessage/flags
func (o ECSkyLightEventAuthenticationMessageObject) Flags() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("flags"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECSkyLightEventAuthenticationMessage/gesture
func (o ECSkyLightEventAuthenticationMessageObject) Gesture() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("gesture"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/ECSkyLightEventAuthenticationMessage/key
func (o ECSkyLightEventAuthenticationMessageObject) Key() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("key"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/ECSkyLightEventAuthenticationMessage/location
func (o ECSkyLightEventAuthenticationMessageObject) Location() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("location"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECSkyLightEventAuthenticationMessage/matchesEvent:
func (o ECSkyLightEventAuthenticationMessageObject) MatchesEvent(event objectivec.IObject) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("matchesEvent:"), event)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECSkyLightEventAuthenticationMessage/mouse
func (o ECSkyLightEventAuthenticationMessageObject) Mouse() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("mouse"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/ECSkyLightEventAuthenticationMessage/window
func (o ECSkyLightEventAuthenticationMessageObject) Window() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("window"))
	return rv
}

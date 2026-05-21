// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLSXPCServiceProtocol protocol.
type SLSXPCServiceProtocol interface {
	objectivec.IObject

	// Autoreconnect protocol.
	Autoreconnect() bool

	// ClientErrorBlock protocol.
	ClientErrorBlock() unsafe.Pointer

	// ClientNotificationBlock protocol.
	ClientNotificationBlock() unsafe.Pointer

	// Connected protocol.
	Connected() bool

	// Enabled protocol.
	Enabled() bool

	// ErrorBlock protocol.
	ErrorBlock() unsafe.Pointer

	// NotificationBlock protocol.
	NotificationBlock() unsafe.Pointer

	// SetAutoreconnect protocol.
	SetAutoreconnect(autoreconnect bool)

	// SetTarget protocol.
	SetTarget() bool
}

// SLSXPCServiceProtocolObject wraps an existing Objective-C object that conforms to the SLSXPCServiceProtocol protocol.
type SLSXPCServiceProtocolObject struct {
	objectivec.Object
}

func (o SLSXPCServiceProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLSXPCServiceProtocolObjectFromID constructs a [SLSXPCServiceProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLSXPCServiceProtocolObjectFromID(id objc.ID) SLSXPCServiceProtocolObject {
	return SLSXPCServiceProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o SLSXPCServiceProtocolObject) Autoreconnect() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("autoreconnect"))
	return rv
}
func (o SLSXPCServiceProtocolObject) ClientErrorBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("clientErrorBlock"))
	return rv
}
func (o SLSXPCServiceProtocolObject) ClientNotificationBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("clientNotificationBlock"))
	return rv
}
func (o SLSXPCServiceProtocolObject) Connected() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("connected"))
	return rv
}
func (o SLSXPCServiceProtocolObject) Connection() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("connection"))
	return objectivec.Object{ID: rv}
}
func (o SLSXPCServiceProtocolObject) CreateXPCDictionary(xPCDictionary uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("createXPCDictionary:"), xPCDictionary)
	return objectivec.Object{ID: rv}
}
func (o SLSXPCServiceProtocolObject) Enabled() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("enabled"))
	return rv
}
func (o SLSXPCServiceProtocolObject) ErrorBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("errorBlock"))
	return rv
}
func (o SLSXPCServiceProtocolObject) NotificationBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("notificationBlock"))
	return rv
}
func (o SLSXPCServiceProtocolObject) NotifyQueue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("notifyQueue"))
	return objectivec.Object{ID: rv}
}
func (o SLSXPCServiceProtocolObject) SendXPCDictionary(xPCDictionary objectivec.IObject) int {
	rv := objc.Send[int](o.ID, objc.Sel("sendXPCDictionary:"), xPCDictionary)
	return rv
}
func (o SLSXPCServiceProtocolObject) SendXPCDictionarySync(sync objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sendXPCDictionarySync:"), sync)
	return objectivec.Object{ID: rv}
}
func (o SLSXPCServiceProtocolObject) SetAutoreconnect(autoreconnect bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAutoreconnect:"), autoreconnect)
}
func (o SLSXPCServiceProtocolObject) SetTarget() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("setTarget"))
	return rv
}

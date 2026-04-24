// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLSXPCServiceProtocol protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol
type SLSXPCServiceProtocol interface {
	objectivec.IObject

	// Autoreconnect protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/autoreconnect
	Autoreconnect() bool

	// ClientErrorBlock protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/clientErrorBlock
	ClientErrorBlock() objectivec.IObject

	// ClientNotificationBlock protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/clientNotificationBlock
	ClientNotificationBlock() objectivec.IObject

	// Connected protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/connected
	Connected() bool

	// Enabled protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/enabled
	Enabled() bool

	// ErrorBlock protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/errorBlock
	ErrorBlock() objectivec.IObject

	// NotificationBlock protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/notificationBlock
	NotificationBlock() objectivec.IObject

	// SetAutoreconnect protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/setAutoreconnect:
	SetAutoreconnect(autoreconnect bool)

	// SetTarget protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/setTarget
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

// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/autoreconnect
func (o SLSXPCServiceProtocolObject) Autoreconnect() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("autoreconnect"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/clientErrorBlock
func (o SLSXPCServiceProtocolObject) ClientErrorBlock() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("clientErrorBlock"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/clientNotificationBlock
func (o SLSXPCServiceProtocolObject) ClientNotificationBlock() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("clientNotificationBlock"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/connected
func (o SLSXPCServiceProtocolObject) Connected() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("connected"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/connection
func (o SLSXPCServiceProtocolObject) Connection() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("connection"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/createXPCDictionary:
func (o SLSXPCServiceProtocolObject) CreateXPCDictionary(xPCDictionary uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("createXPCDictionary:"), xPCDictionary)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/enabled
func (o SLSXPCServiceProtocolObject) Enabled() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("enabled"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/errorBlock
func (o SLSXPCServiceProtocolObject) ErrorBlock() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("errorBlock"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/notificationBlock
func (o SLSXPCServiceProtocolObject) NotificationBlock() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("notificationBlock"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/notifyQueue
func (o SLSXPCServiceProtocolObject) NotifyQueue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("notifyQueue"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/sendXPCDictionary:
func (o SLSXPCServiceProtocolObject) SendXPCDictionary(xPCDictionary objectivec.IObject) int {
	rv := objc.Send[int](o.ID, objc.Sel("sendXPCDictionary:"), xPCDictionary)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/sendXPCDictionarySync:
func (o SLSXPCServiceProtocolObject) SendXPCDictionarySync(sync objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sendXPCDictionarySync:"), sync)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/setAutoreconnect:
func (o SLSXPCServiceProtocolObject) SetAutoreconnect(autoreconnect bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAutoreconnect:"), autoreconnect)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSXPCServiceProtocol/setTarget
func (o SLSXPCServiceProtocolObject) SetTarget() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("setTarget"))
	return rv
}

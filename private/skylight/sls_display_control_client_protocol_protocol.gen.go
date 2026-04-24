// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLSDisplayControlClientProtocol protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClientProtocol
type SLSDisplayControlClientProtocol interface {
	objectivec.IObject

	// Configured protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClientProtocol/configured
	Configured() bool

	// Enabled protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClientProtocol/enabled
	Enabled() bool

	// Notification protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClientProtocol/notification
	Notification() objectivec.IObject

	// TerminateConnection protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClientProtocol/terminateConnection
	TerminateConnection()
}

// SLSDisplayControlClientProtocolObject wraps an existing Objective-C object that conforms to the SLSDisplayControlClientProtocol protocol.
type SLSDisplayControlClientProtocolObject struct {
	objectivec.Object
}

func (o SLSDisplayControlClientProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLSDisplayControlClientProtocolObjectFromID constructs a [SLSDisplayControlClientProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLSDisplayControlClientProtocolObjectFromID(id objc.ID) SLSDisplayControlClientProtocolObject {
	return SLSDisplayControlClientProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClientProtocol/configured
func (o SLSDisplayControlClientProtocolObject) Configured() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("configured"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClientProtocol/enabled
func (o SLSDisplayControlClientProtocolObject) Enabled() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("enabled"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClientProtocol/notification
func (o SLSDisplayControlClientProtocolObject) Notification() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("notification"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClientProtocol/registerDaemonClient:withAutoreconnect:error:notifyQueue:notificationType:notificationBlock:
func (o SLSDisplayControlClientProtocolObject) RegisterDaemonClientWithAutoreconnectErrorNotifyQueueNotificationTypeNotificationBlock(client objectivec.IObject, autoreconnect bool, error_ []objectivec.IObject, queue objectivec.IObject, type_ uint64, block VoidHandler) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("registerDaemonClient:withAutoreconnect:error:notifyQueue:notificationType:notificationBlock:"), client, autoreconnect, objectivec.IObjectSliceToNSArray(error_), queue, type_, block)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClientProtocol/registerGUIClient:connectionPort:error:notifyQueue:notificationType:notificationBlock:
func (o SLSDisplayControlClientProtocolObject) RegisterGUIClientConnectionPortErrorNotifyQueueNotificationTypeNotificationBlock(gUIClient objectivec.IObject, port uint32, error_ []objectivec.IObject, queue objectivec.IObject, type_ uint64, block VoidHandler) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("registerGUIClient:connectionPort:error:notifyQueue:notificationType:notificationBlock:"), gUIClient, port, objectivec.IObjectSliceToNSArray(error_), queue, type_, block)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClientProtocol/terminateConnection
func (o SLSDisplayControlClientProtocolObject) TerminateConnection() {
	objc.Send[struct{}](o.ID, objc.Sel("terminateConnection"))
}

// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLSDisplayControlClientProtocol protocol.
type SLSDisplayControlClientProtocol interface {
	objectivec.IObject

	// Configured protocol.
	Configured() bool

	// Enabled protocol.
	Enabled() bool

	// Notification protocol.
	Notification() unsafe.Pointer

	// TerminateConnection protocol.
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

func (o SLSDisplayControlClientProtocolObject) Configured() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("configured"))
	return rv
}
func (o SLSDisplayControlClientProtocolObject) Enabled() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("enabled"))
	return rv
}
func (o SLSDisplayControlClientProtocolObject) Notification() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("notification"))
	return rv
}
func (o SLSDisplayControlClientProtocolObject) RegisterDaemonClientWithAutoreconnectErrorNotifyQueueNotificationTypeNotificationBlock(client objectivec.IObject, autoreconnect bool, error_ []objectivec.IObject, queue objectivec.IObject, type_ uint64, block UnsafePointerHandler) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("registerDaemonClient:withAutoreconnect:error:notifyQueue:notificationType:notificationBlock:"), client, autoreconnect, objectivec.IObjectSliceToNSArray(error_), queue, type_, block)
	return objectivec.Object{ID: rv}
}
func (o SLSDisplayControlClientProtocolObject) RegisterGUIClientConnectionPortErrorNotifyQueueNotificationTypeNotificationBlock(gUIClient objectivec.IObject, port uint32, error_ []objectivec.IObject, queue objectivec.IObject, type_ uint64, block UnsafePointerHandler) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("registerGUIClient:connectionPort:error:notifyQueue:notificationType:notificationBlock:"), gUIClient, port, objectivec.IObjectSliceToNSArray(error_), queue, type_, block)
	return objectivec.Object{ID: rv}
}
func (o SLSDisplayControlClientProtocolObject) TerminateConnection() {
	objc.Send[struct{}](o.ID, objc.Sel("terminateConnection"))
}

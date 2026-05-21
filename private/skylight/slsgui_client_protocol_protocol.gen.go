// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLSGUIClientProtocol protocol.
type SLSGUIClientProtocol interface {
	objectivec.IObject

	// TerminateConnection protocol.
	TerminateConnection(connection []objectivec.IObject)
}

// SLSGUIClientProtocolObject wraps an existing Objective-C object that conforms to the SLSGUIClientProtocol protocol.
type SLSGUIClientProtocolObject struct {
	objectivec.Object
}

func (o SLSGUIClientProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLSGUIClientProtocolObjectFromID constructs a [SLSGUIClientProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLSGUIClientProtocolObjectFromID(id objc.ID) SLSGUIClientProtocolObject {
	return SLSGUIClientProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o SLSGUIClientProtocolObject) ConfigGUIClientErrorNotifyQueueNotificationTypeNotificationBlock(gUIClient objectivec.IObject, error_ []objectivec.IObject, queue objectivec.IObject, type_ uint64, block UnsafePointerHandler) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("configGUIClient:error:notifyQueue:notificationType:notificationBlock:"), gUIClient, objectivec.IObjectSliceToNSArray(error_), queue, type_, block)
	return objectivec.Object{ID: rv}
}
func (o SLSGUIClientProtocolObject) RequestDisplaysIdleError(idle objectivec.IObject) (uint64, error) {
	rv, err := objc.SendWithError[uint64](o.ID, objc.Sel("requestDisplaysIdle:error:"), idle)
	if err != nil {
		return 0, err
	}
	return rv, nil
}
func (o SLSGUIClientProtocolObject) TerminateConnection(connection []objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("terminateConnection:"), objectivec.IObjectSliceToNSArray(connection))
}

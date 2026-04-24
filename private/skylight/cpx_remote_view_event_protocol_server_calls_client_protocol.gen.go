// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXRemoteViewEventProtocolServerCallsClient protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventProtocolServerCallsClient
type CPXRemoteViewEventProtocolServerCallsClient interface {
	objectivec.IObject
}

// CPXRemoteViewEventProtocolServerCallsClientObject wraps an existing Objective-C object that conforms to the CPXRemoteViewEventProtocolServerCallsClient protocol.
type CPXRemoteViewEventProtocolServerCallsClientObject struct {
	objectivec.Object
}

func (o CPXRemoteViewEventProtocolServerCallsClientObject) BaseObject() objectivec.Object {
	return o.Object
}

// CPXRemoteViewEventProtocolServerCallsClientObjectFromID constructs a [CPXRemoteViewEventProtocolServerCallsClientObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CPXRemoteViewEventProtocolServerCallsClientObjectFromID(id objc.ID) CPXRemoteViewEventProtocolServerCallsClientObject {
	return CPXRemoteViewEventProtocolServerCallsClientObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventProtocolServerCallsClient/sendEventToHost:fullDispatch:reply:
func (o CPXRemoteViewEventProtocolServerCallsClientObject) SendEventToHostFullDispatchReply(host objectivec.IObject, dispatch objectivec.IObject, reply NumberErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("sendEventToHost:fullDispatch:reply:"), host, dispatch, reply)
}

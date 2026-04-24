// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXRemoteViewEventServerConfig protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventServerConfig
type CPXRemoteViewEventServerConfig interface {
	objectivec.IObject
}

// CPXRemoteViewEventServerConfigObject wraps an existing Objective-C object that conforms to the CPXRemoteViewEventServerConfig protocol.
type CPXRemoteViewEventServerConfigObject struct {
	objectivec.Object
}

func (o CPXRemoteViewEventServerConfigObject) BaseObject() objectivec.Object {
	return o.Object
}

// CPXRemoteViewEventServerConfigObjectFromID constructs a [CPXRemoteViewEventServerConfigObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CPXRemoteViewEventServerConfigObjectFromID(id objc.ID) CPXRemoteViewEventServerConfigObject {
	return CPXRemoteViewEventServerConfigObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventServerConfig/listenerDomain
func (o CPXRemoteViewEventServerConfigObject) ListenerDomain() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("listenerDomain"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventServerConfig/listenerService
func (o CPXRemoteViewEventServerConfigObject) ListenerService() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("listenerService"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventServerConfig/remoteViewEventManagerForConnection:
func (o CPXRemoteViewEventServerConfigObject) RemoteViewEventManagerForConnection(connection objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("remoteViewEventManagerForConnection:"), connection)
	return objectivec.Object{ID: rv}
}

// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLSRemoteViewEventClientConfig protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/SLSRemoteViewEventClientConfig
type SLSRemoteViewEventClientConfig interface {
	objectivec.IObject
}

// SLSRemoteViewEventClientConfigObject wraps an existing Objective-C object that conforms to the SLSRemoteViewEventClientConfig protocol.
type SLSRemoteViewEventClientConfigObject struct {
	objectivec.Object
}

func (o SLSRemoteViewEventClientConfigObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLSRemoteViewEventClientConfigObjectFromID constructs a [SLSRemoteViewEventClientConfigObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLSRemoteViewEventClientConfigObjectFromID(id objc.ID) SLSRemoteViewEventClientConfigObject {
	return SLSRemoteViewEventClientConfigObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSRemoteViewEventClientConfig/connection
func (o SLSRemoteViewEventClientConfigObject) Connection() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("connection"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSRemoteViewEventClientConfig/serviceInterface
func (o SLSRemoteViewEventClientConfigObject) ServiceInterface() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("serviceInterface"))
	return objectivec.Object{ID: rv}
}

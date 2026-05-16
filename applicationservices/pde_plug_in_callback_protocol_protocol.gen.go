// Code generated from Apple documentation for ApplicationServices. DO NOT EDIT.

package applicationservices

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// PDEPlugInCallbackProtocol protocol.
//
// See: https://developer.apple.com/documentation/applicationservices/pdeplugincallbackprotocol
type PDEPlugInCallbackProtocol interface {
	objectivec.IObject
}

// PDEPlugInCallbackProtocolObject wraps an existing Objective-C object that conforms to the PDEPlugInCallbackProtocol protocol.
type PDEPlugInCallbackProtocolObject struct {
	objectivec.Object
}

func (o PDEPlugInCallbackProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// PDEPlugInCallbackProtocolObjectFromID constructs a [PDEPlugInCallbackProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func PDEPlugInCallbackProtocolObjectFromID(id objc.ID) PDEPlugInCallbackProtocolObject {
	return PDEPlugInCallbackProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Code generated from Apple documentation for ApplicationServices. DO NOT EDIT.

package applicationservices

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// PDEPlugIn protocol.
//
// See: https://developer.apple.com/documentation/applicationservices/pdeplugin
type PDEPlugIn interface {
	objectivec.IObject
}

// PDEPlugInObject wraps an existing Objective-C object that conforms to the PDEPlugIn protocol.
type PDEPlugInObject struct {
	objectivec.Object
}

func (o PDEPlugInObject) BaseObject() objectivec.Object {
	return o.Object
}

// PDEPlugInObjectFromID constructs a [PDEPlugInObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func PDEPlugInObjectFromID(id objc.ID) PDEPlugInObject {
	return PDEPlugInObject{
		Object: objectivec.ObjectFromID(id),
	}
}

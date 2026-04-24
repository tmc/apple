// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLXPCExportable protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/SLXPCExportable
type SLXPCExportable interface {
	objectivec.IObject
}

// SLXPCExportableObject wraps an existing Objective-C object that conforms to the SLXPCExportable protocol.
type SLXPCExportableObject struct {
	objectivec.Object
}

func (o SLXPCExportableObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLXPCExportableObjectFromID constructs a [SLXPCExportableObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLXPCExportableObjectFromID(id objc.ID) SLXPCExportableObject {
	return SLXPCExportableObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/SLXPCExportable/createXPCObject
func (o SLXPCExportableObject) CreateXPCObject() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("createXPCObject"))
	return objectivec.Object{ID: rv}
}

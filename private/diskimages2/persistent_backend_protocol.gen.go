// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// PersistentBackend protocol.
type PersistentBackend interface {
	objectivec.IObject

	// SandboxExtensionTokenWithError protocol.
	SandboxExtensionTokenWithError() (objectivec.IObject, error)
}

// PersistentBackendObject wraps an existing Objective-C object that conforms to the PersistentBackend protocol.
type PersistentBackendObject struct {
	objectivec.Object
}

func (o PersistentBackendObject) BaseObject() objectivec.Object {
	return o.Object
}

// PersistentBackendObjectFromID constructs a [PersistentBackendObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func PersistentBackendObjectFromID(id objc.ID) PersistentBackendObject {
	return PersistentBackendObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o PersistentBackendObject) SandboxExtensionTokenWithError() (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("sandboxExtensionTokenWithError:"))
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}

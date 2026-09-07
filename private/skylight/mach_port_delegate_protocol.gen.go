// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSMachPortDelegate protocol.
type NSMachPortDelegate interface {
	objectivec.IObject

	// HandleMachMessage protocol.
	HandleMachMessage(message unsafe.Pointer)
}

// NSMachPortDelegateObject wraps an existing Objective-C object that conforms to the NSMachPortDelegate protocol.
type NSMachPortDelegateObject struct {
	objectivec.Object
}

func (o NSMachPortDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSMachPortDelegateObjectFromID constructs a [NSMachPortDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSMachPortDelegateObjectFromID(id objc.ID) NSMachPortDelegateObject {
	return NSMachPortDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o NSMachPortDelegateObject) HandleMachMessage(message unsafe.Pointer) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("handleMachMessage:"), message)
}

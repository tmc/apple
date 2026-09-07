// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZVirtualMachineAccessorObserver protocol.
type VZVirtualMachineAccessorObserver interface {
	objectivec.IObject

	// VirtualMachineAccessorAssociateWithDisplayPresenter protocol.
	VirtualMachineAccessorAssociateWithDisplayPresenter(accessor objectivec.IObject, presenter objectivec.IObject)
}

// VZVirtualMachineAccessorObserverObject wraps an existing Objective-C object that conforms to the VZVirtualMachineAccessorObserver protocol.
type VZVirtualMachineAccessorObserverObject struct {
	objectivec.Object
}

func (o VZVirtualMachineAccessorObserverObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZVirtualMachineAccessorObserverObjectFromID constructs a [VZVirtualMachineAccessorObserverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZVirtualMachineAccessorObserverObjectFromID(id objc.ID) VZVirtualMachineAccessorObserverObject {
	return VZVirtualMachineAccessorObserverObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o VZVirtualMachineAccessorObserverObject) VirtualMachineAccessorAssociateWithDisplayPresenter(accessor objectivec.IObject, presenter objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("virtualMachineAccessor:associateWithDisplayPresenter:"), accessor, presenter)
}

// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtualMachineAccessorManager] class.
var (
	_VZVirtualMachineAccessorManagerClass     VZVirtualMachineAccessorManagerClass
	_VZVirtualMachineAccessorManagerClassOnce sync.Once
)

func getVZVirtualMachineAccessorManagerClass() VZVirtualMachineAccessorManagerClass {
	_VZVirtualMachineAccessorManagerClassOnce.Do(func() {
		_VZVirtualMachineAccessorManagerClass = VZVirtualMachineAccessorManagerClass{class: objc.GetClass("_VZVirtualMachineAccessorManager")}
	})
	return _VZVirtualMachineAccessorManagerClass
}

// GetVZVirtualMachineAccessorManagerClass returns the class object for _VZVirtualMachineAccessorManager.
func GetVZVirtualMachineAccessorManagerClass() VZVirtualMachineAccessorManagerClass {
	return getVZVirtualMachineAccessorManagerClass()
}

type VZVirtualMachineAccessorManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtualMachineAccessorManagerClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtualMachineAccessorManagerClass) Alloc() VZVirtualMachineAccessorManager {
	rv := objc.Send[VZVirtualMachineAccessorManager](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineAccessorManager
type VZVirtualMachineAccessorManager struct {
	objectivec.Object
}

// VZVirtualMachineAccessorManagerFromID constructs a [VZVirtualMachineAccessorManager] from an objc.ID.
func VZVirtualMachineAccessorManagerFromID(id objc.ID) VZVirtualMachineAccessorManager {
	return VZVirtualMachineAccessorManager{objectivec.Object{ID: id}}
}
// Ensure VZVirtualMachineAccessorManager implements IVZVirtualMachineAccessorManager.
var _ IVZVirtualMachineAccessorManager = VZVirtualMachineAccessorManager{}

// An interface definition for the [VZVirtualMachineAccessorManager] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineAccessorManager
type IVZVirtualMachineAccessorManager interface {
	objectivec.IObject
}

// Init initializes the instance.
func (v VZVirtualMachineAccessorManager) Init() VZVirtualMachineAccessorManager {
	rv := objc.Send[VZVirtualMachineAccessorManager](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtualMachineAccessorManager) Autorelease() VZVirtualMachineAccessorManager {
	rv := objc.Send[VZVirtualMachineAccessorManager](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachineAccessorManager creates a new VZVirtualMachineAccessorManager instance.
func NewVZVirtualMachineAccessorManager() VZVirtualMachineAccessorManager {
	class := getVZVirtualMachineAccessorManagerClass()
	rv := objc.Send[VZVirtualMachineAccessorManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}


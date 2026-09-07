// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtualMachineViewAdaptorInternal] class.
var (
	_VZVirtualMachineViewAdaptorInternalClass     VZVirtualMachineViewAdaptorInternalClass
	_VZVirtualMachineViewAdaptorInternalClassOnce sync.Once
)

func getVZVirtualMachineViewAdaptorInternalClass() VZVirtualMachineViewAdaptorInternalClass {
	_VZVirtualMachineViewAdaptorInternalClassOnce.Do(func() {
		_VZVirtualMachineViewAdaptorInternalClass = VZVirtualMachineViewAdaptorInternalClass{class: objc.GetClass("_VZVirtualMachineViewAdaptorInternal")}
	})
	return _VZVirtualMachineViewAdaptorInternalClass
}

// GetVZVirtualMachineViewAdaptorInternalClass returns the class object for _VZVirtualMachineViewAdaptorInternal.
func GetVZVirtualMachineViewAdaptorInternalClass() VZVirtualMachineViewAdaptorInternalClass {
	return getVZVirtualMachineViewAdaptorInternalClass()
}

type VZVirtualMachineViewAdaptorInternalClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtualMachineViewAdaptorInternalClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtualMachineViewAdaptorInternalClass) Alloc() VZVirtualMachineViewAdaptorInternal {
	rv := objc.SendIfResponds[VZVirtualMachineViewAdaptorInternal](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVirtualMachineViewAdaptorInternal.InitWithGraphicsDisplay]
//   - [VZVirtualMachineViewAdaptorInternal.InitWithVirtualMachine]
type VZVirtualMachineViewAdaptorInternal struct {
	objectivec.Object
}

// VZVirtualMachineViewAdaptorInternalFromID constructs a [VZVirtualMachineViewAdaptorInternal] from an objc.ID.
func VZVirtualMachineViewAdaptorInternalFromID(id objc.ID) VZVirtualMachineViewAdaptorInternal {
	return VZVirtualMachineViewAdaptorInternal{objectivec.Object{ID: id}}
}

// Ensure VZVirtualMachineViewAdaptorInternal implements IVZVirtualMachineViewAdaptorInternal.
var _ IVZVirtualMachineViewAdaptorInternal = VZVirtualMachineViewAdaptorInternal{}

// An interface definition for the [VZVirtualMachineViewAdaptorInternal] class.
//
// # Methods
//
//   - [IVZVirtualMachineViewAdaptorInternal.InitWithGraphicsDisplay]
//   - [IVZVirtualMachineViewAdaptorInternal.InitWithVirtualMachine]
type IVZVirtualMachineViewAdaptorInternal interface {
	objectivec.IObject

	// Topic: Methods

	InitWithGraphicsDisplay(display objectivec.IObject) VZVirtualMachineViewAdaptorInternal
	InitWithVirtualMachine(machine objectivec.IObject) VZVirtualMachineViewAdaptorInternal
}

// Init initializes the instance.
func (v VZVirtualMachineViewAdaptorInternal) Init() VZVirtualMachineViewAdaptorInternal {
	rv := objc.SendIfResponds[VZVirtualMachineViewAdaptorInternal](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtualMachineViewAdaptorInternal) Autorelease() VZVirtualMachineViewAdaptorInternal {
	rv := objc.SendIfResponds[VZVirtualMachineViewAdaptorInternal](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachineViewAdaptorInternal creates a new VZVirtualMachineViewAdaptorInternal instance.
func NewVZVirtualMachineViewAdaptorInternal() VZVirtualMachineViewAdaptorInternal {
	class := getVZVirtualMachineViewAdaptorInternalClass()
	rv := objc.SendIfResponds[VZVirtualMachineViewAdaptorInternal](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZVirtualMachineViewAdaptorInternalWithGraphicsDisplay(display objectivec.IObject) VZVirtualMachineViewAdaptorInternal {
	instance := getVZVirtualMachineViewAdaptorInternalClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithGraphicsDisplay:"), display)
	return VZVirtualMachineViewAdaptorInternalFromID(rv)
}

func NewVZVirtualMachineViewAdaptorInternalWithVirtualMachine(machine objectivec.IObject) VZVirtualMachineViewAdaptorInternal {
	instance := getVZVirtualMachineViewAdaptorInternalClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithVirtualMachine:"), machine)
	return VZVirtualMachineViewAdaptorInternalFromID(rv)
}

func (v VZVirtualMachineViewAdaptorInternal) InitWithGraphicsDisplay(display objectivec.IObject) VZVirtualMachineViewAdaptorInternal {
	rv := objc.SendIfResponds[VZVirtualMachineViewAdaptorInternal](v.ID, objc.Sel("initWithGraphicsDisplay:"), display)
	return rv
}
func (v VZVirtualMachineViewAdaptorInternal) InitWithVirtualMachine(machine objectivec.IObject) VZVirtualMachineViewAdaptorInternal {
	rv := objc.SendIfResponds[VZVirtualMachineViewAdaptorInternal](v.ID, objc.Sel("initWithVirtualMachine:"), machine)
	return rv
}

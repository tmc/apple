// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtualMachineStartOptions] class.
var (
	_VZVirtualMachineStartOptionsClass     VZVirtualMachineStartOptionsClass
	_VZVirtualMachineStartOptionsClassOnce sync.Once
)

func getVZVirtualMachineStartOptionsClass() VZVirtualMachineStartOptionsClass {
	_VZVirtualMachineStartOptionsClassOnce.Do(func() {
		_VZVirtualMachineStartOptionsClass = VZVirtualMachineStartOptionsClass{class: objc.GetClass("VZVirtualMachineStartOptions")}
	})
	return _VZVirtualMachineStartOptionsClass
}

// GetVZVirtualMachineStartOptionsClass returns the class object for VZVirtualMachineStartOptions.
func GetVZVirtualMachineStartOptionsClass() VZVirtualMachineStartOptionsClass {
	return getVZVirtualMachineStartOptionsClass()
}

type VZVirtualMachineStartOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtualMachineStartOptionsClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtualMachineStartOptionsClass) Alloc() VZVirtualMachineStartOptions {
	rv := objc.Send[VZVirtualMachineStartOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The abstract class for VM start options.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineStartOptions
type VZVirtualMachineStartOptions struct {
	objectivec.Object
}

// VZVirtualMachineStartOptionsFromID constructs a [VZVirtualMachineStartOptions] from an objc.ID.
//
// The abstract class for VM start options.
func VZVirtualMachineStartOptionsFromID(id objc.ID) VZVirtualMachineStartOptions {
	return VZVirtualMachineStartOptions{objectivec.Object{ID: id}}
}
// NOTE: VZVirtualMachineStartOptions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtualMachineStartOptions] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineStartOptions
type IVZVirtualMachineStartOptions interface {
	objectivec.IObject
}

// Init initializes the instance.
func (v VZVirtualMachineStartOptions) Init() VZVirtualMachineStartOptions {
	rv := objc.Send[VZVirtualMachineStartOptions](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtualMachineStartOptions) Autorelease() VZVirtualMachineStartOptions {
	rv := objc.Send[VZVirtualMachineStartOptions](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachineStartOptions creates a new VZVirtualMachineStartOptions instance.
func NewVZVirtualMachineStartOptions() VZVirtualMachineStartOptions {
	class := getVZVirtualMachineStartOptionsClass()
	rv := objc.Send[VZVirtualMachineStartOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}


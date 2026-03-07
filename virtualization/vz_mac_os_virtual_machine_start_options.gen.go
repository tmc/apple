// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacOSVirtualMachineStartOptions] class.
var (
	_VZMacOSVirtualMachineStartOptionsClass     VZMacOSVirtualMachineStartOptionsClass
	_VZMacOSVirtualMachineStartOptionsClassOnce sync.Once
)

func getVZMacOSVirtualMachineStartOptionsClass() VZMacOSVirtualMachineStartOptionsClass {
	_VZMacOSVirtualMachineStartOptionsClassOnce.Do(func() {
		_VZMacOSVirtualMachineStartOptionsClass = VZMacOSVirtualMachineStartOptionsClass{class: objc.GetClass("VZMacOSVirtualMachineStartOptions")}
	})
	return _VZMacOSVirtualMachineStartOptionsClass
}

// GetVZMacOSVirtualMachineStartOptionsClass returns the class object for VZMacOSVirtualMachineStartOptions.
func GetVZMacOSVirtualMachineStartOptionsClass() VZMacOSVirtualMachineStartOptionsClass {
	return getVZMacOSVirtualMachineStartOptionsClass()
}

type VZMacOSVirtualMachineStartOptionsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacOSVirtualMachineStartOptionsClass) Alloc() VZMacOSVirtualMachineStartOptions {
	rv := objc.Send[VZMacOSVirtualMachineStartOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// A class that describes start options for macOS VMs.
//
// # Setting recovery mode
//
//   - [VZMacOSVirtualMachineStartOptions.StartUpFromMacOSRecovery]: A Boolean value that indicates whether the macOS guest should start in recovery mode.
//   - [VZMacOSVirtualMachineStartOptions.SetStartUpFromMacOSRecovery]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions
type VZMacOSVirtualMachineStartOptions struct {
	VZVirtualMachineStartOptions
}

// VZMacOSVirtualMachineStartOptionsFromID constructs a [VZMacOSVirtualMachineStartOptions] from an objc.ID.
//
// A class that describes start options for macOS VMs.
func VZMacOSVirtualMachineStartOptionsFromID(id objc.ID) VZMacOSVirtualMachineStartOptions {
	return VZMacOSVirtualMachineStartOptions{VZVirtualMachineStartOptions: VZVirtualMachineStartOptionsFromID(id)}
}
// NOTE: VZMacOSVirtualMachineStartOptions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZMacOSVirtualMachineStartOptions] class.
//
// # Setting recovery mode
//
//   - [IVZMacOSVirtualMachineStartOptions.StartUpFromMacOSRecovery]: A Boolean value that indicates whether the macOS guest should start in recovery mode.
//   - [IVZMacOSVirtualMachineStartOptions.SetStartUpFromMacOSRecovery]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions
type IVZMacOSVirtualMachineStartOptions interface {
	IVZVirtualMachineStartOptions

	// Topic: Setting recovery mode

	// A Boolean value that indicates whether the macOS guest should start in recovery mode.
	StartUpFromMacOSRecovery() bool
	SetStartUpFromMacOSRecovery(value bool)
}





// Init initializes the instance.
func (m VZMacOSVirtualMachineStartOptions) Init() VZMacOSVirtualMachineStartOptions {
	rv := objc.Send[VZMacOSVirtualMachineStartOptions](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacOSVirtualMachineStartOptions) Autorelease() VZMacOSVirtualMachineStartOptions {
	rv := objc.Send[VZMacOSVirtualMachineStartOptions](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacOSVirtualMachineStartOptions creates a new VZMacOSVirtualMachineStartOptions instance.
func NewVZMacOSVirtualMachineStartOptions() VZMacOSVirtualMachineStartOptions {
	class := getVZMacOSVirtualMachineStartOptionsClass()
	rv := objc.Send[VZMacOSVirtualMachineStartOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// A Boolean value that indicates whether the macOS guest should start in
// recovery mode.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions/startUpFromMacOSRecovery
func (m VZMacOSVirtualMachineStartOptions) StartUpFromMacOSRecovery() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("startUpFromMacOSRecovery"))
	return rv
}
func (m VZMacOSVirtualMachineStartOptions) SetStartUpFromMacOSRecovery(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setStartUpFromMacOSRecovery:"), value)
}
























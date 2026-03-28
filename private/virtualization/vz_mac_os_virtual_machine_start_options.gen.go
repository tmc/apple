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

// Class returns the underlying Objective-C class pointer.
func (vc VZMacOSVirtualMachineStartOptionsClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacOSVirtualMachineStartOptionsClass) Alloc() VZMacOSVirtualMachineStartOptions {
	rv := objc.Send[VZMacOSVirtualMachineStartOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZMacOSVirtualMachineStartOptions._forceDFU]
//   - [VZMacOSVirtualMachineStartOptions.Set_forceDFU]
//   - [VZMacOSVirtualMachineStartOptions._setForceDFU]
//   - [VZMacOSVirtualMachineStartOptions._setStopInIBootStage1]
//   - [VZMacOSVirtualMachineStartOptions._setStopInIBootStage2]
//   - [VZMacOSVirtualMachineStartOptions._stopInIBootStage1]
//   - [VZMacOSVirtualMachineStartOptions.Set_stopInIBootStage1]
//   - [VZMacOSVirtualMachineStartOptions._stopInIBootStage2]
//   - [VZMacOSVirtualMachineStartOptions.Set_stopInIBootStage2]
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions
type VZMacOSVirtualMachineStartOptions struct {
	VZVirtualMachineStartOptions
}

// VZMacOSVirtualMachineStartOptionsFromID constructs a [VZMacOSVirtualMachineStartOptions] from an objc.ID.
func VZMacOSVirtualMachineStartOptionsFromID(id objc.ID) VZMacOSVirtualMachineStartOptions {
	return VZMacOSVirtualMachineStartOptions{VZVirtualMachineStartOptions: VZVirtualMachineStartOptionsFromID(id)}
}
// Ensure VZMacOSVirtualMachineStartOptions implements IVZMacOSVirtualMachineStartOptions.
var _ IVZMacOSVirtualMachineStartOptions = VZMacOSVirtualMachineStartOptions{}

// An interface definition for the [VZMacOSVirtualMachineStartOptions] class.
//
// # Methods
//
//   - [IVZMacOSVirtualMachineStartOptions._forceDFU]
//   - [IVZMacOSVirtualMachineStartOptions.Set_forceDFU]
//   - [IVZMacOSVirtualMachineStartOptions._setForceDFU]
//   - [IVZMacOSVirtualMachineStartOptions._setStopInIBootStage1]
//   - [IVZMacOSVirtualMachineStartOptions._setStopInIBootStage2]
//   - [IVZMacOSVirtualMachineStartOptions._stopInIBootStage1]
//   - [IVZMacOSVirtualMachineStartOptions.Set_stopInIBootStage1]
//   - [IVZMacOSVirtualMachineStartOptions._stopInIBootStage2]
//   - [IVZMacOSVirtualMachineStartOptions.Set_stopInIBootStage2]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions
type IVZMacOSVirtualMachineStartOptions interface {
	IVZVirtualMachineStartOptions

	// Topic: Methods

	_forceDFU() bool
	Set_forceDFU(value bool)
	_setForceDFU(dfu bool)
	_setStopInIBootStage1(stage1 bool)
	_setStopInIBootStage2(stage2 bool)
	_stopInIBootStage1() bool
	Set_stopInIBootStage1(value bool)
	_stopInIBootStage2() bool
	Set_stopInIBootStage2(value bool)
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

//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions/_setForceDFU:
func (m VZMacOSVirtualMachineStartOptions) _setForceDFU(dfu bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setForceDFU:"), dfu)
}

// SetForceDFU is an exported wrapper for the private method _setForceDFU.
func (m VZMacOSVirtualMachineStartOptions) SetForceDFU(dfu bool) {
	m._setForceDFU(dfu)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions/_setStopInIBootStage1:
func (m VZMacOSVirtualMachineStartOptions) _setStopInIBootStage1(stage1 bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setStopInIBootStage1:"), stage1)
}

// SetStopInIBootStage1 is an exported wrapper for the private method _setStopInIBootStage1.
func (m VZMacOSVirtualMachineStartOptions) SetStopInIBootStage1(stage1 bool) {
	m._setStopInIBootStage1(stage1)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions/_setStopInIBootStage2:
func (m VZMacOSVirtualMachineStartOptions) _setStopInIBootStage2(stage2 bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setStopInIBootStage2:"), stage2)
}

// SetStopInIBootStage2 is an exported wrapper for the private method _setStopInIBootStage2.
func (m VZMacOSVirtualMachineStartOptions) SetStopInIBootStage2(stage2 bool) {
	m._setStopInIBootStage2(stage2)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions/_forceDFU
func (m VZMacOSVirtualMachineStartOptions) _forceDFU() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_forceDFU"))
	return rv
}
func (m VZMacOSVirtualMachineStartOptions) Set_forceDFU(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("set_forceDFU:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions/_stopInIBootStage1
func (m VZMacOSVirtualMachineStartOptions) _stopInIBootStage1() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_stopInIBootStage1"))
	return rv
}
func (m VZMacOSVirtualMachineStartOptions) Set_stopInIBootStage1(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("set_stopInIBootStage1:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions/_stopInIBootStage2
func (m VZMacOSVirtualMachineStartOptions) _stopInIBootStage2() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_stopInIBootStage2"))
	return rv
}
func (m VZMacOSVirtualMachineStartOptions) Set_stopInIBootStage2(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("set_stopInIBootStage2:"), value)
}


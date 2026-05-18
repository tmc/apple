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
//
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
func (v VZMacOSVirtualMachineStartOptions) Init() VZMacOSVirtualMachineStartOptions {
	rv := objc.Send[VZMacOSVirtualMachineStartOptions](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacOSVirtualMachineStartOptions) Autorelease() VZMacOSVirtualMachineStartOptions {
	rv := objc.Send[VZMacOSVirtualMachineStartOptions](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacOSVirtualMachineStartOptions creates a new VZMacOSVirtualMachineStartOptions instance.
func NewVZMacOSVirtualMachineStartOptions() VZMacOSVirtualMachineStartOptions {
	class := getVZMacOSVirtualMachineStartOptionsClass()
	rv := objc.Send[VZMacOSVirtualMachineStartOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions/_setForceDFU:
func (v VZMacOSVirtualMachineStartOptions) _setForceDFU(dfu bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setForceDFU:"), dfu)
}

// SetForceDFU is an exported wrapper for the private method _setForceDFU.
func (v VZMacOSVirtualMachineStartOptions) SetForceDFU(dfu bool) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setForceDFU:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setForceDFU:"}
		return err
	}
	v._setForceDFU(dfu)
	return nil
}

// CanSetForceDFU reports whether the receiver responds to the private selector _setForceDFU:.
func (v VZMacOSVirtualMachineStartOptions) CanSetForceDFU() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setForceDFU:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions/_setStopInIBootStage1:
func (v VZMacOSVirtualMachineStartOptions) _setStopInIBootStage1(stage1 bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setStopInIBootStage1:"), stage1)
}

// SetStopInIBootStage1 is an exported wrapper for the private method _setStopInIBootStage1.
func (v VZMacOSVirtualMachineStartOptions) SetStopInIBootStage1(stage1 bool) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setStopInIBootStage1:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setStopInIBootStage1:"}
		return err
	}
	v._setStopInIBootStage1(stage1)
	return nil
}

// CanSetStopInIBootStage1 reports whether the receiver responds to the private selector _setStopInIBootStage1:.
func (v VZMacOSVirtualMachineStartOptions) CanSetStopInIBootStage1() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setStopInIBootStage1:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions/_setStopInIBootStage2:
func (v VZMacOSVirtualMachineStartOptions) _setStopInIBootStage2(stage2 bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setStopInIBootStage2:"), stage2)
}

// SetStopInIBootStage2 is an exported wrapper for the private method _setStopInIBootStage2.
func (v VZMacOSVirtualMachineStartOptions) SetStopInIBootStage2(stage2 bool) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setStopInIBootStage2:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setStopInIBootStage2:"}
		return err
	}
	v._setStopInIBootStage2(stage2)
	return nil
}

// CanSetStopInIBootStage2 reports whether the receiver responds to the private selector _setStopInIBootStage2:.
func (v VZMacOSVirtualMachineStartOptions) CanSetStopInIBootStage2() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setStopInIBootStage2:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions/_forceDFU
func (v VZMacOSVirtualMachineStartOptions) _forceDFU() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_forceDFU"))
	return rv
}

// CanForceDFU reports whether the receiver responds to the private selector _forceDFU.
func (v VZMacOSVirtualMachineStartOptions) CanForceDFU() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_forceDFU"))
}

// ForceDFU is an exported wrapper for the private property _forceDFU.
func (v VZMacOSVirtualMachineStartOptions) ForceDFU() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_forceDFU")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_forceDFU"}
	}
	return v._forceDFU(), nil
}
func (v VZMacOSVirtualMachineStartOptions) Set_forceDFU(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("set_forceDFU:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions/_stopInIBootStage1
func (v VZMacOSVirtualMachineStartOptions) _stopInIBootStage1() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_stopInIBootStage1"))
	return rv
}

// CanStopInIBootStage1 reports whether the receiver responds to the private selector _stopInIBootStage1.
func (v VZMacOSVirtualMachineStartOptions) CanStopInIBootStage1() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_stopInIBootStage1"))
}

// StopInIBootStage1 is an exported wrapper for the private property _stopInIBootStage1.
func (v VZMacOSVirtualMachineStartOptions) StopInIBootStage1() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_stopInIBootStage1")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_stopInIBootStage1"}
	}
	return v._stopInIBootStage1(), nil
}
func (v VZMacOSVirtualMachineStartOptions) Set_stopInIBootStage1(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("set_stopInIBootStage1:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacOSVirtualMachineStartOptions/_stopInIBootStage2
func (v VZMacOSVirtualMachineStartOptions) _stopInIBootStage2() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_stopInIBootStage2"))
	return rv
}

// CanStopInIBootStage2 reports whether the receiver responds to the private selector _stopInIBootStage2.
func (v VZMacOSVirtualMachineStartOptions) CanStopInIBootStage2() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_stopInIBootStage2"))
}

// StopInIBootStage2 is an exported wrapper for the private property _stopInIBootStage2.
func (v VZMacOSVirtualMachineStartOptions) StopInIBootStage2() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_stopInIBootStage2")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_stopInIBootStage2"}
	}
	return v._stopInIBootStage2(), nil
}
func (v VZMacOSVirtualMachineStartOptions) Set_stopInIBootStage2(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("set_stopInIBootStage2:"), value)
}

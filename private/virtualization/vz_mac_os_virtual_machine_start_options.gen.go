// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
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
	rv := objc.SendIfResponds[VZMacOSVirtualMachineStartOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMacOSVirtualMachineStartOptions._forceDFU]
//   - [VZMacOSVirtualMachineStartOptions.Set_forceDFU]
//   - [VZMacOSVirtualMachineStartOptions._guestProvisioningOptionsForApplicationIdentifier]
//   - [VZMacOSVirtualMachineStartOptions._setForceDFU]
//   - [VZMacOSVirtualMachineStartOptions._setGuestProvisioningOptionsForApplicationIdentifierError]
//   - [VZMacOSVirtualMachineStartOptions._setStopInIBootStage1]
//   - [VZMacOSVirtualMachineStartOptions._setStopInIBootStage2]
//   - [VZMacOSVirtualMachineStartOptions._stopInIBootStage1]
//   - [VZMacOSVirtualMachineStartOptions.Set_stopInIBootStage1]
//   - [VZMacOSVirtualMachineStartOptions._stopInIBootStage2]
//   - [VZMacOSVirtualMachineStartOptions.Set_stopInIBootStage2]
//   - [VZMacOSVirtualMachineStartOptions.GuestProvisioningOptions]
//   - [VZMacOSVirtualMachineStartOptions.SetGuestProvisioningOptionsError]
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
//   - [IVZMacOSVirtualMachineStartOptions._guestProvisioningOptionsForApplicationIdentifier]
//   - [IVZMacOSVirtualMachineStartOptions._setForceDFU]
//   - [IVZMacOSVirtualMachineStartOptions._setGuestProvisioningOptionsForApplicationIdentifierError]
//   - [IVZMacOSVirtualMachineStartOptions._setStopInIBootStage1]
//   - [IVZMacOSVirtualMachineStartOptions._setStopInIBootStage2]
//   - [IVZMacOSVirtualMachineStartOptions._stopInIBootStage1]
//   - [IVZMacOSVirtualMachineStartOptions.Set_stopInIBootStage1]
//   - [IVZMacOSVirtualMachineStartOptions._stopInIBootStage2]
//   - [IVZMacOSVirtualMachineStartOptions.Set_stopInIBootStage2]
//   - [IVZMacOSVirtualMachineStartOptions.GuestProvisioningOptions]
//   - [IVZMacOSVirtualMachineStartOptions.SetGuestProvisioningOptionsError]
type IVZMacOSVirtualMachineStartOptions interface {
	IVZVirtualMachineStartOptions

	// Topic: Methods

	_forceDFU() bool
	Set_forceDFU(value bool)
	_guestProvisioningOptionsForApplicationIdentifier(identifier objectivec.IObject) objectivec.IObject
	_setForceDFU(dfu bool)
	_setGuestProvisioningOptionsForApplicationIdentifierError(options objectivec.IObject, identifier objectivec.IObject) (bool, error)
	_setStopInIBootStage1(stage1 bool)
	_setStopInIBootStage2(stage2 bool)
	_stopInIBootStage1() bool
	Set_stopInIBootStage1(value bool)
	_stopInIBootStage2() bool
	Set_stopInIBootStage2(value bool)
	GuestProvisioningOptions() IVZMacGuestProvisioningOptions
	SetGuestProvisioningOptionsError(options objectivec.IObject) (bool, error)
}

// Init initializes the instance.
func (v VZMacOSVirtualMachineStartOptions) Init() VZMacOSVirtualMachineStartOptions {
	rv := objc.SendIfResponds[VZMacOSVirtualMachineStartOptions](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacOSVirtualMachineStartOptions) Autorelease() VZMacOSVirtualMachineStartOptions {
	rv := objc.SendIfResponds[VZMacOSVirtualMachineStartOptions](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacOSVirtualMachineStartOptions creates a new VZMacOSVirtualMachineStartOptions instance.
func NewVZMacOSVirtualMachineStartOptions() VZMacOSVirtualMachineStartOptions {
	class := getVZMacOSVirtualMachineStartOptionsClass()
	rv := objc.SendIfResponds[VZMacOSVirtualMachineStartOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZMacOSVirtualMachineStartOptions) _guestProvisioningOptionsForApplicationIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_guestProvisioningOptionsForApplicationIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}

// GuestProvisioningOptionsForApplicationIdentifier is an exported wrapper for the private method _guestProvisioningOptionsForApplicationIdentifier.
func (v VZMacOSVirtualMachineStartOptions) GuestProvisioningOptionsForApplicationIdentifier(identifier objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_guestProvisioningOptionsForApplicationIdentifier:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_guestProvisioningOptionsForApplicationIdentifier:"}
		return nil, err
	}
	return v._guestProvisioningOptionsForApplicationIdentifier(identifier), nil
}

// CanGuestProvisioningOptionsForApplicationIdentifier reports whether the receiver responds to the private selector _guestProvisioningOptionsForApplicationIdentifier:.
func (v VZMacOSVirtualMachineStartOptions) CanGuestProvisioningOptionsForApplicationIdentifier() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_guestProvisioningOptionsForApplicationIdentifier:"))
}
func (v VZMacOSVirtualMachineStartOptions) _setForceDFU(dfu bool) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_setForceDFU:"), dfu)
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
func (v VZMacOSVirtualMachineStartOptions) _setGuestProvisioningOptionsForApplicationIdentifierError(options objectivec.IObject, identifier objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("_setGuestProvisioningOptions:forApplicationIdentifier:error:"), options, identifier, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_setGuestProvisioningOptions:forApplicationIdentifier:error: returned NO with nil NSError")
	}
	return rv, nil

}

// SetGuestProvisioningOptionsForApplicationIdentifierError is an exported wrapper for the private method _setGuestProvisioningOptionsForApplicationIdentifierError.
func (v VZMacOSVirtualMachineStartOptions) SetGuestProvisioningOptionsForApplicationIdentifierError(options objectivec.IObject, identifier objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setGuestProvisioningOptions:forApplicationIdentifier:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setGuestProvisioningOptions:forApplicationIdentifier:error:"}
		return false, err
	}
	return v._setGuestProvisioningOptionsForApplicationIdentifierError(options, identifier)
}

// CanSetGuestProvisioningOptionsForApplicationIdentifierError reports whether the receiver responds to the private selector _setGuestProvisioningOptions:forApplicationIdentifier:error:.
func (v VZMacOSVirtualMachineStartOptions) CanSetGuestProvisioningOptionsForApplicationIdentifierError() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setGuestProvisioningOptions:forApplicationIdentifier:error:"))
}
func (v VZMacOSVirtualMachineStartOptions) _setStopInIBootStage1(stage1 bool) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_setStopInIBootStage1:"), stage1)
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
func (v VZMacOSVirtualMachineStartOptions) _setStopInIBootStage2(stage2 bool) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_setStopInIBootStage2:"), stage2)
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
func (v VZMacOSVirtualMachineStartOptions) SetGuestProvisioningOptionsError(options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("setGuestProvisioningOptions:error:"), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setGuestProvisioningOptions:error: returned NO with nil NSError")
	}
	return rv, nil

}

func (v VZMacOSVirtualMachineStartOptions) _forceDFU() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("_forceDFU"))
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
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("set_forceDFU:"), value)
}
func (v VZMacOSVirtualMachineStartOptions) _stopInIBootStage1() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("_stopInIBootStage1"))
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
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("set_stopInIBootStage1:"), value)
}
func (v VZMacOSVirtualMachineStartOptions) _stopInIBootStage2() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("_stopInIBootStage2"))
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
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("set_stopInIBootStage2:"), value)
}
func (v VZMacOSVirtualMachineStartOptions) GuestProvisioningOptions() IVZMacGuestProvisioningOptions {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("guestProvisioningOptions"))
	return VZMacGuestProvisioningOptionsFromID(objc.ID(rv))
}

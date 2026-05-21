// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtualMachineConfiguration] class.
var (
	_VZVirtualMachineConfigurationClass     VZVirtualMachineConfigurationClass
	_VZVirtualMachineConfigurationClassOnce sync.Once
)

func getVZVirtualMachineConfigurationClass() VZVirtualMachineConfigurationClass {
	_VZVirtualMachineConfigurationClassOnce.Do(func() {
		_VZVirtualMachineConfigurationClass = VZVirtualMachineConfigurationClass{class: objc.GetClass("VZVirtualMachineConfiguration")}
	})
	return _VZVirtualMachineConfigurationClass
}

// GetVZVirtualMachineConfigurationClass returns the class object for VZVirtualMachineConfiguration.
func GetVZVirtualMachineConfigurationClass() VZVirtualMachineConfigurationClass {
	return getVZVirtualMachineConfigurationClass()
}

type VZVirtualMachineConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtualMachineConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtualMachineConfigurationClass) Alloc() VZVirtualMachineConfiguration {
	rv := objc.Send[VZVirtualMachineConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVirtualMachineConfiguration._acceleratorDevices]
//   - [VZVirtualMachineConfiguration.Set_acceleratorDevices]
//   - [VZVirtualMachineConfiguration._bifrostDevices]
//   - [VZVirtualMachineConfiguration.Set_bifrostDevices]
//   - [VZVirtualMachineConfiguration._biometricDevices]
//   - [VZVirtualMachineConfiguration.Set_biometricDevices]
//   - [VZVirtualMachineConfiguration._coprocessors]
//   - [VZVirtualMachineConfiguration.Set_coprocessors]
//   - [VZVirtualMachineConfiguration._cpuEmulator]
//   - [VZVirtualMachineConfiguration.Set_cpuEmulator]
//   - [VZVirtualMachineConfiguration._customMMIODevices]
//   - [VZVirtualMachineConfiguration.Set_customMMIODevices]
//   - [VZVirtualMachineConfiguration._customVirtioDevices]
//   - [VZVirtualMachineConfiguration.Set_customVirtioDevices]
//   - [VZVirtualMachineConfiguration._debugStub]
//   - [VZVirtualMachineConfiguration.Set_debugStub]
//   - [VZVirtualMachineConfiguration._fatalErrorAction]
//   - [VZVirtualMachineConfiguration.Set_fatalErrorAction]
//   - [VZVirtualMachineConfiguration._hidDevices]
//   - [VZVirtualMachineConfiguration.Set_hidDevices]
//   - [VZVirtualMachineConfiguration._isDuplicateUSBDeviceConfigurationAtUsbDeviceIndex]
//   - [VZVirtualMachineConfiguration._mailboxDevices]
//   - [VZVirtualMachineConfiguration.Set_mailboxDevices]
//   - [VZVirtualMachineConfiguration._memoryOvercommitmentAllowed]
//   - [VZVirtualMachineConfiguration.Set_memoryOvercommitmentAllowed]
//   - [VZVirtualMachineConfiguration._multiTouchDevices]
//   - [VZVirtualMachineConfiguration.Set_multiTouchDevices]
//   - [VZVirtualMachineConfiguration._panicAction]
//   - [VZVirtualMachineConfiguration.Set_panicAction]
//   - [VZVirtualMachineConfiguration._panicDevice]
//   - [VZVirtualMachineConfiguration.Set_panicDevice]
//   - [VZVirtualMachineConfiguration._pciPassthroughDevices]
//   - [VZVirtualMachineConfiguration.Set_pciPassthroughDevices]
//   - [VZVirtualMachineConfiguration._powerSourceDevices]
//   - [VZVirtualMachineConfiguration.Set_powerSourceDevices]
//   - [VZVirtualMachineConfiguration._restartAction]
//   - [VZVirtualMachineConfiguration.Set_restartAction]
//   - [VZVirtualMachineConfiguration._setAcceleratorDevices]
//   - [VZVirtualMachineConfiguration._setBifrostDevices]
//   - [VZVirtualMachineConfiguration._setBiometricDevices]
//   - [VZVirtualMachineConfiguration._setCPUEmulator]
//   - [VZVirtualMachineConfiguration._setCoprocessors]
//   - [VZVirtualMachineConfiguration._setCustomMMIODevices]
//   - [VZVirtualMachineConfiguration._setCustomVirtioDevices]
//   - [VZVirtualMachineConfiguration._setDebugStub]
//   - [VZVirtualMachineConfiguration._setFatalErrorAction]
//   - [VZVirtualMachineConfiguration._setHIDDevices]
//   - [VZVirtualMachineConfiguration._setMailboxDevices]
//   - [VZVirtualMachineConfiguration._setMemoryOvercommitmentAllowed]
//   - [VZVirtualMachineConfiguration._setMultiTouchDevices]
//   - [VZVirtualMachineConfiguration._setPCIPassthroughDevices]
//   - [VZVirtualMachineConfiguration._setPanicAction]
//   - [VZVirtualMachineConfiguration._setPanicDevice]
//   - [VZVirtualMachineConfiguration._setPowerSourceDevices]
//   - [VZVirtualMachineConfiguration._setRestartAction]
//   - [VZVirtualMachineConfiguration._setTerminationUnderMemoryPressureEnabled]
//   - [VZVirtualMachineConfiguration._setTestIgnoreEntitlementChecks]
//   - [VZVirtualMachineConfiguration._terminationUnderMemoryPressureEnabled]
//   - [VZVirtualMachineConfiguration.Set_terminationUnderMemoryPressureEnabled]
//   - [VZVirtualMachineConfiguration._testIgnoreEntitlementChecks]
//   - [VZVirtualMachineConfiguration.Set_testIgnoreEntitlementChecks]
//   - [VZVirtualMachineConfiguration._sharedRamRegions]
//   - [VZVirtualMachineConfiguration.Set_sharedRamRegions]
type VZVirtualMachineConfiguration struct {
	objectivec.Object
}

// VZVirtualMachineConfigurationFromID constructs a [VZVirtualMachineConfiguration] from an objc.ID.
func VZVirtualMachineConfigurationFromID(id objc.ID) VZVirtualMachineConfiguration {
	return VZVirtualMachineConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZVirtualMachineConfiguration implements IVZVirtualMachineConfiguration.
var _ IVZVirtualMachineConfiguration = VZVirtualMachineConfiguration{}

// An interface definition for the [VZVirtualMachineConfiguration] class.
//
// # Methods
//
//   - [IVZVirtualMachineConfiguration._acceleratorDevices]
//   - [IVZVirtualMachineConfiguration.Set_acceleratorDevices]
//   - [IVZVirtualMachineConfiguration._bifrostDevices]
//   - [IVZVirtualMachineConfiguration.Set_bifrostDevices]
//   - [IVZVirtualMachineConfiguration._biometricDevices]
//   - [IVZVirtualMachineConfiguration.Set_biometricDevices]
//   - [IVZVirtualMachineConfiguration._coprocessors]
//   - [IVZVirtualMachineConfiguration.Set_coprocessors]
//   - [IVZVirtualMachineConfiguration._cpuEmulator]
//   - [IVZVirtualMachineConfiguration.Set_cpuEmulator]
//   - [IVZVirtualMachineConfiguration._customMMIODevices]
//   - [IVZVirtualMachineConfiguration.Set_customMMIODevices]
//   - [IVZVirtualMachineConfiguration._customVirtioDevices]
//   - [IVZVirtualMachineConfiguration.Set_customVirtioDevices]
//   - [IVZVirtualMachineConfiguration._debugStub]
//   - [IVZVirtualMachineConfiguration.Set_debugStub]
//   - [IVZVirtualMachineConfiguration._fatalErrorAction]
//   - [IVZVirtualMachineConfiguration.Set_fatalErrorAction]
//   - [IVZVirtualMachineConfiguration._hidDevices]
//   - [IVZVirtualMachineConfiguration.Set_hidDevices]
//   - [IVZVirtualMachineConfiguration._isDuplicateUSBDeviceConfigurationAtUsbDeviceIndex]
//   - [IVZVirtualMachineConfiguration._mailboxDevices]
//   - [IVZVirtualMachineConfiguration.Set_mailboxDevices]
//   - [IVZVirtualMachineConfiguration._memoryOvercommitmentAllowed]
//   - [IVZVirtualMachineConfiguration.Set_memoryOvercommitmentAllowed]
//   - [IVZVirtualMachineConfiguration._multiTouchDevices]
//   - [IVZVirtualMachineConfiguration.Set_multiTouchDevices]
//   - [IVZVirtualMachineConfiguration._panicAction]
//   - [IVZVirtualMachineConfiguration.Set_panicAction]
//   - [IVZVirtualMachineConfiguration._panicDevice]
//   - [IVZVirtualMachineConfiguration.Set_panicDevice]
//   - [IVZVirtualMachineConfiguration._pciPassthroughDevices]
//   - [IVZVirtualMachineConfiguration.Set_pciPassthroughDevices]
//   - [IVZVirtualMachineConfiguration._powerSourceDevices]
//   - [IVZVirtualMachineConfiguration.Set_powerSourceDevices]
//   - [IVZVirtualMachineConfiguration._restartAction]
//   - [IVZVirtualMachineConfiguration.Set_restartAction]
//   - [IVZVirtualMachineConfiguration._setAcceleratorDevices]
//   - [IVZVirtualMachineConfiguration._setBifrostDevices]
//   - [IVZVirtualMachineConfiguration._setBiometricDevices]
//   - [IVZVirtualMachineConfiguration._setCPUEmulator]
//   - [IVZVirtualMachineConfiguration._setCoprocessors]
//   - [IVZVirtualMachineConfiguration._setCustomMMIODevices]
//   - [IVZVirtualMachineConfiguration._setCustomVirtioDevices]
//   - [IVZVirtualMachineConfiguration._setDebugStub]
//   - [IVZVirtualMachineConfiguration._setFatalErrorAction]
//   - [IVZVirtualMachineConfiguration._setHIDDevices]
//   - [IVZVirtualMachineConfiguration._setMailboxDevices]
//   - [IVZVirtualMachineConfiguration._setMemoryOvercommitmentAllowed]
//   - [IVZVirtualMachineConfiguration._setMultiTouchDevices]
//   - [IVZVirtualMachineConfiguration._setPCIPassthroughDevices]
//   - [IVZVirtualMachineConfiguration._setPanicAction]
//   - [IVZVirtualMachineConfiguration._setPanicDevice]
//   - [IVZVirtualMachineConfiguration._setPowerSourceDevices]
//   - [IVZVirtualMachineConfiguration._setRestartAction]
//   - [IVZVirtualMachineConfiguration._setTerminationUnderMemoryPressureEnabled]
//   - [IVZVirtualMachineConfiguration._setTestIgnoreEntitlementChecks]
//   - [IVZVirtualMachineConfiguration._terminationUnderMemoryPressureEnabled]
//   - [IVZVirtualMachineConfiguration.Set_terminationUnderMemoryPressureEnabled]
//   - [IVZVirtualMachineConfiguration._testIgnoreEntitlementChecks]
//   - [IVZVirtualMachineConfiguration.Set_testIgnoreEntitlementChecks]
//   - [IVZVirtualMachineConfiguration._sharedRamRegions]
//   - [IVZVirtualMachineConfiguration.Set_sharedRamRegions]
type IVZVirtualMachineConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_acceleratorDevices() foundation.INSArray
	Set_acceleratorDevices(value foundation.INSArray)
	_bifrostDevices() foundation.INSArray
	Set_bifrostDevices(value foundation.INSArray)
	_biometricDevices() foundation.INSArray
	Set_biometricDevices(value foundation.INSArray)
	_coprocessors() foundation.INSArray
	Set_coprocessors(value foundation.INSArray)
	_cpuEmulator() IVZCPUEmulatorConfiguration
	Set_cpuEmulator(value IVZCPUEmulatorConfiguration)
	_customMMIODevices() foundation.INSArray
	Set_customMMIODevices(value foundation.INSArray)
	_customVirtioDevices() foundation.INSArray
	Set_customVirtioDevices(value foundation.INSArray)
	_debugStub() IVZDebugStubConfiguration
	Set_debugStub(value IVZDebugStubConfiguration)
	_fatalErrorAction() int64
	Set_fatalErrorAction(value int64)
	_hidDevices() foundation.INSArray
	Set_hidDevices(value foundation.INSArray)
	_isDuplicateUSBDeviceConfigurationAtUsbDeviceIndex(at uint64, index uint64) bool
	_mailboxDevices() foundation.INSArray
	Set_mailboxDevices(value foundation.INSArray)
	_memoryOvercommitmentAllowed() bool
	Set_memoryOvercommitmentAllowed(value bool)
	_multiTouchDevices() foundation.INSArray
	Set_multiTouchDevices(value foundation.INSArray)
	_panicAction() int64
	Set_panicAction(value int64)
	_panicDevice() IVZPanicDeviceConfiguration
	Set_panicDevice(value IVZPanicDeviceConfiguration)
	_pciPassthroughDevices() foundation.INSArray
	Set_pciPassthroughDevices(value foundation.INSArray)
	_powerSourceDevices() foundation.INSArray
	Set_powerSourceDevices(value foundation.INSArray)
	_restartAction() int64
	Set_restartAction(value int64)
	_setAcceleratorDevices(devices objectivec.IObject)
	_setBifrostDevices(devices objectivec.IObject)
	_setBiometricDevices(devices objectivec.IObject)
	_setCPUEmulator(cPUEmulator objectivec.IObject)
	_setCoprocessors(coprocessors objectivec.IObject)
	_setCustomMMIODevices(mMIODevices objectivec.IObject)
	_setCustomVirtioDevices(devices objectivec.IObject)
	_setDebugStub(stub objectivec.IObject)
	_setFatalErrorAction(action int64)
	_setHIDDevices(hIDDevices objectivec.IObject)
	_setMailboxDevices(devices objectivec.IObject)
	_setMemoryOvercommitmentAllowed(allowed bool)
	_setMultiTouchDevices(devices objectivec.IObject)
	_setPCIPassthroughDevices(devices objectivec.IObject)
	_setPanicAction(action int64)
	_setPanicDevice(device objectivec.IObject)
	_setPowerSourceDevices(devices objectivec.IObject)
	_setRestartAction(action int64)
	_setTerminationUnderMemoryPressureEnabled(enabled bool)
	_setTestIgnoreEntitlementChecks(checks bool)
	_terminationUnderMemoryPressureEnabled() bool
	Set_terminationUnderMemoryPressureEnabled(value bool)
	_testIgnoreEntitlementChecks() bool
	Set_testIgnoreEntitlementChecks(value bool)
	_sharedRamRegions() unsafe.Pointer
	Set_sharedRamRegions(value unsafe.Pointer)

	// The number of CPUs for the virtual machine. Must be between minimumAllowedCPUCount and maximumAllowedCPUCount.
	CPUCount() uint
	SetCPUCount(value uint)
	// The memory size in bytes for the virtual machine. Must be a multiple of 1MB and between minimumAllowedMemorySize and maximumAllowedMemorySize.
	MemorySize() uint64
	SetMemorySize(value uint64)
}

// Init initializes the instance.
func (v VZVirtualMachineConfiguration) Init() VZVirtualMachineConfiguration {
	rv := objc.Send[VZVirtualMachineConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtualMachineConfiguration) Autorelease() VZVirtualMachineConfiguration {
	rv := objc.Send[VZVirtualMachineConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachineConfiguration creates a new VZVirtualMachineConfiguration instance.
func NewVZVirtualMachineConfiguration() VZVirtualMachineConfiguration {
	class := getVZVirtualMachineConfigurationClass()
	rv := objc.Send[VZVirtualMachineConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZVirtualMachineConfiguration) _isDuplicateUSBDeviceConfigurationAtUsbDeviceIndex(at uint64, index uint64) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_isDuplicateUSBDeviceConfigurationAt:usbDeviceIndex:"), at, index)
	return rv
}

// IsDuplicateUSBDeviceConfigurationAtUsbDeviceIndex is an exported wrapper for the private method _isDuplicateUSBDeviceConfigurationAtUsbDeviceIndex.
func (v VZVirtualMachineConfiguration) IsDuplicateUSBDeviceConfigurationAtUsbDeviceIndex(at uint64, index uint64) (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_isDuplicateUSBDeviceConfigurationAt:usbDeviceIndex:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_isDuplicateUSBDeviceConfigurationAt:usbDeviceIndex:"}
		return false, err
	}
	return v._isDuplicateUSBDeviceConfigurationAtUsbDeviceIndex(at, index), nil
}

// CanIsDuplicateUSBDeviceConfigurationAtUsbDeviceIndex reports whether the receiver responds to the private selector _isDuplicateUSBDeviceConfigurationAt:usbDeviceIndex:.
func (v VZVirtualMachineConfiguration) CanIsDuplicateUSBDeviceConfigurationAtUsbDeviceIndex() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_isDuplicateUSBDeviceConfigurationAt:usbDeviceIndex:"))
}
func (v VZVirtualMachineConfiguration) _setAcceleratorDevices(devices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setAcceleratorDevices:"), devices)
}

// SetAcceleratorDevices is an exported wrapper for the private method _setAcceleratorDevices.
func (v VZVirtualMachineConfiguration) SetAcceleratorDevices(devices objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setAcceleratorDevices:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setAcceleratorDevices:"}
		return err
	}
	v._setAcceleratorDevices(devices)
	return nil
}

// CanSetAcceleratorDevices reports whether the receiver responds to the private selector _setAcceleratorDevices:.
func (v VZVirtualMachineConfiguration) CanSetAcceleratorDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setAcceleratorDevices:"))
}
func (v VZVirtualMachineConfiguration) _setBifrostDevices(devices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setBifrostDevices:"), devices)
}

// SetBifrostDevices is an exported wrapper for the private method _setBifrostDevices.
func (v VZVirtualMachineConfiguration) SetBifrostDevices(devices objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setBifrostDevices:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setBifrostDevices:"}
		return err
	}
	v._setBifrostDevices(devices)
	return nil
}

// CanSetBifrostDevices reports whether the receiver responds to the private selector _setBifrostDevices:.
func (v VZVirtualMachineConfiguration) CanSetBifrostDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setBifrostDevices:"))
}
func (v VZVirtualMachineConfiguration) _setBiometricDevices(devices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setBiometricDevices:"), devices)
}

// SetBiometricDevices is an exported wrapper for the private method _setBiometricDevices.
func (v VZVirtualMachineConfiguration) SetBiometricDevices(devices objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setBiometricDevices:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setBiometricDevices:"}
		return err
	}
	v._setBiometricDevices(devices)
	return nil
}

// CanSetBiometricDevices reports whether the receiver responds to the private selector _setBiometricDevices:.
func (v VZVirtualMachineConfiguration) CanSetBiometricDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setBiometricDevices:"))
}
func (v VZVirtualMachineConfiguration) _setCPUEmulator(cPUEmulator objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setCPUEmulator:"), cPUEmulator)
}

// SetCPUEmulator is an exported wrapper for the private method _setCPUEmulator.
func (v VZVirtualMachineConfiguration) SetCPUEmulator(cPUEmulator objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setCPUEmulator:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setCPUEmulator:"}
		return err
	}
	v._setCPUEmulator(cPUEmulator)
	return nil
}

// CanSetCPUEmulator reports whether the receiver responds to the private selector _setCPUEmulator:.
func (v VZVirtualMachineConfiguration) CanSetCPUEmulator() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setCPUEmulator:"))
}
func (v VZVirtualMachineConfiguration) _setCoprocessors(coprocessors objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setCoprocessors:"), coprocessors)
}

// SetCoprocessors is an exported wrapper for the private method _setCoprocessors.
func (v VZVirtualMachineConfiguration) SetCoprocessors(coprocessors objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setCoprocessors:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setCoprocessors:"}
		return err
	}
	v._setCoprocessors(coprocessors)
	return nil
}

// CanSetCoprocessors reports whether the receiver responds to the private selector _setCoprocessors:.
func (v VZVirtualMachineConfiguration) CanSetCoprocessors() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setCoprocessors:"))
}
func (v VZVirtualMachineConfiguration) _setCustomMMIODevices(mMIODevices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setCustomMMIODevices:"), mMIODevices)
}

// SetCustomMMIODevices is an exported wrapper for the private method _setCustomMMIODevices.
func (v VZVirtualMachineConfiguration) SetCustomMMIODevices(mMIODevices objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setCustomMMIODevices:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setCustomMMIODevices:"}
		return err
	}
	v._setCustomMMIODevices(mMIODevices)
	return nil
}

// CanSetCustomMMIODevices reports whether the receiver responds to the private selector _setCustomMMIODevices:.
func (v VZVirtualMachineConfiguration) CanSetCustomMMIODevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setCustomMMIODevices:"))
}
func (v VZVirtualMachineConfiguration) _setCustomVirtioDevices(devices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setCustomVirtioDevices:"), devices)
}

// SetCustomVirtioDevices is an exported wrapper for the private method _setCustomVirtioDevices.
func (v VZVirtualMachineConfiguration) SetCustomVirtioDevices(devices objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setCustomVirtioDevices:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setCustomVirtioDevices:"}
		return err
	}
	v._setCustomVirtioDevices(devices)
	return nil
}

// CanSetCustomVirtioDevices reports whether the receiver responds to the private selector _setCustomVirtioDevices:.
func (v VZVirtualMachineConfiguration) CanSetCustomVirtioDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setCustomVirtioDevices:"))
}
func (v VZVirtualMachineConfiguration) _setDebugStub(stub objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setDebugStub:"), stub)
}

// SetDebugStub is an exported wrapper for the private method _setDebugStub.
func (v VZVirtualMachineConfiguration) SetDebugStub(stub objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setDebugStub:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setDebugStub:"}
		return err
	}
	v._setDebugStub(stub)
	return nil
}

// CanSetDebugStub reports whether the receiver responds to the private selector _setDebugStub:.
func (v VZVirtualMachineConfiguration) CanSetDebugStub() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setDebugStub:"))
}
func (v VZVirtualMachineConfiguration) _setFatalErrorAction(action int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setFatalErrorAction:"), action)
}

// SetFatalErrorAction is an exported wrapper for the private method _setFatalErrorAction.
func (v VZVirtualMachineConfiguration) SetFatalErrorAction(action int64) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setFatalErrorAction:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setFatalErrorAction:"}
		return err
	}
	v._setFatalErrorAction(action)
	return nil
}

// CanSetFatalErrorAction reports whether the receiver responds to the private selector _setFatalErrorAction:.
func (v VZVirtualMachineConfiguration) CanSetFatalErrorAction() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setFatalErrorAction:"))
}
func (v VZVirtualMachineConfiguration) _setHIDDevices(hIDDevices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setHIDDevices:"), hIDDevices)
}

// SetHIDDevices is an exported wrapper for the private method _setHIDDevices.
func (v VZVirtualMachineConfiguration) SetHIDDevices(hIDDevices objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setHIDDevices:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setHIDDevices:"}
		return err
	}
	v._setHIDDevices(hIDDevices)
	return nil
}

// CanSetHIDDevices reports whether the receiver responds to the private selector _setHIDDevices:.
func (v VZVirtualMachineConfiguration) CanSetHIDDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setHIDDevices:"))
}
func (v VZVirtualMachineConfiguration) _setMailboxDevices(devices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setMailboxDevices:"), devices)
}

// SetMailboxDevices is an exported wrapper for the private method _setMailboxDevices.
func (v VZVirtualMachineConfiguration) SetMailboxDevices(devices objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setMailboxDevices:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setMailboxDevices:"}
		return err
	}
	v._setMailboxDevices(devices)
	return nil
}

// CanSetMailboxDevices reports whether the receiver responds to the private selector _setMailboxDevices:.
func (v VZVirtualMachineConfiguration) CanSetMailboxDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setMailboxDevices:"))
}
func (v VZVirtualMachineConfiguration) _setMemoryOvercommitmentAllowed(allowed bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setMemoryOvercommitmentAllowed:"), allowed)
}

// SetMemoryOvercommitmentAllowed is an exported wrapper for the private method _setMemoryOvercommitmentAllowed.
func (v VZVirtualMachineConfiguration) SetMemoryOvercommitmentAllowed(allowed bool) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setMemoryOvercommitmentAllowed:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setMemoryOvercommitmentAllowed:"}
		return err
	}
	v._setMemoryOvercommitmentAllowed(allowed)
	return nil
}

// CanSetMemoryOvercommitmentAllowed reports whether the receiver responds to the private selector _setMemoryOvercommitmentAllowed:.
func (v VZVirtualMachineConfiguration) CanSetMemoryOvercommitmentAllowed() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setMemoryOvercommitmentAllowed:"))
}
func (v VZVirtualMachineConfiguration) _setMultiTouchDevices(devices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setMultiTouchDevices:"), devices)
}

// SetMultiTouchDevices is an exported wrapper for the private method _setMultiTouchDevices.
func (v VZVirtualMachineConfiguration) SetMultiTouchDevices(devices objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setMultiTouchDevices:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setMultiTouchDevices:"}
		return err
	}
	v._setMultiTouchDevices(devices)
	return nil
}

// CanSetMultiTouchDevices reports whether the receiver responds to the private selector _setMultiTouchDevices:.
func (v VZVirtualMachineConfiguration) CanSetMultiTouchDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setMultiTouchDevices:"))
}
func (v VZVirtualMachineConfiguration) _setPCIPassthroughDevices(devices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setPCIPassthroughDevices:"), devices)
}

// SetPCIPassthroughDevices is an exported wrapper for the private method _setPCIPassthroughDevices.
func (v VZVirtualMachineConfiguration) SetPCIPassthroughDevices(devices objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setPCIPassthroughDevices:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setPCIPassthroughDevices:"}
		return err
	}
	v._setPCIPassthroughDevices(devices)
	return nil
}

// CanSetPCIPassthroughDevices reports whether the receiver responds to the private selector _setPCIPassthroughDevices:.
func (v VZVirtualMachineConfiguration) CanSetPCIPassthroughDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setPCIPassthroughDevices:"))
}
func (v VZVirtualMachineConfiguration) _setPanicAction(action int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setPanicAction:"), action)
}

// SetPanicAction is an exported wrapper for the private method _setPanicAction.
func (v VZVirtualMachineConfiguration) SetPanicAction(action int64) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setPanicAction:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setPanicAction:"}
		return err
	}
	v._setPanicAction(action)
	return nil
}

// CanSetPanicAction reports whether the receiver responds to the private selector _setPanicAction:.
func (v VZVirtualMachineConfiguration) CanSetPanicAction() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setPanicAction:"))
}
func (v VZVirtualMachineConfiguration) _setPanicDevice(device objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setPanicDevice:"), device)
}

// SetPanicDevice is an exported wrapper for the private method _setPanicDevice.
func (v VZVirtualMachineConfiguration) SetPanicDevice(device objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setPanicDevice:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setPanicDevice:"}
		return err
	}
	v._setPanicDevice(device)
	return nil
}

// CanSetPanicDevice reports whether the receiver responds to the private selector _setPanicDevice:.
func (v VZVirtualMachineConfiguration) CanSetPanicDevice() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setPanicDevice:"))
}
func (v VZVirtualMachineConfiguration) _setPowerSourceDevices(devices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setPowerSourceDevices:"), devices)
}

// SetPowerSourceDevices is an exported wrapper for the private method _setPowerSourceDevices.
func (v VZVirtualMachineConfiguration) SetPowerSourceDevices(devices objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setPowerSourceDevices:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setPowerSourceDevices:"}
		return err
	}
	v._setPowerSourceDevices(devices)
	return nil
}

// CanSetPowerSourceDevices reports whether the receiver responds to the private selector _setPowerSourceDevices:.
func (v VZVirtualMachineConfiguration) CanSetPowerSourceDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setPowerSourceDevices:"))
}
func (v VZVirtualMachineConfiguration) _setRestartAction(action int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setRestartAction:"), action)
}

// SetRestartAction is an exported wrapper for the private method _setRestartAction.
func (v VZVirtualMachineConfiguration) SetRestartAction(action int64) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setRestartAction:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setRestartAction:"}
		return err
	}
	v._setRestartAction(action)
	return nil
}

// CanSetRestartAction reports whether the receiver responds to the private selector _setRestartAction:.
func (v VZVirtualMachineConfiguration) CanSetRestartAction() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setRestartAction:"))
}
func (v VZVirtualMachineConfiguration) _setTerminationUnderMemoryPressureEnabled(enabled bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setTerminationUnderMemoryPressureEnabled:"), enabled)
}

// SetTerminationUnderMemoryPressureEnabled is an exported wrapper for the private method _setTerminationUnderMemoryPressureEnabled.
func (v VZVirtualMachineConfiguration) SetTerminationUnderMemoryPressureEnabled(enabled bool) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setTerminationUnderMemoryPressureEnabled:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setTerminationUnderMemoryPressureEnabled:"}
		return err
	}
	v._setTerminationUnderMemoryPressureEnabled(enabled)
	return nil
}

// CanSetTerminationUnderMemoryPressureEnabled reports whether the receiver responds to the private selector _setTerminationUnderMemoryPressureEnabled:.
func (v VZVirtualMachineConfiguration) CanSetTerminationUnderMemoryPressureEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setTerminationUnderMemoryPressureEnabled:"))
}
func (v VZVirtualMachineConfiguration) _setTestIgnoreEntitlementChecks(checks bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setTestIgnoreEntitlementChecks:"), checks)
}

// SetTestIgnoreEntitlementChecks is an exported wrapper for the private method _setTestIgnoreEntitlementChecks.
func (v VZVirtualMachineConfiguration) SetTestIgnoreEntitlementChecks(checks bool) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setTestIgnoreEntitlementChecks:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setTestIgnoreEntitlementChecks:"}
		return err
	}
	v._setTestIgnoreEntitlementChecks(checks)
	return nil
}

// CanSetTestIgnoreEntitlementChecks reports whether the receiver responds to the private selector _setTestIgnoreEntitlementChecks:.
func (v VZVirtualMachineConfiguration) CanSetTestIgnoreEntitlementChecks() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setTestIgnoreEntitlementChecks:"))
}

func (_VZVirtualMachineConfigurationClass VZVirtualMachineConfigurationClass) _maximumAllowedOvercommittedMemorySize() uint64 {
	rv := objc.Send[uint64](objc.ID(_VZVirtualMachineConfigurationClass.class), objc.Sel("_maximumAllowedOvercommittedMemorySize"))
	return rv
}

// MaximumAllowedOvercommittedMemorySize is an exported wrapper for the private method _maximumAllowedOvercommittedMemorySize.
func (_VZVirtualMachineConfigurationClass VZVirtualMachineConfigurationClass) MaximumAllowedOvercommittedMemorySize() (uint64, error) {
	if !objc.RespondsToSelector(objc.ID(_VZVirtualMachineConfigurationClass.class), objc.Sel("_maximumAllowedOvercommittedMemorySize")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_maximumAllowedOvercommittedMemorySize"}
		return 0, err
	}
	return _VZVirtualMachineConfigurationClass._maximumAllowedOvercommittedMemorySize(), nil
}

// CanMaximumAllowedOvercommittedMemorySize reports whether the receiver responds to the private selector _maximumAllowedOvercommittedMemorySize.
func (_VZVirtualMachineConfigurationClass VZVirtualMachineConfigurationClass) CanMaximumAllowedOvercommittedMemorySize() bool {
	return objc.RespondsToSelector(objc.ID(_VZVirtualMachineConfigurationClass.class), objc.Sel("_maximumAllowedOvercommittedMemorySize"))
}

func (v VZVirtualMachineConfiguration) _acceleratorDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_acceleratorDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanAcceleratorDevices reports whether the receiver responds to the private selector _acceleratorDevices.
func (v VZVirtualMachineConfiguration) CanAcceleratorDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_acceleratorDevices"))
}

// AcceleratorDevices is an exported wrapper for the private property _acceleratorDevices.
func (v VZVirtualMachineConfiguration) AcceleratorDevices() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_acceleratorDevices")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_acceleratorDevices"}
	}
	return v._acceleratorDevices(), nil
}
func (v VZVirtualMachineConfiguration) Set_acceleratorDevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_acceleratorDevices:"), value)
}
func (v VZVirtualMachineConfiguration) _bifrostDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_bifrostDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanBifrostDevices reports whether the receiver responds to the private selector _bifrostDevices.
func (v VZVirtualMachineConfiguration) CanBifrostDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_bifrostDevices"))
}

// BifrostDevices is an exported wrapper for the private property _bifrostDevices.
func (v VZVirtualMachineConfiguration) BifrostDevices() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_bifrostDevices")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_bifrostDevices"}
	}
	return v._bifrostDevices(), nil
}
func (v VZVirtualMachineConfiguration) Set_bifrostDevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_bifrostDevices:"), value)
}
func (v VZVirtualMachineConfiguration) _biometricDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_biometricDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanBiometricDevices reports whether the receiver responds to the private selector _biometricDevices.
func (v VZVirtualMachineConfiguration) CanBiometricDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_biometricDevices"))
}

// BiometricDevices is an exported wrapper for the private property _biometricDevices.
func (v VZVirtualMachineConfiguration) BiometricDevices() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_biometricDevices")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_biometricDevices"}
	}
	return v._biometricDevices(), nil
}
func (v VZVirtualMachineConfiguration) Set_biometricDevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_biometricDevices:"), value)
}
func (v VZVirtualMachineConfiguration) _coprocessors() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_coprocessors"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanCoprocessors reports whether the receiver responds to the private selector _coprocessors.
func (v VZVirtualMachineConfiguration) CanCoprocessors() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_coprocessors"))
}

// Coprocessors is an exported wrapper for the private property _coprocessors.
func (v VZVirtualMachineConfiguration) Coprocessors() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_coprocessors")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_coprocessors"}
	}
	return v._coprocessors(), nil
}
func (v VZVirtualMachineConfiguration) Set_coprocessors(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_coprocessors:"), value)
}
func (v VZVirtualMachineConfiguration) _cpuEmulator() IVZCPUEmulatorConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_cpuEmulator"))
	return VZCPUEmulatorConfigurationFromID(objc.ID(rv))
}

// CanCpuEmulator reports whether the receiver responds to the private selector _cpuEmulator.
func (v VZVirtualMachineConfiguration) CanCpuEmulator() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_cpuEmulator"))
}

// CpuEmulator is an exported wrapper for the private property _cpuEmulator.
func (v VZVirtualMachineConfiguration) CpuEmulator() (IVZCPUEmulatorConfiguration, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_cpuEmulator")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_cpuEmulator"}
	}
	return v._cpuEmulator(), nil
}
func (v VZVirtualMachineConfiguration) Set_cpuEmulator(value IVZCPUEmulatorConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("set_cpuEmulator:"), value)
}
func (v VZVirtualMachineConfiguration) _customMMIODevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_customMMIODevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanCustomMMIODevices reports whether the receiver responds to the private selector _customMMIODevices.
func (v VZVirtualMachineConfiguration) CanCustomMMIODevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_customMMIODevices"))
}

// CustomMMIODevices is an exported wrapper for the private property _customMMIODevices.
func (v VZVirtualMachineConfiguration) CustomMMIODevices() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_customMMIODevices")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_customMMIODevices"}
	}
	return v._customMMIODevices(), nil
}
func (v VZVirtualMachineConfiguration) Set_customMMIODevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_customMMIODevices:"), value)
}
func (v VZVirtualMachineConfiguration) _customVirtioDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_customVirtioDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanCustomVirtioDevices reports whether the receiver responds to the private selector _customVirtioDevices.
func (v VZVirtualMachineConfiguration) CanCustomVirtioDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_customVirtioDevices"))
}

// CustomVirtioDevices is an exported wrapper for the private property _customVirtioDevices.
func (v VZVirtualMachineConfiguration) CustomVirtioDevices() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_customVirtioDevices")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_customVirtioDevices"}
	}
	return v._customVirtioDevices(), nil
}
func (v VZVirtualMachineConfiguration) Set_customVirtioDevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_customVirtioDevices:"), value)
}
func (v VZVirtualMachineConfiguration) _debugStub() IVZDebugStubConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_debugStub"))
	return VZDebugStubConfigurationFromID(objc.ID(rv))
}

// CanDebugStub reports whether the receiver responds to the private selector _debugStub.
func (v VZVirtualMachineConfiguration) CanDebugStub() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_debugStub"))
}

// DebugStub is an exported wrapper for the private property _debugStub.
func (v VZVirtualMachineConfiguration) DebugStub() (IVZDebugStubConfiguration, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_debugStub")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_debugStub"}
	}
	return v._debugStub(), nil
}
func (v VZVirtualMachineConfiguration) Set_debugStub(value IVZDebugStubConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("set_debugStub:"), value)
}
func (v VZVirtualMachineConfiguration) _fatalErrorAction() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("_fatalErrorAction"))
	return rv
}

// CanFatalErrorAction reports whether the receiver responds to the private selector _fatalErrorAction.
func (v VZVirtualMachineConfiguration) CanFatalErrorAction() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_fatalErrorAction"))
}

// FatalErrorAction is an exported wrapper for the private property _fatalErrorAction.
func (v VZVirtualMachineConfiguration) FatalErrorAction() (int64, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_fatalErrorAction")) {
		return 0, &objc.UnrecognizedSelectorError{Selector: "_fatalErrorAction"}
	}
	return v._fatalErrorAction(), nil
}
func (v VZVirtualMachineConfiguration) Set_fatalErrorAction(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("set_fatalErrorAction:"), value)
}
func (v VZVirtualMachineConfiguration) _hidDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_hidDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanHidDevices reports whether the receiver responds to the private selector _hidDevices.
func (v VZVirtualMachineConfiguration) CanHidDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_hidDevices"))
}

// HidDevices is an exported wrapper for the private property _hidDevices.
func (v VZVirtualMachineConfiguration) HidDevices() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_hidDevices")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_hidDevices"}
	}
	return v._hidDevices(), nil
}
func (v VZVirtualMachineConfiguration) Set_hidDevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_hidDevices:"), value)
}
func (v VZVirtualMachineConfiguration) _mailboxDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_mailboxDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanMailboxDevices reports whether the receiver responds to the private selector _mailboxDevices.
func (v VZVirtualMachineConfiguration) CanMailboxDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_mailboxDevices"))
}

// MailboxDevices is an exported wrapper for the private property _mailboxDevices.
func (v VZVirtualMachineConfiguration) MailboxDevices() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_mailboxDevices")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_mailboxDevices"}
	}
	return v._mailboxDevices(), nil
}
func (v VZVirtualMachineConfiguration) Set_mailboxDevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_mailboxDevices:"), value)
}
func (v VZVirtualMachineConfiguration) _memoryOvercommitmentAllowed() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_memoryOvercommitmentAllowed"))
	return rv
}

// CanMemoryOvercommitmentAllowed reports whether the receiver responds to the private selector _memoryOvercommitmentAllowed.
func (v VZVirtualMachineConfiguration) CanMemoryOvercommitmentAllowed() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_memoryOvercommitmentAllowed"))
}

// MemoryOvercommitmentAllowed is an exported wrapper for the private property _memoryOvercommitmentAllowed.
func (v VZVirtualMachineConfiguration) MemoryOvercommitmentAllowed() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_memoryOvercommitmentAllowed")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_memoryOvercommitmentAllowed"}
	}
	return v._memoryOvercommitmentAllowed(), nil
}
func (v VZVirtualMachineConfiguration) Set_memoryOvercommitmentAllowed(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("set_memoryOvercommitmentAllowed:"), value)
}
func (v VZVirtualMachineConfiguration) _multiTouchDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_multiTouchDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanMultiTouchDevices reports whether the receiver responds to the private selector _multiTouchDevices.
func (v VZVirtualMachineConfiguration) CanMultiTouchDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_multiTouchDevices"))
}

// MultiTouchDevices is an exported wrapper for the private property _multiTouchDevices.
func (v VZVirtualMachineConfiguration) MultiTouchDevices() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_multiTouchDevices")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_multiTouchDevices"}
	}
	return v._multiTouchDevices(), nil
}
func (v VZVirtualMachineConfiguration) Set_multiTouchDevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_multiTouchDevices:"), value)
}
func (v VZVirtualMachineConfiguration) _panicAction() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("_panicAction"))
	return rv
}

// CanPanicAction reports whether the receiver responds to the private selector _panicAction.
func (v VZVirtualMachineConfiguration) CanPanicAction() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_panicAction"))
}

// PanicAction is an exported wrapper for the private property _panicAction.
func (v VZVirtualMachineConfiguration) PanicAction() (int64, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_panicAction")) {
		return 0, &objc.UnrecognizedSelectorError{Selector: "_panicAction"}
	}
	return v._panicAction(), nil
}
func (v VZVirtualMachineConfiguration) Set_panicAction(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("set_panicAction:"), value)
}
func (v VZVirtualMachineConfiguration) _panicDevice() IVZPanicDeviceConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_panicDevice"))
	return VZPanicDeviceConfigurationFromID(objc.ID(rv))
}

// CanPanicDevice reports whether the receiver responds to the private selector _panicDevice.
func (v VZVirtualMachineConfiguration) CanPanicDevice() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_panicDevice"))
}

// PanicDevice is an exported wrapper for the private property _panicDevice.
func (v VZVirtualMachineConfiguration) PanicDevice() (IVZPanicDeviceConfiguration, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_panicDevice")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_panicDevice"}
	}
	return v._panicDevice(), nil
}
func (v VZVirtualMachineConfiguration) Set_panicDevice(value IVZPanicDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("set_panicDevice:"), value)
}
func (v VZVirtualMachineConfiguration) _pciPassthroughDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_pciPassthroughDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanPciPassthroughDevices reports whether the receiver responds to the private selector _pciPassthroughDevices.
func (v VZVirtualMachineConfiguration) CanPciPassthroughDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_pciPassthroughDevices"))
}

// PciPassthroughDevices is an exported wrapper for the private property _pciPassthroughDevices.
func (v VZVirtualMachineConfiguration) PciPassthroughDevices() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_pciPassthroughDevices")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_pciPassthroughDevices"}
	}
	return v._pciPassthroughDevices(), nil
}
func (v VZVirtualMachineConfiguration) Set_pciPassthroughDevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_pciPassthroughDevices:"), value)
}
func (v VZVirtualMachineConfiguration) _powerSourceDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_powerSourceDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanPowerSourceDevices reports whether the receiver responds to the private selector _powerSourceDevices.
func (v VZVirtualMachineConfiguration) CanPowerSourceDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_powerSourceDevices"))
}

// PowerSourceDevices is an exported wrapper for the private property _powerSourceDevices.
func (v VZVirtualMachineConfiguration) PowerSourceDevices() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_powerSourceDevices")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_powerSourceDevices"}
	}
	return v._powerSourceDevices(), nil
}
func (v VZVirtualMachineConfiguration) Set_powerSourceDevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_powerSourceDevices:"), value)
}
func (v VZVirtualMachineConfiguration) _restartAction() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("_restartAction"))
	return rv
}

// CanRestartAction reports whether the receiver responds to the private selector _restartAction.
func (v VZVirtualMachineConfiguration) CanRestartAction() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_restartAction"))
}

// RestartAction is an exported wrapper for the private property _restartAction.
func (v VZVirtualMachineConfiguration) RestartAction() (int64, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_restartAction")) {
		return 0, &objc.UnrecognizedSelectorError{Selector: "_restartAction"}
	}
	return v._restartAction(), nil
}
func (v VZVirtualMachineConfiguration) Set_restartAction(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("set_restartAction:"), value)
}
func (v VZVirtualMachineConfiguration) _sharedRamRegions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("_sharedRamRegions"))
	return rv
}

// CanSharedRamRegions reports whether the receiver responds to the private selector _sharedRamRegions.
func (v VZVirtualMachineConfiguration) CanSharedRamRegions() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_sharedRamRegions"))
}

// SharedRamRegions is an exported wrapper for the private property _sharedRamRegions.
func (v VZVirtualMachineConfiguration) SharedRamRegions() (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_sharedRamRegions")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_sharedRamRegions"}
	}
	return v._sharedRamRegions(), nil
}
func (v VZVirtualMachineConfiguration) Set_sharedRamRegions(value unsafe.Pointer) {
	objc.Send[struct{}](v.ID, objc.Sel("set_sharedRamRegions:"), value)
}
func (v VZVirtualMachineConfiguration) _terminationUnderMemoryPressureEnabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_terminationUnderMemoryPressureEnabled"))
	return rv
}

// CanTerminationUnderMemoryPressureEnabled reports whether the receiver responds to the private selector _terminationUnderMemoryPressureEnabled.
func (v VZVirtualMachineConfiguration) CanTerminationUnderMemoryPressureEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_terminationUnderMemoryPressureEnabled"))
}

// TerminationUnderMemoryPressureEnabled is an exported wrapper for the private property _terminationUnderMemoryPressureEnabled.
func (v VZVirtualMachineConfiguration) TerminationUnderMemoryPressureEnabled() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_terminationUnderMemoryPressureEnabled")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_terminationUnderMemoryPressureEnabled"}
	}
	return v._terminationUnderMemoryPressureEnabled(), nil
}
func (v VZVirtualMachineConfiguration) Set_terminationUnderMemoryPressureEnabled(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("set_terminationUnderMemoryPressureEnabled:"), value)
}
func (v VZVirtualMachineConfiguration) _testIgnoreEntitlementChecks() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_testIgnoreEntitlementChecks"))
	return rv
}

// CanTestIgnoreEntitlementChecks reports whether the receiver responds to the private selector _testIgnoreEntitlementChecks.
func (v VZVirtualMachineConfiguration) CanTestIgnoreEntitlementChecks() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_testIgnoreEntitlementChecks"))
}

// TestIgnoreEntitlementChecks is an exported wrapper for the private property _testIgnoreEntitlementChecks.
func (v VZVirtualMachineConfiguration) TestIgnoreEntitlementChecks() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_testIgnoreEntitlementChecks")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_testIgnoreEntitlementChecks"}
	}
	return v._testIgnoreEntitlementChecks(), nil
}
func (v VZVirtualMachineConfiguration) Set_testIgnoreEntitlementChecks(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("set_testIgnoreEntitlementChecks:"), value)
}

// The number of CPUs for the virtual machine. Must be between
// minimumAllowedCPUCount and maximumAllowedCPUCount. [Full Topic]
func (v VZVirtualMachineConfiguration) CPUCount() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("CPUCount"))
	return rv
}
func (v VZVirtualMachineConfiguration) SetCPUCount(value uint) {
	objc.Send[struct{}](v.ID, objc.Sel("setCPUCount:"), value)
}

// The memory size in bytes for the virtual machine. Must be a multiple of 1MB
// and between minimumAllowedMemorySize and maximumAllowedMemorySize. [Full Topic]
func (v VZVirtualMachineConfiguration) MemorySize() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("memorySize"))
	return rv
}
func (v VZVirtualMachineConfiguration) SetMemorySize(value uint64) {
	objc.Send[struct{}](v.ID, objc.Sel("setMemorySize:"), value)
}

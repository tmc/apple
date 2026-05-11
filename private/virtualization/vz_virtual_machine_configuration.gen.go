// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

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
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration
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
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration
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
	_cpuEmulator() *VZCPUEmulatorConfiguration
	Set_cpuEmulator(value *VZCPUEmulatorConfiguration)
	_customMMIODevices() foundation.INSArray
	Set_customMMIODevices(value foundation.INSArray)
	_customVirtioDevices() foundation.INSArray
	Set_customVirtioDevices(value foundation.INSArray)
	_debugStub() *VZDebugStubConfiguration
	Set_debugStub(value *VZDebugStubConfiguration)
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
	_panicDevice() *VZPanicDeviceConfiguration
	Set_panicDevice(value *VZPanicDeviceConfiguration)
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
	_sharedRamRegions() objectivec.IObject
	Set_sharedRamRegions(value objectivec.IObject)

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

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_isDuplicateUSBDeviceConfigurationAt:usbDeviceIndex:
func (v VZVirtualMachineConfiguration) _isDuplicateUSBDeviceConfigurationAtUsbDeviceIndex(at uint64, index uint64) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_isDuplicateUSBDeviceConfigurationAt:usbDeviceIndex:"), at, index)
	return rv
}

// IsDuplicateUSBDeviceConfigurationAtUsbDeviceIndex is an exported wrapper for the private method _isDuplicateUSBDeviceConfigurationAtUsbDeviceIndex.
func (v VZVirtualMachineConfiguration) IsDuplicateUSBDeviceConfigurationAtUsbDeviceIndex(at uint64, index uint64) bool {
	return v._isDuplicateUSBDeviceConfigurationAtUsbDeviceIndex(at, index)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setAcceleratorDevices:
func (v VZVirtualMachineConfiguration) _setAcceleratorDevices(devices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setAcceleratorDevices:"), devices)
}

// SetAcceleratorDevices is an exported wrapper for the private method _setAcceleratorDevices.
func (v VZVirtualMachineConfiguration) SetAcceleratorDevices(devices objectivec.IObject) {
	v._setAcceleratorDevices(devices)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setBifrostDevices:
func (v VZVirtualMachineConfiguration) _setBifrostDevices(devices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setBifrostDevices:"), devices)
}

// SetBifrostDevices is an exported wrapper for the private method _setBifrostDevices.
func (v VZVirtualMachineConfiguration) SetBifrostDevices(devices objectivec.IObject) {
	v._setBifrostDevices(devices)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setBiometricDevices:
func (v VZVirtualMachineConfiguration) _setBiometricDevices(devices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setBiometricDevices:"), devices)
}

// SetBiometricDevices is an exported wrapper for the private method _setBiometricDevices.
func (v VZVirtualMachineConfiguration) SetBiometricDevices(devices objectivec.IObject) {
	v._setBiometricDevices(devices)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setCPUEmulator:
func (v VZVirtualMachineConfiguration) _setCPUEmulator(cPUEmulator objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setCPUEmulator:"), cPUEmulator)
}

// SetCPUEmulator is an exported wrapper for the private method _setCPUEmulator.
func (v VZVirtualMachineConfiguration) SetCPUEmulator(cPUEmulator objectivec.IObject) {
	v._setCPUEmulator(cPUEmulator)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setCoprocessors:
func (v VZVirtualMachineConfiguration) _setCoprocessors(coprocessors objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setCoprocessors:"), coprocessors)
}

// SetCoprocessors is an exported wrapper for the private method _setCoprocessors.
func (v VZVirtualMachineConfiguration) SetCoprocessors(coprocessors objectivec.IObject) {
	v._setCoprocessors(coprocessors)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setCustomMMIODevices:
func (v VZVirtualMachineConfiguration) _setCustomMMIODevices(mMIODevices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setCustomMMIODevices:"), mMIODevices)
}

// SetCustomMMIODevices is an exported wrapper for the private method _setCustomMMIODevices.
func (v VZVirtualMachineConfiguration) SetCustomMMIODevices(mMIODevices objectivec.IObject) {
	v._setCustomMMIODevices(mMIODevices)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setCustomVirtioDevices:
func (v VZVirtualMachineConfiguration) _setCustomVirtioDevices(devices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setCustomVirtioDevices:"), devices)
}

// SetCustomVirtioDevices is an exported wrapper for the private method _setCustomVirtioDevices.
func (v VZVirtualMachineConfiguration) SetCustomVirtioDevices(devices objectivec.IObject) {
	v._setCustomVirtioDevices(devices)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setDebugStub:
func (v VZVirtualMachineConfiguration) _setDebugStub(stub objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setDebugStub:"), stub)
}

// SetDebugStub is an exported wrapper for the private method _setDebugStub.
func (v VZVirtualMachineConfiguration) SetDebugStub(stub objectivec.IObject) {
	v._setDebugStub(stub)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setFatalErrorAction:
func (v VZVirtualMachineConfiguration) _setFatalErrorAction(action int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setFatalErrorAction:"), action)
}

// SetFatalErrorAction is an exported wrapper for the private method _setFatalErrorAction.
func (v VZVirtualMachineConfiguration) SetFatalErrorAction(action int64) {
	v._setFatalErrorAction(action)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setHIDDevices:
func (v VZVirtualMachineConfiguration) _setHIDDevices(hIDDevices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setHIDDevices:"), hIDDevices)
}

// SetHIDDevices is an exported wrapper for the private method _setHIDDevices.
func (v VZVirtualMachineConfiguration) SetHIDDevices(hIDDevices objectivec.IObject) {
	v._setHIDDevices(hIDDevices)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setMailboxDevices:
func (v VZVirtualMachineConfiguration) _setMailboxDevices(devices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setMailboxDevices:"), devices)
}

// SetMailboxDevices is an exported wrapper for the private method _setMailboxDevices.
func (v VZVirtualMachineConfiguration) SetMailboxDevices(devices objectivec.IObject) {
	v._setMailboxDevices(devices)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setMemoryOvercommitmentAllowed:
func (v VZVirtualMachineConfiguration) _setMemoryOvercommitmentAllowed(allowed bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setMemoryOvercommitmentAllowed:"), allowed)
}

// SetMemoryOvercommitmentAllowed is an exported wrapper for the private method _setMemoryOvercommitmentAllowed.
func (v VZVirtualMachineConfiguration) SetMemoryOvercommitmentAllowed(allowed bool) {
	v._setMemoryOvercommitmentAllowed(allowed)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setMultiTouchDevices:
func (v VZVirtualMachineConfiguration) _setMultiTouchDevices(devices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setMultiTouchDevices:"), devices)
}

// SetMultiTouchDevices is an exported wrapper for the private method _setMultiTouchDevices.
func (v VZVirtualMachineConfiguration) SetMultiTouchDevices(devices objectivec.IObject) {
	v._setMultiTouchDevices(devices)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setPCIPassthroughDevices:
func (v VZVirtualMachineConfiguration) _setPCIPassthroughDevices(devices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setPCIPassthroughDevices:"), devices)
}

// SetPCIPassthroughDevices is an exported wrapper for the private method _setPCIPassthroughDevices.
func (v VZVirtualMachineConfiguration) SetPCIPassthroughDevices(devices objectivec.IObject) {
	v._setPCIPassthroughDevices(devices)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setPanicAction:
func (v VZVirtualMachineConfiguration) _setPanicAction(action int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setPanicAction:"), action)
}

// SetPanicAction is an exported wrapper for the private method _setPanicAction.
func (v VZVirtualMachineConfiguration) SetPanicAction(action int64) {
	v._setPanicAction(action)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setPanicDevice:
func (v VZVirtualMachineConfiguration) _setPanicDevice(device objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setPanicDevice:"), device)
}

// SetPanicDevice is an exported wrapper for the private method _setPanicDevice.
func (v VZVirtualMachineConfiguration) SetPanicDevice(device objectivec.IObject) {
	v._setPanicDevice(device)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setPowerSourceDevices:
func (v VZVirtualMachineConfiguration) _setPowerSourceDevices(devices objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setPowerSourceDevices:"), devices)
}

// SetPowerSourceDevices is an exported wrapper for the private method _setPowerSourceDevices.
func (v VZVirtualMachineConfiguration) SetPowerSourceDevices(devices objectivec.IObject) {
	v._setPowerSourceDevices(devices)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setRestartAction:
func (v VZVirtualMachineConfiguration) _setRestartAction(action int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setRestartAction:"), action)
}

// SetRestartAction is an exported wrapper for the private method _setRestartAction.
func (v VZVirtualMachineConfiguration) SetRestartAction(action int64) {
	v._setRestartAction(action)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setTerminationUnderMemoryPressureEnabled:
func (v VZVirtualMachineConfiguration) _setTerminationUnderMemoryPressureEnabled(enabled bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setTerminationUnderMemoryPressureEnabled:"), enabled)
}

// SetTerminationUnderMemoryPressureEnabled is an exported wrapper for the private method _setTerminationUnderMemoryPressureEnabled.
func (v VZVirtualMachineConfiguration) SetTerminationUnderMemoryPressureEnabled(enabled bool) {
	v._setTerminationUnderMemoryPressureEnabled(enabled)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_setTestIgnoreEntitlementChecks:
func (v VZVirtualMachineConfiguration) _setTestIgnoreEntitlementChecks(checks bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setTestIgnoreEntitlementChecks:"), checks)
}

// SetTestIgnoreEntitlementChecks is an exported wrapper for the private method _setTestIgnoreEntitlementChecks.
func (v VZVirtualMachineConfiguration) SetTestIgnoreEntitlementChecks(checks bool) {
	v._setTestIgnoreEntitlementChecks(checks)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_maximumAllowedOvercommittedMemorySize
func (_VZVirtualMachineConfigurationClass VZVirtualMachineConfigurationClass) _maximumAllowedOvercommittedMemorySize() uint64 {
	rv := objc.Send[uint64](objc.ID(_VZVirtualMachineConfigurationClass.class), objc.Sel("_maximumAllowedOvercommittedMemorySize"))
	return rv
}

// MaximumAllowedOvercommittedMemorySize is an exported wrapper for the private method _maximumAllowedOvercommittedMemorySize.
func (_VZVirtualMachineConfigurationClass VZVirtualMachineConfigurationClass) MaximumAllowedOvercommittedMemorySize() uint64 {
	return _VZVirtualMachineConfigurationClass._maximumAllowedOvercommittedMemorySize()
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_acceleratorDevices
func (v VZVirtualMachineConfiguration) _acceleratorDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_acceleratorDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZVirtualMachineConfiguration) Set_acceleratorDevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_acceleratorDevices:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_bifrostDevices
func (v VZVirtualMachineConfiguration) _bifrostDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_bifrostDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZVirtualMachineConfiguration) Set_bifrostDevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_bifrostDevices:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_biometricDevices
func (v VZVirtualMachineConfiguration) _biometricDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_biometricDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZVirtualMachineConfiguration) Set_biometricDevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_biometricDevices:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_coprocessors
func (v VZVirtualMachineConfiguration) _coprocessors() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_coprocessors"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZVirtualMachineConfiguration) Set_coprocessors(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_coprocessors:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_cpuEmulator
func (v VZVirtualMachineConfiguration) _cpuEmulator() *VZCPUEmulatorConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_cpuEmulator"))
	if rv == 0 {
		return nil
	}
	val := VZCPUEmulatorConfigurationFromID(objc.ID(rv))
	return &val
}
func (v VZVirtualMachineConfiguration) Set_cpuEmulator(value *VZCPUEmulatorConfiguration) {
	if value == nil {
		objc.Send[struct{}](v.ID, objc.Sel("set_cpuEmulator:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](v.ID, objc.Sel("set_cpuEmulator:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_customMMIODevices
func (v VZVirtualMachineConfiguration) _customMMIODevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_customMMIODevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZVirtualMachineConfiguration) Set_customMMIODevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_customMMIODevices:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_customVirtioDevices
func (v VZVirtualMachineConfiguration) _customVirtioDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_customVirtioDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZVirtualMachineConfiguration) Set_customVirtioDevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_customVirtioDevices:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_debugStub
func (v VZVirtualMachineConfiguration) _debugStub() *VZDebugStubConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_debugStub"))
	if rv == 0 {
		return nil
	}
	val := VZDebugStubConfigurationFromID(objc.ID(rv))
	return &val
}
func (v VZVirtualMachineConfiguration) Set_debugStub(value *VZDebugStubConfiguration) {
	if value == nil {
		objc.Send[struct{}](v.ID, objc.Sel("set_debugStub:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](v.ID, objc.Sel("set_debugStub:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_fatalErrorAction
func (v VZVirtualMachineConfiguration) _fatalErrorAction() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("_fatalErrorAction"))
	return rv
}
func (v VZVirtualMachineConfiguration) Set_fatalErrorAction(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("set_fatalErrorAction:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_hidDevices
func (v VZVirtualMachineConfiguration) _hidDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_hidDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZVirtualMachineConfiguration) Set_hidDevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_hidDevices:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_mailboxDevices
func (v VZVirtualMachineConfiguration) _mailboxDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_mailboxDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZVirtualMachineConfiguration) Set_mailboxDevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_mailboxDevices:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_memoryOvercommitmentAllowed
func (v VZVirtualMachineConfiguration) _memoryOvercommitmentAllowed() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_memoryOvercommitmentAllowed"))
	return rv
}
func (v VZVirtualMachineConfiguration) Set_memoryOvercommitmentAllowed(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("set_memoryOvercommitmentAllowed:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_multiTouchDevices
func (v VZVirtualMachineConfiguration) _multiTouchDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_multiTouchDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZVirtualMachineConfiguration) Set_multiTouchDevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_multiTouchDevices:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_panicAction
func (v VZVirtualMachineConfiguration) _panicAction() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("_panicAction"))
	return rv
}
func (v VZVirtualMachineConfiguration) Set_panicAction(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("set_panicAction:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_panicDevice
func (v VZVirtualMachineConfiguration) _panicDevice() *VZPanicDeviceConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_panicDevice"))
	if rv == 0 {
		return nil
	}
	val := VZPanicDeviceConfigurationFromID(objc.ID(rv))
	return &val
}
func (v VZVirtualMachineConfiguration) Set_panicDevice(value *VZPanicDeviceConfiguration) {
	if value == nil {
		objc.Send[struct{}](v.ID, objc.Sel("set_panicDevice:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](v.ID, objc.Sel("set_panicDevice:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_pciPassthroughDevices
func (v VZVirtualMachineConfiguration) _pciPassthroughDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_pciPassthroughDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZVirtualMachineConfiguration) Set_pciPassthroughDevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_pciPassthroughDevices:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_powerSourceDevices
func (v VZVirtualMachineConfiguration) _powerSourceDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_powerSourceDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZVirtualMachineConfiguration) Set_powerSourceDevices(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("set_powerSourceDevices:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_restartAction
func (v VZVirtualMachineConfiguration) _restartAction() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("_restartAction"))
	return rv
}
func (v VZVirtualMachineConfiguration) Set_restartAction(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("set_restartAction:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_sharedRamRegions
func (v VZVirtualMachineConfiguration) _sharedRamRegions() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_sharedRamRegions"))
	return objectivec.Object{ID: rv}
}
func (v VZVirtualMachineConfiguration) Set_sharedRamRegions(value objectivec.IObject) {
	objc.Send[struct{}](v.ID, objc.Sel("set_sharedRamRegions:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_terminationUnderMemoryPressureEnabled
func (v VZVirtualMachineConfiguration) _terminationUnderMemoryPressureEnabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_terminationUnderMemoryPressureEnabled"))
	return rv
}
func (v VZVirtualMachineConfiguration) Set_terminationUnderMemoryPressureEnabled(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("set_terminationUnderMemoryPressureEnabled:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineConfiguration/_testIgnoreEntitlementChecks
func (v VZVirtualMachineConfiguration) _testIgnoreEntitlementChecks() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_testIgnoreEntitlementChecks"))
	return rv
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

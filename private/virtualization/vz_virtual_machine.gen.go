// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtualMachine] class.
var (
	_VZVirtualMachineClass     VZVirtualMachineClass
	_VZVirtualMachineClassOnce sync.Once
)

func getVZVirtualMachineClass() VZVirtualMachineClass {
	_VZVirtualMachineClassOnce.Do(func() {
		_VZVirtualMachineClass = VZVirtualMachineClass{class: objc.GetClass("VZVirtualMachine")}
	})
	return _VZVirtualMachineClass
}

// GetVZVirtualMachineClass returns the class object for VZVirtualMachine.
func GetVZVirtualMachineClass() VZVirtualMachineClass {
	return getVZVirtualMachineClass()
}

type VZVirtualMachineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtualMachineClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtualMachineClass) Alloc() VZVirtualMachine {
	rv := objc.Send[VZVirtualMachine](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZVirtualMachine._audioDevices]
//   - [VZVirtualMachine._canCreateCore]
//   - [VZVirtualMachine._coprocessors]
//   - [VZVirtualMachine._crashContextMessage]
//   - [VZVirtualMachine.Set_crashContextMessage]
//   - [VZVirtualMachine._createCoreWithCompletionHandler]
//   - [VZVirtualMachine._createCoresWithCompletionHandler]
//   - [VZVirtualMachine._createViewEndpointWithOptions]
//   - [VZVirtualMachine._currentConfiguration]
//   - [VZVirtualMachine._debugStub]
//   - [VZVirtualMachine._enterRestrictedModeWithCompletionHandler]
//   - [VZVirtualMachine._getUSBControllerLocationIDWithCompletionHandler]
//   - [VZVirtualMachine._hidDevices]
//   - [VZVirtualMachine._hidEventMonitor]
//   - [VZVirtualMachine._keyboards]
//   - [VZVirtualMachine._multiTouchDevices]
//   - [VZVirtualMachine._name]
//   - [VZVirtualMachine.Set_name]
//   - [VZVirtualMachine._overrideConnectionForTesting]
//   - [VZVirtualMachine._pointingDevices]
//   - [VZVirtualMachine._powerSourceDevices]
//   - [VZVirtualMachine._processHIDReportsForDeviceDeviceType]
//   - [VZVirtualMachine._resetWithTypeCompletionHandler]
//   - [VZVirtualMachine._saveMachineStateToURLOptionsCompletionHandler]
//   - [VZVirtualMachine._serialPorts]
//   - [VZVirtualMachine._serviceProcessIdentifier]
//   - [VZVirtualMachine._setCrashContextMessage]
//   - [VZVirtualMachine._setName]
//   - [VZVirtualMachine._shouldSendHIDReports]
//   - [VZVirtualMachine._stateDescription]
//   - [VZVirtualMachine._storageDevices]
//   - [VZVirtualMachine._validateRestrictedModeSupportWithError]
//   - [VZVirtualMachine.SendDigitizerEventsPointingDeviceIndex]
//   - [VZVirtualMachine.SendIOHIDEventsHidDeviceIndex]
//   - [VZVirtualMachine.SendKeyboardEventsKeyboardID]
//   - [VZVirtualMachine.SendMagnifyEventsPointingDeviceIndex]
//   - [VZVirtualMachine.SendMouseEventsPointingDeviceIndex]
//   - [VZVirtualMachine.SendMultiTouchEventsMultiTouchDeviceIndex]
//   - [VZVirtualMachine.SendPointerNSEventPointingDeviceIndex]
//   - [VZVirtualMachine.SendQuickLookEventsPointingDeviceIndex]
//   - [VZVirtualMachine.SendRotationEventsPointingDeviceIndex]
//   - [VZVirtualMachine.SendScrollWheelEventsPointingDeviceIndex]
//   - [VZVirtualMachine.SendSmartMagnifyEventsPointingDeviceIndex]
//   - [VZVirtualMachine.State]
//   - [VZVirtualMachine.SetState]
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine
type VZVirtualMachine struct {
	objectivec.Object
}

// VZVirtualMachineFromID constructs a [VZVirtualMachine] from an objc.ID.
func VZVirtualMachineFromID(id objc.ID) VZVirtualMachine {
	return VZVirtualMachine{objectivec.Object{ID: id}}
}
// Ensure VZVirtualMachine implements IVZVirtualMachine.
var _ IVZVirtualMachine = VZVirtualMachine{}

// An interface definition for the [VZVirtualMachine] class.
//
// # Methods
//
//   - [IVZVirtualMachine._audioDevices]
//   - [IVZVirtualMachine._canCreateCore]
//   - [IVZVirtualMachine._coprocessors]
//   - [IVZVirtualMachine._crashContextMessage]
//   - [IVZVirtualMachine.Set_crashContextMessage]
//   - [IVZVirtualMachine._createCoreWithCompletionHandler]
//   - [IVZVirtualMachine._createCoresWithCompletionHandler]
//   - [IVZVirtualMachine._createViewEndpointWithOptions]
//   - [IVZVirtualMachine._currentConfiguration]
//   - [IVZVirtualMachine._debugStub]
//   - [IVZVirtualMachine._enterRestrictedModeWithCompletionHandler]
//   - [IVZVirtualMachine._getUSBControllerLocationIDWithCompletionHandler]
//   - [IVZVirtualMachine._hidDevices]
//   - [IVZVirtualMachine._hidEventMonitor]
//   - [IVZVirtualMachine._keyboards]
//   - [IVZVirtualMachine._multiTouchDevices]
//   - [IVZVirtualMachine._name]
//   - [IVZVirtualMachine.Set_name]
//   - [IVZVirtualMachine._overrideConnectionForTesting]
//   - [IVZVirtualMachine._pointingDevices]
//   - [IVZVirtualMachine._powerSourceDevices]
//   - [IVZVirtualMachine._processHIDReportsForDeviceDeviceType]
//   - [IVZVirtualMachine._resetWithTypeCompletionHandler]
//   - [IVZVirtualMachine._saveMachineStateToURLOptionsCompletionHandler]
//   - [IVZVirtualMachine._serialPorts]
//   - [IVZVirtualMachine._serviceProcessIdentifier]
//   - [IVZVirtualMachine._setCrashContextMessage]
//   - [IVZVirtualMachine._setName]
//   - [IVZVirtualMachine._shouldSendHIDReports]
//   - [IVZVirtualMachine._stateDescription]
//   - [IVZVirtualMachine._storageDevices]
//   - [IVZVirtualMachine._validateRestrictedModeSupportWithError]
//   - [IVZVirtualMachine.SendDigitizerEventsPointingDeviceIndex]
//   - [IVZVirtualMachine.SendIOHIDEventsHidDeviceIndex]
//   - [IVZVirtualMachine.SendKeyboardEventsKeyboardID]
//   - [IVZVirtualMachine.SendMagnifyEventsPointingDeviceIndex]
//   - [IVZVirtualMachine.SendMouseEventsPointingDeviceIndex]
//   - [IVZVirtualMachine.SendMultiTouchEventsMultiTouchDeviceIndex]
//   - [IVZVirtualMachine.SendPointerNSEventPointingDeviceIndex]
//   - [IVZVirtualMachine.SendQuickLookEventsPointingDeviceIndex]
//   - [IVZVirtualMachine.SendRotationEventsPointingDeviceIndex]
//   - [IVZVirtualMachine.SendScrollWheelEventsPointingDeviceIndex]
//   - [IVZVirtualMachine.SendSmartMagnifyEventsPointingDeviceIndex]
//   - [IVZVirtualMachine.State]
//   - [IVZVirtualMachine.SetState]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine
type IVZVirtualMachine interface {
	objectivec.IObject

	// Topic: Methods

	_audioDevices() foundation.INSArray
	_canCreateCore() bool
	_coprocessors() foundation.INSArray
	_crashContextMessage() string
	Set_crashContextMessage(value string)
	_createCoreWithCompletionHandler(handler ErrorHandler)
	_createCoresWithCompletionHandler(handler ErrorHandler)
	_createViewEndpointWithOptions(options uint64) objectivec.IObject
	_currentConfiguration() IVZVirtualMachineConfiguration
	_debugStub() *VZDebugStub
	_enterRestrictedModeWithCompletionHandler(handler ErrorHandler)
	_getUSBControllerLocationIDWithCompletionHandler(handler ErrorHandler)
	_hidDevices() foundation.INSArray
	_hidEventMonitor() *VZHIDEventMonitor
	_keyboards() foundation.INSArray
	_multiTouchDevices() foundation.INSArray
	_name() string
	Set_name(value string)
	_overrideConnectionForTesting(testing objectivec.IObject)
	_pointingDevices() foundation.INSArray
	_powerSourceDevices() foundation.INSArray
	_processHIDReportsForDeviceDeviceType(hIDReports unsafe.Pointer, device uint32, type_ int)
	_resetWithTypeCompletionHandler(type_ int64, handler ErrorHandler)
	_saveMachineStateToURLOptionsCompletionHandler(url foundation.INSURL, options objectivec.IObject, handler ErrorHandler)
	_serialPorts() foundation.INSArray
	_serviceProcessIdentifier() int
	_setCrashContextMessage(message objectivec.IObject)
	_setName(name objectivec.IObject)
	_shouldSendHIDReports() bool
	_stateDescription() string
	_storageDevices() foundation.INSArray
	_validateRestrictedModeSupportWithError() (bool, error)
	SendDigitizerEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)
	SendIOHIDEventsHidDeviceIndex(iOHIDEvents unsafe.Pointer, index uint32)
	SendKeyboardEventsKeyboardID(events unsafe.Pointer, id uint32)
	SendMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)
	SendMouseEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)
	SendMultiTouchEventsMultiTouchDeviceIndex(events unsafe.Pointer, index uint32)
	SendPointerNSEventPointingDeviceIndex(nSEvent objectivec.IObject, index uint32)
	SendQuickLookEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)
	SendRotationEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)
	SendScrollWheelEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)
	SendSmartMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)
	State() int64
	SetState(value int64)
}

// Init initializes the instance.
func (v VZVirtualMachine) Init() VZVirtualMachine {
	rv := objc.Send[VZVirtualMachine](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtualMachine) Autorelease() VZVirtualMachine {
	rv := objc.Send[VZVirtualMachine](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachine creates a new VZVirtualMachine instance.
func NewVZVirtualMachine() VZVirtualMachine {
	class := getVZVirtualMachineClass()
	rv := objc.Send[VZVirtualMachine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_createCoreWithCompletionHandler:
func (v VZVirtualMachine) _createCoreWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](v.ID, objc.Sel("_createCoreWithCompletionHandler:"), _block0)
}

// CreateCoreWithCompletionHandler is an exported wrapper for the private method _createCoreWithCompletionHandler.
func (v VZVirtualMachine) CreateCoreWithCompletionHandler(handler ErrorHandler) {
	v._createCoreWithCompletionHandler(handler)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_createCoresWithCompletionHandler:
func (v VZVirtualMachine) _createCoresWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](v.ID, objc.Sel("_createCoresWithCompletionHandler:"), _block0)
}

// CreateCoresWithCompletionHandler is an exported wrapper for the private method _createCoresWithCompletionHandler.
func (v VZVirtualMachine) CreateCoresWithCompletionHandler(handler ErrorHandler) {
	v._createCoresWithCompletionHandler(handler)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_createViewEndpointWithOptions:
func (v VZVirtualMachine) _createViewEndpointWithOptions(options uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_createViewEndpointWithOptions:"), options)
	return objectivec.Object{ID: rv}
}

// CreateViewEndpointWithOptions is an exported wrapper for the private method _createViewEndpointWithOptions.
func (v VZVirtualMachine) CreateViewEndpointWithOptions(options uint64) objectivec.IObject {
	return v._createViewEndpointWithOptions(options)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_enterRestrictedModeWithCompletionHandler:
func (v VZVirtualMachine) _enterRestrictedModeWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](v.ID, objc.Sel("_enterRestrictedModeWithCompletionHandler:"), _block0)
}

// EnterRestrictedModeWithCompletionHandler is an exported wrapper for the private method _enterRestrictedModeWithCompletionHandler.
func (v VZVirtualMachine) EnterRestrictedModeWithCompletionHandler(handler ErrorHandler) {
	v._enterRestrictedModeWithCompletionHandler(handler)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_getUSBControllerLocationIDWithCompletionHandler:
func (v VZVirtualMachine) _getUSBControllerLocationIDWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](v.ID, objc.Sel("_getUSBControllerLocationIDWithCompletionHandler:"), _block0)
}

// GetUSBControllerLocationIDWithCompletionHandler is an exported wrapper for the private method _getUSBControllerLocationIDWithCompletionHandler.
func (v VZVirtualMachine) GetUSBControllerLocationIDWithCompletionHandler(handler ErrorHandler) {
	v._getUSBControllerLocationIDWithCompletionHandler(handler)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_overrideConnectionForTesting:
func (v VZVirtualMachine) _overrideConnectionForTesting(testing objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_overrideConnectionForTesting:"), testing)
}

// OverrideConnectionForTesting is an exported wrapper for the private method _overrideConnectionForTesting.
func (v VZVirtualMachine) OverrideConnectionForTesting(testing objectivec.IObject) {
	v._overrideConnectionForTesting(testing)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_processHIDReports:forDevice:deviceType:
func (v VZVirtualMachine) _processHIDReportsForDeviceDeviceType(hIDReports unsafe.Pointer, device uint32, type_ int) {
	objc.Send[objc.ID](v.ID, objc.Sel("_processHIDReports:forDevice:deviceType:"), hIDReports, device, type_)
}

// ProcessHIDReportsForDeviceDeviceType is an exported wrapper for the private method _processHIDReportsForDeviceDeviceType.
func (v VZVirtualMachine) ProcessHIDReportsForDeviceDeviceType(hIDReports unsafe.Pointer, device uint32, type_ int) {
	v._processHIDReportsForDeviceDeviceType(hIDReports, device, type_)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_resetWithType:completionHandler:
func (v VZVirtualMachine) _resetWithTypeCompletionHandler(type_ int64, handler ErrorHandler) {
_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](v.ID, objc.Sel("_resetWithType:completionHandler:"), type_, _block1)
}

// ResetWithTypeCompletionHandler is an exported wrapper for the private method _resetWithTypeCompletionHandler.
func (v VZVirtualMachine) ResetWithTypeCompletionHandler(type_ int64, handler ErrorHandler) {
	v._resetWithTypeCompletionHandler(type_, handler)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_saveMachineStateToURL:options:completionHandler:
func (v VZVirtualMachine) _saveMachineStateToURLOptionsCompletionHandler(url foundation.INSURL, options objectivec.IObject, handler ErrorHandler) {
_block2, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](v.ID, objc.Sel("_saveMachineStateToURL:options:completionHandler:"), url, options, _block2)
}

// SaveMachineStateToURLOptionsCompletionHandler is an exported wrapper for the private method _saveMachineStateToURLOptionsCompletionHandler.
func (v VZVirtualMachine) SaveMachineStateToURLOptionsCompletionHandler(url foundation.INSURL, options objectivec.IObject, handler ErrorHandler) {
	v._saveMachineStateToURLOptionsCompletionHandler(url, options, handler)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_setCrashContextMessage:
func (v VZVirtualMachine) _setCrashContextMessage(message objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setCrashContextMessage:"), message)
}

// SetCrashContextMessage is an exported wrapper for the private method _setCrashContextMessage.
func (v VZVirtualMachine) SetCrashContextMessage(message objectivec.IObject) {
	v._setCrashContextMessage(message)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_setName:
func (v VZVirtualMachine) _setName(name objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setName:"), name)
}

// SetName is an exported wrapper for the private method _setName.
func (v VZVirtualMachine) SetName(name objectivec.IObject) {
	v._setName(name)
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_shouldSendHIDReports
func (v VZVirtualMachine) _shouldSendHIDReports() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_shouldSendHIDReports"))
	return rv
}

// ShouldSendHIDReports is an exported wrapper for the private method _shouldSendHIDReports.
func (v VZVirtualMachine) ShouldSendHIDReports() bool {
	return v._shouldSendHIDReports()
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_validateRestrictedModeSupportWithError:
func (v VZVirtualMachine) _validateRestrictedModeSupportWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("_validateRestrictedModeSupportWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_validateRestrictedModeSupportWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// ValidateRestrictedModeSupportWithError is an exported wrapper for the private method _validateRestrictedModeSupportWithError.
func (v VZVirtualMachine) ValidateRestrictedModeSupportWithError() (bool, error) {
	return v._validateRestrictedModeSupportWithError()
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/sendDigitizerEvents:pointingDeviceIndex:
func (v VZVirtualMachine) SendDigitizerEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendDigitizerEvents:pointingDeviceIndex:"), events, index)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/sendIOHIDEvents:hidDeviceIndex:
func (v VZVirtualMachine) SendIOHIDEventsHidDeviceIndex(iOHIDEvents unsafe.Pointer, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendIOHIDEvents:hidDeviceIndex:"), iOHIDEvents, index)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/sendKeyboardEvents:keyboardID:
func (v VZVirtualMachine) SendKeyboardEventsKeyboardID(events unsafe.Pointer, id uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendKeyboardEvents:keyboardID:"), events, id)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/sendMagnifyEvents:pointingDeviceIndex:
func (v VZVirtualMachine) SendMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendMagnifyEvents:pointingDeviceIndex:"), events, index)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/sendMouseEvents:pointingDeviceIndex:
func (v VZVirtualMachine) SendMouseEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendMouseEvents:pointingDeviceIndex:"), events, index)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/sendMultiTouchEvents:multiTouchDeviceIndex:
func (v VZVirtualMachine) SendMultiTouchEventsMultiTouchDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendMultiTouchEvents:multiTouchDeviceIndex:"), events, index)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/sendPointerNSEvent:pointingDeviceIndex:
func (v VZVirtualMachine) SendPointerNSEventPointingDeviceIndex(nSEvent objectivec.IObject, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendPointerNSEvent:pointingDeviceIndex:"), nSEvent, index)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/sendQuickLookEvents:pointingDeviceIndex:
func (v VZVirtualMachine) SendQuickLookEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendQuickLookEvents:pointingDeviceIndex:"), events, index)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/sendRotationEvents:pointingDeviceIndex:
func (v VZVirtualMachine) SendRotationEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendRotationEvents:pointingDeviceIndex:"), events, index)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/sendScrollWheelEvents:pointingDeviceIndex:
func (v VZVirtualMachine) SendScrollWheelEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendScrollWheelEvents:pointingDeviceIndex:"), events, index)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/sendSmartMagnifyEvents:pointingDeviceIndex:
func (v VZVirtualMachine) SendSmartMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendSmartMagnifyEvents:pointingDeviceIndex:"), events, index)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_audioDevices
func (v VZVirtualMachine) _audioDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_audioDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_canCreateCore
func (v VZVirtualMachine) _canCreateCore() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_canCreateCore"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_coprocessors
func (v VZVirtualMachine) _coprocessors() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_coprocessors"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_crashContextMessage
func (v VZVirtualMachine) _crashContextMessage() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_crashContextMessage"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZVirtualMachine) Set_crashContextMessage(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("set_crashContextMessage:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_currentConfiguration
func (v VZVirtualMachine) _currentConfiguration() IVZVirtualMachineConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_currentConfiguration"))
	return VZVirtualMachineConfigurationFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_debugStub
func (v VZVirtualMachine) _debugStub() *VZDebugStub {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_debugStub"))
	if rv == 0 {
		return nil
	}
	val := VZDebugStubFromID(objc.ID(rv))
	return &val
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_hidDevices
func (v VZVirtualMachine) _hidDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_hidDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_hidEventMonitor
func (v VZVirtualMachine) _hidEventMonitor() *VZHIDEventMonitor {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_hidEventMonitor"))
	if rv == 0 {
		return nil
	}
	val := VZHIDEventMonitorFromID(objc.ID(rv))
	return &val
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_keyboards
func (v VZVirtualMachine) _keyboards() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_keyboards"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_multiTouchDevices
func (v VZVirtualMachine) _multiTouchDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_multiTouchDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_name
func (v VZVirtualMachine) _name() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_name"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZVirtualMachine) Set_name(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("set_name:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_pointingDevices
func (v VZVirtualMachine) _pointingDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_pointingDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_powerSourceDevices
func (v VZVirtualMachine) _powerSourceDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_powerSourceDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_serialPorts
func (v VZVirtualMachine) _serialPorts() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_serialPorts"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_serviceProcessIdentifier
func (v VZVirtualMachine) _serviceProcessIdentifier() int {
	rv := objc.Send[int](v.ID, objc.Sel("_serviceProcessIdentifier"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_stateDescription
func (v VZVirtualMachine) _stateDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_stateDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/_storageDevices
func (v VZVirtualMachine) _storageDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_storageDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/state
func (v VZVirtualMachine) State() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("state"))
	return rv
}
func (v VZVirtualMachine) SetState(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("setState:"), value)
}

// _createCore is a synchronous wrapper around [VZVirtualMachine._createCoreWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZVirtualMachine) _createCore(ctx context.Context) error {
	done := make(chan error, 1)
	v._createCoreWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// _createCores is a synchronous wrapper around [VZVirtualMachine._createCoresWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZVirtualMachine) _createCores(ctx context.Context) error {
	done := make(chan error, 1)
	v._createCoresWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// _enterRestrictedMode is a synchronous wrapper around [VZVirtualMachine._enterRestrictedModeWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZVirtualMachine) _enterRestrictedMode(ctx context.Context) error {
	done := make(chan error, 1)
	v._enterRestrictedModeWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// _getUSBControllerLocationID is a synchronous wrapper around [VZVirtualMachine._getUSBControllerLocationIDWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZVirtualMachine) _getUSBControllerLocationID(ctx context.Context) error {
	done := make(chan error, 1)
	v._getUSBControllerLocationIDWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// _resetWithType is a synchronous wrapper around [VZVirtualMachine._resetWithTypeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZVirtualMachine) _resetWithType(ctx context.Context, type_ int64) error {
	done := make(chan error, 1)
	v._resetWithTypeCompletionHandler(type_, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// _saveMachineStateToURLOptions is a synchronous wrapper around [VZVirtualMachine._saveMachineStateToURLOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZVirtualMachine) _saveMachineStateToURLOptions(ctx context.Context, url foundation.INSURL, options objectivec.IObject) error {
	done := make(chan error, 1)
	v._saveMachineStateToURLOptionsCompletionHandler(url, options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}


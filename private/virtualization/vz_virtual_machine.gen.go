// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[VZVirtualMachine](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVirtualMachine._audioDevices]
//   - [VZVirtualMachine._canCreateCore]
//   - [VZVirtualMachine._coprocessors]
//   - [VZVirtualMachine._crashContextMessage]
//   - [VZVirtualMachine.Set_crashContextMessage]
//   - [VZVirtualMachine._createCoreWithCompletionHandler]
//   - [VZVirtualMachine._createCoresWithCompletionHandler]
//   - [VZVirtualMachine._createSharedMemoryCoresWithOptionsCompletionHandler]
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
//   - [IVZVirtualMachine._createSharedMemoryCoresWithOptionsCompletionHandler]
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
	_createSharedMemoryCoresWithOptionsCompletionHandler(options unsafe.Pointer, handler ErrorHandler)
	_createViewEndpointWithOptions(options uint64) objectivec.IObject
	_currentConfiguration() IVZVirtualMachineConfiguration
	_debugStub() IVZDebugStub
	_enterRestrictedModeWithCompletionHandler(handler ErrorHandler)
	_getUSBControllerLocationIDWithCompletionHandler(handler ErrorHandler)
	_hidDevices() foundation.INSArray
	_hidEventMonitor() IVZHIDEventMonitor
	_keyboards() foundation.INSArray
	_multiTouchDevices() foundation.INSArray
	_name() string
	Set_name(value string)
	_overrideConnectionForTesting(testing objectivec.IObject)
	_pointingDevices() foundation.INSArray
	_powerSourceDevices() foundation.INSArray
	_processHIDReportsForDeviceDeviceType(hIDReports VZOpaqueHIDReports, device uint32, type_ int32)
	_resetWithTypeCompletionHandler(type_ int64, handler ErrorHandler)
	_saveMachineStateToURLOptionsCompletionHandler(url foundation.NSURL, options objectivec.IObject, handler ErrorHandler)
	_serialPorts() foundation.INSArray
	_serviceProcessIdentifier() int
	_setCrashContextMessage(message objectivec.IObject)
	_setName(name objectivec.IObject)
	_shouldSendHIDReports() bool
	_stateDescription() string
	_storageDevices() foundation.INSArray
	_validateRestrictedModeSupportWithError() (bool, error)
	SendDigitizerEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)
	SendIOHIDEventsHidDeviceIndex(iOHIDEvents VZOpaqueIOHIDEvents, index uint32)
	SendKeyboardEventsKeyboardID(events VZOpaqueKeyboardEvents, id uint32)
	SendMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)
	SendMouseEventsPointingDeviceIndex(events VZOpaqueMouseEvents, index uint32)
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
	rv := objc.SendIfResponds[VZVirtualMachine](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtualMachine) Autorelease() VZVirtualMachine {
	rv := objc.SendIfResponds[VZVirtualMachine](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachine creates a new VZVirtualMachine instance.
func NewVZVirtualMachine() VZVirtualMachine {
	class := getVZVirtualMachineClass()
	rv := objc.SendIfResponds[VZVirtualMachine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZVirtualMachine) _createCoreWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_createCoreWithCompletionHandler:"), _block0)
}

// CreateCoreWithCompletionHandler is an exported wrapper for the private method _createCoreWithCompletionHandler.
func (v VZVirtualMachine) CreateCoreWithCompletionHandler(handler ErrorHandler) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_createCoreWithCompletionHandler:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_createCoreWithCompletionHandler:"}
		return err
	}
	v._createCoreWithCompletionHandler(handler)
	return nil
}

// CanCreateCoreWithCompletionHandler reports whether the receiver responds to the private selector _createCoreWithCompletionHandler:.
func (v VZVirtualMachine) CanCreateCoreWithCompletionHandler() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_createCoreWithCompletionHandler:"))
}
func (v VZVirtualMachine) _createCoresWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_createCoresWithCompletionHandler:"), _block0)
}

// CreateCoresWithCompletionHandler is an exported wrapper for the private method _createCoresWithCompletionHandler.
func (v VZVirtualMachine) CreateCoresWithCompletionHandler(handler ErrorHandler) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_createCoresWithCompletionHandler:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_createCoresWithCompletionHandler:"}
		return err
	}
	v._createCoresWithCompletionHandler(handler)
	return nil
}

// CanCreateCoresWithCompletionHandler reports whether the receiver responds to the private selector _createCoresWithCompletionHandler:.
func (v VZVirtualMachine) CanCreateCoresWithCompletionHandler() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_createCoresWithCompletionHandler:"))
}
func (v VZVirtualMachine) _createSharedMemoryCoresWithOptionsCompletionHandler(options unsafe.Pointer, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_createSharedMemoryCoresWithOptions:completionHandler:"), options, _block1)
}

// CreateSharedMemoryCoresWithOptionsCompletionHandler is an exported wrapper for the private method _createSharedMemoryCoresWithOptionsCompletionHandler.
func (v VZVirtualMachine) CreateSharedMemoryCoresWithOptionsCompletionHandler(options unsafe.Pointer, handler ErrorHandler) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_createSharedMemoryCoresWithOptions:completionHandler:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_createSharedMemoryCoresWithOptions:completionHandler:"}
		return err
	}
	v._createSharedMemoryCoresWithOptionsCompletionHandler(options, handler)
	return nil
}

// CanCreateSharedMemoryCoresWithOptionsCompletionHandler reports whether the receiver responds to the private selector _createSharedMemoryCoresWithOptions:completionHandler:.
func (v VZVirtualMachine) CanCreateSharedMemoryCoresWithOptionsCompletionHandler() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_createSharedMemoryCoresWithOptions:completionHandler:"))
}
func (v VZVirtualMachine) _createViewEndpointWithOptions(options uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_createViewEndpointWithOptions:"), options)
	return objectivec.Object{ID: rv}
}

// CreateViewEndpointWithOptions is an exported wrapper for the private method _createViewEndpointWithOptions.
func (v VZVirtualMachine) CreateViewEndpointWithOptions(options uint64) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_createViewEndpointWithOptions:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_createViewEndpointWithOptions:"}
		return nil, err
	}
	return v._createViewEndpointWithOptions(options), nil
}

// CanCreateViewEndpointWithOptions reports whether the receiver responds to the private selector _createViewEndpointWithOptions:.
func (v VZVirtualMachine) CanCreateViewEndpointWithOptions() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_createViewEndpointWithOptions:"))
}
func (v VZVirtualMachine) _enterRestrictedModeWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_enterRestrictedModeWithCompletionHandler:"), _block0)
}

// EnterRestrictedModeWithCompletionHandler is an exported wrapper for the private method _enterRestrictedModeWithCompletionHandler.
func (v VZVirtualMachine) EnterRestrictedModeWithCompletionHandler(handler ErrorHandler) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_enterRestrictedModeWithCompletionHandler:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_enterRestrictedModeWithCompletionHandler:"}
		return err
	}
	v._enterRestrictedModeWithCompletionHandler(handler)
	return nil
}

// CanEnterRestrictedModeWithCompletionHandler reports whether the receiver responds to the private selector _enterRestrictedModeWithCompletionHandler:.
func (v VZVirtualMachine) CanEnterRestrictedModeWithCompletionHandler() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_enterRestrictedModeWithCompletionHandler:"))
}
func (v VZVirtualMachine) _getUSBControllerLocationIDWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_getUSBControllerLocationIDWithCompletionHandler:"), _block0)
}

// GetUSBControllerLocationIDWithCompletionHandler is an exported wrapper for the private method _getUSBControllerLocationIDWithCompletionHandler.
func (v VZVirtualMachine) GetUSBControllerLocationIDWithCompletionHandler(handler ErrorHandler) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_getUSBControllerLocationIDWithCompletionHandler:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_getUSBControllerLocationIDWithCompletionHandler:"}
		return err
	}
	v._getUSBControllerLocationIDWithCompletionHandler(handler)
	return nil
}

// CanGetUSBControllerLocationIDWithCompletionHandler reports whether the receiver responds to the private selector _getUSBControllerLocationIDWithCompletionHandler:.
func (v VZVirtualMachine) CanGetUSBControllerLocationIDWithCompletionHandler() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_getUSBControllerLocationIDWithCompletionHandler:"))
}
func (v VZVirtualMachine) _overrideConnectionForTesting(testing objectivec.IObject) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_overrideConnectionForTesting:"), testing)
}

// OverrideConnectionForTesting is an exported wrapper for the private method _overrideConnectionForTesting.
func (v VZVirtualMachine) OverrideConnectionForTesting(testing objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_overrideConnectionForTesting:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_overrideConnectionForTesting:"}
		return err
	}
	v._overrideConnectionForTesting(testing)
	return nil
}

// CanOverrideConnectionForTesting reports whether the receiver responds to the private selector _overrideConnectionForTesting:.
func (v VZVirtualMachine) CanOverrideConnectionForTesting() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_overrideConnectionForTesting:"))
}
func (v VZVirtualMachine) _processHIDReportsForDeviceDeviceType(hIDReports VZOpaqueHIDReports, device uint32, type_ int32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_processHIDReports:forDevice:deviceType:"), hIDReports.UnsafePointer(), device, type_)
}

// ProcessHIDReportsForDeviceDeviceType is an exported wrapper for the private method _processHIDReportsForDeviceDeviceType.
func (v VZVirtualMachine) ProcessHIDReportsForDeviceDeviceType(hIDReports VZOpaqueHIDReports, device uint32, type_ int32) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_processHIDReports:forDevice:deviceType:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_processHIDReports:forDevice:deviceType:"}
		return err
	}
	v._processHIDReportsForDeviceDeviceType(hIDReports, device, type_)
	return nil
}

// CanProcessHIDReportsForDeviceDeviceType reports whether the receiver responds to the private selector _processHIDReports:forDevice:deviceType:.
func (v VZVirtualMachine) CanProcessHIDReportsForDeviceDeviceType() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_processHIDReports:forDevice:deviceType:"))
}
func (v VZVirtualMachine) _resetWithTypeCompletionHandler(type_ int64, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_resetWithType:completionHandler:"), type_, _block1)
}

// ResetWithTypeCompletionHandler is an exported wrapper for the private method _resetWithTypeCompletionHandler.
func (v VZVirtualMachine) ResetWithTypeCompletionHandler(type_ int64, handler ErrorHandler) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_resetWithType:completionHandler:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_resetWithType:completionHandler:"}
		return err
	}
	v._resetWithTypeCompletionHandler(type_, handler)
	return nil
}

// CanResetWithTypeCompletionHandler reports whether the receiver responds to the private selector _resetWithType:completionHandler:.
func (v VZVirtualMachine) CanResetWithTypeCompletionHandler() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_resetWithType:completionHandler:"))
}
func (v VZVirtualMachine) _saveMachineStateToURLOptionsCompletionHandler(url foundation.NSURL, options objectivec.IObject, handler ErrorHandler) {
	_block2, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_saveMachineStateToURL:options:completionHandler:"), url, options, _block2)
}

// SaveMachineStateToURLOptionsCompletionHandler is an exported wrapper for the private method _saveMachineStateToURLOptionsCompletionHandler.
func (v VZVirtualMachine) SaveMachineStateToURLOptionsCompletionHandler(url foundation.NSURL, options objectivec.IObject, handler ErrorHandler) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_saveMachineStateToURL:options:completionHandler:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_saveMachineStateToURL:options:completionHandler:"}
		return err
	}
	v._saveMachineStateToURLOptionsCompletionHandler(url, options, handler)
	return nil
}

// CanSaveMachineStateToURLOptionsCompletionHandler reports whether the receiver responds to the private selector _saveMachineStateToURL:options:completionHandler:.
func (v VZVirtualMachine) CanSaveMachineStateToURLOptionsCompletionHandler() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_saveMachineStateToURL:options:completionHandler:"))
}
func (v VZVirtualMachine) _setCrashContextMessage(message objectivec.IObject) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_setCrashContextMessage:"), message)
}

// SetCrashContextMessage is an exported wrapper for the private method _setCrashContextMessage.
func (v VZVirtualMachine) SetCrashContextMessage(message objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setCrashContextMessage:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setCrashContextMessage:"}
		return err
	}
	v._setCrashContextMessage(message)
	return nil
}

// CanSetCrashContextMessage reports whether the receiver responds to the private selector _setCrashContextMessage:.
func (v VZVirtualMachine) CanSetCrashContextMessage() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setCrashContextMessage:"))
}
func (v VZVirtualMachine) _setName(name objectivec.IObject) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_setName:"), name)
}

// SetName is an exported wrapper for the private method _setName.
func (v VZVirtualMachine) SetName(name objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setName:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setName:"}
		return err
	}
	v._setName(name)
	return nil
}

// CanSetName reports whether the receiver responds to the private selector _setName:.
func (v VZVirtualMachine) CanSetName() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setName:"))
}
func (v VZVirtualMachine) _shouldSendHIDReports() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("_shouldSendHIDReports"))
	return rv
}

// ShouldSendHIDReports is an exported wrapper for the private method _shouldSendHIDReports.
func (v VZVirtualMachine) ShouldSendHIDReports() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_shouldSendHIDReports")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_shouldSendHIDReports"}
		return false, err
	}
	return v._shouldSendHIDReports(), nil
}

// CanShouldSendHIDReports reports whether the receiver responds to the private selector _shouldSendHIDReports.
func (v VZVirtualMachine) CanShouldSendHIDReports() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_shouldSendHIDReports"))
}
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
	if !objc.RespondsToSelector(v.ID, objc.Sel("_validateRestrictedModeSupportWithError:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_validateRestrictedModeSupportWithError:"}
		return false, err
	}
	return v._validateRestrictedModeSupportWithError()
}

// CanValidateRestrictedModeSupportWithError reports whether the receiver responds to the private selector _validateRestrictedModeSupportWithError:.
func (v VZVirtualMachine) CanValidateRestrictedModeSupportWithError() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_validateRestrictedModeSupportWithError:"))
}
func (v VZVirtualMachine) SendDigitizerEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendDigitizerEvents:pointingDeviceIndex:"), events, index)
}
func (v VZVirtualMachine) SendIOHIDEventsHidDeviceIndex(iOHIDEvents VZOpaqueIOHIDEvents, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendIOHIDEvents:hidDeviceIndex:"), iOHIDEvents.UnsafePointer(), index)
}
func (v VZVirtualMachine) SendKeyboardEventsKeyboardID(events VZOpaqueKeyboardEvents, id uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendKeyboardEvents:keyboardID:"), events.UnsafePointer(), id)
}
func (v VZVirtualMachine) SendMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendMagnifyEvents:pointingDeviceIndex:"), events, index)
}
func (v VZVirtualMachine) SendMouseEventsPointingDeviceIndex(events VZOpaqueMouseEvents, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendMouseEvents:pointingDeviceIndex:"), events.UnsafePointer(), index)
}
func (v VZVirtualMachine) SendMultiTouchEventsMultiTouchDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendMultiTouchEvents:multiTouchDeviceIndex:"), events, index)
}
func (v VZVirtualMachine) SendPointerNSEventPointingDeviceIndex(nSEvent objectivec.IObject, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendPointerNSEvent:pointingDeviceIndex:"), nSEvent, index)
}
func (v VZVirtualMachine) SendQuickLookEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendQuickLookEvents:pointingDeviceIndex:"), events, index)
}
func (v VZVirtualMachine) SendRotationEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendRotationEvents:pointingDeviceIndex:"), events, index)
}
func (v VZVirtualMachine) SendScrollWheelEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendScrollWheelEvents:pointingDeviceIndex:"), events, index)
}
func (v VZVirtualMachine) SendSmartMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendSmartMagnifyEvents:pointingDeviceIndex:"), events, index)
}

func (v VZVirtualMachine) _audioDevices() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_audioDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanAudioDevices reports whether the receiver responds to the private selector _audioDevices.
func (v VZVirtualMachine) CanAudioDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_audioDevices"))
}

// AudioDevices is an exported wrapper for the private property _audioDevices.
func (v VZVirtualMachine) AudioDevices() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_audioDevices")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_audioDevices"}
	}
	return v._audioDevices(), nil
}
func (v VZVirtualMachine) _canCreateCore() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("_canCreateCore"))
	return rv
}

// CanCanCreateCore reports whether the receiver responds to the private selector _canCreateCore.
func (v VZVirtualMachine) CanCanCreateCore() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_canCreateCore"))
}

// CanCreateCore is an exported wrapper for the private property _canCreateCore.
func (v VZVirtualMachine) CanCreateCore() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_canCreateCore")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_canCreateCore"}
	}
	return v._canCreateCore(), nil
}
func (v VZVirtualMachine) _coprocessors() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_coprocessors"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanCoprocessors reports whether the receiver responds to the private selector _coprocessors.
func (v VZVirtualMachine) CanCoprocessors() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_coprocessors"))
}

// Coprocessors is an exported wrapper for the private property _coprocessors.
func (v VZVirtualMachine) Coprocessors() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_coprocessors")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_coprocessors"}
	}
	return v._coprocessors(), nil
}
func (v VZVirtualMachine) _crashContextMessage() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_crashContextMessage"))
	return foundation.NSStringFromID(rv).String()
}

// CanCrashContextMessage reports whether the receiver responds to the private selector _crashContextMessage.
func (v VZVirtualMachine) CanCrashContextMessage() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_crashContextMessage"))
}

// CrashContextMessage is an exported wrapper for the private property _crashContextMessage.
func (v VZVirtualMachine) CrashContextMessage() (string, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_crashContextMessage")) {
		return "", &objc.UnrecognizedSelectorError{Selector: "_crashContextMessage"}
	}
	return v._crashContextMessage(), nil
}
func (v VZVirtualMachine) Set_crashContextMessage(value string) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("set_crashContextMessage:"), objc.String(value))
}
func (v VZVirtualMachine) _currentConfiguration() IVZVirtualMachineConfiguration {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_currentConfiguration"))
	return VZVirtualMachineConfigurationFromID(objc.ID(rv))
}

// CanCurrentConfiguration reports whether the receiver responds to the private selector _currentConfiguration.
func (v VZVirtualMachine) CanCurrentConfiguration() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_currentConfiguration"))
}

// CurrentConfiguration is an exported wrapper for the private property _currentConfiguration.
func (v VZVirtualMachine) CurrentConfiguration() (IVZVirtualMachineConfiguration, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_currentConfiguration")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_currentConfiguration"}
	}
	return v._currentConfiguration(), nil
}
func (v VZVirtualMachine) _debugStub() IVZDebugStub {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_debugStub"))
	return VZDebugStubFromID(objc.ID(rv))
}

// CanDebugStub reports whether the receiver responds to the private selector _debugStub.
func (v VZVirtualMachine) CanDebugStub() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_debugStub"))
}

// DebugStub is an exported wrapper for the private property _debugStub.
func (v VZVirtualMachine) DebugStub() (IVZDebugStub, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_debugStub")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_debugStub"}
	}
	return v._debugStub(), nil
}
func (v VZVirtualMachine) _hidDevices() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_hidDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanHidDevices reports whether the receiver responds to the private selector _hidDevices.
func (v VZVirtualMachine) CanHidDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_hidDevices"))
}

// HidDevices is an exported wrapper for the private property _hidDevices.
func (v VZVirtualMachine) HidDevices() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_hidDevices")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_hidDevices"}
	}
	return v._hidDevices(), nil
}
func (v VZVirtualMachine) _hidEventMonitor() IVZHIDEventMonitor {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_hidEventMonitor"))
	return VZHIDEventMonitorFromID(objc.ID(rv))
}

// CanHidEventMonitor reports whether the receiver responds to the private selector _hidEventMonitor.
func (v VZVirtualMachine) CanHidEventMonitor() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_hidEventMonitor"))
}

// HidEventMonitor is an exported wrapper for the private property _hidEventMonitor.
func (v VZVirtualMachine) HidEventMonitor() (IVZHIDEventMonitor, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_hidEventMonitor")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_hidEventMonitor"}
	}
	return v._hidEventMonitor(), nil
}
func (v VZVirtualMachine) _keyboards() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_keyboards"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanKeyboards reports whether the receiver responds to the private selector _keyboards.
func (v VZVirtualMachine) CanKeyboards() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_keyboards"))
}

// Keyboards is an exported wrapper for the private property _keyboards.
func (v VZVirtualMachine) Keyboards() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_keyboards")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_keyboards"}
	}
	return v._keyboards(), nil
}
func (v VZVirtualMachine) _multiTouchDevices() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_multiTouchDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanMultiTouchDevices reports whether the receiver responds to the private selector _multiTouchDevices.
func (v VZVirtualMachine) CanMultiTouchDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_multiTouchDevices"))
}

// MultiTouchDevices is an exported wrapper for the private property _multiTouchDevices.
func (v VZVirtualMachine) MultiTouchDevices() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_multiTouchDevices")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_multiTouchDevices"}
	}
	return v._multiTouchDevices(), nil
}
func (v VZVirtualMachine) _name() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_name"))
	return foundation.NSStringFromID(rv).String()
}

// CanName reports whether the receiver responds to the private selector _name.
func (v VZVirtualMachine) CanName() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_name"))
}

// Name is an exported wrapper for the private property _name.
func (v VZVirtualMachine) Name() (string, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_name")) {
		return "", &objc.UnrecognizedSelectorError{Selector: "_name"}
	}
	return v._name(), nil
}
func (v VZVirtualMachine) Set_name(value string) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("set_name:"), objc.String(value))
}
func (v VZVirtualMachine) _pointingDevices() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_pointingDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanPointingDevices reports whether the receiver responds to the private selector _pointingDevices.
func (v VZVirtualMachine) CanPointingDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_pointingDevices"))
}

// PointingDevices is an exported wrapper for the private property _pointingDevices.
func (v VZVirtualMachine) PointingDevices() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_pointingDevices")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_pointingDevices"}
	}
	return v._pointingDevices(), nil
}
func (v VZVirtualMachine) _powerSourceDevices() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_powerSourceDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanPowerSourceDevices reports whether the receiver responds to the private selector _powerSourceDevices.
func (v VZVirtualMachine) CanPowerSourceDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_powerSourceDevices"))
}

// PowerSourceDevices is an exported wrapper for the private property _powerSourceDevices.
func (v VZVirtualMachine) PowerSourceDevices() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_powerSourceDevices")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_powerSourceDevices"}
	}
	return v._powerSourceDevices(), nil
}
func (v VZVirtualMachine) _serialPorts() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_serialPorts"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanSerialPorts reports whether the receiver responds to the private selector _serialPorts.
func (v VZVirtualMachine) CanSerialPorts() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_serialPorts"))
}

// SerialPorts is an exported wrapper for the private property _serialPorts.
func (v VZVirtualMachine) SerialPorts() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_serialPorts")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_serialPorts"}
	}
	return v._serialPorts(), nil
}
func (v VZVirtualMachine) _serviceProcessIdentifier() int {
	rv := objc.SendIfResponds[int](v.ID, objc.Sel("_serviceProcessIdentifier"))
	return rv
}

// CanServiceProcessIdentifier reports whether the receiver responds to the private selector _serviceProcessIdentifier.
func (v VZVirtualMachine) CanServiceProcessIdentifier() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_serviceProcessIdentifier"))
}

// ServiceProcessIdentifier is an exported wrapper for the private property _serviceProcessIdentifier.
func (v VZVirtualMachine) ServiceProcessIdentifier() (int, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_serviceProcessIdentifier")) {
		return 0, &objc.UnrecognizedSelectorError{Selector: "_serviceProcessIdentifier"}
	}
	return v._serviceProcessIdentifier(), nil
}
func (v VZVirtualMachine) _stateDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_stateDescription"))
	return foundation.NSStringFromID(rv).String()
}

// CanStateDescription reports whether the receiver responds to the private selector _stateDescription.
func (v VZVirtualMachine) CanStateDescription() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_stateDescription"))
}

// StateDescription is an exported wrapper for the private property _stateDescription.
func (v VZVirtualMachine) StateDescription() (string, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_stateDescription")) {
		return "", &objc.UnrecognizedSelectorError{Selector: "_stateDescription"}
	}
	return v._stateDescription(), nil
}
func (v VZVirtualMachine) _storageDevices() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_storageDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanStorageDevices reports whether the receiver responds to the private selector _storageDevices.
func (v VZVirtualMachine) CanStorageDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_storageDevices"))
}

// StorageDevices is an exported wrapper for the private property _storageDevices.
func (v VZVirtualMachine) StorageDevices() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_storageDevices")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_storageDevices"}
	}
	return v._storageDevices(), nil
}
func (v VZVirtualMachine) State() int64 {
	rv := objc.SendIfResponds[int64](v.ID, objc.Sel("state"))
	return rv
}
func (v VZVirtualMachine) SetState(value int64) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setState:"), value)
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

// _createSharedMemoryCoresWithOptions is a synchronous wrapper around [VZVirtualMachine._createSharedMemoryCoresWithOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZVirtualMachine) _createSharedMemoryCoresWithOptions(ctx context.Context, options unsafe.Pointer) error {
	done := make(chan error, 1)
	v._createSharedMemoryCoresWithOptionsCompletionHandler(options, func(err error) {
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
func (v VZVirtualMachine) _saveMachineStateToURLOptions(ctx context.Context, url foundation.NSURL, options objectivec.IObject) error {
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

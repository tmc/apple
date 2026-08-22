// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
)

type unavailableSymbolError struct {
	symbol     string
	introduced string
	cause      error
}

func (e *unavailableSymbolError) Error() string {
	if e == nil {
		return ""
	}
	if e.introduced != "" {
		return fmt.Sprintf("CoreMIDI: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("CoreMIDI: symbol %s unavailable on this system", e.symbol)
}

func (e *unavailableSymbolError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	if err != nil {
		return err
	}
	if frameworkHandle == 0 {
		return fmt.Errorf("CoreMIDI: symbol %s unavailable because the framework could not be loaded", name)
	}
	return missingSymbolError(name, introduced, nil)
}

// registerFunc resolves a framework symbol and registers it as a Go function.
func registerFunc(fptr any, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			*errDst = fmt.Errorf("CoreMIDI: register symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	*errDst = nil
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	*dst = sym
	*errDst = nil
}

var _mIDIBluetoothDriverActivateAllConnections func() int32
var _mIDIBluetoothDriverActivateAllConnectionsErr error

func tryMIDIBluetoothDriverActivateAllConnections() (int32, error) {
	if _mIDIBluetoothDriverActivateAllConnections == nil {
		return 0, symbolCallError("MIDIBluetoothDriverActivateAllConnections", "13.0", _mIDIBluetoothDriverActivateAllConnectionsErr)
	}
	return _mIDIBluetoothDriverActivateAllConnections(), nil
}

// MIDIBluetoothDriverActivateAllConnections promote all active Bluetooth connections into an online MIDI device capable of input and output.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIBluetoothDriverActivateAllConnections()
func MIDIBluetoothDriverActivateAllConnections() int32 {
	result, callErr := tryMIDIBluetoothDriverActivateAllConnections()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIBluetoothDriverDisconnect func(uuid corefoundation.CFStringRef) int32
var _mIDIBluetoothDriverDisconnectErr error

func tryMIDIBluetoothDriverDisconnect(uuid corefoundation.CFStringRef) (int32, error) {
	if _mIDIBluetoothDriverDisconnect == nil {
		return 0, symbolCallError("MIDIBluetoothDriverDisconnect", "13.0", _mIDIBluetoothDriverDisconnectErr)
	}
	return _mIDIBluetoothDriverDisconnect(uuid), nil
}

// MIDIBluetoothDriverDisconnect disconnect the Bluetooth MIDI driver from a Bluetooth Low Energy MIDI peripheral.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIBluetoothDriverDisconnect(_:)
func MIDIBluetoothDriverDisconnect(uuid corefoundation.CFStringRef) int32 {
	result, callErr := tryMIDIBluetoothDriverDisconnect(uuid)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIClientCreate func(name corefoundation.CFStringRef, notifyProc MIDINotifyProc, notifyRefCon unsafe.Pointer, outClient *MIDIClientRef) int32
var _mIDIClientCreateErr error

func tryMIDIClientCreate(name corefoundation.CFStringRef, notifyProc MIDINotifyProc, notifyRefCon unsafe.Pointer, outClient *MIDIClientRef) (int32, error) {
	if _mIDIClientCreate == nil {
		return 0, symbolCallError("MIDIClientCreate", "10.0", _mIDIClientCreateErr)
	}
	return _mIDIClientCreate(name, notifyProc, notifyRefCon, outClient), nil
}

// MIDIClientCreate creates a MIDI client.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIClientCreate(_:_:_:_:)
func MIDIClientCreate(name corefoundation.CFStringRef, notifyProc MIDINotifyProc, notifyRefCon unsafe.Pointer, outClient *MIDIClientRef) int32 {
	result, callErr := tryMIDIClientCreate(name, notifyProc, notifyRefCon, outClient)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIClientCreateWithBlock func(name corefoundation.CFStringRef, outClient *MIDIClientRef, notifyBlock unsafe.Pointer) int32
var _mIDIClientCreateWithBlockErr error

func tryMIDIClientCreateWithBlock(name corefoundation.CFStringRef, outClient *MIDIClientRef, notifyBlock MIDINotifyBlock) (int32, error) {
	if _mIDIClientCreateWithBlock == nil {
		return 0, symbolCallError("MIDIClientCreateWithBlock", "10.11", _mIDIClientCreateWithBlockErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *MIDINotification) { notifyBlock(blockArg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _mIDIClientCreateWithBlock(name, outClient, _block0), nil
}

// MIDIClientCreateWithBlock creates a MIDI client with a callback block.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIClientCreateWithBlock(_:_:_:)
func MIDIClientCreateWithBlock(name corefoundation.CFStringRef, outClient *MIDIClientRef, notifyBlock MIDINotifyBlock) int32 {
	result, callErr := tryMIDIClientCreateWithBlock(name, outClient, notifyBlock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIClientDispose func(client MIDIClientRef) int32
var _mIDIClientDisposeErr error

func tryMIDIClientDispose(client MIDIClientRef) (int32, error) {
	if _mIDIClientDispose == nil {
		return 0, symbolCallError("MIDIClientDispose", "10.0", _mIDIClientDisposeErr)
	}
	return _mIDIClientDispose(client), nil
}

// MIDIClientDispose disposes of a MIDI client.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIClientDispose(_:)
func MIDIClientDispose(client MIDIClientRef) int32 {
	result, callErr := tryMIDIClientDispose(client)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIDestinationCreateWithProtocol func(client MIDIClientRef, name corefoundation.CFStringRef, protocol_ MIDIProtocolID, outDest *MIDIEndpointRef, readBlock unsafe.Pointer) int32
var _mIDIDestinationCreateWithProtocolErr error

func tryMIDIDestinationCreateWithProtocol(client MIDIClientRef, name corefoundation.CFStringRef, protocol_ MIDIProtocolID, outDest *MIDIEndpointRef, readBlock MIDIReceiveBlock) (int32, error) {
	if _mIDIDestinationCreateWithProtocol == nil {
		return 0, symbolCallError("MIDIDestinationCreateWithProtocol", "11.0", _mIDIDestinationCreateWithProtocolErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *MIDIEventList, blockArg1 unsafe.Pointer) {
		readBlock(blockArg0, blockArg1)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _mIDIDestinationCreateWithProtocol(client, name, protocol_, outDest, _block0), nil
}

// MIDIDestinationCreateWithProtocol creates a virtual destination in a client.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIDestinationCreateWithProtocol(_:_:_:_:_:)
func MIDIDestinationCreateWithProtocol(client MIDIClientRef, name corefoundation.CFStringRef, protocol_ MIDIProtocolID, outDest *MIDIEndpointRef, readBlock MIDIReceiveBlock) int32 {
	result, callErr := tryMIDIDestinationCreateWithProtocol(client, name, protocol_, outDest, readBlock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIDeviceAddEntity func(device MIDIDeviceRef, name corefoundation.CFStringRef, embedded bool, numSourceEndpoints uint, numDestinationEndpoints uint, newEntity *MIDIEntityRef) int32
var _mIDIDeviceAddEntityErr error

func tryMIDIDeviceAddEntity(device MIDIDeviceRef, name corefoundation.CFStringRef, embedded bool, numSourceEndpoints uint, numDestinationEndpoints uint, newEntity *MIDIEntityRef) (int32, error) {
	if _mIDIDeviceAddEntity == nil {
		return 0, symbolCallError("MIDIDeviceAddEntity", "10.0", _mIDIDeviceAddEntityErr)
	}
	return _mIDIDeviceAddEntity(device, name, embedded, numSourceEndpoints, numDestinationEndpoints, newEntity), nil
}

// MIDIDeviceAddEntity specifies one of the entities that make up a device.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceAddEntity(_:_:_:_:_:_:)
func MIDIDeviceAddEntity(device MIDIDeviceRef, name corefoundation.CFStringRef, embedded bool, numSourceEndpoints uint, numDestinationEndpoints uint, newEntity *MIDIEntityRef) int32 {
	result, callErr := tryMIDIDeviceAddEntity(device, name, embedded, numSourceEndpoints, numDestinationEndpoints, newEntity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIDeviceCreate func(owner MIDIDriverRef, name corefoundation.CFStringRef, manufacturer corefoundation.CFStringRef, model corefoundation.CFStringRef, outDevice *MIDIDeviceRef) int32
var _mIDIDeviceCreateErr error

func tryMIDIDeviceCreate(owner MIDIDriverRef, name corefoundation.CFStringRef, manufacturer corefoundation.CFStringRef, model corefoundation.CFStringRef, outDevice *MIDIDeviceRef) (int32, error) {
	if _mIDIDeviceCreate == nil {
		return 0, symbolCallError("MIDIDeviceCreate", "10.0", _mIDIDeviceCreateErr)
	}
	return _mIDIDeviceCreate(owner, name, manufacturer, model, outDevice), nil
}

// MIDIDeviceCreate creates a new device object that corresponds to the available hardware.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceCreate(_:_:_:_:_:)
func MIDIDeviceCreate(owner MIDIDriverRef, name corefoundation.CFStringRef, manufacturer corefoundation.CFStringRef, model corefoundation.CFStringRef, outDevice *MIDIDeviceRef) int32 {
	result, callErr := tryMIDIDeviceCreate(owner, name, manufacturer, model, outDevice)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIDeviceDispose func(device MIDIDeviceRef) int32
var _mIDIDeviceDisposeErr error

func tryMIDIDeviceDispose(device MIDIDeviceRef) (int32, error) {
	if _mIDIDeviceDispose == nil {
		return 0, symbolCallError("MIDIDeviceDispose", "10.3", _mIDIDeviceDisposeErr)
	}
	return _mIDIDeviceDispose(device), nil
}

// MIDIDeviceDispose disposes of a MIDI device.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceDispose(_:)
func MIDIDeviceDispose(device MIDIDeviceRef) int32 {
	result, callErr := tryMIDIDeviceDispose(device)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIDeviceGetEntity func(device MIDIDeviceRef, entityIndex0 uint) MIDIEntityRef
var _mIDIDeviceGetEntityErr error

func tryMIDIDeviceGetEntity(device MIDIDeviceRef, entityIndex0 uint) (MIDIEntityRef, error) {
	if _mIDIDeviceGetEntity == nil {
		return *new(MIDIEntityRef), symbolCallError("MIDIDeviceGetEntity", "10.0", _mIDIDeviceGetEntityErr)
	}
	return _mIDIDeviceGetEntity(device, entityIndex0), nil
}

// MIDIDeviceGetEntity returns the device’s entity at a specific index.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceGetEntity(_:_:)
func MIDIDeviceGetEntity(device MIDIDeviceRef, entityIndex0 uint) MIDIEntityRef {
	result, callErr := tryMIDIDeviceGetEntity(device, entityIndex0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIDeviceGetNumberOfEntities func(device MIDIDeviceRef) uint
var _mIDIDeviceGetNumberOfEntitiesErr error

func tryMIDIDeviceGetNumberOfEntities(device MIDIDeviceRef) (uint, error) {
	if _mIDIDeviceGetNumberOfEntities == nil {
		return 0, symbolCallError("MIDIDeviceGetNumberOfEntities", "10.0", _mIDIDeviceGetNumberOfEntitiesErr)
	}
	return _mIDIDeviceGetNumberOfEntities(device), nil
}

// MIDIDeviceGetNumberOfEntities returns the number of entities in a device.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceGetNumberOfEntities(_:)
func MIDIDeviceGetNumberOfEntities(device MIDIDeviceRef) uint {
	result, callErr := tryMIDIDeviceGetNumberOfEntities(device)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIDeviceListAddDevice func(devList MIDIDeviceListRef, dev MIDIDeviceRef) int32
var _mIDIDeviceListAddDeviceErr error

func tryMIDIDeviceListAddDevice(devList MIDIDeviceListRef, dev MIDIDeviceRef) (int32, error) {
	if _mIDIDeviceListAddDevice == nil {
		return 0, symbolCallError("MIDIDeviceListAddDevice", "10.0", _mIDIDeviceListAddDeviceErr)
	}
	return _mIDIDeviceListAddDevice(devList, dev), nil
}

// MIDIDeviceListAddDevice adds the specified device to the device list.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceListAddDevice(_:_:)
func MIDIDeviceListAddDevice(devList MIDIDeviceListRef, dev MIDIDeviceRef) int32 {
	result, callErr := tryMIDIDeviceListAddDevice(devList, dev)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIDeviceListDispose func(devList MIDIDeviceListRef) int32
var _mIDIDeviceListDisposeErr error

func tryMIDIDeviceListDispose(devList MIDIDeviceListRef) (int32, error) {
	if _mIDIDeviceListDispose == nil {
		return 0, symbolCallError("MIDIDeviceListDispose", "10.1", _mIDIDeviceListDisposeErr)
	}
	return _mIDIDeviceListDispose(devList), nil
}

// MIDIDeviceListDispose disposes of a device list, but not its devices.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceListDispose(_:)
func MIDIDeviceListDispose(devList MIDIDeviceListRef) int32 {
	result, callErr := tryMIDIDeviceListDispose(devList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIDeviceListGetDevice func(devList MIDIDeviceListRef, index0 uint) MIDIDeviceRef
var _mIDIDeviceListGetDeviceErr error

func tryMIDIDeviceListGetDevice(devList MIDIDeviceListRef, index0 uint) (MIDIDeviceRef, error) {
	if _mIDIDeviceListGetDevice == nil {
		return *new(MIDIDeviceRef), symbolCallError("MIDIDeviceListGetDevice", "10.0", _mIDIDeviceListGetDeviceErr)
	}
	return _mIDIDeviceListGetDevice(devList, index0), nil
}

// MIDIDeviceListGetDevice retrieves a MIDI device from a device list.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceListGetDevice(_:_:)
func MIDIDeviceListGetDevice(devList MIDIDeviceListRef, index0 uint) MIDIDeviceRef {
	result, callErr := tryMIDIDeviceListGetDevice(devList, index0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIDeviceListGetNumberOfDevices func(devList MIDIDeviceListRef) uint
var _mIDIDeviceListGetNumberOfDevicesErr error

func tryMIDIDeviceListGetNumberOfDevices(devList MIDIDeviceListRef) (uint, error) {
	if _mIDIDeviceListGetNumberOfDevices == nil {
		return 0, symbolCallError("MIDIDeviceListGetNumberOfDevices", "10.0", _mIDIDeviceListGetNumberOfDevicesErr)
	}
	return _mIDIDeviceListGetNumberOfDevices(devList), nil
}

// MIDIDeviceListGetNumberOfDevices retrieves the number of devices in a device list.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceListGetNumberOfDevices(_:)
func MIDIDeviceListGetNumberOfDevices(devList MIDIDeviceListRef) uint {
	result, callErr := tryMIDIDeviceListGetNumberOfDevices(devList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIDeviceNewEntity func(device MIDIDeviceRef, name corefoundation.CFStringRef, protocol_ MIDIProtocolID, embedded bool, numSourceEndpoints uint, numDestinationEndpoints uint, newEntity *MIDIEntityRef) int32
var _mIDIDeviceNewEntityErr error

func tryMIDIDeviceNewEntity(device MIDIDeviceRef, name corefoundation.CFStringRef, protocol_ MIDIProtocolID, embedded bool, numSourceEndpoints uint, numDestinationEndpoints uint, newEntity *MIDIEntityRef) (int32, error) {
	if _mIDIDeviceNewEntity == nil {
		return 0, symbolCallError("MIDIDeviceNewEntity", "11.0", _mIDIDeviceNewEntityErr)
	}
	return _mIDIDeviceNewEntity(device, name, protocol_, embedded, numSourceEndpoints, numDestinationEndpoints, newEntity), nil
}

// MIDIDeviceNewEntity adds a new entity to a device.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceNewEntity(_:_:_:_:_:_:_:)
func MIDIDeviceNewEntity(device MIDIDeviceRef, name corefoundation.CFStringRef, protocol_ MIDIProtocolID, embedded bool, numSourceEndpoints uint, numDestinationEndpoints uint, newEntity *MIDIEntityRef) int32 {
	result, callErr := tryMIDIDeviceNewEntity(device, name, protocol_, embedded, numSourceEndpoints, numDestinationEndpoints, newEntity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIDeviceRemoveEntity func(device MIDIDeviceRef, entity MIDIEntityRef) int32
var _mIDIDeviceRemoveEntityErr error

func tryMIDIDeviceRemoveEntity(device MIDIDeviceRef, entity MIDIEntityRef) (int32, error) {
	if _mIDIDeviceRemoveEntity == nil {
		return 0, symbolCallError("MIDIDeviceRemoveEntity", "10.1", _mIDIDeviceRemoveEntityErr)
	}
	return _mIDIDeviceRemoveEntity(device, entity), nil
}

// MIDIDeviceRemoveEntity removes an entity from a device.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceRemoveEntity(_:_:)
func MIDIDeviceRemoveEntity(device MIDIDeviceRef, entity MIDIEntityRef) int32 {
	result, callErr := tryMIDIDeviceRemoveEntity(device, entity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIDriverEnableMonitoring func(driver MIDIDriverRef, enabled bool) int32
var _mIDIDriverEnableMonitoringErr error

func tryMIDIDriverEnableMonitoring(driver MIDIDriverRef, enabled bool) (int32, error) {
	if _mIDIDriverEnableMonitoring == nil {
		return 0, symbolCallError("MIDIDriverEnableMonitoring", "10.1", _mIDIDriverEnableMonitoringErr)
	}
	return _mIDIDriverEnableMonitoring(driver, enabled), nil
}

// MIDIDriverEnableMonitoring enables monitoring of all outgoing MIDI packets.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIDriverEnableMonitoring(_:_:)
func MIDIDriverEnableMonitoring(driver MIDIDriverRef, enabled bool) int32 {
	result, callErr := tryMIDIDriverEnableMonitoring(driver, enabled)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIEndpointDispose func(endpt MIDIEndpointRef) int32
var _mIDIEndpointDisposeErr error

func tryMIDIEndpointDispose(endpt MIDIEndpointRef) (int32, error) {
	if _mIDIEndpointDispose == nil {
		return 0, symbolCallError("MIDIEndpointDispose", "10.0", _mIDIEndpointDisposeErr)
	}
	return _mIDIEndpointDispose(endpt), nil
}

// MIDIEndpointDispose disposes of a virtual source or destination.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIEndpointDispose(_:)
func MIDIEndpointDispose(endpt MIDIEndpointRef) int32 {
	result, callErr := tryMIDIEndpointDispose(endpt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIEndpointGetEntity func(inEndpoint MIDIEndpointRef, outEntity *MIDIEntityRef) int32
var _mIDIEndpointGetEntityErr error

func tryMIDIEndpointGetEntity(inEndpoint MIDIEndpointRef, outEntity *MIDIEntityRef) (int32, error) {
	if _mIDIEndpointGetEntity == nil {
		return 0, symbolCallError("MIDIEndpointGetEntity", "10.2", _mIDIEndpointGetEntityErr)
	}
	return _mIDIEndpointGetEntity(inEndpoint, outEntity), nil
}

// MIDIEndpointGetEntity returns an endpoint’s entity.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIEndpointGetEntity(_:_:)
func MIDIEndpointGetEntity(inEndpoint MIDIEndpointRef, outEntity *MIDIEntityRef) int32 {
	result, callErr := tryMIDIEndpointGetEntity(inEndpoint, outEntity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIEndpointGetRefCons func(endpt MIDIEndpointRef, ref1 unsafe.Pointer, ref2 unsafe.Pointer) int32
var _mIDIEndpointGetRefConsErr error

func tryMIDIEndpointGetRefCons(endpt MIDIEndpointRef, ref1 unsafe.Pointer, ref2 unsafe.Pointer) (int32, error) {
	if _mIDIEndpointGetRefCons == nil {
		return 0, symbolCallError("MIDIEndpointGetRefCons", "10.0", _mIDIEndpointGetRefConsErr)
	}
	return _mIDIEndpointGetRefCons(endpt, ref1, ref2), nil
}

// MIDIEndpointGetRefCons returns contextual data assigned to an endpoint.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIEndpointGetRefCons(_:_:_:)
func MIDIEndpointGetRefCons(endpt MIDIEndpointRef, ref1 unsafe.Pointer, ref2 unsafe.Pointer) int32 {
	result, callErr := tryMIDIEndpointGetRefCons(endpt, ref1, ref2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIEndpointSetRefCons func(endpt MIDIEndpointRef, ref1 unsafe.Pointer, ref2 unsafe.Pointer) int32
var _mIDIEndpointSetRefConsErr error

func tryMIDIEndpointSetRefCons(endpt MIDIEndpointRef, ref1 unsafe.Pointer, ref2 unsafe.Pointer) (int32, error) {
	if _mIDIEndpointSetRefCons == nil {
		return 0, symbolCallError("MIDIEndpointSetRefCons", "10.0", _mIDIEndpointSetRefConsErr)
	}
	return _mIDIEndpointSetRefCons(endpt, ref1, ref2), nil
}

// MIDIEndpointSetRefCons sets contextual data on an endpoint.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIEndpointSetRefCons(_:_:_:)
func MIDIEndpointSetRefCons(endpt MIDIEndpointRef, ref1 unsafe.Pointer, ref2 unsafe.Pointer) int32 {
	result, callErr := tryMIDIEndpointSetRefCons(endpt, ref1, ref2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIEntityAddOrRemoveEndpoints func(entity MIDIEntityRef, numSourceEndpoints uint, numDestinationEndpoints uint) int32
var _mIDIEntityAddOrRemoveEndpointsErr error

func tryMIDIEntityAddOrRemoveEndpoints(entity MIDIEntityRef, numSourceEndpoints uint, numDestinationEndpoints uint) (int32, error) {
	if _mIDIEntityAddOrRemoveEndpoints == nil {
		return 0, symbolCallError("MIDIEntityAddOrRemoveEndpoints", "10.2", _mIDIEntityAddOrRemoveEndpointsErr)
	}
	return _mIDIEntityAddOrRemoveEndpoints(entity, numSourceEndpoints, numDestinationEndpoints), nil
}

// MIDIEntityAddOrRemoveEndpoints adds or removes an entity’s endpoints.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityAddOrRemoveEndpoints(_:_:_:)
func MIDIEntityAddOrRemoveEndpoints(entity MIDIEntityRef, numSourceEndpoints uint, numDestinationEndpoints uint) int32 {
	result, callErr := tryMIDIEntityAddOrRemoveEndpoints(entity, numSourceEndpoints, numDestinationEndpoints)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIEntityGetDestination func(entity MIDIEntityRef, destIndex0 uint) MIDIEndpointRef
var _mIDIEntityGetDestinationErr error

func tryMIDIEntityGetDestination(entity MIDIEntityRef, destIndex0 uint) (MIDIEndpointRef, error) {
	if _mIDIEntityGetDestination == nil {
		return *new(MIDIEndpointRef), symbolCallError("MIDIEntityGetDestination", "10.0", _mIDIEntityGetDestinationErr)
	}
	return _mIDIEntityGetDestination(entity, destIndex0), nil
}

// MIDIEntityGetDestination returns one of an entity’s destinations.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityGetDestination(_:_:)
func MIDIEntityGetDestination(entity MIDIEntityRef, destIndex0 uint) MIDIEndpointRef {
	result, callErr := tryMIDIEntityGetDestination(entity, destIndex0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIEntityGetDevice func(inEntity MIDIEntityRef, outDevice *MIDIDeviceRef) int32
var _mIDIEntityGetDeviceErr error

func tryMIDIEntityGetDevice(inEntity MIDIEntityRef, outDevice *MIDIDeviceRef) (int32, error) {
	if _mIDIEntityGetDevice == nil {
		return 0, symbolCallError("MIDIEntityGetDevice", "10.2", _mIDIEntityGetDeviceErr)
	}
	return _mIDIEntityGetDevice(inEntity, outDevice), nil
}

// MIDIEntityGetDevice returns an entity’s device.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityGetDevice(_:_:)
func MIDIEntityGetDevice(inEntity MIDIEntityRef, outDevice *MIDIDeviceRef) int32 {
	result, callErr := tryMIDIEntityGetDevice(inEntity, outDevice)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIEntityGetNumberOfDestinations func(entity MIDIEntityRef) uint
var _mIDIEntityGetNumberOfDestinationsErr error

func tryMIDIEntityGetNumberOfDestinations(entity MIDIEntityRef) (uint, error) {
	if _mIDIEntityGetNumberOfDestinations == nil {
		return 0, symbolCallError("MIDIEntityGetNumberOfDestinations", "10.0", _mIDIEntityGetNumberOfDestinationsErr)
	}
	return _mIDIEntityGetNumberOfDestinations(entity), nil
}

// MIDIEntityGetNumberOfDestinations returns the number of destinations in an entity.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityGetNumberOfDestinations(_:)
func MIDIEntityGetNumberOfDestinations(entity MIDIEntityRef) uint {
	result, callErr := tryMIDIEntityGetNumberOfDestinations(entity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIEntityGetNumberOfSources func(entity MIDIEntityRef) uint
var _mIDIEntityGetNumberOfSourcesErr error

func tryMIDIEntityGetNumberOfSources(entity MIDIEntityRef) (uint, error) {
	if _mIDIEntityGetNumberOfSources == nil {
		return 0, symbolCallError("MIDIEntityGetNumberOfSources", "10.0", _mIDIEntityGetNumberOfSourcesErr)
	}
	return _mIDIEntityGetNumberOfSources(entity), nil
}

// MIDIEntityGetNumberOfSources returns the number of sources in an entity.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityGetNumberOfSources(_:)
func MIDIEntityGetNumberOfSources(entity MIDIEntityRef) uint {
	result, callErr := tryMIDIEntityGetNumberOfSources(entity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIEntityGetSource func(entity MIDIEntityRef, sourceIndex0 uint) MIDIEndpointRef
var _mIDIEntityGetSourceErr error

func tryMIDIEntityGetSource(entity MIDIEntityRef, sourceIndex0 uint) (MIDIEndpointRef, error) {
	if _mIDIEntityGetSource == nil {
		return *new(MIDIEndpointRef), symbolCallError("MIDIEntityGetSource", "10.0", _mIDIEntityGetSourceErr)
	}
	return _mIDIEntityGetSource(entity, sourceIndex0), nil
}

// MIDIEntityGetSource returns one of an entity’s sources.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityGetSource(_:_:)
func MIDIEntityGetSource(entity MIDIEntityRef, sourceIndex0 uint) MIDIEndpointRef {
	result, callErr := tryMIDIEntityGetSource(entity, sourceIndex0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIEventListAdd func(evtlist *MIDIEventList, listSize uint, curPacket *MIDIEventPacket, time MIDITimeStamp, wordCount uint, words *uint32) *MIDIEventPacket
var _mIDIEventListAddErr error

func tryMIDIEventListAdd(evtlist *MIDIEventList, listSize uint, curPacket *MIDIEventPacket, time MIDITimeStamp, wordCount uint, words *uint32) (*MIDIEventPacket, error) {
	if _mIDIEventListAdd == nil {
		return nil, symbolCallError("MIDIEventListAdd", "11.0", _mIDIEventListAddErr)
	}
	return _mIDIEventListAdd(evtlist, listSize, curPacket, time, wordCount, words), nil
}

// MIDIEventListAdd adds an event to an event list.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIEventListAdd(_:_:_:_:_:_:)
func MIDIEventListAdd(evtlist *MIDIEventList, listSize uint, curPacket *MIDIEventPacket, time MIDITimeStamp, wordCount uint, words *uint32) *MIDIEventPacket {
	result, callErr := tryMIDIEventListAdd(evtlist, listSize, curPacket, time, wordCount, words)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIEventListForEachEvent func(evtlist *MIDIEventList, visitor MIDIEventVisitor, visitorContext unsafe.Pointer)
var _mIDIEventListForEachEventErr error

func tryMIDIEventListForEachEvent(evtlist *MIDIEventList, visitor MIDIEventVisitor, visitorContext unsafe.Pointer) error {
	if _mIDIEventListForEachEvent == nil {
		return symbolCallError("MIDIEventListForEachEvent", "12.0", _mIDIEventListForEachEventErr)
	}
	_mIDIEventListForEachEvent(evtlist, visitor, visitorContext)
	return nil
}

// MIDIEventListForEachEvent.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIEventListForEachEvent(_:_:_:)
func MIDIEventListForEachEvent(evtlist *MIDIEventList, visitor MIDIEventVisitor, visitorContext unsafe.Pointer) {
	if callErr := tryMIDIEventListForEachEvent(evtlist, visitor, visitorContext); callErr != nil {
		panic(callErr)
	}
}

var _mIDIEventListInit func(evtlist *MIDIEventList, protocol_ MIDIProtocolID) *MIDIEventPacket
var _mIDIEventListInitErr error

func tryMIDIEventListInit(evtlist *MIDIEventList, protocol_ MIDIProtocolID) (*MIDIEventPacket, error) {
	if _mIDIEventListInit == nil {
		return nil, symbolCallError("MIDIEventListInit", "11.0", _mIDIEventListInitErr)
	}
	return _mIDIEventListInit(evtlist, protocol_), nil
}

// MIDIEventListInit initializes an event list.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIEventListInit(_:_:)
func MIDIEventListInit(evtlist *MIDIEventList, protocol_ MIDIProtocolID) *MIDIEventPacket {
	result, callErr := tryMIDIEventListInit(evtlist, protocol_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIEventPacketSysexBytesForGroup func(pkt *MIDIEventPacket, groupIndex uint8, outData *corefoundation.CFDataRef) int32
var _mIDIEventPacketSysexBytesForGroupErr error

func tryMIDIEventPacketSysexBytesForGroup(pkt *MIDIEventPacket, groupIndex uint8, outData *corefoundation.CFDataRef) (int32, error) {
	if _mIDIEventPacketSysexBytesForGroup == nil {
		return 0, symbolCallError("MIDIEventPacketSysexBytesForGroup", "14.0", _mIDIEventPacketSysexBytesForGroupErr)
	}
	return _mIDIEventPacketSysexBytesForGroup(pkt, groupIndex, outData), nil
}

// MIDIEventPacketSysexBytesForGroup gets MIDI 1.0 system-exclusive (SysEx) bytes on the indicated group.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIEventPacketSysexBytesForGroup(_:_:_:)
func MIDIEventPacketSysexBytesForGroup(pkt *MIDIEventPacket, groupIndex uint8, outData *corefoundation.CFDataRef) int32 {
	result, callErr := tryMIDIEventPacketSysexBytesForGroup(pkt, groupIndex, outData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIExternalDeviceCreate func(name corefoundation.CFStringRef, manufacturer corefoundation.CFStringRef, model corefoundation.CFStringRef, outDevice *MIDIDeviceRef) int32
var _mIDIExternalDeviceCreateErr error

func tryMIDIExternalDeviceCreate(name corefoundation.CFStringRef, manufacturer corefoundation.CFStringRef, model corefoundation.CFStringRef, outDevice *MIDIDeviceRef) (int32, error) {
	if _mIDIExternalDeviceCreate == nil {
		return 0, symbolCallError("MIDIExternalDeviceCreate", "10.1", _mIDIExternalDeviceCreateErr)
	}
	return _mIDIExternalDeviceCreate(name, manufacturer, model, outDevice), nil
}

// MIDIExternalDeviceCreate creates an external MIDI device.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIExternalDeviceCreate(_:_:_:_:)
func MIDIExternalDeviceCreate(name corefoundation.CFStringRef, manufacturer corefoundation.CFStringRef, model corefoundation.CFStringRef, outDevice *MIDIDeviceRef) int32 {
	result, callErr := tryMIDIExternalDeviceCreate(name, manufacturer, model, outDevice)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIFlushOutput func(dest MIDIEndpointRef) int32
var _mIDIFlushOutputErr error

func tryMIDIFlushOutput(dest MIDIEndpointRef) (int32, error) {
	if _mIDIFlushOutput == nil {
		return 0, symbolCallError("MIDIFlushOutput", "10.1", _mIDIFlushOutputErr)
	}
	return _mIDIFlushOutput(dest), nil
}

// MIDIFlushOutput cancels all pending events that were previously scheduled to send.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIFlushOutput(_:)
func MIDIFlushOutput(dest MIDIEndpointRef) int32 {
	result, callErr := tryMIDIFlushOutput(dest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIGetDestination func(destIndex0 uint) MIDIEndpointRef
var _mIDIGetDestinationErr error

func tryMIDIGetDestination(destIndex0 uint) (MIDIEndpointRef, error) {
	if _mIDIGetDestination == nil {
		return *new(MIDIEndpointRef), symbolCallError("MIDIGetDestination", "10.0", _mIDIGetDestinationErr)
	}
	return _mIDIGetDestination(destIndex0), nil
}

// MIDIGetDestination returns a destination in the system.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIGetDestination(_:)
func MIDIGetDestination(destIndex0 uint) MIDIEndpointRef {
	result, callErr := tryMIDIGetDestination(destIndex0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIGetDevice func(deviceIndex0 uint) MIDIDeviceRef
var _mIDIGetDeviceErr error

func tryMIDIGetDevice(deviceIndex0 uint) (MIDIDeviceRef, error) {
	if _mIDIGetDevice == nil {
		return *new(MIDIDeviceRef), symbolCallError("MIDIGetDevice", "10.0", _mIDIGetDeviceErr)
	}
	return _mIDIGetDevice(deviceIndex0), nil
}

// MIDIGetDevice returns a device from the system.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIGetDevice(_:)
func MIDIGetDevice(deviceIndex0 uint) MIDIDeviceRef {
	result, callErr := tryMIDIGetDevice(deviceIndex0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIGetDriverDeviceList func(driver MIDIDriverRef) MIDIDeviceListRef
var _mIDIGetDriverDeviceListErr error

func tryMIDIGetDriverDeviceList(driver MIDIDriverRef) (MIDIDeviceListRef, error) {
	if _mIDIGetDriverDeviceList == nil {
		return *new(MIDIDeviceListRef), symbolCallError("MIDIGetDriverDeviceList", "10.1", _mIDIGetDriverDeviceListErr)
	}
	return _mIDIGetDriverDeviceList(driver), nil
}

// MIDIGetDriverDeviceList returns the list of driver-created devices in the current MIDI setup.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIGetDriverDeviceList(_:)
func MIDIGetDriverDeviceList(driver MIDIDriverRef) MIDIDeviceListRef {
	result, callErr := tryMIDIGetDriverDeviceList(driver)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIGetDriverIORunLoop func() corefoundation.CFRunLoopRef
var _mIDIGetDriverIORunLoopErr error

func tryMIDIGetDriverIORunLoop() (corefoundation.CFRunLoopRef, error) {
	if _mIDIGetDriverIORunLoop == nil {
		return *new(corefoundation.CFRunLoopRef), symbolCallError("MIDIGetDriverIORunLoop", "10.0", _mIDIGetDriverIORunLoopErr)
	}
	return _mIDIGetDriverIORunLoop(), nil
}

// MIDIGetDriverIORunLoop returns the server’s driver I/O thread.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIGetDriverIORunLoop()
func MIDIGetDriverIORunLoop() corefoundation.CFRunLoopRef {
	result, callErr := tryMIDIGetDriverIORunLoop()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIGetExternalDevice func(deviceIndex0 uint) MIDIDeviceRef
var _mIDIGetExternalDeviceErr error

func tryMIDIGetExternalDevice(deviceIndex0 uint) (MIDIDeviceRef, error) {
	if _mIDIGetExternalDevice == nil {
		return *new(MIDIDeviceRef), symbolCallError("MIDIGetExternalDevice", "10.1", _mIDIGetExternalDeviceErr)
	}
	return _mIDIGetExternalDevice(deviceIndex0), nil
}

// MIDIGetExternalDevice returns one of the external devices in the system.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIGetExternalDevice(_:)
func MIDIGetExternalDevice(deviceIndex0 uint) MIDIDeviceRef {
	result, callErr := tryMIDIGetExternalDevice(deviceIndex0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIGetNumberOfDestinations func() uint
var _mIDIGetNumberOfDestinationsErr error

func tryMIDIGetNumberOfDestinations() (uint, error) {
	if _mIDIGetNumberOfDestinations == nil {
		return 0, symbolCallError("MIDIGetNumberOfDestinations", "10.0", _mIDIGetNumberOfDestinationsErr)
	}
	return _mIDIGetNumberOfDestinations(), nil
}

// MIDIGetNumberOfDestinations returns the number of destinations in the system.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIGetNumberOfDestinations()
func MIDIGetNumberOfDestinations() uint {
	result, callErr := tryMIDIGetNumberOfDestinations()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIGetNumberOfDevices func() uint
var _mIDIGetNumberOfDevicesErr error

func tryMIDIGetNumberOfDevices() (uint, error) {
	if _mIDIGetNumberOfDevices == nil {
		return 0, symbolCallError("MIDIGetNumberOfDevices", "10.0", _mIDIGetNumberOfDevicesErr)
	}
	return _mIDIGetNumberOfDevices(), nil
}

// MIDIGetNumberOfDevices returns the number of devices in the system.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIGetNumberOfDevices()
func MIDIGetNumberOfDevices() uint {
	result, callErr := tryMIDIGetNumberOfDevices()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIGetNumberOfExternalDevices func() uint
var _mIDIGetNumberOfExternalDevicesErr error

func tryMIDIGetNumberOfExternalDevices() (uint, error) {
	if _mIDIGetNumberOfExternalDevices == nil {
		return 0, symbolCallError("MIDIGetNumberOfExternalDevices", "10.1", _mIDIGetNumberOfExternalDevicesErr)
	}
	return _mIDIGetNumberOfExternalDevices(), nil
}

// MIDIGetNumberOfExternalDevices returns the number of external MIDI devices in the system.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIGetNumberOfExternalDevices()
func MIDIGetNumberOfExternalDevices() uint {
	result, callErr := tryMIDIGetNumberOfExternalDevices()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIGetNumberOfSources func() uint
var _mIDIGetNumberOfSourcesErr error

func tryMIDIGetNumberOfSources() (uint, error) {
	if _mIDIGetNumberOfSources == nil {
		return 0, symbolCallError("MIDIGetNumberOfSources", "10.0", _mIDIGetNumberOfSourcesErr)
	}
	return _mIDIGetNumberOfSources(), nil
}

// MIDIGetNumberOfSources returns the number of sources in the system.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIGetNumberOfSources()
func MIDIGetNumberOfSources() uint {
	result, callErr := tryMIDIGetNumberOfSources()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIGetSerialPortDrivers func(outDriverNames *corefoundation.CFArrayRef) int32
var _mIDIGetSerialPortDriversErr error

func tryMIDIGetSerialPortDrivers(outDriverNames *corefoundation.CFArrayRef) (int32, error) {
	if _mIDIGetSerialPortDrivers == nil {
		return 0, symbolCallError("MIDIGetSerialPortDrivers", "10.1", _mIDIGetSerialPortDriversErr)
	}
	return _mIDIGetSerialPortDrivers(outDriverNames), nil
}

// MIDIGetSerialPortDrivers returns a list of installed MIDI drivers for serial port MIDI devices.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIGetSerialPortDrivers
func MIDIGetSerialPortDrivers(outDriverNames *corefoundation.CFArrayRef) int32 {
	result, callErr := tryMIDIGetSerialPortDrivers(outDriverNames)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIGetSerialPortOwner func(portName corefoundation.CFStringRef, outDriverName *corefoundation.CFStringRef) int32
var _mIDIGetSerialPortOwnerErr error

func tryMIDIGetSerialPortOwner(portName corefoundation.CFStringRef, outDriverName *corefoundation.CFStringRef) (int32, error) {
	if _mIDIGetSerialPortOwner == nil {
		return 0, symbolCallError("MIDIGetSerialPortOwner", "10.1", _mIDIGetSerialPortOwnerErr)
	}
	return _mIDIGetSerialPortOwner(portName, outDriverName), nil
}

// MIDIGetSerialPortOwner returns the MIDI driver that owns a serial port.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIGetSerialPortOwner
func MIDIGetSerialPortOwner(portName corefoundation.CFStringRef, outDriverName *corefoundation.CFStringRef) int32 {
	result, callErr := tryMIDIGetSerialPortOwner(portName, outDriverName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIGetSource func(sourceIndex0 uint) MIDIEndpointRef
var _mIDIGetSourceErr error

func tryMIDIGetSource(sourceIndex0 uint) (MIDIEndpointRef, error) {
	if _mIDIGetSource == nil {
		return *new(MIDIEndpointRef), symbolCallError("MIDIGetSource", "10.0", _mIDIGetSourceErr)
	}
	return _mIDIGetSource(sourceIndex0), nil
}

// MIDIGetSource returns a source in the system.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIGetSource(_:)
func MIDIGetSource(sourceIndex0 uint) MIDIEndpointRef {
	result, callErr := tryMIDIGetSource(sourceIndex0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIInputPortCreateWithProtocol func(client MIDIClientRef, portName corefoundation.CFStringRef, protocol_ MIDIProtocolID, outPort *MIDIPortRef, receiveBlock unsafe.Pointer) int32
var _mIDIInputPortCreateWithProtocolErr error

func tryMIDIInputPortCreateWithProtocol(client MIDIClientRef, portName corefoundation.CFStringRef, protocol_ MIDIProtocolID, outPort *MIDIPortRef, receiveBlock MIDIReceiveBlock) (int32, error) {
	if _mIDIInputPortCreateWithProtocol == nil {
		return 0, symbolCallError("MIDIInputPortCreateWithProtocol", "11.0", _mIDIInputPortCreateWithProtocolErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *MIDIEventList, blockArg1 unsafe.Pointer) {
		receiveBlock(blockArg0, blockArg1)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _mIDIInputPortCreateWithProtocol(client, portName, protocol_, outPort, _block0), nil
}

// MIDIInputPortCreateWithProtocol creates an input port through which the client may receive incoming MIDI messages from any MIDI source.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIInputPortCreateWithProtocol(_:_:_:_:_:)
func MIDIInputPortCreateWithProtocol(client MIDIClientRef, portName corefoundation.CFStringRef, protocol_ MIDIProtocolID, outPort *MIDIPortRef, receiveBlock MIDIReceiveBlock) int32 {
	result, callErr := tryMIDIInputPortCreateWithProtocol(client, portName, protocol_, outPort, receiveBlock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIObjectFindByUniqueID func(inUniqueID MIDIUniqueID, outObject *MIDIObjectRef, outObjectType *MIDIObjectType) int32
var _mIDIObjectFindByUniqueIDErr error

func tryMIDIObjectFindByUniqueID(inUniqueID MIDIUniqueID, outObject *MIDIObjectRef, outObjectType *MIDIObjectType) (int32, error) {
	if _mIDIObjectFindByUniqueID == nil {
		return 0, symbolCallError("MIDIObjectFindByUniqueID", "10.2", _mIDIObjectFindByUniqueIDErr)
	}
	return _mIDIObjectFindByUniqueID(inUniqueID, outObject, outObjectType), nil
}

// MIDIObjectFindByUniqueID locates a device, entity, or endpoint by its unique identifier.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectFindByUniqueID(_:_:_:)
func MIDIObjectFindByUniqueID(inUniqueID MIDIUniqueID, outObject *MIDIObjectRef, outObjectType *MIDIObjectType) int32 {
	result, callErr := tryMIDIObjectFindByUniqueID(inUniqueID, outObject, outObjectType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIObjectGetDataProperty func(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, outData *corefoundation.CFDataRef) int32
var _mIDIObjectGetDataPropertyErr error

func tryMIDIObjectGetDataProperty(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, outData *corefoundation.CFDataRef) (int32, error) {
	if _mIDIObjectGetDataProperty == nil {
		return 0, symbolCallError("MIDIObjectGetDataProperty", "10.0", _mIDIObjectGetDataPropertyErr)
	}
	return _mIDIObjectGetDataProperty(obj, propertyID, outData), nil
}

// MIDIObjectGetDataProperty gets an object’s data-type property.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectGetDataProperty(_:_:_:)
func MIDIObjectGetDataProperty(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, outData *corefoundation.CFDataRef) int32 {
	result, callErr := tryMIDIObjectGetDataProperty(obj, propertyID, outData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIObjectGetDictionaryProperty func(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, outDict *corefoundation.CFDictionaryRef) int32
var _mIDIObjectGetDictionaryPropertyErr error

func tryMIDIObjectGetDictionaryProperty(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, outDict *corefoundation.CFDictionaryRef) (int32, error) {
	if _mIDIObjectGetDictionaryProperty == nil {
		return 0, symbolCallError("MIDIObjectGetDictionaryProperty", "10.2", _mIDIObjectGetDictionaryPropertyErr)
	}
	return _mIDIObjectGetDictionaryProperty(obj, propertyID, outDict), nil
}

// MIDIObjectGetDictionaryProperty gets an object’s dictionary-type property.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectGetDictionaryProperty(_:_:_:)
func MIDIObjectGetDictionaryProperty(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, outDict *corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryMIDIObjectGetDictionaryProperty(obj, propertyID, outDict)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIObjectGetIntegerProperty func(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, outValue *int32) int32
var _mIDIObjectGetIntegerPropertyErr error

func tryMIDIObjectGetIntegerProperty(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, outValue *int32) (int32, error) {
	if _mIDIObjectGetIntegerProperty == nil {
		return 0, symbolCallError("MIDIObjectGetIntegerProperty", "10.0", _mIDIObjectGetIntegerPropertyErr)
	}
	return _mIDIObjectGetIntegerProperty(obj, propertyID, outValue), nil
}

// MIDIObjectGetIntegerProperty gets an object’s integer-type property.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectGetIntegerProperty(_:_:_:)
func MIDIObjectGetIntegerProperty(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, outValue *int32) int32 {
	result, callErr := tryMIDIObjectGetIntegerProperty(obj, propertyID, outValue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIObjectGetProperties func(obj MIDIObjectRef, outProperties *corefoundation.CFPropertyListRef, deep bool) int32
var _mIDIObjectGetPropertiesErr error

func tryMIDIObjectGetProperties(obj MIDIObjectRef, outProperties *corefoundation.CFPropertyListRef, deep bool) (int32, error) {
	if _mIDIObjectGetProperties == nil {
		return 0, symbolCallError("MIDIObjectGetProperties", "10.1", _mIDIObjectGetPropertiesErr)
	}
	return _mIDIObjectGetProperties(obj, outProperties, deep), nil
}

// MIDIObjectGetProperties returns all properties of an object.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectGetProperties(_:_:_:)
func MIDIObjectGetProperties(obj MIDIObjectRef, outProperties *corefoundation.CFPropertyListRef, deep bool) int32 {
	result, callErr := tryMIDIObjectGetProperties(obj, outProperties, deep)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIObjectGetStringProperty func(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, str *corefoundation.CFStringRef) int32
var _mIDIObjectGetStringPropertyErr error

func tryMIDIObjectGetStringProperty(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, str *corefoundation.CFStringRef) (int32, error) {
	if _mIDIObjectGetStringProperty == nil {
		return 0, symbolCallError("MIDIObjectGetStringProperty", "10.0", _mIDIObjectGetStringPropertyErr)
	}
	return _mIDIObjectGetStringProperty(obj, propertyID, str), nil
}

// MIDIObjectGetStringProperty gets an object’s string-type property.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectGetStringProperty(_:_:_:)
func MIDIObjectGetStringProperty(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, str *corefoundation.CFStringRef) int32 {
	result, callErr := tryMIDIObjectGetStringProperty(obj, propertyID, str)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIObjectRemoveProperty func(obj MIDIObjectRef, propertyID corefoundation.CFStringRef) int32
var _mIDIObjectRemovePropertyErr error

func tryMIDIObjectRemoveProperty(obj MIDIObjectRef, propertyID corefoundation.CFStringRef) (int32, error) {
	if _mIDIObjectRemoveProperty == nil {
		return 0, symbolCallError("MIDIObjectRemoveProperty", "10.2", _mIDIObjectRemovePropertyErr)
	}
	return _mIDIObjectRemoveProperty(obj, propertyID), nil
}

// MIDIObjectRemoveProperty removes an object’s property.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectRemoveProperty(_:_:)
func MIDIObjectRemoveProperty(obj MIDIObjectRef, propertyID corefoundation.CFStringRef) int32 {
	result, callErr := tryMIDIObjectRemoveProperty(obj, propertyID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIObjectSetDataProperty func(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, data corefoundation.CFDataRef) int32
var _mIDIObjectSetDataPropertyErr error

func tryMIDIObjectSetDataProperty(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, data corefoundation.CFDataRef) (int32, error) {
	if _mIDIObjectSetDataProperty == nil {
		return 0, symbolCallError("MIDIObjectSetDataProperty", "10.0", _mIDIObjectSetDataPropertyErr)
	}
	return _mIDIObjectSetDataProperty(obj, propertyID, data), nil
}

// MIDIObjectSetDataProperty sets an object’s data-type property.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectSetDataProperty(_:_:_:)
func MIDIObjectSetDataProperty(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, data corefoundation.CFDataRef) int32 {
	result, callErr := tryMIDIObjectSetDataProperty(obj, propertyID, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIObjectSetDictionaryProperty func(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, dict corefoundation.CFDictionaryRef) int32
var _mIDIObjectSetDictionaryPropertyErr error

func tryMIDIObjectSetDictionaryProperty(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, dict corefoundation.CFDictionaryRef) (int32, error) {
	if _mIDIObjectSetDictionaryProperty == nil {
		return 0, symbolCallError("MIDIObjectSetDictionaryProperty", "10.2", _mIDIObjectSetDictionaryPropertyErr)
	}
	return _mIDIObjectSetDictionaryProperty(obj, propertyID, dict), nil
}

// MIDIObjectSetDictionaryProperty sets an object’s dictionary-type property.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectSetDictionaryProperty(_:_:_:)
func MIDIObjectSetDictionaryProperty(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, dict corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryMIDIObjectSetDictionaryProperty(obj, propertyID, dict)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIObjectSetIntegerProperty func(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, value int32) int32
var _mIDIObjectSetIntegerPropertyErr error

func tryMIDIObjectSetIntegerProperty(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, value int32) (int32, error) {
	if _mIDIObjectSetIntegerProperty == nil {
		return 0, symbolCallError("MIDIObjectSetIntegerProperty", "10.0", _mIDIObjectSetIntegerPropertyErr)
	}
	return _mIDIObjectSetIntegerProperty(obj, propertyID, value), nil
}

// MIDIObjectSetIntegerProperty sets an object’s integer-type property.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectSetIntegerProperty(_:_:_:)
func MIDIObjectSetIntegerProperty(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, value int32) int32 {
	result, callErr := tryMIDIObjectSetIntegerProperty(obj, propertyID, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIObjectSetStringProperty func(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, str corefoundation.CFStringRef) int32
var _mIDIObjectSetStringPropertyErr error

func tryMIDIObjectSetStringProperty(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, str corefoundation.CFStringRef) (int32, error) {
	if _mIDIObjectSetStringProperty == nil {
		return 0, symbolCallError("MIDIObjectSetStringProperty", "10.0", _mIDIObjectSetStringPropertyErr)
	}
	return _mIDIObjectSetStringProperty(obj, propertyID, str), nil
}

// MIDIObjectSetStringProperty sets an object’s string-type property.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectSetStringProperty(_:_:_:)
func MIDIObjectSetStringProperty(obj MIDIObjectRef, propertyID corefoundation.CFStringRef, str corefoundation.CFStringRef) int32 {
	result, callErr := tryMIDIObjectSetStringProperty(obj, propertyID, str)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIOutputPortCreate func(client MIDIClientRef, portName corefoundation.CFStringRef, outPort *MIDIPortRef) int32
var _mIDIOutputPortCreateErr error

func tryMIDIOutputPortCreate(client MIDIClientRef, portName corefoundation.CFStringRef, outPort *MIDIPortRef) (int32, error) {
	if _mIDIOutputPortCreate == nil {
		return 0, symbolCallError("MIDIOutputPortCreate", "10.0", _mIDIOutputPortCreateErr)
	}
	return _mIDIOutputPortCreate(client, portName, outPort), nil
}

// MIDIOutputPortCreate creates an output port through which a client sends outgoing MIDI messages to any MIDI destination.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIOutputPortCreate(_:_:_:)
func MIDIOutputPortCreate(client MIDIClientRef, portName corefoundation.CFStringRef, outPort *MIDIPortRef) int32 {
	result, callErr := tryMIDIOutputPortCreate(client, portName, outPort)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIPortConnectSource func(port MIDIPortRef, source MIDIEndpointRef, connRefCon unsafe.Pointer) int32
var _mIDIPortConnectSourceErr error

func tryMIDIPortConnectSource(port MIDIPortRef, source MIDIEndpointRef, connRefCon unsafe.Pointer) (int32, error) {
	if _mIDIPortConnectSource == nil {
		return 0, symbolCallError("MIDIPortConnectSource", "10.0", _mIDIPortConnectSourceErr)
	}
	return _mIDIPortConnectSource(port, source, connRefCon), nil
}

// MIDIPortConnectSource makes a connection from a source to a client input port.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIPortConnectSource(_:_:_:)
func MIDIPortConnectSource(port MIDIPortRef, source MIDIEndpointRef, connRefCon unsafe.Pointer) int32 {
	result, callErr := tryMIDIPortConnectSource(port, source, connRefCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIPortDisconnectSource func(port MIDIPortRef, source MIDIEndpointRef) int32
var _mIDIPortDisconnectSourceErr error

func tryMIDIPortDisconnectSource(port MIDIPortRef, source MIDIEndpointRef) (int32, error) {
	if _mIDIPortDisconnectSource == nil {
		return 0, symbolCallError("MIDIPortDisconnectSource", "10.0", _mIDIPortDisconnectSourceErr)
	}
	return _mIDIPortDisconnectSource(port, source), nil
}

// MIDIPortDisconnectSource closes a previously established source-to-input port connection.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIPortDisconnectSource(_:_:)
func MIDIPortDisconnectSource(port MIDIPortRef, source MIDIEndpointRef) int32 {
	result, callErr := tryMIDIPortDisconnectSource(port, source)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIPortDispose func(port MIDIPortRef) int32
var _mIDIPortDisposeErr error

func tryMIDIPortDispose(port MIDIPortRef) (int32, error) {
	if _mIDIPortDispose == nil {
		return 0, symbolCallError("MIDIPortDispose", "10.0", _mIDIPortDisposeErr)
	}
	return _mIDIPortDispose(port), nil
}

// MIDIPortDispose disposes of a MIDI port.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIPortDispose(_:)
func MIDIPortDispose(port MIDIPortRef) int32 {
	result, callErr := tryMIDIPortDispose(port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIReceivedEventList func(src MIDIEndpointRef, evtlist *MIDIEventList) int32
var _mIDIReceivedEventListErr error

func tryMIDIReceivedEventList(src MIDIEndpointRef, evtlist *MIDIEventList) (int32, error) {
	if _mIDIReceivedEventList == nil {
		return 0, symbolCallError("MIDIReceivedEventList", "11.0", _mIDIReceivedEventListErr)
	}
	return _mIDIReceivedEventList(src, evtlist), nil
}

// MIDIReceivedEventList distributes incoming MIDI events from a source to its connected client input ports.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIReceivedEventList(_:_:)
func MIDIReceivedEventList(src MIDIEndpointRef, evtlist *MIDIEventList) int32 {
	result, callErr := tryMIDIReceivedEventList(src, evtlist)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIRestart func() int32
var _mIDIRestartErr error

func tryMIDIRestart() (int32, error) {
	if _mIDIRestart == nil {
		return 0, symbolCallError("MIDIRestart", "10.1", _mIDIRestartErr)
	}
	return _mIDIRestart(), nil
}

// MIDIRestart stops and restarts MIDI I/O.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIRestart()
func MIDIRestart() int32 {
	result, callErr := tryMIDIRestart()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDISendEventList func(port MIDIPortRef, dest MIDIEndpointRef, evtlist *MIDIEventList) int32
var _mIDISendEventListErr error

func tryMIDISendEventList(port MIDIPortRef, dest MIDIEndpointRef, evtlist *MIDIEventList) (int32, error) {
	if _mIDISendEventList == nil {
		return 0, symbolCallError("MIDISendEventList", "11.0", _mIDISendEventListErr)
	}
	return _mIDISendEventList(port, dest, evtlist), nil
}

// MIDISendEventList sends MIDI events to a destination.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDISendEventList(_:_:_:)
func MIDISendEventList(port MIDIPortRef, dest MIDIEndpointRef, evtlist *MIDIEventList) int32 {
	result, callErr := tryMIDISendEventList(port, dest, evtlist)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDISendSysex func(request *MIDISysexSendRequest) int32
var _mIDISendSysexErr error

func tryMIDISendSysex(request *MIDISysexSendRequest) (int32, error) {
	if _mIDISendSysex == nil {
		return 0, symbolCallError("MIDISendSysex", "10.0", _mIDISendSysexErr)
	}
	return _mIDISendSysex(request), nil
}

// MIDISendSysex asynchronously sends a single system-exclusive (SysEx) event.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDISendSysex(_:)
func MIDISendSysex(request *MIDISysexSendRequest) int32 {
	result, callErr := tryMIDISendSysex(request)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDISendUMPSysex func(umpRequest *MIDISysexSendRequestUMP) int32
var _mIDISendUMPSysexErr error

func tryMIDISendUMPSysex(umpRequest *MIDISysexSendRequestUMP) (int32, error) {
	if _mIDISendUMPSysex == nil {
		return 0, symbolCallError("MIDISendUMPSysex", "14.0", _mIDISendUMPSysexErr)
	}
	return _mIDISendUMPSysex(umpRequest), nil
}

// MIDISendUMPSysex asynchronously sends a single universal MIDI packet (UMP) system-exclusive (SysEx) event.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDISendUMPSysex(_:)
func MIDISendUMPSysex(umpRequest *MIDISysexSendRequestUMP) int32 {
	result, callErr := tryMIDISendUMPSysex(umpRequest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDISendUMPSysex8 func(umpRequest *MIDISysexSendRequestUMP) int32
var _mIDISendUMPSysex8Err error

func tryMIDISendUMPSysex8(umpRequest *MIDISysexSendRequestUMP) (int32, error) {
	if _mIDISendUMPSysex8 == nil {
		return 0, symbolCallError("MIDISendUMPSysex8", "14.0", _mIDISendUMPSysex8Err)
	}
	return _mIDISendUMPSysex8(umpRequest), nil
}

// MIDISendUMPSysex8 asynchronously sends a single universal MIDI packet (UMP) system-exclusive (SysEx) event with an 8-bit message.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDISendUMPSysex8(_:)
func MIDISendUMPSysex8(umpRequest *MIDISysexSendRequestUMP) int32 {
	result, callErr := tryMIDISendUMPSysex8(umpRequest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDISetSerialPortOwner func(portName corefoundation.CFStringRef, driverName corefoundation.CFStringRef) int32
var _mIDISetSerialPortOwnerErr error

func tryMIDISetSerialPortOwner(portName corefoundation.CFStringRef, driverName corefoundation.CFStringRef) (int32, error) {
	if _mIDISetSerialPortOwner == nil {
		return 0, symbolCallError("MIDISetSerialPortOwner", "10.1", _mIDISetSerialPortOwnerErr)
	}
	return _mIDISetSerialPortOwner(portName, driverName), nil
}

// MIDISetSerialPortOwner specifies the MIDI driver that owns a serial port.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDISetSerialPortOwner
func MIDISetSerialPortOwner(portName corefoundation.CFStringRef, driverName corefoundation.CFStringRef) int32 {
	result, callErr := tryMIDISetSerialPortOwner(portName, driverName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDISetupAddDevice func(device MIDIDeviceRef) int32
var _mIDISetupAddDeviceErr error

func tryMIDISetupAddDevice(device MIDIDeviceRef) (int32, error) {
	if _mIDISetupAddDevice == nil {
		return 0, symbolCallError("MIDISetupAddDevice", "10.1", _mIDISetupAddDeviceErr)
	}
	return _mIDISetupAddDevice(device), nil
}

// MIDISetupAddDevice adds a driver-owned MIDI device to the current MIDI setup.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDISetupAddDevice(_:)
func MIDISetupAddDevice(device MIDIDeviceRef) int32 {
	result, callErr := tryMIDISetupAddDevice(device)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDISetupAddExternalDevice func(device MIDIDeviceRef) int32
var _mIDISetupAddExternalDeviceErr error

func tryMIDISetupAddExternalDevice(device MIDIDeviceRef) (int32, error) {
	if _mIDISetupAddExternalDevice == nil {
		return 0, symbolCallError("MIDISetupAddExternalDevice", "10.1", _mIDISetupAddExternalDeviceErr)
	}
	return _mIDISetupAddExternalDevice(device), nil
}

// MIDISetupAddExternalDevice adds an external MIDI device to the current MIDI setup.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDISetupAddExternalDevice(_:)
func MIDISetupAddExternalDevice(device MIDIDeviceRef) int32 {
	result, callErr := tryMIDISetupAddExternalDevice(device)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDISetupCreate func(outSetup *MIDISetupRef) int32
var _mIDISetupCreateErr error

func tryMIDISetupCreate(outSetup *MIDISetupRef) (int32, error) {
	if _mIDISetupCreate == nil {
		return 0, symbolCallError("MIDISetupCreate", "10.0", _mIDISetupCreateErr)
	}
	return _mIDISetupCreate(outSetup), nil
}

// MIDISetupCreate queries drivers to discover what hardware is available.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDISetupCreate
func MIDISetupCreate(outSetup *MIDISetupRef) int32 {
	result, callErr := tryMIDISetupCreate(outSetup)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDISetupDispose func(setup MIDISetupRef) int32
var _mIDISetupDisposeErr error

func tryMIDISetupDispose(setup MIDISetupRef) (int32, error) {
	if _mIDISetupDispose == nil {
		return 0, symbolCallError("MIDISetupDispose", "10.0", _mIDISetupDisposeErr)
	}
	return _mIDISetupDispose(setup), nil
}

// MIDISetupDispose disposes the specified setup object.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDISetupDispose
func MIDISetupDispose(setup MIDISetupRef) int32 {
	result, callErr := tryMIDISetupDispose(setup)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDISetupFromData func(data corefoundation.CFDataRef, outSetup *MIDISetupRef) int32
var _mIDISetupFromDataErr error

func tryMIDISetupFromData(data corefoundation.CFDataRef, outSetup *MIDISetupRef) (int32, error) {
	if _mIDISetupFromData == nil {
		return 0, symbolCallError("MIDISetupFromData", "10.0", _mIDISetupFromDataErr)
	}
	return _mIDISetupFromData(data, outSetup), nil
}

// MIDISetupFromData creates a MIDISetup object from an XML stream.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDISetupFromData
func MIDISetupFromData(data corefoundation.CFDataRef, outSetup *MIDISetupRef) int32 {
	result, callErr := tryMIDISetupFromData(data, outSetup)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDISetupGetCurrent func(outSetup *MIDISetupRef) int32
var _mIDISetupGetCurrentErr error

func tryMIDISetupGetCurrent(outSetup *MIDISetupRef) (int32, error) {
	if _mIDISetupGetCurrent == nil {
		return 0, symbolCallError("MIDISetupGetCurrent", "10.0", _mIDISetupGetCurrentErr)
	}
	return _mIDISetupGetCurrent(outSetup), nil
}

// MIDISetupGetCurrent returns the system’s current MIDISetup.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDISetupGetCurrent
func MIDISetupGetCurrent(outSetup *MIDISetupRef) int32 {
	result, callErr := tryMIDISetupGetCurrent(outSetup)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDISetupInstall func(setup MIDISetupRef) int32
var _mIDISetupInstallErr error

func tryMIDISetupInstall(setup MIDISetupRef) (int32, error) {
	if _mIDISetupInstall == nil {
		return 0, symbolCallError("MIDISetupInstall", "10.0", _mIDISetupInstallErr)
	}
	return _mIDISetupInstall(setup), nil
}

// MIDISetupInstall installs a MIDISetup as the system’s current state.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDISetupInstall
func MIDISetupInstall(setup MIDISetupRef) int32 {
	result, callErr := tryMIDISetupInstall(setup)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDISetupRemoveDevice func(device MIDIDeviceRef) int32
var _mIDISetupRemoveDeviceErr error

func tryMIDISetupRemoveDevice(device MIDIDeviceRef) (int32, error) {
	if _mIDISetupRemoveDevice == nil {
		return 0, symbolCallError("MIDISetupRemoveDevice", "10.1", _mIDISetupRemoveDeviceErr)
	}
	return _mIDISetupRemoveDevice(device), nil
}

// MIDISetupRemoveDevice removes a driver-owned MIDI device from the current MIDI setup.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDISetupRemoveDevice(_:)
func MIDISetupRemoveDevice(device MIDIDeviceRef) int32 {
	result, callErr := tryMIDISetupRemoveDevice(device)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDISetupRemoveExternalDevice func(device MIDIDeviceRef) int32
var _mIDISetupRemoveExternalDeviceErr error

func tryMIDISetupRemoveExternalDevice(device MIDIDeviceRef) (int32, error) {
	if _mIDISetupRemoveExternalDevice == nil {
		return 0, symbolCallError("MIDISetupRemoveExternalDevice", "10.1", _mIDISetupRemoveExternalDeviceErr)
	}
	return _mIDISetupRemoveExternalDevice(device), nil
}

// MIDISetupRemoveExternalDevice removes an external MIDI device from the current MIDI setup.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDISetupRemoveExternalDevice(_:)
func MIDISetupRemoveExternalDevice(device MIDIDeviceRef) int32 {
	result, callErr := tryMIDISetupRemoveExternalDevice(device)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDISetupToData func(setup MIDISetupRef, outData *corefoundation.CFDataRef) int32
var _mIDISetupToDataErr error

func tryMIDISetupToData(setup MIDISetupRef, outData *corefoundation.CFDataRef) (int32, error) {
	if _mIDISetupToData == nil {
		return 0, symbolCallError("MIDISetupToData", "10.0", _mIDISetupToDataErr)
	}
	return _mIDISetupToData(setup, outData), nil
}

// MIDISetupToData creates an XML representation of a MIDISetup object.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDISetupToData
func MIDISetupToData(setup MIDISetupRef, outData *corefoundation.CFDataRef) int32 {
	result, callErr := tryMIDISetupToData(setup, outData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDISourceCreateWithProtocol func(client MIDIClientRef, name corefoundation.CFStringRef, protocol_ MIDIProtocolID, outSrc *MIDIEndpointRef) int32
var _mIDISourceCreateWithProtocolErr error

func tryMIDISourceCreateWithProtocol(client MIDIClientRef, name corefoundation.CFStringRef, protocol_ MIDIProtocolID, outSrc *MIDIEndpointRef) (int32, error) {
	if _mIDISourceCreateWithProtocol == nil {
		return 0, symbolCallError("MIDISourceCreateWithProtocol", "11.0", _mIDISourceCreateWithProtocolErr)
	}
	return _mIDISourceCreateWithProtocol(client, name, protocol_, outSrc), nil
}

// MIDISourceCreateWithProtocol creates a virtual source in a client.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDISourceCreateWithProtocol(_:_:_:_:)
func MIDISourceCreateWithProtocol(client MIDIClientRef, name corefoundation.CFStringRef, protocol_ MIDIProtocolID, outSrc *MIDIEndpointRef) int32 {
	result, callErr := tryMIDISourceCreateWithProtocol(client, name, protocol_, outSrc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIThruConnectionCreate func(inPersistentOwnerID corefoundation.CFStringRef, inConnectionParams corefoundation.CFDataRef, outConnection *MIDIThruConnectionRef) int32
var _mIDIThruConnectionCreateErr error

func tryMIDIThruConnectionCreate(inPersistentOwnerID corefoundation.CFStringRef, inConnectionParams corefoundation.CFDataRef, outConnection *MIDIThruConnectionRef) (int32, error) {
	if _mIDIThruConnectionCreate == nil {
		return 0, symbolCallError("MIDIThruConnectionCreate", "10.2", _mIDIThruConnectionCreateErr)
	}
	return _mIDIThruConnectionCreate(inPersistentOwnerID, inConnectionParams, outConnection), nil
}

// MIDIThruConnectionCreate creates a MIDI thru connection.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionCreate(_:_:_:)
func MIDIThruConnectionCreate(inPersistentOwnerID corefoundation.CFStringRef, inConnectionParams corefoundation.CFDataRef, outConnection *MIDIThruConnectionRef) int32 {
	result, callErr := tryMIDIThruConnectionCreate(inPersistentOwnerID, inConnectionParams, outConnection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIThruConnectionDispose func(connection MIDIThruConnectionRef) int32
var _mIDIThruConnectionDisposeErr error

func tryMIDIThruConnectionDispose(connection MIDIThruConnectionRef) (int32, error) {
	if _mIDIThruConnectionDispose == nil {
		return 0, symbolCallError("MIDIThruConnectionDispose", "10.2", _mIDIThruConnectionDisposeErr)
	}
	return _mIDIThruConnectionDispose(connection), nil
}

// MIDIThruConnectionDispose disposes a MIDI thru connection.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionDispose(_:)
func MIDIThruConnectionDispose(connection MIDIThruConnectionRef) int32 {
	result, callErr := tryMIDIThruConnectionDispose(connection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIThruConnectionFind func(inPersistentOwnerID corefoundation.CFStringRef, outConnectionList *corefoundation.CFDataRef) int32
var _mIDIThruConnectionFindErr error

func tryMIDIThruConnectionFind(inPersistentOwnerID corefoundation.CFStringRef, outConnectionList *corefoundation.CFDataRef) (int32, error) {
	if _mIDIThruConnectionFind == nil {
		return 0, symbolCallError("MIDIThruConnectionFind", "10.2", _mIDIThruConnectionFindErr)
	}
	return _mIDIThruConnectionFind(inPersistentOwnerID, outConnectionList), nil
}

// MIDIThruConnectionFind finds the persistent thru connections for the specified client.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionFind(_:_:)
func MIDIThruConnectionFind(inPersistentOwnerID corefoundation.CFStringRef, outConnectionList *corefoundation.CFDataRef) int32 {
	result, callErr := tryMIDIThruConnectionFind(inPersistentOwnerID, outConnectionList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIThruConnectionGetParams func(connection MIDIThruConnectionRef, outConnectionParams *corefoundation.CFDataRef) int32
var _mIDIThruConnectionGetParamsErr error

func tryMIDIThruConnectionGetParams(connection MIDIThruConnectionRef, outConnectionParams *corefoundation.CFDataRef) (int32, error) {
	if _mIDIThruConnectionGetParams == nil {
		return 0, symbolCallError("MIDIThruConnectionGetParams", "10.2", _mIDIThruConnectionGetParamsErr)
	}
	return _mIDIThruConnectionGetParams(connection, outConnectionParams), nil
}

// MIDIThruConnectionGetParams returns the thru connection’s parameters.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionGetParams(_:_:)
func MIDIThruConnectionGetParams(connection MIDIThruConnectionRef, outConnectionParams *corefoundation.CFDataRef) int32 {
	result, callErr := tryMIDIThruConnectionGetParams(connection, outConnectionParams)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mIDIThruConnectionParamsInitialize func(inConnectionParams *MIDIThruConnectionParams)
var _mIDIThruConnectionParamsInitializeErr error

func tryMIDIThruConnectionParamsInitialize(inConnectionParams *MIDIThruConnectionParams) error {
	if _mIDIThruConnectionParamsInitialize == nil {
		return symbolCallError("MIDIThruConnectionParamsInitialize", "10.2", _mIDIThruConnectionParamsInitializeErr)
	}
	_mIDIThruConnectionParamsInitialize(inConnectionParams)
	return nil
}

// MIDIThruConnectionParamsInitialize initializes a parameters object with its default values.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionParamsInitialize(_:)
func MIDIThruConnectionParamsInitialize(inConnectionParams *MIDIThruConnectionParams) {
	if callErr := tryMIDIThruConnectionParamsInitialize(inConnectionParams); callErr != nil {
		panic(callErr)
	}
}

var _mIDIThruConnectionSetParams func(connection MIDIThruConnectionRef, inConnectionParams corefoundation.CFDataRef) int32
var _mIDIThruConnectionSetParamsErr error

func tryMIDIThruConnectionSetParams(connection MIDIThruConnectionRef, inConnectionParams corefoundation.CFDataRef) (int32, error) {
	if _mIDIThruConnectionSetParams == nil {
		return 0, symbolCallError("MIDIThruConnectionSetParams", "10.2", _mIDIThruConnectionSetParamsErr)
	}
	return _mIDIThruConnectionSetParams(connection, inConnectionParams), nil
}

// MIDIThruConnectionSetParams updates a thru connection’s parameters.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionSetParams(_:_:)
func MIDIThruConnectionSetParams(connection MIDIThruConnectionRef, inConnectionParams corefoundation.CFDataRef) int32 {
	result, callErr := tryMIDIThruConnectionSetParams(connection, inConnectionParams)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_mIDIBluetoothDriverActivateAllConnections, &_mIDIBluetoothDriverActivateAllConnectionsErr, frameworkHandle, "MIDIBluetoothDriverActivateAllConnections", "13.0")
	registerFunc(&_mIDIBluetoothDriverDisconnect, &_mIDIBluetoothDriverDisconnectErr, frameworkHandle, "MIDIBluetoothDriverDisconnect", "13.0")
	registerFunc(&_mIDIClientCreate, &_mIDIClientCreateErr, frameworkHandle, "MIDIClientCreate", "10.0")
	registerFunc(&_mIDIClientCreateWithBlock, &_mIDIClientCreateWithBlockErr, frameworkHandle, "MIDIClientCreateWithBlock", "10.11")
	registerFunc(&_mIDIClientDispose, &_mIDIClientDisposeErr, frameworkHandle, "MIDIClientDispose", "10.0")
	registerFunc(&_mIDIDestinationCreateWithProtocol, &_mIDIDestinationCreateWithProtocolErr, frameworkHandle, "MIDIDestinationCreateWithProtocol", "11.0")
	registerFunc(&_mIDIDeviceAddEntity, &_mIDIDeviceAddEntityErr, frameworkHandle, "MIDIDeviceAddEntity", "10.0")
	registerFunc(&_mIDIDeviceCreate, &_mIDIDeviceCreateErr, frameworkHandle, "MIDIDeviceCreate", "10.0")
	registerFunc(&_mIDIDeviceDispose, &_mIDIDeviceDisposeErr, frameworkHandle, "MIDIDeviceDispose", "10.3")
	registerFunc(&_mIDIDeviceGetEntity, &_mIDIDeviceGetEntityErr, frameworkHandle, "MIDIDeviceGetEntity", "10.0")
	registerFunc(&_mIDIDeviceGetNumberOfEntities, &_mIDIDeviceGetNumberOfEntitiesErr, frameworkHandle, "MIDIDeviceGetNumberOfEntities", "10.0")
	registerFunc(&_mIDIDeviceListAddDevice, &_mIDIDeviceListAddDeviceErr, frameworkHandle, "MIDIDeviceListAddDevice", "10.0")
	registerFunc(&_mIDIDeviceListDispose, &_mIDIDeviceListDisposeErr, frameworkHandle, "MIDIDeviceListDispose", "10.1")
	registerFunc(&_mIDIDeviceListGetDevice, &_mIDIDeviceListGetDeviceErr, frameworkHandle, "MIDIDeviceListGetDevice", "10.0")
	registerFunc(&_mIDIDeviceListGetNumberOfDevices, &_mIDIDeviceListGetNumberOfDevicesErr, frameworkHandle, "MIDIDeviceListGetNumberOfDevices", "10.0")
	registerFunc(&_mIDIDeviceNewEntity, &_mIDIDeviceNewEntityErr, frameworkHandle, "MIDIDeviceNewEntity", "11.0")
	registerFunc(&_mIDIDeviceRemoveEntity, &_mIDIDeviceRemoveEntityErr, frameworkHandle, "MIDIDeviceRemoveEntity", "10.1")
	registerFunc(&_mIDIDriverEnableMonitoring, &_mIDIDriverEnableMonitoringErr, frameworkHandle, "MIDIDriverEnableMonitoring", "10.1")
	registerFunc(&_mIDIEndpointDispose, &_mIDIEndpointDisposeErr, frameworkHandle, "MIDIEndpointDispose", "10.0")
	registerFunc(&_mIDIEndpointGetEntity, &_mIDIEndpointGetEntityErr, frameworkHandle, "MIDIEndpointGetEntity", "10.2")
	registerFunc(&_mIDIEndpointGetRefCons, &_mIDIEndpointGetRefConsErr, frameworkHandle, "MIDIEndpointGetRefCons", "10.0")
	registerFunc(&_mIDIEndpointSetRefCons, &_mIDIEndpointSetRefConsErr, frameworkHandle, "MIDIEndpointSetRefCons", "10.0")
	registerFunc(&_mIDIEntityAddOrRemoveEndpoints, &_mIDIEntityAddOrRemoveEndpointsErr, frameworkHandle, "MIDIEntityAddOrRemoveEndpoints", "10.2")
	registerFunc(&_mIDIEntityGetDestination, &_mIDIEntityGetDestinationErr, frameworkHandle, "MIDIEntityGetDestination", "10.0")
	registerFunc(&_mIDIEntityGetDevice, &_mIDIEntityGetDeviceErr, frameworkHandle, "MIDIEntityGetDevice", "10.2")
	registerFunc(&_mIDIEntityGetNumberOfDestinations, &_mIDIEntityGetNumberOfDestinationsErr, frameworkHandle, "MIDIEntityGetNumberOfDestinations", "10.0")
	registerFunc(&_mIDIEntityGetNumberOfSources, &_mIDIEntityGetNumberOfSourcesErr, frameworkHandle, "MIDIEntityGetNumberOfSources", "10.0")
	registerFunc(&_mIDIEntityGetSource, &_mIDIEntityGetSourceErr, frameworkHandle, "MIDIEntityGetSource", "10.0")
	registerFunc(&_mIDIEventListAdd, &_mIDIEventListAddErr, frameworkHandle, "MIDIEventListAdd", "11.0")
	registerFunc(&_mIDIEventListForEachEvent, &_mIDIEventListForEachEventErr, frameworkHandle, "MIDIEventListForEachEvent", "12.0")
	registerFunc(&_mIDIEventListInit, &_mIDIEventListInitErr, frameworkHandle, "MIDIEventListInit", "11.0")
	registerFunc(&_mIDIEventPacketSysexBytesForGroup, &_mIDIEventPacketSysexBytesForGroupErr, frameworkHandle, "MIDIEventPacketSysexBytesForGroup", "14.0")
	registerFunc(&_mIDIExternalDeviceCreate, &_mIDIExternalDeviceCreateErr, frameworkHandle, "MIDIExternalDeviceCreate", "10.1")
	registerFunc(&_mIDIFlushOutput, &_mIDIFlushOutputErr, frameworkHandle, "MIDIFlushOutput", "10.1")
	registerFunc(&_mIDIGetDestination, &_mIDIGetDestinationErr, frameworkHandle, "MIDIGetDestination", "10.0")
	registerFunc(&_mIDIGetDevice, &_mIDIGetDeviceErr, frameworkHandle, "MIDIGetDevice", "10.0")
	registerFunc(&_mIDIGetDriverDeviceList, &_mIDIGetDriverDeviceListErr, frameworkHandle, "MIDIGetDriverDeviceList", "10.1")
	registerFunc(&_mIDIGetDriverIORunLoop, &_mIDIGetDriverIORunLoopErr, frameworkHandle, "MIDIGetDriverIORunLoop", "10.0")
	registerFunc(&_mIDIGetExternalDevice, &_mIDIGetExternalDeviceErr, frameworkHandle, "MIDIGetExternalDevice", "10.1")
	registerFunc(&_mIDIGetNumberOfDestinations, &_mIDIGetNumberOfDestinationsErr, frameworkHandle, "MIDIGetNumberOfDestinations", "10.0")
	registerFunc(&_mIDIGetNumberOfDevices, &_mIDIGetNumberOfDevicesErr, frameworkHandle, "MIDIGetNumberOfDevices", "10.0")
	registerFunc(&_mIDIGetNumberOfExternalDevices, &_mIDIGetNumberOfExternalDevicesErr, frameworkHandle, "MIDIGetNumberOfExternalDevices", "10.1")
	registerFunc(&_mIDIGetNumberOfSources, &_mIDIGetNumberOfSourcesErr, frameworkHandle, "MIDIGetNumberOfSources", "10.0")
	registerFunc(&_mIDIGetSerialPortDrivers, &_mIDIGetSerialPortDriversErr, frameworkHandle, "MIDIGetSerialPortDrivers", "10.1")
	registerFunc(&_mIDIGetSerialPortOwner, &_mIDIGetSerialPortOwnerErr, frameworkHandle, "MIDIGetSerialPortOwner", "10.1")
	registerFunc(&_mIDIGetSource, &_mIDIGetSourceErr, frameworkHandle, "MIDIGetSource", "10.0")
	registerFunc(&_mIDIInputPortCreateWithProtocol, &_mIDIInputPortCreateWithProtocolErr, frameworkHandle, "MIDIInputPortCreateWithProtocol", "11.0")
	registerFunc(&_mIDIObjectFindByUniqueID, &_mIDIObjectFindByUniqueIDErr, frameworkHandle, "MIDIObjectFindByUniqueID", "10.2")
	registerFunc(&_mIDIObjectGetDataProperty, &_mIDIObjectGetDataPropertyErr, frameworkHandle, "MIDIObjectGetDataProperty", "10.0")
	registerFunc(&_mIDIObjectGetDictionaryProperty, &_mIDIObjectGetDictionaryPropertyErr, frameworkHandle, "MIDIObjectGetDictionaryProperty", "10.2")
	registerFunc(&_mIDIObjectGetIntegerProperty, &_mIDIObjectGetIntegerPropertyErr, frameworkHandle, "MIDIObjectGetIntegerProperty", "10.0")
	registerFunc(&_mIDIObjectGetProperties, &_mIDIObjectGetPropertiesErr, frameworkHandle, "MIDIObjectGetProperties", "10.1")
	registerFunc(&_mIDIObjectGetStringProperty, &_mIDIObjectGetStringPropertyErr, frameworkHandle, "MIDIObjectGetStringProperty", "10.0")
	registerFunc(&_mIDIObjectRemoveProperty, &_mIDIObjectRemovePropertyErr, frameworkHandle, "MIDIObjectRemoveProperty", "10.2")
	registerFunc(&_mIDIObjectSetDataProperty, &_mIDIObjectSetDataPropertyErr, frameworkHandle, "MIDIObjectSetDataProperty", "10.0")
	registerFunc(&_mIDIObjectSetDictionaryProperty, &_mIDIObjectSetDictionaryPropertyErr, frameworkHandle, "MIDIObjectSetDictionaryProperty", "10.2")
	registerFunc(&_mIDIObjectSetIntegerProperty, &_mIDIObjectSetIntegerPropertyErr, frameworkHandle, "MIDIObjectSetIntegerProperty", "10.0")
	registerFunc(&_mIDIObjectSetStringProperty, &_mIDIObjectSetStringPropertyErr, frameworkHandle, "MIDIObjectSetStringProperty", "10.0")
	registerFunc(&_mIDIOutputPortCreate, &_mIDIOutputPortCreateErr, frameworkHandle, "MIDIOutputPortCreate", "10.0")
	registerFunc(&_mIDIPortConnectSource, &_mIDIPortConnectSourceErr, frameworkHandle, "MIDIPortConnectSource", "10.0")
	registerFunc(&_mIDIPortDisconnectSource, &_mIDIPortDisconnectSourceErr, frameworkHandle, "MIDIPortDisconnectSource", "10.0")
	registerFunc(&_mIDIPortDispose, &_mIDIPortDisposeErr, frameworkHandle, "MIDIPortDispose", "10.0")
	registerFunc(&_mIDIReceivedEventList, &_mIDIReceivedEventListErr, frameworkHandle, "MIDIReceivedEventList", "11.0")
	registerFunc(&_mIDIRestart, &_mIDIRestartErr, frameworkHandle, "MIDIRestart", "10.1")
	registerFunc(&_mIDISendEventList, &_mIDISendEventListErr, frameworkHandle, "MIDISendEventList", "11.0")
	registerFunc(&_mIDISendSysex, &_mIDISendSysexErr, frameworkHandle, "MIDISendSysex", "10.0")
	registerFunc(&_mIDISendUMPSysex, &_mIDISendUMPSysexErr, frameworkHandle, "MIDISendUMPSysex", "14.0")
	registerFunc(&_mIDISendUMPSysex8, &_mIDISendUMPSysex8Err, frameworkHandle, "MIDISendUMPSysex8", "14.0")
	registerFunc(&_mIDISetSerialPortOwner, &_mIDISetSerialPortOwnerErr, frameworkHandle, "MIDISetSerialPortOwner", "10.1")
	registerFunc(&_mIDISetupAddDevice, &_mIDISetupAddDeviceErr, frameworkHandle, "MIDISetupAddDevice", "10.1")
	registerFunc(&_mIDISetupAddExternalDevice, &_mIDISetupAddExternalDeviceErr, frameworkHandle, "MIDISetupAddExternalDevice", "10.1")
	registerFunc(&_mIDISetupCreate, &_mIDISetupCreateErr, frameworkHandle, "MIDISetupCreate", "10.0")
	registerFunc(&_mIDISetupDispose, &_mIDISetupDisposeErr, frameworkHandle, "MIDISetupDispose", "10.0")
	registerFunc(&_mIDISetupFromData, &_mIDISetupFromDataErr, frameworkHandle, "MIDISetupFromData", "10.0")
	registerFunc(&_mIDISetupGetCurrent, &_mIDISetupGetCurrentErr, frameworkHandle, "MIDISetupGetCurrent", "10.0")
	registerFunc(&_mIDISetupInstall, &_mIDISetupInstallErr, frameworkHandle, "MIDISetupInstall", "10.0")
	registerFunc(&_mIDISetupRemoveDevice, &_mIDISetupRemoveDeviceErr, frameworkHandle, "MIDISetupRemoveDevice", "10.1")
	registerFunc(&_mIDISetupRemoveExternalDevice, &_mIDISetupRemoveExternalDeviceErr, frameworkHandle, "MIDISetupRemoveExternalDevice", "10.1")
	registerFunc(&_mIDISetupToData, &_mIDISetupToDataErr, frameworkHandle, "MIDISetupToData", "10.0")
	registerFunc(&_mIDISourceCreateWithProtocol, &_mIDISourceCreateWithProtocolErr, frameworkHandle, "MIDISourceCreateWithProtocol", "11.0")
	registerFunc(&_mIDIThruConnectionCreate, &_mIDIThruConnectionCreateErr, frameworkHandle, "MIDIThruConnectionCreate", "10.2")
	registerFunc(&_mIDIThruConnectionDispose, &_mIDIThruConnectionDisposeErr, frameworkHandle, "MIDIThruConnectionDispose", "10.2")
	registerFunc(&_mIDIThruConnectionFind, &_mIDIThruConnectionFindErr, frameworkHandle, "MIDIThruConnectionFind", "10.2")
	registerFunc(&_mIDIThruConnectionGetParams, &_mIDIThruConnectionGetParamsErr, frameworkHandle, "MIDIThruConnectionGetParams", "10.2")
	registerFunc(&_mIDIThruConnectionParamsInitialize, &_mIDIThruConnectionParamsInitializeErr, frameworkHandle, "MIDIThruConnectionParamsInitialize", "10.2")
	registerFunc(&_mIDIThruConnectionSetParams, &_mIDIThruConnectionSetParamsErr, frameworkHandle, "MIDIThruConnectionSetParams", "10.2")
}

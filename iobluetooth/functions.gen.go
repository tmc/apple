// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/kernel"
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
		return fmt.Sprintf("IOBluetooth: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("IOBluetooth: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("IOBluetooth: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("IOBluetooth: register symbol %s: %v", name, r)
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

var _iOBluetoothAddSCOAudioDevice func(device IOBluetoothDeviceRef, configDict corefoundation.CFDictionaryRef) int32
var _iOBluetoothAddSCOAudioDeviceErr error

func tryIOBluetoothAddSCOAudioDevice(device IOBluetoothDeviceRef, configDict corefoundation.CFDictionaryRef) (int32, error) {
	if _iOBluetoothAddSCOAudioDevice == nil {
		return 0, symbolCallError("IOBluetoothAddSCOAudioDevice", "10.0", _iOBluetoothAddSCOAudioDeviceErr)
	}
	return _iOBluetoothAddSCOAudioDevice(device, configDict), nil
}

// IOBluetoothAddSCOAudioDevice creates a persistent audio driver that will route audio data to/from the specified device.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothAddSCOAudioDevice
func IOBluetoothAddSCOAudioDevice(device IOBluetoothDeviceRef, configDict corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryIOBluetoothAddSCOAudioDevice(device, configDict)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothFindNumberOfRegistryEntriesOfClassName func(deviceType string) int
var _iOBluetoothFindNumberOfRegistryEntriesOfClassNameErr error

func tryIOBluetoothFindNumberOfRegistryEntriesOfClassName(deviceType string) (int, error) {
	if _iOBluetoothFindNumberOfRegistryEntriesOfClassName == nil {
		return 0, symbolCallError("IOBluetoothFindNumberOfRegistryEntriesOfClassName", "", _iOBluetoothFindNumberOfRegistryEntriesOfClassNameErr)
	}
	return _iOBluetoothFindNumberOfRegistryEntriesOfClassName(deviceType), nil
}

// IOBluetoothFindNumberOfRegistryEntriesOfClassName the number of registry entries with a device classname.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothFindNumberOfRegistryEntriesOfClassName(_:)
func IOBluetoothFindNumberOfRegistryEntriesOfClassName(deviceType string) int {
	result, callErr := tryIOBluetoothFindNumberOfRegistryEntriesOfClassName(deviceType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothGetUniqueFileNameAndPath func(inName *foundation.NSString, inPath *foundation.NSString) *foundation.NSString
var _iOBluetoothGetUniqueFileNameAndPathErr error

func tryIOBluetoothGetUniqueFileNameAndPath(inName *foundation.NSString, inPath *foundation.NSString) (*foundation.NSString, error) {
	if _iOBluetoothGetUniqueFileNameAndPath == nil {
		return nil, symbolCallError("IOBluetoothGetUniqueFileNameAndPath", "", _iOBluetoothGetUniqueFileNameAndPathErr)
	}
	return _iOBluetoothGetUniqueFileNameAndPath(inName, inPath), nil
}

// IOBluetoothGetUniqueFileNameAndPath.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothGetUniqueFileNameAndPath(_:_:)
func IOBluetoothGetUniqueFileNameAndPath(inName *foundation.NSString, inPath *foundation.NSString) *foundation.NSString {
	result, callErr := tryIOBluetoothGetUniqueFileNameAndPath(inName, inPath)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothIgnoreHIDDevice func(device IOBluetoothDeviceRef)
var _iOBluetoothIgnoreHIDDeviceErr error

func tryIOBluetoothIgnoreHIDDevice(device IOBluetoothDeviceRef) error {
	if _iOBluetoothIgnoreHIDDevice == nil {
		return symbolCallError("IOBluetoothIgnoreHIDDevice", "", _iOBluetoothIgnoreHIDDeviceErr)
	}
	_iOBluetoothIgnoreHIDDevice(device)
	return nil
}

// IOBluetoothIgnoreHIDDevice hints that the macOS Bluetooth software should ignore a HID device that connects up.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothIgnoreHIDDevice(_:)
func IOBluetoothIgnoreHIDDevice(device IOBluetoothDeviceRef) {
	if callErr := tryIOBluetoothIgnoreHIDDevice(device); callErr != nil {
		panic(callErr)
	}
}

var _iOBluetoothIsFileAppleDesignatedPIMData func(inFileName *foundation.NSString) bool
var _iOBluetoothIsFileAppleDesignatedPIMDataErr error

func tryIOBluetoothIsFileAppleDesignatedPIMData(inFileName *foundation.NSString) (bool, error) {
	if _iOBluetoothIsFileAppleDesignatedPIMData == nil {
		return false, symbolCallError("IOBluetoothIsFileAppleDesignatedPIMData", "", _iOBluetoothIsFileAppleDesignatedPIMDataErr)
	}
	return _iOBluetoothIsFileAppleDesignatedPIMData(inFileName), nil
}

// IOBluetoothIsFileAppleDesignatedPIMData apple designated PIM data is classified as: .vcard, .vcal, .vcf, .vnote, .vmsg, .vcs
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothIsFileAppleDesignatedPIMData(_:)
func IOBluetoothIsFileAppleDesignatedPIMData(inFileName *foundation.NSString) bool {
	result, callErr := tryIOBluetoothIsFileAppleDesignatedPIMData(inFileName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothL2CAPChannelRegisterForChannelCloseNotification func(channel IOBluetoothL2CAPChannelRef, callback IOBluetoothUserNotificationCallback, inRefCon unsafe.Pointer) IOBluetoothUserNotificationRef
var _iOBluetoothL2CAPChannelRegisterForChannelCloseNotificationErr error

func tryIOBluetoothL2CAPChannelRegisterForChannelCloseNotification(channel IOBluetoothL2CAPChannelRef, callback IOBluetoothUserNotificationCallback, inRefCon unsafe.Pointer) (IOBluetoothUserNotificationRef, error) {
	if _iOBluetoothL2CAPChannelRegisterForChannelCloseNotification == nil {
		return *new(IOBluetoothUserNotificationRef), symbolCallError("IOBluetoothL2CAPChannelRegisterForChannelCloseNotification", "", _iOBluetoothL2CAPChannelRegisterForChannelCloseNotificationErr)
	}
	return _iOBluetoothL2CAPChannelRegisterForChannelCloseNotification(channel, callback, inRefCon), nil
}

// IOBluetoothL2CAPChannelRegisterForChannelCloseNotification allows a client to register for a channel close notification.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannelRegisterForChannelCloseNotification(_:_:_:)
func IOBluetoothL2CAPChannelRegisterForChannelCloseNotification(channel IOBluetoothL2CAPChannelRef, callback IOBluetoothUserNotificationCallback, inRefCon unsafe.Pointer) IOBluetoothUserNotificationRef {
	result, callErr := tryIOBluetoothL2CAPChannelRegisterForChannelCloseNotification(channel, callback, inRefCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothNSStringFromDeviceAddress func(deviceAddress *BluetoothDeviceAddress) *foundation.NSString
var _iOBluetoothNSStringFromDeviceAddressErr error

func tryIOBluetoothNSStringFromDeviceAddress(deviceAddress *BluetoothDeviceAddress) (*foundation.NSString, error) {
	if _iOBluetoothNSStringFromDeviceAddress == nil {
		return nil, symbolCallError("IOBluetoothNSStringFromDeviceAddress", "", _iOBluetoothNSStringFromDeviceAddressErr)
	}
	return _iOBluetoothNSStringFromDeviceAddress(deviceAddress), nil
}

// IOBluetoothNSStringFromDeviceAddress convenience routine to take a device address structure and create an NSString.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNSStringFromDeviceAddress(_:)
func IOBluetoothNSStringFromDeviceAddress(deviceAddress *BluetoothDeviceAddress) *foundation.NSString {
	result, callErr := tryIOBluetoothNSStringFromDeviceAddress(deviceAddress)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothNSStringFromDeviceAddressColon func(deviceAddress *BluetoothDeviceAddress) *foundation.NSString
var _iOBluetoothNSStringFromDeviceAddressColonErr error

func tryIOBluetoothNSStringFromDeviceAddressColon(deviceAddress *BluetoothDeviceAddress) (*foundation.NSString, error) {
	if _iOBluetoothNSStringFromDeviceAddressColon == nil {
		return nil, symbolCallError("IOBluetoothNSStringFromDeviceAddressColon", "", _iOBluetoothNSStringFromDeviceAddressColonErr)
	}
	return _iOBluetoothNSStringFromDeviceAddressColon(deviceAddress), nil
}

// IOBluetoothNSStringFromDeviceAddressColon.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNSStringFromDeviceAddressColon(_:)
func IOBluetoothNSStringFromDeviceAddressColon(deviceAddress *BluetoothDeviceAddress) *foundation.NSString {
	result, callErr := tryIOBluetoothNSStringFromDeviceAddressColon(deviceAddress)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothNSStringToDeviceAddress func(inNameString *foundation.NSString, outDeviceAddress *BluetoothDeviceAddress) int32
var _iOBluetoothNSStringToDeviceAddressErr error

func tryIOBluetoothNSStringToDeviceAddress(inNameString *foundation.NSString, outDeviceAddress *BluetoothDeviceAddress) (int32, error) {
	if _iOBluetoothNSStringToDeviceAddress == nil {
		return 0, symbolCallError("IOBluetoothNSStringToDeviceAddress", "", _iOBluetoothNSStringToDeviceAddressErr)
	}
	return _iOBluetoothNSStringToDeviceAddress(inNameString, outDeviceAddress), nil
}

// IOBluetoothNSStringToDeviceAddress convenience routine to take an NSString and turn it into a BluetoothDeviceAddress structure.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNSStringToDeviceAddress(_:_:)
func IOBluetoothNSStringToDeviceAddress(inNameString *foundation.NSString, outDeviceAddress *BluetoothDeviceAddress) int32 {
	result, callErr := tryIOBluetoothNSStringToDeviceAddress(inNameString, outDeviceAddress)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothNumberOfAvailableHIDDevices func() int
var _iOBluetoothNumberOfAvailableHIDDevicesErr error

func tryIOBluetoothNumberOfAvailableHIDDevices() (int, error) {
	if _iOBluetoothNumberOfAvailableHIDDevices == nil {
		return 0, symbolCallError("IOBluetoothNumberOfAvailableHIDDevices", "", _iOBluetoothNumberOfAvailableHIDDevicesErr)
	}
	return _iOBluetoothNumberOfAvailableHIDDevices(), nil
}

// IOBluetoothNumberOfAvailableHIDDevices returns total number of HID devices on the system (Bluetooth + USB)
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNumberOfAvailableHIDDevices()
func IOBluetoothNumberOfAvailableHIDDevices() int {
	result, callErr := tryIOBluetoothNumberOfAvailableHIDDevices()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothNumberOfKeyboardHIDDevices func() int
var _iOBluetoothNumberOfKeyboardHIDDevicesErr error

func tryIOBluetoothNumberOfKeyboardHIDDevices() (int, error) {
	if _iOBluetoothNumberOfKeyboardHIDDevices == nil {
		return 0, symbolCallError("IOBluetoothNumberOfKeyboardHIDDevices", "", _iOBluetoothNumberOfKeyboardHIDDevicesErr)
	}
	return _iOBluetoothNumberOfKeyboardHIDDevices(), nil
}

// IOBluetoothNumberOfKeyboardHIDDevices returns number of keyboard HID devices on the system (Bluetooth + USB)
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNumberOfKeyboardHIDDevices()
func IOBluetoothNumberOfKeyboardHIDDevices() int {
	result, callErr := tryIOBluetoothNumberOfKeyboardHIDDevices()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothNumberOfPointingHIDDevices func() int
var _iOBluetoothNumberOfPointingHIDDevicesErr error

func tryIOBluetoothNumberOfPointingHIDDevices() (int, error) {
	if _iOBluetoothNumberOfPointingHIDDevices == nil {
		return 0, symbolCallError("IOBluetoothNumberOfPointingHIDDevices", "", _iOBluetoothNumberOfPointingHIDDevicesErr)
	}
	return _iOBluetoothNumberOfPointingHIDDevices(), nil
}

// IOBluetoothNumberOfPointingHIDDevices returns number of “pointing” HID devices on the system (Bluetooth + USB)
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNumberOfPointingHIDDevices()
func IOBluetoothNumberOfPointingHIDDevices() int {
	result, callErr := tryIOBluetoothNumberOfPointingHIDDevices()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothNumberOfTabletHIDDevices func() int
var _iOBluetoothNumberOfTabletHIDDevicesErr error

func tryIOBluetoothNumberOfTabletHIDDevices() (int, error) {
	if _iOBluetoothNumberOfTabletHIDDevices == nil {
		return 0, symbolCallError("IOBluetoothNumberOfTabletHIDDevices", "", _iOBluetoothNumberOfTabletHIDDevicesErr)
	}
	return _iOBluetoothNumberOfTabletHIDDevices(), nil
}

// IOBluetoothNumberOfTabletHIDDevices returns number of “Tablet” HID devices on the system (Bluetooth + USB)
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothNumberOfTabletHIDDevices()
func IOBluetoothNumberOfTabletHIDDevices() int {
	result, callErr := tryIOBluetoothNumberOfTabletHIDDevices()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber func(inDeviceRef IOBluetoothDeviceRef, inChannelID kernel.BluetoothRFCOMMChannelID, outSessionRef *OBEXSessionRef) OBEXError
var _iOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumberErr error

func tryIOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber(inDeviceRef IOBluetoothDeviceRef, inChannelID kernel.BluetoothRFCOMMChannelID, outSessionRef *OBEXSessionRef) (OBEXError, error) {
	if _iOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber == nil {
		return *new(OBEXError), symbolCallError("IOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber", "10.0", _iOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumberErr)
	}
	return _iOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber(inDeviceRef, inChannelID, outSessionRef), nil
}

// IOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber create an OBEX session with a device ref and an RFCOMM channel ID.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber
func IOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber(inDeviceRef IOBluetoothDeviceRef, inChannelID kernel.BluetoothRFCOMMChannelID, outSessionRef *OBEXSessionRef) OBEXError {
	result, callErr := tryIOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber(inDeviceRef, inChannelID, outSessionRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef func(inSDPServiceRef IOBluetoothSDPServiceRecordRef, outSessionRef *OBEXSessionRef) OBEXError
var _iOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRefErr error

func tryIOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef(inSDPServiceRef IOBluetoothSDPServiceRecordRef, outSessionRef *OBEXSessionRef) (OBEXError, error) {
	if _iOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef == nil {
		return *new(OBEXError), symbolCallError("IOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef", "10.0", _iOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRefErr)
	}
	return _iOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef(inSDPServiceRef, outSessionRef), nil
}

// IOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef create an OBEX session with a service ref, usually obtained from the device browser.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef
func IOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef(inSDPServiceRef IOBluetoothSDPServiceRecordRef, outSessionRef *OBEXSessionRef) OBEXError {
	result, callErr := tryIOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef(inSDPServiceRef, outSessionRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel func(inRFCOMMChannelRef IOBluetoothRFCOMMChannelRef, inCallback OBEXSessionEventCallback, inUserRefCon uintptr, outSessionRef *OBEXSessionRef) OBEXError
var _iOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannelErr error

func tryIOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel(inRFCOMMChannelRef IOBluetoothRFCOMMChannelRef, inCallback OBEXSessionEventCallback, inUserRefCon uintptr, outSessionRef *OBEXSessionRef) (OBEXError, error) {
	if _iOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel == nil {
		return *new(OBEXError), symbolCallError("IOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel", "10.0", _iOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannelErr)
	}
	return _iOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel(inRFCOMMChannelRef, inCallback, inUserRefCon, outSessionRef), nil
}

// IOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel create an OBEX session with an IOBluetoothRFCOMMchannel.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel
func IOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel(inRFCOMMChannelRef IOBluetoothRFCOMMChannelRef, inCallback OBEXSessionEventCallback, inUserRefCon uintptr, outSessionRef *OBEXSessionRef) OBEXError {
	result, callErr := tryIOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel(inRFCOMMChannelRef, inCallback, inUserRefCon, outSessionRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothOBEXSessionOpenTransportConnection func(inSessionRef OBEXSessionRef, inCallback IOBluetoothOBEXSessionOpenConnectionCallback, inUserRefCon uintptr) OBEXError
var _iOBluetoothOBEXSessionOpenTransportConnectionErr error

func tryIOBluetoothOBEXSessionOpenTransportConnection(inSessionRef OBEXSessionRef, inCallback IOBluetoothOBEXSessionOpenConnectionCallback, inUserRefCon uintptr) (OBEXError, error) {
	if _iOBluetoothOBEXSessionOpenTransportConnection == nil {
		return *new(OBEXError), symbolCallError("IOBluetoothOBEXSessionOpenTransportConnection", "10.0", _iOBluetoothOBEXSessionOpenTransportConnectionErr)
	}
	return _iOBluetoothOBEXSessionOpenTransportConnection(inSessionRef, inCallback, inUserRefCon), nil
}

// IOBluetoothOBEXSessionOpenTransportConnection.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSessionOpenTransportConnection
func IOBluetoothOBEXSessionOpenTransportConnection(inSessionRef OBEXSessionRef, inCallback IOBluetoothOBEXSessionOpenConnectionCallback, inUserRefCon uintptr) OBEXError {
	result, callErr := tryIOBluetoothOBEXSessionOpenTransportConnection(inSessionRef, inCallback, inUserRefCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothPackData func(ioBuffer unsafe.Pointer, inFormat string) int
var _iOBluetoothPackDataErr error

func tryIOBluetoothPackData(ioBuffer unsafe.Pointer, inFormat string) (int, error) {
	if _iOBluetoothPackData == nil {
		return 0, symbolCallError("IOBluetoothPackData", "", _iOBluetoothPackDataErr)
	}
	return _iOBluetoothPackData(ioBuffer, inFormat), nil
}

// IOBluetoothPackData packs a variable amount of parameters into a buffer according to a printf-style format string.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothPackData
func IOBluetoothPackData(ioBuffer unsafe.Pointer, inFormat string) int {
	result, callErr := tryIOBluetoothPackData(ioBuffer, inFormat)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothPackDataList func(ioBuffer unsafe.Pointer, inFormat string, inArgs kernel.Va_list) int
var _iOBluetoothPackDataListErr error

func tryIOBluetoothPackDataList(ioBuffer unsafe.Pointer, inFormat string, inArgs kernel.Va_list) (int, error) {
	if _iOBluetoothPackDataList == nil {
		return 0, symbolCallError("IOBluetoothPackDataList", "", _iOBluetoothPackDataListErr)
	}
	return _iOBluetoothPackDataList(ioBuffer, inFormat, inArgs), nil
}

// IOBluetoothPackDataList.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothPackDataList(_:_:_:)
func IOBluetoothPackDataList(ioBuffer unsafe.Pointer, inFormat string, inArgs kernel.Va_list) int {
	result, callErr := tryIOBluetoothPackDataList(ioBuffer, inFormat, inArgs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothRemoveIgnoredHIDDevice func(device IOBluetoothDeviceRef)
var _iOBluetoothRemoveIgnoredHIDDeviceErr error

func tryIOBluetoothRemoveIgnoredHIDDevice(device IOBluetoothDeviceRef) error {
	if _iOBluetoothRemoveIgnoredHIDDevice == nil {
		return symbolCallError("IOBluetoothRemoveIgnoredHIDDevice", "", _iOBluetoothRemoveIgnoredHIDDeviceErr)
	}
	_iOBluetoothRemoveIgnoredHIDDevice(device)
	return nil
}

// IOBluetoothRemoveIgnoredHIDDevice the counterpart to the above IOBluetoothIgnoreHIDDevice() API.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRemoveIgnoredHIDDevice(_:)
func IOBluetoothRemoveIgnoredHIDDevice(device IOBluetoothDeviceRef) {
	if callErr := tryIOBluetoothRemoveIgnoredHIDDevice(device); callErr != nil {
		panic(callErr)
	}
}

var _iOBluetoothRemoveSCOAudioDevice func(device IOBluetoothDeviceRef) int32
var _iOBluetoothRemoveSCOAudioDeviceErr error

func tryIOBluetoothRemoveSCOAudioDevice(device IOBluetoothDeviceRef) (int32, error) {
	if _iOBluetoothRemoveSCOAudioDevice == nil {
		return 0, symbolCallError("IOBluetoothRemoveSCOAudioDevice", "10.0", _iOBluetoothRemoveSCOAudioDeviceErr)
	}
	return _iOBluetoothRemoveSCOAudioDevice(device), nil
}

// IOBluetoothRemoveSCOAudioDevice removes a persistent audio driver for a device that had already been added using IOBluetoothAddAudioDevice().
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRemoveSCOAudioDevice
func IOBluetoothRemoveSCOAudioDevice(device IOBluetoothDeviceRef) int32 {
	result, callErr := tryIOBluetoothRemoveSCOAudioDevice(device)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothUnpackData func(inBufferSize uint, inBuffer unsafe.Pointer, inFormat *byte) int
var _iOBluetoothUnpackDataErr error

func tryIOBluetoothUnpackData(inBufferSize uint, inBuffer unsafe.Pointer, inFormat *byte) (int, error) {
	if _iOBluetoothUnpackData == nil {
		return 0, symbolCallError("IOBluetoothUnpackData", "", _iOBluetoothUnpackDataErr)
	}
	return _iOBluetoothUnpackData(inBufferSize, inBuffer, inFormat), nil
}

// IOBluetoothUnpackData unpacks a variable amount of data from a buffer into a variable number of parameters according to a printf-style format string.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothUnpackData
func IOBluetoothUnpackData(inBufferSize uint, inBuffer unsafe.Pointer, inFormat *byte) int {
	result, callErr := tryIOBluetoothUnpackData(inBufferSize, inBuffer, inFormat)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothUnpackDataList func(inBufferSize uint, inBuffer unsafe.Pointer, inFormat *byte, inArgs kernel.Va_list) int
var _iOBluetoothUnpackDataListErr error

func tryIOBluetoothUnpackDataList(inBufferSize uint, inBuffer unsafe.Pointer, inFormat *byte, inArgs kernel.Va_list) (int, error) {
	if _iOBluetoothUnpackDataList == nil {
		return 0, symbolCallError("IOBluetoothUnpackDataList", "", _iOBluetoothUnpackDataListErr)
	}
	return _iOBluetoothUnpackDataList(inBufferSize, inBuffer, inFormat, inArgs), nil
}

// IOBluetoothUnpackDataList.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothUnpackDataList(_:_:_:_:)
func IOBluetoothUnpackDataList(inBufferSize uint, inBuffer unsafe.Pointer, inFormat *byte, inArgs kernel.Va_list) int {
	result, callErr := tryIOBluetoothUnpackDataList(inBufferSize, inBuffer, inFormat, inArgs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBluetoothUserNotificationUnregister func(notificationRef IOBluetoothUserNotificationRef)
var _iOBluetoothUserNotificationUnregisterErr error

func tryIOBluetoothUserNotificationUnregister(notificationRef IOBluetoothUserNotificationRef) error {
	if _iOBluetoothUserNotificationUnregister == nil {
		return symbolCallError("IOBluetoothUserNotificationUnregister", "", _iOBluetoothUserNotificationUnregisterErr)
	}
	_iOBluetoothUserNotificationUnregister(notificationRef)
	return nil
}

// IOBluetoothUserNotificationUnregister unregisters the target notification.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothUserNotificationUnregister(_:)
func IOBluetoothUserNotificationUnregister(notificationRef IOBluetoothUserNotificationRef) {
	if callErr := tryIOBluetoothUserNotificationUnregister(notificationRef); callErr != nil {
		panic(callErr)
	}
}

var _oBEXAddApplicationParameterHeader func(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError
var _oBEXAddApplicationParameterHeaderErr error

func tryOBEXAddApplicationParameterHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) (OBEXError, error) {
	if _oBEXAddApplicationParameterHeader == nil {
		return *new(OBEXError), symbolCallError("OBEXAddApplicationParameterHeader", "", _oBEXAddApplicationParameterHeaderErr)
	}
	return _oBEXAddApplicationParameterHeader(inHeaderData, inHeaderDataLength, dictRef), nil
}

// OBEXAddApplicationParameterHeader add bytes representing an application parameter to a dictionary of OBEX headers.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXAddApplicationParameterHeader(_:_:_:)
func OBEXAddApplicationParameterHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError {
	result, callErr := tryOBEXAddApplicationParameterHeader(inHeaderData, inHeaderDataLength, dictRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXAddAuthorizationChallengeHeader func(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError
var _oBEXAddAuthorizationChallengeHeaderErr error

func tryOBEXAddAuthorizationChallengeHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) (OBEXError, error) {
	if _oBEXAddAuthorizationChallengeHeader == nil {
		return *new(OBEXError), symbolCallError("OBEXAddAuthorizationChallengeHeader", "", _oBEXAddAuthorizationChallengeHeaderErr)
	}
	return _oBEXAddAuthorizationChallengeHeader(inHeaderData, inHeaderDataLength, dictRef), nil
}

// OBEXAddAuthorizationChallengeHeader add an authorization challenge header to a dictionary of OBEXheaders.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXAddAuthorizationChallengeHeader(_:_:_:)
func OBEXAddAuthorizationChallengeHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError {
	result, callErr := tryOBEXAddAuthorizationChallengeHeader(inHeaderData, inHeaderDataLength, dictRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXAddAuthorizationResponseHeader func(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError
var _oBEXAddAuthorizationResponseHeaderErr error

func tryOBEXAddAuthorizationResponseHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) (OBEXError, error) {
	if _oBEXAddAuthorizationResponseHeader == nil {
		return *new(OBEXError), symbolCallError("OBEXAddAuthorizationResponseHeader", "", _oBEXAddAuthorizationResponseHeaderErr)
	}
	return _oBEXAddAuthorizationResponseHeader(inHeaderData, inHeaderDataLength, dictRef), nil
}

// OBEXAddAuthorizationResponseHeader add an authorization Response header to a dictionary of OBEXheaders.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXAddAuthorizationResponseHeader(_:_:_:)
func OBEXAddAuthorizationResponseHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError {
	result, callErr := tryOBEXAddAuthorizationResponseHeader(inHeaderData, inHeaderDataLength, dictRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXAddBodyHeader func(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, isEndOfBody bool, dictRef corefoundation.CFMutableDictionaryRef) OBEXError
var _oBEXAddBodyHeaderErr error

func tryOBEXAddBodyHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, isEndOfBody bool, dictRef corefoundation.CFMutableDictionaryRef) (OBEXError, error) {
	if _oBEXAddBodyHeader == nil {
		return *new(OBEXError), symbolCallError("OBEXAddBodyHeader", "", _oBEXAddBodyHeaderErr)
	}
	return _oBEXAddBodyHeader(inHeaderData, inHeaderDataLength, isEndOfBody, dictRef), nil
}

// OBEXAddBodyHeader add bytes of data to a dictionary of OBEXheaders.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXAddBodyHeader(_:_:_:_:)
func OBEXAddBodyHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, isEndOfBody bool, dictRef corefoundation.CFMutableDictionaryRef) OBEXError {
	result, callErr := tryOBEXAddBodyHeader(inHeaderData, inHeaderDataLength, isEndOfBody, dictRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXAddByteSequenceHeader func(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError
var _oBEXAddByteSequenceHeaderErr error

func tryOBEXAddByteSequenceHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) (OBEXError, error) {
	if _oBEXAddByteSequenceHeader == nil {
		return *new(OBEXError), symbolCallError("OBEXAddByteSequenceHeader", "", _oBEXAddByteSequenceHeaderErr)
	}
	return _oBEXAddByteSequenceHeader(inHeaderData, inHeaderDataLength, dictRef), nil
}

// OBEXAddByteSequenceHeader add a byte sequence header to a dictionary of OBEXheaders.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXAddByteSequenceHeader(_:_:_:)
func OBEXAddByteSequenceHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError {
	result, callErr := tryOBEXAddByteSequenceHeader(inHeaderData, inHeaderDataLength, dictRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXAddConnectionIDHeader func(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError
var _oBEXAddConnectionIDHeaderErr error

func tryOBEXAddConnectionIDHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) (OBEXError, error) {
	if _oBEXAddConnectionIDHeader == nil {
		return *new(OBEXError), symbolCallError("OBEXAddConnectionIDHeader", "", _oBEXAddConnectionIDHeaderErr)
	}
	return _oBEXAddConnectionIDHeader(inHeaderData, inHeaderDataLength, dictRef), nil
}

// OBEXAddConnectionIDHeader add bytes representing a connection ID to a dictionary of OBEX headers.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXAddConnectionIDHeader(_:_:_:)
func OBEXAddConnectionIDHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError {
	result, callErr := tryOBEXAddConnectionIDHeader(inHeaderData, inHeaderDataLength, dictRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXAddCountHeader func(count uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError
var _oBEXAddCountHeaderErr error

func tryOBEXAddCountHeader(count uint32, dictRef corefoundation.CFMutableDictionaryRef) (OBEXError, error) {
	if _oBEXAddCountHeader == nil {
		return *new(OBEXError), symbolCallError("OBEXAddCountHeader", "", _oBEXAddCountHeaderErr)
	}
	return _oBEXAddCountHeader(count, dictRef), nil
}

// OBEXAddCountHeader add a CFStringRef to a dictionary of OBEXheaders.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXAddCountHeader(_:_:)
func OBEXAddCountHeader(count uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError {
	result, callErr := tryOBEXAddCountHeader(count, dictRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXAddDescriptionHeader func(description corefoundation.CFStringRef, dictRef corefoundation.CFMutableDictionaryRef) OBEXError
var _oBEXAddDescriptionHeaderErr error

func tryOBEXAddDescriptionHeader(description corefoundation.CFStringRef, dictRef corefoundation.CFMutableDictionaryRef) (OBEXError, error) {
	if _oBEXAddDescriptionHeader == nil {
		return *new(OBEXError), symbolCallError("OBEXAddDescriptionHeader", "", _oBEXAddDescriptionHeaderErr)
	}
	return _oBEXAddDescriptionHeader(description, dictRef), nil
}

// OBEXAddDescriptionHeader add a CFStringRef to a dictionary of OBEXheaders.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXAddDescriptionHeader(_:_:)
func OBEXAddDescriptionHeader(description corefoundation.CFStringRef, dictRef corefoundation.CFMutableDictionaryRef) OBEXError {
	result, callErr := tryOBEXAddDescriptionHeader(description, dictRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXAddHTTPHeader func(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError
var _oBEXAddHTTPHeaderErr error

func tryOBEXAddHTTPHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) (OBEXError, error) {
	if _oBEXAddHTTPHeader == nil {
		return *new(OBEXError), symbolCallError("OBEXAddHTTPHeader", "", _oBEXAddHTTPHeaderErr)
	}
	return _oBEXAddHTTPHeader(inHeaderData, inHeaderDataLength, dictRef), nil
}

// OBEXAddHTTPHeader add bytes of data to a dictionary of OBEXheaders.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXAddHTTPHeader(_:_:_:)
func OBEXAddHTTPHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError {
	result, callErr := tryOBEXAddHTTPHeader(inHeaderData, inHeaderDataLength, dictRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXAddLengthHeader func(length uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError
var _oBEXAddLengthHeaderErr error

func tryOBEXAddLengthHeader(length uint32, dictRef corefoundation.CFMutableDictionaryRef) (OBEXError, error) {
	if _oBEXAddLengthHeader == nil {
		return *new(OBEXError), symbolCallError("OBEXAddLengthHeader", "", _oBEXAddLengthHeaderErr)
	}
	return _oBEXAddLengthHeader(length, dictRef), nil
}

// OBEXAddLengthHeader add a CFStringRef to a dictionary of OBEXheaders.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXAddLengthHeader(_:_:)
func OBEXAddLengthHeader(length uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError {
	result, callErr := tryOBEXAddLengthHeader(length, dictRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXAddNameHeader func(name corefoundation.CFStringRef, dictRef corefoundation.CFMutableDictionaryRef) OBEXError
var _oBEXAddNameHeaderErr error

func tryOBEXAddNameHeader(name corefoundation.CFStringRef, dictRef corefoundation.CFMutableDictionaryRef) (OBEXError, error) {
	if _oBEXAddNameHeader == nil {
		return *new(OBEXError), symbolCallError("OBEXAddNameHeader", "", _oBEXAddNameHeaderErr)
	}
	return _oBEXAddNameHeader(name, dictRef), nil
}

// OBEXAddNameHeader add a CFStringRef to a dictionary of OBEXheaders.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXAddNameHeader(_:_:)
func OBEXAddNameHeader(name corefoundation.CFStringRef, dictRef corefoundation.CFMutableDictionaryRef) OBEXError {
	result, callErr := tryOBEXAddNameHeader(name, dictRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXAddObjectClassHeader func(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError
var _oBEXAddObjectClassHeaderErr error

func tryOBEXAddObjectClassHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) (OBEXError, error) {
	if _oBEXAddObjectClassHeader == nil {
		return *new(OBEXError), symbolCallError("OBEXAddObjectClassHeader", "", _oBEXAddObjectClassHeaderErr)
	}
	return _oBEXAddObjectClassHeader(inHeaderData, inHeaderDataLength, dictRef), nil
}

// OBEXAddObjectClassHeader add an object class header to a dictionary of OBEXheaders.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXAddObjectClassHeader(_:_:_:)
func OBEXAddObjectClassHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError {
	result, callErr := tryOBEXAddObjectClassHeader(inHeaderData, inHeaderDataLength, dictRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXAddTargetHeader func(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError
var _oBEXAddTargetHeaderErr error

func tryOBEXAddTargetHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) (OBEXError, error) {
	if _oBEXAddTargetHeader == nil {
		return *new(OBEXError), symbolCallError("OBEXAddTargetHeader", "", _oBEXAddTargetHeaderErr)
	}
	return _oBEXAddTargetHeader(inHeaderData, inHeaderDataLength, dictRef), nil
}

// OBEXAddTargetHeader add bytes of data to a dictionary of OBEXheaders.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXAddTargetHeader(_:_:_:)
func OBEXAddTargetHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError {
	result, callErr := tryOBEXAddTargetHeader(inHeaderData, inHeaderDataLength, dictRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXAddTime4ByteHeader func(time4Byte uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError
var _oBEXAddTime4ByteHeaderErr error

func tryOBEXAddTime4ByteHeader(time4Byte uint32, dictRef corefoundation.CFMutableDictionaryRef) (OBEXError, error) {
	if _oBEXAddTime4ByteHeader == nil {
		return *new(OBEXError), symbolCallError("OBEXAddTime4ByteHeader", "", _oBEXAddTime4ByteHeaderErr)
	}
	return _oBEXAddTime4ByteHeader(time4Byte, dictRef), nil
}

// OBEXAddTime4ByteHeader add a CFStringRef to a dictionary of OBEXheaders.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXAddTime4ByteHeader(_:_:)
func OBEXAddTime4ByteHeader(time4Byte uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError {
	result, callErr := tryOBEXAddTime4ByteHeader(time4Byte, dictRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXAddTimeISOHeader func(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError
var _oBEXAddTimeISOHeaderErr error

func tryOBEXAddTimeISOHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) (OBEXError, error) {
	if _oBEXAddTimeISOHeader == nil {
		return *new(OBEXError), symbolCallError("OBEXAddTimeISOHeader", "", _oBEXAddTimeISOHeaderErr)
	}
	return _oBEXAddTimeISOHeader(inHeaderData, inHeaderDataLength, dictRef), nil
}

// OBEXAddTimeISOHeader add bytes to a dictionary of OBEXheaders.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXAddTimeISOHeader(_:_:_:)
func OBEXAddTimeISOHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError {
	result, callErr := tryOBEXAddTimeISOHeader(inHeaderData, inHeaderDataLength, dictRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXAddTypeHeader func(type_ corefoundation.CFStringRef, dictRef corefoundation.CFMutableDictionaryRef) OBEXError
var _oBEXAddTypeHeaderErr error

func tryOBEXAddTypeHeader(type_ corefoundation.CFStringRef, dictRef corefoundation.CFMutableDictionaryRef) (OBEXError, error) {
	if _oBEXAddTypeHeader == nil {
		return *new(OBEXError), symbolCallError("OBEXAddTypeHeader", "", _oBEXAddTypeHeaderErr)
	}
	return _oBEXAddTypeHeader(type_, dictRef), nil
}

// OBEXAddTypeHeader add a CFStringRef to a dictionary of OBEXheaders.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXAddTypeHeader(_:_:)
func OBEXAddTypeHeader(type_ corefoundation.CFStringRef, dictRef corefoundation.CFMutableDictionaryRef) OBEXError {
	result, callErr := tryOBEXAddTypeHeader(type_, dictRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXAddUserDefinedHeader func(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError
var _oBEXAddUserDefinedHeaderErr error

func tryOBEXAddUserDefinedHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) (OBEXError, error) {
	if _oBEXAddUserDefinedHeader == nil {
		return *new(OBEXError), symbolCallError("OBEXAddUserDefinedHeader", "", _oBEXAddUserDefinedHeaderErr)
	}
	return _oBEXAddUserDefinedHeader(inHeaderData, inHeaderDataLength, dictRef), nil
}

// OBEXAddUserDefinedHeader add a user-defined custom header to a dictionary of OBEXheaders.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXAddUserDefinedHeader(_:_:_:)
func OBEXAddUserDefinedHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError {
	result, callErr := tryOBEXAddUserDefinedHeader(inHeaderData, inHeaderDataLength, dictRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXAddWhoHeader func(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError
var _oBEXAddWhoHeaderErr error

func tryOBEXAddWhoHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) (OBEXError, error) {
	if _oBEXAddWhoHeader == nil {
		return *new(OBEXError), symbolCallError("OBEXAddWhoHeader", "", _oBEXAddWhoHeaderErr)
	}
	return _oBEXAddWhoHeader(inHeaderData, inHeaderDataLength, dictRef), nil
}

// OBEXAddWhoHeader add bytes of data to a dictionary of OBEXheaders.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXAddWhoHeader(_:_:_:)
func OBEXAddWhoHeader(inHeaderData unsafe.Pointer, inHeaderDataLength uint32, dictRef corefoundation.CFMutableDictionaryRef) OBEXError {
	result, callErr := tryOBEXAddWhoHeader(inHeaderData, inHeaderDataLength, dictRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXCreateVCard func(inFirstName unsafe.Pointer, inFirstNameLength uint32, inLastName unsafe.Pointer, inLastNameLength uint32, inFriendlyName unsafe.Pointer, inFriendlyNameLength uint32, inNameCharset unsafe.Pointer, inNameCharsetLength uint32, inHomePhone unsafe.Pointer, inHomePhoneLength uint32, inWorkPhone unsafe.Pointer, inWorkPhoneLength uint32, inCellPhone unsafe.Pointer, inCellPhoneLength uint32, inFaxPhone unsafe.Pointer, inFaxPhoneLength uint32, inEMailAddress unsafe.Pointer, inEMailAddressLength uint32, inEMailAddressCharset unsafe.Pointer, inEMailAddressCharsetLength uint32, inOrganization unsafe.Pointer, inOrganizationLength uint32, inOrganizationCharset unsafe.Pointer, inOrganizationCharsetLength uint32, inTitle unsafe.Pointer, inTitleLength uint32, inTitleCharset unsafe.Pointer, inTitleCharsetLength uint32) corefoundation.CFDataRef
var _oBEXCreateVCardErr error

func tryOBEXCreateVCard(inFirstName unsafe.Pointer, inFirstNameLength uint32, inLastName unsafe.Pointer, inLastNameLength uint32, inFriendlyName unsafe.Pointer, inFriendlyNameLength uint32, inNameCharset unsafe.Pointer, inNameCharsetLength uint32, inHomePhone unsafe.Pointer, inHomePhoneLength uint32, inWorkPhone unsafe.Pointer, inWorkPhoneLength uint32, inCellPhone unsafe.Pointer, inCellPhoneLength uint32, inFaxPhone unsafe.Pointer, inFaxPhoneLength uint32, inEMailAddress unsafe.Pointer, inEMailAddressLength uint32, inEMailAddressCharset unsafe.Pointer, inEMailAddressCharsetLength uint32, inOrganization unsafe.Pointer, inOrganizationLength uint32, inOrganizationCharset unsafe.Pointer, inOrganizationCharsetLength uint32, inTitle unsafe.Pointer, inTitleLength uint32, inTitleCharset unsafe.Pointer, inTitleCharsetLength uint32) (corefoundation.CFDataRef, error) {
	if _oBEXCreateVCard == nil {
		return *new(corefoundation.CFDataRef), symbolCallError("OBEXCreateVCard", "10.0", _oBEXCreateVCardErr)
	}
	return _oBEXCreateVCard(inFirstName, inFirstNameLength, inLastName, inLastNameLength, inFriendlyName, inFriendlyNameLength, inNameCharset, inNameCharsetLength, inHomePhone, inHomePhoneLength, inWorkPhone, inWorkPhoneLength, inCellPhone, inCellPhoneLength, inFaxPhone, inFaxPhoneLength, inEMailAddress, inEMailAddressLength, inEMailAddressCharset, inEMailAddressCharsetLength, inOrganization, inOrganizationLength, inOrganizationCharset, inOrganizationCharsetLength, inTitle, inTitleLength, inTitleCharset, inTitleCharsetLength), nil
}

// OBEXCreateVCard creates a formatted vCard, ready to be sent over OBEX or whatever.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXCreateVCard
func OBEXCreateVCard(inFirstName unsafe.Pointer, inFirstNameLength uint32, inLastName unsafe.Pointer, inLastNameLength uint32, inFriendlyName unsafe.Pointer, inFriendlyNameLength uint32, inNameCharset unsafe.Pointer, inNameCharsetLength uint32, inHomePhone unsafe.Pointer, inHomePhoneLength uint32, inWorkPhone unsafe.Pointer, inWorkPhoneLength uint32, inCellPhone unsafe.Pointer, inCellPhoneLength uint32, inFaxPhone unsafe.Pointer, inFaxPhoneLength uint32, inEMailAddress unsafe.Pointer, inEMailAddressLength uint32, inEMailAddressCharset unsafe.Pointer, inEMailAddressCharsetLength uint32, inOrganization unsafe.Pointer, inOrganizationLength uint32, inOrganizationCharset unsafe.Pointer, inOrganizationCharsetLength uint32, inTitle unsafe.Pointer, inTitleLength uint32, inTitleCharset unsafe.Pointer, inTitleCharsetLength uint32) corefoundation.CFDataRef {
	result, callErr := tryOBEXCreateVCard(inFirstName, inFirstNameLength, inLastName, inLastNameLength, inFriendlyName, inFriendlyNameLength, inNameCharset, inNameCharsetLength, inHomePhone, inHomePhoneLength, inWorkPhone, inWorkPhoneLength, inCellPhone, inCellPhoneLength, inFaxPhone, inFaxPhoneLength, inEMailAddress, inEMailAddressLength, inEMailAddressCharset, inEMailAddressCharsetLength, inOrganization, inOrganizationLength, inOrganizationCharset, inOrganizationCharsetLength, inTitle, inTitleLength, inTitleCharset, inTitleCharsetLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXCreateVEvent func(inCharset string, inCharsetLength uint32, inEncoding string, inEncodingLength uint32, inEventStartDate string, inEventStartDateLength uint32, inEventEndDate string, inEventEndDateLength uint32, inAlarmDate string, inAlarmDateLength uint32, inCategory string, inCategoryLength uint32, inSummary string, inSummaryLength uint32, inLocation string, inLocationLength uint32, inXIRMCLUID string, inXIRMCLUIDLength uint32) corefoundation.CFDataRef
var _oBEXCreateVEventErr error

func tryOBEXCreateVEvent(inCharset string, inCharsetLength uint32, inEncoding string, inEncodingLength uint32, inEventStartDate string, inEventStartDateLength uint32, inEventEndDate string, inEventEndDateLength uint32, inAlarmDate string, inAlarmDateLength uint32, inCategory string, inCategoryLength uint32, inSummary string, inSummaryLength uint32, inLocation string, inLocationLength uint32, inXIRMCLUID string, inXIRMCLUIDLength uint32) (corefoundation.CFDataRef, error) {
	if _oBEXCreateVEvent == nil {
		return *new(corefoundation.CFDataRef), symbolCallError("OBEXCreateVEvent", "10.0", _oBEXCreateVEventErr)
	}
	return _oBEXCreateVEvent(inCharset, inCharsetLength, inEncoding, inEncodingLength, inEventStartDate, inEventStartDateLength, inEventEndDate, inEventEndDateLength, inAlarmDate, inAlarmDateLength, inCategory, inCategoryLength, inSummary, inSummaryLength, inLocation, inLocationLength, inXIRMCLUID, inXIRMCLUIDLength), nil
}

// OBEXCreateVEvent creates a formatted vEvent, ready to be sent over OBEX or whatever.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXCreateVEvent
func OBEXCreateVEvent(inCharset string, inCharsetLength uint32, inEncoding string, inEncodingLength uint32, inEventStartDate string, inEventStartDateLength uint32, inEventEndDate string, inEventEndDateLength uint32, inAlarmDate string, inAlarmDateLength uint32, inCategory string, inCategoryLength uint32, inSummary string, inSummaryLength uint32, inLocation string, inLocationLength uint32, inXIRMCLUID string, inXIRMCLUIDLength uint32) corefoundation.CFDataRef {
	result, callErr := tryOBEXCreateVEvent(inCharset, inCharsetLength, inEncoding, inEncodingLength, inEventStartDate, inEventStartDateLength, inEventEndDate, inEventEndDateLength, inAlarmDate, inAlarmDateLength, inCategory, inCategoryLength, inSummary, inSummaryLength, inLocation, inLocationLength, inXIRMCLUID, inXIRMCLUIDLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXGetHeaders func(inData unsafe.Pointer, inDataSize uintptr) corefoundation.CFDictionaryRef
var _oBEXGetHeadersErr error

func tryOBEXGetHeaders(inData unsafe.Pointer, inDataSize uintptr) (corefoundation.CFDictionaryRef, error) {
	if _oBEXGetHeaders == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("OBEXGetHeaders", "", _oBEXGetHeadersErr)
	}
	return _oBEXGetHeaders(inData, inDataSize), nil
}

// OBEXGetHeaders take a data blob and looks for OBEX headers.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXGetHeaders(_:_:)
func OBEXGetHeaders(inData unsafe.Pointer, inDataSize uintptr) corefoundation.CFDictionaryRef {
	result, callErr := tryOBEXGetHeaders(inData, inDataSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXHeadersToBytes func(dictionaryOfHeaders corefoundation.CFDictionaryRef) corefoundation.CFMutableDataRef
var _oBEXHeadersToBytesErr error

func tryOBEXHeadersToBytes(dictionaryOfHeaders corefoundation.CFDictionaryRef) (corefoundation.CFMutableDataRef, error) {
	if _oBEXHeadersToBytes == nil {
		return *new(corefoundation.CFMutableDataRef), symbolCallError("OBEXHeadersToBytes", "", _oBEXHeadersToBytesErr)
	}
	return _oBEXHeadersToBytes(dictionaryOfHeaders), nil
}

// OBEXHeadersToBytes converts a dictionary of headers to a data pointer, from which you can extract as bytes and pass to the OBEX command/response functions.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXHeadersToBytes(_:)
func OBEXHeadersToBytes(dictionaryOfHeaders corefoundation.CFDictionaryRef) corefoundation.CFMutableDataRef {
	result, callErr := tryOBEXHeadersToBytes(dictionaryOfHeaders)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXSessionAbort func(inSessionRef OBEXSessionRef, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError
var _oBEXSessionAbortErr error

func tryOBEXSessionAbort(inSessionRef OBEXSessionRef, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) (OBEXError, error) {
	if _oBEXSessionAbort == nil {
		return *new(OBEXError), symbolCallError("OBEXSessionAbort", "10.0", _oBEXSessionAbortErr)
	}
	return _oBEXSessionAbort(inSessionRef, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon), nil
}

// OBEXSessionAbort send an abort command to a remote OBEX server.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionAbort
func OBEXSessionAbort(inSessionRef OBEXSessionRef, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError {
	result, callErr := tryOBEXSessionAbort(inSessionRef, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXSessionAbortResponse func(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError
var _oBEXSessionAbortResponseErr error

func tryOBEXSessionAbortResponse(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) (OBEXError, error) {
	if _oBEXSessionAbortResponse == nil {
		return *new(OBEXError), symbolCallError("OBEXSessionAbortResponse", "10.0", _oBEXSessionAbortResponseErr)
	}
	return _oBEXSessionAbortResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon), nil
}

// OBEXSessionAbortResponse send a response to a abort command to the remote client.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionAbortResponse
func OBEXSessionAbortResponse(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError {
	result, callErr := tryOBEXSessionAbortResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXSessionConnect func(inSessionRef OBEXSessionRef, inFlags OBEXFlags, inMaxPacketLength OBEXMaxPacketLength, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError
var _oBEXSessionConnectErr error

func tryOBEXSessionConnect(inSessionRef OBEXSessionRef, inFlags OBEXFlags, inMaxPacketLength OBEXMaxPacketLength, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) (OBEXError, error) {
	if _oBEXSessionConnect == nil {
		return *new(OBEXError), symbolCallError("OBEXSessionConnect", "10.0", _oBEXSessionConnectErr)
	}
	return _oBEXSessionConnect(inSessionRef, inFlags, inMaxPacketLength, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon), nil
}

// OBEXSessionConnect establishes an OBEX connection to the target device for the session.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionConnect
func OBEXSessionConnect(inSessionRef OBEXSessionRef, inFlags OBEXFlags, inMaxPacketLength OBEXMaxPacketLength, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError {
	result, callErr := tryOBEXSessionConnect(inSessionRef, inFlags, inMaxPacketLength, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXSessionConnectResponse func(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inFlags OBEXFlags, inMaxPacketLength OBEXMaxPacketLength, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError
var _oBEXSessionConnectResponseErr error

func tryOBEXSessionConnectResponse(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inFlags OBEXFlags, inMaxPacketLength OBEXMaxPacketLength, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) (OBEXError, error) {
	if _oBEXSessionConnectResponse == nil {
		return *new(OBEXError), symbolCallError("OBEXSessionConnectResponse", "10.0", _oBEXSessionConnectResponseErr)
	}
	return _oBEXSessionConnectResponse(inSessionRef, inResponseOpCode, inFlags, inMaxPacketLength, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon), nil
}

// OBEXSessionConnectResponse send a response to a connect command to the remote client.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionConnectResponse
func OBEXSessionConnectResponse(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inFlags OBEXFlags, inMaxPacketLength OBEXMaxPacketLength, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError {
	result, callErr := tryOBEXSessionConnectResponse(inSessionRef, inResponseOpCode, inFlags, inMaxPacketLength, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXSessionDelete func(inSessionRef OBEXSessionRef) OBEXError
var _oBEXSessionDeleteErr error

func tryOBEXSessionDelete(inSessionRef OBEXSessionRef) (OBEXError, error) {
	if _oBEXSessionDelete == nil {
		return *new(OBEXError), symbolCallError("OBEXSessionDelete", "10.0", _oBEXSessionDeleteErr)
	}
	return _oBEXSessionDelete(inSessionRef), nil
}

// OBEXSessionDelete destroy an OBEX session.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionDelete
func OBEXSessionDelete(inSessionRef OBEXSessionRef) OBEXError {
	result, callErr := tryOBEXSessionDelete(inSessionRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXSessionDisconnect func(inSessionRef OBEXSessionRef, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError
var _oBEXSessionDisconnectErr error

func tryOBEXSessionDisconnect(inSessionRef OBEXSessionRef, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) (OBEXError, error) {
	if _oBEXSessionDisconnect == nil {
		return *new(OBEXError), symbolCallError("OBEXSessionDisconnect", "10.0", _oBEXSessionDisconnectErr)
	}
	return _oBEXSessionDisconnect(inSessionRef, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon), nil
}

// OBEXSessionDisconnect send a disconnect command to a remote OBEX server.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionDisconnect
func OBEXSessionDisconnect(inSessionRef OBEXSessionRef, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError {
	result, callErr := tryOBEXSessionDisconnect(inSessionRef, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXSessionDisconnectResponse func(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError
var _oBEXSessionDisconnectResponseErr error

func tryOBEXSessionDisconnectResponse(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) (OBEXError, error) {
	if _oBEXSessionDisconnectResponse == nil {
		return *new(OBEXError), symbolCallError("OBEXSessionDisconnectResponse", "10.0", _oBEXSessionDisconnectResponseErr)
	}
	return _oBEXSessionDisconnectResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon), nil
}

// OBEXSessionDisconnectResponse send a response to a disconnect command to the remote client.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionDisconnectResponse
func OBEXSessionDisconnectResponse(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError {
	result, callErr := tryOBEXSessionDisconnectResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXSessionGet func(inSessionRef OBEXSessionRef, inIsFinalChunk bool, inHeadersData unsafe.Pointer, inHeadersDataLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError
var _oBEXSessionGetErr error

func tryOBEXSessionGet(inSessionRef OBEXSessionRef, inIsFinalChunk bool, inHeadersData unsafe.Pointer, inHeadersDataLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) (OBEXError, error) {
	if _oBEXSessionGet == nil {
		return *new(OBEXError), symbolCallError("OBEXSessionGet", "10.0", _oBEXSessionGetErr)
	}
	return _oBEXSessionGet(inSessionRef, inIsFinalChunk, inHeadersData, inHeadersDataLength, inCallback, inUserRefCon), nil
}

// OBEXSessionGet send a get command to a remote OBEX server.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionGet
func OBEXSessionGet(inSessionRef OBEXSessionRef, inIsFinalChunk bool, inHeadersData unsafe.Pointer, inHeadersDataLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError {
	result, callErr := tryOBEXSessionGet(inSessionRef, inIsFinalChunk, inHeadersData, inHeadersDataLength, inCallback, inUserRefCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXSessionGetAvailableCommandPayloadLength func(inSessionRef OBEXSessionRef, inOpCode OBEXOpCode, outLength *OBEXMaxPacketLength) OBEXError
var _oBEXSessionGetAvailableCommandPayloadLengthErr error

func tryOBEXSessionGetAvailableCommandPayloadLength(inSessionRef OBEXSessionRef, inOpCode OBEXOpCode, outLength *OBEXMaxPacketLength) (OBEXError, error) {
	if _oBEXSessionGetAvailableCommandPayloadLength == nil {
		return *new(OBEXError), symbolCallError("OBEXSessionGetAvailableCommandPayloadLength", "10.0", _oBEXSessionGetAvailableCommandPayloadLengthErr)
	}
	return _oBEXSessionGetAvailableCommandPayloadLength(inSessionRef, inOpCode, outLength), nil
}

// OBEXSessionGetAvailableCommandPayloadLength gets space available for your data for a particular command response you are trying to send.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionGetAvailableCommandPayloadLength
func OBEXSessionGetAvailableCommandPayloadLength(inSessionRef OBEXSessionRef, inOpCode OBEXOpCode, outLength *OBEXMaxPacketLength) OBEXError {
	result, callErr := tryOBEXSessionGetAvailableCommandPayloadLength(inSessionRef, inOpCode, outLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXSessionGetAvailableCommandResponsePayloadLength func(inSessionRef OBEXSessionRef, inOpCode OBEXOpCode, outLength *OBEXMaxPacketLength) OBEXError
var _oBEXSessionGetAvailableCommandResponsePayloadLengthErr error

func tryOBEXSessionGetAvailableCommandResponsePayloadLength(inSessionRef OBEXSessionRef, inOpCode OBEXOpCode, outLength *OBEXMaxPacketLength) (OBEXError, error) {
	if _oBEXSessionGetAvailableCommandResponsePayloadLength == nil {
		return *new(OBEXError), symbolCallError("OBEXSessionGetAvailableCommandResponsePayloadLength", "10.0", _oBEXSessionGetAvailableCommandResponsePayloadLengthErr)
	}
	return _oBEXSessionGetAvailableCommandResponsePayloadLength(inSessionRef, inOpCode, outLength), nil
}

// OBEXSessionGetAvailableCommandResponsePayloadLength gets space available for your data for a particular command response you are trying to send.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionGetAvailableCommandResponsePayloadLength
func OBEXSessionGetAvailableCommandResponsePayloadLength(inSessionRef OBEXSessionRef, inOpCode OBEXOpCode, outLength *OBEXMaxPacketLength) OBEXError {
	result, callErr := tryOBEXSessionGetAvailableCommandResponsePayloadLength(inSessionRef, inOpCode, outLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXSessionGetMaxPacketLength func(inSessionRef OBEXSessionRef, outLength *OBEXMaxPacketLength) OBEXError
var _oBEXSessionGetMaxPacketLengthErr error

func tryOBEXSessionGetMaxPacketLength(inSessionRef OBEXSessionRef, outLength *OBEXMaxPacketLength) (OBEXError, error) {
	if _oBEXSessionGetMaxPacketLength == nil {
		return *new(OBEXError), symbolCallError("OBEXSessionGetMaxPacketLength", "10.0", _oBEXSessionGetMaxPacketLengthErr)
	}
	return _oBEXSessionGetMaxPacketLength(inSessionRef, outLength), nil
}

// OBEXSessionGetMaxPacketLength gets current max packet length.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionGetMaxPacketLength
func OBEXSessionGetMaxPacketLength(inSessionRef OBEXSessionRef, outLength *OBEXMaxPacketLength) OBEXError {
	result, callErr := tryOBEXSessionGetMaxPacketLength(inSessionRef, outLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXSessionGetResponse func(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError
var _oBEXSessionGetResponseErr error

func tryOBEXSessionGetResponse(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) (OBEXError, error) {
	if _oBEXSessionGetResponse == nil {
		return *new(OBEXError), symbolCallError("OBEXSessionGetResponse", "10.0", _oBEXSessionGetResponseErr)
	}
	return _oBEXSessionGetResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon), nil
}

// OBEXSessionGetResponse send a response to a get command to the remote client.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionGetResponse
func OBEXSessionGetResponse(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError {
	result, callErr := tryOBEXSessionGetResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXSessionHasOpenOBEXConnection func(inSessionRef OBEXSessionRef, outIsConnected *bool) OBEXError
var _oBEXSessionHasOpenOBEXConnectionErr error

func tryOBEXSessionHasOpenOBEXConnection(inSessionRef OBEXSessionRef, outIsConnected *bool) (OBEXError, error) {
	if _oBEXSessionHasOpenOBEXConnection == nil {
		return *new(OBEXError), symbolCallError("OBEXSessionHasOpenOBEXConnection", "10.0", _oBEXSessionHasOpenOBEXConnectionErr)
	}
	return _oBEXSessionHasOpenOBEXConnection(inSessionRef, outIsConnected), nil
}

// OBEXSessionHasOpenOBEXConnection allows you to test the session for an open OBEX connection for a particular session.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionHasOpenOBEXConnection
func OBEXSessionHasOpenOBEXConnection(inSessionRef OBEXSessionRef, outIsConnected *bool) OBEXError {
	result, callErr := tryOBEXSessionHasOpenOBEXConnection(inSessionRef, outIsConnected)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXSessionPut func(inSessionRef OBEXSessionRef, inIsFinalChunk bool, inHeadersData unsafe.Pointer, inHeadersDataLength uintptr, inBodyData unsafe.Pointer, inBodyDataLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError
var _oBEXSessionPutErr error

func tryOBEXSessionPut(inSessionRef OBEXSessionRef, inIsFinalChunk bool, inHeadersData unsafe.Pointer, inHeadersDataLength uintptr, inBodyData unsafe.Pointer, inBodyDataLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) (OBEXError, error) {
	if _oBEXSessionPut == nil {
		return *new(OBEXError), symbolCallError("OBEXSessionPut", "10.0", _oBEXSessionPutErr)
	}
	return _oBEXSessionPut(inSessionRef, inIsFinalChunk, inHeadersData, inHeadersDataLength, inBodyData, inBodyDataLength, inCallback, inUserRefCon), nil
}

// OBEXSessionPut send a put command to a remote OBEX server.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionPut
func OBEXSessionPut(inSessionRef OBEXSessionRef, inIsFinalChunk bool, inHeadersData unsafe.Pointer, inHeadersDataLength uintptr, inBodyData unsafe.Pointer, inBodyDataLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError {
	result, callErr := tryOBEXSessionPut(inSessionRef, inIsFinalChunk, inHeadersData, inHeadersDataLength, inBodyData, inBodyDataLength, inCallback, inUserRefCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXSessionPutResponse func(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError
var _oBEXSessionPutResponseErr error

func tryOBEXSessionPutResponse(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) (OBEXError, error) {
	if _oBEXSessionPutResponse == nil {
		return *new(OBEXError), symbolCallError("OBEXSessionPutResponse", "10.0", _oBEXSessionPutResponseErr)
	}
	return _oBEXSessionPutResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon), nil
}

// OBEXSessionPutResponse send a response to a put command to the remote client.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionPutResponse
func OBEXSessionPutResponse(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError {
	result, callErr := tryOBEXSessionPutResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXSessionSetPath func(inSessionRef OBEXSessionRef, inFlags OBEXFlags, inConstants OBEXConstants, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError
var _oBEXSessionSetPathErr error

func tryOBEXSessionSetPath(inSessionRef OBEXSessionRef, inFlags OBEXFlags, inConstants OBEXConstants, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) (OBEXError, error) {
	if _oBEXSessionSetPath == nil {
		return *new(OBEXError), symbolCallError("OBEXSessionSetPath", "10.0", _oBEXSessionSetPathErr)
	}
	return _oBEXSessionSetPath(inSessionRef, inFlags, inConstants, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon), nil
}

// OBEXSessionSetPath send a set path command to a remote OBEX server.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionSetPath
func OBEXSessionSetPath(inSessionRef OBEXSessionRef, inFlags OBEXFlags, inConstants OBEXConstants, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError {
	result, callErr := tryOBEXSessionSetPath(inSessionRef, inFlags, inConstants, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXSessionSetPathResponse func(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError
var _oBEXSessionSetPathResponseErr error

func tryOBEXSessionSetPathResponse(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) (OBEXError, error) {
	if _oBEXSessionSetPathResponse == nil {
		return *new(OBEXError), symbolCallError("OBEXSessionSetPathResponse", "10.0", _oBEXSessionSetPathResponseErr)
	}
	return _oBEXSessionSetPathResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon), nil
}

// OBEXSessionSetPathResponse send a response to a set path command to the remote client.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionSetPathResponse
func OBEXSessionSetPathResponse(inSessionRef OBEXSessionRef, inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError {
	result, callErr := tryOBEXSessionSetPathResponse(inSessionRef, inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inCallback, inUserRefCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oBEXSessionSetServerCallback func(inSessionRef OBEXSessionRef, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError
var _oBEXSessionSetServerCallbackErr error

func tryOBEXSessionSetServerCallback(inSessionRef OBEXSessionRef, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) (OBEXError, error) {
	if _oBEXSessionSetServerCallback == nil {
		return *new(OBEXError), symbolCallError("OBEXSessionSetServerCallback", "10.0", _oBEXSessionSetServerCallbackErr)
	}
	return _oBEXSessionSetServerCallback(inSessionRef, inCallback, inUserRefCon), nil
}

// OBEXSessionSetServerCallback.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionSetServerCallback
func OBEXSessionSetServerCallback(inSessionRef OBEXSessionRef, inCallback OBEXSessionEventCallback, inUserRefCon uintptr) OBEXError {
	result, callErr := tryOBEXSessionSetServerCallback(inSessionRef, inCallback, inUserRefCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_iOBluetoothAddSCOAudioDevice, &_iOBluetoothAddSCOAudioDeviceErr, frameworkHandle, "IOBluetoothAddSCOAudioDevice", "10.0")
	registerFunc(&_iOBluetoothFindNumberOfRegistryEntriesOfClassName, &_iOBluetoothFindNumberOfRegistryEntriesOfClassNameErr, frameworkHandle, "IOBluetoothFindNumberOfRegistryEntriesOfClassName", "")
	registerFunc(&_iOBluetoothGetUniqueFileNameAndPath, &_iOBluetoothGetUniqueFileNameAndPathErr, frameworkHandle, "IOBluetoothGetUniqueFileNameAndPath", "")
	registerFunc(&_iOBluetoothIgnoreHIDDevice, &_iOBluetoothIgnoreHIDDeviceErr, frameworkHandle, "IOBluetoothIgnoreHIDDevice", "")
	registerFunc(&_iOBluetoothIsFileAppleDesignatedPIMData, &_iOBluetoothIsFileAppleDesignatedPIMDataErr, frameworkHandle, "IOBluetoothIsFileAppleDesignatedPIMData", "")
	registerFunc(&_iOBluetoothL2CAPChannelRegisterForChannelCloseNotification, &_iOBluetoothL2CAPChannelRegisterForChannelCloseNotificationErr, frameworkHandle, "IOBluetoothL2CAPChannelRegisterForChannelCloseNotification", "")
	registerFunc(&_iOBluetoothNSStringFromDeviceAddress, &_iOBluetoothNSStringFromDeviceAddressErr, frameworkHandle, "IOBluetoothNSStringFromDeviceAddress", "")
	registerFunc(&_iOBluetoothNSStringFromDeviceAddressColon, &_iOBluetoothNSStringFromDeviceAddressColonErr, frameworkHandle, "IOBluetoothNSStringFromDeviceAddressColon", "")
	registerFunc(&_iOBluetoothNSStringToDeviceAddress, &_iOBluetoothNSStringToDeviceAddressErr, frameworkHandle, "IOBluetoothNSStringToDeviceAddress", "")
	registerFunc(&_iOBluetoothNumberOfAvailableHIDDevices, &_iOBluetoothNumberOfAvailableHIDDevicesErr, frameworkHandle, "IOBluetoothNumberOfAvailableHIDDevices", "")
	registerFunc(&_iOBluetoothNumberOfKeyboardHIDDevices, &_iOBluetoothNumberOfKeyboardHIDDevicesErr, frameworkHandle, "IOBluetoothNumberOfKeyboardHIDDevices", "")
	registerFunc(&_iOBluetoothNumberOfPointingHIDDevices, &_iOBluetoothNumberOfPointingHIDDevicesErr, frameworkHandle, "IOBluetoothNumberOfPointingHIDDevices", "")
	registerFunc(&_iOBluetoothNumberOfTabletHIDDevices, &_iOBluetoothNumberOfTabletHIDDevicesErr, frameworkHandle, "IOBluetoothNumberOfTabletHIDDevices", "")
	registerFunc(&_iOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber, &_iOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumberErr, frameworkHandle, "IOBluetoothOBEXSessionCreateWithIOBluetoothDeviceRefAndChannelNumber", "10.0")
	registerFunc(&_iOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef, &_iOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRefErr, frameworkHandle, "IOBluetoothOBEXSessionCreateWithIOBluetoothSDPServiceRecordRef", "10.0")
	registerFunc(&_iOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel, &_iOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannelErr, frameworkHandle, "IOBluetoothOBEXSessionCreateWithIncomingIOBluetoothRFCOMMChannel", "10.0")
	registerFunc(&_iOBluetoothOBEXSessionOpenTransportConnection, &_iOBluetoothOBEXSessionOpenTransportConnectionErr, frameworkHandle, "IOBluetoothOBEXSessionOpenTransportConnection", "10.0")
	registerFunc(&_iOBluetoothPackData, &_iOBluetoothPackDataErr, frameworkHandle, "IOBluetoothPackData", "")
	registerFunc(&_iOBluetoothPackDataList, &_iOBluetoothPackDataListErr, frameworkHandle, "IOBluetoothPackDataList", "")
	registerFunc(&_iOBluetoothRemoveIgnoredHIDDevice, &_iOBluetoothRemoveIgnoredHIDDeviceErr, frameworkHandle, "IOBluetoothRemoveIgnoredHIDDevice", "")
	registerFunc(&_iOBluetoothRemoveSCOAudioDevice, &_iOBluetoothRemoveSCOAudioDeviceErr, frameworkHandle, "IOBluetoothRemoveSCOAudioDevice", "10.0")
	registerFunc(&_iOBluetoothUnpackData, &_iOBluetoothUnpackDataErr, frameworkHandle, "IOBluetoothUnpackData", "")
	registerFunc(&_iOBluetoothUnpackDataList, &_iOBluetoothUnpackDataListErr, frameworkHandle, "IOBluetoothUnpackDataList", "")
	registerFunc(&_iOBluetoothUserNotificationUnregister, &_iOBluetoothUserNotificationUnregisterErr, frameworkHandle, "IOBluetoothUserNotificationUnregister", "")
	registerFunc(&_oBEXAddApplicationParameterHeader, &_oBEXAddApplicationParameterHeaderErr, frameworkHandle, "OBEXAddApplicationParameterHeader", "")
	registerFunc(&_oBEXAddAuthorizationChallengeHeader, &_oBEXAddAuthorizationChallengeHeaderErr, frameworkHandle, "OBEXAddAuthorizationChallengeHeader", "")
	registerFunc(&_oBEXAddAuthorizationResponseHeader, &_oBEXAddAuthorizationResponseHeaderErr, frameworkHandle, "OBEXAddAuthorizationResponseHeader", "")
	registerFunc(&_oBEXAddBodyHeader, &_oBEXAddBodyHeaderErr, frameworkHandle, "OBEXAddBodyHeader", "")
	registerFunc(&_oBEXAddByteSequenceHeader, &_oBEXAddByteSequenceHeaderErr, frameworkHandle, "OBEXAddByteSequenceHeader", "")
	registerFunc(&_oBEXAddConnectionIDHeader, &_oBEXAddConnectionIDHeaderErr, frameworkHandle, "OBEXAddConnectionIDHeader", "")
	registerFunc(&_oBEXAddCountHeader, &_oBEXAddCountHeaderErr, frameworkHandle, "OBEXAddCountHeader", "")
	registerFunc(&_oBEXAddDescriptionHeader, &_oBEXAddDescriptionHeaderErr, frameworkHandle, "OBEXAddDescriptionHeader", "")
	registerFunc(&_oBEXAddHTTPHeader, &_oBEXAddHTTPHeaderErr, frameworkHandle, "OBEXAddHTTPHeader", "")
	registerFunc(&_oBEXAddLengthHeader, &_oBEXAddLengthHeaderErr, frameworkHandle, "OBEXAddLengthHeader", "")
	registerFunc(&_oBEXAddNameHeader, &_oBEXAddNameHeaderErr, frameworkHandle, "OBEXAddNameHeader", "")
	registerFunc(&_oBEXAddObjectClassHeader, &_oBEXAddObjectClassHeaderErr, frameworkHandle, "OBEXAddObjectClassHeader", "")
	registerFunc(&_oBEXAddTargetHeader, &_oBEXAddTargetHeaderErr, frameworkHandle, "OBEXAddTargetHeader", "")
	registerFunc(&_oBEXAddTime4ByteHeader, &_oBEXAddTime4ByteHeaderErr, frameworkHandle, "OBEXAddTime4ByteHeader", "")
	registerFunc(&_oBEXAddTimeISOHeader, &_oBEXAddTimeISOHeaderErr, frameworkHandle, "OBEXAddTimeISOHeader", "")
	registerFunc(&_oBEXAddTypeHeader, &_oBEXAddTypeHeaderErr, frameworkHandle, "OBEXAddTypeHeader", "")
	registerFunc(&_oBEXAddUserDefinedHeader, &_oBEXAddUserDefinedHeaderErr, frameworkHandle, "OBEXAddUserDefinedHeader", "")
	registerFunc(&_oBEXAddWhoHeader, &_oBEXAddWhoHeaderErr, frameworkHandle, "OBEXAddWhoHeader", "")
	registerFunc(&_oBEXCreateVCard, &_oBEXCreateVCardErr, frameworkHandle, "OBEXCreateVCard", "10.0")
	registerFunc(&_oBEXCreateVEvent, &_oBEXCreateVEventErr, frameworkHandle, "OBEXCreateVEvent", "10.0")
	registerFunc(&_oBEXGetHeaders, &_oBEXGetHeadersErr, frameworkHandle, "OBEXGetHeaders", "")
	registerFunc(&_oBEXHeadersToBytes, &_oBEXHeadersToBytesErr, frameworkHandle, "OBEXHeadersToBytes", "")
	registerFunc(&_oBEXSessionAbort, &_oBEXSessionAbortErr, frameworkHandle, "OBEXSessionAbort", "10.0")
	registerFunc(&_oBEXSessionAbortResponse, &_oBEXSessionAbortResponseErr, frameworkHandle, "OBEXSessionAbortResponse", "10.0")
	registerFunc(&_oBEXSessionConnect, &_oBEXSessionConnectErr, frameworkHandle, "OBEXSessionConnect", "10.0")
	registerFunc(&_oBEXSessionConnectResponse, &_oBEXSessionConnectResponseErr, frameworkHandle, "OBEXSessionConnectResponse", "10.0")
	registerFunc(&_oBEXSessionDelete, &_oBEXSessionDeleteErr, frameworkHandle, "OBEXSessionDelete", "10.0")
	registerFunc(&_oBEXSessionDisconnect, &_oBEXSessionDisconnectErr, frameworkHandle, "OBEXSessionDisconnect", "10.0")
	registerFunc(&_oBEXSessionDisconnectResponse, &_oBEXSessionDisconnectResponseErr, frameworkHandle, "OBEXSessionDisconnectResponse", "10.0")
	registerFunc(&_oBEXSessionGet, &_oBEXSessionGetErr, frameworkHandle, "OBEXSessionGet", "10.0")
	registerFunc(&_oBEXSessionGetAvailableCommandPayloadLength, &_oBEXSessionGetAvailableCommandPayloadLengthErr, frameworkHandle, "OBEXSessionGetAvailableCommandPayloadLength", "10.0")
	registerFunc(&_oBEXSessionGetAvailableCommandResponsePayloadLength, &_oBEXSessionGetAvailableCommandResponsePayloadLengthErr, frameworkHandle, "OBEXSessionGetAvailableCommandResponsePayloadLength", "10.0")
	registerFunc(&_oBEXSessionGetMaxPacketLength, &_oBEXSessionGetMaxPacketLengthErr, frameworkHandle, "OBEXSessionGetMaxPacketLength", "10.0")
	registerFunc(&_oBEXSessionGetResponse, &_oBEXSessionGetResponseErr, frameworkHandle, "OBEXSessionGetResponse", "10.0")
	registerFunc(&_oBEXSessionHasOpenOBEXConnection, &_oBEXSessionHasOpenOBEXConnectionErr, frameworkHandle, "OBEXSessionHasOpenOBEXConnection", "10.0")
	registerFunc(&_oBEXSessionPut, &_oBEXSessionPutErr, frameworkHandle, "OBEXSessionPut", "10.0")
	registerFunc(&_oBEXSessionPutResponse, &_oBEXSessionPutResponseErr, frameworkHandle, "OBEXSessionPutResponse", "10.0")
	registerFunc(&_oBEXSessionSetPath, &_oBEXSessionSetPathErr, frameworkHandle, "OBEXSessionSetPath", "10.0")
	registerFunc(&_oBEXSessionSetPathResponse, &_oBEXSessionSetPathResponseErr, frameworkHandle, "OBEXSessionSetPathResponse", "10.0")
	registerFunc(&_oBEXSessionSetServerCallback, &_oBEXSessionSetServerCallbackErr, frameworkHandle, "OBEXSessionSetServerCallback", "10.0")
}

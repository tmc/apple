// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"fmt"

	"github.com/ebitengine/purego"
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
		return fmt.Sprintf("IOUSBHost: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("IOUSBHost: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("IOUSBHost: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("IOUSBHost: register symbol %s: %v", name, r)
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

var _iOUSBGetBillboardDescriptor func(bosDescriptor uintptr) uintptr
var _iOUSBGetBillboardDescriptorErr error

func tryIOUSBGetBillboardDescriptor(bosDescriptor uintptr) (uintptr, error) {
	if _iOUSBGetBillboardDescriptor == nil {
		return 0, symbolCallError("IOUSBGetBillboardDescriptor", "10.15", _iOUSBGetBillboardDescriptorErr)
	}
	return _iOUSBGetBillboardDescriptor(bosDescriptor), nil
}

// IOUSBGetBillboardDescriptor obtains the first billboard capability descriptor in a BOS descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetBillboardDescriptor(_:)
func IOUSBGetBillboardDescriptor(bosDescriptor uintptr) uintptr {
	result, callErr := tryIOUSBGetBillboardDescriptor(bosDescriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetConfigurationMaxPowerMilliAmps func(usbDeviceSpeed uint32, descriptor uintptr) uint32
var _iOUSBGetConfigurationMaxPowerMilliAmpsErr error

func tryIOUSBGetConfigurationMaxPowerMilliAmps(usbDeviceSpeed uint32, descriptor uintptr) (uint32, error) {
	if _iOUSBGetConfigurationMaxPowerMilliAmps == nil {
		return 0, symbolCallError("IOUSBGetConfigurationMaxPowerMilliAmps", "10.15", _iOUSBGetConfigurationMaxPowerMilliAmpsErr)
	}
	return _iOUSBGetConfigurationMaxPowerMilliAmps(usbDeviceSpeed, descriptor), nil
}

// IOUSBGetConfigurationMaxPowerMilliAmps obtains the maximum bus current that a configuration descriptor requires.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetConfigurationMaxPowerMilliAmps(_:_:)
func IOUSBGetConfigurationMaxPowerMilliAmps(usbDeviceSpeed uint32, descriptor uintptr) uint32 {
	result, callErr := tryIOUSBGetConfigurationMaxPowerMilliAmps(usbDeviceSpeed, descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetContainerIDDescriptor func(bosDescriptor uintptr) string
var _iOUSBGetContainerIDDescriptorErr error

func tryIOUSBGetContainerIDDescriptor(bosDescriptor uintptr) (string, error) {
	if _iOUSBGetContainerIDDescriptor == nil {
		return "", symbolCallError("IOUSBGetContainerIDDescriptor", "10.15", _iOUSBGetContainerIDDescriptorErr)
	}
	return _iOUSBGetContainerIDDescriptor(bosDescriptor), nil
}

// IOUSBGetContainerIDDescriptor obtains the first container ID capability descriptor in a BOS descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetContainerIDDescriptor(_:)
func IOUSBGetContainerIDDescriptor(bosDescriptor uintptr) string {
	result, callErr := tryIOUSBGetContainerIDDescriptor(bosDescriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetEndpointAddress func(descriptor uintptr) uint8
var _iOUSBGetEndpointAddressErr error

func tryIOUSBGetEndpointAddress(descriptor uintptr) (uint8, error) {
	if _iOUSBGetEndpointAddress == nil {
		return 0, symbolCallError("IOUSBGetEndpointAddress", "10.15", _iOUSBGetEndpointAddressErr)
	}
	return _iOUSBGetEndpointAddress(descriptor), nil
}

// IOUSBGetEndpointAddress obtains the direction and number of an endpoint from an endpoint descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointAddress(_:)
func IOUSBGetEndpointAddress(descriptor uintptr) uint8 {
	result, callErr := tryIOUSBGetEndpointAddress(descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetEndpointBurstSize func(usbDeviceSpeed uint32, descriptor uintptr, companionDescriptor uintptr, sspCompanionDescriptor uintptr) uint32
var _iOUSBGetEndpointBurstSizeErr error

func tryIOUSBGetEndpointBurstSize(usbDeviceSpeed uint32, descriptor uintptr, companionDescriptor uintptr, sspCompanionDescriptor uintptr) (uint32, error) {
	if _iOUSBGetEndpointBurstSize == nil {
		return 0, symbolCallError("IOUSBGetEndpointBurstSize", "10.15", _iOUSBGetEndpointBurstSizeErr)
	}
	return _iOUSBGetEndpointBurstSize(usbDeviceSpeed, descriptor, companionDescriptor, sspCompanionDescriptor), nil
}

// IOUSBGetEndpointBurstSize.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointBurstSize(_:_:_:_:)
func IOUSBGetEndpointBurstSize(usbDeviceSpeed uint32, descriptor uintptr, companionDescriptor uintptr, sspCompanionDescriptor uintptr) uint32 {
	result, callErr := tryIOUSBGetEndpointBurstSize(usbDeviceSpeed, descriptor, companionDescriptor, sspCompanionDescriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetEndpointDirection func(descriptor uintptr) uint8
var _iOUSBGetEndpointDirectionErr error

func tryIOUSBGetEndpointDirection(descriptor uintptr) (uint8, error) {
	if _iOUSBGetEndpointDirection == nil {
		return 0, symbolCallError("IOUSBGetEndpointDirection", "10.15", _iOUSBGetEndpointDirectionErr)
	}
	return _iOUSBGetEndpointDirection(descriptor), nil
}

// IOUSBGetEndpointDirection obtains the direction of an endpoint from an endpoint descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointDirection(_:)
func IOUSBGetEndpointDirection(descriptor uintptr) uint8 {
	result, callErr := tryIOUSBGetEndpointDirection(descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetEndpointIntervalEncodedMicroframes func(usbDeviceSpeed uint32, descriptor uintptr) uint32
var _iOUSBGetEndpointIntervalEncodedMicroframesErr error

func tryIOUSBGetEndpointIntervalEncodedMicroframes(usbDeviceSpeed uint32, descriptor uintptr) (uint32, error) {
	if _iOUSBGetEndpointIntervalEncodedMicroframes == nil {
		return 0, symbolCallError("IOUSBGetEndpointIntervalEncodedMicroframes", "10.15", _iOUSBGetEndpointIntervalEncodedMicroframesErr)
	}
	return _iOUSBGetEndpointIntervalEncodedMicroframes(usbDeviceSpeed, descriptor), nil
}

// IOUSBGetEndpointIntervalEncodedMicroframes obtains the interval of an endpoint descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointIntervalEncodedMicroframes(_:_:)
func IOUSBGetEndpointIntervalEncodedMicroframes(usbDeviceSpeed uint32, descriptor uintptr) uint32 {
	result, callErr := tryIOUSBGetEndpointIntervalEncodedMicroframes(usbDeviceSpeed, descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetEndpointIntervalFrames func(usbDeviceSpeed uint32, descriptor uintptr) uint32
var _iOUSBGetEndpointIntervalFramesErr error

func tryIOUSBGetEndpointIntervalFrames(usbDeviceSpeed uint32, descriptor uintptr) (uint32, error) {
	if _iOUSBGetEndpointIntervalFrames == nil {
		return 0, symbolCallError("IOUSBGetEndpointIntervalFrames", "10.15", _iOUSBGetEndpointIntervalFramesErr)
	}
	return _iOUSBGetEndpointIntervalFrames(usbDeviceSpeed, descriptor), nil
}

// IOUSBGetEndpointIntervalFrames obtains the interval of an endpoint descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointIntervalFrames(_:_:)
func IOUSBGetEndpointIntervalFrames(usbDeviceSpeed uint32, descriptor uintptr) uint32 {
	result, callErr := tryIOUSBGetEndpointIntervalFrames(usbDeviceSpeed, descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetEndpointIntervalMicroframes func(usbDeviceSpeed uint32, descriptor uintptr) uint32
var _iOUSBGetEndpointIntervalMicroframesErr error

func tryIOUSBGetEndpointIntervalMicroframes(usbDeviceSpeed uint32, descriptor uintptr) (uint32, error) {
	if _iOUSBGetEndpointIntervalMicroframes == nil {
		return 0, symbolCallError("IOUSBGetEndpointIntervalMicroframes", "10.15", _iOUSBGetEndpointIntervalMicroframesErr)
	}
	return _iOUSBGetEndpointIntervalMicroframes(usbDeviceSpeed, descriptor), nil
}

// IOUSBGetEndpointIntervalMicroframes obtains the interval of an endpoint descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointIntervalMicroframes(_:_:)
func IOUSBGetEndpointIntervalMicroframes(usbDeviceSpeed uint32, descriptor uintptr) uint32 {
	result, callErr := tryIOUSBGetEndpointIntervalMicroframes(usbDeviceSpeed, descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetEndpointMaxPacketSize func(usbDeviceSpeed uint32, descriptor uintptr) uint16
var _iOUSBGetEndpointMaxPacketSizeErr error

func tryIOUSBGetEndpointMaxPacketSize(usbDeviceSpeed uint32, descriptor uintptr) (uint16, error) {
	if _iOUSBGetEndpointMaxPacketSize == nil {
		return 0, symbolCallError("IOUSBGetEndpointMaxPacketSize", "10.15", _iOUSBGetEndpointMaxPacketSizeErr)
	}
	return _iOUSBGetEndpointMaxPacketSize(usbDeviceSpeed, descriptor), nil
}

// IOUSBGetEndpointMaxPacketSize obtains the maximum packet size from an endpoint descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointMaxPacketSize(_:_:)
func IOUSBGetEndpointMaxPacketSize(usbDeviceSpeed uint32, descriptor uintptr) uint16 {
	result, callErr := tryIOUSBGetEndpointMaxPacketSize(usbDeviceSpeed, descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetEndpointMaxStreams func(usbDeviceSpeed uint32, descriptor uintptr, companionDescriptor uintptr) uint32
var _iOUSBGetEndpointMaxStreamsErr error

func tryIOUSBGetEndpointMaxStreams(usbDeviceSpeed uint32, descriptor uintptr, companionDescriptor uintptr) (uint32, error) {
	if _iOUSBGetEndpointMaxStreams == nil {
		return 0, symbolCallError("IOUSBGetEndpointMaxStreams", "10.15", _iOUSBGetEndpointMaxStreamsErr)
	}
	return _iOUSBGetEndpointMaxStreams(usbDeviceSpeed, descriptor, companionDescriptor), nil
}

// IOUSBGetEndpointMaxStreams obtains the number of supported streams.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointMaxStreams(_:_:_:)
func IOUSBGetEndpointMaxStreams(usbDeviceSpeed uint32, descriptor uintptr, companionDescriptor uintptr) uint32 {
	result, callErr := tryIOUSBGetEndpointMaxStreams(usbDeviceSpeed, descriptor, companionDescriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetEndpointMaxStreamsEncoded func(usbDeviceSpeed uint32, descriptor uintptr, companionDescriptor uintptr) uint32
var _iOUSBGetEndpointMaxStreamsEncodedErr error

func tryIOUSBGetEndpointMaxStreamsEncoded(usbDeviceSpeed uint32, descriptor uintptr, companionDescriptor uintptr) (uint32, error) {
	if _iOUSBGetEndpointMaxStreamsEncoded == nil {
		return 0, symbolCallError("IOUSBGetEndpointMaxStreamsEncoded", "10.15", _iOUSBGetEndpointMaxStreamsEncodedErr)
	}
	return _iOUSBGetEndpointMaxStreamsEncoded(usbDeviceSpeed, descriptor, companionDescriptor), nil
}

// IOUSBGetEndpointMaxStreamsEncoded obtains the number of streams that an endpoint supports.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointMaxStreamsEncoded(_:_:_:)
func IOUSBGetEndpointMaxStreamsEncoded(usbDeviceSpeed uint32, descriptor uintptr, companionDescriptor uintptr) uint32 {
	result, callErr := tryIOUSBGetEndpointMaxStreamsEncoded(usbDeviceSpeed, descriptor, companionDescriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetEndpointMult func(usbDeviceSpeed uint32, descriptor uintptr, companionDescriptor uintptr, sspCompanionDescriptor uintptr) uint8
var _iOUSBGetEndpointMultErr error

func tryIOUSBGetEndpointMult(usbDeviceSpeed uint32, descriptor uintptr, companionDescriptor uintptr, sspCompanionDescriptor uintptr) (uint8, error) {
	if _iOUSBGetEndpointMult == nil {
		return 0, symbolCallError("IOUSBGetEndpointMult", "10.15", _iOUSBGetEndpointMultErr)
	}
	return _iOUSBGetEndpointMult(usbDeviceSpeed, descriptor, companionDescriptor, sspCompanionDescriptor), nil
}

// IOUSBGetEndpointMult.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointMult(_:_:_:_:)
func IOUSBGetEndpointMult(usbDeviceSpeed uint32, descriptor uintptr, companionDescriptor uintptr, sspCompanionDescriptor uintptr) uint8 {
	result, callErr := tryIOUSBGetEndpointMult(usbDeviceSpeed, descriptor, companionDescriptor, sspCompanionDescriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetEndpointNumber func(descriptor uintptr) uint8
var _iOUSBGetEndpointNumberErr error

func tryIOUSBGetEndpointNumber(descriptor uintptr) (uint8, error) {
	if _iOUSBGetEndpointNumber == nil {
		return 0, symbolCallError("IOUSBGetEndpointNumber", "10.15", _iOUSBGetEndpointNumberErr)
	}
	return _iOUSBGetEndpointNumber(descriptor), nil
}

// IOUSBGetEndpointNumber obtains the number of an endpoint from an endpoint descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointNumber(_:)
func IOUSBGetEndpointNumber(descriptor uintptr) uint8 {
	result, callErr := tryIOUSBGetEndpointNumber(descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetEndpointSynchronizationType func(descriptor uintptr) uint8
var _iOUSBGetEndpointSynchronizationTypeErr error

func tryIOUSBGetEndpointSynchronizationType(descriptor uintptr) (uint8, error) {
	if _iOUSBGetEndpointSynchronizationType == nil {
		return 0, symbolCallError("IOUSBGetEndpointSynchronizationType", "10.15", _iOUSBGetEndpointSynchronizationTypeErr)
	}
	return _iOUSBGetEndpointSynchronizationType(descriptor), nil
}

// IOUSBGetEndpointSynchronizationType.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointSynchronizationType(_:)
func IOUSBGetEndpointSynchronizationType(descriptor uintptr) uint8 {
	result, callErr := tryIOUSBGetEndpointSynchronizationType(descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetEndpointType func(descriptor uintptr) uint8
var _iOUSBGetEndpointTypeErr error

func tryIOUSBGetEndpointType(descriptor uintptr) (uint8, error) {
	if _iOUSBGetEndpointType == nil {
		return 0, symbolCallError("IOUSBGetEndpointType", "10.15", _iOUSBGetEndpointTypeErr)
	}
	return _iOUSBGetEndpointType(descriptor), nil
}

// IOUSBGetEndpointType obtains the type of an endpoint from an endpoint descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointType(_:)
func IOUSBGetEndpointType(descriptor uintptr) uint8 {
	result, callErr := tryIOUSBGetEndpointType(descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetEndpointUsageType func(descriptor uintptr) uint8
var _iOUSBGetEndpointUsageTypeErr error

func tryIOUSBGetEndpointUsageType(descriptor uintptr) (uint8, error) {
	if _iOUSBGetEndpointUsageType == nil {
		return 0, symbolCallError("IOUSBGetEndpointUsageType", "10.15", _iOUSBGetEndpointUsageTypeErr)
	}
	return _iOUSBGetEndpointUsageType(descriptor), nil
}

// IOUSBGetEndpointUsageType.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointUsageType(_:)
func IOUSBGetEndpointUsageType(descriptor uintptr) uint8 {
	result, callErr := tryIOUSBGetEndpointUsageType(descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetNextAssociatedDescriptor func(configurationDescriptor uintptr, parentDescriptor uintptr, currentDescriptor uintptr) uintptr
var _iOUSBGetNextAssociatedDescriptorErr error

func tryIOUSBGetNextAssociatedDescriptor(configurationDescriptor uintptr, parentDescriptor uintptr, currentDescriptor uintptr) (uintptr, error) {
	if _iOUSBGetNextAssociatedDescriptor == nil {
		return 0, symbolCallError("IOUSBGetNextAssociatedDescriptor", "10.15", _iOUSBGetNextAssociatedDescriptorErr)
	}
	return _iOUSBGetNextAssociatedDescriptor(configurationDescriptor, parentDescriptor, currentDescriptor), nil
}

// IOUSBGetNextAssociatedDescriptor obtains the next associated descriptor in a configuration descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetNextAssociatedDescriptor(_:_:_:)
func IOUSBGetNextAssociatedDescriptor(configurationDescriptor uintptr, parentDescriptor uintptr, currentDescriptor uintptr) uintptr {
	result, callErr := tryIOUSBGetNextAssociatedDescriptor(configurationDescriptor, parentDescriptor, currentDescriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetNextAssociatedDescriptorWithType func(configurationDescriptor uintptr, parentDescriptor uintptr, currentDescriptor uintptr, type_ uint8) uintptr
var _iOUSBGetNextAssociatedDescriptorWithTypeErr error

func tryIOUSBGetNextAssociatedDescriptorWithType(configurationDescriptor uintptr, parentDescriptor uintptr, currentDescriptor uintptr, type_ uint8) (uintptr, error) {
	if _iOUSBGetNextAssociatedDescriptorWithType == nil {
		return 0, symbolCallError("IOUSBGetNextAssociatedDescriptorWithType", "10.15", _iOUSBGetNextAssociatedDescriptorWithTypeErr)
	}
	return _iOUSBGetNextAssociatedDescriptorWithType(configurationDescriptor, parentDescriptor, currentDescriptor, type_), nil
}

// IOUSBGetNextAssociatedDescriptorWithType obtains the next associated descriptor in a configuration descriptor and matches the type.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetNextAssociatedDescriptorWithType(_:_:_:_:)
func IOUSBGetNextAssociatedDescriptorWithType(configurationDescriptor uintptr, parentDescriptor uintptr, currentDescriptor uintptr, type_ uint8) uintptr {
	result, callErr := tryIOUSBGetNextAssociatedDescriptorWithType(configurationDescriptor, parentDescriptor, currentDescriptor, type_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetNextCapabilityDescriptor func(bosDescriptor uintptr, currentDescriptor uintptr) uintptr
var _iOUSBGetNextCapabilityDescriptorErr error

func tryIOUSBGetNextCapabilityDescriptor(bosDescriptor uintptr, currentDescriptor uintptr) (uintptr, error) {
	if _iOUSBGetNextCapabilityDescriptor == nil {
		return 0, symbolCallError("IOUSBGetNextCapabilityDescriptor", "10.15", _iOUSBGetNextCapabilityDescriptorErr)
	}
	return _iOUSBGetNextCapabilityDescriptor(bosDescriptor, currentDescriptor), nil
}

// IOUSBGetNextCapabilityDescriptor obtains the next device capability descriptor in a BOS descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetNextCapabilityDescriptor(_:_:)
func IOUSBGetNextCapabilityDescriptor(bosDescriptor uintptr, currentDescriptor uintptr) uintptr {
	result, callErr := tryIOUSBGetNextCapabilityDescriptor(bosDescriptor, currentDescriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetNextCapabilityDescriptorWithType func(bosDescriptor uintptr, currentDescriptor uintptr, type_ uint8) uintptr
var _iOUSBGetNextCapabilityDescriptorWithTypeErr error

func tryIOUSBGetNextCapabilityDescriptorWithType(bosDescriptor uintptr, currentDescriptor uintptr, type_ uint8) (uintptr, error) {
	if _iOUSBGetNextCapabilityDescriptorWithType == nil {
		return 0, symbolCallError("IOUSBGetNextCapabilityDescriptorWithType", "10.15", _iOUSBGetNextCapabilityDescriptorWithTypeErr)
	}
	return _iOUSBGetNextCapabilityDescriptorWithType(bosDescriptor, currentDescriptor, type_), nil
}

// IOUSBGetNextCapabilityDescriptorWithType obtains the next descriptor matching a specific type within a BOS descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetNextCapabilityDescriptorWithType(_:_:_:)
func IOUSBGetNextCapabilityDescriptorWithType(bosDescriptor uintptr, currentDescriptor uintptr, type_ uint8) uintptr {
	result, callErr := tryIOUSBGetNextCapabilityDescriptorWithType(bosDescriptor, currentDescriptor, type_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetNextDescriptor func(configurationDescriptor uintptr, currentDescriptor uintptr) uintptr
var _iOUSBGetNextDescriptorErr error

func tryIOUSBGetNextDescriptor(configurationDescriptor uintptr, currentDescriptor uintptr) (uintptr, error) {
	if _iOUSBGetNextDescriptor == nil {
		return 0, symbolCallError("IOUSBGetNextDescriptor", "10.15", _iOUSBGetNextDescriptorErr)
	}
	return _iOUSBGetNextDescriptor(configurationDescriptor, currentDescriptor), nil
}

// IOUSBGetNextDescriptor obtains the next descriptor in a configuration descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetNextDescriptor(_:_:)
func IOUSBGetNextDescriptor(configurationDescriptor uintptr, currentDescriptor uintptr) uintptr {
	result, callErr := tryIOUSBGetNextDescriptor(configurationDescriptor, currentDescriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetNextDescriptorWithType func(configurationDescriptor uintptr, currentDescriptor uintptr, type_ uint8) uintptr
var _iOUSBGetNextDescriptorWithTypeErr error

func tryIOUSBGetNextDescriptorWithType(configurationDescriptor uintptr, currentDescriptor uintptr, type_ uint8) (uintptr, error) {
	if _iOUSBGetNextDescriptorWithType == nil {
		return 0, symbolCallError("IOUSBGetNextDescriptorWithType", "10.15", _iOUSBGetNextDescriptorWithTypeErr)
	}
	return _iOUSBGetNextDescriptorWithType(configurationDescriptor, currentDescriptor, type_), nil
}

// IOUSBGetNextDescriptorWithType obtains the next descriptor in a configuration descriptor that matches the type.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetNextDescriptorWithType(_:_:_:)
func IOUSBGetNextDescriptorWithType(configurationDescriptor uintptr, currentDescriptor uintptr, type_ uint8) uintptr {
	result, callErr := tryIOUSBGetNextDescriptorWithType(configurationDescriptor, currentDescriptor, type_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetNextEndpointDescriptor func(configurationDescriptor uintptr, interfaceDescriptor uintptr, currentDescriptor uintptr) uintptr
var _iOUSBGetNextEndpointDescriptorErr error

func tryIOUSBGetNextEndpointDescriptor(configurationDescriptor uintptr, interfaceDescriptor uintptr, currentDescriptor uintptr) (uintptr, error) {
	if _iOUSBGetNextEndpointDescriptor == nil {
		return 0, symbolCallError("IOUSBGetNextEndpointDescriptor", "10.15", _iOUSBGetNextEndpointDescriptorErr)
	}
	return _iOUSBGetNextEndpointDescriptor(configurationDescriptor, interfaceDescriptor, currentDescriptor), nil
}

// IOUSBGetNextEndpointDescriptor obtains the next endpoint descriptor for an interface descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetNextEndpointDescriptor(_:_:_:)
func IOUSBGetNextEndpointDescriptor(configurationDescriptor uintptr, interfaceDescriptor uintptr, currentDescriptor uintptr) uintptr {
	result, callErr := tryIOUSBGetNextEndpointDescriptor(configurationDescriptor, interfaceDescriptor, currentDescriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetNextInterfaceAssociationDescriptor func(configurationDescriptor uintptr, currentDescriptor uintptr) uintptr
var _iOUSBGetNextInterfaceAssociationDescriptorErr error

func tryIOUSBGetNextInterfaceAssociationDescriptor(configurationDescriptor uintptr, currentDescriptor uintptr) (uintptr, error) {
	if _iOUSBGetNextInterfaceAssociationDescriptor == nil {
		return 0, symbolCallError("IOUSBGetNextInterfaceAssociationDescriptor", "10.15", _iOUSBGetNextInterfaceAssociationDescriptorErr)
	}
	return _iOUSBGetNextInterfaceAssociationDescriptor(configurationDescriptor, currentDescriptor), nil
}

// IOUSBGetNextInterfaceAssociationDescriptor obtains the next interface association descriptor in a configuration descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetNextInterfaceAssociationDescriptor(_:_:)
func IOUSBGetNextInterfaceAssociationDescriptor(configurationDescriptor uintptr, currentDescriptor uintptr) uintptr {
	result, callErr := tryIOUSBGetNextInterfaceAssociationDescriptor(configurationDescriptor, currentDescriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetNextInterfaceDescriptor func(configurationDescriptor uintptr, currentDescriptor uintptr) uintptr
var _iOUSBGetNextInterfaceDescriptorErr error

func tryIOUSBGetNextInterfaceDescriptor(configurationDescriptor uintptr, currentDescriptor uintptr) (uintptr, error) {
	if _iOUSBGetNextInterfaceDescriptor == nil {
		return 0, symbolCallError("IOUSBGetNextInterfaceDescriptor", "10.15", _iOUSBGetNextInterfaceDescriptorErr)
	}
	return _iOUSBGetNextInterfaceDescriptor(configurationDescriptor, currentDescriptor), nil
}

// IOUSBGetNextInterfaceDescriptor obtains the next interface descriptor in a configuration descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetNextInterfaceDescriptor(_:_:)
func IOUSBGetNextInterfaceDescriptor(configurationDescriptor uintptr, currentDescriptor uintptr) uintptr {
	result, callErr := tryIOUSBGetNextInterfaceDescriptor(configurationDescriptor, currentDescriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetPlatformCapabilityDescriptor func(bosDescriptor uintptr) uintptr
var _iOUSBGetPlatformCapabilityDescriptorErr error

func tryIOUSBGetPlatformCapabilityDescriptor(bosDescriptor uintptr) (uintptr, error) {
	if _iOUSBGetPlatformCapabilityDescriptor == nil {
		return 0, symbolCallError("IOUSBGetPlatformCapabilityDescriptor", "10.15", _iOUSBGetPlatformCapabilityDescriptorErr)
	}
	return _iOUSBGetPlatformCapabilityDescriptor(bosDescriptor), nil
}

// IOUSBGetPlatformCapabilityDescriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetPlatformCapabilityDescriptor(_:)
func IOUSBGetPlatformCapabilityDescriptor(bosDescriptor uintptr) uintptr {
	result, callErr := tryIOUSBGetPlatformCapabilityDescriptor(bosDescriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetPlatformCapabilityDescriptorWithUUID func(bosDescriptor uintptr, uuid *[16]byte) uintptr
var _iOUSBGetPlatformCapabilityDescriptorWithUUIDErr error

func tryIOUSBGetPlatformCapabilityDescriptorWithUUID(bosDescriptor uintptr, uuid [16]byte) (uintptr, error) {
	if _iOUSBGetPlatformCapabilityDescriptorWithUUID == nil {
		return 0, symbolCallError("IOUSBGetPlatformCapabilityDescriptorWithUUID", "10.15", _iOUSBGetPlatformCapabilityDescriptorWithUUIDErr)
	}
	return _iOUSBGetPlatformCapabilityDescriptorWithUUID(bosDescriptor, &uuid), nil
}

// IOUSBGetPlatformCapabilityDescriptorWithUUID.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetPlatformCapabilityDescriptorWithUUID(_:_:)
func IOUSBGetPlatformCapabilityDescriptorWithUUID(bosDescriptor uintptr, uuid [16]byte) uintptr {
	result, callErr := tryIOUSBGetPlatformCapabilityDescriptorWithUUID(bosDescriptor, uuid)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetSuperSpeedDeviceCapabilityDescriptor func(bosDescriptor uintptr) uintptr
var _iOUSBGetSuperSpeedDeviceCapabilityDescriptorErr error

func tryIOUSBGetSuperSpeedDeviceCapabilityDescriptor(bosDescriptor uintptr) (uintptr, error) {
	if _iOUSBGetSuperSpeedDeviceCapabilityDescriptor == nil {
		return 0, symbolCallError("IOUSBGetSuperSpeedDeviceCapabilityDescriptor", "10.15", _iOUSBGetSuperSpeedDeviceCapabilityDescriptorErr)
	}
	return _iOUSBGetSuperSpeedDeviceCapabilityDescriptor(bosDescriptor), nil
}

// IOUSBGetSuperSpeedDeviceCapabilityDescriptor obtains the first SuperSpeed capability descriptor in a BOS descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetSuperSpeedDeviceCapabilityDescriptor(_:)
func IOUSBGetSuperSpeedDeviceCapabilityDescriptor(bosDescriptor uintptr) uintptr {
	result, callErr := tryIOUSBGetSuperSpeedDeviceCapabilityDescriptor(bosDescriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetSuperSpeedPlusDeviceCapabilityDescriptor func(bosDescriptor uintptr) uintptr
var _iOUSBGetSuperSpeedPlusDeviceCapabilityDescriptorErr error

func tryIOUSBGetSuperSpeedPlusDeviceCapabilityDescriptor(bosDescriptor uintptr) (uintptr, error) {
	if _iOUSBGetSuperSpeedPlusDeviceCapabilityDescriptor == nil {
		return 0, symbolCallError("IOUSBGetSuperSpeedPlusDeviceCapabilityDescriptor", "10.15", _iOUSBGetSuperSpeedPlusDeviceCapabilityDescriptorErr)
	}
	return _iOUSBGetSuperSpeedPlusDeviceCapabilityDescriptor(bosDescriptor), nil
}

// IOUSBGetSuperSpeedPlusDeviceCapabilityDescriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetSuperSpeedPlusDeviceCapabilityDescriptor(_:)
func IOUSBGetSuperSpeedPlusDeviceCapabilityDescriptor(bosDescriptor uintptr) uintptr {
	result, callErr := tryIOUSBGetSuperSpeedPlusDeviceCapabilityDescriptor(bosDescriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBGetUSB20ExtensionDeviceCapabilityDescriptor func(bosDescriptor uintptr) uintptr
var _iOUSBGetUSB20ExtensionDeviceCapabilityDescriptorErr error

func tryIOUSBGetUSB20ExtensionDeviceCapabilityDescriptor(bosDescriptor uintptr) (uintptr, error) {
	if _iOUSBGetUSB20ExtensionDeviceCapabilityDescriptor == nil {
		return 0, symbolCallError("IOUSBGetUSB20ExtensionDeviceCapabilityDescriptor", "10.15", _iOUSBGetUSB20ExtensionDeviceCapabilityDescriptorErr)
	}
	return _iOUSBGetUSB20ExtensionDeviceCapabilityDescriptor(bosDescriptor), nil
}

// IOUSBGetUSB20ExtensionDeviceCapabilityDescriptor obtains the first USB 2.0 extension capability descriptor in a BOS descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetUSB20ExtensionDeviceCapabilityDescriptor(_:)
func IOUSBGetUSB20ExtensionDeviceCapabilityDescriptor(bosDescriptor uintptr) uintptr {
	result, callErr := tryIOUSBGetUSB20ExtensionDeviceCapabilityDescriptor(bosDescriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBHostCIControllerStateToString func(controllerState IOUSBHostCIControllerState) *byte
var _iOUSBHostCIControllerStateToStringErr error

func tryIOUSBHostCIControllerStateToString(controllerState IOUSBHostCIControllerState) (*byte, error) {
	if _iOUSBHostCIControllerStateToString == nil {
		return nil, symbolCallError("IOUSBHostCIControllerStateToString", "10.15", _iOUSBHostCIControllerStateToStringErr)
	}
	return _iOUSBHostCIControllerStateToString(controllerState), nil
}

// IOUSBHostCIControllerStateToString.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateToString(_:)
func IOUSBHostCIControllerStateToString(controllerState IOUSBHostCIControllerState) *byte {
	result, callErr := tryIOUSBHostCIControllerStateToString(controllerState)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBHostCIDeviceSpeedToString func(speed IOUSBHostCIDeviceSpeed) *byte
var _iOUSBHostCIDeviceSpeedToStringErr error

func tryIOUSBHostCIDeviceSpeedToString(speed IOUSBHostCIDeviceSpeed) (*byte, error) {
	if _iOUSBHostCIDeviceSpeedToString == nil {
		return nil, symbolCallError("IOUSBHostCIDeviceSpeedToString", "10.15", _iOUSBHostCIDeviceSpeedToStringErr)
	}
	return _iOUSBHostCIDeviceSpeedToString(speed), nil
}

// IOUSBHostCIDeviceSpeedToString.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceSpeedToString(_:)
func IOUSBHostCIDeviceSpeedToString(speed IOUSBHostCIDeviceSpeed) *byte {
	result, callErr := tryIOUSBHostCIDeviceSpeedToString(speed)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBHostCIDeviceStateToString func(deviceState IOUSBHostCIDeviceState) *byte
var _iOUSBHostCIDeviceStateToStringErr error

func tryIOUSBHostCIDeviceStateToString(deviceState IOUSBHostCIDeviceState) (*byte, error) {
	if _iOUSBHostCIDeviceStateToString == nil {
		return nil, symbolCallError("IOUSBHostCIDeviceStateToString", "10.15", _iOUSBHostCIDeviceStateToStringErr)
	}
	return _iOUSBHostCIDeviceStateToString(deviceState), nil
}

// IOUSBHostCIDeviceStateToString.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateToString(_:)
func IOUSBHostCIDeviceStateToString(deviceState IOUSBHostCIDeviceState) *byte {
	result, callErr := tryIOUSBHostCIDeviceStateToString(deviceState)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBHostCIEndpointStateToString func(endpointState IOUSBHostCIEndpointState) *byte
var _iOUSBHostCIEndpointStateToStringErr error

func tryIOUSBHostCIEndpointStateToString(endpointState IOUSBHostCIEndpointState) (*byte, error) {
	if _iOUSBHostCIEndpointStateToString == nil {
		return nil, symbolCallError("IOUSBHostCIEndpointStateToString", "10.15", _iOUSBHostCIEndpointStateToStringErr)
	}
	return _iOUSBHostCIEndpointStateToString(endpointState), nil
}

// IOUSBHostCIEndpointStateToString.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateToString(_:)
func IOUSBHostCIEndpointStateToString(endpointState IOUSBHostCIEndpointState) *byte {
	result, callErr := tryIOUSBHostCIEndpointStateToString(endpointState)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBHostCIExceptionTypeToString func(exceptionType IOUSBHostCIExceptionType) *byte
var _iOUSBHostCIExceptionTypeToStringErr error

func tryIOUSBHostCIExceptionTypeToString(exceptionType IOUSBHostCIExceptionType) (*byte, error) {
	if _iOUSBHostCIExceptionTypeToString == nil {
		return nil, symbolCallError("IOUSBHostCIExceptionTypeToString", "10.15", _iOUSBHostCIExceptionTypeToStringErr)
	}
	return _iOUSBHostCIExceptionTypeToString(exceptionType), nil
}

// IOUSBHostCIExceptionTypeToString.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIExceptionTypeToString(_:)
func IOUSBHostCIExceptionTypeToString(exceptionType IOUSBHostCIExceptionType) *byte {
	result, callErr := tryIOUSBHostCIExceptionTypeToString(exceptionType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBHostCILinkStateEnabled func(linkState IOUSBHostCILinkState) bool
var _iOUSBHostCILinkStateEnabledErr error

func tryIOUSBHostCILinkStateEnabled(linkState IOUSBHostCILinkState) (bool, error) {
	if _iOUSBHostCILinkStateEnabled == nil {
		return false, symbolCallError("IOUSBHostCILinkStateEnabled", "10.15", _iOUSBHostCILinkStateEnabledErr)
	}
	return _iOUSBHostCILinkStateEnabled(linkState), nil
}

// IOUSBHostCILinkStateEnabled.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCILinkStateEnabled(_:)
func IOUSBHostCILinkStateEnabled(linkState IOUSBHostCILinkState) bool {
	result, callErr := tryIOUSBHostCILinkStateEnabled(linkState)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBHostCILinkStateToString func(linkState IOUSBHostCILinkState) *byte
var _iOUSBHostCILinkStateToStringErr error

func tryIOUSBHostCILinkStateToString(linkState IOUSBHostCILinkState) (*byte, error) {
	if _iOUSBHostCILinkStateToString == nil {
		return nil, symbolCallError("IOUSBHostCILinkStateToString", "10.15", _iOUSBHostCILinkStateToStringErr)
	}
	return _iOUSBHostCILinkStateToString(linkState), nil
}

// IOUSBHostCILinkStateToString.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCILinkStateToString(_:)
func IOUSBHostCILinkStateToString(linkState IOUSBHostCILinkState) *byte {
	result, callErr := tryIOUSBHostCILinkStateToString(linkState)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBHostCIMessageStatusFromIOReturn func(status int32) IOUSBHostCIMessageStatus
var _iOUSBHostCIMessageStatusFromIOReturnErr error

func tryIOUSBHostCIMessageStatusFromIOReturn(status int32) (IOUSBHostCIMessageStatus, error) {
	if _iOUSBHostCIMessageStatusFromIOReturn == nil {
		return *new(IOUSBHostCIMessageStatus), symbolCallError("IOUSBHostCIMessageStatusFromIOReturn", "10.15", _iOUSBHostCIMessageStatusFromIOReturnErr)
	}
	return _iOUSBHostCIMessageStatusFromIOReturn(status), nil
}

// IOUSBHostCIMessageStatusFromIOReturn.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIMessageStatusFromIOReturn(_:)
func IOUSBHostCIMessageStatusFromIOReturn(status int32) IOUSBHostCIMessageStatus {
	result, callErr := tryIOUSBHostCIMessageStatusFromIOReturn(status)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBHostCIMessageStatusToIOReturn func(status IOUSBHostCIMessageStatus) int32
var _iOUSBHostCIMessageStatusToIOReturnErr error

func tryIOUSBHostCIMessageStatusToIOReturn(status IOUSBHostCIMessageStatus) (int32, error) {
	if _iOUSBHostCIMessageStatusToIOReturn == nil {
		return 0, symbolCallError("IOUSBHostCIMessageStatusToIOReturn", "10.15", _iOUSBHostCIMessageStatusToIOReturnErr)
	}
	return _iOUSBHostCIMessageStatusToIOReturn(status), nil
}

// IOUSBHostCIMessageStatusToIOReturn.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIMessageStatusToIOReturn(_:)
func IOUSBHostCIMessageStatusToIOReturn(status IOUSBHostCIMessageStatus) int32 {
	result, callErr := tryIOUSBHostCIMessageStatusToIOReturn(status)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBHostCIMessageStatusToString func(status IOUSBHostCIMessageStatus) *byte
var _iOUSBHostCIMessageStatusToStringErr error

func tryIOUSBHostCIMessageStatusToString(status IOUSBHostCIMessageStatus) (*byte, error) {
	if _iOUSBHostCIMessageStatusToString == nil {
		return nil, symbolCallError("IOUSBHostCIMessageStatusToString", "10.15", _iOUSBHostCIMessageStatusToStringErr)
	}
	return _iOUSBHostCIMessageStatusToString(status), nil
}

// IOUSBHostCIMessageStatusToString.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIMessageStatusToString(_:)
func IOUSBHostCIMessageStatusToString(status IOUSBHostCIMessageStatus) *byte {
	result, callErr := tryIOUSBHostCIMessageStatusToString(status)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBHostCIMessageTypeToString func(type_ IOUSBHostCIMessageType) *byte
var _iOUSBHostCIMessageTypeToStringErr error

func tryIOUSBHostCIMessageTypeToString(type_ IOUSBHostCIMessageType) (*byte, error) {
	if _iOUSBHostCIMessageTypeToString == nil {
		return nil, symbolCallError("IOUSBHostCIMessageTypeToString", "10.15", _iOUSBHostCIMessageTypeToStringErr)
	}
	return _iOUSBHostCIMessageTypeToString(type_), nil
}

// IOUSBHostCIMessageTypeToString.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIMessageTypeToString(_:)
func IOUSBHostCIMessageTypeToString(type_ IOUSBHostCIMessageType) *byte {
	result, callErr := tryIOUSBHostCIMessageTypeToString(type_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOUSBHostCIPortStateToString func(portState IOUSBHostCIPortState) *byte
var _iOUSBHostCIPortStateToStringErr error

func tryIOUSBHostCIPortStateToString(portState IOUSBHostCIPortState) (*byte, error) {
	if _iOUSBHostCIPortStateToString == nil {
		return nil, symbolCallError("IOUSBHostCIPortStateToString", "10.15", _iOUSBHostCIPortStateToStringErr)
	}
	return _iOUSBHostCIPortStateToString(portState), nil
}

// IOUSBHostCIPortStateToString.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateToString(_:)
func IOUSBHostCIPortStateToString(portState IOUSBHostCIPortState) *byte {
	result, callErr := tryIOUSBHostCIPortStateToString(portState)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_iOUSBGetBillboardDescriptor, &_iOUSBGetBillboardDescriptorErr, frameworkHandle, "IOUSBGetBillboardDescriptor", "10.15")
	registerFunc(&_iOUSBGetConfigurationMaxPowerMilliAmps, &_iOUSBGetConfigurationMaxPowerMilliAmpsErr, frameworkHandle, "IOUSBGetConfigurationMaxPowerMilliAmps", "10.15")
	registerFunc(&_iOUSBGetContainerIDDescriptor, &_iOUSBGetContainerIDDescriptorErr, frameworkHandle, "IOUSBGetContainerIDDescriptor", "10.15")
	registerFunc(&_iOUSBGetEndpointAddress, &_iOUSBGetEndpointAddressErr, frameworkHandle, "IOUSBGetEndpointAddress", "10.15")
	registerFunc(&_iOUSBGetEndpointBurstSize, &_iOUSBGetEndpointBurstSizeErr, frameworkHandle, "IOUSBGetEndpointBurstSize", "10.15")
	registerFunc(&_iOUSBGetEndpointDirection, &_iOUSBGetEndpointDirectionErr, frameworkHandle, "IOUSBGetEndpointDirection", "10.15")
	registerFunc(&_iOUSBGetEndpointIntervalEncodedMicroframes, &_iOUSBGetEndpointIntervalEncodedMicroframesErr, frameworkHandle, "IOUSBGetEndpointIntervalEncodedMicroframes", "10.15")
	registerFunc(&_iOUSBGetEndpointIntervalFrames, &_iOUSBGetEndpointIntervalFramesErr, frameworkHandle, "IOUSBGetEndpointIntervalFrames", "10.15")
	registerFunc(&_iOUSBGetEndpointIntervalMicroframes, &_iOUSBGetEndpointIntervalMicroframesErr, frameworkHandle, "IOUSBGetEndpointIntervalMicroframes", "10.15")
	registerFunc(&_iOUSBGetEndpointMaxPacketSize, &_iOUSBGetEndpointMaxPacketSizeErr, frameworkHandle, "IOUSBGetEndpointMaxPacketSize", "10.15")
	registerFunc(&_iOUSBGetEndpointMaxStreams, &_iOUSBGetEndpointMaxStreamsErr, frameworkHandle, "IOUSBGetEndpointMaxStreams", "10.15")
	registerFunc(&_iOUSBGetEndpointMaxStreamsEncoded, &_iOUSBGetEndpointMaxStreamsEncodedErr, frameworkHandle, "IOUSBGetEndpointMaxStreamsEncoded", "10.15")
	registerFunc(&_iOUSBGetEndpointMult, &_iOUSBGetEndpointMultErr, frameworkHandle, "IOUSBGetEndpointMult", "10.15")
	registerFunc(&_iOUSBGetEndpointNumber, &_iOUSBGetEndpointNumberErr, frameworkHandle, "IOUSBGetEndpointNumber", "10.15")
	registerFunc(&_iOUSBGetEndpointSynchronizationType, &_iOUSBGetEndpointSynchronizationTypeErr, frameworkHandle, "IOUSBGetEndpointSynchronizationType", "10.15")
	registerFunc(&_iOUSBGetEndpointType, &_iOUSBGetEndpointTypeErr, frameworkHandle, "IOUSBGetEndpointType", "10.15")
	registerFunc(&_iOUSBGetEndpointUsageType, &_iOUSBGetEndpointUsageTypeErr, frameworkHandle, "IOUSBGetEndpointUsageType", "10.15")
	registerFunc(&_iOUSBGetNextAssociatedDescriptor, &_iOUSBGetNextAssociatedDescriptorErr, frameworkHandle, "IOUSBGetNextAssociatedDescriptor", "10.15")
	registerFunc(&_iOUSBGetNextAssociatedDescriptorWithType, &_iOUSBGetNextAssociatedDescriptorWithTypeErr, frameworkHandle, "IOUSBGetNextAssociatedDescriptorWithType", "10.15")
	registerFunc(&_iOUSBGetNextCapabilityDescriptor, &_iOUSBGetNextCapabilityDescriptorErr, frameworkHandle, "IOUSBGetNextCapabilityDescriptor", "10.15")
	registerFunc(&_iOUSBGetNextCapabilityDescriptorWithType, &_iOUSBGetNextCapabilityDescriptorWithTypeErr, frameworkHandle, "IOUSBGetNextCapabilityDescriptorWithType", "10.15")
	registerFunc(&_iOUSBGetNextDescriptor, &_iOUSBGetNextDescriptorErr, frameworkHandle, "IOUSBGetNextDescriptor", "10.15")
	registerFunc(&_iOUSBGetNextDescriptorWithType, &_iOUSBGetNextDescriptorWithTypeErr, frameworkHandle, "IOUSBGetNextDescriptorWithType", "10.15")
	registerFunc(&_iOUSBGetNextEndpointDescriptor, &_iOUSBGetNextEndpointDescriptorErr, frameworkHandle, "IOUSBGetNextEndpointDescriptor", "10.15")
	registerFunc(&_iOUSBGetNextInterfaceAssociationDescriptor, &_iOUSBGetNextInterfaceAssociationDescriptorErr, frameworkHandle, "IOUSBGetNextInterfaceAssociationDescriptor", "10.15")
	registerFunc(&_iOUSBGetNextInterfaceDescriptor, &_iOUSBGetNextInterfaceDescriptorErr, frameworkHandle, "IOUSBGetNextInterfaceDescriptor", "10.15")
	registerFunc(&_iOUSBGetPlatformCapabilityDescriptor, &_iOUSBGetPlatformCapabilityDescriptorErr, frameworkHandle, "IOUSBGetPlatformCapabilityDescriptor", "10.15")
	registerFunc(&_iOUSBGetPlatformCapabilityDescriptorWithUUID, &_iOUSBGetPlatformCapabilityDescriptorWithUUIDErr, frameworkHandle, "IOUSBGetPlatformCapabilityDescriptorWithUUID", "10.15")
	registerFunc(&_iOUSBGetSuperSpeedDeviceCapabilityDescriptor, &_iOUSBGetSuperSpeedDeviceCapabilityDescriptorErr, frameworkHandle, "IOUSBGetSuperSpeedDeviceCapabilityDescriptor", "10.15")
	registerFunc(&_iOUSBGetSuperSpeedPlusDeviceCapabilityDescriptor, &_iOUSBGetSuperSpeedPlusDeviceCapabilityDescriptorErr, frameworkHandle, "IOUSBGetSuperSpeedPlusDeviceCapabilityDescriptor", "10.15")
	registerFunc(&_iOUSBGetUSB20ExtensionDeviceCapabilityDescriptor, &_iOUSBGetUSB20ExtensionDeviceCapabilityDescriptorErr, frameworkHandle, "IOUSBGetUSB20ExtensionDeviceCapabilityDescriptor", "10.15")
	registerFunc(&_iOUSBHostCIControllerStateToString, &_iOUSBHostCIControllerStateToStringErr, frameworkHandle, "IOUSBHostCIControllerStateToString", "10.15")
	registerFunc(&_iOUSBHostCIDeviceSpeedToString, &_iOUSBHostCIDeviceSpeedToStringErr, frameworkHandle, "IOUSBHostCIDeviceSpeedToString", "10.15")
	registerFunc(&_iOUSBHostCIDeviceStateToString, &_iOUSBHostCIDeviceStateToStringErr, frameworkHandle, "IOUSBHostCIDeviceStateToString", "10.15")
	registerFunc(&_iOUSBHostCIEndpointStateToString, &_iOUSBHostCIEndpointStateToStringErr, frameworkHandle, "IOUSBHostCIEndpointStateToString", "10.15")
	registerFunc(&_iOUSBHostCIExceptionTypeToString, &_iOUSBHostCIExceptionTypeToStringErr, frameworkHandle, "IOUSBHostCIExceptionTypeToString", "10.15")
	registerFunc(&_iOUSBHostCILinkStateEnabled, &_iOUSBHostCILinkStateEnabledErr, frameworkHandle, "IOUSBHostCILinkStateEnabled", "10.15")
	registerFunc(&_iOUSBHostCILinkStateToString, &_iOUSBHostCILinkStateToStringErr, frameworkHandle, "IOUSBHostCILinkStateToString", "10.15")
	registerFunc(&_iOUSBHostCIMessageStatusFromIOReturn, &_iOUSBHostCIMessageStatusFromIOReturnErr, frameworkHandle, "IOUSBHostCIMessageStatusFromIOReturn", "10.15")
	registerFunc(&_iOUSBHostCIMessageStatusToIOReturn, &_iOUSBHostCIMessageStatusToIOReturnErr, frameworkHandle, "IOUSBHostCIMessageStatusToIOReturn", "10.15")
	registerFunc(&_iOUSBHostCIMessageStatusToString, &_iOUSBHostCIMessageStatusToStringErr, frameworkHandle, "IOUSBHostCIMessageStatusToString", "10.15")
	registerFunc(&_iOUSBHostCIMessageTypeToString, &_iOUSBHostCIMessageTypeToStringErr, frameworkHandle, "IOUSBHostCIMessageTypeToString", "10.15")
	registerFunc(&_iOUSBHostCIPortStateToString, &_iOUSBHostCIPortStateToStringErr, frameworkHandle, "IOUSBHostCIPortStateToString", "10.15")
}

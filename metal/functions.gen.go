// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
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
		return fmt.Sprintf("Metal: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("Metal: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("Metal: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("Metal: register symbol %s: %v", name, r)
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

var _mTLCopyAllDevices func() []unsafe.Pointer
var _mTLCopyAllDevicesErr error

func tryMTLCopyAllDevices() ([]MTLDeviceObject, error) {
	if _mTLCopyAllDevices == nil {
		return nil, symbolCallError("MTLCopyAllDevices", "10.11", _mTLCopyAllDevicesErr)
	}
	ptrs := _mTLCopyAllDevices()
	result := make([]MTLDeviceObject, len(ptrs))
	for i, p := range ptrs {
		result[i] = MTLDeviceObjectFromID(objc.IDFrom(p))
	}
	return result, nil
}

// MTLCopyAllDevices returns an array of all the Metal device instances in the system.
//
// See: https://developer.apple.com/documentation/Metal/MTLCopyAllDevices()
func MTLCopyAllDevices() []MTLDeviceObject {
	result, callErr := tryMTLCopyAllDevices()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mTLCopyAllDevicesWithObserver func(observer objectivec.Object, handler unsafe.Pointer) []unsafe.Pointer
var _mTLCopyAllDevicesWithObserverErr error

func tryMTLCopyAllDevicesWithObserver(observer objectivec.Object, handler MTLDeviceNotificationHandler) ([]MTLDeviceObject, error) {
	if _mTLCopyAllDevicesWithObserver == nil {
		return nil, symbolCallError("MTLCopyAllDevicesWithObserver", "10.13", _mTLCopyAllDevicesWithObserverErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 MTLDevice, arg1 objc.ID) { handler(arg0, objc.IDToString(arg1)) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	ptrs := _mTLCopyAllDevicesWithObserver(observer, _block0)
	result := make([]MTLDeviceObject, len(ptrs))
	for i, p := range ptrs {
		result[i] = MTLDeviceObjectFromID(objc.IDFrom(p))
	}
	return result, nil
}

// MTLCopyAllDevicesWithObserver returns an array of all the Metal GPU devices in the system and registers a notification handler that Metal calls when the device list changes.
//
// See: https://developer.apple.com/documentation/Metal/MTLCopyAllDevicesWithObserver
func MTLCopyAllDevicesWithObserver(observer objectivec.Object, handler MTLDeviceNotificationHandler) []MTLDeviceObject {
	result, callErr := tryMTLCopyAllDevicesWithObserver(observer, handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mTLCreateSystemDefaultDevice func() unsafe.Pointer
var _mTLCreateSystemDefaultDeviceErr error

func tryMTLCreateSystemDefaultDevice() (MTLDeviceObject, error) {
	if _mTLCreateSystemDefaultDevice == nil {
		return MTLDeviceObject{}, symbolCallError("MTLCreateSystemDefaultDevice", "10.11", _mTLCreateSystemDefaultDeviceErr)
	}
	rv := _mTLCreateSystemDefaultDevice()
	return MTLDeviceObjectFromID(objc.IDFrom(rv)), nil
}

// MTLCreateSystemDefaultDevice returns the device instance Metal selects as the default.
//
// See: https://developer.apple.com/documentation/Metal/MTLCreateSystemDefaultDevice()
func MTLCreateSystemDefaultDevice() MTLDeviceObject {
	result, callErr := tryMTLCreateSystemDefaultDevice()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mTLIOCompressionContextAppendData func(context MTLIOCompressionContext, data unsafe.Pointer, size uintptr)
var _mTLIOCompressionContextAppendDataErr error

func tryMTLIOCompressionContextAppendData(context MTLIOCompressionContext, data unsafe.Pointer, size uintptr) error {
	if _mTLIOCompressionContextAppendData == nil {
		return symbolCallError("MTLIOCompressionContextAppendData", "13.0", _mTLIOCompressionContextAppendDataErr)
	}
	_mTLIOCompressionContextAppendData(context, data, size)
	return nil
}

// MTLIOCompressionContextAppendData adds data to a compression context.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCompressionContextAppendData(_:_:_:)
func MTLIOCompressionContextAppendData(context MTLIOCompressionContext, data unsafe.Pointer, size uintptr) {
	if callErr := tryMTLIOCompressionContextAppendData(context, data, size); callErr != nil {
		panic(callErr)
	}
}

var _mTLIOCompressionContextDefaultChunkSize func() uintptr
var _mTLIOCompressionContextDefaultChunkSizeErr error

func tryMTLIOCompressionContextDefaultChunkSize() (uintptr, error) {
	if _mTLIOCompressionContextDefaultChunkSize == nil {
		return 0, symbolCallError("MTLIOCompressionContextDefaultChunkSize", "", _mTLIOCompressionContextDefaultChunkSizeErr)
	}
	return _mTLIOCompressionContextDefaultChunkSize(), nil
}

// MTLIOCompressionContextDefaultChunkSize returns a compression chunk size you can use as a default for creating a compression context.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCompressionContextDefaultChunkSize()
func MTLIOCompressionContextDefaultChunkSize() uintptr {
	result, callErr := tryMTLIOCompressionContextDefaultChunkSize()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mTLIOCreateCompressionContext func(path string, type_ MTLIOCompressionMethod, chunkSize uintptr) MTLIOCompressionContext
var _mTLIOCreateCompressionContextErr error

func tryMTLIOCreateCompressionContext(path string, type_ MTLIOCompressionMethod, chunkSize uintptr) (MTLIOCompressionContext, error) {
	if _mTLIOCreateCompressionContext == nil {
		return *new(MTLIOCompressionContext), symbolCallError("MTLIOCreateCompressionContext", "13.0", _mTLIOCreateCompressionContextErr)
	}
	return _mTLIOCreateCompressionContext(path, type_, chunkSize), nil
}

// MTLIOCreateCompressionContext creates a compression context that you use to compress data into a single file.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCreateCompressionContext
func MTLIOCreateCompressionContext(path string, type_ MTLIOCompressionMethod, chunkSize uintptr) MTLIOCompressionContext {
	result, callErr := tryMTLIOCreateCompressionContext(path, type_, chunkSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mTLIOFlushAndDestroyCompressionContext func(context MTLIOCompressionContext) MTLIOCompressionStatus
var _mTLIOFlushAndDestroyCompressionContextErr error

func tryMTLIOFlushAndDestroyCompressionContext(context MTLIOCompressionContext) (MTLIOCompressionStatus, error) {
	if _mTLIOFlushAndDestroyCompressionContext == nil {
		return *new(MTLIOCompressionStatus), symbolCallError("MTLIOFlushAndDestroyCompressionContext", "13.0", _mTLIOFlushAndDestroyCompressionContextErr)
	}
	return _mTLIOFlushAndDestroyCompressionContext(context), nil
}

// MTLIOFlushAndDestroyCompressionContext finishes compressing and saves the file that a compression context represents.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOFlushAndDestroyCompressionContext(_:)
func MTLIOFlushAndDestroyCompressionContext(context MTLIOCompressionContext) MTLIOCompressionStatus {
	result, callErr := tryMTLIOFlushAndDestroyCompressionContext(context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mTLRemoveDeviceObserver func(observer objectivec.Object)
var _mTLRemoveDeviceObserverErr error

func tryMTLRemoveDeviceObserver(observer objectivec.Object) error {
	if _mTLRemoveDeviceObserver == nil {
		return symbolCallError("MTLRemoveDeviceObserver", "10.13", _mTLRemoveDeviceObserverErr)
	}
	_mTLRemoveDeviceObserver(observer)
	return nil
}

// MTLRemoveDeviceObserver removes a registered observer of device notifications.
//
// See: https://developer.apple.com/documentation/Metal/MTLRemoveDeviceObserver(_:)
func MTLRemoveDeviceObserver(observer objectivec.Object) {
	if callErr := tryMTLRemoveDeviceObserver(observer); callErr != nil {
		panic(callErr)
	}
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_mTLCopyAllDevices, &_mTLCopyAllDevicesErr, frameworkHandle, "MTLCopyAllDevices", "10.11")
	registerFunc(&_mTLCopyAllDevicesWithObserver, &_mTLCopyAllDevicesWithObserverErr, frameworkHandle, "MTLCopyAllDevicesWithObserver", "10.13")
	registerFunc(&_mTLCreateSystemDefaultDevice, &_mTLCreateSystemDefaultDeviceErr, frameworkHandle, "MTLCreateSystemDefaultDevice", "10.11")
	registerFunc(&_mTLIOCompressionContextAppendData, &_mTLIOCompressionContextAppendDataErr, frameworkHandle, "MTLIOCompressionContextAppendData", "13.0")
	registerFunc(&_mTLIOCompressionContextDefaultChunkSize, &_mTLIOCompressionContextDefaultChunkSizeErr, frameworkHandle, "MTLIOCompressionContextDefaultChunkSize", "")
	registerFunc(&_mTLIOCreateCompressionContext, &_mTLIOCreateCompressionContextErr, frameworkHandle, "MTLIOCreateCompressionContext", "13.0")
	registerFunc(&_mTLIOFlushAndDestroyCompressionContext, &_mTLIOFlushAndDestroyCompressionContextErr, frameworkHandle, "MTLIOFlushAndDestroyCompressionContext", "13.0")
	registerFunc(&_mTLRemoveDeviceObserver, &_mTLRemoveDeviceObserverErr, frameworkHandle, "MTLRemoveDeviceObserver", "10.13")
}

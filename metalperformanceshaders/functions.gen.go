// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/metal"
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
		return fmt.Sprintf("MetalPerformanceShaders: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("MetalPerformanceShaders: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("MetalPerformanceShaders: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("MetalPerformanceShaders: register symbol %s: %v", name, r)
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

var _mPSGetImageType func(image *MPSImage) MPSImageType
var _mPSGetImageTypeErr error

func tryMPSGetImageType(image *MPSImage) (MPSImageType, error) {
	if _mPSGetImageType == nil {
		return *new(MPSImageType), symbolCallError("MPSGetImageType", "", _mPSGetImageTypeErr)
	}
	return _mPSGetImageType(image), nil
}

// MPSGetImageType.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSGetImageType(_:)
func MPSGetImageType(image *MPSImage) MPSImageType {
	result, callErr := tryMPSGetImageType(image)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mPSGetPreferredDevice func(options MPSDeviceOptions) unsafe.Pointer
var _mPSGetPreferredDeviceErr error

func tryMPSGetPreferredDevice(options MPSDeviceOptions) (metal.MTLDeviceObject, error) {
	if _mPSGetPreferredDevice == nil {
		return metal.MTLDeviceObject{}, symbolCallError("MPSGetPreferredDevice", "10.14.4", _mPSGetPreferredDeviceErr)
	}
	rv := _mPSGetPreferredDevice(options)
	return metal.MTLDeviceObjectFromID(objc.IDFrom(rv)), nil
}

// MPSGetPreferredDevice.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSGetPreferredDevice(_:)
func MPSGetPreferredDevice(options MPSDeviceOptions) metal.MTLDeviceObject {
	result, callErr := tryMPSGetPreferredDevice(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mPSHintTemporaryMemoryHighWaterMark func(cmdBuf metal.MTLCommandBufferObject, bytes uint)
var _mPSHintTemporaryMemoryHighWaterMarkErr error

func tryMPSHintTemporaryMemoryHighWaterMark(cmdBuf metal.MTLCommandBufferObject, bytes uint) error {
	if _mPSHintTemporaryMemoryHighWaterMark == nil {
		return symbolCallError("MPSHintTemporaryMemoryHighWaterMark", "", _mPSHintTemporaryMemoryHighWaterMarkErr)
	}
	_mPSHintTemporaryMemoryHighWaterMark(cmdBuf, bytes)
	return nil
}

// MPSHintTemporaryMemoryHighWaterMark triggers Metal Performance Shaders to prefetch a Metal heap of the indicated size into its internal cache.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSHintTemporaryMemoryHighWaterMark(_:_:)
func MPSHintTemporaryMemoryHighWaterMark(cmdBuf metal.MTLCommandBufferObject, bytes uint) {
	if callErr := tryMPSHintTemporaryMemoryHighWaterMark(cmdBuf, bytes); callErr != nil {
		panic(callErr)
	}
}

var _mPSImageBatchIncrementReadCount func(batch *MPSImageBatch, amount int) uint
var _mPSImageBatchIncrementReadCountErr error

func tryMPSImageBatchIncrementReadCount(batch *MPSImageBatch, amount int) (uint, error) {
	if _mPSImageBatchIncrementReadCount == nil {
		return 0, symbolCallError("MPSImageBatchIncrementReadCount", "10.13.4", _mPSImageBatchIncrementReadCountErr)
	}
	return _mPSImageBatchIncrementReadCount(batch, amount), nil
}

// MPSImageBatchIncrementReadCount increments or decrements the read count of an image batch by a specified amount.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageBatchIncrementReadCount(_:_:)
func MPSImageBatchIncrementReadCount(batch *MPSImageBatch, amount int) uint {
	result, callErr := tryMPSImageBatchIncrementReadCount(batch, amount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mPSImageBatchIterate func(batch *MPSImageBatch, iteratorBlock int) int
var _mPSImageBatchIterateErr error

func tryMPSImageBatchIterate(batch *MPSImageBatch, iteratorBlock int) (int, error) {
	if _mPSImageBatchIterate == nil {
		return 0, symbolCallError("MPSImageBatchIterate", "10.15", _mPSImageBatchIterateErr)
	}
	return _mPSImageBatchIterate(batch, iteratorBlock), nil
}

// MPSImageBatchIterate executes a callback block once for each unique image in a batch.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageBatchIterate(_:_:)
func MPSImageBatchIterate(batch *MPSImageBatch, iteratorBlock int) int {
	result, callErr := tryMPSImageBatchIterate(batch, iteratorBlock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mPSImageBatchResourceSize func(batch *MPSImageBatch) uint
var _mPSImageBatchResourceSizeErr error

func tryMPSImageBatchResourceSize(batch *MPSImageBatch) (uint, error) {
	if _mPSImageBatchResourceSize == nil {
		return 0, symbolCallError("MPSImageBatchResourceSize", "10.14", _mPSImageBatchResourceSizeErr)
	}
	return _mPSImageBatchResourceSize(batch), nil
}

// MPSImageBatchResourceSize returns the number of bytes used to allocate the specified image batch.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageBatchResourceSize(_:)
func MPSImageBatchResourceSize(batch *MPSImageBatch) uint {
	result, callErr := tryMPSImageBatchResourceSize(batch)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mPSImageBatchSynchronize func(batch *MPSImageBatch, cmdBuf metal.MTLCommandBufferObject)
var _mPSImageBatchSynchronizeErr error

func tryMPSImageBatchSynchronize(batch *MPSImageBatch, cmdBuf metal.MTLCommandBufferObject) error {
	if _mPSImageBatchSynchronize == nil {
		return symbolCallError("MPSImageBatchSynchronize", "10.13.4", _mPSImageBatchSynchronizeErr)
	}
	_mPSImageBatchSynchronize(batch, cmdBuf)
	return nil
}

// MPSImageBatchSynchronize removes any copy of the specified image batch from the device’s caches, and, if needed, invalidates any CPU caches.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageBatchSynchronize(_:_:)
func MPSImageBatchSynchronize(batch *MPSImageBatch, cmdBuf metal.MTLCommandBufferObject) {
	if callErr := tryMPSImageBatchSynchronize(batch, cmdBuf); callErr != nil {
		panic(callErr)
	}
}

var _mPSSetHeapCacheDuration func(cmdBuf metal.MTLCommandBufferObject, seconds float64)
var _mPSSetHeapCacheDurationErr error

func tryMPSSetHeapCacheDuration(cmdBuf metal.MTLCommandBufferObject, seconds float64) error {
	if _mPSSetHeapCacheDuration == nil {
		return symbolCallError("MPSSetHeapCacheDuration", "", _mPSSetHeapCacheDurationErr)
	}
	_mPSSetHeapCacheDuration(cmdBuf, seconds)
	return nil
}

// MPSSetHeapCacheDuration sets the timeout after which unused cached Metal heaps are released.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSetHeapCacheDuration(_:_:)
func MPSSetHeapCacheDuration(cmdBuf metal.MTLCommandBufferObject, seconds float64) {
	if callErr := tryMPSSetHeapCacheDuration(cmdBuf, seconds); callErr != nil {
		panic(callErr)
	}
}

var _mPSStateBatchIncrementReadCount func(batch *MPSStateBatch, amount int) uint
var _mPSStateBatchIncrementReadCountErr error

func tryMPSStateBatchIncrementReadCount(batch *MPSStateBatch, amount int) (uint, error) {
	if _mPSStateBatchIncrementReadCount == nil {
		return 0, symbolCallError("MPSStateBatchIncrementReadCount", "10.13.4", _mPSStateBatchIncrementReadCountErr)
	}
	return _mPSStateBatchIncrementReadCount(batch, amount), nil
}

// MPSStateBatchIncrementReadCount increments or decrements the read count of a state batch by a specified amount.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateBatchIncrementReadCount(_:_:)
func MPSStateBatchIncrementReadCount(batch *MPSStateBatch, amount int) uint {
	result, callErr := tryMPSStateBatchIncrementReadCount(batch, amount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mPSStateBatchResourceSize func(batch *MPSStateBatch) uint
var _mPSStateBatchResourceSizeErr error

func tryMPSStateBatchResourceSize(batch *MPSStateBatch) (uint, error) {
	if _mPSStateBatchResourceSize == nil {
		return 0, symbolCallError("MPSStateBatchResourceSize", "10.14", _mPSStateBatchResourceSizeErr)
	}
	return _mPSStateBatchResourceSize(batch), nil
}

// MPSStateBatchResourceSize returns the number of bytes used to allocate the specified state batch.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateBatchResourceSize(_:)
func MPSStateBatchResourceSize(batch *MPSStateBatch) uint {
	result, callErr := tryMPSStateBatchResourceSize(batch)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mPSStateBatchSynchronize func(batch *MPSStateBatch, cmdBuf metal.MTLCommandBufferObject)
var _mPSStateBatchSynchronizeErr error

func tryMPSStateBatchSynchronize(batch *MPSStateBatch, cmdBuf metal.MTLCommandBufferObject) error {
	if _mPSStateBatchSynchronize == nil {
		return symbolCallError("MPSStateBatchSynchronize", "10.13.4", _mPSStateBatchSynchronizeErr)
	}
	_mPSStateBatchSynchronize(batch, cmdBuf)
	return nil
}

// MPSStateBatchSynchronize removes any copy of the specified state batch from the device’s caches, and, if needed, invalidates any CPU caches.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateBatchSynchronize(_:_:)
func MPSStateBatchSynchronize(batch *MPSStateBatch, cmdBuf metal.MTLCommandBufferObject) {
	if callErr := tryMPSStateBatchSynchronize(batch, cmdBuf); callErr != nil {
		panic(callErr)
	}
}

var _mPSSupportsMTLDevice func(device metal.MTLDeviceObject) bool
var _mPSSupportsMTLDeviceErr error

func tryMPSSupportsMTLDevice(device metal.MTLDeviceObject) (bool, error) {
	if _mPSSupportsMTLDevice == nil {
		return false, symbolCallError("MPSSupportsMTLDevice", "10.13", _mPSSupportsMTLDeviceErr)
	}
	return _mPSSupportsMTLDevice(device), nil
}

// MPSSupportsMTLDevice determines whether the Metal Performance Shaders framework supports a Metal device.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSupportsMTLDevice(_:)
func MPSSupportsMTLDevice(device metal.MTLDeviceObject) bool {
	result, callErr := tryMPSSupportsMTLDevice(device)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_mPSGetImageType, &_mPSGetImageTypeErr, frameworkHandle, "MPSGetImageType", "")
	registerFunc(&_mPSGetPreferredDevice, &_mPSGetPreferredDeviceErr, frameworkHandle, "MPSGetPreferredDevice", "10.14.4")
	registerFunc(&_mPSHintTemporaryMemoryHighWaterMark, &_mPSHintTemporaryMemoryHighWaterMarkErr, frameworkHandle, "MPSHintTemporaryMemoryHighWaterMark", "")
	registerFunc(&_mPSImageBatchIncrementReadCount, &_mPSImageBatchIncrementReadCountErr, frameworkHandle, "MPSImageBatchIncrementReadCount", "10.13.4")
	registerFunc(&_mPSImageBatchIterate, &_mPSImageBatchIterateErr, frameworkHandle, "MPSImageBatchIterate", "10.15")
	registerFunc(&_mPSImageBatchResourceSize, &_mPSImageBatchResourceSizeErr, frameworkHandle, "MPSImageBatchResourceSize", "10.14")
	registerFunc(&_mPSImageBatchSynchronize, &_mPSImageBatchSynchronizeErr, frameworkHandle, "MPSImageBatchSynchronize", "10.13.4")
	registerFunc(&_mPSSetHeapCacheDuration, &_mPSSetHeapCacheDurationErr, frameworkHandle, "MPSSetHeapCacheDuration", "")
	registerFunc(&_mPSStateBatchIncrementReadCount, &_mPSStateBatchIncrementReadCountErr, frameworkHandle, "MPSStateBatchIncrementReadCount", "10.13.4")
	registerFunc(&_mPSStateBatchResourceSize, &_mPSStateBatchResourceSizeErr, frameworkHandle, "MPSStateBatchResourceSize", "10.14")
	registerFunc(&_mPSStateBatchSynchronize, &_mPSStateBatchSynchronizeErr, frameworkHandle, "MPSStateBatchSynchronize", "10.13.4")
	registerFunc(&_mPSSupportsMTLDevice, &_mPSSupportsMTLDeviceErr, frameworkHandle, "MPSSupportsMTLDevice", "10.13")
}

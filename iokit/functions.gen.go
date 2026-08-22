// Code generated from Apple documentation for iokit. DO NOT EDIT.

package iokit

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/kernel"
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
		return fmt.Sprintf("iokit: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("iokit: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("iokit: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("iokit: register symbol %s: %v", name, r)
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

var _iOAccelFindAccelerator func(framebuffer uint32, pAccelerator *uint32, pFramebufferIndex *uint32) kernel.IOReturn
var _iOAccelFindAcceleratorErr error

func tryIOAccelFindAccelerator(framebuffer uint32, pAccelerator *uint32, pFramebufferIndex *uint32) (kernel.IOReturn, error) {
	if _iOAccelFindAccelerator == nil {
		return *new(kernel.IOReturn), symbolCallError("IOAccelFindAccelerator", "10.0", _iOAccelFindAcceleratorErr)
	}
	return _iOAccelFindAccelerator(framebuffer, pAccelerator, pFramebufferIndex), nil
}

// IOAccelFindAccelerator.
//
// See: https://developer.apple.com/documentation/iokit/1502405-ioaccelfindaccelerator
func IOAccelFindAccelerator(framebuffer uint32, pAccelerator *uint32, pFramebufferIndex *uint32) kernel.IOReturn {
	result, callErr := tryIOAccelFindAccelerator(framebuffer, pAccelerator, pFramebufferIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOAllowPowerChange func(kernelPort uint32, notificationID int) kernel.IOReturn
var _iOAllowPowerChangeErr error

func tryIOAllowPowerChange(kernelPort uint32, notificationID int) (kernel.IOReturn, error) {
	if _iOAllowPowerChange == nil {
		return *new(kernel.IOReturn), symbolCallError("IOAllowPowerChange", "10.0", _iOAllowPowerChangeErr)
	}
	return _iOAllowPowerChange(kernelPort, notificationID), nil
}

// IOAllowPowerChange the caller acknowledges notification of a power state change on a device it has registered for notifications for via IORegisterForSystemPower or IORegisterApp.
//
// See: https://developer.apple.com/documentation/iokit/1557064-ioallowpowerchange
func IOAllowPowerChange(kernelPort uint32, notificationID int) kernel.IOReturn {
	result, callErr := tryIOAllowPowerChange(kernelPort, notificationID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOBSDNameMatching func(mainPort uint32, options uint32, bsdName string) corefoundation.CFMutableDictionary
var _iOBSDNameMatchingErr error

func tryIOBSDNameMatching(mainPort uint32, options uint32, bsdName string) (corefoundation.CFMutableDictionary, error) {
	if _iOBSDNameMatching == nil {
		return *new(corefoundation.CFMutableDictionary), symbolCallError("IOBSDNameMatching", "10.0", _iOBSDNameMatchingErr)
	}
	return _iOBSDNameMatching(mainPort, options, bsdName), nil
}

// IOBSDNameMatching create a matching dictionary that specifies an IOService match based on BSD device name.
//
// See: https://developer.apple.com/documentation/iokit/1514486-iobsdnamematching
func IOBSDNameMatching(mainPort uint32, options uint32, bsdName string) corefoundation.CFMutableDictionary {
	result, callErr := tryIOBSDNameMatching(mainPort, options, bsdName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOCFSerialize func(object corefoundation.CFTypeRef, options uint) corefoundation.CFDataRef
var _iOCFSerializeErr error

func tryIOCFSerialize(object corefoundation.CFTypeRef, options uint) (corefoundation.CFDataRef, error) {
	if _iOCFSerialize == nil {
		return *new(corefoundation.CFDataRef), symbolCallError("IOCFSerialize", "10.0", _iOCFSerializeErr)
	}
	return _iOCFSerialize(object, options), nil
}

// IOCFSerialize.
//
// See: https://developer.apple.com/documentation/iokit/1403329-iocfserialize
func IOCFSerialize(object corefoundation.CFTypeRef, options uint) corefoundation.CFDataRef {
	result, callErr := tryIOCFSerialize(object, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOCFUnserialize func(buffer string, allocator corefoundation.CFAllocatorRef, options uint, errorString *corefoundation.CFStringRef) corefoundation.CFTypeRef
var _iOCFUnserializeErr error

func tryIOCFUnserialize(buffer string, allocator corefoundation.CFAllocatorRef, options uint, errorString *corefoundation.CFStringRef) (corefoundation.CFTypeRef, error) {
	if _iOCFUnserialize == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("IOCFUnserialize", "10.0", _iOCFUnserializeErr)
	}
	return _iOCFUnserialize(buffer, allocator, options, errorString), nil
}

// IOCFUnserialize.
//
// See: https://developer.apple.com/documentation/iokit/1514265-iocfunserialize
func IOCFUnserialize(buffer string, allocator corefoundation.CFAllocatorRef, options uint, errorString *corefoundation.CFStringRef) corefoundation.CFTypeRef {
	result, callErr := tryIOCFUnserialize(buffer, allocator, options, errorString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOCFUnserializeBinary func(buffer string, bufferSize unsafe.Pointer, allocator corefoundation.CFAllocatorRef, options uint, errorString *corefoundation.CFStringRef) corefoundation.CFTypeRef
var _iOCFUnserializeBinaryErr error

func tryIOCFUnserializeBinary(buffer string, bufferSize unsafe.Pointer, allocator corefoundation.CFAllocatorRef, options uint, errorString *corefoundation.CFStringRef) (corefoundation.CFTypeRef, error) {
	if _iOCFUnserializeBinary == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("IOCFUnserializeBinary", "10.10", _iOCFUnserializeBinaryErr)
	}
	return _iOCFUnserializeBinary(buffer, bufferSize, allocator, options, errorString), nil
}

// IOCFUnserializeBinary.
//
// See: https://developer.apple.com/documentation/iokit/1514876-iocfunserializebinary
func IOCFUnserializeBinary(buffer string, bufferSize unsafe.Pointer, allocator corefoundation.CFAllocatorRef, options uint, errorString *corefoundation.CFStringRef) corefoundation.CFTypeRef {
	result, callErr := tryIOCFUnserializeBinary(buffer, bufferSize, allocator, options, errorString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOCFUnserializeWithSize func(buffer string, bufferSize unsafe.Pointer, allocator corefoundation.CFAllocatorRef, options uint, errorString *corefoundation.CFStringRef) corefoundation.CFTypeRef
var _iOCFUnserializeWithSizeErr error

func tryIOCFUnserializeWithSize(buffer string, bufferSize unsafe.Pointer, allocator corefoundation.CFAllocatorRef, options uint, errorString *corefoundation.CFStringRef) (corefoundation.CFTypeRef, error) {
	if _iOCFUnserializeWithSize == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("IOCFUnserializeWithSize", "10.10", _iOCFUnserializeWithSizeErr)
	}
	return _iOCFUnserializeWithSize(buffer, bufferSize, allocator, options, errorString), nil
}

// IOCFUnserializeWithSize.
//
// See: https://developer.apple.com/documentation/iokit/1514745-iocfunserializewithsize
func IOCFUnserializeWithSize(buffer string, bufferSize unsafe.Pointer, allocator corefoundation.CFAllocatorRef, options uint, errorString *corefoundation.CFStringRef) corefoundation.CFTypeRef {
	result, callErr := tryIOCFUnserializeWithSize(buffer, bufferSize, allocator, options, errorString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOCancelPowerChange func(kernelPort uint32, notificationID int) kernel.IOReturn
var _iOCancelPowerChangeErr error

func tryIOCancelPowerChange(kernelPort uint32, notificationID int) (kernel.IOReturn, error) {
	if _iOCancelPowerChange == nil {
		return *new(kernel.IOReturn), symbolCallError("IOCancelPowerChange", "10.0", _iOCancelPowerChangeErr)
	}
	return _iOCancelPowerChange(kernelPort, notificationID), nil
}

// IOCancelPowerChange the caller denies an idle system sleep power state change.
//
// See: https://developer.apple.com/documentation/iokit/1557115-iocancelpowerchange
func IOCancelPowerChange(kernelPort uint32, notificationID int) kernel.IOReturn {
	result, callErr := tryIOCancelPowerChange(kernelPort, notificationID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOCatalogueGetData func(mainPort uint32, flag uint32, buffer string, size *uint32) int32
var _iOCatalogueGetDataErr error

func tryIOCatalogueGetData(mainPort uint32, flag uint32, buffer string, size *uint32) (int32, error) {
	if _iOCatalogueGetData == nil {
		return 0, symbolCallError("IOCatalogueGetData", "10.0", _iOCatalogueGetDataErr)
	}
	return _iOCatalogueGetData(mainPort, flag, buffer, size), nil
}

// IOCatalogueGetData.
//
// See: https://developer.apple.com/documentation/iokit/1514233-iocataloguegetdata
func IOCatalogueGetData(mainPort uint32, flag uint32, buffer string, size *uint32) int32 {
	result, callErr := tryIOCatalogueGetData(mainPort, flag, buffer, size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOCatalogueModuleLoaded func(mainPort uint32, name string) int32
var _iOCatalogueModuleLoadedErr error

func tryIOCatalogueModuleLoaded(mainPort uint32, name string) (int32, error) {
	if _iOCatalogueModuleLoaded == nil {
		return 0, symbolCallError("IOCatalogueModuleLoaded", "10.0", _iOCatalogueModuleLoadedErr)
	}
	return _iOCatalogueModuleLoaded(mainPort, name), nil
}

// IOCatalogueModuleLoaded.
//
// See: https://developer.apple.com/documentation/iokit/1514886-iocataloguemoduleloaded
func IOCatalogueModuleLoaded(mainPort uint32, name string) int32 {
	result, callErr := tryIOCatalogueModuleLoaded(mainPort, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOCatalogueReset func(mainPort uint32, flag uint32) int32
var _iOCatalogueResetErr error

func tryIOCatalogueReset(mainPort uint32, flag uint32) (int32, error) {
	if _iOCatalogueReset == nil {
		return 0, symbolCallError("IOCatalogueReset", "10.0", _iOCatalogueResetErr)
	}
	return _iOCatalogueReset(mainPort, flag), nil
}

// IOCatalogueReset.
//
// See: https://developer.apple.com/documentation/iokit/1514702-iocataloguereset
func IOCatalogueReset(mainPort uint32, flag uint32) int32 {
	result, callErr := tryIOCatalogueReset(mainPort, flag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOCatalogueSendData func(mainPort uint32, flag uint32, buffer string, size uint32) int32
var _iOCatalogueSendDataErr error

func tryIOCatalogueSendData(mainPort uint32, flag uint32, buffer string, size uint32) (int32, error) {
	if _iOCatalogueSendData == nil {
		return 0, symbolCallError("IOCatalogueSendData", "10.0", _iOCatalogueSendDataErr)
	}
	return _iOCatalogueSendData(mainPort, flag, buffer, size), nil
}

// IOCatalogueSendData.
//
// See: https://developer.apple.com/documentation/iokit/1514405-iocataloguesenddata
func IOCatalogueSendData(mainPort uint32, flag uint32, buffer string, size uint32) int32 {
	result, callErr := tryIOCatalogueSendData(mainPort, flag, buffer, size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOCatalogueTerminate func(mainPort uint32, flag uint32, description string) int32
var _iOCatalogueTerminateErr error

func tryIOCatalogueTerminate(mainPort uint32, flag uint32, description string) (int32, error) {
	if _iOCatalogueTerminate == nil {
		return 0, symbolCallError("IOCatalogueTerminate", "10.0", _iOCatalogueTerminateErr)
	}
	return _iOCatalogueTerminate(mainPort, flag, description), nil
}

// IOCatalogueTerminate.
//
// See: https://developer.apple.com/documentation/iokit/1514665-iocatalogueterminate
func IOCatalogueTerminate(mainPort uint32, flag uint32, description string) int32 {
	result, callErr := tryIOCatalogueTerminate(mainPort, flag, description)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectAddClient func(connect uint32, client uint32) int32
var _iOConnectAddClientErr error

func tryIOConnectAddClient(connect uint32, client uint32) (int32, error) {
	if _iOConnectAddClient == nil {
		return 0, symbolCallError("IOConnectAddClient", "10.0", _iOConnectAddClientErr)
	}
	return _iOConnectAddClient(connect, client), nil
}

// IOConnectAddClient inform a connection of a second connection.
//
// See: https://developer.apple.com/documentation/iokit/1514609-ioconnectaddclient
func IOConnectAddClient(connect uint32, client uint32) int32 {
	result, callErr := tryIOConnectAddClient(connect, client)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectAddRef func(connect uint32) int32
var _iOConnectAddRefErr error

func tryIOConnectAddRef(connect uint32) (int32, error) {
	if _iOConnectAddRef == nil {
		return 0, symbolCallError("IOConnectAddRef", "10.0", _iOConnectAddRefErr)
	}
	return _iOConnectAddRef(connect), nil
}

// IOConnectAddRef adds a reference to the connect handle.
//
// See: https://developer.apple.com/documentation/iokit/1514739-ioconnectaddref
func IOConnectAddRef(connect uint32) int32 {
	result, callErr := tryIOConnectAddRef(connect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectCallAsyncMethod func(connection uint32, selector uint32, wake_port uint32, reference *uint64, referenceCnt uint32, input *uint64, inputCnt uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, output *uint64, outputCnt *uint32, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) int32
var _iOConnectCallAsyncMethodErr error

func tryIOConnectCallAsyncMethod(connection uint32, selector uint32, wake_port uint32, reference *uint64, referenceCnt uint32, input *uint64, inputCnt uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, output *uint64, outputCnt *uint32, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) (int32, error) {
	if _iOConnectCallAsyncMethod == nil {
		return 0, symbolCallError("IOConnectCallAsyncMethod", "10.5", _iOConnectCallAsyncMethodErr)
	}
	return _iOConnectCallAsyncMethod(connection, selector, wake_port, reference, referenceCnt, input, inputCnt, inputStruct, inputStructCnt, output, outputCnt, outputStruct, outputStructCnt), nil
}

// IOConnectCallAsyncMethod.
//
// See: https://developer.apple.com/documentation/iokit/1514418-ioconnectcallasyncmethod
func IOConnectCallAsyncMethod(connection uint32, selector uint32, wake_port uint32, reference *uint64, referenceCnt uint32, input *uint64, inputCnt uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, output *uint64, outputCnt *uint32, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) int32 {
	result, callErr := tryIOConnectCallAsyncMethod(connection, selector, wake_port, reference, referenceCnt, input, inputCnt, inputStruct, inputStructCnt, output, outputCnt, outputStruct, outputStructCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectCallAsyncScalarMethod func(connection uint32, selector uint32, wake_port uint32, reference *uint64, referenceCnt uint32, input *uint64, inputCnt uint32, output *uint64, outputCnt *uint32) int32
var _iOConnectCallAsyncScalarMethodErr error

func tryIOConnectCallAsyncScalarMethod(connection uint32, selector uint32, wake_port uint32, reference *uint64, referenceCnt uint32, input *uint64, inputCnt uint32, output *uint64, outputCnt *uint32) (int32, error) {
	if _iOConnectCallAsyncScalarMethod == nil {
		return 0, symbolCallError("IOConnectCallAsyncScalarMethod", "10.5", _iOConnectCallAsyncScalarMethodErr)
	}
	return _iOConnectCallAsyncScalarMethod(connection, selector, wake_port, reference, referenceCnt, input, inputCnt, output, outputCnt), nil
}

// IOConnectCallAsyncScalarMethod.
//
// See: https://developer.apple.com/documentation/iokit/1514884-ioconnectcallasyncscalarmethod
func IOConnectCallAsyncScalarMethod(connection uint32, selector uint32, wake_port uint32, reference *uint64, referenceCnt uint32, input *uint64, inputCnt uint32, output *uint64, outputCnt *uint32) int32 {
	result, callErr := tryIOConnectCallAsyncScalarMethod(connection, selector, wake_port, reference, referenceCnt, input, inputCnt, output, outputCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectCallAsyncStructMethod func(connection uint32, selector uint32, wake_port uint32, reference *uint64, referenceCnt uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) int32
var _iOConnectCallAsyncStructMethodErr error

func tryIOConnectCallAsyncStructMethod(connection uint32, selector uint32, wake_port uint32, reference *uint64, referenceCnt uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) (int32, error) {
	if _iOConnectCallAsyncStructMethod == nil {
		return 0, symbolCallError("IOConnectCallAsyncStructMethod", "10.5", _iOConnectCallAsyncStructMethodErr)
	}
	return _iOConnectCallAsyncStructMethod(connection, selector, wake_port, reference, referenceCnt, inputStruct, inputStructCnt, outputStruct, outputStructCnt), nil
}

// IOConnectCallAsyncStructMethod.
//
// See: https://developer.apple.com/documentation/iokit/1514403-ioconnectcallasyncstructmethod
func IOConnectCallAsyncStructMethod(connection uint32, selector uint32, wake_port uint32, reference *uint64, referenceCnt uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) int32 {
	result, callErr := tryIOConnectCallAsyncStructMethod(connection, selector, wake_port, reference, referenceCnt, inputStruct, inputStructCnt, outputStruct, outputStructCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectCallMethod func(connection uint32, selector uint32, input *uint64, inputCnt uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, output *uint64, outputCnt *uint32, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) int32
var _iOConnectCallMethodErr error

func tryIOConnectCallMethod(connection uint32, selector uint32, input *uint64, inputCnt uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, output *uint64, outputCnt *uint32, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) (int32, error) {
	if _iOConnectCallMethod == nil {
		return 0, symbolCallError("IOConnectCallMethod", "10.5", _iOConnectCallMethodErr)
	}
	return _iOConnectCallMethod(connection, selector, input, inputCnt, inputStruct, inputStructCnt, output, outputCnt, outputStruct, outputStructCnt), nil
}

// IOConnectCallMethod.
//
// See: https://developer.apple.com/documentation/iokit/1514240-ioconnectcallmethod
func IOConnectCallMethod(connection uint32, selector uint32, input *uint64, inputCnt uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, output *uint64, outputCnt *uint32, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) int32 {
	result, callErr := tryIOConnectCallMethod(connection, selector, input, inputCnt, inputStruct, inputStructCnt, output, outputCnt, outputStruct, outputStructCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectCallScalarMethod func(connection uint32, selector uint32, input *uint64, inputCnt uint32, output *uint64, outputCnt *uint32) int32
var _iOConnectCallScalarMethodErr error

func tryIOConnectCallScalarMethod(connection uint32, selector uint32, input *uint64, inputCnt uint32, output *uint64, outputCnt *uint32) (int32, error) {
	if _iOConnectCallScalarMethod == nil {
		return 0, symbolCallError("IOConnectCallScalarMethod", "10.5", _iOConnectCallScalarMethodErr)
	}
	return _iOConnectCallScalarMethod(connection, selector, input, inputCnt, output, outputCnt), nil
}

// IOConnectCallScalarMethod.
//
// See: https://developer.apple.com/documentation/iokit/1514793-ioconnectcallscalarmethod
func IOConnectCallScalarMethod(connection uint32, selector uint32, input *uint64, inputCnt uint32, output *uint64, outputCnt *uint32) int32 {
	result, callErr := tryIOConnectCallScalarMethod(connection, selector, input, inputCnt, output, outputCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectCallStructMethod func(connection uint32, selector uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) int32
var _iOConnectCallStructMethodErr error

func tryIOConnectCallStructMethod(connection uint32, selector uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) (int32, error) {
	if _iOConnectCallStructMethod == nil {
		return 0, symbolCallError("IOConnectCallStructMethod", "10.5", _iOConnectCallStructMethodErr)
	}
	return _iOConnectCallStructMethod(connection, selector, inputStruct, inputStructCnt, outputStruct, outputStructCnt), nil
}

// IOConnectCallStructMethod.
//
// See: https://developer.apple.com/documentation/iokit/1514274-ioconnectcallstructmethod
func IOConnectCallStructMethod(connection uint32, selector uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) int32 {
	result, callErr := tryIOConnectCallStructMethod(connection, selector, inputStruct, inputStructCnt, outputStruct, outputStructCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectGetService func(connect uint32, service *uint32) int32
var _iOConnectGetServiceErr error

func tryIOConnectGetService(connect uint32, service *uint32) (int32, error) {
	if _iOConnectGetService == nil {
		return 0, symbolCallError("IOConnectGetService", "10.0", _iOConnectGetServiceErr)
	}
	return _iOConnectGetService(connect, service), nil
}

// IOConnectGetService returns the IOService a connect handle was opened on.
//
// See: https://developer.apple.com/documentation/iokit/1514438-ioconnectgetservice
func IOConnectGetService(connect uint32, service *uint32) int32 {
	result, callErr := tryIOConnectGetService(connect, service)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectMapMemory func(connect uint32, memoryType uint32, intoTask uint32, atAddress *uint64, ofSize *uint64, options uint32) int32
var _iOConnectMapMemoryErr error

func tryIOConnectMapMemory(connect uint32, memoryType uint32, intoTask uint32, atAddress *uint64, ofSize *uint64, options uint32) (int32, error) {
	if _iOConnectMapMemory == nil {
		return 0, symbolCallError("IOConnectMapMemory", "10.0", _iOConnectMapMemoryErr)
	}
	return _iOConnectMapMemory(connect, memoryType, intoTask, atAddress, ofSize, options), nil
}

// IOConnectMapMemory map hardware or shared memory into the caller's task.
//
// See: https://developer.apple.com/documentation/iokit/1514377-ioconnectmapmemory
func IOConnectMapMemory(connect uint32, memoryType uint32, intoTask uint32, atAddress *uint64, ofSize *uint64, options uint32) int32 {
	result, callErr := tryIOConnectMapMemory(connect, memoryType, intoTask, atAddress, ofSize, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectMapMemory64 func(connect uint32, memoryType uint32, intoTask uint32, atAddress *uint64, ofSize *uint64, options uint32) int32
var _iOConnectMapMemory64Err error

func tryIOConnectMapMemory64(connect uint32, memoryType uint32, intoTask uint32, atAddress *uint64, ofSize *uint64, options uint32) (int32, error) {
	if _iOConnectMapMemory64 == nil {
		return 0, symbolCallError("IOConnectMapMemory64", "10.5", _iOConnectMapMemory64Err)
	}
	return _iOConnectMapMemory64(connect, memoryType, intoTask, atAddress, ofSize, options), nil
}

// IOConnectMapMemory64 map hardware or shared memory into the caller's task.
//
// See: https://developer.apple.com/documentation/iokit/1514862-ioconnectmapmemory64
func IOConnectMapMemory64(connect uint32, memoryType uint32, intoTask uint32, atAddress *uint64, ofSize *uint64, options uint32) int32 {
	result, callErr := tryIOConnectMapMemory64(connect, memoryType, intoTask, atAddress, ofSize, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectRelease func(connect uint32) int32
var _iOConnectReleaseErr error

func tryIOConnectRelease(connect uint32) (int32, error) {
	if _iOConnectRelease == nil {
		return 0, symbolCallError("IOConnectRelease", "10.0", _iOConnectReleaseErr)
	}
	return _iOConnectRelease(connect), nil
}

// IOConnectRelease remove a reference to the connect handle.
//
// See: https://developer.apple.com/documentation/iokit/1514511-ioconnectrelease
func IOConnectRelease(connect uint32) int32 {
	result, callErr := tryIOConnectRelease(connect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectSetCFProperties func(connect uint32, properties corefoundation.CFTypeRef) int32
var _iOConnectSetCFPropertiesErr error

func tryIOConnectSetCFProperties(connect uint32, properties corefoundation.CFTypeRef) (int32, error) {
	if _iOConnectSetCFProperties == nil {
		return 0, symbolCallError("IOConnectSetCFProperties", "10.0", _iOConnectSetCFPropertiesErr)
	}
	return _iOConnectSetCFProperties(connect, properties), nil
}

// IOConnectSetCFProperties set CF container based properties on a connection.
//
// See: https://developer.apple.com/documentation/iokit/1514713-ioconnectsetcfproperties
func IOConnectSetCFProperties(connect uint32, properties corefoundation.CFTypeRef) int32 {
	result, callErr := tryIOConnectSetCFProperties(connect, properties)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectSetCFProperty func(connect uint32, propertyName corefoundation.CFStringRef, property corefoundation.CFTypeRef) int32
var _iOConnectSetCFPropertyErr error

func tryIOConnectSetCFProperty(connect uint32, propertyName corefoundation.CFStringRef, property corefoundation.CFTypeRef) (int32, error) {
	if _iOConnectSetCFProperty == nil {
		return 0, symbolCallError("IOConnectSetCFProperty", "10.0", _iOConnectSetCFPropertyErr)
	}
	return _iOConnectSetCFProperty(connect, propertyName, property), nil
}

// IOConnectSetCFProperty set a CF container based property on a connection.
//
// See: https://developer.apple.com/documentation/iokit/1514796-ioconnectsetcfproperty
func IOConnectSetCFProperty(connect uint32, propertyName corefoundation.CFStringRef, property corefoundation.CFTypeRef) int32 {
	result, callErr := tryIOConnectSetCFProperty(connect, propertyName, property)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectSetNotificationPort func(connect uint32, type_ uint32, port uint32, reference unsafe.Pointer) int32
var _iOConnectSetNotificationPortErr error

func tryIOConnectSetNotificationPort(connect uint32, type_ uint32, port uint32, reference unsafe.Pointer) (int32, error) {
	if _iOConnectSetNotificationPort == nil {
		return 0, symbolCallError("IOConnectSetNotificationPort", "10.0", _iOConnectSetNotificationPortErr)
	}
	return _iOConnectSetNotificationPort(connect, type_, port, reference), nil
}

// IOConnectSetNotificationPort set a port to receive family specific notifications.
//
// See: https://developer.apple.com/documentation/iokit/1514541-ioconnectsetnotificationport
func IOConnectSetNotificationPort(connect uint32, type_ uint32, port uint32, reference unsafe.Pointer) int32 {
	result, callErr := tryIOConnectSetNotificationPort(connect, type_, port, reference)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectTrap0 func(connect uint32, index uint32) int32
var _iOConnectTrap0Err error

func tryIOConnectTrap0(connect uint32, index uint32) (int32, error) {
	if _iOConnectTrap0 == nil {
		return 0, symbolCallError("IOConnectTrap0", "10.0", _iOConnectTrap0Err)
	}
	return _iOConnectTrap0(connect, index), nil
}

// IOConnectTrap0.
//
// See: https://developer.apple.com/documentation/iokit/1514674-ioconnecttrap0
func IOConnectTrap0(connect uint32, index uint32) int32 {
	result, callErr := tryIOConnectTrap0(connect, index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectTrap1 func(connect uint32, index uint32, p1 unsafe.Pointer) int32
var _iOConnectTrap1Err error

func tryIOConnectTrap1(connect uint32, index uint32, p1 unsafe.Pointer) (int32, error) {
	if _iOConnectTrap1 == nil {
		return 0, symbolCallError("IOConnectTrap1", "10.0", _iOConnectTrap1Err)
	}
	return _iOConnectTrap1(connect, index, p1), nil
}

// IOConnectTrap1.
//
// See: https://developer.apple.com/documentation/iokit/1514816-ioconnecttrap1
func IOConnectTrap1(connect uint32, index uint32, p1 unsafe.Pointer) int32 {
	result, callErr := tryIOConnectTrap1(connect, index, p1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectTrap2 func(connect uint32, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer) int32
var _iOConnectTrap2Err error

func tryIOConnectTrap2(connect uint32, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer) (int32, error) {
	if _iOConnectTrap2 == nil {
		return 0, symbolCallError("IOConnectTrap2", "10.0", _iOConnectTrap2Err)
	}
	return _iOConnectTrap2(connect, index, p1, p2), nil
}

// IOConnectTrap2.
//
// See: https://developer.apple.com/documentation/iokit/1514354-ioconnecttrap2
func IOConnectTrap2(connect uint32, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer) int32 {
	result, callErr := tryIOConnectTrap2(connect, index, p1, p2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectTrap3 func(connect uint32, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) int32
var _iOConnectTrap3Err error

func tryIOConnectTrap3(connect uint32, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) (int32, error) {
	if _iOConnectTrap3 == nil {
		return 0, symbolCallError("IOConnectTrap3", "10.0", _iOConnectTrap3Err)
	}
	return _iOConnectTrap3(connect, index, p1, p2, p3), nil
}

// IOConnectTrap3.
//
// See: https://developer.apple.com/documentation/iokit/1514833-ioconnecttrap3
func IOConnectTrap3(connect uint32, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) int32 {
	result, callErr := tryIOConnectTrap3(connect, index, p1, p2, p3)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectTrap4 func(connect uint32, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer, p4 unsafe.Pointer) int32
var _iOConnectTrap4Err error

func tryIOConnectTrap4(connect uint32, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer, p4 unsafe.Pointer) (int32, error) {
	if _iOConnectTrap4 == nil {
		return 0, symbolCallError("IOConnectTrap4", "10.0", _iOConnectTrap4Err)
	}
	return _iOConnectTrap4(connect, index, p1, p2, p3, p4), nil
}

// IOConnectTrap4.
//
// See: https://developer.apple.com/documentation/iokit/1514410-ioconnecttrap4
func IOConnectTrap4(connect uint32, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer, p4 unsafe.Pointer) int32 {
	result, callErr := tryIOConnectTrap4(connect, index, p1, p2, p3, p4)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectTrap5 func(connect uint32, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer, p4 unsafe.Pointer, p5 unsafe.Pointer) int32
var _iOConnectTrap5Err error

func tryIOConnectTrap5(connect uint32, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer, p4 unsafe.Pointer, p5 unsafe.Pointer) (int32, error) {
	if _iOConnectTrap5 == nil {
		return 0, symbolCallError("IOConnectTrap5", "10.0", _iOConnectTrap5Err)
	}
	return _iOConnectTrap5(connect, index, p1, p2, p3, p4, p5), nil
}

// IOConnectTrap5.
//
// See: https://developer.apple.com/documentation/iokit/1514346-ioconnecttrap5
func IOConnectTrap5(connect uint32, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer, p4 unsafe.Pointer, p5 unsafe.Pointer) int32 {
	result, callErr := tryIOConnectTrap5(connect, index, p1, p2, p3, p4, p5)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectTrap6 func(connect uint32, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer, p4 unsafe.Pointer, p5 unsafe.Pointer, p6 unsafe.Pointer) int32
var _iOConnectTrap6Err error

func tryIOConnectTrap6(connect uint32, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer, p4 unsafe.Pointer, p5 unsafe.Pointer, p6 unsafe.Pointer) (int32, error) {
	if _iOConnectTrap6 == nil {
		return 0, symbolCallError("IOConnectTrap6", "10.0", _iOConnectTrap6Err)
	}
	return _iOConnectTrap6(connect, index, p1, p2, p3, p4, p5, p6), nil
}

// IOConnectTrap6.
//
// See: https://developer.apple.com/documentation/iokit/1514493-ioconnecttrap6
func IOConnectTrap6(connect uint32, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer, p4 unsafe.Pointer, p5 unsafe.Pointer, p6 unsafe.Pointer) int32 {
	result, callErr := tryIOConnectTrap6(connect, index, p1, p2, p3, p4, p5, p6)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectUnmapMemory func(connect uint32, memoryType uint32, fromTask uint32, atAddress uint64) int32
var _iOConnectUnmapMemoryErr error

func tryIOConnectUnmapMemory(connect uint32, memoryType uint32, fromTask uint32, atAddress uint64) (int32, error) {
	if _iOConnectUnmapMemory == nil {
		return 0, symbolCallError("IOConnectUnmapMemory", "10.0", _iOConnectUnmapMemoryErr)
	}
	return _iOConnectUnmapMemory(connect, memoryType, fromTask, atAddress), nil
}

// IOConnectUnmapMemory remove a mapping made with IOConnectMapMemory.
//
// See: https://developer.apple.com/documentation/iokit/1514527-ioconnectunmapmemory
func IOConnectUnmapMemory(connect uint32, memoryType uint32, fromTask uint32, atAddress uint64) int32 {
	result, callErr := tryIOConnectUnmapMemory(connect, memoryType, fromTask, atAddress)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOConnectUnmapMemory64 func(connect uint32, memoryType uint32, fromTask uint32, atAddress uint64) int32
var _iOConnectUnmapMemory64Err error

func tryIOConnectUnmapMemory64(connect uint32, memoryType uint32, fromTask uint32, atAddress uint64) (int32, error) {
	if _iOConnectUnmapMemory64 == nil {
		return 0, symbolCallError("IOConnectUnmapMemory64", "10.5", _iOConnectUnmapMemory64Err)
	}
	return _iOConnectUnmapMemory64(connect, memoryType, fromTask, atAddress), nil
}

// IOConnectUnmapMemory64 remove a mapping made with IOConnectMapMemory64.
//
// See: https://developer.apple.com/documentation/iokit/1514760-ioconnectunmapmemory64
func IOConnectUnmapMemory64(connect uint32, memoryType uint32, fromTask uint32, atAddress uint64) int32 {
	result, callErr := tryIOConnectUnmapMemory64(connect, memoryType, fromTask, atAddress)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOCopySystemLoadAdvisoryDetailed func() corefoundation.CFDictionaryRef
var _iOCopySystemLoadAdvisoryDetailedErr error

func tryIOCopySystemLoadAdvisoryDetailed() (corefoundation.CFDictionaryRef, error) {
	if _iOCopySystemLoadAdvisoryDetailed == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("IOCopySystemLoadAdvisoryDetailed", "10.6", _iOCopySystemLoadAdvisoryDetailedErr)
	}
	return _iOCopySystemLoadAdvisoryDetailed(), nil
}

// IOCopySystemLoadAdvisoryDetailed indicates how user activity, battery level, and thermal level each contribute to the overall "SystemLoadAdvisory" level.
//
// See: https://developer.apple.com/documentation/iokit/1557099-iocopysystemloadadvisorydetailed
func IOCopySystemLoadAdvisoryDetailed() corefoundation.CFDictionaryRef {
	result, callErr := tryIOCopySystemLoadAdvisoryDetailed()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOCreatePlugInInterfaceForService func(service uint32, pluginType corefoundation.CFUUIDRef, interfaceType corefoundation.CFUUIDRef, theInterface *IOCFPlugInInterface, theScore int32) int32
var _iOCreatePlugInInterfaceForServiceErr error

func tryIOCreatePlugInInterfaceForService(service uint32, pluginType corefoundation.CFUUIDRef, interfaceType corefoundation.CFUUIDRef, theInterface *IOCFPlugInInterface, theScore int32) (int32, error) {
	if _iOCreatePlugInInterfaceForService == nil {
		return 0, symbolCallError("IOCreatePlugInInterfaceForService", "10.0", _iOCreatePlugInInterfaceForServiceErr)
	}
	return _iOCreatePlugInInterfaceForService(service, pluginType, interfaceType, theInterface, theScore), nil
}

// IOCreatePlugInInterfaceForService.
//
// See: https://developer.apple.com/documentation/iokit/1412429-iocreateplugininterfaceforservic
func IOCreatePlugInInterfaceForService(service uint32, pluginType corefoundation.CFUUIDRef, interfaceType corefoundation.CFUUIDRef, theInterface *IOCFPlugInInterface, theScore int32) int32 {
	result, callErr := tryIOCreatePlugInInterfaceForService(service, pluginType, interfaceType, theInterface, theScore)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOCreateReceivePort func(msgType uint32, recvPort *uint32) int32
var _iOCreateReceivePortErr error

func tryIOCreateReceivePort(msgType uint32, recvPort *uint32) (int32, error) {
	if _iOCreateReceivePort == nil {
		return 0, symbolCallError("IOCreateReceivePort", "10.0", _iOCreateReceivePortErr)
	}
	return _iOCreateReceivePort(msgType, recvPort), nil
}

// IOCreateReceivePort creates and returns a mach port suitable for receiving IOKit messages of the specified type.
//
// See: https://developer.apple.com/documentation/iokit/1514698-iocreatereceiveport
func IOCreateReceivePort(msgType uint32, recvPort *uint32) int32 {
	result, callErr := tryIOCreateReceivePort(msgType, recvPort)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODataQueueAllocateNotificationPort func() uint32
var _iODataQueueAllocateNotificationPortErr error

func tryIODataQueueAllocateNotificationPort() (uint32, error) {
	if _iODataQueueAllocateNotificationPort == nil {
		return 0, symbolCallError("IODataQueueAllocateNotificationPort", "10.0", _iODataQueueAllocateNotificationPortErr)
	}
	return _iODataQueueAllocateNotificationPort(), nil
}

// IODataQueueAllocateNotificationPort allocates and returns a new mach port able to receive data available notifications from an IODataQueue.
//
// See: https://developer.apple.com/documentation/iokit/1514495-iodataqueueallocatenotificationp
func IODataQueueAllocateNotificationPort() uint32 {
	result, callErr := tryIODataQueueAllocateNotificationPort()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODataQueueDataAvailable func(dataQueue *IODataQueueMemory) bool
var _iODataQueueDataAvailableErr error

func tryIODataQueueDataAvailable(dataQueue *IODataQueueMemory) (bool, error) {
	if _iODataQueueDataAvailable == nil {
		return false, symbolCallError("IODataQueueDataAvailable", "10.0", _iODataQueueDataAvailableErr)
	}
	return _iODataQueueDataAvailable(dataQueue), nil
}

// IODataQueueDataAvailable used to determine if more data is avilable on the queue.
//
// See: https://developer.apple.com/documentation/iokit/1514386-iodataqueuedataavailable
func IODataQueueDataAvailable(dataQueue *IODataQueueMemory) bool {
	result, callErr := tryIODataQueueDataAvailable(dataQueue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODataQueueDequeue func(dataQueue *IODataQueueMemory, data unsafe.Pointer, dataSize *uint32) kernel.IOReturn
var _iODataQueueDequeueErr error

func tryIODataQueueDequeue(dataQueue *IODataQueueMemory, data unsafe.Pointer, dataSize *uint32) (kernel.IOReturn, error) {
	if _iODataQueueDequeue == nil {
		return *new(kernel.IOReturn), symbolCallError("IODataQueueDequeue", "10.0", _iODataQueueDequeueErr)
	}
	return _iODataQueueDequeue(dataQueue, data, dataSize), nil
}

// IODataQueueDequeue dequeues the next available entry on the queue and copies it into the given data pointer.
//
// See: https://developer.apple.com/documentation/iokit/1514287-iodataqueuedequeue
func IODataQueueDequeue(dataQueue *IODataQueueMemory, data unsafe.Pointer, dataSize *uint32) kernel.IOReturn {
	result, callErr := tryIODataQueueDequeue(dataQueue, data, dataSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODataQueueEnqueue func(dataQueue *IODataQueueMemory, data unsafe.Pointer, dataSize uint32) kernel.IOReturn
var _iODataQueueEnqueueErr error

func tryIODataQueueEnqueue(dataQueue *IODataQueueMemory, data unsafe.Pointer, dataSize uint32) (kernel.IOReturn, error) {
	if _iODataQueueEnqueue == nil {
		return *new(kernel.IOReturn), symbolCallError("IODataQueueEnqueue", "10.5", _iODataQueueEnqueueErr)
	}
	return _iODataQueueEnqueue(dataQueue, data, dataSize), nil
}

// IODataQueueEnqueue enqueues a new entry on the queue.
//
// See: https://developer.apple.com/documentation/iokit/1514482-iodataqueueenqueue
func IODataQueueEnqueue(dataQueue *IODataQueueMemory, data unsafe.Pointer, dataSize uint32) kernel.IOReturn {
	result, callErr := tryIODataQueueEnqueue(dataQueue, data, dataSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODataQueuePeek func(dataQueue *IODataQueueMemory) IODataQueueEntry
var _iODataQueuePeekErr error

func tryIODataQueuePeek(dataQueue *IODataQueueMemory) (IODataQueueEntry, error) {
	if _iODataQueuePeek == nil {
		return *new(IODataQueueEntry), symbolCallError("IODataQueuePeek", "10.0", _iODataQueuePeekErr)
	}
	return _iODataQueuePeek(dataQueue), nil
}

// IODataQueuePeek used to peek at the next entry on the queue.
//
// See: https://developer.apple.com/documentation/iokit/1514649-iodataqueuepeek
func IODataQueuePeek(dataQueue *IODataQueueMemory) IODataQueueEntry {
	result, callErr := tryIODataQueuePeek(dataQueue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODataQueueSetNotificationPort func(dataQueue *IODataQueueMemory, notifyPort uint32) kernel.IOReturn
var _iODataQueueSetNotificationPortErr error

func tryIODataQueueSetNotificationPort(dataQueue *IODataQueueMemory, notifyPort uint32) (kernel.IOReturn, error) {
	if _iODataQueueSetNotificationPort == nil {
		return *new(kernel.IOReturn), symbolCallError("IODataQueueSetNotificationPort", "10.5", _iODataQueueSetNotificationPortErr)
	}
	return _iODataQueueSetNotificationPort(dataQueue, notifyPort), nil
}

// IODataQueueSetNotificationPort creates a simple mach message targeting the mach port specified in port.
//
// See: https://developer.apple.com/documentation/iokit/1514301-iodataqueuesetnotificationport
func IODataQueueSetNotificationPort(dataQueue *IODataQueueMemory, notifyPort uint32) kernel.IOReturn {
	result, callErr := tryIODataQueueSetNotificationPort(dataQueue, notifyPort)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODataQueueWaitForAvailableData func(dataQueue *IODataQueueMemory, notificationPort uint32) kernel.IOReturn
var _iODataQueueWaitForAvailableDataErr error

func tryIODataQueueWaitForAvailableData(dataQueue *IODataQueueMemory, notificationPort uint32) (kernel.IOReturn, error) {
	if _iODataQueueWaitForAvailableData == nil {
		return *new(kernel.IOReturn), symbolCallError("IODataQueueWaitForAvailableData", "10.0", _iODataQueueWaitForAvailableDataErr)
	}
	return _iODataQueueWaitForAvailableData(dataQueue, notificationPort), nil
}

// IODataQueueWaitForAvailableData wait for an incoming dataAvailable message on the given notifyPort.
//
// See: https://developer.apple.com/documentation/iokit/1514696-iodataqueuewaitforavailabledata
func IODataQueueWaitForAvailableData(dataQueue *IODataQueueMemory, notificationPort uint32) kernel.IOReturn {
	result, callErr := tryIODataQueueWaitForAvailableData(dataQueue, notificationPort)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODeregisterApp func(notifier *uint32) kernel.IOReturn
var _iODeregisterAppErr error

func tryIODeregisterApp(notifier *uint32) (kernel.IOReturn, error) {
	if _iODeregisterApp == nil {
		return *new(kernel.IOReturn), symbolCallError("IODeregisterApp", "10.0", _iODeregisterAppErr)
	}
	return _iODeregisterApp(notifier), nil
}

// IODeregisterApp disconnects the caller from an IOService after receiving power state change notifications from the IOService.
//
// See: https://developer.apple.com/documentation/iokit/1557129-ioderegisterapp
func IODeregisterApp(notifier *uint32) kernel.IOReturn {
	result, callErr := tryIODeregisterApp(notifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODeregisterForSystemPower func(notifier *uint32) kernel.IOReturn
var _iODeregisterForSystemPowerErr error

func tryIODeregisterForSystemPower(notifier *uint32) (kernel.IOReturn, error) {
	if _iODeregisterForSystemPower == nil {
		return *new(kernel.IOReturn), symbolCallError("IODeregisterForSystemPower", "10.0", _iODeregisterForSystemPowerErr)
	}
	return _iODeregisterForSystemPower(notifier), nil
}

// IODeregisterForSystemPower disconnects the caller from the Root Power Domain IOService after receiving system power state change notifications.
//
// See: https://developer.apple.com/documentation/iokit/1557132-ioderegisterforsystempower
func IODeregisterForSystemPower(notifier *uint32) kernel.IOReturn {
	result, callErr := tryIODeregisterForSystemPower(notifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODestroyPlugInInterface func(interface_ *IOCFPlugInInterface) int32
var _iODestroyPlugInInterfaceErr error

func tryIODestroyPlugInInterface(interface_ *IOCFPlugInInterface) (int32, error) {
	if _iODestroyPlugInInterface == nil {
		return 0, symbolCallError("IODestroyPlugInInterface", "10.0", _iODestroyPlugInInterfaceErr)
	}
	return _iODestroyPlugInInterface(interface_), nil
}

// IODestroyPlugInInterface.
//
// See: https://developer.apple.com/documentation/iokit/1412425-iodestroyplugininterface
func IODestroyPlugInInterface(interface_ *IOCFPlugInInterface) int32 {
	result, callErr := tryIODestroyPlugInInterface(interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODispatchCalloutFromMessage func(unused unsafe.Pointer, msg unsafe.Pointer, reference unsafe.Pointer)
var _iODispatchCalloutFromMessageErr error

func tryIODispatchCalloutFromMessage(unused unsafe.Pointer, msg unsafe.Pointer, reference unsafe.Pointer) error {
	if _iODispatchCalloutFromMessage == nil {
		return symbolCallError("IODispatchCalloutFromMessage", "10.0", _iODispatchCalloutFromMessageErr)
	}
	_iODispatchCalloutFromMessage(unused, msg, reference)
	return nil
}

// IODispatchCalloutFromMessage dispatches callback notifications from a mach message.
//
// See: https://developer.apple.com/documentation/iokit/1514775-iodispatchcalloutfrommessage
func IODispatchCalloutFromMessage(unused unsafe.Pointer, msg unsafe.Pointer, reference unsafe.Pointer) {
	if callErr := tryIODispatchCalloutFromMessage(unused, msg, reference); callErr != nil {
		panic(callErr)
	}
}

var _iODisplayCommitParameters func(service uint32, options uint32) kernel.IOReturn
var _iODisplayCommitParametersErr error

func tryIODisplayCommitParameters(service uint32, options uint32) (kernel.IOReturn, error) {
	if _iODisplayCommitParameters == nil {
		return *new(kernel.IOReturn), symbolCallError("IODisplayCommitParameters", "10.0", _iODisplayCommitParametersErr)
	}
	return _iODisplayCommitParameters(service, options), nil
}

// IODisplayCommitParameters.
//
// See: https://developer.apple.com/documentation/iokit/1574906-iodisplaycommitparameters
func IODisplayCommitParameters(service uint32, options uint32) kernel.IOReturn {
	result, callErr := tryIODisplayCommitParameters(service, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODisplayCopyFloatParameters func(service uint32, options uint32, params *corefoundation.CFDictionaryRef) kernel.IOReturn
var _iODisplayCopyFloatParametersErr error

func tryIODisplayCopyFloatParameters(service uint32, options uint32, params *corefoundation.CFDictionaryRef) (kernel.IOReturn, error) {
	if _iODisplayCopyFloatParameters == nil {
		return *new(kernel.IOReturn), symbolCallError("IODisplayCopyFloatParameters", "10.0", _iODisplayCopyFloatParametersErr)
	}
	return _iODisplayCopyFloatParameters(service, options, params), nil
}

// IODisplayCopyFloatParameters.
//
// See: https://developer.apple.com/documentation/iokit/1574891-iodisplaycopyfloatparameters
func IODisplayCopyFloatParameters(service uint32, options uint32, params *corefoundation.CFDictionaryRef) kernel.IOReturn {
	result, callErr := tryIODisplayCopyFloatParameters(service, options, params)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODisplayCopyParameters func(service uint32, options uint32, params *corefoundation.CFDictionaryRef) kernel.IOReturn
var _iODisplayCopyParametersErr error

func tryIODisplayCopyParameters(service uint32, options uint32, params *corefoundation.CFDictionaryRef) (kernel.IOReturn, error) {
	if _iODisplayCopyParameters == nil {
		return *new(kernel.IOReturn), symbolCallError("IODisplayCopyParameters", "10.0", _iODisplayCopyParametersErr)
	}
	return _iODisplayCopyParameters(service, options, params), nil
}

// IODisplayCopyParameters.
//
// See: https://developer.apple.com/documentation/iokit/1574865-iodisplaycopyparameters
func IODisplayCopyParameters(service uint32, options uint32, params *corefoundation.CFDictionaryRef) kernel.IOReturn {
	result, callErr := tryIODisplayCopyParameters(service, options, params)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODisplayCreateInfoDictionary func(framebuffer uint32, options uint32) corefoundation.CFDictionaryRef
var _iODisplayCreateInfoDictionaryErr error

func tryIODisplayCreateInfoDictionary(framebuffer uint32, options uint32) (corefoundation.CFDictionaryRef, error) {
	if _iODisplayCreateInfoDictionary == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("IODisplayCreateInfoDictionary", "10.0", _iODisplayCreateInfoDictionaryErr)
	}
	return _iODisplayCreateInfoDictionary(framebuffer, options), nil
}

// IODisplayCreateInfoDictionary create a CFDictionary with information about display hardware.
//
// See: https://developer.apple.com/documentation/iokit/1574917-iodisplaycreateinfodictionary
func IODisplayCreateInfoDictionary(framebuffer uint32, options uint32) corefoundation.CFDictionaryRef {
	result, callErr := tryIODisplayCreateInfoDictionary(framebuffer, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODisplayForFramebuffer func(framebuffer uint32, options uint32) uint32
var _iODisplayForFramebufferErr error

func tryIODisplayForFramebuffer(framebuffer uint32, options uint32) (uint32, error) {
	if _iODisplayForFramebuffer == nil {
		return 0, symbolCallError("IODisplayForFramebuffer", "10.7", _iODisplayForFramebufferErr)
	}
	return _iODisplayForFramebuffer(framebuffer, options), nil
}

// IODisplayForFramebuffer.
//
// See: https://developer.apple.com/documentation/iokit/1574899-iodisplayforframebuffer
func IODisplayForFramebuffer(framebuffer uint32, options uint32) uint32 {
	result, callErr := tryIODisplayForFramebuffer(framebuffer, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODisplayGetFloatParameter func(service uint32, options uint32, parameterName corefoundation.CFStringRef, value *float32) kernel.IOReturn
var _iODisplayGetFloatParameterErr error

func tryIODisplayGetFloatParameter(service uint32, options uint32, parameterName corefoundation.CFStringRef, value []float32) (kernel.IOReturn, error) {
	if _iODisplayGetFloatParameter == nil {
		return *new(kernel.IOReturn), symbolCallError("IODisplayGetFloatParameter", "10.0", _iODisplayGetFloatParameterErr)
	}
	return _iODisplayGetFloatParameter(service, options, parameterName, unsafe.SliceData(value)), nil
}

// IODisplayGetFloatParameter.
//
// See: https://developer.apple.com/documentation/iokit/1574900-iodisplaygetfloatparameter
func IODisplayGetFloatParameter(service uint32, options uint32, parameterName corefoundation.CFStringRef, value []float32) kernel.IOReturn {
	result, callErr := tryIODisplayGetFloatParameter(service, options, parameterName, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODisplayGetIntegerRangeParameter func(service uint32, options uint32, parameterName corefoundation.CFStringRef, value *int32, min *int32, max *int32) kernel.IOReturn
var _iODisplayGetIntegerRangeParameterErr error

func tryIODisplayGetIntegerRangeParameter(service uint32, options uint32, parameterName corefoundation.CFStringRef, value *int32, min *int32, max *int32) (kernel.IOReturn, error) {
	if _iODisplayGetIntegerRangeParameter == nil {
		return *new(kernel.IOReturn), symbolCallError("IODisplayGetIntegerRangeParameter", "10.0", _iODisplayGetIntegerRangeParameterErr)
	}
	return _iODisplayGetIntegerRangeParameter(service, options, parameterName, value, min, max), nil
}

// IODisplayGetIntegerRangeParameter.
//
// See: https://developer.apple.com/documentation/iokit/1574908-iodisplaygetintegerrangeparamete
func IODisplayGetIntegerRangeParameter(service uint32, options uint32, parameterName corefoundation.CFStringRef, value *int32, min *int32, max *int32) kernel.IOReturn {
	result, callErr := tryIODisplayGetIntegerRangeParameter(service, options, parameterName, value, min, max)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODisplayMatchDictionaries func(matching1 corefoundation.CFDictionaryRef, matching2 corefoundation.CFDictionaryRef, options uint32) int32
var _iODisplayMatchDictionariesErr error

func tryIODisplayMatchDictionaries(matching1 corefoundation.CFDictionaryRef, matching2 corefoundation.CFDictionaryRef, options uint32) (int32, error) {
	if _iODisplayMatchDictionaries == nil {
		return 0, symbolCallError("IODisplayMatchDictionaries", "10.0", _iODisplayMatchDictionariesErr)
	}
	return _iODisplayMatchDictionaries(matching1, matching2, options), nil
}

// IODisplayMatchDictionaries match two display information dictionaries to see if they are for the same display.
//
// See: https://developer.apple.com/documentation/iokit/1574879-iodisplaymatchdictionaries
func IODisplayMatchDictionaries(matching1 corefoundation.CFDictionaryRef, matching2 corefoundation.CFDictionaryRef, options uint32) int32 {
	result, callErr := tryIODisplayMatchDictionaries(matching1, matching2, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODisplaySetFloatParameter func(service uint32, options uint32, parameterName corefoundation.CFStringRef, value float32) kernel.IOReturn
var _iODisplaySetFloatParameterErr error

func tryIODisplaySetFloatParameter(service uint32, options uint32, parameterName corefoundation.CFStringRef, value float32) (kernel.IOReturn, error) {
	if _iODisplaySetFloatParameter == nil {
		return *new(kernel.IOReturn), symbolCallError("IODisplaySetFloatParameter", "10.0", _iODisplaySetFloatParameterErr)
	}
	return _iODisplaySetFloatParameter(service, options, parameterName, value), nil
}

// IODisplaySetFloatParameter.
//
// See: https://developer.apple.com/documentation/iokit/1574926-iodisplaysetfloatparameter
func IODisplaySetFloatParameter(service uint32, options uint32, parameterName corefoundation.CFStringRef, value float32) kernel.IOReturn {
	result, callErr := tryIODisplaySetFloatParameter(service, options, parameterName, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODisplaySetIntegerParameter func(service uint32, options uint32, parameterName corefoundation.CFStringRef, value int32) kernel.IOReturn
var _iODisplaySetIntegerParameterErr error

func tryIODisplaySetIntegerParameter(service uint32, options uint32, parameterName corefoundation.CFStringRef, value int32) (kernel.IOReturn, error) {
	if _iODisplaySetIntegerParameter == nil {
		return *new(kernel.IOReturn), symbolCallError("IODisplaySetIntegerParameter", "10.0", _iODisplaySetIntegerParameterErr)
	}
	return _iODisplaySetIntegerParameter(service, options, parameterName, value), nil
}

// IODisplaySetIntegerParameter.
//
// See: https://developer.apple.com/documentation/iokit/1574915-iodisplaysetintegerparameter
func IODisplaySetIntegerParameter(service uint32, options uint32, parameterName corefoundation.CFStringRef, value int32) kernel.IOReturn {
	result, callErr := tryIODisplaySetIntegerParameter(service, options, parameterName, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iODisplaySetParameters func(service uint32, options uint32, params corefoundation.CFDictionaryRef) kernel.IOReturn
var _iODisplaySetParametersErr error

func tryIODisplaySetParameters(service uint32, options uint32, params corefoundation.CFDictionaryRef) (kernel.IOReturn, error) {
	if _iODisplaySetParameters == nil {
		return *new(kernel.IOReturn), symbolCallError("IODisplaySetParameters", "10.0", _iODisplaySetParametersErr)
	}
	return _iODisplaySetParameters(service, options, params), nil
}

// IODisplaySetParameters.
//
// See: https://developer.apple.com/documentation/iokit/1574878-iodisplaysetparameters
func IODisplaySetParameters(service uint32, options uint32, params corefoundation.CFDictionaryRef) kernel.IOReturn {
	result, callErr := tryIODisplaySetParameters(service, options, params)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOFBCopyI2CInterfaceForBus func(framebuffer uint32, bus uint32, interface_ *uint32) kernel.IOReturn
var _iOFBCopyI2CInterfaceForBusErr error

func tryIOFBCopyI2CInterfaceForBus(framebuffer uint32, bus uint32, interface_ *uint32) (kernel.IOReturn, error) {
	if _iOFBCopyI2CInterfaceForBus == nil {
		return *new(kernel.IOReturn), symbolCallError("IOFBCopyI2CInterfaceForBus", "10.3", _iOFBCopyI2CInterfaceForBusErr)
	}
	return _iOFBCopyI2CInterfaceForBus(framebuffer, bus, interface_), nil
}

// IOFBCopyI2CInterfaceForBus returns an instance of an I2C bus interface, associated with an IOFramebuffer instance / bus index pair.
//
// See: https://developer.apple.com/documentation/iokit/1410345-iofbcopyi2cinterfaceforbus
func IOFBCopyI2CInterfaceForBus(framebuffer uint32, bus uint32, interface_ *uint32) kernel.IOReturn {
	result, callErr := tryIOFBCopyI2CInterfaceForBus(framebuffer, bus, interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOFBGetI2CInterfaceCount func(framebuffer uint32, count *uint32) kernel.IOReturn
var _iOFBGetI2CInterfaceCountErr error

func tryIOFBGetI2CInterfaceCount(framebuffer uint32, count *uint32) (kernel.IOReturn, error) {
	if _iOFBGetI2CInterfaceCount == nil {
		return *new(kernel.IOReturn), symbolCallError("IOFBGetI2CInterfaceCount", "10.3", _iOFBGetI2CInterfaceCountErr)
	}
	return _iOFBGetI2CInterfaceCount(framebuffer, count), nil
}

// IOFBGetI2CInterfaceCount returns a count of I2C interfaces available associated with an IOFramebuffer instance.
//
// See: https://developer.apple.com/documentation/iokit/1410333-iofbgeti2cinterfacecount
func IOFBGetI2CInterfaceCount(framebuffer uint32, count *uint32) kernel.IOReturn {
	result, callErr := tryIOFBGetI2CInterfaceCount(framebuffer, count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOFramebufferOpen func(service uint32, owningTask uint32, type_ uint32, connect *uint32) int32
var _iOFramebufferOpenErr error

func tryIOFramebufferOpen(service uint32, owningTask uint32, type_ uint32, connect *uint32) (int32, error) {
	if _iOFramebufferOpen == nil {
		return 0, symbolCallError("IOFramebufferOpen", "10.0", _iOFramebufferOpenErr)
	}
	return _iOFramebufferOpen(service, owningTask, type_, connect), nil
}

// IOFramebufferOpen.
//
// See: https://developer.apple.com/documentation/iokit/1574872-ioframebufferopen
func IOFramebufferOpen(service uint32, owningTask uint32, type_ uint32, connect *uint32) int32 {
	result, callErr := tryIOFramebufferOpen(service, owningTask, type_, connect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOGetSystemLoadAdvisory func() IOSystemLoadAdvisoryLevel
var _iOGetSystemLoadAdvisoryErr error

func tryIOGetSystemLoadAdvisory() (IOSystemLoadAdvisoryLevel, error) {
	if _iOGetSystemLoadAdvisory == nil {
		return *new(IOSystemLoadAdvisoryLevel), symbolCallError("IOGetSystemLoadAdvisory", "10.6", _iOGetSystemLoadAdvisoryErr)
	}
	return _iOGetSystemLoadAdvisory(), nil
}

// IOGetSystemLoadAdvisory returns a hint about whether now would be a good time to perform time-insensitive work.
//
// See: https://developer.apple.com/documentation/iokit/1557110-iogetsystemloadadvisory
func IOGetSystemLoadAdvisory() IOSystemLoadAdvisoryLevel {
	result, callErr := tryIOGetSystemLoadAdvisory()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDAccessCheckAuditToken func(requestType IOHIDRequestType, token *[32]byte) bool
var _iOHIDAccessCheckAuditTokenErr error

func tryIOHIDAccessCheckAuditToken(requestType IOHIDRequestType, token [32]byte) (bool, error) {
	if _iOHIDAccessCheckAuditToken == nil {
		return false, symbolCallError("IOHIDAccessCheckAuditToken", "10.15", _iOHIDAccessCheckAuditTokenErr)
	}
	return _iOHIDAccessCheckAuditToken(requestType, &token), nil
}

// IOHIDAccessCheckAuditToken.
//
// See: https://developer.apple.com/documentation/iokit/4423108-iohidaccesscheckaudittoken
func IOHIDAccessCheckAuditToken(requestType IOHIDRequestType, token [32]byte) bool {
	result, callErr := tryIOHIDAccessCheckAuditToken(requestType, token)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDCheckAccess func(requestType IOHIDRequestType) IOHIDAccessType
var _iOHIDCheckAccessErr error

func tryIOHIDCheckAccess(requestType IOHIDRequestType) (IOHIDAccessType, error) {
	if _iOHIDCheckAccess == nil {
		return *new(IOHIDAccessType), symbolCallError("IOHIDCheckAccess", "10.15", _iOHIDCheckAccessErr)
	}
	return _iOHIDCheckAccess(requestType), nil
}

// IOHIDCheckAccess.
//
// See: https://developer.apple.com/documentation/iokit/3181573-iohidcheckaccess
func IOHIDCheckAccess(requestType IOHIDRequestType) IOHIDAccessType {
	result, callErr := tryIOHIDCheckAccess(requestType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDCopyCFTypeParameter func(handle uint32, key corefoundation.CFStringRef, parameter *corefoundation.CFTypeRef) int32
var _iOHIDCopyCFTypeParameterErr error

func tryIOHIDCopyCFTypeParameter(handle uint32, key corefoundation.CFStringRef, parameter *corefoundation.CFTypeRef) (int32, error) {
	if _iOHIDCopyCFTypeParameter == nil {
		return 0, symbolCallError("IOHIDCopyCFTypeParameter", "10.3", _iOHIDCopyCFTypeParameterErr)
	}
	return _iOHIDCopyCFTypeParameter(handle, key, parameter), nil
}

// IOHIDCopyCFTypeParameter.
//
// See: https://developer.apple.com/documentation/iokit/1555415-iohidcopycftypeparameter
func IOHIDCopyCFTypeParameter(handle uint32, key corefoundation.CFStringRef, parameter *corefoundation.CFTypeRef) int32 {
	result, callErr := tryIOHIDCopyCFTypeParameter(handle, key, parameter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDCreateSharedMemory func(connect uint32, version uint32) int32
var _iOHIDCreateSharedMemoryErr error

func tryIOHIDCreateSharedMemory(connect uint32, version uint32) (int32, error) {
	if _iOHIDCreateSharedMemory == nil {
		return 0, symbolCallError("IOHIDCreateSharedMemory", "10.0", _iOHIDCreateSharedMemoryErr)
	}
	return _iOHIDCreateSharedMemory(connect, version), nil
}

// IOHIDCreateSharedMemory.
//
// See: https://developer.apple.com/documentation/iokit/1555393-iohidcreatesharedmemory
func IOHIDCreateSharedMemory(connect uint32, version uint32) int32 {
	result, callErr := tryIOHIDCreateSharedMemory(connect, version)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceActivate func(device IOHIDDeviceRef)
var _iOHIDDeviceActivateErr error

func tryIOHIDDeviceActivate(device IOHIDDeviceRef) error {
	if _iOHIDDeviceActivate == nil {
		return symbolCallError("IOHIDDeviceActivate", "10.15", _iOHIDDeviceActivateErr)
	}
	_iOHIDDeviceActivate(device)
	return nil
}

// IOHIDDeviceActivate.
//
// See: https://developer.apple.com/documentation/iokit/3042781-iohiddeviceactivate
func IOHIDDeviceActivate(device IOHIDDeviceRef) {
	if callErr := tryIOHIDDeviceActivate(device); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDDeviceCancel func(device IOHIDDeviceRef)
var _iOHIDDeviceCancelErr error

func tryIOHIDDeviceCancel(device IOHIDDeviceRef) error {
	if _iOHIDDeviceCancel == nil {
		return symbolCallError("IOHIDDeviceCancel", "10.15", _iOHIDDeviceCancelErr)
	}
	_iOHIDDeviceCancel(device)
	return nil
}

// IOHIDDeviceCancel.
//
// See: https://developer.apple.com/documentation/iokit/3042782-iohiddevicecancel
func IOHIDDeviceCancel(device IOHIDDeviceRef) {
	if callErr := tryIOHIDDeviceCancel(device); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDDeviceClose func(device IOHIDDeviceRef, options uint32) kernel.IOReturn
var _iOHIDDeviceCloseErr error

func tryIOHIDDeviceClose(device IOHIDDeviceRef, options uint32) (kernel.IOReturn, error) {
	if _iOHIDDeviceClose == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDDeviceClose", "10.5", _iOHIDDeviceCloseErr)
	}
	return _iOHIDDeviceClose(device, options), nil
}

// IOHIDDeviceClose closes communication with a HID device.
//
// See: https://developer.apple.com/documentation/iokit/1588668-iohiddeviceclose
func IOHIDDeviceClose(device IOHIDDeviceRef, options uint32) kernel.IOReturn {
	result, callErr := tryIOHIDDeviceClose(device, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceConformsTo func(device IOHIDDeviceRef, usagePage uint32, usage uint32) bool
var _iOHIDDeviceConformsToErr error

func tryIOHIDDeviceConformsTo(device IOHIDDeviceRef, usagePage uint32, usage uint32) (bool, error) {
	if _iOHIDDeviceConformsTo == nil {
		return false, symbolCallError("IOHIDDeviceConformsTo", "10.5", _iOHIDDeviceConformsToErr)
	}
	return _iOHIDDeviceConformsTo(device, usagePage, usage), nil
}

// IOHIDDeviceConformsTo convenience function that scans the Application Collection elements to see if it conforms to the provided usagePage and usage.
//
// See: https://developer.apple.com/documentation/iokit/1588665-iohiddeviceconformsto
func IOHIDDeviceConformsTo(device IOHIDDeviceRef, usagePage uint32, usage uint32) bool {
	result, callErr := tryIOHIDDeviceConformsTo(device, usagePage, usage)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceCopyMatchingElements func(device IOHIDDeviceRef, matching corefoundation.CFDictionaryRef, options uint32) corefoundation.CFArrayRef
var _iOHIDDeviceCopyMatchingElementsErr error

func tryIOHIDDeviceCopyMatchingElements(device IOHIDDeviceRef, matching corefoundation.CFDictionaryRef, options uint32) (corefoundation.CFArrayRef, error) {
	if _iOHIDDeviceCopyMatchingElements == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("IOHIDDeviceCopyMatchingElements", "10.5", _iOHIDDeviceCopyMatchingElementsErr)
	}
	return _iOHIDDeviceCopyMatchingElements(device, matching, options), nil
}

// IOHIDDeviceCopyMatchingElements obtains HID elements that match the criteria contained in the matching dictionary.
//
// See: https://developer.apple.com/documentation/iokit/1588671-iohiddevicecopymatchingelements
func IOHIDDeviceCopyMatchingElements(device IOHIDDeviceRef, matching corefoundation.CFDictionaryRef, options uint32) corefoundation.CFArrayRef {
	result, callErr := tryIOHIDDeviceCopyMatchingElements(device, matching, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceCopyValueMultiple func(device IOHIDDeviceRef, elements corefoundation.CFArrayRef, pMultiple *corefoundation.CFDictionaryRef) kernel.IOReturn
var _iOHIDDeviceCopyValueMultipleErr error

func tryIOHIDDeviceCopyValueMultiple(device IOHIDDeviceRef, elements corefoundation.CFArrayRef, pMultiple *corefoundation.CFDictionaryRef) (kernel.IOReturn, error) {
	if _iOHIDDeviceCopyValueMultiple == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDDeviceCopyValueMultiple", "10.5", _iOHIDDeviceCopyValueMultipleErr)
	}
	return _iOHIDDeviceCopyValueMultiple(device, elements, pMultiple), nil
}

// IOHIDDeviceCopyValueMultiple copies a values for multiple elements.
//
// See: https://developer.apple.com/documentation/iokit/1588652-iohiddevicecopyvaluemultiple
func IOHIDDeviceCopyValueMultiple(device IOHIDDeviceRef, elements corefoundation.CFArrayRef, pMultiple *corefoundation.CFDictionaryRef) kernel.IOReturn {
	result, callErr := tryIOHIDDeviceCopyValueMultiple(device, elements, pMultiple)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceCopyValueMultipleWithCallback func(device IOHIDDeviceRef, elements corefoundation.CFArrayRef, pMultiple *corefoundation.CFDictionaryRef, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) kernel.IOReturn
var _iOHIDDeviceCopyValueMultipleWithCallbackErr error

func tryIOHIDDeviceCopyValueMultipleWithCallback(device IOHIDDeviceRef, elements corefoundation.CFArrayRef, pMultiple *corefoundation.CFDictionaryRef, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) (kernel.IOReturn, error) {
	if _iOHIDDeviceCopyValueMultipleWithCallback == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDDeviceCopyValueMultipleWithCallback", "10.5", _iOHIDDeviceCopyValueMultipleWithCallbackErr)
	}
	return _iOHIDDeviceCopyValueMultipleWithCallback(device, elements, pMultiple, timeout, callback, context), nil
}

// IOHIDDeviceCopyValueMultipleWithCallback copies a values for multiple elements and returns status via a completion callback.
//
// See: https://developer.apple.com/documentation/iokit/1588655-iohiddevicecopyvaluemultiplewith
func IOHIDDeviceCopyValueMultipleWithCallback(device IOHIDDeviceRef, elements corefoundation.CFArrayRef, pMultiple *corefoundation.CFDictionaryRef, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) kernel.IOReturn {
	result, callErr := tryIOHIDDeviceCopyValueMultipleWithCallback(device, elements, pMultiple, timeout, callback, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceCreate func(allocator corefoundation.CFAllocatorRef, service uint32) IOHIDDeviceRef
var _iOHIDDeviceCreateErr error

func tryIOHIDDeviceCreate(allocator corefoundation.CFAllocatorRef, service uint32) (IOHIDDeviceRef, error) {
	if _iOHIDDeviceCreate == nil {
		return *new(IOHIDDeviceRef), symbolCallError("IOHIDDeviceCreate", "10.5", _iOHIDDeviceCreateErr)
	}
	return _iOHIDDeviceCreate(allocator, service), nil
}

// IOHIDDeviceCreate creates an element from an io_service_t.
//
// See: https://developer.apple.com/documentation/iokit/1588663-iohiddevicecreate
func IOHIDDeviceCreate(allocator corefoundation.CFAllocatorRef, service uint32) IOHIDDeviceRef {
	result, callErr := tryIOHIDDeviceCreate(allocator, service)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceGetProperty func(device IOHIDDeviceRef, key corefoundation.CFStringRef) corefoundation.CFTypeRef
var _iOHIDDeviceGetPropertyErr error

func tryIOHIDDeviceGetProperty(device IOHIDDeviceRef, key corefoundation.CFStringRef) (corefoundation.CFTypeRef, error) {
	if _iOHIDDeviceGetProperty == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("IOHIDDeviceGetProperty", "10.5", _iOHIDDeviceGetPropertyErr)
	}
	return _iOHIDDeviceGetProperty(device, key), nil
}

// IOHIDDeviceGetProperty obtains a property from an IOHIDDevice.
//
// See: https://developer.apple.com/documentation/iokit/1588648-iohiddevicegetproperty
func IOHIDDeviceGetProperty(device IOHIDDeviceRef, key corefoundation.CFStringRef) corefoundation.CFTypeRef {
	result, callErr := tryIOHIDDeviceGetProperty(device, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceGetReport func(device IOHIDDeviceRef, reportType IOHIDReportType, reportID corefoundation.CFIndex, report *byte, pReportLength *int) kernel.IOReturn
var _iOHIDDeviceGetReportErr error

func tryIOHIDDeviceGetReport(device IOHIDDeviceRef, reportType IOHIDReportType, reportID corefoundation.CFIndex, report []byte, pReportLength *int) (kernel.IOReturn, error) {
	if _iOHIDDeviceGetReport == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDDeviceGetReport", "10.5", _iOHIDDeviceGetReportErr)
	}
	return _iOHIDDeviceGetReport(device, reportType, reportID, unsafe.SliceData(report), pReportLength), nil
}

// IOHIDDeviceGetReport obtains a report from the device.
//
// See: https://developer.apple.com/documentation/iokit/1588659-iohiddevicegetreport
func IOHIDDeviceGetReport(device IOHIDDeviceRef, reportType IOHIDReportType, reportID corefoundation.CFIndex, report []byte, pReportLength *int) kernel.IOReturn {
	result, callErr := tryIOHIDDeviceGetReport(device, reportType, reportID, report, pReportLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceGetReportWithCallback func(device IOHIDDeviceRef, reportType IOHIDReportType, reportID corefoundation.CFIndex, report *byte, pReportLength *int, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) kernel.IOReturn
var _iOHIDDeviceGetReportWithCallbackErr error

func tryIOHIDDeviceGetReportWithCallback(device IOHIDDeviceRef, reportType IOHIDReportType, reportID corefoundation.CFIndex, report []byte, pReportLength *int, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) (kernel.IOReturn, error) {
	if _iOHIDDeviceGetReportWithCallback == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDDeviceGetReportWithCallback", "10.5", _iOHIDDeviceGetReportWithCallbackErr)
	}
	return _iOHIDDeviceGetReportWithCallback(device, reportType, reportID, unsafe.SliceData(report), pReportLength, timeout, callback, context), nil
}

// IOHIDDeviceGetReportWithCallback obtains a report from the device.
//
// See: https://developer.apple.com/documentation/iokit/1588662-iohiddevicegetreportwithcallback
func IOHIDDeviceGetReportWithCallback(device IOHIDDeviceRef, reportType IOHIDReportType, reportID corefoundation.CFIndex, report []byte, pReportLength *int, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) kernel.IOReturn {
	result, callErr := tryIOHIDDeviceGetReportWithCallback(device, reportType, reportID, report, pReportLength, timeout, callback, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceGetService func(device IOHIDDeviceRef) uint32
var _iOHIDDeviceGetServiceErr error

func tryIOHIDDeviceGetService(device IOHIDDeviceRef) (uint32, error) {
	if _iOHIDDeviceGetService == nil {
		return 0, symbolCallError("IOHIDDeviceGetService", "10.6", _iOHIDDeviceGetServiceErr)
	}
	return _iOHIDDeviceGetService(device), nil
}

// IOHIDDeviceGetService returns the io_service_t for an IOHIDDevice, if it has one.
//
// See: https://developer.apple.com/documentation/iokit/1588646-iohiddevicegetservice
func IOHIDDeviceGetService(device IOHIDDeviceRef) uint32 {
	result, callErr := tryIOHIDDeviceGetService(device)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceGetTypeID func() uint
var _iOHIDDeviceGetTypeIDErr error

func tryIOHIDDeviceGetTypeID() (uint, error) {
	if _iOHIDDeviceGetTypeID == nil {
		return 0, symbolCallError("IOHIDDeviceGetTypeID", "10.5", _iOHIDDeviceGetTypeIDErr)
	}
	return _iOHIDDeviceGetTypeID(), nil
}

// IOHIDDeviceGetTypeID returns the type identifier of all IOHIDDevice instances.
//
// See: https://developer.apple.com/documentation/iokit/1588664-iohiddevicegettypeid
func IOHIDDeviceGetTypeID() uint {
	result, callErr := tryIOHIDDeviceGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceGetValue func(device IOHIDDeviceRef, element IOHIDElementRef, pValue *IOHIDValueRef) kernel.IOReturn
var _iOHIDDeviceGetValueErr error

func tryIOHIDDeviceGetValue(device IOHIDDeviceRef, element IOHIDElementRef, pValue *IOHIDValueRef) (kernel.IOReturn, error) {
	if _iOHIDDeviceGetValue == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDDeviceGetValue", "10.5", _iOHIDDeviceGetValueErr)
	}
	return _iOHIDDeviceGetValue(device, element, pValue), nil
}

// IOHIDDeviceGetValue gets a value for an element.
//
// See: https://developer.apple.com/documentation/iokit/1588657-iohiddevicegetvalue
func IOHIDDeviceGetValue(device IOHIDDeviceRef, element IOHIDElementRef, pValue *IOHIDValueRef) kernel.IOReturn {
	result, callErr := tryIOHIDDeviceGetValue(device, element, pValue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceGetValueWithCallback func(device IOHIDDeviceRef, element IOHIDElementRef, pValue *IOHIDValueRef, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) kernel.IOReturn
var _iOHIDDeviceGetValueWithCallbackErr error

func tryIOHIDDeviceGetValueWithCallback(device IOHIDDeviceRef, element IOHIDElementRef, pValue *IOHIDValueRef, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) (kernel.IOReturn, error) {
	if _iOHIDDeviceGetValueWithCallback == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDDeviceGetValueWithCallback", "10.5", _iOHIDDeviceGetValueWithCallbackErr)
	}
	return _iOHIDDeviceGetValueWithCallback(device, element, pValue, timeout, callback, context), nil
}

// IOHIDDeviceGetValueWithCallback gets a value for an element and returns status via a completion callback.
//
// See: https://developer.apple.com/documentation/iokit/1588647-iohiddevicegetvaluewithcallback
func IOHIDDeviceGetValueWithCallback(device IOHIDDeviceRef, element IOHIDElementRef, pValue *IOHIDValueRef, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) kernel.IOReturn {
	result, callErr := tryIOHIDDeviceGetValueWithCallback(device, element, pValue, timeout, callback, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceGetValueWithOptions func(device IOHIDDeviceRef, element IOHIDElementRef, pValue *IOHIDValueRef, options uint32) kernel.IOReturn
var _iOHIDDeviceGetValueWithOptionsErr error

func tryIOHIDDeviceGetValueWithOptions(device IOHIDDeviceRef, element IOHIDElementRef, pValue *IOHIDValueRef, options uint32) (kernel.IOReturn, error) {
	if _iOHIDDeviceGetValueWithOptions == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDDeviceGetValueWithOptions", "10.13", _iOHIDDeviceGetValueWithOptionsErr)
	}
	return _iOHIDDeviceGetValueWithOptions(device, element, pValue, options), nil
}

// IOHIDDeviceGetValueWithOptions.
//
// See: https://developer.apple.com/documentation/iokit/2937933-iohiddevicegetvaluewithoptions
func IOHIDDeviceGetValueWithOptions(device IOHIDDeviceRef, element IOHIDElementRef, pValue *IOHIDValueRef, options uint32) kernel.IOReturn {
	result, callErr := tryIOHIDDeviceGetValueWithOptions(device, element, pValue, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceOpen func(device IOHIDDeviceRef, options uint32) kernel.IOReturn
var _iOHIDDeviceOpenErr error

func tryIOHIDDeviceOpen(device IOHIDDeviceRef, options uint32) (kernel.IOReturn, error) {
	if _iOHIDDeviceOpen == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDDeviceOpen", "10.5", _iOHIDDeviceOpenErr)
	}
	return _iOHIDDeviceOpen(device, options), nil
}

// IOHIDDeviceOpen opens a HID device for communication.
//
// See: https://developer.apple.com/documentation/iokit/1588670-iohiddeviceopen
func IOHIDDeviceOpen(device IOHIDDeviceRef, options uint32) kernel.IOReturn {
	result, callErr := tryIOHIDDeviceOpen(device, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceRegisterInputReportCallback func(device IOHIDDeviceRef, report *byte, reportLength corefoundation.CFIndex, callback unsafe.Pointer, context unsafe.Pointer)
var _iOHIDDeviceRegisterInputReportCallbackErr error

func tryIOHIDDeviceRegisterInputReportCallback(device IOHIDDeviceRef, report []byte, reportLength corefoundation.CFIndex, callback unsafe.Pointer, context unsafe.Pointer) error {
	if _iOHIDDeviceRegisterInputReportCallback == nil {
		return symbolCallError("IOHIDDeviceRegisterInputReportCallback", "10.5", _iOHIDDeviceRegisterInputReportCallbackErr)
	}
	_iOHIDDeviceRegisterInputReportCallback(device, unsafe.SliceData(report), reportLength, callback, context)
	return nil
}

// IOHIDDeviceRegisterInputReportCallback registers a callback to be used when an input report is issued by the device.
//
// See: https://developer.apple.com/documentation/iokit/1588666-iohiddeviceregisterinputreportca
func IOHIDDeviceRegisterInputReportCallback(device IOHIDDeviceRef, report []byte, reportLength corefoundation.CFIndex, callback unsafe.Pointer, context unsafe.Pointer) {
	if callErr := tryIOHIDDeviceRegisterInputReportCallback(device, report, reportLength, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDDeviceRegisterInputReportWithTimeStampCallback func(device IOHIDDeviceRef, report *byte, reportLength corefoundation.CFIndex, callback unsafe.Pointer, context unsafe.Pointer)
var _iOHIDDeviceRegisterInputReportWithTimeStampCallbackErr error

func tryIOHIDDeviceRegisterInputReportWithTimeStampCallback(device IOHIDDeviceRef, report []byte, reportLength corefoundation.CFIndex, callback unsafe.Pointer, context unsafe.Pointer) error {
	if _iOHIDDeviceRegisterInputReportWithTimeStampCallback == nil {
		return symbolCallError("IOHIDDeviceRegisterInputReportWithTimeStampCallback", "10.10", _iOHIDDeviceRegisterInputReportWithTimeStampCallbackErr)
	}
	_iOHIDDeviceRegisterInputReportWithTimeStampCallback(device, unsafe.SliceData(report), reportLength, callback, context)
	return nil
}

// IOHIDDeviceRegisterInputReportWithTimeStampCallback.
//
// See: https://developer.apple.com/documentation/iokit/1588649-iohiddeviceregisterinputreportwi
func IOHIDDeviceRegisterInputReportWithTimeStampCallback(device IOHIDDeviceRef, report []byte, reportLength corefoundation.CFIndex, callback unsafe.Pointer, context unsafe.Pointer) {
	if callErr := tryIOHIDDeviceRegisterInputReportWithTimeStampCallback(device, report, reportLength, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDDeviceRegisterInputValueCallback func(device IOHIDDeviceRef, callback unsafe.Pointer, context unsafe.Pointer)
var _iOHIDDeviceRegisterInputValueCallbackErr error

func tryIOHIDDeviceRegisterInputValueCallback(device IOHIDDeviceRef, callback unsafe.Pointer, context unsafe.Pointer) error {
	if _iOHIDDeviceRegisterInputValueCallback == nil {
		return symbolCallError("IOHIDDeviceRegisterInputValueCallback", "10.5", _iOHIDDeviceRegisterInputValueCallbackErr)
	}
	_iOHIDDeviceRegisterInputValueCallback(device, callback, context)
	return nil
}

// IOHIDDeviceRegisterInputValueCallback registers a callback to be used when an input value is issued by the device.
//
// See: https://developer.apple.com/documentation/iokit/1588672-iohiddeviceregisterinputvaluecal
func IOHIDDeviceRegisterInputValueCallback(device IOHIDDeviceRef, callback unsafe.Pointer, context unsafe.Pointer) {
	if callErr := tryIOHIDDeviceRegisterInputValueCallback(device, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDDeviceRegisterRemovalCallback func(device IOHIDDeviceRef, callback unsafe.Pointer, context unsafe.Pointer)
var _iOHIDDeviceRegisterRemovalCallbackErr error

func tryIOHIDDeviceRegisterRemovalCallback(device IOHIDDeviceRef, callback unsafe.Pointer, context unsafe.Pointer) error {
	if _iOHIDDeviceRegisterRemovalCallback == nil {
		return symbolCallError("IOHIDDeviceRegisterRemovalCallback", "10.5", _iOHIDDeviceRegisterRemovalCallbackErr)
	}
	_iOHIDDeviceRegisterRemovalCallback(device, callback, context)
	return nil
}

// IOHIDDeviceRegisterRemovalCallback registers a callback to be used when a IOHIDDevice is removed.
//
// See: https://developer.apple.com/documentation/iokit/1588673-iohiddeviceregisterremovalcallba
func IOHIDDeviceRegisterRemovalCallback(device IOHIDDeviceRef, callback unsafe.Pointer, context unsafe.Pointer) {
	if callErr := tryIOHIDDeviceRegisterRemovalCallback(device, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDDeviceScheduleWithRunLoop func(device IOHIDDeviceRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef)
var _iOHIDDeviceScheduleWithRunLoopErr error

func tryIOHIDDeviceScheduleWithRunLoop(device IOHIDDeviceRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) error {
	if _iOHIDDeviceScheduleWithRunLoop == nil {
		return symbolCallError("IOHIDDeviceScheduleWithRunLoop", "10.5", _iOHIDDeviceScheduleWithRunLoopErr)
	}
	_iOHIDDeviceScheduleWithRunLoop(device, runLoop, runLoopMode)
	return nil
}

// IOHIDDeviceScheduleWithRunLoop schedules HID device with run loop.
//
// See: https://developer.apple.com/documentation/iokit/1588660-iohiddeviceschedulewithrunloop
func IOHIDDeviceScheduleWithRunLoop(device IOHIDDeviceRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) {
	if callErr := tryIOHIDDeviceScheduleWithRunLoop(device, runLoop, runLoopMode); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDDeviceSetCancelHandler func(device IOHIDDeviceRef, handler unsafe.Pointer)
var _iOHIDDeviceSetCancelHandlerErr error

func tryIOHIDDeviceSetCancelHandler(device IOHIDDeviceRef, handler unsafe.Pointer) error {
	if _iOHIDDeviceSetCancelHandler == nil {
		return symbolCallError("IOHIDDeviceSetCancelHandler", "10.15", _iOHIDDeviceSetCancelHandlerErr)
	}
	_iOHIDDeviceSetCancelHandler(device, handler)
	return nil
}

// IOHIDDeviceSetCancelHandler.
//
// See: https://developer.apple.com/documentation/iokit/3042783-iohiddevicesetcancelhandler
func IOHIDDeviceSetCancelHandler(device IOHIDDeviceRef, handler unsafe.Pointer) {
	if callErr := tryIOHIDDeviceSetCancelHandler(device, handler); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDDeviceSetDispatchQueue func(device IOHIDDeviceRef, queue uintptr)
var _iOHIDDeviceSetDispatchQueueErr error

func tryIOHIDDeviceSetDispatchQueue(device IOHIDDeviceRef, queue dispatch.Queue) error {
	if _iOHIDDeviceSetDispatchQueue == nil {
		return symbolCallError("IOHIDDeviceSetDispatchQueue", "10.15", _iOHIDDeviceSetDispatchQueueErr)
	}
	_iOHIDDeviceSetDispatchQueue(device, uintptr(queue.Handle()))
	return nil
}

// IOHIDDeviceSetDispatchQueue.
//
// See: https://developer.apple.com/documentation/iokit/3042784-iohiddevicesetdispatchqueue
func IOHIDDeviceSetDispatchQueue(device IOHIDDeviceRef, queue dispatch.Queue) {
	if callErr := tryIOHIDDeviceSetDispatchQueue(device, queue); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDDeviceSetInputValueMatching func(device IOHIDDeviceRef, matching corefoundation.CFDictionaryRef)
var _iOHIDDeviceSetInputValueMatchingErr error

func tryIOHIDDeviceSetInputValueMatching(device IOHIDDeviceRef, matching corefoundation.CFDictionaryRef) error {
	if _iOHIDDeviceSetInputValueMatching == nil {
		return symbolCallError("IOHIDDeviceSetInputValueMatching", "10.5", _iOHIDDeviceSetInputValueMatchingErr)
	}
	_iOHIDDeviceSetInputValueMatching(device, matching)
	return nil
}

// IOHIDDeviceSetInputValueMatching sets matching criteria for input values received via IOHIDDeviceRegisterInputValueCallback.
//
// See: https://developer.apple.com/documentation/iokit/1588654-iohiddevicesetinputvaluematching
func IOHIDDeviceSetInputValueMatching(device IOHIDDeviceRef, matching corefoundation.CFDictionaryRef) {
	if callErr := tryIOHIDDeviceSetInputValueMatching(device, matching); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDDeviceSetInputValueMatchingMultiple func(device IOHIDDeviceRef, multiple corefoundation.CFArrayRef)
var _iOHIDDeviceSetInputValueMatchingMultipleErr error

func tryIOHIDDeviceSetInputValueMatchingMultiple(device IOHIDDeviceRef, multiple corefoundation.CFArrayRef) error {
	if _iOHIDDeviceSetInputValueMatchingMultiple == nil {
		return symbolCallError("IOHIDDeviceSetInputValueMatchingMultiple", "10.5", _iOHIDDeviceSetInputValueMatchingMultipleErr)
	}
	_iOHIDDeviceSetInputValueMatchingMultiple(device, multiple)
	return nil
}

// IOHIDDeviceSetInputValueMatchingMultiple sets multiple matching criteria for input values received via IOHIDDeviceRegisterInputValueCallback.
//
// See: https://developer.apple.com/documentation/iokit/1588645-iohiddevicesetinputvaluematching
func IOHIDDeviceSetInputValueMatchingMultiple(device IOHIDDeviceRef, multiple corefoundation.CFArrayRef) {
	if callErr := tryIOHIDDeviceSetInputValueMatchingMultiple(device, multiple); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDDeviceSetProperty func(device IOHIDDeviceRef, key corefoundation.CFStringRef, property corefoundation.CFTypeRef) bool
var _iOHIDDeviceSetPropertyErr error

func tryIOHIDDeviceSetProperty(device IOHIDDeviceRef, key corefoundation.CFStringRef, property corefoundation.CFTypeRef) (bool, error) {
	if _iOHIDDeviceSetProperty == nil {
		return false, symbolCallError("IOHIDDeviceSetProperty", "10.5", _iOHIDDeviceSetPropertyErr)
	}
	return _iOHIDDeviceSetProperty(device, key, property), nil
}

// IOHIDDeviceSetProperty sets a property for an IOHIDDevice.
//
// See: https://developer.apple.com/documentation/iokit/1588653-iohiddevicesetproperty
func IOHIDDeviceSetProperty(device IOHIDDeviceRef, key corefoundation.CFStringRef, property corefoundation.CFTypeRef) bool {
	result, callErr := tryIOHIDDeviceSetProperty(device, key, property)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceSetReport func(device IOHIDDeviceRef, reportType IOHIDReportType, reportID corefoundation.CFIndex, report *byte, reportLength corefoundation.CFIndex) kernel.IOReturn
var _iOHIDDeviceSetReportErr error

func tryIOHIDDeviceSetReport(device IOHIDDeviceRef, reportType IOHIDReportType, reportID corefoundation.CFIndex, report []byte, reportLength corefoundation.CFIndex) (kernel.IOReturn, error) {
	if _iOHIDDeviceSetReport == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDDeviceSetReport", "10.5", _iOHIDDeviceSetReportErr)
	}
	return _iOHIDDeviceSetReport(device, reportType, reportID, unsafe.SliceData(report), reportLength), nil
}

// IOHIDDeviceSetReport sends a report to the device.
//
// See: https://developer.apple.com/documentation/iokit/1588656-iohiddevicesetreport
func IOHIDDeviceSetReport(device IOHIDDeviceRef, reportType IOHIDReportType, reportID corefoundation.CFIndex, report []byte, reportLength corefoundation.CFIndex) kernel.IOReturn {
	result, callErr := tryIOHIDDeviceSetReport(device, reportType, reportID, report, reportLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceSetReportWithCallback func(device IOHIDDeviceRef, reportType IOHIDReportType, reportID corefoundation.CFIndex, report *byte, reportLength corefoundation.CFIndex, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) kernel.IOReturn
var _iOHIDDeviceSetReportWithCallbackErr error

func tryIOHIDDeviceSetReportWithCallback(device IOHIDDeviceRef, reportType IOHIDReportType, reportID corefoundation.CFIndex, report []byte, reportLength corefoundation.CFIndex, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) (kernel.IOReturn, error) {
	if _iOHIDDeviceSetReportWithCallback == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDDeviceSetReportWithCallback", "10.5", _iOHIDDeviceSetReportWithCallbackErr)
	}
	return _iOHIDDeviceSetReportWithCallback(device, reportType, reportID, unsafe.SliceData(report), reportLength, timeout, callback, context), nil
}

// IOHIDDeviceSetReportWithCallback sends a report to the device.
//
// See: https://developer.apple.com/documentation/iokit/1588661-iohiddevicesetreportwithcallback
func IOHIDDeviceSetReportWithCallback(device IOHIDDeviceRef, reportType IOHIDReportType, reportID corefoundation.CFIndex, report []byte, reportLength corefoundation.CFIndex, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) kernel.IOReturn {
	result, callErr := tryIOHIDDeviceSetReportWithCallback(device, reportType, reportID, report, reportLength, timeout, callback, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceSetValue func(device IOHIDDeviceRef, element IOHIDElementRef, value IOHIDValueRef) kernel.IOReturn
var _iOHIDDeviceSetValueErr error

func tryIOHIDDeviceSetValue(device IOHIDDeviceRef, element IOHIDElementRef, value IOHIDValueRef) (kernel.IOReturn, error) {
	if _iOHIDDeviceSetValue == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDDeviceSetValue", "10.5", _iOHIDDeviceSetValueErr)
	}
	return _iOHIDDeviceSetValue(device, element, value), nil
}

// IOHIDDeviceSetValue sets a value for an element.
//
// See: https://developer.apple.com/documentation/iokit/1588651-iohiddevicesetvalue
func IOHIDDeviceSetValue(device IOHIDDeviceRef, element IOHIDElementRef, value IOHIDValueRef) kernel.IOReturn {
	result, callErr := tryIOHIDDeviceSetValue(device, element, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceSetValueMultiple func(device IOHIDDeviceRef, multiple corefoundation.CFDictionaryRef) kernel.IOReturn
var _iOHIDDeviceSetValueMultipleErr error

func tryIOHIDDeviceSetValueMultiple(device IOHIDDeviceRef, multiple corefoundation.CFDictionaryRef) (kernel.IOReturn, error) {
	if _iOHIDDeviceSetValueMultiple == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDDeviceSetValueMultiple", "10.5", _iOHIDDeviceSetValueMultipleErr)
	}
	return _iOHIDDeviceSetValueMultiple(device, multiple), nil
}

// IOHIDDeviceSetValueMultiple sets multiple values for multiple elements.
//
// See: https://developer.apple.com/documentation/iokit/1588669-iohiddevicesetvaluemultiple
func IOHIDDeviceSetValueMultiple(device IOHIDDeviceRef, multiple corefoundation.CFDictionaryRef) kernel.IOReturn {
	result, callErr := tryIOHIDDeviceSetValueMultiple(device, multiple)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceSetValueMultipleWithCallback func(device IOHIDDeviceRef, multiple corefoundation.CFDictionaryRef, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) kernel.IOReturn
var _iOHIDDeviceSetValueMultipleWithCallbackErr error

func tryIOHIDDeviceSetValueMultipleWithCallback(device IOHIDDeviceRef, multiple corefoundation.CFDictionaryRef, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) (kernel.IOReturn, error) {
	if _iOHIDDeviceSetValueMultipleWithCallback == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDDeviceSetValueMultipleWithCallback", "10.5", _iOHIDDeviceSetValueMultipleWithCallbackErr)
	}
	return _iOHIDDeviceSetValueMultipleWithCallback(device, multiple, timeout, callback, context), nil
}

// IOHIDDeviceSetValueMultipleWithCallback sets multiple values for multiple elements and returns status via a completion callback.
//
// See: https://developer.apple.com/documentation/iokit/1588658-iohiddevicesetvaluemultiplewithc
func IOHIDDeviceSetValueMultipleWithCallback(device IOHIDDeviceRef, multiple corefoundation.CFDictionaryRef, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) kernel.IOReturn {
	result, callErr := tryIOHIDDeviceSetValueMultipleWithCallback(device, multiple, timeout, callback, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceSetValueWithCallback func(device IOHIDDeviceRef, element IOHIDElementRef, value IOHIDValueRef, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) kernel.IOReturn
var _iOHIDDeviceSetValueWithCallbackErr error

func tryIOHIDDeviceSetValueWithCallback(device IOHIDDeviceRef, element IOHIDElementRef, value IOHIDValueRef, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) (kernel.IOReturn, error) {
	if _iOHIDDeviceSetValueWithCallback == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDDeviceSetValueWithCallback", "10.5", _iOHIDDeviceSetValueWithCallbackErr)
	}
	return _iOHIDDeviceSetValueWithCallback(device, element, value, timeout, callback, context), nil
}

// IOHIDDeviceSetValueWithCallback sets a value for an element and returns status via a completion callback.
//
// See: https://developer.apple.com/documentation/iokit/1588667-iohiddevicesetvaluewithcallback
func IOHIDDeviceSetValueWithCallback(device IOHIDDeviceRef, element IOHIDElementRef, value IOHIDValueRef, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) kernel.IOReturn {
	result, callErr := tryIOHIDDeviceSetValueWithCallback(device, element, value, timeout, callback, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDDeviceUnscheduleFromRunLoop func(device IOHIDDeviceRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef)
var _iOHIDDeviceUnscheduleFromRunLoopErr error

func tryIOHIDDeviceUnscheduleFromRunLoop(device IOHIDDeviceRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) error {
	if _iOHIDDeviceUnscheduleFromRunLoop == nil {
		return symbolCallError("IOHIDDeviceUnscheduleFromRunLoop", "10.5", _iOHIDDeviceUnscheduleFromRunLoopErr)
	}
	_iOHIDDeviceUnscheduleFromRunLoop(device, runLoop, runLoopMode)
	return nil
}

// IOHIDDeviceUnscheduleFromRunLoop unschedules HID device with run loop.
//
// See: https://developer.apple.com/documentation/iokit/1588650-iohiddeviceunschedulefromrunloop
func IOHIDDeviceUnscheduleFromRunLoop(device IOHIDDeviceRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) {
	if callErr := tryIOHIDDeviceUnscheduleFromRunLoop(device, runLoop, runLoopMode); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDElementAttach func(element IOHIDElementRef, toAttach IOHIDElementRef)
var _iOHIDElementAttachErr error

func tryIOHIDElementAttach(element IOHIDElementRef, toAttach IOHIDElementRef) error {
	if _iOHIDElementAttach == nil {
		return symbolCallError("IOHIDElementAttach", "10.5", _iOHIDElementAttachErr)
	}
	_iOHIDElementAttach(element, toAttach)
	return nil
}

// IOHIDElementAttach establish a relationship between one or more elements.
//
// See: https://developer.apple.com/documentation/iokit/1564146-iohidelementattach
func IOHIDElementAttach(element IOHIDElementRef, toAttach IOHIDElementRef) {
	if callErr := tryIOHIDElementAttach(element, toAttach); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDElementCopyAttached func(element IOHIDElementRef) corefoundation.CFArrayRef
var _iOHIDElementCopyAttachedErr error

func tryIOHIDElementCopyAttached(element IOHIDElementRef) (corefoundation.CFArrayRef, error) {
	if _iOHIDElementCopyAttached == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("IOHIDElementCopyAttached", "10.5", _iOHIDElementCopyAttachedErr)
	}
	return _iOHIDElementCopyAttached(element), nil
}

// IOHIDElementCopyAttached obtain attached elements.
//
// See: https://developer.apple.com/documentation/iokit/1564123-iohidelementcopyattached
func IOHIDElementCopyAttached(element IOHIDElementRef) corefoundation.CFArrayRef {
	result, callErr := tryIOHIDElementCopyAttached(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementCreateWithDictionary func(allocator corefoundation.CFAllocatorRef, dictionary corefoundation.CFDictionaryRef) IOHIDElementRef
var _iOHIDElementCreateWithDictionaryErr error

func tryIOHIDElementCreateWithDictionary(allocator corefoundation.CFAllocatorRef, dictionary corefoundation.CFDictionaryRef) (IOHIDElementRef, error) {
	if _iOHIDElementCreateWithDictionary == nil {
		return *new(IOHIDElementRef), symbolCallError("IOHIDElementCreateWithDictionary", "10.5", _iOHIDElementCreateWithDictionaryErr)
	}
	return _iOHIDElementCreateWithDictionary(allocator, dictionary), nil
}

// IOHIDElementCreateWithDictionary creates an element from a dictionary.
//
// See: https://developer.apple.com/documentation/iokit/1564115-iohidelementcreatewithdictionary
func IOHIDElementCreateWithDictionary(allocator corefoundation.CFAllocatorRef, dictionary corefoundation.CFDictionaryRef) IOHIDElementRef {
	result, callErr := tryIOHIDElementCreateWithDictionary(allocator, dictionary)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementDetach func(element IOHIDElementRef, toDetach IOHIDElementRef)
var _iOHIDElementDetachErr error

func tryIOHIDElementDetach(element IOHIDElementRef, toDetach IOHIDElementRef) error {
	if _iOHIDElementDetach == nil {
		return symbolCallError("IOHIDElementDetach", "10.5", _iOHIDElementDetachErr)
	}
	_iOHIDElementDetach(element, toDetach)
	return nil
}

// IOHIDElementDetach remove a relationship between one or more elements.
//
// See: https://developer.apple.com/documentation/iokit/1564116-iohidelementdetach
func IOHIDElementDetach(element IOHIDElementRef, toDetach IOHIDElementRef) {
	if callErr := tryIOHIDElementDetach(element, toDetach); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDElementGetChildren func(element IOHIDElementRef) corefoundation.CFArrayRef
var _iOHIDElementGetChildrenErr error

func tryIOHIDElementGetChildren(element IOHIDElementRef) (corefoundation.CFArrayRef, error) {
	if _iOHIDElementGetChildren == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("IOHIDElementGetChildren", "10.5", _iOHIDElementGetChildrenErr)
	}
	return _iOHIDElementGetChildren(element), nil
}

// IOHIDElementGetChildren returns the children for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564119-iohidelementgetchildren
func IOHIDElementGetChildren(element IOHIDElementRef) corefoundation.CFArrayRef {
	result, callErr := tryIOHIDElementGetChildren(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementGetCollectionType func(element IOHIDElementRef) IOHIDElementCollectionType
var _iOHIDElementGetCollectionTypeErr error

func tryIOHIDElementGetCollectionType(element IOHIDElementRef) (IOHIDElementCollectionType, error) {
	if _iOHIDElementGetCollectionType == nil {
		return *new(IOHIDElementCollectionType), symbolCallError("IOHIDElementGetCollectionType", "10.5", _iOHIDElementGetCollectionTypeErr)
	}
	return _iOHIDElementGetCollectionType(element), nil
}

// IOHIDElementGetCollectionType retrieves the collection type for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564132-iohidelementgetcollectiontype
func IOHIDElementGetCollectionType(element IOHIDElementRef) IOHIDElementCollectionType {
	result, callErr := tryIOHIDElementGetCollectionType(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementGetCookie func(element IOHIDElementRef) kernel.IOHIDElementCookie
var _iOHIDElementGetCookieErr error

func tryIOHIDElementGetCookie(element IOHIDElementRef) (kernel.IOHIDElementCookie, error) {
	if _iOHIDElementGetCookie == nil {
		return *new(kernel.IOHIDElementCookie), symbolCallError("IOHIDElementGetCookie", "10.5", _iOHIDElementGetCookieErr)
	}
	return _iOHIDElementGetCookie(element), nil
}

// IOHIDElementGetCookie retrieves the cookie for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564124-iohidelementgetcookie
func IOHIDElementGetCookie(element IOHIDElementRef) kernel.IOHIDElementCookie {
	result, callErr := tryIOHIDElementGetCookie(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementGetDevice func(element IOHIDElementRef) IOHIDDeviceRef
var _iOHIDElementGetDeviceErr error

func tryIOHIDElementGetDevice(element IOHIDElementRef) (IOHIDDeviceRef, error) {
	if _iOHIDElementGetDevice == nil {
		return *new(IOHIDDeviceRef), symbolCallError("IOHIDElementGetDevice", "10.5", _iOHIDElementGetDeviceErr)
	}
	return _iOHIDElementGetDevice(element), nil
}

// IOHIDElementGetDevice obtain the device associated with the element.
//
// See: https://developer.apple.com/documentation/iokit/1564139-iohidelementgetdevice
func IOHIDElementGetDevice(element IOHIDElementRef) IOHIDDeviceRef {
	result, callErr := tryIOHIDElementGetDevice(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementGetLogicalMax func(element IOHIDElementRef) corefoundation.CFIndex
var _iOHIDElementGetLogicalMaxErr error

func tryIOHIDElementGetLogicalMax(element IOHIDElementRef) (corefoundation.CFIndex, error) {
	if _iOHIDElementGetLogicalMax == nil {
		return *new(corefoundation.CFIndex), symbolCallError("IOHIDElementGetLogicalMax", "10.5", _iOHIDElementGetLogicalMaxErr)
	}
	return _iOHIDElementGetLogicalMax(element), nil
}

// IOHIDElementGetLogicalMax returns the maximum value possible for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564143-iohidelementgetlogicalmax
func IOHIDElementGetLogicalMax(element IOHIDElementRef) corefoundation.CFIndex {
	result, callErr := tryIOHIDElementGetLogicalMax(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementGetLogicalMin func(element IOHIDElementRef) corefoundation.CFIndex
var _iOHIDElementGetLogicalMinErr error

func tryIOHIDElementGetLogicalMin(element IOHIDElementRef) (corefoundation.CFIndex, error) {
	if _iOHIDElementGetLogicalMin == nil {
		return *new(corefoundation.CFIndex), symbolCallError("IOHIDElementGetLogicalMin", "10.5", _iOHIDElementGetLogicalMinErr)
	}
	return _iOHIDElementGetLogicalMin(element), nil
}

// IOHIDElementGetLogicalMin returns the minimum value possible for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564137-iohidelementgetlogicalmin
func IOHIDElementGetLogicalMin(element IOHIDElementRef) corefoundation.CFIndex {
	result, callErr := tryIOHIDElementGetLogicalMin(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementGetName func(element IOHIDElementRef) corefoundation.CFStringRef
var _iOHIDElementGetNameErr error

func tryIOHIDElementGetName(element IOHIDElementRef) (corefoundation.CFStringRef, error) {
	if _iOHIDElementGetName == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("IOHIDElementGetName", "10.5", _iOHIDElementGetNameErr)
	}
	return _iOHIDElementGetName(element), nil
}

// IOHIDElementGetName returns the name for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564117-iohidelementgetname
func IOHIDElementGetName(element IOHIDElementRef) corefoundation.CFStringRef {
	result, callErr := tryIOHIDElementGetName(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementGetParent func(element IOHIDElementRef) IOHIDElementRef
var _iOHIDElementGetParentErr error

func tryIOHIDElementGetParent(element IOHIDElementRef) (IOHIDElementRef, error) {
	if _iOHIDElementGetParent == nil {
		return *new(IOHIDElementRef), symbolCallError("IOHIDElementGetParent", "10.5", _iOHIDElementGetParentErr)
	}
	return _iOHIDElementGetParent(element), nil
}

// IOHIDElementGetParent returns the parent for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564144-iohidelementgetparent
func IOHIDElementGetParent(element IOHIDElementRef) IOHIDElementRef {
	result, callErr := tryIOHIDElementGetParent(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementGetPhysicalMax func(element IOHIDElementRef) corefoundation.CFIndex
var _iOHIDElementGetPhysicalMaxErr error

func tryIOHIDElementGetPhysicalMax(element IOHIDElementRef) (corefoundation.CFIndex, error) {
	if _iOHIDElementGetPhysicalMax == nil {
		return *new(corefoundation.CFIndex), symbolCallError("IOHIDElementGetPhysicalMax", "10.5", _iOHIDElementGetPhysicalMaxErr)
	}
	return _iOHIDElementGetPhysicalMax(element), nil
}

// IOHIDElementGetPhysicalMax returns the scaled maximum value possible for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564134-iohidelementgetphysicalmax
func IOHIDElementGetPhysicalMax(element IOHIDElementRef) corefoundation.CFIndex {
	result, callErr := tryIOHIDElementGetPhysicalMax(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementGetPhysicalMin func(element IOHIDElementRef) corefoundation.CFIndex
var _iOHIDElementGetPhysicalMinErr error

func tryIOHIDElementGetPhysicalMin(element IOHIDElementRef) (corefoundation.CFIndex, error) {
	if _iOHIDElementGetPhysicalMin == nil {
		return *new(corefoundation.CFIndex), symbolCallError("IOHIDElementGetPhysicalMin", "10.5", _iOHIDElementGetPhysicalMinErr)
	}
	return _iOHIDElementGetPhysicalMin(element), nil
}

// IOHIDElementGetPhysicalMin returns the scaled minimum value possible for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564140-iohidelementgetphysicalmin
func IOHIDElementGetPhysicalMin(element IOHIDElementRef) corefoundation.CFIndex {
	result, callErr := tryIOHIDElementGetPhysicalMin(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementGetProperty func(element IOHIDElementRef, key corefoundation.CFStringRef) corefoundation.CFTypeRef
var _iOHIDElementGetPropertyErr error

func tryIOHIDElementGetProperty(element IOHIDElementRef, key corefoundation.CFStringRef) (corefoundation.CFTypeRef, error) {
	if _iOHIDElementGetProperty == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("IOHIDElementGetProperty", "10.5", _iOHIDElementGetPropertyErr)
	}
	return _iOHIDElementGetProperty(element, key), nil
}

// IOHIDElementGetProperty returns the an element property.
//
// See: https://developer.apple.com/documentation/iokit/1564118-iohidelementgetproperty
func IOHIDElementGetProperty(element IOHIDElementRef, key corefoundation.CFStringRef) corefoundation.CFTypeRef {
	result, callErr := tryIOHIDElementGetProperty(element, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementGetReportCount func(element IOHIDElementRef) uint32
var _iOHIDElementGetReportCountErr error

func tryIOHIDElementGetReportCount(element IOHIDElementRef) (uint32, error) {
	if _iOHIDElementGetReportCount == nil {
		return 0, symbolCallError("IOHIDElementGetReportCount", "10.5", _iOHIDElementGetReportCountErr)
	}
	return _iOHIDElementGetReportCount(element), nil
}

// IOHIDElementGetReportCount returns the report count for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564142-iohidelementgetreportcount
func IOHIDElementGetReportCount(element IOHIDElementRef) uint32 {
	result, callErr := tryIOHIDElementGetReportCount(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementGetReportID func(element IOHIDElementRef) uint32
var _iOHIDElementGetReportIDErr error

func tryIOHIDElementGetReportID(element IOHIDElementRef) (uint32, error) {
	if _iOHIDElementGetReportID == nil {
		return 0, symbolCallError("IOHIDElementGetReportID", "10.5", _iOHIDElementGetReportIDErr)
	}
	return _iOHIDElementGetReportID(element), nil
}

// IOHIDElementGetReportID returns the report ID for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564122-iohidelementgetreportid
func IOHIDElementGetReportID(element IOHIDElementRef) uint32 {
	result, callErr := tryIOHIDElementGetReportID(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementGetReportSize func(element IOHIDElementRef) uint32
var _iOHIDElementGetReportSizeErr error

func tryIOHIDElementGetReportSize(element IOHIDElementRef) (uint32, error) {
	if _iOHIDElementGetReportSize == nil {
		return 0, symbolCallError("IOHIDElementGetReportSize", "10.5", _iOHIDElementGetReportSizeErr)
	}
	return _iOHIDElementGetReportSize(element), nil
}

// IOHIDElementGetReportSize returns the report size in bits for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564130-iohidelementgetreportsize
func IOHIDElementGetReportSize(element IOHIDElementRef) uint32 {
	result, callErr := tryIOHIDElementGetReportSize(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementGetType func(element IOHIDElementRef) IOHIDElementType
var _iOHIDElementGetTypeErr error

func tryIOHIDElementGetType(element IOHIDElementRef) (IOHIDElementType, error) {
	if _iOHIDElementGetType == nil {
		return *new(IOHIDElementType), symbolCallError("IOHIDElementGetType", "10.5", _iOHIDElementGetTypeErr)
	}
	return _iOHIDElementGetType(element), nil
}

// IOHIDElementGetType retrieves the type for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564135-iohidelementgettype
func IOHIDElementGetType(element IOHIDElementRef) IOHIDElementType {
	result, callErr := tryIOHIDElementGetType(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementGetTypeID func() uint
var _iOHIDElementGetTypeIDErr error

func tryIOHIDElementGetTypeID() (uint, error) {
	if _iOHIDElementGetTypeID == nil {
		return 0, symbolCallError("IOHIDElementGetTypeID", "10.5", _iOHIDElementGetTypeIDErr)
	}
	return _iOHIDElementGetTypeID(), nil
}

// IOHIDElementGetTypeID returns the type identifier of all IOHIDElement instances.
//
// See: https://developer.apple.com/documentation/iokit/1564120-iohidelementgettypeid
func IOHIDElementGetTypeID() uint {
	result, callErr := tryIOHIDElementGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementGetUnit func(element IOHIDElementRef) uint32
var _iOHIDElementGetUnitErr error

func tryIOHIDElementGetUnit(element IOHIDElementRef) (uint32, error) {
	if _iOHIDElementGetUnit == nil {
		return 0, symbolCallError("IOHIDElementGetUnit", "10.5", _iOHIDElementGetUnitErr)
	}
	return _iOHIDElementGetUnit(element), nil
}

// IOHIDElementGetUnit returns the unit property for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564136-iohidelementgetunit
func IOHIDElementGetUnit(element IOHIDElementRef) uint32 {
	result, callErr := tryIOHIDElementGetUnit(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementGetUnitExponent func(element IOHIDElementRef) uint32
var _iOHIDElementGetUnitExponentErr error

func tryIOHIDElementGetUnitExponent(element IOHIDElementRef) (uint32, error) {
	if _iOHIDElementGetUnitExponent == nil {
		return 0, symbolCallError("IOHIDElementGetUnitExponent", "10.5", _iOHIDElementGetUnitExponentErr)
	}
	return _iOHIDElementGetUnitExponent(element), nil
}

// IOHIDElementGetUnitExponent returns the unit exponenet in base 10 for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564121-iohidelementgetunitexponent
func IOHIDElementGetUnitExponent(element IOHIDElementRef) uint32 {
	result, callErr := tryIOHIDElementGetUnitExponent(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementGetUsage func(element IOHIDElementRef) uint32
var _iOHIDElementGetUsageErr error

func tryIOHIDElementGetUsage(element IOHIDElementRef) (uint32, error) {
	if _iOHIDElementGetUsage == nil {
		return 0, symbolCallError("IOHIDElementGetUsage", "10.5", _iOHIDElementGetUsageErr)
	}
	return _iOHIDElementGetUsage(element), nil
}

// IOHIDElementGetUsage retrieves the usage for an element.
//
// See: https://developer.apple.com/documentation/iokit/1564126-iohidelementgetusage
func IOHIDElementGetUsage(element IOHIDElementRef) uint32 {
	result, callErr := tryIOHIDElementGetUsage(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementGetUsagePage func(element IOHIDElementRef) uint32
var _iOHIDElementGetUsagePageErr error

func tryIOHIDElementGetUsagePage(element IOHIDElementRef) (uint32, error) {
	if _iOHIDElementGetUsagePage == nil {
		return 0, symbolCallError("IOHIDElementGetUsagePage", "10.5", _iOHIDElementGetUsagePageErr)
	}
	return _iOHIDElementGetUsagePage(element), nil
}

// IOHIDElementGetUsagePage retrieves the usage page for an element.
//
// See: https://developer.apple.com/documentation/iokit/1564128-iohidelementgetusagepage
func IOHIDElementGetUsagePage(element IOHIDElementRef) uint32 {
	result, callErr := tryIOHIDElementGetUsagePage(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementHasNullState func(element IOHIDElementRef) bool
var _iOHIDElementHasNullStateErr error

func tryIOHIDElementHasNullState(element IOHIDElementRef) (bool, error) {
	if _iOHIDElementHasNullState == nil {
		return false, symbolCallError("IOHIDElementHasNullState", "10.5", _iOHIDElementHasNullStateErr)
	}
	return _iOHIDElementHasNullState(element), nil
}

// IOHIDElementHasNullState returns the null state property for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564145-iohidelementhasnullstate
func IOHIDElementHasNullState(element IOHIDElementRef) bool {
	result, callErr := tryIOHIDElementHasNullState(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementHasPreferredState func(element IOHIDElementRef) bool
var _iOHIDElementHasPreferredStateErr error

func tryIOHIDElementHasPreferredState(element IOHIDElementRef) (bool, error) {
	if _iOHIDElementHasPreferredState == nil {
		return false, symbolCallError("IOHIDElementHasPreferredState", "10.5", _iOHIDElementHasPreferredStateErr)
	}
	return _iOHIDElementHasPreferredState(element), nil
}

// IOHIDElementHasPreferredState returns the preferred state property for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564141-iohidelementhaspreferredstate
func IOHIDElementHasPreferredState(element IOHIDElementRef) bool {
	result, callErr := tryIOHIDElementHasPreferredState(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementIsArray func(element IOHIDElementRef) bool
var _iOHIDElementIsArrayErr error

func tryIOHIDElementIsArray(element IOHIDElementRef) (bool, error) {
	if _iOHIDElementIsArray == nil {
		return false, symbolCallError("IOHIDElementIsArray", "10.5", _iOHIDElementIsArrayErr)
	}
	return _iOHIDElementIsArray(element), nil
}

// IOHIDElementIsArray returns the array property for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564133-iohidelementisarray
func IOHIDElementIsArray(element IOHIDElementRef) bool {
	result, callErr := tryIOHIDElementIsArray(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementIsNonLinear func(element IOHIDElementRef) bool
var _iOHIDElementIsNonLinearErr error

func tryIOHIDElementIsNonLinear(element IOHIDElementRef) (bool, error) {
	if _iOHIDElementIsNonLinear == nil {
		return false, symbolCallError("IOHIDElementIsNonLinear", "10.5", _iOHIDElementIsNonLinearErr)
	}
	return _iOHIDElementIsNonLinear(element), nil
}

// IOHIDElementIsNonLinear returns the linear property for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564131-iohidelementisnonlinear
func IOHIDElementIsNonLinear(element IOHIDElementRef) bool {
	result, callErr := tryIOHIDElementIsNonLinear(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementIsRelative func(element IOHIDElementRef) bool
var _iOHIDElementIsRelativeErr error

func tryIOHIDElementIsRelative(element IOHIDElementRef) (bool, error) {
	if _iOHIDElementIsRelative == nil {
		return false, symbolCallError("IOHIDElementIsRelative", "10.5", _iOHIDElementIsRelativeErr)
	}
	return _iOHIDElementIsRelative(element), nil
}

// IOHIDElementIsRelative returns the relative property for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564129-iohidelementisrelative
func IOHIDElementIsRelative(element IOHIDElementRef) bool {
	result, callErr := tryIOHIDElementIsRelative(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementIsVirtual func(element IOHIDElementRef) bool
var _iOHIDElementIsVirtualErr error

func tryIOHIDElementIsVirtual(element IOHIDElementRef) (bool, error) {
	if _iOHIDElementIsVirtual == nil {
		return false, symbolCallError("IOHIDElementIsVirtual", "10.5", _iOHIDElementIsVirtualErr)
	}
	return _iOHIDElementIsVirtual(element), nil
}

// IOHIDElementIsVirtual returns the virtual property for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564125-iohidelementisvirtual
func IOHIDElementIsVirtual(element IOHIDElementRef) bool {
	result, callErr := tryIOHIDElementIsVirtual(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementIsWrapping func(element IOHIDElementRef) bool
var _iOHIDElementIsWrappingErr error

func tryIOHIDElementIsWrapping(element IOHIDElementRef) (bool, error) {
	if _iOHIDElementIsWrapping == nil {
		return false, symbolCallError("IOHIDElementIsWrapping", "10.5", _iOHIDElementIsWrappingErr)
	}
	return _iOHIDElementIsWrapping(element), nil
}

// IOHIDElementIsWrapping returns the wrap property for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564127-iohidelementiswrapping
func IOHIDElementIsWrapping(element IOHIDElementRef) bool {
	result, callErr := tryIOHIDElementIsWrapping(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDElementSetProperty func(element IOHIDElementRef, key corefoundation.CFStringRef, property corefoundation.CFTypeRef) bool
var _iOHIDElementSetPropertyErr error

func tryIOHIDElementSetProperty(element IOHIDElementRef, key corefoundation.CFStringRef, property corefoundation.CFTypeRef) (bool, error) {
	if _iOHIDElementSetProperty == nil {
		return false, symbolCallError("IOHIDElementSetProperty", "10.5", _iOHIDElementSetPropertyErr)
	}
	return _iOHIDElementSetProperty(element, key, property), nil
}

// IOHIDElementSetProperty sets an element property.
//
// See: https://developer.apple.com/documentation/iokit/1564138-iohidelementsetproperty
func IOHIDElementSetProperty(element IOHIDElementRef, key corefoundation.CFStringRef, property corefoundation.CFTypeRef) bool {
	result, callErr := tryIOHIDElementSetProperty(element, key, property)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDEventSystemClientCopyProperty func(client IOHIDEventSystemClientRef, key corefoundation.CFStringRef) corefoundation.CFTypeRef
var _iOHIDEventSystemClientCopyPropertyErr error

func tryIOHIDEventSystemClientCopyProperty(client IOHIDEventSystemClientRef, key corefoundation.CFStringRef) (corefoundation.CFTypeRef, error) {
	if _iOHIDEventSystemClientCopyProperty == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("IOHIDEventSystemClientCopyProperty", "10.12", _iOHIDEventSystemClientCopyPropertyErr)
	}
	return _iOHIDEventSystemClientCopyProperty(client, key), nil
}

// IOHIDEventSystemClientCopyProperty.
//
// See: https://developer.apple.com/documentation/iokit/2269513-iohideventsystemclientcopyproper
func IOHIDEventSystemClientCopyProperty(client IOHIDEventSystemClientRef, key corefoundation.CFStringRef) corefoundation.CFTypeRef {
	result, callErr := tryIOHIDEventSystemClientCopyProperty(client, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDEventSystemClientCopyServices func(client IOHIDEventSystemClientRef) corefoundation.CFArrayRef
var _iOHIDEventSystemClientCopyServicesErr error

func tryIOHIDEventSystemClientCopyServices(client IOHIDEventSystemClientRef) (corefoundation.CFArrayRef, error) {
	if _iOHIDEventSystemClientCopyServices == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("IOHIDEventSystemClientCopyServices", "10.12", _iOHIDEventSystemClientCopyServicesErr)
	}
	return _iOHIDEventSystemClientCopyServices(client), nil
}

// IOHIDEventSystemClientCopyServices.
//
// See: https://developer.apple.com/documentation/iokit/2269511-iohideventsystemclientcopyservic
func IOHIDEventSystemClientCopyServices(client IOHIDEventSystemClientRef) corefoundation.CFArrayRef {
	result, callErr := tryIOHIDEventSystemClientCopyServices(client)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDEventSystemClientCreateSimpleClient func(allocator corefoundation.CFAllocatorRef) IOHIDEventSystemClientRef
var _iOHIDEventSystemClientCreateSimpleClientErr error

func tryIOHIDEventSystemClientCreateSimpleClient(allocator corefoundation.CFAllocatorRef) (IOHIDEventSystemClientRef, error) {
	if _iOHIDEventSystemClientCreateSimpleClient == nil {
		return *new(IOHIDEventSystemClientRef), symbolCallError("IOHIDEventSystemClientCreateSimpleClient", "10.12", _iOHIDEventSystemClientCreateSimpleClientErr)
	}
	return _iOHIDEventSystemClientCreateSimpleClient(allocator), nil
}

// IOHIDEventSystemClientCreateSimpleClient.
//
// See: https://developer.apple.com/documentation/iokit/2269514-iohideventsystemclientcreatesimp
func IOHIDEventSystemClientCreateSimpleClient(allocator corefoundation.CFAllocatorRef) IOHIDEventSystemClientRef {
	result, callErr := tryIOHIDEventSystemClientCreateSimpleClient(allocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDEventSystemClientGetTypeID func() uint
var _iOHIDEventSystemClientGetTypeIDErr error

func tryIOHIDEventSystemClientGetTypeID() (uint, error) {
	if _iOHIDEventSystemClientGetTypeID == nil {
		return 0, symbolCallError("IOHIDEventSystemClientGetTypeID", "10.12", _iOHIDEventSystemClientGetTypeIDErr)
	}
	return _iOHIDEventSystemClientGetTypeID(), nil
}

// IOHIDEventSystemClientGetTypeID.
//
// See: https://developer.apple.com/documentation/iokit/2269512-iohideventsystemclientgettypeid
func IOHIDEventSystemClientGetTypeID() uint {
	result, callErr := tryIOHIDEventSystemClientGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDEventSystemClientSetProperty func(client IOHIDEventSystemClientRef, key corefoundation.CFStringRef, property corefoundation.CFTypeRef) bool
var _iOHIDEventSystemClientSetPropertyErr error

func tryIOHIDEventSystemClientSetProperty(client IOHIDEventSystemClientRef, key corefoundation.CFStringRef, property corefoundation.CFTypeRef) (bool, error) {
	if _iOHIDEventSystemClientSetProperty == nil {
		return false, symbolCallError("IOHIDEventSystemClientSetProperty", "10.12", _iOHIDEventSystemClientSetPropertyErr)
	}
	return _iOHIDEventSystemClientSetProperty(client, key, property), nil
}

// IOHIDEventSystemClientSetProperty.
//
// See: https://developer.apple.com/documentation/iokit/2269517-iohideventsystemclientsetpropert
func IOHIDEventSystemClientSetProperty(client IOHIDEventSystemClientRef, key corefoundation.CFStringRef, property corefoundation.CFTypeRef) bool {
	result, callErr := tryIOHIDEventSystemClientSetProperty(client, key, property)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDGetAccelerationWithKey func(handle uint32, key corefoundation.CFStringRef, acceleration *float64) int32
var _iOHIDGetAccelerationWithKeyErr error

func tryIOHIDGetAccelerationWithKey(handle uint32, key corefoundation.CFStringRef, acceleration []float64) (int32, error) {
	if _iOHIDGetAccelerationWithKey == nil {
		return 0, symbolCallError("IOHIDGetAccelerationWithKey", "10.0", _iOHIDGetAccelerationWithKeyErr)
	}
	return _iOHIDGetAccelerationWithKey(handle, key, unsafe.SliceData(acceleration)), nil
}

// IOHIDGetAccelerationWithKey.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555418-iohidgetaccelerationwithkey
func IOHIDGetAccelerationWithKey(handle uint32, key corefoundation.CFStringRef, acceleration []float64) int32 {
	result, callErr := tryIOHIDGetAccelerationWithKey(handle, key, acceleration)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDGetActivityState func(handle uint32, hidActivityIdle *bool) int32
var _iOHIDGetActivityStateErr error

func tryIOHIDGetActivityState(handle uint32, hidActivityIdle *bool) (int32, error) {
	if _iOHIDGetActivityState == nil {
		return 0, symbolCallError("IOHIDGetActivityState", "10.9", _iOHIDGetActivityStateErr)
	}
	return _iOHIDGetActivityState(handle, hidActivityIdle), nil
}

// IOHIDGetActivityState.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/iokit/1555412-iohidgetactivitystate
func IOHIDGetActivityState(handle uint32, hidActivityIdle *bool) int32 {
	result, callErr := tryIOHIDGetActivityState(handle, hidActivityIdle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDGetButtonEventNum func(connect uint32, button NXMouseButton, eventNum *int32) int32
var _iOHIDGetButtonEventNumErr error

func tryIOHIDGetButtonEventNum(connect uint32, button NXMouseButton, eventNum *int32) (int32, error) {
	if _iOHIDGetButtonEventNum == nil {
		return 0, symbolCallError("IOHIDGetButtonEventNum", "10.0", _iOHIDGetButtonEventNumErr)
	}
	return _iOHIDGetButtonEventNum(connect, button, eventNum), nil
}

// IOHIDGetButtonEventNum.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555407-iohidgetbuttoneventnum
func IOHIDGetButtonEventNum(connect uint32, button NXMouseButton, eventNum *int32) int32 {
	result, callErr := tryIOHIDGetButtonEventNum(connect, button, eventNum)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDGetModifierLockState func(handle uint32, selector int32, state *bool) int32
var _iOHIDGetModifierLockStateErr error

func tryIOHIDGetModifierLockState(handle uint32, selector int32, state *bool) (int32, error) {
	if _iOHIDGetModifierLockState == nil {
		return 0, symbolCallError("IOHIDGetModifierLockState", "10.6", _iOHIDGetModifierLockStateErr)
	}
	return _iOHIDGetModifierLockState(handle, selector, state), nil
}

// IOHIDGetModifierLockState.
//
// See: https://developer.apple.com/documentation/iokit/1555400-iohidgetmodifierlockstate
func IOHIDGetModifierLockState(handle uint32, selector int32, state *bool) int32 {
	result, callErr := tryIOHIDGetModifierLockState(handle, selector, state)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDGetMouseAcceleration func(handle uint32, acceleration *float64) int32
var _iOHIDGetMouseAccelerationErr error

func tryIOHIDGetMouseAcceleration(handle uint32, acceleration []float64) (int32, error) {
	if _iOHIDGetMouseAcceleration == nil {
		return 0, symbolCallError("IOHIDGetMouseAcceleration", "10.0", _iOHIDGetMouseAccelerationErr)
	}
	return _iOHIDGetMouseAcceleration(handle, unsafe.SliceData(acceleration)), nil
}

// IOHIDGetMouseAcceleration.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555417-iohidgetmouseacceleration
func IOHIDGetMouseAcceleration(handle uint32, acceleration []float64) int32 {
	result, callErr := tryIOHIDGetMouseAcceleration(handle, acceleration)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDGetMouseButtonMode func(handle uint32, mode *int32) int32
var _iOHIDGetMouseButtonModeErr error

func tryIOHIDGetMouseButtonMode(handle uint32, mode *int32) (int32, error) {
	if _iOHIDGetMouseButtonMode == nil {
		return 0, symbolCallError("IOHIDGetMouseButtonMode", "10.2", _iOHIDGetMouseButtonModeErr)
	}
	return _iOHIDGetMouseButtonMode(handle, mode), nil
}

// IOHIDGetMouseButtonMode.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/iokit/1555413-iohidgetmousebuttonmode
func IOHIDGetMouseButtonMode(handle uint32, mode *int32) int32 {
	result, callErr := tryIOHIDGetMouseButtonMode(handle, mode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDGetParameter func(handle uint32, key corefoundation.CFStringRef, maxSize uint, bytes unsafe.Pointer, actualSize *uint) int32
var _iOHIDGetParameterErr error

func tryIOHIDGetParameter(handle uint32, key corefoundation.CFStringRef, maxSize uint, bytes unsafe.Pointer, actualSize *uint) (int32, error) {
	if _iOHIDGetParameter == nil {
		return 0, symbolCallError("IOHIDGetParameter", "10.0", _iOHIDGetParameterErr)
	}
	return _iOHIDGetParameter(handle, key, maxSize, bytes, actualSize), nil
}

// IOHIDGetParameter.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555405-iohidgetparameter
func IOHIDGetParameter(handle uint32, key corefoundation.CFStringRef, maxSize uint, bytes unsafe.Pointer, actualSize *uint) int32 {
	result, callErr := tryIOHIDGetParameter(handle, key, maxSize, bytes, actualSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDGetScrollAcceleration func(handle uint32, acceleration *float64) int32
var _iOHIDGetScrollAccelerationErr error

func tryIOHIDGetScrollAcceleration(handle uint32, acceleration []float64) (int32, error) {
	if _iOHIDGetScrollAcceleration == nil {
		return 0, symbolCallError("IOHIDGetScrollAcceleration", "10.0", _iOHIDGetScrollAccelerationErr)
	}
	return _iOHIDGetScrollAcceleration(handle, unsafe.SliceData(acceleration)), nil
}

// IOHIDGetScrollAcceleration.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555389-iohidgetscrollacceleration
func IOHIDGetScrollAcceleration(handle uint32, acceleration []float64) int32 {
	result, callErr := tryIOHIDGetScrollAcceleration(handle, acceleration)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDGetStateForSelector func(handle uint32, selector int32, state *uint32) int32
var _iOHIDGetStateForSelectorErr error

func tryIOHIDGetStateForSelector(handle uint32, selector int32, state *uint32) (int32, error) {
	if _iOHIDGetStateForSelector == nil {
		return 0, symbolCallError("IOHIDGetStateForSelector", "10.9", _iOHIDGetStateForSelectorErr)
	}
	return _iOHIDGetStateForSelector(handle, selector, state), nil
}

// IOHIDGetStateForSelector.
//
// See: https://developer.apple.com/documentation/iokit/1555411-iohidgetstateforselector
func IOHIDGetStateForSelector(handle uint32, selector int32, state *uint32) int32 {
	result, callErr := tryIOHIDGetStateForSelector(handle, selector, state)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDManagerActivate func(manager IOHIDManagerRef)
var _iOHIDManagerActivateErr error

func tryIOHIDManagerActivate(manager IOHIDManagerRef) error {
	if _iOHIDManagerActivate == nil {
		return symbolCallError("IOHIDManagerActivate", "10.15", _iOHIDManagerActivateErr)
	}
	_iOHIDManagerActivate(manager)
	return nil
}

// IOHIDManagerActivate.
//
// See: https://developer.apple.com/documentation/iokit/3042786-iohidmanageractivate
func IOHIDManagerActivate(manager IOHIDManagerRef) {
	if callErr := tryIOHIDManagerActivate(manager); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDManagerCancel func(manager IOHIDManagerRef)
var _iOHIDManagerCancelErr error

func tryIOHIDManagerCancel(manager IOHIDManagerRef) error {
	if _iOHIDManagerCancel == nil {
		return symbolCallError("IOHIDManagerCancel", "10.15", _iOHIDManagerCancelErr)
	}
	_iOHIDManagerCancel(manager)
	return nil
}

// IOHIDManagerCancel.
//
// See: https://developer.apple.com/documentation/iokit/3042787-iohidmanagercancel
func IOHIDManagerCancel(manager IOHIDManagerRef) {
	if callErr := tryIOHIDManagerCancel(manager); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDManagerClose func(manager IOHIDManagerRef, options uint32) kernel.IOReturn
var _iOHIDManagerCloseErr error

func tryIOHIDManagerClose(manager IOHIDManagerRef, options uint32) (kernel.IOReturn, error) {
	if _iOHIDManagerClose == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDManagerClose", "10.5", _iOHIDManagerCloseErr)
	}
	return _iOHIDManagerClose(manager, options), nil
}

// IOHIDManagerClose closes the IOHIDManager.
//
// See: https://developer.apple.com/documentation/iokit/1438405-iohidmanagerclose
func IOHIDManagerClose(manager IOHIDManagerRef, options uint32) kernel.IOReturn {
	result, callErr := tryIOHIDManagerClose(manager, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDManagerCopyDevices func(manager IOHIDManagerRef) corefoundation.CFSet
var _iOHIDManagerCopyDevicesErr error

func tryIOHIDManagerCopyDevices(manager IOHIDManagerRef) (corefoundation.CFSet, error) {
	if _iOHIDManagerCopyDevices == nil {
		return *new(corefoundation.CFSet), symbolCallError("IOHIDManagerCopyDevices", "10.5", _iOHIDManagerCopyDevicesErr)
	}
	return _iOHIDManagerCopyDevices(manager), nil
}

// IOHIDManagerCopyDevices obtains currently enumerated devices.
//
// See: https://developer.apple.com/documentation/iokit/1438391-iohidmanagercopydevices
func IOHIDManagerCopyDevices(manager IOHIDManagerRef) corefoundation.CFSet {
	result, callErr := tryIOHIDManagerCopyDevices(manager)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDManagerCreate func(allocator corefoundation.CFAllocatorRef, options uint32) IOHIDManagerRef
var _iOHIDManagerCreateErr error

func tryIOHIDManagerCreate(allocator corefoundation.CFAllocatorRef, options uint32) (IOHIDManagerRef, error) {
	if _iOHIDManagerCreate == nil {
		return *new(IOHIDManagerRef), symbolCallError("IOHIDManagerCreate", "10.5", _iOHIDManagerCreateErr)
	}
	return _iOHIDManagerCreate(allocator, options), nil
}

// IOHIDManagerCreate creates an IOHIDManager object.
//
// See: https://developer.apple.com/documentation/iokit/1438383-iohidmanagercreate
func IOHIDManagerCreate(allocator corefoundation.CFAllocatorRef, options uint32) IOHIDManagerRef {
	result, callErr := tryIOHIDManagerCreate(allocator, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDManagerGetProperty func(manager IOHIDManagerRef, key corefoundation.CFStringRef) corefoundation.CFTypeRef
var _iOHIDManagerGetPropertyErr error

func tryIOHIDManagerGetProperty(manager IOHIDManagerRef, key corefoundation.CFStringRef) (corefoundation.CFTypeRef, error) {
	if _iOHIDManagerGetProperty == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("IOHIDManagerGetProperty", "10.5", _iOHIDManagerGetPropertyErr)
	}
	return _iOHIDManagerGetProperty(manager, key), nil
}

// IOHIDManagerGetProperty obtains a property of an IOHIDManager.
//
// See: https://developer.apple.com/documentation/iokit/1438403-iohidmanagergetproperty
func IOHIDManagerGetProperty(manager IOHIDManagerRef, key corefoundation.CFStringRef) corefoundation.CFTypeRef {
	result, callErr := tryIOHIDManagerGetProperty(manager, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDManagerGetTypeID func() uint
var _iOHIDManagerGetTypeIDErr error

func tryIOHIDManagerGetTypeID() (uint, error) {
	if _iOHIDManagerGetTypeID == nil {
		return 0, symbolCallError("IOHIDManagerGetTypeID", "10.5", _iOHIDManagerGetTypeIDErr)
	}
	return _iOHIDManagerGetTypeID(), nil
}

// IOHIDManagerGetTypeID returns the type identifier of all IOHIDManager instances.
//
// See: https://developer.apple.com/documentation/iokit/1438375-iohidmanagergettypeid
func IOHIDManagerGetTypeID() uint {
	result, callErr := tryIOHIDManagerGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDManagerOpen func(manager IOHIDManagerRef, options uint32) kernel.IOReturn
var _iOHIDManagerOpenErr error

func tryIOHIDManagerOpen(manager IOHIDManagerRef, options uint32) (kernel.IOReturn, error) {
	if _iOHIDManagerOpen == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDManagerOpen", "10.5", _iOHIDManagerOpenErr)
	}
	return _iOHIDManagerOpen(manager, options), nil
}

// IOHIDManagerOpen opens the IOHIDManager.
//
// See: https://developer.apple.com/documentation/iokit/1438369-iohidmanageropen
func IOHIDManagerOpen(manager IOHIDManagerRef, options uint32) kernel.IOReturn {
	result, callErr := tryIOHIDManagerOpen(manager, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDManagerRegisterDeviceMatchingCallback func(manager IOHIDManagerRef, callback unsafe.Pointer, context unsafe.Pointer)
var _iOHIDManagerRegisterDeviceMatchingCallbackErr error

func tryIOHIDManagerRegisterDeviceMatchingCallback(manager IOHIDManagerRef, callback unsafe.Pointer, context unsafe.Pointer) error {
	if _iOHIDManagerRegisterDeviceMatchingCallback == nil {
		return symbolCallError("IOHIDManagerRegisterDeviceMatchingCallback", "10.5", _iOHIDManagerRegisterDeviceMatchingCallbackErr)
	}
	_iOHIDManagerRegisterDeviceMatchingCallback(manager, callback, context)
	return nil
}

// IOHIDManagerRegisterDeviceMatchingCallback registers a callback to be used a device is enumerated.
//
// See: https://developer.apple.com/documentation/iokit/1438399-iohidmanagerregisterdevicematchi
func IOHIDManagerRegisterDeviceMatchingCallback(manager IOHIDManagerRef, callback unsafe.Pointer, context unsafe.Pointer) {
	if callErr := tryIOHIDManagerRegisterDeviceMatchingCallback(manager, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDManagerRegisterDeviceRemovalCallback func(manager IOHIDManagerRef, callback unsafe.Pointer, context unsafe.Pointer)
var _iOHIDManagerRegisterDeviceRemovalCallbackErr error

func tryIOHIDManagerRegisterDeviceRemovalCallback(manager IOHIDManagerRef, callback unsafe.Pointer, context unsafe.Pointer) error {
	if _iOHIDManagerRegisterDeviceRemovalCallback == nil {
		return symbolCallError("IOHIDManagerRegisterDeviceRemovalCallback", "10.5", _iOHIDManagerRegisterDeviceRemovalCallbackErr)
	}
	_iOHIDManagerRegisterDeviceRemovalCallback(manager, callback, context)
	return nil
}

// IOHIDManagerRegisterDeviceRemovalCallback registers a callback to be used when any enumerated device is removed.
//
// See: https://developer.apple.com/documentation/iokit/1438376-iohidmanagerregisterdeviceremova
func IOHIDManagerRegisterDeviceRemovalCallback(manager IOHIDManagerRef, callback unsafe.Pointer, context unsafe.Pointer) {
	if callErr := tryIOHIDManagerRegisterDeviceRemovalCallback(manager, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDManagerRegisterInputReportCallback func(manager IOHIDManagerRef, callback unsafe.Pointer, context unsafe.Pointer)
var _iOHIDManagerRegisterInputReportCallbackErr error

func tryIOHIDManagerRegisterInputReportCallback(manager IOHIDManagerRef, callback unsafe.Pointer, context unsafe.Pointer) error {
	if _iOHIDManagerRegisterInputReportCallback == nil {
		return symbolCallError("IOHIDManagerRegisterInputReportCallback", "10.5", _iOHIDManagerRegisterInputReportCallbackErr)
	}
	_iOHIDManagerRegisterInputReportCallback(manager, callback, context)
	return nil
}

// IOHIDManagerRegisterInputReportCallback registers a callback to be used when an input report is issued by any enumerated device.
//
// See: https://developer.apple.com/documentation/iokit/1438397-iohidmanagerregisterinputreportc
func IOHIDManagerRegisterInputReportCallback(manager IOHIDManagerRef, callback unsafe.Pointer, context unsafe.Pointer) {
	if callErr := tryIOHIDManagerRegisterInputReportCallback(manager, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDManagerRegisterInputReportWithTimeStampCallback func(manager IOHIDManagerRef, callback unsafe.Pointer, context unsafe.Pointer)
var _iOHIDManagerRegisterInputReportWithTimeStampCallbackErr error

func tryIOHIDManagerRegisterInputReportWithTimeStampCallback(manager IOHIDManagerRef, callback unsafe.Pointer, context unsafe.Pointer) error {
	if _iOHIDManagerRegisterInputReportWithTimeStampCallback == nil {
		return symbolCallError("IOHIDManagerRegisterInputReportWithTimeStampCallback", "10.15", _iOHIDManagerRegisterInputReportWithTimeStampCallbackErr)
	}
	_iOHIDManagerRegisterInputReportWithTimeStampCallback(manager, callback, context)
	return nil
}

// IOHIDManagerRegisterInputReportWithTimeStampCallback.
//
// See: https://developer.apple.com/documentation/iokit/3042788-iohidmanagerregisterinputreportw
func IOHIDManagerRegisterInputReportWithTimeStampCallback(manager IOHIDManagerRef, callback unsafe.Pointer, context unsafe.Pointer) {
	if callErr := tryIOHIDManagerRegisterInputReportWithTimeStampCallback(manager, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDManagerRegisterInputValueCallback func(manager IOHIDManagerRef, callback unsafe.Pointer, context unsafe.Pointer)
var _iOHIDManagerRegisterInputValueCallbackErr error

func tryIOHIDManagerRegisterInputValueCallback(manager IOHIDManagerRef, callback unsafe.Pointer, context unsafe.Pointer) error {
	if _iOHIDManagerRegisterInputValueCallback == nil {
		return symbolCallError("IOHIDManagerRegisterInputValueCallback", "10.5", _iOHIDManagerRegisterInputValueCallbackErr)
	}
	_iOHIDManagerRegisterInputValueCallback(manager, callback, context)
	return nil
}

// IOHIDManagerRegisterInputValueCallback registers a callback to be used when an input value is issued by any enumerated device.
//
// See: https://developer.apple.com/documentation/iokit/1438367-iohidmanagerregisterinputvalueca
func IOHIDManagerRegisterInputValueCallback(manager IOHIDManagerRef, callback unsafe.Pointer, context unsafe.Pointer) {
	if callErr := tryIOHIDManagerRegisterInputValueCallback(manager, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDManagerSaveToPropertyDomain func(manager IOHIDManagerRef, applicationID corefoundation.CFStringRef, userName corefoundation.CFStringRef, hostName corefoundation.CFStringRef, options uint32)
var _iOHIDManagerSaveToPropertyDomainErr error

func tryIOHIDManagerSaveToPropertyDomain(manager IOHIDManagerRef, applicationID corefoundation.CFStringRef, userName corefoundation.CFStringRef, hostName corefoundation.CFStringRef, options uint32) error {
	if _iOHIDManagerSaveToPropertyDomain == nil {
		return symbolCallError("IOHIDManagerSaveToPropertyDomain", "10.6", _iOHIDManagerSaveToPropertyDomainErr)
	}
	_iOHIDManagerSaveToPropertyDomain(manager, applicationID, userName, hostName, options)
	return nil
}

// IOHIDManagerSaveToPropertyDomain used to write out the current properties to a specific domain.
//
// See: https://developer.apple.com/documentation/iokit/1438395-iohidmanagersavetopropertydomain
func IOHIDManagerSaveToPropertyDomain(manager IOHIDManagerRef, applicationID corefoundation.CFStringRef, userName corefoundation.CFStringRef, hostName corefoundation.CFStringRef, options uint32) {
	if callErr := tryIOHIDManagerSaveToPropertyDomain(manager, applicationID, userName, hostName, options); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDManagerScheduleWithRunLoop func(manager IOHIDManagerRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef)
var _iOHIDManagerScheduleWithRunLoopErr error

func tryIOHIDManagerScheduleWithRunLoop(manager IOHIDManagerRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) error {
	if _iOHIDManagerScheduleWithRunLoop == nil {
		return symbolCallError("IOHIDManagerScheduleWithRunLoop", "10.5", _iOHIDManagerScheduleWithRunLoopErr)
	}
	_iOHIDManagerScheduleWithRunLoop(manager, runLoop, runLoopMode)
	return nil
}

// IOHIDManagerScheduleWithRunLoop schedules HID manager with run loop.
//
// See: https://developer.apple.com/documentation/iokit/1438409-iohidmanagerschedulewithrunloop
func IOHIDManagerScheduleWithRunLoop(manager IOHIDManagerRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) {
	if callErr := tryIOHIDManagerScheduleWithRunLoop(manager, runLoop, runLoopMode); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDManagerSetCancelHandler func(manager IOHIDManagerRef, handler unsafe.Pointer)
var _iOHIDManagerSetCancelHandlerErr error

func tryIOHIDManagerSetCancelHandler(manager IOHIDManagerRef, handler unsafe.Pointer) error {
	if _iOHIDManagerSetCancelHandler == nil {
		return symbolCallError("IOHIDManagerSetCancelHandler", "10.15", _iOHIDManagerSetCancelHandlerErr)
	}
	_iOHIDManagerSetCancelHandler(manager, handler)
	return nil
}

// IOHIDManagerSetCancelHandler.
//
// See: https://developer.apple.com/documentation/iokit/3042789-iohidmanagersetcancelhandler
func IOHIDManagerSetCancelHandler(manager IOHIDManagerRef, handler unsafe.Pointer) {
	if callErr := tryIOHIDManagerSetCancelHandler(manager, handler); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDManagerSetDeviceMatching func(manager IOHIDManagerRef, matching corefoundation.CFDictionaryRef)
var _iOHIDManagerSetDeviceMatchingErr error

func tryIOHIDManagerSetDeviceMatching(manager IOHIDManagerRef, matching corefoundation.CFDictionaryRef) error {
	if _iOHIDManagerSetDeviceMatching == nil {
		return symbolCallError("IOHIDManagerSetDeviceMatching", "10.5", _iOHIDManagerSetDeviceMatchingErr)
	}
	_iOHIDManagerSetDeviceMatching(manager, matching)
	return nil
}

// IOHIDManagerSetDeviceMatching sets matching criteria for device enumeration.
//
// See: https://developer.apple.com/documentation/iokit/1438371-iohidmanagersetdevicematching
func IOHIDManagerSetDeviceMatching(manager IOHIDManagerRef, matching corefoundation.CFDictionaryRef) {
	if callErr := tryIOHIDManagerSetDeviceMatching(manager, matching); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDManagerSetDeviceMatchingMultiple func(manager IOHIDManagerRef, multiple corefoundation.CFArrayRef)
var _iOHIDManagerSetDeviceMatchingMultipleErr error

func tryIOHIDManagerSetDeviceMatchingMultiple(manager IOHIDManagerRef, multiple corefoundation.CFArrayRef) error {
	if _iOHIDManagerSetDeviceMatchingMultiple == nil {
		return symbolCallError("IOHIDManagerSetDeviceMatchingMultiple", "10.5", _iOHIDManagerSetDeviceMatchingMultipleErr)
	}
	_iOHIDManagerSetDeviceMatchingMultiple(manager, multiple)
	return nil
}

// IOHIDManagerSetDeviceMatchingMultiple sets multiple matching criteria for device enumeration.
//
// See: https://developer.apple.com/documentation/iokit/1438387-iohidmanagersetdevicematchingmul
func IOHIDManagerSetDeviceMatchingMultiple(manager IOHIDManagerRef, multiple corefoundation.CFArrayRef) {
	if callErr := tryIOHIDManagerSetDeviceMatchingMultiple(manager, multiple); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDManagerSetDispatchQueue func(manager IOHIDManagerRef, queue uintptr)
var _iOHIDManagerSetDispatchQueueErr error

func tryIOHIDManagerSetDispatchQueue(manager IOHIDManagerRef, queue dispatch.Queue) error {
	if _iOHIDManagerSetDispatchQueue == nil {
		return symbolCallError("IOHIDManagerSetDispatchQueue", "10.15", _iOHIDManagerSetDispatchQueueErr)
	}
	_iOHIDManagerSetDispatchQueue(manager, uintptr(queue.Handle()))
	return nil
}

// IOHIDManagerSetDispatchQueue.
//
// See: https://developer.apple.com/documentation/iokit/3042790-iohidmanagersetdispatchqueue
func IOHIDManagerSetDispatchQueue(manager IOHIDManagerRef, queue dispatch.Queue) {
	if callErr := tryIOHIDManagerSetDispatchQueue(manager, queue); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDManagerSetInputValueMatching func(manager IOHIDManagerRef, matching corefoundation.CFDictionaryRef)
var _iOHIDManagerSetInputValueMatchingErr error

func tryIOHIDManagerSetInputValueMatching(manager IOHIDManagerRef, matching corefoundation.CFDictionaryRef) error {
	if _iOHIDManagerSetInputValueMatching == nil {
		return symbolCallError("IOHIDManagerSetInputValueMatching", "10.5", _iOHIDManagerSetInputValueMatchingErr)
	}
	_iOHIDManagerSetInputValueMatching(manager, matching)
	return nil
}

// IOHIDManagerSetInputValueMatching sets matching criteria for input values received via IOHIDManagerRegisterInputValueCallback.
//
// See: https://developer.apple.com/documentation/iokit/1438389-iohidmanagersetinputvaluematchin
func IOHIDManagerSetInputValueMatching(manager IOHIDManagerRef, matching corefoundation.CFDictionaryRef) {
	if callErr := tryIOHIDManagerSetInputValueMatching(manager, matching); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDManagerSetInputValueMatchingMultiple func(manager IOHIDManagerRef, multiple corefoundation.CFArrayRef)
var _iOHIDManagerSetInputValueMatchingMultipleErr error

func tryIOHIDManagerSetInputValueMatchingMultiple(manager IOHIDManagerRef, multiple corefoundation.CFArrayRef) error {
	if _iOHIDManagerSetInputValueMatchingMultiple == nil {
		return symbolCallError("IOHIDManagerSetInputValueMatchingMultiple", "10.5", _iOHIDManagerSetInputValueMatchingMultipleErr)
	}
	_iOHIDManagerSetInputValueMatchingMultiple(manager, multiple)
	return nil
}

// IOHIDManagerSetInputValueMatchingMultiple sets multiple matching criteria for input values received via IOHIDManagerRegisterInputValueCallback.
//
// See: https://developer.apple.com/documentation/iokit/1438379-iohidmanagersetinputvaluematchin
func IOHIDManagerSetInputValueMatchingMultiple(manager IOHIDManagerRef, multiple corefoundation.CFArrayRef) {
	if callErr := tryIOHIDManagerSetInputValueMatchingMultiple(manager, multiple); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDManagerSetProperty func(manager IOHIDManagerRef, key corefoundation.CFStringRef, value corefoundation.CFTypeRef) bool
var _iOHIDManagerSetPropertyErr error

func tryIOHIDManagerSetProperty(manager IOHIDManagerRef, key corefoundation.CFStringRef, value corefoundation.CFTypeRef) (bool, error) {
	if _iOHIDManagerSetProperty == nil {
		return false, symbolCallError("IOHIDManagerSetProperty", "10.5", _iOHIDManagerSetPropertyErr)
	}
	return _iOHIDManagerSetProperty(manager, key, value), nil
}

// IOHIDManagerSetProperty sets a property for an IOHIDManager.
//
// See: https://developer.apple.com/documentation/iokit/1438401-iohidmanagersetproperty
func IOHIDManagerSetProperty(manager IOHIDManagerRef, key corefoundation.CFStringRef, value corefoundation.CFTypeRef) bool {
	result, callErr := tryIOHIDManagerSetProperty(manager, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDManagerUnscheduleFromRunLoop func(manager IOHIDManagerRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef)
var _iOHIDManagerUnscheduleFromRunLoopErr error

func tryIOHIDManagerUnscheduleFromRunLoop(manager IOHIDManagerRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) error {
	if _iOHIDManagerUnscheduleFromRunLoop == nil {
		return symbolCallError("IOHIDManagerUnscheduleFromRunLoop", "10.5", _iOHIDManagerUnscheduleFromRunLoopErr)
	}
	_iOHIDManagerUnscheduleFromRunLoop(manager, runLoop, runLoopMode)
	return nil
}

// IOHIDManagerUnscheduleFromRunLoop unschedules HID manager with run loop.
//
// See: https://developer.apple.com/documentation/iokit/1438378-iohidmanagerunschedulefromrunloo
func IOHIDManagerUnscheduleFromRunLoop(manager IOHIDManagerRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) {
	if callErr := tryIOHIDManagerUnscheduleFromRunLoop(manager, runLoop, runLoopMode); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDPostEvent func(connect uint32, eventType uint32, location IOGPoint, eventData *kernel.NXEventData, eventDataVersion uint32, eventFlags uint32, options uint32) int32
var _iOHIDPostEventErr error

func tryIOHIDPostEvent(connect uint32, eventType uint32, location IOGPoint, eventData *kernel.NXEventData, eventDataVersion uint32, eventFlags uint32, options uint32) (int32, error) {
	if _iOHIDPostEvent == nil {
		return 0, symbolCallError("IOHIDPostEvent", "10.0", _iOHIDPostEventErr)
	}
	return _iOHIDPostEvent(connect, eventType, location, eventData, eventDataVersion, eventFlags, options), nil
}

// IOHIDPostEvent.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/iokit/1555406-iohidpostevent
func IOHIDPostEvent(connect uint32, eventType uint32, location IOGPoint, eventData *kernel.NXEventData, eventDataVersion uint32, eventFlags uint32, options uint32) int32 {
	result, callErr := tryIOHIDPostEvent(connect, eventType, location, eventData, eventDataVersion, eventFlags, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDQueueActivate func(queue IOHIDQueueRef)
var _iOHIDQueueActivateErr error

func tryIOHIDQueueActivate(queue IOHIDQueueRef) error {
	if _iOHIDQueueActivate == nil {
		return symbolCallError("IOHIDQueueActivate", "10.15", _iOHIDQueueActivateErr)
	}
	_iOHIDQueueActivate(queue)
	return nil
}

// IOHIDQueueActivate.
//
// See: https://developer.apple.com/documentation/iokit/3042791-iohidqueueactivate
func IOHIDQueueActivate(queue IOHIDQueueRef) {
	if callErr := tryIOHIDQueueActivate(queue); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDQueueAddElement func(queue IOHIDQueueRef, element IOHIDElementRef)
var _iOHIDQueueAddElementErr error

func tryIOHIDQueueAddElement(queue IOHIDQueueRef, element IOHIDElementRef) error {
	if _iOHIDQueueAddElement == nil {
		return symbolCallError("IOHIDQueueAddElement", "10.5", _iOHIDQueueAddElementErr)
	}
	_iOHIDQueueAddElement(queue, element)
	return nil
}

// IOHIDQueueAddElement adds an element to the queue
//
// See: https://developer.apple.com/documentation/iokit/1545835-iohidqueueaddelement
func IOHIDQueueAddElement(queue IOHIDQueueRef, element IOHIDElementRef) {
	if callErr := tryIOHIDQueueAddElement(queue, element); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDQueueCancel func(queue IOHIDQueueRef)
var _iOHIDQueueCancelErr error

func tryIOHIDQueueCancel(queue IOHIDQueueRef) error {
	if _iOHIDQueueCancel == nil {
		return symbolCallError("IOHIDQueueCancel", "10.15", _iOHIDQueueCancelErr)
	}
	_iOHIDQueueCancel(queue)
	return nil
}

// IOHIDQueueCancel.
//
// See: https://developer.apple.com/documentation/iokit/3042792-iohidqueuecancel
func IOHIDQueueCancel(queue IOHIDQueueRef) {
	if callErr := tryIOHIDQueueCancel(queue); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDQueueContainsElement func(queue IOHIDQueueRef, element IOHIDElementRef) bool
var _iOHIDQueueContainsElementErr error

func tryIOHIDQueueContainsElement(queue IOHIDQueueRef, element IOHIDElementRef) (bool, error) {
	if _iOHIDQueueContainsElement == nil {
		return false, symbolCallError("IOHIDQueueContainsElement", "10.5", _iOHIDQueueContainsElementErr)
	}
	return _iOHIDQueueContainsElement(queue, element), nil
}

// IOHIDQueueContainsElement queries the queue to determine if elemement has been added.
//
// See: https://developer.apple.com/documentation/iokit/1545842-iohidqueuecontainselement
func IOHIDQueueContainsElement(queue IOHIDQueueRef, element IOHIDElementRef) bool {
	result, callErr := tryIOHIDQueueContainsElement(queue, element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDQueueCopyNextValue func(queue IOHIDQueueRef) IOHIDValueRef
var _iOHIDQueueCopyNextValueErr error

func tryIOHIDQueueCopyNextValue(queue IOHIDQueueRef) (IOHIDValueRef, error) {
	if _iOHIDQueueCopyNextValue == nil {
		return *new(IOHIDValueRef), symbolCallError("IOHIDQueueCopyNextValue", "10.5", _iOHIDQueueCopyNextValueErr)
	}
	return _iOHIDQueueCopyNextValue(queue), nil
}

// IOHIDQueueCopyNextValue dequeues a retained copy of an element value from the head of an IOHIDQueue.
//
// See: https://developer.apple.com/documentation/iokit/1545844-iohidqueuecopynextvalue
func IOHIDQueueCopyNextValue(queue IOHIDQueueRef) IOHIDValueRef {
	result, callErr := tryIOHIDQueueCopyNextValue(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDQueueCopyNextValueWithTimeout func(queue IOHIDQueueRef, timeout corefoundation.CFTimeInterval) IOHIDValueRef
var _iOHIDQueueCopyNextValueWithTimeoutErr error

func tryIOHIDQueueCopyNextValueWithTimeout(queue IOHIDQueueRef, timeout corefoundation.CFTimeInterval) (IOHIDValueRef, error) {
	if _iOHIDQueueCopyNextValueWithTimeout == nil {
		return *new(IOHIDValueRef), symbolCallError("IOHIDQueueCopyNextValueWithTimeout", "10.5", _iOHIDQueueCopyNextValueWithTimeoutErr)
	}
	return _iOHIDQueueCopyNextValueWithTimeout(queue, timeout), nil
}

// IOHIDQueueCopyNextValueWithTimeout dequeues a retained copy of an element value from the head of an IOHIDQueue.
//
// See: https://developer.apple.com/documentation/iokit/1545832-iohidqueuecopynextvaluewithtimeo
func IOHIDQueueCopyNextValueWithTimeout(queue IOHIDQueueRef, timeout corefoundation.CFTimeInterval) IOHIDValueRef {
	result, callErr := tryIOHIDQueueCopyNextValueWithTimeout(queue, timeout)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDQueueCreate func(allocator corefoundation.CFAllocatorRef, device IOHIDDeviceRef, depth corefoundation.CFIndex, options uint32) IOHIDQueueRef
var _iOHIDQueueCreateErr error

func tryIOHIDQueueCreate(allocator corefoundation.CFAllocatorRef, device IOHIDDeviceRef, depth corefoundation.CFIndex, options uint32) (IOHIDQueueRef, error) {
	if _iOHIDQueueCreate == nil {
		return *new(IOHIDQueueRef), symbolCallError("IOHIDQueueCreate", "10.5", _iOHIDQueueCreateErr)
	}
	return _iOHIDQueueCreate(allocator, device, depth, options), nil
}

// IOHIDQueueCreate creates an IOHIDQueue object for the specified device.
//
// See: https://developer.apple.com/documentation/iokit/1545840-iohidqueuecreate
func IOHIDQueueCreate(allocator corefoundation.CFAllocatorRef, device IOHIDDeviceRef, depth corefoundation.CFIndex, options uint32) IOHIDQueueRef {
	result, callErr := tryIOHIDQueueCreate(allocator, device, depth, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDQueueGetDepth func(queue IOHIDQueueRef) corefoundation.CFIndex
var _iOHIDQueueGetDepthErr error

func tryIOHIDQueueGetDepth(queue IOHIDQueueRef) (corefoundation.CFIndex, error) {
	if _iOHIDQueueGetDepth == nil {
		return *new(corefoundation.CFIndex), symbolCallError("IOHIDQueueGetDepth", "10.5", _iOHIDQueueGetDepthErr)
	}
	return _iOHIDQueueGetDepth(queue), nil
}

// IOHIDQueueGetDepth obtain the depth of the queue.
//
// See: https://developer.apple.com/documentation/iokit/1545833-iohidqueuegetdepth
func IOHIDQueueGetDepth(queue IOHIDQueueRef) corefoundation.CFIndex {
	result, callErr := tryIOHIDQueueGetDepth(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDQueueGetDevice func(queue IOHIDQueueRef) IOHIDDeviceRef
var _iOHIDQueueGetDeviceErr error

func tryIOHIDQueueGetDevice(queue IOHIDQueueRef) (IOHIDDeviceRef, error) {
	if _iOHIDQueueGetDevice == nil {
		return *new(IOHIDDeviceRef), symbolCallError("IOHIDQueueGetDevice", "10.5", _iOHIDQueueGetDeviceErr)
	}
	return _iOHIDQueueGetDevice(queue), nil
}

// IOHIDQueueGetDevice obtain the device associated with the queue.
//
// See: https://developer.apple.com/documentation/iokit/1545839-iohidqueuegetdevice
func IOHIDQueueGetDevice(queue IOHIDQueueRef) IOHIDDeviceRef {
	result, callErr := tryIOHIDQueueGetDevice(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDQueueGetTypeID func() uint
var _iOHIDQueueGetTypeIDErr error

func tryIOHIDQueueGetTypeID() (uint, error) {
	if _iOHIDQueueGetTypeID == nil {
		return 0, symbolCallError("IOHIDQueueGetTypeID", "10.5", _iOHIDQueueGetTypeIDErr)
	}
	return _iOHIDQueueGetTypeID(), nil
}

// IOHIDQueueGetTypeID returns the type identifier of all IOHIDQueue instances.
//
// See: https://developer.apple.com/documentation/iokit/1545836-iohidqueuegettypeid
func IOHIDQueueGetTypeID() uint {
	result, callErr := tryIOHIDQueueGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDQueueRegisterValueAvailableCallback func(queue IOHIDQueueRef, callback unsafe.Pointer, context unsafe.Pointer)
var _iOHIDQueueRegisterValueAvailableCallbackErr error

func tryIOHIDQueueRegisterValueAvailableCallback(queue IOHIDQueueRef, callback unsafe.Pointer, context unsafe.Pointer) error {
	if _iOHIDQueueRegisterValueAvailableCallback == nil {
		return symbolCallError("IOHIDQueueRegisterValueAvailableCallback", "10.5", _iOHIDQueueRegisterValueAvailableCallbackErr)
	}
	_iOHIDQueueRegisterValueAvailableCallback(queue, callback, context)
	return nil
}

// IOHIDQueueRegisterValueAvailableCallback sets callback to be used when the queue transitions to non-empty.
//
// See: https://developer.apple.com/documentation/iokit/1545829-iohidqueueregistervalueavailable
func IOHIDQueueRegisterValueAvailableCallback(queue IOHIDQueueRef, callback unsafe.Pointer, context unsafe.Pointer) {
	if callErr := tryIOHIDQueueRegisterValueAvailableCallback(queue, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDQueueRemoveElement func(queue IOHIDQueueRef, element IOHIDElementRef)
var _iOHIDQueueRemoveElementErr error

func tryIOHIDQueueRemoveElement(queue IOHIDQueueRef, element IOHIDElementRef) error {
	if _iOHIDQueueRemoveElement == nil {
		return symbolCallError("IOHIDQueueRemoveElement", "10.5", _iOHIDQueueRemoveElementErr)
	}
	_iOHIDQueueRemoveElement(queue, element)
	return nil
}

// IOHIDQueueRemoveElement removes an element from the queue
//
// See: https://developer.apple.com/documentation/iokit/1545838-iohidqueueremoveelement
func IOHIDQueueRemoveElement(queue IOHIDQueueRef, element IOHIDElementRef) {
	if callErr := tryIOHIDQueueRemoveElement(queue, element); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDQueueScheduleWithRunLoop func(queue IOHIDQueueRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef)
var _iOHIDQueueScheduleWithRunLoopErr error

func tryIOHIDQueueScheduleWithRunLoop(queue IOHIDQueueRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) error {
	if _iOHIDQueueScheduleWithRunLoop == nil {
		return symbolCallError("IOHIDQueueScheduleWithRunLoop", "10.5", _iOHIDQueueScheduleWithRunLoopErr)
	}
	_iOHIDQueueScheduleWithRunLoop(queue, runLoop, runLoopMode)
	return nil
}

// IOHIDQueueScheduleWithRunLoop schedules queue with run loop.
//
// See: https://developer.apple.com/documentation/iokit/1545841-iohidqueueschedulewithrunloop
func IOHIDQueueScheduleWithRunLoop(queue IOHIDQueueRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) {
	if callErr := tryIOHIDQueueScheduleWithRunLoop(queue, runLoop, runLoopMode); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDQueueSetCancelHandler func(queue IOHIDQueueRef, handler unsafe.Pointer)
var _iOHIDQueueSetCancelHandlerErr error

func tryIOHIDQueueSetCancelHandler(queue IOHIDQueueRef, handler unsafe.Pointer) error {
	if _iOHIDQueueSetCancelHandler == nil {
		return symbolCallError("IOHIDQueueSetCancelHandler", "10.15", _iOHIDQueueSetCancelHandlerErr)
	}
	_iOHIDQueueSetCancelHandler(queue, handler)
	return nil
}

// IOHIDQueueSetCancelHandler.
//
// See: https://developer.apple.com/documentation/iokit/3042793-iohidqueuesetcancelhandler
func IOHIDQueueSetCancelHandler(queue IOHIDQueueRef, handler unsafe.Pointer) {
	if callErr := tryIOHIDQueueSetCancelHandler(queue, handler); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDQueueSetDepth func(queue IOHIDQueueRef, depth corefoundation.CFIndex)
var _iOHIDQueueSetDepthErr error

func tryIOHIDQueueSetDepth(queue IOHIDQueueRef, depth corefoundation.CFIndex) error {
	if _iOHIDQueueSetDepth == nil {
		return symbolCallError("IOHIDQueueSetDepth", "10.5", _iOHIDQueueSetDepthErr)
	}
	_iOHIDQueueSetDepth(queue, depth)
	return nil
}

// IOHIDQueueSetDepth sets the depth of the queue.
//
// See: https://developer.apple.com/documentation/iokit/1545846-iohidqueuesetdepth
func IOHIDQueueSetDepth(queue IOHIDQueueRef, depth corefoundation.CFIndex) {
	if callErr := tryIOHIDQueueSetDepth(queue, depth); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDQueueSetDispatchQueue func(queue IOHIDQueueRef, dispatchQueue uintptr)
var _iOHIDQueueSetDispatchQueueErr error

func tryIOHIDQueueSetDispatchQueue(queue IOHIDQueueRef, dispatchQueue dispatch.Queue) error {
	if _iOHIDQueueSetDispatchQueue == nil {
		return symbolCallError("IOHIDQueueSetDispatchQueue", "10.15", _iOHIDQueueSetDispatchQueueErr)
	}
	_iOHIDQueueSetDispatchQueue(queue, uintptr(dispatchQueue.Handle()))
	return nil
}

// IOHIDQueueSetDispatchQueue.
//
// See: https://developer.apple.com/documentation/iokit/3042794-iohidqueuesetdispatchqueue
func IOHIDQueueSetDispatchQueue(queue IOHIDQueueRef, dispatchQueue dispatch.Queue) {
	if callErr := tryIOHIDQueueSetDispatchQueue(queue, dispatchQueue); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDQueueStart func(queue IOHIDQueueRef)
var _iOHIDQueueStartErr error

func tryIOHIDQueueStart(queue IOHIDQueueRef) error {
	if _iOHIDQueueStart == nil {
		return symbolCallError("IOHIDQueueStart", "10.5", _iOHIDQueueStartErr)
	}
	_iOHIDQueueStart(queue)
	return nil
}

// IOHIDQueueStart starts element value delivery to the queue.
//
// See: https://developer.apple.com/documentation/iokit/1545843-iohidqueuestart
func IOHIDQueueStart(queue IOHIDQueueRef) {
	if callErr := tryIOHIDQueueStart(queue); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDQueueStop func(queue IOHIDQueueRef)
var _iOHIDQueueStopErr error

func tryIOHIDQueueStop(queue IOHIDQueueRef) error {
	if _iOHIDQueueStop == nil {
		return symbolCallError("IOHIDQueueStop", "10.5", _iOHIDQueueStopErr)
	}
	_iOHIDQueueStop(queue)
	return nil
}

// IOHIDQueueStop stops element value delivery to the queue.
//
// See: https://developer.apple.com/documentation/iokit/1545830-iohidqueuestop
func IOHIDQueueStop(queue IOHIDQueueRef) {
	if callErr := tryIOHIDQueueStop(queue); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDQueueUnscheduleFromRunLoop func(queue IOHIDQueueRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef)
var _iOHIDQueueUnscheduleFromRunLoopErr error

func tryIOHIDQueueUnscheduleFromRunLoop(queue IOHIDQueueRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) error {
	if _iOHIDQueueUnscheduleFromRunLoop == nil {
		return symbolCallError("IOHIDQueueUnscheduleFromRunLoop", "10.5", _iOHIDQueueUnscheduleFromRunLoopErr)
	}
	_iOHIDQueueUnscheduleFromRunLoop(queue, runLoop, runLoopMode)
	return nil
}

// IOHIDQueueUnscheduleFromRunLoop unschedules queue with run loop.
//
// See: https://developer.apple.com/documentation/iokit/1545834-iohidqueueunschedulefromrunloop
func IOHIDQueueUnscheduleFromRunLoop(queue IOHIDQueueRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) {
	if callErr := tryIOHIDQueueUnscheduleFromRunLoop(queue, runLoop, runLoopMode); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDRegisterVirtualDisplay func(handle uint32, display_token *uint32) int32
var _iOHIDRegisterVirtualDisplayErr error

func tryIOHIDRegisterVirtualDisplay(handle uint32, display_token *uint32) (int32, error) {
	if _iOHIDRegisterVirtualDisplay == nil {
		return 0, symbolCallError("IOHIDRegisterVirtualDisplay", "10.0", _iOHIDRegisterVirtualDisplayErr)
	}
	return _iOHIDRegisterVirtualDisplay(handle, display_token), nil
}

// IOHIDRegisterVirtualDisplay.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/iokit/1555416-iohidregistervirtualdisplay
func IOHIDRegisterVirtualDisplay(handle uint32, display_token *uint32) int32 {
	result, callErr := tryIOHIDRegisterVirtualDisplay(handle, display_token)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDRequestAccess func(requestType IOHIDRequestType) bool
var _iOHIDRequestAccessErr error

func tryIOHIDRequestAccess(requestType IOHIDRequestType) (bool, error) {
	if _iOHIDRequestAccess == nil {
		return false, symbolCallError("IOHIDRequestAccess", "10.15", _iOHIDRequestAccessErr)
	}
	return _iOHIDRequestAccess(requestType), nil
}

// IOHIDRequestAccess.
//
// See: https://developer.apple.com/documentation/iokit/3181574-iohidrequestaccess
func IOHIDRequestAccess(requestType IOHIDRequestType) bool {
	result, callErr := tryIOHIDRequestAccess(requestType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDServiceClientConformsTo func(service IOHIDServiceClientRef, usagePage uint32, usage uint32) int32
var _iOHIDServiceClientConformsToErr error

func tryIOHIDServiceClientConformsTo(service IOHIDServiceClientRef, usagePage uint32, usage uint32) (int32, error) {
	if _iOHIDServiceClientConformsTo == nil {
		return 0, symbolCallError("IOHIDServiceClientConformsTo", "10.12", _iOHIDServiceClientConformsToErr)
	}
	return _iOHIDServiceClientConformsTo(service, usagePage, usage), nil
}

// IOHIDServiceClientConformsTo.
//
// See: https://developer.apple.com/documentation/iokit/2269428-iohidserviceclientconformsto
func IOHIDServiceClientConformsTo(service IOHIDServiceClientRef, usagePage uint32, usage uint32) int32 {
	result, callErr := tryIOHIDServiceClientConformsTo(service, usagePage, usage)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDServiceClientCopyProperty func(service IOHIDServiceClientRef, key corefoundation.CFStringRef) corefoundation.CFTypeRef
var _iOHIDServiceClientCopyPropertyErr error

func tryIOHIDServiceClientCopyProperty(service IOHIDServiceClientRef, key corefoundation.CFStringRef) (corefoundation.CFTypeRef, error) {
	if _iOHIDServiceClientCopyProperty == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("IOHIDServiceClientCopyProperty", "10.12", _iOHIDServiceClientCopyPropertyErr)
	}
	return _iOHIDServiceClientCopyProperty(service, key), nil
}

// IOHIDServiceClientCopyProperty.
//
// See: https://developer.apple.com/documentation/iokit/2269430-iohidserviceclientcopyproperty
func IOHIDServiceClientCopyProperty(service IOHIDServiceClientRef, key corefoundation.CFStringRef) corefoundation.CFTypeRef {
	result, callErr := tryIOHIDServiceClientCopyProperty(service, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDServiceClientGetRegistryID func(service IOHIDServiceClientRef) corefoundation.CFTypeRef
var _iOHIDServiceClientGetRegistryIDErr error

func tryIOHIDServiceClientGetRegistryID(service IOHIDServiceClientRef) (corefoundation.CFTypeRef, error) {
	if _iOHIDServiceClientGetRegistryID == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("IOHIDServiceClientGetRegistryID", "10.12", _iOHIDServiceClientGetRegistryIDErr)
	}
	return _iOHIDServiceClientGetRegistryID(service), nil
}

// IOHIDServiceClientGetRegistryID.
//
// See: https://developer.apple.com/documentation/iokit/2269426-iohidserviceclientgetregistryid
func IOHIDServiceClientGetRegistryID(service IOHIDServiceClientRef) corefoundation.CFTypeRef {
	result, callErr := tryIOHIDServiceClientGetRegistryID(service)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDServiceClientGetTypeID func() uint
var _iOHIDServiceClientGetTypeIDErr error

func tryIOHIDServiceClientGetTypeID() (uint, error) {
	if _iOHIDServiceClientGetTypeID == nil {
		return 0, symbolCallError("IOHIDServiceClientGetTypeID", "10.12", _iOHIDServiceClientGetTypeIDErr)
	}
	return _iOHIDServiceClientGetTypeID(), nil
}

// IOHIDServiceClientGetTypeID.
//
// See: https://developer.apple.com/documentation/iokit/2269431-iohidserviceclientgettypeid
func IOHIDServiceClientGetTypeID() uint {
	result, callErr := tryIOHIDServiceClientGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDServiceClientSetProperty func(service IOHIDServiceClientRef, key corefoundation.CFStringRef, property corefoundation.CFTypeRef) bool
var _iOHIDServiceClientSetPropertyErr error

func tryIOHIDServiceClientSetProperty(service IOHIDServiceClientRef, key corefoundation.CFStringRef, property corefoundation.CFTypeRef) (bool, error) {
	if _iOHIDServiceClientSetProperty == nil {
		return false, symbolCallError("IOHIDServiceClientSetProperty", "10.12", _iOHIDServiceClientSetPropertyErr)
	}
	return _iOHIDServiceClientSetProperty(service, key, property), nil
}

// IOHIDServiceClientSetProperty.
//
// See: https://developer.apple.com/documentation/iokit/2269429-iohidserviceclientsetproperty
func IOHIDServiceClientSetProperty(service IOHIDServiceClientRef, key corefoundation.CFStringRef, property corefoundation.CFTypeRef) bool {
	result, callErr := tryIOHIDServiceClientSetProperty(service, key, property)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDSetAccelerationWithKey func(handle uint32, key corefoundation.CFStringRef, acceleration float64) int32
var _iOHIDSetAccelerationWithKeyErr error

func tryIOHIDSetAccelerationWithKey(handle uint32, key corefoundation.CFStringRef, acceleration float64) (int32, error) {
	if _iOHIDSetAccelerationWithKey == nil {
		return 0, symbolCallError("IOHIDSetAccelerationWithKey", "10.0", _iOHIDSetAccelerationWithKeyErr)
	}
	return _iOHIDSetAccelerationWithKey(handle, key, acceleration), nil
}

// IOHIDSetAccelerationWithKey.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555398-iohidsetaccelerationwithkey
func IOHIDSetAccelerationWithKey(handle uint32, key corefoundation.CFStringRef, acceleration float64) int32 {
	result, callErr := tryIOHIDSetAccelerationWithKey(handle, key, acceleration)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDSetCFTypeParameter func(handle uint32, key corefoundation.CFStringRef, parameter corefoundation.CFTypeRef) int32
var _iOHIDSetCFTypeParameterErr error

func tryIOHIDSetCFTypeParameter(handle uint32, key corefoundation.CFStringRef, parameter corefoundation.CFTypeRef) (int32, error) {
	if _iOHIDSetCFTypeParameter == nil {
		return 0, symbolCallError("IOHIDSetCFTypeParameter", "10.3", _iOHIDSetCFTypeParameterErr)
	}
	return _iOHIDSetCFTypeParameter(handle, key, parameter), nil
}

// IOHIDSetCFTypeParameter.
//
// See: https://developer.apple.com/documentation/iokit/1555395-iohidsetcftypeparameter
func IOHIDSetCFTypeParameter(handle uint32, key corefoundation.CFStringRef, parameter corefoundation.CFTypeRef) int32 {
	result, callErr := tryIOHIDSetCFTypeParameter(handle, key, parameter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDSetCursorEnable func(connect uint32, enable int32) int32
var _iOHIDSetCursorEnableErr error

func tryIOHIDSetCursorEnable(connect uint32, enable int32) (int32, error) {
	if _iOHIDSetCursorEnable == nil {
		return 0, symbolCallError("IOHIDSetCursorEnable", "10.0", _iOHIDSetCursorEnableErr)
	}
	return _iOHIDSetCursorEnable(connect, enable), nil
}

// IOHIDSetCursorEnable.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/iokit/1555409-iohidsetcursorenable
func IOHIDSetCursorEnable(connect uint32, enable int32) int32 {
	result, callErr := tryIOHIDSetCursorEnable(connect, enable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDSetEventsEnable func(connect uint32, enable int32) int32
var _iOHIDSetEventsEnableErr error

func tryIOHIDSetEventsEnable(connect uint32, enable int32) (int32, error) {
	if _iOHIDSetEventsEnable == nil {
		return 0, symbolCallError("IOHIDSetEventsEnable", "10.0", _iOHIDSetEventsEnableErr)
	}
	return _iOHIDSetEventsEnable(connect, enable), nil
}

// IOHIDSetEventsEnable.
//
// See: https://developer.apple.com/documentation/iokit/1555396-iohidseteventsenable
func IOHIDSetEventsEnable(connect uint32, enable int32) int32 {
	result, callErr := tryIOHIDSetEventsEnable(connect, enable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDSetModifierLockState func(handle uint32, selector int32, state bool) int32
var _iOHIDSetModifierLockStateErr error

func tryIOHIDSetModifierLockState(handle uint32, selector int32, state bool) (int32, error) {
	if _iOHIDSetModifierLockState == nil {
		return 0, symbolCallError("IOHIDSetModifierLockState", "10.6", _iOHIDSetModifierLockStateErr)
	}
	return _iOHIDSetModifierLockState(handle, selector, state), nil
}

// IOHIDSetModifierLockState.
//
// See: https://developer.apple.com/documentation/iokit/1555420-iohidsetmodifierlockstate
func IOHIDSetModifierLockState(handle uint32, selector int32, state bool) int32 {
	result, callErr := tryIOHIDSetModifierLockState(handle, selector, state)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDSetMouseAcceleration func(handle uint32, acceleration float64) int32
var _iOHIDSetMouseAccelerationErr error

func tryIOHIDSetMouseAcceleration(handle uint32, acceleration float64) (int32, error) {
	if _iOHIDSetMouseAcceleration == nil {
		return 0, symbolCallError("IOHIDSetMouseAcceleration", "10.0", _iOHIDSetMouseAccelerationErr)
	}
	return _iOHIDSetMouseAcceleration(handle, acceleration), nil
}

// IOHIDSetMouseAcceleration.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555390-iohidsetmouseacceleration
func IOHIDSetMouseAcceleration(handle uint32, acceleration float64) int32 {
	result, callErr := tryIOHIDSetMouseAcceleration(handle, acceleration)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDSetMouseButtonMode func(handle uint32, mode int32) int32
var _iOHIDSetMouseButtonModeErr error

func tryIOHIDSetMouseButtonMode(handle uint32, mode int32) (int32, error) {
	if _iOHIDSetMouseButtonMode == nil {
		return 0, symbolCallError("IOHIDSetMouseButtonMode", "10.0", _iOHIDSetMouseButtonModeErr)
	}
	return _iOHIDSetMouseButtonMode(handle, mode), nil
}

// IOHIDSetMouseButtonMode.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555399-iohidsetmousebuttonmode
func IOHIDSetMouseButtonMode(handle uint32, mode int32) int32 {
	result, callErr := tryIOHIDSetMouseButtonMode(handle, mode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDSetMouseLocation func(connect uint32, x int32, y int32) int32
var _iOHIDSetMouseLocationErr error

func tryIOHIDSetMouseLocation(connect uint32, x int32, y int32) (int32, error) {
	if _iOHIDSetMouseLocation == nil {
		return 0, symbolCallError("IOHIDSetMouseLocation", "10.0", _iOHIDSetMouseLocationErr)
	}
	return _iOHIDSetMouseLocation(connect, x, y), nil
}

// IOHIDSetMouseLocation.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/iokit/1555402-iohidsetmouselocation
func IOHIDSetMouseLocation(connect uint32, x int32, y int32) int32 {
	result, callErr := tryIOHIDSetMouseLocation(connect, x, y)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDSetParameter func(handle uint32, key corefoundation.CFStringRef, bytes unsafe.Pointer, size uint) int32
var _iOHIDSetParameterErr error

func tryIOHIDSetParameter(handle uint32, key corefoundation.CFStringRef, bytes unsafe.Pointer, size uint) (int32, error) {
	if _iOHIDSetParameter == nil {
		return 0, symbolCallError("IOHIDSetParameter", "10.0", _iOHIDSetParameterErr)
	}
	return _iOHIDSetParameter(handle, key, bytes, size), nil
}

// IOHIDSetParameter.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555394-iohidsetparameter
func IOHIDSetParameter(handle uint32, key corefoundation.CFStringRef, bytes unsafe.Pointer, size uint) int32 {
	result, callErr := tryIOHIDSetParameter(handle, key, bytes, size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDSetScrollAcceleration func(handle uint32, acceleration float64) int32
var _iOHIDSetScrollAccelerationErr error

func tryIOHIDSetScrollAcceleration(handle uint32, acceleration float64) (int32, error) {
	if _iOHIDSetScrollAcceleration == nil {
		return 0, symbolCallError("IOHIDSetScrollAcceleration", "10.0", _iOHIDSetScrollAccelerationErr)
	}
	return _iOHIDSetScrollAcceleration(handle, acceleration), nil
}

// IOHIDSetScrollAcceleration.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555391-iohidsetscrollacceleration
func IOHIDSetScrollAcceleration(handle uint32, acceleration float64) int32 {
	result, callErr := tryIOHIDSetScrollAcceleration(handle, acceleration)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDSetStateForSelector func(handle uint32, selector int32, state uint32) int32
var _iOHIDSetStateForSelectorErr error

func tryIOHIDSetStateForSelector(handle uint32, selector int32, state uint32) (int32, error) {
	if _iOHIDSetStateForSelector == nil {
		return 0, symbolCallError("IOHIDSetStateForSelector", "10.9", _iOHIDSetStateForSelectorErr)
	}
	return _iOHIDSetStateForSelector(handle, selector, state), nil
}

// IOHIDSetStateForSelector.
//
// See: https://developer.apple.com/documentation/iokit/1555397-iohidsetstateforselector
func IOHIDSetStateForSelector(handle uint32, selector int32, state uint32) int32 {
	result, callErr := tryIOHIDSetStateForSelector(handle, selector, state)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDSetVirtualDisplayBounds func(handle uint32, display_token uint32, bounds *IOGBounds) int32
var _iOHIDSetVirtualDisplayBoundsErr error

func tryIOHIDSetVirtualDisplayBounds(handle uint32, display_token uint32, bounds *IOGBounds) (int32, error) {
	if _iOHIDSetVirtualDisplayBounds == nil {
		return 0, symbolCallError("IOHIDSetVirtualDisplayBounds", "10.0", _iOHIDSetVirtualDisplayBoundsErr)
	}
	return _iOHIDSetVirtualDisplayBounds(handle, display_token, bounds), nil
}

// IOHIDSetVirtualDisplayBounds.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/iokit/1555392-iohidsetvirtualdisplaybounds
func IOHIDSetVirtualDisplayBounds(handle uint32, display_token uint32, bounds *IOGBounds) int32 {
	result, callErr := tryIOHIDSetVirtualDisplayBounds(handle, display_token, bounds)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDTransactionAddElement func(transaction IOHIDTransactionRef, element IOHIDElementRef)
var _iOHIDTransactionAddElementErr error

func tryIOHIDTransactionAddElement(transaction IOHIDTransactionRef, element IOHIDElementRef) error {
	if _iOHIDTransactionAddElement == nil {
		return symbolCallError("IOHIDTransactionAddElement", "10.5", _iOHIDTransactionAddElementErr)
	}
	_iOHIDTransactionAddElement(transaction, element)
	return nil
}

// IOHIDTransactionAddElement adds an element to the transaction @disussion To minimize device traffic it is important to add elements that share a common report type and report id.
//
// See: https://developer.apple.com/documentation/iokit/1561679-iohidtransactionaddelement
func IOHIDTransactionAddElement(transaction IOHIDTransactionRef, element IOHIDElementRef) {
	if callErr := tryIOHIDTransactionAddElement(transaction, element); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDTransactionClear func(transaction IOHIDTransactionRef)
var _iOHIDTransactionClearErr error

func tryIOHIDTransactionClear(transaction IOHIDTransactionRef) error {
	if _iOHIDTransactionClear == nil {
		return symbolCallError("IOHIDTransactionClear", "10.5", _iOHIDTransactionClearErr)
	}
	_iOHIDTransactionClear(transaction)
	return nil
}

// IOHIDTransactionClear clears element transaction values.
//
// See: https://developer.apple.com/documentation/iokit/1561687-iohidtransactionclear
func IOHIDTransactionClear(transaction IOHIDTransactionRef) {
	if callErr := tryIOHIDTransactionClear(transaction); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDTransactionCommit func(transaction IOHIDTransactionRef) kernel.IOReturn
var _iOHIDTransactionCommitErr error

func tryIOHIDTransactionCommit(transaction IOHIDTransactionRef) (kernel.IOReturn, error) {
	if _iOHIDTransactionCommit == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDTransactionCommit", "10.5", _iOHIDTransactionCommitErr)
	}
	return _iOHIDTransactionCommit(transaction), nil
}

// IOHIDTransactionCommit synchronously commits element transaction to the device.
//
// See: https://developer.apple.com/documentation/iokit/1561681-iohidtransactioncommit
func IOHIDTransactionCommit(transaction IOHIDTransactionRef) kernel.IOReturn {
	result, callErr := tryIOHIDTransactionCommit(transaction)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDTransactionCommitWithCallback func(transaction IOHIDTransactionRef, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) kernel.IOReturn
var _iOHIDTransactionCommitWithCallbackErr error

func tryIOHIDTransactionCommitWithCallback(transaction IOHIDTransactionRef, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) (kernel.IOReturn, error) {
	if _iOHIDTransactionCommitWithCallback == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDTransactionCommitWithCallback", "10.5", _iOHIDTransactionCommitWithCallbackErr)
	}
	return _iOHIDTransactionCommitWithCallback(transaction, timeout, callback, context), nil
}

// IOHIDTransactionCommitWithCallback commits element transaction to the device.
//
// See: https://developer.apple.com/documentation/iokit/1561677-iohidtransactioncommitwithcallba
func IOHIDTransactionCommitWithCallback(transaction IOHIDTransactionRef, timeout corefoundation.CFTimeInterval, callback unsafe.Pointer, context unsafe.Pointer) kernel.IOReturn {
	result, callErr := tryIOHIDTransactionCommitWithCallback(transaction, timeout, callback, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDTransactionContainsElement func(transaction IOHIDTransactionRef, element IOHIDElementRef) bool
var _iOHIDTransactionContainsElementErr error

func tryIOHIDTransactionContainsElement(transaction IOHIDTransactionRef, element IOHIDElementRef) (bool, error) {
	if _iOHIDTransactionContainsElement == nil {
		return false, symbolCallError("IOHIDTransactionContainsElement", "10.5", _iOHIDTransactionContainsElementErr)
	}
	return _iOHIDTransactionContainsElement(transaction, element), nil
}

// IOHIDTransactionContainsElement queries the transaction to determine if elemement has been added.
//
// See: https://developer.apple.com/documentation/iokit/1561680-iohidtransactioncontainselement
func IOHIDTransactionContainsElement(transaction IOHIDTransactionRef, element IOHIDElementRef) bool {
	result, callErr := tryIOHIDTransactionContainsElement(transaction, element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDTransactionCreate func(allocator corefoundation.CFAllocatorRef, device IOHIDDeviceRef, direction IOHIDTransactionDirectionType, options uint32) IOHIDTransactionRef
var _iOHIDTransactionCreateErr error

func tryIOHIDTransactionCreate(allocator corefoundation.CFAllocatorRef, device IOHIDDeviceRef, direction IOHIDTransactionDirectionType, options uint32) (IOHIDTransactionRef, error) {
	if _iOHIDTransactionCreate == nil {
		return *new(IOHIDTransactionRef), symbolCallError("IOHIDTransactionCreate", "10.5", _iOHIDTransactionCreateErr)
	}
	return _iOHIDTransactionCreate(allocator, device, direction, options), nil
}

// IOHIDTransactionCreate creates an IOHIDTransaction object for the specified device.
//
// See: https://developer.apple.com/documentation/iokit/1561689-iohidtransactioncreate
func IOHIDTransactionCreate(allocator corefoundation.CFAllocatorRef, device IOHIDDeviceRef, direction IOHIDTransactionDirectionType, options uint32) IOHIDTransactionRef {
	result, callErr := tryIOHIDTransactionCreate(allocator, device, direction, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDTransactionGetDevice func(transaction IOHIDTransactionRef) IOHIDDeviceRef
var _iOHIDTransactionGetDeviceErr error

func tryIOHIDTransactionGetDevice(transaction IOHIDTransactionRef) (IOHIDDeviceRef, error) {
	if _iOHIDTransactionGetDevice == nil {
		return *new(IOHIDDeviceRef), symbolCallError("IOHIDTransactionGetDevice", "10.5", _iOHIDTransactionGetDeviceErr)
	}
	return _iOHIDTransactionGetDevice(transaction), nil
}

// IOHIDTransactionGetDevice obtain the device associated with the transaction.
//
// See: https://developer.apple.com/documentation/iokit/1561685-iohidtransactiongetdevice
func IOHIDTransactionGetDevice(transaction IOHIDTransactionRef) IOHIDDeviceRef {
	result, callErr := tryIOHIDTransactionGetDevice(transaction)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDTransactionGetDirection func(transaction IOHIDTransactionRef) IOHIDTransactionDirectionType
var _iOHIDTransactionGetDirectionErr error

func tryIOHIDTransactionGetDirection(transaction IOHIDTransactionRef) (IOHIDTransactionDirectionType, error) {
	if _iOHIDTransactionGetDirection == nil {
		return *new(IOHIDTransactionDirectionType), symbolCallError("IOHIDTransactionGetDirection", "10.5", _iOHIDTransactionGetDirectionErr)
	}
	return _iOHIDTransactionGetDirection(transaction), nil
}

// IOHIDTransactionGetDirection obtain the direction of the transaction.
//
// See: https://developer.apple.com/documentation/iokit/1561674-iohidtransactiongetdirection
func IOHIDTransactionGetDirection(transaction IOHIDTransactionRef) IOHIDTransactionDirectionType {
	result, callErr := tryIOHIDTransactionGetDirection(transaction)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDTransactionGetTypeID func() uint
var _iOHIDTransactionGetTypeIDErr error

func tryIOHIDTransactionGetTypeID() (uint, error) {
	if _iOHIDTransactionGetTypeID == nil {
		return 0, symbolCallError("IOHIDTransactionGetTypeID", "10.5", _iOHIDTransactionGetTypeIDErr)
	}
	return _iOHIDTransactionGetTypeID(), nil
}

// IOHIDTransactionGetTypeID returns the type identifier of all IOHIDTransaction instances.
//
// See: https://developer.apple.com/documentation/iokit/1561678-iohidtransactiongettypeid
func IOHIDTransactionGetTypeID() uint {
	result, callErr := tryIOHIDTransactionGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDTransactionGetValue func(transaction IOHIDTransactionRef, element IOHIDElementRef, options uint32) IOHIDValueRef
var _iOHIDTransactionGetValueErr error

func tryIOHIDTransactionGetValue(transaction IOHIDTransactionRef, element IOHIDElementRef, options uint32) (IOHIDValueRef, error) {
	if _iOHIDTransactionGetValue == nil {
		return *new(IOHIDValueRef), symbolCallError("IOHIDTransactionGetValue", "10.5", _iOHIDTransactionGetValueErr)
	}
	return _iOHIDTransactionGetValue(transaction, element, options), nil
}

// IOHIDTransactionGetValue obtains the value for a transaction element.
//
// See: https://developer.apple.com/documentation/iokit/1561683-iohidtransactiongetvalue
func IOHIDTransactionGetValue(transaction IOHIDTransactionRef, element IOHIDElementRef, options uint32) IOHIDValueRef {
	result, callErr := tryIOHIDTransactionGetValue(transaction, element, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDTransactionRemoveElement func(transaction IOHIDTransactionRef, element IOHIDElementRef)
var _iOHIDTransactionRemoveElementErr error

func tryIOHIDTransactionRemoveElement(transaction IOHIDTransactionRef, element IOHIDElementRef) error {
	if _iOHIDTransactionRemoveElement == nil {
		return symbolCallError("IOHIDTransactionRemoveElement", "10.5", _iOHIDTransactionRemoveElementErr)
	}
	_iOHIDTransactionRemoveElement(transaction, element)
	return nil
}

// IOHIDTransactionRemoveElement removes an element to the transaction
//
// See: https://developer.apple.com/documentation/iokit/1561686-iohidtransactionremoveelement
func IOHIDTransactionRemoveElement(transaction IOHIDTransactionRef, element IOHIDElementRef) {
	if callErr := tryIOHIDTransactionRemoveElement(transaction, element); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDTransactionScheduleWithRunLoop func(transaction IOHIDTransactionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef)
var _iOHIDTransactionScheduleWithRunLoopErr error

func tryIOHIDTransactionScheduleWithRunLoop(transaction IOHIDTransactionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) error {
	if _iOHIDTransactionScheduleWithRunLoop == nil {
		return symbolCallError("IOHIDTransactionScheduleWithRunLoop", "10.5", _iOHIDTransactionScheduleWithRunLoopErr)
	}
	_iOHIDTransactionScheduleWithRunLoop(transaction, runLoop, runLoopMode)
	return nil
}

// IOHIDTransactionScheduleWithRunLoop schedules transaction with run loop.
//
// See: https://developer.apple.com/documentation/iokit/1561675-iohidtransactionschedulewithrunl
func IOHIDTransactionScheduleWithRunLoop(transaction IOHIDTransactionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) {
	if callErr := tryIOHIDTransactionScheduleWithRunLoop(transaction, runLoop, runLoopMode); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDTransactionSetDirection func(transaction IOHIDTransactionRef, direction IOHIDTransactionDirectionType)
var _iOHIDTransactionSetDirectionErr error

func tryIOHIDTransactionSetDirection(transaction IOHIDTransactionRef, direction IOHIDTransactionDirectionType) error {
	if _iOHIDTransactionSetDirection == nil {
		return symbolCallError("IOHIDTransactionSetDirection", "10.5", _iOHIDTransactionSetDirectionErr)
	}
	_iOHIDTransactionSetDirection(transaction, direction)
	return nil
}

// IOHIDTransactionSetDirection sets the direction of the transaction @disussion This method is useful for manipulating bi-direction (feature) elements such that you can set or get element values without creating an additional transaction object.
//
// See: https://developer.apple.com/documentation/iokit/1561688-iohidtransactionsetdirection
func IOHIDTransactionSetDirection(transaction IOHIDTransactionRef, direction IOHIDTransactionDirectionType) {
	if callErr := tryIOHIDTransactionSetDirection(transaction, direction); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDTransactionSetValue func(transaction IOHIDTransactionRef, element IOHIDElementRef, value IOHIDValueRef, options uint32)
var _iOHIDTransactionSetValueErr error

func tryIOHIDTransactionSetValue(transaction IOHIDTransactionRef, element IOHIDElementRef, value IOHIDValueRef, options uint32) error {
	if _iOHIDTransactionSetValue == nil {
		return symbolCallError("IOHIDTransactionSetValue", "10.5", _iOHIDTransactionSetValueErr)
	}
	_iOHIDTransactionSetValue(transaction, element, value, options)
	return nil
}

// IOHIDTransactionSetValue sets the value for a transaction element.
//
// See: https://developer.apple.com/documentation/iokit/1561676-iohidtransactionsetvalue
func IOHIDTransactionSetValue(transaction IOHIDTransactionRef, element IOHIDElementRef, value IOHIDValueRef, options uint32) {
	if callErr := tryIOHIDTransactionSetValue(transaction, element, value, options); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDTransactionUnscheduleFromRunLoop func(transaction IOHIDTransactionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef)
var _iOHIDTransactionUnscheduleFromRunLoopErr error

func tryIOHIDTransactionUnscheduleFromRunLoop(transaction IOHIDTransactionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) error {
	if _iOHIDTransactionUnscheduleFromRunLoop == nil {
		return symbolCallError("IOHIDTransactionUnscheduleFromRunLoop", "10.5", _iOHIDTransactionUnscheduleFromRunLoopErr)
	}
	_iOHIDTransactionUnscheduleFromRunLoop(transaction, runLoop, runLoopMode)
	return nil
}

// IOHIDTransactionUnscheduleFromRunLoop unschedules transaction with run loop.
//
// See: https://developer.apple.com/documentation/iokit/1561682-iohidtransactionunschedulefromru
func IOHIDTransactionUnscheduleFromRunLoop(transaction IOHIDTransactionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) {
	if callErr := tryIOHIDTransactionUnscheduleFromRunLoop(transaction, runLoop, runLoopMode); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDUnregisterVirtualDisplay func(handle uint32, display_token uint32) int32
var _iOHIDUnregisterVirtualDisplayErr error

func tryIOHIDUnregisterVirtualDisplay(handle uint32, display_token uint32) (int32, error) {
	if _iOHIDUnregisterVirtualDisplay == nil {
		return 0, symbolCallError("IOHIDUnregisterVirtualDisplay", "10.0", _iOHIDUnregisterVirtualDisplayErr)
	}
	return _iOHIDUnregisterVirtualDisplay(handle, display_token), nil
}

// IOHIDUnregisterVirtualDisplay.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/iokit/1555410-iohidunregistervirtualdisplay
func IOHIDUnregisterVirtualDisplay(handle uint32, display_token uint32) int32 {
	result, callErr := tryIOHIDUnregisterVirtualDisplay(handle, display_token)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDUserDeviceActivate func(device IOHIDUserDeviceRef)
var _iOHIDUserDeviceActivateErr error

func tryIOHIDUserDeviceActivate(device IOHIDUserDeviceRef) error {
	if _iOHIDUserDeviceActivate == nil {
		return symbolCallError("IOHIDUserDeviceActivate", "10.15", _iOHIDUserDeviceActivateErr)
	}
	_iOHIDUserDeviceActivate(device)
	return nil
}

// IOHIDUserDeviceActivate.
//
// See: https://developer.apple.com/documentation/iokit/3334949-iohiduserdeviceactivate
func IOHIDUserDeviceActivate(device IOHIDUserDeviceRef) {
	if callErr := tryIOHIDUserDeviceActivate(device); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDUserDeviceCancel func(device IOHIDUserDeviceRef)
var _iOHIDUserDeviceCancelErr error

func tryIOHIDUserDeviceCancel(device IOHIDUserDeviceRef) error {
	if _iOHIDUserDeviceCancel == nil {
		return symbolCallError("IOHIDUserDeviceCancel", "10.15", _iOHIDUserDeviceCancelErr)
	}
	_iOHIDUserDeviceCancel(device)
	return nil
}

// IOHIDUserDeviceCancel.
//
// See: https://developer.apple.com/documentation/iokit/3334950-iohiduserdevicecancel
func IOHIDUserDeviceCancel(device IOHIDUserDeviceRef) {
	if callErr := tryIOHIDUserDeviceCancel(device); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDUserDeviceCopyProperty func(device IOHIDUserDeviceRef, key corefoundation.CFStringRef) corefoundation.CFTypeRef
var _iOHIDUserDeviceCopyPropertyErr error

func tryIOHIDUserDeviceCopyProperty(device IOHIDUserDeviceRef, key corefoundation.CFStringRef) (corefoundation.CFTypeRef, error) {
	if _iOHIDUserDeviceCopyProperty == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("IOHIDUserDeviceCopyProperty", "10.15", _iOHIDUserDeviceCopyPropertyErr)
	}
	return _iOHIDUserDeviceCopyProperty(device, key), nil
}

// IOHIDUserDeviceCopyProperty.
//
// See: https://developer.apple.com/documentation/iokit/3334951-iohiduserdevicecopyproperty
func IOHIDUserDeviceCopyProperty(device IOHIDUserDeviceRef, key corefoundation.CFStringRef) corefoundation.CFTypeRef {
	result, callErr := tryIOHIDUserDeviceCopyProperty(device, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDUserDeviceCreateWithProperties func(allocator corefoundation.CFAllocatorRef, properties corefoundation.CFDictionaryRef, options uint32) IOHIDUserDeviceRef
var _iOHIDUserDeviceCreateWithPropertiesErr error

func tryIOHIDUserDeviceCreateWithProperties(allocator corefoundation.CFAllocatorRef, properties corefoundation.CFDictionaryRef, options uint32) (IOHIDUserDeviceRef, error) {
	if _iOHIDUserDeviceCreateWithProperties == nil {
		return *new(IOHIDUserDeviceRef), symbolCallError("IOHIDUserDeviceCreateWithProperties", "10.15", _iOHIDUserDeviceCreateWithPropertiesErr)
	}
	return _iOHIDUserDeviceCreateWithProperties(allocator, properties, options), nil
}

// IOHIDUserDeviceCreateWithProperties.
//
// See: https://developer.apple.com/documentation/iokit/3334952-iohiduserdevicecreatewithpropert
func IOHIDUserDeviceCreateWithProperties(allocator corefoundation.CFAllocatorRef, properties corefoundation.CFDictionaryRef, options uint32) IOHIDUserDeviceRef {
	result, callErr := tryIOHIDUserDeviceCreateWithProperties(allocator, properties, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDUserDeviceGetTypeID func() uint
var _iOHIDUserDeviceGetTypeIDErr error

func tryIOHIDUserDeviceGetTypeID() (uint, error) {
	if _iOHIDUserDeviceGetTypeID == nil {
		return 0, symbolCallError("IOHIDUserDeviceGetTypeID", "10.15", _iOHIDUserDeviceGetTypeIDErr)
	}
	return _iOHIDUserDeviceGetTypeID(), nil
}

// IOHIDUserDeviceGetTypeID.
//
// See: https://developer.apple.com/documentation/iokit/3334954-iohiduserdevicegettypeid
func IOHIDUserDeviceGetTypeID() uint {
	result, callErr := tryIOHIDUserDeviceGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDUserDeviceHandleReportWithTimeStamp func(device IOHIDUserDeviceRef, timestamp uint64, report *byte, reportLength corefoundation.CFIndex) kernel.IOReturn
var _iOHIDUserDeviceHandleReportWithTimeStampErr error

func tryIOHIDUserDeviceHandleReportWithTimeStamp(device IOHIDUserDeviceRef, timestamp uint64, report []byte, reportLength corefoundation.CFIndex) (kernel.IOReturn, error) {
	if _iOHIDUserDeviceHandleReportWithTimeStamp == nil {
		return *new(kernel.IOReturn), symbolCallError("IOHIDUserDeviceHandleReportWithTimeStamp", "10.15", _iOHIDUserDeviceHandleReportWithTimeStampErr)
	}
	return _iOHIDUserDeviceHandleReportWithTimeStamp(device, timestamp, unsafe.SliceData(report), reportLength), nil
}

// IOHIDUserDeviceHandleReportWithTimeStamp.
//
// See: https://developer.apple.com/documentation/iokit/3334955-iohiduserdevicehandlereportwitht
func IOHIDUserDeviceHandleReportWithTimeStamp(device IOHIDUserDeviceRef, timestamp uint64, report []byte, reportLength corefoundation.CFIndex) kernel.IOReturn {
	result, callErr := tryIOHIDUserDeviceHandleReportWithTimeStamp(device, timestamp, report, reportLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDUserDeviceRegisterGetReportBlock func(device IOHIDUserDeviceRef, block unsafe.Pointer)
var _iOHIDUserDeviceRegisterGetReportBlockErr error

func tryIOHIDUserDeviceRegisterGetReportBlock(device IOHIDUserDeviceRef, block IOHIDUserDeviceGetReportBlock) error {
	if _iOHIDUserDeviceRegisterGetReportBlock == nil {
		return symbolCallError("IOHIDUserDeviceRegisterGetReportBlock", "10.15", _iOHIDUserDeviceRegisterGetReportBlockErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 IOHIDReportType, blockArg1 uint32, blockArg2 *uint8, blockArg3 unsafe.Pointer) int32 {
		return block(blockArg0, blockArg1, blockArg2, (*corefoundation.CFIndex)(blockArg3))
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_iOHIDUserDeviceRegisterGetReportBlock(device, _block0)
	return nil
}

// IOHIDUserDeviceRegisterGetReportBlock.
//
// See: https://developer.apple.com/documentation/iokit/3334959-iohiduserdeviceregistergetreport
func IOHIDUserDeviceRegisterGetReportBlock(device IOHIDUserDeviceRef, block IOHIDUserDeviceGetReportBlock) {
	if callErr := tryIOHIDUserDeviceRegisterGetReportBlock(device, block); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDUserDeviceRegisterSetReportBlock func(device IOHIDUserDeviceRef, block unsafe.Pointer)
var _iOHIDUserDeviceRegisterSetReportBlockErr error

func tryIOHIDUserDeviceRegisterSetReportBlock(device IOHIDUserDeviceRef, block IOHIDUserDeviceSetReportBlock) error {
	if _iOHIDUserDeviceRegisterSetReportBlock == nil {
		return symbolCallError("IOHIDUserDeviceRegisterSetReportBlock", "10.15", _iOHIDUserDeviceRegisterSetReportBlockErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 IOHIDReportType, blockArg1 uint32, blockArg2 *uint8, blockArg3 corefoundation.CFIndex) int32 {
		return block(blockArg0, blockArg1, blockArg2, blockArg3)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_iOHIDUserDeviceRegisterSetReportBlock(device, _block0)
	return nil
}

// IOHIDUserDeviceRegisterSetReportBlock.
//
// See: https://developer.apple.com/documentation/iokit/3334960-iohiduserdeviceregistersetreport
func IOHIDUserDeviceRegisterSetReportBlock(device IOHIDUserDeviceRef, block IOHIDUserDeviceSetReportBlock) {
	if callErr := tryIOHIDUserDeviceRegisterSetReportBlock(device, block); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDUserDeviceSetCancelHandler func(device IOHIDUserDeviceRef, handler unsafe.Pointer)
var _iOHIDUserDeviceSetCancelHandlerErr error

func tryIOHIDUserDeviceSetCancelHandler(device IOHIDUserDeviceRef, handler unsafe.Pointer) error {
	if _iOHIDUserDeviceSetCancelHandler == nil {
		return symbolCallError("IOHIDUserDeviceSetCancelHandler", "10.15", _iOHIDUserDeviceSetCancelHandlerErr)
	}
	_iOHIDUserDeviceSetCancelHandler(device, handler)
	return nil
}

// IOHIDUserDeviceSetCancelHandler.
//
// See: https://developer.apple.com/documentation/iokit/3334961-iohiduserdevicesetcancelhandler
func IOHIDUserDeviceSetCancelHandler(device IOHIDUserDeviceRef, handler unsafe.Pointer) {
	if callErr := tryIOHIDUserDeviceSetCancelHandler(device, handler); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDUserDeviceSetDispatchQueue func(device IOHIDUserDeviceRef, queue uintptr)
var _iOHIDUserDeviceSetDispatchQueueErr error

func tryIOHIDUserDeviceSetDispatchQueue(device IOHIDUserDeviceRef, queue dispatch.Queue) error {
	if _iOHIDUserDeviceSetDispatchQueue == nil {
		return symbolCallError("IOHIDUserDeviceSetDispatchQueue", "10.15", _iOHIDUserDeviceSetDispatchQueueErr)
	}
	_iOHIDUserDeviceSetDispatchQueue(device, uintptr(queue.Handle()))
	return nil
}

// IOHIDUserDeviceSetDispatchQueue.
//
// See: https://developer.apple.com/documentation/iokit/3334962-iohiduserdevicesetdispatchqueue
func IOHIDUserDeviceSetDispatchQueue(device IOHIDUserDeviceRef, queue dispatch.Queue) {
	if callErr := tryIOHIDUserDeviceSetDispatchQueue(device, queue); callErr != nil {
		panic(callErr)
	}
}

var _iOHIDUserDeviceSetProperty func(device IOHIDUserDeviceRef, key corefoundation.CFStringRef, property corefoundation.CFTypeRef) bool
var _iOHIDUserDeviceSetPropertyErr error

func tryIOHIDUserDeviceSetProperty(device IOHIDUserDeviceRef, key corefoundation.CFStringRef, property corefoundation.CFTypeRef) (bool, error) {
	if _iOHIDUserDeviceSetProperty == nil {
		return false, symbolCallError("IOHIDUserDeviceSetProperty", "10.15", _iOHIDUserDeviceSetPropertyErr)
	}
	return _iOHIDUserDeviceSetProperty(device, key, property), nil
}

// IOHIDUserDeviceSetProperty.
//
// See: https://developer.apple.com/documentation/iokit/3334963-iohiduserdevicesetproperty
func IOHIDUserDeviceSetProperty(device IOHIDUserDeviceRef, key corefoundation.CFStringRef, property corefoundation.CFTypeRef) bool {
	result, callErr := tryIOHIDUserDeviceSetProperty(device, key, property)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDValueCreateWithBytes func(allocator corefoundation.CFAllocatorRef, element IOHIDElementRef, timeStamp uint64, bytes *byte, length corefoundation.CFIndex) IOHIDValueRef
var _iOHIDValueCreateWithBytesErr error

func tryIOHIDValueCreateWithBytes(allocator corefoundation.CFAllocatorRef, element IOHIDElementRef, timeStamp uint64, bytes []byte, length corefoundation.CFIndex) (IOHIDValueRef, error) {
	if _iOHIDValueCreateWithBytes == nil {
		return *new(IOHIDValueRef), symbolCallError("IOHIDValueCreateWithBytes", "10.5", _iOHIDValueCreateWithBytesErr)
	}
	return _iOHIDValueCreateWithBytes(allocator, element, timeStamp, unsafe.SliceData(bytes), length), nil
}

// IOHIDValueCreateWithBytes creates a new element value using byte data.
//
// See: https://developer.apple.com/documentation/iokit/1433290-iohidvaluecreatewithbytes
func IOHIDValueCreateWithBytes(allocator corefoundation.CFAllocatorRef, element IOHIDElementRef, timeStamp uint64, bytes []byte, length corefoundation.CFIndex) IOHIDValueRef {
	result, callErr := tryIOHIDValueCreateWithBytes(allocator, element, timeStamp, bytes, length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDValueCreateWithBytesNoCopy func(allocator corefoundation.CFAllocatorRef, element IOHIDElementRef, timeStamp uint64, bytes *byte, length corefoundation.CFIndex) IOHIDValueRef
var _iOHIDValueCreateWithBytesNoCopyErr error

func tryIOHIDValueCreateWithBytesNoCopy(allocator corefoundation.CFAllocatorRef, element IOHIDElementRef, timeStamp uint64, bytes []byte, length corefoundation.CFIndex) (IOHIDValueRef, error) {
	if _iOHIDValueCreateWithBytesNoCopy == nil {
		return *new(IOHIDValueRef), symbolCallError("IOHIDValueCreateWithBytesNoCopy", "10.5", _iOHIDValueCreateWithBytesNoCopyErr)
	}
	return _iOHIDValueCreateWithBytesNoCopy(allocator, element, timeStamp, unsafe.SliceData(bytes), length), nil
}

// IOHIDValueCreateWithBytesNoCopy creates a new element value using byte data without performing a copy.
//
// See: https://developer.apple.com/documentation/iokit/1433287-iohidvaluecreatewithbytesnocopy
func IOHIDValueCreateWithBytesNoCopy(allocator corefoundation.CFAllocatorRef, element IOHIDElementRef, timeStamp uint64, bytes []byte, length corefoundation.CFIndex) IOHIDValueRef {
	result, callErr := tryIOHIDValueCreateWithBytesNoCopy(allocator, element, timeStamp, bytes, length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDValueCreateWithIntegerValue func(allocator corefoundation.CFAllocatorRef, element IOHIDElementRef, timeStamp uint64, value corefoundation.CFIndex) IOHIDValueRef
var _iOHIDValueCreateWithIntegerValueErr error

func tryIOHIDValueCreateWithIntegerValue(allocator corefoundation.CFAllocatorRef, element IOHIDElementRef, timeStamp uint64, value corefoundation.CFIndex) (IOHIDValueRef, error) {
	if _iOHIDValueCreateWithIntegerValue == nil {
		return *new(IOHIDValueRef), symbolCallError("IOHIDValueCreateWithIntegerValue", "10.5", _iOHIDValueCreateWithIntegerValueErr)
	}
	return _iOHIDValueCreateWithIntegerValue(allocator, element, timeStamp, value), nil
}

// IOHIDValueCreateWithIntegerValue creates a new element value using an integer value.
//
// See: https://developer.apple.com/documentation/iokit/1433294-iohidvaluecreatewithintegervalue
func IOHIDValueCreateWithIntegerValue(allocator corefoundation.CFAllocatorRef, element IOHIDElementRef, timeStamp uint64, value corefoundation.CFIndex) IOHIDValueRef {
	result, callErr := tryIOHIDValueCreateWithIntegerValue(allocator, element, timeStamp, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDValueGetBytePtr func(value IOHIDValueRef) *uint8
var _iOHIDValueGetBytePtrErr error

func tryIOHIDValueGetBytePtr(value IOHIDValueRef) (*uint8, error) {
	if _iOHIDValueGetBytePtr == nil {
		return nil, symbolCallError("IOHIDValueGetBytePtr", "10.5", _iOHIDValueGetBytePtrErr)
	}
	return _iOHIDValueGetBytePtr(value), nil
}

// IOHIDValueGetBytePtr returns a byte pointer to the value contained in this IOHIDValueRef.
//
// See: https://developer.apple.com/documentation/iokit/1433292-iohidvaluegetbyteptr
func IOHIDValueGetBytePtr(value IOHIDValueRef) *uint8 {
	result, callErr := tryIOHIDValueGetBytePtr(value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDValueGetElement func(value IOHIDValueRef) IOHIDElementRef
var _iOHIDValueGetElementErr error

func tryIOHIDValueGetElement(value IOHIDValueRef) (IOHIDElementRef, error) {
	if _iOHIDValueGetElement == nil {
		return *new(IOHIDElementRef), symbolCallError("IOHIDValueGetElement", "10.5", _iOHIDValueGetElementErr)
	}
	return _iOHIDValueGetElement(value), nil
}

// IOHIDValueGetElement returns the element value associated with this IOHIDValueRef.
//
// See: https://developer.apple.com/documentation/iokit/1433285-iohidvaluegetelement
func IOHIDValueGetElement(value IOHIDValueRef) IOHIDElementRef {
	result, callErr := tryIOHIDValueGetElement(value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDValueGetIntegerValue func(value IOHIDValueRef) corefoundation.CFIndex
var _iOHIDValueGetIntegerValueErr error

func tryIOHIDValueGetIntegerValue(value IOHIDValueRef) (corefoundation.CFIndex, error) {
	if _iOHIDValueGetIntegerValue == nil {
		return *new(corefoundation.CFIndex), symbolCallError("IOHIDValueGetIntegerValue", "10.5", _iOHIDValueGetIntegerValueErr)
	}
	return _iOHIDValueGetIntegerValue(value), nil
}

// IOHIDValueGetIntegerValue returns an integer representaion of the value contained in this IOHIDValueRef.
//
// See: https://developer.apple.com/documentation/iokit/1433289-iohidvaluegetintegervalue
func IOHIDValueGetIntegerValue(value IOHIDValueRef) corefoundation.CFIndex {
	result, callErr := tryIOHIDValueGetIntegerValue(value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDValueGetLength func(value IOHIDValueRef) corefoundation.CFIndex
var _iOHIDValueGetLengthErr error

func tryIOHIDValueGetLength(value IOHIDValueRef) (corefoundation.CFIndex, error) {
	if _iOHIDValueGetLength == nil {
		return *new(corefoundation.CFIndex), symbolCallError("IOHIDValueGetLength", "10.5", _iOHIDValueGetLengthErr)
	}
	return _iOHIDValueGetLength(value), nil
}

// IOHIDValueGetLength returns the size, in bytes, of the value contained in this IOHIDValueRef.
//
// See: https://developer.apple.com/documentation/iokit/1433291-iohidvaluegetlength
func IOHIDValueGetLength(value IOHIDValueRef) corefoundation.CFIndex {
	result, callErr := tryIOHIDValueGetLength(value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDValueGetScaledValue func(value IOHIDValueRef, type_ kernel.IOHIDValueScaleType) float64
var _iOHIDValueGetScaledValueErr error

func tryIOHIDValueGetScaledValue(value IOHIDValueRef, type_ kernel.IOHIDValueScaleType) (float64, error) {
	if _iOHIDValueGetScaledValue == nil {
		return 0.0, symbolCallError("IOHIDValueGetScaledValue", "10.5", _iOHIDValueGetScaledValueErr)
	}
	return _iOHIDValueGetScaledValue(value, type_), nil
}

// IOHIDValueGetScaledValue returns an scaled representaion of the value contained in this IOHIDValueRef based on the scale type.
//
// See: https://developer.apple.com/documentation/iokit/1433288-iohidvaluegetscaledvalue
func IOHIDValueGetScaledValue(value IOHIDValueRef, type_ kernel.IOHIDValueScaleType) float64 {
	result, callErr := tryIOHIDValueGetScaledValue(value, type_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDValueGetTimeStamp func(value IOHIDValueRef) uint64
var _iOHIDValueGetTimeStampErr error

func tryIOHIDValueGetTimeStamp(value IOHIDValueRef) (uint64, error) {
	if _iOHIDValueGetTimeStamp == nil {
		return 0, symbolCallError("IOHIDValueGetTimeStamp", "10.5", _iOHIDValueGetTimeStampErr)
	}
	return _iOHIDValueGetTimeStamp(value), nil
}

// IOHIDValueGetTimeStamp returns the timestamp value contained in this IOHIDValueRef.
//
// See: https://developer.apple.com/documentation/iokit/1433286-iohidvaluegettimestamp
func IOHIDValueGetTimeStamp(value IOHIDValueRef) uint64 {
	result, callErr := tryIOHIDValueGetTimeStamp(value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOHIDValueGetTypeID func() uint
var _iOHIDValueGetTypeIDErr error

func tryIOHIDValueGetTypeID() (uint, error) {
	if _iOHIDValueGetTypeID == nil {
		return 0, symbolCallError("IOHIDValueGetTypeID", "10.5", _iOHIDValueGetTypeIDErr)
	}
	return _iOHIDValueGetTypeID(), nil
}

// IOHIDValueGetTypeID returns the type identifier of all IOHIDValue instances.
//
// See: https://developer.apple.com/documentation/iokit/1433293-iohidvaluegettypeid
func IOHIDValueGetTypeID() uint {
	result, callErr := tryIOHIDValueGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOI2CCopyInterfaceForID func(identifier corefoundation.CFTypeRef, interface_ *uint32) kernel.IOReturn
var _iOI2CCopyInterfaceForIDErr error

func tryIOI2CCopyInterfaceForID(identifier corefoundation.CFTypeRef, interface_ *uint32) (kernel.IOReturn, error) {
	if _iOI2CCopyInterfaceForID == nil {
		return *new(kernel.IOReturn), symbolCallError("IOI2CCopyInterfaceForID", "10.3", _iOI2CCopyInterfaceForIDErr)
	}
	return _iOI2CCopyInterfaceForID(identifier, interface_), nil
}

// IOI2CCopyInterfaceForID.
//
// See: https://developer.apple.com/documentation/iokit/1410384-ioi2ccopyinterfaceforid
func IOI2CCopyInterfaceForID(identifier corefoundation.CFTypeRef, interface_ *uint32) kernel.IOReturn {
	result, callErr := tryIOI2CCopyInterfaceForID(identifier, interface_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOI2CInterfaceClose func(connect IOI2CConnectRef, options uint32) kernel.IOReturn
var _iOI2CInterfaceCloseErr error

func tryIOI2CInterfaceClose(connect IOI2CConnectRef, options uint32) (kernel.IOReturn, error) {
	if _iOI2CInterfaceClose == nil {
		return *new(kernel.IOReturn), symbolCallError("IOI2CInterfaceClose", "10.3", _iOI2CInterfaceCloseErr)
	}
	return _iOI2CInterfaceClose(connect, options), nil
}

// IOI2CInterfaceClose closes an IOI2CConnectRef.
//
// See: https://developer.apple.com/documentation/iokit/1410390-ioi2cinterfaceclose
func IOI2CInterfaceClose(connect IOI2CConnectRef, options uint32) kernel.IOReturn {
	result, callErr := tryIOI2CInterfaceClose(connect, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOI2CInterfaceOpen func(interface_ uint32, options uint32, connect *IOI2CConnectRef) kernel.IOReturn
var _iOI2CInterfaceOpenErr error

func tryIOI2CInterfaceOpen(interface_ uint32, options uint32, connect *IOI2CConnectRef) (kernel.IOReturn, error) {
	if _iOI2CInterfaceOpen == nil {
		return *new(kernel.IOReturn), symbolCallError("IOI2CInterfaceOpen", "10.3", _iOI2CInterfaceOpenErr)
	}
	return _iOI2CInterfaceOpen(interface_, options, connect), nil
}

// IOI2CInterfaceOpen opens an instance of an I2C bus interface, allowing I2C requests to be made.
//
// See: https://developer.apple.com/documentation/iokit/1410388-ioi2cinterfaceopen
func IOI2CInterfaceOpen(interface_ uint32, options uint32, connect *IOI2CConnectRef) kernel.IOReturn {
	result, callErr := tryIOI2CInterfaceOpen(interface_, options, connect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOI2CSendRequest func(connect IOI2CConnectRef, options uint32, request *IOI2CRequest) kernel.IOReturn
var _iOI2CSendRequestErr error

func tryIOI2CSendRequest(connect IOI2CConnectRef, options uint32, request *IOI2CRequest) (kernel.IOReturn, error) {
	if _iOI2CSendRequest == nil {
		return *new(kernel.IOReturn), symbolCallError("IOI2CSendRequest", "10.3", _iOI2CSendRequestErr)
	}
	return _iOI2CSendRequest(connect, options, request), nil
}

// IOI2CSendRequest carries out the I2C transaction specified by an IOI2CRequest structure.
//
// See: https://developer.apple.com/documentation/iokit/1410373-ioi2csendrequest
func IOI2CSendRequest(connect IOI2CConnectRef, options uint32, request *IOI2CRequest) kernel.IOReturn {
	result, callErr := tryIOI2CSendRequest(connect, options, request)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOIteratorIsValid func(iterator uint32) int32
var _iOIteratorIsValidErr error

func tryIOIteratorIsValid(iterator uint32) (int32, error) {
	if _iOIteratorIsValid == nil {
		return 0, symbolCallError("IOIteratorIsValid", "10.0", _iOIteratorIsValidErr)
	}
	return _iOIteratorIsValid(iterator), nil
}

// IOIteratorIsValid checks an iterator is still valid.
//
// See: https://developer.apple.com/documentation/iokit/1514556-ioiteratorisvalid
func IOIteratorIsValid(iterator uint32) int32 {
	result, callErr := tryIOIteratorIsValid(iterator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOIteratorNext func(iterator uint32) uint32
var _iOIteratorNextErr error

func tryIOIteratorNext(iterator uint32) (uint32, error) {
	if _iOIteratorNext == nil {
		return 0, symbolCallError("IOIteratorNext", "10.0", _iOIteratorNextErr)
	}
	return _iOIteratorNext(iterator), nil
}

// IOIteratorNext returns the next object in an iteration.
//
// See: https://developer.apple.com/documentation/iokit/1514741-ioiteratornext
func IOIteratorNext(iterator uint32) uint32 {
	result, callErr := tryIOIteratorNext(iterator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOIteratorReset func(iterator uint32)
var _iOIteratorResetErr error

func tryIOIteratorReset(iterator uint32) error {
	if _iOIteratorReset == nil {
		return symbolCallError("IOIteratorReset", "10.0", _iOIteratorResetErr)
	}
	_iOIteratorReset(iterator)
	return nil
}

// IOIteratorReset resets an iteration back to the beginning.
//
// See: https://developer.apple.com/documentation/iokit/1514379-ioiteratorreset
func IOIteratorReset(iterator uint32) {
	if callErr := tryIOIteratorReset(iterator); callErr != nil {
		panic(callErr)
	}
}

var _iOKitGetBusyState func(mainPort uint32, busyState *uint32) int32
var _iOKitGetBusyStateErr error

func tryIOKitGetBusyState(mainPort uint32, busyState *uint32) (int32, error) {
	if _iOKitGetBusyState == nil {
		return 0, symbolCallError("IOKitGetBusyState", "10.0", _iOKitGetBusyStateErr)
	}
	return _iOKitGetBusyState(mainPort, busyState), nil
}

// IOKitGetBusyState returns the busyState of all IOServices.
//
// See: https://developer.apple.com/documentation/iokit/1514460-iokitgetbusystate
func IOKitGetBusyState(mainPort uint32, busyState *uint32) int32 {
	result, callErr := tryIOKitGetBusyState(mainPort, busyState)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOKitWaitQuiet func(mainPort uint32, waitTime *kernel.MachTimespec) int32
var _iOKitWaitQuietErr error

func tryIOKitWaitQuiet(mainPort uint32, waitTime *kernel.MachTimespec) (int32, error) {
	if _iOKitWaitQuiet == nil {
		return 0, symbolCallError("IOKitWaitQuiet", "10.0", _iOKitWaitQuietErr)
	}
	return _iOKitWaitQuiet(mainPort, waitTime), nil
}

// IOKitWaitQuiet wait for a all IOServices' busyState to be zero.
//
// See: https://developer.apple.com/documentation/iokit/1514440-iokitwaitquiet
func IOKitWaitQuiet(mainPort uint32, waitTime *kernel.MachTimespec) int32 {
	result, callErr := tryIOKitWaitQuiet(mainPort, waitTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOMainPort func(bootstrapPort uint32, mainPort *uint32) int32
var _iOMainPortErr error

func tryIOMainPort(bootstrapPort uint32, mainPort *uint32) (int32, error) {
	if _iOMainPort == nil {
		return 0, symbolCallError("IOMainPort", "12.0", _iOMainPortErr)
	}
	return _iOMainPort(bootstrapPort, mainPort), nil
}

// IOMainPort.
//
// See: https://developer.apple.com/documentation/iokit/3753260-iomainport
func IOMainPort(bootstrapPort uint32, mainPort *uint32) int32 {
	result, callErr := tryIOMainPort(bootstrapPort, mainPort)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOMasterPort func(bootstrapPort uint32, mainPort *uint32) int32
var _iOMasterPortErr error

func tryIOMasterPort(bootstrapPort uint32, mainPort *uint32) (int32, error) {
	if _iOMasterPort == nil {
		return 0, symbolCallError("IOMasterPort", "10.0", _iOMasterPortErr)
	}
	return _iOMasterPort(bootstrapPort, mainPort), nil
}

// IOMasterPort returns the mach port used to initiate communication with IOKit.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/iokit/1514652-iomasterport
func IOMasterPort(bootstrapPort uint32, mainPort *uint32) int32 {
	result, callErr := tryIOMasterPort(bootstrapPort, mainPort)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iONetworkClose func(con uint32) kernel.IOReturn
var _iONetworkCloseErr error

func tryIONetworkClose(con uint32) (kernel.IOReturn, error) {
	if _iONetworkClose == nil {
		return *new(kernel.IOReturn), symbolCallError("IONetworkClose", "10.0", _iONetworkCloseErr)
	}
	return _iONetworkClose(con), nil
}

// IONetworkClose close the connection to an IONetworkInterface object.
//
// See: https://developer.apple.com/documentation/iokit/1572704-ionetworkclose
func IONetworkClose(con uint32) kernel.IOReturn {
	result, callErr := tryIONetworkClose(con)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iONetworkGetDataCapacity func(conObject uint32, dataHandle IONDHandle, capacityP *uint32) kernel.IOReturn
var _iONetworkGetDataCapacityErr error

func tryIONetworkGetDataCapacity(conObject uint32, dataHandle IONDHandle, capacityP *uint32) (kernel.IOReturn, error) {
	if _iONetworkGetDataCapacity == nil {
		return *new(kernel.IOReturn), symbolCallError("IONetworkGetDataCapacity", "10.0", _iONetworkGetDataCapacityErr)
	}
	return _iONetworkGetDataCapacity(conObject, dataHandle, capacityP), nil
}

// IONetworkGetDataCapacity get the capacity (in bytes) of a network data object.
//
// See: https://developer.apple.com/documentation/iokit/1572712-ionetworkgetdatacapacity
func IONetworkGetDataCapacity(conObject uint32, dataHandle IONDHandle, capacityP *uint32) kernel.IOReturn {
	result, callErr := tryIONetworkGetDataCapacity(conObject, dataHandle, capacityP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iONetworkGetDataHandle func(conObject uint32, dataName string, dataHandleP *IONDHandle) kernel.IOReturn
var _iONetworkGetDataHandleErr error

func tryIONetworkGetDataHandle(conObject uint32, dataName string, dataHandleP *IONDHandle) (kernel.IOReturn, error) {
	if _iONetworkGetDataHandle == nil {
		return *new(kernel.IOReturn), symbolCallError("IONetworkGetDataHandle", "10.0", _iONetworkGetDataHandleErr)
	}
	return _iONetworkGetDataHandle(conObject, dataName, dataHandleP), nil
}

// IONetworkGetDataHandle get the handle of a network data object with the given name.
//
// See: https://developer.apple.com/documentation/iokit/1572708-ionetworkgetdatahandle
func IONetworkGetDataHandle(conObject uint32, dataName string, dataHandleP *IONDHandle) kernel.IOReturn {
	result, callErr := tryIONetworkGetDataHandle(conObject, dataName, dataHandleP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iONetworkGetPacketFiltersMask func(connect uint32, filterGroup unsafe.Pointer, filtersMask *uint32, options uint32) kernel.IOReturn
var _iONetworkGetPacketFiltersMaskErr error

func tryIONetworkGetPacketFiltersMask(connect uint32, filterGroup unsafe.Pointer, filtersMask *uint32, options uint32) (kernel.IOReturn, error) {
	if _iONetworkGetPacketFiltersMask == nil {
		return *new(kernel.IOReturn), symbolCallError("IONetworkGetPacketFiltersMask", "10.1", _iONetworkGetPacketFiltersMaskErr)
	}
	return _iONetworkGetPacketFiltersMask(connect, filterGroup, filtersMask, options), nil
}

// IONetworkGetPacketFiltersMask get the packet filters for a given filter group.
//
// See: https://developer.apple.com/documentation/iokit/1572711-ionetworkgetpacketfiltersmask
func IONetworkGetPacketFiltersMask(connect uint32, filterGroup unsafe.Pointer, filtersMask *uint32, options uint32) kernel.IOReturn {
	result, callErr := tryIONetworkGetPacketFiltersMask(connect, filterGroup, filtersMask, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iONetworkOpen func(obj uint32, con *uint32) kernel.IOReturn
var _iONetworkOpenErr error

func tryIONetworkOpen(obj uint32, con *uint32) (kernel.IOReturn, error) {
	if _iONetworkOpen == nil {
		return *new(kernel.IOReturn), symbolCallError("IONetworkOpen", "10.0", _iONetworkOpenErr)
	}
	return _iONetworkOpen(obj, con), nil
}

// IONetworkOpen open a connection to an IONetworkInterface object.
//
// See: https://developer.apple.com/documentation/iokit/1572709-ionetworkopen
func IONetworkOpen(obj uint32, con *uint32) kernel.IOReturn {
	result, callErr := tryIONetworkOpen(obj, con)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iONetworkReadData func(conObj uint32, dataHandle IONDHandle, destBuf *byte, inOutSizeP *uint32) kernel.IOReturn
var _iONetworkReadDataErr error

func tryIONetworkReadData(conObj uint32, dataHandle IONDHandle, destBuf []byte, inOutSizeP *uint32) (kernel.IOReturn, error) {
	if _iONetworkReadData == nil {
		return *new(kernel.IOReturn), symbolCallError("IONetworkReadData", "10.0", _iONetworkReadDataErr)
	}
	return _iONetworkReadData(conObj, dataHandle, unsafe.SliceData(destBuf), inOutSizeP), nil
}

// IONetworkReadData read the buffer of a network data object.
//
// See: https://developer.apple.com/documentation/iokit/1572706-ionetworkreaddata
func IONetworkReadData(conObj uint32, dataHandle IONDHandle, destBuf []byte, inOutSizeP *uint32) kernel.IOReturn {
	result, callErr := tryIONetworkReadData(conObj, dataHandle, destBuf, inOutSizeP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iONetworkResetData func(conObject uint32, dataHandle IONDHandle) kernel.IOReturn
var _iONetworkResetDataErr error

func tryIONetworkResetData(conObject uint32, dataHandle IONDHandle) (kernel.IOReturn, error) {
	if _iONetworkResetData == nil {
		return *new(kernel.IOReturn), symbolCallError("IONetworkResetData", "10.0", _iONetworkResetDataErr)
	}
	return _iONetworkResetData(conObject, dataHandle), nil
}

// IONetworkResetData fill the buffer of a network data object with zeroes.
//
// See: https://developer.apple.com/documentation/iokit/1572710-ionetworkresetdata
func IONetworkResetData(conObject uint32, dataHandle IONDHandle) kernel.IOReturn {
	result, callErr := tryIONetworkResetData(conObject, dataHandle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iONetworkSetPacketFiltersMask func(connect uint32, filterGroup unsafe.Pointer, filtersMask uint32, options uint32) kernel.IOReturn
var _iONetworkSetPacketFiltersMaskErr error

func tryIONetworkSetPacketFiltersMask(connect uint32, filterGroup unsafe.Pointer, filtersMask uint32, options uint32) (kernel.IOReturn, error) {
	if _iONetworkSetPacketFiltersMask == nil {
		return *new(kernel.IOReturn), symbolCallError("IONetworkSetPacketFiltersMask", "10.1", _iONetworkSetPacketFiltersMaskErr)
	}
	return _iONetworkSetPacketFiltersMask(connect, filterGroup, filtersMask, options), nil
}

// IONetworkSetPacketFiltersMask set the packet filters for a given filter group.
//
// See: https://developer.apple.com/documentation/iokit/1572703-ionetworksetpacketfiltersmask
func IONetworkSetPacketFiltersMask(connect uint32, filterGroup unsafe.Pointer, filtersMask uint32, options uint32) kernel.IOReturn {
	result, callErr := tryIONetworkSetPacketFiltersMask(connect, filterGroup, filtersMask, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iONetworkWriteData func(conObj uint32, dataHandle IONDHandle, srcBuf *byte, inSize uint32) kernel.IOReturn
var _iONetworkWriteDataErr error

func tryIONetworkWriteData(conObj uint32, dataHandle IONDHandle, srcBuf []byte, inSize uint32) (kernel.IOReturn, error) {
	if _iONetworkWriteData == nil {
		return *new(kernel.IOReturn), symbolCallError("IONetworkWriteData", "10.0", _iONetworkWriteDataErr)
	}
	return _iONetworkWriteData(conObj, dataHandle, unsafe.SliceData(srcBuf), inSize), nil
}

// IONetworkWriteData write to the buffer of a network data object.
//
// See: https://developer.apple.com/documentation/iokit/1572707-ionetworkwritedata
func IONetworkWriteData(conObj uint32, dataHandle IONDHandle, srcBuf []byte, inSize uint32) kernel.IOReturn {
	result, callErr := tryIONetworkWriteData(conObj, dataHandle, srcBuf, inSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iONotificationPortCreate func(mainPort uint32) IONotificationPortRef
var _iONotificationPortCreateErr error

func tryIONotificationPortCreate(mainPort uint32) (IONotificationPortRef, error) {
	if _iONotificationPortCreate == nil {
		return *new(IONotificationPortRef), symbolCallError("IONotificationPortCreate", "10.0", _iONotificationPortCreateErr)
	}
	return _iONotificationPortCreate(mainPort), nil
}

// IONotificationPortCreate creates and returns a notification object for receiving IOKit notifications of new devices or state changes.
//
// See: https://developer.apple.com/documentation/iokit/1514480-ionotificationportcreate
func IONotificationPortCreate(mainPort uint32) IONotificationPortRef {
	result, callErr := tryIONotificationPortCreate(mainPort)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iONotificationPortDestroy func(notify IONotificationPortRef)
var _iONotificationPortDestroyErr error

func tryIONotificationPortDestroy(notify IONotificationPortRef) error {
	if _iONotificationPortDestroy == nil {
		return symbolCallError("IONotificationPortDestroy", "10.0", _iONotificationPortDestroyErr)
	}
	_iONotificationPortDestroy(notify)
	return nil
}

// IONotificationPortDestroy destroys a notification object created with IONotificationPortCreate.
//
// See: https://developer.apple.com/documentation/iokit/1514751-ionotificationportdestroy
func IONotificationPortDestroy(notify IONotificationPortRef) {
	if callErr := tryIONotificationPortDestroy(notify); callErr != nil {
		panic(callErr)
	}
}

var _iONotificationPortGetMachPort func(notify IONotificationPortRef) uint32
var _iONotificationPortGetMachPortErr error

func tryIONotificationPortGetMachPort(notify IONotificationPortRef) (uint32, error) {
	if _iONotificationPortGetMachPort == nil {
		return 0, symbolCallError("IONotificationPortGetMachPort", "10.0", _iONotificationPortGetMachPortErr)
	}
	return _iONotificationPortGetMachPort(notify), nil
}

// IONotificationPortGetMachPort returns a mach_port to be used to listen for notifications.
//
// See: https://developer.apple.com/documentation/iokit/1514875-ionotificationportgetmachport
func IONotificationPortGetMachPort(notify IONotificationPortRef) uint32 {
	result, callErr := tryIONotificationPortGetMachPort(notify)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iONotificationPortGetRunLoopSource func(notify IONotificationPortRef) corefoundation.CFRunLoopSourceRef
var _iONotificationPortGetRunLoopSourceErr error

func tryIONotificationPortGetRunLoopSource(notify IONotificationPortRef) (corefoundation.CFRunLoopSourceRef, error) {
	if _iONotificationPortGetRunLoopSource == nil {
		return *new(corefoundation.CFRunLoopSourceRef), symbolCallError("IONotificationPortGetRunLoopSource", "10.0", _iONotificationPortGetRunLoopSourceErr)
	}
	return _iONotificationPortGetRunLoopSource(notify), nil
}

// IONotificationPortGetRunLoopSource returns a CFRunLoopSource to be used to listen for notifications.
//
// See: https://developer.apple.com/documentation/iokit/1514599-ionotificationportgetrunloopsour
func IONotificationPortGetRunLoopSource(notify IONotificationPortRef) corefoundation.CFRunLoopSourceRef {
	result, callErr := tryIONotificationPortGetRunLoopSource(notify)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iONotificationPortSetDispatchQueue func(notify IONotificationPortRef, queue uintptr)
var _iONotificationPortSetDispatchQueueErr error

func tryIONotificationPortSetDispatchQueue(notify IONotificationPortRef, queue dispatch.Queue) error {
	if _iONotificationPortSetDispatchQueue == nil {
		return symbolCallError("IONotificationPortSetDispatchQueue", "10.6", _iONotificationPortSetDispatchQueueErr)
	}
	_iONotificationPortSetDispatchQueue(notify, uintptr(queue.Handle()))
	return nil
}

// IONotificationPortSetDispatchQueue sets a dispatch queue to be used to listen for notifications.
//
// See: https://developer.apple.com/documentation/iokit/1514596-ionotificationportsetdispatchque
func IONotificationPortSetDispatchQueue(notify IONotificationPortRef, queue dispatch.Queue) {
	if callErr := tryIONotificationPortSetDispatchQueue(notify, queue); callErr != nil {
		panic(callErr)
	}
}

var _iONotificationPortSetImportanceReceiver func(notify IONotificationPortRef) int32
var _iONotificationPortSetImportanceReceiverErr error

func tryIONotificationPortSetImportanceReceiver(notify IONotificationPortRef) (int32, error) {
	if _iONotificationPortSetImportanceReceiver == nil {
		return 0, symbolCallError("IONotificationPortSetImportanceReceiver", "10.13", _iONotificationPortSetImportanceReceiverErr)
	}
	return _iONotificationPortSetImportanceReceiver(notify), nil
}

// IONotificationPortSetImportanceReceiver.
//
// See: https://developer.apple.com/documentation/iokit/2870065-ionotificationportsetimportancer
func IONotificationPortSetImportanceReceiver(notify IONotificationPortRef) int32 {
	result, callErr := tryIONotificationPortSetImportanceReceiver(notify)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOObjectConformsTo func(object uint32, className string) int32
var _iOObjectConformsToErr error

func tryIOObjectConformsTo(object uint32, className string) (int32, error) {
	if _iOObjectConformsTo == nil {
		return 0, symbolCallError("IOObjectConformsTo", "10.0", _iOObjectConformsToErr)
	}
	return _iOObjectConformsTo(object, className), nil
}

// IOObjectConformsTo performs an OSDynamicCast operation on an IOKit object.
//
// See: https://developer.apple.com/documentation/iokit/1514505-ioobjectconformsto
func IOObjectConformsTo(object uint32, className string) int32 {
	result, callErr := tryIOObjectConformsTo(object, className)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOObjectCopyBundleIdentifierForClass func(classname corefoundation.CFStringRef) corefoundation.CFStringRef
var _iOObjectCopyBundleIdentifierForClassErr error

func tryIOObjectCopyBundleIdentifierForClass(classname corefoundation.CFStringRef) (corefoundation.CFStringRef, error) {
	if _iOObjectCopyBundleIdentifierForClass == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("IOObjectCopyBundleIdentifierForClass", "10.4", _iOObjectCopyBundleIdentifierForClassErr)
	}
	return _iOObjectCopyBundleIdentifierForClass(classname), nil
}

// IOObjectCopyBundleIdentifierForClass return the bundle identifier of the given class.
//
// See: https://developer.apple.com/documentation/iokit/1514375-ioobjectcopybundleidentifierforc
func IOObjectCopyBundleIdentifierForClass(classname corefoundation.CFStringRef) corefoundation.CFStringRef {
	result, callErr := tryIOObjectCopyBundleIdentifierForClass(classname)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOObjectCopyClass func(object uint32) corefoundation.CFStringRef
var _iOObjectCopyClassErr error

func tryIOObjectCopyClass(object uint32) (corefoundation.CFStringRef, error) {
	if _iOObjectCopyClass == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("IOObjectCopyClass", "10.4", _iOObjectCopyClassErr)
	}
	return _iOObjectCopyClass(object), nil
}

// IOObjectCopyClass return the class name of an IOKit object.
//
// See: https://developer.apple.com/documentation/iokit/1514781-ioobjectcopyclass
func IOObjectCopyClass(object uint32) corefoundation.CFStringRef {
	result, callErr := tryIOObjectCopyClass(object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOObjectCopySuperclassForClass func(classname corefoundation.CFStringRef) corefoundation.CFStringRef
var _iOObjectCopySuperclassForClassErr error

func tryIOObjectCopySuperclassForClass(classname corefoundation.CFStringRef) (corefoundation.CFStringRef, error) {
	if _iOObjectCopySuperclassForClass == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("IOObjectCopySuperclassForClass", "10.4", _iOObjectCopySuperclassForClassErr)
	}
	return _iOObjectCopySuperclassForClass(classname), nil
}

// IOObjectCopySuperclassForClass return the superclass name of the given class.
//
// See: https://developer.apple.com/documentation/iokit/1514635-ioobjectcopysuperclassforclass
func IOObjectCopySuperclassForClass(classname corefoundation.CFStringRef) corefoundation.CFStringRef {
	result, callErr := tryIOObjectCopySuperclassForClass(classname)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOObjectGetClass func(object uint32, className string) int32
var _iOObjectGetClassErr error

func tryIOObjectGetClass(object uint32, className string) (int32, error) {
	if _iOObjectGetClass == nil {
		return 0, symbolCallError("IOObjectGetClass", "10.0", _iOObjectGetClassErr)
	}
	return _iOObjectGetClass(object, className), nil
}

// IOObjectGetClass return the class name of an IOKit object.
//
// See: https://developer.apple.com/documentation/iokit/1514756-ioobjectgetclass
func IOObjectGetClass(object uint32, className string) int32 {
	result, callErr := tryIOObjectGetClass(object, className)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOObjectGetKernelRetainCount func(object uint32) uint32
var _iOObjectGetKernelRetainCountErr error

func tryIOObjectGetKernelRetainCount(object uint32) (uint32, error) {
	if _iOObjectGetKernelRetainCount == nil {
		return 0, symbolCallError("IOObjectGetKernelRetainCount", "10.6", _iOObjectGetKernelRetainCountErr)
	}
	return _iOObjectGetKernelRetainCount(object), nil
}

// IOObjectGetKernelRetainCount returns kernel retain count of an IOKit object.
//
// See: https://developer.apple.com/documentation/iokit/1514325-ioobjectgetkernelretaincount
func IOObjectGetKernelRetainCount(object uint32) uint32 {
	result, callErr := tryIOObjectGetKernelRetainCount(object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOObjectGetRetainCount func(object uint32) uint32
var _iOObjectGetRetainCountErr error

func tryIOObjectGetRetainCount(object uint32) (uint32, error) {
	if _iOObjectGetRetainCount == nil {
		return 0, symbolCallError("IOObjectGetRetainCount", "10.0", _iOObjectGetRetainCountErr)
	}
	return _iOObjectGetRetainCount(object), nil
}

// IOObjectGetRetainCount returns kernel retain count of an IOKit object.
//
// See: https://developer.apple.com/documentation/iokit/1514824-ioobjectgetretaincount
func IOObjectGetRetainCount(object uint32) uint32 {
	result, callErr := tryIOObjectGetRetainCount(object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOObjectGetUserRetainCount func(object uint32) uint32
var _iOObjectGetUserRetainCountErr error

func tryIOObjectGetUserRetainCount(object uint32) (uint32, error) {
	if _iOObjectGetUserRetainCount == nil {
		return 0, symbolCallError("IOObjectGetUserRetainCount", "10.6", _iOObjectGetUserRetainCountErr)
	}
	return _iOObjectGetUserRetainCount(object), nil
}

// IOObjectGetUserRetainCount returns the retain count for the current process of an IOKit object.
//
// See: https://developer.apple.com/documentation/iokit/1514464-ioobjectgetuserretaincount
func IOObjectGetUserRetainCount(object uint32) uint32 {
	result, callErr := tryIOObjectGetUserRetainCount(object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOObjectIsEqualTo func(object uint32, anObject uint32) int32
var _iOObjectIsEqualToErr error

func tryIOObjectIsEqualTo(object uint32, anObject uint32) (int32, error) {
	if _iOObjectIsEqualTo == nil {
		return 0, symbolCallError("IOObjectIsEqualTo", "10.0", _iOObjectIsEqualToErr)
	}
	return _iOObjectIsEqualTo(object, anObject), nil
}

// IOObjectIsEqualTo checks two object handles to see if they represent the same kernel object.
//
// See: https://developer.apple.com/documentation/iokit/1514563-ioobjectisequalto
func IOObjectIsEqualTo(object uint32, anObject uint32) int32 {
	result, callErr := tryIOObjectIsEqualTo(object, anObject)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOObjectRelease func(object uint32) int32
var _iOObjectReleaseErr error

func tryIOObjectRelease(object uint32) (int32, error) {
	if _iOObjectRelease == nil {
		return 0, symbolCallError("IOObjectRelease", "10.0", _iOObjectReleaseErr)
	}
	return _iOObjectRelease(object), nil
}

// IOObjectRelease releases an object handle previously returned by IOKitLib.
//
// See: https://developer.apple.com/documentation/iokit/1514627-ioobjectrelease
func IOObjectRelease(object uint32) int32 {
	result, callErr := tryIOObjectRelease(object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOObjectRetain func(object uint32) int32
var _iOObjectRetainErr error

func tryIOObjectRetain(object uint32) (int32, error) {
	if _iOObjectRetain == nil {
		return 0, symbolCallError("IOObjectRetain", "10.1", _iOObjectRetainErr)
	}
	return _iOObjectRetain(object), nil
}

// IOObjectRetain retains an object handle previously returned by IOKitLib.
//
// See: https://developer.apple.com/documentation/iokit/1514769-ioobjectretain
func IOObjectRetain(object uint32) int32 {
	result, callErr := tryIOObjectRetain(object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOOpenFirmwarePathMatching func(mainPort uint32, options uint32, path string) corefoundation.CFMutableDictionary
var _iOOpenFirmwarePathMatchingErr error

func tryIOOpenFirmwarePathMatching(mainPort uint32, options uint32, path string) (corefoundation.CFMutableDictionary, error) {
	if _iOOpenFirmwarePathMatching == nil {
		return *new(corefoundation.CFMutableDictionary), symbolCallError("IOOpenFirmwarePathMatching", "10.0", _iOOpenFirmwarePathMatchingErr)
	}
	return _iOOpenFirmwarePathMatching(mainPort, options, path), nil
}

// IOOpenFirmwarePathMatching.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/iokit/1514715-ioopenfirmwarepathmatching
func IOOpenFirmwarePathMatching(mainPort uint32, options uint32, path string) corefoundation.CFMutableDictionary {
	result, callErr := tryIOOpenFirmwarePathMatching(mainPort, options, path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMAssertionCopyProperties func(theAssertion IOPMAssertionID) corefoundation.CFDictionaryRef
var _iOPMAssertionCopyPropertiesErr error

func tryIOPMAssertionCopyProperties(theAssertion IOPMAssertionID) (corefoundation.CFDictionaryRef, error) {
	if _iOPMAssertionCopyProperties == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("IOPMAssertionCopyProperties", "10.7", _iOPMAssertionCopyPropertiesErr)
	}
	return _iOPMAssertionCopyProperties(theAssertion), nil
}

// IOPMAssertionCopyProperties copies details about an [IOPMAssertion]
//
// See: https://developer.apple.com/documentation/iokit/1557066-iopmassertioncopyproperties
func IOPMAssertionCopyProperties(theAssertion IOPMAssertionID) corefoundation.CFDictionaryRef {
	result, callErr := tryIOPMAssertionCopyProperties(theAssertion)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMAssertionCreate func(AssertionType corefoundation.CFStringRef, AssertionLevel IOPMAssertionLevel, AssertionID *IOPMAssertionID) kernel.IOReturn
var _iOPMAssertionCreateErr error

func tryIOPMAssertionCreate(AssertionType corefoundation.CFStringRef, AssertionLevel IOPMAssertionLevel, AssertionID *IOPMAssertionID) (kernel.IOReturn, error) {
	if _iOPMAssertionCreate == nil {
		return *new(kernel.IOReturn), symbolCallError("IOPMAssertionCreate", "10.5", _iOPMAssertionCreateErr)
	}
	return _iOPMAssertionCreate(AssertionType, AssertionLevel, AssertionID), nil
}

// IOPMAssertionCreate dynamically requests a system behavior from the power management system.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/iokit/1557118-iopmassertioncreate
func IOPMAssertionCreate(AssertionType corefoundation.CFStringRef, AssertionLevel IOPMAssertionLevel, AssertionID *IOPMAssertionID) kernel.IOReturn {
	result, callErr := tryIOPMAssertionCreate(AssertionType, AssertionLevel, AssertionID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMAssertionCreateWithDescription func(AssertionType corefoundation.CFStringRef, Name corefoundation.CFStringRef, Details corefoundation.CFStringRef, HumanReadableReason corefoundation.CFStringRef, LocalizationBundlePath corefoundation.CFStringRef, Timeout corefoundation.CFTimeInterval, TimeoutAction corefoundation.CFStringRef, AssertionID *IOPMAssertionID) kernel.IOReturn
var _iOPMAssertionCreateWithDescriptionErr error

func tryIOPMAssertionCreateWithDescription(AssertionType corefoundation.CFStringRef, Name corefoundation.CFStringRef, Details corefoundation.CFStringRef, HumanReadableReason corefoundation.CFStringRef, LocalizationBundlePath corefoundation.CFStringRef, Timeout corefoundation.CFTimeInterval, TimeoutAction corefoundation.CFStringRef, AssertionID *IOPMAssertionID) (kernel.IOReturn, error) {
	if _iOPMAssertionCreateWithDescription == nil {
		return *new(kernel.IOReturn), symbolCallError("IOPMAssertionCreateWithDescription", "10.7", _iOPMAssertionCreateWithDescriptionErr)
	}
	return _iOPMAssertionCreateWithDescription(AssertionType, Name, Details, HumanReadableReason, LocalizationBundlePath, Timeout, TimeoutAction, AssertionID), nil
}

// IOPMAssertionCreateWithDescription.
//
// See: https://developer.apple.com/documentation/iokit/1557078-iopmassertioncreatewithdescripti
func IOPMAssertionCreateWithDescription(AssertionType corefoundation.CFStringRef, Name corefoundation.CFStringRef, Details corefoundation.CFStringRef, HumanReadableReason corefoundation.CFStringRef, LocalizationBundlePath corefoundation.CFStringRef, Timeout corefoundation.CFTimeInterval, TimeoutAction corefoundation.CFStringRef, AssertionID *IOPMAssertionID) kernel.IOReturn {
	result, callErr := tryIOPMAssertionCreateWithDescription(AssertionType, Name, Details, HumanReadableReason, LocalizationBundlePath, Timeout, TimeoutAction, AssertionID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMAssertionCreateWithName func(AssertionType corefoundation.CFStringRef, AssertionLevel IOPMAssertionLevel, AssertionName corefoundation.CFStringRef, AssertionID *IOPMAssertionID) kernel.IOReturn
var _iOPMAssertionCreateWithNameErr error

func tryIOPMAssertionCreateWithName(AssertionType corefoundation.CFStringRef, AssertionLevel IOPMAssertionLevel, AssertionName corefoundation.CFStringRef, AssertionID *IOPMAssertionID) (kernel.IOReturn, error) {
	if _iOPMAssertionCreateWithName == nil {
		return *new(kernel.IOReturn), symbolCallError("IOPMAssertionCreateWithName", "10.6", _iOPMAssertionCreateWithNameErr)
	}
	return _iOPMAssertionCreateWithName(AssertionType, AssertionLevel, AssertionName, AssertionID), nil
}

// IOPMAssertionCreateWithName dynamically requests a system behavior from the power management system.
//
// See: https://developer.apple.com/documentation/iokit/1557134-iopmassertioncreatewithname
func IOPMAssertionCreateWithName(AssertionType corefoundation.CFStringRef, AssertionLevel IOPMAssertionLevel, AssertionName corefoundation.CFStringRef, AssertionID *IOPMAssertionID) kernel.IOReturn {
	result, callErr := tryIOPMAssertionCreateWithName(AssertionType, AssertionLevel, AssertionName, AssertionID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMAssertionCreateWithProperties func(AssertionProperties corefoundation.CFDictionaryRef, AssertionID *IOPMAssertionID) kernel.IOReturn
var _iOPMAssertionCreateWithPropertiesErr error

func tryIOPMAssertionCreateWithProperties(AssertionProperties corefoundation.CFDictionaryRef, AssertionID *IOPMAssertionID) (kernel.IOReturn, error) {
	if _iOPMAssertionCreateWithProperties == nil {
		return *new(kernel.IOReturn), symbolCallError("IOPMAssertionCreateWithProperties", "10.7", _iOPMAssertionCreateWithPropertiesErr)
	}
	return _iOPMAssertionCreateWithProperties(AssertionProperties, AssertionID), nil
}

// IOPMAssertionCreateWithProperties creates an IOPMAssertion with more flexibility than IOPMAssertionCreateWithDescription.
//
// See: https://developer.apple.com/documentation/iokit/1557082-iopmassertioncreatewithpropertie
func IOPMAssertionCreateWithProperties(AssertionProperties corefoundation.CFDictionaryRef, AssertionID *IOPMAssertionID) kernel.IOReturn {
	result, callErr := tryIOPMAssertionCreateWithProperties(AssertionProperties, AssertionID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMAssertionDeclareUserActivity func(AssertionName corefoundation.CFStringRef, userType IOPMUserActiveType, AssertionID *IOPMAssertionID) kernel.IOReturn
var _iOPMAssertionDeclareUserActivityErr error

func tryIOPMAssertionDeclareUserActivity(AssertionName corefoundation.CFStringRef, userType IOPMUserActiveType, AssertionID *IOPMAssertionID) (kernel.IOReturn, error) {
	if _iOPMAssertionDeclareUserActivity == nil {
		return *new(kernel.IOReturn), symbolCallError("IOPMAssertionDeclareUserActivity", "10.7", _iOPMAssertionDeclareUserActivityErr)
	}
	return _iOPMAssertionDeclareUserActivity(AssertionName, userType, AssertionID), nil
}

// IOPMAssertionDeclareUserActivity declares that the user is active on the system.
//
// See: https://developer.apple.com/documentation/iokit/1557127-iopmassertiondeclareuseractivity
func IOPMAssertionDeclareUserActivity(AssertionName corefoundation.CFStringRef, userType IOPMUserActiveType, AssertionID *IOPMAssertionID) kernel.IOReturn {
	result, callErr := tryIOPMAssertionDeclareUserActivity(AssertionName, userType, AssertionID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMAssertionRelease func(AssertionID IOPMAssertionID) kernel.IOReturn
var _iOPMAssertionReleaseErr error

func tryIOPMAssertionRelease(AssertionID IOPMAssertionID) (kernel.IOReturn, error) {
	if _iOPMAssertionRelease == nil {
		return *new(kernel.IOReturn), symbolCallError("IOPMAssertionRelease", "10.7", _iOPMAssertionReleaseErr)
	}
	return _iOPMAssertionRelease(AssertionID), nil
}

// IOPMAssertionRelease decrements the assertion's retain count.
//
// See: https://developer.apple.com/documentation/iokit/1557090-iopmassertionrelease
func IOPMAssertionRelease(AssertionID IOPMAssertionID) kernel.IOReturn {
	result, callErr := tryIOPMAssertionRelease(AssertionID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMAssertionRetain func(theAssertion IOPMAssertionID)
var _iOPMAssertionRetainErr error

func tryIOPMAssertionRetain(theAssertion IOPMAssertionID) error {
	if _iOPMAssertionRetain == nil {
		return symbolCallError("IOPMAssertionRetain", "10.7", _iOPMAssertionRetainErr)
	}
	_iOPMAssertionRetain(theAssertion)
	return nil
}

// IOPMAssertionRetain increments the assertion's retain count.
//
// See: https://developer.apple.com/documentation/iokit/1557071-iopmassertionretain
func IOPMAssertionRetain(theAssertion IOPMAssertionID) {
	if callErr := tryIOPMAssertionRetain(theAssertion); callErr != nil {
		panic(callErr)
	}
}

var _iOPMAssertionSetProperty func(theAssertion IOPMAssertionID, theProperty corefoundation.CFStringRef, theValue corefoundation.CFTypeRef) kernel.IOReturn
var _iOPMAssertionSetPropertyErr error

func tryIOPMAssertionSetProperty(theAssertion IOPMAssertionID, theProperty corefoundation.CFStringRef, theValue corefoundation.CFTypeRef) (kernel.IOReturn, error) {
	if _iOPMAssertionSetProperty == nil {
		return *new(kernel.IOReturn), symbolCallError("IOPMAssertionSetProperty", "10.7", _iOPMAssertionSetPropertyErr)
	}
	return _iOPMAssertionSetProperty(theAssertion, theProperty, theValue), nil
}

// IOPMAssertionSetProperty sets a property in the assertion.
//
// See: https://developer.apple.com/documentation/iokit/1557107-iopmassertionsetproperty
func IOPMAssertionSetProperty(theAssertion IOPMAssertionID, theProperty corefoundation.CFStringRef, theValue corefoundation.CFTypeRef) kernel.IOReturn {
	result, callErr := tryIOPMAssertionSetProperty(theAssertion, theProperty, theValue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMCancelScheduledPowerEvent func(time_to_wake corefoundation.CFDateRef, my_id corefoundation.CFStringRef, type_ corefoundation.CFStringRef) kernel.IOReturn
var _iOPMCancelScheduledPowerEventErr error

func tryIOPMCancelScheduledPowerEvent(time_to_wake corefoundation.CFDateRef, my_id corefoundation.CFStringRef, type_ corefoundation.CFStringRef) (kernel.IOReturn, error) {
	if _iOPMCancelScheduledPowerEvent == nil {
		return *new(kernel.IOReturn), symbolCallError("IOPMCancelScheduledPowerEvent", "10.3", _iOPMCancelScheduledPowerEventErr)
	}
	return _iOPMCancelScheduledPowerEvent(time_to_wake, my_id, type_), nil
}

// IOPMCancelScheduledPowerEvent cancel a previously scheduled power event.
//
// See: https://developer.apple.com/documentation/iokit/1557116-iopmcancelscheduledpowerevent
func IOPMCancelScheduledPowerEvent(time_to_wake corefoundation.CFDateRef, my_id corefoundation.CFStringRef, type_ corefoundation.CFStringRef) kernel.IOReturn {
	result, callErr := tryIOPMCancelScheduledPowerEvent(time_to_wake, my_id, type_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMCopyAssertionsByProcess func(AssertionsByPID *corefoundation.CFDictionaryRef) kernel.IOReturn
var _iOPMCopyAssertionsByProcessErr error

func tryIOPMCopyAssertionsByProcess(AssertionsByPID *corefoundation.CFDictionaryRef) (kernel.IOReturn, error) {
	if _iOPMCopyAssertionsByProcess == nil {
		return *new(kernel.IOReturn), symbolCallError("IOPMCopyAssertionsByProcess", "10.7", _iOPMCopyAssertionsByProcessErr)
	}
	return _iOPMCopyAssertionsByProcess(AssertionsByPID), nil
}

// IOPMCopyAssertionsByProcess returns a dictionary listing all assertions, grouped by their owning process.
//
// See: https://developer.apple.com/documentation/iokit/1557130-iopmcopyassertionsbyprocess
func IOPMCopyAssertionsByProcess(AssertionsByPID *corefoundation.CFDictionaryRef) kernel.IOReturn {
	result, callErr := tryIOPMCopyAssertionsByProcess(AssertionsByPID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMCopyAssertionsStatus func(AssertionsStatus *corefoundation.CFDictionaryRef) kernel.IOReturn
var _iOPMCopyAssertionsStatusErr error

func tryIOPMCopyAssertionsStatus(AssertionsStatus *corefoundation.CFDictionaryRef) (kernel.IOReturn, error) {
	if _iOPMCopyAssertionsStatus == nil {
		return *new(kernel.IOReturn), symbolCallError("IOPMCopyAssertionsStatus", "10.7", _iOPMCopyAssertionsStatusErr)
	}
	return _iOPMCopyAssertionsStatus(AssertionsStatus), nil
}

// IOPMCopyAssertionsStatus returns a list of available assertions and their system-wide levels.
//
// See: https://developer.apple.com/documentation/iokit/1557072-iopmcopyassertionsstatus
func IOPMCopyAssertionsStatus(AssertionsStatus *corefoundation.CFDictionaryRef) kernel.IOReturn {
	result, callErr := tryIOPMCopyAssertionsStatus(AssertionsStatus)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMCopyBatteryInfo func(masterPort uint32, info *corefoundation.CFArrayRef) kernel.IOReturn
var _iOPMCopyBatteryInfoErr error

func tryIOPMCopyBatteryInfo(masterPort uint32, info *corefoundation.CFArrayRef) (kernel.IOReturn, error) {
	if _iOPMCopyBatteryInfo == nil {
		return *new(kernel.IOReturn), symbolCallError("IOPMCopyBatteryInfo", "10.0", _iOPMCopyBatteryInfoErr)
	}
	return _iOPMCopyBatteryInfo(masterPort, info), nil
}

// IOPMCopyBatteryInfo request raw battery data from the system.
//
// See: https://developer.apple.com/documentation/iokit/1557138-iopmcopybatteryinfo
func IOPMCopyBatteryInfo(masterPort uint32, info *corefoundation.CFArrayRef) kernel.IOReturn {
	result, callErr := tryIOPMCopyBatteryInfo(masterPort, info)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMCopyCPUPowerStatus func(cpuPowerStatus *corefoundation.CFDictionaryRef) kernel.IOReturn
var _iOPMCopyCPUPowerStatusErr error

func tryIOPMCopyCPUPowerStatus(cpuPowerStatus *corefoundation.CFDictionaryRef) (kernel.IOReturn, error) {
	if _iOPMCopyCPUPowerStatus == nil {
		return *new(kernel.IOReturn), symbolCallError("IOPMCopyCPUPowerStatus", "10.7", _iOPMCopyCPUPowerStatusErr)
	}
	return _iOPMCopyCPUPowerStatus(cpuPowerStatus), nil
}

// IOPMCopyCPUPowerStatus copy status of all current CPU power levels.
//
// See: https://developer.apple.com/documentation/iokit/1557079-iopmcopycpupowerstatus
func IOPMCopyCPUPowerStatus(cpuPowerStatus *corefoundation.CFDictionaryRef) kernel.IOReturn {
	result, callErr := tryIOPMCopyCPUPowerStatus(cpuPowerStatus)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMCopyScheduledPowerEvents func() corefoundation.CFArrayRef
var _iOPMCopyScheduledPowerEventsErr error

func tryIOPMCopyScheduledPowerEvents() (corefoundation.CFArrayRef, error) {
	if _iOPMCopyScheduledPowerEvents == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("IOPMCopyScheduledPowerEvents", "10.3", _iOPMCopyScheduledPowerEventsErr)
	}
	return _iOPMCopyScheduledPowerEvents(), nil
}

// IOPMCopyScheduledPowerEvents list all scheduled system power events
//
// See: https://developer.apple.com/documentation/iokit/1557109-iopmcopyscheduledpowerevents
func IOPMCopyScheduledPowerEvents() corefoundation.CFArrayRef {
	result, callErr := tryIOPMCopyScheduledPowerEvents()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMDeclareNetworkClientActivity func(AssertionName corefoundation.CFStringRef, AssertionID *IOPMAssertionID) kernel.IOReturn
var _iOPMDeclareNetworkClientActivityErr error

func tryIOPMDeclareNetworkClientActivity(AssertionName corefoundation.CFStringRef, AssertionID *IOPMAssertionID) (kernel.IOReturn, error) {
	if _iOPMDeclareNetworkClientActivity == nil {
		return *new(kernel.IOReturn), symbolCallError("IOPMDeclareNetworkClientActivity", "10.9", _iOPMDeclareNetworkClientActivityErr)
	}
	return _iOPMDeclareNetworkClientActivity(AssertionName, AssertionID), nil
}

// IOPMDeclareNetworkClientActivity.
//
// See: https://developer.apple.com/documentation/iokit/1557137-iopmdeclarenetworkclientactivity
func IOPMDeclareNetworkClientActivity(AssertionName corefoundation.CFStringRef, AssertionID *IOPMAssertionID) kernel.IOReturn {
	result, callErr := tryIOPMDeclareNetworkClientActivity(AssertionName, AssertionID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMFindPowerManagement func(master_device_port uint32) uint32
var _iOPMFindPowerManagementErr error

func tryIOPMFindPowerManagement(master_device_port uint32) (uint32, error) {
	if _iOPMFindPowerManagement == nil {
		return 0, symbolCallError("IOPMFindPowerManagement", "10.0", _iOPMFindPowerManagementErr)
	}
	return _iOPMFindPowerManagement(master_device_port), nil
}

// IOPMFindPowerManagement finds the Root Power Domain IOService.
//
// See: https://developer.apple.com/documentation/iokit/1557133-iopmfindpowermanagement
func IOPMFindPowerManagement(master_device_port uint32) uint32 {
	result, callErr := tryIOPMFindPowerManagement(master_device_port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMGetAggressiveness func(fb uint32, type_ uint, aggressiveness *uint) kernel.IOReturn
var _iOPMGetAggressivenessErr error

func tryIOPMGetAggressiveness(fb uint32, type_ uint, aggressiveness *uint) (kernel.IOReturn, error) {
	if _iOPMGetAggressiveness == nil {
		return *new(kernel.IOReturn), symbolCallError("IOPMGetAggressiveness", "10.0", _iOPMGetAggressivenessErr)
	}
	return _iOPMGetAggressiveness(fb, type_, aggressiveness), nil
}

// IOPMGetAggressiveness retrieves the current value of one of the aggressiveness factors in IOKit Power Management.
//
// See: https://developer.apple.com/documentation/iokit/1557117-iopmgetaggressiveness
func IOPMGetAggressiveness(fb uint32, type_ uint, aggressiveness *uint) kernel.IOReturn {
	result, callErr := tryIOPMGetAggressiveness(fb, type_, aggressiveness)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMGetThermalWarningLevel func(thermalLevel *uint32) kernel.IOReturn
var _iOPMGetThermalWarningLevelErr error

func tryIOPMGetThermalWarningLevel(thermalLevel *uint32) (kernel.IOReturn, error) {
	if _iOPMGetThermalWarningLevel == nil {
		return *new(kernel.IOReturn), symbolCallError("IOPMGetThermalWarningLevel", "10.7", _iOPMGetThermalWarningLevelErr)
	}
	return _iOPMGetThermalWarningLevel(thermalLevel), nil
}

// IOPMGetThermalWarningLevel get thermal warning level of the system.
//
// See: https://developer.apple.com/documentation/iokit/1557103-iopmgetthermalwarninglevel
func IOPMGetThermalWarningLevel(thermalLevel *uint32) kernel.IOReturn {
	result, callErr := tryIOPMGetThermalWarningLevel(thermalLevel)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMSchedulePowerEvent func(time_to_wake corefoundation.CFDateRef, my_id corefoundation.CFStringRef, type_ corefoundation.CFStringRef) kernel.IOReturn
var _iOPMSchedulePowerEventErr error

func tryIOPMSchedulePowerEvent(time_to_wake corefoundation.CFDateRef, my_id corefoundation.CFStringRef, type_ corefoundation.CFStringRef) (kernel.IOReturn, error) {
	if _iOPMSchedulePowerEvent == nil {
		return *new(kernel.IOReturn), symbolCallError("IOPMSchedulePowerEvent", "10.3", _iOPMSchedulePowerEventErr)
	}
	return _iOPMSchedulePowerEvent(time_to_wake, my_id, type_), nil
}

// IOPMSchedulePowerEvent schedule the machine to wake from sleep, power on, go to sleep, or shutdown.
//
// See: https://developer.apple.com/documentation/iokit/1557076-iopmschedulepowerevent
func IOPMSchedulePowerEvent(time_to_wake corefoundation.CFDateRef, my_id corefoundation.CFStringRef, type_ corefoundation.CFStringRef) kernel.IOReturn {
	result, callErr := tryIOPMSchedulePowerEvent(time_to_wake, my_id, type_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMSetAggressiveness func(fb uint32, type_ uint, aggressiveness uint) kernel.IOReturn
var _iOPMSetAggressivenessErr error

func tryIOPMSetAggressiveness(fb uint32, type_ uint, aggressiveness uint) (kernel.IOReturn, error) {
	if _iOPMSetAggressiveness == nil {
		return *new(kernel.IOReturn), symbolCallError("IOPMSetAggressiveness", "10.0", _iOPMSetAggressivenessErr)
	}
	return _iOPMSetAggressiveness(fb, type_, aggressiveness), nil
}

// IOPMSetAggressiveness sets one of the aggressiveness factors in IOKit Power Management.
//
// See: https://developer.apple.com/documentation/iokit/1557098-iopmsetaggressiveness
func IOPMSetAggressiveness(fb uint32, type_ uint, aggressiveness uint) kernel.IOReturn {
	result, callErr := tryIOPMSetAggressiveness(fb, type_, aggressiveness)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMSleepEnabled func() int32
var _iOPMSleepEnabledErr error

func tryIOPMSleepEnabled() (int32, error) {
	if _iOPMSleepEnabled == nil {
		return 0, symbolCallError("IOPMSleepEnabled", "10.0", _iOPMSleepEnabledErr)
	}
	return _iOPMSleepEnabled(), nil
}

// IOPMSleepEnabled tells whether the system supports full sleep, or just doze
//
// See: https://developer.apple.com/documentation/iokit/1557074-iopmsleepenabled
func IOPMSleepEnabled() int32 {
	result, callErr := tryIOPMSleepEnabled()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPMSleepSystem func(fb uint32) kernel.IOReturn
var _iOPMSleepSystemErr error

func tryIOPMSleepSystem(fb uint32) (kernel.IOReturn, error) {
	if _iOPMSleepSystem == nil {
		return *new(kernel.IOReturn), symbolCallError("IOPMSleepSystem", "10.0", _iOPMSleepSystemErr)
	}
	return _iOPMSleepSystem(fb), nil
}

// IOPMSleepSystem request that the system initiate sleep.
//
// See: https://developer.apple.com/documentation/iokit/1557121-iopmsleepsystem
func IOPMSleepSystem(fb uint32) kernel.IOReturn {
	result, callErr := tryIOPMSleepSystem(fb)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPSCopyExternalPowerAdapterDetails func() corefoundation.CFDictionaryRef
var _iOPSCopyExternalPowerAdapterDetailsErr error

func tryIOPSCopyExternalPowerAdapterDetails() (corefoundation.CFDictionaryRef, error) {
	if _iOPSCopyExternalPowerAdapterDetails == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("IOPSCopyExternalPowerAdapterDetails", "10.6", _iOPSCopyExternalPowerAdapterDetailsErr)
	}
	return _iOPSCopyExternalPowerAdapterDetails(), nil
}

// IOPSCopyExternalPowerAdapterDetails returns a CFDictionary that describes the attached (AC) external power adapter (if any external power adapter is attached.
//
// See: https://developer.apple.com/documentation/iokit/1523866-iopscopyexternalpoweradapterdeta
func IOPSCopyExternalPowerAdapterDetails() corefoundation.CFDictionaryRef {
	result, callErr := tryIOPSCopyExternalPowerAdapterDetails()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPSCopyPowerSourcesInfo func() corefoundation.CFTypeRef
var _iOPSCopyPowerSourcesInfoErr error

func tryIOPSCopyPowerSourcesInfo() (corefoundation.CFTypeRef, error) {
	if _iOPSCopyPowerSourcesInfo == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("IOPSCopyPowerSourcesInfo", "10.2", _iOPSCopyPowerSourcesInfoErr)
	}
	return _iOPSCopyPowerSourcesInfo(), nil
}

// IOPSCopyPowerSourcesInfo returns a blob of Power Source information in an opaque CFTypeRef.
//
// See: https://developer.apple.com/documentation/iokit/1523839-iopscopypowersourcesinfo
func IOPSCopyPowerSourcesInfo() corefoundation.CFTypeRef {
	result, callErr := tryIOPSCopyPowerSourcesInfo()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPSCopyPowerSourcesList func(blob corefoundation.CFTypeRef) corefoundation.CFArrayRef
var _iOPSCopyPowerSourcesListErr error

func tryIOPSCopyPowerSourcesList(blob corefoundation.CFTypeRef) (corefoundation.CFArrayRef, error) {
	if _iOPSCopyPowerSourcesList == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("IOPSCopyPowerSourcesList", "10.2", _iOPSCopyPowerSourcesListErr)
	}
	return _iOPSCopyPowerSourcesList(blob), nil
}

// IOPSCopyPowerSourcesList returns a CFArray of Power Source handles, each of type CFTypeRef.
//
// See: https://developer.apple.com/documentation/iokit/1523856-iopscopypowersourceslist
func IOPSCopyPowerSourcesList(blob corefoundation.CFTypeRef) corefoundation.CFArrayRef {
	result, callErr := tryIOPSCopyPowerSourcesList(blob)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPSCreateLimitedPowerNotification func(callback unsafe.Pointer, context unsafe.Pointer) corefoundation.CFRunLoopSourceRef
var _iOPSCreateLimitedPowerNotificationErr error

func tryIOPSCreateLimitedPowerNotification(callback unsafe.Pointer, context unsafe.Pointer) (corefoundation.CFRunLoopSourceRef, error) {
	if _iOPSCreateLimitedPowerNotification == nil {
		return *new(corefoundation.CFRunLoopSourceRef), symbolCallError("IOPSCreateLimitedPowerNotification", "10.9", _iOPSCreateLimitedPowerNotificationErr)
	}
	return _iOPSCreateLimitedPowerNotification(callback, context), nil
}

// IOPSCreateLimitedPowerNotification.
//
// See: https://developer.apple.com/documentation/iokit/1523865-iopscreatelimitedpowernotificati
func IOPSCreateLimitedPowerNotification(callback unsafe.Pointer, context unsafe.Pointer) corefoundation.CFRunLoopSourceRef {
	result, callErr := tryIOPSCreateLimitedPowerNotification(callback, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPSGetBatteryWarningLevel func() IOPSLowBatteryWarningLevel
var _iOPSGetBatteryWarningLevelErr error

func tryIOPSGetBatteryWarningLevel() (IOPSLowBatteryWarningLevel, error) {
	if _iOPSGetBatteryWarningLevel == nil {
		return *new(IOPSLowBatteryWarningLevel), symbolCallError("IOPSGetBatteryWarningLevel", "10.6", _iOPSGetBatteryWarningLevelErr)
	}
	return _iOPSGetBatteryWarningLevel(), nil
}

// IOPSGetBatteryWarningLevel indicates whether the system is at a low battery warning level.
//
// See: https://developer.apple.com/documentation/iokit/1523851-iopsgetbatterywarninglevel
func IOPSGetBatteryWarningLevel() IOPSLowBatteryWarningLevel {
	result, callErr := tryIOPSGetBatteryWarningLevel()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPSGetPowerSourceDescription func(blob corefoundation.CFTypeRef, ps corefoundation.CFTypeRef) corefoundation.CFDictionaryRef
var _iOPSGetPowerSourceDescriptionErr error

func tryIOPSGetPowerSourceDescription(blob corefoundation.CFTypeRef, ps corefoundation.CFTypeRef) (corefoundation.CFDictionaryRef, error) {
	if _iOPSGetPowerSourceDescription == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("IOPSGetPowerSourceDescription", "10.2", _iOPSGetPowerSourceDescriptionErr)
	}
	return _iOPSGetPowerSourceDescription(blob, ps), nil
}

// IOPSGetPowerSourceDescription returns a CFDictionary with readable information about the specific power source.
//
// See: https://developer.apple.com/documentation/iokit/1523867-iopsgetpowersourcedescription
func IOPSGetPowerSourceDescription(blob corefoundation.CFTypeRef, ps corefoundation.CFTypeRef) corefoundation.CFDictionaryRef {
	result, callErr := tryIOPSGetPowerSourceDescription(blob, ps)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPSGetProvidingPowerSourceType func(snapshot corefoundation.CFTypeRef) corefoundation.CFStringRef
var _iOPSGetProvidingPowerSourceTypeErr error

func tryIOPSGetProvidingPowerSourceType(snapshot corefoundation.CFTypeRef) (corefoundation.CFStringRef, error) {
	if _iOPSGetProvidingPowerSourceType == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("IOPSGetProvidingPowerSourceType", "10.7", _iOPSGetProvidingPowerSourceTypeErr)
	}
	return _iOPSGetProvidingPowerSourceType(snapshot), nil
}

// IOPSGetProvidingPowerSourceType.
//
// See: https://developer.apple.com/documentation/iokit/1523858-iopsgetprovidingpowersourcetype
func IOPSGetProvidingPowerSourceType(snapshot corefoundation.CFTypeRef) corefoundation.CFStringRef {
	result, callErr := tryIOPSGetProvidingPowerSourceType(snapshot)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPSGetTimeRemainingEstimate func() corefoundation.CFTimeInterval
var _iOPSGetTimeRemainingEstimateErr error

func tryIOPSGetTimeRemainingEstimate() (corefoundation.CFTimeInterval, error) {
	if _iOPSGetTimeRemainingEstimate == nil {
		return *new(corefoundation.CFTimeInterval), symbolCallError("IOPSGetTimeRemainingEstimate", "10.7", _iOPSGetTimeRemainingEstimateErr)
	}
	return _iOPSGetTimeRemainingEstimate(), nil
}

// IOPSGetTimeRemainingEstimate returns the estimated minutes remaining until all power sources (battery and/or UPS's) are empty, or returns kIOPSTimeRemainingUnlimited if attached to an unlimited power source.
//
// See: https://developer.apple.com/documentation/iokit/1523835-iopsgettimeremainingestimate
func IOPSGetTimeRemainingEstimate() corefoundation.CFTimeInterval {
	result, callErr := tryIOPSGetTimeRemainingEstimate()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOPSNotificationCreateRunLoopSource func(callback unsafe.Pointer, context unsafe.Pointer) corefoundation.CFRunLoopSourceRef
var _iOPSNotificationCreateRunLoopSourceErr error

func tryIOPSNotificationCreateRunLoopSource(callback unsafe.Pointer, context unsafe.Pointer) (corefoundation.CFRunLoopSourceRef, error) {
	if _iOPSNotificationCreateRunLoopSource == nil {
		return *new(corefoundation.CFRunLoopSourceRef), symbolCallError("IOPSNotificationCreateRunLoopSource", "10.2", _iOPSNotificationCreateRunLoopSourceErr)
	}
	return _iOPSNotificationCreateRunLoopSource(callback, context), nil
}

// IOPSNotificationCreateRunLoopSource returns a CFRunLoopSourceRef that notifies the caller when power source information changes.
//
// See: https://developer.apple.com/documentation/iokit/1523868-iopsnotificationcreaterunloopsou
func IOPSNotificationCreateRunLoopSource(callback unsafe.Pointer, context unsafe.Pointer) corefoundation.CFRunLoopSourceRef {
	result, callErr := tryIOPSNotificationCreateRunLoopSource(callback, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegisterApp func(refcon uintptr, theDriver uint32, thePortRef *IONotificationPortRef, callback unsafe.Pointer, notifier *uint32) uint32
var _iORegisterAppErr error

func tryIORegisterApp(refcon uintptr, theDriver uint32, thePortRef *IONotificationPortRef, callback IOServiceInterestCallback, notifier *uint32) (uint32, error) {
	if _iORegisterApp == nil {
		return 0, symbolCallError("IORegisterApp", "10.0", _iORegisterAppErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 unsafe.Pointer, blockArg1 uint32, blockArg2 uint32, blockArg3 unsafe.Pointer) {
		callback(blockArg0, blockArg1, blockArg2, blockArg3)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _iORegisterApp(refcon, theDriver, thePortRef, _block0, notifier), nil
}

// IORegisterApp connects the caller to an IOService for the purpose of receiving power state change notifications for the device controlled by the IOService.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/iokit/1557102-ioregisterapp
func IORegisterApp(refcon uintptr, theDriver uint32, thePortRef *IONotificationPortRef, callback IOServiceInterestCallback, notifier *uint32) uint32 {
	result, callErr := tryIORegisterApp(refcon, theDriver, thePortRef, callback, notifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegisterForSystemPower func(refcon uintptr, thePortRef *IONotificationPortRef, callback unsafe.Pointer, notifier *uint32) uint32
var _iORegisterForSystemPowerErr error

func tryIORegisterForSystemPower(refcon uintptr, thePortRef *IONotificationPortRef, callback IOServiceInterestCallback, notifier *uint32) (uint32, error) {
	if _iORegisterForSystemPower == nil {
		return 0, symbolCallError("IORegisterForSystemPower", "10.0", _iORegisterForSystemPowerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 unsafe.Pointer, blockArg1 uint32, blockArg2 uint32, blockArg3 unsafe.Pointer) {
		callback(blockArg0, blockArg1, blockArg2, blockArg3)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _iORegisterForSystemPower(refcon, thePortRef, _block0, notifier), nil
}

// IORegisterForSystemPower connects the caller to the Root Power Domain IOService for the purpose of receiving sleep & wake notifications for the system.
//
// See: https://developer.apple.com/documentation/iokit/1557114-ioregisterforsystempower
func IORegisterForSystemPower(refcon uintptr, thePortRef *IONotificationPortRef, callback IOServiceInterestCallback, notifier *uint32) uint32 {
	result, callErr := tryIORegisterForSystemPower(refcon, thePortRef, callback, notifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryCreateIterator func(mainPort uint32, plane string, options uint32, iterator *uint32) int32
var _iORegistryCreateIteratorErr error

func tryIORegistryCreateIterator(mainPort uint32, plane string, options uint32, iterator *uint32) (int32, error) {
	if _iORegistryCreateIterator == nil {
		return 0, symbolCallError("IORegistryCreateIterator", "10.0", _iORegistryCreateIteratorErr)
	}
	return _iORegistryCreateIterator(mainPort, plane, options, iterator), nil
}

// IORegistryCreateIterator create an iterator rooted at the registry root.
//
// See: https://developer.apple.com/documentation/iokit/1514238-ioregistrycreateiterator
func IORegistryCreateIterator(mainPort uint32, plane string, options uint32, iterator *uint32) int32 {
	result, callErr := tryIORegistryCreateIterator(mainPort, plane, options, iterator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntryCopyFromPath func(mainPort uint32, path corefoundation.CFStringRef) uint32
var _iORegistryEntryCopyFromPathErr error

func tryIORegistryEntryCopyFromPath(mainPort uint32, path corefoundation.CFStringRef) (uint32, error) {
	if _iORegistryEntryCopyFromPath == nil {
		return 0, symbolCallError("IORegistryEntryCopyFromPath", "10.11", _iORegistryEntryCopyFromPathErr)
	}
	return _iORegistryEntryCopyFromPath(mainPort, path), nil
}

// IORegistryEntryCopyFromPath.
//
// See: https://developer.apple.com/documentation/iokit/1514248-ioregistryentrycopyfrompath
func IORegistryEntryCopyFromPath(mainPort uint32, path corefoundation.CFStringRef) uint32 {
	result, callErr := tryIORegistryEntryCopyFromPath(mainPort, path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntryCopyPath func(entry uint32, plane string) corefoundation.CFStringRef
var _iORegistryEntryCopyPathErr error

func tryIORegistryEntryCopyPath(entry uint32, plane string) (corefoundation.CFStringRef, error) {
	if _iORegistryEntryCopyPath == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("IORegistryEntryCopyPath", "10.11", _iORegistryEntryCopyPathErr)
	}
	return _iORegistryEntryCopyPath(entry, plane), nil
}

// IORegistryEntryCopyPath.
//
// See: https://developer.apple.com/documentation/iokit/1514853-ioregistryentrycopypath
func IORegistryEntryCopyPath(entry uint32, plane string) corefoundation.CFStringRef {
	result, callErr := tryIORegistryEntryCopyPath(entry, plane)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntryCreateCFProperties func(entry uint32, properties *corefoundation.CFMutableDictionaryRef, allocator corefoundation.CFAllocatorRef, options uint32) int32
var _iORegistryEntryCreateCFPropertiesErr error

func tryIORegistryEntryCreateCFProperties(entry uint32, properties *corefoundation.CFMutableDictionaryRef, allocator corefoundation.CFAllocatorRef, options uint32) (int32, error) {
	if _iORegistryEntryCreateCFProperties == nil {
		return 0, symbolCallError("IORegistryEntryCreateCFProperties", "10.0", _iORegistryEntryCreateCFPropertiesErr)
	}
	return _iORegistryEntryCreateCFProperties(entry, properties, allocator, options), nil
}

// IORegistryEntryCreateCFProperties create a CF dictionary representation of a registry entry's property table.
//
// See: https://developer.apple.com/documentation/iokit/1514310-ioregistryentrycreatecfpropertie
func IORegistryEntryCreateCFProperties(entry uint32, properties *corefoundation.CFMutableDictionaryRef, allocator corefoundation.CFAllocatorRef, options uint32) int32 {
	result, callErr := tryIORegistryEntryCreateCFProperties(entry, properties, allocator, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntryCreateCFProperty func(entry uint32, key corefoundation.CFStringRef, allocator corefoundation.CFAllocatorRef, options uint32) corefoundation.CFTypeRef
var _iORegistryEntryCreateCFPropertyErr error

func tryIORegistryEntryCreateCFProperty(entry uint32, key corefoundation.CFStringRef, allocator corefoundation.CFAllocatorRef, options uint32) (corefoundation.CFTypeRef, error) {
	if _iORegistryEntryCreateCFProperty == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("IORegistryEntryCreateCFProperty", "10.0", _iORegistryEntryCreateCFPropertyErr)
	}
	return _iORegistryEntryCreateCFProperty(entry, key, allocator, options), nil
}

// IORegistryEntryCreateCFProperty create a CF representation of a registry entry's property.
//
// See: https://developer.apple.com/documentation/iokit/1514293-ioregistryentrycreatecfproperty
func IORegistryEntryCreateCFProperty(entry uint32, key corefoundation.CFStringRef, allocator corefoundation.CFAllocatorRef, options uint32) corefoundation.CFTypeRef {
	result, callErr := tryIORegistryEntryCreateCFProperty(entry, key, allocator, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntryCreateIterator func(entry uint32, plane string, options uint32, iterator *uint32) int32
var _iORegistryEntryCreateIteratorErr error

func tryIORegistryEntryCreateIterator(entry uint32, plane string, options uint32, iterator *uint32) (int32, error) {
	if _iORegistryEntryCreateIterator == nil {
		return 0, symbolCallError("IORegistryEntryCreateIterator", "10.0", _iORegistryEntryCreateIteratorErr)
	}
	return _iORegistryEntryCreateIterator(entry, plane, options, iterator), nil
}

// IORegistryEntryCreateIterator create an iterator rooted at a given registry entry.
//
// See: https://developer.apple.com/documentation/iokit/1514318-ioregistryentrycreateiterator
func IORegistryEntryCreateIterator(entry uint32, plane string, options uint32, iterator *uint32) int32 {
	result, callErr := tryIORegistryEntryCreateIterator(entry, plane, options, iterator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntryFromPath func(mainPort uint32, path string) uint32
var _iORegistryEntryFromPathErr error

func tryIORegistryEntryFromPath(mainPort uint32, path string) (uint32, error) {
	if _iORegistryEntryFromPath == nil {
		return 0, symbolCallError("IORegistryEntryFromPath", "10.0", _iORegistryEntryFromPathErr)
	}
	return _iORegistryEntryFromPath(mainPort, path), nil
}

// IORegistryEntryFromPath looks up a registry entry by path.
//
// See: https://developer.apple.com/documentation/iokit/1514802-ioregistryentryfrompath
func IORegistryEntryFromPath(mainPort uint32, path string) uint32 {
	result, callErr := tryIORegistryEntryFromPath(mainPort, path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntryGetChildEntry func(entry uint32, plane string, child *uint32) int32
var _iORegistryEntryGetChildEntryErr error

func tryIORegistryEntryGetChildEntry(entry uint32, plane string, child *uint32) (int32, error) {
	if _iORegistryEntryGetChildEntry == nil {
		return 0, symbolCallError("IORegistryEntryGetChildEntry", "10.0", _iORegistryEntryGetChildEntryErr)
	}
	return _iORegistryEntryGetChildEntry(entry, plane, child), nil
}

// IORegistryEntryGetChildEntry returns the first child of a registry entry in a plane.
//
// See: https://developer.apple.com/documentation/iokit/1514496-ioregistryentrygetchildentry
func IORegistryEntryGetChildEntry(entry uint32, plane string, child *uint32) int32 {
	result, callErr := tryIORegistryEntryGetChildEntry(entry, plane, child)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntryGetChildIterator func(entry uint32, plane string, iterator *uint32) int32
var _iORegistryEntryGetChildIteratorErr error

func tryIORegistryEntryGetChildIterator(entry uint32, plane string, iterator *uint32) (int32, error) {
	if _iORegistryEntryGetChildIterator == nil {
		return 0, symbolCallError("IORegistryEntryGetChildIterator", "10.0", _iORegistryEntryGetChildIteratorErr)
	}
	return _iORegistryEntryGetChildIterator(entry, plane, iterator), nil
}

// IORegistryEntryGetChildIterator returns an iterator over a registry entry’s child entries in a plane.
//
// See: https://developer.apple.com/documentation/iokit/1514703-ioregistryentrygetchilditerator
func IORegistryEntryGetChildIterator(entry uint32, plane string, iterator *uint32) int32 {
	result, callErr := tryIORegistryEntryGetChildIterator(entry, plane, iterator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntryGetLocationInPlane func(entry uint32, plane string, location string) int32
var _iORegistryEntryGetLocationInPlaneErr error

func tryIORegistryEntryGetLocationInPlane(entry uint32, plane string, location string) (int32, error) {
	if _iORegistryEntryGetLocationInPlane == nil {
		return 0, symbolCallError("IORegistryEntryGetLocationInPlane", "10.1", _iORegistryEntryGetLocationInPlaneErr)
	}
	return _iORegistryEntryGetLocationInPlane(entry, plane, location), nil
}

// IORegistryEntryGetLocationInPlane returns a C-string location assigned to a registry entry, in a specified plane.
//
// See: https://developer.apple.com/documentation/iokit/1514340-ioregistryentrygetlocationinplan
func IORegistryEntryGetLocationInPlane(entry uint32, plane string, location string) int32 {
	result, callErr := tryIORegistryEntryGetLocationInPlane(entry, plane, location)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntryGetName func(entry uint32, name string) int32
var _iORegistryEntryGetNameErr error

func tryIORegistryEntryGetName(entry uint32, name string) (int32, error) {
	if _iORegistryEntryGetName == nil {
		return 0, symbolCallError("IORegistryEntryGetName", "10.0", _iORegistryEntryGetNameErr)
	}
	return _iORegistryEntryGetName(entry, name), nil
}

// IORegistryEntryGetName returns a C-string name assigned to a registry entry.
//
// See: https://developer.apple.com/documentation/iokit/1514323-ioregistryentrygetname
func IORegistryEntryGetName(entry uint32, name string) int32 {
	result, callErr := tryIORegistryEntryGetName(entry, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntryGetNameInPlane func(entry uint32, plane string, name string) int32
var _iORegistryEntryGetNameInPlaneErr error

func tryIORegistryEntryGetNameInPlane(entry uint32, plane string, name string) (int32, error) {
	if _iORegistryEntryGetNameInPlane == nil {
		return 0, symbolCallError("IORegistryEntryGetNameInPlane", "10.0", _iORegistryEntryGetNameInPlaneErr)
	}
	return _iORegistryEntryGetNameInPlane(entry, plane, name), nil
}

// IORegistryEntryGetNameInPlane returns a C-string name assigned to a registry entry, in a specified plane.
//
// See: https://developer.apple.com/documentation/iokit/1514475-ioregistryentrygetnameinplane
func IORegistryEntryGetNameInPlane(entry uint32, plane string, name string) int32 {
	result, callErr := tryIORegistryEntryGetNameInPlane(entry, plane, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntryGetParentEntry func(entry uint32, plane string, parent *uint32) int32
var _iORegistryEntryGetParentEntryErr error

func tryIORegistryEntryGetParentEntry(entry uint32, plane string, parent *uint32) (int32, error) {
	if _iORegistryEntryGetParentEntry == nil {
		return 0, symbolCallError("IORegistryEntryGetParentEntry", "10.0", _iORegistryEntryGetParentEntryErr)
	}
	return _iORegistryEntryGetParentEntry(entry, plane, parent), nil
}

// IORegistryEntryGetParentEntry returns the first parent of a registry entry in a plane.
//
// See: https://developer.apple.com/documentation/iokit/1514454-ioregistryentrygetparententry
func IORegistryEntryGetParentEntry(entry uint32, plane string, parent *uint32) int32 {
	result, callErr := tryIORegistryEntryGetParentEntry(entry, plane, parent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntryGetParentIterator func(entry uint32, plane string, iterator *uint32) int32
var _iORegistryEntryGetParentIteratorErr error

func tryIORegistryEntryGetParentIterator(entry uint32, plane string, iterator *uint32) (int32, error) {
	if _iORegistryEntryGetParentIterator == nil {
		return 0, symbolCallError("IORegistryEntryGetParentIterator", "10.0", _iORegistryEntryGetParentIteratorErr)
	}
	return _iORegistryEntryGetParentIterator(entry, plane, iterator), nil
}

// IORegistryEntryGetParentIterator returns an iterator over a registry entry’s parent entries in a plane.
//
// See: https://developer.apple.com/documentation/iokit/1514366-ioregistryentrygetparentiterator
func IORegistryEntryGetParentIterator(entry uint32, plane string, iterator *uint32) int32 {
	result, callErr := tryIORegistryEntryGetParentIterator(entry, plane, iterator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntryGetPath func(entry uint32, plane string, path string) int32
var _iORegistryEntryGetPathErr error

func tryIORegistryEntryGetPath(entry uint32, plane string, path string) (int32, error) {
	if _iORegistryEntryGetPath == nil {
		return 0, symbolCallError("IORegistryEntryGetPath", "10.0", _iORegistryEntryGetPathErr)
	}
	return _iORegistryEntryGetPath(entry, plane, path), nil
}

// IORegistryEntryGetPath create a path for a registry entry.
//
// See: https://developer.apple.com/documentation/iokit/1514229-ioregistryentrygetpath
func IORegistryEntryGetPath(entry uint32, plane string, path string) int32 {
	result, callErr := tryIORegistryEntryGetPath(entry, plane, path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntryGetProperty func(entry uint32, propertyName string, buffer string, size *uint32) int32
var _iORegistryEntryGetPropertyErr error

func tryIORegistryEntryGetProperty(entry uint32, propertyName string, buffer string, size *uint32) (int32, error) {
	if _iORegistryEntryGetProperty == nil {
		return 0, symbolCallError("IORegistryEntryGetProperty", "10.0", _iORegistryEntryGetPropertyErr)
	}
	return _iORegistryEntryGetProperty(entry, propertyName, buffer, size), nil
}

// IORegistryEntryGetProperty.
//
// See: https://developer.apple.com/documentation/iokit/1514254-ioregistryentrygetproperty
func IORegistryEntryGetProperty(entry uint32, propertyName string, buffer string, size *uint32) int32 {
	result, callErr := tryIORegistryEntryGetProperty(entry, propertyName, buffer, size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntryGetRegistryEntryID func(entry uint32, entryID *uint64) int32
var _iORegistryEntryGetRegistryEntryIDErr error

func tryIORegistryEntryGetRegistryEntryID(entry uint32, entryID *uint64) (int32, error) {
	if _iORegistryEntryGetRegistryEntryID == nil {
		return 0, symbolCallError("IORegistryEntryGetRegistryEntryID", "10.6", _iORegistryEntryGetRegistryEntryIDErr)
	}
	return _iORegistryEntryGetRegistryEntryID(entry, entryID), nil
}

// IORegistryEntryGetRegistryEntryID returns an ID for the registry entry that is global to all tasks.
//
// See: https://developer.apple.com/documentation/iokit/1514719-ioregistryentrygetregistryentryi
func IORegistryEntryGetRegistryEntryID(entry uint32, entryID *uint64) int32 {
	result, callErr := tryIORegistryEntryGetRegistryEntryID(entry, entryID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntryIDMatching func(entryID uint64) corefoundation.CFMutableDictionary
var _iORegistryEntryIDMatchingErr error

func tryIORegistryEntryIDMatching(entryID uint64) (corefoundation.CFMutableDictionary, error) {
	if _iORegistryEntryIDMatching == nil {
		return *new(corefoundation.CFMutableDictionary), symbolCallError("IORegistryEntryIDMatching", "10.6", _iORegistryEntryIDMatchingErr)
	}
	return _iORegistryEntryIDMatching(entryID), nil
}

// IORegistryEntryIDMatching create a matching dictionary that specifies an IOService match based on a registry entry ID.
//
// See: https://developer.apple.com/documentation/iokit/1514880-ioregistryentryidmatching
func IORegistryEntryIDMatching(entryID uint64) corefoundation.CFMutableDictionary {
	result, callErr := tryIORegistryEntryIDMatching(entryID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntryInPlane func(entry uint32, plane string) int32
var _iORegistryEntryInPlaneErr error

func tryIORegistryEntryInPlane(entry uint32, plane string) (int32, error) {
	if _iORegistryEntryInPlane == nil {
		return 0, symbolCallError("IORegistryEntryInPlane", "10.0", _iORegistryEntryInPlaneErr)
	}
	return _iORegistryEntryInPlane(entry, plane), nil
}

// IORegistryEntryInPlane determines if the registry entry is attached in a plane.
//
// See: https://developer.apple.com/documentation/iokit/1514668-ioregistryentryinplane
func IORegistryEntryInPlane(entry uint32, plane string) int32 {
	result, callErr := tryIORegistryEntryInPlane(entry, plane)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntrySearchCFProperty func(entry uint32, plane string, key corefoundation.CFStringRef, allocator corefoundation.CFAllocatorRef, options uint32) corefoundation.CFTypeRef
var _iORegistryEntrySearchCFPropertyErr error

func tryIORegistryEntrySearchCFProperty(entry uint32, plane string, key corefoundation.CFStringRef, allocator corefoundation.CFAllocatorRef, options uint32) (corefoundation.CFTypeRef, error) {
	if _iORegistryEntrySearchCFProperty == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("IORegistryEntrySearchCFProperty", "10.1", _iORegistryEntrySearchCFPropertyErr)
	}
	return _iORegistryEntrySearchCFProperty(entry, plane, key, allocator, options), nil
}

// IORegistryEntrySearchCFProperty create a CF representation of a registry entry's property.
//
// See: https://developer.apple.com/documentation/iokit/1514537-ioregistryentrysearchcfproperty
func IORegistryEntrySearchCFProperty(entry uint32, plane string, key corefoundation.CFStringRef, allocator corefoundation.CFAllocatorRef, options uint32) corefoundation.CFTypeRef {
	result, callErr := tryIORegistryEntrySearchCFProperty(entry, plane, key, allocator, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntrySetCFProperties func(entry uint32, properties corefoundation.CFTypeRef) int32
var _iORegistryEntrySetCFPropertiesErr error

func tryIORegistryEntrySetCFProperties(entry uint32, properties corefoundation.CFTypeRef) (int32, error) {
	if _iORegistryEntrySetCFProperties == nil {
		return 0, symbolCallError("IORegistryEntrySetCFProperties", "10.0", _iORegistryEntrySetCFPropertiesErr)
	}
	return _iORegistryEntrySetCFProperties(entry, properties), nil
}

// IORegistryEntrySetCFProperties set CF container based properties in a registry entry.
//
// See: https://developer.apple.com/documentation/iokit/1514414-ioregistryentrysetcfproperties
func IORegistryEntrySetCFProperties(entry uint32, properties corefoundation.CFTypeRef) int32 {
	result, callErr := tryIORegistryEntrySetCFProperties(entry, properties)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryEntrySetCFProperty func(entry uint32, propertyName corefoundation.CFStringRef, property corefoundation.CFTypeRef) int32
var _iORegistryEntrySetCFPropertyErr error

func tryIORegistryEntrySetCFProperty(entry uint32, propertyName corefoundation.CFStringRef, property corefoundation.CFTypeRef) (int32, error) {
	if _iORegistryEntrySetCFProperty == nil {
		return 0, symbolCallError("IORegistryEntrySetCFProperty", "10.0", _iORegistryEntrySetCFPropertyErr)
	}
	return _iORegistryEntrySetCFProperty(entry, propertyName, property), nil
}

// IORegistryEntrySetCFProperty set a CF container based property in a registry entry.
//
// See: https://developer.apple.com/documentation/iokit/1514882-ioregistryentrysetcfproperty
func IORegistryEntrySetCFProperty(entry uint32, propertyName corefoundation.CFStringRef, property corefoundation.CFTypeRef) int32 {
	result, callErr := tryIORegistryEntrySetCFProperty(entry, propertyName, property)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryGetRootEntry func(mainPort uint32) uint32
var _iORegistryGetRootEntryErr error

func tryIORegistryGetRootEntry(mainPort uint32) (uint32, error) {
	if _iORegistryGetRootEntry == nil {
		return 0, symbolCallError("IORegistryGetRootEntry", "10.0", _iORegistryGetRootEntryErr)
	}
	return _iORegistryGetRootEntry(mainPort), nil
}

// IORegistryGetRootEntry return a handle to the registry root.
//
// See: https://developer.apple.com/documentation/iokit/1514878-ioregistrygetrootentry
func IORegistryGetRootEntry(mainPort uint32) uint32 {
	result, callErr := tryIORegistryGetRootEntry(mainPort)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryIteratorEnterEntry func(iterator uint32) int32
var _iORegistryIteratorEnterEntryErr error

func tryIORegistryIteratorEnterEntry(iterator uint32) (int32, error) {
	if _iORegistryIteratorEnterEntry == nil {
		return 0, symbolCallError("IORegistryIteratorEnterEntry", "10.0", _iORegistryIteratorEnterEntryErr)
	}
	return _iORegistryIteratorEnterEntry(iterator), nil
}

// IORegistryIteratorEnterEntry recurse into the current entry in the registry iteration.
//
// See: https://developer.apple.com/documentation/iokit/1514822-ioregistryiteratorenterentry
func IORegistryIteratorEnterEntry(iterator uint32) int32 {
	result, callErr := tryIORegistryIteratorEnterEntry(iterator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iORegistryIteratorExitEntry func(iterator uint32) int32
var _iORegistryIteratorExitEntryErr error

func tryIORegistryIteratorExitEntry(iterator uint32) (int32, error) {
	if _iORegistryIteratorExitEntry == nil {
		return 0, symbolCallError("IORegistryIteratorExitEntry", "10.0", _iORegistryIteratorExitEntryErr)
	}
	return _iORegistryIteratorExitEntry(iterator), nil
}

// IORegistryIteratorExitEntry exits a level of recursion, restoring the current entry.
//
// See: https://developer.apple.com/documentation/iokit/1514334-ioregistryiteratorexitentry
func IORegistryIteratorExitEntry(iterator uint32) int32 {
	result, callErr := tryIORegistryIteratorExitEntry(iterator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOServiceAddInterestNotification func(notifyPort IONotificationPortRef, service uint32, interestType string, callback unsafe.Pointer, refCon uintptr, notification *uint32) int32
var _iOServiceAddInterestNotificationErr error

func tryIOServiceAddInterestNotification(notifyPort IONotificationPortRef, service uint32, interestType string, callback IOServiceInterestCallback, refCon uintptr, notification *uint32) (int32, error) {
	if _iOServiceAddInterestNotification == nil {
		return 0, symbolCallError("IOServiceAddInterestNotification", "10.0", _iOServiceAddInterestNotificationErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 unsafe.Pointer, blockArg1 uint32, blockArg2 uint32, blockArg3 unsafe.Pointer) {
		callback(blockArg0, blockArg1, blockArg2, blockArg3)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _iOServiceAddInterestNotification(notifyPort, service, interestType, _block0, refCon, notification), nil
}

// IOServiceAddInterestNotification register for notification of state changes in an IOService.
//
// See: https://developer.apple.com/documentation/iokit/1514866-ioserviceaddinterestnotification
func IOServiceAddInterestNotification(notifyPort IONotificationPortRef, service uint32, interestType string, callback IOServiceInterestCallback, refCon uintptr, notification *uint32) int32 {
	result, callErr := tryIOServiceAddInterestNotification(notifyPort, service, interestType, callback, refCon, notification)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOServiceAddMatchingNotification func(notifyPort IONotificationPortRef, notificationType string, matching corefoundation.CFDictionaryRef, callback unsafe.Pointer, refCon uintptr, notification *uint32) int32
var _iOServiceAddMatchingNotificationErr error

func tryIOServiceAddMatchingNotification(notifyPort IONotificationPortRef, notificationType string, matching corefoundation.CFDictionaryRef, callback IOServiceMatchingCallback, refCon uintptr, notification *uint32) (int32, error) {
	if _iOServiceAddMatchingNotification == nil {
		return 0, symbolCallError("IOServiceAddMatchingNotification", "10.0", _iOServiceAddMatchingNotificationErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 unsafe.Pointer, blockArg1 uint32) { callback(blockArg0, blockArg1) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _iOServiceAddMatchingNotification(notifyPort, notificationType, matching, _block0, refCon, notification), nil
}

// IOServiceAddMatchingNotification look up registered IOService objects that match a matching dictionary, and install a notification request of new IOServices that match.
//
// See: https://developer.apple.com/documentation/iokit/1514362-ioserviceaddmatchingnotification
func IOServiceAddMatchingNotification(notifyPort IONotificationPortRef, notificationType string, matching corefoundation.CFDictionaryRef, callback IOServiceMatchingCallback, refCon uintptr, notification *uint32) int32 {
	result, callErr := tryIOServiceAddMatchingNotification(notifyPort, notificationType, matching, callback, refCon, notification)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOServiceAddNotification func(mainPort uint32, notificationType string, matching corefoundation.CFDictionaryRef, wakePort uint32, reference unsafe.Pointer, notification *uint32) int32
var _iOServiceAddNotificationErr error

func tryIOServiceAddNotification(mainPort uint32, notificationType string, matching corefoundation.CFDictionaryRef, wakePort uint32, reference unsafe.Pointer, notification *uint32) (int32, error) {
	if _iOServiceAddNotification == nil {
		return 0, symbolCallError("IOServiceAddNotification", "10.0", _iOServiceAddNotificationErr)
	}
	return _iOServiceAddNotification(mainPort, notificationType, matching, wakePort, reference, notification), nil
}

// IOServiceAddNotification.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/iokit/1514382-ioserviceaddnotification
func IOServiceAddNotification(mainPort uint32, notificationType string, matching corefoundation.CFDictionaryRef, wakePort uint32, reference unsafe.Pointer, notification *uint32) int32 {
	result, callErr := tryIOServiceAddNotification(mainPort, notificationType, matching, wakePort, reference, notification)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOServiceAuthorize func(service uint32, options uint32) int32
var _iOServiceAuthorizeErr error

func tryIOServiceAuthorize(service uint32, options uint32) (int32, error) {
	if _iOServiceAuthorize == nil {
		return 0, symbolCallError("IOServiceAuthorize", "10.0", _iOServiceAuthorizeErr)
	}
	return _iOServiceAuthorize(service, options), nil
}

// IOServiceAuthorize.
//
// See: https://developer.apple.com/documentation/iokit/1514533-ioserviceauthorize
func IOServiceAuthorize(service uint32, options uint32) int32 {
	result, callErr := tryIOServiceAuthorize(service, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOServiceClose func(connect uint32) int32
var _iOServiceCloseErr error

func tryIOServiceClose(connect uint32) (int32, error) {
	if _iOServiceClose == nil {
		return 0, symbolCallError("IOServiceClose", "10.0", _iOServiceCloseErr)
	}
	return _iOServiceClose(connect), nil
}

// IOServiceClose close a connection to an IOService and destroy the connect handle.
//
// See: https://developer.apple.com/documentation/iokit/1514646-ioserviceclose
func IOServiceClose(connect uint32) int32 {
	result, callErr := tryIOServiceClose(connect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOServiceGetBusyState func(service uint32, busyState *uint32) int32
var _iOServiceGetBusyStateErr error

func tryIOServiceGetBusyState(service uint32, busyState *uint32) (int32, error) {
	if _iOServiceGetBusyState == nil {
		return 0, symbolCallError("IOServiceGetBusyState", "10.0", _iOServiceGetBusyStateErr)
	}
	return _iOServiceGetBusyState(service, busyState), nil
}

// IOServiceGetBusyState returns the busyState of an IOService.
//
// See: https://developer.apple.com/documentation/iokit/1514607-ioservicegetbusystate
func IOServiceGetBusyState(service uint32, busyState *uint32) int32 {
	result, callErr := tryIOServiceGetBusyState(service, busyState)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOServiceGetMatchingService func(mainPort uint32, matching corefoundation.CFDictionaryRef) uint32
var _iOServiceGetMatchingServiceErr error

func tryIOServiceGetMatchingService(mainPort uint32, matching corefoundation.CFDictionaryRef) (uint32, error) {
	if _iOServiceGetMatchingService == nil {
		return 0, symbolCallError("IOServiceGetMatchingService", "10.2", _iOServiceGetMatchingServiceErr)
	}
	return _iOServiceGetMatchingService(mainPort, matching), nil
}

// IOServiceGetMatchingService look up a registered IOService object that matches a matching dictionary.
//
// See: https://developer.apple.com/documentation/iokit/1514535-ioservicegetmatchingservice
func IOServiceGetMatchingService(mainPort uint32, matching corefoundation.CFDictionaryRef) uint32 {
	result, callErr := tryIOServiceGetMatchingService(mainPort, matching)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOServiceGetMatchingServices func(mainPort uint32, matching corefoundation.CFDictionaryRef, existing *uint32) int32
var _iOServiceGetMatchingServicesErr error

func tryIOServiceGetMatchingServices(mainPort uint32, matching corefoundation.CFDictionaryRef, existing *uint32) (int32, error) {
	if _iOServiceGetMatchingServices == nil {
		return 0, symbolCallError("IOServiceGetMatchingServices", "10.0", _iOServiceGetMatchingServicesErr)
	}
	return _iOServiceGetMatchingServices(mainPort, matching, existing), nil
}

// IOServiceGetMatchingServices look up registered IOService objects that match a matching dictionary.
//
// See: https://developer.apple.com/documentation/iokit/1514494-ioservicegetmatchingservices
func IOServiceGetMatchingServices(mainPort uint32, matching corefoundation.CFDictionaryRef, existing *uint32) int32 {
	result, callErr := tryIOServiceGetMatchingServices(mainPort, matching, existing)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOServiceMatchPropertyTable func(service uint32, matching corefoundation.CFDictionaryRef, matches *bool) int32
var _iOServiceMatchPropertyTableErr error

func tryIOServiceMatchPropertyTable(service uint32, matching corefoundation.CFDictionaryRef, matches *bool) (int32, error) {
	if _iOServiceMatchPropertyTable == nil {
		return 0, symbolCallError("IOServiceMatchPropertyTable", "10.0", _iOServiceMatchPropertyTableErr)
	}
	return _iOServiceMatchPropertyTable(service, matching, matches), nil
}

// IOServiceMatchPropertyTable match an IOService objects with matching dictionary.
//
// See: https://developer.apple.com/documentation/iokit/1514685-ioservicematchpropertytable
func IOServiceMatchPropertyTable(service uint32, matching corefoundation.CFDictionaryRef, matches *bool) int32 {
	result, callErr := tryIOServiceMatchPropertyTable(service, matching, matches)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOServiceMatching func(name string) corefoundation.CFMutableDictionary
var _iOServiceMatchingErr error

func tryIOServiceMatching(name string) (corefoundation.CFMutableDictionary, error) {
	if _iOServiceMatching == nil {
		return *new(corefoundation.CFMutableDictionary), symbolCallError("IOServiceMatching", "10.0", _iOServiceMatchingErr)
	}
	return _iOServiceMatching(name), nil
}

// IOServiceMatching create a matching dictionary that specifies an IOService class match.
//
// See: https://developer.apple.com/documentation/iokit/1514687-ioservicematching
func IOServiceMatching(name string) corefoundation.CFMutableDictionary {
	result, callErr := tryIOServiceMatching(name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOServiceNameMatching func(name string) corefoundation.CFMutableDictionary
var _iOServiceNameMatchingErr error

func tryIOServiceNameMatching(name string) (corefoundation.CFMutableDictionary, error) {
	if _iOServiceNameMatching == nil {
		return *new(corefoundation.CFMutableDictionary), symbolCallError("IOServiceNameMatching", "10.0", _iOServiceNameMatchingErr)
	}
	return _iOServiceNameMatching(name), nil
}

// IOServiceNameMatching create a matching dictionary that specifies an IOService name match.
//
// See: https://developer.apple.com/documentation/iokit/1514416-ioservicenamematching
func IOServiceNameMatching(name string) corefoundation.CFMutableDictionary {
	result, callErr := tryIOServiceNameMatching(name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOServiceOFPathToBSDName func(mainPort uint32, openFirmwarePath string, bsdName string) int32
var _iOServiceOFPathToBSDNameErr error

func tryIOServiceOFPathToBSDName(mainPort uint32, openFirmwarePath string, bsdName string) (int32, error) {
	if _iOServiceOFPathToBSDName == nil {
		return 0, symbolCallError("IOServiceOFPathToBSDName", "10.0", _iOServiceOFPathToBSDNameErr)
	}
	return _iOServiceOFPathToBSDName(mainPort, openFirmwarePath, bsdName), nil
}

// IOServiceOFPathToBSDName.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/iokit/1514661-ioserviceofpathtobsdname
func IOServiceOFPathToBSDName(mainPort uint32, openFirmwarePath string, bsdName string) int32 {
	result, callErr := tryIOServiceOFPathToBSDName(mainPort, openFirmwarePath, bsdName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOServiceOpen func(service uint32, owningTask uint32, type_ uint32, connect *uint32) int32
var _iOServiceOpenErr error

func tryIOServiceOpen(service uint32, owningTask uint32, type_ uint32, connect *uint32) (int32, error) {
	if _iOServiceOpen == nil {
		return 0, symbolCallError("IOServiceOpen", "10.0", _iOServiceOpenErr)
	}
	return _iOServiceOpen(service, owningTask, type_, connect), nil
}

// IOServiceOpen a request to create a connection to an IOService.
//
// See: https://developer.apple.com/documentation/iokit/1514515-ioserviceopen
func IOServiceOpen(service uint32, owningTask uint32, type_ uint32, connect *uint32) int32 {
	result, callErr := tryIOServiceOpen(service, owningTask, type_, connect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOServiceOpenAsFileDescriptor func(service uint32, oflag int32) int32
var _iOServiceOpenAsFileDescriptorErr error

func tryIOServiceOpenAsFileDescriptor(service uint32, oflag int32) (int32, error) {
	if _iOServiceOpenAsFileDescriptor == nil {
		return 0, symbolCallError("IOServiceOpenAsFileDescriptor", "10.0", _iOServiceOpenAsFileDescriptorErr)
	}
	return _iOServiceOpenAsFileDescriptor(service, oflag), nil
}

// IOServiceOpenAsFileDescriptor.
//
// See: https://developer.apple.com/documentation/iokit/1514879-ioserviceopenasfiledescriptor
func IOServiceOpenAsFileDescriptor(service uint32, oflag int32) int32 {
	result, callErr := tryIOServiceOpenAsFileDescriptor(service, oflag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOServiceRequestProbe func(service uint32, options uint32) int32
var _iOServiceRequestProbeErr error

func tryIOServiceRequestProbe(service uint32, options uint32) (int32, error) {
	if _iOServiceRequestProbe == nil {
		return 0, symbolCallError("IOServiceRequestProbe", "10.0", _iOServiceRequestProbeErr)
	}
	return _iOServiceRequestProbe(service, options), nil
}

// IOServiceRequestProbe a request to rescan a bus for device changes.
//
// See: https://developer.apple.com/documentation/iokit/1514364-ioservicerequestprobe
func IOServiceRequestProbe(service uint32, options uint32) int32 {
	result, callErr := tryIOServiceRequestProbe(service, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOServiceWaitQuiet func(service uint32, waitTime *kernel.MachTimespec) int32
var _iOServiceWaitQuietErr error

func tryIOServiceWaitQuiet(service uint32, waitTime *kernel.MachTimespec) (int32, error) {
	if _iOServiceWaitQuiet == nil {
		return 0, symbolCallError("IOServiceWaitQuiet", "10.0", _iOServiceWaitQuietErr)
	}
	return _iOServiceWaitQuiet(service, waitTime), nil
}

// IOServiceWaitQuiet wait for an IOService's busyState to be zero.
//
// See: https://developer.apple.com/documentation/iokit/1514573-ioservicewaitquiet
func IOServiceWaitQuiet(service uint32, waitTime *kernel.MachTimespec) int32 {
	result, callErr := tryIOServiceWaitQuiet(service, waitTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOURLCreateDataAndPropertiesFromResource func(alloc corefoundation.CFAllocatorRef, url corefoundation.CFURLRef, resourceData *corefoundation.CFDataRef, properties *corefoundation.CFDictionaryRef, desiredProperties corefoundation.CFArrayRef, errorCode *int32) bool
var _iOURLCreateDataAndPropertiesFromResourceErr error

func tryIOURLCreateDataAndPropertiesFromResource(alloc corefoundation.CFAllocatorRef, url corefoundation.CFURLRef, resourceData *corefoundation.CFDataRef, properties *corefoundation.CFDictionaryRef, desiredProperties corefoundation.CFArrayRef, errorCode *int32) (bool, error) {
	if _iOURLCreateDataAndPropertiesFromResource == nil {
		return false, symbolCallError("IOURLCreateDataAndPropertiesFromResource", "10.0", _iOURLCreateDataAndPropertiesFromResourceErr)
	}
	return _iOURLCreateDataAndPropertiesFromResource(alloc, url, resourceData, properties, desiredProperties, errorCode), nil
}

// IOURLCreateDataAndPropertiesFromResource.
//
// See: https://developer.apple.com/documentation/iokit/1390335-iourlcreatedataandpropertiesfrom
func IOURLCreateDataAndPropertiesFromResource(alloc corefoundation.CFAllocatorRef, url corefoundation.CFURLRef, resourceData *corefoundation.CFDataRef, properties *corefoundation.CFDictionaryRef, desiredProperties corefoundation.CFArrayRef, errorCode *int32) bool {
	result, callErr := tryIOURLCreateDataAndPropertiesFromResource(alloc, url, resourceData, properties, desiredProperties, errorCode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iOURLCreatePropertyFromResource func(alloc corefoundation.CFAllocatorRef, url corefoundation.CFURLRef, property corefoundation.CFStringRef, errorCode *int32) corefoundation.CFTypeRef
var _iOURLCreatePropertyFromResourceErr error

func tryIOURLCreatePropertyFromResource(alloc corefoundation.CFAllocatorRef, url corefoundation.CFURLRef, property corefoundation.CFStringRef, errorCode *int32) (corefoundation.CFTypeRef, error) {
	if _iOURLCreatePropertyFromResource == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("IOURLCreatePropertyFromResource", "10.0", _iOURLCreatePropertyFromResourceErr)
	}
	return _iOURLCreatePropertyFromResource(alloc, url, property, errorCode), nil
}

// IOURLCreatePropertyFromResource.
//
// See: https://developer.apple.com/documentation/iokit/1390322-iourlcreatepropertyfromresource
func IOURLCreatePropertyFromResource(alloc corefoundation.CFAllocatorRef, url corefoundation.CFURLRef, property corefoundation.CFStringRef, errorCode *int32) corefoundation.CFTypeRef {
	result, callErr := tryIOURLCreatePropertyFromResource(alloc, url, property, errorCode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _kextManagerCopyLoadedKextInfo func(kextIdentifiers corefoundation.CFArrayRef, infoKeys corefoundation.CFArrayRef) corefoundation.CFDictionaryRef
var _kextManagerCopyLoadedKextInfoErr error

func tryKextManagerCopyLoadedKextInfo(kextIdentifiers corefoundation.CFArrayRef, infoKeys corefoundation.CFArrayRef) (corefoundation.CFDictionaryRef, error) {
	if _kextManagerCopyLoadedKextInfo == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("KextManagerCopyLoadedKextInfo", "10.7", _kextManagerCopyLoadedKextInfoErr)
	}
	return _kextManagerCopyLoadedKextInfo(kextIdentifiers, infoKeys), nil
}

// KextManagerCopyLoadedKextInfo returns information about loaded kexts in a dictionary.
//
// See: https://developer.apple.com/documentation/iokit/1562908-kextmanagercopyloadedkextinfo
func KextManagerCopyLoadedKextInfo(kextIdentifiers corefoundation.CFArrayRef, infoKeys corefoundation.CFArrayRef) corefoundation.CFDictionaryRef {
	result, callErr := tryKextManagerCopyLoadedKextInfo(kextIdentifiers, infoKeys)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _kextManagerCreateURLForBundleIdentifier func(allocator corefoundation.CFAllocatorRef, kextIdentifier corefoundation.CFStringRef) corefoundation.CFURLRef
var _kextManagerCreateURLForBundleIdentifierErr error

func tryKextManagerCreateURLForBundleIdentifier(allocator corefoundation.CFAllocatorRef, kextIdentifier corefoundation.CFStringRef) (corefoundation.CFURLRef, error) {
	if _kextManagerCreateURLForBundleIdentifier == nil {
		return *new(corefoundation.CFURLRef), symbolCallError("KextManagerCreateURLForBundleIdentifier", "10.2", _kextManagerCreateURLForBundleIdentifierErr)
	}
	return _kextManagerCreateURLForBundleIdentifier(allocator, kextIdentifier), nil
}

// KextManagerCreateURLForBundleIdentifier create a URL locating a kext with a given bundle identifier.
//
// See: https://developer.apple.com/documentation/iokit/1562905-kextmanagercreateurlforbundleide
func KextManagerCreateURLForBundleIdentifier(allocator corefoundation.CFAllocatorRef, kextIdentifier corefoundation.CFStringRef) corefoundation.CFURLRef {
	result, callErr := tryKextManagerCreateURLForBundleIdentifier(allocator, kextIdentifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _kextManagerLoadKextWithIdentifier func(kextIdentifier corefoundation.CFStringRef, dependencyKextAndFolderURLs corefoundation.CFArrayRef) int32
var _kextManagerLoadKextWithIdentifierErr error

func tryKextManagerLoadKextWithIdentifier(kextIdentifier corefoundation.CFStringRef, dependencyKextAndFolderURLs corefoundation.CFArrayRef) (int32, error) {
	if _kextManagerLoadKextWithIdentifier == nil {
		return 0, symbolCallError("KextManagerLoadKextWithIdentifier", "10.6", _kextManagerLoadKextWithIdentifierErr)
	}
	return _kextManagerLoadKextWithIdentifier(kextIdentifier, dependencyKextAndFolderURLs), nil
}

// KextManagerLoadKextWithIdentifier request the kext loading system to load a kext with a given bundle identifier.
//
// See: https://developer.apple.com/documentation/iokit/1562907-kextmanagerloadkextwithidentifie
func KextManagerLoadKextWithIdentifier(kextIdentifier corefoundation.CFStringRef, dependencyKextAndFolderURLs corefoundation.CFArrayRef) int32 {
	result, callErr := tryKextManagerLoadKextWithIdentifier(kextIdentifier, dependencyKextAndFolderURLs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _kextManagerLoadKextWithURL func(kextURL corefoundation.CFURLRef, dependencyKextAndFolderURLs corefoundation.CFArrayRef) int32
var _kextManagerLoadKextWithURLErr error

func tryKextManagerLoadKextWithURL(kextURL corefoundation.CFURLRef, dependencyKextAndFolderURLs corefoundation.CFArrayRef) (int32, error) {
	if _kextManagerLoadKextWithURL == nil {
		return 0, symbolCallError("KextManagerLoadKextWithURL", "10.6", _kextManagerLoadKextWithURLErr)
	}
	return _kextManagerLoadKextWithURL(kextURL, dependencyKextAndFolderURLs), nil
}

// KextManagerLoadKextWithURL request the kext loading system to load a kext with a given URL.
//
// See: https://developer.apple.com/documentation/iokit/1562906-kextmanagerloadkextwithurl
func KextManagerLoadKextWithURL(kextURL corefoundation.CFURLRef, dependencyKextAndFolderURLs corefoundation.CFArrayRef) int32 {
	result, callErr := tryKextManagerLoadKextWithURL(kextURL, dependencyKextAndFolderURLs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _kextManagerUnloadKextWithIdentifier func(kextIdentifier corefoundation.CFStringRef) int32
var _kextManagerUnloadKextWithIdentifierErr error

func tryKextManagerUnloadKextWithIdentifier(kextIdentifier corefoundation.CFStringRef) (int32, error) {
	if _kextManagerUnloadKextWithIdentifier == nil {
		return 0, symbolCallError("KextManagerUnloadKextWithIdentifier", "10.7", _kextManagerUnloadKextWithIdentifierErr)
	}
	return _kextManagerUnloadKextWithIdentifier(kextIdentifier), nil
}

// KextManagerUnloadKextWithIdentifier request the kernel to unload a kext with a given bundle identifier.
//
// See: https://developer.apple.com/documentation/iokit/1562904-kextmanagerunloadkextwithidentif
func KextManagerUnloadKextWithIdentifier(kextIdentifier corefoundation.CFStringRef) int32 {
	result, callErr := tryKextManagerUnloadKextWithIdentifier(kextIdentifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXClickTime func(handle NXEventHandle) float64
var _nXClickTimeErr error

func tryNXClickTime(handle NXEventHandle) (float64, error) {
	if _nXClickTime == nil {
		return 0.0, symbolCallError("NXClickTime", "10.0", _nXClickTimeErr)
	}
	return _nXClickTime(handle), nil
}

// NXClickTime.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574526-nxclicktime
func NXClickTime(handle NXEventHandle) float64 {
	result, callErr := tryNXClickTime(handle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXCloseEventStatus func(handle NXEventHandle)
var _nXCloseEventStatusErr error

func tryNXCloseEventStatus(handle NXEventHandle) error {
	if _nXCloseEventStatus == nil {
		return symbolCallError("NXCloseEventStatus", "10.0", _nXCloseEventStatusErr)
	}
	_nXCloseEventStatus(handle)
	return nil
}

// NXCloseEventStatus.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574524-nxcloseeventstatus
func NXCloseEventStatus(handle NXEventHandle) {
	if callErr := tryNXCloseEventStatus(handle); callErr != nil {
		panic(callErr)
	}
}

var _nXEventSystemInfo func(handle NXEventHandle, flavor string, evs_info *int32, evs_info_cnt *uint32) NXEventSystemInfoType
var _nXEventSystemInfoErr error

func tryNXEventSystemInfo(handle NXEventHandle, flavor string, evs_info *int32, evs_info_cnt *uint32) (NXEventSystemInfoType, error) {
	if _nXEventSystemInfo == nil {
		return *new(NXEventSystemInfoType), symbolCallError("NXEventSystemInfo", "10.3", _nXEventSystemInfoErr)
	}
	return _nXEventSystemInfo(handle, flavor, evs_info, evs_info_cnt), nil
}

// NXEventSystemInfo.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574527-nxeventsysteminfo
func NXEventSystemInfo(handle NXEventHandle, flavor string, evs_info *int32, evs_info_cnt *uint32) NXEventSystemInfoType {
	result, callErr := tryNXEventSystemInfo(handle, flavor, evs_info, evs_info_cnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXGetClickSpace func(handle NXEventHandle, area *NXSize)
var _nXGetClickSpaceErr error

func tryNXGetClickSpace(handle NXEventHandle, area *NXSize) error {
	if _nXGetClickSpace == nil {
		return symbolCallError("NXGetClickSpace", "10.0", _nXGetClickSpaceErr)
	}
	_nXGetClickSpace(handle, area)
	return nil
}

// NXGetClickSpace.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574518-nxgetclickspace
func NXGetClickSpace(handle NXEventHandle, area *NXSize) {
	if callErr := tryNXGetClickSpace(handle, area); callErr != nil {
		panic(callErr)
	}
}

var _nXKeyRepeatInterval func(handle NXEventHandle) float64
var _nXKeyRepeatIntervalErr error

func tryNXKeyRepeatInterval(handle NXEventHandle) (float64, error) {
	if _nXKeyRepeatInterval == nil {
		return 0.0, symbolCallError("NXKeyRepeatInterval", "10.0", _nXKeyRepeatIntervalErr)
	}
	return _nXKeyRepeatInterval(handle), nil
}

// NXKeyRepeatInterval.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574515-nxkeyrepeatinterval
func NXKeyRepeatInterval(handle NXEventHandle) float64 {
	result, callErr := tryNXKeyRepeatInterval(handle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXKeyRepeatThreshold func(handle NXEventHandle) float64
var _nXKeyRepeatThresholdErr error

func tryNXKeyRepeatThreshold(handle NXEventHandle) (float64, error) {
	if _nXKeyRepeatThreshold == nil {
		return 0.0, symbolCallError("NXKeyRepeatThreshold", "10.0", _nXKeyRepeatThresholdErr)
	}
	return _nXKeyRepeatThreshold(handle), nil
}

// NXKeyRepeatThreshold.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574514-nxkeyrepeatthreshold
func NXKeyRepeatThreshold(handle NXEventHandle) float64 {
	result, callErr := tryNXKeyRepeatThreshold(handle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXOpenEventStatus func() NXEventHandle
var _nXOpenEventStatusErr error

func tryNXOpenEventStatus() (NXEventHandle, error) {
	if _nXOpenEventStatus == nil {
		return *new(NXEventHandle), symbolCallError("NXOpenEventStatus", "10.0", _nXOpenEventStatusErr)
	}
	return _nXOpenEventStatus(), nil
}

// NXOpenEventStatus.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574517-nxopeneventstatus
func NXOpenEventStatus() NXEventHandle {
	result, callErr := tryNXOpenEventStatus()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXResetKeyboard func(handle NXEventHandle)
var _nXResetKeyboardErr error

func tryNXResetKeyboard(handle NXEventHandle) error {
	if _nXResetKeyboard == nil {
		return symbolCallError("NXResetKeyboard", "10.3", _nXResetKeyboardErr)
	}
	_nXResetKeyboard(handle)
	return nil
}

// NXResetKeyboard.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574523-nxresetkeyboard
func NXResetKeyboard(handle NXEventHandle) {
	if callErr := tryNXResetKeyboard(handle); callErr != nil {
		panic(callErr)
	}
}

var _nXResetMouse func(handle NXEventHandle)
var _nXResetMouseErr error

func tryNXResetMouse(handle NXEventHandle) error {
	if _nXResetMouse == nil {
		return symbolCallError("NXResetMouse", "10.3", _nXResetMouseErr)
	}
	_nXResetMouse(handle)
	return nil
}

// NXResetMouse.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574520-nxresetmouse
func NXResetMouse(handle NXEventHandle) {
	if callErr := tryNXResetMouse(handle); callErr != nil {
		panic(callErr)
	}
}

var _nXSetClickSpace func(handle NXEventHandle, area *NXSize)
var _nXSetClickSpaceErr error

func tryNXSetClickSpace(handle NXEventHandle, area *NXSize) error {
	if _nXSetClickSpace == nil {
		return symbolCallError("NXSetClickSpace", "10.0", _nXSetClickSpaceErr)
	}
	_nXSetClickSpace(handle, area)
	return nil
}

// NXSetClickSpace.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574525-nxsetclickspace
func NXSetClickSpace(handle NXEventHandle, area *NXSize) {
	if callErr := tryNXSetClickSpace(handle, area); callErr != nil {
		panic(callErr)
	}
}

var _nXSetClickTime func(handle NXEventHandle, seconds float64)
var _nXSetClickTimeErr error

func tryNXSetClickTime(handle NXEventHandle, seconds float64) error {
	if _nXSetClickTime == nil {
		return symbolCallError("NXSetClickTime", "10.0", _nXSetClickTimeErr)
	}
	_nXSetClickTime(handle, seconds)
	return nil
}

// NXSetClickTime.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574522-nxsetclicktime
func NXSetClickTime(handle NXEventHandle, seconds float64) {
	if callErr := tryNXSetClickTime(handle, seconds); callErr != nil {
		panic(callErr)
	}
}

var _nXSetKeyRepeatInterval func(handle NXEventHandle, seconds float64)
var _nXSetKeyRepeatIntervalErr error

func tryNXSetKeyRepeatInterval(handle NXEventHandle, seconds float64) error {
	if _nXSetKeyRepeatInterval == nil {
		return symbolCallError("NXSetKeyRepeatInterval", "10.0", _nXSetKeyRepeatIntervalErr)
	}
	_nXSetKeyRepeatInterval(handle, seconds)
	return nil
}

// NXSetKeyRepeatInterval.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574519-nxsetkeyrepeatinterval
func NXSetKeyRepeatInterval(handle NXEventHandle, seconds float64) {
	if callErr := tryNXSetKeyRepeatInterval(handle, seconds); callErr != nil {
		panic(callErr)
	}
}

var _nXSetKeyRepeatThreshold func(handle NXEventHandle, threshold float64)
var _nXSetKeyRepeatThresholdErr error

func tryNXSetKeyRepeatThreshold(handle NXEventHandle, threshold float64) error {
	if _nXSetKeyRepeatThreshold == nil {
		return symbolCallError("NXSetKeyRepeatThreshold", "10.0", _nXSetKeyRepeatThresholdErr)
	}
	_nXSetKeyRepeatThreshold(handle, threshold)
	return nil
}

// NXSetKeyRepeatThreshold.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574516-nxsetkeyrepeatthreshold
func NXSetKeyRepeatThreshold(handle NXEventHandle, threshold float64) {
	if callErr := tryNXSetKeyRepeatThreshold(handle, threshold); callErr != nil {
		panic(callErr)
	}
}

var _oSGetNotificationFromMessage func(msg unsafe.Pointer, index uint32, type_ *uint32, reference unsafe.Pointer, content unsafe.Pointer, size *uintptr) int32
var _oSGetNotificationFromMessageErr error

func tryOSGetNotificationFromMessage(msg unsafe.Pointer, index uint32, type_ *uint32, reference unsafe.Pointer, content unsafe.Pointer, size *uintptr) (int32, error) {
	if _oSGetNotificationFromMessage == nil {
		return 0, symbolCallError("OSGetNotificationFromMessage", "10.0", _oSGetNotificationFromMessageErr)
	}
	return _oSGetNotificationFromMessage(msg, index, type_, reference, content, size), nil
}

// OSGetNotificationFromMessage.
//
// See: https://developer.apple.com/documentation/iokit/1514263-osgetnotificationfrommessage
func OSGetNotificationFromMessage(msg unsafe.Pointer, index uint32, type_ *uint32, reference unsafe.Pointer, content unsafe.Pointer, size *uintptr) int32 {
	result, callErr := tryOSGetNotificationFromMessage(msg, index, type_, reference, content, size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_iOAccelFindAccelerator, &_iOAccelFindAcceleratorErr, frameworkHandle, "IOAccelFindAccelerator", "10.0")
	registerFunc(&_iOAllowPowerChange, &_iOAllowPowerChangeErr, frameworkHandle, "IOAllowPowerChange", "10.0")
	registerFunc(&_iOBSDNameMatching, &_iOBSDNameMatchingErr, frameworkHandle, "IOBSDNameMatching", "10.0")
	registerFunc(&_iOCFSerialize, &_iOCFSerializeErr, frameworkHandle, "IOCFSerialize", "10.0")
	registerFunc(&_iOCFUnserialize, &_iOCFUnserializeErr, frameworkHandle, "IOCFUnserialize", "10.0")
	registerFunc(&_iOCFUnserializeBinary, &_iOCFUnserializeBinaryErr, frameworkHandle, "IOCFUnserializeBinary", "10.10")
	registerFunc(&_iOCFUnserializeWithSize, &_iOCFUnserializeWithSizeErr, frameworkHandle, "IOCFUnserializeWithSize", "10.10")
	registerFunc(&_iOCancelPowerChange, &_iOCancelPowerChangeErr, frameworkHandle, "IOCancelPowerChange", "10.0")
	registerFunc(&_iOCatalogueGetData, &_iOCatalogueGetDataErr, frameworkHandle, "IOCatalogueGetData", "10.0")
	registerFunc(&_iOCatalogueModuleLoaded, &_iOCatalogueModuleLoadedErr, frameworkHandle, "IOCatalogueModuleLoaded", "10.0")
	registerFunc(&_iOCatalogueReset, &_iOCatalogueResetErr, frameworkHandle, "IOCatalogueReset", "10.0")
	registerFunc(&_iOCatalogueSendData, &_iOCatalogueSendDataErr, frameworkHandle, "IOCatalogueSendData", "10.0")
	registerFunc(&_iOCatalogueTerminate, &_iOCatalogueTerminateErr, frameworkHandle, "IOCatalogueTerminate", "10.0")
	registerFunc(&_iOConnectAddClient, &_iOConnectAddClientErr, frameworkHandle, "IOConnectAddClient", "10.0")
	registerFunc(&_iOConnectAddRef, &_iOConnectAddRefErr, frameworkHandle, "IOConnectAddRef", "10.0")
	registerFunc(&_iOConnectCallAsyncMethod, &_iOConnectCallAsyncMethodErr, frameworkHandle, "IOConnectCallAsyncMethod", "10.5")
	registerFunc(&_iOConnectCallAsyncScalarMethod, &_iOConnectCallAsyncScalarMethodErr, frameworkHandle, "IOConnectCallAsyncScalarMethod", "10.5")
	registerFunc(&_iOConnectCallAsyncStructMethod, &_iOConnectCallAsyncStructMethodErr, frameworkHandle, "IOConnectCallAsyncStructMethod", "10.5")
	registerFunc(&_iOConnectCallMethod, &_iOConnectCallMethodErr, frameworkHandle, "IOConnectCallMethod", "10.5")
	registerFunc(&_iOConnectCallScalarMethod, &_iOConnectCallScalarMethodErr, frameworkHandle, "IOConnectCallScalarMethod", "10.5")
	registerFunc(&_iOConnectCallStructMethod, &_iOConnectCallStructMethodErr, frameworkHandle, "IOConnectCallStructMethod", "10.5")
	registerFunc(&_iOConnectGetService, &_iOConnectGetServiceErr, frameworkHandle, "IOConnectGetService", "10.0")
	registerFunc(&_iOConnectMapMemory, &_iOConnectMapMemoryErr, frameworkHandle, "IOConnectMapMemory", "10.0")
	registerFunc(&_iOConnectMapMemory64, &_iOConnectMapMemory64Err, frameworkHandle, "IOConnectMapMemory64", "10.5")
	registerFunc(&_iOConnectRelease, &_iOConnectReleaseErr, frameworkHandle, "IOConnectRelease", "10.0")
	registerFunc(&_iOConnectSetCFProperties, &_iOConnectSetCFPropertiesErr, frameworkHandle, "IOConnectSetCFProperties", "10.0")
	registerFunc(&_iOConnectSetCFProperty, &_iOConnectSetCFPropertyErr, frameworkHandle, "IOConnectSetCFProperty", "10.0")
	registerFunc(&_iOConnectSetNotificationPort, &_iOConnectSetNotificationPortErr, frameworkHandle, "IOConnectSetNotificationPort", "10.0")
	registerFunc(&_iOConnectTrap0, &_iOConnectTrap0Err, frameworkHandle, "IOConnectTrap0", "10.0")
	registerFunc(&_iOConnectTrap1, &_iOConnectTrap1Err, frameworkHandle, "IOConnectTrap1", "10.0")
	registerFunc(&_iOConnectTrap2, &_iOConnectTrap2Err, frameworkHandle, "IOConnectTrap2", "10.0")
	registerFunc(&_iOConnectTrap3, &_iOConnectTrap3Err, frameworkHandle, "IOConnectTrap3", "10.0")
	registerFunc(&_iOConnectTrap4, &_iOConnectTrap4Err, frameworkHandle, "IOConnectTrap4", "10.0")
	registerFunc(&_iOConnectTrap5, &_iOConnectTrap5Err, frameworkHandle, "IOConnectTrap5", "10.0")
	registerFunc(&_iOConnectTrap6, &_iOConnectTrap6Err, frameworkHandle, "IOConnectTrap6", "10.0")
	registerFunc(&_iOConnectUnmapMemory, &_iOConnectUnmapMemoryErr, frameworkHandle, "IOConnectUnmapMemory", "10.0")
	registerFunc(&_iOConnectUnmapMemory64, &_iOConnectUnmapMemory64Err, frameworkHandle, "IOConnectUnmapMemory64", "10.5")
	registerFunc(&_iOCopySystemLoadAdvisoryDetailed, &_iOCopySystemLoadAdvisoryDetailedErr, frameworkHandle, "IOCopySystemLoadAdvisoryDetailed", "10.6")
	registerFunc(&_iOCreatePlugInInterfaceForService, &_iOCreatePlugInInterfaceForServiceErr, frameworkHandle, "IOCreatePlugInInterfaceForService", "10.0")
	registerFunc(&_iOCreateReceivePort, &_iOCreateReceivePortErr, frameworkHandle, "IOCreateReceivePort", "10.0")
	registerFunc(&_iODataQueueAllocateNotificationPort, &_iODataQueueAllocateNotificationPortErr, frameworkHandle, "IODataQueueAllocateNotificationPort", "10.0")
	registerFunc(&_iODataQueueDataAvailable, &_iODataQueueDataAvailableErr, frameworkHandle, "IODataQueueDataAvailable", "10.0")
	registerFunc(&_iODataQueueDequeue, &_iODataQueueDequeueErr, frameworkHandle, "IODataQueueDequeue", "10.0")
	registerFunc(&_iODataQueueEnqueue, &_iODataQueueEnqueueErr, frameworkHandle, "IODataQueueEnqueue", "10.5")
	registerFunc(&_iODataQueuePeek, &_iODataQueuePeekErr, frameworkHandle, "IODataQueuePeek", "10.0")
	registerFunc(&_iODataQueueSetNotificationPort, &_iODataQueueSetNotificationPortErr, frameworkHandle, "IODataQueueSetNotificationPort", "10.5")
	registerFunc(&_iODataQueueWaitForAvailableData, &_iODataQueueWaitForAvailableDataErr, frameworkHandle, "IODataQueueWaitForAvailableData", "10.0")
	registerFunc(&_iODeregisterApp, &_iODeregisterAppErr, frameworkHandle, "IODeregisterApp", "10.0")
	registerFunc(&_iODeregisterForSystemPower, &_iODeregisterForSystemPowerErr, frameworkHandle, "IODeregisterForSystemPower", "10.0")
	registerFunc(&_iODestroyPlugInInterface, &_iODestroyPlugInInterfaceErr, frameworkHandle, "IODestroyPlugInInterface", "10.0")
	registerFunc(&_iODispatchCalloutFromMessage, &_iODispatchCalloutFromMessageErr, frameworkHandle, "IODispatchCalloutFromMessage", "10.0")
	registerFunc(&_iODisplayCommitParameters, &_iODisplayCommitParametersErr, frameworkHandle, "IODisplayCommitParameters", "10.0")
	registerFunc(&_iODisplayCopyFloatParameters, &_iODisplayCopyFloatParametersErr, frameworkHandle, "IODisplayCopyFloatParameters", "10.0")
	registerFunc(&_iODisplayCopyParameters, &_iODisplayCopyParametersErr, frameworkHandle, "IODisplayCopyParameters", "10.0")
	registerFunc(&_iODisplayCreateInfoDictionary, &_iODisplayCreateInfoDictionaryErr, frameworkHandle, "IODisplayCreateInfoDictionary", "10.0")
	registerFunc(&_iODisplayForFramebuffer, &_iODisplayForFramebufferErr, frameworkHandle, "IODisplayForFramebuffer", "10.7")
	registerFunc(&_iODisplayGetFloatParameter, &_iODisplayGetFloatParameterErr, frameworkHandle, "IODisplayGetFloatParameter", "10.0")
	registerFunc(&_iODisplayGetIntegerRangeParameter, &_iODisplayGetIntegerRangeParameterErr, frameworkHandle, "IODisplayGetIntegerRangeParameter", "10.0")
	registerFunc(&_iODisplayMatchDictionaries, &_iODisplayMatchDictionariesErr, frameworkHandle, "IODisplayMatchDictionaries", "10.0")
	registerFunc(&_iODisplaySetFloatParameter, &_iODisplaySetFloatParameterErr, frameworkHandle, "IODisplaySetFloatParameter", "10.0")
	registerFunc(&_iODisplaySetIntegerParameter, &_iODisplaySetIntegerParameterErr, frameworkHandle, "IODisplaySetIntegerParameter", "10.0")
	registerFunc(&_iODisplaySetParameters, &_iODisplaySetParametersErr, frameworkHandle, "IODisplaySetParameters", "10.0")
	registerFunc(&_iOFBCopyI2CInterfaceForBus, &_iOFBCopyI2CInterfaceForBusErr, frameworkHandle, "IOFBCopyI2CInterfaceForBus", "10.3")
	registerFunc(&_iOFBGetI2CInterfaceCount, &_iOFBGetI2CInterfaceCountErr, frameworkHandle, "IOFBGetI2CInterfaceCount", "10.3")
	registerFunc(&_iOFramebufferOpen, &_iOFramebufferOpenErr, frameworkHandle, "IOFramebufferOpen", "10.0")
	registerFunc(&_iOGetSystemLoadAdvisory, &_iOGetSystemLoadAdvisoryErr, frameworkHandle, "IOGetSystemLoadAdvisory", "10.6")
	registerFunc(&_iOHIDAccessCheckAuditToken, &_iOHIDAccessCheckAuditTokenErr, frameworkHandle, "IOHIDAccessCheckAuditToken", "10.15")
	registerFunc(&_iOHIDCheckAccess, &_iOHIDCheckAccessErr, frameworkHandle, "IOHIDCheckAccess", "10.15")
	registerFunc(&_iOHIDCopyCFTypeParameter, &_iOHIDCopyCFTypeParameterErr, frameworkHandle, "IOHIDCopyCFTypeParameter", "10.3")
	registerFunc(&_iOHIDCreateSharedMemory, &_iOHIDCreateSharedMemoryErr, frameworkHandle, "IOHIDCreateSharedMemory", "10.0")
	registerFunc(&_iOHIDDeviceActivate, &_iOHIDDeviceActivateErr, frameworkHandle, "IOHIDDeviceActivate", "10.15")
	registerFunc(&_iOHIDDeviceCancel, &_iOHIDDeviceCancelErr, frameworkHandle, "IOHIDDeviceCancel", "10.15")
	registerFunc(&_iOHIDDeviceClose, &_iOHIDDeviceCloseErr, frameworkHandle, "IOHIDDeviceClose", "10.5")
	registerFunc(&_iOHIDDeviceConformsTo, &_iOHIDDeviceConformsToErr, frameworkHandle, "IOHIDDeviceConformsTo", "10.5")
	registerFunc(&_iOHIDDeviceCopyMatchingElements, &_iOHIDDeviceCopyMatchingElementsErr, frameworkHandle, "IOHIDDeviceCopyMatchingElements", "10.5")
	registerFunc(&_iOHIDDeviceCopyValueMultiple, &_iOHIDDeviceCopyValueMultipleErr, frameworkHandle, "IOHIDDeviceCopyValueMultiple", "10.5")
	registerFunc(&_iOHIDDeviceCopyValueMultipleWithCallback, &_iOHIDDeviceCopyValueMultipleWithCallbackErr, frameworkHandle, "IOHIDDeviceCopyValueMultipleWithCallback", "10.5")
	registerFunc(&_iOHIDDeviceCreate, &_iOHIDDeviceCreateErr, frameworkHandle, "IOHIDDeviceCreate", "10.5")
	registerFunc(&_iOHIDDeviceGetProperty, &_iOHIDDeviceGetPropertyErr, frameworkHandle, "IOHIDDeviceGetProperty", "10.5")
	registerFunc(&_iOHIDDeviceGetReport, &_iOHIDDeviceGetReportErr, frameworkHandle, "IOHIDDeviceGetReport", "10.5")
	registerFunc(&_iOHIDDeviceGetReportWithCallback, &_iOHIDDeviceGetReportWithCallbackErr, frameworkHandle, "IOHIDDeviceGetReportWithCallback", "10.5")
	registerFunc(&_iOHIDDeviceGetService, &_iOHIDDeviceGetServiceErr, frameworkHandle, "IOHIDDeviceGetService", "10.6")
	registerFunc(&_iOHIDDeviceGetTypeID, &_iOHIDDeviceGetTypeIDErr, frameworkHandle, "IOHIDDeviceGetTypeID", "10.5")
	registerFunc(&_iOHIDDeviceGetValue, &_iOHIDDeviceGetValueErr, frameworkHandle, "IOHIDDeviceGetValue", "10.5")
	registerFunc(&_iOHIDDeviceGetValueWithCallback, &_iOHIDDeviceGetValueWithCallbackErr, frameworkHandle, "IOHIDDeviceGetValueWithCallback", "10.5")
	registerFunc(&_iOHIDDeviceGetValueWithOptions, &_iOHIDDeviceGetValueWithOptionsErr, frameworkHandle, "IOHIDDeviceGetValueWithOptions", "10.13")
	registerFunc(&_iOHIDDeviceOpen, &_iOHIDDeviceOpenErr, frameworkHandle, "IOHIDDeviceOpen", "10.5")
	registerFunc(&_iOHIDDeviceRegisterInputReportCallback, &_iOHIDDeviceRegisterInputReportCallbackErr, frameworkHandle, "IOHIDDeviceRegisterInputReportCallback", "10.5")
	registerFunc(&_iOHIDDeviceRegisterInputReportWithTimeStampCallback, &_iOHIDDeviceRegisterInputReportWithTimeStampCallbackErr, frameworkHandle, "IOHIDDeviceRegisterInputReportWithTimeStampCallback", "10.10")
	registerFunc(&_iOHIDDeviceRegisterInputValueCallback, &_iOHIDDeviceRegisterInputValueCallbackErr, frameworkHandle, "IOHIDDeviceRegisterInputValueCallback", "10.5")
	registerFunc(&_iOHIDDeviceRegisterRemovalCallback, &_iOHIDDeviceRegisterRemovalCallbackErr, frameworkHandle, "IOHIDDeviceRegisterRemovalCallback", "10.5")
	registerFunc(&_iOHIDDeviceScheduleWithRunLoop, &_iOHIDDeviceScheduleWithRunLoopErr, frameworkHandle, "IOHIDDeviceScheduleWithRunLoop", "10.5")
	registerFunc(&_iOHIDDeviceSetCancelHandler, &_iOHIDDeviceSetCancelHandlerErr, frameworkHandle, "IOHIDDeviceSetCancelHandler", "10.15")
	registerFunc(&_iOHIDDeviceSetDispatchQueue, &_iOHIDDeviceSetDispatchQueueErr, frameworkHandle, "IOHIDDeviceSetDispatchQueue", "10.15")
	registerFunc(&_iOHIDDeviceSetInputValueMatching, &_iOHIDDeviceSetInputValueMatchingErr, frameworkHandle, "IOHIDDeviceSetInputValueMatching", "10.5")
	registerFunc(&_iOHIDDeviceSetInputValueMatchingMultiple, &_iOHIDDeviceSetInputValueMatchingMultipleErr, frameworkHandle, "IOHIDDeviceSetInputValueMatchingMultiple", "10.5")
	registerFunc(&_iOHIDDeviceSetProperty, &_iOHIDDeviceSetPropertyErr, frameworkHandle, "IOHIDDeviceSetProperty", "10.5")
	registerFunc(&_iOHIDDeviceSetReport, &_iOHIDDeviceSetReportErr, frameworkHandle, "IOHIDDeviceSetReport", "10.5")
	registerFunc(&_iOHIDDeviceSetReportWithCallback, &_iOHIDDeviceSetReportWithCallbackErr, frameworkHandle, "IOHIDDeviceSetReportWithCallback", "10.5")
	registerFunc(&_iOHIDDeviceSetValue, &_iOHIDDeviceSetValueErr, frameworkHandle, "IOHIDDeviceSetValue", "10.5")
	registerFunc(&_iOHIDDeviceSetValueMultiple, &_iOHIDDeviceSetValueMultipleErr, frameworkHandle, "IOHIDDeviceSetValueMultiple", "10.5")
	registerFunc(&_iOHIDDeviceSetValueMultipleWithCallback, &_iOHIDDeviceSetValueMultipleWithCallbackErr, frameworkHandle, "IOHIDDeviceSetValueMultipleWithCallback", "10.5")
	registerFunc(&_iOHIDDeviceSetValueWithCallback, &_iOHIDDeviceSetValueWithCallbackErr, frameworkHandle, "IOHIDDeviceSetValueWithCallback", "10.5")
	registerFunc(&_iOHIDDeviceUnscheduleFromRunLoop, &_iOHIDDeviceUnscheduleFromRunLoopErr, frameworkHandle, "IOHIDDeviceUnscheduleFromRunLoop", "10.5")
	registerFunc(&_iOHIDElementAttach, &_iOHIDElementAttachErr, frameworkHandle, "IOHIDElementAttach", "10.5")
	registerFunc(&_iOHIDElementCopyAttached, &_iOHIDElementCopyAttachedErr, frameworkHandle, "IOHIDElementCopyAttached", "10.5")
	registerFunc(&_iOHIDElementCreateWithDictionary, &_iOHIDElementCreateWithDictionaryErr, frameworkHandle, "IOHIDElementCreateWithDictionary", "10.5")
	registerFunc(&_iOHIDElementDetach, &_iOHIDElementDetachErr, frameworkHandle, "IOHIDElementDetach", "10.5")
	registerFunc(&_iOHIDElementGetChildren, &_iOHIDElementGetChildrenErr, frameworkHandle, "IOHIDElementGetChildren", "10.5")
	registerFunc(&_iOHIDElementGetCollectionType, &_iOHIDElementGetCollectionTypeErr, frameworkHandle, "IOHIDElementGetCollectionType", "10.5")
	registerFunc(&_iOHIDElementGetCookie, &_iOHIDElementGetCookieErr, frameworkHandle, "IOHIDElementGetCookie", "10.5")
	registerFunc(&_iOHIDElementGetDevice, &_iOHIDElementGetDeviceErr, frameworkHandle, "IOHIDElementGetDevice", "10.5")
	registerFunc(&_iOHIDElementGetLogicalMax, &_iOHIDElementGetLogicalMaxErr, frameworkHandle, "IOHIDElementGetLogicalMax", "10.5")
	registerFunc(&_iOHIDElementGetLogicalMin, &_iOHIDElementGetLogicalMinErr, frameworkHandle, "IOHIDElementGetLogicalMin", "10.5")
	registerFunc(&_iOHIDElementGetName, &_iOHIDElementGetNameErr, frameworkHandle, "IOHIDElementGetName", "10.5")
	registerFunc(&_iOHIDElementGetParent, &_iOHIDElementGetParentErr, frameworkHandle, "IOHIDElementGetParent", "10.5")
	registerFunc(&_iOHIDElementGetPhysicalMax, &_iOHIDElementGetPhysicalMaxErr, frameworkHandle, "IOHIDElementGetPhysicalMax", "10.5")
	registerFunc(&_iOHIDElementGetPhysicalMin, &_iOHIDElementGetPhysicalMinErr, frameworkHandle, "IOHIDElementGetPhysicalMin", "10.5")
	registerFunc(&_iOHIDElementGetProperty, &_iOHIDElementGetPropertyErr, frameworkHandle, "IOHIDElementGetProperty", "10.5")
	registerFunc(&_iOHIDElementGetReportCount, &_iOHIDElementGetReportCountErr, frameworkHandle, "IOHIDElementGetReportCount", "10.5")
	registerFunc(&_iOHIDElementGetReportID, &_iOHIDElementGetReportIDErr, frameworkHandle, "IOHIDElementGetReportID", "10.5")
	registerFunc(&_iOHIDElementGetReportSize, &_iOHIDElementGetReportSizeErr, frameworkHandle, "IOHIDElementGetReportSize", "10.5")
	registerFunc(&_iOHIDElementGetType, &_iOHIDElementGetTypeErr, frameworkHandle, "IOHIDElementGetType", "10.5")
	registerFunc(&_iOHIDElementGetTypeID, &_iOHIDElementGetTypeIDErr, frameworkHandle, "IOHIDElementGetTypeID", "10.5")
	registerFunc(&_iOHIDElementGetUnit, &_iOHIDElementGetUnitErr, frameworkHandle, "IOHIDElementGetUnit", "10.5")
	registerFunc(&_iOHIDElementGetUnitExponent, &_iOHIDElementGetUnitExponentErr, frameworkHandle, "IOHIDElementGetUnitExponent", "10.5")
	registerFunc(&_iOHIDElementGetUsage, &_iOHIDElementGetUsageErr, frameworkHandle, "IOHIDElementGetUsage", "10.5")
	registerFunc(&_iOHIDElementGetUsagePage, &_iOHIDElementGetUsagePageErr, frameworkHandle, "IOHIDElementGetUsagePage", "10.5")
	registerFunc(&_iOHIDElementHasNullState, &_iOHIDElementHasNullStateErr, frameworkHandle, "IOHIDElementHasNullState", "10.5")
	registerFunc(&_iOHIDElementHasPreferredState, &_iOHIDElementHasPreferredStateErr, frameworkHandle, "IOHIDElementHasPreferredState", "10.5")
	registerFunc(&_iOHIDElementIsArray, &_iOHIDElementIsArrayErr, frameworkHandle, "IOHIDElementIsArray", "10.5")
	registerFunc(&_iOHIDElementIsNonLinear, &_iOHIDElementIsNonLinearErr, frameworkHandle, "IOHIDElementIsNonLinear", "10.5")
	registerFunc(&_iOHIDElementIsRelative, &_iOHIDElementIsRelativeErr, frameworkHandle, "IOHIDElementIsRelative", "10.5")
	registerFunc(&_iOHIDElementIsVirtual, &_iOHIDElementIsVirtualErr, frameworkHandle, "IOHIDElementIsVirtual", "10.5")
	registerFunc(&_iOHIDElementIsWrapping, &_iOHIDElementIsWrappingErr, frameworkHandle, "IOHIDElementIsWrapping", "10.5")
	registerFunc(&_iOHIDElementSetProperty, &_iOHIDElementSetPropertyErr, frameworkHandle, "IOHIDElementSetProperty", "10.5")
	registerFunc(&_iOHIDEventSystemClientCopyProperty, &_iOHIDEventSystemClientCopyPropertyErr, frameworkHandle, "IOHIDEventSystemClientCopyProperty", "10.12")
	registerFunc(&_iOHIDEventSystemClientCopyServices, &_iOHIDEventSystemClientCopyServicesErr, frameworkHandle, "IOHIDEventSystemClientCopyServices", "10.12")
	registerFunc(&_iOHIDEventSystemClientCreateSimpleClient, &_iOHIDEventSystemClientCreateSimpleClientErr, frameworkHandle, "IOHIDEventSystemClientCreateSimpleClient", "10.12")
	registerFunc(&_iOHIDEventSystemClientGetTypeID, &_iOHIDEventSystemClientGetTypeIDErr, frameworkHandle, "IOHIDEventSystemClientGetTypeID", "10.12")
	registerFunc(&_iOHIDEventSystemClientSetProperty, &_iOHIDEventSystemClientSetPropertyErr, frameworkHandle, "IOHIDEventSystemClientSetProperty", "10.12")
	registerFunc(&_iOHIDGetAccelerationWithKey, &_iOHIDGetAccelerationWithKeyErr, frameworkHandle, "IOHIDGetAccelerationWithKey", "10.0")
	registerFunc(&_iOHIDGetActivityState, &_iOHIDGetActivityStateErr, frameworkHandle, "IOHIDGetActivityState", "10.9")
	registerFunc(&_iOHIDGetButtonEventNum, &_iOHIDGetButtonEventNumErr, frameworkHandle, "IOHIDGetButtonEventNum", "10.0")
	registerFunc(&_iOHIDGetModifierLockState, &_iOHIDGetModifierLockStateErr, frameworkHandle, "IOHIDGetModifierLockState", "10.6")
	registerFunc(&_iOHIDGetMouseAcceleration, &_iOHIDGetMouseAccelerationErr, frameworkHandle, "IOHIDGetMouseAcceleration", "10.0")
	registerFunc(&_iOHIDGetMouseButtonMode, &_iOHIDGetMouseButtonModeErr, frameworkHandle, "IOHIDGetMouseButtonMode", "10.2")
	registerFunc(&_iOHIDGetParameter, &_iOHIDGetParameterErr, frameworkHandle, "IOHIDGetParameter", "10.0")
	registerFunc(&_iOHIDGetScrollAcceleration, &_iOHIDGetScrollAccelerationErr, frameworkHandle, "IOHIDGetScrollAcceleration", "10.0")
	registerFunc(&_iOHIDGetStateForSelector, &_iOHIDGetStateForSelectorErr, frameworkHandle, "IOHIDGetStateForSelector", "10.9")
	registerFunc(&_iOHIDManagerActivate, &_iOHIDManagerActivateErr, frameworkHandle, "IOHIDManagerActivate", "10.15")
	registerFunc(&_iOHIDManagerCancel, &_iOHIDManagerCancelErr, frameworkHandle, "IOHIDManagerCancel", "10.15")
	registerFunc(&_iOHIDManagerClose, &_iOHIDManagerCloseErr, frameworkHandle, "IOHIDManagerClose", "10.5")
	registerFunc(&_iOHIDManagerCopyDevices, &_iOHIDManagerCopyDevicesErr, frameworkHandle, "IOHIDManagerCopyDevices", "10.5")
	registerFunc(&_iOHIDManagerCreate, &_iOHIDManagerCreateErr, frameworkHandle, "IOHIDManagerCreate", "10.5")
	registerFunc(&_iOHIDManagerGetProperty, &_iOHIDManagerGetPropertyErr, frameworkHandle, "IOHIDManagerGetProperty", "10.5")
	registerFunc(&_iOHIDManagerGetTypeID, &_iOHIDManagerGetTypeIDErr, frameworkHandle, "IOHIDManagerGetTypeID", "10.5")
	registerFunc(&_iOHIDManagerOpen, &_iOHIDManagerOpenErr, frameworkHandle, "IOHIDManagerOpen", "10.5")
	registerFunc(&_iOHIDManagerRegisterDeviceMatchingCallback, &_iOHIDManagerRegisterDeviceMatchingCallbackErr, frameworkHandle, "IOHIDManagerRegisterDeviceMatchingCallback", "10.5")
	registerFunc(&_iOHIDManagerRegisterDeviceRemovalCallback, &_iOHIDManagerRegisterDeviceRemovalCallbackErr, frameworkHandle, "IOHIDManagerRegisterDeviceRemovalCallback", "10.5")
	registerFunc(&_iOHIDManagerRegisterInputReportCallback, &_iOHIDManagerRegisterInputReportCallbackErr, frameworkHandle, "IOHIDManagerRegisterInputReportCallback", "10.5")
	registerFunc(&_iOHIDManagerRegisterInputReportWithTimeStampCallback, &_iOHIDManagerRegisterInputReportWithTimeStampCallbackErr, frameworkHandle, "IOHIDManagerRegisterInputReportWithTimeStampCallback", "10.15")
	registerFunc(&_iOHIDManagerRegisterInputValueCallback, &_iOHIDManagerRegisterInputValueCallbackErr, frameworkHandle, "IOHIDManagerRegisterInputValueCallback", "10.5")
	registerFunc(&_iOHIDManagerSaveToPropertyDomain, &_iOHIDManagerSaveToPropertyDomainErr, frameworkHandle, "IOHIDManagerSaveToPropertyDomain", "10.6")
	registerFunc(&_iOHIDManagerScheduleWithRunLoop, &_iOHIDManagerScheduleWithRunLoopErr, frameworkHandle, "IOHIDManagerScheduleWithRunLoop", "10.5")
	registerFunc(&_iOHIDManagerSetCancelHandler, &_iOHIDManagerSetCancelHandlerErr, frameworkHandle, "IOHIDManagerSetCancelHandler", "10.15")
	registerFunc(&_iOHIDManagerSetDeviceMatching, &_iOHIDManagerSetDeviceMatchingErr, frameworkHandle, "IOHIDManagerSetDeviceMatching", "10.5")
	registerFunc(&_iOHIDManagerSetDeviceMatchingMultiple, &_iOHIDManagerSetDeviceMatchingMultipleErr, frameworkHandle, "IOHIDManagerSetDeviceMatchingMultiple", "10.5")
	registerFunc(&_iOHIDManagerSetDispatchQueue, &_iOHIDManagerSetDispatchQueueErr, frameworkHandle, "IOHIDManagerSetDispatchQueue", "10.15")
	registerFunc(&_iOHIDManagerSetInputValueMatching, &_iOHIDManagerSetInputValueMatchingErr, frameworkHandle, "IOHIDManagerSetInputValueMatching", "10.5")
	registerFunc(&_iOHIDManagerSetInputValueMatchingMultiple, &_iOHIDManagerSetInputValueMatchingMultipleErr, frameworkHandle, "IOHIDManagerSetInputValueMatchingMultiple", "10.5")
	registerFunc(&_iOHIDManagerSetProperty, &_iOHIDManagerSetPropertyErr, frameworkHandle, "IOHIDManagerSetProperty", "10.5")
	registerFunc(&_iOHIDManagerUnscheduleFromRunLoop, &_iOHIDManagerUnscheduleFromRunLoopErr, frameworkHandle, "IOHIDManagerUnscheduleFromRunLoop", "10.5")
	registerFunc(&_iOHIDPostEvent, &_iOHIDPostEventErr, frameworkHandle, "IOHIDPostEvent", "10.0")
	registerFunc(&_iOHIDQueueActivate, &_iOHIDQueueActivateErr, frameworkHandle, "IOHIDQueueActivate", "10.15")
	registerFunc(&_iOHIDQueueAddElement, &_iOHIDQueueAddElementErr, frameworkHandle, "IOHIDQueueAddElement", "10.5")
	registerFunc(&_iOHIDQueueCancel, &_iOHIDQueueCancelErr, frameworkHandle, "IOHIDQueueCancel", "10.15")
	registerFunc(&_iOHIDQueueContainsElement, &_iOHIDQueueContainsElementErr, frameworkHandle, "IOHIDQueueContainsElement", "10.5")
	registerFunc(&_iOHIDQueueCopyNextValue, &_iOHIDQueueCopyNextValueErr, frameworkHandle, "IOHIDQueueCopyNextValue", "10.5")
	registerFunc(&_iOHIDQueueCopyNextValueWithTimeout, &_iOHIDQueueCopyNextValueWithTimeoutErr, frameworkHandle, "IOHIDQueueCopyNextValueWithTimeout", "10.5")
	registerFunc(&_iOHIDQueueCreate, &_iOHIDQueueCreateErr, frameworkHandle, "IOHIDQueueCreate", "10.5")
	registerFunc(&_iOHIDQueueGetDepth, &_iOHIDQueueGetDepthErr, frameworkHandle, "IOHIDQueueGetDepth", "10.5")
	registerFunc(&_iOHIDQueueGetDevice, &_iOHIDQueueGetDeviceErr, frameworkHandle, "IOHIDQueueGetDevice", "10.5")
	registerFunc(&_iOHIDQueueGetTypeID, &_iOHIDQueueGetTypeIDErr, frameworkHandle, "IOHIDQueueGetTypeID", "10.5")
	registerFunc(&_iOHIDQueueRegisterValueAvailableCallback, &_iOHIDQueueRegisterValueAvailableCallbackErr, frameworkHandle, "IOHIDQueueRegisterValueAvailableCallback", "10.5")
	registerFunc(&_iOHIDQueueRemoveElement, &_iOHIDQueueRemoveElementErr, frameworkHandle, "IOHIDQueueRemoveElement", "10.5")
	registerFunc(&_iOHIDQueueScheduleWithRunLoop, &_iOHIDQueueScheduleWithRunLoopErr, frameworkHandle, "IOHIDQueueScheduleWithRunLoop", "10.5")
	registerFunc(&_iOHIDQueueSetCancelHandler, &_iOHIDQueueSetCancelHandlerErr, frameworkHandle, "IOHIDQueueSetCancelHandler", "10.15")
	registerFunc(&_iOHIDQueueSetDepth, &_iOHIDQueueSetDepthErr, frameworkHandle, "IOHIDQueueSetDepth", "10.5")
	registerFunc(&_iOHIDQueueSetDispatchQueue, &_iOHIDQueueSetDispatchQueueErr, frameworkHandle, "IOHIDQueueSetDispatchQueue", "10.15")
	registerFunc(&_iOHIDQueueStart, &_iOHIDQueueStartErr, frameworkHandle, "IOHIDQueueStart", "10.5")
	registerFunc(&_iOHIDQueueStop, &_iOHIDQueueStopErr, frameworkHandle, "IOHIDQueueStop", "10.5")
	registerFunc(&_iOHIDQueueUnscheduleFromRunLoop, &_iOHIDQueueUnscheduleFromRunLoopErr, frameworkHandle, "IOHIDQueueUnscheduleFromRunLoop", "10.5")
	registerFunc(&_iOHIDRegisterVirtualDisplay, &_iOHIDRegisterVirtualDisplayErr, frameworkHandle, "IOHIDRegisterVirtualDisplay", "10.0")
	registerFunc(&_iOHIDRequestAccess, &_iOHIDRequestAccessErr, frameworkHandle, "IOHIDRequestAccess", "10.15")
	registerFunc(&_iOHIDServiceClientConformsTo, &_iOHIDServiceClientConformsToErr, frameworkHandle, "IOHIDServiceClientConformsTo", "10.12")
	registerFunc(&_iOHIDServiceClientCopyProperty, &_iOHIDServiceClientCopyPropertyErr, frameworkHandle, "IOHIDServiceClientCopyProperty", "10.12")
	registerFunc(&_iOHIDServiceClientGetRegistryID, &_iOHIDServiceClientGetRegistryIDErr, frameworkHandle, "IOHIDServiceClientGetRegistryID", "10.12")
	registerFunc(&_iOHIDServiceClientGetTypeID, &_iOHIDServiceClientGetTypeIDErr, frameworkHandle, "IOHIDServiceClientGetTypeID", "10.12")
	registerFunc(&_iOHIDServiceClientSetProperty, &_iOHIDServiceClientSetPropertyErr, frameworkHandle, "IOHIDServiceClientSetProperty", "10.12")
	registerFunc(&_iOHIDSetAccelerationWithKey, &_iOHIDSetAccelerationWithKeyErr, frameworkHandle, "IOHIDSetAccelerationWithKey", "10.0")
	registerFunc(&_iOHIDSetCFTypeParameter, &_iOHIDSetCFTypeParameterErr, frameworkHandle, "IOHIDSetCFTypeParameter", "10.3")
	registerFunc(&_iOHIDSetCursorEnable, &_iOHIDSetCursorEnableErr, frameworkHandle, "IOHIDSetCursorEnable", "10.0")
	registerFunc(&_iOHIDSetEventsEnable, &_iOHIDSetEventsEnableErr, frameworkHandle, "IOHIDSetEventsEnable", "10.0")
	registerFunc(&_iOHIDSetModifierLockState, &_iOHIDSetModifierLockStateErr, frameworkHandle, "IOHIDSetModifierLockState", "10.6")
	registerFunc(&_iOHIDSetMouseAcceleration, &_iOHIDSetMouseAccelerationErr, frameworkHandle, "IOHIDSetMouseAcceleration", "10.0")
	registerFunc(&_iOHIDSetMouseButtonMode, &_iOHIDSetMouseButtonModeErr, frameworkHandle, "IOHIDSetMouseButtonMode", "10.0")
	registerFunc(&_iOHIDSetMouseLocation, &_iOHIDSetMouseLocationErr, frameworkHandle, "IOHIDSetMouseLocation", "10.0")
	registerFunc(&_iOHIDSetParameter, &_iOHIDSetParameterErr, frameworkHandle, "IOHIDSetParameter", "10.0")
	registerFunc(&_iOHIDSetScrollAcceleration, &_iOHIDSetScrollAccelerationErr, frameworkHandle, "IOHIDSetScrollAcceleration", "10.0")
	registerFunc(&_iOHIDSetStateForSelector, &_iOHIDSetStateForSelectorErr, frameworkHandle, "IOHIDSetStateForSelector", "10.9")
	registerFunc(&_iOHIDSetVirtualDisplayBounds, &_iOHIDSetVirtualDisplayBoundsErr, frameworkHandle, "IOHIDSetVirtualDisplayBounds", "10.0")
	registerFunc(&_iOHIDTransactionAddElement, &_iOHIDTransactionAddElementErr, frameworkHandle, "IOHIDTransactionAddElement", "10.5")
	registerFunc(&_iOHIDTransactionClear, &_iOHIDTransactionClearErr, frameworkHandle, "IOHIDTransactionClear", "10.5")
	registerFunc(&_iOHIDTransactionCommit, &_iOHIDTransactionCommitErr, frameworkHandle, "IOHIDTransactionCommit", "10.5")
	registerFunc(&_iOHIDTransactionCommitWithCallback, &_iOHIDTransactionCommitWithCallbackErr, frameworkHandle, "IOHIDTransactionCommitWithCallback", "10.5")
	registerFunc(&_iOHIDTransactionContainsElement, &_iOHIDTransactionContainsElementErr, frameworkHandle, "IOHIDTransactionContainsElement", "10.5")
	registerFunc(&_iOHIDTransactionCreate, &_iOHIDTransactionCreateErr, frameworkHandle, "IOHIDTransactionCreate", "10.5")
	registerFunc(&_iOHIDTransactionGetDevice, &_iOHIDTransactionGetDeviceErr, frameworkHandle, "IOHIDTransactionGetDevice", "10.5")
	registerFunc(&_iOHIDTransactionGetDirection, &_iOHIDTransactionGetDirectionErr, frameworkHandle, "IOHIDTransactionGetDirection", "10.5")
	registerFunc(&_iOHIDTransactionGetTypeID, &_iOHIDTransactionGetTypeIDErr, frameworkHandle, "IOHIDTransactionGetTypeID", "10.5")
	registerFunc(&_iOHIDTransactionGetValue, &_iOHIDTransactionGetValueErr, frameworkHandle, "IOHIDTransactionGetValue", "10.5")
	registerFunc(&_iOHIDTransactionRemoveElement, &_iOHIDTransactionRemoveElementErr, frameworkHandle, "IOHIDTransactionRemoveElement", "10.5")
	registerFunc(&_iOHIDTransactionScheduleWithRunLoop, &_iOHIDTransactionScheduleWithRunLoopErr, frameworkHandle, "IOHIDTransactionScheduleWithRunLoop", "10.5")
	registerFunc(&_iOHIDTransactionSetDirection, &_iOHIDTransactionSetDirectionErr, frameworkHandle, "IOHIDTransactionSetDirection", "10.5")
	registerFunc(&_iOHIDTransactionSetValue, &_iOHIDTransactionSetValueErr, frameworkHandle, "IOHIDTransactionSetValue", "10.5")
	registerFunc(&_iOHIDTransactionUnscheduleFromRunLoop, &_iOHIDTransactionUnscheduleFromRunLoopErr, frameworkHandle, "IOHIDTransactionUnscheduleFromRunLoop", "10.5")
	registerFunc(&_iOHIDUnregisterVirtualDisplay, &_iOHIDUnregisterVirtualDisplayErr, frameworkHandle, "IOHIDUnregisterVirtualDisplay", "10.0")
	registerFunc(&_iOHIDUserDeviceActivate, &_iOHIDUserDeviceActivateErr, frameworkHandle, "IOHIDUserDeviceActivate", "10.15")
	registerFunc(&_iOHIDUserDeviceCancel, &_iOHIDUserDeviceCancelErr, frameworkHandle, "IOHIDUserDeviceCancel", "10.15")
	registerFunc(&_iOHIDUserDeviceCopyProperty, &_iOHIDUserDeviceCopyPropertyErr, frameworkHandle, "IOHIDUserDeviceCopyProperty", "10.15")
	registerFunc(&_iOHIDUserDeviceCreateWithProperties, &_iOHIDUserDeviceCreateWithPropertiesErr, frameworkHandle, "IOHIDUserDeviceCreateWithProperties", "10.15")
	registerFunc(&_iOHIDUserDeviceGetTypeID, &_iOHIDUserDeviceGetTypeIDErr, frameworkHandle, "IOHIDUserDeviceGetTypeID", "10.15")
	registerFunc(&_iOHIDUserDeviceHandleReportWithTimeStamp, &_iOHIDUserDeviceHandleReportWithTimeStampErr, frameworkHandle, "IOHIDUserDeviceHandleReportWithTimeStamp", "10.15")
	registerFunc(&_iOHIDUserDeviceRegisterGetReportBlock, &_iOHIDUserDeviceRegisterGetReportBlockErr, frameworkHandle, "IOHIDUserDeviceRegisterGetReportBlock", "10.15")
	registerFunc(&_iOHIDUserDeviceRegisterSetReportBlock, &_iOHIDUserDeviceRegisterSetReportBlockErr, frameworkHandle, "IOHIDUserDeviceRegisterSetReportBlock", "10.15")
	registerFunc(&_iOHIDUserDeviceSetCancelHandler, &_iOHIDUserDeviceSetCancelHandlerErr, frameworkHandle, "IOHIDUserDeviceSetCancelHandler", "10.15")
	registerFunc(&_iOHIDUserDeviceSetDispatchQueue, &_iOHIDUserDeviceSetDispatchQueueErr, frameworkHandle, "IOHIDUserDeviceSetDispatchQueue", "10.15")
	registerFunc(&_iOHIDUserDeviceSetProperty, &_iOHIDUserDeviceSetPropertyErr, frameworkHandle, "IOHIDUserDeviceSetProperty", "10.15")
	registerFunc(&_iOHIDValueCreateWithBytes, &_iOHIDValueCreateWithBytesErr, frameworkHandle, "IOHIDValueCreateWithBytes", "10.5")
	registerFunc(&_iOHIDValueCreateWithBytesNoCopy, &_iOHIDValueCreateWithBytesNoCopyErr, frameworkHandle, "IOHIDValueCreateWithBytesNoCopy", "10.5")
	registerFunc(&_iOHIDValueCreateWithIntegerValue, &_iOHIDValueCreateWithIntegerValueErr, frameworkHandle, "IOHIDValueCreateWithIntegerValue", "10.5")
	registerFunc(&_iOHIDValueGetBytePtr, &_iOHIDValueGetBytePtrErr, frameworkHandle, "IOHIDValueGetBytePtr", "10.5")
	registerFunc(&_iOHIDValueGetElement, &_iOHIDValueGetElementErr, frameworkHandle, "IOHIDValueGetElement", "10.5")
	registerFunc(&_iOHIDValueGetIntegerValue, &_iOHIDValueGetIntegerValueErr, frameworkHandle, "IOHIDValueGetIntegerValue", "10.5")
	registerFunc(&_iOHIDValueGetLength, &_iOHIDValueGetLengthErr, frameworkHandle, "IOHIDValueGetLength", "10.5")
	registerFunc(&_iOHIDValueGetScaledValue, &_iOHIDValueGetScaledValueErr, frameworkHandle, "IOHIDValueGetScaledValue", "10.5")
	registerFunc(&_iOHIDValueGetTimeStamp, &_iOHIDValueGetTimeStampErr, frameworkHandle, "IOHIDValueGetTimeStamp", "10.5")
	registerFunc(&_iOHIDValueGetTypeID, &_iOHIDValueGetTypeIDErr, frameworkHandle, "IOHIDValueGetTypeID", "10.5")
	registerFunc(&_iOI2CCopyInterfaceForID, &_iOI2CCopyInterfaceForIDErr, frameworkHandle, "IOI2CCopyInterfaceForID", "10.3")
	registerFunc(&_iOI2CInterfaceClose, &_iOI2CInterfaceCloseErr, frameworkHandle, "IOI2CInterfaceClose", "10.3")
	registerFunc(&_iOI2CInterfaceOpen, &_iOI2CInterfaceOpenErr, frameworkHandle, "IOI2CInterfaceOpen", "10.3")
	registerFunc(&_iOI2CSendRequest, &_iOI2CSendRequestErr, frameworkHandle, "IOI2CSendRequest", "10.3")
	registerFunc(&_iOIteratorIsValid, &_iOIteratorIsValidErr, frameworkHandle, "IOIteratorIsValid", "10.0")
	registerFunc(&_iOIteratorNext, &_iOIteratorNextErr, frameworkHandle, "IOIteratorNext", "10.0")
	registerFunc(&_iOIteratorReset, &_iOIteratorResetErr, frameworkHandle, "IOIteratorReset", "10.0")
	registerFunc(&_iOKitGetBusyState, &_iOKitGetBusyStateErr, frameworkHandle, "IOKitGetBusyState", "10.0")
	registerFunc(&_iOKitWaitQuiet, &_iOKitWaitQuietErr, frameworkHandle, "IOKitWaitQuiet", "10.0")
	registerFunc(&_iOMainPort, &_iOMainPortErr, frameworkHandle, "IOMainPort", "12.0")
	registerFunc(&_iOMasterPort, &_iOMasterPortErr, frameworkHandle, "IOMasterPort", "10.0")
	registerFunc(&_iONetworkClose, &_iONetworkCloseErr, frameworkHandle, "IONetworkClose", "10.0")
	registerFunc(&_iONetworkGetDataCapacity, &_iONetworkGetDataCapacityErr, frameworkHandle, "IONetworkGetDataCapacity", "10.0")
	registerFunc(&_iONetworkGetDataHandle, &_iONetworkGetDataHandleErr, frameworkHandle, "IONetworkGetDataHandle", "10.0")
	registerFunc(&_iONetworkGetPacketFiltersMask, &_iONetworkGetPacketFiltersMaskErr, frameworkHandle, "IONetworkGetPacketFiltersMask", "10.1")
	registerFunc(&_iONetworkOpen, &_iONetworkOpenErr, frameworkHandle, "IONetworkOpen", "10.0")
	registerFunc(&_iONetworkReadData, &_iONetworkReadDataErr, frameworkHandle, "IONetworkReadData", "10.0")
	registerFunc(&_iONetworkResetData, &_iONetworkResetDataErr, frameworkHandle, "IONetworkResetData", "10.0")
	registerFunc(&_iONetworkSetPacketFiltersMask, &_iONetworkSetPacketFiltersMaskErr, frameworkHandle, "IONetworkSetPacketFiltersMask", "10.1")
	registerFunc(&_iONetworkWriteData, &_iONetworkWriteDataErr, frameworkHandle, "IONetworkWriteData", "10.0")
	registerFunc(&_iONotificationPortCreate, &_iONotificationPortCreateErr, frameworkHandle, "IONotificationPortCreate", "10.0")
	registerFunc(&_iONotificationPortDestroy, &_iONotificationPortDestroyErr, frameworkHandle, "IONotificationPortDestroy", "10.0")
	registerFunc(&_iONotificationPortGetMachPort, &_iONotificationPortGetMachPortErr, frameworkHandle, "IONotificationPortGetMachPort", "10.0")
	registerFunc(&_iONotificationPortGetRunLoopSource, &_iONotificationPortGetRunLoopSourceErr, frameworkHandle, "IONotificationPortGetRunLoopSource", "10.0")
	registerFunc(&_iONotificationPortSetDispatchQueue, &_iONotificationPortSetDispatchQueueErr, frameworkHandle, "IONotificationPortSetDispatchQueue", "10.6")
	registerFunc(&_iONotificationPortSetImportanceReceiver, &_iONotificationPortSetImportanceReceiverErr, frameworkHandle, "IONotificationPortSetImportanceReceiver", "10.13")
	registerFunc(&_iOObjectConformsTo, &_iOObjectConformsToErr, frameworkHandle, "IOObjectConformsTo", "10.0")
	registerFunc(&_iOObjectCopyBundleIdentifierForClass, &_iOObjectCopyBundleIdentifierForClassErr, frameworkHandle, "IOObjectCopyBundleIdentifierForClass", "10.4")
	registerFunc(&_iOObjectCopyClass, &_iOObjectCopyClassErr, frameworkHandle, "IOObjectCopyClass", "10.4")
	registerFunc(&_iOObjectCopySuperclassForClass, &_iOObjectCopySuperclassForClassErr, frameworkHandle, "IOObjectCopySuperclassForClass", "10.4")
	registerFunc(&_iOObjectGetClass, &_iOObjectGetClassErr, frameworkHandle, "IOObjectGetClass", "10.0")
	registerFunc(&_iOObjectGetKernelRetainCount, &_iOObjectGetKernelRetainCountErr, frameworkHandle, "IOObjectGetKernelRetainCount", "10.6")
	registerFunc(&_iOObjectGetRetainCount, &_iOObjectGetRetainCountErr, frameworkHandle, "IOObjectGetRetainCount", "10.0")
	registerFunc(&_iOObjectGetUserRetainCount, &_iOObjectGetUserRetainCountErr, frameworkHandle, "IOObjectGetUserRetainCount", "10.6")
	registerFunc(&_iOObjectIsEqualTo, &_iOObjectIsEqualToErr, frameworkHandle, "IOObjectIsEqualTo", "10.0")
	registerFunc(&_iOObjectRelease, &_iOObjectReleaseErr, frameworkHandle, "IOObjectRelease", "10.0")
	registerFunc(&_iOObjectRetain, &_iOObjectRetainErr, frameworkHandle, "IOObjectRetain", "10.1")
	registerFunc(&_iOOpenFirmwarePathMatching, &_iOOpenFirmwarePathMatchingErr, frameworkHandle, "IOOpenFirmwarePathMatching", "10.0")
	registerFunc(&_iOPMAssertionCopyProperties, &_iOPMAssertionCopyPropertiesErr, frameworkHandle, "IOPMAssertionCopyProperties", "10.7")
	registerFunc(&_iOPMAssertionCreate, &_iOPMAssertionCreateErr, frameworkHandle, "IOPMAssertionCreate", "10.5")
	registerFunc(&_iOPMAssertionCreateWithDescription, &_iOPMAssertionCreateWithDescriptionErr, frameworkHandle, "IOPMAssertionCreateWithDescription", "10.7")
	registerFunc(&_iOPMAssertionCreateWithName, &_iOPMAssertionCreateWithNameErr, frameworkHandle, "IOPMAssertionCreateWithName", "10.6")
	registerFunc(&_iOPMAssertionCreateWithProperties, &_iOPMAssertionCreateWithPropertiesErr, frameworkHandle, "IOPMAssertionCreateWithProperties", "10.7")
	registerFunc(&_iOPMAssertionDeclareUserActivity, &_iOPMAssertionDeclareUserActivityErr, frameworkHandle, "IOPMAssertionDeclareUserActivity", "10.7")
	registerFunc(&_iOPMAssertionRelease, &_iOPMAssertionReleaseErr, frameworkHandle, "IOPMAssertionRelease", "10.7")
	registerFunc(&_iOPMAssertionRetain, &_iOPMAssertionRetainErr, frameworkHandle, "IOPMAssertionRetain", "10.7")
	registerFunc(&_iOPMAssertionSetProperty, &_iOPMAssertionSetPropertyErr, frameworkHandle, "IOPMAssertionSetProperty", "10.7")
	registerFunc(&_iOPMCancelScheduledPowerEvent, &_iOPMCancelScheduledPowerEventErr, frameworkHandle, "IOPMCancelScheduledPowerEvent", "10.3")
	registerFunc(&_iOPMCopyAssertionsByProcess, &_iOPMCopyAssertionsByProcessErr, frameworkHandle, "IOPMCopyAssertionsByProcess", "10.7")
	registerFunc(&_iOPMCopyAssertionsStatus, &_iOPMCopyAssertionsStatusErr, frameworkHandle, "IOPMCopyAssertionsStatus", "10.7")
	registerFunc(&_iOPMCopyBatteryInfo, &_iOPMCopyBatteryInfoErr, frameworkHandle, "IOPMCopyBatteryInfo", "10.0")
	registerFunc(&_iOPMCopyCPUPowerStatus, &_iOPMCopyCPUPowerStatusErr, frameworkHandle, "IOPMCopyCPUPowerStatus", "10.7")
	registerFunc(&_iOPMCopyScheduledPowerEvents, &_iOPMCopyScheduledPowerEventsErr, frameworkHandle, "IOPMCopyScheduledPowerEvents", "10.3")
	registerFunc(&_iOPMDeclareNetworkClientActivity, &_iOPMDeclareNetworkClientActivityErr, frameworkHandle, "IOPMDeclareNetworkClientActivity", "10.9")
	registerFunc(&_iOPMFindPowerManagement, &_iOPMFindPowerManagementErr, frameworkHandle, "IOPMFindPowerManagement", "10.0")
	registerFunc(&_iOPMGetAggressiveness, &_iOPMGetAggressivenessErr, frameworkHandle, "IOPMGetAggressiveness", "10.0")
	registerFunc(&_iOPMGetThermalWarningLevel, &_iOPMGetThermalWarningLevelErr, frameworkHandle, "IOPMGetThermalWarningLevel", "10.7")
	registerFunc(&_iOPMSchedulePowerEvent, &_iOPMSchedulePowerEventErr, frameworkHandle, "IOPMSchedulePowerEvent", "10.3")
	registerFunc(&_iOPMSetAggressiveness, &_iOPMSetAggressivenessErr, frameworkHandle, "IOPMSetAggressiveness", "10.0")
	registerFunc(&_iOPMSleepEnabled, &_iOPMSleepEnabledErr, frameworkHandle, "IOPMSleepEnabled", "10.0")
	registerFunc(&_iOPMSleepSystem, &_iOPMSleepSystemErr, frameworkHandle, "IOPMSleepSystem", "10.0")
	registerFunc(&_iOPSCopyExternalPowerAdapterDetails, &_iOPSCopyExternalPowerAdapterDetailsErr, frameworkHandle, "IOPSCopyExternalPowerAdapterDetails", "10.6")
	registerFunc(&_iOPSCopyPowerSourcesInfo, &_iOPSCopyPowerSourcesInfoErr, frameworkHandle, "IOPSCopyPowerSourcesInfo", "10.2")
	registerFunc(&_iOPSCopyPowerSourcesList, &_iOPSCopyPowerSourcesListErr, frameworkHandle, "IOPSCopyPowerSourcesList", "10.2")
	registerFunc(&_iOPSCreateLimitedPowerNotification, &_iOPSCreateLimitedPowerNotificationErr, frameworkHandle, "IOPSCreateLimitedPowerNotification", "10.9")
	registerFunc(&_iOPSGetBatteryWarningLevel, &_iOPSGetBatteryWarningLevelErr, frameworkHandle, "IOPSGetBatteryWarningLevel", "10.6")
	registerFunc(&_iOPSGetPowerSourceDescription, &_iOPSGetPowerSourceDescriptionErr, frameworkHandle, "IOPSGetPowerSourceDescription", "10.2")
	registerFunc(&_iOPSGetProvidingPowerSourceType, &_iOPSGetProvidingPowerSourceTypeErr, frameworkHandle, "IOPSGetProvidingPowerSourceType", "10.7")
	registerFunc(&_iOPSGetTimeRemainingEstimate, &_iOPSGetTimeRemainingEstimateErr, frameworkHandle, "IOPSGetTimeRemainingEstimate", "10.7")
	registerFunc(&_iOPSNotificationCreateRunLoopSource, &_iOPSNotificationCreateRunLoopSourceErr, frameworkHandle, "IOPSNotificationCreateRunLoopSource", "10.2")
	registerFunc(&_iORegisterApp, &_iORegisterAppErr, frameworkHandle, "IORegisterApp", "10.0")
	registerFunc(&_iORegisterForSystemPower, &_iORegisterForSystemPowerErr, frameworkHandle, "IORegisterForSystemPower", "10.0")
	registerFunc(&_iORegistryCreateIterator, &_iORegistryCreateIteratorErr, frameworkHandle, "IORegistryCreateIterator", "10.0")
	registerFunc(&_iORegistryEntryCopyFromPath, &_iORegistryEntryCopyFromPathErr, frameworkHandle, "IORegistryEntryCopyFromPath", "10.11")
	registerFunc(&_iORegistryEntryCopyPath, &_iORegistryEntryCopyPathErr, frameworkHandle, "IORegistryEntryCopyPath", "10.11")
	registerFunc(&_iORegistryEntryCreateCFProperties, &_iORegistryEntryCreateCFPropertiesErr, frameworkHandle, "IORegistryEntryCreateCFProperties", "10.0")
	registerFunc(&_iORegistryEntryCreateCFProperty, &_iORegistryEntryCreateCFPropertyErr, frameworkHandle, "IORegistryEntryCreateCFProperty", "10.0")
	registerFunc(&_iORegistryEntryCreateIterator, &_iORegistryEntryCreateIteratorErr, frameworkHandle, "IORegistryEntryCreateIterator", "10.0")
	registerFunc(&_iORegistryEntryFromPath, &_iORegistryEntryFromPathErr, frameworkHandle, "IORegistryEntryFromPath", "10.0")
	registerFunc(&_iORegistryEntryGetChildEntry, &_iORegistryEntryGetChildEntryErr, frameworkHandle, "IORegistryEntryGetChildEntry", "10.0")
	registerFunc(&_iORegistryEntryGetChildIterator, &_iORegistryEntryGetChildIteratorErr, frameworkHandle, "IORegistryEntryGetChildIterator", "10.0")
	registerFunc(&_iORegistryEntryGetLocationInPlane, &_iORegistryEntryGetLocationInPlaneErr, frameworkHandle, "IORegistryEntryGetLocationInPlane", "10.1")
	registerFunc(&_iORegistryEntryGetName, &_iORegistryEntryGetNameErr, frameworkHandle, "IORegistryEntryGetName", "10.0")
	registerFunc(&_iORegistryEntryGetNameInPlane, &_iORegistryEntryGetNameInPlaneErr, frameworkHandle, "IORegistryEntryGetNameInPlane", "10.0")
	registerFunc(&_iORegistryEntryGetParentEntry, &_iORegistryEntryGetParentEntryErr, frameworkHandle, "IORegistryEntryGetParentEntry", "10.0")
	registerFunc(&_iORegistryEntryGetParentIterator, &_iORegistryEntryGetParentIteratorErr, frameworkHandle, "IORegistryEntryGetParentIterator", "10.0")
	registerFunc(&_iORegistryEntryGetPath, &_iORegistryEntryGetPathErr, frameworkHandle, "IORegistryEntryGetPath", "10.0")
	registerFunc(&_iORegistryEntryGetProperty, &_iORegistryEntryGetPropertyErr, frameworkHandle, "IORegistryEntryGetProperty", "10.0")
	registerFunc(&_iORegistryEntryGetRegistryEntryID, &_iORegistryEntryGetRegistryEntryIDErr, frameworkHandle, "IORegistryEntryGetRegistryEntryID", "10.6")
	registerFunc(&_iORegistryEntryIDMatching, &_iORegistryEntryIDMatchingErr, frameworkHandle, "IORegistryEntryIDMatching", "10.6")
	registerFunc(&_iORegistryEntryInPlane, &_iORegistryEntryInPlaneErr, frameworkHandle, "IORegistryEntryInPlane", "10.0")
	registerFunc(&_iORegistryEntrySearchCFProperty, &_iORegistryEntrySearchCFPropertyErr, frameworkHandle, "IORegistryEntrySearchCFProperty", "10.1")
	registerFunc(&_iORegistryEntrySetCFProperties, &_iORegistryEntrySetCFPropertiesErr, frameworkHandle, "IORegistryEntrySetCFProperties", "10.0")
	registerFunc(&_iORegistryEntrySetCFProperty, &_iORegistryEntrySetCFPropertyErr, frameworkHandle, "IORegistryEntrySetCFProperty", "10.0")
	registerFunc(&_iORegistryGetRootEntry, &_iORegistryGetRootEntryErr, frameworkHandle, "IORegistryGetRootEntry", "10.0")
	registerFunc(&_iORegistryIteratorEnterEntry, &_iORegistryIteratorEnterEntryErr, frameworkHandle, "IORegistryIteratorEnterEntry", "10.0")
	registerFunc(&_iORegistryIteratorExitEntry, &_iORegistryIteratorExitEntryErr, frameworkHandle, "IORegistryIteratorExitEntry", "10.0")
	registerFunc(&_iOServiceAddInterestNotification, &_iOServiceAddInterestNotificationErr, frameworkHandle, "IOServiceAddInterestNotification", "10.0")
	registerFunc(&_iOServiceAddMatchingNotification, &_iOServiceAddMatchingNotificationErr, frameworkHandle, "IOServiceAddMatchingNotification", "10.0")
	registerFunc(&_iOServiceAddNotification, &_iOServiceAddNotificationErr, frameworkHandle, "IOServiceAddNotification", "10.0")
	registerFunc(&_iOServiceAuthorize, &_iOServiceAuthorizeErr, frameworkHandle, "IOServiceAuthorize", "10.0")
	registerFunc(&_iOServiceClose, &_iOServiceCloseErr, frameworkHandle, "IOServiceClose", "10.0")
	registerFunc(&_iOServiceGetBusyState, &_iOServiceGetBusyStateErr, frameworkHandle, "IOServiceGetBusyState", "10.0")
	registerFunc(&_iOServiceGetMatchingService, &_iOServiceGetMatchingServiceErr, frameworkHandle, "IOServiceGetMatchingService", "10.2")
	registerFunc(&_iOServiceGetMatchingServices, &_iOServiceGetMatchingServicesErr, frameworkHandle, "IOServiceGetMatchingServices", "10.0")
	registerFunc(&_iOServiceMatchPropertyTable, &_iOServiceMatchPropertyTableErr, frameworkHandle, "IOServiceMatchPropertyTable", "10.0")
	registerFunc(&_iOServiceMatching, &_iOServiceMatchingErr, frameworkHandle, "IOServiceMatching", "10.0")
	registerFunc(&_iOServiceNameMatching, &_iOServiceNameMatchingErr, frameworkHandle, "IOServiceNameMatching", "10.0")
	registerFunc(&_iOServiceOFPathToBSDName, &_iOServiceOFPathToBSDNameErr, frameworkHandle, "IOServiceOFPathToBSDName", "10.0")
	registerFunc(&_iOServiceOpen, &_iOServiceOpenErr, frameworkHandle, "IOServiceOpen", "10.0")
	registerFunc(&_iOServiceOpenAsFileDescriptor, &_iOServiceOpenAsFileDescriptorErr, frameworkHandle, "IOServiceOpenAsFileDescriptor", "10.0")
	registerFunc(&_iOServiceRequestProbe, &_iOServiceRequestProbeErr, frameworkHandle, "IOServiceRequestProbe", "10.0")
	registerFunc(&_iOServiceWaitQuiet, &_iOServiceWaitQuietErr, frameworkHandle, "IOServiceWaitQuiet", "10.0")
	registerFunc(&_iOURLCreateDataAndPropertiesFromResource, &_iOURLCreateDataAndPropertiesFromResourceErr, frameworkHandle, "IOURLCreateDataAndPropertiesFromResource", "10.0")
	registerFunc(&_iOURLCreatePropertyFromResource, &_iOURLCreatePropertyFromResourceErr, frameworkHandle, "IOURLCreatePropertyFromResource", "10.0")
	registerFunc(&_kextManagerCopyLoadedKextInfo, &_kextManagerCopyLoadedKextInfoErr, frameworkHandle, "KextManagerCopyLoadedKextInfo", "10.7")
	registerFunc(&_kextManagerCreateURLForBundleIdentifier, &_kextManagerCreateURLForBundleIdentifierErr, frameworkHandle, "KextManagerCreateURLForBundleIdentifier", "10.2")
	registerFunc(&_kextManagerLoadKextWithIdentifier, &_kextManagerLoadKextWithIdentifierErr, frameworkHandle, "KextManagerLoadKextWithIdentifier", "10.6")
	registerFunc(&_kextManagerLoadKextWithURL, &_kextManagerLoadKextWithURLErr, frameworkHandle, "KextManagerLoadKextWithURL", "10.6")
	registerFunc(&_kextManagerUnloadKextWithIdentifier, &_kextManagerUnloadKextWithIdentifierErr, frameworkHandle, "KextManagerUnloadKextWithIdentifier", "10.7")
	registerFunc(&_nXClickTime, &_nXClickTimeErr, frameworkHandle, "NXClickTime", "10.0")
	registerFunc(&_nXCloseEventStatus, &_nXCloseEventStatusErr, frameworkHandle, "NXCloseEventStatus", "10.0")
	registerFunc(&_nXEventSystemInfo, &_nXEventSystemInfoErr, frameworkHandle, "NXEventSystemInfo", "10.3")
	registerFunc(&_nXGetClickSpace, &_nXGetClickSpaceErr, frameworkHandle, "NXGetClickSpace", "10.0")
	registerFunc(&_nXKeyRepeatInterval, &_nXKeyRepeatIntervalErr, frameworkHandle, "NXKeyRepeatInterval", "10.0")
	registerFunc(&_nXKeyRepeatThreshold, &_nXKeyRepeatThresholdErr, frameworkHandle, "NXKeyRepeatThreshold", "10.0")
	registerFunc(&_nXOpenEventStatus, &_nXOpenEventStatusErr, frameworkHandle, "NXOpenEventStatus", "10.0")
	registerFunc(&_nXResetKeyboard, &_nXResetKeyboardErr, frameworkHandle, "NXResetKeyboard", "10.3")
	registerFunc(&_nXResetMouse, &_nXResetMouseErr, frameworkHandle, "NXResetMouse", "10.3")
	registerFunc(&_nXSetClickSpace, &_nXSetClickSpaceErr, frameworkHandle, "NXSetClickSpace", "10.0")
	registerFunc(&_nXSetClickTime, &_nXSetClickTimeErr, frameworkHandle, "NXSetClickTime", "10.0")
	registerFunc(&_nXSetKeyRepeatInterval, &_nXSetKeyRepeatIntervalErr, frameworkHandle, "NXSetKeyRepeatInterval", "10.0")
	registerFunc(&_nXSetKeyRepeatThreshold, &_nXSetKeyRepeatThresholdErr, frameworkHandle, "NXSetKeyRepeatThreshold", "10.0")
	registerFunc(&_oSGetNotificationFromMessage, &_oSGetNotificationFromMessageErr, frameworkHandle, "OSGetNotificationFromMessage", "10.0")
}

// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"fmt"
	"os"
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: Metal: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: Metal: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: Metal: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _mTLCopyAllDevices func() []unsafe.Pointer

// MTLCopyAllDevices returns an array of all the Metal device instances in the system.
//
// See: https://developer.apple.com/documentation/Metal/MTLCopyAllDevices()
func MTLCopyAllDevices() []MTLDeviceObject {
	if _mTLCopyAllDevices == nil {
		panic("Metal: symbol MTLCopyAllDevices not loaded")
	}
	ptrs := _mTLCopyAllDevices()
	result := make([]MTLDeviceObject, len(ptrs))
	for i, p := range ptrs {
		result[i] = MTLDeviceObjectFromID(objc.IDFrom(p))
	}
	return result
}

var _mTLCopyAllDevicesWithObserver func(observer objectivec.Object, handler MTLDeviceNotificationHandler) []unsafe.Pointer

// MTLCopyAllDevicesWithObserver returns an array of all the Metal GPU devices in the system and registers a notification handler that Metal calls when the device list changes.
//
// See: https://developer.apple.com/documentation/Metal/MTLCopyAllDevicesWithObserver
func MTLCopyAllDevicesWithObserver(observer objectivec.Object, handler MTLDeviceNotificationHandler) []MTLDeviceObject {
	if _mTLCopyAllDevicesWithObserver == nil {
		panic("Metal: symbol MTLCopyAllDevicesWithObserver not loaded")
	}
	ptrs := _mTLCopyAllDevicesWithObserver(observer, handler)
	result := make([]MTLDeviceObject, len(ptrs))
	for i, p := range ptrs {
		result[i] = MTLDeviceObjectFromID(objc.IDFrom(p))
	}
	return result
}

var _mTLCreateSystemDefaultDevice func() unsafe.Pointer

// MTLCreateSystemDefaultDevice returns the device instance Metal selects as the default.
//
// See: https://developer.apple.com/documentation/Metal/MTLCreateSystemDefaultDevice()
func MTLCreateSystemDefaultDevice() MTLDeviceObject {
	if _mTLCreateSystemDefaultDevice == nil {
		panic("Metal: symbol MTLCreateSystemDefaultDevice not loaded")
	}
	rv := _mTLCreateSystemDefaultDevice()
	return MTLDeviceObjectFromID(objc.IDFrom(rv))
}

var _mTLIOCompressionContextAppendData func(context MTLIOCompressionContext, data unsafe.Pointer, size uintptr)

// MTLIOCompressionContextAppendData adds data to a compression context.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCompressionContextAppendData(_:_:_:)
func MTLIOCompressionContextAppendData(context MTLIOCompressionContext, data unsafe.Pointer, size uintptr) {
	if _mTLIOCompressionContextAppendData == nil {
		panic("Metal: symbol MTLIOCompressionContextAppendData not loaded")
	}
	_mTLIOCompressionContextAppendData(context, data, size)
}

var _mTLIOCompressionContextDefaultChunkSize func() uintptr

// MTLIOCompressionContextDefaultChunkSize returns a compression chunk size you can use as a default for creating a compression context.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCompressionContextDefaultChunkSize()
func MTLIOCompressionContextDefaultChunkSize() uintptr {
	if _mTLIOCompressionContextDefaultChunkSize == nil {
		panic("Metal: symbol MTLIOCompressionContextDefaultChunkSize not loaded")
	}
	return _mTLIOCompressionContextDefaultChunkSize()
}

var _mTLIOCreateCompressionContext func(path string, type_ MTLIOCompressionMethod, chunkSize uintptr) MTLIOCompressionContext

// MTLIOCreateCompressionContext creates a compression context that you use to compress data into a single file.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCreateCompressionContext
func MTLIOCreateCompressionContext(path string, type_ MTLIOCompressionMethod, chunkSize uintptr) MTLIOCompressionContext {
	if _mTLIOCreateCompressionContext == nil {
		panic("Metal: symbol MTLIOCreateCompressionContext not loaded")
	}
	return _mTLIOCreateCompressionContext(path, type_, chunkSize)
}

var _mTLIOFlushAndDestroyCompressionContext func(context MTLIOCompressionContext) MTLIOCompressionStatus

// MTLIOFlushAndDestroyCompressionContext finishes compressing and saves the file that a compression context represents.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOFlushAndDestroyCompressionContext(_:)
func MTLIOFlushAndDestroyCompressionContext(context MTLIOCompressionContext) MTLIOCompressionStatus {
	if _mTLIOFlushAndDestroyCompressionContext == nil {
		panic("Metal: symbol MTLIOFlushAndDestroyCompressionContext not loaded")
	}
	return _mTLIOFlushAndDestroyCompressionContext(context)
}

var _mTLRemoveDeviceObserver func(observer objectivec.Object)

// MTLRemoveDeviceObserver removes a registered observer of device notifications.
//
// See: https://developer.apple.com/documentation/Metal/MTLRemoveDeviceObserver(_:)
func MTLRemoveDeviceObserver(observer objectivec.Object) {
	if _mTLRemoveDeviceObserver == nil {
		panic("Metal: symbol MTLRemoveDeviceObserver not loaded")
	}
	_mTLRemoveDeviceObserver(observer)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
		registerFunc(&_mTLCopyAllDevices, frameworkHandle, "MTLCopyAllDevices")
		registerFunc(&_mTLCopyAllDevicesWithObserver, frameworkHandle, "MTLCopyAllDevicesWithObserver")
		registerFunc(&_mTLCreateSystemDefaultDevice, frameworkHandle, "MTLCreateSystemDefaultDevice")
		registerFunc(&_mTLIOCompressionContextAppendData, frameworkHandle, "MTLIOCompressionContextAppendData")
		registerFunc(&_mTLIOCompressionContextDefaultChunkSize, frameworkHandle, "MTLIOCompressionContextDefaultChunkSize")
		registerFunc(&_mTLIOCreateCompressionContext, frameworkHandle, "MTLIOCreateCompressionContext")
		registerFunc(&_mTLIOFlushAndDestroyCompressionContext, frameworkHandle, "MTLIOFlushAndDestroyCompressionContext")
		registerFunc(&_mTLRemoveDeviceObserver, frameworkHandle, "MTLRemoveDeviceObserver")
	}


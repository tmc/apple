// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

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
		return fmt.Sprintf("espresso: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("espresso: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("espresso: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("espresso: register symbol %s: %v", name, r)
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

// SymbolAddress returns the address of name in espresso, whether or not
// this package generated a binding for it.
//
// What is generated is bounded by what is documented, and for a private
// framework that is whatever a manifest happened to enumerate. The dylib
// usually exports far more. Without this, a symbol nobody wrote down is
// unreachable from a package that has already loaded the image holding it, and
// the generated surface becomes a ceiling instead of a floor.
//
// The lookup is scoped to this framework's handle, not RTLD_DEFAULT, so a
// symbol some other loaded image exports is not reported as this one's.
func SymbolAddress(name string) (uintptr, error) {
	if frameworkHandle == 0 {
		return 0, fmt.Errorf("espresso: symbol %s unavailable because the framework could not be loaded", name)
	}
	sym, err := purego.Dlsym(frameworkHandle, name)
	if err != nil || sym == 0 {
		return 0, missingSymbolError(name, "", err)
	}
	return sym, nil
}

// BindFunc binds the espresso symbol name into fptr, which must be a
// pointer to a func variable.
//
// The caller supplies the signature, and nothing checks it. A dylib records no
// argument count or types for a C symbol, so a wrong signature here is not a
// type error: it is the wrong number of machine words moved on a live stack,
// and the failure surfaces somewhere else entirely. Prefer a generated binding,
// whose signature carries recorded evidence, and reach for this only for a
// symbol that has none.
// purego.RegisterFunc panics on a signature it cannot lower; that is recovered
// and returned, because an escape hatch that takes down the process on a
// mistyped experiment is not one anybody can experiment with.
func BindFunc(fptr any, name string) (err error) {
	sym, err := SymbolAddress(name)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("espresso: bind symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	return nil
}

var _e5rtAneMemoryProviderCreate func(a0 uintptr, a1 uintptr) int32
var _e5rtAneMemoryProviderCreateErr error

func tryE5rtAneMemoryProviderCreate(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtAneMemoryProviderCreate == nil {
		return 0, symbolCallError("e5rt_ane_memory_provider_create", "", _e5rtAneMemoryProviderCreateErr)
	}
	return _e5rtAneMemoryProviderCreate(a0, a1), nil
}

// E5rtAneMemoryProviderCreate.
func E5rtAneMemoryProviderCreate(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtAneMemoryProviderCreate(a0, a1)
}

var _e5rtAneMemoryProviderRelease func(out *uintptr) int32
var _e5rtAneMemoryProviderReleaseErr error

func tryE5rtAneMemoryProviderRelease(out *uintptr) (int32, error) {
	if _e5rtAneMemoryProviderRelease == nil {
		return 0, symbolCallError("e5rt_ane_memory_provider_release", "", _e5rtAneMemoryProviderReleaseErr)
	}
	return _e5rtAneMemoryProviderRelease(out), nil
}

// E5rtAneMemoryProviderRelease.
func E5rtAneMemoryProviderRelease(out *uintptr) (int32, error) {
	return tryE5rtAneMemoryProviderRelease(out)
}

var _e5rtAsyncEventAsyncNotify func(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtAsyncEventAsyncNotifyErr error

func tryE5rtAsyncEventAsyncNotify(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtAsyncEventAsyncNotify == nil {
		return 0, symbolCallError("e5rt_async_event_async_notify", "", _e5rtAsyncEventAsyncNotifyErr)
	}
	return _e5rtAsyncEventAsyncNotify(out, a1, a2, a3), nil
}

// E5rtAsyncEventAsyncNotify.
func E5rtAsyncEventAsyncNotify(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtAsyncEventAsyncNotify(out, a1, a2, a3)
}

var _e5rtAsyncEventCreate func(out *uintptr, a1 *byte, a2 uint32) int32
var _e5rtAsyncEventCreateErr error

func tryE5rtAsyncEventCreate(out *uintptr, a1 *byte, a2 uint32) (int32, error) {
	if _e5rtAsyncEventCreate == nil {
		return 0, symbolCallError("e5rt_async_event_create", "", _e5rtAsyncEventCreateErr)
	}
	return _e5rtAsyncEventCreate(out, a1, a2), nil
}

// E5rtAsyncEventCreate.
func E5rtAsyncEventCreate(out *uintptr, a1 *byte, a2 uint32) (int32, error) {
	return tryE5rtAsyncEventCreate(out, a1, a2)
}

var _e5rtAsyncEventCreateFromIosurfaceSharedEvent func(out *uintptr, a1 *byte, a2 uintptr) int32
var _e5rtAsyncEventCreateFromIosurfaceSharedEventErr error

func tryE5rtAsyncEventCreateFromIosurfaceSharedEvent(out *uintptr, a1 *byte, a2 uintptr) (int32, error) {
	if _e5rtAsyncEventCreateFromIosurfaceSharedEvent == nil {
		return 0, symbolCallError("e5rt_async_event_create_from_iosurface_shared_event", "", _e5rtAsyncEventCreateFromIosurfaceSharedEventErr)
	}
	return _e5rtAsyncEventCreateFromIosurfaceSharedEvent(out, a1, a2), nil
}

// E5rtAsyncEventCreateFromIosurfaceSharedEvent.
func E5rtAsyncEventCreateFromIosurfaceSharedEvent(out *uintptr, a1 *byte, a2 uintptr) (int32, error) {
	return tryE5rtAsyncEventCreateFromIosurfaceSharedEvent(out, a1, a2)
}

var _e5rtAsyncEventGetActiveFutureValue func(a0 uintptr, out *uint64) int32
var _e5rtAsyncEventGetActiveFutureValueErr error

func tryE5rtAsyncEventGetActiveFutureValue(a0 uintptr, out *uint64) (int32, error) {
	if _e5rtAsyncEventGetActiveFutureValue == nil {
		return 0, symbolCallError("e5rt_async_event_get_active_future_value", "", _e5rtAsyncEventGetActiveFutureValueErr)
	}
	return _e5rtAsyncEventGetActiveFutureValue(a0, out), nil
}

// E5rtAsyncEventGetActiveFutureValue.
func E5rtAsyncEventGetActiveFutureValue(a0 uintptr, out *uint64) (int32, error) {
	return tryE5rtAsyncEventGetActiveFutureValue(a0, out)
}

var _e5rtAsyncEventGetIosurfaceSharedEvent func(a0 uintptr, out *uintptr) int32
var _e5rtAsyncEventGetIosurfaceSharedEventErr error

func tryE5rtAsyncEventGetIosurfaceSharedEvent(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtAsyncEventGetIosurfaceSharedEvent == nil {
		return 0, symbolCallError("e5rt_async_event_get_iosurface_shared_event", "", _e5rtAsyncEventGetIosurfaceSharedEventErr)
	}
	return _e5rtAsyncEventGetIosurfaceSharedEvent(a0, out), nil
}

// E5rtAsyncEventGetIosurfaceSharedEvent.
func E5rtAsyncEventGetIosurfaceSharedEvent(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtAsyncEventGetIosurfaceSharedEvent(a0, out)
}

var _e5rtAsyncEventGetLastSignaledValue func(a0 uintptr, out *uint64) int32
var _e5rtAsyncEventGetLastSignaledValueErr error

func tryE5rtAsyncEventGetLastSignaledValue(a0 uintptr, out *uint64) (int32, error) {
	if _e5rtAsyncEventGetLastSignaledValue == nil {
		return 0, symbolCallError("e5rt_async_event_get_last_signaled_value", "", _e5rtAsyncEventGetLastSignaledValueErr)
	}
	return _e5rtAsyncEventGetLastSignaledValue(a0, out), nil
}

// E5rtAsyncEventGetLastSignaledValue.
func E5rtAsyncEventGetLastSignaledValue(a0 uintptr, out *uint64) (int32, error) {
	return tryE5rtAsyncEventGetLastSignaledValue(a0, out)
}

var _e5rtAsyncEventGetName func(a0 uintptr, a1 **byte) int32
var _e5rtAsyncEventGetNameErr error

func tryE5rtAsyncEventGetName(a0 uintptr, a1 **byte) (int32, error) {
	if _e5rtAsyncEventGetName == nil {
		return 0, symbolCallError("e5rt_async_event_get_name", "", _e5rtAsyncEventGetNameErr)
	}
	return _e5rtAsyncEventGetName(a0, a1), nil
}

// E5rtAsyncEventGetName.
func E5rtAsyncEventGetName(a0 uintptr, a1 **byte) (int32, error) {
	return tryE5rtAsyncEventGetName(a0, a1)
}

var _e5rtAsyncEventRelease func(out *uintptr) int32
var _e5rtAsyncEventReleaseErr error

func tryE5rtAsyncEventRelease(out *uintptr) (int32, error) {
	if _e5rtAsyncEventRelease == nil {
		return 0, symbolCallError("e5rt_async_event_release", "", _e5rtAsyncEventReleaseErr)
	}
	return _e5rtAsyncEventRelease(out), nil
}

// E5rtAsyncEventRelease.
func E5rtAsyncEventRelease(out *uintptr) (int32, error) {
	return tryE5rtAsyncEventRelease(out)
}

var _e5rtAsyncEventSetActiveFutureValue func(a0 uintptr, a1 uint64) int32
var _e5rtAsyncEventSetActiveFutureValueErr error

func tryE5rtAsyncEventSetActiveFutureValue(a0 uintptr, a1 uint64) (int32, error) {
	if _e5rtAsyncEventSetActiveFutureValue == nil {
		return 0, symbolCallError("e5rt_async_event_set_active_future_value", "", _e5rtAsyncEventSetActiveFutureValueErr)
	}
	return _e5rtAsyncEventSetActiveFutureValue(a0, a1), nil
}

// E5rtAsyncEventSetActiveFutureValue.
func E5rtAsyncEventSetActiveFutureValue(a0 uintptr, a1 uint64) (int32, error) {
	return tryE5rtAsyncEventSetActiveFutureValue(a0, a1)
}

var _e5rtAsyncEventSignal func(a0 uintptr, a1 uint64) int32
var _e5rtAsyncEventSignalErr error

func tryE5rtAsyncEventSignal(a0 uintptr, a1 uint64) (int32, error) {
	if _e5rtAsyncEventSignal == nil {
		return 0, symbolCallError("e5rt_async_event_signal", "", _e5rtAsyncEventSignalErr)
	}
	return _e5rtAsyncEventSignal(a0, a1), nil
}

// E5rtAsyncEventSignal.
func E5rtAsyncEventSignal(a0 uintptr, a1 uint64) (int32, error) {
	return tryE5rtAsyncEventSignal(a0, a1)
}

var _e5rtAsyncEventSyncWait func(a0 uintptr, a1 uint64, a2 uint64) int32
var _e5rtAsyncEventSyncWaitErr error

func tryE5rtAsyncEventSyncWait(a0 uintptr, a1 uint64, a2 uint64) (int32, error) {
	if _e5rtAsyncEventSyncWait == nil {
		return 0, symbolCallError("e5rt_async_event_sync_wait", "", _e5rtAsyncEventSyncWaitErr)
	}
	return _e5rtAsyncEventSyncWait(a0, a1, a2), nil
}

// E5rtAsyncEventSyncWait.
func E5rtAsyncEventSyncWait(a0 uintptr, a1 uint64, a2 uint64) (int32, error) {
	return tryE5rtAsyncEventSyncWait(a0, a1, a2)
}

var _e5rtBufferObjectAlloc func(out *uintptr, a1 uint64, a2 uint32) int32
var _e5rtBufferObjectAllocErr error

func tryE5rtBufferObjectAlloc(out *uintptr, a1 uint64, a2 uint32) (int32, error) {
	if _e5rtBufferObjectAlloc == nil {
		return 0, symbolCallError("e5rt_buffer_object_alloc", "", _e5rtBufferObjectAllocErr)
	}
	return _e5rtBufferObjectAlloc(out, a1, a2), nil
}

// E5rtBufferObjectAlloc.
func E5rtBufferObjectAlloc(out *uintptr, a1 uint64, a2 uint32) (int32, error) {
	return tryE5rtBufferObjectAlloc(out, a1, a2)
}

var _e5rtBufferObjectCreateAsAlias func(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtBufferObjectCreateAsAliasErr error

func tryE5rtBufferObjectCreateAsAlias(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtBufferObjectCreateAsAlias == nil {
		return 0, symbolCallError("e5rt_buffer_object_create_as_alias", "", _e5rtBufferObjectCreateAsAliasErr)
	}
	return _e5rtBufferObjectCreateAsAlias(out, a1, a2, a3), nil
}

// E5rtBufferObjectCreateAsAlias.
func E5rtBufferObjectCreateAsAlias(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtBufferObjectCreateAsAlias(out, a1, a2, a3)
}

var _e5rtBufferObjectCreateFromDataPointer func(out *uintptr, a1 uintptr, a2 uint64) int32
var _e5rtBufferObjectCreateFromDataPointerErr error

func tryE5rtBufferObjectCreateFromDataPointer(out *uintptr, a1 uintptr, a2 uint64) (int32, error) {
	if _e5rtBufferObjectCreateFromDataPointer == nil {
		return 0, symbolCallError("e5rt_buffer_object_create_from_data_pointer", "", _e5rtBufferObjectCreateFromDataPointerErr)
	}
	return _e5rtBufferObjectCreateFromDataPointer(out, a1, a2), nil
}

// E5rtBufferObjectCreateFromDataPointer.
func E5rtBufferObjectCreateFromDataPointer(out *uintptr, a1 uintptr, a2 uint64) (int32, error) {
	return tryE5rtBufferObjectCreateFromDataPointer(out, a1, a2)
}

var _e5rtBufferObjectCreateFromIosurface func(out *uintptr, a1 uintptr) int32
var _e5rtBufferObjectCreateFromIosurfaceErr error

func tryE5rtBufferObjectCreateFromIosurface(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtBufferObjectCreateFromIosurface == nil {
		return 0, symbolCallError("e5rt_buffer_object_create_from_iosurface", "", _e5rtBufferObjectCreateFromIosurfaceErr)
	}
	return _e5rtBufferObjectCreateFromIosurface(out, a1), nil
}

// E5rtBufferObjectCreateFromIosurface.
func E5rtBufferObjectCreateFromIosurface(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtBufferObjectCreateFromIosurface(out, a1)
}

var _e5rtBufferObjectCreateFromMtlbuffer func(out *uintptr, a1 uintptr) int32
var _e5rtBufferObjectCreateFromMtlbufferErr error

func tryE5rtBufferObjectCreateFromMtlbuffer(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtBufferObjectCreateFromMtlbuffer == nil {
		return 0, symbolCallError("e5rt_buffer_object_create_from_mtlbuffer", "", _e5rtBufferObjectCreateFromMtlbufferErr)
	}
	return _e5rtBufferObjectCreateFromMtlbuffer(out, a1), nil
}

// E5rtBufferObjectCreateFromMtlbuffer.
func E5rtBufferObjectCreateFromMtlbuffer(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtBufferObjectCreateFromMtlbuffer(out, a1)
}

var _e5rtBufferObjectGetDataPtr func(a0 uintptr, out *uintptr) int32
var _e5rtBufferObjectGetDataPtrErr error

func tryE5rtBufferObjectGetDataPtr(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtBufferObjectGetDataPtr == nil {
		return 0, symbolCallError("e5rt_buffer_object_get_data_ptr", "", _e5rtBufferObjectGetDataPtrErr)
	}
	return _e5rtBufferObjectGetDataPtr(a0, out), nil
}

// E5rtBufferObjectGetDataPtr.
func E5rtBufferObjectGetDataPtr(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtBufferObjectGetDataPtr(a0, out)
}

var _e5rtBufferObjectGetIosurface func(a0 uintptr, out *uintptr) int32
var _e5rtBufferObjectGetIosurfaceErr error

func tryE5rtBufferObjectGetIosurface(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtBufferObjectGetIosurface == nil {
		return 0, symbolCallError("e5rt_buffer_object_get_iosurface", "", _e5rtBufferObjectGetIosurfaceErr)
	}
	return _e5rtBufferObjectGetIosurface(a0, out), nil
}

// E5rtBufferObjectGetIosurface.
func E5rtBufferObjectGetIosurface(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtBufferObjectGetIosurface(a0, out)
}

var _e5rtBufferObjectGetMtlbuffer func(a0 uintptr, a1 *uintptr) int32
var _e5rtBufferObjectGetMtlbufferErr error

func tryE5rtBufferObjectGetMtlbuffer(a0 uintptr, a1 *uintptr) (int32, error) {
	if _e5rtBufferObjectGetMtlbuffer == nil {
		return 0, symbolCallError("e5rt_buffer_object_get_mtlbuffer", "", _e5rtBufferObjectGetMtlbufferErr)
	}
	return _e5rtBufferObjectGetMtlbuffer(a0, a1), nil
}

// E5rtBufferObjectGetMtlbuffer.
func E5rtBufferObjectGetMtlbuffer(a0 uintptr, a1 *uintptr) (int32, error) {
	return tryE5rtBufferObjectGetMtlbuffer(a0, a1)
}

var _e5rtBufferObjectGetSize func(a0 uintptr, out *uint64) int32
var _e5rtBufferObjectGetSizeErr error

func tryE5rtBufferObjectGetSize(a0 uintptr, out *uint64) (int32, error) {
	if _e5rtBufferObjectGetSize == nil {
		return 0, symbolCallError("e5rt_buffer_object_get_size", "", _e5rtBufferObjectGetSizeErr)
	}
	return _e5rtBufferObjectGetSize(a0, out), nil
}

// E5rtBufferObjectGetSize.
func E5rtBufferObjectGetSize(a0 uintptr, out *uint64) (int32, error) {
	return tryE5rtBufferObjectGetSize(a0, out)
}

var _e5rtBufferObjectGetType func(a0 uintptr, out *uint32) int32
var _e5rtBufferObjectGetTypeErr error

func tryE5rtBufferObjectGetType(a0 uintptr, out *uint32) (int32, error) {
	if _e5rtBufferObjectGetType == nil {
		return 0, symbolCallError("e5rt_buffer_object_get_type", "", _e5rtBufferObjectGetTypeErr)
	}
	return _e5rtBufferObjectGetType(a0, out), nil
}

// E5rtBufferObjectGetType.
func E5rtBufferObjectGetType(a0 uintptr, out *uint32) (int32, error) {
	return tryE5rtBufferObjectGetType(a0, out)
}

var _e5rtBufferObjectRelease func(out *uintptr) int32
var _e5rtBufferObjectReleaseErr error

func tryE5rtBufferObjectRelease(out *uintptr) (int32, error) {
	if _e5rtBufferObjectRelease == nil {
		return 0, symbolCallError("e5rt_buffer_object_release", "", _e5rtBufferObjectReleaseErr)
	}
	return _e5rtBufferObjectRelease(out), nil
}

// E5rtBufferObjectRelease.
func E5rtBufferObjectRelease(out *uintptr) (int32, error) {
	return tryE5rtBufferObjectRelease(out)
}

var _e5rtComputeGPUDeviceGetMtlDevice func(a0 uintptr, out *uintptr) int32
var _e5rtComputeGPUDeviceGetMtlDeviceErr error

func tryE5rtComputeGPUDeviceGetMtlDevice(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtComputeGPUDeviceGetMtlDevice == nil {
		return 0, symbolCallError("e5rt_compute_gpu_device_get_mtl_device", "", _e5rtComputeGPUDeviceGetMtlDeviceErr)
	}
	return _e5rtComputeGPUDeviceGetMtlDevice(a0, out), nil
}

// E5rtComputeGPUDeviceGetMtlDevice.
func E5rtComputeGPUDeviceGetMtlDevice(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtComputeGPUDeviceGetMtlDevice(a0, out)
}

var _e5rtComputeGPUDeviceRelease func(out *uintptr) int32
var _e5rtComputeGPUDeviceReleaseErr error

func tryE5rtComputeGPUDeviceRelease(out *uintptr) (int32, error) {
	if _e5rtComputeGPUDeviceRelease == nil {
		return 0, symbolCallError("e5rt_compute_gpu_device_release", "", _e5rtComputeGPUDeviceReleaseErr)
	}
	return _e5rtComputeGPUDeviceRelease(out), nil
}

// E5rtComputeGPUDeviceRelease.
func E5rtComputeGPUDeviceRelease(out *uintptr) (int32, error) {
	return tryE5rtComputeGPUDeviceRelease(out)
}

var _e5rtComputeGPUDeviceRetainAll func(a0 *uintptr, a1 *uint64) int32
var _e5rtComputeGPUDeviceRetainAllErr error

func tryE5rtComputeGPUDeviceRetainAll(a0 *uintptr, a1 *uint64) (int32, error) {
	if _e5rtComputeGPUDeviceRetainAll == nil {
		return 0, symbolCallError("e5rt_compute_gpu_device_retain_all", "", _e5rtComputeGPUDeviceRetainAllErr)
	}
	return _e5rtComputeGPUDeviceRetainAll(a0, a1), nil
}

// E5rtComputeGPUDeviceRetainAll.
func E5rtComputeGPUDeviceRetainAll(a0 *uintptr, a1 *uint64) (int32, error) {
	return tryE5rtComputeGPUDeviceRetainAll(a0, a1)
}

var _e5rtComputeGPUDeviceRetainFromMtlDevice func(out *uintptr, a1 uintptr) int32
var _e5rtComputeGPUDeviceRetainFromMtlDeviceErr error

func tryE5rtComputeGPUDeviceRetainFromMtlDevice(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtComputeGPUDeviceRetainFromMtlDevice == nil {
		return 0, symbolCallError("e5rt_compute_gpu_device_retain_from_mtl_device", "", _e5rtComputeGPUDeviceRetainFromMtlDeviceErr)
	}
	return _e5rtComputeGPUDeviceRetainFromMtlDevice(out, a1), nil
}

// E5rtComputeGPUDeviceRetainFromMtlDevice.
func E5rtComputeGPUDeviceRetainFromMtlDevice(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtComputeGPUDeviceRetainFromMtlDevice(out, a1)
}

var _e5rtCreateSurfaceObjectFromIosurface func(a0 *uintptr, a1 uintptr) int32
var _e5rtCreateSurfaceObjectFromIosurfaceErr error

func tryE5rtCreateSurfaceObjectFromIosurface(a0 *uintptr, a1 uintptr) (int32, error) {
	if _e5rtCreateSurfaceObjectFromIosurface == nil {
		return 0, symbolCallError("e5rt_create_surface_object_from_iosurface", "", _e5rtCreateSurfaceObjectFromIosurfaceErr)
	}
	return _e5rtCreateSurfaceObjectFromIosurface(a0, a1), nil
}

// E5rtCreateSurfaceObjectFromIosurface.
func E5rtCreateSurfaceObjectFromIosurface(a0 *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtCreateSurfaceObjectFromIosurface(a0, a1)
}

var _e5rtCvpb4ccToSurfaceFormat func(a0 uint32, out *uint32) int32
var _e5rtCvpb4ccToSurfaceFormatErr error

func tryE5rtCvpb4ccToSurfaceFormat(a0 uint32, out *uint32) (int32, error) {
	if _e5rtCvpb4ccToSurfaceFormat == nil {
		return 0, symbolCallError("e5rt_cvpb_4cc_to_surface_format", "", _e5rtCvpb4ccToSurfaceFormatErr)
	}
	return _e5rtCvpb4ccToSurfaceFormat(a0, out), nil
}

// E5rtCvpb4ccToSurfaceFormat.
func E5rtCvpb4ccToSurfaceFormat(a0 uint32, out *uint32) (int32, error) {
	return tryE5rtCvpb4ccToSurfaceFormat(a0, out)
}

var _e5rtE5CompilerCompile func(a0 uintptr, a1 *byte, a2 uintptr, out *uintptr) int32
var _e5rtE5CompilerCompileErr error

func tryE5rtE5CompilerCompile(a0 uintptr, a1 *byte, a2 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerCompile == nil {
		return 0, symbolCallError("e5rt_e5_compiler_compile", "", _e5rtE5CompilerCompileErr)
	}
	return _e5rtE5CompilerCompile(a0, a1, a2, out), nil
}

// E5rtE5CompilerCompile.
func E5rtE5CompilerCompile(a0 uintptr, a1 *byte, a2 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerCompile(a0, a1, a2, out)
}

var _e5rtE5CompilerCompileFromIrProgram func(a0 uintptr, a1 uintptr, a2 uintptr, out *uintptr) int32
var _e5rtE5CompilerCompileFromIrProgramErr error

func tryE5rtE5CompilerCompileFromIrProgram(a0 uintptr, a1 uintptr, a2 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerCompileFromIrProgram == nil {
		return 0, symbolCallError("e5rt_e5_compiler_compile_from_ir_program", "", _e5rtE5CompilerCompileFromIrProgramErr)
	}
	return _e5rtE5CompilerCompileFromIrProgram(a0, a1, a2, out), nil
}

// E5rtE5CompilerCompileFromIrProgram.
func E5rtE5CompilerCompileFromIrProgram(a0 uintptr, a1 uintptr, a2 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerCompileFromIrProgram(a0, a1, a2, out)
}

var _e5rtE5CompilerConfigOptionsCreate func(out *uintptr) int32
var _e5rtE5CompilerConfigOptionsCreateErr error

func tryE5rtE5CompilerConfigOptionsCreate(out *uintptr) (int32, error) {
	if _e5rtE5CompilerConfigOptionsCreate == nil {
		return 0, symbolCallError("e5rt_e5_compiler_config_options_create", "", _e5rtE5CompilerConfigOptionsCreateErr)
	}
	return _e5rtE5CompilerConfigOptionsCreate(out), nil
}

// E5rtE5CompilerConfigOptionsCreate.
func E5rtE5CompilerConfigOptionsCreate(out *uintptr) (int32, error) {
	return tryE5rtE5CompilerConfigOptionsCreate(out)
}

var _e5rtE5CompilerConfigOptionsGetBundleCacheApfsPurgeable func(a0 uintptr, a1 uintptr) int32
var _e5rtE5CompilerConfigOptionsGetBundleCacheApfsPurgeableErr error

func tryE5rtE5CompilerConfigOptionsGetBundleCacheApfsPurgeable(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerConfigOptionsGetBundleCacheApfsPurgeable == nil {
		return 0, symbolCallError("e5rt_e5_compiler_config_options_get_bundle_cache_apfs_purgeable", "", _e5rtE5CompilerConfigOptionsGetBundleCacheApfsPurgeableErr)
	}
	return _e5rtE5CompilerConfigOptionsGetBundleCacheApfsPurgeable(a0, a1), nil
}

// E5rtE5CompilerConfigOptionsGetBundleCacheApfsPurgeable.
func E5rtE5CompilerConfigOptionsGetBundleCacheApfsPurgeable(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerConfigOptionsGetBundleCacheApfsPurgeable(a0, a1)
}

var _e5rtE5CompilerConfigOptionsGetCacheBundleLocation func(a0 uintptr, a1 uintptr) int32
var _e5rtE5CompilerConfigOptionsGetCacheBundleLocationErr error

func tryE5rtE5CompilerConfigOptionsGetCacheBundleLocation(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerConfigOptionsGetCacheBundleLocation == nil {
		return 0, symbolCallError("e5rt_e5_compiler_config_options_get_cache_bundle_location", "", _e5rtE5CompilerConfigOptionsGetCacheBundleLocationErr)
	}
	return _e5rtE5CompilerConfigOptionsGetCacheBundleLocation(a0, a1), nil
}

// E5rtE5CompilerConfigOptionsGetCacheBundleLocation.
func E5rtE5CompilerConfigOptionsGetCacheBundleLocation(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerConfigOptionsGetCacheBundleLocation(a0, a1)
}

var _e5rtE5CompilerConfigOptionsRelease func(out *uintptr) int32
var _e5rtE5CompilerConfigOptionsReleaseErr error

func tryE5rtE5CompilerConfigOptionsRelease(out *uintptr) (int32, error) {
	if _e5rtE5CompilerConfigOptionsRelease == nil {
		return 0, symbolCallError("e5rt_e5_compiler_config_options_release", "", _e5rtE5CompilerConfigOptionsReleaseErr)
	}
	return _e5rtE5CompilerConfigOptionsRelease(out), nil
}

// E5rtE5CompilerConfigOptionsRelease.
func E5rtE5CompilerConfigOptionsRelease(out *uintptr) (int32, error) {
	return tryE5rtE5CompilerConfigOptionsRelease(out)
}

var _e5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeable func(a0 uintptr, a1 bool) int32
var _e5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeableErr error

func tryE5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeable(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeable == nil {
		return 0, symbolCallError("e5rt_e5_compiler_config_options_set_bundle_cache_apfs_purgeable", "", _e5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeableErr)
	}
	return _e5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeable(a0, a1), nil
}

// E5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeable.
func E5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeable(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeable(a0, a1)
}

var _e5rtE5CompilerConfigOptionsSetCacheBundleLocation func(a0 uintptr, a1 *byte) int32
var _e5rtE5CompilerConfigOptionsSetCacheBundleLocationErr error

func tryE5rtE5CompilerConfigOptionsSetCacheBundleLocation(a0 uintptr, a1 *byte) (int32, error) {
	if _e5rtE5CompilerConfigOptionsSetCacheBundleLocation == nil {
		return 0, symbolCallError("e5rt_e5_compiler_config_options_set_cache_bundle_location", "", _e5rtE5CompilerConfigOptionsSetCacheBundleLocationErr)
	}
	return _e5rtE5CompilerConfigOptionsSetCacheBundleLocation(a0, a1), nil
}

// E5rtE5CompilerConfigOptionsSetCacheBundleLocation.
func E5rtE5CompilerConfigOptionsSetCacheBundleLocation(a0 uintptr, a1 *byte) (int32, error) {
	return tryE5rtE5CompilerConfigOptionsSetCacheBundleLocation(a0, a1)
}

var _e5rtE5CompilerCreate func(out *uintptr) int32
var _e5rtE5CompilerCreateErr error

func tryE5rtE5CompilerCreate(out *uintptr) (int32, error) {
	if _e5rtE5CompilerCreate == nil {
		return 0, symbolCallError("e5rt_e5_compiler_create", "", _e5rtE5CompilerCreateErr)
	}
	return _e5rtE5CompilerCreate(out), nil
}

// E5rtE5CompilerCreate.
func E5rtE5CompilerCreate(out *uintptr) (int32, error) {
	return tryE5rtE5CompilerCreate(out)
}

var _e5rtE5CompilerCreateWithConfig func(out *uintptr, a1 uintptr) int32
var _e5rtE5CompilerCreateWithConfigErr error

func tryE5rtE5CompilerCreateWithConfig(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerCreateWithConfig == nil {
		return 0, symbolCallError("e5rt_e5_compiler_create_with_config", "", _e5rtE5CompilerCreateWithConfigErr)
	}
	return _e5rtE5CompilerCreateWithConfig(out, a1), nil
}

// E5rtE5CompilerCreateWithConfig.
func E5rtE5CompilerCreateWithConfig(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerCreateWithConfig(out, a1)
}

var _e5rtE5CompilerIsNewCompileRequired func(a0 uintptr, a1 *byte, a2 uintptr, out *uintptr) int32
var _e5rtE5CompilerIsNewCompileRequiredErr error

func tryE5rtE5CompilerIsNewCompileRequired(a0 uintptr, a1 *byte, a2 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerIsNewCompileRequired == nil {
		return 0, symbolCallError("e5rt_e5_compiler_is_new_compile_required", "", _e5rtE5CompilerIsNewCompileRequiredErr)
	}
	return _e5rtE5CompilerIsNewCompileRequired(a0, a1, a2, out), nil
}

// E5rtE5CompilerIsNewCompileRequired.
func E5rtE5CompilerIsNewCompileRequired(a0 uintptr, a1 *byte, a2 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerIsNewCompileRequired(a0, a1, a2, out)
}

var _e5rtE5CompilerOptionsCreate func(out *uintptr) int32
var _e5rtE5CompilerOptionsCreateErr error

func tryE5rtE5CompilerOptionsCreate(out *uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsCreate == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_create", "", _e5rtE5CompilerOptionsCreateErr)
	}
	return _e5rtE5CompilerOptionsCreate(out), nil
}

// E5rtE5CompilerOptionsCreate.
func E5rtE5CompilerOptionsCreate(out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsCreate(out)
}

var _e5rtE5CompilerOptionsGetComputeDeviceTypesMask func(a0 uintptr, out *uint64) int32
var _e5rtE5CompilerOptionsGetComputeDeviceTypesMaskErr error

func tryE5rtE5CompilerOptionsGetComputeDeviceTypesMask(a0 uintptr, out *uint64) (int32, error) {
	if _e5rtE5CompilerOptionsGetComputeDeviceTypesMask == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_compute_device_types_mask", "", _e5rtE5CompilerOptionsGetComputeDeviceTypesMaskErr)
	}
	return _e5rtE5CompilerOptionsGetComputeDeviceTypesMask(a0, out), nil
}

// E5rtE5CompilerOptionsGetComputeDeviceTypesMask.
func E5rtE5CompilerOptionsGetComputeDeviceTypesMask(a0 uintptr, out *uint64) (int32, error) {
	return tryE5rtE5CompilerOptionsGetComputeDeviceTypesMask(a0, out)
}

var _e5rtE5CompilerOptionsGetCreateProtectedAssets func(a0 uintptr, out *uintptr) int32
var _e5rtE5CompilerOptionsGetCreateProtectedAssetsErr error

func tryE5rtE5CompilerOptionsGetCreateProtectedAssets(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsGetCreateProtectedAssets == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_create_protected_assets", "", _e5rtE5CompilerOptionsGetCreateProtectedAssetsErr)
	}
	return _e5rtE5CompilerOptionsGetCreateProtectedAssets(a0, out), nil
}

// E5rtE5CompilerOptionsGetCreateProtectedAssets.
func E5rtE5CompilerOptionsGetCreateProtectedAssets(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsGetCreateProtectedAssets(a0, out)
}

var _e5rtE5CompilerOptionsGetCustomAneCompilerOptions func(a0 uintptr, out *uintptr) int32
var _e5rtE5CompilerOptionsGetCustomAneCompilerOptionsErr error

func tryE5rtE5CompilerOptionsGetCustomAneCompilerOptions(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsGetCustomAneCompilerOptions == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_custom_ane_compiler_options", "", _e5rtE5CompilerOptionsGetCustomAneCompilerOptionsErr)
	}
	return _e5rtE5CompilerOptionsGetCustomAneCompilerOptions(a0, out), nil
}

// E5rtE5CompilerOptionsGetCustomAneCompilerOptions.
func E5rtE5CompilerOptionsGetCustomAneCompilerOptions(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsGetCustomAneCompilerOptions(a0, out)
}

var _e5rtE5CompilerOptionsGetEnableMpsgraphPackage func(a0 uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsGetEnableMpsgraphPackageErr error

func tryE5rtE5CompilerOptionsGetEnableMpsgraphPackage(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsGetEnableMpsgraphPackage == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_enable_mpsgraph_package", "", _e5rtE5CompilerOptionsGetEnableMpsgraphPackageErr)
	}
	return _e5rtE5CompilerOptionsGetEnableMpsgraphPackage(a0, a1), nil
}

// E5rtE5CompilerOptionsGetEnableMpsgraphPackage.
func E5rtE5CompilerOptionsGetEnableMpsgraphPackage(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsGetEnableMpsgraphPackage(a0, a1)
}

var _e5rtE5CompilerOptionsGetEnableProfiling func(a0 uintptr, out *uintptr) int32
var _e5rtE5CompilerOptionsGetEnableProfilingErr error

func tryE5rtE5CompilerOptionsGetEnableProfiling(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsGetEnableProfiling == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_enable_profiling", "", _e5rtE5CompilerOptionsGetEnableProfilingErr)
	}
	return _e5rtE5CompilerOptionsGetEnableProfiling(a0, out), nil
}

// E5rtE5CompilerOptionsGetEnableProfiling.
func E5rtE5CompilerOptionsGetEnableProfiling(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsGetEnableProfiling(a0, out)
}

var _e5rtE5CompilerOptionsGetEnableReshapeWithMinimalAllocations func(a0 uintptr, out *uintptr) int32
var _e5rtE5CompilerOptionsGetEnableReshapeWithMinimalAllocationsErr error

func tryE5rtE5CompilerOptionsGetEnableReshapeWithMinimalAllocations(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsGetEnableReshapeWithMinimalAllocations == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_enable_reshape_with_minimal_allocations", "", _e5rtE5CompilerOptionsGetEnableReshapeWithMinimalAllocationsErr)
	}
	return _e5rtE5CompilerOptionsGetEnableReshapeWithMinimalAllocations(a0, out), nil
}

// E5rtE5CompilerOptionsGetEnableReshapeWithMinimalAllocations.
func E5rtE5CompilerOptionsGetEnableReshapeWithMinimalAllocations(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsGetEnableReshapeWithMinimalAllocations(a0, out)
}

var _e5rtE5CompilerOptionsGetExperimentalDisableCompileTimeMpsgraphTypeInference func(a0 uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsGetExperimentalDisableCompileTimeMpsgraphTypeInferenceErr error

func tryE5rtE5CompilerOptionsGetExperimentalDisableCompileTimeMpsgraphTypeInference(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsGetExperimentalDisableCompileTimeMpsgraphTypeInference == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_experimental_disable_compile_time_mpsgraph_type_inference", "", _e5rtE5CompilerOptionsGetExperimentalDisableCompileTimeMpsgraphTypeInferenceErr)
	}
	return _e5rtE5CompilerOptionsGetExperimentalDisableCompileTimeMpsgraphTypeInference(a0, a1), nil
}

// E5rtE5CompilerOptionsGetExperimentalDisableCompileTimeMpsgraphTypeInference.
func E5rtE5CompilerOptionsGetExperimentalDisableCompileTimeMpsgraphTypeInference(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsGetExperimentalDisableCompileTimeMpsgraphTypeInference(a0, a1)
}

var _e5rtE5CompilerOptionsGetExperimentalDisableDataDependentShape func(a0 uintptr, out *uintptr) int32
var _e5rtE5CompilerOptionsGetExperimentalDisableDataDependentShapeErr error

func tryE5rtE5CompilerOptionsGetExperimentalDisableDataDependentShape(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsGetExperimentalDisableDataDependentShape == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_experimental_disable_data_dependent_shape", "", _e5rtE5CompilerOptionsGetExperimentalDisableDataDependentShapeErr)
	}
	return _e5rtE5CompilerOptionsGetExperimentalDisableDataDependentShape(a0, out), nil
}

// E5rtE5CompilerOptionsGetExperimentalDisableDataDependentShape.
func E5rtE5CompilerOptionsGetExperimentalDisableDataDependentShape(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsGetExperimentalDisableDataDependentShape(a0, out)
}

var _e5rtE5CompilerOptionsGetExperimentalEnableDefaultFunctionForRangeDim func(a0 uintptr, out *uintptr) int32
var _e5rtE5CompilerOptionsGetExperimentalEnableDefaultFunctionForRangeDimErr error

func tryE5rtE5CompilerOptionsGetExperimentalEnableDefaultFunctionForRangeDim(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsGetExperimentalEnableDefaultFunctionForRangeDim == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_experimental_enable_default_function_for_range_dim", "", _e5rtE5CompilerOptionsGetExperimentalEnableDefaultFunctionForRangeDimErr)
	}
	return _e5rtE5CompilerOptionsGetExperimentalEnableDefaultFunctionForRangeDim(a0, out), nil
}

// E5rtE5CompilerOptionsGetExperimentalEnableDefaultFunctionForRangeDim.
func E5rtE5CompilerOptionsGetExperimentalEnableDefaultFunctionForRangeDim(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsGetExperimentalEnableDefaultFunctionForRangeDim(a0, out)
}

var _e5rtE5CompilerOptionsGetExperimentalForceClassicCPUBackend func(a0 uintptr, out *uintptr) int32
var _e5rtE5CompilerOptionsGetExperimentalForceClassicCPUBackendErr error

func tryE5rtE5CompilerOptionsGetExperimentalForceClassicCPUBackend(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsGetExperimentalForceClassicCPUBackend == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_experimental_force_classic_cpu_backend", "", _e5rtE5CompilerOptionsGetExperimentalForceClassicCPUBackendErr)
	}
	return _e5rtE5CompilerOptionsGetExperimentalForceClassicCPUBackend(a0, out), nil
}

// E5rtE5CompilerOptionsGetExperimentalForceClassicCPUBackend.
func E5rtE5CompilerOptionsGetExperimentalForceClassicCPUBackend(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsGetExperimentalForceClassicCPUBackend(a0, out)
}

var _e5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatterns func(a0 uintptr, out *uintptr) int32
var _e5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatternsErr error

func tryE5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatterns(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatterns == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_experimental_match_e5_minimal_cpu_patterns", "", _e5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatternsErr)
	}
	return _e5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatterns(a0, out), nil
}

// E5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatterns.
func E5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatterns(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatterns(a0, out)
}

var _e5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatternsForStates func(a0 uintptr, out *uintptr) int32
var _e5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatternsForStatesErr error

func tryE5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatternsForStates(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatternsForStates == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_experimental_match_e5_minimal_cpu_patterns_for_states", "", _e5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatternsForStatesErr)
	}
	return _e5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatternsForStates(a0, out), nil
}

// E5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatternsForStates.
func E5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatternsForStates(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatternsForStates(a0, out)
}

var _e5rtE5CompilerOptionsGetForceBnnsGraph func(a0 uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsGetForceBnnsGraphErr error

func tryE5rtE5CompilerOptionsGetForceBnnsGraph(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsGetForceBnnsGraph == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_force_bnns_graph", "", _e5rtE5CompilerOptionsGetForceBnnsGraphErr)
	}
	return _e5rtE5CompilerOptionsGetForceBnnsGraph(a0, a1), nil
}

// E5rtE5CompilerOptionsGetForceBnnsGraph.
func E5rtE5CompilerOptionsGetForceBnnsGraph(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsGetForceBnnsGraph(a0, a1)
}

var _e5rtE5CompilerOptionsGetForceClassicAotOldHw func(a0 uintptr, out *uintptr) int32
var _e5rtE5CompilerOptionsGetForceClassicAotOldHwErr error

func tryE5rtE5CompilerOptionsGetForceClassicAotOldHw(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsGetForceClassicAotOldHw == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_force_classic_aot_old_hw", "", _e5rtE5CompilerOptionsGetForceClassicAotOldHwErr)
	}
	return _e5rtE5CompilerOptionsGetForceClassicAotOldHw(a0, out), nil
}

// E5rtE5CompilerOptionsGetForceClassicAotOldHw.
func E5rtE5CompilerOptionsGetForceClassicAotOldHw(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsGetForceClassicAotOldHw(a0, out)
}

var _e5rtE5CompilerOptionsGetForceFetchFromCache func(a0 uintptr, out *uintptr) int32
var _e5rtE5CompilerOptionsGetForceFetchFromCacheErr error

func tryE5rtE5CompilerOptionsGetForceFetchFromCache(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsGetForceFetchFromCache == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_force_fetch_from_cache", "", _e5rtE5CompilerOptionsGetForceFetchFromCacheErr)
	}
	return _e5rtE5CompilerOptionsGetForceFetchFromCache(a0, out), nil
}

// E5rtE5CompilerOptionsGetForceFetchFromCache.
func E5rtE5CompilerOptionsGetForceFetchFromCache(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsGetForceFetchFromCache(a0, out)
}

var _e5rtE5CompilerOptionsGetForceRecompilation func(a0 uintptr, out *uintptr) int32
var _e5rtE5CompilerOptionsGetForceRecompilationErr error

func tryE5rtE5CompilerOptionsGetForceRecompilation(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsGetForceRecompilation == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_force_recompilation", "", _e5rtE5CompilerOptionsGetForceRecompilationErr)
	}
	return _e5rtE5CompilerOptionsGetForceRecompilation(a0, out), nil
}

// E5rtE5CompilerOptionsGetForceRecompilation.
func E5rtE5CompilerOptionsGetForceRecompilation(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsGetForceRecompilation(a0, out)
}

var _e5rtE5CompilerOptionsGetPreferredCPUBackend func(a0 uintptr, out *uintptr) int32
var _e5rtE5CompilerOptionsGetPreferredCPUBackendErr error

func tryE5rtE5CompilerOptionsGetPreferredCPUBackend(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsGetPreferredCPUBackend == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_preferred_cpu_backend", "", _e5rtE5CompilerOptionsGetPreferredCPUBackendErr)
	}
	return _e5rtE5CompilerOptionsGetPreferredCPUBackend(a0, out), nil
}

// E5rtE5CompilerOptionsGetPreferredCPUBackend.
func E5rtE5CompilerOptionsGetPreferredCPUBackend(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsGetPreferredCPUBackend(a0, out)
}

var _e5rtE5CompilerOptionsGetPreferredCPUBackends func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtE5CompilerOptionsGetPreferredCPUBackendsErr error

func tryE5rtE5CompilerOptionsGetPreferredCPUBackends(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsGetPreferredCPUBackends == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_preferred_cpu_backends", "", _e5rtE5CompilerOptionsGetPreferredCPUBackendsErr)
	}
	return _e5rtE5CompilerOptionsGetPreferredCPUBackends(a0, a1, a2), nil
}

// E5rtE5CompilerOptionsGetPreferredCPUBackends.
func E5rtE5CompilerOptionsGetPreferredCPUBackends(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsGetPreferredCPUBackends(a0, a1, a2)
}

var _e5rtE5CompilerOptionsGetSegmenter func(a0 uintptr, out *uintptr) int32
var _e5rtE5CompilerOptionsGetSegmenterErr error

func tryE5rtE5CompilerOptionsGetSegmenter(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsGetSegmenter == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_segmenter", "", _e5rtE5CompilerOptionsGetSegmenterErr)
	}
	return _e5rtE5CompilerOptionsGetSegmenter(a0, out), nil
}

// E5rtE5CompilerOptionsGetSegmenter.
func E5rtE5CompilerOptionsGetSegmenter(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsGetSegmenter(a0, out)
}

var _e5rtE5CompilerOptionsRelease func(out *uintptr) int32
var _e5rtE5CompilerOptionsReleaseErr error

func tryE5rtE5CompilerOptionsRelease(out *uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsRelease == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_release", "", _e5rtE5CompilerOptionsReleaseErr)
	}
	return _e5rtE5CompilerOptionsRelease(out), nil
}

// E5rtE5CompilerOptionsRelease.
func E5rtE5CompilerOptionsRelease(out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsRelease(out)
}

var _e5rtE5CompilerOptionsRetainMilEntryPoints func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtE5CompilerOptionsRetainMilEntryPointsErr error

func tryE5rtE5CompilerOptionsRetainMilEntryPoints(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsRetainMilEntryPoints == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_retain_mil_entry_points", "", _e5rtE5CompilerOptionsRetainMilEntryPointsErr)
	}
	return _e5rtE5CompilerOptionsRetainMilEntryPoints(a0, a1, a2), nil
}

// E5rtE5CompilerOptionsRetainMilEntryPoints.
func E5rtE5CompilerOptionsRetainMilEntryPoints(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsRetainMilEntryPoints(a0, a1, a2)
}

var _e5rtE5CompilerOptionsSetComputeDeviceTypesMask func(a0 uintptr, a1 uint64) int32
var _e5rtE5CompilerOptionsSetComputeDeviceTypesMaskErr error

func tryE5rtE5CompilerOptionsSetComputeDeviceTypesMask(a0 uintptr, a1 uint64) (int32, error) {
	if _e5rtE5CompilerOptionsSetComputeDeviceTypesMask == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_compute_device_types_mask", "", _e5rtE5CompilerOptionsSetComputeDeviceTypesMaskErr)
	}
	return _e5rtE5CompilerOptionsSetComputeDeviceTypesMask(a0, a1), nil
}

// E5rtE5CompilerOptionsSetComputeDeviceTypesMask.
func E5rtE5CompilerOptionsSetComputeDeviceTypesMask(a0 uintptr, a1 uint64) (int32, error) {
	return tryE5rtE5CompilerOptionsSetComputeDeviceTypesMask(a0, a1)
}

var _e5rtE5CompilerOptionsSetCreateProtectedAssets func(a0 uintptr, a1 bool) int32
var _e5rtE5CompilerOptionsSetCreateProtectedAssetsErr error

func tryE5rtE5CompilerOptionsSetCreateProtectedAssets(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtE5CompilerOptionsSetCreateProtectedAssets == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_create_protected_assets", "", _e5rtE5CompilerOptionsSetCreateProtectedAssetsErr)
	}
	return _e5rtE5CompilerOptionsSetCreateProtectedAssets(a0, a1), nil
}

// E5rtE5CompilerOptionsSetCreateProtectedAssets.
func E5rtE5CompilerOptionsSetCreateProtectedAssets(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtE5CompilerOptionsSetCreateProtectedAssets(a0, a1)
}

var _e5rtE5CompilerOptionsSetCustomAneCompilerOptions func(a0 uintptr, a1 *byte) int32
var _e5rtE5CompilerOptionsSetCustomAneCompilerOptionsErr error

func tryE5rtE5CompilerOptionsSetCustomAneCompilerOptions(a0 uintptr, a1 *byte) (int32, error) {
	if _e5rtE5CompilerOptionsSetCustomAneCompilerOptions == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_custom_ane_compiler_options", "", _e5rtE5CompilerOptionsSetCustomAneCompilerOptionsErr)
	}
	return _e5rtE5CompilerOptionsSetCustomAneCompilerOptions(a0, a1), nil
}

// E5rtE5CompilerOptionsSetCustomAneCompilerOptions.
func E5rtE5CompilerOptionsSetCustomAneCompilerOptions(a0 uintptr, a1 *byte) (int32, error) {
	return tryE5rtE5CompilerOptionsSetCustomAneCompilerOptions(a0, a1)
}

var _e5rtE5CompilerOptionsSetEnableMpsgraphPackage func(a0 uintptr, a1 bool) int32
var _e5rtE5CompilerOptionsSetEnableMpsgraphPackageErr error

func tryE5rtE5CompilerOptionsSetEnableMpsgraphPackage(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtE5CompilerOptionsSetEnableMpsgraphPackage == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_enable_mpsgraph_package", "", _e5rtE5CompilerOptionsSetEnableMpsgraphPackageErr)
	}
	return _e5rtE5CompilerOptionsSetEnableMpsgraphPackage(a0, a1), nil
}

// E5rtE5CompilerOptionsSetEnableMpsgraphPackage.
func E5rtE5CompilerOptionsSetEnableMpsgraphPackage(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtE5CompilerOptionsSetEnableMpsgraphPackage(a0, a1)
}

var _e5rtE5CompilerOptionsSetEnableProfiling func(a0 uintptr, a1 bool) int32
var _e5rtE5CompilerOptionsSetEnableProfilingErr error

func tryE5rtE5CompilerOptionsSetEnableProfiling(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtE5CompilerOptionsSetEnableProfiling == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_enable_profiling", "", _e5rtE5CompilerOptionsSetEnableProfilingErr)
	}
	return _e5rtE5CompilerOptionsSetEnableProfiling(a0, a1), nil
}

// E5rtE5CompilerOptionsSetEnableProfiling.
func E5rtE5CompilerOptionsSetEnableProfiling(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtE5CompilerOptionsSetEnableProfiling(a0, a1)
}

var _e5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocations func(a0 uintptr, a1 bool) int32
var _e5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocationsErr error

func tryE5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocations(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocations == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_enable_reshape_with_minimal_allocations", "", _e5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocationsErr)
	}
	return _e5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocations(a0, a1), nil
}

// E5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocations.
func E5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocations(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocations(a0, a1)
}

var _e5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference func(a0 uintptr, a1 bool) int32
var _e5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInferenceErr error

func tryE5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_experimental_disable_compile_time_mpsgraph_type_inference", "", _e5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInferenceErr)
	}
	return _e5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference(a0, a1), nil
}

// E5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference.
func E5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference(a0, a1)
}

var _e5rtE5CompilerOptionsSetExperimentalDisableDataDependentShape func(a0 uintptr, a1 bool) int32
var _e5rtE5CompilerOptionsSetExperimentalDisableDataDependentShapeErr error

func tryE5rtE5CompilerOptionsSetExperimentalDisableDataDependentShape(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtE5CompilerOptionsSetExperimentalDisableDataDependentShape == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_experimental_disable_data_dependent_shape", "", _e5rtE5CompilerOptionsSetExperimentalDisableDataDependentShapeErr)
	}
	return _e5rtE5CompilerOptionsSetExperimentalDisableDataDependentShape(a0, a1), nil
}

// E5rtE5CompilerOptionsSetExperimentalDisableDataDependentShape.
func E5rtE5CompilerOptionsSetExperimentalDisableDataDependentShape(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtE5CompilerOptionsSetExperimentalDisableDataDependentShape(a0, a1)
}

var _e5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDim func(a0 uintptr, a1 bool) int32
var _e5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDimErr error

func tryE5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDim(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDim == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_experimental_enable_default_function_for_range_dim", "", _e5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDimErr)
	}
	return _e5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDim(a0, a1), nil
}

// E5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDim.
func E5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDim(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDim(a0, a1)
}

var _e5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackend func(a0 uintptr, a1 bool) int32
var _e5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackendErr error

func tryE5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackend(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackend == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_experimental_force_classic_cpu_backend", "", _e5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackendErr)
	}
	return _e5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackend(a0, a1), nil
}

// E5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackend.
func E5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackend(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackend(a0, a1)
}

var _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatterns func(a0 uintptr, a1 bool) int32
var _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsErr error

func tryE5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatterns(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatterns == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_experimental_match_e5_minimal_cpu_patterns", "", _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsErr)
	}
	return _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatterns(a0, a1), nil
}

// E5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatterns.
func E5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatterns(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatterns(a0, a1)
}

var _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStates func(a0 uintptr, a1 bool) int32
var _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStatesErr error

func tryE5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStates(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStates == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_experimental_match_e5_minimal_cpu_patterns_for_states", "", _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStatesErr)
	}
	return _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStates(a0, a1), nil
}

// E5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStates.
func E5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStates(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStates(a0, a1)
}

var _e5rtE5CompilerOptionsSetForceBnnsGraph func(a0 uintptr, a1 bool) int32
var _e5rtE5CompilerOptionsSetForceBnnsGraphErr error

func tryE5rtE5CompilerOptionsSetForceBnnsGraph(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtE5CompilerOptionsSetForceBnnsGraph == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_force_bnns_graph", "", _e5rtE5CompilerOptionsSetForceBnnsGraphErr)
	}
	return _e5rtE5CompilerOptionsSetForceBnnsGraph(a0, a1), nil
}

// E5rtE5CompilerOptionsSetForceBnnsGraph.
func E5rtE5CompilerOptionsSetForceBnnsGraph(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtE5CompilerOptionsSetForceBnnsGraph(a0, a1)
}

var _e5rtE5CompilerOptionsSetForceClassicAotOldHw func(a0 uintptr, a1 bool) int32
var _e5rtE5CompilerOptionsSetForceClassicAotOldHwErr error

func tryE5rtE5CompilerOptionsSetForceClassicAotOldHw(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtE5CompilerOptionsSetForceClassicAotOldHw == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_force_classic_aot_old_hw", "", _e5rtE5CompilerOptionsSetForceClassicAotOldHwErr)
	}
	return _e5rtE5CompilerOptionsSetForceClassicAotOldHw(a0, a1), nil
}

// E5rtE5CompilerOptionsSetForceClassicAotOldHw.
func E5rtE5CompilerOptionsSetForceClassicAotOldHw(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtE5CompilerOptionsSetForceClassicAotOldHw(a0, a1)
}

var _e5rtE5CompilerOptionsSetForceFetchFromCache func(a0 uintptr, a1 bool) int32
var _e5rtE5CompilerOptionsSetForceFetchFromCacheErr error

func tryE5rtE5CompilerOptionsSetForceFetchFromCache(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtE5CompilerOptionsSetForceFetchFromCache == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_force_fetch_from_cache", "", _e5rtE5CompilerOptionsSetForceFetchFromCacheErr)
	}
	return _e5rtE5CompilerOptionsSetForceFetchFromCache(a0, a1), nil
}

// E5rtE5CompilerOptionsSetForceFetchFromCache.
func E5rtE5CompilerOptionsSetForceFetchFromCache(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtE5CompilerOptionsSetForceFetchFromCache(a0, a1)
}

var _e5rtE5CompilerOptionsSetForceRecompilation func(a0 uintptr, a1 bool) int32
var _e5rtE5CompilerOptionsSetForceRecompilationErr error

func tryE5rtE5CompilerOptionsSetForceRecompilation(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtE5CompilerOptionsSetForceRecompilation == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_force_recompilation", "", _e5rtE5CompilerOptionsSetForceRecompilationErr)
	}
	return _e5rtE5CompilerOptionsSetForceRecompilation(a0, a1), nil
}

// E5rtE5CompilerOptionsSetForceRecompilation.
func E5rtE5CompilerOptionsSetForceRecompilation(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtE5CompilerOptionsSetForceRecompilation(a0, a1)
}

var _e5rtE5CompilerOptionsSetMilEntryPoints func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtE5CompilerOptionsSetMilEntryPointsErr error

func tryE5rtE5CompilerOptionsSetMilEntryPoints(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetMilEntryPoints == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_mil_entry_points", "", _e5rtE5CompilerOptionsSetMilEntryPointsErr)
	}
	return _e5rtE5CompilerOptionsSetMilEntryPoints(a0, a1, a2), nil
}

// E5rtE5CompilerOptionsSetMilEntryPoints.
func E5rtE5CompilerOptionsSetMilEntryPoints(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetMilEntryPoints(a0, a1, a2)
}

var _e5rtE5CompilerOptionsSetPreferredCPUBackend func(a0 uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsSetPreferredCPUBackendErr error

func tryE5rtE5CompilerOptionsSetPreferredCPUBackend(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetPreferredCPUBackend == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_preferred_cpu_backend", "", _e5rtE5CompilerOptionsSetPreferredCPUBackendErr)
	}
	return _e5rtE5CompilerOptionsSetPreferredCPUBackend(a0, a1), nil
}

// E5rtE5CompilerOptionsSetPreferredCPUBackend.
func E5rtE5CompilerOptionsSetPreferredCPUBackend(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetPreferredCPUBackend(a0, a1)
}

var _e5rtE5CompilerOptionsSetPreferredCPUBackends func(a0 uintptr, out *uintptr, a2 uintptr) int32
var _e5rtE5CompilerOptionsSetPreferredCPUBackendsErr error

func tryE5rtE5CompilerOptionsSetPreferredCPUBackends(a0 uintptr, out *uintptr, a2 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetPreferredCPUBackends == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_preferred_cpu_backends", "", _e5rtE5CompilerOptionsSetPreferredCPUBackendsErr)
	}
	return _e5rtE5CompilerOptionsSetPreferredCPUBackends(a0, out, a2), nil
}

// E5rtE5CompilerOptionsSetPreferredCPUBackends.
func E5rtE5CompilerOptionsSetPreferredCPUBackends(a0 uintptr, out *uintptr, a2 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetPreferredCPUBackends(a0, out, a2)
}

var _e5rtE5CompilerOptionsSetSegmenter func(a0 uintptr, a1 *byte) int32
var _e5rtE5CompilerOptionsSetSegmenterErr error

func tryE5rtE5CompilerOptionsSetSegmenter(a0 uintptr, a1 *byte) (int32, error) {
	if _e5rtE5CompilerOptionsSetSegmenter == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_segmenter", "", _e5rtE5CompilerOptionsSetSegmenterErr)
	}
	return _e5rtE5CompilerOptionsSetSegmenter(a0, a1), nil
}

// E5rtE5CompilerOptionsSetSegmenter.
func E5rtE5CompilerOptionsSetSegmenter(a0 uintptr, a1 *byte) (int32, error) {
	return tryE5rtE5CompilerOptionsSetSegmenter(a0, a1)
}

var _e5rtE5CompilerPurgeE5BundlesForInputModel func(out *uintptr, a1 *byte) int32
var _e5rtE5CompilerPurgeE5BundlesForInputModelErr error

func tryE5rtE5CompilerPurgeE5BundlesForInputModel(out *uintptr, a1 *byte) (int32, error) {
	if _e5rtE5CompilerPurgeE5BundlesForInputModel == nil {
		return 0, symbolCallError("e5rt_e5_compiler_purge_e5_bundles_for_input_model", "", _e5rtE5CompilerPurgeE5BundlesForInputModelErr)
	}
	return _e5rtE5CompilerPurgeE5BundlesForInputModel(out, a1), nil
}

// E5rtE5CompilerPurgeE5BundlesForInputModel.
func E5rtE5CompilerPurgeE5BundlesForInputModel(out *uintptr, a1 *byte) (int32, error) {
	return tryE5rtE5CompilerPurgeE5BundlesForInputModel(out, a1)
}

var _e5rtE5CompilerRelease func(out *uintptr) int32
var _e5rtE5CompilerReleaseErr error

func tryE5rtE5CompilerRelease(out *uintptr) (int32, error) {
	if _e5rtE5CompilerRelease == nil {
		return 0, symbolCallError("e5rt_e5_compiler_release", "", _e5rtE5CompilerReleaseErr)
	}
	return _e5rtE5CompilerRelease(out), nil
}

// E5rtE5CompilerRelease.
func E5rtE5CompilerRelease(out *uintptr) (int32, error) {
	return tryE5rtE5CompilerRelease(out)
}

var _e5rtErrorCodeGetString func(a0 int32) *byte
var _e5rtErrorCodeGetStringErr error

func tryE5rtErrorCodeGetString(a0 int32) (*byte, error) {
	if _e5rtErrorCodeGetString == nil {
		return nil, symbolCallError("e5rt_error_code_get_string", "", _e5rtErrorCodeGetStringErr)
	}
	return _e5rtErrorCodeGetString(a0), nil
}

// E5rtErrorCodeGetString.
func E5rtErrorCodeGetString(a0 int32) (*byte, error) {
	return tryE5rtErrorCodeGetString(a0)
}

var _e5rtExecutionStreamAsyncSubmit func(a0 uintptr) int32
var _e5rtExecutionStreamAsyncSubmitErr error

func tryE5rtExecutionStreamAsyncSubmit(a0 uintptr) (int32, error) {
	if _e5rtExecutionStreamAsyncSubmit == nil {
		return 0, symbolCallError("e5rt_execution_stream_async_submit", "", _e5rtExecutionStreamAsyncSubmitErr)
	}
	return _e5rtExecutionStreamAsyncSubmit(a0), nil
}

// E5rtExecutionStreamAsyncSubmit.
func E5rtExecutionStreamAsyncSubmit(a0 uintptr) (int32, error) {
	return tryE5rtExecutionStreamAsyncSubmit(a0)
}

var _e5rtExecutionStreamConfigOptionsCreate func(a0 *uintptr) int32
var _e5rtExecutionStreamConfigOptionsCreateErr error

func tryE5rtExecutionStreamConfigOptionsCreate(a0 *uintptr) (int32, error) {
	if _e5rtExecutionStreamConfigOptionsCreate == nil {
		return 0, symbolCallError("e5rt_execution_stream_config_options_create", "", _e5rtExecutionStreamConfigOptionsCreateErr)
	}
	return _e5rtExecutionStreamConfigOptionsCreate(a0), nil
}

// E5rtExecutionStreamConfigOptionsCreate.
func E5rtExecutionStreamConfigOptionsCreate(a0 *uintptr) (int32, error) {
	return tryE5rtExecutionStreamConfigOptionsCreate(a0)
}

var _e5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecution func(a0 uintptr, out *bool) int32
var _e5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecutionErr error

func tryE5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecution(a0 uintptr, out *bool) (int32, error) {
	if _e5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecution == nil {
		return 0, symbolCallError("e5rt_execution_stream_config_options_get_enable_concurrent_sync_execution", "", _e5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecutionErr)
	}
	return _e5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecution(a0, out), nil
}

// E5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecution.
func E5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecution(a0 uintptr, out *bool) (int32, error) {
	return tryE5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecution(a0, out)
}

var _e5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEvents func(a0 uintptr, out *bool) int32
var _e5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEventsErr error

func tryE5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEvents(a0 uintptr, out *bool) (int32, error) {
	if _e5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEvents == nil {
		return 0, symbolCallError("e5rt_execution_stream_config_options_get_enable_low_latency_async_events", "", _e5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEventsErr)
	}
	return _e5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEvents(a0, out), nil
}

// E5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEvents.
func E5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEvents(a0 uintptr, out *bool) (int32, error) {
	return tryE5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEvents(a0, out)
}

var _e5rtExecutionStreamConfigOptionsGetSkipIOFences func(a0 uintptr, out *bool) int32
var _e5rtExecutionStreamConfigOptionsGetSkipIOFencesErr error

func tryE5rtExecutionStreamConfigOptionsGetSkipIOFences(a0 uintptr, out *bool) (int32, error) {
	if _e5rtExecutionStreamConfigOptionsGetSkipIOFences == nil {
		return 0, symbolCallError("e5rt_execution_stream_config_options_get_skip_io_fences", "", _e5rtExecutionStreamConfigOptionsGetSkipIOFencesErr)
	}
	return _e5rtExecutionStreamConfigOptionsGetSkipIOFences(a0, out), nil
}

// E5rtExecutionStreamConfigOptionsGetSkipIOFences.
func E5rtExecutionStreamConfigOptionsGetSkipIOFences(a0 uintptr, out *bool) (int32, error) {
	return tryE5rtExecutionStreamConfigOptionsGetSkipIOFences(a0, out)
}

var _e5rtExecutionStreamConfigOptionsRelease func(out *uintptr) int32
var _e5rtExecutionStreamConfigOptionsReleaseErr error

func tryE5rtExecutionStreamConfigOptionsRelease(out *uintptr) (int32, error) {
	if _e5rtExecutionStreamConfigOptionsRelease == nil {
		return 0, symbolCallError("e5rt_execution_stream_config_options_release", "", _e5rtExecutionStreamConfigOptionsReleaseErr)
	}
	return _e5rtExecutionStreamConfigOptionsRelease(out), nil
}

// E5rtExecutionStreamConfigOptionsRelease.
func E5rtExecutionStreamConfigOptionsRelease(out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamConfigOptionsRelease(out)
}

var _e5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecution func(a0 uintptr, a1 bool) int32
var _e5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecutionErr error

func tryE5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecution(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecution == nil {
		return 0, symbolCallError("e5rt_execution_stream_config_options_set_enable_concurrent_sync_execution", "", _e5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecutionErr)
	}
	return _e5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecution(a0, a1), nil
}

// E5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecution.
func E5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecution(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecution(a0, a1)
}

var _e5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEvents func(a0 uintptr, a1 bool) int32
var _e5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEventsErr error

func tryE5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEvents(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEvents == nil {
		return 0, symbolCallError("e5rt_execution_stream_config_options_set_enable_low_latency_async_events", "", _e5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEventsErr)
	}
	return _e5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEvents(a0, a1), nil
}

// E5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEvents.
func E5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEvents(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEvents(a0, a1)
}

var _e5rtExecutionStreamConfigOptionsSetSkipIOFences func(a0 uintptr, a1 bool) int32
var _e5rtExecutionStreamConfigOptionsSetSkipIOFencesErr error

func tryE5rtExecutionStreamConfigOptionsSetSkipIOFences(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtExecutionStreamConfigOptionsSetSkipIOFences == nil {
		return 0, symbolCallError("e5rt_execution_stream_config_options_set_skip_io_fences", "", _e5rtExecutionStreamConfigOptionsSetSkipIOFencesErr)
	}
	return _e5rtExecutionStreamConfigOptionsSetSkipIOFences(a0, a1), nil
}

// E5rtExecutionStreamConfigOptionsSetSkipIOFences.
func E5rtExecutionStreamConfigOptionsSetSkipIOFences(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtExecutionStreamConfigOptionsSetSkipIOFences(a0, a1)
}

var _e5rtExecutionStreamCreate func(out *uintptr) int32
var _e5rtExecutionStreamCreateErr error

func tryE5rtExecutionStreamCreate(out *uintptr) (int32, error) {
	if _e5rtExecutionStreamCreate == nil {
		return 0, symbolCallError("e5rt_execution_stream_create", "", _e5rtExecutionStreamCreateErr)
	}
	return _e5rtExecutionStreamCreate(out), nil
}

// E5rtExecutionStreamCreate.
func E5rtExecutionStreamCreate(out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamCreate(out)
}

var _e5rtExecutionStreamEncodeOperation func(a0 uintptr, a1 uintptr) int32
var _e5rtExecutionStreamEncodeOperationErr error

func tryE5rtExecutionStreamEncodeOperation(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamEncodeOperation == nil {
		return 0, symbolCallError("e5rt_execution_stream_encode_operation", "", _e5rtExecutionStreamEncodeOperationErr)
	}
	return _e5rtExecutionStreamEncodeOperation(a0, a1), nil
}

// E5rtExecutionStreamEncodeOperation.
func E5rtExecutionStreamEncodeOperation(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamEncodeOperation(a0, a1)
}

var _e5rtExecutionStreamEncodeWorkload func(a0 uintptr, a1 uintptr) int32
var _e5rtExecutionStreamEncodeWorkloadErr error

func tryE5rtExecutionStreamEncodeWorkload(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamEncodeWorkload == nil {
		return 0, symbolCallError("e5rt_execution_stream_encode_workload", "", _e5rtExecutionStreamEncodeWorkloadErr)
	}
	return _e5rtExecutionStreamEncodeWorkload(a0, a1), nil
}

// E5rtExecutionStreamEncodeWorkload.
func E5rtExecutionStreamEncodeWorkload(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamEncodeWorkload(a0, a1)
}

var _e5rtExecutionStreamExecuteSync func(a0 uintptr) int32
var _e5rtExecutionStreamExecuteSyncErr error

func tryE5rtExecutionStreamExecuteSync(a0 uintptr) (int32, error) {
	if _e5rtExecutionStreamExecuteSync == nil {
		return 0, symbolCallError("e5rt_execution_stream_execute_sync", "", _e5rtExecutionStreamExecuteSyncErr)
	}
	return _e5rtExecutionStreamExecuteSync(a0), nil
}

// E5rtExecutionStreamExecuteSync.
func E5rtExecutionStreamExecuteSync(a0 uintptr) (int32, error) {
	return tryE5rtExecutionStreamExecuteSync(a0)
}

var _e5rtExecutionStreamGetInternalAsyncComputeRequestIDForLastSubmit func(a0 uintptr, out *uintptr) int32
var _e5rtExecutionStreamGetInternalAsyncComputeRequestIDForLastSubmitErr error

func tryE5rtExecutionStreamGetInternalAsyncComputeRequestIDForLastSubmit(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtExecutionStreamGetInternalAsyncComputeRequestIDForLastSubmit == nil {
		return 0, symbolCallError("e5rt_execution_stream_get_internal_async_compute_request_id_for_last_submit", "", _e5rtExecutionStreamGetInternalAsyncComputeRequestIDForLastSubmitErr)
	}
	return _e5rtExecutionStreamGetInternalAsyncComputeRequestIDForLastSubmit(a0, out), nil
}

// E5rtExecutionStreamGetInternalAsyncComputeRequestIDForLastSubmit.
func E5rtExecutionStreamGetInternalAsyncComputeRequestIDForLastSubmit(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamGetInternalAsyncComputeRequestIDForLastSubmit(a0, out)
}

var _e5rtExecutionStreamGetStreamID func(a0 uintptr, out *uintptr) int32
var _e5rtExecutionStreamGetStreamIDErr error

func tryE5rtExecutionStreamGetStreamID(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtExecutionStreamGetStreamID == nil {
		return 0, symbolCallError("e5rt_execution_stream_get_stream_id", "", _e5rtExecutionStreamGetStreamIDErr)
	}
	return _e5rtExecutionStreamGetStreamID(a0, out), nil
}

// E5rtExecutionStreamGetStreamID.
func E5rtExecutionStreamGetStreamID(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamGetStreamID(a0, out)
}

var _e5rtExecutionStreamOperationBindCompletionEvent func(a0 uintptr, a1 uintptr) int32
var _e5rtExecutionStreamOperationBindCompletionEventErr error

func tryE5rtExecutionStreamOperationBindCompletionEvent(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationBindCompletionEvent == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_bind_completion_event", "", _e5rtExecutionStreamOperationBindCompletionEventErr)
	}
	return _e5rtExecutionStreamOperationBindCompletionEvent(a0, a1), nil
}

// E5rtExecutionStreamOperationBindCompletionEvent.
func E5rtExecutionStreamOperationBindCompletionEvent(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationBindCompletionEvent(a0, a1)
}

var _e5rtExecutionStreamOperationBindDependentEvents func(a0 uintptr, a1 *uintptr, a2 uint64) int32
var _e5rtExecutionStreamOperationBindDependentEventsErr error

func tryE5rtExecutionStreamOperationBindDependentEvents(a0 uintptr, a1 *uintptr, a2 uint64) (int32, error) {
	if _e5rtExecutionStreamOperationBindDependentEvents == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_bind_dependent_events", "", _e5rtExecutionStreamOperationBindDependentEventsErr)
	}
	return _e5rtExecutionStreamOperationBindDependentEvents(a0, a1, a2), nil
}

// E5rtExecutionStreamOperationBindDependentEvents.
func E5rtExecutionStreamOperationBindDependentEvents(a0 uintptr, a1 *uintptr, a2 uint64) (int32, error) {
	return tryE5rtExecutionStreamOperationBindDependentEvents(a0, a1, a2)
}

var _e5rtExecutionStreamOperationConfigOptionsCreate func(a0 uintptr) int32
var _e5rtExecutionStreamOperationConfigOptionsCreateErr error

func tryE5rtExecutionStreamOperationConfigOptionsCreate(a0 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationConfigOptionsCreate == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_config_options_create", "", _e5rtExecutionStreamOperationConfigOptionsCreateErr)
	}
	return _e5rtExecutionStreamOperationConfigOptionsCreate(a0), nil
}

// E5rtExecutionStreamOperationConfigOptionsCreate.
func E5rtExecutionStreamOperationConfigOptionsCreate(a0 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationConfigOptionsCreate(a0)
}

var _e5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemory func(a0 uintptr, out *uintptr) int32
var _e5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemoryErr error

func tryE5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemory(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemory == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_config_options_get_prewire_model_memory", "", _e5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemoryErr)
	}
	return _e5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemory(a0, out), nil
}

// E5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemory.
func E5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemory(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemory(a0, out)
}

var _e5rtExecutionStreamOperationConfigOptionsRelease func(out *uintptr) int32
var _e5rtExecutionStreamOperationConfigOptionsReleaseErr error

func tryE5rtExecutionStreamOperationConfigOptionsRelease(out *uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationConfigOptionsRelease == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_config_options_release", "", _e5rtExecutionStreamOperationConfigOptionsReleaseErr)
	}
	return _e5rtExecutionStreamOperationConfigOptionsRelease(out), nil
}

// E5rtExecutionStreamOperationConfigOptionsRelease.
func E5rtExecutionStreamOperationConfigOptionsRelease(out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationConfigOptionsRelease(out)
}

var _e5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemory func(a0 uintptr, a1 bool) int32
var _e5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemoryErr error

func tryE5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemory(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemory == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_config_options_set_prewire_model_memory", "", _e5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemoryErr)
	}
	return _e5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemory(a0, a1), nil
}

// E5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemory.
func E5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemory(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemory(a0, a1)
}

var _e5rtExecutionStreamOperationCreatePrecompiledComputeOperation func(a0 uintptr, a1 *byte, a2 *byte, a3 *byte, a4 uintptr, a5 bool) int32
var _e5rtExecutionStreamOperationCreatePrecompiledComputeOperationErr error

func tryE5rtExecutionStreamOperationCreatePrecompiledComputeOperation(a0 uintptr, a1 *byte, a2 *byte, a3 *byte, a4 uintptr, a5 bool) (int32, error) {
	if _e5rtExecutionStreamOperationCreatePrecompiledComputeOperation == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_create_precompiled_compute_operation", "", _e5rtExecutionStreamOperationCreatePrecompiledComputeOperationErr)
	}
	return _e5rtExecutionStreamOperationCreatePrecompiledComputeOperation(a0, a1, a2, a3, a4, a5), nil
}

// E5rtExecutionStreamOperationCreatePrecompiledComputeOperation.
func E5rtExecutionStreamOperationCreatePrecompiledComputeOperation(a0 uintptr, a1 *byte, a2 *byte, a3 *byte, a4 uintptr, a5 bool) (int32, error) {
	return tryE5rtExecutionStreamOperationCreatePrecompiledComputeOperation(a0, a1, a2, a3, a4, a5)
}

var _e5rtExecutionStreamOperationCreatePrecompiledComputeOperationWithOptions func(out *uintptr, a1 uintptr) int32
var _e5rtExecutionStreamOperationCreatePrecompiledComputeOperationWithOptionsErr error

func tryE5rtExecutionStreamOperationCreatePrecompiledComputeOperationWithOptions(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationCreatePrecompiledComputeOperationWithOptions == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_create_precompiled_compute_operation_with_options", "", _e5rtExecutionStreamOperationCreatePrecompiledComputeOperationWithOptionsErr)
	}
	return _e5rtExecutionStreamOperationCreatePrecompiledComputeOperationWithOptions(out, a1), nil
}

// E5rtExecutionStreamOperationCreatePrecompiledComputeOperationWithOptions.
func E5rtExecutionStreamOperationCreatePrecompiledComputeOperationWithOptions(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationCreatePrecompiledComputeOperationWithOptions(out, a1)
}

var _e5rtExecutionStreamOperationCreateResourceSharingPrecompiledComputeOperationsWithMultipleOptions func(out *uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtExecutionStreamOperationCreateResourceSharingPrecompiledComputeOperationsWithMultipleOptionsErr error

func tryE5rtExecutionStreamOperationCreateResourceSharingPrecompiledComputeOperationsWithMultipleOptions(out *uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationCreateResourceSharingPrecompiledComputeOperationsWithMultipleOptions == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_create_resource_sharing_precompiled_compute_operations_with_multiple_options", "", _e5rtExecutionStreamOperationCreateResourceSharingPrecompiledComputeOperationsWithMultipleOptionsErr)
	}
	return _e5rtExecutionStreamOperationCreateResourceSharingPrecompiledComputeOperationsWithMultipleOptions(out, a1, a2), nil
}

// E5rtExecutionStreamOperationCreateResourceSharingPrecompiledComputeOperationsWithMultipleOptions.
func E5rtExecutionStreamOperationCreateResourceSharingPrecompiledComputeOperationsWithMultipleOptions(out *uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationCreateResourceSharingPrecompiledComputeOperationsWithMultipleOptions(out, a1, a2)
}

var _e5rtExecutionStreamOperationGetDependentEventCount func(a0 uintptr, out *uintptr) int32
var _e5rtExecutionStreamOperationGetDependentEventCountErr error

func tryE5rtExecutionStreamOperationGetDependentEventCount(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationGetDependentEventCount == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_get_dependent_event_count", "", _e5rtExecutionStreamOperationGetDependentEventCountErr)
	}
	return _e5rtExecutionStreamOperationGetDependentEventCount(a0, out), nil
}

// E5rtExecutionStreamOperationGetDependentEventCount.
func E5rtExecutionStreamOperationGetDependentEventCount(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationGetDependentEventCount(a0, out)
}

var _e5rtExecutionStreamOperationGetInoutNames func(a0 uintptr, a1 uint64, a2 **byte) int32
var _e5rtExecutionStreamOperationGetInoutNamesErr error

func tryE5rtExecutionStreamOperationGetInoutNames(a0 uintptr, a1 uint64, a2 **byte) (int32, error) {
	if _e5rtExecutionStreamOperationGetInoutNames == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_get_inout_names", "", _e5rtExecutionStreamOperationGetInoutNamesErr)
	}
	return _e5rtExecutionStreamOperationGetInoutNames(a0, a1, a2), nil
}

// E5rtExecutionStreamOperationGetInoutNames.
func E5rtExecutionStreamOperationGetInoutNames(a0 uintptr, a1 uint64, a2 **byte) (int32, error) {
	return tryE5rtExecutionStreamOperationGetInoutNames(a0, a1, a2)
}

var _e5rtExecutionStreamOperationGetInputNames func(a0 uintptr, a1 uint64, a2 **byte) int32
var _e5rtExecutionStreamOperationGetInputNamesErr error

func tryE5rtExecutionStreamOperationGetInputNames(a0 uintptr, a1 uint64, a2 **byte) (int32, error) {
	if _e5rtExecutionStreamOperationGetInputNames == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_get_input_names", "", _e5rtExecutionStreamOperationGetInputNamesErr)
	}
	return _e5rtExecutionStreamOperationGetInputNames(a0, a1, a2), nil
}

// E5rtExecutionStreamOperationGetInputNames.
func E5rtExecutionStreamOperationGetInputNames(a0 uintptr, a1 uint64, a2 **byte) (int32, error) {
	return tryE5rtExecutionStreamOperationGetInputNames(a0, a1, a2)
}

var _e5rtExecutionStreamOperationGetNumInouts func(a0 uintptr, a1 *uint64) int32
var _e5rtExecutionStreamOperationGetNumInoutsErr error

func tryE5rtExecutionStreamOperationGetNumInouts(a0 uintptr, a1 *uint64) (int32, error) {
	if _e5rtExecutionStreamOperationGetNumInouts == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_get_num_inouts", "", _e5rtExecutionStreamOperationGetNumInoutsErr)
	}
	return _e5rtExecutionStreamOperationGetNumInouts(a0, a1), nil
}

// E5rtExecutionStreamOperationGetNumInouts.
func E5rtExecutionStreamOperationGetNumInouts(a0 uintptr, a1 *uint64) (int32, error) {
	return tryE5rtExecutionStreamOperationGetNumInouts(a0, a1)
}

var _e5rtExecutionStreamOperationGetNumInputs func(a0 uintptr, out *uint64) int32
var _e5rtExecutionStreamOperationGetNumInputsErr error

func tryE5rtExecutionStreamOperationGetNumInputs(a0 uintptr, out *uint64) (int32, error) {
	if _e5rtExecutionStreamOperationGetNumInputs == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_get_num_inputs", "", _e5rtExecutionStreamOperationGetNumInputsErr)
	}
	return _e5rtExecutionStreamOperationGetNumInputs(a0, out), nil
}

// E5rtExecutionStreamOperationGetNumInputs.
func E5rtExecutionStreamOperationGetNumInputs(a0 uintptr, out *uint64) (int32, error) {
	return tryE5rtExecutionStreamOperationGetNumInputs(a0, out)
}

var _e5rtExecutionStreamOperationGetNumOutputs func(a0 uintptr, out *uint64) int32
var _e5rtExecutionStreamOperationGetNumOutputsErr error

func tryE5rtExecutionStreamOperationGetNumOutputs(a0 uintptr, out *uint64) (int32, error) {
	if _e5rtExecutionStreamOperationGetNumOutputs == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_get_num_outputs", "", _e5rtExecutionStreamOperationGetNumOutputsErr)
	}
	return _e5rtExecutionStreamOperationGetNumOutputs(a0, out), nil
}

// E5rtExecutionStreamOperationGetNumOutputs.
func E5rtExecutionStreamOperationGetNumOutputs(a0 uintptr, out *uint64) (int32, error) {
	return tryE5rtExecutionStreamOperationGetNumOutputs(a0, out)
}

var _e5rtExecutionStreamOperationGetOpname func(a0 uintptr, a1 **byte) int32
var _e5rtExecutionStreamOperationGetOpnameErr error

func tryE5rtExecutionStreamOperationGetOpname(a0 uintptr, a1 **byte) (int32, error) {
	if _e5rtExecutionStreamOperationGetOpname == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_get_opname", "", _e5rtExecutionStreamOperationGetOpnameErr)
	}
	return _e5rtExecutionStreamOperationGetOpname(a0, a1), nil
}

// E5rtExecutionStreamOperationGetOpname.
func E5rtExecutionStreamOperationGetOpname(a0 uintptr, a1 **byte) (int32, error) {
	return tryE5rtExecutionStreamOperationGetOpname(a0, a1)
}

var _e5rtExecutionStreamOperationGetOutputNames func(a0 uintptr, a1 uint64, a2 **byte) int32
var _e5rtExecutionStreamOperationGetOutputNamesErr error

func tryE5rtExecutionStreamOperationGetOutputNames(a0 uintptr, a1 uint64, a2 **byte) (int32, error) {
	if _e5rtExecutionStreamOperationGetOutputNames == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_get_output_names", "", _e5rtExecutionStreamOperationGetOutputNamesErr)
	}
	return _e5rtExecutionStreamOperationGetOutputNames(a0, a1, a2), nil
}

// E5rtExecutionStreamOperationGetOutputNames.
func E5rtExecutionStreamOperationGetOutputNames(a0 uintptr, a1 uint64, a2 **byte) (int32, error) {
	return tryE5rtExecutionStreamOperationGetOutputNames(a0, a1, a2)
}

var _e5rtExecutionStreamOperationPrepareOpForEncode func(a0 uintptr) int32
var _e5rtExecutionStreamOperationPrepareOpForEncodeErr error

func tryE5rtExecutionStreamOperationPrepareOpForEncode(a0 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationPrepareOpForEncode == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_prepare_op_for_encode", "", _e5rtExecutionStreamOperationPrepareOpForEncodeErr)
	}
	return _e5rtExecutionStreamOperationPrepareOpForEncode(a0), nil
}

// E5rtExecutionStreamOperationPrepareOpForEncode.
func E5rtExecutionStreamOperationPrepareOpForEncode(a0 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationPrepareOpForEncode(a0)
}

var _e5rtExecutionStreamOperationRelease func(out *uintptr) int32
var _e5rtExecutionStreamOperationReleaseErr error

func tryE5rtExecutionStreamOperationRelease(out *uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationRelease == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_release", "", _e5rtExecutionStreamOperationReleaseErr)
	}
	return _e5rtExecutionStreamOperationRelease(out), nil
}

// E5rtExecutionStreamOperationRelease.
func E5rtExecutionStreamOperationRelease(out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationRelease(out)
}

var _e5rtExecutionStreamOperationReshapeOperation func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtExecutionStreamOperationReshapeOperationErr error

func tryE5rtExecutionStreamOperationReshapeOperation(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationReshapeOperation == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_reshape_operation", "", _e5rtExecutionStreamOperationReshapeOperationErr)
	}
	return _e5rtExecutionStreamOperationReshapeOperation(a0, a1, a2, a3), nil
}

// E5rtExecutionStreamOperationReshapeOperation.
func E5rtExecutionStreamOperationReshapeOperation(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationReshapeOperation(a0, a1, a2, a3)
}

var _e5rtExecutionStreamOperationRetainCompletionEvent func(a0 uintptr, a1 uintptr) int32
var _e5rtExecutionStreamOperationRetainCompletionEventErr error

func tryE5rtExecutionStreamOperationRetainCompletionEvent(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationRetainCompletionEvent == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_retain_completion_event", "", _e5rtExecutionStreamOperationRetainCompletionEventErr)
	}
	return _e5rtExecutionStreamOperationRetainCompletionEvent(a0, a1), nil
}

// E5rtExecutionStreamOperationRetainCompletionEvent.
func E5rtExecutionStreamOperationRetainCompletionEvent(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationRetainCompletionEvent(a0, a1)
}

var _e5rtExecutionStreamOperationRetainDependentEvents func(a0 uintptr, a1 uintptr) int32
var _e5rtExecutionStreamOperationRetainDependentEventsErr error

func tryE5rtExecutionStreamOperationRetainDependentEvents(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationRetainDependentEvents == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_retain_dependent_events", "", _e5rtExecutionStreamOperationRetainDependentEventsErr)
	}
	return _e5rtExecutionStreamOperationRetainDependentEvents(a0, a1), nil
}

// E5rtExecutionStreamOperationRetainDependentEvents.
func E5rtExecutionStreamOperationRetainDependentEvents(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationRetainDependentEvents(a0, a1)
}

var _e5rtExecutionStreamOperationRetainInoutPort func(a0 uintptr, a1 *byte, out *uintptr) int32
var _e5rtExecutionStreamOperationRetainInoutPortErr error

func tryE5rtExecutionStreamOperationRetainInoutPort(a0 uintptr, a1 *byte, out *uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationRetainInoutPort == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_retain_inout_port", "", _e5rtExecutionStreamOperationRetainInoutPortErr)
	}
	return _e5rtExecutionStreamOperationRetainInoutPort(a0, a1, out), nil
}

// E5rtExecutionStreamOperationRetainInoutPort.
func E5rtExecutionStreamOperationRetainInoutPort(a0 uintptr, a1 *byte, out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationRetainInoutPort(a0, a1, out)
}

var _e5rtExecutionStreamOperationRetainInputPort func(a0 uintptr, a1 *byte, out *uintptr) int32
var _e5rtExecutionStreamOperationRetainInputPortErr error

func tryE5rtExecutionStreamOperationRetainInputPort(a0 uintptr, a1 *byte, out *uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationRetainInputPort == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_retain_input_port", "", _e5rtExecutionStreamOperationRetainInputPortErr)
	}
	return _e5rtExecutionStreamOperationRetainInputPort(a0, a1, out), nil
}

// E5rtExecutionStreamOperationRetainInputPort.
func E5rtExecutionStreamOperationRetainInputPort(a0 uintptr, a1 *byte, out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationRetainInputPort(a0, a1, out)
}

var _e5rtExecutionStreamOperationRetainOutputPort func(a0 uintptr, a1 *byte, out *uintptr) int32
var _e5rtExecutionStreamOperationRetainOutputPortErr error

func tryE5rtExecutionStreamOperationRetainOutputPort(a0 uintptr, a1 *byte, out *uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationRetainOutputPort == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_retain_output_port", "", _e5rtExecutionStreamOperationRetainOutputPortErr)
	}
	return _e5rtExecutionStreamOperationRetainOutputPort(a0, a1, out), nil
}

// E5rtExecutionStreamOperationRetainOutputPort.
func E5rtExecutionStreamOperationRetainOutputPort(a0 uintptr, a1 *byte, out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationRetainOutputPort(a0, a1, out)
}

var _e5rtExecutionStreamOperationSerializeInferenceFrameData func(a0 uintptr, a1 *byte, a2 *byte, a3 bool) int32
var _e5rtExecutionStreamOperationSerializeInferenceFrameDataErr error

func tryE5rtExecutionStreamOperationSerializeInferenceFrameData(a0 uintptr, a1 *byte, a2 *byte, a3 bool) (int32, error) {
	if _e5rtExecutionStreamOperationSerializeInferenceFrameData == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_serialize_inference_frame_data", "", _e5rtExecutionStreamOperationSerializeInferenceFrameDataErr)
	}
	return _e5rtExecutionStreamOperationSerializeInferenceFrameData(a0, a1, a2, a3), nil
}

// E5rtExecutionStreamOperationSerializeInferenceFrameData.
func E5rtExecutionStreamOperationSerializeInferenceFrameData(a0 uintptr, a1 *byte, a2 *byte, a3 bool) (int32, error) {
	return tryE5rtExecutionStreamOperationSerializeInferenceFrameData(a0, a1, a2, a3)
}

var _e5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegment func(a0 uintptr, a1 *byte, a2 *byte, a3 bool) int32
var _e5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegmentErr error

func tryE5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegment(a0 uintptr, a1 *byte, a2 *byte, a3 bool) (int32, error) {
	if _e5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegment == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_serialize_inference_frame_data_per_segment", "", _e5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegmentErr)
	}
	return _e5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegment(a0, a1, a2, a3), nil
}

// E5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegment.
func E5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegment(a0 uintptr, a1 *byte, a2 *byte, a3 bool) (int32, error) {
	return tryE5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegment(a0, a1, a2, a3)
}

var _e5rtExecutionStreamPrewireInUseAllocations func(a0 uintptr) int32
var _e5rtExecutionStreamPrewireInUseAllocationsErr error

func tryE5rtExecutionStreamPrewireInUseAllocations(a0 uintptr) (int32, error) {
	if _e5rtExecutionStreamPrewireInUseAllocations == nil {
		return 0, symbolCallError("e5rt_execution_stream_prewire_in_use_allocations", "", _e5rtExecutionStreamPrewireInUseAllocationsErr)
	}
	return _e5rtExecutionStreamPrewireInUseAllocations(a0), nil
}

// E5rtExecutionStreamPrewireInUseAllocations.
func E5rtExecutionStreamPrewireInUseAllocations(a0 uintptr) (int32, error) {
	return tryE5rtExecutionStreamPrewireInUseAllocations(a0)
}

var _e5rtExecutionStreamRelease func(out *uintptr) int32
var _e5rtExecutionStreamReleaseErr error

func tryE5rtExecutionStreamRelease(out *uintptr) (int32, error) {
	if _e5rtExecutionStreamRelease == nil {
		return 0, symbolCallError("e5rt_execution_stream_release", "", _e5rtExecutionStreamReleaseErr)
	}
	return _e5rtExecutionStreamRelease(out), nil
}

// E5rtExecutionStreamRelease.
func E5rtExecutionStreamRelease(out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamRelease(out)
}

var _e5rtExecutionStreamReset func(a0 uintptr) int32
var _e5rtExecutionStreamResetErr error

func tryE5rtExecutionStreamReset(a0 uintptr) (int32, error) {
	if _e5rtExecutionStreamReset == nil {
		return 0, symbolCallError("e5rt_execution_stream_reset", "", _e5rtExecutionStreamResetErr)
	}
	return _e5rtExecutionStreamReset(a0), nil
}

// E5rtExecutionStreamReset.
func E5rtExecutionStreamReset(a0 uintptr) (int32, error) {
	return tryE5rtExecutionStreamReset(a0)
}

var _e5rtExecutionStreamResetConfigOptions func(a0 uintptr, out *uintptr) int32
var _e5rtExecutionStreamResetConfigOptionsErr error

func tryE5rtExecutionStreamResetConfigOptions(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtExecutionStreamResetConfigOptions == nil {
		return 0, symbolCallError("e5rt_execution_stream_reset_config_options", "", _e5rtExecutionStreamResetConfigOptionsErr)
	}
	return _e5rtExecutionStreamResetConfigOptions(a0, out), nil
}

// E5rtExecutionStreamResetConfigOptions.
func E5rtExecutionStreamResetConfigOptions(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamResetConfigOptions(a0, out)
}

var _e5rtExecutionStreamSetAneExecutionPriority func(a0 uintptr, a1 uint32) int32
var _e5rtExecutionStreamSetAneExecutionPriorityErr error

func tryE5rtExecutionStreamSetAneExecutionPriority(a0 uintptr, a1 uint32) (int32, error) {
	if _e5rtExecutionStreamSetAneExecutionPriority == nil {
		return 0, symbolCallError("e5rt_execution_stream_set_ane_execution_priority", "", _e5rtExecutionStreamSetAneExecutionPriorityErr)
	}
	return _e5rtExecutionStreamSetAneExecutionPriority(a0, a1), nil
}

// E5rtExecutionStreamSetAneExecutionPriority.
func E5rtExecutionStreamSetAneExecutionPriority(a0 uintptr, a1 uint32) (int32, error) {
	return tryE5rtExecutionStreamSetAneExecutionPriority(a0, a1)
}

var _e5rtExecutionStreamSetConfigOptions func(a0 uintptr, a1 uintptr) int32
var _e5rtExecutionStreamSetConfigOptionsErr error

func tryE5rtExecutionStreamSetConfigOptions(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamSetConfigOptions == nil {
		return 0, symbolCallError("e5rt_execution_stream_set_config_options", "", _e5rtExecutionStreamSetConfigOptionsErr)
	}
	return _e5rtExecutionStreamSetConfigOptions(a0, a1), nil
}

// E5rtExecutionStreamSetConfigOptions.
func E5rtExecutionStreamSetConfigOptions(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamSetConfigOptions(a0, a1)
}

var _e5rtExecutionStreamSetQualityOfService func(a0 uintptr, a1 uint32) int32
var _e5rtExecutionStreamSetQualityOfServiceErr error

func tryE5rtExecutionStreamSetQualityOfService(a0 uintptr, a1 uint32) (int32, error) {
	if _e5rtExecutionStreamSetQualityOfService == nil {
		return 0, symbolCallError("e5rt_execution_stream_set_quality_of_service", "", _e5rtExecutionStreamSetQualityOfServiceErr)
	}
	return _e5rtExecutionStreamSetQualityOfService(a0, a1), nil
}

// E5rtExecutionStreamSetQualityOfService.
func E5rtExecutionStreamSetQualityOfService(a0 uintptr, a1 uint32) (int32, error) {
	return tryE5rtExecutionStreamSetQualityOfService(a0, a1)
}

var _e5rtExecutionStreamStepExecuteSync func(a0 uintptr, a1 uintptr) int32
var _e5rtExecutionStreamStepExecuteSyncErr error

func tryE5rtExecutionStreamStepExecuteSync(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamStepExecuteSync == nil {
		return 0, symbolCallError("e5rt_execution_stream_step_execute_sync", "", _e5rtExecutionStreamStepExecuteSyncErr)
	}
	return _e5rtExecutionStreamStepExecuteSync(a0, a1), nil
}

// E5rtExecutionStreamStepExecuteSync.
func E5rtExecutionStreamStepExecuteSync(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamStepExecuteSync(a0, a1)
}

var _e5rtExecutionStreamSubmitAsync func(a0 uintptr, a1 uintptr) int32
var _e5rtExecutionStreamSubmitAsyncErr error

func tryE5rtExecutionStreamSubmitAsync(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamSubmitAsync == nil {
		return 0, symbolCallError("e5rt_execution_stream_submit_async", "", _e5rtExecutionStreamSubmitAsyncErr)
	}
	return _e5rtExecutionStreamSubmitAsync(a0, a1), nil
}

// E5rtExecutionStreamSubmitAsync.
func E5rtExecutionStreamSubmitAsync(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamSubmitAsync(a0, a1)
}

var _e5rtExecutionStreamSubmitAsyncWithTimeout func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtExecutionStreamSubmitAsyncWithTimeoutErr error

func tryE5rtExecutionStreamSubmitAsyncWithTimeout(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtExecutionStreamSubmitAsyncWithTimeout == nil {
		return 0, symbolCallError("e5rt_execution_stream_submit_async_with_timeout", "", _e5rtExecutionStreamSubmitAsyncWithTimeoutErr)
	}
	return _e5rtExecutionStreamSubmitAsyncWithTimeout(a0, a1, a2), nil
}

// E5rtExecutionStreamSubmitAsyncWithTimeout.
func E5rtExecutionStreamSubmitAsyncWithTimeout(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtExecutionStreamSubmitAsyncWithTimeout(a0, a1, a2)
}

var _e5rtGetLastErrorMessage func() *byte
var _e5rtGetLastErrorMessageErr error

func tryE5rtGetLastErrorMessage() (*byte, error) {
	if _e5rtGetLastErrorMessage == nil {
		return nil, symbolCallError("e5rt_get_last_error_message", "", _e5rtGetLastErrorMessageErr)
	}
	return _e5rtGetLastErrorMessage(), nil
}

// E5rtGetLastErrorMessage.
func E5rtGetLastErrorMessage() (*byte, error) {
	return tryE5rtGetLastErrorMessage()
}

var _e5rtIOPortBindBufferObject func(a0 uintptr, a1 uintptr) int32
var _e5rtIOPortBindBufferObjectErr error

func tryE5rtIOPortBindBufferObject(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtIOPortBindBufferObject == nil {
		return 0, symbolCallError("e5rt_io_port_bind_buffer_object", "", _e5rtIOPortBindBufferObjectErr)
	}
	return _e5rtIOPortBindBufferObject(a0, a1), nil
}

// E5rtIOPortBindBufferObject.
func E5rtIOPortBindBufferObject(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtIOPortBindBufferObject(a0, a1)
}

var _e5rtIOPortBindMemoryObject func(a0 uintptr, a1 uintptr) int32
var _e5rtIOPortBindMemoryObjectErr error

func tryE5rtIOPortBindMemoryObject(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtIOPortBindMemoryObject == nil {
		return 0, symbolCallError("e5rt_io_port_bind_memory_object", "", _e5rtIOPortBindMemoryObjectErr)
	}
	return _e5rtIOPortBindMemoryObject(a0, a1), nil
}

// E5rtIOPortBindMemoryObject.
func E5rtIOPortBindMemoryObject(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtIOPortBindMemoryObject(a0, a1)
}

var _e5rtIOPortBindSurfaceObject func(a0 uintptr, a1 uintptr) int32
var _e5rtIOPortBindSurfaceObjectErr error

func tryE5rtIOPortBindSurfaceObject(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtIOPortBindSurfaceObject == nil {
		return 0, symbolCallError("e5rt_io_port_bind_surface_object", "", _e5rtIOPortBindSurfaceObjectErr)
	}
	return _e5rtIOPortBindSurfaceObject(a0, a1), nil
}

// E5rtIOPortBindSurfaceObject.
func E5rtIOPortBindSurfaceObject(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtIOPortBindSurfaceObject(a0, a1)
}

var _e5rtIOPortGetSupportedBufferTypes func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtIOPortGetSupportedBufferTypesErr error

func tryE5rtIOPortGetSupportedBufferTypes(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtIOPortGetSupportedBufferTypes == nil {
		return 0, symbolCallError("e5rt_io_port_get_supported_buffer_types", "", _e5rtIOPortGetSupportedBufferTypesErr)
	}
	return _e5rtIOPortGetSupportedBufferTypes(a0, a1, a2), nil
}

// E5rtIOPortGetSupportedBufferTypes.
func E5rtIOPortGetSupportedBufferTypes(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtIOPortGetSupportedBufferTypes(a0, a1, a2)
}

var _e5rtIOPortHasKnownShape func(a0 uintptr, out *uintptr) int32
var _e5rtIOPortHasKnownShapeErr error

func tryE5rtIOPortHasKnownShape(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtIOPortHasKnownShape == nil {
		return 0, symbolCallError("e5rt_io_port_has_known_shape", "", _e5rtIOPortHasKnownShapeErr)
	}
	return _e5rtIOPortHasKnownShape(a0, out), nil
}

// E5rtIOPortHasKnownShape.
func E5rtIOPortHasKnownShape(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtIOPortHasKnownShape(a0, out)
}

var _e5rtIOPortIsDynamic func(a0 uintptr, out *uintptr) int32
var _e5rtIOPortIsDynamicErr error

func tryE5rtIOPortIsDynamic(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtIOPortIsDynamic == nil {
		return 0, symbolCallError("e5rt_io_port_is_dynamic", "", _e5rtIOPortIsDynamicErr)
	}
	return _e5rtIOPortIsDynamic(a0, out), nil
}

// E5rtIOPortIsDynamic.
func E5rtIOPortIsDynamic(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtIOPortIsDynamic(a0, out)
}

var _e5rtIOPortIsSurface func(a0 uintptr, out *bool) int32
var _e5rtIOPortIsSurfaceErr error

func tryE5rtIOPortIsSurface(a0 uintptr, out *bool) (int32, error) {
	if _e5rtIOPortIsSurface == nil {
		return 0, symbolCallError("e5rt_io_port_is_surface", "", _e5rtIOPortIsSurfaceErr)
	}
	return _e5rtIOPortIsSurface(a0, out), nil
}

// E5rtIOPortIsSurface.
func E5rtIOPortIsSurface(a0 uintptr, out *bool) (int32, error) {
	return tryE5rtIOPortIsSurface(a0, out)
}

var _e5rtIOPortIsTensor func(a0 uintptr, out *bool) int32
var _e5rtIOPortIsTensorErr error

func tryE5rtIOPortIsTensor(a0 uintptr, out *bool) (int32, error) {
	if _e5rtIOPortIsTensor == nil {
		return 0, symbolCallError("e5rt_io_port_is_tensor", "", _e5rtIOPortIsTensorErr)
	}
	return _e5rtIOPortIsTensor(a0, out), nil
}

// E5rtIOPortIsTensor.
func E5rtIOPortIsTensor(a0 uintptr, out *bool) (int32, error) {
	return tryE5rtIOPortIsTensor(a0, out)
}

var _e5rtIOPortRelease func(out *uintptr) int32
var _e5rtIOPortReleaseErr error

func tryE5rtIOPortRelease(out *uintptr) (int32, error) {
	if _e5rtIOPortRelease == nil {
		return 0, symbolCallError("e5rt_io_port_release", "", _e5rtIOPortReleaseErr)
	}
	return _e5rtIOPortRelease(out), nil
}

// E5rtIOPortRelease.
func E5rtIOPortRelease(out *uintptr) (int32, error) {
	return tryE5rtIOPortRelease(out)
}

var _e5rtIOPortRetainBufferObject func(a0 uintptr, out *uintptr) int32
var _e5rtIOPortRetainBufferObjectErr error

func tryE5rtIOPortRetainBufferObject(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtIOPortRetainBufferObject == nil {
		return 0, symbolCallError("e5rt_io_port_retain_buffer_object", "", _e5rtIOPortRetainBufferObjectErr)
	}
	return _e5rtIOPortRetainBufferObject(a0, out), nil
}

// E5rtIOPortRetainBufferObject.
func E5rtIOPortRetainBufferObject(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtIOPortRetainBufferObject(a0, out)
}

var _e5rtIOPortRetainMemoryObject func(a0 uintptr, a1 uintptr) int32
var _e5rtIOPortRetainMemoryObjectErr error

func tryE5rtIOPortRetainMemoryObject(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtIOPortRetainMemoryObject == nil {
		return 0, symbolCallError("e5rt_io_port_retain_memory_object", "", _e5rtIOPortRetainMemoryObjectErr)
	}
	return _e5rtIOPortRetainMemoryObject(a0, a1), nil
}

// E5rtIOPortRetainMemoryObject.
func E5rtIOPortRetainMemoryObject(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtIOPortRetainMemoryObject(a0, a1)
}

var _e5rtIOPortRetainSurfaceDesc func(a0 uintptr, out *uintptr) int32
var _e5rtIOPortRetainSurfaceDescErr error

func tryE5rtIOPortRetainSurfaceDesc(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtIOPortRetainSurfaceDesc == nil {
		return 0, symbolCallError("e5rt_io_port_retain_surface_desc", "", _e5rtIOPortRetainSurfaceDescErr)
	}
	return _e5rtIOPortRetainSurfaceDesc(a0, out), nil
}

// E5rtIOPortRetainSurfaceDesc.
func E5rtIOPortRetainSurfaceDesc(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtIOPortRetainSurfaceDesc(a0, out)
}

var _e5rtIOPortRetainSurfaceObject func(a0 uintptr, out *uintptr) int32
var _e5rtIOPortRetainSurfaceObjectErr error

func tryE5rtIOPortRetainSurfaceObject(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtIOPortRetainSurfaceObject == nil {
		return 0, symbolCallError("e5rt_io_port_retain_surface_object", "", _e5rtIOPortRetainSurfaceObjectErr)
	}
	return _e5rtIOPortRetainSurfaceObject(a0, out), nil
}

// E5rtIOPortRetainSurfaceObject.
func E5rtIOPortRetainSurfaceObject(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtIOPortRetainSurfaceObject(a0, out)
}

var _e5rtIOPortRetainTensorDesc func(a0 uintptr, out *uintptr) int32
var _e5rtIOPortRetainTensorDescErr error

func tryE5rtIOPortRetainTensorDesc(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtIOPortRetainTensorDesc == nil {
		return 0, symbolCallError("e5rt_io_port_retain_tensor_desc", "", _e5rtIOPortRetainTensorDescErr)
	}
	return _e5rtIOPortRetainTensorDesc(a0, out), nil
}

// E5rtIOPortRetainTensorDesc.
func E5rtIOPortRetainTensorDesc(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtIOPortRetainTensorDesc(a0, out)
}

var _e5rtMemoryObjectCreate func(out *uintptr, a1 uint64, a2 uint32) int32
var _e5rtMemoryObjectCreateErr error

func tryE5rtMemoryObjectCreate(out *uintptr, a1 uint64, a2 uint32) (int32, error) {
	if _e5rtMemoryObjectCreate == nil {
		return 0, symbolCallError("e5rt_memory_object_create", "", _e5rtMemoryObjectCreateErr)
	}
	return _e5rtMemoryObjectCreate(out, a1, a2), nil
}

// E5rtMemoryObjectCreate.
func E5rtMemoryObjectCreate(out *uintptr, a1 uint64, a2 uint32) (int32, error) {
	return tryE5rtMemoryObjectCreate(out, a1, a2)
}

var _e5rtMemoryObjectCreateAsAlias func(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtMemoryObjectCreateAsAliasErr error

func tryE5rtMemoryObjectCreateAsAlias(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtMemoryObjectCreateAsAlias == nil {
		return 0, symbolCallError("e5rt_memory_object_create_as_alias", "", _e5rtMemoryObjectCreateAsAliasErr)
	}
	return _e5rtMemoryObjectCreateAsAlias(out, a1, a2, a3), nil
}

// E5rtMemoryObjectCreateAsAlias.
func E5rtMemoryObjectCreateAsAlias(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtMemoryObjectCreateAsAlias(out, a1, a2, a3)
}

var _e5rtMemoryObjectCreateFromIosurface func(out *uintptr, a1 uintptr) int32
var _e5rtMemoryObjectCreateFromIosurfaceErr error

func tryE5rtMemoryObjectCreateFromIosurface(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtMemoryObjectCreateFromIosurface == nil {
		return 0, symbolCallError("e5rt_memory_object_create_from_iosurface", "", _e5rtMemoryObjectCreateFromIosurfaceErr)
	}
	return _e5rtMemoryObjectCreateFromIosurface(out, a1), nil
}

// E5rtMemoryObjectCreateFromIosurface.
func E5rtMemoryObjectCreateFromIosurface(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtMemoryObjectCreateFromIosurface(out, a1)
}

var _e5rtMemoryObjectGetDataPtr func(a0 uintptr, a1 *uintptr) int32
var _e5rtMemoryObjectGetDataPtrErr error

func tryE5rtMemoryObjectGetDataPtr(a0 uintptr, a1 *uintptr) (int32, error) {
	if _e5rtMemoryObjectGetDataPtr == nil {
		return 0, symbolCallError("e5rt_memory_object_get_data_ptr", "", _e5rtMemoryObjectGetDataPtrErr)
	}
	return _e5rtMemoryObjectGetDataPtr(a0, a1), nil
}

// E5rtMemoryObjectGetDataPtr.
func E5rtMemoryObjectGetDataPtr(a0 uintptr, a1 *uintptr) (int32, error) {
	return tryE5rtMemoryObjectGetDataPtr(a0, a1)
}

var _e5rtMemoryObjectGetIosurface func(a0 uintptr, a1 *uintptr) int32
var _e5rtMemoryObjectGetIosurfaceErr error

func tryE5rtMemoryObjectGetIosurface(a0 uintptr, a1 *uintptr) (int32, error) {
	if _e5rtMemoryObjectGetIosurface == nil {
		return 0, symbolCallError("e5rt_memory_object_get_iosurface", "", _e5rtMemoryObjectGetIosurfaceErr)
	}
	return _e5rtMemoryObjectGetIosurface(a0, a1), nil
}

// E5rtMemoryObjectGetIosurface.
func E5rtMemoryObjectGetIosurface(a0 uintptr, a1 *uintptr) (int32, error) {
	return tryE5rtMemoryObjectGetIosurface(a0, a1)
}

var _e5rtMemoryObjectGetSize func(a0 uintptr, a1 *uint64) int32
var _e5rtMemoryObjectGetSizeErr error

func tryE5rtMemoryObjectGetSize(a0 uintptr, a1 *uint64) (int32, error) {
	if _e5rtMemoryObjectGetSize == nil {
		return 0, symbolCallError("e5rt_memory_object_get_size", "", _e5rtMemoryObjectGetSizeErr)
	}
	return _e5rtMemoryObjectGetSize(a0, a1), nil
}

// E5rtMemoryObjectGetSize.
func E5rtMemoryObjectGetSize(a0 uintptr, a1 *uint64) (int32, error) {
	return tryE5rtMemoryObjectGetSize(a0, a1)
}

var _e5rtMemoryObjectRelease func(out *uintptr) int32
var _e5rtMemoryObjectReleaseErr error

func tryE5rtMemoryObjectRelease(out *uintptr) (int32, error) {
	if _e5rtMemoryObjectRelease == nil {
		return 0, symbolCallError("e5rt_memory_object_release", "", _e5rtMemoryObjectReleaseErr)
	}
	return _e5rtMemoryObjectRelease(out), nil
}

// E5rtMemoryObjectRelease.
func E5rtMemoryObjectRelease(out *uintptr) (int32, error) {
	return tryE5rtMemoryObjectRelease(out)
}

var _e5rtOperandDescIsSurfaceDesc func(a0 uintptr, a1 *bool) int32
var _e5rtOperandDescIsSurfaceDescErr error

func tryE5rtOperandDescIsSurfaceDesc(a0 uintptr, a1 *bool) (int32, error) {
	if _e5rtOperandDescIsSurfaceDesc == nil {
		return 0, symbolCallError("e5rt_operand_desc_is_surface_desc", "", _e5rtOperandDescIsSurfaceDescErr)
	}
	return _e5rtOperandDescIsSurfaceDesc(a0, a1), nil
}

// E5rtOperandDescIsSurfaceDesc.
func E5rtOperandDescIsSurfaceDesc(a0 uintptr, a1 *bool) (int32, error) {
	return tryE5rtOperandDescIsSurfaceDesc(a0, a1)
}

var _e5rtOperandDescIsTensorDesc func(a0 uintptr, a1 *bool) int32
var _e5rtOperandDescIsTensorDescErr error

func tryE5rtOperandDescIsTensorDesc(a0 uintptr, a1 *bool) (int32, error) {
	if _e5rtOperandDescIsTensorDesc == nil {
		return 0, symbolCallError("e5rt_operand_desc_is_tensor_desc", "", _e5rtOperandDescIsTensorDescErr)
	}
	return _e5rtOperandDescIsTensorDesc(a0, a1), nil
}

// E5rtOperandDescIsTensorDesc.
func E5rtOperandDescIsTensorDesc(a0 uintptr, a1 *bool) (int32, error) {
	return tryE5rtOperandDescIsTensorDesc(a0, a1)
}

var _e5rtOperandDescRelease func(out *uintptr) int32
var _e5rtOperandDescReleaseErr error

func tryE5rtOperandDescRelease(out *uintptr) (int32, error) {
	if _e5rtOperandDescRelease == nil {
		return 0, symbolCallError("e5rt_operand_desc_release", "", _e5rtOperandDescReleaseErr)
	}
	return _e5rtOperandDescRelease(out), nil
}

// E5rtOperandDescRelease.
func E5rtOperandDescRelease(out *uintptr) (int32, error) {
	return tryE5rtOperandDescRelease(out)
}

var _e5rtOperandDescRetainFromSurfaceDesc func(out *uintptr, a1 uintptr) int32
var _e5rtOperandDescRetainFromSurfaceDescErr error

func tryE5rtOperandDescRetainFromSurfaceDesc(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtOperandDescRetainFromSurfaceDesc == nil {
		return 0, symbolCallError("e5rt_operand_desc_retain_from_surface_desc", "", _e5rtOperandDescRetainFromSurfaceDescErr)
	}
	return _e5rtOperandDescRetainFromSurfaceDesc(out, a1), nil
}

// E5rtOperandDescRetainFromSurfaceDesc.
func E5rtOperandDescRetainFromSurfaceDesc(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtOperandDescRetainFromSurfaceDesc(out, a1)
}

var _e5rtOperandDescRetainFromTensorDesc func(a0 uintptr, a1 uintptr) int32
var _e5rtOperandDescRetainFromTensorDescErr error

func tryE5rtOperandDescRetainFromTensorDesc(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtOperandDescRetainFromTensorDesc == nil {
		return 0, symbolCallError("e5rt_operand_desc_retain_from_tensor_desc", "", _e5rtOperandDescRetainFromTensorDescErr)
	}
	return _e5rtOperandDescRetainFromTensorDesc(a0, a1), nil
}

// E5rtOperandDescRetainFromTensorDesc.
func E5rtOperandDescRetainFromTensorDesc(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtOperandDescRetainFromTensorDesc(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsCopyDynamicCallables func(a0 uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsCopyDynamicCallablesErr error

func tryE5rtPrecompiledComputeOpCreateOptionsCopyDynamicCallables(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsCopyDynamicCallables == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_copy_dynamic_callables", "", _e5rtPrecompiledComputeOpCreateOptionsCopyDynamicCallablesErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsCopyDynamicCallables(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsCopyDynamicCallables.
func E5rtPrecompiledComputeOpCreateOptionsCopyDynamicCallables(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsCopyDynamicCallables(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsCopyMutableMilWeightPaths func(a0 uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsCopyMutableMilWeightPathsErr error

func tryE5rtPrecompiledComputeOpCreateOptionsCopyMutableMilWeightPaths(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsCopyMutableMilWeightPaths == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_copy_mutable_mil_weight_paths", "", _e5rtPrecompiledComputeOpCreateOptionsCopyMutableMilWeightPathsErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsCopyMutableMilWeightPaths(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsCopyMutableMilWeightPaths.
func E5rtPrecompiledComputeOpCreateOptionsCopyMutableMilWeightPaths(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsCopyMutableMilWeightPaths(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsCreate func(a0 uintptr, a1 *byte, a2 *byte) int32
var _e5rtPrecompiledComputeOpCreateOptionsCreateErr error

func tryE5rtPrecompiledComputeOpCreateOptionsCreate(a0 uintptr, a1 *byte, a2 *byte) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsCreate == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_create", "", _e5rtPrecompiledComputeOpCreateOptionsCreateErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsCreate(a0, a1, a2), nil
}

// E5rtPrecompiledComputeOpCreateOptionsCreate.
func E5rtPrecompiledComputeOpCreateOptionsCreate(a0 uintptr, a1 *byte, a2 *byte) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsCreate(a0, a1, a2)
}

var _e5rtPrecompiledComputeOpCreateOptionsCreateWithProgramFunction func(out *uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsCreateWithProgramFunctionErr error

func tryE5rtPrecompiledComputeOpCreateOptionsCreateWithProgramFunction(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsCreateWithProgramFunction == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_create_with_program_function", "", _e5rtPrecompiledComputeOpCreateOptionsCreateWithProgramFunctionErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsCreateWithProgramFunction(out, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsCreateWithProgramFunction.
func E5rtPrecompiledComputeOpCreateOptionsCreateWithProgramFunction(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsCreateWithProgramFunction(out, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsGetAllocateIntermediateBuffers func(a0 uintptr, out *uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsGetAllocateIntermediateBuffersErr error

func tryE5rtPrecompiledComputeOpCreateOptionsGetAllocateIntermediateBuffers(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsGetAllocateIntermediateBuffers == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_get_allocate_intermediate_buffers", "", _e5rtPrecompiledComputeOpCreateOptionsGetAllocateIntermediateBuffersErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsGetAllocateIntermediateBuffers(a0, out), nil
}

// E5rtPrecompiledComputeOpCreateOptionsGetAllocateIntermediateBuffers.
func E5rtPrecompiledComputeOpCreateOptionsGetAllocateIntermediateBuffers(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsGetAllocateIntermediateBuffers(a0, out)
}

var _e5rtPrecompiledComputeOpCreateOptionsGetExperimentalEnableMpsgraphParallelEncode func(a0 uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsGetExperimentalEnableMpsgraphParallelEncodeErr error

func tryE5rtPrecompiledComputeOpCreateOptionsGetExperimentalEnableMpsgraphParallelEncode(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsGetExperimentalEnableMpsgraphParallelEncode == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_get_experimental_enable_mpsgraph_parallel_encode", "", _e5rtPrecompiledComputeOpCreateOptionsGetExperimentalEnableMpsgraphParallelEncodeErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsGetExperimentalEnableMpsgraphParallelEncode(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsGetExperimentalEnableMpsgraphParallelEncode.
func E5rtPrecompiledComputeOpCreateOptionsGetExperimentalEnableMpsgraphParallelEncode(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsGetExperimentalEnableMpsgraphParallelEncode(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsGetIosurfaceMemoryPoolID func(a0 uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsGetIosurfaceMemoryPoolIDErr error

func tryE5rtPrecompiledComputeOpCreateOptionsGetIosurfaceMemoryPoolID(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsGetIosurfaceMemoryPoolID == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_get_iosurface_memory_pool_id", "", _e5rtPrecompiledComputeOpCreateOptionsGetIosurfaceMemoryPoolIDErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsGetIosurfaceMemoryPoolID(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsGetIosurfaceMemoryPoolID.
func E5rtPrecompiledComputeOpCreateOptionsGetIosurfaceMemoryPoolID(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsGetIosurfaceMemoryPoolID(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsGetLazyPrepareOpForEncode func(a0 uintptr, out *uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsGetLazyPrepareOpForEncodeErr error

func tryE5rtPrecompiledComputeOpCreateOptionsGetLazyPrepareOpForEncode(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsGetLazyPrepareOpForEncode == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_get_lazy_prepare_op_for_encode", "", _e5rtPrecompiledComputeOpCreateOptionsGetLazyPrepareOpForEncodeErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsGetLazyPrepareOpForEncode(a0, out), nil
}

// E5rtPrecompiledComputeOpCreateOptionsGetLazyPrepareOpForEncode.
func E5rtPrecompiledComputeOpCreateOptionsGetLazyPrepareOpForEncode(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsGetLazyPrepareOpForEncode(a0, out)
}

var _e5rtPrecompiledComputeOpCreateOptionsGetOperationName func(a0 uintptr, out *uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsGetOperationNameErr error

func tryE5rtPrecompiledComputeOpCreateOptionsGetOperationName(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsGetOperationName == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_get_operation_name", "", _e5rtPrecompiledComputeOpCreateOptionsGetOperationNameErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsGetOperationName(a0, out), nil
}

// E5rtPrecompiledComputeOpCreateOptionsGetOperationName.
func E5rtPrecompiledComputeOpCreateOptionsGetOperationName(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsGetOperationName(a0, out)
}

var _e5rtPrecompiledComputeOpCreateOptionsRelease func(out *uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsReleaseErr error

func tryE5rtPrecompiledComputeOpCreateOptionsRelease(out *uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsRelease == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_release", "", _e5rtPrecompiledComputeOpCreateOptionsReleaseErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsRelease(out), nil
}

// E5rtPrecompiledComputeOpCreateOptionsRelease.
func E5rtPrecompiledComputeOpCreateOptionsRelease(out *uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsRelease(out)
}

var _e5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDevice func(a0 uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDeviceErr error

func tryE5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDevice(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDevice == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_retain_override_compute_gpu_device", "", _e5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDeviceErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDevice(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDevice.
func E5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDevice(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDevice(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffers func(a0 uintptr, a1 bool) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffersErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffers(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffers == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_allocate_intermediate_buffers", "", _e5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffersErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffers(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffers.
func E5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffers(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffers(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProvider func(a0 uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProviderErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProvider(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProvider == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_custom_ane_memory_provider", "", _e5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProviderErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProvider(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProvider.
func E5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProvider(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProvider(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetDynamicCallables func(a0 uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetDynamicCallablesErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetDynamicCallables(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetDynamicCallables == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_dynamic_callables", "", _e5rtPrecompiledComputeOpCreateOptionsSetDynamicCallablesErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetDynamicCallables(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetDynamicCallables.
func E5rtPrecompiledComputeOpCreateOptionsSetDynamicCallables(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetDynamicCallables(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference func(a0 uintptr, a1 bool) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInferenceErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_experimental_disable_compile_time_mpsgraph_type_inference", "", _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInferenceErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference.
func E5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOps func(a0 uintptr, a1 bool) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOpsErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOps(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOps == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_experimental_enable_gpu_quant_ops", "", _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOpsErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOps(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOps.
func E5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOps(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOps(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecision func(a0 uintptr, a1 bool) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecisionErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecision(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecision == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_experimental_enable_mps_reduced_precision", "", _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecisionErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecision(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecision.
func E5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecision(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecision(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsgraphParallelEncode func(a0 uintptr, a1 bool) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsgraphParallelEncodeErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsgraphParallelEncode(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsgraphParallelEncode == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_experimental_enable_mpsgraph_parallel_encode", "", _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsgraphParallelEncodeErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsgraphParallelEncode(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsgraphParallelEncode.
func E5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsgraphParallelEncode(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsgraphParallelEncode(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalMpsgraphMaximumNumberOfEncodingThreads func(a0 uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalMpsgraphMaximumNumberOfEncodingThreadsErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetExperimentalMpsgraphMaximumNumberOfEncodingThreads(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalMpsgraphMaximumNumberOfEncodingThreads == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_experimental_mpsgraph_maximum_number_of_encoding_threads", "", _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalMpsgraphMaximumNumberOfEncodingThreadsErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalMpsgraphMaximumNumberOfEncodingThreads(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetExperimentalMpsgraphMaximumNumberOfEncodingThreads.
func E5rtPrecompiledComputeOpCreateOptionsSetExperimentalMpsgraphMaximumNumberOfEncodingThreads(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetExperimentalMpsgraphMaximumNumberOfEncodingThreads(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetIosurfaceMemoryPoolID func(a0 uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetIosurfaceMemoryPoolIDErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetIosurfaceMemoryPoolID(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetIosurfaceMemoryPoolID == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_iosurface_memory_pool_id", "", _e5rtPrecompiledComputeOpCreateOptionsSetIosurfaceMemoryPoolIDErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetIosurfaceMemoryPoolID(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetIosurfaceMemoryPoolID.
func E5rtPrecompiledComputeOpCreateOptionsSetIosurfaceMemoryPoolID(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetIosurfaceMemoryPoolID(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncode func(a0 uintptr, a1 bool) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncodeErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncode(a0 uintptr, a1 bool) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncode == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_lazy_prepare_op_for_encode", "", _e5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncodeErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncode(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncode.
func E5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncode(a0 uintptr, a1 bool) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncode(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPaths func(a0 uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPathsErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPaths(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPaths == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_mutable_mil_weight_paths", "", _e5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPathsErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPaths(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPaths.
func E5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPaths(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPaths(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetOperationName func(a0 uintptr, a1 *byte) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetOperationNameErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetOperationName(a0 uintptr, a1 *byte) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetOperationName == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_operation_name", "", _e5rtPrecompiledComputeOpCreateOptionsSetOperationNameErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetOperationName(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetOperationName.
func E5rtPrecompiledComputeOpCreateOptionsSetOperationName(a0 uintptr, a1 *byte) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetOperationName(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDevice func(a0 uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDeviceErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDevice(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDevice == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_override_compute_gpu_device", "", _e5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDeviceErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDevice(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDevice.
func E5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDevice(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDevice(a0, a1)
}

var _e5rtProgramFunctionGetExternInoutNames func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtProgramFunctionGetExternInoutNamesErr error

func tryE5rtProgramFunctionGetExternInoutNames(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtProgramFunctionGetExternInoutNames == nil {
		return 0, symbolCallError("e5rt_program_function_get_extern_inout_names", "", _e5rtProgramFunctionGetExternInoutNamesErr)
	}
	return _e5rtProgramFunctionGetExternInoutNames(a0, a1, a2), nil
}

// E5rtProgramFunctionGetExternInoutNames.
func E5rtProgramFunctionGetExternInoutNames(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtProgramFunctionGetExternInoutNames(a0, a1, a2)
}

var _e5rtProgramFunctionGetExternInputNames func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtProgramFunctionGetExternInputNamesErr error

func tryE5rtProgramFunctionGetExternInputNames(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtProgramFunctionGetExternInputNames == nil {
		return 0, symbolCallError("e5rt_program_function_get_extern_input_names", "", _e5rtProgramFunctionGetExternInputNamesErr)
	}
	return _e5rtProgramFunctionGetExternInputNames(a0, a1, a2), nil
}

// E5rtProgramFunctionGetExternInputNames.
func E5rtProgramFunctionGetExternInputNames(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtProgramFunctionGetExternInputNames(a0, a1, a2)
}

var _e5rtProgramFunctionGetExternOutputNames func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtProgramFunctionGetExternOutputNamesErr error

func tryE5rtProgramFunctionGetExternOutputNames(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtProgramFunctionGetExternOutputNames == nil {
		return 0, symbolCallError("e5rt_program_function_get_extern_output_names", "", _e5rtProgramFunctionGetExternOutputNamesErr)
	}
	return _e5rtProgramFunctionGetExternOutputNames(a0, a1, a2), nil
}

// E5rtProgramFunctionGetExternOutputNames.
func E5rtProgramFunctionGetExternOutputNames(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtProgramFunctionGetExternOutputNames(a0, a1, a2)
}

var _e5rtProgramFunctionGetName func(a0 uintptr, out **byte) int32
var _e5rtProgramFunctionGetNameErr error

func tryE5rtProgramFunctionGetName(a0 uintptr, out **byte) (int32, error) {
	if _e5rtProgramFunctionGetName == nil {
		return 0, symbolCallError("e5rt_program_function_get_name", "", _e5rtProgramFunctionGetNameErr)
	}
	return _e5rtProgramFunctionGetName(a0, out), nil
}

// E5rtProgramFunctionGetName.
func E5rtProgramFunctionGetName(a0 uintptr, out **byte) (int32, error) {
	return tryE5rtProgramFunctionGetName(a0, out)
}

var _e5rtProgramFunctionGetNumExternInouts func(a0 uintptr, a1 uintptr) int32
var _e5rtProgramFunctionGetNumExternInoutsErr error

func tryE5rtProgramFunctionGetNumExternInouts(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtProgramFunctionGetNumExternInouts == nil {
		return 0, symbolCallError("e5rt_program_function_get_num_extern_inouts", "", _e5rtProgramFunctionGetNumExternInoutsErr)
	}
	return _e5rtProgramFunctionGetNumExternInouts(a0, a1), nil
}

// E5rtProgramFunctionGetNumExternInouts.
func E5rtProgramFunctionGetNumExternInouts(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtProgramFunctionGetNumExternInouts(a0, a1)
}

var _e5rtProgramFunctionGetNumExternInputs func(a0 uintptr, a1 uintptr) int32
var _e5rtProgramFunctionGetNumExternInputsErr error

func tryE5rtProgramFunctionGetNumExternInputs(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtProgramFunctionGetNumExternInputs == nil {
		return 0, symbolCallError("e5rt_program_function_get_num_extern_inputs", "", _e5rtProgramFunctionGetNumExternInputsErr)
	}
	return _e5rtProgramFunctionGetNumExternInputs(a0, a1), nil
}

// E5rtProgramFunctionGetNumExternInputs.
func E5rtProgramFunctionGetNumExternInputs(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtProgramFunctionGetNumExternInputs(a0, a1)
}

var _e5rtProgramFunctionGetNumExternOutputs func(a0 uintptr, a1 uintptr) int32
var _e5rtProgramFunctionGetNumExternOutputsErr error

func tryE5rtProgramFunctionGetNumExternOutputs(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtProgramFunctionGetNumExternOutputs == nil {
		return 0, symbolCallError("e5rt_program_function_get_num_extern_outputs", "", _e5rtProgramFunctionGetNumExternOutputsErr)
	}
	return _e5rtProgramFunctionGetNumExternOutputs(a0, a1), nil
}

// E5rtProgramFunctionGetNumExternOutputs.
func E5rtProgramFunctionGetNumExternOutputs(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtProgramFunctionGetNumExternOutputs(a0, a1)
}

var _e5rtProgramFunctionLoadForExecution func(a0 uintptr) int32
var _e5rtProgramFunctionLoadForExecutionErr error

func tryE5rtProgramFunctionLoadForExecution(a0 uintptr) (int32, error) {
	if _e5rtProgramFunctionLoadForExecution == nil {
		return 0, symbolCallError("e5rt_program_function_load_for_execution", "", _e5rtProgramFunctionLoadForExecutionErr)
	}
	return _e5rtProgramFunctionLoadForExecution(a0), nil
}

// E5rtProgramFunctionLoadForExecution this function has a runtime HAZARD: runtime observation from the ANE lane on macOS 26.x: removed on macOS 26.x and returns status 2; harmless but dead.
func E5rtProgramFunctionLoadForExecution(a0 uintptr) (int32, error) {
	return tryE5rtProgramFunctionLoadForExecution(a0)
}

var _e5rtProgramFunctionRelease func(out *uintptr) int32
var _e5rtProgramFunctionReleaseErr error

func tryE5rtProgramFunctionRelease(out *uintptr) (int32, error) {
	if _e5rtProgramFunctionRelease == nil {
		return 0, symbolCallError("e5rt_program_function_release", "", _e5rtProgramFunctionReleaseErr)
	}
	return _e5rtProgramFunctionRelease(out), nil
}

// E5rtProgramFunctionRelease.
func E5rtProgramFunctionRelease(out *uintptr) (int32, error) {
	return tryE5rtProgramFunctionRelease(out)
}

var _e5rtProgramFunctionRetainExternInputIOPort func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtProgramFunctionRetainExternInputIOPortErr error

func tryE5rtProgramFunctionRetainExternInputIOPort(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtProgramFunctionRetainExternInputIOPort == nil {
		return 0, symbolCallError("e5rt_program_function_retain_extern_input_io_port", "", _e5rtProgramFunctionRetainExternInputIOPortErr)
	}
	return _e5rtProgramFunctionRetainExternInputIOPort(a0, a1, a2), nil
}

// E5rtProgramFunctionRetainExternInputIOPort.
func E5rtProgramFunctionRetainExternInputIOPort(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtProgramFunctionRetainExternInputIOPort(a0, a1, a2)
}

var _e5rtProgramFunctionRetainExternOutputIOPort func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtProgramFunctionRetainExternOutputIOPortErr error

func tryE5rtProgramFunctionRetainExternOutputIOPort(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtProgramFunctionRetainExternOutputIOPort == nil {
		return 0, symbolCallError("e5rt_program_function_retain_extern_output_io_port", "", _e5rtProgramFunctionRetainExternOutputIOPortErr)
	}
	return _e5rtProgramFunctionRetainExternOutputIOPort(a0, a1, a2), nil
}

// E5rtProgramFunctionRetainExternOutputIOPort.
func E5rtProgramFunctionRetainExternOutputIOPort(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtProgramFunctionRetainExternOutputIOPort(a0, a1, a2)
}

var _e5rtProgramFunctionRetainInoutSurfaceDesc func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtProgramFunctionRetainInoutSurfaceDescErr error

func tryE5rtProgramFunctionRetainInoutSurfaceDesc(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtProgramFunctionRetainInoutSurfaceDesc == nil {
		return 0, symbolCallError("e5rt_program_function_retain_inout_surface_desc", "", _e5rtProgramFunctionRetainInoutSurfaceDescErr)
	}
	return _e5rtProgramFunctionRetainInoutSurfaceDesc(a0, a1, a2), nil
}

// E5rtProgramFunctionRetainInoutSurfaceDesc.
func E5rtProgramFunctionRetainInoutSurfaceDesc(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtProgramFunctionRetainInoutSurfaceDesc(a0, a1, a2)
}

var _e5rtProgramFunctionRetainInoutTensorDesc func(a0 uintptr, a1 uintptr, out *uintptr) int32
var _e5rtProgramFunctionRetainInoutTensorDescErr error

func tryE5rtProgramFunctionRetainInoutTensorDesc(a0 uintptr, a1 uintptr, out *uintptr) (int32, error) {
	if _e5rtProgramFunctionRetainInoutTensorDesc == nil {
		return 0, symbolCallError("e5rt_program_function_retain_inout_tensor_desc", "", _e5rtProgramFunctionRetainInoutTensorDescErr)
	}
	return _e5rtProgramFunctionRetainInoutTensorDesc(a0, a1, out), nil
}

// E5rtProgramFunctionRetainInoutTensorDesc.
func E5rtProgramFunctionRetainInoutTensorDesc(a0 uintptr, a1 uintptr, out *uintptr) (int32, error) {
	return tryE5rtProgramFunctionRetainInoutTensorDesc(a0, a1, out)
}

var _e5rtProgramFunctionRetainInputSurfaceDesc func(a0 uintptr, a1 *byte, a2 uintptr) int32
var _e5rtProgramFunctionRetainInputSurfaceDescErr error

func tryE5rtProgramFunctionRetainInputSurfaceDesc(a0 uintptr, a1 *byte, a2 uintptr) (int32, error) {
	if _e5rtProgramFunctionRetainInputSurfaceDesc == nil {
		return 0, symbolCallError("e5rt_program_function_retain_input_surface_desc", "", _e5rtProgramFunctionRetainInputSurfaceDescErr)
	}
	return _e5rtProgramFunctionRetainInputSurfaceDesc(a0, a1, a2), nil
}

// E5rtProgramFunctionRetainInputSurfaceDesc.
func E5rtProgramFunctionRetainInputSurfaceDesc(a0 uintptr, a1 *byte, a2 uintptr) (int32, error) {
	return tryE5rtProgramFunctionRetainInputSurfaceDesc(a0, a1, a2)
}

var _e5rtProgramFunctionRetainInputTensorDesc func(a0 uintptr, a1 *byte, a2 uintptr) int32
var _e5rtProgramFunctionRetainInputTensorDescErr error

func tryE5rtProgramFunctionRetainInputTensorDesc(a0 uintptr, a1 *byte, a2 uintptr) (int32, error) {
	if _e5rtProgramFunctionRetainInputTensorDesc == nil {
		return 0, symbolCallError("e5rt_program_function_retain_input_tensor_desc", "", _e5rtProgramFunctionRetainInputTensorDescErr)
	}
	return _e5rtProgramFunctionRetainInputTensorDesc(a0, a1, a2), nil
}

// E5rtProgramFunctionRetainInputTensorDesc.
func E5rtProgramFunctionRetainInputTensorDesc(a0 uintptr, a1 *byte, a2 uintptr) (int32, error) {
	return tryE5rtProgramFunctionRetainInputTensorDesc(a0, a1, a2)
}

var _e5rtProgramFunctionRetainOutputSurfaceDesc func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtProgramFunctionRetainOutputSurfaceDescErr error

func tryE5rtProgramFunctionRetainOutputSurfaceDesc(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtProgramFunctionRetainOutputSurfaceDesc == nil {
		return 0, symbolCallError("e5rt_program_function_retain_output_surface_desc", "", _e5rtProgramFunctionRetainOutputSurfaceDescErr)
	}
	return _e5rtProgramFunctionRetainOutputSurfaceDesc(a0, a1, a2), nil
}

// E5rtProgramFunctionRetainOutputSurfaceDesc.
func E5rtProgramFunctionRetainOutputSurfaceDesc(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtProgramFunctionRetainOutputSurfaceDesc(a0, a1, a2)
}

var _e5rtProgramFunctionRetainOutputTensorDesc func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtProgramFunctionRetainOutputTensorDescErr error

func tryE5rtProgramFunctionRetainOutputTensorDesc(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtProgramFunctionRetainOutputTensorDesc == nil {
		return 0, symbolCallError("e5rt_program_function_retain_output_tensor_desc", "", _e5rtProgramFunctionRetainOutputTensorDescErr)
	}
	return _e5rtProgramFunctionRetainOutputTensorDesc(a0, a1, a2), nil
}

// E5rtProgramFunctionRetainOutputTensorDesc.
func E5rtProgramFunctionRetainOutputTensorDesc(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtProgramFunctionRetainOutputTensorDesc(a0, a1, a2)
}

var _e5rtProgramLibraryCreate func(out *uintptr, a1 *byte) int32
var _e5rtProgramLibraryCreateErr error

func tryE5rtProgramLibraryCreate(out *uintptr, a1 *byte) (int32, error) {
	if _e5rtProgramLibraryCreate == nil {
		return 0, symbolCallError("e5rt_program_library_create", "", _e5rtProgramLibraryCreateErr)
	}
	return _e5rtProgramLibraryCreate(out, a1), nil
}

// E5rtProgramLibraryCreate.
func E5rtProgramLibraryCreate(out *uintptr, a1 *byte) (int32, error) {
	return tryE5rtProgramLibraryCreate(out, a1)
}

var _e5rtProgramLibraryGetBuildInfo func(a0 uintptr, out *uintptr) int32
var _e5rtProgramLibraryGetBuildInfoErr error

func tryE5rtProgramLibraryGetBuildInfo(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtProgramLibraryGetBuildInfo == nil {
		return 0, symbolCallError("e5rt_program_library_get_build_info", "", _e5rtProgramLibraryGetBuildInfoErr)
	}
	return _e5rtProgramLibraryGetBuildInfo(a0, out), nil
}

// E5rtProgramLibraryGetBuildInfo.
func E5rtProgramLibraryGetBuildInfo(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtProgramLibraryGetBuildInfo(a0, out)
}

var _e5rtProgramLibraryGetE5BundlePath func(a0 uintptr, a1 **byte) int32
var _e5rtProgramLibraryGetE5BundlePathErr error

func tryE5rtProgramLibraryGetE5BundlePath(a0 uintptr, a1 **byte) (int32, error) {
	if _e5rtProgramLibraryGetE5BundlePath == nil {
		return 0, symbolCallError("e5rt_program_library_get_e5_bundle_path", "", _e5rtProgramLibraryGetE5BundlePathErr)
	}
	return _e5rtProgramLibraryGetE5BundlePath(a0, a1), nil
}

// E5rtProgramLibraryGetE5BundlePath.
func E5rtProgramLibraryGetE5BundlePath(a0 uintptr, a1 **byte) (int32, error) {
	return tryE5rtProgramLibraryGetE5BundlePath(a0, a1)
}

var _e5rtProgramLibraryGetFunctionMetadata func(a0 uintptr, a1 *byte, out *uintptr) int32
var _e5rtProgramLibraryGetFunctionMetadataErr error

func tryE5rtProgramLibraryGetFunctionMetadata(a0 uintptr, a1 *byte, out *uintptr) (int32, error) {
	if _e5rtProgramLibraryGetFunctionMetadata == nil {
		return 0, symbolCallError("e5rt_program_library_get_function_metadata", "", _e5rtProgramLibraryGetFunctionMetadataErr)
	}
	return _e5rtProgramLibraryGetFunctionMetadata(a0, a1, out), nil
}

// E5rtProgramLibraryGetFunctionMetadata.
func E5rtProgramLibraryGetFunctionMetadata(a0 uintptr, a1 *byte, out *uintptr) (int32, error) {
	return tryE5rtProgramLibraryGetFunctionMetadata(a0, a1, out)
}

var _e5rtProgramLibraryGetFunctionNames func(a0 uintptr, a1 uint64, out **byte) int32
var _e5rtProgramLibraryGetFunctionNamesErr error

func tryE5rtProgramLibraryGetFunctionNames(a0 uintptr, a1 uint64, out **byte) (int32, error) {
	if _e5rtProgramLibraryGetFunctionNames == nil {
		return 0, symbolCallError("e5rt_program_library_get_function_names", "", _e5rtProgramLibraryGetFunctionNamesErr)
	}
	return _e5rtProgramLibraryGetFunctionNames(a0, a1, out), nil
}

// E5rtProgramLibraryGetFunctionNames.
func E5rtProgramLibraryGetFunctionNames(a0 uintptr, a1 uint64, out **byte) (int32, error) {
	return tryE5rtProgramLibraryGetFunctionNames(a0, a1, out)
}

var _e5rtProgramLibraryGetNumFunctions func(a0 uintptr, out *uint64) int32
var _e5rtProgramLibraryGetNumFunctionsErr error

func tryE5rtProgramLibraryGetNumFunctions(a0 uintptr, out *uint64) (int32, error) {
	if _e5rtProgramLibraryGetNumFunctions == nil {
		return 0, symbolCallError("e5rt_program_library_get_num_functions", "", _e5rtProgramLibraryGetNumFunctionsErr)
	}
	return _e5rtProgramLibraryGetNumFunctions(a0, out), nil
}

// E5rtProgramLibraryGetNumFunctions.
func E5rtProgramLibraryGetNumFunctions(a0 uintptr, out *uint64) (int32, error) {
	return tryE5rtProgramLibraryGetNumFunctions(a0, out)
}

var _e5rtProgramLibraryGetSegmentationAnalytics func(a0 uintptr, out *uintptr) int32
var _e5rtProgramLibraryGetSegmentationAnalyticsErr error

func tryE5rtProgramLibraryGetSegmentationAnalytics(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtProgramLibraryGetSegmentationAnalytics == nil {
		return 0, symbolCallError("e5rt_program_library_get_segmentation_analytics", "", _e5rtProgramLibraryGetSegmentationAnalyticsErr)
	}
	return _e5rtProgramLibraryGetSegmentationAnalytics(a0, out), nil
}

// E5rtProgramLibraryGetSegmentationAnalytics.
func E5rtProgramLibraryGetSegmentationAnalytics(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtProgramLibraryGetSegmentationAnalytics(a0, out)
}

var _e5rtProgramLibraryRelease func(out *uintptr) int32
var _e5rtProgramLibraryReleaseErr error

func tryE5rtProgramLibraryRelease(out *uintptr) (int32, error) {
	if _e5rtProgramLibraryRelease == nil {
		return 0, symbolCallError("e5rt_program_library_release", "", _e5rtProgramLibraryReleaseErr)
	}
	return _e5rtProgramLibraryRelease(out), nil
}

// E5rtProgramLibraryRelease.
func E5rtProgramLibraryRelease(out *uintptr) (int32, error) {
	return tryE5rtProgramLibraryRelease(out)
}

var _e5rtProgramLibraryRetainProgramFunction func(a0 uintptr, a1 *byte, out *uintptr) int32
var _e5rtProgramLibraryRetainProgramFunctionErr error

func tryE5rtProgramLibraryRetainProgramFunction(a0 uintptr, a1 *byte, out *uintptr) (int32, error) {
	if _e5rtProgramLibraryRetainProgramFunction == nil {
		return 0, symbolCallError("e5rt_program_library_retain_program_function", "", _e5rtProgramLibraryRetainProgramFunctionErr)
	}
	return _e5rtProgramLibraryRetainProgramFunction(a0, a1, out), nil
}

// E5rtProgramLibraryRetainProgramFunction.
func E5rtProgramLibraryRetainProgramFunction(a0 uintptr, a1 *byte, out *uintptr) (int32, error) {
	return tryE5rtProgramLibraryRetainProgramFunction(a0, a1, out)
}

var _e5rtSurfaceDescCreate func(out *uintptr, a1 uint32, a2 uintptr, a3 uintptr) int32
var _e5rtSurfaceDescCreateErr error

func tryE5rtSurfaceDescCreate(out *uintptr, a1 uint32, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtSurfaceDescCreate == nil {
		return 0, symbolCallError("e5rt_surface_desc_create", "", _e5rtSurfaceDescCreateErr)
	}
	return _e5rtSurfaceDescCreate(out, a1, a2, a3), nil
}

// E5rtSurfaceDescCreate.
func E5rtSurfaceDescCreate(out *uintptr, a1 uint32, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtSurfaceDescCreate(out, a1, a2, a3)
}

var _e5rtSurfaceDescCreateFromOperandDesc func(a0 uintptr, a1 uintptr) int32
var _e5rtSurfaceDescCreateFromOperandDescErr error

func tryE5rtSurfaceDescCreateFromOperandDesc(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtSurfaceDescCreateFromOperandDesc == nil {
		return 0, symbolCallError("e5rt_surface_desc_create_from_operand_desc", "", _e5rtSurfaceDescCreateFromOperandDescErr)
	}
	return _e5rtSurfaceDescCreateFromOperandDesc(a0, a1), nil
}

// E5rtSurfaceDescCreateFromOperandDesc.
func E5rtSurfaceDescCreateFromOperandDesc(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtSurfaceDescCreateFromOperandDesc(a0, a1)
}

var _e5rtSurfaceDescCreateWithSlices func(a0 uintptr, a1 uint32, a2 uintptr, a3 uintptr, a4 uintptr) int32
var _e5rtSurfaceDescCreateWithSlicesErr error

func tryE5rtSurfaceDescCreateWithSlices(a0 uintptr, a1 uint32, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	if _e5rtSurfaceDescCreateWithSlices == nil {
		return 0, symbolCallError("e5rt_surface_desc_create_with_slices", "", _e5rtSurfaceDescCreateWithSlicesErr)
	}
	return _e5rtSurfaceDescCreateWithSlices(a0, a1, a2, a3, a4), nil
}

// E5rtSurfaceDescCreateWithSlices.
func E5rtSurfaceDescCreateWithSlices(a0 uintptr, a1 uint32, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	return tryE5rtSurfaceDescCreateWithSlices(a0, a1, a2, a3, a4)
}

var _e5rtSurfaceDescCreateWithStrides func(a0 uintptr, a1 uint32, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) int32
var _e5rtSurfaceDescCreateWithStridesErr error

func tryE5rtSurfaceDescCreateWithStrides(a0 uintptr, a1 uint32, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	if _e5rtSurfaceDescCreateWithStrides == nil {
		return 0, symbolCallError("e5rt_surface_desc_create_with_strides", "", _e5rtSurfaceDescCreateWithStridesErr)
	}
	return _e5rtSurfaceDescCreateWithStrides(a0, a1, a2, a3, a4, a5), nil
}

// E5rtSurfaceDescCreateWithStrides.
func E5rtSurfaceDescCreateWithStrides(a0 uintptr, a1 uint32, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	return tryE5rtSurfaceDescCreateWithStrides(a0, a1, a2, a3, a4, a5)
}

var _e5rtSurfaceDescCreateWithStridesAndSlices func(a0 uintptr, a1 uint32, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) int32
var _e5rtSurfaceDescCreateWithStridesAndSlicesErr error

func tryE5rtSurfaceDescCreateWithStridesAndSlices(a0 uintptr, a1 uint32, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) (int32, error) {
	if _e5rtSurfaceDescCreateWithStridesAndSlices == nil {
		return 0, symbolCallError("e5rt_surface_desc_create_with_strides_and_slices", "", _e5rtSurfaceDescCreateWithStridesAndSlicesErr)
	}
	return _e5rtSurfaceDescCreateWithStridesAndSlices(a0, a1, a2, a3, a4, a5, a6), nil
}

// E5rtSurfaceDescCreateWithStridesAndSlices.
func E5rtSurfaceDescCreateWithStridesAndSlices(a0 uintptr, a1 uint32, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) (int32, error) {
	return tryE5rtSurfaceDescCreateWithStridesAndSlices(a0, a1, a2, a3, a4, a5, a6)
}

var _e5rtSurfaceDescGetCustomRowStrides func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtSurfaceDescGetCustomRowStridesErr error

func tryE5rtSurfaceDescGetCustomRowStrides(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtSurfaceDescGetCustomRowStrides == nil {
		return 0, symbolCallError("e5rt_surface_desc_get_custom_row_strides", "", _e5rtSurfaceDescGetCustomRowStridesErr)
	}
	return _e5rtSurfaceDescGetCustomRowStrides(a0, a1, a2), nil
}

// E5rtSurfaceDescGetCustomRowStrides.
func E5rtSurfaceDescGetCustomRowStrides(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtSurfaceDescGetCustomRowStrides(a0, a1, a2)
}

var _e5rtSurfaceDescGetFormat func(a0 uintptr, a1 uintptr) int32
var _e5rtSurfaceDescGetFormatErr error

func tryE5rtSurfaceDescGetFormat(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtSurfaceDescGetFormat == nil {
		return 0, symbolCallError("e5rt_surface_desc_get_format", "", _e5rtSurfaceDescGetFormatErr)
	}
	return _e5rtSurfaceDescGetFormat(a0, a1), nil
}

// E5rtSurfaceDescGetFormat.
func E5rtSurfaceDescGetFormat(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtSurfaceDescGetFormat(a0, a1)
}

var _e5rtSurfaceDescGetHeight func(a0 uintptr, out *uintptr) int32
var _e5rtSurfaceDescGetHeightErr error

func tryE5rtSurfaceDescGetHeight(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtSurfaceDescGetHeight == nil {
		return 0, symbolCallError("e5rt_surface_desc_get_height", "", _e5rtSurfaceDescGetHeightErr)
	}
	return _e5rtSurfaceDescGetHeight(a0, out), nil
}

// E5rtSurfaceDescGetHeight.
func E5rtSurfaceDescGetHeight(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtSurfaceDescGetHeight(a0, out)
}

var _e5rtSurfaceDescGetPlaneCount func(a0 uintptr, a1 uintptr) int32
var _e5rtSurfaceDescGetPlaneCountErr error

func tryE5rtSurfaceDescGetPlaneCount(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtSurfaceDescGetPlaneCount == nil {
		return 0, symbolCallError("e5rt_surface_desc_get_plane_count", "", _e5rtSurfaceDescGetPlaneCountErr)
	}
	return _e5rtSurfaceDescGetPlaneCount(a0, a1), nil
}

// E5rtSurfaceDescGetPlaneCount.
func E5rtSurfaceDescGetPlaneCount(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtSurfaceDescGetPlaneCount(a0, a1)
}

var _e5rtSurfaceDescGetSliceCount func(a0 uintptr, a1 uintptr) int32
var _e5rtSurfaceDescGetSliceCountErr error

func tryE5rtSurfaceDescGetSliceCount(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtSurfaceDescGetSliceCount == nil {
		return 0, symbolCallError("e5rt_surface_desc_get_slice_count", "", _e5rtSurfaceDescGetSliceCountErr)
	}
	return _e5rtSurfaceDescGetSliceCount(a0, a1), nil
}

// E5rtSurfaceDescGetSliceCount.
func E5rtSurfaceDescGetSliceCount(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtSurfaceDescGetSliceCount(a0, a1)
}

var _e5rtSurfaceDescGetWidth func(a0 uintptr, out *uintptr) int32
var _e5rtSurfaceDescGetWidthErr error

func tryE5rtSurfaceDescGetWidth(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtSurfaceDescGetWidth == nil {
		return 0, symbolCallError("e5rt_surface_desc_get_width", "", _e5rtSurfaceDescGetWidthErr)
	}
	return _e5rtSurfaceDescGetWidth(a0, out), nil
}

// E5rtSurfaceDescGetWidth.
func E5rtSurfaceDescGetWidth(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtSurfaceDescGetWidth(a0, out)
}

var _e5rtSurfaceDescRelease func(out *uintptr) int32
var _e5rtSurfaceDescReleaseErr error

func tryE5rtSurfaceDescRelease(out *uintptr) (int32, error) {
	if _e5rtSurfaceDescRelease == nil {
		return 0, symbolCallError("e5rt_surface_desc_release", "", _e5rtSurfaceDescReleaseErr)
	}
	return _e5rtSurfaceDescRelease(out), nil
}

// E5rtSurfaceDescRelease.
func E5rtSurfaceDescRelease(out *uintptr) (int32, error) {
	return tryE5rtSurfaceDescRelease(out)
}

var _e5rtSurfaceFormatToCvpb4cc func(a0 uint32, a1 *uint32) int32
var _e5rtSurfaceFormatToCvpb4ccErr error

func tryE5rtSurfaceFormatToCvpb4cc(a0 uint32, a1 *uint32) (int32, error) {
	if _e5rtSurfaceFormatToCvpb4cc == nil {
		return 0, symbolCallError("e5rt_surface_format_to_cvpb_4cc", "", _e5rtSurfaceFormatToCvpb4ccErr)
	}
	return _e5rtSurfaceFormatToCvpb4cc(a0, a1), nil
}

// E5rtSurfaceFormatToCvpb4cc.
func E5rtSurfaceFormatToCvpb4cc(a0 uint32, a1 *uint32) (int32, error) {
	return tryE5rtSurfaceFormatToCvpb4cc(a0, a1)
}

var _e5rtSurfaceObjectAlloc func(out *uintptr, a1 uintptr, a2 uint32) int32
var _e5rtSurfaceObjectAllocErr error

func tryE5rtSurfaceObjectAlloc(out *uintptr, a1 uintptr, a2 uint32) (int32, error) {
	if _e5rtSurfaceObjectAlloc == nil {
		return 0, symbolCallError("e5rt_surface_object_alloc", "", _e5rtSurfaceObjectAllocErr)
	}
	return _e5rtSurfaceObjectAlloc(out, a1, a2), nil
}

// E5rtSurfaceObjectAlloc.
func E5rtSurfaceObjectAlloc(out *uintptr, a1 uintptr, a2 uint32) (int32, error) {
	return tryE5rtSurfaceObjectAlloc(out, a1, a2)
}

var _e5rtSurfaceObjectCreateFromIosurface func(out *uintptr, a1 uintptr) int32
var _e5rtSurfaceObjectCreateFromIosurfaceErr error

func tryE5rtSurfaceObjectCreateFromIosurface(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtSurfaceObjectCreateFromIosurface == nil {
		return 0, symbolCallError("e5rt_surface_object_create_from_iosurface", "", _e5rtSurfaceObjectCreateFromIosurfaceErr)
	}
	return _e5rtSurfaceObjectCreateFromIosurface(out, a1), nil
}

// E5rtSurfaceObjectCreateFromIosurface.
func E5rtSurfaceObjectCreateFromIosurface(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtSurfaceObjectCreateFromIosurface(out, a1)
}

var _e5rtSurfaceObjectGetIosurface func(a0 uintptr, a1 *uintptr) int32
var _e5rtSurfaceObjectGetIosurfaceErr error

func tryE5rtSurfaceObjectGetIosurface(a0 uintptr, a1 *uintptr) (int32, error) {
	if _e5rtSurfaceObjectGetIosurface == nil {
		return 0, symbolCallError("e5rt_surface_object_get_iosurface", "", _e5rtSurfaceObjectGetIosurfaceErr)
	}
	return _e5rtSurfaceObjectGetIosurface(a0, a1), nil
}

// E5rtSurfaceObjectGetIosurface.
func E5rtSurfaceObjectGetIosurface(a0 uintptr, a1 *uintptr) (int32, error) {
	return tryE5rtSurfaceObjectGetIosurface(a0, a1)
}

var _e5rtSurfaceObjectRelease func(out *uintptr) int32
var _e5rtSurfaceObjectReleaseErr error

func tryE5rtSurfaceObjectRelease(out *uintptr) (int32, error) {
	if _e5rtSurfaceObjectRelease == nil {
		return 0, symbolCallError("e5rt_surface_object_release", "", _e5rtSurfaceObjectReleaseErr)
	}
	return _e5rtSurfaceObjectRelease(out), nil
}

// E5rtSurfaceObjectRelease.
func E5rtSurfaceObjectRelease(out *uintptr) (int32, error) {
	return tryE5rtSurfaceObjectRelease(out)
}

var _e5rtTensorDescAllocBufferObject func(a0 uintptr, a1 uint32, a2 uint64, out *uintptr) int32
var _e5rtTensorDescAllocBufferObjectErr error

func tryE5rtTensorDescAllocBufferObject(a0 uintptr, a1 uint32, a2 uint64, out *uintptr) (int32, error) {
	if _e5rtTensorDescAllocBufferObject == nil {
		return 0, symbolCallError("e5rt_tensor_desc_alloc_buffer_object", "", _e5rtTensorDescAllocBufferObjectErr)
	}
	return _e5rtTensorDescAllocBufferObject(a0, a1, a2, out), nil
}

// E5rtTensorDescAllocBufferObject.
func E5rtTensorDescAllocBufferObject(a0 uintptr, a1 uint32, a2 uint64, out *uintptr) (int32, error) {
	return tryE5rtTensorDescAllocBufferObject(a0, a1, a2, out)
}

var _e5rtTensorDescCreate func(out *uintptr, a1 *uint64, a2 uint64, a3 uintptr) int32
var _e5rtTensorDescCreateErr error

func tryE5rtTensorDescCreate(out *uintptr, a1 *uint64, a2 uint64, a3 uintptr) (int32, error) {
	if _e5rtTensorDescCreate == nil {
		return 0, symbolCallError("e5rt_tensor_desc_create", "", _e5rtTensorDescCreateErr)
	}
	return _e5rtTensorDescCreate(out, a1, a2, a3), nil
}

// E5rtTensorDescCreate.
func E5rtTensorDescCreate(out *uintptr, a1 *uint64, a2 uint64, a3 uintptr) (int32, error) {
	return tryE5rtTensorDescCreate(out, a1, a2, a3)
}

var _e5rtTensorDescCreateFromOperandDesc func(a0 uintptr, a1 uintptr) int32
var _e5rtTensorDescCreateFromOperandDescErr error

func tryE5rtTensorDescCreateFromOperandDesc(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtTensorDescCreateFromOperandDesc == nil {
		return 0, symbolCallError("e5rt_tensor_desc_create_from_operand_desc", "", _e5rtTensorDescCreateFromOperandDescErr)
	}
	return _e5rtTensorDescCreateFromOperandDesc(a0, a1), nil
}

// E5rtTensorDescCreateFromOperandDesc.
func E5rtTensorDescCreateFromOperandDesc(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtTensorDescCreateFromOperandDesc(a0, a1)
}

var _e5rtTensorDescCreateMemoryObject func(a0 uintptr, a1 uint32, a2 uintptr, a3 uintptr) int32
var _e5rtTensorDescCreateMemoryObjectErr error

func tryE5rtTensorDescCreateMemoryObject(a0 uintptr, a1 uint32, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtTensorDescCreateMemoryObject == nil {
		return 0, symbolCallError("e5rt_tensor_desc_create_memory_object", "", _e5rtTensorDescCreateMemoryObjectErr)
	}
	return _e5rtTensorDescCreateMemoryObject(a0, a1, a2, a3), nil
}

// E5rtTensorDescCreateMemoryObject.
func E5rtTensorDescCreateMemoryObject(a0 uintptr, a1 uint32, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtTensorDescCreateMemoryObject(a0, a1, a2, a3)
}

var _e5rtTensorDescCreateSlice func(a0 uintptr, a1 uint64, a2 *uint64, out *uintptr) int32
var _e5rtTensorDescCreateSliceErr error

func tryE5rtTensorDescCreateSlice(a0 uintptr, a1 uint64, a2 *uint64, out *uintptr) (int32, error) {
	if _e5rtTensorDescCreateSlice == nil {
		return 0, symbolCallError("e5rt_tensor_desc_create_slice", "", _e5rtTensorDescCreateSliceErr)
	}
	return _e5rtTensorDescCreateSlice(a0, a1, a2, out), nil
}

// E5rtTensorDescCreateSlice.
func E5rtTensorDescCreateSlice(a0 uintptr, a1 uint64, a2 *uint64, out *uintptr) (int32, error) {
	return tryE5rtTensorDescCreateSlice(a0, a1, a2, out)
}

var _e5rtTensorDescCreateSliceWithLengths func(a0 uintptr, a1 uint64, a2 *uint64, a3 *uint64, out *uintptr) int32
var _e5rtTensorDescCreateSliceWithLengthsErr error

func tryE5rtTensorDescCreateSliceWithLengths(a0 uintptr, a1 uint64, a2 *uint64, a3 *uint64, out *uintptr) (int32, error) {
	if _e5rtTensorDescCreateSliceWithLengths == nil {
		return 0, symbolCallError("e5rt_tensor_desc_create_slice_with_lengths", "", _e5rtTensorDescCreateSliceWithLengthsErr)
	}
	return _e5rtTensorDescCreateSliceWithLengths(a0, a1, a2, a3, out), nil
}

// E5rtTensorDescCreateSliceWithLengths.
func E5rtTensorDescCreateSliceWithLengths(a0 uintptr, a1 uint64, a2 *uint64, a3 *uint64, out *uintptr) (int32, error) {
	return tryE5rtTensorDescCreateSliceWithLengths(a0, a1, a2, a3, out)
}

var _e5rtTensorDescCreateWithAlignments func(out *uintptr, a1 *uint64, a2 *uint64, a3 uint64, a4 uintptr) int32
var _e5rtTensorDescCreateWithAlignmentsErr error

func tryE5rtTensorDescCreateWithAlignments(out *uintptr, a1 *uint64, a2 *uint64, a3 uint64, a4 uintptr) (int32, error) {
	if _e5rtTensorDescCreateWithAlignments == nil {
		return 0, symbolCallError("e5rt_tensor_desc_create_with_alignments", "", _e5rtTensorDescCreateWithAlignmentsErr)
	}
	return _e5rtTensorDescCreateWithAlignments(out, a1, a2, a3, a4), nil
}

// E5rtTensorDescCreateWithAlignments.
func E5rtTensorDescCreateWithAlignments(out *uintptr, a1 *uint64, a2 *uint64, a3 uint64, a4 uintptr) (int32, error) {
	return tryE5rtTensorDescCreateWithAlignments(out, a1, a2, a3, a4)
}

var _e5rtTensorDescCreateWithStrides func(out *uintptr, a1 *uint64, a2 *int64, a3 uint64, a4 uintptr) int32
var _e5rtTensorDescCreateWithStridesErr error

func tryE5rtTensorDescCreateWithStrides(out *uintptr, a1 *uint64, a2 *int64, a3 uint64, a4 uintptr) (int32, error) {
	if _e5rtTensorDescCreateWithStrides == nil {
		return 0, symbolCallError("e5rt_tensor_desc_create_with_strides", "", _e5rtTensorDescCreateWithStridesErr)
	}
	return _e5rtTensorDescCreateWithStrides(out, a1, a2, a3, a4), nil
}

// E5rtTensorDescCreateWithStrides.
func E5rtTensorDescCreateWithStrides(out *uintptr, a1 *uint64, a2 *int64, a3 uint64, a4 uintptr) (int32, error) {
	return tryE5rtTensorDescCreateWithStrides(out, a1, a2, a3, a4)
}

var _e5rtTensorDescDtypeAreEqual func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtTensorDescDtypeAreEqualErr error

func tryE5rtTensorDescDtypeAreEqual(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtTensorDescDtypeAreEqual == nil {
		return 0, symbolCallError("e5rt_tensor_desc_dtype_are_equal", "", _e5rtTensorDescDtypeAreEqualErr)
	}
	return _e5rtTensorDescDtypeAreEqual(a0, a1, a2), nil
}

// E5rtTensorDescDtypeAreEqual.
func E5rtTensorDescDtypeAreEqual(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtTensorDescDtypeAreEqual(a0, a1, a2)
}

var _e5rtTensorDescDtypeCreate func(out *uintptr, a1 uint32, a2 uint32) int32
var _e5rtTensorDescDtypeCreateErr error

func tryE5rtTensorDescDtypeCreate(out *uintptr, a1 uint32, a2 uint32) (int32, error) {
	if _e5rtTensorDescDtypeCreate == nil {
		return 0, symbolCallError("e5rt_tensor_desc_dtype_create", "", _e5rtTensorDescDtypeCreateErr)
	}
	return _e5rtTensorDescDtypeCreate(out, a1, a2), nil
}

// E5rtTensorDescDtypeCreate.
func E5rtTensorDescDtypeCreate(out *uintptr, a1 uint32, a2 uint32) (int32, error) {
	return tryE5rtTensorDescDtypeCreate(out, a1, a2)
}

var _e5rtTensorDescDtypeGetComponentDtype func(a0 uintptr, out *uint32) int32
var _e5rtTensorDescDtypeGetComponentDtypeErr error

func tryE5rtTensorDescDtypeGetComponentDtype(a0 uintptr, out *uint32) (int32, error) {
	if _e5rtTensorDescDtypeGetComponentDtype == nil {
		return 0, symbolCallError("e5rt_tensor_desc_dtype_get_component_dtype", "", _e5rtTensorDescDtypeGetComponentDtypeErr)
	}
	return _e5rtTensorDescDtypeGetComponentDtype(a0, out), nil
}

// E5rtTensorDescDtypeGetComponentDtype.
func E5rtTensorDescDtypeGetComponentDtype(a0 uintptr, out *uint32) (int32, error) {
	return tryE5rtTensorDescDtypeGetComponentDtype(a0, out)
}

var _e5rtTensorDescDtypeGetComponentPack func(a0 uintptr, out *uint32) int32
var _e5rtTensorDescDtypeGetComponentPackErr error

func tryE5rtTensorDescDtypeGetComponentPack(a0 uintptr, out *uint32) (int32, error) {
	if _e5rtTensorDescDtypeGetComponentPack == nil {
		return 0, symbolCallError("e5rt_tensor_desc_dtype_get_component_pack", "", _e5rtTensorDescDtypeGetComponentPackErr)
	}
	return _e5rtTensorDescDtypeGetComponentPack(a0, out), nil
}

// E5rtTensorDescDtypeGetComponentPack.
func E5rtTensorDescDtypeGetComponentPack(a0 uintptr, out *uint32) (int32, error) {
	return tryE5rtTensorDescDtypeGetComponentPack(a0, out)
}

var _e5rtTensorDescDtypeGetComponentSize func(a0 uintptr, out *uint64) int32
var _e5rtTensorDescDtypeGetComponentSizeErr error

func tryE5rtTensorDescDtypeGetComponentSize(a0 uintptr, out *uint64) (int32, error) {
	if _e5rtTensorDescDtypeGetComponentSize == nil {
		return 0, symbolCallError("e5rt_tensor_desc_dtype_get_component_size", "", _e5rtTensorDescDtypeGetComponentSizeErr)
	}
	return _e5rtTensorDescDtypeGetComponentSize(a0, out), nil
}

// E5rtTensorDescDtypeGetComponentSize.
func E5rtTensorDescDtypeGetComponentSize(a0 uintptr, out *uint64) (int32, error) {
	return tryE5rtTensorDescDtypeGetComponentSize(a0, out)
}

var _e5rtTensorDescDtypeGetElementSize func(a0 uintptr, out *uint64) int32
var _e5rtTensorDescDtypeGetElementSizeErr error

func tryE5rtTensorDescDtypeGetElementSize(a0 uintptr, out *uint64) (int32, error) {
	if _e5rtTensorDescDtypeGetElementSize == nil {
		return 0, symbolCallError("e5rt_tensor_desc_dtype_get_element_size", "", _e5rtTensorDescDtypeGetElementSizeErr)
	}
	return _e5rtTensorDescDtypeGetElementSize(a0, out), nil
}

// E5rtTensorDescDtypeGetElementSize.
func E5rtTensorDescDtypeGetElementSize(a0 uintptr, out *uint64) (int32, error) {
	return tryE5rtTensorDescDtypeGetElementSize(a0, out)
}

var _e5rtTensorDescDtypeGetNumComponents func(a0 uintptr, out *uint8) int32
var _e5rtTensorDescDtypeGetNumComponentsErr error

func tryE5rtTensorDescDtypeGetNumComponents(a0 uintptr, out *uint8) (int32, error) {
	if _e5rtTensorDescDtypeGetNumComponents == nil {
		return 0, symbolCallError("e5rt_tensor_desc_dtype_get_num_components", "", _e5rtTensorDescDtypeGetNumComponentsErr)
	}
	return _e5rtTensorDescDtypeGetNumComponents(a0, out), nil
}

// E5rtTensorDescDtypeGetNumComponents.
func E5rtTensorDescDtypeGetNumComponents(a0 uintptr, out *uint8) (int32, error) {
	return tryE5rtTensorDescDtypeGetNumComponents(a0, out)
}

var _e5rtTensorDescDtypeRelease func(out *uintptr) int32
var _e5rtTensorDescDtypeReleaseErr error

func tryE5rtTensorDescDtypeRelease(out *uintptr) (int32, error) {
	if _e5rtTensorDescDtypeRelease == nil {
		return 0, symbolCallError("e5rt_tensor_desc_dtype_release", "", _e5rtTensorDescDtypeReleaseErr)
	}
	return _e5rtTensorDescDtypeRelease(out), nil
}

// E5rtTensorDescDtypeRelease.
func E5rtTensorDescDtypeRelease(out *uintptr) (int32, error) {
	return tryE5rtTensorDescDtypeRelease(out)
}

var _e5rtTensorDescDtypeValidateSpec func(component_data_type uint32, component_pack uint32, out *uint8) int32
var _e5rtTensorDescDtypeValidateSpecErr error

func tryE5rtTensorDescDtypeValidateSpec(component_data_type uint32, component_pack uint32, out *uint8) (int32, error) {
	if _e5rtTensorDescDtypeValidateSpec == nil {
		return 0, symbolCallError("e5rt_tensor_desc_dtype_validate_spec", "", _e5rtTensorDescDtypeValidateSpecErr)
	}
	return _e5rtTensorDescDtypeValidateSpec(component_data_type, component_pack, out), nil
}

// E5rtTensorDescDtypeValidateSpec.
func E5rtTensorDescDtypeValidateSpec(component_data_type uint32, component_pack uint32, out *uint8) (int32, error) {
	return tryE5rtTensorDescDtypeValidateSpec(component_data_type, component_pack, out)
}

var _e5rtTensorDescGetByteOffset func(a0 uintptr, a1 *uint64, a2 uint64, out *uint64) int32
var _e5rtTensorDescGetByteOffsetErr error

func tryE5rtTensorDescGetByteOffset(a0 uintptr, a1 *uint64, a2 uint64, out *uint64) (int32, error) {
	if _e5rtTensorDescGetByteOffset == nil {
		return 0, symbolCallError("e5rt_tensor_desc_get_byte_offset", "", _e5rtTensorDescGetByteOffsetErr)
	}
	return _e5rtTensorDescGetByteOffset(a0, a1, a2, out), nil
}

// E5rtTensorDescGetByteOffset.
func E5rtTensorDescGetByteOffset(a0 uintptr, a1 *uint64, a2 uint64, out *uint64) (int32, error) {
	return tryE5rtTensorDescGetByteOffset(a0, a1, a2, out)
}

var _e5rtTensorDescGetDimensionLength func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtTensorDescGetDimensionLengthErr error

func tryE5rtTensorDescGetDimensionLength(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtTensorDescGetDimensionLength == nil {
		return 0, symbolCallError("e5rt_tensor_desc_get_dimension_length", "", _e5rtTensorDescGetDimensionLengthErr)
	}
	return _e5rtTensorDescGetDimensionLength(a0, a1, a2), nil
}

// E5rtTensorDescGetDimensionLength.
func E5rtTensorDescGetDimensionLength(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtTensorDescGetDimensionLength(a0, a1, a2)
}

var _e5rtTensorDescGetDimensionStride func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtTensorDescGetDimensionStrideErr error

func tryE5rtTensorDescGetDimensionStride(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtTensorDescGetDimensionStride == nil {
		return 0, symbolCallError("e5rt_tensor_desc_get_dimension_stride", "", _e5rtTensorDescGetDimensionStrideErr)
	}
	return _e5rtTensorDescGetDimensionStride(a0, a1, a2), nil
}

// E5rtTensorDescGetDimensionStride.
func E5rtTensorDescGetDimensionStride(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtTensorDescGetDimensionStride(a0, a1, a2)
}

var _e5rtTensorDescGetNumElements func(a0 uintptr, out *uint64) int32
var _e5rtTensorDescGetNumElementsErr error

func tryE5rtTensorDescGetNumElements(a0 uintptr, out *uint64) (int32, error) {
	if _e5rtTensorDescGetNumElements == nil {
		return 0, symbolCallError("e5rt_tensor_desc_get_num_elements", "", _e5rtTensorDescGetNumElementsErr)
	}
	return _e5rtTensorDescGetNumElements(a0, out), nil
}

// E5rtTensorDescGetNumElements.
func E5rtTensorDescGetNumElements(a0 uintptr, out *uint64) (int32, error) {
	return tryE5rtTensorDescGetNumElements(a0, out)
}

var _e5rtTensorDescGetRank func(a0 uintptr, out *uint64) int32
var _e5rtTensorDescGetRankErr error

func tryE5rtTensorDescGetRank(a0 uintptr, out *uint64) (int32, error) {
	if _e5rtTensorDescGetRank == nil {
		return 0, symbolCallError("e5rt_tensor_desc_get_rank", "", _e5rtTensorDescGetRankErr)
	}
	return _e5rtTensorDescGetRank(a0, out), nil
}

// E5rtTensorDescGetRank.
func E5rtTensorDescGetRank(a0 uintptr, out *uint64) (int32, error) {
	return tryE5rtTensorDescGetRank(a0, out)
}

var _e5rtTensorDescGetShape func(a0 uintptr, a1 *uint64, a2 **uint64) int32
var _e5rtTensorDescGetShapeErr error

func tryE5rtTensorDescGetShape(a0 uintptr, a1 *uint64, a2 **uint64) (int32, error) {
	if _e5rtTensorDescGetShape == nil {
		return 0, symbolCallError("e5rt_tensor_desc_get_shape", "", _e5rtTensorDescGetShapeErr)
	}
	return _e5rtTensorDescGetShape(a0, a1, a2), nil
}

// E5rtTensorDescGetShape.
func E5rtTensorDescGetShape(a0 uintptr, a1 *uint64, a2 **uint64) (int32, error) {
	return tryE5rtTensorDescGetShape(a0, a1, a2)
}

var _e5rtTensorDescGetSize func(a0 uintptr, out *uint64) int32
var _e5rtTensorDescGetSizeErr error

func tryE5rtTensorDescGetSize(a0 uintptr, out *uint64) (int32, error) {
	if _e5rtTensorDescGetSize == nil {
		return 0, symbolCallError("e5rt_tensor_desc_get_size", "", _e5rtTensorDescGetSizeErr)
	}
	return _e5rtTensorDescGetSize(a0, out), nil
}

// E5rtTensorDescGetSize.
func E5rtTensorDescGetSize(a0 uintptr, out *uint64) (int32, error) {
	return tryE5rtTensorDescGetSize(a0, out)
}

var _e5rtTensorDescGetStrides func(a0 uintptr, a1 *uint64, a2 **int64) int32
var _e5rtTensorDescGetStridesErr error

func tryE5rtTensorDescGetStrides(a0 uintptr, a1 *uint64, a2 **int64) (int32, error) {
	if _e5rtTensorDescGetStrides == nil {
		return 0, symbolCallError("e5rt_tensor_desc_get_strides", "", _e5rtTensorDescGetStridesErr)
	}
	return _e5rtTensorDescGetStrides(a0, a1, a2), nil
}

// E5rtTensorDescGetStrides.
func E5rtTensorDescGetStrides(a0 uintptr, a1 *uint64, a2 **int64) (int32, error) {
	return tryE5rtTensorDescGetStrides(a0, a1, a2)
}

var _e5rtTensorDescHasKnownShape func(a0 uintptr, out *uint8) int32
var _e5rtTensorDescHasKnownShapeErr error

func tryE5rtTensorDescHasKnownShape(a0 uintptr, out *uint8) (int32, error) {
	if _e5rtTensorDescHasKnownShape == nil {
		return 0, symbolCallError("e5rt_tensor_desc_has_known_shape", "", _e5rtTensorDescHasKnownShapeErr)
	}
	return _e5rtTensorDescHasKnownShape(a0, out), nil
}

// E5rtTensorDescHasKnownShape.
func E5rtTensorDescHasKnownShape(a0 uintptr, out *uint8) (int32, error) {
	return tryE5rtTensorDescHasKnownShape(a0, out)
}

var _e5rtTensorDescRelease func(out *uintptr) int32
var _e5rtTensorDescReleaseErr error

func tryE5rtTensorDescRelease(out *uintptr) (int32, error) {
	if _e5rtTensorDescRelease == nil {
		return 0, symbolCallError("e5rt_tensor_desc_release", "", _e5rtTensorDescReleaseErr)
	}
	return _e5rtTensorDescRelease(out), nil
}

// E5rtTensorDescRelease.
func E5rtTensorDescRelease(out *uintptr) (int32, error) {
	return tryE5rtTensorDescRelease(out)
}

var _e5rtTensorDescRetainDtype func(a0 uintptr, a1 uintptr) int32
var _e5rtTensorDescRetainDtypeErr error

func tryE5rtTensorDescRetainDtype(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtTensorDescRetainDtype == nil {
		return 0, symbolCallError("e5rt_tensor_desc_retain_dtype", "", _e5rtTensorDescRetainDtypeErr)
	}
	return _e5rtTensorDescRetainDtype(a0, a1), nil
}

// E5rtTensorDescRetainDtype.
func E5rtTensorDescRetainDtype(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtTensorDescRetainDtype(a0, a1)
}

var _e5rtTensorUtilsAreTensorsEqual func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) int32
var _e5rtTensorUtilsAreTensorsEqualErr error

func tryE5rtTensorUtilsAreTensorsEqual(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	if _e5rtTensorUtilsAreTensorsEqual == nil {
		return 0, symbolCallError("e5rt_tensor_utils_are_tensors_equal", "", _e5rtTensorUtilsAreTensorsEqualErr)
	}
	return _e5rtTensorUtilsAreTensorsEqual(a0, a1, a2, a3, a4), nil
}

// E5rtTensorUtilsAreTensorsEqual.
func E5rtTensorUtilsAreTensorsEqual(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	return tryE5rtTensorUtilsAreTensorsEqual(a0, a1, a2, a3, a4)
}

var _e5rtTensorUtilsCastFromFp16ToFp32 func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtTensorUtilsCastFromFp16ToFp32Err error

func tryE5rtTensorUtilsCastFromFp16ToFp32(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtTensorUtilsCastFromFp16ToFp32 == nil {
		return 0, symbolCallError("e5rt_tensor_utils_cast_from_fp16_to_fp32", "", _e5rtTensorUtilsCastFromFp16ToFp32Err)
	}
	return _e5rtTensorUtilsCastFromFp16ToFp32(a0, a1, a2, a3), nil
}

// E5rtTensorUtilsCastFromFp16ToFp32.
func E5rtTensorUtilsCastFromFp16ToFp32(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtTensorUtilsCastFromFp16ToFp32(a0, a1, a2, a3)
}

var _e5rtTensorUtilsCastFromFp32ToFp16 func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtTensorUtilsCastFromFp32ToFp16Err error

func tryE5rtTensorUtilsCastFromFp32ToFp16(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtTensorUtilsCastFromFp32ToFp16 == nil {
		return 0, symbolCallError("e5rt_tensor_utils_cast_from_fp32_to_fp16", "", _e5rtTensorUtilsCastFromFp32ToFp16Err)
	}
	return _e5rtTensorUtilsCastFromFp32ToFp16(a0, a1, a2, a3), nil
}

// E5rtTensorUtilsCastFromFp32ToFp16.
func E5rtTensorUtilsCastFromFp32ToFp16(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtTensorUtilsCastFromFp32ToFp16(a0, a1, a2, a3)
}

var _e5rtTensorUtilsCopyTensor func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtTensorUtilsCopyTensorErr error

func tryE5rtTensorUtilsCopyTensor(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtTensorUtilsCopyTensor == nil {
		return 0, symbolCallError("e5rt_tensor_utils_copy_tensor", "", _e5rtTensorUtilsCopyTensorErr)
	}
	return _e5rtTensorUtilsCopyTensor(a0, a1, a2, a3), nil
}

// E5rtTensorUtilsCopyTensor.
func E5rtTensorUtilsCopyTensor(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtTensorUtilsCopyTensor(a0, a1, a2, a3)
}

var _e5rtTensorUtilsGetFp16Element func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uint8, a5 *uintptr) int32
var _e5rtTensorUtilsGetFp16ElementErr error

func tryE5rtTensorUtilsGetFp16Element(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uint8, a5 *uintptr) (int32, error) {
	if _e5rtTensorUtilsGetFp16Element == nil {
		return 0, symbolCallError("e5rt_tensor_utils_get_fp16_element", "", _e5rtTensorUtilsGetFp16ElementErr)
	}
	return _e5rtTensorUtilsGetFp16Element(a0, a1, a2, a3, a4, a5), nil
}

// E5rtTensorUtilsGetFp16Element.
func E5rtTensorUtilsGetFp16Element(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uint8, a5 *uintptr) (int32, error) {
	return tryE5rtTensorUtilsGetFp16Element(a0, a1, a2, a3, a4, a5)
}

var _e5rtTensorUtilsGetFp32Element func(a0 uintptr, a1 uintptr, a2 *uint64, a3 uint64, a4 uint8, out *float32) int32
var _e5rtTensorUtilsGetFp32ElementErr error

func tryE5rtTensorUtilsGetFp32Element(a0 uintptr, a1 uintptr, a2 *uint64, a3 uint64, a4 uint8, out *float32) (int32, error) {
	if _e5rtTensorUtilsGetFp32Element == nil {
		return 0, symbolCallError("e5rt_tensor_utils_get_fp32_element", "", _e5rtTensorUtilsGetFp32ElementErr)
	}
	return _e5rtTensorUtilsGetFp32Element(a0, a1, a2, a3, a4, out), nil
}

// E5rtTensorUtilsGetFp32Element.
func E5rtTensorUtilsGetFp32Element(a0 uintptr, a1 uintptr, a2 *uint64, a3 uint64, a4 uint8, out *float32) (int32, error) {
	return tryE5rtTensorUtilsGetFp32Element(a0, a1, a2, a3, a4, out)
}

var _e5rtTensorUtilsGetS8Element func(a0 uintptr, a1 uintptr, a2 *uint64, a3 uint64, a4 uint8, out *int8) int32
var _e5rtTensorUtilsGetS8ElementErr error

func tryE5rtTensorUtilsGetS8Element(a0 uintptr, a1 uintptr, a2 *uint64, a3 uint64, a4 uint8, out *int8) (int32, error) {
	if _e5rtTensorUtilsGetS8Element == nil {
		return 0, symbolCallError("e5rt_tensor_utils_get_s8_element", "", _e5rtTensorUtilsGetS8ElementErr)
	}
	return _e5rtTensorUtilsGetS8Element(a0, a1, a2, a3, a4, out), nil
}

// E5rtTensorUtilsGetS8Element.
func E5rtTensorUtilsGetS8Element(a0 uintptr, a1 uintptr, a2 *uint64, a3 uint64, a4 uint8, out *int8) (int32, error) {
	return tryE5rtTensorUtilsGetS8Element(a0, a1, a2, a3, a4, out)
}

var _e5rtTensorUtilsGetU8Element func(a0 uintptr, a1 uintptr, a2 *uint64, a3 uint64, a4 uint8, out *uint8) int32
var _e5rtTensorUtilsGetU8ElementErr error

func tryE5rtTensorUtilsGetU8Element(a0 uintptr, a1 uintptr, a2 *uint64, a3 uint64, a4 uint8, out *uint8) (int32, error) {
	if _e5rtTensorUtilsGetU8Element == nil {
		return 0, symbolCallError("e5rt_tensor_utils_get_u8_element", "", _e5rtTensorUtilsGetU8ElementErr)
	}
	return _e5rtTensorUtilsGetU8Element(a0, a1, a2, a3, a4, out), nil
}

// E5rtTensorUtilsGetU8Element.
func E5rtTensorUtilsGetU8Element(a0 uintptr, a1 uintptr, a2 *uint64, a3 uint64, a4 uint8, out *uint8) (int32, error) {
	return tryE5rtTensorUtilsGetU8Element(a0, a1, a2, a3, a4, out)
}

var _e5rtTensorUtilsSetFp32Element func(a0 uintptr, a1 uintptr, a2 *uint64, a3 uint64, a4 uint8, a5 float32) int32
var _e5rtTensorUtilsSetFp32ElementErr error

func tryE5rtTensorUtilsSetFp32Element(a0 uintptr, a1 uintptr, a2 *uint64, a3 uint64, a4 uint8, a5 float32) (int32, error) {
	if _e5rtTensorUtilsSetFp32Element == nil {
		return 0, symbolCallError("e5rt_tensor_utils_set_fp32_element", "", _e5rtTensorUtilsSetFp32ElementErr)
	}
	return _e5rtTensorUtilsSetFp32Element(a0, a1, a2, a3, a4, a5), nil
}

// E5rtTensorUtilsSetFp32Element.
func E5rtTensorUtilsSetFp32Element(a0 uintptr, a1 uintptr, a2 *uint64, a3 uint64, a4 uint8, a5 float32) (int32, error) {
	return tryE5rtTensorUtilsSetFp32Element(a0, a1, a2, a3, a4, a5)
}

var _e5rtTensorUtilsSetS8Element func(a0 uintptr, a1 uintptr, a2 *uint64, a3 uint64, a4 uint8, a5 int8) int32
var _e5rtTensorUtilsSetS8ElementErr error

func tryE5rtTensorUtilsSetS8Element(a0 uintptr, a1 uintptr, a2 *uint64, a3 uint64, a4 uint8, a5 int8) (int32, error) {
	if _e5rtTensorUtilsSetS8Element == nil {
		return 0, symbolCallError("e5rt_tensor_utils_set_s8_element", "", _e5rtTensorUtilsSetS8ElementErr)
	}
	return _e5rtTensorUtilsSetS8Element(a0, a1, a2, a3, a4, a5), nil
}

// E5rtTensorUtilsSetS8Element.
func E5rtTensorUtilsSetS8Element(a0 uintptr, a1 uintptr, a2 *uint64, a3 uint64, a4 uint8, a5 int8) (int32, error) {
	return tryE5rtTensorUtilsSetS8Element(a0, a1, a2, a3, a4, a5)
}

var _e5rtTensorUtilsSetU8Element func(a0 uintptr, a1 uintptr, a2 *uint64, a3 uint64, a4 uint8, a5 uint8) int32
var _e5rtTensorUtilsSetU8ElementErr error

func tryE5rtTensorUtilsSetU8Element(a0 uintptr, a1 uintptr, a2 *uint64, a3 uint64, a4 uint8, a5 uint8) (int32, error) {
	if _e5rtTensorUtilsSetU8Element == nil {
		return 0, symbolCallError("e5rt_tensor_utils_set_u8_element", "", _e5rtTensorUtilsSetU8ElementErr)
	}
	return _e5rtTensorUtilsSetU8Element(a0, a1, a2, a3, a4, a5), nil
}

// E5rtTensorUtilsSetU8Element.
func E5rtTensorUtilsSetU8Element(a0 uintptr, a1 uintptr, a2 *uint64, a3 uint64, a4 uint8, a5 uint8) (int32, error) {
	return tryE5rtTensorUtilsSetU8Element(a0, a1, a2, a3, a4, a5)
}

var _espressoAneCacheHasNetwork func(a0 *byte, out *bool) int32
var _espressoAneCacheHasNetworkErr error

func tryEspressoAneCacheHasNetwork(a0 *byte, out *bool) (int32, error) {
	if _espressoAneCacheHasNetwork == nil {
		return 0, symbolCallError("espresso_ane_cache_has_network", "", _espressoAneCacheHasNetworkErr)
	}
	return _espressoAneCacheHasNetwork(a0, out), nil
}

// EspressoAneCacheHasNetwork.
func EspressoAneCacheHasNetwork(a0 *byte, out *bool) (int32, error) {
	return tryEspressoAneCacheHasNetwork(a0, out)
}

var _espressoAneCachePurgeNetwork func(a0 *byte) int32
var _espressoAneCachePurgeNetworkErr error

func tryEspressoAneCachePurgeNetwork(a0 *byte) (int32, error) {
	if _espressoAneCachePurgeNetwork == nil {
		return 0, symbolCallError("espresso_ane_cache_purge_network", "", _espressoAneCachePurgeNetworkErr)
	}
	return _espressoAneCachePurgeNetwork(a0), nil
}

// EspressoAneCachePurgeNetwork.
func EspressoAneCachePurgeNetwork(a0 *byte) (int32, error) {
	return tryEspressoAneCachePurgeNetwork(a0)
}

var _espressoBlobSetIntOption func(a0 EspressoNetworkCStruct, a1 uintptr, a2 *byte, a3 uintptr) uintptr
var _espressoBlobSetIntOptionErr error

func tryEspressoBlobSetIntOption(a0 EspressoNetworkCStruct, a1 uintptr, a2 *byte, a3 uintptr) (uintptr, error) {
	if _espressoBlobSetIntOption == nil {
		return 0, symbolCallError("espresso_blob_set_int_option", "", _espressoBlobSetIntOptionErr)
	}
	return _espressoBlobSetIntOption(a0, a1, a2, a3), nil
}

// EspressoBlobSetIntOption.
func EspressoBlobSetIntOption(a0 EspressoNetworkCStruct, a1 uintptr, a2 *byte, a3 uintptr) (uintptr, error) {
	return tryEspressoBlobSetIntOption(a0, a1, a2, a3)
}

var _espressoBufferGetCount func(a0 uintptr) uintptr
var _espressoBufferGetCountErr error

func tryEspressoBufferGetCount(a0 uintptr) (uintptr, error) {
	if _espressoBufferGetCount == nil {
		return 0, symbolCallError("espresso_buffer_get_count", "", _espressoBufferGetCountErr)
	}
	return _espressoBufferGetCount(a0), nil
}

// EspressoBufferGetCount.
func EspressoBufferGetCount(a0 uintptr) (uintptr, error) {
	return tryEspressoBufferGetCount(a0)
}

var _espressoBufferGetRank func(a0 uintptr) uintptr
var _espressoBufferGetRankErr error

func tryEspressoBufferGetRank(a0 uintptr) (uintptr, error) {
	if _espressoBufferGetRank == nil {
		return 0, symbolCallError("espresso_buffer_get_rank", "", _espressoBufferGetRankErr)
	}
	return _espressoBufferGetRank(a0), nil
}

// EspressoBufferGetRank.
func EspressoBufferGetRank(a0 uintptr) (uintptr, error) {
	return tryEspressoBufferGetRank(a0)
}

var _espressoBufferGetSize func(a0 uintptr) uintptr
var _espressoBufferGetSizeErr error

func tryEspressoBufferGetSize(a0 uintptr) (uintptr, error) {
	if _espressoBufferGetSize == nil {
		return 0, symbolCallError("espresso_buffer_get_size", "", _espressoBufferGetSizeErr)
	}
	return _espressoBufferGetSize(a0), nil
}

// EspressoBufferGetSize.
func EspressoBufferGetSize(a0 uintptr) (uintptr, error) {
	return tryEspressoBufferGetSize(a0)
}

var _espressoBufferPackTensorShape func(a0 uintptr, a1 uint64, a2 uintptr) uintptr
var _espressoBufferPackTensorShapeErr error

func tryEspressoBufferPackTensorShape(a0 uintptr, a1 uint64, a2 uintptr) (uintptr, error) {
	if _espressoBufferPackTensorShape == nil {
		return 0, symbolCallError("espresso_buffer_pack_tensor_shape", "", _espressoBufferPackTensorShapeErr)
	}
	return _espressoBufferPackTensorShape(a0, a1, a2), nil
}

// EspressoBufferPackTensorShape.
func EspressoBufferPackTensorShape(a0 uintptr, a1 uint64, a2 uintptr) (uintptr, error) {
	return tryEspressoBufferPackTensorShape(a0, a1, a2)
}

var _espressoBufferSetRank func(a0 uintptr, a1 uintptr) uintptr
var _espressoBufferSetRankErr error

func tryEspressoBufferSetRank(a0 uintptr, a1 uintptr) (uintptr, error) {
	if _espressoBufferSetRank == nil {
		return 0, symbolCallError("espresso_buffer_set_rank", "", _espressoBufferSetRankErr)
	}
	return _espressoBufferSetRank(a0, a1), nil
}

// EspressoBufferSetRank.
func EspressoBufferSetRank(a0 uintptr, a1 uintptr) (uintptr, error) {
	return tryEspressoBufferSetRank(a0, a1)
}

var _espressoBufferUnpackTensorShape func(a0 uintptr, a1 uintptr, a2 uintptr) uintptr
var _espressoBufferUnpackTensorShapeErr error

func tryEspressoBufferUnpackTensorShape(a0 uintptr, a1 uintptr, a2 uintptr) (uintptr, error) {
	if _espressoBufferUnpackTensorShape == nil {
		return 0, symbolCallError("espresso_buffer_unpack_tensor_shape", "", _espressoBufferUnpackTensorShapeErr)
	}
	return _espressoBufferUnpackTensorShape(a0, a1, a2), nil
}

// EspressoBufferUnpackTensorShape.
func EspressoBufferUnpackTensorShape(a0 uintptr, a1 uintptr, a2 uintptr) (uintptr, error) {
	return tryEspressoBufferUnpackTensorShape(a0, a1, a2)
}

var _espressoCompileMilToEir func(a0 uintptr, a1 uintptr) uintptr
var _espressoCompileMilToEirErr error

func tryEspressoCompileMilToEir(a0 uintptr, a1 uintptr) (uintptr, error) {
	if _espressoCompileMilToEir == nil {
		return 0, symbolCallError("espresso_compile_mil_to_eir", "", _espressoCompileMilToEirErr)
	}
	return _espressoCompileMilToEir(a0, a1), nil
}

// EspressoCompileMilToEir.
func EspressoCompileMilToEir(a0 uintptr, a1 uintptr) (uintptr, error) {
	return tryEspressoCompileMilToEir(a0, a1)
}

var _espressoContextDestroy func(ctx EspressoContextRef)
var _espressoContextDestroyErr error

func tryEspressoContextDestroy(ctx EspressoContextRef) error {
	if _espressoContextDestroy == nil {
		return symbolCallError("espresso_context_destroy", "", _espressoContextDestroyErr)
	}
	_espressoContextDestroy(ctx)
	return nil
}

// EspressoContextDestroy.
func EspressoContextDestroy(ctx EspressoContextRef) error {
	return tryEspressoContextDestroy(ctx)
}

var _espressoContextReportBench func(a0 uintptr, a1 uintptr)
var _espressoContextReportBenchErr error

func tryEspressoContextReportBench(a0 uintptr, a1 uintptr) error {
	if _espressoContextReportBench == nil {
		return symbolCallError("espresso_context_report_bench", "", _espressoContextReportBenchErr)
	}
	_espressoContextReportBench(a0, a1)
	return nil
}

// EspressoContextReportBench.
func EspressoContextReportBench(a0 uintptr, a1 uintptr) error {
	return tryEspressoContextReportBench(a0, a1)
}

var _espressoContextSetIntOption func(a0 uintptr, a1 *byte, a2 int32) uintptr
var _espressoContextSetIntOptionErr error

func tryEspressoContextSetIntOption(a0 uintptr, a1 *byte, a2 int32) (uintptr, error) {
	if _espressoContextSetIntOption == nil {
		return 0, symbolCallError("espresso_context_set_int_option", "", _espressoContextSetIntOptionErr)
	}
	return _espressoContextSetIntOption(a0, a1, a2), nil
}

// EspressoContextSetIntOption.
func EspressoContextSetIntOption(a0 uintptr, a1 *byte, a2 int32) (uintptr, error) {
	return tryEspressoContextSetIntOption(a0, a1, a2)
}

var _espressoContextSetLowPrecisionAccumulation func(a0 uintptr, a1 bool) uintptr
var _espressoContextSetLowPrecisionAccumulationErr error

func tryEspressoContextSetLowPrecisionAccumulation(a0 uintptr, a1 bool) (uintptr, error) {
	if _espressoContextSetLowPrecisionAccumulation == nil {
		return 0, symbolCallError("espresso_context_set_low_precision_accumulation", "", _espressoContextSetLowPrecisionAccumulationErr)
	}
	return _espressoContextSetLowPrecisionAccumulation(a0, a1), nil
}

// EspressoContextSetLowPrecisionAccumulation.
func EspressoContextSetLowPrecisionAccumulation(a0 uintptr, a1 bool) (uintptr, error) {
	return tryEspressoContextSetLowPrecisionAccumulation(a0, a1)
}

var _espressoCreateContext func(platform int32, options int32) EspressoContextRef
var _espressoCreateContextErr error

func tryEspressoCreateContext(platform int32, options int32) (EspressoContextRef, error) {
	if _espressoCreateContext == nil {
		return *new(EspressoContextRef), symbolCallError("espresso_create_context", "", _espressoCreateContextErr)
	}
	return _espressoCreateContext(platform, options), nil
}

// EspressoCreateContext.
func EspressoCreateContext(platform int32, options int32) (EspressoContextRef, error) {
	return tryEspressoCreateContext(platform, options)
}

var _espressoCreateContextAuto func() uintptr
var _espressoCreateContextAutoErr error

func tryEspressoCreateContextAuto() (uintptr, error) {
	if _espressoCreateContextAuto == nil {
		return 0, symbolCallError("espresso_create_context_auto", "", _espressoCreateContextAutoErr)
	}
	return _espressoCreateContextAuto(), nil
}

// EspressoCreateContextAuto.
func EspressoCreateContextAuto() (uintptr, error) {
	return tryEspressoCreateContextAuto()
}

var _espressoCreateContextWithArgs func(a0 uintptr, a1 uintptr, a2 uintptr) uintptr
var _espressoCreateContextWithArgsErr error

func tryEspressoCreateContextWithArgs(a0 uintptr, a1 uintptr, a2 uintptr) (uintptr, error) {
	if _espressoCreateContextWithArgs == nil {
		return 0, symbolCallError("espresso_create_context_with_args", "", _espressoCreateContextWithArgsErr)
	}
	return _espressoCreateContextWithArgs(a0, a1, a2), nil
}

// EspressoCreateContextWithArgs.
func EspressoCreateContextWithArgs(a0 uintptr, a1 uintptr, a2 uintptr) (uintptr, error) {
	return tryEspressoCreateContextWithArgs(a0, a1, a2)
}

var _espressoCreatePlan func(ctx EspressoContextRef, platform int32) EspressoPlanRef
var _espressoCreatePlanErr error

func tryEspressoCreatePlan(ctx EspressoContextRef, platform int32) (EspressoPlanRef, error) {
	if _espressoCreatePlan == nil {
		return *new(EspressoPlanRef), symbolCallError("espresso_create_plan", "", _espressoCreatePlanErr)
	}
	return _espressoCreatePlan(ctx, platform), nil
}

// EspressoCreatePlan.
func EspressoCreatePlan(ctx EspressoContextRef, platform int32) (EspressoPlanRef, error) {
	return tryEspressoCreatePlan(ctx, platform)
}

var _espressoCreatePlanAndLoadNetwork func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) uintptr
var _espressoCreatePlanAndLoadNetworkErr error

func tryEspressoCreatePlanAndLoadNetwork(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (uintptr, error) {
	if _espressoCreatePlanAndLoadNetwork == nil {
		return 0, symbolCallError("espresso_create_plan_and_load_network", "", _espressoCreatePlanAndLoadNetworkErr)
	}
	return _espressoCreatePlanAndLoadNetwork(a0, a1, a2, a3), nil
}

// EspressoCreatePlanAndLoadNetwork.
func EspressoCreatePlanAndLoadNetwork(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (uintptr, error) {
	return tryEspressoCreatePlanAndLoadNetwork(a0, a1, a2, a3)
}

var _espressoDeviceIDForMetalDevice func(a0 uintptr) uintptr
var _espressoDeviceIDForMetalDeviceErr error

func tryEspressoDeviceIDForMetalDevice(a0 uintptr) (uintptr, error) {
	if _espressoDeviceIDForMetalDevice == nil {
		return 0, symbolCallError("espresso_device_id_for_metal_device", "", _espressoDeviceIDForMetalDeviceErr)
	}
	return _espressoDeviceIDForMetalDevice(a0), nil
}

// EspressoDeviceIDForMetalDevice.
func EspressoDeviceIDForMetalDevice(a0 uintptr) (uintptr, error) {
	return tryEspressoDeviceIDForMetalDevice(a0)
}

var _espressoDumpIr func(a0 uintptr, out **byte) uintptr
var _espressoDumpIrErr error

func tryEspressoDumpIr(a0 uintptr, out **byte) (uintptr, error) {
	if _espressoDumpIr == nil {
		return 0, symbolCallError("espresso_dump_ir", "", _espressoDumpIrErr)
	}
	return _espressoDumpIr(a0, out), nil
}

// EspressoDumpIr.
func EspressoDumpIr(a0 uintptr, out **byte) (uintptr, error) {
	return tryEspressoDumpIr(a0, out)
}

var _espressoEnableAutoinitialize func(a0 uintptr) uintptr
var _espressoEnableAutoinitializeErr error

func tryEspressoEnableAutoinitialize(a0 uintptr) (uintptr, error) {
	if _espressoEnableAutoinitialize == nil {
		return 0, symbolCallError("espresso_enable_autoinitialize", "", _espressoEnableAutoinitializeErr)
	}
	return _espressoEnableAutoinitialize(a0), nil
}

// EspressoEnableAutoinitialize.
func EspressoEnableAutoinitialize(a0 uintptr) (uintptr, error) {
	return tryEspressoEnableAutoinitialize(a0)
}

var _espressoEnableTestVectorMode func(a0 uintptr, a1 uintptr) uintptr
var _espressoEnableTestVectorModeErr error

func tryEspressoEnableTestVectorMode(a0 uintptr, a1 uintptr) (uintptr, error) {
	if _espressoEnableTestVectorMode == nil {
		return 0, symbolCallError("espresso_enable_test_vector_mode", "", _espressoEnableTestVectorModeErr)
	}
	return _espressoEnableTestVectorMode(a0, a1), nil
}

// EspressoEnableTestVectorMode.
func EspressoEnableTestVectorMode(a0 uintptr, a1 uintptr) (uintptr, error) {
	return tryEspressoEnableTestVectorMode(a0, a1)
}

var _espressoGenerateTrainingProgram func(a0 uintptr, a1 uintptr) uintptr
var _espressoGenerateTrainingProgramErr error

func tryEspressoGenerateTrainingProgram(a0 uintptr, a1 uintptr) (uintptr, error) {
	if _espressoGenerateTrainingProgram == nil {
		return 0, symbolCallError("espresso_generate_training_program", "", _espressoGenerateTrainingProgramErr)
	}
	return _espressoGenerateTrainingProgram(a0, a1), nil
}

// EspressoGenerateTrainingProgram.
func EspressoGenerateTrainingProgram(a0 uintptr, a1 uintptr) (uintptr, error) {
	return tryEspressoGenerateTrainingProgram(a0, a1)
}

var _espressoGetAnalysisModelMetadataForKey func(a0 EspressoNetworkCStruct, a1 *byte) uintptr
var _espressoGetAnalysisModelMetadataForKeyErr error

func tryEspressoGetAnalysisModelMetadataForKey(a0 EspressoNetworkCStruct, a1 *byte) (uintptr, error) {
	if _espressoGetAnalysisModelMetadataForKey == nil {
		return 0, symbolCallError("espresso_get_analysis_model_metadata_for_key", "", _espressoGetAnalysisModelMetadataForKeyErr)
	}
	return _espressoGetAnalysisModelMetadataForKey(a0, a1), nil
}

// EspressoGetAnalysisModelMetadataForKey.
func EspressoGetAnalysisModelMetadataForKey(a0 EspressoNetworkCStruct, a1 *byte) (uintptr, error) {
	return tryEspressoGetAnalysisModelMetadataForKey(a0, a1)
}

var _espressoGetDefaultStorageType func(a0 uintptr) uint64
var _espressoGetDefaultStorageTypeErr error

func tryEspressoGetDefaultStorageType(a0 uintptr) (uint64, error) {
	if _espressoGetDefaultStorageType == nil {
		return 0, symbolCallError("espresso_get_default_storage_type", "", _espressoGetDefaultStorageTypeErr)
	}
	return _espressoGetDefaultStorageType(a0), nil
}

// EspressoGetDefaultStorageType.
func EspressoGetDefaultStorageType(a0 uintptr) (uint64, error) {
	return tryEspressoGetDefaultStorageType(a0)
}

var _espressoGetMetadataForKey func(a0 EspressoNetworkCStruct, a1 uintptr, a2 uintptr, a3 uintptr) uintptr
var _espressoGetMetadataForKeyErr error

func tryEspressoGetMetadataForKey(a0 EspressoNetworkCStruct, a1 uintptr, a2 uintptr, a3 uintptr) (uintptr, error) {
	if _espressoGetMetadataForKey == nil {
		return 0, symbolCallError("espresso_get_metadata_for_key", "", _espressoGetMetadataForKeyErr)
	}
	return _espressoGetMetadataForKey(a0, a1, a2, a3), nil
}

// EspressoGetMetadataForKey.
func EspressoGetMetadataForKey(a0 EspressoNetworkCStruct, a1 uintptr, a2 uintptr, a3 uintptr) (uintptr, error) {
	return tryEspressoGetMetadataForKey(a0, a1, a2, a3)
}

var _espressoGetStatusString func(a0 uintptr) uintptr
var _espressoGetStatusStringErr error

func tryEspressoGetStatusString(a0 uintptr) (uintptr, error) {
	if _espressoGetStatusString == nil {
		return 0, symbolCallError("espresso_get_status_string", "", _espressoGetStatusStringErr)
	}
	return _espressoGetStatusString(a0), nil
}

// EspressoGetStatusString.
func EspressoGetStatusString(a0 uintptr) (uintptr, error) {
	return tryEspressoGetStatusString(a0)
}

var _espressoGetVersionString func() *byte
var _espressoGetVersionStringErr error

func tryEspressoGetVersionString() (*byte, error) {
	if _espressoGetVersionString == nil {
		return nil, symbolCallError("espresso_get_version_string", "", _espressoGetVersionStringErr)
	}
	return _espressoGetVersionString(), nil
}

// EspressoGetVersionString.
func EspressoGetVersionString() (*byte, error) {
	return tryEspressoGetVersionString()
}

var _espressoGPUPreferIntegrated func(a0 uintptr) uintptr
var _espressoGPUPreferIntegratedErr error

func tryEspressoGPUPreferIntegrated(a0 uintptr) (uintptr, error) {
	if _espressoGPUPreferIntegrated == nil {
		return 0, symbolCallError("espresso_gpu_prefer_integrated", "", _espressoGPUPreferIntegratedErr)
	}
	return _espressoGPUPreferIntegrated(a0), nil
}

// EspressoGPUPreferIntegrated.
func EspressoGPUPreferIntegrated(a0 uintptr) (uintptr, error) {
	return tryEspressoGPUPreferIntegrated(a0)
}

var _espressoIsAneArchGreaterThanOrEqual func(a0 *byte, a1 uintptr) uintptr
var _espressoIsAneArchGreaterThanOrEqualErr error

func tryEspressoIsAneArchGreaterThanOrEqual(a0 *byte, a1 uintptr) (uintptr, error) {
	if _espressoIsAneArchGreaterThanOrEqual == nil {
		return 0, symbolCallError("espresso_is_ane_arch_greater_than_or_equal", "", _espressoIsAneArchGreaterThanOrEqualErr)
	}
	return _espressoIsAneArchGreaterThanOrEqual(a0, a1), nil
}

// EspressoIsAneArchGreaterThanOrEqual.
func EspressoIsAneArchGreaterThanOrEqual(a0 *byte, a1 uintptr) (uintptr, error) {
	return tryEspressoIsAneArchGreaterThanOrEqual(a0, a1)
}

var _espressoNetworkBindBuffer func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) int32
var _espressoNetworkBindBufferErr error

func tryEspressoNetworkBindBuffer(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) (int32, error) {
	if _espressoNetworkBindBuffer == nil {
		return 0, symbolCallError("espresso_network_bind_buffer", "", _espressoNetworkBindBufferErr)
	}
	return _espressoNetworkBindBuffer(a0, a1, a2, a3, a4, a5, a6), nil
}

// EspressoNetworkBindBuffer.
func EspressoNetworkBindBuffer(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) (int32, error) {
	return tryEspressoNetworkBindBuffer(a0, a1, a2, a3, a4, a5, a6)
}

var _espressoNetworkBindCvpixelbuffer func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) int32
var _espressoNetworkBindCvpixelbufferErr error

func tryEspressoNetworkBindCvpixelbuffer(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	if _espressoNetworkBindCvpixelbuffer == nil {
		return 0, symbolCallError("espresso_network_bind_cvpixelbuffer", "", _espressoNetworkBindCvpixelbufferErr)
	}
	return _espressoNetworkBindCvpixelbuffer(a0, a1, a2, a3, a4), nil
}

// EspressoNetworkBindCvpixelbuffer.
func EspressoNetworkBindCvpixelbuffer(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	return tryEspressoNetworkBindCvpixelbuffer(a0, a1, a2, a3, a4)
}

var _espressoNetworkBindDirectCvpixelbuffer func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _espressoNetworkBindDirectCvpixelbufferErr error

func tryEspressoNetworkBindDirectCvpixelbuffer(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _espressoNetworkBindDirectCvpixelbuffer == nil {
		return 0, symbolCallError("espresso_network_bind_direct_cvpixelbuffer", "", _espressoNetworkBindDirectCvpixelbufferErr)
	}
	return _espressoNetworkBindDirectCvpixelbuffer(a0, a1, a2, a3), nil
}

// EspressoNetworkBindDirectCvpixelbuffer.
func EspressoNetworkBindDirectCvpixelbuffer(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryEspressoNetworkBindDirectCvpixelbuffer(a0, a1, a2, a3)
}

var _espressoNetworkBindInputCvpixelbuffer func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) int32
var _espressoNetworkBindInputCvpixelbufferErr error

func tryEspressoNetworkBindInputCvpixelbuffer(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) (int32, error) {
	if _espressoNetworkBindInputCvpixelbuffer == nil {
		return 0, symbolCallError("espresso_network_bind_input_cvpixelbuffer", "", _espressoNetworkBindInputCvpixelbufferErr)
	}
	return _espressoNetworkBindInputCvpixelbuffer(a0, a1, a2, a3, a4, a5, a6), nil
}

// EspressoNetworkBindInputCvpixelbuffer.
func EspressoNetworkBindInputCvpixelbuffer(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) (int32, error) {
	return tryEspressoNetworkBindInputCvpixelbuffer(a0, a1, a2, a3, a4, a5, a6)
}

var _espressoNetworkBindInputMetaltexture func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) int32
var _espressoNetworkBindInputMetaltextureErr error

func tryEspressoNetworkBindInputMetaltexture(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) (int32, error) {
	if _espressoNetworkBindInputMetaltexture == nil {
		return 0, symbolCallError("espresso_network_bind_input_metaltexture", "", _espressoNetworkBindInputMetaltextureErr)
	}
	return _espressoNetworkBindInputMetaltexture(a0, a1, a2, a3, a4, a5, a6), nil
}

// EspressoNetworkBindInputMetaltexture.
func EspressoNetworkBindInputMetaltexture(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) (int32, error) {
	return tryEspressoNetworkBindInputMetaltexture(a0, a1, a2, a3, a4, a5, a6)
}

var _espressoNetworkBindInputVimagebufferArgb8 func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, out *uintptr, a5 uintptr) int32
var _espressoNetworkBindInputVimagebufferArgb8Err error

func tryEspressoNetworkBindInputVimagebufferArgb8(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, out *uintptr, a5 uintptr) (int32, error) {
	if _espressoNetworkBindInputVimagebufferArgb8 == nil {
		return 0, symbolCallError("espresso_network_bind_input_vimagebuffer_argb8", "", _espressoNetworkBindInputVimagebufferArgb8Err)
	}
	return _espressoNetworkBindInputVimagebufferArgb8(a0, a1, a2, a3, out, a5), nil
}

// EspressoNetworkBindInputVimagebufferArgb8.
func EspressoNetworkBindInputVimagebufferArgb8(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, out *uintptr, a5 uintptr) (int32, error) {
	return tryEspressoNetworkBindInputVimagebufferArgb8(a0, a1, a2, a3, out, a5)
}

var _espressoNetworkBindInputVimagebufferBgra8 func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) int32
var _espressoNetworkBindInputVimagebufferBgra8Err error

func tryEspressoNetworkBindInputVimagebufferBgra8(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	if _espressoNetworkBindInputVimagebufferBgra8 == nil {
		return 0, symbolCallError("espresso_network_bind_input_vimagebuffer_bgra8", "", _espressoNetworkBindInputVimagebufferBgra8Err)
	}
	return _espressoNetworkBindInputVimagebufferBgra8(a0, a1, a2, a3, a4, a5), nil
}

// EspressoNetworkBindInputVimagebufferBgra8.
func EspressoNetworkBindInputVimagebufferBgra8(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	return tryEspressoNetworkBindInputVimagebufferBgra8(a0, a1, a2, a3, a4, a5)
}

var _espressoNetworkBindInputVimagebufferPlanar8 func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, out *uintptr, a5 uintptr) int32
var _espressoNetworkBindInputVimagebufferPlanar8Err error

func tryEspressoNetworkBindInputVimagebufferPlanar8(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, out *uintptr, a5 uintptr) (int32, error) {
	if _espressoNetworkBindInputVimagebufferPlanar8 == nil {
		return 0, symbolCallError("espresso_network_bind_input_vimagebuffer_planar8", "", _espressoNetworkBindInputVimagebufferPlanar8Err)
	}
	return _espressoNetworkBindInputVimagebufferPlanar8(a0, a1, a2, a3, out, a5), nil
}

// EspressoNetworkBindInputVimagebufferPlanar8.
func EspressoNetworkBindInputVimagebufferPlanar8(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, out *uintptr, a5 uintptr) (int32, error) {
	return tryEspressoNetworkBindInputVimagebufferPlanar8(a0, a1, a2, a3, out, a5)
}

var _espressoNetworkBindInputVimagebufferRgba8 func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, out *uintptr, a5 uintptr) int32
var _espressoNetworkBindInputVimagebufferRgba8Err error

func tryEspressoNetworkBindInputVimagebufferRgba8(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, out *uintptr, a5 uintptr) (int32, error) {
	if _espressoNetworkBindInputVimagebufferRgba8 == nil {
		return 0, symbolCallError("espresso_network_bind_input_vimagebuffer_rgba8", "", _espressoNetworkBindInputVimagebufferRgba8Err)
	}
	return _espressoNetworkBindInputVimagebufferRgba8(a0, a1, a2, a3, out, a5), nil
}

// EspressoNetworkBindInputVimagebufferRgba8.
func EspressoNetworkBindInputVimagebufferRgba8(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, out *uintptr, a5 uintptr) (int32, error) {
	return tryEspressoNetworkBindInputVimagebufferRgba8(a0, a1, a2, a3, out, a5)
}

var _espressoNetworkChangeBlobShape func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr, a7 uintptr) uintptr
var _espressoNetworkChangeBlobShapeErr error

func tryEspressoNetworkChangeBlobShape(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr, a7 uintptr) (uintptr, error) {
	if _espressoNetworkChangeBlobShape == nil {
		return 0, symbolCallError("espresso_network_change_blob_shape", "", _espressoNetworkChangeBlobShapeErr)
	}
	return _espressoNetworkChangeBlobShape(a0, a1, a2, a3, a4, a5, a6, a7), nil
}

// EspressoNetworkChangeBlobShape.
func EspressoNetworkChangeBlobShape(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr, a7 uintptr) (uintptr, error) {
	return tryEspressoNetworkChangeBlobShape(a0, a1, a2, a3, a4, a5, a6, a7)
}

var _espressoNetworkChangeInputBlobShapes func(a0 EspressoNetworkCStruct, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) uintptr
var _espressoNetworkChangeInputBlobShapesErr error

func tryEspressoNetworkChangeInputBlobShapes(a0 EspressoNetworkCStruct, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) (uintptr, error) {
	if _espressoNetworkChangeInputBlobShapes == nil {
		return 0, symbolCallError("espresso_network_change_input_blob_shapes", "", _espressoNetworkChangeInputBlobShapesErr)
	}
	return _espressoNetworkChangeInputBlobShapes(a0, a1, a2, a3, a4, a5, a6), nil
}

// EspressoNetworkChangeInputBlobShapes.
func EspressoNetworkChangeInputBlobShapes(a0 EspressoNetworkCStruct, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) (uintptr, error) {
	return tryEspressoNetworkChangeInputBlobShapes(a0, a1, a2, a3, a4, a5, a6)
}

var _espressoNetworkDeclareInput func(a0 uintptr, a1 int32, a2 *byte) int32
var _espressoNetworkDeclareInputErr error

func tryEspressoNetworkDeclareInput(a0 uintptr, a1 int32, a2 *byte) (int32, error) {
	if _espressoNetworkDeclareInput == nil {
		return 0, symbolCallError("espresso_network_declare_input", "", _espressoNetworkDeclareInputErr)
	}
	return _espressoNetworkDeclareInput(a0, a1, a2), nil
}

// EspressoNetworkDeclareInput.
func EspressoNetworkDeclareInput(a0 uintptr, a1 int32, a2 *byte) (int32, error) {
	return tryEspressoNetworkDeclareInput(a0, a1, a2)
}

var _espressoNetworkDeclareOutput func(a0 uintptr, a1 int32, a2 *byte) int32
var _espressoNetworkDeclareOutputErr error

func tryEspressoNetworkDeclareOutput(a0 uintptr, a1 int32, a2 *byte) (int32, error) {
	if _espressoNetworkDeclareOutput == nil {
		return 0, symbolCallError("espresso_network_declare_output", "", _espressoNetworkDeclareOutputErr)
	}
	return _espressoNetworkDeclareOutput(a0, a1, a2), nil
}

// EspressoNetworkDeclareOutput.
func EspressoNetworkDeclareOutput(a0 uintptr, a1 int32, a2 *byte) (int32, error) {
	return tryEspressoNetworkDeclareOutput(a0, a1, a2)
}

var _espressoNetworkDumpTestVector func(a0 EspressoNetworkCStruct, a1 uintptr) int32
var _espressoNetworkDumpTestVectorErr error

func tryEspressoNetworkDumpTestVector(a0 EspressoNetworkCStruct, a1 uintptr) (int32, error) {
	if _espressoNetworkDumpTestVector == nil {
		return 0, symbolCallError("espresso_network_dump_test_vector", "", _espressoNetworkDumpTestVectorErr)
	}
	return _espressoNetworkDumpTestVector(a0, a1), nil
}

// EspressoNetworkDumpTestVector.
func EspressoNetworkDumpTestVector(a0 EspressoNetworkCStruct, a1 uintptr) (int32, error) {
	return tryEspressoNetworkDumpTestVector(a0, a1)
}

var _espressoNetworkGetVersion func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) uintptr
var _espressoNetworkGetVersionErr error

func tryEspressoNetworkGetVersion(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (uintptr, error) {
	if _espressoNetworkGetVersion == nil {
		return 0, symbolCallError("espresso_network_get_version", "", _espressoNetworkGetVersionErr)
	}
	return _espressoNetworkGetVersion(a0, a1, a2, a3), nil
}

// EspressoNetworkGetVersion.
func EspressoNetworkGetVersion(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (uintptr, error) {
	return tryEspressoNetworkGetVersion(a0, a1, a2, a3)
}

var _espressoNetworkQueryBlobDimensions func(a0 uintptr, a1 uintptr, a2 uintptr, out *uintptr) int32
var _espressoNetworkQueryBlobDimensionsErr error

func tryEspressoNetworkQueryBlobDimensions(a0 uintptr, a1 uintptr, a2 uintptr, out *uintptr) (int32, error) {
	if _espressoNetworkQueryBlobDimensions == nil {
		return 0, symbolCallError("espresso_network_query_blob_dimensions", "", _espressoNetworkQueryBlobDimensionsErr)
	}
	return _espressoNetworkQueryBlobDimensions(a0, a1, a2, out), nil
}

// EspressoNetworkQueryBlobDimensions.
func EspressoNetworkQueryBlobDimensions(a0 uintptr, a1 uintptr, a2 uintptr, out *uintptr) (int32, error) {
	return tryEspressoNetworkQueryBlobDimensions(a0, a1, a2, out)
}

var _espressoNetworkQueryBlobShape func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) int32
var _espressoNetworkQueryBlobShapeErr error

func tryEspressoNetworkQueryBlobShape(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	if _espressoNetworkQueryBlobShape == nil {
		return 0, symbolCallError("espresso_network_query_blob_shape", "", _espressoNetworkQueryBlobShapeErr)
	}
	return _espressoNetworkQueryBlobShape(a0, a1, a2, a3, a4), nil
}

// EspressoNetworkQueryBlobShape.
func EspressoNetworkQueryBlobShape(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	return tryEspressoNetworkQueryBlobShape(a0, a1, a2, a3, a4)
}

var _espressoNetworkSetMemoryPoolID func(a0 EspressoNetworkCStruct, a1 uintptr) uintptr
var _espressoNetworkSetMemoryPoolIDErr error

func tryEspressoNetworkSetMemoryPoolID(a0 EspressoNetworkCStruct, a1 uintptr) (uintptr, error) {
	if _espressoNetworkSetMemoryPoolID == nil {
		return 0, symbolCallError("espresso_network_set_memory_pool_id", "", _espressoNetworkSetMemoryPoolIDErr)
	}
	return _espressoNetworkSetMemoryPoolID(a0, a1), nil
}

// EspressoNetworkSetMemoryPoolID.
func EspressoNetworkSetMemoryPoolID(a0 EspressoNetworkCStruct, a1 uintptr) (uintptr, error) {
	return tryEspressoNetworkSetMemoryPoolID(a0, a1)
}

var _espressoNetworkSetTracingName func(a0 EspressoNetworkCStruct, a1 uintptr) int32
var _espressoNetworkSetTracingNameErr error

func tryEspressoNetworkSetTracingName(a0 EspressoNetworkCStruct, a1 uintptr) (int32, error) {
	if _espressoNetworkSetTracingName == nil {
		return 0, symbolCallError("espresso_network_set_tracing_name", "", _espressoNetworkSetTracingNameErr)
	}
	return _espressoNetworkSetTracingName(a0, a1), nil
}

// EspressoNetworkSetTracingName.
func EspressoNetworkSetTracingName(a0 EspressoNetworkCStruct, a1 uintptr) (int32, error) {
	return tryEspressoNetworkSetTracingName(a0, a1)
}

var _espressoNetworkSwapGlobal func(a0 EspressoNetworkCStruct, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _espressoNetworkSwapGlobalErr error

func tryEspressoNetworkSwapGlobal(a0 EspressoNetworkCStruct, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _espressoNetworkSwapGlobal == nil {
		return 0, symbolCallError("espresso_network_swap_global", "", _espressoNetworkSwapGlobalErr)
	}
	return _espressoNetworkSwapGlobal(a0, a1, a2, a3), nil
}

// EspressoNetworkSwapGlobal.
func EspressoNetworkSwapGlobal(a0 EspressoNetworkCStruct, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryEspressoNetworkSwapGlobal(a0, a1, a2, a3)
}

var _espressoNetworkSyncCopyGlobal func(a0 EspressoNetworkCStruct, a1 uintptr, a2 uintptr) int32
var _espressoNetworkSyncCopyGlobalErr error

func tryEspressoNetworkSyncCopyGlobal(a0 EspressoNetworkCStruct, a1 uintptr, a2 uintptr) (int32, error) {
	if _espressoNetworkSyncCopyGlobal == nil {
		return 0, symbolCallError("espresso_network_sync_copy_global", "", _espressoNetworkSyncCopyGlobalErr)
	}
	return _espressoNetworkSyncCopyGlobal(a0, a1, a2), nil
}

// EspressoNetworkSyncCopyGlobal.
func EspressoNetworkSyncCopyGlobal(a0 EspressoNetworkCStruct, a1 uintptr, a2 uintptr) (int32, error) {
	return tryEspressoNetworkSyncCopyGlobal(a0, a1, a2)
}

var _espressoNetworkTemporalStateReset func(a0 uintptr, a1 uintptr, a2 uintptr) uintptr
var _espressoNetworkTemporalStateResetErr error

func tryEspressoNetworkTemporalStateReset(a0 uintptr, a1 uintptr, a2 uintptr) (uintptr, error) {
	if _espressoNetworkTemporalStateReset == nil {
		return 0, symbolCallError("espresso_network_temporal_state_reset", "", _espressoNetworkTemporalStateResetErr)
	}
	return _espressoNetworkTemporalStateReset(a0, a1, a2), nil
}

// EspressoNetworkTemporalStateReset.
func EspressoNetworkTemporalStateReset(a0 uintptr, a1 uintptr, a2 uintptr) (uintptr, error) {
	return tryEspressoNetworkTemporalStateReset(a0, a1, a2)
}

var _espressoNetworkUnbindBuffer func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _espressoNetworkUnbindBufferErr error

func tryEspressoNetworkUnbindBuffer(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _espressoNetworkUnbindBuffer == nil {
		return 0, symbolCallError("espresso_network_unbind_buffer", "", _espressoNetworkUnbindBufferErr)
	}
	return _espressoNetworkUnbindBuffer(a0, a1, a2, a3), nil
}

// EspressoNetworkUnbindBuffer.
func EspressoNetworkUnbindBuffer(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryEspressoNetworkUnbindBuffer(a0, a1, a2, a3)
}

var _espressoPlanAddNetwork func(a0 uintptr, a1 *byte, a2 uint64, out *uint64) int32
var _espressoPlanAddNetworkErr error

func tryEspressoPlanAddNetwork(a0 uintptr, a1 *byte, a2 uint64, out *uint64) (int32, error) {
	if _espressoPlanAddNetwork == nil {
		return 0, symbolCallError("espresso_plan_add_network", "", _espressoPlanAddNetworkErr)
	}
	return _espressoPlanAddNetwork(a0, a1, a2, out), nil
}

// EspressoPlanAddNetwork.
func EspressoPlanAddNetwork(a0 uintptr, a1 *byte, a2 uint64, out *uint64) (int32, error) {
	return tryEspressoPlanAddNetwork(a0, a1, a2, out)
}

var _espressoPlanAutoProfile func(a0 uintptr) uintptr
var _espressoPlanAutoProfileErr error

func tryEspressoPlanAutoProfile(a0 uintptr) (uintptr, error) {
	if _espressoPlanAutoProfile == nil {
		return 0, symbolCallError("espresso_plan_auto_profile", "", _espressoPlanAutoProfileErr)
	}
	return _espressoPlanAutoProfile(a0), nil
}

// EspressoPlanAutoProfile.
func EspressoPlanAutoProfile(a0 uintptr) (uintptr, error) {
	return tryEspressoPlanAutoProfile(a0)
}

var _espressoPlanBuild func(plan EspressoPlanRef) int32
var _espressoPlanBuildErr error

func tryEspressoPlanBuild(plan EspressoPlanRef) (int32, error) {
	if _espressoPlanBuild == nil {
		return 0, symbolCallError("espresso_plan_build", "", _espressoPlanBuildErr)
	}
	return _espressoPlanBuild(plan), nil
}

// EspressoPlanBuild.
func EspressoPlanBuild(plan EspressoPlanRef) (int32, error) {
	return tryEspressoPlanBuild(plan)
}

var _espressoPlanBuildClean func(a0 uintptr) int32
var _espressoPlanBuildCleanErr error

func tryEspressoPlanBuildClean(a0 uintptr) (int32, error) {
	if _espressoPlanBuildClean == nil {
		return 0, symbolCallError("espresso_plan_build_clean", "", _espressoPlanBuildCleanErr)
	}
	return _espressoPlanBuildClean(a0), nil
}

// EspressoPlanBuildClean.
func EspressoPlanBuildClean(a0 uintptr) (int32, error) {
	return tryEspressoPlanBuildClean(a0)
}

var _espressoPlanBuildWithOptions func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _espressoPlanBuildWithOptionsErr error

func tryEspressoPlanBuildWithOptions(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _espressoPlanBuildWithOptions == nil {
		return 0, symbolCallError("espresso_plan_build_with_options", "", _espressoPlanBuildWithOptionsErr)
	}
	return _espressoPlanBuildWithOptions(a0, a1, a2), nil
}

// EspressoPlanBuildWithOptions.
func EspressoPlanBuildWithOptions(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryEspressoPlanBuildWithOptions(a0, a1, a2)
}

var _espressoPlanCanUseSubmit func(a0 uintptr) int32
var _espressoPlanCanUseSubmitErr error

func tryEspressoPlanCanUseSubmit(a0 uintptr) (int32, error) {
	if _espressoPlanCanUseSubmit == nil {
		return 0, symbolCallError("espresso_plan_can_use_submit", "", _espressoPlanCanUseSubmitErr)
	}
	return _espressoPlanCanUseSubmit(a0), nil
}

// EspressoPlanCanUseSubmit.
func EspressoPlanCanUseSubmit(a0 uintptr) (int32, error) {
	return tryEspressoPlanCanUseSubmit(a0)
}

var _espressoPlanDestroy func(plan EspressoPlanRef)
var _espressoPlanDestroyErr error

func tryEspressoPlanDestroy(plan EspressoPlanRef) error {
	if _espressoPlanDestroy == nil {
		return symbolCallError("espresso_plan_destroy", "", _espressoPlanDestroyErr)
	}
	_espressoPlanDestroy(plan)
	return nil
}

// EspressoPlanDestroy.
func EspressoPlanDestroy(plan EspressoPlanRef) error {
	return tryEspressoPlanDestroy(plan)
}

var _espressoPlanExecuteSync func(plan EspressoPlanRef) int32
var _espressoPlanExecuteSyncErr error

func tryEspressoPlanExecuteSync(plan EspressoPlanRef) (int32, error) {
	if _espressoPlanExecuteSync == nil {
		return 0, symbolCallError("espresso_plan_execute_sync", "", _espressoPlanExecuteSyncErr)
	}
	return _espressoPlanExecuteSync(plan), nil
}

// EspressoPlanExecuteSync.
func EspressoPlanExecuteSync(plan EspressoPlanRef) (int32, error) {
	return tryEspressoPlanExecuteSync(plan)
}

var _espressoPlanFinishProfiling func(a0 uintptr) uintptr
var _espressoPlanFinishProfilingErr error

func tryEspressoPlanFinishProfiling(a0 uintptr) (uintptr, error) {
	if _espressoPlanFinishProfiling == nil {
		return 0, symbolCallError("espresso_plan_finish_profiling", "", _espressoPlanFinishProfilingErr)
	}
	return _espressoPlanFinishProfiling(a0), nil
}

// EspressoPlanFinishProfiling.
func EspressoPlanFinishProfiling(a0 uintptr) (uintptr, error) {
	return tryEspressoPlanFinishProfiling(a0)
}

var _espressoPlanGetErrorInfo func(a0 uintptr) uintptr
var _espressoPlanGetErrorInfoErr error

func tryEspressoPlanGetErrorInfo(a0 uintptr) (uintptr, error) {
	if _espressoPlanGetErrorInfo == nil {
		return 0, symbolCallError("espresso_plan_get_error_info", "", _espressoPlanGetErrorInfoErr)
	}
	return _espressoPlanGetErrorInfo(a0), nil
}

// EspressoPlanGetErrorInfo.
func EspressoPlanGetErrorInfo(a0 uintptr) (uintptr, error) {
	return tryEspressoPlanGetErrorInfo(a0)
}

var _espressoPlanGetPhase func(a0 uintptr) uintptr
var _espressoPlanGetPhaseErr error

func tryEspressoPlanGetPhase(a0 uintptr) (uintptr, error) {
	if _espressoPlanGetPhase == nil {
		return 0, symbolCallError("espresso_plan_get_phase", "", _espressoPlanGetPhaseErr)
	}
	return _espressoPlanGetPhase(a0), nil
}

// EspressoPlanGetPhase.
func EspressoPlanGetPhase(a0 uintptr) (uintptr, error) {
	return tryEspressoPlanGetPhase(a0)
}

var _espressoPlanSetExecutionQueue func(a0 uintptr, a1 uintptr) int32
var _espressoPlanSetExecutionQueueErr error

func tryEspressoPlanSetExecutionQueue(a0 uintptr, a1 uintptr) (int32, error) {
	if _espressoPlanSetExecutionQueue == nil {
		return 0, symbolCallError("espresso_plan_set_execution_queue", "", _espressoPlanSetExecutionQueueErr)
	}
	return _espressoPlanSetExecutionQueue(a0, a1), nil
}

// EspressoPlanSetExecutionQueue.
func EspressoPlanSetExecutionQueue(a0 uintptr, a1 uintptr) (int32, error) {
	return tryEspressoPlanSetExecutionQueue(a0, a1)
}

var _espressoPlanSetPriority func(a0 uintptr, a1 int32) int32
var _espressoPlanSetPriorityErr error

func tryEspressoPlanSetPriority(a0 uintptr, a1 int32) (int32, error) {
	if _espressoPlanSetPriority == nil {
		return 0, symbolCallError("espresso_plan_set_priority", "", _espressoPlanSetPriorityErr)
	}
	return _espressoPlanSetPriority(a0, a1), nil
}

// EspressoPlanSetPriority.
func EspressoPlanSetPriority(a0 uintptr, a1 int32) (int32, error) {
	return tryEspressoPlanSetPriority(a0, a1)
}

var _espressoPlanShareIntermediateBuffer func(a0 uintptr, a1 uintptr) uintptr
var _espressoPlanShareIntermediateBufferErr error

func tryEspressoPlanShareIntermediateBuffer(a0 uintptr, a1 uintptr) (uintptr, error) {
	if _espressoPlanShareIntermediateBuffer == nil {
		return 0, symbolCallError("espresso_plan_share_intermediate_buffer", "", _espressoPlanShareIntermediateBufferErr)
	}
	return _espressoPlanShareIntermediateBuffer(a0, a1), nil
}

// EspressoPlanShareIntermediateBuffer.
func EspressoPlanShareIntermediateBuffer(a0 uintptr, a1 uintptr) (uintptr, error) {
	return tryEspressoPlanShareIntermediateBuffer(a0, a1)
}

var _espressoPlanStartProfiling func(a0 uintptr)
var _espressoPlanStartProfilingErr error

func tryEspressoPlanStartProfiling(a0 uintptr) error {
	if _espressoPlanStartProfiling == nil {
		return symbolCallError("espresso_plan_start_profiling", "", _espressoPlanStartProfilingErr)
	}
	_espressoPlanStartProfiling(a0)
	return nil
}

// EspressoPlanStartProfiling.
func EspressoPlanStartProfiling(a0 uintptr) error {
	return tryEspressoPlanStartProfiling(a0)
}

var _espressoPlanStartProfilingWithOptions func(a0 uintptr, a1 uintptr) uintptr
var _espressoPlanStartProfilingWithOptionsErr error

func tryEspressoPlanStartProfilingWithOptions(a0 uintptr, a1 uintptr) (uintptr, error) {
	if _espressoPlanStartProfilingWithOptions == nil {
		return 0, symbolCallError("espresso_plan_start_profiling_with_options", "", _espressoPlanStartProfilingWithOptionsErr)
	}
	return _espressoPlanStartProfilingWithOptions(a0, a1), nil
}

// EspressoPlanStartProfilingWithOptions.
func EspressoPlanStartProfilingWithOptions(a0 uintptr, a1 uintptr) (uintptr, error) {
	return tryEspressoPlanStartProfilingWithOptions(a0, a1)
}

var _espressoPlanStaticProfilingInfo func(a0 uintptr) uintptr
var _espressoPlanStaticProfilingInfoErr error

func tryEspressoPlanStaticProfilingInfo(a0 uintptr) (uintptr, error) {
	if _espressoPlanStaticProfilingInfo == nil {
		return 0, symbolCallError("espresso_plan_static_profiling_info", "", _espressoPlanStaticProfilingInfoErr)
	}
	return _espressoPlanStaticProfilingInfo(a0), nil
}

// EspressoPlanStaticProfilingInfo.
func EspressoPlanStaticProfilingInfo(a0 uintptr) (uintptr, error) {
	return tryEspressoPlanStaticProfilingInfo(a0)
}

var _espressoPlanSubmit func(a0 uintptr, a1 uintptr, a2 uintptr) uintptr
var _espressoPlanSubmitErr error

func tryEspressoPlanSubmit(a0 uintptr, a1 uintptr, a2 uintptr) (uintptr, error) {
	if _espressoPlanSubmit == nil {
		return 0, symbolCallError("espresso_plan_submit", "", _espressoPlanSubmitErr)
	}
	return _espressoPlanSubmit(a0, a1, a2), nil
}

// EspressoPlanSubmit.
func EspressoPlanSubmit(a0 uintptr, a1 uintptr, a2 uintptr) (uintptr, error) {
	return tryEspressoPlanSubmit(a0, a1, a2)
}

var _espressoPlanSubmitCamera func(a0 uintptr, a1 uintptr) uintptr
var _espressoPlanSubmitCameraErr error

func tryEspressoPlanSubmitCamera(a0 uintptr, a1 uintptr) (uintptr, error) {
	if _espressoPlanSubmitCamera == nil {
		return 0, symbolCallError("espresso_plan_submit_camera", "", _espressoPlanSubmitCameraErr)
	}
	return _espressoPlanSubmitCamera(a0, a1), nil
}

// EspressoPlanSubmitCamera.
func EspressoPlanSubmitCamera(a0 uintptr, a1 uintptr) (uintptr, error) {
	return tryEspressoPlanSubmitCamera(a0, a1)
}

var _espressoPlanSubmitSetMultipleBuffering func(a0 uintptr, a1 uint64) int32
var _espressoPlanSubmitSetMultipleBufferingErr error

func tryEspressoPlanSubmitSetMultipleBuffering(a0 uintptr, a1 uint64) (int32, error) {
	if _espressoPlanSubmitSetMultipleBuffering == nil {
		return 0, symbolCallError("espresso_plan_submit_set_multiple_buffering", "", _espressoPlanSubmitSetMultipleBufferingErr)
	}
	return _espressoPlanSubmitSetMultipleBuffering(a0, a1), nil
}

// EspressoPlanSubmitSetMultipleBuffering.
func EspressoPlanSubmitSetMultipleBuffering(a0 uintptr, a1 uint64) (int32, error) {
	return tryEspressoPlanSubmitSetMultipleBuffering(a0, a1)
}

var _espressoPlanSubmitWithArgs func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) int32
var _espressoPlanSubmitWithArgsErr error

func tryEspressoPlanSubmitWithArgs(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	if _espressoPlanSubmitWithArgs == nil {
		return 0, symbolCallError("espresso_plan_submit_with_args", "", _espressoPlanSubmitWithArgsErr)
	}
	return _espressoPlanSubmitWithArgs(a0, a1, a2, a3, a4), nil
}

// EspressoPlanSubmitWithArgs.
func EspressoPlanSubmitWithArgs(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	return tryEspressoPlanSubmitWithArgs(a0, a1, a2, a3, a4)
}

var _espressoPlanWipeTempoary func(a0 uintptr) uintptr
var _espressoPlanWipeTempoaryErr error

func tryEspressoPlanWipeTempoary(a0 uintptr) (uintptr, error) {
	if _espressoPlanWipeTempoary == nil {
		return 0, symbolCallError("espresso_plan_wipe_tempoary", "", _espressoPlanWipeTempoaryErr)
	}
	return _espressoPlanWipeTempoary(a0), nil
}

// EspressoPlanWipeTempoary.
func EspressoPlanWipeTempoary(a0 uintptr) (uintptr, error) {
	return tryEspressoPlanWipeTempoary(a0)
}

var _espressoSetAnalysisModelMetadataForKey func(a0 EspressoNetworkCStruct, a1 *byte, a2 uintptr) uintptr
var _espressoSetAnalysisModelMetadataForKeyErr error

func tryEspressoSetAnalysisModelMetadataForKey(a0 EspressoNetworkCStruct, a1 *byte, a2 uintptr) (uintptr, error) {
	if _espressoSetAnalysisModelMetadataForKey == nil {
		return 0, symbolCallError("espresso_set_analysis_model_metadata_for_key", "", _espressoSetAnalysisModelMetadataForKeyErr)
	}
	return _espressoSetAnalysisModelMetadataForKey(a0, a1, a2), nil
}

// EspressoSetAnalysisModelMetadataForKey.
func EspressoSetAnalysisModelMetadataForKey(a0 EspressoNetworkCStruct, a1 *byte, a2 uintptr) (uintptr, error) {
	return tryEspressoSetAnalysisModelMetadataForKey(a0, a1, a2)
}

var _espressoSetCompilationPlatform func(a0 uint32, a1 uint32) uintptr
var _espressoSetCompilationPlatformErr error

func tryEspressoSetCompilationPlatform(a0 uint32, a1 uint32) (uintptr, error) {
	if _espressoSetCompilationPlatform == nil {
		return 0, symbolCallError("espresso_set_compilation_platform", "", _espressoSetCompilationPlatformErr)
	}
	return _espressoSetCompilationPlatform(a0, a1), nil
}

// EspressoSetCompilationPlatform.
func EspressoSetCompilationPlatform(a0 uint32, a1 uint32) (uintptr, error) {
	return tryEspressoSetCompilationPlatform(a0, a1)
}

var _espressoSetMilConstValues func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) uintptr
var _espressoSetMilConstValuesErr error

func tryEspressoSetMilConstValues(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (uintptr, error) {
	if _espressoSetMilConstValues == nil {
		return 0, symbolCallError("espresso_set_mil_const_values", "", _espressoSetMilConstValuesErr)
	}
	return _espressoSetMilConstValues(a0, a1, a2, a3, a4, a5), nil
}

// EspressoSetMilConstValues.
func EspressoSetMilConstValues(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (uintptr, error) {
	return tryEspressoSetMilConstValues(a0, a1, a2, a3, a4, a5)
}

var _espressoTmpEnableMontreal func(a0 uintptr) uintptr
var _espressoTmpEnableMontrealErr error

func tryEspressoTmpEnableMontreal(a0 uintptr) (uintptr, error) {
	if _espressoTmpEnableMontreal == nil {
		return 0, symbolCallError("espresso_tmp_enable_montreal", "", _espressoTmpEnableMontrealErr)
	}
	return _espressoTmpEnableMontreal(a0), nil
}

// EspressoTmpEnableMontreal.
func EspressoTmpEnableMontreal(a0 uintptr) (uintptr, error) {
	return tryEspressoTmpEnableMontreal(a0)
}

var _espressoUpgradeEirToMil func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr, a7 uintptr) uintptr
var _espressoUpgradeEirToMilErr error

func tryEspressoUpgradeEirToMil(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr, a7 uintptr) (uintptr, error) {
	if _espressoUpgradeEirToMil == nil {
		return 0, symbolCallError("espresso_upgrade_eir_to_mil", "", _espressoUpgradeEirToMilErr)
	}
	return _espressoUpgradeEirToMil(a0, a1, a2, a3, a4, a5, a6, a7), nil
}

// EspressoUpgradeEirToMil.
func EspressoUpgradeEirToMil(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr, a7 uintptr) (uintptr, error) {
	return tryEspressoUpgradeEirToMil(a0, a1, a2, a3, a4, a5, a6, a7)
}

var _espressoUpgradeNetToMil func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) uintptr
var _espressoUpgradeNetToMilErr error

func tryEspressoUpgradeNetToMil(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (uintptr, error) {
	if _espressoUpgradeNetToMil == nil {
		return 0, symbolCallError("espresso_upgrade_net_to_mil", "", _espressoUpgradeNetToMilErr)
	}
	return _espressoUpgradeNetToMil(a0, a1, a2, a3, a4), nil
}

// EspressoUpgradeNetToMil.
func EspressoUpgradeNetToMil(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (uintptr, error) {
	return tryEspressoUpgradeNetToMil(a0, a1, a2, a3, a4)
}

var _espressoUpgradeNetToMilProgram func(a0 uintptr, a1 uintptr, a2 uintptr) uintptr
var _espressoUpgradeNetToMilProgramErr error

func tryEspressoUpgradeNetToMilProgram(a0 uintptr, a1 uintptr, a2 uintptr) (uintptr, error) {
	if _espressoUpgradeNetToMilProgram == nil {
		return 0, symbolCallError("espresso_upgrade_net_to_mil_program", "", _espressoUpgradeNetToMilProgramErr)
	}
	return _espressoUpgradeNetToMilProgram(a0, a1, a2), nil
}

// EspressoUpgradeNetToMilProgram.
func EspressoUpgradeNetToMilProgram(a0 uintptr, a1 uintptr, a2 uintptr) (uintptr, error) {
	return tryEspressoUpgradeNetToMilProgram(a0, a1, a2)
}

var _espressoUpgradeToMil func(a0 uintptr, a1 uintptr, a2 uintptr) uintptr
var _espressoUpgradeToMilErr error

func tryEspressoUpgradeToMil(a0 uintptr, a1 uintptr, a2 uintptr) (uintptr, error) {
	if _espressoUpgradeToMil == nil {
		return 0, symbolCallError("espresso_upgrade_to_mil", "", _espressoUpgradeToMilErr)
	}
	return _espressoUpgradeToMil(a0, a1, a2), nil
}

// EspressoUpgradeToMil.
func EspressoUpgradeToMil(a0 uintptr, a1 uintptr, a2 uintptr) (uintptr, error) {
	return tryEspressoUpgradeToMil(a0, a1, a2)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_e5rtAneMemoryProviderCreate, &_e5rtAneMemoryProviderCreateErr, frameworkHandle, "e5rt_ane_memory_provider_create", "")
	registerFunc(&_e5rtAneMemoryProviderRelease, &_e5rtAneMemoryProviderReleaseErr, frameworkHandle, "e5rt_ane_memory_provider_release", "")
	registerFunc(&_e5rtAsyncEventAsyncNotify, &_e5rtAsyncEventAsyncNotifyErr, frameworkHandle, "e5rt_async_event_async_notify", "")
	registerFunc(&_e5rtAsyncEventCreate, &_e5rtAsyncEventCreateErr, frameworkHandle, "e5rt_async_event_create", "")
	registerFunc(&_e5rtAsyncEventCreateFromIosurfaceSharedEvent, &_e5rtAsyncEventCreateFromIosurfaceSharedEventErr, frameworkHandle, "e5rt_async_event_create_from_iosurface_shared_event", "")
	registerFunc(&_e5rtAsyncEventGetActiveFutureValue, &_e5rtAsyncEventGetActiveFutureValueErr, frameworkHandle, "e5rt_async_event_get_active_future_value", "")
	registerFunc(&_e5rtAsyncEventGetIosurfaceSharedEvent, &_e5rtAsyncEventGetIosurfaceSharedEventErr, frameworkHandle, "e5rt_async_event_get_iosurface_shared_event", "")
	registerFunc(&_e5rtAsyncEventGetLastSignaledValue, &_e5rtAsyncEventGetLastSignaledValueErr, frameworkHandle, "e5rt_async_event_get_last_signaled_value", "")
	registerFunc(&_e5rtAsyncEventGetName, &_e5rtAsyncEventGetNameErr, frameworkHandle, "e5rt_async_event_get_name", "")
	registerFunc(&_e5rtAsyncEventRelease, &_e5rtAsyncEventReleaseErr, frameworkHandle, "e5rt_async_event_release", "")
	registerFunc(&_e5rtAsyncEventSetActiveFutureValue, &_e5rtAsyncEventSetActiveFutureValueErr, frameworkHandle, "e5rt_async_event_set_active_future_value", "")
	registerFunc(&_e5rtAsyncEventSignal, &_e5rtAsyncEventSignalErr, frameworkHandle, "e5rt_async_event_signal", "")
	registerFunc(&_e5rtAsyncEventSyncWait, &_e5rtAsyncEventSyncWaitErr, frameworkHandle, "e5rt_async_event_sync_wait", "")
	registerFunc(&_e5rtBufferObjectAlloc, &_e5rtBufferObjectAllocErr, frameworkHandle, "e5rt_buffer_object_alloc", "")
	registerFunc(&_e5rtBufferObjectCreateAsAlias, &_e5rtBufferObjectCreateAsAliasErr, frameworkHandle, "e5rt_buffer_object_create_as_alias", "")
	registerFunc(&_e5rtBufferObjectCreateFromDataPointer, &_e5rtBufferObjectCreateFromDataPointerErr, frameworkHandle, "e5rt_buffer_object_create_from_data_pointer", "")
	registerFunc(&_e5rtBufferObjectCreateFromIosurface, &_e5rtBufferObjectCreateFromIosurfaceErr, frameworkHandle, "e5rt_buffer_object_create_from_iosurface", "")
	registerFunc(&_e5rtBufferObjectCreateFromMtlbuffer, &_e5rtBufferObjectCreateFromMtlbufferErr, frameworkHandle, "e5rt_buffer_object_create_from_mtlbuffer", "")
	registerFunc(&_e5rtBufferObjectGetDataPtr, &_e5rtBufferObjectGetDataPtrErr, frameworkHandle, "e5rt_buffer_object_get_data_ptr", "")
	registerFunc(&_e5rtBufferObjectGetIosurface, &_e5rtBufferObjectGetIosurfaceErr, frameworkHandle, "e5rt_buffer_object_get_iosurface", "")
	registerFunc(&_e5rtBufferObjectGetMtlbuffer, &_e5rtBufferObjectGetMtlbufferErr, frameworkHandle, "e5rt_buffer_object_get_mtlbuffer", "")
	registerFunc(&_e5rtBufferObjectGetSize, &_e5rtBufferObjectGetSizeErr, frameworkHandle, "e5rt_buffer_object_get_size", "")
	registerFunc(&_e5rtBufferObjectGetType, &_e5rtBufferObjectGetTypeErr, frameworkHandle, "e5rt_buffer_object_get_type", "")
	registerFunc(&_e5rtBufferObjectRelease, &_e5rtBufferObjectReleaseErr, frameworkHandle, "e5rt_buffer_object_release", "")
	registerFunc(&_e5rtComputeGPUDeviceGetMtlDevice, &_e5rtComputeGPUDeviceGetMtlDeviceErr, frameworkHandle, "e5rt_compute_gpu_device_get_mtl_device", "")
	registerFunc(&_e5rtComputeGPUDeviceRelease, &_e5rtComputeGPUDeviceReleaseErr, frameworkHandle, "e5rt_compute_gpu_device_release", "")
	registerFunc(&_e5rtComputeGPUDeviceRetainAll, &_e5rtComputeGPUDeviceRetainAllErr, frameworkHandle, "e5rt_compute_gpu_device_retain_all", "")
	registerFunc(&_e5rtComputeGPUDeviceRetainFromMtlDevice, &_e5rtComputeGPUDeviceRetainFromMtlDeviceErr, frameworkHandle, "e5rt_compute_gpu_device_retain_from_mtl_device", "")
	registerFunc(&_e5rtCreateSurfaceObjectFromIosurface, &_e5rtCreateSurfaceObjectFromIosurfaceErr, frameworkHandle, "e5rt_create_surface_object_from_iosurface", "")
	registerFunc(&_e5rtCvpb4ccToSurfaceFormat, &_e5rtCvpb4ccToSurfaceFormatErr, frameworkHandle, "e5rt_cvpb_4cc_to_surface_format", "")
	registerFunc(&_e5rtE5CompilerCompile, &_e5rtE5CompilerCompileErr, frameworkHandle, "e5rt_e5_compiler_compile", "")
	registerFunc(&_e5rtE5CompilerCompileFromIrProgram, &_e5rtE5CompilerCompileFromIrProgramErr, frameworkHandle, "e5rt_e5_compiler_compile_from_ir_program", "")
	registerFunc(&_e5rtE5CompilerConfigOptionsCreate, &_e5rtE5CompilerConfigOptionsCreateErr, frameworkHandle, "e5rt_e5_compiler_config_options_create", "")
	registerFunc(&_e5rtE5CompilerConfigOptionsGetBundleCacheApfsPurgeable, &_e5rtE5CompilerConfigOptionsGetBundleCacheApfsPurgeableErr, frameworkHandle, "e5rt_e5_compiler_config_options_get_bundle_cache_apfs_purgeable", "")
	registerFunc(&_e5rtE5CompilerConfigOptionsGetCacheBundleLocation, &_e5rtE5CompilerConfigOptionsGetCacheBundleLocationErr, frameworkHandle, "e5rt_e5_compiler_config_options_get_cache_bundle_location", "")
	registerFunc(&_e5rtE5CompilerConfigOptionsRelease, &_e5rtE5CompilerConfigOptionsReleaseErr, frameworkHandle, "e5rt_e5_compiler_config_options_release", "")
	registerFunc(&_e5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeable, &_e5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeableErr, frameworkHandle, "e5rt_e5_compiler_config_options_set_bundle_cache_apfs_purgeable", "")
	registerFunc(&_e5rtE5CompilerConfigOptionsSetCacheBundleLocation, &_e5rtE5CompilerConfigOptionsSetCacheBundleLocationErr, frameworkHandle, "e5rt_e5_compiler_config_options_set_cache_bundle_location", "")
	registerFunc(&_e5rtE5CompilerCreate, &_e5rtE5CompilerCreateErr, frameworkHandle, "e5rt_e5_compiler_create", "")
	registerFunc(&_e5rtE5CompilerCreateWithConfig, &_e5rtE5CompilerCreateWithConfigErr, frameworkHandle, "e5rt_e5_compiler_create_with_config", "")
	registerFunc(&_e5rtE5CompilerIsNewCompileRequired, &_e5rtE5CompilerIsNewCompileRequiredErr, frameworkHandle, "e5rt_e5_compiler_is_new_compile_required", "")
	registerFunc(&_e5rtE5CompilerOptionsCreate, &_e5rtE5CompilerOptionsCreateErr, frameworkHandle, "e5rt_e5_compiler_options_create", "")
	registerFunc(&_e5rtE5CompilerOptionsGetComputeDeviceTypesMask, &_e5rtE5CompilerOptionsGetComputeDeviceTypesMaskErr, frameworkHandle, "e5rt_e5_compiler_options_get_compute_device_types_mask", "")
	registerFunc(&_e5rtE5CompilerOptionsGetCreateProtectedAssets, &_e5rtE5CompilerOptionsGetCreateProtectedAssetsErr, frameworkHandle, "e5rt_e5_compiler_options_get_create_protected_assets", "")
	registerFunc(&_e5rtE5CompilerOptionsGetCustomAneCompilerOptions, &_e5rtE5CompilerOptionsGetCustomAneCompilerOptionsErr, frameworkHandle, "e5rt_e5_compiler_options_get_custom_ane_compiler_options", "")
	registerFunc(&_e5rtE5CompilerOptionsGetEnableMpsgraphPackage, &_e5rtE5CompilerOptionsGetEnableMpsgraphPackageErr, frameworkHandle, "e5rt_e5_compiler_options_get_enable_mpsgraph_package", "")
	registerFunc(&_e5rtE5CompilerOptionsGetEnableProfiling, &_e5rtE5CompilerOptionsGetEnableProfilingErr, frameworkHandle, "e5rt_e5_compiler_options_get_enable_profiling", "")
	registerFunc(&_e5rtE5CompilerOptionsGetEnableReshapeWithMinimalAllocations, &_e5rtE5CompilerOptionsGetEnableReshapeWithMinimalAllocationsErr, frameworkHandle, "e5rt_e5_compiler_options_get_enable_reshape_with_minimal_allocations", "")
	registerFunc(&_e5rtE5CompilerOptionsGetExperimentalDisableCompileTimeMpsgraphTypeInference, &_e5rtE5CompilerOptionsGetExperimentalDisableCompileTimeMpsgraphTypeInferenceErr, frameworkHandle, "e5rt_e5_compiler_options_get_experimental_disable_compile_time_mpsgraph_type_inference", "")
	registerFunc(&_e5rtE5CompilerOptionsGetExperimentalDisableDataDependentShape, &_e5rtE5CompilerOptionsGetExperimentalDisableDataDependentShapeErr, frameworkHandle, "e5rt_e5_compiler_options_get_experimental_disable_data_dependent_shape", "")
	registerFunc(&_e5rtE5CompilerOptionsGetExperimentalEnableDefaultFunctionForRangeDim, &_e5rtE5CompilerOptionsGetExperimentalEnableDefaultFunctionForRangeDimErr, frameworkHandle, "e5rt_e5_compiler_options_get_experimental_enable_default_function_for_range_dim", "")
	registerFunc(&_e5rtE5CompilerOptionsGetExperimentalForceClassicCPUBackend, &_e5rtE5CompilerOptionsGetExperimentalForceClassicCPUBackendErr, frameworkHandle, "e5rt_e5_compiler_options_get_experimental_force_classic_cpu_backend", "")
	registerFunc(&_e5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatterns, &_e5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatternsErr, frameworkHandle, "e5rt_e5_compiler_options_get_experimental_match_e5_minimal_cpu_patterns", "")
	registerFunc(&_e5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatternsForStates, &_e5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatternsForStatesErr, frameworkHandle, "e5rt_e5_compiler_options_get_experimental_match_e5_minimal_cpu_patterns_for_states", "")
	registerFunc(&_e5rtE5CompilerOptionsGetForceBnnsGraph, &_e5rtE5CompilerOptionsGetForceBnnsGraphErr, frameworkHandle, "e5rt_e5_compiler_options_get_force_bnns_graph", "")
	registerFunc(&_e5rtE5CompilerOptionsGetForceClassicAotOldHw, &_e5rtE5CompilerOptionsGetForceClassicAotOldHwErr, frameworkHandle, "e5rt_e5_compiler_options_get_force_classic_aot_old_hw", "")
	registerFunc(&_e5rtE5CompilerOptionsGetForceFetchFromCache, &_e5rtE5CompilerOptionsGetForceFetchFromCacheErr, frameworkHandle, "e5rt_e5_compiler_options_get_force_fetch_from_cache", "")
	registerFunc(&_e5rtE5CompilerOptionsGetForceRecompilation, &_e5rtE5CompilerOptionsGetForceRecompilationErr, frameworkHandle, "e5rt_e5_compiler_options_get_force_recompilation", "")
	registerFunc(&_e5rtE5CompilerOptionsGetPreferredCPUBackend, &_e5rtE5CompilerOptionsGetPreferredCPUBackendErr, frameworkHandle, "e5rt_e5_compiler_options_get_preferred_cpu_backend", "")
	registerFunc(&_e5rtE5CompilerOptionsGetPreferredCPUBackends, &_e5rtE5CompilerOptionsGetPreferredCPUBackendsErr, frameworkHandle, "e5rt_e5_compiler_options_get_preferred_cpu_backends", "")
	registerFunc(&_e5rtE5CompilerOptionsGetSegmenter, &_e5rtE5CompilerOptionsGetSegmenterErr, frameworkHandle, "e5rt_e5_compiler_options_get_segmenter", "")
	registerFunc(&_e5rtE5CompilerOptionsRelease, &_e5rtE5CompilerOptionsReleaseErr, frameworkHandle, "e5rt_e5_compiler_options_release", "")
	registerFunc(&_e5rtE5CompilerOptionsRetainMilEntryPoints, &_e5rtE5CompilerOptionsRetainMilEntryPointsErr, frameworkHandle, "e5rt_e5_compiler_options_retain_mil_entry_points", "")
	registerFunc(&_e5rtE5CompilerOptionsSetComputeDeviceTypesMask, &_e5rtE5CompilerOptionsSetComputeDeviceTypesMaskErr, frameworkHandle, "e5rt_e5_compiler_options_set_compute_device_types_mask", "")
	registerFunc(&_e5rtE5CompilerOptionsSetCreateProtectedAssets, &_e5rtE5CompilerOptionsSetCreateProtectedAssetsErr, frameworkHandle, "e5rt_e5_compiler_options_set_create_protected_assets", "")
	registerFunc(&_e5rtE5CompilerOptionsSetCustomAneCompilerOptions, &_e5rtE5CompilerOptionsSetCustomAneCompilerOptionsErr, frameworkHandle, "e5rt_e5_compiler_options_set_custom_ane_compiler_options", "")
	registerFunc(&_e5rtE5CompilerOptionsSetEnableMpsgraphPackage, &_e5rtE5CompilerOptionsSetEnableMpsgraphPackageErr, frameworkHandle, "e5rt_e5_compiler_options_set_enable_mpsgraph_package", "")
	registerFunc(&_e5rtE5CompilerOptionsSetEnableProfiling, &_e5rtE5CompilerOptionsSetEnableProfilingErr, frameworkHandle, "e5rt_e5_compiler_options_set_enable_profiling", "")
	registerFunc(&_e5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocations, &_e5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocationsErr, frameworkHandle, "e5rt_e5_compiler_options_set_enable_reshape_with_minimal_allocations", "")
	registerFunc(&_e5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference, &_e5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInferenceErr, frameworkHandle, "e5rt_e5_compiler_options_set_experimental_disable_compile_time_mpsgraph_type_inference", "")
	registerFunc(&_e5rtE5CompilerOptionsSetExperimentalDisableDataDependentShape, &_e5rtE5CompilerOptionsSetExperimentalDisableDataDependentShapeErr, frameworkHandle, "e5rt_e5_compiler_options_set_experimental_disable_data_dependent_shape", "")
	registerFunc(&_e5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDim, &_e5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDimErr, frameworkHandle, "e5rt_e5_compiler_options_set_experimental_enable_default_function_for_range_dim", "")
	registerFunc(&_e5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackend, &_e5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackendErr, frameworkHandle, "e5rt_e5_compiler_options_set_experimental_force_classic_cpu_backend", "")
	registerFunc(&_e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatterns, &_e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsErr, frameworkHandle, "e5rt_e5_compiler_options_set_experimental_match_e5_minimal_cpu_patterns", "")
	registerFunc(&_e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStates, &_e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStatesErr, frameworkHandle, "e5rt_e5_compiler_options_set_experimental_match_e5_minimal_cpu_patterns_for_states", "")
	registerFunc(&_e5rtE5CompilerOptionsSetForceBnnsGraph, &_e5rtE5CompilerOptionsSetForceBnnsGraphErr, frameworkHandle, "e5rt_e5_compiler_options_set_force_bnns_graph", "")
	registerFunc(&_e5rtE5CompilerOptionsSetForceClassicAotOldHw, &_e5rtE5CompilerOptionsSetForceClassicAotOldHwErr, frameworkHandle, "e5rt_e5_compiler_options_set_force_classic_aot_old_hw", "")
	registerFunc(&_e5rtE5CompilerOptionsSetForceFetchFromCache, &_e5rtE5CompilerOptionsSetForceFetchFromCacheErr, frameworkHandle, "e5rt_e5_compiler_options_set_force_fetch_from_cache", "")
	registerFunc(&_e5rtE5CompilerOptionsSetForceRecompilation, &_e5rtE5CompilerOptionsSetForceRecompilationErr, frameworkHandle, "e5rt_e5_compiler_options_set_force_recompilation", "")
	registerFunc(&_e5rtE5CompilerOptionsSetMilEntryPoints, &_e5rtE5CompilerOptionsSetMilEntryPointsErr, frameworkHandle, "e5rt_e5_compiler_options_set_mil_entry_points", "")
	registerFunc(&_e5rtE5CompilerOptionsSetPreferredCPUBackend, &_e5rtE5CompilerOptionsSetPreferredCPUBackendErr, frameworkHandle, "e5rt_e5_compiler_options_set_preferred_cpu_backend", "")
	registerFunc(&_e5rtE5CompilerOptionsSetPreferredCPUBackends, &_e5rtE5CompilerOptionsSetPreferredCPUBackendsErr, frameworkHandle, "e5rt_e5_compiler_options_set_preferred_cpu_backends", "")
	registerFunc(&_e5rtE5CompilerOptionsSetSegmenter, &_e5rtE5CompilerOptionsSetSegmenterErr, frameworkHandle, "e5rt_e5_compiler_options_set_segmenter", "")
	registerFunc(&_e5rtE5CompilerPurgeE5BundlesForInputModel, &_e5rtE5CompilerPurgeE5BundlesForInputModelErr, frameworkHandle, "e5rt_e5_compiler_purge_e5_bundles_for_input_model", "")
	registerFunc(&_e5rtE5CompilerRelease, &_e5rtE5CompilerReleaseErr, frameworkHandle, "e5rt_e5_compiler_release", "")
	registerFunc(&_e5rtErrorCodeGetString, &_e5rtErrorCodeGetStringErr, frameworkHandle, "e5rt_error_code_get_string", "")
	registerFunc(&_e5rtExecutionStreamAsyncSubmit, &_e5rtExecutionStreamAsyncSubmitErr, frameworkHandle, "e5rt_execution_stream_async_submit", "")
	registerFunc(&_e5rtExecutionStreamConfigOptionsCreate, &_e5rtExecutionStreamConfigOptionsCreateErr, frameworkHandle, "e5rt_execution_stream_config_options_create", "")
	registerFunc(&_e5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecution, &_e5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecutionErr, frameworkHandle, "e5rt_execution_stream_config_options_get_enable_concurrent_sync_execution", "")
	registerFunc(&_e5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEvents, &_e5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEventsErr, frameworkHandle, "e5rt_execution_stream_config_options_get_enable_low_latency_async_events", "")
	registerFunc(&_e5rtExecutionStreamConfigOptionsGetSkipIOFences, &_e5rtExecutionStreamConfigOptionsGetSkipIOFencesErr, frameworkHandle, "e5rt_execution_stream_config_options_get_skip_io_fences", "")
	registerFunc(&_e5rtExecutionStreamConfigOptionsRelease, &_e5rtExecutionStreamConfigOptionsReleaseErr, frameworkHandle, "e5rt_execution_stream_config_options_release", "")
	registerFunc(&_e5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecution, &_e5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecutionErr, frameworkHandle, "e5rt_execution_stream_config_options_set_enable_concurrent_sync_execution", "")
	registerFunc(&_e5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEvents, &_e5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEventsErr, frameworkHandle, "e5rt_execution_stream_config_options_set_enable_low_latency_async_events", "")
	registerFunc(&_e5rtExecutionStreamConfigOptionsSetSkipIOFences, &_e5rtExecutionStreamConfigOptionsSetSkipIOFencesErr, frameworkHandle, "e5rt_execution_stream_config_options_set_skip_io_fences", "")
	registerFunc(&_e5rtExecutionStreamCreate, &_e5rtExecutionStreamCreateErr, frameworkHandle, "e5rt_execution_stream_create", "")
	registerFunc(&_e5rtExecutionStreamEncodeOperation, &_e5rtExecutionStreamEncodeOperationErr, frameworkHandle, "e5rt_execution_stream_encode_operation", "")
	registerFunc(&_e5rtExecutionStreamEncodeWorkload, &_e5rtExecutionStreamEncodeWorkloadErr, frameworkHandle, "e5rt_execution_stream_encode_workload", "")
	registerFunc(&_e5rtExecutionStreamExecuteSync, &_e5rtExecutionStreamExecuteSyncErr, frameworkHandle, "e5rt_execution_stream_execute_sync", "")
	registerFunc(&_e5rtExecutionStreamGetInternalAsyncComputeRequestIDForLastSubmit, &_e5rtExecutionStreamGetInternalAsyncComputeRequestIDForLastSubmitErr, frameworkHandle, "e5rt_execution_stream_get_internal_async_compute_request_id_for_last_submit", "")
	registerFunc(&_e5rtExecutionStreamGetStreamID, &_e5rtExecutionStreamGetStreamIDErr, frameworkHandle, "e5rt_execution_stream_get_stream_id", "")
	registerFunc(&_e5rtExecutionStreamOperationBindCompletionEvent, &_e5rtExecutionStreamOperationBindCompletionEventErr, frameworkHandle, "e5rt_execution_stream_operation_bind_completion_event", "")
	registerFunc(&_e5rtExecutionStreamOperationBindDependentEvents, &_e5rtExecutionStreamOperationBindDependentEventsErr, frameworkHandle, "e5rt_execution_stream_operation_bind_dependent_events", "")
	registerFunc(&_e5rtExecutionStreamOperationConfigOptionsCreate, &_e5rtExecutionStreamOperationConfigOptionsCreateErr, frameworkHandle, "e5rt_execution_stream_operation_config_options_create", "")
	registerFunc(&_e5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemory, &_e5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemoryErr, frameworkHandle, "e5rt_execution_stream_operation_config_options_get_prewire_model_memory", "")
	registerFunc(&_e5rtExecutionStreamOperationConfigOptionsRelease, &_e5rtExecutionStreamOperationConfigOptionsReleaseErr, frameworkHandle, "e5rt_execution_stream_operation_config_options_release", "")
	registerFunc(&_e5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemory, &_e5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemoryErr, frameworkHandle, "e5rt_execution_stream_operation_config_options_set_prewire_model_memory", "")
	registerFunc(&_e5rtExecutionStreamOperationCreatePrecompiledComputeOperation, &_e5rtExecutionStreamOperationCreatePrecompiledComputeOperationErr, frameworkHandle, "e5rt_execution_stream_operation_create_precompiled_compute_operation", "")
	registerFunc(&_e5rtExecutionStreamOperationCreatePrecompiledComputeOperationWithOptions, &_e5rtExecutionStreamOperationCreatePrecompiledComputeOperationWithOptionsErr, frameworkHandle, "e5rt_execution_stream_operation_create_precompiled_compute_operation_with_options", "")
	registerFunc(&_e5rtExecutionStreamOperationCreateResourceSharingPrecompiledComputeOperationsWithMultipleOptions, &_e5rtExecutionStreamOperationCreateResourceSharingPrecompiledComputeOperationsWithMultipleOptionsErr, frameworkHandle, "e5rt_execution_stream_operation_create_resource_sharing_precompiled_compute_operations_with_multiple_options", "")
	registerFunc(&_e5rtExecutionStreamOperationGetDependentEventCount, &_e5rtExecutionStreamOperationGetDependentEventCountErr, frameworkHandle, "e5rt_execution_stream_operation_get_dependent_event_count", "")
	registerFunc(&_e5rtExecutionStreamOperationGetInoutNames, &_e5rtExecutionStreamOperationGetInoutNamesErr, frameworkHandle, "e5rt_execution_stream_operation_get_inout_names", "")
	registerFunc(&_e5rtExecutionStreamOperationGetInputNames, &_e5rtExecutionStreamOperationGetInputNamesErr, frameworkHandle, "e5rt_execution_stream_operation_get_input_names", "")
	registerFunc(&_e5rtExecutionStreamOperationGetNumInouts, &_e5rtExecutionStreamOperationGetNumInoutsErr, frameworkHandle, "e5rt_execution_stream_operation_get_num_inouts", "")
	registerFunc(&_e5rtExecutionStreamOperationGetNumInputs, &_e5rtExecutionStreamOperationGetNumInputsErr, frameworkHandle, "e5rt_execution_stream_operation_get_num_inputs", "")
	registerFunc(&_e5rtExecutionStreamOperationGetNumOutputs, &_e5rtExecutionStreamOperationGetNumOutputsErr, frameworkHandle, "e5rt_execution_stream_operation_get_num_outputs", "")
	registerFunc(&_e5rtExecutionStreamOperationGetOpname, &_e5rtExecutionStreamOperationGetOpnameErr, frameworkHandle, "e5rt_execution_stream_operation_get_opname", "")
	registerFunc(&_e5rtExecutionStreamOperationGetOutputNames, &_e5rtExecutionStreamOperationGetOutputNamesErr, frameworkHandle, "e5rt_execution_stream_operation_get_output_names", "")
	registerFunc(&_e5rtExecutionStreamOperationPrepareOpForEncode, &_e5rtExecutionStreamOperationPrepareOpForEncodeErr, frameworkHandle, "e5rt_execution_stream_operation_prepare_op_for_encode", "")
	registerFunc(&_e5rtExecutionStreamOperationRelease, &_e5rtExecutionStreamOperationReleaseErr, frameworkHandle, "e5rt_execution_stream_operation_release", "")
	registerFunc(&_e5rtExecutionStreamOperationReshapeOperation, &_e5rtExecutionStreamOperationReshapeOperationErr, frameworkHandle, "e5rt_execution_stream_operation_reshape_operation", "")
	registerFunc(&_e5rtExecutionStreamOperationRetainCompletionEvent, &_e5rtExecutionStreamOperationRetainCompletionEventErr, frameworkHandle, "e5rt_execution_stream_operation_retain_completion_event", "")
	registerFunc(&_e5rtExecutionStreamOperationRetainDependentEvents, &_e5rtExecutionStreamOperationRetainDependentEventsErr, frameworkHandle, "e5rt_execution_stream_operation_retain_dependent_events", "")
	registerFunc(&_e5rtExecutionStreamOperationRetainInoutPort, &_e5rtExecutionStreamOperationRetainInoutPortErr, frameworkHandle, "e5rt_execution_stream_operation_retain_inout_port", "")
	registerFunc(&_e5rtExecutionStreamOperationRetainInputPort, &_e5rtExecutionStreamOperationRetainInputPortErr, frameworkHandle, "e5rt_execution_stream_operation_retain_input_port", "")
	registerFunc(&_e5rtExecutionStreamOperationRetainOutputPort, &_e5rtExecutionStreamOperationRetainOutputPortErr, frameworkHandle, "e5rt_execution_stream_operation_retain_output_port", "")
	registerFunc(&_e5rtExecutionStreamOperationSerializeInferenceFrameData, &_e5rtExecutionStreamOperationSerializeInferenceFrameDataErr, frameworkHandle, "e5rt_execution_stream_operation_serialize_inference_frame_data", "")
	registerFunc(&_e5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegment, &_e5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegmentErr, frameworkHandle, "e5rt_execution_stream_operation_serialize_inference_frame_data_per_segment", "")
	registerFunc(&_e5rtExecutionStreamPrewireInUseAllocations, &_e5rtExecutionStreamPrewireInUseAllocationsErr, frameworkHandle, "e5rt_execution_stream_prewire_in_use_allocations", "")
	registerFunc(&_e5rtExecutionStreamRelease, &_e5rtExecutionStreamReleaseErr, frameworkHandle, "e5rt_execution_stream_release", "")
	registerFunc(&_e5rtExecutionStreamReset, &_e5rtExecutionStreamResetErr, frameworkHandle, "e5rt_execution_stream_reset", "")
	registerFunc(&_e5rtExecutionStreamResetConfigOptions, &_e5rtExecutionStreamResetConfigOptionsErr, frameworkHandle, "e5rt_execution_stream_reset_config_options", "")
	registerFunc(&_e5rtExecutionStreamSetAneExecutionPriority, &_e5rtExecutionStreamSetAneExecutionPriorityErr, frameworkHandle, "e5rt_execution_stream_set_ane_execution_priority", "")
	registerFunc(&_e5rtExecutionStreamSetConfigOptions, &_e5rtExecutionStreamSetConfigOptionsErr, frameworkHandle, "e5rt_execution_stream_set_config_options", "")
	registerFunc(&_e5rtExecutionStreamSetQualityOfService, &_e5rtExecutionStreamSetQualityOfServiceErr, frameworkHandle, "e5rt_execution_stream_set_quality_of_service", "")
	registerFunc(&_e5rtExecutionStreamStepExecuteSync, &_e5rtExecutionStreamStepExecuteSyncErr, frameworkHandle, "e5rt_execution_stream_step_execute_sync", "")
	registerFunc(&_e5rtExecutionStreamSubmitAsync, &_e5rtExecutionStreamSubmitAsyncErr, frameworkHandle, "e5rt_execution_stream_submit_async", "")
	registerFunc(&_e5rtExecutionStreamSubmitAsyncWithTimeout, &_e5rtExecutionStreamSubmitAsyncWithTimeoutErr, frameworkHandle, "e5rt_execution_stream_submit_async_with_timeout", "")
	registerFunc(&_e5rtGetLastErrorMessage, &_e5rtGetLastErrorMessageErr, frameworkHandle, "e5rt_get_last_error_message", "")
	registerFunc(&_e5rtIOPortBindBufferObject, &_e5rtIOPortBindBufferObjectErr, frameworkHandle, "e5rt_io_port_bind_buffer_object", "")
	registerFunc(&_e5rtIOPortBindMemoryObject, &_e5rtIOPortBindMemoryObjectErr, frameworkHandle, "e5rt_io_port_bind_memory_object", "")
	registerFunc(&_e5rtIOPortBindSurfaceObject, &_e5rtIOPortBindSurfaceObjectErr, frameworkHandle, "e5rt_io_port_bind_surface_object", "")
	registerFunc(&_e5rtIOPortGetSupportedBufferTypes, &_e5rtIOPortGetSupportedBufferTypesErr, frameworkHandle, "e5rt_io_port_get_supported_buffer_types", "")
	registerFunc(&_e5rtIOPortHasKnownShape, &_e5rtIOPortHasKnownShapeErr, frameworkHandle, "e5rt_io_port_has_known_shape", "")
	registerFunc(&_e5rtIOPortIsDynamic, &_e5rtIOPortIsDynamicErr, frameworkHandle, "e5rt_io_port_is_dynamic", "")
	registerFunc(&_e5rtIOPortIsSurface, &_e5rtIOPortIsSurfaceErr, frameworkHandle, "e5rt_io_port_is_surface", "")
	registerFunc(&_e5rtIOPortIsTensor, &_e5rtIOPortIsTensorErr, frameworkHandle, "e5rt_io_port_is_tensor", "")
	registerFunc(&_e5rtIOPortRelease, &_e5rtIOPortReleaseErr, frameworkHandle, "e5rt_io_port_release", "")
	registerFunc(&_e5rtIOPortRetainBufferObject, &_e5rtIOPortRetainBufferObjectErr, frameworkHandle, "e5rt_io_port_retain_buffer_object", "")
	registerFunc(&_e5rtIOPortRetainMemoryObject, &_e5rtIOPortRetainMemoryObjectErr, frameworkHandle, "e5rt_io_port_retain_memory_object", "")
	registerFunc(&_e5rtIOPortRetainSurfaceDesc, &_e5rtIOPortRetainSurfaceDescErr, frameworkHandle, "e5rt_io_port_retain_surface_desc", "")
	registerFunc(&_e5rtIOPortRetainSurfaceObject, &_e5rtIOPortRetainSurfaceObjectErr, frameworkHandle, "e5rt_io_port_retain_surface_object", "")
	registerFunc(&_e5rtIOPortRetainTensorDesc, &_e5rtIOPortRetainTensorDescErr, frameworkHandle, "e5rt_io_port_retain_tensor_desc", "")
	registerFunc(&_e5rtMemoryObjectCreate, &_e5rtMemoryObjectCreateErr, frameworkHandle, "e5rt_memory_object_create", "")
	registerFunc(&_e5rtMemoryObjectCreateAsAlias, &_e5rtMemoryObjectCreateAsAliasErr, frameworkHandle, "e5rt_memory_object_create_as_alias", "")
	registerFunc(&_e5rtMemoryObjectCreateFromIosurface, &_e5rtMemoryObjectCreateFromIosurfaceErr, frameworkHandle, "e5rt_memory_object_create_from_iosurface", "")
	registerFunc(&_e5rtMemoryObjectGetDataPtr, &_e5rtMemoryObjectGetDataPtrErr, frameworkHandle, "e5rt_memory_object_get_data_ptr", "")
	registerFunc(&_e5rtMemoryObjectGetIosurface, &_e5rtMemoryObjectGetIosurfaceErr, frameworkHandle, "e5rt_memory_object_get_iosurface", "")
	registerFunc(&_e5rtMemoryObjectGetSize, &_e5rtMemoryObjectGetSizeErr, frameworkHandle, "e5rt_memory_object_get_size", "")
	registerFunc(&_e5rtMemoryObjectRelease, &_e5rtMemoryObjectReleaseErr, frameworkHandle, "e5rt_memory_object_release", "")
	registerFunc(&_e5rtOperandDescIsSurfaceDesc, &_e5rtOperandDescIsSurfaceDescErr, frameworkHandle, "e5rt_operand_desc_is_surface_desc", "")
	registerFunc(&_e5rtOperandDescIsTensorDesc, &_e5rtOperandDescIsTensorDescErr, frameworkHandle, "e5rt_operand_desc_is_tensor_desc", "")
	registerFunc(&_e5rtOperandDescRelease, &_e5rtOperandDescReleaseErr, frameworkHandle, "e5rt_operand_desc_release", "")
	registerFunc(&_e5rtOperandDescRetainFromSurfaceDesc, &_e5rtOperandDescRetainFromSurfaceDescErr, frameworkHandle, "e5rt_operand_desc_retain_from_surface_desc", "")
	registerFunc(&_e5rtOperandDescRetainFromTensorDesc, &_e5rtOperandDescRetainFromTensorDescErr, frameworkHandle, "e5rt_operand_desc_retain_from_tensor_desc", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsCopyDynamicCallables, &_e5rtPrecompiledComputeOpCreateOptionsCopyDynamicCallablesErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_copy_dynamic_callables", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsCopyMutableMilWeightPaths, &_e5rtPrecompiledComputeOpCreateOptionsCopyMutableMilWeightPathsErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_copy_mutable_mil_weight_paths", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsCreate, &_e5rtPrecompiledComputeOpCreateOptionsCreateErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_create", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsCreateWithProgramFunction, &_e5rtPrecompiledComputeOpCreateOptionsCreateWithProgramFunctionErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_create_with_program_function", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsGetAllocateIntermediateBuffers, &_e5rtPrecompiledComputeOpCreateOptionsGetAllocateIntermediateBuffersErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_get_allocate_intermediate_buffers", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsGetExperimentalEnableMpsgraphParallelEncode, &_e5rtPrecompiledComputeOpCreateOptionsGetExperimentalEnableMpsgraphParallelEncodeErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_get_experimental_enable_mpsgraph_parallel_encode", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsGetIosurfaceMemoryPoolID, &_e5rtPrecompiledComputeOpCreateOptionsGetIosurfaceMemoryPoolIDErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_get_iosurface_memory_pool_id", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsGetLazyPrepareOpForEncode, &_e5rtPrecompiledComputeOpCreateOptionsGetLazyPrepareOpForEncodeErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_get_lazy_prepare_op_for_encode", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsGetOperationName, &_e5rtPrecompiledComputeOpCreateOptionsGetOperationNameErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_get_operation_name", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsRelease, &_e5rtPrecompiledComputeOpCreateOptionsReleaseErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_release", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDevice, &_e5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDeviceErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_retain_override_compute_gpu_device", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffers, &_e5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffersErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_set_allocate_intermediate_buffers", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProvider, &_e5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProviderErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_set_custom_ane_memory_provider", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsSetDynamicCallables, &_e5rtPrecompiledComputeOpCreateOptionsSetDynamicCallablesErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_set_dynamic_callables", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference, &_e5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInferenceErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_set_experimental_disable_compile_time_mpsgraph_type_inference", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOps, &_e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOpsErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_set_experimental_enable_gpu_quant_ops", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecision, &_e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecisionErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_set_experimental_enable_mps_reduced_precision", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsgraphParallelEncode, &_e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsgraphParallelEncodeErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_set_experimental_enable_mpsgraph_parallel_encode", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsSetExperimentalMpsgraphMaximumNumberOfEncodingThreads, &_e5rtPrecompiledComputeOpCreateOptionsSetExperimentalMpsgraphMaximumNumberOfEncodingThreadsErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_set_experimental_mpsgraph_maximum_number_of_encoding_threads", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsSetIosurfaceMemoryPoolID, &_e5rtPrecompiledComputeOpCreateOptionsSetIosurfaceMemoryPoolIDErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_set_iosurface_memory_pool_id", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncode, &_e5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncodeErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_set_lazy_prepare_op_for_encode", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPaths, &_e5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPathsErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_set_mutable_mil_weight_paths", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsSetOperationName, &_e5rtPrecompiledComputeOpCreateOptionsSetOperationNameErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_set_operation_name", "")
	registerFunc(&_e5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDevice, &_e5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDeviceErr, frameworkHandle, "e5rt_precompiled_compute_op_create_options_set_override_compute_gpu_device", "")
	registerFunc(&_e5rtProgramFunctionGetExternInoutNames, &_e5rtProgramFunctionGetExternInoutNamesErr, frameworkHandle, "e5rt_program_function_get_extern_inout_names", "")
	registerFunc(&_e5rtProgramFunctionGetExternInputNames, &_e5rtProgramFunctionGetExternInputNamesErr, frameworkHandle, "e5rt_program_function_get_extern_input_names", "")
	registerFunc(&_e5rtProgramFunctionGetExternOutputNames, &_e5rtProgramFunctionGetExternOutputNamesErr, frameworkHandle, "e5rt_program_function_get_extern_output_names", "")
	registerFunc(&_e5rtProgramFunctionGetName, &_e5rtProgramFunctionGetNameErr, frameworkHandle, "e5rt_program_function_get_name", "")
	registerFunc(&_e5rtProgramFunctionGetNumExternInouts, &_e5rtProgramFunctionGetNumExternInoutsErr, frameworkHandle, "e5rt_program_function_get_num_extern_inouts", "")
	registerFunc(&_e5rtProgramFunctionGetNumExternInputs, &_e5rtProgramFunctionGetNumExternInputsErr, frameworkHandle, "e5rt_program_function_get_num_extern_inputs", "")
	registerFunc(&_e5rtProgramFunctionGetNumExternOutputs, &_e5rtProgramFunctionGetNumExternOutputsErr, frameworkHandle, "e5rt_program_function_get_num_extern_outputs", "")
	registerFunc(&_e5rtProgramFunctionLoadForExecution, &_e5rtProgramFunctionLoadForExecutionErr, frameworkHandle, "e5rt_program_function_load_for_execution", "")
	registerFunc(&_e5rtProgramFunctionRelease, &_e5rtProgramFunctionReleaseErr, frameworkHandle, "e5rt_program_function_release", "")
	registerFunc(&_e5rtProgramFunctionRetainExternInputIOPort, &_e5rtProgramFunctionRetainExternInputIOPortErr, frameworkHandle, "e5rt_program_function_retain_extern_input_io_port", "")
	registerFunc(&_e5rtProgramFunctionRetainExternOutputIOPort, &_e5rtProgramFunctionRetainExternOutputIOPortErr, frameworkHandle, "e5rt_program_function_retain_extern_output_io_port", "")
	registerFunc(&_e5rtProgramFunctionRetainInoutSurfaceDesc, &_e5rtProgramFunctionRetainInoutSurfaceDescErr, frameworkHandle, "e5rt_program_function_retain_inout_surface_desc", "")
	registerFunc(&_e5rtProgramFunctionRetainInoutTensorDesc, &_e5rtProgramFunctionRetainInoutTensorDescErr, frameworkHandle, "e5rt_program_function_retain_inout_tensor_desc", "")
	registerFunc(&_e5rtProgramFunctionRetainInputSurfaceDesc, &_e5rtProgramFunctionRetainInputSurfaceDescErr, frameworkHandle, "e5rt_program_function_retain_input_surface_desc", "")
	registerFunc(&_e5rtProgramFunctionRetainInputTensorDesc, &_e5rtProgramFunctionRetainInputTensorDescErr, frameworkHandle, "e5rt_program_function_retain_input_tensor_desc", "")
	registerFunc(&_e5rtProgramFunctionRetainOutputSurfaceDesc, &_e5rtProgramFunctionRetainOutputSurfaceDescErr, frameworkHandle, "e5rt_program_function_retain_output_surface_desc", "")
	registerFunc(&_e5rtProgramFunctionRetainOutputTensorDesc, &_e5rtProgramFunctionRetainOutputTensorDescErr, frameworkHandle, "e5rt_program_function_retain_output_tensor_desc", "")
	registerFunc(&_e5rtProgramLibraryCreate, &_e5rtProgramLibraryCreateErr, frameworkHandle, "e5rt_program_library_create", "")
	registerFunc(&_e5rtProgramLibraryGetBuildInfo, &_e5rtProgramLibraryGetBuildInfoErr, frameworkHandle, "e5rt_program_library_get_build_info", "")
	registerFunc(&_e5rtProgramLibraryGetE5BundlePath, &_e5rtProgramLibraryGetE5BundlePathErr, frameworkHandle, "e5rt_program_library_get_e5_bundle_path", "")
	registerFunc(&_e5rtProgramLibraryGetFunctionMetadata, &_e5rtProgramLibraryGetFunctionMetadataErr, frameworkHandle, "e5rt_program_library_get_function_metadata", "")
	registerFunc(&_e5rtProgramLibraryGetFunctionNames, &_e5rtProgramLibraryGetFunctionNamesErr, frameworkHandle, "e5rt_program_library_get_function_names", "")
	registerFunc(&_e5rtProgramLibraryGetNumFunctions, &_e5rtProgramLibraryGetNumFunctionsErr, frameworkHandle, "e5rt_program_library_get_num_functions", "")
	registerFunc(&_e5rtProgramLibraryGetSegmentationAnalytics, &_e5rtProgramLibraryGetSegmentationAnalyticsErr, frameworkHandle, "e5rt_program_library_get_segmentation_analytics", "")
	registerFunc(&_e5rtProgramLibraryRelease, &_e5rtProgramLibraryReleaseErr, frameworkHandle, "e5rt_program_library_release", "")
	registerFunc(&_e5rtProgramLibraryRetainProgramFunction, &_e5rtProgramLibraryRetainProgramFunctionErr, frameworkHandle, "e5rt_program_library_retain_program_function", "")
	registerFunc(&_e5rtSurfaceDescCreate, &_e5rtSurfaceDescCreateErr, frameworkHandle, "e5rt_surface_desc_create", "")
	registerFunc(&_e5rtSurfaceDescCreateFromOperandDesc, &_e5rtSurfaceDescCreateFromOperandDescErr, frameworkHandle, "e5rt_surface_desc_create_from_operand_desc", "")
	registerFunc(&_e5rtSurfaceDescCreateWithSlices, &_e5rtSurfaceDescCreateWithSlicesErr, frameworkHandle, "e5rt_surface_desc_create_with_slices", "")
	registerFunc(&_e5rtSurfaceDescCreateWithStrides, &_e5rtSurfaceDescCreateWithStridesErr, frameworkHandle, "e5rt_surface_desc_create_with_strides", "")
	registerFunc(&_e5rtSurfaceDescCreateWithStridesAndSlices, &_e5rtSurfaceDescCreateWithStridesAndSlicesErr, frameworkHandle, "e5rt_surface_desc_create_with_strides_and_slices", "")
	registerFunc(&_e5rtSurfaceDescGetCustomRowStrides, &_e5rtSurfaceDescGetCustomRowStridesErr, frameworkHandle, "e5rt_surface_desc_get_custom_row_strides", "")
	registerFunc(&_e5rtSurfaceDescGetFormat, &_e5rtSurfaceDescGetFormatErr, frameworkHandle, "e5rt_surface_desc_get_format", "")
	registerFunc(&_e5rtSurfaceDescGetHeight, &_e5rtSurfaceDescGetHeightErr, frameworkHandle, "e5rt_surface_desc_get_height", "")
	registerFunc(&_e5rtSurfaceDescGetPlaneCount, &_e5rtSurfaceDescGetPlaneCountErr, frameworkHandle, "e5rt_surface_desc_get_plane_count", "")
	registerFunc(&_e5rtSurfaceDescGetSliceCount, &_e5rtSurfaceDescGetSliceCountErr, frameworkHandle, "e5rt_surface_desc_get_slice_count", "")
	registerFunc(&_e5rtSurfaceDescGetWidth, &_e5rtSurfaceDescGetWidthErr, frameworkHandle, "e5rt_surface_desc_get_width", "")
	registerFunc(&_e5rtSurfaceDescRelease, &_e5rtSurfaceDescReleaseErr, frameworkHandle, "e5rt_surface_desc_release", "")
	registerFunc(&_e5rtSurfaceFormatToCvpb4cc, &_e5rtSurfaceFormatToCvpb4ccErr, frameworkHandle, "e5rt_surface_format_to_cvpb_4cc", "")
	registerFunc(&_e5rtSurfaceObjectAlloc, &_e5rtSurfaceObjectAllocErr, frameworkHandle, "e5rt_surface_object_alloc", "")
	registerFunc(&_e5rtSurfaceObjectCreateFromIosurface, &_e5rtSurfaceObjectCreateFromIosurfaceErr, frameworkHandle, "e5rt_surface_object_create_from_iosurface", "")
	registerFunc(&_e5rtSurfaceObjectGetIosurface, &_e5rtSurfaceObjectGetIosurfaceErr, frameworkHandle, "e5rt_surface_object_get_iosurface", "")
	registerFunc(&_e5rtSurfaceObjectRelease, &_e5rtSurfaceObjectReleaseErr, frameworkHandle, "e5rt_surface_object_release", "")
	registerFunc(&_e5rtTensorDescAllocBufferObject, &_e5rtTensorDescAllocBufferObjectErr, frameworkHandle, "e5rt_tensor_desc_alloc_buffer_object", "")
	registerFunc(&_e5rtTensorDescCreate, &_e5rtTensorDescCreateErr, frameworkHandle, "e5rt_tensor_desc_create", "")
	registerFunc(&_e5rtTensorDescCreateFromOperandDesc, &_e5rtTensorDescCreateFromOperandDescErr, frameworkHandle, "e5rt_tensor_desc_create_from_operand_desc", "")
	registerFunc(&_e5rtTensorDescCreateMemoryObject, &_e5rtTensorDescCreateMemoryObjectErr, frameworkHandle, "e5rt_tensor_desc_create_memory_object", "")
	registerFunc(&_e5rtTensorDescCreateSlice, &_e5rtTensorDescCreateSliceErr, frameworkHandle, "e5rt_tensor_desc_create_slice", "")
	registerFunc(&_e5rtTensorDescCreateSliceWithLengths, &_e5rtTensorDescCreateSliceWithLengthsErr, frameworkHandle, "e5rt_tensor_desc_create_slice_with_lengths", "")
	registerFunc(&_e5rtTensorDescCreateWithAlignments, &_e5rtTensorDescCreateWithAlignmentsErr, frameworkHandle, "e5rt_tensor_desc_create_with_alignments", "")
	registerFunc(&_e5rtTensorDescCreateWithStrides, &_e5rtTensorDescCreateWithStridesErr, frameworkHandle, "e5rt_tensor_desc_create_with_strides", "")
	registerFunc(&_e5rtTensorDescDtypeAreEqual, &_e5rtTensorDescDtypeAreEqualErr, frameworkHandle, "e5rt_tensor_desc_dtype_are_equal", "")
	registerFunc(&_e5rtTensorDescDtypeCreate, &_e5rtTensorDescDtypeCreateErr, frameworkHandle, "e5rt_tensor_desc_dtype_create", "")
	registerFunc(&_e5rtTensorDescDtypeGetComponentDtype, &_e5rtTensorDescDtypeGetComponentDtypeErr, frameworkHandle, "e5rt_tensor_desc_dtype_get_component_dtype", "")
	registerFunc(&_e5rtTensorDescDtypeGetComponentPack, &_e5rtTensorDescDtypeGetComponentPackErr, frameworkHandle, "e5rt_tensor_desc_dtype_get_component_pack", "")
	registerFunc(&_e5rtTensorDescDtypeGetComponentSize, &_e5rtTensorDescDtypeGetComponentSizeErr, frameworkHandle, "e5rt_tensor_desc_dtype_get_component_size", "")
	registerFunc(&_e5rtTensorDescDtypeGetElementSize, &_e5rtTensorDescDtypeGetElementSizeErr, frameworkHandle, "e5rt_tensor_desc_dtype_get_element_size", "")
	registerFunc(&_e5rtTensorDescDtypeGetNumComponents, &_e5rtTensorDescDtypeGetNumComponentsErr, frameworkHandle, "e5rt_tensor_desc_dtype_get_num_components", "")
	registerFunc(&_e5rtTensorDescDtypeRelease, &_e5rtTensorDescDtypeReleaseErr, frameworkHandle, "e5rt_tensor_desc_dtype_release", "")
	registerFunc(&_e5rtTensorDescDtypeValidateSpec, &_e5rtTensorDescDtypeValidateSpecErr, frameworkHandle, "e5rt_tensor_desc_dtype_validate_spec", "")
	registerFunc(&_e5rtTensorDescGetByteOffset, &_e5rtTensorDescGetByteOffsetErr, frameworkHandle, "e5rt_tensor_desc_get_byte_offset", "")
	registerFunc(&_e5rtTensorDescGetDimensionLength, &_e5rtTensorDescGetDimensionLengthErr, frameworkHandle, "e5rt_tensor_desc_get_dimension_length", "")
	registerFunc(&_e5rtTensorDescGetDimensionStride, &_e5rtTensorDescGetDimensionStrideErr, frameworkHandle, "e5rt_tensor_desc_get_dimension_stride", "")
	registerFunc(&_e5rtTensorDescGetNumElements, &_e5rtTensorDescGetNumElementsErr, frameworkHandle, "e5rt_tensor_desc_get_num_elements", "")
	registerFunc(&_e5rtTensorDescGetRank, &_e5rtTensorDescGetRankErr, frameworkHandle, "e5rt_tensor_desc_get_rank", "")
	registerFunc(&_e5rtTensorDescGetShape, &_e5rtTensorDescGetShapeErr, frameworkHandle, "e5rt_tensor_desc_get_shape", "")
	registerFunc(&_e5rtTensorDescGetSize, &_e5rtTensorDescGetSizeErr, frameworkHandle, "e5rt_tensor_desc_get_size", "")
	registerFunc(&_e5rtTensorDescGetStrides, &_e5rtTensorDescGetStridesErr, frameworkHandle, "e5rt_tensor_desc_get_strides", "")
	registerFunc(&_e5rtTensorDescHasKnownShape, &_e5rtTensorDescHasKnownShapeErr, frameworkHandle, "e5rt_tensor_desc_has_known_shape", "")
	registerFunc(&_e5rtTensorDescRelease, &_e5rtTensorDescReleaseErr, frameworkHandle, "e5rt_tensor_desc_release", "")
	registerFunc(&_e5rtTensorDescRetainDtype, &_e5rtTensorDescRetainDtypeErr, frameworkHandle, "e5rt_tensor_desc_retain_dtype", "")
	registerFunc(&_e5rtTensorUtilsAreTensorsEqual, &_e5rtTensorUtilsAreTensorsEqualErr, frameworkHandle, "e5rt_tensor_utils_are_tensors_equal", "")
	registerFunc(&_e5rtTensorUtilsCastFromFp16ToFp32, &_e5rtTensorUtilsCastFromFp16ToFp32Err, frameworkHandle, "e5rt_tensor_utils_cast_from_fp16_to_fp32", "")
	registerFunc(&_e5rtTensorUtilsCastFromFp32ToFp16, &_e5rtTensorUtilsCastFromFp32ToFp16Err, frameworkHandle, "e5rt_tensor_utils_cast_from_fp32_to_fp16", "")
	registerFunc(&_e5rtTensorUtilsCopyTensor, &_e5rtTensorUtilsCopyTensorErr, frameworkHandle, "e5rt_tensor_utils_copy_tensor", "")
	registerFunc(&_e5rtTensorUtilsGetFp16Element, &_e5rtTensorUtilsGetFp16ElementErr, frameworkHandle, "e5rt_tensor_utils_get_fp16_element", "")
	registerFunc(&_e5rtTensorUtilsGetFp32Element, &_e5rtTensorUtilsGetFp32ElementErr, frameworkHandle, "e5rt_tensor_utils_get_fp32_element", "")
	registerFunc(&_e5rtTensorUtilsGetS8Element, &_e5rtTensorUtilsGetS8ElementErr, frameworkHandle, "e5rt_tensor_utils_get_s8_element", "")
	registerFunc(&_e5rtTensorUtilsGetU8Element, &_e5rtTensorUtilsGetU8ElementErr, frameworkHandle, "e5rt_tensor_utils_get_u8_element", "")
	registerFunc(&_e5rtTensorUtilsSetFp32Element, &_e5rtTensorUtilsSetFp32ElementErr, frameworkHandle, "e5rt_tensor_utils_set_fp32_element", "")
	registerFunc(&_e5rtTensorUtilsSetS8Element, &_e5rtTensorUtilsSetS8ElementErr, frameworkHandle, "e5rt_tensor_utils_set_s8_element", "")
	registerFunc(&_e5rtTensorUtilsSetU8Element, &_e5rtTensorUtilsSetU8ElementErr, frameworkHandle, "e5rt_tensor_utils_set_u8_element", "")
	registerFunc(&_espressoAneCacheHasNetwork, &_espressoAneCacheHasNetworkErr, frameworkHandle, "espresso_ane_cache_has_network", "")
	registerFunc(&_espressoAneCachePurgeNetwork, &_espressoAneCachePurgeNetworkErr, frameworkHandle, "espresso_ane_cache_purge_network", "")
	registerFunc(&_espressoBlobSetIntOption, &_espressoBlobSetIntOptionErr, frameworkHandle, "espresso_blob_set_int_option", "")
	registerFunc(&_espressoBufferGetCount, &_espressoBufferGetCountErr, frameworkHandle, "espresso_buffer_get_count", "")
	registerFunc(&_espressoBufferGetRank, &_espressoBufferGetRankErr, frameworkHandle, "espresso_buffer_get_rank", "")
	registerFunc(&_espressoBufferGetSize, &_espressoBufferGetSizeErr, frameworkHandle, "espresso_buffer_get_size", "")
	registerFunc(&_espressoBufferPackTensorShape, &_espressoBufferPackTensorShapeErr, frameworkHandle, "espresso_buffer_pack_tensor_shape", "")
	registerFunc(&_espressoBufferSetRank, &_espressoBufferSetRankErr, frameworkHandle, "espresso_buffer_set_rank", "")
	registerFunc(&_espressoBufferUnpackTensorShape, &_espressoBufferUnpackTensorShapeErr, frameworkHandle, "espresso_buffer_unpack_tensor_shape", "")
	registerFunc(&_espressoCompileMilToEir, &_espressoCompileMilToEirErr, frameworkHandle, "espresso_compile_mil_to_eir", "")
	registerFunc(&_espressoContextDestroy, &_espressoContextDestroyErr, frameworkHandle, "espresso_context_destroy", "")
	registerFunc(&_espressoContextReportBench, &_espressoContextReportBenchErr, frameworkHandle, "espresso_context_report_bench", "")
	registerFunc(&_espressoContextSetIntOption, &_espressoContextSetIntOptionErr, frameworkHandle, "espresso_context_set_int_option", "")
	registerFunc(&_espressoContextSetLowPrecisionAccumulation, &_espressoContextSetLowPrecisionAccumulationErr, frameworkHandle, "espresso_context_set_low_precision_accumulation", "")
	registerFunc(&_espressoCreateContext, &_espressoCreateContextErr, frameworkHandle, "espresso_create_context", "")
	registerFunc(&_espressoCreateContextAuto, &_espressoCreateContextAutoErr, frameworkHandle, "espresso_create_context_auto", "")
	registerFunc(&_espressoCreateContextWithArgs, &_espressoCreateContextWithArgsErr, frameworkHandle, "espresso_create_context_with_args", "")
	registerFunc(&_espressoCreatePlan, &_espressoCreatePlanErr, frameworkHandle, "espresso_create_plan", "")
	registerFunc(&_espressoCreatePlanAndLoadNetwork, &_espressoCreatePlanAndLoadNetworkErr, frameworkHandle, "espresso_create_plan_and_load_network", "")
	registerFunc(&_espressoDeviceIDForMetalDevice, &_espressoDeviceIDForMetalDeviceErr, frameworkHandle, "espresso_device_id_for_metal_device", "")
	registerFunc(&_espressoDumpIr, &_espressoDumpIrErr, frameworkHandle, "espresso_dump_ir", "")
	registerFunc(&_espressoEnableAutoinitialize, &_espressoEnableAutoinitializeErr, frameworkHandle, "espresso_enable_autoinitialize", "")
	registerFunc(&_espressoEnableTestVectorMode, &_espressoEnableTestVectorModeErr, frameworkHandle, "espresso_enable_test_vector_mode", "")
	registerFunc(&_espressoGenerateTrainingProgram, &_espressoGenerateTrainingProgramErr, frameworkHandle, "espresso_generate_training_program", "")
	registerFunc(&_espressoGetAnalysisModelMetadataForKey, &_espressoGetAnalysisModelMetadataForKeyErr, frameworkHandle, "espresso_get_analysis_model_metadata_for_key", "")
	registerFunc(&_espressoGetDefaultStorageType, &_espressoGetDefaultStorageTypeErr, frameworkHandle, "espresso_get_default_storage_type", "")
	registerFunc(&_espressoGetMetadataForKey, &_espressoGetMetadataForKeyErr, frameworkHandle, "espresso_get_metadata_for_key", "")
	registerFunc(&_espressoGetStatusString, &_espressoGetStatusStringErr, frameworkHandle, "espresso_get_status_string", "")
	registerFunc(&_espressoGetVersionString, &_espressoGetVersionStringErr, frameworkHandle, "espresso_get_version_string", "")
	registerFunc(&_espressoGPUPreferIntegrated, &_espressoGPUPreferIntegratedErr, frameworkHandle, "espresso_gpu_prefer_integrated", "")
	registerFunc(&_espressoIsAneArchGreaterThanOrEqual, &_espressoIsAneArchGreaterThanOrEqualErr, frameworkHandle, "espresso_is_ane_arch_greater_than_or_equal", "")
	registerFunc(&_espressoNetworkBindBuffer, &_espressoNetworkBindBufferErr, frameworkHandle, "espresso_network_bind_buffer", "")
	registerFunc(&_espressoNetworkBindCvpixelbuffer, &_espressoNetworkBindCvpixelbufferErr, frameworkHandle, "espresso_network_bind_cvpixelbuffer", "")
	registerFunc(&_espressoNetworkBindDirectCvpixelbuffer, &_espressoNetworkBindDirectCvpixelbufferErr, frameworkHandle, "espresso_network_bind_direct_cvpixelbuffer", "")
	registerFunc(&_espressoNetworkBindInputCvpixelbuffer, &_espressoNetworkBindInputCvpixelbufferErr, frameworkHandle, "espresso_network_bind_input_cvpixelbuffer", "")
	registerFunc(&_espressoNetworkBindInputMetaltexture, &_espressoNetworkBindInputMetaltextureErr, frameworkHandle, "espresso_network_bind_input_metaltexture", "")
	registerFunc(&_espressoNetworkBindInputVimagebufferArgb8, &_espressoNetworkBindInputVimagebufferArgb8Err, frameworkHandle, "espresso_network_bind_input_vimagebuffer_argb8", "")
	registerFunc(&_espressoNetworkBindInputVimagebufferBgra8, &_espressoNetworkBindInputVimagebufferBgra8Err, frameworkHandle, "espresso_network_bind_input_vimagebuffer_bgra8", "")
	registerFunc(&_espressoNetworkBindInputVimagebufferPlanar8, &_espressoNetworkBindInputVimagebufferPlanar8Err, frameworkHandle, "espresso_network_bind_input_vimagebuffer_planar8", "")
	registerFunc(&_espressoNetworkBindInputVimagebufferRgba8, &_espressoNetworkBindInputVimagebufferRgba8Err, frameworkHandle, "espresso_network_bind_input_vimagebuffer_rgba8", "")
	registerFunc(&_espressoNetworkChangeBlobShape, &_espressoNetworkChangeBlobShapeErr, frameworkHandle, "espresso_network_change_blob_shape", "")
	registerFunc(&_espressoNetworkChangeInputBlobShapes, &_espressoNetworkChangeInputBlobShapesErr, frameworkHandle, "espresso_network_change_input_blob_shapes", "")
	registerFunc(&_espressoNetworkDeclareInput, &_espressoNetworkDeclareInputErr, frameworkHandle, "espresso_network_declare_input", "")
	registerFunc(&_espressoNetworkDeclareOutput, &_espressoNetworkDeclareOutputErr, frameworkHandle, "espresso_network_declare_output", "")
	registerFunc(&_espressoNetworkDumpTestVector, &_espressoNetworkDumpTestVectorErr, frameworkHandle, "espresso_network_dump_test_vector", "")
	registerFunc(&_espressoNetworkGetVersion, &_espressoNetworkGetVersionErr, frameworkHandle, "espresso_network_get_version", "")
	registerFunc(&_espressoNetworkQueryBlobDimensions, &_espressoNetworkQueryBlobDimensionsErr, frameworkHandle, "espresso_network_query_blob_dimensions", "")
	registerFunc(&_espressoNetworkQueryBlobShape, &_espressoNetworkQueryBlobShapeErr, frameworkHandle, "espresso_network_query_blob_shape", "")
	registerFunc(&_espressoNetworkSetMemoryPoolID, &_espressoNetworkSetMemoryPoolIDErr, frameworkHandle, "espresso_network_set_memory_pool_id", "")
	registerFunc(&_espressoNetworkSetTracingName, &_espressoNetworkSetTracingNameErr, frameworkHandle, "espresso_network_set_tracing_name", "")
	registerFunc(&_espressoNetworkSwapGlobal, &_espressoNetworkSwapGlobalErr, frameworkHandle, "espresso_network_swap_global", "")
	registerFunc(&_espressoNetworkSyncCopyGlobal, &_espressoNetworkSyncCopyGlobalErr, frameworkHandle, "espresso_network_sync_copy_global", "")
	registerFunc(&_espressoNetworkTemporalStateReset, &_espressoNetworkTemporalStateResetErr, frameworkHandle, "espresso_network_temporal_state_reset", "")
	registerFunc(&_espressoNetworkUnbindBuffer, &_espressoNetworkUnbindBufferErr, frameworkHandle, "espresso_network_unbind_buffer", "")
	registerFunc(&_espressoPlanAddNetwork, &_espressoPlanAddNetworkErr, frameworkHandle, "espresso_plan_add_network", "")
	registerFunc(&_espressoPlanAutoProfile, &_espressoPlanAutoProfileErr, frameworkHandle, "espresso_plan_auto_profile", "")
	registerFunc(&_espressoPlanBuild, &_espressoPlanBuildErr, frameworkHandle, "espresso_plan_build", "")
	registerFunc(&_espressoPlanBuildClean, &_espressoPlanBuildCleanErr, frameworkHandle, "espresso_plan_build_clean", "")
	registerFunc(&_espressoPlanBuildWithOptions, &_espressoPlanBuildWithOptionsErr, frameworkHandle, "espresso_plan_build_with_options", "")
	registerFunc(&_espressoPlanCanUseSubmit, &_espressoPlanCanUseSubmitErr, frameworkHandle, "espresso_plan_can_use_submit", "")
	registerFunc(&_espressoPlanDestroy, &_espressoPlanDestroyErr, frameworkHandle, "espresso_plan_destroy", "")
	registerFunc(&_espressoPlanExecuteSync, &_espressoPlanExecuteSyncErr, frameworkHandle, "espresso_plan_execute_sync", "")
	registerFunc(&_espressoPlanFinishProfiling, &_espressoPlanFinishProfilingErr, frameworkHandle, "espresso_plan_finish_profiling", "")
	registerFunc(&_espressoPlanGetErrorInfo, &_espressoPlanGetErrorInfoErr, frameworkHandle, "espresso_plan_get_error_info", "")
	registerFunc(&_espressoPlanGetPhase, &_espressoPlanGetPhaseErr, frameworkHandle, "espresso_plan_get_phase", "")
	registerFunc(&_espressoPlanSetExecutionQueue, &_espressoPlanSetExecutionQueueErr, frameworkHandle, "espresso_plan_set_execution_queue", "")
	registerFunc(&_espressoPlanSetPriority, &_espressoPlanSetPriorityErr, frameworkHandle, "espresso_plan_set_priority", "")
	registerFunc(&_espressoPlanShareIntermediateBuffer, &_espressoPlanShareIntermediateBufferErr, frameworkHandle, "espresso_plan_share_intermediate_buffer", "")
	registerFunc(&_espressoPlanStartProfiling, &_espressoPlanStartProfilingErr, frameworkHandle, "espresso_plan_start_profiling", "")
	registerFunc(&_espressoPlanStartProfilingWithOptions, &_espressoPlanStartProfilingWithOptionsErr, frameworkHandle, "espresso_plan_start_profiling_with_options", "")
	registerFunc(&_espressoPlanStaticProfilingInfo, &_espressoPlanStaticProfilingInfoErr, frameworkHandle, "espresso_plan_static_profiling_info", "")
	registerFunc(&_espressoPlanSubmit, &_espressoPlanSubmitErr, frameworkHandle, "espresso_plan_submit", "")
	registerFunc(&_espressoPlanSubmitCamera, &_espressoPlanSubmitCameraErr, frameworkHandle, "espresso_plan_submit_camera", "")
	registerFunc(&_espressoPlanSubmitSetMultipleBuffering, &_espressoPlanSubmitSetMultipleBufferingErr, frameworkHandle, "espresso_plan_submit_set_multiple_buffering", "")
	registerFunc(&_espressoPlanSubmitWithArgs, &_espressoPlanSubmitWithArgsErr, frameworkHandle, "espresso_plan_submit_with_args", "")
	registerFunc(&_espressoPlanWipeTempoary, &_espressoPlanWipeTempoaryErr, frameworkHandle, "espresso_plan_wipe_tempoary", "")
	registerFunc(&_espressoSetAnalysisModelMetadataForKey, &_espressoSetAnalysisModelMetadataForKeyErr, frameworkHandle, "espresso_set_analysis_model_metadata_for_key", "")
	registerFunc(&_espressoSetCompilationPlatform, &_espressoSetCompilationPlatformErr, frameworkHandle, "espresso_set_compilation_platform", "")
	registerFunc(&_espressoSetMilConstValues, &_espressoSetMilConstValuesErr, frameworkHandle, "espresso_set_mil_const_values", "")
	registerFunc(&_espressoTmpEnableMontreal, &_espressoTmpEnableMontrealErr, frameworkHandle, "espresso_tmp_enable_montreal", "")
	registerFunc(&_espressoUpgradeEirToMil, &_espressoUpgradeEirToMilErr, frameworkHandle, "espresso_upgrade_eir_to_mil", "")
	registerFunc(&_espressoUpgradeNetToMil, &_espressoUpgradeNetToMilErr, frameworkHandle, "espresso_upgrade_net_to_mil", "")
	registerFunc(&_espressoUpgradeNetToMilProgram, &_espressoUpgradeNetToMilProgramErr, frameworkHandle, "espresso_upgrade_net_to_mil_program", "")
	registerFunc(&_espressoUpgradeToMil, &_espressoUpgradeToMilErr, frameworkHandle, "espresso_upgrade_to_mil", "")
}

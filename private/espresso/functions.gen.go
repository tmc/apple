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

var _e5rtAneMemoryProviderCreate func(out *uintptr, a1 uintptr) int32
var _e5rtAneMemoryProviderCreateErr error

func tryE5rtAneMemoryProviderCreate(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtAneMemoryProviderCreate == nil {
		return 0, symbolCallError("e5rt_ane_memory_provider_create", "", _e5rtAneMemoryProviderCreateErr)
	}
	return _e5rtAneMemoryProviderCreate(out, a1), nil
}

// E5rtAneMemoryProviderCreate signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::ANEMemoryProvider::Create takes 1.
func E5rtAneMemoryProviderCreate(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtAneMemoryProviderCreate(out, a1)
}

var _e5rtAneMemoryProviderRelease func(a0 uintptr) int32
var _e5rtAneMemoryProviderReleaseErr error

func tryE5rtAneMemoryProviderRelease(a0 uintptr) (int32, error) {
	if _e5rtAneMemoryProviderRelease == nil {
		return 0, symbolCallError("e5rt_ane_memory_provider_release", "", _e5rtAneMemoryProviderReleaseErr)
	}
	return _e5rtAneMemoryProviderRelease(a0), nil
}

// E5rtAneMemoryProviderRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtAneMemoryProviderRelease(a0 uintptr) (int32, error) {
	return tryE5rtAneMemoryProviderRelease(a0)
}

var _e5rtAsyncEventAsyncNotify func(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtAsyncEventAsyncNotifyErr error

func tryE5rtAsyncEventAsyncNotify(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtAsyncEventAsyncNotify == nil {
		return 0, symbolCallError("e5rt_async_event_async_notify", "", _e5rtAsyncEventAsyncNotifyErr)
	}
	return _e5rtAsyncEventAsyncNotify(out, a1, a2, a3), nil
}

// E5rtAsyncEventAsyncNotify signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 4, E5RT::AsyncEvent::AsyncNotify takes 3.
func E5rtAsyncEventAsyncNotify(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtAsyncEventAsyncNotify(out, a1, a2, a3)
}

var _e5rtAsyncEventCreate func(out *uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtAsyncEventCreateErr error

func tryE5rtAsyncEventCreate(out *uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtAsyncEventCreate == nil {
		return 0, symbolCallError("e5rt_async_event_create", "", _e5rtAsyncEventCreateErr)
	}
	return _e5rtAsyncEventCreate(out, a1, a2), nil
}

// E5rtAsyncEventCreate signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: wraps a constructor, so the out-parameter is the object being made.
func E5rtAsyncEventCreate(out *uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtAsyncEventCreate(out, a1, a2)
}

var _e5rtAsyncEventCreateFromIosurfaceSharedEvent func(out *uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtAsyncEventCreateFromIosurfaceSharedEventErr error

func tryE5rtAsyncEventCreateFromIosurfaceSharedEvent(out *uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtAsyncEventCreateFromIosurfaceSharedEvent == nil {
		return 0, symbolCallError("e5rt_async_event_create_from_iosurface_shared_event", "", _e5rtAsyncEventCreateFromIosurfaceSharedEventErr)
	}
	return _e5rtAsyncEventCreateFromIosurfaceSharedEvent(out, a1, a2), nil
}

// E5rtAsyncEventCreateFromIosurfaceSharedEvent signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: wraps a constructor, so the out-parameter is the object being made.
func E5rtAsyncEventCreateFromIosurfaceSharedEvent(out *uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtAsyncEventCreateFromIosurfaceSharedEvent(out, a1, a2)
}

var _e5rtAsyncEventGetActiveFutureValue func(a0 uintptr, out *uintptr) int32
var _e5rtAsyncEventGetActiveFutureValueErr error

func tryE5rtAsyncEventGetActiveFutureValue(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtAsyncEventGetActiveFutureValue == nil {
		return 0, symbolCallError("e5rt_async_event_get_active_future_value", "", _e5rtAsyncEventGetActiveFutureValueErr)
	}
	return _e5rtAsyncEventGetActiveFutureValue(a0, out), nil
}

// E5rtAsyncEventGetActiveFutureValue signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::AsyncEvent::GetActiveFutureValue takes 0.
func E5rtAsyncEventGetActiveFutureValue(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtAsyncEventGetActiveFutureValue(a0, out)
}

var _e5rtAsyncEventGetIosurfaceSharedEvent func(a0 uintptr, a1 uintptr) int32
var _e5rtAsyncEventGetIosurfaceSharedEventErr error

func tryE5rtAsyncEventGetIosurfaceSharedEvent(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtAsyncEventGetIosurfaceSharedEvent == nil {
		return 0, symbolCallError("e5rt_async_event_get_iosurface_shared_event", "", _e5rtAsyncEventGetIosurfaceSharedEventErr)
	}
	return _e5rtAsyncEventGetIosurfaceSharedEvent(a0, a1), nil
}

// E5rtAsyncEventGetIosurfaceSharedEvent signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtAsyncEventGetIosurfaceSharedEvent(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtAsyncEventGetIosurfaceSharedEvent(a0, a1)
}

var _e5rtAsyncEventGetLastSignaledValue func(a0 uintptr, out *uintptr) int32
var _e5rtAsyncEventGetLastSignaledValueErr error

func tryE5rtAsyncEventGetLastSignaledValue(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtAsyncEventGetLastSignaledValue == nil {
		return 0, symbolCallError("e5rt_async_event_get_last_signaled_value", "", _e5rtAsyncEventGetLastSignaledValueErr)
	}
	return _e5rtAsyncEventGetLastSignaledValue(a0, out), nil
}

// E5rtAsyncEventGetLastSignaledValue signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::AsyncEvent::GetLastSignaledValue takes 0.
func E5rtAsyncEventGetLastSignaledValue(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtAsyncEventGetLastSignaledValue(a0, out)
}

var _e5rtAsyncEventGetName func(a0 uintptr, a1 uintptr) int32
var _e5rtAsyncEventGetNameErr error

func tryE5rtAsyncEventGetName(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtAsyncEventGetName == nil {
		return 0, symbolCallError("e5rt_async_event_get_name", "", _e5rtAsyncEventGetNameErr)
	}
	return _e5rtAsyncEventGetName(a0, a1), nil
}

// E5rtAsyncEventGetName signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtAsyncEventGetName(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtAsyncEventGetName(a0, a1)
}

var _e5rtAsyncEventRelease func(a0 uintptr) int32
var _e5rtAsyncEventReleaseErr error

func tryE5rtAsyncEventRelease(a0 uintptr) (int32, error) {
	if _e5rtAsyncEventRelease == nil {
		return 0, symbolCallError("e5rt_async_event_release", "", _e5rtAsyncEventReleaseErr)
	}
	return _e5rtAsyncEventRelease(a0), nil
}

// E5rtAsyncEventRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtAsyncEventRelease(a0 uintptr) (int32, error) {
	return tryE5rtAsyncEventRelease(a0)
}

var _e5rtAsyncEventSetActiveFutureValue func(out *uintptr, a1 uintptr) int32
var _e5rtAsyncEventSetActiveFutureValueErr error

func tryE5rtAsyncEventSetActiveFutureValue(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtAsyncEventSetActiveFutureValue == nil {
		return 0, symbolCallError("e5rt_async_event_set_active_future_value", "", _e5rtAsyncEventSetActiveFutureValueErr)
	}
	return _e5rtAsyncEventSetActiveFutureValue(out, a1), nil
}

// E5rtAsyncEventSetActiveFutureValue signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::AsyncEvent::SetActiveFutureValue takes 1.
func E5rtAsyncEventSetActiveFutureValue(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtAsyncEventSetActiveFutureValue(out, a1)
}

var _e5rtAsyncEventSignal func(out *uintptr, a1 uintptr) int32
var _e5rtAsyncEventSignalErr error

func tryE5rtAsyncEventSignal(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtAsyncEventSignal == nil {
		return 0, symbolCallError("e5rt_async_event_signal", "", _e5rtAsyncEventSignalErr)
	}
	return _e5rtAsyncEventSignal(out, a1), nil
}

// E5rtAsyncEventSignal signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::AsyncEvent::Signal takes 1.
func E5rtAsyncEventSignal(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtAsyncEventSignal(out, a1)
}

var _e5rtAsyncEventSyncWait func(out *uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtAsyncEventSyncWaitErr error

func tryE5rtAsyncEventSyncWait(out *uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtAsyncEventSyncWait == nil {
		return 0, symbolCallError("e5rt_async_event_sync_wait", "", _e5rtAsyncEventSyncWaitErr)
	}
	return _e5rtAsyncEventSyncWait(out, a1, a2), nil
}

// E5rtAsyncEventSyncWait signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 3, E5RT::AsyncEvent::SyncWait takes 2.
func E5rtAsyncEventSyncWait(out *uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtAsyncEventSyncWait(out, a1, a2)
}

var _e5rtBufferObjectAlloc func(out *uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtBufferObjectAllocErr error

func tryE5rtBufferObjectAlloc(out *uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtBufferObjectAlloc == nil {
		return 0, symbolCallError("e5rt_buffer_object_alloc", "", _e5rtBufferObjectAllocErr)
	}
	return _e5rtBufferObjectAlloc(out, a1, a2), nil
}

// E5rtBufferObjectAlloc signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:567:9, where argument 0 is the out-parameter.
func E5rtBufferObjectAlloc(out *uintptr, a1 uintptr, a2 uintptr) (int32, error) {
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

// E5rtBufferObjectCreateAsAlias signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: wraps a constructor, so the out-parameter is the object being made.
func E5rtBufferObjectCreateAsAlias(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtBufferObjectCreateAsAlias(out, a1, a2, a3)
}

var _e5rtBufferObjectCreateFromDataPointer func(out *uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtBufferObjectCreateFromDataPointerErr error

func tryE5rtBufferObjectCreateFromDataPointer(out *uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtBufferObjectCreateFromDataPointer == nil {
		return 0, symbolCallError("e5rt_buffer_object_create_from_data_pointer", "", _e5rtBufferObjectCreateFromDataPointerErr)
	}
	return _e5rtBufferObjectCreateFromDataPointer(out, a1, a2), nil
}

// E5rtBufferObjectCreateFromDataPointer signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: wraps a constructor, so the out-parameter is the object being made.
func E5rtBufferObjectCreateFromDataPointer(out *uintptr, a1 uintptr, a2 uintptr) (int32, error) {
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

// E5rtBufferObjectCreateFromIosurface signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: wraps a constructor, so the out-parameter is the object being made.
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

// E5rtBufferObjectCreateFromMtlbuffer signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: wraps a constructor, so the out-parameter is the object being made.
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

// E5rtBufferObjectGetDataPtr signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:580:9, where argument 1 is the out-parameter.
func E5rtBufferObjectGetDataPtr(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtBufferObjectGetDataPtr(a0, out)
}

var _e5rtBufferObjectGetIosurface func(a0 uintptr, a1 uintptr) int32
var _e5rtBufferObjectGetIosurfaceErr error

func tryE5rtBufferObjectGetIosurface(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtBufferObjectGetIosurface == nil {
		return 0, symbolCallError("e5rt_buffer_object_get_iosurface", "", _e5rtBufferObjectGetIosurfaceErr)
	}
	return _e5rtBufferObjectGetIosurface(a0, a1), nil
}

// E5rtBufferObjectGetIosurface signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtBufferObjectGetIosurface(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtBufferObjectGetIosurface(a0, a1)
}

var _e5rtBufferObjectGetMtlbuffer func(a0 uintptr, a1 uintptr) int32
var _e5rtBufferObjectGetMtlbufferErr error

func tryE5rtBufferObjectGetMtlbuffer(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtBufferObjectGetMtlbuffer == nil {
		return 0, symbolCallError("e5rt_buffer_object_get_mtlbuffer", "", _e5rtBufferObjectGetMtlbufferErr)
	}
	return _e5rtBufferObjectGetMtlbuffer(a0, a1), nil
}

// E5rtBufferObjectGetMtlbuffer signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtBufferObjectGetMtlbuffer(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtBufferObjectGetMtlbuffer(a0, a1)
}

var _e5rtBufferObjectGetSize func(a0 uintptr, out *uintptr) int32
var _e5rtBufferObjectGetSizeErr error

func tryE5rtBufferObjectGetSize(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtBufferObjectGetSize == nil {
		return 0, symbolCallError("e5rt_buffer_object_get_size", "", _e5rtBufferObjectGetSizeErr)
	}
	return _e5rtBufferObjectGetSize(a0, out), nil
}

// E5rtBufferObjectGetSize signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::BufferObject::GetSize takes 0.
func E5rtBufferObjectGetSize(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtBufferObjectGetSize(a0, out)
}

var _e5rtBufferObjectGetType func(a0 uintptr, a1 uintptr) int32
var _e5rtBufferObjectGetTypeErr error

func tryE5rtBufferObjectGetType(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtBufferObjectGetType == nil {
		return 0, symbolCallError("e5rt_buffer_object_get_type", "", _e5rtBufferObjectGetTypeErr)
	}
	return _e5rtBufferObjectGetType(a0, a1), nil
}

// E5rtBufferObjectGetType signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtBufferObjectGetType(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtBufferObjectGetType(a0, a1)
}

var _e5rtBufferObjectRelease func(a0 uintptr) int32
var _e5rtBufferObjectReleaseErr error

func tryE5rtBufferObjectRelease(a0 uintptr) (int32, error) {
	if _e5rtBufferObjectRelease == nil {
		return 0, symbolCallError("e5rt_buffer_object_release", "", _e5rtBufferObjectReleaseErr)
	}
	return _e5rtBufferObjectRelease(a0), nil
}

// E5rtBufferObjectRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtBufferObjectRelease(a0 uintptr) (int32, error) {
	return tryE5rtBufferObjectRelease(a0)
}

var _e5rtComputeGPUDeviceGetMtlDevice func(a0 uintptr, out *uintptr) int32
var _e5rtComputeGPUDeviceGetMtlDeviceErr error

func tryE5rtComputeGPUDeviceGetMtlDevice(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtComputeGPUDeviceGetMtlDevice == nil {
		return 0, symbolCallError("e5rt_compute_gpu_device_get_mtl_device", "", _e5rtComputeGPUDeviceGetMtlDeviceErr)
	}
	return _e5rtComputeGPUDeviceGetMtlDevice(a0, out), nil
}

// E5rtComputeGPUDeviceGetMtlDevice signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::ComputeGPUDevice::GetMTLDevice takes 0.
func E5rtComputeGPUDeviceGetMtlDevice(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtComputeGPUDeviceGetMtlDevice(a0, out)
}

var _e5rtComputeGPUDeviceRelease func(a0 uintptr) int32
var _e5rtComputeGPUDeviceReleaseErr error

func tryE5rtComputeGPUDeviceRelease(a0 uintptr) (int32, error) {
	if _e5rtComputeGPUDeviceRelease == nil {
		return 0, symbolCallError("e5rt_compute_gpu_device_release", "", _e5rtComputeGPUDeviceReleaseErr)
	}
	return _e5rtComputeGPUDeviceRelease(a0), nil
}

// E5rtComputeGPUDeviceRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtComputeGPUDeviceRelease(a0 uintptr) (int32, error) {
	return tryE5rtComputeGPUDeviceRelease(a0)
}

var _e5rtComputeGPUDeviceRetainAll func(a0 uintptr, a1 uintptr) int32
var _e5rtComputeGPUDeviceRetainAllErr error

func tryE5rtComputeGPUDeviceRetainAll(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtComputeGPUDeviceRetainAll == nil {
		return 0, symbolCallError("e5rt_compute_gpu_device_retain_all", "", _e5rtComputeGPUDeviceRetainAllErr)
	}
	return _e5rtComputeGPUDeviceRetainAll(a0, a1), nil
}

// E5rtComputeGPUDeviceRetainAll signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtComputeGPUDeviceRetainAll(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtComputeGPUDeviceRetainAll(a0, a1)
}

var _e5rtComputeGPUDeviceRetainFromMtlDevice func(a0 uintptr, a1 uintptr) int32
var _e5rtComputeGPUDeviceRetainFromMtlDeviceErr error

func tryE5rtComputeGPUDeviceRetainFromMtlDevice(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtComputeGPUDeviceRetainFromMtlDevice == nil {
		return 0, symbolCallError("e5rt_compute_gpu_device_retain_from_mtl_device", "", _e5rtComputeGPUDeviceRetainFromMtlDeviceErr)
	}
	return _e5rtComputeGPUDeviceRetainFromMtlDevice(a0, a1), nil
}

// E5rtComputeGPUDeviceRetainFromMtlDevice signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtComputeGPUDeviceRetainFromMtlDevice(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtComputeGPUDeviceRetainFromMtlDevice(a0, a1)
}

var _e5rtCreateSurfaceObjectFromIosurface func(a0 uintptr, a1 uintptr) int32
var _e5rtCreateSurfaceObjectFromIosurfaceErr error

func tryE5rtCreateSurfaceObjectFromIosurface(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtCreateSurfaceObjectFromIosurface == nil {
		return 0, symbolCallError("e5rt_create_surface_object_from_iosurface", "", _e5rtCreateSurfaceObjectFromIosurfaceErr)
	}
	return _e5rtCreateSurfaceObjectFromIosurface(a0, a1), nil
}

// E5rtCreateSurfaceObjectFromIosurface signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtCreateSurfaceObjectFromIosurface(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtCreateSurfaceObjectFromIosurface(a0, a1)
}

var _e5rtCvpb4ccToSurfaceFormat func(a0 uintptr, a1 uintptr) int32
var _e5rtCvpb4ccToSurfaceFormatErr error

func tryE5rtCvpb4ccToSurfaceFormat(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtCvpb4ccToSurfaceFormat == nil {
		return 0, symbolCallError("e5rt_cvpb_4cc_to_surface_format", "", _e5rtCvpb4ccToSurfaceFormatErr)
	}
	return _e5rtCvpb4ccToSurfaceFormat(a0, a1), nil
}

// E5rtCvpb4ccToSurfaceFormat signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtCvpb4ccToSurfaceFormat(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtCvpb4ccToSurfaceFormat(a0, a1)
}

var _e5rtE5CompilerCompile func(a0 uintptr, a1 uintptr, a2 uintptr, out *uintptr) int32
var _e5rtE5CompilerCompileErr error

func tryE5rtE5CompilerCompile(a0 uintptr, a1 uintptr, a2 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerCompile == nil {
		return 0, symbolCallError("e5rt_e5_compiler_compile", "", _e5rtE5CompilerCompileErr)
	}
	return _e5rtE5CompilerCompile(a0, a1, a2, out), nil
}

// E5rtE5CompilerCompile signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:425:9, where argument 3 is the out-parameter.
func E5rtE5CompilerCompile(a0 uintptr, a1 uintptr, a2 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerCompile(a0, a1, a2, out)
}

var _e5rtE5CompilerCompileFromIrProgram func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtE5CompilerCompileFromIrProgramErr error

func tryE5rtE5CompilerCompileFromIrProgram(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtE5CompilerCompileFromIrProgram == nil {
		return 0, symbolCallError("e5rt_e5_compiler_compile_from_ir_program", "", _e5rtE5CompilerCompileFromIrProgramErr)
	}
	return _e5rtE5CompilerCompileFromIrProgram(a0, a1, a2, a3), nil
}

// E5rtE5CompilerCompileFromIrProgram signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtE5CompilerCompileFromIrProgram(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtE5CompilerCompileFromIrProgram(a0, a1, a2, a3)
}

var _e5rtE5CompilerConfigOptionsCreate func(out *uintptr) int32
var _e5rtE5CompilerConfigOptionsCreateErr error

func tryE5rtE5CompilerConfigOptionsCreate(out *uintptr) (int32, error) {
	if _e5rtE5CompilerConfigOptionsCreate == nil {
		return 0, symbolCallError("e5rt_e5_compiler_config_options_create", "", _e5rtE5CompilerConfigOptionsCreateErr)
	}
	return _e5rtE5CompilerConfigOptionsCreate(out), nil
}

// E5rtE5CompilerConfigOptionsCreate signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:303:9, where argument 0 is the out-parameter.
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

// E5rtE5CompilerConfigOptionsGetBundleCacheApfsPurgeable signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtE5CompilerConfigOptionsGetCacheBundleLocation signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtE5CompilerConfigOptionsGetCacheBundleLocation(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerConfigOptionsGetCacheBundleLocation(a0, a1)
}

var _e5rtE5CompilerConfigOptionsRelease func(a0 uintptr) int32
var _e5rtE5CompilerConfigOptionsReleaseErr error

func tryE5rtE5CompilerConfigOptionsRelease(a0 uintptr) (int32, error) {
	if _e5rtE5CompilerConfigOptionsRelease == nil {
		return 0, symbolCallError("e5rt_e5_compiler_config_options_release", "", _e5rtE5CompilerConfigOptionsReleaseErr)
	}
	return _e5rtE5CompilerConfigOptionsRelease(a0), nil
}

// E5rtE5CompilerConfigOptionsRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtE5CompilerConfigOptionsRelease(a0 uintptr) (int32, error) {
	return tryE5rtE5CompilerConfigOptionsRelease(a0)
}

var _e5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeable func(a0 uintptr, a1 uintptr) int32
var _e5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeableErr error

func tryE5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeable(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeable == nil {
		return 0, symbolCallError("e5rt_e5_compiler_config_options_set_bundle_cache_apfs_purgeable", "", _e5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeableErr)
	}
	return _e5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeable(a0, a1), nil
}

// E5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeable signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeable(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerConfigOptionsSetBundleCacheApfsPurgeable(a0, a1)
}

var _e5rtE5CompilerConfigOptionsSetCacheBundleLocation func(a0 uintptr, a1 uintptr) int32
var _e5rtE5CompilerConfigOptionsSetCacheBundleLocationErr error

func tryE5rtE5CompilerConfigOptionsSetCacheBundleLocation(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerConfigOptionsSetCacheBundleLocation == nil {
		return 0, symbolCallError("e5rt_e5_compiler_config_options_set_cache_bundle_location", "", _e5rtE5CompilerConfigOptionsSetCacheBundleLocationErr)
	}
	return _e5rtE5CompilerConfigOptionsSetCacheBundleLocation(a0, a1), nil
}

// E5rtE5CompilerConfigOptionsSetCacheBundleLocation signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:314:9.
func E5rtE5CompilerConfigOptionsSetCacheBundleLocation(a0 uintptr, a1 uintptr) (int32, error) {
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

// E5rtE5CompilerCreate signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: wraps a constructor, so the out-parameter is the object being made.
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

// E5rtE5CompilerCreateWithConfig signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:337:9, where argument 0 is the out-parameter.
func E5rtE5CompilerCreateWithConfig(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerCreateWithConfig(out, a1)
}

var _e5rtE5CompilerIsNewCompileRequired func(a0 uintptr, a1 uintptr, a2 uintptr, out *uintptr) int32
var _e5rtE5CompilerIsNewCompileRequiredErr error

func tryE5rtE5CompilerIsNewCompileRequired(a0 uintptr, a1 uintptr, a2 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerIsNewCompileRequired == nil {
		return 0, symbolCallError("e5rt_e5_compiler_is_new_compile_required", "", _e5rtE5CompilerIsNewCompileRequiredErr)
	}
	return _e5rtE5CompilerIsNewCompileRequired(a0, a1, a2, out), nil
}

// E5rtE5CompilerIsNewCompileRequired signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 4, E5RT::E5Compiler::IsNewCompileRequired takes 2.
func E5rtE5CompilerIsNewCompileRequired(a0 uintptr, a1 uintptr, a2 uintptr, out *uintptr) (int32, error) {
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

// E5rtE5CompilerOptionsCreate signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:353:9, where argument 0 is the out-parameter.
func E5rtE5CompilerOptionsCreate(out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsCreate(out)
}

var _e5rtE5CompilerOptionsGetComputeDeviceTypesMask func(a0 uintptr, out *uintptr) int32
var _e5rtE5CompilerOptionsGetComputeDeviceTypesMaskErr error

func tryE5rtE5CompilerOptionsGetComputeDeviceTypesMask(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsGetComputeDeviceTypesMask == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_get_compute_device_types_mask", "", _e5rtE5CompilerOptionsGetComputeDeviceTypesMaskErr)
	}
	return _e5rtE5CompilerOptionsGetComputeDeviceTypesMask(a0, out), nil
}

// E5rtE5CompilerOptionsGetComputeDeviceTypesMask signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:373:9, where argument 1 is the out-parameter.
func E5rtE5CompilerOptionsGetComputeDeviceTypesMask(a0 uintptr, out *uintptr) (int32, error) {
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

// E5rtE5CompilerOptionsGetCreateProtectedAssets signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::GetCreateProtectedAssets takes 0.
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

// E5rtE5CompilerOptionsGetCustomAneCompilerOptions signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::GetCustomAneCompilerOptions takes 0.
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

// E5rtE5CompilerOptionsGetEnableMpsgraphPackage signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtE5CompilerOptionsGetEnableProfiling signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::GetEnableProfiling takes 0.
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

// E5rtE5CompilerOptionsGetEnableReshapeWithMinimalAllocations signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::GetEnableReshapeWithMinimalAllocations takes 0.
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

// E5rtE5CompilerOptionsGetExperimentalDisableCompileTimeMpsgraphTypeInference signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtE5CompilerOptionsGetExperimentalDisableDataDependentShape signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::GetExperimentalDisableDataDependentShape takes 0.
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

// E5rtE5CompilerOptionsGetExperimentalEnableDefaultFunctionForRangeDim signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::GetExperimentalEnableDefaultFunctionForRangeDim takes 0.
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

// E5rtE5CompilerOptionsGetExperimentalForceClassicCPUBackend signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::GetExperimentalForceClassicCpuBackend takes 0.
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

// E5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatterns signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::GetExperimentalMatchE5MinimalCpuPatterns takes 0.
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

// E5rtE5CompilerOptionsGetExperimentalMatchE5MinimalCPUPatternsForStates signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::GetExperimentalMatchE5MinimalCpuPatternsForStates takes 0.
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

// E5rtE5CompilerOptionsGetForceBnnsGraph signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtE5CompilerOptionsGetForceClassicAotOldHw signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::GetForceClassicAotOldHw takes 0.
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

// E5rtE5CompilerOptionsGetForceFetchFromCache signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::GetForceFetchFromCache takes 0.
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

// E5rtE5CompilerOptionsGetForceRecompilation signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::GetForceRecompilation takes 0.
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

// E5rtE5CompilerOptionsGetPreferredCPUBackend signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::GetPreferredCpuBackend takes 0.
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

// E5rtE5CompilerOptionsGetPreferredCPUBackends signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtE5CompilerOptionsGetSegmenter signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::GetSegmenter takes 0.
func E5rtE5CompilerOptionsGetSegmenter(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsGetSegmenter(a0, out)
}

var _e5rtE5CompilerOptionsRelease func(a0 uintptr) int32
var _e5rtE5CompilerOptionsReleaseErr error

func tryE5rtE5CompilerOptionsRelease(a0 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsRelease == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_release", "", _e5rtE5CompilerOptionsReleaseErr)
	}
	return _e5rtE5CompilerOptionsRelease(a0), nil
}

// E5rtE5CompilerOptionsRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtE5CompilerOptionsRelease(a0 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsRelease(a0)
}

var _e5rtE5CompilerOptionsRetainMilEntryPoints func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtE5CompilerOptionsRetainMilEntryPointsErr error

func tryE5rtE5CompilerOptionsRetainMilEntryPoints(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsRetainMilEntryPoints == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_retain_mil_entry_points", "", _e5rtE5CompilerOptionsRetainMilEntryPointsErr)
	}
	return _e5rtE5CompilerOptionsRetainMilEntryPoints(a0, a1, a2), nil
}

// E5rtE5CompilerOptionsRetainMilEntryPoints signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtE5CompilerOptionsRetainMilEntryPoints(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsRetainMilEntryPoints(a0, a1, a2)
}

var _e5rtE5CompilerOptionsSetComputeDeviceTypesMask func(a0 uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsSetComputeDeviceTypesMaskErr error

func tryE5rtE5CompilerOptionsSetComputeDeviceTypesMask(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetComputeDeviceTypesMask == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_compute_device_types_mask", "", _e5rtE5CompilerOptionsSetComputeDeviceTypesMaskErr)
	}
	return _e5rtE5CompilerOptionsSetComputeDeviceTypesMask(a0, a1), nil
}

// E5rtE5CompilerOptionsSetComputeDeviceTypesMask signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:364:9.
func E5rtE5CompilerOptionsSetComputeDeviceTypesMask(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetComputeDeviceTypesMask(a0, a1)
}

var _e5rtE5CompilerOptionsSetCreateProtectedAssets func(out *uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsSetCreateProtectedAssetsErr error

func tryE5rtE5CompilerOptionsSetCreateProtectedAssets(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetCreateProtectedAssets == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_create_protected_assets", "", _e5rtE5CompilerOptionsSetCreateProtectedAssetsErr)
	}
	return _e5rtE5CompilerOptionsSetCreateProtectedAssets(out, a1), nil
}

// E5rtE5CompilerOptionsSetCreateProtectedAssets signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::SetCreateProtectedAssets takes 1.
func E5rtE5CompilerOptionsSetCreateProtectedAssets(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetCreateProtectedAssets(out, a1)
}

var _e5rtE5CompilerOptionsSetCustomAneCompilerOptions func(a0 uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsSetCustomAneCompilerOptionsErr error

func tryE5rtE5CompilerOptionsSetCustomAneCompilerOptions(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetCustomAneCompilerOptions == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_custom_ane_compiler_options", "", _e5rtE5CompilerOptionsSetCustomAneCompilerOptionsErr)
	}
	return _e5rtE5CompilerOptionsSetCustomAneCompilerOptions(a0, a1), nil
}

// E5rtE5CompilerOptionsSetCustomAneCompilerOptions signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:399:9.
func E5rtE5CompilerOptionsSetCustomAneCompilerOptions(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetCustomAneCompilerOptions(a0, a1)
}

var _e5rtE5CompilerOptionsSetEnableMpsgraphPackage func(a0 uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsSetEnableMpsgraphPackageErr error

func tryE5rtE5CompilerOptionsSetEnableMpsgraphPackage(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetEnableMpsgraphPackage == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_enable_mpsgraph_package", "", _e5rtE5CompilerOptionsSetEnableMpsgraphPackageErr)
	}
	return _e5rtE5CompilerOptionsSetEnableMpsgraphPackage(a0, a1), nil
}

// E5rtE5CompilerOptionsSetEnableMpsgraphPackage signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtE5CompilerOptionsSetEnableMpsgraphPackage(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetEnableMpsgraphPackage(a0, a1)
}

var _e5rtE5CompilerOptionsSetEnableProfiling func(out *uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsSetEnableProfilingErr error

func tryE5rtE5CompilerOptionsSetEnableProfiling(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetEnableProfiling == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_enable_profiling", "", _e5rtE5CompilerOptionsSetEnableProfilingErr)
	}
	return _e5rtE5CompilerOptionsSetEnableProfiling(out, a1), nil
}

// E5rtE5CompilerOptionsSetEnableProfiling signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::SetEnableProfiling takes 1.
func E5rtE5CompilerOptionsSetEnableProfiling(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetEnableProfiling(out, a1)
}

var _e5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocations func(out *uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocationsErr error

func tryE5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocations(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocations == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_enable_reshape_with_minimal_allocations", "", _e5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocationsErr)
	}
	return _e5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocations(out, a1), nil
}

// E5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocations signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::SetEnableReshapeWithMinimalAllocations takes 1.
func E5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocations(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetEnableReshapeWithMinimalAllocations(out, a1)
}

var _e5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference func(a0 uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInferenceErr error

func tryE5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_experimental_disable_compile_time_mpsgraph_type_inference", "", _e5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInferenceErr)
	}
	return _e5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference(a0, a1), nil
}

// E5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference(a0, a1)
}

var _e5rtE5CompilerOptionsSetExperimentalDisableDataDependentShape func(out *uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsSetExperimentalDisableDataDependentShapeErr error

func tryE5rtE5CompilerOptionsSetExperimentalDisableDataDependentShape(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetExperimentalDisableDataDependentShape == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_experimental_disable_data_dependent_shape", "", _e5rtE5CompilerOptionsSetExperimentalDisableDataDependentShapeErr)
	}
	return _e5rtE5CompilerOptionsSetExperimentalDisableDataDependentShape(out, a1), nil
}

// E5rtE5CompilerOptionsSetExperimentalDisableDataDependentShape signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::SetExperimentalDisableDataDependentShape takes 1.
func E5rtE5CompilerOptionsSetExperimentalDisableDataDependentShape(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetExperimentalDisableDataDependentShape(out, a1)
}

var _e5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDim func(out *uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDimErr error

func tryE5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDim(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDim == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_experimental_enable_default_function_for_range_dim", "", _e5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDimErr)
	}
	return _e5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDim(out, a1), nil
}

// E5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDim signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::SetExperimentalEnableDefaultFunctionForRangeDim takes 1.
func E5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDim(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetExperimentalEnableDefaultFunctionForRangeDim(out, a1)
}

var _e5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackend func(out *uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackendErr error

func tryE5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackend(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackend == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_experimental_force_classic_cpu_backend", "", _e5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackendErr)
	}
	return _e5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackend(out, a1), nil
}

// E5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackend signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::SetExperimentalForceClassicCpuBackend takes 1.
func E5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackend(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetExperimentalForceClassicCPUBackend(out, a1)
}

var _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatterns func(out *uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsErr error

func tryE5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatterns(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatterns == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_experimental_match_e5_minimal_cpu_patterns", "", _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsErr)
	}
	return _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatterns(out, a1), nil
}

// E5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatterns signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::SetExperimentalMatchE5MinimalCpuPatterns takes 1.
func E5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatterns(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatterns(out, a1)
}

var _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStates func(out *uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStatesErr error

func tryE5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStates(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStates == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_experimental_match_e5_minimal_cpu_patterns_for_states", "", _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStatesErr)
	}
	return _e5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStates(out, a1), nil
}

// E5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStates signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::SetExperimentalMatchE5MinimalCpuPatternsForStates takes 1.
func E5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStates(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetExperimentalMatchE5MinimalCPUPatternsForStates(out, a1)
}

var _e5rtE5CompilerOptionsSetForceBnnsGraph func(a0 uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsSetForceBnnsGraphErr error

func tryE5rtE5CompilerOptionsSetForceBnnsGraph(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetForceBnnsGraph == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_force_bnns_graph", "", _e5rtE5CompilerOptionsSetForceBnnsGraphErr)
	}
	return _e5rtE5CompilerOptionsSetForceBnnsGraph(a0, a1), nil
}

// E5rtE5CompilerOptionsSetForceBnnsGraph signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtE5CompilerOptionsSetForceBnnsGraph(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetForceBnnsGraph(a0, a1)
}

var _e5rtE5CompilerOptionsSetForceClassicAotOldHw func(out *uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsSetForceClassicAotOldHwErr error

func tryE5rtE5CompilerOptionsSetForceClassicAotOldHw(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetForceClassicAotOldHw == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_force_classic_aot_old_hw", "", _e5rtE5CompilerOptionsSetForceClassicAotOldHwErr)
	}
	return _e5rtE5CompilerOptionsSetForceClassicAotOldHw(out, a1), nil
}

// E5rtE5CompilerOptionsSetForceClassicAotOldHw signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::SetForceClassicAotOldHw takes 1.
func E5rtE5CompilerOptionsSetForceClassicAotOldHw(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetForceClassicAotOldHw(out, a1)
}

var _e5rtE5CompilerOptionsSetForceFetchFromCache func(out *uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsSetForceFetchFromCacheErr error

func tryE5rtE5CompilerOptionsSetForceFetchFromCache(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetForceFetchFromCache == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_force_fetch_from_cache", "", _e5rtE5CompilerOptionsSetForceFetchFromCacheErr)
	}
	return _e5rtE5CompilerOptionsSetForceFetchFromCache(out, a1), nil
}

// E5rtE5CompilerOptionsSetForceFetchFromCache signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::SetForceFetchFromCache takes 1.
func E5rtE5CompilerOptionsSetForceFetchFromCache(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetForceFetchFromCache(out, a1)
}

var _e5rtE5CompilerOptionsSetForceRecompilation func(a0 uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsSetForceRecompilationErr error

func tryE5rtE5CompilerOptionsSetForceRecompilation(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetForceRecompilation == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_force_recompilation", "", _e5rtE5CompilerOptionsSetForceRecompilationErr)
	}
	return _e5rtE5CompilerOptionsSetForceRecompilation(a0, a1), nil
}

// E5rtE5CompilerOptionsSetForceRecompilation signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:381:9.
func E5rtE5CompilerOptionsSetForceRecompilation(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetForceRecompilation(a0, a1)
}

var _e5rtE5CompilerOptionsSetMilEntryPoints func(a0 uintptr, a1 uintptr, out *uintptr) int32
var _e5rtE5CompilerOptionsSetMilEntryPointsErr error

func tryE5rtE5CompilerOptionsSetMilEntryPoints(a0 uintptr, a1 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetMilEntryPoints == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_mil_entry_points", "", _e5rtE5CompilerOptionsSetMilEntryPointsErr)
	}
	return _e5rtE5CompilerOptionsSetMilEntryPoints(a0, a1, out), nil
}

// E5rtE5CompilerOptionsSetMilEntryPoints signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 3, E5RT::E5CompilerOptions::SetMilEntryPoints takes 1.
func E5rtE5CompilerOptionsSetMilEntryPoints(a0 uintptr, a1 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetMilEntryPoints(a0, a1, out)
}

var _e5rtE5CompilerOptionsSetPreferredCPUBackend func(out *uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsSetPreferredCPUBackendErr error

func tryE5rtE5CompilerOptionsSetPreferredCPUBackend(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetPreferredCPUBackend == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_preferred_cpu_backend", "", _e5rtE5CompilerOptionsSetPreferredCPUBackendErr)
	}
	return _e5rtE5CompilerOptionsSetPreferredCPUBackend(out, a1), nil
}

// E5rtE5CompilerOptionsSetPreferredCPUBackend signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::E5CompilerOptions::SetPreferredCpuBackend takes 1.
func E5rtE5CompilerOptionsSetPreferredCPUBackend(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetPreferredCPUBackend(out, a1)
}

var _e5rtE5CompilerOptionsSetPreferredCPUBackends func(a0 uintptr, a1 uintptr, out *uintptr) int32
var _e5rtE5CompilerOptionsSetPreferredCPUBackendsErr error

func tryE5rtE5CompilerOptionsSetPreferredCPUBackends(a0 uintptr, a1 uintptr, out *uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetPreferredCPUBackends == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_preferred_cpu_backends", "", _e5rtE5CompilerOptionsSetPreferredCPUBackendsErr)
	}
	return _e5rtE5CompilerOptionsSetPreferredCPUBackends(a0, a1, out), nil
}

// E5rtE5CompilerOptionsSetPreferredCPUBackends signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 3, E5RT::E5CompilerOptions::SetPreferredCpuBackends takes 1.
func E5rtE5CompilerOptionsSetPreferredCPUBackends(a0 uintptr, a1 uintptr, out *uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetPreferredCPUBackends(a0, a1, out)
}

var _e5rtE5CompilerOptionsSetSegmenter func(a0 uintptr, a1 uintptr) int32
var _e5rtE5CompilerOptionsSetSegmenterErr error

func tryE5rtE5CompilerOptionsSetSegmenter(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerOptionsSetSegmenter == nil {
		return 0, symbolCallError("e5rt_e5_compiler_options_set_segmenter", "", _e5rtE5CompilerOptionsSetSegmenterErr)
	}
	return _e5rtE5CompilerOptionsSetSegmenter(a0, a1), nil
}

// E5rtE5CompilerOptionsSetSegmenter signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:388:9.
func E5rtE5CompilerOptionsSetSegmenter(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerOptionsSetSegmenter(a0, a1)
}

var _e5rtE5CompilerPurgeE5BundlesForInputModel func(out *uintptr, a1 uintptr) int32
var _e5rtE5CompilerPurgeE5BundlesForInputModelErr error

func tryE5rtE5CompilerPurgeE5BundlesForInputModel(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtE5CompilerPurgeE5BundlesForInputModel == nil {
		return 0, symbolCallError("e5rt_e5_compiler_purge_e5_bundles_for_input_model", "", _e5rtE5CompilerPurgeE5BundlesForInputModelErr)
	}
	return _e5rtE5CompilerPurgeE5BundlesForInputModel(out, a1), nil
}

// E5rtE5CompilerPurgeE5BundlesForInputModel signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, C++ counterpart takes 1.
func E5rtE5CompilerPurgeE5BundlesForInputModel(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtE5CompilerPurgeE5BundlesForInputModel(out, a1)
}

var _e5rtE5CompilerRelease func(a0 uintptr) int32
var _e5rtE5CompilerReleaseErr error

func tryE5rtE5CompilerRelease(a0 uintptr) (int32, error) {
	if _e5rtE5CompilerRelease == nil {
		return 0, symbolCallError("e5rt_e5_compiler_release", "", _e5rtE5CompilerReleaseErr)
	}
	return _e5rtE5CompilerRelease(a0), nil
}

// E5rtE5CompilerRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtE5CompilerRelease(a0 uintptr) (int32, error) {
	return tryE5rtE5CompilerRelease(a0)
}

var _e5rtErrorCodeGetString func() int32
var _e5rtErrorCodeGetStringErr error

func tryE5rtErrorCodeGetString() (int32, error) {
	if _e5rtErrorCodeGetString == nil {
		return 0, symbolCallError("e5rt_error_code_get_string", "", _e5rtErrorCodeGetStringErr)
	}
	return _e5rtErrorCodeGetString(), nil
}

// E5rtErrorCodeGetString signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtErrorCodeGetString() (int32, error) {
	return tryE5rtErrorCodeGetString()
}

var _e5rtExecutionStreamConfigOptionsCreate func(out *uintptr) int32
var _e5rtExecutionStreamConfigOptionsCreateErr error

func tryE5rtExecutionStreamConfigOptionsCreate(out *uintptr) (int32, error) {
	if _e5rtExecutionStreamConfigOptionsCreate == nil {
		return 0, symbolCallError("e5rt_execution_stream_config_options_create", "", _e5rtExecutionStreamConfigOptionsCreateErr)
	}
	return _e5rtExecutionStreamConfigOptionsCreate(out), nil
}

// E5rtExecutionStreamConfigOptionsCreate signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 1, E5RT::ExecutionStreamConfigOptions::Create takes 0.
func E5rtExecutionStreamConfigOptionsCreate(out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamConfigOptionsCreate(out)
}

var _e5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecution func(a0 uintptr, out *uintptr) int32
var _e5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecutionErr error

func tryE5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecution(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecution == nil {
		return 0, symbolCallError("e5rt_execution_stream_config_options_get_enable_concurrent_sync_execution", "", _e5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecutionErr)
	}
	return _e5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecution(a0, out), nil
}

// E5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecution signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::ExecutionStreamConfigOptions::GetEnableConcurrentSyncExecution takes 0.
func E5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecution(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamConfigOptionsGetEnableConcurrentSyncExecution(a0, out)
}

var _e5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEvents func(a0 uintptr, out *uintptr) int32
var _e5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEventsErr error

func tryE5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEvents(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEvents == nil {
		return 0, symbolCallError("e5rt_execution_stream_config_options_get_enable_low_latency_async_events", "", _e5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEventsErr)
	}
	return _e5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEvents(a0, out), nil
}

// E5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEvents signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::ExecutionStreamConfigOptions::GetEnableLowLatencyAsyncEvents takes 0.
func E5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEvents(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamConfigOptionsGetEnableLowLatencyAsyncEvents(a0, out)
}

var _e5rtExecutionStreamConfigOptionsGetSkipIOFences func(a0 uintptr, out *uintptr) int32
var _e5rtExecutionStreamConfigOptionsGetSkipIOFencesErr error

func tryE5rtExecutionStreamConfigOptionsGetSkipIOFences(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtExecutionStreamConfigOptionsGetSkipIOFences == nil {
		return 0, symbolCallError("e5rt_execution_stream_config_options_get_skip_io_fences", "", _e5rtExecutionStreamConfigOptionsGetSkipIOFencesErr)
	}
	return _e5rtExecutionStreamConfigOptionsGetSkipIOFences(a0, out), nil
}

// E5rtExecutionStreamConfigOptionsGetSkipIOFences signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::ExecutionStreamConfigOptions::GetSkipIOFences takes 0.
func E5rtExecutionStreamConfigOptionsGetSkipIOFences(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamConfigOptionsGetSkipIOFences(a0, out)
}

var _e5rtExecutionStreamConfigOptionsRelease func(a0 uintptr) int32
var _e5rtExecutionStreamConfigOptionsReleaseErr error

func tryE5rtExecutionStreamConfigOptionsRelease(a0 uintptr) (int32, error) {
	if _e5rtExecutionStreamConfigOptionsRelease == nil {
		return 0, symbolCallError("e5rt_execution_stream_config_options_release", "", _e5rtExecutionStreamConfigOptionsReleaseErr)
	}
	return _e5rtExecutionStreamConfigOptionsRelease(a0), nil
}

// E5rtExecutionStreamConfigOptionsRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamConfigOptionsRelease(a0 uintptr) (int32, error) {
	return tryE5rtExecutionStreamConfigOptionsRelease(a0)
}

var _e5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecution func(out *uintptr, a1 uintptr) int32
var _e5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecutionErr error

func tryE5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecution(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecution == nil {
		return 0, symbolCallError("e5rt_execution_stream_config_options_set_enable_concurrent_sync_execution", "", _e5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecutionErr)
	}
	return _e5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecution(out, a1), nil
}

// E5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecution signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::ExecutionStreamConfigOptions::SetEnableConcurrentSyncExecution takes 1.
func E5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecution(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamConfigOptionsSetEnableConcurrentSyncExecution(out, a1)
}

var _e5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEvents func(out *uintptr, a1 uintptr) int32
var _e5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEventsErr error

func tryE5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEvents(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEvents == nil {
		return 0, symbolCallError("e5rt_execution_stream_config_options_set_enable_low_latency_async_events", "", _e5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEventsErr)
	}
	return _e5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEvents(out, a1), nil
}

// E5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEvents signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::ExecutionStreamConfigOptions::SetEnableLowLatencyAsyncEvents takes 1.
func E5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEvents(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamConfigOptionsSetEnableLowLatencyAsyncEvents(out, a1)
}

var _e5rtExecutionStreamConfigOptionsSetSkipIOFences func(out *uintptr, a1 uintptr) int32
var _e5rtExecutionStreamConfigOptionsSetSkipIOFencesErr error

func tryE5rtExecutionStreamConfigOptionsSetSkipIOFences(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamConfigOptionsSetSkipIOFences == nil {
		return 0, symbolCallError("e5rt_execution_stream_config_options_set_skip_io_fences", "", _e5rtExecutionStreamConfigOptionsSetSkipIOFencesErr)
	}
	return _e5rtExecutionStreamConfigOptionsSetSkipIOFences(out, a1), nil
}

// E5rtExecutionStreamConfigOptionsSetSkipIOFences signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::ExecutionStreamConfigOptions::SetSkipIOFences takes 1.
func E5rtExecutionStreamConfigOptionsSetSkipIOFences(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamConfigOptionsSetSkipIOFences(out, a1)
}

var _e5rtExecutionStreamCreate func(out *uintptr) int32
var _e5rtExecutionStreamCreateErr error

func tryE5rtExecutionStreamCreate(out *uintptr) (int32, error) {
	if _e5rtExecutionStreamCreate == nil {
		return 0, symbolCallError("e5rt_execution_stream_create", "", _e5rtExecutionStreamCreateErr)
	}
	return _e5rtExecutionStreamCreate(out), nil
}

// E5rtExecutionStreamCreate signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:642:9, where argument 0 is the out-parameter.
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

// E5rtExecutionStreamEncodeOperation signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:664:9.
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

// E5rtExecutionStreamEncodeWorkload signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtExecutionStreamExecuteSync signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:676:9.
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

// E5rtExecutionStreamGetInternalAsyncComputeRequestIDForLastSubmit signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::ExecutionStream::GetInternalAsyncComputeRequestIdForLastSubmit takes 0.
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

// E5rtExecutionStreamGetStreamID signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::ExecutionStream::GetStreamId takes 0.
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

// E5rtExecutionStreamOperationBindCompletionEvent signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamOperationBindCompletionEvent(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationBindCompletionEvent(a0, a1)
}

var _e5rtExecutionStreamOperationBindDependentEvents func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtExecutionStreamOperationBindDependentEventsErr error

func tryE5rtExecutionStreamOperationBindDependentEvents(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationBindDependentEvents == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_bind_dependent_events", "", _e5rtExecutionStreamOperationBindDependentEventsErr)
	}
	return _e5rtExecutionStreamOperationBindDependentEvents(a0, a1, a2), nil
}

// E5rtExecutionStreamOperationBindDependentEvents signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamOperationBindDependentEvents(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationBindDependentEvents(a0, a1, a2)
}

var _e5rtExecutionStreamOperationConfigOptionsCreate func(out *uintptr) int32
var _e5rtExecutionStreamOperationConfigOptionsCreateErr error

func tryE5rtExecutionStreamOperationConfigOptionsCreate(out *uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationConfigOptionsCreate == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_config_options_create", "", _e5rtExecutionStreamOperationConfigOptionsCreateErr)
	}
	return _e5rtExecutionStreamOperationConfigOptionsCreate(out), nil
}

// E5rtExecutionStreamOperationConfigOptionsCreate signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 1, E5RT::ExecutionStreamOperationConfigOptions::Create takes 0.
func E5rtExecutionStreamOperationConfigOptionsCreate(out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationConfigOptionsCreate(out)
}

var _e5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemory func(a0 uintptr, out *uintptr) int32
var _e5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemoryErr error

func tryE5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemory(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemory == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_config_options_get_prewire_model_memory", "", _e5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemoryErr)
	}
	return _e5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemory(a0, out), nil
}

// E5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemory signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::ExecutionStreamOperationConfigOptions::GetPrewireModelMemory takes 0.
func E5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemory(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationConfigOptionsGetPrewireModelMemory(a0, out)
}

var _e5rtExecutionStreamOperationConfigOptionsRelease func(a0 uintptr) int32
var _e5rtExecutionStreamOperationConfigOptionsReleaseErr error

func tryE5rtExecutionStreamOperationConfigOptionsRelease(a0 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationConfigOptionsRelease == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_config_options_release", "", _e5rtExecutionStreamOperationConfigOptionsReleaseErr)
	}
	return _e5rtExecutionStreamOperationConfigOptionsRelease(a0), nil
}

// E5rtExecutionStreamOperationConfigOptionsRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamOperationConfigOptionsRelease(a0 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationConfigOptionsRelease(a0)
}

var _e5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemory func(out *uintptr, a1 uintptr) int32
var _e5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemoryErr error

func tryE5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemory(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemory == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_config_options_set_prewire_model_memory", "", _e5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemoryErr)
	}
	return _e5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemory(out, a1), nil
}

// E5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemory signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::ExecutionStreamOperationConfigOptions::SetPrewireModelMemory takes 1.
func E5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemory(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationConfigOptionsSetPrewireModelMemory(out, a1)
}

var _e5rtExecutionStreamOperationCreatePrecompiledComputeOperation func(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) int32
var _e5rtExecutionStreamOperationCreatePrecompiledComputeOperationErr error

func tryE5rtExecutionStreamOperationCreatePrecompiledComputeOperation(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationCreatePrecompiledComputeOperation == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_create_precompiled_compute_operation", "", _e5rtExecutionStreamOperationCreatePrecompiledComputeOperationErr)
	}
	return _e5rtExecutionStreamOperationCreatePrecompiledComputeOperation(out, a1, a2, a3, a4, a5), nil
}

// E5rtExecutionStreamOperationCreatePrecompiledComputeOperation signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: wraps a constructor, so the out-parameter is the object being made.
func E5rtExecutionStreamOperationCreatePrecompiledComputeOperation(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationCreatePrecompiledComputeOperation(out, a1, a2, a3, a4, a5)
}

var _e5rtExecutionStreamOperationCreatePrecompiledComputeOperationWithOptions func(out *uintptr, a1 uintptr) int32
var _e5rtExecutionStreamOperationCreatePrecompiledComputeOperationWithOptionsErr error

func tryE5rtExecutionStreamOperationCreatePrecompiledComputeOperationWithOptions(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationCreatePrecompiledComputeOperationWithOptions == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_create_precompiled_compute_operation_with_options", "", _e5rtExecutionStreamOperationCreatePrecompiledComputeOperationWithOptionsErr)
	}
	return _e5rtExecutionStreamOperationCreatePrecompiledComputeOperationWithOptions(out, a1), nil
}

// E5rtExecutionStreamOperationCreatePrecompiledComputeOperationWithOptions signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:538:9, where argument 0 is the out-parameter.
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

// E5rtExecutionStreamOperationCreateResourceSharingPrecompiledComputeOperationsWithMultipleOptions signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: wraps a constructor, so the out-parameter is the object being made.
func E5rtExecutionStreamOperationCreateResourceSharingPrecompiledComputeOperationsWithMultipleOptions(out *uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationCreateResourceSharingPrecompiledComputeOperationsWithMultipleOptions(out, a1, a2)
}

var _e5rtExecutionStreamOperationGetDependentEventCount func(a0 uintptr, a1 uintptr) int32
var _e5rtExecutionStreamOperationGetDependentEventCountErr error

func tryE5rtExecutionStreamOperationGetDependentEventCount(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationGetDependentEventCount == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_get_dependent_event_count", "", _e5rtExecutionStreamOperationGetDependentEventCountErr)
	}
	return _e5rtExecutionStreamOperationGetDependentEventCount(a0, a1), nil
}

// E5rtExecutionStreamOperationGetDependentEventCount signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamOperationGetDependentEventCount(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationGetDependentEventCount(a0, a1)
}

var _e5rtExecutionStreamOperationGetInoutNames func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtExecutionStreamOperationGetInoutNamesErr error

func tryE5rtExecutionStreamOperationGetInoutNames(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationGetInoutNames == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_get_inout_names", "", _e5rtExecutionStreamOperationGetInoutNamesErr)
	}
	return _e5rtExecutionStreamOperationGetInoutNames(a0, a1, a2), nil
}

// E5rtExecutionStreamOperationGetInoutNames signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamOperationGetInoutNames(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationGetInoutNames(a0, a1, a2)
}

var _e5rtExecutionStreamOperationGetInputNames func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtExecutionStreamOperationGetInputNamesErr error

func tryE5rtExecutionStreamOperationGetInputNames(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationGetInputNames == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_get_input_names", "", _e5rtExecutionStreamOperationGetInputNamesErr)
	}
	return _e5rtExecutionStreamOperationGetInputNames(a0, a1, a2), nil
}

// E5rtExecutionStreamOperationGetInputNames signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamOperationGetInputNames(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationGetInputNames(a0, a1, a2)
}

var _e5rtExecutionStreamOperationGetNumInouts func(a0 uintptr, a1 uintptr) int32
var _e5rtExecutionStreamOperationGetNumInoutsErr error

func tryE5rtExecutionStreamOperationGetNumInouts(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationGetNumInouts == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_get_num_inouts", "", _e5rtExecutionStreamOperationGetNumInoutsErr)
	}
	return _e5rtExecutionStreamOperationGetNumInouts(a0, a1), nil
}

// E5rtExecutionStreamOperationGetNumInouts signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamOperationGetNumInouts(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationGetNumInouts(a0, a1)
}

var _e5rtExecutionStreamOperationGetNumInputs func(a0 uintptr, a1 uintptr) int32
var _e5rtExecutionStreamOperationGetNumInputsErr error

func tryE5rtExecutionStreamOperationGetNumInputs(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationGetNumInputs == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_get_num_inputs", "", _e5rtExecutionStreamOperationGetNumInputsErr)
	}
	return _e5rtExecutionStreamOperationGetNumInputs(a0, a1), nil
}

// E5rtExecutionStreamOperationGetNumInputs signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamOperationGetNumInputs(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationGetNumInputs(a0, a1)
}

var _e5rtExecutionStreamOperationGetNumOutputs func(a0 uintptr, a1 uintptr) int32
var _e5rtExecutionStreamOperationGetNumOutputsErr error

func tryE5rtExecutionStreamOperationGetNumOutputs(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationGetNumOutputs == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_get_num_outputs", "", _e5rtExecutionStreamOperationGetNumOutputsErr)
	}
	return _e5rtExecutionStreamOperationGetNumOutputs(a0, a1), nil
}

// E5rtExecutionStreamOperationGetNumOutputs signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamOperationGetNumOutputs(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationGetNumOutputs(a0, a1)
}

var _e5rtExecutionStreamOperationGetOpname func(a0 uintptr, a1 uintptr) int32
var _e5rtExecutionStreamOperationGetOpnameErr error

func tryE5rtExecutionStreamOperationGetOpname(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationGetOpname == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_get_opname", "", _e5rtExecutionStreamOperationGetOpnameErr)
	}
	return _e5rtExecutionStreamOperationGetOpname(a0, a1), nil
}

// E5rtExecutionStreamOperationGetOpname signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamOperationGetOpname(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationGetOpname(a0, a1)
}

var _e5rtExecutionStreamOperationGetOutputNames func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtExecutionStreamOperationGetOutputNamesErr error

func tryE5rtExecutionStreamOperationGetOutputNames(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationGetOutputNames == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_get_output_names", "", _e5rtExecutionStreamOperationGetOutputNamesErr)
	}
	return _e5rtExecutionStreamOperationGetOutputNames(a0, a1, a2), nil
}

// E5rtExecutionStreamOperationGetOutputNames signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamOperationGetOutputNames(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
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

// E5rtExecutionStreamOperationPrepareOpForEncode signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:655:9.
func E5rtExecutionStreamOperationPrepareOpForEncode(a0 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationPrepareOpForEncode(a0)
}

var _e5rtExecutionStreamOperationRelease func(a0 uintptr) int32
var _e5rtExecutionStreamOperationReleaseErr error

func tryE5rtExecutionStreamOperationRelease(a0 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationRelease == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_release", "", _e5rtExecutionStreamOperationReleaseErr)
	}
	return _e5rtExecutionStreamOperationRelease(a0), nil
}

// E5rtExecutionStreamOperationRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamOperationRelease(a0 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationRelease(a0)
}

var _e5rtExecutionStreamOperationReshapeOperation func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtExecutionStreamOperationReshapeOperationErr error

func tryE5rtExecutionStreamOperationReshapeOperation(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationReshapeOperation == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_reshape_operation", "", _e5rtExecutionStreamOperationReshapeOperationErr)
	}
	return _e5rtExecutionStreamOperationReshapeOperation(a0, a1, a2, a3), nil
}

// E5rtExecutionStreamOperationReshapeOperation signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtExecutionStreamOperationRetainCompletionEvent signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtExecutionStreamOperationRetainDependentEvents signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamOperationRetainDependentEvents(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationRetainDependentEvents(a0, a1)
}

var _e5rtExecutionStreamOperationRetainInoutPort func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtExecutionStreamOperationRetainInoutPortErr error

func tryE5rtExecutionStreamOperationRetainInoutPort(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationRetainInoutPort == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_retain_inout_port", "", _e5rtExecutionStreamOperationRetainInoutPortErr)
	}
	return _e5rtExecutionStreamOperationRetainInoutPort(a0, a1, a2), nil
}

// E5rtExecutionStreamOperationRetainInoutPort signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamOperationRetainInoutPort(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationRetainInoutPort(a0, a1, a2)
}

var _e5rtExecutionStreamOperationRetainInputPort func(a0 uintptr, a1 uintptr, out *uintptr) int32
var _e5rtExecutionStreamOperationRetainInputPortErr error

func tryE5rtExecutionStreamOperationRetainInputPort(a0 uintptr, a1 uintptr, out *uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationRetainInputPort == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_retain_input_port", "", _e5rtExecutionStreamOperationRetainInputPortErr)
	}
	return _e5rtExecutionStreamOperationRetainInputPort(a0, a1, out), nil
}

// E5rtExecutionStreamOperationRetainInputPort signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:598:9, where argument 2 is the out-parameter.
func E5rtExecutionStreamOperationRetainInputPort(a0 uintptr, a1 uintptr, out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationRetainInputPort(a0, a1, out)
}

var _e5rtExecutionStreamOperationRetainOutputPort func(a0 uintptr, a1 uintptr, out *uintptr) int32
var _e5rtExecutionStreamOperationRetainOutputPortErr error

func tryE5rtExecutionStreamOperationRetainOutputPort(a0 uintptr, a1 uintptr, out *uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationRetainOutputPort == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_retain_output_port", "", _e5rtExecutionStreamOperationRetainOutputPortErr)
	}
	return _e5rtExecutionStreamOperationRetainOutputPort(a0, a1, out), nil
}

// E5rtExecutionStreamOperationRetainOutputPort signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:612:9, where argument 2 is the out-parameter.
func E5rtExecutionStreamOperationRetainOutputPort(a0 uintptr, a1 uintptr, out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationRetainOutputPort(a0, a1, out)
}

var _e5rtExecutionStreamOperationSerializeInferenceFrameData func(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtExecutionStreamOperationSerializeInferenceFrameDataErr error

func tryE5rtExecutionStreamOperationSerializeInferenceFrameData(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationSerializeInferenceFrameData == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_serialize_inference_frame_data", "", _e5rtExecutionStreamOperationSerializeInferenceFrameDataErr)
	}
	return _e5rtExecutionStreamOperationSerializeInferenceFrameData(out, a1, a2, a3), nil
}

// E5rtExecutionStreamOperationSerializeInferenceFrameData signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 4, E5RT::ExecutionStreamOperation::SerializeInferenceFrameData takes 3.
func E5rtExecutionStreamOperationSerializeInferenceFrameData(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationSerializeInferenceFrameData(out, a1, a2, a3)
}

var _e5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegment func(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegmentErr error

func tryE5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegment(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegment == nil {
		return 0, symbolCallError("e5rt_execution_stream_operation_serialize_inference_frame_data_per_segment", "", _e5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegmentErr)
	}
	return _e5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegment(out, a1, a2, a3), nil
}

// E5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegment signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 4, E5RT::ExecutionStreamOperation::SerializeInferenceFrameDataPerSegment takes 3.
func E5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegment(out *uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtExecutionStreamOperationSerializeInferenceFrameDataPerSegment(out, a1, a2, a3)
}

var _e5rtExecutionStreamPrewireInUseAllocations func(a0 uintptr) int32
var _e5rtExecutionStreamPrewireInUseAllocationsErr error

func tryE5rtExecutionStreamPrewireInUseAllocations(a0 uintptr) (int32, error) {
	if _e5rtExecutionStreamPrewireInUseAllocations == nil {
		return 0, symbolCallError("e5rt_execution_stream_prewire_in_use_allocations", "", _e5rtExecutionStreamPrewireInUseAllocationsErr)
	}
	return _e5rtExecutionStreamPrewireInUseAllocations(a0), nil
}

// E5rtExecutionStreamPrewireInUseAllocations signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamPrewireInUseAllocations(a0 uintptr) (int32, error) {
	return tryE5rtExecutionStreamPrewireInUseAllocations(a0)
}

var _e5rtExecutionStreamRelease func(a0 uintptr) int32
var _e5rtExecutionStreamReleaseErr error

func tryE5rtExecutionStreamRelease(a0 uintptr) (int32, error) {
	if _e5rtExecutionStreamRelease == nil {
		return 0, symbolCallError("e5rt_execution_stream_release", "", _e5rtExecutionStreamReleaseErr)
	}
	return _e5rtExecutionStreamRelease(a0), nil
}

// E5rtExecutionStreamRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamRelease(a0 uintptr) (int32, error) {
	return tryE5rtExecutionStreamRelease(a0)
}

var _e5rtExecutionStreamReset func(a0 uintptr) int32
var _e5rtExecutionStreamResetErr error

func tryE5rtExecutionStreamReset(a0 uintptr) (int32, error) {
	if _e5rtExecutionStreamReset == nil {
		return 0, symbolCallError("e5rt_execution_stream_reset", "", _e5rtExecutionStreamResetErr)
	}
	return _e5rtExecutionStreamReset(a0), nil
}

// E5rtExecutionStreamReset signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:692:9.
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

// E5rtExecutionStreamResetConfigOptions signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::ExecutionStream::ResetConfigOptions takes 0.
func E5rtExecutionStreamResetConfigOptions(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtExecutionStreamResetConfigOptions(a0, out)
}

var _e5rtExecutionStreamSetAneExecutionPriority func(a0 uintptr, a1 uintptr) int32
var _e5rtExecutionStreamSetAneExecutionPriorityErr error

func tryE5rtExecutionStreamSetAneExecutionPriority(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamSetAneExecutionPriority == nil {
		return 0, symbolCallError("e5rt_execution_stream_set_ane_execution_priority", "", _e5rtExecutionStreamSetAneExecutionPriorityErr)
	}
	return _e5rtExecutionStreamSetAneExecutionPriority(a0, a1), nil
}

// E5rtExecutionStreamSetAneExecutionPriority signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamSetAneExecutionPriority(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamSetAneExecutionPriority(a0, a1)
}

var _e5rtExecutionStreamSetConfigOptions func(out *uintptr, a1 uintptr) int32
var _e5rtExecutionStreamSetConfigOptionsErr error

func tryE5rtExecutionStreamSetConfigOptions(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamSetConfigOptions == nil {
		return 0, symbolCallError("e5rt_execution_stream_set_config_options", "", _e5rtExecutionStreamSetConfigOptionsErr)
	}
	return _e5rtExecutionStreamSetConfigOptions(out, a1), nil
}

// E5rtExecutionStreamSetConfigOptions signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::ExecutionStream::SetConfigOptions takes 1.
func E5rtExecutionStreamSetConfigOptions(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtExecutionStreamSetConfigOptions(out, a1)
}

var _e5rtExecutionStreamSetQualityOfService func(a0 uintptr, a1 uintptr) int32
var _e5rtExecutionStreamSetQualityOfServiceErr error

func tryE5rtExecutionStreamSetQualityOfService(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtExecutionStreamSetQualityOfService == nil {
		return 0, symbolCallError("e5rt_execution_stream_set_quality_of_service", "", _e5rtExecutionStreamSetQualityOfServiceErr)
	}
	return _e5rtExecutionStreamSetQualityOfService(a0, a1), nil
}

// E5rtExecutionStreamSetQualityOfService signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamSetQualityOfService(a0 uintptr, a1 uintptr) (int32, error) {
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

// E5rtExecutionStreamStepExecuteSync signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtExecutionStreamSubmitAsync signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtExecutionStreamSubmitAsyncWithTimeout signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtExecutionStreamSubmitAsyncWithTimeout(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtExecutionStreamSubmitAsyncWithTimeout(a0, a1, a2)
}

var _e5rtGetLastErrorMessage func() int32
var _e5rtGetLastErrorMessageErr error

func tryE5rtGetLastErrorMessage() (int32, error) {
	if _e5rtGetLastErrorMessage == nil {
		return 0, symbolCallError("e5rt_get_last_error_message", "", _e5rtGetLastErrorMessageErr)
	}
	return _e5rtGetLastErrorMessage(), nil
}

// E5rtGetLastErrorMessage signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtGetLastErrorMessage() (int32, error) {
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

// E5rtIOPortBindBufferObject signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:628:9.
func E5rtIOPortBindBufferObject(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtIOPortBindBufferObject(a0, a1)
}

var _e5rtIOPortBindMemoryObject func(out *uintptr, a1 uintptr) int32
var _e5rtIOPortBindMemoryObjectErr error

func tryE5rtIOPortBindMemoryObject(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtIOPortBindMemoryObject == nil {
		return 0, symbolCallError("e5rt_io_port_bind_memory_object", "", _e5rtIOPortBindMemoryObjectErr)
	}
	return _e5rtIOPortBindMemoryObject(out, a1), nil
}

// E5rtIOPortBindMemoryObject signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::IOPort::BindMemoryObject takes 1.
func E5rtIOPortBindMemoryObject(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtIOPortBindMemoryObject(out, a1)
}

var _e5rtIOPortBindSurfaceObject func(a0 uintptr, a1 uintptr) int32
var _e5rtIOPortBindSurfaceObjectErr error

func tryE5rtIOPortBindSurfaceObject(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtIOPortBindSurfaceObject == nil {
		return 0, symbolCallError("e5rt_io_port_bind_surface_object", "", _e5rtIOPortBindSurfaceObjectErr)
	}
	return _e5rtIOPortBindSurfaceObject(a0, a1), nil
}

// E5rtIOPortBindSurfaceObject signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtIOPortGetSupportedBufferTypes signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtIOPortHasKnownShape signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::IOPort::HasKnownShape takes 0.
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

// E5rtIOPortIsDynamic signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::IOPort::IsDynamic takes 0.
func E5rtIOPortIsDynamic(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtIOPortIsDynamic(a0, out)
}

var _e5rtIOPortIsSurface func(a0 uintptr, a1 uintptr) int32
var _e5rtIOPortIsSurfaceErr error

func tryE5rtIOPortIsSurface(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtIOPortIsSurface == nil {
		return 0, symbolCallError("e5rt_io_port_is_surface", "", _e5rtIOPortIsSurfaceErr)
	}
	return _e5rtIOPortIsSurface(a0, a1), nil
}

// E5rtIOPortIsSurface signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtIOPortIsSurface(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtIOPortIsSurface(a0, a1)
}

var _e5rtIOPortIsTensor func(a0 uintptr, a1 uintptr) int32
var _e5rtIOPortIsTensorErr error

func tryE5rtIOPortIsTensor(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtIOPortIsTensor == nil {
		return 0, symbolCallError("e5rt_io_port_is_tensor", "", _e5rtIOPortIsTensorErr)
	}
	return _e5rtIOPortIsTensor(a0, a1), nil
}

// E5rtIOPortIsTensor signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtIOPortIsTensor(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtIOPortIsTensor(a0, a1)
}

var _e5rtIOPortRelease func(a0 uintptr) int32
var _e5rtIOPortReleaseErr error

func tryE5rtIOPortRelease(a0 uintptr) (int32, error) {
	if _e5rtIOPortRelease == nil {
		return 0, symbolCallError("e5rt_io_port_release", "", _e5rtIOPortReleaseErr)
	}
	return _e5rtIOPortRelease(a0), nil
}

// E5rtIOPortRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtIOPortRelease(a0 uintptr) (int32, error) {
	return tryE5rtIOPortRelease(a0)
}

var _e5rtIOPortRetainBufferObject func(a0 uintptr, a1 uintptr) int32
var _e5rtIOPortRetainBufferObjectErr error

func tryE5rtIOPortRetainBufferObject(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtIOPortRetainBufferObject == nil {
		return 0, symbolCallError("e5rt_io_port_retain_buffer_object", "", _e5rtIOPortRetainBufferObjectErr)
	}
	return _e5rtIOPortRetainBufferObject(a0, a1), nil
}

// E5rtIOPortRetainBufferObject signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtIOPortRetainBufferObject(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtIOPortRetainBufferObject(a0, a1)
}

var _e5rtIOPortRetainMemoryObject func(a0 uintptr, a1 uintptr) int32
var _e5rtIOPortRetainMemoryObjectErr error

func tryE5rtIOPortRetainMemoryObject(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtIOPortRetainMemoryObject == nil {
		return 0, symbolCallError("e5rt_io_port_retain_memory_object", "", _e5rtIOPortRetainMemoryObjectErr)
	}
	return _e5rtIOPortRetainMemoryObject(a0, a1), nil
}

// E5rtIOPortRetainMemoryObject signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtIOPortRetainMemoryObject(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtIOPortRetainMemoryObject(a0, a1)
}

var _e5rtIOPortRetainSurfaceDesc func(a0 uintptr, a1 uintptr) int32
var _e5rtIOPortRetainSurfaceDescErr error

func tryE5rtIOPortRetainSurfaceDesc(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtIOPortRetainSurfaceDesc == nil {
		return 0, symbolCallError("e5rt_io_port_retain_surface_desc", "", _e5rtIOPortRetainSurfaceDescErr)
	}
	return _e5rtIOPortRetainSurfaceDesc(a0, a1), nil
}

// E5rtIOPortRetainSurfaceDesc signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtIOPortRetainSurfaceDesc(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtIOPortRetainSurfaceDesc(a0, a1)
}

var _e5rtIOPortRetainSurfaceObject func(a0 uintptr, a1 uintptr) int32
var _e5rtIOPortRetainSurfaceObjectErr error

func tryE5rtIOPortRetainSurfaceObject(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtIOPortRetainSurfaceObject == nil {
		return 0, symbolCallError("e5rt_io_port_retain_surface_object", "", _e5rtIOPortRetainSurfaceObjectErr)
	}
	return _e5rtIOPortRetainSurfaceObject(a0, a1), nil
}

// E5rtIOPortRetainSurfaceObject signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtIOPortRetainSurfaceObject(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtIOPortRetainSurfaceObject(a0, a1)
}

var _e5rtIOPortRetainTensorDesc func(a0 uintptr, a1 uintptr) int32
var _e5rtIOPortRetainTensorDescErr error

func tryE5rtIOPortRetainTensorDesc(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtIOPortRetainTensorDesc == nil {
		return 0, symbolCallError("e5rt_io_port_retain_tensor_desc", "", _e5rtIOPortRetainTensorDescErr)
	}
	return _e5rtIOPortRetainTensorDesc(a0, a1), nil
}

// E5rtIOPortRetainTensorDesc signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtIOPortRetainTensorDesc(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtIOPortRetainTensorDesc(a0, a1)
}

var _e5rtMemoryObjectCreate func(out *uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtMemoryObjectCreateErr error

func tryE5rtMemoryObjectCreate(out *uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtMemoryObjectCreate == nil {
		return 0, symbolCallError("e5rt_memory_object_create", "", _e5rtMemoryObjectCreateErr)
	}
	return _e5rtMemoryObjectCreate(out, a1, a2), nil
}

// E5rtMemoryObjectCreate signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: wraps a constructor, so the out-parameter is the object being made.
func E5rtMemoryObjectCreate(out *uintptr, a1 uintptr, a2 uintptr) (int32, error) {
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

// E5rtMemoryObjectCreateAsAlias signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: wraps a constructor, so the out-parameter is the object being made.
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

// E5rtMemoryObjectCreateFromIosurface signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: wraps a constructor, so the out-parameter is the object being made.
func E5rtMemoryObjectCreateFromIosurface(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtMemoryObjectCreateFromIosurface(out, a1)
}

var _e5rtMemoryObjectGetDataPtr func(a0 uintptr, a1 uintptr) int32
var _e5rtMemoryObjectGetDataPtrErr error

func tryE5rtMemoryObjectGetDataPtr(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtMemoryObjectGetDataPtr == nil {
		return 0, symbolCallError("e5rt_memory_object_get_data_ptr", "", _e5rtMemoryObjectGetDataPtrErr)
	}
	return _e5rtMemoryObjectGetDataPtr(a0, a1), nil
}

// E5rtMemoryObjectGetDataPtr signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtMemoryObjectGetDataPtr(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtMemoryObjectGetDataPtr(a0, a1)
}

var _e5rtMemoryObjectGetIosurface func(a0 uintptr, a1 uintptr) int32
var _e5rtMemoryObjectGetIosurfaceErr error

func tryE5rtMemoryObjectGetIosurface(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtMemoryObjectGetIosurface == nil {
		return 0, symbolCallError("e5rt_memory_object_get_iosurface", "", _e5rtMemoryObjectGetIosurfaceErr)
	}
	return _e5rtMemoryObjectGetIosurface(a0, a1), nil
}

// E5rtMemoryObjectGetIosurface signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtMemoryObjectGetIosurface(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtMemoryObjectGetIosurface(a0, a1)
}

var _e5rtMemoryObjectGetSize func(a0 uintptr, a1 uintptr) int32
var _e5rtMemoryObjectGetSizeErr error

func tryE5rtMemoryObjectGetSize(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtMemoryObjectGetSize == nil {
		return 0, symbolCallError("e5rt_memory_object_get_size", "", _e5rtMemoryObjectGetSizeErr)
	}
	return _e5rtMemoryObjectGetSize(a0, a1), nil
}

// E5rtMemoryObjectGetSize signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtMemoryObjectGetSize(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtMemoryObjectGetSize(a0, a1)
}

var _e5rtMemoryObjectRelease func(a0 uintptr) int32
var _e5rtMemoryObjectReleaseErr error

func tryE5rtMemoryObjectRelease(a0 uintptr) (int32, error) {
	if _e5rtMemoryObjectRelease == nil {
		return 0, symbolCallError("e5rt_memory_object_release", "", _e5rtMemoryObjectReleaseErr)
	}
	return _e5rtMemoryObjectRelease(a0), nil
}

// E5rtMemoryObjectRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtMemoryObjectRelease(a0 uintptr) (int32, error) {
	return tryE5rtMemoryObjectRelease(a0)
}

var _e5rtOperandDescIsSurfaceDesc func(a0 uintptr, a1 uintptr) int32
var _e5rtOperandDescIsSurfaceDescErr error

func tryE5rtOperandDescIsSurfaceDesc(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtOperandDescIsSurfaceDesc == nil {
		return 0, symbolCallError("e5rt_operand_desc_is_surface_desc", "", _e5rtOperandDescIsSurfaceDescErr)
	}
	return _e5rtOperandDescIsSurfaceDesc(a0, a1), nil
}

// E5rtOperandDescIsSurfaceDesc signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtOperandDescIsSurfaceDesc(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtOperandDescIsSurfaceDesc(a0, a1)
}

var _e5rtOperandDescIsTensorDesc func(a0 uintptr, a1 uintptr) int32
var _e5rtOperandDescIsTensorDescErr error

func tryE5rtOperandDescIsTensorDesc(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtOperandDescIsTensorDesc == nil {
		return 0, symbolCallError("e5rt_operand_desc_is_tensor_desc", "", _e5rtOperandDescIsTensorDescErr)
	}
	return _e5rtOperandDescIsTensorDesc(a0, a1), nil
}

// E5rtOperandDescIsTensorDesc signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtOperandDescIsTensorDesc(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtOperandDescIsTensorDesc(a0, a1)
}

var _e5rtOperandDescRelease func(a0 uintptr) int32
var _e5rtOperandDescReleaseErr error

func tryE5rtOperandDescRelease(a0 uintptr) (int32, error) {
	if _e5rtOperandDescRelease == nil {
		return 0, symbolCallError("e5rt_operand_desc_release", "", _e5rtOperandDescReleaseErr)
	}
	return _e5rtOperandDescRelease(a0), nil
}

// E5rtOperandDescRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtOperandDescRelease(a0 uintptr) (int32, error) {
	return tryE5rtOperandDescRelease(a0)
}

var _e5rtOperandDescRetainFromSurfaceDesc func(a0 uintptr, a1 uintptr) int32
var _e5rtOperandDescRetainFromSurfaceDescErr error

func tryE5rtOperandDescRetainFromSurfaceDesc(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtOperandDescRetainFromSurfaceDesc == nil {
		return 0, symbolCallError("e5rt_operand_desc_retain_from_surface_desc", "", _e5rtOperandDescRetainFromSurfaceDescErr)
	}
	return _e5rtOperandDescRetainFromSurfaceDesc(a0, a1), nil
}

// E5rtOperandDescRetainFromSurfaceDesc signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtOperandDescRetainFromSurfaceDesc(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtOperandDescRetainFromSurfaceDesc(a0, a1)
}

var _e5rtOperandDescRetainFromTensorDesc func(a0 uintptr, a1 uintptr) int32
var _e5rtOperandDescRetainFromTensorDescErr error

func tryE5rtOperandDescRetainFromTensorDesc(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtOperandDescRetainFromTensorDesc == nil {
		return 0, symbolCallError("e5rt_operand_desc_retain_from_tensor_desc", "", _e5rtOperandDescRetainFromTensorDescErr)
	}
	return _e5rtOperandDescRetainFromTensorDesc(a0, a1), nil
}

// E5rtOperandDescRetainFromTensorDesc signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtPrecompiledComputeOpCreateOptionsCopyDynamicCallables signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtPrecompiledComputeOpCreateOptionsCopyMutableMilWeightPaths signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtPrecompiledComputeOpCreateOptionsCopyMutableMilWeightPaths(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsCopyMutableMilWeightPaths(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsCreate func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsCreateErr error

func tryE5rtPrecompiledComputeOpCreateOptionsCreate(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsCreate == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_create", "", _e5rtPrecompiledComputeOpCreateOptionsCreateErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsCreate(a0, a1, a2), nil
}

// E5rtPrecompiledComputeOpCreateOptionsCreate signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtPrecompiledComputeOpCreateOptionsCreate(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
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

// E5rtPrecompiledComputeOpCreateOptionsCreateWithProgramFunction signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:500:9, where argument 0 is the out-parameter.
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

// E5rtPrecompiledComputeOpCreateOptionsGetAllocateIntermediateBuffers signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::PrecompiledComputeOpCreateOptions::GetAllocateIntermediateBuffers takes 0.
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

// E5rtPrecompiledComputeOpCreateOptionsGetExperimentalEnableMpsgraphParallelEncode signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtPrecompiledComputeOpCreateOptionsGetIosurfaceMemoryPoolID signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtPrecompiledComputeOpCreateOptionsGetLazyPrepareOpForEncode signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::PrecompiledComputeOpCreateOptions::GetLazyPrepareOpForEncode takes 0.
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

// E5rtPrecompiledComputeOpCreateOptionsGetOperationName signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::PrecompiledComputeOpCreateOptions::GetOperationName takes 0.
func E5rtPrecompiledComputeOpCreateOptionsGetOperationName(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsGetOperationName(a0, out)
}

var _e5rtPrecompiledComputeOpCreateOptionsRelease func(a0 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsReleaseErr error

func tryE5rtPrecompiledComputeOpCreateOptionsRelease(a0 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsRelease == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_release", "", _e5rtPrecompiledComputeOpCreateOptionsReleaseErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsRelease(a0), nil
}

// E5rtPrecompiledComputeOpCreateOptionsRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtPrecompiledComputeOpCreateOptionsRelease(a0 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsRelease(a0)
}

var _e5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDevice func(a0 uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDeviceErr error

func tryE5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDevice(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDevice == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_retain_override_compute_gpu_device", "", _e5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDeviceErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDevice(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDevice signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDevice(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsRetainOverrideComputeGPUDevice(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffers func(a0 uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffersErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffers(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffers == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_allocate_intermediate_buffers", "", _e5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffersErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffers(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffers signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:520:9.
func E5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffers(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetAllocateIntermediateBuffers(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProvider func(out *uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProviderErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProvider(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProvider == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_custom_ane_memory_provider", "", _e5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProviderErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProvider(out, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProvider signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::PrecompiledComputeOpCreateOptions::SetCustomANEMemoryProvider takes 1.
func E5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProvider(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetCustomAneMemoryProvider(out, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetDynamicCallables func(out *uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetDynamicCallablesErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetDynamicCallables(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetDynamicCallables == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_dynamic_callables", "", _e5rtPrecompiledComputeOpCreateOptionsSetDynamicCallablesErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetDynamicCallables(out, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetDynamicCallables signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::PrecompiledComputeOpCreateOptions::SetDynamicCallables takes 1.
func E5rtPrecompiledComputeOpCreateOptionsSetDynamicCallables(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetDynamicCallables(out, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference func(a0 uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInferenceErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_experimental_disable_compile_time_mpsgraph_type_inference", "", _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInferenceErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetExperimentalDisableCompileTimeMpsgraphTypeInference(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOps func(out *uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOpsErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOps(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOps == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_experimental_enable_gpu_quant_ops", "", _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOpsErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOps(out, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOps signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::PrecompiledComputeOpCreateOptions::SetExperimentalEnableGPUQuantOps takes 1.
func E5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOps(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableGPUQuantOps(out, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecision func(a0 uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecisionErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecision(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecision == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_experimental_enable_mps_reduced_precision", "", _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecisionErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecision(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecision signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecision(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsReducedPrecision(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsgraphParallelEncode func(a0 uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsgraphParallelEncodeErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsgraphParallelEncode(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsgraphParallelEncode == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_experimental_enable_mpsgraph_parallel_encode", "", _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsgraphParallelEncodeErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsgraphParallelEncode(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsgraphParallelEncode signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtPrecompiledComputeOpCreateOptionsSetExperimentalEnableMpsgraphParallelEncode(a0 uintptr, a1 uintptr) (int32, error) {
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

// E5rtPrecompiledComputeOpCreateOptionsSetExperimentalMpsgraphMaximumNumberOfEncodingThreads signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtPrecompiledComputeOpCreateOptionsSetIosurfaceMemoryPoolID signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtPrecompiledComputeOpCreateOptionsSetIosurfaceMemoryPoolID(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetIosurfaceMemoryPoolID(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncode func(out *uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncodeErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncode(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncode == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_lazy_prepare_op_for_encode", "", _e5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncodeErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncode(out, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncode signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::PrecompiledComputeOpCreateOptions::SetLazyPrepareOpForEncode takes 1.
func E5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncode(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetLazyPrepareOpForEncode(out, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPaths func(out *uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPathsErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPaths(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPaths == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_mutable_mil_weight_paths", "", _e5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPathsErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPaths(out, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPaths signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::PrecompiledComputeOpCreateOptions::SetMutableMILWeightPaths takes 1.
func E5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPaths(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetMutableMilWeightPaths(out, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetOperationName func(a0 uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetOperationNameErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetOperationName(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetOperationName == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_operation_name", "", _e5rtPrecompiledComputeOpCreateOptionsSetOperationNameErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetOperationName(a0, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetOperationName signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:510:9.
func E5rtPrecompiledComputeOpCreateOptionsSetOperationName(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetOperationName(a0, a1)
}

var _e5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDevice func(out *uintptr, a1 uintptr) int32
var _e5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDeviceErr error

func tryE5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDevice(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDevice == nil {
		return 0, symbolCallError("e5rt_precompiled_compute_op_create_options_set_override_compute_gpu_device", "", _e5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDeviceErr)
	}
	return _e5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDevice(out, a1), nil
}

// E5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDevice signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::PrecompiledComputeOpCreateOptions::SetOverrideComputeGPUDevice takes 1.
func E5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDevice(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtPrecompiledComputeOpCreateOptionsSetOverrideComputeGPUDevice(out, a1)
}

var _e5rtProgramFunctionGetExternInoutNames func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtProgramFunctionGetExternInoutNamesErr error

func tryE5rtProgramFunctionGetExternInoutNames(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtProgramFunctionGetExternInoutNames == nil {
		return 0, symbolCallError("e5rt_program_function_get_extern_inout_names", "", _e5rtProgramFunctionGetExternInoutNamesErr)
	}
	return _e5rtProgramFunctionGetExternInoutNames(a0, a1, a2), nil
}

// E5rtProgramFunctionGetExternInoutNames signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtProgramFunctionGetExternInputNames signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtProgramFunctionGetExternOutputNames signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtProgramFunctionGetExternOutputNames(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtProgramFunctionGetExternOutputNames(a0, a1, a2)
}

var _e5rtProgramFunctionGetName func(a0 uintptr, a1 uintptr) int32
var _e5rtProgramFunctionGetNameErr error

func tryE5rtProgramFunctionGetName(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtProgramFunctionGetName == nil {
		return 0, symbolCallError("e5rt_program_function_get_name", "", _e5rtProgramFunctionGetNameErr)
	}
	return _e5rtProgramFunctionGetName(a0, a1), nil
}

// E5rtProgramFunctionGetName signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtProgramFunctionGetName(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtProgramFunctionGetName(a0, a1)
}

var _e5rtProgramFunctionGetNumExternInouts func(a0 uintptr, a1 uintptr) int32
var _e5rtProgramFunctionGetNumExternInoutsErr error

func tryE5rtProgramFunctionGetNumExternInouts(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtProgramFunctionGetNumExternInouts == nil {
		return 0, symbolCallError("e5rt_program_function_get_num_extern_inouts", "", _e5rtProgramFunctionGetNumExternInoutsErr)
	}
	return _e5rtProgramFunctionGetNumExternInouts(a0, a1), nil
}

// E5rtProgramFunctionGetNumExternInouts signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtProgramFunctionGetNumExternInputs signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtProgramFunctionGetNumExternOutputs signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtProgramFunctionLoadForExecution signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:481:9.
func E5rtProgramFunctionLoadForExecution(a0 uintptr) (int32, error) {
	return tryE5rtProgramFunctionLoadForExecution(a0)
}

var _e5rtProgramFunctionRelease func(a0 uintptr) int32
var _e5rtProgramFunctionReleaseErr error

func tryE5rtProgramFunctionRelease(a0 uintptr) (int32, error) {
	if _e5rtProgramFunctionRelease == nil {
		return 0, symbolCallError("e5rt_program_function_release", "", _e5rtProgramFunctionReleaseErr)
	}
	return _e5rtProgramFunctionRelease(a0), nil
}

// E5rtProgramFunctionRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtProgramFunctionRelease(a0 uintptr) (int32, error) {
	return tryE5rtProgramFunctionRelease(a0)
}

var _e5rtProgramFunctionRetainExternInputIOPort func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtProgramFunctionRetainExternInputIOPortErr error

func tryE5rtProgramFunctionRetainExternInputIOPort(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtProgramFunctionRetainExternInputIOPort == nil {
		return 0, symbolCallError("e5rt_program_function_retain_extern_input_io_port", "", _e5rtProgramFunctionRetainExternInputIOPortErr)
	}
	return _e5rtProgramFunctionRetainExternInputIOPort(a0, a1, a2), nil
}

// E5rtProgramFunctionRetainExternInputIOPort signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtProgramFunctionRetainExternOutputIOPort signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtProgramFunctionRetainInoutSurfaceDesc signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtProgramFunctionRetainInoutSurfaceDesc(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtProgramFunctionRetainInoutSurfaceDesc(a0, a1, a2)
}

var _e5rtProgramFunctionRetainInoutTensorDesc func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtProgramFunctionRetainInoutTensorDescErr error

func tryE5rtProgramFunctionRetainInoutTensorDesc(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtProgramFunctionRetainInoutTensorDesc == nil {
		return 0, symbolCallError("e5rt_program_function_retain_inout_tensor_desc", "", _e5rtProgramFunctionRetainInoutTensorDescErr)
	}
	return _e5rtProgramFunctionRetainInoutTensorDesc(a0, a1, a2), nil
}

// E5rtProgramFunctionRetainInoutTensorDesc signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtProgramFunctionRetainInoutTensorDesc(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtProgramFunctionRetainInoutTensorDesc(a0, a1, a2)
}

var _e5rtProgramFunctionRetainInputSurfaceDesc func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtProgramFunctionRetainInputSurfaceDescErr error

func tryE5rtProgramFunctionRetainInputSurfaceDesc(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtProgramFunctionRetainInputSurfaceDesc == nil {
		return 0, symbolCallError("e5rt_program_function_retain_input_surface_desc", "", _e5rtProgramFunctionRetainInputSurfaceDescErr)
	}
	return _e5rtProgramFunctionRetainInputSurfaceDesc(a0, a1, a2), nil
}

// E5rtProgramFunctionRetainInputSurfaceDesc signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtProgramFunctionRetainInputSurfaceDesc(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtProgramFunctionRetainInputSurfaceDesc(a0, a1, a2)
}

var _e5rtProgramFunctionRetainInputTensorDesc func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtProgramFunctionRetainInputTensorDescErr error

func tryE5rtProgramFunctionRetainInputTensorDesc(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtProgramFunctionRetainInputTensorDesc == nil {
		return 0, symbolCallError("e5rt_program_function_retain_input_tensor_desc", "", _e5rtProgramFunctionRetainInputTensorDescErr)
	}
	return _e5rtProgramFunctionRetainInputTensorDesc(a0, a1, a2), nil
}

// E5rtProgramFunctionRetainInputTensorDesc signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtProgramFunctionRetainInputTensorDesc(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
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

// E5rtProgramFunctionRetainOutputSurfaceDesc signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtProgramFunctionRetainOutputTensorDesc signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtProgramFunctionRetainOutputTensorDesc(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtProgramFunctionRetainOutputTensorDesc(a0, a1, a2)
}

var _e5rtProgramLibraryCreate func(out *uintptr, a1 uintptr) int32
var _e5rtProgramLibraryCreateErr error

func tryE5rtProgramLibraryCreate(out *uintptr, a1 uintptr) (int32, error) {
	if _e5rtProgramLibraryCreate == nil {
		return 0, symbolCallError("e5rt_program_library_create", "", _e5rtProgramLibraryCreateErr)
	}
	return _e5rtProgramLibraryCreate(out, a1), nil
}

// E5rtProgramLibraryCreate signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:444:9, where argument 0 is the out-parameter.
func E5rtProgramLibraryCreate(out *uintptr, a1 uintptr) (int32, error) {
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

// E5rtProgramLibraryGetBuildInfo signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::ProgramLibrary::GetBuildInfo takes 0.
func E5rtProgramLibraryGetBuildInfo(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtProgramLibraryGetBuildInfo(a0, out)
}

var _e5rtProgramLibraryGetE5BundlePath func(a0 uintptr, a1 uintptr) int32
var _e5rtProgramLibraryGetE5BundlePathErr error

func tryE5rtProgramLibraryGetE5BundlePath(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtProgramLibraryGetE5BundlePath == nil {
		return 0, symbolCallError("e5rt_program_library_get_e5_bundle_path", "", _e5rtProgramLibraryGetE5BundlePathErr)
	}
	return _e5rtProgramLibraryGetE5BundlePath(a0, a1), nil
}

// E5rtProgramLibraryGetE5BundlePath signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtProgramLibraryGetE5BundlePath(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtProgramLibraryGetE5BundlePath(a0, a1)
}

var _e5rtProgramLibraryGetFunctionMetadata func(a0 uintptr, a1 uintptr, out *uintptr) int32
var _e5rtProgramLibraryGetFunctionMetadataErr error

func tryE5rtProgramLibraryGetFunctionMetadata(a0 uintptr, a1 uintptr, out *uintptr) (int32, error) {
	if _e5rtProgramLibraryGetFunctionMetadata == nil {
		return 0, symbolCallError("e5rt_program_library_get_function_metadata", "", _e5rtProgramLibraryGetFunctionMetadataErr)
	}
	return _e5rtProgramLibraryGetFunctionMetadata(a0, a1, out), nil
}

// E5rtProgramLibraryGetFunctionMetadata signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 3, E5RT::ProgramLibrary::GetFunctionMetadata takes 1.
func E5rtProgramLibraryGetFunctionMetadata(a0 uintptr, a1 uintptr, out *uintptr) (int32, error) {
	return tryE5rtProgramLibraryGetFunctionMetadata(a0, a1, out)
}

var _e5rtProgramLibraryGetFunctionNames func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtProgramLibraryGetFunctionNamesErr error

func tryE5rtProgramLibraryGetFunctionNames(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtProgramLibraryGetFunctionNames == nil {
		return 0, symbolCallError("e5rt_program_library_get_function_names", "", _e5rtProgramLibraryGetFunctionNamesErr)
	}
	return _e5rtProgramLibraryGetFunctionNames(a0, a1, a2), nil
}

// E5rtProgramLibraryGetFunctionNames signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtProgramLibraryGetFunctionNames(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtProgramLibraryGetFunctionNames(a0, a1, a2)
}

var _e5rtProgramLibraryGetNumFunctions func(a0 uintptr, a1 uintptr) int32
var _e5rtProgramLibraryGetNumFunctionsErr error

func tryE5rtProgramLibraryGetNumFunctions(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtProgramLibraryGetNumFunctions == nil {
		return 0, symbolCallError("e5rt_program_library_get_num_functions", "", _e5rtProgramLibraryGetNumFunctionsErr)
	}
	return _e5rtProgramLibraryGetNumFunctions(a0, a1), nil
}

// E5rtProgramLibraryGetNumFunctions signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtProgramLibraryGetNumFunctions(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtProgramLibraryGetNumFunctions(a0, a1)
}

var _e5rtProgramLibraryGetSegmentationAnalytics func(a0 uintptr, out *uintptr) int32
var _e5rtProgramLibraryGetSegmentationAnalyticsErr error

func tryE5rtProgramLibraryGetSegmentationAnalytics(a0 uintptr, out *uintptr) (int32, error) {
	if _e5rtProgramLibraryGetSegmentationAnalytics == nil {
		return 0, symbolCallError("e5rt_program_library_get_segmentation_analytics", "", _e5rtProgramLibraryGetSegmentationAnalyticsErr)
	}
	return _e5rtProgramLibraryGetSegmentationAnalytics(a0, out), nil
}

// E5rtProgramLibraryGetSegmentationAnalytics signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: shim takes 2, E5RT::ProgramLibrary::GetSegmentationAnalytics takes 0.
func E5rtProgramLibraryGetSegmentationAnalytics(a0 uintptr, out *uintptr) (int32, error) {
	return tryE5rtProgramLibraryGetSegmentationAnalytics(a0, out)
}

var _e5rtProgramLibraryRelease func(a0 uintptr) int32
var _e5rtProgramLibraryReleaseErr error

func tryE5rtProgramLibraryRelease(a0 uintptr) (int32, error) {
	if _e5rtProgramLibraryRelease == nil {
		return 0, symbolCallError("e5rt_program_library_release", "", _e5rtProgramLibraryReleaseErr)
	}
	return _e5rtProgramLibraryRelease(a0), nil
}

// E5rtProgramLibraryRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtProgramLibraryRelease(a0 uintptr) (int32, error) {
	return tryE5rtProgramLibraryRelease(a0)
}

var _e5rtProgramLibraryRetainProgramFunction func(a0 uintptr, a1 uintptr, out *uintptr) int32
var _e5rtProgramLibraryRetainProgramFunctionErr error

func tryE5rtProgramLibraryRetainProgramFunction(a0 uintptr, a1 uintptr, out *uintptr) (int32, error) {
	if _e5rtProgramLibraryRetainProgramFunction == nil {
		return 0, symbolCallError("e5rt_program_library_retain_program_function", "", _e5rtProgramLibraryRetainProgramFunctionErr)
	}
	return _e5rtProgramLibraryRetainProgramFunction(a0, a1, out), nil
}

// E5rtProgramLibraryRetainProgramFunction signature verified against an observed call site at internal/signatureoracle/testdata/e5rt.go.txt:465:9, where argument 2 is the out-parameter.
func E5rtProgramLibraryRetainProgramFunction(a0 uintptr, a1 uintptr, out *uintptr) (int32, error) {
	return tryE5rtProgramLibraryRetainProgramFunction(a0, a1, out)
}

var _e5rtSurfaceDescCreate func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtSurfaceDescCreateErr error

func tryE5rtSurfaceDescCreate(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtSurfaceDescCreate == nil {
		return 0, symbolCallError("e5rt_surface_desc_create", "", _e5rtSurfaceDescCreateErr)
	}
	return _e5rtSurfaceDescCreate(a0, a1, a2, a3), nil
}

// E5rtSurfaceDescCreate signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtSurfaceDescCreate(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtSurfaceDescCreate(a0, a1, a2, a3)
}

var _e5rtSurfaceDescCreateFromOperandDesc func(a0 uintptr, a1 uintptr) int32
var _e5rtSurfaceDescCreateFromOperandDescErr error

func tryE5rtSurfaceDescCreateFromOperandDesc(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtSurfaceDescCreateFromOperandDesc == nil {
		return 0, symbolCallError("e5rt_surface_desc_create_from_operand_desc", "", _e5rtSurfaceDescCreateFromOperandDescErr)
	}
	return _e5rtSurfaceDescCreateFromOperandDesc(a0, a1), nil
}

// E5rtSurfaceDescCreateFromOperandDesc signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtSurfaceDescCreateFromOperandDesc(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtSurfaceDescCreateFromOperandDesc(a0, a1)
}

var _e5rtSurfaceDescCreateWithSlices func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) int32
var _e5rtSurfaceDescCreateWithSlicesErr error

func tryE5rtSurfaceDescCreateWithSlices(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	if _e5rtSurfaceDescCreateWithSlices == nil {
		return 0, symbolCallError("e5rt_surface_desc_create_with_slices", "", _e5rtSurfaceDescCreateWithSlicesErr)
	}
	return _e5rtSurfaceDescCreateWithSlices(a0, a1, a2, a3, a4), nil
}

// E5rtSurfaceDescCreateWithSlices signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtSurfaceDescCreateWithSlices(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	return tryE5rtSurfaceDescCreateWithSlices(a0, a1, a2, a3, a4)
}

var _e5rtSurfaceDescCreateWithStrides func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) int32
var _e5rtSurfaceDescCreateWithStridesErr error

func tryE5rtSurfaceDescCreateWithStrides(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	if _e5rtSurfaceDescCreateWithStrides == nil {
		return 0, symbolCallError("e5rt_surface_desc_create_with_strides", "", _e5rtSurfaceDescCreateWithStridesErr)
	}
	return _e5rtSurfaceDescCreateWithStrides(a0, a1, a2, a3, a4, a5), nil
}

// E5rtSurfaceDescCreateWithStrides signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtSurfaceDescCreateWithStrides(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	return tryE5rtSurfaceDescCreateWithStrides(a0, a1, a2, a3, a4, a5)
}

var _e5rtSurfaceDescCreateWithStridesAndSlices func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) int32
var _e5rtSurfaceDescCreateWithStridesAndSlicesErr error

func tryE5rtSurfaceDescCreateWithStridesAndSlices(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) (int32, error) {
	if _e5rtSurfaceDescCreateWithStridesAndSlices == nil {
		return 0, symbolCallError("e5rt_surface_desc_create_with_strides_and_slices", "", _e5rtSurfaceDescCreateWithStridesAndSlicesErr)
	}
	return _e5rtSurfaceDescCreateWithStridesAndSlices(a0, a1, a2, a3, a4, a5, a6), nil
}

// E5rtSurfaceDescCreateWithStridesAndSlices signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtSurfaceDescCreateWithStridesAndSlices(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) (int32, error) {
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

// E5rtSurfaceDescGetCustomRowStrides signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtSurfaceDescGetFormat signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtSurfaceDescGetFormat(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtSurfaceDescGetFormat(a0, a1)
}

var _e5rtSurfaceDescGetHeight func(a0 uintptr, a1 uintptr) int32
var _e5rtSurfaceDescGetHeightErr error

func tryE5rtSurfaceDescGetHeight(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtSurfaceDescGetHeight == nil {
		return 0, symbolCallError("e5rt_surface_desc_get_height", "", _e5rtSurfaceDescGetHeightErr)
	}
	return _e5rtSurfaceDescGetHeight(a0, a1), nil
}

// E5rtSurfaceDescGetHeight signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtSurfaceDescGetHeight(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtSurfaceDescGetHeight(a0, a1)
}

var _e5rtSurfaceDescGetPlaneCount func(a0 uintptr, a1 uintptr) int32
var _e5rtSurfaceDescGetPlaneCountErr error

func tryE5rtSurfaceDescGetPlaneCount(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtSurfaceDescGetPlaneCount == nil {
		return 0, symbolCallError("e5rt_surface_desc_get_plane_count", "", _e5rtSurfaceDescGetPlaneCountErr)
	}
	return _e5rtSurfaceDescGetPlaneCount(a0, a1), nil
}

// E5rtSurfaceDescGetPlaneCount signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtSurfaceDescGetSliceCount signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtSurfaceDescGetSliceCount(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtSurfaceDescGetSliceCount(a0, a1)
}

var _e5rtSurfaceDescGetWidth func(a0 uintptr, a1 uintptr) int32
var _e5rtSurfaceDescGetWidthErr error

func tryE5rtSurfaceDescGetWidth(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtSurfaceDescGetWidth == nil {
		return 0, symbolCallError("e5rt_surface_desc_get_width", "", _e5rtSurfaceDescGetWidthErr)
	}
	return _e5rtSurfaceDescGetWidth(a0, a1), nil
}

// E5rtSurfaceDescGetWidth signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtSurfaceDescGetWidth(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtSurfaceDescGetWidth(a0, a1)
}

var _e5rtSurfaceDescRelease func(a0 uintptr) int32
var _e5rtSurfaceDescReleaseErr error

func tryE5rtSurfaceDescRelease(a0 uintptr) (int32, error) {
	if _e5rtSurfaceDescRelease == nil {
		return 0, symbolCallError("e5rt_surface_desc_release", "", _e5rtSurfaceDescReleaseErr)
	}
	return _e5rtSurfaceDescRelease(a0), nil
}

// E5rtSurfaceDescRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtSurfaceDescRelease(a0 uintptr) (int32, error) {
	return tryE5rtSurfaceDescRelease(a0)
}

var _e5rtSurfaceFormatToCvpb4cc func(a0 uintptr, a1 uintptr) int32
var _e5rtSurfaceFormatToCvpb4ccErr error

func tryE5rtSurfaceFormatToCvpb4cc(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtSurfaceFormatToCvpb4cc == nil {
		return 0, symbolCallError("e5rt_surface_format_to_cvpb_4cc", "", _e5rtSurfaceFormatToCvpb4ccErr)
	}
	return _e5rtSurfaceFormatToCvpb4cc(a0, a1), nil
}

// E5rtSurfaceFormatToCvpb4cc signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtSurfaceFormatToCvpb4cc(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtSurfaceFormatToCvpb4cc(a0, a1)
}

var _e5rtSurfaceObjectAlloc func(out *uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtSurfaceObjectAllocErr error

func tryE5rtSurfaceObjectAlloc(out *uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtSurfaceObjectAlloc == nil {
		return 0, symbolCallError("e5rt_surface_object_alloc", "", _e5rtSurfaceObjectAllocErr)
	}
	return _e5rtSurfaceObjectAlloc(out, a1, a2), nil
}

// E5rtSurfaceObjectAlloc signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: wraps a constructor, so the out-parameter is the object being made.
func E5rtSurfaceObjectAlloc(out *uintptr, a1 uintptr, a2 uintptr) (int32, error) {
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

// E5rtSurfaceObjectCreateFromIosurface signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated); out-parameter position: wraps a constructor, so the out-parameter is the object being made.
func E5rtSurfaceObjectCreateFromIosurface(out *uintptr, a1 uintptr) (int32, error) {
	return tryE5rtSurfaceObjectCreateFromIosurface(out, a1)
}

var _e5rtSurfaceObjectGetIosurface func(a0 uintptr, a1 uintptr) int32
var _e5rtSurfaceObjectGetIosurfaceErr error

func tryE5rtSurfaceObjectGetIosurface(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtSurfaceObjectGetIosurface == nil {
		return 0, symbolCallError("e5rt_surface_object_get_iosurface", "", _e5rtSurfaceObjectGetIosurfaceErr)
	}
	return _e5rtSurfaceObjectGetIosurface(a0, a1), nil
}

// E5rtSurfaceObjectGetIosurface signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtSurfaceObjectGetIosurface(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtSurfaceObjectGetIosurface(a0, a1)
}

var _e5rtSurfaceObjectRelease func(a0 uintptr) int32
var _e5rtSurfaceObjectReleaseErr error

func tryE5rtSurfaceObjectRelease(a0 uintptr) (int32, error) {
	if _e5rtSurfaceObjectRelease == nil {
		return 0, symbolCallError("e5rt_surface_object_release", "", _e5rtSurfaceObjectReleaseErr)
	}
	return _e5rtSurfaceObjectRelease(a0), nil
}

// E5rtSurfaceObjectRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtSurfaceObjectRelease(a0 uintptr) (int32, error) {
	return tryE5rtSurfaceObjectRelease(a0)
}

var _e5rtTensorDescAllocBufferObject func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtTensorDescAllocBufferObjectErr error

func tryE5rtTensorDescAllocBufferObject(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtTensorDescAllocBufferObject == nil {
		return 0, symbolCallError("e5rt_tensor_desc_alloc_buffer_object", "", _e5rtTensorDescAllocBufferObjectErr)
	}
	return _e5rtTensorDescAllocBufferObject(a0, a1, a2, a3), nil
}

// E5rtTensorDescAllocBufferObject signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescAllocBufferObject(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtTensorDescAllocBufferObject(a0, a1, a2, a3)
}

var _e5rtTensorDescCreate func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtTensorDescCreateErr error

func tryE5rtTensorDescCreate(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtTensorDescCreate == nil {
		return 0, symbolCallError("e5rt_tensor_desc_create", "", _e5rtTensorDescCreateErr)
	}
	return _e5rtTensorDescCreate(a0, a1, a2, a3), nil
}

// E5rtTensorDescCreate signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescCreate(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtTensorDescCreate(a0, a1, a2, a3)
}

var _e5rtTensorDescCreateFromOperandDesc func(a0 uintptr, a1 uintptr) int32
var _e5rtTensorDescCreateFromOperandDescErr error

func tryE5rtTensorDescCreateFromOperandDesc(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtTensorDescCreateFromOperandDesc == nil {
		return 0, symbolCallError("e5rt_tensor_desc_create_from_operand_desc", "", _e5rtTensorDescCreateFromOperandDescErr)
	}
	return _e5rtTensorDescCreateFromOperandDesc(a0, a1), nil
}

// E5rtTensorDescCreateFromOperandDesc signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescCreateFromOperandDesc(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtTensorDescCreateFromOperandDesc(a0, a1)
}

var _e5rtTensorDescCreateMemoryObject func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtTensorDescCreateMemoryObjectErr error

func tryE5rtTensorDescCreateMemoryObject(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtTensorDescCreateMemoryObject == nil {
		return 0, symbolCallError("e5rt_tensor_desc_create_memory_object", "", _e5rtTensorDescCreateMemoryObjectErr)
	}
	return _e5rtTensorDescCreateMemoryObject(a0, a1, a2, a3), nil
}

// E5rtTensorDescCreateMemoryObject signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescCreateMemoryObject(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtTensorDescCreateMemoryObject(a0, a1, a2, a3)
}

var _e5rtTensorDescCreateSlice func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtTensorDescCreateSliceErr error

func tryE5rtTensorDescCreateSlice(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtTensorDescCreateSlice == nil {
		return 0, symbolCallError("e5rt_tensor_desc_create_slice", "", _e5rtTensorDescCreateSliceErr)
	}
	return _e5rtTensorDescCreateSlice(a0, a1, a2, a3), nil
}

// E5rtTensorDescCreateSlice signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescCreateSlice(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtTensorDescCreateSlice(a0, a1, a2, a3)
}

var _e5rtTensorDescCreateSliceWithLengths func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) int32
var _e5rtTensorDescCreateSliceWithLengthsErr error

func tryE5rtTensorDescCreateSliceWithLengths(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	if _e5rtTensorDescCreateSliceWithLengths == nil {
		return 0, symbolCallError("e5rt_tensor_desc_create_slice_with_lengths", "", _e5rtTensorDescCreateSliceWithLengthsErr)
	}
	return _e5rtTensorDescCreateSliceWithLengths(a0, a1, a2, a3, a4), nil
}

// E5rtTensorDescCreateSliceWithLengths signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescCreateSliceWithLengths(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	return tryE5rtTensorDescCreateSliceWithLengths(a0, a1, a2, a3, a4)
}

var _e5rtTensorDescCreateWithAlignments func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) int32
var _e5rtTensorDescCreateWithAlignmentsErr error

func tryE5rtTensorDescCreateWithAlignments(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	if _e5rtTensorDescCreateWithAlignments == nil {
		return 0, symbolCallError("e5rt_tensor_desc_create_with_alignments", "", _e5rtTensorDescCreateWithAlignmentsErr)
	}
	return _e5rtTensorDescCreateWithAlignments(a0, a1, a2, a3, a4), nil
}

// E5rtTensorDescCreateWithAlignments signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescCreateWithAlignments(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	return tryE5rtTensorDescCreateWithAlignments(a0, a1, a2, a3, a4)
}

var _e5rtTensorDescCreateWithStrides func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) int32
var _e5rtTensorDescCreateWithStridesErr error

func tryE5rtTensorDescCreateWithStrides(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	if _e5rtTensorDescCreateWithStrides == nil {
		return 0, symbolCallError("e5rt_tensor_desc_create_with_strides", "", _e5rtTensorDescCreateWithStridesErr)
	}
	return _e5rtTensorDescCreateWithStrides(a0, a1, a2, a3, a4), nil
}

// E5rtTensorDescCreateWithStrides signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescCreateWithStrides(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	return tryE5rtTensorDescCreateWithStrides(a0, a1, a2, a3, a4)
}

var _e5rtTensorDescDtypeAreEqual func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtTensorDescDtypeAreEqualErr error

func tryE5rtTensorDescDtypeAreEqual(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtTensorDescDtypeAreEqual == nil {
		return 0, symbolCallError("e5rt_tensor_desc_dtype_are_equal", "", _e5rtTensorDescDtypeAreEqualErr)
	}
	return _e5rtTensorDescDtypeAreEqual(a0, a1, a2), nil
}

// E5rtTensorDescDtypeAreEqual signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescDtypeAreEqual(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtTensorDescDtypeAreEqual(a0, a1, a2)
}

var _e5rtTensorDescDtypeCreate func(a0 uintptr, a1 uintptr) int32
var _e5rtTensorDescDtypeCreateErr error

func tryE5rtTensorDescDtypeCreate(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtTensorDescDtypeCreate == nil {
		return 0, symbolCallError("e5rt_tensor_desc_dtype_create", "", _e5rtTensorDescDtypeCreateErr)
	}
	return _e5rtTensorDescDtypeCreate(a0, a1), nil
}

// E5rtTensorDescDtypeCreate signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescDtypeCreate(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtTensorDescDtypeCreate(a0, a1)
}

var _e5rtTensorDescDtypeGetComponentDtype func(a0 uintptr, a1 uintptr) int32
var _e5rtTensorDescDtypeGetComponentDtypeErr error

func tryE5rtTensorDescDtypeGetComponentDtype(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtTensorDescDtypeGetComponentDtype == nil {
		return 0, symbolCallError("e5rt_tensor_desc_dtype_get_component_dtype", "", _e5rtTensorDescDtypeGetComponentDtypeErr)
	}
	return _e5rtTensorDescDtypeGetComponentDtype(a0, a1), nil
}

// E5rtTensorDescDtypeGetComponentDtype signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescDtypeGetComponentDtype(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtTensorDescDtypeGetComponentDtype(a0, a1)
}

var _e5rtTensorDescDtypeGetComponentPack func(a0 uintptr, a1 uintptr) int32
var _e5rtTensorDescDtypeGetComponentPackErr error

func tryE5rtTensorDescDtypeGetComponentPack(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtTensorDescDtypeGetComponentPack == nil {
		return 0, symbolCallError("e5rt_tensor_desc_dtype_get_component_pack", "", _e5rtTensorDescDtypeGetComponentPackErr)
	}
	return _e5rtTensorDescDtypeGetComponentPack(a0, a1), nil
}

// E5rtTensorDescDtypeGetComponentPack signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescDtypeGetComponentPack(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtTensorDescDtypeGetComponentPack(a0, a1)
}

var _e5rtTensorDescDtypeGetComponentSize func(a0 uintptr, a1 uintptr) int32
var _e5rtTensorDescDtypeGetComponentSizeErr error

func tryE5rtTensorDescDtypeGetComponentSize(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtTensorDescDtypeGetComponentSize == nil {
		return 0, symbolCallError("e5rt_tensor_desc_dtype_get_component_size", "", _e5rtTensorDescDtypeGetComponentSizeErr)
	}
	return _e5rtTensorDescDtypeGetComponentSize(a0, a1), nil
}

// E5rtTensorDescDtypeGetComponentSize signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescDtypeGetComponentSize(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtTensorDescDtypeGetComponentSize(a0, a1)
}

var _e5rtTensorDescDtypeGetElementSize func(a0 uintptr, a1 uintptr) int32
var _e5rtTensorDescDtypeGetElementSizeErr error

func tryE5rtTensorDescDtypeGetElementSize(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtTensorDescDtypeGetElementSize == nil {
		return 0, symbolCallError("e5rt_tensor_desc_dtype_get_element_size", "", _e5rtTensorDescDtypeGetElementSizeErr)
	}
	return _e5rtTensorDescDtypeGetElementSize(a0, a1), nil
}

// E5rtTensorDescDtypeGetElementSize signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescDtypeGetElementSize(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtTensorDescDtypeGetElementSize(a0, a1)
}

var _e5rtTensorDescDtypeGetNumComponents func(a0 uintptr, a1 uintptr) int32
var _e5rtTensorDescDtypeGetNumComponentsErr error

func tryE5rtTensorDescDtypeGetNumComponents(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtTensorDescDtypeGetNumComponents == nil {
		return 0, symbolCallError("e5rt_tensor_desc_dtype_get_num_components", "", _e5rtTensorDescDtypeGetNumComponentsErr)
	}
	return _e5rtTensorDescDtypeGetNumComponents(a0, a1), nil
}

// E5rtTensorDescDtypeGetNumComponents signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescDtypeGetNumComponents(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtTensorDescDtypeGetNumComponents(a0, a1)
}

var _e5rtTensorDescDtypeRelease func(a0 uintptr) int32
var _e5rtTensorDescDtypeReleaseErr error

func tryE5rtTensorDescDtypeRelease(a0 uintptr) (int32, error) {
	if _e5rtTensorDescDtypeRelease == nil {
		return 0, symbolCallError("e5rt_tensor_desc_dtype_release", "", _e5rtTensorDescDtypeReleaseErr)
	}
	return _e5rtTensorDescDtypeRelease(a0), nil
}

// E5rtTensorDescDtypeRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescDtypeRelease(a0 uintptr) (int32, error) {
	return tryE5rtTensorDescDtypeRelease(a0)
}

var _e5rtTensorDescGetByteOffset func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtTensorDescGetByteOffsetErr error

func tryE5rtTensorDescGetByteOffset(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtTensorDescGetByteOffset == nil {
		return 0, symbolCallError("e5rt_tensor_desc_get_byte_offset", "", _e5rtTensorDescGetByteOffsetErr)
	}
	return _e5rtTensorDescGetByteOffset(a0, a1, a2, a3), nil
}

// E5rtTensorDescGetByteOffset signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescGetByteOffset(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtTensorDescGetByteOffset(a0, a1, a2, a3)
}

var _e5rtTensorDescGetDimensionLength func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtTensorDescGetDimensionLengthErr error

func tryE5rtTensorDescGetDimensionLength(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtTensorDescGetDimensionLength == nil {
		return 0, symbolCallError("e5rt_tensor_desc_get_dimension_length", "", _e5rtTensorDescGetDimensionLengthErr)
	}
	return _e5rtTensorDescGetDimensionLength(a0, a1, a2), nil
}

// E5rtTensorDescGetDimensionLength signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtTensorDescGetDimensionStride signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescGetDimensionStride(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtTensorDescGetDimensionStride(a0, a1, a2)
}

var _e5rtTensorDescGetNumElements func(a0 uintptr, a1 uintptr) int32
var _e5rtTensorDescGetNumElementsErr error

func tryE5rtTensorDescGetNumElements(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtTensorDescGetNumElements == nil {
		return 0, symbolCallError("e5rt_tensor_desc_get_num_elements", "", _e5rtTensorDescGetNumElementsErr)
	}
	return _e5rtTensorDescGetNumElements(a0, a1), nil
}

// E5rtTensorDescGetNumElements signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescGetNumElements(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtTensorDescGetNumElements(a0, a1)
}

var _e5rtTensorDescGetRank func(a0 uintptr, a1 uintptr) int32
var _e5rtTensorDescGetRankErr error

func tryE5rtTensorDescGetRank(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtTensorDescGetRank == nil {
		return 0, symbolCallError("e5rt_tensor_desc_get_rank", "", _e5rtTensorDescGetRankErr)
	}
	return _e5rtTensorDescGetRank(a0, a1), nil
}

// E5rtTensorDescGetRank signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescGetRank(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtTensorDescGetRank(a0, a1)
}

var _e5rtTensorDescGetShape func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtTensorDescGetShapeErr error

func tryE5rtTensorDescGetShape(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtTensorDescGetShape == nil {
		return 0, symbolCallError("e5rt_tensor_desc_get_shape", "", _e5rtTensorDescGetShapeErr)
	}
	return _e5rtTensorDescGetShape(a0, a1, a2), nil
}

// E5rtTensorDescGetShape signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescGetShape(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtTensorDescGetShape(a0, a1, a2)
}

var _e5rtTensorDescGetSize func(a0 uintptr, a1 uintptr) int32
var _e5rtTensorDescGetSizeErr error

func tryE5rtTensorDescGetSize(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtTensorDescGetSize == nil {
		return 0, symbolCallError("e5rt_tensor_desc_get_size", "", _e5rtTensorDescGetSizeErr)
	}
	return _e5rtTensorDescGetSize(a0, a1), nil
}

// E5rtTensorDescGetSize signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescGetSize(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtTensorDescGetSize(a0, a1)
}

var _e5rtTensorDescGetStrides func(a0 uintptr, a1 uintptr, a2 uintptr) int32
var _e5rtTensorDescGetStridesErr error

func tryE5rtTensorDescGetStrides(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	if _e5rtTensorDescGetStrides == nil {
		return 0, symbolCallError("e5rt_tensor_desc_get_strides", "", _e5rtTensorDescGetStridesErr)
	}
	return _e5rtTensorDescGetStrides(a0, a1, a2), nil
}

// E5rtTensorDescGetStrides signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescGetStrides(a0 uintptr, a1 uintptr, a2 uintptr) (int32, error) {
	return tryE5rtTensorDescGetStrides(a0, a1, a2)
}

var _e5rtTensorDescHasKnownShape func(a0 uintptr, a1 uintptr) int32
var _e5rtTensorDescHasKnownShapeErr error

func tryE5rtTensorDescHasKnownShape(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtTensorDescHasKnownShape == nil {
		return 0, symbolCallError("e5rt_tensor_desc_has_known_shape", "", _e5rtTensorDescHasKnownShapeErr)
	}
	return _e5rtTensorDescHasKnownShape(a0, a1), nil
}

// E5rtTensorDescHasKnownShape signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescHasKnownShape(a0 uintptr, a1 uintptr) (int32, error) {
	return tryE5rtTensorDescHasKnownShape(a0, a1)
}

var _e5rtTensorDescRelease func(a0 uintptr) int32
var _e5rtTensorDescReleaseErr error

func tryE5rtTensorDescRelease(a0 uintptr) (int32, error) {
	if _e5rtTensorDescRelease == nil {
		return 0, symbolCallError("e5rt_tensor_desc_release", "", _e5rtTensorDescReleaseErr)
	}
	return _e5rtTensorDescRelease(a0), nil
}

// E5rtTensorDescRelease signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorDescRelease(a0 uintptr) (int32, error) {
	return tryE5rtTensorDescRelease(a0)
}

var _e5rtTensorDescRetainDtype func(a0 uintptr, a1 uintptr) int32
var _e5rtTensorDescRetainDtypeErr error

func tryE5rtTensorDescRetainDtype(a0 uintptr, a1 uintptr) (int32, error) {
	if _e5rtTensorDescRetainDtype == nil {
		return 0, symbolCallError("e5rt_tensor_desc_retain_dtype", "", _e5rtTensorDescRetainDtypeErr)
	}
	return _e5rtTensorDescRetainDtype(a0, a1), nil
}

// E5rtTensorDescRetainDtype signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtTensorUtilsAreTensorsEqual signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtTensorUtilsCastFromFp16ToFp32 signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtTensorUtilsCastFromFp32ToFp16 signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
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

// E5rtTensorUtilsCopyTensor signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorUtilsCopyTensor(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtTensorUtilsCopyTensor(a0, a1, a2, a3)
}

var _e5rtTensorUtilsDequantizeFromS8ToFp32 func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtTensorUtilsDequantizeFromS8ToFp32Err error

func tryE5rtTensorUtilsDequantizeFromS8ToFp32(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtTensorUtilsDequantizeFromS8ToFp32 == nil {
		return 0, symbolCallError("e5rt_tensor_utils_dequantize_from_s8_to_fp32", "", _e5rtTensorUtilsDequantizeFromS8ToFp32Err)
	}
	return _e5rtTensorUtilsDequantizeFromS8ToFp32(a0, a1, a2, a3), nil
}

// E5rtTensorUtilsDequantizeFromS8ToFp32 signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorUtilsDequantizeFromS8ToFp32(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtTensorUtilsDequantizeFromS8ToFp32(a0, a1, a2, a3)
}

var _e5rtTensorUtilsDequantizeFromU8ToFp32 func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtTensorUtilsDequantizeFromU8ToFp32Err error

func tryE5rtTensorUtilsDequantizeFromU8ToFp32(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtTensorUtilsDequantizeFromU8ToFp32 == nil {
		return 0, symbolCallError("e5rt_tensor_utils_dequantize_from_u8_to_fp32", "", _e5rtTensorUtilsDequantizeFromU8ToFp32Err)
	}
	return _e5rtTensorUtilsDequantizeFromU8ToFp32(a0, a1, a2, a3), nil
}

// E5rtTensorUtilsDequantizeFromU8ToFp32 signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorUtilsDequantizeFromU8ToFp32(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtTensorUtilsDequantizeFromU8ToFp32(a0, a1, a2, a3)
}

var _e5rtTensorUtilsGetFp16Element func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) int32
var _e5rtTensorUtilsGetFp16ElementErr error

func tryE5rtTensorUtilsGetFp16Element(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	if _e5rtTensorUtilsGetFp16Element == nil {
		return 0, symbolCallError("e5rt_tensor_utils_get_fp16_element", "", _e5rtTensorUtilsGetFp16ElementErr)
	}
	return _e5rtTensorUtilsGetFp16Element(a0, a1, a2, a3, a4, a5), nil
}

// E5rtTensorUtilsGetFp16Element signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorUtilsGetFp16Element(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	return tryE5rtTensorUtilsGetFp16Element(a0, a1, a2, a3, a4, a5)
}

var _e5rtTensorUtilsGetFp32Element func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) int32
var _e5rtTensorUtilsGetFp32ElementErr error

func tryE5rtTensorUtilsGetFp32Element(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	if _e5rtTensorUtilsGetFp32Element == nil {
		return 0, symbolCallError("e5rt_tensor_utils_get_fp32_element", "", _e5rtTensorUtilsGetFp32ElementErr)
	}
	return _e5rtTensorUtilsGetFp32Element(a0, a1, a2, a3, a4, a5), nil
}

// E5rtTensorUtilsGetFp32Element signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorUtilsGetFp32Element(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	return tryE5rtTensorUtilsGetFp32Element(a0, a1, a2, a3, a4, a5)
}

var _e5rtTensorUtilsGetS8Element func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) int32
var _e5rtTensorUtilsGetS8ElementErr error

func tryE5rtTensorUtilsGetS8Element(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	if _e5rtTensorUtilsGetS8Element == nil {
		return 0, symbolCallError("e5rt_tensor_utils_get_s8_element", "", _e5rtTensorUtilsGetS8ElementErr)
	}
	return _e5rtTensorUtilsGetS8Element(a0, a1, a2, a3, a4, a5), nil
}

// E5rtTensorUtilsGetS8Element signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorUtilsGetS8Element(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	return tryE5rtTensorUtilsGetS8Element(a0, a1, a2, a3, a4, a5)
}

var _e5rtTensorUtilsGetU8Element func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) int32
var _e5rtTensorUtilsGetU8ElementErr error

func tryE5rtTensorUtilsGetU8Element(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	if _e5rtTensorUtilsGetU8Element == nil {
		return 0, symbolCallError("e5rt_tensor_utils_get_u8_element", "", _e5rtTensorUtilsGetU8ElementErr)
	}
	return _e5rtTensorUtilsGetU8Element(a0, a1, a2, a3, a4, a5), nil
}

// E5rtTensorUtilsGetU8Element signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorUtilsGetU8Element(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	return tryE5rtTensorUtilsGetU8Element(a0, a1, a2, a3, a4, a5)
}

var _e5rtTensorUtilsQuantizeFromFp32ToU8 func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) int32
var _e5rtTensorUtilsQuantizeFromFp32ToU8Err error

func tryE5rtTensorUtilsQuantizeFromFp32ToU8(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	if _e5rtTensorUtilsQuantizeFromFp32ToU8 == nil {
		return 0, symbolCallError("e5rt_tensor_utils_quantize_from_fp32_to_u8", "", _e5rtTensorUtilsQuantizeFromFp32ToU8Err)
	}
	return _e5rtTensorUtilsQuantizeFromFp32ToU8(a0, a1, a2, a3), nil
}

// E5rtTensorUtilsQuantizeFromFp32ToU8 signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorUtilsQuantizeFromFp32ToU8(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (int32, error) {
	return tryE5rtTensorUtilsQuantizeFromFp32ToU8(a0, a1, a2, a3)
}

var _e5rtTensorUtilsSetFp16Element func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) int32
var _e5rtTensorUtilsSetFp16ElementErr error

func tryE5rtTensorUtilsSetFp16Element(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	if _e5rtTensorUtilsSetFp16Element == nil {
		return 0, symbolCallError("e5rt_tensor_utils_set_fp16_element", "", _e5rtTensorUtilsSetFp16ElementErr)
	}
	return _e5rtTensorUtilsSetFp16Element(a0, a1, a2, a3, a4), nil
}

// E5rtTensorUtilsSetFp16Element signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorUtilsSetFp16Element(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	return tryE5rtTensorUtilsSetFp16Element(a0, a1, a2, a3, a4)
}

var _e5rtTensorUtilsSetFp32Element func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) int32
var _e5rtTensorUtilsSetFp32ElementErr error

func tryE5rtTensorUtilsSetFp32Element(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	if _e5rtTensorUtilsSetFp32Element == nil {
		return 0, symbolCallError("e5rt_tensor_utils_set_fp32_element", "", _e5rtTensorUtilsSetFp32ElementErr)
	}
	return _e5rtTensorUtilsSetFp32Element(a0, a1, a2, a3, a4), nil
}

// E5rtTensorUtilsSetFp32Element signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorUtilsSetFp32Element(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr) (int32, error) {
	return tryE5rtTensorUtilsSetFp32Element(a0, a1, a2, a3, a4)
}

var _e5rtTensorUtilsSetS8Element func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) int32
var _e5rtTensorUtilsSetS8ElementErr error

func tryE5rtTensorUtilsSetS8Element(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	if _e5rtTensorUtilsSetS8Element == nil {
		return 0, symbolCallError("e5rt_tensor_utils_set_s8_element", "", _e5rtTensorUtilsSetS8ElementErr)
	}
	return _e5rtTensorUtilsSetS8Element(a0, a1, a2, a3, a4, a5), nil
}

// E5rtTensorUtilsSetS8Element signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorUtilsSetS8Element(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	return tryE5rtTensorUtilsSetS8Element(a0, a1, a2, a3, a4, a5)
}

var _e5rtTensorUtilsSetU8Element func(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) int32
var _e5rtTensorUtilsSetU8ElementErr error

func tryE5rtTensorUtilsSetU8Element(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	if _e5rtTensorUtilsSetU8Element == nil {
		return 0, symbolCallError("e5rt_tensor_utils_set_u8_element", "", _e5rtTensorUtilsSetU8ElementErr)
	}
	return _e5rtTensorUtilsSetU8Element(a0, a1, a2, a3, a4, a5), nil
}

// E5rtTensorUtilsSetU8Element signature constrained, not determined, by reading x0-x7 liveness in the shim prologue (shimtext, 75.9% accurate on this image, uncorroborated).
func E5rtTensorUtilsSetU8Element(a0 uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr) (int32, error) {
	return tryE5rtTensorUtilsSetU8Element(a0, a1, a2, a3, a4, a5)
}

var _espressoContextDestroy func(ctx EspressoContext)
var _espressoContextDestroyErr error

func tryEspressoContextDestroy(ctx EspressoContext) error {
	if _espressoContextDestroy == nil {
		return symbolCallError("espresso_context_destroy", "", _espressoContextDestroyErr)
	}
	_espressoContextDestroy(ctx)
	return nil
}

// EspressoContextDestroy signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
func EspressoContextDestroy(ctx EspressoContext) error {
	return tryEspressoContextDestroy(ctx)
}

var _espressoCreateContext func(platform int32, options int32) EspressoContext
var _espressoCreateContextErr error

func tryEspressoCreateContext(platform int32, options int32) (EspressoContext, error) {
	if _espressoCreateContext == nil {
		return EspressoContext{}, symbolCallError("espresso_create_context", "", _espressoCreateContextErr)
	}
	return _espressoCreateContext(platform, options), nil
}

// EspressoCreateContext signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
func EspressoCreateContext(platform int32, options int32) (EspressoContext, error) {
	return tryEspressoCreateContext(platform, options)
}

var _espressoCreatePlan func(ctx EspressoContext, platform int32) EspressoPlan
var _espressoCreatePlanErr error

func tryEspressoCreatePlan(ctx EspressoContext, platform int32) (EspressoPlan, error) {
	if _espressoCreatePlan == nil {
		return *new(EspressoPlan), symbolCallError("espresso_create_plan", "", _espressoCreatePlanErr)
	}
	return _espressoCreatePlan(ctx, platform), nil
}

// EspressoCreatePlan signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
func EspressoCreatePlan(ctx EspressoContext, platform int32) (EspressoPlan, error) {
	return tryEspressoCreatePlan(ctx, platform)
}

var _espressoGetVersionString func() uintptr
var _espressoGetVersionStringErr error

func tryEspressoGetVersionString() (uintptr, error) {
	if _espressoGetVersionString == nil {
		return 0, symbolCallError("espresso_get_version_string", "", _espressoGetVersionStringErr)
	}
	return _espressoGetVersionString(), nil
}

// EspressoGetVersionString signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
func EspressoGetVersionString() (uintptr, error) {
	return tryEspressoGetVersionString()
}

var _espressoPlanBuild func(plan EspressoPlan) int32
var _espressoPlanBuildErr error

func tryEspressoPlanBuild(plan EspressoPlan) (int32, error) {
	if _espressoPlanBuild == nil {
		return 0, symbolCallError("espresso_plan_build", "", _espressoPlanBuildErr)
	}
	return _espressoPlanBuild(plan), nil
}

// EspressoPlanBuild signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
func EspressoPlanBuild(plan EspressoPlan) (int32, error) {
	return tryEspressoPlanBuild(plan)
}

var _espressoPlanDestroy func(plan EspressoPlan)
var _espressoPlanDestroyErr error

func tryEspressoPlanDestroy(plan EspressoPlan) error {
	if _espressoPlanDestroy == nil {
		return symbolCallError("espresso_plan_destroy", "", _espressoPlanDestroyErr)
	}
	_espressoPlanDestroy(plan)
	return nil
}

// EspressoPlanDestroy signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
func EspressoPlanDestroy(plan EspressoPlan) error {
	return tryEspressoPlanDestroy(plan)
}

var _espressoPlanExecuteSync func(plan EspressoPlan) int32
var _espressoPlanExecuteSyncErr error

func tryEspressoPlanExecuteSync(plan EspressoPlan) (int32, error) {
	if _espressoPlanExecuteSync == nil {
		return 0, symbolCallError("espresso_plan_execute_sync", "", _espressoPlanExecuteSyncErr)
	}
	return _espressoPlanExecuteSync(plan), nil
}

// EspressoPlanExecuteSync signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
func EspressoPlanExecuteSync(plan EspressoPlan) (int32, error) {
	return tryEspressoPlanExecuteSync(plan)
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
	registerFunc(&_e5rtTensorUtilsDequantizeFromS8ToFp32, &_e5rtTensorUtilsDequantizeFromS8ToFp32Err, frameworkHandle, "e5rt_tensor_utils_dequantize_from_s8_to_fp32", "")
	registerFunc(&_e5rtTensorUtilsDequantizeFromU8ToFp32, &_e5rtTensorUtilsDequantizeFromU8ToFp32Err, frameworkHandle, "e5rt_tensor_utils_dequantize_from_u8_to_fp32", "")
	registerFunc(&_e5rtTensorUtilsGetFp16Element, &_e5rtTensorUtilsGetFp16ElementErr, frameworkHandle, "e5rt_tensor_utils_get_fp16_element", "")
	registerFunc(&_e5rtTensorUtilsGetFp32Element, &_e5rtTensorUtilsGetFp32ElementErr, frameworkHandle, "e5rt_tensor_utils_get_fp32_element", "")
	registerFunc(&_e5rtTensorUtilsGetS8Element, &_e5rtTensorUtilsGetS8ElementErr, frameworkHandle, "e5rt_tensor_utils_get_s8_element", "")
	registerFunc(&_e5rtTensorUtilsGetU8Element, &_e5rtTensorUtilsGetU8ElementErr, frameworkHandle, "e5rt_tensor_utils_get_u8_element", "")
	registerFunc(&_e5rtTensorUtilsQuantizeFromFp32ToU8, &_e5rtTensorUtilsQuantizeFromFp32ToU8Err, frameworkHandle, "e5rt_tensor_utils_quantize_from_fp32_to_u8", "")
	registerFunc(&_e5rtTensorUtilsSetFp16Element, &_e5rtTensorUtilsSetFp16ElementErr, frameworkHandle, "e5rt_tensor_utils_set_fp16_element", "")
	registerFunc(&_e5rtTensorUtilsSetFp32Element, &_e5rtTensorUtilsSetFp32ElementErr, frameworkHandle, "e5rt_tensor_utils_set_fp32_element", "")
	registerFunc(&_e5rtTensorUtilsSetS8Element, &_e5rtTensorUtilsSetS8ElementErr, frameworkHandle, "e5rt_tensor_utils_set_s8_element", "")
	registerFunc(&_e5rtTensorUtilsSetU8Element, &_e5rtTensorUtilsSetU8ElementErr, frameworkHandle, "e5rt_tensor_utils_set_u8_element", "")
	registerFunc(&_espressoContextDestroy, &_espressoContextDestroyErr, frameworkHandle, "espresso_context_destroy", "")
	registerFunc(&_espressoCreateContext, &_espressoCreateContextErr, frameworkHandle, "espresso_create_context", "")
	registerFunc(&_espressoCreatePlan, &_espressoCreatePlanErr, frameworkHandle, "espresso_create_plan", "")
	registerFunc(&_espressoGetVersionString, &_espressoGetVersionStringErr, frameworkHandle, "espresso_get_version_string", "")
	registerFunc(&_espressoPlanBuild, &_espressoPlanBuildErr, frameworkHandle, "espresso_plan_build", "")
	registerFunc(&_espressoPlanDestroy, &_espressoPlanDestroyErr, frameworkHandle, "espresso_plan_destroy", "")
	registerFunc(&_espressoPlanExecuteSync, &_espressoPlanExecuteSyncErr, frameworkHandle, "espresso_plan_execute_sync", "")
}

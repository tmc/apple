// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoANEIOSurface] class.
var (
	_EspressoANEIOSurfaceClass     EspressoANEIOSurfaceClass
	_EspressoANEIOSurfaceClassOnce sync.Once
)

func getEspressoANEIOSurfaceClass() EspressoANEIOSurfaceClass {
	_EspressoANEIOSurfaceClassOnce.Do(func() {
		_EspressoANEIOSurfaceClass = EspressoANEIOSurfaceClass{class: objc.GetClass("EspressoANEIOSurface")}
	})
	return _EspressoANEIOSurfaceClass
}

// GetEspressoANEIOSurfaceClass returns the class object for EspressoANEIOSurface.
func GetEspressoANEIOSurfaceClass() EspressoANEIOSurfaceClass {
	return getEspressoANEIOSurfaceClass()
}

type EspressoANEIOSurfaceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoANEIOSurfaceClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoANEIOSurfaceClass) Alloc() EspressoANEIOSurface {
	rv := objc.SendIfResponds[EspressoANEIOSurface](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [EspressoANEIOSurface.Ane_io_surfaceForMultiBufferFrame]
//   - [EspressoANEIOSurface.CheckIfMatches]
//   - [EspressoANEIOSurface.CheckIfMatchesIOSurface]
//   - [EspressoANEIOSurface.Cleanup]
//   - [EspressoANEIOSurface.CreateIOSurfaceWithExtraProperties]
//   - [EspressoANEIOSurface.DoNonLazyAllocation]
//   - [EspressoANEIOSurface.External_storage_blob_for_aliasing_mem]
//   - [EspressoANEIOSurface.SetExternal_storage_blob_for_aliasing_mem]
//   - [EspressoANEIOSurface.IoSurfaceForMultiBufferFrame]
//   - [EspressoANEIOSurface.IoSurfaceForMultiBufferFrameNoLazyForTesting]
//   - [EspressoANEIOSurface.LazilyAutoCreateSurfaceForFrame]
//   - [EspressoANEIOSurface.MetalBufferWithDeviceMultiBufferFrame]
//   - [EspressoANEIOSurface.NFrames]
//   - [EspressoANEIOSurface.PixelFormat]
//   - [EspressoANEIOSurface.ResizeForMultipleAsyncBuffers]
//   - [EspressoANEIOSurface.RestoreInternalStorage]
//   - [EspressoANEIOSurface.RestoreInternalStorageForAllMultiBufferFrames]
//   - [EspressoANEIOSurface.SetExternalStorageIoSurface]
//   - [EspressoANEIOSurface.InitWithIOSurfacePropertiesAndPixelFormats]
type EspressoANEIOSurface struct {
	objectivec.Object
}

// EspressoANEIOSurfaceFromID constructs a [EspressoANEIOSurface] from an objc.ID.
func EspressoANEIOSurfaceFromID(id objc.ID) EspressoANEIOSurface {
	return EspressoANEIOSurface{objectivec.Object{ID: id}}
}

// Ensure EspressoANEIOSurface implements IEspressoANEIOSurface.
var _ IEspressoANEIOSurface = EspressoANEIOSurface{}

// An interface definition for the [EspressoANEIOSurface] class.
//
// # Methods
//
//   - [IEspressoANEIOSurface.Ane_io_surfaceForMultiBufferFrame]
//   - [IEspressoANEIOSurface.CheckIfMatches]
//   - [IEspressoANEIOSurface.CheckIfMatchesIOSurface]
//   - [IEspressoANEIOSurface.Cleanup]
//   - [IEspressoANEIOSurface.CreateIOSurfaceWithExtraProperties]
//   - [IEspressoANEIOSurface.DoNonLazyAllocation]
//   - [IEspressoANEIOSurface.External_storage_blob_for_aliasing_mem]
//   - [IEspressoANEIOSurface.SetExternal_storage_blob_for_aliasing_mem]
//   - [IEspressoANEIOSurface.IoSurfaceForMultiBufferFrame]
//   - [IEspressoANEIOSurface.IoSurfaceForMultiBufferFrameNoLazyForTesting]
//   - [IEspressoANEIOSurface.LazilyAutoCreateSurfaceForFrame]
//   - [IEspressoANEIOSurface.MetalBufferWithDeviceMultiBufferFrame]
//   - [IEspressoANEIOSurface.NFrames]
//   - [IEspressoANEIOSurface.PixelFormat]
//   - [IEspressoANEIOSurface.ResizeForMultipleAsyncBuffers]
//   - [IEspressoANEIOSurface.RestoreInternalStorage]
//   - [IEspressoANEIOSurface.RestoreInternalStorageForAllMultiBufferFrames]
//   - [IEspressoANEIOSurface.SetExternalStorageIoSurface]
//   - [IEspressoANEIOSurface.InitWithIOSurfacePropertiesAndPixelFormats]
type IEspressoANEIOSurface interface {
	objectivec.IObject

	// Topic: Methods

	Ane_io_surfaceForMultiBufferFrame(frame uint64) objectivec.IObject
	CheckIfMatches(matches corevideo.CVImageBufferRef) bool
	CheckIfMatchesIOSurface(iOSurface iosurface.IOSurfaceRef) bool
	Cleanup()
	CreateIOSurfaceWithExtraProperties(properties objectivec.IObject) iosurface.IOSurfaceRef
	DoNonLazyAllocation(allocation objectivec.IObject)
	External_storage_blob_for_aliasing_mem() unsafe.Pointer
	SetExternal_storage_blob_for_aliasing_mem(value unsafe.Pointer)
	IoSurfaceForMultiBufferFrame(frame uint64) iosurface.IOSurfaceRef
	IoSurfaceForMultiBufferFrameNoLazyForTesting(testing uint64) iosurface.IOSurfaceRef
	LazilyAutoCreateSurfaceForFrame(frame uint64)
	MetalBufferWithDeviceMultiBufferFrame(device objectivec.IObject, frame uint64) metal.MTLBuffer
	NFrames() uint64
	PixelFormat() uint32
	ResizeForMultipleAsyncBuffers(buffers uint64)
	RestoreInternalStorage(storage uint64)
	RestoreInternalStorageForAllMultiBufferFrames()
	SetExternalStorageIoSurface(storage uint64, surface iosurface.IOSurfaceRef)
	InitWithIOSurfacePropertiesAndPixelFormats(properties objectivec.IObject, formats objectivec.IObject) EspressoANEIOSurface
}

// Init initializes the instance.
func (e EspressoANEIOSurface) Init() EspressoANEIOSurface {
	rv := objc.SendIfResponds[EspressoANEIOSurface](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoANEIOSurface) Autorelease() EspressoANEIOSurface {
	rv := objc.SendIfResponds[EspressoANEIOSurface](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoANEIOSurface creates a new EspressoANEIOSurface instance.
func NewEspressoANEIOSurface() EspressoANEIOSurface {
	class := getEspressoANEIOSurfaceClass()
	rv := objc.SendIfResponds[EspressoANEIOSurface](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewEspressoANEIOSurfaceWithIOSurfacePropertiesAndPixelFormats(properties objectivec.IObject, formats objectivec.IObject) EspressoANEIOSurface {
	instance := getEspressoANEIOSurfaceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithIOSurfaceProperties:andPixelFormats:"), properties, formats)
	return EspressoANEIOSurfaceFromID(rv)
}

func (e EspressoANEIOSurface) Ane_io_surfaceForMultiBufferFrame(frame uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("ane_io_surfaceForMultiBufferFrame:"), frame)
	return objectivec.Object{ID: rv}
}
func (e EspressoANEIOSurface) CheckIfMatches(matches corevideo.CVImageBufferRef) bool {
	rv := objc.SendIfResponds[bool](e.ID, objc.Sel("checkIfMatches:"), matches)
	return rv
}
func (e EspressoANEIOSurface) CheckIfMatchesIOSurface(iOSurface iosurface.IOSurfaceRef) bool {
	rv := objc.SendIfResponds[bool](e.ID, objc.Sel("checkIfMatchesIOSurface:"), iOSurface)
	return rv
}
func (e EspressoANEIOSurface) Cleanup() {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("cleanup"))
}
func (e EspressoANEIOSurface) CreateIOSurfaceWithExtraProperties(properties objectivec.IObject) iosurface.IOSurfaceRef {
	rv := objc.SendIfResponds[iosurface.IOSurfaceRef](e.ID, objc.Sel("createIOSurfaceWithExtraProperties:"), properties)
	return iosurface.IOSurfaceRef(rv)
}
func (e EspressoANEIOSurface) DoNonLazyAllocation(allocation objectivec.IObject) {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("doNonLazyAllocation:"), allocation)
}
func (e EspressoANEIOSurface) IoSurfaceForMultiBufferFrame(frame uint64) iosurface.IOSurfaceRef {
	rv := objc.SendIfResponds[iosurface.IOSurfaceRef](e.ID, objc.Sel("ioSurfaceForMultiBufferFrame:"), frame)
	return iosurface.IOSurfaceRef(rv)
}
func (e EspressoANEIOSurface) IoSurfaceForMultiBufferFrameNoLazyForTesting(testing uint64) iosurface.IOSurfaceRef {
	rv := objc.SendIfResponds[iosurface.IOSurfaceRef](e.ID, objc.Sel("ioSurfaceForMultiBufferFrameNoLazyForTesting:"), testing)
	return iosurface.IOSurfaceRef(rv)
}
func (e EspressoANEIOSurface) LazilyAutoCreateSurfaceForFrame(frame uint64) {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("lazilyAutoCreateSurfaceForFrame:"), frame)
}
func (e EspressoANEIOSurface) MetalBufferWithDeviceMultiBufferFrame(device objectivec.IObject, frame uint64) metal.MTLBuffer {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("metalBufferWithDevice:multiBufferFrame:"), device, frame)
	return metal.MTLBufferObjectFromID(rv)
}
func (e EspressoANEIOSurface) NFrames() uint64 {
	rv := objc.SendIfResponds[uint64](e.ID, objc.Sel("nFrames"))
	return rv
}
func (e EspressoANEIOSurface) ResizeForMultipleAsyncBuffers(buffers uint64) {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("resizeForMultipleAsyncBuffers:"), buffers)
}
func (e EspressoANEIOSurface) RestoreInternalStorage(storage uint64) {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("restoreInternalStorage:"), storage)
}
func (e EspressoANEIOSurface) RestoreInternalStorageForAllMultiBufferFrames() {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("restoreInternalStorageForAllMultiBufferFrames"))
}
func (e EspressoANEIOSurface) SetExternalStorageIoSurface(storage uint64, surface iosurface.IOSurfaceRef) {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("setExternalStorage:ioSurface:"), storage, surface)
}
func (e EspressoANEIOSurface) InitWithIOSurfacePropertiesAndPixelFormats(properties objectivec.IObject, formats objectivec.IObject) EspressoANEIOSurface {
	rv := objc.SendIfResponds[EspressoANEIOSurface](e.ID, objc.Sel("initWithIOSurfaceProperties:andPixelFormats:"), properties, formats)
	return rv
}

func (e EspressoANEIOSurface) External_storage_blob_for_aliasing_mem() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("external_storage_blob_for_aliasing_mem"))
	return rv
}
func (e EspressoANEIOSurface) SetExternal_storage_blob_for_aliasing_mem(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setExternal_storage_blob_for_aliasing_mem:"), value)
}
func (e EspressoANEIOSurface) PixelFormat() uint32 {
	rv := objc.SendIfResponds[uint32](e.ID, objc.Sel("pixelFormat"))
	return rv
}

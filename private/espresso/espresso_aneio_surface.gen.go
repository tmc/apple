// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/corevideo"
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
	rv := objc.Send[EspressoANEIOSurface](objc.ID(ec.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface
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
//
// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface
type IEspressoANEIOSurface interface {
	objectivec.IObject

	// Topic: Methods

	Ane_io_surfaceForMultiBufferFrame(frame uint64) objectivec.IObject
	CheckIfMatches(matches corevideo.CVImageBufferRef) bool
	CheckIfMatchesIOSurface(iOSurface coregraphics.IOSurfaceRef) bool
	Cleanup()
	CreateIOSurfaceWithExtraProperties(properties objectivec.IObject) coregraphics.IOSurfaceRef
	DoNonLazyAllocation(allocation objectivec.IObject)
	External_storage_blob_for_aliasing_mem() objectivec.IObject
	SetExternal_storage_blob_for_aliasing_mem(value objectivec.IObject)
	IoSurfaceForMultiBufferFrame(frame uint64) coregraphics.IOSurfaceRef
	IoSurfaceForMultiBufferFrameNoLazyForTesting(testing uint64) coregraphics.IOSurfaceRef
	LazilyAutoCreateSurfaceForFrame(frame uint64)
	MetalBufferWithDeviceMultiBufferFrame(device objectivec.IObject, frame uint64) metal.MTLBuffer
	NFrames() uint64
	PixelFormat() uint32
	ResizeForMultipleAsyncBuffers(buffers uint64)
	RestoreInternalStorage(storage uint64)
	RestoreInternalStorageForAllMultiBufferFrames()
	SetExternalStorageIoSurface(storage uint64, surface coregraphics.IOSurfaceRef)
	InitWithIOSurfacePropertiesAndPixelFormats(properties objectivec.IObject, formats objectivec.IObject) EspressoANEIOSurface
}

// Init initializes the instance.
func (e EspressoANEIOSurface) Init() EspressoANEIOSurface {
	rv := objc.Send[EspressoANEIOSurface](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoANEIOSurface) Autorelease() EspressoANEIOSurface {
	rv := objc.Send[EspressoANEIOSurface](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoANEIOSurface creates a new EspressoANEIOSurface instance.
func NewEspressoANEIOSurface() EspressoANEIOSurface {
	class := getEspressoANEIOSurfaceClass()
	rv := objc.Send[EspressoANEIOSurface](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface/initWithIOSurfaceProperties:andPixelFormats:
func NewEspressoANEIOSurfaceWithIOSurfacePropertiesAndPixelFormats(properties objectivec.IObject, formats objectivec.IObject) EspressoANEIOSurface {
	instance := getEspressoANEIOSurfaceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIOSurfaceProperties:andPixelFormats:"), properties, formats)
	return EspressoANEIOSurfaceFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface/ane_io_surfaceForMultiBufferFrame:
func (e EspressoANEIOSurface) Ane_io_surfaceForMultiBufferFrame(frame uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("ane_io_surfaceForMultiBufferFrame:"), frame)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface/checkIfMatches:
func (e EspressoANEIOSurface) CheckIfMatches(matches corevideo.CVImageBufferRef) bool {
	rv := objc.Send[bool](e.ID, objc.Sel("checkIfMatches:"), matches)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface/checkIfMatchesIOSurface:
func (e EspressoANEIOSurface) CheckIfMatchesIOSurface(iOSurface coregraphics.IOSurfaceRef) bool {
	rv := objc.Send[bool](e.ID, objc.Sel("checkIfMatchesIOSurface:"), iOSurface)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface/cleanup
func (e EspressoANEIOSurface) Cleanup() {
	objc.Send[objc.ID](e.ID, objc.Sel("cleanup"))
}

// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface/createIOSurfaceWithExtraProperties:
func (e EspressoANEIOSurface) CreateIOSurfaceWithExtraProperties(properties objectivec.IObject) coregraphics.IOSurfaceRef {
	rv := objc.Send[coregraphics.IOSurfaceRef](e.ID, objc.Sel("createIOSurfaceWithExtraProperties:"), properties)
	return coregraphics.IOSurfaceRef(rv)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface/doNonLazyAllocation:
func (e EspressoANEIOSurface) DoNonLazyAllocation(allocation objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("doNonLazyAllocation:"), allocation)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface/ioSurfaceForMultiBufferFrame:
func (e EspressoANEIOSurface) IoSurfaceForMultiBufferFrame(frame uint64) coregraphics.IOSurfaceRef {
	rv := objc.Send[coregraphics.IOSurfaceRef](e.ID, objc.Sel("ioSurfaceForMultiBufferFrame:"), frame)
	return coregraphics.IOSurfaceRef(rv)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface/ioSurfaceForMultiBufferFrameNoLazyForTesting:
func (e EspressoANEIOSurface) IoSurfaceForMultiBufferFrameNoLazyForTesting(testing uint64) coregraphics.IOSurfaceRef {
	rv := objc.Send[coregraphics.IOSurfaceRef](e.ID, objc.Sel("ioSurfaceForMultiBufferFrameNoLazyForTesting:"), testing)
	return coregraphics.IOSurfaceRef(rv)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface/lazilyAutoCreateSurfaceForFrame:
func (e EspressoANEIOSurface) LazilyAutoCreateSurfaceForFrame(frame uint64) {
	objc.Send[objc.ID](e.ID, objc.Sel("lazilyAutoCreateSurfaceForFrame:"), frame)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface/metalBufferWithDevice:multiBufferFrame:
func (e EspressoANEIOSurface) MetalBufferWithDeviceMultiBufferFrame(device objectivec.IObject, frame uint64) metal.MTLBuffer {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("metalBufferWithDevice:multiBufferFrame:"), device, frame)
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface/nFrames
func (e EspressoANEIOSurface) NFrames() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("nFrames"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface/resizeForMultipleAsyncBuffers:
func (e EspressoANEIOSurface) ResizeForMultipleAsyncBuffers(buffers uint64) {
	objc.Send[objc.ID](e.ID, objc.Sel("resizeForMultipleAsyncBuffers:"), buffers)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface/restoreInternalStorage:
func (e EspressoANEIOSurface) RestoreInternalStorage(storage uint64) {
	objc.Send[objc.ID](e.ID, objc.Sel("restoreInternalStorage:"), storage)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface/restoreInternalStorageForAllMultiBufferFrames
func (e EspressoANEIOSurface) RestoreInternalStorageForAllMultiBufferFrames() {
	objc.Send[objc.ID](e.ID, objc.Sel("restoreInternalStorageForAllMultiBufferFrames"))
}

// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface/setExternalStorage:ioSurface:
func (e EspressoANEIOSurface) SetExternalStorageIoSurface(storage uint64, surface coregraphics.IOSurfaceRef) {
	objc.Send[objc.ID](e.ID, objc.Sel("setExternalStorage:ioSurface:"), storage, surface)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface/initWithIOSurfaceProperties:andPixelFormats:
func (e EspressoANEIOSurface) InitWithIOSurfacePropertiesAndPixelFormats(properties objectivec.IObject, formats objectivec.IObject) EspressoANEIOSurface {
	rv := objc.Send[EspressoANEIOSurface](e.ID, objc.Sel("initWithIOSurfaceProperties:andPixelFormats:"), properties, formats)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface/external_storage_blob_for_aliasing_mem
func (e EspressoANEIOSurface) External_storage_blob_for_aliasing_mem() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("external_storage_blob_for_aliasing_mem"))
	return objectivec.Object{ID: rv}
}
func (e EspressoANEIOSurface) SetExternal_storage_blob_for_aliasing_mem(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setExternal_storage_blob_for_aliasing_mem:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoANEIOSurface/pixelFormat
func (e EspressoANEIOSurface) PixelFormat() uint32 {
	rv := objc.Send[uint32](e.ID, objc.Sel("pixelFormat"))
	return rv
}

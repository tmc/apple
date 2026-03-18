//go:build darwin

package espresso

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	pe "github.com/tmc/apple/private/espresso"
)

// MetalDevice wraps a Metal GPU device for zero-copy interop with Espresso.
type MetalDevice struct {
	device metal.MTLDeviceObject
	closed bool
}

// OpenMetal opens the system default Metal device.
func OpenMetal() (*MetalDevice, error) {
	dev := metal.MTLCreateSystemDefaultDevice()
	if dev.GetID() == 0 {
		return nil, fmt.Errorf("espresso: no Metal device available")
	}
	d := &MetalDevice{device: dev}
	runtime.SetFinalizer(d, (*MetalDevice).Close)
	return d, nil
}

// Close releases the Metal device reference.
func (d *MetalDevice) Close() error {
	if d.closed {
		return nil
	}
	d.closed = true
	objc.Send[struct{}](d.device.GetID(), objc.Sel("release"))
	runtime.SetFinalizer(d, nil)
	return nil
}

// Device returns the underlying MTLDeviceObject.
func (d *MetalDevice) Device() metal.MTLDeviceObject { return d.device }

// IsMemoryTight returns whether Espresso's shared Metal singleton reports memory pressure.
func IsMemoryTight() bool {
	s := pe.GetEspressoMetalSingletonClass().Shared()
	if s.GetID() == 0 {
		return false
	}
	return s.Is_memory_tight() != 0
}

// ANESurface provides access to Espresso's ANE IOSurface for zero-copy
// Metal buffer creation and multi-buffer async execution.
type ANESurface struct {
	s      pe.EspressoANEIOSurface
	closed bool
}

// NewANESurface creates a new ANE IOSurface wrapper.
func NewANESurface() *ANESurface {
	a := &ANESurface{s: pe.NewEspressoANEIOSurface()}
	runtime.SetFinalizer(a, (*ANESurface).Cleanup)
	return a
}

// NewANESurfaceWithProperties creates an ANE IOSurface with specific properties and pixel formats.
func NewANESurfaceWithProperties(properties, formats objectivec.IObject) *ANESurface {
	a := &ANESurface{
		s: pe.NewEspressoANEIOSurfaceWithIOSurfacePropertiesAndPixelFormats(properties, formats),
	}
	runtime.SetFinalizer(a, (*ANESurface).Cleanup)
	return a
}

// SurfaceLayout describes the IOSurface dimensions required by a compiled model.
type SurfaceLayout struct {
	Width      int // spatial or sequence dimension
	Height     int // channels * height (flattened for IOSurface)
	BytesPerRow int // row stride (must be 64-byte aligned)
	AllocSize  int // total allocation size
	ElemSize   int // bytes per element (2 for fp16, 4 for fp32)
}

// NewANESurfaceFromLayout creates an ANE IOSurface with explicit dimensions
// from the compiled model's attributes. This replaces the flat allocator
// with a model-driven allocator that matches Width, Height, BytesPerRow,
// and AllocSize from the compiled artifact.
func NewANESurfaceFromLayout(layout SurfaceLayout) *ANESurface {
	// Build IOSurface properties dictionary.
	props := foundation.NewNSMutableDictionary()
	setIntKey := func(key string, val int) {
		props.SetObjectForKey(
			foundation.GetNSNumberClass().NumberWithInt(val),
			foundation.NSCopyingObject{Object: objectivec.Object{ID: objc.String(key)}},
		)
	}
	setIntKey("IOSurfaceWidth", layout.Width)
	setIntKey("IOSurfaceHeight", layout.Height)
	setIntKey("IOSurfaceBytesPerElement", layout.ElemSize)
	setIntKey("IOSurfaceBytesPerRow", layout.BytesPerRow)
	setIntKey("IOSurfaceAllocSize", layout.AllocSize)
	setIntKey("IOSurfacePixelFormat", 0)

	a := &ANESurface{
		s: pe.NewEspressoANEIOSurfaceWithIOSurfacePropertiesAndPixelFormats(props, nil),
	}
	runtime.SetFinalizer(a, (*ANESurface).Cleanup)
	return a
}

// MetalBuffer returns a Metal buffer backed by the ANE IOSurface for
// zero-copy GPU access. The frame parameter selects the multi-buffer
// frame index (0 for single-buffer execution).
func (a *ANESurface) MetalBuffer(d *MetalDevice, frame uint64) (metal.MTLBuffer, error) {
	buf := a.s.MetalBufferWithDeviceMultiBufferFrame(d.device, frame)
	if buf == nil || buf.GetID() == 0 {
		return nil, fmt.Errorf("espresso: failed to create Metal buffer for frame %d", frame)
	}
	return buf, nil
}

// NFrames returns the number of allocated multi-buffer frames.
func (a *ANESurface) NFrames() uint64 {
	return a.s.NFrames()
}

// ResizeForAsync resizes the surface for multiple async buffer frames.
func (a *ANESurface) ResizeForAsync(n uint64) {
	a.s.ResizeForMultipleAsyncBuffers(n)
}

// PixelFormat returns the surface pixel format.
func (a *ANESurface) PixelFormat() uint32 {
	return a.s.PixelFormat()
}

// IOSurfaceForFrame returns the IOSurface backing the given multi-buffer frame.
func (a *ANESurface) IOSurfaceForFrame(frame uint64) (coregraphics.IOSurfaceRef, error) {
	ref := a.s.IoSurfaceForMultiBufferFrame(frame)
	if ref == 0 {
		return 0, fmt.Errorf("espresso: no IOSurface for frame %d", frame)
	}
	return ref, nil
}

// SetExternalStorage replaces the IOSurface backing a storage slot with an external one.
func (a *ANESurface) SetExternalStorage(storage uint64, surface coregraphics.IOSurfaceRef) {
	a.s.SetExternalStorageIoSurface(storage, surface)
}

// RestoreInternalStorage restores internal storage for a single slot.
func (a *ANESurface) RestoreInternalStorage(storage uint64) {
	a.s.RestoreInternalStorage(storage)
}

// RestoreAllInternalStorage restores internal storage for all multi-buffer frames.
func (a *ANESurface) RestoreAllInternalStorage() {
	a.s.RestoreInternalStorageForAllMultiBufferFrames()
}

// CreateIOSurface creates a new IOSurface with optional extra properties.
func (a *ANESurface) CreateIOSurface(extraProperties objectivec.IObject) coregraphics.IOSurfaceRef {
	return a.s.CreateIOSurfaceWithExtraProperties(extraProperties)
}

// LazilyAllocFrame lazily allocates an IOSurface for the given frame.
func (a *ANESurface) LazilyAllocFrame(frame uint64) {
	a.s.LazilyAutoCreateSurfaceForFrame(frame)
}

// ForceAlloc forces non-lazy allocation with the given allocation parameters.
func (a *ANESurface) ForceAlloc(allocation objectivec.IObject) {
	a.s.DoNonLazyAllocation(allocation)
}

// MatchesCVImageBuffer returns whether the surface matches a CVImageBuffer.
func (a *ANESurface) MatchesCVImageBuffer(buf corevideo.CVImageBufferRef) bool {
	return a.s.CheckIfMatches(buf)
}

// MatchesIOSurface returns whether the surface matches an IOSurfaceRef.
func (a *ANESurface) MatchesIOSurface(ref coregraphics.IOSurfaceRef) bool {
	return a.s.CheckIfMatchesIOSurface(ref)
}

// AliasingMem returns the external storage blob used for aliasing memory.
func (a *ANESurface) AliasingMem() objectivec.IObject {
	return a.s.External_storage_blob_for_aliasing_mem()
}

// SetAliasingMem sets the external storage blob for aliasing memory.
func (a *ANESurface) SetAliasingMem(blob objectivec.IObject) {
	a.s.SetExternal_storage_blob_for_aliasing_mem(blob)
}

// WriteFrame writes raw bytes to the IOSurface backing the given frame.
func (a *ANESurface) WriteFrame(frame uint64, data []byte) error {
	ref := a.s.IoSurfaceForMultiBufferFrame(frame)
	if ref == 0 {
		return fmt.Errorf("espresso: no IOSurface for frame %d", frame)
	}
	surfRef := iosurface.IOSurfaceRef(ref)
	iosurface.IOSurfaceLock(surfRef, 0, nil)
	base := iosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		iosurface.IOSurfaceUnlock(surfRef, 0, nil)
		return fmt.Errorf("espresso: IOSurface base address is nil for frame %d", frame)
	}
	dst := unsafe.Slice((*byte)(base), len(data))
	copy(dst, data)
	iosurface.IOSurfaceUnlock(surfRef, 0, nil)
	return nil
}

// ReadFrame reads raw bytes from the IOSurface backing the given frame.
func (a *ANESurface) ReadFrame(frame uint64, data []byte) error {
	ref := a.s.IoSurfaceForMultiBufferFrame(frame)
	if ref == 0 {
		return fmt.Errorf("espresso: no IOSurface for frame %d", frame)
	}
	surfRef := iosurface.IOSurfaceRef(ref)
	iosurface.IOSurfaceLock(surfRef, iosurface.KIOSurfaceLockReadOnly, nil)
	base := iosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		iosurface.IOSurfaceUnlock(surfRef, iosurface.KIOSurfaceLockReadOnly, nil)
		return fmt.Errorf("espresso: IOSurface base address is nil for frame %d", frame)
	}
	src := unsafe.Slice((*byte)(base), len(data))
	copy(data, src)
	iosurface.IOSurfaceUnlock(surfRef, iosurface.KIOSurfaceLockReadOnly, nil)
	return nil
}

// WriteFrameF32 writes float32 data to the IOSurface backing the given frame.
func (a *ANESurface) WriteFrameF32(frame uint64, data []float32) error {
	return a.WriteFrame(frame, unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(data))), len(data)*4))
}

// ReadFrameF32 reads float32 data from the IOSurface backing the given frame.
func (a *ANESurface) ReadFrameF32(frame uint64, data []float32) error {
	return a.ReadFrame(frame, unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(data))), len(data)*4))
}

// Cleanup releases surface resources.
func (a *ANESurface) Cleanup() {
	if a.closed {
		return
	}
	a.closed = true
	a.s.Cleanup()
	runtime.SetFinalizer(a, nil)
}

//go:build darwin

package ane

import (
	"fmt"
	"math"
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objectivec"
)

// SurfaceBinding binds an IOSurface to a compiled model symbol index.
type SurfaceBinding struct {
	Surface     coregraphics.IOSurfaceRef
	SymbolIndex int
}

// IOSurfaceFloat32 wraps an IOSurface that stores float32 data.
type IOSurfaceFloat32 struct {
	mu          sync.Mutex
	surface     coregraphics.IOSurfaceRef
	count       int
	layout      *TensorLayout
	owned       bool
	readLocked  bool
	writeLocked bool
}

// NewIOSurfaceFloat32 allocates a float32 IOSurface with count elements.
func NewIOSurfaceFloat32(count int) (*IOSurfaceFloat32, error) {
	surf, err := newFloatSurface(count)
	if err != nil {
		return nil, err
	}
	return &IOSurfaceFloat32{surface: surf, count: count, owned: true}, nil
}

// NewIOSurfaceFloat32WithLayout allocates a float32 IOSurface with the given layout.
func NewIOSurfaceFloat32WithLayout(layout TensorLayout) (*IOSurfaceFloat32, error) {
	if layout.ElemSize != 4 {
		return nil, fmt.Errorf("ane: float32 IOSurface requires ElemSize=4, got %d", layout.ElemSize)
	}
	if err := validateLayout(layout); err != nil {
		return nil, err
	}
	surf, err := createSurfaceForLayout(layout)
	if err != nil {
		return nil, err
	}
	layoutCopy := layout
	return &IOSurfaceFloat32{
		surface: surf,
		count:   layout.LogicalElements(),
		layout:  &layoutCopy,
		owned:   true,
	}, nil
}

// WrapIOSurfaceFloat32 wraps an existing IOSurface with a logical element count.
func WrapIOSurfaceFloat32(surface coregraphics.IOSurfaceRef, count int) (*IOSurfaceFloat32, error) {
	if surface == 0 {
		return nil, fmt.Errorf("ane: surface is nil")
	}
	if count <= 0 {
		return nil, fmt.Errorf("ane: invalid float32 IOSurface count %d", count)
	}
	return &IOSurfaceFloat32{surface: surface, count: count}, nil
}

// WrapIOSurfaceFloat32WithLayout wraps an existing IOSurface with a compiled layout.
func WrapIOSurfaceFloat32WithLayout(surface coregraphics.IOSurfaceRef, layout TensorLayout) (*IOSurfaceFloat32, error) {
	if surface == 0 {
		return nil, fmt.Errorf("ane: surface is nil")
	}
	if layout.ElemSize != 4 {
		return nil, fmt.Errorf("ane: float32 IOSurface requires ElemSize=4, got %d", layout.ElemSize)
	}
	if err := validateLayout(layout); err != nil {
		return nil, err
	}
	layoutCopy := layout
	return &IOSurfaceFloat32{
		surface: surface,
		count:   layout.LogicalElements(),
		layout:  &layoutCopy,
	}, nil
}

// Ref returns the wrapped IOSurface.
func (s *IOSurfaceFloat32) Ref() coregraphics.IOSurfaceRef {
	if s == nil {
		return 0
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.surface
}

// Count returns the logical float32 element count.
func (s *IOSurfaceFloat32) Count() int {
	if s == nil {
		return 0
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.count
}

// ByteLen returns the IOSurface allocation size in bytes.
func (s *IOSurfaceFloat32) ByteLen() int {
	if s == nil {
		return 0
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.surface == 0 {
		if s.layout != nil {
			return s.layout.AllocSize()
		}
		return s.count * 4
	}
	return surfaceSize(s.surface)
}

// Layout reports the compiled layout attached to this wrapper, if any.
func (s *IOSurfaceFloat32) Layout() (TensorLayout, bool) {
	if s == nil {
		return TensorLayout{}, false
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.layout == nil {
		return TensorLayout{}, false
	}
	return *s.layout, true
}

// Binding returns a request binding for the given compiled symbol index.
func (s *IOSurfaceFloat32) Binding(symbolIndex int) SurfaceBinding {
	return SurfaceBinding{Surface: s.Ref(), SymbolIndex: symbolIndex}
}

// LayoutBinding returns a request binding using the attached layout's symbol index.
func (s *IOSurfaceFloat32) LayoutBinding() (SurfaceBinding, error) {
	layout, ok := s.Layout()
	if !ok {
		return SurfaceBinding{}, fmt.Errorf("ane: surface has no attached layout")
	}
	if layout.SymbolIndex < 0 {
		return SurfaceBinding{}, fmt.Errorf("ane: surface layout has no symbol index")
	}
	return SurfaceBinding{Surface: s.Ref(), SymbolIndex: layout.SymbolIndex}, nil
}

// Write copies data into the IOSurface.
func (s *IOSurfaceFloat32) Write(data []float32) error {
	if s == nil {
		return fmt.Errorf("ane: float32 IOSurface is nil")
	}
	s.mu.Lock()
	surf := s.surface
	count := s.count
	layout := s.layout
	s.mu.Unlock()
	if surf == 0 {
		return fmt.Errorf("ane: float32 IOSurface is closed")
	}
	if len(data) != count {
		return fmt.Errorf("ane: float32 IOSurface write len=%d want=%d", len(data), count)
	}
	if layout != nil {
		return writeFloat32IOSurfaceWithLayout(surf, data, *layout)
	}
	return writeFloat32IOSurface(surf, data)
}

// WriteAt copies data into the IOSurface starting at the logical element offset.
func (s *IOSurfaceFloat32) WriteAt(offset int, data []float32) error {
	if s == nil {
		return fmt.Errorf("ane: float32 IOSurface is nil")
	}
	if offset < 0 {
		return fmt.Errorf("ane: invalid write offset %d", offset)
	}
	if len(data) == 0 {
		return nil
	}
	s.mu.Lock()
	surf := s.surface
	count := s.count
	layout := s.layout
	s.mu.Unlock()
	if surf == 0 {
		return fmt.Errorf("ane: float32 IOSurface is closed")
	}
	if offset+len(data) > count {
		return fmt.Errorf("ane: write range [%d,%d) out of bounds [0,%d)", offset, offset+len(data), count)
	}
	if layout != nil {
		return writeFloat32IOSurfaceAtWithLayout(surf, offset, data, count, *layout)
	}
	raw := iosurface.IOSurfaceRef(surf)
	if rc := iosurface.IOSurfaceLock(raw, 0, nil); rc != 0 {
		return fmt.Errorf("ane: IOSurface lock failed rc=%d", rc)
	}
	defer iosurface.IOSurfaceUnlock(raw, 0, nil)
	base := iosurface.IOSurfaceGetBaseAddress(raw)
	if base == nil {
		return fmt.Errorf("ane: IOSurface base address is nil")
	}
	dst := unsafe.Slice((*float32)(base), count)
	copy(dst[offset:offset+len(data)], data)
	return nil
}

// Read reads the full logical tensor contents from the IOSurface.
func (s *IOSurfaceFloat32) Read() ([]float32, error) {
	if s == nil {
		return nil, fmt.Errorf("ane: float32 IOSurface is nil")
	}
	s.mu.Lock()
	surf := s.surface
	count := s.count
	layout := s.layout
	s.mu.Unlock()
	if surf == 0 {
		return nil, fmt.Errorf("ane: float32 IOSurface is closed")
	}
	if layout != nil {
		return readFloat32IOSurfaceWithLayout(surf, count, *layout)
	}
	return readFloat32IOSurface(surf, count)
}

// ReadAt reads n logical elements from the IOSurface starting at offset.
func (s *IOSurfaceFloat32) ReadAt(offset, n int) ([]float32, error) {
	if s == nil {
		return nil, fmt.Errorf("ane: float32 IOSurface is nil")
	}
	if offset < 0 || n < 0 {
		return nil, fmt.Errorf("ane: invalid read range offset=%d n=%d", offset, n)
	}
	s.mu.Lock()
	surf := s.surface
	count := s.count
	layout := s.layout
	s.mu.Unlock()
	if surf == 0 {
		return nil, fmt.Errorf("ane: float32 IOSurface is closed")
	}
	if offset+n > count {
		return nil, fmt.Errorf("ane: read range [%d,%d) out of bounds [0,%d)", offset, offset+n, count)
	}
	if layout != nil {
		return readFloat32IOSurfaceAtWithLayout(surf, offset, n, count, *layout)
	}
	out := make([]float32, n)
	if n == 0 {
		return out, nil
	}
	raw := iosurface.IOSurfaceRef(surf)
	if rc := iosurface.IOSurfaceLock(raw, iosurface.KIOSurfaceLockReadOnly, nil); rc != 0 {
		return nil, fmt.Errorf("ane: IOSurface read lock failed rc=%d", rc)
	}
	defer iosurface.IOSurfaceUnlock(raw, iosurface.KIOSurfaceLockReadOnly, nil)
	base := iosurface.IOSurfaceGetBaseAddress(raw)
	if base == nil {
		return nil, fmt.Errorf("ane: IOSurface base address is nil")
	}
	src := unsafe.Slice((*float32)(base), count)
	copy(out, src[offset:offset+n])
	return out, nil
}

// Fill overwrites the IOSurface with one value.
func (s *IOSurfaceFloat32) Fill(v float32) error {
	if s == nil {
		return fmt.Errorf("ane: float32 IOSurface is nil")
	}
	if v == 0 {
		s.mu.Lock()
		surf := s.surface
		layout := s.layout
		count := s.count
		s.mu.Unlock()
		if surf == 0 {
			return fmt.Errorf("ane: float32 IOSurface is closed")
		}
		if layout != nil {
			return zeroIOSurfaceWithLayout(surf, *layout)
		}
		return writeFloat32IOSurface(surf, make([]float32, count))
	}
	count := s.Count()
	data := make([]float32, count)
	for i := range data {
		data[i] = v
	}
	return s.Write(data)
}

// FillNaN overwrites the IOSurface with NaNs.
func (s *IOSurfaceFloat32) FillNaN() error {
	return s.Fill(float32(math.NaN()))
}

// LockReadOnly locks the IOSurface for read-only zero-copy access.
func (s *IOSurfaceFloat32) LockReadOnly() (unsafe.Pointer, int, error) {
	if s == nil {
		return nil, 0, fmt.Errorf("ane: float32 IOSurface is nil")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.surface == 0 {
		return nil, 0, fmt.Errorf("ane: float32 IOSurface is closed")
	}
	if s.readLocked || s.writeLocked {
		return nil, 0, fmt.Errorf("ane: float32 IOSurface is already locked")
	}
	raw := iosurface.IOSurfaceRef(s.surface)
	if rc := iosurface.IOSurfaceLock(raw, iosurface.KIOSurfaceLockReadOnly, nil); rc != 0 {
		return nil, 0, fmt.Errorf("ane: IOSurface read lock failed rc=%d", rc)
	}
	base := iosurface.IOSurfaceGetBaseAddress(raw)
	if base == nil {
		iosurface.IOSurfaceUnlock(raw, iosurface.KIOSurfaceLockReadOnly, nil)
		return nil, 0, fmt.Errorf("ane: IOSurface base address is nil")
	}
	s.readLocked = true
	return base, int(iosurface.IOSurfaceGetAllocSize(raw)), nil
}

// UnlockReadOnly unlocks a previous LockReadOnly call.
func (s *IOSurfaceFloat32) UnlockReadOnly() error {
	if s == nil {
		return nil
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.surface == 0 || !s.readLocked {
		return nil
	}
	raw := iosurface.IOSurfaceRef(s.surface)
	if rc := iosurface.IOSurfaceUnlock(raw, iosurface.KIOSurfaceLockReadOnly, nil); rc != 0 {
		return fmt.Errorf("ane: IOSurface read unlock failed rc=%d", rc)
	}
	s.readLocked = false
	return nil
}

// LockWritable locks the IOSurface for read-write zero-copy access.
func (s *IOSurfaceFloat32) LockWritable() (unsafe.Pointer, int, error) {
	if s == nil {
		return nil, 0, fmt.Errorf("ane: float32 IOSurface is nil")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.surface == 0 {
		return nil, 0, fmt.Errorf("ane: float32 IOSurface is closed")
	}
	if s.readLocked || s.writeLocked {
		return nil, 0, fmt.Errorf("ane: float32 IOSurface is already locked")
	}
	raw := iosurface.IOSurfaceRef(s.surface)
	if rc := iosurface.IOSurfaceLock(raw, 0, nil); rc != 0 {
		return nil, 0, fmt.Errorf("ane: IOSurface write lock failed rc=%d", rc)
	}
	base := iosurface.IOSurfaceGetBaseAddress(raw)
	if base == nil {
		iosurface.IOSurfaceUnlock(raw, 0, nil)
		return nil, 0, fmt.Errorf("ane: IOSurface base address is nil")
	}
	s.writeLocked = true
	return base, int(iosurface.IOSurfaceGetAllocSize(raw)), nil
}

// UnlockWritable unlocks a previous LockWritable call.
func (s *IOSurfaceFloat32) UnlockWritable() error {
	if s == nil {
		return nil
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.surface == 0 || !s.writeLocked {
		return nil
	}
	raw := iosurface.IOSurfaceRef(s.surface)
	if rc := iosurface.IOSurfaceUnlock(raw, 0, nil); rc != 0 {
		return fmt.Errorf("ane: IOSurface write unlock failed rc=%d", rc)
	}
	s.writeLocked = false
	return nil
}

// Close releases the wrapped IOSurface if this wrapper owns it.
func (s *IOSurfaceFloat32) Close() error {
	if s == nil {
		return nil
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.surface == 0 {
		return nil
	}
	raw := iosurface.IOSurfaceRef(s.surface)
	if s.readLocked {
		_ = iosurface.IOSurfaceUnlock(raw, iosurface.KIOSurfaceLockReadOnly, nil)
		s.readLocked = false
	}
	if s.writeLocked {
		_ = iosurface.IOSurfaceUnlock(raw, 0, nil)
		s.writeLocked = false
	}
	if s.owned {
		corefoundation.CFRelease(corefoundation.CFTypeRef(s.surface))
	}
	s.surface = 0
	return nil
}

// MetalBufferBinding is a reusable no-copy Metal view over an IOSurface.
type MetalBufferBinding struct {
	buffer  metal.MTLBufferObject
	surface *IOSurfaceFloat32
	ptr     unsafe.Pointer
	byteLen int
}

// Buffer returns the underlying Metal buffer.
func (b *MetalBufferBinding) Buffer() metal.MTLBufferObject {
	if b == nil {
		return metal.MTLBufferObject{}
	}
	return b.buffer
}

// Pointer returns the IOSurface base address used for the Metal buffer.
func (b *MetalBufferBinding) Pointer() unsafe.Pointer {
	if b == nil {
		return nil
	}
	return b.ptr
}

// ByteLen returns the byte length of the Metal mapping.
func (b *MetalBufferBinding) ByteLen() int {
	if b == nil {
		return 0
	}
	return b.byteLen
}

// LockReadOnly locks the IOSurface before Metal reads.
func (b *MetalBufferBinding) LockReadOnly() error {
	if b == nil || b.surface == nil {
		return fmt.Errorf("ane: metal buffer binding is nil")
	}
	ptr, n, err := b.surface.LockReadOnly()
	if err != nil {
		return err
	}
	if ptr != b.ptr || n != b.byteLen {
		_ = b.surface.UnlockReadOnly()
		return fmt.Errorf("ane: IOSurface mapping changed")
	}
	return nil
}

// UnlockReadOnly unlocks a previous LockReadOnly call.
func (b *MetalBufferBinding) UnlockReadOnly() error {
	if b == nil || b.surface == nil {
		return nil
	}
	return b.surface.UnlockReadOnly()
}

// Close releases the Metal buffer and unlocks the IOSurface if needed.
func (b *MetalBufferBinding) Close() error {
	if b == nil {
		return nil
	}
	if b.buffer.GetID() != 0 {
		objectivec.ObjectFromID(b.buffer.GetID()).Release()
		b.buffer = metal.MTLBufferObject{}
	}
	err := b.UnlockReadOnly()
	b.surface = nil
	b.ptr = nil
	b.byteLen = 0
	return err
}

// NewMetalBufferBinding creates a reusable no-copy Metal view over the IOSurface.
func (s *IOSurfaceFloat32) NewMetalBufferBinding(device *MetalDevice) (*MetalBufferBinding, error) {
	if s == nil {
		return nil, fmt.Errorf("ane: float32 IOSurface is nil")
	}
	if device == nil || device.device.GetID() == 0 {
		return nil, fmt.Errorf("ane: metal device is nil")
	}
	ptr, n, err := s.LockReadOnly()
	if err != nil {
		return nil, err
	}
	if err := s.UnlockReadOnly(); err != nil {
		return nil, err
	}
	buf := device.device.NewBufferWithBytesNoCopyLengthOptionsDeallocator(
		ptr,
		uint(n),
		metal.MTLResourceStorageModeShared,
		nil,
	)
	if buf == nil || buf.GetID() == 0 {
		return nil, fmt.Errorf("ane: failed to create Metal buffer")
	}
	return &MetalBufferBinding{
		buffer:  metal.MTLBufferObjectFromID(buf.GetID()),
		surface: s,
		ptr:     ptr,
		byteLen: n,
	}, nil
}

func newFloatSurface(count int) (coregraphics.IOSurfaceRef, error) {
	if count <= 0 {
		return 0, fmt.Errorf("ane: invalid float32 IOSurface count %d", count)
	}
	layout := TensorLayout{
		Channels:    1,
		Width:       count,
		Height:      1,
		ElemSize:    4,
		RowStride:   RowStrideFor(count, 4),
		PlaneStride: RowStrideFor(count, 4),
	}
	return createSurfaceForLayout(layout)
}

func writeFloat32IOSurface(surface coregraphics.IOSurfaceRef, data []float32) error {
	if len(data) == 0 {
		return nil
	}
	raw := iosurface.IOSurfaceRef(surface)
	if rc := iosurface.IOSurfaceLock(raw, 0, nil); rc != 0 {
		return fmt.Errorf("ane: IOSurface lock failed rc=%d", rc)
	}
	defer iosurface.IOSurfaceUnlock(raw, 0, nil)
	base := iosurface.IOSurfaceGetBaseAddress(raw)
	if base == nil {
		return fmt.Errorf("ane: IOSurface base address is nil")
	}
	n := len(data) * 4
	dst := unsafe.Slice((*byte)(base), n)
	src := unsafe.Slice((*byte)(unsafe.Pointer(&data[0])), n)
	copy(dst, src)
	return nil
}

func readFloat32IOSurface(surface coregraphics.IOSurfaceRef, count int) ([]float32, error) {
	if count < 0 {
		return nil, fmt.Errorf("ane: invalid float32 IOSurface count %d", count)
	}
	out := make([]float32, count)
	if count == 0 {
		return out, nil
	}
	raw := iosurface.IOSurfaceRef(surface)
	if rc := iosurface.IOSurfaceLock(raw, iosurface.KIOSurfaceLockReadOnly, nil); rc != 0 {
		return nil, fmt.Errorf("ane: IOSurface read lock failed rc=%d", rc)
	}
	defer iosurface.IOSurfaceUnlock(raw, iosurface.KIOSurfaceLockReadOnly, nil)
	base := iosurface.IOSurfaceGetBaseAddress(raw)
	if base == nil {
		return nil, fmt.Errorf("ane: IOSurface base address is nil")
	}
	n := count * 4
	src := unsafe.Slice((*byte)(base), n)
	dst := unsafe.Slice((*byte)(unsafe.Pointer(&out[0])), n)
	copy(dst, src)
	return out, nil
}

func writeFloat32IOSurfaceWithLayout(surface coregraphics.IOSurfaceRef, data []float32, layout TensorLayout) error {
	if err := validateLayout(layout); err != nil {
		return err
	}
	if layout.ElemSize != 4 {
		return fmt.Errorf("ane: float32 IOSurface layout requires ElemSize=4, got %d", layout.ElemSize)
	}
	if len(data) != layout.LogicalElements() {
		return fmt.Errorf("ane: float32 IOSurface write len=%d want=%d", len(data), layout.LogicalElements())
	}
	return writeStridedF32WithLayout(surface, data, layout)
}

func readFloat32IOSurfaceWithLayout(surface coregraphics.IOSurfaceRef, count int, layout TensorLayout) ([]float32, error) {
	if err := validateLayout(layout); err != nil {
		return nil, err
	}
	if layout.ElemSize != 4 {
		return nil, fmt.Errorf("ane: float32 IOSurface layout requires ElemSize=4, got %d", layout.ElemSize)
	}
	if count != layout.LogicalElements() {
		return nil, fmt.Errorf("ane: float32 IOSurface read count=%d want=%d", count, layout.LogicalElements())
	}
	out := make([]float32, count)
	if err := readStridedF32WithLayout(surface, out, layout); err != nil {
		return nil, err
	}
	return out, nil
}

func writeFloat32IOSurfaceAtWithLayout(surface coregraphics.IOSurfaceRef, offset int, data []float32, count int, layout TensorLayout) error {
	if err := validateLayout(layout); err != nil {
		return err
	}
	if layout.ElemSize != 4 {
		return fmt.Errorf("ane: float32 IOSurface layout requires ElemSize=4, got %d", layout.ElemSize)
	}
	raw := iosurface.IOSurfaceRef(surface)
	if rc := iosurface.IOSurfaceLock(raw, 0, nil); rc != 0 {
		return fmt.Errorf("ane: IOSurface lock failed rc=%d", rc)
	}
	defer iosurface.IOSurfaceUnlock(raw, 0, nil)
	base := iosurface.IOSurfaceGetBaseAddress(raw)
	if base == nil {
		return fmt.Errorf("ane: IOSurface base address is nil")
	}
	dst := unsafe.Slice((*byte)(base), layout.AllocSize())
	for i, v := range data {
		off, err := layoutByteOffset(layout, offset+i, count)
		if err != nil {
			return err
		}
		*(*float32)(unsafe.Pointer(&dst[off])) = v
	}
	return nil
}

func readFloat32IOSurfaceAtWithLayout(surface coregraphics.IOSurfaceRef, offset, n, count int, layout TensorLayout) ([]float32, error) {
	if err := validateLayout(layout); err != nil {
		return nil, err
	}
	if layout.ElemSize != 4 {
		return nil, fmt.Errorf("ane: float32 IOSurface layout requires ElemSize=4, got %d", layout.ElemSize)
	}
	raw := iosurface.IOSurfaceRef(surface)
	if rc := iosurface.IOSurfaceLock(raw, iosurface.KIOSurfaceLockReadOnly, nil); rc != 0 {
		return nil, fmt.Errorf("ane: IOSurface read lock failed rc=%d", rc)
	}
	defer iosurface.IOSurfaceUnlock(raw, iosurface.KIOSurfaceLockReadOnly, nil)
	base := iosurface.IOSurfaceGetBaseAddress(raw)
	if base == nil {
		return nil, fmt.Errorf("ane: IOSurface base address is nil")
	}
	src := unsafe.Slice((*byte)(base), layout.AllocSize())
	out := make([]float32, n)
	for i := range out {
		off, err := layoutByteOffset(layout, offset+i, count)
		if err != nil {
			return nil, err
		}
		out[i] = *(*float32)(unsafe.Pointer(&src[off]))
	}
	return out, nil
}

func zeroIOSurfaceWithLayout(surface coregraphics.IOSurfaceRef, layout TensorLayout) error {
	raw := iosurface.IOSurfaceRef(surface)
	if rc := iosurface.IOSurfaceLock(raw, 0, nil); rc != 0 {
		return fmt.Errorf("ane: IOSurface lock failed rc=%d", rc)
	}
	defer iosurface.IOSurfaceUnlock(raw, 0, nil)
	base := iosurface.IOSurfaceGetBaseAddress(raw)
	if base == nil {
		return fmt.Errorf("ane: IOSurface base address is nil")
	}
	dst := unsafe.Slice((*byte)(base), layout.AllocSize())
	clear(dst)
	return nil
}

func layoutByteOffset(layout TensorLayout, logicalIndex, count int) (int, error) {
	if logicalIndex < 0 || logicalIndex >= count {
		return 0, fmt.Errorf("ane: logical index %d out of bounds [0,%d)", logicalIndex, count)
	}
	planeElems := layout.Height * layout.Width
	if planeElems <= 0 {
		return 0, fmt.Errorf("ane: invalid layout plane size %d", planeElems)
	}
	channel := logicalIndex / planeElems
	rem := logicalIndex % planeElems
	row := rem / layout.Width
	col := rem % layout.Width
	return channel*layout.PlaneStride + row*layout.RowStride + col*layout.ElemSize, nil
}

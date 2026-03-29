//go:build darwin

package ane

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
)

// IOSurfaceTensor is an IOSurface-backed tensor that can be used as a
// zero-copy CoreML MLMultiArray input. CoreML reads directly from the
// IOSurface, avoiding a buffer copy into the compute unit.
//
// Only fp16 tensors are supported (kCVPixelFormatType_OneComponent16Half),
// matching the CoreML IOSurface requirement.
//
// The tensor owns its CVPixelBuffer and underlying IOSurface. Call Close
// to release them.
type IOSurfaceTensor struct {
	pixelBuf corevideo.CVPixelBufferRef
	width    int // last dimension of shape
	height   int // product of all other dimensions
	shape    []int
}

// NewIOSurfaceTensor creates an IOSurface-backed fp16 tensor with the given
// shape. The last dimension becomes the pixel buffer width; the product of
// all preceding dimensions becomes the height.
//
// The created pixel buffer uses kCVPixelFormatType_OneComponent16Half and is
// IOSurface-backed, making it suitable for passing to
// MLMultiArray.initWithPixelBuffer:shape:.
func NewIOSurfaceTensor(shape []int) (*IOSurfaceTensor, error) {
	if len(shape) == 0 {
		return nil, fmt.Errorf("ane: empty shape")
	}

	width := shape[len(shape)-1]
	height := 1
	for _, d := range shape[:len(shape)-1] {
		height *= d
	}
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("ane: invalid shape %v (width=%d, height=%d)", shape, width, height)
	}

	// Create IOSurface-backed CVPixelBuffer.
	attrs := iosurfacePixelBufferAttrs()
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(attrs))

	var pixelBuf corevideo.CVPixelBufferRef
	status := corevideo.CVPixelBufferCreate(
		0, // default allocator
		uintptr(width),
		uintptr(height),
		corevideo.CVPixelFormatOneComponent16Half,
		attrs,
		&pixelBuf,
	)
	if status != 0 {
		return nil, fmt.Errorf("ane: CVPixelBufferCreate failed (status %d) for shape %v", status, shape)
	}
	if pixelBuf == 0 {
		return nil, fmt.Errorf("ane: CVPixelBufferCreate returned nil for shape %v", shape)
	}

	// Verify IOSurface backing.
	ioSurf := corevideo.CVPixelBufferGetIOSurface(pixelBuf)
	if ioSurf == nil {
		corefoundation.CFRelease(corefoundation.CFTypeRef(pixelBuf))
		return nil, fmt.Errorf("ane: pixel buffer is not IOSurface-backed for shape %v", shape)
	}

	shapeCopy := make([]int, len(shape))
	copy(shapeCopy, shape)

	return &IOSurfaceTensor{
		pixelBuf: pixelBuf,
		width:    width,
		height:   height,
		shape:    shapeCopy,
	}, nil
}

// Close releases the CVPixelBuffer and its backing IOSurface.
func (t *IOSurfaceTensor) Close() {
	if t.pixelBuf != 0 {
		corefoundation.CFRelease(corefoundation.CFTypeRef(t.pixelBuf))
		t.pixelBuf = 0
	}
}

// Shape returns the tensor's logical shape.
func (t *IOSurfaceTensor) Shape() []int { return t.shape }

// PixelBuffer returns the underlying CVPixelBufferRef.
func (t *IOSurfaceTensor) PixelBuffer() corevideo.CVPixelBufferRef { return t.pixelBuf }

// MLMultiArray creates an MLMultiArray backed by this tensor's IOSurface.
// The MLMultiArray borrows the pixel buffer — the IOSurfaceTensor must
// outlive the MLMultiArray.
//
// The returned array has dtype float16 and the tensor's shape.
func (t *IOSurfaceTensor) MLMultiArray() (coreml.MLMultiArray, error) {
	if t.pixelBuf == 0 {
		return coreml.MLMultiArray{}, fmt.Errorf("ane: IOSurfaceTensor is closed")
	}

	nsShape := make([]foundation.NSNumber, len(t.shape))
	for i, d := range t.shape {
		nsShape[i] = foundation.GetNSNumberClass().NumberWithInteger(d)
	}

	arr := coreml.NewMultiArrayWithPixelBufferShape(
		corevideo.CVImageBufferRef(t.pixelBuf),
		nsShape,
	)
	if arr.ID == 0 {
		return coreml.MLMultiArray{}, fmt.Errorf("ane: initWithPixelBuffer:shape: returned nil for shape %v", t.shape)
	}
	return arr, nil
}

// WriteFloat32 writes float32 data into the tensor, converting to fp16.
// The data must have exactly NumElements() values in row-major order.
func (t *IOSurfaceTensor) WriteFloat32(data []float32) error {
	n := t.NumElements()
	if len(data) != n {
		return fmt.Errorf("ane: WriteFloat32 got %d values, want %d", len(data), n)
	}
	if t.pixelBuf == 0 {
		return fmt.Errorf("ane: IOSurfaceTensor is closed")
	}

	status := corevideo.CVPixelBufferLockBaseAddress(t.pixelBuf, 0)
	if status != 0 {
		return fmt.Errorf("ane: CVPixelBufferLockBaseAddress failed (status %d)", status)
	}

	base := corevideo.CVPixelBufferGetBaseAddress(t.pixelBuf)
	if base == nil {
		corevideo.CVPixelBufferUnlockBaseAddress(t.pixelBuf, 0)
		return fmt.Errorf("ane: pixel buffer base address is nil")
	}

	dst := unsafe.Slice((*uint16)(base), n)
	for i, v := range data {
		dst[i] = float32ToFP16(v)
	}

	corevideo.CVPixelBufferUnlockBaseAddress(t.pixelBuf, 0)
	runtime.KeepAlive(t)
	return nil
}

// ReadFloat32 reads the tensor's fp16 data into float32 values.
// dst must have exactly NumElements() capacity.
func (t *IOSurfaceTensor) ReadFloat32(dst []float32) error {
	n := t.NumElements()
	if len(dst) != n {
		return fmt.Errorf("ane: ReadFloat32 got %d slots, want %d", len(dst), n)
	}
	if t.pixelBuf == 0 {
		return fmt.Errorf("ane: IOSurfaceTensor is closed")
	}

	status := corevideo.CVPixelBufferLockBaseAddress(t.pixelBuf, corevideo.KCVPixelBufferLock_ReadOnly)
	if status != 0 {
		return fmt.Errorf("ane: CVPixelBufferLockBaseAddress failed (status %d)", status)
	}

	base := corevideo.CVPixelBufferGetBaseAddress(t.pixelBuf)
	if base == nil {
		corevideo.CVPixelBufferUnlockBaseAddress(t.pixelBuf, corevideo.KCVPixelBufferLock_ReadOnly)
		return fmt.Errorf("ane: pixel buffer base address is nil")
	}

	src := unsafe.Slice((*uint16)(base), n)
	for i, bits := range src {
		dst[i] = fp16ToFloat32(bits)
	}

	corevideo.CVPixelBufferUnlockBaseAddress(t.pixelBuf, corevideo.KCVPixelBufferLock_ReadOnly)
	runtime.KeepAlive(t)
	return nil
}

// DataPointer returns the raw fp16 data pointer after locking the base
// address. The caller must call the returned unlock function after use.
//
// This enables zero-copy reads by Metal or MLX without float conversion.
func (t *IOSurfaceTensor) DataPointer() (ptr unsafe.Pointer, unlock func(), err error) {
	if t.pixelBuf == 0 {
		return nil, nil, fmt.Errorf("ane: IOSurfaceTensor is closed")
	}

	status := corevideo.CVPixelBufferLockBaseAddress(t.pixelBuf, corevideo.KCVPixelBufferLock_ReadOnly)
	if status != 0 {
		return nil, nil, fmt.Errorf("ane: CVPixelBufferLockBaseAddress failed (status %d)", status)
	}

	base := corevideo.CVPixelBufferGetBaseAddress(t.pixelBuf)
	if base == nil {
		corevideo.CVPixelBufferUnlockBaseAddress(t.pixelBuf, corevideo.KCVPixelBufferLock_ReadOnly)
		return nil, nil, fmt.Errorf("ane: pixel buffer base address is nil")
	}

	pinner := &runtime.Pinner{}
	pinner.Pin(t)
	return base, func() {
		corevideo.CVPixelBufferUnlockBaseAddress(t.pixelBuf, corevideo.KCVPixelBufferLock_ReadOnly)
		pinner.Unpin()
	}, nil
}

// IOSurface returns the raw IOSurface pointer backing this tensor.
// Returns nil if the tensor is closed.
func (t *IOSurfaceTensor) IOSurface() unsafe.Pointer {
	if t.pixelBuf == 0 {
		return nil
	}
	return corevideo.CVPixelBufferGetIOSurface(t.pixelBuf)
}

// NumElements returns the product of the shape dimensions.
func (t *IOSurfaceTensor) NumElements() int {
	n := 1
	for _, d := range t.shape {
		n *= d
	}
	return n
}

// ByteLen returns the tensor's data size in bytes (elements * 2 for fp16).
func (t *IOSurfaceTensor) ByteLen() int {
	return t.NumElements() * 2
}

// MLFeatureValue creates an MLFeatureValue wrapping this tensor's pixel buffer.
// Convenience for building input feature dictionaries.
func (t *IOSurfaceTensor) MLFeatureValue() (coreml.MLFeatureValue, error) {
	arr, err := t.MLMultiArray()
	if err != nil {
		return coreml.MLFeatureValue{}, err
	}
	fv := coreml.NewFeatureValueWithMultiArray(arr)
	if fv.ID == 0 {
		return coreml.MLFeatureValue{}, fmt.Errorf("ane: NewFeatureValueWithMultiArray returned nil")
	}
	return fv, nil
}

// Retain prevents the underlying pixel buffer from being released by adding
// a CFRetain. Returns a release function that must be called exactly once.
// Useful when passing the pixel buffer to CoreML and needing to ensure it
// survives until prediction completes.
func (t *IOSurfaceTensor) Retain() func() {
	corefoundation.CFRetain(corefoundation.CFTypeRef(t.pixelBuf))
	return func() {
		corefoundation.CFRelease(corefoundation.CFTypeRef(t.pixelBuf))
	}
}

// iosurfacePixelBufferAttrs returns a CFDictionary with
// kCVPixelBufferIOSurfacePropertiesKey set to an empty dictionary,
// which tells CVPixelBufferCreate to back the buffer with an IOSurface.
func iosurfacePixelBufferAttrs() corefoundation.CFDictionaryRef {
	key := cfString("IOSurfacePropertiesKey")
	emptyDict := corefoundation.CFDictionaryCreate(0, nil, nil, 0, nil, nil)
	// Convert CFDictionaryRef (uintptr) to unsafe.Pointer via a local,
	// matching the pattern in surface.go's cfString/cfInt helpers.
	emptyDictPtr := uintptr(emptyDict)
	val := *(*unsafe.Pointer)(unsafe.Pointer(&emptyDictPtr))

	dict := corefoundation.CFDictionaryCreate(
		0,
		unsafe.Pointer(&key),
		unsafe.Pointer(&val),
		1,
		nil,
		nil,
	)

	corefoundation.CFRelease(corefoundation.CFTypeRef(uintptr(key)))
	corefoundation.CFRelease(corefoundation.CFTypeRef(emptyDict))
	return dict
}

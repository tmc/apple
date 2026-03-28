//go:build darwin

package ane

import (
	"fmt"
	"math"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/iosurface"
)

// TensorLayout describes the compiled model's memory layout for a single tensor.
// Parsed from ANEModel.ModelAttributes() after compilation.
//
// For [1, C, H, W] tensors: each channel plane occupies H rows with
// PlaneStride = H * RowStride. The minimum row stride is 64 bytes
// (the ANE's alignment granularity). For the common [1, C, 1, S] case,
// RowStride == PlaneStride.
type TensorLayout struct {
	Channels    int // number of channels (C in NCHW)
	Width       int // spatial dimension (W in NCHW)
	Height      int // row dimension (H in NCHW), may be >1 for LLM sequence/heads
	ElemSize    int // bytes per element (2 for fp16, 4 for fp32)
	RowStride   int // byte stride between rows (64-byte aligned)
	PlaneStride int // byte stride between channels (= Height * RowStride)
}

// AllocSize returns the total IOSurface allocation size for this layout.
func (l TensorLayout) AllocSize() int {
	return l.Channels * l.PlaneStride
}

// RowStrideFor computes the minimum 64-byte-aligned row stride for a given
// width and element size.
func RowStrideFor(width, elemSize int) int {
	raw := width * elemSize
	const align = 64
	return (raw + align - 1) &^ (align - 1)
}

// LogicalBytes returns the number of logical data bytes (Channels * Width * Height * ElemSize),
// excluding stride padding.
func (l TensorLayout) LogicalBytes() int {
	return l.Channels * l.Width * l.Height * l.ElemSize
}

// LogicalElements returns the number of logical elements (Channels * Width * Height).
func (l TensorLayout) LogicalElements() int {
	return l.Channels * l.Width * l.Height
}

// createSurfaceForLayout creates an IOSurface matching the compiled model's layout.
//
// IOSurface dimensions map the NCHW tensor as:
//   - IOSurfaceWidth  = Width (spatial W)
//   - IOSurfaceHeight = Channels * Height (flattened C*H rows)
//   - BytesPerElement = ElemSize
//   - BytesPerRow     = RowStride (must be 64-byte aligned)
//   - AllocSize       = Channels * PlaneStride
//
// The RowStride is validated against the 64-byte ANE alignment granularity.
func createSurfaceForLayout(layout TensorLayout) (coregraphics.IOSurfaceRef, error) {
	// Validate stride alignment.
	if layout.RowStride%64 != 0 {
		return 0, fmt.Errorf("ane: RowStride %d is not 64-byte aligned for tensor [%d, %d, %d] x %d",
			layout.RowStride, layout.Channels, layout.Height, layout.Width, layout.ElemSize)
	}
	minStride := RowStrideFor(layout.Width, layout.ElemSize)
	if layout.RowStride < minStride {
		return 0, fmt.Errorf("ane: RowStride %d < minimum %d for Width=%d ElemSize=%d",
			layout.RowStride, minStride, layout.Width, layout.ElemSize)
	}

	alloc := layout.AllocSize()
	// IOSurface treats the surface as a 2D plane; we flatten C*H into Height.
	surfHeight := layout.Channels * layout.Height
	keys := [6]unsafe.Pointer{
		cfString("IOSurfaceWidth"),
		cfString("IOSurfaceHeight"),
		cfString("IOSurfaceBytesPerElement"),
		cfString("IOSurfaceBytesPerRow"),
		cfString("IOSurfaceAllocSize"),
		cfString("IOSurfacePixelFormat"),
	}
	values := [6]unsafe.Pointer{
		cfInt(layout.Width),
		cfInt(surfHeight),
		cfInt(layout.ElemSize),
		cfInt(layout.RowStride),
		cfInt(alloc),
		cfInt(0),
	}
	dict := corefoundation.CFDictionaryCreate(
		0,
		unsafe.Pointer(&keys[0]),
		unsafe.Pointer(&values[0]),
		6,
		nil,
		nil,
	)
	ref := iosurface.IOSurfaceCreate(corefoundation.CFDictionaryRef(dict))

	// Release the dictionary and all keys/values. CFDictionaryCreate with nil
	// callbacks does not retain, so these are still at +1 from creation.
	corefoundation.CFRelease(corefoundation.CFTypeRef(dict))
	for i := range keys {
		corefoundation.CFRelease(corefoundation.CFTypeRef(uintptr(keys[i])))
		corefoundation.CFRelease(corefoundation.CFTypeRef(uintptr(values[i])))
	}

	if ref == 0 {
		return 0, fmt.Errorf("ane: failed to create IOSurface for tensor [%d, %d, %d] x %d (alloc=%d stride=%d)",
			layout.Channels, layout.Height, layout.Width, layout.ElemSize, alloc, layout.RowStride)
	}
	return coregraphics.IOSurfaceRef(ref), nil
}

// pageAlign rounds up to the nearest page boundary (16384 bytes, matching ANE).
func pageAlign(n int) int {
	const pageSize = 16384
	return (n + pageSize - 1) &^ (pageSize - 1)
}

// cfString creates a CFString from a Go string.
//
//go:nosplit
func cfString(s string) unsafe.Pointer {
	ref := corefoundation.CFStringCreateWithCString(0, s, 0x08000100) // kCFStringEncodingUTF8
	p := uintptr(ref)
	return *(*unsafe.Pointer)(unsafe.Pointer(&p))
}

// cfInt creates a CFNumber from an int.
//
//go:nosplit
func cfInt(v int) unsafe.Pointer {
	val := int64(v)
	ref := corefoundation.CFNumberCreate(0, corefoundation.KCFNumberSInt64Type, unsafe.Pointer(&val))
	p := uintptr(ref)
	return *(*unsafe.Pointer)(unsafe.Pointer(&p))
}

// writeSurface copies data from src into the IOSurface.
func writeSurface(ref coregraphics.IOSurfaceRef, src []byte) error {
	surfRef := iosurface.IOSurfaceRef(ref)
	iosurface.IOSurfaceLock(surfRef, 0, nil)
	base := iosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		iosurface.IOSurfaceUnlock(surfRef, 0, nil)
		return fmt.Errorf("ane: IOSurface base address is nil")
	}
	dst := unsafe.Slice((*byte)(base), len(src))
	copy(dst, src)
	iosurface.IOSurfaceUnlock(surfRef, 0, nil)
	return nil
}

// readSurface copies data from the IOSurface into dst.
func readSurface(ref coregraphics.IOSurfaceRef, dst []byte) error {
	surfRef := iosurface.IOSurfaceRef(ref)
	iosurface.IOSurfaceLock(surfRef, iosurface.KIOSurfaceLockReadOnly, nil)
	base := iosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		iosurface.IOSurfaceUnlock(surfRef, iosurface.KIOSurfaceLockReadOnly, nil)
		return fmt.Errorf("ane: IOSurface base address is nil")
	}
	src := unsafe.Slice((*byte)(base), len(dst))
	copy(dst, src)
	iosurface.IOSurfaceUnlock(surfRef, iosurface.KIOSurfaceLockReadOnly, nil)
	return nil
}

// surfaceSize returns the allocation size of an IOSurface.
func surfaceSize(ref coregraphics.IOSurfaceRef) int {
	return int(iosurface.IOSurfaceGetAllocSize(iosurface.IOSurfaceRef(ref)))
}

// writeStridedFP16 writes float32 data as fp16 to an IOSurface using the given layout.
// Data is in channel-first order: data[c*(H*W) + h*W + w].
// Byte offset: c*PlaneStride + h*RowStride + w*ElemSize.
func writeStridedFP16WithLayout(ref coregraphics.IOSurfaceRef, data []float32, layout TensorLayout) error {
	surfRef := iosurface.IOSurfaceRef(ref)
	iosurface.IOSurfaceLock(surfRef, 0, nil)
	base := iosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		iosurface.IOSurfaceUnlock(surfRef, 0, nil)
		return fmt.Errorf("ane: IOSurface base address is nil")
	}
	alloc := layout.AllocSize()
	dst := unsafe.Slice((*byte)(base), alloc)
	writeStridedFP16Bytes(dst, data, layout)

	iosurface.IOSurfaceUnlock(surfRef, 0, nil)
	return nil
}

// writeStridedFP16ChannelsWithLayout writes float32 data as fp16 starting at the
// given channel offset. Data is in channel-first order with the same per-channel
// shape as layout.
func writeStridedFP16ChannelsWithLayout(ref coregraphics.IOSurfaceRef, data []float32, layout TensorLayout, channel int) error {
	channelElems := layout.Height * layout.Width
	if channelElems <= 0 {
		return fmt.Errorf("ane: invalid fp16 layout channel elements %d", channelElems)
	}
	if len(data)%channelElems != 0 {
		return fmt.Errorf("ane: fp16 channel data length %d is not a multiple of channel size %d", len(data), channelElems)
	}
	channels := len(data) / channelElems
	if channel < 0 || channel+channels > layout.Channels {
		return fmt.Errorf("ane: fp16 channel range [%d,%d) out of bounds [0,%d)", channel, channel+channels, layout.Channels)
	}
	sub := layout
	sub.Channels = channels
	surfRef := iosurface.IOSurfaceRef(ref)
	iosurface.IOSurfaceLock(surfRef, 0, nil)
	base := iosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		iosurface.IOSurfaceUnlock(surfRef, 0, nil)
		return fmt.Errorf("ane: IOSurface base address is nil")
	}
	alloc := layout.AllocSize()
	dst := unsafe.Slice((*byte)(base), alloc)
	off := channel * layout.PlaneStride
	if off < 0 || off > len(dst) {
		iosurface.IOSurfaceUnlock(surfRef, 0, nil)
		return fmt.Errorf("ane: fp16 channel offset %d out of bounds", off)
	}
	writeStridedFP16Bytes(dst[off:], data, sub)
	iosurface.IOSurfaceUnlock(surfRef, 0, nil)
	return nil
}

// readStridedFP16WithLayout reads fp16 data from an IOSurface into float32s using the given layout.
func readStridedFP16WithLayout(ref coregraphics.IOSurfaceRef, data []float32, layout TensorLayout) error {
	surfRef := iosurface.IOSurfaceRef(ref)
	iosurface.IOSurfaceLock(surfRef, iosurface.KIOSurfaceLockReadOnly, nil)
	base := iosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		iosurface.IOSurfaceUnlock(surfRef, iosurface.KIOSurfaceLockReadOnly, nil)
		return fmt.Errorf("ane: IOSurface base address is nil")
	}
	alloc := layout.AllocSize()
	src := unsafe.Slice((*byte)(base), alloc)
	readStridedFP16Bytes(data, src, layout)

	iosurface.IOSurfaceUnlock(surfRef, iosurface.KIOSurfaceLockReadOnly, nil)
	return nil
}

// readStridedFP16ChannelsWithLayout reads fp16 data from the given channel
// offset into float32s using the same per-channel shape as layout.
func readStridedFP16ChannelsWithLayout(ref coregraphics.IOSurfaceRef, data []float32, layout TensorLayout, channel int) error {
	channelElems := layout.Height * layout.Width
	if channelElems <= 0 {
		return fmt.Errorf("ane: invalid fp16 layout channel elements %d", channelElems)
	}
	if len(data)%channelElems != 0 {
		return fmt.Errorf("ane: fp16 channel data length %d is not a multiple of channel size %d", len(data), channelElems)
	}
	channels := len(data) / channelElems
	if channel < 0 || channel+channels > layout.Channels {
		return fmt.Errorf("ane: fp16 channel range [%d,%d) out of bounds [0,%d)", channel, channel+channels, layout.Channels)
	}
	sub := layout
	sub.Channels = channels
	surfRef := iosurface.IOSurfaceRef(ref)
	iosurface.IOSurfaceLock(surfRef, iosurface.KIOSurfaceLockReadOnly, nil)
	base := iosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		iosurface.IOSurfaceUnlock(surfRef, iosurface.KIOSurfaceLockReadOnly, nil)
		return fmt.Errorf("ane: IOSurface base address is nil")
	}
	alloc := layout.AllocSize()
	src := unsafe.Slice((*byte)(base), alloc)
	off := channel * layout.PlaneStride
	if off < 0 || off > len(src) {
		iosurface.IOSurfaceUnlock(surfRef, iosurface.KIOSurfaceLockReadOnly, nil)
		return fmt.Errorf("ane: fp16 channel offset %d out of bounds", off)
	}
	readStridedFP16Bytes(data, src[off:], sub)
	iosurface.IOSurfaceUnlock(surfRef, iosurface.KIOSurfaceLockReadOnly, nil)
	return nil
}

// writeStridedF32WithLayout writes float32 data to an IOSurface using the given layout.
func writeStridedF32WithLayout(ref coregraphics.IOSurfaceRef, data []float32, layout TensorLayout) error {
	surfRef := iosurface.IOSurfaceRef(ref)
	iosurface.IOSurfaceLock(surfRef, 0, nil)
	base := iosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		iosurface.IOSurfaceUnlock(surfRef, 0, nil)
		return fmt.Errorf("ane: IOSurface base address is nil")
	}
	alloc := layout.AllocSize()
	dst := unsafe.Slice((*byte)(base), alloc)
	writeStridedF32Bytes(dst, data, layout)

	iosurface.IOSurfaceUnlock(surfRef, 0, nil)
	return nil
}

// readStridedF32WithLayout reads float32 data from an IOSurface using the given layout.
func readStridedF32WithLayout(ref coregraphics.IOSurfaceRef, data []float32, layout TensorLayout) error {
	surfRef := iosurface.IOSurfaceRef(ref)
	iosurface.IOSurfaceLock(surfRef, iosurface.KIOSurfaceLockReadOnly, nil)
	base := iosurface.IOSurfaceGetBaseAddress(surfRef)
	if base == nil {
		iosurface.IOSurfaceUnlock(surfRef, iosurface.KIOSurfaceLockReadOnly, nil)
		return fmt.Errorf("ane: IOSurface base address is nil")
	}
	alloc := layout.AllocSize()
	src := unsafe.Slice((*byte)(base), alloc)
	readStridedF32Bytes(data, src, layout)

	iosurface.IOSurfaceUnlock(surfRef, iosurface.KIOSurfaceLockReadOnly, nil)
	return nil
}

// float32ToFP16 converts a float32 to IEEE 754 half-precision (package-internal).
func float32ToFP16(f float32) uint16 {
	b := math.Float32bits(f)
	sign := (b >> 16) & 0x8000
	exp := int((b>>23)&0xFF) - 127 + 15
	frac := b & 0x7FFFFF

	switch {
	case exp <= 0:
		return uint16(sign)
	case exp >= 31:
		return uint16(sign | 0x7C00)
	default:
		return uint16(sign | uint32(exp)<<10 | (frac >> 13))
	}
}

// fp16ToFloat32 converts an IEEE 754 half-precision value to float32 (package-internal).
func fp16ToFloat32(h uint16) float32 {
	sign := uint32(h>>15) & 1
	exp := uint32(h>>10) & 0x1F
	frac := uint32(h) & 0x3FF

	switch {
	case exp == 0:
		if frac == 0 {
			return math.Float32frombits(sign << 31)
		}
		for frac&0x400 == 0 {
			frac <<= 1
			exp--
		}
		exp++
		frac &= 0x3FF
		fallthrough
	case exp < 31:
		return math.Float32frombits(sign<<31 | (exp+127-15)<<23 | frac<<13)
	default:
		return math.Float32frombits(sign<<31 | 0x7F800000 | frac<<13)
	}
}

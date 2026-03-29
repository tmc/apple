//go:build darwin

package ane

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Dtype describes the element type of a TensorView.
type Dtype uint8

const (
	DtypeFloat16 Dtype = iota
	DtypeFloat32
	DtypeFloat64
	DtypeInt32
)

// ByteSize returns the number of bytes per element.
func (d Dtype) ByteSize() int {
	switch d {
	case DtypeFloat16:
		return 2
	case DtypeFloat32, DtypeInt32:
		return 4
	case DtypeFloat64:
		return 8
	default:
		return 0
	}
}

func (d Dtype) String() string {
	switch d {
	case DtypeFloat16:
		return "float16"
	case DtypeFloat32:
		return "float32"
	case DtypeFloat64:
		return "float64"
	case DtypeInt32:
		return "int32"
	default:
		return fmt.Sprintf("dtype(%d)", d)
	}
}

// TensorView is a read-only, non-owning view over contiguous or strided
// tensor memory. It captures everything a downstream consumer (MLX, Metal,
// memcpy) needs to interpret the buffer without copying.
//
// A TensorView does not own the memory it points to. The lifetime of Ptr
// depends on how the view was created:
//   - WithMultiArrayBytes: Ptr valid only within the callback.
//   - WithMultiArrayMutableBytes: Ptr valid only within the callback.
//   - OutputTensorView: Ptr valid while the backing MLMultiArray is retained
//     and the returned unlock function has not been called.
//   - DeprecatedTensorViewFromMultiArray: Ptr valid while the MLMultiArray
//     is retained. Uses deprecated dataPointer.
type TensorView struct {
	// Ptr is the raw pointer to the first element.
	Ptr unsafe.Pointer

	// Shape is the logical shape (e.g. [1, 4096] for a hidden state).
	Shape []int

	// Strides is the per-dimension stride in elements (not bytes).
	// len(Strides) == len(Shape). For a contiguous row-major tensor,
	// Strides[i] == product(Shape[i+1:]).
	Strides []int

	// Dtype is the element type.
	Dtype Dtype

	// ByteLen is the total number of bytes accessible from Ptr.
	// For contiguous tensors this equals product(Shape) * Dtype.ByteSize().
	// For strided tensors it is the distance from Ptr to the end of the
	// last element, which may exceed the logical element count due to padding.
	ByteLen int
}

// IsContiguous reports whether the tensor is laid out in row-major order
// with no padding between elements.
func (v *TensorView) IsContiguous() bool {
	if len(v.Shape) == 0 {
		return true
	}
	expected := 1
	for i := len(v.Shape) - 1; i >= 0; i-- {
		if v.Strides[i] != expected {
			return false
		}
		expected *= v.Shape[i]
	}
	return true
}

// NumElements returns the product of the shape dimensions.
func (v *TensorView) NumElements() int {
	n := 1
	for _, d := range v.Shape {
		n *= d
	}
	return n
}

// Float32Slice returns the view's data as a []float32 slice backed by the
// same memory (no copy). Panics if Dtype is not DtypeFloat32 or the view
// is not contiguous.
//
// The returned slice is valid only while the backing memory is valid.
func (v *TensorView) Float32Slice() []float32 {
	if v.Dtype != DtypeFloat32 {
		panic(fmt.Sprintf("ane: Float32Slice called on %s tensor", v.Dtype))
	}
	if !v.IsContiguous() {
		panic("ane: Float32Slice called on non-contiguous tensor")
	}
	n := v.NumElements()
	return unsafe.Slice((*float32)(v.Ptr), n)
}

// --- Supported (non-deprecated) scoped access ---

// WithMultiArrayBytes calls fn with a read-only scoped TensorView over an
// MLMultiArray's data using getBytesWithHandler: (macOS 12.3+, iOS 15.4+).
//
// The TensorView and its Ptr are valid only for the duration of fn.
// Do not retain the view or its Ptr after fn returns.
//
// The strides come from the MLMultiArray's Strides property. The buffer
// may not be contiguous; check IsContiguous before assuming dense layout.
func WithMultiArrayBytes(arr coreml.MLMultiArray, fn func(view *TensorView)) error {
	if arr.ID == 0 {
		return fmt.Errorf("ane: nil MLMultiArray")
	}

	dt, err := mlDtype(arr.DataType())
	if err != nil {
		return err
	}

	shapeNums := arr.Shape()
	strideNums := arr.Strides()
	if len(shapeNums) == 0 {
		return fmt.Errorf("ane: MLMultiArray shape is empty")
	}
	if len(strideNums) != len(shapeNums) {
		return fmt.Errorf("ane: MLMultiArray strides length %d != shape length %d", len(strideNums), len(shapeNums))
	}

	shape := nsNumbersToInts(shapeNums)
	strides := nsNumbersToInts(strideNums)

	block := objc.NewBlock(func(_ objc.Block, bytes unsafe.Pointer, size int) {
		fn(&TensorView{
			Ptr:     bytes,
			Shape:   shape,
			Strides: strides,
			Dtype:   dt,
			ByteLen: size,
		})
	})
	defer block.Release()
	objc.Send[objc.ID](arr.ID, objc.Sel("getBytesWithHandler:"), objc.ID(block))
	return nil
}

// WithMultiArrayMutableBytes calls fn with a mutable scoped TensorView over
// an MLMultiArray's data using getMutableBytesWithHandler: (macOS 12.3+).
//
// The strides in the TensorView come from the handler's strides parameter,
// NOT from arr.Strides. getMutableBytesWithHandler: may copy the buffer
// (copy-on-write) and the resulting strides can differ from the original.
// Always use the view's Strides, not the array's.
//
// The TensorView and its Ptr are valid only for the duration of fn.
func WithMultiArrayMutableBytes(arr coreml.MLMultiArray, fn func(view *TensorView)) error {
	if arr.ID == 0 {
		return fmt.Errorf("ane: nil MLMultiArray")
	}

	dt, err := mlDtype(arr.DataType())
	if err != nil {
		return err
	}

	shapeNums := arr.Shape()
	if len(shapeNums) == 0 {
		return fmt.Errorf("ane: MLMultiArray shape is empty")
	}
	shape := nsNumbersToInts(shapeNums)

	block := objc.NewBlock(func(_ objc.Block, mutableBytes unsafe.Pointer, size int, stridesID objc.ID) {
		stridesArr := foundation.NSArrayFromID(stridesID)
		count := stridesArr.Count()
		strides := make([]int, count)
		for i := uint(0); i < count; i++ {
			obj := stridesArr.ObjectAtIndex(i)
			strides[i] = foundation.NSNumberFromID(obj.GetID()).IntegerValue()
		}
		fn(&TensorView{
			Ptr:     mutableBytes,
			Shape:   shape,
			Strides: strides,
			Dtype:   dt,
			ByteLen: size,
		})
	})
	defer block.Release()
	objc.Send[objc.ID](arr.ID, objc.Sel("getMutableBytesWithHandler:"), objc.ID(block))
	return nil
}

// --- Non-deprecated long-lived access via PixelBuffer ---

// OutputTensorView creates a long-lived TensorView from an MLMultiArray
// that is backed by a CVPixelBuffer (IOSurface). This is the non-deprecated
// path for persistent pointer access to CoreML outputs.
//
// Returns an error if the MLMultiArray's pixelBuffer property is nil, which
// means CoreML did not use IOSurface-backed storage for this output.
// IOSurface backing is typical for ANE fp16 outputs but not guaranteed.
//
// The returned view is valid until unlock is called. The caller must call
// unlock exactly once after all access to Ptr is complete:
//
//	view, unlock, err := OutputTensorView(arr)
//	if err != nil { ... }
//	defer unlock()
//	// use view.Ptr — valid until unlock()
//
// The MLMultiArray (and its parent prediction result) must remain retained
// for the lifetime of the view. Call runtime.KeepAlive(arr) after unlock
// if there is any risk the GC could collect it.
func OutputTensorView(arr coreml.MLMultiArray) (view *TensorView, unlock func(), err error) {
	if arr.ID == 0 {
		return nil, nil, fmt.Errorf("ane: nil MLMultiArray")
	}

	pb := arr.PixelBuffer()
	if pb == 0 {
		return nil, nil, fmt.Errorf("ane: MLMultiArray has no backing pixelBuffer (not IOSurface-backed)")
	}

	dt, err := mlDtype(arr.DataType())
	if err != nil {
		return nil, nil, err
	}

	shapeNums := arr.Shape()
	strideNums := arr.Strides()
	if len(shapeNums) == 0 {
		return nil, nil, fmt.Errorf("ane: MLMultiArray shape is empty")
	}
	if len(strideNums) != len(shapeNums) {
		return nil, nil, fmt.Errorf("ane: MLMultiArray strides length %d != shape length %d", len(strideNums), len(shapeNums))
	}

	shape := nsNumbersToInts(shapeNums)
	strides := nsNumbersToInts(strideNums)

	pixelBuf := corevideo.CVPixelBufferRef(pb)
	status := corevideo.CVPixelBufferLockBaseAddress(pixelBuf, corevideo.KCVPixelBufferLock_ReadOnly)
	if status != 0 {
		return nil, nil, fmt.Errorf("ane: CVPixelBufferLockBaseAddress failed (status %d)", status)
	}

	base := corevideo.CVPixelBufferGetBaseAddress(pixelBuf)
	if base == nil {
		corevideo.CVPixelBufferUnlockBaseAddress(pixelBuf, corevideo.KCVPixelBufferLock_ReadOnly)
		return nil, nil, fmt.Errorf("ane: pixel buffer base address is nil")
	}

	// Compute byte length from pixel buffer geometry.
	// bytesPerRow may include row padding beyond width*elementSize.
	bytesPerRow := int(corevideo.CVPixelBufferGetBytesPerRow(pixelBuf))
	height := int(corevideo.CVPixelBufferGetHeight(pixelBuf))
	byteLen := bytesPerRow * height

	unlocked := false
	unlock = func() {
		if !unlocked {
			unlocked = true
			corevideo.CVPixelBufferUnlockBaseAddress(pixelBuf, corevideo.KCVPixelBufferLock_ReadOnly)
			runtime.KeepAlive(arr)
		}
	}

	return &TensorView{
		Ptr:     base,
		Shape:   shape,
		Strides: strides,
		Dtype:   dt,
		ByteLen: byteLen,
	}, unlock, nil
}

// --- Deprecated long-lived access via dataPointer ---

// DeprecatedTensorViewFromMultiArray creates a TensorView using the
// deprecated dataPointer property.
//
// Deprecated: Use WithMultiArrayBytes for scoped access, or
// OutputTensorView for long-lived access to IOSurface-backed outputs.
// dataPointer will be removed in a future macOS/iOS release.
//
// The returned view borrows memory from arr. The caller must keep arr
// retained for the lifetime of the view. After the last access through
// the view's Ptr, call runtime.KeepAlive(arr).
func DeprecatedTensorViewFromMultiArray(arr coreml.MLMultiArray) (*TensorView, error) {
	if arr.ID == 0 {
		return nil, fmt.Errorf("ane: nil MLMultiArray")
	}

	ptr := objc.Send[unsafe.Pointer](arr.ID, objc.Sel("dataPointer"))
	if ptr == nil {
		return nil, fmt.Errorf("ane: MLMultiArray data pointer is nil")
	}

	dt, err := mlDtype(arr.DataType())
	if err != nil {
		return nil, err
	}

	shapeNums := arr.Shape()
	strideNums := arr.Strides()
	if len(shapeNums) == 0 {
		return nil, fmt.Errorf("ane: MLMultiArray shape is empty")
	}
	if len(strideNums) != len(shapeNums) {
		return nil, fmt.Errorf("ane: MLMultiArray strides length %d != shape length %d", len(strideNums), len(shapeNums))
	}

	shape := nsNumbersToInts(shapeNums)
	strides := nsNumbersToInts(strideNums)

	var maxOffset int
	for i := range shape {
		if shape[i] > 0 {
			maxOffset += (shape[i] - 1) * strides[i]
		}
	}
	byteLen := (maxOffset + 1) * dt.ByteSize()

	runtime.KeepAlive(arr)

	return &TensorView{
		Ptr:     ptr,
		Shape:   shape,
		Strides: strides,
		Dtype:   dt,
		ByteLen: byteLen,
	}, nil
}

// DeprecatedRetainedTensorView creates a TensorView and retains the
// MLMultiArray for pointer validity. Uses deprecated dataPointer.
//
// Deprecated: Use WithMultiArrayBytes or OutputTensorView instead.
func DeprecatedRetainedTensorView(arr coreml.MLMultiArray) (view *TensorView, release func(), err error) {
	view, err = DeprecatedTensorViewFromMultiArray(arr)
	if err != nil {
		return nil, nil, err
	}

	owner := objectivec.Object{ID: arr.ID}
	owner.Retain()
	release = func() {
		owner.Release()
	}
	return view, release, nil
}

// --- Compatibility aliases ---

// TensorViewFromMultiArray is an alias for DeprecatedTensorViewFromMultiArray.
//
// Deprecated: Use WithMultiArrayBytes or OutputTensorView instead.
var TensorViewFromMultiArray = DeprecatedTensorViewFromMultiArray

// RetainedTensorView is an alias for DeprecatedRetainedTensorView.
//
// Deprecated: Use WithMultiArrayBytes or OutputTensorView instead.
var RetainedTensorView = DeprecatedRetainedTensorView

// --- Internal helpers ---

func mlDtype(dt coreml.MLMultiArrayDataType) (Dtype, error) {
	switch dt {
	case coreml.MLMultiArrayDataTypeFloat16:
		return DtypeFloat16, nil
	case coreml.MLMultiArrayDataTypeFloat32:
		return DtypeFloat32, nil
	case coreml.MLMultiArrayDataTypeDouble:
		return DtypeFloat64, nil
	case coreml.MLMultiArrayDataTypeInt32:
		return DtypeInt32, nil
	default:
		return 0, fmt.Errorf("ane: unsupported MLMultiArrayDataType %d", dt)
	}
}

func nsNumbersToInts(nums []foundation.NSNumber) []int {
	out := make([]int, len(nums))
	for i, n := range nums {
		out[i] = n.IntegerValue()
	}
	return out
}

//go:build darwin

package ane

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/foundation"
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
// A TensorView does not own the memory it points to. The caller must ensure
// the backing object (typically an MLMultiArray retained via its prediction
// result) outlives all use of the view. Call runtime.KeepAlive on the owner
// after the last access through Ptr.
type TensorView struct {
	// Ptr is the raw pointer to the first element. Read-only — do not write
	// through this pointer for CoreML-backed views.
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

// TensorViewFromMultiArray creates a TensorView from a CoreML MLMultiArray.
//
// The returned view borrows memory from arr. The caller must keep arr (and
// its parent prediction result) retained for the lifetime of the view.
// After the last access through the view's Ptr, call:
//
//	runtime.KeepAlive(arr)
//
// Returns an error if arr is nil, has an unsupported data type, or has a
// nil data pointer.
func TensorViewFromMultiArray(arr coreml.MLMultiArray) (*TensorView, error) {
	if arr.ID == 0 {
		return nil, fmt.Errorf("ane: nil MLMultiArray")
	}

	ptr := arr.DataPointer()
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

	// Compute byte length: offset of last element + element size.
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

// RetainedTensorView creates a TensorView from an MLMultiArray and retains
// the array to guarantee pointer validity. The returned release function
// must be called exactly once after all access to the view is complete.
//
// This is a convenience for callers that want self-contained ownership:
//
//	view, release, err := RetainedTensorView(arr)
//	defer release()
//	// use view.Ptr ...
func RetainedTensorView(arr coreml.MLMultiArray) (view *TensorView, release func(), err error) {
	view, err = TensorViewFromMultiArray(arr)
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

// Float32Slice returns the view's data as a []float32 slice backed by the
// same memory (no copy). Panics if Dtype is not DtypeFloat32 or the view
// is not contiguous.
//
// The returned slice is valid only while the backing MLMultiArray is retained.
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

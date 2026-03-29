// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLMultiArray] class.
var (
	_MLMultiArrayClass     MLMultiArrayClass
	_MLMultiArrayClassOnce sync.Once
)

func getMLMultiArrayClass() MLMultiArrayClass {
	_MLMultiArrayClassOnce.Do(func() {
		_MLMultiArrayClass = MLMultiArrayClass{class: objc.GetClass("MLMultiArray")}
	})
	return _MLMultiArrayClass
}

// GetMLMultiArrayClass returns the class object for MLMultiArray.
func GetMLMultiArrayClass() MLMultiArrayClass {
	return getMLMultiArrayClass()
}

type MLMultiArrayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLMultiArrayClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLMultiArrayClass) Alloc() MLMultiArray {
	rv := objc.Send[MLMultiArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLMultiArray.BackingPixelBufferWasLocked]
//   - [MLMultiArray.Bytes]
//   - [MLMultiArray.CopyIntoMultiArrayError]
//   - [MLMultiArray.DataPointer]
//   - [MLMultiArray.DebugQuickLookObject]
//   - [MLMultiArray.DoublePointer]
//   - [MLMultiArray.FillWithNumber]
//   - [MLMultiArray.Float32Pointer]
//   - [MLMultiArray.GetBytesWithHandler]
//   - [MLMultiArray.GetContiguousFirstMajorFloat32BufferWithHandler]
//   - [MLMultiArray.IsContiguous]
//   - [MLMultiArray.IsContiguousInOrder]
//   - [MLMultiArray.IsEqualToMultiArray]
//   - [MLMultiArray.MtlBuffer]
//   - [MLMultiArray.MultiArrayBuffer]
//   - [MLMultiArray.MultiArrayViewExpandingDimensionsAtAxis]
//   - [MLMultiArray.MutableBytes]
//   - [MLMultiArray.NumberArray]
//   - [MLMultiArray.NumberAtOffset]
//   - [MLMultiArray.NumberOfBytesPerElement]
//   - [MLMultiArray.ObjectAtIndexedSubscript]
//   - [MLMultiArray.ObjectForKeyedSubscript]
//   - [MLMultiArray.OffsetForKeyedSubscript]
//   - [MLMultiArray.SetNumberAtOffset]
//   - [MLMultiArray.SetObjectAtIndexedSubscript]
//   - [MLMultiArray.SetObjectForKeyedSubscript]
//   - [MLMultiArray.SetRangeWithRawDataDestIndexError]
//   - [MLMultiArray.SliceAtOriginShapeSqueezeError]
//   - [MLMultiArray.Squeeze]
//   - [MLMultiArray.SqueezeDimensionsError]
//   - [MLMultiArray.VectorizeIntoMultiArrayStorageOrderError]
//   - [MLMultiArray.InitWithArrayDataType]
//   - [MLMultiArray.InitWithBytesNoCopyShapeDataTypeStridesDeallocatorMutableShapedBufferProviderError]
//   - [MLMultiArray.InitWithBytesNoCopyShapeDataTypeStridesMutableShapedBufferProvider]
//   - [MLMultiArray.InitWithCoder]
//   - [MLMultiArray.InitWithScalarsShapeDataType]
//   - [MLMultiArray.InitWithShapeDataTypeStorageOrderBufferAlignment]
//   - [MLMultiArray.InitWithShapeDataTypeStorageOrderError]
//   - [MLMultiArray.InitWithShapeDataTypeStrides]
//   - [MLMultiArray.Contiguous]
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray
type MLMultiArray struct {
	objectivec.Object
}

// MLMultiArrayFromID constructs a [MLMultiArray] from an objc.ID.
func MLMultiArrayFromID(id objc.ID) MLMultiArray {
	return MLMultiArray{objectivec.Object{ID: id}}
}
// Ensure MLMultiArray implements IMLMultiArray.
var _ IMLMultiArray = MLMultiArray{}

// An interface definition for the [MLMultiArray] class.
//
// # Methods
//
//   - [IMLMultiArray.BackingPixelBufferWasLocked]
//   - [IMLMultiArray.Bytes]
//   - [IMLMultiArray.CopyIntoMultiArrayError]
//   - [IMLMultiArray.DataPointer]
//   - [IMLMultiArray.DebugQuickLookObject]
//   - [IMLMultiArray.DoublePointer]
//   - [IMLMultiArray.FillWithNumber]
//   - [IMLMultiArray.Float32Pointer]
//   - [IMLMultiArray.GetBytesWithHandler]
//   - [IMLMultiArray.GetContiguousFirstMajorFloat32BufferWithHandler]
//   - [IMLMultiArray.IsContiguous]
//   - [IMLMultiArray.IsContiguousInOrder]
//   - [IMLMultiArray.IsEqualToMultiArray]
//   - [IMLMultiArray.MtlBuffer]
//   - [IMLMultiArray.MultiArrayBuffer]
//   - [IMLMultiArray.MultiArrayViewExpandingDimensionsAtAxis]
//   - [IMLMultiArray.MutableBytes]
//   - [IMLMultiArray.NumberArray]
//   - [IMLMultiArray.NumberAtOffset]
//   - [IMLMultiArray.NumberOfBytesPerElement]
//   - [IMLMultiArray.ObjectAtIndexedSubscript]
//   - [IMLMultiArray.ObjectForKeyedSubscript]
//   - [IMLMultiArray.OffsetForKeyedSubscript]
//   - [IMLMultiArray.SetNumberAtOffset]
//   - [IMLMultiArray.SetObjectAtIndexedSubscript]
//   - [IMLMultiArray.SetObjectForKeyedSubscript]
//   - [IMLMultiArray.SetRangeWithRawDataDestIndexError]
//   - [IMLMultiArray.SliceAtOriginShapeSqueezeError]
//   - [IMLMultiArray.Squeeze]
//   - [IMLMultiArray.SqueezeDimensionsError]
//   - [IMLMultiArray.VectorizeIntoMultiArrayStorageOrderError]
//   - [IMLMultiArray.InitWithArrayDataType]
//   - [IMLMultiArray.InitWithBytesNoCopyShapeDataTypeStridesDeallocatorMutableShapedBufferProviderError]
//   - [IMLMultiArray.InitWithBytesNoCopyShapeDataTypeStridesMutableShapedBufferProvider]
//   - [IMLMultiArray.InitWithCoder]
//   - [IMLMultiArray.InitWithScalarsShapeDataType]
//   - [IMLMultiArray.InitWithShapeDataTypeStorageOrderBufferAlignment]
//   - [IMLMultiArray.InitWithShapeDataTypeStorageOrderError]
//   - [IMLMultiArray.InitWithShapeDataTypeStrides]
//   - [IMLMultiArray.Contiguous]
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray
type IMLMultiArray interface {
	objectivec.IObject

	// Topic: Methods

	BackingPixelBufferWasLocked() bool
	Bytes() objectivec.IObject
	CopyIntoMultiArrayError(array objectivec.IObject) (bool, error)
	DataPointer() unsafe.Pointer
	DebugQuickLookObject() objectivec.IObject
	DoublePointer() []float64
	FillWithNumber(number objectivec.IObject) bool
	Float32Pointer() unsafe.Pointer
	GetBytesWithHandler(handler func(unsafe.Pointer, int64))
	GetContiguousFirstMajorFloat32BufferWithHandler(handler VoidHandler)
	IsContiguous() bool
	IsContiguousInOrder(order int64) bool
	IsEqualToMultiArray(array objectivec.IObject) bool
	MtlBuffer() objectivec.IObject
	MultiArrayBuffer() unsafe.Pointer
	MultiArrayViewExpandingDimensionsAtAxis(axis int64) objectivec.IObject
	MutableBytes() unsafe.Pointer
	NumberArray() objectivec.IObject
	NumberAtOffset(offset uint64) objectivec.IObject
	NumberOfBytesPerElement() uint64
	ObjectAtIndexedSubscript(subscript int64) objectivec.IObject
	ObjectForKeyedSubscript(subscript objectivec.IObject) objectivec.IObject
	OffsetForKeyedSubscript(subscript objectivec.IObject) uint64
	SetNumberAtOffset(number objectivec.IObject, offset uint64)
	SetObjectAtIndexedSubscript(object objectivec.IObject, subscript int64)
	SetObjectForKeyedSubscript(object objectivec.IObject, subscript objectivec.IObject)
	SetRangeWithRawDataDestIndexError(data objectivec.IObject, index uint64) (bool, error)
	SliceAtOriginShapeSqueezeError(origin objectivec.IObject, shape objectivec.IObject, squeeze bool) (objectivec.IObject, error)
	Squeeze() objectivec.IObject
	SqueezeDimensionsError(dimensions objectivec.IObject) (objectivec.IObject, error)
	VectorizeIntoMultiArrayStorageOrderError(array objectivec.IObject, order int64) (bool, error)
	InitWithArrayDataType(array objectivec.IObject, type_ int64) MLMultiArray
	InitWithBytesNoCopyShapeDataTypeStridesDeallocatorMutableShapedBufferProviderError(copy_ unsafe.Pointer, shape objectivec.IObject, type_ int64, strides objectivec.IObject, deallocator func(), provider func()) (MLMultiArray, error)
	InitWithBytesNoCopyShapeDataTypeStridesMutableShapedBufferProvider(copy_ unsafe.Pointer, shape objectivec.IObject, type_ int64, strides objectivec.IObject, provider VoidHandler) MLMultiArray
	InitWithCoder(coder foundation.INSCoder) MLMultiArray
	InitWithScalarsShapeDataType(scalars objectivec.IObject, shape objectivec.IObject, type_ int64) MLMultiArray
	InitWithShapeDataTypeStorageOrderBufferAlignment(shape objectivec.IObject, type_ int64, order int64, alignment uint64) MLMultiArray
	InitWithShapeDataTypeStorageOrderError(shape objectivec.IObject, type_ int64, order int64) (MLMultiArray, error)
	InitWithShapeDataTypeStrides(shape objectivec.IObject, type_ int64, strides objectivec.IObject) MLMultiArray
	Contiguous() bool
}

// Init initializes the instance.
func (m MLMultiArray) Init() MLMultiArray {
	rv := objc.Send[MLMultiArray](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLMultiArray) Autorelease() MLMultiArray {
	rv := objc.Send[MLMultiArray](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLMultiArray creates a new MLMultiArray instance.
func NewMLMultiArray() MLMultiArray {
	class := getMLMultiArrayClass()
	rv := objc.Send[MLMultiArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithArray:dataType:
func NewMultiArrayWithArrayDataType(array objectivec.IObject, type_ int64) MLMultiArray {
	instance := getMLMultiArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArray:dataType:"), array, type_)
	return MLMultiArrayFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithCoder:
func NewMultiArrayWithCoder(coder objectivec.IObject) MLMultiArray {
	instance := getMLMultiArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLMultiArrayFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithScalars:shape:dataType:
func NewMultiArrayWithScalarsShapeDataType(scalars objectivec.IObject, shape objectivec.IObject, type_ int64) MLMultiArray {
	instance := getMLMultiArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithScalars:shape:dataType:"), scalars, shape, type_)
	return MLMultiArrayFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithShape:dataType:storageOrder:bufferAlignment:
func NewMultiArrayWithShapeDataTypeStorageOrderBufferAlignment(shape objectivec.IObject, type_ int64, order int64, alignment uint64) MLMultiArray {
	instance := getMLMultiArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithShape:dataType:storageOrder:bufferAlignment:"), shape, type_, order, alignment)
	return MLMultiArrayFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithShape:dataType:storageOrder:error:
func NewMultiArrayWithShapeDataTypeStorageOrderError(shape objectivec.IObject, type_ int64, order int64) (MLMultiArray, error) {
	var errorPtr objc.ID
	instance := getMLMultiArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithShape:dataType:storageOrder:error:"), shape, type_, order, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLMultiArray{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLMultiArrayFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithShape:dataType:strides:
func NewMultiArrayWithShapeDataTypeStrides(shape objectivec.IObject, type_ int64, strides objectivec.IObject) MLMultiArray {
	instance := getMLMultiArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithShape:dataType:strides:"), shape, type_, strides)
	return MLMultiArrayFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/copyIntoMultiArray:error:
func (m MLMultiArray) CopyIntoMultiArrayError(array objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("copyIntoMultiArray:error:"), array, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("copyIntoMultiArray:error: returned NO with nil NSError")
	}
	return rv, nil

}
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/debugQuickLookObject
func (m MLMultiArray) DebugQuickLookObject() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugQuickLookObject"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/doublePointer
func (m MLMultiArray) DoublePointer() []float64 {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("doublePointer"))
	return objc.ConvertSlice(rv, func(id objc.ID) float64 {
		return float64(id)
	})
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/fillWithNumber:
func (m MLMultiArray) FillWithNumber(number objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("fillWithNumber:"), number)
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/float32Pointer
func (m MLMultiArray) Float32Pointer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("float32Pointer"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/getBytesWithHandler:
func (m MLMultiArray) GetBytesWithHandler(handler func(unsafe.Pointer, int64)) {
	_block0 := objc.NewBlock(func(_ objc.Block, arg0 unsafe.Pointer, arg1 int64) { handler(arg0, arg1) })
	defer _block0.Release()
	objc.Send[objc.ID](m.ID, objc.Sel("getBytesWithHandler:"), objc.ID(_block0))
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/getContiguousFirstMajorFloat32BufferWithHandler:
func (m MLMultiArray) GetContiguousFirstMajorFloat32BufferWithHandler(handler VoidHandler) {
_block0, _ := NewVoidBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("getContiguousFirstMajorFloat32BufferWithHandler:"), _block0)
}
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/isContiguous
func (m MLMultiArray) IsContiguous() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isContiguous"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/isContiguousInOrder:
func (m MLMultiArray) IsContiguousInOrder(order int64) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isContiguousInOrder:"), order)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/isEqualToMultiArray:
func (m MLMultiArray) IsEqualToMultiArray(array objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isEqualToMultiArray:"), array)
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/multiArrayBuffer
func (m MLMultiArray) MultiArrayBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("multiArrayBuffer"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/multiArrayViewExpandingDimensionsAtAxis:
func (m MLMultiArray) MultiArrayViewExpandingDimensionsAtAxis(axis int64) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("multiArrayViewExpandingDimensionsAtAxis:"), axis)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/numberArray
func (m MLMultiArray) NumberArray() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("numberArray"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/numberAtOffset:
func (m MLMultiArray) NumberAtOffset(offset uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("numberAtOffset:"), offset)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/objectAtIndexedSubscript:
func (m MLMultiArray) ObjectAtIndexedSubscript(subscript int64) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectAtIndexedSubscript:"), subscript)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/objectForKeyedSubscript:
func (m MLMultiArray) ObjectForKeyedSubscript(subscript objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectForKeyedSubscript:"), subscript)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/offsetForKeyedSubscript:
func (m MLMultiArray) OffsetForKeyedSubscript(subscript objectivec.IObject) uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("offsetForKeyedSubscript:"), subscript)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/setNumber:atOffset:
func (m MLMultiArray) SetNumberAtOffset(number objectivec.IObject, offset uint64) {
	objc.Send[objc.ID](m.ID, objc.Sel("setNumber:atOffset:"), number, offset)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/setObject:atIndexedSubscript:
func (m MLMultiArray) SetObjectAtIndexedSubscript(object objectivec.IObject, subscript int64) {
	objc.Send[objc.ID](m.ID, objc.Sel("setObject:atIndexedSubscript:"), object, subscript)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/setObject:forKeyedSubscript:
func (m MLMultiArray) SetObjectForKeyedSubscript(object objectivec.IObject, subscript objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setObject:forKeyedSubscript:"), object, subscript)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/setRangeWithRawData:destIndex:error:
func (m MLMultiArray) SetRangeWithRawDataDestIndexError(data objectivec.IObject, index uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("setRangeWithRawData:destIndex:error:"), data, index, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setRangeWithRawData:destIndex:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/sliceAtOrigin:shape:squeeze:error:
func (m MLMultiArray) SliceAtOriginShapeSqueezeError(origin objectivec.IObject, shape objectivec.IObject, squeeze bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("sliceAtOrigin:shape:squeeze:error:"), origin, shape, squeeze, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/squeeze
func (m MLMultiArray) Squeeze() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("squeeze"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/squeezeDimensions:error:
func (m MLMultiArray) SqueezeDimensionsError(dimensions objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("squeezeDimensions:error:"), dimensions, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/vectorizeIntoMultiArray:storageOrder:error:
func (m MLMultiArray) VectorizeIntoMultiArrayStorageOrderError(array objectivec.IObject, order int64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("vectorizeIntoMultiArray:storageOrder:error:"), array, order, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("vectorizeIntoMultiArray:storageOrder:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithArray:dataType:
func (m MLMultiArray) InitWithArrayDataType(array objectivec.IObject, type_ int64) MLMultiArray {
	rv := objc.Send[MLMultiArray](m.ID, objc.Sel("initWithArray:dataType:"), array, type_)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithBytesNoCopy:shape:dataType:strides:deallocator:mutableShapedBufferProvider:error:
func (m MLMultiArray) InitWithBytesNoCopyShapeDataTypeStridesDeallocatorMutableShapedBufferProviderError(copy_ unsafe.Pointer, shape objectivec.IObject, type_ int64, strides objectivec.IObject, deallocator func(), provider func()) (MLMultiArray, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithBytesNoCopy:shape:dataType:strides:deallocator:mutableShapedBufferProvider:error:"), copy_, shape, type_, strides, deallocator, provider, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLMultiArray{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLMultiArrayFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithBytesNoCopy:shape:dataType:strides:mutableShapedBufferProvider:
func (m MLMultiArray) InitWithBytesNoCopyShapeDataTypeStridesMutableShapedBufferProvider(copy_ unsafe.Pointer, shape objectivec.IObject, type_ int64, strides objectivec.IObject, provider VoidHandler) MLMultiArray {
_block4, _ := NewVoidBlock(provider)
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithBytesNoCopy:shape:dataType:strides:mutableShapedBufferProvider:"), copy_, shape, type_, strides, _block4)
	return MLMultiArrayFromID(rv)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithCoder:
func (m MLMultiArray) InitWithCoder(coder foundation.INSCoder) MLMultiArray {
	rv := objc.Send[MLMultiArray](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithScalars:shape:dataType:
func (m MLMultiArray) InitWithScalarsShapeDataType(scalars objectivec.IObject, shape objectivec.IObject, type_ int64) MLMultiArray {
	rv := objc.Send[MLMultiArray](m.ID, objc.Sel("initWithScalars:shape:dataType:"), scalars, shape, type_)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithShape:dataType:storageOrder:bufferAlignment:
func (m MLMultiArray) InitWithShapeDataTypeStorageOrderBufferAlignment(shape objectivec.IObject, type_ int64, order int64, alignment uint64) MLMultiArray {
	rv := objc.Send[MLMultiArray](m.ID, objc.Sel("initWithShape:dataType:storageOrder:bufferAlignment:"), shape, type_, order, alignment)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithShape:dataType:storageOrder:error:
func (m MLMultiArray) InitWithShapeDataTypeStorageOrderError(shape objectivec.IObject, type_ int64, order int64) (MLMultiArray, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithShape:dataType:storageOrder:error:"), shape, type_, order, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLMultiArray{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLMultiArrayFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithShape:dataType:strides:
func (m MLMultiArray) InitWithShapeDataTypeStrides(shape objectivec.IObject, type_ int64, strides objectivec.IObject) MLMultiArray {
	rv := objc.Send[MLMultiArray](m.ID, objc.Sel("initWithShape:dataType:strides:"), shape, type_, strides)
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/_multiArrayByConcatenatingMultiArrays:alongAxis:dataType:
func (_MLMultiArrayClass MLMultiArrayClass) _multiArrayByConcatenatingMultiArraysAlongAxisDataType(arrays objectivec.IObject, axis int64, type_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("_multiArrayByConcatenatingMultiArrays:alongAxis:dataType:"), arrays, axis, type_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/_shapeOfNestedArray:error:
func (_MLMultiArrayClass MLMultiArrayClass) _shapeOfNestedArrayError(array objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("_shapeOfNestedArray:error:"), array, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// ShapeOfNestedArrayError is an exported wrapper for the private method _shapeOfNestedArrayError.
func (_MLMultiArrayClass MLMultiArrayClass) ShapeOfNestedArrayError(array objectivec.IObject) (objectivec.IObject, error) {
	return _MLMultiArrayClass._shapeOfNestedArrayError(array)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/cppStorageOrder:
func (_MLMultiArrayClass MLMultiArrayClass) CppStorageOrder(order int64) int {
	rv := objc.Send[int](objc.ID(_MLMultiArrayClass.class), objc.Sel("cppStorageOrder:"), order)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/doubleMatrixWithValues:error:
func (_MLMultiArrayClass MLMultiArrayClass) DoubleMatrixWithValuesError(values objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("doubleMatrixWithValues:error:"), values, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/doubleMultiArrayWithCopyOfMultiArray:
func (_MLMultiArrayClass MLMultiArrayClass) DoubleMultiArrayWithCopyOfMultiArray(array objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("doubleMultiArrayWithCopyOfMultiArray:"), array)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/doubleMultiArrayWithShape:valueArray:error:
func (_MLMultiArrayClass MLMultiArrayClass) DoubleMultiArrayWithShapeValueArrayError(shape objectivec.IObject, array objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("doubleMultiArrayWithShape:valueArray:error:"), shape, array, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/doubleVectorWithValues:
func (_MLMultiArrayClass MLMultiArrayClass) DoubleVectorWithValues(values objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("doubleVectorWithValues:"), values)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/float32MatrixWithValues:error:
func (_MLMultiArrayClass MLMultiArrayClass) Float32MatrixWithValuesError(values objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("float32MatrixWithValues:error:"), values, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/getShapeOfArrayOfSameLengthArrays:numberOfRows:numberOfColumns:error:
func (_MLMultiArrayClass MLMultiArrayClass) GetShapeOfArrayOfSameLengthArraysNumberOfRowsNumberOfColumnsError(arrays objectivec.IObject) (uint64, uint64, error) {
	var rows uint64
	var columns uint64
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLMultiArrayClass.class), objc.Sel("getShapeOfArrayOfSameLengthArrays:numberOfRows:numberOfColumns:error:"), arrays, unsafe.Pointer(&rows), unsafe.Pointer(&columns), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, 0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0, 0, errors.New("getShapeOfArrayOfSameLengthArrays:numberOfRows:numberOfColumns:error: returned NO with nil NSError")
	}
	return rows, columns, nil
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/multiArrayByConcatenatingMultiArrays:alongAxis:dataType:
func (_MLMultiArrayClass MLMultiArrayClass) MultiArrayByConcatenatingMultiArraysAlongAxisDataType(arrays objectivec.IObject, axis int64, type_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("multiArrayByConcatenatingMultiArrays:alongAxis:dataType:"), arrays, axis, type_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/newMultiArrayForState:strides:dtype:error:
func (_MLMultiArrayClass MLMultiArrayClass) NewMultiArrayForStateStridesDtypeError(state objectivec.IObject, strides objectivec.IObject, dtype int64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("newMultiArrayForState:strides:dtype:error:"), state, strides, dtype, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/stringForDataType:
func (_MLMultiArrayClass MLMultiArrayClass) StringForDataType(type_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("stringForDataType:"), type_)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/supportsSecureCoding
func (_MLMultiArrayClass MLMultiArrayClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLMultiArrayClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/validateMultiArrays:forConcatenatingAlongAxis:normalizedAxis:reason:
func (_MLMultiArrayClass MLMultiArrayClass) ValidateMultiArraysForConcatenatingAlongAxisNormalizedAxisReason(arrays objectivec.IObject, axis int64, reason []objectivec.IObject) (uint64, bool) {
	var axis2 uint64
	rv := objc.Send[bool](objc.ID(_MLMultiArrayClass.class), objc.Sel("validateMultiArrays:forConcatenatingAlongAxis:normalizedAxis:reason:"), arrays, axis, unsafe.Pointer(&axis2), reason)
	return axis2, rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/validateNestedArray:error:
func (_MLMultiArrayClass MLMultiArrayClass) ValidateNestedArrayError(array objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLMultiArrayClass.class), objc.Sel("validateNestedArray:error:"), array, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateNestedArray:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/backingPixelBufferWasLocked
func (m MLMultiArray) BackingPixelBufferWasLocked() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("backingPixelBufferWasLocked"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/bytes
func (m MLMultiArray) Bytes() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("bytes"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/contiguous
func (m MLMultiArray) Contiguous() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("contiguous"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/dataPointer
func (m MLMultiArray) DataPointer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("dataPointer"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/mtlBuffer
func (m MLMultiArray) MtlBuffer() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mtlBuffer"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/mutableBytes
func (m MLMultiArray) MutableBytes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("mutableBytes"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/numberOfBytesPerElement
func (m MLMultiArray) NumberOfBytesPerElement() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("numberOfBytesPerElement"))
	return rv
}

// GetContiguousFirstMajorFloat32BufferWithHandlerSync is a synchronous wrapper around [MLMultiArray.GetContiguousFirstMajorFloat32BufferWithHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLMultiArray) GetContiguousFirstMajorFloat32BufferWithHandlerSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	m.GetContiguousFirstMajorFloat32BufferWithHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// InitWithBytesNoCopyShapeDataTypeStridesMutableShapedBufferProviderSync is a synchronous wrapper around [MLMultiArray.InitWithBytesNoCopyShapeDataTypeStridesMutableShapedBufferProvider].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLMultiArray) InitWithBytesNoCopyShapeDataTypeStridesMutableShapedBufferProviderSync(ctx context.Context, copy_ unsafe.Pointer, shape objectivec.IObject, type_ int64, strides objectivec.IObject) error {
	done := make(chan struct{}, 1)
	m.InitWithBytesNoCopyShapeDataTypeStridesMutableShapedBufferProvider(copy_, shape, type_, strides, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}


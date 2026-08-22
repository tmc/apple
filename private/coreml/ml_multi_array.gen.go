// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[MLMultiArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLMultiArray.BackingPixelBufferWasLocked]
//   - [MLMultiArray.Bytes]
//   - [MLMultiArray.CopyIntoMultiArrayError]
//   - [MLMultiArray.DebugQuickLookObject]
//   - [MLMultiArray.DoublePointer]
//   - [MLMultiArray.FillWithNumber]
//   - [MLMultiArray.Float32Pointer]
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
//   - [MLMultiArray.OffsetForKeyedSubscript]
//   - [MLMultiArray.RenderTo32BGRAPixelBufferChannelOrderIsBGRError]
//   - [MLMultiArray.RenderToCVPixelBufferChannelOrderIsBGRError]
//   - [MLMultiArray.RenderToOneComponent16HalfPixelBufferError]
//   - [MLMultiArray.RenderToOneComponent8PixelBufferError]
//   - [MLMultiArray.SetNumberAtOffset]
//   - [MLMultiArray.SetRangeWithRawDataDestIndexError]
//   - [MLMultiArray.SliceAtOriginShapeSqueezeError]
//   - [MLMultiArray.Squeeze]
//   - [MLMultiArray.SqueezeDimensionsError]
//   - [MLMultiArray.VectorizeIntoMultiArrayStorageOrderError]
//   - [MLMultiArray.InitWithArrayDataType]
//   - [MLMultiArray.InitWithBytesNoCopyShapeDataTypeStridesDeallocatorMutableShapedBufferProviderError]
//   - [MLMultiArray.InitWithBytesNoCopyShapeDataTypeStridesMutableShapedBufferProvider]
//   - [MLMultiArray.InitWithPixelBufferShapeStrides]
//   - [MLMultiArray.InitWithPixelBufferShapeStridesMutableShapedBufferProvider]
//   - [MLMultiArray.InitWithScalarsShapeDataType]
//   - [MLMultiArray.InitWithShapeDataTypeStorageOrderBufferAlignment]
//   - [MLMultiArray.InitWithShapeDataTypeStorageOrderError]
//   - [MLMultiArray.Contiguous]
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
//   - [IMLMultiArray.DebugQuickLookObject]
//   - [IMLMultiArray.DoublePointer]
//   - [IMLMultiArray.FillWithNumber]
//   - [IMLMultiArray.Float32Pointer]
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
//   - [IMLMultiArray.OffsetForKeyedSubscript]
//   - [IMLMultiArray.RenderTo32BGRAPixelBufferChannelOrderIsBGRError]
//   - [IMLMultiArray.RenderToCVPixelBufferChannelOrderIsBGRError]
//   - [IMLMultiArray.RenderToOneComponent16HalfPixelBufferError]
//   - [IMLMultiArray.RenderToOneComponent8PixelBufferError]
//   - [IMLMultiArray.SetNumberAtOffset]
//   - [IMLMultiArray.SetRangeWithRawDataDestIndexError]
//   - [IMLMultiArray.SliceAtOriginShapeSqueezeError]
//   - [IMLMultiArray.Squeeze]
//   - [IMLMultiArray.SqueezeDimensionsError]
//   - [IMLMultiArray.VectorizeIntoMultiArrayStorageOrderError]
//   - [IMLMultiArray.InitWithArrayDataType]
//   - [IMLMultiArray.InitWithBytesNoCopyShapeDataTypeStridesDeallocatorMutableShapedBufferProviderError]
//   - [IMLMultiArray.InitWithBytesNoCopyShapeDataTypeStridesMutableShapedBufferProvider]
//   - [IMLMultiArray.InitWithPixelBufferShapeStrides]
//   - [IMLMultiArray.InitWithPixelBufferShapeStridesMutableShapedBufferProvider]
//   - [IMLMultiArray.InitWithScalarsShapeDataType]
//   - [IMLMultiArray.InitWithShapeDataTypeStorageOrderBufferAlignment]
//   - [IMLMultiArray.InitWithShapeDataTypeStorageOrderError]
//   - [IMLMultiArray.Contiguous]
type IMLMultiArray interface {
	objectivec.IObject

	// Topic: Methods

	BackingPixelBufferWasLocked() bool
	Bytes() objectivec.IObject
	CopyIntoMultiArrayError(array objectivec.IObject) (bool, error)
	DebugQuickLookObject() objectivec.IObject
	DoublePointer() unsafe.Pointer
	FillWithNumber(number objectivec.IObject) bool
	Float32Pointer() unsafe.Pointer
	GetContiguousFirstMajorFloat32BufferWithHandler(handler VoidHandler)
	IsContiguous() bool
	IsContiguousInOrder(order int64) bool
	IsEqualToMultiArray(array objectivec.IObject) bool
	MtlBuffer() unsafe.Pointer
	MultiArrayBuffer() unsafe.Pointer
	MultiArrayViewExpandingDimensionsAtAxis(axis int64) objectivec.IObject
	MutableBytes() unsafe.Pointer
	NumberArray() objectivec.IObject
	NumberAtOffset(offset uint64) objectivec.IObject
	NumberOfBytesPerElement() uint64
	OffsetForKeyedSubscript(subscript objectivec.IObject) uint64
	RenderTo32BGRAPixelBufferChannelOrderIsBGRError(buffer corevideo.CVImageBufferRef, bgr bool) (bool, error)
	RenderToCVPixelBufferChannelOrderIsBGRError(buffer corevideo.CVImageBufferRef, bgr bool) (bool, error)
	RenderToOneComponent16HalfPixelBufferError(buffer corevideo.CVImageBufferRef) (bool, error)
	RenderToOneComponent8PixelBufferError(buffer corevideo.CVImageBufferRef) (bool, error)
	SetNumberAtOffset(number objectivec.IObject, offset uint64)
	SetRangeWithRawDataDestIndexError(data objectivec.IObject, index uint64) (bool, error)
	SliceAtOriginShapeSqueezeError(origin objectivec.IObject, shape objectivec.IObject, squeeze bool) (objectivec.IObject, error)
	Squeeze() objectivec.IObject
	SqueezeDimensionsError(dimensions objectivec.IObject) (objectivec.IObject, error)
	VectorizeIntoMultiArrayStorageOrderError(array objectivec.IObject, order int64) (bool, error)
	InitWithArrayDataType(array objectivec.IObject, type_ int64) MLMultiArray
	InitWithBytesNoCopyShapeDataTypeStridesDeallocatorMutableShapedBufferProviderError(copy_ unsafe.Pointer, shape objectivec.IObject, type_ int64, strides objectivec.IObject, deallocator func(), provider func()) (MLMultiArray, error)
	InitWithBytesNoCopyShapeDataTypeStridesMutableShapedBufferProvider(copy_ unsafe.Pointer, shape objectivec.IObject, type_ int64, strides objectivec.IObject, provider VoidHandler) MLMultiArray
	InitWithPixelBufferShapeStrides(buffer corevideo.CVImageBufferRef, shape objectivec.IObject, strides objectivec.IObject) MLMultiArray
	InitWithPixelBufferShapeStridesMutableShapedBufferProvider(buffer corevideo.CVImageBufferRef, shape objectivec.IObject, strides objectivec.IObject, provider VoidHandler) MLMultiArray
	InitWithScalarsShapeDataType(scalars objectivec.IObject, shape objectivec.IObject, type_ int64) MLMultiArray
	InitWithShapeDataTypeStorageOrderBufferAlignment(shape objectivec.IObject, type_ int64, order int64, alignment uint64) MLMultiArray
	InitWithShapeDataTypeStorageOrderError(shape objectivec.IObject, type_ int64, order int64) (MLMultiArray, error)
	Contiguous() bool
}

// Init initializes the instance.
func (m MLMultiArray) Init() MLMultiArray {
	rv := objc.SendIfResponds[MLMultiArray](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLMultiArray) Autorelease() MLMultiArray {
	rv := objc.SendIfResponds[MLMultiArray](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLMultiArray creates a new MLMultiArray instance.
func NewMLMultiArray() MLMultiArray {
	class := getMLMultiArrayClass()
	rv := objc.SendIfResponds[MLMultiArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewMultiArrayWithArrayDataType(array objectivec.IObject, type_ int64) MLMultiArray {
	instance := getMLMultiArrayClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithArray:dataType:"), array, type_)
	return MLMultiArrayFromID(rv)
}

func NewMultiArrayWithMultiArrayBuffer(buffer unsafe.Pointer) MLMultiArray {
	instance := getMLMultiArrayClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithMultiArrayBuffer:"), buffer)
	return MLMultiArrayFromID(rv)
}

func NewMultiArrayWithPixelBufferShapeStrides(buffer corevideo.CVImageBufferRef, shape objectivec.IObject, strides objectivec.IObject) MLMultiArray {
	instance := getMLMultiArrayClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithPixelBuffer:shape:strides:"), buffer, shape, strides)
	return MLMultiArrayFromID(rv)
}

func NewMultiArrayWithScalarsShapeDataType(scalars objectivec.IObject, shape objectivec.IObject, type_ int64) MLMultiArray {
	instance := getMLMultiArrayClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithScalars:shape:dataType:"), scalars, shape, type_)
	return MLMultiArrayFromID(rv)
}

func NewMultiArrayWithShapeDataTypeStorageOrderBufferAlignment(shape objectivec.IObject, type_ int64, order int64, alignment uint64) MLMultiArray {
	instance := getMLMultiArrayClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithShape:dataType:storageOrder:bufferAlignment:"), shape, type_, order, alignment)
	return MLMultiArrayFromID(rv)
}

func NewMultiArrayWithShapeDataTypeStorageOrderError(shape objectivec.IObject, type_ int64, order int64) (MLMultiArray, error) {
	var errorPtr objc.ID
	instance := getMLMultiArrayClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithShape:dataType:storageOrder:error:"), shape, type_, order, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLMultiArray{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLMultiArray{}, objc.ErrInitFailed
	}
	return MLMultiArrayFromID(rv), nil
}

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
func (m MLMultiArray) DebugQuickLookObject() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugQuickLookObject"))
	return objectivec.Object{ID: rv}
}
func (m MLMultiArray) DoublePointer() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("doublePointer"))
	return rv
}
func (m MLMultiArray) FillWithNumber(number objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("fillWithNumber:"), number)
	return rv
}
func (m MLMultiArray) Float32Pointer() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("float32Pointer"))
	return rv
}
func (m MLMultiArray) GetContiguousFirstMajorFloat32BufferWithHandler(handler VoidHandler) {
	_block0, _ := NewVoidBlock(handler)
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("getContiguousFirstMajorFloat32BufferWithHandler:"), _block0)
}
func (m MLMultiArray) IsContiguous() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("isContiguous"))
	return rv
}
func (m MLMultiArray) IsContiguousInOrder(order int64) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("isContiguousInOrder:"), order)
	return rv
}
func (m MLMultiArray) IsEqualToMultiArray(array objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("isEqualToMultiArray:"), array)
	return rv
}
func (m MLMultiArray) MultiArrayBuffer() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("multiArrayBuffer"))
	return rv
}
func (m MLMultiArray) MultiArrayViewExpandingDimensionsAtAxis(axis int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("multiArrayViewExpandingDimensionsAtAxis:"), axis)
	return objectivec.Object{ID: rv}
}
func (m MLMultiArray) NumberArray() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("numberArray"))
	return objectivec.Object{ID: rv}
}
func (m MLMultiArray) NumberAtOffset(offset uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("numberAtOffset:"), offset)
	return objectivec.Object{ID: rv}
}
func (m MLMultiArray) OffsetForKeyedSubscript(subscript objectivec.IObject) uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("offsetForKeyedSubscript:"), subscript)
	return rv
}
func (m MLMultiArray) RenderTo32BGRAPixelBufferChannelOrderIsBGRError(buffer corevideo.CVImageBufferRef, bgr bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("renderTo32BGRAPixelBuffer:channelOrderIsBGR:error:"), buffer, bgr, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("renderTo32BGRAPixelBuffer:channelOrderIsBGR:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLMultiArray) RenderToCVPixelBufferChannelOrderIsBGRError(buffer corevideo.CVImageBufferRef, bgr bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("renderToCVPixelBuffer:channelOrderIsBGR:error:"), buffer, bgr, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("renderToCVPixelBuffer:channelOrderIsBGR:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLMultiArray) RenderToOneComponent16HalfPixelBufferError(buffer corevideo.CVImageBufferRef) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("renderToOneComponent16HalfPixelBuffer:error:"), buffer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("renderToOneComponent16HalfPixelBuffer:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLMultiArray) RenderToOneComponent8PixelBufferError(buffer corevideo.CVImageBufferRef) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("renderToOneComponent8PixelBuffer:error:"), buffer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("renderToOneComponent8PixelBuffer:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLMultiArray) SetNumberAtOffset(number objectivec.IObject, offset uint64) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setNumber:atOffset:"), number, offset)
}
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
func (m MLMultiArray) SliceAtOriginShapeSqueezeError(origin objectivec.IObject, shape objectivec.IObject, squeeze bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("sliceAtOrigin:shape:squeeze:error:"), origin, shape, squeeze, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLMultiArray) Squeeze() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("squeeze"))
	return objectivec.Object{ID: rv}
}
func (m MLMultiArray) SqueezeDimensionsError(dimensions objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("squeezeDimensions:error:"), dimensions, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
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
func (m MLMultiArray) InitWithArrayDataType(array objectivec.IObject, type_ int64) MLMultiArray {
	rv := objc.SendIfResponds[MLMultiArray](m.ID, objc.Sel("initWithArray:dataType:"), array, type_)
	return rv
}

var _mlmultiarray_initwithbytesnocopy_shape_datatype_strides_deallocator_mutableshapedbufferprovider_error_p4_key byte
var _mlmultiarray_initwithbytesnocopy_shape_datatype_strides_deallocator_mutableshapedbufferprovider_error_p5_key byte

func (m MLMultiArray) InitWithBytesNoCopyShapeDataTypeStridesDeallocatorMutableShapedBufferProviderError(copy_ unsafe.Pointer, shape objectivec.IObject, type_ int64, strides objectivec.IObject, deallocator func(), provider func()) (MLMultiArray, error) {
	_block4, _cleanup4 := NewVoidBlock(deallocator)
	defer _cleanup4()
	_block5, _cleanup5 := NewVoidBlock(provider)
	defer _cleanup5()
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithBytesNoCopy:shape:dataType:strides:deallocator:mutableShapedBufferProvider:error:"), copy_, shape, type_, strides, objc.ID(_block4), objc.ID(_block5), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLMultiArray{}, foundation.NSErrorFrom(errorPtr)
	}
	objc.AssociateBlockWithReceiver(rv, &_mlmultiarray_initwithbytesnocopy_shape_datatype_strides_deallocator_mutableshapedbufferprovider_error_p4_key, objc.Block(_block4))
	objc.AssociateBlockWithReceiver(rv, &_mlmultiarray_initwithbytesnocopy_shape_datatype_strides_deallocator_mutableshapedbufferprovider_error_p5_key, objc.Block(_block5))
	return MLMultiArrayFromID(rv), nil

}
func (m MLMultiArray) InitWithBytesNoCopyShapeDataTypeStridesMutableShapedBufferProvider(copy_ unsafe.Pointer, shape objectivec.IObject, type_ int64, strides objectivec.IObject, provider VoidHandler) MLMultiArray {
	_block4, _ := NewVoidBlock(provider)
	rv := objc.SendIfResponds[MLMultiArray](m.ID, objc.Sel("initWithBytesNoCopy:shape:dataType:strides:mutableShapedBufferProvider:"), copy_, shape, type_, strides, _block4)
	return rv
}
func (m MLMultiArray) InitWithPixelBufferShapeStrides(buffer corevideo.CVImageBufferRef, shape objectivec.IObject, strides objectivec.IObject) MLMultiArray {
	rv := objc.SendIfResponds[MLMultiArray](m.ID, objc.Sel("initWithPixelBuffer:shape:strides:"), buffer, shape, strides)
	return rv
}
func (m MLMultiArray) InitWithPixelBufferShapeStridesMutableShapedBufferProvider(buffer corevideo.CVImageBufferRef, shape objectivec.IObject, strides objectivec.IObject, provider VoidHandler) MLMultiArray {
	_block3, _ := NewVoidBlock(provider)
	rv := objc.SendIfResponds[MLMultiArray](m.ID, objc.Sel("initWithPixelBuffer:shape:strides:mutableShapedBufferProvider:"), buffer, shape, strides, _block3)
	return rv
}
func (m MLMultiArray) InitWithScalarsShapeDataType(scalars objectivec.IObject, shape objectivec.IObject, type_ int64) MLMultiArray {
	rv := objc.SendIfResponds[MLMultiArray](m.ID, objc.Sel("initWithScalars:shape:dataType:"), scalars, shape, type_)
	return rv
}
func (m MLMultiArray) InitWithShapeDataTypeStorageOrderBufferAlignment(shape objectivec.IObject, type_ int64, order int64, alignment uint64) MLMultiArray {
	rv := objc.SendIfResponds[MLMultiArray](m.ID, objc.Sel("initWithShape:dataType:storageOrder:bufferAlignment:"), shape, type_, order, alignment)
	return rv
}
func (m MLMultiArray) InitWithShapeDataTypeStorageOrderError(shape objectivec.IObject, type_ int64, order int64) (MLMultiArray, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithShape:dataType:storageOrder:error:"), shape, type_, order, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLMultiArray{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLMultiArrayFromID(rv), nil

}

func (_MLMultiArrayClass MLMultiArrayClass) _multiArrayByConcatenatingMultiArraysAlongAxisDataType(arrays objectivec.IObject, axis int64, type_ int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("_multiArrayByConcatenatingMultiArrays:alongAxis:dataType:"), arrays, axis, type_)
	return objectivec.Object{ID: rv}
}
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
	if !objc.RespondsToSelector(objc.ID(_MLMultiArrayClass.class), objc.Sel("_shapeOfNestedArray:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_shapeOfNestedArray:error:"}
		return nil, err
	}
	return _MLMultiArrayClass._shapeOfNestedArrayError(array)
}

// CanShapeOfNestedArrayError reports whether the receiver responds to the private selector _shapeOfNestedArray:error:.
func (_MLMultiArrayClass MLMultiArrayClass) CanShapeOfNestedArrayError() bool {
	return objc.RespondsToSelector(objc.ID(_MLMultiArrayClass.class), objc.Sel("_shapeOfNestedArray:error:"))
}
func (_MLMultiArrayClass MLMultiArrayClass) CppStorageOrder(order int64) int {
	rv := objc.SendIfResponds[int](objc.ID(_MLMultiArrayClass.class), objc.Sel("cppStorageOrder:"), order)
	return rv
}
func (_MLMultiArrayClass MLMultiArrayClass) DoubleMatrixWithValuesError(values objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("doubleMatrixWithValues:error:"), values, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLMultiArrayClass MLMultiArrayClass) DoubleMultiArrayWithCopyOfMultiArray(array objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("doubleMultiArrayWithCopyOfMultiArray:"), array)
	return objectivec.Object{ID: rv}
}
func (_MLMultiArrayClass MLMultiArrayClass) DoubleMultiArrayWithShapeValueArrayError(shape objectivec.IObject, array objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("doubleMultiArrayWithShape:valueArray:error:"), shape, array, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLMultiArrayClass MLMultiArrayClass) DoubleVectorWithValues(values objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("doubleVectorWithValues:"), values)
	return objectivec.Object{ID: rv}
}
func (_MLMultiArrayClass MLMultiArrayClass) Float32MatrixWithValuesError(values objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("float32MatrixWithValues:error:"), values, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
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
func (_MLMultiArrayClass MLMultiArrayClass) MultiArrayByConcatenatingMultiArraysAlongAxisDataType(arrays objectivec.IObject, axis int64, type_ int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("multiArrayByConcatenatingMultiArrays:alongAxis:dataType:"), arrays, axis, type_)
	return objectivec.Object{ID: rv}
}
func (_MLMultiArrayClass MLMultiArrayClass) MultiArrayOwningBufferObjectOfPortError(port E5rtIOPortRef) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("multiArrayOwningBufferObjectOfPort:error:"), port, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLMultiArrayClass MLMultiArrayClass) NewMultiArrayForStateStridesDtypeError(state objectivec.IObject, strides objectivec.IObject, dtype int64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("newMultiArrayForState:strides:dtype:error:"), state, strides, dtype, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLMultiArrayClass MLMultiArrayClass) PixelBufferBGRA8FromMultiArrayCHWChannelOrderIsBGRError(chw objectivec.IObject, bgr bool) (corevideo.CVImageBufferRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[corevideo.CVImageBufferRef](objc.ID(_MLMultiArrayClass.class), objc.Sel("pixelBufferBGRA8FromMultiArrayCHW:channelOrderIsBGR:error:"), chw, bgr, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(corevideo.CVImageBufferRef), foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}
func (_MLMultiArrayClass MLMultiArrayClass) PixelBufferGray16HalfFromMultiArrayHWError(hw objectivec.IObject) (corevideo.CVImageBufferRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[corevideo.CVImageBufferRef](objc.ID(_MLMultiArrayClass.class), objc.Sel("pixelBufferGray16HalfFromMultiArrayHW:error:"), hw, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(corevideo.CVImageBufferRef), foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}
func (_MLMultiArrayClass MLMultiArrayClass) PixelBufferGray8FromMultiArrayHWError(hw objectivec.IObject) (corevideo.CVImageBufferRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[corevideo.CVImageBufferRef](objc.ID(_MLMultiArrayClass.class), objc.Sel("pixelBufferGray8FromMultiArrayHW:error:"), hw, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(corevideo.CVImageBufferRef), foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}
func (_MLMultiArrayClass MLMultiArrayClass) StringForDataType(type_ int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLMultiArrayClass.class), objc.Sel("stringForDataType:"), type_)
	return objectivec.Object{ID: rv}
}
func (_MLMultiArrayClass MLMultiArrayClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLMultiArrayClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
func (_MLMultiArrayClass MLMultiArrayClass) ValidateMultiArraysForConcatenatingAlongAxisNormalizedAxisReason(arrays objectivec.IObject, axis int64, reason []objectivec.IObject) (uint64, bool) {
	var axis2 uint64
	rv := objc.Send[bool](objc.ID(_MLMultiArrayClass.class), objc.Sel("validateMultiArrays:forConcatenatingAlongAxis:normalizedAxis:reason:"), arrays, axis, unsafe.Pointer(&axis2), reason)
	return axis2, rv
}
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

func (m MLMultiArray) BackingPixelBufferWasLocked() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("backingPixelBufferWasLocked"))
	return rv
}
func (m MLMultiArray) Bytes() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("bytes"))
	return objectivec.Object{ID: rv}
}
func (m MLMultiArray) Contiguous() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("contiguous"))
	return rv
}
func (m MLMultiArray) MtlBuffer() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("mtlBuffer"))
	return rv
}
func (m MLMultiArray) MutableBytes() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("mutableBytes"))
	return rv
}
func (m MLMultiArray) NumberOfBytesPerElement() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("numberOfBytesPerElement"))
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

// InitWithPixelBufferShapeStridesMutableShapedBufferProviderSync is a synchronous wrapper around [MLMultiArray.InitWithPixelBufferShapeStridesMutableShapedBufferProvider].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLMultiArray) InitWithPixelBufferShapeStridesMutableShapedBufferProviderSync(ctx context.Context, buffer corevideo.CVImageBufferRef, shape objectivec.IObject, strides objectivec.IObject) error {
	done := make(chan struct{}, 1)
	m.InitWithPixelBufferShapeStridesMutableShapedBufferProvider(buffer, shape, strides, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

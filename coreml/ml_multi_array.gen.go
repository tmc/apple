// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
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
	rv := objc.Send[MLMultiArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A machine learning collection type that stores numeric values in an array
// with multiple dimensions.
//
// # Overview
//
// A multidimensional array, or , is one of the underlying types of an
// [MLFeatureValue] that stores numeric values in multiple dimensions. All
// elements in an [MLMultiArray] instance are one of the same type, and one of
// the types that [MLMultiArrayDataType] defines:
//
// [MLMultiArrayDataTypeInt32]: 32-bit integer [MLMultiArrayDataTypeFloat16]:
// 16-bit floating point number [MLMultiArrayDataTypeFloat32]: 32-bit floating
// point number (also known as [MLMultiArrayDataTypeFloat])
// [MLMultiArrayDataTypeFloat64]: 64-bit floating point number (also known as
// `double` in Swift or [MLMultiArrayDataTypeDouble] in Objective-C)
//
// Each dimension in a multiarray is typically significant or meaningful. For
// example, a model could have an input that accepts images as a multiarray of
// pixels with three dimensions, C x H x W. The first dimension, ,_
// _represents the number of color channels, and the second and third
// dimensions, and , represent the image’s height and width, respectively.
// The number of dimensions and size of each dimension define the
// multiarray’s .
//
// The [MLMultiArray.Shape] property is an integer array that has an element for each
// dimension in the multiarray. Each element in [MLMultiArray.Shape] defines the size of
// the corresponding dimension. To inspect the shape and constraints of a
// model’s multiarray input or output feature:
//
// - Access the model’s [ModelDescription] property. - Find the multiarray
// input or output feature in the model description’s
// [MLMultiArray.InputDescriptionsByName] or [MLMultiArray.OutputDescriptionsByName] property,
// respectively. - Access the feature description’s [MultiArrayConstraint]
// property. - Inspect the multiarray constraint’s [MLMultiArray.Shape] and
// [MLMultiArray.ShapeConstraint].
//
// # Creating a multiarray
//
//   - [MLMultiArray.InitWithShapeDataTypeError]: Creates a multidimensional array with a shape and type.
//   - [MLMultiArray.InitWithDataPointerShapeDataTypeStridesDeallocatorError]: Creates a multiarray from a data pointer.
//   - [MLMultiArray.InitWithPixelBufferShape]: Creates a multiarray sharing the surface of a pixel buffer.
//
// # Inspecting a multiarray
//
//   - [MLMultiArray.Count]: The total number of elements in the multiarray.
//   - [MLMultiArray.DataType]: The underlying type of the multiarray.
//   - [MLMultiArray.Shape]: The multiarray’s multidimensional shape as a number array in which each element’s value is the size of the corresponding dimension.
//   - [MLMultiArray.Strides]: A number array in which each element is the number of memory locations that span the length of the corresponding dimension.
//
// # Transfering the contents
//
//   - [MLMultiArray.TransferToMultiArray]: Transfer the contents to the destination multi-array.
//
// # Accessing a multiarray’s elements
//
//   - [MLMultiArray.PixelBuffer]: A reference to the multiarray’s underlying pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray
//
// [MLMultiArrayDataType]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType
type MLMultiArray struct {
	objectivec.Object
}

// MLMultiArrayFromID constructs a [MLMultiArray] from an objc.ID.
//
// A machine learning collection type that stores numeric values in an array
// with multiple dimensions.
func MLMultiArrayFromID(id objc.ID) MLMultiArray {
	return MLMultiArray{objectivec.Object{ID: id}}
}

// NOTE: MLMultiArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLMultiArray] class.
//
// # Creating a multiarray
//
//   - [IMLMultiArray.InitWithShapeDataTypeError]: Creates a multidimensional array with a shape and type.
//   - [IMLMultiArray.InitWithDataPointerShapeDataTypeStridesDeallocatorError]: Creates a multiarray from a data pointer.
//   - [IMLMultiArray.InitWithPixelBufferShape]: Creates a multiarray sharing the surface of a pixel buffer.
//
// # Inspecting a multiarray
//
//   - [IMLMultiArray.Count]: The total number of elements in the multiarray.
//   - [IMLMultiArray.DataType]: The underlying type of the multiarray.
//   - [IMLMultiArray.Shape]: The multiarray’s multidimensional shape as a number array in which each element’s value is the size of the corresponding dimension.
//   - [IMLMultiArray.Strides]: A number array in which each element is the number of memory locations that span the length of the corresponding dimension.
//
// # Transfering the contents
//
//   - [IMLMultiArray.TransferToMultiArray]: Transfer the contents to the destination multi-array.
//
// # Accessing a multiarray’s elements
//
//   - [IMLMultiArray.PixelBuffer]: A reference to the multiarray’s underlying pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray
type IMLMultiArray interface {
	objectivec.IObject

	// Topic: Creating a multiarray

	// Creates a multidimensional array with a shape and type.
	InitWithShapeDataTypeError(shape []foundation.NSNumber, dataType MLMultiArrayDataType) (MLMultiArray, error)
	// Creates a multiarray from a data pointer.
	InitWithDataPointerShapeDataTypeStridesDeallocatorError(dataPointer unsafe.Pointer, shape []foundation.NSNumber, dataType MLMultiArrayDataType, strides []foundation.NSNumber, deallocator func(unsafe.Pointer)) (MLMultiArray, error)
	// Creates a multiarray sharing the surface of a pixel buffer.
	InitWithPixelBufferShape(pixelBuffer corevideo.CVImageBufferRef, shape []foundation.NSNumber) MLMultiArray

	// Topic: Inspecting a multiarray

	// The total number of elements in the multiarray.
	Count() int
	// The underlying type of the multiarray.
	DataType() MLMultiArrayDataType
	// The multiarray’s multidimensional shape as a number array in which each element’s value is the size of the corresponding dimension.
	Shape() []foundation.NSNumber
	// A number array in which each element is the number of memory locations that span the length of the corresponding dimension.
	Strides() []foundation.NSNumber

	// Topic: Transfering the contents

	// Transfer the contents to the destination multi-array.
	TransferToMultiArray(destinationMultiArray IMLMultiArray)

	// Topic: Accessing a multiarray’s elements

	// A reference to the multiarray’s underlying pixel buffer.
	PixelBuffer() corevideo.CVImageBufferRef

	// A dictionary of input feature descriptions, which the model keys by the input’s name.
	InputDescriptionsByName() IMLFeatureDescription
	SetInputDescriptionsByName(value IMLFeatureDescription)
	// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
	ModelDescription() IMLModelDescription
	SetModelDescription(value IMLModelDescription)
	// The constraints on a multidimensional array feature.
	MultiArrayConstraint() IMLMultiArrayConstraint
	SetMultiArrayConstraint(value IMLMultiArrayConstraint)
	// A dictionary of output feature descriptions, which the model keys by the output’s name.
	OutputDescriptionsByName() IMLFeatureDescription
	SetOutputDescriptionsByName(value IMLFeatureDescription)
	// The constraint on the shape of the multiarray.
	ShapeConstraint() IMLMultiArrayShapeConstraint
	SetShapeConstraint(value IMLMultiArrayShapeConstraint)
	// Get the underlying buffer pointer to read.
	GetBytesWithHandler(handler func(unsafe.Pointer, int64))
	// Creates the object with specified strides.
	InitWithShapeDataTypeStrides(shape []foundation.NSNumber, dataType MLMultiArrayDataType, strides []foundation.NSNumber) MLMultiArray
	// Accesses the multiarray by using a linear offset.
	ObjectAtIndexedSubscript(idx int) foundation.NSNumber
	// Accesses the multiarray by using a number array that has an element for each dimension.
	ObjectForKeyedSubscript(key []foundation.NSNumber) foundation.NSNumber
	// Assigns a number to the multiarray’s element at the location that the linear offset defines.
	SetObjectAtIndexedSubscript(obj foundation.NSNumber, idx int)
	// Assigns a number to the multiarray’s element at the location that the number array defines.
	SetObjectForKeyedSubscript(obj foundation.NSNumber, key []foundation.NSNumber)
	EncodeWithCoder(coder foundation.INSCoder)
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

// Merges an array of multiarrays into one multiarray along an axis.
//
// multiArrays: An [MLMultiArray] array.
//
// axis: A zero-based axis number the instances in `multiArray` merge along.
//
// dataType: An [MLMultiArrayDataType] instance that represents the underlying type of
// all the instances in `multiArrays`.
//
// # Discussion
//
// All multiarray instances in `multiArrays` must have:
//
// - The same data type - The same number of dimensions - The same size for
// each corresponding dimension, except for the concatenation axis
//
// For example, this code concatenates two multiarrays along their first
// dimension:
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(byConcatenatingMultiArrays:alongAxis:dataType:)
//
// [MLMultiArrayDataType]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType
func NewMultiArrayByConcatenatingMultiArraysAlongAxisDataType(multiArrays []MLMultiArray, axis int, dataType MLMultiArrayDataType) MLMultiArray {
	rv := objc.Send[objc.ID](objc.ID(getMLMultiArrayClass().class), objc.Sel("multiArrayByConcatenatingMultiArrays:alongAxis:dataType:"), objectivec.IObjectSliceToNSArray(multiArrays), axis, dataType)
	return MLMultiArrayFromID(rv)
}

// Creates a multiarray sharing the surface of a pixel buffer.
//
// pixelBuffer: The pixel buffer owned by the instance.
//
// shape: The shape of the [MLMultiArray]. The last dimension of `shape` must match
// the pixel buffer’s width. The product of the rest of the dimensions must
// match the height.
//
// # Discussion
//
// Use this initializer to create an [IOSurface]-backed [MLMultiArray] that
// reduces the inference latency by avoiding the buffer copy to and from some
// compute units.
//
// The instance will own the pixel buffer and release it on the deallocation.
//
// The pixel buffer’s pixel format type must be
// [kCVPixelFormatType_OneComponent16Half]. The [MLMultiArray] data type is
// [MLMultiArrayDataTypeFloat16].
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(pixelBuffer:shape:)
//
// [IOSurface]: https://developer.apple.com/documentation/IOSurface
// [kCVPixelFormatType_OneComponent16Half]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_OneComponent16Half
func NewMultiArrayWithPixelBufferShape(pixelBuffer corevideo.CVImageBufferRef, shape []foundation.NSNumber) MLMultiArray {
	instance := getMLMultiArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPixelBuffer:shape:"), pixelBuffer, objectivec.IObjectSliceToNSArray(shape))
	return MLMultiArrayFromID(rv)
}

// Creates a multidimensional array with a shape and type.
//
// shape: An integer array that has an element for each dimension in a multiarray
// that represents its length.
//
// dataType: An element type defined by [MLMultiArrayDataType].
//
// # Discussion
//
// This method allocates a contiguous region of memory for the multiarray’s
// shape. You must set the contents of memory. The multiarray frees the memory
// in its deinitializer (Swift) or [dealloc] method (Objective-C).
//
// The following code creates a 3 x 3 multiarray and sets its contents to the
// value 3.14159.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(shape:dataType:)
//
// [MLMultiArrayDataType]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType
// [dealloc]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/dealloc
func NewMultiArrayWithShapeDataTypeError(shape []foundation.NSNumber, dataType MLMultiArrayDataType) (MLMultiArray, error) {
	var errorPtr objc.ID
	instance := getMLMultiArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithShape:dataType:error:"), objectivec.IObjectSliceToNSArray(shape), dataType, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLMultiArray{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLMultiArrayFromID(rv), nil
}

// Creates the object with specified strides.
//
// shape: The shape
//
// dataType: The data type
//
// strides: The strides.
//
// # Discussion
//
// The contents of the object are left uninitialized; the client must
// initialize it.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithShape:dataType:strides:
func NewMultiArrayWithShapeDataTypeStrides(shape []foundation.NSNumber, dataType MLMultiArrayDataType, strides []foundation.NSNumber) MLMultiArray {
	instance := getMLMultiArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithShape:dataType:strides:"), objectivec.IObjectSliceToNSArray(shape), dataType, objectivec.IObjectSliceToNSArray(strides))
	return MLMultiArrayFromID(rv)
}

// Creates a multidimensional array with a shape and type.
//
// shape: An integer array that has an element for each dimension in a multiarray
// that represents its length.
//
// dataType: An element type defined by [MLMultiArrayDataType].
//
// # Discussion
//
// This method allocates a contiguous region of memory for the multiarray’s
// shape. You must set the contents of memory. The multiarray frees the memory
// in its deinitializer (Swift) or [dealloc] method (Objective-C).
//
// The following code creates a 3 x 3 multiarray and sets its contents to the
// value 3.14159.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(shape:dataType:)
//
// [MLMultiArrayDataType]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType
// [dealloc]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/dealloc
func (m MLMultiArray) InitWithShapeDataTypeError(shape []foundation.NSNumber, dataType MLMultiArrayDataType) (MLMultiArray, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithShape:dataType:error:"), objectivec.IObjectSliceToNSArray(shape), dataType, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLMultiArray{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLMultiArrayFromID(rv), nil

}

// Creates a multiarray from a data pointer.
//
// dataPointer: A pointer to data in memory.
//
// shape: An integer array with an element for each dimension. An element represents
// the size of the corresponding dimension.
//
// dataType: An [MLMultiArrayDataType] instance that represents the pointer’s data
// type.
//
// strides: An integer array with an element for each dimension. An element represents
// the number of memory locations that span the length of the corresponding
// dimension.
//
// deallocator: In Swift, a closure the multiarray calls in its deinitializer. In
// Objective-C, a block the multiarray calls in its [dealloc] method.
//
// # Discussion
//
// The caller is responsible for freeing the memory the `dataPointer` points
// to, by providing a `deallocator` closure.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(dataPointer:shape:dataType:strides:deallocator:)
//
// [MLMultiArrayDataType]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType
// [dealloc]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/dealloc
func (m MLMultiArray) InitWithDataPointerShapeDataTypeStridesDeallocatorError(dataPointer unsafe.Pointer, shape []foundation.NSNumber, dataType MLMultiArrayDataType, strides []foundation.NSNumber, deallocator func(unsafe.Pointer)) (MLMultiArray, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithDataPointer:shape:dataType:strides:deallocator:error:"), dataPointer, objectivec.IObjectSliceToNSArray(shape), dataType, objectivec.IObjectSliceToNSArray(strides), deallocator, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLMultiArray{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLMultiArrayFromID(rv), nil

}

// Creates a multiarray sharing the surface of a pixel buffer.
//
// pixelBuffer: The pixel buffer owned by the instance.
//
// shape: The shape of the [MLMultiArray]. The last dimension of `shape` must match
// the pixel buffer’s width. The product of the rest of the dimensions must
// match the height.
//
// # Discussion
//
// Use this initializer to create an [IOSurface]-backed [MLMultiArray] that
// reduces the inference latency by avoiding the buffer copy to and from some
// compute units.
//
// The instance will own the pixel buffer and release it on the deallocation.
//
// The pixel buffer’s pixel format type must be
// [kCVPixelFormatType_OneComponent16Half]. The [MLMultiArray] data type is
// [MLMultiArrayDataTypeFloat16].
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/init(pixelBuffer:shape:)
//
// [IOSurface]: https://developer.apple.com/documentation/IOSurface
// [kCVPixelFormatType_OneComponent16Half]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_OneComponent16Half
func (m MLMultiArray) InitWithPixelBufferShape(pixelBuffer corevideo.CVImageBufferRef, shape []foundation.NSNumber) MLMultiArray {
	rv := objc.Send[MLMultiArray](m.ID, objc.Sel("initWithPixelBuffer:shape:"), pixelBuffer, objectivec.IObjectSliceToNSArray(shape))
	return rv
}

// Transfer the contents to the destination multi-array.
//
// destinationMultiArray: The transfer destination.
//
// # Discussion
//
// Numeric data will be up or down casted as needed. It can transfer to a
// multi-array with different layout (strides).
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/transfer(to:)
func (m MLMultiArray) TransferToMultiArray(destinationMultiArray IMLMultiArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("transferToMultiArray:"), destinationMultiArray)
}

// Get the underlying buffer pointer to read.
//
// handler: The block to receive the buffer pointer and its size in bytes. This block
// has no return value and takes the following parameters:
//
// `bytes`: The pointer to the buffer.
// `size`: The size of the buffer.
//
// # Discussion
//
// The buffer contains a collection of `int32`, `float16`, `float32`, or
// `float64` values, depending on the multiarray’s data type. It may not
// store these scalar values contiguously; use [Strides] to get the buffer
// layout.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/getBytesWithHandler:
func (m MLMultiArray) GetBytesWithHandler(handler func(unsafe.Pointer, int64)) {
	_block0 := objc.NewBlock(func(_ objc.Block, arg0 unsafe.Pointer, arg1 int64) { handler(arg0, arg1) })
	defer _block0.Release()
	objc.Send[objc.ID](m.ID, objc.Sel("getBytesWithHandler:"), objc.ID(_block0))
}

// Creates the object with specified strides.
//
// shape: The shape
//
// dataType: The data type
//
// strides: The strides.
//
// # Discussion
//
// The contents of the object are left uninitialized; the client must
// initialize it.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/initWithShape:dataType:strides:
func (m MLMultiArray) InitWithShapeDataTypeStrides(shape []foundation.NSNumber, dataType MLMultiArrayDataType, strides []foundation.NSNumber) MLMultiArray {
	rv := objc.Send[MLMultiArray](m.ID, objc.Sel("initWithShape:dataType:strides:"), objectivec.IObjectSliceToNSArray(shape), dataType, objectivec.IObjectSliceToNSArray(strides))
	return rv
}

// Accesses the multiarray by using a linear offset.
//
// idx: A linear offset index that represents a position the multiarray.
//
// # Return Value
//
// A number.
//
// # Discussion
//
// Use the subscript to linearly index a one-dimensional multiarray in the
// same way as a conventional array.
//
// Multiarrays with more than one dimension order their elements in row-major
// order. In these cases, calculate a linear offset by summing the products of
// each dimension’s index with the dimension’s stride (See [Strides]).
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/subscript(_:)-2hh91
func (m MLMultiArray) ObjectAtIndexedSubscript(idx int) foundation.NSNumber {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectAtIndexedSubscript:"), idx)
	return foundation.NSNumberFromID(rv)
}

// Accesses the multiarray by using a number array that has an element for
// each dimension.
//
// key: An [NSNumber] array that represents a position in a multiarray in which
// each element is an index in the corresponding dimension.
//
// # Return Value
//
// A number.
//
// # Discussion
//
// Use this subscript to access the multiarray by using an [NSNumber] array in
// which each element is an index in the multiarray’s corresponding
// dimension.
//
// For better performance, directly access the multiarray’s underlying
// pointer by using a linear offset. Calculate the linear offset by summing
// the products of each dimension’s index with the dimension’s stride (See
// [Strides]).
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/subscript(_:)-3d9el
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (m MLMultiArray) ObjectForKeyedSubscript(key []foundation.NSNumber) foundation.NSNumber {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectForKeyedSubscript:"), objectivec.IObjectSliceToNSArray(key))
	return foundation.NSNumberFromID(rv)
}

// Assigns a number to the multiarray’s element at the location that the
// linear offset defines.
//
// obj: An [NSNumber] to assign to the element.
//
// idx: A linear offset index that represents a position the multiarray in the
// first major order.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/setObject:atIndexedSubscript:
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (m MLMultiArray) SetObjectAtIndexedSubscript(obj foundation.NSNumber, idx int) {
	objc.Send[objc.ID](m.ID, objc.Sel("setObject:atIndexedSubscript:"), obj, idx)
}

// Assigns a number to the multiarray’s element at the location that the
// number array defines.
//
// obj: An [NSNumber] to assign to the element.
//
// key: An [NSNumber] array that represents an element’s position in the
// multiarray. Each element of `key` is an index into the multiarray’s
// corresponding dimension.
//
// # Discussion
//
// See [ObjectForKeyedSubscript] for the method’s getter counterpart.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/setObject:forKeyedSubscript:
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (m MLMultiArray) SetObjectForKeyedSubscript(obj foundation.NSNumber, key []foundation.NSNumber) {
	objc.Send[objc.ID](m.ID, objc.Sel("setObject:forKeyedSubscript:"), obj, objectivec.IObjectSliceToNSArray(key))
}
func (m MLMultiArray) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The total number of elements in the multiarray.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/count
func (m MLMultiArray) Count() int {
	rv := objc.Send[int](m.ID, objc.Sel("count"))
	return rv
}

// The underlying type of the multiarray.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/dataType
func (m MLMultiArray) DataType() MLMultiArrayDataType {
	rv := objc.Send[MLMultiArrayDataType](m.ID, objc.Sel("dataType"))
	return MLMultiArrayDataType(rv)
}

// The multiarray’s multidimensional shape as a number array in which each
// element’s value is the size of the corresponding dimension.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/shape
func (m MLMultiArray) Shape() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("shape"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// A number array in which each element is the number of memory locations that
// span the length of the corresponding dimension.
//
// # Discussion
//
// See [ObjectAtIndexedSubscript] and [ObjectForKeyedSubscript] for code
// examples that use `strides`.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/strides
func (m MLMultiArray) Strides() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("strides"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// A reference to the multiarray’s underlying pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArray/pixelBuffer
func (m MLMultiArray) PixelBuffer() corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](m.ID, objc.Sel("pixelBuffer"))
	return corevideo.CVImageBufferRef(rv)
}

// A dictionary of input feature descriptions, which the model keys by the
// input’s name.
//
// See: https://developer.apple.com/documentation/coreml/mlmodeldescription/inputdescriptionsbyname
func (m MLMultiArray) InputDescriptionsByName() IMLFeatureDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputDescriptionsByName"))
	return MLFeatureDescriptionFromID(objc.ID(rv))
}
func (m MLMultiArray) SetInputDescriptionsByName(value IMLFeatureDescription) {
	objc.Send[struct{}](m.ID, objc.Sel("setInputDescriptionsByName:"), value)
}

// Model information you use at runtime during development, which Xcode also
// displays in its Core ML model editor view.
//
// See: https://developer.apple.com/documentation/coreml/mlmodel/modeldescription
func (m MLMultiArray) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}
func (m MLMultiArray) SetModelDescription(value IMLModelDescription) {
	objc.Send[struct{}](m.ID, objc.Sel("setModelDescription:"), value)
}

// The constraints on a multidimensional array feature.
//
// See: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint
func (m MLMultiArray) MultiArrayConstraint() IMLMultiArrayConstraint {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("multiArrayConstraint"))
	return MLMultiArrayConstraintFromID(objc.ID(rv))
}
func (m MLMultiArray) SetMultiArrayConstraint(value IMLMultiArrayConstraint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMultiArrayConstraint:"), value)
}

// A dictionary of output feature descriptions, which the model keys by the
// output’s name.
//
// See: https://developer.apple.com/documentation/coreml/mlmodeldescription/outputdescriptionsbyname
func (m MLMultiArray) OutputDescriptionsByName() IMLFeatureDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputDescriptionsByName"))
	return MLFeatureDescriptionFromID(objc.ID(rv))
}
func (m MLMultiArray) SetOutputDescriptionsByName(value IMLFeatureDescription) {
	objc.Send[struct{}](m.ID, objc.Sel("setOutputDescriptionsByName:"), value)
}

// The constraint on the shape of the multiarray.
//
// See: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/shapeconstraint
func (m MLMultiArray) ShapeConstraint() IMLMultiArrayShapeConstraint {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("shapeConstraint"))
	return MLMultiArrayShapeConstraintFromID(objc.ID(rv))
}
func (m MLMultiArray) SetShapeConstraint(value IMLMultiArrayShapeConstraint) {
	objc.Send[struct{}](m.ID, objc.Sel("setShapeConstraint:"), value)
}

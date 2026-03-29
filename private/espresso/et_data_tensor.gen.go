// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETDataTensor] class.
var (
	_ETDataTensorClass     ETDataTensorClass
	_ETDataTensorClassOnce sync.Once
)

func getETDataTensorClass() ETDataTensorClass {
	_ETDataTensorClassOnce.Do(func() {
		_ETDataTensorClass = ETDataTensorClass{class: objc.GetClass("ETDataTensor")}
	})
	return _ETDataTensorClass
}

// GetETDataTensorClass returns the class object for ETDataTensor.
func GetETDataTensorClass() ETDataTensorClass {
	return getETDataTensorClass()
}

type ETDataTensorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETDataTensorClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETDataTensorClass) Alloc() ETDataTensor {
	rv := objc.Send[ETDataTensor](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ETDataTensor.AllocatedImageData]
//   - [ETDataTensor.SetAllocatedImageData]
//   - [ETDataTensor.Blob]
//   - [ETDataTensor.SetBlob]
//   - [ETDataTensor.DataArray]
//   - [ETDataTensor.SetDataArray]
//   - [ETDataTensor.DataPointer]
//   - [ETDataTensor.SetDataPointer]
//   - [ETDataTensor.Float_buffer]
//   - [ETDataTensor.SetFloat_buffer]
//   - [ETDataTensor.ImageBuffer]
//   - [ETDataTensor.SetImageBuffer]
//   - [ETDataTensor.MaxNumberOfElements]
//   - [ETDataTensor.SetMaxNumberOfElements]
//   - [ETDataTensor.Shape]
//   - [ETDataTensor.SetShape]
//   - [ETDataTensor.Strides]
//   - [ETDataTensor.SetStrides]
//   - [ETDataTensor.Type]
//   - [ETDataTensor.SetType]
//   - [ETDataTensor.InitWithBlobContainer]
//   - [ETDataTensor.InitWithBlobContainerDirectBind]
//   - [ETDataTensor.InitWithCVPixelBufferImageParametersError]
//   - [ETDataTensor.InitWithDataTypeShapeStrides]
// See: https://developer.apple.com/documentation/Espresso/ETDataTensor
type ETDataTensor struct {
	objectivec.Object
}

// ETDataTensorFromID constructs a [ETDataTensor] from an objc.ID.
func ETDataTensorFromID(id objc.ID) ETDataTensor {
	return ETDataTensor{objectivec.Object{ID: id}}
}
// Ensure ETDataTensor implements IETDataTensor.
var _ IETDataTensor = ETDataTensor{}

// An interface definition for the [ETDataTensor] class.
//
// # Methods
//
//   - [IETDataTensor.AllocatedImageData]
//   - [IETDataTensor.SetAllocatedImageData]
//   - [IETDataTensor.Blob]
//   - [IETDataTensor.SetBlob]
//   - [IETDataTensor.DataArray]
//   - [IETDataTensor.SetDataArray]
//   - [IETDataTensor.DataPointer]
//   - [IETDataTensor.SetDataPointer]
//   - [IETDataTensor.Float_buffer]
//   - [IETDataTensor.SetFloat_buffer]
//   - [IETDataTensor.ImageBuffer]
//   - [IETDataTensor.SetImageBuffer]
//   - [IETDataTensor.MaxNumberOfElements]
//   - [IETDataTensor.SetMaxNumberOfElements]
//   - [IETDataTensor.Shape]
//   - [IETDataTensor.SetShape]
//   - [IETDataTensor.Strides]
//   - [IETDataTensor.SetStrides]
//   - [IETDataTensor.Type]
//   - [IETDataTensor.SetType]
//   - [IETDataTensor.InitWithBlobContainer]
//   - [IETDataTensor.InitWithBlobContainerDirectBind]
//   - [IETDataTensor.InitWithCVPixelBufferImageParametersError]
//   - [IETDataTensor.InitWithDataTypeShapeStrides]
//
// See: https://developer.apple.com/documentation/Espresso/ETDataTensor
type IETDataTensor interface {
	objectivec.IObject

	// Topic: Methods

	AllocatedImageData() objectivec.IObject
	SetAllocatedImageData(value objectivec.IObject)
	Blob() objectivec.IObject
	SetBlob(value objectivec.IObject)
	DataArray() foundation.INSArray
	SetDataArray(value foundation.INSArray)
	DataPointer() unsafe.Pointer
	SetDataPointer(value unsafe.Pointer)
	Float_buffer() objectivec.IObject
	SetFloat_buffer(value objectivec.IObject)
	ImageBuffer() unsafe.Pointer
	SetImageBuffer(value unsafe.Pointer)
	MaxNumberOfElements() foundation.NSNumber
	SetMaxNumberOfElements(value foundation.NSNumber)
	Shape() foundation.INSArray
	SetShape(value foundation.INSArray)
	Strides() foundation.INSArray
	SetStrides(value foundation.INSArray)
	Type() uint64
	SetType(value uint64)
	InitWithBlobContainer(container objectivec.IObject) ETDataTensor
	InitWithBlobContainerDirectBind(container objectivec.IObject, bind bool) ETDataTensor
	InitWithCVPixelBufferImageParametersError(buffer corevideo.CVImageBufferRef, parameters objectivec.IObject) (ETDataTensor, error)
	InitWithDataTypeShapeStrides(data unsafe.Pointer, type_ uint64, shape objectivec.IObject, strides objectivec.IObject) ETDataTensor
}

// Init initializes the instance.
func (e ETDataTensor) Init() ETDataTensor {
	rv := objc.Send[ETDataTensor](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETDataTensor) Autorelease() ETDataTensor {
	rv := objc.Send[ETDataTensor](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETDataTensor creates a new ETDataTensor instance.
func NewETDataTensor() ETDataTensor {
	class := getETDataTensorClass()
	rv := objc.Send[ETDataTensor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/ETDataTensor/initWithBlobContainer:
func NewETDataTensorWithBlobContainer(container objectivec.IObject) ETDataTensor {
	instance := getETDataTensorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBlobContainer:"), container)
	return ETDataTensorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/ETDataTensor/initWithBlobContainer:directBind:
func NewETDataTensorWithBlobContainerDirectBind(container objectivec.IObject, bind bool) ETDataTensor {
	instance := getETDataTensorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBlobContainer:directBind:"), container, bind)
	return ETDataTensorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/ETDataTensor/initWithCVPixelBuffer:imageParameters:error:
func NewETDataTensorWithCVPixelBufferImageParametersError(buffer corevideo.CVImageBufferRef, parameters objectivec.IObject) (ETDataTensor, error) {
	var errorPtr objc.ID
	instance := getETDataTensorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCVPixelBuffer:imageParameters:error:"), buffer, parameters, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETDataTensor{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETDataTensorFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/Espresso/ETDataTensor/initWithData:type:shape:strides:
func NewETDataTensorWithDataTypeShapeStrides(data unsafe.Pointer, type_ uint64, shape objectivec.IObject, strides objectivec.IObject) ETDataTensor {
	instance := getETDataTensorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:type:shape:strides:"), data, type_, shape, strides)
	return ETDataTensorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/ETDataTensor/initWithBlobContainer:
func (e ETDataTensor) InitWithBlobContainer(container objectivec.IObject) ETDataTensor {
	rv := objc.Send[ETDataTensor](e.ID, objc.Sel("initWithBlobContainer:"), container)
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/ETDataTensor/initWithBlobContainer:directBind:
func (e ETDataTensor) InitWithBlobContainerDirectBind(container objectivec.IObject, bind bool) ETDataTensor {
	rv := objc.Send[ETDataTensor](e.ID, objc.Sel("initWithBlobContainer:directBind:"), container, bind)
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/ETDataTensor/initWithCVPixelBuffer:imageParameters:error:
func (e ETDataTensor) InitWithCVPixelBufferImageParametersError(buffer corevideo.CVImageBufferRef, parameters objectivec.IObject) (ETDataTensor, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initWithCVPixelBuffer:imageParameters:error:"), buffer, parameters, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETDataTensor{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETDataTensorFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/Espresso/ETDataTensor/initWithData:type:shape:strides:
func (e ETDataTensor) InitWithDataTypeShapeStrides(data unsafe.Pointer, type_ uint64, shape objectivec.IObject, strides objectivec.IObject) ETDataTensor {
	rv := objc.Send[ETDataTensor](e.ID, objc.Sel("initWithData:type:shape:strides:"), data, type_, shape, strides)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETDataTensor/allocatedImageData
func (e ETDataTensor) AllocatedImageData() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("allocatedImageData"))
	return objectivec.Object{ID: rv}
}
func (e ETDataTensor) SetAllocatedImageData(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setAllocatedImageData:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETDataTensor/blob
func (e ETDataTensor) Blob() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("blob"))
	return objectivec.Object{ID: rv}
}
func (e ETDataTensor) SetBlob(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setBlob:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETDataTensor/dataArray
func (e ETDataTensor) DataArray() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("dataArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e ETDataTensor) SetDataArray(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setDataArray:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETDataTensor/dataPointer
func (e ETDataTensor) DataPointer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("dataPointer"))
	return rv
}
func (e ETDataTensor) SetDataPointer(value unsafe.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setDataPointer:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETDataTensor/float_buffer
func (e ETDataTensor) Float_buffer() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("float_buffer"))
	return objectivec.Object{ID: rv}
}
func (e ETDataTensor) SetFloat_buffer(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setFloat_buffer:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETDataTensor/imageBuffer
func (e ETDataTensor) ImageBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("imageBuffer"))
	return rv
}
func (e ETDataTensor) SetImageBuffer(value unsafe.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setImageBuffer:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETDataTensor/maxNumberOfElements
func (e ETDataTensor) MaxNumberOfElements() foundation.NSNumber {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("maxNumberOfElements"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (e ETDataTensor) SetMaxNumberOfElements(value foundation.NSNumber) {
	objc.Send[struct{}](e.ID, objc.Sel("setMaxNumberOfElements:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETDataTensor/shape
func (e ETDataTensor) Shape() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("shape"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e ETDataTensor) SetShape(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setShape:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETDataTensor/strides
func (e ETDataTensor) Strides() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("strides"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e ETDataTensor) SetStrides(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setStrides:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETDataTensor/type
func (e ETDataTensor) Type() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("type"))
	return rv
}
func (e ETDataTensor) SetType(value uint64) {
	objc.Send[struct{}](e.ID, objc.Sel("setType:"), value)
}


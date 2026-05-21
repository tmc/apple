// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
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
type IETDataTensor interface {
	objectivec.IObject

	// Topic: Methods

	AllocatedImageData() unsafe.Pointer
	SetAllocatedImageData(value kernel.Pointer)
	Blob() unsafe.Pointer
	SetBlob(value kernel.Pointer)
	DataArray() foundation.INSArray
	SetDataArray(value foundation.INSArray)
	DataPointer() unsafe.Pointer
	SetDataPointer(value kernel.Pointer)
	Float_buffer() FloatBuffer
	SetFloat_buffer(value FloatBuffer)
	ImageBuffer() unsafe.Pointer
	SetImageBuffer(value kernel.Pointer)
	MaxNumberOfElements() foundation.NSNumber
	SetMaxNumberOfElements(value foundation.NSNumber)
	Shape() foundation.INSArray
	SetShape(value foundation.INSArray)
	Strides() foundation.INSArray
	SetStrides(value foundation.INSArray)
	Type() uint64
	SetType(value uint64)
	InitWithBlobContainer(container unsafe.Pointer) ETDataTensor
	InitWithBlobContainerDirectBind(container unsafe.Pointer, bind bool) ETDataTensor
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

func NewETDataTensorWithBlobContainer(container unsafe.Pointer) ETDataTensor {
	instance := getETDataTensorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBlobContainer:"), container)
	return ETDataTensorFromID(rv)
}

func NewETDataTensorWithBlobContainerDirectBind(container unsafe.Pointer, bind bool) ETDataTensor {
	instance := getETDataTensorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBlobContainer:directBind:"), container, bind)
	return ETDataTensorFromID(rv)
}

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

func NewETDataTensorWithDataTypeShapeStrides(data unsafe.Pointer, type_ uint64, shape objectivec.IObject, strides objectivec.IObject) ETDataTensor {
	instance := getETDataTensorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:type:shape:strides:"), data, type_, shape, strides)
	return ETDataTensorFromID(rv)
}

func (e ETDataTensor) InitWithBlobContainer(container unsafe.Pointer) ETDataTensor {
	rv := objc.Send[ETDataTensor](e.ID, objc.Sel("initWithBlobContainer:"), container)
	return rv
}
func (e ETDataTensor) InitWithBlobContainerDirectBind(container unsafe.Pointer, bind bool) ETDataTensor {
	rv := objc.Send[ETDataTensor](e.ID, objc.Sel("initWithBlobContainer:directBind:"), container, bind)
	return rv
}
func (e ETDataTensor) InitWithCVPixelBufferImageParametersError(buffer corevideo.CVImageBufferRef, parameters objectivec.IObject) (ETDataTensor, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initWithCVPixelBuffer:imageParameters:error:"), buffer, parameters, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETDataTensor{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETDataTensorFromID(rv), nil

}
func (e ETDataTensor) InitWithDataTypeShapeStrides(data unsafe.Pointer, type_ uint64, shape objectivec.IObject, strides objectivec.IObject) ETDataTensor {
	rv := objc.Send[ETDataTensor](e.ID, objc.Sel("initWithData:type:shape:strides:"), data, type_, shape, strides)
	return rv
}

func (e ETDataTensor) AllocatedImageData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("allocatedImageData"))
	return rv
}
func (e ETDataTensor) SetAllocatedImageData(value kernel.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setAllocatedImageData:"), value)
}
func (e ETDataTensor) Blob() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("blob"))
	return rv
}
func (e ETDataTensor) SetBlob(value kernel.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setBlob:"), value)
}
func (e ETDataTensor) DataArray() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("dataArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e ETDataTensor) SetDataArray(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setDataArray:"), value)
}
func (e ETDataTensor) DataPointer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("dataPointer"))
	return rv
}
func (e ETDataTensor) SetDataPointer(value kernel.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setDataPointer:"), value)
}
func (e ETDataTensor) Float_buffer() FloatBuffer {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("float_buffer"))
	_ = rv
	return FloatBuffer{}
}
func (e ETDataTensor) SetFloat_buffer(value FloatBuffer) {
	objc.Send[struct{}](e.ID, objc.Sel("setFloat_buffer:"), value)
}
func (e ETDataTensor) ImageBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("imageBuffer"))
	return rv
}
func (e ETDataTensor) SetImageBuffer(value kernel.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setImageBuffer:"), value)
}
func (e ETDataTensor) MaxNumberOfElements() foundation.NSNumber {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("maxNumberOfElements"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (e ETDataTensor) SetMaxNumberOfElements(value foundation.NSNumber) {
	objc.Send[struct{}](e.ID, objc.Sel("setMaxNumberOfElements:"), value)
}
func (e ETDataTensor) Shape() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("shape"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e ETDataTensor) SetShape(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setShape:"), value)
}
func (e ETDataTensor) Strides() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("strides"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e ETDataTensor) SetStrides(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setStrides:"), value)
}
func (e ETDataTensor) Type() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("type"))
	return rv
}
func (e ETDataTensor) SetType(value uint64) {
	objc.Send[struct{}](e.ID, objc.Sel("setType:"), value)
}

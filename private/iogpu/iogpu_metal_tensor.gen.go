// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalTensor] class.
var (
	_IOGPUMetalTensorClass     IOGPUMetalTensorClass
	_IOGPUMetalTensorClassOnce sync.Once
)

func getIOGPUMetalTensorClass() IOGPUMetalTensorClass {
	_IOGPUMetalTensorClassOnce.Do(func() {
		_IOGPUMetalTensorClass = IOGPUMetalTensorClass{class: objc.GetClass("IOGPUMetalTensor")}
	})
	return _IOGPUMetalTensorClass
}

// GetIOGPUMetalTensorClass returns the class object for IOGPUMetalTensor.
func GetIOGPUMetalTensorClass() IOGPUMetalTensorClass {
	return getIOGPUMetalTensorClass()
}

type IOGPUMetalTensorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalTensorClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalTensorClass) Alloc() IOGPUMetalTensor {
	rv := objc.SendIfResponds[IOGPUMetalTensor](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalTensor.Buffer]
//   - [IOGPUMetalTensor.DataType]
//   - [IOGPUMetalTensor.Dimensions]
//   - [IOGPUMetalTensor.GetBytesStridesFromSlice]
//   - [IOGPUMetalTensor.GetBytesStridesFromSliceOriginSliceDimensions]
//   - [IOGPUMetalTensor.GpuResourceID]
//   - [IOGPUMetalTensor.InternalMTLBuffer]
//   - [IOGPUMetalTensor.Iosurface]
//   - [IOGPUMetalTensor.IsTensorViewableWithReshapedDescriptor]
//   - [IOGPUMetalTensor.NewTensorViewWithReshapedDescriptorError]
//   - [IOGPUMetalTensor.NewTensorViewWithSliceError]
//   - [IOGPUMetalTensor.Offset]
//   - [IOGPUMetalTensor.ParentTensor]
//   - [IOGPUMetalTensor.Plane]
//   - [IOGPUMetalTensor.ReplaceSliceWithBytesStrides]
//   - [IOGPUMetalTensor.ReplaceSliceOriginSliceDimensionsWithBytesStrides]
//   - [IOGPUMetalTensor.ResourceIndex]
//   - [IOGPUMetalTensor.Strides]
//   - [IOGPUMetalTensor.Usage]
//   - [IOGPUMetalTensor.InitWithBuffer]
//   - [IOGPUMetalTensor.BufferOffset]
type IOGPUMetalTensor struct {
	IOGPUMetalResource
}

// IOGPUMetalTensorFromID constructs a [IOGPUMetalTensor] from an objc.ID.
func IOGPUMetalTensorFromID(id objc.ID) IOGPUMetalTensor {
	return IOGPUMetalTensor{IOGPUMetalResource: IOGPUMetalResourceFromID(id)}
}

// Ensure IOGPUMetalTensor implements IIOGPUMetalTensor.
var _ IIOGPUMetalTensor = IOGPUMetalTensor{}

// An interface definition for the [IOGPUMetalTensor] class.
//
// # Methods
//
//   - [IIOGPUMetalTensor.Buffer]
//   - [IIOGPUMetalTensor.DataType]
//   - [IIOGPUMetalTensor.Dimensions]
//   - [IIOGPUMetalTensor.GetBytesStridesFromSlice]
//   - [IIOGPUMetalTensor.GetBytesStridesFromSliceOriginSliceDimensions]
//   - [IIOGPUMetalTensor.GpuResourceID]
//   - [IIOGPUMetalTensor.InternalMTLBuffer]
//   - [IIOGPUMetalTensor.Iosurface]
//   - [IIOGPUMetalTensor.IsTensorViewableWithReshapedDescriptor]
//   - [IIOGPUMetalTensor.NewTensorViewWithReshapedDescriptorError]
//   - [IIOGPUMetalTensor.NewTensorViewWithSliceError]
//   - [IIOGPUMetalTensor.Offset]
//   - [IIOGPUMetalTensor.ParentTensor]
//   - [IIOGPUMetalTensor.Plane]
//   - [IIOGPUMetalTensor.ReplaceSliceWithBytesStrides]
//   - [IIOGPUMetalTensor.ReplaceSliceOriginSliceDimensionsWithBytesStrides]
//   - [IIOGPUMetalTensor.ResourceIndex]
//   - [IIOGPUMetalTensor.Strides]
//   - [IIOGPUMetalTensor.Usage]
//   - [IIOGPUMetalTensor.InitWithBuffer]
//   - [IIOGPUMetalTensor.BufferOffset]
type IIOGPUMetalTensor interface {
	IIOGPUMetalResource

	// Topic: Methods

	Buffer() unsafe.Pointer
	DataType() int64
	Dimensions() metal.MTLTensorExtents
	GetBytesStridesFromSlice(bytes unsafe.Pointer, strides objectivec.IObject, slice MTLTensorSlice)
	GetBytesStridesFromSliceOriginSliceDimensions(bytes unsafe.Pointer, strides objectivec.IObject, origin objectivec.IObject, dimensions objectivec.IObject)
	GpuResourceID() metal.MTLResourceID
	InternalMTLBuffer() objectivec.IObject
	Iosurface() iosurface.IOSurfaceRef
	IsTensorViewableWithReshapedDescriptor(descriptor objectivec.IObject) bool
	NewTensorViewWithReshapedDescriptorError(descriptor objectivec.IObject) (objectivec.IObject, error)
	NewTensorViewWithSliceError(slice MTLTensorSlice) (objectivec.IObject, error)
	Offset() uint64
	ParentTensor() unsafe.Pointer
	Plane() uint64
	ReplaceSliceWithBytesStrides(slice MTLTensorSlice, bytes unsafe.Pointer, strides objectivec.IObject)
	ReplaceSliceOriginSliceDimensionsWithBytesStrides(origin objectivec.IObject, dimensions objectivec.IObject, bytes unsafe.Pointer, strides objectivec.IObject)
	ResourceIndex() uint64
	Strides() metal.MTLTensorExtents
	Usage() uint64
	InitWithBuffer(buffer objectivec.IObject) IOGPUMetalTensor
	BufferOffset() uint64
}

// Init initializes the instance.
func (i IOGPUMetalTensor) Init() IOGPUMetalTensor {
	rv := objc.SendIfResponds[IOGPUMetalTensor](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalTensor) Autorelease() IOGPUMetalTensor {
	rv := objc.SendIfResponds[IOGPUMetalTensor](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalTensor creates a new IOGPUMetalTensor instance.
func NewIOGPUMetalTensor() IOGPUMetalTensor {
	class := getIOGPUMetalTensorClass()
	rv := objc.SendIfResponds[IOGPUMetalTensor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalTensorMemorylessDescriptor(memoryless objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalTensor {
	instance := getIOGPUMetalTensorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initMemoryless:descriptor:"), memoryless, descriptor)
	return IOGPUMetalTensorFromID(rv)
}

func NewGPUMetalTensorStandinWithDevice(device objectivec.IObject) IOGPUMetalTensor {
	instance := getIOGPUMetalTensorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initStandinWithDevice:"), device)
	return IOGPUMetalTensorFromID(rv)
}

func NewGPUMetalTensorWithBuffer(buffer objectivec.IObject) IOGPUMetalTensor {
	instance := getIOGPUMetalTensorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBuffer:"), buffer)
	return IOGPUMetalTensorFromID(rv)
}

func NewGPUMetalTensorWithDeviceOptionsArgsArgsSize(device objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalTensor {
	instance := getIOGPUMetalTensorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:options:args:argsSize:"), device, options, unsafe.Pointer(args), size)
	return IOGPUMetalTensorFromID(rv)
}

func NewGPUMetalTensorWithDeviceRemoteStorageResourceOptionsArgsArgsSize(device objectivec.IObject, resource objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalTensor {
	instance := getIOGPUMetalTensorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:remoteStorageResource:options:args:argsSize:"), device, resource, options, unsafe.Pointer(args), size)
	return IOGPUMetalTensorFromID(rv)
}

func NewGPUMetalTensorWithResource(resource objectivec.IObject) IOGPUMetalTensor {
	instance := getIOGPUMetalTensorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return IOGPUMetalTensorFromID(rv)
}

func (i IOGPUMetalTensor) GetBytesStridesFromSlice(bytes unsafe.Pointer, strides objectivec.IObject, slice MTLTensorSlice) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("getBytes:strides:fromSlice:"), bytes, strides, slice)
}
func (i IOGPUMetalTensor) GetBytesStridesFromSliceOriginSliceDimensions(bytes unsafe.Pointer, strides objectivec.IObject, origin objectivec.IObject, dimensions objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("getBytes:strides:fromSliceOrigin:sliceDimensions:"), bytes, strides, origin, dimensions)
}
func (i IOGPUMetalTensor) InternalMTLBuffer() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("internalMTLBuffer"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalTensor) Iosurface() iosurface.IOSurfaceRef {
	rv := objc.SendIfResponds[iosurface.IOSurfaceRef](i.ID, objc.Sel("iosurface"))
	return iosurface.IOSurfaceRef(rv)
}
func (i IOGPUMetalTensor) IsTensorViewableWithReshapedDescriptor(descriptor objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("isTensorViewableWithReshapedDescriptor:"), descriptor)
	return rv
}
func (i IOGPUMetalTensor) NewTensorViewWithReshapedDescriptorError(descriptor objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("newTensorViewWithReshapedDescriptor:error:"), descriptor, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (i IOGPUMetalTensor) NewTensorViewWithSliceError(slice MTLTensorSlice) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("newTensorViewWithSlice:error:"), slice, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (i IOGPUMetalTensor) ReplaceSliceWithBytesStrides(slice MTLTensorSlice, bytes unsafe.Pointer, strides objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("replaceSlice:withBytes:strides:"), slice, bytes, strides)
}
func (i IOGPUMetalTensor) ReplaceSliceOriginSliceDimensionsWithBytesStrides(origin objectivec.IObject, dimensions objectivec.IObject, bytes unsafe.Pointer, strides objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("replaceSliceOrigin:sliceDimensions:withBytes:strides:"), origin, dimensions, bytes, strides)
}
func (i IOGPUMetalTensor) InitWithBuffer(buffer objectivec.IObject) IOGPUMetalTensor {
	rv := objc.SendIfResponds[IOGPUMetalTensor](i.ID, objc.Sel("initWithBuffer:"), buffer)
	return rv
}

func (i IOGPUMetalTensor) Buffer() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("buffer"))
	return rv
}
func (i IOGPUMetalTensor) BufferOffset() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("bufferOffset"))
	return rv
}
func (i IOGPUMetalTensor) DataType() int64 {
	rv := objc.SendIfResponds[int64](i.ID, objc.Sel("dataType"))
	return rv
}
func (i IOGPUMetalTensor) Dimensions() metal.MTLTensorExtents {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("dimensions"))
	return metal.MTLTensorExtentsFromID(objc.ID(rv))
}
func (i IOGPUMetalTensor) GpuResourceID() metal.MTLResourceID {
	rv := objc.SendIfResponds[metal.MTLResourceID](i.ID, objc.Sel("gpuResourceID"))
	return metal.MTLResourceID(rv)
}
func (i IOGPUMetalTensor) Offset() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("offset"))
	return rv
}
func (i IOGPUMetalTensor) ParentTensor() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("parentTensor"))
	return rv
}
func (i IOGPUMetalTensor) Plane() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("plane"))
	return rv
}
func (i IOGPUMetalTensor) ResourceIndex() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("resourceIndex"))
	return rv
}
func (i IOGPUMetalTensor) Strides() metal.MTLTensorExtents {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("strides"))
	return metal.MTLTensorExtentsFromID(objc.ID(rv))
}
func (i IOGPUMetalTensor) Usage() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("usage"))
	return rv
}

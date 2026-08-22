// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNDArray] class.
var (
	_MPSNDArrayClass     MPSNDArrayClass
	_MPSNDArrayClassOnce sync.Once
)

func getMPSNDArrayClass() MPSNDArrayClass {
	_MPSNDArrayClassOnce.Do(func() {
		_MPSNDArrayClass = MPSNDArrayClass{class: objc.GetClass("MPSNDArray")}
	})
	return _MPSNDArrayClass
}

// GetMPSNDArrayClass returns the class object for MPSNDArray.
func GetMPSNDArrayClass() MPSNDArrayClass {
	return getMPSNDArrayClass()
}

type MPSNDArrayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayClass) Alloc() MPSNDArray {
	rv := objc.Send[MPSNDArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNDArray.InitWithBufferOffsetDescriptor]
//   - [MPSNDArray.InitWithDeviceDescriptor]
//   - [MPSNDArray.InitWithDeviceScalar]
//
// # Instance Properties
//
//   - [MPSNDArray.DataType]
//   - [MPSNDArray.DataTypeSize]
//   - [MPSNDArray.Device]
//   - [MPSNDArray.Label]
//   - [MPSNDArray.SetLabel]
//   - [MPSNDArray.NumberOfDimensions]
//   - [MPSNDArray.Parent]
//
// # Instance Methods
//
//   - [MPSNDArray.ArrayViewWithDescriptor]
//   - [MPSNDArray.ArrayViewWithCommandBufferDescriptorAliasing]
//   - [MPSNDArray.ArrayViewWithDimensionCountDimensionSizesStrides]
//   - [MPSNDArray.ArrayViewWithShapeStrides]
//   - [MPSNDArray.Descriptor]
//   - [MPSNDArray.ExportDataWithCommandBufferToBufferDestinationDataTypeOffsetRowStrides]
//   - [MPSNDArray.ExportDataWithCommandBufferToImagesOffset]
//   - [MPSNDArray.ImportDataWithCommandBufferFromImagesOffset]
//   - [MPSNDArray.ImportDataWithCommandBufferFromBufferSourceDataTypeOffsetRowStrides]
//   - [MPSNDArray.LengthOfDimension]
//   - [MPSNDArray.ReadBytesStrideBytes]
//   - [MPSNDArray.ResourceSize]
//   - [MPSNDArray.SynchronizeOnCommandBuffer]
//   - [MPSNDArray.UserBuffer]
//   - [MPSNDArray.WriteBytesStrideBytes]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray
type MPSNDArray struct {
	objectivec.Object
}

// MPSNDArrayFromID constructs a [MPSNDArray] from an objc.ID.
func MPSNDArrayFromID(id objc.ID) MPSNDArray {
	return MPSNDArray{objectivec.Object{ID: id}}
}

// NOTE: MPSNDArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArray] class.
//
// # Initializers
//
//   - [IMPSNDArray.InitWithBufferOffsetDescriptor]
//   - [IMPSNDArray.InitWithDeviceDescriptor]
//   - [IMPSNDArray.InitWithDeviceScalar]
//
// # Instance Properties
//
//   - [IMPSNDArray.DataType]
//   - [IMPSNDArray.DataTypeSize]
//   - [IMPSNDArray.Device]
//   - [IMPSNDArray.Label]
//   - [IMPSNDArray.SetLabel]
//   - [IMPSNDArray.NumberOfDimensions]
//   - [IMPSNDArray.Parent]
//
// # Instance Methods
//
//   - [IMPSNDArray.ArrayViewWithDescriptor]
//   - [IMPSNDArray.ArrayViewWithCommandBufferDescriptorAliasing]
//   - [IMPSNDArray.ArrayViewWithDimensionCountDimensionSizesStrides]
//   - [IMPSNDArray.ArrayViewWithShapeStrides]
//   - [IMPSNDArray.Descriptor]
//   - [IMPSNDArray.ExportDataWithCommandBufferToBufferDestinationDataTypeOffsetRowStrides]
//   - [IMPSNDArray.ExportDataWithCommandBufferToImagesOffset]
//   - [IMPSNDArray.ImportDataWithCommandBufferFromImagesOffset]
//   - [IMPSNDArray.ImportDataWithCommandBufferFromBufferSourceDataTypeOffsetRowStrides]
//   - [IMPSNDArray.LengthOfDimension]
//   - [IMPSNDArray.ReadBytesStrideBytes]
//   - [IMPSNDArray.ResourceSize]
//   - [IMPSNDArray.SynchronizeOnCommandBuffer]
//   - [IMPSNDArray.UserBuffer]
//   - [IMPSNDArray.WriteBytesStrideBytes]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray
type IMPSNDArray interface {
	objectivec.IObject

	// Topic: Initializers

	InitWithBufferOffsetDescriptor(buffer metal.MTLBuffer, offset uint, descriptor IMPSNDArrayDescriptor) MPSNDArray
	InitWithDeviceDescriptor(device metal.MTLDevice, descriptor IMPSNDArrayDescriptor) MPSNDArray
	InitWithDeviceScalar(device metal.MTLDevice, value float64) MPSNDArray

	// Topic: Instance Properties

	DataType() MPSDataType
	DataTypeSize() uintptr
	Device() metal.MTLDevice
	Label() string
	SetLabel(value string)
	NumberOfDimensions() uint
	Parent() IMPSNDArray

	// Topic: Instance Methods

	ArrayViewWithDescriptor(descriptor IMPSNDArrayDescriptor) IMPSNDArray
	ArrayViewWithCommandBufferDescriptorAliasing(cmdBuf metal.MTLCommandBuffer, descriptor IMPSNDArrayDescriptor, aliasing MPSAliasingStrategy) IMPSNDArray
	ArrayViewWithDimensionCountDimensionSizesStrides(numberOfDimensions uint, dimensionSizes *uint, dimStrides *uint) IMPSNDArray
	ArrayViewWithShapeStrides(shape MPSShape, strides MPSShape) IMPSNDArray
	Descriptor() IMPSNDArrayDescriptor
	ExportDataWithCommandBufferToBufferDestinationDataTypeOffsetRowStrides(cmdBuf metal.MTLCommandBuffer, buffer metal.MTLBuffer, destinationDataType MPSDataType, offset uint, rowStrides *int)
	ExportDataWithCommandBufferToImagesOffset(cmdBuf metal.MTLCommandBuffer, images MPSImageBatch, offset MPSImageCoordinate)
	ImportDataWithCommandBufferFromImagesOffset(cmdBuf metal.MTLCommandBuffer, images MPSImageBatch, offset MPSImageCoordinate)
	ImportDataWithCommandBufferFromBufferSourceDataTypeOffsetRowStrides(cmdBuf metal.MTLCommandBuffer, buffer metal.MTLBuffer, sourceDataType MPSDataType, offset uint, rowStrides *int)
	LengthOfDimension(dimensionIndex uint) uint
	ReadBytesStrideBytes(buffer unsafe.Pointer, strideBytesPerDimension *int)
	ResourceSize() uint
	SynchronizeOnCommandBuffer(commandBuffer metal.MTLCommandBuffer)
	UserBuffer() metal.MTLBuffer
	WriteBytesStrideBytes(buffer unsafe.Pointer, strideBytesPerDimension *int)
}

// Init initializes the instance.
func (n MPSNDArray) Init() MPSNDArray {
	rv := objc.Send[MPSNDArray](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArray) Autorelease() MPSNDArray {
	rv := objc.Send[MPSNDArray](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArray creates a new MPSNDArray instance.
func NewMPSNDArray() MPSNDArray {
	class := getMPSNDArrayClass()
	rv := objc.Send[MPSNDArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/init(buffer:offset:descriptor:)
func NewNDArrayWithBufferOffsetDescriptor(buffer metal.MTLBuffer, offset uint, descriptor IMPSNDArrayDescriptor) MPSNDArray {
	instance := getMPSNDArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBuffer:offset:descriptor:"), buffer, offset, descriptor)
	return MPSNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/init(device:descriptor:)
func NewNDArrayWithDeviceDescriptor(device metal.MTLDevice, descriptor IMPSNDArrayDescriptor) MPSNDArray {
	instance := getMPSNDArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	return MPSNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/init(device:scalar:)
func NewNDArrayWithDeviceScalar(device metal.MTLDevice, value float64) MPSNDArray {
	instance := getMPSNDArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:scalar:"), device, value)
	return MPSNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/init(buffer:offset:descriptor:)
func (n MPSNDArray) InitWithBufferOffsetDescriptor(buffer metal.MTLBuffer, offset uint, descriptor IMPSNDArrayDescriptor) MPSNDArray {
	rv := objc.Send[MPSNDArray](n.ID, objc.Sel("initWithBuffer:offset:descriptor:"), buffer, offset, descriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/init(device:descriptor:)
func (n MPSNDArray) InitWithDeviceDescriptor(device metal.MTLDevice, descriptor IMPSNDArrayDescriptor) MPSNDArray {
	rv := objc.Send[MPSNDArray](n.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/init(device:scalar:)
func (n MPSNDArray) InitWithDeviceScalar(device metal.MTLDevice, value float64) MPSNDArray {
	rv := objc.Send[MPSNDArray](n.ID, objc.Sel("initWithDevice:scalar:"), device, value)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/arrayView(with:)
func (n MPSNDArray) ArrayViewWithDescriptor(descriptor IMPSNDArrayDescriptor) IMPSNDArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("arrayViewWithDescriptor:"), descriptor)
	return MPSNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/arrayView(with:descriptor:aliasing:)
func (n MPSNDArray) ArrayViewWithCommandBufferDescriptorAliasing(cmdBuf metal.MTLCommandBuffer, descriptor IMPSNDArrayDescriptor, aliasing MPSAliasingStrategy) IMPSNDArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("arrayViewWithCommandBuffer:descriptor:aliasing:"), cmdBuf, descriptor, aliasing)
	return MPSNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/arrayView(withDimensionCount:dimensionSizes:strides:)
func (n MPSNDArray) ArrayViewWithDimensionCountDimensionSizesStrides(numberOfDimensions uint, dimensionSizes *uint, dimStrides *uint) IMPSNDArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("arrayViewWithDimensionCount:dimensionSizes:strides:"), numberOfDimensions, unsafe.Pointer(dimensionSizes), unsafe.Pointer(dimStrides))
	return MPSNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/arrayView(withShape:strides:)
func (n MPSNDArray) ArrayViewWithShapeStrides(shape MPSShape, strides MPSShape) IMPSNDArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("arrayViewWithShape:strides:"), shape, strides)
	return MPSNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/descriptor()
func (n MPSNDArray) Descriptor() IMPSNDArrayDescriptor {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("descriptor"))
	return MPSNDArrayDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/exportData(with:to:destinationDataType:offset:rowStrides:)
func (n MPSNDArray) ExportDataWithCommandBufferToBufferDestinationDataTypeOffsetRowStrides(cmdBuf metal.MTLCommandBuffer, buffer metal.MTLBuffer, destinationDataType MPSDataType, offset uint, rowStrides *int) {
	objc.Send[objc.ID](n.ID, objc.Sel("exportDataWithCommandBuffer:toBuffer:destinationDataType:offset:rowStrides:"), cmdBuf, buffer, destinationDataType, offset, rowStrides)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/exportData(with:to:offset:)
func (n MPSNDArray) ExportDataWithCommandBufferToImagesOffset(cmdBuf metal.MTLCommandBuffer, images MPSImageBatch, offset MPSImageCoordinate) {
	objc.Send[objc.ID](n.ID, objc.Sel("exportDataWithCommandBuffer:toImages:offset:"), cmdBuf, images, offset)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/importData(with:from:offset:)
func (n MPSNDArray) ImportDataWithCommandBufferFromImagesOffset(cmdBuf metal.MTLCommandBuffer, images MPSImageBatch, offset MPSImageCoordinate) {
	objc.Send[objc.ID](n.ID, objc.Sel("importDataWithCommandBuffer:fromImages:offset:"), cmdBuf, images, offset)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/importData(with:from:sourceDataType:offset:rowStrides:)
func (n MPSNDArray) ImportDataWithCommandBufferFromBufferSourceDataTypeOffsetRowStrides(cmdBuf metal.MTLCommandBuffer, buffer metal.MTLBuffer, sourceDataType MPSDataType, offset uint, rowStrides *int) {
	objc.Send[objc.ID](n.ID, objc.Sel("importDataWithCommandBuffer:fromBuffer:sourceDataType:offset:rowStrides:"), cmdBuf, buffer, sourceDataType, offset, rowStrides)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/length(ofDimension:)
func (n MPSNDArray) LengthOfDimension(dimensionIndex uint) uint {
	rv := objc.Send[uint](n.ID, objc.Sel("lengthOfDimension:"), dimensionIndex)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/readBytes(_:strideBytes:)
func (n MPSNDArray) ReadBytesStrideBytes(buffer unsafe.Pointer, strideBytesPerDimension *int) {
	objc.Send[objc.ID](n.ID, objc.Sel("readBytes:strideBytes:"), buffer, strideBytesPerDimension)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/resourceSize()
func (n MPSNDArray) ResourceSize() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("resourceSize"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/synchronize(on:)
func (n MPSNDArray) SynchronizeOnCommandBuffer(commandBuffer metal.MTLCommandBuffer) {
	objc.Send[objc.ID](n.ID, objc.Sel("synchronizeOnCommandBuffer:"), commandBuffer)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/userBuffer()
func (n MPSNDArray) UserBuffer() metal.MTLBuffer {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("userBuffer"))
	return metal.MTLBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/writeBytes(_:strideBytes:)
func (n MPSNDArray) WriteBytesStrideBytes(buffer unsafe.Pointer, strideBytesPerDimension *int) {
	objc.Send[objc.ID](n.ID, objc.Sel("writeBytes:strideBytes:"), buffer, strideBytesPerDimension)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/defaultAllocator()
func (_MPSNDArrayClass MPSNDArrayClass) DefaultAllocator() MPSNDArrayAllocator {
	rv := objc.Send[objc.ID](objc.ID(_MPSNDArrayClass.class), objc.Sel("defaultAllocator"))
	return MPSNDArrayAllocatorObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/dataType
func (n MPSNDArray) DataType() MPSDataType {
	rv := objc.Send[MPSDataType](n.ID, objc.Sel("dataType"))
	return MPSDataType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/dataTypeSize
func (n MPSNDArray) DataTypeSize() uintptr {
	rv := objc.Send[uintptr](n.ID, objc.Sel("dataTypeSize"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/device
func (n MPSNDArray) Device() metal.MTLDevice {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("device"))
	return metal.MTLDeviceObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/label
func (n MPSNDArray) Label() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (n MPSNDArray) SetLabel(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setLabel:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/numberOfDimensions
func (n MPSNDArray) NumberOfDimensions() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("numberOfDimensions"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/parent
func (n MPSNDArray) Parent() IMPSNDArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("parent"))
	return MPSNDArrayFromID(objc.ID(rv))
}

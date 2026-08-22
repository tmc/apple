// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/metalperformanceshaders"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphTensorData] class.
var (
	_MPSGraphTensorDataClass     MPSGraphTensorDataClass
	_MPSGraphTensorDataClassOnce sync.Once
)

func getMPSGraphTensorDataClass() MPSGraphTensorDataClass {
	_MPSGraphTensorDataClassOnce.Do(func() {
		_MPSGraphTensorDataClass = MPSGraphTensorDataClass{class: objc.GetClass("MPSGraphTensorData")}
	})
	return _MPSGraphTensorDataClass
}

// GetMPSGraphTensorDataClass returns the class object for MPSGraphTensorData.
func GetMPSGraphTensorDataClass() MPSGraphTensorDataClass {
	return getMPSGraphTensorDataClass()
}

type MPSGraphTensorDataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphTensorDataClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphTensorDataClass) Alloc() MPSGraphTensorData {
	rv := objc.Send[MPSGraphTensorData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The representation of a compute data type.
//
// # Overview
//
// Pass data to a graph using a tensor data, a reference will be taken to your
// data and used just in time when the graph is run.
//
// # Initializers
//
//   - [MPSGraphTensorData.InitWithMPSMatrix]: Initializes a tensor data with an MPS matrix.
//   - [MPSGraphTensorData.InitWithMPSNDArray]: Initializes an MPSGraphTensorData with an MPS ndarray.
//   - [MPSGraphTensorData.InitWithMPSImageBatch]: Initializes a tensor data with an MPS image batch.
//   - [MPSGraphTensorData.InitWithMTLTensor]: Initializes an MPSGraphTensorData with an MTLTensor.
//   - [MPSGraphTensorData.InitWithMPSVector]: Initializes a tensor data with an MPS vector.
//   - [MPSGraphTensorData.InitWithMPSVectorRank]: Initializes a tensor data with an MPS vector enforcing rank of the result.
//   - [MPSGraphTensorData.InitWithMPSMatrixRank]: Initializes a tensor data with an MPS matrix enforcing rank of the result.
//   - [MPSGraphTensorData.InitWithMTLBufferShapeDataType]: Initializes an tensor data with a metal buffer.
//   - [MPSGraphTensorData.InitWithMTLBufferShapeDataTypeRowBytes]: Initializes an tensor data with a metal buffer.
//   - [MPSGraphTensorData.InitWithDeviceDataShapeDataType]: Initializes the tensor data with an [NSData] on a device.
//
// # Instance Properties
//
//   - [MPSGraphTensorData.DataType]: The data type of the tensor data.
//   - [MPSGraphTensorData.Device]: The device of the tensor data.
//   - [MPSGraphTensorData.Shape]: The shape of the tensor data.
//
// # Instance Methods
//
//   - [MPSGraphTensorData.Mpsndarray]: Return an mpsndarray object will copy contents if the contents are not stored in an MPS ndarray.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData
type MPSGraphTensorData struct {
	MPSGraphObject
}

// MPSGraphTensorDataFromID constructs a [MPSGraphTensorData] from an objc.ID.
//
// The representation of a compute data type.
func MPSGraphTensorDataFromID(id objc.ID) MPSGraphTensorData {
	return MPSGraphTensorData{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphTensorData adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphTensorData] class.
//
// # Initializers
//
//   - [IMPSGraphTensorData.InitWithMPSMatrix]: Initializes a tensor data with an MPS matrix.
//   - [IMPSGraphTensorData.InitWithMPSNDArray]: Initializes an MPSGraphTensorData with an MPS ndarray.
//   - [IMPSGraphTensorData.InitWithMPSImageBatch]: Initializes a tensor data with an MPS image batch.
//   - [IMPSGraphTensorData.InitWithMTLTensor]: Initializes an MPSGraphTensorData with an MTLTensor.
//   - [IMPSGraphTensorData.InitWithMPSVector]: Initializes a tensor data with an MPS vector.
//   - [IMPSGraphTensorData.InitWithMPSVectorRank]: Initializes a tensor data with an MPS vector enforcing rank of the result.
//   - [IMPSGraphTensorData.InitWithMPSMatrixRank]: Initializes a tensor data with an MPS matrix enforcing rank of the result.
//   - [IMPSGraphTensorData.InitWithMTLBufferShapeDataType]: Initializes an tensor data with a metal buffer.
//   - [IMPSGraphTensorData.InitWithMTLBufferShapeDataTypeRowBytes]: Initializes an tensor data with a metal buffer.
//   - [IMPSGraphTensorData.InitWithDeviceDataShapeDataType]: Initializes the tensor data with an [NSData] on a device.
//
// # Instance Properties
//
//   - [IMPSGraphTensorData.DataType]: The data type of the tensor data.
//   - [IMPSGraphTensorData.Device]: The device of the tensor data.
//   - [IMPSGraphTensorData.Shape]: The shape of the tensor data.
//
// # Instance Methods
//
//   - [IMPSGraphTensorData.Mpsndarray]: Return an mpsndarray object will copy contents if the contents are not stored in an MPS ndarray.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData
type IMPSGraphTensorData interface {
	IMPSGraphObject

	// Topic: Initializers

	// Initializes a tensor data with an MPS matrix.
	InitWithMPSMatrix(matrix *metalperformanceshaders.MPSMatrix) MPSGraphTensorData
	// Initializes an MPSGraphTensorData with an MPS ndarray.
	InitWithMPSNDArray(ndarray *metalperformanceshaders.MPSNDArray) MPSGraphTensorData
	// Initializes a tensor data with an MPS image batch.
	InitWithMPSImageBatch(imageBatch foundation.NSArray) MPSGraphTensorData
	// Initializes an MPSGraphTensorData with an MTLTensor.
	InitWithMTLTensor(tensor metal.MTLTensor) MPSGraphTensorData
	// Initializes a tensor data with an MPS vector.
	InitWithMPSVector(vector *metalperformanceshaders.MPSVector) MPSGraphTensorData
	// Initializes a tensor data with an MPS vector enforcing rank of the result.
	InitWithMPSVectorRank(vector *metalperformanceshaders.MPSVector, rank uint) MPSGraphTensorData
	// Initializes a tensor data with an MPS matrix enforcing rank of the result.
	InitWithMPSMatrixRank(matrix *metalperformanceshaders.MPSMatrix, rank uint) MPSGraphTensorData
	// Initializes an tensor data with a metal buffer.
	InitWithMTLBufferShapeDataType(buffer metal.MTLBuffer, shape foundation.NSArray, dataType uint32) MPSGraphTensorData
	// Initializes an tensor data with a metal buffer.
	InitWithMTLBufferShapeDataTypeRowBytes(buffer metal.MTLBuffer, shape foundation.NSArray, dataType uint32, rowBytes uint) MPSGraphTensorData
	// Initializes the tensor data with an [NSData] on a device.
	InitWithDeviceDataShapeDataType(device IMPSGraphDevice, data foundation.NSData, shape foundation.NSArray, dataType uint32) MPSGraphTensorData

	// Topic: Instance Properties

	// The data type of the tensor data.
	DataType() uint32
	// The device of the tensor data.
	Device() IMPSGraphDevice
	// The shape of the tensor data.
	Shape() foundation.NSArray

	// Topic: Instance Methods

	// Return an mpsndarray object will copy contents if the contents are not stored in an MPS ndarray.
	Mpsndarray() metalperformanceshaders.MPSNDArray
}

// Init initializes the instance.
func (g MPSGraphTensorData) Init() MPSGraphTensorData {
	rv := objc.Send[MPSGraphTensorData](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphTensorData) Autorelease() MPSGraphTensorData {
	rv := objc.Send[MPSGraphTensorData](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphTensorData creates a new MPSGraphTensorData instance.
func NewMPSGraphTensorData() MPSGraphTensorData {
	class := getMPSGraphTensorDataClass()
	rv := objc.Send[MPSGraphTensorData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes the tensor data with an [NSData] on a device.
//
// device: MPSDevice on which the MPSGraphTensorData exists
//
// data: NSData from which to copy the contents
//
// shape: Shape of the output tensor
//
// dataType: dataType of the placeholder tensor
//
// # Return Value
//
// A valid MPSGraphTensorData, or nil if allocation failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(device:data:shape:dataType:)
func NewGraphTensorDataWithDeviceDataShapeDataType(device IMPSGraphDevice, data foundation.NSData, shape foundation.NSArray, dataType uint32) MPSGraphTensorData {
	instance := getMPSGraphTensorDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:data:shape:dataType:"), device, data, shape, dataType)
	return MPSGraphTensorDataFromID(rv)
}

// Initializes a tensor data with an MPS image batch.
//
// imageBatch: The device on which the kernel will run, unorm8 and unorm16 images will
// create a float32 tensorData
//
// # Return Value
//
// A valid MPSGraphTensorData, or nil if allocation failure.
//
// # Discussion
//
// The dataLayout used will be NHWC, call a transpose or permute to change to
// a layout of your choice.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:)-511a
func NewGraphTensorDataWithMPSImageBatch(imageBatch foundation.NSArray) MPSGraphTensorData {
	instance := getMPSGraphTensorDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMPSImageBatch:"), imageBatch)
	return MPSGraphTensorDataFromID(rv)
}

// Initializes a tensor data with an MPS matrix.
//
// matrix: MPSMatrix to be used within the MPSGraphTensorData
//
// # Return Value
//
// A valid MPSGraphTensorData, or nil if allocation failure.
//
// # Discussion
//
// The device of the MPSMatrix will be used to get the MPSDevice for this
// MPSGraphTensorData.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:)-2go2
func NewGraphTensorDataWithMPSMatrix(matrix *metalperformanceshaders.MPSMatrix) MPSGraphTensorData {
	instance := getMPSGraphTensorDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMPSMatrix:"), matrix.ID)
	return MPSGraphTensorDataFromID(rv)
}

// Initializes a tensor data with an MPS matrix enforcing rank of the result.
//
// matrix: MPSMatrix to be used within the MPSGraphTensorData
//
// rank: The rank of the resulting TensorData tensor. NOTE: must be within { 1, …
// ,16 }.
//
// # Return Value
//
// A valid MPSGraphTensorData of given rank, or nil if allocation failure.
//
// # Discussion
//
// The device of the MPSMatrix will be used to get the MPSDevice for this
// MPSGraphTensorData.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:rank:)-1lnxg
func NewGraphTensorDataWithMPSMatrixRank(matrix *metalperformanceshaders.MPSMatrix, rank uint) MPSGraphTensorData {
	instance := getMPSGraphTensorDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMPSMatrix:rank:"), matrix.ID, rank)
	return MPSGraphTensorDataFromID(rv)
}

// Initializes an MPSGraphTensorData with an MPS ndarray.
//
// ndarray: MPSNDArray to be used within the MPSGraphTensorData.
//
// # Return Value
//
// A valid MPSGraphTensorData, or nil if allocation failure.
//
// # Discussion
//
// The device of the MPSNDArray will be used to get the MPSDevice for this
// MPSGraphTensorData.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:)-4bnfb
func NewGraphTensorDataWithMPSNDArray(ndarray *metalperformanceshaders.MPSNDArray) MPSGraphTensorData {
	instance := getMPSGraphTensorDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMPSNDArray:"), ndarray.ID)
	return MPSGraphTensorDataFromID(rv)
}

// Initializes a tensor data with an MPS vector.
//
// vector: MPSVector to be used within the MPSGraphTensorData
//
// # Return Value
//
// A valid MPSGraphTensorData, or nil if allocation failure.
//
// # Discussion
//
// The device of the MPSVector will be used to get the MPSDevice for this
// MPSGraphTensorData.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:)-9kgoe
func NewGraphTensorDataWithMPSVector(vector *metalperformanceshaders.MPSVector) MPSGraphTensorData {
	instance := getMPSGraphTensorDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMPSVector:"), vector.ID)
	return MPSGraphTensorDataFromID(rv)
}

// Initializes a tensor data with an MPS vector enforcing rank of the result.
//
// vector: MPSVector to be used within the MPSGraphTensorData
//
// rank: The rank of the resulting TensorData tensor. NOTE: must be within { 1, …
// ,16 }.
//
// # Return Value
//
// A valid MPSGraphTensorData, or nil if allocation failure.
//
// # Discussion
//
// The device of the MPSVector will be used to get the MPSDevice for this
// MPSGraphTensorData.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:rank:)-1e4ks
func NewGraphTensorDataWithMPSVectorRank(vector *metalperformanceshaders.MPSVector, rank uint) MPSGraphTensorData {
	instance := getMPSGraphTensorDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMPSVector:rank:"), vector.ID, rank)
	return MPSGraphTensorDataFromID(rv)
}

// Initializes an tensor data with a metal buffer.
//
// buffer: MTLBuffer to be used within the MPSGraphTensorData
//
// shape: Shape of the output tensor
//
// dataType: dataType of the placeholder tensor
//
// # Return Value
//
// A valid MPSGraphTensorData, or nil if allocation failure.
//
// # Discussion
//
// The device of the MTLBuffer will be used to get the MPSDevice for this
// MPSGraphTensorData.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:shape:dataType:)
func NewGraphTensorDataWithMTLBufferShapeDataType(buffer metal.MTLBuffer, shape foundation.NSArray, dataType uint32) MPSGraphTensorData {
	instance := getMPSGraphTensorDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMTLBuffer:shape:dataType:"), buffer, shape, dataType)
	return MPSGraphTensorDataFromID(rv)
}

// Initializes an tensor data with a metal buffer.
//
// buffer: MTLBuffer to be used within the MPSGraphTensorData
//
// shape: Shape of the output tensor
//
// dataType: dataType of the placeholder tensor
//
// rowBytes: rowBytes for the fastest moving dimension, must be larger than or equal to
// sizeOf(dataType)shape[rank - 1] and must be a multiple of sizeOf(dataType)
//
// # Return Value
//
// A valid MPSGraphTensorData, or nil if allocation failure.
//
// # Discussion
//
// The device of the MTLBuffer will be used to get the MPSDevice for this
// MPSGraphTensorData.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:shape:dataType:rowBytes:)
func NewGraphTensorDataWithMTLBufferShapeDataTypeRowBytes(buffer metal.MTLBuffer, shape foundation.NSArray, dataType uint32, rowBytes uint) MPSGraphTensorData {
	instance := getMPSGraphTensorDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMTLBuffer:shape:dataType:rowBytes:"), buffer, shape, dataType, rowBytes)
	return MPSGraphTensorDataFromID(rv)
}

// Initializes an MPSGraphTensorData with an MTLTensor.
//
// tensor: MTLTensor to be used within the MPSGraphTensorData
//
// # Return Value
//
// A valid MPSGraphTensorData, or nil if allocation failure.
//
// # Discussion
//
// The internal storage of the MTLTensor will be aliased. Requires tensor to
// support MTLTensorUsageMachineLearning.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:)-60j6x
func NewGraphTensorDataWithMTLTensor(tensor metal.MTLTensor) MPSGraphTensorData {
	instance := getMPSGraphTensorDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMTLTensor:"), tensor)
	return MPSGraphTensorDataFromID(rv)
}

// Initializes a tensor data with an MPS matrix.
//
// matrix: MPSMatrix to be used within the MPSGraphTensorData
//
// # Return Value
//
// A valid MPSGraphTensorData, or nil if allocation failure.
//
// # Discussion
//
// The device of the MPSMatrix will be used to get the MPSDevice for this
// MPSGraphTensorData.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:)-2go2
func (g MPSGraphTensorData) InitWithMPSMatrix(matrix *metalperformanceshaders.MPSMatrix) MPSGraphTensorData {
	rv := objc.Send[MPSGraphTensorData](g.ID, objc.Sel("initWithMPSMatrix:"), matrix.ID)
	return rv
}

// Initializes an MPSGraphTensorData with an MPS ndarray.
//
// ndarray: MPSNDArray to be used within the MPSGraphTensorData.
//
// # Return Value
//
// A valid MPSGraphTensorData, or nil if allocation failure.
//
// # Discussion
//
// The device of the MPSNDArray will be used to get the MPSDevice for this
// MPSGraphTensorData.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:)-4bnfb
func (g MPSGraphTensorData) InitWithMPSNDArray(ndarray *metalperformanceshaders.MPSNDArray) MPSGraphTensorData {
	rv := objc.Send[MPSGraphTensorData](g.ID, objc.Sel("initWithMPSNDArray:"), ndarray.ID)
	return rv
}

// Initializes a tensor data with an MPS image batch.
//
// imageBatch: The device on which the kernel will run, unorm8 and unorm16 images will
// create a float32 tensorData
//
// # Return Value
//
// A valid MPSGraphTensorData, or nil if allocation failure.
//
// # Discussion
//
// The dataLayout used will be NHWC, call a transpose or permute to change to
// a layout of your choice.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:)-511a
func (g MPSGraphTensorData) InitWithMPSImageBatch(imageBatch foundation.NSArray) MPSGraphTensorData {
	rv := objc.Send[MPSGraphTensorData](g.ID, objc.Sel("initWithMPSImageBatch:"), imageBatch)
	return rv
}

// Initializes an MPSGraphTensorData with an MTLTensor.
//
// tensor: MTLTensor to be used within the MPSGraphTensorData
//
// # Return Value
//
// A valid MPSGraphTensorData, or nil if allocation failure.
//
// # Discussion
//
// The internal storage of the MTLTensor will be aliased. Requires tensor to
// support MTLTensorUsageMachineLearning.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:)-60j6x
func (g MPSGraphTensorData) InitWithMTLTensor(tensor metal.MTLTensor) MPSGraphTensorData {
	rv := objc.Send[MPSGraphTensorData](g.ID, objc.Sel("initWithMTLTensor:"), tensor)
	return rv
}

// Initializes a tensor data with an MPS vector.
//
// vector: MPSVector to be used within the MPSGraphTensorData
//
// # Return Value
//
// A valid MPSGraphTensorData, or nil if allocation failure.
//
// # Discussion
//
// The device of the MPSVector will be used to get the MPSDevice for this
// MPSGraphTensorData.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:)-9kgoe
func (g MPSGraphTensorData) InitWithMPSVector(vector *metalperformanceshaders.MPSVector) MPSGraphTensorData {
	rv := objc.Send[MPSGraphTensorData](g.ID, objc.Sel("initWithMPSVector:"), vector.ID)
	return rv
}

// Initializes a tensor data with an MPS vector enforcing rank of the result.
//
// vector: MPSVector to be used within the MPSGraphTensorData
//
// rank: The rank of the resulting TensorData tensor. NOTE: must be within { 1, …
// ,16 }.
//
// # Return Value
//
// A valid MPSGraphTensorData, or nil if allocation failure.
//
// # Discussion
//
// The device of the MPSVector will be used to get the MPSDevice for this
// MPSGraphTensorData.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:rank:)-1e4ks
func (g MPSGraphTensorData) InitWithMPSVectorRank(vector *metalperformanceshaders.MPSVector, rank uint) MPSGraphTensorData {
	rv := objc.Send[MPSGraphTensorData](g.ID, objc.Sel("initWithMPSVector:rank:"), vector.ID, rank)
	return rv
}

// Initializes a tensor data with an MPS matrix enforcing rank of the result.
//
// matrix: MPSMatrix to be used within the MPSGraphTensorData
//
// rank: The rank of the resulting TensorData tensor. NOTE: must be within { 1, …
// ,16 }.
//
// # Return Value
//
// A valid MPSGraphTensorData of given rank, or nil if allocation failure.
//
// # Discussion
//
// The device of the MPSMatrix will be used to get the MPSDevice for this
// MPSGraphTensorData.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:rank:)-1lnxg
func (g MPSGraphTensorData) InitWithMPSMatrixRank(matrix *metalperformanceshaders.MPSMatrix, rank uint) MPSGraphTensorData {
	rv := objc.Send[MPSGraphTensorData](g.ID, objc.Sel("initWithMPSMatrix:rank:"), matrix.ID, rank)
	return rv
}

// Initializes an tensor data with a metal buffer.
//
// buffer: MTLBuffer to be used within the MPSGraphTensorData
//
// shape: Shape of the output tensor
//
// dataType: dataType of the placeholder tensor
//
// # Return Value
//
// A valid MPSGraphTensorData, or nil if allocation failure.
//
// # Discussion
//
// The device of the MTLBuffer will be used to get the MPSDevice for this
// MPSGraphTensorData.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:shape:dataType:)
func (g MPSGraphTensorData) InitWithMTLBufferShapeDataType(buffer metal.MTLBuffer, shape foundation.NSArray, dataType uint32) MPSGraphTensorData {
	rv := objc.Send[MPSGraphTensorData](g.ID, objc.Sel("initWithMTLBuffer:shape:dataType:"), buffer, shape, dataType)
	return rv
}

// Initializes an tensor data with a metal buffer.
//
// buffer: MTLBuffer to be used within the MPSGraphTensorData
//
// shape: Shape of the output tensor
//
// dataType: dataType of the placeholder tensor
//
// rowBytes: rowBytes for the fastest moving dimension, must be larger than or equal to
// sizeOf(dataType)shape[rank - 1] and must be a multiple of sizeOf(dataType)
//
// # Return Value
//
// A valid MPSGraphTensorData, or nil if allocation failure.
//
// # Discussion
//
// The device of the MTLBuffer will be used to get the MPSDevice for this
// MPSGraphTensorData.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:shape:dataType:rowBytes:)
func (g MPSGraphTensorData) InitWithMTLBufferShapeDataTypeRowBytes(buffer metal.MTLBuffer, shape foundation.NSArray, dataType uint32, rowBytes uint) MPSGraphTensorData {
	rv := objc.Send[MPSGraphTensorData](g.ID, objc.Sel("initWithMTLBuffer:shape:dataType:rowBytes:"), buffer, shape, dataType, rowBytes)
	return rv
}

// Initializes the tensor data with an [NSData] on a device.
//
// device: MPSDevice on which the MPSGraphTensorData exists
//
// data: NSData from which to copy the contents
//
// shape: Shape of the output tensor
//
// dataType: dataType of the placeholder tensor
//
// # Return Value
//
// A valid MPSGraphTensorData, or nil if allocation failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(device:data:shape:dataType:)
func (g MPSGraphTensorData) InitWithDeviceDataShapeDataType(device IMPSGraphDevice, data foundation.NSData, shape foundation.NSArray, dataType uint32) MPSGraphTensorData {
	rv := objc.Send[MPSGraphTensorData](g.ID, objc.Sel("initWithDevice:data:shape:dataType:"), device, data, shape, dataType)
	return rv
}

// Return an mpsndarray object will copy contents if the contents are not
// stored in an MPS ndarray.
//
// # Return Value
//
// A valid MPSNDArray, or nil if allocation fails.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/mpsndarray()
func (g MPSGraphTensorData) Mpsndarray() metalperformanceshaders.MPSNDArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("mpsndarray"))
	return metalperformanceshaders.MPSNDArrayFromID(rv)
}

// The data type of the tensor data.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/dataType
func (g MPSGraphTensorData) DataType() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("dataType"))
	return rv
}

// The device of the tensor data.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/device
func (g MPSGraphTensorData) Device() IMPSGraphDevice {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("device"))
	return MPSGraphDeviceFromID(objc.ID(rv))
}

// The shape of the tensor data.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/shape
func (g MPSGraphTensorData) Shape() foundation.NSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("shape"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

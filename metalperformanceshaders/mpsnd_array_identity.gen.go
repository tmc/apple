// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNDArrayIdentity] class.
var (
	_MPSNDArrayIdentityClass     MPSNDArrayIdentityClass
	_MPSNDArrayIdentityClassOnce sync.Once
)

func getMPSNDArrayIdentityClass() MPSNDArrayIdentityClass {
	_MPSNDArrayIdentityClassOnce.Do(func() {
		_MPSNDArrayIdentityClass = MPSNDArrayIdentityClass{class: objc.GetClass("MPSNDArrayIdentity")}
	})
	return _MPSNDArrayIdentityClass
}

// GetMPSNDArrayIdentityClass returns the class object for MPSNDArrayIdentity.
func GetMPSNDArrayIdentityClass() MPSNDArrayIdentityClass {
	return getMPSNDArrayIdentityClass()
}

type MPSNDArrayIdentityClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayIdentityClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayIdentityClass) Alloc() MPSNDArrayIdentity {
	rv := objc.Send[MPSNDArrayIdentity](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Methods
//
//   - [MPSNDArrayIdentity.ReshapeWithCommandEncoderCommandBufferSourceArrayDimensionCountDimensionSizesDestinationArray]
//   - [MPSNDArrayIdentity.ReshapeWithCommandEncoderCommandBufferSourceArrayShapeDestinationArray]
//   - [MPSNDArrayIdentity.ReshapeWithCommandBufferSourceArrayDimensionCountDimensionSizesDestinationArray]
//   - [MPSNDArrayIdentity.ReshapeWithCommandBufferSourceArrayShapeDestinationArray]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayIdentity
type MPSNDArrayIdentity struct {
	MPSNDArrayUnaryKernel
}

// MPSNDArrayIdentityFromID constructs a [MPSNDArrayIdentity] from an objc.ID.
func MPSNDArrayIdentityFromID(id objc.ID) MPSNDArrayIdentity {
	return MPSNDArrayIdentity{MPSNDArrayUnaryKernel: MPSNDArrayUnaryKernelFromID(id)}
}

// NOTE: MPSNDArrayIdentity adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayIdentity] class.
//
// # Instance Methods
//
//   - [IMPSNDArrayIdentity.ReshapeWithCommandEncoderCommandBufferSourceArrayDimensionCountDimensionSizesDestinationArray]
//   - [IMPSNDArrayIdentity.ReshapeWithCommandEncoderCommandBufferSourceArrayShapeDestinationArray]
//   - [IMPSNDArrayIdentity.ReshapeWithCommandBufferSourceArrayDimensionCountDimensionSizesDestinationArray]
//   - [IMPSNDArrayIdentity.ReshapeWithCommandBufferSourceArrayShapeDestinationArray]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayIdentity
type IMPSNDArrayIdentity interface {
	IMPSNDArrayUnaryKernel

	// Topic: Instance Methods

	ReshapeWithCommandEncoderCommandBufferSourceArrayDimensionCountDimensionSizesDestinationArray(encoder metal.MTLComputeCommandEncoder, cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray, numberOfDimensions uint, dimensionSizes *uint, destinationArray IMPSNDArray) IMPSNDArray
	ReshapeWithCommandEncoderCommandBufferSourceArrayShapeDestinationArray(encoder metal.MTLComputeCommandEncoder, cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray, shape MPSShape, destinationArray IMPSNDArray) IMPSNDArray
	ReshapeWithCommandBufferSourceArrayDimensionCountDimensionSizesDestinationArray(cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray, numberOfDimensions uint, dimensionSizes *uint, destinationArray IMPSNDArray) IMPSNDArray
	ReshapeWithCommandBufferSourceArrayShapeDestinationArray(cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray, shape MPSShape, destinationArray IMPSNDArray) IMPSNDArray
}

// Init initializes the instance.
func (n MPSNDArrayIdentity) Init() MPSNDArrayIdentity {
	rv := objc.Send[MPSNDArrayIdentity](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayIdentity) Autorelease() MPSNDArrayIdentity {
	rv := objc.Send[MPSNDArrayIdentity](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayIdentity creates a new MPSNDArrayIdentity instance.
func NewMPSNDArrayIdentity() MPSNDArrayIdentity {
	class := getMPSNDArrayIdentityClass()
	rv := objc.Send[MPSNDArrayIdentity](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewNDArrayIdentityWithCoder(aDecoder foundation.INSCoder) MPSNDArrayIdentity {
	instance := getMPSNDArrayIdentityClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNDArrayIdentityFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryKernel/init(coder:device:)
func NewNDArrayIdentityWithCoderDevice(coder foundation.INSCoder, device metal.MTLDevice) MPSNDArrayIdentity {
	instance := getMPSNDArrayIdentityClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNDArrayIdentityFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayIdentity/init(device:)
func NewNDArrayIdentityWithDevice(device metal.MTLDevice) MPSNDArrayIdentity {
	instance := getMPSNDArrayIdentityClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNDArrayIdentityFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/init(device:sourceCount:)
func NewNDArrayIdentityWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayIdentity {
	instance := getMPSNDArrayIdentityClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	return MPSNDArrayIdentityFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayIdentity/reshape(with:commandBuffer:sourceArray:dimensionCount:dimensionSizes:destinationArray:)
func (n MPSNDArrayIdentity) ReshapeWithCommandEncoderCommandBufferSourceArrayDimensionCountDimensionSizesDestinationArray(encoder metal.MTLComputeCommandEncoder, cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray, numberOfDimensions uint, dimensionSizes *uint, destinationArray IMPSNDArray) IMPSNDArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("reshapeWithCommandEncoder:commandBuffer:sourceArray:dimensionCount:dimensionSizes:destinationArray:"), encoder, cmdBuf, sourceArray, numberOfDimensions, unsafe.Pointer(dimensionSizes), destinationArray)
	return MPSNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayIdentity/reshape(with:commandBuffer:sourceArray:shape:destinationArray:)
func (n MPSNDArrayIdentity) ReshapeWithCommandEncoderCommandBufferSourceArrayShapeDestinationArray(encoder metal.MTLComputeCommandEncoder, cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray, shape MPSShape, destinationArray IMPSNDArray) IMPSNDArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("reshapeWithCommandEncoder:commandBuffer:sourceArray:shape:destinationArray:"), encoder, cmdBuf, sourceArray, shape, destinationArray)
	return MPSNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayIdentity/reshape(with:sourceArray:dimensionCount:dimensionSizes:destinationArray:)
func (n MPSNDArrayIdentity) ReshapeWithCommandBufferSourceArrayDimensionCountDimensionSizesDestinationArray(cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray, numberOfDimensions uint, dimensionSizes *uint, destinationArray IMPSNDArray) IMPSNDArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("reshapeWithCommandBuffer:sourceArray:dimensionCount:dimensionSizes:destinationArray:"), cmdBuf, sourceArray, numberOfDimensions, unsafe.Pointer(dimensionSizes), destinationArray)
	return MPSNDArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayIdentity/reshape(with:sourceArray:shape:destinationArray:)
func (n MPSNDArrayIdentity) ReshapeWithCommandBufferSourceArrayShapeDestinationArray(cmdBuf metal.MTLCommandBuffer, sourceArray IMPSNDArray, shape MPSShape, destinationArray IMPSNDArray) IMPSNDArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("reshapeWithCommandBuffer:sourceArray:shape:destinationArray:"), cmdBuf, sourceArray, shape, destinationArray)
	return MPSNDArrayFromID(rv)
}

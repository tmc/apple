// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixRandom] class.
var (
	_MPSMatrixRandomClass     MPSMatrixRandomClass
	_MPSMatrixRandomClassOnce sync.Once
)

func getMPSMatrixRandomClass() MPSMatrixRandomClass {
	_MPSMatrixRandomClassOnce.Do(func() {
		_MPSMatrixRandomClass = MPSMatrixRandomClass{class: objc.GetClass("MPSMatrixRandom")}
	})
	return _MPSMatrixRandomClass
}

// GetMPSMatrixRandomClass returns the class object for MPSMatrixRandom.
func GetMPSMatrixRandomClass() MPSMatrixRandomClass {
	return getMPSMatrixRandomClass()
}

type MPSMatrixRandomClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixRandomClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixRandomClass) Alloc() MPSMatrixRandom {
	rv := objc.Send[MPSMatrixRandom](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MPSMatrixRandom.BatchSize]
//   - [MPSMatrixRandom.SetBatchSize]
//   - [MPSMatrixRandom.BatchStart]
//   - [MPSMatrixRandom.SetBatchStart]
//   - [MPSMatrixRandom.DestinationDataType]
//   - [MPSMatrixRandom.DistributionType]
//
// # Instance Methods
//
//   - [MPSMatrixRandom.EncodeToCommandBufferDestinationMatrix]
//   - [MPSMatrixRandom.EncodeToCommandBufferDestinationVector]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandom
type MPSMatrixRandom struct {
	MPSKernel
}

// MPSMatrixRandomFromID constructs a [MPSMatrixRandom] from an objc.ID.
func MPSMatrixRandomFromID(id objc.ID) MPSMatrixRandom {
	return MPSMatrixRandom{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSMatrixRandom adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixRandom] class.
//
// # Instance Properties
//
//   - [IMPSMatrixRandom.BatchSize]
//   - [IMPSMatrixRandom.SetBatchSize]
//   - [IMPSMatrixRandom.BatchStart]
//   - [IMPSMatrixRandom.SetBatchStart]
//   - [IMPSMatrixRandom.DestinationDataType]
//   - [IMPSMatrixRandom.DistributionType]
//
// # Instance Methods
//
//   - [IMPSMatrixRandom.EncodeToCommandBufferDestinationMatrix]
//   - [IMPSMatrixRandom.EncodeToCommandBufferDestinationVector]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandom
type IMPSMatrixRandom interface {
	IMPSKernel

	// Topic: Instance Properties

	BatchSize() uint
	SetBatchSize(value uint)
	BatchStart() uint
	SetBatchStart(value uint)
	DestinationDataType() MPSDataType
	DistributionType() MPSMatrixRandomDistribution

	// Topic: Instance Methods

	EncodeToCommandBufferDestinationMatrix(commandBuffer metal.MTLCommandBuffer, destinationMatrix IMPSMatrix)
	EncodeToCommandBufferDestinationVector(commandBuffer metal.MTLCommandBuffer, destinationVector IMPSVector)
}

// Init initializes the instance.
func (m MPSMatrixRandom) Init() MPSMatrixRandom {
	rv := objc.Send[MPSMatrixRandom](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixRandom) Autorelease() MPSMatrixRandom {
	rv := objc.Send[MPSMatrixRandom](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixRandom creates a new MPSMatrixRandom instance.
func NewMPSMatrixRandom() MPSMatrixRandom {
	class := getMPSMatrixRandomClass()
	rv := objc.Send[MPSMatrixRandom](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixRandomWithCoder(aDecoder foundation.INSCoder) MPSMatrixRandom {
	instance := getMPSMatrixRandomClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixRandomFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:device:)
func NewMatrixRandomWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixRandom {
	instance := getMPSMatrixRandomClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixRandomFromID(rv)
}

// Initializes a new kernel object.
//
// device: The Metal device on which the kernel will be used.
//
// # Return Value
//
// An initialized kernel object.
//
// # Discussion
//
// This method fails if the device is not supported. Query the
// [MPSSupportsMTLDevice] function to determine whether the device is
// supported.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(device:)
func NewMatrixRandomWithDevice(device metal.MTLDevice) MPSMatrixRandom {
	instance := getMPSMatrixRandomClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixRandomFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandom/encode(commandBuffer:destinationMatrix:)
func (m MPSMatrixRandom) EncodeToCommandBufferDestinationMatrix(commandBuffer metal.MTLCommandBuffer, destinationMatrix IMPSMatrix) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:destinationMatrix:"), commandBuffer, destinationMatrix)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandom/encode(commandBuffer:destinationVector:)
func (m MPSMatrixRandom) EncodeToCommandBufferDestinationVector(commandBuffer metal.MTLCommandBuffer, destinationVector IMPSVector) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeToCommandBuffer:destinationVector:"), commandBuffer, destinationVector)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandom/batchSize
func (m MPSMatrixRandom) BatchSize() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("batchSize"))
	return rv
}
func (m MPSMatrixRandom) SetBatchSize(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setBatchSize:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandom/batchStart
func (m MPSMatrixRandom) BatchStart() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("batchStart"))
	return rv
}
func (m MPSMatrixRandom) SetBatchStart(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setBatchStart:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandom/destinationDataType
func (m MPSMatrixRandom) DestinationDataType() MPSDataType {
	rv := objc.Send[MPSDataType](m.ID, objc.Sel("destinationDataType"))
	return MPSDataType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandom/distributionType
func (m MPSMatrixRandom) DistributionType() MPSMatrixRandomDistribution {
	rv := objc.Send[MPSMatrixRandomDistribution](m.ID, objc.Sel("distributionType"))
	return MPSMatrixRandomDistribution(rv)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNBatchNormalizationStatistics] class.
var (
	_MPSCNNBatchNormalizationStatisticsClass     MPSCNNBatchNormalizationStatisticsClass
	_MPSCNNBatchNormalizationStatisticsClassOnce sync.Once
)

func getMPSCNNBatchNormalizationStatisticsClass() MPSCNNBatchNormalizationStatisticsClass {
	_MPSCNNBatchNormalizationStatisticsClassOnce.Do(func() {
		_MPSCNNBatchNormalizationStatisticsClass = MPSCNNBatchNormalizationStatisticsClass{class: objc.GetClass("MPSCNNBatchNormalizationStatistics")}
	})
	return _MPSCNNBatchNormalizationStatisticsClass
}

// GetMPSCNNBatchNormalizationStatisticsClass returns the class object for MPSCNNBatchNormalizationStatistics.
func GetMPSCNNBatchNormalizationStatisticsClass() MPSCNNBatchNormalizationStatisticsClass {
	return getMPSCNNBatchNormalizationStatisticsClass()
}

type MPSCNNBatchNormalizationStatisticsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNBatchNormalizationStatisticsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNBatchNormalizationStatisticsClass) Alloc() MPSCNNBatchNormalizationStatistics {
	rv := objc.Send[MPSCNNBatchNormalizationStatistics](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that stores statistics required to execute batch normalization.
//
// # Instance Methods
//
//   - [MPSCNNBatchNormalizationStatistics.EncodeBatchToCommandBufferSourceImagesBatchNormalizationState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationStatistics
type MPSCNNBatchNormalizationStatistics struct {
	MPSCNNKernel
}

// MPSCNNBatchNormalizationStatisticsFromID constructs a [MPSCNNBatchNormalizationStatistics] from an objc.ID.
//
// An object that stores statistics required to execute batch normalization.
func MPSCNNBatchNormalizationStatisticsFromID(id objc.ID) MPSCNNBatchNormalizationStatistics {
	return MPSCNNBatchNormalizationStatistics{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSCNNBatchNormalizationStatistics adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNBatchNormalizationStatistics] class.
//
// # Instance Methods
//
//   - [IMPSCNNBatchNormalizationStatistics.EncodeBatchToCommandBufferSourceImagesBatchNormalizationState]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationStatistics
type IMPSCNNBatchNormalizationStatistics interface {
	IMPSCNNKernel

	// Topic: Instance Methods

	EncodeBatchToCommandBufferSourceImagesBatchNormalizationState(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, batchNormalizationState IMPSCNNBatchNormalizationState)
}

// Init initializes the instance.
func (c MPSCNNBatchNormalizationStatistics) Init() MPSCNNBatchNormalizationStatistics {
	rv := objc.Send[MPSCNNBatchNormalizationStatistics](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNBatchNormalizationStatistics) Autorelease() MPSCNNBatchNormalizationStatistics {
	rv := objc.Send[MPSCNNBatchNormalizationStatistics](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNBatchNormalizationStatistics creates a new MPSCNNBatchNormalizationStatistics instance.
func NewMPSCNNBatchNormalizationStatistics() MPSCNNBatchNormalizationStatistics {
	class := getMPSCNNBatchNormalizationStatisticsClass()
	rv := objc.Send[MPSCNNBatchNormalizationStatistics](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNBatchNormalizationStatisticsWithCoder(aDecoder foundation.INSCoder) MPSCNNBatchNormalizationStatistics {
	instance := getMPSCNNBatchNormalizationStatisticsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNBatchNormalizationStatisticsFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationStatistics/init(coder:device:)
func NewCNNBatchNormalizationStatisticsWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNBatchNormalizationStatistics {
	instance := getMPSCNNBatchNormalizationStatisticsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNBatchNormalizationStatisticsFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationStatistics/init(device:)
func NewCNNBatchNormalizationStatisticsWithDevice(device metal.MTLDevice) MPSCNNBatchNormalizationStatistics {
	instance := getMPSCNNBatchNormalizationStatisticsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNBatchNormalizationStatisticsFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationStatistics/encodeBatch(to:sourceImages:batchNormalizationState:)
func (c MPSCNNBatchNormalizationStatistics) EncodeBatchToCommandBufferSourceImagesBatchNormalizationState(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, batchNormalizationState IMPSCNNBatchNormalizationState) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:batchNormalizationState:"), commandBuffer, sourceImages, batchNormalizationState)
}

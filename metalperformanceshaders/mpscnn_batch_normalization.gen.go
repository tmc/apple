// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNBatchNormalization] class.
var (
	_MPSCNNBatchNormalizationClass     MPSCNNBatchNormalizationClass
	_MPSCNNBatchNormalizationClassOnce sync.Once
)

func getMPSCNNBatchNormalizationClass() MPSCNNBatchNormalizationClass {
	_MPSCNNBatchNormalizationClassOnce.Do(func() {
		_MPSCNNBatchNormalizationClass = MPSCNNBatchNormalizationClass{class: objc.GetClass("MPSCNNBatchNormalization")}
	})
	return _MPSCNNBatchNormalizationClass
}

// GetMPSCNNBatchNormalizationClass returns the class object for MPSCNNBatchNormalization.
func GetMPSCNNBatchNormalizationClass() MPSCNNBatchNormalizationClass {
	return getMPSCNNBatchNormalizationClass()
}

type MPSCNNBatchNormalizationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNBatchNormalizationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNBatchNormalizationClass) Alloc() MPSCNNBatchNormalization {
	rv := objc.Send[MPSCNNBatchNormalization](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A batch normalization kernel.
//
// # Initializers
//
//   - [MPSCNNBatchNormalization.InitWithDeviceDataSource]
//   - [MPSCNNBatchNormalization.InitWithDeviceDataSourceFusedNeuronDescriptor]
//
// # Instance Properties
//
//   - [MPSCNNBatchNormalization.DataSource]
//   - [MPSCNNBatchNormalization.Epsilon]
//   - [MPSCNNBatchNormalization.SetEpsilon]
//   - [MPSCNNBatchNormalization.NumberOfFeatureChannels]
//
// # Instance Methods
//
//   - [MPSCNNBatchNormalization.EncodeToCommandBufferSourceImageBatchNormalizationStateDestinationImage]
//   - [MPSCNNBatchNormalization.EncodeBatchToCommandBufferSourceImagesBatchNormalizationStateDestinationImages]
//   - [MPSCNNBatchNormalization.ReloadGammaAndBetaWithCommandBufferGammaAndBetaState]
//   - [MPSCNNBatchNormalization.ReloadGammaAndBetaFromDataSource]
//   - [MPSCNNBatchNormalization.ReloadMeanAndVarianceWithCommandBufferMeanAndVarianceState]
//   - [MPSCNNBatchNormalization.ReloadMeanAndVarianceFromDataSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalization
type MPSCNNBatchNormalization struct {
	MPSCNNKernel
}

// MPSCNNBatchNormalizationFromID constructs a [MPSCNNBatchNormalization] from an objc.ID.
//
// A batch normalization kernel.
func MPSCNNBatchNormalizationFromID(id objc.ID) MPSCNNBatchNormalization {
	return MPSCNNBatchNormalization{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSCNNBatchNormalization adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNBatchNormalization] class.
//
// # Initializers
//
//   - [IMPSCNNBatchNormalization.InitWithDeviceDataSource]
//   - [IMPSCNNBatchNormalization.InitWithDeviceDataSourceFusedNeuronDescriptor]
//
// # Instance Properties
//
//   - [IMPSCNNBatchNormalization.DataSource]
//   - [IMPSCNNBatchNormalization.Epsilon]
//   - [IMPSCNNBatchNormalization.SetEpsilon]
//   - [IMPSCNNBatchNormalization.NumberOfFeatureChannels]
//
// # Instance Methods
//
//   - [IMPSCNNBatchNormalization.EncodeToCommandBufferSourceImageBatchNormalizationStateDestinationImage]
//   - [IMPSCNNBatchNormalization.EncodeBatchToCommandBufferSourceImagesBatchNormalizationStateDestinationImages]
//   - [IMPSCNNBatchNormalization.ReloadGammaAndBetaWithCommandBufferGammaAndBetaState]
//   - [IMPSCNNBatchNormalization.ReloadGammaAndBetaFromDataSource]
//   - [IMPSCNNBatchNormalization.ReloadMeanAndVarianceWithCommandBufferMeanAndVarianceState]
//   - [IMPSCNNBatchNormalization.ReloadMeanAndVarianceFromDataSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalization
type IMPSCNNBatchNormalization interface {
	IMPSCNNKernel

	// Topic: Initializers

	InitWithDeviceDataSource(device metal.MTLDevice, dataSource MPSCNNBatchNormalizationDataSource) MPSCNNBatchNormalization
	InitWithDeviceDataSourceFusedNeuronDescriptor(device metal.MTLDevice, dataSource MPSCNNBatchNormalizationDataSource, fusedNeuronDescriptor IMPSNNNeuronDescriptor) MPSCNNBatchNormalization

	// Topic: Instance Properties

	DataSource() MPSCNNBatchNormalizationDataSource
	Epsilon() float32
	SetEpsilon(value float32)
	NumberOfFeatureChannels() uint

	// Topic: Instance Methods

	EncodeToCommandBufferSourceImageBatchNormalizationStateDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, batchNormalizationState IMPSCNNBatchNormalizationState, destinationImage IMPSImage)
	EncodeBatchToCommandBufferSourceImagesBatchNormalizationStateDestinationImages(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, batchNormalizationState IMPSCNNBatchNormalizationState, destinationImages MPSImageBatch)
	ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer metal.MTLCommandBuffer, gammaAndBetaState IMPSCNNNormalizationGammaAndBetaState)
	ReloadGammaAndBetaFromDataSource()
	ReloadMeanAndVarianceWithCommandBufferMeanAndVarianceState(commandBuffer metal.MTLCommandBuffer, meanAndVarianceState IMPSCNNNormalizationMeanAndVarianceState)
	ReloadMeanAndVarianceFromDataSource()
}

// Init initializes the instance.
func (c MPSCNNBatchNormalization) Init() MPSCNNBatchNormalization {
	rv := objc.Send[MPSCNNBatchNormalization](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNBatchNormalization) Autorelease() MPSCNNBatchNormalization {
	rv := objc.Send[MPSCNNBatchNormalization](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNBatchNormalization creates a new MPSCNNBatchNormalization instance.
func NewMPSCNNBatchNormalization() MPSCNNBatchNormalization {
	class := getMPSCNNBatchNormalizationClass()
	rv := objc.Send[MPSCNNBatchNormalization](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNBatchNormalizationWithCoder(aDecoder foundation.INSCoder) MPSCNNBatchNormalization {
	instance := getMPSCNNBatchNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNBatchNormalizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalization/init(coder:device:)
func NewCNNBatchNormalizationWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNBatchNormalization {
	instance := getMPSCNNBatchNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNBatchNormalizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNBatchNormalizationWithDevice(device metal.MTLDevice) MPSCNNBatchNormalization {
	instance := getMPSCNNBatchNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNBatchNormalizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalization/init(device:dataSource:)
func NewCNNBatchNormalizationWithDeviceDataSource(device metal.MTLDevice, dataSource MPSCNNBatchNormalizationDataSource) MPSCNNBatchNormalization {
	instance := getMPSCNNBatchNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:dataSource:"), device, dataSource)
	return MPSCNNBatchNormalizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalization/init(device:dataSource:fusedNeuronDescriptor:)
func NewCNNBatchNormalizationWithDeviceDataSourceFusedNeuronDescriptor(device metal.MTLDevice, dataSource MPSCNNBatchNormalizationDataSource, fusedNeuronDescriptor IMPSNNNeuronDescriptor) MPSCNNBatchNormalization {
	instance := getMPSCNNBatchNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:dataSource:fusedNeuronDescriptor:"), device, dataSource, fusedNeuronDescriptor)
	return MPSCNNBatchNormalizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalization/init(device:dataSource:)
func (c MPSCNNBatchNormalization) InitWithDeviceDataSource(device metal.MTLDevice, dataSource MPSCNNBatchNormalizationDataSource) MPSCNNBatchNormalization {
	rv := objc.Send[MPSCNNBatchNormalization](c.ID, objc.Sel("initWithDevice:dataSource:"), device, dataSource)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalization/init(device:dataSource:fusedNeuronDescriptor:)
func (c MPSCNNBatchNormalization) InitWithDeviceDataSourceFusedNeuronDescriptor(device metal.MTLDevice, dataSource MPSCNNBatchNormalizationDataSource, fusedNeuronDescriptor IMPSNNNeuronDescriptor) MPSCNNBatchNormalization {
	rv := objc.Send[MPSCNNBatchNormalization](c.ID, objc.Sel("initWithDevice:dataSource:fusedNeuronDescriptor:"), device, dataSource, fusedNeuronDescriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalization/encode(to:sourceImage:batchNormalizationState:destinationImage:)
func (c MPSCNNBatchNormalization) EncodeToCommandBufferSourceImageBatchNormalizationStateDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, batchNormalizationState IMPSCNNBatchNormalizationState, destinationImage IMPSImage) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceImage:batchNormalizationState:destinationImage:"), commandBuffer, sourceImage, batchNormalizationState, destinationImage)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalization/encodeBatch(to:sourceImages:batchNormalizationState:destinationImages:)
func (c MPSCNNBatchNormalization) EncodeBatchToCommandBufferSourceImagesBatchNormalizationStateDestinationImages(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, batchNormalizationState IMPSCNNBatchNormalizationState, destinationImages MPSImageBatch) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:batchNormalizationState:destinationImages:"), commandBuffer, sourceImages, batchNormalizationState, destinationImages)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalization/reloadGammaAndBeta(with:gammaAndBetaState:)
func (c MPSCNNBatchNormalization) ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer metal.MTLCommandBuffer, gammaAndBetaState IMPSCNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadGammaAndBetaWithCommandBuffer:gammaAndBetaState:"), commandBuffer, gammaAndBetaState)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalization/reloadGammaAndBetaFromDataSource()
func (c MPSCNNBatchNormalization) ReloadGammaAndBetaFromDataSource() {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadGammaAndBetaFromDataSource"))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalization/reloadMeanAndVariance(with:meanAndVarianceState:)
func (c MPSCNNBatchNormalization) ReloadMeanAndVarianceWithCommandBufferMeanAndVarianceState(commandBuffer metal.MTLCommandBuffer, meanAndVarianceState IMPSCNNNormalizationMeanAndVarianceState) {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadMeanAndVarianceWithCommandBuffer:meanAndVarianceState:"), commandBuffer, meanAndVarianceState)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalization/reloadMeanAndVarianceFromDataSource()
func (c MPSCNNBatchNormalization) ReloadMeanAndVarianceFromDataSource() {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadMeanAndVarianceFromDataSource"))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalization/dataSource
func (c MPSCNNBatchNormalization) DataSource() MPSCNNBatchNormalizationDataSource {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dataSource"))
	return MPSCNNBatchNormalizationDataSourceObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalization/epsilon
func (c MPSCNNBatchNormalization) Epsilon() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("epsilon"))
	return rv
}
func (c MPSCNNBatchNormalization) SetEpsilon(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setEpsilon:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalization/numberOfFeatureChannels
func (c MPSCNNBatchNormalization) NumberOfFeatureChannels() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("numberOfFeatureChannels"))
	return rv
}

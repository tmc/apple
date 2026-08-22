// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNConvolutionTranspose] class.
var (
	_MPSCNNConvolutionTransposeClass     MPSCNNConvolutionTransposeClass
	_MPSCNNConvolutionTransposeClassOnce sync.Once
)

func getMPSCNNConvolutionTransposeClass() MPSCNNConvolutionTransposeClass {
	_MPSCNNConvolutionTransposeClassOnce.Do(func() {
		_MPSCNNConvolutionTransposeClass = MPSCNNConvolutionTransposeClass{class: objc.GetClass("MPSCNNConvolutionTranspose")}
	})
	return _MPSCNNConvolutionTransposeClass
}

// GetMPSCNNConvolutionTransposeClass returns the class object for MPSCNNConvolutionTranspose.
func GetMPSCNNConvolutionTransposeClass() MPSCNNConvolutionTransposeClass {
	return getMPSCNNConvolutionTransposeClass()
}

type MPSCNNConvolutionTransposeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNConvolutionTransposeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNConvolutionTransposeClass) Alloc() MPSCNNConvolutionTranspose {
	rv := objc.Send[MPSCNNConvolutionTranspose](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A transposed convolution kernel.
//
// # Initializers
//
//   - [MPSCNNConvolutionTranspose.InitWithDeviceWeights]: Initializes a transposed convolution kernel.
//
// # Instance Properties
//
//   - [MPSCNNConvolutionTranspose.Groups]
//   - [MPSCNNConvolutionTranspose.InputFeatureChannels]
//   - [MPSCNNConvolutionTranspose.KernelOffsetX]
//   - [MPSCNNConvolutionTranspose.SetKernelOffsetX]
//   - [MPSCNNConvolutionTranspose.KernelOffsetY]
//   - [MPSCNNConvolutionTranspose.SetKernelOffsetY]
//   - [MPSCNNConvolutionTranspose.OutputFeatureChannels]
//   - [MPSCNNConvolutionTranspose.AccumulatorPrecisionOption]
//   - [MPSCNNConvolutionTranspose.SetAccumulatorPrecisionOption]
//   - [MPSCNNConvolutionTranspose.DataSource]
//
// # Instance Methods
//
//   - [MPSCNNConvolutionTranspose.EncodeToCommandBufferSourceImageConvolutionGradientState]
//   - [MPSCNNConvolutionTranspose.EncodeToCommandBufferSourceImageConvolutionGradientStateDestinationImage]
//   - [MPSCNNConvolutionTranspose.EncodeToCommandBufferSourceImageConvolutionGradientStateDestinationStateDestinationStateIsTemporary]
//   - [MPSCNNConvolutionTranspose.EncodeBatchToCommandBufferSourceImagesConvolutionGradientStates]
//   - [MPSCNNConvolutionTranspose.EncodeBatchToCommandBufferSourceImagesConvolutionGradientStatesDestinationImages]
//   - [MPSCNNConvolutionTranspose.EncodeBatchToCommandBufferSourceImagesConvolutionGradientStatesDestinationStatesDestinationStateIsTemporary]
//   - [MPSCNNConvolutionTranspose.ExportWeightsAndBiasesWithCommandBufferResultStateCanBeTemporary]
//   - [MPSCNNConvolutionTranspose.ReloadWeightsAndBiasesWithCommandBufferState]
//   - [MPSCNNConvolutionTranspose.ReloadWeightsAndBiasesFromDataSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose
type MPSCNNConvolutionTranspose struct {
	MPSCNNKernel
}

// MPSCNNConvolutionTransposeFromID constructs a [MPSCNNConvolutionTranspose] from an objc.ID.
//
// A transposed convolution kernel.
func MPSCNNConvolutionTransposeFromID(id objc.ID) MPSCNNConvolutionTranspose {
	return MPSCNNConvolutionTranspose{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSCNNConvolutionTranspose adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNConvolutionTranspose] class.
//
// # Initializers
//
//   - [IMPSCNNConvolutionTranspose.InitWithDeviceWeights]: Initializes a transposed convolution kernel.
//
// # Instance Properties
//
//   - [IMPSCNNConvolutionTranspose.Groups]
//   - [IMPSCNNConvolutionTranspose.InputFeatureChannels]
//   - [IMPSCNNConvolutionTranspose.KernelOffsetX]
//   - [IMPSCNNConvolutionTranspose.SetKernelOffsetX]
//   - [IMPSCNNConvolutionTranspose.KernelOffsetY]
//   - [IMPSCNNConvolutionTranspose.SetKernelOffsetY]
//   - [IMPSCNNConvolutionTranspose.OutputFeatureChannels]
//   - [IMPSCNNConvolutionTranspose.AccumulatorPrecisionOption]
//   - [IMPSCNNConvolutionTranspose.SetAccumulatorPrecisionOption]
//   - [IMPSCNNConvolutionTranspose.DataSource]
//
// # Instance Methods
//
//   - [IMPSCNNConvolutionTranspose.EncodeToCommandBufferSourceImageConvolutionGradientState]
//   - [IMPSCNNConvolutionTranspose.EncodeToCommandBufferSourceImageConvolutionGradientStateDestinationImage]
//   - [IMPSCNNConvolutionTranspose.EncodeToCommandBufferSourceImageConvolutionGradientStateDestinationStateDestinationStateIsTemporary]
//   - [IMPSCNNConvolutionTranspose.EncodeBatchToCommandBufferSourceImagesConvolutionGradientStates]
//   - [IMPSCNNConvolutionTranspose.EncodeBatchToCommandBufferSourceImagesConvolutionGradientStatesDestinationImages]
//   - [IMPSCNNConvolutionTranspose.EncodeBatchToCommandBufferSourceImagesConvolutionGradientStatesDestinationStatesDestinationStateIsTemporary]
//   - [IMPSCNNConvolutionTranspose.ExportWeightsAndBiasesWithCommandBufferResultStateCanBeTemporary]
//   - [IMPSCNNConvolutionTranspose.ReloadWeightsAndBiasesWithCommandBufferState]
//   - [IMPSCNNConvolutionTranspose.ReloadWeightsAndBiasesFromDataSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose
type IMPSCNNConvolutionTranspose interface {
	IMPSCNNKernel

	// Topic: Initializers

	// Initializes a transposed convolution kernel.
	InitWithDeviceWeights(device metal.MTLDevice, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionTranspose

	// Topic: Instance Properties

	Groups() uint
	InputFeatureChannels() uint
	KernelOffsetX() int
	SetKernelOffsetX(value int)
	KernelOffsetY() int
	SetKernelOffsetY(value int)
	OutputFeatureChannels() uint
	AccumulatorPrecisionOption() MPSNNConvolutionAccumulatorPrecisionOption
	SetAccumulatorPrecisionOption(value MPSNNConvolutionAccumulatorPrecisionOption)
	DataSource() MPSCNNConvolutionDataSource

	// Topic: Instance Methods

	EncodeToCommandBufferSourceImageConvolutionGradientState(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, convolutionGradientState IMPSCNNConvolutionGradientState) IMPSImage
	EncodeToCommandBufferSourceImageConvolutionGradientStateDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, convolutionGradientState IMPSCNNConvolutionGradientState, destinationImage IMPSImage)
	EncodeToCommandBufferSourceImageConvolutionGradientStateDestinationStateDestinationStateIsTemporary(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, convolutionGradientState IMPSCNNConvolutionGradientState, outState IMPSCNNConvolutionTransposeGradientState, isTemporary bool) IMPSImage
	EncodeBatchToCommandBufferSourceImagesConvolutionGradientStates(commandBuffer metal.MTLCommandBuffer, sourceImage MPSImageBatch, convolutionGradientState MPSCNNConvolutionGradientStateBatch) MPSImageBatch
	EncodeBatchToCommandBufferSourceImagesConvolutionGradientStatesDestinationImages(commandBuffer metal.MTLCommandBuffer, sourceImage MPSImageBatch, convolutionGradientState MPSCNNConvolutionGradientStateBatch, destinationImage MPSImageBatch)
	EncodeBatchToCommandBufferSourceImagesConvolutionGradientStatesDestinationStatesDestinationStateIsTemporary(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, convolutionGradientStates MPSCNNConvolutionGradientStateBatch, outStates MPSCNNConvolutionTransposeGradientStateBatch, isTemporary bool) MPSImageBatch
	ExportWeightsAndBiasesWithCommandBufferResultStateCanBeTemporary(commandBuffer metal.MTLCommandBuffer, resultStateCanBeTemporary bool) IMPSCNNConvolutionWeightsAndBiasesState
	ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer metal.MTLCommandBuffer, state IMPSCNNConvolutionWeightsAndBiasesState)
	ReloadWeightsAndBiasesFromDataSource()
}

// Init initializes the instance.
func (c MPSCNNConvolutionTranspose) Init() MPSCNNConvolutionTranspose {
	rv := objc.Send[MPSCNNConvolutionTranspose](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNConvolutionTranspose) Autorelease() MPSCNNConvolutionTranspose {
	rv := objc.Send[MPSCNNConvolutionTranspose](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNConvolutionTranspose creates a new MPSCNNConvolutionTranspose instance.
func NewMPSCNNConvolutionTranspose() MPSCNNConvolutionTranspose {
	class := getMPSCNNConvolutionTransposeClass()
	rv := objc.Send[MPSCNNConvolutionTranspose](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNConvolutionTransposeWithCoder(aDecoder foundation.INSCoder) MPSCNNConvolutionTranspose {
	instance := getMPSCNNConvolutionTransposeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNConvolutionTransposeFromID(rv)
}

// Initializes a transposed convolution kernel.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose/init(coder:device:)
func NewCNNConvolutionTransposeWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNConvolutionTranspose {
	instance := getMPSCNNConvolutionTransposeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNConvolutionTransposeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNConvolutionTransposeWithDevice(device metal.MTLDevice) MPSCNNConvolutionTranspose {
	instance := getMPSCNNConvolutionTransposeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNConvolutionTransposeFromID(rv)
}

// Initializes a transposed convolution kernel.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose/init(device:weights:)
func NewCNNConvolutionTransposeWithDeviceWeights(device metal.MTLDevice, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionTranspose {
	instance := getMPSCNNConvolutionTransposeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:weights:"), device, weights)
	return MPSCNNConvolutionTransposeFromID(rv)
}

// Initializes a transposed convolution kernel.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose/init(device:weights:)
func (c MPSCNNConvolutionTranspose) InitWithDeviceWeights(device metal.MTLDevice, weights MPSCNNConvolutionDataSource) MPSCNNConvolutionTranspose {
	rv := objc.Send[MPSCNNConvolutionTranspose](c.ID, objc.Sel("initWithDevice:weights:"), device, weights)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose/encode(commandBuffer:sourceImage:convolutionGradientState:)
func (c MPSCNNConvolutionTranspose) EncodeToCommandBufferSourceImageConvolutionGradientState(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, convolutionGradientState IMPSCNNConvolutionGradientState) IMPSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceImage:convolutionGradientState:"), commandBuffer, sourceImage, convolutionGradientState)
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose/encode(commandBuffer:sourceImage:convolutionGradientState:destinationImage:)
func (c MPSCNNConvolutionTranspose) EncodeToCommandBufferSourceImageConvolutionGradientStateDestinationImage(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, convolutionGradientState IMPSCNNConvolutionGradientState, destinationImage IMPSImage) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceImage:convolutionGradientState:destinationImage:"), commandBuffer, sourceImage, convolutionGradientState, destinationImage)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose/encode(commandBuffer:sourceImage:convolutionGradientState:destinationState:destinationStateIsTemporary:)
func (c MPSCNNConvolutionTranspose) EncodeToCommandBufferSourceImageConvolutionGradientStateDestinationStateDestinationStateIsTemporary(commandBuffer metal.MTLCommandBuffer, sourceImage IMPSImage, convolutionGradientState IMPSCNNConvolutionGradientState, outState IMPSCNNConvolutionTransposeGradientState, isTemporary bool) IMPSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("encodeToCommandBuffer:sourceImage:convolutionGradientState:destinationState:destinationStateIsTemporary:"), commandBuffer, sourceImage, convolutionGradientState, outState, isTemporary)
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose/encodeBatch(commandBuffer:sourceImages:convolutionGradientStates:)
func (c MPSCNNConvolutionTranspose) EncodeBatchToCommandBufferSourceImagesConvolutionGradientStates(commandBuffer metal.MTLCommandBuffer, sourceImage MPSImageBatch, convolutionGradientState MPSCNNConvolutionGradientStateBatch) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:convolutionGradientStates:"), commandBuffer, sourceImage, convolutionGradientState)
	return MPSImageBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose/encodeBatch(commandBuffer:sourceImages:convolutionGradientStates:destinationImages:)
func (c MPSCNNConvolutionTranspose) EncodeBatchToCommandBufferSourceImagesConvolutionGradientStatesDestinationImages(commandBuffer metal.MTLCommandBuffer, sourceImage MPSImageBatch, convolutionGradientState MPSCNNConvolutionGradientStateBatch, destinationImage MPSImageBatch) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:convolutionGradientStates:destinationImages:"), commandBuffer, sourceImage, convolutionGradientState, destinationImage)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose/encodeBatch(commandBuffer:sourceImages:convolutionGradientStates:destinationStates:destinationStateIsTemporary:)
func (c MPSCNNConvolutionTranspose) EncodeBatchToCommandBufferSourceImagesConvolutionGradientStatesDestinationStatesDestinationStateIsTemporary(commandBuffer metal.MTLCommandBuffer, sourceImages MPSImageBatch, convolutionGradientStates MPSCNNConvolutionGradientStateBatch, outStates MPSCNNConvolutionTransposeGradientStateBatch, isTemporary bool) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](c.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:convolutionGradientStates:destinationStates:destinationStateIsTemporary:"), commandBuffer, sourceImages, convolutionGradientStates, outStates, isTemporary)
	return MPSImageBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose/exportWeightsAndBiases(with:resultStateCanBeTemporary:)
func (c MPSCNNConvolutionTranspose) ExportWeightsAndBiasesWithCommandBufferResultStateCanBeTemporary(commandBuffer metal.MTLCommandBuffer, resultStateCanBeTemporary bool) IMPSCNNConvolutionWeightsAndBiasesState {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("exportWeightsAndBiasesWithCommandBuffer:resultStateCanBeTemporary:"), commandBuffer, resultStateCanBeTemporary)
	return MPSCNNConvolutionWeightsAndBiasesStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose/reloadWeightsAndBiases(with:state:)
func (c MPSCNNConvolutionTranspose) ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer metal.MTLCommandBuffer, state IMPSCNNConvolutionWeightsAndBiasesState) {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadWeightsAndBiasesWithCommandBuffer:state:"), commandBuffer, state)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose/reloadWeightsAndBiasesFromDataSource()
func (c MPSCNNConvolutionTranspose) ReloadWeightsAndBiasesFromDataSource() {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadWeightsAndBiasesFromDataSource"))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose/groups
func (c MPSCNNConvolutionTranspose) Groups() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("groups"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose/inputFeatureChannels
func (c MPSCNNConvolutionTranspose) InputFeatureChannels() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("inputFeatureChannels"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose/kernelOffsetX
func (c MPSCNNConvolutionTranspose) KernelOffsetX() int {
	rv := objc.Send[int](c.ID, objc.Sel("kernelOffsetX"))
	return rv
}
func (c MPSCNNConvolutionTranspose) SetKernelOffsetX(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setKernelOffsetX:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose/kernelOffsetY
func (c MPSCNNConvolutionTranspose) KernelOffsetY() int {
	rv := objc.Send[int](c.ID, objc.Sel("kernelOffsetY"))
	return rv
}
func (c MPSCNNConvolutionTranspose) SetKernelOffsetY(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setKernelOffsetY:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose/outputFeatureChannels
func (c MPSCNNConvolutionTranspose) OutputFeatureChannels() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("outputFeatureChannels"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose/accumulatorPrecisionOption
func (c MPSCNNConvolutionTranspose) AccumulatorPrecisionOption() MPSNNConvolutionAccumulatorPrecisionOption {
	rv := objc.Send[MPSNNConvolutionAccumulatorPrecisionOption](c.ID, objc.Sel("accumulatorPrecisionOption"))
	return MPSNNConvolutionAccumulatorPrecisionOption(rv)
}
func (c MPSCNNConvolutionTranspose) SetAccumulatorPrecisionOption(value MPSNNConvolutionAccumulatorPrecisionOption) {
	objc.Send[struct{}](c.ID, objc.Sel("setAccumulatorPrecisionOption:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose/dataSource
func (c MPSCNNConvolutionTranspose) DataSource() MPSCNNConvolutionDataSource {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dataSource"))
	return MPSCNNConvolutionDataSourceObjectFromID(rv)
}

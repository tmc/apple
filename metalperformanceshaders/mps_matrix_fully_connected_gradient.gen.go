// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixFullyConnectedGradient] class.
var (
	_MPSMatrixFullyConnectedGradientClass     MPSMatrixFullyConnectedGradientClass
	_MPSMatrixFullyConnectedGradientClassOnce sync.Once
)

func getMPSMatrixFullyConnectedGradientClass() MPSMatrixFullyConnectedGradientClass {
	_MPSMatrixFullyConnectedGradientClassOnce.Do(func() {
		_MPSMatrixFullyConnectedGradientClass = MPSMatrixFullyConnectedGradientClass{class: objc.GetClass("MPSMatrixFullyConnectedGradient")}
	})
	return _MPSMatrixFullyConnectedGradientClass
}

// GetMPSMatrixFullyConnectedGradientClass returns the class object for MPSMatrixFullyConnectedGradient.
func GetMPSMatrixFullyConnectedGradientClass() MPSMatrixFullyConnectedGradientClass {
	return getMPSMatrixFullyConnectedGradientClass()
}

type MPSMatrixFullyConnectedGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixFullyConnectedGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixFullyConnectedGradientClass) Alloc() MPSMatrixFullyConnectedGradient {
	rv := objc.Send[MPSMatrixFullyConnectedGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel for applying a fully gradient connected neural network layer.
//
// # Instance Properties
//
//   - [MPSMatrixFullyConnectedGradient.Alpha]
//   - [MPSMatrixFullyConnectedGradient.SetAlpha]
//   - [MPSMatrixFullyConnectedGradient.SourceInputFeatureChannels]
//   - [MPSMatrixFullyConnectedGradient.SetSourceInputFeatureChannels]
//   - [MPSMatrixFullyConnectedGradient.SourceNumberOfFeatureVectors]
//   - [MPSMatrixFullyConnectedGradient.SetSourceNumberOfFeatureVectors]
//   - [MPSMatrixFullyConnectedGradient.SourceOutputFeatureChannels]
//   - [MPSMatrixFullyConnectedGradient.SetSourceOutputFeatureChannels]
//
// # Instance Methods
//
//   - [MPSMatrixFullyConnectedGradient.EncodeGradientForDataToCommandBufferGradientMatrixWeightMatrixResultGradientForDataMatrix]
//   - [MPSMatrixFullyConnectedGradient.EncodeGradientForWeightsAndBiasToCommandBufferGradientMatrixInputMatrixResultGradientForWeightMatrixResultGradientForBiasVector]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnectedGradient
type MPSMatrixFullyConnectedGradient struct {
	MPSMatrixBinaryKernel
}

// MPSMatrixFullyConnectedGradientFromID constructs a [MPSMatrixFullyConnectedGradient] from an objc.ID.
//
// A kernel for applying a fully gradient connected neural network layer.
func MPSMatrixFullyConnectedGradientFromID(id objc.ID) MPSMatrixFullyConnectedGradient {
	return MPSMatrixFullyConnectedGradient{MPSMatrixBinaryKernel: MPSMatrixBinaryKernelFromID(id)}
}

// NOTE: MPSMatrixFullyConnectedGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixFullyConnectedGradient] class.
//
// # Instance Properties
//
//   - [IMPSMatrixFullyConnectedGradient.Alpha]
//   - [IMPSMatrixFullyConnectedGradient.SetAlpha]
//   - [IMPSMatrixFullyConnectedGradient.SourceInputFeatureChannels]
//   - [IMPSMatrixFullyConnectedGradient.SetSourceInputFeatureChannels]
//   - [IMPSMatrixFullyConnectedGradient.SourceNumberOfFeatureVectors]
//   - [IMPSMatrixFullyConnectedGradient.SetSourceNumberOfFeatureVectors]
//   - [IMPSMatrixFullyConnectedGradient.SourceOutputFeatureChannels]
//   - [IMPSMatrixFullyConnectedGradient.SetSourceOutputFeatureChannels]
//
// # Instance Methods
//
//   - [IMPSMatrixFullyConnectedGradient.EncodeGradientForDataToCommandBufferGradientMatrixWeightMatrixResultGradientForDataMatrix]
//   - [IMPSMatrixFullyConnectedGradient.EncodeGradientForWeightsAndBiasToCommandBufferGradientMatrixInputMatrixResultGradientForWeightMatrixResultGradientForBiasVector]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnectedGradient
type IMPSMatrixFullyConnectedGradient interface {
	IMPSMatrixBinaryKernel

	// Topic: Instance Properties

	Alpha() float64
	SetAlpha(value float64)
	SourceInputFeatureChannels() uint
	SetSourceInputFeatureChannels(value uint)
	SourceNumberOfFeatureVectors() uint
	SetSourceNumberOfFeatureVectors(value uint)
	SourceOutputFeatureChannels() uint
	SetSourceOutputFeatureChannels(value uint)

	// Topic: Instance Methods

	EncodeGradientForDataToCommandBufferGradientMatrixWeightMatrixResultGradientForDataMatrix(commandBuffer metal.MTLCommandBuffer, gradientMatrix IMPSMatrix, weightMatrix IMPSMatrix, resultGradientForDataMatrix IMPSMatrix)
	EncodeGradientForWeightsAndBiasToCommandBufferGradientMatrixInputMatrixResultGradientForWeightMatrixResultGradientForBiasVector(commandBuffer metal.MTLCommandBuffer, gradientMatrix IMPSMatrix, inputMatrix IMPSMatrix, resultGradientForWeightMatrix IMPSMatrix, resultGradientForBiasVector IMPSVector)
}

// Init initializes the instance.
func (m MPSMatrixFullyConnectedGradient) Init() MPSMatrixFullyConnectedGradient {
	rv := objc.Send[MPSMatrixFullyConnectedGradient](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixFullyConnectedGradient) Autorelease() MPSMatrixFullyConnectedGradient {
	rv := objc.Send[MPSMatrixFullyConnectedGradient](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixFullyConnectedGradient creates a new MPSMatrixFullyConnectedGradient instance.
func NewMPSMatrixFullyConnectedGradient() MPSMatrixFullyConnectedGradient {
	class := getMPSMatrixFullyConnectedGradientClass()
	rv := objc.Send[MPSMatrixFullyConnectedGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixFullyConnectedGradientWithCoder(aDecoder foundation.INSCoder) MPSMatrixFullyConnectedGradient {
	instance := getMPSMatrixFullyConnectedGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixFullyConnectedGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnectedGradient/init(coder:device:)
func NewMatrixFullyConnectedGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixFullyConnectedGradient {
	instance := getMPSMatrixFullyConnectedGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixFullyConnectedGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnectedGradient/init(device:)
func NewMatrixFullyConnectedGradientWithDevice(device metal.MTLDevice) MPSMatrixFullyConnectedGradient {
	instance := getMPSMatrixFullyConnectedGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixFullyConnectedGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnectedGradient/encodeForData(to:gradientMatrix:weightMatrix:resultGradientForDataMatrix:)
func (m MPSMatrixFullyConnectedGradient) EncodeGradientForDataToCommandBufferGradientMatrixWeightMatrixResultGradientForDataMatrix(commandBuffer metal.MTLCommandBuffer, gradientMatrix IMPSMatrix, weightMatrix IMPSMatrix, resultGradientForDataMatrix IMPSMatrix) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeGradientForDataToCommandBuffer:gradientMatrix:weightMatrix:resultGradientForDataMatrix:"), commandBuffer, gradientMatrix, weightMatrix, resultGradientForDataMatrix)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnectedGradient/encodeForWeightsAndBias(to:gradientMatrix:inputMatrix:resultGradientForWeightMatrix:resultGradientForBiasVector:)
func (m MPSMatrixFullyConnectedGradient) EncodeGradientForWeightsAndBiasToCommandBufferGradientMatrixInputMatrixResultGradientForWeightMatrixResultGradientForBiasVector(commandBuffer metal.MTLCommandBuffer, gradientMatrix IMPSMatrix, inputMatrix IMPSMatrix, resultGradientForWeightMatrix IMPSMatrix, resultGradientForBiasVector IMPSVector) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeGradientForWeightsAndBiasToCommandBuffer:gradientMatrix:inputMatrix:resultGradientForWeightMatrix:resultGradientForBiasVector:"), commandBuffer, gradientMatrix, inputMatrix, resultGradientForWeightMatrix, resultGradientForBiasVector)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnectedGradient/alpha
func (m MPSMatrixFullyConnectedGradient) Alpha() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("alpha"))
	return rv
}
func (m MPSMatrixFullyConnectedGradient) SetAlpha(value float64) {
	objc.Send[struct{}](m.ID, objc.Sel("setAlpha:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnectedGradient/sourceInputFeatureChannels
func (m MPSMatrixFullyConnectedGradient) SourceInputFeatureChannels() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceInputFeatureChannels"))
	return rv
}
func (m MPSMatrixFullyConnectedGradient) SetSourceInputFeatureChannels(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceInputFeatureChannels:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnectedGradient/sourceNumberOfFeatureVectors
func (m MPSMatrixFullyConnectedGradient) SourceNumberOfFeatureVectors() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}
func (m MPSMatrixFullyConnectedGradient) SetSourceNumberOfFeatureVectors(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixFullyConnectedGradient/sourceOutputFeatureChannels
func (m MPSMatrixFullyConnectedGradient) SourceOutputFeatureChannels() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("sourceOutputFeatureChannels"))
	return rv
}
func (m MPSMatrixFullyConnectedGradient) SetSourceOutputFeatureChannels(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setSourceOutputFeatureChannels:"), value)
}

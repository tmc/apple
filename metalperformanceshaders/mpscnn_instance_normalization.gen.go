// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNInstanceNormalization] class.
var (
	_MPSCNNInstanceNormalizationClass     MPSCNNInstanceNormalizationClass
	_MPSCNNInstanceNormalizationClassOnce sync.Once
)

func getMPSCNNInstanceNormalizationClass() MPSCNNInstanceNormalizationClass {
	_MPSCNNInstanceNormalizationClassOnce.Do(func() {
		_MPSCNNInstanceNormalizationClass = MPSCNNInstanceNormalizationClass{class: objc.GetClass("MPSCNNInstanceNormalization")}
	})
	return _MPSCNNInstanceNormalizationClass
}

// GetMPSCNNInstanceNormalizationClass returns the class object for MPSCNNInstanceNormalization.
func GetMPSCNNInstanceNormalizationClass() MPSCNNInstanceNormalizationClass {
	return getMPSCNNInstanceNormalizationClass()
}

type MPSCNNInstanceNormalizationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNInstanceNormalizationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNInstanceNormalizationClass) Alloc() MPSCNNInstanceNormalization {
	rv := objc.Send[MPSCNNInstanceNormalization](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An instance normalization kernel.
//
// # Initializers
//
//   - [MPSCNNInstanceNormalization.InitWithDeviceDataSource]
//
// # Instance Properties
//
//   - [MPSCNNInstanceNormalization.DataSource]
//   - [MPSCNNInstanceNormalization.Epsilon]
//   - [MPSCNNInstanceNormalization.SetEpsilon]
//
// # Instance Methods
//
//   - [MPSCNNInstanceNormalization.ReloadGammaAndBetaWithCommandBufferGammaAndBetaState]
//   - [MPSCNNInstanceNormalization.ReloadGammaAndBetaFromDataSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalization
type MPSCNNInstanceNormalization struct {
	MPSCNNKernel
}

// MPSCNNInstanceNormalizationFromID constructs a [MPSCNNInstanceNormalization] from an objc.ID.
//
// An instance normalization kernel.
func MPSCNNInstanceNormalizationFromID(id objc.ID) MPSCNNInstanceNormalization {
	return MPSCNNInstanceNormalization{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSCNNInstanceNormalization adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNInstanceNormalization] class.
//
// # Initializers
//
//   - [IMPSCNNInstanceNormalization.InitWithDeviceDataSource]
//
// # Instance Properties
//
//   - [IMPSCNNInstanceNormalization.DataSource]
//   - [IMPSCNNInstanceNormalization.Epsilon]
//   - [IMPSCNNInstanceNormalization.SetEpsilon]
//
// # Instance Methods
//
//   - [IMPSCNNInstanceNormalization.ReloadGammaAndBetaWithCommandBufferGammaAndBetaState]
//   - [IMPSCNNInstanceNormalization.ReloadGammaAndBetaFromDataSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalization
type IMPSCNNInstanceNormalization interface {
	IMPSCNNKernel

	// Topic: Initializers

	InitWithDeviceDataSource(device metal.MTLDevice, dataSource MPSCNNInstanceNormalizationDataSource) MPSCNNInstanceNormalization

	// Topic: Instance Properties

	DataSource() MPSCNNInstanceNormalizationDataSource
	Epsilon() float32
	SetEpsilon(value float32)

	// Topic: Instance Methods

	ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer metal.MTLCommandBuffer, gammaAndBetaState IMPSCNNNormalizationGammaAndBetaState)
	ReloadGammaAndBetaFromDataSource()
}

// Init initializes the instance.
func (c MPSCNNInstanceNormalization) Init() MPSCNNInstanceNormalization {
	rv := objc.Send[MPSCNNInstanceNormalization](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNInstanceNormalization) Autorelease() MPSCNNInstanceNormalization {
	rv := objc.Send[MPSCNNInstanceNormalization](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNInstanceNormalization creates a new MPSCNNInstanceNormalization instance.
func NewMPSCNNInstanceNormalization() MPSCNNInstanceNormalization {
	class := getMPSCNNInstanceNormalizationClass()
	rv := objc.Send[MPSCNNInstanceNormalization](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNInstanceNormalizationWithCoder(aDecoder foundation.INSCoder) MPSCNNInstanceNormalization {
	instance := getMPSCNNInstanceNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNInstanceNormalizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalization/init(coder:device:)
func NewCNNInstanceNormalizationWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNInstanceNormalization {
	instance := getMPSCNNInstanceNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNInstanceNormalizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNInstanceNormalizationWithDevice(device metal.MTLDevice) MPSCNNInstanceNormalization {
	instance := getMPSCNNInstanceNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNInstanceNormalizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalization/init(device:dataSource:)
func NewCNNInstanceNormalizationWithDeviceDataSource(device metal.MTLDevice, dataSource MPSCNNInstanceNormalizationDataSource) MPSCNNInstanceNormalization {
	instance := getMPSCNNInstanceNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:dataSource:"), device, dataSource)
	return MPSCNNInstanceNormalizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalization/init(device:dataSource:)
func (c MPSCNNInstanceNormalization) InitWithDeviceDataSource(device metal.MTLDevice, dataSource MPSCNNInstanceNormalizationDataSource) MPSCNNInstanceNormalization {
	rv := objc.Send[MPSCNNInstanceNormalization](c.ID, objc.Sel("initWithDevice:dataSource:"), device, dataSource)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalization/reloadGammaAndBeta(with:gammaAndBetaState:)
func (c MPSCNNInstanceNormalization) ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer metal.MTLCommandBuffer, gammaAndBetaState IMPSCNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadGammaAndBetaWithCommandBuffer:gammaAndBetaState:"), commandBuffer, gammaAndBetaState)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalization/reloadGammaAndBetaFromDataSource()
func (c MPSCNNInstanceNormalization) ReloadGammaAndBetaFromDataSource() {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadGammaAndBetaFromDataSource"))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalization/dataSource
func (c MPSCNNInstanceNormalization) DataSource() MPSCNNInstanceNormalizationDataSource {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dataSource"))
	return MPSCNNInstanceNormalizationDataSourceObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalization/epsilon
func (c MPSCNNInstanceNormalization) Epsilon() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("epsilon"))
	return rv
}
func (c MPSCNNInstanceNormalization) SetEpsilon(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setEpsilon:"), value)
}

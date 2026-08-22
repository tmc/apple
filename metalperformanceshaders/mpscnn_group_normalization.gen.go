// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNGroupNormalization] class.
var (
	_MPSCNNGroupNormalizationClass     MPSCNNGroupNormalizationClass
	_MPSCNNGroupNormalizationClassOnce sync.Once
)

func getMPSCNNGroupNormalizationClass() MPSCNNGroupNormalizationClass {
	_MPSCNNGroupNormalizationClassOnce.Do(func() {
		_MPSCNNGroupNormalizationClass = MPSCNNGroupNormalizationClass{class: objc.GetClass("MPSCNNGroupNormalization")}
	})
	return _MPSCNNGroupNormalizationClass
}

// GetMPSCNNGroupNormalizationClass returns the class object for MPSCNNGroupNormalization.
func GetMPSCNNGroupNormalizationClass() MPSCNNGroupNormalizationClass {
	return getMPSCNNGroupNormalizationClass()
}

type MPSCNNGroupNormalizationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNGroupNormalizationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNGroupNormalizationClass) Alloc() MPSCNNGroupNormalization {
	rv := objc.Send[MPSCNNGroupNormalization](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSCNNGroupNormalization.InitWithDeviceDataSource]
//
// # Instance Properties
//
//   - [MPSCNNGroupNormalization.DataSource]
//   - [MPSCNNGroupNormalization.Epsilon]
//   - [MPSCNNGroupNormalization.SetEpsilon]
//
// # Instance Methods
//
//   - [MPSCNNGroupNormalization.ReloadGammaAndBetaWithCommandBufferGammaAndBetaState]
//   - [MPSCNNGroupNormalization.ReloadGammaAndBetaFromDataSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalization
type MPSCNNGroupNormalization struct {
	MPSCNNKernel
}

// MPSCNNGroupNormalizationFromID constructs a [MPSCNNGroupNormalization] from an objc.ID.
func MPSCNNGroupNormalizationFromID(id objc.ID) MPSCNNGroupNormalization {
	return MPSCNNGroupNormalization{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSCNNGroupNormalization adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNGroupNormalization] class.
//
// # Initializers
//
//   - [IMPSCNNGroupNormalization.InitWithDeviceDataSource]
//
// # Instance Properties
//
//   - [IMPSCNNGroupNormalization.DataSource]
//   - [IMPSCNNGroupNormalization.Epsilon]
//   - [IMPSCNNGroupNormalization.SetEpsilon]
//
// # Instance Methods
//
//   - [IMPSCNNGroupNormalization.ReloadGammaAndBetaWithCommandBufferGammaAndBetaState]
//   - [IMPSCNNGroupNormalization.ReloadGammaAndBetaFromDataSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalization
type IMPSCNNGroupNormalization interface {
	IMPSCNNKernel

	// Topic: Initializers

	InitWithDeviceDataSource(device metal.MTLDevice, dataSource MPSCNNGroupNormalizationDataSource) MPSCNNGroupNormalization

	// Topic: Instance Properties

	DataSource() MPSCNNGroupNormalizationDataSource
	Epsilon() float32
	SetEpsilon(value float32)

	// Topic: Instance Methods

	ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer metal.MTLCommandBuffer, gammaAndBetaState IMPSCNNNormalizationGammaAndBetaState)
	ReloadGammaAndBetaFromDataSource()
}

// Init initializes the instance.
func (c MPSCNNGroupNormalization) Init() MPSCNNGroupNormalization {
	rv := objc.Send[MPSCNNGroupNormalization](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNGroupNormalization) Autorelease() MPSCNNGroupNormalization {
	rv := objc.Send[MPSCNNGroupNormalization](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNGroupNormalization creates a new MPSCNNGroupNormalization instance.
func NewMPSCNNGroupNormalization() MPSCNNGroupNormalization {
	class := getMPSCNNGroupNormalizationClass()
	rv := objc.Send[MPSCNNGroupNormalization](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNGroupNormalizationWithCoder(aDecoder foundation.INSCoder) MPSCNNGroupNormalization {
	instance := getMPSCNNGroupNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNGroupNormalizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalization/init(coder:device:)
func NewCNNGroupNormalizationWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNGroupNormalization {
	instance := getMPSCNNGroupNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNGroupNormalizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNGroupNormalizationWithDevice(device metal.MTLDevice) MPSCNNGroupNormalization {
	instance := getMPSCNNGroupNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNGroupNormalizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalization/init(device:dataSource:)
func NewCNNGroupNormalizationWithDeviceDataSource(device metal.MTLDevice, dataSource MPSCNNGroupNormalizationDataSource) MPSCNNGroupNormalization {
	instance := getMPSCNNGroupNormalizationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:dataSource:"), device, dataSource)
	return MPSCNNGroupNormalizationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalization/init(device:dataSource:)
func (c MPSCNNGroupNormalization) InitWithDeviceDataSource(device metal.MTLDevice, dataSource MPSCNNGroupNormalizationDataSource) MPSCNNGroupNormalization {
	rv := objc.Send[MPSCNNGroupNormalization](c.ID, objc.Sel("initWithDevice:dataSource:"), device, dataSource)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalization/reloadGammaAndBeta(with:gammaAndBetaState:)
func (c MPSCNNGroupNormalization) ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer metal.MTLCommandBuffer, gammaAndBetaState IMPSCNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadGammaAndBetaWithCommandBuffer:gammaAndBetaState:"), commandBuffer, gammaAndBetaState)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalization/reloadGammaAndBetaFromDataSource()
func (c MPSCNNGroupNormalization) ReloadGammaAndBetaFromDataSource() {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadGammaAndBetaFromDataSource"))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalization/dataSource
func (c MPSCNNGroupNormalization) DataSource() MPSCNNGroupNormalizationDataSource {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dataSource"))
	return MPSCNNGroupNormalizationDataSourceObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalization/epsilon
func (c MPSCNNGroupNormalization) Epsilon() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("epsilon"))
	return rv
}
func (c MPSCNNGroupNormalization) SetEpsilon(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setEpsilon:"), value)
}

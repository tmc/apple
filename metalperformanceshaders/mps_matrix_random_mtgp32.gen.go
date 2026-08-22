// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixRandomMTGP32] class.
var (
	_MPSMatrixRandomMTGP32Class     MPSMatrixRandomMTGP32Class
	_MPSMatrixRandomMTGP32ClassOnce sync.Once
)

func getMPSMatrixRandomMTGP32Class() MPSMatrixRandomMTGP32Class {
	_MPSMatrixRandomMTGP32ClassOnce.Do(func() {
		_MPSMatrixRandomMTGP32Class = MPSMatrixRandomMTGP32Class{class: objc.GetClass("MPSMatrixRandomMTGP32")}
	})
	return _MPSMatrixRandomMTGP32Class
}

// GetMPSMatrixRandomMTGP32Class returns the class object for MPSMatrixRandomMTGP32.
func GetMPSMatrixRandomMTGP32Class() MPSMatrixRandomMTGP32Class {
	return getMPSMatrixRandomMTGP32Class()
}

type MPSMatrixRandomMTGP32Class struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixRandomMTGP32Class) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixRandomMTGP32Class) Alloc() MPSMatrixRandomMTGP32 {
	rv := objc.Send[MPSMatrixRandomMTGP32](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSMatrixRandomMTGP32.InitWithDeviceDestinationDataTypeSeed]
//   - [MPSMatrixRandomMTGP32.InitWithDeviceDestinationDataTypeSeedDistributionDescriptor]
//
// # Instance Methods
//
//   - [MPSMatrixRandomMTGP32.SynchronizeStateOnCommandBuffer]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomMTGP32
type MPSMatrixRandomMTGP32 struct {
	MPSMatrixRandom
}

// MPSMatrixRandomMTGP32FromID constructs a [MPSMatrixRandomMTGP32] from an objc.ID.
func MPSMatrixRandomMTGP32FromID(id objc.ID) MPSMatrixRandomMTGP32 {
	return MPSMatrixRandomMTGP32{MPSMatrixRandom: MPSMatrixRandomFromID(id)}
}

// NOTE: MPSMatrixRandomMTGP32 adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixRandomMTGP32] class.
//
// # Initializers
//
//   - [IMPSMatrixRandomMTGP32.InitWithDeviceDestinationDataTypeSeed]
//   - [IMPSMatrixRandomMTGP32.InitWithDeviceDestinationDataTypeSeedDistributionDescriptor]
//
// # Instance Methods
//
//   - [IMPSMatrixRandomMTGP32.SynchronizeStateOnCommandBuffer]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomMTGP32
type IMPSMatrixRandomMTGP32 interface {
	IMPSMatrixRandom

	// Topic: Initializers

	InitWithDeviceDestinationDataTypeSeed(device metal.MTLDevice, destinationDataType MPSDataType, seed uint) MPSMatrixRandomMTGP32
	InitWithDeviceDestinationDataTypeSeedDistributionDescriptor(device metal.MTLDevice, destinationDataType MPSDataType, seed uint, distributionDescriptor IMPSMatrixRandomDistributionDescriptor) MPSMatrixRandomMTGP32

	// Topic: Instance Methods

	SynchronizeStateOnCommandBuffer(commandBuffer metal.MTLCommandBuffer)
}

// Init initializes the instance.
func (m MPSMatrixRandomMTGP32) Init() MPSMatrixRandomMTGP32 {
	rv := objc.Send[MPSMatrixRandomMTGP32](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixRandomMTGP32) Autorelease() MPSMatrixRandomMTGP32 {
	rv := objc.Send[MPSMatrixRandomMTGP32](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixRandomMTGP32 creates a new MPSMatrixRandomMTGP32 instance.
func NewMPSMatrixRandomMTGP32() MPSMatrixRandomMTGP32 {
	class := getMPSMatrixRandomMTGP32Class()
	rv := objc.Send[MPSMatrixRandomMTGP32](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixRandomMTGP32WithCoder(aDecoder foundation.INSCoder) MPSMatrixRandomMTGP32 {
	instance := getMPSMatrixRandomMTGP32Class().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixRandomMTGP32FromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomMTGP32/init(coder:device:)
func NewMatrixRandomMTGP32WithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixRandomMTGP32 {
	instance := getMPSMatrixRandomMTGP32Class().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixRandomMTGP32FromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomMTGP32/init(device:)
func NewMatrixRandomMTGP32WithDevice(device metal.MTLDevice) MPSMatrixRandomMTGP32 {
	instance := getMPSMatrixRandomMTGP32Class().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixRandomMTGP32FromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomMTGP32/init(device:destinationDataType:seed:)
func NewMatrixRandomMTGP32WithDeviceDestinationDataTypeSeed(device metal.MTLDevice, destinationDataType MPSDataType, seed uint) MPSMatrixRandomMTGP32 {
	instance := getMPSMatrixRandomMTGP32Class().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:destinationDataType:seed:"), device, destinationDataType, seed)
	return MPSMatrixRandomMTGP32FromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomMTGP32/init(device:destinationDataType:seed:distributionDescriptor:)
func NewMatrixRandomMTGP32WithDeviceDestinationDataTypeSeedDistributionDescriptor(device metal.MTLDevice, destinationDataType MPSDataType, seed uint, distributionDescriptor IMPSMatrixRandomDistributionDescriptor) MPSMatrixRandomMTGP32 {
	instance := getMPSMatrixRandomMTGP32Class().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:destinationDataType:seed:distributionDescriptor:"), device, destinationDataType, seed, distributionDescriptor)
	return MPSMatrixRandomMTGP32FromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomMTGP32/init(device:destinationDataType:seed:)
func (m MPSMatrixRandomMTGP32) InitWithDeviceDestinationDataTypeSeed(device metal.MTLDevice, destinationDataType MPSDataType, seed uint) MPSMatrixRandomMTGP32 {
	rv := objc.Send[MPSMatrixRandomMTGP32](m.ID, objc.Sel("initWithDevice:destinationDataType:seed:"), device, destinationDataType, seed)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomMTGP32/init(device:destinationDataType:seed:distributionDescriptor:)
func (m MPSMatrixRandomMTGP32) InitWithDeviceDestinationDataTypeSeedDistributionDescriptor(device metal.MTLDevice, destinationDataType MPSDataType, seed uint, distributionDescriptor IMPSMatrixRandomDistributionDescriptor) MPSMatrixRandomMTGP32 {
	rv := objc.Send[MPSMatrixRandomMTGP32](m.ID, objc.Sel("initWithDevice:destinationDataType:seed:distributionDescriptor:"), device, destinationDataType, seed, distributionDescriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomMTGP32/synchronizeState(on:)
func (m MPSMatrixRandomMTGP32) SynchronizeStateOnCommandBuffer(commandBuffer metal.MTLCommandBuffer) {
	objc.Send[objc.ID](m.ID, objc.Sel("synchronizeStateOnCommandBuffer:"), commandBuffer)
}

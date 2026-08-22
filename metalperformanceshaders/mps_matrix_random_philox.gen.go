// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSMatrixRandomPhilox] class.
var (
	_MPSMatrixRandomPhiloxClass     MPSMatrixRandomPhiloxClass
	_MPSMatrixRandomPhiloxClassOnce sync.Once
)

func getMPSMatrixRandomPhiloxClass() MPSMatrixRandomPhiloxClass {
	_MPSMatrixRandomPhiloxClassOnce.Do(func() {
		_MPSMatrixRandomPhiloxClass = MPSMatrixRandomPhiloxClass{class: objc.GetClass("MPSMatrixRandomPhilox")}
	})
	return _MPSMatrixRandomPhiloxClass
}

// GetMPSMatrixRandomPhiloxClass returns the class object for MPSMatrixRandomPhilox.
func GetMPSMatrixRandomPhiloxClass() MPSMatrixRandomPhiloxClass {
	return getMPSMatrixRandomPhiloxClass()
}

type MPSMatrixRandomPhiloxClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixRandomPhiloxClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixRandomPhiloxClass) Alloc() MPSMatrixRandomPhilox {
	rv := objc.Send[MPSMatrixRandomPhilox](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSMatrixRandomPhilox.InitWithDeviceDestinationDataTypeSeed]
//   - [MPSMatrixRandomPhilox.InitWithDeviceDestinationDataTypeSeedDistributionDescriptor]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomPhilox
type MPSMatrixRandomPhilox struct {
	MPSMatrixRandom
}

// MPSMatrixRandomPhiloxFromID constructs a [MPSMatrixRandomPhilox] from an objc.ID.
func MPSMatrixRandomPhiloxFromID(id objc.ID) MPSMatrixRandomPhilox {
	return MPSMatrixRandomPhilox{MPSMatrixRandom: MPSMatrixRandomFromID(id)}
}

// NOTE: MPSMatrixRandomPhilox adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixRandomPhilox] class.
//
// # Initializers
//
//   - [IMPSMatrixRandomPhilox.InitWithDeviceDestinationDataTypeSeed]
//   - [IMPSMatrixRandomPhilox.InitWithDeviceDestinationDataTypeSeedDistributionDescriptor]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomPhilox
type IMPSMatrixRandomPhilox interface {
	IMPSMatrixRandom

	// Topic: Initializers

	InitWithDeviceDestinationDataTypeSeed(device metal.MTLDevice, destinationDataType MPSDataType, seed uint) MPSMatrixRandomPhilox
	InitWithDeviceDestinationDataTypeSeedDistributionDescriptor(device metal.MTLDevice, destinationDataType MPSDataType, seed uint, distributionDescriptor IMPSMatrixRandomDistributionDescriptor) MPSMatrixRandomPhilox
}

// Init initializes the instance.
func (m MPSMatrixRandomPhilox) Init() MPSMatrixRandomPhilox {
	rv := objc.Send[MPSMatrixRandomPhilox](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixRandomPhilox) Autorelease() MPSMatrixRandomPhilox {
	rv := objc.Send[MPSMatrixRandomPhilox](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixRandomPhilox creates a new MPSMatrixRandomPhilox instance.
func NewMPSMatrixRandomPhilox() MPSMatrixRandomPhilox {
	class := getMPSMatrixRandomPhiloxClass()
	rv := objc.Send[MPSMatrixRandomPhilox](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewMatrixRandomPhiloxWithCoder(aDecoder foundation.INSCoder) MPSMatrixRandomPhilox {
	instance := getMPSMatrixRandomPhiloxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSMatrixRandomPhiloxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomPhilox/init(coder:device:)
func NewMatrixRandomPhiloxWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSMatrixRandomPhilox {
	instance := getMPSMatrixRandomPhiloxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSMatrixRandomPhiloxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomPhilox/init(device:)
func NewMatrixRandomPhiloxWithDevice(device metal.MTLDevice) MPSMatrixRandomPhilox {
	instance := getMPSMatrixRandomPhiloxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSMatrixRandomPhiloxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomPhilox/init(device:destinationDataType:seed:)
func NewMatrixRandomPhiloxWithDeviceDestinationDataTypeSeed(device metal.MTLDevice, destinationDataType MPSDataType, seed uint) MPSMatrixRandomPhilox {
	instance := getMPSMatrixRandomPhiloxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:destinationDataType:seed:"), device, destinationDataType, seed)
	return MPSMatrixRandomPhiloxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomPhilox/init(device:destinationDataType:seed:distributionDescriptor:)
func NewMatrixRandomPhiloxWithDeviceDestinationDataTypeSeedDistributionDescriptor(device metal.MTLDevice, destinationDataType MPSDataType, seed uint, distributionDescriptor IMPSMatrixRandomDistributionDescriptor) MPSMatrixRandomPhilox {
	instance := getMPSMatrixRandomPhiloxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:destinationDataType:seed:distributionDescriptor:"), device, destinationDataType, seed, distributionDescriptor)
	return MPSMatrixRandomPhiloxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomPhilox/init(device:destinationDataType:seed:)
func (m MPSMatrixRandomPhilox) InitWithDeviceDestinationDataTypeSeed(device metal.MTLDevice, destinationDataType MPSDataType, seed uint) MPSMatrixRandomPhilox {
	rv := objc.Send[MPSMatrixRandomPhilox](m.ID, objc.Sel("initWithDevice:destinationDataType:seed:"), device, destinationDataType, seed)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomPhilox/init(device:destinationDataType:seed:distributionDescriptor:)
func (m MPSMatrixRandomPhilox) InitWithDeviceDestinationDataTypeSeedDistributionDescriptor(device metal.MTLDevice, destinationDataType MPSDataType, seed uint, distributionDescriptor IMPSMatrixRandomDistributionDescriptor) MPSMatrixRandomPhilox {
	rv := objc.Send[MPSMatrixRandomPhilox](m.ID, objc.Sel("initWithDevice:destinationDataType:seed:distributionDescriptor:"), device, destinationDataType, seed, distributionDescriptor)
	return rv
}

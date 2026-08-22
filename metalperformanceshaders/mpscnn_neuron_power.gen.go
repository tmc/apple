// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronPower] class.
var (
	_MPSCNNNeuronPowerClass     MPSCNNNeuronPowerClass
	_MPSCNNNeuronPowerClassOnce sync.Once
)

func getMPSCNNNeuronPowerClass() MPSCNNNeuronPowerClass {
	_MPSCNNNeuronPowerClassOnce.Do(func() {
		_MPSCNNNeuronPowerClass = MPSCNNNeuronPowerClass{class: objc.GetClass("MPSCNNNeuronPower")}
	})
	return _MPSCNNNeuronPowerClass
}

// GetMPSCNNNeuronPowerClass returns the class object for MPSCNNNeuronPower.
func GetMPSCNNNeuronPowerClass() MPSCNNNeuronPowerClass {
	return getMPSCNNNeuronPowerClass()
}

type MPSCNNNeuronPowerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronPowerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronPowerClass) Alloc() MPSCNNNeuronPower {
	rv := objc.Send[MPSCNNNeuronPower](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A power neuron filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPower
type MPSCNNNeuronPower struct {
	MPSCNNNeuron
}

// MPSCNNNeuronPowerFromID constructs a [MPSCNNNeuronPower] from an objc.ID.
//
// A power neuron filter.
func MPSCNNNeuronPowerFromID(id objc.ID) MPSCNNNeuronPower {
	return MPSCNNNeuronPower{MPSCNNNeuron: MPSCNNNeuronFromID(id)}
}

// NOTE: MPSCNNNeuronPower adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronPower] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPower
type IMPSCNNNeuronPower interface {
	IMPSCNNNeuron
}

// Init initializes the instance.
func (c MPSCNNNeuronPower) Init() MPSCNNNeuronPower {
	rv := objc.Send[MPSCNNNeuronPower](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronPower) Autorelease() MPSCNNNeuronPower {
	rv := objc.Send[MPSCNNNeuronPower](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronPower creates a new MPSCNNNeuronPower instance.
func NewMPSCNNNeuronPower() MPSCNNNeuronPower {
	class := getMPSCNNNeuronPowerClass()
	rv := objc.Send[MPSCNNNeuronPower](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNNeuronPowerWithCoder(aDecoder foundation.INSCoder) MPSCNNNeuronPower {
	instance := getMPSCNNNeuronPowerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNNeuronPowerFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(coder:device:)
func NewCNNNeuronPowerWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNNeuronPower {
	instance := getMPSCNNNeuronPowerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNNeuronPowerFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNNeuronPowerWithDevice(device metal.MTLDevice) MPSCNNNeuronPower {
	instance := getMPSCNNNeuronPowerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNNeuronPowerFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(device:neuronDescriptor:)
func NewCNNNeuronPowerWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuronPower {
	instance := getMPSCNNNeuronPowerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	return MPSCNNNeuronPowerFromID(rv)
}

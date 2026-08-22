// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronLinear] class.
var (
	_MPSCNNNeuronLinearClass     MPSCNNNeuronLinearClass
	_MPSCNNNeuronLinearClassOnce sync.Once
)

func getMPSCNNNeuronLinearClass() MPSCNNNeuronLinearClass {
	_MPSCNNNeuronLinearClassOnce.Do(func() {
		_MPSCNNNeuronLinearClass = MPSCNNNeuronLinearClass{class: objc.GetClass("MPSCNNNeuronLinear")}
	})
	return _MPSCNNNeuronLinearClass
}

// GetMPSCNNNeuronLinearClass returns the class object for MPSCNNNeuronLinear.
func GetMPSCNNNeuronLinearClass() MPSCNNNeuronLinearClass {
	return getMPSCNNNeuronLinearClass()
}

type MPSCNNNeuronLinearClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronLinearClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronLinearClass) Alloc() MPSCNNNeuronLinear {
	rv := objc.Send[MPSCNNNeuronLinear](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A linear neuron filter.
//
// # Overview
//
// For each pixel in an image, the filter applies the following function:
//
// [media-2903542]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLinear
type MPSCNNNeuronLinear struct {
	MPSCNNNeuron
}

// MPSCNNNeuronLinearFromID constructs a [MPSCNNNeuronLinear] from an objc.ID.
//
// A linear neuron filter.
func MPSCNNNeuronLinearFromID(id objc.ID) MPSCNNNeuronLinear {
	return MPSCNNNeuronLinear{MPSCNNNeuron: MPSCNNNeuronFromID(id)}
}

// NOTE: MPSCNNNeuronLinear adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronLinear] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLinear
type IMPSCNNNeuronLinear interface {
	IMPSCNNNeuron
}

// Init initializes the instance.
func (c MPSCNNNeuronLinear) Init() MPSCNNNeuronLinear {
	rv := objc.Send[MPSCNNNeuronLinear](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronLinear) Autorelease() MPSCNNNeuronLinear {
	rv := objc.Send[MPSCNNNeuronLinear](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronLinear creates a new MPSCNNNeuronLinear instance.
func NewMPSCNNNeuronLinear() MPSCNNNeuronLinear {
	class := getMPSCNNNeuronLinearClass()
	rv := objc.Send[MPSCNNNeuronLinear](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNNeuronLinearWithCoder(aDecoder foundation.INSCoder) MPSCNNNeuronLinear {
	instance := getMPSCNNNeuronLinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNNeuronLinearFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(coder:device:)
func NewCNNNeuronLinearWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNNeuronLinear {
	instance := getMPSCNNNeuronLinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNNeuronLinearFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNNeuronLinearWithDevice(device metal.MTLDevice) MPSCNNNeuronLinear {
	instance := getMPSCNNNeuronLinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNNeuronLinearFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(device:neuronDescriptor:)
func NewCNNNeuronLinearWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuronLinear {
	instance := getMPSCNNNeuronLinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	return MPSCNNNeuronLinearFromID(rv)
}

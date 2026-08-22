// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronELU] class.
var (
	_MPSCNNNeuronELUClass     MPSCNNNeuronELUClass
	_MPSCNNNeuronELUClassOnce sync.Once
)

func getMPSCNNNeuronELUClass() MPSCNNNeuronELUClass {
	_MPSCNNNeuronELUClassOnce.Do(func() {
		_MPSCNNNeuronELUClass = MPSCNNNeuronELUClass{class: objc.GetClass("MPSCNNNeuronELU")}
	})
	return _MPSCNNNeuronELUClass
}

// GetMPSCNNNeuronELUClass returns the class object for MPSCNNNeuronELU.
func GetMPSCNNNeuronELUClass() MPSCNNNeuronELUClass {
	return getMPSCNNNeuronELUClass()
}

type MPSCNNNeuronELUClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronELUClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronELUClass) Alloc() MPSCNNNeuronELU {
	rv := objc.Send[MPSCNNNeuronELU](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A parametric ELU neuron filter.
//
// # Overview
//
// For each pixel in an image, the filter applies the following function:
//
// ![f(x) = a * (exp(x) - 1) if x < 0 | f(x) = [ a * (exp(x) - 1) if x < 0]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronELU
type MPSCNNNeuronELU struct {
	MPSCNNNeuron
}

// MPSCNNNeuronELUFromID constructs a [MPSCNNNeuronELU] from an objc.ID.
//
// A parametric ELU neuron filter.
func MPSCNNNeuronELUFromID(id objc.ID) MPSCNNNeuronELU {
	return MPSCNNNeuronELU{MPSCNNNeuron: MPSCNNNeuronFromID(id)}
}

// NOTE: MPSCNNNeuronELU adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronELU] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronELU
type IMPSCNNNeuronELU interface {
	IMPSCNNNeuron
}

// Init initializes the instance.
func (c MPSCNNNeuronELU) Init() MPSCNNNeuronELU {
	rv := objc.Send[MPSCNNNeuronELU](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronELU) Autorelease() MPSCNNNeuronELU {
	rv := objc.Send[MPSCNNNeuronELU](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronELU creates a new MPSCNNNeuronELU instance.
func NewMPSCNNNeuronELU() MPSCNNNeuronELU {
	class := getMPSCNNNeuronELUClass()
	rv := objc.Send[MPSCNNNeuronELU](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNNeuronELUWithCoder(aDecoder foundation.INSCoder) MPSCNNNeuronELU {
	instance := getMPSCNNNeuronELUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNNeuronELUFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(coder:device:)
func NewCNNNeuronELUWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNNeuronELU {
	instance := getMPSCNNNeuronELUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNNeuronELUFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNNeuronELUWithDevice(device metal.MTLDevice) MPSCNNNeuronELU {
	instance := getMPSCNNNeuronELUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNNeuronELUFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(device:neuronDescriptor:)
func NewCNNNeuronELUWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuronELU {
	instance := getMPSCNNNeuronELUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	return MPSCNNNeuronELUFromID(rv)
}

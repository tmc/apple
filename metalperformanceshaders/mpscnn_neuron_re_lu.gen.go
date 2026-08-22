// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronReLU] class.
var (
	_MPSCNNNeuronReLUClass     MPSCNNNeuronReLUClass
	_MPSCNNNeuronReLUClassOnce sync.Once
)

func getMPSCNNNeuronReLUClass() MPSCNNNeuronReLUClass {
	_MPSCNNNeuronReLUClassOnce.Do(func() {
		_MPSCNNNeuronReLUClass = MPSCNNNeuronReLUClass{class: objc.GetClass("MPSCNNNeuronReLU")}
	})
	return _MPSCNNNeuronReLUClass
}

// GetMPSCNNNeuronReLUClass returns the class object for MPSCNNNeuronReLU.
func GetMPSCNNNeuronReLUClass() MPSCNNNeuronReLUClass {
	return getMPSCNNNeuronReLUClass()
}

type MPSCNNNeuronReLUClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronReLUClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronReLUClass) Alloc() MPSCNNNeuronReLU {
	rv := objc.Send[MPSCNNNeuronReLU](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A ReLU (Rectified Linear Unit) neuron filter.
//
// # Overview
//
// For each pixel in an image, the filter applies the following function:
//
// [media-2903544]
//
// This filter is called l__eaky ReLU in CNN literature. Some CNN literature
// defines classical ReLU as `max(0, x)`. If you want this behavior, simply
// set the `a` property to `0`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLU
type MPSCNNNeuronReLU struct {
	MPSCNNNeuron
}

// MPSCNNNeuronReLUFromID constructs a [MPSCNNNeuronReLU] from an objc.ID.
//
// A ReLU (Rectified Linear Unit) neuron filter.
func MPSCNNNeuronReLUFromID(id objc.ID) MPSCNNNeuronReLU {
	return MPSCNNNeuronReLU{MPSCNNNeuron: MPSCNNNeuronFromID(id)}
}

// NOTE: MPSCNNNeuronReLU adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronReLU] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLU
type IMPSCNNNeuronReLU interface {
	IMPSCNNNeuron
}

// Init initializes the instance.
func (c MPSCNNNeuronReLU) Init() MPSCNNNeuronReLU {
	rv := objc.Send[MPSCNNNeuronReLU](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronReLU) Autorelease() MPSCNNNeuronReLU {
	rv := objc.Send[MPSCNNNeuronReLU](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronReLU creates a new MPSCNNNeuronReLU instance.
func NewMPSCNNNeuronReLU() MPSCNNNeuronReLU {
	class := getMPSCNNNeuronReLUClass()
	rv := objc.Send[MPSCNNNeuronReLU](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNNeuronReLUWithCoder(aDecoder foundation.INSCoder) MPSCNNNeuronReLU {
	instance := getMPSCNNNeuronReLUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNNeuronReLUFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(coder:device:)
func NewCNNNeuronReLUWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNNeuronReLU {
	instance := getMPSCNNNeuronReLUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNNeuronReLUFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNNeuronReLUWithDevice(device metal.MTLDevice) MPSCNNNeuronReLU {
	instance := getMPSCNNNeuronReLUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNNeuronReLUFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(device:neuronDescriptor:)
func NewCNNNeuronReLUWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuronReLU {
	instance := getMPSCNNNeuronReLUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	return MPSCNNNeuronReLUFromID(rv)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronPReLU] class.
var (
	_MPSCNNNeuronPReLUClass     MPSCNNNeuronPReLUClass
	_MPSCNNNeuronPReLUClassOnce sync.Once
)

func getMPSCNNNeuronPReLUClass() MPSCNNNeuronPReLUClass {
	_MPSCNNNeuronPReLUClassOnce.Do(func() {
		_MPSCNNNeuronPReLUClass = MPSCNNNeuronPReLUClass{class: objc.GetClass("MPSCNNNeuronPReLU")}
	})
	return _MPSCNNNeuronPReLUClass
}

// GetMPSCNNNeuronPReLUClass returns the class object for MPSCNNNeuronPReLU.
func GetMPSCNNNeuronPReLUClass() MPSCNNNeuronPReLUClass {
	return getMPSCNNNeuronPReLUClass()
}

type MPSCNNNeuronPReLUClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronPReLUClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronPReLUClass) Alloc() MPSCNNNeuronPReLU {
	rv := objc.Send[MPSCNNNeuronPReLU](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A parametric ReLU (Rectified Linear Unit) neuron filter.
//
// # Overview
//
// For each pixel in an image, the filter applies the following function:
//
// [media-2923194]
//
// Where `i` in `[0 ... channels - 1]`. That is, parameters `a“ᵢ` are
// learned and applied to each channel separately. Compare this to
// [MPSCNNNeuronReLU] where parameter `a` is shared across all channels.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPReLU
type MPSCNNNeuronPReLU struct {
	MPSCNNNeuron
}

// MPSCNNNeuronPReLUFromID constructs a [MPSCNNNeuronPReLU] from an objc.ID.
//
// A parametric ReLU (Rectified Linear Unit) neuron filter.
func MPSCNNNeuronPReLUFromID(id objc.ID) MPSCNNNeuronPReLU {
	return MPSCNNNeuronPReLU{MPSCNNNeuron: MPSCNNNeuronFromID(id)}
}

// NOTE: MPSCNNNeuronPReLU adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronPReLU] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPReLU
type IMPSCNNNeuronPReLU interface {
	IMPSCNNNeuron
}

// Init initializes the instance.
func (c MPSCNNNeuronPReLU) Init() MPSCNNNeuronPReLU {
	rv := objc.Send[MPSCNNNeuronPReLU](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronPReLU) Autorelease() MPSCNNNeuronPReLU {
	rv := objc.Send[MPSCNNNeuronPReLU](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronPReLU creates a new MPSCNNNeuronPReLU instance.
func NewMPSCNNNeuronPReLU() MPSCNNNeuronPReLU {
	class := getMPSCNNNeuronPReLUClass()
	rv := objc.Send[MPSCNNNeuronPReLU](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNNeuronPReLUWithCoder(aDecoder foundation.INSCoder) MPSCNNNeuronPReLU {
	instance := getMPSCNNNeuronPReLUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNNeuronPReLUFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(coder:device:)
func NewCNNNeuronPReLUWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNNeuronPReLU {
	instance := getMPSCNNNeuronPReLUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNNeuronPReLUFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNNeuronPReLUWithDevice(device metal.MTLDevice) MPSCNNNeuronPReLU {
	instance := getMPSCNNNeuronPReLUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNNeuronPReLUFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(device:neuronDescriptor:)
func NewCNNNeuronPReLUWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuronPReLU {
	instance := getMPSCNNNeuronPReLUClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	return MPSCNNNeuronPReLUFromID(rv)
}

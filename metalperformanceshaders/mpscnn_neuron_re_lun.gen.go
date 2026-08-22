// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronReLUN] class.
var (
	_MPSCNNNeuronReLUNClass     MPSCNNNeuronReLUNClass
	_MPSCNNNeuronReLUNClassOnce sync.Once
)

func getMPSCNNNeuronReLUNClass() MPSCNNNeuronReLUNClass {
	_MPSCNNNeuronReLUNClassOnce.Do(func() {
		_MPSCNNNeuronReLUNClass = MPSCNNNeuronReLUNClass{class: objc.GetClass("MPSCNNNeuronReLUN")}
	})
	return _MPSCNNNeuronReLUNClass
}

// GetMPSCNNNeuronReLUNClass returns the class object for MPSCNNNeuronReLUN.
func GetMPSCNNNeuronReLUNClass() MPSCNNNeuronReLUNClass {
	return getMPSCNNNeuronReLUNClass()
}

type MPSCNNNeuronReLUNClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronReLUNClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronReLUNClass) Alloc() MPSCNNNeuronReLUN {
	rv := objc.Send[MPSCNNNeuronReLUN](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A ReLUN neuron filter.
//
// # Overview
//
// For each pixel in an image, the filter applies the following function:
//
// [media-2923195]
//
// The default value of `a` is 1.0 and the default value of `b` is 6.0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUN
type MPSCNNNeuronReLUN struct {
	MPSCNNNeuron
}

// MPSCNNNeuronReLUNFromID constructs a [MPSCNNNeuronReLUN] from an objc.ID.
//
// A ReLUN neuron filter.
func MPSCNNNeuronReLUNFromID(id objc.ID) MPSCNNNeuronReLUN {
	return MPSCNNNeuronReLUN{MPSCNNNeuron: MPSCNNNeuronFromID(id)}
}

// NOTE: MPSCNNNeuronReLUN adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronReLUN] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUN
type IMPSCNNNeuronReLUN interface {
	IMPSCNNNeuron
}

// Init initializes the instance.
func (c MPSCNNNeuronReLUN) Init() MPSCNNNeuronReLUN {
	rv := objc.Send[MPSCNNNeuronReLUN](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronReLUN) Autorelease() MPSCNNNeuronReLUN {
	rv := objc.Send[MPSCNNNeuronReLUN](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronReLUN creates a new MPSCNNNeuronReLUN instance.
func NewMPSCNNNeuronReLUN() MPSCNNNeuronReLUN {
	class := getMPSCNNNeuronReLUNClass()
	rv := objc.Send[MPSCNNNeuronReLUN](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNNeuronReLUNWithCoder(aDecoder foundation.INSCoder) MPSCNNNeuronReLUN {
	instance := getMPSCNNNeuronReLUNClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNNeuronReLUNFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(coder:device:)
func NewCNNNeuronReLUNWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNNeuronReLUN {
	instance := getMPSCNNNeuronReLUNClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNNeuronReLUNFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNNeuronReLUNWithDevice(device metal.MTLDevice) MPSCNNNeuronReLUN {
	instance := getMPSCNNNeuronReLUNClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNNeuronReLUNFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(device:neuronDescriptor:)
func NewCNNNeuronReLUNWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuronReLUN {
	instance := getMPSCNNNeuronReLUNClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	return MPSCNNNeuronReLUNFromID(rv)
}

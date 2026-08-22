// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronHardSigmoid] class.
var (
	_MPSCNNNeuronHardSigmoidClass     MPSCNNNeuronHardSigmoidClass
	_MPSCNNNeuronHardSigmoidClassOnce sync.Once
)

func getMPSCNNNeuronHardSigmoidClass() MPSCNNNeuronHardSigmoidClass {
	_MPSCNNNeuronHardSigmoidClassOnce.Do(func() {
		_MPSCNNNeuronHardSigmoidClass = MPSCNNNeuronHardSigmoidClass{class: objc.GetClass("MPSCNNNeuronHardSigmoid")}
	})
	return _MPSCNNNeuronHardSigmoidClass
}

// GetMPSCNNNeuronHardSigmoidClass returns the class object for MPSCNNNeuronHardSigmoid.
func GetMPSCNNNeuronHardSigmoidClass() MPSCNNNeuronHardSigmoidClass {
	return getMPSCNNNeuronHardSigmoidClass()
}

type MPSCNNNeuronHardSigmoidClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronHardSigmoidClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronHardSigmoidClass) Alloc() MPSCNNNeuronHardSigmoid {
	rv := objc.Send[MPSCNNNeuronHardSigmoid](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A hard sigmoid neuron filter.
//
// # Overview
//
// For each pixel in an image, the filter applies the following function:
//
// [media-2903540]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronHardSigmoid
type MPSCNNNeuronHardSigmoid struct {
	MPSCNNNeuron
}

// MPSCNNNeuronHardSigmoidFromID constructs a [MPSCNNNeuronHardSigmoid] from an objc.ID.
//
// A hard sigmoid neuron filter.
func MPSCNNNeuronHardSigmoidFromID(id objc.ID) MPSCNNNeuronHardSigmoid {
	return MPSCNNNeuronHardSigmoid{MPSCNNNeuron: MPSCNNNeuronFromID(id)}
}

// NOTE: MPSCNNNeuronHardSigmoid adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronHardSigmoid] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronHardSigmoid
type IMPSCNNNeuronHardSigmoid interface {
	IMPSCNNNeuron
}

// Init initializes the instance.
func (c MPSCNNNeuronHardSigmoid) Init() MPSCNNNeuronHardSigmoid {
	rv := objc.Send[MPSCNNNeuronHardSigmoid](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronHardSigmoid) Autorelease() MPSCNNNeuronHardSigmoid {
	rv := objc.Send[MPSCNNNeuronHardSigmoid](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronHardSigmoid creates a new MPSCNNNeuronHardSigmoid instance.
func NewMPSCNNNeuronHardSigmoid() MPSCNNNeuronHardSigmoid {
	class := getMPSCNNNeuronHardSigmoidClass()
	rv := objc.Send[MPSCNNNeuronHardSigmoid](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNNeuronHardSigmoidWithCoder(aDecoder foundation.INSCoder) MPSCNNNeuronHardSigmoid {
	instance := getMPSCNNNeuronHardSigmoidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNNeuronHardSigmoidFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(coder:device:)
func NewCNNNeuronHardSigmoidWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNNeuronHardSigmoid {
	instance := getMPSCNNNeuronHardSigmoidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNNeuronHardSigmoidFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNNeuronHardSigmoidWithDevice(device metal.MTLDevice) MPSCNNNeuronHardSigmoid {
	instance := getMPSCNNNeuronHardSigmoidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNNeuronHardSigmoidFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(device:neuronDescriptor:)
func NewCNNNeuronHardSigmoidWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuronHardSigmoid {
	instance := getMPSCNNNeuronHardSigmoidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	return MPSCNNNeuronHardSigmoidFromID(rv)
}

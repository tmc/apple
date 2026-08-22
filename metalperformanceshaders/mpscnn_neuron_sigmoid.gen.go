// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronSigmoid] class.
var (
	_MPSCNNNeuronSigmoidClass     MPSCNNNeuronSigmoidClass
	_MPSCNNNeuronSigmoidClassOnce sync.Once
)

func getMPSCNNNeuronSigmoidClass() MPSCNNNeuronSigmoidClass {
	_MPSCNNNeuronSigmoidClassOnce.Do(func() {
		_MPSCNNNeuronSigmoidClass = MPSCNNNeuronSigmoidClass{class: objc.GetClass("MPSCNNNeuronSigmoid")}
	})
	return _MPSCNNNeuronSigmoidClass
}

// GetMPSCNNNeuronSigmoidClass returns the class object for MPSCNNNeuronSigmoid.
func GetMPSCNNNeuronSigmoidClass() MPSCNNNeuronSigmoidClass {
	return getMPSCNNNeuronSigmoidClass()
}

type MPSCNNNeuronSigmoidClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronSigmoidClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronSigmoidClass) Alloc() MPSCNNNeuronSigmoid {
	rv := objc.Send[MPSCNNNeuronSigmoid](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A sigmoid neuron filter.
//
// # Overview
//
// For each pixel in an image, the filter applies the following function:
//
// [media-2903545]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSigmoid
type MPSCNNNeuronSigmoid struct {
	MPSCNNNeuron
}

// MPSCNNNeuronSigmoidFromID constructs a [MPSCNNNeuronSigmoid] from an objc.ID.
//
// A sigmoid neuron filter.
func MPSCNNNeuronSigmoidFromID(id objc.ID) MPSCNNNeuronSigmoid {
	return MPSCNNNeuronSigmoid{MPSCNNNeuron: MPSCNNNeuronFromID(id)}
}

// NOTE: MPSCNNNeuronSigmoid adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronSigmoid] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSigmoid
type IMPSCNNNeuronSigmoid interface {
	IMPSCNNNeuron
}

// Init initializes the instance.
func (c MPSCNNNeuronSigmoid) Init() MPSCNNNeuronSigmoid {
	rv := objc.Send[MPSCNNNeuronSigmoid](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronSigmoid) Autorelease() MPSCNNNeuronSigmoid {
	rv := objc.Send[MPSCNNNeuronSigmoid](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronSigmoid creates a new MPSCNNNeuronSigmoid instance.
func NewMPSCNNNeuronSigmoid() MPSCNNNeuronSigmoid {
	class := getMPSCNNNeuronSigmoidClass()
	rv := objc.Send[MPSCNNNeuronSigmoid](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNNeuronSigmoidWithCoder(aDecoder foundation.INSCoder) MPSCNNNeuronSigmoid {
	instance := getMPSCNNNeuronSigmoidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNNeuronSigmoidFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(coder:device:)
func NewCNNNeuronSigmoidWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNNeuronSigmoid {
	instance := getMPSCNNNeuronSigmoidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNNeuronSigmoidFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(device:neuronDescriptor:)
func NewCNNNeuronSigmoidWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuronSigmoid {
	instance := getMPSCNNNeuronSigmoidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	return MPSCNNNeuronSigmoidFromID(rv)
}

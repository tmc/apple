// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronAbsolute] class.
var (
	_MPSCNNNeuronAbsoluteClass     MPSCNNNeuronAbsoluteClass
	_MPSCNNNeuronAbsoluteClassOnce sync.Once
)

func getMPSCNNNeuronAbsoluteClass() MPSCNNNeuronAbsoluteClass {
	_MPSCNNNeuronAbsoluteClassOnce.Do(func() {
		_MPSCNNNeuronAbsoluteClass = MPSCNNNeuronAbsoluteClass{class: objc.GetClass("MPSCNNNeuronAbsolute")}
	})
	return _MPSCNNNeuronAbsoluteClass
}

// GetMPSCNNNeuronAbsoluteClass returns the class object for MPSCNNNeuronAbsolute.
func GetMPSCNNNeuronAbsoluteClass() MPSCNNNeuronAbsoluteClass {
	return getMPSCNNNeuronAbsoluteClass()
}

type MPSCNNNeuronAbsoluteClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronAbsoluteClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronAbsoluteClass) Alloc() MPSCNNNeuronAbsolute {
	rv := objc.Send[MPSCNNNeuronAbsolute](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An absolute neuron filter.
//
// # Overview
//
// For each pixel in an image, the filter applies the following function:
//
// [media-2903538]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronAbsolute
type MPSCNNNeuronAbsolute struct {
	MPSCNNNeuron
}

// MPSCNNNeuronAbsoluteFromID constructs a [MPSCNNNeuronAbsolute] from an objc.ID.
//
// An absolute neuron filter.
func MPSCNNNeuronAbsoluteFromID(id objc.ID) MPSCNNNeuronAbsolute {
	return MPSCNNNeuronAbsolute{MPSCNNNeuron: MPSCNNNeuronFromID(id)}
}

// NOTE: MPSCNNNeuronAbsolute adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronAbsolute] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronAbsolute
type IMPSCNNNeuronAbsolute interface {
	IMPSCNNNeuron
}

// Init initializes the instance.
func (c MPSCNNNeuronAbsolute) Init() MPSCNNNeuronAbsolute {
	rv := objc.Send[MPSCNNNeuronAbsolute](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronAbsolute) Autorelease() MPSCNNNeuronAbsolute {
	rv := objc.Send[MPSCNNNeuronAbsolute](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronAbsolute creates a new MPSCNNNeuronAbsolute instance.
func NewMPSCNNNeuronAbsolute() MPSCNNNeuronAbsolute {
	class := getMPSCNNNeuronAbsoluteClass()
	rv := objc.Send[MPSCNNNeuronAbsolute](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNNeuronAbsoluteWithCoder(aDecoder foundation.INSCoder) MPSCNNNeuronAbsolute {
	instance := getMPSCNNNeuronAbsoluteClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNNeuronAbsoluteFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(coder:device:)
func NewCNNNeuronAbsoluteWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNNeuronAbsolute {
	instance := getMPSCNNNeuronAbsoluteClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNNeuronAbsoluteFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(device:neuronDescriptor:)
func NewCNNNeuronAbsoluteWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuronAbsolute {
	instance := getMPSCNNNeuronAbsoluteClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	return MPSCNNNeuronAbsoluteFromID(rv)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronLogarithm] class.
var (
	_MPSCNNNeuronLogarithmClass     MPSCNNNeuronLogarithmClass
	_MPSCNNNeuronLogarithmClassOnce sync.Once
)

func getMPSCNNNeuronLogarithmClass() MPSCNNNeuronLogarithmClass {
	_MPSCNNNeuronLogarithmClassOnce.Do(func() {
		_MPSCNNNeuronLogarithmClass = MPSCNNNeuronLogarithmClass{class: objc.GetClass("MPSCNNNeuronLogarithm")}
	})
	return _MPSCNNNeuronLogarithmClass
}

// GetMPSCNNNeuronLogarithmClass returns the class object for MPSCNNNeuronLogarithm.
func GetMPSCNNNeuronLogarithmClass() MPSCNNNeuronLogarithmClass {
	return getMPSCNNNeuronLogarithmClass()
}

type MPSCNNNeuronLogarithmClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronLogarithmClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronLogarithmClass) Alloc() MPSCNNNeuronLogarithm {
	rv := objc.Send[MPSCNNNeuronLogarithm](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A logarithm neuron filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLogarithm
type MPSCNNNeuronLogarithm struct {
	MPSCNNNeuron
}

// MPSCNNNeuronLogarithmFromID constructs a [MPSCNNNeuronLogarithm] from an objc.ID.
//
// A logarithm neuron filter.
func MPSCNNNeuronLogarithmFromID(id objc.ID) MPSCNNNeuronLogarithm {
	return MPSCNNNeuronLogarithm{MPSCNNNeuron: MPSCNNNeuronFromID(id)}
}

// NOTE: MPSCNNNeuronLogarithm adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronLogarithm] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLogarithm
type IMPSCNNNeuronLogarithm interface {
	IMPSCNNNeuron
}

// Init initializes the instance.
func (c MPSCNNNeuronLogarithm) Init() MPSCNNNeuronLogarithm {
	rv := objc.Send[MPSCNNNeuronLogarithm](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronLogarithm) Autorelease() MPSCNNNeuronLogarithm {
	rv := objc.Send[MPSCNNNeuronLogarithm](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronLogarithm creates a new MPSCNNNeuronLogarithm instance.
func NewMPSCNNNeuronLogarithm() MPSCNNNeuronLogarithm {
	class := getMPSCNNNeuronLogarithmClass()
	rv := objc.Send[MPSCNNNeuronLogarithm](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNNeuronLogarithmWithCoder(aDecoder foundation.INSCoder) MPSCNNNeuronLogarithm {
	instance := getMPSCNNNeuronLogarithmClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNNeuronLogarithmFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(coder:device:)
func NewCNNNeuronLogarithmWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNNeuronLogarithm {
	instance := getMPSCNNNeuronLogarithmClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNNeuronLogarithmFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNNeuronLogarithmWithDevice(device metal.MTLDevice) MPSCNNNeuronLogarithm {
	instance := getMPSCNNNeuronLogarithmClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNNeuronLogarithmFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(device:neuronDescriptor:)
func NewCNNNeuronLogarithmWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuronLogarithm {
	instance := getMPSCNNNeuronLogarithmClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	return MPSCNNNeuronLogarithmFromID(rv)
}

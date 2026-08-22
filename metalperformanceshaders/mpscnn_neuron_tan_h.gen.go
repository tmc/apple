// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronTanH] class.
var (
	_MPSCNNNeuronTanHClass     MPSCNNNeuronTanHClass
	_MPSCNNNeuronTanHClassOnce sync.Once
)

func getMPSCNNNeuronTanHClass() MPSCNNNeuronTanHClass {
	_MPSCNNNeuronTanHClassOnce.Do(func() {
		_MPSCNNNeuronTanHClass = MPSCNNNeuronTanHClass{class: objc.GetClass("MPSCNNNeuronTanH")}
	})
	return _MPSCNNNeuronTanHClass
}

// GetMPSCNNNeuronTanHClass returns the class object for MPSCNNNeuronTanH.
func GetMPSCNNNeuronTanHClass() MPSCNNNeuronTanHClass {
	return getMPSCNNNeuronTanHClass()
}

type MPSCNNNeuronTanHClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronTanHClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronTanHClass) Alloc() MPSCNNNeuronTanH {
	rv := objc.Send[MPSCNNNeuronTanH](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A hyperbolic tangent neuron filter.
//
// # Overview
//
// For each pixel in an image, the filter applies the following function:
//
// [media-2903548]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronTanH
type MPSCNNNeuronTanH struct {
	MPSCNNNeuron
}

// MPSCNNNeuronTanHFromID constructs a [MPSCNNNeuronTanH] from an objc.ID.
//
// A hyperbolic tangent neuron filter.
func MPSCNNNeuronTanHFromID(id objc.ID) MPSCNNNeuronTanH {
	return MPSCNNNeuronTanH{MPSCNNNeuron: MPSCNNNeuronFromID(id)}
}

// NOTE: MPSCNNNeuronTanH adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronTanH] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronTanH
type IMPSCNNNeuronTanH interface {
	IMPSCNNNeuron
}

// Init initializes the instance.
func (c MPSCNNNeuronTanH) Init() MPSCNNNeuronTanH {
	rv := objc.Send[MPSCNNNeuronTanH](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronTanH) Autorelease() MPSCNNNeuronTanH {
	rv := objc.Send[MPSCNNNeuronTanH](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronTanH creates a new MPSCNNNeuronTanH instance.
func NewMPSCNNNeuronTanH() MPSCNNNeuronTanH {
	class := getMPSCNNNeuronTanHClass()
	rv := objc.Send[MPSCNNNeuronTanH](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNNeuronTanHWithCoder(aDecoder foundation.INSCoder) MPSCNNNeuronTanH {
	instance := getMPSCNNNeuronTanHClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNNeuronTanHFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(coder:device:)
func NewCNNNeuronTanHWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNNeuronTanH {
	instance := getMPSCNNNeuronTanHClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNNeuronTanHFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNNeuronTanHWithDevice(device metal.MTLDevice) MPSCNNNeuronTanH {
	instance := getMPSCNNNeuronTanHClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNNeuronTanHFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(device:neuronDescriptor:)
func NewCNNNeuronTanHWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuronTanH {
	instance := getMPSCNNNeuronTanHClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	return MPSCNNNeuronTanHFromID(rv)
}

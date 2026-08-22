// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronSoftPlus] class.
var (
	_MPSCNNNeuronSoftPlusClass     MPSCNNNeuronSoftPlusClass
	_MPSCNNNeuronSoftPlusClassOnce sync.Once
)

func getMPSCNNNeuronSoftPlusClass() MPSCNNNeuronSoftPlusClass {
	_MPSCNNNeuronSoftPlusClassOnce.Do(func() {
		_MPSCNNNeuronSoftPlusClass = MPSCNNNeuronSoftPlusClass{class: objc.GetClass("MPSCNNNeuronSoftPlus")}
	})
	return _MPSCNNNeuronSoftPlusClass
}

// GetMPSCNNNeuronSoftPlusClass returns the class object for MPSCNNNeuronSoftPlus.
func GetMPSCNNNeuronSoftPlusClass() MPSCNNNeuronSoftPlusClass {
	return getMPSCNNNeuronSoftPlusClass()
}

type MPSCNNNeuronSoftPlusClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronSoftPlusClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronSoftPlusClass) Alloc() MPSCNNNeuronSoftPlus {
	rv := objc.Send[MPSCNNNeuronSoftPlus](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A parametric softplus neuron filter.
//
// # Overview
//
// For each pixel in an image, the filter applies the following function:
//
// [media-2903546]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftPlus
type MPSCNNNeuronSoftPlus struct {
	MPSCNNNeuron
}

// MPSCNNNeuronSoftPlusFromID constructs a [MPSCNNNeuronSoftPlus] from an objc.ID.
//
// A parametric softplus neuron filter.
func MPSCNNNeuronSoftPlusFromID(id objc.ID) MPSCNNNeuronSoftPlus {
	return MPSCNNNeuronSoftPlus{MPSCNNNeuron: MPSCNNNeuronFromID(id)}
}

// NOTE: MPSCNNNeuronSoftPlus adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronSoftPlus] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftPlus
type IMPSCNNNeuronSoftPlus interface {
	IMPSCNNNeuron
}

// Init initializes the instance.
func (c MPSCNNNeuronSoftPlus) Init() MPSCNNNeuronSoftPlus {
	rv := objc.Send[MPSCNNNeuronSoftPlus](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronSoftPlus) Autorelease() MPSCNNNeuronSoftPlus {
	rv := objc.Send[MPSCNNNeuronSoftPlus](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronSoftPlus creates a new MPSCNNNeuronSoftPlus instance.
func NewMPSCNNNeuronSoftPlus() MPSCNNNeuronSoftPlus {
	class := getMPSCNNNeuronSoftPlusClass()
	rv := objc.Send[MPSCNNNeuronSoftPlus](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNNeuronSoftPlusWithCoder(aDecoder foundation.INSCoder) MPSCNNNeuronSoftPlus {
	instance := getMPSCNNNeuronSoftPlusClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNNeuronSoftPlusFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(coder:device:)
func NewCNNNeuronSoftPlusWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNNeuronSoftPlus {
	instance := getMPSCNNNeuronSoftPlusClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNNeuronSoftPlusFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNNeuronSoftPlusWithDevice(device metal.MTLDevice) MPSCNNNeuronSoftPlus {
	instance := getMPSCNNNeuronSoftPlusClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNNeuronSoftPlusFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(device:neuronDescriptor:)
func NewCNNNeuronSoftPlusWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuronSoftPlus {
	instance := getMPSCNNNeuronSoftPlusClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	return MPSCNNNeuronSoftPlusFromID(rv)
}

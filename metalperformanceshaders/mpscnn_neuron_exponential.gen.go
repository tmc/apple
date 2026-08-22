// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronExponential] class.
var (
	_MPSCNNNeuronExponentialClass     MPSCNNNeuronExponentialClass
	_MPSCNNNeuronExponentialClassOnce sync.Once
)

func getMPSCNNNeuronExponentialClass() MPSCNNNeuronExponentialClass {
	_MPSCNNNeuronExponentialClassOnce.Do(func() {
		_MPSCNNNeuronExponentialClass = MPSCNNNeuronExponentialClass{class: objc.GetClass("MPSCNNNeuronExponential")}
	})
	return _MPSCNNNeuronExponentialClass
}

// GetMPSCNNNeuronExponentialClass returns the class object for MPSCNNNeuronExponential.
func GetMPSCNNNeuronExponentialClass() MPSCNNNeuronExponentialClass {
	return getMPSCNNNeuronExponentialClass()
}

type MPSCNNNeuronExponentialClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronExponentialClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronExponentialClass) Alloc() MPSCNNNeuronExponential {
	rv := objc.Send[MPSCNNNeuronExponential](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An exponential neuron filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronExponential
type MPSCNNNeuronExponential struct {
	MPSCNNNeuron
}

// MPSCNNNeuronExponentialFromID constructs a [MPSCNNNeuronExponential] from an objc.ID.
//
// An exponential neuron filter.
func MPSCNNNeuronExponentialFromID(id objc.ID) MPSCNNNeuronExponential {
	return MPSCNNNeuronExponential{MPSCNNNeuron: MPSCNNNeuronFromID(id)}
}

// NOTE: MPSCNNNeuronExponential adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronExponential] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronExponential
type IMPSCNNNeuronExponential interface {
	IMPSCNNNeuron
}

// Init initializes the instance.
func (c MPSCNNNeuronExponential) Init() MPSCNNNeuronExponential {
	rv := objc.Send[MPSCNNNeuronExponential](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronExponential) Autorelease() MPSCNNNeuronExponential {
	rv := objc.Send[MPSCNNNeuronExponential](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronExponential creates a new MPSCNNNeuronExponential instance.
func NewMPSCNNNeuronExponential() MPSCNNNeuronExponential {
	class := getMPSCNNNeuronExponentialClass()
	rv := objc.Send[MPSCNNNeuronExponential](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNNeuronExponentialWithCoder(aDecoder foundation.INSCoder) MPSCNNNeuronExponential {
	instance := getMPSCNNNeuronExponentialClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNNeuronExponentialFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(coder:device:)
func NewCNNNeuronExponentialWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNNeuronExponential {
	instance := getMPSCNNNeuronExponentialClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNNeuronExponentialFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNNeuronExponentialWithDevice(device metal.MTLDevice) MPSCNNNeuronExponential {
	instance := getMPSCNNNeuronExponentialClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNNeuronExponentialFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(device:neuronDescriptor:)
func NewCNNNeuronExponentialWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuronExponential {
	instance := getMPSCNNNeuronExponentialClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	return MPSCNNNeuronExponentialFromID(rv)
}

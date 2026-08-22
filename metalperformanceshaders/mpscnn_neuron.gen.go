// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuron] class.
var (
	_MPSCNNNeuronClass     MPSCNNNeuronClass
	_MPSCNNNeuronClassOnce sync.Once
)

func getMPSCNNNeuronClass() MPSCNNNeuronClass {
	_MPSCNNNeuronClassOnce.Do(func() {
		_MPSCNNNeuronClass = MPSCNNNeuronClass{class: objc.GetClass("MPSCNNNeuron")}
	})
	return _MPSCNNNeuronClass
}

// GetMPSCNNNeuronClass returns the class object for MPSCNNNeuron.
func GetMPSCNNNeuronClass() MPSCNNNeuronClass {
	return getMPSCNNNeuronClass()
}

type MPSCNNNeuronClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronClass) Alloc() MPSCNNNeuron {
	rv := objc.Send[MPSCNNNeuron](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that applies a neuron activation function.
//
// # Overview
//
// Do not use this class directly; use one of the [MPSCNNNeuron] subclasses
// instead.
//
// # Initializers
//
//   - [MPSCNNNeuron.InitWithDeviceNeuronDescriptor]
//
// # Instance Properties
//
//   - [MPSCNNNeuron.A]
//   - [MPSCNNNeuron.B]
//   - [MPSCNNNeuron.C]
//   - [MPSCNNNeuron.Data]
//   - [MPSCNNNeuron.NeuronType]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron
type MPSCNNNeuron struct {
	MPSCNNKernel
}

// MPSCNNNeuronFromID constructs a [MPSCNNNeuron] from an objc.ID.
//
// A filter that applies a neuron activation function.
func MPSCNNNeuronFromID(id objc.ID) MPSCNNNeuron {
	return MPSCNNNeuron{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSCNNNeuron adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuron] class.
//
// # Initializers
//
//   - [IMPSCNNNeuron.InitWithDeviceNeuronDescriptor]
//
// # Instance Properties
//
//   - [IMPSCNNNeuron.A]
//   - [IMPSCNNNeuron.B]
//   - [IMPSCNNNeuron.C]
//   - [IMPSCNNNeuron.Data]
//   - [IMPSCNNNeuron.NeuronType]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron
type IMPSCNNNeuron interface {
	IMPSCNNKernel

	// Topic: Initializers

	InitWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuron

	// Topic: Instance Properties

	A() float32
	B() float32
	C() float32
	Data() foundation.NSData
	NeuronType() MPSCNNNeuronType
}

// Init initializes the instance.
func (c MPSCNNNeuron) Init() MPSCNNNeuron {
	rv := objc.Send[MPSCNNNeuron](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuron) Autorelease() MPSCNNNeuron {
	rv := objc.Send[MPSCNNNeuron](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuron creates a new MPSCNNNeuron instance.
func NewMPSCNNNeuron() MPSCNNNeuron {
	class := getMPSCNNNeuronClass()
	rv := objc.Send[MPSCNNNeuron](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNNeuronWithCoder(aDecoder foundation.INSCoder) MPSCNNNeuron {
	instance := getMPSCNNNeuronClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNNeuronFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(coder:device:)
func NewCNNNeuronWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNNeuron {
	instance := getMPSCNNNeuronClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNNeuronFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNNeuronWithDevice(device metal.MTLDevice) MPSCNNNeuron {
	instance := getMPSCNNNeuronClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNNeuronFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(device:neuronDescriptor:)
func NewCNNNeuronWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuron {
	instance := getMPSCNNNeuronClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	return MPSCNNNeuronFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(device:neuronDescriptor:)
func (c MPSCNNNeuron) InitWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuron {
	rv := objc.Send[MPSCNNNeuron](c.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/a
func (c MPSCNNNeuron) A() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("a"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/b
func (c MPSCNNNeuron) B() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("b"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/c
func (c MPSCNNNeuron) C() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("c"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/data
func (c MPSCNNNeuron) Data() foundation.NSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("data"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/neuronType
func (c MPSCNNNeuron) NeuronType() MPSCNNNeuronType {
	rv := objc.Send[MPSCNNNeuronType](c.ID, objc.Sel("neuronType"))
	return MPSCNNNeuronType(rv)
}

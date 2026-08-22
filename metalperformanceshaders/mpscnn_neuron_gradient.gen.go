// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronGradient] class.
var (
	_MPSCNNNeuronGradientClass     MPSCNNNeuronGradientClass
	_MPSCNNNeuronGradientClassOnce sync.Once
)

func getMPSCNNNeuronGradientClass() MPSCNNNeuronGradientClass {
	_MPSCNNNeuronGradientClassOnce.Do(func() {
		_MPSCNNNeuronGradientClass = MPSCNNNeuronGradientClass{class: objc.GetClass("MPSCNNNeuronGradient")}
	})
	return _MPSCNNNeuronGradientClass
}

// GetMPSCNNNeuronGradientClass returns the class object for MPSCNNNeuronGradient.
func GetMPSCNNNeuronGradientClass() MPSCNNNeuronGradientClass {
	return getMPSCNNNeuronGradientClass()
}

type MPSCNNNeuronGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronGradientClass) Alloc() MPSCNNNeuronGradient {
	rv := objc.Send[MPSCNNNeuronGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient neuron filter.
//
// # Initializers
//
//   - [MPSCNNNeuronGradient.InitWithDeviceNeuronDescriptor]
//
// # Instance Properties
//
//   - [MPSCNNNeuronGradient.A]
//   - [MPSCNNNeuronGradient.B]
//   - [MPSCNNNeuronGradient.C]
//   - [MPSCNNNeuronGradient.Data]
//   - [MPSCNNNeuronGradient.NeuronType]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGradient
type MPSCNNNeuronGradient struct {
	MPSCNNGradientKernel
}

// MPSCNNNeuronGradientFromID constructs a [MPSCNNNeuronGradient] from an objc.ID.
//
// A gradient neuron filter.
func MPSCNNNeuronGradientFromID(id objc.ID) MPSCNNNeuronGradient {
	return MPSCNNNeuronGradient{MPSCNNGradientKernel: MPSCNNGradientKernelFromID(id)}
}

// NOTE: MPSCNNNeuronGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronGradient] class.
//
// # Initializers
//
//   - [IMPSCNNNeuronGradient.InitWithDeviceNeuronDescriptor]
//
// # Instance Properties
//
//   - [IMPSCNNNeuronGradient.A]
//   - [IMPSCNNNeuronGradient.B]
//   - [IMPSCNNNeuronGradient.C]
//   - [IMPSCNNNeuronGradient.Data]
//   - [IMPSCNNNeuronGradient.NeuronType]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGradient
type IMPSCNNNeuronGradient interface {
	IMPSCNNGradientKernel

	// Topic: Initializers

	InitWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuronGradient

	// Topic: Instance Properties

	A() float32
	B() float32
	C() float32
	Data() foundation.NSData
	NeuronType() MPSCNNNeuronType
}

// Init initializes the instance.
func (c MPSCNNNeuronGradient) Init() MPSCNNNeuronGradient {
	rv := objc.Send[MPSCNNNeuronGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronGradient) Autorelease() MPSCNNNeuronGradient {
	rv := objc.Send[MPSCNNNeuronGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronGradient creates a new MPSCNNNeuronGradient instance.
func NewMPSCNNNeuronGradient() MPSCNNNeuronGradient {
	class := getMPSCNNNeuronGradientClass()
	rv := objc.Send[MPSCNNNeuronGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNNeuronGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNNeuronGradient {
	instance := getMPSCNNNeuronGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNNeuronGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGradient/init(coder:device:)
func NewCNNNeuronGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNNeuronGradient {
	instance := getMPSCNNNeuronGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNNeuronGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNNeuronGradientWithDevice(device metal.MTLDevice) MPSCNNNeuronGradient {
	instance := getMPSCNNNeuronGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNNeuronGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGradient/init(device:neuronDescriptor:)
func NewCNNNeuronGradientWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuronGradient {
	instance := getMPSCNNNeuronGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	return MPSCNNNeuronGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGradient/init(device:neuronDescriptor:)
func (c MPSCNNNeuronGradient) InitWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuronGradient {
	rv := objc.Send[MPSCNNNeuronGradient](c.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGradient/a
func (c MPSCNNNeuronGradient) A() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("a"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGradient/b
func (c MPSCNNNeuronGradient) B() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("b"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGradient/c
func (c MPSCNNNeuronGradient) C() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("c"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGradient/data
func (c MPSCNNNeuronGradient) Data() foundation.NSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("data"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGradient/neuronType
func (c MPSCNNNeuronGradient) NeuronType() MPSCNNNeuronType {
	rv := objc.Send[MPSCNNNeuronType](c.ID, objc.Sel("neuronType"))
	return MPSCNNNeuronType(rv)
}

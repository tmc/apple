// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNNeuronSoftSign] class.
var (
	_MPSCNNNeuronSoftSignClass     MPSCNNNeuronSoftSignClass
	_MPSCNNNeuronSoftSignClassOnce sync.Once
)

func getMPSCNNNeuronSoftSignClass() MPSCNNNeuronSoftSignClass {
	_MPSCNNNeuronSoftSignClassOnce.Do(func() {
		_MPSCNNNeuronSoftSignClass = MPSCNNNeuronSoftSignClass{class: objc.GetClass("MPSCNNNeuronSoftSign")}
	})
	return _MPSCNNNeuronSoftSignClass
}

// GetMPSCNNNeuronSoftSignClass returns the class object for MPSCNNNeuronSoftSign.
func GetMPSCNNNeuronSoftSignClass() MPSCNNNeuronSoftSignClass {
	return getMPSCNNNeuronSoftSignClass()
}

type MPSCNNNeuronSoftSignClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNNeuronSoftSignClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNNeuronSoftSignClass) Alloc() MPSCNNNeuronSoftSign {
	rv := objc.Send[MPSCNNNeuronSoftSign](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A softsign neuron filter.
//
// # Overview
//
// For each pixel in an image, the filter applies the following function:
//
// [media-2903547]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftSign
type MPSCNNNeuronSoftSign struct {
	MPSCNNNeuron
}

// MPSCNNNeuronSoftSignFromID constructs a [MPSCNNNeuronSoftSign] from an objc.ID.
//
// A softsign neuron filter.
func MPSCNNNeuronSoftSignFromID(id objc.ID) MPSCNNNeuronSoftSign {
	return MPSCNNNeuronSoftSign{MPSCNNNeuron: MPSCNNNeuronFromID(id)}
}

// NOTE: MPSCNNNeuronSoftSign adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNNeuronSoftSign] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftSign
type IMPSCNNNeuronSoftSign interface {
	IMPSCNNNeuron
}

// Init initializes the instance.
func (c MPSCNNNeuronSoftSign) Init() MPSCNNNeuronSoftSign {
	rv := objc.Send[MPSCNNNeuronSoftSign](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNNeuronSoftSign) Autorelease() MPSCNNNeuronSoftSign {
	rv := objc.Send[MPSCNNNeuronSoftSign](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNNeuronSoftSign creates a new MPSCNNNeuronSoftSign instance.
func NewMPSCNNNeuronSoftSign() MPSCNNNeuronSoftSign {
	class := getMPSCNNNeuronSoftSignClass()
	rv := objc.Send[MPSCNNNeuronSoftSign](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNNeuronSoftSignWithCoder(aDecoder foundation.INSCoder) MPSCNNNeuronSoftSign {
	instance := getMPSCNNNeuronSoftSignClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNNeuronSoftSignFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(coder:device:)
func NewCNNNeuronSoftSignWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNNeuronSoftSign {
	instance := getMPSCNNNeuronSoftSignClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNNeuronSoftSignFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron/init(device:neuronDescriptor:)
func NewCNNNeuronSoftSignWithDeviceNeuronDescriptor(device metal.MTLDevice, neuronDescriptor IMPSNNNeuronDescriptor) MPSCNNNeuronSoftSign {
	instance := getMPSCNNNeuronSoftSignClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	return MPSCNNNeuronSoftSignFromID(rv)
}

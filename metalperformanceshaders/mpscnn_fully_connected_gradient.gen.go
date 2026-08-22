// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNFullyConnectedGradient] class.
var (
	_MPSCNNFullyConnectedGradientClass     MPSCNNFullyConnectedGradientClass
	_MPSCNNFullyConnectedGradientClassOnce sync.Once
)

func getMPSCNNFullyConnectedGradientClass() MPSCNNFullyConnectedGradientClass {
	_MPSCNNFullyConnectedGradientClassOnce.Do(func() {
		_MPSCNNFullyConnectedGradientClass = MPSCNNFullyConnectedGradientClass{class: objc.GetClass("MPSCNNFullyConnectedGradient")}
	})
	return _MPSCNNFullyConnectedGradientClass
}

// GetMPSCNNFullyConnectedGradientClass returns the class object for MPSCNNFullyConnectedGradient.
func GetMPSCNNFullyConnectedGradientClass() MPSCNNFullyConnectedGradientClass {
	return getMPSCNNFullyConnectedGradientClass()
}

type MPSCNNFullyConnectedGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNFullyConnectedGradientClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNFullyConnectedGradientClass) Alloc() MPSCNNFullyConnectedGradient {
	rv := objc.Send[MPSCNNFullyConnectedGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A gradient fully connected convolution layer.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnectedGradient
type MPSCNNFullyConnectedGradient struct {
	MPSCNNConvolutionGradient
}

// MPSCNNFullyConnectedGradientFromID constructs a [MPSCNNFullyConnectedGradient] from an objc.ID.
//
// A gradient fully connected convolution layer.
func MPSCNNFullyConnectedGradientFromID(id objc.ID) MPSCNNFullyConnectedGradient {
	return MPSCNNFullyConnectedGradient{MPSCNNConvolutionGradient: MPSCNNConvolutionGradientFromID(id)}
}

// NOTE: MPSCNNFullyConnectedGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNFullyConnectedGradient] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnectedGradient
type IMPSCNNFullyConnectedGradient interface {
	IMPSCNNConvolutionGradient
}

// Init initializes the instance.
func (c MPSCNNFullyConnectedGradient) Init() MPSCNNFullyConnectedGradient {
	rv := objc.Send[MPSCNNFullyConnectedGradient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNFullyConnectedGradient) Autorelease() MPSCNNFullyConnectedGradient {
	rv := objc.Send[MPSCNNFullyConnectedGradient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNFullyConnectedGradient creates a new MPSCNNFullyConnectedGradient instance.
func NewMPSCNNFullyConnectedGradient() MPSCNNFullyConnectedGradient {
	class := getMPSCNNFullyConnectedGradientClass()
	rv := objc.Send[MPSCNNFullyConnectedGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNFullyConnectedGradientWithCoder(aDecoder foundation.INSCoder) MPSCNNFullyConnectedGradient {
	instance := getMPSCNNFullyConnectedGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNFullyConnectedGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnectedGradient/init(coder:device:)
func NewCNNFullyConnectedGradientWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNFullyConnectedGradient {
	instance := getMPSCNNFullyConnectedGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNFullyConnectedGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel/init(device:)
func NewCNNFullyConnectedGradientWithDevice(device metal.MTLDevice) MPSCNNFullyConnectedGradient {
	instance := getMPSCNNFullyConnectedGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNFullyConnectedGradientFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnectedGradient/init(device:weights:)
func NewCNNFullyConnectedGradientWithDeviceWeights(device metal.MTLDevice, weights MPSCNNConvolutionDataSource) MPSCNNFullyConnectedGradient {
	instance := getMPSCNNFullyConnectedGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:weights:"), device, weights)
	return MPSCNNFullyConnectedGradientFromID(rv)
}

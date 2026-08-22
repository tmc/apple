// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNLogSoftMax] class.
var (
	_MPSCNNLogSoftMaxClass     MPSCNNLogSoftMaxClass
	_MPSCNNLogSoftMaxClassOnce sync.Once
)

func getMPSCNNLogSoftMaxClass() MPSCNNLogSoftMaxClass {
	_MPSCNNLogSoftMaxClassOnce.Do(func() {
		_MPSCNNLogSoftMaxClass = MPSCNNLogSoftMaxClass{class: objc.GetClass("MPSCNNLogSoftMax")}
	})
	return _MPSCNNLogSoftMaxClass
}

// GetMPSCNNLogSoftMaxClass returns the class object for MPSCNNLogSoftMax.
func GetMPSCNNLogSoftMaxClass() MPSCNNLogSoftMaxClass {
	return getMPSCNNLogSoftMaxClass()
}

type MPSCNNLogSoftMaxClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNLogSoftMaxClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNLogSoftMaxClass) Alloc() MPSCNNLogSoftMax {
	rv := objc.Send[MPSCNNLogSoftMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A neural transfer function that is useful for constructing a loss function
// to be minimized when training neural networks.
//
// # Overview
//
// The logarithmic softmax filter is calculated by taking the natural
// logarithm of the result of a softmax filter.
//
// For each feature channel per pixel in an image in a feature map, the
// logarithmic softmax filter computes the following:
//
// [media-2903560]
//
// Where [R] is the result channel in the pixel, [N] is the number of feature
// channels, and `y=ln(x)` satisfies `e“ʸ“=x`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMax
type MPSCNNLogSoftMax struct {
	MPSCNNKernel
}

// MPSCNNLogSoftMaxFromID constructs a [MPSCNNLogSoftMax] from an objc.ID.
//
// A neural transfer function that is useful for constructing a loss function
// to be minimized when training neural networks.
func MPSCNNLogSoftMaxFromID(id objc.ID) MPSCNNLogSoftMax {
	return MPSCNNLogSoftMax{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSCNNLogSoftMax adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNLogSoftMax] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMax
type IMPSCNNLogSoftMax interface {
	IMPSCNNKernel
}

// Init initializes the instance.
func (c MPSCNNLogSoftMax) Init() MPSCNNLogSoftMax {
	rv := objc.Send[MPSCNNLogSoftMax](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNLogSoftMax) Autorelease() MPSCNNLogSoftMax {
	rv := objc.Send[MPSCNNLogSoftMax](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNLogSoftMax creates a new MPSCNNLogSoftMax instance.
func NewMPSCNNLogSoftMax() MPSCNNLogSoftMax {
	class := getMPSCNNLogSoftMaxClass()
	rv := objc.Send[MPSCNNLogSoftMax](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNLogSoftMaxWithCoder(aDecoder foundation.INSCoder) MPSCNNLogSoftMax {
	instance := getMPSCNNLogSoftMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNLogSoftMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(coder:device:)
func NewCNNLogSoftMaxWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNLogSoftMax {
	instance := getMPSCNNLogSoftMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNLogSoftMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNLogSoftMaxWithDevice(device metal.MTLDevice) MPSCNNLogSoftMax {
	instance := getMPSCNNLogSoftMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNLogSoftMaxFromID(rv)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNSoftMax] class.
var (
	_MPSCNNSoftMaxClass     MPSCNNSoftMaxClass
	_MPSCNNSoftMaxClassOnce sync.Once
)

func getMPSCNNSoftMaxClass() MPSCNNSoftMaxClass {
	_MPSCNNSoftMaxClassOnce.Do(func() {
		_MPSCNNSoftMaxClass = MPSCNNSoftMaxClass{class: objc.GetClass("MPSCNNSoftMax")}
	})
	return _MPSCNNSoftMaxClass
}

// GetMPSCNNSoftMaxClass returns the class object for MPSCNNSoftMax.
func GetMPSCNNSoftMaxClass() MPSCNNSoftMaxClass {
	return getMPSCNNSoftMaxClass()
}

type MPSCNNSoftMaxClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNSoftMaxClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNSoftMaxClass) Alloc() MPSCNNSoftMax {
	rv := objc.Send[MPSCNNSoftMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A neural transfer function that is useful for classification tasks.
//
// # Overview
//
// The softmax filter is applied across feature channels in a convolutional
// manner at all spatial locations. The softmax filter can be seen as the
// combination of an activation function (exponential) and a normalization
// operator.
//
// For each feature channel per pixel in an image in a feature map, the
// softmax filter computes the following:
//
// [media-2903559]
//
// Where [R] is the result channel in the pixel and [N] is the number of
// feature channels.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMax
type MPSCNNSoftMax struct {
	MPSCNNKernel
}

// MPSCNNSoftMaxFromID constructs a [MPSCNNSoftMax] from an objc.ID.
//
// A neural transfer function that is useful for classification tasks.
func MPSCNNSoftMaxFromID(id objc.ID) MPSCNNSoftMax {
	return MPSCNNSoftMax{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSCNNSoftMax adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNSoftMax] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMax
type IMPSCNNSoftMax interface {
	IMPSCNNKernel
}

// Init initializes the instance.
func (c MPSCNNSoftMax) Init() MPSCNNSoftMax {
	rv := objc.Send[MPSCNNSoftMax](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNSoftMax) Autorelease() MPSCNNSoftMax {
	rv := objc.Send[MPSCNNSoftMax](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNSoftMax creates a new MPSCNNSoftMax instance.
func NewMPSCNNSoftMax() MPSCNNSoftMax {
	class := getMPSCNNSoftMaxClass()
	rv := objc.Send[MPSCNNSoftMax](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNSoftMaxWithCoder(aDecoder foundation.INSCoder) MPSCNNSoftMax {
	instance := getMPSCNNSoftMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNSoftMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(coder:device:)
func NewCNNSoftMaxWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNSoftMax {
	instance := getMPSCNNSoftMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNSoftMaxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCNNSoftMaxWithDevice(device metal.MTLDevice) MPSCNNSoftMax {
	instance := getMPSCNNSoftMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNSoftMaxFromID(rv)
}

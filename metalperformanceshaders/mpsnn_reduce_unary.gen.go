// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNReduceUnary] class.
var (
	_MPSNNReduceUnaryClass     MPSNNReduceUnaryClass
	_MPSNNReduceUnaryClassOnce sync.Once
)

func getMPSNNReduceUnaryClass() MPSNNReduceUnaryClass {
	_MPSNNReduceUnaryClassOnce.Do(func() {
		_MPSNNReduceUnaryClass = MPSNNReduceUnaryClass{class: objc.GetClass("MPSNNReduceUnary")}
	})
	return _MPSNNReduceUnaryClass
}

// GetMPSNNReduceUnaryClass returns the class object for MPSNNReduceUnary.
func GetMPSNNReduceUnaryClass() MPSNNReduceUnaryClass {
	return getMPSNNReduceUnaryClass()
}

type MPSNNReduceUnaryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNReduceUnaryClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceUnaryClass) Alloc() MPSNNReduceUnary {
	rv := objc.Send[MPSNNReduceUnary](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The base class for unary reduction filters.
//
// # Instance Properties
//
//   - [MPSNNReduceUnary.ClipRectSource]
//   - [MPSNNReduceUnary.SetClipRectSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceUnary
type MPSNNReduceUnary struct {
	MPSCNNKernel
}

// MPSNNReduceUnaryFromID constructs a [MPSNNReduceUnary] from an objc.ID.
//
// The base class for unary reduction filters.
func MPSNNReduceUnaryFromID(id objc.ID) MPSNNReduceUnary {
	return MPSNNReduceUnary{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSNNReduceUnary adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNReduceUnary] class.
//
// # Instance Properties
//
//   - [IMPSNNReduceUnary.ClipRectSource]
//   - [IMPSNNReduceUnary.SetClipRectSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceUnary
type IMPSNNReduceUnary interface {
	IMPSCNNKernel

	// Topic: Instance Properties

	ClipRectSource() metal.MTLRegion
	SetClipRectSource(value metal.MTLRegion)
}

// Init initializes the instance.
func (r MPSNNReduceUnary) Init() MPSNNReduceUnary {
	rv := objc.Send[MPSNNReduceUnary](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceUnary) Autorelease() MPSNNReduceUnary {
	rv := objc.Send[MPSNNReduceUnary](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceUnary creates a new MPSNNReduceUnary instance.
func NewMPSNNReduceUnary() MPSNNReduceUnary {
	class := getMPSNNReduceUnaryClass()
	rv := objc.Send[MPSNNReduceUnary](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewReduceUnaryWithCoder(aDecoder foundation.INSCoder) MPSNNReduceUnary {
	instance := getMPSNNReduceUnaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNReduceUnaryFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(coder:device:)
func NewReduceUnaryWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNReduceUnary {
	instance := getMPSNNReduceUnaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNReduceUnaryFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewReduceUnaryWithDevice(device metal.MTLDevice) MPSNNReduceUnary {
	instance := getMPSNNReduceUnaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceUnaryFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceUnary/clipRectSource
func (r MPSNNReduceUnary) ClipRectSource() metal.MTLRegion {
	rv := objc.Send[metal.MTLRegion](r.ID, objc.Sel("clipRectSource"))
	return metal.MTLRegion(rv)
}
func (r MPSNNReduceUnary) SetClipRectSource(value metal.MTLRegion) {
	objc.Send[struct{}](r.ID, objc.Sel("setClipRectSource:"), value)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageReduceUnary] class.
var (
	_MPSImageReduceUnaryClass     MPSImageReduceUnaryClass
	_MPSImageReduceUnaryClassOnce sync.Once
)

func getMPSImageReduceUnaryClass() MPSImageReduceUnaryClass {
	_MPSImageReduceUnaryClassOnce.Do(func() {
		_MPSImageReduceUnaryClass = MPSImageReduceUnaryClass{class: objc.GetClass("MPSImageReduceUnary")}
	})
	return _MPSImageReduceUnaryClass
}

// GetMPSImageReduceUnaryClass returns the class object for MPSImageReduceUnary.
func GetMPSImageReduceUnaryClass() MPSImageReduceUnaryClass {
	return getMPSImageReduceUnaryClass()
}

type MPSImageReduceUnaryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageReduceUnaryClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageReduceUnaryClass) Alloc() MPSImageReduceUnary {
	rv := objc.Send[MPSImageReduceUnary](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The base class for reduction filters that take a single source as input.
//
// # Instance Properties
//
//   - [MPSImageReduceUnary.ClipRectSource]
//   - [MPSImageReduceUnary.SetClipRectSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceUnary
type MPSImageReduceUnary struct {
	MPSUnaryImageKernel
}

// MPSImageReduceUnaryFromID constructs a [MPSImageReduceUnary] from an objc.ID.
//
// The base class for reduction filters that take a single source as input.
func MPSImageReduceUnaryFromID(id objc.ID) MPSImageReduceUnary {
	return MPSImageReduceUnary{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageReduceUnary adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageReduceUnary] class.
//
// # Instance Properties
//
//   - [IMPSImageReduceUnary.ClipRectSource]
//   - [IMPSImageReduceUnary.SetClipRectSource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceUnary
type IMPSImageReduceUnary interface {
	IMPSUnaryImageKernel

	// Topic: Instance Properties

	ClipRectSource() metal.MTLRegion
	SetClipRectSource(value metal.MTLRegion)
}

// Init initializes the instance.
func (i MPSImageReduceUnary) Init() MPSImageReduceUnary {
	rv := objc.Send[MPSImageReduceUnary](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageReduceUnary) Autorelease() MPSImageReduceUnary {
	rv := objc.Send[MPSImageReduceUnary](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageReduceUnary creates a new MPSImageReduceUnary instance.
func NewMPSImageReduceUnary() MPSImageReduceUnary {
	class := getMPSImageReduceUnaryClass()
	rv := objc.Send[MPSImageReduceUnary](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageReduceUnaryWithCoder(aDecoder foundation.INSCoder) MPSImageReduceUnary {
	instance := getMPSImageReduceUnaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageReduceUnaryFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(coder:device:)
func NewImageReduceUnaryWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageReduceUnary {
	instance := getMPSImageReduceUnaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageReduceUnaryFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageReduceUnaryWithDevice(device metal.MTLDevice) MPSImageReduceUnary {
	instance := getMPSImageReduceUnaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageReduceUnaryFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceUnary/clipRectSource
func (i MPSImageReduceUnary) ClipRectSource() metal.MTLRegion {
	rv := objc.Send[metal.MTLRegion](i.ID, objc.Sel("clipRectSource"))
	return metal.MTLRegion(rv)
}
func (i MPSImageReduceUnary) SetClipRectSource(value metal.MTLRegion) {
	objc.Send[struct{}](i.ID, objc.Sel("setClipRectSource:"), value)
}

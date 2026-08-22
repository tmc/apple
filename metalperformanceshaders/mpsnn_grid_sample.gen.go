// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNGridSample] class.
var (
	_MPSNNGridSampleClass     MPSNNGridSampleClass
	_MPSNNGridSampleClassOnce sync.Once
)

func getMPSNNGridSampleClass() MPSNNGridSampleClass {
	_MPSNNGridSampleClassOnce.Do(func() {
		_MPSNNGridSampleClass = MPSNNGridSampleClass{class: objc.GetClass("MPSNNGridSample")}
	})
	return _MPSNNGridSampleClass
}

// GetMPSNNGridSampleClass returns the class object for MPSNNGridSample.
func GetMPSNNGridSampleClass() MPSNNGridSampleClass {
	return getMPSNNGridSampleClass()
}

type MPSNNGridSampleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNGridSampleClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNGridSampleClass) Alloc() MPSNNGridSample {
	rv := objc.Send[MPSNNGridSample](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MPSNNGridSample.UseGridValueAsInputCoordinate]
//   - [MPSNNGridSample.SetUseGridValueAsInputCoordinate]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGridSample
type MPSNNGridSample struct {
	MPSCNNBinaryKernel
}

// MPSNNGridSampleFromID constructs a [MPSNNGridSample] from an objc.ID.
func MPSNNGridSampleFromID(id objc.ID) MPSNNGridSample {
	return MPSNNGridSample{MPSCNNBinaryKernel: MPSCNNBinaryKernelFromID(id)}
}

// NOTE: MPSNNGridSample adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNGridSample] class.
//
// # Instance Properties
//
//   - [IMPSNNGridSample.UseGridValueAsInputCoordinate]
//   - [IMPSNNGridSample.SetUseGridValueAsInputCoordinate]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGridSample
type IMPSNNGridSample interface {
	IMPSCNNBinaryKernel

	// Topic: Instance Properties

	UseGridValueAsInputCoordinate() bool
	SetUseGridValueAsInputCoordinate(value bool)
}

// Init initializes the instance.
func (g MPSNNGridSample) Init() MPSNNGridSample {
	rv := objc.Send[MPSNNGridSample](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSNNGridSample) Autorelease() MPSNNGridSample {
	rv := objc.Send[MPSNNGridSample](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNGridSample creates a new MPSNNGridSample instance.
func NewMPSNNGridSample() MPSNNGridSample {
	class := getMPSNNGridSampleClass()
	rv := objc.Send[MPSNNGridSample](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewGridSampleWithCoder(aDecoder foundation.INSCoder) MPSNNGridSample {
	instance := getMPSNNGridSampleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNGridSampleFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGridSample/init(coder:device:)
func NewGridSampleWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNGridSample {
	instance := getMPSNNGridSampleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNGridSampleFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGridSample/init(device:)
func NewGridSampleWithDevice(device metal.MTLDevice) MPSNNGridSample {
	instance := getMPSNNGridSampleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNGridSampleFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGridSample/useGridValueAsInputCoordinate
func (g MPSNNGridSample) UseGridValueAsInputCoordinate() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("useGridValueAsInputCoordinate"))
	return rv
}
func (g MPSNNGridSample) SetUseGridValueAsInputCoordinate(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setUseGridValueAsInputCoordinate:"), value)
}

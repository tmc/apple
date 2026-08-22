// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNDArrayGather] class.
var (
	_MPSNDArrayGatherClass     MPSNDArrayGatherClass
	_MPSNDArrayGatherClassOnce sync.Once
)

func getMPSNDArrayGatherClass() MPSNDArrayGatherClass {
	_MPSNDArrayGatherClassOnce.Do(func() {
		_MPSNDArrayGatherClass = MPSNDArrayGatherClass{class: objc.GetClass("MPSNDArrayGather")}
	})
	return _MPSNDArrayGatherClass
}

// GetMPSNDArrayGatherClass returns the class object for MPSNDArrayGather.
func GetMPSNDArrayGatherClass() MPSNDArrayGatherClass {
	return getMPSNDArrayGatherClass()
}

type MPSNDArrayGatherClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayGatherClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayGatherClass) Alloc() MPSNDArrayGather {
	rv := objc.Send[MPSNDArrayGather](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MPSNDArrayGather.Axis]
//   - [MPSNDArrayGather.SetAxis]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayGather
type MPSNDArrayGather struct {
	MPSNDArrayBinaryKernel
}

// MPSNDArrayGatherFromID constructs a [MPSNDArrayGather] from an objc.ID.
func MPSNDArrayGatherFromID(id objc.ID) MPSNDArrayGather {
	return MPSNDArrayGather{MPSNDArrayBinaryKernel: MPSNDArrayBinaryKernelFromID(id)}
}

// NOTE: MPSNDArrayGather adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayGather] class.
//
// # Instance Properties
//
//   - [IMPSNDArrayGather.Axis]
//   - [IMPSNDArrayGather.SetAxis]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayGather
type IMPSNDArrayGather interface {
	IMPSNDArrayBinaryKernel

	// Topic: Instance Properties

	Axis() uint
	SetAxis(value uint)
}

// Init initializes the instance.
func (n MPSNDArrayGather) Init() MPSNDArrayGather {
	rv := objc.Send[MPSNDArrayGather](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayGather) Autorelease() MPSNDArrayGather {
	rv := objc.Send[MPSNDArrayGather](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayGather creates a new MPSNDArrayGather instance.
func NewMPSNDArrayGather() MPSNDArrayGather {
	class := getMPSNDArrayGatherClass()
	rv := objc.Send[MPSNDArrayGather](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewNDArrayGatherWithCoder(aDecoder foundation.INSCoder) MPSNDArrayGather {
	instance := getMPSNDArrayGatherClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNDArrayGatherFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryKernel/init(coder:device:)
func NewNDArrayGatherWithCoderDevice(coder foundation.INSCoder, device metal.MTLDevice) MPSNDArrayGather {
	instance := getMPSNDArrayGatherClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNDArrayGatherFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryKernel/init(device:)
func NewNDArrayGatherWithDevice(device metal.MTLDevice) MPSNDArrayGather {
	instance := getMPSNDArrayGatherClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNDArrayGatherFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/init(device:sourceCount:)
func NewNDArrayGatherWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayGather {
	instance := getMPSNDArrayGatherClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	return MPSNDArrayGatherFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayGather/axis
func (n MPSNDArrayGather) Axis() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("axis"))
	return rv
}
func (n MPSNDArrayGather) SetAxis(value uint) {
	objc.Send[struct{}](n.ID, objc.Sel("setAxis:"), value)
}

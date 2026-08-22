// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNDArrayStridedSlice] class.
var (
	_MPSNDArrayStridedSliceClass     MPSNDArrayStridedSliceClass
	_MPSNDArrayStridedSliceClassOnce sync.Once
)

func getMPSNDArrayStridedSliceClass() MPSNDArrayStridedSliceClass {
	_MPSNDArrayStridedSliceClassOnce.Do(func() {
		_MPSNDArrayStridedSliceClass = MPSNDArrayStridedSliceClass{class: objc.GetClass("MPSNDArrayStridedSlice")}
	})
	return _MPSNDArrayStridedSliceClass
}

// GetMPSNDArrayStridedSliceClass returns the class object for MPSNDArrayStridedSlice.
func GetMPSNDArrayStridedSliceClass() MPSNDArrayStridedSliceClass {
	return getMPSNDArrayStridedSliceClass()
}

type MPSNDArrayStridedSliceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayStridedSliceClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayStridedSliceClass) Alloc() MPSNDArrayStridedSlice {
	rv := objc.Send[MPSNDArrayStridedSlice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MPSNDArrayStridedSlice.Strides]
//   - [MPSNDArrayStridedSlice.SetStrides]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayStridedSlice
type MPSNDArrayStridedSlice struct {
	MPSNDArrayUnaryKernel
}

// MPSNDArrayStridedSliceFromID constructs a [MPSNDArrayStridedSlice] from an objc.ID.
func MPSNDArrayStridedSliceFromID(id objc.ID) MPSNDArrayStridedSlice {
	return MPSNDArrayStridedSlice{MPSNDArrayUnaryKernel: MPSNDArrayUnaryKernelFromID(id)}
}

// NOTE: MPSNDArrayStridedSlice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayStridedSlice] class.
//
// # Instance Properties
//
//   - [IMPSNDArrayStridedSlice.Strides]
//   - [IMPSNDArrayStridedSlice.SetStrides]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayStridedSlice
type IMPSNDArrayStridedSlice interface {
	IMPSNDArrayUnaryKernel

	// Topic: Instance Properties

	Strides() MPSNDArrayOffsets
	SetStrides(value MPSNDArrayOffsets)
}

// Init initializes the instance.
func (n MPSNDArrayStridedSlice) Init() MPSNDArrayStridedSlice {
	rv := objc.Send[MPSNDArrayStridedSlice](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayStridedSlice) Autorelease() MPSNDArrayStridedSlice {
	rv := objc.Send[MPSNDArrayStridedSlice](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayStridedSlice creates a new MPSNDArrayStridedSlice instance.
func NewMPSNDArrayStridedSlice() MPSNDArrayStridedSlice {
	class := getMPSNDArrayStridedSliceClass()
	rv := objc.Send[MPSNDArrayStridedSlice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewNDArrayStridedSliceWithCoder(aDecoder foundation.INSCoder) MPSNDArrayStridedSlice {
	instance := getMPSNDArrayStridedSliceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNDArrayStridedSliceFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryKernel/init(coder:device:)
func NewNDArrayStridedSliceWithCoderDevice(coder foundation.INSCoder, device metal.MTLDevice) MPSNDArrayStridedSlice {
	instance := getMPSNDArrayStridedSliceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNDArrayStridedSliceFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryKernel/init(device:)
func NewNDArrayStridedSliceWithDevice(device metal.MTLDevice) MPSNDArrayStridedSlice {
	instance := getMPSNDArrayStridedSliceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNDArrayStridedSliceFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel/init(device:sourceCount:)
func NewNDArrayStridedSliceWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayStridedSlice {
	instance := getMPSNDArrayStridedSliceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	return MPSNDArrayStridedSliceFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayStridedSlice/strides
func (n MPSNDArrayStridedSlice) Strides() MPSNDArrayOffsets {
	rv := objc.Send[MPSNDArrayOffsets](n.ID, objc.Sel("strides"))
	return MPSNDArrayOffsets(rv)
}
func (n MPSNDArrayStridedSlice) SetStrides(value MPSNDArrayOffsets) {
	objc.Send[struct{}](n.ID, objc.Sel("setStrides:"), value)
}

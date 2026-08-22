// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNSlice] class.
var (
	_MPSNNSliceClass     MPSNNSliceClass
	_MPSNNSliceClassOnce sync.Once
)

func getMPSNNSliceClass() MPSNNSliceClass {
	_MPSNNSliceClassOnce.Do(func() {
		_MPSNNSliceClass = MPSNNSliceClass{class: objc.GetClass("MPSNNSlice")}
	})
	return _MPSNNSliceClass
}

// GetMPSNNSliceClass returns the class object for MPSNNSlice.
func GetMPSNNSliceClass() MPSNNSliceClass {
	return getMPSNNSliceClass()
}

type MPSNNSliceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNSliceClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNSliceClass) Alloc() MPSNNSlice {
	rv := objc.Send[MPSNNSlice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A kernel that extracts a slice from an image.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNSlice
type MPSNNSlice struct {
	MPSCNNKernel
}

// MPSNNSliceFromID constructs a [MPSNNSlice] from an objc.ID.
//
// A kernel that extracts a slice from an image.
func MPSNNSliceFromID(id objc.ID) MPSNNSlice {
	return MPSNNSlice{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSNNSlice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNSlice] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNSlice
type IMPSNNSlice interface {
	IMPSCNNKernel
}

// Init initializes the instance.
func (s MPSNNSlice) Init() MPSNNSlice {
	rv := objc.Send[MPSNNSlice](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MPSNNSlice) Autorelease() MPSNNSlice {
	rv := objc.Send[MPSNNSlice](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNSlice creates a new MPSNNSlice instance.
func NewMPSNNSlice() MPSNNSlice {
	class := getMPSNNSliceClass()
	rv := objc.Send[MPSNNSlice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewSliceWithCoder(aDecoder foundation.INSCoder) MPSNNSlice {
	instance := getMPSNNSliceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNSliceFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNSlice/init(coder:device:)
func NewSliceWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNSlice {
	instance := getMPSNNSliceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNSliceFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNSlice/init(device:)
func NewSliceWithDevice(device metal.MTLDevice) MPSNNSlice {
	instance := getMPSNNSliceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNSliceFromID(rv)
}

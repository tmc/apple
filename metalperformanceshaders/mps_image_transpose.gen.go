// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageTranspose] class.
var (
	_MPSImageTransposeClass     MPSImageTransposeClass
	_MPSImageTransposeClassOnce sync.Once
)

func getMPSImageTransposeClass() MPSImageTransposeClass {
	_MPSImageTransposeClassOnce.Do(func() {
		_MPSImageTransposeClass = MPSImageTransposeClass{class: objc.GetClass("MPSImageTranspose")}
	})
	return _MPSImageTransposeClass
}

// GetMPSImageTransposeClass returns the class object for MPSImageTranspose.
func GetMPSImageTransposeClass() MPSImageTransposeClass {
	return getMPSImageTransposeClass()
}

type MPSImageTransposeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageTransposeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageTransposeClass) Alloc() MPSImageTranspose {
	rv := objc.Send[MPSImageTranspose](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that transposes an image.
//
// # Overview
//
// An [MPSImageTranspose] filter applies a matrix transposition to the source
// image by exchanging its rows with its columns.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageTranspose
type MPSImageTranspose struct {
	MPSUnaryImageKernel
}

// MPSImageTransposeFromID constructs a [MPSImageTranspose] from an objc.ID.
//
// A filter that transposes an image.
func MPSImageTransposeFromID(id objc.ID) MPSImageTranspose {
	return MPSImageTranspose{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageTranspose adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageTranspose] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageTranspose
type IMPSImageTranspose interface {
	IMPSUnaryImageKernel
}

// Init initializes the instance.
func (i MPSImageTranspose) Init() MPSImageTranspose {
	rv := objc.Send[MPSImageTranspose](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageTranspose) Autorelease() MPSImageTranspose {
	rv := objc.Send[MPSImageTranspose](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageTranspose creates a new MPSImageTranspose instance.
func NewMPSImageTranspose() MPSImageTranspose {
	class := getMPSImageTransposeClass()
	rv := objc.Send[MPSImageTranspose](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageTransposeWithCoder(aDecoder foundation.INSCoder) MPSImageTranspose {
	instance := getMPSImageTransposeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageTransposeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(coder:device:)
func NewImageTransposeWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageTranspose {
	instance := getMPSImageTransposeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageTransposeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageTransposeWithDevice(device metal.MTLDevice) MPSImageTranspose {
	instance := getMPSImageTransposeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageTransposeFromID(rv)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageTent] class.
var (
	_MPSImageTentClass     MPSImageTentClass
	_MPSImageTentClassOnce sync.Once
)

func getMPSImageTentClass() MPSImageTentClass {
	_MPSImageTentClassOnce.Do(func() {
		_MPSImageTentClass = MPSImageTentClass{class: objc.GetClass("MPSImageTent")}
	})
	return _MPSImageTentClass
}

// GetMPSImageTentClass returns the class object for MPSImageTent.
func GetMPSImageTentClass() MPSImageTentClass {
	return getMPSImageTentClass()
}

type MPSImageTentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageTentClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageTentClass) Alloc() MPSImageTent {
	rv := objc.Send[MPSImageTent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that convolves an image with a tent filter.
//
// # Overview
//
// The kernel elements of the filter form a tent shape with increasing sides,
// for example:
//
// [media-2556918]
//
// Like a box filter, this arrangement allows for much faster algorithms,
// especially for larger blur radii but with a more pleasing appearance.
//
// The tent blur is a separable filter and the Metal Performance Shaders
// framework will act accordingly to give the best performance for
// multi-dimensional blurs.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageTent
type MPSImageTent struct {
	MPSImageBox
}

// MPSImageTentFromID constructs a [MPSImageTent] from an objc.ID.
//
// A filter that convolves an image with a tent filter.
func MPSImageTentFromID(id objc.ID) MPSImageTent {
	return MPSImageTent{MPSImageBox: MPSImageBoxFromID(id)}
}

// NOTE: MPSImageTent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageTent] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageTent
type IMPSImageTent interface {
	IMPSImageBox
}

// Init initializes the instance.
func (i MPSImageTent) Init() MPSImageTent {
	rv := objc.Send[MPSImageTent](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageTent) Autorelease() MPSImageTent {
	rv := objc.Send[MPSImageTent](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageTent creates a new MPSImageTent instance.
func NewMPSImageTent() MPSImageTent {
	class := getMPSImageTentClass()
	rv := objc.Send[MPSImageTent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageTentWithCoder(aDecoder foundation.INSCoder) MPSImageTent {
	instance := getMPSImageTentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageTentFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageBox/init(coder:device:)
func NewImageTentWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageTent {
	instance := getMPSImageTentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageTentFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageTentWithDevice(device metal.MTLDevice) MPSImageTent {
	instance := getMPSImageTentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageTentFromID(rv)
}

// Initializes a box filter.
//
// device: The Metal device the filter will run on.
//
// kernelWidth: The width of the kernel. Must be an odd number.
//
// kernelHeight: The height of the kernel. Must be an odd number.
//
// # Return Value
//
// An initialized box filter object.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageBox/init(device:kernelWidth:kernelHeight:)
func NewImageTentWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSImageTent {
	instance := getMPSImageTentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return MPSImageTentFromID(rv)
}

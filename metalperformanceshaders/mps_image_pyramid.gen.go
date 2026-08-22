// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImagePyramid] class.
var (
	_MPSImagePyramidClass     MPSImagePyramidClass
	_MPSImagePyramidClassOnce sync.Once
)

func getMPSImagePyramidClass() MPSImagePyramidClass {
	_MPSImagePyramidClassOnce.Do(func() {
		_MPSImagePyramidClass = MPSImagePyramidClass{class: objc.GetClass("MPSImagePyramid")}
	})
	return _MPSImagePyramidClass
}

// GetMPSImagePyramidClass returns the class object for MPSImagePyramid.
func GetMPSImagePyramidClass() MPSImagePyramidClass {
	return getMPSImagePyramidClass()
}

type MPSImagePyramidClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImagePyramidClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImagePyramidClass) Alloc() MPSImagePyramid {
	rv := objc.Send[MPSImagePyramid](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A base class for creating different kinds of pyramid images.
//
// # Methods
//
//   - [MPSImagePyramid.InitWithDeviceCenterWeight]: Initialize a downwards 5-tap image pyramid with a central weight parameter and device.
//   - [MPSImagePyramid.InitWithDeviceKernelWidthKernelHeightWeights]: Initialize a downwards n-tap image pyramid with a custom filter kernel and device.
//
// # Properties
//
//   - [MPSImagePyramid.KernelWidth]: The width of the filter window. Must be an odd number.
//   - [MPSImagePyramid.KernelHeight]: The height of the filter window. Must be an odd number.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid
type MPSImagePyramid struct {
	MPSUnaryImageKernel
}

// MPSImagePyramidFromID constructs a [MPSImagePyramid] from an objc.ID.
//
// A base class for creating different kinds of pyramid images.
func MPSImagePyramidFromID(id objc.ID) MPSImagePyramid {
	return MPSImagePyramid{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImagePyramid adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImagePyramid] class.
//
// # Methods
//
//   - [IMPSImagePyramid.InitWithDeviceCenterWeight]: Initialize a downwards 5-tap image pyramid with a central weight parameter and device.
//   - [IMPSImagePyramid.InitWithDeviceKernelWidthKernelHeightWeights]: Initialize a downwards n-tap image pyramid with a custom filter kernel and device.
//
// # Properties
//
//   - [IMPSImagePyramid.KernelWidth]: The width of the filter window. Must be an odd number.
//   - [IMPSImagePyramid.KernelHeight]: The height of the filter window. Must be an odd number.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid
type IMPSImagePyramid interface {
	IMPSUnaryImageKernel

	// Topic: Methods

	// Initialize a downwards 5-tap image pyramid with a central weight parameter and device.
	InitWithDeviceCenterWeight(device metal.MTLDevice, centerWeight float32) MPSImagePyramid
	// Initialize a downwards n-tap image pyramid with a custom filter kernel and device.
	InitWithDeviceKernelWidthKernelHeightWeights(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, kernelWeights *float32) MPSImagePyramid

	// Topic: Properties

	// The width of the filter window. Must be an odd number.
	KernelWidth() uint
	// The height of the filter window. Must be an odd number.
	KernelHeight() uint
}

// Init initializes the instance.
func (i MPSImagePyramid) Init() MPSImagePyramid {
	rv := objc.Send[MPSImagePyramid](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImagePyramid) Autorelease() MPSImagePyramid {
	rv := objc.Send[MPSImagePyramid](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImagePyramid creates a new MPSImagePyramid instance.
func NewMPSImagePyramid() MPSImagePyramid {
	class := getMPSImagePyramidClass()
	rv := objc.Send[MPSImagePyramid](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImagePyramidWithCoder(aDecoder foundation.INSCoder) MPSImagePyramid {
	instance := getMPSImagePyramidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImagePyramidFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid/init(coder:device:)
func NewImagePyramidWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImagePyramid {
	instance := getMPSImagePyramidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImagePyramidFromID(rv)
}

// Initializes a downwards 5-tap image pyramid with the default filter kernel
// and device.
//
// device: The device the filter will run on.
//
// # Return Value
//
// A valid [MPSImagePyramid] object or `nil`, if failure.
//
// # Discussion
//
// The filter kernel is the outer product of `w = [1/16, 1/4, 3/8, 1/4,
// 1/16]^T`, with itself.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid/init(device:)
func NewImagePyramidWithDevice(device metal.MTLDevice) MPSImagePyramid {
	instance := getMPSImagePyramidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImagePyramidFromID(rv)
}

// Initialize a downwards 5-tap image pyramid with a central weight parameter
// and device.
//
// device: The device the filter will run on.
//
// centerWeight: Defines the form of the filter kernel through the outer product `ww^T`,
// where `w = [(1/4 - a/2), 1/4, a, 1/4, (1/4 - a/2)]^T` and `a` is the value
// of `centerWeight`.
//
// # Return Value
//
// A valid [MPSImagePyramid] object or `nil`, if failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid/init(device:centerWeight:)
func NewImagePyramidWithDeviceCenterWeight(device metal.MTLDevice, centerWeight float32) MPSImagePyramid {
	instance := getMPSImagePyramidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:centerWeight:"), device, centerWeight)
	return MPSImagePyramidFromID(rv)
}

// Initialize a downwards n-tap image pyramid with a custom filter kernel and
// device.
//
// device: The device the filter will run on.
//
// kernelWidth: The width of the filter kernel.
//
// kernelHeight: The height of the filter kernel.
//
// kernelWeights: A pointer to an array of `kernelWidth*kernelHeight` values to be used as
// the kernel. These values are in row-major order.
//
// # Return Value
//
// A valid [MPSImagePyramid] object or `nil`, if failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid/init(device:kernelWidth:kernelHeight:weights:)
func NewImagePyramidWithDeviceKernelWidthKernelHeightWeights(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, kernelWeights *float32) MPSImagePyramid {
	instance := getMPSImagePyramidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:weights:"), device, kernelWidth, kernelHeight, kernelWeights)
	return MPSImagePyramidFromID(rv)
}

// Initialize a downwards 5-tap image pyramid with a central weight parameter
// and device.
//
// device: The device the filter will run on.
//
// centerWeight: Defines the form of the filter kernel through the outer product `ww^T`,
// where `w = [(1/4 - a/2), 1/4, a, 1/4, (1/4 - a/2)]^T` and `a` is the value
// of `centerWeight`.
//
// # Return Value
//
// A valid [MPSImagePyramid] object or `nil`, if failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid/init(device:centerWeight:)
func (i MPSImagePyramid) InitWithDeviceCenterWeight(device metal.MTLDevice, centerWeight float32) MPSImagePyramid {
	rv := objc.Send[MPSImagePyramid](i.ID, objc.Sel("initWithDevice:centerWeight:"), device, centerWeight)
	return rv
}

// Initialize a downwards n-tap image pyramid with a custom filter kernel and
// device.
//
// device: The device the filter will run on.
//
// kernelWidth: The width of the filter kernel.
//
// kernelHeight: The height of the filter kernel.
//
// kernelWeights: A pointer to an array of `kernelWidth*kernelHeight` values to be used as
// the kernel. These values are in row-major order.
//
// # Return Value
//
// A valid [MPSImagePyramid] object or `nil`, if failure.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid/init(device:kernelWidth:kernelHeight:weights:)
func (i MPSImagePyramid) InitWithDeviceKernelWidthKernelHeightWeights(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, kernelWeights *float32) MPSImagePyramid {
	rv := objc.Send[MPSImagePyramid](i.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:weights:"), device, kernelWidth, kernelHeight, kernelWeights)
	return rv
}

// The width of the filter window. Must be an odd number.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid/kernelWidth
func (i MPSImagePyramid) KernelWidth() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("kernelWidth"))
	return rv
}

// The height of the filter window. Must be an odd number.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid/kernelHeight
func (i MPSImagePyramid) KernelHeight() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("kernelHeight"))
	return rv
}

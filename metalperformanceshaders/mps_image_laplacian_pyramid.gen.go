// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageLaplacianPyramid] class.
var (
	_MPSImageLaplacianPyramidClass     MPSImageLaplacianPyramidClass
	_MPSImageLaplacianPyramidClassOnce sync.Once
)

func getMPSImageLaplacianPyramidClass() MPSImageLaplacianPyramidClass {
	_MPSImageLaplacianPyramidClassOnce.Do(func() {
		_MPSImageLaplacianPyramidClass = MPSImageLaplacianPyramidClass{class: objc.GetClass("MPSImageLaplacianPyramid")}
	})
	return _MPSImageLaplacianPyramidClass
}

// GetMPSImageLaplacianPyramidClass returns the class object for MPSImageLaplacianPyramid.
func GetMPSImageLaplacianPyramidClass() MPSImageLaplacianPyramidClass {
	return getMPSImageLaplacianPyramidClass()
}

type MPSImageLaplacianPyramidClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageLaplacianPyramidClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageLaplacianPyramidClass) Alloc() MPSImageLaplacianPyramid {
	rv := objc.Send[MPSImageLaplacianPyramid](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that convolves an image with a Laplacian filter.
//
// # Instance Properties
//
//   - [MPSImageLaplacianPyramid.LaplacianBias]
//   - [MPSImageLaplacianPyramid.SetLaplacianBias]
//   - [MPSImageLaplacianPyramid.LaplacianScale]
//   - [MPSImageLaplacianPyramid.SetLaplacianScale]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLaplacianPyramid
type MPSImageLaplacianPyramid struct {
	MPSImagePyramid
}

// MPSImageLaplacianPyramidFromID constructs a [MPSImageLaplacianPyramid] from an objc.ID.
//
// A filter that convolves an image with a Laplacian filter.
func MPSImageLaplacianPyramidFromID(id objc.ID) MPSImageLaplacianPyramid {
	return MPSImageLaplacianPyramid{MPSImagePyramid: MPSImagePyramidFromID(id)}
}

// NOTE: MPSImageLaplacianPyramid adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageLaplacianPyramid] class.
//
// # Instance Properties
//
//   - [IMPSImageLaplacianPyramid.LaplacianBias]
//   - [IMPSImageLaplacianPyramid.SetLaplacianBias]
//   - [IMPSImageLaplacianPyramid.LaplacianScale]
//   - [IMPSImageLaplacianPyramid.SetLaplacianScale]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLaplacianPyramid
type IMPSImageLaplacianPyramid interface {
	IMPSImagePyramid

	// Topic: Instance Properties

	LaplacianBias() float32
	SetLaplacianBias(value float32)
	LaplacianScale() float32
	SetLaplacianScale(value float32)
}

// Init initializes the instance.
func (i MPSImageLaplacianPyramid) Init() MPSImageLaplacianPyramid {
	rv := objc.Send[MPSImageLaplacianPyramid](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageLaplacianPyramid) Autorelease() MPSImageLaplacianPyramid {
	rv := objc.Send[MPSImageLaplacianPyramid](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageLaplacianPyramid creates a new MPSImageLaplacianPyramid instance.
func NewMPSImageLaplacianPyramid() MPSImageLaplacianPyramid {
	class := getMPSImageLaplacianPyramidClass()
	rv := objc.Send[MPSImageLaplacianPyramid](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageLaplacianPyramidWithCoder(aDecoder foundation.INSCoder) MPSImageLaplacianPyramid {
	instance := getMPSImageLaplacianPyramidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageLaplacianPyramidFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid/init(coder:device:)
func NewImageLaplacianPyramidWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageLaplacianPyramid {
	instance := getMPSImageLaplacianPyramidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageLaplacianPyramidFromID(rv)
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
func NewImageLaplacianPyramidWithDevice(device metal.MTLDevice) MPSImageLaplacianPyramid {
	instance := getMPSImageLaplacianPyramidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageLaplacianPyramidFromID(rv)
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
func NewImageLaplacianPyramidWithDeviceCenterWeight(device metal.MTLDevice, centerWeight float32) MPSImageLaplacianPyramid {
	instance := getMPSImageLaplacianPyramidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:centerWeight:"), device, centerWeight)
	return MPSImageLaplacianPyramidFromID(rv)
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
func NewImageLaplacianPyramidWithDeviceKernelWidthKernelHeightWeights(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, kernelWeights *float32) MPSImageLaplacianPyramid {
	instance := getMPSImageLaplacianPyramidClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:weights:"), device, kernelWidth, kernelHeight, kernelWeights)
	return MPSImageLaplacianPyramidFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLaplacianPyramid/laplacianBias
func (i MPSImageLaplacianPyramid) LaplacianBias() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("getLaplacianBias"))
	return rv
}
func (i MPSImageLaplacianPyramid) SetLaplacianBias(value float32) {
	objc.Send[struct{}](i.ID, objc.Sel("setLaplacianBias:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLaplacianPyramid/laplacianScale
func (i MPSImageLaplacianPyramid) LaplacianScale() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("getLaplacianScale"))
	return rv
}
func (i MPSImageLaplacianPyramid) SetLaplacianScale(value float32) {
	objc.Send[struct{}](i.ID, objc.Sel("setLaplacianScale:"), value)
}

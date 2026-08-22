// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageErode] class.
var (
	_MPSImageErodeClass     MPSImageErodeClass
	_MPSImageErodeClassOnce sync.Once
)

func getMPSImageErodeClass() MPSImageErodeClass {
	_MPSImageErodeClassOnce.Do(func() {
		_MPSImageErodeClass = MPSImageErodeClass{class: objc.GetClass("MPSImageErode")}
	})
	return _MPSImageErodeClass
}

// GetMPSImageErodeClass returns the class object for MPSImageErode.
func GetMPSImageErodeClass() MPSImageErodeClass {
	return getMPSImageErodeClass()
}

type MPSImageErodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageErodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageErodeClass) Alloc() MPSImageErode {
	rv := objc.Send[MPSImageErode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that finds the minimum pixel value in a rectangular region by
// applying an erosion function.
//
// # Overview
//
// An [MPSImageErode] behaves like the [MPSImageAreaMin] filter, except that
// Metal calculates the intensity at each position relative to a different
// value before determining which is the maximum pixel value, allowing for
// shaped, nonrectangular morphological probes.
//
// The code example below shows pseudocode for the calculation that returns
// each pixel value:
//
// The definition of the [MPSImageErode] filter is different from its `vImage`
// counterpart (`MPSImageErode_filter_value =
// 1.0f-vImageErode_filter_value.`). This allows [MPSImageDilate] and
// [MPSImageErode] to use the same filter, making open and close operators
// easier to write.
//
// A filter that contains all zeros is identical to an [MPSImageAreaMin]
// filter. Metal handles the center filter element as `0` to avoid causing a
// general lightening of the image, and it handles the
// [MPSUnaryImageKernel.EdgeMode] property as [MPSImageEdgeModeClamp] for this
// filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageErode
type MPSImageErode struct {
	MPSImageDilate
}

// MPSImageErodeFromID constructs a [MPSImageErode] from an objc.ID.
//
// A filter that finds the minimum pixel value in a rectangular region by
// applying an erosion function.
func MPSImageErodeFromID(id objc.ID) MPSImageErode {
	return MPSImageErode{MPSImageDilate: MPSImageDilateFromID(id)}
}

// NOTE: MPSImageErode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageErode] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageErode
type IMPSImageErode interface {
	IMPSImageDilate
}

// Init initializes the instance.
func (i MPSImageErode) Init() MPSImageErode {
	rv := objc.Send[MPSImageErode](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageErode) Autorelease() MPSImageErode {
	rv := objc.Send[MPSImageErode](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageErode creates a new MPSImageErode instance.
func NewMPSImageErode() MPSImageErode {
	class := getMPSImageErodeClass()
	rv := objc.Send[MPSImageErode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageErodeWithCoder(aDecoder foundation.INSCoder) MPSImageErode {
	instance := getMPSImageErodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageErodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDilate/init(coder:device:)
func NewImageErodeWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageErode {
	instance := getMPSImageErodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageErodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageErodeWithDevice(device metal.MTLDevice) MPSImageErode {
	instance := getMPSImageErodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageErodeFromID(rv)
}

// Initializes the kernel with a specified width, height, and weight values.
//
// device: The Metal device the filter will run on.
//
// kernelWidth: The width of the kernel. Must be an odd number.
//
// kernelHeight: The height of the kernel. Must be an odd number.
//
// values: The set of values to use as the dilate probe. The values are copied into
// the filter. To avoid image lightening or darkening, the center value should
// be `0.0f`.
//
// # Return Value
//
// Returns an initialized kernel object with specific width, height, and
// weight values.
//
// # Discussion
//
// Each dilate shape probe defines a 3D surface of values. These are arranged
// in order left to right, then top to bottom in a 1D array.
// (`values[kernelWidth*y+x] = probe[y][x]`)
//
// Values should be generally be in the range `[0,1]` with the center pixel
// tending towards `0` and edges towards `1`. However, any numerical value is
// allowed. Calculations are subject to the usual floating-point rounding
// error.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDilate/init(device:kernelWidth:kernelHeight:values:)
func NewImageErodeWithDeviceKernelWidthKernelHeightValues(device metal.MTLDevice, kernelWidth uint, kernelHeight uint, values *float32) MPSImageErode {
	instance := getMPSImageErodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:values:"), device, kernelWidth, kernelHeight, values)
	return MPSImageErodeFromID(rv)
}

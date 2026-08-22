// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageAreaMin] class.
var (
	_MPSImageAreaMinClass     MPSImageAreaMinClass
	_MPSImageAreaMinClassOnce sync.Once
)

func getMPSImageAreaMinClass() MPSImageAreaMinClass {
	_MPSImageAreaMinClassOnce.Do(func() {
		_MPSImageAreaMinClass = MPSImageAreaMinClass{class: objc.GetClass("MPSImageAreaMin")}
	})
	return _MPSImageAreaMinClass
}

// GetMPSImageAreaMinClass returns the class object for MPSImageAreaMin.
func GetMPSImageAreaMinClass() MPSImageAreaMinClass {
	return getMPSImageAreaMinClass()
}

type MPSImageAreaMinClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageAreaMinClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageAreaMinClass) Alloc() MPSImageAreaMin {
	rv := objc.Send[MPSImageAreaMin](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that finds the minimum pixel value in a rectangular region
// centered around each pixel in the source image.
//
// # Overview
//
// An [MPSImageAreaMin] filter has the same methods and properties as the
// [MPSImageAreaMax] class.
//
// If there are multiple channels in the source image, each channel is
// processed independently. The [MPSUnaryImageKernel.EdgeMode] property value
// is assumed to always be [MPSImageEdgeModeClamp] for this filter.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAreaMin
type MPSImageAreaMin struct {
	MPSImageAreaMax
}

// MPSImageAreaMinFromID constructs a [MPSImageAreaMin] from an objc.ID.
//
// A filter that finds the minimum pixel value in a rectangular region
// centered around each pixel in the source image.
func MPSImageAreaMinFromID(id objc.ID) MPSImageAreaMin {
	return MPSImageAreaMin{MPSImageAreaMax: MPSImageAreaMaxFromID(id)}
}

// NOTE: MPSImageAreaMin adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageAreaMin] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAreaMin
type IMPSImageAreaMin interface {
	IMPSImageAreaMax
}

// Init initializes the instance.
func (i MPSImageAreaMin) Init() MPSImageAreaMin {
	rv := objc.Send[MPSImageAreaMin](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageAreaMin) Autorelease() MPSImageAreaMin {
	rv := objc.Send[MPSImageAreaMin](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageAreaMin creates a new MPSImageAreaMin instance.
func NewMPSImageAreaMin() MPSImageAreaMin {
	class := getMPSImageAreaMinClass()
	rv := objc.Send[MPSImageAreaMin](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageAreaMinWithCoder(aDecoder foundation.INSCoder) MPSImageAreaMin {
	instance := getMPSImageAreaMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageAreaMinFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAreaMax/init(coder:device:)
func NewImageAreaMinWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageAreaMin {
	instance := getMPSImageAreaMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageAreaMinFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageAreaMinWithDevice(device metal.MTLDevice) MPSImageAreaMin {
	instance := getMPSImageAreaMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageAreaMinFromID(rv)
}

// Initializes the kernel with a specified width and height.
//
// device: The Metal device the filter will run on.
//
// kernelWidth: The width of the kernel. Must be an odd number.
//
// kernelHeight: The height of the kernel. Must be an odd number.
//
// # Return Value
//
// Returns an initialized kernel object with a specific width and height.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAreaMax/init(device:kernelWidth:kernelHeight:)
func NewImageAreaMinWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSImageAreaMin {
	instance := getMPSImageAreaMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return MPSImageAreaMinFromID(rv)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageBox] class.
var (
	_MPSImageBoxClass     MPSImageBoxClass
	_MPSImageBoxClassOnce sync.Once
)

func getMPSImageBoxClass() MPSImageBoxClass {
	_MPSImageBoxClassOnce.Do(func() {
		_MPSImageBoxClass = MPSImageBoxClass{class: objc.GetClass("MPSImageBox")}
	})
	return _MPSImageBoxClass
}

// GetMPSImageBoxClass returns the class object for MPSImageBox.
func GetMPSImageBoxClass() MPSImageBoxClass {
	return getMPSImageBoxClass()
}

type MPSImageBoxClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageBoxClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageBoxClass) Alloc() MPSImageBox {
	rv := objc.Send[MPSImageBox](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that convolves an image with a given kernel of odd width and
// height.
//
// # Overview
//
// The kernel elements all have equal weight, achieving a blur effect (each
// result is the unweighted average of the surrounding pixels). This allows
// for much faster algorithms, especially for larger blur radii. The box
// height and width must be odd numbers.
//
// The box blur is a separable filter and the Metal Performance Shaders
// framework will act accordingly to give best performance for
// multi-dimensional blurs.
//
// # Methods
//
//   - [MPSImageBox.InitWithDeviceKernelWidthKernelHeight]: Initializes a box filter.
//
// # Properties
//
//   - [MPSImageBox.KernelHeight]: The height of the filter window. Must be an odd number.
//   - [MPSImageBox.KernelWidth]: The width of the filter window. Must be an odd number.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageBox
type MPSImageBox struct {
	MPSUnaryImageKernel
}

// MPSImageBoxFromID constructs a [MPSImageBox] from an objc.ID.
//
// A filter that convolves an image with a given kernel of odd width and
// height.
func MPSImageBoxFromID(id objc.ID) MPSImageBox {
	return MPSImageBox{MPSUnaryImageKernel: MPSUnaryImageKernelFromID(id)}
}

// NOTE: MPSImageBox adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageBox] class.
//
// # Methods
//
//   - [IMPSImageBox.InitWithDeviceKernelWidthKernelHeight]: Initializes a box filter.
//
// # Properties
//
//   - [IMPSImageBox.KernelHeight]: The height of the filter window. Must be an odd number.
//   - [IMPSImageBox.KernelWidth]: The width of the filter window. Must be an odd number.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageBox
type IMPSImageBox interface {
	IMPSUnaryImageKernel

	// Topic: Methods

	// Initializes a box filter.
	InitWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSImageBox

	// Topic: Properties

	// The height of the filter window. Must be an odd number.
	KernelHeight() uint
	// The width of the filter window. Must be an odd number.
	KernelWidth() uint
}

// Init initializes the instance.
func (i MPSImageBox) Init() MPSImageBox {
	rv := objc.Send[MPSImageBox](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageBox) Autorelease() MPSImageBox {
	rv := objc.Send[MPSImageBox](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageBox creates a new MPSImageBox instance.
func NewMPSImageBox() MPSImageBox {
	class := getMPSImageBoxClass()
	rv := objc.Send[MPSImageBox](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageBoxWithCoder(aDecoder foundation.INSCoder) MPSImageBox {
	instance := getMPSImageBoxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageBoxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageBox/init(coder:device:)
func NewImageBoxWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageBox {
	instance := getMPSImageBoxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageBoxFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSUnaryImageKernel/init(device:)
func NewImageBoxWithDevice(device metal.MTLDevice) MPSImageBox {
	instance := getMPSImageBoxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageBoxFromID(rv)
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
func NewImageBoxWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSImageBox {
	instance := getMPSImageBoxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return MPSImageBoxFromID(rv)
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
func (i MPSImageBox) InitWithDeviceKernelWidthKernelHeight(device metal.MTLDevice, kernelWidth uint, kernelHeight uint) MPSImageBox {
	rv := objc.Send[MPSImageBox](i.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	return rv
}

// The height of the filter window. Must be an odd number.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageBox/kernelHeight
func (i MPSImageBox) KernelHeight() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("kernelHeight"))
	return rv
}

// The width of the filter window. Must be an odd number.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageBox/kernelWidth
func (i MPSImageBox) KernelWidth() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("kernelWidth"))
	return rv
}

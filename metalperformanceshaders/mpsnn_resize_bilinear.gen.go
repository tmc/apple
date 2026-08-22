// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNResizeBilinear] class.
var (
	_MPSNNResizeBilinearClass     MPSNNResizeBilinearClass
	_MPSNNResizeBilinearClassOnce sync.Once
)

func getMPSNNResizeBilinearClass() MPSNNResizeBilinearClass {
	_MPSNNResizeBilinearClassOnce.Do(func() {
		_MPSNNResizeBilinearClass = MPSNNResizeBilinearClass{class: objc.GetClass("MPSNNResizeBilinear")}
	})
	return _MPSNNResizeBilinearClass
}

// GetMPSNNResizeBilinearClass returns the class object for MPSNNResizeBilinear.
func GetMPSNNResizeBilinearClass() MPSNNResizeBilinearClass {
	return getMPSNNResizeBilinearClass()
}

type MPSNNResizeBilinearClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNResizeBilinearClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNResizeBilinearClass) Alloc() MPSNNResizeBilinear {
	rv := objc.Send[MPSNNResizeBilinear](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A bilinear resizing filter.
//
// # Initializers
//
//   - [MPSNNResizeBilinear.InitWithDeviceResizeWidthResizeHeightAlignCorners]
//
// # Instance Properties
//
//   - [MPSNNResizeBilinear.AlignCorners]
//   - [MPSNNResizeBilinear.ResizeHeight]
//   - [MPSNNResizeBilinear.ResizeWidth]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNResizeBilinear
type MPSNNResizeBilinear struct {
	MPSCNNKernel
}

// MPSNNResizeBilinearFromID constructs a [MPSNNResizeBilinear] from an objc.ID.
//
// A bilinear resizing filter.
func MPSNNResizeBilinearFromID(id objc.ID) MPSNNResizeBilinear {
	return MPSNNResizeBilinear{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSNNResizeBilinear adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNResizeBilinear] class.
//
// # Initializers
//
//   - [IMPSNNResizeBilinear.InitWithDeviceResizeWidthResizeHeightAlignCorners]
//
// # Instance Properties
//
//   - [IMPSNNResizeBilinear.AlignCorners]
//   - [IMPSNNResizeBilinear.ResizeHeight]
//   - [IMPSNNResizeBilinear.ResizeWidth]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNResizeBilinear
type IMPSNNResizeBilinear interface {
	IMPSCNNKernel

	// Topic: Initializers

	InitWithDeviceResizeWidthResizeHeightAlignCorners(device metal.MTLDevice, resizeWidth uint, resizeHeight uint, alignCorners bool) MPSNNResizeBilinear

	// Topic: Instance Properties

	AlignCorners() bool
	ResizeHeight() uint
	ResizeWidth() uint
}

// Init initializes the instance.
func (r MPSNNResizeBilinear) Init() MPSNNResizeBilinear {
	rv := objc.Send[MPSNNResizeBilinear](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNResizeBilinear) Autorelease() MPSNNResizeBilinear {
	rv := objc.Send[MPSNNResizeBilinear](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNResizeBilinear creates a new MPSNNResizeBilinear instance.
func NewMPSNNResizeBilinear() MPSNNResizeBilinear {
	class := getMPSNNResizeBilinearClass()
	rv := objc.Send[MPSNNResizeBilinear](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewResizeBilinearWithCoder(aDecoder foundation.INSCoder) MPSNNResizeBilinear {
	instance := getMPSNNResizeBilinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNResizeBilinearFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNResizeBilinear/init(coder:device:)
func NewResizeBilinearWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNResizeBilinear {
	instance := getMPSNNResizeBilinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNResizeBilinearFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewResizeBilinearWithDevice(device metal.MTLDevice) MPSNNResizeBilinear {
	instance := getMPSNNResizeBilinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNResizeBilinearFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNResizeBilinear/init(device:resizeWidth:resizeHeight:alignCorners:)
func NewResizeBilinearWithDeviceResizeWidthResizeHeightAlignCorners(device metal.MTLDevice, resizeWidth uint, resizeHeight uint, alignCorners bool) MPSNNResizeBilinear {
	instance := getMPSNNResizeBilinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resizeWidth:resizeHeight:alignCorners:"), device, resizeWidth, resizeHeight, alignCorners)
	return MPSNNResizeBilinearFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNResizeBilinear/init(device:resizeWidth:resizeHeight:alignCorners:)
func (r MPSNNResizeBilinear) InitWithDeviceResizeWidthResizeHeightAlignCorners(device metal.MTLDevice, resizeWidth uint, resizeHeight uint, alignCorners bool) MPSNNResizeBilinear {
	rv := objc.Send[MPSNNResizeBilinear](r.ID, objc.Sel("initWithDevice:resizeWidth:resizeHeight:alignCorners:"), device, resizeWidth, resizeHeight, alignCorners)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNResizeBilinear/alignCorners
func (r MPSNNResizeBilinear) AlignCorners() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("alignCorners"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNResizeBilinear/resizeHeight
func (r MPSNNResizeBilinear) ResizeHeight() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("resizeHeight"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNResizeBilinear/resizeWidth
func (r MPSNNResizeBilinear) ResizeWidth() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("resizeWidth"))
	return rv
}

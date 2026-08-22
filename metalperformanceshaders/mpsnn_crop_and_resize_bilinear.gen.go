// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNCropAndResizeBilinear] class.
var (
	_MPSNNCropAndResizeBilinearClass     MPSNNCropAndResizeBilinearClass
	_MPSNNCropAndResizeBilinearClassOnce sync.Once
)

func getMPSNNCropAndResizeBilinearClass() MPSNNCropAndResizeBilinearClass {
	_MPSNNCropAndResizeBilinearClassOnce.Do(func() {
		_MPSNNCropAndResizeBilinearClass = MPSNNCropAndResizeBilinearClass{class: objc.GetClass("MPSNNCropAndResizeBilinear")}
	})
	return _MPSNNCropAndResizeBilinearClass
}

// GetMPSNNCropAndResizeBilinearClass returns the class object for MPSNNCropAndResizeBilinear.
func GetMPSNNCropAndResizeBilinearClass() MPSNNCropAndResizeBilinearClass {
	return getMPSNNCropAndResizeBilinearClass()
}

type MPSNNCropAndResizeBilinearClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNCropAndResizeBilinearClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNCropAndResizeBilinearClass) Alloc() MPSNNCropAndResizeBilinear {
	rv := objc.Send[MPSNNCropAndResizeBilinear](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A cropping and bilinear resizing filter.
//
// # Initializers
//
//   - [MPSNNCropAndResizeBilinear.InitWithDeviceResizeWidthResizeHeightNumberOfRegionsRegions]
//
// # Instance Properties
//
//   - [MPSNNCropAndResizeBilinear.NumberOfRegions]
//   - [MPSNNCropAndResizeBilinear.Regions]
//   - [MPSNNCropAndResizeBilinear.ResizeHeight]
//   - [MPSNNCropAndResizeBilinear.ResizeWidth]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCropAndResizeBilinear
type MPSNNCropAndResizeBilinear struct {
	MPSCNNKernel
}

// MPSNNCropAndResizeBilinearFromID constructs a [MPSNNCropAndResizeBilinear] from an objc.ID.
//
// A cropping and bilinear resizing filter.
func MPSNNCropAndResizeBilinearFromID(id objc.ID) MPSNNCropAndResizeBilinear {
	return MPSNNCropAndResizeBilinear{MPSCNNKernel: MPSCNNKernelFromID(id)}
}

// NOTE: MPSNNCropAndResizeBilinear adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNCropAndResizeBilinear] class.
//
// # Initializers
//
//   - [IMPSNNCropAndResizeBilinear.InitWithDeviceResizeWidthResizeHeightNumberOfRegionsRegions]
//
// # Instance Properties
//
//   - [IMPSNNCropAndResizeBilinear.NumberOfRegions]
//   - [IMPSNNCropAndResizeBilinear.Regions]
//   - [IMPSNNCropAndResizeBilinear.ResizeHeight]
//   - [IMPSNNCropAndResizeBilinear.ResizeWidth]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCropAndResizeBilinear
type IMPSNNCropAndResizeBilinear interface {
	IMPSCNNKernel

	// Topic: Initializers

	InitWithDeviceResizeWidthResizeHeightNumberOfRegionsRegions(device metal.MTLDevice, resizeWidth uint, resizeHeight uint, numberOfRegions uint, regions *MPSRegion) MPSNNCropAndResizeBilinear

	// Topic: Instance Properties

	NumberOfRegions() uint
	Regions() *MPSRegion
	ResizeHeight() uint
	ResizeWidth() uint
}

// Init initializes the instance.
func (c MPSNNCropAndResizeBilinear) Init() MPSNNCropAndResizeBilinear {
	rv := objc.Send[MPSNNCropAndResizeBilinear](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSNNCropAndResizeBilinear) Autorelease() MPSNNCropAndResizeBilinear {
	rv := objc.Send[MPSNNCropAndResizeBilinear](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNCropAndResizeBilinear creates a new MPSNNCropAndResizeBilinear instance.
func NewMPSNNCropAndResizeBilinear() MPSNNCropAndResizeBilinear {
	class := getMPSNNCropAndResizeBilinearClass()
	rv := objc.Send[MPSNNCropAndResizeBilinear](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCropAndResizeBilinearWithCoder(aDecoder foundation.INSCoder) MPSNNCropAndResizeBilinear {
	instance := getMPSNNCropAndResizeBilinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNCropAndResizeBilinearFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCropAndResizeBilinear/init(coder:device:)
func NewCropAndResizeBilinearWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNCropAndResizeBilinear {
	instance := getMPSNNCropAndResizeBilinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNCropAndResizeBilinearFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNKernel/init(device:)
func NewCropAndResizeBilinearWithDevice(device metal.MTLDevice) MPSNNCropAndResizeBilinear {
	instance := getMPSNNCropAndResizeBilinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNCropAndResizeBilinearFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCropAndResizeBilinear/init(device:resizeWidth:resizeHeight:numberOfRegions:regions:)
func NewCropAndResizeBilinearWithDeviceResizeWidthResizeHeightNumberOfRegionsRegions(device metal.MTLDevice, resizeWidth uint, resizeHeight uint, numberOfRegions uint, regions *MPSRegion) MPSNNCropAndResizeBilinear {
	instance := getMPSNNCropAndResizeBilinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resizeWidth:resizeHeight:numberOfRegions:regions:"), device, resizeWidth, resizeHeight, numberOfRegions, unsafe.Pointer(regions))
	return MPSNNCropAndResizeBilinearFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCropAndResizeBilinear/init(device:resizeWidth:resizeHeight:numberOfRegions:regions:)
func (c MPSNNCropAndResizeBilinear) InitWithDeviceResizeWidthResizeHeightNumberOfRegionsRegions(device metal.MTLDevice, resizeWidth uint, resizeHeight uint, numberOfRegions uint, regions *MPSRegion) MPSNNCropAndResizeBilinear {
	rv := objc.Send[MPSNNCropAndResizeBilinear](c.ID, objc.Sel("initWithDevice:resizeWidth:resizeHeight:numberOfRegions:regions:"), device, resizeWidth, resizeHeight, numberOfRegions, unsafe.Pointer(regions))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCropAndResizeBilinear/numberOfRegions
func (c MPSNNCropAndResizeBilinear) NumberOfRegions() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("numberOfRegions"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCropAndResizeBilinear/regions
func (c MPSNNCropAndResizeBilinear) Regions() *MPSRegion {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("regions"))
	return (*MPSRegion)(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCropAndResizeBilinear/resizeHeight
func (c MPSNNCropAndResizeBilinear) ResizeHeight() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("resizeHeight"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCropAndResizeBilinear/resizeWidth
func (c MPSNNCropAndResizeBilinear) ResizeWidth() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("resizeWidth"))
	return rv
}

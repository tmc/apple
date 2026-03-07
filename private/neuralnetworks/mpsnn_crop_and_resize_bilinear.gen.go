// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
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

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNCropAndResizeBilinearClass) Alloc() MPSNNCropAndResizeBilinear {
	rv := objc.Send[MPSNNCropAndResizeBilinear](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNCropAndResizeBilinear.CopyWithZoneDevice]
//   - [MPSNNCropAndResizeBilinear.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//   - [MPSNNCropAndResizeBilinear.NumberOfRegions]
//   - [MPSNNCropAndResizeBilinear.Regions]
//   - [MPSNNCropAndResizeBilinear.ResizeHeight]
//   - [MPSNNCropAndResizeBilinear.ResizeWidth]
//   - [MPSNNCropAndResizeBilinear.InitWithCoderDevice]
//   - [MPSNNCropAndResizeBilinear.InitWithDeviceResizeWidthResizeHeightNumberOfRegionsRegions]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNCropAndResizeBilinear
type MPSNNCropAndResizeBilinear struct {
	MPSCNNKernel
}

// MPSNNCropAndResizeBilinearFromID constructs a [MPSNNCropAndResizeBilinear] from an objc.ID.
func MPSNNCropAndResizeBilinearFromID(id objc.ID) MPSNNCropAndResizeBilinear {
	return MPSNNCropAndResizeBilinear{MPSCNNKernel: MPSCNNKernelFromID(id)}
}
// Ensure MPSNNCropAndResizeBilinear implements IMPSNNCropAndResizeBilinear.
var _ IMPSNNCropAndResizeBilinear = MPSNNCropAndResizeBilinear{}





// An interface definition for the [MPSNNCropAndResizeBilinear] class.
//
// # Methods
//
//   - [IMPSNNCropAndResizeBilinear.CopyWithZoneDevice]
//   - [IMPSNNCropAndResizeBilinear.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//   - [IMPSNNCropAndResizeBilinear.NumberOfRegions]
//   - [IMPSNNCropAndResizeBilinear.Regions]
//   - [IMPSNNCropAndResizeBilinear.ResizeHeight]
//   - [IMPSNNCropAndResizeBilinear.ResizeWidth]
//   - [IMPSNNCropAndResizeBilinear.InitWithCoderDevice]
//   - [IMPSNNCropAndResizeBilinear.InitWithDeviceResizeWidthResizeHeightNumberOfRegionsRegions]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNCropAndResizeBilinear
type IMPSNNCropAndResizeBilinear interface {
	IMPSCNNKernel

	// Topic: Methods

	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
	NumberOfRegions() uint64
	Regions() objectivec.IObject
	ResizeHeight() uint64
	ResizeWidth() uint64
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNCropAndResizeBilinear
	InitWithDeviceResizeWidthResizeHeightNumberOfRegionsRegions(device objectivec.IObject, width uint64, height uint64, regions uint64, regions2 unsafe.Pointer) MPSNNCropAndResizeBilinear
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






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNCropAndResizeBilinear/initWithCoder:device:
func NewCropAndResizeBilinearWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNCropAndResizeBilinear {
	instance := getMPSNNCropAndResizeBilinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNCropAndResizeBilinearFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNCropAndResizeBilinear/initWithDevice:resizeWidth:resizeHeight:numberOfRegions:regions:
func NewCropAndResizeBilinearWithDeviceResizeWidthResizeHeightNumberOfRegionsRegions(device objectivec.IObject, width uint64, height uint64, regions uint64, regions2 unsafe.Pointer) MPSNNCropAndResizeBilinear {
	instance := getMPSNNCropAndResizeBilinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resizeWidth:resizeHeight:numberOfRegions:regions:"), device, width, height, regions, regions2)
	return MPSNNCropAndResizeBilinearFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNCropAndResizeBilinear/copyWithZone:device:
func (c MPSNNCropAndResizeBilinear) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNCropAndResizeBilinear/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (c MPSNNCropAndResizeBilinear) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNCropAndResizeBilinear/initWithCoder:device:
func (c MPSNNCropAndResizeBilinear) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNCropAndResizeBilinear {
	rv := objc.Send[MPSNNCropAndResizeBilinear](c.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNCropAndResizeBilinear/initWithDevice:resizeWidth:resizeHeight:numberOfRegions:regions:
func (c MPSNNCropAndResizeBilinear) InitWithDeviceResizeWidthResizeHeightNumberOfRegionsRegions(device objectivec.IObject, width uint64, height uint64, regions uint64, regions2 unsafe.Pointer) MPSNNCropAndResizeBilinear {
	rv := objc.Send[MPSNNCropAndResizeBilinear](c.ID, objc.Sel("initWithDevice:resizeWidth:resizeHeight:numberOfRegions:regions:"), device, width, height, regions, regions2)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNCropAndResizeBilinear/libraryInfo:
func (_MPSNNCropAndResizeBilinearClass MPSNNCropAndResizeBilinearClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNCropAndResizeBilinearClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNCropAndResizeBilinear/numberOfRegions
func (c MPSNNCropAndResizeBilinear) NumberOfRegions() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("numberOfRegions"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNCropAndResizeBilinear/regions
func (c MPSNNCropAndResizeBilinear) Regions() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("regions"))
	return objectivec.Object{ID: rv}
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNCropAndResizeBilinear/resizeHeight
func (c MPSNNCropAndResizeBilinear) ResizeHeight() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("resizeHeight"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNCropAndResizeBilinear/resizeWidth
func (c MPSNNCropAndResizeBilinear) ResizeWidth() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("resizeWidth"))
	return rv
}


















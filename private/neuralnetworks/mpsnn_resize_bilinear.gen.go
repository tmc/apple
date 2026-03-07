// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
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

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNResizeBilinearClass) Alloc() MPSNNResizeBilinear {
	rv := objc.Send[MPSNNResizeBilinear](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNResizeBilinear.AlignCorners]
//   - [MPSNNResizeBilinear.CopyWithZoneDevice]
//   - [MPSNNResizeBilinear.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//   - [MPSNNResizeBilinear.ResizeHeight]
//   - [MPSNNResizeBilinear.ResizeWidth]
//   - [MPSNNResizeBilinear.InitWithCoderDevice]
//   - [MPSNNResizeBilinear.InitWithDeviceResizeWidthResizeHeightAlignCorners]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNResizeBilinear
type MPSNNResizeBilinear struct {
	MPSCNNKernel
}

// MPSNNResizeBilinearFromID constructs a [MPSNNResizeBilinear] from an objc.ID.
func MPSNNResizeBilinearFromID(id objc.ID) MPSNNResizeBilinear {
	return MPSNNResizeBilinear{MPSCNNKernel: MPSCNNKernelFromID(id)}
}
// Ensure MPSNNResizeBilinear implements IMPSNNResizeBilinear.
var _ IMPSNNResizeBilinear = MPSNNResizeBilinear{}





// An interface definition for the [MPSNNResizeBilinear] class.
//
// # Methods
//
//   - [IMPSNNResizeBilinear.AlignCorners]
//   - [IMPSNNResizeBilinear.CopyWithZoneDevice]
//   - [IMPSNNResizeBilinear.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//   - [IMPSNNResizeBilinear.ResizeHeight]
//   - [IMPSNNResizeBilinear.ResizeWidth]
//   - [IMPSNNResizeBilinear.InitWithCoderDevice]
//   - [IMPSNNResizeBilinear.InitWithDeviceResizeWidthResizeHeightAlignCorners]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNResizeBilinear
type IMPSNNResizeBilinear interface {
	IMPSCNNKernel

	// Topic: Methods

	AlignCorners() bool
	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
	ResizeHeight() uint64
	ResizeWidth() uint64
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNResizeBilinear
	InitWithDeviceResizeWidthResizeHeightAlignCorners(device objectivec.IObject, width uint64, height uint64, corners bool) MPSNNResizeBilinear
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






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNResizeBilinear/initWithCoder:device:
func NewResizeBilinearWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNResizeBilinear {
	instance := getMPSNNResizeBilinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNResizeBilinearFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNResizeBilinear/initWithDevice:resizeWidth:resizeHeight:alignCorners:
func NewResizeBilinearWithDeviceResizeWidthResizeHeightAlignCorners(device objectivec.IObject, width uint64, height uint64, corners bool) MPSNNResizeBilinear {
	instance := getMPSNNResizeBilinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resizeWidth:resizeHeight:alignCorners:"), device, width, height, corners)
	return MPSNNResizeBilinearFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNResizeBilinear/copyWithZone:device:
func (r MPSNNResizeBilinear) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNResizeBilinear/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (r MPSNNResizeBilinear) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNResizeBilinear/initWithCoder:device:
func (r MPSNNResizeBilinear) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNResizeBilinear {
	rv := objc.Send[MPSNNResizeBilinear](r.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNResizeBilinear/initWithDevice:resizeWidth:resizeHeight:alignCorners:
func (r MPSNNResizeBilinear) InitWithDeviceResizeWidthResizeHeightAlignCorners(device objectivec.IObject, width uint64, height uint64, corners bool) MPSNNResizeBilinear {
	rv := objc.Send[MPSNNResizeBilinear](r.ID, objc.Sel("initWithDevice:resizeWidth:resizeHeight:alignCorners:"), device, width, height, corners)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNResizeBilinear/libraryInfo:
func (_MPSNNResizeBilinearClass MPSNNResizeBilinearClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNResizeBilinearClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNResizeBilinear/alignCorners
func (r MPSNNResizeBilinear) AlignCorners() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("alignCorners"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNResizeBilinear/resizeHeight
func (r MPSNNResizeBilinear) ResizeHeight() uint64 {
	rv := objc.Send[uint64](r.ID, objc.Sel("resizeHeight"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNResizeBilinear/resizeWidth
func (r MPSNNResizeBilinear) ResizeWidth() uint64 {
	rv := objc.Send[uint64](r.ID, objc.Sel("resizeWidth"))
	return rv
}


















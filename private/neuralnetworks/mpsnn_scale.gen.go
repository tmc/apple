// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNScale] class.
var (
	_MPSNNScaleClass     MPSNNScaleClass
	_MPSNNScaleClassOnce sync.Once
)

func getMPSNNScaleClass() MPSNNScaleClass {
	_MPSNNScaleClassOnce.Do(func() {
		_MPSNNScaleClass = MPSNNScaleClass{class: objc.GetClass("MPSNNScale")}
	})
	return _MPSNNScaleClass
}

// GetMPSNNScaleClass returns the class object for MPSNNScale.
func GetMPSNNScaleClass() MPSNNScaleClass {
	return getMPSNNScaleClass()
}

type MPSNNScaleClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNScaleClass) Alloc() MPSNNScale {
	rv := objc.Send[MPSNNScale](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNScale.CopyWithZoneDevice]
//   - [MPSNNScale.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//   - [MPSNNScale.SetEdgeMode]
//   - [MPSNNScale.SetLabel]
//   - [MPSNNScale.SetOptions]
//   - [MPSNNScale.InitWithCoderDevice]
//   - [MPSNNScale.InitWithDevice]
//   - [MPSNNScale.InitWithDeviceTransformProviderHandleOutputSizeScaleClass]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScale
type MPSNNScale struct {
	MPSCNNKernel
}

// MPSNNScaleFromID constructs a [MPSNNScale] from an objc.ID.
func MPSNNScaleFromID(id objc.ID) MPSNNScale {
	return MPSNNScale{MPSCNNKernel: MPSCNNKernelFromID(id)}
}
// Ensure MPSNNScale implements IMPSNNScale.
var _ IMPSNNScale = MPSNNScale{}





// An interface definition for the [MPSNNScale] class.
//
// # Methods
//
//   - [IMPSNNScale.CopyWithZoneDevice]
//   - [IMPSNNScale.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//   - [IMPSNNScale.SetEdgeMode]
//   - [IMPSNNScale.SetLabel]
//   - [IMPSNNScale.SetOptions]
//   - [IMPSNNScale.InitWithCoderDevice]
//   - [IMPSNNScale.InitWithDevice]
//   - [IMPSNNScale.InitWithDeviceTransformProviderHandleOutputSizeScaleClass]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScale
type IMPSNNScale interface {
	IMPSCNNKernel

	// Topic: Methods

	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
	SetEdgeMode(mode uint64)
	SetLabel(label objectivec.IObject)
	SetOptions(options uint64)
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNScale
	InitWithDevice(device objectivec.IObject) MPSNNScale
	InitWithDeviceTransformProviderHandleOutputSizeScaleClass(device objectivec.IObject, provider objectivec.IObject, handle objectivec.IObject, size objectivec.IObject, class objc.Class) MPSNNScale
}





// Init initializes the instance.
func (s MPSNNScale) Init() MPSNNScale {
	rv := objc.Send[MPSNNScale](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MPSNNScale) Autorelease() MPSNNScale {
	rv := objc.Send[MPSNNScale](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNScale creates a new MPSNNScale instance.
func NewMPSNNScale() MPSNNScale {
	class := getMPSNNScaleClass()
	rv := objc.Send[MPSNNScale](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScale/initWithCoder:device:
func NewScaleWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNScale {
	instance := getMPSNNScaleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNScaleFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScale/initWithDevice:
func NewScaleWithDevice(device objectivec.IObject) MPSNNScale {
	instance := getMPSNNScaleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNScaleFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScale/initWithDevice:transformProvider:handle:outputSize:scaleClass:
func NewScaleWithDeviceTransformProviderHandleOutputSizeScaleClass(device objectivec.IObject, provider objectivec.IObject, handle objectivec.IObject, size objectivec.IObject, class objc.Class) MPSNNScale {
	instance := getMPSNNScaleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:transformProvider:handle:outputSize:scaleClass:"), device, provider, handle, size, class)
	return MPSNNScaleFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScale/copyWithZone:device:
func (s MPSNNScale) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScale/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (s MPSNNScale) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScale/setEdgeMode:
func (s MPSNNScale) SetEdgeMode(mode uint64) {
	objc.Send[objc.ID](s.ID, objc.Sel("setEdgeMode:"), mode)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScale/setLabel:
func (s MPSNNScale) SetLabel(label objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("setLabel:"), label)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScale/setOptions:
func (s MPSNNScale) SetOptions(options uint64) {
	objc.Send[objc.ID](s.ID, objc.Sel("setOptions:"), options)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScale/initWithCoder:device:
func (s MPSNNScale) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNScale {
	rv := objc.Send[MPSNNScale](s.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScale/initWithDevice:
func (s MPSNNScale) InitWithDevice(device objectivec.IObject) MPSNNScale {
	rv := objc.Send[MPSNNScale](s.ID, objc.Sel("initWithDevice:"), device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScale/initWithDevice:transformProvider:handle:outputSize:scaleClass:
func (s MPSNNScale) InitWithDeviceTransformProviderHandleOutputSizeScaleClass(device objectivec.IObject, provider objectivec.IObject, handle objectivec.IObject, size objectivec.IObject, class objc.Class) MPSNNScale {
	rv := objc.Send[MPSNNScale](s.ID, objc.Sel("initWithDevice:transformProvider:handle:outputSize:scaleClass:"), device, provider, handle, size, class)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNScale/libraryInfo:
func (_MPSNNScaleClass MPSNNScaleClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNScaleClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}























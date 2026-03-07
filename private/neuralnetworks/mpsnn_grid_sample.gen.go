// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNGridSample] class.
var (
	_MPSNNGridSampleClass     MPSNNGridSampleClass
	_MPSNNGridSampleClassOnce sync.Once
)

func getMPSNNGridSampleClass() MPSNNGridSampleClass {
	_MPSNNGridSampleClassOnce.Do(func() {
		_MPSNNGridSampleClass = MPSNNGridSampleClass{class: objc.GetClass("MPSNNGridSample")}
	})
	return _MPSNNGridSampleClass
}

// GetMPSNNGridSampleClass returns the class object for MPSNNGridSample.
func GetMPSNNGridSampleClass() MPSNNGridSampleClass {
	return getMPSNNGridSampleClass()
}

type MPSNNGridSampleClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNGridSampleClass) Alloc() MPSNNGridSample {
	rv := objc.Send[MPSNNGridSample](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNGridSample.CopyWithZoneDevice]
//   - [MPSNNGridSample.UseGridValueAsInputCoordinate]
//   - [MPSNNGridSample.SetUseGridValueAsInputCoordinate]
//   - [MPSNNGridSample.InitWithCoderDevice]
//   - [MPSNNGridSample.InitWithDevice]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGridSample
type MPSNNGridSample struct {
	MPSCNNBinaryKernel
}

// MPSNNGridSampleFromID constructs a [MPSNNGridSample] from an objc.ID.
func MPSNNGridSampleFromID(id objc.ID) MPSNNGridSample {
	return MPSNNGridSample{MPSCNNBinaryKernel: MPSCNNBinaryKernelFromID(id)}
}
// Ensure MPSNNGridSample implements IMPSNNGridSample.
var _ IMPSNNGridSample = MPSNNGridSample{}





// An interface definition for the [MPSNNGridSample] class.
//
// # Methods
//
//   - [IMPSNNGridSample.CopyWithZoneDevice]
//   - [IMPSNNGridSample.UseGridValueAsInputCoordinate]
//   - [IMPSNNGridSample.SetUseGridValueAsInputCoordinate]
//   - [IMPSNNGridSample.InitWithCoderDevice]
//   - [IMPSNNGridSample.InitWithDevice]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGridSample
type IMPSNNGridSample interface {
	IMPSCNNBinaryKernel

	// Topic: Methods

	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	UseGridValueAsInputCoordinate() bool
	SetUseGridValueAsInputCoordinate(value bool)
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNGridSample
	InitWithDevice(device objectivec.IObject) MPSNNGridSample
}





// Init initializes the instance.
func (g MPSNNGridSample) Init() MPSNNGridSample {
	rv := objc.Send[MPSNNGridSample](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSNNGridSample) Autorelease() MPSNNGridSample {
	rv := objc.Send[MPSNNGridSample](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNGridSample creates a new MPSNNGridSample instance.
func NewMPSNNGridSample() MPSNNGridSample {
	class := getMPSNNGridSampleClass()
	rv := objc.Send[MPSNNGridSample](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGridSample/initWithCoder:device:
func NewGridSampleWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNGridSample {
	instance := getMPSNNGridSampleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNGridSampleFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGridSample/initWithDevice:
func NewGridSampleWithDevice(device objectivec.IObject) MPSNNGridSample {
	instance := getMPSNNGridSampleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNGridSampleFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGridSample/copyWithZone:device:
func (g MPSNNGridSample) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGridSample/initWithCoder:device:
func (g MPSNNGridSample) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNGridSample {
	rv := objc.Send[MPSNNGridSample](g.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGridSample/initWithDevice:
func (g MPSNNGridSample) InitWithDevice(device objectivec.IObject) MPSNNGridSample {
	rv := objc.Send[MPSNNGridSample](g.ID, objc.Sel("initWithDevice:"), device)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGridSample/libraryInfo:
func (_MPSNNGridSampleClass MPSNNGridSampleClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNGridSampleClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGridSample/useGridValueAsInputCoordinate
func (g MPSNNGridSample) UseGridValueAsInputCoordinate() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("useGridValueAsInputCoordinate"))
	return rv
}
func (g MPSNNGridSample) SetUseGridValueAsInputCoordinate(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setUseGridValueAsInputCoordinate:"), value)
}


















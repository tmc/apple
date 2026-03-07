// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNInitialGradient] class.
var (
	_MPSNNInitialGradientClass     MPSNNInitialGradientClass
	_MPSNNInitialGradientClassOnce sync.Once
)

func getMPSNNInitialGradientClass() MPSNNInitialGradientClass {
	_MPSNNInitialGradientClassOnce.Do(func() {
		_MPSNNInitialGradientClass = MPSNNInitialGradientClass{class: objc.GetClass("MPSNNInitialGradient")}
	})
	return _MPSNNInitialGradientClass
}

// GetMPSNNInitialGradientClass returns the class object for MPSNNInitialGradient.
func GetMPSNNInitialGradientClass() MPSNNInitialGradientClass {
	return getMPSNNInitialGradientClass()
}

type MPSNNInitialGradientClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNInitialGradientClass) Alloc() MPSNNInitialGradient {
	rv := objc.Send[MPSNNInitialGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNInitialGradient.CopyWithZoneDevice]
//   - [MPSNNInitialGradient.InitWithCoderDevice]
//   - [MPSNNInitialGradient.InitWithDevice]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNInitialGradient
type MPSNNInitialGradient struct {
	MPSCNNKernel
}

// MPSNNInitialGradientFromID constructs a [MPSNNInitialGradient] from an objc.ID.
func MPSNNInitialGradientFromID(id objc.ID) MPSNNInitialGradient {
	return MPSNNInitialGradient{MPSCNNKernel: MPSCNNKernelFromID(id)}
}
// Ensure MPSNNInitialGradient implements IMPSNNInitialGradient.
var _ IMPSNNInitialGradient = MPSNNInitialGradient{}





// An interface definition for the [MPSNNInitialGradient] class.
//
// # Methods
//
//   - [IMPSNNInitialGradient.CopyWithZoneDevice]
//   - [IMPSNNInitialGradient.InitWithCoderDevice]
//   - [IMPSNNInitialGradient.InitWithDevice]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNInitialGradient
type IMPSNNInitialGradient interface {
	IMPSCNNKernel

	// Topic: Methods

	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNInitialGradient
	InitWithDevice(device objectivec.IObject) MPSNNInitialGradient
}





// Init initializes the instance.
func (i MPSNNInitialGradient) Init() MPSNNInitialGradient {
	rv := objc.Send[MPSNNInitialGradient](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSNNInitialGradient) Autorelease() MPSNNInitialGradient {
	rv := objc.Send[MPSNNInitialGradient](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNInitialGradient creates a new MPSNNInitialGradient instance.
func NewMPSNNInitialGradient() MPSNNInitialGradient {
	class := getMPSNNInitialGradientClass()
	rv := objc.Send[MPSNNInitialGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNInitialGradient/initWithCoder:device:
func NewInitialGradientWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNInitialGradient {
	instance := getMPSNNInitialGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNInitialGradientFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNInitialGradient/initWithDevice:
func NewInitialGradientWithDevice(device objectivec.IObject) MPSNNInitialGradient {
	instance := getMPSNNInitialGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNInitialGradientFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNInitialGradient/copyWithZone:device:
func (i MPSNNInitialGradient) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNInitialGradient/initWithCoder:device:
func (i MPSNNInitialGradient) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNInitialGradient {
	rv := objc.Send[MPSNNInitialGradient](i.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNInitialGradient/initWithDevice:
func (i MPSNNInitialGradient) InitWithDevice(device objectivec.IObject) MPSNNInitialGradient {
	rv := objc.Send[MPSNNInitialGradient](i.ID, objc.Sel("initWithDevice:"), device)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNInitialGradient/libraryInfo:
func (_MPSNNInitialGradientClass MPSNNInitialGradientClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNInitialGradientClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}























// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNSlice] class.
var (
	_MPSNNSliceClass     MPSNNSliceClass
	_MPSNNSliceClassOnce sync.Once
)

func getMPSNNSliceClass() MPSNNSliceClass {
	_MPSNNSliceClassOnce.Do(func() {
		_MPSNNSliceClass = MPSNNSliceClass{class: objc.GetClass("MPSNNSlice")}
	})
	return _MPSNNSliceClass
}

// GetMPSNNSliceClass returns the class object for MPSNNSlice.
func GetMPSNNSliceClass() MPSNNSliceClass {
	return getMPSNNSliceClass()
}

type MPSNNSliceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNSliceClass) Alloc() MPSNNSlice {
	rv := objc.Send[MPSNNSlice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNSlice.InitWithCoderDevice]
//   - [MPSNNSlice.InitWithDevice]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNSlice
type MPSNNSlice struct {
	MPSCNNKernel
}

// MPSNNSliceFromID constructs a [MPSNNSlice] from an objc.ID.
func MPSNNSliceFromID(id objc.ID) MPSNNSlice {
	return MPSNNSlice{MPSCNNKernel: MPSCNNKernelFromID(id)}
}
// Ensure MPSNNSlice implements IMPSNNSlice.
var _ IMPSNNSlice = MPSNNSlice{}





// An interface definition for the [MPSNNSlice] class.
//
// # Methods
//
//   - [IMPSNNSlice.InitWithCoderDevice]
//   - [IMPSNNSlice.InitWithDevice]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNSlice
type IMPSNNSlice interface {
	IMPSCNNKernel

	// Topic: Methods

	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNSlice
	InitWithDevice(device objectivec.IObject) MPSNNSlice
}





// Init initializes the instance.
func (s MPSNNSlice) Init() MPSNNSlice {
	rv := objc.Send[MPSNNSlice](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MPSNNSlice) Autorelease() MPSNNSlice {
	rv := objc.Send[MPSNNSlice](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNSlice creates a new MPSNNSlice instance.
func NewMPSNNSlice() MPSNNSlice {
	class := getMPSNNSliceClass()
	rv := objc.Send[MPSNNSlice](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNSlice/initWithCoder:device:
func NewSliceWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNSlice {
	instance := getMPSNNSliceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNSliceFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNSlice/initWithDevice:
func NewSliceWithDevice(device objectivec.IObject) MPSNNSlice {
	instance := getMPSNNSliceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNSliceFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNSlice/initWithCoder:device:
func (s MPSNNSlice) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNSlice {
	rv := objc.Send[MPSNNSlice](s.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNSlice/initWithDevice:
func (s MPSNNSlice) InitWithDevice(device objectivec.IObject) MPSNNSlice {
	rv := objc.Send[MPSNNSlice](s.ID, objc.Sel("initWithDevice:"), device)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNSlice/libraryInfo:
func (_MPSNNSliceClass MPSNNSliceClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNSliceClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}























// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReduceUnary] class.
var (
	_MPSNNReduceUnaryClass     MPSNNReduceUnaryClass
	_MPSNNReduceUnaryClassOnce sync.Once
)

func getMPSNNReduceUnaryClass() MPSNNReduceUnaryClass {
	_MPSNNReduceUnaryClassOnce.Do(func() {
		_MPSNNReduceUnaryClass = MPSNNReduceUnaryClass{class: objc.GetClass("MPSNNReduceUnary")}
	})
	return _MPSNNReduceUnaryClass
}

// GetMPSNNReduceUnaryClass returns the class object for MPSNNReduceUnary.
func GetMPSNNReduceUnaryClass() MPSNNReduceUnaryClass {
	return getMPSNNReduceUnaryClass()
}

type MPSNNReduceUnaryClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceUnaryClass) Alloc() MPSNNReduceUnary {
	rv := objc.Send[MPSNNReduceUnary](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReduceUnary.ClipRectSource]
//   - [MPSNNReduceUnary.SetClipRectSource]
//   - [MPSNNReduceUnary.CopyWithZoneDevice]
//   - [MPSNNReduceUnary.SetWeightValue]
//   - [MPSNNReduceUnary.InitWithCoderDevice]
//   - [MPSNNReduceUnary.InitWithDevice]
//   - [MPSNNReduceUnary.InitWithDeviceReduceOperation]
//   - [MPSNNReduceUnary.Offset]
//   - [MPSNNReduceUnary.SetOffset]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary
type MPSNNReduceUnary struct {
	MPSCNNKernel
}

// MPSNNReduceUnaryFromID constructs a [MPSNNReduceUnary] from an objc.ID.
func MPSNNReduceUnaryFromID(id objc.ID) MPSNNReduceUnary {
	return MPSNNReduceUnary{MPSCNNKernel: MPSCNNKernelFromID(id)}
}
// Ensure MPSNNReduceUnary implements IMPSNNReduceUnary.
var _ IMPSNNReduceUnary = MPSNNReduceUnary{}





// An interface definition for the [MPSNNReduceUnary] class.
//
// # Methods
//
//   - [IMPSNNReduceUnary.ClipRectSource]
//   - [IMPSNNReduceUnary.SetClipRectSource]
//   - [IMPSNNReduceUnary.CopyWithZoneDevice]
//   - [IMPSNNReduceUnary.SetWeightValue]
//   - [IMPSNNReduceUnary.InitWithCoderDevice]
//   - [IMPSNNReduceUnary.InitWithDevice]
//   - [IMPSNNReduceUnary.InitWithDeviceReduceOperation]
//   - [IMPSNNReduceUnary.Offset]
//   - [IMPSNNReduceUnary.SetOffset]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary
type IMPSNNReduceUnary interface {
	IMPSCNNKernel

	// Topic: Methods

	ClipRectSource() objectivec.IObject
	SetClipRectSource(value objectivec.IObject)
	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	SetWeightValue(value float32)
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceUnary
	InitWithDevice(device objectivec.IObject) MPSNNReduceUnary
	InitWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceUnary
	Offset() objectivec.IObject
	SetOffset(value objectivec.IObject)
}





// Init initializes the instance.
func (r MPSNNReduceUnary) Init() MPSNNReduceUnary {
	rv := objc.Send[MPSNNReduceUnary](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceUnary) Autorelease() MPSNNReduceUnary {
	rv := objc.Send[MPSNNReduceUnary](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceUnary creates a new MPSNNReduceUnary instance.
func NewMPSNNReduceUnary() MPSNNReduceUnary {
	class := getMPSNNReduceUnaryClass()
	rv := objc.Send[MPSNNReduceUnary](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithCoder:device:
func NewReduceUnaryWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceUnary {
	instance := getMPSNNReduceUnaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReduceUnaryFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithDevice:
func NewReduceUnaryWithDevice(device objectivec.IObject) MPSNNReduceUnary {
	instance := getMPSNNReduceUnaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceUnaryFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithDevice:reduceOperation:
func NewReduceUnaryWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceUnary {
	instance := getMPSNNReduceUnaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return MPSNNReduceUnaryFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/copyWithZone:device:
func (r MPSNNReduceUnary) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/setWeightValue:
func (r MPSNNReduceUnary) SetWeightValue(value float32) {
	objc.Send[objc.ID](r.ID, objc.Sel("setWeightValue:"), value)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithCoder:device:
func (r MPSNNReduceUnary) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceUnary {
	rv := objc.Send[MPSNNReduceUnary](r.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithDevice:
func (r MPSNNReduceUnary) InitWithDevice(device objectivec.IObject) MPSNNReduceUnary {
	rv := objc.Send[MPSNNReduceUnary](r.ID, objc.Sel("initWithDevice:"), device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithDevice:reduceOperation:
func (r MPSNNReduceUnary) InitWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceUnary {
	rv := objc.Send[MPSNNReduceUnary](r.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/libraryInfo:
func (_MPSNNReduceUnaryClass MPSNNReduceUnaryClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNReduceUnaryClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/clipRectSource
func (r MPSNNReduceUnary) ClipRectSource() objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("clipRectSource"))
	return objectivec.Object{ID: rv}
}
func (r MPSNNReduceUnary) SetClipRectSource(value objectivec.IObject) {
	objc.Send[struct{}](r.ID, objc.Sel("setClipRectSource:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/offset
func (r MPSNNReduceUnary) Offset() objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("offset"))
	return objectivec.Object{ID: rv}
}
func (r MPSNNReduceUnary) SetOffset(value objectivec.IObject) {
	objc.Send[struct{}](r.ID, objc.Sel("setOffset:"), value)
}


















// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReduceBinary] class.
var (
	_MPSNNReduceBinaryClass     MPSNNReduceBinaryClass
	_MPSNNReduceBinaryClassOnce sync.Once
)

func getMPSNNReduceBinaryClass() MPSNNReduceBinaryClass {
	_MPSNNReduceBinaryClassOnce.Do(func() {
		_MPSNNReduceBinaryClass = MPSNNReduceBinaryClass{class: objc.GetClass("MPSNNReduceBinary")}
	})
	return _MPSNNReduceBinaryClass
}

// GetMPSNNReduceBinaryClass returns the class object for MPSNNReduceBinary.
func GetMPSNNReduceBinaryClass() MPSNNReduceBinaryClass {
	return getMPSNNReduceBinaryClass()
}

type MPSNNReduceBinaryClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceBinaryClass) Alloc() MPSNNReduceBinary {
	rv := objc.Send[MPSNNReduceBinary](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReduceBinary.CopyWithZoneDevice]
//   - [MPSNNReduceBinary.PrimarySourceClipRect]
//   - [MPSNNReduceBinary.SetPrimarySourceClipRect]
//   - [MPSNNReduceBinary.ReduceOp]
//   - [MPSNNReduceBinary.SecondarySourceClipRect]
//   - [MPSNNReduceBinary.SetSecondarySourceClipRect]
//   - [MPSNNReduceBinary.InitWithCoderDevice]
//   - [MPSNNReduceBinary.InitWithDevice]
//   - [MPSNNReduceBinary.InitWithDeviceReduceOperation]
//   - [MPSNNReduceBinary.PrimaryOffset]
//   - [MPSNNReduceBinary.SetPrimaryOffset]
//   - [MPSNNReduceBinary.SecondaryOffset]
//   - [MPSNNReduceBinary.SetSecondaryOffset]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceBinary
type MPSNNReduceBinary struct {
	MPSCNNBinaryKernel
}

// MPSNNReduceBinaryFromID constructs a [MPSNNReduceBinary] from an objc.ID.
func MPSNNReduceBinaryFromID(id objc.ID) MPSNNReduceBinary {
	return MPSNNReduceBinary{MPSCNNBinaryKernel: MPSCNNBinaryKernelFromID(id)}
}
// Ensure MPSNNReduceBinary implements IMPSNNReduceBinary.
var _ IMPSNNReduceBinary = MPSNNReduceBinary{}





// An interface definition for the [MPSNNReduceBinary] class.
//
// # Methods
//
//   - [IMPSNNReduceBinary.CopyWithZoneDevice]
//   - [IMPSNNReduceBinary.PrimarySourceClipRect]
//   - [IMPSNNReduceBinary.SetPrimarySourceClipRect]
//   - [IMPSNNReduceBinary.ReduceOp]
//   - [IMPSNNReduceBinary.SecondarySourceClipRect]
//   - [IMPSNNReduceBinary.SetSecondarySourceClipRect]
//   - [IMPSNNReduceBinary.InitWithCoderDevice]
//   - [IMPSNNReduceBinary.InitWithDevice]
//   - [IMPSNNReduceBinary.InitWithDeviceReduceOperation]
//   - [IMPSNNReduceBinary.PrimaryOffset]
//   - [IMPSNNReduceBinary.SetPrimaryOffset]
//   - [IMPSNNReduceBinary.SecondaryOffset]
//   - [IMPSNNReduceBinary.SetSecondaryOffset]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceBinary
type IMPSNNReduceBinary interface {
	IMPSCNNBinaryKernel

	// Topic: Methods

	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	PrimarySourceClipRect() objectivec.IObject
	SetPrimarySourceClipRect(value objectivec.IObject)
	ReduceOp() int
	SecondarySourceClipRect() objectivec.IObject
	SetSecondarySourceClipRect(value objectivec.IObject)
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceBinary
	InitWithDevice(device objectivec.IObject) MPSNNReduceBinary
	InitWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceBinary
	PrimaryOffset() objectivec.IObject
	SetPrimaryOffset(value objectivec.IObject)
	SecondaryOffset() objectivec.IObject
	SetSecondaryOffset(value objectivec.IObject)
}





// Init initializes the instance.
func (r MPSNNReduceBinary) Init() MPSNNReduceBinary {
	rv := objc.Send[MPSNNReduceBinary](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceBinary) Autorelease() MPSNNReduceBinary {
	rv := objc.Send[MPSNNReduceBinary](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceBinary creates a new MPSNNReduceBinary instance.
func NewMPSNNReduceBinary() MPSNNReduceBinary {
	class := getMPSNNReduceBinaryClass()
	rv := objc.Send[MPSNNReduceBinary](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceBinary/initWithCoder:device:
func NewReduceBinaryWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceBinary {
	instance := getMPSNNReduceBinaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReduceBinaryFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceBinary/initWithDevice:
func NewReduceBinaryWithDevice(device objectivec.IObject) MPSNNReduceBinary {
	instance := getMPSNNReduceBinaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceBinaryFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceBinary/initWithDevice:reduceOperation:
func NewReduceBinaryWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceBinary {
	instance := getMPSNNReduceBinaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return MPSNNReduceBinaryFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceBinary/copyWithZone:device:
func (r MPSNNReduceBinary) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceBinary/reduceOp
func (r MPSNNReduceBinary) ReduceOp() int {
	rv := objc.Send[int](r.ID, objc.Sel("reduceOp"))
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceBinary/initWithCoder:device:
func (r MPSNNReduceBinary) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceBinary {
	rv := objc.Send[MPSNNReduceBinary](r.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceBinary/initWithDevice:
func (r MPSNNReduceBinary) InitWithDevice(device objectivec.IObject) MPSNNReduceBinary {
	rv := objc.Send[MPSNNReduceBinary](r.ID, objc.Sel("initWithDevice:"), device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceBinary/initWithDevice:reduceOperation:
func (r MPSNNReduceBinary) InitWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceBinary {
	rv := objc.Send[MPSNNReduceBinary](r.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceBinary/libraryInfo:
func (_MPSNNReduceBinaryClass MPSNNReduceBinaryClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNReduceBinaryClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceBinary/primaryOffset
func (r MPSNNReduceBinary) PrimaryOffset() objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("primaryOffset"))
	return objectivec.Object{ID: rv}
}
func (r MPSNNReduceBinary) SetPrimaryOffset(value objectivec.IObject) {
	objc.Send[struct{}](r.ID, objc.Sel("setPrimaryOffset:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceBinary/primarySourceClipRect
func (r MPSNNReduceBinary) PrimarySourceClipRect() objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("primarySourceClipRect"))
	return objectivec.Object{ID: rv}
}
func (r MPSNNReduceBinary) SetPrimarySourceClipRect(value objectivec.IObject) {
	objc.Send[struct{}](r.ID, objc.Sel("setPrimarySourceClipRect:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceBinary/secondaryOffset
func (r MPSNNReduceBinary) SecondaryOffset() objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("secondaryOffset"))
	return objectivec.Object{ID: rv}
}
func (r MPSNNReduceBinary) SetSecondaryOffset(value objectivec.IObject) {
	objc.Send[struct{}](r.ID, objc.Sel("setSecondaryOffset:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceBinary/secondarySourceClipRect
func (r MPSNNReduceBinary) SecondarySourceClipRect() objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("secondarySourceClipRect"))
	return objectivec.Object{ID: rv}
}
func (r MPSNNReduceBinary) SetSecondarySourceClipRect(value objectivec.IObject) {
	objc.Send[struct{}](r.ID, objc.Sel("setSecondarySourceClipRect:"), value)
}


















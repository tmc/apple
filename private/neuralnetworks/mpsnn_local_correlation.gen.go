// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNLocalCorrelation] class.
var (
	_MPSNNLocalCorrelationClass     MPSNNLocalCorrelationClass
	_MPSNNLocalCorrelationClassOnce sync.Once
)

func getMPSNNLocalCorrelationClass() MPSNNLocalCorrelationClass {
	_MPSNNLocalCorrelationClassOnce.Do(func() {
		_MPSNNLocalCorrelationClass = MPSNNLocalCorrelationClass{class: objc.GetClass("MPSNNLocalCorrelation")}
	})
	return _MPSNNLocalCorrelationClass
}

// GetMPSNNLocalCorrelationClass returns the class object for MPSNNLocalCorrelation.
func GetMPSNNLocalCorrelationClass() MPSNNLocalCorrelationClass {
	return getMPSNNLocalCorrelationClass()
}

type MPSNNLocalCorrelationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNLocalCorrelationClass) Alloc() MPSNNLocalCorrelation {
	rv := objc.Send[MPSNNLocalCorrelation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNLocalCorrelation.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset]
//   - [MPSNNLocalCorrelation.StrideInX]
//   - [MPSNNLocalCorrelation.SetStrideInX]
//   - [MPSNNLocalCorrelation.StrideInY]
//   - [MPSNNLocalCorrelation.SetStrideInY]
//   - [MPSNNLocalCorrelation.WindowInX]
//   - [MPSNNLocalCorrelation.SetWindowInX]
//   - [MPSNNLocalCorrelation.WindowInY]
//   - [MPSNNLocalCorrelation.SetWindowInY]
//   - [MPSNNLocalCorrelation.InitWithDeviceWindowInXWindowInYStrideInXStrideInY]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLocalCorrelation
type MPSNNLocalCorrelation struct {
	MPSNNReduceBinary
}

// MPSNNLocalCorrelationFromID constructs a [MPSNNLocalCorrelation] from an objc.ID.
func MPSNNLocalCorrelationFromID(id objc.ID) MPSNNLocalCorrelation {
	return MPSNNLocalCorrelation{MPSNNReduceBinary: MPSNNReduceBinaryFromID(id)}
}
// Ensure MPSNNLocalCorrelation implements IMPSNNLocalCorrelation.
var _ IMPSNNLocalCorrelation = MPSNNLocalCorrelation{}





// An interface definition for the [MPSNNLocalCorrelation] class.
//
// # Methods
//
//   - [IMPSNNLocalCorrelation.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset]
//   - [IMPSNNLocalCorrelation.StrideInX]
//   - [IMPSNNLocalCorrelation.SetStrideInX]
//   - [IMPSNNLocalCorrelation.StrideInY]
//   - [IMPSNNLocalCorrelation.SetStrideInY]
//   - [IMPSNNLocalCorrelation.WindowInX]
//   - [IMPSNNLocalCorrelation.SetWindowInX]
//   - [IMPSNNLocalCorrelation.WindowInY]
//   - [IMPSNNLocalCorrelation.SetWindowInY]
//   - [IMPSNNLocalCorrelation.InitWithDeviceWindowInXWindowInYStrideInXStrideInY]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLocalCorrelation
type IMPSNNLocalCorrelation interface {
	IMPSNNReduceBinary

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject, offset2 objectivec.IObject, offset3 objectivec.IObject) objectivec.IObject
	StrideInX() uint64
	SetStrideInX(value uint64)
	StrideInY() uint64
	SetStrideInY(value uint64)
	WindowInX() uint64
	SetWindowInX(value uint64)
	WindowInY() uint64
	SetWindowInY(value uint64)
	InitWithDeviceWindowInXWindowInYStrideInXStrideInY(device objectivec.IObject, x uint64, y uint64, x2 uint64, y2 uint64) MPSNNLocalCorrelation
}





// Init initializes the instance.
func (l MPSNNLocalCorrelation) Init() MPSNNLocalCorrelation {
	rv := objc.Send[MPSNNLocalCorrelation](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MPSNNLocalCorrelation) Autorelease() MPSNNLocalCorrelation {
	rv := objc.Send[MPSNNLocalCorrelation](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNLocalCorrelation creates a new MPSNNLocalCorrelation instance.
func NewMPSNNLocalCorrelation() MPSNNLocalCorrelation {
	class := getMPSNNLocalCorrelationClass()
	rv := objc.Send[MPSNNLocalCorrelation](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLocalCorrelation/initWithCoder:device:
func NewLocalCorrelationWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNLocalCorrelation {
	instance := getMPSNNLocalCorrelationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNLocalCorrelationFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLocalCorrelation/initWithDevice:
func NewLocalCorrelationWithDevice(device objectivec.IObject) MPSNNLocalCorrelation {
	instance := getMPSNNLocalCorrelationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNLocalCorrelationFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceBinary/initWithDevice:reduceOperation:
func NewLocalCorrelationWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNLocalCorrelation {
	instance := getMPSNNLocalCorrelationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return MPSNNLocalCorrelationFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLocalCorrelation/initWithDevice:windowInX:windowInY:strideInX:strideInY:
func NewLocalCorrelationWithDeviceWindowInXWindowInYStrideInXStrideInY(device objectivec.IObject, x uint64, y uint64, x2 uint64, y2 uint64) MPSNNLocalCorrelation {
	instance := getMPSNNLocalCorrelationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:windowInX:windowInY:strideInX:strideInY:"), device, x, y, x2, y2)
	return MPSNNLocalCorrelationFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLocalCorrelation/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:primaryOffset:secondaryOffset:kernelOffset:
func (l MPSNNLocalCorrelation) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject, offset2 objectivec.IObject, offset3 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:primaryOffset:secondaryOffset:kernelOffset:"), images, states, method, offset, offset2, offset3)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLocalCorrelation/initWithDevice:windowInX:windowInY:strideInX:strideInY:
func (l MPSNNLocalCorrelation) InitWithDeviceWindowInXWindowInYStrideInXStrideInY(device objectivec.IObject, x uint64, y uint64, x2 uint64, y2 uint64) MPSNNLocalCorrelation {
	rv := objc.Send[MPSNNLocalCorrelation](l.ID, objc.Sel("initWithDevice:windowInX:windowInY:strideInX:strideInY:"), device, x, y, x2, y2)
	return rv
}












// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLocalCorrelation/strideInX
func (l MPSNNLocalCorrelation) StrideInX() uint64 {
	rv := objc.Send[uint64](l.ID, objc.Sel("strideInX"))
	return rv
}
func (l MPSNNLocalCorrelation) SetStrideInX(value uint64) {
	objc.Send[struct{}](l.ID, objc.Sel("setStrideInX:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLocalCorrelation/strideInY
func (l MPSNNLocalCorrelation) StrideInY() uint64 {
	rv := objc.Send[uint64](l.ID, objc.Sel("strideInY"))
	return rv
}
func (l MPSNNLocalCorrelation) SetStrideInY(value uint64) {
	objc.Send[struct{}](l.ID, objc.Sel("setStrideInY:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLocalCorrelation/windowInX
func (l MPSNNLocalCorrelation) WindowInX() uint64 {
	rv := objc.Send[uint64](l.ID, objc.Sel("windowInX"))
	return rv
}
func (l MPSNNLocalCorrelation) SetWindowInX(value uint64) {
	objc.Send[struct{}](l.ID, objc.Sel("setWindowInX:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNLocalCorrelation/windowInY
func (l MPSNNLocalCorrelation) WindowInY() uint64 {
	rv := objc.Send[uint64](l.ID, objc.Sel("windowInY"))
	return rv
}
func (l MPSNNLocalCorrelation) SetWindowInY(value uint64) {
	objc.Send[struct{}](l.ID, objc.Sel("setWindowInY:"), value)
}


















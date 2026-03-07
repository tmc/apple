// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReduceRowMax] class.
var (
	_MPSNNReduceRowMaxClass     MPSNNReduceRowMaxClass
	_MPSNNReduceRowMaxClassOnce sync.Once
)

func getMPSNNReduceRowMaxClass() MPSNNReduceRowMaxClass {
	_MPSNNReduceRowMaxClassOnce.Do(func() {
		_MPSNNReduceRowMaxClass = MPSNNReduceRowMaxClass{class: objc.GetClass("MPSNNReduceRowMax")}
	})
	return _MPSNNReduceRowMaxClass
}

// GetMPSNNReduceRowMaxClass returns the class object for MPSNNReduceRowMax.
func GetMPSNNReduceRowMaxClass() MPSNNReduceRowMaxClass {
	return getMPSNNReduceRowMaxClass()
}

type MPSNNReduceRowMaxClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceRowMaxClass) Alloc() MPSNNReduceRowMax {
	rv := objc.Send[MPSNNReduceRowMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReduceRowMax.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowMax
type MPSNNReduceRowMax struct {
	MPSNNReduceUnary
}

// MPSNNReduceRowMaxFromID constructs a [MPSNNReduceRowMax] from an objc.ID.
func MPSNNReduceRowMaxFromID(id objc.ID) MPSNNReduceRowMax {
	return MPSNNReduceRowMax{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}
// Ensure MPSNNReduceRowMax implements IMPSNNReduceRowMax.
var _ IMPSNNReduceRowMax = MPSNNReduceRowMax{}





// An interface definition for the [MPSNNReduceRowMax] class.
//
// # Methods
//
//   - [IMPSNNReduceRowMax.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowMax
type IMPSNNReduceRowMax interface {
	IMPSNNReduceUnary

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
}





// Init initializes the instance.
func (r MPSNNReduceRowMax) Init() MPSNNReduceRowMax {
	rv := objc.Send[MPSNNReduceRowMax](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceRowMax) Autorelease() MPSNNReduceRowMax {
	rv := objc.Send[MPSNNReduceRowMax](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceRowMax creates a new MPSNNReduceRowMax instance.
func NewMPSNNReduceRowMax() MPSNNReduceRowMax {
	class := getMPSNNReduceRowMaxClass()
	rv := objc.Send[MPSNNReduceRowMax](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowMax/initWithCoder:device:
func NewReduceRowMaxWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceRowMax {
	instance := getMPSNNReduceRowMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReduceRowMaxFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowMax/initWithDevice:
func NewReduceRowMaxWithDevice(device objectivec.IObject) MPSNNReduceRowMax {
	instance := getMPSNNReduceRowMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceRowMaxFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithDevice:reduceOperation:
func NewReduceRowMaxWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceRowMax {
	instance := getMPSNNReduceRowMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return MPSNNReduceRowMaxFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowMax/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (r MPSNNReduceRowMax) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}



























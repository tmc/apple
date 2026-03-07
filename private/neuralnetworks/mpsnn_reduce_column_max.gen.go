// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReduceColumnMax] class.
var (
	_MPSNNReduceColumnMaxClass     MPSNNReduceColumnMaxClass
	_MPSNNReduceColumnMaxClassOnce sync.Once
)

func getMPSNNReduceColumnMaxClass() MPSNNReduceColumnMaxClass {
	_MPSNNReduceColumnMaxClassOnce.Do(func() {
		_MPSNNReduceColumnMaxClass = MPSNNReduceColumnMaxClass{class: objc.GetClass("MPSNNReduceColumnMax")}
	})
	return _MPSNNReduceColumnMaxClass
}

// GetMPSNNReduceColumnMaxClass returns the class object for MPSNNReduceColumnMax.
func GetMPSNNReduceColumnMaxClass() MPSNNReduceColumnMaxClass {
	return getMPSNNReduceColumnMaxClass()
}

type MPSNNReduceColumnMaxClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceColumnMaxClass) Alloc() MPSNNReduceColumnMax {
	rv := objc.Send[MPSNNReduceColumnMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReduceColumnMax.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnMax
type MPSNNReduceColumnMax struct {
	MPSNNReduceUnary
}

// MPSNNReduceColumnMaxFromID constructs a [MPSNNReduceColumnMax] from an objc.ID.
func MPSNNReduceColumnMaxFromID(id objc.ID) MPSNNReduceColumnMax {
	return MPSNNReduceColumnMax{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}
// Ensure MPSNNReduceColumnMax implements IMPSNNReduceColumnMax.
var _ IMPSNNReduceColumnMax = MPSNNReduceColumnMax{}





// An interface definition for the [MPSNNReduceColumnMax] class.
//
// # Methods
//
//   - [IMPSNNReduceColumnMax.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnMax
type IMPSNNReduceColumnMax interface {
	IMPSNNReduceUnary

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
}





// Init initializes the instance.
func (r MPSNNReduceColumnMax) Init() MPSNNReduceColumnMax {
	rv := objc.Send[MPSNNReduceColumnMax](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceColumnMax) Autorelease() MPSNNReduceColumnMax {
	rv := objc.Send[MPSNNReduceColumnMax](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceColumnMax creates a new MPSNNReduceColumnMax instance.
func NewMPSNNReduceColumnMax() MPSNNReduceColumnMax {
	class := getMPSNNReduceColumnMaxClass()
	rv := objc.Send[MPSNNReduceColumnMax](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnMax/initWithCoder:device:
func NewReduceColumnMaxWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceColumnMax {
	instance := getMPSNNReduceColumnMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReduceColumnMaxFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnMax/initWithDevice:
func NewReduceColumnMaxWithDevice(device objectivec.IObject) MPSNNReduceColumnMax {
	instance := getMPSNNReduceColumnMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceColumnMaxFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithDevice:reduceOperation:
func NewReduceColumnMaxWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceColumnMax {
	instance := getMPSNNReduceColumnMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return MPSNNReduceColumnMaxFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnMax/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (r MPSNNReduceColumnMax) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}



























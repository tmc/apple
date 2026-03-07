// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReduceFeatureChannelsArgumentMax] class.
var (
	_MPSNNReduceFeatureChannelsArgumentMaxClass     MPSNNReduceFeatureChannelsArgumentMaxClass
	_MPSNNReduceFeatureChannelsArgumentMaxClassOnce sync.Once
)

func getMPSNNReduceFeatureChannelsArgumentMaxClass() MPSNNReduceFeatureChannelsArgumentMaxClass {
	_MPSNNReduceFeatureChannelsArgumentMaxClassOnce.Do(func() {
		_MPSNNReduceFeatureChannelsArgumentMaxClass = MPSNNReduceFeatureChannelsArgumentMaxClass{class: objc.GetClass("MPSNNReduceFeatureChannelsArgumentMax")}
	})
	return _MPSNNReduceFeatureChannelsArgumentMaxClass
}

// GetMPSNNReduceFeatureChannelsArgumentMaxClass returns the class object for MPSNNReduceFeatureChannelsArgumentMax.
func GetMPSNNReduceFeatureChannelsArgumentMaxClass() MPSNNReduceFeatureChannelsArgumentMaxClass {
	return getMPSNNReduceFeatureChannelsArgumentMaxClass()
}

type MPSNNReduceFeatureChannelsArgumentMaxClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceFeatureChannelsArgumentMaxClass) Alloc() MPSNNReduceFeatureChannelsArgumentMax {
	rv := objc.Send[MPSNNReduceFeatureChannelsArgumentMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReduceFeatureChannelsArgumentMax.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsArgumentMax
type MPSNNReduceFeatureChannelsArgumentMax struct {
	MPSNNReduceUnary
}

// MPSNNReduceFeatureChannelsArgumentMaxFromID constructs a [MPSNNReduceFeatureChannelsArgumentMax] from an objc.ID.
func MPSNNReduceFeatureChannelsArgumentMaxFromID(id objc.ID) MPSNNReduceFeatureChannelsArgumentMax {
	return MPSNNReduceFeatureChannelsArgumentMax{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}
// Ensure MPSNNReduceFeatureChannelsArgumentMax implements IMPSNNReduceFeatureChannelsArgumentMax.
var _ IMPSNNReduceFeatureChannelsArgumentMax = MPSNNReduceFeatureChannelsArgumentMax{}





// An interface definition for the [MPSNNReduceFeatureChannelsArgumentMax] class.
//
// # Methods
//
//   - [IMPSNNReduceFeatureChannelsArgumentMax.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsArgumentMax
type IMPSNNReduceFeatureChannelsArgumentMax interface {
	IMPSNNReduceUnary

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
}





// Init initializes the instance.
func (r MPSNNReduceFeatureChannelsArgumentMax) Init() MPSNNReduceFeatureChannelsArgumentMax {
	rv := objc.Send[MPSNNReduceFeatureChannelsArgumentMax](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceFeatureChannelsArgumentMax) Autorelease() MPSNNReduceFeatureChannelsArgumentMax {
	rv := objc.Send[MPSNNReduceFeatureChannelsArgumentMax](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceFeatureChannelsArgumentMax creates a new MPSNNReduceFeatureChannelsArgumentMax instance.
func NewMPSNNReduceFeatureChannelsArgumentMax() MPSNNReduceFeatureChannelsArgumentMax {
	class := getMPSNNReduceFeatureChannelsArgumentMaxClass()
	rv := objc.Send[MPSNNReduceFeatureChannelsArgumentMax](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsArgumentMax/initWithCoder:device:
func NewReduceFeatureChannelsArgumentMaxWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceFeatureChannelsArgumentMax {
	instance := getMPSNNReduceFeatureChannelsArgumentMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReduceFeatureChannelsArgumentMaxFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsArgumentMax/initWithDevice:
func NewReduceFeatureChannelsArgumentMaxWithDevice(device objectivec.IObject) MPSNNReduceFeatureChannelsArgumentMax {
	instance := getMPSNNReduceFeatureChannelsArgumentMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceFeatureChannelsArgumentMaxFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithDevice:reduceOperation:
func NewReduceFeatureChannelsArgumentMaxWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceFeatureChannelsArgumentMax {
	instance := getMPSNNReduceFeatureChannelsArgumentMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return MPSNNReduceFeatureChannelsArgumentMaxFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsArgumentMax/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (r MPSNNReduceFeatureChannelsArgumentMax) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}



























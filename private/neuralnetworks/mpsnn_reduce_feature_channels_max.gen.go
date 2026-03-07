// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReduceFeatureChannelsMax] class.
var (
	_MPSNNReduceFeatureChannelsMaxClass     MPSNNReduceFeatureChannelsMaxClass
	_MPSNNReduceFeatureChannelsMaxClassOnce sync.Once
)

func getMPSNNReduceFeatureChannelsMaxClass() MPSNNReduceFeatureChannelsMaxClass {
	_MPSNNReduceFeatureChannelsMaxClassOnce.Do(func() {
		_MPSNNReduceFeatureChannelsMaxClass = MPSNNReduceFeatureChannelsMaxClass{class: objc.GetClass("MPSNNReduceFeatureChannelsMax")}
	})
	return _MPSNNReduceFeatureChannelsMaxClass
}

// GetMPSNNReduceFeatureChannelsMaxClass returns the class object for MPSNNReduceFeatureChannelsMax.
func GetMPSNNReduceFeatureChannelsMaxClass() MPSNNReduceFeatureChannelsMaxClass {
	return getMPSNNReduceFeatureChannelsMaxClass()
}

type MPSNNReduceFeatureChannelsMaxClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceFeatureChannelsMaxClass) Alloc() MPSNNReduceFeatureChannelsMax {
	rv := objc.Send[MPSNNReduceFeatureChannelsMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReduceFeatureChannelsMax.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsMax
type MPSNNReduceFeatureChannelsMax struct {
	MPSNNReduceUnary
}

// MPSNNReduceFeatureChannelsMaxFromID constructs a [MPSNNReduceFeatureChannelsMax] from an objc.ID.
func MPSNNReduceFeatureChannelsMaxFromID(id objc.ID) MPSNNReduceFeatureChannelsMax {
	return MPSNNReduceFeatureChannelsMax{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}
// Ensure MPSNNReduceFeatureChannelsMax implements IMPSNNReduceFeatureChannelsMax.
var _ IMPSNNReduceFeatureChannelsMax = MPSNNReduceFeatureChannelsMax{}





// An interface definition for the [MPSNNReduceFeatureChannelsMax] class.
//
// # Methods
//
//   - [IMPSNNReduceFeatureChannelsMax.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsMax
type IMPSNNReduceFeatureChannelsMax interface {
	IMPSNNReduceUnary

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
}





// Init initializes the instance.
func (r MPSNNReduceFeatureChannelsMax) Init() MPSNNReduceFeatureChannelsMax {
	rv := objc.Send[MPSNNReduceFeatureChannelsMax](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceFeatureChannelsMax) Autorelease() MPSNNReduceFeatureChannelsMax {
	rv := objc.Send[MPSNNReduceFeatureChannelsMax](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceFeatureChannelsMax creates a new MPSNNReduceFeatureChannelsMax instance.
func NewMPSNNReduceFeatureChannelsMax() MPSNNReduceFeatureChannelsMax {
	class := getMPSNNReduceFeatureChannelsMaxClass()
	rv := objc.Send[MPSNNReduceFeatureChannelsMax](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsMax/initWithCoder:device:
func NewReduceFeatureChannelsMaxWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceFeatureChannelsMax {
	instance := getMPSNNReduceFeatureChannelsMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReduceFeatureChannelsMaxFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsMax/initWithDevice:
func NewReduceFeatureChannelsMaxWithDevice(device objectivec.IObject) MPSNNReduceFeatureChannelsMax {
	instance := getMPSNNReduceFeatureChannelsMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceFeatureChannelsMaxFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithDevice:reduceOperation:
func NewReduceFeatureChannelsMaxWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceFeatureChannelsMax {
	instance := getMPSNNReduceFeatureChannelsMaxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return MPSNNReduceFeatureChannelsMaxFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsMax/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (r MPSNNReduceFeatureChannelsMax) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}



























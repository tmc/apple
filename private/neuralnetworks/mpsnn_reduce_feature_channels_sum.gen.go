// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReduceFeatureChannelsSum] class.
var (
	_MPSNNReduceFeatureChannelsSumClass     MPSNNReduceFeatureChannelsSumClass
	_MPSNNReduceFeatureChannelsSumClassOnce sync.Once
)

func getMPSNNReduceFeatureChannelsSumClass() MPSNNReduceFeatureChannelsSumClass {
	_MPSNNReduceFeatureChannelsSumClassOnce.Do(func() {
		_MPSNNReduceFeatureChannelsSumClass = MPSNNReduceFeatureChannelsSumClass{class: objc.GetClass("MPSNNReduceFeatureChannelsSum")}
	})
	return _MPSNNReduceFeatureChannelsSumClass
}

// GetMPSNNReduceFeatureChannelsSumClass returns the class object for MPSNNReduceFeatureChannelsSum.
func GetMPSNNReduceFeatureChannelsSumClass() MPSNNReduceFeatureChannelsSumClass {
	return getMPSNNReduceFeatureChannelsSumClass()
}

type MPSNNReduceFeatureChannelsSumClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceFeatureChannelsSumClass) Alloc() MPSNNReduceFeatureChannelsSum {
	rv := objc.Send[MPSNNReduceFeatureChannelsSum](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReduceFeatureChannelsSum.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//   - [MPSNNReduceFeatureChannelsSum.Weight]
//   - [MPSNNReduceFeatureChannelsSum.SetWeight]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsSum
type MPSNNReduceFeatureChannelsSum struct {
	MPSNNReduceUnary
}

// MPSNNReduceFeatureChannelsSumFromID constructs a [MPSNNReduceFeatureChannelsSum] from an objc.ID.
func MPSNNReduceFeatureChannelsSumFromID(id objc.ID) MPSNNReduceFeatureChannelsSum {
	return MPSNNReduceFeatureChannelsSum{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}
// Ensure MPSNNReduceFeatureChannelsSum implements IMPSNNReduceFeatureChannelsSum.
var _ IMPSNNReduceFeatureChannelsSum = MPSNNReduceFeatureChannelsSum{}





// An interface definition for the [MPSNNReduceFeatureChannelsSum] class.
//
// # Methods
//
//   - [IMPSNNReduceFeatureChannelsSum.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//   - [IMPSNNReduceFeatureChannelsSum.Weight]
//   - [IMPSNNReduceFeatureChannelsSum.SetWeight]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsSum
type IMPSNNReduceFeatureChannelsSum interface {
	IMPSNNReduceUnary

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
	Weight() float32
	SetWeight(value float32)
}





// Init initializes the instance.
func (r MPSNNReduceFeatureChannelsSum) Init() MPSNNReduceFeatureChannelsSum {
	rv := objc.Send[MPSNNReduceFeatureChannelsSum](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceFeatureChannelsSum) Autorelease() MPSNNReduceFeatureChannelsSum {
	rv := objc.Send[MPSNNReduceFeatureChannelsSum](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceFeatureChannelsSum creates a new MPSNNReduceFeatureChannelsSum instance.
func NewMPSNNReduceFeatureChannelsSum() MPSNNReduceFeatureChannelsSum {
	class := getMPSNNReduceFeatureChannelsSumClass()
	rv := objc.Send[MPSNNReduceFeatureChannelsSum](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsSum/initWithCoder:device:
func NewReduceFeatureChannelsSumWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceFeatureChannelsSum {
	instance := getMPSNNReduceFeatureChannelsSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReduceFeatureChannelsSumFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsSum/initWithDevice:
func NewReduceFeatureChannelsSumWithDevice(device objectivec.IObject) MPSNNReduceFeatureChannelsSum {
	instance := getMPSNNReduceFeatureChannelsSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceFeatureChannelsSumFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithDevice:reduceOperation:
func NewReduceFeatureChannelsSumWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceFeatureChannelsSum {
	instance := getMPSNNReduceFeatureChannelsSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return MPSNNReduceFeatureChannelsSumFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsSum/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (r MPSNNReduceFeatureChannelsSum) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}












// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsSum/weight
func (r MPSNNReduceFeatureChannelsSum) Weight() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("weight"))
	return rv
}
func (r MPSNNReduceFeatureChannelsSum) SetWeight(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setWeight:"), value)
}


















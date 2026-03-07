// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReduceFeatureChannelsMin] class.
var (
	_MPSNNReduceFeatureChannelsMinClass     MPSNNReduceFeatureChannelsMinClass
	_MPSNNReduceFeatureChannelsMinClassOnce sync.Once
)

func getMPSNNReduceFeatureChannelsMinClass() MPSNNReduceFeatureChannelsMinClass {
	_MPSNNReduceFeatureChannelsMinClassOnce.Do(func() {
		_MPSNNReduceFeatureChannelsMinClass = MPSNNReduceFeatureChannelsMinClass{class: objc.GetClass("MPSNNReduceFeatureChannelsMin")}
	})
	return _MPSNNReduceFeatureChannelsMinClass
}

// GetMPSNNReduceFeatureChannelsMinClass returns the class object for MPSNNReduceFeatureChannelsMin.
func GetMPSNNReduceFeatureChannelsMinClass() MPSNNReduceFeatureChannelsMinClass {
	return getMPSNNReduceFeatureChannelsMinClass()
}

type MPSNNReduceFeatureChannelsMinClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceFeatureChannelsMinClass) Alloc() MPSNNReduceFeatureChannelsMin {
	rv := objc.Send[MPSNNReduceFeatureChannelsMin](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReduceFeatureChannelsMin.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsMin
type MPSNNReduceFeatureChannelsMin struct {
	MPSNNReduceUnary
}

// MPSNNReduceFeatureChannelsMinFromID constructs a [MPSNNReduceFeatureChannelsMin] from an objc.ID.
func MPSNNReduceFeatureChannelsMinFromID(id objc.ID) MPSNNReduceFeatureChannelsMin {
	return MPSNNReduceFeatureChannelsMin{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}
// Ensure MPSNNReduceFeatureChannelsMin implements IMPSNNReduceFeatureChannelsMin.
var _ IMPSNNReduceFeatureChannelsMin = MPSNNReduceFeatureChannelsMin{}





// An interface definition for the [MPSNNReduceFeatureChannelsMin] class.
//
// # Methods
//
//   - [IMPSNNReduceFeatureChannelsMin.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsMin
type IMPSNNReduceFeatureChannelsMin interface {
	IMPSNNReduceUnary

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
}





// Init initializes the instance.
func (r MPSNNReduceFeatureChannelsMin) Init() MPSNNReduceFeatureChannelsMin {
	rv := objc.Send[MPSNNReduceFeatureChannelsMin](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceFeatureChannelsMin) Autorelease() MPSNNReduceFeatureChannelsMin {
	rv := objc.Send[MPSNNReduceFeatureChannelsMin](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceFeatureChannelsMin creates a new MPSNNReduceFeatureChannelsMin instance.
func NewMPSNNReduceFeatureChannelsMin() MPSNNReduceFeatureChannelsMin {
	class := getMPSNNReduceFeatureChannelsMinClass()
	rv := objc.Send[MPSNNReduceFeatureChannelsMin](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsMin/initWithCoder:device:
func NewReduceFeatureChannelsMinWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceFeatureChannelsMin {
	instance := getMPSNNReduceFeatureChannelsMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReduceFeatureChannelsMinFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsMin/initWithDevice:
func NewReduceFeatureChannelsMinWithDevice(device objectivec.IObject) MPSNNReduceFeatureChannelsMin {
	instance := getMPSNNReduceFeatureChannelsMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceFeatureChannelsMinFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithDevice:reduceOperation:
func NewReduceFeatureChannelsMinWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceFeatureChannelsMin {
	instance := getMPSNNReduceFeatureChannelsMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return MPSNNReduceFeatureChannelsMinFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsMin/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (r MPSNNReduceFeatureChannelsMin) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}



























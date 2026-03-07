// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReduceFeatureChannelsArgumentMin] class.
var (
	_MPSNNReduceFeatureChannelsArgumentMinClass     MPSNNReduceFeatureChannelsArgumentMinClass
	_MPSNNReduceFeatureChannelsArgumentMinClassOnce sync.Once
)

func getMPSNNReduceFeatureChannelsArgumentMinClass() MPSNNReduceFeatureChannelsArgumentMinClass {
	_MPSNNReduceFeatureChannelsArgumentMinClassOnce.Do(func() {
		_MPSNNReduceFeatureChannelsArgumentMinClass = MPSNNReduceFeatureChannelsArgumentMinClass{class: objc.GetClass("MPSNNReduceFeatureChannelsArgumentMin")}
	})
	return _MPSNNReduceFeatureChannelsArgumentMinClass
}

// GetMPSNNReduceFeatureChannelsArgumentMinClass returns the class object for MPSNNReduceFeatureChannelsArgumentMin.
func GetMPSNNReduceFeatureChannelsArgumentMinClass() MPSNNReduceFeatureChannelsArgumentMinClass {
	return getMPSNNReduceFeatureChannelsArgumentMinClass()
}

type MPSNNReduceFeatureChannelsArgumentMinClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceFeatureChannelsArgumentMinClass) Alloc() MPSNNReduceFeatureChannelsArgumentMin {
	rv := objc.Send[MPSNNReduceFeatureChannelsArgumentMin](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReduceFeatureChannelsArgumentMin.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsArgumentMin
type MPSNNReduceFeatureChannelsArgumentMin struct {
	MPSNNReduceUnary
}

// MPSNNReduceFeatureChannelsArgumentMinFromID constructs a [MPSNNReduceFeatureChannelsArgumentMin] from an objc.ID.
func MPSNNReduceFeatureChannelsArgumentMinFromID(id objc.ID) MPSNNReduceFeatureChannelsArgumentMin {
	return MPSNNReduceFeatureChannelsArgumentMin{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}
// Ensure MPSNNReduceFeatureChannelsArgumentMin implements IMPSNNReduceFeatureChannelsArgumentMin.
var _ IMPSNNReduceFeatureChannelsArgumentMin = MPSNNReduceFeatureChannelsArgumentMin{}





// An interface definition for the [MPSNNReduceFeatureChannelsArgumentMin] class.
//
// # Methods
//
//   - [IMPSNNReduceFeatureChannelsArgumentMin.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsArgumentMin
type IMPSNNReduceFeatureChannelsArgumentMin interface {
	IMPSNNReduceUnary

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
}





// Init initializes the instance.
func (r MPSNNReduceFeatureChannelsArgumentMin) Init() MPSNNReduceFeatureChannelsArgumentMin {
	rv := objc.Send[MPSNNReduceFeatureChannelsArgumentMin](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceFeatureChannelsArgumentMin) Autorelease() MPSNNReduceFeatureChannelsArgumentMin {
	rv := objc.Send[MPSNNReduceFeatureChannelsArgumentMin](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceFeatureChannelsArgumentMin creates a new MPSNNReduceFeatureChannelsArgumentMin instance.
func NewMPSNNReduceFeatureChannelsArgumentMin() MPSNNReduceFeatureChannelsArgumentMin {
	class := getMPSNNReduceFeatureChannelsArgumentMinClass()
	rv := objc.Send[MPSNNReduceFeatureChannelsArgumentMin](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsArgumentMin/initWithCoder:device:
func NewReduceFeatureChannelsArgumentMinWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceFeatureChannelsArgumentMin {
	instance := getMPSNNReduceFeatureChannelsArgumentMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReduceFeatureChannelsArgumentMinFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsArgumentMin/initWithDevice:
func NewReduceFeatureChannelsArgumentMinWithDevice(device objectivec.IObject) MPSNNReduceFeatureChannelsArgumentMin {
	instance := getMPSNNReduceFeatureChannelsArgumentMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceFeatureChannelsArgumentMinFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithDevice:reduceOperation:
func NewReduceFeatureChannelsArgumentMinWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceFeatureChannelsArgumentMin {
	instance := getMPSNNReduceFeatureChannelsArgumentMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return MPSNNReduceFeatureChannelsArgumentMinFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsArgumentMin/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (r MPSNNReduceFeatureChannelsArgumentMin) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}



























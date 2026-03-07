// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReduceFeatureChannelsMean] class.
var (
	_MPSNNReduceFeatureChannelsMeanClass     MPSNNReduceFeatureChannelsMeanClass
	_MPSNNReduceFeatureChannelsMeanClassOnce sync.Once
)

func getMPSNNReduceFeatureChannelsMeanClass() MPSNNReduceFeatureChannelsMeanClass {
	_MPSNNReduceFeatureChannelsMeanClassOnce.Do(func() {
		_MPSNNReduceFeatureChannelsMeanClass = MPSNNReduceFeatureChannelsMeanClass{class: objc.GetClass("MPSNNReduceFeatureChannelsMean")}
	})
	return _MPSNNReduceFeatureChannelsMeanClass
}

// GetMPSNNReduceFeatureChannelsMeanClass returns the class object for MPSNNReduceFeatureChannelsMean.
func GetMPSNNReduceFeatureChannelsMeanClass() MPSNNReduceFeatureChannelsMeanClass {
	return getMPSNNReduceFeatureChannelsMeanClass()
}

type MPSNNReduceFeatureChannelsMeanClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceFeatureChannelsMeanClass) Alloc() MPSNNReduceFeatureChannelsMean {
	rv := objc.Send[MPSNNReduceFeatureChannelsMean](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReduceFeatureChannelsMean.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsMean
type MPSNNReduceFeatureChannelsMean struct {
	MPSNNReduceUnary
}

// MPSNNReduceFeatureChannelsMeanFromID constructs a [MPSNNReduceFeatureChannelsMean] from an objc.ID.
func MPSNNReduceFeatureChannelsMeanFromID(id objc.ID) MPSNNReduceFeatureChannelsMean {
	return MPSNNReduceFeatureChannelsMean{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}
// Ensure MPSNNReduceFeatureChannelsMean implements IMPSNNReduceFeatureChannelsMean.
var _ IMPSNNReduceFeatureChannelsMean = MPSNNReduceFeatureChannelsMean{}





// An interface definition for the [MPSNNReduceFeatureChannelsMean] class.
//
// # Methods
//
//   - [IMPSNNReduceFeatureChannelsMean.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsMean
type IMPSNNReduceFeatureChannelsMean interface {
	IMPSNNReduceUnary

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
}





// Init initializes the instance.
func (r MPSNNReduceFeatureChannelsMean) Init() MPSNNReduceFeatureChannelsMean {
	rv := objc.Send[MPSNNReduceFeatureChannelsMean](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceFeatureChannelsMean) Autorelease() MPSNNReduceFeatureChannelsMean {
	rv := objc.Send[MPSNNReduceFeatureChannelsMean](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceFeatureChannelsMean creates a new MPSNNReduceFeatureChannelsMean instance.
func NewMPSNNReduceFeatureChannelsMean() MPSNNReduceFeatureChannelsMean {
	class := getMPSNNReduceFeatureChannelsMeanClass()
	rv := objc.Send[MPSNNReduceFeatureChannelsMean](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsMean/initWithCoder:device:
func NewReduceFeatureChannelsMeanWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceFeatureChannelsMean {
	instance := getMPSNNReduceFeatureChannelsMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReduceFeatureChannelsMeanFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsMean/initWithDevice:
func NewReduceFeatureChannelsMeanWithDevice(device objectivec.IObject) MPSNNReduceFeatureChannelsMean {
	instance := getMPSNNReduceFeatureChannelsMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceFeatureChannelsMeanFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithDevice:reduceOperation:
func NewReduceFeatureChannelsMeanWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceFeatureChannelsMean {
	instance := getMPSNNReduceFeatureChannelsMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return MPSNNReduceFeatureChannelsMeanFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsMean/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (r MPSNNReduceFeatureChannelsMean) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}



























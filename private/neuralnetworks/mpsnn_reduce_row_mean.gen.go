// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReduceRowMean] class.
var (
	_MPSNNReduceRowMeanClass     MPSNNReduceRowMeanClass
	_MPSNNReduceRowMeanClassOnce sync.Once
)

func getMPSNNReduceRowMeanClass() MPSNNReduceRowMeanClass {
	_MPSNNReduceRowMeanClassOnce.Do(func() {
		_MPSNNReduceRowMeanClass = MPSNNReduceRowMeanClass{class: objc.GetClass("MPSNNReduceRowMean")}
	})
	return _MPSNNReduceRowMeanClass
}

// GetMPSNNReduceRowMeanClass returns the class object for MPSNNReduceRowMean.
func GetMPSNNReduceRowMeanClass() MPSNNReduceRowMeanClass {
	return getMPSNNReduceRowMeanClass()
}

type MPSNNReduceRowMeanClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceRowMeanClass) Alloc() MPSNNReduceRowMean {
	rv := objc.Send[MPSNNReduceRowMean](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReduceRowMean.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowMean
type MPSNNReduceRowMean struct {
	MPSNNReduceUnary
}

// MPSNNReduceRowMeanFromID constructs a [MPSNNReduceRowMean] from an objc.ID.
func MPSNNReduceRowMeanFromID(id objc.ID) MPSNNReduceRowMean {
	return MPSNNReduceRowMean{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}
// Ensure MPSNNReduceRowMean implements IMPSNNReduceRowMean.
var _ IMPSNNReduceRowMean = MPSNNReduceRowMean{}





// An interface definition for the [MPSNNReduceRowMean] class.
//
// # Methods
//
//   - [IMPSNNReduceRowMean.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowMean
type IMPSNNReduceRowMean interface {
	IMPSNNReduceUnary

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
}





// Init initializes the instance.
func (r MPSNNReduceRowMean) Init() MPSNNReduceRowMean {
	rv := objc.Send[MPSNNReduceRowMean](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceRowMean) Autorelease() MPSNNReduceRowMean {
	rv := objc.Send[MPSNNReduceRowMean](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceRowMean creates a new MPSNNReduceRowMean instance.
func NewMPSNNReduceRowMean() MPSNNReduceRowMean {
	class := getMPSNNReduceRowMeanClass()
	rv := objc.Send[MPSNNReduceRowMean](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowMean/initWithCoder:device:
func NewReduceRowMeanWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceRowMean {
	instance := getMPSNNReduceRowMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReduceRowMeanFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowMean/initWithDevice:
func NewReduceRowMeanWithDevice(device objectivec.IObject) MPSNNReduceRowMean {
	instance := getMPSNNReduceRowMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceRowMeanFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithDevice:reduceOperation:
func NewReduceRowMeanWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceRowMean {
	instance := getMPSNNReduceRowMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return MPSNNReduceRowMeanFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowMean/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (r MPSNNReduceRowMean) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}



























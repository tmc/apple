// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReduceColumnMean] class.
var (
	_MPSNNReduceColumnMeanClass     MPSNNReduceColumnMeanClass
	_MPSNNReduceColumnMeanClassOnce sync.Once
)

func getMPSNNReduceColumnMeanClass() MPSNNReduceColumnMeanClass {
	_MPSNNReduceColumnMeanClassOnce.Do(func() {
		_MPSNNReduceColumnMeanClass = MPSNNReduceColumnMeanClass{class: objc.GetClass("MPSNNReduceColumnMean")}
	})
	return _MPSNNReduceColumnMeanClass
}

// GetMPSNNReduceColumnMeanClass returns the class object for MPSNNReduceColumnMean.
func GetMPSNNReduceColumnMeanClass() MPSNNReduceColumnMeanClass {
	return getMPSNNReduceColumnMeanClass()
}

type MPSNNReduceColumnMeanClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceColumnMeanClass) Alloc() MPSNNReduceColumnMean {
	rv := objc.Send[MPSNNReduceColumnMean](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReduceColumnMean.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnMean
type MPSNNReduceColumnMean struct {
	MPSNNReduceUnary
}

// MPSNNReduceColumnMeanFromID constructs a [MPSNNReduceColumnMean] from an objc.ID.
func MPSNNReduceColumnMeanFromID(id objc.ID) MPSNNReduceColumnMean {
	return MPSNNReduceColumnMean{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}
// Ensure MPSNNReduceColumnMean implements IMPSNNReduceColumnMean.
var _ IMPSNNReduceColumnMean = MPSNNReduceColumnMean{}





// An interface definition for the [MPSNNReduceColumnMean] class.
//
// # Methods
//
//   - [IMPSNNReduceColumnMean.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnMean
type IMPSNNReduceColumnMean interface {
	IMPSNNReduceUnary

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
}





// Init initializes the instance.
func (r MPSNNReduceColumnMean) Init() MPSNNReduceColumnMean {
	rv := objc.Send[MPSNNReduceColumnMean](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceColumnMean) Autorelease() MPSNNReduceColumnMean {
	rv := objc.Send[MPSNNReduceColumnMean](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceColumnMean creates a new MPSNNReduceColumnMean instance.
func NewMPSNNReduceColumnMean() MPSNNReduceColumnMean {
	class := getMPSNNReduceColumnMeanClass()
	rv := objc.Send[MPSNNReduceColumnMean](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnMean/initWithCoder:device:
func NewReduceColumnMeanWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceColumnMean {
	instance := getMPSNNReduceColumnMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReduceColumnMeanFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnMean/initWithDevice:
func NewReduceColumnMeanWithDevice(device objectivec.IObject) MPSNNReduceColumnMean {
	instance := getMPSNNReduceColumnMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceColumnMeanFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithDevice:reduceOperation:
func NewReduceColumnMeanWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceColumnMean {
	instance := getMPSNNReduceColumnMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return MPSNNReduceColumnMeanFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnMean/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (r MPSNNReduceColumnMean) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}



























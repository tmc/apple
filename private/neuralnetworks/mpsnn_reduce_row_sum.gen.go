// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReduceRowSum] class.
var (
	_MPSNNReduceRowSumClass     MPSNNReduceRowSumClass
	_MPSNNReduceRowSumClassOnce sync.Once
)

func getMPSNNReduceRowSumClass() MPSNNReduceRowSumClass {
	_MPSNNReduceRowSumClassOnce.Do(func() {
		_MPSNNReduceRowSumClass = MPSNNReduceRowSumClass{class: objc.GetClass("MPSNNReduceRowSum")}
	})
	return _MPSNNReduceRowSumClass
}

// GetMPSNNReduceRowSumClass returns the class object for MPSNNReduceRowSum.
func GetMPSNNReduceRowSumClass() MPSNNReduceRowSumClass {
	return getMPSNNReduceRowSumClass()
}

type MPSNNReduceRowSumClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceRowSumClass) Alloc() MPSNNReduceRowSum {
	rv := objc.Send[MPSNNReduceRowSum](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReduceRowSum.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowSum
type MPSNNReduceRowSum struct {
	MPSNNReduceUnary
}

// MPSNNReduceRowSumFromID constructs a [MPSNNReduceRowSum] from an objc.ID.
func MPSNNReduceRowSumFromID(id objc.ID) MPSNNReduceRowSum {
	return MPSNNReduceRowSum{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}
// Ensure MPSNNReduceRowSum implements IMPSNNReduceRowSum.
var _ IMPSNNReduceRowSum = MPSNNReduceRowSum{}





// An interface definition for the [MPSNNReduceRowSum] class.
//
// # Methods
//
//   - [IMPSNNReduceRowSum.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowSum
type IMPSNNReduceRowSum interface {
	IMPSNNReduceUnary

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
}





// Init initializes the instance.
func (r MPSNNReduceRowSum) Init() MPSNNReduceRowSum {
	rv := objc.Send[MPSNNReduceRowSum](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceRowSum) Autorelease() MPSNNReduceRowSum {
	rv := objc.Send[MPSNNReduceRowSum](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceRowSum creates a new MPSNNReduceRowSum instance.
func NewMPSNNReduceRowSum() MPSNNReduceRowSum {
	class := getMPSNNReduceRowSumClass()
	rv := objc.Send[MPSNNReduceRowSum](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowSum/initWithCoder:device:
func NewReduceRowSumWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceRowSum {
	instance := getMPSNNReduceRowSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReduceRowSumFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowSum/initWithDevice:
func NewReduceRowSumWithDevice(device objectivec.IObject) MPSNNReduceRowSum {
	instance := getMPSNNReduceRowSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceRowSumFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithDevice:reduceOperation:
func NewReduceRowSumWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceRowSum {
	instance := getMPSNNReduceRowSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return MPSNNReduceRowSumFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowSum/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (r MPSNNReduceRowSum) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}



























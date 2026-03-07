// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReduceColumnSum] class.
var (
	_MPSNNReduceColumnSumClass     MPSNNReduceColumnSumClass
	_MPSNNReduceColumnSumClassOnce sync.Once
)

func getMPSNNReduceColumnSumClass() MPSNNReduceColumnSumClass {
	_MPSNNReduceColumnSumClassOnce.Do(func() {
		_MPSNNReduceColumnSumClass = MPSNNReduceColumnSumClass{class: objc.GetClass("MPSNNReduceColumnSum")}
	})
	return _MPSNNReduceColumnSumClass
}

// GetMPSNNReduceColumnSumClass returns the class object for MPSNNReduceColumnSum.
func GetMPSNNReduceColumnSumClass() MPSNNReduceColumnSumClass {
	return getMPSNNReduceColumnSumClass()
}

type MPSNNReduceColumnSumClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceColumnSumClass) Alloc() MPSNNReduceColumnSum {
	rv := objc.Send[MPSNNReduceColumnSum](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReduceColumnSum.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnSum
type MPSNNReduceColumnSum struct {
	MPSNNReduceUnary
}

// MPSNNReduceColumnSumFromID constructs a [MPSNNReduceColumnSum] from an objc.ID.
func MPSNNReduceColumnSumFromID(id objc.ID) MPSNNReduceColumnSum {
	return MPSNNReduceColumnSum{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}
// Ensure MPSNNReduceColumnSum implements IMPSNNReduceColumnSum.
var _ IMPSNNReduceColumnSum = MPSNNReduceColumnSum{}





// An interface definition for the [MPSNNReduceColumnSum] class.
//
// # Methods
//
//   - [IMPSNNReduceColumnSum.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnSum
type IMPSNNReduceColumnSum interface {
	IMPSNNReduceUnary

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
}





// Init initializes the instance.
func (r MPSNNReduceColumnSum) Init() MPSNNReduceColumnSum {
	rv := objc.Send[MPSNNReduceColumnSum](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceColumnSum) Autorelease() MPSNNReduceColumnSum {
	rv := objc.Send[MPSNNReduceColumnSum](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceColumnSum creates a new MPSNNReduceColumnSum instance.
func NewMPSNNReduceColumnSum() MPSNNReduceColumnSum {
	class := getMPSNNReduceColumnSumClass()
	rv := objc.Send[MPSNNReduceColumnSum](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnSum/initWithCoder:device:
func NewReduceColumnSumWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceColumnSum {
	instance := getMPSNNReduceColumnSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReduceColumnSumFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnSum/initWithDevice:
func NewReduceColumnSumWithDevice(device objectivec.IObject) MPSNNReduceColumnSum {
	instance := getMPSNNReduceColumnSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceColumnSumFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithDevice:reduceOperation:
func NewReduceColumnSumWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceColumnSum {
	instance := getMPSNNReduceColumnSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return MPSNNReduceColumnSumFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnSum/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (r MPSNNReduceColumnSum) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}



























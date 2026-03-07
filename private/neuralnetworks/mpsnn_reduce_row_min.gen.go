// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReduceRowMin] class.
var (
	_MPSNNReduceRowMinClass     MPSNNReduceRowMinClass
	_MPSNNReduceRowMinClassOnce sync.Once
)

func getMPSNNReduceRowMinClass() MPSNNReduceRowMinClass {
	_MPSNNReduceRowMinClassOnce.Do(func() {
		_MPSNNReduceRowMinClass = MPSNNReduceRowMinClass{class: objc.GetClass("MPSNNReduceRowMin")}
	})
	return _MPSNNReduceRowMinClass
}

// GetMPSNNReduceRowMinClass returns the class object for MPSNNReduceRowMin.
func GetMPSNNReduceRowMinClass() MPSNNReduceRowMinClass {
	return getMPSNNReduceRowMinClass()
}

type MPSNNReduceRowMinClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceRowMinClass) Alloc() MPSNNReduceRowMin {
	rv := objc.Send[MPSNNReduceRowMin](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReduceRowMin.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowMin
type MPSNNReduceRowMin struct {
	MPSNNReduceUnary
}

// MPSNNReduceRowMinFromID constructs a [MPSNNReduceRowMin] from an objc.ID.
func MPSNNReduceRowMinFromID(id objc.ID) MPSNNReduceRowMin {
	return MPSNNReduceRowMin{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}
// Ensure MPSNNReduceRowMin implements IMPSNNReduceRowMin.
var _ IMPSNNReduceRowMin = MPSNNReduceRowMin{}





// An interface definition for the [MPSNNReduceRowMin] class.
//
// # Methods
//
//   - [IMPSNNReduceRowMin.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowMin
type IMPSNNReduceRowMin interface {
	IMPSNNReduceUnary

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
}





// Init initializes the instance.
func (r MPSNNReduceRowMin) Init() MPSNNReduceRowMin {
	rv := objc.Send[MPSNNReduceRowMin](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceRowMin) Autorelease() MPSNNReduceRowMin {
	rv := objc.Send[MPSNNReduceRowMin](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceRowMin creates a new MPSNNReduceRowMin instance.
func NewMPSNNReduceRowMin() MPSNNReduceRowMin {
	class := getMPSNNReduceRowMinClass()
	rv := objc.Send[MPSNNReduceRowMin](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowMin/initWithCoder:device:
func NewReduceRowMinWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceRowMin {
	instance := getMPSNNReduceRowMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReduceRowMinFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowMin/initWithDevice:
func NewReduceRowMinWithDevice(device objectivec.IObject) MPSNNReduceRowMin {
	instance := getMPSNNReduceRowMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceRowMinFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithDevice:reduceOperation:
func NewReduceRowMinWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceRowMin {
	instance := getMPSNNReduceRowMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return MPSNNReduceRowMinFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceRowMin/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (r MPSNNReduceRowMin) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}



























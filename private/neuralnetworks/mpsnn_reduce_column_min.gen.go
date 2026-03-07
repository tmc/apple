// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReduceColumnMin] class.
var (
	_MPSNNReduceColumnMinClass     MPSNNReduceColumnMinClass
	_MPSNNReduceColumnMinClassOnce sync.Once
)

func getMPSNNReduceColumnMinClass() MPSNNReduceColumnMinClass {
	_MPSNNReduceColumnMinClassOnce.Do(func() {
		_MPSNNReduceColumnMinClass = MPSNNReduceColumnMinClass{class: objc.GetClass("MPSNNReduceColumnMin")}
	})
	return _MPSNNReduceColumnMinClass
}

// GetMPSNNReduceColumnMinClass returns the class object for MPSNNReduceColumnMin.
func GetMPSNNReduceColumnMinClass() MPSNNReduceColumnMinClass {
	return getMPSNNReduceColumnMinClass()
}

type MPSNNReduceColumnMinClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceColumnMinClass) Alloc() MPSNNReduceColumnMin {
	rv := objc.Send[MPSNNReduceColumnMin](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReduceColumnMin.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnMin
type MPSNNReduceColumnMin struct {
	MPSNNReduceUnary
}

// MPSNNReduceColumnMinFromID constructs a [MPSNNReduceColumnMin] from an objc.ID.
func MPSNNReduceColumnMinFromID(id objc.ID) MPSNNReduceColumnMin {
	return MPSNNReduceColumnMin{MPSNNReduceUnary: MPSNNReduceUnaryFromID(id)}
}
// Ensure MPSNNReduceColumnMin implements IMPSNNReduceColumnMin.
var _ IMPSNNReduceColumnMin = MPSNNReduceColumnMin{}





// An interface definition for the [MPSNNReduceColumnMin] class.
//
// # Methods
//
//   - [IMPSNNReduceColumnMin.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnMin
type IMPSNNReduceColumnMin interface {
	IMPSNNReduceUnary

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
}





// Init initializes the instance.
func (r MPSNNReduceColumnMin) Init() MPSNNReduceColumnMin {
	rv := objc.Send[MPSNNReduceColumnMin](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceColumnMin) Autorelease() MPSNNReduceColumnMin {
	rv := objc.Send[MPSNNReduceColumnMin](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceColumnMin creates a new MPSNNReduceColumnMin instance.
func NewMPSNNReduceColumnMin() MPSNNReduceColumnMin {
	class := getMPSNNReduceColumnMinClass()
	rv := objc.Send[MPSNNReduceColumnMin](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnMin/initWithCoder:device:
func NewReduceColumnMinWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceColumnMin {
	instance := getMPSNNReduceColumnMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReduceColumnMinFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnMin/initWithDevice:
func NewReduceColumnMinWithDevice(device objectivec.IObject) MPSNNReduceColumnMin {
	instance := getMPSNNReduceColumnMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceColumnMinFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceUnary/initWithDevice:reduceOperation:
func NewReduceColumnMinWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceColumnMin {
	instance := getMPSNNReduceColumnMinClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return MPSNNReduceColumnMinFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceColumnMin/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (r MPSNNReduceColumnMin) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}



























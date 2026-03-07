// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNTensorFlowPoolingPadding] class.
var (
	_MPSNNTensorFlowPoolingPaddingClass     MPSNNTensorFlowPoolingPaddingClass
	_MPSNNTensorFlowPoolingPaddingClassOnce sync.Once
)

func getMPSNNTensorFlowPoolingPaddingClass() MPSNNTensorFlowPoolingPaddingClass {
	_MPSNNTensorFlowPoolingPaddingClassOnce.Do(func() {
		_MPSNNTensorFlowPoolingPaddingClass = MPSNNTensorFlowPoolingPaddingClass{class: objc.GetClass("MPSNNTensorFlowPoolingPadding")}
	})
	return _MPSNNTensorFlowPoolingPaddingClass
}

// GetMPSNNTensorFlowPoolingPaddingClass returns the class object for MPSNNTensorFlowPoolingPadding.
func GetMPSNNTensorFlowPoolingPaddingClass() MPSNNTensorFlowPoolingPaddingClass {
	return getMPSNNTensorFlowPoolingPaddingClass()
}

type MPSNNTensorFlowPoolingPaddingClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNTensorFlowPoolingPaddingClass) Alloc() MPSNNTensorFlowPoolingPadding {
	rv := objc.Send[MPSNNTensorFlowPoolingPadding](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNTensorFlowPoolingPadding.DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNTensorFlowPoolingPadding
type MPSNNTensorFlowPoolingPadding struct {
	MPSNNDefaultPadding
}

// MPSNNTensorFlowPoolingPaddingFromID constructs a [MPSNNTensorFlowPoolingPadding] from an objc.ID.
func MPSNNTensorFlowPoolingPaddingFromID(id objc.ID) MPSNNTensorFlowPoolingPadding {
	return MPSNNTensorFlowPoolingPadding{MPSNNDefaultPadding: MPSNNDefaultPaddingFromID(id)}
}
// Ensure MPSNNTensorFlowPoolingPadding implements IMPSNNTensorFlowPoolingPadding.
var _ IMPSNNTensorFlowPoolingPadding = MPSNNTensorFlowPoolingPadding{}





// An interface definition for the [MPSNNTensorFlowPoolingPadding] class.
//
// # Methods
//
//   - [IMPSNNTensorFlowPoolingPadding.DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNTensorFlowPoolingPadding
type IMPSNNTensorFlowPoolingPadding interface {
	IMPSNNDefaultPadding

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(images objectivec.IObject, states objectivec.IObject, kernel objectivec.IObject, descriptor objectivec.IObject) objectivec.IObject
}





// Init initializes the instance.
func (t MPSNNTensorFlowPoolingPadding) Init() MPSNNTensorFlowPoolingPadding {
	rv := objc.Send[MPSNNTensorFlowPoolingPadding](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MPSNNTensorFlowPoolingPadding) Autorelease() MPSNNTensorFlowPoolingPadding {
	rv := objc.Send[MPSNNTensorFlowPoolingPadding](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNTensorFlowPoolingPadding creates a new MPSNNTensorFlowPoolingPadding instance.
func NewMPSNNTensorFlowPoolingPadding() MPSNNTensorFlowPoolingPadding {
	class := getMPSNNTensorFlowPoolingPaddingClass()
	rv := objc.Send[MPSNNTensorFlowPoolingPadding](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNTensorFlowPoolingPadding/initWithCoder:
func NewTensorFlowPoolingPaddingWithCoder(coder objectivec.IObject) MPSNNTensorFlowPoolingPadding {
	instance := getMPSNNTensorFlowPoolingPaddingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MPSNNTensorFlowPoolingPaddingFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding/initWithPaddingMethod:
func NewTensorFlowPoolingPaddingWithPaddingMethod(method uint64) MPSNNTensorFlowPoolingPadding {
	instance := getMPSNNTensorFlowPoolingPaddingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPaddingMethod:"), method)
	return MPSNNTensorFlowPoolingPaddingFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNTensorFlowPoolingPadding/destinationImageDescriptorForSourceImages:sourceStates:forKernel:suggestedDescriptor:
func (t MPSNNTensorFlowPoolingPadding) DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(images objectivec.IObject, states objectivec.IObject, kernel objectivec.IObject, descriptor objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:forKernel:suggestedDescriptor:"), images, states, kernel, descriptor)
	return objectivec.Object{ID: rv}
}



























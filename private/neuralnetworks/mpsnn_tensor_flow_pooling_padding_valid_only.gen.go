// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNTensorFlowPoolingPaddingValidOnly] class.
var (
	_MPSNNTensorFlowPoolingPaddingValidOnlyClass     MPSNNTensorFlowPoolingPaddingValidOnlyClass
	_MPSNNTensorFlowPoolingPaddingValidOnlyClassOnce sync.Once
)

func getMPSNNTensorFlowPoolingPaddingValidOnlyClass() MPSNNTensorFlowPoolingPaddingValidOnlyClass {
	_MPSNNTensorFlowPoolingPaddingValidOnlyClassOnce.Do(func() {
		_MPSNNTensorFlowPoolingPaddingValidOnlyClass = MPSNNTensorFlowPoolingPaddingValidOnlyClass{class: objc.GetClass("MPSNNTensorFlowPoolingPaddingValidOnly")}
	})
	return _MPSNNTensorFlowPoolingPaddingValidOnlyClass
}

// GetMPSNNTensorFlowPoolingPaddingValidOnlyClass returns the class object for MPSNNTensorFlowPoolingPaddingValidOnly.
func GetMPSNNTensorFlowPoolingPaddingValidOnlyClass() MPSNNTensorFlowPoolingPaddingValidOnlyClass {
	return getMPSNNTensorFlowPoolingPaddingValidOnlyClass()
}

type MPSNNTensorFlowPoolingPaddingValidOnlyClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNTensorFlowPoolingPaddingValidOnlyClass) Alloc() MPSNNTensorFlowPoolingPaddingValidOnly {
	rv := objc.Send[MPSNNTensorFlowPoolingPaddingValidOnly](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNTensorFlowPoolingPaddingValidOnly.DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNTensorFlowPoolingPaddingValidOnly
type MPSNNTensorFlowPoolingPaddingValidOnly struct {
	MPSNNDefaultPadding
}

// MPSNNTensorFlowPoolingPaddingValidOnlyFromID constructs a [MPSNNTensorFlowPoolingPaddingValidOnly] from an objc.ID.
func MPSNNTensorFlowPoolingPaddingValidOnlyFromID(id objc.ID) MPSNNTensorFlowPoolingPaddingValidOnly {
	return MPSNNTensorFlowPoolingPaddingValidOnly{MPSNNDefaultPadding: MPSNNDefaultPaddingFromID(id)}
}
// Ensure MPSNNTensorFlowPoolingPaddingValidOnly implements IMPSNNTensorFlowPoolingPaddingValidOnly.
var _ IMPSNNTensorFlowPoolingPaddingValidOnly = MPSNNTensorFlowPoolingPaddingValidOnly{}





// An interface definition for the [MPSNNTensorFlowPoolingPaddingValidOnly] class.
//
// # Methods
//
//   - [IMPSNNTensorFlowPoolingPaddingValidOnly.DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNTensorFlowPoolingPaddingValidOnly
type IMPSNNTensorFlowPoolingPaddingValidOnly interface {
	IMPSNNDefaultPadding

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(images objectivec.IObject, states objectivec.IObject, kernel objectivec.IObject, descriptor objectivec.IObject) objectivec.IObject
}





// Init initializes the instance.
func (t MPSNNTensorFlowPoolingPaddingValidOnly) Init() MPSNNTensorFlowPoolingPaddingValidOnly {
	rv := objc.Send[MPSNNTensorFlowPoolingPaddingValidOnly](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MPSNNTensorFlowPoolingPaddingValidOnly) Autorelease() MPSNNTensorFlowPoolingPaddingValidOnly {
	rv := objc.Send[MPSNNTensorFlowPoolingPaddingValidOnly](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNTensorFlowPoolingPaddingValidOnly creates a new MPSNNTensorFlowPoolingPaddingValidOnly instance.
func NewMPSNNTensorFlowPoolingPaddingValidOnly() MPSNNTensorFlowPoolingPaddingValidOnly {
	class := getMPSNNTensorFlowPoolingPaddingValidOnlyClass()
	rv := objc.Send[MPSNNTensorFlowPoolingPaddingValidOnly](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNTensorFlowPoolingPaddingValidOnly/initWithCoder:
func NewTensorFlowPoolingPaddingValidOnlyWithCoder(coder objectivec.IObject) MPSNNTensorFlowPoolingPaddingValidOnly {
	instance := getMPSNNTensorFlowPoolingPaddingValidOnlyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MPSNNTensorFlowPoolingPaddingValidOnlyFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNDefaultPadding/initWithPaddingMethod:
func NewTensorFlowPoolingPaddingValidOnlyWithPaddingMethod(method uint64) MPSNNTensorFlowPoolingPaddingValidOnly {
	instance := getMPSNNTensorFlowPoolingPaddingValidOnlyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPaddingMethod:"), method)
	return MPSNNTensorFlowPoolingPaddingValidOnlyFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNTensorFlowPoolingPaddingValidOnly/destinationImageDescriptorForSourceImages:sourceStates:forKernel:suggestedDescriptor:
func (t MPSNNTensorFlowPoolingPaddingValidOnly) DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(images objectivec.IObject, states objectivec.IObject, kernel objectivec.IObject, descriptor objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:forKernel:suggestedDescriptor:"), images, states, kernel, descriptor)
	return objectivec.Object{ID: rv}
}



























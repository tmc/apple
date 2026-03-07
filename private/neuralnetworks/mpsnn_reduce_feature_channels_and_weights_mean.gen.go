// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReduceFeatureChannelsAndWeightsMean] class.
var (
	_MPSNNReduceFeatureChannelsAndWeightsMeanClass     MPSNNReduceFeatureChannelsAndWeightsMeanClass
	_MPSNNReduceFeatureChannelsAndWeightsMeanClassOnce sync.Once
)

func getMPSNNReduceFeatureChannelsAndWeightsMeanClass() MPSNNReduceFeatureChannelsAndWeightsMeanClass {
	_MPSNNReduceFeatureChannelsAndWeightsMeanClassOnce.Do(func() {
		_MPSNNReduceFeatureChannelsAndWeightsMeanClass = MPSNNReduceFeatureChannelsAndWeightsMeanClass{class: objc.GetClass("MPSNNReduceFeatureChannelsAndWeightsMean")}
	})
	return _MPSNNReduceFeatureChannelsAndWeightsMeanClass
}

// GetMPSNNReduceFeatureChannelsAndWeightsMeanClass returns the class object for MPSNNReduceFeatureChannelsAndWeightsMean.
func GetMPSNNReduceFeatureChannelsAndWeightsMeanClass() MPSNNReduceFeatureChannelsAndWeightsMeanClass {
	return getMPSNNReduceFeatureChannelsAndWeightsMeanClass()
}

type MPSNNReduceFeatureChannelsAndWeightsMeanClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceFeatureChannelsAndWeightsMeanClass) Alloc() MPSNNReduceFeatureChannelsAndWeightsMean {
	rv := objc.Send[MPSNNReduceFeatureChannelsAndWeightsMean](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReduceFeatureChannelsAndWeightsMean.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsAndWeightsMean
type MPSNNReduceFeatureChannelsAndWeightsMean struct {
	MPSNNReduceBinary
}

// MPSNNReduceFeatureChannelsAndWeightsMeanFromID constructs a [MPSNNReduceFeatureChannelsAndWeightsMean] from an objc.ID.
func MPSNNReduceFeatureChannelsAndWeightsMeanFromID(id objc.ID) MPSNNReduceFeatureChannelsAndWeightsMean {
	return MPSNNReduceFeatureChannelsAndWeightsMean{MPSNNReduceBinary: MPSNNReduceBinaryFromID(id)}
}
// Ensure MPSNNReduceFeatureChannelsAndWeightsMean implements IMPSNNReduceFeatureChannelsAndWeightsMean.
var _ IMPSNNReduceFeatureChannelsAndWeightsMean = MPSNNReduceFeatureChannelsAndWeightsMean{}





// An interface definition for the [MPSNNReduceFeatureChannelsAndWeightsMean] class.
//
// # Methods
//
//   - [IMPSNNReduceFeatureChannelsAndWeightsMean.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsAndWeightsMean
type IMPSNNReduceFeatureChannelsAndWeightsMean interface {
	IMPSNNReduceBinary

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject, offset2 objectivec.IObject, offset3 objectivec.IObject) objectivec.IObject
}





// Init initializes the instance.
func (r MPSNNReduceFeatureChannelsAndWeightsMean) Init() MPSNNReduceFeatureChannelsAndWeightsMean {
	rv := objc.Send[MPSNNReduceFeatureChannelsAndWeightsMean](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceFeatureChannelsAndWeightsMean) Autorelease() MPSNNReduceFeatureChannelsAndWeightsMean {
	rv := objc.Send[MPSNNReduceFeatureChannelsAndWeightsMean](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceFeatureChannelsAndWeightsMean creates a new MPSNNReduceFeatureChannelsAndWeightsMean instance.
func NewMPSNNReduceFeatureChannelsAndWeightsMean() MPSNNReduceFeatureChannelsAndWeightsMean {
	class := getMPSNNReduceFeatureChannelsAndWeightsMeanClass()
	rv := objc.Send[MPSNNReduceFeatureChannelsAndWeightsMean](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsAndWeightsMean/initWithCoder:device:
func NewReduceFeatureChannelsAndWeightsMeanWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceFeatureChannelsAndWeightsMean {
	instance := getMPSNNReduceFeatureChannelsAndWeightsMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReduceFeatureChannelsAndWeightsMeanFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsAndWeightsMean/initWithDevice:
func NewReduceFeatureChannelsAndWeightsMeanWithDevice(device objectivec.IObject) MPSNNReduceFeatureChannelsAndWeightsMean {
	instance := getMPSNNReduceFeatureChannelsAndWeightsMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceFeatureChannelsAndWeightsMeanFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceBinary/initWithDevice:reduceOperation:
func NewReduceFeatureChannelsAndWeightsMeanWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceFeatureChannelsAndWeightsMean {
	instance := getMPSNNReduceFeatureChannelsAndWeightsMeanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return MPSNNReduceFeatureChannelsAndWeightsMeanFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsAndWeightsMean/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:primaryOffset:secondaryOffset:kernelOffset:
func (r MPSNNReduceFeatureChannelsAndWeightsMean) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject, offset2 objectivec.IObject, offset3 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:primaryOffset:secondaryOffset:kernelOffset:"), images, states, method, offset, offset2, offset3)
	return objectivec.Object{ID: rv}
}



























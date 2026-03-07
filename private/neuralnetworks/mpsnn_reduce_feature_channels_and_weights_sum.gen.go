// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReduceFeatureChannelsAndWeightsSum] class.
var (
	_MPSNNReduceFeatureChannelsAndWeightsSumClass     MPSNNReduceFeatureChannelsAndWeightsSumClass
	_MPSNNReduceFeatureChannelsAndWeightsSumClassOnce sync.Once
)

func getMPSNNReduceFeatureChannelsAndWeightsSumClass() MPSNNReduceFeatureChannelsAndWeightsSumClass {
	_MPSNNReduceFeatureChannelsAndWeightsSumClassOnce.Do(func() {
		_MPSNNReduceFeatureChannelsAndWeightsSumClass = MPSNNReduceFeatureChannelsAndWeightsSumClass{class: objc.GetClass("MPSNNReduceFeatureChannelsAndWeightsSum")}
	})
	return _MPSNNReduceFeatureChannelsAndWeightsSumClass
}

// GetMPSNNReduceFeatureChannelsAndWeightsSumClass returns the class object for MPSNNReduceFeatureChannelsAndWeightsSum.
func GetMPSNNReduceFeatureChannelsAndWeightsSumClass() MPSNNReduceFeatureChannelsAndWeightsSumClass {
	return getMPSNNReduceFeatureChannelsAndWeightsSumClass()
}

type MPSNNReduceFeatureChannelsAndWeightsSumClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReduceFeatureChannelsAndWeightsSumClass) Alloc() MPSNNReduceFeatureChannelsAndWeightsSum {
	rv := objc.Send[MPSNNReduceFeatureChannelsAndWeightsSum](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReduceFeatureChannelsAndWeightsSum.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset]
//   - [MPSNNReduceFeatureChannelsAndWeightsSum.DoWeightedSumByNonZeroWeights]
//   - [MPSNNReduceFeatureChannelsAndWeightsSum.InitWithDeviceDoWeightedSumByNonZeroWeights]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsAndWeightsSum
type MPSNNReduceFeatureChannelsAndWeightsSum struct {
	MPSNNReduceBinary
}

// MPSNNReduceFeatureChannelsAndWeightsSumFromID constructs a [MPSNNReduceFeatureChannelsAndWeightsSum] from an objc.ID.
func MPSNNReduceFeatureChannelsAndWeightsSumFromID(id objc.ID) MPSNNReduceFeatureChannelsAndWeightsSum {
	return MPSNNReduceFeatureChannelsAndWeightsSum{MPSNNReduceBinary: MPSNNReduceBinaryFromID(id)}
}
// Ensure MPSNNReduceFeatureChannelsAndWeightsSum implements IMPSNNReduceFeatureChannelsAndWeightsSum.
var _ IMPSNNReduceFeatureChannelsAndWeightsSum = MPSNNReduceFeatureChannelsAndWeightsSum{}





// An interface definition for the [MPSNNReduceFeatureChannelsAndWeightsSum] class.
//
// # Methods
//
//   - [IMPSNNReduceFeatureChannelsAndWeightsSum.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset]
//   - [IMPSNNReduceFeatureChannelsAndWeightsSum.DoWeightedSumByNonZeroWeights]
//   - [IMPSNNReduceFeatureChannelsAndWeightsSum.InitWithDeviceDoWeightedSumByNonZeroWeights]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsAndWeightsSum
type IMPSNNReduceFeatureChannelsAndWeightsSum interface {
	IMPSNNReduceBinary

	// Topic: Methods

	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject, offset2 objectivec.IObject, offset3 objectivec.IObject) objectivec.IObject
	DoWeightedSumByNonZeroWeights() bool
	InitWithDeviceDoWeightedSumByNonZeroWeights(device objectivec.IObject, weights bool) MPSNNReduceFeatureChannelsAndWeightsSum
}





// Init initializes the instance.
func (r MPSNNReduceFeatureChannelsAndWeightsSum) Init() MPSNNReduceFeatureChannelsAndWeightsSum {
	rv := objc.Send[MPSNNReduceFeatureChannelsAndWeightsSum](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReduceFeatureChannelsAndWeightsSum) Autorelease() MPSNNReduceFeatureChannelsAndWeightsSum {
	rv := objc.Send[MPSNNReduceFeatureChannelsAndWeightsSum](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReduceFeatureChannelsAndWeightsSum creates a new MPSNNReduceFeatureChannelsAndWeightsSum instance.
func NewMPSNNReduceFeatureChannelsAndWeightsSum() MPSNNReduceFeatureChannelsAndWeightsSum {
	class := getMPSNNReduceFeatureChannelsAndWeightsSumClass()
	rv := objc.Send[MPSNNReduceFeatureChannelsAndWeightsSum](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsAndWeightsSum/initWithCoder:device:
func NewReduceFeatureChannelsAndWeightsSumWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReduceFeatureChannelsAndWeightsSum {
	instance := getMPSNNReduceFeatureChannelsAndWeightsSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReduceFeatureChannelsAndWeightsSumFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsAndWeightsSum/initWithDevice:
func NewReduceFeatureChannelsAndWeightsSumWithDevice(device objectivec.IObject) MPSNNReduceFeatureChannelsAndWeightsSum {
	instance := getMPSNNReduceFeatureChannelsAndWeightsSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReduceFeatureChannelsAndWeightsSumFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsAndWeightsSum/initWithDevice:doWeightedSumByNonZeroWeights:
func NewReduceFeatureChannelsAndWeightsSumWithDeviceDoWeightedSumByNonZeroWeights(device objectivec.IObject, weights bool) MPSNNReduceFeatureChannelsAndWeightsSum {
	instance := getMPSNNReduceFeatureChannelsAndWeightsSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:doWeightedSumByNonZeroWeights:"), device, weights)
	return MPSNNReduceFeatureChannelsAndWeightsSumFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceBinary/initWithDevice:reduceOperation:
func NewReduceFeatureChannelsAndWeightsSumWithDeviceReduceOperation(device objectivec.IObject, operation int) MPSNNReduceFeatureChannelsAndWeightsSum {
	instance := getMPSNNReduceFeatureChannelsAndWeightsSumClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:reduceOperation:"), device, operation)
	return MPSNNReduceFeatureChannelsAndWeightsSumFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsAndWeightsSum/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:primaryOffset:secondaryOffset:kernelOffset:
func (r MPSNNReduceFeatureChannelsAndWeightsSum) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodPrimaryOffsetSecondaryOffsetKernelOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject, offset2 objectivec.IObject, offset3 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:primaryOffset:secondaryOffset:kernelOffset:"), images, states, method, offset, offset2, offset3)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsAndWeightsSum/initWithDevice:doWeightedSumByNonZeroWeights:
func (r MPSNNReduceFeatureChannelsAndWeightsSum) InitWithDeviceDoWeightedSumByNonZeroWeights(device objectivec.IObject, weights bool) MPSNNReduceFeatureChannelsAndWeightsSum {
	rv := objc.Send[MPSNNReduceFeatureChannelsAndWeightsSum](r.ID, objc.Sel("initWithDevice:doWeightedSumByNonZeroWeights:"), device, weights)
	return rv
}












// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReduceFeatureChannelsAndWeightsSum/doWeightedSumByNonZeroWeights
func (r MPSNNReduceFeatureChannelsAndWeightsSum) DoWeightedSumByNonZeroWeights() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("doWeightedSumByNonZeroWeights"))
	return rv
}


















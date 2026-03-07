// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNReshape] class.
var (
	_MPSNNReshapeClass     MPSNNReshapeClass
	_MPSNNReshapeClassOnce sync.Once
)

func getMPSNNReshapeClass() MPSNNReshapeClass {
	_MPSNNReshapeClassOnce.Do(func() {
		_MPSNNReshapeClass = MPSNNReshapeClass{class: objc.GetClass("MPSNNReshape")}
	})
	return _MPSNNReshapeClass
}

// GetMPSNNReshapeClass returns the class object for MPSNNReshape.
func GetMPSNNReshapeClass() MPSNNReshapeClass {
	return getMPSNNReshapeClass()
}

type MPSNNReshapeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNReshapeClass) Alloc() MPSNNReshape {
	rv := objc.Send[MPSNNReshape](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNReshape.CopyWithZoneDevice]
//   - [MPSNNReshape.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//   - [MPSNNReshape.EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels]
//   - [MPSNNReshape.EncodeBatchToCommandBufferSourceImagesReshapedWidthReshapedHeightReshapedFeatureChannels]
//   - [MPSNNReshape.EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels]
//   - [MPSNNReshape.EncodeToCommandBufferSourceImageReshapedWidthReshapedHeightReshapedFeatureChannels]
//   - [MPSNNReshape.InitWithCoderDevice]
//   - [MPSNNReshape.InitWithDevice]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshape
type MPSNNReshape struct {
	MPSCNNKernel
}

// MPSNNReshapeFromID constructs a [MPSNNReshape] from an objc.ID.
func MPSNNReshapeFromID(id objc.ID) MPSNNReshape {
	return MPSNNReshape{MPSCNNKernel: MPSCNNKernelFromID(id)}
}
// Ensure MPSNNReshape implements IMPSNNReshape.
var _ IMPSNNReshape = MPSNNReshape{}





// An interface definition for the [MPSNNReshape] class.
//
// # Methods
//
//   - [IMPSNNReshape.CopyWithZoneDevice]
//   - [IMPSNNReshape.DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset]
//   - [IMPSNNReshape.EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels]
//   - [IMPSNNReshape.EncodeBatchToCommandBufferSourceImagesReshapedWidthReshapedHeightReshapedFeatureChannels]
//   - [IMPSNNReshape.EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels]
//   - [IMPSNNReshape.EncodeToCommandBufferSourceImageReshapedWidthReshapedHeightReshapedFeatureChannels]
//   - [IMPSNNReshape.InitWithCoderDevice]
//   - [IMPSNNReshape.InitWithDevice]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshape
type IMPSNNReshape interface {
	IMPSCNNKernel

	// Topic: Methods

	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject
	EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels(buffer objectivec.IObject, images objectivec.IObject, states []objectivec.IObject, temporary bool, width uint64, height uint64, channels uint64) objectivec.IObject
	EncodeBatchToCommandBufferSourceImagesReshapedWidthReshapedHeightReshapedFeatureChannels(buffer objectivec.IObject, images objectivec.IObject, width uint64, height uint64, channels uint64) objectivec.IObject
	EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels(buffer objectivec.IObject, image objectivec.IObject, state []objectivec.IObject, temporary bool, width uint64, height uint64, channels uint64) objectivec.IObject
	EncodeToCommandBufferSourceImageReshapedWidthReshapedHeightReshapedFeatureChannels(buffer objectivec.IObject, image objectivec.IObject, width uint64, height uint64, channels uint64) objectivec.IObject
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReshape
	InitWithDevice(device objectivec.IObject) MPSNNReshape
}





// Init initializes the instance.
func (r MPSNNReshape) Init() MPSNNReshape {
	rv := objc.Send[MPSNNReshape](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSNNReshape) Autorelease() MPSNNReshape {
	rv := objc.Send[MPSNNReshape](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNReshape creates a new MPSNNReshape instance.
func NewMPSNNReshape() MPSNNReshape {
	class := getMPSNNReshapeClass()
	rv := objc.Send[MPSNNReshape](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshape/initWithCoder:device:
func NewReshapeWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReshape {
	instance := getMPSNNReshapeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNReshapeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshape/initWithDevice:
func NewReshapeWithDevice(device objectivec.IObject) MPSNNReshape {
	instance := getMPSNNReshapeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNReshapeFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshape/copyWithZone:device:
func (r MPSNNReshape) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshape/destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:
func (r MPSNNReshape) DestinationImageDescriptorForSourceImagesSourceStatesPaddingMethodSourceOffset(images objectivec.IObject, states objectivec.IObject, method uint64, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:paddingMethod:sourceOffset:"), images, states, method, offset)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshape/encodeBatchToCommandBuffer:sourceImages:destinationStates:destinationStateIsTemporary:reshapedWidth:reshapedHeight:reshapedFeatureChannels:
func (r MPSNNReshape) EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels(buffer objectivec.IObject, images objectivec.IObject, states []objectivec.IObject, temporary bool, width uint64, height uint64, channels uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationStates:destinationStateIsTemporary:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), buffer, images, objectivec.IObjectSliceToNSArray(states), temporary, width, height, channels)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshape/encodeBatchToCommandBuffer:sourceImages:reshapedWidth:reshapedHeight:reshapedFeatureChannels:
func (r MPSNNReshape) EncodeBatchToCommandBufferSourceImagesReshapedWidthReshapedHeightReshapedFeatureChannels(buffer objectivec.IObject, images objectivec.IObject, width uint64, height uint64, channels uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), buffer, images, width, height, channels)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshape/encodeToCommandBuffer:sourceImage:destinationState:destinationStateIsTemporary:reshapedWidth:reshapedHeight:reshapedFeatureChannels:
func (r MPSNNReshape) EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels(buffer objectivec.IObject, image objectivec.IObject, state []objectivec.IObject, temporary bool, width uint64, height uint64, channels uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("encodeToCommandBuffer:sourceImage:destinationState:destinationStateIsTemporary:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), buffer, image, objectivec.IObjectSliceToNSArray(state), temporary, width, height, channels)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshape/encodeToCommandBuffer:sourceImage:reshapedWidth:reshapedHeight:reshapedFeatureChannels:
func (r MPSNNReshape) EncodeToCommandBufferSourceImageReshapedWidthReshapedHeightReshapedFeatureChannels(buffer objectivec.IObject, image objectivec.IObject, width uint64, height uint64, channels uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("encodeToCommandBuffer:sourceImage:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), buffer, image, width, height, channels)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshape/initWithCoder:device:
func (r MPSNNReshape) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNReshape {
	rv := objc.Send[MPSNNReshape](r.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshape/initWithDevice:
func (r MPSNNReshape) InitWithDevice(device objectivec.IObject) MPSNNReshape {
	rv := objc.Send[MPSNNReshape](r.ID, objc.Sel("initWithDevice:"), device)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNReshape/libraryInfo:
func (_MPSNNReshapeClass MPSNNReshapeClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNReshapeClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}























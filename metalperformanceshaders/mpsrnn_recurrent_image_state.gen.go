// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSRNNRecurrentImageState] class.
var (
	_MPSRNNRecurrentImageStateClass     MPSRNNRecurrentImageStateClass
	_MPSRNNRecurrentImageStateClassOnce sync.Once
)

func getMPSRNNRecurrentImageStateClass() MPSRNNRecurrentImageStateClass {
	_MPSRNNRecurrentImageStateClassOnce.Do(func() {
		_MPSRNNRecurrentImageStateClass = MPSRNNRecurrentImageStateClass{class: objc.GetClass("MPSRNNRecurrentImageState")}
	})
	return _MPSRNNRecurrentImageStateClass
}

// GetMPSRNNRecurrentImageStateClass returns the class object for MPSRNNRecurrentImageState.
func GetMPSRNNRecurrentImageStateClass() MPSRNNRecurrentImageStateClass {
	return getMPSRNNRecurrentImageStateClass()
}

type MPSRNNRecurrentImageStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSRNNRecurrentImageStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSRNNRecurrentImageStateClass) Alloc() MPSRNNRecurrentImageState {
	rv := objc.Send[MPSRNNRecurrentImageState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that holds all the data that’s passed from one sequence iteration
// of the image-based recurrent neural network layer (stack) to the next.
//
// # Instance Methods
//
//   - [MPSRNNRecurrentImageState.GetMemoryCellImageForLayerIndex]
//   - [MPSRNNRecurrentImageState.GetRecurrentOutputImageForLayerIndex]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNRecurrentImageState
type MPSRNNRecurrentImageState struct {
	MPSState
}

// MPSRNNRecurrentImageStateFromID constructs a [MPSRNNRecurrentImageState] from an objc.ID.
//
// A class that holds all the data that’s passed from one sequence iteration
// of the image-based recurrent neural network layer (stack) to the next.
func MPSRNNRecurrentImageStateFromID(id objc.ID) MPSRNNRecurrentImageState {
	return MPSRNNRecurrentImageState{MPSState: MPSStateFromID(id)}
}

// NOTE: MPSRNNRecurrentImageState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSRNNRecurrentImageState] class.
//
// # Instance Methods
//
//   - [IMPSRNNRecurrentImageState.GetMemoryCellImageForLayerIndex]
//   - [IMPSRNNRecurrentImageState.GetRecurrentOutputImageForLayerIndex]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNRecurrentImageState
type IMPSRNNRecurrentImageState interface {
	IMPSState

	// Topic: Instance Methods

	GetMemoryCellImageForLayerIndex(layerIndex uint) IMPSImage
	GetRecurrentOutputImageForLayerIndex(layerIndex uint) IMPSImage
}

// Init initializes the instance.
func (r MPSRNNRecurrentImageState) Init() MPSRNNRecurrentImageState {
	rv := objc.Send[MPSRNNRecurrentImageState](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSRNNRecurrentImageState) Autorelease() MPSRNNRecurrentImageState {
	rv := objc.Send[MPSRNNRecurrentImageState](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSRNNRecurrentImageState creates a new MPSRNNRecurrentImageState instance.
func NewMPSRNNRecurrentImageState() MPSRNNRecurrentImageState {
	class := getMPSRNNRecurrentImageStateClass()
	rv := objc.Send[MPSRNNRecurrentImageState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewRNNRecurrentImageStateWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSRNNRecurrentImageState {
	instance := getMPSRNNRecurrentImageStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSRNNRecurrentImageStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewRNNRecurrentImageStateWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSRNNRecurrentImageState {
	instance := getMPSRNNRecurrentImageStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSRNNRecurrentImageStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewRNNRecurrentImageStateWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSRNNRecurrentImageState {
	instance := getMPSRNNRecurrentImageStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSRNNRecurrentImageStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewRNNRecurrentImageStateWithResource(resource metal.MTLResource) MPSRNNRecurrentImageState {
	instance := getMPSRNNRecurrentImageStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSRNNRecurrentImageStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewRNNRecurrentImageStateWithResources(resources []objectivec.IObject) MPSRNNRecurrentImageState {
	instance := getMPSRNNRecurrentImageStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSRNNRecurrentImageStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNRecurrentImageState/getMemoryCellImage(forLayerIndex:)
func (r MPSRNNRecurrentImageState) GetMemoryCellImageForLayerIndex(layerIndex uint) IMPSImage {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("getMemoryCellImageForLayerIndex:"), layerIndex)
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNRecurrentImageState/getRecurrentOutputImage(forLayerIndex:)
func (r MPSRNNRecurrentImageState) GetRecurrentOutputImageForLayerIndex(layerIndex uint) IMPSImage {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("getRecurrentOutputImageForLayerIndex:"), layerIndex)
	return MPSImageFromID(rv)
}

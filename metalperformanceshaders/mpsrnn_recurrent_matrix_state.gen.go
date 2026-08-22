// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSRNNRecurrentMatrixState] class.
var (
	_MPSRNNRecurrentMatrixStateClass     MPSRNNRecurrentMatrixStateClass
	_MPSRNNRecurrentMatrixStateClassOnce sync.Once
)

func getMPSRNNRecurrentMatrixStateClass() MPSRNNRecurrentMatrixStateClass {
	_MPSRNNRecurrentMatrixStateClassOnce.Do(func() {
		_MPSRNNRecurrentMatrixStateClass = MPSRNNRecurrentMatrixStateClass{class: objc.GetClass("MPSRNNRecurrentMatrixState")}
	})
	return _MPSRNNRecurrentMatrixStateClass
}

// GetMPSRNNRecurrentMatrixStateClass returns the class object for MPSRNNRecurrentMatrixState.
func GetMPSRNNRecurrentMatrixStateClass() MPSRNNRecurrentMatrixStateClass {
	return getMPSRNNRecurrentMatrixStateClass()
}

type MPSRNNRecurrentMatrixStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSRNNRecurrentMatrixStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSRNNRecurrentMatrixStateClass) Alloc() MPSRNNRecurrentMatrixState {
	rv := objc.Send[MPSRNNRecurrentMatrixState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class holds all the data that’s passed from one sequence iteration of
// the matrix-based recurrent neural network layer to the next.
//
// # Instance Methods
//
//   - [MPSRNNRecurrentMatrixState.GetMemoryCellMatrixForLayerIndex]
//   - [MPSRNNRecurrentMatrixState.GetRecurrentOutputMatrixForLayerIndex]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNRecurrentMatrixState
type MPSRNNRecurrentMatrixState struct {
	MPSState
}

// MPSRNNRecurrentMatrixStateFromID constructs a [MPSRNNRecurrentMatrixState] from an objc.ID.
//
// A class holds all the data that’s passed from one sequence iteration of
// the matrix-based recurrent neural network layer to the next.
func MPSRNNRecurrentMatrixStateFromID(id objc.ID) MPSRNNRecurrentMatrixState {
	return MPSRNNRecurrentMatrixState{MPSState: MPSStateFromID(id)}
}

// NOTE: MPSRNNRecurrentMatrixState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSRNNRecurrentMatrixState] class.
//
// # Instance Methods
//
//   - [IMPSRNNRecurrentMatrixState.GetMemoryCellMatrixForLayerIndex]
//   - [IMPSRNNRecurrentMatrixState.GetRecurrentOutputMatrixForLayerIndex]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNRecurrentMatrixState
type IMPSRNNRecurrentMatrixState interface {
	IMPSState

	// Topic: Instance Methods

	GetMemoryCellMatrixForLayerIndex(layerIndex uint) IMPSMatrix
	GetRecurrentOutputMatrixForLayerIndex(layerIndex uint) IMPSMatrix
}

// Init initializes the instance.
func (r MPSRNNRecurrentMatrixState) Init() MPSRNNRecurrentMatrixState {
	rv := objc.Send[MPSRNNRecurrentMatrixState](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSRNNRecurrentMatrixState) Autorelease() MPSRNNRecurrentMatrixState {
	rv := objc.Send[MPSRNNRecurrentMatrixState](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSRNNRecurrentMatrixState creates a new MPSRNNRecurrentMatrixState instance.
func NewMPSRNNRecurrentMatrixState() MPSRNNRecurrentMatrixState {
	class := getMPSRNNRecurrentMatrixStateClass()
	rv := objc.Send[MPSRNNRecurrentMatrixState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewRNNRecurrentMatrixStateWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSRNNRecurrentMatrixState {
	instance := getMPSRNNRecurrentMatrixStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSRNNRecurrentMatrixStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewRNNRecurrentMatrixStateWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSRNNRecurrentMatrixState {
	instance := getMPSRNNRecurrentMatrixStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSRNNRecurrentMatrixStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewRNNRecurrentMatrixStateWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSRNNRecurrentMatrixState {
	instance := getMPSRNNRecurrentMatrixStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSRNNRecurrentMatrixStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewRNNRecurrentMatrixStateWithResource(resource metal.MTLResource) MPSRNNRecurrentMatrixState {
	instance := getMPSRNNRecurrentMatrixStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSRNNRecurrentMatrixStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewRNNRecurrentMatrixStateWithResources(resources []objectivec.IObject) MPSRNNRecurrentMatrixState {
	instance := getMPSRNNRecurrentMatrixStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSRNNRecurrentMatrixStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNRecurrentMatrixState/getMemoryCellMatrix(forLayerIndex:)
func (r MPSRNNRecurrentMatrixState) GetMemoryCellMatrixForLayerIndex(layerIndex uint) IMPSMatrix {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("getMemoryCellMatrixForLayerIndex:"), layerIndex)
	return MPSMatrixFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNRecurrentMatrixState/getRecurrentOutputMatrix(forLayerIndex:)
func (r MPSRNNRecurrentMatrixState) GetRecurrentOutputMatrixForLayerIndex(layerIndex uint) IMPSMatrix {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("getRecurrentOutputMatrixForLayerIndex:"), layerIndex)
	return MPSMatrixFromID(rv)
}

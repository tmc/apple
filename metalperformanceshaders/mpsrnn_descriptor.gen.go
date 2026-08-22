// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSRNNDescriptor] class.
var (
	_MPSRNNDescriptorClass     MPSRNNDescriptorClass
	_MPSRNNDescriptorClassOnce sync.Once
)

func getMPSRNNDescriptorClass() MPSRNNDescriptorClass {
	_MPSRNNDescriptorClassOnce.Do(func() {
		_MPSRNNDescriptorClass = MPSRNNDescriptorClass{class: objc.GetClass("MPSRNNDescriptor")}
	})
	return _MPSRNNDescriptorClass
}

// GetMPSRNNDescriptorClass returns the class object for MPSRNNDescriptor.
func GetMPSRNNDescriptorClass() MPSRNNDescriptorClass {
	return getMPSRNNDescriptorClass()
}

type MPSRNNDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSRNNDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSRNNDescriptorClass) Alloc() MPSRNNDescriptor {
	rv := objc.Send[MPSRNNDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of a recursive neural network block or layer.
//
// # Instance Properties
//
//   - [MPSRNNDescriptor.InputFeatureChannels]
//   - [MPSRNNDescriptor.SetInputFeatureChannels]
//   - [MPSRNNDescriptor.LayerSequenceDirection]
//   - [MPSRNNDescriptor.SetLayerSequenceDirection]
//   - [MPSRNNDescriptor.OutputFeatureChannels]
//   - [MPSRNNDescriptor.SetOutputFeatureChannels]
//   - [MPSRNNDescriptor.UseFloat32Weights]
//   - [MPSRNNDescriptor.SetUseFloat32Weights]
//   - [MPSRNNDescriptor.UseLayerInputUnitTransformMode]
//   - [MPSRNNDescriptor.SetUseLayerInputUnitTransformMode]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNDescriptor
type MPSRNNDescriptor struct {
	objectivec.Object
}

// MPSRNNDescriptorFromID constructs a [MPSRNNDescriptor] from an objc.ID.
//
// A description of a recursive neural network block or layer.
func MPSRNNDescriptorFromID(id objc.ID) MPSRNNDescriptor {
	return MPSRNNDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MPSRNNDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSRNNDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSRNNDescriptor.InputFeatureChannels]
//   - [IMPSRNNDescriptor.SetInputFeatureChannels]
//   - [IMPSRNNDescriptor.LayerSequenceDirection]
//   - [IMPSRNNDescriptor.SetLayerSequenceDirection]
//   - [IMPSRNNDescriptor.OutputFeatureChannels]
//   - [IMPSRNNDescriptor.SetOutputFeatureChannels]
//   - [IMPSRNNDescriptor.UseFloat32Weights]
//   - [IMPSRNNDescriptor.SetUseFloat32Weights]
//   - [IMPSRNNDescriptor.UseLayerInputUnitTransformMode]
//   - [IMPSRNNDescriptor.SetUseLayerInputUnitTransformMode]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNDescriptor
type IMPSRNNDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	InputFeatureChannels() uint
	SetInputFeatureChannels(value uint)
	LayerSequenceDirection() MPSRNNSequenceDirection
	SetLayerSequenceDirection(value MPSRNNSequenceDirection)
	OutputFeatureChannels() uint
	SetOutputFeatureChannels(value uint)
	UseFloat32Weights() bool
	SetUseFloat32Weights(value bool)
	UseLayerInputUnitTransformMode() bool
	SetUseLayerInputUnitTransformMode(value bool)
}

// Init initializes the instance.
func (r MPSRNNDescriptor) Init() MPSRNNDescriptor {
	rv := objc.Send[MPSRNNDescriptor](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MPSRNNDescriptor) Autorelease() MPSRNNDescriptor {
	rv := objc.Send[MPSRNNDescriptor](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSRNNDescriptor creates a new MPSRNNDescriptor instance.
func NewMPSRNNDescriptor() MPSRNNDescriptor {
	class := getMPSRNNDescriptorClass()
	rv := objc.Send[MPSRNNDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNDescriptor/inputFeatureChannels
func (r MPSRNNDescriptor) InputFeatureChannels() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("inputFeatureChannels"))
	return rv
}
func (r MPSRNNDescriptor) SetInputFeatureChannels(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setInputFeatureChannels:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNDescriptor/layerSequenceDirection
func (r MPSRNNDescriptor) LayerSequenceDirection() MPSRNNSequenceDirection {
	rv := objc.Send[MPSRNNSequenceDirection](r.ID, objc.Sel("layerSequenceDirection"))
	return MPSRNNSequenceDirection(rv)
}
func (r MPSRNNDescriptor) SetLayerSequenceDirection(value MPSRNNSequenceDirection) {
	objc.Send[struct{}](r.ID, objc.Sel("setLayerSequenceDirection:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNDescriptor/outputFeatureChannels
func (r MPSRNNDescriptor) OutputFeatureChannels() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("outputFeatureChannels"))
	return rv
}
func (r MPSRNNDescriptor) SetOutputFeatureChannels(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setOutputFeatureChannels:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNDescriptor/useFloat32Weights
func (r MPSRNNDescriptor) UseFloat32Weights() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("useFloat32Weights"))
	return rv
}
func (r MPSRNNDescriptor) SetUseFloat32Weights(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setUseFloat32Weights:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNDescriptor/useLayerInputUnitTransformMode
func (r MPSRNNDescriptor) UseLayerInputUnitTransformMode() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("useLayerInputUnitTransformMode"))
	return rv
}
func (r MPSRNNDescriptor) SetUseLayerInputUnitTransformMode(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setUseLayerInputUnitTransformMode:"), value)
}

// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelStructureNeuralNetwork] class.
var (
	_MLModelStructureNeuralNetworkClass     MLModelStructureNeuralNetworkClass
	_MLModelStructureNeuralNetworkClassOnce sync.Once
)

func getMLModelStructureNeuralNetworkClass() MLModelStructureNeuralNetworkClass {
	_MLModelStructureNeuralNetworkClassOnce.Do(func() {
		_MLModelStructureNeuralNetworkClass = MLModelStructureNeuralNetworkClass{class: objc.GetClass("MLModelStructureNeuralNetwork")}
	})
	return _MLModelStructureNeuralNetworkClass
}

// GetMLModelStructureNeuralNetworkClass returns the class object for MLModelStructureNeuralNetwork.
func GetMLModelStructureNeuralNetworkClass() MLModelStructureNeuralNetworkClass {
	return getMLModelStructureNeuralNetworkClass()
}

type MLModelStructureNeuralNetworkClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelStructureNeuralNetworkClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureNeuralNetworkClass) Alloc() MLModelStructureNeuralNetwork {
	rv := objc.Send[MLModelStructureNeuralNetwork](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class representing the structure of a NeuralNetwork model.
//
// # Accessing the layers
//
//   - [MLModelStructureNeuralNetwork.Layers]: The topologically sorted layers in the NeuralNetwork.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetwork
type MLModelStructureNeuralNetwork struct {
	objectivec.Object
}

// MLModelStructureNeuralNetworkFromID constructs a [MLModelStructureNeuralNetwork] from an objc.ID.
//
// A class representing the structure of a NeuralNetwork model.
func MLModelStructureNeuralNetworkFromID(id objc.ID) MLModelStructureNeuralNetwork {
	return MLModelStructureNeuralNetwork{objectivec.Object{ID: id}}
}
// Ensure MLModelStructureNeuralNetwork implements IMLModelStructureNeuralNetwork.
var _ IMLModelStructureNeuralNetwork = MLModelStructureNeuralNetwork{}

// An interface definition for the [MLModelStructureNeuralNetwork] class.
//
// # Accessing the layers
//
//   - [IMLModelStructureNeuralNetwork.Layers]: The topologically sorted layers in the NeuralNetwork.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetwork
type IMLModelStructureNeuralNetwork interface {
	objectivec.IObject

	// Topic: Accessing the layers

	// The topologically sorted layers in the NeuralNetwork.
	Layers() []MLModelStructureNeuralNetworkLayer
}

// Init initializes the instance.
func (m MLModelStructureNeuralNetwork) Init() MLModelStructureNeuralNetwork {
	rv := objc.Send[MLModelStructureNeuralNetwork](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelStructureNeuralNetwork) Autorelease() MLModelStructureNeuralNetwork {
	rv := objc.Send[MLModelStructureNeuralNetwork](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelStructureNeuralNetwork creates a new MLModelStructureNeuralNetwork instance.
func NewMLModelStructureNeuralNetwork() MLModelStructureNeuralNetwork {
	class := getMLModelStructureNeuralNetworkClass()
	rv := objc.Send[MLModelStructureNeuralNetwork](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The topologically sorted layers in the NeuralNetwork.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetwork/layers
func (m MLModelStructureNeuralNetwork) Layers() []MLModelStructureNeuralNetworkLayer {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("layers"))
	return objc.ConvertSlice(rv, func(id objc.ID) MLModelStructureNeuralNetworkLayer {
		return MLModelStructureNeuralNetworkLayerFromID(id)
	})
}


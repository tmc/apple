// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelStructureNeuralNetworkLayer] class.
var (
	_MLModelStructureNeuralNetworkLayerClass     MLModelStructureNeuralNetworkLayerClass
	_MLModelStructureNeuralNetworkLayerClassOnce sync.Once
)

func getMLModelStructureNeuralNetworkLayerClass() MLModelStructureNeuralNetworkLayerClass {
	_MLModelStructureNeuralNetworkLayerClassOnce.Do(func() {
		_MLModelStructureNeuralNetworkLayerClass = MLModelStructureNeuralNetworkLayerClass{class: objc.GetClass("MLModelStructureNeuralNetworkLayer")}
	})
	return _MLModelStructureNeuralNetworkLayerClass
}

// GetMLModelStructureNeuralNetworkLayerClass returns the class object for MLModelStructureNeuralNetworkLayer.
func GetMLModelStructureNeuralNetworkLayerClass() MLModelStructureNeuralNetworkLayerClass {
	return getMLModelStructureNeuralNetworkLayerClass()
}

type MLModelStructureNeuralNetworkLayerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureNeuralNetworkLayerClass) Alloc() MLModelStructureNeuralNetworkLayer {
	rv := objc.Send[MLModelStructureNeuralNetworkLayer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A class representing a layer in a NeuralNetwork.
//
// # Accessing the network layer properties
//
//   - [MLModelStructureNeuralNetworkLayer.InputNames]: The input names.
//   - [MLModelStructureNeuralNetworkLayer.Name]: The layer name.
//   - [MLModelStructureNeuralNetworkLayer.OutputNames]: The output names.
//   - [MLModelStructureNeuralNetworkLayer.Type]: The type of the layer, e,g, “elementwise”, “pooling”, etc.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer
type MLModelStructureNeuralNetworkLayer struct {
	objectivec.Object
}

// MLModelStructureNeuralNetworkLayerFromID constructs a [MLModelStructureNeuralNetworkLayer] from an objc.ID.
//
// A class representing a layer in a NeuralNetwork.
func MLModelStructureNeuralNetworkLayerFromID(id objc.ID) MLModelStructureNeuralNetworkLayer {
	return MLModelStructureNeuralNetworkLayer{objectivec.Object{id}}
}
// Ensure MLModelStructureNeuralNetworkLayer implements IMLModelStructureNeuralNetworkLayer.
var _ IMLModelStructureNeuralNetworkLayer = MLModelStructureNeuralNetworkLayer{}





// An interface definition for the [MLModelStructureNeuralNetworkLayer] class.
//
// # Accessing the network layer properties
//
//   - [IMLModelStructureNeuralNetworkLayer.InputNames]: The input names.
//   - [IMLModelStructureNeuralNetworkLayer.Name]: The layer name.
//   - [IMLModelStructureNeuralNetworkLayer.OutputNames]: The output names.
//   - [IMLModelStructureNeuralNetworkLayer.Type]: The type of the layer, e,g, “elementwise”, “pooling”, etc.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer
type IMLModelStructureNeuralNetworkLayer interface {
	objectivec.IObject

	// Topic: Accessing the network layer properties

	// The input names.
	InputNames() []string
	// The layer name.
	Name() string
	// The output names.
	OutputNames() []string
	// The type of the layer, e,g, “elementwise”, “pooling”, etc.
	Type() string
}





// Init initializes the instance.
func (m MLModelStructureNeuralNetworkLayer) Init() MLModelStructureNeuralNetworkLayer {
	rv := objc.Send[MLModelStructureNeuralNetworkLayer](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelStructureNeuralNetworkLayer) Autorelease() MLModelStructureNeuralNetworkLayer {
	rv := objc.Send[MLModelStructureNeuralNetworkLayer](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelStructureNeuralNetworkLayer creates a new MLModelStructureNeuralNetworkLayer instance.
func NewMLModelStructureNeuralNetworkLayer() MLModelStructureNeuralNetworkLayer {
	class := getMLModelStructureNeuralNetworkLayerClass()
	rv := objc.Send[MLModelStructureNeuralNetworkLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// The input names.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer/inputNames
func (m MLModelStructureNeuralNetworkLayer) InputNames() []string {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("inputNames"))
	return objc.ConvertSliceToStrings(rv)
}



// The layer name.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer/name
func (m MLModelStructureNeuralNetworkLayer) Name() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}



// The output names.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer/outputNames
func (m MLModelStructureNeuralNetworkLayer) OutputNames() []string {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("outputNames"))
	return objc.ConvertSliceToStrings(rv)
}



// The type of the layer, e,g, “elementwise”, “pooling”, etc.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer/type
func (m MLModelStructureNeuralNetworkLayer) Type() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("type"))
	return foundation.NSStringFromID(rv).String()
}


















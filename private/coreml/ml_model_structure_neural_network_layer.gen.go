// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
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

// Class returns the underlying Objective-C class pointer.
func (mc MLModelStructureNeuralNetworkLayerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureNeuralNetworkLayerClass) Alloc() MLModelStructureNeuralNetworkLayer {
	rv := objc.Send[MLModelStructureNeuralNetworkLayer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLModelStructureNeuralNetworkLayer.Path]
//   - [MLModelStructureNeuralNetworkLayer.InitWithNameTypeInputNamesOutputNamesPath]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer
type MLModelStructureNeuralNetworkLayer struct {
	objectivec.Object
}

// MLModelStructureNeuralNetworkLayerFromID constructs a [MLModelStructureNeuralNetworkLayer] from an objc.ID.
func MLModelStructureNeuralNetworkLayerFromID(id objc.ID) MLModelStructureNeuralNetworkLayer {
	return MLModelStructureNeuralNetworkLayer{objectivec.Object{ID: id}}
}

// Ensure MLModelStructureNeuralNetworkLayer implements IMLModelStructureNeuralNetworkLayer.
var _ IMLModelStructureNeuralNetworkLayer = MLModelStructureNeuralNetworkLayer{}

// An interface definition for the [MLModelStructureNeuralNetworkLayer] class.
//
// # Methods
//
//   - [IMLModelStructureNeuralNetworkLayer.Path]
//   - [IMLModelStructureNeuralNetworkLayer.InitWithNameTypeInputNamesOutputNamesPath]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer
type IMLModelStructureNeuralNetworkLayer interface {
	objectivec.IObject

	// Topic: Methods

	Path() IMLModelStructurePath
	InitWithNameTypeInputNamesOutputNamesPath(name objectivec.IObject, type_ objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, path objectivec.IObject) MLModelStructureNeuralNetworkLayer
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

// See: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer/initWithName:type:inputNames:outputNames:path:
func NewModelStructureNeuralNetworkLayerWithNameTypeInputNamesOutputNamesPath(name objectivec.IObject, type_ objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, path objectivec.IObject) MLModelStructureNeuralNetworkLayer {
	instance := getMLModelStructureNeuralNetworkLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:type:inputNames:outputNames:path:"), name, type_, names, names2, path)
	return MLModelStructureNeuralNetworkLayerFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer/initWithName:type:inputNames:outputNames:path:
func (m MLModelStructureNeuralNetworkLayer) InitWithNameTypeInputNamesOutputNamesPath(name objectivec.IObject, type_ objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, path objectivec.IObject) MLModelStructureNeuralNetworkLayer {
	rv := objc.Send[MLModelStructureNeuralNetworkLayer](m.ID, objc.Sel("initWithName:type:inputNames:outputNames:path:"), name, type_, names, names2, path)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer/path
func (m MLModelStructureNeuralNetworkLayer) Path() IMLModelStructurePath {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("path"))
	return MLModelStructurePathFromID(objc.ID(rv))
}

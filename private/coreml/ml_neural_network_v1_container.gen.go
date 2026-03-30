// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNeuralNetworkV1Container] class.
var (
	_MLNeuralNetworkV1ContainerClass     MLNeuralNetworkV1ContainerClass
	_MLNeuralNetworkV1ContainerClassOnce sync.Once
)

func getMLNeuralNetworkV1ContainerClass() MLNeuralNetworkV1ContainerClass {
	_MLNeuralNetworkV1ContainerClassOnce.Do(func() {
		_MLNeuralNetworkV1ContainerClass = MLNeuralNetworkV1ContainerClass{class: objc.GetClass("MLNeuralNetworkV1Container")}
	})
	return _MLNeuralNetworkV1ContainerClass
}

// GetMLNeuralNetworkV1ContainerClass returns the class object for MLNeuralNetworkV1Container.
func GetMLNeuralNetworkV1ContainerClass() MLNeuralNetworkV1ContainerClass {
	return getMLNeuralNetworkV1ContainerClass()
}

type MLNeuralNetworkV1ContainerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNeuralNetworkV1ContainerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNeuralNetworkV1ContainerClass) Alloc() MLNeuralNetworkV1Container {
	rv := objc.Send[MLNeuralNetworkV1Container](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkV1Container
type MLNeuralNetworkV1Container struct {
	MLNeuralNetworkContainer
}

// MLNeuralNetworkV1ContainerFromID constructs a [MLNeuralNetworkV1Container] from an objc.ID.
func MLNeuralNetworkV1ContainerFromID(id objc.ID) MLNeuralNetworkV1Container {
	return MLNeuralNetworkV1Container{MLNeuralNetworkContainer: MLNeuralNetworkContainerFromID(id)}
}

// Ensure MLNeuralNetworkV1Container implements IMLNeuralNetworkV1Container.
var _ IMLNeuralNetworkV1Container = MLNeuralNetworkV1Container{}

// An interface definition for the [MLNeuralNetworkV1Container] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkV1Container
type IMLNeuralNetworkV1Container interface {
	IMLNeuralNetworkContainer
}

// Init initializes the instance.
func (n MLNeuralNetworkV1Container) Init() MLNeuralNetworkV1Container {
	rv := objc.Send[MLNeuralNetworkV1Container](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MLNeuralNetworkV1Container) Autorelease() MLNeuralNetworkV1Container {
	rv := objc.Send[MLNeuralNetworkV1Container](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralNetworkV1Container creates a new MLNeuralNetworkV1Container instance.
func NewMLNeuralNetworkV1Container() MLNeuralNetworkV1Container {
	class := getMLNeuralNetworkV1ContainerClass()
	rv := objc.Send[MLNeuralNetworkV1Container](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:
func NewNeuralNetworkV1ContainerWithFeatureDescriptionsModelDescriptionOutputLayerNamesClassScoreVectorNameClassLabelsIsEncryptedModelVersionInfo(descriptions objectivec.IObject, description objectivec.IObject, names objectivec.IObject, name objectivec.IObject, labels objectivec.IObject, encrypted bool, info objectivec.IObject) MLNeuralNetworkV1Container {
	instance := getMLNeuralNetworkV1ContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:"), descriptions, description, names, name, labels, encrypted, info)
	return MLNeuralNetworkV1ContainerFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:compilerVersionInfo:
func NewNeuralNetworkV1ContainerWithFeatureDescriptionsModelDescriptionOutputLayerNamesClassScoreVectorNameClassLabelsIsEncryptedModelVersionInfoCompilerVersionInfo(descriptions objectivec.IObject, description objectivec.IObject, names objectivec.IObject, name objectivec.IObject, labels objectivec.IObject, encrypted bool, info objectivec.IObject, info2 objectivec.IObject) MLNeuralNetworkV1Container {
	instance := getMLNeuralNetworkV1ContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:compilerVersionInfo:"), descriptions, description, names, name, labels, encrypted, info, info2)
	return MLNeuralNetworkV1ContainerFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/initWithFilePath:inputLayerNames:outputLayerNames:parameters:
func NewNeuralNetworkV1ContainerWithFilePathInputLayerNamesOutputLayerNamesParameters(path objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, parameters objectivec.IObject) MLNeuralNetworkV1Container {
	instance := getMLNeuralNetworkV1ContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFilePath:inputLayerNames:outputLayerNames:parameters:"), path, names, names2, parameters)
	return MLNeuralNetworkV1ContainerFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkV1Container/readerFromArchive:error:
func (_MLNeuralNetworkV1ContainerClass MLNeuralNetworkV1ContainerClass) ReaderFromArchiveError(archive unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkV1ContainerClass.class), objc.Sel("readerFromArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

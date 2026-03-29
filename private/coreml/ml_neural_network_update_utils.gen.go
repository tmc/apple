// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNeuralNetworkUpdateUtils] class.
var (
	_MLNeuralNetworkUpdateUtilsClass     MLNeuralNetworkUpdateUtilsClass
	_MLNeuralNetworkUpdateUtilsClassOnce sync.Once
)

func getMLNeuralNetworkUpdateUtilsClass() MLNeuralNetworkUpdateUtilsClass {
	_MLNeuralNetworkUpdateUtilsClassOnce.Do(func() {
		_MLNeuralNetworkUpdateUtilsClass = MLNeuralNetworkUpdateUtilsClass{class: objc.GetClass("MLNeuralNetworkUpdateUtils")}
	})
	return _MLNeuralNetworkUpdateUtilsClass
}

// GetMLNeuralNetworkUpdateUtilsClass returns the class object for MLNeuralNetworkUpdateUtils.
func GetMLNeuralNetworkUpdateUtilsClass() MLNeuralNetworkUpdateUtilsClass {
	return getMLNeuralNetworkUpdateUtilsClass()
}

type MLNeuralNetworkUpdateUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNeuralNetworkUpdateUtilsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNeuralNetworkUpdateUtilsClass) Alloc() MLNeuralNetworkUpdateUtils {
	rv := objc.Send[MLNeuralNetworkUpdateUtils](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateUtils
type MLNeuralNetworkUpdateUtils struct {
	objectivec.Object
}

// MLNeuralNetworkUpdateUtilsFromID constructs a [MLNeuralNetworkUpdateUtils] from an objc.ID.
func MLNeuralNetworkUpdateUtilsFromID(id objc.ID) MLNeuralNetworkUpdateUtils {
	return MLNeuralNetworkUpdateUtils{objectivec.Object{ID: id}}
}
// Ensure MLNeuralNetworkUpdateUtils implements IMLNeuralNetworkUpdateUtils.
var _ IMLNeuralNetworkUpdateUtils = MLNeuralNetworkUpdateUtils{}

// An interface definition for the [MLNeuralNetworkUpdateUtils] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateUtils
type IMLNeuralNetworkUpdateUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (n MLNeuralNetworkUpdateUtils) Init() MLNeuralNetworkUpdateUtils {
	rv := objc.Send[MLNeuralNetworkUpdateUtils](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MLNeuralNetworkUpdateUtils) Autorelease() MLNeuralNetworkUpdateUtils {
	rv := objc.Send[MLNeuralNetworkUpdateUtils](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralNetworkUpdateUtils creates a new MLNeuralNetworkUpdateUtils instance.
func NewMLNeuralNetworkUpdateUtils() MLNeuralNetworkUpdateUtils {
	class := getMLNeuralNetworkUpdateUtilsClass()
	rv := objc.Send[MLNeuralNetworkUpdateUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateUtils/createClassLabelToIndexMapWith:
func (_MLNeuralNetworkUpdateUtilsClass MLNeuralNetworkUpdateUtilsClass) CreateClassLabelToIndexMapWith(with objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkUpdateUtilsClass.class), objc.Sel("createClassLabelToIndexMapWith:"), with)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateUtils/loadParameterDescriptionsAndContainerFromUpdateParameters:modelDescription:
func (_MLNeuralNetworkUpdateUtilsClass MLNeuralNetworkUpdateUtilsClass) LoadParameterDescriptionsAndContainerFromUpdateParametersModelDescription(parameters unsafe.Pointer, description objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkUpdateUtilsClass.class), objc.Sel("loadParameterDescriptionsAndContainerFromUpdateParameters:modelDescription:"), parameters, description)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateUtils/loadUpdateParameters:fromCompiledArchive:error:
func (_MLNeuralNetworkUpdateUtilsClass MLNeuralNetworkUpdateUtilsClass) LoadUpdateParametersFromCompiledArchiveError(parameters unsafe.Pointer, archive unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLNeuralNetworkUpdateUtilsClass.class), objc.Sel("loadUpdateParameters:fromCompiledArchive:error:"), parameters, archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadUpdateParameters:fromCompiledArchive:error: returned NO with nil NSError")
	}
	return rv, nil

}


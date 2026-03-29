// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLAppleImageFeatureExtractorScenePrintParameters] class.
var (
	_MLAppleImageFeatureExtractorScenePrintParametersClass     MLAppleImageFeatureExtractorScenePrintParametersClass
	_MLAppleImageFeatureExtractorScenePrintParametersClassOnce sync.Once
)

func getMLAppleImageFeatureExtractorScenePrintParametersClass() MLAppleImageFeatureExtractorScenePrintParametersClass {
	_MLAppleImageFeatureExtractorScenePrintParametersClassOnce.Do(func() {
		_MLAppleImageFeatureExtractorScenePrintParametersClass = MLAppleImageFeatureExtractorScenePrintParametersClass{class: objc.GetClass("MLAppleImageFeatureExtractorScenePrintParameters")}
	})
	return _MLAppleImageFeatureExtractorScenePrintParametersClass
}

// GetMLAppleImageFeatureExtractorScenePrintParametersClass returns the class object for MLAppleImageFeatureExtractorScenePrintParameters.
func GetMLAppleImageFeatureExtractorScenePrintParametersClass() MLAppleImageFeatureExtractorScenePrintParametersClass {
	return getMLAppleImageFeatureExtractorScenePrintParametersClass()
}

type MLAppleImageFeatureExtractorScenePrintParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLAppleImageFeatureExtractorScenePrintParametersClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLAppleImageFeatureExtractorScenePrintParametersClass) Alloc() MLAppleImageFeatureExtractorScenePrintParameters {
	rv := objc.Send[MLAppleImageFeatureExtractorScenePrintParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLAppleImageFeatureExtractorScenePrintParameters.RequestClassName]
//   - [MLAppleImageFeatureExtractorScenePrintParameters.ScenePrintVersion]
//   - [MLAppleImageFeatureExtractorScenePrintParameters.InitScenePrintParametersRequestClassError]
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorScenePrintParameters
type MLAppleImageFeatureExtractorScenePrintParameters struct {
	objectivec.Object
}

// MLAppleImageFeatureExtractorScenePrintParametersFromID constructs a [MLAppleImageFeatureExtractorScenePrintParameters] from an objc.ID.
func MLAppleImageFeatureExtractorScenePrintParametersFromID(id objc.ID) MLAppleImageFeatureExtractorScenePrintParameters {
	return MLAppleImageFeatureExtractorScenePrintParameters{objectivec.Object{ID: id}}
}
// Ensure MLAppleImageFeatureExtractorScenePrintParameters implements IMLAppleImageFeatureExtractorScenePrintParameters.
var _ IMLAppleImageFeatureExtractorScenePrintParameters = MLAppleImageFeatureExtractorScenePrintParameters{}

// An interface definition for the [MLAppleImageFeatureExtractorScenePrintParameters] class.
//
// # Methods
//
//   - [IMLAppleImageFeatureExtractorScenePrintParameters.RequestClassName]
//   - [IMLAppleImageFeatureExtractorScenePrintParameters.ScenePrintVersion]
//   - [IMLAppleImageFeatureExtractorScenePrintParameters.InitScenePrintParametersRequestClassError]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorScenePrintParameters
type IMLAppleImageFeatureExtractorScenePrintParameters interface {
	objectivec.IObject

	// Topic: Methods

	RequestClassName() string
	ScenePrintVersion() uint64
	InitScenePrintParametersRequestClassError(parameters uint64, class objectivec.IObject) (MLAppleImageFeatureExtractorScenePrintParameters, error)
}

// Init initializes the instance.
func (a MLAppleImageFeatureExtractorScenePrintParameters) Init() MLAppleImageFeatureExtractorScenePrintParameters {
	rv := objc.Send[MLAppleImageFeatureExtractorScenePrintParameters](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLAppleImageFeatureExtractorScenePrintParameters) Autorelease() MLAppleImageFeatureExtractorScenePrintParameters {
	rv := objc.Send[MLAppleImageFeatureExtractorScenePrintParameters](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAppleImageFeatureExtractorScenePrintParameters creates a new MLAppleImageFeatureExtractorScenePrintParameters instance.
func NewMLAppleImageFeatureExtractorScenePrintParameters() MLAppleImageFeatureExtractorScenePrintParameters {
	class := getMLAppleImageFeatureExtractorScenePrintParametersClass()
	rv := objc.Send[MLAppleImageFeatureExtractorScenePrintParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorScenePrintParameters/initScenePrintParameters:requestClass:error:
func NewAppleImageFeatureExtractorScenePrintParametersScenePrintParametersRequestClassError(parameters uint64, class objectivec.IObject) (MLAppleImageFeatureExtractorScenePrintParameters, error) {
	var errorPtr objc.ID
	instance := getMLAppleImageFeatureExtractorScenePrintParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initScenePrintParameters:requestClass:error:"), parameters, class, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleImageFeatureExtractorScenePrintParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleImageFeatureExtractorScenePrintParametersFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorScenePrintParameters/initScenePrintParameters:requestClass:error:
func (a MLAppleImageFeatureExtractorScenePrintParameters) InitScenePrintParametersRequestClassError(parameters uint64, class objectivec.IObject) (MLAppleImageFeatureExtractorScenePrintParameters, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initScenePrintParameters:requestClass:error:"), parameters, class, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleImageFeatureExtractorScenePrintParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleImageFeatureExtractorScenePrintParametersFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorScenePrintParameters/requestClassName
func (a MLAppleImageFeatureExtractorScenePrintParameters) RequestClassName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("requestClassName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorScenePrintParameters/scenePrintVersion
func (a MLAppleImageFeatureExtractorScenePrintParameters) ScenePrintVersion() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("scenePrintVersion"))
	return rv
}


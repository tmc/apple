// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLAppleImageFeatureExtractorParameters] class.
var (
	_MLAppleImageFeatureExtractorParametersClass     MLAppleImageFeatureExtractorParametersClass
	_MLAppleImageFeatureExtractorParametersClassOnce sync.Once
)

func getMLAppleImageFeatureExtractorParametersClass() MLAppleImageFeatureExtractorParametersClass {
	_MLAppleImageFeatureExtractorParametersClassOnce.Do(func() {
		_MLAppleImageFeatureExtractorParametersClass = MLAppleImageFeatureExtractorParametersClass{class: objc.GetClass("MLAppleImageFeatureExtractorParameters")}
	})
	return _MLAppleImageFeatureExtractorParametersClass
}

// GetMLAppleImageFeatureExtractorParametersClass returns the class object for MLAppleImageFeatureExtractorParameters.
func GetMLAppleImageFeatureExtractorParametersClass() MLAppleImageFeatureExtractorParametersClass {
	return getMLAppleImageFeatureExtractorParametersClass()
}

type MLAppleImageFeatureExtractorParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLAppleImageFeatureExtractorParametersClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLAppleImageFeatureExtractorParametersClass) Alloc() MLAppleImageFeatureExtractorParameters {
	rv := objc.Send[MLAppleImageFeatureExtractorParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLAppleImageFeatureExtractorParameters.FeatureExtractorParameters]
//   - [MLAppleImageFeatureExtractorParameters.InitWithObjectPrintParametersError]
//   - [MLAppleImageFeatureExtractorParameters.InitWithScenePrintParametersError]
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorParameters
type MLAppleImageFeatureExtractorParameters struct {
	objectivec.Object
}

// MLAppleImageFeatureExtractorParametersFromID constructs a [MLAppleImageFeatureExtractorParameters] from an objc.ID.
func MLAppleImageFeatureExtractorParametersFromID(id objc.ID) MLAppleImageFeatureExtractorParameters {
	return MLAppleImageFeatureExtractorParameters{objectivec.Object{ID: id}}
}
// Ensure MLAppleImageFeatureExtractorParameters implements IMLAppleImageFeatureExtractorParameters.
var _ IMLAppleImageFeatureExtractorParameters = MLAppleImageFeatureExtractorParameters{}

// An interface definition for the [MLAppleImageFeatureExtractorParameters] class.
//
// # Methods
//
//   - [IMLAppleImageFeatureExtractorParameters.FeatureExtractorParameters]
//   - [IMLAppleImageFeatureExtractorParameters.InitWithObjectPrintParametersError]
//   - [IMLAppleImageFeatureExtractorParameters.InitWithScenePrintParametersError]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorParameters
type IMLAppleImageFeatureExtractorParameters interface {
	objectivec.IObject

	// Topic: Methods

	FeatureExtractorParameters() objectivec.IObject
	InitWithObjectPrintParametersError(parameters objectivec.IObject) (MLAppleImageFeatureExtractorParameters, error)
	InitWithScenePrintParametersError(parameters objectivec.IObject) (MLAppleImageFeatureExtractorParameters, error)
}

// Init initializes the instance.
func (a MLAppleImageFeatureExtractorParameters) Init() MLAppleImageFeatureExtractorParameters {
	rv := objc.Send[MLAppleImageFeatureExtractorParameters](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLAppleImageFeatureExtractorParameters) Autorelease() MLAppleImageFeatureExtractorParameters {
	rv := objc.Send[MLAppleImageFeatureExtractorParameters](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAppleImageFeatureExtractorParameters creates a new MLAppleImageFeatureExtractorParameters instance.
func NewMLAppleImageFeatureExtractorParameters() MLAppleImageFeatureExtractorParameters {
	class := getMLAppleImageFeatureExtractorParametersClass()
	rv := objc.Send[MLAppleImageFeatureExtractorParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorParameters/initWithObjectPrintParameters:error:
func NewAppleImageFeatureExtractorParametersWithObjectPrintParametersError(parameters objectivec.IObject) (MLAppleImageFeatureExtractorParameters, error) {
	var errorPtr objc.ID
	instance := getMLAppleImageFeatureExtractorParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObjectPrintParameters:error:"), parameters, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleImageFeatureExtractorParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleImageFeatureExtractorParametersFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorParameters/initWithScenePrintParameters:error:
func NewAppleImageFeatureExtractorParametersWithScenePrintParametersError(parameters objectivec.IObject) (MLAppleImageFeatureExtractorParameters, error) {
	var errorPtr objc.ID
	instance := getMLAppleImageFeatureExtractorParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithScenePrintParameters:error:"), parameters, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleImageFeatureExtractorParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleImageFeatureExtractorParametersFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorParameters/initWithObjectPrintParameters:error:
func (a MLAppleImageFeatureExtractorParameters) InitWithObjectPrintParametersError(parameters objectivec.IObject) (MLAppleImageFeatureExtractorParameters, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithObjectPrintParameters:error:"), parameters, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleImageFeatureExtractorParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleImageFeatureExtractorParametersFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorParameters/initWithScenePrintParameters:error:
func (a MLAppleImageFeatureExtractorParameters) InitWithScenePrintParametersError(parameters objectivec.IObject) (MLAppleImageFeatureExtractorParameters, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithScenePrintParameters:error:"), parameters, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleImageFeatureExtractorParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleImageFeatureExtractorParametersFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorParameters/featureExtractorParameters
func (a MLAppleImageFeatureExtractorParameters) FeatureExtractorParameters() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("featureExtractorParameters"))
	return objectivec.Object{ID: rv}
}


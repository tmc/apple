// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNonMaximumSuppression] class.
var (
	_MLNonMaximumSuppressionClass     MLNonMaximumSuppressionClass
	_MLNonMaximumSuppressionClassOnce sync.Once
)

func getMLNonMaximumSuppressionClass() MLNonMaximumSuppressionClass {
	_MLNonMaximumSuppressionClassOnce.Do(func() {
		_MLNonMaximumSuppressionClass = MLNonMaximumSuppressionClass{class: objc.GetClass("MLNonMaximumSuppression")}
	})
	return _MLNonMaximumSuppressionClass
}

// GetMLNonMaximumSuppressionClass returns the class object for MLNonMaximumSuppression.
func GetMLNonMaximumSuppressionClass() MLNonMaximumSuppressionClass {
	return getMLNonMaximumSuppressionClass()
}

type MLNonMaximumSuppressionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNonMaximumSuppressionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNonMaximumSuppressionClass) Alloc() MLNonMaximumSuppression {
	rv := objc.SendIfResponds[MLNonMaximumSuppression](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLNonMaximumSuppression.Parameters]
//   - [MLNonMaximumSuppression.InitWithParametersModelDescriptionConfigurationError]
type MLNonMaximumSuppression struct {
	MLModelEngine
}

// MLNonMaximumSuppressionFromID constructs a [MLNonMaximumSuppression] from an objc.ID.
func MLNonMaximumSuppressionFromID(id objc.ID) MLNonMaximumSuppression {
	return MLNonMaximumSuppression{MLModelEngine: MLModelEngineFromID(id)}
}

// Ensure MLNonMaximumSuppression implements IMLNonMaximumSuppression.
var _ IMLNonMaximumSuppression = MLNonMaximumSuppression{}

// An interface definition for the [MLNonMaximumSuppression] class.
//
// # Methods
//
//   - [IMLNonMaximumSuppression.Parameters]
//   - [IMLNonMaximumSuppression.InitWithParametersModelDescriptionConfigurationError]
type IMLNonMaximumSuppression interface {
	IMLModelEngine

	// Topic: Methods

	Parameters() IMLNonMaximumSuppressionParameters
	InitWithParametersModelDescriptionConfigurationError(parameters objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLNonMaximumSuppression, error)
}

// Init initializes the instance.
func (m MLNonMaximumSuppression) Init() MLNonMaximumSuppression {
	rv := objc.SendIfResponds[MLNonMaximumSuppression](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNonMaximumSuppression) Autorelease() MLNonMaximumSuppression {
	rv := objc.SendIfResponds[MLNonMaximumSuppression](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNonMaximumSuppression creates a new MLNonMaximumSuppression instance.
func NewMLNonMaximumSuppression() MLNonMaximumSuppression {
	class := getMLNonMaximumSuppressionClass()
	rv := objc.SendIfResponds[MLNonMaximumSuppression](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewNonMaximumSuppressionWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLNonMaximumSuppression {
	instance := getMLNonMaximumSuppressionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLNonMaximumSuppressionFromID(rv)
}

func NewNonMaximumSuppressionWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLNonMaximumSuppression {
	instance := getMLNonMaximumSuppressionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLNonMaximumSuppressionFromID(rv)
}

func NewNonMaximumSuppressionWithParametersModelDescriptionConfigurationError(parameters objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLNonMaximumSuppression, error) {
	var errorPtr objc.ID
	instance := getMLNonMaximumSuppressionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParameters:modelDescription:configuration:error:"), parameters, description, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNonMaximumSuppression{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLNonMaximumSuppression{}, objc.ErrInitFailed
	}
	return MLNonMaximumSuppressionFromID(rv), nil
}

func (m MLNonMaximumSuppression) InitWithParametersModelDescriptionConfigurationError(parameters objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLNonMaximumSuppression, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithParameters:modelDescription:configuration:error:"), parameters, description, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNonMaximumSuppression{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNonMaximumSuppressionFromID(rv), nil

}

func (_MLNonMaximumSuppressionClass MLNonMaximumSuppressionClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNonMaximumSuppressionClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

func (m MLNonMaximumSuppression) Parameters() IMLNonMaximumSuppressionParameters {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("parameters"))
	return MLNonMaximumSuppressionParametersFromID(objc.ID(rv))
}

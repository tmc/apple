// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLIdentity] class.
var (
	_MLIdentityClass     MLIdentityClass
	_MLIdentityClassOnce sync.Once
)

func getMLIdentityClass() MLIdentityClass {
	_MLIdentityClassOnce.Do(func() {
		_MLIdentityClass = MLIdentityClass{class: objc.GetClass("MLIdentity")}
	})
	return _MLIdentityClass
}

// GetMLIdentityClass returns the class object for MLIdentity.
func GetMLIdentityClass() MLIdentityClass {
	return getMLIdentityClass()
}

type MLIdentityClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLIdentityClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLIdentityClass) Alloc() MLIdentity {
	rv := objc.Send[MLIdentity](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLIdentity.PredictionFromFeaturesOptionsError]
//
// See: https://developer.apple.com/documentation/CoreML/MLIdentity
type MLIdentity struct {
	MLModel
}

// MLIdentityFromID constructs a [MLIdentity] from an objc.ID.
func MLIdentityFromID(id objc.ID) MLIdentity {
	return MLIdentity{MLModel: MLModelFromID(id)}
}

// Ensure MLIdentity implements IMLIdentity.
var _ IMLIdentity = MLIdentity{}

// An interface definition for the [MLIdentity] class.
//
// # Methods
//
//   - [IMLIdentity.PredictionFromFeaturesOptionsError]
//
// See: https://developer.apple.com/documentation/CoreML/MLIdentity
type IMLIdentity interface {
	IMLModel

	// Topic: Methods

	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
}

// Init initializes the instance.
func (i MLIdentity) Init() MLIdentity {
	rv := objc.Send[MLIdentity](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MLIdentity) Autorelease() MLIdentity {
	rv := objc.Send[MLIdentity](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLIdentity creates a new MLIdentity instance.
func NewMLIdentity() MLIdentity {
	class := getMLIdentityClass()
	rv := objc.Send[MLIdentity](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewIdentityDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLIdentity, error) {
	var errorPtr objc.ID
	instance := getMLIdentityClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLIdentity{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLIdentityFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewIdentityInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLIdentity, error) {
	var errorPtr objc.ID
	instance := getMLIdentityClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLIdentity{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLIdentityFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewIdentityWithConfiguration(configuration objectivec.IObject) MLIdentity {
	instance := getMLIdentityClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLIdentityFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewIdentityWithDescription(description objectivec.IObject) MLIdentity {
	instance := getMLIdentityClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLIdentityFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewIdentityWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLIdentity {
	instance := getMLIdentityClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLIdentityFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewIdentityWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLIdentity {
	instance := getMLIdentityClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLIdentityFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLIdentity/predictionFromFeatures:options:error:
func (i MLIdentity) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLIdentity/loadModelFromSpecification:configuration:error:
func (_MLIdentityClass MLIdentityClass) LoadModelFromSpecificationConfigurationError(specification MLModelSpecificationRef, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLIdentityClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLFeatureProviderConformer] class.
var (
	_MLFeatureProviderConformerClass     MLFeatureProviderConformerClass
	_MLFeatureProviderConformerClassOnce sync.Once
)

func getMLFeatureProviderConformerClass() MLFeatureProviderConformerClass {
	_MLFeatureProviderConformerClassOnce.Do(func() {
		_MLFeatureProviderConformerClass = MLFeatureProviderConformerClass{class: objc.GetClass("MLFeatureProviderConformer")}
	})
	return _MLFeatureProviderConformerClass
}

// GetMLFeatureProviderConformerClass returns the class object for MLFeatureProviderConformer.
func GetMLFeatureProviderConformerClass() MLFeatureProviderConformerClass {
	return getMLFeatureProviderConformerClass()
}

type MLFeatureProviderConformerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLFeatureProviderConformerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLFeatureProviderConformerClass) Alloc() MLFeatureProviderConformer {
	rv := objc.Send[MLFeatureProviderConformer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLFeatureProviderConformer._fabricateFeatureForDescriptionError]
//   - [MLFeatureProviderConformer._sequenceConcatConsumesOptionalInputNamed]
//   - [MLFeatureProviderConformer.ConformFeaturesError]
//   - [MLFeatureProviderConformer.DefaultValues]
//   - [MLFeatureProviderConformer.FeatureDescriptions]
//   - [MLFeatureProviderConformer.OptionalInputTypes]
//   - [MLFeatureProviderConformer.PassthroughStateFeatures]
//   - [MLFeatureProviderConformer.UsingRank5Mapping]
//   - [MLFeatureProviderConformer.InitWithFeatureDescriptionsDefaultValuesUsingRank5MappingOptionalInputTypesPassthroughStateFeatures]
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderConformer
type MLFeatureProviderConformer struct {
	objectivec.Object
}

// MLFeatureProviderConformerFromID constructs a [MLFeatureProviderConformer] from an objc.ID.
func MLFeatureProviderConformerFromID(id objc.ID) MLFeatureProviderConformer {
	return MLFeatureProviderConformer{objectivec.Object{ID: id}}
}

// Ensure MLFeatureProviderConformer implements IMLFeatureProviderConformer.
var _ IMLFeatureProviderConformer = MLFeatureProviderConformer{}

// An interface definition for the [MLFeatureProviderConformer] class.
//
// # Methods
//
//   - [IMLFeatureProviderConformer._fabricateFeatureForDescriptionError]
//   - [IMLFeatureProviderConformer._sequenceConcatConsumesOptionalInputNamed]
//   - [IMLFeatureProviderConformer.ConformFeaturesError]
//   - [IMLFeatureProviderConformer.DefaultValues]
//   - [IMLFeatureProviderConformer.FeatureDescriptions]
//   - [IMLFeatureProviderConformer.OptionalInputTypes]
//   - [IMLFeatureProviderConformer.PassthroughStateFeatures]
//   - [IMLFeatureProviderConformer.UsingRank5Mapping]
//   - [IMLFeatureProviderConformer.InitWithFeatureDescriptionsDefaultValuesUsingRank5MappingOptionalInputTypesPassthroughStateFeatures]
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderConformer
type IMLFeatureProviderConformer interface {
	objectivec.IObject

	// Topic: Methods

	_fabricateFeatureForDescriptionError(description objectivec.IObject) (objectivec.IObject, error)
	_sequenceConcatConsumesOptionalInputNamed(named objectivec.IObject) bool
	ConformFeaturesError(features objectivec.IObject) (objectivec.IObject, error)
	DefaultValues() foundation.INSDictionary
	FeatureDescriptions() foundation.INSArray
	OptionalInputTypes() foundation.INSDictionary
	PassthroughStateFeatures() bool
	UsingRank5Mapping() bool
	InitWithFeatureDescriptionsDefaultValuesUsingRank5MappingOptionalInputTypesPassthroughStateFeatures(descriptions objectivec.IObject, values objectivec.IObject, rank5Mapping bool, types objectivec.IObject, features bool) MLFeatureProviderConformer
}

// Init initializes the instance.
func (f MLFeatureProviderConformer) Init() MLFeatureProviderConformer {
	rv := objc.Send[MLFeatureProviderConformer](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MLFeatureProviderConformer) Autorelease() MLFeatureProviderConformer {
	rv := objc.Send[MLFeatureProviderConformer](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLFeatureProviderConformer creates a new MLFeatureProviderConformer instance.
func NewMLFeatureProviderConformer() MLFeatureProviderConformer {
	class := getMLFeatureProviderConformerClass()
	rv := objc.Send[MLFeatureProviderConformer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderConformer/initWithFeatureDescriptions:defaultValues:usingRank5Mapping:optionalInputTypes:passthroughStateFeatures:
func NewFeatureProviderConformerWithFeatureDescriptionsDefaultValuesUsingRank5MappingOptionalInputTypesPassthroughStateFeatures(descriptions objectivec.IObject, values objectivec.IObject, rank5Mapping bool, types objectivec.IObject, features bool) MLFeatureProviderConformer {
	instance := getMLFeatureProviderConformerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFeatureDescriptions:defaultValues:usingRank5Mapping:optionalInputTypes:passthroughStateFeatures:"), descriptions, values, rank5Mapping, types, features)
	return MLFeatureProviderConformerFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderConformer/_fabricateFeatureForDescription:error:
func (f MLFeatureProviderConformer) _fabricateFeatureForDescriptionError(description objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("_fabricateFeatureForDescription:error:"), description, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// FabricateFeatureForDescriptionError is an exported wrapper for the private method _fabricateFeatureForDescriptionError.
func (f MLFeatureProviderConformer) FabricateFeatureForDescriptionError(description objectivec.IObject) (objectivec.IObject, error) {
	return f._fabricateFeatureForDescriptionError(description)
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderConformer/_sequenceConcatConsumesOptionalInputNamed:
func (f MLFeatureProviderConformer) _sequenceConcatConsumesOptionalInputNamed(named objectivec.IObject) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("_sequenceConcatConsumesOptionalInputNamed:"), named)
	return rv
}

// SequenceConcatConsumesOptionalInputNamed is an exported wrapper for the private method _sequenceConcatConsumesOptionalInputNamed.
func (f MLFeatureProviderConformer) SequenceConcatConsumesOptionalInputNamed(named objectivec.IObject) bool {
	return f._sequenceConcatConsumesOptionalInputNamed(named)
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderConformer/conformFeatures:error:
func (f MLFeatureProviderConformer) ConformFeaturesError(features objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("conformFeatures:error:"), features, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderConformer/initWithFeatureDescriptions:defaultValues:usingRank5Mapping:optionalInputTypes:passthroughStateFeatures:
func (f MLFeatureProviderConformer) InitWithFeatureDescriptionsDefaultValuesUsingRank5MappingOptionalInputTypesPassthroughStateFeatures(descriptions objectivec.IObject, values objectivec.IObject, rank5Mapping bool, types objectivec.IObject, features bool) MLFeatureProviderConformer {
	rv := objc.Send[MLFeatureProviderConformer](f.ID, objc.Sel("initWithFeatureDescriptions:defaultValues:usingRank5Mapping:optionalInputTypes:passthroughStateFeatures:"), descriptions, values, rank5Mapping, types, features)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderConformer/defaultValues
func (f MLFeatureProviderConformer) DefaultValues() foundation.INSDictionary {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("defaultValues"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderConformer/featureDescriptions
func (f MLFeatureProviderConformer) FeatureDescriptions() foundation.INSArray {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("featureDescriptions"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderConformer/optionalInputTypes
func (f MLFeatureProviderConformer) OptionalInputTypes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("optionalInputTypes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderConformer/passthroughStateFeatures
func (f MLFeatureProviderConformer) PassthroughStateFeatures() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("passthroughStateFeatures"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderConformer/usingRank5Mapping
func (f MLFeatureProviderConformer) UsingRank5Mapping() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("usingRank5Mapping"))
	return rv
}

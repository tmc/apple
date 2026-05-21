// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLScaler] class.
var (
	_MLScalerClass     MLScalerClass
	_MLScalerClassOnce sync.Once
)

func getMLScalerClass() MLScalerClass {
	_MLScalerClassOnce.Do(func() {
		_MLScalerClass = MLScalerClass{class: objc.GetClass("MLScaler")}
	})
	return _MLScalerClass
}

// GetMLScalerClass returns the class object for MLScaler.
func GetMLScalerClass() MLScalerClass {
	return getMLScalerClass()
}

type MLScalerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLScalerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLScalerClass) Alloc() MLScaler {
	rv := objc.Send[MLScaler](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLScaler.ScaleValue]
//   - [MLScaler.ShiftValue]
//   - [MLScaler.InitWithShiftValueScaleValueDescriptionConfiguration]
type MLScaler struct {
	MLModelEngine
}

// MLScalerFromID constructs a [MLScaler] from an objc.ID.
func MLScalerFromID(id objc.ID) MLScaler {
	return MLScaler{MLModelEngine: MLModelEngineFromID(id)}
}

// Ensure MLScaler implements IMLScaler.
var _ IMLScaler = MLScaler{}

// An interface definition for the [MLScaler] class.
//
// # Methods
//
//   - [IMLScaler.ScaleValue]
//   - [IMLScaler.ShiftValue]
//   - [IMLScaler.InitWithShiftValueScaleValueDescriptionConfiguration]
type IMLScaler interface {
	IMLModelEngine

	// Topic: Methods

	ScaleValue() IMLFeatureValue
	ShiftValue() IMLFeatureValue
	InitWithShiftValueScaleValueDescriptionConfiguration(value objectivec.IObject, value2 objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) MLScaler
}

// Init initializes the instance.
func (m MLScaler) Init() MLScaler {
	rv := objc.Send[MLScaler](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLScaler) Autorelease() MLScaler {
	rv := objc.Send[MLScaler](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLScaler creates a new MLScaler instance.
func NewMLScaler() MLScaler {
	class := getMLScalerClass()
	rv := objc.Send[MLScaler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewScalerWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLScaler {
	instance := getMLScalerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLScalerFromID(rv)
}

func NewScalerWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLScaler {
	instance := getMLScalerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLScalerFromID(rv)
}

func NewScalerWithShiftValueScaleValueDescriptionConfiguration(value objectivec.IObject, value2 objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) MLScaler {
	instance := getMLScalerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithShiftValue:scaleValue:description:configuration:"), value, value2, description, configuration)
	return MLScalerFromID(rv)
}

func (m MLScaler) InitWithShiftValueScaleValueDescriptionConfiguration(value objectivec.IObject, value2 objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) MLScaler {
	rv := objc.Send[MLScaler](m.ID, objc.Sel("initWithShiftValue:scaleValue:description:configuration:"), value, value2, description, configuration)
	return rv
}

func (_MLScalerClass MLScalerClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLScalerClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

func (m MLScaler) ScaleValue() IMLFeatureValue {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("scaleValue"))
	return MLFeatureValueFromID(objc.ID(rv))
}
func (m MLScaler) ShiftValue() IMLFeatureValue {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("shiftValue"))
	return MLFeatureValueFromID(objc.ID(rv))
}

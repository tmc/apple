// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelTypeRegistry] class.
var (
	_MLModelTypeRegistryClass     MLModelTypeRegistryClass
	_MLModelTypeRegistryClassOnce sync.Once
)

func getMLModelTypeRegistryClass() MLModelTypeRegistryClass {
	_MLModelTypeRegistryClassOnce.Do(func() {
		_MLModelTypeRegistryClass = MLModelTypeRegistryClass{class: objc.GetClass("MLModelTypeRegistry")}
	})
	return _MLModelTypeRegistryClass
}

// GetMLModelTypeRegistryClass returns the class object for MLModelTypeRegistry.
func GetMLModelTypeRegistryClass() MLModelTypeRegistryClass {
	return getMLModelTypeRegistryClass()
}

type MLModelTypeRegistryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelTypeRegistryClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelTypeRegistryClass) Alloc() MLModelTypeRegistry {
	rv := objc.Send[MLModelTypeRegistry](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLModelTypeRegistry.ClassForCompilingModelType]
//   - [MLModelTypeRegistry.ClassesForLoadingModelType]
//   - [MLModelTypeRegistry.ClassesForLoadingModelTypeConfigurationIsUpdatableIsEncrypted]
//   - [MLModelTypeRegistry.LoadNeuralNetworkClassesTrainWithMLCompute]
// See: https://developer.apple.com/documentation/CoreML/MLModelTypeRegistry
type MLModelTypeRegistry struct {
	objectivec.Object
}

// MLModelTypeRegistryFromID constructs a [MLModelTypeRegistry] from an objc.ID.
func MLModelTypeRegistryFromID(id objc.ID) MLModelTypeRegistry {
	return MLModelTypeRegistry{objectivec.Object{ID: id}}
}
// Ensure MLModelTypeRegistry implements IMLModelTypeRegistry.
var _ IMLModelTypeRegistry = MLModelTypeRegistry{}

// An interface definition for the [MLModelTypeRegistry] class.
//
// # Methods
//
//   - [IMLModelTypeRegistry.ClassForCompilingModelType]
//   - [IMLModelTypeRegistry.ClassesForLoadingModelType]
//   - [IMLModelTypeRegistry.ClassesForLoadingModelTypeConfigurationIsUpdatableIsEncrypted]
//   - [IMLModelTypeRegistry.LoadNeuralNetworkClassesTrainWithMLCompute]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelTypeRegistry
type IMLModelTypeRegistry interface {
	objectivec.IObject

	// Topic: Methods

	ClassForCompilingModelType(type_ int) objc.Class
	ClassesForLoadingModelType(type_ int) objectivec.IObject
	ClassesForLoadingModelTypeConfigurationIsUpdatableIsEncrypted(type_ int, configuration objectivec.IObject, updatable bool, encrypted bool) objectivec.IObject
	LoadNeuralNetworkClassesTrainWithMLCompute(classes bool, mLCompute bool) objectivec.IObject
}

// Init initializes the instance.
func (m MLModelTypeRegistry) Init() MLModelTypeRegistry {
	rv := objc.Send[MLModelTypeRegistry](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelTypeRegistry) Autorelease() MLModelTypeRegistry {
	rv := objc.Send[MLModelTypeRegistry](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelTypeRegistry creates a new MLModelTypeRegistry instance.
func NewMLModelTypeRegistry() MLModelTypeRegistry {
	class := getMLModelTypeRegistryClass()
	rv := objc.Send[MLModelTypeRegistry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelTypeRegistry/classForCompilingModelType:
func (m MLModelTypeRegistry) ClassForCompilingModelType(type_ int) objc.Class {
	rv := objc.Send[objc.Class](m.ID, objc.Sel("classForCompilingModelType:"), type_)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelTypeRegistry/classesForLoadingModelType:
func (m MLModelTypeRegistry) ClassesForLoadingModelType(type_ int) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classesForLoadingModelType:"), type_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelTypeRegistry/classesForLoadingModelType:configuration:isUpdatable:isEncrypted:
func (m MLModelTypeRegistry) ClassesForLoadingModelTypeConfigurationIsUpdatableIsEncrypted(type_ int, configuration objectivec.IObject, updatable bool, encrypted bool) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classesForLoadingModelType:configuration:isUpdatable:isEncrypted:"), type_, configuration, updatable, encrypted)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelTypeRegistry/loadNeuralNetworkClasses:trainWithMLCompute:
func (m MLModelTypeRegistry) LoadNeuralNetworkClassesTrainWithMLCompute(classes bool, mLCompute bool) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("loadNeuralNetworkClasses:trainWithMLCompute:"), classes, mLCompute)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLModelTypeRegistry/sharedInstance
func (_MLModelTypeRegistryClass MLModelTypeRegistryClass) SharedInstance() MLModelTypeRegistry {
	rv := objc.Send[objc.ID](objc.ID(_MLModelTypeRegistryClass.class), objc.Sel("sharedInstance"))
	return MLModelTypeRegistryFromID(rv)
}


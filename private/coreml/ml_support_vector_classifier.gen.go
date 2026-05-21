// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSupportVectorClassifier] class.
var (
	_MLSupportVectorClassifierClass     MLSupportVectorClassifierClass
	_MLSupportVectorClassifierClassOnce sync.Once
)

func getMLSupportVectorClassifierClass() MLSupportVectorClassifierClass {
	_MLSupportVectorClassifierClassOnce.Do(func() {
		_MLSupportVectorClassifierClass = MLSupportVectorClassifierClass{class: objc.GetClass("MLSupportVectorClassifier")}
	})
	return _MLSupportVectorClassifierClass
}

// GetMLSupportVectorClassifierClass returns the class object for MLSupportVectorClassifier.
func GetMLSupportVectorClassifierClass() MLSupportVectorClassifierClass {
	return getMLSupportVectorClassifierClass()
}

type MLSupportVectorClassifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSupportVectorClassifierClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSupportVectorClassifierClass) Alloc() MLSupportVectorClassifier {
	rv := objc.Send[MLSupportVectorClassifier](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLSupportVectorClassifier.ClassifyOptionsError]
//   - [MLSupportVectorClassifier.Engine]
//   - [MLSupportVectorClassifier.SetEngine]
//   - [MLSupportVectorClassifier.InitWithEngineDescriptionConfigurationError]
type MLSupportVectorClassifier struct {
	objectivec.Object
}

// MLSupportVectorClassifierFromID constructs a [MLSupportVectorClassifier] from an objc.ID.
func MLSupportVectorClassifierFromID(id objc.ID) MLSupportVectorClassifier {
	return MLSupportVectorClassifier{objectivec.Object{ID: id}}
}

// NOTE: MLSupportVectorClassifier struct embeds objectivec.Object (parent type unavailable) but
// IMLSupportVectorClassifier embeds the parent interface; skip compile-time assertion.

// An interface definition for the [MLSupportVectorClassifier] class.
//
// # Methods
//
//   - [IMLSupportVectorClassifier.ClassifyOptionsError]
//   - [IMLSupportVectorClassifier.Engine]
//   - [IMLSupportVectorClassifier.SetEngine]
//   - [IMLSupportVectorClassifier.InitWithEngineDescriptionConfigurationError]
type IMLSupportVectorClassifier interface {
	IMLClassifier

	// Topic: Methods

	ClassifyOptionsError(classify objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	Engine() IMLSVMEngine
	SetEngine(value IMLSVMEngine)
	InitWithEngineDescriptionConfigurationError(engine objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLSupportVectorClassifier, error)
}

// Init initializes the instance.
func (m MLSupportVectorClassifier) Init() MLSupportVectorClassifier {
	rv := objc.Send[MLSupportVectorClassifier](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLSupportVectorClassifier) Autorelease() MLSupportVectorClassifier {
	rv := objc.Send[MLSupportVectorClassifier](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSupportVectorClassifier creates a new MLSupportVectorClassifier instance.
func NewMLSupportVectorClassifier() MLSupportVectorClassifier {
	class := getMLSupportVectorClassifierClass()
	rv := objc.Send[MLSupportVectorClassifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSupportVectorClassifierWithEngineDescriptionConfigurationError(engine objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLSupportVectorClassifier, error) {
	var errorPtr objc.ID
	instance := getMLSupportVectorClassifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEngine:description:configuration:error:"), engine, description, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLSupportVectorClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLSupportVectorClassifierFromID(rv), nil
}

func (m MLSupportVectorClassifier) ClassifyOptionsError(classify objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classify:options:error:"), classify, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLSupportVectorClassifier) InitWithEngineDescriptionConfigurationError(engine objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLSupportVectorClassifier, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithEngine:description:configuration:error:"), engine, description, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLSupportVectorClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLSupportVectorClassifierFromID(rv), nil

}

func (_MLSupportVectorClassifierClass MLSupportVectorClassifierClass) FeatureValueWithObject(object objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLSupportVectorClassifierClass.class), objc.Sel("featureValueWithObject:"), object)
	return objectivec.Object{ID: rv}
}

func (m MLSupportVectorClassifier) Engine() IMLSVMEngine {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("engine"))
	return MLSVMEngineFromID(objc.ID(rv))
}
func (m MLSupportVectorClassifier) SetEngine(value IMLSVMEngine) {
	objc.Send[struct{}](m.ID, objc.Sel("setEngine:"), value)
}

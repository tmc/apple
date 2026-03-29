// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSupportVectorRegressor] class.
var (
	_MLSupportVectorRegressorClass     MLSupportVectorRegressorClass
	_MLSupportVectorRegressorClassOnce sync.Once
)

func getMLSupportVectorRegressorClass() MLSupportVectorRegressorClass {
	_MLSupportVectorRegressorClassOnce.Do(func() {
		_MLSupportVectorRegressorClass = MLSupportVectorRegressorClass{class: objc.GetClass("MLSupportVectorRegressor")}
	})
	return _MLSupportVectorRegressorClass
}

// GetMLSupportVectorRegressorClass returns the class object for MLSupportVectorRegressor.
func GetMLSupportVectorRegressorClass() MLSupportVectorRegressorClass {
	return getMLSupportVectorRegressorClass()
}

type MLSupportVectorRegressorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSupportVectorRegressorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSupportVectorRegressorClass) Alloc() MLSupportVectorRegressor {
	rv := objc.Send[MLSupportVectorRegressor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLSupportVectorRegressor.Engine]
//   - [MLSupportVectorRegressor.SetEngine]
//   - [MLSupportVectorRegressor.RegressOptionsError]
//   - [MLSupportVectorRegressor.InitWithEngineDescriptionConfigurationError]
// See: https://developer.apple.com/documentation/CoreML/MLSupportVectorRegressor
type MLSupportVectorRegressor struct {
	objectivec.Object
}

// MLSupportVectorRegressorFromID constructs a [MLSupportVectorRegressor] from an objc.ID.
func MLSupportVectorRegressorFromID(id objc.ID) MLSupportVectorRegressor {
	return MLSupportVectorRegressor{objectivec.Object{ID: id}}
}
// NOTE: MLSupportVectorRegressor struct embeds objectivec.Object (parent type unavailable) but
// IMLSupportVectorRegressor embeds the parent interface; skip compile-time assertion.

// An interface definition for the [MLSupportVectorRegressor] class.
//
// # Methods
//
//   - [IMLSupportVectorRegressor.Engine]
//   - [IMLSupportVectorRegressor.SetEngine]
//   - [IMLSupportVectorRegressor.RegressOptionsError]
//   - [IMLSupportVectorRegressor.InitWithEngineDescriptionConfigurationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLSupportVectorRegressor
type IMLSupportVectorRegressor interface {
	IMLRegressor

	// Topic: Methods

	Engine() IMLSVREngine
	SetEngine(value IMLSVREngine)
	RegressOptionsError(regress objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithEngineDescriptionConfigurationError(engine objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLSupportVectorRegressor, error)
}

// Init initializes the instance.
func (s MLSupportVectorRegressor) Init() MLSupportVectorRegressor {
	rv := objc.Send[MLSupportVectorRegressor](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MLSupportVectorRegressor) Autorelease() MLSupportVectorRegressor {
	rv := objc.Send[MLSupportVectorRegressor](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSupportVectorRegressor creates a new MLSupportVectorRegressor instance.
func NewMLSupportVectorRegressor() MLSupportVectorRegressor {
	class := getMLSupportVectorRegressorClass()
	rv := objc.Send[MLSupportVectorRegressor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLSupportVectorRegressor/initWithEngine:description:configuration:error:
func NewSupportVectorRegressorWithEngineDescriptionConfigurationError(engine objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLSupportVectorRegressor, error) {
	var errorPtr objc.ID
	instance := getMLSupportVectorRegressorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEngine:description:configuration:error:"), engine, description, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLSupportVectorRegressor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLSupportVectorRegressorFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLSupportVectorRegressor/regress:options:error:
func (s MLSupportVectorRegressor) RegressOptionsError(regress objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("regress:options:error:"), regress, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLSupportVectorRegressor/initWithEngine:description:configuration:error:
func (s MLSupportVectorRegressor) InitWithEngineDescriptionConfigurationError(engine objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLSupportVectorRegressor, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithEngine:description:configuration:error:"), engine, description, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLSupportVectorRegressor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLSupportVectorRegressorFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLSupportVectorRegressor/engine
func (s MLSupportVectorRegressor) Engine() IMLSVREngine {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("engine"))
	return MLSVREngineFromID(objc.ID(rv))
}
func (s MLSupportVectorRegressor) SetEngine(value IMLSVREngine) {
	objc.Send[struct{}](s.ID, objc.Sel("setEngine:"), value)
}


// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLPipelineRegressor] class.
var (
	_MLPipelineRegressorClass     MLPipelineRegressorClass
	_MLPipelineRegressorClassOnce sync.Once
)

func getMLPipelineRegressorClass() MLPipelineRegressorClass {
	_MLPipelineRegressorClassOnce.Do(func() {
		_MLPipelineRegressorClass = MLPipelineRegressorClass{class: objc.GetClass("MLPipelineRegressor")}
	})
	return _MLPipelineRegressorClass
}

// GetMLPipelineRegressorClass returns the class object for MLPipelineRegressor.
func GetMLPipelineRegressorClass() MLPipelineRegressorClass {
	return getMLPipelineRegressorClass()
}

type MLPipelineRegressorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLPipelineRegressorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLPipelineRegressorClass) Alloc() MLPipelineRegressor {
	rv := objc.Send[MLPipelineRegressor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLPipelineRegressor.Engine]
//   - [MLPipelineRegressor.SetEngine]
//   - [MLPipelineRegressor.ExecutionSchedule]
//   - [MLPipelineRegressor.Pipeline]
//   - [MLPipelineRegressor.RegressOptionsError]
//   - [MLPipelineRegressor.SignpostID]
//   - [MLPipelineRegressor.InitWithEngineDescriptionConfigurationError]
// See: https://developer.apple.com/documentation/CoreML/MLPipelineRegressor
type MLPipelineRegressor struct {
	objectivec.Object
}

// MLPipelineRegressorFromID constructs a [MLPipelineRegressor] from an objc.ID.
func MLPipelineRegressorFromID(id objc.ID) MLPipelineRegressor {
	return MLPipelineRegressor{objectivec.Object{ID: id}}
}
// NOTE: MLPipelineRegressor struct embeds objectivec.Object (parent type unavailable) but
// IMLPipelineRegressor embeds the parent interface; skip compile-time assertion.

// An interface definition for the [MLPipelineRegressor] class.
//
// # Methods
//
//   - [IMLPipelineRegressor.Engine]
//   - [IMLPipelineRegressor.SetEngine]
//   - [IMLPipelineRegressor.ExecutionSchedule]
//   - [IMLPipelineRegressor.Pipeline]
//   - [IMLPipelineRegressor.RegressOptionsError]
//   - [IMLPipelineRegressor.SignpostID]
//   - [IMLPipelineRegressor.InitWithEngineDescriptionConfigurationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineRegressor
type IMLPipelineRegressor interface {
	IMLRegressor

	// Topic: Methods

	Engine() unsafe.Pointer
	SetEngine(value *MLPipeline)
	ExecutionSchedule() objectivec.IObject
	Pipeline() unsafe.Pointer
	RegressOptionsError(regress objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	SignpostID() uint64
	InitWithEngineDescriptionConfigurationError(engine objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLPipelineRegressor, error)
}

// Init initializes the instance.
func (p MLPipelineRegressor) Init() MLPipelineRegressor {
	rv := objc.Send[MLPipelineRegressor](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLPipelineRegressor) Autorelease() MLPipelineRegressor {
	rv := objc.Send[MLPipelineRegressor](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPipelineRegressor creates a new MLPipelineRegressor instance.
func NewMLPipelineRegressor() MLPipelineRegressor {
	class := getMLPipelineRegressorClass()
	rv := objc.Send[MLPipelineRegressor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineRegressor/initWithEngine:description:configuration:error:
func NewPipelineRegressorWithEngineDescriptionConfigurationError(engine objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLPipelineRegressor, error) {
	var errorPtr objc.ID
	instance := getMLPipelineRegressorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEngine:description:configuration:error:"), engine, description, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLPipelineRegressor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLPipelineRegressorFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineRegressor/executionSchedule
func (p MLPipelineRegressor) ExecutionSchedule() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("executionSchedule"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineRegressor/regress:options:error:
func (p MLPipelineRegressor) RegressOptionsError(regress objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](p.ID, objc.Sel("regress:options:error:"), regress, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/CoreML/MLPipelineRegressor/signpostID
func (p MLPipelineRegressor) SignpostID() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("signpostID"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineRegressor/initWithEngine:description:configuration:error:
func (p MLPipelineRegressor) InitWithEngineDescriptionConfigurationError(engine objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLPipelineRegressor, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](p.ID, objc.Sel("initWithEngine:description:configuration:error:"), engine, description, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLPipelineRegressor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLPipelineRegressorFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineRegressor/engine
func (p MLPipelineRegressor) Engine() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p.ID, objc.Sel("engine"))
	return rv
}
func (p MLPipelineRegressor) SetEngine(value *MLPipeline) {
	objc.Send[struct{}](p.ID, objc.Sel("setEngine:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLPipelineRegressor/pipeline
func (p MLPipelineRegressor) Pipeline() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p.ID, objc.Sel("pipeline"))
	return rv
}


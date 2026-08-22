// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[MLPipelineRegressor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLPipelineRegressor.Engine]
//   - [MLPipelineRegressor.SetEngine]
//   - [MLPipelineRegressor.ExecutionSchedule]
//   - [MLPipelineRegressor.Pipeline]
//   - [MLPipelineRegressor.RegressOptionsError]
//   - [MLPipelineRegressor.SignpostID]
//   - [MLPipelineRegressor.InitWithEngineDescriptionConfigurationError]
type MLPipelineRegressor struct {
	objectivec.Object
}

// MLPipelineRegressorFromID constructs a [MLPipelineRegressor] from an objc.ID.
func MLPipelineRegressorFromID(id objc.ID) MLPipelineRegressor {
	return MLPipelineRegressor{objectivec.Object{ID: id}}
}

// Ensure MLPipelineRegressor implements IMLPipelineRegressor.
var _ IMLPipelineRegressor = MLPipelineRegressor{}

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
type IMLPipelineRegressor interface {
	objectivec.IObject

	// Topic: Methods

	Engine() MLPipeline
	SetEngine(value MLPipeline)
	ExecutionSchedule() objectivec.IObject
	Pipeline() MLPipeline
	RegressOptionsError(regress objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	SignpostID() uint64
	InitWithEngineDescriptionConfigurationError(engine objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLPipelineRegressor, error)
}

// Init initializes the instance.
func (m MLPipelineRegressor) Init() MLPipelineRegressor {
	rv := objc.SendIfResponds[MLPipelineRegressor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLPipelineRegressor) Autorelease() MLPipelineRegressor {
	rv := objc.SendIfResponds[MLPipelineRegressor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPipelineRegressor creates a new MLPipelineRegressor instance.
func NewMLPipelineRegressor() MLPipelineRegressor {
	class := getMLPipelineRegressorClass()
	rv := objc.SendIfResponds[MLPipelineRegressor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewPipelineRegressorWithEngineDescriptionConfigurationError(engine objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLPipelineRegressor, error) {
	var errorPtr objc.ID
	instance := getMLPipelineRegressorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithEngine:description:configuration:error:"), engine, description, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLPipelineRegressor{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLPipelineRegressor{}, objc.ErrInitFailed
	}
	return MLPipelineRegressorFromID(rv), nil
}

func (m MLPipelineRegressor) ExecutionSchedule() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("executionSchedule"))
	return objectivec.Object{ID: rv}
}
func (m MLPipelineRegressor) RegressOptionsError(regress objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("regress:options:error:"), regress, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLPipelineRegressor) SignpostID() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("signpostID"))
	return rv
}
func (m MLPipelineRegressor) InitWithEngineDescriptionConfigurationError(engine objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLPipelineRegressor, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithEngine:description:configuration:error:"), engine, description, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLPipelineRegressor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLPipelineRegressorFromID(rv), nil

}

func (m MLPipelineRegressor) Engine() MLPipeline {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("engine"))
	return MLPipelineObjectFromID(rv)
}
func (m MLPipelineRegressor) SetEngine(value MLPipeline) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setEngine:"), value)
}
func (m MLPipelineRegressor) Pipeline() MLPipeline {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("pipeline"))
	return MLPipelineObjectFromID(rv)
}

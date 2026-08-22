// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLUpdateContext] class.
var (
	_MLUpdateContextClass     MLUpdateContextClass
	_MLUpdateContextClassOnce sync.Once
)

func getMLUpdateContextClass() MLUpdateContextClass {
	_MLUpdateContextClassOnce.Do(func() {
		_MLUpdateContextClass = MLUpdateContextClass{class: objc.GetClass("MLUpdateContext")}
	})
	return _MLUpdateContextClass
}

// GetMLUpdateContextClass returns the class object for MLUpdateContext.
func GetMLUpdateContextClass() MLUpdateContextClass {
	return getMLUpdateContextClass()
}

type MLUpdateContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLUpdateContextClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLUpdateContextClass) Alloc() MLUpdateContext {
	rv := objc.SendIfResponds[MLUpdateContext](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLUpdateContext.Error]
//   - [MLUpdateContext.SetError]
//   - [MLUpdateContext.Event]
//   - [MLUpdateContext.SetEvent]
//   - [MLUpdateContext.Metrics]
//   - [MLUpdateContext.SetMetrics]
//   - [MLUpdateContext.Model]
//   - [MLUpdateContext.SetModel]
//   - [MLUpdateContext.Parameters]
//   - [MLUpdateContext.SetParameters]
//   - [MLUpdateContext.Task]
//   - [MLUpdateContext.SetTask]
type MLUpdateContext struct {
	objectivec.Object
}

// MLUpdateContextFromID constructs a [MLUpdateContext] from an objc.ID.
func MLUpdateContextFromID(id objc.ID) MLUpdateContext {
	return MLUpdateContext{objectivec.Object{ID: id}}
}

// Ensure MLUpdateContext implements IMLUpdateContext.
var _ IMLUpdateContext = MLUpdateContext{}

// An interface definition for the [MLUpdateContext] class.
//
// # Methods
//
//   - [IMLUpdateContext.Error]
//   - [IMLUpdateContext.SetError]
//   - [IMLUpdateContext.Event]
//   - [IMLUpdateContext.SetEvent]
//   - [IMLUpdateContext.Metrics]
//   - [IMLUpdateContext.SetMetrics]
//   - [IMLUpdateContext.Model]
//   - [IMLUpdateContext.SetModel]
//   - [IMLUpdateContext.Parameters]
//   - [IMLUpdateContext.SetParameters]
//   - [IMLUpdateContext.Task]
//   - [IMLUpdateContext.SetTask]
type IMLUpdateContext interface {
	objectivec.IObject

	// Topic: Methods

	Error() foundation.NSError
	SetError(value foundation.NSError)
	Event() int64
	SetEvent(value int64)
	Metrics() foundation.INSDictionary
	SetMetrics(value foundation.INSDictionary)
	Model() IMLModel
	SetModel(value IMLModel)
	Parameters() foundation.INSDictionary
	SetParameters(value foundation.INSDictionary)
	Task() IMLUpdateTask
	SetTask(value IMLUpdateTask)
}

// Init initializes the instance.
func (m MLUpdateContext) Init() MLUpdateContext {
	rv := objc.SendIfResponds[MLUpdateContext](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLUpdateContext) Autorelease() MLUpdateContext {
	rv := objc.SendIfResponds[MLUpdateContext](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLUpdateContext creates a new MLUpdateContext instance.
func NewMLUpdateContext() MLUpdateContext {
	class := getMLUpdateContextClass()
	rv := objc.SendIfResponds[MLUpdateContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_MLUpdateContextClass MLUpdateContextClass) UpdateContextForEventMetricsParametersError(event int64, metrics objectivec.IObject, parameters objectivec.IObject, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLUpdateContextClass.class), objc.Sel("updateContextForEvent:metrics:parameters:error:"), event, metrics, parameters, error_)
	return objectivec.Object{ID: rv}
}
func (_MLUpdateContextClass MLUpdateContextClass) UpdateContextWithTaskModelEventMetricsParameters(task objectivec.IObject, model objectivec.IObject, event int64, metrics objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLUpdateContextClass.class), objc.Sel("updateContextWithTask:model:event:metrics:parameters:"), task, model, event, metrics, parameters)
	return objectivec.Object{ID: rv}
}

func (m MLUpdateContext) Error() foundation.NSError {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
func (m MLUpdateContext) SetError(value foundation.NSError) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setError:"), value)
}
func (m MLUpdateContext) Event() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("event"))
	return rv
}
func (m MLUpdateContext) SetEvent(value int64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setEvent:"), value)
}
func (m MLUpdateContext) Metrics() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("metrics"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLUpdateContext) SetMetrics(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setMetrics:"), value)
}
func (m MLUpdateContext) Model() IMLModel {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("model"))
	return MLModelFromID(objc.ID(rv))
}
func (m MLUpdateContext) SetModel(value IMLModel) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setModel:"), value)
}
func (m MLUpdateContext) Parameters() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("parameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLUpdateContext) SetParameters(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setParameters:"), value)
}
func (m MLUpdateContext) Task() IMLUpdateTask {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("task"))
	return MLUpdateTaskFromID(objc.ID(rv))
}
func (m MLUpdateContext) SetTask(value IMLUpdateTask) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setTask:"), value)
}

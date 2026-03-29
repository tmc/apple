// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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
	rv := objc.Send[MLUpdateContext](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
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
// See: https://developer.apple.com/documentation/CoreML/MLUpdateContext
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
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateContext
type IMLUpdateContext interface {
	objectivec.IObject

	// Topic: Methods

	Error() foundation.INSError
	SetError(value foundation.INSError)
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
func (u MLUpdateContext) Init() MLUpdateContext {
	rv := objc.Send[MLUpdateContext](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u MLUpdateContext) Autorelease() MLUpdateContext {
	rv := objc.Send[MLUpdateContext](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLUpdateContext creates a new MLUpdateContext instance.
func NewMLUpdateContext() MLUpdateContext {
	class := getMLUpdateContextClass()
	rv := objc.Send[MLUpdateContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateContext/updateContextForEvent:metrics:parameters:error:
func (_MLUpdateContextClass MLUpdateContextClass) UpdateContextForEventMetricsParametersError(event int64, metrics objectivec.IObject, parameters objectivec.IObject, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLUpdateContextClass.class), objc.Sel("updateContextForEvent:metrics:parameters:error:"), event, metrics, parameters, error_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateContext/updateContextWithTask:model:event:metrics:parameters:
func (_MLUpdateContextClass MLUpdateContextClass) UpdateContextWithTaskModelEventMetricsParameters(task objectivec.IObject, model objectivec.IObject, event int64, metrics objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLUpdateContextClass.class), objc.Sel("updateContextWithTask:model:event:metrics:parameters:"), task, model, event, metrics, parameters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLUpdateContext/error
func (u MLUpdateContext) Error() foundation.INSError {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
func (u MLUpdateContext) SetError(value foundation.INSError) {
	objc.Send[struct{}](u.ID, objc.Sel("setError:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLUpdateContext/event
func (u MLUpdateContext) Event() int64 {
	rv := objc.Send[int64](u.ID, objc.Sel("event"))
	return rv
}
func (u MLUpdateContext) SetEvent(value int64) {
	objc.Send[struct{}](u.ID, objc.Sel("setEvent:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLUpdateContext/metrics
func (u MLUpdateContext) Metrics() foundation.INSDictionary {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("metrics"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (u MLUpdateContext) SetMetrics(value foundation.INSDictionary) {
	objc.Send[struct{}](u.ID, objc.Sel("setMetrics:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLUpdateContext/model
func (u MLUpdateContext) Model() IMLModel {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("model"))
	return MLModelFromID(objc.ID(rv))
}
func (u MLUpdateContext) SetModel(value IMLModel) {
	objc.Send[struct{}](u.ID, objc.Sel("setModel:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLUpdateContext/parameters
func (u MLUpdateContext) Parameters() foundation.INSDictionary {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("parameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (u MLUpdateContext) SetParameters(value foundation.INSDictionary) {
	objc.Send[struct{}](u.ID, objc.Sel("setParameters:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLUpdateContext/task
func (u MLUpdateContext) Task() IMLUpdateTask {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("task"))
	return MLUpdateTaskFromID(objc.ID(rv))
}
func (u MLUpdateContext) SetTask(value IMLUpdateTask) {
	objc.Send[struct{}](u.ID, objc.Sel("setTask:"), value)
}


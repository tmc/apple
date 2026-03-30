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
	rv := objc.Send[MLUpdateContext](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The context an update task provides to your app’s completion and update
// progress handlers.
//
// # Getting the update context
//
//   - [MLUpdateContext.Event]: The event type that triggered an update task to notify your app’s completion and update progress handlers.
//   - [MLUpdateContext.Task]: The update task that generated the update context.
//   - [MLUpdateContext.Parameters]: The parameters for the update task.
//
// # Evaluating the update
//
//   - [MLUpdateContext.Metrics]: The training metrics of the model for the update task, contained in a dictionary.
//
// # Saving an updated model
//
//   - [MLUpdateContext.Model]: The underlying Core ML model stored in memory.
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateContext
type MLUpdateContext struct {
	objectivec.Object
}

// MLUpdateContextFromID constructs a [MLUpdateContext] from an objc.ID.
//
// The context an update task provides to your app’s completion and update
// progress handlers.
func MLUpdateContextFromID(id objc.ID) MLUpdateContext {
	return MLUpdateContext{objectivec.Object{ID: id}}
}

// NOTE: MLUpdateContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLUpdateContext] class.
//
// # Getting the update context
//
//   - [IMLUpdateContext.Event]: The event type that triggered an update task to notify your app’s completion and update progress handlers.
//   - [IMLUpdateContext.Task]: The update task that generated the update context.
//   - [IMLUpdateContext.Parameters]: The parameters for the update task.
//
// # Evaluating the update
//
//   - [IMLUpdateContext.Metrics]: The training metrics of the model for the update task, contained in a dictionary.
//
// # Saving an updated model
//
//   - [IMLUpdateContext.Model]: The underlying Core ML model stored in memory.
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateContext
type IMLUpdateContext interface {
	objectivec.IObject

	// Topic: Getting the update context

	// The event type that triggered an update task to notify your app’s completion and update progress handlers.
	Event() MLUpdateProgressEvent
	// The update task that generated the update context.
	Task() IMLUpdateTask
	// The parameters for the update task.
	Parameters() foundation.INSDictionary

	// Topic: Evaluating the update

	// The training metrics of the model for the update task, contained in a dictionary.
	Metrics() foundation.INSDictionary

	// Topic: Saving an updated model

	// The underlying Core ML model stored in memory.
	Model() IMLModel
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

// The event type that triggered an update task to notify your app’s
// completion and update progress handlers.
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateContext/event
func (u MLUpdateContext) Event() MLUpdateProgressEvent {
	rv := objc.Send[MLUpdateProgressEvent](u.ID, objc.Sel("event"))
	return MLUpdateProgressEvent(rv)
}

// The update task that generated the update context.
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateContext/task
func (u MLUpdateContext) Task() IMLUpdateTask {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("task"))
	return MLUpdateTaskFromID(objc.ID(rv))
}

// The parameters for the update task.
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateContext/parameters
func (u MLUpdateContext) Parameters() foundation.INSDictionary {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("parameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The training metrics of the model for the update task, contained in a
// dictionary.
//
// # Discussion
//
// Use the [MLMetricKey] to access the values within the dictionary.
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateContext/metrics
func (u MLUpdateContext) Metrics() foundation.INSDictionary {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("metrics"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The underlying Core ML model stored in memory.
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateContext/model
func (u MLUpdateContext) Model() IMLModel {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("model"))
	return MLModelFromID(objc.ID(rv))
}

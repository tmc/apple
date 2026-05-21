// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLBackgroundPredictionTask] class.
var (
	_MLBackgroundPredictionTaskClass     MLBackgroundPredictionTaskClass
	_MLBackgroundPredictionTaskClassOnce sync.Once
)

func getMLBackgroundPredictionTaskClass() MLBackgroundPredictionTaskClass {
	_MLBackgroundPredictionTaskClassOnce.Do(func() {
		_MLBackgroundPredictionTaskClass = MLBackgroundPredictionTaskClass{class: objc.GetClass("MLBackgroundPredictionTask")}
	})
	return _MLBackgroundPredictionTaskClass
}

// GetMLBackgroundPredictionTaskClass returns the class object for MLBackgroundPredictionTask.
func GetMLBackgroundPredictionTaskClass() MLBackgroundPredictionTaskClass {
	return getMLBackgroundPredictionTaskClass()
}

type MLBackgroundPredictionTaskClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLBackgroundPredictionTaskClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLBackgroundPredictionTaskClass) Alloc() MLBackgroundPredictionTask {
	rv := objc.Send[MLBackgroundPredictionTask](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLBackgroundPredictionTask.ModelConfiguration]
//   - [MLBackgroundPredictionTask.SetModelConfiguration]
//   - [MLBackgroundPredictionTask.ModelURL]
//   - [MLBackgroundPredictionTask.SetModelURL]
//   - [MLBackgroundPredictionTask.PredictionOptions]
//   - [MLBackgroundPredictionTask.SetPredictionOptions]
type MLBackgroundPredictionTask struct {
	MLBackgroundTask
}

// MLBackgroundPredictionTaskFromID constructs a [MLBackgroundPredictionTask] from an objc.ID.
func MLBackgroundPredictionTaskFromID(id objc.ID) MLBackgroundPredictionTask {
	return MLBackgroundPredictionTask{MLBackgroundTask: MLBackgroundTaskFromID(id)}
}

// Ensure MLBackgroundPredictionTask implements IMLBackgroundPredictionTask.
var _ IMLBackgroundPredictionTask = MLBackgroundPredictionTask{}

// An interface definition for the [MLBackgroundPredictionTask] class.
//
// # Methods
//
//   - [IMLBackgroundPredictionTask.ModelConfiguration]
//   - [IMLBackgroundPredictionTask.SetModelConfiguration]
//   - [IMLBackgroundPredictionTask.ModelURL]
//   - [IMLBackgroundPredictionTask.SetModelURL]
//   - [IMLBackgroundPredictionTask.PredictionOptions]
//   - [IMLBackgroundPredictionTask.SetPredictionOptions]
type IMLBackgroundPredictionTask interface {
	IMLBackgroundTask

	// Topic: Methods

	ModelConfiguration() IMLModelConfiguration
	SetModelConfiguration(value IMLModelConfiguration)
	ModelURL() foundation.NSURL
	SetModelURL(value foundation.NSURL)
	PredictionOptions() IMLPredictionOptions
	SetPredictionOptions(value IMLPredictionOptions)
}

// Init initializes the instance.
func (m MLBackgroundPredictionTask) Init() MLBackgroundPredictionTask {
	rv := objc.Send[MLBackgroundPredictionTask](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLBackgroundPredictionTask) Autorelease() MLBackgroundPredictionTask {
	rv := objc.Send[MLBackgroundPredictionTask](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLBackgroundPredictionTask creates a new MLBackgroundPredictionTask instance.
func NewMLBackgroundPredictionTask() MLBackgroundPredictionTask {
	class := getMLBackgroundPredictionTaskClass()
	rv := objc.Send[MLBackgroundPredictionTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewBackgroundPredictionTaskWithCoder(coder objectivec.IObject) MLBackgroundPredictionTask {
	instance := getMLBackgroundPredictionTaskClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLBackgroundPredictionTaskFromID(rv)
}

func (_MLBackgroundPredictionTaskClass MLBackgroundPredictionTaskClass) TaskRunnerClass() objectivec.Class {
	rv := objc.Send[objectivec.Class](objc.ID(_MLBackgroundPredictionTaskClass.class), objc.Sel("taskRunnerClass"))
	return objectivec.Class(rv)
}

func (m MLBackgroundPredictionTask) ModelConfiguration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelConfiguration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
func (m MLBackgroundPredictionTask) SetModelConfiguration(value IMLModelConfiguration) {
	objc.Send[struct{}](m.ID, objc.Sel("setModelConfiguration:"), value)
}
func (m MLBackgroundPredictionTask) ModelURL() foundation.NSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (m MLBackgroundPredictionTask) SetModelURL(value foundation.NSURL) {
	objc.Send[struct{}](m.ID, objc.Sel("setModelURL:"), value)
}
func (m MLBackgroundPredictionTask) PredictionOptions() IMLPredictionOptions {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionOptions"))
	return MLPredictionOptionsFromID(objc.ID(rv))
}
func (m MLBackgroundPredictionTask) SetPredictionOptions(value IMLPredictionOptions) {
	objc.Send[struct{}](m.ID, objc.Sel("setPredictionOptions:"), value)
}

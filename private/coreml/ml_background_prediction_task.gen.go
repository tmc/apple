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
//
// See: https://developer.apple.com/documentation/CoreML/MLBackgroundPredictionTask
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
//
// See: https://developer.apple.com/documentation/CoreML/MLBackgroundPredictionTask
type IMLBackgroundPredictionTask interface {
	IMLBackgroundTask

	// Topic: Methods

	ModelConfiguration() IMLModelConfiguration
	SetModelConfiguration(value IMLModelConfiguration)
	ModelURL() foundation.INSURL
	SetModelURL(value foundation.INSURL)
	PredictionOptions() IMLPredictionOptions
	SetPredictionOptions(value IMLPredictionOptions)
}

// Init initializes the instance.
func (b MLBackgroundPredictionTask) Init() MLBackgroundPredictionTask {
	rv := objc.Send[MLBackgroundPredictionTask](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MLBackgroundPredictionTask) Autorelease() MLBackgroundPredictionTask {
	rv := objc.Send[MLBackgroundPredictionTask](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLBackgroundPredictionTask creates a new MLBackgroundPredictionTask instance.
func NewMLBackgroundPredictionTask() MLBackgroundPredictionTask {
	class := getMLBackgroundPredictionTaskClass()
	rv := objc.Send[MLBackgroundPredictionTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundPredictionTask/initWithCoder:
func NewBackgroundPredictionTaskWithCoder(coder objectivec.IObject) MLBackgroundPredictionTask {
	instance := getMLBackgroundPredictionTaskClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLBackgroundPredictionTaskFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundPredictionTask/taskRunnerClass
func (_MLBackgroundPredictionTaskClass MLBackgroundPredictionTaskClass) TaskRunnerClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_MLBackgroundPredictionTaskClass.class), objc.Sel("taskRunnerClass"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundPredictionTask/modelConfiguration
func (b MLBackgroundPredictionTask) ModelConfiguration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("modelConfiguration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
func (b MLBackgroundPredictionTask) SetModelConfiguration(value IMLModelConfiguration) {
	objc.Send[struct{}](b.ID, objc.Sel("setModelConfiguration:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundPredictionTask/modelURL
func (b MLBackgroundPredictionTask) ModelURL() foundation.INSURL {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("modelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (b MLBackgroundPredictionTask) SetModelURL(value foundation.INSURL) {
	objc.Send[struct{}](b.ID, objc.Sel("setModelURL:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundPredictionTask/predictionOptions
func (b MLBackgroundPredictionTask) PredictionOptions() IMLPredictionOptions {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("predictionOptions"))
	return MLPredictionOptionsFromID(objc.ID(rv))
}
func (b MLBackgroundPredictionTask) SetPredictionOptions(value IMLPredictionOptions) {
	objc.Send[struct{}](b.ID, objc.Sel("setPredictionOptions:"), value)
}

// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [MLUpdateTask] class.
var (
	_MLUpdateTaskClass     MLUpdateTaskClass
	_MLUpdateTaskClassOnce sync.Once
)

func getMLUpdateTaskClass() MLUpdateTaskClass {
	_MLUpdateTaskClassOnce.Do(func() {
		_MLUpdateTaskClass = MLUpdateTaskClass{class: objc.GetClass("MLUpdateTask")}
	})
	return _MLUpdateTaskClass
}

// GetMLUpdateTaskClass returns the class object for MLUpdateTask.
func GetMLUpdateTaskClass() MLUpdateTaskClass {
	return getMLUpdateTaskClass()
}

type MLUpdateTaskClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLUpdateTaskClass) Alloc() MLUpdateTask {
	rv := objc.Send[MLUpdateTask](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A task that updates a model with additional training data.
//
// # Overview
// 
// Use an [MLUpdateTask] to update a machine learning model on a user’s
// device.
//
// # Starting and Resuming an Update
//
//   - [MLUpdateTask.ResumeWithParameters]: Resumes a model update with updated parameter values.
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateTask
type MLUpdateTask struct {
	MLTask
}

// MLUpdateTaskFromID constructs a [MLUpdateTask] from an objc.ID.
//
// A task that updates a model with additional training data.
func MLUpdateTaskFromID(id objc.ID) MLUpdateTask {
	return MLUpdateTask{MLTask: MLTaskFromID(id)}
}
// NOTE: MLUpdateTask adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLUpdateTask] class.
//
// # Starting and Resuming an Update
//
//   - [IMLUpdateTask.ResumeWithParameters]: Resumes a model update with updated parameter values.
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateTask
type IMLUpdateTask interface {
	IMLTask

	// Topic: Starting and Resuming an Update

	// Resumes a model update with updated parameter values.
	ResumeWithParameters(updateParameters foundation.INSDictionary)
}

// Init initializes the instance.
func (u MLUpdateTask) Init() MLUpdateTask {
	rv := objc.Send[MLUpdateTask](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u MLUpdateTask) Autorelease() MLUpdateTask {
	rv := objc.Send[MLUpdateTask](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLUpdateTask creates a new MLUpdateTask instance.
func NewMLUpdateTask() MLUpdateTask {
	class := getMLUpdateTaskClass()
	rv := objc.Send[MLUpdateTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a task that updates the model at the URL with the training data and
// configuration, and calls the progress handlers during and after the update.
//
// modelURL: The location in the file system of a model file
// (`XCUIElementTypeMlmodelc`).
//
// trainingData: The update data for the model, contained in a batch provider.
//
// configuration: The model settings for an updated model object.
//
// progressHandlers: The closures the task calls during the update process.
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:configuration:progressHandlers:)
func NewUpdateTaskForModelAtURLTrainingDataConfigurationProgressHandlersError(modelURL foundation.INSURL, trainingData MLBatchProvider, configuration IMLModelConfiguration, progressHandlers IMLUpdateProgressHandlers) (MLUpdateTask, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getMLUpdateTaskClass().class), objc.Sel("updateTaskForModelAtURL:trainingData:configuration:progressHandlers:error:"), modelURL, trainingData, configuration, progressHandlers, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLUpdateTaskFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return MLUpdateTaskFromID(rv), nil
}

// Creates a task that updates the model at the URL with the training data,
// and calls the progress handlers during and after the update.
//
// modelURL: The location in the file system of a model file
// (`XCUIElementTypeMlmodelc`).
//
// trainingData: The update data for the model, contained in a batch provider.
//
// progressHandlers: The closures the task calls during the update process.
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:progressHandlers:)
func NewUpdateTaskForModelAtURLTrainingDataProgressHandlersError(modelURL foundation.INSURL, trainingData MLBatchProvider, progressHandlers IMLUpdateProgressHandlers) (MLUpdateTask, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getMLUpdateTaskClass().class), objc.Sel("updateTaskForModelAtURL:trainingData:progressHandlers:error:"), modelURL, trainingData, progressHandlers, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLUpdateTaskFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return MLUpdateTaskFromID(rv), nil
}

// Resumes a model update with updated parameter values.
//
// updateParameters: Model training parameter values to replace those currently set in the
// update task.
//
// # Discussion
// 
// Use this method to resume the model update task with newer parameter
// values. You use this method within the closures you provide in an
// [MLUpdateProgressHandlers] instance to resume the [MLUpdateTask].
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateTask/resume(withParameters:)
func (u MLUpdateTask) ResumeWithParameters(updateParameters foundation.INSDictionary) {
	objc.Send[objc.ID](u.ID, objc.Sel("resumeWithParameters:"), updateParameters)
}

// Creates a task that updates the model at the URL with the training data,
// and calls the completion handler when the update completes.
//
// modelURL: The location in the file system of a model file
// (`XCUIElementTypeMlmodelc`).
//
// trainingData: The update data for the model, contained in a batch provider.
//
// completionHandler: The closure the task calls when it finishes.
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:completionHandler:)
func (_MLUpdateTaskClass MLUpdateTaskClass) UpdateTaskForModelAtURLTrainingDataCompletionHandlerError(modelURL foundation.INSURL, trainingData MLBatchProvider, completionHandler func(*MLUpdateContext)) (MLUpdateTask, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLUpdateTaskClass.class), objc.Sel("updateTaskForModelAtURL:trainingData:completionHandler:error:"), modelURL, trainingData, completionHandler, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLUpdateTask{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLUpdateTaskFromID(rv), nil

}
// Creates a task that updates the model at the URL with the training data and
// configuration, and calls the completion handler when the update completes.
//
// modelURL: The location in the file system of a model file
// (`XCUIElementTypeMlmodelc`).
//
// trainingData: The update data for the model, contained in a batch provider.
//
// configuration: The model settings for an updated model object.
//
// completionHandler: The closure the task calls when it finishes.
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:configuration:completionHandler:)
func (_MLUpdateTaskClass MLUpdateTaskClass) UpdateTaskForModelAtURLTrainingDataConfigurationCompletionHandlerError(modelURL foundation.INSURL, trainingData MLBatchProvider, configuration IMLModelConfiguration, completionHandler func(*MLUpdateContext)) (MLUpdateTask, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLUpdateTaskClass.class), objc.Sel("updateTaskForModelAtURL:trainingData:configuration:completionHandler:error:"), modelURL, trainingData, configuration, completionHandler, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLUpdateTask{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLUpdateTaskFromID(rv), nil

}


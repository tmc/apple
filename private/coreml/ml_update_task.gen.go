// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
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

// Class returns the underlying Objective-C class pointer.
func (mc MLUpdateTaskClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLUpdateTaskClass) Alloc() MLUpdateTask {
	rv := objc.SendIfResponds[MLUpdateTask](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLUpdateTask._completionHandlerBlock]
//   - [MLUpdateTask._invokeProgressHandlerForContext]
//   - [MLUpdateTask._progressHandlerBlock]
//   - [MLUpdateTask.OnCancellation]
//   - [MLUpdateTask.OnCompletionWithTaskContext]
//   - [MLUpdateTask.OnFailureWithTaskContext]
//   - [MLUpdateTask.OnResumptionWithTaskContext]
//   - [MLUpdateTask.OnSuspensionWithTaskContext]
//   - [MLUpdateTask.ProgressHandlers]
//   - [MLUpdateTask.TrainingData]
//   - [MLUpdateTask.UpdatableModel]
//   - [MLUpdateTask.UpdatableModelURL]
//   - [MLUpdateTask.UpdateHasStarted]
//   - [MLUpdateTask.SetUpdateHasStarted]
//   - [MLUpdateTask.UpdateQueue]
//   - [MLUpdateTask.InitWithModelAtURLTrainingDataConfigurationProgressHandlersError]
//   - [MLUpdateTask.DebugDescription]
//   - [MLUpdateTask.Description]
//   - [MLUpdateTask.Hash]
//   - [MLUpdateTask.Superclass]
type MLUpdateTask struct {
	MLTask
}

// MLUpdateTaskFromID constructs a [MLUpdateTask] from an objc.ID.
func MLUpdateTaskFromID(id objc.ID) MLUpdateTask {
	return MLUpdateTask{MLTask: MLTaskFromID(id)}
}

// Ensure MLUpdateTask implements IMLUpdateTask.
var _ IMLUpdateTask = MLUpdateTask{}

// An interface definition for the [MLUpdateTask] class.
//
// # Methods
//
//   - [IMLUpdateTask._completionHandlerBlock]
//   - [IMLUpdateTask._invokeProgressHandlerForContext]
//   - [IMLUpdateTask._progressHandlerBlock]
//   - [IMLUpdateTask.OnCancellation]
//   - [IMLUpdateTask.OnCompletionWithTaskContext]
//   - [IMLUpdateTask.OnFailureWithTaskContext]
//   - [IMLUpdateTask.OnResumptionWithTaskContext]
//   - [IMLUpdateTask.OnSuspensionWithTaskContext]
//   - [IMLUpdateTask.ProgressHandlers]
//   - [IMLUpdateTask.TrainingData]
//   - [IMLUpdateTask.UpdatableModel]
//   - [IMLUpdateTask.UpdatableModelURL]
//   - [IMLUpdateTask.UpdateHasStarted]
//   - [IMLUpdateTask.SetUpdateHasStarted]
//   - [IMLUpdateTask.UpdateQueue]
//   - [IMLUpdateTask.InitWithModelAtURLTrainingDataConfigurationProgressHandlersError]
//   - [IMLUpdateTask.DebugDescription]
//   - [IMLUpdateTask.Description]
//   - [IMLUpdateTask.Hash]
//   - [IMLUpdateTask.Superclass]
type IMLUpdateTask interface {
	IMLTask

	// Topic: Methods

	_completionHandlerBlock()
	_invokeProgressHandlerForContext(context objectivec.IObject)
	_progressHandlerBlock()
	OnCancellation()
	OnCompletionWithTaskContext(context objectivec.IObject)
	OnFailureWithTaskContext(context objectivec.IObject)
	OnResumptionWithTaskContext(context objectivec.IObject)
	OnSuspensionWithTaskContext(context objectivec.IObject)
	ProgressHandlers() IMLUpdateProgressHandlers
	TrainingData() unsafe.Pointer
	UpdatableModel() IMLModel
	UpdatableModelURL() foundation.NSURL
	UpdateHasStarted() bool
	SetUpdateHasStarted(value bool)
	UpdateQueue() objectivec.Object
	InitWithModelAtURLTrainingDataConfigurationProgressHandlersError(url foundation.NSURL, data objectivec.IObject, configuration objectivec.IObject, handlers objectivec.IObject) (MLUpdateTask, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLUpdateTask) Init() MLUpdateTask {
	rv := objc.SendIfResponds[MLUpdateTask](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLUpdateTask) Autorelease() MLUpdateTask {
	rv := objc.SendIfResponds[MLUpdateTask](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLUpdateTask creates a new MLUpdateTask instance.
func NewMLUpdateTask() MLUpdateTask {
	class := getMLUpdateTaskClass()
	rv := objc.SendIfResponds[MLUpdateTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewUpdateTaskWithModelAtURLTrainingDataConfigurationProgressHandlersError(url foundation.NSURL, data objectivec.IObject, configuration objectivec.IObject, handlers objectivec.IObject) (MLUpdateTask, error) {
	var errorPtr objc.ID
	instance := getMLUpdateTaskClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithModelAtURL:trainingData:configuration:progressHandlers:error:"), url, data, configuration, handlers, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLUpdateTask{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLUpdateTask{}, objc.ErrInitFailed
	}
	return MLUpdateTaskFromID(rv), nil
}

func NewUpdateTaskWithState(state int64) MLUpdateTask {
	instance := getMLUpdateTaskClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithState:"), state)
	return MLUpdateTaskFromID(rv)
}

func (m MLUpdateTask) _completionHandlerBlock() {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("_completionHandlerBlock"))
}

// CompletionHandlerBlock is an exported wrapper for the private method _completionHandlerBlock.
func (m MLUpdateTask) CompletionHandlerBlock() error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_completionHandlerBlock")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_completionHandlerBlock"}
		return err
	}
	m._completionHandlerBlock()
	return nil
}

// CanCompletionHandlerBlock reports whether the receiver responds to the private selector _completionHandlerBlock.
func (m MLUpdateTask) CanCompletionHandlerBlock() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_completionHandlerBlock"))
}
func (m MLUpdateTask) _invokeProgressHandlerForContext(context objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("_invokeProgressHandlerForContext:"), context)
}

// InvokeProgressHandlerForContext is an exported wrapper for the private method _invokeProgressHandlerForContext.
func (m MLUpdateTask) InvokeProgressHandlerForContext(context objectivec.IObject) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_invokeProgressHandlerForContext:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_invokeProgressHandlerForContext:"}
		return err
	}
	m._invokeProgressHandlerForContext(context)
	return nil
}

// CanInvokeProgressHandlerForContext reports whether the receiver responds to the private selector _invokeProgressHandlerForContext:.
func (m MLUpdateTask) CanInvokeProgressHandlerForContext() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_invokeProgressHandlerForContext:"))
}
func (m MLUpdateTask) _progressHandlerBlock() {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("_progressHandlerBlock"))
}

// ProgressHandlerBlock is an exported wrapper for the private method _progressHandlerBlock.
func (m MLUpdateTask) ProgressHandlerBlock() error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_progressHandlerBlock")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_progressHandlerBlock"}
		return err
	}
	m._progressHandlerBlock()
	return nil
}

// CanProgressHandlerBlock reports whether the receiver responds to the private selector _progressHandlerBlock.
func (m MLUpdateTask) CanProgressHandlerBlock() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_progressHandlerBlock"))
}
func (m MLUpdateTask) OnCancellation() {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("onCancellation"))
}
func (m MLUpdateTask) OnCompletionWithTaskContext(context objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("onCompletionWithTaskContext:"), context)
}
func (m MLUpdateTask) OnFailureWithTaskContext(context objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("onFailureWithTaskContext:"), context)
}
func (m MLUpdateTask) OnResumptionWithTaskContext(context objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("onResumptionWithTaskContext:"), context)
}
func (m MLUpdateTask) OnSuspensionWithTaskContext(context objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("onSuspensionWithTaskContext:"), context)
}
func (m MLUpdateTask) InitWithModelAtURLTrainingDataConfigurationProgressHandlersError(url foundation.NSURL, data objectivec.IObject, configuration objectivec.IObject, handlers objectivec.IObject) (MLUpdateTask, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithModelAtURL:trainingData:configuration:progressHandlers:error:"), url, data, configuration, handlers, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLUpdateTask{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLUpdateTaskFromID(rv), nil

}

func (_MLUpdateTaskClass MLUpdateTaskClass) UpdateModelAtURLTrainingDataConfigurationWriteToURLError(url foundation.NSURL, data objectivec.IObject, configuration objectivec.IObject, url2 foundation.NSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLUpdateTaskClass.class), objc.Sel("updateModelAtURL:trainingData:configuration:writeToURL:error:"), url, data, configuration, url2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateModelAtURL:trainingData:configuration:writeToURL:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (_MLUpdateTaskClass MLUpdateTaskClass) UpdateTaskForModelAtURLTrainingDataConfigurationProgressHandlersError(url foundation.NSURL, data objectivec.IObject, configuration objectivec.IObject, handlers objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLUpdateTaskClass.class), objc.Sel("updateTaskForModelAtURL:trainingData:configuration:progressHandlers:error:"), url, data, configuration, handlers, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLUpdateTaskClass MLUpdateTaskClass) UpdateTaskForModelAtURLTrainingDataProgressHandlersError(url foundation.NSURL, data objectivec.IObject, handlers objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLUpdateTaskClass.class), objc.Sel("updateTaskForModelAtURL:trainingData:progressHandlers:error:"), url, data, handlers, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

func (m MLUpdateTask) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLUpdateTask) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLUpdateTask) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLUpdateTask) ProgressHandlers() IMLUpdateProgressHandlers {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("progressHandlers"))
	return MLUpdateProgressHandlersFromID(objc.ID(rv))
}
func (m MLUpdateTask) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (m MLUpdateTask) TrainingData() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("trainingData"))
	return rv
}
func (m MLUpdateTask) UpdatableModel() IMLModel {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("updatableModel"))
	return MLModelFromID(objc.ID(rv))
}
func (m MLUpdateTask) UpdatableModelURL() foundation.NSURL {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("updatableModelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (m MLUpdateTask) UpdateHasStarted() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("updateHasStarted"))
	return rv
}
func (m MLUpdateTask) SetUpdateHasStarted(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setUpdateHasStarted:"), value)
}
func (m MLUpdateTask) UpdateQueue() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("updateQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

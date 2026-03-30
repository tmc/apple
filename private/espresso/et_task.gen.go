// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETTask] class.
var (
	_ETTaskClass     ETTaskClass
	_ETTaskClassOnce sync.Once
)

func getETTaskClass() ETTaskClass {
	_ETTaskClassOnce.Do(func() {
		_ETTaskClass = ETTaskClass{class: objc.GetClass("ETTask")}
	})
	return _ETTaskClass
}

// GetETTaskClass returns the class object for ETTask.
func GetETTaskClass() ETTaskClass {
	return getETTaskClass()
}

type ETTaskClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETTaskClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETTaskClass) Alloc() ETTask {
	rv := objc.Send[ETTask](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ETTask.DumpData]
//   - [ETTask.SetDumpData]
//   - [ETTask.Evaluate]
//   - [ETTask.Extractor]
//   - [ETTask.SetExtractor]
//   - [ETTask.FitNumberOfBatchesOutputNamesBatchCallback]
//   - [ETTask.FitNumberOfBatchesWithProgress]
//   - [ETTask.FitNumberOfEpochsOutputNamesBatchCallback]
//   - [ETTask.FitNumberOfEpochsWithProgress]
//   - [ETTask.Model]
//   - [ETTask.SetModel]
//   - [ETTask.MoveToGPUError]
//   - [ETTask.Optimizer]
//   - [ETTask.SetOptimizer]
//   - [ETTask.ReinitializeVariables]
//   - [ETTask.RunBatchesNumberOfBatchesOutputNamesBatchCallback]
//   - [ETTask.RunInferenceOutputNamesBatchCallback]
//   - [ETTask.SaveNetwork]
//   - [ETTask.SaveNetworkRevertToInferenceMode]
//   - [ETTask.InitWithModelDefOptimizerDefExtractor]
//   - [ETTask.InitWithModelDefOptimizerDefExtractorNeedWeightsInitialization]
//   - [ETTask.InitWithModelDefOptimizerDefLossConfig]
//   - [ETTask.InitWithModelDefOptimizerDefLossConfigExtractor]
//
// See: https://developer.apple.com/documentation/Espresso/ETTask
type ETTask struct {
	objectivec.Object
}

// ETTaskFromID constructs a [ETTask] from an objc.ID.
func ETTaskFromID(id objc.ID) ETTask {
	return ETTask{objectivec.Object{ID: id}}
}

// Ensure ETTask implements IETTask.
var _ IETTask = ETTask{}

// An interface definition for the [ETTask] class.
//
// # Methods
//
//   - [IETTask.DumpData]
//   - [IETTask.SetDumpData]
//   - [IETTask.Evaluate]
//   - [IETTask.Extractor]
//   - [IETTask.SetExtractor]
//   - [IETTask.FitNumberOfBatchesOutputNamesBatchCallback]
//   - [IETTask.FitNumberOfBatchesWithProgress]
//   - [IETTask.FitNumberOfEpochsOutputNamesBatchCallback]
//   - [IETTask.FitNumberOfEpochsWithProgress]
//   - [IETTask.Model]
//   - [IETTask.SetModel]
//   - [IETTask.MoveToGPUError]
//   - [IETTask.Optimizer]
//   - [IETTask.SetOptimizer]
//   - [IETTask.ReinitializeVariables]
//   - [IETTask.RunBatchesNumberOfBatchesOutputNamesBatchCallback]
//   - [IETTask.RunInferenceOutputNamesBatchCallback]
//   - [IETTask.SaveNetwork]
//   - [IETTask.SaveNetworkRevertToInferenceMode]
//   - [IETTask.InitWithModelDefOptimizerDefExtractor]
//   - [IETTask.InitWithModelDefOptimizerDefExtractorNeedWeightsInitialization]
//   - [IETTask.InitWithModelDefOptimizerDefLossConfig]
//   - [IETTask.InitWithModelDefOptimizerDefLossConfigExtractor]
//
// See: https://developer.apple.com/documentation/Espresso/ETTask
type IETTask interface {
	objectivec.IObject

	// Topic: Methods

	DumpData() bool
	SetDumpData(value bool)
	Evaluate(evaluate objectivec.IObject) objectivec.IObject
	Extractor() IETImageDescriptorExtractor
	SetExtractor(value IETImageDescriptorExtractor)
	FitNumberOfBatchesOutputNamesBatchCallback(fit objectivec.IObject, batches uint32, names objectivec.IObject, callback VoidHandler) bool
	FitNumberOfBatchesWithProgress(fit objectivec.IObject, batches uint32, progress VoidHandler) float32
	FitNumberOfEpochsOutputNamesBatchCallback(fit objectivec.IObject, epochs int, names objectivec.IObject, callback VoidHandler) bool
	FitNumberOfEpochsWithProgress(fit objectivec.IObject, epochs int, progress VoidHandler) float32
	Model() IETModelDef
	SetModel(value IETModelDef)
	MoveToGPUError(gpu int) (bool, error)
	Optimizer() IETOptimizerDef
	SetOptimizer(value IETOptimizerDef)
	ReinitializeVariables() objectivec.IObject
	RunBatchesNumberOfBatchesOutputNamesBatchCallback(batches objectivec.IObject, batches2 uint32, names objectivec.IObject, callback VoidHandler) bool
	RunInferenceOutputNamesBatchCallback(inference objectivec.IObject, names objectivec.IObject, callback VoidHandler) bool
	SaveNetwork(network objectivec.IObject)
	SaveNetworkRevertToInferenceMode(network objectivec.IObject, mode bool)
	InitWithModelDefOptimizerDefExtractor(def objectivec.IObject, def2 objectivec.IObject, extractor objectivec.IObject) ETTask
	InitWithModelDefOptimizerDefExtractorNeedWeightsInitialization(def objectivec.IObject, def2 objectivec.IObject, extractor objectivec.IObject, initialization bool) ETTask
	InitWithModelDefOptimizerDefLossConfig(def objectivec.IObject, def2 objectivec.IObject, config objectivec.IObject) ETTask
	InitWithModelDefOptimizerDefLossConfigExtractor(def objectivec.IObject, def2 objectivec.IObject, config objectivec.IObject, extractor objectivec.IObject) ETTask
}

// Init initializes the instance.
func (e ETTask) Init() ETTask {
	rv := objc.Send[ETTask](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETTask) Autorelease() ETTask {
	rv := objc.Send[ETTask](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETTask creates a new ETTask instance.
func NewETTask() ETTask {
	class := getETTaskClass()
	rv := objc.Send[ETTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/initWithModelDef:optimizerDef:extractor:
func NewETTaskWithModelDefOptimizerDefExtractor(def objectivec.IObject, def2 objectivec.IObject, extractor objectivec.IObject) ETTask {
	instance := getETTaskClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelDef:optimizerDef:extractor:"), def, def2, extractor)
	return ETTaskFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/initWithModelDef:optimizerDef:extractor:needWeightsInitialization:
func NewETTaskWithModelDefOptimizerDefExtractorNeedWeightsInitialization(def objectivec.IObject, def2 objectivec.IObject, extractor objectivec.IObject, initialization bool) ETTask {
	instance := getETTaskClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelDef:optimizerDef:extractor:needWeightsInitialization:"), def, def2, extractor, initialization)
	return ETTaskFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/initWithModelDef:optimizerDef:lossConfig:
func NewETTaskWithModelDefOptimizerDefLossConfig(def objectivec.IObject, def2 objectivec.IObject, config objectivec.IObject) ETTask {
	instance := getETTaskClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelDef:optimizerDef:lossConfig:"), def, def2, config)
	return ETTaskFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/initWithModelDef:optimizerDef:lossConfig:extractor:
func NewETTaskWithModelDefOptimizerDefLossConfigExtractor(def objectivec.IObject, def2 objectivec.IObject, config objectivec.IObject, extractor objectivec.IObject) ETTask {
	instance := getETTaskClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelDef:optimizerDef:lossConfig:extractor:"), def, def2, config, extractor)
	return ETTaskFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/evaluate:
func (e ETTask) Evaluate(evaluate objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("evaluate:"), evaluate)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/fit:numberOfBatches:outputNames:batchCallback:
func (e ETTask) FitNumberOfBatchesOutputNamesBatchCallback(fit objectivec.IObject, batches uint32, names objectivec.IObject, callback VoidHandler) bool {
	_block3, _ := NewVoidBlock(callback)
	rv := objc.Send[bool](e.ID, objc.Sel("fit:numberOfBatches:outputNames:batchCallback:"), fit, batches, names, _block3)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/fit:numberOfBatches:withProgress:
func (e ETTask) FitNumberOfBatchesWithProgress(fit objectivec.IObject, batches uint32, progress VoidHandler) float32 {
	_block2, _ := NewVoidBlock(progress)
	rv := objc.Send[float32](e.ID, objc.Sel("fit:numberOfBatches:withProgress:"), fit, batches, _block2)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/fit:numberOfEpochs:outputNames:batchCallback:
func (e ETTask) FitNumberOfEpochsOutputNamesBatchCallback(fit objectivec.IObject, epochs int, names objectivec.IObject, callback VoidHandler) bool {
	_block3, _ := NewVoidBlock(callback)
	rv := objc.Send[bool](e.ID, objc.Sel("fit:numberOfEpochs:outputNames:batchCallback:"), fit, epochs, names, _block3)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/fit:numberOfEpochs:withProgress:
func (e ETTask) FitNumberOfEpochsWithProgress(fit objectivec.IObject, epochs int, progress VoidHandler) float32 {
	_block2, _ := NewVoidBlock(progress)
	rv := objc.Send[float32](e.ID, objc.Sel("fit:numberOfEpochs:withProgress:"), fit, epochs, _block2)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/moveToGPU:error:
func (e ETTask) MoveToGPUError(gpu int) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("moveToGPU:error:"), gpu, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("moveToGPU:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/Espresso/ETTask/reinitializeVariables
func (e ETTask) ReinitializeVariables() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("reinitializeVariables"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/runBatches:numberOfBatches:outputNames:batchCallback:
func (e ETTask) RunBatchesNumberOfBatchesOutputNamesBatchCallback(batches objectivec.IObject, batches2 uint32, names objectivec.IObject, callback VoidHandler) bool {
	_block3, _ := NewVoidBlock(callback)
	rv := objc.Send[bool](e.ID, objc.Sel("runBatches:numberOfBatches:outputNames:batchCallback:"), batches, batches2, names, _block3)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/runInference:outputNames:batchCallback:
func (e ETTask) RunInferenceOutputNamesBatchCallback(inference objectivec.IObject, names objectivec.IObject, callback VoidHandler) bool {
	_block2, _ := NewVoidBlock(callback)
	rv := objc.Send[bool](e.ID, objc.Sel("runInference:outputNames:batchCallback:"), inference, names, _block2)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/saveNetwork:
func (e ETTask) SaveNetwork(network objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("saveNetwork:"), network)
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/saveNetwork:revertToInferenceMode:
func (e ETTask) SaveNetworkRevertToInferenceMode(network objectivec.IObject, mode bool) {
	objc.Send[objc.ID](e.ID, objc.Sel("saveNetwork:revertToInferenceMode:"), network, mode)
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/initWithModelDef:optimizerDef:extractor:
func (e ETTask) InitWithModelDefOptimizerDefExtractor(def objectivec.IObject, def2 objectivec.IObject, extractor objectivec.IObject) ETTask {
	rv := objc.Send[ETTask](e.ID, objc.Sel("initWithModelDef:optimizerDef:extractor:"), def, def2, extractor)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/initWithModelDef:optimizerDef:extractor:needWeightsInitialization:
func (e ETTask) InitWithModelDefOptimizerDefExtractorNeedWeightsInitialization(def objectivec.IObject, def2 objectivec.IObject, extractor objectivec.IObject, initialization bool) ETTask {
	rv := objc.Send[ETTask](e.ID, objc.Sel("initWithModelDef:optimizerDef:extractor:needWeightsInitialization:"), def, def2, extractor, initialization)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/initWithModelDef:optimizerDef:lossConfig:
func (e ETTask) InitWithModelDefOptimizerDefLossConfig(def objectivec.IObject, def2 objectivec.IObject, config objectivec.IObject) ETTask {
	rv := objc.Send[ETTask](e.ID, objc.Sel("initWithModelDef:optimizerDef:lossConfig:"), def, def2, config)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/initWithModelDef:optimizerDef:lossConfig:extractor:
func (e ETTask) InitWithModelDefOptimizerDefLossConfigExtractor(def objectivec.IObject, def2 objectivec.IObject, config objectivec.IObject, extractor objectivec.IObject) ETTask {
	rv := objc.Send[ETTask](e.ID, objc.Sel("initWithModelDef:optimizerDef:lossConfig:extractor:"), def, def2, config, extractor)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/dumpData
func (e ETTask) DumpData() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("dumpData"))
	return rv
}
func (e ETTask) SetDumpData(value bool) {
	objc.Send[struct{}](e.ID, objc.Sel("setDumpData:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/extractor
func (e ETTask) Extractor() IETImageDescriptorExtractor {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("extractor"))
	return ETImageDescriptorExtractorFromID(objc.ID(rv))
}
func (e ETTask) SetExtractor(value IETImageDescriptorExtractor) {
	objc.Send[struct{}](e.ID, objc.Sel("setExtractor:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/model
func (e ETTask) Model() IETModelDef {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("model"))
	return ETModelDefFromID(objc.ID(rv))
}
func (e ETTask) SetModel(value IETModelDef) {
	objc.Send[struct{}](e.ID, objc.Sel("setModel:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/optimizer
func (e ETTask) Optimizer() IETOptimizerDef {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("optimizer"))
	return ETOptimizerDefFromID(objc.ID(rv))
}
func (e ETTask) SetOptimizer(value IETOptimizerDef) {
	objc.Send[struct{}](e.ID, objc.Sel("setOptimizer:"), value)
}

// FitNumberOfBatchesOutputNamesBatchCallbackSync is a synchronous wrapper around [ETTask.FitNumberOfBatchesOutputNamesBatchCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (e ETTask) FitNumberOfBatchesOutputNamesBatchCallbackSync(ctx context.Context, fit objectivec.IObject, batches uint32, names objectivec.IObject) error {
	done := make(chan struct{}, 1)
	e.FitNumberOfBatchesOutputNamesBatchCallback(fit, batches, names, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// FitNumberOfBatchesWithProgressSync is a synchronous wrapper around [ETTask.FitNumberOfBatchesWithProgress].
// It blocks until the completion handler fires or the context is cancelled.
func (e ETTask) FitNumberOfBatchesWithProgressSync(ctx context.Context, fit objectivec.IObject, batches uint32) error {
	done := make(chan struct{}, 1)
	e.FitNumberOfBatchesWithProgress(fit, batches, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// FitNumberOfEpochsOutputNamesBatchCallbackSync is a synchronous wrapper around [ETTask.FitNumberOfEpochsOutputNamesBatchCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (e ETTask) FitNumberOfEpochsOutputNamesBatchCallbackSync(ctx context.Context, fit objectivec.IObject, epochs int, names objectivec.IObject) error {
	done := make(chan struct{}, 1)
	e.FitNumberOfEpochsOutputNamesBatchCallback(fit, epochs, names, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// FitNumberOfEpochsWithProgressSync is a synchronous wrapper around [ETTask.FitNumberOfEpochsWithProgress].
// It blocks until the completion handler fires or the context is cancelled.
func (e ETTask) FitNumberOfEpochsWithProgressSync(ctx context.Context, fit objectivec.IObject, epochs int) error {
	done := make(chan struct{}, 1)
	e.FitNumberOfEpochsWithProgress(fit, epochs, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RunBatchesNumberOfBatchesOutputNamesBatchCallbackSync is a synchronous wrapper around [ETTask.RunBatchesNumberOfBatchesOutputNamesBatchCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (e ETTask) RunBatchesNumberOfBatchesOutputNamesBatchCallbackSync(ctx context.Context, batches objectivec.IObject, batches2 uint32, names objectivec.IObject) error {
	done := make(chan struct{}, 1)
	e.RunBatchesNumberOfBatchesOutputNamesBatchCallback(batches, batches2, names, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RunInferenceOutputNamesBatchCallbackSync is a synchronous wrapper around [ETTask.RunInferenceOutputNamesBatchCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (e ETTask) RunInferenceOutputNamesBatchCallbackSync(ctx context.Context, inference objectivec.IObject, names objectivec.IObject) error {
	done := make(chan struct{}, 1)
	e.RunInferenceOutputNamesBatchCallback(inference, names, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

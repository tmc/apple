// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLPipelineClassifier] class.
var (
	_MLPipelineClassifierClass     MLPipelineClassifierClass
	_MLPipelineClassifierClassOnce sync.Once
)

func getMLPipelineClassifierClass() MLPipelineClassifierClass {
	_MLPipelineClassifierClassOnce.Do(func() {
		_MLPipelineClassifierClass = MLPipelineClassifierClass{class: objc.GetClass("MLPipelineClassifier")}
	})
	return _MLPipelineClassifierClass
}

// GetMLPipelineClassifierClass returns the class object for MLPipelineClassifier.
func GetMLPipelineClassifierClass() MLPipelineClassifierClass {
	return getMLPipelineClassifierClass()
}

type MLPipelineClassifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLPipelineClassifierClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLPipelineClassifierClass) Alloc() MLPipelineClassifier {
	rv := objc.Send[MLPipelineClassifier](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLPipelineClassifier.ClassifyOptionsCompletionHandler]
//   - [MLPipelineClassifier.ClassifyOptionsError]
//   - [MLPipelineClassifier.EnableInstrumentsTracing]
//   - [MLPipelineClassifier.Engine]
//   - [MLPipelineClassifier.SetEngine]
//   - [MLPipelineClassifier.ExecutionSchedule]
//   - [MLPipelineClassifier.Pipeline]
//   - [MLPipelineClassifier.SignpostID]
//   - [MLPipelineClassifier.InitWithEngineDescriptionConfigurationError]
//   - [MLPipelineClassifier.Configuration]
//   - [MLPipelineClassifier.DebugDescription]
//   - [MLPipelineClassifier.Description]
//   - [MLPipelineClassifier.Hash]
//   - [MLPipelineClassifier.Metadata]
//   - [MLPipelineClassifier.ModelDescription]
//   - [MLPipelineClassifier.PredictionTypeForKTrace]
//   - [MLPipelineClassifier.RecordsPredictionEvent]
//   - [MLPipelineClassifier.Superclass]
//   - [MLPipelineClassifier.SupportsConcurrentSubmissions]
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier
type MLPipelineClassifier struct {
	objectivec.Object
}

// MLPipelineClassifierFromID constructs a [MLPipelineClassifier] from an objc.ID.
func MLPipelineClassifierFromID(id objc.ID) MLPipelineClassifier {
	return MLPipelineClassifier{objectivec.Object{ID: id}}
}

// NOTE: MLPipelineClassifier struct embeds objectivec.Object (parent type unavailable) but
// IMLPipelineClassifier embeds the parent interface; skip compile-time assertion.

// An interface definition for the [MLPipelineClassifier] class.
//
// # Methods
//
//   - [IMLPipelineClassifier.ClassifyOptionsCompletionHandler]
//   - [IMLPipelineClassifier.ClassifyOptionsError]
//   - [IMLPipelineClassifier.EnableInstrumentsTracing]
//   - [IMLPipelineClassifier.Engine]
//   - [IMLPipelineClassifier.SetEngine]
//   - [IMLPipelineClassifier.ExecutionSchedule]
//   - [IMLPipelineClassifier.Pipeline]
//   - [IMLPipelineClassifier.SignpostID]
//   - [IMLPipelineClassifier.InitWithEngineDescriptionConfigurationError]
//   - [IMLPipelineClassifier.Configuration]
//   - [IMLPipelineClassifier.DebugDescription]
//   - [IMLPipelineClassifier.Description]
//   - [IMLPipelineClassifier.Hash]
//   - [IMLPipelineClassifier.Metadata]
//   - [IMLPipelineClassifier.ModelDescription]
//   - [IMLPipelineClassifier.PredictionTypeForKTrace]
//   - [IMLPipelineClassifier.RecordsPredictionEvent]
//   - [IMLPipelineClassifier.Superclass]
//   - [IMLPipelineClassifier.SupportsConcurrentSubmissions]
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier
type IMLPipelineClassifier interface {
	IMLClassifier

	// Topic: Methods

	ClassifyOptionsCompletionHandler(classify objectivec.IObject, options objectivec.IObject, handler ErrorHandler)
	ClassifyOptionsError(classify objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	EnableInstrumentsTracing()
	Engine() unsafe.Pointer
	SetEngine(value *MLPipeline)
	ExecutionSchedule() objectivec.IObject
	Pipeline() unsafe.Pointer
	SignpostID() uint64
	InitWithEngineDescriptionConfigurationError(engine objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLPipelineClassifier, error)
	Configuration() IMLModelConfiguration
	DebugDescription() string
	Description() string
	Hash() uint64
	Metadata() IMLModelMetadata
	ModelDescription() IMLModelDescription
	PredictionTypeForKTrace() uint64
	RecordsPredictionEvent() bool
	Superclass() objc.Class
	SupportsConcurrentSubmissions() bool
}

// Init initializes the instance.
func (m MLPipelineClassifier) Init() MLPipelineClassifier {
	rv := objc.Send[MLPipelineClassifier](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLPipelineClassifier) Autorelease() MLPipelineClassifier {
	rv := objc.Send[MLPipelineClassifier](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPipelineClassifier creates a new MLPipelineClassifier instance.
func NewMLPipelineClassifier() MLPipelineClassifier {
	class := getMLPipelineClassifierClass()
	rv := objc.Send[MLPipelineClassifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier/initWithEngine:description:configuration:error:
func NewPipelineClassifierWithEngineDescriptionConfigurationError(engine objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLPipelineClassifier, error) {
	var errorPtr objc.ID
	instance := getMLPipelineClassifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEngine:description:configuration:error:"), engine, description, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLPipelineClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLPipelineClassifierFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier/classify:options:completionHandler:
func (m MLPipelineClassifier) ClassifyOptionsCompletionHandler(classify objectivec.IObject, options objectivec.IObject, handler ErrorHandler) {
	_block2, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("classify:options:completionHandler:"), classify, options, _block2)
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier/classify:options:error:
func (m MLPipelineClassifier) ClassifyOptionsError(classify objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classify:options:error:"), classify, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier/enableInstrumentsTracing
func (m MLPipelineClassifier) EnableInstrumentsTracing() {
	objc.Send[objc.ID](m.ID, objc.Sel("enableInstrumentsTracing"))
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier/executionSchedule
func (m MLPipelineClassifier) ExecutionSchedule() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("executionSchedule"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier/initWithEngine:description:configuration:error:
func (m MLPipelineClassifier) InitWithEngineDescriptionConfigurationError(engine objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLPipelineClassifier, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithEngine:description:configuration:error:"), engine, description, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLPipelineClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLPipelineClassifierFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier/configuration
func (m MLPipelineClassifier) Configuration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier/debugDescription
func (m MLPipelineClassifier) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier/description
func (m MLPipelineClassifier) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier/engine
func (m MLPipelineClassifier) Engine() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("engine"))
	return rv
}
func (m MLPipelineClassifier) SetEngine(value *MLPipeline) {
	objc.Send[struct{}](m.ID, objc.Sel("setEngine:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier/hash
func (m MLPipelineClassifier) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier/metadata
func (m MLPipelineClassifier) Metadata() IMLModelMetadata {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("metadata"))
	return MLModelMetadataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier/modelDescription
func (m MLPipelineClassifier) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier/pipeline
func (m MLPipelineClassifier) Pipeline() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("pipeline"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier/predictionTypeForKTrace
func (m MLPipelineClassifier) PredictionTypeForKTrace() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("predictionTypeForKTrace"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier/recordsPredictionEvent
func (m MLPipelineClassifier) RecordsPredictionEvent() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("recordsPredictionEvent"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier/signpostID
func (m MLPipelineClassifier) SignpostID() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("signpostID"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier/superclass
func (m MLPipelineClassifier) Superclass() objc.Class {
	rv := objc.Send[objc.Class](m.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineClassifier/supportsConcurrentSubmissions
func (m MLPipelineClassifier) SupportsConcurrentSubmissions() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportsConcurrentSubmissions"))
	return rv
}

// ClassifyOptions is a synchronous wrapper around [MLPipelineClassifier.ClassifyOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLPipelineClassifier) ClassifyOptions(ctx context.Context, classify objectivec.IObject, options objectivec.IObject) error {
	done := make(chan error, 1)
	m.ClassifyOptionsCompletionHandler(classify, options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

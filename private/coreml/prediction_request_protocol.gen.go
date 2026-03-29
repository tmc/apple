// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLPredictionRequest protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLPredictionRequest
type MLPredictionRequest interface {
	objectivec.IObject

	// Cancel protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLPredictionRequest/cancel
	Cancel()

	// IsCancelled protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLPredictionRequest/isCancelled
	IsCancelled() bool

	// SubmitWithCompletionHandler protocol.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLPredictionRequest/submitWithCompletionHandler:
	SubmitWithCompletionHandler(handler ErrorHandler)
}

// MLPredictionRequestObject wraps an existing Objective-C object that conforms to the MLPredictionRequest protocol.
type MLPredictionRequestObject struct {
	objectivec.Object
}
func (o MLPredictionRequestObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLPredictionRequestObjectFromID constructs a [MLPredictionRequestObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLPredictionRequestObjectFromID(id objc.ID) MLPredictionRequestObject {
	return MLPredictionRequestObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionRequest/cancel
func (o MLPredictionRequestObject) Cancel() {
	objc.Send[struct{}](o.ID, objc.Sel("cancel"))
	}
// See: https://developer.apple.com/documentation/CoreML/MLPredictionRequest/inputFeatures
func (o MLPredictionRequestObject) InputFeatures() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputFeatures"))
	return objectivec.Object{ID: rv}
	}
// See: https://developer.apple.com/documentation/CoreML/MLPredictionRequest/isCancelled
func (o MLPredictionRequestObject) IsCancelled() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isCancelled"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreML/MLPredictionRequest/predictionOptions
func (o MLPredictionRequestObject) PredictionOptions() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("predictionOptions"))
	return objectivec.Object{ID: rv}
	}
//
// See: https://developer.apple.com/documentation/CoreML/MLPredictionRequest/submitWithCompletionHandler:
func (o MLPredictionRequestObject) SubmitWithCompletionHandler(handler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("submitWithCompletionHandler:"), handler)
	}


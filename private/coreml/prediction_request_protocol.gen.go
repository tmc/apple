// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLPredictionRequest protocol.
type MLPredictionRequest interface {
	objectivec.IObject

	// Cancel protocol.
	Cancel()

	// InputFeatures protocol.
	InputFeatures() objectivec.IObject

	// IsCancelled protocol.
	IsCancelled() bool

	// PredictionOptions protocol.
	PredictionOptions() objectivec.IObject
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

func (o MLPredictionRequestObject) Cancel() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("cancel"))
}
func (o MLPredictionRequestObject) InputFeatures() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("inputFeatures"))
	return objectivec.Object{ID: rv}
}
func (o MLPredictionRequestObject) IsCancelled() bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("isCancelled"))
	return rv
}
func (o MLPredictionRequestObject) PredictionOptions() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("predictionOptions"))
	return objectivec.Object{ID: rv}
}

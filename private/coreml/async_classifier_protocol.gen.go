// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLAsyncClassifier protocol.
type MLAsyncClassifier interface {
	objectivec.IObject
}

// MLAsyncClassifierObject wraps an existing Objective-C object that conforms to the MLAsyncClassifier protocol.
type MLAsyncClassifierObject struct {
	objectivec.Object
}

func (o MLAsyncClassifierObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLAsyncClassifierObjectFromID constructs a [MLAsyncClassifierObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLAsyncClassifierObjectFromID(id objc.ID) MLAsyncClassifierObject {
	return MLAsyncClassifierObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MLAsyncClassifierObject) ClassifyOptionsCompletionHandler(classify objectivec.IObject, options objectivec.IObject, handler MLClassifierResultErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("classify:options:completionHandler:"), classify, options, handler)
}

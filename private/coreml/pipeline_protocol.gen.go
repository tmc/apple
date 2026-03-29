// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLPipeline protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLPipeline
type MLPipelineProtocol interface {
	objectivec.IObject
}

// MLPipelineProtocolObject wraps an existing Objective-C object that conforms to the MLPipelineProtocol protocol.
type MLPipelineProtocolObject struct {
	objectivec.Object
}
func (o MLPipelineProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLPipelineProtocolObjectFromID constructs a [MLPipelineProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLPipelineProtocolObjectFromID(id objc.ID) MLPipelineProtocolObject {
	return MLPipelineProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLPipeline/modelNames
func (o MLPipelineProtocolObject) ModelNames() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("modelNames"))
	return objectivec.Object{ID: rv}
	}
// See: https://developer.apple.com/documentation/CoreML/MLPipeline/models
func (o MLPipelineProtocolObject) Models() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("models"))
	return objectivec.Object{ID: rv}
	}


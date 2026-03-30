// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLClassifier protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLClassifier
type MLClassifierProtocol interface {
	objectivec.IObject
}

// MLClassifierProtocolObject wraps an existing Objective-C object that conforms to the MLClassifierProtocol protocol.
type MLClassifierProtocolObject struct {
	objectivec.Object
}

func (o MLClassifierProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLClassifierProtocolObjectFromID constructs a [MLClassifierProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLClassifierProtocolObjectFromID(id objc.ID) MLClassifierProtocolObject {
	return MLClassifierProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreML/MLClassifier/classLabels
func (o MLClassifierProtocolObject) ClassLabels() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("classLabels"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLClassifier/classify:options:error:
func (o MLClassifierProtocolObject) ClassifyOptionsError(classify objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("classify:options:error:"), classify, options)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}

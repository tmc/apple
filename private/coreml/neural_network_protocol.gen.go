// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLNeuralNetwork protocol.
type MLNeuralNetwork interface {
	objectivec.IObject

	// EvaluateError protocol.
	EvaluateError(evaluate objectivec.IObject) (objectivec.IObject, error)
}

// MLNeuralNetworkObject wraps an existing Objective-C object that conforms to the MLNeuralNetwork protocol.
type MLNeuralNetworkObject struct {
	objectivec.Object
}

func (o MLNeuralNetworkObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLNeuralNetworkObjectFromID constructs a [MLNeuralNetworkObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLNeuralNetworkObjectFromID(id objc.ID) MLNeuralNetworkObject {
	return MLNeuralNetworkObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MLNeuralNetworkObject) EvaluateError(evaluate objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("evaluate:error:"), evaluate)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}

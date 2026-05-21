// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _MLNeuralNetworkTensorDataStoring protocol.
type MLNeuralNetworkTensorDataStoring interface {
	objectivec.IObject
}

// MLNeuralNetworkTensorDataStoringObject wraps an existing Objective-C object that conforms to the MLNeuralNetworkTensorDataStoring protocol.
type MLNeuralNetworkTensorDataStoringObject struct {
	objectivec.Object
}

func (o MLNeuralNetworkTensorDataStoringObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLNeuralNetworkTensorDataStoringObjectFromID constructs a [MLNeuralNetworkTensorDataStoringObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLNeuralNetworkTensorDataStoringObjectFromID(id objc.ID) MLNeuralNetworkTensorDataStoringObject {
	return MLNeuralNetworkTensorDataStoringObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MLNeuralNetworkTensorDataStoringObject) TensorDataForOffsetExpectedLength(offset uint64, length uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tensorDataForOffset:expectedLength:"), offset, length)
	return objectivec.Object{ID: rv}
}
func (o MLNeuralNetworkTensorDataStoringObject) WriteToFileError(file objectivec.IObject) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("writeToFile:error:"), file)
	if err != nil {
		return false, err
	}
	return rv, nil
}

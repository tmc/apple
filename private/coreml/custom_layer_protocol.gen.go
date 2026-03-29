// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLCustomLayer protocol.
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomLayer
type MLCustomLayer interface {
	objectivec.IObject
}

// MLCustomLayerObject wraps an existing Objective-C object that conforms to the MLCustomLayer protocol.
type MLCustomLayerObject struct {
	objectivec.Object
}
func (o MLCustomLayerObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLCustomLayerObjectFromID constructs a [MLCustomLayerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLCustomLayerObjectFromID(id objc.ID) MLCustomLayerObject {
	return MLCustomLayerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/CoreML/MLCustomLayer/encodeToCommandBuffer:inputs:outputs:error:
func (o MLCustomLayerObject) EncodeToCommandBufferInputsOutputsError(buffer objectivec.IObject, inputs objectivec.IObject, outputs objectivec.IObject) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("encodeToCommandBuffer:inputs:outputs:error:"), buffer, inputs, outputs)
	if err != nil {
		return false, err
	}
	return rv, nil
	}
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomLayer/evaluateOnCPUWithInputs:outputs:error:
func (o MLCustomLayerObject) EvaluateOnCPUWithInputsOutputsError(inputs objectivec.IObject, outputs objectivec.IObject) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("evaluateOnCPUWithInputs:outputs:error:"), inputs, outputs)
	if err != nil {
		return false, err
	}
	return rv, nil
	}
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomLayer/outputShapesForInputShapes:error:
func (o MLCustomLayerObject) OutputShapesForInputShapesError(shapes objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("outputShapesForInputShapes:error:"), shapes)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
	}
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomLayer/setWeightData:error:
func (o MLCustomLayerObject) SetWeightDataError(data objectivec.IObject) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("setWeightData:error:"), data)
	if err != nil {
		return false, err
	}
	return rv, nil
	}


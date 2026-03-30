// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An interface that defines the behavior of a custom layer in your neural network model.
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomLayer
type MLCustomLayer interface {
	objectivec.IObject

	// Assigns the weights for the connections within the layer.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLCustomLayer/setWeightData(_:)
	SetWeightDataError(weights []foundation.NSData) (bool, error)

	// Calculates the shapes of the output of this layer for the given input shapes.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLCustomLayer/outputShapes(forInputShapes:)
	OutputShapesForInputShapesError(inputShapes []foundation.NSArray) ([]foundation.NSArray, error)

	// Evaluates the custom layer with the given inputs.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLCustomLayer/evaluate(inputs:outputs:)
	EvaluateOnCPUWithInputsOutputsError(inputs []MLMultiArray, outputs []MLMultiArray) (bool, error)
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

// Assigns the weights for the connections within the layer.
//
// weights: The data encoded in the `weights` field of the model specification.
//
// # Discussion
//
// Implement this method to assign the weights for all the connections between
// nodes in your layer. This method is called once after the initialization
// call. Your implementation should validate the weights and throw an error if
// the weights do not have the expected shape or values.
//
// The data encoded in the `weights` field of the `XCUIElementTypeMlmodel`
// file is loaded and passed into this method. If there are repeated weights
// in the `XCUIElementTypeMlmodel` file, they will be listed explicitly in the
// `weights` array. The weight values are provided in the order that they were
// defined during the custom layer conversion process. Keep a reference to the
// `weights` passed in because copying the `weights` array can significantly
// increase an app’s memory. Avoid modifying values of the weights.
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomLayer/setWeightData(_:)
func (o MLCustomLayerObject) SetWeightDataError(weights []foundation.NSData) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("setWeightData:error:"), objectivec.IObjectSliceToNSArray(weights))
	if err != nil {
		return false, err
	}
	return rv, nil
}

// Calculates the shapes of the output of this layer for the given input
// shapes.
//
// inputShapes: The shapes of the input for this layer.
//
// # Return Value
//
// The shapes of the output for the given input shapes.
//
// # Discussion
//
// Implement this method to define the layer’s interface with the rest of
// the network. It will be called at least once at load time and any time the
// size of the inputs changes in a call to [PredictionFromFeaturesError].
//
// This method consumes and returns arrays of shapes, for inputs and outputs
// of the custom layer, respectively. See the [Core ML Neural Network
// specification] for more details about shapes and how layers use them.
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomLayer/outputShapes(forInputShapes:)
//
// [Core ML Neural Network specification]: https://mlmodel.readme.io/reference/neuralnetwork
func (o MLCustomLayerObject) OutputShapesForInputShapesError(inputShapes []foundation.NSArray) ([]foundation.NSArray, error) {
	rv, err := objc.SendWithError[[]objc.ID](o.ID, objc.Sel("outputShapesForInputShapes:error:"), objectivec.IObjectSliceToNSArray(inputShapes))
	if err != nil {
		return nil, err
	}
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSArray {
		return foundation.NSArrayFromID(id)
	}), nil
}

// Evaluates the custom layer with the given inputs.
//
// inputs: The array of inputs to be evaluated.
//
// outputs: The array of outputs to be populated by evaluating the given inputs.
//
// # Discussion
//
// Implement this method to evaluate the inputs using your layer’s custom
// behavior and to populate the output arrays. It will be called for each
// evaluation of your model performed on the CPU.
//
// The memory for both input and output arrays is preallocated; don’t copy
// or move it. The inputs and outputs will have the shapes of the most recent
// call to [OutputShapesForInputShapesError]. Don’t modify the input values.
//
// Investigate [vecLib] for methods that could optimize your implementation
// significantly.
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomLayer/evaluate(inputs:outputs:)
//
// [vecLib]: https://developer.apple.com/documentation/Accelerate/veclib
func (o MLCustomLayerObject) EvaluateOnCPUWithInputsOutputsError(inputs []MLMultiArray, outputs []MLMultiArray) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("evaluateOnCPUWithInputs:outputs:error:"), objectivec.IObjectSliceToNSArray(inputs), objectivec.IObjectSliceToNSArray(outputs))
	if err != nil {
		return false, err
	}
	return rv, nil
}

// Encodes GPU commands to evaluate the custom layer.
//
// commandBuffer: A command buffer that defines the work the layer performs on the GPU.
//
// inputs: A texture array that represents the layer’s inputs.
//
// outputs: A texture array that represents the layer’s outputs.
//
// # Discussion
//
// Implement this method to use the GPU to evaluate your layer. Fill
// `commandBuffer` with the GPU commands that evaluate the layer. Don’t
// commit the command buffer in this method; Core ML executes the command
// buffer after this method returns.
//
// Improve your layer’s performance by caching the [MTLComputePipelineState]
// instances you create and intend to reuse in subsequent calls.
//
// Implementing this method doesn’t guarantee that Core ML evaluates this
// layer on the GPU. For example, Core ML may evaluate the layer on the CPU if
// the system doesn’t have enough GPU’s resources to run the custom layer.
//
// If you don’t implement this method, Core ML instead uses
// [EvaluateOnCPUWithInputsOutputsError].
//
// For more information about using the GPU for general purpose programming,
// see `Compute Processing`.
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomLayer/encode(commandBuffer:inputs:outputs:)
//
// [MTLComputePipelineState]: https://developer.apple.com/documentation/Metal/MTLComputePipelineState
func (o MLCustomLayerObject) EncodeToCommandBufferInputsOutputsError(commandBuffer metal.MTLCommandBuffer, inputs []objectivec.IObject, outputs []objectivec.IObject) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("encodeToCommandBuffer:inputs:outputs:error:"), commandBuffer, objectivec.IObjectSliceToNSArray(inputs), objectivec.IObjectSliceToNSArray(outputs))
	if err != nil {
		return false, err
	}
	return rv, nil
}

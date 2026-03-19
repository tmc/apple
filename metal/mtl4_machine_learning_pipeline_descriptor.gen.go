// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4MachineLearningPipelineDescriptor] class.
var (
	_MTL4MachineLearningPipelineDescriptorClass     MTL4MachineLearningPipelineDescriptorClass
	_MTL4MachineLearningPipelineDescriptorClassOnce sync.Once
)

func getMTL4MachineLearningPipelineDescriptorClass() MTL4MachineLearningPipelineDescriptorClass {
	_MTL4MachineLearningPipelineDescriptorClassOnce.Do(func() {
		_MTL4MachineLearningPipelineDescriptorClass = MTL4MachineLearningPipelineDescriptorClass{class: objc.GetClass("MTL4MachineLearningPipelineDescriptor")}
	})
	return _MTL4MachineLearningPipelineDescriptorClass
}

// GetMTL4MachineLearningPipelineDescriptorClass returns the class object for MTL4MachineLearningPipelineDescriptor.
func GetMTL4MachineLearningPipelineDescriptorClass() MTL4MachineLearningPipelineDescriptorClass {
	return getMTL4MachineLearningPipelineDescriptorClass()
}

type MTL4MachineLearningPipelineDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4MachineLearningPipelineDescriptorClass) Alloc() MTL4MachineLearningPipelineDescriptor {
	rv := objc.Send[MTL4MachineLearningPipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Description for a machine learning pipeline state.
//
// # Instance Properties
//
//   - [MTL4MachineLearningPipelineDescriptor.MachineLearningFunctionDescriptor]: Assigns the function that the machine learning pipeline you create from this descriptor executes.
//   - [MTL4MachineLearningPipelineDescriptor.SetMachineLearningFunctionDescriptor]
//
// # Instance Methods
//
//   - [MTL4MachineLearningPipelineDescriptor.InputDimensionsAtBufferIndex]: Obtains the dimensions of the input tensor at `bufferIndex` if set, `nil` otherwise.
//   - [MTL4MachineLearningPipelineDescriptor.Reset]: Resets the descriptor to its default values.
//   - [MTL4MachineLearningPipelineDescriptor.SetInputDimensionsAtBufferIndex]: Sets the dimension of an input tensor at a buffer index.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor
type MTL4MachineLearningPipelineDescriptor struct {
	MTL4PipelineDescriptor
}

// MTL4MachineLearningPipelineDescriptorFromID constructs a [MTL4MachineLearningPipelineDescriptor] from an objc.ID.
//
// Description for a machine learning pipeline state.
func MTL4MachineLearningPipelineDescriptorFromID(id objc.ID) MTL4MachineLearningPipelineDescriptor {
	return MTL4MachineLearningPipelineDescriptor{MTL4PipelineDescriptor: MTL4PipelineDescriptorFromID(id)}
}
// NOTE: MTL4MachineLearningPipelineDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4MachineLearningPipelineDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4MachineLearningPipelineDescriptor.MachineLearningFunctionDescriptor]: Assigns the function that the machine learning pipeline you create from this descriptor executes.
//   - [IMTL4MachineLearningPipelineDescriptor.SetMachineLearningFunctionDescriptor]
//
// # Instance Methods
//
//   - [IMTL4MachineLearningPipelineDescriptor.InputDimensionsAtBufferIndex]: Obtains the dimensions of the input tensor at `bufferIndex` if set, `nil` otherwise.
//   - [IMTL4MachineLearningPipelineDescriptor.Reset]: Resets the descriptor to its default values.
//   - [IMTL4MachineLearningPipelineDescriptor.SetInputDimensionsAtBufferIndex]: Sets the dimension of an input tensor at a buffer index.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor
type IMTL4MachineLearningPipelineDescriptor interface {
	IMTL4PipelineDescriptor

	// Topic: Instance Properties

	// Assigns the function that the machine learning pipeline you create from this descriptor executes.
	MachineLearningFunctionDescriptor() IMTL4FunctionDescriptor
	SetMachineLearningFunctionDescriptor(value IMTL4FunctionDescriptor)

	// Topic: Instance Methods

	// Obtains the dimensions of the input tensor at `bufferIndex` if set, `nil` otherwise.
	InputDimensionsAtBufferIndex(bufferIndex int) IMTLTensorExtents
	// Resets the descriptor to its default values.
	Reset()
	// Sets the dimension of an input tensor at a buffer index.
	SetInputDimensionsAtBufferIndex(dimensions IMTLTensorExtents, bufferIndex int)

	// Sets the dimensions of multiple input tensors on a range of buffer bindings.
	SetInputDimensionsWithRange(dimensions []MTLTensorExtents, range_ foundation.NSRange)
}

// Init initializes the instance.
func (m MTL4MachineLearningPipelineDescriptor) Init() MTL4MachineLearningPipelineDescriptor {
	rv := objc.Send[MTL4MachineLearningPipelineDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4MachineLearningPipelineDescriptor) Autorelease() MTL4MachineLearningPipelineDescriptor {
	rv := objc.Send[MTL4MachineLearningPipelineDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4MachineLearningPipelineDescriptor creates a new MTL4MachineLearningPipelineDescriptor instance.
func NewMTL4MachineLearningPipelineDescriptor() MTL4MachineLearningPipelineDescriptor {
	class := getMTL4MachineLearningPipelineDescriptorClass()
	rv := objc.Send[MTL4MachineLearningPipelineDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Obtains the dimensions of the input tensor at `bufferIndex` if set, `nil`
// otherwise.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/inputDimensions(bufferIndex:)
func (m MTL4MachineLearningPipelineDescriptor) InputDimensionsAtBufferIndex(bufferIndex int) IMTLTensorExtents {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputDimensionsAtBufferIndex:"), bufferIndex)
	return MTLTensorExtentsFromID(rv)
}
// Resets the descriptor to its default values.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/reset()
func (m MTL4MachineLearningPipelineDescriptor) Reset() {
	objc.Send[objc.ID](m.ID, objc.Sel("reset"))
}
// Sets the dimension of an input tensor at a buffer index.
//
// dimensions: The dimensions of the tensor.
//
// bufferIndex: Index of the tensor to modify.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/setInputDimensions(_:bufferIndex:)-34gir
func (m MTL4MachineLearningPipelineDescriptor) SetInputDimensionsAtBufferIndex(dimensions IMTLTensorExtents, bufferIndex int) {
	objc.Send[objc.ID](m.ID, objc.Sel("setInputDimensions:atBufferIndex:"), dimensions, bufferIndex)
}
// Sets the dimensions of multiple input tensors on a range of buffer
// bindings.
//
// dimensions: An array of tensor extents.
//
// range: The range of inputs of the `dimensions` argument. The range’s `length`
// needs to match the dimensions’ `count` property.
//
// # Discussion
// 
// Use this method to specify the dimensions of multiple input tensors at a
// range of indices in a single call.
// 
// You can indicate that any tensors in the range have unspecified dimensions
// by providing [NSNull] at the their corresponding index location in the
// array.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/setInputDimensions:withRange:
func (m MTL4MachineLearningPipelineDescriptor) SetInputDimensionsWithRange(dimensions []MTLTensorExtents, range_ foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("setInputDimensions:withRange:"), objectivec.IObjectSliceToNSArray(dimensions), range_)
}

// Assigns the function that the machine learning pipeline you create from
// this descriptor executes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/machineLearningFunctionDescriptor
func (m MTL4MachineLearningPipelineDescriptor) MachineLearningFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("machineLearningFunctionDescriptor"))
	return MTL4FunctionDescriptorFromID(objc.ID(rv))
}
func (m MTL4MachineLearningPipelineDescriptor) SetMachineLearningFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setMachineLearningFunctionDescriptor:"), value)
}


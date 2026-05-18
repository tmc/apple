// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A interface that represents a public shader function in a Metal library.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunction
type MTLFunction interface {
	objectivec.IObject

	// Creates an argument encoder for an argument buffer that’s one of this function’s arguments.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunction/makeArgumentEncoder(bufferIndex:)
	NewArgumentEncoderWithBufferIndex(bufferIndex uint) MTLArgumentEncoder

	// The device object that created the shader function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunction/device
	Device() MTLDevice

	// A string that identifies the shader function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunction/label
	Label() string
	SetLabel(value string)

	// The shader function’s type.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunction/functionType
	FunctionType() MTLFunctionType

	// The function’s name.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunction/name
	Name() string

	// The options that Metal used to compile this function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunction/options
	Options() MTLFunctionOptions

	// The tessellation patch type of a post-tessellation vertex function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunction/patchType
	PatchType() MTLPatchType

	// The number of patch control points in the post-tessellation vertex function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunction/patchControlPointCount
	PatchControlPointCount() int

	// An array that describes the vertex input attributes to a vertex function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunction/vertexAttributes
	VertexAttributes() []MTLVertexAttribute

	// An array that describes the input attributes to the function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunction/stageInputAttributes
	StageInputAttributes() []MTLAttribute

	// A dictionary of function constants for a specialized function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunction/functionConstantsDictionary
	FunctionConstantsDictionary() foundation.INSDictionary
}

// MTLFunctionObject wraps an existing Objective-C object that conforms to the MTLFunction protocol.
type MTLFunctionObject struct {
	objectivec.Object
}

func (o MTLFunctionObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLFunctionObjectFromID constructs a [MTLFunctionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLFunctionObjectFromID(id objc.ID) MTLFunctionObject {
	return MTLFunctionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Creates an argument encoder for an argument buffer that’s one of this
// function’s arguments.
//
// bufferIndex: The index of an argument buffer in the function’s argument list. This
// method fails if the specified index doesn’t correspond to an argument
// buffer.
//
// # Discussion
//
// Resources encoded into an argument buffer by the [MTLArgumentEncoder]
// object need to match the structure of the argument buffer located at the
// specified buffer index. If you want to interpret a regular structure as an
// argument buffer, at least one of the members of the structure needs to have
// an `[[id(n)]]` attribute.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunction/makeArgumentEncoder(bufferIndex:)
func (o MTLFunctionObject) NewArgumentEncoderWithBufferIndex(bufferIndex uint) MTLArgumentEncoder {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newArgumentEncoderWithBufferIndex:"), bufferIndex)
	return MTLArgumentEncoderObjectFromID(rv)
}

// The device object that created the shader function.
//
// # Discussion
//
// You can only use this function object with this [MTLDevice].
//
// See: https://developer.apple.com/documentation/Metal/MTLFunction/device
func (o MTLFunctionObject) Device() MTLDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
}

// A string that identifies the shader function.
//
// # Discussion
//
// Object and command labels are useful identifiers at runtime or when
// profiling and debugging your app using any Metal tool. See [Naming
// resources and commands].
//
// See: https://developer.apple.com/documentation/Metal/MTLFunction/label
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
func (o MTLFunctionObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

func (o MTLFunctionObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

// The shader function’s type.
//
// # Discussion
//
// A function’s type determines what kind of pipeline state objects you can
// create from it and whether you can use it as a callable function in a
// function table.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunction/functionType
func (o MTLFunctionObject) FunctionType() MTLFunctionType {
	rv := objc.Send[MTLFunctionType](o.ID, objc.Sel("functionType"))
	return MTLFunctionType(rv)
}

// The function’s name.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunction/name
func (o MTLFunctionObject) Name() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The options that Metal used to compile this function.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunction/options
func (o MTLFunctionObject) Options() MTLFunctionOptions {
	rv := objc.Send[MTLFunctionOptions](o.ID, objc.Sel("options"))
	return MTLFunctionOptions(rv)
}

// The tessellation patch type of a post-tessellation vertex function.
//
// # Discussion
//
// This value is [MTLPatchTypeNone] if the function isn’t a
// post-tessellation vertex function.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunction/patchType
func (o MTLFunctionObject) PatchType() MTLPatchType {
	rv := objc.Send[MTLPatchType](o.ID, objc.Sel("patchType"))
	return MTLPatchType(rv)
}

// The number of patch control points in the post-tessellation vertex
// function.
//
// # Discussion
//
// This value is `-1` if the number of patch control points wasn’t specified
// or if the function isn’t a post-tessellation vertex function.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunction/patchControlPointCount
func (o MTLFunctionObject) PatchControlPointCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("patchControlPointCount"))
	return int(rv)
}

// An array that describes the vertex input attributes to a vertex function.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunction/vertexAttributes
func (o MTLFunctionObject) VertexAttributes() []MTLVertexAttribute {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("vertexAttributes"))
	result := make([]MTLVertexAttribute, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = MTLVertexAttributeFromID(id)
	}
	return result
}

// An array that describes the input attributes to the function.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunction/stageInputAttributes
func (o MTLFunctionObject) StageInputAttributes() []MTLAttribute {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("stageInputAttributes"))
	result := make([]MTLAttribute, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = MTLAttributeFromID(id)
	}
	return result
}

// A dictionary of function constants for a specialized function.
//
// # Discussion
//
// This property returns a dictionary of the function constants that you need
// to provide to specialize this function. This property returns an empty
// dictionary if this function is already specialized or doesn’t declare any
// function constants.
//
// To create the specialized function, set these constant values in a new
// [MTLFunctionConstantValues] object and call the
// [NewFunctionWithNameConstantValuesCompletionHandler] method.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunction/functionConstantsDictionary
func (o MTLFunctionObject) FunctionConstantsDictionary() foundation.INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("functionConstantsDictionary"))
	return foundation.NSDictionaryFromID(rv)
}

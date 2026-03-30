// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTShaderProfilerGPUCommand protocol.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand
type GTShaderProfilerGPUCommand interface {
	objectivec.IObject

	// CommandBufferIndex protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/commandBufferIndex
	CommandBufferIndex() uint32

	// CommandType protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/commandType
	CommandType() uint32

	// EncoderInfoIndex protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/encoderInfoIndex
	EncoderInfoIndex() uint32

	// EncoderObjectId protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/encoderObjectId
	EncoderObjectId() uint64

	// FunctionIndex protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/functionIndex
	FunctionIndex() uint64

	// Index protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/index
	Index() uint32

	// PipelineInfoIndex protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/pipelineInfoIndex
	PipelineInfoIndex() uint32

	// PipelineStateObjectId protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/pipelineStateObjectId
	PipelineStateObjectId() uint64

	// SubCommandIndex protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/subCommandIndex
	SubCommandIndex() int
}

// GTShaderProfilerGPUCommandObject wraps an existing Objective-C object that conforms to the GTShaderProfilerGPUCommand protocol.
type GTShaderProfilerGPUCommandObject struct {
	objectivec.Object
}

func (o GTShaderProfilerGPUCommandObject) BaseObject() objectivec.Object {
	return o.Object
}

// GTShaderProfilerGPUCommandObjectFromID constructs a [GTShaderProfilerGPUCommandObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GTShaderProfilerGPUCommandObjectFromID(id objc.ID) GTShaderProfilerGPUCommandObject {
	return GTShaderProfilerGPUCommandObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/allBinaryKeys
func (o GTShaderProfilerGPUCommandObject) AllBinaryKeys() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("allBinaryKeys"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/binaryKeys
func (o GTShaderProfilerGPUCommandObject) BinaryKeys() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("binaryKeys"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/commandBufferIndex
func (o GTShaderProfilerGPUCommandObject) CommandBufferIndex() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("commandBufferIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/commandType
func (o GTShaderProfilerGPUCommandObject) CommandType() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("commandType"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/encoderInfoIndex
func (o GTShaderProfilerGPUCommandObject) EncoderInfoIndex() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("encoderInfoIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/encoderObjectId
func (o GTShaderProfilerGPUCommandObject) EncoderObjectId() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("encoderObjectId"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/functionIndex
func (o GTShaderProfilerGPUCommandObject) FunctionIndex() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("functionIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/index
func (o GTShaderProfilerGPUCommandObject) Index() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("index"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/pipelineInfoIndex
func (o GTShaderProfilerGPUCommandObject) PipelineInfoIndex() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("pipelineInfoIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/pipelineStateObjectId
func (o GTShaderProfilerGPUCommandObject) PipelineStateObjectId() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("pipelineStateObjectId"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/subCommandIndex
func (o GTShaderProfilerGPUCommandObject) SubCommandIndex() int {
	rv := objc.Send[int](o.ID, objc.Sel("subCommandIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerGPUCommand/timingInfo
func (o GTShaderProfilerGPUCommandObject) TimingInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("timingInfo"))
	return objectivec.Object{ID: rv}
}

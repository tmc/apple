// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTShaderProfilerGPUCommand protocol.
type GTShaderProfilerGPUCommand interface {
	objectivec.IObject

	// CommandBufferIndex protocol.
	CommandBufferIndex() uint32

	// CommandType protocol.
	CommandType() uint32

	// EncoderInfoIndex protocol.
	EncoderInfoIndex() uint32

	// EncoderObjectId protocol.
	EncoderObjectId() uint64

	// FunctionIndex protocol.
	FunctionIndex() uint64

	// Index protocol.
	Index() uint32

	// PipelineInfoIndex protocol.
	PipelineInfoIndex() uint32

	// PipelineStateObjectId protocol.
	PipelineStateObjectId() uint64

	// SubCommandIndex protocol.
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

func (o GTShaderProfilerGPUCommandObject) AllBinaryKeys() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("allBinaryKeys"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerGPUCommandObject) BinaryKeys() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("binaryKeys"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerGPUCommandObject) CommandBufferIndex() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("commandBufferIndex"))
	return rv
}
func (o GTShaderProfilerGPUCommandObject) CommandType() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("commandType"))
	return rv
}
func (o GTShaderProfilerGPUCommandObject) EncoderInfoIndex() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("encoderInfoIndex"))
	return rv
}
func (o GTShaderProfilerGPUCommandObject) EncoderObjectId() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("encoderObjectId"))
	return rv
}
func (o GTShaderProfilerGPUCommandObject) FunctionIndex() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("functionIndex"))
	return rv
}
func (o GTShaderProfilerGPUCommandObject) Index() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("index"))
	return rv
}
func (o GTShaderProfilerGPUCommandObject) PipelineInfoIndex() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("pipelineInfoIndex"))
	return rv
}
func (o GTShaderProfilerGPUCommandObject) PipelineStateObjectId() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("pipelineStateObjectId"))
	return rv
}
func (o GTShaderProfilerGPUCommandObject) SubCommandIndex() int {
	rv := objc.Send[int](o.ID, objc.Sel("subCommandIndex"))
	return rv
}
func (o GTShaderProfilerGPUCommandObject) TimingInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("timingInfo"))
	return objectivec.Object{ID: rv}
}

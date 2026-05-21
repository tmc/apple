// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTShaderProfilerPipelineState protocol.
type GTShaderProfilerPipelineState interface {
	objectivec.IObject

	// FunctionIndex protocol.
	FunctionIndex() uint64

	// Index protocol.
	Index() uint32

	// NumGPUCommands protocol.
	NumGPUCommands() uint32

	// ObjectId protocol.
	ObjectId() uint64

	// PointerId protocol.
	PointerId() uint64
}

// GTShaderProfilerPipelineStateObject wraps an existing Objective-C object that conforms to the GTShaderProfilerPipelineState protocol.
type GTShaderProfilerPipelineStateObject struct {
	objectivec.Object
}

func (o GTShaderProfilerPipelineStateObject) BaseObject() objectivec.Object {
	return o.Object
}

// GTShaderProfilerPipelineStateObjectFromID constructs a [GTShaderProfilerPipelineStateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GTShaderProfilerPipelineStateObjectFromID(id objc.ID) GTShaderProfilerPipelineStateObject {
	return GTShaderProfilerPipelineStateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o GTShaderProfilerPipelineStateObject) AllBinaryKeys() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("allBinaryKeys"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerPipelineStateObject) BinaryKeys() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("binaryKeys"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerPipelineStateObject) FunctionIndex() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("functionIndex"))
	return rv
}
func (o GTShaderProfilerPipelineStateObject) Index() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("index"))
	return rv
}
func (o GTShaderProfilerPipelineStateObject) NumGPUCommands() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("numGPUCommands"))
	return rv
}
func (o GTShaderProfilerPipelineStateObject) ObjectId() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("objectId"))
	return rv
}
func (o GTShaderProfilerPipelineStateObject) PointerId() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("pointerId"))
	return rv
}
func (o GTShaderProfilerPipelineStateObject) ShaderFunctions() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("shaderFunctions"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerPipelineStateObject) TimingInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("timingInfo"))
	return objectivec.Object{ID: rv}
}

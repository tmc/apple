// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTShaderProfilerEncoder protocol.
type GTShaderProfilerEncoder interface {
	objectivec.IObject

	// FunctionIndex protocol.
	FunctionIndex() uint64

	// GpuCommandStartIndex protocol.
	GpuCommandStartIndex() uint32

	// Index protocol.
	Index() uint32

	// LoadTime protocol.
	LoadTime() uint64

	// NumGPUCommands protocol.
	NumGPUCommands() uint32

	// PointerId protocol.
	PointerId() uint64

	// StoreTime protocol.
	StoreTime() uint64
}

// GTShaderProfilerEncoderObject wraps an existing Objective-C object that conforms to the GTShaderProfilerEncoder protocol.
type GTShaderProfilerEncoderObject struct {
	objectivec.Object
}

func (o GTShaderProfilerEncoderObject) BaseObject() objectivec.Object {
	return o.Object
}

// GTShaderProfilerEncoderObjectFromID constructs a [GTShaderProfilerEncoderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GTShaderProfilerEncoderObjectFromID(id objc.ID) GTShaderProfilerEncoderObject {
	return GTShaderProfilerEncoderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o GTShaderProfilerEncoderObject) FunctionIndex() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("functionIndex"))
	return rv
}
func (o GTShaderProfilerEncoderObject) GpuCommandStartIndex() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("gpuCommandStartIndex"))
	return rv
}
func (o GTShaderProfilerEncoderObject) Index() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("index"))
	return rv
}
func (o GTShaderProfilerEncoderObject) LoadTime() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("loadTime"))
	return rv
}
func (o GTShaderProfilerEncoderObject) NumGPUCommands() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("numGPUCommands"))
	return rv
}
func (o GTShaderProfilerEncoderObject) PointerId() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("pointerId"))
	return rv
}
func (o GTShaderProfilerEncoderObject) StoreTime() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("storeTime"))
	return rv
}
func (o GTShaderProfilerEncoderObject) TimingInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("timingInfo"))
	return objectivec.Object{ID: rv}
}

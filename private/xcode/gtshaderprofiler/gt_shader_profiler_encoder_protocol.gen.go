// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTShaderProfilerEncoder protocol.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerEncoder
type GTShaderProfilerEncoder interface {
	objectivec.IObject

	// FunctionIndex protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerEncoder/functionIndex
	FunctionIndex() uint64

	// GpuCommandStartIndex protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerEncoder/gpuCommandStartIndex
	GpuCommandStartIndex() uint32

	// Index protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerEncoder/index
	Index() uint32

	// LoadTime protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerEncoder/loadTime
	LoadTime() uint64

	// NumGPUCommands protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerEncoder/numGPUCommands
	NumGPUCommands() uint32

	// PointerId protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerEncoder/pointerId
	PointerId() uint64

	// StoreTime protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerEncoder/storeTime
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

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerEncoder/functionIndex
func (o GTShaderProfilerEncoderObject) FunctionIndex() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("functionIndex"))
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerEncoder/gpuCommandStartIndex
func (o GTShaderProfilerEncoderObject) GpuCommandStartIndex() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("gpuCommandStartIndex"))
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerEncoder/index
func (o GTShaderProfilerEncoderObject) Index() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("index"))
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerEncoder/loadTime
func (o GTShaderProfilerEncoderObject) LoadTime() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("loadTime"))
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerEncoder/numGPUCommands
func (o GTShaderProfilerEncoderObject) NumGPUCommands() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("numGPUCommands"))
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerEncoder/pointerId
func (o GTShaderProfilerEncoderObject) PointerId() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("pointerId"))
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerEncoder/storeTime
func (o GTShaderProfilerEncoderObject) StoreTime() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("storeTime"))
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerEncoder/timingInfo
func (o GTShaderProfilerEncoderObject) TimingInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("timingInfo"))
	return objectivec.Object{ID: rv}
	}


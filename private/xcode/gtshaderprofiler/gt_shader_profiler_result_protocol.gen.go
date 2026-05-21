// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTShaderProfilerResult protocol.
type GTShaderProfilerResult interface {
	objectivec.IObject

	// Gpu protocol.
	Gpu() uint32

	// GpuGeneration protocol.
	GpuGeneration() uint32

	// PerformanceState protocol.
	PerformanceState() uint32

	// ProfilerMode protocol.
	ProfilerMode() uint32

	// UnixTimestamp protocol.
	UnixTimestamp() int64

	// WasPerformanceStateConsistent protocol.
	WasPerformanceStateConsistent() bool
}

// GTShaderProfilerResultObject wraps an existing Objective-C object that conforms to the GTShaderProfilerResult protocol.
type GTShaderProfilerResultObject struct {
	objectivec.Object
}

func (o GTShaderProfilerResultObject) BaseObject() objectivec.Object {
	return o.Object
}

// GTShaderProfilerResultObjectFromID constructs a [GTShaderProfilerResultObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GTShaderProfilerResultObjectFromID(id objc.ID) GTShaderProfilerResultObject {
	return GTShaderProfilerResultObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o GTShaderProfilerResultObject) DerivedCountersData() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("derivedCountersData"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerResultObject) EncoderForFunctionIndex(index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("encoderForFunctionIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerResultObject) Encoders() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("encoders"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerResultObject) Gpu() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("gpu"))
	return rv
}
func (o GTShaderProfilerResultObject) GpuCommandForFunctionIndexSubCommandIndex(index uint64, index2 int) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("gpuCommandForFunctionIndex:subCommandIndex:"), index, index2)
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerResultObject) GpuCommands() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("gpuCommands"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerResultObject) GpuGeneration() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("gpuGeneration"))
	return rv
}
func (o GTShaderProfilerResultObject) GpuName(name bool) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("gpuName:"), name)
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerResultObject) McaBinaryForBinaryKey(key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("mcaBinaryForBinaryKey:"), key)
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerResultObject) MetalPluginName() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("metalPluginName"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerResultObject) MioData() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("mioData"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerResultObject) PerformanceState() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("performanceState"))
	return rv
}
func (o GTShaderProfilerResultObject) PipelineStateForId(id uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("pipelineStateForId:"), id)
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerResultObject) PipelineStates() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("pipelineStates"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerResultObject) ProfilerMode() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("profilerMode"))
	return rv
}
func (o GTShaderProfilerResultObject) ShaderBinaries() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("shaderBinaries"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerResultObject) TimingInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("timingInfo"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerResultObject) UnixTimestamp() int64 {
	rv := objc.Send[int64](o.ID, objc.Sel("unixTimestamp"))
	return rv
}
func (o GTShaderProfilerResultObject) WasPerformanceStateConsistent() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("wasPerformanceStateConsistent"))
	return rv
}

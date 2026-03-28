// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTAGX2ShaderProfilerResult] class.
var (
	_GTAGX2ShaderProfilerResultClass     GTAGX2ShaderProfilerResultClass
	_GTAGX2ShaderProfilerResultClassOnce sync.Once
)

func getGTAGX2ShaderProfilerResultClass() GTAGX2ShaderProfilerResultClass {
	_GTAGX2ShaderProfilerResultClassOnce.Do(func() {
		_GTAGX2ShaderProfilerResultClass = GTAGX2ShaderProfilerResultClass{class: objc.GetClass("GTAGX2ShaderProfilerResult")}
	})
	return _GTAGX2ShaderProfilerResultClass
}

// GetGTAGX2ShaderProfilerResultClass returns the class object for GTAGX2ShaderProfilerResult.
func GetGTAGX2ShaderProfilerResultClass() GTAGX2ShaderProfilerResultClass {
	return getGTAGX2ShaderProfilerResultClass()
}

type GTAGX2ShaderProfilerResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTAGX2ShaderProfilerResultClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTAGX2ShaderProfilerResultClass) Alloc() GTAGX2ShaderProfilerResult {
	rv := objc.Send[GTAGX2ShaderProfilerResult](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTAGX2ShaderProfilerResult._cacheObjects]
//   - [GTAGX2ShaderProfilerResult.DerivedCountersData]
//   - [GTAGX2ShaderProfilerResult.SetDerivedCountersData]
//   - [GTAGX2ShaderProfilerResult.EncodeWithCoder]
//   - [GTAGX2ShaderProfilerResult.EncoderForFunctionIndex]
//   - [GTAGX2ShaderProfilerResult.Encoders]
//   - [GTAGX2ShaderProfilerResult.SetEncoders]
//   - [GTAGX2ShaderProfilerResult.Gpu]
//   - [GTAGX2ShaderProfilerResult.GpuCommandForFunctionIndexSubCommandIndex]
//   - [GTAGX2ShaderProfilerResult.GpuCommands]
//   - [GTAGX2ShaderProfilerResult.SetGpuCommands]
//   - [GTAGX2ShaderProfilerResult.GpuGeneration]
//   - [GTAGX2ShaderProfilerResult.SetGpuGeneration]
//   - [GTAGX2ShaderProfilerResult.GpuName]
//   - [GTAGX2ShaderProfilerResult.McaBinaryForBinaryKey]
//   - [GTAGX2ShaderProfilerResult.MetalPluginName]
//   - [GTAGX2ShaderProfilerResult.SetMetalPluginName]
//   - [GTAGX2ShaderProfilerResult.MioData]
//   - [GTAGX2ShaderProfilerResult.PerformanceState]
//   - [GTAGX2ShaderProfilerResult.SetPerformanceState]
//   - [GTAGX2ShaderProfilerResult.PipelineStateForId]
//   - [GTAGX2ShaderProfilerResult.PipelineStates]
//   - [GTAGX2ShaderProfilerResult.SetPipelineStates]
//   - [GTAGX2ShaderProfilerResult.ProfilerMode]
//   - [GTAGX2ShaderProfilerResult.ShaderBinaries]
//   - [GTAGX2ShaderProfilerResult.SetShaderBinaries]
//   - [GTAGX2ShaderProfilerResult.TimelineGPUDuration]
//   - [GTAGX2ShaderProfilerResult.SetTimelineGPUDuration]
//   - [GTAGX2ShaderProfilerResult.TimingInfo]
//   - [GTAGX2ShaderProfilerResult.SetTimingInfo]
//   - [GTAGX2ShaderProfilerResult.UnixTimestamp]
//   - [GTAGX2ShaderProfilerResult.SetUnixTimestamp]
//   - [GTAGX2ShaderProfilerResult.WasPerformanceStateConsistent]
//   - [GTAGX2ShaderProfilerResult.SetWasPerformanceStateConsistent]
//   - [GTAGX2ShaderProfilerResult.InitWithCoder]
//   - [GTAGX2ShaderProfilerResult.DebugDescription]
//   - [GTAGX2ShaderProfilerResult.Description]
//   - [GTAGX2ShaderProfilerResult.Hash]
//   - [GTAGX2ShaderProfilerResult.Superclass]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult
type GTAGX2ShaderProfilerResult struct {
	objectivec.Object
}

// GTAGX2ShaderProfilerResultFromID constructs a [GTAGX2ShaderProfilerResult] from an objc.ID.
func GTAGX2ShaderProfilerResultFromID(id objc.ID) GTAGX2ShaderProfilerResult {
	return GTAGX2ShaderProfilerResult{objectivec.Object{ID: id}}
}
// Ensure GTAGX2ShaderProfilerResult implements IGTAGX2ShaderProfilerResult.
var _ IGTAGX2ShaderProfilerResult = GTAGX2ShaderProfilerResult{}

// An interface definition for the [GTAGX2ShaderProfilerResult] class.
//
// # Methods
//
//   - [IGTAGX2ShaderProfilerResult._cacheObjects]
//   - [IGTAGX2ShaderProfilerResult.DerivedCountersData]
//   - [IGTAGX2ShaderProfilerResult.SetDerivedCountersData]
//   - [IGTAGX2ShaderProfilerResult.EncodeWithCoder]
//   - [IGTAGX2ShaderProfilerResult.EncoderForFunctionIndex]
//   - [IGTAGX2ShaderProfilerResult.Encoders]
//   - [IGTAGX2ShaderProfilerResult.SetEncoders]
//   - [IGTAGX2ShaderProfilerResult.Gpu]
//   - [IGTAGX2ShaderProfilerResult.GpuCommandForFunctionIndexSubCommandIndex]
//   - [IGTAGX2ShaderProfilerResult.GpuCommands]
//   - [IGTAGX2ShaderProfilerResult.SetGpuCommands]
//   - [IGTAGX2ShaderProfilerResult.GpuGeneration]
//   - [IGTAGX2ShaderProfilerResult.SetGpuGeneration]
//   - [IGTAGX2ShaderProfilerResult.GpuName]
//   - [IGTAGX2ShaderProfilerResult.McaBinaryForBinaryKey]
//   - [IGTAGX2ShaderProfilerResult.MetalPluginName]
//   - [IGTAGX2ShaderProfilerResult.SetMetalPluginName]
//   - [IGTAGX2ShaderProfilerResult.MioData]
//   - [IGTAGX2ShaderProfilerResult.PerformanceState]
//   - [IGTAGX2ShaderProfilerResult.SetPerformanceState]
//   - [IGTAGX2ShaderProfilerResult.PipelineStateForId]
//   - [IGTAGX2ShaderProfilerResult.PipelineStates]
//   - [IGTAGX2ShaderProfilerResult.SetPipelineStates]
//   - [IGTAGX2ShaderProfilerResult.ProfilerMode]
//   - [IGTAGX2ShaderProfilerResult.ShaderBinaries]
//   - [IGTAGX2ShaderProfilerResult.SetShaderBinaries]
//   - [IGTAGX2ShaderProfilerResult.TimelineGPUDuration]
//   - [IGTAGX2ShaderProfilerResult.SetTimelineGPUDuration]
//   - [IGTAGX2ShaderProfilerResult.TimingInfo]
//   - [IGTAGX2ShaderProfilerResult.SetTimingInfo]
//   - [IGTAGX2ShaderProfilerResult.UnixTimestamp]
//   - [IGTAGX2ShaderProfilerResult.SetUnixTimestamp]
//   - [IGTAGX2ShaderProfilerResult.WasPerformanceStateConsistent]
//   - [IGTAGX2ShaderProfilerResult.SetWasPerformanceStateConsistent]
//   - [IGTAGX2ShaderProfilerResult.InitWithCoder]
//   - [IGTAGX2ShaderProfilerResult.DebugDescription]
//   - [IGTAGX2ShaderProfilerResult.Description]
//   - [IGTAGX2ShaderProfilerResult.Hash]
//   - [IGTAGX2ShaderProfilerResult.Superclass]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult
type IGTAGX2ShaderProfilerResult interface {
	objectivec.IObject

	// Topic: Methods

	_cacheObjects()
	DerivedCountersData() foundation.INSDictionary
	SetDerivedCountersData(value foundation.INSDictionary)
	EncodeWithCoder(coder foundation.INSCoder)
	EncoderForFunctionIndex(index uint64) objectivec.IObject
	Encoders() foundation.INSArray
	SetEncoders(value foundation.INSArray)
	Gpu() uint32
	GpuCommandForFunctionIndexSubCommandIndex(index uint64, index2 int) objectivec.IObject
	GpuCommands() foundation.INSArray
	SetGpuCommands(value foundation.INSArray)
	GpuGeneration() uint32
	SetGpuGeneration(value uint32)
	GpuName(name bool) objectivec.IObject
	McaBinaryForBinaryKey(key objectivec.IObject) objectivec.IObject
	MetalPluginName() string
	SetMetalPluginName(value string)
	MioData() IGTMioTraceData
	PerformanceState() uint32
	SetPerformanceState(value uint32)
	PipelineStateForId(id uint64) objectivec.IObject
	PipelineStates() foundation.INSArray
	SetPipelineStates(value foundation.INSArray)
	ProfilerMode() uint32
	ShaderBinaries() foundation.INSDictionary
	SetShaderBinaries(value foundation.INSDictionary)
	TimelineGPUDuration() uint64
	SetTimelineGPUDuration(value uint64)
	TimingInfo() IGTShaderProfilerTimingInfo
	SetTimingInfo(value IGTShaderProfilerTimingInfo)
	UnixTimestamp() int64
	SetUnixTimestamp(value int64)
	WasPerformanceStateConsistent() bool
	SetWasPerformanceStateConsistent(value bool)
	InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderProfilerResult
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (g GTAGX2ShaderProfilerResult) Init() GTAGX2ShaderProfilerResult {
	rv := objc.Send[GTAGX2ShaderProfilerResult](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTAGX2ShaderProfilerResult) Autorelease() GTAGX2ShaderProfilerResult {
	rv := objc.Send[GTAGX2ShaderProfilerResult](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTAGX2ShaderProfilerResult creates a new GTAGX2ShaderProfilerResult instance.
func NewGTAGX2ShaderProfilerResult() GTAGX2ShaderProfilerResult {
	class := getGTAGX2ShaderProfilerResultClass()
	rv := objc.Send[GTAGX2ShaderProfilerResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/initWithCoder:
func NewGTAGX2ShaderProfilerResultWithCoder(coder objectivec.IObject) GTAGX2ShaderProfilerResult {
	instance := getGTAGX2ShaderProfilerResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTAGX2ShaderProfilerResultFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/_cacheObjects
func (g GTAGX2ShaderProfilerResult) _cacheObjects() {
	objc.Send[objc.ID](g.ID, objc.Sel("_cacheObjects"))
}

// CacheObjects is an exported wrapper for the private method _cacheObjects.
func (g GTAGX2ShaderProfilerResult) CacheObjects() {
	g._cacheObjects()
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/encodeWithCoder:
func (g GTAGX2ShaderProfilerResult) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/encoderForFunctionIndex:
func (g GTAGX2ShaderProfilerResult) EncoderForFunctionIndex(index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("encoderForFunctionIndex:"), index)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/gpuCommandForFunctionIndex:subCommandIndex:
func (g GTAGX2ShaderProfilerResult) GpuCommandForFunctionIndexSubCommandIndex(index uint64, index2 int) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gpuCommandForFunctionIndex:subCommandIndex:"), index, index2)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/gpuName:
func (g GTAGX2ShaderProfilerResult) GpuName(name bool) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gpuName:"), name)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/mcaBinaryForBinaryKey:
func (g GTAGX2ShaderProfilerResult) McaBinaryForBinaryKey(key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("mcaBinaryForBinaryKey:"), key)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/pipelineStateForId:
func (g GTAGX2ShaderProfilerResult) PipelineStateForId(id uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("pipelineStateForId:"), id)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/initWithCoder:
func (g GTAGX2ShaderProfilerResult) InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderProfilerResult {
	rv := objc.Send[GTAGX2ShaderProfilerResult](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/supportsSecureCoding
func (_GTAGX2ShaderProfilerResultClass GTAGX2ShaderProfilerResultClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_GTAGX2ShaderProfilerResultClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/debugDescription
func (g GTAGX2ShaderProfilerResult) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/derivedCountersData
func (g GTAGX2ShaderProfilerResult) DerivedCountersData() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("derivedCountersData"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (g GTAGX2ShaderProfilerResult) SetDerivedCountersData(value foundation.INSDictionary) {
	objc.Send[struct{}](g.ID, objc.Sel("setDerivedCountersData:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/description
func (g GTAGX2ShaderProfilerResult) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/encoders
func (g GTAGX2ShaderProfilerResult) Encoders() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("encoders"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTAGX2ShaderProfilerResult) SetEncoders(value foundation.INSArray) {
	objc.Send[struct{}](g.ID, objc.Sel("setEncoders:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/gpu
func (g GTAGX2ShaderProfilerResult) Gpu() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("gpu"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/gpuCommands
func (g GTAGX2ShaderProfilerResult) GpuCommands() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gpuCommands"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTAGX2ShaderProfilerResult) SetGpuCommands(value foundation.INSArray) {
	objc.Send[struct{}](g.ID, objc.Sel("setGpuCommands:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/gpuGeneration
func (g GTAGX2ShaderProfilerResult) GpuGeneration() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("gpuGeneration"))
	return rv
}
func (g GTAGX2ShaderProfilerResult) SetGpuGeneration(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setGpuGeneration:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/hash
func (g GTAGX2ShaderProfilerResult) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/metalPluginName
func (g GTAGX2ShaderProfilerResult) MetalPluginName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("metalPluginName"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTAGX2ShaderProfilerResult) SetMetalPluginName(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setMetalPluginName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/mioData
func (g GTAGX2ShaderProfilerResult) MioData() IGTMioTraceData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("mioData"))
	return GTMioTraceDataFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/performanceState
func (g GTAGX2ShaderProfilerResult) PerformanceState() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("performanceState"))
	return rv
}
func (g GTAGX2ShaderProfilerResult) SetPerformanceState(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setPerformanceState:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/pipelineStates
func (g GTAGX2ShaderProfilerResult) PipelineStates() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("pipelineStates"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTAGX2ShaderProfilerResult) SetPipelineStates(value foundation.INSArray) {
	objc.Send[struct{}](g.ID, objc.Sel("setPipelineStates:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/profilerMode
func (g GTAGX2ShaderProfilerResult) ProfilerMode() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("profilerMode"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/shaderBinaries
func (g GTAGX2ShaderProfilerResult) ShaderBinaries() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("shaderBinaries"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (g GTAGX2ShaderProfilerResult) SetShaderBinaries(value foundation.INSDictionary) {
	objc.Send[struct{}](g.ID, objc.Sel("setShaderBinaries:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/superclass
func (g GTAGX2ShaderProfilerResult) Superclass() objc.Class {
	rv := objc.Send[objc.Class](g.ID, objc.Sel("superclass"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/timelineGPUDuration
func (g GTAGX2ShaderProfilerResult) TimelineGPUDuration() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("timelineGPUDuration"))
	return rv
}
func (g GTAGX2ShaderProfilerResult) SetTimelineGPUDuration(value uint64) {
	objc.Send[struct{}](g.ID, objc.Sel("setTimelineGPUDuration:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/timingInfo
func (g GTAGX2ShaderProfilerResult) TimingInfo() IGTShaderProfilerTimingInfo {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("timingInfo"))
	return GTShaderProfilerTimingInfoFromID(objc.ID(rv))
}
func (g GTAGX2ShaderProfilerResult) SetTimingInfo(value IGTShaderProfilerTimingInfo) {
	objc.Send[struct{}](g.ID, objc.Sel("setTimingInfo:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/unixTimestamp
func (g GTAGX2ShaderProfilerResult) UnixTimestamp() int64 {
	rv := objc.Send[int64](g.ID, objc.Sel("unixTimestamp"))
	return rv
}
func (g GTAGX2ShaderProfilerResult) SetUnixTimestamp(value int64) {
	objc.Send[struct{}](g.ID, objc.Sel("setUnixTimestamp:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerResult/wasPerformanceStateConsistent
func (g GTAGX2ShaderProfilerResult) WasPerformanceStateConsistent() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("wasPerformanceStateConsistent"))
	return rv
}
func (g GTAGX2ShaderProfilerResult) SetWasPerformanceStateConsistent(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setWasPerformanceStateConsistent:"), value)
}


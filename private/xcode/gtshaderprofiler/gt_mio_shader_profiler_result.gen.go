// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioShaderProfilerResult] class.
var (
	_GTMioShaderProfilerResultClass     GTMioShaderProfilerResultClass
	_GTMioShaderProfilerResultClassOnce sync.Once
)

func getGTMioShaderProfilerResultClass() GTMioShaderProfilerResultClass {
	_GTMioShaderProfilerResultClassOnce.Do(func() {
		_GTMioShaderProfilerResultClass = GTMioShaderProfilerResultClass{class: objc.GetClass("GTMioShaderProfilerResult")}
	})
	return _GTMioShaderProfilerResultClass
}

// GetGTMioShaderProfilerResultClass returns the class object for GTMioShaderProfilerResult.
func GetGTMioShaderProfilerResultClass() GTMioShaderProfilerResultClass {
	return getGTMioShaderProfilerResultClass()
}

type GTMioShaderProfilerResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioShaderProfilerResultClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioShaderProfilerResultClass) Alloc() GTMioShaderProfilerResult {
	rv := objc.Send[GTMioShaderProfilerResult](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioShaderProfilerResult._cacheObjects]
//   - [GTMioShaderProfilerResult.DerivedCountersData]
//   - [GTMioShaderProfilerResult.EncodeWithCoder]
//   - [GTMioShaderProfilerResult.EncoderForFunctionIndex]
//   - [GTMioShaderProfilerResult.Encoders]
//   - [GTMioShaderProfilerResult.SetEncoders]
//   - [GTMioShaderProfilerResult.Gpu]
//   - [GTMioShaderProfilerResult.GpuCommandForFunctionIndexSubCommandIndex]
//   - [GTMioShaderProfilerResult.GpuCommands]
//   - [GTMioShaderProfilerResult.SetGpuCommands]
//   - [GTMioShaderProfilerResult.GpuGeneration]
//   - [GTMioShaderProfilerResult.GpuName]
//   - [GTMioShaderProfilerResult.GpuTime]
//   - [GTMioShaderProfilerResult.LoadFromTraceData]
//   - [GTMioShaderProfilerResult.McaBinaryForBinaryKey]
//   - [GTMioShaderProfilerResult.MetalPluginName]
//   - [GTMioShaderProfilerResult.MioData]
//   - [GTMioShaderProfilerResult.PerformanceState]
//   - [GTMioShaderProfilerResult.PipelineStateForId]
//   - [GTMioShaderProfilerResult.PipelineStates]
//   - [GTMioShaderProfilerResult.SetPipelineStates]
//   - [GTMioShaderProfilerResult.ProfilerMode]
//   - [GTMioShaderProfilerResult.ShaderBinaries]
//   - [GTMioShaderProfilerResult.SetShaderBinaries]
//   - [GTMioShaderProfilerResult.TimingInfo]
//   - [GTMioShaderProfilerResult.UnixTimestamp]
//   - [GTMioShaderProfilerResult.WasPerformanceStateConsistent]
//   - [GTMioShaderProfilerResult.InitWithCoder]
//   - [GTMioShaderProfilerResult.InitWithTraceData]
//   - [GTMioShaderProfilerResult.DebugDescription]
//   - [GTMioShaderProfilerResult.Description]
//   - [GTMioShaderProfilerResult.Hash]
//   - [GTMioShaderProfilerResult.Superclass]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult
type GTMioShaderProfilerResult struct {
	objectivec.Object
}

// GTMioShaderProfilerResultFromID constructs a [GTMioShaderProfilerResult] from an objc.ID.
func GTMioShaderProfilerResultFromID(id objc.ID) GTMioShaderProfilerResult {
	return GTMioShaderProfilerResult{objectivec.Object{ID: id}}
}

// Ensure GTMioShaderProfilerResult implements IGTMioShaderProfilerResult.
var _ IGTMioShaderProfilerResult = GTMioShaderProfilerResult{}

// An interface definition for the [GTMioShaderProfilerResult] class.
//
// # Methods
//
//   - [IGTMioShaderProfilerResult._cacheObjects]
//   - [IGTMioShaderProfilerResult.DerivedCountersData]
//   - [IGTMioShaderProfilerResult.EncodeWithCoder]
//   - [IGTMioShaderProfilerResult.EncoderForFunctionIndex]
//   - [IGTMioShaderProfilerResult.Encoders]
//   - [IGTMioShaderProfilerResult.SetEncoders]
//   - [IGTMioShaderProfilerResult.Gpu]
//   - [IGTMioShaderProfilerResult.GpuCommandForFunctionIndexSubCommandIndex]
//   - [IGTMioShaderProfilerResult.GpuCommands]
//   - [IGTMioShaderProfilerResult.SetGpuCommands]
//   - [IGTMioShaderProfilerResult.GpuGeneration]
//   - [IGTMioShaderProfilerResult.GpuName]
//   - [IGTMioShaderProfilerResult.GpuTime]
//   - [IGTMioShaderProfilerResult.LoadFromTraceData]
//   - [IGTMioShaderProfilerResult.McaBinaryForBinaryKey]
//   - [IGTMioShaderProfilerResult.MetalPluginName]
//   - [IGTMioShaderProfilerResult.MioData]
//   - [IGTMioShaderProfilerResult.PerformanceState]
//   - [IGTMioShaderProfilerResult.PipelineStateForId]
//   - [IGTMioShaderProfilerResult.PipelineStates]
//   - [IGTMioShaderProfilerResult.SetPipelineStates]
//   - [IGTMioShaderProfilerResult.ProfilerMode]
//   - [IGTMioShaderProfilerResult.ShaderBinaries]
//   - [IGTMioShaderProfilerResult.SetShaderBinaries]
//   - [IGTMioShaderProfilerResult.TimingInfo]
//   - [IGTMioShaderProfilerResult.UnixTimestamp]
//   - [IGTMioShaderProfilerResult.WasPerformanceStateConsistent]
//   - [IGTMioShaderProfilerResult.InitWithCoder]
//   - [IGTMioShaderProfilerResult.InitWithTraceData]
//   - [IGTMioShaderProfilerResult.DebugDescription]
//   - [IGTMioShaderProfilerResult.Description]
//   - [IGTMioShaderProfilerResult.Hash]
//   - [IGTMioShaderProfilerResult.Superclass]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult
type IGTMioShaderProfilerResult interface {
	objectivec.IObject

	// Topic: Methods

	_cacheObjects()
	DerivedCountersData() foundation.INSDictionary
	EncodeWithCoder(coder foundation.INSCoder)
	EncoderForFunctionIndex(index uint64) objectivec.IObject
	Encoders() foundation.INSArray
	SetEncoders(value foundation.INSArray)
	Gpu() uint32
	GpuCommandForFunctionIndexSubCommandIndex(index uint64, index2 int) objectivec.IObject
	GpuCommands() foundation.INSArray
	SetGpuCommands(value foundation.INSArray)
	GpuGeneration() uint32
	GpuName(name bool) objectivec.IObject
	GpuTime() uint64
	LoadFromTraceData(data objectivec.IObject)
	McaBinaryForBinaryKey(key objectivec.IObject) objectivec.IObject
	MetalPluginName() string
	MioData() IGTMioTraceData
	PerformanceState() uint32
	PipelineStateForId(id uint64) objectivec.IObject
	PipelineStates() foundation.INSArray
	SetPipelineStates(value foundation.INSArray)
	ProfilerMode() uint32
	ShaderBinaries() foundation.INSDictionary
	SetShaderBinaries(value foundation.INSDictionary)
	TimingInfo() IGTShaderProfilerTimingInfo
	UnixTimestamp() int64
	WasPerformanceStateConsistent() bool
	InitWithCoder(coder foundation.INSCoder) GTMioShaderProfilerResult
	InitWithTraceData(data objectivec.IObject) GTMioShaderProfilerResult
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (g GTMioShaderProfilerResult) Init() GTMioShaderProfilerResult {
	rv := objc.Send[GTMioShaderProfilerResult](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioShaderProfilerResult) Autorelease() GTMioShaderProfilerResult {
	rv := objc.Send[GTMioShaderProfilerResult](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioShaderProfilerResult creates a new GTMioShaderProfilerResult instance.
func NewGTMioShaderProfilerResult() GTMioShaderProfilerResult {
	class := getGTMioShaderProfilerResultClass()
	rv := objc.Send[GTMioShaderProfilerResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/initWithCoder:
func NewGTMioShaderProfilerResultWithCoder(coder objectivec.IObject) GTMioShaderProfilerResult {
	instance := getGTMioShaderProfilerResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTMioShaderProfilerResultFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/initWithTraceData:
func NewGTMioShaderProfilerResultWithTraceData(data objectivec.IObject) GTMioShaderProfilerResult {
	instance := getGTMioShaderProfilerResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTraceData:"), data)
	return GTMioShaderProfilerResultFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/_cacheObjects
func (g GTMioShaderProfilerResult) _cacheObjects() {
	objc.Send[objc.ID](g.ID, objc.Sel("_cacheObjects"))
}

// CacheObjects is an exported wrapper for the private method _cacheObjects.
func (g GTMioShaderProfilerResult) CacheObjects() {
	g._cacheObjects()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/encodeWithCoder:
func (g GTMioShaderProfilerResult) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/encoderForFunctionIndex:
func (g GTMioShaderProfilerResult) EncoderForFunctionIndex(index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("encoderForFunctionIndex:"), index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/gpuCommandForFunctionIndex:subCommandIndex:
func (g GTMioShaderProfilerResult) GpuCommandForFunctionIndexSubCommandIndex(index uint64, index2 int) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gpuCommandForFunctionIndex:subCommandIndex:"), index, index2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/gpuName:
func (g GTMioShaderProfilerResult) GpuName(name bool) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gpuName:"), name)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/loadFromTraceData:
func (g GTMioShaderProfilerResult) LoadFromTraceData(data objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("loadFromTraceData:"), data)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/mcaBinaryForBinaryKey:
func (g GTMioShaderProfilerResult) McaBinaryForBinaryKey(key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("mcaBinaryForBinaryKey:"), key)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/pipelineStateForId:
func (g GTMioShaderProfilerResult) PipelineStateForId(id uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("pipelineStateForId:"), id)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/initWithCoder:
func (g GTMioShaderProfilerResult) InitWithCoder(coder foundation.INSCoder) GTMioShaderProfilerResult {
	rv := objc.Send[GTMioShaderProfilerResult](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/initWithTraceData:
func (g GTMioShaderProfilerResult) InitWithTraceData(data objectivec.IObject) GTMioShaderProfilerResult {
	rv := objc.Send[GTMioShaderProfilerResult](g.ID, objc.Sel("initWithTraceData:"), data)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/supportsSecureCoding
func (_GTMioShaderProfilerResultClass GTMioShaderProfilerResultClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_GTMioShaderProfilerResultClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/debugDescription
func (g GTMioShaderProfilerResult) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/derivedCountersData
func (g GTMioShaderProfilerResult) DerivedCountersData() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("derivedCountersData"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/description
func (g GTMioShaderProfilerResult) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/encoders
func (g GTMioShaderProfilerResult) Encoders() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("encoders"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTMioShaderProfilerResult) SetEncoders(value foundation.INSArray) {
	objc.Send[struct{}](g.ID, objc.Sel("setEncoders:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/gpu
func (g GTMioShaderProfilerResult) Gpu() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("gpu"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/gpuCommands
func (g GTMioShaderProfilerResult) GpuCommands() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gpuCommands"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTMioShaderProfilerResult) SetGpuCommands(value foundation.INSArray) {
	objc.Send[struct{}](g.ID, objc.Sel("setGpuCommands:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/gpuGeneration
func (g GTMioShaderProfilerResult) GpuGeneration() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("gpuGeneration"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/gpuTime
func (g GTMioShaderProfilerResult) GpuTime() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("gpuTime"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/hash
func (g GTMioShaderProfilerResult) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/metalPluginName
func (g GTMioShaderProfilerResult) MetalPluginName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("metalPluginName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/mioData
func (g GTMioShaderProfilerResult) MioData() IGTMioTraceData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("mioData"))
	return GTMioTraceDataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/performanceState
func (g GTMioShaderProfilerResult) PerformanceState() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("performanceState"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/pipelineStates
func (g GTMioShaderProfilerResult) PipelineStates() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("pipelineStates"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTMioShaderProfilerResult) SetPipelineStates(value foundation.INSArray) {
	objc.Send[struct{}](g.ID, objc.Sel("setPipelineStates:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/profilerMode
func (g GTMioShaderProfilerResult) ProfilerMode() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("profilerMode"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/shaderBinaries
func (g GTMioShaderProfilerResult) ShaderBinaries() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("shaderBinaries"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (g GTMioShaderProfilerResult) SetShaderBinaries(value foundation.INSDictionary) {
	objc.Send[struct{}](g.ID, objc.Sel("setShaderBinaries:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/superclass
func (g GTMioShaderProfilerResult) Superclass() objc.Class {
	rv := objc.Send[objc.Class](g.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/timingInfo
func (g GTMioShaderProfilerResult) TimingInfo() IGTShaderProfilerTimingInfo {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("timingInfo"))
	return GTShaderProfilerTimingInfoFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/unixTimestamp
func (g GTMioShaderProfilerResult) UnixTimestamp() int64 {
	rv := objc.Send[int64](g.ID, objc.Sel("unixTimestamp"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerResult/wasPerformanceStateConsistent
func (g GTMioShaderProfilerResult) WasPerformanceStateConsistent() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("wasPerformanceStateConsistent"))
	return rv
}

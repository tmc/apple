// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTAGX2ShaderProfilerPipelineState] class.
var (
	_GTAGX2ShaderProfilerPipelineStateClass     GTAGX2ShaderProfilerPipelineStateClass
	_GTAGX2ShaderProfilerPipelineStateClassOnce sync.Once
)

func getGTAGX2ShaderProfilerPipelineStateClass() GTAGX2ShaderProfilerPipelineStateClass {
	_GTAGX2ShaderProfilerPipelineStateClassOnce.Do(func() {
		_GTAGX2ShaderProfilerPipelineStateClass = GTAGX2ShaderProfilerPipelineStateClass{class: objc.GetClass("GTAGX2ShaderProfilerPipelineState")}
	})
	return _GTAGX2ShaderProfilerPipelineStateClass
}

// GetGTAGX2ShaderProfilerPipelineStateClass returns the class object for GTAGX2ShaderProfilerPipelineState.
func GetGTAGX2ShaderProfilerPipelineStateClass() GTAGX2ShaderProfilerPipelineStateClass {
	return getGTAGX2ShaderProfilerPipelineStateClass()
}

type GTAGX2ShaderProfilerPipelineStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTAGX2ShaderProfilerPipelineStateClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTAGX2ShaderProfilerPipelineStateClass) Alloc() GTAGX2ShaderProfilerPipelineState {
	rv := objc.Send[GTAGX2ShaderProfilerPipelineState](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTAGX2ShaderProfilerPipelineState.AddBinaryKeyForType]
//   - [GTAGX2ShaderProfilerPipelineState.AddFunctionForType]
//   - [GTAGX2ShaderProfilerPipelineState.AllBinaryKeys]
//   - [GTAGX2ShaderProfilerPipelineState.BinaryKeys]
//   - [GTAGX2ShaderProfilerPipelineState.ComputeTiming]
//   - [GTAGX2ShaderProfilerPipelineState.SetComputeTiming]
//   - [GTAGX2ShaderProfilerPipelineState.EncodeWithCoder]
//   - [GTAGX2ShaderProfilerPipelineState.FragmentTiming]
//   - [GTAGX2ShaderProfilerPipelineState.SetFragmentTiming]
//   - [GTAGX2ShaderProfilerPipelineState.FunctionIndex]
//   - [GTAGX2ShaderProfilerPipelineState.SetFunctionIndex]
//   - [GTAGX2ShaderProfilerPipelineState.NumGPUCommands]
//   - [GTAGX2ShaderProfilerPipelineState.SetNumGPUCommands]
//   - [GTAGX2ShaderProfilerPipelineState.ObjectId]
//   - [GTAGX2ShaderProfilerPipelineState.SetObjectId]
//   - [GTAGX2ShaderProfilerPipelineState.PointerId]
//   - [GTAGX2ShaderProfilerPipelineState.SetPointerId]
//   - [GTAGX2ShaderProfilerPipelineState.SetIndex]
//   - [GTAGX2ShaderProfilerPipelineState.ShaderFunctions]
//   - [GTAGX2ShaderProfilerPipelineState.Timing]
//   - [GTAGX2ShaderProfilerPipelineState.SetTiming]
//   - [GTAGX2ShaderProfilerPipelineState.TimingInfo]
//   - [GTAGX2ShaderProfilerPipelineState.VertexTiming]
//   - [GTAGX2ShaderProfilerPipelineState.SetVertexTiming]
//   - [GTAGX2ShaderProfilerPipelineState.InitWithCoder]
//   - [GTAGX2ShaderProfilerPipelineState.DebugDescription]
//   - [GTAGX2ShaderProfilerPipelineState.Description]
//   - [GTAGX2ShaderProfilerPipelineState.Hash]
//   - [GTAGX2ShaderProfilerPipelineState.Superclass]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState
type GTAGX2ShaderProfilerPipelineState struct {
	objectivec.Object
}

// GTAGX2ShaderProfilerPipelineStateFromID constructs a [GTAGX2ShaderProfilerPipelineState] from an objc.ID.
func GTAGX2ShaderProfilerPipelineStateFromID(id objc.ID) GTAGX2ShaderProfilerPipelineState {
	return GTAGX2ShaderProfilerPipelineState{objectivec.Object{ID: id}}
}
// Ensure GTAGX2ShaderProfilerPipelineState implements IGTAGX2ShaderProfilerPipelineState.
var _ IGTAGX2ShaderProfilerPipelineState = GTAGX2ShaderProfilerPipelineState{}

// An interface definition for the [GTAGX2ShaderProfilerPipelineState] class.
//
// # Methods
//
//   - [IGTAGX2ShaderProfilerPipelineState.AddBinaryKeyForType]
//   - [IGTAGX2ShaderProfilerPipelineState.AddFunctionForType]
//   - [IGTAGX2ShaderProfilerPipelineState.AllBinaryKeys]
//   - [IGTAGX2ShaderProfilerPipelineState.BinaryKeys]
//   - [IGTAGX2ShaderProfilerPipelineState.ComputeTiming]
//   - [IGTAGX2ShaderProfilerPipelineState.SetComputeTiming]
//   - [IGTAGX2ShaderProfilerPipelineState.EncodeWithCoder]
//   - [IGTAGX2ShaderProfilerPipelineState.FragmentTiming]
//   - [IGTAGX2ShaderProfilerPipelineState.SetFragmentTiming]
//   - [IGTAGX2ShaderProfilerPipelineState.FunctionIndex]
//   - [IGTAGX2ShaderProfilerPipelineState.SetFunctionIndex]
//   - [IGTAGX2ShaderProfilerPipelineState.NumGPUCommands]
//   - [IGTAGX2ShaderProfilerPipelineState.SetNumGPUCommands]
//   - [IGTAGX2ShaderProfilerPipelineState.ObjectId]
//   - [IGTAGX2ShaderProfilerPipelineState.SetObjectId]
//   - [IGTAGX2ShaderProfilerPipelineState.PointerId]
//   - [IGTAGX2ShaderProfilerPipelineState.SetPointerId]
//   - [IGTAGX2ShaderProfilerPipelineState.SetIndex]
//   - [IGTAGX2ShaderProfilerPipelineState.ShaderFunctions]
//   - [IGTAGX2ShaderProfilerPipelineState.Timing]
//   - [IGTAGX2ShaderProfilerPipelineState.SetTiming]
//   - [IGTAGX2ShaderProfilerPipelineState.TimingInfo]
//   - [IGTAGX2ShaderProfilerPipelineState.VertexTiming]
//   - [IGTAGX2ShaderProfilerPipelineState.SetVertexTiming]
//   - [IGTAGX2ShaderProfilerPipelineState.InitWithCoder]
//   - [IGTAGX2ShaderProfilerPipelineState.DebugDescription]
//   - [IGTAGX2ShaderProfilerPipelineState.Description]
//   - [IGTAGX2ShaderProfilerPipelineState.Hash]
//   - [IGTAGX2ShaderProfilerPipelineState.Superclass]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState
type IGTAGX2ShaderProfilerPipelineState interface {
	objectivec.IObject

	// Topic: Methods

	AddBinaryKeyForType(key objectivec.IObject, type_ uint32)
	AddFunctionForType(function objectivec.IObject, type_ uint32)
	AllBinaryKeys() foundation.INSDictionary
	BinaryKeys() foundation.INSDictionary
	ComputeTiming() IGTAGX2ShaderProfilerTiming
	SetComputeTiming(value IGTAGX2ShaderProfilerTiming)
	EncodeWithCoder(coder foundation.INSCoder)
	FragmentTiming() IGTAGX2ShaderProfilerTiming
	SetFragmentTiming(value IGTAGX2ShaderProfilerTiming)
	FunctionIndex() uint64
	SetFunctionIndex(value uint64)
	NumGPUCommands() uint32
	SetNumGPUCommands(value uint32)
	ObjectId() uint64
	SetObjectId(value uint64)
	PointerId() uint64
	SetPointerId(value uint64)
	SetIndex(index uint32)
	ShaderFunctions() foundation.INSDictionary
	Timing() IGTAGX2ShaderProfilerTiming
	SetTiming(value IGTAGX2ShaderProfilerTiming)
	TimingInfo() IGTShaderProfilerTimingInfo
	VertexTiming() IGTAGX2ShaderProfilerTiming
	SetVertexTiming(value IGTAGX2ShaderProfilerTiming)
	InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderProfilerPipelineState
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (g GTAGX2ShaderProfilerPipelineState) Init() GTAGX2ShaderProfilerPipelineState {
	rv := objc.Send[GTAGX2ShaderProfilerPipelineState](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTAGX2ShaderProfilerPipelineState) Autorelease() GTAGX2ShaderProfilerPipelineState {
	rv := objc.Send[GTAGX2ShaderProfilerPipelineState](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTAGX2ShaderProfilerPipelineState creates a new GTAGX2ShaderProfilerPipelineState instance.
func NewGTAGX2ShaderProfilerPipelineState() GTAGX2ShaderProfilerPipelineState {
	class := getGTAGX2ShaderProfilerPipelineStateClass()
	rv := objc.Send[GTAGX2ShaderProfilerPipelineState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/initWithCoder:
func NewGTAGX2ShaderProfilerPipelineStateWithCoder(coder objectivec.IObject) GTAGX2ShaderProfilerPipelineState {
	instance := getGTAGX2ShaderProfilerPipelineStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTAGX2ShaderProfilerPipelineStateFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/addBinaryKey:forType:
func (g GTAGX2ShaderProfilerPipelineState) AddBinaryKeyForType(key objectivec.IObject, type_ uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("addBinaryKey:forType:"), key, type_)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/addFunction:forType:
func (g GTAGX2ShaderProfilerPipelineState) AddFunctionForType(function objectivec.IObject, type_ uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("addFunction:forType:"), function, type_)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/encodeWithCoder:
func (g GTAGX2ShaderProfilerPipelineState) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/setIndex:
func (g GTAGX2ShaderProfilerPipelineState) SetIndex(index uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("setIndex:"), index)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/initWithCoder:
func (g GTAGX2ShaderProfilerPipelineState) InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderProfilerPipelineState {
	rv := objc.Send[GTAGX2ShaderProfilerPipelineState](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/supportsSecureCoding
func (_GTAGX2ShaderProfilerPipelineStateClass GTAGX2ShaderProfilerPipelineStateClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_GTAGX2ShaderProfilerPipelineStateClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/allBinaryKeys
func (g GTAGX2ShaderProfilerPipelineState) AllBinaryKeys() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("allBinaryKeys"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/binaryKeys
func (g GTAGX2ShaderProfilerPipelineState) BinaryKeys() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("binaryKeys"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/computeTiming
func (g GTAGX2ShaderProfilerPipelineState) ComputeTiming() IGTAGX2ShaderProfilerTiming {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("computeTiming"))
	return GTAGX2ShaderProfilerTimingFromID(objc.ID(rv))
}
func (g GTAGX2ShaderProfilerPipelineState) SetComputeTiming(value IGTAGX2ShaderProfilerTiming) {
	objc.Send[struct{}](g.ID, objc.Sel("setComputeTiming:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/debugDescription
func (g GTAGX2ShaderProfilerPipelineState) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/description
func (g GTAGX2ShaderProfilerPipelineState) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/fragmentTiming
func (g GTAGX2ShaderProfilerPipelineState) FragmentTiming() IGTAGX2ShaderProfilerTiming {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("fragmentTiming"))
	return GTAGX2ShaderProfilerTimingFromID(objc.ID(rv))
}
func (g GTAGX2ShaderProfilerPipelineState) SetFragmentTiming(value IGTAGX2ShaderProfilerTiming) {
	objc.Send[struct{}](g.ID, objc.Sel("setFragmentTiming:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/functionIndex
func (g GTAGX2ShaderProfilerPipelineState) FunctionIndex() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("functionIndex"))
	return rv
}
func (g GTAGX2ShaderProfilerPipelineState) SetFunctionIndex(value uint64) {
	objc.Send[struct{}](g.ID, objc.Sel("setFunctionIndex:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/hash
func (g GTAGX2ShaderProfilerPipelineState) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/numGPUCommands
func (g GTAGX2ShaderProfilerPipelineState) NumGPUCommands() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("numGPUCommands"))
	return rv
}
func (g GTAGX2ShaderProfilerPipelineState) SetNumGPUCommands(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setNumGPUCommands:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/objectId
func (g GTAGX2ShaderProfilerPipelineState) ObjectId() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("objectId"))
	return rv
}
func (g GTAGX2ShaderProfilerPipelineState) SetObjectId(value uint64) {
	objc.Send[struct{}](g.ID, objc.Sel("setObjectId:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/pointerId
func (g GTAGX2ShaderProfilerPipelineState) PointerId() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("pointerId"))
	return rv
}
func (g GTAGX2ShaderProfilerPipelineState) SetPointerId(value uint64) {
	objc.Send[struct{}](g.ID, objc.Sel("setPointerId:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/shaderFunctions
func (g GTAGX2ShaderProfilerPipelineState) ShaderFunctions() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("shaderFunctions"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/superclass
func (g GTAGX2ShaderProfilerPipelineState) Superclass() objc.Class {
	rv := objc.Send[objc.Class](g.ID, objc.Sel("superclass"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/timing
func (g GTAGX2ShaderProfilerPipelineState) Timing() IGTAGX2ShaderProfilerTiming {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("timing"))
	return GTAGX2ShaderProfilerTimingFromID(objc.ID(rv))
}
func (g GTAGX2ShaderProfilerPipelineState) SetTiming(value IGTAGX2ShaderProfilerTiming) {
	objc.Send[struct{}](g.ID, objc.Sel("setTiming:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/timingInfo
func (g GTAGX2ShaderProfilerPipelineState) TimingInfo() IGTShaderProfilerTimingInfo {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("timingInfo"))
	return GTShaderProfilerTimingInfoFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerPipelineState/vertexTiming
func (g GTAGX2ShaderProfilerPipelineState) VertexTiming() IGTAGX2ShaderProfilerTiming {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("vertexTiming"))
	return GTAGX2ShaderProfilerTimingFromID(objc.ID(rv))
}
func (g GTAGX2ShaderProfilerPipelineState) SetVertexTiming(value IGTAGX2ShaderProfilerTiming) {
	objc.Send[struct{}](g.ID, objc.Sel("setVertexTiming:"), value)
}


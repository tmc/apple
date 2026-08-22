// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioShaderProfilerPipelineState] class.
var (
	_GTMioShaderProfilerPipelineStateClass     GTMioShaderProfilerPipelineStateClass
	_GTMioShaderProfilerPipelineStateClassOnce sync.Once
)

func getGTMioShaderProfilerPipelineStateClass() GTMioShaderProfilerPipelineStateClass {
	_GTMioShaderProfilerPipelineStateClassOnce.Do(func() {
		_GTMioShaderProfilerPipelineStateClass = GTMioShaderProfilerPipelineStateClass{class: objc.GetClass("GTMioShaderProfilerPipelineState")}
	})
	return _GTMioShaderProfilerPipelineStateClass
}

// GetGTMioShaderProfilerPipelineStateClass returns the class object for GTMioShaderProfilerPipelineState.
func GetGTMioShaderProfilerPipelineStateClass() GTMioShaderProfilerPipelineStateClass {
	return getGTMioShaderProfilerPipelineStateClass()
}

type GTMioShaderProfilerPipelineStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioShaderProfilerPipelineStateClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioShaderProfilerPipelineStateClass) Alloc() GTMioShaderProfilerPipelineState {
	rv := objc.SendIfResponds[GTMioShaderProfilerPipelineState](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioShaderProfilerPipelineState.AddBinaryKeyForType]
//   - [GTMioShaderProfilerPipelineState.AddFunctionForType]
//   - [GTMioShaderProfilerPipelineState.AllBinaryKeys]
//   - [GTMioShaderProfilerPipelineState.BinaryKeys]
//   - [GTMioShaderProfilerPipelineState.FunctionIndex]
//   - [GTMioShaderProfilerPipelineState.Index]
//   - [GTMioShaderProfilerPipelineState.NumGPUCommands]
//   - [GTMioShaderProfilerPipelineState.ObjectId]
//   - [GTMioShaderProfilerPipelineState.PointerId]
//   - [GTMioShaderProfilerPipelineState.ShaderFunctions]
//   - [GTMioShaderProfilerPipelineState.TimingInfo]
//   - [GTMioShaderProfilerPipelineState.InitWithInfoTraceData]
//   - [GTMioShaderProfilerPipelineState.DebugDescription]
//   - [GTMioShaderProfilerPipelineState.Description]
//   - [GTMioShaderProfilerPipelineState.Hash]
//   - [GTMioShaderProfilerPipelineState.Superclass]
type GTMioShaderProfilerPipelineState struct {
	objectivec.Object
}

// GTMioShaderProfilerPipelineStateFromID constructs a [GTMioShaderProfilerPipelineState] from an objc.ID.
func GTMioShaderProfilerPipelineStateFromID(id objc.ID) GTMioShaderProfilerPipelineState {
	return GTMioShaderProfilerPipelineState{objectivec.Object{ID: id}}
}

// Ensure GTMioShaderProfilerPipelineState implements IGTMioShaderProfilerPipelineState.
var _ IGTMioShaderProfilerPipelineState = GTMioShaderProfilerPipelineState{}

// An interface definition for the [GTMioShaderProfilerPipelineState] class.
//
// # Methods
//
//   - [IGTMioShaderProfilerPipelineState.AddBinaryKeyForType]
//   - [IGTMioShaderProfilerPipelineState.AddFunctionForType]
//   - [IGTMioShaderProfilerPipelineState.AllBinaryKeys]
//   - [IGTMioShaderProfilerPipelineState.BinaryKeys]
//   - [IGTMioShaderProfilerPipelineState.FunctionIndex]
//   - [IGTMioShaderProfilerPipelineState.Index]
//   - [IGTMioShaderProfilerPipelineState.NumGPUCommands]
//   - [IGTMioShaderProfilerPipelineState.ObjectId]
//   - [IGTMioShaderProfilerPipelineState.PointerId]
//   - [IGTMioShaderProfilerPipelineState.ShaderFunctions]
//   - [IGTMioShaderProfilerPipelineState.TimingInfo]
//   - [IGTMioShaderProfilerPipelineState.InitWithInfoTraceData]
//   - [IGTMioShaderProfilerPipelineState.DebugDescription]
//   - [IGTMioShaderProfilerPipelineState.Description]
//   - [IGTMioShaderProfilerPipelineState.Hash]
//   - [IGTMioShaderProfilerPipelineState.Superclass]
type IGTMioShaderProfilerPipelineState interface {
	objectivec.IObject

	// Topic: Methods

	AddBinaryKeyForType(key objectivec.IObject, type_ uint32)
	AddFunctionForType(function objectivec.IObject, type_ uint32)
	AllBinaryKeys() foundation.INSDictionary
	BinaryKeys() foundation.INSDictionary
	FunctionIndex() uint64
	Index() uint32
	NumGPUCommands() uint32
	ObjectId() uint64
	PointerId() uint64
	ShaderFunctions() foundation.INSDictionary
	TimingInfo() IGTShaderProfilerTimingInfo
	InitWithInfoTraceData(info unsafe.Pointer, data objectivec.IObject) GTMioShaderProfilerPipelineState
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (g GTMioShaderProfilerPipelineState) Init() GTMioShaderProfilerPipelineState {
	rv := objc.SendIfResponds[GTMioShaderProfilerPipelineState](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioShaderProfilerPipelineState) Autorelease() GTMioShaderProfilerPipelineState {
	rv := objc.SendIfResponds[GTMioShaderProfilerPipelineState](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioShaderProfilerPipelineState creates a new GTMioShaderProfilerPipelineState instance.
func NewGTMioShaderProfilerPipelineState() GTMioShaderProfilerPipelineState {
	class := getGTMioShaderProfilerPipelineStateClass()
	rv := objc.SendIfResponds[GTMioShaderProfilerPipelineState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTMioShaderProfilerPipelineStateWithInfoTraceData(info unsafe.Pointer, data objectivec.IObject) GTMioShaderProfilerPipelineState {
	instance := getGTMioShaderProfilerPipelineStateClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithInfo:traceData:"), info, data)
	return GTMioShaderProfilerPipelineStateFromID(rv)
}

func (g GTMioShaderProfilerPipelineState) AddBinaryKeyForType(key objectivec.IObject, type_ uint32) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("addBinaryKey:forType:"), key, type_)
}
func (g GTMioShaderProfilerPipelineState) AddFunctionForType(function objectivec.IObject, type_ uint32) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("addFunction:forType:"), function, type_)
}
func (g GTMioShaderProfilerPipelineState) InitWithInfoTraceData(info unsafe.Pointer, data objectivec.IObject) GTMioShaderProfilerPipelineState {
	rv := objc.SendIfResponds[GTMioShaderProfilerPipelineState](g.ID, objc.Sel("initWithInfo:traceData:"), info, data)
	return rv
}

func (g GTMioShaderProfilerPipelineState) AllBinaryKeys() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("allBinaryKeys"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (g GTMioShaderProfilerPipelineState) BinaryKeys() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("binaryKeys"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (g GTMioShaderProfilerPipelineState) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTMioShaderProfilerPipelineState) Description() string {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTMioShaderProfilerPipelineState) FunctionIndex() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("functionIndex"))
	return rv
}
func (g GTMioShaderProfilerPipelineState) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("hash"))
	return rv
}
func (g GTMioShaderProfilerPipelineState) Index() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("index"))
	return rv
}
func (g GTMioShaderProfilerPipelineState) NumGPUCommands() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("numGPUCommands"))
	return rv
}
func (g GTMioShaderProfilerPipelineState) ObjectId() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("objectId"))
	return rv
}
func (g GTMioShaderProfilerPipelineState) PointerId() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("pointerId"))
	return rv
}
func (g GTMioShaderProfilerPipelineState) ShaderFunctions() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("shaderFunctions"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (g GTMioShaderProfilerPipelineState) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](g.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (g GTMioShaderProfilerPipelineState) TimingInfo() IGTShaderProfilerTimingInfo {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("timingInfo"))
	return GTShaderProfilerTimingInfoFromID(objc.ID(rv))
}

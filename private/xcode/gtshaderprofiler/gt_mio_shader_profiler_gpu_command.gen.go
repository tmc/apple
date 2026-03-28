// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioShaderProfilerGPUCommand] class.
var (
	_GTMioShaderProfilerGPUCommandClass     GTMioShaderProfilerGPUCommandClass
	_GTMioShaderProfilerGPUCommandClassOnce sync.Once
)

func getGTMioShaderProfilerGPUCommandClass() GTMioShaderProfilerGPUCommandClass {
	_GTMioShaderProfilerGPUCommandClassOnce.Do(func() {
		_GTMioShaderProfilerGPUCommandClass = GTMioShaderProfilerGPUCommandClass{class: objc.GetClass("GTMioShaderProfilerGPUCommand")}
	})
	return _GTMioShaderProfilerGPUCommandClass
}

// GetGTMioShaderProfilerGPUCommandClass returns the class object for GTMioShaderProfilerGPUCommand.
func GetGTMioShaderProfilerGPUCommandClass() GTMioShaderProfilerGPUCommandClass {
	return getGTMioShaderProfilerGPUCommandClass()
}

type GTMioShaderProfilerGPUCommandClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioShaderProfilerGPUCommandClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioShaderProfilerGPUCommandClass) Alloc() GTMioShaderProfilerGPUCommand {
	rv := objc.Send[GTMioShaderProfilerGPUCommand](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTMioShaderProfilerGPUCommand.AddBinaryKeyForType]
//   - [GTMioShaderProfilerGPUCommand.AllBinaryKeys]
//   - [GTMioShaderProfilerGPUCommand.BinaryKeys]
//   - [GTMioShaderProfilerGPUCommand.CommandBufferIndex]
//   - [GTMioShaderProfilerGPUCommand.CommandType]
//   - [GTMioShaderProfilerGPUCommand.EncoderInfoIndex]
//   - [GTMioShaderProfilerGPUCommand.EncoderObjectId]
//   - [GTMioShaderProfilerGPUCommand.FunctionIndex]
//   - [GTMioShaderProfilerGPUCommand.PipelineInfoIndex]
//   - [GTMioShaderProfilerGPUCommand.PipelineStateObjectId]
//   - [GTMioShaderProfilerGPUCommand.SubCommandIndex]
//   - [GTMioShaderProfilerGPUCommand.TimingInfo]
//   - [GTMioShaderProfilerGPUCommand.InitWithMioGPUCommandStreamGPUCommandTraceData]
//   - [GTMioShaderProfilerGPUCommand.DebugDescription]
//   - [GTMioShaderProfilerGPUCommand.Description]
//   - [GTMioShaderProfilerGPUCommand.Hash]
//   - [GTMioShaderProfilerGPUCommand.Superclass]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand
type GTMioShaderProfilerGPUCommand struct {
	objectivec.Object
}

// GTMioShaderProfilerGPUCommandFromID constructs a [GTMioShaderProfilerGPUCommand] from an objc.ID.
func GTMioShaderProfilerGPUCommandFromID(id objc.ID) GTMioShaderProfilerGPUCommand {
	return GTMioShaderProfilerGPUCommand{objectivec.Object{ID: id}}
}
// Ensure GTMioShaderProfilerGPUCommand implements IGTMioShaderProfilerGPUCommand.
var _ IGTMioShaderProfilerGPUCommand = GTMioShaderProfilerGPUCommand{}

// An interface definition for the [GTMioShaderProfilerGPUCommand] class.
//
// # Methods
//
//   - [IGTMioShaderProfilerGPUCommand.AddBinaryKeyForType]
//   - [IGTMioShaderProfilerGPUCommand.AllBinaryKeys]
//   - [IGTMioShaderProfilerGPUCommand.BinaryKeys]
//   - [IGTMioShaderProfilerGPUCommand.CommandBufferIndex]
//   - [IGTMioShaderProfilerGPUCommand.CommandType]
//   - [IGTMioShaderProfilerGPUCommand.EncoderInfoIndex]
//   - [IGTMioShaderProfilerGPUCommand.EncoderObjectId]
//   - [IGTMioShaderProfilerGPUCommand.FunctionIndex]
//   - [IGTMioShaderProfilerGPUCommand.PipelineInfoIndex]
//   - [IGTMioShaderProfilerGPUCommand.PipelineStateObjectId]
//   - [IGTMioShaderProfilerGPUCommand.SubCommandIndex]
//   - [IGTMioShaderProfilerGPUCommand.TimingInfo]
//   - [IGTMioShaderProfilerGPUCommand.InitWithMioGPUCommandStreamGPUCommandTraceData]
//   - [IGTMioShaderProfilerGPUCommand.DebugDescription]
//   - [IGTMioShaderProfilerGPUCommand.Description]
//   - [IGTMioShaderProfilerGPUCommand.Hash]
//   - [IGTMioShaderProfilerGPUCommand.Superclass]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand
type IGTMioShaderProfilerGPUCommand interface {
	objectivec.IObject

	// Topic: Methods

	AddBinaryKeyForType(key objectivec.IObject, type_ uint32)
	AllBinaryKeys() foundation.INSDictionary
	BinaryKeys() foundation.INSDictionary
	CommandBufferIndex() uint32
	CommandType() uint32
	EncoderInfoIndex() uint32
	EncoderObjectId() uint64
	FunctionIndex() uint64
	PipelineInfoIndex() uint32
	PipelineStateObjectId() uint64
	SubCommandIndex() int
	TimingInfo() IGTShaderProfilerTimingInfo
	InitWithMioGPUCommandStreamGPUCommandTraceData(gPUCommand unsafe.Pointer, gPUCommand2 objectivec.IObject, data objectivec.IObject) GTMioShaderProfilerGPUCommand
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (g GTMioShaderProfilerGPUCommand) Init() GTMioShaderProfilerGPUCommand {
	rv := objc.Send[GTMioShaderProfilerGPUCommand](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioShaderProfilerGPUCommand) Autorelease() GTMioShaderProfilerGPUCommand {
	rv := objc.Send[GTMioShaderProfilerGPUCommand](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioShaderProfilerGPUCommand creates a new GTMioShaderProfilerGPUCommand instance.
func NewGTMioShaderProfilerGPUCommand() GTMioShaderProfilerGPUCommand {
	class := getGTMioShaderProfilerGPUCommandClass()
	rv := objc.Send[GTMioShaderProfilerGPUCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand/initWithMioGPUCommand:streamGPUCommand:traceData:
func NewGTMioShaderProfilerGPUCommandWithMioGPUCommandStreamGPUCommandTraceData(gPUCommand unsafe.Pointer, gPUCommand2 objectivec.IObject, data objectivec.IObject) GTMioShaderProfilerGPUCommand {
	instance := getGTMioShaderProfilerGPUCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMioGPUCommand:streamGPUCommand:traceData:"), gPUCommand, gPUCommand2, data)
	return GTMioShaderProfilerGPUCommandFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand/addBinaryKey:forType:
func (g GTMioShaderProfilerGPUCommand) AddBinaryKeyForType(key objectivec.IObject, type_ uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("addBinaryKey:forType:"), key, type_)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand/initWithMioGPUCommand:streamGPUCommand:traceData:
func (g GTMioShaderProfilerGPUCommand) InitWithMioGPUCommandStreamGPUCommandTraceData(gPUCommand unsafe.Pointer, gPUCommand2 objectivec.IObject, data objectivec.IObject) GTMioShaderProfilerGPUCommand {
	rv := objc.Send[GTMioShaderProfilerGPUCommand](g.ID, objc.Sel("initWithMioGPUCommand:streamGPUCommand:traceData:"), gPUCommand, gPUCommand2, data)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand/allBinaryKeys
func (g GTMioShaderProfilerGPUCommand) AllBinaryKeys() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("allBinaryKeys"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand/binaryKeys
func (g GTMioShaderProfilerGPUCommand) BinaryKeys() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("binaryKeys"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand/commandBufferIndex
func (g GTMioShaderProfilerGPUCommand) CommandBufferIndex() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("commandBufferIndex"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand/commandType
func (g GTMioShaderProfilerGPUCommand) CommandType() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("commandType"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand/debugDescription
func (g GTMioShaderProfilerGPUCommand) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand/description
func (g GTMioShaderProfilerGPUCommand) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand/encoderInfoIndex
func (g GTMioShaderProfilerGPUCommand) EncoderInfoIndex() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("encoderInfoIndex"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand/encoderObjectId
func (g GTMioShaderProfilerGPUCommand) EncoderObjectId() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("encoderObjectId"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand/functionIndex
func (g GTMioShaderProfilerGPUCommand) FunctionIndex() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("functionIndex"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand/hash
func (g GTMioShaderProfilerGPUCommand) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand/pipelineInfoIndex
func (g GTMioShaderProfilerGPUCommand) PipelineInfoIndex() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("pipelineInfoIndex"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand/pipelineStateObjectId
func (g GTMioShaderProfilerGPUCommand) PipelineStateObjectId() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("pipelineStateObjectId"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand/subCommandIndex
func (g GTMioShaderProfilerGPUCommand) SubCommandIndex() int {
	rv := objc.Send[int](g.ID, objc.Sel("subCommandIndex"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand/superclass
func (g GTMioShaderProfilerGPUCommand) Superclass() objc.Class {
	rv := objc.Send[objc.Class](g.ID, objc.Sel("superclass"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerGPUCommand/timingInfo
func (g GTMioShaderProfilerGPUCommand) TimingInfo() IGTShaderProfilerTimingInfo {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("timingInfo"))
	return GTShaderProfilerTimingInfoFromID(objc.ID(rv))
}


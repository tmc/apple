// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
//   - [GTMioShaderProfilerGPUCommand.Index]
//   - [GTMioShaderProfilerGPUCommand.PipelineInfoIndex]
//   - [GTMioShaderProfilerGPUCommand.PipelineStateObjectId]
//   - [GTMioShaderProfilerGPUCommand.SubCommandIndex]
//   - [GTMioShaderProfilerGPUCommand.TimingInfo]
//   - [GTMioShaderProfilerGPUCommand.InitWithMioGPUCommandStreamGPUCommandTraceData]
//   - [GTMioShaderProfilerGPUCommand.DebugDescription]
//   - [GTMioShaderProfilerGPUCommand.Description]
//   - [GTMioShaderProfilerGPUCommand.Hash]
//   - [GTMioShaderProfilerGPUCommand.Superclass]
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
//   - [IGTMioShaderProfilerGPUCommand.Index]
//   - [IGTMioShaderProfilerGPUCommand.PipelineInfoIndex]
//   - [IGTMioShaderProfilerGPUCommand.PipelineStateObjectId]
//   - [IGTMioShaderProfilerGPUCommand.SubCommandIndex]
//   - [IGTMioShaderProfilerGPUCommand.TimingInfo]
//   - [IGTMioShaderProfilerGPUCommand.InitWithMioGPUCommandStreamGPUCommandTraceData]
//   - [IGTMioShaderProfilerGPUCommand.DebugDescription]
//   - [IGTMioShaderProfilerGPUCommand.Description]
//   - [IGTMioShaderProfilerGPUCommand.Hash]
//   - [IGTMioShaderProfilerGPUCommand.Superclass]
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
	Index() uint32
	PipelineInfoIndex() uint32
	PipelineStateObjectId() uint64
	SubCommandIndex() int
	TimingInfo() IGTShaderProfilerTimingInfo
	InitWithMioGPUCommandStreamGPUCommandTraceData(gPUCommand *GTMioDrawMetadata, gPUCommand2 unsafe.Pointer, data objectivec.IObject) GTMioShaderProfilerGPUCommand
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
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

func NewGTMioShaderProfilerGPUCommandWithMioGPUCommandStreamGPUCommandTraceData(gPUCommand *GTMioDrawMetadata, gPUCommand2 unsafe.Pointer, data objectivec.IObject) GTMioShaderProfilerGPUCommand {
	instance := getGTMioShaderProfilerGPUCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMioGPUCommand:streamGPUCommand:traceData:"), gPUCommand, gPUCommand2, data)
	return GTMioShaderProfilerGPUCommandFromID(rv)
}

func (g GTMioShaderProfilerGPUCommand) AddBinaryKeyForType(key objectivec.IObject, type_ uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("addBinaryKey:forType:"), key, type_)
}
func (g GTMioShaderProfilerGPUCommand) InitWithMioGPUCommandStreamGPUCommandTraceData(gPUCommand *GTMioDrawMetadata, gPUCommand2 unsafe.Pointer, data objectivec.IObject) GTMioShaderProfilerGPUCommand {
	rv := objc.Send[GTMioShaderProfilerGPUCommand](g.ID, objc.Sel("initWithMioGPUCommand:streamGPUCommand:traceData:"), gPUCommand, gPUCommand2, data)
	return rv
}

func (g GTMioShaderProfilerGPUCommand) AllBinaryKeys() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("allBinaryKeys"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (g GTMioShaderProfilerGPUCommand) BinaryKeys() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("binaryKeys"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (g GTMioShaderProfilerGPUCommand) CommandBufferIndex() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("commandBufferIndex"))
	return rv
}
func (g GTMioShaderProfilerGPUCommand) CommandType() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("commandType"))
	return rv
}
func (g GTMioShaderProfilerGPUCommand) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTMioShaderProfilerGPUCommand) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTMioShaderProfilerGPUCommand) EncoderInfoIndex() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("encoderInfoIndex"))
	return rv
}
func (g GTMioShaderProfilerGPUCommand) EncoderObjectId() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("encoderObjectId"))
	return rv
}
func (g GTMioShaderProfilerGPUCommand) FunctionIndex() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("functionIndex"))
	return rv
}
func (g GTMioShaderProfilerGPUCommand) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}
func (g GTMioShaderProfilerGPUCommand) Index() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("index"))
	return rv
}
func (g GTMioShaderProfilerGPUCommand) PipelineInfoIndex() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("pipelineInfoIndex"))
	return rv
}
func (g GTMioShaderProfilerGPUCommand) PipelineStateObjectId() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("pipelineStateObjectId"))
	return rv
}
func (g GTMioShaderProfilerGPUCommand) SubCommandIndex() int {
	rv := objc.Send[int](g.ID, objc.Sel("subCommandIndex"))
	return rv
}
func (g GTMioShaderProfilerGPUCommand) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](g.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (g GTMioShaderProfilerGPUCommand) TimingInfo() IGTShaderProfilerTimingInfo {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("timingInfo"))
	return GTShaderProfilerTimingInfoFromID(objc.ID(rv))
}

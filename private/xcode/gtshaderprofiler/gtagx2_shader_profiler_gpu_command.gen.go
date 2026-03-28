// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTAGX2ShaderProfilerGPUCommand] class.
var (
	_GTAGX2ShaderProfilerGPUCommandClass     GTAGX2ShaderProfilerGPUCommandClass
	_GTAGX2ShaderProfilerGPUCommandClassOnce sync.Once
)

func getGTAGX2ShaderProfilerGPUCommandClass() GTAGX2ShaderProfilerGPUCommandClass {
	_GTAGX2ShaderProfilerGPUCommandClassOnce.Do(func() {
		_GTAGX2ShaderProfilerGPUCommandClass = GTAGX2ShaderProfilerGPUCommandClass{class: objc.GetClass("GTAGX2ShaderProfilerGPUCommand")}
	})
	return _GTAGX2ShaderProfilerGPUCommandClass
}

// GetGTAGX2ShaderProfilerGPUCommandClass returns the class object for GTAGX2ShaderProfilerGPUCommand.
func GetGTAGX2ShaderProfilerGPUCommandClass() GTAGX2ShaderProfilerGPUCommandClass {
	return getGTAGX2ShaderProfilerGPUCommandClass()
}

type GTAGX2ShaderProfilerGPUCommandClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTAGX2ShaderProfilerGPUCommandClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTAGX2ShaderProfilerGPUCommandClass) Alloc() GTAGX2ShaderProfilerGPUCommand {
	rv := objc.Send[GTAGX2ShaderProfilerGPUCommand](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTAGX2ShaderProfilerGPUCommand.AddBinaryKeyForType]
//   - [GTAGX2ShaderProfilerGPUCommand.AllBinaryKeys]
//   - [GTAGX2ShaderProfilerGPUCommand.BinaryKeys]
//   - [GTAGX2ShaderProfilerGPUCommand.CommandBufferIndex]
//   - [GTAGX2ShaderProfilerGPUCommand.SetCommandBufferIndex]
//   - [GTAGX2ShaderProfilerGPUCommand.CommandType]
//   - [GTAGX2ShaderProfilerGPUCommand.SetCommandType]
//   - [GTAGX2ShaderProfilerGPUCommand.EncodeWithCoder]
//   - [GTAGX2ShaderProfilerGPUCommand.EncoderInfoIndex]
//   - [GTAGX2ShaderProfilerGPUCommand.SetEncoderInfoIndex]
//   - [GTAGX2ShaderProfilerGPUCommand.EncoderObjectId]
//   - [GTAGX2ShaderProfilerGPUCommand.SetEncoderObjectId]
//   - [GTAGX2ShaderProfilerGPUCommand.FunctionIndex]
//   - [GTAGX2ShaderProfilerGPUCommand.SetFunctionIndex]
//   - [GTAGX2ShaderProfilerGPUCommand.MioData]
//   - [GTAGX2ShaderProfilerGPUCommand.PipelineInfoIndex]
//   - [GTAGX2ShaderProfilerGPUCommand.SetPipelineInfoIndex]
//   - [GTAGX2ShaderProfilerGPUCommand.PipelineStateObjectId]
//   - [GTAGX2ShaderProfilerGPUCommand.SetPipelineStateObjectId]
//   - [GTAGX2ShaderProfilerGPUCommand.SetIndex]
//   - [GTAGX2ShaderProfilerGPUCommand.SubCommandIndex]
//   - [GTAGX2ShaderProfilerGPUCommand.SetSubCommandIndex]
//   - [GTAGX2ShaderProfilerGPUCommand.TimingInfo]
//   - [GTAGX2ShaderProfilerGPUCommand.SetTimingInfo]
//   - [GTAGX2ShaderProfilerGPUCommand.InitWithCoder]
//   - [GTAGX2ShaderProfilerGPUCommand.DebugDescription]
//   - [GTAGX2ShaderProfilerGPUCommand.Description]
//   - [GTAGX2ShaderProfilerGPUCommand.Hash]
//   - [GTAGX2ShaderProfilerGPUCommand.Superclass]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand
type GTAGX2ShaderProfilerGPUCommand struct {
	objectivec.Object
}

// GTAGX2ShaderProfilerGPUCommandFromID constructs a [GTAGX2ShaderProfilerGPUCommand] from an objc.ID.
func GTAGX2ShaderProfilerGPUCommandFromID(id objc.ID) GTAGX2ShaderProfilerGPUCommand {
	return GTAGX2ShaderProfilerGPUCommand{objectivec.Object{ID: id}}
}
// Ensure GTAGX2ShaderProfilerGPUCommand implements IGTAGX2ShaderProfilerGPUCommand.
var _ IGTAGX2ShaderProfilerGPUCommand = GTAGX2ShaderProfilerGPUCommand{}

// An interface definition for the [GTAGX2ShaderProfilerGPUCommand] class.
//
// # Methods
//
//   - [IGTAGX2ShaderProfilerGPUCommand.AddBinaryKeyForType]
//   - [IGTAGX2ShaderProfilerGPUCommand.AllBinaryKeys]
//   - [IGTAGX2ShaderProfilerGPUCommand.BinaryKeys]
//   - [IGTAGX2ShaderProfilerGPUCommand.CommandBufferIndex]
//   - [IGTAGX2ShaderProfilerGPUCommand.SetCommandBufferIndex]
//   - [IGTAGX2ShaderProfilerGPUCommand.CommandType]
//   - [IGTAGX2ShaderProfilerGPUCommand.SetCommandType]
//   - [IGTAGX2ShaderProfilerGPUCommand.EncodeWithCoder]
//   - [IGTAGX2ShaderProfilerGPUCommand.EncoderInfoIndex]
//   - [IGTAGX2ShaderProfilerGPUCommand.SetEncoderInfoIndex]
//   - [IGTAGX2ShaderProfilerGPUCommand.EncoderObjectId]
//   - [IGTAGX2ShaderProfilerGPUCommand.SetEncoderObjectId]
//   - [IGTAGX2ShaderProfilerGPUCommand.FunctionIndex]
//   - [IGTAGX2ShaderProfilerGPUCommand.SetFunctionIndex]
//   - [IGTAGX2ShaderProfilerGPUCommand.MioData]
//   - [IGTAGX2ShaderProfilerGPUCommand.PipelineInfoIndex]
//   - [IGTAGX2ShaderProfilerGPUCommand.SetPipelineInfoIndex]
//   - [IGTAGX2ShaderProfilerGPUCommand.PipelineStateObjectId]
//   - [IGTAGX2ShaderProfilerGPUCommand.SetPipelineStateObjectId]
//   - [IGTAGX2ShaderProfilerGPUCommand.SetIndex]
//   - [IGTAGX2ShaderProfilerGPUCommand.SubCommandIndex]
//   - [IGTAGX2ShaderProfilerGPUCommand.SetSubCommandIndex]
//   - [IGTAGX2ShaderProfilerGPUCommand.TimingInfo]
//   - [IGTAGX2ShaderProfilerGPUCommand.SetTimingInfo]
//   - [IGTAGX2ShaderProfilerGPUCommand.InitWithCoder]
//   - [IGTAGX2ShaderProfilerGPUCommand.DebugDescription]
//   - [IGTAGX2ShaderProfilerGPUCommand.Description]
//   - [IGTAGX2ShaderProfilerGPUCommand.Hash]
//   - [IGTAGX2ShaderProfilerGPUCommand.Superclass]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand
type IGTAGX2ShaderProfilerGPUCommand interface {
	objectivec.IObject

	// Topic: Methods

	AddBinaryKeyForType(key objectivec.IObject, type_ uint32)
	AllBinaryKeys() foundation.INSDictionary
	BinaryKeys() foundation.INSDictionary
	CommandBufferIndex() uint32
	SetCommandBufferIndex(value uint32)
	CommandType() uint32
	SetCommandType(value uint32)
	EncodeWithCoder(coder foundation.INSCoder)
	EncoderInfoIndex() uint32
	SetEncoderInfoIndex(value uint32)
	EncoderObjectId() uint64
	SetEncoderObjectId(value uint64)
	FunctionIndex() uint64
	SetFunctionIndex(value uint64)
	MioData() objectivec.IObject
	PipelineInfoIndex() uint32
	SetPipelineInfoIndex(value uint32)
	PipelineStateObjectId() uint64
	SetPipelineStateObjectId(value uint64)
	SetIndex(index uint32)
	SubCommandIndex() int
	SetSubCommandIndex(value int)
	TimingInfo() IGTShaderProfilerTimingInfo
	SetTimingInfo(value IGTShaderProfilerTimingInfo)
	InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderProfilerGPUCommand
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (g GTAGX2ShaderProfilerGPUCommand) Init() GTAGX2ShaderProfilerGPUCommand {
	rv := objc.Send[GTAGX2ShaderProfilerGPUCommand](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTAGX2ShaderProfilerGPUCommand) Autorelease() GTAGX2ShaderProfilerGPUCommand {
	rv := objc.Send[GTAGX2ShaderProfilerGPUCommand](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTAGX2ShaderProfilerGPUCommand creates a new GTAGX2ShaderProfilerGPUCommand instance.
func NewGTAGX2ShaderProfilerGPUCommand() GTAGX2ShaderProfilerGPUCommand {
	class := getGTAGX2ShaderProfilerGPUCommandClass()
	rv := objc.Send[GTAGX2ShaderProfilerGPUCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/initWithCoder:
func NewGTAGX2ShaderProfilerGPUCommandWithCoder(coder objectivec.IObject) GTAGX2ShaderProfilerGPUCommand {
	instance := getGTAGX2ShaderProfilerGPUCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTAGX2ShaderProfilerGPUCommandFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/addBinaryKey:forType:
func (g GTAGX2ShaderProfilerGPUCommand) AddBinaryKeyForType(key objectivec.IObject, type_ uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("addBinaryKey:forType:"), key, type_)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/encodeWithCoder:
func (g GTAGX2ShaderProfilerGPUCommand) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/mioData
func (g GTAGX2ShaderProfilerGPUCommand) MioData() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("mioData"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/setIndex:
func (g GTAGX2ShaderProfilerGPUCommand) SetIndex(index uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("setIndex:"), index)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/initWithCoder:
func (g GTAGX2ShaderProfilerGPUCommand) InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderProfilerGPUCommand {
	rv := objc.Send[GTAGX2ShaderProfilerGPUCommand](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/supportsSecureCoding
func (_GTAGX2ShaderProfilerGPUCommandClass GTAGX2ShaderProfilerGPUCommandClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_GTAGX2ShaderProfilerGPUCommandClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/allBinaryKeys
func (g GTAGX2ShaderProfilerGPUCommand) AllBinaryKeys() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("allBinaryKeys"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/binaryKeys
func (g GTAGX2ShaderProfilerGPUCommand) BinaryKeys() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("binaryKeys"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/commandBufferIndex
func (g GTAGX2ShaderProfilerGPUCommand) CommandBufferIndex() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("commandBufferIndex"))
	return rv
}
func (g GTAGX2ShaderProfilerGPUCommand) SetCommandBufferIndex(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setCommandBufferIndex:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/commandType
func (g GTAGX2ShaderProfilerGPUCommand) CommandType() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("commandType"))
	return rv
}
func (g GTAGX2ShaderProfilerGPUCommand) SetCommandType(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setCommandType:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/debugDescription
func (g GTAGX2ShaderProfilerGPUCommand) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/description
func (g GTAGX2ShaderProfilerGPUCommand) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/encoderInfoIndex
func (g GTAGX2ShaderProfilerGPUCommand) EncoderInfoIndex() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("encoderInfoIndex"))
	return rv
}
func (g GTAGX2ShaderProfilerGPUCommand) SetEncoderInfoIndex(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setEncoderInfoIndex:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/encoderObjectId
func (g GTAGX2ShaderProfilerGPUCommand) EncoderObjectId() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("encoderObjectId"))
	return rv
}
func (g GTAGX2ShaderProfilerGPUCommand) SetEncoderObjectId(value uint64) {
	objc.Send[struct{}](g.ID, objc.Sel("setEncoderObjectId:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/functionIndex
func (g GTAGX2ShaderProfilerGPUCommand) FunctionIndex() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("functionIndex"))
	return rv
}
func (g GTAGX2ShaderProfilerGPUCommand) SetFunctionIndex(value uint64) {
	objc.Send[struct{}](g.ID, objc.Sel("setFunctionIndex:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/hash
func (g GTAGX2ShaderProfilerGPUCommand) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/pipelineInfoIndex
func (g GTAGX2ShaderProfilerGPUCommand) PipelineInfoIndex() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("pipelineInfoIndex"))
	return rv
}
func (g GTAGX2ShaderProfilerGPUCommand) SetPipelineInfoIndex(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setPipelineInfoIndex:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/pipelineStateObjectId
func (g GTAGX2ShaderProfilerGPUCommand) PipelineStateObjectId() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("pipelineStateObjectId"))
	return rv
}
func (g GTAGX2ShaderProfilerGPUCommand) SetPipelineStateObjectId(value uint64) {
	objc.Send[struct{}](g.ID, objc.Sel("setPipelineStateObjectId:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/subCommandIndex
func (g GTAGX2ShaderProfilerGPUCommand) SubCommandIndex() int {
	rv := objc.Send[int](g.ID, objc.Sel("subCommandIndex"))
	return rv
}
func (g GTAGX2ShaderProfilerGPUCommand) SetSubCommandIndex(value int) {
	objc.Send[struct{}](g.ID, objc.Sel("setSubCommandIndex:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/superclass
func (g GTAGX2ShaderProfilerGPUCommand) Superclass() objc.Class {
	rv := objc.Send[objc.Class](g.ID, objc.Sel("superclass"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerGPUCommand/timingInfo
func (g GTAGX2ShaderProfilerGPUCommand) TimingInfo() IGTShaderProfilerTimingInfo {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("timingInfo"))
	return GTShaderProfilerTimingInfoFromID(objc.ID(rv))
}
func (g GTAGX2ShaderProfilerGPUCommand) SetTimingInfo(value IGTShaderProfilerTimingInfo) {
	objc.Send[struct{}](g.ID, objc.Sel("setTimingInfo:"), value)
}


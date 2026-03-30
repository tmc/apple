// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioShaderAnalyzer] class.
var (
	_GTMioShaderAnalyzerClass     GTMioShaderAnalyzerClass
	_GTMioShaderAnalyzerClassOnce sync.Once
)

func getGTMioShaderAnalyzerClass() GTMioShaderAnalyzerClass {
	_GTMioShaderAnalyzerClassOnce.Do(func() {
		_GTMioShaderAnalyzerClass = GTMioShaderAnalyzerClass{class: objc.GetClass("GTMioShaderAnalyzer")}
	})
	return _GTMioShaderAnalyzerClass
}

// GetGTMioShaderAnalyzerClass returns the class object for GTMioShaderAnalyzer.
func GetGTMioShaderAnalyzerClass() GTMioShaderAnalyzerClass {
	return getGTMioShaderAnalyzerClass()
}

type GTMioShaderAnalyzerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioShaderAnalyzerClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioShaderAnalyzerClass) Alloc() GTMioShaderAnalyzer {
	rv := objc.Send[GTMioShaderAnalyzer](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioShaderAnalyzer.BuildDrawProgramTypeTraceData]
//   - [GTMioShaderAnalyzer.BuildEncoderProgramTypeTraceData]
//   - [GTMioShaderAnalyzer.BuildPipelineProgramTypeTraceData]
//   - [GTMioShaderAnalyzer.FinalizeBuild]
//   - [GTMioShaderAnalyzer.HandleBinary]
//   - [GTMioShaderAnalyzer.InstructionDataTypeInfo]
//   - [GTMioShaderAnalyzer.InstructionMemoryTypeInfo]
//   - [GTMioShaderAnalyzer.InstructionScopeInfo]
//   - [GTMioShaderAnalyzer.InstructionTypeInfo]
//   - [GTMioShaderAnalyzer.NumInstructionDataTypeInfo]
//   - [GTMioShaderAnalyzer.NumInstructionMemoryTypeInfo]
//   - [GTMioShaderAnalyzer.NumInstructionScopeInfo]
//   - [GTMioShaderAnalyzer.NumInstructionTypeInfo]
//   - [GTMioShaderAnalyzer.ProgramType]
//   - [GTMioShaderAnalyzer.Scope]
//   - [GTMioShaderAnalyzer.ScopeIdentifier]
//   - [GTMioShaderAnalyzer.TraceData]
//   - [GTMioShaderAnalyzer.UseBaseProgramType]
//   - [GTMioShaderAnalyzer.InitWithDrawIndexProgramTypeUseBaseProgramTypeTraceData]
//   - [GTMioShaderAnalyzer.InitWithEncoderFunctionIndexProgramTypeUseBaseProgramTypeTraceData]
//   - [GTMioShaderAnalyzer.InitWithPipelineStateIdProgramTypeUseBaseProgramTypeTraceData]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer
type GTMioShaderAnalyzer struct {
	objectivec.Object
}

// GTMioShaderAnalyzerFromID constructs a [GTMioShaderAnalyzer] from an objc.ID.
func GTMioShaderAnalyzerFromID(id objc.ID) GTMioShaderAnalyzer {
	return GTMioShaderAnalyzer{objectivec.Object{ID: id}}
}

// Ensure GTMioShaderAnalyzer implements IGTMioShaderAnalyzer.
var _ IGTMioShaderAnalyzer = GTMioShaderAnalyzer{}

// An interface definition for the [GTMioShaderAnalyzer] class.
//
// # Methods
//
//   - [IGTMioShaderAnalyzer.BuildDrawProgramTypeTraceData]
//   - [IGTMioShaderAnalyzer.BuildEncoderProgramTypeTraceData]
//   - [IGTMioShaderAnalyzer.BuildPipelineProgramTypeTraceData]
//   - [IGTMioShaderAnalyzer.FinalizeBuild]
//   - [IGTMioShaderAnalyzer.HandleBinary]
//   - [IGTMioShaderAnalyzer.InstructionDataTypeInfo]
//   - [IGTMioShaderAnalyzer.InstructionMemoryTypeInfo]
//   - [IGTMioShaderAnalyzer.InstructionScopeInfo]
//   - [IGTMioShaderAnalyzer.InstructionTypeInfo]
//   - [IGTMioShaderAnalyzer.NumInstructionDataTypeInfo]
//   - [IGTMioShaderAnalyzer.NumInstructionMemoryTypeInfo]
//   - [IGTMioShaderAnalyzer.NumInstructionScopeInfo]
//   - [IGTMioShaderAnalyzer.NumInstructionTypeInfo]
//   - [IGTMioShaderAnalyzer.ProgramType]
//   - [IGTMioShaderAnalyzer.Scope]
//   - [IGTMioShaderAnalyzer.ScopeIdentifier]
//   - [IGTMioShaderAnalyzer.TraceData]
//   - [IGTMioShaderAnalyzer.UseBaseProgramType]
//   - [IGTMioShaderAnalyzer.InitWithDrawIndexProgramTypeUseBaseProgramTypeTraceData]
//   - [IGTMioShaderAnalyzer.InitWithEncoderFunctionIndexProgramTypeUseBaseProgramTypeTraceData]
//   - [IGTMioShaderAnalyzer.InitWithPipelineStateIdProgramTypeUseBaseProgramTypeTraceData]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer
type IGTMioShaderAnalyzer interface {
	objectivec.IObject

	// Topic: Methods

	BuildDrawProgramTypeTraceData(draw uint32, type_ uint16, data objectivec.IObject) bool
	BuildEncoderProgramTypeTraceData(encoder uint32, type_ uint16, data objectivec.IObject) bool
	BuildPipelineProgramTypeTraceData(pipeline uint64, type_ uint16, data objectivec.IObject) bool
	FinalizeBuild()
	HandleBinary(binary objectivec.IObject)
	InstructionDataTypeInfo() objectivec.IObject
	InstructionMemoryTypeInfo() objectivec.IObject
	InstructionScopeInfo() objectivec.IObject
	InstructionTypeInfo() objectivec.IObject
	NumInstructionDataTypeInfo() int64
	NumInstructionMemoryTypeInfo() int64
	NumInstructionScopeInfo() int64
	NumInstructionTypeInfo() int64
	ProgramType() uint16
	Scope() uint16
	ScopeIdentifier() uint64
	TraceData() objectivec.IObject
	UseBaseProgramType() bool
	InitWithDrawIndexProgramTypeUseBaseProgramTypeTraceData(index uint32, type_ uint16, type_2 bool, data objectivec.IObject) GTMioShaderAnalyzer
	InitWithEncoderFunctionIndexProgramTypeUseBaseProgramTypeTraceData(index uint32, type_ uint16, type_2 bool, data objectivec.IObject) GTMioShaderAnalyzer
	InitWithPipelineStateIdProgramTypeUseBaseProgramTypeTraceData(id uint64, type_ uint16, type_2 bool, data objectivec.IObject) GTMioShaderAnalyzer
}

// Init initializes the instance.
func (g GTMioShaderAnalyzer) Init() GTMioShaderAnalyzer {
	rv := objc.Send[GTMioShaderAnalyzer](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioShaderAnalyzer) Autorelease() GTMioShaderAnalyzer {
	rv := objc.Send[GTMioShaderAnalyzer](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioShaderAnalyzer creates a new GTMioShaderAnalyzer instance.
func NewGTMioShaderAnalyzer() GTMioShaderAnalyzer {
	class := getGTMioShaderAnalyzerClass()
	rv := objc.Send[GTMioShaderAnalyzer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/initWithDrawIndex:programType:useBaseProgramType:traceData:
func NewGTMioShaderAnalyzerWithDrawIndexProgramTypeUseBaseProgramTypeTraceData(index uint32, type_ uint16, type_2 bool, data objectivec.IObject) GTMioShaderAnalyzer {
	instance := getGTMioShaderAnalyzerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDrawIndex:programType:useBaseProgramType:traceData:"), index, type_, type_2, data)
	return GTMioShaderAnalyzerFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/initWithEncoderFunctionIndex:programType:useBaseProgramType:traceData:
func NewGTMioShaderAnalyzerWithEncoderFunctionIndexProgramTypeUseBaseProgramTypeTraceData(index uint32, type_ uint16, type_2 bool, data objectivec.IObject) GTMioShaderAnalyzer {
	instance := getGTMioShaderAnalyzerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEncoderFunctionIndex:programType:useBaseProgramType:traceData:"), index, type_, type_2, data)
	return GTMioShaderAnalyzerFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/initWithPipelineStateId:programType:useBaseProgramType:traceData:
func NewGTMioShaderAnalyzerWithPipelineStateIdProgramTypeUseBaseProgramTypeTraceData(id uint64, type_ uint16, type_2 bool, data objectivec.IObject) GTMioShaderAnalyzer {
	instance := getGTMioShaderAnalyzerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPipelineStateId:programType:useBaseProgramType:traceData:"), id, type_, type_2, data)
	return GTMioShaderAnalyzerFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/buildDraw:programType:traceData:
func (g GTMioShaderAnalyzer) BuildDrawProgramTypeTraceData(draw uint32, type_ uint16, data objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("buildDraw:programType:traceData:"), draw, type_, data)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/buildEncoder:programType:traceData:
func (g GTMioShaderAnalyzer) BuildEncoderProgramTypeTraceData(encoder uint32, type_ uint16, data objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("buildEncoder:programType:traceData:"), encoder, type_, data)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/buildPipeline:programType:traceData:
func (g GTMioShaderAnalyzer) BuildPipelineProgramTypeTraceData(pipeline uint64, type_ uint16, data objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("buildPipeline:programType:traceData:"), pipeline, type_, data)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/finalizeBuild
func (g GTMioShaderAnalyzer) FinalizeBuild() {
	objc.Send[objc.ID](g.ID, objc.Sel("finalizeBuild"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/handleBinary:
func (g GTMioShaderAnalyzer) HandleBinary(binary objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("handleBinary:"), binary)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/initWithDrawIndex:programType:useBaseProgramType:traceData:
func (g GTMioShaderAnalyzer) InitWithDrawIndexProgramTypeUseBaseProgramTypeTraceData(index uint32, type_ uint16, type_2 bool, data objectivec.IObject) GTMioShaderAnalyzer {
	rv := objc.Send[GTMioShaderAnalyzer](g.ID, objc.Sel("initWithDrawIndex:programType:useBaseProgramType:traceData:"), index, type_, type_2, data)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/initWithEncoderFunctionIndex:programType:useBaseProgramType:traceData:
func (g GTMioShaderAnalyzer) InitWithEncoderFunctionIndexProgramTypeUseBaseProgramTypeTraceData(index uint32, type_ uint16, type_2 bool, data objectivec.IObject) GTMioShaderAnalyzer {
	rv := objc.Send[GTMioShaderAnalyzer](g.ID, objc.Sel("initWithEncoderFunctionIndex:programType:useBaseProgramType:traceData:"), index, type_, type_2, data)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/initWithPipelineStateId:programType:useBaseProgramType:traceData:
func (g GTMioShaderAnalyzer) InitWithPipelineStateIdProgramTypeUseBaseProgramTypeTraceData(id uint64, type_ uint16, type_2 bool, data objectivec.IObject) GTMioShaderAnalyzer {
	rv := objc.Send[GTMioShaderAnalyzer](g.ID, objc.Sel("initWithPipelineStateId:programType:useBaseProgramType:traceData:"), id, type_, type_2, data)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/instructionDataTypeInfo
func (g GTMioShaderAnalyzer) InstructionDataTypeInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("instructionDataTypeInfo"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/instructionMemoryTypeInfo
func (g GTMioShaderAnalyzer) InstructionMemoryTypeInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("instructionMemoryTypeInfo"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/instructionScopeInfo
func (g GTMioShaderAnalyzer) InstructionScopeInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("instructionScopeInfo"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/instructionTypeInfo
func (g GTMioShaderAnalyzer) InstructionTypeInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("instructionTypeInfo"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/numInstructionDataTypeInfo
func (g GTMioShaderAnalyzer) NumInstructionDataTypeInfo() int64 {
	rv := objc.Send[int64](g.ID, objc.Sel("numInstructionDataTypeInfo"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/numInstructionMemoryTypeInfo
func (g GTMioShaderAnalyzer) NumInstructionMemoryTypeInfo() int64 {
	rv := objc.Send[int64](g.ID, objc.Sel("numInstructionMemoryTypeInfo"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/numInstructionScopeInfo
func (g GTMioShaderAnalyzer) NumInstructionScopeInfo() int64 {
	rv := objc.Send[int64](g.ID, objc.Sel("numInstructionScopeInfo"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/numInstructionTypeInfo
func (g GTMioShaderAnalyzer) NumInstructionTypeInfo() int64 {
	rv := objc.Send[int64](g.ID, objc.Sel("numInstructionTypeInfo"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/programType
func (g GTMioShaderAnalyzer) ProgramType() uint16 {
	rv := objc.Send[uint16](g.ID, objc.Sel("programType"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/scope
func (g GTMioShaderAnalyzer) Scope() uint16 {
	rv := objc.Send[uint16](g.ID, objc.Sel("scope"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/scopeIdentifier
func (g GTMioShaderAnalyzer) ScopeIdentifier() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("scopeIdentifier"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/traceData
func (g GTMioShaderAnalyzer) TraceData() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("traceData"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderAnalyzer/useBaseProgramType
func (g GTMioShaderAnalyzer) UseBaseProgramType() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("useBaseProgramType"))
	return rv
}

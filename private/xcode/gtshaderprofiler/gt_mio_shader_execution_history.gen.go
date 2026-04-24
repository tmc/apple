// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioShaderExecutionHistory] class.
var (
	_GTMioShaderExecutionHistoryClass     GTMioShaderExecutionHistoryClass
	_GTMioShaderExecutionHistoryClassOnce sync.Once
)

func getGTMioShaderExecutionHistoryClass() GTMioShaderExecutionHistoryClass {
	_GTMioShaderExecutionHistoryClassOnce.Do(func() {
		_GTMioShaderExecutionHistoryClass = GTMioShaderExecutionHistoryClass{class: objc.GetClass("GTMioShaderExecutionHistory")}
	})
	return _GTMioShaderExecutionHistoryClass
}

// GetGTMioShaderExecutionHistoryClass returns the class object for GTMioShaderExecutionHistory.
func GetGTMioShaderExecutionHistoryClass() GTMioShaderExecutionHistoryClass {
	return getGTMioShaderExecutionHistoryClass()
}

type GTMioShaderExecutionHistoryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioShaderExecutionHistoryClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioShaderExecutionHistoryClass) Alloc() GTMioShaderExecutionHistory {
	rv := objc.Send[GTMioShaderExecutionHistory](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioShaderExecutionHistory.AllocStyles]
//   - [GTMioShaderExecutionHistory.CacheKey]
//   - [GTMioShaderExecutionHistory.CacheObject]
//   - [GTMioShaderExecutionHistory.CallStack]
//   - [GTMioShaderExecutionHistory.CliqueExecutionHistoryBeginUsc]
//   - [GTMioShaderExecutionHistory.CliqueExecutionHistoryEndUsc]
//   - [GTMioShaderExecutionHistory.CliqueExecutionHistoryStyle]
//   - [GTMioShaderExecutionHistory.Compact]
//   - [GTMioShaderExecutionHistory.Delegate]
//   - [GTMioShaderExecutionHistory.SetDelegate]
//   - [GTMioShaderExecutionHistory.DumpTree]
//   - [GTMioShaderExecutionHistory.Full]
//   - [GTMioShaderExecutionHistory.FullInstructionHistory]
//   - [GTMioShaderExecutionHistory.GenerateClique]
//   - [GTMioShaderExecutionHistory.GenerateCliqueIndexUscIndex]
//   - [GTMioShaderExecutionHistory.GenerateDrawIndexProgramType]
//   - [GTMioShaderExecutionHistory.GenerateDrawIndexProgramTypeProgressController]
//   - [GTMioShaderExecutionHistory.GeneratePipelineStateIdProgramType]
//   - [GTMioShaderExecutionHistory.GeneratePipelineStateIdProgramTypeProgressController]
//   - [GTMioShaderExecutionHistory.HandleCachedObject]
//   - [GTMioShaderExecutionHistory.LoopBackInstructionEndLoopCountCurrentLoopCountBinary]
//   - [GTMioShaderExecutionHistory.NodeForStyle]
//   - [GTMioShaderExecutionHistory.Options]
//   - [GTMioShaderExecutionHistory.PopFunctionBinaryRangeBinary]
//   - [GTMioShaderExecutionHistory.PopLoopInstructionEndLoopCountBinary]
//   - [GTMioShaderExecutionHistory.ProcessInstructionBinaryRangeBinaryNumHits]
//   - [GTMioShaderExecutionHistory.PushFunctionBinaryRangeIndexInlinedBinaryCallerLocationCallerBinaryRangeIndexCallerBinary]
//   - [GTMioShaderExecutionHistory.PushLoopInstructionEndLoopCountBinary]
//   - [GTMioShaderExecutionHistory.Style]
//   - [GTMioShaderExecutionHistory.TimestampNextInstructionCount]
//   - [GTMioShaderExecutionHistory.TraceData]
//   - [GTMioShaderExecutionHistory.InitWithTraceDataStyleOptionsDelegate]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory
type GTMioShaderExecutionHistory struct {
	objectivec.Object
}

// GTMioShaderExecutionHistoryFromID constructs a [GTMioShaderExecutionHistory] from an objc.ID.
func GTMioShaderExecutionHistoryFromID(id objc.ID) GTMioShaderExecutionHistory {
	return GTMioShaderExecutionHistory{objectivec.Object{ID: id}}
}

// Ensure GTMioShaderExecutionHistory implements IGTMioShaderExecutionHistory.
var _ IGTMioShaderExecutionHistory = GTMioShaderExecutionHistory{}

// An interface definition for the [GTMioShaderExecutionHistory] class.
//
// # Methods
//
//   - [IGTMioShaderExecutionHistory.AllocStyles]
//   - [IGTMioShaderExecutionHistory.CacheKey]
//   - [IGTMioShaderExecutionHistory.CacheObject]
//   - [IGTMioShaderExecutionHistory.CallStack]
//   - [IGTMioShaderExecutionHistory.CliqueExecutionHistoryBeginUsc]
//   - [IGTMioShaderExecutionHistory.CliqueExecutionHistoryEndUsc]
//   - [IGTMioShaderExecutionHistory.CliqueExecutionHistoryStyle]
//   - [IGTMioShaderExecutionHistory.Compact]
//   - [IGTMioShaderExecutionHistory.Delegate]
//   - [IGTMioShaderExecutionHistory.SetDelegate]
//   - [IGTMioShaderExecutionHistory.DumpTree]
//   - [IGTMioShaderExecutionHistory.Full]
//   - [IGTMioShaderExecutionHistory.FullInstructionHistory]
//   - [IGTMioShaderExecutionHistory.GenerateClique]
//   - [IGTMioShaderExecutionHistory.GenerateCliqueIndexUscIndex]
//   - [IGTMioShaderExecutionHistory.GenerateDrawIndexProgramType]
//   - [IGTMioShaderExecutionHistory.GenerateDrawIndexProgramTypeProgressController]
//   - [IGTMioShaderExecutionHistory.GeneratePipelineStateIdProgramType]
//   - [IGTMioShaderExecutionHistory.GeneratePipelineStateIdProgramTypeProgressController]
//   - [IGTMioShaderExecutionHistory.HandleCachedObject]
//   - [IGTMioShaderExecutionHistory.LoopBackInstructionEndLoopCountCurrentLoopCountBinary]
//   - [IGTMioShaderExecutionHistory.NodeForStyle]
//   - [IGTMioShaderExecutionHistory.Options]
//   - [IGTMioShaderExecutionHistory.PopFunctionBinaryRangeBinary]
//   - [IGTMioShaderExecutionHistory.PopLoopInstructionEndLoopCountBinary]
//   - [IGTMioShaderExecutionHistory.ProcessInstructionBinaryRangeBinaryNumHits]
//   - [IGTMioShaderExecutionHistory.PushFunctionBinaryRangeIndexInlinedBinaryCallerLocationCallerBinaryRangeIndexCallerBinary]
//   - [IGTMioShaderExecutionHistory.PushLoopInstructionEndLoopCountBinary]
//   - [IGTMioShaderExecutionHistory.Style]
//   - [IGTMioShaderExecutionHistory.TimestampNextInstructionCount]
//   - [IGTMioShaderExecutionHistory.TraceData]
//   - [IGTMioShaderExecutionHistory.InitWithTraceDataStyleOptionsDelegate]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory
type IGTMioShaderExecutionHistory interface {
	objectivec.IObject

	// Topic: Methods

	AllocStyles()
	CacheKey() objectivec.IObject
	CacheObject() objectivec.IObject
	CallStack() IGTMioShaderExecutionHistoryRootNode
	CliqueExecutionHistoryBeginUsc(begin *GTMioUSCCliqueMetadataRef, usc objectivec.IObject)
	CliqueExecutionHistoryEndUsc(end *GTMioUSCCliqueMetadataRef, usc objectivec.IObject)
	CliqueExecutionHistoryStyle() uint32
	Compact() IGTMioShaderExecutionHistoryRootNode
	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)
	DumpTree(tree objectivec.IObject)
	Full() IGTMioShaderExecutionHistoryRootNode
	FullInstructionHistory() bool
	GenerateClique(clique *GTMioUSCCliqueMetadataRef) bool
	GenerateCliqueIndexUscIndex(index uint32, index2 uint32) bool
	GenerateDrawIndexProgramType(index uint32, type_ uint16) bool
	GenerateDrawIndexProgramTypeProgressController(index uint32, type_ uint16, controller objectivec.IObject) bool
	GeneratePipelineStateIdProgramType(id uint64, type_ uint16) bool
	GeneratePipelineStateIdProgramTypeProgressController(id uint64, type_ uint16, controller objectivec.IObject) bool
	HandleCachedObject(object objectivec.IObject) bool
	LoopBackInstructionEndLoopCountCurrentLoopCountBinary(back uint32, end uint32, count uint32, count2 uint32, binary objectivec.IObject)
	NodeForStyle(style uint32) objectivec.IObject
	Options() uint32
	PopFunctionBinaryRangeBinary(function unsafe.Pointer, range_ unsafe.Pointer, binary objectivec.IObject)
	PopLoopInstructionEndLoopCountBinary(loop uint32, end uint32, count uint32, binary objectivec.IObject)
	ProcessInstructionBinaryRangeBinaryNumHits(instruction uint32, range_ unsafe.Pointer, binary objectivec.IObject, hits uint32)
	PushFunctionBinaryRangeIndexInlinedBinaryCallerLocationCallerBinaryRangeIndexCallerBinary(function unsafe.Pointer, index uint32, inlined bool, binary objectivec.IObject, location unsafe.Pointer, index2 uint32, binary2 objectivec.IObject)
	PushLoopInstructionEndLoopCountBinary(loop uint32, end uint32, count uint32, binary objectivec.IObject)
	Style() uint32
	TimestampNextInstructionCount(timestamp uint64, next uint64, count uint32)
	TraceData() objectivec.IObject
	InitWithTraceDataStyleOptionsDelegate(data objectivec.IObject, style uint32, options uint32, delegate objectivec.IObject) GTMioShaderExecutionHistory
}

// Init initializes the instance.
func (g GTMioShaderExecutionHistory) Init() GTMioShaderExecutionHistory {
	rv := objc.Send[GTMioShaderExecutionHistory](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioShaderExecutionHistory) Autorelease() GTMioShaderExecutionHistory {
	rv := objc.Send[GTMioShaderExecutionHistory](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioShaderExecutionHistory creates a new GTMioShaderExecutionHistory instance.
func NewGTMioShaderExecutionHistory() GTMioShaderExecutionHistory {
	class := getGTMioShaderExecutionHistoryClass()
	rv := objc.Send[GTMioShaderExecutionHistory](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/initWithTraceData:style:options:delegate:
func NewGTMioShaderExecutionHistoryWithTraceDataStyleOptionsDelegate(data objectivec.IObject, style uint32, options uint32, delegate objectivec.IObject) GTMioShaderExecutionHistory {
	instance := getGTMioShaderExecutionHistoryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTraceData:style:options:delegate:"), data, style, options, delegate)
	return GTMioShaderExecutionHistoryFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/allocStyles
func (g GTMioShaderExecutionHistory) AllocStyles() {
	objc.Send[objc.ID](g.ID, objc.Sel("allocStyles"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/cacheKey
func (g GTMioShaderExecutionHistory) CacheKey() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cacheKey"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/cacheObject
func (g GTMioShaderExecutionHistory) CacheObject() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cacheObject"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/cliqueExecutionHistoryBegin:usc:
func (g GTMioShaderExecutionHistory) CliqueExecutionHistoryBeginUsc(begin *GTMioUSCCliqueMetadataRef, usc objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("cliqueExecutionHistoryBegin:usc:"), begin, usc)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/cliqueExecutionHistoryEnd:usc:
func (g GTMioShaderExecutionHistory) CliqueExecutionHistoryEndUsc(end *GTMioUSCCliqueMetadataRef, usc objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("cliqueExecutionHistoryEnd:usc:"), end, usc)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/cliqueExecutionHistoryStyle
func (g GTMioShaderExecutionHistory) CliqueExecutionHistoryStyle() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("cliqueExecutionHistoryStyle"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/dumpTree:
func (g GTMioShaderExecutionHistory) DumpTree(tree objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("dumpTree:"), tree)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/generateClique:
func (g GTMioShaderExecutionHistory) GenerateClique(clique *GTMioUSCCliqueMetadataRef) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("generateClique:"), clique)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/generateCliqueIndex:uscIndex:
func (g GTMioShaderExecutionHistory) GenerateCliqueIndexUscIndex(index uint32, index2 uint32) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("generateCliqueIndex:uscIndex:"), index, index2)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/generateDrawIndex:programType:
func (g GTMioShaderExecutionHistory) GenerateDrawIndexProgramType(index uint32, type_ uint16) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("generateDrawIndex:programType:"), index, type_)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/generateDrawIndex:programType:progressController:
func (g GTMioShaderExecutionHistory) GenerateDrawIndexProgramTypeProgressController(index uint32, type_ uint16, controller objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("generateDrawIndex:programType:progressController:"), index, type_, controller)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/generatePipelineStateId:programType:
func (g GTMioShaderExecutionHistory) GeneratePipelineStateIdProgramType(id uint64, type_ uint16) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("generatePipelineStateId:programType:"), id, type_)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/generatePipelineStateId:programType:progressController:
func (g GTMioShaderExecutionHistory) GeneratePipelineStateIdProgramTypeProgressController(id uint64, type_ uint16, controller objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("generatePipelineStateId:programType:progressController:"), id, type_, controller)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/handleCachedObject:
func (g GTMioShaderExecutionHistory) HandleCachedObject(object objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("handleCachedObject:"), object)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/loopBack:instructionEnd:loopCount:currentLoopCount:binary:
func (g GTMioShaderExecutionHistory) LoopBackInstructionEndLoopCountCurrentLoopCountBinary(back uint32, end uint32, count uint32, count2 uint32, binary objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("loopBack:instructionEnd:loopCount:currentLoopCount:binary:"), back, end, count, count2, binary)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/nodeForStyle:
func (g GTMioShaderExecutionHistory) NodeForStyle(style uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("nodeForStyle:"), style)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/popFunction:binaryRange:binary:
func (g GTMioShaderExecutionHistory) PopFunctionBinaryRangeBinary(function unsafe.Pointer, range_ unsafe.Pointer, binary objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("popFunction:binaryRange:binary:"), function, range_, binary)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/popLoop:instructionEnd:loopCount:binary:
func (g GTMioShaderExecutionHistory) PopLoopInstructionEndLoopCountBinary(loop uint32, end uint32, count uint32, binary objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("popLoop:instructionEnd:loopCount:binary:"), loop, end, count, binary)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/processInstruction:binaryRange:binary:numHits:
func (g GTMioShaderExecutionHistory) ProcessInstructionBinaryRangeBinaryNumHits(instruction uint32, range_ unsafe.Pointer, binary objectivec.IObject, hits uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("processInstruction:binaryRange:binary:numHits:"), instruction, range_, binary, hits)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/pushFunction:binaryRangeIndex:inlined:binary:callerLocation:callerBinaryRangeIndex:callerBinary:
func (g GTMioShaderExecutionHistory) PushFunctionBinaryRangeIndexInlinedBinaryCallerLocationCallerBinaryRangeIndexCallerBinary(function unsafe.Pointer, index uint32, inlined bool, binary objectivec.IObject, location unsafe.Pointer, index2 uint32, binary2 objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("pushFunction:binaryRangeIndex:inlined:binary:callerLocation:callerBinaryRangeIndex:callerBinary:"), function, index, inlined, binary, location, index2, binary2)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/pushLoop:instructionEnd:loopCount:binary:
func (g GTMioShaderExecutionHistory) PushLoopInstructionEndLoopCountBinary(loop uint32, end uint32, count uint32, binary objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("pushLoop:instructionEnd:loopCount:binary:"), loop, end, count, binary)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/timestamp:next:instructionCount:
func (g GTMioShaderExecutionHistory) TimestampNextInstructionCount(timestamp uint64, next uint64, count uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("timestamp:next:instructionCount:"), timestamp, next, count)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/initWithTraceData:style:options:delegate:
func (g GTMioShaderExecutionHistory) InitWithTraceDataStyleOptionsDelegate(data objectivec.IObject, style uint32, options uint32, delegate objectivec.IObject) GTMioShaderExecutionHistory {
	rv := objc.Send[GTMioShaderExecutionHistory](g.ID, objc.Sel("initWithTraceData:style:options:delegate:"), data, style, options, delegate)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/callStack
func (g GTMioShaderExecutionHistory) CallStack() IGTMioShaderExecutionHistoryRootNode {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("callStack"))
	return GTMioShaderExecutionHistoryRootNodeFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/compact
func (g GTMioShaderExecutionHistory) Compact() IGTMioShaderExecutionHistoryRootNode {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("compact"))
	return GTMioShaderExecutionHistoryRootNodeFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/delegate
func (g GTMioShaderExecutionHistory) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (g GTMioShaderExecutionHistory) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](g.ID, objc.Sel("setDelegate:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/full
func (g GTMioShaderExecutionHistory) Full() IGTMioShaderExecutionHistoryRootNode {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("full"))
	return GTMioShaderExecutionHistoryRootNodeFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/fullInstructionHistory
func (g GTMioShaderExecutionHistory) FullInstructionHistory() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("fullInstructionHistory"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/options
func (g GTMioShaderExecutionHistory) Options() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("options"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/style
func (g GTMioShaderExecutionHistory) Style() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("style"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistory/traceData
func (g GTMioShaderExecutionHistory) TraceData() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("traceData"))
	return objectivec.Object{ID: rv}
}

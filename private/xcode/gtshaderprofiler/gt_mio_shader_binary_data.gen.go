// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioShaderBinaryData] class.
var (
	_GTMioShaderBinaryDataClass     GTMioShaderBinaryDataClass
	_GTMioShaderBinaryDataClassOnce sync.Once
)

func getGTMioShaderBinaryDataClass() GTMioShaderBinaryDataClass {
	_GTMioShaderBinaryDataClassOnce.Do(func() {
		_GTMioShaderBinaryDataClass = GTMioShaderBinaryDataClass{class: objc.GetClass("GTMioShaderBinaryData")}
	})
	return _GTMioShaderBinaryDataClass
}

// GetGTMioShaderBinaryDataClass returns the class object for GTMioShaderBinaryData.
func GetGTMioShaderBinaryDataClass() GTMioShaderBinaryDataClass {
	return getGTMioShaderBinaryDataClass()
}

type GTMioShaderBinaryDataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioShaderBinaryDataClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioShaderBinaryDataClass) Alloc() GTMioShaderBinaryData {
	rv := objc.Send[GTMioShaderBinaryData](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioShaderBinaryData.Address]
//   - [GTMioShaderBinaryData.AddressForInstructionAtIndex]
//   - [GTMioShaderBinaryData.BinaryRangeIndexForInstructionIndex]
//   - [GTMioShaderBinaryData.CachedISAFileURL]
//   - [GTMioShaderBinaryData.Cost]
//   - [GTMioShaderBinaryData.CostCount]
//   - [GTMioShaderBinaryData.CostForBinaryRangeScopeScopeIdentifierCostNumInstructions]
//   - [GTMioShaderBinaryData.CostForLineFullPathIndexScopeScopeIdentifierCostNumInstructions]
//   - [GTMioShaderBinaryData.CostForScopeScopeIdentifierCost]
//   - [GTMioShaderBinaryData.Costs]
//   - [GTMioShaderBinaryData.DebugBinaryRangeCount]
//   - [GTMioShaderBinaryData.DebugBinaryRanges]
//   - [GTMioShaderBinaryData.DebugFilePathForDebugLocationAtIndex]
//   - [GTMioShaderBinaryData.DebugFunctionNameForDebugLocationAtIndex]
//   - [GTMioShaderBinaryData.DebugLocationAtIndex]
//   - [GTMioShaderBinaryData.DebugLocationCount]
//   - [GTMioShaderBinaryData.DebugLocations]
//   - [GTMioShaderBinaryData.DebugRangeForInstructionAtIndex]
//   - [GTMioShaderBinaryData.DebugStringForStringIndex]
//   - [GTMioShaderBinaryData.DebugStrings]
//   - [GTMioShaderBinaryData.Duration]
//   - [GTMioShaderBinaryData.EnumerateBinaryRangesForFileLineEnumerator]
//   - [GTMioShaderBinaryData.EnumerateDrawsWithProgramTypeEnumerator]
//   - [GTMioShaderBinaryData.EnumerateEncoderCosts]
//   - [GTMioShaderBinaryData.EnumerateEntryPoints]
//   - [GTMioShaderBinaryData.EnumerateInstructionsForBinaryRangeEnumerator]
//   - [GTMioShaderBinaryData.EnumerateLinesForFileEnumerator]
//   - [GTMioShaderBinaryData.EnumeratePerDrawCosts]
//   - [GTMioShaderBinaryData.EnumeratePipelineStateCosts]
//   - [GTMioShaderBinaryData.EnumeratePipelineStatesWithProgramTypeEnumerator]
//   - [GTMioShaderBinaryData.EnumerateTraces]
//   - [GTMioShaderBinaryData.FullPathIndexForFilePath]
//   - [GTMioShaderBinaryData.HasRaytracing]
//   - [GTMioShaderBinaryData.Index]
//   - [GTMioShaderBinaryData.InstructionCostScopeScopeIdentifierCost]
//   - [GTMioShaderBinaryData.InstructionCosts]
//   - [GTMioShaderBinaryData.InstructionCostsForDraw]
//   - [GTMioShaderBinaryData.InstructionCostsForEncoder]
//   - [GTMioShaderBinaryData.InstructionCostsForPipelineState]
//   - [GTMioShaderBinaryData.InstructionCostsForScopeScopeIdentifier]
//   - [GTMioShaderBinaryData.InstructionCountForScopeScopeIdentifierDataMaster]
//   - [GTMioShaderBinaryData.InstructionExecuted]
//   - [GTMioShaderBinaryData.InstructionInfo]
//   - [GTMioShaderBinaryData.InstructionInfoCount]
//   - [GTMioShaderBinaryData.Internal]
//   - [GTMioShaderBinaryData.IsaForInstructionAtIndex]
//   - [GTMioShaderBinaryData.IsaForInstructionAtIndexCount]
//   - [GTMioShaderBinaryData.IsaPrefixForInstructionAtIndex]
//   - [GTMioShaderBinaryData.IsaStrings]
//   - [GTMioShaderBinaryData.LiveRegisterForInstructionAtIndex]
//   - [GTMioShaderBinaryData.ProgramType]
//   - [GTMioShaderBinaryData.TotalCostForScopeScopeIdentifierDataMaster]
//   - [GTMioShaderBinaryData.TraceCount]
//   - [GTMioShaderBinaryData.TraceData]
//   - [GTMioShaderBinaryData.Traces]
//   - [GTMioShaderBinaryData.TracesForProgramTypeCount]
//   - [GTMioShaderBinaryData.UsedInDataMaster]
//   - [GTMioShaderBinaryData.UsedInDylib]
//   - [GTMioShaderBinaryData.UsedInEncoder]
//   - [GTMioShaderBinaryData.UsedInPipelineState]
//   - [GTMioShaderBinaryData.UsedInProgramType]
//   - [GTMioShaderBinaryData.InitWithBinaryDataParentIndex]
//   - [GTMioShaderBinaryData.DebugDescription]
//   - [GTMioShaderBinaryData.Description]
//   - [GTMioShaderBinaryData.Hash]
//   - [GTMioShaderBinaryData.Superclass]
type GTMioShaderBinaryData struct {
	objectivec.Object
}

// GTMioShaderBinaryDataFromID constructs a [GTMioShaderBinaryData] from an objc.ID.
func GTMioShaderBinaryDataFromID(id objc.ID) GTMioShaderBinaryData {
	return GTMioShaderBinaryData{objectivec.Object{ID: id}}
}

// Ensure GTMioShaderBinaryData implements IGTMioShaderBinaryData.
var _ IGTMioShaderBinaryData = GTMioShaderBinaryData{}

// An interface definition for the [GTMioShaderBinaryData] class.
//
// # Methods
//
//   - [IGTMioShaderBinaryData.Address]
//   - [IGTMioShaderBinaryData.AddressForInstructionAtIndex]
//   - [IGTMioShaderBinaryData.BinaryRangeIndexForInstructionIndex]
//   - [IGTMioShaderBinaryData.CachedISAFileURL]
//   - [IGTMioShaderBinaryData.Cost]
//   - [IGTMioShaderBinaryData.CostCount]
//   - [IGTMioShaderBinaryData.CostForBinaryRangeScopeScopeIdentifierCostNumInstructions]
//   - [IGTMioShaderBinaryData.CostForLineFullPathIndexScopeScopeIdentifierCostNumInstructions]
//   - [IGTMioShaderBinaryData.CostForScopeScopeIdentifierCost]
//   - [IGTMioShaderBinaryData.Costs]
//   - [IGTMioShaderBinaryData.DebugBinaryRangeCount]
//   - [IGTMioShaderBinaryData.DebugBinaryRanges]
//   - [IGTMioShaderBinaryData.DebugFilePathForDebugLocationAtIndex]
//   - [IGTMioShaderBinaryData.DebugFunctionNameForDebugLocationAtIndex]
//   - [IGTMioShaderBinaryData.DebugLocationAtIndex]
//   - [IGTMioShaderBinaryData.DebugLocationCount]
//   - [IGTMioShaderBinaryData.DebugLocations]
//   - [IGTMioShaderBinaryData.DebugRangeForInstructionAtIndex]
//   - [IGTMioShaderBinaryData.DebugStringForStringIndex]
//   - [IGTMioShaderBinaryData.DebugStrings]
//   - [IGTMioShaderBinaryData.Duration]
//   - [IGTMioShaderBinaryData.EnumerateBinaryRangesForFileLineEnumerator]
//   - [IGTMioShaderBinaryData.EnumerateDrawsWithProgramTypeEnumerator]
//   - [IGTMioShaderBinaryData.EnumerateEncoderCosts]
//   - [IGTMioShaderBinaryData.EnumerateEntryPoints]
//   - [IGTMioShaderBinaryData.EnumerateInstructionsForBinaryRangeEnumerator]
//   - [IGTMioShaderBinaryData.EnumerateLinesForFileEnumerator]
//   - [IGTMioShaderBinaryData.EnumeratePerDrawCosts]
//   - [IGTMioShaderBinaryData.EnumeratePipelineStateCosts]
//   - [IGTMioShaderBinaryData.EnumeratePipelineStatesWithProgramTypeEnumerator]
//   - [IGTMioShaderBinaryData.EnumerateTraces]
//   - [IGTMioShaderBinaryData.FullPathIndexForFilePath]
//   - [IGTMioShaderBinaryData.HasRaytracing]
//   - [IGTMioShaderBinaryData.Index]
//   - [IGTMioShaderBinaryData.InstructionCostScopeScopeIdentifierCost]
//   - [IGTMioShaderBinaryData.InstructionCosts]
//   - [IGTMioShaderBinaryData.InstructionCostsForDraw]
//   - [IGTMioShaderBinaryData.InstructionCostsForEncoder]
//   - [IGTMioShaderBinaryData.InstructionCostsForPipelineState]
//   - [IGTMioShaderBinaryData.InstructionCostsForScopeScopeIdentifier]
//   - [IGTMioShaderBinaryData.InstructionCountForScopeScopeIdentifierDataMaster]
//   - [IGTMioShaderBinaryData.InstructionExecuted]
//   - [IGTMioShaderBinaryData.InstructionInfo]
//   - [IGTMioShaderBinaryData.InstructionInfoCount]
//   - [IGTMioShaderBinaryData.Internal]
//   - [IGTMioShaderBinaryData.IsaForInstructionAtIndex]
//   - [IGTMioShaderBinaryData.IsaForInstructionAtIndexCount]
//   - [IGTMioShaderBinaryData.IsaPrefixForInstructionAtIndex]
//   - [IGTMioShaderBinaryData.IsaStrings]
//   - [IGTMioShaderBinaryData.LiveRegisterForInstructionAtIndex]
//   - [IGTMioShaderBinaryData.ProgramType]
//   - [IGTMioShaderBinaryData.TotalCostForScopeScopeIdentifierDataMaster]
//   - [IGTMioShaderBinaryData.TraceCount]
//   - [IGTMioShaderBinaryData.TraceData]
//   - [IGTMioShaderBinaryData.Traces]
//   - [IGTMioShaderBinaryData.TracesForProgramTypeCount]
//   - [IGTMioShaderBinaryData.UsedInDataMaster]
//   - [IGTMioShaderBinaryData.UsedInDylib]
//   - [IGTMioShaderBinaryData.UsedInEncoder]
//   - [IGTMioShaderBinaryData.UsedInPipelineState]
//   - [IGTMioShaderBinaryData.UsedInProgramType]
//   - [IGTMioShaderBinaryData.InitWithBinaryDataParentIndex]
//   - [IGTMioShaderBinaryData.DebugDescription]
//   - [IGTMioShaderBinaryData.Description]
//   - [IGTMioShaderBinaryData.Hash]
//   - [IGTMioShaderBinaryData.Superclass]
type IGTMioShaderBinaryData interface {
	objectivec.IObject

	// Topic: Methods

	Address() uint64
	AddressForInstructionAtIndex(index uint32) uint64
	BinaryRangeIndexForInstructionIndex(index uint32) uint32
	CachedISAFileURL() foundation.NSURL
	Cost() unsafe.Pointer
	CostCount() uint64
	CostForBinaryRangeScopeScopeIdentifierCostNumInstructions(range_ uint32, scope uint16, identifier uint64, cost GTMioCostInfo, instructions *uint32) bool
	CostForLineFullPathIndexScopeScopeIdentifierCostNumInstructions(line uint32, index uint32, scope uint16, identifier uint64, cost GTMioCostInfo, instructions *uint32) bool
	CostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost GTMioCostInfo) bool
	Costs() unsafe.Pointer
	DebugBinaryRangeCount() uint64
	DebugBinaryRanges() unsafe.Pointer
	DebugFilePathForDebugLocationAtIndex(index uint32) objectivec.IObject
	DebugFunctionNameForDebugLocationAtIndex(index uint32) objectivec.IObject
	DebugLocationAtIndex(index uint32) unsafe.Pointer
	DebugLocationCount() uint64
	DebugLocations() unsafe.Pointer
	DebugRangeForInstructionAtIndex(index uint32) unsafe.Pointer
	DebugStringForStringIndex(index uint32) objectivec.IObject
	DebugStrings() foundation.INSArray
	Duration() uint64
	EnumerateBinaryRangesForFileLineEnumerator(file uint32, line uint32, enumerator VoidHandler)
	EnumerateDrawsWithProgramTypeEnumerator(type_ uint16, enumerator VoidHandler)
	EnumerateEncoderCosts(costs VoidHandler)
	EnumerateEntryPoints(points VoidHandler)
	EnumerateInstructionsForBinaryRangeEnumerator(range_ uint32, enumerator VoidHandler)
	EnumerateLinesForFileEnumerator(file uint32, enumerator VoidHandler)
	EnumeratePerDrawCosts(costs VoidHandler)
	EnumeratePipelineStateCosts(costs VoidHandler)
	EnumeratePipelineStatesWithProgramTypeEnumerator(type_ uint16, enumerator VoidHandler)
	EnumerateTraces(traces VoidHandler)
	FullPathIndexForFilePath(path objectivec.IObject) uint32
	HasRaytracing() bool
	Index() uint64
	InstructionCostScopeScopeIdentifierCost(cost uint32, scope uint16, identifier uint64, cost2 GTMioCostInfo) bool
	InstructionCosts() unsafe.Pointer
	InstructionCostsForDraw(draw uint32) unsafe.Pointer
	InstructionCostsForEncoder(encoder uint32) unsafe.Pointer
	InstructionCostsForPipelineState(state uint64) unsafe.Pointer
	InstructionCostsForScopeScopeIdentifier(scope uint16, identifier uint64) unsafe.Pointer
	InstructionCountForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) uint64
	InstructionExecuted() uint64
	InstructionInfo() unsafe.Pointer
	InstructionInfoCount() uint64
	Internal() unsafe.Pointer
	IsaForInstructionAtIndex(index uint32) objectivec.IObject
	IsaForInstructionAtIndexCount(index uint32, count uint32) objectivec.IObject
	IsaPrefixForInstructionAtIndex(index uint32) objectivec.IObject
	IsaStrings() foundation.INSArray
	LiveRegisterForInstructionAtIndex(index uint32) int
	ProgramType() uint16
	TotalCostForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) float64
	TraceCount() uint64
	TraceData() unsafe.Pointer
	Traces() unsafe.Pointer
	TracesForProgramTypeCount(type_ uint16, count *uint64) unsafe.Pointer
	UsedInDataMaster(master uint16) bool
	UsedInDylib() bool
	UsedInEncoder(encoder uint64) bool
	UsedInPipelineState(state uint64) bool
	UsedInProgramType(type_ uint16) bool
	InitWithBinaryDataParentIndex(data unsafe.Pointer, parent objectivec.IObject, index uint64) GTMioShaderBinaryData
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (g GTMioShaderBinaryData) Init() GTMioShaderBinaryData {
	rv := objc.Send[GTMioShaderBinaryData](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioShaderBinaryData) Autorelease() GTMioShaderBinaryData {
	rv := objc.Send[GTMioShaderBinaryData](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioShaderBinaryData creates a new GTMioShaderBinaryData instance.
func NewGTMioShaderBinaryData() GTMioShaderBinaryData {
	class := getGTMioShaderBinaryDataClass()
	rv := objc.Send[GTMioShaderBinaryData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTMioShaderBinaryDataWithBinaryDataParentIndex(data unsafe.Pointer, parent objectivec.IObject, index uint64) GTMioShaderBinaryData {
	instance := getGTMioShaderBinaryDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBinaryData:parent:index:"), data, parent, index)
	return GTMioShaderBinaryDataFromID(rv)
}

func (g GTMioShaderBinaryData) AddressForInstructionAtIndex(index uint32) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("addressForInstructionAtIndex:"), index)
	return rv
}
func (g GTMioShaderBinaryData) BinaryRangeIndexForInstructionIndex(index uint32) uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("binaryRangeIndexForInstructionIndex:"), index)
	return rv
}
func (g GTMioShaderBinaryData) CachedISAFileURL() foundation.NSURL {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cachedISAFileURL"))
	return foundation.NSURLFromID(rv)
}
func (g GTMioShaderBinaryData) CostForBinaryRangeScopeScopeIdentifierCostNumInstructions(range_ uint32, scope uint16, identifier uint64, cost GTMioCostInfo, instructions *uint32) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("costForBinaryRange:scope:scopeIdentifier:cost:numInstructions:"), range_, scope, identifier, cost, instructions)
	return rv
}
func (g GTMioShaderBinaryData) CostForLineFullPathIndexScopeScopeIdentifierCostNumInstructions(line uint32, index uint32, scope uint16, identifier uint64, cost GTMioCostInfo, instructions *uint32) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("costForLine:fullPathIndex:scope:scopeIdentifier:cost:numInstructions:"), line, index, scope, identifier, cost, instructions)
	return rv
}
func (g GTMioShaderBinaryData) CostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost GTMioCostInfo) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("costForScope:scopeIdentifier:cost:"), scope, identifier, cost)
	return rv
}
func (g GTMioShaderBinaryData) DebugFilePathForDebugLocationAtIndex(index uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugFilePathForDebugLocationAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (g GTMioShaderBinaryData) DebugFunctionNameForDebugLocationAtIndex(index uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugFunctionNameForDebugLocationAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (g GTMioShaderBinaryData) DebugLocationAtIndex(index uint32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("debugLocationAtIndex:"), index)
	return rv
}
func (g GTMioShaderBinaryData) DebugRangeForInstructionAtIndex(index uint32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("debugRangeForInstructionAtIndex:"), index)
	return rv
}
func (g GTMioShaderBinaryData) DebugStringForStringIndex(index uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugStringForStringIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (g GTMioShaderBinaryData) EnumerateBinaryRangesForFileLineEnumerator(file uint32, line uint32, enumerator VoidHandler) {
	_block2, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateBinaryRangesForFile:line:enumerator:"), file, line, _block2)
}
func (g GTMioShaderBinaryData) EnumerateDrawsWithProgramTypeEnumerator(type_ uint16, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateDrawsWithProgramType:enumerator:"), type_, _block1)
}
func (g GTMioShaderBinaryData) EnumerateEncoderCosts(costs VoidHandler) {
	_block0, _ := NewVoidBlock(costs)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateEncoderCosts:"), _block0)
}
func (g GTMioShaderBinaryData) EnumerateEntryPoints(points VoidHandler) {
	_block0, _ := NewVoidBlock(points)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateEntryPoints:"), _block0)
}
func (g GTMioShaderBinaryData) EnumerateInstructionsForBinaryRangeEnumerator(range_ uint32, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateInstructionsForBinaryRange:enumerator:"), range_, _block1)
}
func (g GTMioShaderBinaryData) EnumerateLinesForFileEnumerator(file uint32, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateLinesForFile:enumerator:"), file, _block1)
}
func (g GTMioShaderBinaryData) EnumeratePerDrawCosts(costs VoidHandler) {
	_block0, _ := NewVoidBlock(costs)
	objc.Send[objc.ID](g.ID, objc.Sel("enumeratePerDrawCosts:"), _block0)
}
func (g GTMioShaderBinaryData) EnumeratePipelineStateCosts(costs VoidHandler) {
	_block0, _ := NewVoidBlock(costs)
	objc.Send[objc.ID](g.ID, objc.Sel("enumeratePipelineStateCosts:"), _block0)
}
func (g GTMioShaderBinaryData) EnumeratePipelineStatesWithProgramTypeEnumerator(type_ uint16, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumeratePipelineStatesWithProgramType:enumerator:"), type_, _block1)
}
func (g GTMioShaderBinaryData) EnumerateTraces(traces VoidHandler) {
	_block0, _ := NewVoidBlock(traces)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateTraces:"), _block0)
}
func (g GTMioShaderBinaryData) FullPathIndexForFilePath(path objectivec.IObject) uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("fullPathIndexForFilePath:"), path)
	return rv
}
func (g GTMioShaderBinaryData) InstructionCostScopeScopeIdentifierCost(cost uint32, scope uint16, identifier uint64, cost2 GTMioCostInfo) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("instructionCost:scope:scopeIdentifier:cost:"), cost, scope, identifier, cost2)
	return rv
}
func (g GTMioShaderBinaryData) InstructionCostsForDraw(draw uint32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("instructionCostsForDraw:"), draw)
	return rv
}
func (g GTMioShaderBinaryData) InstructionCostsForEncoder(encoder uint32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("instructionCostsForEncoder:"), encoder)
	return rv
}
func (g GTMioShaderBinaryData) InstructionCostsForPipelineState(state uint64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("instructionCostsForPipelineState:"), state)
	return rv
}
func (g GTMioShaderBinaryData) InstructionCostsForScopeScopeIdentifier(scope uint16, identifier uint64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("instructionCostsForScope:scopeIdentifier:"), scope, identifier)
	return rv
}
func (g GTMioShaderBinaryData) InstructionCountForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("instructionCountForScope:scopeIdentifier:dataMaster:"), scope, identifier, master)
	return rv
}
func (g GTMioShaderBinaryData) IsaForInstructionAtIndex(index uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("isaForInstructionAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (g GTMioShaderBinaryData) IsaForInstructionAtIndexCount(index uint32, count uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("isaForInstructionAtIndex:count:"), index, count)
	return objectivec.Object{ID: rv}
}
func (g GTMioShaderBinaryData) IsaPrefixForInstructionAtIndex(index uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("isaPrefixForInstructionAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (g GTMioShaderBinaryData) LiveRegisterForInstructionAtIndex(index uint32) int {
	rv := objc.Send[int](g.ID, objc.Sel("liveRegisterForInstructionAtIndex:"), index)
	return rv
}
func (g GTMioShaderBinaryData) TotalCostForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("totalCostForScope:scopeIdentifier:dataMaster:"), scope, identifier, master)
	return rv
}
func (g GTMioShaderBinaryData) TracesForProgramTypeCount(type_ uint16, count *uint64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("tracesForProgramType:count:"), type_, count)
	return rv
}
func (g GTMioShaderBinaryData) UsedInDataMaster(master uint16) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("usedInDataMaster:"), master)
	return rv
}
func (g GTMioShaderBinaryData) UsedInEncoder(encoder uint64) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("usedInEncoder:"), encoder)
	return rv
}
func (g GTMioShaderBinaryData) UsedInPipelineState(state uint64) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("usedInPipelineState:"), state)
	return rv
}
func (g GTMioShaderBinaryData) UsedInProgramType(type_ uint16) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("usedInProgramType:"), type_)
	return rv
}
func (g GTMioShaderBinaryData) InitWithBinaryDataParentIndex(data unsafe.Pointer, parent objectivec.IObject, index uint64) GTMioShaderBinaryData {
	rv := objc.Send[GTMioShaderBinaryData](g.ID, objc.Sel("initWithBinaryData:parent:index:"), data, parent, index)
	return rv
}

func (g GTMioShaderBinaryData) Address() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("address"))
	return rv
}
func (g GTMioShaderBinaryData) Cost() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("cost"))
	return rv
}
func (g GTMioShaderBinaryData) CostCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("costCount"))
	return rv
}
func (g GTMioShaderBinaryData) Costs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("costs"))
	return rv
}
func (g GTMioShaderBinaryData) DebugBinaryRangeCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("debugBinaryRangeCount"))
	return rv
}
func (g GTMioShaderBinaryData) DebugBinaryRanges() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("debugBinaryRanges"))
	return rv
}
func (g GTMioShaderBinaryData) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTMioShaderBinaryData) DebugLocationCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("debugLocationCount"))
	return rv
}
func (g GTMioShaderBinaryData) DebugLocations() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("debugLocations"))
	return rv
}
func (g GTMioShaderBinaryData) DebugStrings() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugStrings"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTMioShaderBinaryData) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTMioShaderBinaryData) Duration() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("duration"))
	return rv
}
func (g GTMioShaderBinaryData) HasRaytracing() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("hasRaytracing"))
	return rv
}
func (g GTMioShaderBinaryData) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}
func (g GTMioShaderBinaryData) Index() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("index"))
	return rv
}
func (g GTMioShaderBinaryData) InstructionCosts() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("instructionCosts"))
	return rv
}
func (g GTMioShaderBinaryData) InstructionExecuted() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("instructionExecuted"))
	return rv
}
func (g GTMioShaderBinaryData) InstructionInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("instructionInfo"))
	return rv
}
func (g GTMioShaderBinaryData) InstructionInfoCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("instructionInfoCount"))
	return rv
}
func (g GTMioShaderBinaryData) Internal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("internal"))
	return rv
}
func (g GTMioShaderBinaryData) IsaStrings() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("isaStrings"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTMioShaderBinaryData) ProgramType() uint16 {
	rv := objc.Send[uint16](g.ID, objc.Sel("programType"))
	return rv
}
func (g GTMioShaderBinaryData) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](g.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (g GTMioShaderBinaryData) TraceCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("traceCount"))
	return rv
}
func (g GTMioShaderBinaryData) TraceData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("traceData"))
	return rv
}
func (g GTMioShaderBinaryData) Traces() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("traces"))
	return rv
}
func (g GTMioShaderBinaryData) UsedInDylib() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("usedInDylib"))
	return rv
}

// EnumerateBinaryRangesForFileLineEnumeratorSync is a synchronous wrapper around [GTMioShaderBinaryData.EnumerateBinaryRangesForFileLineEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioShaderBinaryData) EnumerateBinaryRangesForFileLineEnumeratorSync(ctx context.Context, file uint32, line uint32) error {
	done := make(chan struct{}, 1)
	g.EnumerateBinaryRangesForFileLineEnumerator(file, line, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateDrawsWithProgramTypeEnumeratorSync is a synchronous wrapper around [GTMioShaderBinaryData.EnumerateDrawsWithProgramTypeEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioShaderBinaryData) EnumerateDrawsWithProgramTypeEnumeratorSync(ctx context.Context, type_ uint16) error {
	done := make(chan struct{}, 1)
	g.EnumerateDrawsWithProgramTypeEnumerator(type_, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateEncoderCostsSync is a synchronous wrapper around [GTMioShaderBinaryData.EnumerateEncoderCosts].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioShaderBinaryData) EnumerateEncoderCostsSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.EnumerateEncoderCosts(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateEntryPointsSync is a synchronous wrapper around [GTMioShaderBinaryData.EnumerateEntryPoints].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioShaderBinaryData) EnumerateEntryPointsSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.EnumerateEntryPoints(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateInstructionsForBinaryRangeEnumeratorSync is a synchronous wrapper around [GTMioShaderBinaryData.EnumerateInstructionsForBinaryRangeEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioShaderBinaryData) EnumerateInstructionsForBinaryRangeEnumeratorSync(ctx context.Context, range_ uint32) error {
	done := make(chan struct{}, 1)
	g.EnumerateInstructionsForBinaryRangeEnumerator(range_, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateLinesForFileEnumeratorSync is a synchronous wrapper around [GTMioShaderBinaryData.EnumerateLinesForFileEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioShaderBinaryData) EnumerateLinesForFileEnumeratorSync(ctx context.Context, file uint32) error {
	done := make(chan struct{}, 1)
	g.EnumerateLinesForFileEnumerator(file, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumeratePerDrawCostsSync is a synchronous wrapper around [GTMioShaderBinaryData.EnumeratePerDrawCosts].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioShaderBinaryData) EnumeratePerDrawCostsSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.EnumeratePerDrawCosts(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumeratePipelineStateCostsSync is a synchronous wrapper around [GTMioShaderBinaryData.EnumeratePipelineStateCosts].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioShaderBinaryData) EnumeratePipelineStateCostsSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.EnumeratePipelineStateCosts(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumeratePipelineStatesWithProgramTypeEnumeratorSync is a synchronous wrapper around [GTMioShaderBinaryData.EnumeratePipelineStatesWithProgramTypeEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioShaderBinaryData) EnumeratePipelineStatesWithProgramTypeEnumeratorSync(ctx context.Context, type_ uint16) error {
	done := make(chan struct{}, 1)
	g.EnumeratePipelineStatesWithProgramTypeEnumerator(type_, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateTracesSync is a synchronous wrapper around [GTMioShaderBinaryData.EnumerateTraces].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioShaderBinaryData) EnumerateTracesSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.EnumerateTraces(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

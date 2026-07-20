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

// The class instance for the [GTMioUSCTraceData] class.
var (
	_GTMioUSCTraceDataClass     GTMioUSCTraceDataClass
	_GTMioUSCTraceDataClassOnce sync.Once
)

func getGTMioUSCTraceDataClass() GTMioUSCTraceDataClass {
	_GTMioUSCTraceDataClassOnce.Do(func() {
		_GTMioUSCTraceDataClass = GTMioUSCTraceDataClass{class: objc.GetClass("GTMioUSCTraceData")}
	})
	return _GTMioUSCTraceDataClass
}

// GetGTMioUSCTraceDataClass returns the class object for GTMioUSCTraceData.
func GetGTMioUSCTraceDataClass() GTMioUSCTraceDataClass {
	return getGTMioUSCTraceDataClass()
}

type GTMioUSCTraceDataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioUSCTraceDataClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioUSCTraceDataClass) Alloc() GTMioUSCTraceData {
	rv := objc.Send[GTMioUSCTraceData](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioUSCTraceData.BinaryTrace]
//   - [GTMioUSCTraceData.BinaryTraceCount]
//   - [GTMioUSCTraceData.ChildCliqueOfClique]
//   - [GTMioUSCTraceData.CliqueFirstPCCount]
//   - [GTMioUSCTraceData.CliqueFirstPCs]
//   - [GTMioUSCTraceData.Cliques]
//   - [GTMioUSCTraceData.CliquesCount]
//   - [GTMioUSCTraceData.CostCount]
//   - [GTMioUSCTraceData.CostForScopeScopeIdentifierCost]
//   - [GTMioUSCTraceData.Costs]
//   - [GTMioUSCTraceData.DatabaseInternal]
//   - [GTMioUSCTraceData.DrawTrace]
//   - [GTMioUSCTraceData.DrawTraceCount]
//   - [GTMioUSCTraceData.EnumerateCliquesForTimeRangeBeginEndEnumerator]
//   - [GTMioUSCTraceData.EnumerateInstructionTracesForCliqueRequiresTimestampsEnumerator]
//   - [GTMioUSCTraceData.EnumerateKickAtFunctionIndexEnumerator]
//   - [GTMioUSCTraceData.EnumerateKickCliquesEnumerator]
//   - [GTMioUSCTraceData.EnumerateKickCliquesAtFunctionIndexDataMasterEnumerator]
//   - [GTMioUSCTraceData.EnumerateKickTilesEnumerator]
//   - [GTMioUSCTraceData.EnumerateKickTilesAtFunctionIndexDataMasterEnumerator]
//   - [GTMioUSCTraceData.EnumerateTileCliquesEnumerator]
//   - [GTMioUSCTraceData.FirstBinaryAddressForCliqueAtIndex]
//   - [GTMioUSCTraceData.FirstBinaryIndexForCliqueAtIndex]
//   - [GTMioUSCTraceData.FirstPCForCliqueAtIndex]
//   - [GTMioUSCTraceData.Index]
//   - [GTMioUSCTraceData.InstructionCountForScopeScopeIdentifierDataMaster]
//   - [GTMioUSCTraceData.InstructionTraceInfoForClique]
//   - [GTMioUSCTraceData.InstructionTraceInfoForCliqueAtIndex]
//   - [GTMioUSCTraceData.KickCliqueIndexes]
//   - [GTMioUSCTraceData.KickCliqueIndexesCount]
//   - [GTMioUSCTraceData.KickDurationForEncoder]
//   - [GTMioUSCTraceData.KickDurationForEncoderDataMaster]
//   - [GTMioUSCTraceData.KickTileIndexes]
//   - [GTMioUSCTraceData.KickTileIndexesCount]
//   - [GTMioUSCTraceData.Kicks]
//   - [GTMioUSCTraceData.KicksCount]
//   - [GTMioUSCTraceData.MGPUIndex]
//   - [GTMioUSCTraceData.PcForInstructionBinaryIndex]
//   - [GTMioUSCTraceData.PipelineStateIdForCliqueAtIndex]
//   - [GTMioUSCTraceData.TileCliqueIndexes]
//   - [GTMioUSCTraceData.TileCliqueIndexesCount]
//   - [GTMioUSCTraceData.Tiles]
//   - [GTMioUSCTraceData.TilesCount]
//   - [GTMioUSCTraceData.TimelineCounters]
//   - [GTMioUSCTraceData.TotalCostForScopeScopeIdentifierDataMaster]
//   - [GTMioUSCTraceData.TraceData]
//   - [GTMioUSCTraceData.UscCost]
//   - [GTMioUSCTraceData.InitWithUSCDataMGPUIndexParent]
//   - [GTMioUSCTraceData.DebugDescription]
//   - [GTMioUSCTraceData.Description]
//   - [GTMioUSCTraceData.Hash]
//   - [GTMioUSCTraceData.Superclass]
type GTMioUSCTraceData struct {
	objectivec.Object
}

// GTMioUSCTraceDataFromID constructs a [GTMioUSCTraceData] from an objc.ID.
func GTMioUSCTraceDataFromID(id objc.ID) GTMioUSCTraceData {
	return GTMioUSCTraceData{objectivec.Object{ID: id}}
}

// Ensure GTMioUSCTraceData implements IGTMioUSCTraceData.
var _ IGTMioUSCTraceData = GTMioUSCTraceData{}

// An interface definition for the [GTMioUSCTraceData] class.
//
// # Methods
//
//   - [IGTMioUSCTraceData.BinaryTrace]
//   - [IGTMioUSCTraceData.BinaryTraceCount]
//   - [IGTMioUSCTraceData.ChildCliqueOfClique]
//   - [IGTMioUSCTraceData.CliqueFirstPCCount]
//   - [IGTMioUSCTraceData.CliqueFirstPCs]
//   - [IGTMioUSCTraceData.Cliques]
//   - [IGTMioUSCTraceData.CliquesCount]
//   - [IGTMioUSCTraceData.CostCount]
//   - [IGTMioUSCTraceData.CostForScopeScopeIdentifierCost]
//   - [IGTMioUSCTraceData.Costs]
//   - [IGTMioUSCTraceData.DatabaseInternal]
//   - [IGTMioUSCTraceData.DrawTrace]
//   - [IGTMioUSCTraceData.DrawTraceCount]
//   - [IGTMioUSCTraceData.EnumerateCliquesForTimeRangeBeginEndEnumerator]
//   - [IGTMioUSCTraceData.EnumerateInstructionTracesForCliqueRequiresTimestampsEnumerator]
//   - [IGTMioUSCTraceData.EnumerateKickAtFunctionIndexEnumerator]
//   - [IGTMioUSCTraceData.EnumerateKickCliquesEnumerator]
//   - [IGTMioUSCTraceData.EnumerateKickCliquesAtFunctionIndexDataMasterEnumerator]
//   - [IGTMioUSCTraceData.EnumerateKickTilesEnumerator]
//   - [IGTMioUSCTraceData.EnumerateKickTilesAtFunctionIndexDataMasterEnumerator]
//   - [IGTMioUSCTraceData.EnumerateTileCliquesEnumerator]
//   - [IGTMioUSCTraceData.FirstBinaryAddressForCliqueAtIndex]
//   - [IGTMioUSCTraceData.FirstBinaryIndexForCliqueAtIndex]
//   - [IGTMioUSCTraceData.FirstPCForCliqueAtIndex]
//   - [IGTMioUSCTraceData.Index]
//   - [IGTMioUSCTraceData.InstructionCountForScopeScopeIdentifierDataMaster]
//   - [IGTMioUSCTraceData.InstructionTraceInfoForClique]
//   - [IGTMioUSCTraceData.InstructionTraceInfoForCliqueAtIndex]
//   - [IGTMioUSCTraceData.KickCliqueIndexes]
//   - [IGTMioUSCTraceData.KickCliqueIndexesCount]
//   - [IGTMioUSCTraceData.KickDurationForEncoder]
//   - [IGTMioUSCTraceData.KickDurationForEncoderDataMaster]
//   - [IGTMioUSCTraceData.KickTileIndexes]
//   - [IGTMioUSCTraceData.KickTileIndexesCount]
//   - [IGTMioUSCTraceData.Kicks]
//   - [IGTMioUSCTraceData.KicksCount]
//   - [IGTMioUSCTraceData.MGPUIndex]
//   - [IGTMioUSCTraceData.PcForInstructionBinaryIndex]
//   - [IGTMioUSCTraceData.PipelineStateIdForCliqueAtIndex]
//   - [IGTMioUSCTraceData.TileCliqueIndexes]
//   - [IGTMioUSCTraceData.TileCliqueIndexesCount]
//   - [IGTMioUSCTraceData.Tiles]
//   - [IGTMioUSCTraceData.TilesCount]
//   - [IGTMioUSCTraceData.TimelineCounters]
//   - [IGTMioUSCTraceData.TotalCostForScopeScopeIdentifierDataMaster]
//   - [IGTMioUSCTraceData.TraceData]
//   - [IGTMioUSCTraceData.UscCost]
//   - [IGTMioUSCTraceData.InitWithUSCDataMGPUIndexParent]
//   - [IGTMioUSCTraceData.DebugDescription]
//   - [IGTMioUSCTraceData.Description]
//   - [IGTMioUSCTraceData.Hash]
//   - [IGTMioUSCTraceData.Superclass]
type IGTMioUSCTraceData interface {
	objectivec.IObject

	// Topic: Methods

	BinaryTrace() unsafe.Pointer
	BinaryTraceCount() uint64
	ChildCliqueOfClique(clique *GTMioUSCCliqueMetadata) unsafe.Pointer
	CliqueFirstPCCount() uint64
	CliqueFirstPCs() unsafe.Pointer
	Cliques() unsafe.Pointer
	CliquesCount() uint64
	CostCount() uint64
	CostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost *GTMioCostInfo) bool
	Costs() unsafe.Pointer
	DatabaseInternal() uint64
	DrawTrace() unsafe.Pointer
	DrawTraceCount() uint64
	EnumerateCliquesForTimeRangeBeginEndEnumerator(begin uint64, end uint64, enumerator VoidHandler)
	EnumerateInstructionTracesForCliqueRequiresTimestampsEnumerator(clique *GTMioUSCCliqueMetadata, timestamps bool, enumerator VoidHandler)
	EnumerateKickAtFunctionIndexEnumerator(index uint32, enumerator VoidHandler)
	EnumerateKickCliquesEnumerator(cliques *GTMioUSCKickMetadata, enumerator VoidHandler)
	EnumerateKickCliquesAtFunctionIndexDataMasterEnumerator(index uint32, master uint16, enumerator VoidHandler)
	EnumerateKickTilesEnumerator(tiles *GTMioUSCKickMetadata, enumerator VoidHandler)
	EnumerateKickTilesAtFunctionIndexDataMasterEnumerator(index uint32, master uint16, enumerator VoidHandler)
	EnumerateTileCliquesEnumerator(cliques *GTMioUSCTileMetadata, enumerator VoidHandler)
	FirstBinaryAddressForCliqueAtIndex(index uint32) uint64
	FirstBinaryIndexForCliqueAtIndex(index uint32) uint32
	FirstPCForCliqueAtIndex(index uint32) uint64
	Index() uint64
	InstructionCountForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) uint64
	InstructionTraceInfoForClique(clique *GTMioUSCCliqueMetadata) unsafe.Pointer
	InstructionTraceInfoForCliqueAtIndex(index uint32) unsafe.Pointer
	KickCliqueIndexes() unsafe.Pointer
	KickCliqueIndexesCount() uint64
	KickDurationForEncoder(encoder uint32) uint64
	KickDurationForEncoderDataMaster(encoder uint32, master uint16) uint64
	KickTileIndexes() unsafe.Pointer
	KickTileIndexesCount() uint64
	Kicks() unsafe.Pointer
	KicksCount() uint64
	MGPUIndex() uint64
	PcForInstructionBinaryIndex(instruction uint32, index uint32) uint64
	PipelineStateIdForCliqueAtIndex(index uint32) uint64
	TileCliqueIndexes() unsafe.Pointer
	TileCliqueIndexesCount() uint64
	Tiles() unsafe.Pointer
	TilesCount() uint64
	TimelineCounters() IGTMioTimelineCounters
	TotalCostForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) float64
	TraceData() unsafe.Pointer
	UscCost() unsafe.Pointer
	InitWithUSCDataMGPUIndexParent(uSCData unsafe.Pointer, gPUIndex uint64, parent objectivec.IObject) GTMioUSCTraceData
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (g GTMioUSCTraceData) Init() GTMioUSCTraceData {
	rv := objc.Send[GTMioUSCTraceData](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioUSCTraceData) Autorelease() GTMioUSCTraceData {
	rv := objc.Send[GTMioUSCTraceData](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioUSCTraceData creates a new GTMioUSCTraceData instance.
func NewGTMioUSCTraceData() GTMioUSCTraceData {
	class := getGTMioUSCTraceDataClass()
	rv := objc.Send[GTMioUSCTraceData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTMioUSCTraceDataWithUSCDataMGPUIndexParent(uSCData unsafe.Pointer, gPUIndex uint64, parent objectivec.IObject) GTMioUSCTraceData {
	instance := getGTMioUSCTraceDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUSCData:mGPUIndex:parent:"), uSCData, gPUIndex, parent)
	return GTMioUSCTraceDataFromID(rv)
}

func (g GTMioUSCTraceData) ChildCliqueOfClique(clique *GTMioUSCCliqueMetadata) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("childCliqueOfClique:"), clique)
	return rv
}
func (g GTMioUSCTraceData) CostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost *GTMioCostInfo) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("costForScope:scopeIdentifier:cost:"), scope, identifier, cost)
	return rv
}
func (g GTMioUSCTraceData) EnumerateCliquesForTimeRangeBeginEndEnumerator(begin uint64, end uint64, enumerator VoidHandler) {
	_block2, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateCliquesForTimeRangeBegin:end:enumerator:"), begin, end, _block2)
}
func (g GTMioUSCTraceData) EnumerateInstructionTracesForCliqueRequiresTimestampsEnumerator(clique *GTMioUSCCliqueMetadata, timestamps bool, enumerator VoidHandler) {
	_block2, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateInstructionTracesForClique:requiresTimestamps:enumerator:"), clique, timestamps, _block2)
}
func (g GTMioUSCTraceData) EnumerateKickAtFunctionIndexEnumerator(index uint32, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateKickAtFunctionIndex:enumerator:"), index, _block1)
}
func (g GTMioUSCTraceData) EnumerateKickCliquesEnumerator(cliques *GTMioUSCKickMetadata, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateKickCliques:enumerator:"), cliques, _block1)
}
func (g GTMioUSCTraceData) EnumerateKickCliquesAtFunctionIndexDataMasterEnumerator(index uint32, master uint16, enumerator VoidHandler) {
	_block2, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateKickCliquesAtFunctionIndex:dataMaster:enumerator:"), index, master, _block2)
}
func (g GTMioUSCTraceData) EnumerateKickTilesEnumerator(tiles *GTMioUSCKickMetadata, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateKickTiles:enumerator:"), tiles, _block1)
}
func (g GTMioUSCTraceData) EnumerateKickTilesAtFunctionIndexDataMasterEnumerator(index uint32, master uint16, enumerator VoidHandler) {
	_block2, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateKickTilesAtFunctionIndex:dataMaster:enumerator:"), index, master, _block2)
}
func (g GTMioUSCTraceData) EnumerateTileCliquesEnumerator(cliques *GTMioUSCTileMetadata, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateTileCliques:enumerator:"), cliques, _block1)
}
func (g GTMioUSCTraceData) FirstBinaryAddressForCliqueAtIndex(index uint32) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("firstBinaryAddressForCliqueAtIndex:"), index)
	return rv
}
func (g GTMioUSCTraceData) FirstBinaryIndexForCliqueAtIndex(index uint32) uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("firstBinaryIndexForCliqueAtIndex:"), index)
	return rv
}
func (g GTMioUSCTraceData) FirstPCForCliqueAtIndex(index uint32) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("firstPCForCliqueAtIndex:"), index)
	return rv
}
func (g GTMioUSCTraceData) InstructionCountForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("instructionCountForScope:scopeIdentifier:dataMaster:"), scope, identifier, master)
	return rv
}
func (g GTMioUSCTraceData) InstructionTraceInfoForClique(clique *GTMioUSCCliqueMetadata) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("instructionTraceInfoForClique:"), clique)
	return rv
}
func (g GTMioUSCTraceData) InstructionTraceInfoForCliqueAtIndex(index uint32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("instructionTraceInfoForCliqueAtIndex:"), index)
	return rv
}
func (g GTMioUSCTraceData) KickDurationForEncoder(encoder uint32) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("kickDurationForEncoder:"), encoder)
	return rv
}
func (g GTMioUSCTraceData) KickDurationForEncoderDataMaster(encoder uint32, master uint16) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("kickDurationForEncoder:dataMaster:"), encoder, master)
	return rv
}
func (g GTMioUSCTraceData) PcForInstructionBinaryIndex(instruction uint32, index uint32) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("pcForInstruction:binaryIndex:"), instruction, index)
	return rv
}
func (g GTMioUSCTraceData) PipelineStateIdForCliqueAtIndex(index uint32) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("pipelineStateIdForCliqueAtIndex:"), index)
	return rv
}
func (g GTMioUSCTraceData) TotalCostForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("totalCostForScope:scopeIdentifier:dataMaster:"), scope, identifier, master)
	return rv
}
func (g GTMioUSCTraceData) InitWithUSCDataMGPUIndexParent(uSCData unsafe.Pointer, gPUIndex uint64, parent objectivec.IObject) GTMioUSCTraceData {
	rv := objc.Send[GTMioUSCTraceData](g.ID, objc.Sel("initWithUSCData:mGPUIndex:parent:"), uSCData, gPUIndex, parent)
	return rv
}

func (g GTMioUSCTraceData) BinaryTrace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("binaryTrace"))
	return rv
}
func (g GTMioUSCTraceData) BinaryTraceCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("binaryTraceCount"))
	return rv
}
func (g GTMioUSCTraceData) CliqueFirstPCCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("cliqueFirstPCCount"))
	return rv
}
func (g GTMioUSCTraceData) CliqueFirstPCs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("cliqueFirstPCs"))
	return rv
}
func (g GTMioUSCTraceData) Cliques() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("cliques"))
	return rv
}
func (g GTMioUSCTraceData) CliquesCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("cliquesCount"))
	return rv
}
func (g GTMioUSCTraceData) CostCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("costCount"))
	return rv
}
func (g GTMioUSCTraceData) Costs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("costs"))
	return rv
}
func (g GTMioUSCTraceData) DatabaseInternal() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("databaseInternal"))
	return rv
}
func (g GTMioUSCTraceData) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTMioUSCTraceData) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTMioUSCTraceData) DrawTrace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("drawTrace"))
	return rv
}
func (g GTMioUSCTraceData) DrawTraceCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("drawTraceCount"))
	return rv
}
func (g GTMioUSCTraceData) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}
func (g GTMioUSCTraceData) Index() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("index"))
	return rv
}
func (g GTMioUSCTraceData) KickCliqueIndexes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("kickCliqueIndexes"))
	return rv
}
func (g GTMioUSCTraceData) KickCliqueIndexesCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("kickCliqueIndexesCount"))
	return rv
}
func (g GTMioUSCTraceData) KickTileIndexes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("kickTileIndexes"))
	return rv
}
func (g GTMioUSCTraceData) KickTileIndexesCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("kickTileIndexesCount"))
	return rv
}
func (g GTMioUSCTraceData) Kicks() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("kicks"))
	return rv
}
func (g GTMioUSCTraceData) KicksCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("kicksCount"))
	return rv
}
func (g GTMioUSCTraceData) MGPUIndex() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("mGPUIndex"))
	return rv
}
func (g GTMioUSCTraceData) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](g.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (g GTMioUSCTraceData) TileCliqueIndexes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("tileCliqueIndexes"))
	return rv
}
func (g GTMioUSCTraceData) TileCliqueIndexesCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("tileCliqueIndexesCount"))
	return rv
}
func (g GTMioUSCTraceData) Tiles() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("tiles"))
	return rv
}
func (g GTMioUSCTraceData) TilesCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("tilesCount"))
	return rv
}
func (g GTMioUSCTraceData) TimelineCounters() IGTMioTimelineCounters {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("timelineCounters"))
	return GTMioTimelineCountersFromID(objc.ID(rv))
}
func (g GTMioUSCTraceData) TraceData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("traceData"))
	return rv
}
func (g GTMioUSCTraceData) UscCost() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("uscCost"))
	return rv
}

// EnumerateCliquesForTimeRangeBeginEndEnumeratorSync is a synchronous wrapper around [GTMioUSCTraceData.EnumerateCliquesForTimeRangeBeginEndEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioUSCTraceData) EnumerateCliquesForTimeRangeBeginEndEnumeratorSync(ctx context.Context, begin uint64, end uint64) error {
	done := make(chan struct{}, 1)
	g.EnumerateCliquesForTimeRangeBeginEndEnumerator(begin, end, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateInstructionTracesForCliqueRequiresTimestampsEnumeratorSync is a synchronous wrapper around [GTMioUSCTraceData.EnumerateInstructionTracesForCliqueRequiresTimestampsEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioUSCTraceData) EnumerateInstructionTracesForCliqueRequiresTimestampsEnumeratorSync(ctx context.Context, clique *GTMioUSCCliqueMetadata, timestamps bool) error {
	done := make(chan struct{}, 1)
	g.EnumerateInstructionTracesForCliqueRequiresTimestampsEnumerator(clique, timestamps, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateKickAtFunctionIndexEnumeratorSync is a synchronous wrapper around [GTMioUSCTraceData.EnumerateKickAtFunctionIndexEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioUSCTraceData) EnumerateKickAtFunctionIndexEnumeratorSync(ctx context.Context, index uint32) error {
	done := make(chan struct{}, 1)
	g.EnumerateKickAtFunctionIndexEnumerator(index, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateKickCliquesEnumeratorSync is a synchronous wrapper around [GTMioUSCTraceData.EnumerateKickCliquesEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioUSCTraceData) EnumerateKickCliquesEnumeratorSync(ctx context.Context, cliques *GTMioUSCKickMetadata) error {
	done := make(chan struct{}, 1)
	g.EnumerateKickCliquesEnumerator(cliques, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateKickCliquesAtFunctionIndexDataMasterEnumeratorSync is a synchronous wrapper around [GTMioUSCTraceData.EnumerateKickCliquesAtFunctionIndexDataMasterEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioUSCTraceData) EnumerateKickCliquesAtFunctionIndexDataMasterEnumeratorSync(ctx context.Context, index uint32, master uint16) error {
	done := make(chan struct{}, 1)
	g.EnumerateKickCliquesAtFunctionIndexDataMasterEnumerator(index, master, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateKickTilesEnumeratorSync is a synchronous wrapper around [GTMioUSCTraceData.EnumerateKickTilesEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioUSCTraceData) EnumerateKickTilesEnumeratorSync(ctx context.Context, tiles *GTMioUSCKickMetadata) error {
	done := make(chan struct{}, 1)
	g.EnumerateKickTilesEnumerator(tiles, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateKickTilesAtFunctionIndexDataMasterEnumeratorSync is a synchronous wrapper around [GTMioUSCTraceData.EnumerateKickTilesAtFunctionIndexDataMasterEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioUSCTraceData) EnumerateKickTilesAtFunctionIndexDataMasterEnumeratorSync(ctx context.Context, index uint32, master uint16) error {
	done := make(chan struct{}, 1)
	g.EnumerateKickTilesAtFunctionIndexDataMasterEnumerator(index, master, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateTileCliquesEnumeratorSync is a synchronous wrapper around [GTMioUSCTraceData.EnumerateTileCliquesEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioUSCTraceData) EnumerateTileCliquesEnumeratorSync(ctx context.Context, cliques *GTMioUSCTileMetadata) error {
	done := make(chan struct{}, 1)
	g.EnumerateTileCliquesEnumerator(cliques, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

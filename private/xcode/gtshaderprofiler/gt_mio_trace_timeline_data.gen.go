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

// The class instance for the [GTMioTraceTimelineData] class.
var (
	_GTMioTraceTimelineDataClass     GTMioTraceTimelineDataClass
	_GTMioTraceTimelineDataClassOnce sync.Once
)

func getGTMioTraceTimelineDataClass() GTMioTraceTimelineDataClass {
	_GTMioTraceTimelineDataClassOnce.Do(func() {
		_GTMioTraceTimelineDataClass = GTMioTraceTimelineDataClass{class: objc.GetClass("GTMioTraceTimelineData")}
	})
	return _GTMioTraceTimelineDataClass
}

// GetGTMioTraceTimelineDataClass returns the class object for GTMioTraceTimelineData.
func GetGTMioTraceTimelineDataClass() GTMioTraceTimelineDataClass {
	return getGTMioTraceTimelineDataClass()
}

type GTMioTraceTimelineDataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioTraceTimelineDataClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioTraceTimelineDataClass) Alloc() GTMioTraceTimelineData {
	rv := objc.SendIfResponds[GTMioTraceTimelineData](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioTraceTimelineData._cacheExeuctionHistory]
//   - [GTMioTraceTimelineData._waitPendingExecutionHistory]
//   - [GTMioTraceTimelineData.ArchivedData]
//   - [GTMioTraceTimelineData.ArchivedDataWithCompressionOriginalSizeError]
//   - [GTMioTraceTimelineData.Binaries]
//   - [GTMioTraceTimelineData.BinaryForDrawProgramType]
//   - [GTMioTraceTimelineData.BinaryForPipelineStateProgramType]
//   - [GTMioTraceTimelineData.ChildCliqueOfClique]
//   - [GTMioTraceTimelineData.CliqueFromCliqueIndex]
//   - [GTMioTraceTimelineData.CoalescedFunctionIndexForEncoderFunctionIndex]
//   - [GTMioTraceTimelineData.ComputePositionCount]
//   - [GTMioTraceTimelineData.ComputePositions]
//   - [GTMioTraceTimelineData.ConsistentStateAchieved]
//   - [GTMioTraceTimelineData.CostCount]
//   - [GTMioTraceTimelineData.CostForContextCost]
//   - [GTMioTraceTimelineData.CostForLevelLevelIdentifierScopeScopeIdentifierCost]
//   - [GTMioTraceTimelineData.CostForScopeScopeIdentifierCost]
//   - [GTMioTraceTimelineData.Costs]
//   - [GTMioTraceTimelineData.DataType]
//   - [GTMioTraceTimelineData.DatabaseInternal]
//   - [GTMioTraceTimelineData.DrawCount]
//   - [GTMioTraceTimelineData.DrawTraceCount]
//   - [GTMioTraceTimelineData.DrawTraces]
//   - [GTMioTraceTimelineData.Draws]
//   - [GTMioTraceTimelineData.DurationForDrawDataMaster]
//   - [GTMioTraceTimelineData.EncodeError]
//   - [GTMioTraceTimelineData.EncodeWithCoder]
//   - [GTMioTraceTimelineData.EncoderCount]
//   - [GTMioTraceTimelineData.EncoderFromFunctionIndex]
//   - [GTMioTraceTimelineData.Encoders]
//   - [GTMioTraceTimelineData.EnumerateBinariesForDrawEnumerator]
//   - [GTMioTraceTimelineData.EnumerateBinariesForEncoderEnumerator]
//   - [GTMioTraceTimelineData.EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceTimelineData.EnumerateBinariesForPipelineStateEnumerator]
//   - [GTMioTraceTimelineData.EnumerateBinaryRangesForCliqueUscDataEnumerator]
//   - [GTMioTraceTimelineData.EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceTimelineData.EnumerateDrawsForEncoderEnumerator]
//   - [GTMioTraceTimelineData.EnumerateDrawsForPipelineStateEnumerator]
//   - [GTMioTraceTimelineData.EnumerateEncoders]
//   - [GTMioTraceTimelineData.EnumerateInstructionsForCliqueUscDataEnumerator]
//   - [GTMioTraceTimelineData.EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceTimelineData.EnumerateKickAtFunctionIndexEnumerator]
//   - [GTMioTraceTimelineData.EnumeratePipelineStates]
//   - [GTMioTraceTimelineData.EnumerateUniqueTracesForBinaryEnumerator]
//   - [GTMioTraceTimelineData.ExecutionHistoryForCliqueUscDelegate]
//   - [GTMioTraceTimelineData.ExecutionHistoryForCliqueUscDelegateRequiresTimestampCount]
//   - [GTMioTraceTimelineData.ExecutionHistoryForCliqueUscIndexDelegate]
//   - [GTMioTraceTimelineData.ExecutionHistoryForDrawProgramTypeDelegateProgressController]
//   - [GTMioTraceTimelineData.ExecutionHistoryForPipelineStateDelegateProgressControllerCliqueFilter]
//   - [GTMioTraceTimelineData.ExecutionHistoryForPipelineStateProgramTypeDelegateProgressController]
//   - [GTMioTraceTimelineData.FailedUSCIndexCount]
//   - [GTMioTraceTimelineData.FailedUSCIndexes]
//   - [GTMioTraceTimelineData.FragmentPositionCount]
//   - [GTMioTraceTimelineData.FragmentPositions]
//   - [GTMioTraceTimelineData.GlobalGPUTime]
//   - [GTMioTraceTimelineData.GpuCost]
//   - [GTMioTraceTimelineData.GpuInfo]
//   - [GTMioTraceTimelineData.GpuTime]
//   - [GTMioTraceTimelineData.InitializeFromDatabase]
//   - [GTMioTraceTimelineData.InstructionCountForScopeScopeIdentifierDataMaster]
//   - [GTMioTraceTimelineData.KickDurationForEncoder]
//   - [GTMioTraceTimelineData.KickDurationForEncoderDataMaster]
//   - [GTMioTraceTimelineData.Kicks]
//   - [GTMioTraceTimelineData.KicksCount]
//   - [GTMioTraceTimelineData.MGPUs]
//   - [GTMioTraceTimelineData.MetalFXInfo]
//   - [GTMioTraceTimelineData.MetalFXInfoCount]
//   - [GTMioTraceTimelineData.NumDrawsForEncoder]
//   - [GTMioTraceTimelineData.NumDrawsForPipelineState]
//   - [GTMioTraceTimelineData.ParentData]
//   - [GTMioTraceTimelineData.SetParentData]
//   - [GTMioTraceTimelineData.PipelineStateCount]
//   - [GTMioTraceTimelineData.PipelineStateIdForCliqueIndex]
//   - [GTMioTraceTimelineData.ProfiledState]
//   - [GTMioTraceTimelineData.ProfiledWithOverlapEnabled]
//   - [GTMioTraceTimelineData.ReferenceComputePositionForClique]
//   - [GTMioTraceTimelineData.RiaTraceCount]
//   - [GTMioTraceTimelineData.RiaTraces]
//   - [GTMioTraceTimelineData.SampledCores]
//   - [GTMioTraceTimelineData.ShaderBinaryInfo]
//   - [GTMioTraceTimelineData.ShaderBinaryInfoCount]
//   - [GTMioTraceTimelineData.SignpostPipelineStateCount]
//   - [GTMioTraceTimelineData.SignpostPipelineStates]
//   - [GTMioTraceTimelineData.SignpostProcessCount]
//   - [GTMioTraceTimelineData.SignpostProcesses]
//   - [GTMioTraceTimelineData.SignpostShaderCount]
//   - [GTMioTraceTimelineData.SignpostShaders]
//   - [GTMioTraceTimelineData.SignpostStrings]
//   - [GTMioTraceTimelineData.StreamData]
//   - [GTMioTraceTimelineData.TimelineCounters]
//   - [GTMioTraceTimelineData.TimelineDuration]
//   - [GTMioTraceTimelineData.TimestampBegin]
//   - [GTMioTraceTimelineData.TimestampEnd]
//   - [GTMioTraceTimelineData.TotalCliqueCost]
//   - [GTMioTraceTimelineData.TotalCores]
//   - [GTMioTraceTimelineData.TotalCostForScopeScopeIdentifierDataMaster]
//   - [GTMioTraceTimelineData.TotalCostForScopeScopeIdentifierProgramType]
//   - [GTMioTraceTimelineData.Uscs]
//   - [GTMioTraceTimelineData.InitWithAPSTraceDataTimelineDataStreamDataTimelineTypeOptionsParentData]
//   - [GTMioTraceTimelineData.InitWithCoder]
//   - [GTMioTraceTimelineData.InitWithDecodedDictionaryStreamDataParentData]
//   - [GTMioTraceTimelineData.InitWithSerializedDataStreamDataParentData]
//   - [GTMioTraceTimelineData.InitWithTraceDatabaseDeallocatorParentData]
//   - [GTMioTraceTimelineData.DebugDescription]
//   - [GTMioTraceTimelineData.Description]
//   - [GTMioTraceTimelineData.Hash]
//   - [GTMioTraceTimelineData.Superclass]
type GTMioTraceTimelineData struct {
	objectivec.Object
}

// GTMioTraceTimelineDataFromID constructs a [GTMioTraceTimelineData] from an objc.ID.
func GTMioTraceTimelineDataFromID(id objc.ID) GTMioTraceTimelineData {
	return GTMioTraceTimelineData{objectivec.Object{ID: id}}
}

// Ensure GTMioTraceTimelineData implements IGTMioTraceTimelineData.
var _ IGTMioTraceTimelineData = GTMioTraceTimelineData{}

// An interface definition for the [GTMioTraceTimelineData] class.
//
// # Methods
//
//   - [IGTMioTraceTimelineData._cacheExeuctionHistory]
//   - [IGTMioTraceTimelineData._waitPendingExecutionHistory]
//   - [IGTMioTraceTimelineData.ArchivedData]
//   - [IGTMioTraceTimelineData.ArchivedDataWithCompressionOriginalSizeError]
//   - [IGTMioTraceTimelineData.Binaries]
//   - [IGTMioTraceTimelineData.BinaryForDrawProgramType]
//   - [IGTMioTraceTimelineData.BinaryForPipelineStateProgramType]
//   - [IGTMioTraceTimelineData.ChildCliqueOfClique]
//   - [IGTMioTraceTimelineData.CliqueFromCliqueIndex]
//   - [IGTMioTraceTimelineData.CoalescedFunctionIndexForEncoderFunctionIndex]
//   - [IGTMioTraceTimelineData.ComputePositionCount]
//   - [IGTMioTraceTimelineData.ComputePositions]
//   - [IGTMioTraceTimelineData.ConsistentStateAchieved]
//   - [IGTMioTraceTimelineData.CostCount]
//   - [IGTMioTraceTimelineData.CostForContextCost]
//   - [IGTMioTraceTimelineData.CostForLevelLevelIdentifierScopeScopeIdentifierCost]
//   - [IGTMioTraceTimelineData.CostForScopeScopeIdentifierCost]
//   - [IGTMioTraceTimelineData.Costs]
//   - [IGTMioTraceTimelineData.DataType]
//   - [IGTMioTraceTimelineData.DatabaseInternal]
//   - [IGTMioTraceTimelineData.DrawCount]
//   - [IGTMioTraceTimelineData.DrawTraceCount]
//   - [IGTMioTraceTimelineData.DrawTraces]
//   - [IGTMioTraceTimelineData.Draws]
//   - [IGTMioTraceTimelineData.DurationForDrawDataMaster]
//   - [IGTMioTraceTimelineData.EncodeError]
//   - [IGTMioTraceTimelineData.EncodeWithCoder]
//   - [IGTMioTraceTimelineData.EncoderCount]
//   - [IGTMioTraceTimelineData.EncoderFromFunctionIndex]
//   - [IGTMioTraceTimelineData.Encoders]
//   - [IGTMioTraceTimelineData.EnumerateBinariesForDrawEnumerator]
//   - [IGTMioTraceTimelineData.EnumerateBinariesForEncoderEnumerator]
//   - [IGTMioTraceTimelineData.EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator]
//   - [IGTMioTraceTimelineData.EnumerateBinariesForPipelineStateEnumerator]
//   - [IGTMioTraceTimelineData.EnumerateBinaryRangesForCliqueUscDataEnumerator]
//   - [IGTMioTraceTimelineData.EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator]
//   - [IGTMioTraceTimelineData.EnumerateDrawsForEncoderEnumerator]
//   - [IGTMioTraceTimelineData.EnumerateDrawsForPipelineStateEnumerator]
//   - [IGTMioTraceTimelineData.EnumerateEncoders]
//   - [IGTMioTraceTimelineData.EnumerateInstructionsForCliqueUscDataEnumerator]
//   - [IGTMioTraceTimelineData.EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator]
//   - [IGTMioTraceTimelineData.EnumerateKickAtFunctionIndexEnumerator]
//   - [IGTMioTraceTimelineData.EnumeratePipelineStates]
//   - [IGTMioTraceTimelineData.EnumerateUniqueTracesForBinaryEnumerator]
//   - [IGTMioTraceTimelineData.ExecutionHistoryForCliqueUscDelegate]
//   - [IGTMioTraceTimelineData.ExecutionHistoryForCliqueUscDelegateRequiresTimestampCount]
//   - [IGTMioTraceTimelineData.ExecutionHistoryForCliqueUscIndexDelegate]
//   - [IGTMioTraceTimelineData.ExecutionHistoryForDrawProgramTypeDelegateProgressController]
//   - [IGTMioTraceTimelineData.ExecutionHistoryForPipelineStateDelegateProgressControllerCliqueFilter]
//   - [IGTMioTraceTimelineData.ExecutionHistoryForPipelineStateProgramTypeDelegateProgressController]
//   - [IGTMioTraceTimelineData.FailedUSCIndexCount]
//   - [IGTMioTraceTimelineData.FailedUSCIndexes]
//   - [IGTMioTraceTimelineData.FragmentPositionCount]
//   - [IGTMioTraceTimelineData.FragmentPositions]
//   - [IGTMioTraceTimelineData.GlobalGPUTime]
//   - [IGTMioTraceTimelineData.GpuCost]
//   - [IGTMioTraceTimelineData.GpuInfo]
//   - [IGTMioTraceTimelineData.GpuTime]
//   - [IGTMioTraceTimelineData.InitializeFromDatabase]
//   - [IGTMioTraceTimelineData.InstructionCountForScopeScopeIdentifierDataMaster]
//   - [IGTMioTraceTimelineData.KickDurationForEncoder]
//   - [IGTMioTraceTimelineData.KickDurationForEncoderDataMaster]
//   - [IGTMioTraceTimelineData.Kicks]
//   - [IGTMioTraceTimelineData.KicksCount]
//   - [IGTMioTraceTimelineData.MGPUs]
//   - [IGTMioTraceTimelineData.MetalFXInfo]
//   - [IGTMioTraceTimelineData.MetalFXInfoCount]
//   - [IGTMioTraceTimelineData.NumDrawsForEncoder]
//   - [IGTMioTraceTimelineData.NumDrawsForPipelineState]
//   - [IGTMioTraceTimelineData.ParentData]
//   - [IGTMioTraceTimelineData.SetParentData]
//   - [IGTMioTraceTimelineData.PipelineStateCount]
//   - [IGTMioTraceTimelineData.PipelineStateIdForCliqueIndex]
//   - [IGTMioTraceTimelineData.ProfiledState]
//   - [IGTMioTraceTimelineData.ProfiledWithOverlapEnabled]
//   - [IGTMioTraceTimelineData.ReferenceComputePositionForClique]
//   - [IGTMioTraceTimelineData.RiaTraceCount]
//   - [IGTMioTraceTimelineData.RiaTraces]
//   - [IGTMioTraceTimelineData.SampledCores]
//   - [IGTMioTraceTimelineData.ShaderBinaryInfo]
//   - [IGTMioTraceTimelineData.ShaderBinaryInfoCount]
//   - [IGTMioTraceTimelineData.SignpostPipelineStateCount]
//   - [IGTMioTraceTimelineData.SignpostPipelineStates]
//   - [IGTMioTraceTimelineData.SignpostProcessCount]
//   - [IGTMioTraceTimelineData.SignpostProcesses]
//   - [IGTMioTraceTimelineData.SignpostShaderCount]
//   - [IGTMioTraceTimelineData.SignpostShaders]
//   - [IGTMioTraceTimelineData.SignpostStrings]
//   - [IGTMioTraceTimelineData.StreamData]
//   - [IGTMioTraceTimelineData.TimelineCounters]
//   - [IGTMioTraceTimelineData.TimelineDuration]
//   - [IGTMioTraceTimelineData.TimestampBegin]
//   - [IGTMioTraceTimelineData.TimestampEnd]
//   - [IGTMioTraceTimelineData.TotalCliqueCost]
//   - [IGTMioTraceTimelineData.TotalCores]
//   - [IGTMioTraceTimelineData.TotalCostForScopeScopeIdentifierDataMaster]
//   - [IGTMioTraceTimelineData.TotalCostForScopeScopeIdentifierProgramType]
//   - [IGTMioTraceTimelineData.Uscs]
//   - [IGTMioTraceTimelineData.InitWithAPSTraceDataTimelineDataStreamDataTimelineTypeOptionsParentData]
//   - [IGTMioTraceTimelineData.InitWithCoder]
//   - [IGTMioTraceTimelineData.InitWithDecodedDictionaryStreamDataParentData]
//   - [IGTMioTraceTimelineData.InitWithSerializedDataStreamDataParentData]
//   - [IGTMioTraceTimelineData.InitWithTraceDatabaseDeallocatorParentData]
//   - [IGTMioTraceTimelineData.DebugDescription]
//   - [IGTMioTraceTimelineData.Description]
//   - [IGTMioTraceTimelineData.Hash]
//   - [IGTMioTraceTimelineData.Superclass]
type IGTMioTraceTimelineData interface {
	objectivec.IObject

	// Topic: Methods

	_cacheExeuctionHistory(history objectivec.IObject)
	_waitPendingExecutionHistory(history objectivec.IObject) objectivec.IObject
	ArchivedData(data []objectivec.IObject) objectivec.IObject
	ArchivedDataWithCompressionOriginalSizeError(compression int64, size *uint64) (objectivec.IObject, error)
	Binaries() foundation.INSArray
	BinaryForDrawProgramType(draw uint32, type_ uint16) objectivec.IObject
	BinaryForPipelineStateProgramType(state uint64, type_ uint16) objectivec.IObject
	ChildCliqueOfClique(clique *GTMioUSCCliqueMetadata) *GTMioUSCCliqueMetadata
	CliqueFromCliqueIndex(index *GTMioUSCCliqueIndex) *GTMioUSCCliqueMetadata
	CoalescedFunctionIndexForEncoderFunctionIndex(index uint32) uint32
	ComputePositionCount() uint64
	ComputePositions() *GTMioCliqueComputePosition
	ConsistentStateAchieved() bool
	CostCount() uint64
	CostForContextCost(context *GTMioCostContext, cost *GTMioCostInfo) bool
	CostForLevelLevelIdentifierScopeScopeIdentifierCost(level uint16, identifier uint32, scope uint16, identifier2 uint64, cost *GTMioCostInfo) bool
	CostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost *GTMioCostInfo) bool
	Costs() *GTMioCostInfo
	DataType() uint32
	DatabaseInternal() uint64
	DrawCount() uint64
	DrawTraceCount() uint64
	DrawTraces() *GTMioDrawTrace
	Draws() *GTMioDrawMetadata
	DurationForDrawDataMaster(draw uint32, master uint16) uint64
	EncodeError(encode bool) (objectivec.IObject, error)
	EncodeWithCoder(coder foundation.INSCoder)
	EncoderCount() uint64
	EncoderFromFunctionIndex(index uint32) *GTMioEncoderMetadata
	Encoders() *GTMioEncoderMetadata
	EnumerateBinariesForDrawEnumerator(draw uint32, enumerator VoidHandler)
	EnumerateBinariesForEncoderEnumerator(encoder uint32, enumerator VoidHandler)
	EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator VoidHandler)
	EnumerateBinariesForPipelineStateEnumerator(state uint64, enumerator VoidHandler)
	EnumerateBinaryRangesForCliqueUscDataEnumerator(clique *GTMioUSCCliqueMetadata, data objectivec.IObject, enumerator VoidHandler)
	EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator VoidHandler)
	EnumerateDrawsForEncoderEnumerator(encoder uint32, enumerator VoidHandler)
	EnumerateDrawsForPipelineStateEnumerator(state uint64, enumerator VoidHandler)
	EnumerateEncoders(encoders VoidHandler)
	EnumerateInstructionsForCliqueUscDataEnumerator(clique *GTMioUSCCliqueMetadata, data objectivec.IObject, enumerator VoidHandler)
	EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator VoidHandler)
	EnumerateKickAtFunctionIndexEnumerator(index uint32, enumerator VoidHandler)
	EnumeratePipelineStates(states VoidHandler)
	EnumerateUniqueTracesForBinaryEnumerator(binary uint32, enumerator VoidHandler)
	ExecutionHistoryForCliqueUscDelegate(clique *GTMioUSCCliqueMetadata, usc unsafe.Pointer, delegate objectivec.IObject)
	ExecutionHistoryForCliqueUscDelegateRequiresTimestampCount(clique *GTMioUSCCliqueMetadata, usc unsafe.Pointer, delegate objectivec.IObject, timestamp bool, count uint32)
	ExecutionHistoryForCliqueUscIndexDelegate(clique uint32, index uint32, delegate objectivec.IObject)
	ExecutionHistoryForDrawProgramTypeDelegateProgressController(draw uint32, type_ uint16, delegate objectivec.IObject, controller objectivec.IObject)
	ExecutionHistoryForPipelineStateDelegateProgressControllerCliqueFilter(state uint64, delegate objectivec.IObject, controller objectivec.IObject, filter VoidHandler)
	ExecutionHistoryForPipelineStateProgramTypeDelegateProgressController(state uint64, type_ uint16, delegate objectivec.IObject, controller objectivec.IObject)
	FailedUSCIndexCount() uint64
	FailedUSCIndexes() unsafe.Pointer
	FragmentPositionCount() uint64
	FragmentPositions() *GTMioCliqueFragmentPosition
	GlobalGPUTime() uint64
	GpuCost() *GTMioCostInfo
	GpuInfo() IGTMioGPUInfo
	GpuTime() uint64
	InitializeFromDatabase() GTMioTraceTimelineData
	InstructionCountForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) uint64
	KickDurationForEncoder(encoder uint32) uint64
	KickDurationForEncoderDataMaster(encoder uint32, master uint16) uint64
	Kicks() *GTMioKickTrace
	KicksCount() uint64
	MGPUs() foundation.INSArray
	MetalFXInfo() *GTMioMetalFXInfo
	MetalFXInfoCount() uint64
	NumDrawsForEncoder(encoder uint32) uint64
	NumDrawsForPipelineState(state uint64) uint64
	ParentData() IGTMioTraceData
	SetParentData(value IGTMioTraceData)
	PipelineStateCount() uint64
	PipelineStateIdForCliqueIndex(index *GTMioUSCCliqueIndex) uint64
	ProfiledState() uint32
	ProfiledWithOverlapEnabled() bool
	ReferenceComputePositionForClique(clique *GTMioUSCCliqueMetadata) *GTMioCliqueComputePosition
	RiaTraceCount() uint64
	RiaTraces() *GTMioRIATrace
	SampledCores() uint32
	ShaderBinaryInfo() *GTMioShaderBinaryInfo
	ShaderBinaryInfoCount() uint64
	SignpostPipelineStateCount() uint64
	SignpostPipelineStates() *GTMioShaderTimelineSignpostPipelineState
	SignpostProcessCount() uint64
	SignpostProcesses() *GTMioShaderTimelineSignpostProcess
	SignpostShaderCount() uint64
	SignpostShaders() *GTMioShaderTimelineSignpostShader
	SignpostStrings() foundation.INSArray
	StreamData() IGTShaderProfilerStreamData
	TimelineCounters() IGTMioTimelineCounters
	TimelineDuration() uint64
	TimestampBegin() uint64
	TimestampEnd() uint64
	TotalCliqueCost() uint64
	TotalCores() uint32
	TotalCostForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) float64
	TotalCostForScopeScopeIdentifierProgramType(scope uint16, identifier uint64, type_ uint16) float64
	Uscs() foundation.INSArray
	InitWithAPSTraceDataTimelineDataStreamDataTimelineTypeOptionsParentData(data unsafe.Pointer, data2 unsafe.Pointer, data3 objectivec.IObject, type_ uint32, options GTMioTraceDataBuilderOptions, data4 objectivec.IObject) GTMioTraceTimelineData
	InitWithCoder(coder foundation.INSCoder) GTMioTraceTimelineData
	InitWithDecodedDictionaryStreamDataParentData(dictionary objectivec.IObject, data objectivec.IObject, data2 objectivec.IObject) GTMioTraceTimelineData
	InitWithSerializedDataStreamDataParentData(data objectivec.IObject, data2 objectivec.IObject, data3 objectivec.IObject) GTMioTraceTimelineData
	InitWithTraceDatabaseDeallocatorParentData(database uint64, deallocator VoidHandler, data objectivec.IObject) GTMioTraceTimelineData
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (g GTMioTraceTimelineData) Init() GTMioTraceTimelineData {
	rv := objc.SendIfResponds[GTMioTraceTimelineData](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioTraceTimelineData) Autorelease() GTMioTraceTimelineData {
	rv := objc.SendIfResponds[GTMioTraceTimelineData](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioTraceTimelineData creates a new GTMioTraceTimelineData instance.
func NewGTMioTraceTimelineData() GTMioTraceTimelineData {
	class := getGTMioTraceTimelineDataClass()
	rv := objc.SendIfResponds[GTMioTraceTimelineData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTMioTraceTimelineDataWithAPSTraceDataTimelineDataStreamDataTimelineTypeOptionsParentData(data unsafe.Pointer, data2 unsafe.Pointer, data3 objectivec.IObject, type_ uint32, options GTMioTraceDataBuilderOptions, data4 objectivec.IObject) GTMioTraceTimelineData {
	instance := getGTMioTraceTimelineDataClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithAPSTraceData:timelineData:streamData:timelineType:options:parentData:"), data, data2, data3, type_, options, data4)
	return GTMioTraceTimelineDataFromID(rv)
}

func NewGTMioTraceTimelineDataWithCoder(coder objectivec.IObject) GTMioTraceTimelineData {
	instance := getGTMioTraceTimelineDataClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTMioTraceTimelineDataFromID(rv)
}

func NewGTMioTraceTimelineDataWithDecodedDictionaryStreamDataParentData(dictionary objectivec.IObject, data objectivec.IObject, data2 objectivec.IObject) GTMioTraceTimelineData {
	instance := getGTMioTraceTimelineDataClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDecodedDictionary:streamData:parentData:"), dictionary, data, data2)
	return GTMioTraceTimelineDataFromID(rv)
}

func NewGTMioTraceTimelineDataWithSerializedDataStreamDataParentData(data objectivec.IObject, data2 objectivec.IObject, data3 objectivec.IObject) GTMioTraceTimelineData {
	instance := getGTMioTraceTimelineDataClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSerializedData:streamData:parentData:"), data, data2, data3)
	return GTMioTraceTimelineDataFromID(rv)
}

func (g GTMioTraceTimelineData) _cacheExeuctionHistory(history objectivec.IObject) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("_cacheExeuctionHistory:"), history)
}

// CacheExeuctionHistory is an exported wrapper for the private method _cacheExeuctionHistory.
func (g GTMioTraceTimelineData) CacheExeuctionHistory(history objectivec.IObject) error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_cacheExeuctionHistory:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_cacheExeuctionHistory:"}
		return err
	}
	g._cacheExeuctionHistory(history)
	return nil
}

// CanCacheExeuctionHistory reports whether the receiver responds to the private selector _cacheExeuctionHistory:.
func (g GTMioTraceTimelineData) CanCacheExeuctionHistory() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_cacheExeuctionHistory:"))
}
func (g GTMioTraceTimelineData) _waitPendingExecutionHistory(history objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("_waitPendingExecutionHistory:"), history)
	return objectivec.Object{ID: rv}
}

// WaitPendingExecutionHistory is an exported wrapper for the private method _waitPendingExecutionHistory.
func (g GTMioTraceTimelineData) WaitPendingExecutionHistory(history objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_waitPendingExecutionHistory:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_waitPendingExecutionHistory:"}
		return nil, err
	}
	return g._waitPendingExecutionHistory(history), nil
}

// CanWaitPendingExecutionHistory reports whether the receiver responds to the private selector _waitPendingExecutionHistory:.
func (g GTMioTraceTimelineData) CanWaitPendingExecutionHistory() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_waitPendingExecutionHistory:"))
}
func (g GTMioTraceTimelineData) ArchivedData(data []objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("archivedData:"), objectivec.IObjectSliceToNSArray(data))
	return objectivec.Object{ID: rv}
}
func (g GTMioTraceTimelineData) ArchivedDataWithCompressionOriginalSizeError(compression int64, size *uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](g.ID, objc.Sel("archivedDataWithCompression:originalSize:error:"), compression, unsafe.Pointer(size), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (g GTMioTraceTimelineData) BinaryForDrawProgramType(draw uint32, type_ uint16) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("binaryForDraw:programType:"), draw, type_)
	return objectivec.Object{ID: rv}
}
func (g GTMioTraceTimelineData) BinaryForPipelineStateProgramType(state uint64, type_ uint16) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("binaryForPipelineState:programType:"), state, type_)
	return objectivec.Object{ID: rv}
}
func (g GTMioTraceTimelineData) ChildCliqueOfClique(clique *GTMioUSCCliqueMetadata) *GTMioUSCCliqueMetadata {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("childCliqueOfClique:"), unsafe.Pointer(clique))
	return (*GTMioUSCCliqueMetadata)(rv)
}
func (g GTMioTraceTimelineData) CliqueFromCliqueIndex(index *GTMioUSCCliqueIndex) *GTMioUSCCliqueMetadata {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("cliqueFromCliqueIndex:"), unsafe.Pointer(index))
	return (*GTMioUSCCliqueMetadata)(rv)
}
func (g GTMioTraceTimelineData) CoalescedFunctionIndexForEncoderFunctionIndex(index uint32) uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("coalescedFunctionIndexForEncoderFunctionIndex:"), index)
	return rv
}
func (g GTMioTraceTimelineData) CostForContextCost(context *GTMioCostContext, cost *GTMioCostInfo) bool {
	rv := objc.SendIfResponds[bool](g.ID, objc.Sel("costForContext:cost:"), unsafe.Pointer(context), unsafe.Pointer(cost))
	return rv
}
func (g GTMioTraceTimelineData) CostForLevelLevelIdentifierScopeScopeIdentifierCost(level uint16, identifier uint32, scope uint16, identifier2 uint64, cost *GTMioCostInfo) bool {
	rv := objc.SendIfResponds[bool](g.ID, objc.Sel("costForLevel:levelIdentifier:scope:scopeIdentifier:cost:"), level, identifier, scope, identifier2, unsafe.Pointer(cost))
	return rv
}
func (g GTMioTraceTimelineData) CostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost *GTMioCostInfo) bool {
	rv := objc.SendIfResponds[bool](g.ID, objc.Sel("costForScope:scopeIdentifier:cost:"), scope, identifier, unsafe.Pointer(cost))
	return rv
}
func (g GTMioTraceTimelineData) DurationForDrawDataMaster(draw uint32, master uint16) uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("durationForDraw:dataMaster:"), draw, master)
	return rv
}
func (g GTMioTraceTimelineData) EncodeError(encode bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](g.ID, objc.Sel("encode:error:"), encode, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (g GTMioTraceTimelineData) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (g GTMioTraceTimelineData) EncoderFromFunctionIndex(index uint32) *GTMioEncoderMetadata {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("encoderFromFunctionIndex:"), index)
	return (*GTMioEncoderMetadata)(rv)
}
func (g GTMioTraceTimelineData) EnumerateBinariesForDrawEnumerator(draw uint32, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("enumerateBinariesForDraw:enumerator:"), draw, _block1)
}
func (g GTMioTraceTimelineData) EnumerateBinariesForEncoderEnumerator(encoder uint32, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("enumerateBinariesForEncoder:enumerator:"), encoder, _block1)
}
func (g GTMioTraceTimelineData) EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator VoidHandler) {
	_block2, _ := NewVoidBlock(enumerator)
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("enumerateBinariesForForCliqueAtIndex:uscIndex:enumerator:"), index, index2, _block2)
}
func (g GTMioTraceTimelineData) EnumerateBinariesForPipelineStateEnumerator(state uint64, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("enumerateBinariesForPipelineState:enumerator:"), state, _block1)
}
func (g GTMioTraceTimelineData) EnumerateBinaryRangesForCliqueUscDataEnumerator(clique *GTMioUSCCliqueMetadata, data objectivec.IObject, enumerator VoidHandler) {
	_block2, _ := NewVoidBlock(enumerator)
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("enumerateBinaryRangesForClique:uscData:enumerator:"), clique, data, _block2)
}
func (g GTMioTraceTimelineData) EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator VoidHandler) {
	_block2, _ := NewVoidBlock(enumerator)
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("enumerateBinaryRangesForCliqueAtIndex:uscIndex:enumerator:"), index, index2, _block2)
}
func (g GTMioTraceTimelineData) EnumerateDrawsForEncoderEnumerator(encoder uint32, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("enumerateDrawsForEncoder:enumerator:"), encoder, _block1)
}
func (g GTMioTraceTimelineData) EnumerateDrawsForPipelineStateEnumerator(state uint64, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("enumerateDrawsForPipelineState:enumerator:"), state, _block1)
}
func (g GTMioTraceTimelineData) EnumerateEncoders(encoders VoidHandler) {
	_block0, _ := NewVoidBlock(encoders)
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("enumerateEncoders:"), _block0)
}
func (g GTMioTraceTimelineData) EnumerateInstructionsForCliqueUscDataEnumerator(clique *GTMioUSCCliqueMetadata, data objectivec.IObject, enumerator VoidHandler) {
	_block2, _ := NewVoidBlock(enumerator)
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("enumerateInstructionsForClique:uscData:enumerator:"), clique, data, _block2)
}
func (g GTMioTraceTimelineData) EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator VoidHandler) {
	_block2, _ := NewVoidBlock(enumerator)
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("enumerateInstructionsForCliqueAtIndex:uscIndex:enumerator:"), index, index2, _block2)
}
func (g GTMioTraceTimelineData) EnumerateKickAtFunctionIndexEnumerator(index uint32, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("enumerateKickAtFunctionIndex:enumerator:"), index, _block1)
}
func (g GTMioTraceTimelineData) EnumeratePipelineStates(states VoidHandler) {
	_block0, _ := NewVoidBlock(states)
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("enumeratePipelineStates:"), _block0)
}
func (g GTMioTraceTimelineData) EnumerateUniqueTracesForBinaryEnumerator(binary uint32, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("enumerateUniqueTracesForBinary:enumerator:"), binary, _block1)
}
func (g GTMioTraceTimelineData) ExecutionHistoryForCliqueUscDelegate(clique *GTMioUSCCliqueMetadata, usc unsafe.Pointer, delegate objectivec.IObject) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("executionHistoryForClique:usc:delegate:"), unsafe.Pointer(clique), usc, delegate)
}
func (g GTMioTraceTimelineData) ExecutionHistoryForCliqueUscDelegateRequiresTimestampCount(clique *GTMioUSCCliqueMetadata, usc unsafe.Pointer, delegate objectivec.IObject, timestamp bool, count uint32) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("executionHistoryForClique:usc:delegate:requiresTimestamp:count:"), objc.CArray(clique), objc.CArray(usc), delegate, timestamp, count)
}
func (g GTMioTraceTimelineData) ExecutionHistoryForCliqueUscIndexDelegate(clique uint32, index uint32, delegate objectivec.IObject) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("executionHistoryForClique:uscIndex:delegate:"), clique, index, delegate)
}
func (g GTMioTraceTimelineData) ExecutionHistoryForDrawProgramTypeDelegateProgressController(draw uint32, type_ uint16, delegate objectivec.IObject, controller objectivec.IObject) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("executionHistoryForDraw:programType:delegate:progressController:"), draw, type_, delegate, controller)
}
func (g GTMioTraceTimelineData) ExecutionHistoryForPipelineStateDelegateProgressControllerCliqueFilter(state uint64, delegate objectivec.IObject, controller objectivec.IObject, filter VoidHandler) {
	_block3, _ := NewVoidBlock(filter)
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("executionHistoryForPipelineState:delegate:progressController:cliqueFilter:"), state, delegate, controller, _block3)
}
func (g GTMioTraceTimelineData) ExecutionHistoryForPipelineStateProgramTypeDelegateProgressController(state uint64, type_ uint16, delegate objectivec.IObject, controller objectivec.IObject) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("executionHistoryForPipelineState:programType:delegate:progressController:"), state, type_, delegate, controller)
}
func (g GTMioTraceTimelineData) InitializeFromDatabase() GTMioTraceTimelineData {
	rv := objc.SendIfResponds[GTMioTraceTimelineData](g.ID, objc.Sel("initializeFromDatabase"))
	return rv
}
func (g GTMioTraceTimelineData) InstructionCountForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("instructionCountForScope:scopeIdentifier:dataMaster:"), scope, identifier, master)
	return rv
}
func (g GTMioTraceTimelineData) KickDurationForEncoder(encoder uint32) uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("kickDurationForEncoder:"), encoder)
	return rv
}
func (g GTMioTraceTimelineData) KickDurationForEncoderDataMaster(encoder uint32, master uint16) uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("kickDurationForEncoder:dataMaster:"), encoder, master)
	return rv
}
func (g GTMioTraceTimelineData) NumDrawsForEncoder(encoder uint32) uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("numDrawsForEncoder:"), encoder)
	return rv
}
func (g GTMioTraceTimelineData) NumDrawsForPipelineState(state uint64) uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("numDrawsForPipelineState:"), state)
	return rv
}
func (g GTMioTraceTimelineData) PipelineStateIdForCliqueIndex(index *GTMioUSCCliqueIndex) uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("pipelineStateIdForCliqueIndex:"), unsafe.Pointer(index))
	return rv
}
func (g GTMioTraceTimelineData) ReferenceComputePositionForClique(clique *GTMioUSCCliqueMetadata) *GTMioCliqueComputePosition {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("referenceComputePositionForClique:"), unsafe.Pointer(clique))
	return (*GTMioCliqueComputePosition)(rv)
}
func (g GTMioTraceTimelineData) TotalCostForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) float64 {
	rv := objc.SendIfResponds[float64](g.ID, objc.Sel("totalCostForScope:scopeIdentifier:dataMaster:"), scope, identifier, master)
	return rv
}
func (g GTMioTraceTimelineData) TotalCostForScopeScopeIdentifierProgramType(scope uint16, identifier uint64, type_ uint16) float64 {
	rv := objc.SendIfResponds[float64](g.ID, objc.Sel("totalCostForScope:scopeIdentifier:programType:"), scope, identifier, type_)
	return rv
}
func (g GTMioTraceTimelineData) InitWithAPSTraceDataTimelineDataStreamDataTimelineTypeOptionsParentData(data unsafe.Pointer, data2 unsafe.Pointer, data3 objectivec.IObject, type_ uint32, options GTMioTraceDataBuilderOptions, data4 objectivec.IObject) GTMioTraceTimelineData {
	rv := objc.SendIfResponds[GTMioTraceTimelineData](g.ID, objc.Sel("initWithAPSTraceData:timelineData:streamData:timelineType:options:parentData:"), data, data2, data3, type_, options, data4)
	return rv
}
func (g GTMioTraceTimelineData) InitWithCoder(coder foundation.INSCoder) GTMioTraceTimelineData {
	rv := objc.SendIfResponds[GTMioTraceTimelineData](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (g GTMioTraceTimelineData) InitWithDecodedDictionaryStreamDataParentData(dictionary objectivec.IObject, data objectivec.IObject, data2 objectivec.IObject) GTMioTraceTimelineData {
	rv := objc.SendIfResponds[GTMioTraceTimelineData](g.ID, objc.Sel("initWithDecodedDictionary:streamData:parentData:"), dictionary, data, data2)
	return rv
}
func (g GTMioTraceTimelineData) InitWithSerializedDataStreamDataParentData(data objectivec.IObject, data2 objectivec.IObject, data3 objectivec.IObject) GTMioTraceTimelineData {
	rv := objc.SendIfResponds[GTMioTraceTimelineData](g.ID, objc.Sel("initWithSerializedData:streamData:parentData:"), data, data2, data3)
	return rv
}

var _gtmiotracetimelinedata_initwithtracedatabase_deallocator_parentdata_p1_key byte

func (g GTMioTraceTimelineData) InitWithTraceDatabaseDeallocatorParentData(database uint64, deallocator VoidHandler, data objectivec.IObject) GTMioTraceTimelineData {
	_block1, _ := NewVoidBlock(deallocator)
	rv := objc.SendIfResponds[GTMioTraceTimelineData](g.ID, objc.Sel("initWithTraceDatabase:deallocator:parentData:"), database, _block1, data)
	return rv
}

func (_GTMioTraceTimelineDataClass GTMioTraceTimelineDataClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_GTMioTraceTimelineDataClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (g GTMioTraceTimelineData) Binaries() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("binaries"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTMioTraceTimelineData) ComputePositionCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("computePositionCount"))
	return rv
}
func (g GTMioTraceTimelineData) ComputePositions() *GTMioCliqueComputePosition {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("computePositions"))
	return (*GTMioCliqueComputePosition)(rv)
}
func (g GTMioTraceTimelineData) ConsistentStateAchieved() bool {
	rv := objc.SendIfResponds[bool](g.ID, objc.Sel("consistentStateAchieved"))
	return rv
}
func (g GTMioTraceTimelineData) CostCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("costCount"))
	return rv
}
func (g GTMioTraceTimelineData) Costs() *GTMioCostInfo {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("costs"))
	return (*GTMioCostInfo)(rv)
}
func (g GTMioTraceTimelineData) DataType() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("dataType"))
	return rv
}
func (g GTMioTraceTimelineData) DatabaseInternal() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("databaseInternal"))
	return rv
}
func (g GTMioTraceTimelineData) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTMioTraceTimelineData) Description() string {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTMioTraceTimelineData) DrawCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("drawCount"))
	return rv
}
func (g GTMioTraceTimelineData) DrawTraceCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("drawTraceCount"))
	return rv
}
func (g GTMioTraceTimelineData) DrawTraces() *GTMioDrawTrace {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("drawTraces"))
	return (*GTMioDrawTrace)(rv)
}
func (g GTMioTraceTimelineData) Draws() *GTMioDrawMetadata {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("draws"))
	return (*GTMioDrawMetadata)(rv)
}
func (g GTMioTraceTimelineData) EncoderCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("encoderCount"))
	return rv
}
func (g GTMioTraceTimelineData) Encoders() *GTMioEncoderMetadata {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("encoders"))
	return (*GTMioEncoderMetadata)(rv)
}
func (g GTMioTraceTimelineData) FailedUSCIndexCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("failedUSCIndexCount"))
	return rv
}
func (g GTMioTraceTimelineData) FailedUSCIndexes() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("failedUSCIndexes"))
	return rv
}
func (g GTMioTraceTimelineData) FragmentPositionCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("fragmentPositionCount"))
	return rv
}
func (g GTMioTraceTimelineData) FragmentPositions() *GTMioCliqueFragmentPosition {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("fragmentPositions"))
	return (*GTMioCliqueFragmentPosition)(rv)
}
func (g GTMioTraceTimelineData) GlobalGPUTime() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("globalGPUTime"))
	return rv
}
func (g GTMioTraceTimelineData) GpuCost() *GTMioCostInfo {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("gpuCost"))
	return (*GTMioCostInfo)(rv)
}
func (g GTMioTraceTimelineData) GpuInfo() IGTMioGPUInfo {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("gpuInfo"))
	return GTMioGPUInfoFromID(objc.ID(rv))
}
func (g GTMioTraceTimelineData) GpuTime() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("gpuTime"))
	return rv
}
func (g GTMioTraceTimelineData) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("hash"))
	return rv
}
func (g GTMioTraceTimelineData) Kicks() *GTMioKickTrace {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("kicks"))
	return (*GTMioKickTrace)(rv)
}
func (g GTMioTraceTimelineData) KicksCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("kicksCount"))
	return rv
}
func (g GTMioTraceTimelineData) MGPUs() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("mGPUs"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTMioTraceTimelineData) MetalFXInfo() *GTMioMetalFXInfo {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("metalFXInfo"))
	return (*GTMioMetalFXInfo)(rv)
}
func (g GTMioTraceTimelineData) MetalFXInfoCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("metalFXInfoCount"))
	return rv
}
func (g GTMioTraceTimelineData) ParentData() IGTMioTraceData {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("parentData"))
	return GTMioTraceDataFromID(objc.ID(rv))
}
func (g GTMioTraceTimelineData) SetParentData(value IGTMioTraceData) {
	objc.SendIfResponds[struct{}](g.ID, objc.Sel("setParentData:"), value)
}
func (g GTMioTraceTimelineData) PipelineStateCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("pipelineStateCount"))
	return rv
}
func (g GTMioTraceTimelineData) ProfiledState() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("profiledState"))
	return rv
}
func (g GTMioTraceTimelineData) ProfiledWithOverlapEnabled() bool {
	rv := objc.SendIfResponds[bool](g.ID, objc.Sel("profiledWithOverlapEnabled"))
	return rv
}
func (g GTMioTraceTimelineData) RiaTraceCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("riaTraceCount"))
	return rv
}
func (g GTMioTraceTimelineData) RiaTraces() *GTMioRIATrace {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("riaTraces"))
	return (*GTMioRIATrace)(rv)
}
func (g GTMioTraceTimelineData) SampledCores() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("sampledCores"))
	return rv
}
func (g GTMioTraceTimelineData) ShaderBinaryInfo() *GTMioShaderBinaryInfo {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("shaderBinaryInfo"))
	return (*GTMioShaderBinaryInfo)(rv)
}
func (g GTMioTraceTimelineData) ShaderBinaryInfoCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("shaderBinaryInfoCount"))
	return rv
}
func (g GTMioTraceTimelineData) SignpostPipelineStateCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("signpostPipelineStateCount"))
	return rv
}
func (g GTMioTraceTimelineData) SignpostPipelineStates() *GTMioShaderTimelineSignpostPipelineState {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("signpostPipelineStates"))
	return (*GTMioShaderTimelineSignpostPipelineState)(rv)
}
func (g GTMioTraceTimelineData) SignpostProcessCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("signpostProcessCount"))
	return rv
}
func (g GTMioTraceTimelineData) SignpostProcesses() *GTMioShaderTimelineSignpostProcess {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("signpostProcesses"))
	return (*GTMioShaderTimelineSignpostProcess)(rv)
}
func (g GTMioTraceTimelineData) SignpostShaderCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("signpostShaderCount"))
	return rv
}
func (g GTMioTraceTimelineData) SignpostShaders() *GTMioShaderTimelineSignpostShader {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("signpostShaders"))
	return (*GTMioShaderTimelineSignpostShader)(rv)
}
func (g GTMioTraceTimelineData) SignpostStrings() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("signpostStrings"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTMioTraceTimelineData) StreamData() IGTShaderProfilerStreamData {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("streamData"))
	return GTShaderProfilerStreamDataFromID(objc.ID(rv))
}
func (g GTMioTraceTimelineData) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](g.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (g GTMioTraceTimelineData) TimelineCounters() IGTMioTimelineCounters {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("timelineCounters"))
	return GTMioTimelineCountersFromID(objc.ID(rv))
}
func (g GTMioTraceTimelineData) TimelineDuration() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("timelineDuration"))
	return rv
}
func (g GTMioTraceTimelineData) TimestampBegin() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("timestampBegin"))
	return rv
}
func (g GTMioTraceTimelineData) TimestampEnd() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("timestampEnd"))
	return rv
}
func (g GTMioTraceTimelineData) TotalCliqueCost() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("totalCliqueCost"))
	return rv
}
func (g GTMioTraceTimelineData) TotalCores() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("totalCores"))
	return rv
}
func (g GTMioTraceTimelineData) Uscs() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("uscs"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// EnumerateBinariesForDrawEnumeratorSync is a synchronous wrapper around [GTMioTraceTimelineData.EnumerateBinariesForDrawEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceTimelineData) EnumerateBinariesForDrawEnumeratorSync(ctx context.Context, draw uint32) error {
	done := make(chan struct{}, 1)
	g.EnumerateBinariesForDrawEnumerator(draw, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateBinariesForEncoderEnumeratorSync is a synchronous wrapper around [GTMioTraceTimelineData.EnumerateBinariesForEncoderEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceTimelineData) EnumerateBinariesForEncoderEnumeratorSync(ctx context.Context, encoder uint32) error {
	done := make(chan struct{}, 1)
	g.EnumerateBinariesForEncoderEnumerator(encoder, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateBinariesForForCliqueAtIndexUscIndexEnumeratorSync is a synchronous wrapper around [GTMioTraceTimelineData.EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceTimelineData) EnumerateBinariesForForCliqueAtIndexUscIndexEnumeratorSync(ctx context.Context, index uint32, index2 uint32) error {
	done := make(chan struct{}, 1)
	g.EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator(index, index2, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateBinariesForPipelineStateEnumeratorSync is a synchronous wrapper around [GTMioTraceTimelineData.EnumerateBinariesForPipelineStateEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceTimelineData) EnumerateBinariesForPipelineStateEnumeratorSync(ctx context.Context, state uint64) error {
	done := make(chan struct{}, 1)
	g.EnumerateBinariesForPipelineStateEnumerator(state, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateBinaryRangesForCliqueUscDataEnumeratorSync is a synchronous wrapper around [GTMioTraceTimelineData.EnumerateBinaryRangesForCliqueUscDataEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceTimelineData) EnumerateBinaryRangesForCliqueUscDataEnumeratorSync(ctx context.Context, clique *GTMioUSCCliqueMetadata, data objectivec.IObject) error {
	done := make(chan struct{}, 1)
	g.EnumerateBinaryRangesForCliqueUscDataEnumerator(clique, data, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumeratorSync is a synchronous wrapper around [GTMioTraceTimelineData.EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceTimelineData) EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumeratorSync(ctx context.Context, index uint32, index2 uint32) error {
	done := make(chan struct{}, 1)
	g.EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator(index, index2, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateDrawsForEncoderEnumeratorSync is a synchronous wrapper around [GTMioTraceTimelineData.EnumerateDrawsForEncoderEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceTimelineData) EnumerateDrawsForEncoderEnumeratorSync(ctx context.Context, encoder uint32) error {
	done := make(chan struct{}, 1)
	g.EnumerateDrawsForEncoderEnumerator(encoder, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateDrawsForPipelineStateEnumeratorSync is a synchronous wrapper around [GTMioTraceTimelineData.EnumerateDrawsForPipelineStateEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceTimelineData) EnumerateDrawsForPipelineStateEnumeratorSync(ctx context.Context, state uint64) error {
	done := make(chan struct{}, 1)
	g.EnumerateDrawsForPipelineStateEnumerator(state, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateEncodersSync is a synchronous wrapper around [GTMioTraceTimelineData.EnumerateEncoders].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceTimelineData) EnumerateEncodersSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.EnumerateEncoders(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateInstructionsForCliqueUscDataEnumeratorSync is a synchronous wrapper around [GTMioTraceTimelineData.EnumerateInstructionsForCliqueUscDataEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceTimelineData) EnumerateInstructionsForCliqueUscDataEnumeratorSync(ctx context.Context, clique *GTMioUSCCliqueMetadata, data objectivec.IObject) error {
	done := make(chan struct{}, 1)
	g.EnumerateInstructionsForCliqueUscDataEnumerator(clique, data, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateInstructionsForCliqueAtIndexUscIndexEnumeratorSync is a synchronous wrapper around [GTMioTraceTimelineData.EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceTimelineData) EnumerateInstructionsForCliqueAtIndexUscIndexEnumeratorSync(ctx context.Context, index uint32, index2 uint32) error {
	done := make(chan struct{}, 1)
	g.EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator(index, index2, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateKickAtFunctionIndexEnumeratorSync is a synchronous wrapper around [GTMioTraceTimelineData.EnumerateKickAtFunctionIndexEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceTimelineData) EnumerateKickAtFunctionIndexEnumeratorSync(ctx context.Context, index uint32) error {
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

// EnumeratePipelineStatesSync is a synchronous wrapper around [GTMioTraceTimelineData.EnumeratePipelineStates].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceTimelineData) EnumeratePipelineStatesSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.EnumeratePipelineStates(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateUniqueTracesForBinaryEnumeratorSync is a synchronous wrapper around [GTMioTraceTimelineData.EnumerateUniqueTracesForBinaryEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceTimelineData) EnumerateUniqueTracesForBinaryEnumeratorSync(ctx context.Context, binary uint32) error {
	done := make(chan struct{}, 1)
	g.EnumerateUniqueTracesForBinaryEnumerator(binary, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ExecutionHistoryForPipelineStateDelegateProgressControllerCliqueFilterSync is a synchronous wrapper around [GTMioTraceTimelineData.ExecutionHistoryForPipelineStateDelegateProgressControllerCliqueFilter].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceTimelineData) ExecutionHistoryForPipelineStateDelegateProgressControllerCliqueFilterSync(ctx context.Context, state uint64, delegate objectivec.IObject, controller objectivec.IObject) error {
	done := make(chan struct{}, 1)
	g.ExecutionHistoryForPipelineStateDelegateProgressControllerCliqueFilter(state, delegate, controller, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

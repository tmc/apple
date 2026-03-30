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
	rv := objc.Send[GTMioTraceTimelineData](objc.ID(gc.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData
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
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData
type IGTMioTraceTimelineData interface {
	objectivec.IObject

	// Topic: Methods

	_cacheExeuctionHistory(history objectivec.IObject)
	_waitPendingExecutionHistory(history objectivec.IObject) objectivec.IObject
	ArchivedData(data []objectivec.IObject) objectivec.IObject
	ArchivedDataWithCompressionOriginalSizeError(compression int64, size unsafe.Pointer) (objectivec.IObject, error)
	Binaries() foundation.INSArray
	BinaryForDrawProgramType(draw uint32, type_ uint16) objectivec.IObject
	BinaryForPipelineStateProgramType(state uint64, type_ uint16) objectivec.IObject
	ChildCliqueOfClique(clique unsafe.Pointer) unsafe.Pointer
	CliqueFromCliqueIndex(index unsafe.Pointer) unsafe.Pointer
	CoalescedFunctionIndexForEncoderFunctionIndex(index uint32) uint32
	ComputePositionCount() uint64
	ComputePositions() unsafe.Pointer
	ConsistentStateAchieved() bool
	CostCount() uint64
	CostForContextCost(context unsafe.Pointer, cost unsafe.Pointer) bool
	CostForLevelLevelIdentifierScopeScopeIdentifierCost(level uint16, identifier uint32, scope uint16, identifier2 uint64, cost unsafe.Pointer) bool
	CostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost unsafe.Pointer) bool
	Costs() unsafe.Pointer
	DataType() uint32
	DatabaseInternal() uint64
	DrawCount() uint64
	DrawTraceCount() uint64
	DrawTraces() unsafe.Pointer
	Draws() unsafe.Pointer
	DurationForDrawDataMaster(draw uint32, master uint16) uint64
	EncodeError(encode bool) (objectivec.IObject, error)
	EncodeWithCoder(coder foundation.INSCoder)
	EncoderCount() uint64
	EncoderFromFunctionIndex(index uint32) unsafe.Pointer
	Encoders() unsafe.Pointer
	EnumerateBinariesForDrawEnumerator(draw uint32, enumerator VoidHandler)
	EnumerateBinariesForEncoderEnumerator(encoder uint32, enumerator VoidHandler)
	EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator VoidHandler)
	EnumerateBinariesForPipelineStateEnumerator(state uint64, enumerator VoidHandler)
	EnumerateBinaryRangesForCliqueUscDataEnumerator(clique unsafe.Pointer, data objectivec.IObject, enumerator VoidHandler)
	EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator VoidHandler)
	EnumerateDrawsForEncoderEnumerator(encoder uint32, enumerator VoidHandler)
	EnumerateDrawsForPipelineStateEnumerator(state uint64, enumerator VoidHandler)
	EnumerateEncoders(encoders VoidHandler)
	EnumerateInstructionsForCliqueUscDataEnumerator(clique unsafe.Pointer, data objectivec.IObject, enumerator VoidHandler)
	EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator VoidHandler)
	EnumerateKickAtFunctionIndexEnumerator(index uint32, enumerator VoidHandler)
	EnumeratePipelineStates(states VoidHandler)
	EnumerateUniqueTracesForBinaryEnumerator(binary uint32, enumerator VoidHandler)
	ExecutionHistoryForCliqueUscDelegate(clique unsafe.Pointer, usc unsafe.Pointer, delegate objectivec.IObject)
	ExecutionHistoryForCliqueUscDelegateRequiresTimestampCount(clique unsafe.Pointer, usc unsafe.Pointer, delegate objectivec.IObject, timestamp bool, count uint32)
	ExecutionHistoryForCliqueUscIndexDelegate(clique uint32, index uint32, delegate objectivec.IObject)
	ExecutionHistoryForDrawProgramTypeDelegateProgressController(draw uint32, type_ uint16, delegate objectivec.IObject, controller objectivec.IObject)
	ExecutionHistoryForPipelineStateDelegateProgressControllerCliqueFilter(state uint64, delegate objectivec.IObject, controller objectivec.IObject, filter VoidHandler)
	ExecutionHistoryForPipelineStateProgramTypeDelegateProgressController(state uint64, type_ uint16, delegate objectivec.IObject, controller objectivec.IObject)
	FailedUSCIndexCount() uint64
	FailedUSCIndexes() unsafe.Pointer
	FragmentPositionCount() uint64
	FragmentPositions() unsafe.Pointer
	GlobalGPUTime() uint64
	GpuCost() unsafe.Pointer
	GpuInfo() IGTMioGPUInfo
	GpuTime() uint64
	InitializeFromDatabase() GTMioTraceTimelineData
	InstructionCountForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) uint64
	KickDurationForEncoder(encoder uint32) uint64
	KickDurationForEncoderDataMaster(encoder uint32, master uint16) uint64
	Kicks() unsafe.Pointer
	KicksCount() uint64
	MGPUs() foundation.INSArray
	MetalFXInfo() unsafe.Pointer
	MetalFXInfoCount() uint64
	NumDrawsForEncoder(encoder uint32) uint64
	NumDrawsForPipelineState(state uint64) uint64
	ParentData() IGTMioTraceData
	SetParentData(value IGTMioTraceData)
	PipelineStateCount() uint64
	PipelineStateIdForCliqueIndex(index unsafe.Pointer) uint64
	ProfiledState() uint32
	ProfiledWithOverlapEnabled() bool
	ReferenceComputePositionForClique(clique unsafe.Pointer) unsafe.Pointer
	RiaTraceCount() uint64
	RiaTraces() unsafe.Pointer
	SampledCores() uint32
	ShaderBinaryInfo() unsafe.Pointer
	ShaderBinaryInfoCount() uint64
	SignpostPipelineStateCount() uint64
	SignpostPipelineStates() unsafe.Pointer
	SignpostProcessCount() uint64
	SignpostProcesses() unsafe.Pointer
	SignpostShaderCount() uint64
	SignpostShaders() unsafe.Pointer
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
	InitWithAPSTraceDataTimelineDataStreamDataTimelineTypeOptionsParentData(data unsafe.Pointer, data2 unsafe.Pointer, data3 objectivec.IObject, type_ uint32, options uint, data4 objectivec.IObject) GTMioTraceTimelineData
	InitWithCoder(coder foundation.INSCoder) GTMioTraceTimelineData
	InitWithDecodedDictionaryStreamDataParentData(dictionary objectivec.IObject, data objectivec.IObject, data2 objectivec.IObject) GTMioTraceTimelineData
	InitWithSerializedDataStreamDataParentData(data objectivec.IObject, data2 objectivec.IObject, data3 objectivec.IObject) GTMioTraceTimelineData
	InitWithTraceDatabaseDeallocatorParentData(database uint64, deallocator VoidHandler, data objectivec.IObject) GTMioTraceTimelineData
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (g GTMioTraceTimelineData) Init() GTMioTraceTimelineData {
	rv := objc.Send[GTMioTraceTimelineData](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioTraceTimelineData) Autorelease() GTMioTraceTimelineData {
	rv := objc.Send[GTMioTraceTimelineData](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioTraceTimelineData creates a new GTMioTraceTimelineData instance.
func NewGTMioTraceTimelineData() GTMioTraceTimelineData {
	class := getGTMioTraceTimelineDataClass()
	rv := objc.Send[GTMioTraceTimelineData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/initWithAPSTraceData:timelineData:streamData:timelineType:options:parentData:
func NewGTMioTraceTimelineDataWithAPSTraceDataTimelineDataStreamDataTimelineTypeOptionsParentData(data unsafe.Pointer, data2 unsafe.Pointer, data3 objectivec.IObject, type_ uint32, options uint, data4 objectivec.IObject) GTMioTraceTimelineData {
	instance := getGTMioTraceTimelineDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAPSTraceData:timelineData:streamData:timelineType:options:parentData:"), data, data2, data3, type_, options, data4)
	return GTMioTraceTimelineDataFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/initWithCoder:
func NewGTMioTraceTimelineDataWithCoder(coder objectivec.IObject) GTMioTraceTimelineData {
	instance := getGTMioTraceTimelineDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTMioTraceTimelineDataFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/initWithDecodedDictionary:streamData:parentData:
func NewGTMioTraceTimelineDataWithDecodedDictionaryStreamDataParentData(dictionary objectivec.IObject, data objectivec.IObject, data2 objectivec.IObject) GTMioTraceTimelineData {
	instance := getGTMioTraceTimelineDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDecodedDictionary:streamData:parentData:"), dictionary, data, data2)
	return GTMioTraceTimelineDataFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/initWithSerializedData:streamData:parentData:
func NewGTMioTraceTimelineDataWithSerializedDataStreamDataParentData(data objectivec.IObject, data2 objectivec.IObject, data3 objectivec.IObject) GTMioTraceTimelineData {
	instance := getGTMioTraceTimelineDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSerializedData:streamData:parentData:"), data, data2, data3)
	return GTMioTraceTimelineDataFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/_cacheExeuctionHistory:
func (g GTMioTraceTimelineData) _cacheExeuctionHistory(history objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("_cacheExeuctionHistory:"), history)
}

// CacheExeuctionHistory is an exported wrapper for the private method _cacheExeuctionHistory.
func (g GTMioTraceTimelineData) CacheExeuctionHistory(history objectivec.IObject) {
	g._cacheExeuctionHistory(history)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/_waitPendingExecutionHistory:
func (g GTMioTraceTimelineData) _waitPendingExecutionHistory(history objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_waitPendingExecutionHistory:"), history)
	return objectivec.Object{ID: rv}
}

// WaitPendingExecutionHistory is an exported wrapper for the private method _waitPendingExecutionHistory.
func (g GTMioTraceTimelineData) WaitPendingExecutionHistory(history objectivec.IObject) objectivec.IObject {
	return g._waitPendingExecutionHistory(history)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/archivedData:
func (g GTMioTraceTimelineData) ArchivedData(data []objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("archivedData:"), objectivec.IObjectSliceToNSArray(data))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/archivedDataWithCompression:originalSize:error:
func (g GTMioTraceTimelineData) ArchivedDataWithCompressionOriginalSizeError(compression int64, size unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](g.ID, objc.Sel("archivedDataWithCompression:originalSize:error:"), compression, size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/binaryForDraw:programType:
func (g GTMioTraceTimelineData) BinaryForDrawProgramType(draw uint32, type_ uint16) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("binaryForDraw:programType:"), draw, type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/binaryForPipelineState:programType:
func (g GTMioTraceTimelineData) BinaryForPipelineStateProgramType(state uint64, type_ uint16) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("binaryForPipelineState:programType:"), state, type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/childCliqueOfClique:
func (g GTMioTraceTimelineData) ChildCliqueOfClique(clique unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("childCliqueOfClique:"), clique)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/cliqueFromCliqueIndex:
func (g GTMioTraceTimelineData) CliqueFromCliqueIndex(index unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("cliqueFromCliqueIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/coalescedFunctionIndexForEncoderFunctionIndex:
func (g GTMioTraceTimelineData) CoalescedFunctionIndexForEncoderFunctionIndex(index uint32) uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("coalescedFunctionIndexForEncoderFunctionIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/costForContext:cost:
func (g GTMioTraceTimelineData) CostForContextCost(context unsafe.Pointer, cost unsafe.Pointer) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("costForContext:cost:"), context, cost)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/costForLevel:levelIdentifier:scope:scopeIdentifier:cost:
func (g GTMioTraceTimelineData) CostForLevelLevelIdentifierScopeScopeIdentifierCost(level uint16, identifier uint32, scope uint16, identifier2 uint64, cost unsafe.Pointer) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("costForLevel:levelIdentifier:scope:scopeIdentifier:cost:"), level, identifier, scope, identifier2, cost)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/costForScope:scopeIdentifier:cost:
func (g GTMioTraceTimelineData) CostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost unsafe.Pointer) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("costForScope:scopeIdentifier:cost:"), scope, identifier, cost)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/durationForDraw:dataMaster:
func (g GTMioTraceTimelineData) DurationForDrawDataMaster(draw uint32, master uint16) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("durationForDraw:dataMaster:"), draw, master)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/encode:error:
func (g GTMioTraceTimelineData) EncodeError(encode bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](g.ID, objc.Sel("encode:error:"), encode, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/encodeWithCoder:
func (g GTMioTraceTimelineData) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/encoderFromFunctionIndex:
func (g GTMioTraceTimelineData) EncoderFromFunctionIndex(index uint32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("encoderFromFunctionIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/enumerateBinariesForDraw:enumerator:
func (g GTMioTraceTimelineData) EnumerateBinariesForDrawEnumerator(draw uint32, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateBinariesForDraw:enumerator:"), draw, _block1)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/enumerateBinariesForEncoder:enumerator:
func (g GTMioTraceTimelineData) EnumerateBinariesForEncoderEnumerator(encoder uint32, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateBinariesForEncoder:enumerator:"), encoder, _block1)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/enumerateBinariesForForCliqueAtIndex:uscIndex:enumerator:
func (g GTMioTraceTimelineData) EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator VoidHandler) {
	_block2, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateBinariesForForCliqueAtIndex:uscIndex:enumerator:"), index, index2, _block2)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/enumerateBinariesForPipelineState:enumerator:
func (g GTMioTraceTimelineData) EnumerateBinariesForPipelineStateEnumerator(state uint64, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateBinariesForPipelineState:enumerator:"), state, _block1)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/enumerateBinaryRangesForClique:uscData:enumerator:
func (g GTMioTraceTimelineData) EnumerateBinaryRangesForCliqueUscDataEnumerator(clique unsafe.Pointer, data objectivec.IObject, enumerator VoidHandler) {
	_block2, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateBinaryRangesForClique:uscData:enumerator:"), clique, data, _block2)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/enumerateBinaryRangesForCliqueAtIndex:uscIndex:enumerator:
func (g GTMioTraceTimelineData) EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator VoidHandler) {
	_block2, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateBinaryRangesForCliqueAtIndex:uscIndex:enumerator:"), index, index2, _block2)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/enumerateDrawsForEncoder:enumerator:
func (g GTMioTraceTimelineData) EnumerateDrawsForEncoderEnumerator(encoder uint32, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateDrawsForEncoder:enumerator:"), encoder, _block1)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/enumerateDrawsForPipelineState:enumerator:
func (g GTMioTraceTimelineData) EnumerateDrawsForPipelineStateEnumerator(state uint64, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateDrawsForPipelineState:enumerator:"), state, _block1)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/enumerateEncoders:
func (g GTMioTraceTimelineData) EnumerateEncoders(encoders VoidHandler) {
	_block0, _ := NewVoidBlock(encoders)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateEncoders:"), _block0)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/enumerateInstructionsForClique:uscData:enumerator:
func (g GTMioTraceTimelineData) EnumerateInstructionsForCliqueUscDataEnumerator(clique unsafe.Pointer, data objectivec.IObject, enumerator VoidHandler) {
	_block2, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateInstructionsForClique:uscData:enumerator:"), clique, data, _block2)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/enumerateInstructionsForCliqueAtIndex:uscIndex:enumerator:
func (g GTMioTraceTimelineData) EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator VoidHandler) {
	_block2, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateInstructionsForCliqueAtIndex:uscIndex:enumerator:"), index, index2, _block2)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/enumerateKickAtFunctionIndex:enumerator:
func (g GTMioTraceTimelineData) EnumerateKickAtFunctionIndexEnumerator(index uint32, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateKickAtFunctionIndex:enumerator:"), index, _block1)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/enumeratePipelineStates:
func (g GTMioTraceTimelineData) EnumeratePipelineStates(states VoidHandler) {
	_block0, _ := NewVoidBlock(states)
	objc.Send[objc.ID](g.ID, objc.Sel("enumeratePipelineStates:"), _block0)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/enumerateUniqueTracesForBinary:enumerator:
func (g GTMioTraceTimelineData) EnumerateUniqueTracesForBinaryEnumerator(binary uint32, enumerator VoidHandler) {
	_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateUniqueTracesForBinary:enumerator:"), binary, _block1)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/executionHistoryForClique:usc:delegate:
func (g GTMioTraceTimelineData) ExecutionHistoryForCliqueUscDelegate(clique unsafe.Pointer, usc unsafe.Pointer, delegate objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("executionHistoryForClique:usc:delegate:"), clique, usc, delegate)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/executionHistoryForClique:usc:delegate:requiresTimestamp:count:
func (g GTMioTraceTimelineData) ExecutionHistoryForCliqueUscDelegateRequiresTimestampCount(clique unsafe.Pointer, usc unsafe.Pointer, delegate objectivec.IObject, timestamp bool, count uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("executionHistoryForClique:usc:delegate:requiresTimestamp:count:"), objc.CArray(clique), objc.CArray(usc), delegate, timestamp, count)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/executionHistoryForClique:uscIndex:delegate:
func (g GTMioTraceTimelineData) ExecutionHistoryForCliqueUscIndexDelegate(clique uint32, index uint32, delegate objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("executionHistoryForClique:uscIndex:delegate:"), clique, index, delegate)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/executionHistoryForDraw:programType:delegate:progressController:
func (g GTMioTraceTimelineData) ExecutionHistoryForDrawProgramTypeDelegateProgressController(draw uint32, type_ uint16, delegate objectivec.IObject, controller objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("executionHistoryForDraw:programType:delegate:progressController:"), draw, type_, delegate, controller)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/executionHistoryForPipelineState:delegate:progressController:cliqueFilter:
func (g GTMioTraceTimelineData) ExecutionHistoryForPipelineStateDelegateProgressControllerCliqueFilter(state uint64, delegate objectivec.IObject, controller objectivec.IObject, filter VoidHandler) {
	_block3, _ := NewVoidBlock(filter)
	objc.Send[objc.ID](g.ID, objc.Sel("executionHistoryForPipelineState:delegate:progressController:cliqueFilter:"), state, delegate, controller, _block3)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/executionHistoryForPipelineState:programType:delegate:progressController:
func (g GTMioTraceTimelineData) ExecutionHistoryForPipelineStateProgramTypeDelegateProgressController(state uint64, type_ uint16, delegate objectivec.IObject, controller objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("executionHistoryForPipelineState:programType:delegate:progressController:"), state, type_, delegate, controller)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/initializeFromDatabase
func (g GTMioTraceTimelineData) InitializeFromDatabase() GTMioTraceTimelineData {
	rv := objc.Send[GTMioTraceTimelineData](g.ID, objc.Sel("initializeFromDatabase"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/instructionCountForScope:scopeIdentifier:dataMaster:
func (g GTMioTraceTimelineData) InstructionCountForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("instructionCountForScope:scopeIdentifier:dataMaster:"), scope, identifier, master)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/kickDurationForEncoder:
func (g GTMioTraceTimelineData) KickDurationForEncoder(encoder uint32) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("kickDurationForEncoder:"), encoder)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/kickDurationForEncoder:dataMaster:
func (g GTMioTraceTimelineData) KickDurationForEncoderDataMaster(encoder uint32, master uint16) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("kickDurationForEncoder:dataMaster:"), encoder, master)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/numDrawsForEncoder:
func (g GTMioTraceTimelineData) NumDrawsForEncoder(encoder uint32) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("numDrawsForEncoder:"), encoder)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/numDrawsForPipelineState:
func (g GTMioTraceTimelineData) NumDrawsForPipelineState(state uint64) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("numDrawsForPipelineState:"), state)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/pipelineStateIdForCliqueIndex:
func (g GTMioTraceTimelineData) PipelineStateIdForCliqueIndex(index unsafe.Pointer) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("pipelineStateIdForCliqueIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/referenceComputePositionForClique:
func (g GTMioTraceTimelineData) ReferenceComputePositionForClique(clique unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("referenceComputePositionForClique:"), clique)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/totalCostForScope:scopeIdentifier:dataMaster:
func (g GTMioTraceTimelineData) TotalCostForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("totalCostForScope:scopeIdentifier:dataMaster:"), scope, identifier, master)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/totalCostForScope:scopeIdentifier:programType:
func (g GTMioTraceTimelineData) TotalCostForScopeScopeIdentifierProgramType(scope uint16, identifier uint64, type_ uint16) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("totalCostForScope:scopeIdentifier:programType:"), scope, identifier, type_)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/initWithAPSTraceData:timelineData:streamData:timelineType:options:parentData:
func (g GTMioTraceTimelineData) InitWithAPSTraceDataTimelineDataStreamDataTimelineTypeOptionsParentData(data unsafe.Pointer, data2 unsafe.Pointer, data3 objectivec.IObject, type_ uint32, options uint, data4 objectivec.IObject) GTMioTraceTimelineData {
	rv := objc.Send[GTMioTraceTimelineData](g.ID, objc.Sel("initWithAPSTraceData:timelineData:streamData:timelineType:options:parentData:"), data, data2, data3, type_, options, data4)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/initWithCoder:
func (g GTMioTraceTimelineData) InitWithCoder(coder foundation.INSCoder) GTMioTraceTimelineData {
	rv := objc.Send[GTMioTraceTimelineData](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/initWithDecodedDictionary:streamData:parentData:
func (g GTMioTraceTimelineData) InitWithDecodedDictionaryStreamDataParentData(dictionary objectivec.IObject, data objectivec.IObject, data2 objectivec.IObject) GTMioTraceTimelineData {
	rv := objc.Send[GTMioTraceTimelineData](g.ID, objc.Sel("initWithDecodedDictionary:streamData:parentData:"), dictionary, data, data2)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/initWithSerializedData:streamData:parentData:
func (g GTMioTraceTimelineData) InitWithSerializedDataStreamDataParentData(data objectivec.IObject, data2 objectivec.IObject, data3 objectivec.IObject) GTMioTraceTimelineData {
	rv := objc.Send[GTMioTraceTimelineData](g.ID, objc.Sel("initWithSerializedData:streamData:parentData:"), data, data2, data3)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/initWithTraceDatabase:deallocator:parentData:
func (g GTMioTraceTimelineData) InitWithTraceDatabaseDeallocatorParentData(database uint64, deallocator VoidHandler, data objectivec.IObject) GTMioTraceTimelineData {
	_block1, _ := NewVoidBlock(deallocator)
	rv := objc.Send[objc.ID](g.ID, objc.Sel("initWithTraceDatabase:deallocator:parentData:"), database, _block1, data)
	return GTMioTraceTimelineDataFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/supportsSecureCoding
func (_GTMioTraceTimelineDataClass GTMioTraceTimelineDataClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_GTMioTraceTimelineDataClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/binaries
func (g GTMioTraceTimelineData) Binaries() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("binaries"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/computePositionCount
func (g GTMioTraceTimelineData) ComputePositionCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("computePositionCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/computePositions
func (g GTMioTraceTimelineData) ComputePositions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("computePositions"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/consistentStateAchieved
func (g GTMioTraceTimelineData) ConsistentStateAchieved() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("consistentStateAchieved"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/costCount
func (g GTMioTraceTimelineData) CostCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("costCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/costs
func (g GTMioTraceTimelineData) Costs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("costs"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/dataType
func (g GTMioTraceTimelineData) DataType() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("dataType"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/databaseInternal
func (g GTMioTraceTimelineData) DatabaseInternal() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("databaseInternal"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/debugDescription
func (g GTMioTraceTimelineData) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/description
func (g GTMioTraceTimelineData) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/drawCount
func (g GTMioTraceTimelineData) DrawCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("drawCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/drawTraceCount
func (g GTMioTraceTimelineData) DrawTraceCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("drawTraceCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/drawTraces
func (g GTMioTraceTimelineData) DrawTraces() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("drawTraces"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/draws
func (g GTMioTraceTimelineData) Draws() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("draws"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/encoderCount
func (g GTMioTraceTimelineData) EncoderCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("encoderCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/encoders
func (g GTMioTraceTimelineData) Encoders() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("encoders"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/failedUSCIndexCount
func (g GTMioTraceTimelineData) FailedUSCIndexCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("failedUSCIndexCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/failedUSCIndexes
func (g GTMioTraceTimelineData) FailedUSCIndexes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("failedUSCIndexes"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/fragmentPositionCount
func (g GTMioTraceTimelineData) FragmentPositionCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("fragmentPositionCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/fragmentPositions
func (g GTMioTraceTimelineData) FragmentPositions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("fragmentPositions"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/globalGPUTime
func (g GTMioTraceTimelineData) GlobalGPUTime() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("globalGPUTime"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/gpuCost
func (g GTMioTraceTimelineData) GpuCost() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("gpuCost"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/gpuInfo
func (g GTMioTraceTimelineData) GpuInfo() IGTMioGPUInfo {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gpuInfo"))
	return GTMioGPUInfoFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/gpuTime
func (g GTMioTraceTimelineData) GpuTime() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("gpuTime"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/hash
func (g GTMioTraceTimelineData) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/kicks
func (g GTMioTraceTimelineData) Kicks() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("kicks"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/kicksCount
func (g GTMioTraceTimelineData) KicksCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("kicksCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/mGPUs
func (g GTMioTraceTimelineData) MGPUs() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("mGPUs"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/metalFXInfo
func (g GTMioTraceTimelineData) MetalFXInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("metalFXInfo"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/metalFXInfoCount
func (g GTMioTraceTimelineData) MetalFXInfoCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("metalFXInfoCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/parentData
func (g GTMioTraceTimelineData) ParentData() IGTMioTraceData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("parentData"))
	return GTMioTraceDataFromID(objc.ID(rv))
}
func (g GTMioTraceTimelineData) SetParentData(value IGTMioTraceData) {
	objc.Send[struct{}](g.ID, objc.Sel("setParentData:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/pipelineStateCount
func (g GTMioTraceTimelineData) PipelineStateCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("pipelineStateCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/profiledState
func (g GTMioTraceTimelineData) ProfiledState() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("profiledState"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/profiledWithOverlapEnabled
func (g GTMioTraceTimelineData) ProfiledWithOverlapEnabled() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("profiledWithOverlapEnabled"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/riaTraceCount
func (g GTMioTraceTimelineData) RiaTraceCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("riaTraceCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/riaTraces
func (g GTMioTraceTimelineData) RiaTraces() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("riaTraces"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/sampledCores
func (g GTMioTraceTimelineData) SampledCores() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("sampledCores"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/shaderBinaryInfo
func (g GTMioTraceTimelineData) ShaderBinaryInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("shaderBinaryInfo"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/shaderBinaryInfoCount
func (g GTMioTraceTimelineData) ShaderBinaryInfoCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("shaderBinaryInfoCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/signpostPipelineStateCount
func (g GTMioTraceTimelineData) SignpostPipelineStateCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("signpostPipelineStateCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/signpostPipelineStates
func (g GTMioTraceTimelineData) SignpostPipelineStates() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("signpostPipelineStates"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/signpostProcessCount
func (g GTMioTraceTimelineData) SignpostProcessCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("signpostProcessCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/signpostProcesses
func (g GTMioTraceTimelineData) SignpostProcesses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("signpostProcesses"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/signpostShaderCount
func (g GTMioTraceTimelineData) SignpostShaderCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("signpostShaderCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/signpostShaders
func (g GTMioTraceTimelineData) SignpostShaders() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("signpostShaders"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/signpostStrings
func (g GTMioTraceTimelineData) SignpostStrings() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("signpostStrings"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/streamData
func (g GTMioTraceTimelineData) StreamData() IGTShaderProfilerStreamData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("streamData"))
	return GTShaderProfilerStreamDataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/superclass
func (g GTMioTraceTimelineData) Superclass() objc.Class {
	rv := objc.Send[objc.Class](g.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/timelineCounters
func (g GTMioTraceTimelineData) TimelineCounters() IGTMioTimelineCounters {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("timelineCounters"))
	return GTMioTimelineCountersFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/timelineDuration
func (g GTMioTraceTimelineData) TimelineDuration() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("timelineDuration"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/timestampBegin
func (g GTMioTraceTimelineData) TimestampBegin() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("timestampBegin"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/timestampEnd
func (g GTMioTraceTimelineData) TimestampEnd() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("timestampEnd"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/totalCliqueCost
func (g GTMioTraceTimelineData) TotalCliqueCost() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("totalCliqueCost"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/totalCores
func (g GTMioTraceTimelineData) TotalCores() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("totalCores"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTimelineData/uscs
func (g GTMioTraceTimelineData) Uscs() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("uscs"))
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
func (g GTMioTraceTimelineData) EnumerateBinaryRangesForCliqueUscDataEnumeratorSync(ctx context.Context, clique unsafe.Pointer, data objectivec.IObject) error {
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
func (g GTMioTraceTimelineData) EnumerateInstructionsForCliqueUscDataEnumeratorSync(ctx context.Context, clique unsafe.Pointer, data objectivec.IObject) error {
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

// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioTraceData] class.
var (
	_GTMioTraceDataClass     GTMioTraceDataClass
	_GTMioTraceDataClassOnce sync.Once
)

func getGTMioTraceDataClass() GTMioTraceDataClass {
	_GTMioTraceDataClassOnce.Do(func() {
		_GTMioTraceDataClass = GTMioTraceDataClass{class: objc.GetClass("GTMioTraceData")}
	})
	return _GTMioTraceDataClass
}

// GetGTMioTraceDataClass returns the class object for GTMioTraceData.
func GetGTMioTraceDataClass() GTMioTraceDataClass {
	return getGTMioTraceDataClass()
}

type GTMioTraceDataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioTraceDataClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioTraceDataClass) Alloc() GTMioTraceData {
	rv := objc.Send[GTMioTraceData](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTMioTraceData.LLVMConnectionHeld]
//   - [GTMioTraceData._storeError]
//   - [GTMioTraceData.ArchiveToFileCompressDataError]
//   - [GTMioTraceData.ArchivedDataError]
//   - [GTMioTraceData.Binaries]
//   - [GTMioTraceData.BinaryForDrawProgramType]
//   - [GTMioTraceData.BinaryForPipelineStateProgramType]
//   - [GTMioTraceData.BuildNonOverlappingCounters]
//   - [GTMioTraceData.Cancel]
//   - [GTMioTraceData.CancelToken]
//   - [GTMioTraceData.CliqueFromCliqueIndex]
//   - [GTMioTraceData.ComputePositionCount]
//   - [GTMioTraceData.ComputePositions]
//   - [GTMioTraceData.ConsistentStateAchieved]
//   - [GTMioTraceData.CostCount]
//   - [GTMioTraceData.CostForContextCost]
//   - [GTMioTraceData.CostForLevelLevelIdentifierScopeScopeIdentifierCost]
//   - [GTMioTraceData.CostForScopeScopeIdentifierCost]
//   - [GTMioTraceData.CostTimeline]
//   - [GTMioTraceData.Costs]
//   - [GTMioTraceData.DrawCount]
//   - [GTMioTraceData.DrawTraceCount]
//   - [GTMioTraceData.DrawTraces]
//   - [GTMioTraceData.Draws]
//   - [GTMioTraceData.DurationForDrawDataMaster]
//   - [GTMioTraceData.EncodeWithCoder]
//   - [GTMioTraceData.EncoderCount]
//   - [GTMioTraceData.EncoderFromFunctionIndex]
//   - [GTMioTraceData.Encoders]
//   - [GTMioTraceData.EnumerateBinariesForDrawEnumerator]
//   - [GTMioTraceData.EnumerateBinariesForEncoderEnumerator]
//   - [GTMioTraceData.EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceData.EnumerateBinariesForPipelineStateEnumerator]
//   - [GTMioTraceData.EnumerateBinaryRangesForCliqueUscDataEnumerator]
//   - [GTMioTraceData.EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceData.EnumerateDrawsForEncoderEnumerator]
//   - [GTMioTraceData.EnumerateDrawsForPipelineStateEnumerator]
//   - [GTMioTraceData.EnumerateEncoders]
//   - [GTMioTraceData.EnumerateInstructionsForCliqueUscDataEnumerator]
//   - [GTMioTraceData.EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceData.EnumerateKickAtFunctionIndexEnumerator]
//   - [GTMioTraceData.EnumeratePipelineStates]
//   - [GTMioTraceData.EnumerateUniqueTracesForBinaryEnumerator]
//   - [GTMioTraceData.ExecutionHistoryForCliqueUscIndexDelegate]
//   - [GTMioTraceData.ExecutionHistoryForDrawProgramTypeDelegateProgressController]
//   - [GTMioTraceData.ExecutionHistoryForPipelineStateProgramTypeDelegateProgressController]
//   - [GTMioTraceData.FragmentPositionCount]
//   - [GTMioTraceData.FragmentPositions]
//   - [GTMioTraceData.GpuCost]
//   - [GTMioTraceData.GpuInfo]
//   - [GTMioTraceData.GpuTime]
//   - [GTMioTraceData.HandleBatchIdFilteredData]
//   - [GTMioTraceData.HasSeparateCostsTimeline]
//   - [GTMioTraceData.InstructionCountForScopeScopeIdentifierDataMaster]
//   - [GTMioTraceData.IsMio]
//   - [GTMioTraceData.KickDurationForEncoder]
//   - [GTMioTraceData.KickDurationForEncoderDataMaster]
//   - [GTMioTraceData.Kicks]
//   - [GTMioTraceData.KicksCount]
//   - [GTMioTraceData.LoadCostTimeline]
//   - [GTMioTraceData.LoadTimeline]
//   - [GTMioTraceData.LoadingCostTimeline]
//   - [GTMioTraceData.MGPUs]
//   - [GTMioTraceData.MetalFXCallDuration]
//   - [GTMioTraceData.NonOverlappingCounters]
//   - [GTMioTraceData.NonOverlappingTimeline]
//   - [GTMioTraceData.NumDrawsForPipelineState]
//   - [GTMioTraceData.ObservePerDrawCounterUpdates]
//   - [GTMioTraceData.OverlappingTimeline]
//   - [GTMioTraceData.PipelineStateCount]
//   - [GTMioTraceData.ProfiledState]
//   - [GTMioTraceData.RequestCostTimeline]
//   - [GTMioTraceData.RiaTraceCount]
//   - [GTMioTraceData.RiaTraces]
//   - [GTMioTraceData.ShaderBinaryInfo]
//   - [GTMioTraceData.ShaderBinaryInfoCount]
//   - [GTMioTraceData.StreamData]
//   - [GTMioTraceData.SyncedBatchIDObserverCopy]
//   - [GTMioTraceData.TimelineCounters]
//   - [GTMioTraceData.TimestampBegin]
//   - [GTMioTraceData.TimestampEnd]
//   - [GTMioTraceData.TotalCostForScopeScopeIdentifierDataMaster]
//   - [GTMioTraceData.UnixTimestamp]
//   - [GTMioTraceData.Uscs]
//   - [GTMioTraceData.UseMinimumTracingMode]
//   - [GTMioTraceData.InitWithArchivedDataError]
//   - [GTMioTraceData.InitWithCoder]
//   - [GTMioTraceData.InitWithStreamDataLlvmHelperPathOptions]
//   - [GTMioTraceData.InitWithTraceDatabaseDeallocator]
//   - [GTMioTraceData.Version]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData
type GTMioTraceData struct {
	objectivec.Object
}

// GTMioTraceDataFromID constructs a [GTMioTraceData] from an objc.ID.
func GTMioTraceDataFromID(id objc.ID) GTMioTraceData {
	return GTMioTraceData{objectivec.Object{ID: id}}
}
// Ensure GTMioTraceData implements IGTMioTraceData.
var _ IGTMioTraceData = GTMioTraceData{}

// An interface definition for the [GTMioTraceData] class.
//
// # Methods
//
//   - [IGTMioTraceData.LLVMConnectionHeld]
//   - [IGTMioTraceData._storeError]
//   - [IGTMioTraceData.ArchiveToFileCompressDataError]
//   - [IGTMioTraceData.ArchivedDataError]
//   - [IGTMioTraceData.Binaries]
//   - [IGTMioTraceData.BinaryForDrawProgramType]
//   - [IGTMioTraceData.BinaryForPipelineStateProgramType]
//   - [IGTMioTraceData.BuildNonOverlappingCounters]
//   - [IGTMioTraceData.Cancel]
//   - [IGTMioTraceData.CancelToken]
//   - [IGTMioTraceData.CliqueFromCliqueIndex]
//   - [IGTMioTraceData.ComputePositionCount]
//   - [IGTMioTraceData.ComputePositions]
//   - [IGTMioTraceData.ConsistentStateAchieved]
//   - [IGTMioTraceData.CostCount]
//   - [IGTMioTraceData.CostForContextCost]
//   - [IGTMioTraceData.CostForLevelLevelIdentifierScopeScopeIdentifierCost]
//   - [IGTMioTraceData.CostForScopeScopeIdentifierCost]
//   - [IGTMioTraceData.CostTimeline]
//   - [IGTMioTraceData.Costs]
//   - [IGTMioTraceData.DrawCount]
//   - [IGTMioTraceData.DrawTraceCount]
//   - [IGTMioTraceData.DrawTraces]
//   - [IGTMioTraceData.Draws]
//   - [IGTMioTraceData.DurationForDrawDataMaster]
//   - [IGTMioTraceData.EncodeWithCoder]
//   - [IGTMioTraceData.EncoderCount]
//   - [IGTMioTraceData.EncoderFromFunctionIndex]
//   - [IGTMioTraceData.Encoders]
//   - [IGTMioTraceData.EnumerateBinariesForDrawEnumerator]
//   - [IGTMioTraceData.EnumerateBinariesForEncoderEnumerator]
//   - [IGTMioTraceData.EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator]
//   - [IGTMioTraceData.EnumerateBinariesForPipelineStateEnumerator]
//   - [IGTMioTraceData.EnumerateBinaryRangesForCliqueUscDataEnumerator]
//   - [IGTMioTraceData.EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator]
//   - [IGTMioTraceData.EnumerateDrawsForEncoderEnumerator]
//   - [IGTMioTraceData.EnumerateDrawsForPipelineStateEnumerator]
//   - [IGTMioTraceData.EnumerateEncoders]
//   - [IGTMioTraceData.EnumerateInstructionsForCliqueUscDataEnumerator]
//   - [IGTMioTraceData.EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator]
//   - [IGTMioTraceData.EnumerateKickAtFunctionIndexEnumerator]
//   - [IGTMioTraceData.EnumeratePipelineStates]
//   - [IGTMioTraceData.EnumerateUniqueTracesForBinaryEnumerator]
//   - [IGTMioTraceData.ExecutionHistoryForCliqueUscIndexDelegate]
//   - [IGTMioTraceData.ExecutionHistoryForDrawProgramTypeDelegateProgressController]
//   - [IGTMioTraceData.ExecutionHistoryForPipelineStateProgramTypeDelegateProgressController]
//   - [IGTMioTraceData.FragmentPositionCount]
//   - [IGTMioTraceData.FragmentPositions]
//   - [IGTMioTraceData.GpuCost]
//   - [IGTMioTraceData.GpuInfo]
//   - [IGTMioTraceData.GpuTime]
//   - [IGTMioTraceData.HandleBatchIdFilteredData]
//   - [IGTMioTraceData.HasSeparateCostsTimeline]
//   - [IGTMioTraceData.InstructionCountForScopeScopeIdentifierDataMaster]
//   - [IGTMioTraceData.IsMio]
//   - [IGTMioTraceData.KickDurationForEncoder]
//   - [IGTMioTraceData.KickDurationForEncoderDataMaster]
//   - [IGTMioTraceData.Kicks]
//   - [IGTMioTraceData.KicksCount]
//   - [IGTMioTraceData.LoadCostTimeline]
//   - [IGTMioTraceData.LoadTimeline]
//   - [IGTMioTraceData.LoadingCostTimeline]
//   - [IGTMioTraceData.MGPUs]
//   - [IGTMioTraceData.MetalFXCallDuration]
//   - [IGTMioTraceData.NonOverlappingCounters]
//   - [IGTMioTraceData.NonOverlappingTimeline]
//   - [IGTMioTraceData.NumDrawsForPipelineState]
//   - [IGTMioTraceData.ObservePerDrawCounterUpdates]
//   - [IGTMioTraceData.OverlappingTimeline]
//   - [IGTMioTraceData.PipelineStateCount]
//   - [IGTMioTraceData.ProfiledState]
//   - [IGTMioTraceData.RequestCostTimeline]
//   - [IGTMioTraceData.RiaTraceCount]
//   - [IGTMioTraceData.RiaTraces]
//   - [IGTMioTraceData.ShaderBinaryInfo]
//   - [IGTMioTraceData.ShaderBinaryInfoCount]
//   - [IGTMioTraceData.StreamData]
//   - [IGTMioTraceData.SyncedBatchIDObserverCopy]
//   - [IGTMioTraceData.TimelineCounters]
//   - [IGTMioTraceData.TimestampBegin]
//   - [IGTMioTraceData.TimestampEnd]
//   - [IGTMioTraceData.TotalCostForScopeScopeIdentifierDataMaster]
//   - [IGTMioTraceData.UnixTimestamp]
//   - [IGTMioTraceData.Uscs]
//   - [IGTMioTraceData.UseMinimumTracingMode]
//   - [IGTMioTraceData.InitWithArchivedDataError]
//   - [IGTMioTraceData.InitWithCoder]
//   - [IGTMioTraceData.InitWithStreamDataLlvmHelperPathOptions]
//   - [IGTMioTraceData.InitWithTraceDatabaseDeallocator]
//   - [IGTMioTraceData.Version]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData
type IGTMioTraceData interface {
	objectivec.IObject

	// Topic: Methods

	LLVMConnectionHeld() bool
	_storeError(_store bool) (objectivec.IObject, error)
	ArchiveToFileCompressDataError(file objectivec.IObject, data bool) (bool, error)
	ArchivedDataError(data bool) (objectivec.IObject, error)
	Binaries() objectivec.IObject
	BinaryForDrawProgramType(draw uint32, type_ uint16) objectivec.IObject
	BinaryForPipelineStateProgramType(state uint64, type_ uint16) objectivec.IObject
	BuildNonOverlappingCounters(counters unsafe.Pointer)
	Cancel()
	CancelToken(token uint64)
	CliqueFromCliqueIndex(index unsafe.Pointer) unsafe.Pointer
	ComputePositionCount() uint64
	ComputePositions() unsafe.Pointer
	ConsistentStateAchieved() bool
	CostCount() uint64
	CostForContextCost(context unsafe.Pointer, cost unsafe.Pointer) bool
	CostForLevelLevelIdentifierScopeScopeIdentifierCost(level uint16, identifier uint32, scope uint16, identifier2 uint64, cost unsafe.Pointer) bool
	CostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost unsafe.Pointer) bool
	CostTimeline() objectivec.IObject
	Costs() unsafe.Pointer
	DrawCount() uint64
	DrawTraceCount() uint64
	DrawTraces() unsafe.Pointer
	Draws() unsafe.Pointer
	DurationForDrawDataMaster(draw uint32, master uint16) uint64
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
	ExecutionHistoryForCliqueUscIndexDelegate(clique uint32, index uint32, delegate objectivec.IObject)
	ExecutionHistoryForDrawProgramTypeDelegateProgressController(draw uint32, type_ uint16, delegate objectivec.IObject, controller objectivec.IObject)
	ExecutionHistoryForPipelineStateProgramTypeDelegateProgressController(state uint64, type_ uint16, delegate objectivec.IObject, controller objectivec.IObject)
	FragmentPositionCount() uint64
	FragmentPositions() unsafe.Pointer
	GpuCost() unsafe.Pointer
	GpuInfo() IGTMioGPUInfo
	GpuTime() uint64
	HandleBatchIdFilteredData(data objectivec.IObject)
	HasSeparateCostsTimeline() bool
	InstructionCountForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) uint64
	IsMio() bool
	KickDurationForEncoder(encoder uint32) uint64
	KickDurationForEncoderDataMaster(encoder uint32, master uint16) uint64
	Kicks() unsafe.Pointer
	KicksCount() uint64
	LoadCostTimeline()
	LoadTimeline()
	LoadingCostTimeline() bool
	MGPUs() objectivec.IObject
	MetalFXCallDuration(duration uint64) uint64
	NonOverlappingCounters() IGTMioNonOverlappingCounters
	NonOverlappingTimeline() objectivec.IObject
	NumDrawsForPipelineState(state uint64) uint64
	ObservePerDrawCounterUpdates(updates objectivec.IObject) objectivec.IObject
	OverlappingTimeline() objectivec.IObject
	PipelineStateCount() uint64
	ProfiledState() uint32
	RequestCostTimeline(timeline VoidHandler) objectivec.IObject
	RiaTraceCount() uint64
	RiaTraces() unsafe.Pointer
	ShaderBinaryInfo() unsafe.Pointer
	ShaderBinaryInfoCount() uint64
	StreamData() IGTShaderProfilerStreamData
	SyncedBatchIDObserverCopy() objectivec.IObject
	TimelineCounters() objectivec.IObject
	TimestampBegin() uint64
	TimestampEnd() uint64
	TotalCostForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) float64
	UnixTimestamp() int64
	Uscs() objectivec.IObject
	UseMinimumTracingMode() bool
	InitWithArchivedDataError(data objectivec.IObject) (GTMioTraceData, error)
	InitWithCoder(coder foundation.INSCoder) GTMioTraceData
	InitWithStreamDataLlvmHelperPathOptions(data objectivec.IObject, path objectivec.IObject, options uint) GTMioTraceData
	InitWithTraceDatabaseDeallocator(database uint64, deallocator VoidHandler) GTMioTraceData
	Version() uint32
}

// Init initializes the instance.
func (g GTMioTraceData) Init() GTMioTraceData {
	rv := objc.Send[GTMioTraceData](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioTraceData) Autorelease() GTMioTraceData {
	rv := objc.Send[GTMioTraceData](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioTraceData creates a new GTMioTraceData instance.
func NewGTMioTraceData() GTMioTraceData {
	class := getGTMioTraceDataClass()
	rv := objc.Send[GTMioTraceData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/initWithArchivedData:error:
func NewGTMioTraceDataWithArchivedDataError(data objectivec.IObject) (GTMioTraceData, error) {
	var errorPtr objc.ID
	instance := getGTMioTraceDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArchivedData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return GTMioTraceData{}, foundation.NSErrorFrom(errorPtr)
	}
	return GTMioTraceDataFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/initWithCoder:
func NewGTMioTraceDataWithCoder(coder objectivec.IObject) GTMioTraceData {
	instance := getGTMioTraceDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTMioTraceDataFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/initWithStreamData:llvmHelperPath:options:
func NewGTMioTraceDataWithStreamDataLlvmHelperPathOptions(data objectivec.IObject, path objectivec.IObject, options uint) GTMioTraceData {
	instance := getGTMioTraceDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStreamData:llvmHelperPath:options:"), data, path, options)
	return GTMioTraceDataFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/_store:error:
func (g GTMioTraceData) _storeError(_store bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_store:error:"), _store, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// StoreError is an exported wrapper for the private method _storeError.
func (g GTMioTraceData) StoreError(_store bool) (objectivec.IObject, error) {
	return g._storeError(_store)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/archiveToFile:compressData:error:
func (g GTMioTraceData) ArchiveToFileCompressDataError(file objectivec.IObject, data bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](g.ID, objc.Sel("archiveToFile:compressData:error:"), file, data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("archiveToFile:compressData:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/archivedData:error:
func (g GTMioTraceData) ArchivedDataError(data bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](g.ID, objc.Sel("archivedData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/binaries
func (g GTMioTraceData) Binaries() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("binaries"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/binaryForDraw:programType:
func (g GTMioTraceData) BinaryForDrawProgramType(draw uint32, type_ uint16) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("binaryForDraw:programType:"), draw, type_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/binaryForPipelineState:programType:
func (g GTMioTraceData) BinaryForPipelineStateProgramType(state uint64, type_ uint16) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("binaryForPipelineState:programType:"), state, type_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/buildNonOverlappingCounters:
func (g GTMioTraceData) BuildNonOverlappingCounters(counters unsafe.Pointer) {
	objc.Send[objc.ID](g.ID, objc.Sel("buildNonOverlappingCounters:"), counters)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/cancel
func (g GTMioTraceData) Cancel() {
	objc.Send[objc.ID](g.ID, objc.Sel("cancel"))
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/cancelToken:
func (g GTMioTraceData) CancelToken(token uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("cancelToken:"), token)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/cliqueFromCliqueIndex:
func (g GTMioTraceData) CliqueFromCliqueIndex(index unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("cliqueFromCliqueIndex:"), index)
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/computePositionCount
func (g GTMioTraceData) ComputePositionCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("computePositionCount"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/computePositions
func (g GTMioTraceData) ComputePositions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("computePositions"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/costCount
func (g GTMioTraceData) CostCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("costCount"))
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/costForContext:cost:
func (g GTMioTraceData) CostForContextCost(context unsafe.Pointer, cost unsafe.Pointer) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("costForContext:cost:"), context, cost)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/costForLevel:levelIdentifier:scope:scopeIdentifier:cost:
func (g GTMioTraceData) CostForLevelLevelIdentifierScopeScopeIdentifierCost(level uint16, identifier uint32, scope uint16, identifier2 uint64, cost unsafe.Pointer) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("costForLevel:levelIdentifier:scope:scopeIdentifier:cost:"), level, identifier, scope, identifier2, cost)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/costForScope:scopeIdentifier:cost:
func (g GTMioTraceData) CostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost unsafe.Pointer) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("costForScope:scopeIdentifier:cost:"), scope, identifier, cost)
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/costs
func (g GTMioTraceData) Costs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("costs"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/drawCount
func (g GTMioTraceData) DrawCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("drawCount"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/drawTraceCount
func (g GTMioTraceData) DrawTraceCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("drawTraceCount"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/drawTraces
func (g GTMioTraceData) DrawTraces() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("drawTraces"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/draws
func (g GTMioTraceData) Draws() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("draws"))
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/durationForDraw:dataMaster:
func (g GTMioTraceData) DurationForDrawDataMaster(draw uint32, master uint16) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("durationForDraw:dataMaster:"), draw, master)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/encodeWithCoder:
func (g GTMioTraceData) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/encoderCount
func (g GTMioTraceData) EncoderCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("encoderCount"))
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/encoderFromFunctionIndex:
func (g GTMioTraceData) EncoderFromFunctionIndex(index uint32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("encoderFromFunctionIndex:"), index)
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/encoders
func (g GTMioTraceData) Encoders() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("encoders"))
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/enumerateBinariesForDraw:enumerator:
func (g GTMioTraceData) EnumerateBinariesForDrawEnumerator(draw uint32, enumerator VoidHandler) {
_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateBinariesForDraw:enumerator:"), draw, _block1)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/enumerateBinariesForEncoder:enumerator:
func (g GTMioTraceData) EnumerateBinariesForEncoderEnumerator(encoder uint32, enumerator VoidHandler) {
_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateBinariesForEncoder:enumerator:"), encoder, _block1)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/enumerateBinariesForForCliqueAtIndex:uscIndex:enumerator:
func (g GTMioTraceData) EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator VoidHandler) {
_block2, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateBinariesForForCliqueAtIndex:uscIndex:enumerator:"), index, index2, _block2)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/enumerateBinariesForPipelineState:enumerator:
func (g GTMioTraceData) EnumerateBinariesForPipelineStateEnumerator(state uint64, enumerator VoidHandler) {
_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateBinariesForPipelineState:enumerator:"), state, _block1)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/enumerateBinaryRangesForClique:uscData:enumerator:
func (g GTMioTraceData) EnumerateBinaryRangesForCliqueUscDataEnumerator(clique unsafe.Pointer, data objectivec.IObject, enumerator VoidHandler) {
_block2, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateBinaryRangesForClique:uscData:enumerator:"), clique, data, _block2)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/enumerateBinaryRangesForCliqueAtIndex:uscIndex:enumerator:
func (g GTMioTraceData) EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator VoidHandler) {
_block2, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateBinaryRangesForCliqueAtIndex:uscIndex:enumerator:"), index, index2, _block2)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/enumerateDrawsForEncoder:enumerator:
func (g GTMioTraceData) EnumerateDrawsForEncoderEnumerator(encoder uint32, enumerator VoidHandler) {
_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateDrawsForEncoder:enumerator:"), encoder, _block1)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/enumerateDrawsForPipelineState:enumerator:
func (g GTMioTraceData) EnumerateDrawsForPipelineStateEnumerator(state uint64, enumerator VoidHandler) {
_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateDrawsForPipelineState:enumerator:"), state, _block1)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/enumerateEncoders:
func (g GTMioTraceData) EnumerateEncoders(encoders VoidHandler) {
_block0, _ := NewVoidBlock(encoders)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateEncoders:"), _block0)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/enumerateInstructionsForClique:uscData:enumerator:
func (g GTMioTraceData) EnumerateInstructionsForCliqueUscDataEnumerator(clique unsafe.Pointer, data objectivec.IObject, enumerator VoidHandler) {
_block2, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateInstructionsForClique:uscData:enumerator:"), clique, data, _block2)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/enumerateInstructionsForCliqueAtIndex:uscIndex:enumerator:
func (g GTMioTraceData) EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator VoidHandler) {
_block2, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateInstructionsForCliqueAtIndex:uscIndex:enumerator:"), index, index2, _block2)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/enumerateKickAtFunctionIndex:enumerator:
func (g GTMioTraceData) EnumerateKickAtFunctionIndexEnumerator(index uint32, enumerator VoidHandler) {
_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateKickAtFunctionIndex:enumerator:"), index, _block1)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/enumeratePipelineStates:
func (g GTMioTraceData) EnumeratePipelineStates(states VoidHandler) {
_block0, _ := NewVoidBlock(states)
	objc.Send[objc.ID](g.ID, objc.Sel("enumeratePipelineStates:"), _block0)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/enumerateUniqueTracesForBinary:enumerator:
func (g GTMioTraceData) EnumerateUniqueTracesForBinaryEnumerator(binary uint32, enumerator VoidHandler) {
_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateUniqueTracesForBinary:enumerator:"), binary, _block1)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/executionHistoryForClique:uscIndex:delegate:
func (g GTMioTraceData) ExecutionHistoryForCliqueUscIndexDelegate(clique uint32, index uint32, delegate objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("executionHistoryForClique:uscIndex:delegate:"), clique, index, delegate)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/executionHistoryForDraw:programType:delegate:progressController:
func (g GTMioTraceData) ExecutionHistoryForDrawProgramTypeDelegateProgressController(draw uint32, type_ uint16, delegate objectivec.IObject, controller objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("executionHistoryForDraw:programType:delegate:progressController:"), draw, type_, delegate, controller)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/executionHistoryForPipelineState:programType:delegate:progressController:
func (g GTMioTraceData) ExecutionHistoryForPipelineStateProgramTypeDelegateProgressController(state uint64, type_ uint16, delegate objectivec.IObject, controller objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("executionHistoryForPipelineState:programType:delegate:progressController:"), state, type_, delegate, controller)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/fragmentPositionCount
func (g GTMioTraceData) FragmentPositionCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("fragmentPositionCount"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/fragmentPositions
func (g GTMioTraceData) FragmentPositions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("fragmentPositions"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/gpuCost
func (g GTMioTraceData) GpuCost() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("gpuCost"))
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/handleBatchIdFilteredData:
func (g GTMioTraceData) HandleBatchIdFilteredData(data objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("handleBatchIdFilteredData:"), data)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/instructionCountForScope:scopeIdentifier:dataMaster:
func (g GTMioTraceData) InstructionCountForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("instructionCountForScope:scopeIdentifier:dataMaster:"), scope, identifier, master)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/kickDurationForEncoder:
func (g GTMioTraceData) KickDurationForEncoder(encoder uint32) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("kickDurationForEncoder:"), encoder)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/kickDurationForEncoder:dataMaster:
func (g GTMioTraceData) KickDurationForEncoderDataMaster(encoder uint32, master uint16) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("kickDurationForEncoder:dataMaster:"), encoder, master)
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/kicks
func (g GTMioTraceData) Kicks() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("kicks"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/kicksCount
func (g GTMioTraceData) KicksCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("kicksCount"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/loadCostTimeline
func (g GTMioTraceData) LoadCostTimeline() {
	objc.Send[objc.ID](g.ID, objc.Sel("loadCostTimeline"))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/loadTimeline
func (g GTMioTraceData) LoadTimeline() {
	objc.Send[objc.ID](g.ID, objc.Sel("loadTimeline"))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/mGPUs
func (g GTMioTraceData) MGPUs() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("mGPUs"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/metalFXCallDuration:
func (g GTMioTraceData) MetalFXCallDuration(duration uint64) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("metalFXCallDuration:"), duration)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/numDrawsForPipelineState:
func (g GTMioTraceData) NumDrawsForPipelineState(state uint64) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("numDrawsForPipelineState:"), state)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/observePerDrawCounterUpdates:
func (g GTMioTraceData) ObservePerDrawCounterUpdates(updates objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("observePerDrawCounterUpdates:"), updates)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/pipelineStateCount
func (g GTMioTraceData) PipelineStateCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("pipelineStateCount"))
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/requestCostTimeline:
func (g GTMioTraceData) RequestCostTimeline(timeline VoidHandler) objectivec.IObject {
_block0, _ := NewVoidBlock(timeline)
	rv := objc.Send[objc.ID](g.ID, objc.Sel("requestCostTimeline:"), _block0)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/riaTraceCount
func (g GTMioTraceData) RiaTraceCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("riaTraceCount"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/riaTraces
func (g GTMioTraceData) RiaTraces() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("riaTraces"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/shaderBinaryInfo
func (g GTMioTraceData) ShaderBinaryInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("shaderBinaryInfo"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/shaderBinaryInfoCount
func (g GTMioTraceData) ShaderBinaryInfoCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("shaderBinaryInfoCount"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/syncedBatchIDObserverCopy
func (g GTMioTraceData) SyncedBatchIDObserverCopy() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("syncedBatchIDObserverCopy"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/timelineCounters
func (g GTMioTraceData) TimelineCounters() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("timelineCounters"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/timestampBegin
func (g GTMioTraceData) TimestampBegin() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("timestampBegin"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/timestampEnd
func (g GTMioTraceData) TimestampEnd() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("timestampEnd"))
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/totalCostForScope:scopeIdentifier:dataMaster:
func (g GTMioTraceData) TotalCostForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("totalCostForScope:scopeIdentifier:dataMaster:"), scope, identifier, master)
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/uscs
func (g GTMioTraceData) Uscs() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("uscs"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/initWithArchivedData:error:
func (g GTMioTraceData) InitWithArchivedDataError(data objectivec.IObject) (GTMioTraceData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](g.ID, objc.Sel("initWithArchivedData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return GTMioTraceData{}, foundation.NSErrorFrom(errorPtr)
	}
	return GTMioTraceDataFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/initWithCoder:
func (g GTMioTraceData) InitWithCoder(coder foundation.INSCoder) GTMioTraceData {
	rv := objc.Send[GTMioTraceData](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/initWithStreamData:llvmHelperPath:options:
func (g GTMioTraceData) InitWithStreamDataLlvmHelperPathOptions(data objectivec.IObject, path objectivec.IObject, options uint) GTMioTraceData {
	rv := objc.Send[GTMioTraceData](g.ID, objc.Sel("initWithStreamData:llvmHelperPath:options:"), data, path, options)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/initWithTraceDatabase:deallocator:
func (g GTMioTraceData) InitWithTraceDatabaseDeallocator(database uint64, deallocator VoidHandler) GTMioTraceData {
_block1, _ := NewVoidBlock(deallocator)
	rv := objc.Send[objc.ID](g.ID, objc.Sel("initWithTraceDatabase:deallocator:"), database, _block1)
	return GTMioTraceDataFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/supportsSecureCoding
func (_GTMioTraceDataClass GTMioTraceDataClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_GTMioTraceDataClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/traceDataFromCompressedData:originalSize:compressionAlgorithm:error:
func (_GTMioTraceDataClass GTMioTraceDataClass) TraceDataFromCompressedDataOriginalSizeCompressionAlgorithmError(data objectivec.IObject, size uint64, algorithm int64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_GTMioTraceDataClass.class), objc.Sel("traceDataFromCompressedData:originalSize:compressionAlgorithm:error:"), data, size, algorithm, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/traceDataFromData:error:
func (_GTMioTraceDataClass GTMioTraceDataClass) TraceDataFromDataError(data objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_GTMioTraceDataClass.class), objc.Sel("traceDataFromData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/traceDataFromURL:error:
func (_GTMioTraceDataClass GTMioTraceDataClass) TraceDataFromURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_GTMioTraceDataClass.class), objc.Sel("traceDataFromURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/unarchiverFromData:error:
func (_GTMioTraceDataClass GTMioTraceDataClass) UnarchiverFromDataError(data objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_GTMioTraceDataClass.class), objc.Sel("unarchiverFromData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/LLVMConnectionHeld
func (g GTMioTraceData) LLVMConnectionHeld() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("LLVMConnectionHeld"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/consistentStateAchieved
func (g GTMioTraceData) ConsistentStateAchieved() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("consistentStateAchieved"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/costTimeline
func (g GTMioTraceData) CostTimeline() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("costTimeline"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/gpuInfo
func (g GTMioTraceData) GpuInfo() IGTMioGPUInfo {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gpuInfo"))
	return GTMioGPUInfoFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/gpuTime
func (g GTMioTraceData) GpuTime() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("gpuTime"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/hasSeparateCostsTimeline
func (g GTMioTraceData) HasSeparateCostsTimeline() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("hasSeparateCostsTimeline"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/isMio
func (g GTMioTraceData) IsMio() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isMio"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/loadingCostTimeline
func (g GTMioTraceData) LoadingCostTimeline() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("loadingCostTimeline"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/nonOverlappingCounters
func (g GTMioTraceData) NonOverlappingCounters() IGTMioNonOverlappingCounters {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("nonOverlappingCounters"))
	return GTMioNonOverlappingCountersFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/nonOverlappingTimeline
func (g GTMioTraceData) NonOverlappingTimeline() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("nonOverlappingTimeline"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/overlappingTimeline
func (g GTMioTraceData) OverlappingTimeline() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("overlappingTimeline"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/profiledState
func (g GTMioTraceData) ProfiledState() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("profiledState"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/streamData
func (g GTMioTraceData) StreamData() IGTShaderProfilerStreamData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("streamData"))
	return GTShaderProfilerStreamDataFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/unixTimestamp
func (g GTMioTraceData) UnixTimestamp() int64 {
	rv := objc.Send[int64](g.ID, objc.Sel("unixTimestamp"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/useMinimumTracingMode
func (g GTMioTraceData) UseMinimumTracingMode() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("useMinimumTracingMode"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceData/version
func (g GTMioTraceData) Version() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("version"))
	return rv
}

// EnumerateBinariesForDrawEnumeratorSync is a synchronous wrapper around [GTMioTraceData.EnumerateBinariesForDrawEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceData) EnumerateBinariesForDrawEnumeratorSync(ctx context.Context, draw uint32) error {
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

// EnumerateBinariesForEncoderEnumeratorSync is a synchronous wrapper around [GTMioTraceData.EnumerateBinariesForEncoderEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceData) EnumerateBinariesForEncoderEnumeratorSync(ctx context.Context, encoder uint32) error {
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

// EnumerateBinariesForForCliqueAtIndexUscIndexEnumeratorSync is a synchronous wrapper around [GTMioTraceData.EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceData) EnumerateBinariesForForCliqueAtIndexUscIndexEnumeratorSync(ctx context.Context, index uint32, index2 uint32) error {
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

// EnumerateBinariesForPipelineStateEnumeratorSync is a synchronous wrapper around [GTMioTraceData.EnumerateBinariesForPipelineStateEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceData) EnumerateBinariesForPipelineStateEnumeratorSync(ctx context.Context, state uint64) error {
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

// EnumerateBinaryRangesForCliqueUscDataEnumeratorSync is a synchronous wrapper around [GTMioTraceData.EnumerateBinaryRangesForCliqueUscDataEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceData) EnumerateBinaryRangesForCliqueUscDataEnumeratorSync(ctx context.Context, clique unsafe.Pointer, data objectivec.IObject) error {
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

// EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumeratorSync is a synchronous wrapper around [GTMioTraceData.EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceData) EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumeratorSync(ctx context.Context, index uint32, index2 uint32) error {
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

// EnumerateDrawsForEncoderEnumeratorSync is a synchronous wrapper around [GTMioTraceData.EnumerateDrawsForEncoderEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceData) EnumerateDrawsForEncoderEnumeratorSync(ctx context.Context, encoder uint32) error {
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

// EnumerateDrawsForPipelineStateEnumeratorSync is a synchronous wrapper around [GTMioTraceData.EnumerateDrawsForPipelineStateEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceData) EnumerateDrawsForPipelineStateEnumeratorSync(ctx context.Context, state uint64) error {
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

// EnumerateEncodersSync is a synchronous wrapper around [GTMioTraceData.EnumerateEncoders].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceData) EnumerateEncodersSync(ctx context.Context) error {
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

// EnumerateInstructionsForCliqueUscDataEnumeratorSync is a synchronous wrapper around [GTMioTraceData.EnumerateInstructionsForCliqueUscDataEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceData) EnumerateInstructionsForCliqueUscDataEnumeratorSync(ctx context.Context, clique unsafe.Pointer, data objectivec.IObject) error {
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

// EnumerateInstructionsForCliqueAtIndexUscIndexEnumeratorSync is a synchronous wrapper around [GTMioTraceData.EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceData) EnumerateInstructionsForCliqueAtIndexUscIndexEnumeratorSync(ctx context.Context, index uint32, index2 uint32) error {
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

// EnumerateKickAtFunctionIndexEnumeratorSync is a synchronous wrapper around [GTMioTraceData.EnumerateKickAtFunctionIndexEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceData) EnumerateKickAtFunctionIndexEnumeratorSync(ctx context.Context, index uint32) error {
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

// EnumeratePipelineStatesSync is a synchronous wrapper around [GTMioTraceData.EnumeratePipelineStates].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceData) EnumeratePipelineStatesSync(ctx context.Context) error {
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

// EnumerateUniqueTracesForBinaryEnumeratorSync is a synchronous wrapper around [GTMioTraceData.EnumerateUniqueTracesForBinaryEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceData) EnumerateUniqueTracesForBinaryEnumeratorSync(ctx context.Context, binary uint32) error {
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

// RequestCostTimelineSync is a synchronous wrapper around [GTMioTraceData.RequestCostTimeline].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceData) RequestCostTimelineSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.RequestCostTimeline(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// InitWithTraceDatabaseDeallocatorSync is a synchronous wrapper around [GTMioTraceData.InitWithTraceDatabaseDeallocator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceData) InitWithTraceDatabaseDeallocatorSync(ctx context.Context, database uint64) error {
	done := make(chan struct{}, 1)
	g.InitWithTraceDatabaseDeallocator(database, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}


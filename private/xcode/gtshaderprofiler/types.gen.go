// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// C struct types

// CGColor
type CGColor struct {
}

// CGImage
type CGImage struct {
}

// CallbackHandler
type CallbackHandler struct {
}

// CliqueInstructionInfo
type CliqueInstructionInfo struct {
}

// CounterValuePerDM
type CounterValuePerDM struct {
}

// ExecutionHistoryCacheKey
type ExecutionHistoryCacheKey struct {
	Key    unsafe.Pointer
	Opaque uint64
}

// GPUCommandTimingInfo
type GPUCommandTimingInfo struct {
}

// GRCPerFrameData
type GRCPerFrameData struct {
}

// GTAGX2GRCStreamingSampleHelper
type GTAGX2GRCStreamingSampleHelper struct {
}

// GTAGX2SBALUTargetBlockInfo
type GTAGX2SBALUTargetBlockInfo struct {
}

// GTAGX2ShaderProfilerGRCProcessedUSCSample
type GTAGX2ShaderProfilerGRCProcessedUSCSample struct {
}

// GTAGX2ShaderProfilerProgramAddress
type GTAGX2ShaderProfilerProgramAddress struct {
	Field1 unsafe.Pointer
	Field2 unsafe.Pointer
}

// GTAPSBinaryInfo
type GTAPSBinaryInfo struct {
	Field1  [3]uint64
	Field2  GTRegisterPressureInfo
	Field3  [3]uint64
	Field4  [3]uint64
	Field5  [3]uint64
	Field6  [3]uint64
	Field7  [3]uint64
	Field8  [3]uint64
	Field9  uint32
	Field10 uint32
}

// GTAPSShaderInstructionInfo
type GTAPSShaderInstructionInfo struct {
}

// GTMessageTransportIPC
type GTMessageTransportIPC struct {
	ClientIndexToFileDescriptorMap [3]uint64
	FileDescriptorToClientIndex    [5]uint64
	SocketFileDescriptor           int32
	CallbackHandler                CallbackHandlerRef
	ServerCommunicationSemaphores  *foundation.NSMutableArray
	SemaMutex                      objectivec.NSObject
	SocketName                     unsafe.Pointer
	Listen                         bool
}

// GTMioBinaryTrace
type GTMioBinaryTrace struct {
	Field1 uint64
	Field2 uint64
	Field3 uint64
	Field4 uint32
	Field5 uint32
	Field6 uint16
}

// GTMioCliqueComputePosition
type GTMioCliqueComputePosition struct {
	Field1 uint32
	Field2 uint32
	Field3 uint32
	Field4 uint16
	Field5 uint16
	Field6 uint16
	Field7 uint8
}

// GTMioCliqueFragmentPosition
type GTMioCliqueFragmentPosition struct {
	Field1 uint16
	Field2 uint16
	Field3 uint16
	Field4 [8]uint16
	Field5 [8]uint16
}

// GTMioCostContext
type GTMioCostContext struct {
	Field1 uint16
	Field2 uint16
	Field3 unsafe.Pointer
	Field4 unsafe.Pointer
}

// GTMioCostInfo
type GTMioCostInfo struct {
	Field1  GTMioCostContext
	Field2  float64
	Field3  [10]float64
	Field4  float64
	Field5  [10]float64
	Field6  uint64
	Field7  [10]uint64
	Field8  uint64
	Field9  uint64
	Field10 uint64
}

// GTMioDrawMetadata
type GTMioDrawMetadata struct {
	Field1  uint32
	Field2  uint32
	Field3  uint32
	Field4  uint32
	Field5  int32
	Field6  uint32
	Field7  uint64
	Field8  uint32
	Field9  uint32
	Field10 uint32
}

// GTMioDrawTrace
type GTMioDrawTrace struct {
	Field1 uint64
	Field2 uint64
	Field3 uint32
	Field4 uint16
}

// GTMioEncoderMetadata
type GTMioEncoderMetadata struct {
	Field1 [4]float64
	Field2 uint32
	Field3 uint32
	Field4 uint16
	Field5 uint16
	Field6 uint32
	Field7 uint16
}

// GTMioGPUInfoInternal
type GTMioGPUInfoInternal struct {
	Field1 uint32
	Field2 uint32
	Field3 uint32
	Field4 uint32
	Field5 uint32
	Field6 uint32
}

// GTMioHeatmapBuilderGenerationOptions
type GTMioHeatmapBuilderGenerationOptions struct {
	Field1  uint64
	Field2  uint64
	Field3  uint64
	Field4  uint64
	Field5  uint64
	Field6  uint64
	Field7  uint32
	Field8  uint32
	Field9  uint64
	Field10 uint64
	Field11 bool
}

// GTMioKVDataBlock
type GTMioKVDataBlock struct {
	Field1 uint32
	Field2 uint32
	Field3 uint64
	Field4 uint64
	Field5 uint32
	Field6 uint64
	Field7 uint64
}

// GTMioKVDataReader
type GTMioKVDataReader struct {
	Data             *byte
	Length           uint64
	Blocks           [3]uint64
	NameToBlockIndex *foundation.NSDictionary
	Meta             *foundation.NSDictionary
	Header           unsafe.Pointer
}

// GTMioKickTrace
type GTMioKickTrace struct {
	Field1  uint64
	Field2  uint64
	Field3  uint32
	Field4  uint32
	Field5  uint32
	Field6  uint32
	Field7  uint32
	Field8  uint32
	Field9  uint16
	Field10 uint16
	Field11 uint16
}

// GTMioMetalFXInfo
type GTMioMetalFXInfo struct {
	Field1 uint64
	Field2 uint64
}

// GTMioNonOverlappingCounterContainerInternal
type GTMioNonOverlappingCounterContainerInternal struct {
	Guard [8]uint64
}

// GTMioNonOverlappingCountersInternal
type GTMioNonOverlappingCountersInternal struct {
	EncoderCounters    GTMioNonOverlappingCounterContainerInternal
	GpuCommandCounters GTMioNonOverlappingCounterContainerInternal
}

// GTMioOrderedTrackAssigner
type GTMioOrderedTrackAssigner struct {
	TrackStates [3]uint64
}

// GTMioPerEncoderShaderTrackStore
type GTMioPerEncoderShaderTrackStore struct {
	Aggregators [5]uint64
}

// GTMioQuadLocation
type GTMioQuadLocation struct {
	Field1 unsafe.Pointer
}

// GTMioRIATrace
type GTMioRIATrace struct {
	Field1 uint64
	Field2 uint64
	Field3 uint16
}

// GTMioShaderBinaryDebugBinaryRange
type GTMioShaderBinaryDebugBinaryRange struct {
	Field1 uint32
	Field2 uint64
	Field3 uint64
	Field4 uint32
	Field5 uint32
}

// GTMioShaderBinaryDebugLocation
type GTMioShaderBinaryDebugLocation struct {
	Field1 uint32
	Field2 uint32
	Field3 uint32
	Field4 uint32
}

// GTMioShaderBinaryInfo
type GTMioShaderBinaryInfo struct {
	Field1 uint32
	Field2 uint32
	Field3 uint64
	Field4 uint64
	Field5 uint16
	Field6 uint16
	Field7 uint16
}

// GTMioShaderInstructionInfo
type GTMioShaderInstructionInfo struct {
	Field1  uint32
	Field2  uint32
	Field3  uint32
	Field4  uint32
	Field5  uint16
	Field6  uint16
	Field7  uint16
	Field8  uint16
	Field9  uint16
	Field10 uint16
}

// GTMioShaderTimelineSignpostPipelineState
type GTMioShaderTimelineSignpostPipelineState struct {
	Field1 uint32
	Field2 uint32
}

// GTMioShaderTimelineSignpostProcess
type GTMioShaderTimelineSignpostProcess struct {
	Field1 uint32
	Field2 uint64
}

// GTMioShaderTimelineSignpostShader
type GTMioShaderTimelineSignpostShader struct {
	Field1 uint32
	Field2 uint32
	Field3 uint32
	Field4 uint64
	Field5 uint64
	Field6 uint64
	Field7 uint32
	Field8 uint32
}

// GTMioTraceDataBuilderOptions
type GTMioTraceDataBuilderOptions struct {
	Field1 bool
	Field2 bool
	Field3 bool
	Field4 bool
}

// GTMioTraceDataShaderStatCollectorInternal
type GTMioTraceDataShaderStatCollectorInternal struct {
	PipelineIdToShaderStat [5]uint64
}

// GTMioTraceDataShaderStatInternal
type GTMioTraceDataShaderStatInternal struct {
	Field1 uint64
	Field2 uint64
	Field3 uint64
}

// GTMioUSCCliqueIndex
type GTMioUSCCliqueIndex struct {
	Field1 unsafe.Pointer
}

// GTMioUSCCliqueMetadata
type GTMioUSCCliqueMetadata struct {
	Field1  uint64
	Field2  uint64
	Field3  uint32
	Field4  uint32
	Field5  uint32
	Field6  uint32
	Field7  uint32
	Field8  uint32
	Field9  uint32
	Field10 uint32
	Field11 uint32
	Field12 uint32
	Field13 uint32
	Field14 uint32
	Field15 uint32
	Field16 uint32
	Field17 uint16
	Field18 uint16
	Field19 uint16
}

// GTMioUSCInstructionTraceInfo
type GTMioUSCInstructionTraceInfo struct {
	Field1 uint64
	Field2 uint32
	Field3 uint64
	Field4 uint32
	Field5 uint32
	Field6 uint32
	Field7 uint32
	Field8 uint64
	Field9 uint64
}

// GTMioUSCInstructionTraceTrackRecord
type GTMioUSCInstructionTraceTrackRecord struct {
	Field1 uint64
	Field2 uint64
	Field3 uint32
	Field4 uint32
	Field5 uint64
	Field6 uint64
}

// GTMioUSCInstructionTraceTrackTrace
type GTMioUSCInstructionTraceTrackTrace struct {
	Field1 uint64
	Field2 uint32
	Field3 uint32
	Field4 uint32
	Field5 uint16
	Field6 uint8
	Field7 uint8
	Field8 uint32
	Field9 uint16
}

// GTMioUSCKickMetadata
type GTMioUSCKickMetadata struct {
	Field1  uint64
	Field2  uint64
	Field3  uint64
	Field4  uint64
	Field5  uint32
	Field6  uint32
	Field7  uint32
	Field8  uint32
	Field9  uint32
	Field10 uint32
	Field11 uint32
	Field12 uint32
	Field13 uint32
	Field14 uint32
	Field15 uint16
	Field16 uint16
	Field17 uint16
	Field18 uint16
}

// GTMioUSCTileMetadata
type GTMioUSCTileMetadata struct {
	Field1  uint64
	Field2  uint64
	Field3  uint64
	Field4  uint64
	Field5  uint32
	Field6  uint32
	Field7  uint32
	Field8  uint32
	Field9  uint32
	Field10 uint16
	Field11 uint16
}

// GTRegisterPressureInfo
type GTRegisterPressureInfo struct {
	Field1 uint32
	Field2 [3]uint64
}

// GTRegisterPressureInstructionInfo
type GTRegisterPressureInstructionInfo struct {
}

// GTShaderProfilerBinaryInfo
type GTShaderProfilerBinaryInfo struct {
	Field1 [3]uint64
	Field2 GTRegisterPressureInfo
	Field3 [3]uint64
	Field4 [3]uint64
	Field5 [3]uint64
	Field6 [3]uint64
	Field7 [3]uint64
	Field8 [3]uint64
	Field9 [3]uint64
}

// GTWaitInstructionInfo
type GTWaitInstructionInfo struct {
}

// InstructionPCStatInfo
type InstructionPCStatInfo struct {
	Field1 ShaderBinaryStatsInfoRef
	Field2 uint32
	Field3 uint32
	Field4 objectivec.Object
	Field5 objectivec.Object
	Field6 [4]int8
	Field7 float64
	Field8 float64
	Field9 [3]uint64
}

// KickShaders
type KickShaders struct {
}

// MCAOutput
type MCAOutput struct {
	Field1 objectivec.Object
	Field2 objectivec.Object
}

// OpaqueJSString
type OpaqueJSString struct {
}

// OpaqueJSValue
type OpaqueJSValue struct {
}

// ProcessedCliqueTraceData
type ProcessedCliqueTraceData struct {
}

// QuadOrder
type QuadOrder struct {
}

// SampleNormFactor
type SampleNormFactor struct {
	Field1 objectivec.Object
	Field2 objectivec.Object
	Field3 objectivec.Object
	Field4 objectivec.Object
	Field5 objectivec.Object
}

// ShaderBinaryStatsInfo
type ShaderBinaryStatsInfo struct {
}

// ShaderProfilerUSCSampleInfo
type ShaderProfilerUSCSampleInfo struct {
	Field1 uint32
	Field2 uint32
	Field3 uint32
	Field4 uint32
	Field5 uint32
	Field6 unsafe.Pointer
	Field7 unsafe.Pointer
	Field8 uint32
	Field9 uint32
}

// Statistics
type Statistics struct {
	Field1 float64
	Field2 float64
	Field3 float64
}

// TimeStats
type TimeStats struct {
}

// XRGPUAGXShaderTimelineSignpostProcess
type XRGPUAGXShaderTimelineSignpostProcess struct {
}

// XRGPUAPSCounterContainer
type XRGPUAPSCounterContainer struct {
	ApsTimestamps               [3]uint64
	ApsRawCounterValues         [3]uint64
	ApsDerivedValues            [3]uint64
	RawCounterNames             *foundation.NSMutableArray
	DeobfuscatedRawCounterNames *foundation.NSMutableArray
	DerivedCounters             *foundation.NSMutableArray
	ProfileData                 [1]uint64
	SystemTimestamps            [3]uint64
	UscTimestamps               [3]uint64
	SyncTimestamps              [3]uint64
	KickMap                     [5]uint64
	KickTraceIds                [3]uint64
	GlobalTraceIds              [3]uint64
	KickShaders                 [3]uint64
}

// XRGPUInterval
type XRGPUInterval struct {
}

// XRGPURDECounterSource
type XRGPURDECounterSource struct {
}

// Long
type Long struct {
	__data_    *byte
	__size_    uint64
	__cap_     objectivec.Object
	__is_long_ objectivec.Object
}

// Rep
type Rep struct {
	__s int16
	__l int
}

// SFILE
type SFILE struct {
	Field1  *byte
	Field2  int32
	Field3  int32
	Field4  int16
	Field5  int16
	Field6  Sbuf
	Field7  int32
	Field8  unsafe.Pointer
	Field9  unsafe.Pointer
	Field10 unsafe.Pointer
	Field11 unsafe.Pointer
	Field12 unsafe.Pointer
	Field13 Sbuf
	Field14 SFILEXRef
	Field15 int32
	Field16 [3]uint8
	Field17 [1]uint8
	Field18 Sbuf
	Field19 int32
	Field20 int64
}

// SFILEX
type SFILEX struct {
}

// Sbuf
type Sbuf struct {
	Field1 *byte
	Field2 int32
}

// Short
type Short struct {
	__data_    [23]int8
	__size_    objectivec.Object
	__is_long_ objectivec.Object
}

// OpaquePthreadMutex
type OpaquePthreadMutex struct {
	__sig    int64
	__opaque [56]int8
}

// Opaque_pthread_mutex_t is a type alias for OpaquePthreadMutex for use in objc.Send[T] calls.
type Opaque_pthread_mutex_t = OpaquePthreadMutex

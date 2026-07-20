// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objectivec"
)

// C struct types

// CGColor
type CGColor struct {
}

// CGImage
type CGImage struct {
}

// ExecutionHistoryCacheKey
type ExecutionHistoryCacheKey struct {
	Key    unsafe.Pointer
	Opaque uint64
}

// GTAGX2ShaderProfilerProgramAddress
type GTAGX2ShaderProfilerProgramAddress struct {
	Field1 unsafe.Pointer
	Field2 unsafe.Pointer
}

// GTAPSBinaryInfo
type GTAPSBinaryInfo struct {
	Field1  unsafe.Pointer
	Field2  unsafe.Pointer
	Field3  unsafe.Pointer
	Field4  unsafe.Pointer
	Field5  unsafe.Pointer
	Field6  unsafe.Pointer
	Field7  unsafe.Pointer
	Field8  unsafe.Pointer
	Field9  uint
	Field10 uint
}

// GTMessageTransportIPC
type GTMessageTransportIPC struct {
	ClientIndexToFileDescriptorMap unsafe.Pointer
	FileDescriptorToClientIndex    unsafe.Pointer
	SocketFileDescriptor           int
	CallbackHandler                CallbackHandlerRef
	ServerCommunicationSemaphores  foundation.NSMutableArray
	SemaMutex                      objectivec.NSObject
	SocketName                     unsafe.Pointer
	Listen                         bool
}

// GTMioBinaryTrace
type GTMioBinaryTrace struct {
	Field1 uint64
	Field2 uint64
	Field3 uint64
	Field4 uint
	Field5 uint
	Field6 uint16
}

// GTMioCliqueComputePosition
type GTMioCliqueComputePosition struct {
	Field1 uint
	Field2 uint
	Field3 uint
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
	Field4 unsafe.Pointer
	Field5 unsafe.Pointer
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
	Field3  unsafe.Pointer
	Field4  float64
	Field5  unsafe.Pointer
	Field6  uint64
	Field7  unsafe.Pointer
	Field8  uint64
	Field9  uint64
	Field10 uint64
}

// GTMioDrawMetadata
type GTMioDrawMetadata struct {
	Field1  uint
	Field2  uint
	Field3  uint
	Field4  uint
	Field5  int
	Field6  uint
	Field7  uint64
	Field8  uint
	Field9  uint
	Field10 uint
}

// GTMioDrawTrace
type GTMioDrawTrace struct {
	Field1 uint64
	Field2 uint64
	Field3 uint
	Field4 uint16
}

// GTMioEncoderMetadata
type GTMioEncoderMetadata struct {
	Field1 unsafe.Pointer
	Field2 uint
	Field3 uint
	Field4 uint16
	Field5 uint16
	Field6 uint
	Field7 uint16
}

// GTMioGPUInfoInternal
type GTMioGPUInfoInternal struct {
	Field1 uint
	Field2 uint
	Field3 uint
	Field4 uint
	Field5 uint
	Field6 uint
}

// GTMioHeatmapBuilderGenerationOptions
type GTMioHeatmapBuilderGenerationOptions struct {
	Field1  uint64
	Field2  uint64
	Field3  uint64
	Field4  uint64
	Field5  uint64
	Field6  uint64
	Field7  uint
	Field8  uint
	Field9  uint64
	Field10 uint64
	Field11 bool
}

// GTMioKVDataBlock
type GTMioKVDataBlock struct {
	Field1 uint
	Field2 uint
	Field3 uint64
	Field4 uint64
	Field5 uint
	Field6 uint64
	Field7 uint64
}

// GTMioKVDataReader
type GTMioKVDataReader struct {
	Data             *byte
	Length           uint64
	Blocks           unsafe.Pointer
	NameToBlockIndex foundation.NSDictionary
	Meta             foundation.NSDictionary
	Header           unsafe.Pointer
}

// GTMioKickTrace
type GTMioKickTrace struct {
	Field1  uint64
	Field2  uint64
	Field3  uint
	Field4  uint
	Field5  uint
	Field6  uint
	Field7  uint
	Field8  uint
	Field9  uint16
	Field10 uint16
	Field11 uint16
}

// GTMioMetalFXInfo
type GTMioMetalFXInfo struct {
	Field1 uint64
	Field2 uint64
}

// GTMioNonOverlappingCountersInternal
type GTMioNonOverlappingCountersInternal struct {
	EncoderCounters    unsafe.Pointer
	GpuCommandCounters unsafe.Pointer
}

// GTMioOrderedTrackAssigner
type GTMioOrderedTrackAssigner struct {
	TrackStates unsafe.Pointer
}

// GTMioPerEncoderShaderTrackStore
type GTMioPerEncoderShaderTrackStore struct {
	Aggregators unsafe.Pointer
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
	Field1 uint
	Field2 uint64
	Field3 uint64
	Field4 uint
	Field5 uint
}

// GTMioShaderBinaryDebugLocation
type GTMioShaderBinaryDebugLocation struct {
	Field1 uint
	Field2 uint
	Field3 uint
	Field4 uint
}

// GTMioShaderBinaryInfo
type GTMioShaderBinaryInfo struct {
	Field1 uint
	Field2 uint
	Field3 uint64
	Field4 uint64
	Field5 uint16
	Field6 uint16
	Field7 uint16
}

// GTMioShaderInstructionInfo
type GTMioShaderInstructionInfo struct {
	Field1  uint
	Field2  uint
	Field3  uint
	Field4  uint
	Field5  uint16
	Field6  uint16
	Field7  uint16
	Field8  uint16
	Field9  uint16
	Field10 uint16
}

// GTMioShaderTimelineSignpostPipelineState
type GTMioShaderTimelineSignpostPipelineState struct {
	Field1 uint
	Field2 uint
}

// GTMioShaderTimelineSignpostProcess
type GTMioShaderTimelineSignpostProcess struct {
	Field1 uint
	Field2 uint64
}

// GTMioShaderTimelineSignpostShader
type GTMioShaderTimelineSignpostShader struct {
	Field1 uint
	Field2 uint
	Field3 uint
	Field4 uint64
	Field5 uint64
	Field6 uint64
	Field7 uint
	Field8 uint
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
	PipelineIdToShaderStat unsafe.Pointer
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
	Field3  uint
	Field4  uint
	Field5  uint
	Field6  uint
	Field7  uint
	Field8  uint
	Field9  uint
	Field10 uint
	Field11 uint
	Field12 uint
	Field13 uint
	Field14 uint
	Field15 uint
	Field16 uint
	Field17 uint16
	Field18 uint16
	Field19 uint16
}

// GTMioUSCInstructionTraceInfo
type GTMioUSCInstructionTraceInfo struct {
	Field1 uint64
	Field2 uint
	Field3 uint64
	Field4 uint
	Field5 uint
	Field6 uint
	Field7 uint
	Field8 uint64
	Field9 uint64
}

// GTMioUSCInstructionTraceTrackRecord
type GTMioUSCInstructionTraceTrackRecord struct {
	Field1 uint64
	Field2 uint64
	Field3 uint
	Field4 uint
	Field5 uint64
	Field6 uint64
}

// GTMioUSCInstructionTraceTrackTrace
type GTMioUSCInstructionTraceTrackTrace struct {
	Field1 uint64
	Field2 uint
	Field3 uint
	Field4 uint
	Field5 uint16
	Field6 uint8
	Field7 uint8
	Field8 uint
	Field9 uint16
}

// GTMioUSCKickMetadata
type GTMioUSCKickMetadata struct {
	Field1  uint64
	Field2  uint64
	Field3  uint64
	Field4  uint64
	Field5  uint
	Field6  uint
	Field7  uint
	Field8  uint
	Field9  uint
	Field10 uint
	Field11 uint
	Field12 uint
	Field13 uint
	Field14 uint
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
	Field5  uint
	Field6  uint
	Field7  uint
	Field8  uint
	Field9  uint
	Field10 uint16
	Field11 uint16
}

// GTShaderProfilerBinaryInfo
type GTShaderProfilerBinaryInfo struct {
	Field1 unsafe.Pointer
	Field2 unsafe.Pointer
	Field3 unsafe.Pointer
	Field4 unsafe.Pointer
	Field5 unsafe.Pointer
	Field6 unsafe.Pointer
	Field7 unsafe.Pointer
	Field8 unsafe.Pointer
	Field9 unsafe.Pointer
}

// InstructionPCStatInfo
type InstructionPCStatInfo struct {
	Field1 ShaderBinaryStatsInfoRef
	Field2 uint
	Field3 uint
	Field4 objectivec.Object
	Field5 objectivec.Object
	Field6 unsafe.Pointer
	Field7 float64
	Field8 float64
	Field9 unsafe.Pointer
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

// ShaderProfilerUSCSampleInfo
type ShaderProfilerUSCSampleInfo struct {
	Field1 uint
	Field2 uint
	Field3 uint
	Field4 uint
	Field5 uint
	Field6 unsafe.Pointer
	Field7 unsafe.Pointer
	Field8 uint
	Field9 uint
}

// XRGPUAGXShaderTimelineSignpostProcess
type XRGPUAGXShaderTimelineSignpostProcess struct {
}

// XRGPUAPSCounterContainer
type XRGPUAPSCounterContainer struct {
	ApsTimestamps               unsafe.Pointer
	ApsRawCounterValues         unsafe.Pointer
	ApsDerivedValues            unsafe.Pointer
	RawCounterNames             foundation.NSMutableArray
	DeobfuscatedRawCounterNames foundation.NSMutableArray
	DerivedCounters             foundation.NSMutableArray
	ProfileData                 unsafe.Pointer
	SystemTimestamps            unsafe.Pointer
	UscTimestamps               unsafe.Pointer
	SyncTimestamps              unsafe.Pointer
	KickMap                     unsafe.Pointer
	KickTraceIds                unsafe.Pointer
	GlobalTraceIds              unsafe.Pointer
	KickShaders                 unsafe.Pointer
}

// XRGPUInterval
type XRGPUInterval struct {
}

// SFILE
type SFILE struct {
	Field1  *byte
	Field2  int
	Field3  int
	Field4  int16
	Field5  int16
	Field6  kernel.Sbuf
	Field7  int
	Field8  unsafe.Pointer
	Field9  unsafe.Pointer
	Field10 unsafe.Pointer
	Field11 unsafe.Pointer
	Field12 unsafe.Pointer
	Field13 kernel.Sbuf
	Field14 SFILEXRef
	Field15 int
	Field16 unsafe.Pointer
	Field17 unsafe.Pointer
	Field18 kernel.Sbuf
	Field19 int
	Field20 int64
}

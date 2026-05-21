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

// GTAGX2ShaderProfilerProgramAddress
type GTAGX2ShaderProfilerProgramAddress struct {
}

// GTAPSBinaryInfo
type GTAPSBinaryInfo struct {
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
}

// GTMioCliqueComputePosition
type GTMioCliqueComputePosition struct {
}

// GTMioCliqueFragmentPosition
type GTMioCliqueFragmentPosition struct {
}

// GTMioCostContext
type GTMioCostContext struct {
}

// GTMioCostInfo
type GTMioCostInfo struct {
}

// GTMioDrawMetadata
type GTMioDrawMetadata struct {
}

// GTMioDrawTrace
type GTMioDrawTrace struct {
}

// GTMioEncoderMetadata
type GTMioEncoderMetadata struct {
}

// GTMioGPUInfoInternal
type GTMioGPUInfoInternal struct {
}

// GTMioHeatmapBuilderGenerationOptions
type GTMioHeatmapBuilderGenerationOptions struct {
}

// GTMioKVDataBlock
type GTMioKVDataBlock struct {
}

// GTMioKVDataReader
type GTMioKVDataReader struct {
	Data             *byte
	Length           uint64
	Blocks           unsafe.Pointer
	NameToBlockIndex foundation.NSDictionary
	Meta             foundation.NSDictionary
	Header           kernel.Pointer
}

// GTMioKickTrace
type GTMioKickTrace struct {
}

// GTMioMetalFXInfo
type GTMioMetalFXInfo struct {
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
}

// GTMioRIATrace
type GTMioRIATrace struct {
}

// GTMioShaderBinaryDebugBinaryRange
type GTMioShaderBinaryDebugBinaryRange struct {
}

// GTMioShaderBinaryDebugLocation
type GTMioShaderBinaryDebugLocation struct {
}

// GTMioShaderBinaryInfo
type GTMioShaderBinaryInfo struct {
}

// GTMioShaderInstructionInfo
type GTMioShaderInstructionInfo struct {
}

// GTMioShaderTimelineSignpostPipelineState
type GTMioShaderTimelineSignpostPipelineState struct {
}

// GTMioShaderTimelineSignpostProcess
type GTMioShaderTimelineSignpostProcess struct {
}

// GTMioShaderTimelineSignpostShader
type GTMioShaderTimelineSignpostShader struct {
}

// GTMioTraceDataBuilderOptions
type GTMioTraceDataBuilderOptions struct {
}

// GTMioTraceDataShaderStatCollectorInternal
type GTMioTraceDataShaderStatCollectorInternal struct {
	PipelineIdToShaderStat unsafe.Pointer
}

// GTMioTraceDataShaderStatInternal
type GTMioTraceDataShaderStatInternal struct {
}

// GTMioUSCCliqueIndex
type GTMioUSCCliqueIndex struct {
}

// GTMioUSCCliqueMetadata
type GTMioUSCCliqueMetadata struct {
}

// GTMioUSCInstructionTraceInfo
type GTMioUSCInstructionTraceInfo struct {
}

// GTMioUSCInstructionTraceTrackRecord
type GTMioUSCInstructionTraceTrackRecord struct {
}

// GTMioUSCInstructionTraceTrackTrace
type GTMioUSCInstructionTraceTrackTrace struct {
}

// GTMioUSCKickMetadata
type GTMioUSCKickMetadata struct {
}

// GTMioUSCTileMetadata
type GTMioUSCTileMetadata struct {
}

// GTShaderProfilerBinaryInfo
type GTShaderProfilerBinaryInfo struct {
}

// InstructionPCStatInfo
type InstructionPCStatInfo struct {
}

// MCAOutput
type MCAOutput struct {
}

// OpaqueJSString
type OpaqueJSString struct {
}

// OpaqueJSValue
type OpaqueJSValue struct {
}

// ShaderProfilerUSCSampleInfo
type ShaderProfilerUSCSampleInfo struct {
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
}

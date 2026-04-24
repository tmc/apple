// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// C struct types

// CGColor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/CGColor
type CGColor struct {
}

// CGImage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/CGImage
type CGImage struct {
}

// GTAGX2ShaderProfilerProgramAddress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerProgramAddress
type GTAGX2ShaderProfilerProgramAddress struct {
}

// GTAPSBinaryInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTAPSBinaryInfo
type GTAPSBinaryInfo struct {
}

// GTMessageTransportIPC
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMessageTransportIPC
type GTMessageTransportIPC struct {
	ClientIndexToFileDescriptorMap unsafe.Pointer
	FileDescriptorToClientIndex    unsafe.Pointer
	SocketFileDescriptor           int
	CallbackHandler                CallbackHandlerRef
	ServerCommunicationSemaphores  *foundation.NSMutableArray
	SemaMutex                      objectivec.NSObject
	SocketName                     unsafe.Pointer
	Listen                         bool
}

// GTMioBinaryTrace
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioBinaryTrace
type GTMioBinaryTrace struct {
}

// GTMioCliqueComputePosition
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCliqueComputePosition
type GTMioCliqueComputePosition struct {
}

// GTMioCliqueFragmentPosition
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCliqueFragmentPosition
type GTMioCliqueFragmentPosition struct {
}

// GTMioCostContext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCostContext
type GTMioCostContext struct {
}

// GTMioCostInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCostInfo
type GTMioCostInfo struct {
}

// GTMioDrawMetadata
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioDrawMetadata
type GTMioDrawMetadata struct {
}

// GTMioDrawTrace
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioDrawTrace
type GTMioDrawTrace struct {
}

// GTMioEncoderMetadata
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderMetadata
type GTMioEncoderMetadata struct {
}

// GTMioGPUInfoInternal
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioGPUInfoInternal
type GTMioGPUInfoInternal struct {
}

// GTMioHeatmapBuilderGenerationOptions
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilderGenerationOptions
type GTMioHeatmapBuilderGenerationOptions struct {
}

// GTMioKVDataBlock
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataBlock
type GTMioKVDataBlock struct {
}

// GTMioKVDataReader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKVDataReader
type GTMioKVDataReader struct {
	Data             *byte
	Length           uint64
	Blocks           unsafe.Pointer
	NameToBlockIndex *foundation.NSDictionary
	Meta             *foundation.NSDictionary
	Header           objectivec.IObject
}

// GTMioKickTrace
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioKickTrace
type GTMioKickTrace struct {
}

// GTMioMetalFXInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioMetalFXInfo
type GTMioMetalFXInfo struct {
}

// GTMioNonOverlappingCountersInternal
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioNonOverlappingCountersInternal
type GTMioNonOverlappingCountersInternal struct {
	EncoderCounters    unsafe.Pointer
	GpuCommandCounters unsafe.Pointer
}

// GTMioOrderedTrackAssigner
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioOrderedTrackAssigner
type GTMioOrderedTrackAssigner struct {
	TrackStates unsafe.Pointer
}

// GTMioPerEncoderShaderTrackStore
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioPerEncoderShaderTrackStore
type GTMioPerEncoderShaderTrackStore struct {
	Aggregators unsafe.Pointer
}

// GTMioQuadLocation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioQuadLocation
type GTMioQuadLocation struct {
}

// GTMioRIATrace
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioRIATrace
type GTMioRIATrace struct {
}

// GTMioShaderBinaryDebugBinaryRange
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderBinaryDebugBinaryRange
type GTMioShaderBinaryDebugBinaryRange struct {
}

// GTMioShaderBinaryDebugLocation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderBinaryDebugLocation
type GTMioShaderBinaryDebugLocation struct {
}

// GTMioShaderBinaryInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderBinaryInfo
type GTMioShaderBinaryInfo struct {
}

// GTMioShaderInstructionInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderInstructionInfo
type GTMioShaderInstructionInfo struct {
}

// GTMioShaderTimelineSignpostPipelineState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderTimelineSignpostPipelineState
type GTMioShaderTimelineSignpostPipelineState struct {
}

// GTMioShaderTimelineSignpostProcess
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderTimelineSignpostProcess
type GTMioShaderTimelineSignpostProcess struct {
}

// GTMioShaderTimelineSignpostShader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderTimelineSignpostShader
type GTMioShaderTimelineSignpostShader struct {
}

// GTMioTraceDataBuilderOptions
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataBuilderOptions
type GTMioTraceDataBuilderOptions struct {
}

// GTMioTraceDataShaderStatCollectorInternal
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataShaderStatCollectorInternal
type GTMioTraceDataShaderStatCollectorInternal struct {
	PipelineIdToShaderStat unsafe.Pointer
}

// GTMioTraceDataShaderStatInternal
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataShaderStatInternal
type GTMioTraceDataShaderStatInternal struct {
}

// GTMioUSCCliqueIndex
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioUSCCliqueIndex
type GTMioUSCCliqueIndex struct {
}

// GTMioUSCCliqueMetadata
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioUSCCliqueMetadata
type GTMioUSCCliqueMetadata struct {
}

// GTMioUSCInstructionTraceInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioUSCInstructionTraceInfo
type GTMioUSCInstructionTraceInfo struct {
}

// GTMioUSCInstructionTraceTrackRecord
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioUSCInstructionTraceTrackRecord
type GTMioUSCInstructionTraceTrackRecord struct {
}

// GTMioUSCInstructionTraceTrackTrace
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioUSCInstructionTraceTrackTrace
type GTMioUSCInstructionTraceTrackTrace struct {
}

// GTMioUSCKickMetadata
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioUSCKickMetadata
type GTMioUSCKickMetadata struct {
}

// GTMioUSCTileMetadata
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTMioUSCTileMetadata
type GTMioUSCTileMetadata struct {
}

// GTShaderProfilerBinaryInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerBinaryInfo
type GTShaderProfilerBinaryInfo struct {
}

// InstructionPCStatInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/InstructionPCStatInfo
type InstructionPCStatInfo struct {
}

// MCAOutput
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/MCAOutput
type MCAOutput struct {
}

// OpaqueJSString
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/OpaqueJSString
type OpaqueJSString struct {
}

// OpaqueJSValue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/OpaqueJSValue
type OpaqueJSValue struct {
}

// ShaderProfilerUSCSampleInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/ShaderProfilerUSCSampleInfo
type ShaderProfilerUSCSampleInfo struct {
}

// XRGPUAGXShaderTimelineSignpostProcess
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAGXShaderTimelineSignpostProcess
type XRGPUAGXShaderTimelineSignpostProcess struct {
}

// XRGPUAPSCounterContainer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSCounterContainer
type XRGPUAPSCounterContainer struct {
	ApsTimestamps               unsafe.Pointer
	ApsRawCounterValues         unsafe.Pointer
	ApsDerivedValues            unsafe.Pointer
	RawCounterNames             *foundation.NSMutableArray
	DeobfuscatedRawCounterNames *foundation.NSMutableArray
	DerivedCounters             *foundation.NSMutableArray
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
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUInterval
type XRGPUInterval struct {
}

// _sFILE
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GTShaderProfiler/__sFILE
type _sFILE struct {
}

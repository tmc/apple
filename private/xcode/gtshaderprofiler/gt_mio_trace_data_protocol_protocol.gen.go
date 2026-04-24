// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTMioTraceDataProtocol protocol.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol
type GTMioTraceDataProtocol interface {
	objectivec.IObject

	// ChildCliqueOfClique protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/childCliqueOfClique:
	ChildCliqueOfClique(clique *GTMioUSCCliqueMetadataRef) unsafe.Pointer

	// CliqueFromCliqueIndex protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/cliqueFromCliqueIndex:
	CliqueFromCliqueIndex(index unsafe.Pointer) unsafe.Pointer

	// CoalescedFunctionIndexForEncoderFunctionIndex protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/coalescedFunctionIndexForEncoderFunctionIndex:
	CoalescedFunctionIndexForEncoderFunctionIndex(index uint32) uint32

	// ComputePositionCount protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/computePositionCount
	ComputePositionCount() uint64

	// ComputePositions protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/computePositions
	ComputePositions() unsafe.Pointer

	// ConsistentStateAchieved protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/consistentStateAchieved
	ConsistentStateAchieved() bool

	// CostForContextCost protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/costForContext:cost:
	CostForContextCost(context unsafe.Pointer, cost unsafe.Pointer) bool

	// CostForLevelLevelIdentifierScopeScopeIdentifierCost protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/costForLevel:levelIdentifier:scope:scopeIdentifier:cost:
	CostForLevelLevelIdentifierScopeScopeIdentifierCost(level uint16, identifier uint32, scope uint16, identifier2 uint64, cost unsafe.Pointer) bool

	// DataType protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/dataType
	DataType() uint32

	// DatabaseInternal protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/databaseInternal
	DatabaseInternal() uint64

	// DrawCount protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/drawCount
	DrawCount() uint64

	// DrawTraceCount protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/drawTraceCount
	DrawTraceCount() uint64

	// DrawTraces protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/drawTraces
	DrawTraces() unsafe.Pointer

	// Draws protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/draws
	Draws() unsafe.Pointer

	// DurationForDrawDataMaster protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/durationForDraw:dataMaster:
	DurationForDrawDataMaster(draw uint32, master uint16) uint64

	// EncoderCount protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/encoderCount
	EncoderCount() uint64

	// EncoderFromFunctionIndex protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/encoderFromFunctionIndex:
	EncoderFromFunctionIndex(index uint32) unsafe.Pointer

	// Encoders protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/encoders
	Encoders() unsafe.Pointer

	// EnumerateBinariesForDrawEnumerator protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateBinariesForDraw:enumerator:
	EnumerateBinariesForDrawEnumerator(draw uint32, enumerator GTMioShaderBinaryDataHandler)

	// EnumerateBinariesForEncoderEnumerator protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateBinariesForEncoder:enumerator:
	EnumerateBinariesForEncoderEnumerator(encoder uint32, enumerator GTMioShaderBinaryDataHandler)

	// EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateBinariesForForCliqueAtIndex:uscIndex:enumerator:
	EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator GTMioShaderBinaryDataHandler)

	// EnumerateBinariesForPipelineStateEnumerator protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateBinariesForPipelineState:enumerator:
	EnumerateBinariesForPipelineStateEnumerator(state uint64, enumerator GTMioShaderBinaryDataHandler)

	// EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateBinaryRangesForCliqueAtIndex:uscIndex:enumerator:
	EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator GTMioShaderBinaryDebugBinaryRangeHandler)

	// EnumerateDrawsForEncoderEnumerator protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateDrawsForEncoder:enumerator:
	EnumerateDrawsForEncoderEnumerator(encoder uint32, enumerator UintHandler)

	// EnumerateDrawsForPipelineStateEnumerator protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateDrawsForPipelineState:enumerator:
	EnumerateDrawsForPipelineStateEnumerator(state uint64, enumerator UintHandler)

	// EnumerateEncoders protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateEncoders:
	EnumerateEncoders(encoders UintHandler)

	// EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateInstructionsForCliqueAtIndex:uscIndex:enumerator:
	EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator GTMioShaderInstructionInfoHandler)

	// EnumerateKickAtFunctionIndexEnumerator protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateKickAtFunctionIndex:enumerator:
	EnumerateKickAtFunctionIndexEnumerator(index uint32, enumerator GTMioUSCKickMetadataHandler)

	// EnumeratePipelineStates protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumeratePipelineStates:
	EnumeratePipelineStates(states unsignedlongHandler)

	// EnumerateUniqueTracesForBinaryEnumerator protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateUniqueTracesForBinary:enumerator:
	EnumerateUniqueTracesForBinaryEnumerator(binary uint32, enumerator GTMioUSCTraceDataHandler)

	// FailedUSCIndexCount protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/failedUSCIndexCount
	FailedUSCIndexCount() uint64

	// FailedUSCIndexes protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/failedUSCIndexes
	FailedUSCIndexes() unsafe.Pointer

	// FragmentPositionCount protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/fragmentPositionCount
	FragmentPositionCount() uint64

	// FragmentPositions protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/fragmentPositions
	FragmentPositions() unsafe.Pointer

	// GlobalGPUTime protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/globalGPUTime
	GlobalGPUTime() uint64

	// GpuCost protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/gpuCost
	GpuCost() unsafe.Pointer

	// GpuTime protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/gpuTime
	GpuTime() uint64

	// Kicks protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/kicks
	Kicks() unsafe.Pointer

	// KicksCount protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/kicksCount
	KicksCount() uint64

	// MetalFXInfo protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/metalFXInfo
	MetalFXInfo() unsafe.Pointer

	// MetalFXInfoCount protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/metalFXInfoCount
	MetalFXInfoCount() uint64

	// NumDrawsForEncoder protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/numDrawsForEncoder:
	NumDrawsForEncoder(encoder uint32) uint64

	// NumDrawsForPipelineState protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/numDrawsForPipelineState:
	NumDrawsForPipelineState(state uint64) uint64

	// PipelineStateCount protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/pipelineStateCount
	PipelineStateCount() uint64

	// PipelineStateIdForCliqueIndex protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/pipelineStateIdForCliqueIndex:
	PipelineStateIdForCliqueIndex(index unsafe.Pointer) uint64

	// ProfiledState protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/profiledState
	ProfiledState() uint32

	// ProfiledWithOverlapEnabled protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/profiledWithOverlapEnabled
	ProfiledWithOverlapEnabled() bool

	// ReferenceComputePositionForClique protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/referenceComputePositionForClique:
	ReferenceComputePositionForClique(clique *GTMioUSCCliqueMetadataRef) unsafe.Pointer

	// RiaTraceCount protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/riaTraceCount
	RiaTraceCount() uint64

	// RiaTraces protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/riaTraces
	RiaTraces() unsafe.Pointer

	// SampledCores protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/sampledCores
	SampledCores() uint32

	// ShaderBinaryInfo protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/shaderBinaryInfo
	ShaderBinaryInfo() unsafe.Pointer

	// ShaderBinaryInfoCount protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/shaderBinaryInfoCount
	ShaderBinaryInfoCount() uint64

	// SignpostPipelineStateCount protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/signpostPipelineStateCount
	SignpostPipelineStateCount() uint64

	// SignpostPipelineStates protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/signpostPipelineStates
	SignpostPipelineStates() unsafe.Pointer

	// SignpostProcessCount protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/signpostProcessCount
	SignpostProcessCount() uint64

	// SignpostProcesses protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/signpostProcesses
	SignpostProcesses() unsafe.Pointer

	// SignpostShaderCount protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/signpostShaderCount
	SignpostShaderCount() uint64

	// SignpostShaders protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/signpostShaders
	SignpostShaders() unsafe.Pointer

	// TimelineDuration protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/timelineDuration
	TimelineDuration() uint64

	// TimestampBegin protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/timestampBegin
	TimestampBegin() uint64

	// TimestampEnd protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/timestampEnd
	TimestampEnd() uint64

	// TotalCliqueCost protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/totalCliqueCost
	TotalCliqueCost() uint64

	// TotalCores protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/totalCores
	TotalCores() uint32

	// TotalCostForScopeScopeIdentifierProgramType protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/totalCostForScope:scopeIdentifier:programType:
	TotalCostForScopeScopeIdentifierProgramType(scope uint16, identifier uint64, type_ uint16) float64
}

// GTMioTraceDataProtocolObject wraps an existing Objective-C object that conforms to the GTMioTraceDataProtocol protocol.
type GTMioTraceDataProtocolObject struct {
	objectivec.Object
}

func (o GTMioTraceDataProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// GTMioTraceDataProtocolObjectFromID constructs a [GTMioTraceDataProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GTMioTraceDataProtocolObjectFromID(id objc.ID) GTMioTraceDataProtocolObject {
	return GTMioTraceDataProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/binaries
func (o GTMioTraceDataProtocolObject) Binaries() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("binaries"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/binaryForDraw:programType:
func (o GTMioTraceDataProtocolObject) BinaryForDrawProgramType(draw uint32, type_ uint16) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("binaryForDraw:programType:"), draw, type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/binaryForPipelineState:programType:
func (o GTMioTraceDataProtocolObject) BinaryForPipelineStateProgramType(state uint64, type_ uint16) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("binaryForPipelineState:programType:"), state, type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/childCliqueOfClique:
func (o GTMioTraceDataProtocolObject) ChildCliqueOfClique(clique *GTMioUSCCliqueMetadataRef) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("childCliqueOfClique:"), clique)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/cliqueFromCliqueIndex:
func (o GTMioTraceDataProtocolObject) CliqueFromCliqueIndex(index unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("cliqueFromCliqueIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/coalescedFunctionIndexForEncoderFunctionIndex:
func (o GTMioTraceDataProtocolObject) CoalescedFunctionIndexForEncoderFunctionIndex(index uint32) uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("coalescedFunctionIndexForEncoderFunctionIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/computePositionCount
func (o GTMioTraceDataProtocolObject) ComputePositionCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("computePositionCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/computePositions
func (o GTMioTraceDataProtocolObject) ComputePositions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("computePositions"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/consistentStateAchieved
func (o GTMioTraceDataProtocolObject) ConsistentStateAchieved() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("consistentStateAchieved"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/costForContext:cost:
func (o GTMioTraceDataProtocolObject) CostForContextCost(context unsafe.Pointer, cost unsafe.Pointer) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("costForContext:cost:"), context, cost)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/costForLevel:levelIdentifier:scope:scopeIdentifier:cost:
func (o GTMioTraceDataProtocolObject) CostForLevelLevelIdentifierScopeScopeIdentifierCost(level uint16, identifier uint32, scope uint16, identifier2 uint64, cost unsafe.Pointer) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("costForLevel:levelIdentifier:scope:scopeIdentifier:cost:"), level, identifier, scope, identifier2, cost)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/dataType
func (o GTMioTraceDataProtocolObject) DataType() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("dataType"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/databaseInternal
func (o GTMioTraceDataProtocolObject) DatabaseInternal() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("databaseInternal"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/drawCount
func (o GTMioTraceDataProtocolObject) DrawCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("drawCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/drawTraceCount
func (o GTMioTraceDataProtocolObject) DrawTraceCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("drawTraceCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/drawTraces
func (o GTMioTraceDataProtocolObject) DrawTraces() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("drawTraces"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/draws
func (o GTMioTraceDataProtocolObject) Draws() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("draws"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/durationForDraw:dataMaster:
func (o GTMioTraceDataProtocolObject) DurationForDrawDataMaster(draw uint32, master uint16) uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("durationForDraw:dataMaster:"), draw, master)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/encoderCount
func (o GTMioTraceDataProtocolObject) EncoderCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("encoderCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/encoderFromFunctionIndex:
func (o GTMioTraceDataProtocolObject) EncoderFromFunctionIndex(index uint32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("encoderFromFunctionIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/encoders
func (o GTMioTraceDataProtocolObject) Encoders() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("encoders"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateBinariesForDraw:enumerator:
func (o GTMioTraceDataProtocolObject) EnumerateBinariesForDrawEnumerator(draw uint32, enumerator GTMioShaderBinaryDataHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("enumerateBinariesForDraw:enumerator:"), draw, enumerator)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateBinariesForEncoder:enumerator:
func (o GTMioTraceDataProtocolObject) EnumerateBinariesForEncoderEnumerator(encoder uint32, enumerator GTMioShaderBinaryDataHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("enumerateBinariesForEncoder:enumerator:"), encoder, enumerator)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateBinariesForForCliqueAtIndex:uscIndex:enumerator:
func (o GTMioTraceDataProtocolObject) EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator GTMioShaderBinaryDataHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("enumerateBinariesForForCliqueAtIndex:uscIndex:enumerator:"), index, index2, enumerator)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateBinariesForPipelineState:enumerator:
func (o GTMioTraceDataProtocolObject) EnumerateBinariesForPipelineStateEnumerator(state uint64, enumerator GTMioShaderBinaryDataHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("enumerateBinariesForPipelineState:enumerator:"), state, enumerator)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateBinaryRangesForClique:uscData:enumerator:
func (o GTMioTraceDataProtocolObject) EnumerateBinaryRangesForCliqueUscDataEnumerator(clique *GTMioUSCCliqueMetadataRef, data objectivec.IObject, enumerator GTMioShaderBinaryDebugBinaryRangeHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("enumerateBinaryRangesForClique:uscData:enumerator:"), clique, data, enumerator)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateBinaryRangesForCliqueAtIndex:uscIndex:enumerator:
func (o GTMioTraceDataProtocolObject) EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator GTMioShaderBinaryDebugBinaryRangeHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("enumerateBinaryRangesForCliqueAtIndex:uscIndex:enumerator:"), index, index2, enumerator)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateDrawsForEncoder:enumerator:
func (o GTMioTraceDataProtocolObject) EnumerateDrawsForEncoderEnumerator(encoder uint32, enumerator UintHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("enumerateDrawsForEncoder:enumerator:"), encoder, enumerator)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateDrawsForPipelineState:enumerator:
func (o GTMioTraceDataProtocolObject) EnumerateDrawsForPipelineStateEnumerator(state uint64, enumerator UintHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("enumerateDrawsForPipelineState:enumerator:"), state, enumerator)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateEncoders:
func (o GTMioTraceDataProtocolObject) EnumerateEncoders(encoders UintHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("enumerateEncoders:"), encoders)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateInstructionsForClique:uscData:enumerator:
func (o GTMioTraceDataProtocolObject) EnumerateInstructionsForCliqueUscDataEnumerator(clique *GTMioUSCCliqueMetadataRef, data objectivec.IObject, enumerator GTMioShaderInstructionInfoHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("enumerateInstructionsForClique:uscData:enumerator:"), clique, data, enumerator)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateInstructionsForCliqueAtIndex:uscIndex:enumerator:
func (o GTMioTraceDataProtocolObject) EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator GTMioShaderInstructionInfoHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("enumerateInstructionsForCliqueAtIndex:uscIndex:enumerator:"), index, index2, enumerator)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateKickAtFunctionIndex:enumerator:
func (o GTMioTraceDataProtocolObject) EnumerateKickAtFunctionIndexEnumerator(index uint32, enumerator GTMioUSCKickMetadataHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("enumerateKickAtFunctionIndex:enumerator:"), index, enumerator)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumeratePipelineStates:
func (o GTMioTraceDataProtocolObject) EnumeratePipelineStates(states unsignedlongHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("enumeratePipelineStates:"), states)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/enumerateUniqueTracesForBinary:enumerator:
func (o GTMioTraceDataProtocolObject) EnumerateUniqueTracesForBinaryEnumerator(binary uint32, enumerator GTMioUSCTraceDataHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("enumerateUniqueTracesForBinary:enumerator:"), binary, enumerator)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/executionHistoryForClique:uscIndex:delegate:
func (o GTMioTraceDataProtocolObject) ExecutionHistoryForCliqueUscIndexDelegate(clique uint32, index uint32, delegate objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("executionHistoryForClique:uscIndex:delegate:"), clique, index, delegate)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/executionHistoryForDraw:programType:delegate:progressController:
func (o GTMioTraceDataProtocolObject) ExecutionHistoryForDrawProgramTypeDelegateProgressController(draw uint32, type_ uint16, delegate objectivec.IObject, controller objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("executionHistoryForDraw:programType:delegate:progressController:"), draw, type_, delegate, controller)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/executionHistoryForPipelineState:programType:delegate:progressController:
func (o GTMioTraceDataProtocolObject) ExecutionHistoryForPipelineStateProgramTypeDelegateProgressController(state uint64, type_ uint16, delegate objectivec.IObject, controller objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("executionHistoryForPipelineState:programType:delegate:progressController:"), state, type_, delegate, controller)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/failedUSCIndexCount
func (o GTMioTraceDataProtocolObject) FailedUSCIndexCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("failedUSCIndexCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/failedUSCIndexes
func (o GTMioTraceDataProtocolObject) FailedUSCIndexes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("failedUSCIndexes"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/fragmentPositionCount
func (o GTMioTraceDataProtocolObject) FragmentPositionCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("fragmentPositionCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/fragmentPositions
func (o GTMioTraceDataProtocolObject) FragmentPositions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("fragmentPositions"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/globalGPUTime
func (o GTMioTraceDataProtocolObject) GlobalGPUTime() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("globalGPUTime"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/gpuCost
func (o GTMioTraceDataProtocolObject) GpuCost() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("gpuCost"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/gpuInfo
func (o GTMioTraceDataProtocolObject) GpuInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("gpuInfo"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/gpuTime
func (o GTMioTraceDataProtocolObject) GpuTime() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("gpuTime"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/kicks
func (o GTMioTraceDataProtocolObject) Kicks() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("kicks"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/kicksCount
func (o GTMioTraceDataProtocolObject) KicksCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("kicksCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/mGPUs
func (o GTMioTraceDataProtocolObject) MGPUs() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("mGPUs"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/metalFXInfo
func (o GTMioTraceDataProtocolObject) MetalFXInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("metalFXInfo"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/metalFXInfoCount
func (o GTMioTraceDataProtocolObject) MetalFXInfoCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("metalFXInfoCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/numDrawsForEncoder:
func (o GTMioTraceDataProtocolObject) NumDrawsForEncoder(encoder uint32) uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("numDrawsForEncoder:"), encoder)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/numDrawsForPipelineState:
func (o GTMioTraceDataProtocolObject) NumDrawsForPipelineState(state uint64) uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("numDrawsForPipelineState:"), state)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/parentData
func (o GTMioTraceDataProtocolObject) ParentData() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("parentData"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/pipelineStateCount
func (o GTMioTraceDataProtocolObject) PipelineStateCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("pipelineStateCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/pipelineStateIdForCliqueIndex:
func (o GTMioTraceDataProtocolObject) PipelineStateIdForCliqueIndex(index unsafe.Pointer) uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("pipelineStateIdForCliqueIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/profiledState
func (o GTMioTraceDataProtocolObject) ProfiledState() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("profiledState"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/profiledWithOverlapEnabled
func (o GTMioTraceDataProtocolObject) ProfiledWithOverlapEnabled() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("profiledWithOverlapEnabled"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/referenceComputePositionForClique:
func (o GTMioTraceDataProtocolObject) ReferenceComputePositionForClique(clique *GTMioUSCCliqueMetadataRef) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("referenceComputePositionForClique:"), clique)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/riaTraceCount
func (o GTMioTraceDataProtocolObject) RiaTraceCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("riaTraceCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/riaTraces
func (o GTMioTraceDataProtocolObject) RiaTraces() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("riaTraces"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/sampledCores
func (o GTMioTraceDataProtocolObject) SampledCores() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("sampledCores"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/setParentData:
func (o GTMioTraceDataProtocolObject) SetParentData(data objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setParentData:"), data)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/shaderBinaryInfo
func (o GTMioTraceDataProtocolObject) ShaderBinaryInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("shaderBinaryInfo"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/shaderBinaryInfoCount
func (o GTMioTraceDataProtocolObject) ShaderBinaryInfoCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("shaderBinaryInfoCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/signpostPipelineStateCount
func (o GTMioTraceDataProtocolObject) SignpostPipelineStateCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("signpostPipelineStateCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/signpostPipelineStates
func (o GTMioTraceDataProtocolObject) SignpostPipelineStates() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("signpostPipelineStates"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/signpostProcessCount
func (o GTMioTraceDataProtocolObject) SignpostProcessCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("signpostProcessCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/signpostProcesses
func (o GTMioTraceDataProtocolObject) SignpostProcesses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("signpostProcesses"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/signpostShaderCount
func (o GTMioTraceDataProtocolObject) SignpostShaderCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("signpostShaderCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/signpostShaders
func (o GTMioTraceDataProtocolObject) SignpostShaders() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("signpostShaders"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/signpostStrings
func (o GTMioTraceDataProtocolObject) SignpostStrings() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("signpostStrings"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/streamData
func (o GTMioTraceDataProtocolObject) StreamData() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("streamData"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/timelineCounters
func (o GTMioTraceDataProtocolObject) TimelineCounters() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("timelineCounters"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/timelineDuration
func (o GTMioTraceDataProtocolObject) TimelineDuration() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("timelineDuration"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/timestampBegin
func (o GTMioTraceDataProtocolObject) TimestampBegin() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("timestampBegin"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/timestampEnd
func (o GTMioTraceDataProtocolObject) TimestampEnd() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("timestampEnd"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/totalCliqueCost
func (o GTMioTraceDataProtocolObject) TotalCliqueCost() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("totalCliqueCost"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/totalCores
func (o GTMioTraceDataProtocolObject) TotalCores() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("totalCores"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/totalCostForScope:scopeIdentifier:programType:
func (o GTMioTraceDataProtocolObject) TotalCostForScopeScopeIdentifierProgramType(scope uint16, identifier uint64, type_ uint16) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("totalCostForScope:scopeIdentifier:programType:"), scope, identifier, type_)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataProtocol/uscs
func (o GTMioTraceDataProtocolObject) Uscs() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("uscs"))
	return objectivec.Object{ID: rv}
}

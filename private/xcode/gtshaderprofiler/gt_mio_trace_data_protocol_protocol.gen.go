// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTMioTraceDataProtocol protocol.
type GTMioTraceDataProtocol interface {
	objectivec.IObject

	// Binaries protocol.
	Binaries() objectivec.IObject

	// BinaryForDrawProgramType protocol.
	BinaryForDrawProgramType(draw uint32, type_ uint16) objectivec.IObject

	// BinaryForPipelineStateProgramType protocol.
	BinaryForPipelineStateProgramType(state uint64, type_ uint16) objectivec.IObject

	// ChildCliqueOfClique protocol.
	ChildCliqueOfClique(clique *GTMioUSCCliqueMetadata) *GTMioUSCCliqueMetadata

	// CliqueFromCliqueIndex protocol.
	CliqueFromCliqueIndex(index *GTMioUSCCliqueIndex) *GTMioUSCCliqueMetadata

	// CoalescedFunctionIndexForEncoderFunctionIndex protocol.
	CoalescedFunctionIndexForEncoderFunctionIndex(index uint32) uint32

	// ComputePositionCount protocol.
	ComputePositionCount() uint64

	// ComputePositions protocol.
	ComputePositions() *GTMioCliqueComputePosition

	// ConsistentStateAchieved protocol.
	ConsistentStateAchieved() bool

	// CostForContextCost protocol.
	CostForContextCost(context *GTMioCostContext, cost *GTMioCostInfo) bool

	// CostForLevelLevelIdentifierScopeScopeIdentifierCost protocol.
	CostForLevelLevelIdentifierScopeScopeIdentifierCost(level uint16, identifier uint32, scope uint16, identifier2 uint64, cost *GTMioCostInfo) bool

	// DataType protocol.
	DataType() uint32

	// DatabaseInternal protocol.
	DatabaseInternal() uint64

	// DrawCount protocol.
	DrawCount() uint64

	// DrawTraceCount protocol.
	DrawTraceCount() uint64

	// DrawTraces protocol.
	DrawTraces() *GTMioDrawTrace

	// Draws protocol.
	Draws() *GTMioDrawMetadata

	// DurationForDrawDataMaster protocol.
	DurationForDrawDataMaster(draw uint32, master uint16) uint64

	// EncoderCount protocol.
	EncoderCount() uint64

	// EncoderFromFunctionIndex protocol.
	EncoderFromFunctionIndex(index uint32) *GTMioEncoderMetadata

	// Encoders protocol.
	Encoders() *GTMioEncoderMetadata

	// EnumerateBinariesForDrawEnumerator protocol.
	EnumerateBinariesForDrawEnumerator(draw uint32, enumerator GTMioShaderBinaryInfoGTMioShaderBinaryDataHandler)

	// EnumerateBinariesForEncoderEnumerator protocol.
	EnumerateBinariesForEncoderEnumerator(encoder uint32, enumerator GTMioShaderBinaryInfoGTMioShaderBinaryDataHandler)

	// EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator protocol.
	EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator GTMioShaderBinaryDataHandler)

	// EnumerateBinariesForPipelineStateEnumerator protocol.
	EnumerateBinariesForPipelineStateEnumerator(state uint64, enumerator GTMioShaderBinaryDataHandler)

	// EnumerateBinaryRangesForCliqueUscDataEnumerator protocol.
	EnumerateBinaryRangesForCliqueUscDataEnumerator(clique *GTMioUSCCliqueMetadata, data objectivec.IObject, enumerator GTMioShaderBinaryDataGTMioShaderBinaryDebugBinaryRangeUint32Handler)

	// EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator protocol.
	EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator GTMioShaderBinaryDataGTMioShaderBinaryDebugBinaryRangeUint32Handler)

	// EnumerateDrawsForEncoderEnumerator protocol.
	EnumerateDrawsForEncoderEnumerator(encoder uint32, enumerator Uint32Handler)

	// EnumerateDrawsForPipelineStateEnumerator protocol.
	EnumerateDrawsForPipelineStateEnumerator(state uint64, enumerator Uint32Handler)

	// EnumerateEncoders protocol.
	EnumerateEncoders(encoders Uint32Handler)

	// EnumerateInstructionsForCliqueUscDataEnumerator protocol.
	EnumerateInstructionsForCliqueUscDataEnumerator(clique *GTMioUSCCliqueMetadata, data objectivec.IObject, enumerator GTMioShaderBinaryDataGTMioShaderInstructionInfoUint32Handler)

	// EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator protocol.
	EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator GTMioShaderBinaryDataGTMioShaderInstructionInfoUint32Handler)

	// EnumerateKickAtFunctionIndexEnumerator protocol.
	EnumerateKickAtFunctionIndexEnumerator(index uint32, enumerator GTMioUSCTraceDataGTMioUSCKickMetadataUint64Handler)

	// EnumeratePipelineStates protocol.
	EnumeratePipelineStates(states unsignedlongHandler)

	// EnumerateUniqueTracesForBinaryEnumerator protocol.
	EnumerateUniqueTracesForBinaryEnumerator(binary uint32, enumerator GTMioUSCCliqueMetadataGTMioUSCTraceDataHandler)

	// ExecutionHistoryForCliqueUscIndexDelegate protocol.
	ExecutionHistoryForCliqueUscIndexDelegate(clique uint32, index uint32, delegate objectivec.IObject)

	// ExecutionHistoryForDrawProgramTypeDelegateProgressController protocol.
	ExecutionHistoryForDrawProgramTypeDelegateProgressController(draw uint32, type_ uint16, delegate objectivec.IObject, controller objectivec.IObject)

	// ExecutionHistoryForPipelineStateProgramTypeDelegateProgressController protocol.
	ExecutionHistoryForPipelineStateProgramTypeDelegateProgressController(state uint64, type_ uint16, delegate objectivec.IObject, controller objectivec.IObject)

	// FailedUSCIndexCount protocol.
	FailedUSCIndexCount() uint64

	// FailedUSCIndexes protocol.
	FailedUSCIndexes() unsafe.Pointer

	// FragmentPositionCount protocol.
	FragmentPositionCount() uint64

	// FragmentPositions protocol.
	FragmentPositions() *GTMioCliqueFragmentPosition

	// GlobalGPUTime protocol.
	GlobalGPUTime() uint64

	// GpuCost protocol.
	GpuCost() *GTMioCostInfo

	// GpuInfo protocol.
	GpuInfo() objectivec.IObject

	// GpuTime protocol.
	GpuTime() uint64

	// Kicks protocol.
	Kicks() *GTMioKickTrace

	// KicksCount protocol.
	KicksCount() uint64

	// MGPUs protocol.
	MGPUs() objectivec.IObject

	// MetalFXInfo protocol.
	MetalFXInfo() *GTMioMetalFXInfo

	// MetalFXInfoCount protocol.
	MetalFXInfoCount() uint64

	// NumDrawsForEncoder protocol.
	NumDrawsForEncoder(encoder uint32) uint64

	// NumDrawsForPipelineState protocol.
	NumDrawsForPipelineState(state uint64) uint64

	// ParentData protocol.
	ParentData() objectivec.IObject

	// PipelineStateCount protocol.
	PipelineStateCount() uint64

	// PipelineStateIdForCliqueIndex protocol.
	PipelineStateIdForCliqueIndex(index *GTMioUSCCliqueIndex) uint64

	// ProfiledState protocol.
	ProfiledState() uint32

	// ProfiledWithOverlapEnabled protocol.
	ProfiledWithOverlapEnabled() bool

	// ReferenceComputePositionForClique protocol.
	ReferenceComputePositionForClique(clique *GTMioUSCCliqueMetadata) *GTMioCliqueComputePosition

	// RiaTraceCount protocol.
	RiaTraceCount() uint64

	// RiaTraces protocol.
	RiaTraces() *GTMioRIATrace

	// SampledCores protocol.
	SampledCores() uint32

	// SetParentData protocol.
	SetParentData(data objectivec.IObject)

	// ShaderBinaryInfo protocol.
	ShaderBinaryInfo() *GTMioShaderBinaryInfo

	// ShaderBinaryInfoCount protocol.
	ShaderBinaryInfoCount() uint64

	// SignpostPipelineStateCount protocol.
	SignpostPipelineStateCount() uint64

	// SignpostPipelineStates protocol.
	SignpostPipelineStates() *GTMioShaderTimelineSignpostPipelineState

	// SignpostProcessCount protocol.
	SignpostProcessCount() uint64

	// SignpostProcesses protocol.
	SignpostProcesses() *GTMioShaderTimelineSignpostProcess

	// SignpostShaderCount protocol.
	SignpostShaderCount() uint64

	// SignpostShaders protocol.
	SignpostShaders() *GTMioShaderTimelineSignpostShader

	// SignpostStrings protocol.
	SignpostStrings() objectivec.IObject

	// StreamData protocol.
	StreamData() objectivec.IObject

	// TimelineCounters protocol.
	TimelineCounters() objectivec.IObject

	// TimelineDuration protocol.
	TimelineDuration() uint64

	// TimestampBegin protocol.
	TimestampBegin() uint64

	// TimestampEnd protocol.
	TimestampEnd() uint64

	// TotalCliqueCost protocol.
	TotalCliqueCost() uint64

	// TotalCores protocol.
	TotalCores() uint32

	// TotalCostForScopeScopeIdentifierProgramType protocol.
	TotalCostForScopeScopeIdentifierProgramType(scope uint16, identifier uint64, type_ uint16) float64

	// Uscs protocol.
	Uscs() objectivec.IObject
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

func (o GTMioTraceDataProtocolObject) Binaries() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("binaries"))
	return objectivec.Object{ID: rv}
}
func (o GTMioTraceDataProtocolObject) BinaryForDrawProgramType(draw uint32, type_ uint16) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("binaryForDraw:programType:"), draw, type_)
	return objectivec.Object{ID: rv}
}
func (o GTMioTraceDataProtocolObject) BinaryForPipelineStateProgramType(state uint64, type_ uint16) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("binaryForPipelineState:programType:"), state, type_)
	return objectivec.Object{ID: rv}
}
func (o GTMioTraceDataProtocolObject) ChildCliqueOfClique(clique *GTMioUSCCliqueMetadata) *GTMioUSCCliqueMetadata {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("childCliqueOfClique:"), unsafe.Pointer(clique))
	return (*GTMioUSCCliqueMetadata)(rv)
}
func (o GTMioTraceDataProtocolObject) CliqueFromCliqueIndex(index *GTMioUSCCliqueIndex) *GTMioUSCCliqueMetadata {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("cliqueFromCliqueIndex:"), unsafe.Pointer(index))
	return (*GTMioUSCCliqueMetadata)(rv)
}
func (o GTMioTraceDataProtocolObject) CoalescedFunctionIndexForEncoderFunctionIndex(index uint32) uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("coalescedFunctionIndexForEncoderFunctionIndex:"), index)
	return rv
}
func (o GTMioTraceDataProtocolObject) ComputePositionCount() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("computePositionCount"))
	return rv
}
func (o GTMioTraceDataProtocolObject) ComputePositions() *GTMioCliqueComputePosition {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("computePositions"))
	return (*GTMioCliqueComputePosition)(rv)
}
func (o GTMioTraceDataProtocolObject) ConsistentStateAchieved() bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("consistentStateAchieved"))
	return rv
}
func (o GTMioTraceDataProtocolObject) CostForContextCost(context *GTMioCostContext, cost *GTMioCostInfo) bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("costForContext:cost:"), unsafe.Pointer(context), unsafe.Pointer(cost))
	return rv
}
func (o GTMioTraceDataProtocolObject) CostForLevelLevelIdentifierScopeScopeIdentifierCost(level uint16, identifier uint32, scope uint16, identifier2 uint64, cost *GTMioCostInfo) bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("costForLevel:levelIdentifier:scope:scopeIdentifier:cost:"), level, identifier, scope, identifier2, unsafe.Pointer(cost))
	return rv
}
func (o GTMioTraceDataProtocolObject) DataType() uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("dataType"))
	return rv
}
func (o GTMioTraceDataProtocolObject) DatabaseInternal() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("databaseInternal"))
	return rv
}
func (o GTMioTraceDataProtocolObject) DrawCount() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("drawCount"))
	return rv
}
func (o GTMioTraceDataProtocolObject) DrawTraceCount() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("drawTraceCount"))
	return rv
}
func (o GTMioTraceDataProtocolObject) DrawTraces() *GTMioDrawTrace {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("drawTraces"))
	return (*GTMioDrawTrace)(rv)
}
func (o GTMioTraceDataProtocolObject) Draws() *GTMioDrawMetadata {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("draws"))
	return (*GTMioDrawMetadata)(rv)
}
func (o GTMioTraceDataProtocolObject) DurationForDrawDataMaster(draw uint32, master uint16) uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("durationForDraw:dataMaster:"), draw, master)
	return rv
}
func (o GTMioTraceDataProtocolObject) EncoderCount() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("encoderCount"))
	return rv
}
func (o GTMioTraceDataProtocolObject) EncoderFromFunctionIndex(index uint32) *GTMioEncoderMetadata {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("encoderFromFunctionIndex:"), index)
	return (*GTMioEncoderMetadata)(rv)
}
func (o GTMioTraceDataProtocolObject) Encoders() *GTMioEncoderMetadata {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("encoders"))
	return (*GTMioEncoderMetadata)(rv)
}
func (o GTMioTraceDataProtocolObject) EnumerateBinariesForDrawEnumerator(draw uint32, enumerator GTMioShaderBinaryInfoGTMioShaderBinaryDataHandler) {
	_block1, _cleanup1 := NewGTMioShaderBinaryInfoGTMioShaderBinaryDataBlock(enumerator)
	defer _cleanup1()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("enumerateBinariesForDraw:enumerator:"), draw, objc.ID(_block1))
}
func (o GTMioTraceDataProtocolObject) EnumerateBinariesForEncoderEnumerator(encoder uint32, enumerator GTMioShaderBinaryInfoGTMioShaderBinaryDataHandler) {
	_block1, _cleanup1 := NewGTMioShaderBinaryInfoGTMioShaderBinaryDataBlock(enumerator)
	defer _cleanup1()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("enumerateBinariesForEncoder:enumerator:"), encoder, objc.ID(_block1))
}
func (o GTMioTraceDataProtocolObject) EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator GTMioShaderBinaryDataHandler) {
	_block2, _cleanup2 := NewGTMioShaderBinaryDataBlock(enumerator)
	defer _cleanup2()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("enumerateBinariesForForCliqueAtIndex:uscIndex:enumerator:"), index, index2, objc.ID(_block2))
}
func (o GTMioTraceDataProtocolObject) EnumerateBinariesForPipelineStateEnumerator(state uint64, enumerator GTMioShaderBinaryDataHandler) {
	_block1, _cleanup1 := NewGTMioShaderBinaryDataBlock(enumerator)
	defer _cleanup1()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("enumerateBinariesForPipelineState:enumerator:"), state, objc.ID(_block1))
}
func (o GTMioTraceDataProtocolObject) EnumerateBinaryRangesForCliqueUscDataEnumerator(clique *GTMioUSCCliqueMetadata, data objectivec.IObject, enumerator GTMioShaderBinaryDataGTMioShaderBinaryDebugBinaryRangeUint32Handler) {
	_block2, _cleanup2 := NewGTMioShaderBinaryDataGTMioShaderBinaryDebugBinaryRangeUint32Block(enumerator)
	defer _cleanup2()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("enumerateBinaryRangesForClique:uscData:enumerator:"), clique, data, objc.ID(_block2))
}
func (o GTMioTraceDataProtocolObject) EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator GTMioShaderBinaryDataGTMioShaderBinaryDebugBinaryRangeUint32Handler) {
	_block2, _cleanup2 := NewGTMioShaderBinaryDataGTMioShaderBinaryDebugBinaryRangeUint32Block(enumerator)
	defer _cleanup2()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("enumerateBinaryRangesForCliqueAtIndex:uscIndex:enumerator:"), index, index2, objc.ID(_block2))
}
func (o GTMioTraceDataProtocolObject) EnumerateDrawsForEncoderEnumerator(encoder uint32, enumerator Uint32Handler) {
	_block1, _cleanup1 := NewUint32Block(enumerator)
	defer _cleanup1()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("enumerateDrawsForEncoder:enumerator:"), encoder, objc.ID(_block1))
}
func (o GTMioTraceDataProtocolObject) EnumerateDrawsForPipelineStateEnumerator(state uint64, enumerator Uint32Handler) {
	_block1, _cleanup1 := NewUint32Block(enumerator)
	defer _cleanup1()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("enumerateDrawsForPipelineState:enumerator:"), state, objc.ID(_block1))
}
func (o GTMioTraceDataProtocolObject) EnumerateEncoders(encoders Uint32Handler) {
	_block0, _cleanup0 := NewUint32Block(encoders)
	defer _cleanup0()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("enumerateEncoders:"), objc.ID(_block0))
}
func (o GTMioTraceDataProtocolObject) EnumerateInstructionsForCliqueUscDataEnumerator(clique *GTMioUSCCliqueMetadata, data objectivec.IObject, enumerator GTMioShaderBinaryDataGTMioShaderInstructionInfoUint32Handler) {
	_block2, _cleanup2 := NewGTMioShaderBinaryDataGTMioShaderInstructionInfoUint32Block(enumerator)
	defer _cleanup2()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("enumerateInstructionsForClique:uscData:enumerator:"), clique, data, objc.ID(_block2))
}
func (o GTMioTraceDataProtocolObject) EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator(index uint32, index2 uint32, enumerator GTMioShaderBinaryDataGTMioShaderInstructionInfoUint32Handler) {
	_block2, _cleanup2 := NewGTMioShaderBinaryDataGTMioShaderInstructionInfoUint32Block(enumerator)
	defer _cleanup2()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("enumerateInstructionsForCliqueAtIndex:uscIndex:enumerator:"), index, index2, objc.ID(_block2))
}
func (o GTMioTraceDataProtocolObject) EnumerateKickAtFunctionIndexEnumerator(index uint32, enumerator GTMioUSCTraceDataGTMioUSCKickMetadataUint64Handler) {
	_block1, _cleanup1 := NewGTMioUSCTraceDataGTMioUSCKickMetadataUint64Block(enumerator)
	defer _cleanup1()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("enumerateKickAtFunctionIndex:enumerator:"), index, objc.ID(_block1))
}
func (o GTMioTraceDataProtocolObject) EnumeratePipelineStates(states unsignedlongHandler) {
	_block0, _cleanup0 := NewunsignedlongBlock(states)
	defer _cleanup0()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("enumeratePipelineStates:"), objc.ID(_block0))
}
func (o GTMioTraceDataProtocolObject) EnumerateUniqueTracesForBinaryEnumerator(binary uint32, enumerator GTMioUSCCliqueMetadataGTMioUSCTraceDataHandler) {
	_block1, _cleanup1 := NewGTMioUSCCliqueMetadataGTMioUSCTraceDataBlock(enumerator)
	defer _cleanup1()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("enumerateUniqueTracesForBinary:enumerator:"), binary, objc.ID(_block1))
}
func (o GTMioTraceDataProtocolObject) ExecutionHistoryForCliqueUscIndexDelegate(clique uint32, index uint32, delegate objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("executionHistoryForClique:uscIndex:delegate:"), clique, index, delegate)
}
func (o GTMioTraceDataProtocolObject) ExecutionHistoryForDrawProgramTypeDelegateProgressController(draw uint32, type_ uint16, delegate objectivec.IObject, controller objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("executionHistoryForDraw:programType:delegate:progressController:"), draw, type_, delegate, controller)
}
func (o GTMioTraceDataProtocolObject) ExecutionHistoryForPipelineStateProgramTypeDelegateProgressController(state uint64, type_ uint16, delegate objectivec.IObject, controller objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("executionHistoryForPipelineState:programType:delegate:progressController:"), state, type_, delegate, controller)
}
func (o GTMioTraceDataProtocolObject) FailedUSCIndexCount() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("failedUSCIndexCount"))
	return rv
}
func (o GTMioTraceDataProtocolObject) FailedUSCIndexes() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("failedUSCIndexes"))
	return rv
}
func (o GTMioTraceDataProtocolObject) FragmentPositionCount() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("fragmentPositionCount"))
	return rv
}
func (o GTMioTraceDataProtocolObject) FragmentPositions() *GTMioCliqueFragmentPosition {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("fragmentPositions"))
	return (*GTMioCliqueFragmentPosition)(rv)
}
func (o GTMioTraceDataProtocolObject) GlobalGPUTime() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("globalGPUTime"))
	return rv
}
func (o GTMioTraceDataProtocolObject) GpuCost() *GTMioCostInfo {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("gpuCost"))
	return (*GTMioCostInfo)(rv)
}
func (o GTMioTraceDataProtocolObject) GpuInfo() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("gpuInfo"))
	return objectivec.Object{ID: rv}
}
func (o GTMioTraceDataProtocolObject) GpuTime() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("gpuTime"))
	return rv
}
func (o GTMioTraceDataProtocolObject) Kicks() *GTMioKickTrace {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("kicks"))
	return (*GTMioKickTrace)(rv)
}
func (o GTMioTraceDataProtocolObject) KicksCount() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("kicksCount"))
	return rv
}
func (o GTMioTraceDataProtocolObject) MGPUs() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("mGPUs"))
	return objectivec.Object{ID: rv}
}
func (o GTMioTraceDataProtocolObject) MetalFXInfo() *GTMioMetalFXInfo {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("metalFXInfo"))
	return (*GTMioMetalFXInfo)(rv)
}
func (o GTMioTraceDataProtocolObject) MetalFXInfoCount() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("metalFXInfoCount"))
	return rv
}
func (o GTMioTraceDataProtocolObject) NumDrawsForEncoder(encoder uint32) uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("numDrawsForEncoder:"), encoder)
	return rv
}
func (o GTMioTraceDataProtocolObject) NumDrawsForPipelineState(state uint64) uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("numDrawsForPipelineState:"), state)
	return rv
}
func (o GTMioTraceDataProtocolObject) ParentData() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("parentData"))
	return objectivec.Object{ID: rv}
}
func (o GTMioTraceDataProtocolObject) PipelineStateCount() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("pipelineStateCount"))
	return rv
}
func (o GTMioTraceDataProtocolObject) PipelineStateIdForCliqueIndex(index *GTMioUSCCliqueIndex) uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("pipelineStateIdForCliqueIndex:"), unsafe.Pointer(index))
	return rv
}
func (o GTMioTraceDataProtocolObject) ProfiledState() uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("profiledState"))
	return rv
}
func (o GTMioTraceDataProtocolObject) ProfiledWithOverlapEnabled() bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("profiledWithOverlapEnabled"))
	return rv
}
func (o GTMioTraceDataProtocolObject) ReferenceComputePositionForClique(clique *GTMioUSCCliqueMetadata) *GTMioCliqueComputePosition {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("referenceComputePositionForClique:"), unsafe.Pointer(clique))
	return (*GTMioCliqueComputePosition)(rv)
}
func (o GTMioTraceDataProtocolObject) RiaTraceCount() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("riaTraceCount"))
	return rv
}
func (o GTMioTraceDataProtocolObject) RiaTraces() *GTMioRIATrace {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("riaTraces"))
	return (*GTMioRIATrace)(rv)
}
func (o GTMioTraceDataProtocolObject) SampledCores() uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("sampledCores"))
	return rv
}
func (o GTMioTraceDataProtocolObject) SetParentData(data objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setParentData:"), data)
}
func (o GTMioTraceDataProtocolObject) ShaderBinaryInfo() *GTMioShaderBinaryInfo {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("shaderBinaryInfo"))
	return (*GTMioShaderBinaryInfo)(rv)
}
func (o GTMioTraceDataProtocolObject) ShaderBinaryInfoCount() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("shaderBinaryInfoCount"))
	return rv
}
func (o GTMioTraceDataProtocolObject) SignpostPipelineStateCount() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("signpostPipelineStateCount"))
	return rv
}
func (o GTMioTraceDataProtocolObject) SignpostPipelineStates() *GTMioShaderTimelineSignpostPipelineState {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("signpostPipelineStates"))
	return (*GTMioShaderTimelineSignpostPipelineState)(rv)
}
func (o GTMioTraceDataProtocolObject) SignpostProcessCount() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("signpostProcessCount"))
	return rv
}
func (o GTMioTraceDataProtocolObject) SignpostProcesses() *GTMioShaderTimelineSignpostProcess {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("signpostProcesses"))
	return (*GTMioShaderTimelineSignpostProcess)(rv)
}
func (o GTMioTraceDataProtocolObject) SignpostShaderCount() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("signpostShaderCount"))
	return rv
}
func (o GTMioTraceDataProtocolObject) SignpostShaders() *GTMioShaderTimelineSignpostShader {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("signpostShaders"))
	return (*GTMioShaderTimelineSignpostShader)(rv)
}
func (o GTMioTraceDataProtocolObject) SignpostStrings() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("signpostStrings"))
	return objectivec.Object{ID: rv}
}
func (o GTMioTraceDataProtocolObject) StreamData() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("streamData"))
	return objectivec.Object{ID: rv}
}
func (o GTMioTraceDataProtocolObject) TimelineCounters() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("timelineCounters"))
	return objectivec.Object{ID: rv}
}
func (o GTMioTraceDataProtocolObject) TimelineDuration() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("timelineDuration"))
	return rv
}
func (o GTMioTraceDataProtocolObject) TimestampBegin() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("timestampBegin"))
	return rv
}
func (o GTMioTraceDataProtocolObject) TimestampEnd() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("timestampEnd"))
	return rv
}
func (o GTMioTraceDataProtocolObject) TotalCliqueCost() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("totalCliqueCost"))
	return rv
}
func (o GTMioTraceDataProtocolObject) TotalCores() uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("totalCores"))
	return rv
}
func (o GTMioTraceDataProtocolObject) TotalCostForScopeScopeIdentifierProgramType(scope uint16, identifier uint64, type_ uint16) float64 {
	rv := objc.SendIfResponds[float64](o.ID, objc.Sel("totalCostForScope:scopeIdentifier:programType:"), scope, identifier, type_)
	return rv
}
func (o GTMioTraceDataProtocolObject) Uscs() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("uscs"))
	return objectivec.Object{ID: rv}
}

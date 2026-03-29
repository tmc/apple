// Code generated from Apple documentation. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
)

// GTMioShaderBinaryDataHandler is the signature for a completion handler block.
//
// Used by:
//   - [GTMioTraceDataProtocol.EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateBinariesForPipelineStateEnumerator]
type GTMioShaderBinaryDataHandler = func(*GTMioShaderBinaryData)

// NewGTMioShaderBinaryDataBlock wraps a Go [GTMioShaderBinaryDataHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [GTMioTraceDataProtocol.EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateBinariesForPipelineStateEnumerator]
func NewGTMioShaderBinaryDataBlock(handler GTMioShaderBinaryDataHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *GTMioShaderBinaryData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := GTMioShaderBinaryDataFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// GTMioShaderBinaryInfoGTMioShaderBinaryDataHandler is the signature for a completion handler block.
//
// Used by:
//   - [GTMioTraceDataProtocol.EnumerateBinariesForDrawEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateBinariesForEncoderEnumerator]
type GTMioShaderBinaryInfoGTMioShaderBinaryDataHandler = func(*uintptr, *GTMioShaderBinaryData)

// GTMioUSCCliqueMetadataGTMioUSCTraceDataHandler is the signature for a completion handler block.
//
// Used by:
//   - [GTMioTraceDataProtocol.EnumerateUniqueTracesForBinaryEnumerator]
type GTMioUSCCliqueMetadataGTMioUSCTraceDataHandler = func(*uintptr, *GTMioUSCTraceData)

// UintHandler handles completion with a primitive value.
//
// Used by:
//   - [GTMioTraceDataProtocol.EnumerateDrawsForEncoderEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateDrawsForPipelineStateEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateEncoders]
type UintHandler = func(uint)

// NewUintBlock wraps a Go [UintHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [GTMioTraceDataProtocol.EnumerateDrawsForEncoderEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateDrawsForPipelineStateEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateEncoders]
func NewUintBlock(handler UintHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal uint) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler is the signature for a completion handler block.
//
// Used by:
//   - [DYGPUDerivedEncoderCounterInfo._enumerateEncoderDerivedDataAtIndexWithBlock]
//   - [DYGPUDerivedEncoderCounterInfo._enumerateEncoderDerivedData]
//   - [DYGPUTimelineInfo.EnumerateActiveShadersForAllSamples]
//   - [DYGPUTimelineInfo.EnumerateActiveShadersForSampleAtIndexWithBlock]
//   - [DYWorkloadGPUTimelineInfo.EnumerateEncoderDerivedDataAtIndexWithBlock]
//   - [DYWorkloadGPUTimelineInfo.EnumerateEncoderDerivedData]
//   - [GTJSScriptingContext.SetExceptionHandler]
//   - [GTMioEncoderQuadData.BuildEncoderFunctionIndexCliqueFilter]
//   - [GTMioEncoderQuadData.EnumerateCliquesForQuadEnumerator]
//   - [GTMioEncoderQuadData.EnumerateCliquesForQuadLocationEnumerator]
//   - [GTMioEncoderQuadData.EnumerateOrderedQuads]
//   - [GTMioEncoderQuadData._buildCliquesEncoderFunctionIndexProgramTypeCliqueFilter]
//   - [GTMioEncoderQuadData._buildComputeEncoderFunctionIndexProgramTypeCliqueFilter]
//   - [GTMioEncoderQuadData._buildFragmentEncoderFunctionIndexProgramTypeCliqueFilter]
//   - [GTMioShaderBinaryData.EnumerateBinaryRangesForFileLineEnumerator]
//   - [GTMioShaderBinaryData.EnumerateDrawsWithProgramTypeEnumerator]
//   - [GTMioShaderBinaryData.EnumerateEncoderCosts]
//   - [GTMioShaderBinaryData.EnumerateEntryPoints]
//   - [GTMioShaderBinaryData.EnumerateInstructionsForBinaryRangeEnumerator]
//   - [GTMioShaderBinaryData.EnumerateLinesForFileEnumerator]
//   - [GTMioShaderBinaryData.EnumeratePerDrawCosts]
//   - [GTMioShaderBinaryData.EnumeratePipelineStateCosts]
//   - [GTMioShaderBinaryData.EnumeratePipelineStatesWithProgramTypeEnumerator]
//   - [GTMioShaderBinaryData.EnumerateTraces]
//   - [GTMioShaderExecutionHistoryFunctionNode.EnumerateInstructions]
//   - [GTMioShaderExecutionHistoryNode.Dfs]
//   - [GTMioShaderExecutionHistoryNode.EnumerateInstructions]
//   - [GTMioShaderExecutionHistoryNode._dfsEnumerator]
//   - [GTMioShaderExecutionHistoryRootNode.EnumerateFunctionCallSites]
//   - [GTMioTraceData.EnumerateBinariesForDrawEnumerator]
//   - [GTMioTraceData.EnumerateBinariesForEncoderEnumerator]
//   - [GTMioTraceData.EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceData.EnumerateBinariesForPipelineStateEnumerator]
//   - [GTMioTraceData.EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceData.EnumerateBinaryRangesForCliqueUscDataEnumerator]
//   - [GTMioTraceData.EnumerateDrawsForEncoderEnumerator]
//   - [GTMioTraceData.EnumerateDrawsForPipelineStateEnumerator]
//   - [GTMioTraceData.EnumerateEncoders]
//   - [GTMioTraceData.EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceData.EnumerateInstructionsForCliqueUscDataEnumerator]
//   - [GTMioTraceData.EnumerateKickAtFunctionIndexEnumerator]
//   - [GTMioTraceData.EnumeratePipelineStates]
//   - [GTMioTraceData.EnumerateUniqueTracesForBinaryEnumerator]
//   - [GTMioTraceData.InitWithTraceDatabaseDeallocator]
//   - [GTMioTraceData.RequestCostTimeline]
//   - [GTMioTraceDataHelper.GenerateTrackForCliqueIndexesCountGroup]
//   - [GTMioTraceDataProtocol.EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateBinaryRangesForCliqueUscDataEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateInstructionsForCliqueUscDataEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateKickAtFunctionIndexEnumerator]
//   - [GTMioTraceTimelineData.EnumerateBinariesForDrawEnumerator]
//   - [GTMioTraceTimelineData.EnumerateBinariesForEncoderEnumerator]
//   - [GTMioTraceTimelineData.EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceTimelineData.EnumerateBinariesForPipelineStateEnumerator]
//   - [GTMioTraceTimelineData.EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceTimelineData.EnumerateBinaryRangesForCliqueUscDataEnumerator]
//   - [GTMioTraceTimelineData.EnumerateDrawsForEncoderEnumerator]
//   - [GTMioTraceTimelineData.EnumerateDrawsForPipelineStateEnumerator]
//   - [GTMioTraceTimelineData.EnumerateEncoders]
//   - [GTMioTraceTimelineData.EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceTimelineData.EnumerateInstructionsForCliqueUscDataEnumerator]
//   - [GTMioTraceTimelineData.EnumerateKickAtFunctionIndexEnumerator]
//   - [GTMioTraceTimelineData.EnumeratePipelineStates]
//   - [GTMioTraceTimelineData.EnumerateUniqueTracesForBinaryEnumerator]
//   - [GTMioTraceTimelineData.ExecutionHistoryForPipelineStateDelegateProgressControllerCliqueFilter]
//   - [GTMioTraceTimelineData.InitWithTraceDatabaseDeallocatorParentData]
//   - [GTMioUSCTraceData.EnumerateCliquesForTimeRangeBeginEndEnumerator]
//   - [GTMioUSCTraceData.EnumerateInstructionTracesForCliqueRequiresTimestampsEnumerator]
//   - [GTMioUSCTraceData.EnumerateKickAtFunctionIndexEnumerator]
//   - [GTMioUSCTraceData.EnumerateKickCliquesAtFunctionIndexDataMasterEnumerator]
//   - [GTMioUSCTraceData.EnumerateKickCliquesEnumerator]
//   - [GTMioUSCTraceData.EnumerateKickTilesAtFunctionIndexDataMasterEnumerator]
//   - [GTMioUSCTraceData.EnumerateKickTilesEnumerator]
//   - [GTMioUSCTraceData.EnumerateTileCliquesEnumerator]
//   - [GTShaderProfilerAnalyzer.GenerateFullMCAReport]
//   - [GTShaderProfilerAnalyzer.GenerateMCAOutputCallback]
//   - [GTShaderProfilerAnalyzer.GenerateRegisterPressureView]
//   - [GTShaderProfilerStreamData.EnumerateUnarchivedBatchIdFilteredCounterData]
//   - [GTShaderProfilerStreamData.EnumerateUnarchivedGPUTimelineData]
//   - [GTShaderProfilerStreamData.EnumerateUnarchivedShaderProfilerData]
//   - [XRGPUAGXShaderTimelineSignposts.EnumerateKickIds]
//   - [XRGPUAPSDataContainer.EnumerateRDEData]
//   - [XRGPUAPSDataContainer.EnumerateUSCData]
//   - [XRGPUAPSDataProcessor.EnumerateShaders]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [DYGPUDerivedEncoderCounterInfo._enumerateEncoderDerivedDataAtIndexWithBlock]
//   - [DYGPUDerivedEncoderCounterInfo._enumerateEncoderDerivedData]
//   - [DYGPUTimelineInfo.EnumerateActiveShadersForAllSamples]
//   - [DYGPUTimelineInfo.EnumerateActiveShadersForSampleAtIndexWithBlock]
//   - [DYWorkloadGPUTimelineInfo.EnumerateEncoderDerivedDataAtIndexWithBlock]
//   - [DYWorkloadGPUTimelineInfo.EnumerateEncoderDerivedData]
//   - [GTJSScriptingContext.SetExceptionHandler]
//   - [GTMioEncoderQuadData.BuildEncoderFunctionIndexCliqueFilter]
//   - [GTMioEncoderQuadData.EnumerateCliquesForQuadEnumerator]
//   - [GTMioEncoderQuadData.EnumerateCliquesForQuadLocationEnumerator]
//   - [GTMioEncoderQuadData.EnumerateOrderedQuads]
//   - [GTMioEncoderQuadData._buildCliquesEncoderFunctionIndexProgramTypeCliqueFilter]
//   - [GTMioEncoderQuadData._buildComputeEncoderFunctionIndexProgramTypeCliqueFilter]
//   - [GTMioEncoderQuadData._buildFragmentEncoderFunctionIndexProgramTypeCliqueFilter]
//   - [GTMioShaderBinaryData.EnumerateBinaryRangesForFileLineEnumerator]
//   - [GTMioShaderBinaryData.EnumerateDrawsWithProgramTypeEnumerator]
//   - [GTMioShaderBinaryData.EnumerateEncoderCosts]
//   - [GTMioShaderBinaryData.EnumerateEntryPoints]
//   - [GTMioShaderBinaryData.EnumerateInstructionsForBinaryRangeEnumerator]
//   - [GTMioShaderBinaryData.EnumerateLinesForFileEnumerator]
//   - [GTMioShaderBinaryData.EnumeratePerDrawCosts]
//   - [GTMioShaderBinaryData.EnumeratePipelineStateCosts]
//   - [GTMioShaderBinaryData.EnumeratePipelineStatesWithProgramTypeEnumerator]
//   - [GTMioShaderBinaryData.EnumerateTraces]
//   - [GTMioShaderExecutionHistoryFunctionNode.EnumerateInstructions]
//   - [GTMioShaderExecutionHistoryNode.Dfs]
//   - [GTMioShaderExecutionHistoryNode.EnumerateInstructions]
//   - [GTMioShaderExecutionHistoryNode._dfsEnumerator]
//   - [GTMioShaderExecutionHistoryRootNode.EnumerateFunctionCallSites]
//   - [GTMioTraceData.EnumerateBinariesForDrawEnumerator]
//   - [GTMioTraceData.EnumerateBinariesForEncoderEnumerator]
//   - [GTMioTraceData.EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceData.EnumerateBinariesForPipelineStateEnumerator]
//   - [GTMioTraceData.EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceData.EnumerateBinaryRangesForCliqueUscDataEnumerator]
//   - [GTMioTraceData.EnumerateDrawsForEncoderEnumerator]
//   - [GTMioTraceData.EnumerateDrawsForPipelineStateEnumerator]
//   - [GTMioTraceData.EnumerateEncoders]
//   - [GTMioTraceData.EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceData.EnumerateInstructionsForCliqueUscDataEnumerator]
//   - [GTMioTraceData.EnumerateKickAtFunctionIndexEnumerator]
//   - [GTMioTraceData.EnumeratePipelineStates]
//   - [GTMioTraceData.EnumerateUniqueTracesForBinaryEnumerator]
//   - [GTMioTraceData.InitWithTraceDatabaseDeallocator]
//   - [GTMioTraceData.RequestCostTimeline]
//   - [GTMioTraceDataHelper.GenerateTrackForCliqueIndexesCountGroup]
//   - [GTMioTraceDataProtocol.EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateBinaryRangesForCliqueUscDataEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateInstructionsForCliqueUscDataEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateKickAtFunctionIndexEnumerator]
//   - [GTMioTraceTimelineData.EnumerateBinariesForDrawEnumerator]
//   - [GTMioTraceTimelineData.EnumerateBinariesForEncoderEnumerator]
//   - [GTMioTraceTimelineData.EnumerateBinariesForForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceTimelineData.EnumerateBinariesForPipelineStateEnumerator]
//   - [GTMioTraceTimelineData.EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceTimelineData.EnumerateBinaryRangesForCliqueUscDataEnumerator]
//   - [GTMioTraceTimelineData.EnumerateDrawsForEncoderEnumerator]
//   - [GTMioTraceTimelineData.EnumerateDrawsForPipelineStateEnumerator]
//   - [GTMioTraceTimelineData.EnumerateEncoders]
//   - [GTMioTraceTimelineData.EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceTimelineData.EnumerateInstructionsForCliqueUscDataEnumerator]
//   - [GTMioTraceTimelineData.EnumerateKickAtFunctionIndexEnumerator]
//   - [GTMioTraceTimelineData.EnumeratePipelineStates]
//   - [GTMioTraceTimelineData.EnumerateUniqueTracesForBinaryEnumerator]
//   - [GTMioTraceTimelineData.ExecutionHistoryForPipelineStateDelegateProgressControllerCliqueFilter]
//   - [GTMioTraceTimelineData.InitWithTraceDatabaseDeallocatorParentData]
//   - [GTMioUSCTraceData.EnumerateCliquesForTimeRangeBeginEndEnumerator]
//   - [GTMioUSCTraceData.EnumerateInstructionTracesForCliqueRequiresTimestampsEnumerator]
//   - [GTMioUSCTraceData.EnumerateKickAtFunctionIndexEnumerator]
//   - [GTMioUSCTraceData.EnumerateKickCliquesAtFunctionIndexDataMasterEnumerator]
//   - [GTMioUSCTraceData.EnumerateKickCliquesEnumerator]
//   - [GTMioUSCTraceData.EnumerateKickTilesAtFunctionIndexDataMasterEnumerator]
//   - [GTMioUSCTraceData.EnumerateKickTilesEnumerator]
//   - [GTMioUSCTraceData.EnumerateTileCliquesEnumerator]
//   - [GTShaderProfilerAnalyzer.GenerateFullMCAReport]
//   - [GTShaderProfilerAnalyzer.GenerateMCAOutputCallback]
//   - [GTShaderProfilerAnalyzer.GenerateRegisterPressureView]
//   - [GTShaderProfilerStreamData.EnumerateUnarchivedBatchIdFilteredCounterData]
//   - [GTShaderProfilerStreamData.EnumerateUnarchivedGPUTimelineData]
//   - [GTShaderProfilerStreamData.EnumerateUnarchivedShaderProfilerData]
//   - [XRGPUAGXShaderTimelineSignposts.EnumerateKickIds]
//   - [XRGPUAPSDataContainer.EnumerateRDEData]
//   - [XRGPUAPSDataContainer.EnumerateUSCData]
//   - [XRGPUAPSDataProcessor.EnumerateShaders]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// unsignedlongHandler handles completion with a primitive value.
//
// Used by:
//   - [GTMioTraceDataProtocol.EnumeratePipelineStates]
type unsignedlongHandler = func(uint64)

// NewunsignedlongBlock wraps a Go [unsignedlongHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [GTMioTraceDataProtocol.EnumeratePipelineStates]
func NewunsignedlongBlock(handler unsignedlongHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal uint64) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}


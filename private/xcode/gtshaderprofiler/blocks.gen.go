// Code generated from Apple documentation. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
)

// GTMioShaderBinaryDataGTMioShaderBinaryDebugBinaryRangeUintHandler is the signature for a completion handler block.
//
// Used by:
//   - [GTMioTraceDataProtocol.EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateBinaryRangesForCliqueUscDataEnumerator]
type GTMioShaderBinaryDataGTMioShaderBinaryDebugBinaryRangeUintHandler = func(*GTMioShaderBinaryData, *GTMioShaderBinaryDebugBinaryRange, uint)

// NewGTMioShaderBinaryDataGTMioShaderBinaryDebugBinaryRangeUintBlock wraps a Go [GTMioShaderBinaryDataGTMioShaderBinaryDebugBinaryRangeUintHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [GTMioTraceDataProtocol.EnumerateBinaryRangesForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateBinaryRangesForCliqueUscDataEnumerator]
func NewGTMioShaderBinaryDataGTMioShaderBinaryDebugBinaryRangeUintBlock(handler GTMioShaderBinaryDataGTMioShaderBinaryDebugBinaryRangeUintHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 *GTMioShaderBinaryDebugBinaryRange, extra1 uint) {
		var result *GTMioShaderBinaryData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := GTMioShaderBinaryDataFromID(resultID)
			result = &v
		}
		handler(result, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// GTMioShaderBinaryDataGTMioShaderInstructionInfoUintHandler is the signature for a completion handler block.
//
// Used by:
//   - [GTMioTraceDataProtocol.EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateInstructionsForCliqueUscDataEnumerator]
type GTMioShaderBinaryDataGTMioShaderInstructionInfoUintHandler = func(*GTMioShaderBinaryData, *GTMioShaderInstructionInfo, uint)

// NewGTMioShaderBinaryDataGTMioShaderInstructionInfoUintBlock wraps a Go [GTMioShaderBinaryDataGTMioShaderInstructionInfoUintHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [GTMioTraceDataProtocol.EnumerateInstructionsForCliqueAtIndexUscIndexEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateInstructionsForCliqueUscDataEnumerator]
func NewGTMioShaderBinaryDataGTMioShaderInstructionInfoUintBlock(handler GTMioShaderBinaryDataGTMioShaderInstructionInfoUintHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 *GTMioShaderInstructionInfo, extra1 uint) {
		var result *GTMioShaderBinaryData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := GTMioShaderBinaryDataFromID(resultID)
			result = &v
		}
		handler(result, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

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
	if handler == nil {
		return 0, func() {}
	}
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

// GTMioShaderBinaryInfoGTMioShaderBinaryDataHandler handles completion with primitive and object results.
//
// Used by:
//   - [GTMioTraceDataProtocol.EnumerateBinariesForDrawEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateBinariesForEncoderEnumerator]
type GTMioShaderBinaryInfoGTMioShaderBinaryDataHandler = func(*GTMioShaderBinaryInfo, *GTMioShaderBinaryData)

// NewGTMioShaderBinaryInfoGTMioShaderBinaryDataBlock wraps a Go [GTMioShaderBinaryInfoGTMioShaderBinaryDataHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [GTMioTraceDataProtocol.EnumerateBinariesForDrawEnumerator]
//   - [GTMioTraceDataProtocol.EnumerateBinariesForEncoderEnumerator]
func NewGTMioShaderBinaryInfoGTMioShaderBinaryDataBlock(handler GTMioShaderBinaryInfoGTMioShaderBinaryDataHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *GTMioShaderBinaryInfo, extra0ID objc.ID) {
		var extra0 *GTMioShaderBinaryData
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := GTMioShaderBinaryDataFromID(extra0ID)
			extra0 = &v
		}
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// GTMioUSCCliqueMetadataGTMioUSCTraceDataHandler handles completion with primitive and object results.
//
// Used by:
//   - [GTMioTraceDataProtocol.EnumerateUniqueTracesForBinaryEnumerator]
type GTMioUSCCliqueMetadataGTMioUSCTraceDataHandler = func(*GTMioUSCCliqueMetadata, *GTMioUSCTraceData)

// NewGTMioUSCCliqueMetadataGTMioUSCTraceDataBlock wraps a Go [GTMioUSCCliqueMetadataGTMioUSCTraceDataHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [GTMioTraceDataProtocol.EnumerateUniqueTracesForBinaryEnumerator]
func NewGTMioUSCCliqueMetadataGTMioUSCTraceDataBlock(handler GTMioUSCCliqueMetadataGTMioUSCTraceDataHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *GTMioUSCCliqueMetadata, extra0ID objc.ID) {
		var extra0 *GTMioUSCTraceData
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := GTMioUSCTraceDataFromID(extra0ID)
			extra0 = &v
		}
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// GTMioUSCTraceDataGTMioUSCKickMetadataUint64Handler is the signature for a completion handler block.
//
// Used by:
//   - [GTMioTraceDataProtocol.EnumerateKickAtFunctionIndexEnumerator]
type GTMioUSCTraceDataGTMioUSCKickMetadataUint64Handler = func(*GTMioUSCTraceData, *GTMioUSCKickMetadata, uint64)

// NewGTMioUSCTraceDataGTMioUSCKickMetadataUint64Block wraps a Go [GTMioUSCTraceDataGTMioUSCKickMetadataUint64Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [GTMioTraceDataProtocol.EnumerateKickAtFunctionIndexEnumerator]
func NewGTMioUSCTraceDataGTMioUSCKickMetadataUint64Block(handler GTMioUSCTraceDataGTMioUSCKickMetadataUint64Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 *GTMioUSCKickMetadata, extra1 uint64) {
		var result *GTMioUSCTraceData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := GTMioUSCTraceDataFromID(resultID)
			result = &v
		}
		handler(result, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

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
	if handler == nil {
		return 0, func() {}
	}
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
	if handler == nil {
		return 0, func() {}
	}
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal uint64) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

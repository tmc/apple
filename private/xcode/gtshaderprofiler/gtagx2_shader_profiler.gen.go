// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTAGX2ShaderProfiler] class.
var (
	_GTAGX2ShaderProfilerClass     GTAGX2ShaderProfilerClass
	_GTAGX2ShaderProfilerClassOnce sync.Once
)

func getGTAGX2ShaderProfilerClass() GTAGX2ShaderProfilerClass {
	_GTAGX2ShaderProfilerClassOnce.Do(func() {
		_GTAGX2ShaderProfilerClass = GTAGX2ShaderProfilerClass{class: objc.GetClass("GTAGX2ShaderProfiler")}
	})
	return _GTAGX2ShaderProfilerClass
}

// GetGTAGX2ShaderProfilerClass returns the class object for GTAGX2ShaderProfiler.
func GetGTAGX2ShaderProfilerClass() GTAGX2ShaderProfilerClass {
	return getGTAGX2ShaderProfilerClass()
}

type GTAGX2ShaderProfilerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTAGX2ShaderProfilerClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTAGX2ShaderProfilerClass) Alloc() GTAGX2ShaderProfiler {
	rv := objc.Send[GTAGX2ShaderProfiler](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTAGX2ShaderProfiler._analyzeShaderBinaries]
//   - [GTAGX2ShaderProfiler._calculateAggregatedEffectiveGPUEncoderCost]
//   - [GTAGX2ShaderProfiler._calculateAndAppendPerBatchCostsForFrameIndexAtTimestampWithLimiterIndicesWithTimestampBuffersWithActiveBatchesWithPerRingBufferLimiterEncoderCosts]
//   - [GTAGX2ShaderProfiler._calculateAverageGPUCommandDuration]
//   - [GTAGX2ShaderProfiler._calculateEffectiveGPUEncoderCostForFrameIndex]
//   - [GTAGX2ShaderProfiler._calculatePerDrawCostsResult]
//   - [GTAGX2ShaderProfiler._cleanup]
//   - [GTAGX2ShaderProfiler._computeSampleNormFactorForSampleForProgramStartAddressForProgramEndAddress]
//   - [GTAGX2ShaderProfiler._conservativeLatencyAdjustmentWithLimiter]
//   - [GTAGX2ShaderProfiler._dumpAggregatedGPUTimePerBatchForFrame]
//   - [GTAGX2ShaderProfiler._dumpLimiterBatchInfoCostsForRingBufferForFrame]
//   - [GTAGX2ShaderProfiler._dumpShaderBinaries]
//   - [GTAGX2ShaderProfiler._dumpTraceBufferPacketsInFileUsingTracePacketsWithExtractedSamplesWithTraceBufferCountForTargetIndex]
//   - [GTAGX2ShaderProfiler._initiateConnectionToLLVMHelper]
//   - [GTAGX2ShaderProfiler._latencyAdjustmentEstimateWithLimiter]
//   - [GTAGX2ShaderProfiler._latencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterType]
//   - [GTAGX2ShaderProfiler._latencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterTypeLoadForLimiterTypeStore]
//   - [GTAGX2ShaderProfiler._latencyAdjustmentFactorWithLimiterDataWithLerpForIndexWithLimiterTypeIndexMapForLimiterType]
//   - [GTAGX2ShaderProfiler._latencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap]
//   - [GTAGX2ShaderProfiler._mergeLegacyAndNewShadersInShaderInfoDictionary]
//   - [GTAGX2ShaderProfiler._processTracePacketsForRenderIndexAndGenerateSampleListForTargetIndexForLimiterIndex]
//   - [GTAGX2ShaderProfiler._setupShaderBinaryInfoWithBinaryKeyAndNumDraws]
//   - [GTAGX2ShaderProfiler._waitLatencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap]
//   - [GTAGX2ShaderProfiler.AdjustHWBiasAndFinalizeResultBlitKickIndicesBlitTimesResult]
//   - [GTAGX2ShaderProfiler.AveragePerDrawKickDurations]
//   - [GTAGX2ShaderProfiler.DumpShaderBinarySamples]
//   - [GTAGX2ShaderProfiler.EffectiveKickTimes]
//   - [GTAGX2ShaderProfiler.EvaluateStreamingSamplesWithUSCSampleNumWithProgramAddressLUTTargetIndexForRingBufferIndexWithMinEncoderIndexWithMaxEncoderIndexWithEncoderIdToEncoderIndexMapWithProfilingData]
//   - [GTAGX2ShaderProfiler.HandleHarvestedBinaryInfo]
//   - [GTAGX2ShaderProfiler.IsaPrinter]
//   - [GTAGX2ShaderProfiler.SetIsaPrinter]
//   - [GTAGX2ShaderProfiler.LoadActionTimes]
//   - [GTAGX2ShaderProfiler.PerRingPerFrameLimiterData]
//   - [GTAGX2ShaderProfiler.SetRingBufferCount]
//   - [GTAGX2ShaderProfiler.StoreActionTimes]
//   - [GTAGX2ShaderProfiler.TimingInfo]
//   - [GTAGX2ShaderProfiler.InitWithStreamDataForTargetIndex]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler
type GTAGX2ShaderProfiler struct {
	objectivec.Object
}

// GTAGX2ShaderProfilerFromID constructs a [GTAGX2ShaderProfiler] from an objc.ID.
func GTAGX2ShaderProfilerFromID(id objc.ID) GTAGX2ShaderProfiler {
	return GTAGX2ShaderProfiler{objectivec.Object{ID: id}}
}

// Ensure GTAGX2ShaderProfiler implements IGTAGX2ShaderProfiler.
var _ IGTAGX2ShaderProfiler = GTAGX2ShaderProfiler{}

// An interface definition for the [GTAGX2ShaderProfiler] class.
//
// # Methods
//
//   - [IGTAGX2ShaderProfiler._analyzeShaderBinaries]
//   - [IGTAGX2ShaderProfiler._calculateAggregatedEffectiveGPUEncoderCost]
//   - [IGTAGX2ShaderProfiler._calculateAndAppendPerBatchCostsForFrameIndexAtTimestampWithLimiterIndicesWithTimestampBuffersWithActiveBatchesWithPerRingBufferLimiterEncoderCosts]
//   - [IGTAGX2ShaderProfiler._calculateAverageGPUCommandDuration]
//   - [IGTAGX2ShaderProfiler._calculateEffectiveGPUEncoderCostForFrameIndex]
//   - [IGTAGX2ShaderProfiler._calculatePerDrawCostsResult]
//   - [IGTAGX2ShaderProfiler._cleanup]
//   - [IGTAGX2ShaderProfiler._computeSampleNormFactorForSampleForProgramStartAddressForProgramEndAddress]
//   - [IGTAGX2ShaderProfiler._conservativeLatencyAdjustmentWithLimiter]
//   - [IGTAGX2ShaderProfiler._dumpAggregatedGPUTimePerBatchForFrame]
//   - [IGTAGX2ShaderProfiler._dumpLimiterBatchInfoCostsForRingBufferForFrame]
//   - [IGTAGX2ShaderProfiler._dumpShaderBinaries]
//   - [IGTAGX2ShaderProfiler._dumpTraceBufferPacketsInFileUsingTracePacketsWithExtractedSamplesWithTraceBufferCountForTargetIndex]
//   - [IGTAGX2ShaderProfiler._initiateConnectionToLLVMHelper]
//   - [IGTAGX2ShaderProfiler._latencyAdjustmentEstimateWithLimiter]
//   - [IGTAGX2ShaderProfiler._latencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterType]
//   - [IGTAGX2ShaderProfiler._latencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterTypeLoadForLimiterTypeStore]
//   - [IGTAGX2ShaderProfiler._latencyAdjustmentFactorWithLimiterDataWithLerpForIndexWithLimiterTypeIndexMapForLimiterType]
//   - [IGTAGX2ShaderProfiler._latencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap]
//   - [IGTAGX2ShaderProfiler._mergeLegacyAndNewShadersInShaderInfoDictionary]
//   - [IGTAGX2ShaderProfiler._processTracePacketsForRenderIndexAndGenerateSampleListForTargetIndexForLimiterIndex]
//   - [IGTAGX2ShaderProfiler._setupShaderBinaryInfoWithBinaryKeyAndNumDraws]
//   - [IGTAGX2ShaderProfiler._waitLatencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap]
//   - [IGTAGX2ShaderProfiler.AdjustHWBiasAndFinalizeResultBlitKickIndicesBlitTimesResult]
//   - [IGTAGX2ShaderProfiler.AveragePerDrawKickDurations]
//   - [IGTAGX2ShaderProfiler.DumpShaderBinarySamples]
//   - [IGTAGX2ShaderProfiler.EffectiveKickTimes]
//   - [IGTAGX2ShaderProfiler.EvaluateStreamingSamplesWithUSCSampleNumWithProgramAddressLUTTargetIndexForRingBufferIndexWithMinEncoderIndexWithMaxEncoderIndexWithEncoderIdToEncoderIndexMapWithProfilingData]
//   - [IGTAGX2ShaderProfiler.HandleHarvestedBinaryInfo]
//   - [IGTAGX2ShaderProfiler.IsaPrinter]
//   - [IGTAGX2ShaderProfiler.SetIsaPrinter]
//   - [IGTAGX2ShaderProfiler.LoadActionTimes]
//   - [IGTAGX2ShaderProfiler.PerRingPerFrameLimiterData]
//   - [IGTAGX2ShaderProfiler.SetRingBufferCount]
//   - [IGTAGX2ShaderProfiler.StoreActionTimes]
//   - [IGTAGX2ShaderProfiler.TimingInfo]
//   - [IGTAGX2ShaderProfiler.InitWithStreamDataForTargetIndex]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler
type IGTAGX2ShaderProfiler interface {
	objectivec.IObject

	// Topic: Methods

	_analyzeShaderBinaries()
	_calculateAggregatedEffectiveGPUEncoderCost()
	_calculateAndAppendPerBatchCostsForFrameIndexAtTimestampWithLimiterIndicesWithTimestampBuffersWithActiveBatchesWithPerRingBufferLimiterEncoderCosts(costs unsafe.Pointer, index uint32, timestamp uint64, indices unsafe.Pointer, buffers unsafe.Pointer, batches unsafe.Pointer, costs2 unsafe.Pointer)
	_calculateAverageGPUCommandDuration()
	_calculateEffectiveGPUEncoderCostForFrameIndex(index uint32)
	_calculatePerDrawCostsResult(costs objectivec.IObject, result objectivec.IObject)
	_cleanup()
	_computeSampleNormFactorForSampleForProgramStartAddressForProgramEndAddress(sample ShaderProfilerUSCSampleInfo, address uint64, address2 uint64) uint32
	_conservativeLatencyAdjustmentWithLimiter(adjustment float64, limiter float64) float64
	_dumpAggregatedGPUTimePerBatchForFrame(frame uint32)
	_dumpLimiterBatchInfoCostsForRingBufferForFrame(costs unsafe.Pointer, buffer uint32, frame uint32)
	_dumpShaderBinaries()
	_dumpTraceBufferPacketsInFileUsingTracePacketsWithExtractedSamplesWithTraceBufferCountForTargetIndex(file objectivec.IObject, packets unsafe.Pointer, samples unsafe.Pointer, count uint32, index int)
	_initiateConnectionToLLVMHelper() bool
	_latencyAdjustmentEstimateWithLimiter(estimate float64, limiter float64) float64
	_latencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterType(data []float64, index uint64, map_ unsafe.Pointer, type_ uint32) float64
	_latencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterTypeLoadForLimiterTypeStore(data []float64, index uint64, map_ unsafe.Pointer, load uint32, store uint32) float64
	_latencyAdjustmentFactorWithLimiterDataWithLerpForIndexWithLimiterTypeIndexMapForLimiterType(lerp []float64, index uint64, map_ unsafe.Pointer, type_ uint32) float64
	_latencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap(data []float64, index uint64, map_ unsafe.Pointer) float64
	_mergeLegacyAndNewShadersInShaderInfoDictionary(dictionary objectivec.IObject) objectivec.IObject
	_processTracePacketsForRenderIndexAndGenerateSampleListForTargetIndexForLimiterIndex(packets unsafe.Pointer, index uint32, list unsafe.Pointer, index2 int, index3 uint32)
	_setupShaderBinaryInfoWithBinaryKeyAndNumDraws(info objectivec.IObject, key unsafe.Pointer, draws uint32)
	_waitLatencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap(data []float64, index uint64, map_ unsafe.Pointer) float64
	AdjustHWBiasAndFinalizeResultBlitKickIndicesBlitTimesResult(result objectivec.IObject, indices objectivec.IObject, times objectivec.IObject, result2 objectivec.IObject) float64
	AveragePerDrawKickDurations() foundation.INSArray
	DumpShaderBinarySamples(samples unsafe.Pointer)
	EffectiveKickTimes() foundation.INSArray
	EvaluateStreamingSamplesWithUSCSampleNumWithProgramAddressLUTTargetIndexForRingBufferIndexWithMinEncoderIndexWithMaxEncoderIndexWithEncoderIdToEncoderIndexMapWithProfilingData(samples unsafe.Pointer, num uint32, lut unsafe.Pointer, index int, index2 uint32, index3 uint32, index4 uint32, map_ unsafe.Pointer, data objectivec.IObject)
	HandleHarvestedBinaryInfo(info objectivec.IObject)
	IsaPrinter() objectivec.IObject
	SetIsaPrinter(value objectivec.IObject)
	LoadActionTimes() foundation.INSArray
	PerRingPerFrameLimiterData() foundation.INSDictionary
	SetRingBufferCount(count uint32)
	StoreActionTimes() foundation.INSArray
	TimingInfo() IGTShaderProfilerTimingInfo
	InitWithStreamDataForTargetIndex(data objectivec.IObject, index int) GTAGX2ShaderProfiler
}

// Init initializes the instance.
func (g GTAGX2ShaderProfiler) Init() GTAGX2ShaderProfiler {
	rv := objc.Send[GTAGX2ShaderProfiler](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTAGX2ShaderProfiler) Autorelease() GTAGX2ShaderProfiler {
	rv := objc.Send[GTAGX2ShaderProfiler](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTAGX2ShaderProfiler creates a new GTAGX2ShaderProfiler instance.
func NewGTAGX2ShaderProfiler() GTAGX2ShaderProfiler {
	class := getGTAGX2ShaderProfilerClass()
	rv := objc.Send[GTAGX2ShaderProfiler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/initWithStreamData:forTargetIndex:
func NewGTAGX2ShaderProfilerWithStreamDataForTargetIndex(data objectivec.IObject, index int) GTAGX2ShaderProfiler {
	instance := getGTAGX2ShaderProfilerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStreamData:forTargetIndex:"), data, index)
	return GTAGX2ShaderProfilerFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_analyzeShaderBinaries
func (g GTAGX2ShaderProfiler) _analyzeShaderBinaries() {
	objc.Send[objc.ID](g.ID, objc.Sel("_analyzeShaderBinaries"))
}

// AnalyzeShaderBinaries is an exported wrapper for the private method _analyzeShaderBinaries.
func (g GTAGX2ShaderProfiler) AnalyzeShaderBinaries() error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_analyzeShaderBinaries")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_analyzeShaderBinaries"}
		return err
	}
	g._analyzeShaderBinaries()
	return nil
}

// CanAnalyzeShaderBinaries reports whether the receiver responds to the private selector _analyzeShaderBinaries.
func (g GTAGX2ShaderProfiler) CanAnalyzeShaderBinaries() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_analyzeShaderBinaries"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_calculateAggregatedEffectiveGPUEncoderCost
func (g GTAGX2ShaderProfiler) _calculateAggregatedEffectiveGPUEncoderCost() {
	objc.Send[objc.ID](g.ID, objc.Sel("_calculateAggregatedEffectiveGPUEncoderCost"))
}

// CalculateAggregatedEffectiveGPUEncoderCost is an exported wrapper for the private method _calculateAggregatedEffectiveGPUEncoderCost.
func (g GTAGX2ShaderProfiler) CalculateAggregatedEffectiveGPUEncoderCost() error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_calculateAggregatedEffectiveGPUEncoderCost")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_calculateAggregatedEffectiveGPUEncoderCost"}
		return err
	}
	g._calculateAggregatedEffectiveGPUEncoderCost()
	return nil
}

// CanCalculateAggregatedEffectiveGPUEncoderCost reports whether the receiver responds to the private selector _calculateAggregatedEffectiveGPUEncoderCost.
func (g GTAGX2ShaderProfiler) CanCalculateAggregatedEffectiveGPUEncoderCost() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_calculateAggregatedEffectiveGPUEncoderCost"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_calculateAndAppendPerBatchCosts:forFrameIndex:atTimestamp:withLimiterIndices:withTimestampBuffers:withActiveBatches:withPerRingBufferLimiterEncoderCosts:
func (g GTAGX2ShaderProfiler) _calculateAndAppendPerBatchCostsForFrameIndexAtTimestampWithLimiterIndicesWithTimestampBuffersWithActiveBatchesWithPerRingBufferLimiterEncoderCosts(costs unsafe.Pointer, index uint32, timestamp uint64, indices unsafe.Pointer, buffers unsafe.Pointer, batches unsafe.Pointer, costs2 unsafe.Pointer) {
	objc.Send[objc.ID](g.ID, objc.Sel("_calculateAndAppendPerBatchCosts:forFrameIndex:atTimestamp:withLimiterIndices:withTimestampBuffers:withActiveBatches:withPerRingBufferLimiterEncoderCosts:"), costs, index, timestamp, indices, buffers, batches, costs2)
}

// CalculateAndAppendPerBatchCostsForFrameIndexAtTimestampWithLimiterIndicesWithTimestampBuffersWithActiveBatchesWithPerRingBufferLimiterEncoderCosts is an exported wrapper for the private method _calculateAndAppendPerBatchCostsForFrameIndexAtTimestampWithLimiterIndicesWithTimestampBuffersWithActiveBatchesWithPerRingBufferLimiterEncoderCosts.
func (g GTAGX2ShaderProfiler) CalculateAndAppendPerBatchCostsForFrameIndexAtTimestampWithLimiterIndicesWithTimestampBuffersWithActiveBatchesWithPerRingBufferLimiterEncoderCosts(costs unsafe.Pointer, index uint32, timestamp uint64, indices unsafe.Pointer, buffers unsafe.Pointer, batches unsafe.Pointer, costs2 unsafe.Pointer) error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_calculateAndAppendPerBatchCosts:forFrameIndex:atTimestamp:withLimiterIndices:withTimestampBuffers:withActiveBatches:withPerRingBufferLimiterEncoderCosts:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_calculateAndAppendPerBatchCosts:forFrameIndex:atTimestamp:withLimiterIndices:withTimestampBuffers:withActiveBatches:withPerRingBufferLimiterEncoderCosts:"}
		return err
	}
	g._calculateAndAppendPerBatchCostsForFrameIndexAtTimestampWithLimiterIndicesWithTimestampBuffersWithActiveBatchesWithPerRingBufferLimiterEncoderCosts(costs, index, timestamp, indices, buffers, batches, costs2)
	return nil
}

// CanCalculateAndAppendPerBatchCostsForFrameIndexAtTimestampWithLimiterIndicesWithTimestampBuffersWithActiveBatchesWithPerRingBufferLimiterEncoderCosts reports whether the receiver responds to the private selector _calculateAndAppendPerBatchCosts:forFrameIndex:atTimestamp:withLimiterIndices:withTimestampBuffers:withActiveBatches:withPerRingBufferLimiterEncoderCosts:.
func (g GTAGX2ShaderProfiler) CanCalculateAndAppendPerBatchCostsForFrameIndexAtTimestampWithLimiterIndicesWithTimestampBuffersWithActiveBatchesWithPerRingBufferLimiterEncoderCosts() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_calculateAndAppendPerBatchCosts:forFrameIndex:atTimestamp:withLimiterIndices:withTimestampBuffers:withActiveBatches:withPerRingBufferLimiterEncoderCosts:"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_calculateAverageGPUCommandDuration
func (g GTAGX2ShaderProfiler) _calculateAverageGPUCommandDuration() {
	objc.Send[objc.ID](g.ID, objc.Sel("_calculateAverageGPUCommandDuration"))
}

// CalculateAverageGPUCommandDuration is an exported wrapper for the private method _calculateAverageGPUCommandDuration.
func (g GTAGX2ShaderProfiler) CalculateAverageGPUCommandDuration() error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_calculateAverageGPUCommandDuration")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_calculateAverageGPUCommandDuration"}
		return err
	}
	g._calculateAverageGPUCommandDuration()
	return nil
}

// CanCalculateAverageGPUCommandDuration reports whether the receiver responds to the private selector _calculateAverageGPUCommandDuration.
func (g GTAGX2ShaderProfiler) CanCalculateAverageGPUCommandDuration() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_calculateAverageGPUCommandDuration"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_calculateEffectiveGPUEncoderCostForFrameIndex:
func (g GTAGX2ShaderProfiler) _calculateEffectiveGPUEncoderCostForFrameIndex(index uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("_calculateEffectiveGPUEncoderCostForFrameIndex:"), index)
}

// CalculateEffectiveGPUEncoderCostForFrameIndex is an exported wrapper for the private method _calculateEffectiveGPUEncoderCostForFrameIndex.
func (g GTAGX2ShaderProfiler) CalculateEffectiveGPUEncoderCostForFrameIndex(index uint32) error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_calculateEffectiveGPUEncoderCostForFrameIndex:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_calculateEffectiveGPUEncoderCostForFrameIndex:"}
		return err
	}
	g._calculateEffectiveGPUEncoderCostForFrameIndex(index)
	return nil
}

// CanCalculateEffectiveGPUEncoderCostForFrameIndex reports whether the receiver responds to the private selector _calculateEffectiveGPUEncoderCostForFrameIndex:.
func (g GTAGX2ShaderProfiler) CanCalculateEffectiveGPUEncoderCostForFrameIndex() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_calculateEffectiveGPUEncoderCostForFrameIndex:"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_calculatePerDrawCosts:result:
func (g GTAGX2ShaderProfiler) _calculatePerDrawCostsResult(costs objectivec.IObject, result objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("_calculatePerDrawCosts:result:"), costs, result)
}

// CalculatePerDrawCostsResult is an exported wrapper for the private method _calculatePerDrawCostsResult.
func (g GTAGX2ShaderProfiler) CalculatePerDrawCostsResult(costs objectivec.IObject, result objectivec.IObject) error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_calculatePerDrawCosts:result:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_calculatePerDrawCosts:result:"}
		return err
	}
	g._calculatePerDrawCostsResult(costs, result)
	return nil
}

// CanCalculatePerDrawCostsResult reports whether the receiver responds to the private selector _calculatePerDrawCosts:result:.
func (g GTAGX2ShaderProfiler) CanCalculatePerDrawCostsResult() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_calculatePerDrawCosts:result:"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_cleanup
func (g GTAGX2ShaderProfiler) _cleanup() {
	objc.Send[objc.ID](g.ID, objc.Sel("_cleanup"))
}

// Cleanup is an exported wrapper for the private method _cleanup.
func (g GTAGX2ShaderProfiler) Cleanup() error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_cleanup")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_cleanup"}
		return err
	}
	g._cleanup()
	return nil
}

// CanCleanup reports whether the receiver responds to the private selector _cleanup.
func (g GTAGX2ShaderProfiler) CanCleanup() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_cleanup"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_computeSampleNormFactorForSample:forProgramStartAddress:forProgramEndAddress:
func (g GTAGX2ShaderProfiler) _computeSampleNormFactorForSampleForProgramStartAddressForProgramEndAddress(sample ShaderProfilerUSCSampleInfo, address uint64, address2 uint64) uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("_computeSampleNormFactorForSample:forProgramStartAddress:forProgramEndAddress:"), sample, address, address2)
	return rv
}

// ComputeSampleNormFactorForSampleForProgramStartAddressForProgramEndAddress is an exported wrapper for the private method _computeSampleNormFactorForSampleForProgramStartAddressForProgramEndAddress.
func (g GTAGX2ShaderProfiler) ComputeSampleNormFactorForSampleForProgramStartAddressForProgramEndAddress(sample ShaderProfilerUSCSampleInfo, address uint64, address2 uint64) (uint32, error) {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_computeSampleNormFactorForSample:forProgramStartAddress:forProgramEndAddress:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_computeSampleNormFactorForSample:forProgramStartAddress:forProgramEndAddress:"}
		return 0, err
	}
	return g._computeSampleNormFactorForSampleForProgramStartAddressForProgramEndAddress(sample, address, address2), nil
}

// CanComputeSampleNormFactorForSampleForProgramStartAddressForProgramEndAddress reports whether the receiver responds to the private selector _computeSampleNormFactorForSample:forProgramStartAddress:forProgramEndAddress:.
func (g GTAGX2ShaderProfiler) CanComputeSampleNormFactorForSampleForProgramStartAddressForProgramEndAddress() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_computeSampleNormFactorForSample:forProgramStartAddress:forProgramEndAddress:"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_conservativeLatencyAdjustment:withLimiter:
func (g GTAGX2ShaderProfiler) _conservativeLatencyAdjustmentWithLimiter(adjustment float64, limiter float64) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("_conservativeLatencyAdjustment:withLimiter:"), adjustment, limiter)
	return rv
}

// ConservativeLatencyAdjustmentWithLimiter is an exported wrapper for the private method _conservativeLatencyAdjustmentWithLimiter.
func (g GTAGX2ShaderProfiler) ConservativeLatencyAdjustmentWithLimiter(adjustment float64, limiter float64) (float64, error) {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_conservativeLatencyAdjustment:withLimiter:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_conservativeLatencyAdjustment:withLimiter:"}
		return 0.0, err
	}
	return g._conservativeLatencyAdjustmentWithLimiter(adjustment, limiter), nil
}

// CanConservativeLatencyAdjustmentWithLimiter reports whether the receiver responds to the private selector _conservativeLatencyAdjustment:withLimiter:.
func (g GTAGX2ShaderProfiler) CanConservativeLatencyAdjustmentWithLimiter() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_conservativeLatencyAdjustment:withLimiter:"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_dumpAggregatedGPUTimePerBatchForFrame:
func (g GTAGX2ShaderProfiler) _dumpAggregatedGPUTimePerBatchForFrame(frame uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("_dumpAggregatedGPUTimePerBatchForFrame:"), frame)
}

// DumpAggregatedGPUTimePerBatchForFrame is an exported wrapper for the private method _dumpAggregatedGPUTimePerBatchForFrame.
func (g GTAGX2ShaderProfiler) DumpAggregatedGPUTimePerBatchForFrame(frame uint32) error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_dumpAggregatedGPUTimePerBatchForFrame:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_dumpAggregatedGPUTimePerBatchForFrame:"}
		return err
	}
	g._dumpAggregatedGPUTimePerBatchForFrame(frame)
	return nil
}

// CanDumpAggregatedGPUTimePerBatchForFrame reports whether the receiver responds to the private selector _dumpAggregatedGPUTimePerBatchForFrame:.
func (g GTAGX2ShaderProfiler) CanDumpAggregatedGPUTimePerBatchForFrame() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_dumpAggregatedGPUTimePerBatchForFrame:"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_dumpLimiterBatchInfoCosts:forRingBuffer:forFrame:
func (g GTAGX2ShaderProfiler) _dumpLimiterBatchInfoCostsForRingBufferForFrame(costs unsafe.Pointer, buffer uint32, frame uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("_dumpLimiterBatchInfoCosts:forRingBuffer:forFrame:"), costs, buffer, frame)
}

// DumpLimiterBatchInfoCostsForRingBufferForFrame is an exported wrapper for the private method _dumpLimiterBatchInfoCostsForRingBufferForFrame.
func (g GTAGX2ShaderProfiler) DumpLimiterBatchInfoCostsForRingBufferForFrame(costs unsafe.Pointer, buffer uint32, frame uint32) error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_dumpLimiterBatchInfoCosts:forRingBuffer:forFrame:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_dumpLimiterBatchInfoCosts:forRingBuffer:forFrame:"}
		return err
	}
	g._dumpLimiterBatchInfoCostsForRingBufferForFrame(costs, buffer, frame)
	return nil
}

// CanDumpLimiterBatchInfoCostsForRingBufferForFrame reports whether the receiver responds to the private selector _dumpLimiterBatchInfoCosts:forRingBuffer:forFrame:.
func (g GTAGX2ShaderProfiler) CanDumpLimiterBatchInfoCostsForRingBufferForFrame() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_dumpLimiterBatchInfoCosts:forRingBuffer:forFrame:"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_dumpShaderBinaries
func (g GTAGX2ShaderProfiler) _dumpShaderBinaries() {
	objc.Send[objc.ID](g.ID, objc.Sel("_dumpShaderBinaries"))
}

// DumpShaderBinaries is an exported wrapper for the private method _dumpShaderBinaries.
func (g GTAGX2ShaderProfiler) DumpShaderBinaries() error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_dumpShaderBinaries")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_dumpShaderBinaries"}
		return err
	}
	g._dumpShaderBinaries()
	return nil
}

// CanDumpShaderBinaries reports whether the receiver responds to the private selector _dumpShaderBinaries.
func (g GTAGX2ShaderProfiler) CanDumpShaderBinaries() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_dumpShaderBinaries"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_dumpTraceBufferPacketsInFile:usingTracePackets:withExtractedSamples:withTraceBufferCount:forTargetIndex:
func (g GTAGX2ShaderProfiler) _dumpTraceBufferPacketsInFileUsingTracePacketsWithExtractedSamplesWithTraceBufferCountForTargetIndex(file objectivec.IObject, packets unsafe.Pointer, samples unsafe.Pointer, count uint32, index int) {
	objc.Send[objc.ID](g.ID, objc.Sel("_dumpTraceBufferPacketsInFile:usingTracePackets:withExtractedSamples:withTraceBufferCount:forTargetIndex:"), file, packets, samples, count, index)
}

// DumpTraceBufferPacketsInFileUsingTracePacketsWithExtractedSamplesWithTraceBufferCountForTargetIndex is an exported wrapper for the private method _dumpTraceBufferPacketsInFileUsingTracePacketsWithExtractedSamplesWithTraceBufferCountForTargetIndex.
func (g GTAGX2ShaderProfiler) DumpTraceBufferPacketsInFileUsingTracePacketsWithExtractedSamplesWithTraceBufferCountForTargetIndex(file objectivec.IObject, packets unsafe.Pointer, samples unsafe.Pointer, count uint32, index int) error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_dumpTraceBufferPacketsInFile:usingTracePackets:withExtractedSamples:withTraceBufferCount:forTargetIndex:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_dumpTraceBufferPacketsInFile:usingTracePackets:withExtractedSamples:withTraceBufferCount:forTargetIndex:"}
		return err
	}
	g._dumpTraceBufferPacketsInFileUsingTracePacketsWithExtractedSamplesWithTraceBufferCountForTargetIndex(file, packets, samples, count, index)
	return nil
}

// CanDumpTraceBufferPacketsInFileUsingTracePacketsWithExtractedSamplesWithTraceBufferCountForTargetIndex reports whether the receiver responds to the private selector _dumpTraceBufferPacketsInFile:usingTracePackets:withExtractedSamples:withTraceBufferCount:forTargetIndex:.
func (g GTAGX2ShaderProfiler) CanDumpTraceBufferPacketsInFileUsingTracePacketsWithExtractedSamplesWithTraceBufferCountForTargetIndex() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_dumpTraceBufferPacketsInFile:usingTracePackets:withExtractedSamples:withTraceBufferCount:forTargetIndex:"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_initiateConnectionToLLVMHelper
func (g GTAGX2ShaderProfiler) _initiateConnectionToLLVMHelper() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("_initiateConnectionToLLVMHelper"))
	return rv
}

// InitiateConnectionToLLVMHelper is an exported wrapper for the private method _initiateConnectionToLLVMHelper.
func (g GTAGX2ShaderProfiler) InitiateConnectionToLLVMHelper() (bool, error) {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_initiateConnectionToLLVMHelper")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_initiateConnectionToLLVMHelper"}
		return false, err
	}
	return g._initiateConnectionToLLVMHelper(), nil
}

// CanInitiateConnectionToLLVMHelper reports whether the receiver responds to the private selector _initiateConnectionToLLVMHelper.
func (g GTAGX2ShaderProfiler) CanInitiateConnectionToLLVMHelper() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_initiateConnectionToLLVMHelper"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_latencyAdjustmentEstimate:withLimiter:
func (g GTAGX2ShaderProfiler) _latencyAdjustmentEstimateWithLimiter(estimate float64, limiter float64) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("_latencyAdjustmentEstimate:withLimiter:"), estimate, limiter)
	return rv
}

// LatencyAdjustmentEstimateWithLimiter is an exported wrapper for the private method _latencyAdjustmentEstimateWithLimiter.
func (g GTAGX2ShaderProfiler) LatencyAdjustmentEstimateWithLimiter(estimate float64, limiter float64) (float64, error) {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_latencyAdjustmentEstimate:withLimiter:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_latencyAdjustmentEstimate:withLimiter:"}
		return 0.0, err
	}
	return g._latencyAdjustmentEstimateWithLimiter(estimate, limiter), nil
}

// CanLatencyAdjustmentEstimateWithLimiter reports whether the receiver responds to the private selector _latencyAdjustmentEstimate:withLimiter:.
func (g GTAGX2ShaderProfiler) CanLatencyAdjustmentEstimateWithLimiter() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_latencyAdjustmentEstimate:withLimiter:"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_latencyAdjustmentFactorWithLimiterData:forIndex:withLimiterTypeIndexMap:forLimiterType:
func (g GTAGX2ShaderProfiler) _latencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterType(data []float64, index uint64, map_ unsafe.Pointer, type_ uint32) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("_latencyAdjustmentFactorWithLimiterData:forIndex:withLimiterTypeIndexMap:forLimiterType:"), data, index, map_, type_)
	return rv
}

// LatencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterType is an exported wrapper for the private method _latencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterType.
func (g GTAGX2ShaderProfiler) LatencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterType(data []float64, index uint64, map_ unsafe.Pointer, type_ uint32) (float64, error) {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_latencyAdjustmentFactorWithLimiterData:forIndex:withLimiterTypeIndexMap:forLimiterType:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_latencyAdjustmentFactorWithLimiterData:forIndex:withLimiterTypeIndexMap:forLimiterType:"}
		return 0.0, err
	}
	return g._latencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterType(data, index, map_, type_), nil
}

// CanLatencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterType reports whether the receiver responds to the private selector _latencyAdjustmentFactorWithLimiterData:forIndex:withLimiterTypeIndexMap:forLimiterType:.
func (g GTAGX2ShaderProfiler) CanLatencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterType() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_latencyAdjustmentFactorWithLimiterData:forIndex:withLimiterTypeIndexMap:forLimiterType:"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_latencyAdjustmentFactorWithLimiterData:forIndex:withLimiterTypeIndexMap:forLimiterTypeLoad:forLimiterTypeStore:
func (g GTAGX2ShaderProfiler) _latencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterTypeLoadForLimiterTypeStore(data []float64, index uint64, map_ unsafe.Pointer, load uint32, store uint32) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("_latencyAdjustmentFactorWithLimiterData:forIndex:withLimiterTypeIndexMap:forLimiterTypeLoad:forLimiterTypeStore:"), data, index, map_, load, store)
	return rv
}

// LatencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterTypeLoadForLimiterTypeStore is an exported wrapper for the private method _latencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterTypeLoadForLimiterTypeStore.
func (g GTAGX2ShaderProfiler) LatencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterTypeLoadForLimiterTypeStore(data []float64, index uint64, map_ unsafe.Pointer, load uint32, store uint32) (float64, error) {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_latencyAdjustmentFactorWithLimiterData:forIndex:withLimiterTypeIndexMap:forLimiterTypeLoad:forLimiterTypeStore:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_latencyAdjustmentFactorWithLimiterData:forIndex:withLimiterTypeIndexMap:forLimiterTypeLoad:forLimiterTypeStore:"}
		return 0.0, err
	}
	return g._latencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterTypeLoadForLimiterTypeStore(data, index, map_, load, store), nil
}

// CanLatencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterTypeLoadForLimiterTypeStore reports whether the receiver responds to the private selector _latencyAdjustmentFactorWithLimiterData:forIndex:withLimiterTypeIndexMap:forLimiterTypeLoad:forLimiterTypeStore:.
func (g GTAGX2ShaderProfiler) CanLatencyAdjustmentFactorWithLimiterDataForIndexWithLimiterTypeIndexMapForLimiterTypeLoadForLimiterTypeStore() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_latencyAdjustmentFactorWithLimiterData:forIndex:withLimiterTypeIndexMap:forLimiterTypeLoad:forLimiterTypeStore:"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_latencyAdjustmentFactorWithLimiterDataWithLerp:forIndex:withLimiterTypeIndexMap:forLimiterType:
func (g GTAGX2ShaderProfiler) _latencyAdjustmentFactorWithLimiterDataWithLerpForIndexWithLimiterTypeIndexMapForLimiterType(lerp []float64, index uint64, map_ unsafe.Pointer, type_ uint32) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("_latencyAdjustmentFactorWithLimiterDataWithLerp:forIndex:withLimiterTypeIndexMap:forLimiterType:"), lerp, index, map_, type_)
	return rv
}

// LatencyAdjustmentFactorWithLimiterDataWithLerpForIndexWithLimiterTypeIndexMapForLimiterType is an exported wrapper for the private method _latencyAdjustmentFactorWithLimiterDataWithLerpForIndexWithLimiterTypeIndexMapForLimiterType.
func (g GTAGX2ShaderProfiler) LatencyAdjustmentFactorWithLimiterDataWithLerpForIndexWithLimiterTypeIndexMapForLimiterType(lerp []float64, index uint64, map_ unsafe.Pointer, type_ uint32) (float64, error) {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_latencyAdjustmentFactorWithLimiterDataWithLerp:forIndex:withLimiterTypeIndexMap:forLimiterType:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_latencyAdjustmentFactorWithLimiterDataWithLerp:forIndex:withLimiterTypeIndexMap:forLimiterType:"}
		return 0.0, err
	}
	return g._latencyAdjustmentFactorWithLimiterDataWithLerpForIndexWithLimiterTypeIndexMapForLimiterType(lerp, index, map_, type_), nil
}

// CanLatencyAdjustmentFactorWithLimiterDataWithLerpForIndexWithLimiterTypeIndexMapForLimiterType reports whether the receiver responds to the private selector _latencyAdjustmentFactorWithLimiterDataWithLerp:forIndex:withLimiterTypeIndexMap:forLimiterType:.
func (g GTAGX2ShaderProfiler) CanLatencyAdjustmentFactorWithLimiterDataWithLerpForIndexWithLimiterTypeIndexMapForLimiterType() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_latencyAdjustmentFactorWithLimiterDataWithLerp:forIndex:withLimiterTypeIndexMap:forLimiterType:"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_latencyAdjustmentWithLimiterData:forIndex:withLimiterTypeIndexMap:
func (g GTAGX2ShaderProfiler) _latencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap(data []float64, index uint64, map_ unsafe.Pointer) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("_latencyAdjustmentWithLimiterData:forIndex:withLimiterTypeIndexMap:"), data, index, map_)
	return rv
}

// LatencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap is an exported wrapper for the private method _latencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap.
func (g GTAGX2ShaderProfiler) LatencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap(data []float64, index uint64, map_ unsafe.Pointer) (float64, error) {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_latencyAdjustmentWithLimiterData:forIndex:withLimiterTypeIndexMap:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_latencyAdjustmentWithLimiterData:forIndex:withLimiterTypeIndexMap:"}
		return 0.0, err
	}
	return g._latencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap(data, index, map_), nil
}

// CanLatencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap reports whether the receiver responds to the private selector _latencyAdjustmentWithLimiterData:forIndex:withLimiterTypeIndexMap:.
func (g GTAGX2ShaderProfiler) CanLatencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_latencyAdjustmentWithLimiterData:forIndex:withLimiterTypeIndexMap:"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_mergeLegacyAndNewShadersInShaderInfoDictionary:
func (g GTAGX2ShaderProfiler) _mergeLegacyAndNewShadersInShaderInfoDictionary(dictionary objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_mergeLegacyAndNewShadersInShaderInfoDictionary:"), dictionary)
	return objectivec.Object{ID: rv}
}

// MergeLegacyAndNewShadersInShaderInfoDictionary is an exported wrapper for the private method _mergeLegacyAndNewShadersInShaderInfoDictionary.
func (g GTAGX2ShaderProfiler) MergeLegacyAndNewShadersInShaderInfoDictionary(dictionary objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_mergeLegacyAndNewShadersInShaderInfoDictionary:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_mergeLegacyAndNewShadersInShaderInfoDictionary:"}
		return nil, err
	}
	return g._mergeLegacyAndNewShadersInShaderInfoDictionary(dictionary), nil
}

// CanMergeLegacyAndNewShadersInShaderInfoDictionary reports whether the receiver responds to the private selector _mergeLegacyAndNewShadersInShaderInfoDictionary:.
func (g GTAGX2ShaderProfiler) CanMergeLegacyAndNewShadersInShaderInfoDictionary() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_mergeLegacyAndNewShadersInShaderInfoDictionary:"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_processTracePackets:forRenderIndex:andGenerateSampleList:forTargetIndex:forLimiterIndex:
func (g GTAGX2ShaderProfiler) _processTracePacketsForRenderIndexAndGenerateSampleListForTargetIndexForLimiterIndex(packets unsafe.Pointer, index uint32, list unsafe.Pointer, index2 int, index3 uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("_processTracePackets:forRenderIndex:andGenerateSampleList:forTargetIndex:forLimiterIndex:"), packets, index, list, index2, index3)
}

// ProcessTracePacketsForRenderIndexAndGenerateSampleListForTargetIndexForLimiterIndex is an exported wrapper for the private method _processTracePacketsForRenderIndexAndGenerateSampleListForTargetIndexForLimiterIndex.
func (g GTAGX2ShaderProfiler) ProcessTracePacketsForRenderIndexAndGenerateSampleListForTargetIndexForLimiterIndex(packets unsafe.Pointer, index uint32, list unsafe.Pointer, index2 int, index3 uint32) error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_processTracePackets:forRenderIndex:andGenerateSampleList:forTargetIndex:forLimiterIndex:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_processTracePackets:forRenderIndex:andGenerateSampleList:forTargetIndex:forLimiterIndex:"}
		return err
	}
	g._processTracePacketsForRenderIndexAndGenerateSampleListForTargetIndexForLimiterIndex(packets, index, list, index2, index3)
	return nil
}

// CanProcessTracePacketsForRenderIndexAndGenerateSampleListForTargetIndexForLimiterIndex reports whether the receiver responds to the private selector _processTracePackets:forRenderIndex:andGenerateSampleList:forTargetIndex:forLimiterIndex:.
func (g GTAGX2ShaderProfiler) CanProcessTracePacketsForRenderIndexAndGenerateSampleListForTargetIndexForLimiterIndex() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_processTracePackets:forRenderIndex:andGenerateSampleList:forTargetIndex:forLimiterIndex:"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_setupShaderBinaryInfo:withBinaryKey:andNumDraws:
func (g GTAGX2ShaderProfiler) _setupShaderBinaryInfoWithBinaryKeyAndNumDraws(info objectivec.IObject, key unsafe.Pointer, draws uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("_setupShaderBinaryInfo:withBinaryKey:andNumDraws:"), info, key, draws)
}

// SetupShaderBinaryInfoWithBinaryKeyAndNumDraws is an exported wrapper for the private method _setupShaderBinaryInfoWithBinaryKeyAndNumDraws.
func (g GTAGX2ShaderProfiler) SetupShaderBinaryInfoWithBinaryKeyAndNumDraws(info objectivec.IObject, key unsafe.Pointer, draws uint32) error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_setupShaderBinaryInfo:withBinaryKey:andNumDraws:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setupShaderBinaryInfo:withBinaryKey:andNumDraws:"}
		return err
	}
	g._setupShaderBinaryInfoWithBinaryKeyAndNumDraws(info, key, draws)
	return nil
}

// CanSetupShaderBinaryInfoWithBinaryKeyAndNumDraws reports whether the receiver responds to the private selector _setupShaderBinaryInfo:withBinaryKey:andNumDraws:.
func (g GTAGX2ShaderProfiler) CanSetupShaderBinaryInfoWithBinaryKeyAndNumDraws() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_setupShaderBinaryInfo:withBinaryKey:andNumDraws:"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/_waitLatencyAdjustmentWithLimiterData:forIndex:withLimiterTypeIndexMap:
func (g GTAGX2ShaderProfiler) _waitLatencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap(data []float64, index uint64, map_ unsafe.Pointer) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("_waitLatencyAdjustmentWithLimiterData:forIndex:withLimiterTypeIndexMap:"), data, index, map_)
	return rv
}

// WaitLatencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap is an exported wrapper for the private method _waitLatencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap.
func (g GTAGX2ShaderProfiler) WaitLatencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap(data []float64, index uint64, map_ unsafe.Pointer) (float64, error) {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_waitLatencyAdjustmentWithLimiterData:forIndex:withLimiterTypeIndexMap:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_waitLatencyAdjustmentWithLimiterData:forIndex:withLimiterTypeIndexMap:"}
		return 0.0, err
	}
	return g._waitLatencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap(data, index, map_), nil
}

// CanWaitLatencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap reports whether the receiver responds to the private selector _waitLatencyAdjustmentWithLimiterData:forIndex:withLimiterTypeIndexMap:.
func (g GTAGX2ShaderProfiler) CanWaitLatencyAdjustmentWithLimiterDataForIndexWithLimiterTypeIndexMap() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_waitLatencyAdjustmentWithLimiterData:forIndex:withLimiterTypeIndexMap:"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/adjustHWBiasAndFinalizeResult:blitKickIndices:blitTimes:result:
func (g GTAGX2ShaderProfiler) AdjustHWBiasAndFinalizeResultBlitKickIndicesBlitTimesResult(result objectivec.IObject, indices objectivec.IObject, times objectivec.IObject, result2 objectivec.IObject) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("adjustHWBiasAndFinalizeResult:blitKickIndices:blitTimes:result:"), result, indices, times, result2)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/dumpShaderBinarySamples:
func (g GTAGX2ShaderProfiler) DumpShaderBinarySamples(samples unsafe.Pointer) {
	objc.Send[objc.ID](g.ID, objc.Sel("dumpShaderBinarySamples:"), samples)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/evaluateStreamingSamples:withUSCSampleNum:withProgramAddressLUT:targetIndex:forRingBufferIndex:withMinEncoderIndex:withMaxEncoderIndex:withEncoderIdToEncoderIndexMap:withProfilingData:
func (g GTAGX2ShaderProfiler) EvaluateStreamingSamplesWithUSCSampleNumWithProgramAddressLUTTargetIndexForRingBufferIndexWithMinEncoderIndexWithMaxEncoderIndexWithEncoderIdToEncoderIndexMapWithProfilingData(samples unsafe.Pointer, num uint32, lut unsafe.Pointer, index int, index2 uint32, index3 uint32, index4 uint32, map_ unsafe.Pointer, data objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("evaluateStreamingSamples:withUSCSampleNum:withProgramAddressLUT:targetIndex:forRingBufferIndex:withMinEncoderIndex:withMaxEncoderIndex:withEncoderIdToEncoderIndexMap:withProfilingData:"), samples, num, lut, index, index2, index3, index4, map_, data)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/handleHarvestedBinaryInfo:
func (g GTAGX2ShaderProfiler) HandleHarvestedBinaryInfo(info objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("handleHarvestedBinaryInfo:"), info)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/setRingBufferCount:
func (g GTAGX2ShaderProfiler) SetRingBufferCount(count uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("setRingBufferCount:"), count)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/initWithStreamData:forTargetIndex:
func (g GTAGX2ShaderProfiler) InitWithStreamDataForTargetIndex(data objectivec.IObject, index int) GTAGX2ShaderProfiler {
	rv := objc.Send[GTAGX2ShaderProfiler](g.ID, objc.Sel("initWithStreamData:forTargetIndex:"), data, index)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/averagePerDrawKickDurations
func (g GTAGX2ShaderProfiler) AveragePerDrawKickDurations() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("averagePerDrawKickDurations"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/effectiveKickTimes
func (g GTAGX2ShaderProfiler) EffectiveKickTimes() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("effectiveKickTimes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/isaPrinter
func (g GTAGX2ShaderProfiler) IsaPrinter() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("isaPrinter"))
	return objectivec.Object{ID: rv}
}
func (g GTAGX2ShaderProfiler) SetIsaPrinter(value objectivec.IObject) {
	objc.Send[struct{}](g.ID, objc.Sel("setIsaPrinter:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/loadActionTimes
func (g GTAGX2ShaderProfiler) LoadActionTimes() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("loadActionTimes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/perRingPerFrameLimiterData
func (g GTAGX2ShaderProfiler) PerRingPerFrameLimiterData() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("perRingPerFrameLimiterData"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/storeActionTimes
func (g GTAGX2ShaderProfiler) StoreActionTimes() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("storeActionTimes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfiler/timingInfo
func (g GTAGX2ShaderProfiler) TimingInfo() IGTShaderProfilerTimingInfo {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("timingInfo"))
	return GTShaderProfilerTimingInfoFromID(objc.ID(rv))
}

// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTAGX2StreamDataShaderProfilerProcessor] class.
var (
	_GTAGX2StreamDataShaderProfilerProcessorClass     GTAGX2StreamDataShaderProfilerProcessorClass
	_GTAGX2StreamDataShaderProfilerProcessorClassOnce sync.Once
)

func getGTAGX2StreamDataShaderProfilerProcessorClass() GTAGX2StreamDataShaderProfilerProcessorClass {
	_GTAGX2StreamDataShaderProfilerProcessorClassOnce.Do(func() {
		_GTAGX2StreamDataShaderProfilerProcessorClass = GTAGX2StreamDataShaderProfilerProcessorClass{class: objc.GetClass("GTAGX2StreamDataShaderProfilerProcessor")}
	})
	return _GTAGX2StreamDataShaderProfilerProcessorClass
}

// GetGTAGX2StreamDataShaderProfilerProcessorClass returns the class object for GTAGX2StreamDataShaderProfilerProcessor.
func GetGTAGX2StreamDataShaderProfilerProcessorClass() GTAGX2StreamDataShaderProfilerProcessorClass {
	return getGTAGX2StreamDataShaderProfilerProcessorClass()
}

type GTAGX2StreamDataShaderProfilerProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTAGX2StreamDataShaderProfilerProcessorClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTAGX2StreamDataShaderProfilerProcessorClass) Alloc() GTAGX2StreamDataShaderProfilerProcessor {
	rv := objc.Send[GTAGX2StreamDataShaderProfilerProcessor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTAGX2StreamDataShaderProfilerProcessor._createPerCounterCommandDataWithPerCommandData]
//   - [GTAGX2StreamDataShaderProfilerProcessor._effectiveBatchDrawKickTimes]
//   - [GTAGX2StreamDataShaderProfilerProcessor._handleStreamingBatchResult]
//   - [GTAGX2StreamDataShaderProfilerProcessor._postProcessData]
//   - [GTAGX2StreamDataShaderProfilerProcessor._preProcessStreamingUSCSampleData]
//   - [GTAGX2StreamDataShaderProfilerProcessor._preProcessStreamingUSCSampleDataWithAddressDataSampleDataFrameIndexRingBufferIdxTargetIndex]
//   - [GTAGX2StreamDataShaderProfilerProcessor._preProcessStreamingUSCSampleDataWithAddressMappingsSampleDataFrameIndexRingBufferIdxTargetIndex]
//   - [GTAGX2StreamDataShaderProfilerProcessor._processDerivedEncoderCounterData]
//   - [GTAGX2StreamDataShaderProfilerProcessor._processFrameTimeData]
//   - [GTAGX2StreamDataShaderProfilerProcessor._processHarvestedBinaryData]
//   - [GTAGX2StreamDataShaderProfilerProcessor._saveAddressListSizeFilename]
//   - [GTAGX2StreamDataShaderProfilerProcessor.AnalyzeBinaryGpuGeneration]
//   - [GTAGX2StreamDataShaderProfilerProcessor.AnalyzeBinaryTypeNameDylibDataGpuGeneration]
//   - [GTAGX2StreamDataShaderProfilerProcessor.Delegate]
//   - [GTAGX2StreamDataShaderProfilerProcessor.SetDelegate]
//   - [GTAGX2StreamDataShaderProfilerProcessor.IsaPrinter]
//   - [GTAGX2StreamDataShaderProfilerProcessor.SetIsaPrinter]
//   - [GTAGX2StreamDataShaderProfilerProcessor.Process]
//   - [GTAGX2StreamDataShaderProfilerProcessor.ProcessBatchIDFilteringData]
//   - [GTAGX2StreamDataShaderProfilerProcessor.ProcessBatchIDFilteringDataWithData]
//   - [GTAGX2StreamDataShaderProfilerProcessor.ProcessBatchIdData]
//   - [GTAGX2StreamDataShaderProfilerProcessor.ProcessShaderProfilerStreamedResult]
//   - [GTAGX2StreamDataShaderProfilerProcessor.ProcessStreamData]
//   - [GTAGX2StreamDataShaderProfilerProcessor.SetupForBatchIDFilteringCounters]
//   - [GTAGX2StreamDataShaderProfilerProcessor.ShaderProfilerResult]
//   - [GTAGX2StreamDataShaderProfilerProcessor.WaitUntilBatchIDCounterFinished]
//   - [GTAGX2StreamDataShaderProfilerProcessor.WaitUntilStreamDataFinished]
//   - [GTAGX2StreamDataShaderProfilerProcessor.InitWithStreamData]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor
type GTAGX2StreamDataShaderProfilerProcessor struct {
	objectivec.Object
}

// GTAGX2StreamDataShaderProfilerProcessorFromID constructs a [GTAGX2StreamDataShaderProfilerProcessor] from an objc.ID.
func GTAGX2StreamDataShaderProfilerProcessorFromID(id objc.ID) GTAGX2StreamDataShaderProfilerProcessor {
	return GTAGX2StreamDataShaderProfilerProcessor{objectivec.Object{ID: id}}
}

// Ensure GTAGX2StreamDataShaderProfilerProcessor implements IGTAGX2StreamDataShaderProfilerProcessor.
var _ IGTAGX2StreamDataShaderProfilerProcessor = GTAGX2StreamDataShaderProfilerProcessor{}

// An interface definition for the [GTAGX2StreamDataShaderProfilerProcessor] class.
//
// # Methods
//
//   - [IGTAGX2StreamDataShaderProfilerProcessor._createPerCounterCommandDataWithPerCommandData]
//   - [IGTAGX2StreamDataShaderProfilerProcessor._effectiveBatchDrawKickTimes]
//   - [IGTAGX2StreamDataShaderProfilerProcessor._handleStreamingBatchResult]
//   - [IGTAGX2StreamDataShaderProfilerProcessor._postProcessData]
//   - [IGTAGX2StreamDataShaderProfilerProcessor._preProcessStreamingUSCSampleData]
//   - [IGTAGX2StreamDataShaderProfilerProcessor._preProcessStreamingUSCSampleDataWithAddressDataSampleDataFrameIndexRingBufferIdxTargetIndex]
//   - [IGTAGX2StreamDataShaderProfilerProcessor._preProcessStreamingUSCSampleDataWithAddressMappingsSampleDataFrameIndexRingBufferIdxTargetIndex]
//   - [IGTAGX2StreamDataShaderProfilerProcessor._processDerivedEncoderCounterData]
//   - [IGTAGX2StreamDataShaderProfilerProcessor._processFrameTimeData]
//   - [IGTAGX2StreamDataShaderProfilerProcessor._processHarvestedBinaryData]
//   - [IGTAGX2StreamDataShaderProfilerProcessor._saveAddressListSizeFilename]
//   - [IGTAGX2StreamDataShaderProfilerProcessor.AnalyzeBinaryGpuGeneration]
//   - [IGTAGX2StreamDataShaderProfilerProcessor.AnalyzeBinaryTypeNameDylibDataGpuGeneration]
//   - [IGTAGX2StreamDataShaderProfilerProcessor.Delegate]
//   - [IGTAGX2StreamDataShaderProfilerProcessor.SetDelegate]
//   - [IGTAGX2StreamDataShaderProfilerProcessor.IsaPrinter]
//   - [IGTAGX2StreamDataShaderProfilerProcessor.SetIsaPrinter]
//   - [IGTAGX2StreamDataShaderProfilerProcessor.Process]
//   - [IGTAGX2StreamDataShaderProfilerProcessor.ProcessBatchIDFilteringData]
//   - [IGTAGX2StreamDataShaderProfilerProcessor.ProcessBatchIDFilteringDataWithData]
//   - [IGTAGX2StreamDataShaderProfilerProcessor.ProcessBatchIdData]
//   - [IGTAGX2StreamDataShaderProfilerProcessor.ProcessShaderProfilerStreamedResult]
//   - [IGTAGX2StreamDataShaderProfilerProcessor.ProcessStreamData]
//   - [IGTAGX2StreamDataShaderProfilerProcessor.SetupForBatchIDFilteringCounters]
//   - [IGTAGX2StreamDataShaderProfilerProcessor.ShaderProfilerResult]
//   - [IGTAGX2StreamDataShaderProfilerProcessor.WaitUntilBatchIDCounterFinished]
//   - [IGTAGX2StreamDataShaderProfilerProcessor.WaitUntilStreamDataFinished]
//   - [IGTAGX2StreamDataShaderProfilerProcessor.InitWithStreamData]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor
type IGTAGX2StreamDataShaderProfilerProcessor interface {
	objectivec.IObject

	// Topic: Methods

	_createPerCounterCommandDataWithPerCommandData(data objectivec.IObject, data2 objectivec.IObject) objectivec.IObject
	_effectiveBatchDrawKickTimes(times objectivec.IObject) objectivec.IObject
	_handleStreamingBatchResult(result objectivec.IObject)
	_postProcessData()
	_preProcessStreamingUSCSampleData(data objectivec.IObject)
	_preProcessStreamingUSCSampleDataWithAddressDataSampleDataFrameIndexRingBufferIdxTargetIndex(data objectivec.IObject, data2 objectivec.IObject, data3 objectivec.IObject, index uint32, idx uint32, index2 int)
	_preProcessStreamingUSCSampleDataWithAddressMappingsSampleDataFrameIndexRingBufferIdxTargetIndex(data objectivec.IObject, mappings objectivec.IObject, data2 objectivec.IObject, index uint32, idx uint32, index2 int)
	_processDerivedEncoderCounterData(data objectivec.IObject)
	_processFrameTimeData(data objectivec.IObject)
	_processHarvestedBinaryData(data objectivec.IObject)
	_saveAddressListSizeFilename(list unsafe.Pointer, size uint32, filename string)
	AnalyzeBinaryGpuGeneration(binary objectivec.IObject, generation uint32) objectivec.IObject
	AnalyzeBinaryTypeNameDylibDataGpuGeneration(binary objectivec.IObject, name objectivec.IObject, dylib bool, data objectivec.IObject, generation uint32) objectivec.IObject
	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)
	IsaPrinter() objectivec.IObject
	SetIsaPrinter(value objectivec.IObject)
	Process(process objectivec.IObject)
	ProcessBatchIDFilteringData()
	ProcessBatchIDFilteringDataWithData(data objectivec.IObject)
	ProcessBatchIdData(data objectivec.IObject) objectivec.IObject
	ProcessShaderProfilerStreamedResult(result objectivec.IObject) objectivec.IObject
	ProcessStreamData()
	SetupForBatchIDFilteringCounters(counters objectivec.IObject) bool
	ShaderProfilerResult() objectivec.IObject
	WaitUntilBatchIDCounterFinished()
	WaitUntilStreamDataFinished()
	InitWithStreamData(data objectivec.IObject) GTAGX2StreamDataShaderProfilerProcessor
}

// Init initializes the instance.
func (g GTAGX2StreamDataShaderProfilerProcessor) Init() GTAGX2StreamDataShaderProfilerProcessor {
	rv := objc.Send[GTAGX2StreamDataShaderProfilerProcessor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTAGX2StreamDataShaderProfilerProcessor) Autorelease() GTAGX2StreamDataShaderProfilerProcessor {
	rv := objc.Send[GTAGX2StreamDataShaderProfilerProcessor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTAGX2StreamDataShaderProfilerProcessor creates a new GTAGX2StreamDataShaderProfilerProcessor instance.
func NewGTAGX2StreamDataShaderProfilerProcessor() GTAGX2StreamDataShaderProfilerProcessor {
	class := getGTAGX2StreamDataShaderProfilerProcessorClass()
	rv := objc.Send[GTAGX2StreamDataShaderProfilerProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/initWithStreamData:
func NewGTAGX2StreamDataShaderProfilerProcessorWithStreamData(data objectivec.IObject) GTAGX2StreamDataShaderProfilerProcessor {
	instance := getGTAGX2StreamDataShaderProfilerProcessorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStreamData:"), data)
	return GTAGX2StreamDataShaderProfilerProcessorFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/_createPerCounterCommandData:withPerCommandData:
func (g GTAGX2StreamDataShaderProfilerProcessor) _createPerCounterCommandDataWithPerCommandData(data objectivec.IObject, data2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_createPerCounterCommandData:withPerCommandData:"), data, data2)
	return objectivec.Object{ID: rv}
}

// CreatePerCounterCommandDataWithPerCommandData is an exported wrapper for the private method _createPerCounterCommandDataWithPerCommandData.
func (g GTAGX2StreamDataShaderProfilerProcessor) CreatePerCounterCommandDataWithPerCommandData(data objectivec.IObject, data2 objectivec.IObject) objectivec.IObject {
	return g._createPerCounterCommandDataWithPerCommandData(data, data2)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/_effectiveBatchDrawKickTimes:
func (g GTAGX2StreamDataShaderProfilerProcessor) _effectiveBatchDrawKickTimes(times objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_effectiveBatchDrawKickTimes:"), times)
	return objectivec.Object{ID: rv}
}

// EffectiveBatchDrawKickTimes is an exported wrapper for the private method _effectiveBatchDrawKickTimes.
func (g GTAGX2StreamDataShaderProfilerProcessor) EffectiveBatchDrawKickTimes(times objectivec.IObject) objectivec.IObject {
	return g._effectiveBatchDrawKickTimes(times)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/_handleStreamingBatchResult:
func (g GTAGX2StreamDataShaderProfilerProcessor) _handleStreamingBatchResult(result objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("_handleStreamingBatchResult:"), result)
}

// HandleStreamingBatchResult is an exported wrapper for the private method _handleStreamingBatchResult.
func (g GTAGX2StreamDataShaderProfilerProcessor) HandleStreamingBatchResult(result objectivec.IObject) {
	g._handleStreamingBatchResult(result)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/_postProcessData
func (g GTAGX2StreamDataShaderProfilerProcessor) _postProcessData() {
	objc.Send[objc.ID](g.ID, objc.Sel("_postProcessData"))
}

// PostProcessData is an exported wrapper for the private method _postProcessData.
func (g GTAGX2StreamDataShaderProfilerProcessor) PostProcessData() {
	g._postProcessData()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/_preProcessStreamingUSCSampleData:
func (g GTAGX2StreamDataShaderProfilerProcessor) _preProcessStreamingUSCSampleData(data objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("_preProcessStreamingUSCSampleData:"), data)
}

// PreProcessStreamingUSCSampleData is an exported wrapper for the private method _preProcessStreamingUSCSampleData.
func (g GTAGX2StreamDataShaderProfilerProcessor) PreProcessStreamingUSCSampleData(data objectivec.IObject) {
	g._preProcessStreamingUSCSampleData(data)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/_preProcessStreamingUSCSampleData:withAddressData:sampleData:frameIndex:ringBufferIdx:targetIndex:
func (g GTAGX2StreamDataShaderProfilerProcessor) _preProcessStreamingUSCSampleDataWithAddressDataSampleDataFrameIndexRingBufferIdxTargetIndex(data objectivec.IObject, data2 objectivec.IObject, data3 objectivec.IObject, index uint32, idx uint32, index2 int) {
	objc.Send[objc.ID](g.ID, objc.Sel("_preProcessStreamingUSCSampleData:withAddressData:sampleData:frameIndex:ringBufferIdx:targetIndex:"), data, data2, data3, index, idx, index2)
}

// PreProcessStreamingUSCSampleDataWithAddressDataSampleDataFrameIndexRingBufferIdxTargetIndex is an exported wrapper for the private method _preProcessStreamingUSCSampleDataWithAddressDataSampleDataFrameIndexRingBufferIdxTargetIndex.
func (g GTAGX2StreamDataShaderProfilerProcessor) PreProcessStreamingUSCSampleDataWithAddressDataSampleDataFrameIndexRingBufferIdxTargetIndex(data objectivec.IObject, data2 objectivec.IObject, data3 objectivec.IObject, index uint32, idx uint32, index2 int) {
	g._preProcessStreamingUSCSampleDataWithAddressDataSampleDataFrameIndexRingBufferIdxTargetIndex(data, data2, data3, index, idx, index2)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/_preProcessStreamingUSCSampleData:withAddressMappings:sampleData:frameIndex:ringBufferIdx:targetIndex:
func (g GTAGX2StreamDataShaderProfilerProcessor) _preProcessStreamingUSCSampleDataWithAddressMappingsSampleDataFrameIndexRingBufferIdxTargetIndex(data objectivec.IObject, mappings objectivec.IObject, data2 objectivec.IObject, index uint32, idx uint32, index2 int) {
	objc.Send[objc.ID](g.ID, objc.Sel("_preProcessStreamingUSCSampleData:withAddressMappings:sampleData:frameIndex:ringBufferIdx:targetIndex:"), data, mappings, data2, index, idx, index2)
}

// PreProcessStreamingUSCSampleDataWithAddressMappingsSampleDataFrameIndexRingBufferIdxTargetIndex is an exported wrapper for the private method _preProcessStreamingUSCSampleDataWithAddressMappingsSampleDataFrameIndexRingBufferIdxTargetIndex.
func (g GTAGX2StreamDataShaderProfilerProcessor) PreProcessStreamingUSCSampleDataWithAddressMappingsSampleDataFrameIndexRingBufferIdxTargetIndex(data objectivec.IObject, mappings objectivec.IObject, data2 objectivec.IObject, index uint32, idx uint32, index2 int) {
	g._preProcessStreamingUSCSampleDataWithAddressMappingsSampleDataFrameIndexRingBufferIdxTargetIndex(data, mappings, data2, index, idx, index2)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/_processDerivedEncoderCounterData:
func (g GTAGX2StreamDataShaderProfilerProcessor) _processDerivedEncoderCounterData(data objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("_processDerivedEncoderCounterData:"), data)
}

// ProcessDerivedEncoderCounterData is an exported wrapper for the private method _processDerivedEncoderCounterData.
func (g GTAGX2StreamDataShaderProfilerProcessor) ProcessDerivedEncoderCounterData(data objectivec.IObject) {
	g._processDerivedEncoderCounterData(data)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/_processFrameTimeData:
func (g GTAGX2StreamDataShaderProfilerProcessor) _processFrameTimeData(data objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("_processFrameTimeData:"), data)
}

// ProcessFrameTimeData is an exported wrapper for the private method _processFrameTimeData.
func (g GTAGX2StreamDataShaderProfilerProcessor) ProcessFrameTimeData(data objectivec.IObject) {
	g._processFrameTimeData(data)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/_processHarvestedBinaryData:
func (g GTAGX2StreamDataShaderProfilerProcessor) _processHarvestedBinaryData(data objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("_processHarvestedBinaryData:"), data)
}

// ProcessHarvestedBinaryData is an exported wrapper for the private method _processHarvestedBinaryData.
func (g GTAGX2StreamDataShaderProfilerProcessor) ProcessHarvestedBinaryData(data objectivec.IObject) {
	g._processHarvestedBinaryData(data)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/_saveAddressList:size:filename:
func (g GTAGX2StreamDataShaderProfilerProcessor) _saveAddressListSizeFilename(list unsafe.Pointer, size uint32, filename string) {
	objc.Send[objc.ID](g.ID, objc.Sel("_saveAddressList:size:filename:"), list, size, unsafe.Pointer(unsafe.StringData(filename+"\x00")))
}

// SaveAddressListSizeFilename is an exported wrapper for the private method _saveAddressListSizeFilename.
func (g GTAGX2StreamDataShaderProfilerProcessor) SaveAddressListSizeFilename(list unsafe.Pointer, size uint32, filename string) {
	g._saveAddressListSizeFilename(list, size, filename)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/analyzeBinary:gpuGeneration:
func (g GTAGX2StreamDataShaderProfilerProcessor) AnalyzeBinaryGpuGeneration(binary objectivec.IObject, generation uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("analyzeBinary:gpuGeneration:"), binary, generation)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/analyzeBinary:typeName:dylib:data:gpuGeneration:
func (g GTAGX2StreamDataShaderProfilerProcessor) AnalyzeBinaryTypeNameDylibDataGpuGeneration(binary objectivec.IObject, name objectivec.IObject, dylib bool, data objectivec.IObject, generation uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("analyzeBinary:typeName:dylib:data:gpuGeneration:"), binary, name, dylib, data, generation)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/process:
func (g GTAGX2StreamDataShaderProfilerProcessor) Process(process objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("process:"), process)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/processBatchIDFilteringData
func (g GTAGX2StreamDataShaderProfilerProcessor) ProcessBatchIDFilteringData() {
	objc.Send[objc.ID](g.ID, objc.Sel("processBatchIDFilteringData"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/processBatchIDFilteringData:
func (g GTAGX2StreamDataShaderProfilerProcessor) ProcessBatchIDFilteringDataWithData(data objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("processBatchIDFilteringData:"), data)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/processBatchIdData:
func (g GTAGX2StreamDataShaderProfilerProcessor) ProcessBatchIdData(data objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("processBatchIdData:"), data)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/processShaderProfilerStreamedResult:
func (g GTAGX2StreamDataShaderProfilerProcessor) ProcessShaderProfilerStreamedResult(result objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("processShaderProfilerStreamedResult:"), result)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/processStreamData
func (g GTAGX2StreamDataShaderProfilerProcessor) ProcessStreamData() {
	objc.Send[objc.ID](g.ID, objc.Sel("processStreamData"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/setupForBatchIDFilteringCounters:
func (g GTAGX2StreamDataShaderProfilerProcessor) SetupForBatchIDFilteringCounters(counters objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("setupForBatchIDFilteringCounters:"), counters)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/waitUntilBatchIDCounterFinished
func (g GTAGX2StreamDataShaderProfilerProcessor) WaitUntilBatchIDCounterFinished() {
	objc.Send[objc.ID](g.ID, objc.Sel("waitUntilBatchIDCounterFinished"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/waitUntilStreamDataFinished
func (g GTAGX2StreamDataShaderProfilerProcessor) WaitUntilStreamDataFinished() {
	objc.Send[objc.ID](g.ID, objc.Sel("waitUntilStreamDataFinished"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/initWithStreamData:
func (g GTAGX2StreamDataShaderProfilerProcessor) InitWithStreamData(data objectivec.IObject) GTAGX2StreamDataShaderProfilerProcessor {
	rv := objc.Send[GTAGX2StreamDataShaderProfilerProcessor](g.ID, objc.Sel("initWithStreamData:"), data)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/delegate
func (g GTAGX2StreamDataShaderProfilerProcessor) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (g GTAGX2StreamDataShaderProfilerProcessor) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](g.ID, objc.Sel("setDelegate:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/isaPrinter
func (g GTAGX2StreamDataShaderProfilerProcessor) IsaPrinter() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("isaPrinter"))
	return objectivec.Object{ID: rv}
}
func (g GTAGX2StreamDataShaderProfilerProcessor) SetIsaPrinter(value objectivec.IObject) {
	objc.Send[struct{}](g.ID, objc.Sel("setIsaPrinter:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2StreamDataShaderProfilerProcessor/shaderProfilerResult
func (g GTAGX2StreamDataShaderProfilerProcessor) ShaderProfilerResult() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("shaderProfilerResult"))
	return objectivec.Object{ID: rv}
}

// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTAGX2StreamDataTimelineProcessor] class.
var (
	_GTAGX2StreamDataTimelineProcessorClass     GTAGX2StreamDataTimelineProcessorClass
	_GTAGX2StreamDataTimelineProcessorClassOnce sync.Once
)

func getGTAGX2StreamDataTimelineProcessorClass() GTAGX2StreamDataTimelineProcessorClass {
	_GTAGX2StreamDataTimelineProcessorClassOnce.Do(func() {
		_GTAGX2StreamDataTimelineProcessorClass = GTAGX2StreamDataTimelineProcessorClass{class: objc.GetClass("GTAGX2StreamDataTimelineProcessor")}
	})
	return _GTAGX2StreamDataTimelineProcessorClass
}

// GetGTAGX2StreamDataTimelineProcessorClass returns the class object for GTAGX2StreamDataTimelineProcessor.
func GetGTAGX2StreamDataTimelineProcessorClass() GTAGX2StreamDataTimelineProcessorClass {
	return getGTAGX2StreamDataTimelineProcessorClass()
}

type GTAGX2StreamDataTimelineProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTAGX2StreamDataTimelineProcessorClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTAGX2StreamDataTimelineProcessorClass) Alloc() GTAGX2StreamDataTimelineProcessor {
	rv := objc.SendIfResponds[GTAGX2StreamDataTimelineProcessor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTAGX2StreamDataTimelineProcessor._addDerivedCounterDataToTimelineInfoForWorkloadInfo]
//   - [GTAGX2StreamDataTimelineProcessor._calculatePerSampleActiveShadersForWorkloadInfo]
//   - [GTAGX2StreamDataTimelineProcessor._calculatePerSampleAggregatedActiveShadersForWorkloadInfo]
//   - [GTAGX2StreamDataTimelineProcessor._updateShaderTimelineInfoWithShaderTimelineDataExForGPUTimelineInfoWithEncoderGlobalTraceIdToStateMirrorIdMappingForRingBuffer]
//   - [GTAGX2StreamDataTimelineProcessor.Process]
//   - [GTAGX2StreamDataTimelineProcessor.ProcessStreamData]
//   - [GTAGX2StreamDataTimelineProcessor.ProcessTimelineStreamedResult]
//   - [GTAGX2StreamDataTimelineProcessor.SaveAddressListSizeFilename]
//   - [GTAGX2StreamDataTimelineProcessor.SaveAddressMappingsFilename]
//   - [GTAGX2StreamDataTimelineProcessor.SaveSampleDataFromGPURawCountersSizeFilenameWithTimeStamps]
//   - [GTAGX2StreamDataTimelineProcessor.TimelineInfo]
//   - [GTAGX2StreamDataTimelineProcessor.WaitUntilFinished]
//   - [GTAGX2StreamDataTimelineProcessor.InitWithStreamData]
type GTAGX2StreamDataTimelineProcessor struct {
	objectivec.Object
}

// GTAGX2StreamDataTimelineProcessorFromID constructs a [GTAGX2StreamDataTimelineProcessor] from an objc.ID.
func GTAGX2StreamDataTimelineProcessorFromID(id objc.ID) GTAGX2StreamDataTimelineProcessor {
	return GTAGX2StreamDataTimelineProcessor{objectivec.Object{ID: id}}
}

// Ensure GTAGX2StreamDataTimelineProcessor implements IGTAGX2StreamDataTimelineProcessor.
var _ IGTAGX2StreamDataTimelineProcessor = GTAGX2StreamDataTimelineProcessor{}

// An interface definition for the [GTAGX2StreamDataTimelineProcessor] class.
//
// # Methods
//
//   - [IGTAGX2StreamDataTimelineProcessor._addDerivedCounterDataToTimelineInfoForWorkloadInfo]
//   - [IGTAGX2StreamDataTimelineProcessor._calculatePerSampleActiveShadersForWorkloadInfo]
//   - [IGTAGX2StreamDataTimelineProcessor._calculatePerSampleAggregatedActiveShadersForWorkloadInfo]
//   - [IGTAGX2StreamDataTimelineProcessor._updateShaderTimelineInfoWithShaderTimelineDataExForGPUTimelineInfoWithEncoderGlobalTraceIdToStateMirrorIdMappingForRingBuffer]
//   - [IGTAGX2StreamDataTimelineProcessor.Process]
//   - [IGTAGX2StreamDataTimelineProcessor.ProcessStreamData]
//   - [IGTAGX2StreamDataTimelineProcessor.ProcessTimelineStreamedResult]
//   - [IGTAGX2StreamDataTimelineProcessor.SaveAddressListSizeFilename]
//   - [IGTAGX2StreamDataTimelineProcessor.SaveAddressMappingsFilename]
//   - [IGTAGX2StreamDataTimelineProcessor.SaveSampleDataFromGPURawCountersSizeFilenameWithTimeStamps]
//   - [IGTAGX2StreamDataTimelineProcessor.TimelineInfo]
//   - [IGTAGX2StreamDataTimelineProcessor.WaitUntilFinished]
//   - [IGTAGX2StreamDataTimelineProcessor.InitWithStreamData]
type IGTAGX2StreamDataTimelineProcessor interface {
	objectivec.IObject

	// Topic: Methods

	_addDerivedCounterDataToTimelineInfoForWorkloadInfo(info objectivec.IObject, info2 objectivec.IObject)
	_calculatePerSampleActiveShadersForWorkloadInfo(shaders objectivec.IObject, info objectivec.IObject)
	_calculatePerSampleAggregatedActiveShadersForWorkloadInfo(shaders objectivec.IObject, info objectivec.IObject)
	_updateShaderTimelineInfoWithShaderTimelineDataExForGPUTimelineInfoWithEncoderGlobalTraceIdToStateMirrorIdMappingForRingBuffer(info objectivec.IObject, ex objectivec.IObject, info2 objectivec.IObject, mapping unsafe.Pointer, buffer uint32)
	Process(process objectivec.IObject)
	ProcessStreamData()
	ProcessTimelineStreamedResult(result objectivec.IObject) objectivec.IObject
	SaveAddressListSizeFilename(list *GTAGX2ShaderProfilerProgramAddress, size uint32, filename objectivec.IObject)
	SaveAddressMappingsFilename(mappings objectivec.IObject, filename objectivec.IObject)
	SaveSampleDataFromGPURawCountersSizeFilenameWithTimeStamps(counters *uint64, size uint32, filename objectivec.IObject, stamps *uint64)
	TimelineInfo() IDYWorkloadGPUTimelineInfo
	WaitUntilFinished()
	InitWithStreamData(data objectivec.IObject) GTAGX2StreamDataTimelineProcessor
}

// Init initializes the instance.
func (g GTAGX2StreamDataTimelineProcessor) Init() GTAGX2StreamDataTimelineProcessor {
	rv := objc.SendIfResponds[GTAGX2StreamDataTimelineProcessor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTAGX2StreamDataTimelineProcessor) Autorelease() GTAGX2StreamDataTimelineProcessor {
	rv := objc.SendIfResponds[GTAGX2StreamDataTimelineProcessor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTAGX2StreamDataTimelineProcessor creates a new GTAGX2StreamDataTimelineProcessor instance.
func NewGTAGX2StreamDataTimelineProcessor() GTAGX2StreamDataTimelineProcessor {
	class := getGTAGX2StreamDataTimelineProcessorClass()
	rv := objc.SendIfResponds[GTAGX2StreamDataTimelineProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTAGX2StreamDataTimelineProcessorWithStreamData(data objectivec.IObject) GTAGX2StreamDataTimelineProcessor {
	instance := getGTAGX2StreamDataTimelineProcessorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithStreamData:"), data)
	return GTAGX2StreamDataTimelineProcessorFromID(rv)
}

func (g GTAGX2StreamDataTimelineProcessor) _addDerivedCounterDataToTimelineInfoForWorkloadInfo(info objectivec.IObject, info2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("_addDerivedCounterDataToTimelineInfo:forWorkloadInfo:"), info, info2)
}

// AddDerivedCounterDataToTimelineInfoForWorkloadInfo is an exported wrapper for the private method _addDerivedCounterDataToTimelineInfoForWorkloadInfo.
func (g GTAGX2StreamDataTimelineProcessor) AddDerivedCounterDataToTimelineInfoForWorkloadInfo(info objectivec.IObject, info2 objectivec.IObject) error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_addDerivedCounterDataToTimelineInfo:forWorkloadInfo:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_addDerivedCounterDataToTimelineInfo:forWorkloadInfo:"}
		return err
	}
	g._addDerivedCounterDataToTimelineInfoForWorkloadInfo(info, info2)
	return nil
}

// CanAddDerivedCounterDataToTimelineInfoForWorkloadInfo reports whether the receiver responds to the private selector _addDerivedCounterDataToTimelineInfo:forWorkloadInfo:.
func (g GTAGX2StreamDataTimelineProcessor) CanAddDerivedCounterDataToTimelineInfoForWorkloadInfo() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_addDerivedCounterDataToTimelineInfo:forWorkloadInfo:"))
}
func (g GTAGX2StreamDataTimelineProcessor) _calculatePerSampleActiveShadersForWorkloadInfo(shaders objectivec.IObject, info objectivec.IObject) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("_calculatePerSampleActiveShaders:forWorkloadInfo:"), shaders, info)
}

// CalculatePerSampleActiveShadersForWorkloadInfo is an exported wrapper for the private method _calculatePerSampleActiveShadersForWorkloadInfo.
func (g GTAGX2StreamDataTimelineProcessor) CalculatePerSampleActiveShadersForWorkloadInfo(shaders objectivec.IObject, info objectivec.IObject) error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_calculatePerSampleActiveShaders:forWorkloadInfo:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_calculatePerSampleActiveShaders:forWorkloadInfo:"}
		return err
	}
	g._calculatePerSampleActiveShadersForWorkloadInfo(shaders, info)
	return nil
}

// CanCalculatePerSampleActiveShadersForWorkloadInfo reports whether the receiver responds to the private selector _calculatePerSampleActiveShaders:forWorkloadInfo:.
func (g GTAGX2StreamDataTimelineProcessor) CanCalculatePerSampleActiveShadersForWorkloadInfo() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_calculatePerSampleActiveShaders:forWorkloadInfo:"))
}
func (g GTAGX2StreamDataTimelineProcessor) _calculatePerSampleAggregatedActiveShadersForWorkloadInfo(shaders objectivec.IObject, info objectivec.IObject) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("_calculatePerSampleAggregatedActiveShaders:forWorkloadInfo:"), shaders, info)
}

// CalculatePerSampleAggregatedActiveShadersForWorkloadInfo is an exported wrapper for the private method _calculatePerSampleAggregatedActiveShadersForWorkloadInfo.
func (g GTAGX2StreamDataTimelineProcessor) CalculatePerSampleAggregatedActiveShadersForWorkloadInfo(shaders objectivec.IObject, info objectivec.IObject) error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_calculatePerSampleAggregatedActiveShaders:forWorkloadInfo:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_calculatePerSampleAggregatedActiveShaders:forWorkloadInfo:"}
		return err
	}
	g._calculatePerSampleAggregatedActiveShadersForWorkloadInfo(shaders, info)
	return nil
}

// CanCalculatePerSampleAggregatedActiveShadersForWorkloadInfo reports whether the receiver responds to the private selector _calculatePerSampleAggregatedActiveShaders:forWorkloadInfo:.
func (g GTAGX2StreamDataTimelineProcessor) CanCalculatePerSampleAggregatedActiveShadersForWorkloadInfo() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_calculatePerSampleAggregatedActiveShaders:forWorkloadInfo:"))
}
func (g GTAGX2StreamDataTimelineProcessor) _updateShaderTimelineInfoWithShaderTimelineDataExForGPUTimelineInfoWithEncoderGlobalTraceIdToStateMirrorIdMappingForRingBuffer(info objectivec.IObject, ex objectivec.IObject, info2 objectivec.IObject, mapping unsafe.Pointer, buffer uint32) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("_updateShaderTimelineInfo:withShaderTimelineDataEx:forGPUTimelineInfo:withEncoderGlobalTraceIdToStateMirrorIdMapping:forRingBuffer:"), info, ex, info2, mapping, buffer)
}

// UpdateShaderTimelineInfoWithShaderTimelineDataExForGPUTimelineInfoWithEncoderGlobalTraceIdToStateMirrorIdMappingForRingBuffer is an exported wrapper for the private method _updateShaderTimelineInfoWithShaderTimelineDataExForGPUTimelineInfoWithEncoderGlobalTraceIdToStateMirrorIdMappingForRingBuffer.
func (g GTAGX2StreamDataTimelineProcessor) UpdateShaderTimelineInfoWithShaderTimelineDataExForGPUTimelineInfoWithEncoderGlobalTraceIdToStateMirrorIdMappingForRingBuffer(info objectivec.IObject, ex objectivec.IObject, info2 objectivec.IObject, mapping unsafe.Pointer, buffer uint32) error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_updateShaderTimelineInfo:withShaderTimelineDataEx:forGPUTimelineInfo:withEncoderGlobalTraceIdToStateMirrorIdMapping:forRingBuffer:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_updateShaderTimelineInfo:withShaderTimelineDataEx:forGPUTimelineInfo:withEncoderGlobalTraceIdToStateMirrorIdMapping:forRingBuffer:"}
		return err
	}
	g._updateShaderTimelineInfoWithShaderTimelineDataExForGPUTimelineInfoWithEncoderGlobalTraceIdToStateMirrorIdMappingForRingBuffer(info, ex, info2, mapping, buffer)
	return nil
}

// CanUpdateShaderTimelineInfoWithShaderTimelineDataExForGPUTimelineInfoWithEncoderGlobalTraceIdToStateMirrorIdMappingForRingBuffer reports whether the receiver responds to the private selector _updateShaderTimelineInfo:withShaderTimelineDataEx:forGPUTimelineInfo:withEncoderGlobalTraceIdToStateMirrorIdMapping:forRingBuffer:.
func (g GTAGX2StreamDataTimelineProcessor) CanUpdateShaderTimelineInfoWithShaderTimelineDataExForGPUTimelineInfoWithEncoderGlobalTraceIdToStateMirrorIdMappingForRingBuffer() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_updateShaderTimelineInfo:withShaderTimelineDataEx:forGPUTimelineInfo:withEncoderGlobalTraceIdToStateMirrorIdMapping:forRingBuffer:"))
}
func (g GTAGX2StreamDataTimelineProcessor) Process(process objectivec.IObject) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("process:"), process)
}
func (g GTAGX2StreamDataTimelineProcessor) ProcessStreamData() {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("processStreamData"))
}
func (g GTAGX2StreamDataTimelineProcessor) ProcessTimelineStreamedResult(result objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("processTimelineStreamedResult:"), result)
	return objectivec.Object{ID: rv}
}
func (g GTAGX2StreamDataTimelineProcessor) SaveAddressListSizeFilename(list *GTAGX2ShaderProfilerProgramAddress, size uint32, filename objectivec.IObject) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("saveAddressList:size:filename:"), unsafe.Pointer(list), size, filename)
}
func (g GTAGX2StreamDataTimelineProcessor) SaveAddressMappingsFilename(mappings objectivec.IObject, filename objectivec.IObject) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("saveAddressMappings:filename:"), mappings, filename)
}
func (g GTAGX2StreamDataTimelineProcessor) SaveSampleDataFromGPURawCountersSizeFilenameWithTimeStamps(counters *uint64, size uint32, filename objectivec.IObject, stamps *uint64) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("saveSampleDataFromGPURawCounters:size:filename:withTimeStamps:"), unsafe.Pointer(counters), size, filename, unsafe.Pointer(stamps))
}
func (g GTAGX2StreamDataTimelineProcessor) WaitUntilFinished() {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("waitUntilFinished"))
}
func (g GTAGX2StreamDataTimelineProcessor) InitWithStreamData(data objectivec.IObject) GTAGX2StreamDataTimelineProcessor {
	rv := objc.SendIfResponds[GTAGX2StreamDataTimelineProcessor](g.ID, objc.Sel("initWithStreamData:"), data)
	return rv
}

func (g GTAGX2StreamDataTimelineProcessor) TimelineInfo() IDYWorkloadGPUTimelineInfo {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("timelineInfo"))
	return DYWorkloadGPUTimelineInfoFromID(objc.ID(rv))
}

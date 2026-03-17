// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerStreamDataProcessor] class.
var (
	_GTShaderProfilerStreamDataProcessorClass     GTShaderProfilerStreamDataProcessorClass
	_GTShaderProfilerStreamDataProcessorClassOnce sync.Once
)

func getGTShaderProfilerStreamDataProcessorClass() GTShaderProfilerStreamDataProcessorClass {
	_GTShaderProfilerStreamDataProcessorClassOnce.Do(func() {
		_GTShaderProfilerStreamDataProcessorClass = GTShaderProfilerStreamDataProcessorClass{class: objc.GetClass("GTShaderProfilerStreamDataProcessor")}
	})
	return _GTShaderProfilerStreamDataProcessorClass
}

// GetGTShaderProfilerStreamDataProcessorClass returns the class object for GTShaderProfilerStreamDataProcessor.
func GetGTShaderProfilerStreamDataProcessorClass() GTShaderProfilerStreamDataProcessorClass {
	return getGTShaderProfilerStreamDataProcessorClass()
}

type GTShaderProfilerStreamDataProcessorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerStreamDataProcessorClass) Alloc() GTShaderProfilerStreamDataProcessor {
	rv := objc.Send[GTShaderProfilerStreamDataProcessor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTShaderProfilerStreamDataProcessor.GpuGeneration]
//   - [GTShaderProfilerStreamDataProcessor.ProcessAPSTimelineData]
//   - [GTShaderProfilerStreamDataProcessor.ProcessShaderProfilerStreamData]
//   - [GTShaderProfilerStreamDataProcessor.ProcessStreamData]
//   - [GTShaderProfilerStreamDataProcessor.ProcessTimelineStreamData]
//   - [GTShaderProfilerStreamDataProcessor.Result]
//   - [GTShaderProfilerStreamDataProcessor.StreamData]
//   - [GTShaderProfilerStreamDataProcessor.WaitUntilFinished]
//   - [GTShaderProfilerStreamDataProcessor.WaitUntilShaderProfilerFinished]
//   - [GTShaderProfilerStreamDataProcessor.WaitUntilTimelineFinished]
//   - [GTShaderProfilerStreamDataProcessor.Delegate]
//   - [GTShaderProfilerStreamDataProcessor.SetDelegate]
//   - [GTShaderProfilerStreamDataProcessor.IsaPrinter]
//   - [GTShaderProfilerStreamDataProcessor.SetIsaPrinter]
//   - [GTShaderProfilerStreamDataProcessor.MioData]
//   - [GTShaderProfilerStreamDataProcessor.ProcessAPSCostData]
//   - [GTShaderProfilerStreamDataProcessor.ProcessBatchIDFilteringData]
//   - [GTShaderProfilerStreamDataProcessor.ProcessBatchIdFilteredCounterStreamData]
//   - [GTShaderProfilerStreamDataProcessor.ProcessGPUTimelineData]
//   - [GTShaderProfilerStreamDataProcessor.ProcessShaderProfilerData]
//   - [GTShaderProfilerStreamDataProcessor.SetupForBatchIDFilteringCounters]
//   - [GTShaderProfilerStreamDataProcessor.StreamDataProcessorBatchIdFilteredCountersUpdatedObserverInfo]
//   - [GTShaderProfilerStreamDataProcessor.WaitUntilBatchIDCounterFinished]
//   - [GTShaderProfilerStreamDataProcessor.InitWithStreamDataLlvmHelperPath]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor
type GTShaderProfilerStreamDataProcessor struct {
	objectivec.Object
}

// GTShaderProfilerStreamDataProcessorFromID constructs a [GTShaderProfilerStreamDataProcessor] from an objc.ID.
func GTShaderProfilerStreamDataProcessorFromID(id objc.ID) GTShaderProfilerStreamDataProcessor {
	return GTShaderProfilerStreamDataProcessor{objectivec.Object{ID: id}}
}
// Ensure GTShaderProfilerStreamDataProcessor implements IGTShaderProfilerStreamDataProcessor.
var _ IGTShaderProfilerStreamDataProcessor = GTShaderProfilerStreamDataProcessor{}

// An interface definition for the [GTShaderProfilerStreamDataProcessor] class.
//
// # Methods
//
//   - [IGTShaderProfilerStreamDataProcessor.GpuGeneration]
//   - [IGTShaderProfilerStreamDataProcessor.ProcessAPSTimelineData]
//   - [IGTShaderProfilerStreamDataProcessor.ProcessShaderProfilerStreamData]
//   - [IGTShaderProfilerStreamDataProcessor.ProcessStreamData]
//   - [IGTShaderProfilerStreamDataProcessor.ProcessTimelineStreamData]
//   - [IGTShaderProfilerStreamDataProcessor.Result]
//   - [IGTShaderProfilerStreamDataProcessor.StreamData]
//   - [IGTShaderProfilerStreamDataProcessor.WaitUntilFinished]
//   - [IGTShaderProfilerStreamDataProcessor.WaitUntilShaderProfilerFinished]
//   - [IGTShaderProfilerStreamDataProcessor.WaitUntilTimelineFinished]
//   - [IGTShaderProfilerStreamDataProcessor.Delegate]
//   - [IGTShaderProfilerStreamDataProcessor.SetDelegate]
//   - [IGTShaderProfilerStreamDataProcessor.IsaPrinter]
//   - [IGTShaderProfilerStreamDataProcessor.SetIsaPrinter]
//   - [IGTShaderProfilerStreamDataProcessor.MioData]
//   - [IGTShaderProfilerStreamDataProcessor.ProcessAPSCostData]
//   - [IGTShaderProfilerStreamDataProcessor.ProcessBatchIDFilteringData]
//   - [IGTShaderProfilerStreamDataProcessor.ProcessBatchIdFilteredCounterStreamData]
//   - [IGTShaderProfilerStreamDataProcessor.ProcessGPUTimelineData]
//   - [IGTShaderProfilerStreamDataProcessor.ProcessShaderProfilerData]
//   - [IGTShaderProfilerStreamDataProcessor.SetupForBatchIDFilteringCounters]
//   - [IGTShaderProfilerStreamDataProcessor.StreamDataProcessorBatchIdFilteredCountersUpdatedObserverInfo]
//   - [IGTShaderProfilerStreamDataProcessor.WaitUntilBatchIDCounterFinished]
//   - [IGTShaderProfilerStreamDataProcessor.InitWithStreamDataLlvmHelperPath]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor
type IGTShaderProfilerStreamDataProcessor interface {
	objectivec.IObject

	// Topic: Methods

	GpuGeneration() uint32
	ProcessAPSTimelineData() bool
	ProcessShaderProfilerStreamData()
	ProcessStreamData()
	ProcessTimelineStreamData()
	Result() objectivec.IObject
	StreamData() IGTShaderProfilerStreamData
	WaitUntilFinished()
	WaitUntilShaderProfilerFinished()
	WaitUntilTimelineFinished()
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	IsaPrinter() unsafe.Pointer
	SetIsaPrinter(value unsafe.Pointer)
	MioData() unsafe.Pointer
	ProcessAPSCostData() bool
	ProcessBatchIDFilteringData(data objectivec.IObject)
	ProcessBatchIdFilteredCounterStreamData()
	ProcessGPUTimelineData(data objectivec.IObject)
	ProcessShaderProfilerData(data objectivec.IObject)
	SetupForBatchIDFilteringCounters(counters objectivec.IObject) bool
	StreamDataProcessorBatchIdFilteredCountersUpdatedObserverInfo(updated objectivec.IObject, info objectivec.IObject)
	WaitUntilBatchIDCounterFinished()
	InitWithStreamDataLlvmHelperPath(data IGTShaderProfilerStreamData, path string) GTShaderProfilerStreamDataProcessor
}

// Init initializes the instance.
func (g GTShaderProfilerStreamDataProcessor) Init() GTShaderProfilerStreamDataProcessor {
	rv := objc.Send[GTShaderProfilerStreamDataProcessor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerStreamDataProcessor) Autorelease() GTShaderProfilerStreamDataProcessor {
	rv := objc.Send[GTShaderProfilerStreamDataProcessor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerStreamDataProcessor creates a new GTShaderProfilerStreamDataProcessor instance.
func NewGTShaderProfilerStreamDataProcessor() GTShaderProfilerStreamDataProcessor {
	class := getGTShaderProfilerStreamDataProcessorClass()
	rv := objc.Send[GTShaderProfilerStreamDataProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/initWithStreamData:llvmHelperPath:
func NewGTShaderProfilerStreamDataProcessorWithStreamDataLlvmHelperPath(data IGTShaderProfilerStreamData, path string) GTShaderProfilerStreamDataProcessor {
	instance := getGTShaderProfilerStreamDataProcessorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStreamData:llvmHelperPath:"), data, objc.String(path))
	return GTShaderProfilerStreamDataProcessorFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/processAPSTimelineData
func (g GTShaderProfilerStreamDataProcessor) ProcessAPSTimelineData() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("processAPSTimelineData"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/processShaderProfilerStreamData
func (g GTShaderProfilerStreamDataProcessor) ProcessShaderProfilerStreamData() {
	objc.Send[objc.ID](g.ID, objc.Sel("processShaderProfilerStreamData"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/processStreamData
func (g GTShaderProfilerStreamDataProcessor) ProcessStreamData() {
	objc.Send[objc.ID](g.ID, objc.Sel("processStreamData"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/processTimelineStreamData
func (g GTShaderProfilerStreamDataProcessor) ProcessTimelineStreamData() {
	objc.Send[objc.ID](g.ID, objc.Sel("processTimelineStreamData"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/waitUntilFinished
func (g GTShaderProfilerStreamDataProcessor) WaitUntilFinished() {
	objc.Send[objc.ID](g.ID, objc.Sel("waitUntilFinished"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/waitUntilShaderProfilerFinished
func (g GTShaderProfilerStreamDataProcessor) WaitUntilShaderProfilerFinished() {
	objc.Send[objc.ID](g.ID, objc.Sel("waitUntilShaderProfilerFinished"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/waitUntilTimelineFinished
func (g GTShaderProfilerStreamDataProcessor) WaitUntilTimelineFinished() {
	objc.Send[objc.ID](g.ID, objc.Sel("waitUntilTimelineFinished"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/processAPSCostData
func (g GTShaderProfilerStreamDataProcessor) ProcessAPSCostData() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("processAPSCostData"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/processBatchIDFilteringData:
func (g GTShaderProfilerStreamDataProcessor) ProcessBatchIDFilteringData(data objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("processBatchIDFilteringData:"), data)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/processBatchIdFilteredCounterStreamData
func (g GTShaderProfilerStreamDataProcessor) ProcessBatchIdFilteredCounterStreamData() {
	objc.Send[objc.ID](g.ID, objc.Sel("processBatchIdFilteredCounterStreamData"))
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/processGPUTimelineData:
func (g GTShaderProfilerStreamDataProcessor) ProcessGPUTimelineData(data objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("processGPUTimelineData:"), data)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/processShaderProfilerData:
func (g GTShaderProfilerStreamDataProcessor) ProcessShaderProfilerData(data objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("processShaderProfilerData:"), data)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/setupForBatchIDFilteringCounters:
func (g GTShaderProfilerStreamDataProcessor) SetupForBatchIDFilteringCounters(counters objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("setupForBatchIDFilteringCounters:"), counters)
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/streamDataProcessorBatchIdFilteredCountersUpdated:observerInfo:
func (g GTShaderProfilerStreamDataProcessor) StreamDataProcessorBatchIdFilteredCountersUpdatedObserverInfo(updated objectivec.IObject, info objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("streamDataProcessorBatchIdFilteredCountersUpdated:observerInfo:"), updated, info)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/waitUntilBatchIDCounterFinished
func (g GTShaderProfilerStreamDataProcessor) WaitUntilBatchIDCounterFinished() {
	objc.Send[objc.ID](g.ID, objc.Sel("waitUntilBatchIDCounterFinished"))
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/initWithStreamData:llvmHelperPath:
func (g GTShaderProfilerStreamDataProcessor) InitWithStreamDataLlvmHelperPath(data IGTShaderProfilerStreamData, path string) GTShaderProfilerStreamDataProcessor {
	rv := objc.Send[GTShaderProfilerStreamDataProcessor](g.ID, objc.Sel("initWithStreamData:llvmHelperPath:"), data, objc.String(path))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/gpuGeneration
func (g GTShaderProfilerStreamDataProcessor) GpuGeneration() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("gpuGeneration"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/result
func (g GTShaderProfilerStreamDataProcessor) Result() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("result"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/streamData
func (g GTShaderProfilerStreamDataProcessor) StreamData() IGTShaderProfilerStreamData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("streamData"))
	return GTShaderProfilerStreamDataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/delegate
func (g GTShaderProfilerStreamDataProcessor) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("delegate"))
	return rv
}
func (g GTShaderProfilerStreamDataProcessor) SetDelegate(value unsafe.Pointer) {
	objc.Send[struct{}](g.ID, objc.Sel("setDelegate:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/isaPrinter
func (g GTShaderProfilerStreamDataProcessor) IsaPrinter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("isaPrinter"))
	return rv
}
func (g GTShaderProfilerStreamDataProcessor) SetIsaPrinter(value unsafe.Pointer) {
	objc.Send[struct{}](g.ID, objc.Sel("setIsaPrinter:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataProcessor/mioData
func (g GTShaderProfilerStreamDataProcessor) MioData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("mioData"))
	return rv
}


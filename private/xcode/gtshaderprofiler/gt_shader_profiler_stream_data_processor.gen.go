// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

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

// Class returns the underlying Objective-C class pointer.
func (gc GTShaderProfilerStreamDataProcessorClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerStreamDataProcessorClass) Alloc() GTShaderProfilerStreamDataProcessor {
	rv := objc.SendIfResponds[GTShaderProfilerStreamDataProcessor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

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
	MioData() IGTMioTraceData
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
	rv := objc.SendIfResponds[GTShaderProfilerStreamDataProcessor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerStreamDataProcessor) Autorelease() GTShaderProfilerStreamDataProcessor {
	rv := objc.SendIfResponds[GTShaderProfilerStreamDataProcessor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerStreamDataProcessor creates a new GTShaderProfilerStreamDataProcessor instance.
func NewGTShaderProfilerStreamDataProcessor() GTShaderProfilerStreamDataProcessor {
	class := getGTShaderProfilerStreamDataProcessorClass()
	rv := objc.SendIfResponds[GTShaderProfilerStreamDataProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTShaderProfilerStreamDataProcessorWithStreamDataLlvmHelperPath(data IGTShaderProfilerStreamData, path string) GTShaderProfilerStreamDataProcessor {
	instance := getGTShaderProfilerStreamDataProcessorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithStreamData:llvmHelperPath:"), data, objc.String(path))
	return GTShaderProfilerStreamDataProcessorFromID(rv)
}

func (g GTShaderProfilerStreamDataProcessor) ProcessAPSTimelineData() bool {
	rv := objc.SendIfResponds[bool](g.ID, objc.Sel("processAPSTimelineData"))
	return rv
}
func (g GTShaderProfilerStreamDataProcessor) ProcessShaderProfilerStreamData() {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("processShaderProfilerStreamData"))
}
func (g GTShaderProfilerStreamDataProcessor) ProcessStreamData() {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("processStreamData"))
}
func (g GTShaderProfilerStreamDataProcessor) ProcessTimelineStreamData() {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("processTimelineStreamData"))
}
func (g GTShaderProfilerStreamDataProcessor) WaitUntilFinished() {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("waitUntilFinished"))
}
func (g GTShaderProfilerStreamDataProcessor) WaitUntilShaderProfilerFinished() {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("waitUntilShaderProfilerFinished"))
}
func (g GTShaderProfilerStreamDataProcessor) WaitUntilTimelineFinished() {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("waitUntilTimelineFinished"))
}
func (g GTShaderProfilerStreamDataProcessor) ProcessAPSCostData() bool {
	rv := objc.SendIfResponds[bool](g.ID, objc.Sel("processAPSCostData"))
	return rv
}
func (g GTShaderProfilerStreamDataProcessor) ProcessBatchIDFilteringData(data objectivec.IObject) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("processBatchIDFilteringData:"), data)
}
func (g GTShaderProfilerStreamDataProcessor) ProcessBatchIdFilteredCounterStreamData() {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("processBatchIdFilteredCounterStreamData"))
}
func (g GTShaderProfilerStreamDataProcessor) ProcessGPUTimelineData(data objectivec.IObject) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("processGPUTimelineData:"), data)
}
func (g GTShaderProfilerStreamDataProcessor) ProcessShaderProfilerData(data objectivec.IObject) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("processShaderProfilerData:"), data)
}
func (g GTShaderProfilerStreamDataProcessor) SetupForBatchIDFilteringCounters(counters objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](g.ID, objc.Sel("setupForBatchIDFilteringCounters:"), counters)
	return rv
}
func (g GTShaderProfilerStreamDataProcessor) StreamDataProcessorBatchIdFilteredCountersUpdatedObserverInfo(updated objectivec.IObject, info objectivec.IObject) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("streamDataProcessorBatchIdFilteredCountersUpdated:observerInfo:"), updated, info)
}
func (g GTShaderProfilerStreamDataProcessor) WaitUntilBatchIDCounterFinished() {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("waitUntilBatchIDCounterFinished"))
}
func (g GTShaderProfilerStreamDataProcessor) InitWithStreamDataLlvmHelperPath(data IGTShaderProfilerStreamData, path string) GTShaderProfilerStreamDataProcessor {
	rv := objc.SendIfResponds[GTShaderProfilerStreamDataProcessor](g.ID, objc.Sel("initWithStreamData:llvmHelperPath:"), data, objc.String(path))
	return rv
}

func (g GTShaderProfilerStreamDataProcessor) GpuGeneration() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("gpuGeneration"))
	return rv
}
func (g GTShaderProfilerStreamDataProcessor) Result() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("result"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerStreamDataProcessor) StreamData() IGTShaderProfilerStreamData {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("streamData"))
	return GTShaderProfilerStreamDataFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamDataProcessor) Delegate() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("delegate"))
	return rv
}
func (g GTShaderProfilerStreamDataProcessor) SetDelegate(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](g.ID, objc.Sel("setDelegate:"), value)
}
func (g GTShaderProfilerStreamDataProcessor) IsaPrinter() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("isaPrinter"))
	return rv
}
func (g GTShaderProfilerStreamDataProcessor) SetIsaPrinter(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](g.ID, objc.Sel("setIsaPrinter:"), value)
}
func (g GTShaderProfilerStreamDataProcessor) MioData() IGTMioTraceData {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("mioData"))
	return GTMioTraceDataFromID(objc.ID(rv))
}

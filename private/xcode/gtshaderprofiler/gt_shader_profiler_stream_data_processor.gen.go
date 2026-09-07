// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
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

// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMutableShaderProfilerStreamData] class.
var (
	_GTMutableShaderProfilerStreamDataClass     GTMutableShaderProfilerStreamDataClass
	_GTMutableShaderProfilerStreamDataClassOnce sync.Once
)

func getGTMutableShaderProfilerStreamDataClass() GTMutableShaderProfilerStreamDataClass {
	_GTMutableShaderProfilerStreamDataClassOnce.Do(func() {
		_GTMutableShaderProfilerStreamDataClass = GTMutableShaderProfilerStreamDataClass{class: objc.GetClass("GTMutableShaderProfilerStreamData")}
	})
	return _GTMutableShaderProfilerStreamDataClass
}

// GetGTMutableShaderProfilerStreamDataClass returns the class object for GTMutableShaderProfilerStreamData.
func GetGTMutableShaderProfilerStreamDataClass() GTMutableShaderProfilerStreamDataClass {
	return getGTMutableShaderProfilerStreamDataClass()
}

type GTMutableShaderProfilerStreamDataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMutableShaderProfilerStreamDataClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMutableShaderProfilerStreamDataClass) Alloc() GTMutableShaderProfilerStreamData {
	rv := objc.Send[GTMutableShaderProfilerStreamData](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMutableShaderProfilerStreamData._commonInit]
//   - [GTMutableShaderProfilerStreamData._copyForAddAPSDataPrefix]
//   - [GTMutableShaderProfilerStreamData.AddAPSCounterData]
//   - [GTMutableShaderProfilerStreamData.AddAPSData]
//   - [GTMutableShaderProfilerStreamData.AddAPSTimelineData]
//   - [GTMutableShaderProfilerStreamData.AddBatchIdFilteredCounterData]
//   - [GTMutableShaderProfilerStreamData.AddCommandBuffersCount]
//   - [GTMutableShaderProfilerStreamData.AddEncodersCount]
//   - [GTMutableShaderProfilerStreamData.AddGPUCommandsCount]
//   - [GTMutableShaderProfilerStreamData.AddGPUTimelineData]
//   - [GTMutableShaderProfilerStreamData.AddPipelinePerformanceStatisticsData]
//   - [GTMutableShaderProfilerStreamData.AddPipelineStatesCount]
//   - [GTMutableShaderProfilerStreamData.AddShaderFunctionInfoCount]
//   - [GTMutableShaderProfilerStreamData.AddShaderProfilerData]
//   - [GTMutableShaderProfilerStreamData.AddString]
//   - [GTMutableShaderProfilerStreamData.RemoveAPSCounterData]
//   - [GTMutableShaderProfilerStreamData.RemoveAPSData]
//   - [GTMutableShaderProfilerStreamData.RemoveAPSTimelineData]
//   - [GTMutableShaderProfilerStreamData.SetDataSourceHasUnusedResourcesCaptureRange]
//   - [GTMutableShaderProfilerStreamData.SetNumBlitCalls]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData
type GTMutableShaderProfilerStreamData struct {
	GTShaderProfilerStreamData
}

// GTMutableShaderProfilerStreamDataFromID constructs a [GTMutableShaderProfilerStreamData] from an objc.ID.
func GTMutableShaderProfilerStreamDataFromID(id objc.ID) GTMutableShaderProfilerStreamData {
	return GTMutableShaderProfilerStreamData{GTShaderProfilerStreamData: GTShaderProfilerStreamDataFromID(id)}
}

// Ensure GTMutableShaderProfilerStreamData implements IGTMutableShaderProfilerStreamData.
var _ IGTMutableShaderProfilerStreamData = GTMutableShaderProfilerStreamData{}

// An interface definition for the [GTMutableShaderProfilerStreamData] class.
//
// # Methods
//
//   - [IGTMutableShaderProfilerStreamData._commonInit]
//   - [IGTMutableShaderProfilerStreamData._copyForAddAPSDataPrefix]
//   - [IGTMutableShaderProfilerStreamData.AddAPSCounterData]
//   - [IGTMutableShaderProfilerStreamData.AddAPSData]
//   - [IGTMutableShaderProfilerStreamData.AddAPSTimelineData]
//   - [IGTMutableShaderProfilerStreamData.AddBatchIdFilteredCounterData]
//   - [IGTMutableShaderProfilerStreamData.AddCommandBuffersCount]
//   - [IGTMutableShaderProfilerStreamData.AddEncodersCount]
//   - [IGTMutableShaderProfilerStreamData.AddGPUCommandsCount]
//   - [IGTMutableShaderProfilerStreamData.AddGPUTimelineData]
//   - [IGTMutableShaderProfilerStreamData.AddPipelinePerformanceStatisticsData]
//   - [IGTMutableShaderProfilerStreamData.AddPipelineStatesCount]
//   - [IGTMutableShaderProfilerStreamData.AddShaderFunctionInfoCount]
//   - [IGTMutableShaderProfilerStreamData.AddShaderProfilerData]
//   - [IGTMutableShaderProfilerStreamData.AddString]
//   - [IGTMutableShaderProfilerStreamData.RemoveAPSCounterData]
//   - [IGTMutableShaderProfilerStreamData.RemoveAPSData]
//   - [IGTMutableShaderProfilerStreamData.RemoveAPSTimelineData]
//   - [IGTMutableShaderProfilerStreamData.SetDataSourceHasUnusedResourcesCaptureRange]
//   - [IGTMutableShaderProfilerStreamData.SetNumBlitCalls]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData
type IGTMutableShaderProfilerStreamData interface {
	IGTShaderProfilerStreamData

	// Topic: Methods

	_commonInit()
	_copyForAddAPSDataPrefix(aPSData objectivec.IObject, prefix objectivec.IObject) objectivec.IObject
	AddAPSCounterData(data objectivec.IObject) bool
	AddAPSData(aPSData objectivec.IObject) bool
	AddAPSTimelineData(data objectivec.IObject) bool
	AddBatchIdFilteredCounterData(data objectivec.IObject) bool
	AddCommandBuffersCount(buffers objectivec.IObject, count uint64)
	AddEncodersCount(encoders objectivec.IObject, count uint64)
	AddGPUCommandsCount(gPUCommands objectivec.IObject, count uint64)
	AddGPUTimelineData(data objectivec.IObject) bool
	AddPipelinePerformanceStatisticsData(data objectivec.IObject)
	AddPipelineStatesCount(states objectivec.IObject, count uint64)
	AddShaderFunctionInfoCount(info objectivec.IObject, count uint64)
	AddShaderProfilerData(data objectivec.IObject) bool
	AddString(string_ objectivec.IObject) uint64
	RemoveAPSCounterData()
	RemoveAPSData()
	RemoveAPSTimelineData()
	SetDataSourceHasUnusedResourcesCaptureRange(resources bool, range_ foundation.NSRange)
	SetNumBlitCalls(calls uint64)
}

// Init initializes the instance.
func (g GTMutableShaderProfilerStreamData) Init() GTMutableShaderProfilerStreamData {
	rv := objc.Send[GTMutableShaderProfilerStreamData](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMutableShaderProfilerStreamData) Autorelease() GTMutableShaderProfilerStreamData {
	rv := objc.Send[GTMutableShaderProfilerStreamData](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMutableShaderProfilerStreamData creates a new GTMutableShaderProfilerStreamData instance.
func NewGTMutableShaderProfilerStreamData() GTMutableShaderProfilerStreamData {
	class := getGTMutableShaderProfilerStreamDataClass()
	rv := objc.Send[GTMutableShaderProfilerStreamData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamData/initWithCoder:
func NewGTMutableShaderProfilerStreamDataWithCoder(coder objectivec.IObject) GTMutableShaderProfilerStreamData {
	instance := getGTMutableShaderProfilerStreamDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTMutableShaderProfilerStreamDataFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/initWithNewFileFormatV2Support:
func NewGTMutableShaderProfilerStreamDataWithNewFileFormatV2Support(v2Support bool) GTMutableShaderProfilerStreamData {
	instance := getGTMutableShaderProfilerStreamDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNewFileFormatV2Support:"), v2Support)
	return GTMutableShaderProfilerStreamDataFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamData/initWithPreSiBundle:
func NewGTMutableShaderProfilerStreamDataWithPreSiBundle(bundle objectivec.IObject) GTMutableShaderProfilerStreamData {
	instance := getGTMutableShaderProfilerStreamDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPreSiBundle:"), bundle)
	return GTMutableShaderProfilerStreamDataFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/_commonInit
func (g GTMutableShaderProfilerStreamData) _commonInit() {
	objc.Send[objc.ID](g.ID, objc.Sel("_commonInit"))
}

// CommonInit is an exported wrapper for the private method _commonInit.
func (g GTMutableShaderProfilerStreamData) CommonInit() {
	g._commonInit()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/_copyForAddAPSData:prefix:
func (g GTMutableShaderProfilerStreamData) _copyForAddAPSDataPrefix(aPSData objectivec.IObject, prefix objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_copyForAddAPSData:prefix:"), aPSData, prefix)
	return objectivec.Object{ID: rv}
}

// CopyForAddAPSDataPrefix is an exported wrapper for the private method _copyForAddAPSDataPrefix.
func (g GTMutableShaderProfilerStreamData) CopyForAddAPSDataPrefix(aPSData objectivec.IObject, prefix objectivec.IObject) objectivec.IObject {
	return g._copyForAddAPSDataPrefix(aPSData, prefix)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/addAPSCounterData:
func (g GTMutableShaderProfilerStreamData) AddAPSCounterData(data objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("addAPSCounterData:"), data)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/addAPSData:
func (g GTMutableShaderProfilerStreamData) AddAPSData(aPSData objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("addAPSData:"), aPSData)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/addAPSTimelineData:
func (g GTMutableShaderProfilerStreamData) AddAPSTimelineData(data objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("addAPSTimelineData:"), data)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/addBatchIdFilteredCounterData:
func (g GTMutableShaderProfilerStreamData) AddBatchIdFilteredCounterData(data objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("addBatchIdFilteredCounterData:"), data)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/addCommandBuffers:count:
func (g GTMutableShaderProfilerStreamData) AddCommandBuffersCount(buffers objectivec.IObject, count uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("addCommandBuffers:count:"), buffers, count)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/addEncoders:count:
func (g GTMutableShaderProfilerStreamData) AddEncodersCount(encoders objectivec.IObject, count uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("addEncoders:count:"), encoders, count)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/addGPUCommands:count:
func (g GTMutableShaderProfilerStreamData) AddGPUCommandsCount(gPUCommands objectivec.IObject, count uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("addGPUCommands:count:"), gPUCommands, count)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/addGPUTimelineData:
func (g GTMutableShaderProfilerStreamData) AddGPUTimelineData(data objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("addGPUTimelineData:"), data)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/addPipelinePerformanceStatisticsData:
func (g GTMutableShaderProfilerStreamData) AddPipelinePerformanceStatisticsData(data objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("addPipelinePerformanceStatisticsData:"), data)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/addPipelineStates:count:
func (g GTMutableShaderProfilerStreamData) AddPipelineStatesCount(states objectivec.IObject, count uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("addPipelineStates:count:"), states, count)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/addShaderFunctionInfo:count:
func (g GTMutableShaderProfilerStreamData) AddShaderFunctionInfoCount(info objectivec.IObject, count uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("addShaderFunctionInfo:count:"), info, count)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/addShaderProfilerData:
func (g GTMutableShaderProfilerStreamData) AddShaderProfilerData(data objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("addShaderProfilerData:"), data)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/addString:
func (g GTMutableShaderProfilerStreamData) AddString(string_ objectivec.IObject) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("addString:"), string_)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/removeAPSCounterData
func (g GTMutableShaderProfilerStreamData) RemoveAPSCounterData() {
	objc.Send[objc.ID](g.ID, objc.Sel("removeAPSCounterData"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/removeAPSData
func (g GTMutableShaderProfilerStreamData) RemoveAPSData() {
	objc.Send[objc.ID](g.ID, objc.Sel("removeAPSData"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/removeAPSTimelineData
func (g GTMutableShaderProfilerStreamData) RemoveAPSTimelineData() {
	objc.Send[objc.ID](g.ID, objc.Sel("removeAPSTimelineData"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/setDataSourceHasUnusedResources:captureRange:
func (g GTMutableShaderProfilerStreamData) SetDataSourceHasUnusedResourcesCaptureRange(resources bool, range_ foundation.NSRange) {
	objc.Send[objc.ID](g.ID, objc.Sel("setDataSourceHasUnusedResources:captureRange:"), resources, range_)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMutableShaderProfilerStreamData/setNumBlitCalls:
func (g GTMutableShaderProfilerStreamData) SetNumBlitCalls(calls uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("setNumBlitCalls:"), calls)
}

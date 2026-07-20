// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

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
type IGTMutableShaderProfilerStreamData interface {
	IGTShaderProfilerStreamData

	// Topic: Methods

	_commonInit()
	_copyForAddAPSDataPrefix(aPSData objectivec.IObject, prefix objectivec.IObject) objectivec.IObject
	AddAPSCounterData(data objectivec.IObject) bool
	AddAPSData(aPSData objectivec.IObject) bool
	AddAPSTimelineData(data objectivec.IObject) bool
	AddBatchIdFilteredCounterData(data objectivec.IObject) bool
	AddCommandBuffersCount(buffers unsafe.Pointer, count uint64)
	AddEncodersCount(encoders unsafe.Pointer, count uint64)
	AddGPUCommandsCount(gPUCommands unsafe.Pointer, count uint64)
	AddGPUTimelineData(data objectivec.IObject) bool
	AddPipelinePerformanceStatisticsData(data objectivec.IObject)
	AddPipelineStatesCount(states unsafe.Pointer, count uint64)
	AddShaderFunctionInfoCount(info unsafe.Pointer, count uint64)
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

func NewGTMutableShaderProfilerStreamDataWithCoder(coder objectivec.IObject) GTMutableShaderProfilerStreamData {
	instance := getGTMutableShaderProfilerStreamDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTMutableShaderProfilerStreamDataFromID(rv)
}

func NewGTMutableShaderProfilerStreamDataWithNewFileFormatV2Support(v2Support bool) GTMutableShaderProfilerStreamData {
	instance := getGTMutableShaderProfilerStreamDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNewFileFormatV2Support:"), v2Support)
	return GTMutableShaderProfilerStreamDataFromID(rv)
}

func NewGTMutableShaderProfilerStreamDataWithPreSiBundle(bundle objectivec.IObject) GTMutableShaderProfilerStreamData {
	instance := getGTMutableShaderProfilerStreamDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPreSiBundle:"), bundle)
	return GTMutableShaderProfilerStreamDataFromID(rv)
}

func (g GTMutableShaderProfilerStreamData) _commonInit() {
	objc.Send[objc.ID](g.ID, objc.Sel("_commonInit"))
}

// CommonInit is an exported wrapper for the private method _commonInit.
func (g GTMutableShaderProfilerStreamData) CommonInit() error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_commonInit")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_commonInit"}
		return err
	}
	g._commonInit()
	return nil
}

// CanCommonInit reports whether the receiver responds to the private selector _commonInit.
func (g GTMutableShaderProfilerStreamData) CanCommonInit() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_commonInit"))
}
func (g GTMutableShaderProfilerStreamData) _copyForAddAPSDataPrefix(aPSData objectivec.IObject, prefix objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_copyForAddAPSData:prefix:"), aPSData, prefix)
	return objectivec.Object{ID: rv}
}

// CopyForAddAPSDataPrefix is an exported wrapper for the private method _copyForAddAPSDataPrefix.
func (g GTMutableShaderProfilerStreamData) CopyForAddAPSDataPrefix(aPSData objectivec.IObject, prefix objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_copyForAddAPSData:prefix:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_copyForAddAPSData:prefix:"}
		return nil, err
	}
	return g._copyForAddAPSDataPrefix(aPSData, prefix), nil
}

// CanCopyForAddAPSDataPrefix reports whether the receiver responds to the private selector _copyForAddAPSData:prefix:.
func (g GTMutableShaderProfilerStreamData) CanCopyForAddAPSDataPrefix() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_copyForAddAPSData:prefix:"))
}
func (g GTMutableShaderProfilerStreamData) AddAPSCounterData(data objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("addAPSCounterData:"), data)
	return rv
}
func (g GTMutableShaderProfilerStreamData) AddAPSData(aPSData objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("addAPSData:"), aPSData)
	return rv
}
func (g GTMutableShaderProfilerStreamData) AddAPSTimelineData(data objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("addAPSTimelineData:"), data)
	return rv
}
func (g GTMutableShaderProfilerStreamData) AddBatchIdFilteredCounterData(data objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("addBatchIdFilteredCounterData:"), data)
	return rv
}
func (g GTMutableShaderProfilerStreamData) AddCommandBuffersCount(buffers unsafe.Pointer, count uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("addCommandBuffers:count:"), objc.CArray(buffers), count)
}
func (g GTMutableShaderProfilerStreamData) AddEncodersCount(encoders unsafe.Pointer, count uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("addEncoders:count:"), objc.CArray(encoders), count)
}
func (g GTMutableShaderProfilerStreamData) AddGPUCommandsCount(gPUCommands unsafe.Pointer, count uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("addGPUCommands:count:"), objc.CArray(gPUCommands), count)
}
func (g GTMutableShaderProfilerStreamData) AddGPUTimelineData(data objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("addGPUTimelineData:"), data)
	return rv
}
func (g GTMutableShaderProfilerStreamData) AddPipelinePerformanceStatisticsData(data objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("addPipelinePerformanceStatisticsData:"), data)
}
func (g GTMutableShaderProfilerStreamData) AddPipelineStatesCount(states unsafe.Pointer, count uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("addPipelineStates:count:"), objc.CArray(states), count)
}
func (g GTMutableShaderProfilerStreamData) AddShaderFunctionInfoCount(info unsafe.Pointer, count uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("addShaderFunctionInfo:count:"), objc.CArray(info), count)
}
func (g GTMutableShaderProfilerStreamData) AddShaderProfilerData(data objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("addShaderProfilerData:"), data)
	return rv
}
func (g GTMutableShaderProfilerStreamData) AddString(string_ objectivec.IObject) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("addString:"), string_)
	return rv
}
func (g GTMutableShaderProfilerStreamData) RemoveAPSCounterData() {
	objc.Send[objc.ID](g.ID, objc.Sel("removeAPSCounterData"))
}
func (g GTMutableShaderProfilerStreamData) RemoveAPSData() {
	objc.Send[objc.ID](g.ID, objc.Sel("removeAPSData"))
}
func (g GTMutableShaderProfilerStreamData) RemoveAPSTimelineData() {
	objc.Send[objc.ID](g.ID, objc.Sel("removeAPSTimelineData"))
}
func (g GTMutableShaderProfilerStreamData) SetDataSourceHasUnusedResourcesCaptureRange(resources bool, range_ foundation.NSRange) {
	objc.Send[objc.ID](g.ID, objc.Sel("setDataSourceHasUnusedResources:captureRange:"), resources, range_)
}
func (g GTMutableShaderProfilerStreamData) SetNumBlitCalls(calls uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("setNumBlitCalls:"), calls)
}

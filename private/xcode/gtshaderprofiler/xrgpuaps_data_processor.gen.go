// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [XRGPUAPSDataProcessor] class.
var (
	_XRGPUAPSDataProcessorClass     XRGPUAPSDataProcessorClass
	_XRGPUAPSDataProcessorClassOnce sync.Once
)

func getXRGPUAPSDataProcessorClass() XRGPUAPSDataProcessorClass {
	_XRGPUAPSDataProcessorClassOnce.Do(func() {
		_XRGPUAPSDataProcessorClass = XRGPUAPSDataProcessorClass{class: objc.GetClass("XRGPUAPSDataProcessor")}
	})
	return _XRGPUAPSDataProcessorClass
}

// GetXRGPUAPSDataProcessorClass returns the class object for XRGPUAPSDataProcessor.
func GetXRGPUAPSDataProcessorClass() XRGPUAPSDataProcessorClass {
	return getXRGPUAPSDataProcessorClass()
}

type XRGPUAPSDataProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (xc XRGPUAPSDataProcessorClass) Class() objc.Class {
	return xc.class
}

// Alloc allocates memory for a new instance of the class.
func (xc XRGPUAPSDataProcessorClass) Alloc() XRGPUAPSDataProcessor {
	rv := objc.Send[XRGPUAPSDataProcessor](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [XRGPUAPSDataProcessor._loadCSVShaderMap]
//   - [XRGPUAPSDataProcessor._loadJSONShaderMap]
//   - [XRGPUAPSDataProcessor._loadOSLogShaderMap]
//   - [XRGPUAPSDataProcessor._supportsDPBRMANCounter]
//   - [XRGPUAPSDataProcessor.AcceleratorID]
//   - [XRGPUAPSDataProcessor.AddBufferAtRDESourceIndexRdeBufferIndexBufferLength]
//   - [XRGPUAPSDataProcessor.AddBufferAtUSCIndexBuffer]
//   - [XRGPUAPSDataProcessor.AddBufferAtUSCIndexBufferLength]
//   - [XRGPUAPSDataProcessor.AddShaderMapDataVariant]
//   - [XRGPUAPSDataProcessor.AggregateAPSCounters]
//   - [XRGPUAPSDataProcessor.AggregateRDESourceBuffer]
//   - [XRGPUAPSDataProcessor.AggregateShaders]
//   - [XRGPUAPSDataProcessor.AgxpsGPU]
//   - [XRGPUAPSDataProcessor.AllShaders]
//   - [XRGPUAPSDataProcessor.ApsDerivedCounters]
//   - [XRGPUAPSDataProcessor.ApsRawCounterNames]
//   - [XRGPUAPSDataProcessor.Config]
//   - [XRGPUAPSDataProcessor.ConvertTimestampToTraceBeginMachAbsoluteBegin]
//   - [XRGPUAPSDataProcessor.CountPeriod]
//   - [XRGPUAPSDataProcessor.CounterConfigForGRCCounterSet]
//   - [XRGPUAPSDataProcessor.CounterTypeFromGroupNameCounterName]
//   - [XRGPUAPSDataProcessor.Delegate]
//   - [XRGPUAPSDataProcessor.SetDelegate]
//   - [XRGPUAPSDataProcessor.DeriveAPSCountersNumCoresCounterSet]
//   - [XRGPUAPSDataProcessor.DeriveRDECountersCounterIndexesRawCounterIdsDerivedCounterIdsDeltaSecondsIndex]
//   - [XRGPUAPSDataProcessor.DeriveRDESourceBuffer]
//   - [XRGPUAPSDataProcessor.EnableUSCEnable]
//   - [XRGPUAPSDataProcessor.EncodeRawData]
//   - [XRGPUAPSDataProcessor.EnumerateShaders]
//   - [XRGPUAPSDataProcessor.FirstAPSTimestamp]
//   - [XRGPUAPSDataProcessor.FirstRDETimestamp]
//   - [XRGPUAPSDataProcessor.GetAPSDerivedCounterDataTimestampsSampleCountCounterIndexCount]
//   - [XRGPUAPSDataProcessor.GetAPSDerivedCounterDataAtUSCIndexBufferTimestampsSampleCountCounterIndexCount]
//   - [XRGPUAPSDataProcessor.GetAPSRawCounterDataTimestampsSampleCountCounterIndexCount]
//   - [XRGPUAPSDataProcessor.GetAPSRawCounterDataAtUSCIndexBufferTimestampsSampleCountCounterIndexCount]
//   - [XRGPUAPSDataProcessor.GetBufferAtRDESourceIndexRdeBufferIndexBufferLength]
//   - [XRGPUAPSDataProcessor.GetBufferAtUSCIndexBufferLength]
//   - [XRGPUAPSDataProcessor.GetKickTraceIdsGlobalTraceIdsCount]
//   - [XRGPUAPSDataProcessor.GetKickTraceIdsAtUSCIndexKickTraceIdsGlobalTraceIdsCount]
//   - [XRGPUAPSDataProcessor.GetProfileDataAtUSCIndex]
//   - [XRGPUAPSDataProcessor.GetRDEDerivedCounterAtSourceIndexBufferTimestampsSampleCountCounterIndexCount]
//   - [XRGPUAPSDataProcessor.GetRDEDerivedCounterAtSourceIndexRdeBufferIndexBufferTimestampsSampleCountCounterIndexCount]
//   - [XRGPUAPSDataProcessor.GetRDERawCounterAtSourceIndexBufferTimestampsSampleCountCounterIndexCount]
//   - [XRGPUAPSDataProcessor.GetRDERawCounterAtSourceIndexRdeBufferIndexBufferTimestampsSampleCountCounterIndexCount]
//   - [XRGPUAPSDataProcessor.GetShaderIdsKickTraceIdShaderIdsShaderPCsDataMastersCount]
//   - [XRGPUAPSDataProcessor.GetShaderIdsShaderIdsShaderPCsDataMastersCount]
//   - [XRGPUAPSDataProcessor.GetShaderIdsAtUSCIndexKickTraceIdShaderIdsShaderPCsDataMastersCount]
//   - [XRGPUAPSDataProcessor.GetShaderIntervalsKickTraceIdShaderIdIntervalsCount]
//   - [XRGPUAPSDataProcessor.GetShaderIntervalsShaderIdIntervalsCount]
//   - [XRGPUAPSDataProcessor.GetShaderIntervalsAtUSCIndexKickTraceIdShaderIdIntervalsCount]
//   - [XRGPUAPSDataProcessor.IsUSCEnabled]
//   - [XRGPUAPSDataProcessor.IsValidUSC]
//   - [XRGPUAPSDataProcessor.LastAPSTimestamp]
//   - [XRGPUAPSDataProcessor.LastRDETimestamp]
//   - [XRGPUAPSDataProcessor.LoadAPSCountersCounterSet]
//   - [XRGPUAPSDataProcessor.LoadCounterGraphConfig]
//   - [XRGPUAPSDataProcessor.LoadCounters]
//   - [XRGPUAPSDataProcessor.LoadRDECounters]
//   - [XRGPUAPSDataProcessor.LoadShaders]
//   - [XRGPUAPSDataProcessor.LoadShadersUscIndex]
//   - [XRGPUAPSDataProcessor.NumAPSDerivedCounters]
//   - [XRGPUAPSDataProcessor.NumAPSRawCounters]
//   - [XRGPUAPSDataProcessor.NumRDEBuffersAtSourceIndex]
//   - [XRGPUAPSDataProcessor.NumRDESources]
//   - [XRGPUAPSDataProcessor.NumUSCs]
//   - [XRGPUAPSDataProcessor.NumValidUSCs]
//   - [XRGPUAPSDataProcessor.ParseData]
//   - [XRGPUAPSDataProcessor.ParseDataLengthUscIndex]
//   - [XRGPUAPSDataProcessor.PidFromKickTraceId]
//   - [XRGPUAPSDataProcessor.PulsePeriod]
//   - [XRGPUAPSDataProcessor.RdeDerivedCountersAtSourceIndex]
//   - [XRGPUAPSDataProcessor.RdeRawCounterNamesAtSourceIndex]
//   - [XRGPUAPSDataProcessor.SampleInterval]
//   - [XRGPUAPSDataProcessor.ShaderFromId]
//   - [XRGPUAPSDataProcessor.ShaderFromPCPidTimeUscIndex]
//   - [XRGPUAPSDataProcessor.TimestampRefToNsUscTimeIndexUscIndex]
//   - [XRGPUAPSDataProcessor.TimestampRefsToNsCountResultUscIndex]
//   - [XRGPUAPSDataProcessor.InitWithGPUGenerationVariantRevConfigOptions]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor
type XRGPUAPSDataProcessor struct {
	objectivec.Object
}

// XRGPUAPSDataProcessorFromID constructs a [XRGPUAPSDataProcessor] from an objc.ID.
func XRGPUAPSDataProcessorFromID(id objc.ID) XRGPUAPSDataProcessor {
	return XRGPUAPSDataProcessor{objectivec.Object{ID: id}}
}

// Ensure XRGPUAPSDataProcessor implements IXRGPUAPSDataProcessor.
var _ IXRGPUAPSDataProcessor = XRGPUAPSDataProcessor{}

// An interface definition for the [XRGPUAPSDataProcessor] class.
//
// # Methods
//
//   - [IXRGPUAPSDataProcessor._loadCSVShaderMap]
//   - [IXRGPUAPSDataProcessor._loadJSONShaderMap]
//   - [IXRGPUAPSDataProcessor._loadOSLogShaderMap]
//   - [IXRGPUAPSDataProcessor._supportsDPBRMANCounter]
//   - [IXRGPUAPSDataProcessor.AcceleratorID]
//   - [IXRGPUAPSDataProcessor.AddBufferAtRDESourceIndexRdeBufferIndexBufferLength]
//   - [IXRGPUAPSDataProcessor.AddBufferAtUSCIndexBuffer]
//   - [IXRGPUAPSDataProcessor.AddBufferAtUSCIndexBufferLength]
//   - [IXRGPUAPSDataProcessor.AddShaderMapDataVariant]
//   - [IXRGPUAPSDataProcessor.AggregateAPSCounters]
//   - [IXRGPUAPSDataProcessor.AggregateRDESourceBuffer]
//   - [IXRGPUAPSDataProcessor.AggregateShaders]
//   - [IXRGPUAPSDataProcessor.AgxpsGPU]
//   - [IXRGPUAPSDataProcessor.AllShaders]
//   - [IXRGPUAPSDataProcessor.ApsDerivedCounters]
//   - [IXRGPUAPSDataProcessor.ApsRawCounterNames]
//   - [IXRGPUAPSDataProcessor.Config]
//   - [IXRGPUAPSDataProcessor.ConvertTimestampToTraceBeginMachAbsoluteBegin]
//   - [IXRGPUAPSDataProcessor.CountPeriod]
//   - [IXRGPUAPSDataProcessor.CounterConfigForGRCCounterSet]
//   - [IXRGPUAPSDataProcessor.CounterTypeFromGroupNameCounterName]
//   - [IXRGPUAPSDataProcessor.Delegate]
//   - [IXRGPUAPSDataProcessor.SetDelegate]
//   - [IXRGPUAPSDataProcessor.DeriveAPSCountersNumCoresCounterSet]
//   - [IXRGPUAPSDataProcessor.DeriveRDECountersCounterIndexesRawCounterIdsDerivedCounterIdsDeltaSecondsIndex]
//   - [IXRGPUAPSDataProcessor.DeriveRDESourceBuffer]
//   - [IXRGPUAPSDataProcessor.EnableUSCEnable]
//   - [IXRGPUAPSDataProcessor.EncodeRawData]
//   - [IXRGPUAPSDataProcessor.EnumerateShaders]
//   - [IXRGPUAPSDataProcessor.FirstAPSTimestamp]
//   - [IXRGPUAPSDataProcessor.FirstRDETimestamp]
//   - [IXRGPUAPSDataProcessor.GetAPSDerivedCounterDataTimestampsSampleCountCounterIndexCount]
//   - [IXRGPUAPSDataProcessor.GetAPSDerivedCounterDataAtUSCIndexBufferTimestampsSampleCountCounterIndexCount]
//   - [IXRGPUAPSDataProcessor.GetAPSRawCounterDataTimestampsSampleCountCounterIndexCount]
//   - [IXRGPUAPSDataProcessor.GetAPSRawCounterDataAtUSCIndexBufferTimestampsSampleCountCounterIndexCount]
//   - [IXRGPUAPSDataProcessor.GetBufferAtRDESourceIndexRdeBufferIndexBufferLength]
//   - [IXRGPUAPSDataProcessor.GetBufferAtUSCIndexBufferLength]
//   - [IXRGPUAPSDataProcessor.GetKickTraceIdsGlobalTraceIdsCount]
//   - [IXRGPUAPSDataProcessor.GetKickTraceIdsAtUSCIndexKickTraceIdsGlobalTraceIdsCount]
//   - [IXRGPUAPSDataProcessor.GetProfileDataAtUSCIndex]
//   - [IXRGPUAPSDataProcessor.GetRDEDerivedCounterAtSourceIndexBufferTimestampsSampleCountCounterIndexCount]
//   - [IXRGPUAPSDataProcessor.GetRDEDerivedCounterAtSourceIndexRdeBufferIndexBufferTimestampsSampleCountCounterIndexCount]
//   - [IXRGPUAPSDataProcessor.GetRDERawCounterAtSourceIndexBufferTimestampsSampleCountCounterIndexCount]
//   - [IXRGPUAPSDataProcessor.GetRDERawCounterAtSourceIndexRdeBufferIndexBufferTimestampsSampleCountCounterIndexCount]
//   - [IXRGPUAPSDataProcessor.GetShaderIdsKickTraceIdShaderIdsShaderPCsDataMastersCount]
//   - [IXRGPUAPSDataProcessor.GetShaderIdsShaderIdsShaderPCsDataMastersCount]
//   - [IXRGPUAPSDataProcessor.GetShaderIdsAtUSCIndexKickTraceIdShaderIdsShaderPCsDataMastersCount]
//   - [IXRGPUAPSDataProcessor.GetShaderIntervalsKickTraceIdShaderIdIntervalsCount]
//   - [IXRGPUAPSDataProcessor.GetShaderIntervalsShaderIdIntervalsCount]
//   - [IXRGPUAPSDataProcessor.GetShaderIntervalsAtUSCIndexKickTraceIdShaderIdIntervalsCount]
//   - [IXRGPUAPSDataProcessor.IsUSCEnabled]
//   - [IXRGPUAPSDataProcessor.IsValidUSC]
//   - [IXRGPUAPSDataProcessor.LastAPSTimestamp]
//   - [IXRGPUAPSDataProcessor.LastRDETimestamp]
//   - [IXRGPUAPSDataProcessor.LoadAPSCountersCounterSet]
//   - [IXRGPUAPSDataProcessor.LoadCounterGraphConfig]
//   - [IXRGPUAPSDataProcessor.LoadCounters]
//   - [IXRGPUAPSDataProcessor.LoadRDECounters]
//   - [IXRGPUAPSDataProcessor.LoadShaders]
//   - [IXRGPUAPSDataProcessor.LoadShadersUscIndex]
//   - [IXRGPUAPSDataProcessor.NumAPSDerivedCounters]
//   - [IXRGPUAPSDataProcessor.NumAPSRawCounters]
//   - [IXRGPUAPSDataProcessor.NumRDEBuffersAtSourceIndex]
//   - [IXRGPUAPSDataProcessor.NumRDESources]
//   - [IXRGPUAPSDataProcessor.NumUSCs]
//   - [IXRGPUAPSDataProcessor.NumValidUSCs]
//   - [IXRGPUAPSDataProcessor.ParseData]
//   - [IXRGPUAPSDataProcessor.ParseDataLengthUscIndex]
//   - [IXRGPUAPSDataProcessor.PidFromKickTraceId]
//   - [IXRGPUAPSDataProcessor.PulsePeriod]
//   - [IXRGPUAPSDataProcessor.RdeDerivedCountersAtSourceIndex]
//   - [IXRGPUAPSDataProcessor.RdeRawCounterNamesAtSourceIndex]
//   - [IXRGPUAPSDataProcessor.SampleInterval]
//   - [IXRGPUAPSDataProcessor.ShaderFromId]
//   - [IXRGPUAPSDataProcessor.ShaderFromPCPidTimeUscIndex]
//   - [IXRGPUAPSDataProcessor.TimestampRefToNsUscTimeIndexUscIndex]
//   - [IXRGPUAPSDataProcessor.TimestampRefsToNsCountResultUscIndex]
//   - [IXRGPUAPSDataProcessor.InitWithGPUGenerationVariantRevConfigOptions]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor
type IXRGPUAPSDataProcessor interface {
	objectivec.IObject

	// Topic: Methods

	_loadCSVShaderMap(map_ objectivec.IObject)
	_loadJSONShaderMap(map_ objectivec.IObject)
	_loadOSLogShaderMap(map_ objectivec.IObject)
	_supportsDPBRMANCounter() bool
	AcceleratorID() uint64
	AddBufferAtRDESourceIndexRdeBufferIndexBufferLength(index uint32, index2 uint32, buffer string, length uint64)
	AddBufferAtUSCIndexBuffer(uSCIndex uint32, buffer objectivec.IObject)
	AddBufferAtUSCIndexBufferLength(uSCIndex uint32, buffer string, length uint64)
	AddShaderMapDataVariant(data objectivec.IObject, variant int)
	AggregateAPSCounters(aPSCounters uint64) bool
	AggregateRDESourceBuffer(buffer unsafe.Pointer) bool
	AggregateShaders()
	AgxpsGPU() uint64
	AllShaders() foundation.INSArray
	ApsDerivedCounters() foundation.INSArray
	ApsRawCounterNames() foundation.INSArray
	Config() foundation.INSDictionary
	ConvertTimestampToTraceBeginMachAbsoluteBegin(begin uint64, begin2 uint64)
	CountPeriod() uint64
	CounterConfigForGRCCounterSet(grc bool, set uint64) objectivec.IObject
	CounterTypeFromGroupNameCounterName(name string, name2 string) uint32
	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)
	DeriveAPSCountersNumCoresCounterSet(aPSCounters unsafe.Pointer, cores uint32, set uint64) bool
	DeriveRDECountersCounterIndexesRawCounterIdsDerivedCounterIdsDeltaSecondsIndex(rDECounters unsafe.Pointer, indexes unsafe.Pointer, ids unsafe.Pointer, ids2 unsafe.Pointer, index uint64) bool
	DeriveRDESourceBuffer(buffer unsafe.Pointer) bool
	EnableUSCEnable(usc uint32, enable bool)
	EncodeRawData() objectivec.IObject
	EnumerateShaders(shaders VoidHandler)
	FirstAPSTimestamp() uint64
	FirstRDETimestamp() uint64
	GetAPSDerivedCounterDataTimestampsSampleCountCounterIndexCount(data []float64, index uint32, count2 uint32) (uint64, uint64, bool)
	GetAPSDerivedCounterDataAtUSCIndexBufferTimestampsSampleCountCounterIndexCount(uSCIndex uint32, buffer []float64, index uint32, count2 uint32) (uint64, uint64, bool)
	GetAPSRawCounterDataTimestampsSampleCountCounterIndexCount(index uint32, count2 uint32) (uint64, uint64, uint64, bool)
	GetAPSRawCounterDataAtUSCIndexBufferTimestampsSampleCountCounterIndexCount(uSCIndex uint32, index uint32, count2 uint32) (uint64, uint64, uint64, bool)
	GetBufferAtRDESourceIndexRdeBufferIndexBufferLength(index uint32, index2 uint32, buffer string) (uint64, bool)
	GetBufferAtUSCIndexBufferLength(uSCIndex uint32, buffer string) (uint64, bool)
	GetKickTraceIdsGlobalTraceIdsCount(ids *uint32, ids2 *uint32) (uint64, bool)
	GetKickTraceIdsAtUSCIndexKickTraceIdsGlobalTraceIdsCount(uSCIndex uint32, ids *uint32, ids2 *uint32) (uint64, bool)
	GetProfileDataAtUSCIndex(uSCIndex uint64) uint64
	GetRDEDerivedCounterAtSourceIndexBufferTimestampsSampleCountCounterIndexCount(index uint32, buffer []float64, index2 uint32, count2 uint32) (uint64, uint64, bool)
	GetRDEDerivedCounterAtSourceIndexRdeBufferIndexBufferTimestampsSampleCountCounterIndexCount(index uint32, index2 uint32, buffer []float64, index3 uint32, count2 uint32) (uint64, uint64, bool)
	GetRDERawCounterAtSourceIndexBufferTimestampsSampleCountCounterIndexCount(index uint32, index2 uint32, count2 uint32) (uint64, uint64, uint64, bool)
	GetRDERawCounterAtSourceIndexRdeBufferIndexBufferTimestampsSampleCountCounterIndexCount(index uint32, index2 uint32, index3 uint32, count2 uint32) (uint64, uint64, uint64, bool)
	GetShaderIdsKickTraceIdShaderIdsShaderPCsDataMastersCount(ids unsafe.Pointer, id uint32, ids2 *uint32, masters *uint32) (uint64, uint64, bool)
	GetShaderIdsShaderIdsShaderPCsDataMastersCount(ids uint32, ids2 *uint32, masters *uint32) (uint64, uint64, bool)
	GetShaderIdsAtUSCIndexKickTraceIdShaderIdsShaderPCsDataMastersCount(uSCIndex uint32, id uint32, ids *uint32, masters *uint32) (uint64, uint64, bool)
	GetShaderIntervalsKickTraceIdShaderIdIntervalsCount(intervals unsafe.Pointer, id uint32, id2 uint32, intervals2 *uintptr) (uint64, bool)
	GetShaderIntervalsShaderIdIntervalsCount(intervals uint32, id uint32, intervals2 *uintptr) (uint64, bool)
	GetShaderIntervalsAtUSCIndexKickTraceIdShaderIdIntervalsCount(uSCIndex uint32, id uint32, id2 uint32, intervals *uintptr) (uint64, bool)
	IsUSCEnabled(uSCEnabled uint32) bool
	IsValidUSC(usc uint32) bool
	LastAPSTimestamp() uint64
	LastRDETimestamp() uint64
	LoadAPSCountersCounterSet(aPSCounters unsafe.Pointer, set uint64) bool
	LoadCounterGraphConfig() objectivec.IObject
	LoadCounters(counters uint64) bool
	LoadRDECounters(rDECounters unsafe.Pointer) bool
	LoadShaders() bool
	LoadShadersUscIndex(shaders unsafe.Pointer, index uint32) bool
	NumAPSDerivedCounters() uint32
	NumAPSRawCounters() uint32
	NumRDEBuffersAtSourceIndex(index uint32) uint32
	NumRDESources() uint32
	NumUSCs() uint32
	NumValidUSCs() uint32
	ParseData() bool
	ParseDataLengthUscIndex(data string, length uint64, index uint32) bool
	PidFromKickTraceId(id uint32) int
	PulsePeriod() uint64
	RdeDerivedCountersAtSourceIndex(index uint32) objectivec.IObject
	RdeRawCounterNamesAtSourceIndex(index uint32) objectivec.IObject
	SampleInterval() uint64
	ShaderFromId(id uint32) objectivec.IObject
	ShaderFromPCPidTimeUscIndex(pc uint64, pid int, time uint64, index uint32) objectivec.IObject
	TimestampRefToNsUscTimeIndexUscIndex(ns uint64, index uint64, index2 uint32) uint64
	TimestampRefsToNsCountResultUscIndex(ns unsafe.Pointer, count uint64, result unsafe.Pointer, index uint32)
	InitWithGPUGenerationVariantRevConfigOptions(gPUGeneration uint32, variant uint32, rev uint32, config objectivec.IObject, options uint32) XRGPUAPSDataProcessor
}

// Init initializes the instance.
func (x XRGPUAPSDataProcessor) Init() XRGPUAPSDataProcessor {
	rv := objc.Send[XRGPUAPSDataProcessor](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x XRGPUAPSDataProcessor) Autorelease() XRGPUAPSDataProcessor {
	rv := objc.Send[XRGPUAPSDataProcessor](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewXRGPUAPSDataProcessor creates a new XRGPUAPSDataProcessor instance.
func NewXRGPUAPSDataProcessor() XRGPUAPSDataProcessor {
	class := getXRGPUAPSDataProcessorClass()
	rv := objc.Send[XRGPUAPSDataProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/initWithGPUGeneration:variant:rev:config:options:
func NewXRGPUAPSDataProcessorWithGPUGenerationVariantRevConfigOptions(gPUGeneration uint32, variant uint32, rev uint32, config objectivec.IObject, options uint32) XRGPUAPSDataProcessor {
	instance := getXRGPUAPSDataProcessorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGPUGeneration:variant:rev:config:options:"), gPUGeneration, variant, rev, config, options)
	return XRGPUAPSDataProcessorFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/_loadCSVShaderMap:
func (x XRGPUAPSDataProcessor) _loadCSVShaderMap(map_ objectivec.IObject) {
	objc.Send[objc.ID](x.ID, objc.Sel("_loadCSVShaderMap:"), map_)
}

// LoadCSVShaderMap is an exported wrapper for the private method _loadCSVShaderMap.
func (x XRGPUAPSDataProcessor) LoadCSVShaderMap(map_ objectivec.IObject) {
	x._loadCSVShaderMap(map_)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/_loadJSONShaderMap:
func (x XRGPUAPSDataProcessor) _loadJSONShaderMap(map_ objectivec.IObject) {
	objc.Send[objc.ID](x.ID, objc.Sel("_loadJSONShaderMap:"), map_)
}

// LoadJSONShaderMap is an exported wrapper for the private method _loadJSONShaderMap.
func (x XRGPUAPSDataProcessor) LoadJSONShaderMap(map_ objectivec.IObject) {
	x._loadJSONShaderMap(map_)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/_loadOSLogShaderMap:
func (x XRGPUAPSDataProcessor) _loadOSLogShaderMap(map_ objectivec.IObject) {
	objc.Send[objc.ID](x.ID, objc.Sel("_loadOSLogShaderMap:"), map_)
}

// LoadOSLogShaderMap is an exported wrapper for the private method _loadOSLogShaderMap.
func (x XRGPUAPSDataProcessor) LoadOSLogShaderMap(map_ objectivec.IObject) {
	x._loadOSLogShaderMap(map_)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/_supportsDPBRMANCounter
func (x XRGPUAPSDataProcessor) _supportsDPBRMANCounter() bool {
	rv := objc.Send[bool](x.ID, objc.Sel("_supportsDPBRMANCounter"))
	return rv
}

// SupportsDPBRMANCounter is an exported wrapper for the private method _supportsDPBRMANCounter.
func (x XRGPUAPSDataProcessor) SupportsDPBRMANCounter() bool {
	return x._supportsDPBRMANCounter()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/addBufferAtRDESourceIndex:rdeBufferIndex:buffer:length:
func (x XRGPUAPSDataProcessor) AddBufferAtRDESourceIndexRdeBufferIndexBufferLength(index uint32, index2 uint32, buffer string, length uint64) {
	objc.Send[objc.ID](x.ID, objc.Sel("addBufferAtRDESourceIndex:rdeBufferIndex:buffer:length:"), index, index2, unsafe.Pointer(unsafe.StringData(buffer+"\x00")), length)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/addBufferAtUSCIndex:buffer:
func (x XRGPUAPSDataProcessor) AddBufferAtUSCIndexBuffer(uSCIndex uint32, buffer objectivec.IObject) {
	objc.Send[objc.ID](x.ID, objc.Sel("addBufferAtUSCIndex:buffer:"), uSCIndex, buffer)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/addBufferAtUSCIndex:buffer:length:
func (x XRGPUAPSDataProcessor) AddBufferAtUSCIndexBufferLength(uSCIndex uint32, buffer string, length uint64) {
	objc.Send[objc.ID](x.ID, objc.Sel("addBufferAtUSCIndex:buffer:length:"), uSCIndex, unsafe.Pointer(unsafe.StringData(buffer+"\x00")), length)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/addShaderMapData:variant:
func (x XRGPUAPSDataProcessor) AddShaderMapDataVariant(data objectivec.IObject, variant int) {
	objc.Send[objc.ID](x.ID, objc.Sel("addShaderMapData:variant:"), data, variant)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/aggregateAPSCounters:
func (x XRGPUAPSDataProcessor) AggregateAPSCounters(aPSCounters uint64) bool {
	rv := objc.Send[bool](x.ID, objc.Sel("aggregateAPSCounters:"), aPSCounters)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/aggregateRDESourceBuffer:
func (x XRGPUAPSDataProcessor) AggregateRDESourceBuffer(buffer unsafe.Pointer) bool {
	rv := objc.Send[bool](x.ID, objc.Sel("aggregateRDESourceBuffer:"), buffer)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/aggregateShaders
func (x XRGPUAPSDataProcessor) AggregateShaders() {
	objc.Send[objc.ID](x.ID, objc.Sel("aggregateShaders"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/convertTimestampToTraceBegin:machAbsoluteBegin:
func (x XRGPUAPSDataProcessor) ConvertTimestampToTraceBeginMachAbsoluteBegin(begin uint64, begin2 uint64) {
	objc.Send[objc.ID](x.ID, objc.Sel("convertTimestampToTraceBegin:machAbsoluteBegin:"), begin, begin2)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/counterConfigForGRC:counterSet:
func (x XRGPUAPSDataProcessor) CounterConfigForGRCCounterSet(grc bool, set uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("counterConfigForGRC:counterSet:"), grc, set)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/counterTypeFromGroupName:counterName:
func (x XRGPUAPSDataProcessor) CounterTypeFromGroupNameCounterName(name string, name2 string) uint32 {
	rv := objc.Send[uint32](x.ID, objc.Sel("counterTypeFromGroupName:counterName:"), unsafe.Pointer(unsafe.StringData(name+"\x00")), unsafe.Pointer(unsafe.StringData(name2+"\x00")))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/deriveAPSCounters:numCores:counterSet:
func (x XRGPUAPSDataProcessor) DeriveAPSCountersNumCoresCounterSet(aPSCounters unsafe.Pointer, cores uint32, set uint64) bool {
	rv := objc.Send[bool](x.ID, objc.Sel("deriveAPSCounters:numCores:counterSet:"), aPSCounters, cores, set)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/deriveRDECounters:counterIndexes:rawCounterIds:derivedCounterIds:deltaSecondsIndex:
func (x XRGPUAPSDataProcessor) DeriveRDECountersCounterIndexesRawCounterIdsDerivedCounterIdsDeltaSecondsIndex(rDECounters unsafe.Pointer, indexes unsafe.Pointer, ids unsafe.Pointer, ids2 unsafe.Pointer, index uint64) bool {
	rv := objc.Send[bool](x.ID, objc.Sel("deriveRDECounters:counterIndexes:rawCounterIds:derivedCounterIds:deltaSecondsIndex:"), rDECounters, indexes, ids, ids2, index)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/deriveRDESourceBuffer:
func (x XRGPUAPSDataProcessor) DeriveRDESourceBuffer(buffer unsafe.Pointer) bool {
	rv := objc.Send[bool](x.ID, objc.Sel("deriveRDESourceBuffer:"), buffer)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/enableUSC:enable:
func (x XRGPUAPSDataProcessor) EnableUSCEnable(usc uint32, enable bool) {
	objc.Send[objc.ID](x.ID, objc.Sel("enableUSC:enable:"), usc, enable)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/encodeRawData
func (x XRGPUAPSDataProcessor) EncodeRawData() objectivec.IObject {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("encodeRawData"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/enumerateShaders:
func (x XRGPUAPSDataProcessor) EnumerateShaders(shaders VoidHandler) {
	_block0, _ := NewVoidBlock(shaders)
	objc.Send[objc.ID](x.ID, objc.Sel("enumerateShaders:"), _block0)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/getAPSDerivedCounterData:timestamps:sampleCount:counterIndex:count:
func (x XRGPUAPSDataProcessor) GetAPSDerivedCounterDataTimestampsSampleCountCounterIndexCount(data []float64, index uint32, count2 uint32) (uint64, uint64, bool) {
	var timestamps uint64
	var count uint64
	rv := objc.Send[bool](x.ID, objc.Sel("getAPSDerivedCounterData:timestamps:sampleCount:counterIndex:count:"), data, unsafe.Pointer(&timestamps), unsafe.Pointer(&count), index, count2)
	return timestamps, count, rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/getAPSDerivedCounterDataAtUSCIndex:buffer:timestamps:sampleCount:counterIndex:count:
func (x XRGPUAPSDataProcessor) GetAPSDerivedCounterDataAtUSCIndexBufferTimestampsSampleCountCounterIndexCount(uSCIndex uint32, buffer []float64, index uint32, count2 uint32) (uint64, uint64, bool) {
	var timestamps uint64
	var count uint64
	rv := objc.Send[bool](x.ID, objc.Sel("getAPSDerivedCounterDataAtUSCIndex:buffer:timestamps:sampleCount:counterIndex:count:"), uSCIndex, buffer, unsafe.Pointer(&timestamps), unsafe.Pointer(&count), index, count2)
	return timestamps, count, rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/getAPSRawCounterData:timestamps:sampleCount:counterIndex:count:
func (x XRGPUAPSDataProcessor) GetAPSRawCounterDataTimestampsSampleCountCounterIndexCount(index uint32, count2 uint32) (uint64, uint64, uint64, bool) {
	var data uint64
	var timestamps uint64
	var count uint64
	rv := objc.Send[bool](x.ID, objc.Sel("getAPSRawCounterData:timestamps:sampleCount:counterIndex:count:"), unsafe.Pointer(&data), unsafe.Pointer(&timestamps), unsafe.Pointer(&count), index, count2)
	return data, timestamps, count, rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/getAPSRawCounterDataAtUSCIndex:buffer:timestamps:sampleCount:counterIndex:count:
func (x XRGPUAPSDataProcessor) GetAPSRawCounterDataAtUSCIndexBufferTimestampsSampleCountCounterIndexCount(uSCIndex uint32, index uint32, count2 uint32) (uint64, uint64, uint64, bool) {
	var buffer uint64
	var timestamps uint64
	var count uint64
	rv := objc.Send[bool](x.ID, objc.Sel("getAPSRawCounterDataAtUSCIndex:buffer:timestamps:sampleCount:counterIndex:count:"), uSCIndex, unsafe.Pointer(&buffer), unsafe.Pointer(&timestamps), unsafe.Pointer(&count), index, count2)
	return buffer, timestamps, count, rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/getBufferAtRDESourceIndex:rdeBufferIndex:buffer:length:
func (x XRGPUAPSDataProcessor) GetBufferAtRDESourceIndexRdeBufferIndexBufferLength(index uint32, index2 uint32, buffer string) (uint64, bool) {
	var length uint64
	rv := objc.Send[bool](x.ID, objc.Sel("getBufferAtRDESourceIndex:rdeBufferIndex:buffer:length:"), index, index2, buffer, unsafe.Pointer(&length))
	return length, rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/getBufferAtUSCIndex:buffer:length:
func (x XRGPUAPSDataProcessor) GetBufferAtUSCIndexBufferLength(uSCIndex uint32, buffer string) (uint64, bool) {
	var length uint64
	rv := objc.Send[bool](x.ID, objc.Sel("getBufferAtUSCIndex:buffer:length:"), uSCIndex, buffer, unsafe.Pointer(&length))
	return length, rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/getKickTraceIds:globalTraceIds:count:
func (x XRGPUAPSDataProcessor) GetKickTraceIdsGlobalTraceIdsCount(ids *uint32, ids2 *uint32) (uint64, bool) {
	var count uint64
	rv := objc.Send[bool](x.ID, objc.Sel("getKickTraceIds:globalTraceIds:count:"), ids, ids2, unsafe.Pointer(&count))
	return count, rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/getKickTraceIdsAtUSCIndex:kickTraceIds:globalTraceIds:count:
func (x XRGPUAPSDataProcessor) GetKickTraceIdsAtUSCIndexKickTraceIdsGlobalTraceIdsCount(uSCIndex uint32, ids *uint32, ids2 *uint32) (uint64, bool) {
	var count uint64
	rv := objc.Send[bool](x.ID, objc.Sel("getKickTraceIdsAtUSCIndex:kickTraceIds:globalTraceIds:count:"), uSCIndex, ids, ids2, unsafe.Pointer(&count))
	return count, rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/getProfileDataAtUSCIndex:
func (x XRGPUAPSDataProcessor) GetProfileDataAtUSCIndex(uSCIndex uint64) uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("getProfileDataAtUSCIndex:"), uSCIndex)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/getRDEDerivedCounterAtSourceIndex:buffer:timestamps:sampleCount:counterIndex:count:
func (x XRGPUAPSDataProcessor) GetRDEDerivedCounterAtSourceIndexBufferTimestampsSampleCountCounterIndexCount(index uint32, buffer []float64, index2 uint32, count2 uint32) (uint64, uint64, bool) {
	var timestamps uint64
	var count uint64
	rv := objc.Send[bool](x.ID, objc.Sel("getRDEDerivedCounterAtSourceIndex:buffer:timestamps:sampleCount:counterIndex:count:"), index, buffer, unsafe.Pointer(&timestamps), unsafe.Pointer(&count), index2, count2)
	return timestamps, count, rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/getRDEDerivedCounterAtSourceIndex:rdeBufferIndex:buffer:timestamps:sampleCount:counterIndex:count:
func (x XRGPUAPSDataProcessor) GetRDEDerivedCounterAtSourceIndexRdeBufferIndexBufferTimestampsSampleCountCounterIndexCount(index uint32, index2 uint32, buffer []float64, index3 uint32, count2 uint32) (uint64, uint64, bool) {
	var timestamps uint64
	var count uint64
	rv := objc.Send[bool](x.ID, objc.Sel("getRDEDerivedCounterAtSourceIndex:rdeBufferIndex:buffer:timestamps:sampleCount:counterIndex:count:"), index, index2, buffer, unsafe.Pointer(&timestamps), unsafe.Pointer(&count), index3, count2)
	return timestamps, count, rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/getRDERawCounterAtSourceIndex:buffer:timestamps:sampleCount:counterIndex:count:
func (x XRGPUAPSDataProcessor) GetRDERawCounterAtSourceIndexBufferTimestampsSampleCountCounterIndexCount(index uint32, index2 uint32, count2 uint32) (uint64, uint64, uint64, bool) {
	var buffer uint64
	var timestamps uint64
	var count uint64
	rv := objc.Send[bool](x.ID, objc.Sel("getRDERawCounterAtSourceIndex:buffer:timestamps:sampleCount:counterIndex:count:"), index, unsafe.Pointer(&buffer), unsafe.Pointer(&timestamps), unsafe.Pointer(&count), index2, count2)
	return buffer, timestamps, count, rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/getRDERawCounterAtSourceIndex:rdeBufferIndex:buffer:timestamps:sampleCount:counterIndex:count:
func (x XRGPUAPSDataProcessor) GetRDERawCounterAtSourceIndexRdeBufferIndexBufferTimestampsSampleCountCounterIndexCount(index uint32, index2 uint32, index3 uint32, count2 uint32) (uint64, uint64, uint64, bool) {
	var buffer uint64
	var timestamps uint64
	var count uint64
	rv := objc.Send[bool](x.ID, objc.Sel("getRDERawCounterAtSourceIndex:rdeBufferIndex:buffer:timestamps:sampleCount:counterIndex:count:"), index, index2, unsafe.Pointer(&buffer), unsafe.Pointer(&timestamps), unsafe.Pointer(&count), index3, count2)
	return buffer, timestamps, count, rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/getShaderIds:kickTraceId:shaderIds:shaderPCs:dataMasters:count:
func (x XRGPUAPSDataProcessor) GetShaderIdsKickTraceIdShaderIdsShaderPCsDataMastersCount(ids unsafe.Pointer, id uint32, ids2 *uint32, masters *uint32) (uint64, uint64, bool) {
	var pCs uint64
	var count uint64
	rv := objc.Send[bool](x.ID, objc.Sel("getShaderIds:kickTraceId:shaderIds:shaderPCs:dataMasters:count:"), ids, id, ids2, unsafe.Pointer(&pCs), masters, unsafe.Pointer(&count))
	return pCs, count, rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/getShaderIds:shaderIds:shaderPCs:dataMasters:count:
func (x XRGPUAPSDataProcessor) GetShaderIdsShaderIdsShaderPCsDataMastersCount(ids uint32, ids2 *uint32, masters *uint32) (uint64, uint64, bool) {
	var pCs uint64
	var count uint64
	rv := objc.Send[bool](x.ID, objc.Sel("getShaderIds:shaderIds:shaderPCs:dataMasters:count:"), ids, ids2, unsafe.Pointer(&pCs), masters, unsafe.Pointer(&count))
	return pCs, count, rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/getShaderIdsAtUSCIndex:kickTraceId:shaderIds:shaderPCs:dataMasters:count:
func (x XRGPUAPSDataProcessor) GetShaderIdsAtUSCIndexKickTraceIdShaderIdsShaderPCsDataMastersCount(uSCIndex uint32, id uint32, ids *uint32, masters *uint32) (uint64, uint64, bool) {
	var pCs uint64
	var count uint64
	rv := objc.Send[bool](x.ID, objc.Sel("getShaderIdsAtUSCIndex:kickTraceId:shaderIds:shaderPCs:dataMasters:count:"), uSCIndex, id, ids, unsafe.Pointer(&pCs), masters, unsafe.Pointer(&count))
	return pCs, count, rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/getShaderIntervals:kickTraceId:shaderId:intervals:count:
func (x XRGPUAPSDataProcessor) GetShaderIntervalsKickTraceIdShaderIdIntervalsCount(intervals unsafe.Pointer, id uint32, id2 uint32, intervals2 *uintptr) (uint64, bool) {
	var count uint64
	rv := objc.Send[bool](x.ID, objc.Sel("getShaderIntervals:kickTraceId:shaderId:intervals:count:"), intervals, id, id2, intervals2, unsafe.Pointer(&count))
	return count, rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/getShaderIntervals:shaderId:intervals:count:
func (x XRGPUAPSDataProcessor) GetShaderIntervalsShaderIdIntervalsCount(intervals uint32, id uint32, intervals2 *uintptr) (uint64, bool) {
	var count uint64
	rv := objc.Send[bool](x.ID, objc.Sel("getShaderIntervals:shaderId:intervals:count:"), intervals, id, intervals2, unsafe.Pointer(&count))
	return count, rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/getShaderIntervalsAtUSCIndex:kickTraceId:shaderId:intervals:count:
func (x XRGPUAPSDataProcessor) GetShaderIntervalsAtUSCIndexKickTraceIdShaderIdIntervalsCount(uSCIndex uint32, id uint32, id2 uint32, intervals *uintptr) (uint64, bool) {
	var count uint64
	rv := objc.Send[bool](x.ID, objc.Sel("getShaderIntervalsAtUSCIndex:kickTraceId:shaderId:intervals:count:"), uSCIndex, id, id2, intervals, unsafe.Pointer(&count))
	return count, rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/isUSCEnabled:
func (x XRGPUAPSDataProcessor) IsUSCEnabled(uSCEnabled uint32) bool {
	rv := objc.Send[bool](x.ID, objc.Sel("isUSCEnabled:"), uSCEnabled)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/isValidUSC:
func (x XRGPUAPSDataProcessor) IsValidUSC(usc uint32) bool {
	rv := objc.Send[bool](x.ID, objc.Sel("isValidUSC:"), usc)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/loadAPSCounters:counterSet:
func (x XRGPUAPSDataProcessor) LoadAPSCountersCounterSet(aPSCounters unsafe.Pointer, set uint64) bool {
	rv := objc.Send[bool](x.ID, objc.Sel("loadAPSCounters:counterSet:"), aPSCounters, set)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/loadCounterGraphConfig
func (x XRGPUAPSDataProcessor) LoadCounterGraphConfig() objectivec.IObject {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("loadCounterGraphConfig"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/loadCounters:
func (x XRGPUAPSDataProcessor) LoadCounters(counters uint64) bool {
	rv := objc.Send[bool](x.ID, objc.Sel("loadCounters:"), counters)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/loadRDECounters:
func (x XRGPUAPSDataProcessor) LoadRDECounters(rDECounters unsafe.Pointer) bool {
	rv := objc.Send[bool](x.ID, objc.Sel("loadRDECounters:"), rDECounters)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/loadShaders
func (x XRGPUAPSDataProcessor) LoadShaders() bool {
	rv := objc.Send[bool](x.ID, objc.Sel("loadShaders"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/loadShaders:uscIndex:
func (x XRGPUAPSDataProcessor) LoadShadersUscIndex(shaders unsafe.Pointer, index uint32) bool {
	rv := objc.Send[bool](x.ID, objc.Sel("loadShaders:uscIndex:"), shaders, index)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/numRDEBuffersAtSourceIndex:
func (x XRGPUAPSDataProcessor) NumRDEBuffersAtSourceIndex(index uint32) uint32 {
	rv := objc.Send[uint32](x.ID, objc.Sel("numRDEBuffersAtSourceIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/parseData
func (x XRGPUAPSDataProcessor) ParseData() bool {
	rv := objc.Send[bool](x.ID, objc.Sel("parseData"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/parseData:length:uscIndex:
func (x XRGPUAPSDataProcessor) ParseDataLengthUscIndex(data string, length uint64, index uint32) bool {
	rv := objc.Send[bool](x.ID, objc.Sel("parseData:length:uscIndex:"), unsafe.Pointer(unsafe.StringData(data+"\x00")), length, index)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/pidFromKickTraceId:
func (x XRGPUAPSDataProcessor) PidFromKickTraceId(id uint32) int {
	rv := objc.Send[int](x.ID, objc.Sel("pidFromKickTraceId:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/rdeDerivedCountersAtSourceIndex:
func (x XRGPUAPSDataProcessor) RdeDerivedCountersAtSourceIndex(index uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("rdeDerivedCountersAtSourceIndex:"), index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/rdeRawCounterNamesAtSourceIndex:
func (x XRGPUAPSDataProcessor) RdeRawCounterNamesAtSourceIndex(index uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("rdeRawCounterNamesAtSourceIndex:"), index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/shaderFromId:
func (x XRGPUAPSDataProcessor) ShaderFromId(id uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("shaderFromId:"), id)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/shaderFromPC:pid:time:uscIndex:
func (x XRGPUAPSDataProcessor) ShaderFromPCPidTimeUscIndex(pc uint64, pid int, time uint64, index uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("shaderFromPC:pid:time:uscIndex:"), pc, pid, time, index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/timestampRefToNs:uscTimeIndex:uscIndex:
func (x XRGPUAPSDataProcessor) TimestampRefToNsUscTimeIndexUscIndex(ns uint64, index uint64, index2 uint32) uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("timestampRefToNs:uscTimeIndex:uscIndex:"), ns, index, index2)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/timestampRefsToNs:count:result:uscIndex:
func (x XRGPUAPSDataProcessor) TimestampRefsToNsCountResultUscIndex(ns unsafe.Pointer, count uint64, result unsafe.Pointer, index uint32) {
	objc.Send[objc.ID](x.ID, objc.Sel("timestampRefsToNs:count:result:uscIndex:"), objc.CArray(ns), count, result, index)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/initWithGPUGeneration:variant:rev:config:options:
func (x XRGPUAPSDataProcessor) InitWithGPUGenerationVariantRevConfigOptions(gPUGeneration uint32, variant uint32, rev uint32, config objectivec.IObject, options uint32) XRGPUAPSDataProcessor {
	rv := objc.Send[XRGPUAPSDataProcessor](x.ID, objc.Sel("initWithGPUGeneration:variant:rev:config:options:"), gPUGeneration, variant, rev, config, options)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/processorFromConfig:options:
func (_XRGPUAPSDataProcessorClass XRGPUAPSDataProcessorClass) ProcessorFromConfigOptions(config objectivec.IObject, options uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_XRGPUAPSDataProcessorClass.class), objc.Sel("processorFromConfig:options:"), config, options)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/processorFromDataContainer:options:
func (_XRGPUAPSDataProcessorClass XRGPUAPSDataProcessorClass) ProcessorFromDataContainerOptions(container objectivec.IObject, options uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_XRGPUAPSDataProcessorClass.class), objc.Sel("processorFromDataContainer:options:"), container, options)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/acceleratorID
func (x XRGPUAPSDataProcessor) AcceleratorID() uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("acceleratorID"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/agxpsGPU
func (x XRGPUAPSDataProcessor) AgxpsGPU() uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("agxpsGPU"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/allShaders
func (x XRGPUAPSDataProcessor) AllShaders() foundation.INSArray {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("allShaders"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/apsDerivedCounters
func (x XRGPUAPSDataProcessor) ApsDerivedCounters() foundation.INSArray {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("apsDerivedCounters"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/apsRawCounterNames
func (x XRGPUAPSDataProcessor) ApsRawCounterNames() foundation.INSArray {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("apsRawCounterNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/config
func (x XRGPUAPSDataProcessor) Config() foundation.INSDictionary {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("config"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/countPeriod
func (x XRGPUAPSDataProcessor) CountPeriod() uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("countPeriod"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/delegate
func (x XRGPUAPSDataProcessor) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (x XRGPUAPSDataProcessor) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](x.ID, objc.Sel("setDelegate:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/firstAPSTimestamp
func (x XRGPUAPSDataProcessor) FirstAPSTimestamp() uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("firstAPSTimestamp"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/firstRDETimestamp
func (x XRGPUAPSDataProcessor) FirstRDETimestamp() uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("firstRDETimestamp"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/lastAPSTimestamp
func (x XRGPUAPSDataProcessor) LastAPSTimestamp() uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("lastAPSTimestamp"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/lastRDETimestamp
func (x XRGPUAPSDataProcessor) LastRDETimestamp() uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("lastRDETimestamp"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/numAPSDerivedCounters
func (x XRGPUAPSDataProcessor) NumAPSDerivedCounters() uint32 {
	rv := objc.Send[uint32](x.ID, objc.Sel("numAPSDerivedCounters"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/numAPSRawCounters
func (x XRGPUAPSDataProcessor) NumAPSRawCounters() uint32 {
	rv := objc.Send[uint32](x.ID, objc.Sel("numAPSRawCounters"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/numRDESources
func (x XRGPUAPSDataProcessor) NumRDESources() uint32 {
	rv := objc.Send[uint32](x.ID, objc.Sel("numRDESources"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/numUSCs
func (x XRGPUAPSDataProcessor) NumUSCs() uint32 {
	rv := objc.Send[uint32](x.ID, objc.Sel("numUSCs"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/numValidUSCs
func (x XRGPUAPSDataProcessor) NumValidUSCs() uint32 {
	rv := objc.Send[uint32](x.ID, objc.Sel("numValidUSCs"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/pulsePeriod
func (x XRGPUAPSDataProcessor) PulsePeriod() uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("pulsePeriod"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataProcessor/sampleInterval
func (x XRGPUAPSDataProcessor) SampleInterval() uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("sampleInterval"))
	return rv
}

// EnumerateShadersSync is a synchronous wrapper around [XRGPUAPSDataProcessor.EnumerateShaders].
// It blocks until the completion handler fires or the context is cancelled.
func (x XRGPUAPSDataProcessor) EnumerateShadersSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	x.EnumerateShaders(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

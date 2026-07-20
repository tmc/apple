// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioNonOverlappingCounters] class.
var (
	_GTMioNonOverlappingCountersClass     GTMioNonOverlappingCountersClass
	_GTMioNonOverlappingCountersClassOnce sync.Once
)

func getGTMioNonOverlappingCountersClass() GTMioNonOverlappingCountersClass {
	_GTMioNonOverlappingCountersClassOnce.Do(func() {
		_GTMioNonOverlappingCountersClass = GTMioNonOverlappingCountersClass{class: objc.GetClass("GTMioNonOverlappingCounters")}
	})
	return _GTMioNonOverlappingCountersClass
}

// GetGTMioNonOverlappingCountersClass returns the class object for GTMioNonOverlappingCounters.
func GetGTMioNonOverlappingCountersClass() GTMioNonOverlappingCountersClass {
	return getGTMioNonOverlappingCountersClass()
}

type GTMioNonOverlappingCountersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioNonOverlappingCountersClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioNonOverlappingCountersClass) Alloc() GTMioNonOverlappingCounters {
	rv := objc.Send[GTMioNonOverlappingCounters](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioNonOverlappingCounters.CounterValuesForEncoderAtFunctionIndex]
//   - [GTMioNonOverlappingCounters.CounterValuesForGPUCommandAtFunctionIndexSubCommandIndex]
//   - [GTMioNonOverlappingCounters.CounterValuesForPipelineStateIdEncoderFunctionIndex]
//   - [GTMioNonOverlappingCounters.DataIndexForEncoderCounterDataMaster]
//   - [GTMioNonOverlappingCounters.DataIndexForGPUCommandCounterDataMasterRequiresBatchIDFiltering]
//   - [GTMioNonOverlappingCounters.DataIndexForPipelineStateCounterDataMaster]
//   - [GTMioNonOverlappingCounters.DerivedEncoderCounters]
//   - [GTMioNonOverlappingCounters.DerivedGPUCommandCounters]
//   - [GTMioNonOverlappingCounters.DrawCounterNames]
//   - [GTMioNonOverlappingCounters.EncoderCounterForNameDataMaster]
//   - [GTMioNonOverlappingCounters.EncoderCounterNames]
//   - [GTMioNonOverlappingCounters.EncoderCounters]
//   - [GTMioNonOverlappingCounters.GpuCommandCounterForNameDataMaster]
//   - [GTMioNonOverlappingCounters.GpuCommandCounters]
//   - [GTMioNonOverlappingCounters.NumDrawCounters]
//   - [GTMioNonOverlappingCounters.NumEncoderCounters]
//   - [GTMioNonOverlappingCounters.NumPipelineStateCounters]
//   - [GTMioNonOverlappingCounters.PipelineStateCounterNames]
//   - [GTMioNonOverlappingCounters.StatsForDrawCounterAtDataIndexMinValueMaxValueTotalMedian]
//   - [GTMioNonOverlappingCounters.StatsForEncoderCounterAtDataIndexMinValueMaxValueTotalMedian]
//   - [GTMioNonOverlappingCounters.UpdatePerDrawCounters]
//   - [GTMioNonOverlappingCounters.InitBatchIDCountersDrawFunctionIndexes]
//   - [GTMioNonOverlappingCounters.InitWithNonOverlappingCountersScopeScopeIndexDatabase]
type GTMioNonOverlappingCounters struct {
	objectivec.Object
}

// GTMioNonOverlappingCountersFromID constructs a [GTMioNonOverlappingCounters] from an objc.ID.
func GTMioNonOverlappingCountersFromID(id objc.ID) GTMioNonOverlappingCounters {
	return GTMioNonOverlappingCounters{objectivec.Object{ID: id}}
}

// Ensure GTMioNonOverlappingCounters implements IGTMioNonOverlappingCounters.
var _ IGTMioNonOverlappingCounters = GTMioNonOverlappingCounters{}

// An interface definition for the [GTMioNonOverlappingCounters] class.
//
// # Methods
//
//   - [IGTMioNonOverlappingCounters.CounterValuesForEncoderAtFunctionIndex]
//   - [IGTMioNonOverlappingCounters.CounterValuesForGPUCommandAtFunctionIndexSubCommandIndex]
//   - [IGTMioNonOverlappingCounters.CounterValuesForPipelineStateIdEncoderFunctionIndex]
//   - [IGTMioNonOverlappingCounters.DataIndexForEncoderCounterDataMaster]
//   - [IGTMioNonOverlappingCounters.DataIndexForGPUCommandCounterDataMasterRequiresBatchIDFiltering]
//   - [IGTMioNonOverlappingCounters.DataIndexForPipelineStateCounterDataMaster]
//   - [IGTMioNonOverlappingCounters.DerivedEncoderCounters]
//   - [IGTMioNonOverlappingCounters.DerivedGPUCommandCounters]
//   - [IGTMioNonOverlappingCounters.DrawCounterNames]
//   - [IGTMioNonOverlappingCounters.EncoderCounterForNameDataMaster]
//   - [IGTMioNonOverlappingCounters.EncoderCounterNames]
//   - [IGTMioNonOverlappingCounters.EncoderCounters]
//   - [IGTMioNonOverlappingCounters.GpuCommandCounterForNameDataMaster]
//   - [IGTMioNonOverlappingCounters.GpuCommandCounters]
//   - [IGTMioNonOverlappingCounters.NumDrawCounters]
//   - [IGTMioNonOverlappingCounters.NumEncoderCounters]
//   - [IGTMioNonOverlappingCounters.NumPipelineStateCounters]
//   - [IGTMioNonOverlappingCounters.PipelineStateCounterNames]
//   - [IGTMioNonOverlappingCounters.StatsForDrawCounterAtDataIndexMinValueMaxValueTotalMedian]
//   - [IGTMioNonOverlappingCounters.StatsForEncoderCounterAtDataIndexMinValueMaxValueTotalMedian]
//   - [IGTMioNonOverlappingCounters.UpdatePerDrawCounters]
//   - [IGTMioNonOverlappingCounters.InitBatchIDCountersDrawFunctionIndexes]
//   - [IGTMioNonOverlappingCounters.InitWithNonOverlappingCountersScopeScopeIndexDatabase]
type IGTMioNonOverlappingCounters interface {
	objectivec.IObject

	// Topic: Methods

	CounterValuesForEncoderAtFunctionIndex(index uint32) []float64
	CounterValuesForGPUCommandAtFunctionIndexSubCommandIndex(index uint32, index2 int) []float64
	CounterValuesForPipelineStateIdEncoderFunctionIndex(id uint64, index uint32) []float64
	DataIndexForEncoderCounterDataMaster(counter objectivec.IObject, master uint16) uint64
	DataIndexForGPUCommandCounterDataMasterRequiresBatchIDFiltering(counter objectivec.IObject, master uint16, iDFiltering *bool) uint64
	DataIndexForPipelineStateCounterDataMaster(counter objectivec.IObject, master uint16) uint64
	DerivedEncoderCounters() objectivec.IObject
	DerivedGPUCommandCounters() objectivec.IObject
	DrawCounterNames() foundation.INSArray
	EncoderCounterForNameDataMaster(name objectivec.IObject, master uint16) objectivec.IObject
	EncoderCounterNames() foundation.INSArray
	EncoderCounters() foundation.INSArray
	GpuCommandCounterForNameDataMaster(name objectivec.IObject, master uint16) objectivec.IObject
	GpuCommandCounters() foundation.INSArray
	NumDrawCounters() uint64
	NumEncoderCounters() uint64
	NumPipelineStateCounters() uint64
	PipelineStateCounterNames() foundation.INSArray
	StatsForDrawCounterAtDataIndexMinValueMaxValueTotalMedian(index uint64, value []float64, value2 []float64, total []float64, median []float64)
	StatsForEncoderCounterAtDataIndexMinValueMaxValueTotalMedian(index uint64, value []float64, value2 []float64, total []float64, median []float64)
	UpdatePerDrawCounters()
	InitBatchIDCountersDrawFunctionIndexes(iDCounters unsafe.Pointer, indexes unsafe.Pointer) GTMioNonOverlappingCounters
	InitWithNonOverlappingCountersScopeScopeIndexDatabase(counters unsafe.Pointer, scope uint16, index uint32, database objectivec.IObject) GTMioNonOverlappingCounters
}

// Init initializes the instance.
func (g GTMioNonOverlappingCounters) Init() GTMioNonOverlappingCounters {
	rv := objc.Send[GTMioNonOverlappingCounters](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioNonOverlappingCounters) Autorelease() GTMioNonOverlappingCounters {
	rv := objc.Send[GTMioNonOverlappingCounters](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioNonOverlappingCounters creates a new GTMioNonOverlappingCounters instance.
func NewGTMioNonOverlappingCounters() GTMioNonOverlappingCounters {
	class := getGTMioNonOverlappingCountersClass()
	rv := objc.Send[GTMioNonOverlappingCounters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTMioNonOverlappingCountersWithNonOverlappingCountersScopeScopeIndexDatabase(counters unsafe.Pointer, scope uint16, index uint32, database objectivec.IObject) GTMioNonOverlappingCounters {
	instance := getGTMioNonOverlappingCountersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNonOverlappingCounters:scope:scopeIndex:database:"), counters, scope, index, database)
	return GTMioNonOverlappingCountersFromID(rv)
}

func (g GTMioNonOverlappingCounters) CounterValuesForEncoderAtFunctionIndex(index uint32) []float64 {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("counterValuesForEncoderAtFunctionIndex:"), index)
	return objc.ConvertSlice(rv, func(id objc.ID) float64 {
		return float64(id)
	})
}
func (g GTMioNonOverlappingCounters) CounterValuesForGPUCommandAtFunctionIndexSubCommandIndex(index uint32, index2 int) []float64 {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("counterValuesForGPUCommandAtFunctionIndex:subCommandIndex:"), index, index2)
	return objc.ConvertSlice(rv, func(id objc.ID) float64 {
		return float64(id)
	})
}
func (g GTMioNonOverlappingCounters) CounterValuesForPipelineStateIdEncoderFunctionIndex(id uint64, index uint32) []float64 {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("counterValuesForPipelineStateId:encoderFunctionIndex:"), id, index)
	return objc.ConvertSlice(rv, func(id objc.ID) float64 {
		return float64(id)
	})
}
func (g GTMioNonOverlappingCounters) DataIndexForEncoderCounterDataMaster(counter objectivec.IObject, master uint16) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("dataIndexForEncoderCounter:dataMaster:"), counter, master)
	return rv
}
func (g GTMioNonOverlappingCounters) DataIndexForGPUCommandCounterDataMasterRequiresBatchIDFiltering(counter objectivec.IObject, master uint16, iDFiltering *bool) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("dataIndexForGPUCommandCounter:dataMaster:requiresBatchIDFiltering:"), counter, master, iDFiltering)
	return rv
}
func (g GTMioNonOverlappingCounters) DataIndexForPipelineStateCounterDataMaster(counter objectivec.IObject, master uint16) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("dataIndexForPipelineStateCounter:dataMaster:"), counter, master)
	return rv
}
func (g GTMioNonOverlappingCounters) DerivedEncoderCounters() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("derivedEncoderCounters"))
	return objectivec.Object{ID: rv}
}
func (g GTMioNonOverlappingCounters) DerivedGPUCommandCounters() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("derivedGPUCommandCounters"))
	return objectivec.Object{ID: rv}
}
func (g GTMioNonOverlappingCounters) EncoderCounterForNameDataMaster(name objectivec.IObject, master uint16) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("encoderCounterForName:dataMaster:"), name, master)
	return objectivec.Object{ID: rv}
}
func (g GTMioNonOverlappingCounters) GpuCommandCounterForNameDataMaster(name objectivec.IObject, master uint16) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gpuCommandCounterForName:dataMaster:"), name, master)
	return objectivec.Object{ID: rv}
}
func (g GTMioNonOverlappingCounters) StatsForDrawCounterAtDataIndexMinValueMaxValueTotalMedian(index uint64, value []float64, value2 []float64, total []float64, median []float64) {
	objc.Send[objc.ID](g.ID, objc.Sel("statsForDrawCounterAtDataIndex:minValue:maxValue:total:median:"), index, value, value2, total, median)
}
func (g GTMioNonOverlappingCounters) StatsForEncoderCounterAtDataIndexMinValueMaxValueTotalMedian(index uint64, value []float64, value2 []float64, total []float64, median []float64) {
	objc.Send[objc.ID](g.ID, objc.Sel("statsForEncoderCounterAtDataIndex:minValue:maxValue:total:median:"), index, value, value2, total, median)
}
func (g GTMioNonOverlappingCounters) UpdatePerDrawCounters() {
	objc.Send[objc.ID](g.ID, objc.Sel("updatePerDrawCounters"))
}
func (g GTMioNonOverlappingCounters) InitBatchIDCountersDrawFunctionIndexes(iDCounters unsafe.Pointer, indexes unsafe.Pointer) GTMioNonOverlappingCounters {
	rv := objc.Send[GTMioNonOverlappingCounters](g.ID, objc.Sel("initBatchIDCounters:drawFunctionIndexes:"), iDCounters, indexes)
	return rv
}
func (g GTMioNonOverlappingCounters) InitWithNonOverlappingCountersScopeScopeIndexDatabase(counters unsafe.Pointer, scope uint16, index uint32, database objectivec.IObject) GTMioNonOverlappingCounters {
	rv := objc.Send[GTMioNonOverlappingCounters](g.ID, objc.Sel("initWithNonOverlappingCounters:scope:scopeIndex:database:"), counters, scope, index, database)
	return rv
}

func (g GTMioNonOverlappingCounters) DrawCounterNames() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("drawCounterNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTMioNonOverlappingCounters) EncoderCounterNames() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("encoderCounterNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTMioNonOverlappingCounters) EncoderCounters() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("encoderCounters"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTMioNonOverlappingCounters) GpuCommandCounters() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gpuCommandCounters"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTMioNonOverlappingCounters) NumDrawCounters() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("numDrawCounters"))
	return rv
}
func (g GTMioNonOverlappingCounters) NumEncoderCounters() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("numEncoderCounters"))
	return rv
}
func (g GTMioNonOverlappingCounters) NumPipelineStateCounters() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("numPipelineStateCounters"))
	return rv
}
func (g GTMioNonOverlappingCounters) PipelineStateCounterNames() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("pipelineStateCounterNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

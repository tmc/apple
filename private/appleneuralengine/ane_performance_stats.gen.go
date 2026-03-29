// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEPerformanceStats] class.
var (
	_ANEPerformanceStatsClass     ANEPerformanceStatsClass
	_ANEPerformanceStatsClassOnce sync.Once
)

func getANEPerformanceStatsClass() ANEPerformanceStatsClass {
	_ANEPerformanceStatsClassOnce.Do(func() {
		_ANEPerformanceStatsClass = ANEPerformanceStatsClass{class: objc.GetClass("_ANEPerformanceStats")}
	})
	return _ANEPerformanceStatsClass
}

// GetANEPerformanceStatsClass returns the class object for _ANEPerformanceStats.
func GetANEPerformanceStatsClass() ANEPerformanceStatsClass {
	return getANEPerformanceStatsClass()
}

type ANEPerformanceStatsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEPerformanceStatsClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEPerformanceStatsClass) Alloc() ANEPerformanceStats {
	rv := objc.Send[ANEPerformanceStats](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ANEPerformanceStats.EmitPerfcounterSignpostsWithModelStringID]
//   - [ANEPerformanceStats.HwExecutionTime]
//   - [ANEPerformanceStats.PStatsRawData]
//   - [ANEPerformanceStats.PerfCounterData]
//   - [ANEPerformanceStats.PerformanceCounters]
//   - [ANEPerformanceStats.StringForPerfCounter]
//   - [ANEPerformanceStats.InitWithHardwareExecutionPerfCounterDataANEStatsRawData]
//   - [ANEPerformanceStats.InitWithReconstructedDataHardwareExecutionNS]
//   - [ANEPerformanceStats.InitWithRequestPerformanceBufferStatsBufferSize]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStats
type ANEPerformanceStats struct {
	objectivec.Object
}

// ANEPerformanceStatsFromID constructs a [ANEPerformanceStats] from an objc.ID.
func ANEPerformanceStatsFromID(id objc.ID) ANEPerformanceStats {
	return ANEPerformanceStats{objectivec.Object{ID: id}}
}
// Ensure ANEPerformanceStats implements IANEPerformanceStats.
var _ IANEPerformanceStats = ANEPerformanceStats{}

// An interface definition for the [ANEPerformanceStats] class.
//
// # Methods
//
//   - [IANEPerformanceStats.EmitPerfcounterSignpostsWithModelStringID]
//   - [IANEPerformanceStats.HwExecutionTime]
//   - [IANEPerformanceStats.PStatsRawData]
//   - [IANEPerformanceStats.PerfCounterData]
//   - [IANEPerformanceStats.PerformanceCounters]
//   - [IANEPerformanceStats.StringForPerfCounter]
//   - [IANEPerformanceStats.InitWithHardwareExecutionPerfCounterDataANEStatsRawData]
//   - [IANEPerformanceStats.InitWithReconstructedDataHardwareExecutionNS]
//   - [IANEPerformanceStats.InitWithRequestPerformanceBufferStatsBufferSize]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStats
type IANEPerformanceStats interface {
	objectivec.IObject

	// Topic: Methods

	EmitPerfcounterSignpostsWithModelStringID(id uint64)
	HwExecutionTime() uint64
	PStatsRawData() foundation.INSData
	PerfCounterData() foundation.INSData
	PerformanceCounters() objectivec.IObject
	StringForPerfCounter(counter int) objectivec.IObject
	InitWithHardwareExecutionPerfCounterDataANEStatsRawData(execution uint64, data objectivec.IObject, data2 objectivec.IObject) ANEPerformanceStats
	InitWithReconstructedDataHardwareExecutionNS(data objectivec.IObject, ns uint64) ANEPerformanceStats
	InitWithRequestPerformanceBufferStatsBufferSize(buffer unsafe.Pointer, size unsafe.Pointer) ANEPerformanceStats
}

// Init initializes the instance.
func (a ANEPerformanceStats) Init() ANEPerformanceStats {
	rv := objc.Send[ANEPerformanceStats](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEPerformanceStats) Autorelease() ANEPerformanceStats {
	rv := objc.Send[ANEPerformanceStats](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEPerformanceStats creates a new ANEPerformanceStats instance.
func NewANEPerformanceStats() ANEPerformanceStats {
	class := getANEPerformanceStatsClass()
	rv := objc.Send[ANEPerformanceStats](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStats/initWithHardwareExecution:perfCounterData:ANEStatsRawData:
func NewANEPerformanceStatsWithHardwareExecutionPerfCounterDataANEStatsRawData(execution uint64, data objectivec.IObject, data2 objectivec.IObject) ANEPerformanceStats {
	instance := getANEPerformanceStatsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithHardwareExecution:perfCounterData:ANEStatsRawData:"), execution, data, data2)
	return ANEPerformanceStatsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStats/initWithReconstructedData:hardwareExecutionNS:
func NewANEPerformanceStatsWithReconstructedDataHardwareExecutionNS(data objectivec.IObject, ns uint64) ANEPerformanceStats {
	instance := getANEPerformanceStatsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithReconstructedData:hardwareExecutionNS:"), data, ns)
	return ANEPerformanceStatsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStats/initWithRequestPerformanceBuffer:statsBufferSize:
func NewANEPerformanceStatsWithRequestPerformanceBufferStatsBufferSize(buffer unsafe.Pointer, size unsafe.Pointer) ANEPerformanceStats {
	instance := getANEPerformanceStatsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRequestPerformanceBuffer:statsBufferSize:"), buffer, size)
	return ANEPerformanceStatsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStats/emitPerfcounterSignpostsWithModelStringID:
func (a ANEPerformanceStats) EmitPerfcounterSignpostsWithModelStringID(id uint64) {
	objc.Send[objc.ID](a.ID, objc.Sel("emitPerfcounterSignpostsWithModelStringID:"), id)
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStats/performanceCounters
func (a ANEPerformanceStats) PerformanceCounters() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("performanceCounters"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStats/stringForPerfCounter:
func (a ANEPerformanceStats) StringForPerfCounter(counter int) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("stringForPerfCounter:"), counter)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStats/initWithHardwareExecution:perfCounterData:ANEStatsRawData:
func (a ANEPerformanceStats) InitWithHardwareExecutionPerfCounterDataANEStatsRawData(execution uint64, data objectivec.IObject, data2 objectivec.IObject) ANEPerformanceStats {
	rv := objc.Send[ANEPerformanceStats](a.ID, objc.Sel("initWithHardwareExecution:perfCounterData:ANEStatsRawData:"), execution, data, data2)
	return rv
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStats/initWithReconstructedData:hardwareExecutionNS:
func (a ANEPerformanceStats) InitWithReconstructedDataHardwareExecutionNS(data objectivec.IObject, ns uint64) ANEPerformanceStats {
	rv := objc.Send[ANEPerformanceStats](a.ID, objc.Sel("initWithReconstructedData:hardwareExecutionNS:"), data, ns)
	return rv
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStats/initWithRequestPerformanceBuffer:statsBufferSize:
func (a ANEPerformanceStats) InitWithRequestPerformanceBufferStatsBufferSize(buffer unsafe.Pointer, size unsafe.Pointer) ANEPerformanceStats {
	rv := objc.Send[ANEPerformanceStats](a.ID, objc.Sel("initWithRequestPerformanceBuffer:statsBufferSize:"), buffer, size)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStats/driverMaskForANEFMask:
func (_ANEPerformanceStatsClass ANEPerformanceStatsClass) DriverMaskForANEFMask(aNEFMask uint32) uint32 {
	rv := objc.Send[uint32](objc.ID(_ANEPerformanceStatsClass.class), objc.Sel("driverMaskForANEFMask:"), aNEFMask)
	return rv
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStats/statsWithHardwareExecutionNS:
func (_ANEPerformanceStatsClass ANEPerformanceStatsClass) StatsWithHardwareExecutionNS(ns uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEPerformanceStatsClass.class), objc.Sel("statsWithHardwareExecutionNS:"), ns)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStats/statsWithReconstructed:hardwareExecutionNS:aneStatsRawData:
func (_ANEPerformanceStatsClass ANEPerformanceStatsClass) StatsWithReconstructedHardwareExecutionNSAneStatsRawData(reconstructed objectivec.IObject, ns uint64, data objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEPerformanceStatsClass.class), objc.Sel("statsWithReconstructed:hardwareExecutionNS:aneStatsRawData:"), reconstructed, ns, data)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStats/statsWithRequestPerformanceBuffer:statsBufferSize:
func (_ANEPerformanceStatsClass ANEPerformanceStatsClass) StatsWithRequestPerformanceBufferStatsBufferSize(buffer unsafe.Pointer, size unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEPerformanceStatsClass.class), objc.Sel("statsWithRequestPerformanceBuffer:statsBufferSize:"), buffer, size)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStats/hwExecutionTime
func (a ANEPerformanceStats) HwExecutionTime() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("hwExecutionTime"))
	return rv
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStats/pStatsRawData
func (a ANEPerformanceStats) PStatsRawData() foundation.INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("pStatsRawData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStats/perfCounterData
func (a ANEPerformanceStats) PerfCounterData() foundation.INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("perfCounterData"))
	return foundation.NSDataFromID(objc.ID(rv))
}


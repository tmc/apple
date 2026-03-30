// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioTraceDataShaderStat] class.
var (
	_GTMioTraceDataShaderStatClass     GTMioTraceDataShaderStatClass
	_GTMioTraceDataShaderStatClassOnce sync.Once
)

func getGTMioTraceDataShaderStatClass() GTMioTraceDataShaderStatClass {
	_GTMioTraceDataShaderStatClassOnce.Do(func() {
		_GTMioTraceDataShaderStatClass = GTMioTraceDataShaderStatClass{class: objc.GetClass("GTMioTraceDataShaderStat")}
	})
	return _GTMioTraceDataShaderStatClass
}

// GetGTMioTraceDataShaderStatClass returns the class object for GTMioTraceDataShaderStat.
func GetGTMioTraceDataShaderStatClass() GTMioTraceDataShaderStatClass {
	return getGTMioTraceDataShaderStatClass()
}

type GTMioTraceDataShaderStatClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioTraceDataShaderStatClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioTraceDataShaderStatClass) Alloc() GTMioTraceDataShaderStat {
	rv := objc.Send[GTMioTraceDataShaderStat](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioTraceDataShaderStat.NumberOfCliques]
//   - [GTMioTraceDataShaderStat.TotalGPUCycles]
//   - [GTMioTraceDataShaderStat.TotalLatency]
//   - [GTMioTraceDataShaderStat.InitWithStat]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataShaderStat
type GTMioTraceDataShaderStat struct {
	objectivec.Object
}

// GTMioTraceDataShaderStatFromID constructs a [GTMioTraceDataShaderStat] from an objc.ID.
func GTMioTraceDataShaderStatFromID(id objc.ID) GTMioTraceDataShaderStat {
	return GTMioTraceDataShaderStat{objectivec.Object{ID: id}}
}

// Ensure GTMioTraceDataShaderStat implements IGTMioTraceDataShaderStat.
var _ IGTMioTraceDataShaderStat = GTMioTraceDataShaderStat{}

// An interface definition for the [GTMioTraceDataShaderStat] class.
//
// # Methods
//
//   - [IGTMioTraceDataShaderStat.NumberOfCliques]
//   - [IGTMioTraceDataShaderStat.TotalGPUCycles]
//   - [IGTMioTraceDataShaderStat.TotalLatency]
//   - [IGTMioTraceDataShaderStat.InitWithStat]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataShaderStat
type IGTMioTraceDataShaderStat interface {
	objectivec.IObject

	// Topic: Methods

	NumberOfCliques() uint64
	TotalGPUCycles() uint64
	TotalLatency() uint64
	InitWithStat(stat objectivec.IObject) GTMioTraceDataShaderStat
}

// Init initializes the instance.
func (g GTMioTraceDataShaderStat) Init() GTMioTraceDataShaderStat {
	rv := objc.Send[GTMioTraceDataShaderStat](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioTraceDataShaderStat) Autorelease() GTMioTraceDataShaderStat {
	rv := objc.Send[GTMioTraceDataShaderStat](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioTraceDataShaderStat creates a new GTMioTraceDataShaderStat instance.
func NewGTMioTraceDataShaderStat() GTMioTraceDataShaderStat {
	class := getGTMioTraceDataShaderStatClass()
	rv := objc.Send[GTMioTraceDataShaderStat](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataShaderStat/initWithStat:
func NewGTMioTraceDataShaderStatWithStat(stat objectivec.IObject) GTMioTraceDataShaderStat {
	instance := getGTMioTraceDataShaderStatClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStat:"), stat)
	return GTMioTraceDataShaderStatFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataShaderStat/initWithStat:
func (g GTMioTraceDataShaderStat) InitWithStat(stat objectivec.IObject) GTMioTraceDataShaderStat {
	rv := objc.Send[GTMioTraceDataShaderStat](g.ID, objc.Sel("initWithStat:"), stat)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataShaderStat/numberOfCliques
func (g GTMioTraceDataShaderStat) NumberOfCliques() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("numberOfCliques"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataShaderStat/totalGPUCycles
func (g GTMioTraceDataShaderStat) TotalGPUCycles() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("totalGPUCycles"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataShaderStat/totalLatency
func (g GTMioTraceDataShaderStat) TotalLatency() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("totalLatency"))
	return rv
}

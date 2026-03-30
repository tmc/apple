// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTAGX2ShaderProfilerTiming] class.
var (
	_GTAGX2ShaderProfilerTimingClass     GTAGX2ShaderProfilerTimingClass
	_GTAGX2ShaderProfilerTimingClassOnce sync.Once
)

func getGTAGX2ShaderProfilerTimingClass() GTAGX2ShaderProfilerTimingClass {
	_GTAGX2ShaderProfilerTimingClassOnce.Do(func() {
		_GTAGX2ShaderProfilerTimingClass = GTAGX2ShaderProfilerTimingClass{class: objc.GetClass("GTAGX2ShaderProfilerTiming")}
	})
	return _GTAGX2ShaderProfilerTimingClass
}

// GetGTAGX2ShaderProfilerTimingClass returns the class object for GTAGX2ShaderProfilerTiming.
func GetGTAGX2ShaderProfilerTimingClass() GTAGX2ShaderProfilerTimingClass {
	return getGTAGX2ShaderProfilerTimingClass()
}

type GTAGX2ShaderProfilerTimingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTAGX2ShaderProfilerTimingClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTAGX2ShaderProfilerTimingClass) Alloc() GTAGX2ShaderProfilerTiming {
	rv := objc.Send[GTAGX2ShaderProfilerTiming](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTAGX2ShaderProfilerTiming.Cycles]
//   - [GTAGX2ShaderProfilerTiming.MaxCycles]
//   - [GTAGX2ShaderProfilerTiming.MaxTime]
//   - [GTAGX2ShaderProfilerTiming.MinCycles]
//   - [GTAGX2ShaderProfilerTiming.MinTime]
//   - [GTAGX2ShaderProfilerTiming.Time]
//   - [GTAGX2ShaderProfilerTiming.InitWithTiming]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerTiming
type GTAGX2ShaderProfilerTiming struct {
	objectivec.Object
}

// GTAGX2ShaderProfilerTimingFromID constructs a [GTAGX2ShaderProfilerTiming] from an objc.ID.
func GTAGX2ShaderProfilerTimingFromID(id objc.ID) GTAGX2ShaderProfilerTiming {
	return GTAGX2ShaderProfilerTiming{objectivec.Object{ID: id}}
}

// Ensure GTAGX2ShaderProfilerTiming implements IGTAGX2ShaderProfilerTiming.
var _ IGTAGX2ShaderProfilerTiming = GTAGX2ShaderProfilerTiming{}

// An interface definition for the [GTAGX2ShaderProfilerTiming] class.
//
// # Methods
//
//   - [IGTAGX2ShaderProfilerTiming.Cycles]
//   - [IGTAGX2ShaderProfilerTiming.MaxCycles]
//   - [IGTAGX2ShaderProfilerTiming.MaxTime]
//   - [IGTAGX2ShaderProfilerTiming.MinCycles]
//   - [IGTAGX2ShaderProfilerTiming.MinTime]
//   - [IGTAGX2ShaderProfilerTiming.Time]
//   - [IGTAGX2ShaderProfilerTiming.InitWithTiming]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerTiming
type IGTAGX2ShaderProfilerTiming interface {
	objectivec.IObject

	// Topic: Methods

	Cycles() float64
	MaxCycles() float64
	MaxTime() float64
	MinCycles() float64
	MinTime() float64
	Time() float64
	InitWithTiming(timing IGTAGX2ShaderProfilerTiming) GTAGX2ShaderProfilerTiming
}

// Init initializes the instance.
func (g GTAGX2ShaderProfilerTiming) Init() GTAGX2ShaderProfilerTiming {
	rv := objc.Send[GTAGX2ShaderProfilerTiming](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTAGX2ShaderProfilerTiming) Autorelease() GTAGX2ShaderProfilerTiming {
	rv := objc.Send[GTAGX2ShaderProfilerTiming](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTAGX2ShaderProfilerTiming creates a new GTAGX2ShaderProfilerTiming instance.
func NewGTAGX2ShaderProfilerTiming() GTAGX2ShaderProfilerTiming {
	class := getGTAGX2ShaderProfilerTimingClass()
	rv := objc.Send[GTAGX2ShaderProfilerTiming](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerTiming/initWithTiming:
func NewGTAGX2ShaderProfilerTimingWithTiming(timing IGTAGX2ShaderProfilerTiming) GTAGX2ShaderProfilerTiming {
	instance := getGTAGX2ShaderProfilerTimingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTiming:"), timing)
	return GTAGX2ShaderProfilerTimingFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerTiming/cycles
func (g GTAGX2ShaderProfilerTiming) Cycles() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("cycles"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerTiming/maxCycles
func (g GTAGX2ShaderProfilerTiming) MaxCycles() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("maxCycles"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerTiming/maxTime
func (g GTAGX2ShaderProfilerTiming) MaxTime() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("maxTime"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerTiming/minCycles
func (g GTAGX2ShaderProfilerTiming) MinCycles() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("minCycles"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerTiming/minTime
func (g GTAGX2ShaderProfilerTiming) MinTime() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("minTime"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerTiming/time
func (g GTAGX2ShaderProfilerTiming) Time() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("time"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerTiming/initWithTiming:
func (g GTAGX2ShaderProfilerTiming) InitWithTiming(timing IGTAGX2ShaderProfilerTiming) GTAGX2ShaderProfilerTiming {
	rv := objc.Send[GTAGX2ShaderProfilerTiming](g.ID, objc.Sel("initWithTiming:"), timing)
	return rv
}

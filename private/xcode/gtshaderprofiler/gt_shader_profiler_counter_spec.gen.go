// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerCounterSpec] class.
var (
	_GTShaderProfilerCounterSpecClass     GTShaderProfilerCounterSpecClass
	_GTShaderProfilerCounterSpecClassOnce sync.Once
)

func getGTShaderProfilerCounterSpecClass() GTShaderProfilerCounterSpecClass {
	_GTShaderProfilerCounterSpecClassOnce.Do(func() {
		_GTShaderProfilerCounterSpecClass = GTShaderProfilerCounterSpecClass{class: objc.GetClass("GTShaderProfilerCounterSpec")}
	})
	return _GTShaderProfilerCounterSpecClass
}

// GetGTShaderProfilerCounterSpecClass returns the class object for GTShaderProfilerCounterSpec.
func GetGTShaderProfilerCounterSpecClass() GTShaderProfilerCounterSpecClass {
	return getGTShaderProfilerCounterSpecClass()
}

type GTShaderProfilerCounterSpecClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTShaderProfilerCounterSpecClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerCounterSpecClass) Alloc() GTShaderProfilerCounterSpec {
	rv := objc.Send[GTShaderProfilerCounterSpec](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTShaderProfilerCounterSpec.BatchIdFilterableCounterNames]
//   - [GTShaderProfilerCounterSpec.CounterFromName]
//   - [GTShaderProfilerCounterSpec.CounterTableGroups]
//   - [GTShaderProfilerCounterSpec.Counters]
//   - [GTShaderProfilerCounterSpec.FilterSynonyms]
//   - [GTShaderProfilerCounterSpec.TimelineGroups]
//   - [GTShaderProfilerCounterSpec.UpdateMioNonOverlappingCounters]
//   - [GTShaderProfilerCounterSpec.InitWithSpecFile]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterSpec
type GTShaderProfilerCounterSpec struct {
	objectivec.Object
}

// GTShaderProfilerCounterSpecFromID constructs a [GTShaderProfilerCounterSpec] from an objc.ID.
func GTShaderProfilerCounterSpecFromID(id objc.ID) GTShaderProfilerCounterSpec {
	return GTShaderProfilerCounterSpec{objectivec.Object{ID: id}}
}

// Ensure GTShaderProfilerCounterSpec implements IGTShaderProfilerCounterSpec.
var _ IGTShaderProfilerCounterSpec = GTShaderProfilerCounterSpec{}

// An interface definition for the [GTShaderProfilerCounterSpec] class.
//
// # Methods
//
//   - [IGTShaderProfilerCounterSpec.BatchIdFilterableCounterNames]
//   - [IGTShaderProfilerCounterSpec.CounterFromName]
//   - [IGTShaderProfilerCounterSpec.CounterTableGroups]
//   - [IGTShaderProfilerCounterSpec.Counters]
//   - [IGTShaderProfilerCounterSpec.FilterSynonyms]
//   - [IGTShaderProfilerCounterSpec.TimelineGroups]
//   - [IGTShaderProfilerCounterSpec.UpdateMioNonOverlappingCounters]
//   - [IGTShaderProfilerCounterSpec.InitWithSpecFile]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterSpec
type IGTShaderProfilerCounterSpec interface {
	objectivec.IObject

	// Topic: Methods

	BatchIdFilterableCounterNames() objectivec.IObject
	CounterFromName(name objectivec.IObject) objectivec.IObject
	CounterTableGroups() foundation.INSArray
	Counters() foundation.INSArray
	FilterSynonyms() foundation.INSDictionary
	TimelineGroups() foundation.INSArray
	UpdateMioNonOverlappingCounters(counters objectivec.IObject)
	InitWithSpecFile(file objectivec.IObject) GTShaderProfilerCounterSpec
}

// Init initializes the instance.
func (g GTShaderProfilerCounterSpec) Init() GTShaderProfilerCounterSpec {
	rv := objc.Send[GTShaderProfilerCounterSpec](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerCounterSpec) Autorelease() GTShaderProfilerCounterSpec {
	rv := objc.Send[GTShaderProfilerCounterSpec](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerCounterSpec creates a new GTShaderProfilerCounterSpec instance.
func NewGTShaderProfilerCounterSpec() GTShaderProfilerCounterSpec {
	class := getGTShaderProfilerCounterSpecClass()
	rv := objc.Send[GTShaderProfilerCounterSpec](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterSpec/initWithSpecFile:
func NewGTShaderProfilerCounterSpecWithSpecFile(file objectivec.IObject) GTShaderProfilerCounterSpec {
	instance := getGTShaderProfilerCounterSpecClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpecFile:"), file)
	return GTShaderProfilerCounterSpecFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterSpec/batchIdFilterableCounterNames
func (g GTShaderProfilerCounterSpec) BatchIdFilterableCounterNames() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("batchIdFilterableCounterNames"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterSpec/counterFromName:
func (g GTShaderProfilerCounterSpec) CounterFromName(name objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("counterFromName:"), name)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterSpec/updateMioNonOverlappingCounters:
func (g GTShaderProfilerCounterSpec) UpdateMioNonOverlappingCounters(counters objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("updateMioNonOverlappingCounters:"), counters)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterSpec/initWithSpecFile:
func (g GTShaderProfilerCounterSpec) InitWithSpecFile(file objectivec.IObject) GTShaderProfilerCounterSpec {
	rv := objc.Send[GTShaderProfilerCounterSpec](g.ID, objc.Sel("initWithSpecFile:"), file)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterSpec/spec
func (_GTShaderProfilerCounterSpecClass GTShaderProfilerCounterSpecClass) Spec() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_GTShaderProfilerCounterSpecClass.class), objc.Sel("spec"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterSpec/counterTableGroups
func (g GTShaderProfilerCounterSpec) CounterTableGroups() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("counterTableGroups"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterSpec/counters
func (g GTShaderProfilerCounterSpec) Counters() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("counters"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterSpec/filterSynonyms
func (g GTShaderProfilerCounterSpec) FilterSynonyms() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("filterSynonyms"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterSpec/timelineGroups
func (g GTShaderProfilerCounterSpec) TimelineGroups() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("timelineGroups"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

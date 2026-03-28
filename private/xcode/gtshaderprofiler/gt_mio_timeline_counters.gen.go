// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioTimelineCounters] class.
var (
	_GTMioTimelineCountersClass     GTMioTimelineCountersClass
	_GTMioTimelineCountersClassOnce sync.Once
)

func getGTMioTimelineCountersClass() GTMioTimelineCountersClass {
	_GTMioTimelineCountersClassOnce.Do(func() {
		_GTMioTimelineCountersClass = GTMioTimelineCountersClass{class: objc.GetClass("GTMioTimelineCounters")}
	})
	return _GTMioTimelineCountersClass
}

// GetGTMioTimelineCountersClass returns the class object for GTMioTimelineCounters.
func GetGTMioTimelineCountersClass() GTMioTimelineCountersClass {
	return getGTMioTimelineCountersClass()
}

type GTMioTimelineCountersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioTimelineCountersClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioTimelineCountersClass) Alloc() GTMioTimelineCounters {
	rv := objc.Send[GTMioTimelineCounters](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTMioTimelineCounters.CounterForName]
//   - [GTMioTimelineCounters.Counters]
//   - [GTMioTimelineCounters.InitWithTimelineCountersScopeScopeIndex]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTimelineCounters
type GTMioTimelineCounters struct {
	objectivec.Object
}

// GTMioTimelineCountersFromID constructs a [GTMioTimelineCounters] from an objc.ID.
func GTMioTimelineCountersFromID(id objc.ID) GTMioTimelineCounters {
	return GTMioTimelineCounters{objectivec.Object{ID: id}}
}
// Ensure GTMioTimelineCounters implements IGTMioTimelineCounters.
var _ IGTMioTimelineCounters = GTMioTimelineCounters{}

// An interface definition for the [GTMioTimelineCounters] class.
//
// # Methods
//
//   - [IGTMioTimelineCounters.CounterForName]
//   - [IGTMioTimelineCounters.Counters]
//   - [IGTMioTimelineCounters.InitWithTimelineCountersScopeScopeIndex]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTimelineCounters
type IGTMioTimelineCounters interface {
	objectivec.IObject

	// Topic: Methods

	CounterForName(name objectivec.IObject) objectivec.IObject
	Counters() foundation.INSDictionary
	InitWithTimelineCountersScopeScopeIndex(counters unsafe.Pointer, scope uint16, index uint32) GTMioTimelineCounters
}

// Init initializes the instance.
func (g GTMioTimelineCounters) Init() GTMioTimelineCounters {
	rv := objc.Send[GTMioTimelineCounters](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioTimelineCounters) Autorelease() GTMioTimelineCounters {
	rv := objc.Send[GTMioTimelineCounters](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioTimelineCounters creates a new GTMioTimelineCounters instance.
func NewGTMioTimelineCounters() GTMioTimelineCounters {
	class := getGTMioTimelineCountersClass()
	rv := objc.Send[GTMioTimelineCounters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTimelineCounters/initWithTimelineCounters:scope:scopeIndex:
func NewGTMioTimelineCountersWithTimelineCountersScopeScopeIndex(counters unsafe.Pointer, scope uint16, index uint32) GTMioTimelineCounters {
	instance := getGTMioTimelineCountersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTimelineCounters:scope:scopeIndex:"), counters, scope, index)
	return GTMioTimelineCountersFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTimelineCounters/counterForName:
func (g GTMioTimelineCounters) CounterForName(name objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("counterForName:"), name)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTimelineCounters/initWithTimelineCounters:scope:scopeIndex:
func (g GTMioTimelineCounters) InitWithTimelineCountersScopeScopeIndex(counters unsafe.Pointer, scope uint16, index uint32) GTMioTimelineCounters {
	rv := objc.Send[GTMioTimelineCounters](g.ID, objc.Sel("initWithTimelineCounters:scope:scopeIndex:"), counters, scope, index)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTimelineCounters/counters
func (g GTMioTimelineCounters) Counters() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("counters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}


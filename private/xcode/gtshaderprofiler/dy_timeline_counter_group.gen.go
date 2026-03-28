// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DYTimelineCounterGroup] class.
var (
	_DYTimelineCounterGroupClass     DYTimelineCounterGroupClass
	_DYTimelineCounterGroupClassOnce sync.Once
)

func getDYTimelineCounterGroupClass() DYTimelineCounterGroupClass {
	_DYTimelineCounterGroupClassOnce.Do(func() {
		_DYTimelineCounterGroupClass = DYTimelineCounterGroupClass{class: objc.GetClass("DYTimelineCounterGroup")}
	})
	return _DYTimelineCounterGroupClass
}

// GetDYTimelineCounterGroupClass returns the class object for DYTimelineCounterGroup.
func GetDYTimelineCounterGroupClass() DYTimelineCounterGroupClass {
	return getDYTimelineCounterGroupClass()
}

type DYTimelineCounterGroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DYTimelineCounterGroupClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DYTimelineCounterGroupClass) Alloc() DYTimelineCounterGroup {
	rv := objc.Send[DYTimelineCounterGroup](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DYTimelineCounterGroup.CounterNames]
//   - [DYTimelineCounterGroup.SetCounterNames]
//   - [DYTimelineCounterGroup.Counters]
//   - [DYTimelineCounterGroup.SetCounters]
//   - [DYTimelineCounterGroup.EncodeWithCoder]
//   - [DYTimelineCounterGroup.Timestamps]
//   - [DYTimelineCounterGroup.SetTimestamps]
//   - [DYTimelineCounterGroup.InitWithCoder]
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYTimelineCounterGroup
type DYTimelineCounterGroup struct {
	objectivec.Object
}

// DYTimelineCounterGroupFromID constructs a [DYTimelineCounterGroup] from an objc.ID.
func DYTimelineCounterGroupFromID(id objc.ID) DYTimelineCounterGroup {
	return DYTimelineCounterGroup{objectivec.Object{ID: id}}
}
// Ensure DYTimelineCounterGroup implements IDYTimelineCounterGroup.
var _ IDYTimelineCounterGroup = DYTimelineCounterGroup{}

// An interface definition for the [DYTimelineCounterGroup] class.
//
// # Methods
//
//   - [IDYTimelineCounterGroup.CounterNames]
//   - [IDYTimelineCounterGroup.SetCounterNames]
//   - [IDYTimelineCounterGroup.Counters]
//   - [IDYTimelineCounterGroup.SetCounters]
//   - [IDYTimelineCounterGroup.EncodeWithCoder]
//   - [IDYTimelineCounterGroup.Timestamps]
//   - [IDYTimelineCounterGroup.SetTimestamps]
//   - [IDYTimelineCounterGroup.InitWithCoder]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYTimelineCounterGroup
type IDYTimelineCounterGroup interface {
	objectivec.IObject

	// Topic: Methods

	CounterNames() foundation.INSArray
	SetCounterNames(value foundation.INSArray)
	Counters() foundation.INSArray
	SetCounters(value foundation.INSArray)
	EncodeWithCoder(coder foundation.INSCoder)
	Timestamps() foundation.INSData
	SetTimestamps(value foundation.INSData)
	InitWithCoder(coder foundation.INSCoder) DYTimelineCounterGroup
}

// Init initializes the instance.
func (d DYTimelineCounterGroup) Init() DYTimelineCounterGroup {
	rv := objc.Send[DYTimelineCounterGroup](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DYTimelineCounterGroup) Autorelease() DYTimelineCounterGroup {
	rv := objc.Send[DYTimelineCounterGroup](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDYTimelineCounterGroup creates a new DYTimelineCounterGroup instance.
func NewDYTimelineCounterGroup() DYTimelineCounterGroup {
	class := getDYTimelineCounterGroupClass()
	rv := objc.Send[DYTimelineCounterGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYTimelineCounterGroup/initWithCoder:
func NewDYTimelineCounterGroupWithCoder(coder objectivec.IObject) DYTimelineCounterGroup {
	instance := getDYTimelineCounterGroupClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DYTimelineCounterGroupFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYTimelineCounterGroup/encodeWithCoder:
func (d DYTimelineCounterGroup) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYTimelineCounterGroup/initWithCoder:
func (d DYTimelineCounterGroup) InitWithCoder(coder foundation.INSCoder) DYTimelineCounterGroup {
	rv := objc.Send[DYTimelineCounterGroup](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYTimelineCounterGroup/supportsSecureCoding
func (_DYTimelineCounterGroupClass DYTimelineCounterGroupClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_DYTimelineCounterGroupClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYTimelineCounterGroup/counterNames
func (d DYTimelineCounterGroup) CounterNames() foundation.INSArray {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("counterNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (d DYTimelineCounterGroup) SetCounterNames(value foundation.INSArray) {
	objc.Send[struct{}](d.ID, objc.Sel("setCounterNames:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYTimelineCounterGroup/counters
func (d DYTimelineCounterGroup) Counters() foundation.INSArray {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("counters"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (d DYTimelineCounterGroup) SetCounters(value foundation.INSArray) {
	objc.Send[struct{}](d.ID, objc.Sel("setCounters:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYTimelineCounterGroup/timestamps
func (d DYTimelineCounterGroup) Timestamps() foundation.INSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("timestamps"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d DYTimelineCounterGroup) SetTimestamps(value foundation.INSData) {
	objc.Send[struct{}](d.ID, objc.Sel("setTimestamps:"), value)
}


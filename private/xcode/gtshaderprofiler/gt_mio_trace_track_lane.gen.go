// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioTraceTrackLane] class.
var (
	_GTMioTraceTrackLaneClass     GTMioTraceTrackLaneClass
	_GTMioTraceTrackLaneClassOnce sync.Once
)

func getGTMioTraceTrackLaneClass() GTMioTraceTrackLaneClass {
	_GTMioTraceTrackLaneClassOnce.Do(func() {
		_GTMioTraceTrackLaneClass = GTMioTraceTrackLaneClass{class: objc.GetClass("GTMioTraceTrackLane")}
	})
	return _GTMioTraceTrackLaneClass
}

// GetGTMioTraceTrackLaneClass returns the class object for GTMioTraceTrackLane.
func GetGTMioTraceTrackLaneClass() GTMioTraceTrackLaneClass {
	return getGTMioTraceTrackLaneClass()
}

type GTMioTraceTrackLaneClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioTraceTrackLaneClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioTraceTrackLaneClass) Alloc() GTMioTraceTrackLane {
	rv := objc.Send[GTMioTraceTrackLane](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioTraceTrackLane.Add]
//   - [GTMioTraceTrackLane.IndexCount]
//   - [GTMioTraceTrackLane.Indexes]
//   - [GTMioTraceTrackLane.IsEmpty]
//   - [GTMioTraceTrackLane.LaneId]
//   - [GTMioTraceTrackLane.InitWithId]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrackLane
type GTMioTraceTrackLane struct {
	objectivec.Object
}

// GTMioTraceTrackLaneFromID constructs a [GTMioTraceTrackLane] from an objc.ID.
func GTMioTraceTrackLaneFromID(id objc.ID) GTMioTraceTrackLane {
	return GTMioTraceTrackLane{objectivec.Object{ID: id}}
}

// Ensure GTMioTraceTrackLane implements IGTMioTraceTrackLane.
var _ IGTMioTraceTrackLane = GTMioTraceTrackLane{}

// An interface definition for the [GTMioTraceTrackLane] class.
//
// # Methods
//
//   - [IGTMioTraceTrackLane.Add]
//   - [IGTMioTraceTrackLane.IndexCount]
//   - [IGTMioTraceTrackLane.Indexes]
//   - [IGTMioTraceTrackLane.IsEmpty]
//   - [IGTMioTraceTrackLane.LaneId]
//   - [IGTMioTraceTrackLane.InitWithId]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrackLane
type IGTMioTraceTrackLane interface {
	objectivec.IObject

	// Topic: Methods

	Add(add uint64)
	IndexCount() uint64
	Indexes() unsafe.Pointer
	IsEmpty() bool
	LaneId() int
	InitWithId(id int) GTMioTraceTrackLane
}

// Init initializes the instance.
func (g GTMioTraceTrackLane) Init() GTMioTraceTrackLane {
	rv := objc.Send[GTMioTraceTrackLane](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioTraceTrackLane) Autorelease() GTMioTraceTrackLane {
	rv := objc.Send[GTMioTraceTrackLane](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioTraceTrackLane creates a new GTMioTraceTrackLane instance.
func NewGTMioTraceTrackLane() GTMioTraceTrackLane {
	class := getGTMioTraceTrackLaneClass()
	rv := objc.Send[GTMioTraceTrackLane](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrackLane/initWithId:
func NewGTMioTraceTrackLaneWithId(id int) GTMioTraceTrackLane {
	instance := getGTMioTraceTrackLaneClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithId:"), id)
	return GTMioTraceTrackLaneFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrackLane/add:
func (g GTMioTraceTrackLane) Add(add uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("add:"), add)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrackLane/isEmpty
func (g GTMioTraceTrackLane) IsEmpty() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isEmpty"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrackLane/initWithId:
func (g GTMioTraceTrackLane) InitWithId(id int) GTMioTraceTrackLane {
	rv := objc.Send[GTMioTraceTrackLane](g.ID, objc.Sel("initWithId:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrackLane/indexCount
func (g GTMioTraceTrackLane) IndexCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("indexCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrackLane/indexes
func (g GTMioTraceTrackLane) Indexes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("indexes"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrackLane/laneId
func (g GTMioTraceTrackLane) LaneId() int {
	rv := objc.Send[int](g.ID, objc.Sel("laneId"))
	return rv
}

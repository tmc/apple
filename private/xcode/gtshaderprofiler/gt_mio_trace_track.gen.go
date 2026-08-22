// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioTraceTrack] class.
var (
	_GTMioTraceTrackClass     GTMioTraceTrackClass
	_GTMioTraceTrackClassOnce sync.Once
)

func getGTMioTraceTrackClass() GTMioTraceTrackClass {
	_GTMioTraceTrackClassOnce.Do(func() {
		_GTMioTraceTrackClass = GTMioTraceTrackClass{class: objc.GetClass("GTMioTraceTrack")}
	})
	return _GTMioTraceTrackClass
}

// GetGTMioTraceTrackClass returns the class object for GTMioTraceTrack.
func GetGTMioTraceTrackClass() GTMioTraceTrackClass {
	return getGTMioTraceTrackClass()
}

type GTMioTraceTrackClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioTraceTrackClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioTraceTrackClass) Alloc() GTMioTraceTrack {
	rv := objc.SendIfResponds[GTMioTraceTrack](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioTraceTrack.Context]
//   - [GTMioTraceTrack.Duration]
//   - [GTMioTraceTrack.EndTimestamp]
//   - [GTMioTraceTrack.FirstIndex]
//   - [GTMioTraceTrack.IsEmpty]
//   - [GTMioTraceTrack.LaneIdForStartEnd]
//   - [GTMioTraceTrack.Lanes]
//   - [GTMioTraceTrack.StartTimestamp]
//   - [GTMioTraceTrack.TakeEndIndex]
//   - [GTMioTraceTrack.TakeEndIndexLaneId]
//   - [GTMioTraceTrack.TakeEndIndexesBeginIndexesEndLaneId]
//   - [GTMioTraceTrack.TrackId]
//   - [GTMioTraceTrack.InitWithIdScopeScopeIdentifierLevelLevelIdentifier]
type GTMioTraceTrack struct {
	objectivec.Object
}

// GTMioTraceTrackFromID constructs a [GTMioTraceTrack] from an objc.ID.
func GTMioTraceTrackFromID(id objc.ID) GTMioTraceTrack {
	return GTMioTraceTrack{objectivec.Object{ID: id}}
}

// Ensure GTMioTraceTrack implements IGTMioTraceTrack.
var _ IGTMioTraceTrack = GTMioTraceTrack{}

// An interface definition for the [GTMioTraceTrack] class.
//
// # Methods
//
//   - [IGTMioTraceTrack.Context]
//   - [IGTMioTraceTrack.Duration]
//   - [IGTMioTraceTrack.EndTimestamp]
//   - [IGTMioTraceTrack.FirstIndex]
//   - [IGTMioTraceTrack.IsEmpty]
//   - [IGTMioTraceTrack.LaneIdForStartEnd]
//   - [IGTMioTraceTrack.Lanes]
//   - [IGTMioTraceTrack.StartTimestamp]
//   - [IGTMioTraceTrack.TakeEndIndex]
//   - [IGTMioTraceTrack.TakeEndIndexLaneId]
//   - [IGTMioTraceTrack.TakeEndIndexesBeginIndexesEndLaneId]
//   - [IGTMioTraceTrack.TrackId]
//   - [IGTMioTraceTrack.InitWithIdScopeScopeIdentifierLevelLevelIdentifier]
type IGTMioTraceTrack interface {
	objectivec.IObject

	// Topic: Methods

	Context() *GTMioCostContext
	Duration() uint64
	EndTimestamp() uint64
	FirstIndex() uint64
	IsEmpty() bool
	LaneIdForStartEnd(start uint64, end uint64) int
	Lanes() foundation.INSArray
	StartTimestamp() uint64
	TakeEndIndex(take uint64, end uint64, index uint64)
	TakeEndIndexLaneId(take uint64, end uint64, index uint64, id int)
	TakeEndIndexesBeginIndexesEndLaneId(take uint64, end uint64, begin uint64, end2 uint64, id int)
	TrackId() int
	InitWithIdScopeScopeIdentifierLevelLevelIdentifier(id int, scope uint16, identifier uint64, level uint16, identifier2 uint32) GTMioTraceTrack
}

// Init initializes the instance.
func (g GTMioTraceTrack) Init() GTMioTraceTrack {
	rv := objc.SendIfResponds[GTMioTraceTrack](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioTraceTrack) Autorelease() GTMioTraceTrack {
	rv := objc.SendIfResponds[GTMioTraceTrack](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioTraceTrack creates a new GTMioTraceTrack instance.
func NewGTMioTraceTrack() GTMioTraceTrack {
	class := getGTMioTraceTrackClass()
	rv := objc.SendIfResponds[GTMioTraceTrack](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTMioTraceTrackWithIdScopeScopeIdentifierLevelLevelIdentifier(id int, scope uint16, identifier uint64, level uint16, identifier2 uint32) GTMioTraceTrack {
	instance := getGTMioTraceTrackClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithId:scope:scopeIdentifier:level:levelIdentifier:"), id, scope, identifier, level, identifier2)
	return GTMioTraceTrackFromID(rv)
}

func (g GTMioTraceTrack) LaneIdForStartEnd(start uint64, end uint64) int {
	rv := objc.SendIfResponds[int](g.ID, objc.Sel("laneIdForStart:end:"), start, end)
	return rv
}
func (g GTMioTraceTrack) TakeEndIndex(take uint64, end uint64, index uint64) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("take:end:index:"), take, end, index)
}
func (g GTMioTraceTrack) TakeEndIndexLaneId(take uint64, end uint64, index uint64, id int) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("take:end:index:laneId:"), take, end, index, id)
}
func (g GTMioTraceTrack) TakeEndIndexesBeginIndexesEndLaneId(take uint64, end uint64, begin uint64, end2 uint64, id int) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("take:end:indexesBegin:indexesEnd:laneId:"), take, end, begin, end2, id)
}
func (g GTMioTraceTrack) InitWithIdScopeScopeIdentifierLevelLevelIdentifier(id int, scope uint16, identifier uint64, level uint16, identifier2 uint32) GTMioTraceTrack {
	rv := objc.SendIfResponds[GTMioTraceTrack](g.ID, objc.Sel("initWithId:scope:scopeIdentifier:level:levelIdentifier:"), id, scope, identifier, level, identifier2)
	return rv
}

func (g GTMioTraceTrack) Context() *GTMioCostContext {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("context"))
	return (*GTMioCostContext)(rv)
}
func (g GTMioTraceTrack) Duration() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("duration"))
	return rv
}
func (g GTMioTraceTrack) EndTimestamp() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("endTimestamp"))
	return rv
}
func (g GTMioTraceTrack) FirstIndex() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("firstIndex"))
	return rv
}
func (g GTMioTraceTrack) IsEmpty() bool {
	rv := objc.SendIfResponds[bool](g.ID, objc.Sel("isEmpty"))
	return rv
}
func (g GTMioTraceTrack) Lanes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("lanes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTMioTraceTrack) StartTimestamp() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("startTimestamp"))
	return rv
}
func (g GTMioTraceTrack) TrackId() int {
	rv := objc.SendIfResponds[int](g.ID, objc.Sel("trackId"))
	return rv
}

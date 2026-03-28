// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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
	rv := objc.Send[GTMioTraceTrack](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
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
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrack
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
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrack
type IGTMioTraceTrack interface {
	objectivec.IObject

	// Topic: Methods

	Context() unsafe.Pointer
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
	rv := objc.Send[GTMioTraceTrack](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioTraceTrack) Autorelease() GTMioTraceTrack {
	rv := objc.Send[GTMioTraceTrack](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioTraceTrack creates a new GTMioTraceTrack instance.
func NewGTMioTraceTrack() GTMioTraceTrack {
	class := getGTMioTraceTrackClass()
	rv := objc.Send[GTMioTraceTrack](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrack/initWithId:scope:scopeIdentifier:level:levelIdentifier:
func NewGTMioTraceTrackWithIdScopeScopeIdentifierLevelLevelIdentifier(id int, scope uint16, identifier uint64, level uint16, identifier2 uint32) GTMioTraceTrack {
	instance := getGTMioTraceTrackClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithId:scope:scopeIdentifier:level:levelIdentifier:"), id, scope, identifier, level, identifier2)
	return GTMioTraceTrackFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrack/laneIdForStart:end:
func (g GTMioTraceTrack) LaneIdForStartEnd(start uint64, end uint64) int {
	rv := objc.Send[int](g.ID, objc.Sel("laneIdForStart:end:"), start, end)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrack/take:end:index:
func (g GTMioTraceTrack) TakeEndIndex(take uint64, end uint64, index uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("take:end:index:"), take, end, index)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrack/take:end:index:laneId:
func (g GTMioTraceTrack) TakeEndIndexLaneId(take uint64, end uint64, index uint64, id int) {
	objc.Send[objc.ID](g.ID, objc.Sel("take:end:index:laneId:"), take, end, index, id)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrack/take:end:indexesBegin:indexesEnd:laneId:
func (g GTMioTraceTrack) TakeEndIndexesBeginIndexesEndLaneId(take uint64, end uint64, begin uint64, end2 uint64, id int) {
	objc.Send[objc.ID](g.ID, objc.Sel("take:end:indexesBegin:indexesEnd:laneId:"), take, end, begin, end2, id)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrack/initWithId:scope:scopeIdentifier:level:levelIdentifier:
func (g GTMioTraceTrack) InitWithIdScopeScopeIdentifierLevelLevelIdentifier(id int, scope uint16, identifier uint64, level uint16, identifier2 uint32) GTMioTraceTrack {
	rv := objc.Send[GTMioTraceTrack](g.ID, objc.Sel("initWithId:scope:scopeIdentifier:level:levelIdentifier:"), id, scope, identifier, level, identifier2)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrack/context
func (g GTMioTraceTrack) Context() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("context"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrack/duration
func (g GTMioTraceTrack) Duration() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("duration"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrack/endTimestamp
func (g GTMioTraceTrack) EndTimestamp() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("endTimestamp"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrack/firstIndex
func (g GTMioTraceTrack) FirstIndex() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("firstIndex"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrack/isEmpty
func (g GTMioTraceTrack) IsEmpty() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isEmpty"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrack/lanes
func (g GTMioTraceTrack) Lanes() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("lanes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrack/startTimestamp
func (g GTMioTraceTrack) StartTimestamp() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("startTimestamp"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrack/trackId
func (g GTMioTraceTrack) TrackId() int {
	rv := objc.Send[int](g.ID, objc.Sel("trackId"))
	return rv
}


// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [GTMioTraceAggregatedDrawTrack] class.
var (
	_GTMioTraceAggregatedDrawTrackClass     GTMioTraceAggregatedDrawTrackClass
	_GTMioTraceAggregatedDrawTrackClassOnce sync.Once
)

func getGTMioTraceAggregatedDrawTrackClass() GTMioTraceAggregatedDrawTrackClass {
	_GTMioTraceAggregatedDrawTrackClassOnce.Do(func() {
		_GTMioTraceAggregatedDrawTrackClass = GTMioTraceAggregatedDrawTrackClass{class: objc.GetClass("GTMioTraceAggregatedDrawTrack")}
	})
	return _GTMioTraceAggregatedDrawTrackClass
}

// GetGTMioTraceAggregatedDrawTrackClass returns the class object for GTMioTraceAggregatedDrawTrack.
func GetGTMioTraceAggregatedDrawTrackClass() GTMioTraceAggregatedDrawTrackClass {
	return getGTMioTraceAggregatedDrawTrackClass()
}

type GTMioTraceAggregatedDrawTrackClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioTraceAggregatedDrawTrackClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioTraceAggregatedDrawTrackClass) Alloc() GTMioTraceAggregatedDrawTrack {
	rv := objc.Send[GTMioTraceAggregatedDrawTrack](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioTraceAggregatedDrawTrack.PostProcess]
//   - [GTMioTraceAggregatedDrawTrack.Take]
//   - [GTMioTraceAggregatedDrawTrack.TraceCount]
//   - [GTMioTraceAggregatedDrawTrack.Traces]
type GTMioTraceAggregatedDrawTrack struct {
	GTMioTraceTrack
}

// GTMioTraceAggregatedDrawTrackFromID constructs a [GTMioTraceAggregatedDrawTrack] from an objc.ID.
func GTMioTraceAggregatedDrawTrackFromID(id objc.ID) GTMioTraceAggregatedDrawTrack {
	return GTMioTraceAggregatedDrawTrack{GTMioTraceTrack: GTMioTraceTrackFromID(id)}
}

// Ensure GTMioTraceAggregatedDrawTrack implements IGTMioTraceAggregatedDrawTrack.
var _ IGTMioTraceAggregatedDrawTrack = GTMioTraceAggregatedDrawTrack{}

// An interface definition for the [GTMioTraceAggregatedDrawTrack] class.
//
// # Methods
//
//   - [IGTMioTraceAggregatedDrawTrack.PostProcess]
//   - [IGTMioTraceAggregatedDrawTrack.Take]
//   - [IGTMioTraceAggregatedDrawTrack.TraceCount]
//   - [IGTMioTraceAggregatedDrawTrack.Traces]
type IGTMioTraceAggregatedDrawTrack interface {
	IGTMioTraceTrack

	// Topic: Methods

	PostProcess()
	Take(take *GTMioDrawTrace)
	TraceCount() uint64
	Traces() unsafe.Pointer
}

// Init initializes the instance.
func (g GTMioTraceAggregatedDrawTrack) Init() GTMioTraceAggregatedDrawTrack {
	rv := objc.Send[GTMioTraceAggregatedDrawTrack](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioTraceAggregatedDrawTrack) Autorelease() GTMioTraceAggregatedDrawTrack {
	rv := objc.Send[GTMioTraceAggregatedDrawTrack](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioTraceAggregatedDrawTrack creates a new GTMioTraceAggregatedDrawTrack instance.
func NewGTMioTraceAggregatedDrawTrack() GTMioTraceAggregatedDrawTrack {
	class := getGTMioTraceAggregatedDrawTrackClass()
	rv := objc.Send[GTMioTraceAggregatedDrawTrack](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTMioTraceAggregatedDrawTrackWithIdScopeScopeIdentifierLevelLevelIdentifier(id int, scope uint16, identifier uint64, level uint16, identifier2 uint32) GTMioTraceAggregatedDrawTrack {
	instance := getGTMioTraceAggregatedDrawTrackClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithId:scope:scopeIdentifier:level:levelIdentifier:"), id, scope, identifier, level, identifier2)
	return GTMioTraceAggregatedDrawTrackFromID(rv)
}

func (g GTMioTraceAggregatedDrawTrack) PostProcess() {
	objc.Send[objc.ID](g.ID, objc.Sel("postProcess"))
}
func (g GTMioTraceAggregatedDrawTrack) Take(take *GTMioDrawTrace) {
	objc.Send[objc.ID](g.ID, objc.Sel("take:"), take)
}

func (g GTMioTraceAggregatedDrawTrack) TraceCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("traceCount"))
	return rv
}
func (g GTMioTraceAggregatedDrawTrack) Traces() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("traces"))
	return rv
}

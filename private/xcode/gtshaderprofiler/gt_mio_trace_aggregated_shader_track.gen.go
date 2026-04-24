// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [GTMioTraceAggregatedShaderTrack] class.
var (
	_GTMioTraceAggregatedShaderTrackClass     GTMioTraceAggregatedShaderTrackClass
	_GTMioTraceAggregatedShaderTrackClassOnce sync.Once
)

func getGTMioTraceAggregatedShaderTrackClass() GTMioTraceAggregatedShaderTrackClass {
	_GTMioTraceAggregatedShaderTrackClassOnce.Do(func() {
		_GTMioTraceAggregatedShaderTrackClass = GTMioTraceAggregatedShaderTrackClass{class: objc.GetClass("GTMioTraceAggregatedShaderTrack")}
	})
	return _GTMioTraceAggregatedShaderTrackClass
}

// GetGTMioTraceAggregatedShaderTrackClass returns the class object for GTMioTraceAggregatedShaderTrack.
func GetGTMioTraceAggregatedShaderTrackClass() GTMioTraceAggregatedShaderTrackClass {
	return getGTMioTraceAggregatedShaderTrackClass()
}

type GTMioTraceAggregatedShaderTrackClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioTraceAggregatedShaderTrackClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioTraceAggregatedShaderTrackClass) Alloc() GTMioTraceAggregatedShaderTrack {
	rv := objc.Send[GTMioTraceAggregatedShaderTrack](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioTraceAggregatedShaderTrack.PostProcess]
//   - [GTMioTraceAggregatedShaderTrack.Take]
//   - [GTMioTraceAggregatedShaderTrack.TraceCount]
//   - [GTMioTraceAggregatedShaderTrack.Traces]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceAggregatedShaderTrack
type GTMioTraceAggregatedShaderTrack struct {
	GTMioTraceTrack
}

// GTMioTraceAggregatedShaderTrackFromID constructs a [GTMioTraceAggregatedShaderTrack] from an objc.ID.
func GTMioTraceAggregatedShaderTrackFromID(id objc.ID) GTMioTraceAggregatedShaderTrack {
	return GTMioTraceAggregatedShaderTrack{GTMioTraceTrack: GTMioTraceTrackFromID(id)}
}

// Ensure GTMioTraceAggregatedShaderTrack implements IGTMioTraceAggregatedShaderTrack.
var _ IGTMioTraceAggregatedShaderTrack = GTMioTraceAggregatedShaderTrack{}

// An interface definition for the [GTMioTraceAggregatedShaderTrack] class.
//
// # Methods
//
//   - [IGTMioTraceAggregatedShaderTrack.PostProcess]
//   - [IGTMioTraceAggregatedShaderTrack.Take]
//   - [IGTMioTraceAggregatedShaderTrack.TraceCount]
//   - [IGTMioTraceAggregatedShaderTrack.Traces]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceAggregatedShaderTrack
type IGTMioTraceAggregatedShaderTrack interface {
	IGTMioTraceTrack

	// Topic: Methods

	PostProcess()
	Take(take *GTMioBinaryTraceRef)
	TraceCount() uint64
	Traces() *GTMioBinaryTraceRef
}

// Init initializes the instance.
func (g GTMioTraceAggregatedShaderTrack) Init() GTMioTraceAggregatedShaderTrack {
	rv := objc.Send[GTMioTraceAggregatedShaderTrack](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioTraceAggregatedShaderTrack) Autorelease() GTMioTraceAggregatedShaderTrack {
	rv := objc.Send[GTMioTraceAggregatedShaderTrack](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioTraceAggregatedShaderTrack creates a new GTMioTraceAggregatedShaderTrack instance.
func NewGTMioTraceAggregatedShaderTrack() GTMioTraceAggregatedShaderTrack {
	class := getGTMioTraceAggregatedShaderTrackClass()
	rv := objc.Send[GTMioTraceAggregatedShaderTrack](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrack/initWithId:scope:scopeIdentifier:level:levelIdentifier:
func NewGTMioTraceAggregatedShaderTrackWithIdScopeScopeIdentifierLevelLevelIdentifier(id int, scope uint16, identifier uint64, level uint16, identifier2 uint32) GTMioTraceAggregatedShaderTrack {
	instance := getGTMioTraceAggregatedShaderTrackClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithId:scope:scopeIdentifier:level:levelIdentifier:"), id, scope, identifier, level, identifier2)
	return GTMioTraceAggregatedShaderTrackFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceAggregatedShaderTrack/postProcess
func (g GTMioTraceAggregatedShaderTrack) PostProcess() {
	objc.Send[objc.ID](g.ID, objc.Sel("postProcess"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceAggregatedShaderTrack/take:
func (g GTMioTraceAggregatedShaderTrack) Take(take *GTMioBinaryTraceRef) {
	objc.Send[objc.ID](g.ID, objc.Sel("take:"), take)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceAggregatedShaderTrack/traceCount
func (g GTMioTraceAggregatedShaderTrack) TraceCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("traceCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceAggregatedShaderTrack/traces
func (g GTMioTraceAggregatedShaderTrack) Traces() *GTMioBinaryTraceRef {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("traces"))
	return (*GTMioBinaryTraceRef)(rv)
}

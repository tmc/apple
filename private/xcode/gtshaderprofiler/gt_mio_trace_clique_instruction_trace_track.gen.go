// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [GTMioTraceCliqueInstructionTraceTrack] class.
var (
	_GTMioTraceCliqueInstructionTraceTrackClass     GTMioTraceCliqueInstructionTraceTrackClass
	_GTMioTraceCliqueInstructionTraceTrackClassOnce sync.Once
)

func getGTMioTraceCliqueInstructionTraceTrackClass() GTMioTraceCliqueInstructionTraceTrackClass {
	_GTMioTraceCliqueInstructionTraceTrackClassOnce.Do(func() {
		_GTMioTraceCliqueInstructionTraceTrackClass = GTMioTraceCliqueInstructionTraceTrackClass{class: objc.GetClass("GTMioTraceCliqueInstructionTraceTrack")}
	})
	return _GTMioTraceCliqueInstructionTraceTrackClass
}

// GetGTMioTraceCliqueInstructionTraceTrackClass returns the class object for GTMioTraceCliqueInstructionTraceTrack.
func GetGTMioTraceCliqueInstructionTraceTrackClass() GTMioTraceCliqueInstructionTraceTrackClass {
	return getGTMioTraceCliqueInstructionTraceTrackClass()
}

type GTMioTraceCliqueInstructionTraceTrackClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioTraceCliqueInstructionTraceTrackClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioTraceCliqueInstructionTraceTrackClass) Alloc() GTMioTraceCliqueInstructionTraceTrack {
	rv := objc.Send[GTMioTraceCliqueInstructionTraceTrack](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioTraceCliqueInstructionTraceTrack.PostProcess]
//   - [GTMioTraceCliqueInstructionTraceTrack.TakeCliqueTraces]
//   - [GTMioTraceCliqueInstructionTraceTrack.TraceCount]
//   - [GTMioTraceCliqueInstructionTraceTrack.Traces]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceCliqueInstructionTraceTrack
type GTMioTraceCliqueInstructionTraceTrack struct {
	GTMioTraceTrack
}

// GTMioTraceCliqueInstructionTraceTrackFromID constructs a [GTMioTraceCliqueInstructionTraceTrack] from an objc.ID.
func GTMioTraceCliqueInstructionTraceTrackFromID(id objc.ID) GTMioTraceCliqueInstructionTraceTrack {
	return GTMioTraceCliqueInstructionTraceTrack{GTMioTraceTrack: GTMioTraceTrackFromID(id)}
}

// Ensure GTMioTraceCliqueInstructionTraceTrack implements IGTMioTraceCliqueInstructionTraceTrack.
var _ IGTMioTraceCliqueInstructionTraceTrack = GTMioTraceCliqueInstructionTraceTrack{}

// An interface definition for the [GTMioTraceCliqueInstructionTraceTrack] class.
//
// # Methods
//
//   - [IGTMioTraceCliqueInstructionTraceTrack.PostProcess]
//   - [IGTMioTraceCliqueInstructionTraceTrack.TakeCliqueTraces]
//   - [IGTMioTraceCliqueInstructionTraceTrack.TraceCount]
//   - [IGTMioTraceCliqueInstructionTraceTrack.Traces]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceCliqueInstructionTraceTrack
type IGTMioTraceCliqueInstructionTraceTrack interface {
	IGTMioTraceTrack

	// Topic: Methods

	PostProcess()
	TakeCliqueTraces(clique uint32, traces unsafe.Pointer)
	TraceCount() uint64
	Traces() unsafe.Pointer
}

// Init initializes the instance.
func (g GTMioTraceCliqueInstructionTraceTrack) Init() GTMioTraceCliqueInstructionTraceTrack {
	rv := objc.Send[GTMioTraceCliqueInstructionTraceTrack](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioTraceCliqueInstructionTraceTrack) Autorelease() GTMioTraceCliqueInstructionTraceTrack {
	rv := objc.Send[GTMioTraceCliqueInstructionTraceTrack](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioTraceCliqueInstructionTraceTrack creates a new GTMioTraceCliqueInstructionTraceTrack instance.
func NewGTMioTraceCliqueInstructionTraceTrack() GTMioTraceCliqueInstructionTraceTrack {
	class := getGTMioTraceCliqueInstructionTraceTrackClass()
	rv := objc.Send[GTMioTraceCliqueInstructionTraceTrack](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrack/initWithId:scope:scopeIdentifier:level:levelIdentifier:
func NewGTMioTraceCliqueInstructionTraceTrackWithIdScopeScopeIdentifierLevelLevelIdentifier(id int, scope uint16, identifier uint64, level uint16, identifier2 uint32) GTMioTraceCliqueInstructionTraceTrack {
	instance := getGTMioTraceCliqueInstructionTraceTrackClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithId:scope:scopeIdentifier:level:levelIdentifier:"), id, scope, identifier, level, identifier2)
	return GTMioTraceCliqueInstructionTraceTrackFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceCliqueInstructionTraceTrack/postProcess
func (g GTMioTraceCliqueInstructionTraceTrack) PostProcess() {
	objc.Send[objc.ID](g.ID, objc.Sel("postProcess"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceCliqueInstructionTraceTrack/takeClique:traces:
func (g GTMioTraceCliqueInstructionTraceTrack) TakeCliqueTraces(clique uint32, traces unsafe.Pointer) {
	objc.Send[objc.ID](g.ID, objc.Sel("takeClique:traces:"), clique, traces)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceCliqueInstructionTraceTrack/traceCount
func (g GTMioTraceCliqueInstructionTraceTrack) TraceCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("traceCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceCliqueInstructionTraceTrack/traces
func (g GTMioTraceCliqueInstructionTraceTrack) Traces() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("traces"))
	return rv
}

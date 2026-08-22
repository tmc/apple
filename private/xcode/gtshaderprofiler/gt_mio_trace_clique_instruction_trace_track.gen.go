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
	rv := objc.SendIfResponds[GTMioTraceCliqueInstructionTraceTrack](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioTraceCliqueInstructionTraceTrack.PostProcess]
//   - [GTMioTraceCliqueInstructionTraceTrack.TakeCliqueTraces]
//   - [GTMioTraceCliqueInstructionTraceTrack.TraceCount]
//   - [GTMioTraceCliqueInstructionTraceTrack.Traces]
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
type IGTMioTraceCliqueInstructionTraceTrack interface {
	IGTMioTraceTrack

	// Topic: Methods

	PostProcess()
	TakeCliqueTraces(clique uint32, traces unsafe.Pointer)
	TraceCount() uint64
	Traces() *GTMioUSCInstructionTraceTrackTrace
}

// Init initializes the instance.
func (g GTMioTraceCliqueInstructionTraceTrack) Init() GTMioTraceCliqueInstructionTraceTrack {
	rv := objc.SendIfResponds[GTMioTraceCliqueInstructionTraceTrack](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioTraceCliqueInstructionTraceTrack) Autorelease() GTMioTraceCliqueInstructionTraceTrack {
	rv := objc.SendIfResponds[GTMioTraceCliqueInstructionTraceTrack](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioTraceCliqueInstructionTraceTrack creates a new GTMioTraceCliqueInstructionTraceTrack instance.
func NewGTMioTraceCliqueInstructionTraceTrack() GTMioTraceCliqueInstructionTraceTrack {
	class := getGTMioTraceCliqueInstructionTraceTrackClass()
	rv := objc.SendIfResponds[GTMioTraceCliqueInstructionTraceTrack](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTMioTraceCliqueInstructionTraceTrackWithIdScopeScopeIdentifierLevelLevelIdentifier(id int, scope uint16, identifier uint64, level uint16, identifier2 uint32) GTMioTraceCliqueInstructionTraceTrack {
	instance := getGTMioTraceCliqueInstructionTraceTrackClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithId:scope:scopeIdentifier:level:levelIdentifier:"), id, scope, identifier, level, identifier2)
	return GTMioTraceCliqueInstructionTraceTrackFromID(rv)
}

func (g GTMioTraceCliqueInstructionTraceTrack) PostProcess() {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("postProcess"))
}
func (g GTMioTraceCliqueInstructionTraceTrack) TakeCliqueTraces(clique uint32, traces unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("takeClique:traces:"), clique, traces)
}

func (g GTMioTraceCliqueInstructionTraceTrack) TraceCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("traceCount"))
	return rv
}
func (g GTMioTraceCliqueInstructionTraceTrack) Traces() *GTMioUSCInstructionTraceTrackTrace {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("traces"))
	return (*GTMioUSCInstructionTraceTrackTrace)(rv)
}

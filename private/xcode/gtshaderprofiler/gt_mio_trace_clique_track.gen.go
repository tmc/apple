// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [GTMioTraceCliqueTrack] class.
var (
	_GTMioTraceCliqueTrackClass     GTMioTraceCliqueTrackClass
	_GTMioTraceCliqueTrackClassOnce sync.Once
)

func getGTMioTraceCliqueTrackClass() GTMioTraceCliqueTrackClass {
	_GTMioTraceCliqueTrackClassOnce.Do(func() {
		_GTMioTraceCliqueTrackClass = GTMioTraceCliqueTrackClass{class: objc.GetClass("GTMioTraceCliqueTrack")}
	})
	return _GTMioTraceCliqueTrackClass
}

// GetGTMioTraceCliqueTrackClass returns the class object for GTMioTraceCliqueTrack.
func GetGTMioTraceCliqueTrackClass() GTMioTraceCliqueTrackClass {
	return getGTMioTraceCliqueTrackClass()
}

type GTMioTraceCliqueTrackClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioTraceCliqueTrackClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioTraceCliqueTrackClass) Alloc() GTMioTraceCliqueTrack {
	rv := objc.SendIfResponds[GTMioTraceCliqueTrack](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioTraceCliqueTrack.PostProcess]
//   - [GTMioTraceCliqueTrack.Take]
//   - [GTMioTraceCliqueTrack.TraceCount]
//   - [GTMioTraceCliqueTrack.Traces]
type GTMioTraceCliqueTrack struct {
	GTMioTraceTrack
}

// GTMioTraceCliqueTrackFromID constructs a [GTMioTraceCliqueTrack] from an objc.ID.
func GTMioTraceCliqueTrackFromID(id objc.ID) GTMioTraceCliqueTrack {
	return GTMioTraceCliqueTrack{GTMioTraceTrack: GTMioTraceTrackFromID(id)}
}

// Ensure GTMioTraceCliqueTrack implements IGTMioTraceCliqueTrack.
var _ IGTMioTraceCliqueTrack = GTMioTraceCliqueTrack{}

// An interface definition for the [GTMioTraceCliqueTrack] class.
//
// # Methods
//
//   - [IGTMioTraceCliqueTrack.PostProcess]
//   - [IGTMioTraceCliqueTrack.Take]
//   - [IGTMioTraceCliqueTrack.TraceCount]
//   - [IGTMioTraceCliqueTrack.Traces]
type IGTMioTraceCliqueTrack interface {
	IGTMioTraceTrack

	// Topic: Methods

	PostProcess()
	Take(take *GTMioUSCCliqueMetadata)
	TraceCount() uint64
	Traces() *GTMioUSCCliqueMetadataRef
}

// Init initializes the instance.
func (g GTMioTraceCliqueTrack) Init() GTMioTraceCliqueTrack {
	rv := objc.SendIfResponds[GTMioTraceCliqueTrack](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioTraceCliqueTrack) Autorelease() GTMioTraceCliqueTrack {
	rv := objc.SendIfResponds[GTMioTraceCliqueTrack](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioTraceCliqueTrack creates a new GTMioTraceCliqueTrack instance.
func NewGTMioTraceCliqueTrack() GTMioTraceCliqueTrack {
	class := getGTMioTraceCliqueTrackClass()
	rv := objc.SendIfResponds[GTMioTraceCliqueTrack](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTMioTraceCliqueTrackWithIdScopeScopeIdentifierLevelLevelIdentifier(id int, scope uint16, identifier uint64, level uint16, identifier2 uint32) GTMioTraceCliqueTrack {
	instance := getGTMioTraceCliqueTrackClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithId:scope:scopeIdentifier:level:levelIdentifier:"), id, scope, identifier, level, identifier2)
	return GTMioTraceCliqueTrackFromID(rv)
}

func (g GTMioTraceCliqueTrack) PostProcess() {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("postProcess"))
}
func (g GTMioTraceCliqueTrack) Take(take *GTMioUSCCliqueMetadata) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("take:"), unsafe.Pointer(take))
}

func (g GTMioTraceCliqueTrack) TraceCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("traceCount"))
	return rv
}
func (g GTMioTraceCliqueTrack) Traces() *GTMioUSCCliqueMetadataRef {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("traces"))
	return (*GTMioUSCCliqueMetadataRef)(rv)
}

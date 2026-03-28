// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [GTMioTraceShaderCliqueInstructionTraceTrackGroup] class.
var (
	_GTMioTraceShaderCliqueInstructionTraceTrackGroupClass     GTMioTraceShaderCliqueInstructionTraceTrackGroupClass
	_GTMioTraceShaderCliqueInstructionTraceTrackGroupClassOnce sync.Once
)

func getGTMioTraceShaderCliqueInstructionTraceTrackGroupClass() GTMioTraceShaderCliqueInstructionTraceTrackGroupClass {
	_GTMioTraceShaderCliqueInstructionTraceTrackGroupClassOnce.Do(func() {
		_GTMioTraceShaderCliqueInstructionTraceTrackGroupClass = GTMioTraceShaderCliqueInstructionTraceTrackGroupClass{class: objc.GetClass("GTMioTraceShaderCliqueInstructionTraceTrackGroup")}
	})
	return _GTMioTraceShaderCliqueInstructionTraceTrackGroupClass
}

// GetGTMioTraceShaderCliqueInstructionTraceTrackGroupClass returns the class object for GTMioTraceShaderCliqueInstructionTraceTrackGroup.
func GetGTMioTraceShaderCliqueInstructionTraceTrackGroupClass() GTMioTraceShaderCliqueInstructionTraceTrackGroupClass {
	return getGTMioTraceShaderCliqueInstructionTraceTrackGroupClass()
}

type GTMioTraceShaderCliqueInstructionTraceTrackGroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioTraceShaderCliqueInstructionTraceTrackGroupClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioTraceShaderCliqueInstructionTraceTrackGroupClass) Alloc() GTMioTraceShaderCliqueInstructionTraceTrackGroup {
	rv := objc.Send[GTMioTraceShaderCliqueInstructionTraceTrackGroup](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTMioTraceShaderCliqueInstructionTraceTrackGroup.EarliestTimestamp]
//   - [GTMioTraceShaderCliqueInstructionTraceTrackGroup.LatestTimestamp]
//   - [GTMioTraceShaderCliqueInstructionTraceTrackGroup.MaxCliqueId]
//   - [GTMioTraceShaderCliqueInstructionTraceTrackGroup.PipelineStateId]
//   - [GTMioTraceShaderCliqueInstructionTraceTrackGroup.ProgramType]
//   - [GTMioTraceShaderCliqueInstructionTraceTrackGroup.RecordCount]
//   - [GTMioTraceShaderCliqueInstructionTraceTrackGroup.Records]
//   - [GTMioTraceShaderCliqueInstructionTraceTrackGroup.TraceCount]
//   - [GTMioTraceShaderCliqueInstructionTraceTrackGroup.Traces]
//   - [GTMioTraceShaderCliqueInstructionTraceTrackGroup.InitWithTracesRecordsPipelineStateIdProgramTypeEarliestTimestampLatestTimestampMaxCliqueId]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceShaderCliqueInstructionTraceTrackGroup
type GTMioTraceShaderCliqueInstructionTraceTrackGroup struct {
	GTMioTraceTrack
}

// GTMioTraceShaderCliqueInstructionTraceTrackGroupFromID constructs a [GTMioTraceShaderCliqueInstructionTraceTrackGroup] from an objc.ID.
func GTMioTraceShaderCliqueInstructionTraceTrackGroupFromID(id objc.ID) GTMioTraceShaderCliqueInstructionTraceTrackGroup {
	return GTMioTraceShaderCliqueInstructionTraceTrackGroup{GTMioTraceTrack: GTMioTraceTrackFromID(id)}
}
// Ensure GTMioTraceShaderCliqueInstructionTraceTrackGroup implements IGTMioTraceShaderCliqueInstructionTraceTrackGroup.
var _ IGTMioTraceShaderCliqueInstructionTraceTrackGroup = GTMioTraceShaderCliqueInstructionTraceTrackGroup{}

// An interface definition for the [GTMioTraceShaderCliqueInstructionTraceTrackGroup] class.
//
// # Methods
//
//   - [IGTMioTraceShaderCliqueInstructionTraceTrackGroup.EarliestTimestamp]
//   - [IGTMioTraceShaderCliqueInstructionTraceTrackGroup.LatestTimestamp]
//   - [IGTMioTraceShaderCliqueInstructionTraceTrackGroup.MaxCliqueId]
//   - [IGTMioTraceShaderCliqueInstructionTraceTrackGroup.PipelineStateId]
//   - [IGTMioTraceShaderCliqueInstructionTraceTrackGroup.ProgramType]
//   - [IGTMioTraceShaderCliqueInstructionTraceTrackGroup.RecordCount]
//   - [IGTMioTraceShaderCliqueInstructionTraceTrackGroup.Records]
//   - [IGTMioTraceShaderCliqueInstructionTraceTrackGroup.TraceCount]
//   - [IGTMioTraceShaderCliqueInstructionTraceTrackGroup.Traces]
//   - [IGTMioTraceShaderCliqueInstructionTraceTrackGroup.InitWithTracesRecordsPipelineStateIdProgramTypeEarliestTimestampLatestTimestampMaxCliqueId]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceShaderCliqueInstructionTraceTrackGroup
type IGTMioTraceShaderCliqueInstructionTraceTrackGroup interface {
	IGTMioTraceTrack

	// Topic: Methods

	EarliestTimestamp() uint64
	LatestTimestamp() uint64
	MaxCliqueId() uint32
	PipelineStateId() uint64
	ProgramType() uint16
	RecordCount() uint64
	Records() unsafe.Pointer
	TraceCount() uint64
	Traces() unsafe.Pointer
	InitWithTracesRecordsPipelineStateIdProgramTypeEarliestTimestampLatestTimestampMaxCliqueId(traces unsafe.Pointer, records unsafe.Pointer, id uint64, type_ uint16, timestamp uint64, timestamp2 uint64, id2 uint32) GTMioTraceShaderCliqueInstructionTraceTrackGroup
}

// Init initializes the instance.
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) Init() GTMioTraceShaderCliqueInstructionTraceTrackGroup {
	rv := objc.Send[GTMioTraceShaderCliqueInstructionTraceTrackGroup](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) Autorelease() GTMioTraceShaderCliqueInstructionTraceTrackGroup {
	rv := objc.Send[GTMioTraceShaderCliqueInstructionTraceTrackGroup](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioTraceShaderCliqueInstructionTraceTrackGroup creates a new GTMioTraceShaderCliqueInstructionTraceTrackGroup instance.
func NewGTMioTraceShaderCliqueInstructionTraceTrackGroup() GTMioTraceShaderCliqueInstructionTraceTrackGroup {
	class := getGTMioTraceShaderCliqueInstructionTraceTrackGroupClass()
	rv := objc.Send[GTMioTraceShaderCliqueInstructionTraceTrackGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceTrack/initWithId:scope:scopeIdentifier:level:levelIdentifier:
func NewGTMioTraceShaderCliqueInstructionTraceTrackGroupWithIdScopeScopeIdentifierLevelLevelIdentifier(id int, scope uint16, identifier uint64, level uint16, identifier2 uint32) GTMioTraceShaderCliqueInstructionTraceTrackGroup {
	instance := getGTMioTraceShaderCliqueInstructionTraceTrackGroupClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithId:scope:scopeIdentifier:level:levelIdentifier:"), id, scope, identifier, level, identifier2)
	return GTMioTraceShaderCliqueInstructionTraceTrackGroupFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceShaderCliqueInstructionTraceTrackGroup/initWithTraces:records:pipelineStateId:programType:earliestTimestamp:latestTimestamp:maxCliqueId:
func NewGTMioTraceShaderCliqueInstructionTraceTrackGroupWithTracesRecordsPipelineStateIdProgramTypeEarliestTimestampLatestTimestampMaxCliqueId(traces unsafe.Pointer, records unsafe.Pointer, id uint64, type_ uint16, timestamp uint64, timestamp2 uint64, id2 uint32) GTMioTraceShaderCliqueInstructionTraceTrackGroup {
	instance := getGTMioTraceShaderCliqueInstructionTraceTrackGroupClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTraces:records:pipelineStateId:programType:earliestTimestamp:latestTimestamp:maxCliqueId:"), traces, records, id, type_, timestamp, timestamp2, id2)
	return GTMioTraceShaderCliqueInstructionTraceTrackGroupFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceShaderCliqueInstructionTraceTrackGroup/initWithTraces:records:pipelineStateId:programType:earliestTimestamp:latestTimestamp:maxCliqueId:
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) InitWithTracesRecordsPipelineStateIdProgramTypeEarliestTimestampLatestTimestampMaxCliqueId(traces unsafe.Pointer, records unsafe.Pointer, id uint64, type_ uint16, timestamp uint64, timestamp2 uint64, id2 uint32) GTMioTraceShaderCliqueInstructionTraceTrackGroup {
	rv := objc.Send[GTMioTraceShaderCliqueInstructionTraceTrackGroup](g.ID, objc.Sel("initWithTraces:records:pipelineStateId:programType:earliestTimestamp:latestTimestamp:maxCliqueId:"), traces, records, id, type_, timestamp, timestamp2, id2)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceShaderCliqueInstructionTraceTrackGroup/earliestTimestamp
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) EarliestTimestamp() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("earliestTimestamp"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceShaderCliqueInstructionTraceTrackGroup/latestTimestamp
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) LatestTimestamp() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("latestTimestamp"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceShaderCliqueInstructionTraceTrackGroup/maxCliqueId
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) MaxCliqueId() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("maxCliqueId"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceShaderCliqueInstructionTraceTrackGroup/pipelineStateId
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) PipelineStateId() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("pipelineStateId"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceShaderCliqueInstructionTraceTrackGroup/programType
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) ProgramType() uint16 {
	rv := objc.Send[uint16](g.ID, objc.Sel("programType"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceShaderCliqueInstructionTraceTrackGroup/recordCount
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) RecordCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("recordCount"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceShaderCliqueInstructionTraceTrackGroup/records
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) Records() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("records"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceShaderCliqueInstructionTraceTrackGroup/traceCount
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) TraceCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("traceCount"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceShaderCliqueInstructionTraceTrackGroup/traces
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) Traces() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("traces"))
	return rv
}


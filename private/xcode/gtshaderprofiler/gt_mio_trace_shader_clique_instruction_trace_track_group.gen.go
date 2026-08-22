// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

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
	rv := objc.SendIfResponds[GTMioTraceShaderCliqueInstructionTraceTrackGroup](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

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
type IGTMioTraceShaderCliqueInstructionTraceTrackGroup interface {
	IGTMioTraceTrack

	// Topic: Methods

	EarliestTimestamp() uint64
	LatestTimestamp() uint64
	MaxCliqueId() uint32
	PipelineStateId() uint64
	ProgramType() uint16
	RecordCount() uint64
	Records() *GTMioUSCInstructionTraceTrackRecord
	TraceCount() uint64
	Traces() *GTMioUSCInstructionTraceTrackTrace
	InitWithTracesRecordsPipelineStateIdProgramTypeEarliestTimestampLatestTimestampMaxCliqueId(traces unsafe.Pointer, records unsafe.Pointer, id uint64, type_ uint16, timestamp uint64, timestamp2 uint64, id2 uint32) GTMioTraceShaderCliqueInstructionTraceTrackGroup
}

// Init initializes the instance.
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) Init() GTMioTraceShaderCliqueInstructionTraceTrackGroup {
	rv := objc.SendIfResponds[GTMioTraceShaderCliqueInstructionTraceTrackGroup](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) Autorelease() GTMioTraceShaderCliqueInstructionTraceTrackGroup {
	rv := objc.SendIfResponds[GTMioTraceShaderCliqueInstructionTraceTrackGroup](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioTraceShaderCliqueInstructionTraceTrackGroup creates a new GTMioTraceShaderCliqueInstructionTraceTrackGroup instance.
func NewGTMioTraceShaderCliqueInstructionTraceTrackGroup() GTMioTraceShaderCliqueInstructionTraceTrackGroup {
	class := getGTMioTraceShaderCliqueInstructionTraceTrackGroupClass()
	rv := objc.SendIfResponds[GTMioTraceShaderCliqueInstructionTraceTrackGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTMioTraceShaderCliqueInstructionTraceTrackGroupWithIdScopeScopeIdentifierLevelLevelIdentifier(id int, scope uint16, identifier uint64, level uint16, identifier2 uint32) GTMioTraceShaderCliqueInstructionTraceTrackGroup {
	instance := getGTMioTraceShaderCliqueInstructionTraceTrackGroupClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithId:scope:scopeIdentifier:level:levelIdentifier:"), id, scope, identifier, level, identifier2)
	return GTMioTraceShaderCliqueInstructionTraceTrackGroupFromID(rv)
}

func NewGTMioTraceShaderCliqueInstructionTraceTrackGroupWithTracesRecordsPipelineStateIdProgramTypeEarliestTimestampLatestTimestampMaxCliqueId(traces unsafe.Pointer, records unsafe.Pointer, id uint64, type_ uint16, timestamp uint64, timestamp2 uint64, id2 uint32) GTMioTraceShaderCliqueInstructionTraceTrackGroup {
	instance := getGTMioTraceShaderCliqueInstructionTraceTrackGroupClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithTraces:records:pipelineStateId:programType:earliestTimestamp:latestTimestamp:maxCliqueId:"), traces, records, id, type_, timestamp, timestamp2, id2)
	return GTMioTraceShaderCliqueInstructionTraceTrackGroupFromID(rv)
}

func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) InitWithTracesRecordsPipelineStateIdProgramTypeEarliestTimestampLatestTimestampMaxCliqueId(traces unsafe.Pointer, records unsafe.Pointer, id uint64, type_ uint16, timestamp uint64, timestamp2 uint64, id2 uint32) GTMioTraceShaderCliqueInstructionTraceTrackGroup {
	rv := objc.SendIfResponds[GTMioTraceShaderCliqueInstructionTraceTrackGroup](g.ID, objc.Sel("initWithTraces:records:pipelineStateId:programType:earliestTimestamp:latestTimestamp:maxCliqueId:"), traces, records, id, type_, timestamp, timestamp2, id2)
	return rv
}

func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) EarliestTimestamp() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("earliestTimestamp"))
	return rv
}
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) LatestTimestamp() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("latestTimestamp"))
	return rv
}
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) MaxCliqueId() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("maxCliqueId"))
	return rv
}
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) PipelineStateId() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("pipelineStateId"))
	return rv
}
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) ProgramType() uint16 {
	rv := objc.SendIfResponds[uint16](g.ID, objc.Sel("programType"))
	return rv
}
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) RecordCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("recordCount"))
	return rv
}
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) Records() *GTMioUSCInstructionTraceTrackRecord {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("records"))
	return (*GTMioUSCInstructionTraceTrackRecord)(rv)
}
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) TraceCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("traceCount"))
	return rv
}
func (g GTMioTraceShaderCliqueInstructionTraceTrackGroup) Traces() *GTMioUSCInstructionTraceTrackTrace {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("traces"))
	return (*GTMioUSCInstructionTraceTrackTrace)(rv)
}

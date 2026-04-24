// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLDataTimelineSessionSnapshotEntry] class.
var (
	_SLDataTimelineSessionSnapshotEntryClass     SLDataTimelineSessionSnapshotEntryClass
	_SLDataTimelineSessionSnapshotEntryClassOnce sync.Once
)

func getSLDataTimelineSessionSnapshotEntryClass() SLDataTimelineSessionSnapshotEntryClass {
	_SLDataTimelineSessionSnapshotEntryClassOnce.Do(func() {
		_SLDataTimelineSessionSnapshotEntryClass = SLDataTimelineSessionSnapshotEntryClass{class: objc.GetClass("SLDataTimelineSessionSnapshotEntry")}
	})
	return _SLDataTimelineSessionSnapshotEntryClass
}

// GetSLDataTimelineSessionSnapshotEntryClass returns the class object for SLDataTimelineSessionSnapshotEntry.
func GetSLDataTimelineSessionSnapshotEntryClass() SLDataTimelineSessionSnapshotEntryClass {
	return getSLDataTimelineSessionSnapshotEntryClass()
}

type SLDataTimelineSessionSnapshotEntryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLDataTimelineSessionSnapshotEntryClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLDataTimelineSessionSnapshotEntryClass) Alloc() SLDataTimelineSessionSnapshotEntry {
	rv := objc.Send[SLDataTimelineSessionSnapshotEntry](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLDataTimelineSessionSnapshotEntry.ForegroundAppPID]
//   - [SLDataTimelineSessionSnapshotEntry.Processes]
//   - [SLDataTimelineSessionSnapshotEntry.ProcessesApplyBlock]
//   - [SLDataTimelineSessionSnapshotEntry.ProcessesArray]
//   - [SLDataTimelineSessionSnapshotEntry.SessionSnapshotIndex]
//   - [SLDataTimelineSessionSnapshotEntry.SessionSnapshotTimestamp]
//
// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionSnapshotEntry
type SLDataTimelineSessionSnapshotEntry struct {
	SLDataTimelineSessionEntry
}

// SLDataTimelineSessionSnapshotEntryFromID constructs a [SLDataTimelineSessionSnapshotEntry] from an objc.ID.
func SLDataTimelineSessionSnapshotEntryFromID(id objc.ID) SLDataTimelineSessionSnapshotEntry {
	return SLDataTimelineSessionSnapshotEntry{SLDataTimelineSessionEntry: SLDataTimelineSessionEntryFromID(id)}
}

// Ensure SLDataTimelineSessionSnapshotEntry implements ISLDataTimelineSessionSnapshotEntry.
var _ ISLDataTimelineSessionSnapshotEntry = SLDataTimelineSessionSnapshotEntry{}

// An interface definition for the [SLDataTimelineSessionSnapshotEntry] class.
//
// # Methods
//
//   - [ISLDataTimelineSessionSnapshotEntry.ForegroundAppPID]
//   - [ISLDataTimelineSessionSnapshotEntry.Processes]
//   - [ISLDataTimelineSessionSnapshotEntry.ProcessesApplyBlock]
//   - [ISLDataTimelineSessionSnapshotEntry.ProcessesArray]
//   - [ISLDataTimelineSessionSnapshotEntry.SessionSnapshotIndex]
//   - [ISLDataTimelineSessionSnapshotEntry.SessionSnapshotTimestamp]
//
// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionSnapshotEntry
type ISLDataTimelineSessionSnapshotEntry interface {
	ISLDataTimelineSessionEntry

	// Topic: Methods

	ForegroundAppPID() int
	Processes() foundation.INSArray
	ProcessesApplyBlock(block VoidHandler)
	ProcessesArray() foundation.INSArray
	SessionSnapshotIndex() uint64
	SessionSnapshotTimestamp() float64
}

// Init initializes the instance.
func (s SLDataTimelineSessionSnapshotEntry) Init() SLDataTimelineSessionSnapshotEntry {
	rv := objc.Send[SLDataTimelineSessionSnapshotEntry](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLDataTimelineSessionSnapshotEntry) Autorelease() SLDataTimelineSessionSnapshotEntry {
	rv := objc.Send[SLDataTimelineSessionSnapshotEntry](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLDataTimelineSessionSnapshotEntry creates a new SLDataTimelineSessionSnapshotEntry instance.
func NewSLDataTimelineSessionSnapshotEntry() SLDataTimelineSessionSnapshotEntry {
	class := getSLDataTimelineSessionSnapshotEntryClass()
	rv := objc.Send[SLDataTimelineSessionSnapshotEntry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionSnapshotEntry/initWithXPCObject:
func NewSLDataTimelineSessionSnapshotEntryWithXPCObject(xPCObject objectivec.IObject) SLDataTimelineSessionSnapshotEntry {
	instance := getSLDataTimelineSessionSnapshotEntryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithXPCObject:"), xPCObject)
	return SLDataTimelineSessionSnapshotEntryFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionSnapshotEntry/processesApplyBlock:
func (s SLDataTimelineSessionSnapshotEntry) ProcessesApplyBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](s.ID, objc.Sel("processesApplyBlock:"), _block0)
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionSnapshotEntry/foregroundAppPID
func (s SLDataTimelineSessionSnapshotEntry) ForegroundAppPID() int {
	rv := objc.Send[int](s.ID, objc.Sel("foregroundAppPID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionSnapshotEntry/processes
func (s SLDataTimelineSessionSnapshotEntry) Processes() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("processes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionSnapshotEntry/processesArray
func (s SLDataTimelineSessionSnapshotEntry) ProcessesArray() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("processesArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionSnapshotEntry/sessionSnapshotIndex
func (s SLDataTimelineSessionSnapshotEntry) SessionSnapshotIndex() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("sessionSnapshotIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionSnapshotEntry/sessionSnapshotTimestamp
func (s SLDataTimelineSessionSnapshotEntry) SessionSnapshotTimestamp() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("sessionSnapshotTimestamp"))
	return rv
}

// ProcessesApplyBlockSync is a synchronous wrapper around [SLDataTimelineSessionSnapshotEntry.ProcessesApplyBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLDataTimelineSessionSnapshotEntry) ProcessesApplyBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s.ProcessesApplyBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLDataTimelineServerSnapshotEntry] class.
var (
	_SLDataTimelineServerSnapshotEntryClass     SLDataTimelineServerSnapshotEntryClass
	_SLDataTimelineServerSnapshotEntryClassOnce sync.Once
)

func getSLDataTimelineServerSnapshotEntryClass() SLDataTimelineServerSnapshotEntryClass {
	_SLDataTimelineServerSnapshotEntryClassOnce.Do(func() {
		_SLDataTimelineServerSnapshotEntryClass = SLDataTimelineServerSnapshotEntryClass{class: objc.GetClass("SLDataTimelineServerSnapshotEntry")}
	})
	return _SLDataTimelineServerSnapshotEntryClass
}

// GetSLDataTimelineServerSnapshotEntryClass returns the class object for SLDataTimelineServerSnapshotEntry.
func GetSLDataTimelineServerSnapshotEntryClass() SLDataTimelineServerSnapshotEntryClass {
	return getSLDataTimelineServerSnapshotEntryClass()
}

type SLDataTimelineServerSnapshotEntryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLDataTimelineServerSnapshotEntryClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLDataTimelineServerSnapshotEntryClass) Alloc() SLDataTimelineServerSnapshotEntry {
	rv := objc.SendIfResponds[SLDataTimelineServerSnapshotEntry](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLDataTimelineServerSnapshotEntry.CreateXPCObject]
//   - [SLDataTimelineServerSnapshotEntry.Index]
//   - [SLDataTimelineServerSnapshotEntry.Sessions]
//   - [SLDataTimelineServerSnapshotEntry.SessionsApplyBlock]
//   - [SLDataTimelineServerSnapshotEntry.SessionsArray]
//   - [SLDataTimelineServerSnapshotEntry.Timestamp]
//   - [SLDataTimelineServerSnapshotEntry.InitWithXPCObject]
type SLDataTimelineServerSnapshotEntry struct {
	objectivec.Object
}

// SLDataTimelineServerSnapshotEntryFromID constructs a [SLDataTimelineServerSnapshotEntry] from an objc.ID.
func SLDataTimelineServerSnapshotEntryFromID(id objc.ID) SLDataTimelineServerSnapshotEntry {
	return SLDataTimelineServerSnapshotEntry{objectivec.Object{ID: id}}
}

// Ensure SLDataTimelineServerSnapshotEntry implements ISLDataTimelineServerSnapshotEntry.
var _ ISLDataTimelineServerSnapshotEntry = SLDataTimelineServerSnapshotEntry{}

// An interface definition for the [SLDataTimelineServerSnapshotEntry] class.
//
// # Methods
//
//   - [ISLDataTimelineServerSnapshotEntry.CreateXPCObject]
//   - [ISLDataTimelineServerSnapshotEntry.Index]
//   - [ISLDataTimelineServerSnapshotEntry.Sessions]
//   - [ISLDataTimelineServerSnapshotEntry.SessionsApplyBlock]
//   - [ISLDataTimelineServerSnapshotEntry.SessionsArray]
//   - [ISLDataTimelineServerSnapshotEntry.Timestamp]
//   - [ISLDataTimelineServerSnapshotEntry.InitWithXPCObject]
type ISLDataTimelineServerSnapshotEntry interface {
	objectivec.IObject

	// Topic: Methods

	CreateXPCObject() objectivec.IObject
	Index() uint64
	Sessions() foundation.INSArray
	SessionsApplyBlock(block VoidHandler)
	SessionsArray() foundation.INSArray
	Timestamp() float64
	InitWithXPCObject(xPCObject objectivec.IObject) SLDataTimelineServerSnapshotEntry
}

// Init initializes the instance.
func (s SLDataTimelineServerSnapshotEntry) Init() SLDataTimelineServerSnapshotEntry {
	rv := objc.SendIfResponds[SLDataTimelineServerSnapshotEntry](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLDataTimelineServerSnapshotEntry) Autorelease() SLDataTimelineServerSnapshotEntry {
	rv := objc.SendIfResponds[SLDataTimelineServerSnapshotEntry](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLDataTimelineServerSnapshotEntry creates a new SLDataTimelineServerSnapshotEntry instance.
func NewSLDataTimelineServerSnapshotEntry() SLDataTimelineServerSnapshotEntry {
	class := getSLDataTimelineServerSnapshotEntryClass()
	rv := objc.SendIfResponds[SLDataTimelineServerSnapshotEntry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLDataTimelineServerSnapshotEntryWithXPCObject(xPCObject objectivec.IObject) SLDataTimelineServerSnapshotEntry {
	instance := getSLDataTimelineServerSnapshotEntryClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithXPCObject:"), xPCObject)
	return SLDataTimelineServerSnapshotEntryFromID(rv)
}

func (s SLDataTimelineServerSnapshotEntry) CreateXPCObject() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("createXPCObject"))
	return objectivec.Object{ID: rv}
}
func (s SLDataTimelineServerSnapshotEntry) SessionsApplyBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("sessionsApplyBlock:"), _block0)
}
func (s SLDataTimelineServerSnapshotEntry) InitWithXPCObject(xPCObject objectivec.IObject) SLDataTimelineServerSnapshotEntry {
	rv := objc.SendIfResponds[SLDataTimelineServerSnapshotEntry](s.ID, objc.Sel("initWithXPCObject:"), xPCObject)
	return rv
}

func (_SLDataTimelineServerSnapshotEntryClass SLDataTimelineServerSnapshotEntryClass) EntryWithXPCObject(xPCObject objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SLDataTimelineServerSnapshotEntryClass.class), objc.Sel("entryWithXPCObject:"), xPCObject)
	return objectivec.Object{ID: rv}
}

func (s SLDataTimelineServerSnapshotEntry) Index() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("index"))
	return rv
}
func (s SLDataTimelineServerSnapshotEntry) Sessions() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("sessions"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (s SLDataTimelineServerSnapshotEntry) SessionsArray() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("sessionsArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (s SLDataTimelineServerSnapshotEntry) Timestamp() float64 {
	rv := objc.SendIfResponds[float64](s.ID, objc.Sel("timestamp"))
	return rv
}

// SessionsApplyBlockSync is a synchronous wrapper around [SLDataTimelineServerSnapshotEntry.SessionsApplyBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLDataTimelineServerSnapshotEntry) SessionsApplyBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s.SessionsApplyBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLDataTimelineSessionEntry] class.
var (
	_SLDataTimelineSessionEntryClass     SLDataTimelineSessionEntryClass
	_SLDataTimelineSessionEntryClassOnce sync.Once
)

func getSLDataTimelineSessionEntryClass() SLDataTimelineSessionEntryClass {
	_SLDataTimelineSessionEntryClassOnce.Do(func() {
		_SLDataTimelineSessionEntryClass = SLDataTimelineSessionEntryClass{class: objc.GetClass("SLDataTimelineSessionEntry")}
	})
	return _SLDataTimelineSessionEntryClass
}

// GetSLDataTimelineSessionEntryClass returns the class object for SLDataTimelineSessionEntry.
func GetSLDataTimelineSessionEntryClass() SLDataTimelineSessionEntryClass {
	return getSLDataTimelineSessionEntryClass()
}

type SLDataTimelineSessionEntryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLDataTimelineSessionEntryClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLDataTimelineSessionEntryClass) Alloc() SLDataTimelineSessionEntry {
	rv := objc.SendIfResponds[SLDataTimelineSessionEntry](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLDataTimelineSessionEntry.AuditID]
//   - [SLDataTimelineSessionEntry.CgID]
//   - [SLDataTimelineSessionEntry.CreateXPCObject]
//   - [SLDataTimelineSessionEntry.CurrentSnapshotMember]
//   - [SLDataTimelineSessionEntry.ProcessData]
//   - [SLDataTimelineSessionEntry.InitWithXPCObject]
type SLDataTimelineSessionEntry struct {
	objectivec.Object
}

// SLDataTimelineSessionEntryFromID constructs a [SLDataTimelineSessionEntry] from an objc.ID.
func SLDataTimelineSessionEntryFromID(id objc.ID) SLDataTimelineSessionEntry {
	return SLDataTimelineSessionEntry{objectivec.Object{ID: id}}
}

// Ensure SLDataTimelineSessionEntry implements ISLDataTimelineSessionEntry.
var _ ISLDataTimelineSessionEntry = SLDataTimelineSessionEntry{}

// An interface definition for the [SLDataTimelineSessionEntry] class.
//
// # Methods
//
//   - [ISLDataTimelineSessionEntry.AuditID]
//   - [ISLDataTimelineSessionEntry.CgID]
//   - [ISLDataTimelineSessionEntry.CreateXPCObject]
//   - [ISLDataTimelineSessionEntry.CurrentSnapshotMember]
//   - [ISLDataTimelineSessionEntry.ProcessData]
//   - [ISLDataTimelineSessionEntry.InitWithXPCObject]
type ISLDataTimelineSessionEntry interface {
	objectivec.IObject

	// Topic: Methods

	AuditID() int
	CgID() uint32
	CreateXPCObject() objectivec.IObject
	CurrentSnapshotMember() bool
	ProcessData() unsafe.Pointer
	InitWithXPCObject(xPCObject objectivec.IObject) SLDataTimelineSessionEntry
}

// Init initializes the instance.
func (s SLDataTimelineSessionEntry) Init() SLDataTimelineSessionEntry {
	rv := objc.SendIfResponds[SLDataTimelineSessionEntry](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLDataTimelineSessionEntry) Autorelease() SLDataTimelineSessionEntry {
	rv := objc.SendIfResponds[SLDataTimelineSessionEntry](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLDataTimelineSessionEntry creates a new SLDataTimelineSessionEntry instance.
func NewSLDataTimelineSessionEntry() SLDataTimelineSessionEntry {
	class := getSLDataTimelineSessionEntryClass()
	rv := objc.SendIfResponds[SLDataTimelineSessionEntry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLDataTimelineSessionEntryWithXPCObject(xPCObject objectivec.IObject) SLDataTimelineSessionEntry {
	instance := getSLDataTimelineSessionEntryClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithXPCObject:"), xPCObject)
	return SLDataTimelineSessionEntryFromID(rv)
}

func (s SLDataTimelineSessionEntry) CreateXPCObject() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("createXPCObject"))
	return objectivec.Object{ID: rv}
}
func (s SLDataTimelineSessionEntry) InitWithXPCObject(xPCObject objectivec.IObject) SLDataTimelineSessionEntry {
	rv := objc.SendIfResponds[SLDataTimelineSessionEntry](s.ID, objc.Sel("initWithXPCObject:"), xPCObject)
	return rv
}

func (_SLDataTimelineSessionEntryClass SLDataTimelineSessionEntryClass) EntryWithXPCObject(xPCObject objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SLDataTimelineSessionEntryClass.class), objc.Sel("entryWithXPCObject:"), xPCObject)
	return objectivec.Object{ID: rv}
}

func (s SLDataTimelineSessionEntry) AuditID() int {
	rv := objc.SendIfResponds[int](s.ID, objc.Sel("auditID"))
	return rv
}
func (s SLDataTimelineSessionEntry) CgID() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("cgID"))
	return rv
}
func (s SLDataTimelineSessionEntry) CurrentSnapshotMember() bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("currentSnapshotMember"))
	return rv
}
func (s SLDataTimelineSessionEntry) ProcessData() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](s.ID, objc.Sel("processData"))
	return rv
}

// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

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
	rv := objc.Send[SLDataTimelineSessionEntry](objc.ID(sc.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionEntry
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
//
// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionEntry
type ISLDataTimelineSessionEntry interface {
	objectivec.IObject

	// Topic: Methods

	AuditID() int
	CgID() uint32
	CreateXPCObject() objectivec.IObject
	CurrentSnapshotMember() bool
	ProcessData() objectivec.IObject
	InitWithXPCObject(xPCObject objectivec.IObject) SLDataTimelineSessionEntry
}

// Init initializes the instance.
func (s SLDataTimelineSessionEntry) Init() SLDataTimelineSessionEntry {
	rv := objc.Send[SLDataTimelineSessionEntry](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLDataTimelineSessionEntry) Autorelease() SLDataTimelineSessionEntry {
	rv := objc.Send[SLDataTimelineSessionEntry](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLDataTimelineSessionEntry creates a new SLDataTimelineSessionEntry instance.
func NewSLDataTimelineSessionEntry() SLDataTimelineSessionEntry {
	class := getSLDataTimelineSessionEntryClass()
	rv := objc.Send[SLDataTimelineSessionEntry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionEntry/initWithXPCObject:
func NewSLDataTimelineSessionEntryWithXPCObject(xPCObject objectivec.IObject) SLDataTimelineSessionEntry {
	instance := getSLDataTimelineSessionEntryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithXPCObject:"), xPCObject)
	return SLDataTimelineSessionEntryFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionEntry/createXPCObject
func (s SLDataTimelineSessionEntry) CreateXPCObject() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("createXPCObject"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionEntry/initWithXPCObject:
func (s SLDataTimelineSessionEntry) InitWithXPCObject(xPCObject objectivec.IObject) SLDataTimelineSessionEntry {
	rv := objc.Send[SLDataTimelineSessionEntry](s.ID, objc.Sel("initWithXPCObject:"), xPCObject)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionEntry/entryWithXPCObject:
func (_SLDataTimelineSessionEntryClass SLDataTimelineSessionEntryClass) EntryWithXPCObject(xPCObject objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLDataTimelineSessionEntryClass.class), objc.Sel("entryWithXPCObject:"), xPCObject)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionEntry/auditID
func (s SLDataTimelineSessionEntry) AuditID() int {
	rv := objc.Send[int](s.ID, objc.Sel("auditID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionEntry/cgID
func (s SLDataTimelineSessionEntry) CgID() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("cgID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionEntry/currentSnapshotMember
func (s SLDataTimelineSessionEntry) CurrentSnapshotMember() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("currentSnapshotMember"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionEntry/processData
func (s SLDataTimelineSessionEntry) ProcessData() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("processData"))
	return objectivec.Object{ID: rv}
}

// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLDataTimelineProcessEntry] class.
var (
	_SLDataTimelineProcessEntryClass     SLDataTimelineProcessEntryClass
	_SLDataTimelineProcessEntryClassOnce sync.Once
)

func getSLDataTimelineProcessEntryClass() SLDataTimelineProcessEntryClass {
	_SLDataTimelineProcessEntryClassOnce.Do(func() {
		_SLDataTimelineProcessEntryClass = SLDataTimelineProcessEntryClass{class: objc.GetClass("SLDataTimelineProcessEntry")}
	})
	return _SLDataTimelineProcessEntryClass
}

// GetSLDataTimelineProcessEntryClass returns the class object for SLDataTimelineProcessEntry.
func GetSLDataTimelineProcessEntryClass() SLDataTimelineProcessEntryClass {
	return getSLDataTimelineProcessEntryClass()
}

type SLDataTimelineProcessEntryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLDataTimelineProcessEntryClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLDataTimelineProcessEntryClass) Alloc() SLDataTimelineProcessEntry {
	rv := objc.Send[SLDataTimelineProcessEntry](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLDataTimelineProcessEntry.CreateXPCObject]
//   - [SLDataTimelineProcessEntry.OffScreen]
//   - [SLDataTimelineProcessEntry.OnScreenOccluded]
//   - [SLDataTimelineProcessEntry.OnScreenVisible]
//   - [SLDataTimelineProcessEntry.OrderedOut]
//   - [SLDataTimelineProcessEntry.Pid]
//   - [SLDataTimelineProcessEntry.WindowData]
//   - [SLDataTimelineProcessEntry.InitWithXPCObject]
type SLDataTimelineProcessEntry struct {
	objectivec.Object
}

// SLDataTimelineProcessEntryFromID constructs a [SLDataTimelineProcessEntry] from an objc.ID.
func SLDataTimelineProcessEntryFromID(id objc.ID) SLDataTimelineProcessEntry {
	return SLDataTimelineProcessEntry{objectivec.Object{ID: id}}
}

// Ensure SLDataTimelineProcessEntry implements ISLDataTimelineProcessEntry.
var _ ISLDataTimelineProcessEntry = SLDataTimelineProcessEntry{}

// An interface definition for the [SLDataTimelineProcessEntry] class.
//
// # Methods
//
//   - [ISLDataTimelineProcessEntry.CreateXPCObject]
//   - [ISLDataTimelineProcessEntry.OffScreen]
//   - [ISLDataTimelineProcessEntry.OnScreenOccluded]
//   - [ISLDataTimelineProcessEntry.OnScreenVisible]
//   - [ISLDataTimelineProcessEntry.OrderedOut]
//   - [ISLDataTimelineProcessEntry.Pid]
//   - [ISLDataTimelineProcessEntry.WindowData]
//   - [ISLDataTimelineProcessEntry.InitWithXPCObject]
type ISLDataTimelineProcessEntry interface {
	objectivec.IObject

	// Topic: Methods

	CreateXPCObject() objectivec.IObject
	OffScreen() uint64
	OnScreenOccluded() uint64
	OnScreenVisible() uint64
	OrderedOut() uint64
	Pid() int
	WindowData() unsafe.Pointer
	InitWithXPCObject(xPCObject objectivec.IObject) SLDataTimelineProcessEntry
}

// Init initializes the instance.
func (s SLDataTimelineProcessEntry) Init() SLDataTimelineProcessEntry {
	rv := objc.Send[SLDataTimelineProcessEntry](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLDataTimelineProcessEntry) Autorelease() SLDataTimelineProcessEntry {
	rv := objc.Send[SLDataTimelineProcessEntry](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLDataTimelineProcessEntry creates a new SLDataTimelineProcessEntry instance.
func NewSLDataTimelineProcessEntry() SLDataTimelineProcessEntry {
	class := getSLDataTimelineProcessEntryClass()
	rv := objc.Send[SLDataTimelineProcessEntry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLDataTimelineProcessEntryWithXPCObject(xPCObject objectivec.IObject) SLDataTimelineProcessEntry {
	instance := getSLDataTimelineProcessEntryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithXPCObject:"), xPCObject)
	return SLDataTimelineProcessEntryFromID(rv)
}

func (s SLDataTimelineProcessEntry) CreateXPCObject() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("createXPCObject"))
	return objectivec.Object{ID: rv}
}
func (s SLDataTimelineProcessEntry) InitWithXPCObject(xPCObject objectivec.IObject) SLDataTimelineProcessEntry {
	rv := objc.Send[SLDataTimelineProcessEntry](s.ID, objc.Sel("initWithXPCObject:"), xPCObject)
	return rv
}

func (_SLDataTimelineProcessEntryClass SLDataTimelineProcessEntryClass) EntryWithXPCObject(xPCObject objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLDataTimelineProcessEntryClass.class), objc.Sel("entryWithXPCObject:"), xPCObject)
	return objectivec.Object{ID: rv}
}

func (s SLDataTimelineProcessEntry) OffScreen() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("offScreen"))
	return rv
}
func (s SLDataTimelineProcessEntry) OnScreenOccluded() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("onScreenOccluded"))
	return rv
}
func (s SLDataTimelineProcessEntry) OnScreenVisible() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("onScreenVisible"))
	return rv
}
func (s SLDataTimelineProcessEntry) OrderedOut() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("orderedOut"))
	return rv
}
func (s SLDataTimelineProcessEntry) Pid() int {
	rv := objc.Send[int](s.ID, objc.Sel("pid"))
	return rv
}
func (s SLDataTimelineProcessEntry) WindowData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s.ID, objc.Sel("windowData"))
	return rv
}

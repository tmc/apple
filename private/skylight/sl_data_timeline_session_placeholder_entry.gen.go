// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLDataTimelineSessionPlaceholderEntry] class.
var (
	_SLDataTimelineSessionPlaceholderEntryClass     SLDataTimelineSessionPlaceholderEntryClass
	_SLDataTimelineSessionPlaceholderEntryClassOnce sync.Once
)

func getSLDataTimelineSessionPlaceholderEntryClass() SLDataTimelineSessionPlaceholderEntryClass {
	_SLDataTimelineSessionPlaceholderEntryClassOnce.Do(func() {
		_SLDataTimelineSessionPlaceholderEntryClass = SLDataTimelineSessionPlaceholderEntryClass{class: objc.GetClass("SLDataTimelineSessionPlaceholderEntry")}
	})
	return _SLDataTimelineSessionPlaceholderEntryClass
}

// GetSLDataTimelineSessionPlaceholderEntryClass returns the class object for SLDataTimelineSessionPlaceholderEntry.
func GetSLDataTimelineSessionPlaceholderEntryClass() SLDataTimelineSessionPlaceholderEntryClass {
	return getSLDataTimelineSessionPlaceholderEntryClass()
}

type SLDataTimelineSessionPlaceholderEntryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLDataTimelineSessionPlaceholderEntryClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLDataTimelineSessionPlaceholderEntryClass) Alloc() SLDataTimelineSessionPlaceholderEntry {
	rv := objc.Send[SLDataTimelineSessionPlaceholderEntry](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionPlaceholderEntry
type SLDataTimelineSessionPlaceholderEntry struct {
	SLDataTimelineSessionEntry
}

// SLDataTimelineSessionPlaceholderEntryFromID constructs a [SLDataTimelineSessionPlaceholderEntry] from an objc.ID.
func SLDataTimelineSessionPlaceholderEntryFromID(id objc.ID) SLDataTimelineSessionPlaceholderEntry {
	return SLDataTimelineSessionPlaceholderEntry{SLDataTimelineSessionEntry: SLDataTimelineSessionEntryFromID(id)}
}

// Ensure SLDataTimelineSessionPlaceholderEntry implements ISLDataTimelineSessionPlaceholderEntry.
var _ ISLDataTimelineSessionPlaceholderEntry = SLDataTimelineSessionPlaceholderEntry{}

// An interface definition for the [SLDataTimelineSessionPlaceholderEntry] class.
//
// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionPlaceholderEntry
type ISLDataTimelineSessionPlaceholderEntry interface {
	ISLDataTimelineSessionEntry
}

// Init initializes the instance.
func (s SLDataTimelineSessionPlaceholderEntry) Init() SLDataTimelineSessionPlaceholderEntry {
	rv := objc.Send[SLDataTimelineSessionPlaceholderEntry](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLDataTimelineSessionPlaceholderEntry) Autorelease() SLDataTimelineSessionPlaceholderEntry {
	rv := objc.Send[SLDataTimelineSessionPlaceholderEntry](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLDataTimelineSessionPlaceholderEntry creates a new SLDataTimelineSessionPlaceholderEntry instance.
func NewSLDataTimelineSessionPlaceholderEntry() SLDataTimelineSessionPlaceholderEntry {
	class := getSLDataTimelineSessionPlaceholderEntryClass()
	rv := objc.Send[SLDataTimelineSessionPlaceholderEntry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionPlaceholderEntry/initWithXPCObject:
func NewSLDataTimelineSessionPlaceholderEntryWithXPCObject(xPCObject objectivec.IObject) SLDataTimelineSessionPlaceholderEntry {
	instance := getSLDataTimelineSessionPlaceholderEntryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithXPCObject:"), xPCObject)
	return SLDataTimelineSessionPlaceholderEntryFromID(rv)
}

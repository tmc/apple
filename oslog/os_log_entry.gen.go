// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [OSLogEntry] class.
var (
	_OSLogEntryClass     OSLogEntryClass
	_OSLogEntryClassOnce sync.Once
)

func getOSLogEntryClass() OSLogEntryClass {
	_OSLogEntryClassOnce.Do(func() {
		_OSLogEntryClass = OSLogEntryClass{class: objc.GetClass("OSLogEntry")}
	})
	return _OSLogEntryClass
}

// GetOSLogEntryClass returns the class object for OSLogEntry.
func GetOSLogEntryClass() OSLogEntryClass {
	return getOSLogEntryClass()
}

type OSLogEntryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OSLogEntryClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OSLogEntryClass) Alloc() OSLogEntry {
	rv := objc.Send[OSLogEntry](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// A single entry from the unified logging system.
//
// # Accessing Log Entry Data
//
//   - [OSLogEntry.ComposedMessage]: The fully formatted message for the entry.
//   - [OSLogEntry.Date]: The timestamp of the entry.
//
// # Accessing Store Categories
//
//   - [OSLogEntry.StoreCategory]: The current log entry’s storage tag.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntry
type OSLogEntry struct {
	objectivec.Object
}

// OSLogEntryFromID constructs a [OSLogEntry] from an objc.ID.
//
// A single entry from the unified logging system.
func OSLogEntryFromID(id objc.ID) OSLogEntry {
	return OSLogEntry{objectivec.Object{ID: id}}
}

// NOTE: OSLogEntry adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OSLogEntry] class.
//
// # Accessing Log Entry Data
//
//   - [IOSLogEntry.ComposedMessage]: The fully formatted message for the entry.
//   - [IOSLogEntry.Date]: The timestamp of the entry.
//
// # Accessing Store Categories
//
//   - [IOSLogEntry.StoreCategory]: The current log entry’s storage tag.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntry
type IOSLogEntry interface {
	objectivec.IObject

	// Topic: Accessing Log Entry Data

	// The fully formatted message for the entry.
	ComposedMessage() string
	// The timestamp of the entry.
	Date() foundation.NSDate

	// Topic: Accessing Store Categories

	// The current log entry’s storage tag.
	StoreCategory() OSLogEntryStoreCategory

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (o OSLogEntry) Init() OSLogEntry {
	rv := objc.Send[OSLogEntry](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OSLogEntry) Autorelease() OSLogEntry {
	rv := objc.Send[OSLogEntry](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogEntry creates a new OSLogEntry instance.
func NewOSLogEntry() OSLogEntry {
	class := getOSLogEntryClass()
	rv := objc.Send[OSLogEntry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (o OSLogEntry) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The fully formatted message for the entry.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntry/composedMessage
func (o OSLogEntry) ComposedMessage() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("composedMessage"))
	return foundation.NSStringFromID(rv).String()
}

// The timestamp of the entry.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntry/date
func (o OSLogEntry) Date() foundation.NSDate {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("date"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// The current log entry’s storage tag.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntry/storeCategory-swift.property
func (o OSLogEntry) StoreCategory() OSLogEntryStoreCategory {
	rv := objc.Send[OSLogEntryStoreCategory](o.ID, objc.Sel("storeCategory"))
	return OSLogEntryStoreCategory(rv)
}

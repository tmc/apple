// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [OSLogStore] class.
var (
	_OSLogStoreClass     OSLogStoreClass
	_OSLogStoreClassOnce sync.Once
)

func getOSLogStoreClass() OSLogStoreClass {
	_OSLogStoreClassOnce.Do(func() {
		_OSLogStoreClass = OSLogStoreClass{class: objc.GetClass("OSLogStore")}
	})
	return _OSLogStoreClass
}

// GetOSLogStoreClass returns the class object for OSLogStore.
func GetOSLogStoreClass() OSLogStoreClass {
	return getOSLogStoreClass()
}

type OSLogStoreClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OSLogStoreClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OSLogStoreClass) Alloc() OSLogStore {
	rv := objc.Send[OSLogStore](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// A set of entries from the unified logging system.
//
// # Overview
//
// Instances of this class represent a fixed range of entries and may be
// backed by a `logarchive` or your Mac’s local store.
//
// In Swift, Use the [getEntries(with:at:matching:)] function to retrieve a
// filtered array of log entries.
//
// In Objective-C, use instances of this class to create [OSLogEnumerator]
// objects. One store can support multiple [OSLogEnumerator] instances
// concurrently.
//
// # Accessing Position
//
//   - [OSLogStore.PositionWithDate]: Returns a position representing the time specified.
//   - [OSLogStore.PositionWithTimeIntervalSinceEnd]: Returns a position representing time since the end of the time range that the entries span.
//   - [OSLogStore.PositionWithTimeIntervalSinceLatestBoot]: Returns a position representing time since the last boot in the series of entries.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogStore
//
// [getEntries(with:at:matching:)]: https://developer.apple.com/documentation/OSLog/OSLogStore/getEntries(with:at:matching:)
type OSLogStore struct {
	objectivec.Object
}

// OSLogStoreFromID constructs a [OSLogStore] from an objc.ID.
//
// A set of entries from the unified logging system.
func OSLogStoreFromID(id objc.ID) OSLogStore {
	return OSLogStore{objectivec.Object{ID: id}}
}

// NOTE: OSLogStore adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OSLogStore] class.
//
// # Accessing Position
//
//   - [IOSLogStore.PositionWithDate]: Returns a position representing the time specified.
//   - [IOSLogStore.PositionWithTimeIntervalSinceEnd]: Returns a position representing time since the end of the time range that the entries span.
//   - [IOSLogStore.PositionWithTimeIntervalSinceLatestBoot]: Returns a position representing time since the last boot in the series of entries.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogStore
type IOSLogStore interface {
	objectivec.IObject

	// Topic: Accessing Position

	// Returns a position representing the time specified.
	PositionWithDate(date foundation.NSDate) IOSLogPosition
	// Returns a position representing time since the end of the time range that the entries span.
	PositionWithTimeIntervalSinceEnd(seconds foundation.NSTimeInterval) IOSLogPosition
	// Returns a position representing time since the last boot in the series of entries.
	PositionWithTimeIntervalSinceLatestBoot(seconds foundation.NSTimeInterval) IOSLogPosition

	// Returns a log enumerator with default options for viewing the entries.
	EntriesEnumeratorAndReturnError() (IOSLogEnumerator, error)
	// Returns a log enumerator based on an underlying store.
	EntriesEnumeratorWithOptionsPositionPredicateError(options OSLogEnumeratorOptions, position IOSLogPosition, predicate foundation.NSPredicate) (IOSLogEnumerator, error)
}

// Init initializes the instance.
func (o OSLogStore) Init() OSLogStore {
	rv := objc.Send[OSLogStore](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OSLogStore) Autorelease() OSLogStore {
	rv := objc.Send[OSLogStore](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogStore creates a new OSLogStore instance.
func NewOSLogStore() OSLogStore {
	class := getOSLogStoreClass()
	rv := objc.Send[OSLogStore](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/OSLog/OSLogStore/init(scope:)
func NewOSLogStoreWithScopeError(scope OSLogStoreScope) (OSLogStore, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getOSLogStoreClass().class), objc.Sel("storeWithScope:error:"), scope, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return OSLogStore{}, foundation.NSErrorFrom(errorPtr)
	}
	return OSLogStoreFromID(rv), nil
}

// Creates a log store based on a log archive.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogStore/init(url:)
func NewOSLogStoreWithURLError(url foundation.NSURL) (OSLogStore, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getOSLogStoreClass().class), objc.Sel("storeWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return OSLogStore{}, foundation.NSErrorFrom(errorPtr)
	}
	return OSLogStoreFromID(rv), nil
}

// Returns a position representing the time specified.
//
// # Discussion
//
// If there are multiple occurrences of the same time, the method returns the
// earliest occurrence.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogStore/position(date:)
func (o OSLogStore) PositionWithDate(date foundation.NSDate) IOSLogPosition {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("positionWithDate:"), date)
	return OSLogPositionFromID(rv)
}

// Returns a position representing time since the end of the time range that
// the entries span.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogStore/position(timeIntervalSinceEnd:)
func (o OSLogStore) PositionWithTimeIntervalSinceEnd(seconds foundation.NSTimeInterval) IOSLogPosition {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("positionWithTimeIntervalSinceEnd:"), seconds)
	return OSLogPositionFromID(rv)
}

// Returns a position representing time since the last boot in the series of
// entries.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogStore/position(timeIntervalSinceLatestBoot:)
func (o OSLogStore) PositionWithTimeIntervalSinceLatestBoot(seconds foundation.NSTimeInterval) IOSLogPosition {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("positionWithTimeIntervalSinceLatestBoot:"), seconds)
	return OSLogPositionFromID(rv)
}

// Returns a log enumerator with default options for viewing the entries.
//
// # Discussion
//
// This method returns all of the entries from earliest to latest. If the
// enumerator can’t be set up, `entriesEnumeratorAndReturnError` returns
// `nil` and sets the error parameter to an error object describing the
// problem.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogStore/entriesEnumeratorAndReturnError:
func (o OSLogStore) EntriesEnumeratorAndReturnError() (IOSLogEnumerator, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("entriesEnumeratorAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return OSLogEnumerator{}, foundation.NSErrorFrom(errorPtr)
	}
	return OSLogEnumeratorFromID(rv), nil

}

// Returns a log enumerator based on an underlying store.
//
// # Discussion
//
// The returned object represents the sequence of entries for the
// [OSLogStore]. Use the additional parameters to control which entries are
// returned and their order.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogStore/entriesEnumeratorWithOptions:position:predicate:error:
func (o OSLogStore) EntriesEnumeratorWithOptionsPositionPredicateError(options OSLogEnumeratorOptions, position IOSLogPosition, predicate foundation.NSPredicate) (IOSLogEnumerator, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("entriesEnumeratorWithOptions:position:predicate:error:"), options, position, predicate, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return OSLogEnumerator{}, foundation.NSErrorFrom(errorPtr)
	}
	return OSLogEnumeratorFromID(rv), nil

}

// Creates a log store representing the Mac’s local store.
//
// # Discussion
//
// Gaining access to the local unified logging system requires permission from
// the system. The caller must be run by an admin account and have the
// `com.AppleXCUIElementTypeLoggingXCUIElementTypeLocal()-store` entitlement.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogStore/local()
func (_OSLogStoreClass OSLogStoreClass) LocalStoreAndReturnError() (OSLogStore, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_OSLogStoreClass.class), objc.Sel("localStoreAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return OSLogStore{}, foundation.NSErrorFrom(errorPtr)
	}
	return OSLogStoreFromID(rv), nil

}

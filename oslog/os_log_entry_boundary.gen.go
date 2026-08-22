// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [OSLogEntryBoundary] class.
var (
	_OSLogEntryBoundaryClass     OSLogEntryBoundaryClass
	_OSLogEntryBoundaryClassOnce sync.Once
)

func getOSLogEntryBoundaryClass() OSLogEntryBoundaryClass {
	_OSLogEntryBoundaryClassOnce.Do(func() {
		_OSLogEntryBoundaryClass = OSLogEntryBoundaryClass{class: objc.GetClass("OSLogEntryBoundary")}
	})
	return _OSLogEntryBoundaryClass
}

// GetOSLogEntryBoundaryClass returns the class object for OSLogEntryBoundary.
func GetOSLogEntryBoundaryClass() OSLogEntryBoundaryClass {
	return getOSLogEntryBoundaryClass()
}

type OSLogEntryBoundaryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OSLogEntryBoundaryClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OSLogEntryBoundaryClass) Alloc() OSLogEntryBoundary {
	rv := objc.Send[OSLogEntryBoundary](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// The metadata that partitions sequences of other entries.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryBoundary
type OSLogEntryBoundary struct {
	OSLogEntry
}

// OSLogEntryBoundaryFromID constructs a [OSLogEntryBoundary] from an objc.ID.
//
// The metadata that partitions sequences of other entries.
func OSLogEntryBoundaryFromID(id objc.ID) OSLogEntryBoundary {
	return OSLogEntryBoundary{OSLogEntry: OSLogEntryFromID(id)}
}

// NOTE: OSLogEntryBoundary adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OSLogEntryBoundary] class.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryBoundary
type IOSLogEntryBoundary interface {
	IOSLogEntry
}

// Init initializes the instance.
func (o OSLogEntryBoundary) Init() OSLogEntryBoundary {
	rv := objc.Send[OSLogEntryBoundary](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OSLogEntryBoundary) Autorelease() OSLogEntryBoundary {
	rv := objc.Send[OSLogEntryBoundary](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogEntryBoundary creates a new OSLogEntryBoundary instance.
func NewOSLogEntryBoundary() OSLogEntryBoundary {
	class := getOSLogEntryBoundaryClass()
	rv := objc.Send[OSLogEntryBoundary](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [OSLogPosition] class.
var (
	_OSLogPositionClass     OSLogPositionClass
	_OSLogPositionClassOnce sync.Once
)

func getOSLogPositionClass() OSLogPositionClass {
	_OSLogPositionClassOnce.Do(func() {
		_OSLogPositionClass = OSLogPositionClass{class: objc.GetClass("OSLogPosition")}
	})
	return _OSLogPositionClass
}

// GetOSLogPositionClass returns the class object for OSLogPosition.
func GetOSLogPositionClass() OSLogPositionClass {
	return getOSLogPositionClass()
}

type OSLogPositionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OSLogPositionClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OSLogPositionClass) Alloc() OSLogPosition {
	rv := objc.Send[OSLogPosition](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a point in a sequence of entries in the unified logging
// system.
//
// # Overview
//
// Generate positions with [OSLogStore] instance methods and use them to view
// entries from a particular starting point.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogPosition
type OSLogPosition struct {
	objectivec.Object
}

// OSLogPositionFromID constructs a [OSLogPosition] from an objc.ID.
//
// A representation of a point in a sequence of entries in the unified logging
// system.
func OSLogPositionFromID(id objc.ID) OSLogPosition {
	return OSLogPosition{objectivec.Object{ID: id}}
}

// NOTE: OSLogPosition adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OSLogPosition] class.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogPosition
type IOSLogPosition interface {
	objectivec.IObject
}

// Init initializes the instance.
func (o OSLogPosition) Init() OSLogPosition {
	rv := objc.Send[OSLogPosition](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OSLogPosition) Autorelease() OSLogPosition {
	rv := objc.Send[OSLogPosition](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogPosition creates a new OSLogPosition instance.
func NewOSLogPosition() OSLogPosition {
	class := getOSLogPositionClass()
	rv := objc.Send[OSLogPosition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [OSLogEnumerator] class.
var (
	_OSLogEnumeratorClass     OSLogEnumeratorClass
	_OSLogEnumeratorClassOnce sync.Once
)

func getOSLogEnumeratorClass() OSLogEnumeratorClass {
	_OSLogEnumeratorClassOnce.Do(func() {
		_OSLogEnumeratorClass = OSLogEnumeratorClass{class: objc.GetClass("OSLogEnumerator")}
	})
	return _OSLogEnumeratorClass
}

// GetOSLogEnumeratorClass returns the class object for OSLogEnumerator.
func GetOSLogEnumeratorClass() OSLogEnumeratorClass {
	return getOSLogEnumeratorClass()
}

type OSLogEnumeratorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OSLogEnumeratorClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OSLogEnumeratorClass) Alloc() OSLogEnumerator {
	rv := objc.Send[OSLogEnumerator](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// An enumerator that can access and list log entries.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEnumerator
type OSLogEnumerator struct {
	foundation.NSEnumerator
}

// OSLogEnumeratorFromID constructs a [OSLogEnumerator] from an objc.ID.
//
// An enumerator that can access and list log entries.
func OSLogEnumeratorFromID(id objc.ID) OSLogEnumerator {
	return OSLogEnumerator{NSEnumerator: foundation.NSEnumeratorFromID(id)}
}

// NOTE: OSLogEnumerator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OSLogEnumerator] class.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEnumerator
type IOSLogEnumerator interface {
	foundation.INSEnumerator
}

// Init initializes the instance.
func (o OSLogEnumerator) Init() OSLogEnumerator {
	rv := objc.Send[OSLogEnumerator](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OSLogEnumerator) Autorelease() OSLogEnumerator {
	rv := objc.Send[OSLogEnumerator](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogEnumerator creates a new OSLogEnumerator instance.
func NewOSLogEnumerator() OSLogEnumerator {
	class := getOSLogEnumeratorClass()
	rv := objc.Send[OSLogEnumerator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

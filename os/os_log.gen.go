// Code generated from Apple documentation for os. DO NOT EDIT.

package os

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [OSLog] class.
var (
	_OSLogClass     OSLogClass
	_OSLogClassOnce sync.Once
)

func getOSLogClass() OSLogClass {
	_OSLogClassOnce.Do(func() {
		_OSLogClass = OSLogClass{class: objc.GetClass("OSLog")}
	})
	return _OSLogClass
}

// GetOSLogClass returns the class object for OSLog.
func GetOSLogClass() OSLogClass {
	return getOSLogClass()
}

type OSLogClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OSLogClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OSLogClass) Alloc() OSLog {
	rv := objc.Send[OSLog](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// A container of related log messages.
//
// # Overview
//
// A log categorizes the messages you write and makes it easy to sort and
// filter them. Each log contains a subsystem and a category, which you
// define. A subsystem identifies a major functional area of your app, which
// you specify using reverse DNS notation, such as
// `com.Your_companyXCUIElementTypeYour_subsystem_name()`. A category
// segregates specific areas within a subsystem.
//
// See: https://developer.apple.com/documentation/os/OSLog
type OSLog struct {
	objectivec.Object
}

// OSLogFromID constructs a [OSLog] from an objc.ID.
//
// A container of related log messages.
func OSLogFromID(id objc.ID) OSLog {
	return OSLog{objectivec.Object{ID: id}}
}

// NOTE: OSLog adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OSLog] class.
//
// See: https://developer.apple.com/documentation/os/OSLog
type IOSLog interface {
	objectivec.IObject
}

// Init initializes the instance.
func (o OSLog) Init() OSLog {
	rv := objc.Send[OSLog](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OSLog) Autorelease() OSLog {
	rv := objc.Send[OSLog](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLog creates a new OSLog instance.
func NewOSLog() OSLog {
	class := getOSLogClass()
	rv := objc.Send[OSLog](objc.ID(class.class), objc.Sel("new"))
	return rv
}

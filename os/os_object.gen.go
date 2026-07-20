// Code generated from Apple documentation for os. DO NOT EDIT.

package os

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [OS_object] class.
var (
	_OS_objectClass     OS_objectClass
	_OS_objectClassOnce sync.Once
)

func getOS_objectClass() OS_objectClass {
	_OS_objectClassOnce.Do(func() {
		_OS_objectClass = OS_objectClass{class: objc.GetClass("OS_object")}
	})
	return _OS_objectClass
}

// GetOS_objectClass returns the class object for OS_object.
func GetOS_objectClass() OS_objectClass {
	return getOS_objectClass()
}

type OS_objectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OS_objectClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OS_objectClass) Alloc() OS_object {
	rv := objc.Send[OS_object](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/os/OS_object
type OS_object struct {
	objectivec.Object
}

// OS_objectFromID constructs a [OS_object] from an objc.ID.
func OS_objectFromID(id objc.ID) OS_object {
	return OS_object{objectivec.Object{ID: id}}
}

// Ensure OS_object implements IOS_object.
var _ IOS_object = OS_object{}

// An interface definition for the [OS_object] class.
//
// See: https://developer.apple.com/documentation/os/OS_object
type IOS_object interface {
	objectivec.IObject
}

// Init initializes the instance.
func (o OS_object) Init() OS_object {
	rv := objc.Send[OS_object](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OS_object) Autorelease() OS_object {
	rv := objc.Send[OS_object](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOS_object creates a new OS_object instance.
func NewOS_object() OS_object {
	class := getOS_objectClass()
	rv := objc.Send[OS_object](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [HostName] class.
var (
	_HostNameClass     HostNameClass
	_HostNameClassOnce sync.Once
)

func getHostNameClass() HostNameClass {
	_HostNameClassOnce.Do(func() {
		_HostNameClass = HostNameClass{class: objc.GetClass("hostName")}
	})
	return _HostNameClass
}

// GetHostNameClass returns the class object for hostName.
func GetHostNameClass() HostNameClass {
	return getHostNameClass()
}

type HostNameClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (hc HostNameClass) Alloc() HostName {
	rv := objc.Send[HostName](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSProcessInfo/hostName-c.ivar
type HostName struct {
	objectivec.Object
}

// HostNameFromID constructs a [HostName] from an objc.ID.
func HostNameFromID(id objc.ID) HostName {
	return HostName{objectivec.Object{ID: id}}
}
// Ensure HostName implements IHostName.
var _ IHostName = HostName{}

// An interface definition for the [HostName] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSProcessInfo/hostName-c.ivar
type IHostName interface {
	objectivec.IObject
}

// Init initializes the instance.
func (h HostName) Init() HostName {
	rv := objc.Send[HostName](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h HostName) Autorelease() HostName {
	rv := objc.Send[HostName](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewHostName creates a new HostName instance.
func NewHostName() HostName {
	class := getHostNameClass()
	rv := objc.Send[HostName](objc.ID(class.class), objc.Sel("new"))
	return rv
}


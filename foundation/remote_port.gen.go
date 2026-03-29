// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [RemotePort] class.
var (
	_RemotePortClass     RemotePortClass
	_RemotePortClassOnce sync.Once
)

func getRemotePortClass() RemotePortClass {
	_RemotePortClassOnce.Do(func() {
		_RemotePortClass = RemotePortClass{class: objc.GetClass("remotePort")}
	})
	return _RemotePortClass
}

// GetRemotePortClass returns the class object for remotePort.
func GetRemotePortClass() RemotePortClass {
	return getRemotePortClass()
}

type RemotePortClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (rc RemotePortClass) Class() objc.Class {
	return rc.class
}

// Alloc allocates memory for a new instance of the class.
func (rc RemotePortClass) Alloc() RemotePort {
	rv := objc.Send[RemotePort](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSPortMessage/remotePort
type RemotePort struct {
	objectivec.Object
}

// RemotePortFromID constructs a [RemotePort] from an objc.ID.
func RemotePortFromID(id objc.ID) RemotePort {
	return RemotePort{objectivec.Object{ID: id}}
}
// Ensure RemotePort implements IRemotePort.
var _ IRemotePort = RemotePort{}

// An interface definition for the [RemotePort] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSPortMessage/remotePort
type IRemotePort interface {
	objectivec.IObject
}

// Init initializes the instance.
func (r RemotePort) Init() RemotePort {
	rv := objc.Send[RemotePort](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r RemotePort) Autorelease() RemotePort {
	rv := objc.Send[RemotePort](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewRemotePort creates a new RemotePort instance.
func NewRemotePort() RemotePort {
	class := getRemotePortClass()
	rv := objc.Send[RemotePort](objc.ID(class.class), objc.Sel("new"))
	return rv
}


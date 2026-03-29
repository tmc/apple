// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [LocalPort] class.
var (
	_LocalPortClass     LocalPortClass
	_LocalPortClassOnce sync.Once
)

func getLocalPortClass() LocalPortClass {
	_LocalPortClassOnce.Do(func() {
		_LocalPortClass = LocalPortClass{class: objc.GetClass("localPort")}
	})
	return _LocalPortClass
}

// GetLocalPortClass returns the class object for localPort.
func GetLocalPortClass() LocalPortClass {
	return getLocalPortClass()
}

type LocalPortClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LocalPortClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LocalPortClass) Alloc() LocalPort {
	rv := objc.Send[LocalPort](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSPortMessage/localPort
type LocalPort struct {
	objectivec.Object
}

// LocalPortFromID constructs a [LocalPort] from an objc.ID.
func LocalPortFromID(id objc.ID) LocalPort {
	return LocalPort{objectivec.Object{ID: id}}
}
// Ensure LocalPort implements ILocalPort.
var _ ILocalPort = LocalPort{}

// An interface definition for the [LocalPort] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSPortMessage/localPort
type ILocalPort interface {
	objectivec.IObject
}

// Init initializes the instance.
func (l LocalPort) Init() LocalPort {
	rv := objc.Send[LocalPort](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l LocalPort) Autorelease() LocalPort {
	rv := objc.Send[LocalPort](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewLocalPort creates a new LocalPort instance.
func NewLocalPort() LocalPort {
	class := getLocalPortClass()
	rv := objc.Send[LocalPort](objc.ID(class.class), objc.Sel("new"))
	return rv
}


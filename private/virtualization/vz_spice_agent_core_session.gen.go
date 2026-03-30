// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSpiceAgentCoreSession] class.
var (
	_VZSpiceAgentCoreSessionClass     VZSpiceAgentCoreSessionClass
	_VZSpiceAgentCoreSessionClassOnce sync.Once
)

func getVZSpiceAgentCoreSessionClass() VZSpiceAgentCoreSessionClass {
	_VZSpiceAgentCoreSessionClassOnce.Do(func() {
		_VZSpiceAgentCoreSessionClass = VZSpiceAgentCoreSessionClass{class: objc.GetClass("_VZSpiceAgentCoreSession")}
	})
	return _VZSpiceAgentCoreSessionClass
}

// GetVZSpiceAgentCoreSessionClass returns the class object for _VZSpiceAgentCoreSession.
func GetVZSpiceAgentCoreSessionClass() VZSpiceAgentCoreSessionClass {
	return getVZSpiceAgentCoreSessionClass()
}

type VZSpiceAgentCoreSessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSpiceAgentCoreSessionClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSpiceAgentCoreSessionClass) Alloc() VZSpiceAgentCoreSession {
	rv := objc.Send[VZSpiceAgentCoreSession](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCoreSession
type VZSpiceAgentCoreSession struct {
	objectivec.Object
}

// VZSpiceAgentCoreSessionFromID constructs a [VZSpiceAgentCoreSession] from an objc.ID.
func VZSpiceAgentCoreSessionFromID(id objc.ID) VZSpiceAgentCoreSession {
	return VZSpiceAgentCoreSession{objectivec.Object{ID: id}}
}

// Ensure VZSpiceAgentCoreSession implements IVZSpiceAgentCoreSession.
var _ IVZSpiceAgentCoreSession = VZSpiceAgentCoreSession{}

// An interface definition for the [VZSpiceAgentCoreSession] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCoreSession
type IVZSpiceAgentCoreSession interface {
	objectivec.IObject
}

// Init initializes the instance.
func (v VZSpiceAgentCoreSession) Init() VZSpiceAgentCoreSession {
	rv := objc.Send[VZSpiceAgentCoreSession](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZSpiceAgentCoreSession) Autorelease() VZSpiceAgentCoreSession {
	rv := objc.Send[VZSpiceAgentCoreSession](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSpiceAgentCoreSession creates a new VZSpiceAgentCoreSession instance.
func NewVZSpiceAgentCoreSession() VZSpiceAgentCoreSession {
	class := getVZSpiceAgentCoreSessionClass()
	rv := objc.Send[VZSpiceAgentCoreSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

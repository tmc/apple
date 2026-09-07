// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PKGCoreWindowInteractionSession] class.
var (
	_PKGCoreWindowInteractionSessionClass     PKGCoreWindowInteractionSessionClass
	_PKGCoreWindowInteractionSessionClassOnce sync.Once
)

func getPKGCoreWindowInteractionSessionClass() PKGCoreWindowInteractionSessionClass {
	_PKGCoreWindowInteractionSessionClassOnce.Do(func() {
		_PKGCoreWindowInteractionSessionClass = PKGCoreWindowInteractionSessionClass{class: objc.GetClass("PKGCore.WindowInteractionSession")}
	})
	return _PKGCoreWindowInteractionSessionClass
}

// GetPKGCoreWindowInteractionSessionClass returns the class object for PKGCore.WindowInteractionSession.
func GetPKGCoreWindowInteractionSessionClass() PKGCoreWindowInteractionSessionClass {
	return getPKGCoreWindowInteractionSessionClass()
}

type PKGCoreWindowInteractionSessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PKGCoreWindowInteractionSessionClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PKGCoreWindowInteractionSessionClass) Alloc() PKGCoreWindowInteractionSession {
	rv := objc.SendIfResponds[PKGCoreWindowInteractionSession](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

type PKGCoreWindowInteractionSession struct {
	objectivec.Object
}

// PKGCoreWindowInteractionSessionFromID constructs a [PKGCoreWindowInteractionSession] from an objc.ID.
func PKGCoreWindowInteractionSessionFromID(id objc.ID) PKGCoreWindowInteractionSession {
	return PKGCoreWindowInteractionSession{objectivec.Object{ID: id}}
}

// Ensure PKGCoreWindowInteractionSession implements IPKGCoreWindowInteractionSession.
var _ IPKGCoreWindowInteractionSession = PKGCoreWindowInteractionSession{}

// An interface definition for the [PKGCoreWindowInteractionSession] class.
type IPKGCoreWindowInteractionSession interface {
	objectivec.IObject
}

// Init initializes the instance.
func (p PKGCoreWindowInteractionSession) Init() PKGCoreWindowInteractionSession {
	rv := objc.SendIfResponds[PKGCoreWindowInteractionSession](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PKGCoreWindowInteractionSession) Autorelease() PKGCoreWindowInteractionSession {
	rv := objc.SendIfResponds[PKGCoreWindowInteractionSession](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPKGCoreWindowInteractionSession creates a new PKGCoreWindowInteractionSession instance.
func NewPKGCoreWindowInteractionSession() PKGCoreWindowInteractionSession {
	class := getPKGCoreWindowInteractionSessionClass()
	rv := objc.SendIfResponds[PKGCoreWindowInteractionSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

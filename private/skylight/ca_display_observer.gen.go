// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CADisplayObserver] class.
var (
	_CADisplayObserverClass     CADisplayObserverClass
	_CADisplayObserverClassOnce sync.Once
)

func getCADisplayObserverClass() CADisplayObserverClass {
	_CADisplayObserverClassOnce.Do(func() {
		_CADisplayObserverClass = CADisplayObserverClass{class: objc.GetClass("CADisplayObserver")}
	})
	return _CADisplayObserverClass
}

// GetCADisplayObserverClass returns the class object for CADisplayObserver.
func GetCADisplayObserverClass() CADisplayObserverClass {
	return getCADisplayObserverClass()
}

type CADisplayObserverClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CADisplayObserverClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CADisplayObserverClass) Alloc() CADisplayObserver {
	rv := objc.SendIfResponds[CADisplayObserver](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

type CADisplayObserver struct {
	objectivec.Object
}

// CADisplayObserverFromID constructs a [CADisplayObserver] from an objc.ID.
func CADisplayObserverFromID(id objc.ID) CADisplayObserver {
	return CADisplayObserver{objectivec.Object{ID: id}}
}

// Ensure CADisplayObserver implements ICADisplayObserver.
var _ ICADisplayObserver = CADisplayObserver{}

// An interface definition for the [CADisplayObserver] class.
type ICADisplayObserver interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CADisplayObserver) Init() CADisplayObserver {
	rv := objc.SendIfResponds[CADisplayObserver](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CADisplayObserver) Autorelease() CADisplayObserver {
	rv := objc.SendIfResponds[CADisplayObserver](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCADisplayObserver creates a new CADisplayObserver instance.
func NewCADisplayObserver() CADisplayObserver {
	class := getCADisplayObserverClass()
	rv := objc.SendIfResponds[CADisplayObserver](objc.ID(class.class), objc.Sel("new"))
	return rv
}

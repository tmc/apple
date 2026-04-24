// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [LPMObserver] class.
var (
	_LPMObserverClass     LPMObserverClass
	_LPMObserverClassOnce sync.Once
)

func getLPMObserverClass() LPMObserverClass {
	_LPMObserverClassOnce.Do(func() {
		_LPMObserverClass = LPMObserverClass{class: objc.GetClass("LPMObserver")}
	})
	return _LPMObserverClass
}

// GetLPMObserverClass returns the class object for LPMObserver.
func GetLPMObserverClass() LPMObserverClass {
	return getLPMObserverClass()
}

type LPMObserverClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LPMObserverClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LPMObserverClass) Alloc() LPMObserver {
	rv := objc.Send[LPMObserver](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [LPMObserver.HandlePowerStateChange]
//   - [LPMObserver.IsLowPowerModeEnabled]
//
// See: https://developer.apple.com/documentation/SkyLight/LPMObserver
type LPMObserver struct {
	objectivec.Object
}

// LPMObserverFromID constructs a [LPMObserver] from an objc.ID.
func LPMObserverFromID(id objc.ID) LPMObserver {
	return LPMObserver{objectivec.Object{ID: id}}
}

// Ensure LPMObserver implements ILPMObserver.
var _ ILPMObserver = LPMObserver{}

// An interface definition for the [LPMObserver] class.
//
// # Methods
//
//   - [ILPMObserver.HandlePowerStateChange]
//   - [ILPMObserver.IsLowPowerModeEnabled]
//
// See: https://developer.apple.com/documentation/SkyLight/LPMObserver
type ILPMObserver interface {
	objectivec.IObject

	// Topic: Methods

	HandlePowerStateChange(change objectivec.IObject)
	IsLowPowerModeEnabled() bool
}

// Init initializes the instance.
func (l LPMObserver) Init() LPMObserver {
	rv := objc.Send[LPMObserver](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l LPMObserver) Autorelease() LPMObserver {
	rv := objc.Send[LPMObserver](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewLPMObserver creates a new LPMObserver instance.
func NewLPMObserver() LPMObserver {
	class := getLPMObserverClass()
	rv := objc.Send[LPMObserver](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/LPMObserver/handlePowerStateChange:
func (l LPMObserver) HandlePowerStateChange(change objectivec.IObject) {
	objc.Send[objc.ID](l.ID, objc.Sel("handlePowerStateChange:"), change)
}

// See: https://developer.apple.com/documentation/SkyLight/LPMObserver/sharedLPMObserver
func (_LPMObserverClass LPMObserverClass) SharedLPMObserver() LPMObserver {
	rv := objc.Send[objc.ID](objc.ID(_LPMObserverClass.class), objc.Sel("sharedLPMObserver"))
	return LPMObserverFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/LPMObserver/isLowPowerModeEnabled
func (l LPMObserver) IsLowPowerModeEnabled() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("isLowPowerModeEnabled"))
	return rv
}

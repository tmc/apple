// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SOSystemBehaviorManager] class.
var (
	_SOSystemBehaviorManagerClass     SOSystemBehaviorManagerClass
	_SOSystemBehaviorManagerClassOnce sync.Once
)

func getSOSystemBehaviorManagerClass() SOSystemBehaviorManagerClass {
	_SOSystemBehaviorManagerClassOnce.Do(func() {
		_SOSystemBehaviorManagerClass = SOSystemBehaviorManagerClass{class: objc.GetClass("SOSystemBehaviorManager")}
	})
	return _SOSystemBehaviorManagerClass
}

// GetSOSystemBehaviorManagerClass returns the class object for SOSystemBehaviorManager.
func GetSOSystemBehaviorManagerClass() SOSystemBehaviorManagerClass {
	return getSOSystemBehaviorManagerClass()
}

type SOSystemBehaviorManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SOSystemBehaviorManagerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SOSystemBehaviorManagerClass) Alloc() SOSystemBehaviorManager {
	rv := objc.Send[SOSystemBehaviorManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [SOSystemBehaviorManager._numberObjectFromTimer]
//   - [SOSystemBehaviorManager.AddTimer]
//   - [SOSystemBehaviorManager.RemoveTimer]
// See: https://developer.apple.com/documentation/SpeechObjects/SOSystemBehaviorManager
type SOSystemBehaviorManager struct {
	objectivec.Object
}

// SOSystemBehaviorManagerFromID constructs a [SOSystemBehaviorManager] from an objc.ID.
func SOSystemBehaviorManagerFromID(id objc.ID) SOSystemBehaviorManager {
	return SOSystemBehaviorManager{objectivec.Object{ID: id}}
}
// Ensure SOSystemBehaviorManager implements ISOSystemBehaviorManager.
var _ ISOSystemBehaviorManager = SOSystemBehaviorManager{}

// An interface definition for the [SOSystemBehaviorManager] class.
//
// # Methods
//
//   - [ISOSystemBehaviorManager._numberObjectFromTimer]
//   - [ISOSystemBehaviorManager.AddTimer]
//   - [ISOSystemBehaviorManager.RemoveTimer]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOSystemBehaviorManager
type ISOSystemBehaviorManager interface {
	objectivec.IObject

	// Topic: Methods

	_numberObjectFromTimer(timer objectivec.IObject) objectivec.IObject
	AddTimer(timer objectivec.IObject)
	RemoveTimer(timer objectivec.IObject)
}

// Init initializes the instance.
func (s SOSystemBehaviorManager) Init() SOSystemBehaviorManager {
	rv := objc.Send[SOSystemBehaviorManager](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOSystemBehaviorManager) Autorelease() SOSystemBehaviorManager {
	rv := objc.Send[SOSystemBehaviorManager](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOSystemBehaviorManager creates a new SOSystemBehaviorManager instance.
func NewSOSystemBehaviorManager() SOSystemBehaviorManager {
	class := getSOSystemBehaviorManagerClass()
	rv := objc.Send[SOSystemBehaviorManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/SpeechObjects/SOSystemBehaviorManager/_numberObjectFromTimer:
func (s SOSystemBehaviorManager) _numberObjectFromTimer(timer objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_numberObjectFromTimer:"), timer)
	return objectivec.Object{ID: rv}
}

// NumberObjectFromTimer is an exported wrapper for the private method _numberObjectFromTimer.
func (s SOSystemBehaviorManager) NumberObjectFromTimer(timer objectivec.IObject) objectivec.IObject {
	return s._numberObjectFromTimer(timer)
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOSystemBehaviorManager/addTimer:
func (s SOSystemBehaviorManager) AddTimer(timer objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("addTimer:"), timer)
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOSystemBehaviorManager/removeTimer:
func (s SOSystemBehaviorManager) RemoveTimer(timer objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("removeTimer:"), timer)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSystemBehaviorManager/sharedSOSystemBehaviorManager
func (_SOSystemBehaviorManagerClass SOSystemBehaviorManagerClass) SharedSOSystemBehaviorManager() SOSystemBehaviorManager {
	rv := objc.Send[objc.ID](objc.ID(_SOSystemBehaviorManagerClass.class), objc.Sel("sharedSOSystemBehaviorManager"))
	return SOSystemBehaviorManagerFromID(rv)
}


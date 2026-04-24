// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXSpecialKeyState] class.
var (
	_CPXSpecialKeyStateClass     CPXSpecialKeyStateClass
	_CPXSpecialKeyStateClassOnce sync.Once
)

func getCPXSpecialKeyStateClass() CPXSpecialKeyStateClass {
	_CPXSpecialKeyStateClassOnce.Do(func() {
		_CPXSpecialKeyStateClass = CPXSpecialKeyStateClass{class: objc.GetClass("CPXSpecialKeyState")}
	})
	return _CPXSpecialKeyStateClass
}

// GetCPXSpecialKeyStateClass returns the class object for CPXSpecialKeyState.
func GetCPXSpecialKeyStateClass() CPXSpecialKeyStateClass {
	return getCPXSpecialKeyStateClass()
}

type CPXSpecialKeyStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXSpecialKeyStateClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXSpecialKeyStateClass) Alloc() CPXSpecialKeyState {
	rv := objc.Send[CPXSpecialKeyState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyState
type CPXSpecialKeyState struct {
	objectivec.Object
}

// CPXSpecialKeyStateFromID constructs a [CPXSpecialKeyState] from an objc.ID.
func CPXSpecialKeyStateFromID(id objc.ID) CPXSpecialKeyState {
	return CPXSpecialKeyState{objectivec.Object{ID: id}}
}

// Ensure CPXSpecialKeyState implements ICPXSpecialKeyState.
var _ ICPXSpecialKeyState = CPXSpecialKeyState{}

// An interface definition for the [CPXSpecialKeyState] class.
//
// See: https://developer.apple.com/documentation/SkyLight/CPXSpecialKeyState
type ICPXSpecialKeyState interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CPXSpecialKeyState) Init() CPXSpecialKeyState {
	rv := objc.Send[CPXSpecialKeyState](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXSpecialKeyState) Autorelease() CPXSpecialKeyState {
	rv := objc.Send[CPXSpecialKeyState](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXSpecialKeyState creates a new CPXSpecialKeyState instance.
func NewCPXSpecialKeyState() CPXSpecialKeyState {
	class := getCPXSpecialKeyStateClass()
	rv := objc.Send[CPXSpecialKeyState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

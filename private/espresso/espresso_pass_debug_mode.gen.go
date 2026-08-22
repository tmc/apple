// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassDebugMode] class.
var (
	_EspressoPassDebugModeClass     EspressoPassDebugModeClass
	_EspressoPassDebugModeClassOnce sync.Once
)

func getEspressoPassDebugModeClass() EspressoPassDebugModeClass {
	_EspressoPassDebugModeClassOnce.Do(func() {
		_EspressoPassDebugModeClass = EspressoPassDebugModeClass{class: objc.GetClass("EspressoPass_debug_mode")}
	})
	return _EspressoPassDebugModeClass
}

// GetEspressoPassDebugModeClass returns the class object for EspressoPass_debug_mode.
func GetEspressoPassDebugModeClass() EspressoPassDebugModeClass {
	return getEspressoPassDebugModeClass()
}

type EspressoPassDebugModeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassDebugModeClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassDebugModeClass) Alloc() EspressoPassDebugMode {
	rv := objc.SendIfResponds[EspressoPassDebugMode](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassDebugMode struct {
	EspressoCustomPass
}

// EspressoPassDebugModeFromID constructs a [EspressoPassDebugMode] from an objc.ID.
func EspressoPassDebugModeFromID(id objc.ID) EspressoPassDebugMode {
	return EspressoPassDebugMode{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_debug_modeFromID is an alias for [EspressoPassDebugModeFromID] for cross-framework compatibility.
func EspressoPass_debug_modeFromID(id objc.ID) EspressoPassDebugMode {
	return EspressoPassDebugModeFromID(id)
}

// Ensure EspressoPassDebugMode implements IEspressoPassDebugMode.
var _ IEspressoPassDebugMode = EspressoPassDebugMode{}

// An interface definition for the [EspressoPassDebugMode] class.
type IEspressoPassDebugMode interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassDebugMode) Init() EspressoPassDebugMode {
	rv := objc.SendIfResponds[EspressoPassDebugMode](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassDebugMode) Autorelease() EspressoPassDebugMode {
	rv := objc.SendIfResponds[EspressoPassDebugMode](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassDebugMode creates a new EspressoPassDebugMode instance.
func NewEspressoPassDebugMode() EspressoPassDebugMode {
	class := getEspressoPassDebugModeClass()
	rv := objc.SendIfResponds[EspressoPassDebugMode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

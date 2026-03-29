// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_debug_mode] class.
var (
	_EspressoPass_debug_modeClass     EspressoPass_debug_modeClass
	_EspressoPass_debug_modeClassOnce sync.Once
)

func getEspressoPass_debug_modeClass() EspressoPass_debug_modeClass {
	_EspressoPass_debug_modeClassOnce.Do(func() {
		_EspressoPass_debug_modeClass = EspressoPass_debug_modeClass{class: objc.GetClass("EspressoPass_debug_mode")}
	})
	return _EspressoPass_debug_modeClass
}

// GetEspressoPass_debug_modeClass returns the class object for EspressoPass_debug_mode.
func GetEspressoPass_debug_modeClass() EspressoPass_debug_modeClass {
	return getEspressoPass_debug_modeClass()
}

type EspressoPass_debug_modeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_debug_modeClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_debug_modeClass) Alloc() EspressoPass_debug_mode {
	rv := objc.Send[EspressoPass_debug_mode](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_debug_mode
type EspressoPass_debug_mode struct {
	EspressoCustomPass
}

// EspressoPass_debug_modeFromID constructs a [EspressoPass_debug_mode] from an objc.ID.
func EspressoPass_debug_modeFromID(id objc.ID) EspressoPass_debug_mode {
	return EspressoPass_debug_mode{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_debug_mode implements IEspressoPass_debug_mode.
var _ IEspressoPass_debug_mode = EspressoPass_debug_mode{}

// An interface definition for the [EspressoPass_debug_mode] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_debug_mode
type IEspressoPass_debug_mode interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_debug_mode) Init() EspressoPass_debug_mode {
	rv := objc.Send[EspressoPass_debug_mode](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_debug_mode) Autorelease() EspressoPass_debug_mode {
	rv := objc.Send[EspressoPass_debug_mode](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_debug_mode creates a new EspressoPass_debug_mode instance.
func NewEspressoPass_debug_mode() EspressoPass_debug_mode {
	class := getEspressoPass_debug_modeClass()
	rv := objc.Send[EspressoPass_debug_mode](objc.ID(class.class), objc.Sel("new"))
	return rv
}


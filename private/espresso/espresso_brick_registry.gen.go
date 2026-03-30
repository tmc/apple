// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoBrickRegistry] class.
var (
	_EspressoBrickRegistryClass     EspressoBrickRegistryClass
	_EspressoBrickRegistryClassOnce sync.Once
)

func getEspressoBrickRegistryClass() EspressoBrickRegistryClass {
	_EspressoBrickRegistryClassOnce.Do(func() {
		_EspressoBrickRegistryClass = EspressoBrickRegistryClass{class: objc.GetClass("EspressoBrickRegistry")}
	})
	return _EspressoBrickRegistryClass
}

// GetEspressoBrickRegistryClass returns the class object for EspressoBrickRegistry.
func GetEspressoBrickRegistryClass() EspressoBrickRegistryClass {
	return getEspressoBrickRegistryClass()
}

type EspressoBrickRegistryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoBrickRegistryClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoBrickRegistryClass) Alloc() EspressoBrickRegistry {
	rv := objc.Send[EspressoBrickRegistry](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoBrickRegistry
type EspressoBrickRegistry struct {
	objectivec.Object
}

// EspressoBrickRegistryFromID constructs a [EspressoBrickRegistry] from an objc.ID.
func EspressoBrickRegistryFromID(id objc.ID) EspressoBrickRegistry {
	return EspressoBrickRegistry{objectivec.Object{ID: id}}
}

// Ensure EspressoBrickRegistry implements IEspressoBrickRegistry.
var _ IEspressoBrickRegistry = EspressoBrickRegistry{}

// An interface definition for the [EspressoBrickRegistry] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoBrickRegistry
type IEspressoBrickRegistry interface {
	objectivec.IObject
}

// Init initializes the instance.
func (e EspressoBrickRegistry) Init() EspressoBrickRegistry {
	rv := objc.Send[EspressoBrickRegistry](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoBrickRegistry) Autorelease() EspressoBrickRegistry {
	rv := objc.Send[EspressoBrickRegistry](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoBrickRegistry creates a new EspressoBrickRegistry instance.
func NewEspressoBrickRegistry() EspressoBrickRegistry {
	class := getEspressoBrickRegistryClass()
	rv := objc.Send[EspressoBrickRegistry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoBrickRegistry/registerBrickClass:
func (_EspressoBrickRegistryClass EspressoBrickRegistryClass) RegisterBrickClass(class objc.Class) {
	objc.Send[objc.ID](objc.ID(_EspressoBrickRegistryClass.class), objc.Sel("registerBrickClass:"), class)
}

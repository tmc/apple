// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoOJBC] class.
var (
	_EspressoOJBCClass     EspressoOJBCClass
	_EspressoOJBCClassOnce sync.Once
)

func getEspressoOJBCClass() EspressoOJBCClass {
	_EspressoOJBCClassOnce.Do(func() {
		_EspressoOJBCClass = EspressoOJBCClass{class: objc.GetClass("EspressoOJBC")}
	})
	return _EspressoOJBCClass
}

// GetEspressoOJBCClass returns the class object for EspressoOJBC.
func GetEspressoOJBCClass() EspressoOJBCClass {
	return getEspressoOJBCClass()
}

type EspressoOJBCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoOJBCClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoOJBCClass) Alloc() EspressoOJBC {
	rv := objc.Send[EspressoOJBC](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoOJBC
type EspressoOJBC struct {
	objectivec.Object
}

// EspressoOJBCFromID constructs a [EspressoOJBC] from an objc.ID.
func EspressoOJBCFromID(id objc.ID) EspressoOJBC {
	return EspressoOJBC{objectivec.Object{ID: id}}
}
// Ensure EspressoOJBC implements IEspressoOJBC.
var _ IEspressoOJBC = EspressoOJBC{}

// An interface definition for the [EspressoOJBC] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoOJBC
type IEspressoOJBC interface {
	objectivec.IObject
}

// Init initializes the instance.
func (e EspressoOJBC) Init() EspressoOJBC {
	rv := objc.Send[EspressoOJBC](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoOJBC) Autorelease() EspressoOJBC {
	rv := objc.Send[EspressoOJBC](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoOJBC creates a new EspressoOJBC instance.
func NewEspressoOJBC() EspressoOJBC {
	class := getEspressoOJBCClass()
	rv := objc.Send[EspressoOJBC](objc.ID(class.class), objc.Sel("new"))
	return rv
}


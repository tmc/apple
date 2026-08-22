// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassFoldConstants] class.
var (
	_EspressoPassFoldConstantsClass     EspressoPassFoldConstantsClass
	_EspressoPassFoldConstantsClassOnce sync.Once
)

func getEspressoPassFoldConstantsClass() EspressoPassFoldConstantsClass {
	_EspressoPassFoldConstantsClassOnce.Do(func() {
		_EspressoPassFoldConstantsClass = EspressoPassFoldConstantsClass{class: objc.GetClass("EspressoPass_fold_constants")}
	})
	return _EspressoPassFoldConstantsClass
}

// GetEspressoPassFoldConstantsClass returns the class object for EspressoPass_fold_constants.
func GetEspressoPassFoldConstantsClass() EspressoPassFoldConstantsClass {
	return getEspressoPassFoldConstantsClass()
}

type EspressoPassFoldConstantsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassFoldConstantsClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassFoldConstantsClass) Alloc() EspressoPassFoldConstants {
	rv := objc.SendIfResponds[EspressoPassFoldConstants](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassFoldConstants struct {
	EspressoCustomPass
}

// EspressoPassFoldConstantsFromID constructs a [EspressoPassFoldConstants] from an objc.ID.
func EspressoPassFoldConstantsFromID(id objc.ID) EspressoPassFoldConstants {
	return EspressoPassFoldConstants{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_fold_constantsFromID is an alias for [EspressoPassFoldConstantsFromID] for cross-framework compatibility.
func EspressoPass_fold_constantsFromID(id objc.ID) EspressoPassFoldConstants {
	return EspressoPassFoldConstantsFromID(id)
}

// Ensure EspressoPassFoldConstants implements IEspressoPassFoldConstants.
var _ IEspressoPassFoldConstants = EspressoPassFoldConstants{}

// An interface definition for the [EspressoPassFoldConstants] class.
type IEspressoPassFoldConstants interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassFoldConstants) Init() EspressoPassFoldConstants {
	rv := objc.SendIfResponds[EspressoPassFoldConstants](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassFoldConstants) Autorelease() EspressoPassFoldConstants {
	rv := objc.SendIfResponds[EspressoPassFoldConstants](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassFoldConstants creates a new EspressoPassFoldConstants instance.
func NewEspressoPassFoldConstants() EspressoPassFoldConstants {
	class := getEspressoPassFoldConstantsClass()
	rv := objc.SendIfResponds[EspressoPassFoldConstants](objc.ID(class.class), objc.Sel("new"))
	return rv
}

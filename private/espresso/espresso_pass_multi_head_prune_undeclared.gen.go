// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_multi_head_prune_undeclared] class.
var (
	_EspressoPass_multi_head_prune_undeclaredClass     EspressoPass_multi_head_prune_undeclaredClass
	_EspressoPass_multi_head_prune_undeclaredClassOnce sync.Once
)

func getEspressoPass_multi_head_prune_undeclaredClass() EspressoPass_multi_head_prune_undeclaredClass {
	_EspressoPass_multi_head_prune_undeclaredClassOnce.Do(func() {
		_EspressoPass_multi_head_prune_undeclaredClass = EspressoPass_multi_head_prune_undeclaredClass{class: objc.GetClass("EspressoPass_multi_head_prune_undeclared")}
	})
	return _EspressoPass_multi_head_prune_undeclaredClass
}

// GetEspressoPass_multi_head_prune_undeclaredClass returns the class object for EspressoPass_multi_head_prune_undeclared.
func GetEspressoPass_multi_head_prune_undeclaredClass() EspressoPass_multi_head_prune_undeclaredClass {
	return getEspressoPass_multi_head_prune_undeclaredClass()
}

type EspressoPass_multi_head_prune_undeclaredClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_multi_head_prune_undeclaredClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_multi_head_prune_undeclaredClass) Alloc() EspressoPass_multi_head_prune_undeclared {
	rv := objc.Send[EspressoPass_multi_head_prune_undeclared](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_multi_head_prune_undeclared
type EspressoPass_multi_head_prune_undeclared struct {
	EspressoCustomPass
}

// EspressoPass_multi_head_prune_undeclaredFromID constructs a [EspressoPass_multi_head_prune_undeclared] from an objc.ID.
func EspressoPass_multi_head_prune_undeclaredFromID(id objc.ID) EspressoPass_multi_head_prune_undeclared {
	return EspressoPass_multi_head_prune_undeclared{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_multi_head_prune_undeclared implements IEspressoPass_multi_head_prune_undeclared.
var _ IEspressoPass_multi_head_prune_undeclared = EspressoPass_multi_head_prune_undeclared{}

// An interface definition for the [EspressoPass_multi_head_prune_undeclared] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_multi_head_prune_undeclared
type IEspressoPass_multi_head_prune_undeclared interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_multi_head_prune_undeclared) Init() EspressoPass_multi_head_prune_undeclared {
	rv := objc.Send[EspressoPass_multi_head_prune_undeclared](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_multi_head_prune_undeclared) Autorelease() EspressoPass_multi_head_prune_undeclared {
	rv := objc.Send[EspressoPass_multi_head_prune_undeclared](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_multi_head_prune_undeclared creates a new EspressoPass_multi_head_prune_undeclared instance.
func NewEspressoPass_multi_head_prune_undeclared() EspressoPass_multi_head_prune_undeclared {
	class := getEspressoPass_multi_head_prune_undeclaredClass()
	rv := objc.Send[EspressoPass_multi_head_prune_undeclared](objc.ID(class.class), objc.Sel("new"))
	return rv
}


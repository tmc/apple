// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassMultiHeadPruneUndeclared] class.
var (
	_EspressoPassMultiHeadPruneUndeclaredClass     EspressoPassMultiHeadPruneUndeclaredClass
	_EspressoPassMultiHeadPruneUndeclaredClassOnce sync.Once
)

func getEspressoPassMultiHeadPruneUndeclaredClass() EspressoPassMultiHeadPruneUndeclaredClass {
	_EspressoPassMultiHeadPruneUndeclaredClassOnce.Do(func() {
		_EspressoPassMultiHeadPruneUndeclaredClass = EspressoPassMultiHeadPruneUndeclaredClass{class: objc.GetClass("EspressoPass_multi_head_prune_undeclared")}
	})
	return _EspressoPassMultiHeadPruneUndeclaredClass
}

// GetEspressoPassMultiHeadPruneUndeclaredClass returns the class object for EspressoPass_multi_head_prune_undeclared.
func GetEspressoPassMultiHeadPruneUndeclaredClass() EspressoPassMultiHeadPruneUndeclaredClass {
	return getEspressoPassMultiHeadPruneUndeclaredClass()
}

type EspressoPassMultiHeadPruneUndeclaredClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassMultiHeadPruneUndeclaredClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassMultiHeadPruneUndeclaredClass) Alloc() EspressoPassMultiHeadPruneUndeclared {
	rv := objc.Send[EspressoPassMultiHeadPruneUndeclared](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_multi_head_prune_undeclared
type EspressoPassMultiHeadPruneUndeclared struct {
	EspressoCustomPass
}

// EspressoPassMultiHeadPruneUndeclaredFromID constructs a [EspressoPassMultiHeadPruneUndeclared] from an objc.ID.
func EspressoPassMultiHeadPruneUndeclaredFromID(id objc.ID) EspressoPassMultiHeadPruneUndeclared {
	return EspressoPassMultiHeadPruneUndeclared{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_multi_head_prune_undeclaredFromID is an alias for [EspressoPassMultiHeadPruneUndeclaredFromID] for cross-framework compatibility.
func EspressoPass_multi_head_prune_undeclaredFromID(id objc.ID) EspressoPassMultiHeadPruneUndeclared {
	return EspressoPassMultiHeadPruneUndeclaredFromID(id)
}

// Ensure EspressoPassMultiHeadPruneUndeclared implements IEspressoPassMultiHeadPruneUndeclared.
var _ IEspressoPassMultiHeadPruneUndeclared = EspressoPassMultiHeadPruneUndeclared{}

// An interface definition for the [EspressoPassMultiHeadPruneUndeclared] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_multi_head_prune_undeclared
type IEspressoPassMultiHeadPruneUndeclared interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassMultiHeadPruneUndeclared) Init() EspressoPassMultiHeadPruneUndeclared {
	rv := objc.Send[EspressoPassMultiHeadPruneUndeclared](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassMultiHeadPruneUndeclared) Autorelease() EspressoPassMultiHeadPruneUndeclared {
	rv := objc.Send[EspressoPassMultiHeadPruneUndeclared](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassMultiHeadPruneUndeclared creates a new EspressoPassMultiHeadPruneUndeclared instance.
func NewEspressoPassMultiHeadPruneUndeclared() EspressoPassMultiHeadPruneUndeclared {
	class := getEspressoPassMultiHeadPruneUndeclaredClass()
	rv := objc.Send[EspressoPassMultiHeadPruneUndeclared](objc.ID(class.class), objc.Sel("new"))
	return rv
}

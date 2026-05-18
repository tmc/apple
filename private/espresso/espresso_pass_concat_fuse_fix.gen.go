// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassConcatFuseFix] class.
var (
	_EspressoPassConcatFuseFixClass     EspressoPassConcatFuseFixClass
	_EspressoPassConcatFuseFixClassOnce sync.Once
)

func getEspressoPassConcatFuseFixClass() EspressoPassConcatFuseFixClass {
	_EspressoPassConcatFuseFixClassOnce.Do(func() {
		_EspressoPassConcatFuseFixClass = EspressoPassConcatFuseFixClass{class: objc.GetClass("EspressoPass_concat_fuse_fix")}
	})
	return _EspressoPassConcatFuseFixClass
}

// GetEspressoPassConcatFuseFixClass returns the class object for EspressoPass_concat_fuse_fix.
func GetEspressoPassConcatFuseFixClass() EspressoPassConcatFuseFixClass {
	return getEspressoPassConcatFuseFixClass()
}

type EspressoPassConcatFuseFixClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassConcatFuseFixClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassConcatFuseFixClass) Alloc() EspressoPassConcatFuseFix {
	rv := objc.Send[EspressoPassConcatFuseFix](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_concat_fuse_fix
type EspressoPassConcatFuseFix struct {
	EspressoCustomPass
}

// EspressoPassConcatFuseFixFromID constructs a [EspressoPassConcatFuseFix] from an objc.ID.
func EspressoPassConcatFuseFixFromID(id objc.ID) EspressoPassConcatFuseFix {
	return EspressoPassConcatFuseFix{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_concat_fuse_fixFromID is an alias for [EspressoPassConcatFuseFixFromID] for cross-framework compatibility.
func EspressoPass_concat_fuse_fixFromID(id objc.ID) EspressoPassConcatFuseFix {
	return EspressoPassConcatFuseFixFromID(id)
}

// Ensure EspressoPassConcatFuseFix implements IEspressoPassConcatFuseFix.
var _ IEspressoPassConcatFuseFix = EspressoPassConcatFuseFix{}

// An interface definition for the [EspressoPassConcatFuseFix] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_concat_fuse_fix
type IEspressoPassConcatFuseFix interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassConcatFuseFix) Init() EspressoPassConcatFuseFix {
	rv := objc.Send[EspressoPassConcatFuseFix](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassConcatFuseFix) Autorelease() EspressoPassConcatFuseFix {
	rv := objc.Send[EspressoPassConcatFuseFix](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassConcatFuseFix creates a new EspressoPassConcatFuseFix instance.
func NewEspressoPassConcatFuseFix() EspressoPassConcatFuseFix {
	class := getEspressoPassConcatFuseFixClass()
	rv := objc.Send[EspressoPassConcatFuseFix](objc.ID(class.class), objc.Sel("new"))
	return rv
}

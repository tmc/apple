// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassFuseGeluWithErf] class.
var (
	_EspressoPassFuseGeluWithErfClass     EspressoPassFuseGeluWithErfClass
	_EspressoPassFuseGeluWithErfClassOnce sync.Once
)

func getEspressoPassFuseGeluWithErfClass() EspressoPassFuseGeluWithErfClass {
	_EspressoPassFuseGeluWithErfClassOnce.Do(func() {
		_EspressoPassFuseGeluWithErfClass = EspressoPassFuseGeluWithErfClass{class: objc.GetClass("EspressoPass_fuse_gelu_with_erf")}
	})
	return _EspressoPassFuseGeluWithErfClass
}

// GetEspressoPassFuseGeluWithErfClass returns the class object for EspressoPass_fuse_gelu_with_erf.
func GetEspressoPassFuseGeluWithErfClass() EspressoPassFuseGeluWithErfClass {
	return getEspressoPassFuseGeluWithErfClass()
}

type EspressoPassFuseGeluWithErfClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassFuseGeluWithErfClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassFuseGeluWithErfClass) Alloc() EspressoPassFuseGeluWithErf {
	rv := objc.Send[EspressoPassFuseGeluWithErf](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassFuseGeluWithErf struct {
	EspressoCustomPass
}

// EspressoPassFuseGeluWithErfFromID constructs a [EspressoPassFuseGeluWithErf] from an objc.ID.
func EspressoPassFuseGeluWithErfFromID(id objc.ID) EspressoPassFuseGeluWithErf {
	return EspressoPassFuseGeluWithErf{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_fuse_gelu_with_erfFromID is an alias for [EspressoPassFuseGeluWithErfFromID] for cross-framework compatibility.
func EspressoPass_fuse_gelu_with_erfFromID(id objc.ID) EspressoPassFuseGeluWithErf {
	return EspressoPassFuseGeluWithErfFromID(id)
}

// Ensure EspressoPassFuseGeluWithErf implements IEspressoPassFuseGeluWithErf.
var _ IEspressoPassFuseGeluWithErf = EspressoPassFuseGeluWithErf{}

// An interface definition for the [EspressoPassFuseGeluWithErf] class.
type IEspressoPassFuseGeluWithErf interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassFuseGeluWithErf) Init() EspressoPassFuseGeluWithErf {
	rv := objc.Send[EspressoPassFuseGeluWithErf](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassFuseGeluWithErf) Autorelease() EspressoPassFuseGeluWithErf {
	rv := objc.Send[EspressoPassFuseGeluWithErf](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassFuseGeluWithErf creates a new EspressoPassFuseGeluWithErf instance.
func NewEspressoPassFuseGeluWithErf() EspressoPassFuseGeluWithErf {
	class := getEspressoPassFuseGeluWithErfClass()
	rv := objc.Send[EspressoPassFuseGeluWithErf](objc.ID(class.class), objc.Sel("new"))
	return rv
}

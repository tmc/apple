// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassFuseGruActivation] class.
var (
	_EspressoPassFuseGruActivationClass     EspressoPassFuseGruActivationClass
	_EspressoPassFuseGruActivationClassOnce sync.Once
)

func getEspressoPassFuseGruActivationClass() EspressoPassFuseGruActivationClass {
	_EspressoPassFuseGruActivationClassOnce.Do(func() {
		_EspressoPassFuseGruActivationClass = EspressoPassFuseGruActivationClass{class: objc.GetClass("EspressoPass_fuse_gru_activation")}
	})
	return _EspressoPassFuseGruActivationClass
}

// GetEspressoPassFuseGruActivationClass returns the class object for EspressoPass_fuse_gru_activation.
func GetEspressoPassFuseGruActivationClass() EspressoPassFuseGruActivationClass {
	return getEspressoPassFuseGruActivationClass()
}

type EspressoPassFuseGruActivationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassFuseGruActivationClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassFuseGruActivationClass) Alloc() EspressoPassFuseGruActivation {
	rv := objc.Send[EspressoPassFuseGruActivation](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassFuseGruActivation struct {
	EspressoCustomPass
}

// EspressoPassFuseGruActivationFromID constructs a [EspressoPassFuseGruActivation] from an objc.ID.
func EspressoPassFuseGruActivationFromID(id objc.ID) EspressoPassFuseGruActivation {
	return EspressoPassFuseGruActivation{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_fuse_gru_activationFromID is an alias for [EspressoPassFuseGruActivationFromID] for cross-framework compatibility.
func EspressoPass_fuse_gru_activationFromID(id objc.ID) EspressoPassFuseGruActivation {
	return EspressoPassFuseGruActivationFromID(id)
}

// Ensure EspressoPassFuseGruActivation implements IEspressoPassFuseGruActivation.
var _ IEspressoPassFuseGruActivation = EspressoPassFuseGruActivation{}

// An interface definition for the [EspressoPassFuseGruActivation] class.
type IEspressoPassFuseGruActivation interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassFuseGruActivation) Init() EspressoPassFuseGruActivation {
	rv := objc.Send[EspressoPassFuseGruActivation](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassFuseGruActivation) Autorelease() EspressoPassFuseGruActivation {
	rv := objc.Send[EspressoPassFuseGruActivation](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassFuseGruActivation creates a new EspressoPassFuseGruActivation instance.
func NewEspressoPassFuseGruActivation() EspressoPassFuseGruActivation {
	class := getEspressoPassFuseGruActivationClass()
	rv := objc.Send[EspressoPassFuseGruActivation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassRemoveSqueezeExpanddims1dOps] class.
var (
	_EspressoPassRemoveSqueezeExpanddims1dOpsClass     EspressoPassRemoveSqueezeExpanddims1dOpsClass
	_EspressoPassRemoveSqueezeExpanddims1dOpsClassOnce sync.Once
)

func getEspressoPassRemoveSqueezeExpanddims1dOpsClass() EspressoPassRemoveSqueezeExpanddims1dOpsClass {
	_EspressoPassRemoveSqueezeExpanddims1dOpsClassOnce.Do(func() {
		_EspressoPassRemoveSqueezeExpanddims1dOpsClass = EspressoPassRemoveSqueezeExpanddims1dOpsClass{class: objc.GetClass("EspressoPass_remove_squeeze_expanddims_1d_ops")}
	})
	return _EspressoPassRemoveSqueezeExpanddims1dOpsClass
}

// GetEspressoPassRemoveSqueezeExpanddims1dOpsClass returns the class object for EspressoPass_remove_squeeze_expanddims_1d_ops.
func GetEspressoPassRemoveSqueezeExpanddims1dOpsClass() EspressoPassRemoveSqueezeExpanddims1dOpsClass {
	return getEspressoPassRemoveSqueezeExpanddims1dOpsClass()
}

type EspressoPassRemoveSqueezeExpanddims1dOpsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassRemoveSqueezeExpanddims1dOpsClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassRemoveSqueezeExpanddims1dOpsClass) Alloc() EspressoPassRemoveSqueezeExpanddims1dOps {
	rv := objc.SendIfResponds[EspressoPassRemoveSqueezeExpanddims1dOps](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassRemoveSqueezeExpanddims1dOps struct {
	EspressoCustomPass
}

// EspressoPassRemoveSqueezeExpanddims1dOpsFromID constructs a [EspressoPassRemoveSqueezeExpanddims1dOps] from an objc.ID.
func EspressoPassRemoveSqueezeExpanddims1dOpsFromID(id objc.ID) EspressoPassRemoveSqueezeExpanddims1dOps {
	return EspressoPassRemoveSqueezeExpanddims1dOps{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_remove_squeeze_expanddims_1d_opsFromID is an alias for [EspressoPassRemoveSqueezeExpanddims1dOpsFromID] for cross-framework compatibility.
func EspressoPass_remove_squeeze_expanddims_1d_opsFromID(id objc.ID) EspressoPassRemoveSqueezeExpanddims1dOps {
	return EspressoPassRemoveSqueezeExpanddims1dOpsFromID(id)
}

// Ensure EspressoPassRemoveSqueezeExpanddims1dOps implements IEspressoPassRemoveSqueezeExpanddims1dOps.
var _ IEspressoPassRemoveSqueezeExpanddims1dOps = EspressoPassRemoveSqueezeExpanddims1dOps{}

// An interface definition for the [EspressoPassRemoveSqueezeExpanddims1dOps] class.
type IEspressoPassRemoveSqueezeExpanddims1dOps interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassRemoveSqueezeExpanddims1dOps) Init() EspressoPassRemoveSqueezeExpanddims1dOps {
	rv := objc.SendIfResponds[EspressoPassRemoveSqueezeExpanddims1dOps](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassRemoveSqueezeExpanddims1dOps) Autorelease() EspressoPassRemoveSqueezeExpanddims1dOps {
	rv := objc.SendIfResponds[EspressoPassRemoveSqueezeExpanddims1dOps](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassRemoveSqueezeExpanddims1dOps creates a new EspressoPassRemoveSqueezeExpanddims1dOps instance.
func NewEspressoPassRemoveSqueezeExpanddims1dOps() EspressoPassRemoveSqueezeExpanddims1dOps {
	class := getEspressoPassRemoveSqueezeExpanddims1dOpsClass()
	rv := objc.SendIfResponds[EspressoPassRemoveSqueezeExpanddims1dOps](objc.ID(class.class), objc.Sel("new"))
	return rv
}

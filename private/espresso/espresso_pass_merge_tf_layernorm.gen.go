// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassMergeTfLayernorm] class.
var (
	_EspressoPassMergeTfLayernormClass     EspressoPassMergeTfLayernormClass
	_EspressoPassMergeTfLayernormClassOnce sync.Once
)

func getEspressoPassMergeTfLayernormClass() EspressoPassMergeTfLayernormClass {
	_EspressoPassMergeTfLayernormClassOnce.Do(func() {
		_EspressoPassMergeTfLayernormClass = EspressoPassMergeTfLayernormClass{class: objc.GetClass("EspressoPass_merge_tf_layernorm")}
	})
	return _EspressoPassMergeTfLayernormClass
}

// GetEspressoPassMergeTfLayernormClass returns the class object for EspressoPass_merge_tf_layernorm.
func GetEspressoPassMergeTfLayernormClass() EspressoPassMergeTfLayernormClass {
	return getEspressoPassMergeTfLayernormClass()
}

type EspressoPassMergeTfLayernormClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassMergeTfLayernormClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassMergeTfLayernormClass) Alloc() EspressoPassMergeTfLayernorm {
	rv := objc.SendIfResponds[EspressoPassMergeTfLayernorm](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassMergeTfLayernorm struct {
	EspressoCustomPass
}

// EspressoPassMergeTfLayernormFromID constructs a [EspressoPassMergeTfLayernorm] from an objc.ID.
func EspressoPassMergeTfLayernormFromID(id objc.ID) EspressoPassMergeTfLayernorm {
	return EspressoPassMergeTfLayernorm{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_merge_tf_layernormFromID is an alias for [EspressoPassMergeTfLayernormFromID] for cross-framework compatibility.
func EspressoPass_merge_tf_layernormFromID(id objc.ID) EspressoPassMergeTfLayernorm {
	return EspressoPassMergeTfLayernormFromID(id)
}

// Ensure EspressoPassMergeTfLayernorm implements IEspressoPassMergeTfLayernorm.
var _ IEspressoPassMergeTfLayernorm = EspressoPassMergeTfLayernorm{}

// An interface definition for the [EspressoPassMergeTfLayernorm] class.
type IEspressoPassMergeTfLayernorm interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassMergeTfLayernorm) Init() EspressoPassMergeTfLayernorm {
	rv := objc.SendIfResponds[EspressoPassMergeTfLayernorm](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassMergeTfLayernorm) Autorelease() EspressoPassMergeTfLayernorm {
	rv := objc.SendIfResponds[EspressoPassMergeTfLayernorm](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassMergeTfLayernorm creates a new EspressoPassMergeTfLayernorm instance.
func NewEspressoPassMergeTfLayernorm() EspressoPassMergeTfLayernorm {
	class := getEspressoPassMergeTfLayernormClass()
	rv := objc.SendIfResponds[EspressoPassMergeTfLayernorm](objc.ID(class.class), objc.Sel("new"))
	return rv
}

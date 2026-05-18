// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassFuseAddAndRelu] class.
var (
	_EspressoPassFuseAddAndReluClass     EspressoPassFuseAddAndReluClass
	_EspressoPassFuseAddAndReluClassOnce sync.Once
)

func getEspressoPassFuseAddAndReluClass() EspressoPassFuseAddAndReluClass {
	_EspressoPassFuseAddAndReluClassOnce.Do(func() {
		_EspressoPassFuseAddAndReluClass = EspressoPassFuseAddAndReluClass{class: objc.GetClass("EspressoPass_fuse_add_and_relu")}
	})
	return _EspressoPassFuseAddAndReluClass
}

// GetEspressoPassFuseAddAndReluClass returns the class object for EspressoPass_fuse_add_and_relu.
func GetEspressoPassFuseAddAndReluClass() EspressoPassFuseAddAndReluClass {
	return getEspressoPassFuseAddAndReluClass()
}

type EspressoPassFuseAddAndReluClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassFuseAddAndReluClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassFuseAddAndReluClass) Alloc() EspressoPassFuseAddAndRelu {
	rv := objc.Send[EspressoPassFuseAddAndRelu](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_add_and_relu
type EspressoPassFuseAddAndRelu struct {
	EspressoCustomPass
}

// EspressoPassFuseAddAndReluFromID constructs a [EspressoPassFuseAddAndRelu] from an objc.ID.
func EspressoPassFuseAddAndReluFromID(id objc.ID) EspressoPassFuseAddAndRelu {
	return EspressoPassFuseAddAndRelu{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_fuse_add_and_reluFromID is an alias for [EspressoPassFuseAddAndReluFromID] for cross-framework compatibility.
func EspressoPass_fuse_add_and_reluFromID(id objc.ID) EspressoPassFuseAddAndRelu {
	return EspressoPassFuseAddAndReluFromID(id)
}

// Ensure EspressoPassFuseAddAndRelu implements IEspressoPassFuseAddAndRelu.
var _ IEspressoPassFuseAddAndRelu = EspressoPassFuseAddAndRelu{}

// An interface definition for the [EspressoPassFuseAddAndRelu] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_add_and_relu
type IEspressoPassFuseAddAndRelu interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassFuseAddAndRelu) Init() EspressoPassFuseAddAndRelu {
	rv := objc.Send[EspressoPassFuseAddAndRelu](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassFuseAddAndRelu) Autorelease() EspressoPassFuseAddAndRelu {
	rv := objc.Send[EspressoPassFuseAddAndRelu](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassFuseAddAndRelu creates a new EspressoPassFuseAddAndRelu instance.
func NewEspressoPassFuseAddAndRelu() EspressoPassFuseAddAndRelu {
	class := getEspressoPassFuseAddAndReluClass()
	rv := objc.Send[EspressoPassFuseAddAndRelu](objc.ID(class.class), objc.Sel("new"))
	return rv
}

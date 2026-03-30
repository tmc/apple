// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_fuse_add_and_relu] class.
var (
	_EspressoPass_fuse_add_and_reluClass     EspressoPass_fuse_add_and_reluClass
	_EspressoPass_fuse_add_and_reluClassOnce sync.Once
)

func getEspressoPass_fuse_add_and_reluClass() EspressoPass_fuse_add_and_reluClass {
	_EspressoPass_fuse_add_and_reluClassOnce.Do(func() {
		_EspressoPass_fuse_add_and_reluClass = EspressoPass_fuse_add_and_reluClass{class: objc.GetClass("EspressoPass_fuse_add_and_relu")}
	})
	return _EspressoPass_fuse_add_and_reluClass
}

// GetEspressoPass_fuse_add_and_reluClass returns the class object for EspressoPass_fuse_add_and_relu.
func GetEspressoPass_fuse_add_and_reluClass() EspressoPass_fuse_add_and_reluClass {
	return getEspressoPass_fuse_add_and_reluClass()
}

type EspressoPass_fuse_add_and_reluClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_fuse_add_and_reluClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_fuse_add_and_reluClass) Alloc() EspressoPass_fuse_add_and_relu {
	rv := objc.Send[EspressoPass_fuse_add_and_relu](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_add_and_relu
type EspressoPass_fuse_add_and_relu struct {
	EspressoCustomPass
}

// EspressoPass_fuse_add_and_reluFromID constructs a [EspressoPass_fuse_add_and_relu] from an objc.ID.
func EspressoPass_fuse_add_and_reluFromID(id objc.ID) EspressoPass_fuse_add_and_relu {
	return EspressoPass_fuse_add_and_relu{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// Ensure EspressoPass_fuse_add_and_relu implements IEspressoPass_fuse_add_and_relu.
var _ IEspressoPass_fuse_add_and_relu = EspressoPass_fuse_add_and_relu{}

// An interface definition for the [EspressoPass_fuse_add_and_relu] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_add_and_relu
type IEspressoPass_fuse_add_and_relu interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_fuse_add_and_relu) Init() EspressoPass_fuse_add_and_relu {
	rv := objc.Send[EspressoPass_fuse_add_and_relu](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_fuse_add_and_relu) Autorelease() EspressoPass_fuse_add_and_relu {
	rv := objc.Send[EspressoPass_fuse_add_and_relu](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_fuse_add_and_relu creates a new EspressoPass_fuse_add_and_relu instance.
func NewEspressoPass_fuse_add_and_relu() EspressoPass_fuse_add_and_relu {
	class := getEspressoPass_fuse_add_and_reluClass()
	rv := objc.Send[EspressoPass_fuse_add_and_relu](objc.ID(class.class), objc.Sel("new"))
	return rv
}

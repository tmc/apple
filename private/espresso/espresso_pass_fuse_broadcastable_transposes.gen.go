// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_fuse_broadcastable_transposes] class.
var (
	_EspressoPass_fuse_broadcastable_transposesClass     EspressoPass_fuse_broadcastable_transposesClass
	_EspressoPass_fuse_broadcastable_transposesClassOnce sync.Once
)

func getEspressoPass_fuse_broadcastable_transposesClass() EspressoPass_fuse_broadcastable_transposesClass {
	_EspressoPass_fuse_broadcastable_transposesClassOnce.Do(func() {
		_EspressoPass_fuse_broadcastable_transposesClass = EspressoPass_fuse_broadcastable_transposesClass{class: objc.GetClass("EspressoPass_fuse_broadcastable_transposes")}
	})
	return _EspressoPass_fuse_broadcastable_transposesClass
}

// GetEspressoPass_fuse_broadcastable_transposesClass returns the class object for EspressoPass_fuse_broadcastable_transposes.
func GetEspressoPass_fuse_broadcastable_transposesClass() EspressoPass_fuse_broadcastable_transposesClass {
	return getEspressoPass_fuse_broadcastable_transposesClass()
}

type EspressoPass_fuse_broadcastable_transposesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_fuse_broadcastable_transposesClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_fuse_broadcastable_transposesClass) Alloc() EspressoPass_fuse_broadcastable_transposes {
	rv := objc.Send[EspressoPass_fuse_broadcastable_transposes](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_broadcastable_transposes
type EspressoPass_fuse_broadcastable_transposes struct {
	EspressoCustomPass
}

// EspressoPass_fuse_broadcastable_transposesFromID constructs a [EspressoPass_fuse_broadcastable_transposes] from an objc.ID.
func EspressoPass_fuse_broadcastable_transposesFromID(id objc.ID) EspressoPass_fuse_broadcastable_transposes {
	return EspressoPass_fuse_broadcastable_transposes{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_fuse_broadcastable_transposes implements IEspressoPass_fuse_broadcastable_transposes.
var _ IEspressoPass_fuse_broadcastable_transposes = EspressoPass_fuse_broadcastable_transposes{}

// An interface definition for the [EspressoPass_fuse_broadcastable_transposes] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_broadcastable_transposes
type IEspressoPass_fuse_broadcastable_transposes interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_fuse_broadcastable_transposes) Init() EspressoPass_fuse_broadcastable_transposes {
	rv := objc.Send[EspressoPass_fuse_broadcastable_transposes](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_fuse_broadcastable_transposes) Autorelease() EspressoPass_fuse_broadcastable_transposes {
	rv := objc.Send[EspressoPass_fuse_broadcastable_transposes](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_fuse_broadcastable_transposes creates a new EspressoPass_fuse_broadcastable_transposes instance.
func NewEspressoPass_fuse_broadcastable_transposes() EspressoPass_fuse_broadcastable_transposes {
	class := getEspressoPass_fuse_broadcastable_transposesClass()
	rv := objc.Send[EspressoPass_fuse_broadcastable_transposes](objc.ID(class.class), objc.Sel("new"))
	return rv
}


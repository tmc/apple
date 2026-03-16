// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_fuse_pad] class.
var (
	_EspressoPass_fuse_padClass     EspressoPass_fuse_padClass
	_EspressoPass_fuse_padClassOnce sync.Once
)

func getEspressoPass_fuse_padClass() EspressoPass_fuse_padClass {
	_EspressoPass_fuse_padClassOnce.Do(func() {
		_EspressoPass_fuse_padClass = EspressoPass_fuse_padClass{class: objc.GetClass("EspressoPass_fuse_pad")}
	})
	return _EspressoPass_fuse_padClass
}

// GetEspressoPass_fuse_padClass returns the class object for EspressoPass_fuse_pad.
func GetEspressoPass_fuse_padClass() EspressoPass_fuse_padClass {
	return getEspressoPass_fuse_padClass()
}

type EspressoPass_fuse_padClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_fuse_padClass) Alloc() EspressoPass_fuse_pad {
	rv := objc.Send[EspressoPass_fuse_pad](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_pad
type EspressoPass_fuse_pad struct {
	EspressoCustomPass
}

// EspressoPass_fuse_padFromID constructs a [EspressoPass_fuse_pad] from an objc.ID.
func EspressoPass_fuse_padFromID(id objc.ID) EspressoPass_fuse_pad {
	return EspressoPass_fuse_pad{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_fuse_pad implements IEspressoPass_fuse_pad.
var _ IEspressoPass_fuse_pad = EspressoPass_fuse_pad{}

// An interface definition for the [EspressoPass_fuse_pad] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_pad
type IEspressoPass_fuse_pad interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_fuse_pad) Init() EspressoPass_fuse_pad {
	rv := objc.Send[EspressoPass_fuse_pad](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_fuse_pad) Autorelease() EspressoPass_fuse_pad {
	rv := objc.Send[EspressoPass_fuse_pad](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_fuse_pad creates a new EspressoPass_fuse_pad instance.
func NewEspressoPass_fuse_pad() EspressoPass_fuse_pad {
	class := getEspressoPass_fuse_padClass()
	rv := objc.Send[EspressoPass_fuse_pad](objc.ID(class.class), objc.Sel("new"))
	return rv
}


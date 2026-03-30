// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_merge_channel_norm] class.
var (
	_EspressoPass_merge_channel_normClass     EspressoPass_merge_channel_normClass
	_EspressoPass_merge_channel_normClassOnce sync.Once
)

func getEspressoPass_merge_channel_normClass() EspressoPass_merge_channel_normClass {
	_EspressoPass_merge_channel_normClassOnce.Do(func() {
		_EspressoPass_merge_channel_normClass = EspressoPass_merge_channel_normClass{class: objc.GetClass("EspressoPass_merge_channel_norm")}
	})
	return _EspressoPass_merge_channel_normClass
}

// GetEspressoPass_merge_channel_normClass returns the class object for EspressoPass_merge_channel_norm.
func GetEspressoPass_merge_channel_normClass() EspressoPass_merge_channel_normClass {
	return getEspressoPass_merge_channel_normClass()
}

type EspressoPass_merge_channel_normClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_merge_channel_normClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_merge_channel_normClass) Alloc() EspressoPass_merge_channel_norm {
	rv := objc.Send[EspressoPass_merge_channel_norm](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_merge_channel_norm
type EspressoPass_merge_channel_norm struct {
	EspressoCustomPass
}

// EspressoPass_merge_channel_normFromID constructs a [EspressoPass_merge_channel_norm] from an objc.ID.
func EspressoPass_merge_channel_normFromID(id objc.ID) EspressoPass_merge_channel_norm {
	return EspressoPass_merge_channel_norm{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// Ensure EspressoPass_merge_channel_norm implements IEspressoPass_merge_channel_norm.
var _ IEspressoPass_merge_channel_norm = EspressoPass_merge_channel_norm{}

// An interface definition for the [EspressoPass_merge_channel_norm] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_merge_channel_norm
type IEspressoPass_merge_channel_norm interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_merge_channel_norm) Init() EspressoPass_merge_channel_norm {
	rv := objc.Send[EspressoPass_merge_channel_norm](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_merge_channel_norm) Autorelease() EspressoPass_merge_channel_norm {
	rv := objc.Send[EspressoPass_merge_channel_norm](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_merge_channel_norm creates a new EspressoPass_merge_channel_norm instance.
func NewEspressoPass_merge_channel_norm() EspressoPass_merge_channel_norm {
	class := getEspressoPass_merge_channel_normClass()
	rv := objc.Send[EspressoPass_merge_channel_norm](objc.ID(class.class), objc.Sel("new"))
	return rv
}

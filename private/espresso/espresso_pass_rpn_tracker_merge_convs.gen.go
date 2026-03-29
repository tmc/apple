// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_rpn_tracker_merge_convs] class.
var (
	_EspressoPass_rpn_tracker_merge_convsClass     EspressoPass_rpn_tracker_merge_convsClass
	_EspressoPass_rpn_tracker_merge_convsClassOnce sync.Once
)

func getEspressoPass_rpn_tracker_merge_convsClass() EspressoPass_rpn_tracker_merge_convsClass {
	_EspressoPass_rpn_tracker_merge_convsClassOnce.Do(func() {
		_EspressoPass_rpn_tracker_merge_convsClass = EspressoPass_rpn_tracker_merge_convsClass{class: objc.GetClass("EspressoPass_rpn_tracker_merge_convs")}
	})
	return _EspressoPass_rpn_tracker_merge_convsClass
}

// GetEspressoPass_rpn_tracker_merge_convsClass returns the class object for EspressoPass_rpn_tracker_merge_convs.
func GetEspressoPass_rpn_tracker_merge_convsClass() EspressoPass_rpn_tracker_merge_convsClass {
	return getEspressoPass_rpn_tracker_merge_convsClass()
}

type EspressoPass_rpn_tracker_merge_convsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_rpn_tracker_merge_convsClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_rpn_tracker_merge_convsClass) Alloc() EspressoPass_rpn_tracker_merge_convs {
	rv := objc.Send[EspressoPass_rpn_tracker_merge_convs](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_rpn_tracker_merge_convs
type EspressoPass_rpn_tracker_merge_convs struct {
	EspressoCustomPass
}

// EspressoPass_rpn_tracker_merge_convsFromID constructs a [EspressoPass_rpn_tracker_merge_convs] from an objc.ID.
func EspressoPass_rpn_tracker_merge_convsFromID(id objc.ID) EspressoPass_rpn_tracker_merge_convs {
	return EspressoPass_rpn_tracker_merge_convs{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_rpn_tracker_merge_convs implements IEspressoPass_rpn_tracker_merge_convs.
var _ IEspressoPass_rpn_tracker_merge_convs = EspressoPass_rpn_tracker_merge_convs{}

// An interface definition for the [EspressoPass_rpn_tracker_merge_convs] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_rpn_tracker_merge_convs
type IEspressoPass_rpn_tracker_merge_convs interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_rpn_tracker_merge_convs) Init() EspressoPass_rpn_tracker_merge_convs {
	rv := objc.Send[EspressoPass_rpn_tracker_merge_convs](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_rpn_tracker_merge_convs) Autorelease() EspressoPass_rpn_tracker_merge_convs {
	rv := objc.Send[EspressoPass_rpn_tracker_merge_convs](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_rpn_tracker_merge_convs creates a new EspressoPass_rpn_tracker_merge_convs instance.
func NewEspressoPass_rpn_tracker_merge_convs() EspressoPass_rpn_tracker_merge_convs {
	class := getEspressoPass_rpn_tracker_merge_convsClass()
	rv := objc.Send[EspressoPass_rpn_tracker_merge_convs](objc.ID(class.class), objc.Sel("new"))
	return rv
}


// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassRpnTrackerMergeConvs] class.
var (
	_EspressoPassRpnTrackerMergeConvsClass     EspressoPassRpnTrackerMergeConvsClass
	_EspressoPassRpnTrackerMergeConvsClassOnce sync.Once
)

func getEspressoPassRpnTrackerMergeConvsClass() EspressoPassRpnTrackerMergeConvsClass {
	_EspressoPassRpnTrackerMergeConvsClassOnce.Do(func() {
		_EspressoPassRpnTrackerMergeConvsClass = EspressoPassRpnTrackerMergeConvsClass{class: objc.GetClass("EspressoPass_rpn_tracker_merge_convs")}
	})
	return _EspressoPassRpnTrackerMergeConvsClass
}

// GetEspressoPassRpnTrackerMergeConvsClass returns the class object for EspressoPass_rpn_tracker_merge_convs.
func GetEspressoPassRpnTrackerMergeConvsClass() EspressoPassRpnTrackerMergeConvsClass {
	return getEspressoPassRpnTrackerMergeConvsClass()
}

type EspressoPassRpnTrackerMergeConvsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassRpnTrackerMergeConvsClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassRpnTrackerMergeConvsClass) Alloc() EspressoPassRpnTrackerMergeConvs {
	rv := objc.Send[EspressoPassRpnTrackerMergeConvs](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_rpn_tracker_merge_convs
type EspressoPassRpnTrackerMergeConvs struct {
	EspressoCustomPass
}

// EspressoPassRpnTrackerMergeConvsFromID constructs a [EspressoPassRpnTrackerMergeConvs] from an objc.ID.
func EspressoPassRpnTrackerMergeConvsFromID(id objc.ID) EspressoPassRpnTrackerMergeConvs {
	return EspressoPassRpnTrackerMergeConvs{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_rpn_tracker_merge_convsFromID is an alias for [EspressoPassRpnTrackerMergeConvsFromID] for cross-framework compatibility.
func EspressoPass_rpn_tracker_merge_convsFromID(id objc.ID) EspressoPassRpnTrackerMergeConvs {
	return EspressoPassRpnTrackerMergeConvsFromID(id)
}

// Ensure EspressoPassRpnTrackerMergeConvs implements IEspressoPassRpnTrackerMergeConvs.
var _ IEspressoPassRpnTrackerMergeConvs = EspressoPassRpnTrackerMergeConvs{}

// An interface definition for the [EspressoPassRpnTrackerMergeConvs] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_rpn_tracker_merge_convs
type IEspressoPassRpnTrackerMergeConvs interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassRpnTrackerMergeConvs) Init() EspressoPassRpnTrackerMergeConvs {
	rv := objc.Send[EspressoPassRpnTrackerMergeConvs](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassRpnTrackerMergeConvs) Autorelease() EspressoPassRpnTrackerMergeConvs {
	rv := objc.Send[EspressoPassRpnTrackerMergeConvs](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassRpnTrackerMergeConvs creates a new EspressoPassRpnTrackerMergeConvs instance.
func NewEspressoPassRpnTrackerMergeConvs() EspressoPassRpnTrackerMergeConvs {
	class := getEspressoPassRpnTrackerMergeConvsClass()
	rv := objc.Send[EspressoPassRpnTrackerMergeConvs](objc.ID(class.class), objc.Sel("new"))
	return rv
}

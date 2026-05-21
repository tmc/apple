// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassFusePad] class.
var (
	_EspressoPassFusePadClass     EspressoPassFusePadClass
	_EspressoPassFusePadClassOnce sync.Once
)

func getEspressoPassFusePadClass() EspressoPassFusePadClass {
	_EspressoPassFusePadClassOnce.Do(func() {
		_EspressoPassFusePadClass = EspressoPassFusePadClass{class: objc.GetClass("EspressoPass_fuse_pad")}
	})
	return _EspressoPassFusePadClass
}

// GetEspressoPassFusePadClass returns the class object for EspressoPass_fuse_pad.
func GetEspressoPassFusePadClass() EspressoPassFusePadClass {
	return getEspressoPassFusePadClass()
}

type EspressoPassFusePadClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassFusePadClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassFusePadClass) Alloc() EspressoPassFusePad {
	rv := objc.Send[EspressoPassFusePad](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassFusePad struct {
	EspressoCustomPass
}

// EspressoPassFusePadFromID constructs a [EspressoPassFusePad] from an objc.ID.
func EspressoPassFusePadFromID(id objc.ID) EspressoPassFusePad {
	return EspressoPassFusePad{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_fuse_padFromID is an alias for [EspressoPassFusePadFromID] for cross-framework compatibility.
func EspressoPass_fuse_padFromID(id objc.ID) EspressoPassFusePad {
	return EspressoPassFusePadFromID(id)
}

// Ensure EspressoPassFusePad implements IEspressoPassFusePad.
var _ IEspressoPassFusePad = EspressoPassFusePad{}

// An interface definition for the [EspressoPassFusePad] class.
type IEspressoPassFusePad interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassFusePad) Init() EspressoPassFusePad {
	rv := objc.Send[EspressoPassFusePad](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassFusePad) Autorelease() EspressoPassFusePad {
	rv := objc.Send[EspressoPassFusePad](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassFusePad creates a new EspressoPassFusePad instance.
func NewEspressoPassFusePad() EspressoPassFusePad {
	class := getEspressoPassFusePadClass()
	rv := objc.Send[EspressoPassFusePad](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassQuantizeGathers] class.
var (
	_EspressoPassQuantizeGathersClass     EspressoPassQuantizeGathersClass
	_EspressoPassQuantizeGathersClassOnce sync.Once
)

func getEspressoPassQuantizeGathersClass() EspressoPassQuantizeGathersClass {
	_EspressoPassQuantizeGathersClassOnce.Do(func() {
		_EspressoPassQuantizeGathersClass = EspressoPassQuantizeGathersClass{class: objc.GetClass("EspressoPass_quantize_gathers")}
	})
	return _EspressoPassQuantizeGathersClass
}

// GetEspressoPassQuantizeGathersClass returns the class object for EspressoPass_quantize_gathers.
func GetEspressoPassQuantizeGathersClass() EspressoPassQuantizeGathersClass {
	return getEspressoPassQuantizeGathersClass()
}

type EspressoPassQuantizeGathersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassQuantizeGathersClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassQuantizeGathersClass) Alloc() EspressoPassQuantizeGathers {
	rv := objc.Send[EspressoPassQuantizeGathers](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassQuantizeGathers struct {
	EspressoCustomPass
}

// EspressoPassQuantizeGathersFromID constructs a [EspressoPassQuantizeGathers] from an objc.ID.
func EspressoPassQuantizeGathersFromID(id objc.ID) EspressoPassQuantizeGathers {
	return EspressoPassQuantizeGathers{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_quantize_gathersFromID is an alias for [EspressoPassQuantizeGathersFromID] for cross-framework compatibility.
func EspressoPass_quantize_gathersFromID(id objc.ID) EspressoPassQuantizeGathers {
	return EspressoPassQuantizeGathersFromID(id)
}

// Ensure EspressoPassQuantizeGathers implements IEspressoPassQuantizeGathers.
var _ IEspressoPassQuantizeGathers = EspressoPassQuantizeGathers{}

// An interface definition for the [EspressoPassQuantizeGathers] class.
type IEspressoPassQuantizeGathers interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassQuantizeGathers) Init() EspressoPassQuantizeGathers {
	rv := objc.Send[EspressoPassQuantizeGathers](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassQuantizeGathers) Autorelease() EspressoPassQuantizeGathers {
	rv := objc.Send[EspressoPassQuantizeGathers](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassQuantizeGathers creates a new EspressoPassQuantizeGathers instance.
func NewEspressoPassQuantizeGathers() EspressoPassQuantizeGathers {
	class := getEspressoPassQuantizeGathersClass()
	rv := objc.Send[EspressoPassQuantizeGathers](objc.ID(class.class), objc.Sel("new"))
	return rv
}

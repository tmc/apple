// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoCustomPass] class.
var (
	_EspressoCustomPassClass     EspressoCustomPassClass
	_EspressoCustomPassClassOnce sync.Once
)

func getEspressoCustomPassClass() EspressoCustomPassClass {
	_EspressoCustomPassClassOnce.Do(func() {
		_EspressoCustomPassClass = EspressoCustomPassClass{class: objc.GetClass("EspressoCustomPass")}
	})
	return _EspressoCustomPassClass
}

// GetEspressoCustomPassClass returns the class object for EspressoCustomPass.
func GetEspressoCustomPassClass() EspressoCustomPassClass {
	return getEspressoCustomPassClass()
}

type EspressoCustomPassClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoCustomPassClass) Alloc() EspressoCustomPass {
	rv := objc.Send[EspressoCustomPass](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [EspressoCustomPass.RunOnNetwork]
// See: https://developer.apple.com/documentation/Espresso/EspressoCustomPass
type EspressoCustomPass struct {
	objectivec.Object
}

// EspressoCustomPassFromID constructs a [EspressoCustomPass] from an objc.ID.
func EspressoCustomPassFromID(id objc.ID) EspressoCustomPass {
	return EspressoCustomPass{objectivec.Object{id}}
}
// Ensure EspressoCustomPass implements IEspressoCustomPass.
var _ IEspressoCustomPass = EspressoCustomPass{}





// An interface definition for the [EspressoCustomPass] class.
//
// # Methods
//
//   - [IEspressoCustomPass.RunOnNetwork]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoCustomPass
type IEspressoCustomPass interface {
	objectivec.IObject

	// Topic: Methods

	RunOnNetwork(network unsafe.Pointer) bool
}





// Init initializes the instance.
func (e EspressoCustomPass) Init() EspressoCustomPass {
	rv := objc.Send[EspressoCustomPass](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoCustomPass) Autorelease() EspressoCustomPass {
	rv := objc.Send[EspressoCustomPass](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoCustomPass creates a new EspressoCustomPass instance.
func NewEspressoCustomPass() EspressoCustomPass {
	class := getEspressoCustomPassClass()
	rv := objc.Send[EspressoCustomPass](objc.ID(class.class), objc.Sel("new"))
	return rv
}










//
// See: https://developer.apple.com/documentation/Espresso/EspressoCustomPass/runOnNetwork:
func (e EspressoCustomPass) RunOnNetwork(network unsafe.Pointer) bool {
	rv := objc.Send[bool](e.ID, objc.Sel("runOnNetwork:"), network)
	return rv
}



























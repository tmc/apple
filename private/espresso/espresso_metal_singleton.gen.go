// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoMetalSingleton] class.
var (
	_EspressoMetalSingletonClass     EspressoMetalSingletonClass
	_EspressoMetalSingletonClassOnce sync.Once
)

func getEspressoMetalSingletonClass() EspressoMetalSingletonClass {
	_EspressoMetalSingletonClassOnce.Do(func() {
		_EspressoMetalSingletonClass = EspressoMetalSingletonClass{class: objc.GetClass("EspressoMetalSingleton")}
	})
	return _EspressoMetalSingletonClass
}

// GetEspressoMetalSingletonClass returns the class object for EspressoMetalSingleton.
func GetEspressoMetalSingletonClass() EspressoMetalSingletonClass {
	return getEspressoMetalSingletonClass()
}

type EspressoMetalSingletonClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoMetalSingletonClass) Alloc() EspressoMetalSingleton {
	rv := objc.Send[EspressoMetalSingleton](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoMetalSingleton.Is_memory_tight]
//   - [EspressoMetalSingleton.SetIs_memory_tight]
// See: https://developer.apple.com/documentation/Espresso/EspressoMetalSingleton
type EspressoMetalSingleton struct {
	objectivec.Object
}

// EspressoMetalSingletonFromID constructs a [EspressoMetalSingleton] from an objc.ID.
func EspressoMetalSingletonFromID(id objc.ID) EspressoMetalSingleton {
	return EspressoMetalSingleton{objectivec.Object{ID: id}}
}
// Ensure EspressoMetalSingleton implements IEspressoMetalSingleton.
var _ IEspressoMetalSingleton = EspressoMetalSingleton{}

// An interface definition for the [EspressoMetalSingleton] class.
//
// # Methods
//
//   - [IEspressoMetalSingleton.Is_memory_tight]
//   - [IEspressoMetalSingleton.SetIs_memory_tight]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoMetalSingleton
type IEspressoMetalSingleton interface {
	objectivec.IObject

	// Topic: Methods

	Is_memory_tight() int
	SetIs_memory_tight(value int)
}

// Init initializes the instance.
func (e EspressoMetalSingleton) Init() EspressoMetalSingleton {
	rv := objc.Send[EspressoMetalSingleton](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoMetalSingleton) Autorelease() EspressoMetalSingleton {
	rv := objc.Send[EspressoMetalSingleton](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoMetalSingleton creates a new EspressoMetalSingleton instance.
func NewEspressoMetalSingleton() EspressoMetalSingleton {
	class := getEspressoMetalSingletonClass()
	rv := objc.Send[EspressoMetalSingleton](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoMetalSingleton/shared
func (_EspressoMetalSingletonClass EspressoMetalSingletonClass) Shared() EspressoMetalSingleton {
	rv := objc.Send[objc.ID](objc.ID(_EspressoMetalSingletonClass.class), objc.Sel("shared"))
	return EspressoMetalSingletonFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoMetalSingleton/is_memory_tight
func (e EspressoMetalSingleton) Is_memory_tight() int {
	rv := objc.Send[int](e.ID, objc.Sel("is_memory_tight"))
	return rv
}
func (e EspressoMetalSingleton) SetIs_memory_tight(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setIs_memory_tight:"), value)
}


// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_fuse_fast_gelu_1] class.
var (
	_EspressoPass_fuse_fast_gelu_1Class     EspressoPass_fuse_fast_gelu_1Class
	_EspressoPass_fuse_fast_gelu_1ClassOnce sync.Once
)

func getEspressoPass_fuse_fast_gelu_1Class() EspressoPass_fuse_fast_gelu_1Class {
	_EspressoPass_fuse_fast_gelu_1ClassOnce.Do(func() {
		_EspressoPass_fuse_fast_gelu_1Class = EspressoPass_fuse_fast_gelu_1Class{class: objc.GetClass("EspressoPass_fuse_fast_gelu_1")}
	})
	return _EspressoPass_fuse_fast_gelu_1Class
}

// GetEspressoPass_fuse_fast_gelu_1Class returns the class object for EspressoPass_fuse_fast_gelu_1.
func GetEspressoPass_fuse_fast_gelu_1Class() EspressoPass_fuse_fast_gelu_1Class {
	return getEspressoPass_fuse_fast_gelu_1Class()
}

type EspressoPass_fuse_fast_gelu_1Class struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_fuse_fast_gelu_1Class) Alloc() EspressoPass_fuse_fast_gelu_1 {
	rv := objc.Send[EspressoPass_fuse_fast_gelu_1](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_fast_gelu_1
type EspressoPass_fuse_fast_gelu_1 struct {
	EspressoCustomPass
}

// EspressoPass_fuse_fast_gelu_1FromID constructs a [EspressoPass_fuse_fast_gelu_1] from an objc.ID.
func EspressoPass_fuse_fast_gelu_1FromID(id objc.ID) EspressoPass_fuse_fast_gelu_1 {
	return EspressoPass_fuse_fast_gelu_1{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_fuse_fast_gelu_1 implements IEspressoPass_fuse_fast_gelu_1.
var _ IEspressoPass_fuse_fast_gelu_1 = EspressoPass_fuse_fast_gelu_1{}





// An interface definition for the [EspressoPass_fuse_fast_gelu_1] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_fast_gelu_1
type IEspressoPass_fuse_fast_gelu_1 interface {
	IEspressoCustomPass
}





// Init initializes the instance.
func (e EspressoPass_fuse_fast_gelu_1) Init() EspressoPass_fuse_fast_gelu_1 {
	rv := objc.Send[EspressoPass_fuse_fast_gelu_1](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_fuse_fast_gelu_1) Autorelease() EspressoPass_fuse_fast_gelu_1 {
	rv := objc.Send[EspressoPass_fuse_fast_gelu_1](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_fuse_fast_gelu_1 creates a new EspressoPass_fuse_fast_gelu_1 instance.
func NewEspressoPass_fuse_fast_gelu_1() EspressoPass_fuse_fast_gelu_1 {
	class := getEspressoPass_fuse_fast_gelu_1Class()
	rv := objc.Send[EspressoPass_fuse_fast_gelu_1](objc.ID(class.class), objc.Sel("new"))
	return rv
}




































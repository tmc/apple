// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_merge_pytorch_layernorm] class.
var (
	_EspressoPass_merge_pytorch_layernormClass     EspressoPass_merge_pytorch_layernormClass
	_EspressoPass_merge_pytorch_layernormClassOnce sync.Once
)

func getEspressoPass_merge_pytorch_layernormClass() EspressoPass_merge_pytorch_layernormClass {
	_EspressoPass_merge_pytorch_layernormClassOnce.Do(func() {
		_EspressoPass_merge_pytorch_layernormClass = EspressoPass_merge_pytorch_layernormClass{class: objc.GetClass("EspressoPass_merge_pytorch_layernorm")}
	})
	return _EspressoPass_merge_pytorch_layernormClass
}

// GetEspressoPass_merge_pytorch_layernormClass returns the class object for EspressoPass_merge_pytorch_layernorm.
func GetEspressoPass_merge_pytorch_layernormClass() EspressoPass_merge_pytorch_layernormClass {
	return getEspressoPass_merge_pytorch_layernormClass()
}

type EspressoPass_merge_pytorch_layernormClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_merge_pytorch_layernormClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_merge_pytorch_layernormClass) Alloc() EspressoPass_merge_pytorch_layernorm {
	rv := objc.Send[EspressoPass_merge_pytorch_layernorm](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_merge_pytorch_layernorm
type EspressoPass_merge_pytorch_layernorm struct {
	EspressoCustomPass
}

// EspressoPass_merge_pytorch_layernormFromID constructs a [EspressoPass_merge_pytorch_layernorm] from an objc.ID.
func EspressoPass_merge_pytorch_layernormFromID(id objc.ID) EspressoPass_merge_pytorch_layernorm {
	return EspressoPass_merge_pytorch_layernorm{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_merge_pytorch_layernorm implements IEspressoPass_merge_pytorch_layernorm.
var _ IEspressoPass_merge_pytorch_layernorm = EspressoPass_merge_pytorch_layernorm{}

// An interface definition for the [EspressoPass_merge_pytorch_layernorm] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_merge_pytorch_layernorm
type IEspressoPass_merge_pytorch_layernorm interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_merge_pytorch_layernorm) Init() EspressoPass_merge_pytorch_layernorm {
	rv := objc.Send[EspressoPass_merge_pytorch_layernorm](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_merge_pytorch_layernorm) Autorelease() EspressoPass_merge_pytorch_layernorm {
	rv := objc.Send[EspressoPass_merge_pytorch_layernorm](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_merge_pytorch_layernorm creates a new EspressoPass_merge_pytorch_layernorm instance.
func NewEspressoPass_merge_pytorch_layernorm() EspressoPass_merge_pytorch_layernorm {
	class := getEspressoPass_merge_pytorch_layernormClass()
	rv := objc.Send[EspressoPass_merge_pytorch_layernorm](objc.ID(class.class), objc.Sel("new"))
	return rv
}


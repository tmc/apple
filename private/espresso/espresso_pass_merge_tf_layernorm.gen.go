// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_merge_tf_layernorm] class.
var (
	_EspressoPass_merge_tf_layernormClass     EspressoPass_merge_tf_layernormClass
	_EspressoPass_merge_tf_layernormClassOnce sync.Once
)

func getEspressoPass_merge_tf_layernormClass() EspressoPass_merge_tf_layernormClass {
	_EspressoPass_merge_tf_layernormClassOnce.Do(func() {
		_EspressoPass_merge_tf_layernormClass = EspressoPass_merge_tf_layernormClass{class: objc.GetClass("EspressoPass_merge_tf_layernorm")}
	})
	return _EspressoPass_merge_tf_layernormClass
}

// GetEspressoPass_merge_tf_layernormClass returns the class object for EspressoPass_merge_tf_layernorm.
func GetEspressoPass_merge_tf_layernormClass() EspressoPass_merge_tf_layernormClass {
	return getEspressoPass_merge_tf_layernormClass()
}

type EspressoPass_merge_tf_layernormClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_merge_tf_layernormClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_merge_tf_layernormClass) Alloc() EspressoPass_merge_tf_layernorm {
	rv := objc.Send[EspressoPass_merge_tf_layernorm](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_merge_tf_layernorm
type EspressoPass_merge_tf_layernorm struct {
	EspressoCustomPass
}

// EspressoPass_merge_tf_layernormFromID constructs a [EspressoPass_merge_tf_layernorm] from an objc.ID.
func EspressoPass_merge_tf_layernormFromID(id objc.ID) EspressoPass_merge_tf_layernorm {
	return EspressoPass_merge_tf_layernorm{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// Ensure EspressoPass_merge_tf_layernorm implements IEspressoPass_merge_tf_layernorm.
var _ IEspressoPass_merge_tf_layernorm = EspressoPass_merge_tf_layernorm{}

// An interface definition for the [EspressoPass_merge_tf_layernorm] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_merge_tf_layernorm
type IEspressoPass_merge_tf_layernorm interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_merge_tf_layernorm) Init() EspressoPass_merge_tf_layernorm {
	rv := objc.Send[EspressoPass_merge_tf_layernorm](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_merge_tf_layernorm) Autorelease() EspressoPass_merge_tf_layernorm {
	rv := objc.Send[EspressoPass_merge_tf_layernorm](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_merge_tf_layernorm creates a new EspressoPass_merge_tf_layernorm instance.
func NewEspressoPass_merge_tf_layernorm() EspressoPass_merge_tf_layernorm {
	class := getEspressoPass_merge_tf_layernormClass()
	rv := objc.Send[EspressoPass_merge_tf_layernorm](objc.ID(class.class), objc.Sel("new"))
	return rv
}

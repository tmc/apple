// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_fuse_conv_batchnorm] class.
var (
	_EspressoPass_fuse_conv_batchnormClass     EspressoPass_fuse_conv_batchnormClass
	_EspressoPass_fuse_conv_batchnormClassOnce sync.Once
)

func getEspressoPass_fuse_conv_batchnormClass() EspressoPass_fuse_conv_batchnormClass {
	_EspressoPass_fuse_conv_batchnormClassOnce.Do(func() {
		_EspressoPass_fuse_conv_batchnormClass = EspressoPass_fuse_conv_batchnormClass{class: objc.GetClass("EspressoPass_fuse_conv_batchnorm")}
	})
	return _EspressoPass_fuse_conv_batchnormClass
}

// GetEspressoPass_fuse_conv_batchnormClass returns the class object for EspressoPass_fuse_conv_batchnorm.
func GetEspressoPass_fuse_conv_batchnormClass() EspressoPass_fuse_conv_batchnormClass {
	return getEspressoPass_fuse_conv_batchnormClass()
}

type EspressoPass_fuse_conv_batchnormClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_fuse_conv_batchnormClass) Alloc() EspressoPass_fuse_conv_batchnorm {
	rv := objc.Send[EspressoPass_fuse_conv_batchnorm](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_conv_batchnorm
type EspressoPass_fuse_conv_batchnorm struct {
	EspressoCustomPass
}

// EspressoPass_fuse_conv_batchnormFromID constructs a [EspressoPass_fuse_conv_batchnorm] from an objc.ID.
func EspressoPass_fuse_conv_batchnormFromID(id objc.ID) EspressoPass_fuse_conv_batchnorm {
	return EspressoPass_fuse_conv_batchnorm{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_fuse_conv_batchnorm implements IEspressoPass_fuse_conv_batchnorm.
var _ IEspressoPass_fuse_conv_batchnorm = EspressoPass_fuse_conv_batchnorm{}

// An interface definition for the [EspressoPass_fuse_conv_batchnorm] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_conv_batchnorm
type IEspressoPass_fuse_conv_batchnorm interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_fuse_conv_batchnorm) Init() EspressoPass_fuse_conv_batchnorm {
	rv := objc.Send[EspressoPass_fuse_conv_batchnorm](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_fuse_conv_batchnorm) Autorelease() EspressoPass_fuse_conv_batchnorm {
	rv := objc.Send[EspressoPass_fuse_conv_batchnorm](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_fuse_conv_batchnorm creates a new EspressoPass_fuse_conv_batchnorm instance.
func NewEspressoPass_fuse_conv_batchnorm() EspressoPass_fuse_conv_batchnorm {
	class := getEspressoPass_fuse_conv_batchnormClass()
	rv := objc.Send[EspressoPass_fuse_conv_batchnorm](objc.ID(class.class), objc.Sel("new"))
	return rv
}


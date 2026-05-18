// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassFuseConvBatchnorm] class.
var (
	_EspressoPassFuseConvBatchnormClass     EspressoPassFuseConvBatchnormClass
	_EspressoPassFuseConvBatchnormClassOnce sync.Once
)

func getEspressoPassFuseConvBatchnormClass() EspressoPassFuseConvBatchnormClass {
	_EspressoPassFuseConvBatchnormClassOnce.Do(func() {
		_EspressoPassFuseConvBatchnormClass = EspressoPassFuseConvBatchnormClass{class: objc.GetClass("EspressoPass_fuse_conv_batchnorm")}
	})
	return _EspressoPassFuseConvBatchnormClass
}

// GetEspressoPassFuseConvBatchnormClass returns the class object for EspressoPass_fuse_conv_batchnorm.
func GetEspressoPassFuseConvBatchnormClass() EspressoPassFuseConvBatchnormClass {
	return getEspressoPassFuseConvBatchnormClass()
}

type EspressoPassFuseConvBatchnormClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassFuseConvBatchnormClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassFuseConvBatchnormClass) Alloc() EspressoPassFuseConvBatchnorm {
	rv := objc.Send[EspressoPassFuseConvBatchnorm](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_conv_batchnorm
type EspressoPassFuseConvBatchnorm struct {
	EspressoCustomPass
}

// EspressoPassFuseConvBatchnormFromID constructs a [EspressoPassFuseConvBatchnorm] from an objc.ID.
func EspressoPassFuseConvBatchnormFromID(id objc.ID) EspressoPassFuseConvBatchnorm {
	return EspressoPassFuseConvBatchnorm{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_fuse_conv_batchnormFromID is an alias for [EspressoPassFuseConvBatchnormFromID] for cross-framework compatibility.
func EspressoPass_fuse_conv_batchnormFromID(id objc.ID) EspressoPassFuseConvBatchnorm {
	return EspressoPassFuseConvBatchnormFromID(id)
}

// Ensure EspressoPassFuseConvBatchnorm implements IEspressoPassFuseConvBatchnorm.
var _ IEspressoPassFuseConvBatchnorm = EspressoPassFuseConvBatchnorm{}

// An interface definition for the [EspressoPassFuseConvBatchnorm] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_conv_batchnorm
type IEspressoPassFuseConvBatchnorm interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassFuseConvBatchnorm) Init() EspressoPassFuseConvBatchnorm {
	rv := objc.Send[EspressoPassFuseConvBatchnorm](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassFuseConvBatchnorm) Autorelease() EspressoPassFuseConvBatchnorm {
	rv := objc.Send[EspressoPassFuseConvBatchnorm](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassFuseConvBatchnorm creates a new EspressoPassFuseConvBatchnorm instance.
func NewEspressoPassFuseConvBatchnorm() EspressoPassFuseConvBatchnorm {
	class := getEspressoPassFuseConvBatchnormClass()
	rv := objc.Send[EspressoPassFuseConvBatchnorm](objc.ID(class.class), objc.Sel("new"))
	return rv
}

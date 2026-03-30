// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_remove_squeeze_expanddims_1d_ops] class.
var (
	_EspressoPass_remove_squeeze_expanddims_1d_opsClass     EspressoPass_remove_squeeze_expanddims_1d_opsClass
	_EspressoPass_remove_squeeze_expanddims_1d_opsClassOnce sync.Once
)

func getEspressoPass_remove_squeeze_expanddims_1d_opsClass() EspressoPass_remove_squeeze_expanddims_1d_opsClass {
	_EspressoPass_remove_squeeze_expanddims_1d_opsClassOnce.Do(func() {
		_EspressoPass_remove_squeeze_expanddims_1d_opsClass = EspressoPass_remove_squeeze_expanddims_1d_opsClass{class: objc.GetClass("EspressoPass_remove_squeeze_expanddims_1d_ops")}
	})
	return _EspressoPass_remove_squeeze_expanddims_1d_opsClass
}

// GetEspressoPass_remove_squeeze_expanddims_1d_opsClass returns the class object for EspressoPass_remove_squeeze_expanddims_1d_ops.
func GetEspressoPass_remove_squeeze_expanddims_1d_opsClass() EspressoPass_remove_squeeze_expanddims_1d_opsClass {
	return getEspressoPass_remove_squeeze_expanddims_1d_opsClass()
}

type EspressoPass_remove_squeeze_expanddims_1d_opsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_remove_squeeze_expanddims_1d_opsClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_remove_squeeze_expanddims_1d_opsClass) Alloc() EspressoPass_remove_squeeze_expanddims_1d_ops {
	rv := objc.Send[EspressoPass_remove_squeeze_expanddims_1d_ops](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_remove_squeeze_expanddims_1d_ops
type EspressoPass_remove_squeeze_expanddims_1d_ops struct {
	EspressoCustomPass
}

// EspressoPass_remove_squeeze_expanddims_1d_opsFromID constructs a [EspressoPass_remove_squeeze_expanddims_1d_ops] from an objc.ID.
func EspressoPass_remove_squeeze_expanddims_1d_opsFromID(id objc.ID) EspressoPass_remove_squeeze_expanddims_1d_ops {
	return EspressoPass_remove_squeeze_expanddims_1d_ops{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// Ensure EspressoPass_remove_squeeze_expanddims_1d_ops implements IEspressoPass_remove_squeeze_expanddims_1d_ops.
var _ IEspressoPass_remove_squeeze_expanddims_1d_ops = EspressoPass_remove_squeeze_expanddims_1d_ops{}

// An interface definition for the [EspressoPass_remove_squeeze_expanddims_1d_ops] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_remove_squeeze_expanddims_1d_ops
type IEspressoPass_remove_squeeze_expanddims_1d_ops interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_remove_squeeze_expanddims_1d_ops) Init() EspressoPass_remove_squeeze_expanddims_1d_ops {
	rv := objc.Send[EspressoPass_remove_squeeze_expanddims_1d_ops](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_remove_squeeze_expanddims_1d_ops) Autorelease() EspressoPass_remove_squeeze_expanddims_1d_ops {
	rv := objc.Send[EspressoPass_remove_squeeze_expanddims_1d_ops](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_remove_squeeze_expanddims_1d_ops creates a new EspressoPass_remove_squeeze_expanddims_1d_ops instance.
func NewEspressoPass_remove_squeeze_expanddims_1d_ops() EspressoPass_remove_squeeze_expanddims_1d_ops {
	class := getEspressoPass_remove_squeeze_expanddims_1d_opsClass()
	rv := objc.Send[EspressoPass_remove_squeeze_expanddims_1d_ops](objc.ID(class.class), objc.Sel("new"))
	return rv
}

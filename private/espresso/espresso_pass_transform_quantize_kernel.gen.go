// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_transform_quantize_kernel] class.
var (
	_EspressoPass_transform_quantize_kernelClass     EspressoPass_transform_quantize_kernelClass
	_EspressoPass_transform_quantize_kernelClassOnce sync.Once
)

func getEspressoPass_transform_quantize_kernelClass() EspressoPass_transform_quantize_kernelClass {
	_EspressoPass_transform_quantize_kernelClassOnce.Do(func() {
		_EspressoPass_transform_quantize_kernelClass = EspressoPass_transform_quantize_kernelClass{class: objc.GetClass("EspressoPass_transform_quantize_kernel")}
	})
	return _EspressoPass_transform_quantize_kernelClass
}

// GetEspressoPass_transform_quantize_kernelClass returns the class object for EspressoPass_transform_quantize_kernel.
func GetEspressoPass_transform_quantize_kernelClass() EspressoPass_transform_quantize_kernelClass {
	return getEspressoPass_transform_quantize_kernelClass()
}

type EspressoPass_transform_quantize_kernelClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_transform_quantize_kernelClass) Alloc() EspressoPass_transform_quantize_kernel {
	rv := objc.Send[EspressoPass_transform_quantize_kernel](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_transform_quantize_kernel
type EspressoPass_transform_quantize_kernel struct {
	EspressoCustomPass
}

// EspressoPass_transform_quantize_kernelFromID constructs a [EspressoPass_transform_quantize_kernel] from an objc.ID.
func EspressoPass_transform_quantize_kernelFromID(id objc.ID) EspressoPass_transform_quantize_kernel {
	return EspressoPass_transform_quantize_kernel{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_transform_quantize_kernel implements IEspressoPass_transform_quantize_kernel.
var _ IEspressoPass_transform_quantize_kernel = EspressoPass_transform_quantize_kernel{}

// An interface definition for the [EspressoPass_transform_quantize_kernel] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_transform_quantize_kernel
type IEspressoPass_transform_quantize_kernel interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_transform_quantize_kernel) Init() EspressoPass_transform_quantize_kernel {
	rv := objc.Send[EspressoPass_transform_quantize_kernel](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_transform_quantize_kernel) Autorelease() EspressoPass_transform_quantize_kernel {
	rv := objc.Send[EspressoPass_transform_quantize_kernel](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_transform_quantize_kernel creates a new EspressoPass_transform_quantize_kernel instance.
func NewEspressoPass_transform_quantize_kernel() EspressoPass_transform_quantize_kernel {
	class := getEspressoPass_transform_quantize_kernelClass()
	rv := objc.Send[EspressoPass_transform_quantize_kernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}


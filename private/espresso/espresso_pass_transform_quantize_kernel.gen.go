// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassTransformQuantizeKernel] class.
var (
	_EspressoPassTransformQuantizeKernelClass     EspressoPassTransformQuantizeKernelClass
	_EspressoPassTransformQuantizeKernelClassOnce sync.Once
)

func getEspressoPassTransformQuantizeKernelClass() EspressoPassTransformQuantizeKernelClass {
	_EspressoPassTransformQuantizeKernelClassOnce.Do(func() {
		_EspressoPassTransformQuantizeKernelClass = EspressoPassTransformQuantizeKernelClass{class: objc.GetClass("EspressoPass_transform_quantize_kernel")}
	})
	return _EspressoPassTransformQuantizeKernelClass
}

// GetEspressoPassTransformQuantizeKernelClass returns the class object for EspressoPass_transform_quantize_kernel.
func GetEspressoPassTransformQuantizeKernelClass() EspressoPassTransformQuantizeKernelClass {
	return getEspressoPassTransformQuantizeKernelClass()
}

type EspressoPassTransformQuantizeKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassTransformQuantizeKernelClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassTransformQuantizeKernelClass) Alloc() EspressoPassTransformQuantizeKernel {
	rv := objc.Send[EspressoPassTransformQuantizeKernel](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_transform_quantize_kernel
type EspressoPassTransformQuantizeKernel struct {
	EspressoCustomPass
}

// EspressoPassTransformQuantizeKernelFromID constructs a [EspressoPassTransformQuantizeKernel] from an objc.ID.
func EspressoPassTransformQuantizeKernelFromID(id objc.ID) EspressoPassTransformQuantizeKernel {
	return EspressoPassTransformQuantizeKernel{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_transform_quantize_kernelFromID is an alias for [EspressoPassTransformQuantizeKernelFromID] for cross-framework compatibility.
func EspressoPass_transform_quantize_kernelFromID(id objc.ID) EspressoPassTransformQuantizeKernel {
	return EspressoPassTransformQuantizeKernelFromID(id)
}

// Ensure EspressoPassTransformQuantizeKernel implements IEspressoPassTransformQuantizeKernel.
var _ IEspressoPassTransformQuantizeKernel = EspressoPassTransformQuantizeKernel{}

// An interface definition for the [EspressoPassTransformQuantizeKernel] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_transform_quantize_kernel
type IEspressoPassTransformQuantizeKernel interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassTransformQuantizeKernel) Init() EspressoPassTransformQuantizeKernel {
	rv := objc.Send[EspressoPassTransformQuantizeKernel](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassTransformQuantizeKernel) Autorelease() EspressoPassTransformQuantizeKernel {
	rv := objc.Send[EspressoPassTransformQuantizeKernel](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassTransformQuantizeKernel creates a new EspressoPassTransformQuantizeKernel instance.
func NewEspressoPassTransformQuantizeKernel() EspressoPassTransformQuantizeKernel {
	class := getEspressoPassTransformQuantizeKernelClass()
	rv := objc.Send[EspressoPassTransformQuantizeKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassTransformStaticQuantizeKernel] class.
var (
	_EspressoPassTransformStaticQuantizeKernelClass     EspressoPassTransformStaticQuantizeKernelClass
	_EspressoPassTransformStaticQuantizeKernelClassOnce sync.Once
)

func getEspressoPassTransformStaticQuantizeKernelClass() EspressoPassTransformStaticQuantizeKernelClass {
	_EspressoPassTransformStaticQuantizeKernelClassOnce.Do(func() {
		_EspressoPassTransformStaticQuantizeKernelClass = EspressoPassTransformStaticQuantizeKernelClass{class: objc.GetClass("EspressoPass_transform_static_quantize_kernel")}
	})
	return _EspressoPassTransformStaticQuantizeKernelClass
}

// GetEspressoPassTransformStaticQuantizeKernelClass returns the class object for EspressoPass_transform_static_quantize_kernel.
func GetEspressoPassTransformStaticQuantizeKernelClass() EspressoPassTransformStaticQuantizeKernelClass {
	return getEspressoPassTransformStaticQuantizeKernelClass()
}

type EspressoPassTransformStaticQuantizeKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassTransformStaticQuantizeKernelClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassTransformStaticQuantizeKernelClass) Alloc() EspressoPassTransformStaticQuantizeKernel {
	rv := objc.Send[EspressoPassTransformStaticQuantizeKernel](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassTransformStaticQuantizeKernel struct {
	EspressoCustomPass
}

// EspressoPassTransformStaticQuantizeKernelFromID constructs a [EspressoPassTransformStaticQuantizeKernel] from an objc.ID.
func EspressoPassTransformStaticQuantizeKernelFromID(id objc.ID) EspressoPassTransformStaticQuantizeKernel {
	return EspressoPassTransformStaticQuantizeKernel{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_transform_static_quantize_kernelFromID is an alias for [EspressoPassTransformStaticQuantizeKernelFromID] for cross-framework compatibility.
func EspressoPass_transform_static_quantize_kernelFromID(id objc.ID) EspressoPassTransformStaticQuantizeKernel {
	return EspressoPassTransformStaticQuantizeKernelFromID(id)
}

// Ensure EspressoPassTransformStaticQuantizeKernel implements IEspressoPassTransformStaticQuantizeKernel.
var _ IEspressoPassTransformStaticQuantizeKernel = EspressoPassTransformStaticQuantizeKernel{}

// An interface definition for the [EspressoPassTransformStaticQuantizeKernel] class.
type IEspressoPassTransformStaticQuantizeKernel interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassTransformStaticQuantizeKernel) Init() EspressoPassTransformStaticQuantizeKernel {
	rv := objc.Send[EspressoPassTransformStaticQuantizeKernel](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassTransformStaticQuantizeKernel) Autorelease() EspressoPassTransformStaticQuantizeKernel {
	rv := objc.Send[EspressoPassTransformStaticQuantizeKernel](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassTransformStaticQuantizeKernel creates a new EspressoPassTransformStaticQuantizeKernel instance.
func NewEspressoPassTransformStaticQuantizeKernel() EspressoPassTransformStaticQuantizeKernel {
	class := getEspressoPassTransformStaticQuantizeKernelClass()
	rv := objc.Send[EspressoPassTransformStaticQuantizeKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSKernel] class.
var (
	_MPSKernelClass     MPSKernelClass
	_MPSKernelClassOnce sync.Once
)

func getMPSKernelClass() MPSKernelClass {
	_MPSKernelClassOnce.Do(func() {
		_MPSKernelClass = MPSKernelClass{class: objc.GetClass("MPSKernel")}
	})
	return _MPSKernelClass
}

// GetMPSKernelClass returns the class object for MPSKernel.
func GetMPSKernelClass() MPSKernelClass {
	return getMPSKernelClass()
}

type MPSKernelClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSKernelClass) Alloc() MPSKernel {
	rv := objc.Send[MPSKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A parent class referenced by other neuralnetworks classes. [Full Topic]
type MPSKernel struct {
	objectivec.Object
}

// MPSKernelFromID constructs a [MPSKernel] from an objc.ID.
//
// A parent class referenced by other neuralnetworks classes.
func MPSKernelFromID(id objc.ID) MPSKernel {
	return MPSKernel{objectivec.Object{id}}
}
// Ensure MPSKernel implements IMPSKernel.
var _ IMPSKernel = MPSKernel{}





// An interface definition for the [MPSKernel] class.
type IMPSKernel interface {
	objectivec.IObject
}





// Init initializes the instance.
func (k MPSKernel) Init() MPSKernel {
	rv := objc.Send[MPSKernel](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k MPSKernel) Autorelease() MPSKernel {
	rv := objc.Send[MPSKernel](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSKernel creates a new MPSKernel instance.
func NewMPSKernel() MPSKernel {
	class := getMPSKernelClass()
	rv := objc.Send[MPSKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}




































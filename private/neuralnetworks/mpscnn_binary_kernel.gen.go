// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNBinaryKernel] class.
var (
	_MPSCNNBinaryKernelClass     MPSCNNBinaryKernelClass
	_MPSCNNBinaryKernelClassOnce sync.Once
)

func getMPSCNNBinaryKernelClass() MPSCNNBinaryKernelClass {
	_MPSCNNBinaryKernelClassOnce.Do(func() {
		_MPSCNNBinaryKernelClass = MPSCNNBinaryKernelClass{class: objc.GetClass("MPSCNNBinaryKernel")}
	})
	return _MPSCNNBinaryKernelClass
}

// GetMPSCNNBinaryKernelClass returns the class object for MPSCNNBinaryKernel.
func GetMPSCNNBinaryKernelClass() MPSCNNBinaryKernelClass {
	return getMPSCNNBinaryKernelClass()
}

type MPSCNNBinaryKernelClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNBinaryKernelClass) Alloc() MPSCNNBinaryKernel {
	rv := objc.Send[MPSCNNBinaryKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A parent class referenced by other neuralnetworks classes. [Full Topic]
type MPSCNNBinaryKernel struct {
	objectivec.Object
}

// MPSCNNBinaryKernelFromID constructs a [MPSCNNBinaryKernel] from an objc.ID.
//
// A parent class referenced by other neuralnetworks classes.
func MPSCNNBinaryKernelFromID(id objc.ID) MPSCNNBinaryKernel {
	return MPSCNNBinaryKernel{objectivec.Object{id}}
}
// Ensure MPSCNNBinaryKernel implements IMPSCNNBinaryKernel.
var _ IMPSCNNBinaryKernel = MPSCNNBinaryKernel{}





// An interface definition for the [MPSCNNBinaryKernel] class.
type IMPSCNNBinaryKernel interface {
	objectivec.IObject
}





// Init initializes the instance.
func (c MPSCNNBinaryKernel) Init() MPSCNNBinaryKernel {
	rv := objc.Send[MPSCNNBinaryKernel](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNBinaryKernel) Autorelease() MPSCNNBinaryKernel {
	rv := objc.Send[MPSCNNBinaryKernel](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNBinaryKernel creates a new MPSCNNBinaryKernel instance.
func NewMPSCNNBinaryKernel() MPSCNNBinaryKernel {
	class := getMPSCNNBinaryKernelClass()
	rv := objc.Send[MPSCNNBinaryKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}




































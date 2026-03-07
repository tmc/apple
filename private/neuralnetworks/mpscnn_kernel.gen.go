// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNKernel] class.
var (
	_MPSCNNKernelClass     MPSCNNKernelClass
	_MPSCNNKernelClassOnce sync.Once
)

func getMPSCNNKernelClass() MPSCNNKernelClass {
	_MPSCNNKernelClassOnce.Do(func() {
		_MPSCNNKernelClass = MPSCNNKernelClass{class: objc.GetClass("MPSCNNKernel")}
	})
	return _MPSCNNKernelClass
}

// GetMPSCNNKernelClass returns the class object for MPSCNNKernel.
func GetMPSCNNKernelClass() MPSCNNKernelClass {
	return getMPSCNNKernelClass()
}

type MPSCNNKernelClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNKernelClass) Alloc() MPSCNNKernel {
	rv := objc.Send[MPSCNNKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A parent class referenced by other neuralnetworks classes. [Full Topic]
type MPSCNNKernel struct {
	objectivec.Object
}

// MPSCNNKernelFromID constructs a [MPSCNNKernel] from an objc.ID.
//
// A parent class referenced by other neuralnetworks classes.
func MPSCNNKernelFromID(id objc.ID) MPSCNNKernel {
	return MPSCNNKernel{objectivec.Object{id}}
}
// Ensure MPSCNNKernel implements IMPSCNNKernel.
var _ IMPSCNNKernel = MPSCNNKernel{}





// An interface definition for the [MPSCNNKernel] class.
type IMPSCNNKernel interface {
	objectivec.IObject
}





// Init initializes the instance.
func (c MPSCNNKernel) Init() MPSCNNKernel {
	rv := objc.Send[MPSCNNKernel](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNKernel) Autorelease() MPSCNNKernel {
	rv := objc.Send[MPSCNNKernel](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNKernel creates a new MPSCNNKernel instance.
func NewMPSCNNKernel() MPSCNNKernel {
	class := getMPSCNNKernelClass()
	rv := objc.Send[MPSCNNKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}




































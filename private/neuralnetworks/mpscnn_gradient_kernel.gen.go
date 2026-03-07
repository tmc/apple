// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNGradientKernel] class.
var (
	_MPSCNNGradientKernelClass     MPSCNNGradientKernelClass
	_MPSCNNGradientKernelClassOnce sync.Once
)

func getMPSCNNGradientKernelClass() MPSCNNGradientKernelClass {
	_MPSCNNGradientKernelClassOnce.Do(func() {
		_MPSCNNGradientKernelClass = MPSCNNGradientKernelClass{class: objc.GetClass("MPSCNNGradientKernel")}
	})
	return _MPSCNNGradientKernelClass
}

// GetMPSCNNGradientKernelClass returns the class object for MPSCNNGradientKernel.
func GetMPSCNNGradientKernelClass() MPSCNNGradientKernelClass {
	return getMPSCNNGradientKernelClass()
}

type MPSCNNGradientKernelClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNGradientKernelClass) Alloc() MPSCNNGradientKernel {
	rv := objc.Send[MPSCNNGradientKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A parent class referenced by other neuralnetworks classes. [Full Topic]
type MPSCNNGradientKernel struct {
	objectivec.Object
}

// MPSCNNGradientKernelFromID constructs a [MPSCNNGradientKernel] from an objc.ID.
//
// A parent class referenced by other neuralnetworks classes.
func MPSCNNGradientKernelFromID(id objc.ID) MPSCNNGradientKernel {
	return MPSCNNGradientKernel{objectivec.Object{id}}
}
// Ensure MPSCNNGradientKernel implements IMPSCNNGradientKernel.
var _ IMPSCNNGradientKernel = MPSCNNGradientKernel{}





// An interface definition for the [MPSCNNGradientKernel] class.
type IMPSCNNGradientKernel interface {
	objectivec.IObject
}





// Init initializes the instance.
func (c MPSCNNGradientKernel) Init() MPSCNNGradientKernel {
	rv := objc.Send[MPSCNNGradientKernel](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNGradientKernel) Autorelease() MPSCNNGradientKernel {
	rv := objc.Send[MPSCNNGradientKernel](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNGradientKernel creates a new MPSCNNGradientKernel instance.
func NewMPSCNNGradientKernel() MPSCNNGradientKernel {
	class := getMPSCNNGradientKernelClass()
	rv := objc.Send[MPSCNNGradientKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}




































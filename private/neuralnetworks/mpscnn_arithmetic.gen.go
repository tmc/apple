// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNArithmetic] class.
var (
	_MPSCNNArithmeticClass     MPSCNNArithmeticClass
	_MPSCNNArithmeticClassOnce sync.Once
)

func getMPSCNNArithmeticClass() MPSCNNArithmeticClass {
	_MPSCNNArithmeticClassOnce.Do(func() {
		_MPSCNNArithmeticClass = MPSCNNArithmeticClass{class: objc.GetClass("MPSCNNArithmetic")}
	})
	return _MPSCNNArithmeticClass
}

// GetMPSCNNArithmeticClass returns the class object for MPSCNNArithmetic.
func GetMPSCNNArithmeticClass() MPSCNNArithmeticClass {
	return getMPSCNNArithmeticClass()
}

type MPSCNNArithmeticClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNArithmeticClass) Alloc() MPSCNNArithmetic {
	rv := objc.Send[MPSCNNArithmetic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A parent class referenced by other neuralnetworks classes. [Full Topic]
type MPSCNNArithmetic struct {
	objectivec.Object
}

// MPSCNNArithmeticFromID constructs a [MPSCNNArithmetic] from an objc.ID.
//
// A parent class referenced by other neuralnetworks classes.
func MPSCNNArithmeticFromID(id objc.ID) MPSCNNArithmetic {
	return MPSCNNArithmetic{objectivec.Object{id}}
}
// Ensure MPSCNNArithmetic implements IMPSCNNArithmetic.
var _ IMPSCNNArithmetic = MPSCNNArithmetic{}





// An interface definition for the [MPSCNNArithmetic] class.
type IMPSCNNArithmetic interface {
	objectivec.IObject
}





// Init initializes the instance.
func (c MPSCNNArithmetic) Init() MPSCNNArithmetic {
	rv := objc.Send[MPSCNNArithmetic](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNArithmetic) Autorelease() MPSCNNArithmetic {
	rv := objc.Send[MPSCNNArithmetic](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNArithmetic creates a new MPSCNNArithmetic instance.
func NewMPSCNNArithmetic() MPSCNNArithmetic {
	class := getMPSCNNArithmeticClass()
	rv := objc.Send[MPSCNNArithmetic](objc.ID(class.class), objc.Sel("new"))
	return rv
}




































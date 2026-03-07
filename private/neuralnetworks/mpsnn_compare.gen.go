// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNCompare] class.
var (
	_MPSNNCompareClass     MPSNNCompareClass
	_MPSNNCompareClassOnce sync.Once
)

func getMPSNNCompareClass() MPSNNCompareClass {
	_MPSNNCompareClassOnce.Do(func() {
		_MPSNNCompareClass = MPSNNCompareClass{class: objc.GetClass("MPSNNCompare")}
	})
	return _MPSNNCompareClass
}

// GetMPSNNCompareClass returns the class object for MPSNNCompare.
func GetMPSNNCompareClass() MPSNNCompareClass {
	return getMPSNNCompareClass()
}

type MPSNNCompareClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNCompareClass) Alloc() MPSNNCompare {
	rv := objc.Send[MPSNNCompare](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNCompare.ComparisonType]
//   - [MPSNNCompare.SetComparisonType]
//   - [MPSNNCompare.Threshold]
//   - [MPSNNCompare.SetThreshold]
//   - [MPSNNCompare.InitWithDevice]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNCompare
type MPSNNCompare struct {
	MPSCNNArithmetic
}

// MPSNNCompareFromID constructs a [MPSNNCompare] from an objc.ID.
func MPSNNCompareFromID(id objc.ID) MPSNNCompare {
	return MPSNNCompare{MPSCNNArithmetic: MPSCNNArithmeticFromID(id)}
}
// Ensure MPSNNCompare implements IMPSNNCompare.
var _ IMPSNNCompare = MPSNNCompare{}





// An interface definition for the [MPSNNCompare] class.
//
// # Methods
//
//   - [IMPSNNCompare.ComparisonType]
//   - [IMPSNNCompare.SetComparisonType]
//   - [IMPSNNCompare.Threshold]
//   - [IMPSNNCompare.SetThreshold]
//   - [IMPSNNCompare.InitWithDevice]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNCompare
type IMPSNNCompare interface {
	IMPSCNNArithmetic

	// Topic: Methods

	ComparisonType() uint64
	SetComparisonType(value uint64)
	Threshold() float32
	SetThreshold(value float32)
	InitWithDevice(device objectivec.IObject) MPSNNCompare
}





// Init initializes the instance.
func (c MPSNNCompare) Init() MPSNNCompare {
	rv := objc.Send[MPSNNCompare](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSNNCompare) Autorelease() MPSNNCompare {
	rv := objc.Send[MPSNNCompare](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNCompare creates a new MPSNNCompare instance.
func NewMPSNNCompare() MPSNNCompare {
	class := getMPSNNCompareClass()
	rv := objc.Send[MPSNNCompare](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNCompare/initWithDevice:
func NewCompareWithDevice(device objectivec.IObject) MPSNNCompare {
	instance := getMPSNNCompareClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNCompareFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNCompare/initWithDevice:
func (c MPSNNCompare) InitWithDevice(device objectivec.IObject) MPSNNCompare {
	rv := objc.Send[MPSNNCompare](c.ID, objc.Sel("initWithDevice:"), device)
	return rv
}












// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNCompare/comparisonType
func (c MPSNNCompare) ComparisonType() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("comparisonType"))
	return rv
}
func (c MPSNNCompare) SetComparisonType(value uint64) {
	objc.Send[struct{}](c.ID, objc.Sel("setComparisonType:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNCompare/threshold
func (c MPSNNCompare) Threshold() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("threshold"))
	return rv
}
func (c MPSNNCompare) SetThreshold(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setThreshold:"), value)
}


















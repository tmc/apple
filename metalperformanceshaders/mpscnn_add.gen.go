// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSCNNAdd] class.
var (
	_MPSCNNAddClass     MPSCNNAddClass
	_MPSCNNAddClassOnce sync.Once
)

func getMPSCNNAddClass() MPSCNNAddClass {
	_MPSCNNAddClassOnce.Do(func() {
		_MPSCNNAddClass = MPSCNNAddClass{class: objc.GetClass("MPSCNNAdd")}
	})
	return _MPSCNNAddClass
}

// GetMPSCNNAddClass returns the class object for MPSCNNAdd.
func GetMPSCNNAddClass() MPSCNNAddClass {
	return getMPSCNNAddClass()
}

type MPSCNNAddClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNAddClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNAddClass) Alloc() MPSCNNAdd {
	rv := objc.Send[MPSCNNAdd](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An addition operator.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNAdd
type MPSCNNAdd struct {
	MPSCNNArithmetic
}

// MPSCNNAddFromID constructs a [MPSCNNAdd] from an objc.ID.
//
// An addition operator.
func MPSCNNAddFromID(id objc.ID) MPSCNNAdd {
	return MPSCNNAdd{MPSCNNArithmetic: MPSCNNArithmeticFromID(id)}
}

// NOTE: MPSCNNAdd adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNAdd] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNAdd
type IMPSCNNAdd interface {
	IMPSCNNArithmetic
}

// Init initializes the instance.
func (c MPSCNNAdd) Init() MPSCNNAdd {
	rv := objc.Send[MPSCNNAdd](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNAdd) Autorelease() MPSCNNAdd {
	rv := objc.Send[MPSCNNAdd](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNAdd creates a new MPSCNNAdd instance.
func NewMPSCNNAdd() MPSCNNAdd {
	class := getMPSCNNAddClass()
	rv := objc.Send[MPSCNNAdd](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewCNNAddWithCoder(aDecoder foundation.INSCoder) MPSCNNAdd {
	instance := getMPSCNNAddClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSCNNAddFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBinaryKernel/init(coder:device:)
func NewCNNAddWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSCNNAdd {
	instance := getMPSCNNAddClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSCNNAddFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNAdd/init(device:)
func NewCNNAddWithDevice(device metal.MTLDevice) MPSCNNAdd {
	instance := getMPSCNNAddClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSCNNAddFromID(rv)
}

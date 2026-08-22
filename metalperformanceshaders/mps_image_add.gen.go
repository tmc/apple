// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageAdd] class.
var (
	_MPSImageAddClass     MPSImageAddClass
	_MPSImageAddClassOnce sync.Once
)

func getMPSImageAddClass() MPSImageAddClass {
	_MPSImageAddClassOnce.Do(func() {
		_MPSImageAddClass = MPSImageAddClass{class: objc.GetClass("MPSImageAdd")}
	})
	return _MPSImageAddClass
}

// GetMPSImageAddClass returns the class object for MPSImageAdd.
func GetMPSImageAddClass() MPSImageAddClass {
	return getMPSImageAddClass()
}

type MPSImageAddClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageAddClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageAddClass) Alloc() MPSImageAdd {
	rv := objc.Send[MPSImageAdd](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A filter that returns the element-wise sum of its two input images.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAdd
type MPSImageAdd struct {
	MPSImageArithmetic
}

// MPSImageAddFromID constructs a [MPSImageAdd] from an objc.ID.
//
// A filter that returns the element-wise sum of its two input images.
func MPSImageAddFromID(id objc.ID) MPSImageAdd {
	return MPSImageAdd{MPSImageArithmetic: MPSImageArithmeticFromID(id)}
}

// NOTE: MPSImageAdd adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageAdd] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAdd
type IMPSImageAdd interface {
	IMPSImageArithmetic
}

// Init initializes the instance.
func (i MPSImageAdd) Init() MPSImageAdd {
	rv := objc.Send[MPSImageAdd](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageAdd) Autorelease() MPSImageAdd {
	rv := objc.Send[MPSImageAdd](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageAdd creates a new MPSImageAdd instance.
func NewMPSImageAdd() MPSImageAdd {
	class := getMPSImageAddClass()
	rv := objc.Send[MPSImageAdd](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageAddWithCoder(aDecoder foundation.INSCoder) MPSImageAdd {
	instance := getMPSImageAddClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageAddFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/init(coder:device:)
func NewImageAddWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageAdd {
	instance := getMPSImageAddClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageAddFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAdd/init(device:)
func NewImageAddWithDevice(device metal.MTLDevice) MPSImageAdd {
	instance := getMPSImageAddClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageAddFromID(rv)
}

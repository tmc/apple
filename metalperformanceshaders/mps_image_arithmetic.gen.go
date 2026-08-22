// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSImageArithmetic] class.
var (
	_MPSImageArithmeticClass     MPSImageArithmeticClass
	_MPSImageArithmeticClassOnce sync.Once
)

func getMPSImageArithmeticClass() MPSImageArithmeticClass {
	_MPSImageArithmeticClassOnce.Do(func() {
		_MPSImageArithmeticClass = MPSImageArithmeticClass{class: objc.GetClass("MPSImageArithmetic")}
	})
	return _MPSImageArithmeticClass
}

// GetMPSImageArithmeticClass returns the class object for MPSImageArithmetic.
func GetMPSImageArithmeticClass() MPSImageArithmeticClass {
	return getMPSImageArithmeticClass()
}

type MPSImageArithmeticClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSImageArithmeticClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSImageArithmeticClass) Alloc() MPSImageArithmetic {
	rv := objc.Send[MPSImageArithmetic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Base class for basic arithmetic nodes
//
// # Instance Properties
//
//   - [MPSImageArithmetic.Bias]
//   - [MPSImageArithmetic.SetBias]
//   - [MPSImageArithmetic.PrimaryScale]
//   - [MPSImageArithmetic.SetPrimaryScale]
//   - [MPSImageArithmetic.PrimaryStrideInPixels]
//   - [MPSImageArithmetic.SetPrimaryStrideInPixels]
//   - [MPSImageArithmetic.SecondaryScale]
//   - [MPSImageArithmetic.SetSecondaryScale]
//   - [MPSImageArithmetic.SecondaryStrideInPixels]
//   - [MPSImageArithmetic.SetSecondaryStrideInPixels]
//   - [MPSImageArithmetic.MaximumValue]
//   - [MPSImageArithmetic.SetMaximumValue]
//   - [MPSImageArithmetic.MinimumValue]
//   - [MPSImageArithmetic.SetMinimumValue]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageArithmetic
type MPSImageArithmetic struct {
	MPSBinaryImageKernel
}

// MPSImageArithmeticFromID constructs a [MPSImageArithmetic] from an objc.ID.
//
// Base class for basic arithmetic nodes
func MPSImageArithmeticFromID(id objc.ID) MPSImageArithmetic {
	return MPSImageArithmetic{MPSBinaryImageKernel: MPSBinaryImageKernelFromID(id)}
}

// NOTE: MPSImageArithmetic adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSImageArithmetic] class.
//
// # Instance Properties
//
//   - [IMPSImageArithmetic.Bias]
//   - [IMPSImageArithmetic.SetBias]
//   - [IMPSImageArithmetic.PrimaryScale]
//   - [IMPSImageArithmetic.SetPrimaryScale]
//   - [IMPSImageArithmetic.PrimaryStrideInPixels]
//   - [IMPSImageArithmetic.SetPrimaryStrideInPixels]
//   - [IMPSImageArithmetic.SecondaryScale]
//   - [IMPSImageArithmetic.SetSecondaryScale]
//   - [IMPSImageArithmetic.SecondaryStrideInPixels]
//   - [IMPSImageArithmetic.SetSecondaryStrideInPixels]
//   - [IMPSImageArithmetic.MaximumValue]
//   - [IMPSImageArithmetic.SetMaximumValue]
//   - [IMPSImageArithmetic.MinimumValue]
//   - [IMPSImageArithmetic.SetMinimumValue]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageArithmetic
type IMPSImageArithmetic interface {
	IMPSBinaryImageKernel

	// Topic: Instance Properties

	Bias() float32
	SetBias(value float32)
	PrimaryScale() float32
	SetPrimaryScale(value float32)
	PrimaryStrideInPixels() metal.MTLSize
	SetPrimaryStrideInPixels(value metal.MTLSize)
	SecondaryScale() float32
	SetSecondaryScale(value float32)
	SecondaryStrideInPixels() metal.MTLSize
	SetSecondaryStrideInPixels(value metal.MTLSize)
	MaximumValue() float32
	SetMaximumValue(value float32)
	MinimumValue() float32
	SetMinimumValue(value float32)
}

// Init initializes the instance.
func (i MPSImageArithmetic) Init() MPSImageArithmetic {
	rv := objc.Send[MPSImageArithmetic](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSImageArithmetic) Autorelease() MPSImageArithmetic {
	rv := objc.Send[MPSImageArithmetic](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSImageArithmetic creates a new MPSImageArithmetic instance.
func NewMPSImageArithmetic() MPSImageArithmetic {
	class := getMPSImageArithmeticClass()
	rv := objc.Send[MPSImageArithmetic](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewImageArithmeticWithCoder(aDecoder foundation.INSCoder) MPSImageArithmetic {
	instance := getMPSImageArithmeticClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSImageArithmeticFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/init(coder:device:)
func NewImageArithmeticWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSImageArithmetic {
	instance := getMPSImageArithmeticClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSImageArithmeticFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSBinaryImageKernel/init(device:)
func NewImageArithmeticWithDevice(device metal.MTLDevice) MPSImageArithmetic {
	instance := getMPSImageArithmeticClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSImageArithmeticFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageArithmetic/bias
func (i MPSImageArithmetic) Bias() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("bias"))
	return rv
}
func (i MPSImageArithmetic) SetBias(value float32) {
	objc.Send[struct{}](i.ID, objc.Sel("setBias:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageArithmetic/primaryScale
func (i MPSImageArithmetic) PrimaryScale() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("primaryScale"))
	return rv
}
func (i MPSImageArithmetic) SetPrimaryScale(value float32) {
	objc.Send[struct{}](i.ID, objc.Sel("setPrimaryScale:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageArithmetic/primaryStrideInPixels
func (i MPSImageArithmetic) PrimaryStrideInPixels() metal.MTLSize {
	rv := objc.Send[metal.MTLSize](i.ID, objc.Sel("primaryStrideInPixels"))
	return metal.MTLSize(rv)
}
func (i MPSImageArithmetic) SetPrimaryStrideInPixels(value metal.MTLSize) {
	objc.Send[struct{}](i.ID, objc.Sel("setPrimaryStrideInPixels:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageArithmetic/secondaryScale
func (i MPSImageArithmetic) SecondaryScale() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("secondaryScale"))
	return rv
}
func (i MPSImageArithmetic) SetSecondaryScale(value float32) {
	objc.Send[struct{}](i.ID, objc.Sel("setSecondaryScale:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageArithmetic/secondaryStrideInPixels
func (i MPSImageArithmetic) SecondaryStrideInPixels() metal.MTLSize {
	rv := objc.Send[metal.MTLSize](i.ID, objc.Sel("secondaryStrideInPixels"))
	return metal.MTLSize(rv)
}
func (i MPSImageArithmetic) SetSecondaryStrideInPixels(value metal.MTLSize) {
	objc.Send[struct{}](i.ID, objc.Sel("setSecondaryStrideInPixels:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageArithmetic/maximumValue
func (i MPSImageArithmetic) MaximumValue() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("maximumValue"))
	return rv
}
func (i MPSImageArithmetic) SetMaximumValue(value float32) {
	objc.Send[struct{}](i.ID, objc.Sel("setMaximumValue:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageArithmetic/minimumValue
func (i MPSImageArithmetic) MinimumValue() float32 {
	rv := objc.Send[float32](i.ID, objc.Sel("minimumValue"))
	return rv
}
func (i MPSImageArithmetic) SetMinimumValue(value float32) {
	objc.Send[struct{}](i.ID, objc.Sel("setMinimumValue:"), value)
}

// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSNNLocalCorrelation] class.
var (
	_MPSNNLocalCorrelationClass     MPSNNLocalCorrelationClass
	_MPSNNLocalCorrelationClassOnce sync.Once
)

func getMPSNNLocalCorrelationClass() MPSNNLocalCorrelationClass {
	_MPSNNLocalCorrelationClassOnce.Do(func() {
		_MPSNNLocalCorrelationClass = MPSNNLocalCorrelationClass{class: objc.GetClass("MPSNNLocalCorrelation")}
	})
	return _MPSNNLocalCorrelationClass
}

// GetMPSNNLocalCorrelationClass returns the class object for MPSNNLocalCorrelation.
func GetMPSNNLocalCorrelationClass() MPSNNLocalCorrelationClass {
	return getMPSNNLocalCorrelationClass()
}

type MPSNNLocalCorrelationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNLocalCorrelationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNLocalCorrelationClass) Alloc() MPSNNLocalCorrelation {
	rv := objc.Send[MPSNNLocalCorrelation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNNLocalCorrelation.InitWithDeviceWindowInXWindowInYStrideInXStrideInY]
//
// # Instance Properties
//
//   - [MPSNNLocalCorrelation.StrideInX]
//   - [MPSNNLocalCorrelation.SetStrideInX]
//   - [MPSNNLocalCorrelation.StrideInY]
//   - [MPSNNLocalCorrelation.SetStrideInY]
//   - [MPSNNLocalCorrelation.WindowInX]
//   - [MPSNNLocalCorrelation.SetWindowInX]
//   - [MPSNNLocalCorrelation.WindowInY]
//   - [MPSNNLocalCorrelation.SetWindowInY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLocalCorrelation
type MPSNNLocalCorrelation struct {
	MPSNNReduceBinary
}

// MPSNNLocalCorrelationFromID constructs a [MPSNNLocalCorrelation] from an objc.ID.
func MPSNNLocalCorrelationFromID(id objc.ID) MPSNNLocalCorrelation {
	return MPSNNLocalCorrelation{MPSNNReduceBinary: MPSNNReduceBinaryFromID(id)}
}

// NOTE: MPSNNLocalCorrelation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNLocalCorrelation] class.
//
// # Initializers
//
//   - [IMPSNNLocalCorrelation.InitWithDeviceWindowInXWindowInYStrideInXStrideInY]
//
// # Instance Properties
//
//   - [IMPSNNLocalCorrelation.StrideInX]
//   - [IMPSNNLocalCorrelation.SetStrideInX]
//   - [IMPSNNLocalCorrelation.StrideInY]
//   - [IMPSNNLocalCorrelation.SetStrideInY]
//   - [IMPSNNLocalCorrelation.WindowInX]
//   - [IMPSNNLocalCorrelation.SetWindowInX]
//   - [IMPSNNLocalCorrelation.WindowInY]
//   - [IMPSNNLocalCorrelation.SetWindowInY]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLocalCorrelation
type IMPSNNLocalCorrelation interface {
	IMPSNNReduceBinary

	// Topic: Initializers

	InitWithDeviceWindowInXWindowInYStrideInXStrideInY(device metal.MTLDevice, windowInX uint, windowInY uint, strideInX uint, strideInY uint) MPSNNLocalCorrelation

	// Topic: Instance Properties

	StrideInX() uint
	SetStrideInX(value uint)
	StrideInY() uint
	SetStrideInY(value uint)
	WindowInX() uint
	SetWindowInX(value uint)
	WindowInY() uint
	SetWindowInY(value uint)
}

// Init initializes the instance.
func (l MPSNNLocalCorrelation) Init() MPSNNLocalCorrelation {
	rv := objc.Send[MPSNNLocalCorrelation](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MPSNNLocalCorrelation) Autorelease() MPSNNLocalCorrelation {
	rv := objc.Send[MPSNNLocalCorrelation](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNLocalCorrelation creates a new MPSNNLocalCorrelation instance.
func NewMPSNNLocalCorrelation() MPSNNLocalCorrelation {
	class := getMPSNNLocalCorrelationClass()
	rv := objc.Send[MPSNNLocalCorrelation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewLocalCorrelationWithCoder(aDecoder foundation.INSCoder) MPSNNLocalCorrelation {
	instance := getMPSNNLocalCorrelationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNLocalCorrelationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLocalCorrelation/init(coder:device:)
func NewLocalCorrelationWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNLocalCorrelation {
	instance := getMPSNNLocalCorrelationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNLocalCorrelationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLocalCorrelation/init(device:)
func NewLocalCorrelationWithDevice(device metal.MTLDevice) MPSNNLocalCorrelation {
	instance := getMPSNNLocalCorrelationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNLocalCorrelationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLocalCorrelation/init(device:windowInX:windowInY:strideInX:strideInY:)
func NewLocalCorrelationWithDeviceWindowInXWindowInYStrideInXStrideInY(device metal.MTLDevice, windowInX uint, windowInY uint, strideInX uint, strideInY uint) MPSNNLocalCorrelation {
	instance := getMPSNNLocalCorrelationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:windowInX:windowInY:strideInX:strideInY:"), device, windowInX, windowInY, strideInX, strideInY)
	return MPSNNLocalCorrelationFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLocalCorrelation/init(device:windowInX:windowInY:strideInX:strideInY:)
func (l MPSNNLocalCorrelation) InitWithDeviceWindowInXWindowInYStrideInXStrideInY(device metal.MTLDevice, windowInX uint, windowInY uint, strideInX uint, strideInY uint) MPSNNLocalCorrelation {
	rv := objc.Send[MPSNNLocalCorrelation](l.ID, objc.Sel("initWithDevice:windowInX:windowInY:strideInX:strideInY:"), device, windowInX, windowInY, strideInX, strideInY)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLocalCorrelation/strideInX
func (l MPSNNLocalCorrelation) StrideInX() uint {
	rv := objc.Send[uint](l.ID, objc.Sel("strideInX"))
	return rv
}
func (l MPSNNLocalCorrelation) SetStrideInX(value uint) {
	objc.Send[struct{}](l.ID, objc.Sel("setStrideInX:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLocalCorrelation/strideInY
func (l MPSNNLocalCorrelation) StrideInY() uint {
	rv := objc.Send[uint](l.ID, objc.Sel("strideInY"))
	return rv
}
func (l MPSNNLocalCorrelation) SetStrideInY(value uint) {
	objc.Send[struct{}](l.ID, objc.Sel("setStrideInY:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLocalCorrelation/windowInX
func (l MPSNNLocalCorrelation) WindowInX() uint {
	rv := objc.Send[uint](l.ID, objc.Sel("windowInX"))
	return rv
}
func (l MPSNNLocalCorrelation) SetWindowInX(value uint) {
	objc.Send[struct{}](l.ID, objc.Sel("setWindowInX:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLocalCorrelation/windowInY
func (l MPSNNLocalCorrelation) WindowInY() uint {
	rv := objc.Send[uint](l.ID, objc.Sel("windowInY"))
	return rv
}
func (l MPSNNLocalCorrelation) SetWindowInY(value uint) {
	objc.Send[struct{}](l.ID, objc.Sel("setWindowInY:"), value)
}

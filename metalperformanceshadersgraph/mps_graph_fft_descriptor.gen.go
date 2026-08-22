// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphFFTDescriptor] class.
var (
	_MPSGraphFFTDescriptorClass     MPSGraphFFTDescriptorClass
	_MPSGraphFFTDescriptorClassOnce sync.Once
)

func getMPSGraphFFTDescriptorClass() MPSGraphFFTDescriptorClass {
	_MPSGraphFFTDescriptorClassOnce.Do(func() {
		_MPSGraphFFTDescriptorClass = MPSGraphFFTDescriptorClass{class: objc.GetClass("MPSGraphFFTDescriptor")}
	})
	return _MPSGraphFFTDescriptorClass
}

// GetMPSGraphFFTDescriptorClass returns the class object for MPSGraphFFTDescriptor.
func GetMPSGraphFFTDescriptorClass() MPSGraphFFTDescriptorClass {
	return getMPSGraphFFTDescriptorClass()
}

type MPSGraphFFTDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphFFTDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphFFTDescriptorClass) Alloc() MPSGraphFFTDescriptor {
	rv := objc.Send[MPSGraphFFTDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The class that defines the parameters for a fast Fourier transform (FFT)
// operation.
//
// # Overview
//
// Use this descriptor with
// [MPSGraph.FastFourierTransformWithTensorAxesDescriptorName],
// [MPSGraph.RealToHermiteanFFTWithTensorAxesTensorDescriptorName], and
// [MPSGraph.HermiteanToRealFFTWithTensorAxesTensorDescriptorName] methods.
//
// # Instance Properties
//
//   - [MPSGraphFFTDescriptor.Inverse]: A Boolean-valued parameter that defines the phase factor sign for Fourier transforms.
//   - [MPSGraphFFTDescriptor.SetInverse]
//   - [MPSGraphFFTDescriptor.RoundToOddHermitean]: A parameter which controls how graph rounds the output tensor size for a Hermitean-to-real Fourier transform.
//   - [MPSGraphFFTDescriptor.SetRoundToOddHermitean]
//   - [MPSGraphFFTDescriptor.ScalingMode]: The scaling mode of the fast fourier transform (FFT) operation.
//   - [MPSGraphFFTDescriptor.SetScalingMode]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTDescriptor
type MPSGraphFFTDescriptor struct {
	MPSGraphObject
}

// MPSGraphFFTDescriptorFromID constructs a [MPSGraphFFTDescriptor] from an objc.ID.
//
// The class that defines the parameters for a fast Fourier transform (FFT)
// operation.
func MPSGraphFFTDescriptorFromID(id objc.ID) MPSGraphFFTDescriptor {
	return MPSGraphFFTDescriptor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphFFTDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphFFTDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGraphFFTDescriptor.Inverse]: A Boolean-valued parameter that defines the phase factor sign for Fourier transforms.
//   - [IMPSGraphFFTDescriptor.SetInverse]
//   - [IMPSGraphFFTDescriptor.RoundToOddHermitean]: A parameter which controls how graph rounds the output tensor size for a Hermitean-to-real Fourier transform.
//   - [IMPSGraphFFTDescriptor.SetRoundToOddHermitean]
//   - [IMPSGraphFFTDescriptor.ScalingMode]: The scaling mode of the fast fourier transform (FFT) operation.
//   - [IMPSGraphFFTDescriptor.SetScalingMode]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTDescriptor
type IMPSGraphFFTDescriptor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// A Boolean-valued parameter that defines the phase factor sign for Fourier transforms.
	Inverse() bool
	SetInverse(value bool)
	// A parameter which controls how graph rounds the output tensor size for a Hermitean-to-real Fourier transform.
	RoundToOddHermitean() bool
	SetRoundToOddHermitean(value bool)
	// The scaling mode of the fast fourier transform (FFT) operation.
	ScalingMode() MPSGraphFFTScalingMode
	SetScalingMode(value MPSGraphFFTScalingMode)
}

// Init initializes the instance.
func (g MPSGraphFFTDescriptor) Init() MPSGraphFFTDescriptor {
	rv := objc.Send[MPSGraphFFTDescriptor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphFFTDescriptor) Autorelease() MPSGraphFFTDescriptor {
	rv := objc.Send[MPSGraphFFTDescriptor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphFFTDescriptor creates a new MPSGraphFFTDescriptor instance.
func NewMPSGraphFFTDescriptor() MPSGraphFFTDescriptor {
	class := getMPSGraphFFTDescriptorClass()
	rv := objc.Send[MPSGraphFFTDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a fast Fourier transform descriptor with default parameter values.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTDescriptor/descriptor
func (_MPSGraphFFTDescriptorClass MPSGraphFFTDescriptorClass) Descriptor() MPSGraphFFTDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSGraphFFTDescriptorClass.class), objc.Sel("descriptor"))
	return MPSGraphFFTDescriptorFromID(rv)
}

// A Boolean-valued parameter that defines the phase factor sign for Fourier
// transforms.
//
// # Discussion
//
// When set to [YES] graph uses the positive phase factor: `exp(+i 2Pi mu nu /
// n)`, when computing the (inverse) Fourier transform. Otherwise MPSGraph
// uses the negative phase factor: `exp(-i 2Pi mu nu / n)`, when computing the
// Fourier transform. Default value: [NO].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTDescriptor/inverse
func (g MPSGraphFFTDescriptor) Inverse() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("inverse"))
	return rv
}
func (g MPSGraphFFTDescriptor) SetInverse(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setInverse:"), value)
}

// A parameter which controls how graph rounds the output tensor size for a
// Hermitean-to-real Fourier transform.
//
// # Discussion
//
// If set to [YES] then MPSGraph rounds the last output dimension of the
// result tensor in
// [MPSGraph.HermiteanToRealFFTWithTensorAxesTensorDescriptorName] to an odd
// value. Has no effect in the other Fourier transform operations. Default
// value: [NO].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTDescriptor/roundToOddHermitean
func (g MPSGraphFFTDescriptor) RoundToOddHermitean() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("roundToOddHermitean"))
	return rv
}
func (g MPSGraphFFTDescriptor) SetRoundToOddHermitean(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setRoundToOddHermitean:"), value)
}

// The scaling mode of the fast fourier transform (FFT) operation.
//
// # Discussion
//
// Note that the scaling mode is independent from the phase factor. Default
// value: [MPSGraphFFTScalingModeNone].
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTDescriptor/scalingMode
func (g MPSGraphFFTDescriptor) ScalingMode() MPSGraphFFTScalingMode {
	rv := objc.Send[MPSGraphFFTScalingMode](g.ID, objc.Sel("scalingMode"))
	return MPSGraphFFTScalingMode(rv)
}
func (g MPSGraphFFTDescriptor) SetScalingMode(value MPSGraphFFTScalingMode) {
	objc.Send[struct{}](g.ID, objc.Sel("setScalingMode:"), value)
}

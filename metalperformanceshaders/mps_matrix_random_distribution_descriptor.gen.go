// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSMatrixRandomDistributionDescriptor] class.
var (
	_MPSMatrixRandomDistributionDescriptorClass     MPSMatrixRandomDistributionDescriptorClass
	_MPSMatrixRandomDistributionDescriptorClassOnce sync.Once
)

func getMPSMatrixRandomDistributionDescriptorClass() MPSMatrixRandomDistributionDescriptorClass {
	_MPSMatrixRandomDistributionDescriptorClassOnce.Do(func() {
		_MPSMatrixRandomDistributionDescriptorClass = MPSMatrixRandomDistributionDescriptorClass{class: objc.GetClass("MPSMatrixRandomDistributionDescriptor")}
	})
	return _MPSMatrixRandomDistributionDescriptorClass
}

// GetMPSMatrixRandomDistributionDescriptorClass returns the class object for MPSMatrixRandomDistributionDescriptor.
func GetMPSMatrixRandomDistributionDescriptorClass() MPSMatrixRandomDistributionDescriptorClass {
	return getMPSMatrixRandomDistributionDescriptorClass()
}

type MPSMatrixRandomDistributionDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSMatrixRandomDistributionDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSMatrixRandomDistributionDescriptorClass) Alloc() MPSMatrixRandomDistributionDescriptor {
	rv := objc.Send[MPSMatrixRandomDistributionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MPSMatrixRandomDistributionDescriptor.DistributionType]
//   - [MPSMatrixRandomDistributionDescriptor.SetDistributionType]
//   - [MPSMatrixRandomDistributionDescriptor.Maximum]
//   - [MPSMatrixRandomDistributionDescriptor.SetMaximum]
//   - [MPSMatrixRandomDistributionDescriptor.Mean]
//   - [MPSMatrixRandomDistributionDescriptor.SetMean]
//   - [MPSMatrixRandomDistributionDescriptor.Minimum]
//   - [MPSMatrixRandomDistributionDescriptor.SetMinimum]
//   - [MPSMatrixRandomDistributionDescriptor.StandardDeviation]
//   - [MPSMatrixRandomDistributionDescriptor.SetStandardDeviation]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor
type MPSMatrixRandomDistributionDescriptor struct {
	objectivec.Object
}

// MPSMatrixRandomDistributionDescriptorFromID constructs a [MPSMatrixRandomDistributionDescriptor] from an objc.ID.
func MPSMatrixRandomDistributionDescriptorFromID(id objc.ID) MPSMatrixRandomDistributionDescriptor {
	return MPSMatrixRandomDistributionDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MPSMatrixRandomDistributionDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSMatrixRandomDistributionDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSMatrixRandomDistributionDescriptor.DistributionType]
//   - [IMPSMatrixRandomDistributionDescriptor.SetDistributionType]
//   - [IMPSMatrixRandomDistributionDescriptor.Maximum]
//   - [IMPSMatrixRandomDistributionDescriptor.SetMaximum]
//   - [IMPSMatrixRandomDistributionDescriptor.Mean]
//   - [IMPSMatrixRandomDistributionDescriptor.SetMean]
//   - [IMPSMatrixRandomDistributionDescriptor.Minimum]
//   - [IMPSMatrixRandomDistributionDescriptor.SetMinimum]
//   - [IMPSMatrixRandomDistributionDescriptor.StandardDeviation]
//   - [IMPSMatrixRandomDistributionDescriptor.SetStandardDeviation]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor
type IMPSMatrixRandomDistributionDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	DistributionType() MPSMatrixRandomDistribution
	SetDistributionType(value MPSMatrixRandomDistribution)
	Maximum() float32
	SetMaximum(value float32)
	Mean() float32
	SetMean(value float32)
	Minimum() float32
	SetMinimum(value float32)
	StandardDeviation() float32
	SetStandardDeviation(value float32)
}

// Init initializes the instance.
func (m MPSMatrixRandomDistributionDescriptor) Init() MPSMatrixRandomDistributionDescriptor {
	rv := objc.Send[MPSMatrixRandomDistributionDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSMatrixRandomDistributionDescriptor) Autorelease() MPSMatrixRandomDistributionDescriptor {
	rv := objc.Send[MPSMatrixRandomDistributionDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSMatrixRandomDistributionDescriptor creates a new MPSMatrixRandomDistributionDescriptor instance.
func NewMPSMatrixRandomDistributionDescriptor() MPSMatrixRandomDistributionDescriptor {
	class := getMPSMatrixRandomDistributionDescriptorClass()
	rv := objc.Send[MPSMatrixRandomDistributionDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor/default()
func (_MPSMatrixRandomDistributionDescriptorClass MPSMatrixRandomDistributionDescriptorClass) DefaultDistributionDescriptor() MPSMatrixRandomDistributionDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSMatrixRandomDistributionDescriptorClass.class), objc.Sel("defaultDistributionDescriptor"))
	return MPSMatrixRandomDistributionDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor/normalDistributionDescriptor(withMean:standardDeviation:)
func (_MPSMatrixRandomDistributionDescriptorClass MPSMatrixRandomDistributionDescriptorClass) NormalDistributionDescriptorWithMeanStandardDeviation(mean float32, standardDeviation float32) MPSMatrixRandomDistributionDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSMatrixRandomDistributionDescriptorClass.class), objc.Sel("normalDistributionDescriptorWithMean:standardDeviation:"), mean, standardDeviation)
	return MPSMatrixRandomDistributionDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor/normalDistributionDescriptor(withMean:standardDeviation:minimum:maximum:)
func (_MPSMatrixRandomDistributionDescriptorClass MPSMatrixRandomDistributionDescriptorClass) NormalDistributionDescriptorWithMeanStandardDeviationMinimumMaximum(mean float32, standardDeviation float32, minimum float32, maximum float32) MPSMatrixRandomDistributionDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSMatrixRandomDistributionDescriptorClass.class), objc.Sel("normalDistributionDescriptorWithMean:standardDeviation:minimum:maximum:"), mean, standardDeviation, minimum, maximum)
	return MPSMatrixRandomDistributionDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor/uniformDistributionDescriptor(withMinimum:maximum:)
func (_MPSMatrixRandomDistributionDescriptorClass MPSMatrixRandomDistributionDescriptorClass) UniformDistributionDescriptorWithMinimumMaximum(minimum float32, maximum float32) MPSMatrixRandomDistributionDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSMatrixRandomDistributionDescriptorClass.class), objc.Sel("uniformDistributionDescriptorWithMinimum:maximum:"), minimum, maximum)
	return MPSMatrixRandomDistributionDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor/distributionType
func (m MPSMatrixRandomDistributionDescriptor) DistributionType() MPSMatrixRandomDistribution {
	rv := objc.Send[MPSMatrixRandomDistribution](m.ID, objc.Sel("distributionType"))
	return MPSMatrixRandomDistribution(rv)
}
func (m MPSMatrixRandomDistributionDescriptor) SetDistributionType(value MPSMatrixRandomDistribution) {
	objc.Send[struct{}](m.ID, objc.Sel("setDistributionType:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor/maximum
func (m MPSMatrixRandomDistributionDescriptor) Maximum() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("maximum"))
	return rv
}
func (m MPSMatrixRandomDistributionDescriptor) SetMaximum(value float32) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaximum:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor/mean
func (m MPSMatrixRandomDistributionDescriptor) Mean() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("mean"))
	return rv
}
func (m MPSMatrixRandomDistributionDescriptor) SetMean(value float32) {
	objc.Send[struct{}](m.ID, objc.Sel("setMean:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor/minimum
func (m MPSMatrixRandomDistributionDescriptor) Minimum() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("minimum"))
	return rv
}
func (m MPSMatrixRandomDistributionDescriptor) SetMinimum(value float32) {
	objc.Send[struct{}](m.ID, objc.Sel("setMinimum:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor/standardDeviation
func (m MPSMatrixRandomDistributionDescriptor) StandardDeviation() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("standardDeviation"))
	return rv
}
func (m MPSMatrixRandomDistributionDescriptor) SetStandardDeviation(value float32) {
	objc.Send[struct{}](m.ID, objc.Sel("setStandardDeviation:"), value)
}

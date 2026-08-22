// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphRandomOpDescriptor] class.
var (
	_MPSGraphRandomOpDescriptorClass     MPSGraphRandomOpDescriptorClass
	_MPSGraphRandomOpDescriptorClassOnce sync.Once
)

func getMPSGraphRandomOpDescriptorClass() MPSGraphRandomOpDescriptorClass {
	_MPSGraphRandomOpDescriptorClassOnce.Do(func() {
		_MPSGraphRandomOpDescriptorClass = MPSGraphRandomOpDescriptorClass{class: objc.GetClass("MPSGraphRandomOpDescriptor")}
	})
	return _MPSGraphRandomOpDescriptorClass
}

// GetMPSGraphRandomOpDescriptorClass returns the class object for MPSGraphRandomOpDescriptor.
func GetMPSGraphRandomOpDescriptorClass() MPSGraphRandomOpDescriptorClass {
	return getMPSGraphRandomOpDescriptorClass()
}

type MPSGraphRandomOpDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphRandomOpDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphRandomOpDescriptorClass) Alloc() MPSGraphRandomOpDescriptor {
	rv := objc.Send[MPSGraphRandomOpDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that describes the random operation.
//
// # Instance Properties
//
//   - [MPSGraphRandomOpDescriptor.DataType]: The data type of the generated result values.
//   - [MPSGraphRandomOpDescriptor.SetDataType]
//   - [MPSGraphRandomOpDescriptor.Distribution]: The type of distribution to draw samples from. See MPSGraphRandomDistribution.
//   - [MPSGraphRandomOpDescriptor.SetDistribution]
//   - [MPSGraphRandomOpDescriptor.Max]: The upper range of the distribution.
//   - [MPSGraphRandomOpDescriptor.SetMax]
//   - [MPSGraphRandomOpDescriptor.MaxInteger]: The upper range of the distribution.
//   - [MPSGraphRandomOpDescriptor.SetMaxInteger]
//   - [MPSGraphRandomOpDescriptor.Mean]: The mean of the distribution.
//   - [MPSGraphRandomOpDescriptor.SetMean]
//   - [MPSGraphRandomOpDescriptor.Min]: The lower range of the distribution.
//   - [MPSGraphRandomOpDescriptor.SetMin]
//   - [MPSGraphRandomOpDescriptor.MinInteger]: The lower range of the distribution.
//   - [MPSGraphRandomOpDescriptor.SetMinInteger]
//   - [MPSGraphRandomOpDescriptor.SamplingMethod]: The sampling method of the distribution.
//   - [MPSGraphRandomOpDescriptor.SetSamplingMethod]
//   - [MPSGraphRandomOpDescriptor.StandardDeviation]: The standard deviation of the distribution.
//   - [MPSGraphRandomOpDescriptor.SetStandardDeviation]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor
type MPSGraphRandomOpDescriptor struct {
	MPSGraphObject
}

// MPSGraphRandomOpDescriptorFromID constructs a [MPSGraphRandomOpDescriptor] from an objc.ID.
//
// A class that describes the random operation.
func MPSGraphRandomOpDescriptorFromID(id objc.ID) MPSGraphRandomOpDescriptor {
	return MPSGraphRandomOpDescriptor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphRandomOpDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphRandomOpDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGraphRandomOpDescriptor.DataType]: The data type of the generated result values.
//   - [IMPSGraphRandomOpDescriptor.SetDataType]
//   - [IMPSGraphRandomOpDescriptor.Distribution]: The type of distribution to draw samples from. See MPSGraphRandomDistribution.
//   - [IMPSGraphRandomOpDescriptor.SetDistribution]
//   - [IMPSGraphRandomOpDescriptor.Max]: The upper range of the distribution.
//   - [IMPSGraphRandomOpDescriptor.SetMax]
//   - [IMPSGraphRandomOpDescriptor.MaxInteger]: The upper range of the distribution.
//   - [IMPSGraphRandomOpDescriptor.SetMaxInteger]
//   - [IMPSGraphRandomOpDescriptor.Mean]: The mean of the distribution.
//   - [IMPSGraphRandomOpDescriptor.SetMean]
//   - [IMPSGraphRandomOpDescriptor.Min]: The lower range of the distribution.
//   - [IMPSGraphRandomOpDescriptor.SetMin]
//   - [IMPSGraphRandomOpDescriptor.MinInteger]: The lower range of the distribution.
//   - [IMPSGraphRandomOpDescriptor.SetMinInteger]
//   - [IMPSGraphRandomOpDescriptor.SamplingMethod]: The sampling method of the distribution.
//   - [IMPSGraphRandomOpDescriptor.SetSamplingMethod]
//   - [IMPSGraphRandomOpDescriptor.StandardDeviation]: The standard deviation of the distribution.
//   - [IMPSGraphRandomOpDescriptor.SetStandardDeviation]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor
type IMPSGraphRandomOpDescriptor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// The data type of the generated result values.
	DataType() uint32
	SetDataType(value uint32)
	// The type of distribution to draw samples from. See MPSGraphRandomDistribution.
	Distribution() MPSGraphRandomDistribution
	SetDistribution(value MPSGraphRandomDistribution)
	// The upper range of the distribution.
	Max() float32
	SetMax(value float32)
	// The upper range of the distribution.
	MaxInteger() int
	SetMaxInteger(value int)
	// The mean of the distribution.
	Mean() float32
	SetMean(value float32)
	// The lower range of the distribution.
	Min() float32
	SetMin(value float32)
	// The lower range of the distribution.
	MinInteger() int
	SetMinInteger(value int)
	// The sampling method of the distribution.
	SamplingMethod() MPSGraphRandomNormalSamplingMethod
	SetSamplingMethod(value MPSGraphRandomNormalSamplingMethod)
	// The standard deviation of the distribution.
	StandardDeviation() float32
	SetStandardDeviation(value float32)
}

// Init initializes the instance.
func (g MPSGraphRandomOpDescriptor) Init() MPSGraphRandomOpDescriptor {
	rv := objc.Send[MPSGraphRandomOpDescriptor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphRandomOpDescriptor) Autorelease() MPSGraphRandomOpDescriptor {
	rv := objc.Send[MPSGraphRandomOpDescriptor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphRandomOpDescriptor creates a new MPSGraphRandomOpDescriptor instance.
func NewMPSGraphRandomOpDescriptor() MPSGraphRandomOpDescriptor {
	class := getMPSGraphRandomOpDescriptorClass()
	rv := objc.Send[MPSGraphRandomOpDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Class method to initialize a distribution descriptor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/init(distribution:dataType:)
func NewGraphRandomOpDescriptorWithDistributionDataType(distribution MPSGraphRandomDistribution, dataType uint32) MPSGraphRandomOpDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSGraphRandomOpDescriptorClass().class), objc.Sel("descriptorWithDistribution:dataType:"), distribution, dataType)
	return MPSGraphRandomOpDescriptorFromID(rv)
}

// The data type of the generated result values.
//
// # Discussion
//
// When sampling from the uniform distribution, valid types are
// MPSDataTypeFloat16, MPSDataTypeFloat32, and MPSDataTypeInt32. When sampling
// from the normal or truncated normal distribution, valid types are
// MPSDataTypeFloat16 and MPSDataTypeFloat32.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/dataType
func (g MPSGraphRandomOpDescriptor) DataType() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("dataType"))
	return rv
}
func (g MPSGraphRandomOpDescriptor) SetDataType(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setDataType:"), value)
}

// The type of distribution to draw samples from. See
// MPSGraphRandomDistribution.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/distribution
func (g MPSGraphRandomOpDescriptor) Distribution() MPSGraphRandomDistribution {
	rv := objc.Send[MPSGraphRandomDistribution](g.ID, objc.Sel("distribution"))
	return MPSGraphRandomDistribution(rv)
}
func (g MPSGraphRandomOpDescriptor) SetDistribution(value MPSGraphRandomDistribution) {
	objc.Send[struct{}](g.ID, objc.Sel("setDistribution:"), value)
}

// The upper range of the distribution.
//
// # Discussion
//
// This value is used for Uniform distributions with float data types and
// Truncated Normal disributions. Defaults to 1 for uniform distributions and
// 2 for normal distributions.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/max
func (g MPSGraphRandomOpDescriptor) Max() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("max"))
	return rv
}
func (g MPSGraphRandomOpDescriptor) SetMax(value float32) {
	objc.Send[struct{}](g.ID, objc.Sel("setMax:"), value)
}

// The upper range of the distribution.
//
// # Discussion
//
// This value is used for Uniform with integer data types Defaults to
// INT32_MAX for uniform distributions and 0 for normal distributions.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/maxInteger
func (g MPSGraphRandomOpDescriptor) MaxInteger() int {
	rv := objc.Send[int](g.ID, objc.Sel("maxInteger"))
	return rv
}
func (g MPSGraphRandomOpDescriptor) SetMaxInteger(value int) {
	objc.Send[struct{}](g.ID, objc.Sel("setMaxInteger:"), value)
}

// The mean of the distribution.
//
// # Discussion
//
// This value is used for Normal and Truncated Normal disributions. Defaults
// to 0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/mean
func (g MPSGraphRandomOpDescriptor) Mean() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("mean"))
	return rv
}
func (g MPSGraphRandomOpDescriptor) SetMean(value float32) {
	objc.Send[struct{}](g.ID, objc.Sel("setMean:"), value)
}

// The lower range of the distribution.
//
// # Discussion
//
// This value is used for Uniform distributions with float data types and
// Truncated Normal disributions. Defaults to 0 for uniform distributions and
// -2 for normal distributions.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/min
func (g MPSGraphRandomOpDescriptor) Min() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("min"))
	return rv
}
func (g MPSGraphRandomOpDescriptor) SetMin(value float32) {
	objc.Send[struct{}](g.ID, objc.Sel("setMin:"), value)
}

// The lower range of the distribution.
//
// # Discussion
//
// This value is used for Uniform with integer data types Defaults to 0.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/minInteger
func (g MPSGraphRandomOpDescriptor) MinInteger() int {
	rv := objc.Send[int](g.ID, objc.Sel("minInteger"))
	return rv
}
func (g MPSGraphRandomOpDescriptor) SetMinInteger(value int) {
	objc.Send[struct{}](g.ID, objc.Sel("setMinInteger:"), value)
}

// The sampling method of the distribution.
//
// # Discussion
//
// This value is used for Normal and Truncated Normal disributions. See
// MPSGraphRandomNormalSamplingMethod. Defaults to
// MPSGraphRandomNormalSamplingInvCDF.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/samplingMethod
func (g MPSGraphRandomOpDescriptor) SamplingMethod() MPSGraphRandomNormalSamplingMethod {
	rv := objc.Send[MPSGraphRandomNormalSamplingMethod](g.ID, objc.Sel("samplingMethod"))
	return MPSGraphRandomNormalSamplingMethod(rv)
}
func (g MPSGraphRandomOpDescriptor) SetSamplingMethod(value MPSGraphRandomNormalSamplingMethod) {
	objc.Send[struct{}](g.ID, objc.Sel("setSamplingMethod:"), value)
}

// The standard deviation of the distribution.
//
// # Discussion
//
// This value is used for Normal and Truncated Normal disributions. For
// Truncated Normal distribution this defines the standard deviation parameter
// of the underlying Normal distribution, that is the width of the Gaussian,
// not the true standard deviation of the truncated distribution which
// typically differs from the standard deviation of the original Normal
// distribution. Defaults to 0 for uniform distributions and 1 for normal
// distributions.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/standardDeviation
func (g MPSGraphRandomOpDescriptor) StandardDeviation() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("standardDeviation"))
	return rv
}
func (g MPSGraphRandomOpDescriptor) SetStandardDeviation(value float32) {
	objc.Send[struct{}](g.ID, objc.Sel("setStandardDeviation:"), value)
}

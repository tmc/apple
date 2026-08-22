// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSFColorConversion] class.
var (
	_MPSFColorConversionClass     MPSFColorConversionClass
	_MPSFColorConversionClassOnce sync.Once
)

func getMPSFColorConversionClass() MPSFColorConversionClass {
	_MPSFColorConversionClassOnce.Do(func() {
		_MPSFColorConversionClass = MPSFColorConversionClass{class: objc.GetClass("MPSFColorConversion")}
	})
	return _MPSFColorConversionClass
}

// GetMPSFColorConversionClass returns the class object for MPSFColorConversion.
func GetMPSFColorConversionClass() MPSFColorConversionClass {
	return getMPSFColorConversionClass()
}

type MPSFColorConversionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSFColorConversionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSFColorConversionClass) Alloc() MPSFColorConversion {
	rv := objc.Send[MPSFColorConversion](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSFColorConversion.InitWithDeviceConversionFunctionNameSourceRangeOptionsError]
//   - [MPSFColorConversion.InitWithDeviceStartColorSpaceEndColorSpaceFunctionNameSourceRangeOptionsError]
//
// # Instance Properties
//
//   - [MPSFColorConversion.InputColorChannels]
//   - [MPSFColorConversion.Options]
//   - [MPSFColorConversion.OutputColorChannels]
//
// # Instance Methods
//
//   - [MPSFColorConversion.DescriptorFor1DTexture1]
//   - [MPSFColorConversion.DescriptorFor3DTexture1]
//   - [MPSFColorConversion.DescriptorFor3DTexture2]
//   - [MPSFColorConversion.EffectiveRange]
//   - [MPSFColorConversion.Initialize1DTexture1]
//   - [MPSFColorConversion.Initialize3DTexture1]
//   - [MPSFColorConversion.Initialize3DTexture2]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFColorConversion
type MPSFColorConversion struct {
	MPSFunction
}

// MPSFColorConversionFromID constructs a [MPSFColorConversion] from an objc.ID.
func MPSFColorConversionFromID(id objc.ID) MPSFColorConversion {
	return MPSFColorConversion{MPSFunction: MPSFunctionFromID(id)}
}

// NOTE: MPSFColorConversion adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSFColorConversion] class.
//
// # Initializers
//
//   - [IMPSFColorConversion.InitWithDeviceConversionFunctionNameSourceRangeOptionsError]
//   - [IMPSFColorConversion.InitWithDeviceStartColorSpaceEndColorSpaceFunctionNameSourceRangeOptionsError]
//
// # Instance Properties
//
//   - [IMPSFColorConversion.InputColorChannels]
//   - [IMPSFColorConversion.Options]
//   - [IMPSFColorConversion.OutputColorChannels]
//
// # Instance Methods
//
//   - [IMPSFColorConversion.DescriptorFor1DTexture1]
//   - [IMPSFColorConversion.DescriptorFor3DTexture1]
//   - [IMPSFColorConversion.DescriptorFor3DTexture2]
//   - [IMPSFColorConversion.EffectiveRange]
//   - [IMPSFColorConversion.Initialize1DTexture1]
//   - [IMPSFColorConversion.Initialize3DTexture1]
//   - [IMPSFColorConversion.Initialize3DTexture2]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFColorConversion
type IMPSFColorConversion interface {
	IMPSFunction

	// Topic: Initializers

	InitWithDeviceConversionFunctionNameSourceRangeOptionsError(device metal.MTLDevice, conversion coregraphics.CGColorConversionInfoRef, name string, sourceRange *MPSFunctions_AABB, options MPSFColorConversionOptions) (MPSFColorConversion, error)
	InitWithDeviceStartColorSpaceEndColorSpaceFunctionNameSourceRangeOptionsError(device metal.MTLDevice, start coregraphics.CGColorSpaceRef, end coregraphics.CGColorSpaceRef, name string, sourceRange *MPSFunctions_AABB, options MPSFColorConversionOptions) (MPSFColorConversion, error)

	// Topic: Instance Properties

	InputColorChannels() uint
	Options() MPSFColorConversionOptions
	OutputColorChannels() uint

	// Topic: Instance Methods

	DescriptorFor1DTexture1() metal.MTLTextureDescriptor
	DescriptorFor3DTexture1() metal.MTLTextureDescriptor
	DescriptorFor3DTexture2() metal.MTLTextureDescriptor
	EffectiveRange(inputRange MPSFunctions_AABB) MPSFunctions_AABB
	Initialize1DTexture1(tex metal.MTLTexture) MPSFColorConversion
	Initialize3DTexture1(tex metal.MTLTexture) MPSFColorConversion
	Initialize3DTexture2(tex metal.MTLTexture) MPSFColorConversion
}

// Init initializes the instance.
func (f MPSFColorConversion) Init() MPSFColorConversion {
	rv := objc.Send[MPSFColorConversion](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MPSFColorConversion) Autorelease() MPSFColorConversion {
	rv := objc.Send[MPSFColorConversion](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSFColorConversion creates a new MPSFColorConversion instance.
func NewMPSFColorConversion() MPSFColorConversion {
	class := getMPSFColorConversionClass()
	rv := objc.Send[MPSFColorConversion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// # Discussion
//
// # Called by NSCoder to decode MPSKernels
//
// This standard method doesn’t allow for control over which device the
// object targets. By default this will be the Metal system default device. If
// you want another device, use the MPSKeyedUnarchiver or other to decode the
// function.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFunction/init(coder:)
func NewFColorConversionWithCoder(aDecoder foundation.INSCoder) MPSFColorConversion {
	instance := getMPSFColorConversionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSFColorConversionFromID(rv)
}

// device: A valid MTLDevice where the conversion will be used
//
// conversion: A CGColorConversionInfoRef to represent the conversion. If NULL, a
// conversion function that returns its argument will be returned.
//
// name: The name of the Metal Shading Language function to build.
//
// sourceRange: If not NULL, the range limit guarantees that the input texels to the
// MTLFunction will not appear outside the given axis aligned bounding box.
// This, in combination with precision limits (see options), may allow for a
// faster conversion calculation. If a rangeLimit is provided, the result of
// the conversion involving out of range inputs is undefined.
//
// options: Options to use when building the conversion CAUTION: when conversion is
// NULL, MPSFunctions has no information about the number of channels in the
// result texel, and so can not intelligently handle
// MPSFColorConversionOptionsReturnGrayscaleAsRGB. In this case, it will
// assume the output content is grayscale and remap it to {Y,Y,Y,A} as
// requested. Your application should either intelligently set the option only
// for grayscale content, or call the other -init method that consumes two
// colorspaces which can manage this detail itself.
//
// # Return Value
//
// On success, a valid MPSFunctionsConversion object. If the conversion
// can’t be done, for example because it consumes or produces more than four
// channels, nil will be returned, and an appropriate error code created.
//
// # Discussion
//
// # Initialize a new MPSFunctionsConversion object
//
// Reads the CGColorConversionInfoRef and creates an internal representation
// Kicks off an asynchronous compilation task to build a MTLFunction
// appropriate for the device. Calling the .function or .error properties will
// stop and wait for it. Since the compilation task may take a few
// milliseconds, your application should create the MPSFunctionsConversion
// object as soon as it knows the conversion will be needed.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFColorConversion/init(device:conversion:functionName:sourceRange:options:)
func NewFColorConversionWithDeviceConversionFunctionNameSourceRangeOptionsError(device metal.MTLDevice, conversion coregraphics.CGColorConversionInfoRef, name string, sourceRange *MPSFunctions_AABB, options MPSFColorConversionOptions) (MPSFColorConversion, error) {
	var errorPtr objc.ID
	instance := getMPSFColorConversionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:conversion:functionName:sourceRange:options:error:"), device, conversion, objc.String(name), unsafe.Pointer(sourceRange), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MPSFColorConversion{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MPSFColorConversion{}, objc.ErrInitFailed
	}
	return MPSFColorConversionFromID(rv), nil
}

// device: A valid MTLDevice where the conversion will be used
//
// start: The CGColorSpaceRef for the input data to the conversion
//
// end: The CGColorSpaceRef for the output data from the conversion
//
// name: The name of the Metal Shading Language function to build.
//
// sourceRange: If not NULL, the range limit guarantees that the input texels to the
// MTLFunction will not appear outside the given axis aligned bounding box.
// This, in combination with precision limits (see options), may allow for a
// faster conversion calculation. If a rangeLimit is provided, the result of
// the conversion involving out of range inputs is undefined.
//
// options: Options to use when building the conversion
//
// # Return Value
//
// On success, a valid MPSFunctionsConversion object. If the conversion
// can’t be done, for example because it consumes or produces more than four
// channels, nil will be returned, and an appropriate error code created.
//
// # Discussion
//
// # Initialize a new MPSFunctionsConversion object
//
// Builds a MPSFunctionsConverison object from a starting and ending
// colorspace. Kicks off an asynchronous compilation task to build a
// MTLFunction appropriate for the device. Calling the .function or .error
// propertywill stop and wait for it. Since the compilation task may take a
// few milliseconds, your application should create the MPSFunctionsConversion
// object as soon as it knows the conversion will be needed.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFColorConversion/init(device:start:end:functionName:sourceRange:options:)
func NewFColorConversionWithDeviceStartColorSpaceEndColorSpaceFunctionNameSourceRangeOptionsError(device metal.MTLDevice, start coregraphics.CGColorSpaceRef, end coregraphics.CGColorSpaceRef, name string, sourceRange *MPSFunctions_AABB, options MPSFColorConversionOptions) (MPSFColorConversion, error) {
	var errorPtr objc.ID
	instance := getMPSFColorConversionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:startColorSpace:endColorSpace:functionName:sourceRange:options:error:"), device, start, end, objc.String(name), unsafe.Pointer(sourceRange), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MPSFColorConversion{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MPSFColorConversion{}, objc.ErrInitFailed
	}
	return MPSFColorConversionFromID(rv), nil
}

// device: A valid MTLDevice where the conversion will be used
//
// conversion: A CGColorConversionInfoRef to represent the conversion. If NULL, a
// conversion function that returns its argument will be returned.
//
// name: The name of the Metal Shading Language function to build.
//
// sourceRange: If not NULL, the range limit guarantees that the input texels to the
// MTLFunction will not appear outside the given axis aligned bounding box.
// This, in combination with precision limits (see options), may allow for a
// faster conversion calculation. If a rangeLimit is provided, the result of
// the conversion involving out of range inputs is undefined.
//
// options: Options to use when building the conversion CAUTION: when conversion is
// NULL, MPSFunctions has no information about the number of channels in the
// result texel, and so can not intelligently handle
// MPSFColorConversionOptionsReturnGrayscaleAsRGB. In this case, it will
// assume the output content is grayscale and remap it to {Y,Y,Y,A} as
// requested. Your application should either intelligently set the option only
// for grayscale content, or call the other -init method that consumes two
// colorspaces which can manage this detail itself.
//
// # Return Value
//
// On success, a valid MPSFunctionsConversion object. If the conversion
// can’t be done, for example because it consumes or produces more than four
// channels, nil will be returned, and an appropriate error code created.
//
// # Discussion
//
// # Initialize a new MPSFunctionsConversion object
//
// Reads the CGColorConversionInfoRef and creates an internal representation
// Kicks off an asynchronous compilation task to build a MTLFunction
// appropriate for the device. Calling the .function or .error properties will
// stop and wait for it. Since the compilation task may take a few
// milliseconds, your application should create the MPSFunctionsConversion
// object as soon as it knows the conversion will be needed.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFColorConversion/init(device:conversion:functionName:sourceRange:options:)
func (f MPSFColorConversion) InitWithDeviceConversionFunctionNameSourceRangeOptionsError(device metal.MTLDevice, conversion coregraphics.CGColorConversionInfoRef, name string, sourceRange *MPSFunctions_AABB, options MPSFColorConversionOptions) (MPSFColorConversion, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("initWithDevice:conversion:functionName:sourceRange:options:error:"), device, conversion, objc.String(name), unsafe.Pointer(sourceRange), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MPSFColorConversion{}, foundation.NSErrorFrom(errorPtr)
	}
	return MPSFColorConversionFromID(rv), nil

}

// device: A valid MTLDevice where the conversion will be used
//
// start: The CGColorSpaceRef for the input data to the conversion
//
// end: The CGColorSpaceRef for the output data from the conversion
//
// name: The name of the Metal Shading Language function to build.
//
// sourceRange: If not NULL, the range limit guarantees that the input texels to the
// MTLFunction will not appear outside the given axis aligned bounding box.
// This, in combination with precision limits (see options), may allow for a
// faster conversion calculation. If a rangeLimit is provided, the result of
// the conversion involving out of range inputs is undefined.
//
// options: Options to use when building the conversion
//
// # Return Value
//
// On success, a valid MPSFunctionsConversion object. If the conversion
// can’t be done, for example because it consumes or produces more than four
// channels, nil will be returned, and an appropriate error code created.
//
// # Discussion
//
// # Initialize a new MPSFunctionsConversion object
//
// Builds a MPSFunctionsConverison object from a starting and ending
// colorspace. Kicks off an asynchronous compilation task to build a
// MTLFunction appropriate for the device. Calling the .function or .error
// propertywill stop and wait for it. Since the compilation task may take a
// few milliseconds, your application should create the MPSFunctionsConversion
// object as soon as it knows the conversion will be needed.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFColorConversion/init(device:start:end:functionName:sourceRange:options:)
func (f MPSFColorConversion) InitWithDeviceStartColorSpaceEndColorSpaceFunctionNameSourceRangeOptionsError(device metal.MTLDevice, start coregraphics.CGColorSpaceRef, end coregraphics.CGColorSpaceRef, name string, sourceRange *MPSFunctions_AABB, options MPSFColorConversionOptions) (MPSFColorConversion, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("initWithDevice:startColorSpace:endColorSpace:functionName:sourceRange:options:error:"), device, start, end, objc.String(name), unsafe.Pointer(sourceRange), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MPSFColorConversion{}, foundation.NSErrorFrom(errorPtr)
	}
	return MPSFColorConversionFromID(rv), nil

}

// # Discussion
//
// A descriptor for mhdr_conversion_data_t.tex1d_1
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFColorConversion/descriptorFor1DTexture1()
func (f MPSFColorConversion) DescriptorFor1DTexture1() metal.MTLTextureDescriptor {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("descriptorFor1DTexture1"))
	return metal.MTLTextureDescriptorFromID(rv)
}

// # Discussion
//
// A descriptor for mhdr_conversion_data_t.tex3d_1
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFColorConversion/descriptorFor3DTexture1()
func (f MPSFColorConversion) DescriptorFor3DTexture1() metal.MTLTextureDescriptor {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("descriptorFor3DTexture1"))
	return metal.MTLTextureDescriptorFromID(rv)
}

// # Discussion
//
// A descriptor for mhdr_conversion_data_t.tex3d_2
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFColorConversion/descriptorFor3DTexture2()
func (f MPSFColorConversion) DescriptorFor3DTexture2() metal.MTLTextureDescriptor {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("descriptorFor3DTexture2"))
	return metal.MTLTextureDescriptorFromID(rv)
}

// # Discussion
//
// # Estimate the gamut produced by the function based on a range of inputs
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFColorConversion/effectiveRange(_:)
func (f MPSFColorConversion) EffectiveRange(inputRange MPSFunctions_AABB) MPSFunctions_AABB {
	rv := objc.Send[MPSFunctions_AABB](f.ID, objc.Sel("effectiveRange:"), inputRange)
	return MPSFunctions_AABB(rv)
}

// # Discussion
//
// Overwrite tex1d_1 will the LUT data needed for the conversion
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFColorConversion/initialize1DTexture1(_:)
func (f MPSFColorConversion) Initialize1DTexture1(tex metal.MTLTexture) MPSFColorConversion {
	rv := objc.Send[MPSFColorConversion](f.ID, objc.Sel("initialize1DTexture1:"), tex)
	return rv
}

// # Discussion
//
// Overwrite tex3d_1 will the LUT data needed for the conversion
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFColorConversion/initialize3DTexture1(_:)
func (f MPSFColorConversion) Initialize3DTexture1(tex metal.MTLTexture) MPSFColorConversion {
	rv := objc.Send[MPSFColorConversion](f.ID, objc.Sel("initialize3DTexture1:"), tex)
	return rv
}

// # Discussion
//
// Overwrite tex3d_2 will the LUT data needed for the conversion
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFColorConversion/initialize3DTexture2(_:)
func (f MPSFColorConversion) Initialize3DTexture2(tex metal.MTLTexture) MPSFColorConversion {
	rv := objc.Send[MPSFColorConversion](f.ID, objc.Sel("initialize3DTexture2:"), tex)
	return rv
}

// # Discussion
//
// # The number of color channels used by the conversion in the float4 texel
//
// When the conversion is initialized with a NULL CGColorConversionInfoRef
// this value will be 0
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFColorConversion/inputColorChannels
func (f MPSFColorConversion) InputColorChannels() uint {
	rv := objc.Send[uint](f.ID, objc.Sel("inputColorChannels"))
	return rv
}

// # Discussion
//
// # The options used when creating the object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFColorConversion/options
func (f MPSFColorConversion) Options() MPSFColorConversionOptions {
	rv := objc.Send[MPSFColorConversionOptions](f.ID, objc.Sel("options"))
	return MPSFColorConversionOptions(rv)
}

// # Discussion
//
// # The number of color channels produced by the conversion in the float4 texel
//
// When the conversion is initialized with a NULL CGColorConversionInfoRef
// this value will be 0
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFColorConversion/outputColorChannels
func (f MPSFColorConversion) OutputColorChannels() uint {
	rv := objc.Send[uint](f.ID, objc.Sel("outputColorChannels"))
	return rv
}

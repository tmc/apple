// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSFunction] class.
var (
	_MPSFunctionClass     MPSFunctionClass
	_MPSFunctionClassOnce sync.Once
)

func getMPSFunctionClass() MPSFunctionClass {
	_MPSFunctionClassOnce.Do(func() {
		_MPSFunctionClass = MPSFunctionClass{class: objc.GetClass("MPSFunction")}
	})
	return _MPSFunctionClass
}

// GetMPSFunctionClass returns the class object for MPSFunction.
func GetMPSFunctionClass() MPSFunctionClass {
	return getMPSFunctionClass()
}

type MPSFunctionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSFunctionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSFunctionClass) Alloc() MPSFunction {
	rv := objc.Send[MPSFunction](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSFunction.InitWithCoder]
//
// # Instance Properties
//
//   - [MPSFunction.Device]
//   - [MPSFunction.Error]: The error produced when attempting to build the function
//   - [MPSFunction.Function]: A MTLFunction that you can link into your shader
//   - [MPSFunction.Name]
//
// # Instance Methods
//
//   - [MPSFunction.CopyWithZoneDevice]
//   - [MPSFunction.FunctionPrototype]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFunction
type MPSFunction struct {
	objectivec.Object
}

// MPSFunctionFromID constructs a [MPSFunction] from an objc.ID.
func MPSFunctionFromID(id objc.ID) MPSFunction {
	return MPSFunction{objectivec.Object{ID: id}}
}

// NOTE: MPSFunction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSFunction] class.
//
// # Initializers
//
//   - [IMPSFunction.InitWithCoder]
//
// # Instance Properties
//
//   - [IMPSFunction.Device]
//   - [IMPSFunction.Error]: The error produced when attempting to build the function
//   - [IMPSFunction.Function]: A MTLFunction that you can link into your shader
//   - [IMPSFunction.Name]
//
// # Instance Methods
//
//   - [IMPSFunction.CopyWithZoneDevice]
//   - [IMPSFunction.FunctionPrototype]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFunction
type IMPSFunction interface {
	objectivec.IObject

	// Topic: Initializers

	InitWithCoder(aDecoder foundation.INSCoder) MPSFunction

	// Topic: Instance Properties

	Device() metal.MTLDevice
	// The error produced when attempting to build the function
	Error() foundation.NSError
	// A MTLFunction that you can link into your shader
	Function() metal.MTLFunction
	Name() string

	// Topic: Instance Methods

	CopyWithZoneDevice(zone unsafe.Pointer, device metal.MTLDevice) IMPSFunction
	FunctionPrototype() string

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f MPSFunction) Init() MPSFunction {
	rv := objc.Send[MPSFunction](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MPSFunction) Autorelease() MPSFunction {
	rv := objc.Send[MPSFunction](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSFunction creates a new MPSFunction instance.
func NewMPSFunction() MPSFunction {
	class := getMPSFunctionClass()
	rv := objc.Send[MPSFunction](objc.ID(class.class), objc.Sel("new"))
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
func NewFunctionWithCoder(aDecoder foundation.INSCoder) MPSFunction {
	instance := getMPSFunctionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSFunctionFromID(rv)
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
func (f MPSFunction) InitWithCoder(aDecoder foundation.INSCoder) MPSFunction {
	rv := objc.Send[MPSFunction](f.ID, objc.Sel("initWithCoder:"), aDecoder)
	return rv
}

// zone: The NSZone in which to allocate the object
//
// device: The device for the new MPSKernel. If nil, then use self.device.
//
// # Return Value
//
// A pointer to a copy of this MPSKernel. This will fail, returning nil if the
// device is not supported. Devices must be MTLFeatureSet_iOS_GPUFamily2_v1 or
// later.
//
// # Discussion
//
// # Make a copy of this MPSFunction for a new device
//
// -copyWithZone: will call this API to make a copy of the MPSKernel on the
// same device. This interface may also be called directly to make a copy of
// the MPSFunction on a new device.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFunction/copy(with:device:)
func (f MPSFunction) CopyWithZoneDevice(zone unsafe.Pointer, device metal.MTLDevice) IMPSFunction {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return MPSFunctionFromID(rv)
}

// # Discussion
//
// # Get a source level representation of the function prototype
//
// If your application is building its shaders from source at run time, this
// string will declare the appropriate function prototypes for the conversion
// routine appropriate to the version of MetalHDR you are currently running.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFunction/functionPrototype()
func (f MPSFunction) FunctionPrototype() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("functionPrototype"))
	return foundation.NSStringFromID(rv).String()
}
func (f MPSFunction) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// # Discussion
//
// # Part of the NSSecureCoding
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFunction/supportsSecureCoding()
func (_MPSFunctionClass MPSFunctionClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MPSFunctionClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// # Discussion
//
// # The device where the Metal Shading Language function will run
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFunction/device
func (f MPSFunction) Device() metal.MTLDevice {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("device"))
	return metal.MTLDeviceObjectFromID(rv)
}

// The error produced when attempting to build the function
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFunction/error
func (f MPSFunction) Error() foundation.NSError {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}

// A MTLFunction that you can link into your shader
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFunction/function
func (f MPSFunction) Function() metal.MTLFunction {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("function"))
	return metal.MTLFunctionObjectFromID(rv)
}

// # Discussion
//
// # The name of the Metal Shading Language function built by this object
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSFunction/name
func (f MPSFunction) Name() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

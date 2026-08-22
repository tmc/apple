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

// The class instance for the [MPSKernel] class.
var (
	_MPSKernelClass     MPSKernelClass
	_MPSKernelClassOnce sync.Once
)

func getMPSKernelClass() MPSKernelClass {
	_MPSKernelClassOnce.Do(func() {
		_MPSKernelClass = MPSKernelClass{class: objc.GetClass("MPSKernel")}
	})
	return _MPSKernelClass
}

// GetMPSKernelClass returns the class object for MPSKernel.
func GetMPSKernelClass() MPSKernelClass {
	return getMPSKernelClass()
}

type MPSKernelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSKernelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSKernelClass) Alloc() MPSKernel {
	rv := objc.Send[MPSKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A standard interface for Metal Performance Shaders kernels.
//
// # Overview
//
// You should not use the [MPSKernel] class directly. Instead, a number of
// subclasses are available that define specific high-performance
// data-parallel operations.
//
// The basic sequence for applying a kernel to an image is as follows:
//
// - Initialize a kernel corresponding to the operation you wish to perform:
//
// - Encode the kernel into a command buffer.
//
// Encoding the kernel merely encodes the operation into a command buffer. It
// does not modify any pixels, yet. All kernel state has been copied to the
// command buffer. Kernels may be reused. If the texture was previously
// operated on by another command encoder (e.g. a render command encoder), you
// should call the [endEncoding()] method on the other encoder before encoding
// the filter.
//
// Some kernels work in place, even in situations where Metal might not
// normally allow in-place operation on textures. If in-place operation is
// desired, you may attempt to call the
// [MPSUnaryImageKernel.EncodeToCommandBufferInPlaceTextureFallbackCopyAllocator]
// method. If the operation cannot be completed in place, then false will be
// returned and you will have to create a new result texture and try again. To
// make an in-place image filter reliable, pass a fallback [MPSCopyAllocator]
// block to the method to create a new texture to write to in the event that a
// filter cannot operate in place.
//
// You may repeat step 2 to encode more kernels, as desired. 3. After encoding
// any additional work to the command buffer using other encoders, submit the
// command buffer to your command queue, using:
//
// # Initializers
//
//   - [MPSKernel.InitWithCoder]
//   - [MPSKernel.InitWithCoderDevice]
//
// # Methods
//
//   - [MPSKernel.InitWithDevice]: Initializes a new kernel object.
//   - [MPSKernel.CopyWithZoneDevice]: Makes a copy of this kernel object for a new device.
//
// # Properties
//
//   - [MPSKernel.Options]: The set of options used to run the kernel.
//   - [MPSKernel.SetOptions]
//   - [MPSKernel.Device]: The device on which the kernel will be used.
//   - [MPSKernel.Label]: The string that identifies the kernel.
//   - [MPSKernel.SetLabel]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel
//
// [endEncoding()]: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/endEncoding()
type MPSKernel struct {
	objectivec.Object
}

// MPSKernelFromID constructs a [MPSKernel] from an objc.ID.
//
// A standard interface for Metal Performance Shaders kernels.
func MPSKernelFromID(id objc.ID) MPSKernel {
	return MPSKernel{objectivec.Object{ID: id}}
}

// NOTE: MPSKernel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSKernel] class.
//
// # Initializers
//
//   - [IMPSKernel.InitWithCoder]
//   - [IMPSKernel.InitWithCoderDevice]
//
// # Methods
//
//   - [IMPSKernel.InitWithDevice]: Initializes a new kernel object.
//   - [IMPSKernel.CopyWithZoneDevice]: Makes a copy of this kernel object for a new device.
//
// # Properties
//
//   - [IMPSKernel.Options]: The set of options used to run the kernel.
//   - [IMPSKernel.SetOptions]
//   - [IMPSKernel.Device]: The device on which the kernel will be used.
//   - [IMPSKernel.Label]: The string that identifies the kernel.
//   - [IMPSKernel.SetLabel]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel
type IMPSKernel interface {
	objectivec.IObject

	// Topic: Initializers

	InitWithCoder(aDecoder foundation.INSCoder) MPSKernel
	InitWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSKernel

	// Topic: Methods

	// Initializes a new kernel object.
	InitWithDevice(device metal.MTLDevice) MPSKernel
	// Makes a copy of this kernel object for a new device.
	CopyWithZoneDevice(zone unsafe.Pointer, device metal.MTLDevice) IMPSKernel

	// Topic: Properties

	// The set of options used to run the kernel.
	Options() MPSKernelOptions
	SetOptions(value MPSKernelOptions)
	// The device on which the kernel will be used.
	Device() metal.MTLDevice
	// The string that identifies the kernel.
	Label() string
	SetLabel(value string)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (k MPSKernel) Init() MPSKernel {
	rv := objc.Send[MPSKernel](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k MPSKernel) Autorelease() MPSKernel {
	rv := objc.Send[MPSKernel](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSKernel creates a new MPSKernel instance.
func NewMPSKernel() MPSKernel {
	class := getMPSKernelClass()
	rv := objc.Send[MPSKernel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewKernelWithCoder(aDecoder foundation.INSCoder) MPSKernel {
	instance := getMPSKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:device:)
func NewKernelWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSKernel {
	instance := getMPSKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSKernelFromID(rv)
}

// Initializes a new kernel object.
//
// device: The Metal device on which the kernel will be used.
//
// # Return Value
//
// An initialized kernel object.
//
// # Discussion
//
// This method fails if the device is not supported. Query the
// [MPSSupportsMTLDevice] function to determine whether the device is
// supported.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(device:)
func NewKernelWithDevice(device metal.MTLDevice) MPSKernel {
	instance := getMPSKernelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSKernelFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func (k MPSKernel) InitWithCoder(aDecoder foundation.INSCoder) MPSKernel {
	rv := objc.Send[MPSKernel](k.ID, objc.Sel("initWithCoder:"), aDecoder)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:device:)
func (k MPSKernel) InitWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSKernel {
	rv := objc.Send[MPSKernel](k.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return rv
}

// Initializes a new kernel object.
//
// device: The Metal device on which the kernel will be used.
//
// # Return Value
//
// An initialized kernel object.
//
// # Discussion
//
// This method fails if the device is not supported. Query the
// [MPSSupportsMTLDevice] function to determine whether the device is
// supported.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(device:)
func (k MPSKernel) InitWithDevice(device metal.MTLDevice) MPSKernel {
	rv := objc.Send[MPSKernel](k.ID, objc.Sel("initWithDevice:"), device)
	return rv
}

// Makes a copy of this kernel object for a new device.
//
// zone: The zone in which to allocate the kernel object.
//
// device: The Metal device for the new kernel object.
//
// # Return Value
//
// A copy of this kernel object.
//
// # Discussion
//
// The same kernel objects should not be used to encode separate kernel
// operations on multiple command buffers from multiple threads. Many kernels
// have mutable properties that might be changed by another thread while the
// kernel is being encoded. If you need to use a kernel from multiple threads,
// make a copy of it for each additional thread using [copy(with:)] or
// [MPSKernel.CopyWithZoneDevice]. Note that the [copy(with:)] method makes a
// copy of the kernel object on the same device.
//
// This method fails if the device is not supported. Query the
// [MPSSupportsMTLDevice] function to determine whether the device is
// supported.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/copy(with:device:)
//
// [copy(with:)]: https://developer.apple.com/documentation/Foundation/NSCopying/copy(with:)
func (k MPSKernel) CopyWithZoneDevice(zone unsafe.Pointer, device metal.MTLDevice) IMPSKernel {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return MPSKernelFromID(rv)
}
func (k MPSKernel) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](k.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The set of options used to run the kernel.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/options
func (k MPSKernel) Options() MPSKernelOptions {
	rv := objc.Send[MPSKernelOptions](k.ID, objc.Sel("options"))
	return MPSKernelOptions(rv)
}
func (k MPSKernel) SetOptions(value MPSKernelOptions) {
	objc.Send[struct{}](k.ID, objc.Sel("setOptions:"), value)
}

// The device on which the kernel will be used.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/device
func (k MPSKernel) Device() metal.MTLDevice {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("device"))
	return metal.MTLDeviceObjectFromID(rv)
}

// The string that identifies the kernel.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/label
func (k MPSKernel) Label() string {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (k MPSKernel) SetLabel(value string) {
	objc.Send[struct{}](k.ID, objc.Sel("setLabel:"), objc.String(value))
}

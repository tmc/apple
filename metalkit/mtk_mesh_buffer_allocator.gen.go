// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTKMeshBufferAllocator] class.
var (
	_MTKMeshBufferAllocatorClass     MTKMeshBufferAllocatorClass
	_MTKMeshBufferAllocatorClassOnce sync.Once
)

func getMTKMeshBufferAllocatorClass() MTKMeshBufferAllocatorClass {
	_MTKMeshBufferAllocatorClassOnce.Do(func() {
		_MTKMeshBufferAllocatorClass = MTKMeshBufferAllocatorClass{class: objc.GetClass("MTKMeshBufferAllocator")}
	})
	return _MTKMeshBufferAllocatorClass
}

// GetMTKMeshBufferAllocatorClass returns the class object for MTKMeshBufferAllocator.
func GetMTKMeshBufferAllocatorClass() MTKMeshBufferAllocatorClass {
	return getMTKMeshBufferAllocatorClass()
}

type MTKMeshBufferAllocatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTKMeshBufferAllocatorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTKMeshBufferAllocatorClass) Alloc() MTKMeshBufferAllocator {
	rv := objc.Send[MTKMeshBufferAllocator](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An interface for allocating a MetalKit buffer that backs the vertex data of
// a Model I/O mesh, suitable for use in a Metal app.
//
// # Initialization
//
//   - [MTKMeshBufferAllocator.InitWithDevice]: Initializes a new allocator object.
//
// # Device
//
//   - [MTKMeshBufferAllocator.Device]: The device used to create Metal objects.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMeshBufferAllocator
type MTKMeshBufferAllocator struct {
	objectivec.Object
}

// MTKMeshBufferAllocatorFromID constructs a [MTKMeshBufferAllocator] from an objc.ID.
//
// An interface for allocating a MetalKit buffer that backs the vertex data of
// a Model I/O mesh, suitable for use in a Metal app.
func MTKMeshBufferAllocatorFromID(id objc.ID) MTKMeshBufferAllocator {
	return MTKMeshBufferAllocator{objectivec.Object{ID: id}}
}
// NOTE: MTKMeshBufferAllocator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTKMeshBufferAllocator] class.
//
// # Initialization
//
//   - [IMTKMeshBufferAllocator.InitWithDevice]: Initializes a new allocator object.
//
// # Device
//
//   - [IMTKMeshBufferAllocator.Device]: The device used to create Metal objects.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMeshBufferAllocator
type IMTKMeshBufferAllocator interface {
	objectivec.IObject

	// Topic: Initialization

	// Initializes a new allocator object.
	InitWithDevice(device metal.MTLDevice) MTKMeshBufferAllocator

	// Topic: Device

	// The device used to create Metal objects.
	Device() metal.MTLDevice
}

// Init initializes the instance.
func (m MTKMeshBufferAllocator) Init() MTKMeshBufferAllocator {
	rv := objc.Send[MTKMeshBufferAllocator](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTKMeshBufferAllocator) Autorelease() MTKMeshBufferAllocator {
	rv := objc.Send[MTKMeshBufferAllocator](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTKMeshBufferAllocator creates a new MTKMeshBufferAllocator instance.
func NewMTKMeshBufferAllocator() MTKMeshBufferAllocator {
	class := getMTKMeshBufferAllocatorClass()
	rv := objc.Send[MTKMeshBufferAllocator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a new allocator object.
//
// device: The Metal device on which to create buffers.
//
// # Return Value
// 
// An initialized allocator object.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMeshBufferAllocator/init(device:)
func NewMeshBufferAllocatorWithDevice(device metal.MTLDevice) MTKMeshBufferAllocator {
	instance := getMTKMeshBufferAllocatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MTKMeshBufferAllocatorFromID(rv)
}

// Initializes a new allocator object.
//
// device: The Metal device on which to create buffers.
//
// # Return Value
// 
// An initialized allocator object.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMeshBufferAllocator/init(device:)
func (m MTKMeshBufferAllocator) InitWithDevice(device metal.MTLDevice) MTKMeshBufferAllocator {
	rv := objc.Send[MTKMeshBufferAllocator](m.ID, objc.Sel("initWithDevice:"), device)
	return rv
}

// The device used to create Metal objects.
//
// # Discussion
// 
// A Metal device must be specialized at initialization.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKMeshBufferAllocator/device
func (m MTKMeshBufferAllocator) Device() metal.MTLDevice {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("device"))
	return metal.MTLDeviceObjectFromID(rv)
}


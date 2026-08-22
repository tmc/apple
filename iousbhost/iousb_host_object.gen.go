// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOUSBHostObject] class.
var (
	_IOUSBHostObjectClass     IOUSBHostObjectClass
	_IOUSBHostObjectClassOnce sync.Once
)

func getIOUSBHostObjectClass() IOUSBHostObjectClass {
	_IOUSBHostObjectClassOnce.Do(func() {
		_IOUSBHostObjectClass = IOUSBHostObjectClass{class: objc.GetClass("IOUSBHostObject")}
	})
	return _IOUSBHostObjectClass
}

// GetIOUSBHostObjectClass returns the class object for IOUSBHostObject.
func GetIOUSBHostObjectClass() IOUSBHostObjectClass {
	return getIOUSBHostObjectClass()
}

type IOUSBHostObjectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOUSBHostObjectClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOUSBHostObjectClass) Alloc() IOUSBHostObject {
	rv := objc.Send[IOUSBHostObject](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// This class provides basic functionality for sending device requests and
// retrieving descriptors.
//
// # Managing the Object Life Cycle
//
//   - [IOUSBHostObject.IoService]: A reference to the kernel object.
//   - [IOUSBHostObject.Queue]: The queue for servicing input/output requests.
//   - [IOUSBHostObject.Destroy]: Removes underlying allocations and connections from the USB host object.
//
// # Creating I/O Buffers
//
//   - [IOUSBHostObject.IoDataWithCapacityError]: Allocates a buffer for input/output requests.
//
// # Getting Host Information
//
//   - [IOUSBHostObject.DeviceAddress]: The device’s bus address.
//
// # Instance Properties
//
//   - [IOUSBHostObject.CapabilityDescriptors]
//   - [IOUSBHostObject.DeviceDescriptor]
//
// # Instance Methods
//
//   - [IOUSBHostObject.ConfigurationDescriptorWithIndexError]
//   - [IOUSBHostObject.ConfigurationDescriptorWithConfigurationValueError]
//   - [IOUSBHostObject.DestroyWithOptions]
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject
type IOUSBHostObject struct {
	objectivec.Object
}

// IOUSBHostObjectFromID constructs a [IOUSBHostObject] from an objc.ID.
//
// This class provides basic functionality for sending device requests and
// retrieving descriptors.
func IOUSBHostObjectFromID(id objc.ID) IOUSBHostObject {
	return IOUSBHostObject{objectivec.Object{ID: id}}
}

// NOTE: IOUSBHostObject adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOUSBHostObject] class.
//
// # Managing the Object Life Cycle
//
//   - [IIOUSBHostObject.IoService]: A reference to the kernel object.
//   - [IIOUSBHostObject.Queue]: The queue for servicing input/output requests.
//   - [IIOUSBHostObject.Destroy]: Removes underlying allocations and connections from the USB host object.
//
// # Creating I/O Buffers
//
//   - [IIOUSBHostObject.IoDataWithCapacityError]: Allocates a buffer for input/output requests.
//
// # Getting Host Information
//
//   - [IIOUSBHostObject.DeviceAddress]: The device’s bus address.
//
// # Instance Properties
//
//   - [IIOUSBHostObject.CapabilityDescriptors]
//   - [IIOUSBHostObject.DeviceDescriptor]
//
// # Instance Methods
//
//   - [IIOUSBHostObject.ConfigurationDescriptorWithIndexError]
//   - [IIOUSBHostObject.ConfigurationDescriptorWithConfigurationValueError]
//   - [IIOUSBHostObject.DestroyWithOptions]
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject
type IIOUSBHostObject interface {
	objectivec.IObject

	// Topic: Managing the Object Life Cycle

	// A reference to the kernel object.
	IoService() uint32
	// The queue for servicing input/output requests.
	Queue() dispatch.Queue
	// Removes underlying allocations and connections from the USB host object.
	Destroy()

	// Topic: Creating I/O Buffers

	// Allocates a buffer for input/output requests.
	IoDataWithCapacityError(capacity uint) (foundation.NSMutableData, error)

	// Topic: Getting Host Information

	// The device’s bus address.
	DeviceAddress() uint

	// Topic: Instance Properties

	CapabilityDescriptors() objectivec.IObject
	DeviceDescriptor() objectivec.IObject

	// Topic: Instance Methods

	ConfigurationDescriptorWithIndexError(index uint) (objectivec.IObject, error)
	ConfigurationDescriptorWithConfigurationValueError(configurationValue uint) (objectivec.IObject, error)
	DestroyWithOptions(options IOUSBHostObjectDestroyOptions)
}

// Init initializes the instance.
func (u IOUSBHostObject) Init() IOUSBHostObject {
	rv := objc.Send[IOUSBHostObject](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u IOUSBHostObject) Autorelease() IOUSBHostObject {
	rv := objc.Send[IOUSBHostObject](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOUSBHostObject creates a new IOUSBHostObject instance.
func NewIOUSBHostObject() IOUSBHostObject {
	class := getIOUSBHostObjectClass()
	rv := objc.Send[IOUSBHostObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Removes underlying allocations and connections from the USB host object.
//
// # Discussion
//
// When you no longer need the [IOUSBHostObject], call
// [IOUSBHostObject.Destroy]. This method destroys the connection with the
// kernel object and deregisters interest on [io_service_t]. Calling
// [IOUSBHostObject.Destroy] multiple times has no effect.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/destroy()
//
// [io_service_t]: https://developer.apple.com/documentation/iokit/io_service_t
func (u IOUSBHostObject) Destroy() {
	objc.Send[objc.ID](u.ID, objc.Sel("destroy"))
}

// Allocates a buffer for input/output requests.
//
// capacity: The size, in bytes, of the buffer to allocate.
//
// # Return Value
//
// A pointer to an allocated buffer.
//
// # Discussion
//
// This method allocates and maps a kernel buffer that the underlying
// controller hardware has optimized. Using this method, the buffer doesn’t
// bounce to perform DMA operations.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/ioData(withCapacity:)
func (u IOUSBHostObject) IoDataWithCapacityError(capacity uint) (foundation.NSMutableData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](u.ID, objc.Sel("ioDataWithCapacity:error:"), capacity, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSMutableData{}, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSMutableDataFromID(rv), nil

}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/configurationDescriptor(with:)
func (u IOUSBHostObject) ConfigurationDescriptorWithIndexError(index uint) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](u.ID, objc.Sel("configurationDescriptorWithIndex:error:"), index, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/configurationDescriptor(withConfigurationValue:)
func (u IOUSBHostObject) ConfigurationDescriptorWithConfigurationValueError(configurationValue uint) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](u.ID, objc.Sel("configurationDescriptorWithConfigurationValue:error:"), configurationValue, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// # Discussion
//
// Removes underlying allocations of the IOUSBHostObject object along with
// user client
//
// Extends destroy to take an options to modify the destroy behavior.
// Currently only the IOUSBHostObjectDestroyOptionsDeviceSurrender is defined
// to support surrendering ownersip of the kernel service. To be used when
// accepting the kUSBHostMessageDeviceIsRequestingClose message.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/destroy(options:)
func (u IOUSBHostObject) DestroyWithOptions(options IOUSBHostObjectDestroyOptions) {
	objc.Send[objc.ID](u.ID, objc.Sel("destroyWithOptions:"), options)
}

// A reference to the kernel object.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/ioService
func (u IOUSBHostObject) IoService() uint32 {
	rv := objc.Send[uint32](u.ID, objc.Sel("ioService"))
	return rv
}

// The queue for servicing input/output requests.
//
// # Discussion
//
// Use this queue only for asynchronous input/output requests.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/queue
func (u IOUSBHostObject) Queue() dispatch.Queue {
	rv := objc.Send[uintptr](u.ID, objc.Sel("queue"))
	return dispatch.QueueFromHandle(rv)
}

// The device’s bus address.
//
// # Return Value
//
// The current bus address of the device.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/deviceAddress
func (u IOUSBHostObject) DeviceAddress() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("deviceAddress"))
	return rv
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/capabilityDescriptors
func (u IOUSBHostObject) CapabilityDescriptors() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("capabilityDescriptors"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/deviceDescriptor
func (u IOUSBHostObject) DeviceDescriptor() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("deviceDescriptor"))
	return objectivec.Object{ID: rv}
}

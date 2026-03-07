// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioConsolePortConfigurationArray] class.
var (
	_VZVirtioConsolePortConfigurationArrayClass     VZVirtioConsolePortConfigurationArrayClass
	_VZVirtioConsolePortConfigurationArrayClassOnce sync.Once
)

func getVZVirtioConsolePortConfigurationArrayClass() VZVirtioConsolePortConfigurationArrayClass {
	_VZVirtioConsolePortConfigurationArrayClassOnce.Do(func() {
		_VZVirtioConsolePortConfigurationArrayClass = VZVirtioConsolePortConfigurationArrayClass{class: objc.GetClass("VZVirtioConsolePortConfigurationArray")}
	})
	return _VZVirtioConsolePortConfigurationArrayClass
}

// GetVZVirtioConsolePortConfigurationArrayClass returns the class object for VZVirtioConsolePortConfigurationArray.
func GetVZVirtioConsolePortConfigurationArrayClass() VZVirtioConsolePortConfigurationArrayClass {
	return getVZVirtioConsolePortConfigurationArrayClass()
}

type VZVirtioConsolePortConfigurationArrayClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioConsolePortConfigurationArrayClass) Alloc() VZVirtioConsolePortConfigurationArray {
	rv := objc.Send[VZVirtioConsolePortConfigurationArray](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// A class that represents a collection of Virtio console port configurations.
//
// # Overview
// 
// This array stores a collection of port configurations for a
// [VZVirtioConsoleDeviceConfiguration]. The index in the array corresponds to
// the port index that the VM uses. You can set a [VZVirtioConsolePortConfigurationArray.MaximumPortCount] value,
// but the value must be larger than the highest indexed port. If there’s no
// `maximumPortCount` value set, the framework uses the value the highest
// indexed port.
//
// # Determining the number of ports
//
//   - [VZVirtioConsolePortConfigurationArray.MaximumPortCount]: An unsigned integer that represents the maximum number of ports allocated by this device.
//   - [VZVirtioConsolePortConfigurationArray.SetMaximumPortCount]
//
// # Accessing a specific port
//
//   - [VZVirtioConsolePortConfigurationArray.ObjectAtIndexedSubscript]: Returns the Virtio console port configuration as the specified index.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfigurationArray
type VZVirtioConsolePortConfigurationArray struct {
	objectivec.Object
}

// VZVirtioConsolePortConfigurationArrayFromID constructs a [VZVirtioConsolePortConfigurationArray] from an objc.ID.
//
// A class that represents a collection of Virtio console port configurations.
func VZVirtioConsolePortConfigurationArrayFromID(id objc.ID) VZVirtioConsolePortConfigurationArray {
	return VZVirtioConsolePortConfigurationArray{objectivec.Object{id}}
}
// NOTE: VZVirtioConsolePortConfigurationArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZVirtioConsolePortConfigurationArray] class.
//
// # Determining the number of ports
//
//   - [IVZVirtioConsolePortConfigurationArray.MaximumPortCount]: An unsigned integer that represents the maximum number of ports allocated by this device.
//   - [IVZVirtioConsolePortConfigurationArray.SetMaximumPortCount]
//
// # Accessing a specific port
//
//   - [IVZVirtioConsolePortConfigurationArray.ObjectAtIndexedSubscript]: Returns the Virtio console port configuration as the specified index.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfigurationArray
type IVZVirtioConsolePortConfigurationArray interface {
	objectivec.IObject

	// Topic: Determining the number of ports

	// An unsigned integer that represents the maximum number of ports allocated by this device.
	MaximumPortCount() uint32
	SetMaximumPortCount(value uint32)

	// Topic: Accessing a specific port

	// Returns the Virtio console port configuration as the specified index.
	ObjectAtIndexedSubscript(portIndex uint) IVZVirtioConsolePortConfiguration

	SetObjectAtIndexedSubscript(configuration IVZVirtioConsolePortConfiguration, portIndex uint)
}





// Init initializes the instance.
func (v VZVirtioConsolePortConfigurationArray) Init() VZVirtioConsolePortConfigurationArray {
	rv := objc.Send[VZVirtioConsolePortConfigurationArray](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioConsolePortConfigurationArray) Autorelease() VZVirtioConsolePortConfigurationArray {
	rv := objc.Send[VZVirtioConsolePortConfigurationArray](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsolePortConfigurationArray creates a new VZVirtioConsolePortConfigurationArray instance.
func NewVZVirtioConsolePortConfigurationArray() VZVirtioConsolePortConfigurationArray {
	class := getVZVirtioConsolePortConfigurationArrayClass()
	rv := objc.Send[VZVirtioConsolePortConfigurationArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Returns the Virtio console port configuration as the specified index.
//
// portIndex: The index of the configuration to retrieve.
//
// # Return Value
// 
// The [VZVirtioConsolePortConfiguration], or nil is the index exceeds the
// number of configurations in the array.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfigurationArray/subscript(_:)
func (v VZVirtioConsolePortConfigurationArray) ObjectAtIndexedSubscript(portIndex uint) IVZVirtioConsolePortConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("objectAtIndexedSubscript:"), portIndex)
	return VZVirtioConsolePortConfigurationFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfigurationArray/setObject:atIndexedSubscript:
func (v VZVirtioConsolePortConfigurationArray) SetObjectAtIndexedSubscript(configuration IVZVirtioConsolePortConfiguration, portIndex uint) {
	objc.Send[objc.ID](v.ID, objc.Sel("setObject:atIndexedSubscript:"), configuration, portIndex)
}












// An unsigned integer that represents the maximum number of ports allocated
// by this device.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfigurationArray/maximumPortCount
func (v VZVirtioConsolePortConfigurationArray) MaximumPortCount() uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("maximumPortCount"))
	return rv
}
func (v VZVirtioConsolePortConfigurationArray) SetMaximumPortCount(value uint32) {
	objc.Send[struct{}](v.ID, objc.Sel("setMaximumPortCount:"), value)
}

























// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSerialPortConfiguration] class.
var (
	_VZSerialPortConfigurationClass     VZSerialPortConfigurationClass
	_VZSerialPortConfigurationClassOnce sync.Once
)

func getVZSerialPortConfigurationClass() VZSerialPortConfigurationClass {
	_VZSerialPortConfigurationClassOnce.Do(func() {
		_VZSerialPortConfigurationClass = VZSerialPortConfigurationClass{class: objc.GetClass("VZSerialPortConfiguration")}
	})
	return _VZSerialPortConfigurationClass
}

// GetVZSerialPortConfigurationClass returns the class object for VZSerialPortConfiguration.
func GetVZSerialPortConfigurationClass() VZSerialPortConfigurationClass {
	return getVZSerialPortConfigurationClass()
}

type VZSerialPortConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSerialPortConfigurationClass) Alloc() VZSerialPortConfiguration {
	rv := objc.Send[VZSerialPortConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The common configuration traits for serial port requests.
//
// # Overview
// 
// Don’t create a [VZSerialPortConfiguration] object directly. Instead,
// instantiate a concrete instance of one of its subclasses, such as
// [VZVirtioConsoleDeviceConfiguration]. Use the [VZSerialPortConfiguration.Attachment] property of this
// class to configure the medium through which serial communication happens.
//
// # Configuring the Attachment Point
//
//   - [VZSerialPortConfiguration.Attachment]: The object that defines how the configuration of the virtual machine’s serial port interfaces.
//   - [VZSerialPortConfiguration.SetAttachment]
//
// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortConfiguration
type VZSerialPortConfiguration struct {
	objectivec.Object
}

// VZSerialPortConfigurationFromID constructs a [VZSerialPortConfiguration] from an objc.ID.
//
// The common configuration traits for serial port requests.
func VZSerialPortConfigurationFromID(id objc.ID) VZSerialPortConfiguration {
	return VZSerialPortConfiguration{objectivec.Object{ID: id}}
}
// NOTE: VZSerialPortConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZSerialPortConfiguration] class.
//
// # Configuring the Attachment Point
//
//   - [IVZSerialPortConfiguration.Attachment]: The object that defines how the configuration of the virtual machine’s serial port interfaces.
//   - [IVZSerialPortConfiguration.SetAttachment]
//
// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortConfiguration
type IVZSerialPortConfiguration interface {
	objectivec.IObject

	// Topic: Configuring the Attachment Point

	// The object that defines how the configuration of the virtual machine’s serial port interfaces.
	Attachment() IVZSerialPortAttachment
	SetAttachment(value IVZSerialPortAttachment)
}

// Init initializes the instance.
func (s VZSerialPortConfiguration) Init() VZSerialPortConfiguration {
	rv := objc.Send[VZSerialPortConfiguration](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s VZSerialPortConfiguration) Autorelease() VZSerialPortConfiguration {
	rv := objc.Send[VZSerialPortConfiguration](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSerialPortConfiguration creates a new VZSerialPortConfiguration instance.
func NewVZSerialPortConfiguration() VZSerialPortConfiguration {
	class := getVZSerialPortConfigurationClass()
	rv := objc.Send[VZSerialPortConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The object that defines how the configuration of the virtual machine’s
// serial port interfaces.
//
// # Discussion
// 
// Assign an appropriate attachment object to this property, such as a
// [VZFileHandleSerialPortAttachment] or [VZFileSerialPortAttachment] object.
// When configuring the serial ports, the virtual machine uses the attachment
// to set up the serial port communications.
//
// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortConfiguration/attachment
func (s VZSerialPortConfiguration) Attachment() IVZSerialPortAttachment {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("attachment"))
	return VZSerialPortAttachmentFromID(objc.ID(rv))
}
func (s VZSerialPortConfiguration) SetAttachment(value IVZSerialPortAttachment) {
	objc.Send[struct{}](s.ID, objc.Sel("setAttachment:"), value)
}


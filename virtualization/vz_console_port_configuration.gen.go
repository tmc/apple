// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZConsolePortConfiguration] class.
var (
	_VZConsolePortConfigurationClass     VZConsolePortConfigurationClass
	_VZConsolePortConfigurationClassOnce sync.Once
)

func getVZConsolePortConfigurationClass() VZConsolePortConfigurationClass {
	_VZConsolePortConfigurationClassOnce.Do(func() {
		_VZConsolePortConfigurationClass = VZConsolePortConfigurationClass{class: objc.GetClass("VZConsolePortConfiguration")}
	})
	return _VZConsolePortConfigurationClass
}

// GetVZConsolePortConfigurationClass returns the class object for VZConsolePortConfiguration.
func GetVZConsolePortConfigurationClass() VZConsolePortConfigurationClass {
	return getVZConsolePortConfigurationClass()
}

type VZConsolePortConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZConsolePortConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZConsolePortConfigurationClass) Alloc() VZConsolePortConfiguration {
	rv := objc.Send[VZConsolePortConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The base class for a console port configuration.
//
// # Overview
//
// Don’t instantiate [VZConsolePortConfiguration] directly, instead use one
// of its subclasses like [VZVirtioConsolePortConfiguration].
//
// # Configuring the attachment
//
//   - [VZConsolePortConfiguration.Attachment]: The serial port attachment.
//   - [VZConsolePortConfiguration.SetAttachment]
//
// See: https://developer.apple.com/documentation/Virtualization/VZConsolePortConfiguration
type VZConsolePortConfiguration struct {
	objectivec.Object
}

// VZConsolePortConfigurationFromID constructs a [VZConsolePortConfiguration] from an objc.ID.
//
// The base class for a console port configuration.
func VZConsolePortConfigurationFromID(id objc.ID) VZConsolePortConfiguration {
	return VZConsolePortConfiguration{objectivec.Object{ID: id}}
}

// NOTE: VZConsolePortConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZConsolePortConfiguration] class.
//
// # Configuring the attachment
//
//   - [IVZConsolePortConfiguration.Attachment]: The serial port attachment.
//   - [IVZConsolePortConfiguration.SetAttachment]
//
// See: https://developer.apple.com/documentation/Virtualization/VZConsolePortConfiguration
type IVZConsolePortConfiguration interface {
	objectivec.IObject

	// Topic: Configuring the attachment

	// The serial port attachment.
	Attachment() IVZSerialPortAttachment
	SetAttachment(value IVZSerialPortAttachment)
}

// Init initializes the instance.
func (c VZConsolePortConfiguration) Init() VZConsolePortConfiguration {
	rv := objc.Send[VZConsolePortConfiguration](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c VZConsolePortConfiguration) Autorelease() VZConsolePortConfiguration {
	rv := objc.Send[VZConsolePortConfiguration](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZConsolePortConfiguration creates a new VZConsolePortConfiguration instance.
func NewVZConsolePortConfiguration() VZConsolePortConfiguration {
	class := getVZConsolePortConfigurationClass()
	rv := objc.Send[VZConsolePortConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The serial port attachment.
//
// See: https://developer.apple.com/documentation/Virtualization/VZConsolePortConfiguration/attachment
func (c VZConsolePortConfiguration) Attachment() IVZSerialPortAttachment {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("attachment"))
	return VZSerialPortAttachmentFromID(objc.ID(rv))
}
func (c VZConsolePortConfiguration) SetAttachment(value IVZSerialPortAttachment) {
	objc.Send[struct{}](c.ID, objc.Sel("setAttachment:"), value)
}

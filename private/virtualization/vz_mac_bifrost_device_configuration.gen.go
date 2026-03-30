// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacBifrostDeviceConfiguration] class.
var (
	_VZMacBifrostDeviceConfigurationClass     VZMacBifrostDeviceConfigurationClass
	_VZMacBifrostDeviceConfigurationClassOnce sync.Once
)

func getVZMacBifrostDeviceConfigurationClass() VZMacBifrostDeviceConfigurationClass {
	_VZMacBifrostDeviceConfigurationClassOnce.Do(func() {
		_VZMacBifrostDeviceConfigurationClass = VZMacBifrostDeviceConfigurationClass{class: objc.GetClass("_VZMacBifrostDeviceConfiguration")}
	})
	return _VZMacBifrostDeviceConfigurationClass
}

// GetVZMacBifrostDeviceConfigurationClass returns the class object for _VZMacBifrostDeviceConfiguration.
func GetVZMacBifrostDeviceConfigurationClass() VZMacBifrostDeviceConfigurationClass {
	return getVZMacBifrostDeviceConfigurationClass()
}

type VZMacBifrostDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacBifrostDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacBifrostDeviceConfigurationClass) Alloc() VZMacBifrostDeviceConfiguration {
	rv := objc.Send[VZMacBifrostDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMacBifrostDeviceConfiguration.InitWithAttachmentMMIOSize]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacBifrostDeviceConfiguration
type VZMacBifrostDeviceConfiguration struct {
	VZBifrostDeviceConfiguration
}

// VZMacBifrostDeviceConfigurationFromID constructs a [VZMacBifrostDeviceConfiguration] from an objc.ID.
func VZMacBifrostDeviceConfigurationFromID(id objc.ID) VZMacBifrostDeviceConfiguration {
	return VZMacBifrostDeviceConfiguration{VZBifrostDeviceConfiguration: VZBifrostDeviceConfigurationFromID(id)}
}

// Ensure VZMacBifrostDeviceConfiguration implements IVZMacBifrostDeviceConfiguration.
var _ IVZMacBifrostDeviceConfiguration = VZMacBifrostDeviceConfiguration{}

// An interface definition for the [VZMacBifrostDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZMacBifrostDeviceConfiguration.InitWithAttachmentMMIOSize]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacBifrostDeviceConfiguration
type IVZMacBifrostDeviceConfiguration interface {
	IVZBifrostDeviceConfiguration

	// Topic: Methods

	InitWithAttachmentMMIOSize(attachment objectivec.IObject, mIOSize uint64) VZMacBifrostDeviceConfiguration
}

// Init initializes the instance.
func (v VZMacBifrostDeviceConfiguration) Init() VZMacBifrostDeviceConfiguration {
	rv := objc.Send[VZMacBifrostDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacBifrostDeviceConfiguration) Autorelease() VZMacBifrostDeviceConfiguration {
	rv := objc.Send[VZMacBifrostDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacBifrostDeviceConfiguration creates a new VZMacBifrostDeviceConfiguration instance.
func NewVZMacBifrostDeviceConfiguration() VZMacBifrostDeviceConfiguration {
	class := getVZMacBifrostDeviceConfigurationClass()
	rv := objc.Send[VZMacBifrostDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBifrostDeviceConfiguration/initWithAttachment:MMIOSize:
func NewVZMacBifrostDeviceConfigurationWithAttachmentMMIOSize(attachment objectivec.IObject, mIOSize uint64) VZMacBifrostDeviceConfiguration {
	instance := getVZMacBifrostDeviceConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAttachment:MMIOSize:"), attachment, mIOSize)
	return VZMacBifrostDeviceConfigurationFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBifrostDeviceConfiguration/initWithAttachment:MMIOSize:
func (v VZMacBifrostDeviceConfiguration) InitWithAttachmentMMIOSize(attachment objectivec.IObject, mIOSize uint64) VZMacBifrostDeviceConfiguration {
	rv := objc.Send[VZMacBifrostDeviceConfiguration](v.ID, objc.Sel("initWithAttachment:MMIOSize:"), attachment, mIOSize)
	return rv
}

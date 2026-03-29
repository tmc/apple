// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZBifrostDeviceConfiguration] class.
var (
	_VZBifrostDeviceConfigurationClass     VZBifrostDeviceConfigurationClass
	_VZBifrostDeviceConfigurationClassOnce sync.Once
)

func getVZBifrostDeviceConfigurationClass() VZBifrostDeviceConfigurationClass {
	_VZBifrostDeviceConfigurationClassOnce.Do(func() {
		_VZBifrostDeviceConfigurationClass = VZBifrostDeviceConfigurationClass{class: objc.GetClass("_VZBifrostDeviceConfiguration")}
	})
	return _VZBifrostDeviceConfigurationClass
}

// GetVZBifrostDeviceConfigurationClass returns the class object for _VZBifrostDeviceConfiguration.
func GetVZBifrostDeviceConfigurationClass() VZBifrostDeviceConfigurationClass {
	return getVZBifrostDeviceConfigurationClass()
}

type VZBifrostDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZBifrostDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZBifrostDeviceConfigurationClass) Alloc() VZBifrostDeviceConfiguration {
	rv := objc.Send[VZBifrostDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZBifrostDeviceConfiguration.MMIOSize]
//   - [VZBifrostDeviceConfiguration.SetMMIOSize]
//   - [VZBifrostDeviceConfiguration._bifrostDevice]
//   - [VZBifrostDeviceConfiguration._initWithAttachmentMMIOSize]
//   - [VZBifrostDeviceConfiguration.Attachment]
//   - [VZBifrostDeviceConfiguration.SetAttachment]
//   - [VZBifrostDeviceConfiguration.EncodeWithEncoder]
//   - [VZBifrostDeviceConfiguration.DebugDescription]
//   - [VZBifrostDeviceConfiguration.Description]
//   - [VZBifrostDeviceConfiguration.Hash]
//   - [VZBifrostDeviceConfiguration.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostDeviceConfiguration
type VZBifrostDeviceConfiguration struct {
	objectivec.Object
}

// VZBifrostDeviceConfigurationFromID constructs a [VZBifrostDeviceConfiguration] from an objc.ID.
func VZBifrostDeviceConfigurationFromID(id objc.ID) VZBifrostDeviceConfiguration {
	return VZBifrostDeviceConfiguration{objectivec.Object{ID: id}}
}
// Ensure VZBifrostDeviceConfiguration implements IVZBifrostDeviceConfiguration.
var _ IVZBifrostDeviceConfiguration = VZBifrostDeviceConfiguration{}

// An interface definition for the [VZBifrostDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZBifrostDeviceConfiguration.MMIOSize]
//   - [IVZBifrostDeviceConfiguration.SetMMIOSize]
//   - [IVZBifrostDeviceConfiguration._bifrostDevice]
//   - [IVZBifrostDeviceConfiguration._initWithAttachmentMMIOSize]
//   - [IVZBifrostDeviceConfiguration.Attachment]
//   - [IVZBifrostDeviceConfiguration.SetAttachment]
//   - [IVZBifrostDeviceConfiguration.EncodeWithEncoder]
//   - [IVZBifrostDeviceConfiguration.DebugDescription]
//   - [IVZBifrostDeviceConfiguration.Description]
//   - [IVZBifrostDeviceConfiguration.Hash]
//   - [IVZBifrostDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostDeviceConfiguration
type IVZBifrostDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	MMIOSize() uint64
	SetMMIOSize(value uint64)
	_bifrostDevice() objectivec.IObject
	_initWithAttachmentMMIOSize(attachment objectivec.IObject, mIOSize uint64) objectivec.IObject
	Attachment() *VZBifrostAttachment
	SetAttachment(value *VZBifrostAttachment)
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZBifrostDeviceConfiguration) Init() VZBifrostDeviceConfiguration {
	rv := objc.Send[VZBifrostDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZBifrostDeviceConfiguration) Autorelease() VZBifrostDeviceConfiguration {
	rv := objc.Send[VZBifrostDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZBifrostDeviceConfiguration creates a new VZBifrostDeviceConfiguration instance.
func NewVZBifrostDeviceConfiguration() VZBifrostDeviceConfiguration {
	class := getVZBifrostDeviceConfigurationClass()
	rv := objc.Send[VZBifrostDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostDeviceConfiguration/_bifrostDevice
func (v VZBifrostDeviceConfiguration) _bifrostDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_bifrostDevice"))
	return objectivec.Object{ID: rv}
}

// BifrostDevice is an exported wrapper for the private method _bifrostDevice.
func (v VZBifrostDeviceConfiguration) BifrostDevice() objectivec.IObject {
	return v._bifrostDevice()
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostDeviceConfiguration/_initWithAttachment:MMIOSize:
func (v VZBifrostDeviceConfiguration) _initWithAttachmentMMIOSize(attachment objectivec.IObject, mIOSize uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_initWithAttachment:MMIOSize:"), attachment, mIOSize)
	return objectivec.Object{ID: rv}
}

// InitWithAttachmentMMIOSize is an exported wrapper for the private method _initWithAttachmentMMIOSize.
func (v VZBifrostDeviceConfiguration) InitWithAttachmentMMIOSize(attachment objectivec.IObject, mIOSize uint64) objectivec.IObject {
	return v._initWithAttachmentMMIOSize(attachment, mIOSize)
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostDeviceConfiguration/encodeWithEncoder:
func (v VZBifrostDeviceConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostDeviceConfiguration/MMIOSize
func (v VZBifrostDeviceConfiguration) MMIOSize() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("MMIOSize"))
	return rv
}
func (v VZBifrostDeviceConfiguration) SetMMIOSize(value uint64) {
	objc.Send[struct{}](v.ID, objc.Sel("setMMIOSize:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostDeviceConfiguration/attachment
func (v VZBifrostDeviceConfiguration) Attachment() *VZBifrostAttachment {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("attachment"))
	if rv == 0 {
		return nil
	}
	val := VZBifrostAttachmentFromID(objc.ID(rv))
	return &val
}
func (v VZBifrostDeviceConfiguration) SetAttachment(value *VZBifrostAttachment) {
	if value == nil {
		objc.Send[struct{}](v.ID, objc.Sel("setAttachment:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](v.ID, objc.Sel("setAttachment:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostDeviceConfiguration/debugDescription
func (v VZBifrostDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostDeviceConfiguration/description
func (v VZBifrostDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostDeviceConfiguration/hash
func (v VZBifrostDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostDeviceConfiguration/superclass
func (v VZBifrostDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}


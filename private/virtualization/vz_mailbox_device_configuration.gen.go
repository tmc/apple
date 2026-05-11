// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMailboxDeviceConfiguration] class.
var (
	_VZMailboxDeviceConfigurationClass     VZMailboxDeviceConfigurationClass
	_VZMailboxDeviceConfigurationClassOnce sync.Once
)

func getVZMailboxDeviceConfigurationClass() VZMailboxDeviceConfigurationClass {
	_VZMailboxDeviceConfigurationClassOnce.Do(func() {
		_VZMailboxDeviceConfigurationClass = VZMailboxDeviceConfigurationClass{class: objc.GetClass("_VZMailboxDeviceConfiguration")}
	})
	return _VZMailboxDeviceConfigurationClass
}

// GetVZMailboxDeviceConfigurationClass returns the class object for _VZMailboxDeviceConfiguration.
func GetVZMailboxDeviceConfigurationClass() VZMailboxDeviceConfigurationClass {
	return getVZMailboxDeviceConfigurationClass()
}

type VZMailboxDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMailboxDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMailboxDeviceConfigurationClass) Alloc() VZMailboxDeviceConfiguration {
	rv := objc.Send[VZMailboxDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMailboxDeviceConfiguration._init]
//   - [VZMailboxDeviceConfiguration._mailboxDevice]
//   - [VZMailboxDeviceConfiguration.Attachment]
//   - [VZMailboxDeviceConfiguration.SetAttachment]
//   - [VZMailboxDeviceConfiguration.EncodeWithEncoder]
//   - [VZMailboxDeviceConfiguration.ValidateWithError]
//   - [VZMailboxDeviceConfiguration.DebugDescription]
//   - [VZMailboxDeviceConfiguration.Description]
//   - [VZMailboxDeviceConfiguration.Hash]
//   - [VZMailboxDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceConfiguration
type VZMailboxDeviceConfiguration struct {
	objectivec.Object
}

// VZMailboxDeviceConfigurationFromID constructs a [VZMailboxDeviceConfiguration] from an objc.ID.
func VZMailboxDeviceConfigurationFromID(id objc.ID) VZMailboxDeviceConfiguration {
	return VZMailboxDeviceConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZMailboxDeviceConfiguration implements IVZMailboxDeviceConfiguration.
var _ IVZMailboxDeviceConfiguration = VZMailboxDeviceConfiguration{}

// An interface definition for the [VZMailboxDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZMailboxDeviceConfiguration._init]
//   - [IVZMailboxDeviceConfiguration._mailboxDevice]
//   - [IVZMailboxDeviceConfiguration.Attachment]
//   - [IVZMailboxDeviceConfiguration.SetAttachment]
//   - [IVZMailboxDeviceConfiguration.EncodeWithEncoder]
//   - [IVZMailboxDeviceConfiguration.ValidateWithError]
//   - [IVZMailboxDeviceConfiguration.DebugDescription]
//   - [IVZMailboxDeviceConfiguration.Description]
//   - [IVZMailboxDeviceConfiguration.Hash]
//   - [IVZMailboxDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceConfiguration
type IVZMailboxDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_mailboxDevice() objectivec.IObject
	Attachment() *VZMailboxDeviceAttachment
	SetAttachment(value *VZMailboxDeviceAttachment)
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	ValidateWithError() (bool, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZMailboxDeviceConfiguration) Init() VZMailboxDeviceConfiguration {
	rv := objc.Send[VZMailboxDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMailboxDeviceConfiguration) Autorelease() VZMailboxDeviceConfiguration {
	rv := objc.Send[VZMailboxDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMailboxDeviceConfiguration creates a new VZMailboxDeviceConfiguration instance.
func NewVZMailboxDeviceConfiguration() VZMailboxDeviceConfiguration {
	class := getVZMailboxDeviceConfigurationClass()
	rv := objc.Send[VZMailboxDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceConfiguration/_init
func (v VZMailboxDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceConfiguration/encodeWithEncoder:
func (v VZMailboxDeviceConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceConfiguration/validateWithError:
func (v VZMailboxDeviceConfiguration) ValidateWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("validateWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceConfiguration/_mailboxDevice
func (v VZMailboxDeviceConfiguration) _mailboxDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_mailboxDevice"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceConfiguration/attachment
func (v VZMailboxDeviceConfiguration) Attachment() *VZMailboxDeviceAttachment {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("attachment"))
	if rv == 0 {
		return nil
	}
	val := VZMailboxDeviceAttachmentFromID(objc.ID(rv))
	return &val
}
func (v VZMailboxDeviceConfiguration) SetAttachment(value *VZMailboxDeviceAttachment) {
	if value == nil {
		objc.Send[struct{}](v.ID, objc.Sel("setAttachment:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](v.ID, objc.Sel("setAttachment:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceConfiguration/debugDescription
func (v VZMailboxDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceConfiguration/description
func (v VZMailboxDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceConfiguration/hash
func (v VZMailboxDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceConfiguration/superclass
func (v VZMailboxDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}

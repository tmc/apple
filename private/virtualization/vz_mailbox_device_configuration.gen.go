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
	rv := objc.SendIfResponds[VZMailboxDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMailboxDeviceConfiguration._init]
//   - [VZMailboxDeviceConfiguration._mailboxDevice]
//   - [VZMailboxDeviceConfiguration.Attachment]
//   - [VZMailboxDeviceConfiguration.SetAttachment]
//   - [VZMailboxDeviceConfiguration.ValidateWithError]
//   - [VZMailboxDeviceConfiguration.DebugDescription]
//   - [VZMailboxDeviceConfiguration.Description]
//   - [VZMailboxDeviceConfiguration.Hash]
//   - [VZMailboxDeviceConfiguration.Superclass]
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
//   - [IVZMailboxDeviceConfiguration.ValidateWithError]
//   - [IVZMailboxDeviceConfiguration.DebugDescription]
//   - [IVZMailboxDeviceConfiguration.Description]
//   - [IVZMailboxDeviceConfiguration.Hash]
//   - [IVZMailboxDeviceConfiguration.Superclass]
type IVZMailboxDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_mailboxDevice() unsafe.Pointer
	Attachment() IVZMailboxDeviceAttachment
	SetAttachment(value IVZMailboxDeviceAttachment)
	ValidateWithError() (bool, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZMailboxDeviceConfiguration) Init() VZMailboxDeviceConfiguration {
	rv := objc.SendIfResponds[VZMailboxDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMailboxDeviceConfiguration) Autorelease() VZMailboxDeviceConfiguration {
	rv := objc.SendIfResponds[VZMailboxDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMailboxDeviceConfiguration creates a new VZMailboxDeviceConfiguration instance.
func NewVZMailboxDeviceConfiguration() VZMailboxDeviceConfiguration {
	class := getVZMailboxDeviceConfigurationClass()
	rv := objc.SendIfResponds[VZMailboxDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZMailboxDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
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

func (v VZMailboxDeviceConfiguration) _mailboxDevice() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](v.ID, objc.Sel("_mailboxDevice"))
	return rv
}

// CanMailboxDevice reports whether the receiver responds to the private selector _mailboxDevice.
func (v VZMailboxDeviceConfiguration) CanMailboxDevice() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_mailboxDevice"))
}

// MailboxDevice is an exported wrapper for the private property _mailboxDevice.
func (v VZMailboxDeviceConfiguration) MailboxDevice() (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_mailboxDevice")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_mailboxDevice"}
	}
	return v._mailboxDevice(), nil
}
func (v VZMailboxDeviceConfiguration) Attachment() IVZMailboxDeviceAttachment {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("attachment"))
	return VZMailboxDeviceAttachmentFromID(objc.ID(rv))
}
func (v VZMailboxDeviceConfiguration) SetAttachment(value IVZMailboxDeviceAttachment) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setAttachment:"), value)
}
func (v VZMailboxDeviceConfiguration) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZMailboxDeviceConfiguration) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZMailboxDeviceConfiguration) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZMailboxDeviceConfiguration) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

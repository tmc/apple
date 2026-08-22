// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[VZBifrostDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZBifrostDeviceConfiguration.MMIOSize]
//   - [VZBifrostDeviceConfiguration.SetMMIOSize]
//   - [VZBifrostDeviceConfiguration._initWithAttachmentMMIOSize]
//   - [VZBifrostDeviceConfiguration.Attachment]
//   - [VZBifrostDeviceConfiguration.SetAttachment]
//   - [VZBifrostDeviceConfiguration.DebugDescription]
//   - [VZBifrostDeviceConfiguration.Description]
//   - [VZBifrostDeviceConfiguration.Hash]
//   - [VZBifrostDeviceConfiguration.Superclass]
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
//   - [IVZBifrostDeviceConfiguration._initWithAttachmentMMIOSize]
//   - [IVZBifrostDeviceConfiguration.Attachment]
//   - [IVZBifrostDeviceConfiguration.SetAttachment]
//   - [IVZBifrostDeviceConfiguration.DebugDescription]
//   - [IVZBifrostDeviceConfiguration.Description]
//   - [IVZBifrostDeviceConfiguration.Hash]
//   - [IVZBifrostDeviceConfiguration.Superclass]
type IVZBifrostDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	MMIOSize() uint64
	SetMMIOSize(value uint64)
	_initWithAttachmentMMIOSize(attachment objectivec.IObject, mIOSize uint64) objectivec.IObject
	Attachment() IVZBifrostAttachment
	SetAttachment(value IVZBifrostAttachment)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZBifrostDeviceConfiguration) Init() VZBifrostDeviceConfiguration {
	rv := objc.SendIfResponds[VZBifrostDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZBifrostDeviceConfiguration) Autorelease() VZBifrostDeviceConfiguration {
	rv := objc.SendIfResponds[VZBifrostDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZBifrostDeviceConfiguration creates a new VZBifrostDeviceConfiguration instance.
func NewVZBifrostDeviceConfiguration() VZBifrostDeviceConfiguration {
	class := getVZBifrostDeviceConfigurationClass()
	rv := objc.SendIfResponds[VZBifrostDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZBifrostDeviceConfiguration) _initWithAttachmentMMIOSize(attachment objectivec.IObject, mIOSize uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_initWithAttachment:MMIOSize:"), attachment, mIOSize)
	return objectivec.Object{ID: rv}
}

// InitWithAttachmentMMIOSize is an exported wrapper for the private method _initWithAttachmentMMIOSize.
func (v VZBifrostDeviceConfiguration) InitWithAttachmentMMIOSize(attachment objectivec.IObject, mIOSize uint64) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_initWithAttachment:MMIOSize:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_initWithAttachment:MMIOSize:"}
		return nil, err
	}
	return v._initWithAttachmentMMIOSize(attachment, mIOSize), nil
}

// CanInitWithAttachmentMMIOSize reports whether the receiver responds to the private selector _initWithAttachment:MMIOSize:.
func (v VZBifrostDeviceConfiguration) CanInitWithAttachmentMMIOSize() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_initWithAttachment:MMIOSize:"))
}

func (v VZBifrostDeviceConfiguration) MMIOSize() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("MMIOSize"))
	return rv
}
func (v VZBifrostDeviceConfiguration) SetMMIOSize(value uint64) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setMMIOSize:"), value)
}
func (v VZBifrostDeviceConfiguration) Attachment() IVZBifrostAttachment {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("attachment"))
	return VZBifrostAttachmentFromID(objc.ID(rv))
}
func (v VZBifrostDeviceConfiguration) SetAttachment(value IVZBifrostAttachment) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setAttachment:"), value)
}
func (v VZBifrostDeviceConfiguration) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZBifrostDeviceConfiguration) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZBifrostDeviceConfiguration) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZBifrostDeviceConfiguration) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

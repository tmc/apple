// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZGraphicsDisplayConfiguration] class.
var (
	_VZGraphicsDisplayConfigurationClass     VZGraphicsDisplayConfigurationClass
	_VZGraphicsDisplayConfigurationClassOnce sync.Once
)

func getVZGraphicsDisplayConfigurationClass() VZGraphicsDisplayConfigurationClass {
	_VZGraphicsDisplayConfigurationClassOnce.Do(func() {
		_VZGraphicsDisplayConfigurationClass = VZGraphicsDisplayConfigurationClass{class: objc.GetClass("VZGraphicsDisplayConfiguration")}
	})
	return _VZGraphicsDisplayConfigurationClass
}

// GetVZGraphicsDisplayConfigurationClass returns the class object for VZGraphicsDisplayConfiguration.
func GetVZGraphicsDisplayConfigurationClass() VZGraphicsDisplayConfigurationClass {
	return getVZGraphicsDisplayConfigurationClass()
}

type VZGraphicsDisplayConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZGraphicsDisplayConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZGraphicsDisplayConfigurationClass) Alloc() VZGraphicsDisplayConfiguration {
	rv := objc.SendIfResponds[VZGraphicsDisplayConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZGraphicsDisplayConfiguration._init]
//   - [VZGraphicsDisplayConfiguration._initWithConfiguration]
//   - [VZGraphicsDisplayConfiguration._setUUID]
//   - [VZGraphicsDisplayConfiguration._uuid]
//   - [VZGraphicsDisplayConfiguration.Set_uuid]
type VZGraphicsDisplayConfiguration struct {
	objectivec.Object
}

// VZGraphicsDisplayConfigurationFromID constructs a [VZGraphicsDisplayConfiguration] from an objc.ID.
func VZGraphicsDisplayConfigurationFromID(id objc.ID) VZGraphicsDisplayConfiguration {
	return VZGraphicsDisplayConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZGraphicsDisplayConfiguration implements IVZGraphicsDisplayConfiguration.
var _ IVZGraphicsDisplayConfiguration = VZGraphicsDisplayConfiguration{}

// An interface definition for the [VZGraphicsDisplayConfiguration] class.
//
// # Methods
//
//   - [IVZGraphicsDisplayConfiguration._init]
//   - [IVZGraphicsDisplayConfiguration._initWithConfiguration]
//   - [IVZGraphicsDisplayConfiguration._setUUID]
//   - [IVZGraphicsDisplayConfiguration._uuid]
//   - [IVZGraphicsDisplayConfiguration.Set_uuid]
type IVZGraphicsDisplayConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_initWithConfiguration(configuration unsafe.Pointer) objectivec.IObject
	_setUUID(uuid objectivec.IObject)
	_uuid() foundation.NSUUID
	Set_uuid(value foundation.NSUUID)
}

// Init initializes the instance.
func (v VZGraphicsDisplayConfiguration) Init() VZGraphicsDisplayConfiguration {
	rv := objc.SendIfResponds[VZGraphicsDisplayConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZGraphicsDisplayConfiguration) Autorelease() VZGraphicsDisplayConfiguration {
	rv := objc.SendIfResponds[VZGraphicsDisplayConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGraphicsDisplayConfiguration creates a new VZGraphicsDisplayConfiguration instance.
func NewVZGraphicsDisplayConfiguration() VZGraphicsDisplayConfiguration {
	class := getVZGraphicsDisplayConfigurationClass()
	rv := objc.SendIfResponds[VZGraphicsDisplayConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZGraphicsDisplayConfiguration) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
func (v VZGraphicsDisplayConfiguration) _initWithConfiguration(configuration unsafe.Pointer) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_initWithConfiguration:"), configuration)
	return objectivec.Object{ID: rv}
}

// InitWithConfiguration is an exported wrapper for the private method _initWithConfiguration.
func (v VZGraphicsDisplayConfiguration) InitWithConfiguration(configuration unsafe.Pointer) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_initWithConfiguration:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_initWithConfiguration:"}
		return nil, err
	}
	return v._initWithConfiguration(configuration), nil
}

// CanInitWithConfiguration reports whether the receiver responds to the private selector _initWithConfiguration:.
func (v VZGraphicsDisplayConfiguration) CanInitWithConfiguration() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_initWithConfiguration:"))
}
func (v VZGraphicsDisplayConfiguration) _setUUID(uuid objectivec.IObject) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_setUUID:"), uuid)
}

// SetUUID is an exported wrapper for the private method _setUUID.
func (v VZGraphicsDisplayConfiguration) SetUUID(uuid objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setUUID:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setUUID:"}
		return err
	}
	v._setUUID(uuid)
	return nil
}

// CanSetUUID reports whether the receiver responds to the private selector _setUUID:.
func (v VZGraphicsDisplayConfiguration) CanSetUUID() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setUUID:"))
}

func (v VZGraphicsDisplayConfiguration) _uuid() foundation.NSUUID {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_uuid"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}

// CanUuid reports whether the receiver responds to the private selector _uuid.
func (v VZGraphicsDisplayConfiguration) CanUuid() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_uuid"))
}

// Uuid is an exported wrapper for the private property _uuid.
func (v VZGraphicsDisplayConfiguration) Uuid() (foundation.NSUUID, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_uuid")) {
		return foundation.NSUUID{}, &objc.UnrecognizedSelectorError{Selector: "_uuid"}
	}
	return v._uuid(), nil
}
func (v VZGraphicsDisplayConfiguration) Set_uuid(value foundation.NSUUID) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("set_uuid:"), value)
}

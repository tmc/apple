// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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
	rv := objc.Send[VZGraphicsDisplayConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZGraphicsDisplayConfiguration._init]
//   - [VZGraphicsDisplayConfiguration._initWithConfiguration]
//   - [VZGraphicsDisplayConfiguration._setUUID]
//   - [VZGraphicsDisplayConfiguration._uuid]
//   - [VZGraphicsDisplayConfiguration.Set_uuid]
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplayConfiguration
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
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplayConfiguration
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
func (g VZGraphicsDisplayConfiguration) Init() VZGraphicsDisplayConfiguration {
	rv := objc.Send[VZGraphicsDisplayConfiguration](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g VZGraphicsDisplayConfiguration) Autorelease() VZGraphicsDisplayConfiguration {
	rv := objc.Send[VZGraphicsDisplayConfiguration](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGraphicsDisplayConfiguration creates a new VZGraphicsDisplayConfiguration instance.
func NewVZGraphicsDisplayConfiguration() VZGraphicsDisplayConfiguration {
	class := getVZGraphicsDisplayConfigurationClass()
	rv := objc.Send[VZGraphicsDisplayConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplayConfiguration/_init
func (g VZGraphicsDisplayConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplayConfiguration/_initWithConfiguration:
func (g VZGraphicsDisplayConfiguration) _initWithConfiguration(configuration unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_initWithConfiguration:"), configuration)
	return objectivec.Object{ID: rv}
}

// InitWithConfiguration is an exported wrapper for the private method _initWithConfiguration.
func (g VZGraphicsDisplayConfiguration) InitWithConfiguration(configuration unsafe.Pointer) objectivec.IObject {
	return g._initWithConfiguration(configuration)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplayConfiguration/_setUUID:
func (g VZGraphicsDisplayConfiguration) _setUUID(uuid objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("_setUUID:"), uuid)
}

// SetUUID is an exported wrapper for the private method _setUUID.
func (g VZGraphicsDisplayConfiguration) SetUUID(uuid objectivec.IObject) {
	g._setUUID(uuid)
}

// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplayConfiguration/_uuid
func (g VZGraphicsDisplayConfiguration) _uuid() foundation.NSUUID {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_uuid"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (g VZGraphicsDisplayConfiguration) Set_uuid(value foundation.NSUUID) {
	objc.Send[struct{}](g.ID, objc.Sel("set_uuid:"), value)
}


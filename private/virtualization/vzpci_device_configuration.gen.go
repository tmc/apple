// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZPCIDeviceConfiguration] class.
var (
	_VZPCIDeviceConfigurationClass     VZPCIDeviceConfigurationClass
	_VZPCIDeviceConfigurationClassOnce sync.Once
)

func getVZPCIDeviceConfigurationClass() VZPCIDeviceConfigurationClass {
	_VZPCIDeviceConfigurationClassOnce.Do(func() {
		_VZPCIDeviceConfigurationClass = VZPCIDeviceConfigurationClass{class: objc.GetClass("_VZPCIDeviceConfiguration")}
	})
	return _VZPCIDeviceConfigurationClass
}

// GetVZPCIDeviceConfigurationClass returns the class object for _VZPCIDeviceConfiguration.
func GetVZPCIDeviceConfigurationClass() VZPCIDeviceConfigurationClass {
	return getVZPCIDeviceConfigurationClass()
}

type VZPCIDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZPCIDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZPCIDeviceConfigurationClass) Alloc() VZPCIDeviceConfiguration {
	rv := objc.SendIfResponds[VZPCIDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZPCIDeviceConfiguration._init]
//   - [VZPCIDeviceConfiguration.DebugDescription]
//   - [VZPCIDeviceConfiguration.Description]
//   - [VZPCIDeviceConfiguration.Hash]
//   - [VZPCIDeviceConfiguration.Superclass]
type VZPCIDeviceConfiguration struct {
	objectivec.Object
}

// VZPCIDeviceConfigurationFromID constructs a [VZPCIDeviceConfiguration] from an objc.ID.
func VZPCIDeviceConfigurationFromID(id objc.ID) VZPCIDeviceConfiguration {
	return VZPCIDeviceConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZPCIDeviceConfiguration implements IVZPCIDeviceConfiguration.
var _ IVZPCIDeviceConfiguration = VZPCIDeviceConfiguration{}

// An interface definition for the [VZPCIDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZPCIDeviceConfiguration._init]
//   - [IVZPCIDeviceConfiguration.DebugDescription]
//   - [IVZPCIDeviceConfiguration.Description]
//   - [IVZPCIDeviceConfiguration.Hash]
//   - [IVZPCIDeviceConfiguration.Superclass]
type IVZPCIDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZPCIDeviceConfiguration) Init() VZPCIDeviceConfiguration {
	rv := objc.SendIfResponds[VZPCIDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZPCIDeviceConfiguration) Autorelease() VZPCIDeviceConfiguration {
	rv := objc.SendIfResponds[VZPCIDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZPCIDeviceConfiguration creates a new VZPCIDeviceConfiguration instance.
func NewVZPCIDeviceConfiguration() VZPCIDeviceConfiguration {
	class := getVZPCIDeviceConfigurationClass()
	rv := objc.SendIfResponds[VZPCIDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZPCIDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

func (v VZPCIDeviceConfiguration) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZPCIDeviceConfiguration) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZPCIDeviceConfiguration) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZPCIDeviceConfiguration) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

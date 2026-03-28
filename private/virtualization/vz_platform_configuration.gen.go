// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZPlatformConfiguration] class.
var (
	_VZPlatformConfigurationClass     VZPlatformConfigurationClass
	_VZPlatformConfigurationClassOnce sync.Once
)

func getVZPlatformConfigurationClass() VZPlatformConfigurationClass {
	_VZPlatformConfigurationClassOnce.Do(func() {
		_VZPlatformConfigurationClass = VZPlatformConfigurationClass{class: objc.GetClass("VZPlatformConfiguration")}
	})
	return _VZPlatformConfigurationClass
}

// GetVZPlatformConfigurationClass returns the class object for VZPlatformConfiguration.
func GetVZPlatformConfigurationClass() VZPlatformConfigurationClass {
	return getVZPlatformConfigurationClass()
}

type VZPlatformConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZPlatformConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZPlatformConfigurationClass) Alloc() VZPlatformConfiguration {
	rv := objc.Send[VZPlatformConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZPlatformConfiguration._init]
//   - [VZPlatformConfiguration._platform]
//   - [VZPlatformConfiguration.DebugDescription]
//   - [VZPlatformConfiguration.Description]
//   - [VZPlatformConfiguration.Hash]
//   - [VZPlatformConfiguration.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/VZPlatformConfiguration
type VZPlatformConfiguration struct {
	objectivec.Object
}

// VZPlatformConfigurationFromID constructs a [VZPlatformConfiguration] from an objc.ID.
func VZPlatformConfigurationFromID(id objc.ID) VZPlatformConfiguration {
	return VZPlatformConfiguration{objectivec.Object{ID: id}}
}
// Ensure VZPlatformConfiguration implements IVZPlatformConfiguration.
var _ IVZPlatformConfiguration = VZPlatformConfiguration{}

// An interface definition for the [VZPlatformConfiguration] class.
//
// # Methods
//
//   - [IVZPlatformConfiguration._init]
//   - [IVZPlatformConfiguration._platform]
//   - [IVZPlatformConfiguration.DebugDescription]
//   - [IVZPlatformConfiguration.Description]
//   - [IVZPlatformConfiguration.Hash]
//   - [IVZPlatformConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZPlatformConfiguration
type IVZPlatformConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_platform() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (p VZPlatformConfiguration) Init() VZPlatformConfiguration {
	rv := objc.Send[VZPlatformConfiguration](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p VZPlatformConfiguration) Autorelease() VZPlatformConfiguration {
	rv := objc.Send[VZPlatformConfiguration](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZPlatformConfiguration creates a new VZPlatformConfiguration instance.
func NewVZPlatformConfiguration() VZPlatformConfiguration {
	class := getVZPlatformConfigurationClass()
	rv := objc.Send[VZPlatformConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZPlatformConfiguration/_init
func (p VZPlatformConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZPlatformConfiguration/_platform
func (p VZPlatformConfiguration) _platform() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("_platform"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Virtualization/VZPlatformConfiguration/debugDescription
func (p VZPlatformConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZPlatformConfiguration/description
func (p VZPlatformConfiguration) Description() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZPlatformConfiguration/hash
func (p VZPlatformConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZPlatformConfiguration/superclass
func (p VZPlatformConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](p.ID, objc.Sel("superclass"))
	return rv
}


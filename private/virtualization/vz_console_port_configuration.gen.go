// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
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

// # Methods
//
//   - [VZConsolePortConfiguration._init]
//   - [VZConsolePortConfiguration.DebugDescription]
//   - [VZConsolePortConfiguration.Description]
//   - [VZConsolePortConfiguration.Hash]
//   - [VZConsolePortConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZConsolePortConfiguration
type VZConsolePortConfiguration struct {
	objectivec.Object
}

// VZConsolePortConfigurationFromID constructs a [VZConsolePortConfiguration] from an objc.ID.
func VZConsolePortConfigurationFromID(id objc.ID) VZConsolePortConfiguration {
	return VZConsolePortConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZConsolePortConfiguration implements IVZConsolePortConfiguration.
var _ IVZConsolePortConfiguration = VZConsolePortConfiguration{}

// An interface definition for the [VZConsolePortConfiguration] class.
//
// # Methods
//
//   - [IVZConsolePortConfiguration._init]
//   - [IVZConsolePortConfiguration.DebugDescription]
//   - [IVZConsolePortConfiguration.Description]
//   - [IVZConsolePortConfiguration.Hash]
//   - [IVZConsolePortConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZConsolePortConfiguration
type IVZConsolePortConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
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

// See: https://developer.apple.com/documentation/Virtualization/VZConsolePortConfiguration/_init
func (c VZConsolePortConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZConsolePortConfiguration/debugDescription
func (c VZConsolePortConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZConsolePortConfiguration/description
func (c VZConsolePortConfiguration) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZConsolePortConfiguration/hash
func (c VZConsolePortConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZConsolePortConfiguration/superclass
func (c VZConsolePortConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}

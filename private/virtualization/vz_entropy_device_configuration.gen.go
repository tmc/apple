// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZEntropyDeviceConfiguration] class.
var (
	_VZEntropyDeviceConfigurationClass     VZEntropyDeviceConfigurationClass
	_VZEntropyDeviceConfigurationClassOnce sync.Once
)

func getVZEntropyDeviceConfigurationClass() VZEntropyDeviceConfigurationClass {
	_VZEntropyDeviceConfigurationClassOnce.Do(func() {
		_VZEntropyDeviceConfigurationClass = VZEntropyDeviceConfigurationClass{class: objc.GetClass("VZEntropyDeviceConfiguration")}
	})
	return _VZEntropyDeviceConfigurationClass
}

// GetVZEntropyDeviceConfigurationClass returns the class object for VZEntropyDeviceConfiguration.
func GetVZEntropyDeviceConfigurationClass() VZEntropyDeviceConfigurationClass {
	return getVZEntropyDeviceConfigurationClass()
}

type VZEntropyDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZEntropyDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZEntropyDeviceConfigurationClass) Alloc() VZEntropyDeviceConfiguration {
	rv := objc.Send[VZEntropyDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZEntropyDeviceConfiguration._entropyDevice]
//   - [VZEntropyDeviceConfiguration._init]
//   - [VZEntropyDeviceConfiguration.DebugDescription]
//   - [VZEntropyDeviceConfiguration.Description]
//   - [VZEntropyDeviceConfiguration.Hash]
//   - [VZEntropyDeviceConfiguration.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/VZEntropyDeviceConfiguration
type VZEntropyDeviceConfiguration struct {
	objectivec.Object
}

// VZEntropyDeviceConfigurationFromID constructs a [VZEntropyDeviceConfiguration] from an objc.ID.
func VZEntropyDeviceConfigurationFromID(id objc.ID) VZEntropyDeviceConfiguration {
	return VZEntropyDeviceConfiguration{objectivec.Object{ID: id}}
}
// Ensure VZEntropyDeviceConfiguration implements IVZEntropyDeviceConfiguration.
var _ IVZEntropyDeviceConfiguration = VZEntropyDeviceConfiguration{}

// An interface definition for the [VZEntropyDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZEntropyDeviceConfiguration._entropyDevice]
//   - [IVZEntropyDeviceConfiguration._init]
//   - [IVZEntropyDeviceConfiguration.DebugDescription]
//   - [IVZEntropyDeviceConfiguration.Description]
//   - [IVZEntropyDeviceConfiguration.Hash]
//   - [IVZEntropyDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZEntropyDeviceConfiguration
type IVZEntropyDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_entropyDevice() int
	_init() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (e VZEntropyDeviceConfiguration) Init() VZEntropyDeviceConfiguration {
	rv := objc.Send[VZEntropyDeviceConfiguration](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e VZEntropyDeviceConfiguration) Autorelease() VZEntropyDeviceConfiguration {
	rv := objc.Send[VZEntropyDeviceConfiguration](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZEntropyDeviceConfiguration creates a new VZEntropyDeviceConfiguration instance.
func NewVZEntropyDeviceConfiguration() VZEntropyDeviceConfiguration {
	class := getVZEntropyDeviceConfigurationClass()
	rv := objc.Send[VZEntropyDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZEntropyDeviceConfiguration/_init
func (e VZEntropyDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZEntropyDeviceConfiguration/_entropyDevice
func (e VZEntropyDeviceConfiguration) _entropyDevice() int {
	rv := objc.Send[int](e.ID, objc.Sel("_entropyDevice"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZEntropyDeviceConfiguration/debugDescription
func (e VZEntropyDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZEntropyDeviceConfiguration/description
func (e VZEntropyDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZEntropyDeviceConfiguration/hash
func (e VZEntropyDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZEntropyDeviceConfiguration/superclass
func (e VZEntropyDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](e.ID, objc.Sel("superclass"))
	return rv
}


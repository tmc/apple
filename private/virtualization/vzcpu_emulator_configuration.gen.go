// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZCPUEmulatorConfiguration] class.
var (
	_VZCPUEmulatorConfigurationClass     VZCPUEmulatorConfigurationClass
	_VZCPUEmulatorConfigurationClassOnce sync.Once
)

func getVZCPUEmulatorConfigurationClass() VZCPUEmulatorConfigurationClass {
	_VZCPUEmulatorConfigurationClassOnce.Do(func() {
		_VZCPUEmulatorConfigurationClass = VZCPUEmulatorConfigurationClass{class: objc.GetClass("_VZCPUEmulatorConfiguration")}
	})
	return _VZCPUEmulatorConfigurationClass
}

// GetVZCPUEmulatorConfigurationClass returns the class object for _VZCPUEmulatorConfiguration.
func GetVZCPUEmulatorConfigurationClass() VZCPUEmulatorConfigurationClass {
	return getVZCPUEmulatorConfigurationClass()
}

type VZCPUEmulatorConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCPUEmulatorConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCPUEmulatorConfigurationClass) Alloc() VZCPUEmulatorConfiguration {
	rv := objc.Send[VZCPUEmulatorConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZCPUEmulatorConfiguration._cpuEmulator]
//   - [VZCPUEmulatorConfiguration._init]
//   - [VZCPUEmulatorConfiguration.EncodeWithEncoder]
//   - [VZCPUEmulatorConfiguration.DebugDescription]
//   - [VZCPUEmulatorConfiguration.Description]
//   - [VZCPUEmulatorConfiguration.Hash]
//   - [VZCPUEmulatorConfiguration.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/_VZCPUEmulatorConfiguration
type VZCPUEmulatorConfiguration struct {
	objectivec.Object
}

// VZCPUEmulatorConfigurationFromID constructs a [VZCPUEmulatorConfiguration] from an objc.ID.
func VZCPUEmulatorConfigurationFromID(id objc.ID) VZCPUEmulatorConfiguration {
	return VZCPUEmulatorConfiguration{objectivec.Object{ID: id}}
}
// Ensure VZCPUEmulatorConfiguration implements IVZCPUEmulatorConfiguration.
var _ IVZCPUEmulatorConfiguration = VZCPUEmulatorConfiguration{}

// An interface definition for the [VZCPUEmulatorConfiguration] class.
//
// # Methods
//
//   - [IVZCPUEmulatorConfiguration._cpuEmulator]
//   - [IVZCPUEmulatorConfiguration._init]
//   - [IVZCPUEmulatorConfiguration.EncodeWithEncoder]
//   - [IVZCPUEmulatorConfiguration.DebugDescription]
//   - [IVZCPUEmulatorConfiguration.Description]
//   - [IVZCPUEmulatorConfiguration.Hash]
//   - [IVZCPUEmulatorConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCPUEmulatorConfiguration
type IVZCPUEmulatorConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_cpuEmulator() objectivec.IObject
	_init() objectivec.IObject
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZCPUEmulatorConfiguration) Init() VZCPUEmulatorConfiguration {
	rv := objc.Send[VZCPUEmulatorConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCPUEmulatorConfiguration) Autorelease() VZCPUEmulatorConfiguration {
	rv := objc.Send[VZCPUEmulatorConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCPUEmulatorConfiguration creates a new VZCPUEmulatorConfiguration instance.
func NewVZCPUEmulatorConfiguration() VZCPUEmulatorConfiguration {
	class := getVZCPUEmulatorConfigurationClass()
	rv := objc.Send[VZCPUEmulatorConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCPUEmulatorConfiguration/_cpuEmulator
func (v VZCPUEmulatorConfiguration) _cpuEmulator() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_cpuEmulator"))
	return objectivec.Object{ID: rv}
}

// CpuEmulator is an exported wrapper for the private method _cpuEmulator.
func (v VZCPUEmulatorConfiguration) CpuEmulator() objectivec.IObject {
	return v._cpuEmulator()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCPUEmulatorConfiguration/_init
func (v VZCPUEmulatorConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCPUEmulatorConfiguration/encodeWithEncoder:
func (v VZCPUEmulatorConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCPUEmulatorConfiguration/debugDescription
func (v VZCPUEmulatorConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCPUEmulatorConfiguration/description
func (v VZCPUEmulatorConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCPUEmulatorConfiguration/hash
func (v VZCPUEmulatorConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCPUEmulatorConfiguration/superclass
func (v VZCPUEmulatorConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}


// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[VZCPUEmulatorConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZCPUEmulatorConfiguration._init]
//   - [VZCPUEmulatorConfiguration.DebugDescription]
//   - [VZCPUEmulatorConfiguration.Description]
//   - [VZCPUEmulatorConfiguration.Hash]
//   - [VZCPUEmulatorConfiguration.Superclass]
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
//   - [IVZCPUEmulatorConfiguration._init]
//   - [IVZCPUEmulatorConfiguration.DebugDescription]
//   - [IVZCPUEmulatorConfiguration.Description]
//   - [IVZCPUEmulatorConfiguration.Hash]
//   - [IVZCPUEmulatorConfiguration.Superclass]
type IVZCPUEmulatorConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZCPUEmulatorConfiguration) Init() VZCPUEmulatorConfiguration {
	rv := objc.SendIfResponds[VZCPUEmulatorConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCPUEmulatorConfiguration) Autorelease() VZCPUEmulatorConfiguration {
	rv := objc.SendIfResponds[VZCPUEmulatorConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCPUEmulatorConfiguration creates a new VZCPUEmulatorConfiguration instance.
func NewVZCPUEmulatorConfiguration() VZCPUEmulatorConfiguration {
	class := getVZCPUEmulatorConfigurationClass()
	rv := objc.SendIfResponds[VZCPUEmulatorConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZCPUEmulatorConfiguration) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

func (v VZCPUEmulatorConfiguration) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZCPUEmulatorConfiguration) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZCPUEmulatorConfiguration) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZCPUEmulatorConfiguration) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

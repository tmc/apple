// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZIOPCoprocessorConfiguration] class.
var (
	_VZIOPCoprocessorConfigurationClass     VZIOPCoprocessorConfigurationClass
	_VZIOPCoprocessorConfigurationClassOnce sync.Once
)

func getVZIOPCoprocessorConfigurationClass() VZIOPCoprocessorConfigurationClass {
	_VZIOPCoprocessorConfigurationClassOnce.Do(func() {
		_VZIOPCoprocessorConfigurationClass = VZIOPCoprocessorConfigurationClass{class: objc.GetClass("_VZIOPCoprocessorConfiguration")}
	})
	return _VZIOPCoprocessorConfigurationClass
}

// GetVZIOPCoprocessorConfigurationClass returns the class object for _VZIOPCoprocessorConfiguration.
func GetVZIOPCoprocessorConfigurationClass() VZIOPCoprocessorConfigurationClass {
	return getVZIOPCoprocessorConfigurationClass()
}

type VZIOPCoprocessorConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZIOPCoprocessorConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZIOPCoprocessorConfigurationClass) Alloc() VZIOPCoprocessorConfiguration {
	rv := objc.SendIfResponds[VZIOPCoprocessorConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZIOPCoprocessorConfiguration.BootLoader]
//   - [VZIOPCoprocessorConfiguration.InitWithBootLoader]
type VZIOPCoprocessorConfiguration struct {
	VZCoprocessorConfiguration
}

// VZIOPCoprocessorConfigurationFromID constructs a [VZIOPCoprocessorConfiguration] from an objc.ID.
func VZIOPCoprocessorConfigurationFromID(id objc.ID) VZIOPCoprocessorConfiguration {
	return VZIOPCoprocessorConfiguration{VZCoprocessorConfiguration: VZCoprocessorConfigurationFromID(id)}
}

// Ensure VZIOPCoprocessorConfiguration implements IVZIOPCoprocessorConfiguration.
var _ IVZIOPCoprocessorConfiguration = VZIOPCoprocessorConfiguration{}

// An interface definition for the [VZIOPCoprocessorConfiguration] class.
//
// # Methods
//
//   - [IVZIOPCoprocessorConfiguration.BootLoader]
//   - [IVZIOPCoprocessorConfiguration.InitWithBootLoader]
type IVZIOPCoprocessorConfiguration interface {
	IVZCoprocessorConfiguration

	// Topic: Methods

	BootLoader() IVZBinaryBootLoader
	InitWithBootLoader(loader objectivec.IObject) VZIOPCoprocessorConfiguration
}

// Init initializes the instance.
func (v VZIOPCoprocessorConfiguration) Init() VZIOPCoprocessorConfiguration {
	rv := objc.SendIfResponds[VZIOPCoprocessorConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZIOPCoprocessorConfiguration) Autorelease() VZIOPCoprocessorConfiguration {
	rv := objc.SendIfResponds[VZIOPCoprocessorConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZIOPCoprocessorConfiguration creates a new VZIOPCoprocessorConfiguration instance.
func NewVZIOPCoprocessorConfiguration() VZIOPCoprocessorConfiguration {
	class := getVZIOPCoprocessorConfigurationClass()
	rv := objc.SendIfResponds[VZIOPCoprocessorConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZIOPCoprocessorConfigurationWithBootLoader(loader objectivec.IObject) VZIOPCoprocessorConfiguration {
	instance := getVZIOPCoprocessorConfigurationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBootLoader:"), loader)
	return VZIOPCoprocessorConfigurationFromID(rv)
}

func (v VZIOPCoprocessorConfiguration) InitWithBootLoader(loader objectivec.IObject) VZIOPCoprocessorConfiguration {
	rv := objc.SendIfResponds[VZIOPCoprocessorConfiguration](v.ID, objc.Sel("initWithBootLoader:"), loader)
	return rv
}

func (v VZIOPCoprocessorConfiguration) BootLoader() IVZBinaryBootLoader {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("bootLoader"))
	return VZBinaryBootLoaderFromID(objc.ID(rv))
}

// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioSharedMemoryRegionConfiguration] class.
var (
	_VZVirtioSharedMemoryRegionConfigurationClass     VZVirtioSharedMemoryRegionConfigurationClass
	_VZVirtioSharedMemoryRegionConfigurationClassOnce sync.Once
)

func getVZVirtioSharedMemoryRegionConfigurationClass() VZVirtioSharedMemoryRegionConfigurationClass {
	_VZVirtioSharedMemoryRegionConfigurationClassOnce.Do(func() {
		_VZVirtioSharedMemoryRegionConfigurationClass = VZVirtioSharedMemoryRegionConfigurationClass{class: objc.GetClass("VZVirtioSharedMemoryRegionConfiguration")}
	})
	return _VZVirtioSharedMemoryRegionConfigurationClass
}

// GetVZVirtioSharedMemoryRegionConfigurationClass returns the class object for VZVirtioSharedMemoryRegionConfiguration.
func GetVZVirtioSharedMemoryRegionConfigurationClass() VZVirtioSharedMemoryRegionConfigurationClass {
	return getVZVirtioSharedMemoryRegionConfigurationClass()
}

type VZVirtioSharedMemoryRegionConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioSharedMemoryRegionConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioSharedMemoryRegionConfigurationClass) Alloc() VZVirtioSharedMemoryRegionConfiguration {
	rv := objc.SendIfResponds[VZVirtioSharedMemoryRegionConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVirtioSharedMemoryRegionConfiguration.RegionID]
//   - [VZVirtioSharedMemoryRegionConfiguration.Size]
//   - [VZVirtioSharedMemoryRegionConfiguration.InitWithRegionIDSize]
type VZVirtioSharedMemoryRegionConfiguration struct {
	objectivec.Object
}

// VZVirtioSharedMemoryRegionConfigurationFromID constructs a [VZVirtioSharedMemoryRegionConfiguration] from an objc.ID.
func VZVirtioSharedMemoryRegionConfigurationFromID(id objc.ID) VZVirtioSharedMemoryRegionConfiguration {
	return VZVirtioSharedMemoryRegionConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZVirtioSharedMemoryRegionConfiguration implements IVZVirtioSharedMemoryRegionConfiguration.
var _ IVZVirtioSharedMemoryRegionConfiguration = VZVirtioSharedMemoryRegionConfiguration{}

// An interface definition for the [VZVirtioSharedMemoryRegionConfiguration] class.
//
// # Methods
//
//   - [IVZVirtioSharedMemoryRegionConfiguration.RegionID]
//   - [IVZVirtioSharedMemoryRegionConfiguration.Size]
//   - [IVZVirtioSharedMemoryRegionConfiguration.InitWithRegionIDSize]
type IVZVirtioSharedMemoryRegionConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	RegionID() byte
	Size() uint64
	InitWithRegionIDSize(id byte, size uint64) VZVirtioSharedMemoryRegionConfiguration
}

// Init initializes the instance.
func (v VZVirtioSharedMemoryRegionConfiguration) Init() VZVirtioSharedMemoryRegionConfiguration {
	rv := objc.SendIfResponds[VZVirtioSharedMemoryRegionConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioSharedMemoryRegionConfiguration) Autorelease() VZVirtioSharedMemoryRegionConfiguration {
	rv := objc.SendIfResponds[VZVirtioSharedMemoryRegionConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSharedMemoryRegionConfiguration creates a new VZVirtioSharedMemoryRegionConfiguration instance.
func NewVZVirtioSharedMemoryRegionConfiguration() VZVirtioSharedMemoryRegionConfiguration {
	class := getVZVirtioSharedMemoryRegionConfigurationClass()
	rv := objc.SendIfResponds[VZVirtioSharedMemoryRegionConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZVirtioSharedMemoryRegionConfigurationWithRegionIDSize(id byte, size uint64) VZVirtioSharedMemoryRegionConfiguration {
	instance := getVZVirtioSharedMemoryRegionConfigurationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithRegionID:size:"), id, size)
	return VZVirtioSharedMemoryRegionConfigurationFromID(rv)
}

func (v VZVirtioSharedMemoryRegionConfiguration) InitWithRegionIDSize(id byte, size uint64) VZVirtioSharedMemoryRegionConfiguration {
	rv := objc.SendIfResponds[VZVirtioSharedMemoryRegionConfiguration](v.ID, objc.Sel("initWithRegionID:size:"), id, size)
	return rv
}

func (v VZVirtioSharedMemoryRegionConfiguration) RegionID() byte {
	rv := objc.SendIfResponds[byte](v.ID, objc.Sel("regionID"))
	return rv
}
func (v VZVirtioSharedMemoryRegionConfiguration) Size() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("size"))
	return rv
}

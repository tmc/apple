// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZLinearFramebufferGraphicsDeviceConfiguration] class.
var (
	_VZLinearFramebufferGraphicsDeviceConfigurationClass     VZLinearFramebufferGraphicsDeviceConfigurationClass
	_VZLinearFramebufferGraphicsDeviceConfigurationClassOnce sync.Once
)

func getVZLinearFramebufferGraphicsDeviceConfigurationClass() VZLinearFramebufferGraphicsDeviceConfigurationClass {
	_VZLinearFramebufferGraphicsDeviceConfigurationClassOnce.Do(func() {
		_VZLinearFramebufferGraphicsDeviceConfigurationClass = VZLinearFramebufferGraphicsDeviceConfigurationClass{class: objc.GetClass("_VZLinearFramebufferGraphicsDeviceConfiguration")}
	})
	return _VZLinearFramebufferGraphicsDeviceConfigurationClass
}

// GetVZLinearFramebufferGraphicsDeviceConfigurationClass returns the class object for _VZLinearFramebufferGraphicsDeviceConfiguration.
func GetVZLinearFramebufferGraphicsDeviceConfigurationClass() VZLinearFramebufferGraphicsDeviceConfigurationClass {
	return getVZLinearFramebufferGraphicsDeviceConfigurationClass()
}

type VZLinearFramebufferGraphicsDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZLinearFramebufferGraphicsDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZLinearFramebufferGraphicsDeviceConfigurationClass) Alloc() VZLinearFramebufferGraphicsDeviceConfiguration {
	rv := objc.Send[VZLinearFramebufferGraphicsDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZLinearFramebufferGraphicsDeviceConfiguration.BackingStoreSize]
//   - [VZLinearFramebufferGraphicsDeviceConfiguration.SetBackingStoreSize]
//   - [VZLinearFramebufferGraphicsDeviceConfiguration.InitWithBackingStoreSize]
type VZLinearFramebufferGraphicsDeviceConfiguration struct {
	VZGraphicsDeviceConfiguration
}

// VZLinearFramebufferGraphicsDeviceConfigurationFromID constructs a [VZLinearFramebufferGraphicsDeviceConfiguration] from an objc.ID.
func VZLinearFramebufferGraphicsDeviceConfigurationFromID(id objc.ID) VZLinearFramebufferGraphicsDeviceConfiguration {
	return VZLinearFramebufferGraphicsDeviceConfiguration{VZGraphicsDeviceConfiguration: VZGraphicsDeviceConfigurationFromID(id)}
}

// Ensure VZLinearFramebufferGraphicsDeviceConfiguration implements IVZLinearFramebufferGraphicsDeviceConfiguration.
var _ IVZLinearFramebufferGraphicsDeviceConfiguration = VZLinearFramebufferGraphicsDeviceConfiguration{}

// An interface definition for the [VZLinearFramebufferGraphicsDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZLinearFramebufferGraphicsDeviceConfiguration.BackingStoreSize]
//   - [IVZLinearFramebufferGraphicsDeviceConfiguration.SetBackingStoreSize]
//   - [IVZLinearFramebufferGraphicsDeviceConfiguration.InitWithBackingStoreSize]
type IVZLinearFramebufferGraphicsDeviceConfiguration interface {
	IVZGraphicsDeviceConfiguration

	// Topic: Methods

	BackingStoreSize() corefoundation.CGSize
	SetBackingStoreSize(value corefoundation.CGSize)
	InitWithBackingStoreSize(size corefoundation.CGSize) VZLinearFramebufferGraphicsDeviceConfiguration
}

// Init initializes the instance.
func (v VZLinearFramebufferGraphicsDeviceConfiguration) Init() VZLinearFramebufferGraphicsDeviceConfiguration {
	rv := objc.Send[VZLinearFramebufferGraphicsDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZLinearFramebufferGraphicsDeviceConfiguration) Autorelease() VZLinearFramebufferGraphicsDeviceConfiguration {
	rv := objc.Send[VZLinearFramebufferGraphicsDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinearFramebufferGraphicsDeviceConfiguration creates a new VZLinearFramebufferGraphicsDeviceConfiguration instance.
func NewVZLinearFramebufferGraphicsDeviceConfiguration() VZLinearFramebufferGraphicsDeviceConfiguration {
	class := getVZLinearFramebufferGraphicsDeviceConfigurationClass()
	rv := objc.Send[VZLinearFramebufferGraphicsDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZLinearFramebufferGraphicsDeviceConfigurationWithBackingStoreSize(size corefoundation.CGSize) VZLinearFramebufferGraphicsDeviceConfiguration {
	instance := getVZLinearFramebufferGraphicsDeviceConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackingStoreSize:"), size)
	return VZLinearFramebufferGraphicsDeviceConfigurationFromID(rv)
}

func (v VZLinearFramebufferGraphicsDeviceConfiguration) InitWithBackingStoreSize(size corefoundation.CGSize) VZLinearFramebufferGraphicsDeviceConfiguration {
	rv := objc.Send[VZLinearFramebufferGraphicsDeviceConfiguration](v.ID, objc.Sel("initWithBackingStoreSize:"), size)
	return rv
}

func (v VZLinearFramebufferGraphicsDeviceConfiguration) BackingStoreSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("backingStoreSize"))
	return corefoundation.CGSize(rv)
}
func (v VZLinearFramebufferGraphicsDeviceConfiguration) SetBackingStoreSize(value corefoundation.CGSize) {
	objc.Send[struct{}](v.ID, objc.Sel("setBackingStoreSize:"), value)
}

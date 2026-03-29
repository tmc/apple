// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
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

//
// # Methods
//
//   - [VZLinearFramebufferGraphicsDeviceConfiguration.BackingStoreSize]
//   - [VZLinearFramebufferGraphicsDeviceConfiguration.SetBackingStoreSize]
//   - [VZLinearFramebufferGraphicsDeviceConfiguration.EncodeWithEncoder]
//   - [VZLinearFramebufferGraphicsDeviceConfiguration.InitWithBackingStoreSize]
// See: https://developer.apple.com/documentation/Virtualization/_VZLinearFramebufferGraphicsDeviceConfiguration
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
//   - [IVZLinearFramebufferGraphicsDeviceConfiguration.EncodeWithEncoder]
//   - [IVZLinearFramebufferGraphicsDeviceConfiguration.InitWithBackingStoreSize]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZLinearFramebufferGraphicsDeviceConfiguration
type IVZLinearFramebufferGraphicsDeviceConfiguration interface {
	IVZGraphicsDeviceConfiguration

	// Topic: Methods

	BackingStoreSize() corefoundation.CGSize
	SetBackingStoreSize(value corefoundation.CGSize)
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
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

//
// See: https://developer.apple.com/documentation/Virtualization/_VZLinearFramebufferGraphicsDeviceConfiguration/initWithBackingStoreSize:
func NewVZLinearFramebufferGraphicsDeviceConfigurationWithBackingStoreSize(size corefoundation.CGSize) VZLinearFramebufferGraphicsDeviceConfiguration {
	instance := getVZLinearFramebufferGraphicsDeviceConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackingStoreSize:"), size)
	return VZLinearFramebufferGraphicsDeviceConfigurationFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZLinearFramebufferGraphicsDeviceConfiguration/encodeWithEncoder:
func (v VZLinearFramebufferGraphicsDeviceConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZLinearFramebufferGraphicsDeviceConfiguration/initWithBackingStoreSize:
func (v VZLinearFramebufferGraphicsDeviceConfiguration) InitWithBackingStoreSize(size corefoundation.CGSize) VZLinearFramebufferGraphicsDeviceConfiguration {
	rv := objc.Send[VZLinearFramebufferGraphicsDeviceConfiguration](v.ID, objc.Sel("initWithBackingStoreSize:"), size)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZLinearFramebufferGraphicsDeviceConfiguration/backingStoreSize
func (v VZLinearFramebufferGraphicsDeviceConfiguration) BackingStoreSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("backingStoreSize"))
	return corefoundation.CGSize(rv)
}
func (v VZLinearFramebufferGraphicsDeviceConfiguration) SetBackingStoreSize(value corefoundation.CGSize) {
	objc.Send[struct{}](v.ID, objc.Sel("setBackingStoreSize:"), value)
}


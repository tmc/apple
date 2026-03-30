// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZForwardingDebugStubConfiguration] class.
var (
	_VZForwardingDebugStubConfigurationClass     VZForwardingDebugStubConfigurationClass
	_VZForwardingDebugStubConfigurationClassOnce sync.Once
)

func getVZForwardingDebugStubConfigurationClass() VZForwardingDebugStubConfigurationClass {
	_VZForwardingDebugStubConfigurationClassOnce.Do(func() {
		_VZForwardingDebugStubConfigurationClass = VZForwardingDebugStubConfigurationClass{class: objc.GetClass("_VZForwardingDebugStubConfiguration")}
	})
	return _VZForwardingDebugStubConfigurationClass
}

// GetVZForwardingDebugStubConfigurationClass returns the class object for _VZForwardingDebugStubConfiguration.
func GetVZForwardingDebugStubConfigurationClass() VZForwardingDebugStubConfigurationClass {
	return getVZForwardingDebugStubConfigurationClass()
}

type VZForwardingDebugStubConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZForwardingDebugStubConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZForwardingDebugStubConfigurationClass) Alloc() VZForwardingDebugStubConfiguration {
	rv := objc.Send[VZForwardingDebugStubConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZForwardingDebugStubConfiguration._initWithDebugStub]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZForwardingDebugStubConfiguration
type VZForwardingDebugStubConfiguration struct {
	VZDebugStubConfiguration
}

// VZForwardingDebugStubConfigurationFromID constructs a [VZForwardingDebugStubConfiguration] from an objc.ID.
func VZForwardingDebugStubConfigurationFromID(id objc.ID) VZForwardingDebugStubConfiguration {
	return VZForwardingDebugStubConfiguration{VZDebugStubConfiguration: VZDebugStubConfigurationFromID(id)}
}

// Ensure VZForwardingDebugStubConfiguration implements IVZForwardingDebugStubConfiguration.
var _ IVZForwardingDebugStubConfiguration = VZForwardingDebugStubConfiguration{}

// An interface definition for the [VZForwardingDebugStubConfiguration] class.
//
// # Methods
//
//   - [IVZForwardingDebugStubConfiguration._initWithDebugStub]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZForwardingDebugStubConfiguration
type IVZForwardingDebugStubConfiguration interface {
	IVZDebugStubConfiguration

	// Topic: Methods

	_initWithDebugStub(stub unsafe.Pointer) objectivec.IObject
}

// Init initializes the instance.
func (v VZForwardingDebugStubConfiguration) Init() VZForwardingDebugStubConfiguration {
	rv := objc.Send[VZForwardingDebugStubConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZForwardingDebugStubConfiguration) Autorelease() VZForwardingDebugStubConfiguration {
	rv := objc.Send[VZForwardingDebugStubConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZForwardingDebugStubConfiguration creates a new VZForwardingDebugStubConfiguration instance.
func NewVZForwardingDebugStubConfiguration() VZForwardingDebugStubConfiguration {
	class := getVZForwardingDebugStubConfigurationClass()
	rv := objc.Send[VZForwardingDebugStubConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZForwardingDebugStubConfiguration/_initWithDebugStub:
func (v VZForwardingDebugStubConfiguration) _initWithDebugStub(stub unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_initWithDebugStub:"), stub)
	return objectivec.Object{ID: rv}
}

// InitWithDebugStub is an exported wrapper for the private method _initWithDebugStub.
func (v VZForwardingDebugStubConfiguration) InitWithDebugStub(stub unsafe.Pointer) objectivec.IObject {
	return v._initWithDebugStub(stub)
}

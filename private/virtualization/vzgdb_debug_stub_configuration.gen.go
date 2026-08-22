// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZGDBDebugStubConfiguration] class.
var (
	_VZGDBDebugStubConfigurationClass     VZGDBDebugStubConfigurationClass
	_VZGDBDebugStubConfigurationClassOnce sync.Once
)

func getVZGDBDebugStubConfigurationClass() VZGDBDebugStubConfigurationClass {
	_VZGDBDebugStubConfigurationClassOnce.Do(func() {
		_VZGDBDebugStubConfigurationClass = VZGDBDebugStubConfigurationClass{class: objc.GetClass("_VZGDBDebugStubConfiguration")}
	})
	return _VZGDBDebugStubConfigurationClass
}

// GetVZGDBDebugStubConfigurationClass returns the class object for _VZGDBDebugStubConfiguration.
func GetVZGDBDebugStubConfigurationClass() VZGDBDebugStubConfigurationClass {
	return getVZGDBDebugStubConfigurationClass()
}

type VZGDBDebugStubConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZGDBDebugStubConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZGDBDebugStubConfigurationClass) Alloc() VZGDBDebugStubConfiguration {
	rv := objc.SendIfResponds[VZGDBDebugStubConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZGDBDebugStubConfiguration.ListensOnAllNetworkInterfaces]
//   - [VZGDBDebugStubConfiguration.SetListensOnAllNetworkInterfaces]
//   - [VZGDBDebugStubConfiguration.Port]
//   - [VZGDBDebugStubConfiguration.SetPort]
//   - [VZGDBDebugStubConfiguration.InitWithPort]
type VZGDBDebugStubConfiguration struct {
	VZDebugStubConfiguration
}

// VZGDBDebugStubConfigurationFromID constructs a [VZGDBDebugStubConfiguration] from an objc.ID.
func VZGDBDebugStubConfigurationFromID(id objc.ID) VZGDBDebugStubConfiguration {
	return VZGDBDebugStubConfiguration{VZDebugStubConfiguration: VZDebugStubConfigurationFromID(id)}
}

// Ensure VZGDBDebugStubConfiguration implements IVZGDBDebugStubConfiguration.
var _ IVZGDBDebugStubConfiguration = VZGDBDebugStubConfiguration{}

// An interface definition for the [VZGDBDebugStubConfiguration] class.
//
// # Methods
//
//   - [IVZGDBDebugStubConfiguration.ListensOnAllNetworkInterfaces]
//   - [IVZGDBDebugStubConfiguration.SetListensOnAllNetworkInterfaces]
//   - [IVZGDBDebugStubConfiguration.Port]
//   - [IVZGDBDebugStubConfiguration.SetPort]
//   - [IVZGDBDebugStubConfiguration.InitWithPort]
type IVZGDBDebugStubConfiguration interface {
	IVZDebugStubConfiguration

	// Topic: Methods

	ListensOnAllNetworkInterfaces() bool
	SetListensOnAllNetworkInterfaces(value bool)
	Port() uint16
	SetPort(value uint16)
	InitWithPort(port uint16) VZGDBDebugStubConfiguration
}

// Init initializes the instance.
func (v VZGDBDebugStubConfiguration) Init() VZGDBDebugStubConfiguration {
	rv := objc.SendIfResponds[VZGDBDebugStubConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZGDBDebugStubConfiguration) Autorelease() VZGDBDebugStubConfiguration {
	rv := objc.SendIfResponds[VZGDBDebugStubConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGDBDebugStubConfiguration creates a new VZGDBDebugStubConfiguration instance.
func NewVZGDBDebugStubConfiguration() VZGDBDebugStubConfiguration {
	class := getVZGDBDebugStubConfigurationClass()
	rv := objc.SendIfResponds[VZGDBDebugStubConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZGDBDebugStubConfigurationWithPort(port uint16) VZGDBDebugStubConfiguration {
	instance := getVZGDBDebugStubConfigurationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithPort:"), port)
	return VZGDBDebugStubConfigurationFromID(rv)
}

func (v VZGDBDebugStubConfiguration) InitWithPort(port uint16) VZGDBDebugStubConfiguration {
	rv := objc.SendIfResponds[VZGDBDebugStubConfiguration](v.ID, objc.Sel("initWithPort:"), port)
	return rv
}

func (v VZGDBDebugStubConfiguration) ListensOnAllNetworkInterfaces() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("listensOnAllNetworkInterfaces"))
	return rv
}
func (v VZGDBDebugStubConfiguration) SetListensOnAllNetworkInterfaces(value bool) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setListensOnAllNetworkInterfaces:"), value)
}
func (v VZGDBDebugStubConfiguration) Port() uint16 {
	rv := objc.SendIfResponds[uint16](v.ID, objc.Sel("port"))
	return rv
}
func (v VZGDBDebugStubConfiguration) SetPort(value uint16) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setPort:"), value)
}

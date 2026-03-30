// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZCustomVirtioDevicePluginProvider] class.
var (
	_VZCustomVirtioDevicePluginProviderClass     VZCustomVirtioDevicePluginProviderClass
	_VZCustomVirtioDevicePluginProviderClassOnce sync.Once
)

func getVZCustomVirtioDevicePluginProviderClass() VZCustomVirtioDevicePluginProviderClass {
	_VZCustomVirtioDevicePluginProviderClassOnce.Do(func() {
		_VZCustomVirtioDevicePluginProviderClass = VZCustomVirtioDevicePluginProviderClass{class: objc.GetClass("_VZCustomVirtioDevicePluginProvider")}
	})
	return _VZCustomVirtioDevicePluginProviderClass
}

// GetVZCustomVirtioDevicePluginProviderClass returns the class object for _VZCustomVirtioDevicePluginProvider.
func GetVZCustomVirtioDevicePluginProviderClass() VZCustomVirtioDevicePluginProviderClass {
	return getVZCustomVirtioDevicePluginProviderClass()
}

type VZCustomVirtioDevicePluginProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCustomVirtioDevicePluginProviderClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCustomVirtioDevicePluginProviderClass) Alloc() VZCustomVirtioDevicePluginProvider {
	rv := objc.Send[VZCustomVirtioDevicePluginProvider](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZCustomVirtioDevicePluginProvider.PluginName]
//   - [VZCustomVirtioDevicePluginProvider.PluginPersonality]
//   - [VZCustomVirtioDevicePluginProvider.InitWithPluginNamePluginPersonality]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDevicePluginProvider
type VZCustomVirtioDevicePluginProvider struct {
	VZCustomVirtioDeviceProvider
}

// VZCustomVirtioDevicePluginProviderFromID constructs a [VZCustomVirtioDevicePluginProvider] from an objc.ID.
func VZCustomVirtioDevicePluginProviderFromID(id objc.ID) VZCustomVirtioDevicePluginProvider {
	return VZCustomVirtioDevicePluginProvider{VZCustomVirtioDeviceProvider: VZCustomVirtioDeviceProviderFromID(id)}
}

// Ensure VZCustomVirtioDevicePluginProvider implements IVZCustomVirtioDevicePluginProvider.
var _ IVZCustomVirtioDevicePluginProvider = VZCustomVirtioDevicePluginProvider{}

// An interface definition for the [VZCustomVirtioDevicePluginProvider] class.
//
// # Methods
//
//   - [IVZCustomVirtioDevicePluginProvider.PluginName]
//   - [IVZCustomVirtioDevicePluginProvider.PluginPersonality]
//   - [IVZCustomVirtioDevicePluginProvider.InitWithPluginNamePluginPersonality]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDevicePluginProvider
type IVZCustomVirtioDevicePluginProvider interface {
	IVZCustomVirtioDeviceProvider

	// Topic: Methods

	PluginName() string
	PluginPersonality() string
	InitWithPluginNamePluginPersonality(name objectivec.IObject, personality objectivec.IObject) VZCustomVirtioDevicePluginProvider
}

// Init initializes the instance.
func (v VZCustomVirtioDevicePluginProvider) Init() VZCustomVirtioDevicePluginProvider {
	rv := objc.Send[VZCustomVirtioDevicePluginProvider](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCustomVirtioDevicePluginProvider) Autorelease() VZCustomVirtioDevicePluginProvider {
	rv := objc.Send[VZCustomVirtioDevicePluginProvider](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCustomVirtioDevicePluginProvider creates a new VZCustomVirtioDevicePluginProvider instance.
func NewVZCustomVirtioDevicePluginProvider() VZCustomVirtioDevicePluginProvider {
	class := getVZCustomVirtioDevicePluginProviderClass()
	rv := objc.Send[VZCustomVirtioDevicePluginProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDevicePluginProvider/initWithPluginName:pluginPersonality:
func NewVZCustomVirtioDevicePluginProviderWithPluginNamePluginPersonality(name objectivec.IObject, personality objectivec.IObject) VZCustomVirtioDevicePluginProvider {
	instance := getVZCustomVirtioDevicePluginProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPluginName:pluginPersonality:"), name, personality)
	return VZCustomVirtioDevicePluginProviderFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDevicePluginProvider/initWithPluginName:pluginPersonality:
func (v VZCustomVirtioDevicePluginProvider) InitWithPluginNamePluginPersonality(name objectivec.IObject, personality objectivec.IObject) VZCustomVirtioDevicePluginProvider {
	rv := objc.Send[VZCustomVirtioDevicePluginProvider](v.ID, objc.Sel("initWithPluginName:pluginPersonality:"), name, personality)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDevicePluginProvider/pluginName
func (v VZCustomVirtioDevicePluginProvider) PluginName() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("pluginName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDevicePluginProvider/pluginPersonality
func (v VZCustomVirtioDevicePluginProvider) PluginPersonality() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("pluginPersonality"))
	return foundation.NSStringFromID(rv).String()
}

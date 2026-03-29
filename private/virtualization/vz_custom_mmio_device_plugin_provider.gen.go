// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZCustomMMIODevicePluginProvider] class.
var (
	_VZCustomMMIODevicePluginProviderClass     VZCustomMMIODevicePluginProviderClass
	_VZCustomMMIODevicePluginProviderClassOnce sync.Once
)

func getVZCustomMMIODevicePluginProviderClass() VZCustomMMIODevicePluginProviderClass {
	_VZCustomMMIODevicePluginProviderClassOnce.Do(func() {
		_VZCustomMMIODevicePluginProviderClass = VZCustomMMIODevicePluginProviderClass{class: objc.GetClass("_VZCustomMMIODevicePluginProvider")}
	})
	return _VZCustomMMIODevicePluginProviderClass
}

// GetVZCustomMMIODevicePluginProviderClass returns the class object for _VZCustomMMIODevicePluginProvider.
func GetVZCustomMMIODevicePluginProviderClass() VZCustomMMIODevicePluginProviderClass {
	return getVZCustomMMIODevicePluginProviderClass()
}

type VZCustomMMIODevicePluginProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCustomMMIODevicePluginProviderClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCustomMMIODevicePluginProviderClass) Alloc() VZCustomMMIODevicePluginProvider {
	rv := objc.Send[VZCustomMMIODevicePluginProvider](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZCustomMMIODevicePluginProvider.PluginName]
//   - [VZCustomMMIODevicePluginProvider.PluginPersonality]
//   - [VZCustomMMIODevicePluginProvider.InitWithPluginNamePluginPersonality]
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODevicePluginProvider
type VZCustomMMIODevicePluginProvider struct {
	VZCustomMMIODeviceProvider
}

// VZCustomMMIODevicePluginProviderFromID constructs a [VZCustomMMIODevicePluginProvider] from an objc.ID.
func VZCustomMMIODevicePluginProviderFromID(id objc.ID) VZCustomMMIODevicePluginProvider {
	return VZCustomMMIODevicePluginProvider{VZCustomMMIODeviceProvider: VZCustomMMIODeviceProviderFromID(id)}
}
// Ensure VZCustomMMIODevicePluginProvider implements IVZCustomMMIODevicePluginProvider.
var _ IVZCustomMMIODevicePluginProvider = VZCustomMMIODevicePluginProvider{}

// An interface definition for the [VZCustomMMIODevicePluginProvider] class.
//
// # Methods
//
//   - [IVZCustomMMIODevicePluginProvider.PluginName]
//   - [IVZCustomMMIODevicePluginProvider.PluginPersonality]
//   - [IVZCustomMMIODevicePluginProvider.InitWithPluginNamePluginPersonality]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODevicePluginProvider
type IVZCustomMMIODevicePluginProvider interface {
	IVZCustomMMIODeviceProvider

	// Topic: Methods

	PluginName() string
	PluginPersonality() string
	InitWithPluginNamePluginPersonality(name objectivec.IObject, personality objectivec.IObject) VZCustomMMIODevicePluginProvider
}

// Init initializes the instance.
func (v VZCustomMMIODevicePluginProvider) Init() VZCustomMMIODevicePluginProvider {
	rv := objc.Send[VZCustomMMIODevicePluginProvider](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCustomMMIODevicePluginProvider) Autorelease() VZCustomMMIODevicePluginProvider {
	rv := objc.Send[VZCustomMMIODevicePluginProvider](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCustomMMIODevicePluginProvider creates a new VZCustomMMIODevicePluginProvider instance.
func NewVZCustomMMIODevicePluginProvider() VZCustomMMIODevicePluginProvider {
	class := getVZCustomMMIODevicePluginProviderClass()
	rv := objc.Send[VZCustomMMIODevicePluginProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODevicePluginProvider/initWithPluginName:pluginPersonality:
func NewVZCustomMMIODevicePluginProviderWithPluginNamePluginPersonality(name objectivec.IObject, personality objectivec.IObject) VZCustomMMIODevicePluginProvider {
	instance := getVZCustomMMIODevicePluginProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPluginName:pluginPersonality:"), name, personality)
	return VZCustomMMIODevicePluginProviderFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODevicePluginProvider/initWithPluginName:pluginPersonality:
func (v VZCustomMMIODevicePluginProvider) InitWithPluginNamePluginPersonality(name objectivec.IObject, personality objectivec.IObject) VZCustomMMIODevicePluginProvider {
	rv := objc.Send[VZCustomMMIODevicePluginProvider](v.ID, objc.Sel("initWithPluginName:pluginPersonality:"), name, personality)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODevicePluginProvider/pluginName
func (v VZCustomMMIODevicePluginProvider) PluginName() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("pluginName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODevicePluginProvider/pluginPersonality
func (v VZCustomMMIODevicePluginProvider) PluginPersonality() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("pluginPersonality"))
	return foundation.NSStringFromID(rv).String()
}


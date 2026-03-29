// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZCustomVirtioDevicePluginBridge] class.
var (
	_VZCustomVirtioDevicePluginBridgeClass     VZCustomVirtioDevicePluginBridgeClass
	_VZCustomVirtioDevicePluginBridgeClassOnce sync.Once
)

func getVZCustomVirtioDevicePluginBridgeClass() VZCustomVirtioDevicePluginBridgeClass {
	_VZCustomVirtioDevicePluginBridgeClassOnce.Do(func() {
		_VZCustomVirtioDevicePluginBridgeClass = VZCustomVirtioDevicePluginBridgeClass{class: objc.GetClass("_VZCustomVirtioDevicePluginBridge")}
	})
	return _VZCustomVirtioDevicePluginBridgeClass
}

// GetVZCustomVirtioDevicePluginBridgeClass returns the class object for _VZCustomVirtioDevicePluginBridge.
func GetVZCustomVirtioDevicePluginBridgeClass() VZCustomVirtioDevicePluginBridgeClass {
	return getVZCustomVirtioDevicePluginBridgeClass()
}

type VZCustomVirtioDevicePluginBridgeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCustomVirtioDevicePluginBridgeClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCustomVirtioDevicePluginBridgeClass) Alloc() VZCustomVirtioDevicePluginBridge {
	rv := objc.Send[VZCustomVirtioDevicePluginBridge](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZCustomVirtioDevicePluginBridge.InitWithXPCConnectionPersonalityDictionaryPersonalityClassConnectionDelegateOptions]
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDevicePluginBridge
type VZCustomVirtioDevicePluginBridge struct {
	objectivec.Object
}

// VZCustomVirtioDevicePluginBridgeFromID constructs a [VZCustomVirtioDevicePluginBridge] from an objc.ID.
func VZCustomVirtioDevicePluginBridgeFromID(id objc.ID) VZCustomVirtioDevicePluginBridge {
	return VZCustomVirtioDevicePluginBridge{objectivec.Object{ID: id}}
}
// Ensure VZCustomVirtioDevicePluginBridge implements IVZCustomVirtioDevicePluginBridge.
var _ IVZCustomVirtioDevicePluginBridge = VZCustomVirtioDevicePluginBridge{}

// An interface definition for the [VZCustomVirtioDevicePluginBridge] class.
//
// # Methods
//
//   - [IVZCustomVirtioDevicePluginBridge.InitWithXPCConnectionPersonalityDictionaryPersonalityClassConnectionDelegateOptions]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDevicePluginBridge
type IVZCustomVirtioDevicePluginBridge interface {
	objectivec.IObject

	// Topic: Methods

	InitWithXPCConnectionPersonalityDictionaryPersonalityClassConnectionDelegateOptions(xPCConnection objectivec.IObject, dictionary objectivec.IObject, class objc.Class, delegate objectivec.IObject, options objectivec.IObject) VZCustomVirtioDevicePluginBridge
}

// Init initializes the instance.
func (v VZCustomVirtioDevicePluginBridge) Init() VZCustomVirtioDevicePluginBridge {
	rv := objc.Send[VZCustomVirtioDevicePluginBridge](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCustomVirtioDevicePluginBridge) Autorelease() VZCustomVirtioDevicePluginBridge {
	rv := objc.Send[VZCustomVirtioDevicePluginBridge](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCustomVirtioDevicePluginBridge creates a new VZCustomVirtioDevicePluginBridge instance.
func NewVZCustomVirtioDevicePluginBridge() VZCustomVirtioDevicePluginBridge {
	class := getVZCustomVirtioDevicePluginBridgeClass()
	rv := objc.Send[VZCustomVirtioDevicePluginBridge](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDevicePluginBridge/initWithXPCConnection:personalityDictionary:personalityClass:connectionDelegate:options:
func NewVZCustomVirtioDevicePluginBridgeWithXPCConnectionPersonalityDictionaryPersonalityClassConnectionDelegateOptions(xPCConnection objectivec.IObject, dictionary objectivec.IObject, class objc.Class, delegate objectivec.IObject, options objectivec.IObject) VZCustomVirtioDevicePluginBridge {
	instance := getVZCustomVirtioDevicePluginBridgeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithXPCConnection:personalityDictionary:personalityClass:connectionDelegate:options:"), xPCConnection, dictionary, class, delegate, options)
	return VZCustomVirtioDevicePluginBridgeFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDevicePluginBridge/initWithXPCConnection:personalityDictionary:personalityClass:connectionDelegate:options:
func (v VZCustomVirtioDevicePluginBridge) InitWithXPCConnectionPersonalityDictionaryPersonalityClassConnectionDelegateOptions(xPCConnection objectivec.IObject, dictionary objectivec.IObject, class objc.Class, delegate objectivec.IObject, options objectivec.IObject) VZCustomVirtioDevicePluginBridge {
	rv := objc.Send[VZCustomVirtioDevicePluginBridge](v.ID, objc.Sel("initWithXPCConnection:personalityDictionary:personalityClass:connectionDelegate:options:"), xPCConnection, dictionary, class, delegate, options)
	return rv
}


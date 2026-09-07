// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZCustomMMIODevicePluginBridge] class.
var (
	_VZCustomMMIODevicePluginBridgeClass     VZCustomMMIODevicePluginBridgeClass
	_VZCustomMMIODevicePluginBridgeClassOnce sync.Once
)

func getVZCustomMMIODevicePluginBridgeClass() VZCustomMMIODevicePluginBridgeClass {
	_VZCustomMMIODevicePluginBridgeClassOnce.Do(func() {
		_VZCustomMMIODevicePluginBridgeClass = VZCustomMMIODevicePluginBridgeClass{class: objc.GetClass("_VZCustomMMIODevicePluginBridge")}
	})
	return _VZCustomMMIODevicePluginBridgeClass
}

// GetVZCustomMMIODevicePluginBridgeClass returns the class object for _VZCustomMMIODevicePluginBridge.
func GetVZCustomMMIODevicePluginBridgeClass() VZCustomMMIODevicePluginBridgeClass {
	return getVZCustomMMIODevicePluginBridgeClass()
}

type VZCustomMMIODevicePluginBridgeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCustomMMIODevicePluginBridgeClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCustomMMIODevicePluginBridgeClass) Alloc() VZCustomMMIODevicePluginBridge {
	rv := objc.SendIfResponds[VZCustomMMIODevicePluginBridge](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZCustomMMIODevicePluginBridge.InitWithXPCConnectionPersonalityDictionaryPersonalityClassConnectionDelegateOptions]
type VZCustomMMIODevicePluginBridge struct {
	objectivec.Object
}

// VZCustomMMIODevicePluginBridgeFromID constructs a [VZCustomMMIODevicePluginBridge] from an objc.ID.
func VZCustomMMIODevicePluginBridgeFromID(id objc.ID) VZCustomMMIODevicePluginBridge {
	return VZCustomMMIODevicePluginBridge{objectivec.Object{ID: id}}
}

// Ensure VZCustomMMIODevicePluginBridge implements IVZCustomMMIODevicePluginBridge.
var _ IVZCustomMMIODevicePluginBridge = VZCustomMMIODevicePluginBridge{}

// An interface definition for the [VZCustomMMIODevicePluginBridge] class.
//
// # Methods
//
//   - [IVZCustomMMIODevicePluginBridge.InitWithXPCConnectionPersonalityDictionaryPersonalityClassConnectionDelegateOptions]
type IVZCustomMMIODevicePluginBridge interface {
	objectivec.IObject

	// Topic: Methods

	InitWithXPCConnectionPersonalityDictionaryPersonalityClassConnectionDelegateOptions(xPCConnection objectivec.IObject, dictionary objectivec.IObject, class objectivec.Class, delegate objectivec.IObject, options objectivec.IObject) VZCustomMMIODevicePluginBridge
}

// Init initializes the instance.
func (v VZCustomMMIODevicePluginBridge) Init() VZCustomMMIODevicePluginBridge {
	rv := objc.SendIfResponds[VZCustomMMIODevicePluginBridge](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCustomMMIODevicePluginBridge) Autorelease() VZCustomMMIODevicePluginBridge {
	rv := objc.SendIfResponds[VZCustomMMIODevicePluginBridge](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCustomMMIODevicePluginBridge creates a new VZCustomMMIODevicePluginBridge instance.
func NewVZCustomMMIODevicePluginBridge() VZCustomMMIODevicePluginBridge {
	class := getVZCustomMMIODevicePluginBridgeClass()
	rv := objc.SendIfResponds[VZCustomMMIODevicePluginBridge](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZCustomMMIODevicePluginBridgeWithXPCConnectionPersonalityDictionaryPersonalityClassConnectionDelegateOptions(xPCConnection objectivec.IObject, dictionary objectivec.IObject, class objectivec.Class, delegate objectivec.IObject, options objectivec.IObject) VZCustomMMIODevicePluginBridge {
	instance := getVZCustomMMIODevicePluginBridgeClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithXPCConnection:personalityDictionary:personalityClass:connectionDelegate:options:"), xPCConnection, dictionary, class, delegate, options)
	return VZCustomMMIODevicePluginBridgeFromID(rv)
}

func (v VZCustomMMIODevicePluginBridge) InitWithXPCConnectionPersonalityDictionaryPersonalityClassConnectionDelegateOptions(xPCConnection objectivec.IObject, dictionary objectivec.IObject, class objectivec.Class, delegate objectivec.IObject, options objectivec.IObject) VZCustomMMIODevicePluginBridge {
	rv := objc.SendIfResponds[VZCustomMMIODevicePluginBridge](v.ID, objc.Sel("initWithXPCConnection:personalityDictionary:personalityClass:connectionDelegate:options:"), xPCConnection, dictionary, class, delegate, options)
	return rv
}

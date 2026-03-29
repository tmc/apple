// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NEAppProxyProviderManager] class.
var (
	_NEAppProxyProviderManagerClass     NEAppProxyProviderManagerClass
	_NEAppProxyProviderManagerClassOnce sync.Once
)

func getNEAppProxyProviderManagerClass() NEAppProxyProviderManagerClass {
	_NEAppProxyProviderManagerClassOnce.Do(func() {
		_NEAppProxyProviderManagerClass = NEAppProxyProviderManagerClass{class: objc.GetClass("NEAppProxyProviderManager")}
	})
	return _NEAppProxyProviderManagerClass
}

// GetNEAppProxyProviderManagerClass returns the class object for NEAppProxyProviderManager.
func GetNEAppProxyProviderManagerClass() NEAppProxyProviderManagerClass {
	return getNEAppProxyProviderManagerClass()
}

type NEAppProxyProviderManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEAppProxyProviderManagerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEAppProxyProviderManagerClass) Alloc() NEAppProxyProviderManager {
	rv := objc.Send[NEAppProxyProviderManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object to create and manage the app proxy provider’s VPN
// configuration.
//
// # Overview
// 
// Objects cannot be directly instantiated. Instead, App Proxy configurations
// are created exclusively from
// `com.AppleXCUIElementTypeVpnXCUIElementTypeManagedXCUIElementTypeApplayer()`
// payloads in configuration profiles.
// 
// App Proxy configurations can only be used with Per-App VPN routing rules.
// For more details about how to create App Proxy configurations and configure
// Per-App VPN, see [NETunnelProviderManager].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyProviderManager
type NEAppProxyProviderManager struct {
	NETunnelProviderManager
}

// NEAppProxyProviderManagerFromID constructs a [NEAppProxyProviderManager] from an objc.ID.
//
// An object to create and manage the app proxy provider’s VPN
// configuration.
func NEAppProxyProviderManagerFromID(id objc.ID) NEAppProxyProviderManager {
	return NEAppProxyProviderManager{NETunnelProviderManager: NETunnelProviderManagerFromID(id)}
}
// NOTE: NEAppProxyProviderManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEAppProxyProviderManager] class.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyProviderManager
type INEAppProxyProviderManager interface {
	INETunnelProviderManager
}

// Init initializes the instance.
func (a NEAppProxyProviderManager) Init() NEAppProxyProviderManager {
	rv := objc.Send[NEAppProxyProviderManager](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NEAppProxyProviderManager) Autorelease() NEAppProxyProviderManager {
	rv := objc.Send[NEAppProxyProviderManager](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppProxyProviderManager creates a new NEAppProxyProviderManager instance.
func NewNEAppProxyProviderManager() NEAppProxyProviderManager {
	class := getNEAppProxyProviderManagerClass()
	rv := objc.Send[NEAppProxyProviderManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}


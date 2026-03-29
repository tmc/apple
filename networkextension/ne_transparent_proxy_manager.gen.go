// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NETransparentProxyManager] class.
var (
	_NETransparentProxyManagerClass     NETransparentProxyManagerClass
	_NETransparentProxyManagerClassOnce sync.Once
)

func getNETransparentProxyManagerClass() NETransparentProxyManagerClass {
	_NETransparentProxyManagerClassOnce.Do(func() {
		_NETransparentProxyManagerClass = NETransparentProxyManagerClass{class: objc.GetClass("NETransparentProxyManager")}
	})
	return _NETransparentProxyManagerClass
}

// GetNETransparentProxyManagerClass returns the class object for NETransparentProxyManager.
func GetNETransparentProxyManagerClass() NETransparentProxyManagerClass {
	return getNETransparentProxyManagerClass()
}

type NETransparentProxyManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NETransparentProxyManagerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NETransparentProxyManagerClass) Alloc() NETransparentProxyManager {
	rv := objc.Send[NETransparentProxyManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that configures and controls transparent proxies.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyManager
type NETransparentProxyManager struct {
	NEVPNManager
}

// NETransparentProxyManagerFromID constructs a [NETransparentProxyManager] from an objc.ID.
//
// An object that configures and controls transparent proxies.
func NETransparentProxyManagerFromID(id objc.ID) NETransparentProxyManager {
	return NETransparentProxyManager{NEVPNManager: NEVPNManagerFromID(id)}
}
// NOTE: NETransparentProxyManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NETransparentProxyManager] class.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyManager
type INETransparentProxyManager interface {
	INEVPNManager
}

// Init initializes the instance.
func (t NETransparentProxyManager) Init() NETransparentProxyManager {
	rv := objc.Send[NETransparentProxyManager](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NETransparentProxyManager) Autorelease() NETransparentProxyManager {
	rv := objc.Send[NETransparentProxyManager](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETransparentProxyManager creates a new NETransparentProxyManager instance.
func NewNETransparentProxyManager() NETransparentProxyManager {
	class := getNETransparentProxyManagerClass()
	rv := objc.Send[NETransparentProxyManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}


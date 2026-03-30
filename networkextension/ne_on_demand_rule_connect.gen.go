// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [NEOnDemandRuleConnect] class.
var (
	_NEOnDemandRuleConnectClass     NEOnDemandRuleConnectClass
	_NEOnDemandRuleConnectClassOnce sync.Once
)

func getNEOnDemandRuleConnectClass() NEOnDemandRuleConnectClass {
	_NEOnDemandRuleConnectClassOnce.Do(func() {
		_NEOnDemandRuleConnectClass = NEOnDemandRuleConnectClass{class: objc.GetClass("NEOnDemandRuleConnect")}
	})
	return _NEOnDemandRuleConnectClass
}

// GetNEOnDemandRuleConnectClass returns the class object for NEOnDemandRuleConnect.
func GetNEOnDemandRuleConnectClass() NEOnDemandRuleConnectClass {
	return getNEOnDemandRuleConnectClass()
}

type NEOnDemandRuleConnectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEOnDemandRuleConnectClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEOnDemandRuleConnectClass) Alloc() NEOnDemandRuleConnect {
	rv := objc.Send[NEOnDemandRuleConnect](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A VPN On Demand rule that connects the VPN.
//
// # Overview
//
// When rules of this class match, the system starts the VPN connection
// whenever an application running on the system opens a network connection.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleConnect
type NEOnDemandRuleConnect struct {
	NEOnDemandRule
}

// NEOnDemandRuleConnectFromID constructs a [NEOnDemandRuleConnect] from an objc.ID.
//
// A VPN On Demand rule that connects the VPN.
func NEOnDemandRuleConnectFromID(id objc.ID) NEOnDemandRuleConnect {
	return NEOnDemandRuleConnect{NEOnDemandRule: NEOnDemandRuleFromID(id)}
}

// NOTE: NEOnDemandRuleConnect adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEOnDemandRuleConnect] class.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleConnect
type INEOnDemandRuleConnect interface {
	INEOnDemandRule
}

// Init initializes the instance.
func (o NEOnDemandRuleConnect) Init() NEOnDemandRuleConnect {
	rv := objc.Send[NEOnDemandRuleConnect](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o NEOnDemandRuleConnect) Autorelease() NEOnDemandRuleConnect {
	rv := objc.Send[NEOnDemandRuleConnect](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEOnDemandRuleConnect creates a new NEOnDemandRuleConnect instance.
func NewNEOnDemandRuleConnect() NEOnDemandRuleConnect {
	class := getNEOnDemandRuleConnectClass()
	rv := objc.Send[NEOnDemandRuleConnect](objc.ID(class.class), objc.Sel("new"))
	return rv
}

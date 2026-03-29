// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NEOnDemandRuleDisconnect] class.
var (
	_NEOnDemandRuleDisconnectClass     NEOnDemandRuleDisconnectClass
	_NEOnDemandRuleDisconnectClassOnce sync.Once
)

func getNEOnDemandRuleDisconnectClass() NEOnDemandRuleDisconnectClass {
	_NEOnDemandRuleDisconnectClassOnce.Do(func() {
		_NEOnDemandRuleDisconnectClass = NEOnDemandRuleDisconnectClass{class: objc.GetClass("NEOnDemandRuleDisconnect")}
	})
	return _NEOnDemandRuleDisconnectClass
}

// GetNEOnDemandRuleDisconnectClass returns the class object for NEOnDemandRuleDisconnect.
func GetNEOnDemandRuleDisconnectClass() NEOnDemandRuleDisconnectClass {
	return getNEOnDemandRuleDisconnectClass()
}

type NEOnDemandRuleDisconnectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEOnDemandRuleDisconnectClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEOnDemandRuleDisconnectClass) Alloc() NEOnDemandRuleDisconnect {
	rv := objc.Send[NEOnDemandRuleDisconnect](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A VPN On Demand rule that disconnects the VPN.
//
// # Overview
// 
// When rules of this class match, the VPN connection is not started, and the
// VPN connection is disconnected if it is not already disconnected.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleDisconnect
type NEOnDemandRuleDisconnect struct {
	NEOnDemandRule
}

// NEOnDemandRuleDisconnectFromID constructs a [NEOnDemandRuleDisconnect] from an objc.ID.
//
// A VPN On Demand rule that disconnects the VPN.
func NEOnDemandRuleDisconnectFromID(id objc.ID) NEOnDemandRuleDisconnect {
	return NEOnDemandRuleDisconnect{NEOnDemandRule: NEOnDemandRuleFromID(id)}
}
// NOTE: NEOnDemandRuleDisconnect adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEOnDemandRuleDisconnect] class.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleDisconnect
type INEOnDemandRuleDisconnect interface {
	INEOnDemandRule
}

// Init initializes the instance.
func (o NEOnDemandRuleDisconnect) Init() NEOnDemandRuleDisconnect {
	rv := objc.Send[NEOnDemandRuleDisconnect](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o NEOnDemandRuleDisconnect) Autorelease() NEOnDemandRuleDisconnect {
	rv := objc.Send[NEOnDemandRuleDisconnect](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEOnDemandRuleDisconnect creates a new NEOnDemandRuleDisconnect instance.
func NewNEOnDemandRuleDisconnect() NEOnDemandRuleDisconnect {
	class := getNEOnDemandRuleDisconnectClass()
	rv := objc.Send[NEOnDemandRuleDisconnect](objc.ID(class.class), objc.Sel("new"))
	return rv
}


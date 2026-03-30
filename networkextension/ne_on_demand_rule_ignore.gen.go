// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [NEOnDemandRuleIgnore] class.
var (
	_NEOnDemandRuleIgnoreClass     NEOnDemandRuleIgnoreClass
	_NEOnDemandRuleIgnoreClassOnce sync.Once
)

func getNEOnDemandRuleIgnoreClass() NEOnDemandRuleIgnoreClass {
	_NEOnDemandRuleIgnoreClassOnce.Do(func() {
		_NEOnDemandRuleIgnoreClass = NEOnDemandRuleIgnoreClass{class: objc.GetClass("NEOnDemandRuleIgnore")}
	})
	return _NEOnDemandRuleIgnoreClass
}

// GetNEOnDemandRuleIgnoreClass returns the class object for NEOnDemandRuleIgnore.
func GetNEOnDemandRuleIgnoreClass() NEOnDemandRuleIgnoreClass {
	return getNEOnDemandRuleIgnoreClass()
}

type NEOnDemandRuleIgnoreClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEOnDemandRuleIgnoreClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEOnDemandRuleIgnoreClass) Alloc() NEOnDemandRuleIgnore {
	rv := objc.Send[NEOnDemandRuleIgnore](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A VPN On Demand rule that doesn’t change the status of the VPN.
//
// # Overview
//
// When rules of this class match, the VPN connection is not started, and the
// current status of the VPN connection is left unchanged.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleIgnore
type NEOnDemandRuleIgnore struct {
	NEOnDemandRule
}

// NEOnDemandRuleIgnoreFromID constructs a [NEOnDemandRuleIgnore] from an objc.ID.
//
// A VPN On Demand rule that doesn’t change the status of the VPN.
func NEOnDemandRuleIgnoreFromID(id objc.ID) NEOnDemandRuleIgnore {
	return NEOnDemandRuleIgnore{NEOnDemandRule: NEOnDemandRuleFromID(id)}
}

// NOTE: NEOnDemandRuleIgnore adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEOnDemandRuleIgnore] class.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleIgnore
type INEOnDemandRuleIgnore interface {
	INEOnDemandRule
}

// Init initializes the instance.
func (o NEOnDemandRuleIgnore) Init() NEOnDemandRuleIgnore {
	rv := objc.Send[NEOnDemandRuleIgnore](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o NEOnDemandRuleIgnore) Autorelease() NEOnDemandRuleIgnore {
	rv := objc.Send[NEOnDemandRuleIgnore](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEOnDemandRuleIgnore creates a new NEOnDemandRuleIgnore instance.
func NewNEOnDemandRuleIgnore() NEOnDemandRuleIgnore {
	class := getNEOnDemandRuleIgnoreClass()
	rv := objc.Send[NEOnDemandRuleIgnore](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CPXMutableEventDeferringPolicy] class.
var (
	_CPXMutableEventDeferringPolicyClass     CPXMutableEventDeferringPolicyClass
	_CPXMutableEventDeferringPolicyClassOnce sync.Once
)

func getCPXMutableEventDeferringPolicyClass() CPXMutableEventDeferringPolicyClass {
	_CPXMutableEventDeferringPolicyClassOnce.Do(func() {
		_CPXMutableEventDeferringPolicyClass = CPXMutableEventDeferringPolicyClass{class: objc.GetClass("CPXMutableEventDeferringPolicy")}
	})
	return _CPXMutableEventDeferringPolicyClass
}

// GetCPXMutableEventDeferringPolicyClass returns the class object for CPXMutableEventDeferringPolicy.
func GetCPXMutableEventDeferringPolicyClass() CPXMutableEventDeferringPolicyClass {
	return getCPXMutableEventDeferringPolicyClass()
}

type CPXMutableEventDeferringPolicyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXMutableEventDeferringPolicyClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXMutableEventDeferringPolicyClass) Alloc() CPXMutableEventDeferringPolicy {
	rv := objc.Send[CPXMutableEventDeferringPolicy](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXMutableEventDeferringPolicy
type CPXMutableEventDeferringPolicy struct {
	CPXEventDeferringPolicy
}

// CPXMutableEventDeferringPolicyFromID constructs a [CPXMutableEventDeferringPolicy] from an objc.ID.
func CPXMutableEventDeferringPolicyFromID(id objc.ID) CPXMutableEventDeferringPolicy {
	return CPXMutableEventDeferringPolicy{CPXEventDeferringPolicy: CPXEventDeferringPolicyFromID(id)}
}

// Ensure CPXMutableEventDeferringPolicy implements ICPXMutableEventDeferringPolicy.
var _ ICPXMutableEventDeferringPolicy = CPXMutableEventDeferringPolicy{}

// An interface definition for the [CPXMutableEventDeferringPolicy] class.
//
// See: https://developer.apple.com/documentation/SkyLight/CPXMutableEventDeferringPolicy
type ICPXMutableEventDeferringPolicy interface {
	ICPXEventDeferringPolicy
}

// Init initializes the instance.
func (c CPXMutableEventDeferringPolicy) Init() CPXMutableEventDeferringPolicy {
	rv := objc.Send[CPXMutableEventDeferringPolicy](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXMutableEventDeferringPolicy) Autorelease() CPXMutableEventDeferringPolicy {
	rv := objc.Send[CPXMutableEventDeferringPolicy](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXMutableEventDeferringPolicy creates a new CPXMutableEventDeferringPolicy instance.
func NewCPXMutableEventDeferringPolicy() CPXMutableEventDeferringPolicy {
	class := getCPXMutableEventDeferringPolicyClass()
	rv := objc.Send[CPXMutableEventDeferringPolicy](objc.ID(class.class), objc.Sel("new"))
	return rv
}

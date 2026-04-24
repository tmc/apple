// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXFocusPolicy] class.
var (
	_CPXFocusPolicyClass     CPXFocusPolicyClass
	_CPXFocusPolicyClassOnce sync.Once
)

func getCPXFocusPolicyClass() CPXFocusPolicyClass {
	_CPXFocusPolicyClassOnce.Do(func() {
		_CPXFocusPolicyClass = CPXFocusPolicyClass{class: objc.GetClass("_CPXFocusPolicy")}
	})
	return _CPXFocusPolicyClass
}

// GetCPXFocusPolicyClass returns the class object for _CPXFocusPolicy.
func GetCPXFocusPolicyClass() CPXFocusPolicyClass {
	return getCPXFocusPolicyClass()
}

type CPXFocusPolicyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXFocusPolicyClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXFocusPolicyClass) Alloc() CPXFocusPolicy {
	rv := objc.Send[CPXFocusPolicy](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXFocusPolicy.BringNextApplicationToFrontInternal]
//   - [CPXFocusPolicy.BringNextProcessToFront]
//   - [CPXFocusPolicy.DebugDescription]
//   - [CPXFocusPolicy.Description]
//   - [CPXFocusPolicy.Hash]
//   - [CPXFocusPolicy.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/_CPXFocusPolicy
type CPXFocusPolicy struct {
	objectivec.Object
}

// CPXFocusPolicyFromID constructs a [CPXFocusPolicy] from an objc.ID.
func CPXFocusPolicyFromID(id objc.ID) CPXFocusPolicy {
	return CPXFocusPolicy{objectivec.Object{ID: id}}
}

// Ensure CPXFocusPolicy implements ICPXFocusPolicy.
var _ ICPXFocusPolicy = CPXFocusPolicy{}

// An interface definition for the [CPXFocusPolicy] class.
//
// # Methods
//
//   - [ICPXFocusPolicy.BringNextApplicationToFrontInternal]
//   - [ICPXFocusPolicy.BringNextProcessToFront]
//   - [ICPXFocusPolicy.DebugDescription]
//   - [ICPXFocusPolicy.Description]
//   - [ICPXFocusPolicy.Hash]
//   - [ICPXFocusPolicy.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/_CPXFocusPolicy
type ICPXFocusPolicy interface {
	objectivec.IObject

	// Topic: Methods

	BringNextApplicationToFrontInternal(internal *CPSProcessRecRef)
	BringNextProcessToFront(front *CPSProcessRecRef)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c CPXFocusPolicy) Init() CPXFocusPolicy {
	rv := objc.Send[CPXFocusPolicy](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXFocusPolicy) Autorelease() CPXFocusPolicy {
	rv := objc.Send[CPXFocusPolicy](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXFocusPolicy creates a new CPXFocusPolicy instance.
func NewCPXFocusPolicy() CPXFocusPolicy {
	class := getCPXFocusPolicyClass()
	rv := objc.Send[CPXFocusPolicy](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXFocusPolicy/bringNextApplicationToFrontInternal:
func (c CPXFocusPolicy) BringNextApplicationToFrontInternal(internal *CPSProcessRecRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("bringNextApplicationToFrontInternal:"), internal)
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXFocusPolicy/bringNextProcessToFront:
func (c CPXFocusPolicy) BringNextProcessToFront(front *CPSProcessRecRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("bringNextProcessToFront:"), front)
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXFocusPolicy/debugDescription
func (c CPXFocusPolicy) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXFocusPolicy/description
func (c CPXFocusPolicy) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXFocusPolicy/hash
func (c CPXFocusPolicy) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXFocusPolicy/superclass
func (c CPXFocusPolicy) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}

// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXPSNWrapper] class.
var (
	_CPXPSNWrapperClass     CPXPSNWrapperClass
	_CPXPSNWrapperClassOnce sync.Once
)

func getCPXPSNWrapperClass() CPXPSNWrapperClass {
	_CPXPSNWrapperClassOnce.Do(func() {
		_CPXPSNWrapperClass = CPXPSNWrapperClass{class: objc.GetClass("_CPXPSNWrapper")}
	})
	return _CPXPSNWrapperClass
}

// GetCPXPSNWrapperClass returns the class object for _CPXPSNWrapper.
func GetCPXPSNWrapperClass() CPXPSNWrapperClass {
	return getCPXPSNWrapperClass()
}

type CPXPSNWrapperClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXPSNWrapperClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXPSNWrapperClass) Alloc() CPXPSNWrapper {
	rv := objc.Send[CPXPSNWrapper](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXPSNWrapper.AppendDescriptionToStream]
//   - [CPXPSNWrapper.DebugDescription]
//   - [CPXPSNWrapper.Description]
//   - [CPXPSNWrapper.Hash]
//   - [CPXPSNWrapper.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/_CPXPSNWrapper
type CPXPSNWrapper struct {
	objectivec.Object
}

// CPXPSNWrapperFromID constructs a [CPXPSNWrapper] from an objc.ID.
func CPXPSNWrapperFromID(id objc.ID) CPXPSNWrapper {
	return CPXPSNWrapper{objectivec.Object{ID: id}}
}

// Ensure CPXPSNWrapper implements ICPXPSNWrapper.
var _ ICPXPSNWrapper = CPXPSNWrapper{}

// An interface definition for the [CPXPSNWrapper] class.
//
// # Methods
//
//   - [ICPXPSNWrapper.AppendDescriptionToStream]
//   - [ICPXPSNWrapper.DebugDescription]
//   - [ICPXPSNWrapper.Description]
//   - [ICPXPSNWrapper.Hash]
//   - [ICPXPSNWrapper.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/_CPXPSNWrapper
type ICPXPSNWrapper interface {
	objectivec.IObject

	// Topic: Methods

	AppendDescriptionToStream(stream objectivec.IObject)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c CPXPSNWrapper) Init() CPXPSNWrapper {
	rv := objc.Send[CPXPSNWrapper](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXPSNWrapper) Autorelease() CPXPSNWrapper {
	rv := objc.Send[CPXPSNWrapper](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXPSNWrapper creates a new CPXPSNWrapper instance.
func NewCPXPSNWrapper() CPXPSNWrapper {
	class := getCPXPSNWrapperClass()
	rv := objc.Send[CPXPSNWrapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXPSNWrapper/appendDescriptionToStream:
func (c CPXPSNWrapper) AppendDescriptionToStream(stream objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("appendDescriptionToStream:"), stream)
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXPSNWrapper/debugDescription
func (c CPXPSNWrapper) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXPSNWrapper/description
func (c CPXPSNWrapper) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXPSNWrapper/hash
func (c CPXPSNWrapper) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXPSNWrapper/superclass
func (c CPXPSNWrapper) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}

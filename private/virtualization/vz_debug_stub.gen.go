// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZDebugStub] class.
var (
	_VZDebugStubClass     VZDebugStubClass
	_VZDebugStubClassOnce sync.Once
)

func getVZDebugStubClass() VZDebugStubClass {
	_VZDebugStubClassOnce.Do(func() {
		_VZDebugStubClass = VZDebugStubClass{class: objc.GetClass("_VZDebugStub")}
	})
	return _VZDebugStubClass
}

// GetVZDebugStubClass returns the class object for _VZDebugStub.
func GetVZDebugStubClass() VZDebugStubClass {
	return getVZDebugStubClass()
}

type VZDebugStubClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZDebugStubClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZDebugStubClass) Alloc() VZDebugStub {
	rv := objc.Send[VZDebugStub](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZDebugStub._debugStub]
//   - [VZDebugStub._init]
// See: https://developer.apple.com/documentation/Virtualization/_VZDebugStub
type VZDebugStub struct {
	objectivec.Object
}

// VZDebugStubFromID constructs a [VZDebugStub] from an objc.ID.
func VZDebugStubFromID(id objc.ID) VZDebugStub {
	return VZDebugStub{objectivec.Object{ID: id}}
}
// Ensure VZDebugStub implements IVZDebugStub.
var _ IVZDebugStub = VZDebugStub{}

// An interface definition for the [VZDebugStub] class.
//
// # Methods
//
//   - [IVZDebugStub._debugStub]
//   - [IVZDebugStub._init]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZDebugStub
type IVZDebugStub interface {
	objectivec.IObject

	// Topic: Methods

	_debugStub() objectivec.IObject
	_init() objectivec.IObject
}

// Init initializes the instance.
func (v VZDebugStub) Init() VZDebugStub {
	rv := objc.Send[VZDebugStub](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZDebugStub) Autorelease() VZDebugStub {
	rv := objc.Send[VZDebugStub](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDebugStub creates a new VZDebugStub instance.
func NewVZDebugStub() VZDebugStub {
	class := getVZDebugStubClass()
	rv := objc.Send[VZDebugStub](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZDebugStub/_init
func (v VZDebugStub) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZDebugStub/_debugStub
func (v VZDebugStub) _debugStub() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_debugStub"))
	return objectivec.Object{ID: rv}
}


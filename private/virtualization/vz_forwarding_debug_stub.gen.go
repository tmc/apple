// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZForwardingDebugStub] class.
var (
	_VZForwardingDebugStubClass     VZForwardingDebugStubClass
	_VZForwardingDebugStubClassOnce sync.Once
)

func getVZForwardingDebugStubClass() VZForwardingDebugStubClass {
	_VZForwardingDebugStubClassOnce.Do(func() {
		_VZForwardingDebugStubClass = VZForwardingDebugStubClass{class: objc.GetClass("_VZForwardingDebugStub")}
	})
	return _VZForwardingDebugStubClass
}

// GetVZForwardingDebugStubClass returns the class object for _VZForwardingDebugStub.
func GetVZForwardingDebugStubClass() VZForwardingDebugStubClass {
	return getVZForwardingDebugStubClass()
}

type VZForwardingDebugStubClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZForwardingDebugStubClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZForwardingDebugStubClass) Alloc() VZForwardingDebugStub {
	rv := objc.Send[VZForwardingDebugStub](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZForwardingDebugStub._initWithDebugStub]
// See: https://developer.apple.com/documentation/Virtualization/_VZForwardingDebugStub
type VZForwardingDebugStub struct {
	VZDebugStub
}

// VZForwardingDebugStubFromID constructs a [VZForwardingDebugStub] from an objc.ID.
func VZForwardingDebugStubFromID(id objc.ID) VZForwardingDebugStub {
	return VZForwardingDebugStub{VZDebugStub: VZDebugStubFromID(id)}
}
// Ensure VZForwardingDebugStub implements IVZForwardingDebugStub.
var _ IVZForwardingDebugStub = VZForwardingDebugStub{}

// An interface definition for the [VZForwardingDebugStub] class.
//
// # Methods
//
//   - [IVZForwardingDebugStub._initWithDebugStub]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZForwardingDebugStub
type IVZForwardingDebugStub interface {
	IVZDebugStub

	// Topic: Methods

	_initWithDebugStub(stub unsafe.Pointer) objectivec.IObject
}

// Init initializes the instance.
func (v VZForwardingDebugStub) Init() VZForwardingDebugStub {
	rv := objc.Send[VZForwardingDebugStub](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZForwardingDebugStub) Autorelease() VZForwardingDebugStub {
	rv := objc.Send[VZForwardingDebugStub](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZForwardingDebugStub creates a new VZForwardingDebugStub instance.
func NewVZForwardingDebugStub() VZForwardingDebugStub {
	class := getVZForwardingDebugStubClass()
	rv := objc.Send[VZForwardingDebugStub](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZForwardingDebugStub/_initWithDebugStub:
func (v VZForwardingDebugStub) _initWithDebugStub(stub unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_initWithDebugStub:"), stub)
	return objectivec.Object{ID: rv}
}

// InitWithDebugStub is an exported wrapper for the private method _initWithDebugStub.
func (v VZForwardingDebugStub) InitWithDebugStub(stub unsafe.Pointer) objectivec.IObject {
	return v._initWithDebugStub(stub)
}


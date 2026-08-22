// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

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
	rv := objc.SendIfResponds[VZForwardingDebugStub](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZForwardingDebugStub._initWithDebugStub]
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
type IVZForwardingDebugStub interface {
	IVZDebugStub

	// Topic: Methods

	_initWithDebugStub(stub *DebugStub) objectivec.IObject
}

// Init initializes the instance.
func (v VZForwardingDebugStub) Init() VZForwardingDebugStub {
	rv := objc.SendIfResponds[VZForwardingDebugStub](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZForwardingDebugStub) Autorelease() VZForwardingDebugStub {
	rv := objc.SendIfResponds[VZForwardingDebugStub](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZForwardingDebugStub creates a new VZForwardingDebugStub instance.
func NewVZForwardingDebugStub() VZForwardingDebugStub {
	class := getVZForwardingDebugStubClass()
	rv := objc.SendIfResponds[VZForwardingDebugStub](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZForwardingDebugStub) _initWithDebugStub(stub *DebugStub) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_initWithDebugStub:"), unsafe.Pointer(stub))
	return objectivec.Object{ID: rv}
}

// InitWithDebugStub is an exported wrapper for the private method _initWithDebugStub.
func (v VZForwardingDebugStub) InitWithDebugStub(stub *DebugStub) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_initWithDebugStub:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_initWithDebugStub:"}
		return nil, err
	}
	return v._initWithDebugStub(stub), nil
}

// CanInitWithDebugStub reports whether the receiver responds to the private selector _initWithDebugStub:.
func (v VZForwardingDebugStub) CanInitWithDebugStub() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_initWithDebugStub:"))
}

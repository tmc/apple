// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZGDBDebugStub] class.
var (
	_VZGDBDebugStubClass     VZGDBDebugStubClass
	_VZGDBDebugStubClassOnce sync.Once
)

func getVZGDBDebugStubClass() VZGDBDebugStubClass {
	_VZGDBDebugStubClassOnce.Do(func() {
		_VZGDBDebugStubClass = VZGDBDebugStubClass{class: objc.GetClass("_VZGDBDebugStub")}
	})
	return _VZGDBDebugStubClass
}

// GetVZGDBDebugStubClass returns the class object for _VZGDBDebugStub.
func GetVZGDBDebugStubClass() VZGDBDebugStubClass {
	return getVZGDBDebugStubClass()
}

type VZGDBDebugStubClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZGDBDebugStubClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZGDBDebugStubClass) Alloc() VZGDBDebugStub {
	rv := objc.Send[VZGDBDebugStub](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZGDBDebugStub.Delegate]
//   - [VZGDBDebugStub.SetDelegate]
//   - [VZGDBDebugStub.Port]
// See: https://developer.apple.com/documentation/Virtualization/_VZGDBDebugStub
type VZGDBDebugStub struct {
	VZDebugStub
}

// VZGDBDebugStubFromID constructs a [VZGDBDebugStub] from an objc.ID.
func VZGDBDebugStubFromID(id objc.ID) VZGDBDebugStub {
	return VZGDBDebugStub{VZDebugStub: VZDebugStubFromID(id)}
}
// Ensure VZGDBDebugStub implements IVZGDBDebugStub.
var _ IVZGDBDebugStub = VZGDBDebugStub{}

// An interface definition for the [VZGDBDebugStub] class.
//
// # Methods
//
//   - [IVZGDBDebugStub.Delegate]
//   - [IVZGDBDebugStub.SetDelegate]
//   - [IVZGDBDebugStub.Port]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZGDBDebugStub
type IVZGDBDebugStub interface {
	IVZDebugStub

	// Topic: Methods

	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)
	Port() uint16
}

// Init initializes the instance.
func (v VZGDBDebugStub) Init() VZGDBDebugStub {
	rv := objc.Send[VZGDBDebugStub](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZGDBDebugStub) Autorelease() VZGDBDebugStub {
	rv := objc.Send[VZGDBDebugStub](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGDBDebugStub creates a new VZGDBDebugStub instance.
func NewVZGDBDebugStub() VZGDBDebugStub {
	class := getVZGDBDebugStubClass()
	rv := objc.Send[VZGDBDebugStub](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZGDBDebugStub/delegate
func (v VZGDBDebugStub) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (v VZGDBDebugStub) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](v.ID, objc.Sel("setDelegate:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZGDBDebugStub/port
func (v VZGDBDebugStub) Port() uint16 {
	rv := objc.Send[uint16](v.ID, objc.Sel("port"))
	return rv
}


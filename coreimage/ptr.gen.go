// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Ptr] class.
var (
	_PtrClass     PtrClass
	_PtrClassOnce sync.Once
)

func getPtrClass() PtrClass {
	_PtrClassOnce.Do(func() {
		_PtrClass = PtrClass{class: objc.GetClass("ptr")}
	})
	return _PtrClass
}

// GetPtrClass returns the class object for ptr.
func GetPtrClass() PtrClass {
	return getPtrClass()
}

type PtrClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PtrClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PtrClass) Alloc() Ptr {
	rv := objc.Send[Ptr](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIVector/union_(unnamed)/ptr
type Ptr struct {
	objectivec.Object
}

// PtrFromID constructs a [Ptr] from an objc.ID.
func PtrFromID(id objc.ID) Ptr {
	return Ptr{objectivec.Object{ID: id}}
}

// Ensure Ptr implements IPtr.
var _ IPtr = Ptr{}

// An interface definition for the [Ptr] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVector/union_(unnamed)/ptr
type IPtr interface {
	objectivec.IObject
}

// Init initializes the instance.
func (p Ptr) Init() Ptr {
	rv := objc.Send[Ptr](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p Ptr) Autorelease() Ptr {
	rv := objc.Send[Ptr](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPtr creates a new Ptr instance.
func NewPtr() Ptr {
	class := getPtrClass()
	rv := objc.Send[Ptr](objc.ID(class.class), objc.Sel("new"))
	return rv
}

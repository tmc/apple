// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZKeyboardEventTap] class.
var (
	_VZKeyboardEventTapClass     VZKeyboardEventTapClass
	_VZKeyboardEventTapClassOnce sync.Once
)

func getVZKeyboardEventTapClass() VZKeyboardEventTapClass {
	_VZKeyboardEventTapClassOnce.Do(func() {
		_VZKeyboardEventTapClass = VZKeyboardEventTapClass{class: objc.GetClass("VZKeyboardEventTap")}
	})
	return _VZKeyboardEventTapClass
}

// GetVZKeyboardEventTapClass returns the class object for VZKeyboardEventTap.
func GetVZKeyboardEventTapClass() VZKeyboardEventTapClass {
	return getVZKeyboardEventTapClass()
}

type VZKeyboardEventTapClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZKeyboardEventTapClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZKeyboardEventTapClass) Alloc() VZKeyboardEventTap {
	rv := objc.Send[VZKeyboardEventTap](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZKeyboardEventTap
type VZKeyboardEventTap struct {
	objectivec.Object
}

// VZKeyboardEventTapFromID constructs a [VZKeyboardEventTap] from an objc.ID.
func VZKeyboardEventTapFromID(id objc.ID) VZKeyboardEventTap {
	return VZKeyboardEventTap{objectivec.Object{ID: id}}
}
// Ensure VZKeyboardEventTap implements IVZKeyboardEventTap.
var _ IVZKeyboardEventTap = VZKeyboardEventTap{}

// An interface definition for the [VZKeyboardEventTap] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZKeyboardEventTap
type IVZKeyboardEventTap interface {
	objectivec.IObject
}

// Init initializes the instance.
func (k VZKeyboardEventTap) Init() VZKeyboardEventTap {
	rv := objc.Send[VZKeyboardEventTap](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k VZKeyboardEventTap) Autorelease() VZKeyboardEventTap {
	rv := objc.Send[VZKeyboardEventTap](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZKeyboardEventTap creates a new VZKeyboardEventTap instance.
func NewVZKeyboardEventTap() VZKeyboardEventTap {
	class := getVZKeyboardEventTapClass()
	rv := objc.Send[VZKeyboardEventTap](objc.ID(class.class), objc.Sel("new"))
	return rv
}


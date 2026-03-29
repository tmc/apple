// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZDirectoryShare] class.
var (
	_VZDirectoryShareClass     VZDirectoryShareClass
	_VZDirectoryShareClassOnce sync.Once
)

func getVZDirectoryShareClass() VZDirectoryShareClass {
	_VZDirectoryShareClassOnce.Do(func() {
		_VZDirectoryShareClass = VZDirectoryShareClass{class: objc.GetClass("VZDirectoryShare")}
	})
	return _VZDirectoryShareClass
}

// GetVZDirectoryShareClass returns the class object for VZDirectoryShare.
func GetVZDirectoryShareClass() VZDirectoryShareClass {
	return getVZDirectoryShareClass()
}

type VZDirectoryShareClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZDirectoryShareClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZDirectoryShareClass) Alloc() VZDirectoryShare {
	rv := objc.Send[VZDirectoryShare](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The base class for a directory share.
//
// # Overview
// 
// A directory share defines how the system exposes host directories to a
// guest VM.
// 
// Don’t instantiate [VZDirectoryShare] directly, use one of its subclasses
// such as [VZSingleDirectoryShare] or [VZMultipleDirectoryShare] instead.
//
// See: https://developer.apple.com/documentation/Virtualization/VZDirectoryShare
type VZDirectoryShare struct {
	objectivec.Object
}

// VZDirectoryShareFromID constructs a [VZDirectoryShare] from an objc.ID.
//
// The base class for a directory share.
func VZDirectoryShareFromID(id objc.ID) VZDirectoryShare {
	return VZDirectoryShare{objectivec.Object{ID: id}}
}
// NOTE: VZDirectoryShare adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZDirectoryShare] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZDirectoryShare
type IVZDirectoryShare interface {
	objectivec.IObject
}

// Init initializes the instance.
func (d VZDirectoryShare) Init() VZDirectoryShare {
	rv := objc.Send[VZDirectoryShare](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VZDirectoryShare) Autorelease() VZDirectoryShare {
	rv := objc.Send[VZDirectoryShare](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDirectoryShare creates a new VZDirectoryShare instance.
func NewVZDirectoryShare() VZDirectoryShare {
	class := getVZDirectoryShareClass()
	rv := objc.Send[VZDirectoryShare](objc.ID(class.class), objc.Sel("new"))
	return rv
}


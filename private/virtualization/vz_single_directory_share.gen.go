// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZSingleDirectoryShare] class.
var (
	_VZSingleDirectoryShareClass     VZSingleDirectoryShareClass
	_VZSingleDirectoryShareClassOnce sync.Once
)

func getVZSingleDirectoryShareClass() VZSingleDirectoryShareClass {
	_VZSingleDirectoryShareClassOnce.Do(func() {
		_VZSingleDirectoryShareClass = VZSingleDirectoryShareClass{class: objc.GetClass("VZSingleDirectoryShare")}
	})
	return _VZSingleDirectoryShareClass
}

// GetVZSingleDirectoryShareClass returns the class object for VZSingleDirectoryShare.
func GetVZSingleDirectoryShareClass() VZSingleDirectoryShareClass {
	return getVZSingleDirectoryShareClass()
}

type VZSingleDirectoryShareClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSingleDirectoryShareClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSingleDirectoryShareClass) Alloc() VZSingleDirectoryShare {
	rv := objc.SendIfResponds[VZSingleDirectoryShare](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZSingleDirectoryShare struct {
	VZDirectoryShare
}

// VZSingleDirectoryShareFromID constructs a [VZSingleDirectoryShare] from an objc.ID.
func VZSingleDirectoryShareFromID(id objc.ID) VZSingleDirectoryShare {
	return VZSingleDirectoryShare{VZDirectoryShare: VZDirectoryShareFromID(id)}
}

// Ensure VZSingleDirectoryShare implements IVZSingleDirectoryShare.
var _ IVZSingleDirectoryShare = VZSingleDirectoryShare{}

// An interface definition for the [VZSingleDirectoryShare] class.
type IVZSingleDirectoryShare interface {
	IVZDirectoryShare
}

// Init initializes the instance.
func (v VZSingleDirectoryShare) Init() VZSingleDirectoryShare {
	rv := objc.SendIfResponds[VZSingleDirectoryShare](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZSingleDirectoryShare) Autorelease() VZSingleDirectoryShare {
	rv := objc.SendIfResponds[VZSingleDirectoryShare](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSingleDirectoryShare creates a new VZSingleDirectoryShare instance.
func NewVZSingleDirectoryShare() VZSingleDirectoryShare {
	class := getVZSingleDirectoryShareClass()
	rv := objc.SendIfResponds[VZSingleDirectoryShare](objc.ID(class.class), objc.Sel("new"))
	return rv
}

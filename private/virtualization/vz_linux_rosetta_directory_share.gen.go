// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZLinuxRosettaDirectoryShare] class.
var (
	_VZLinuxRosettaDirectoryShareClass     VZLinuxRosettaDirectoryShareClass
	_VZLinuxRosettaDirectoryShareClassOnce sync.Once
)

func getVZLinuxRosettaDirectoryShareClass() VZLinuxRosettaDirectoryShareClass {
	_VZLinuxRosettaDirectoryShareClassOnce.Do(func() {
		_VZLinuxRosettaDirectoryShareClass = VZLinuxRosettaDirectoryShareClass{class: objc.GetClass("VZLinuxRosettaDirectoryShare")}
	})
	return _VZLinuxRosettaDirectoryShareClass
}

// GetVZLinuxRosettaDirectoryShareClass returns the class object for VZLinuxRosettaDirectoryShare.
func GetVZLinuxRosettaDirectoryShareClass() VZLinuxRosettaDirectoryShareClass {
	return getVZLinuxRosettaDirectoryShareClass()
}

type VZLinuxRosettaDirectoryShareClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZLinuxRosettaDirectoryShareClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZLinuxRosettaDirectoryShareClass) Alloc() VZLinuxRosettaDirectoryShare {
	rv := objc.SendIfResponds[VZLinuxRosettaDirectoryShare](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZLinuxRosettaDirectoryShare struct {
	VZDirectoryShare
}

// VZLinuxRosettaDirectoryShareFromID constructs a [VZLinuxRosettaDirectoryShare] from an objc.ID.
func VZLinuxRosettaDirectoryShareFromID(id objc.ID) VZLinuxRosettaDirectoryShare {
	return VZLinuxRosettaDirectoryShare{VZDirectoryShare: VZDirectoryShareFromID(id)}
}

// Ensure VZLinuxRosettaDirectoryShare implements IVZLinuxRosettaDirectoryShare.
var _ IVZLinuxRosettaDirectoryShare = VZLinuxRosettaDirectoryShare{}

// An interface definition for the [VZLinuxRosettaDirectoryShare] class.
type IVZLinuxRosettaDirectoryShare interface {
	IVZDirectoryShare
}

// Init initializes the instance.
func (v VZLinuxRosettaDirectoryShare) Init() VZLinuxRosettaDirectoryShare {
	rv := objc.SendIfResponds[VZLinuxRosettaDirectoryShare](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZLinuxRosettaDirectoryShare) Autorelease() VZLinuxRosettaDirectoryShare {
	rv := objc.SendIfResponds[VZLinuxRosettaDirectoryShare](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinuxRosettaDirectoryShare creates a new VZLinuxRosettaDirectoryShare instance.
func NewVZLinuxRosettaDirectoryShare() VZLinuxRosettaDirectoryShare {
	class := getVZLinuxRosettaDirectoryShareClass()
	rv := objc.SendIfResponds[VZLinuxRosettaDirectoryShare](objc.ID(class.class), objc.Sel("new"))
	return rv
}

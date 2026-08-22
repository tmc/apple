// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMultipleDirectoryShare] class.
var (
	_VZMultipleDirectoryShareClass     VZMultipleDirectoryShareClass
	_VZMultipleDirectoryShareClassOnce sync.Once
)

func getVZMultipleDirectoryShareClass() VZMultipleDirectoryShareClass {
	_VZMultipleDirectoryShareClassOnce.Do(func() {
		_VZMultipleDirectoryShareClass = VZMultipleDirectoryShareClass{class: objc.GetClass("VZMultipleDirectoryShare")}
	})
	return _VZMultipleDirectoryShareClass
}

// GetVZMultipleDirectoryShareClass returns the class object for VZMultipleDirectoryShare.
func GetVZMultipleDirectoryShareClass() VZMultipleDirectoryShareClass {
	return getVZMultipleDirectoryShareClass()
}

type VZMultipleDirectoryShareClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMultipleDirectoryShareClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMultipleDirectoryShareClass) Alloc() VZMultipleDirectoryShare {
	rv := objc.SendIfResponds[VZMultipleDirectoryShare](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZMultipleDirectoryShare struct {
	VZDirectoryShare
}

// VZMultipleDirectoryShareFromID constructs a [VZMultipleDirectoryShare] from an objc.ID.
func VZMultipleDirectoryShareFromID(id objc.ID) VZMultipleDirectoryShare {
	return VZMultipleDirectoryShare{VZDirectoryShare: VZDirectoryShareFromID(id)}
}

// Ensure VZMultipleDirectoryShare implements IVZMultipleDirectoryShare.
var _ IVZMultipleDirectoryShare = VZMultipleDirectoryShare{}

// An interface definition for the [VZMultipleDirectoryShare] class.
type IVZMultipleDirectoryShare interface {
	IVZDirectoryShare
}

// Init initializes the instance.
func (v VZMultipleDirectoryShare) Init() VZMultipleDirectoryShare {
	rv := objc.SendIfResponds[VZMultipleDirectoryShare](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMultipleDirectoryShare) Autorelease() VZMultipleDirectoryShare {
	rv := objc.SendIfResponds[VZMultipleDirectoryShare](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMultipleDirectoryShare creates a new VZMultipleDirectoryShare instance.
func NewVZMultipleDirectoryShare() VZMultipleDirectoryShare {
	class := getVZMultipleDirectoryShareClass()
	rv := objc.SendIfResponds[VZMultipleDirectoryShare](objc.ID(class.class), objc.Sel("new"))
	return rv
}

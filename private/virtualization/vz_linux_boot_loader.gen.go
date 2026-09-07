// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZLinuxBootLoader] class.
var (
	_VZLinuxBootLoaderClass     VZLinuxBootLoaderClass
	_VZLinuxBootLoaderClassOnce sync.Once
)

func getVZLinuxBootLoaderClass() VZLinuxBootLoaderClass {
	_VZLinuxBootLoaderClassOnce.Do(func() {
		_VZLinuxBootLoaderClass = VZLinuxBootLoaderClass{class: objc.GetClass("VZLinuxBootLoader")}
	})
	return _VZLinuxBootLoaderClass
}

// GetVZLinuxBootLoaderClass returns the class object for VZLinuxBootLoader.
func GetVZLinuxBootLoaderClass() VZLinuxBootLoaderClass {
	return getVZLinuxBootLoaderClass()
}

type VZLinuxBootLoaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZLinuxBootLoaderClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZLinuxBootLoaderClass) Alloc() VZLinuxBootLoader {
	rv := objc.SendIfResponds[VZLinuxBootLoader](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZLinuxBootLoader struct {
	VZBootLoader
}

// VZLinuxBootLoaderFromID constructs a [VZLinuxBootLoader] from an objc.ID.
func VZLinuxBootLoaderFromID(id objc.ID) VZLinuxBootLoader {
	return VZLinuxBootLoader{VZBootLoader: VZBootLoaderFromID(id)}
}

// Ensure VZLinuxBootLoader implements IVZLinuxBootLoader.
var _ IVZLinuxBootLoader = VZLinuxBootLoader{}

// An interface definition for the [VZLinuxBootLoader] class.
type IVZLinuxBootLoader interface {
	IVZBootLoader
}

// Init initializes the instance.
func (v VZLinuxBootLoader) Init() VZLinuxBootLoader {
	rv := objc.SendIfResponds[VZLinuxBootLoader](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZLinuxBootLoader) Autorelease() VZLinuxBootLoader {
	rv := objc.SendIfResponds[VZLinuxBootLoader](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinuxBootLoader creates a new VZLinuxBootLoader instance.
func NewVZLinuxBootLoader() VZLinuxBootLoader {
	class := getVZLinuxBootLoaderClass()
	rv := objc.SendIfResponds[VZLinuxBootLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

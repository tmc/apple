// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZLinuxRosettaUnixSocketCachingOptions] class.
var (
	_VZLinuxRosettaUnixSocketCachingOptionsClass     VZLinuxRosettaUnixSocketCachingOptionsClass
	_VZLinuxRosettaUnixSocketCachingOptionsClassOnce sync.Once
)

func getVZLinuxRosettaUnixSocketCachingOptionsClass() VZLinuxRosettaUnixSocketCachingOptionsClass {
	_VZLinuxRosettaUnixSocketCachingOptionsClassOnce.Do(func() {
		_VZLinuxRosettaUnixSocketCachingOptionsClass = VZLinuxRosettaUnixSocketCachingOptionsClass{class: objc.GetClass("VZLinuxRosettaUnixSocketCachingOptions")}
	})
	return _VZLinuxRosettaUnixSocketCachingOptionsClass
}

// GetVZLinuxRosettaUnixSocketCachingOptionsClass returns the class object for VZLinuxRosettaUnixSocketCachingOptions.
func GetVZLinuxRosettaUnixSocketCachingOptionsClass() VZLinuxRosettaUnixSocketCachingOptionsClass {
	return getVZLinuxRosettaUnixSocketCachingOptionsClass()
}

type VZLinuxRosettaUnixSocketCachingOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZLinuxRosettaUnixSocketCachingOptionsClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZLinuxRosettaUnixSocketCachingOptionsClass) Alloc() VZLinuxRosettaUnixSocketCachingOptions {
	rv := objc.SendIfResponds[VZLinuxRosettaUnixSocketCachingOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZLinuxRosettaUnixSocketCachingOptions struct {
	VZLinuxRosettaCachingOptions
}

// VZLinuxRosettaUnixSocketCachingOptionsFromID constructs a [VZLinuxRosettaUnixSocketCachingOptions] from an objc.ID.
func VZLinuxRosettaUnixSocketCachingOptionsFromID(id objc.ID) VZLinuxRosettaUnixSocketCachingOptions {
	return VZLinuxRosettaUnixSocketCachingOptions{VZLinuxRosettaCachingOptions: VZLinuxRosettaCachingOptionsFromID(id)}
}

// Ensure VZLinuxRosettaUnixSocketCachingOptions implements IVZLinuxRosettaUnixSocketCachingOptions.
var _ IVZLinuxRosettaUnixSocketCachingOptions = VZLinuxRosettaUnixSocketCachingOptions{}

// An interface definition for the [VZLinuxRosettaUnixSocketCachingOptions] class.
type IVZLinuxRosettaUnixSocketCachingOptions interface {
	IVZLinuxRosettaCachingOptions
}

// Init initializes the instance.
func (v VZLinuxRosettaUnixSocketCachingOptions) Init() VZLinuxRosettaUnixSocketCachingOptions {
	rv := objc.SendIfResponds[VZLinuxRosettaUnixSocketCachingOptions](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZLinuxRosettaUnixSocketCachingOptions) Autorelease() VZLinuxRosettaUnixSocketCachingOptions {
	rv := objc.SendIfResponds[VZLinuxRosettaUnixSocketCachingOptions](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinuxRosettaUnixSocketCachingOptions creates a new VZLinuxRosettaUnixSocketCachingOptions instance.
func NewVZLinuxRosettaUnixSocketCachingOptions() VZLinuxRosettaUnixSocketCachingOptions {
	class := getVZLinuxRosettaUnixSocketCachingOptionsClass()
	rv := objc.SendIfResponds[VZLinuxRosettaUnixSocketCachingOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

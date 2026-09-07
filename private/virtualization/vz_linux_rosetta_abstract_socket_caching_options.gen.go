// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZLinuxRosettaAbstractSocketCachingOptions] class.
var (
	_VZLinuxRosettaAbstractSocketCachingOptionsClass     VZLinuxRosettaAbstractSocketCachingOptionsClass
	_VZLinuxRosettaAbstractSocketCachingOptionsClassOnce sync.Once
)

func getVZLinuxRosettaAbstractSocketCachingOptionsClass() VZLinuxRosettaAbstractSocketCachingOptionsClass {
	_VZLinuxRosettaAbstractSocketCachingOptionsClassOnce.Do(func() {
		_VZLinuxRosettaAbstractSocketCachingOptionsClass = VZLinuxRosettaAbstractSocketCachingOptionsClass{class: objc.GetClass("VZLinuxRosettaAbstractSocketCachingOptions")}
	})
	return _VZLinuxRosettaAbstractSocketCachingOptionsClass
}

// GetVZLinuxRosettaAbstractSocketCachingOptionsClass returns the class object for VZLinuxRosettaAbstractSocketCachingOptions.
func GetVZLinuxRosettaAbstractSocketCachingOptionsClass() VZLinuxRosettaAbstractSocketCachingOptionsClass {
	return getVZLinuxRosettaAbstractSocketCachingOptionsClass()
}

type VZLinuxRosettaAbstractSocketCachingOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZLinuxRosettaAbstractSocketCachingOptionsClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZLinuxRosettaAbstractSocketCachingOptionsClass) Alloc() VZLinuxRosettaAbstractSocketCachingOptions {
	rv := objc.SendIfResponds[VZLinuxRosettaAbstractSocketCachingOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZLinuxRosettaAbstractSocketCachingOptions struct {
	VZLinuxRosettaCachingOptions
}

// VZLinuxRosettaAbstractSocketCachingOptionsFromID constructs a [VZLinuxRosettaAbstractSocketCachingOptions] from an objc.ID.
func VZLinuxRosettaAbstractSocketCachingOptionsFromID(id objc.ID) VZLinuxRosettaAbstractSocketCachingOptions {
	return VZLinuxRosettaAbstractSocketCachingOptions{VZLinuxRosettaCachingOptions: VZLinuxRosettaCachingOptionsFromID(id)}
}

// Ensure VZLinuxRosettaAbstractSocketCachingOptions implements IVZLinuxRosettaAbstractSocketCachingOptions.
var _ IVZLinuxRosettaAbstractSocketCachingOptions = VZLinuxRosettaAbstractSocketCachingOptions{}

// An interface definition for the [VZLinuxRosettaAbstractSocketCachingOptions] class.
type IVZLinuxRosettaAbstractSocketCachingOptions interface {
	IVZLinuxRosettaCachingOptions
}

// Init initializes the instance.
func (v VZLinuxRosettaAbstractSocketCachingOptions) Init() VZLinuxRosettaAbstractSocketCachingOptions {
	rv := objc.SendIfResponds[VZLinuxRosettaAbstractSocketCachingOptions](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZLinuxRosettaAbstractSocketCachingOptions) Autorelease() VZLinuxRosettaAbstractSocketCachingOptions {
	rv := objc.SendIfResponds[VZLinuxRosettaAbstractSocketCachingOptions](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinuxRosettaAbstractSocketCachingOptions creates a new VZLinuxRosettaAbstractSocketCachingOptions instance.
func NewVZLinuxRosettaAbstractSocketCachingOptions() VZLinuxRosettaAbstractSocketCachingOptions {
	class := getVZLinuxRosettaAbstractSocketCachingOptionsClass()
	rv := objc.SendIfResponds[VZLinuxRosettaAbstractSocketCachingOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

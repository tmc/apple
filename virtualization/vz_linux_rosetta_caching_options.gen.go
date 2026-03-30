// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZLinuxRosettaCachingOptions] class.
var (
	_VZLinuxRosettaCachingOptionsClass     VZLinuxRosettaCachingOptionsClass
	_VZLinuxRosettaCachingOptionsClassOnce sync.Once
)

func getVZLinuxRosettaCachingOptionsClass() VZLinuxRosettaCachingOptionsClass {
	_VZLinuxRosettaCachingOptionsClassOnce.Do(func() {
		_VZLinuxRosettaCachingOptionsClass = VZLinuxRosettaCachingOptionsClass{class: objc.GetClass("VZLinuxRosettaCachingOptions")}
	})
	return _VZLinuxRosettaCachingOptionsClass
}

// GetVZLinuxRosettaCachingOptionsClass returns the class object for VZLinuxRosettaCachingOptions.
func GetVZLinuxRosettaCachingOptionsClass() VZLinuxRosettaCachingOptionsClass {
	return getVZLinuxRosettaCachingOptionsClass()
}

type VZLinuxRosettaCachingOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZLinuxRosettaCachingOptionsClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZLinuxRosettaCachingOptionsClass) Alloc() VZLinuxRosettaCachingOptions {
	rv := objc.Send[VZLinuxRosettaCachingOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that defines UNIX socket-based caching options for
// Rosetta.
//
// # Overview
//
// [VZLinuxRosettaCachingOptions] define the communication mechanism between
// the Rosetta daemon and the Rosetta runtime.
//
// Don’t instantiate [VZLinuxRosettaCachingOptions] directly. Use one of its
// subclasses, such as [VZLinuxRosettaUnixSocketCachingOptions] or
// [VZLinuxRosettaAbstractSocketCachingOptions] instead.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaCachingOptions
type VZLinuxRosettaCachingOptions struct {
	objectivec.Object
}

// VZLinuxRosettaCachingOptionsFromID constructs a [VZLinuxRosettaCachingOptions] from an objc.ID.
//
// An abstract class that defines UNIX socket-based caching options for
// Rosetta.
func VZLinuxRosettaCachingOptionsFromID(id objc.ID) VZLinuxRosettaCachingOptions {
	return VZLinuxRosettaCachingOptions{objectivec.Object{ID: id}}
}

// Ensure VZLinuxRosettaCachingOptions implements IVZLinuxRosettaCachingOptions.
var _ IVZLinuxRosettaCachingOptions = VZLinuxRosettaCachingOptions{}

// An interface definition for the [VZLinuxRosettaCachingOptions] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaCachingOptions
type IVZLinuxRosettaCachingOptions interface {
	objectivec.IObject
}

// Init initializes the instance.
func (l VZLinuxRosettaCachingOptions) Init() VZLinuxRosettaCachingOptions {
	rv := objc.Send[VZLinuxRosettaCachingOptions](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l VZLinuxRosettaCachingOptions) Autorelease() VZLinuxRosettaCachingOptions {
	rv := objc.Send[VZLinuxRosettaCachingOptions](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinuxRosettaCachingOptions creates a new VZLinuxRosettaCachingOptions instance.
func NewVZLinuxRosettaCachingOptions() VZLinuxRosettaCachingOptions {
	class := getVZLinuxRosettaCachingOptionsClass()
	rv := objc.Send[VZLinuxRosettaCachingOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

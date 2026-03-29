// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZBootLoader] class.
var (
	_VZBootLoaderClass     VZBootLoaderClass
	_VZBootLoaderClassOnce sync.Once
)

func getVZBootLoaderClass() VZBootLoaderClass {
	_VZBootLoaderClassOnce.Do(func() {
		_VZBootLoaderClass = VZBootLoaderClass{class: objc.GetClass("VZBootLoader")}
	})
	return _VZBootLoaderClass
}

// GetVZBootLoaderClass returns the class object for VZBootLoader.
func GetVZBootLoaderClass() VZBootLoaderClass {
	return getVZBootLoaderClass()
}

type VZBootLoaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZBootLoaderClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZBootLoaderClass) Alloc() VZBootLoader {
	rv := objc.Send[VZBootLoader](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The base class that defines the management of the initial process of the
// guest system.
//
// # Overview
// 
// The [VZBootLoader] abstract class defines the common behaviors for booting
// a guest operating system into a VM. Don’t create instances of this class
// directly. Instead, instantiate the subclass that corresponds to the type of
// operating system you plan to load. For example, to create a boot loader
// object for a Linux kernel, create a [VZLinuxBootLoader] object; to create a
// boot loader object for installation using an ISO image create a
// [VZEFIBootLoader]. For a macOS system create [VZMacOSBootLoader].
//
// See: https://developer.apple.com/documentation/Virtualization/VZBootLoader
type VZBootLoader struct {
	objectivec.Object
}

// VZBootLoaderFromID constructs a [VZBootLoader] from an objc.ID.
//
// The base class that defines the management of the initial process of the
// guest system.
func VZBootLoaderFromID(id objc.ID) VZBootLoader {
	return VZBootLoader{objectivec.Object{ID: id}}
}
// NOTE: VZBootLoader adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZBootLoader] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZBootLoader
type IVZBootLoader interface {
	objectivec.IObject
}

// Init initializes the instance.
func (b VZBootLoader) Init() VZBootLoader {
	rv := objc.Send[VZBootLoader](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b VZBootLoader) Autorelease() VZBootLoader {
	rv := objc.Send[VZBootLoader](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZBootLoader creates a new VZBootLoader instance.
func NewVZBootLoader() VZBootLoader {
	class := getVZBootLoaderClass()
	rv := objc.Send[VZBootLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}


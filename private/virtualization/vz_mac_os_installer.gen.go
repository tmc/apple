// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacOSInstaller] class.
var (
	_VZMacOSInstallerClass     VZMacOSInstallerClass
	_VZMacOSInstallerClassOnce sync.Once
)

func getVZMacOSInstallerClass() VZMacOSInstallerClass {
	_VZMacOSInstallerClassOnce.Do(func() {
		_VZMacOSInstallerClass = VZMacOSInstallerClass{class: objc.GetClass("VZMacOSInstaller")}
	})
	return _VZMacOSInstallerClass
}

// GetVZMacOSInstallerClass returns the class object for VZMacOSInstaller.
func GetVZMacOSInstallerClass() VZMacOSInstallerClass {
	return getVZMacOSInstallerClass()
}

type VZMacOSInstallerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacOSInstallerClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacOSInstallerClass) Alloc() VZMacOSInstaller {
	rv := objc.Send[VZMacOSInstaller](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZMacOSInstaller._disableMobileDeviceUpdate]
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller
type VZMacOSInstaller struct {
	objectivec.Object
}

// VZMacOSInstallerFromID constructs a [VZMacOSInstaller] from an objc.ID.
func VZMacOSInstallerFromID(id objc.ID) VZMacOSInstaller {
	return VZMacOSInstaller{objectivec.Object{ID: id}}
}
// Ensure VZMacOSInstaller implements IVZMacOSInstaller.
var _ IVZMacOSInstaller = VZMacOSInstaller{}

// An interface definition for the [VZMacOSInstaller] class.
//
// # Methods
//
//   - [IVZMacOSInstaller._disableMobileDeviceUpdate]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller
type IVZMacOSInstaller interface {
	objectivec.IObject

	// Topic: Methods

	_disableMobileDeviceUpdate()
}

// Init initializes the instance.
func (m VZMacOSInstaller) Init() VZMacOSInstaller {
	rv := objc.Send[VZMacOSInstaller](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacOSInstaller) Autorelease() VZMacOSInstaller {
	rv := objc.Send[VZMacOSInstaller](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacOSInstaller creates a new VZMacOSInstaller instance.
func NewVZMacOSInstaller() VZMacOSInstaller {
	class := getVZMacOSInstallerClass()
	rv := objc.Send[VZMacOSInstaller](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller/_disableMobileDeviceUpdate
func (m VZMacOSInstaller) _disableMobileDeviceUpdate() {
	objc.Send[objc.ID](m.ID, objc.Sel("_disableMobileDeviceUpdate"))
}

// DisableMobileDeviceUpdate is an exported wrapper for the private method _disableMobileDeviceUpdate.
func (m VZMacOSInstaller) DisableMobileDeviceUpdate() {
	m._disableMobileDeviceUpdate()
}


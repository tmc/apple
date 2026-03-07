// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacOSBootLoader] class.
var (
	_VZMacOSBootLoaderClass     VZMacOSBootLoaderClass
	_VZMacOSBootLoaderClassOnce sync.Once
)

func getVZMacOSBootLoaderClass() VZMacOSBootLoaderClass {
	_VZMacOSBootLoaderClassOnce.Do(func() {
		_VZMacOSBootLoaderClass = VZMacOSBootLoaderClass{class: objc.GetClass("VZMacOSBootLoader")}
	})
	return _VZMacOSBootLoaderClass
}

// GetVZMacOSBootLoaderClass returns the class object for VZMacOSBootLoader.
func GetVZMacOSBootLoaderClass() VZMacOSBootLoaderClass {
	return getVZMacOSBootLoaderClass()
}

type VZMacOSBootLoaderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacOSBootLoaderClass) Alloc() VZMacOSBootLoader {
	rv := objc.Send[VZMacOSBootLoader](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An object that loads and configures a boot loader for running macOS on
// Apple silicon as a guest system of your VM.
//
// # Overview
// 
// You must use a [VZMacPlatformConfiguration] in conjunction with the macOS
// boot loader. It’s invalid to use it with any other platform
// configuration.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSBootLoader
type VZMacOSBootLoader struct {
	VZBootLoader
}

// VZMacOSBootLoaderFromID constructs a [VZMacOSBootLoader] from an objc.ID.
//
// An object that loads and configures a boot loader for running macOS on
// Apple silicon as a guest system of your VM.
func VZMacOSBootLoaderFromID(id objc.ID) VZMacOSBootLoader {
	return VZMacOSBootLoader{VZBootLoader: VZBootLoaderFromID(id)}
}
// NOTE: VZMacOSBootLoader adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZMacOSBootLoader] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSBootLoader
type IVZMacOSBootLoader interface {
	IVZBootLoader

	// The hardware platform to use.
	Platform() IVZPlatformConfiguration
	SetPlatform(value IVZPlatformConfiguration)
}





// Init initializes the instance.
func (m VZMacOSBootLoader) Init() VZMacOSBootLoader {
	rv := objc.Send[VZMacOSBootLoader](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacOSBootLoader) Autorelease() VZMacOSBootLoader {
	rv := objc.Send[VZMacOSBootLoader](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacOSBootLoader creates a new VZMacOSBootLoader instance.
func NewVZMacOSBootLoader() VZMacOSBootLoader {
	class := getVZMacOSBootLoaderClass()
	rv := objc.Send[VZMacOSBootLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}






















// The hardware platform to use.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/platform
func (m VZMacOSBootLoader) Platform() IVZPlatformConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("platform"))
	return VZPlatformConfigurationFromID(objc.ID(rv))
}
func (m VZMacOSBootLoader) SetPlatform(value IVZPlatformConfiguration) {
	objc.Send[struct{}](m.ID, objc.Sel("setPlatform:"), value)
}

























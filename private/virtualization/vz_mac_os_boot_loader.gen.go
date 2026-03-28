// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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

// Class returns the underlying Objective-C class pointer.
func (vc VZMacOSBootLoaderClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacOSBootLoaderClass) Alloc() VZMacOSBootLoader {
	rv := objc.Send[VZMacOSBootLoader](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZMacOSBootLoader._romURL]
//   - [VZMacOSBootLoader.Set_romURL]
//   - [VZMacOSBootLoader._setROMURL]
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSBootLoader
type VZMacOSBootLoader struct {
	VZBootLoader
}

// VZMacOSBootLoaderFromID constructs a [VZMacOSBootLoader] from an objc.ID.
func VZMacOSBootLoaderFromID(id objc.ID) VZMacOSBootLoader {
	return VZMacOSBootLoader{VZBootLoader: VZBootLoaderFromID(id)}
}
// Ensure VZMacOSBootLoader implements IVZMacOSBootLoader.
var _ IVZMacOSBootLoader = VZMacOSBootLoader{}

// An interface definition for the [VZMacOSBootLoader] class.
//
// # Methods
//
//   - [IVZMacOSBootLoader._romURL]
//   - [IVZMacOSBootLoader.Set_romURL]
//   - [IVZMacOSBootLoader._setROMURL]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSBootLoader
type IVZMacOSBootLoader interface {
	IVZBootLoader

	// Topic: Methods

	_romURL() foundation.INSURL
	Set_romURL(value foundation.INSURL)
	_setROMURL(romurl foundation.INSURL)
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

//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSBootLoader/_setROMURL:
func (m VZMacOSBootLoader) _setROMURL(romurl foundation.INSURL) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setROMURL:"), romurl)
}

// SetROMURL is an exported wrapper for the private method _setROMURL.
func (m VZMacOSBootLoader) SetROMURL(romurl foundation.INSURL) {
	m._setROMURL(romurl)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacOSBootLoader/_romURL
func (m VZMacOSBootLoader) _romURL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_romURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (m VZMacOSBootLoader) Set_romURL(value foundation.INSURL) {
	objc.Send[struct{}](m.ID, objc.Sel("set_romURL:"), value)
}


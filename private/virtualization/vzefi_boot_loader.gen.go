// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZEFIBootLoader] class.
var (
	_VZEFIBootLoaderClass     VZEFIBootLoaderClass
	_VZEFIBootLoaderClassOnce sync.Once
)

func getVZEFIBootLoaderClass() VZEFIBootLoaderClass {
	_VZEFIBootLoaderClassOnce.Do(func() {
		_VZEFIBootLoaderClass = VZEFIBootLoaderClass{class: objc.GetClass("VZEFIBootLoader")}
	})
	return _VZEFIBootLoaderClass
}

// GetVZEFIBootLoaderClass returns the class object for VZEFIBootLoader.
func GetVZEFIBootLoaderClass() VZEFIBootLoaderClass {
	return getVZEFIBootLoaderClass()
}

type VZEFIBootLoaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZEFIBootLoaderClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZEFIBootLoaderClass) Alloc() VZEFIBootLoader {
	rv := objc.Send[VZEFIBootLoader](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZEFIBootLoader._ROMImageURL]
//   - [VZEFIBootLoader.Set_ROMImageURL]
//   - [VZEFIBootLoader._setROMImageURL]
//
// See: https://developer.apple.com/documentation/Virtualization/VZEFIBootLoader
type VZEFIBootLoader struct {
	VZBootLoader
}

// VZEFIBootLoaderFromID constructs a [VZEFIBootLoader] from an objc.ID.
func VZEFIBootLoaderFromID(id objc.ID) VZEFIBootLoader {
	return VZEFIBootLoader{VZBootLoader: VZBootLoaderFromID(id)}
}

// Ensure VZEFIBootLoader implements IVZEFIBootLoader.
var _ IVZEFIBootLoader = VZEFIBootLoader{}

// An interface definition for the [VZEFIBootLoader] class.
//
// # Methods
//
//   - [IVZEFIBootLoader._ROMImageURL]
//   - [IVZEFIBootLoader.Set_ROMImageURL]
//   - [IVZEFIBootLoader._setROMImageURL]
//
// See: https://developer.apple.com/documentation/Virtualization/VZEFIBootLoader
type IVZEFIBootLoader interface {
	IVZBootLoader

	// Topic: Methods

	_ROMImageURL() foundation.INSURL
	Set_ROMImageURL(value foundation.INSURL)
	_setROMImageURL(url foundation.INSURL)
}

// Init initializes the instance.
func (e VZEFIBootLoader) Init() VZEFIBootLoader {
	rv := objc.Send[VZEFIBootLoader](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e VZEFIBootLoader) Autorelease() VZEFIBootLoader {
	rv := objc.Send[VZEFIBootLoader](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZEFIBootLoader creates a new VZEFIBootLoader instance.
func NewVZEFIBootLoader() VZEFIBootLoader {
	class := getVZEFIBootLoaderClass()
	rv := objc.Send[VZEFIBootLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZEFIBootLoader/_setROMImageURL:
func (e VZEFIBootLoader) _setROMImageURL(url foundation.INSURL) {
	objc.Send[objc.ID](e.ID, objc.Sel("_setROMImageURL:"), url)
}

// SetROMImageURL is an exported wrapper for the private method _setROMImageURL.
func (e VZEFIBootLoader) SetROMImageURL(url foundation.INSURL) {
	e._setROMImageURL(url)
}

// See: https://developer.apple.com/documentation/Virtualization/VZEFIBootLoader/_ROMImageURL
func (e VZEFIBootLoader) _ROMImageURL() foundation.INSURL {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_ROMImageURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (e VZEFIBootLoader) Set_ROMImageURL(value foundation.INSURL) {
	objc.Send[struct{}](e.ID, objc.Sel("set_ROMImageURL:"), value)
}

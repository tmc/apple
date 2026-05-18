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
func (v VZEFIBootLoader) Init() VZEFIBootLoader {
	rv := objc.Send[VZEFIBootLoader](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZEFIBootLoader) Autorelease() VZEFIBootLoader {
	rv := objc.Send[VZEFIBootLoader](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZEFIBootLoader creates a new VZEFIBootLoader instance.
func NewVZEFIBootLoader() VZEFIBootLoader {
	class := getVZEFIBootLoaderClass()
	rv := objc.Send[VZEFIBootLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZEFIBootLoader/_setROMImageURL:
func (v VZEFIBootLoader) _setROMImageURL(url foundation.INSURL) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setROMImageURL:"), url)
}

// SetROMImageURL is an exported wrapper for the private method _setROMImageURL.
func (v VZEFIBootLoader) SetROMImageURL(url foundation.INSURL) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setROMImageURL:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setROMImageURL:"}
		return err
	}
	v._setROMImageURL(url)
	return nil
}

// CanSetROMImageURL reports whether the receiver responds to the private selector _setROMImageURL:.
func (v VZEFIBootLoader) CanSetROMImageURL() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setROMImageURL:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZEFIBootLoader/_ROMImageURL
func (v VZEFIBootLoader) _ROMImageURL() foundation.INSURL {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_ROMImageURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// CanROMImageURL reports whether the receiver responds to the private selector _ROMImageURL.
func (v VZEFIBootLoader) CanROMImageURL() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_ROMImageURL"))
}

// ROMImageURL is an exported wrapper for the private property _ROMImageURL.
func (v VZEFIBootLoader) ROMImageURL() (foundation.INSURL, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_ROMImageURL")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_ROMImageURL"}
	}
	return v._ROMImageURL(), nil
}
func (v VZEFIBootLoader) Set_ROMImageURL(value foundation.INSURL) {
	objc.Send[struct{}](v.ID, objc.Sel("set_ROMImageURL:"), value)
}

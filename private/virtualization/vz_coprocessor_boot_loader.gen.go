// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZCoprocessorBootLoader] class.
var (
	_VZCoprocessorBootLoaderClass     VZCoprocessorBootLoaderClass
	_VZCoprocessorBootLoaderClassOnce sync.Once
)

func getVZCoprocessorBootLoaderClass() VZCoprocessorBootLoaderClass {
	_VZCoprocessorBootLoaderClassOnce.Do(func() {
		_VZCoprocessorBootLoaderClass = VZCoprocessorBootLoaderClass{class: objc.GetClass("_VZCoprocessorBootLoader")}
	})
	return _VZCoprocessorBootLoaderClass
}

// GetVZCoprocessorBootLoaderClass returns the class object for _VZCoprocessorBootLoader.
func GetVZCoprocessorBootLoaderClass() VZCoprocessorBootLoaderClass {
	return getVZCoprocessorBootLoaderClass()
}

type VZCoprocessorBootLoaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCoprocessorBootLoaderClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCoprocessorBootLoaderClass) Alloc() VZCoprocessorBootLoader {
	rv := objc.Send[VZCoprocessorBootLoader](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZCoprocessorBootLoader._bootLoaderForConfiguration]
//   - [VZCoprocessorBootLoader._romFileDescriptor]
//   - [VZCoprocessorBootLoader.Set_romFileDescriptor]
//   - [VZCoprocessorBootLoader._setROMFileDescriptor]
// See: https://developer.apple.com/documentation/Virtualization/_VZCoprocessorBootLoader
type VZCoprocessorBootLoader struct {
	VZBootLoader
}

// VZCoprocessorBootLoaderFromID constructs a [VZCoprocessorBootLoader] from an objc.ID.
func VZCoprocessorBootLoaderFromID(id objc.ID) VZCoprocessorBootLoader {
	return VZCoprocessorBootLoader{VZBootLoader: VZBootLoaderFromID(id)}
}
// Ensure VZCoprocessorBootLoader implements IVZCoprocessorBootLoader.
var _ IVZCoprocessorBootLoader = VZCoprocessorBootLoader{}

// An interface definition for the [VZCoprocessorBootLoader] class.
//
// # Methods
//
//   - [IVZCoprocessorBootLoader._bootLoaderForConfiguration]
//   - [IVZCoprocessorBootLoader._romFileDescriptor]
//   - [IVZCoprocessorBootLoader.Set_romFileDescriptor]
//   - [IVZCoprocessorBootLoader._setROMFileDescriptor]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCoprocessorBootLoader
type IVZCoprocessorBootLoader interface {
	IVZBootLoader

	// Topic: Methods

	_bootLoaderForConfiguration(configuration objectivec.IObject) objectivec.IObject
	_romFileDescriptor() objectivec.IObject
	Set_romFileDescriptor(value objectivec.IObject)
	_setROMFileDescriptor(descriptor objectivec.IObject)
}

// Init initializes the instance.
func (v VZCoprocessorBootLoader) Init() VZCoprocessorBootLoader {
	rv := objc.Send[VZCoprocessorBootLoader](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCoprocessorBootLoader) Autorelease() VZCoprocessorBootLoader {
	rv := objc.Send[VZCoprocessorBootLoader](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCoprocessorBootLoader creates a new VZCoprocessorBootLoader instance.
func NewVZCoprocessorBootLoader() VZCoprocessorBootLoader {
	class := getVZCoprocessorBootLoaderClass()
	rv := objc.Send[VZCoprocessorBootLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZCoprocessorBootLoader/_bootLoaderForConfiguration:
func (v VZCoprocessorBootLoader) _bootLoaderForConfiguration(configuration objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_bootLoaderForConfiguration:"), configuration)
	return objectivec.Object{ID: rv}
}

// BootLoaderForConfiguration is an exported wrapper for the private method _bootLoaderForConfiguration.
func (v VZCoprocessorBootLoader) BootLoaderForConfiguration(configuration objectivec.IObject) objectivec.IObject {
	return v._bootLoaderForConfiguration(configuration)
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCoprocessorBootLoader/_setROMFileDescriptor:
func (v VZCoprocessorBootLoader) _setROMFileDescriptor(descriptor objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setROMFileDescriptor:"), descriptor)
}

// SetROMFileDescriptor is an exported wrapper for the private method _setROMFileDescriptor.
func (v VZCoprocessorBootLoader) SetROMFileDescriptor(descriptor objectivec.IObject) {
	v._setROMFileDescriptor(descriptor)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCoprocessorBootLoader/_romFileDescriptor
func (v VZCoprocessorBootLoader) _romFileDescriptor() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_romFileDescriptor"))
	return objectivec.Object{ID: rv}
}
func (v VZCoprocessorBootLoader) Set_romFileDescriptor(value objectivec.IObject) {
	objc.Send[struct{}](v.ID, objc.Sel("set_romFileDescriptor:"), value)
}


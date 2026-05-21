// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
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

// # Methods
//
//   - [VZCoprocessorBootLoader._romFileDescriptor]
//   - [VZCoprocessorBootLoader.Set_romFileDescriptor]
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
//   - [IVZCoprocessorBootLoader._romFileDescriptor]
//   - [IVZCoprocessorBootLoader.Set_romFileDescriptor]
type IVZCoprocessorBootLoader interface {
	IVZBootLoader

	// Topic: Methods

	_romFileDescriptor() unsafe.Pointer
	Set_romFileDescriptor(value unsafe.Pointer)
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

func (v VZCoprocessorBootLoader) _romFileDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("_romFileDescriptor"))
	return rv
}

// CanRomFileDescriptor reports whether the receiver responds to the private selector _romFileDescriptor.
func (v VZCoprocessorBootLoader) CanRomFileDescriptor() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_romFileDescriptor"))
}

// RomFileDescriptor is an exported wrapper for the private property _romFileDescriptor.
func (v VZCoprocessorBootLoader) RomFileDescriptor() (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_romFileDescriptor")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_romFileDescriptor"}
	}
	return v._romFileDescriptor(), nil
}
func (v VZCoprocessorBootLoader) Set_romFileDescriptor(value unsafe.Pointer) {
	objc.Send[struct{}](v.ID, objc.Sel("set_romFileDescriptor:"), value)
}

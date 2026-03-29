// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZBinaryBootLoader] class.
var (
	_VZBinaryBootLoaderClass     VZBinaryBootLoaderClass
	_VZBinaryBootLoaderClassOnce sync.Once
)

func getVZBinaryBootLoaderClass() VZBinaryBootLoaderClass {
	_VZBinaryBootLoaderClassOnce.Do(func() {
		_VZBinaryBootLoaderClass = VZBinaryBootLoaderClass{class: objc.GetClass("_VZBinaryBootLoader")}
	})
	return _VZBinaryBootLoaderClass
}

// GetVZBinaryBootLoaderClass returns the class object for _VZBinaryBootLoader.
func GetVZBinaryBootLoaderClass() VZBinaryBootLoaderClass {
	return getVZBinaryBootLoaderClass()
}

type VZBinaryBootLoaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZBinaryBootLoaderClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZBinaryBootLoaderClass) Alloc() VZBinaryBootLoader {
	rv := objc.Send[VZBinaryBootLoader](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZBinaryBootLoader._bootLoaderForConfiguration]
//   - [VZBinaryBootLoader.EncodeWithEncoder]
//   - [VZBinaryBootLoader.EntryPointAddress]
//   - [VZBinaryBootLoader.Segments]
//   - [VZBinaryBootLoader.Validate]
//   - [VZBinaryBootLoader.InitWithSegmentsEntryPointAddress]
// See: https://developer.apple.com/documentation/Virtualization/_VZBinaryBootLoader
type VZBinaryBootLoader struct {
	VZBootLoader
}

// VZBinaryBootLoaderFromID constructs a [VZBinaryBootLoader] from an objc.ID.
func VZBinaryBootLoaderFromID(id objc.ID) VZBinaryBootLoader {
	return VZBinaryBootLoader{VZBootLoader: VZBootLoaderFromID(id)}
}
// Ensure VZBinaryBootLoader implements IVZBinaryBootLoader.
var _ IVZBinaryBootLoader = VZBinaryBootLoader{}

// An interface definition for the [VZBinaryBootLoader] class.
//
// # Methods
//
//   - [IVZBinaryBootLoader._bootLoaderForConfiguration]
//   - [IVZBinaryBootLoader.EncodeWithEncoder]
//   - [IVZBinaryBootLoader.EntryPointAddress]
//   - [IVZBinaryBootLoader.Segments]
//   - [IVZBinaryBootLoader.Validate]
//   - [IVZBinaryBootLoader.InitWithSegmentsEntryPointAddress]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZBinaryBootLoader
type IVZBinaryBootLoader interface {
	IVZBootLoader

	// Topic: Methods

	_bootLoaderForConfiguration(configuration objectivec.IObject) objectivec.IObject
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	EntryPointAddress() uint64
	Segments() foundation.INSArray
	Validate() objectivec.IObject
	InitWithSegmentsEntryPointAddress(segments objectivec.IObject, address uint64) VZBinaryBootLoader
}

// Init initializes the instance.
func (v VZBinaryBootLoader) Init() VZBinaryBootLoader {
	rv := objc.Send[VZBinaryBootLoader](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZBinaryBootLoader) Autorelease() VZBinaryBootLoader {
	rv := objc.Send[VZBinaryBootLoader](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZBinaryBootLoader creates a new VZBinaryBootLoader instance.
func NewVZBinaryBootLoader() VZBinaryBootLoader {
	class := getVZBinaryBootLoaderClass()
	rv := objc.Send[VZBinaryBootLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZBinaryBootLoader/initWithSegments:entryPointAddress:
func NewVZBinaryBootLoaderWithSegmentsEntryPointAddress(segments objectivec.IObject, address uint64) VZBinaryBootLoader {
	instance := getVZBinaryBootLoaderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSegments:entryPointAddress:"), segments, address)
	return VZBinaryBootLoaderFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZBinaryBootLoader/_bootLoaderForConfiguration:
func (v VZBinaryBootLoader) _bootLoaderForConfiguration(configuration objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_bootLoaderForConfiguration:"), configuration)
	return objectivec.Object{ID: rv}
}

// BootLoaderForConfiguration is an exported wrapper for the private method _bootLoaderForConfiguration.
func (v VZBinaryBootLoader) BootLoaderForConfiguration(configuration objectivec.IObject) objectivec.IObject {
	return v._bootLoaderForConfiguration(configuration)
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZBinaryBootLoader/encodeWithEncoder:
func (v VZBinaryBootLoader) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Virtualization/_VZBinaryBootLoader/validate
func (v VZBinaryBootLoader) Validate() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("validate"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZBinaryBootLoader/initWithSegments:entryPointAddress:
func (v VZBinaryBootLoader) InitWithSegmentsEntryPointAddress(segments objectivec.IObject, address uint64) VZBinaryBootLoader {
	rv := objc.Send[VZBinaryBootLoader](v.ID, objc.Sel("initWithSegments:entryPointAddress:"), segments, address)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZBinaryBootLoader/entryPointAddress
func (v VZBinaryBootLoader) EntryPointAddress() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("entryPointAddress"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZBinaryBootLoader/segments
func (v VZBinaryBootLoader) Segments() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("segments"))
	return foundation.NSArrayFromID(objc.ID(rv))
}


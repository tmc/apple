// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMemory] class.
var (
	_VZMemoryClass     VZMemoryClass
	_VZMemoryClassOnce sync.Once
)

func getVZMemoryClass() VZMemoryClass {
	_VZMemoryClassOnce.Do(func() {
		_VZMemoryClass = VZMemoryClass{class: objc.GetClass("_VZMemory")}
	})
	return _VZMemoryClass
}

// GetVZMemoryClass returns the class object for _VZMemory.
func GetVZMemoryClass() VZMemoryClass {
	return getVZMemoryClass()
}

type VZMemoryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMemoryClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMemoryClass) Alloc() VZMemory {
	rv := objc.Send[VZMemory](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZMemory.Length]
//   - [VZMemory.MutableBytes]
//   - [VZMemory.PhysicalAddress]
// See: https://developer.apple.com/documentation/Virtualization/_VZMemory
type VZMemory struct {
	objectivec.Object
}

// VZMemoryFromID constructs a [VZMemory] from an objc.ID.
func VZMemoryFromID(id objc.ID) VZMemory {
	return VZMemory{objectivec.Object{ID: id}}
}
// Ensure VZMemory implements IVZMemory.
var _ IVZMemory = VZMemory{}

// An interface definition for the [VZMemory] class.
//
// # Methods
//
//   - [IVZMemory.Length]
//   - [IVZMemory.MutableBytes]
//   - [IVZMemory.PhysicalAddress]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMemory
type IVZMemory interface {
	objectivec.IObject

	// Topic: Methods

	Length() uint64
	MutableBytes() unsafe.Pointer
	PhysicalAddress() uint64
}

// Init initializes the instance.
func (v VZMemory) Init() VZMemory {
	rv := objc.Send[VZMemory](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMemory) Autorelease() VZMemory {
	rv := objc.Send[VZMemory](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMemory creates a new VZMemory instance.
func NewVZMemory() VZMemory {
	class := getVZMemoryClass()
	rv := objc.Send[VZMemory](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMemory/length
func (v VZMemory) Length() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("length"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZMemory/mutableBytes
func (v VZMemory) MutableBytes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("mutableBytes"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZMemory/physicalAddress
func (v VZMemory) PhysicalAddress() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("physicalAddress"))
	return rv
}


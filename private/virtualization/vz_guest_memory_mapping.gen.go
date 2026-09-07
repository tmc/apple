// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZGuestMemoryMapping] class.
var (
	_VZGuestMemoryMappingClass     VZGuestMemoryMappingClass
	_VZGuestMemoryMappingClassOnce sync.Once
)

func getVZGuestMemoryMappingClass() VZGuestMemoryMappingClass {
	_VZGuestMemoryMappingClassOnce.Do(func() {
		_VZGuestMemoryMappingClass = VZGuestMemoryMappingClass{class: objc.GetClass("VZGuestMemoryMapping")}
	})
	return _VZGuestMemoryMappingClass
}

// GetVZGuestMemoryMappingClass returns the class object for VZGuestMemoryMapping.
func GetVZGuestMemoryMappingClass() VZGuestMemoryMappingClass {
	return getVZGuestMemoryMappingClass()
}

type VZGuestMemoryMappingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZGuestMemoryMappingClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZGuestMemoryMappingClass) Alloc() VZGuestMemoryMapping {
	rv := objc.SendIfResponds[VZGuestMemoryMapping](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZGuestMemoryMapping.Length]
//   - [VZGuestMemoryMapping.MutableBytes]
//   - [VZGuestMemoryMapping.PhysicalAddress]
type VZGuestMemoryMapping struct {
	objectivec.Object
}

// VZGuestMemoryMappingFromID constructs a [VZGuestMemoryMapping] from an objc.ID.
func VZGuestMemoryMappingFromID(id objc.ID) VZGuestMemoryMapping {
	return VZGuestMemoryMapping{objectivec.Object{ID: id}}
}

// Ensure VZGuestMemoryMapping implements IVZGuestMemoryMapping.
var _ IVZGuestMemoryMapping = VZGuestMemoryMapping{}

// An interface definition for the [VZGuestMemoryMapping] class.
//
// # Methods
//
//   - [IVZGuestMemoryMapping.Length]
//   - [IVZGuestMemoryMapping.MutableBytes]
//   - [IVZGuestMemoryMapping.PhysicalAddress]
type IVZGuestMemoryMapping interface {
	objectivec.IObject

	// Topic: Methods

	Length() uint64
	MutableBytes() unsafe.Pointer
	PhysicalAddress() uint64
}

// Init initializes the instance.
func (v VZGuestMemoryMapping) Init() VZGuestMemoryMapping {
	rv := objc.SendIfResponds[VZGuestMemoryMapping](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZGuestMemoryMapping) Autorelease() VZGuestMemoryMapping {
	rv := objc.SendIfResponds[VZGuestMemoryMapping](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGuestMemoryMapping creates a new VZGuestMemoryMapping instance.
func NewVZGuestMemoryMapping() VZGuestMemoryMapping {
	class := getVZGuestMemoryMappingClass()
	rv := objc.SendIfResponds[VZGuestMemoryMapping](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZGuestMemoryMapping) Length() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("length"))
	return rv
}
func (v VZGuestMemoryMapping) MutableBytes() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](v.ID, objc.Sel("mutableBytes"))
	return rv
}
func (v VZGuestMemoryMapping) PhysicalAddress() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("physicalAddress"))
	return rv
}

// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VirtualizationVZDiskImageLayerDescriptor] class.
var (
	_VirtualizationVZDiskImageLayerDescriptorClass     VirtualizationVZDiskImageLayerDescriptorClass
	_VirtualizationVZDiskImageLayerDescriptorClassOnce sync.Once
)

func getVirtualizationVZDiskImageLayerDescriptorClass() VirtualizationVZDiskImageLayerDescriptorClass {
	_VirtualizationVZDiskImageLayerDescriptorClassOnce.Do(func() {
		_VirtualizationVZDiskImageLayerDescriptorClass = VirtualizationVZDiskImageLayerDescriptorClass{class: objc.GetClass("Virtualization._VZDiskImageLayerDescriptor")}
	})
	return _VirtualizationVZDiskImageLayerDescriptorClass
}

// GetVirtualizationVZDiskImageLayerDescriptorClass returns the class object for Virtualization._VZDiskImageLayerDescriptor.
func GetVirtualizationVZDiskImageLayerDescriptorClass() VirtualizationVZDiskImageLayerDescriptorClass {
	return getVirtualizationVZDiskImageLayerDescriptorClass()
}

type VirtualizationVZDiskImageLayerDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VirtualizationVZDiskImageLayerDescriptorClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VirtualizationVZDiskImageLayerDescriptorClass) Alloc() VirtualizationVZDiskImageLayerDescriptor {
	rv := objc.SendIfResponds[VirtualizationVZDiskImageLayerDescriptor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VirtualizationVZDiskImageLayerDescriptor.FileHandle]
//   - [VirtualizationVZDiskImageLayerDescriptor.IsCache]
//   - [VirtualizationVZDiskImageLayerDescriptor.ReadOnly]
type VirtualizationVZDiskImageLayerDescriptor struct {
	objectivec.Object
}

// VirtualizationVZDiskImageLayerDescriptorFromID constructs a [VirtualizationVZDiskImageLayerDescriptor] from an objc.ID.
func VirtualizationVZDiskImageLayerDescriptorFromID(id objc.ID) VirtualizationVZDiskImageLayerDescriptor {
	return VirtualizationVZDiskImageLayerDescriptor{objectivec.Object{ID: id}}
}

// Ensure VirtualizationVZDiskImageLayerDescriptor implements IVirtualizationVZDiskImageLayerDescriptor.
var _ IVirtualizationVZDiskImageLayerDescriptor = VirtualizationVZDiskImageLayerDescriptor{}

// An interface definition for the [VirtualizationVZDiskImageLayerDescriptor] class.
//
// # Methods
//
//   - [IVirtualizationVZDiskImageLayerDescriptor.FileHandle]
//   - [IVirtualizationVZDiskImageLayerDescriptor.IsCache]
//   - [IVirtualizationVZDiskImageLayerDescriptor.ReadOnly]
type IVirtualizationVZDiskImageLayerDescriptor interface {
	objectivec.IObject

	// Topic: Methods

	FileHandle() unsafe.Pointer
	IsCache() bool
	ReadOnly() bool
}

// Init initializes the instance.
func (v VirtualizationVZDiskImageLayerDescriptor) Init() VirtualizationVZDiskImageLayerDescriptor {
	rv := objc.SendIfResponds[VirtualizationVZDiskImageLayerDescriptor](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VirtualizationVZDiskImageLayerDescriptor) Autorelease() VirtualizationVZDiskImageLayerDescriptor {
	rv := objc.SendIfResponds[VirtualizationVZDiskImageLayerDescriptor](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVirtualizationVZDiskImageLayerDescriptor creates a new VirtualizationVZDiskImageLayerDescriptor instance.
func NewVirtualizationVZDiskImageLayerDescriptor() VirtualizationVZDiskImageLayerDescriptor {
	class := getVirtualizationVZDiskImageLayerDescriptorClass()
	rv := objc.SendIfResponds[VirtualizationVZDiskImageLayerDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VirtualizationVZDiskImageLayerDescriptor) FileHandle() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](v.ID, objc.Sel("fileHandle"))
	return rv
}
func (v VirtualizationVZDiskImageLayerDescriptor) IsCache() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("isCache"))
	return rv
}
func (v VirtualizationVZDiskImageLayerDescriptor) ReadOnly() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("readOnly"))
	return rv
}

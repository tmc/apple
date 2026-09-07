// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VirtualizationVZDiskImageWrapper] class.
var (
	_VirtualizationVZDiskImageWrapperClass     VirtualizationVZDiskImageWrapperClass
	_VirtualizationVZDiskImageWrapperClassOnce sync.Once
)

func getVirtualizationVZDiskImageWrapperClass() VirtualizationVZDiskImageWrapperClass {
	_VirtualizationVZDiskImageWrapperClassOnce.Do(func() {
		_VirtualizationVZDiskImageWrapperClass = VirtualizationVZDiskImageWrapperClass{class: objc.GetClass("Virtualization._VZDiskImageWrapper")}
	})
	return _VirtualizationVZDiskImageWrapperClass
}

// GetVirtualizationVZDiskImageWrapperClass returns the class object for Virtualization._VZDiskImageWrapper.
func GetVirtualizationVZDiskImageWrapperClass() VirtualizationVZDiskImageWrapperClass {
	return getVirtualizationVZDiskImageWrapperClass()
}

type VirtualizationVZDiskImageWrapperClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VirtualizationVZDiskImageWrapperClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VirtualizationVZDiskImageWrapperClass) Alloc() VirtualizationVZDiskImageWrapper {
	rv := objc.SendIfResponds[VirtualizationVZDiskImageWrapper](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VirtualizationVZDiskImageWrapper.LayerDescriptors]
//   - [VirtualizationVZDiskImageWrapper.ReadOnly]
//   - [VirtualizationVZDiskImageWrapper.Url]
//   - [VirtualizationVZDiskImageWrapper.InitWithDiskImageError]
type VirtualizationVZDiskImageWrapper struct {
	objectivec.Object
}

// VirtualizationVZDiskImageWrapperFromID constructs a [VirtualizationVZDiskImageWrapper] from an objc.ID.
func VirtualizationVZDiskImageWrapperFromID(id objc.ID) VirtualizationVZDiskImageWrapper {
	return VirtualizationVZDiskImageWrapper{objectivec.Object{ID: id}}
}

// Ensure VirtualizationVZDiskImageWrapper implements IVirtualizationVZDiskImageWrapper.
var _ IVirtualizationVZDiskImageWrapper = VirtualizationVZDiskImageWrapper{}

// An interface definition for the [VirtualizationVZDiskImageWrapper] class.
//
// # Methods
//
//   - [IVirtualizationVZDiskImageWrapper.LayerDescriptors]
//   - [IVirtualizationVZDiskImageWrapper.ReadOnly]
//   - [IVirtualizationVZDiskImageWrapper.Url]
//   - [IVirtualizationVZDiskImageWrapper.InitWithDiskImageError]
type IVirtualizationVZDiskImageWrapper interface {
	objectivec.IObject

	// Topic: Methods

	LayerDescriptors() foundation.INSArray
	ReadOnly() bool
	Url() foundation.NSURL
	InitWithDiskImageError(image objectivec.IObject, error_ *int64) VirtualizationVZDiskImageWrapper
}

// Init initializes the instance.
func (v VirtualizationVZDiskImageWrapper) Init() VirtualizationVZDiskImageWrapper {
	rv := objc.SendIfResponds[VirtualizationVZDiskImageWrapper](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VirtualizationVZDiskImageWrapper) Autorelease() VirtualizationVZDiskImageWrapper {
	rv := objc.SendIfResponds[VirtualizationVZDiskImageWrapper](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVirtualizationVZDiskImageWrapper creates a new VirtualizationVZDiskImageWrapper instance.
func NewVirtualizationVZDiskImageWrapper() VirtualizationVZDiskImageWrapper {
	class := getVirtualizationVZDiskImageWrapperClass()
	rv := objc.SendIfResponds[VirtualizationVZDiskImageWrapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVirtualizationVZDiskImageWrapperWithDiskImageError(image objectivec.IObject, error_ *int64) VirtualizationVZDiskImageWrapper {
	instance := getVirtualizationVZDiskImageWrapperClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDiskImage:error:"), image, error_)
	return VirtualizationVZDiskImageWrapperFromID(rv)
}

func (v VirtualizationVZDiskImageWrapper) InitWithDiskImageError(image objectivec.IObject, error_ *int64) VirtualizationVZDiskImageWrapper {
	rv := objc.SendIfResponds[VirtualizationVZDiskImageWrapper](v.ID, objc.Sel("initWithDiskImage:error:"), image, error_)
	return rv
}

func (v VirtualizationVZDiskImageWrapper) LayerDescriptors() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("layerDescriptors"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VirtualizationVZDiskImageWrapper) ReadOnly() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("readOnly"))
	return rv
}
func (v VirtualizationVZDiskImageWrapper) Url() foundation.NSURL {
	rv := objc.SendIfResponds[foundation.NSURL](v.ID, objc.Sel("URL"))
	return foundation.NSURL(rv)
}

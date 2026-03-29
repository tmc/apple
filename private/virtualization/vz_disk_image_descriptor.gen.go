// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZDiskImageDescriptor] class.
var (
	_VZDiskImageDescriptorClass     VZDiskImageDescriptorClass
	_VZDiskImageDescriptorClassOnce sync.Once
)

func getVZDiskImageDescriptorClass() VZDiskImageDescriptorClass {
	_VZDiskImageDescriptorClassOnce.Do(func() {
		_VZDiskImageDescriptorClass = VZDiskImageDescriptorClass{class: objc.GetClass("_VZDiskImageDescriptor")}
	})
	return _VZDiskImageDescriptorClass
}

// GetVZDiskImageDescriptorClass returns the class object for _VZDiskImageDescriptor.
func GetVZDiskImageDescriptorClass() VZDiskImageDescriptorClass {
	return getVZDiskImageDescriptorClass()
}

type VZDiskImageDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZDiskImageDescriptorClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZDiskImageDescriptorClass) Alloc() VZDiskImageDescriptor {
	rv := objc.Send[VZDiskImageDescriptor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZDiskImageDescriptor.URL]
//   - [VZDiskImageDescriptor.SetURL]
//   - [VZDiskImageDescriptor.BackendType]
//   - [VZDiskImageDescriptor.CachingMode]
//   - [VZDiskImageDescriptor.SetCachingMode]
//   - [VZDiskImageDescriptor.IsReadOnly]
//   - [VZDiskImageDescriptor.SetBackendType]
//   - [VZDiskImageDescriptor.SynchronizationMode]
//   - [VZDiskImageDescriptor.SetSynchronizationMode]
//   - [VZDiskImageDescriptor.InitWithURL]
//   - [VZDiskImageDescriptor.ReadOnly]
//   - [VZDiskImageDescriptor.SetReadOnly]
// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageDescriptor
type VZDiskImageDescriptor struct {
	objectivec.Object
}

// VZDiskImageDescriptorFromID constructs a [VZDiskImageDescriptor] from an objc.ID.
func VZDiskImageDescriptorFromID(id objc.ID) VZDiskImageDescriptor {
	return VZDiskImageDescriptor{objectivec.Object{ID: id}}
}
// Ensure VZDiskImageDescriptor implements IVZDiskImageDescriptor.
var _ IVZDiskImageDescriptor = VZDiskImageDescriptor{}

// An interface definition for the [VZDiskImageDescriptor] class.
//
// # Methods
//
//   - [IVZDiskImageDescriptor.URL]
//   - [IVZDiskImageDescriptor.SetURL]
//   - [IVZDiskImageDescriptor.BackendType]
//   - [IVZDiskImageDescriptor.CachingMode]
//   - [IVZDiskImageDescriptor.SetCachingMode]
//   - [IVZDiskImageDescriptor.IsReadOnly]
//   - [IVZDiskImageDescriptor.SetBackendType]
//   - [IVZDiskImageDescriptor.SynchronizationMode]
//   - [IVZDiskImageDescriptor.SetSynchronizationMode]
//   - [IVZDiskImageDescriptor.InitWithURL]
//   - [IVZDiskImageDescriptor.ReadOnly]
//   - [IVZDiskImageDescriptor.SetReadOnly]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageDescriptor
type IVZDiskImageDescriptor interface {
	objectivec.IObject

	// Topic: Methods

	URL() foundation.INSURL
	SetURL(value foundation.INSURL)
	BackendType() int64
	CachingMode() int64
	SetCachingMode(value int64)
	IsReadOnly() bool
	SetBackendType(type_ int64)
	SynchronizationMode() int64
	SetSynchronizationMode(value int64)
	InitWithURL(url foundation.INSURL) VZDiskImageDescriptor
	ReadOnly() bool
	SetReadOnly(value bool)
}

// Init initializes the instance.
func (v VZDiskImageDescriptor) Init() VZDiskImageDescriptor {
	rv := objc.Send[VZDiskImageDescriptor](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZDiskImageDescriptor) Autorelease() VZDiskImageDescriptor {
	rv := objc.Send[VZDiskImageDescriptor](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDiskImageDescriptor creates a new VZDiskImageDescriptor instance.
func NewVZDiskImageDescriptor() VZDiskImageDescriptor {
	class := getVZDiskImageDescriptorClass()
	rv := objc.Send[VZDiskImageDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageDescriptor/initWithURL:
func NewVZDiskImageDescriptorWithURL(url foundation.INSURL) VZDiskImageDescriptor {
	instance := getVZDiskImageDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return VZDiskImageDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageDescriptor/backendType
func (v VZDiskImageDescriptor) BackendType() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("backendType"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageDescriptor/isReadOnly
func (v VZDiskImageDescriptor) IsReadOnly() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isReadOnly"))
	return rv
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageDescriptor/setBackendType:
func (v VZDiskImageDescriptor) SetBackendType(type_ int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("setBackendType:"), type_)
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageDescriptor/initWithURL:
func (v VZDiskImageDescriptor) InitWithURL(url foundation.INSURL) VZDiskImageDescriptor {
	rv := objc.Send[VZDiskImageDescriptor](v.ID, objc.Sel("initWithURL:"), url)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageDescriptor/URL
func (v VZDiskImageDescriptor) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (v VZDiskImageDescriptor) SetURL(value foundation.INSURL) {
	objc.Send[struct{}](v.ID, objc.Sel("setURL:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageDescriptor/cachingMode
func (v VZDiskImageDescriptor) CachingMode() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("cachingMode"))
	return rv
}
func (v VZDiskImageDescriptor) SetCachingMode(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("setCachingMode:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageDescriptor/readOnly
func (v VZDiskImageDescriptor) ReadOnly() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("readOnly"))
	return rv
}
func (v VZDiskImageDescriptor) SetReadOnly(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setReadOnly:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageDescriptor/synchronizationMode
func (v VZDiskImageDescriptor) SynchronizationMode() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("synchronizationMode"))
	return rv
}
func (v VZDiskImageDescriptor) SetSynchronizationMode(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("setSynchronizationMode:"), value)
}


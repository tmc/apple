// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSocketSerialPortAttachment] class.
var (
	_VZSocketSerialPortAttachmentClass     VZSocketSerialPortAttachmentClass
	_VZSocketSerialPortAttachmentClassOnce sync.Once
)

func getVZSocketSerialPortAttachmentClass() VZSocketSerialPortAttachmentClass {
	_VZSocketSerialPortAttachmentClassOnce.Do(func() {
		_VZSocketSerialPortAttachmentClass = VZSocketSerialPortAttachmentClass{class: objc.GetClass("_VZSocketSerialPortAttachment")}
	})
	return _VZSocketSerialPortAttachmentClass
}

// GetVZSocketSerialPortAttachmentClass returns the class object for _VZSocketSerialPortAttachment.
func GetVZSocketSerialPortAttachmentClass() VZSocketSerialPortAttachmentClass {
	return getVZSocketSerialPortAttachmentClass()
}

type VZSocketSerialPortAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSocketSerialPortAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSocketSerialPortAttachmentClass) Alloc() VZSocketSerialPortAttachment {
	rv := objc.SendIfResponds[VZSocketSerialPortAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZSocketSerialPortAttachment.Address]
//   - [VZSocketSerialPortAttachment.Mode]
//   - [VZSocketSerialPortAttachment.UnixSocketAddress]
//   - [VZSocketSerialPortAttachment.InitWithModeAddress]
//   - [VZSocketSerialPortAttachment.InitWithModeUnixSocketAddress]
type VZSocketSerialPortAttachment struct {
	VZSerialPortAttachment
}

// VZSocketSerialPortAttachmentFromID constructs a [VZSocketSerialPortAttachment] from an objc.ID.
func VZSocketSerialPortAttachmentFromID(id objc.ID) VZSocketSerialPortAttachment {
	return VZSocketSerialPortAttachment{VZSerialPortAttachment: VZSerialPortAttachmentFromID(id)}
}

// Ensure VZSocketSerialPortAttachment implements IVZSocketSerialPortAttachment.
var _ IVZSocketSerialPortAttachment = VZSocketSerialPortAttachment{}

// An interface definition for the [VZSocketSerialPortAttachment] class.
//
// # Methods
//
//   - [IVZSocketSerialPortAttachment.Address]
//   - [IVZSocketSerialPortAttachment.Mode]
//   - [IVZSocketSerialPortAttachment.UnixSocketAddress]
//   - [IVZSocketSerialPortAttachment.InitWithModeAddress]
//   - [IVZSocketSerialPortAttachment.InitWithModeUnixSocketAddress]
type IVZSocketSerialPortAttachment interface {
	IVZSerialPortAttachment

	// Topic: Methods

	Address() objectivec.IObject
	Mode() int64
	UnixSocketAddress() objectivec.IObject
	InitWithModeAddress(mode int64, address *Sockaddr) VZSocketSerialPortAttachment
	InitWithModeUnixSocketAddress(mode int64, address *SockaddrUn) VZSocketSerialPortAttachment
}

// Init initializes the instance.
func (v VZSocketSerialPortAttachment) Init() VZSocketSerialPortAttachment {
	rv := objc.SendIfResponds[VZSocketSerialPortAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZSocketSerialPortAttachment) Autorelease() VZSocketSerialPortAttachment {
	rv := objc.SendIfResponds[VZSocketSerialPortAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSocketSerialPortAttachment creates a new VZSocketSerialPortAttachment instance.
func NewVZSocketSerialPortAttachment() VZSocketSerialPortAttachment {
	class := getVZSocketSerialPortAttachmentClass()
	rv := objc.SendIfResponds[VZSocketSerialPortAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZSocketSerialPortAttachmentWithModeAddress(mode int64, address *Sockaddr) VZSocketSerialPortAttachment {
	instance := getVZSocketSerialPortAttachmentClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithMode:address:"), mode, address)
	return VZSocketSerialPortAttachmentFromID(rv)
}

func NewVZSocketSerialPortAttachmentWithModeUnixSocketAddress(mode int64, address *SockaddrUn) VZSocketSerialPortAttachment {
	instance := getVZSocketSerialPortAttachmentClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithMode:unixSocketAddress:"), mode, unsafe.Pointer(address))
	return VZSocketSerialPortAttachmentFromID(rv)
}

func (v VZSocketSerialPortAttachment) InitWithModeAddress(mode int64, address *Sockaddr) VZSocketSerialPortAttachment {
	rv := objc.SendIfResponds[VZSocketSerialPortAttachment](v.ID, objc.Sel("initWithMode:address:"), mode, address)
	return rv
}
func (v VZSocketSerialPortAttachment) InitWithModeUnixSocketAddress(mode int64, address *SockaddrUn) VZSocketSerialPortAttachment {
	rv := objc.SendIfResponds[VZSocketSerialPortAttachment](v.ID, objc.Sel("initWithMode:unixSocketAddress:"), mode, unsafe.Pointer(address))
	return rv
}

func (v VZSocketSerialPortAttachment) Address() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("address"))
	return objectivec.Object{ID: rv}
}
func (v VZSocketSerialPortAttachment) Mode() int64 {
	rv := objc.SendIfResponds[int64](v.ID, objc.Sel("mode"))
	return rv
}
func (v VZSocketSerialPortAttachment) UnixSocketAddress() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("unixSocketAddress"))
	return objectivec.Object{ID: rv}
}

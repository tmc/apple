// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/kernel"
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
	rv := objc.Send[VZSocketSerialPortAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZSocketSerialPortAttachment.Address]
//   - [VZSocketSerialPortAttachment.EncodeWithEncoder]
//   - [VZSocketSerialPortAttachment.Mode]
//   - [VZSocketSerialPortAttachment.UnixSocketAddress]
//   - [VZSocketSerialPortAttachment.InitWithModeAddress]
//   - [VZSocketSerialPortAttachment.InitWithModeUnixSocketAddress]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZSocketSerialPortAttachment
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
//   - [IVZSocketSerialPortAttachment.EncodeWithEncoder]
//   - [IVZSocketSerialPortAttachment.Mode]
//   - [IVZSocketSerialPortAttachment.UnixSocketAddress]
//   - [IVZSocketSerialPortAttachment.InitWithModeAddress]
//   - [IVZSocketSerialPortAttachment.InitWithModeUnixSocketAddress]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZSocketSerialPortAttachment
type IVZSocketSerialPortAttachment interface {
	IVZSerialPortAttachment

	// Topic: Methods

	Address() objectivec.IObject
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	Mode() int64
	UnixSocketAddress() objectivec.IObject
	InitWithModeAddress(mode int64, address unsafe.Pointer) VZSocketSerialPortAttachment
	InitWithModeUnixSocketAddress(mode int64, address *kernel.Sockaddr_un) VZSocketSerialPortAttachment
}

// Init initializes the instance.
func (v VZSocketSerialPortAttachment) Init() VZSocketSerialPortAttachment {
	rv := objc.Send[VZSocketSerialPortAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZSocketSerialPortAttachment) Autorelease() VZSocketSerialPortAttachment {
	rv := objc.Send[VZSocketSerialPortAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSocketSerialPortAttachment creates a new VZSocketSerialPortAttachment instance.
func NewVZSocketSerialPortAttachment() VZSocketSerialPortAttachment {
	class := getVZSocketSerialPortAttachmentClass()
	rv := objc.Send[VZSocketSerialPortAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSocketSerialPortAttachment/initWithMode:address:
func NewVZSocketSerialPortAttachmentWithModeAddress(mode int64, address unsafe.Pointer) VZSocketSerialPortAttachment {
	instance := getVZSocketSerialPortAttachmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMode:address:"), mode, address)
	return VZSocketSerialPortAttachmentFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSocketSerialPortAttachment/initWithMode:unixSocketAddress:
func NewVZSocketSerialPortAttachmentWithModeUnixSocketAddress(mode int64, address *kernel.Sockaddr_un) VZSocketSerialPortAttachment {
	instance := getVZSocketSerialPortAttachmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMode:unixSocketAddress:"), mode, address)
	return VZSocketSerialPortAttachmentFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSocketSerialPortAttachment/encodeWithEncoder:
func (v VZSocketSerialPortAttachment) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSocketSerialPortAttachment/initWithMode:address:
func (v VZSocketSerialPortAttachment) InitWithModeAddress(mode int64, address unsafe.Pointer) VZSocketSerialPortAttachment {
	rv := objc.Send[VZSocketSerialPortAttachment](v.ID, objc.Sel("initWithMode:address:"), mode, address)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSocketSerialPortAttachment/initWithMode:unixSocketAddress:
func (v VZSocketSerialPortAttachment) InitWithModeUnixSocketAddress(mode int64, address *kernel.Sockaddr_un) VZSocketSerialPortAttachment {
	rv := objc.Send[VZSocketSerialPortAttachment](v.ID, objc.Sel("initWithMode:unixSocketAddress:"), mode, address)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSocketSerialPortAttachment/address
func (v VZSocketSerialPortAttachment) Address() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("address"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSocketSerialPortAttachment/mode
func (v VZSocketSerialPortAttachment) Mode() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("mode"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSocketSerialPortAttachment/unixSocketAddress
func (v VZSocketSerialPortAttachment) UnixSocketAddress() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("unixSocketAddress"))
	return objectivec.Object{ID: rv}
}

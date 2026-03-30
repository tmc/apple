// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSerialPort] class.
var (
	_VZSerialPortClass     VZSerialPortClass
	_VZSerialPortClassOnce sync.Once
)

func getVZSerialPortClass() VZSerialPortClass {
	_VZSerialPortClassOnce.Do(func() {
		_VZSerialPortClass = VZSerialPortClass{class: objc.GetClass("_VZSerialPort")}
	})
	return _VZSerialPortClass
}

// GetVZSerialPortClass returns the class object for _VZSerialPort.
func GetVZSerialPortClass() VZSerialPortClass {
	return getVZSerialPortClass()
}

type VZSerialPortClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSerialPortClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSerialPortClass) Alloc() VZSerialPort {
	rv := objc.Send[VZSerialPort](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZSerialPort.Attachment]
//   - [VZSerialPort.SetAttachment]
//   - [VZSerialPort.Type]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZSerialPort
type VZSerialPort struct {
	objectivec.Object
}

// VZSerialPortFromID constructs a [VZSerialPort] from an objc.ID.
func VZSerialPortFromID(id objc.ID) VZSerialPort {
	return VZSerialPort{objectivec.Object{ID: id}}
}

// Ensure VZSerialPort implements IVZSerialPort.
var _ IVZSerialPort = VZSerialPort{}

// An interface definition for the [VZSerialPort] class.
//
// # Methods
//
//   - [IVZSerialPort.Attachment]
//   - [IVZSerialPort.SetAttachment]
//   - [IVZSerialPort.Type]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZSerialPort
type IVZSerialPort interface {
	objectivec.IObject

	// Topic: Methods

	Attachment() IVZSerialPortAttachment
	SetAttachment(value IVZSerialPortAttachment)
	Type() int64
}

// Init initializes the instance.
func (v VZSerialPort) Init() VZSerialPort {
	rv := objc.Send[VZSerialPort](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZSerialPort) Autorelease() VZSerialPort {
	rv := objc.Send[VZSerialPort](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSerialPort creates a new VZSerialPort instance.
func NewVZSerialPort() VZSerialPort {
	class := getVZSerialPortClass()
	rv := objc.Send[VZSerialPort](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSerialPort/attachment
func (v VZSerialPort) Attachment() IVZSerialPortAttachment {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("attachment"))
	return VZSerialPortAttachmentFromID(objc.ID(rv))
}
func (v VZSerialPort) SetAttachment(value IVZSerialPortAttachment) {
	objc.Send[struct{}](v.ID, objc.Sel("setAttachment:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSerialPort/type
func (v VZSerialPort) Type() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("type"))
	return rv
}

// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioConsolePort] class.
var (
	_VZVirtioConsolePortClass     VZVirtioConsolePortClass
	_VZVirtioConsolePortClassOnce sync.Once
)

func getVZVirtioConsolePortClass() VZVirtioConsolePortClass {
	_VZVirtioConsolePortClassOnce.Do(func() {
		_VZVirtioConsolePortClass = VZVirtioConsolePortClass{class: objc.GetClass("VZVirtioConsolePort")}
	})
	return _VZVirtioConsolePortClass
}

// GetVZVirtioConsolePortClass returns the class object for VZVirtioConsolePort.
func GetVZVirtioConsolePortClass() VZVirtioConsolePortClass {
	return getVZVirtioConsolePortClass()
}

type VZVirtioConsolePortClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioConsolePortClass) Alloc() VZVirtioConsolePort {
	rv := objc.Send[VZVirtioConsolePort](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// A class that represents a Virtio console port in a VM.
//
// # Overview
// 
// Don’t instantiate a [VZVirtioConsolePort] directly. You retrieve this
// object from the [VZVirtioConsoleDevice] [VZVirtioConsolePort.Ports] property.
//
// # Configuring the port
//
//   - [VZVirtioConsolePort.Name]: The name of the port.
//   - [VZVirtioConsolePort.Attachment]: An array of serial port attachments.
//   - [VZVirtioConsolePort.SetAttachment]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePort
type VZVirtioConsolePort struct {
	objectivec.Object
}

// VZVirtioConsolePortFromID constructs a [VZVirtioConsolePort] from an objc.ID.
//
// A class that represents a Virtio console port in a VM.
func VZVirtioConsolePortFromID(id objc.ID) VZVirtioConsolePort {
	return VZVirtioConsolePort{objectivec.Object{id}}
}
// NOTE: VZVirtioConsolePort adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZVirtioConsolePort] class.
//
// # Configuring the port
//
//   - [IVZVirtioConsolePort.Name]: The name of the port.
//   - [IVZVirtioConsolePort.Attachment]: An array of serial port attachments.
//   - [IVZVirtioConsolePort.SetAttachment]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePort
type IVZVirtioConsolePort interface {
	objectivec.IObject

	// Topic: Configuring the port

	// The name of the port.
	Name() string
	// An array of serial port attachments.
	Attachment() IVZSerialPortAttachment
	SetAttachment(value IVZSerialPortAttachment)

	// The array of console ports that a specific device uses.
	Ports() IVZVirtioConsolePortArray
	SetPorts(value IVZVirtioConsolePortArray)
}





// Init initializes the instance.
func (v VZVirtioConsolePort) Init() VZVirtioConsolePort {
	rv := objc.Send[VZVirtioConsolePort](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioConsolePort) Autorelease() VZVirtioConsolePort {
	rv := objc.Send[VZVirtioConsolePort](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsolePort creates a new VZVirtioConsolePort instance.
func NewVZVirtioConsolePort() VZVirtioConsolePort {
	class := getVZVirtioConsolePortClass()
	rv := objc.Send[VZVirtioConsolePort](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// The name of the port.
//
// # Discussion
// 
// This property can’t change while the VM is running. A null value
// indicates a name isn’t set.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePort/name
func (v VZVirtioConsolePort) Name() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}



// An array of serial port attachments.
//
// # Discussion
// 
// This property may change at any time while the VM is running.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePort/attachment
func (v VZVirtioConsolePort) Attachment() IVZSerialPortAttachment {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("attachment"))
	return VZSerialPortAttachmentFromID(objc.ID(rv))
}
func (v VZVirtioConsolePort) SetAttachment(value IVZSerialPortAttachment) {
	objc.Send[struct{}](v.ID, objc.Sel("setAttachment:"), value)
}



// The array of console ports that a specific device uses.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtioconsoledevice/ports
func (v VZVirtioConsolePort) Ports() IVZVirtioConsolePortArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("ports"))
	return VZVirtioConsolePortArrayFromID(objc.ID(rv))
}
func (v VZVirtioConsolePort) SetPorts(value IVZVirtioConsolePortArray) {
	objc.Send[struct{}](v.ID, objc.Sel("setPorts:"), value)
}
























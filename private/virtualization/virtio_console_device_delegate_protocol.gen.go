// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// VZVirtioConsoleDeviceDelegate protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceDelegate
type VZVirtioConsoleDeviceDelegate interface {
	objectivec.IObject
}

// VZVirtioConsoleDeviceDelegateObject wraps an existing Objective-C object that conforms to the VZVirtioConsoleDeviceDelegate protocol.
type VZVirtioConsoleDeviceDelegateObject struct {
	objectivec.Object
}
func (o VZVirtioConsoleDeviceDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZVirtioConsoleDeviceDelegateObjectFromID constructs a [VZVirtioConsoleDeviceDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZVirtioConsoleDeviceDelegateObjectFromID(id objc.ID) VZVirtioConsoleDeviceDelegateObject {
	return VZVirtioConsoleDeviceDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceDelegate/consoleDevice:didClosePort:
func (o VZVirtioConsoleDeviceDelegateObject) ConsoleDeviceDidClosePort(device objectivec.IObject, port objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("consoleDevice:didClosePort:"), device, port)
	}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceDelegate/consoleDevice:didOpenPort:
func (o VZVirtioConsoleDeviceDelegateObject) ConsoleDeviceDidOpenPort(device objectivec.IObject, port objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("consoleDevice:didOpenPort:"), device, port)
	}


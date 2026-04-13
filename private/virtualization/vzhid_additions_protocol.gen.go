// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZHIDAdditions protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZHIDAdditions
type VZHIDAdditions interface {
	objectivec.IObject

	// _processHIDReportsForDeviceDeviceType protocol.
	//
	// See: https://developer.apple.com/documentation/Virtualization/_VZHIDAdditions/_processHIDReports:forDevice:deviceType:
	_processHIDReportsForDeviceDeviceType(hIDReports VZOpaqueHIDReports, device uint32, type_ int32)

	// _shouldSendHIDReports protocol.
	//
	// See: https://developer.apple.com/documentation/Virtualization/_VZHIDAdditions/_shouldSendHIDReports
	_shouldSendHIDReports() bool
}

// VZHIDAdditionsObject wraps an existing Objective-C object that conforms to the VZHIDAdditions protocol.
type VZHIDAdditionsObject struct {
	objectivec.Object
}

func (o VZHIDAdditionsObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZHIDAdditionsObjectFromID constructs a [VZHIDAdditionsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZHIDAdditionsObjectFromID(id objc.ID) VZHIDAdditionsObject {
	return VZHIDAdditionsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZHIDAdditions/_hidEventMonitor
func (o VZHIDAdditionsObject) _hidEventMonitor() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("_hidEventMonitor"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZHIDAdditions/_processHIDReports:forDevice:deviceType:
func (o VZHIDAdditionsObject) _processHIDReportsForDeviceDeviceType(hIDReports VZOpaqueHIDReports, device uint32, type_ int32) {
	objc.Send[struct{}](o.ID, objc.Sel("_processHIDReports:forDevice:deviceType:"), hIDReports.UnsafePointer(), device, type_)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZHIDAdditions/_shouldSendHIDReports
func (o VZHIDAdditionsObject) _shouldSendHIDReports() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("_shouldSendHIDReports"))
	return rv
}

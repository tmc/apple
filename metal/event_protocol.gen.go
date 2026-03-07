// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A type that synchronizes memory operations to one or more resources within a single Metal device.
//
// See: https://developer.apple.com/documentation/Metal/MTLEvent
type MTLEvent interface {
	objectivec.IObject

	// The device object that created the event.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLEvent/device
	Device() MTLDevice

	// A string that identifies the event.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLEvent/label
	Label() string

	// A string that identifies the event.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLEvent/label
	SetLabel(value string)
}



// MTLEventObject wraps an existing Objective-C object that conforms to the MTLEvent protocol.
type MTLEventObject struct {
	objectivec.Object
}
func (o MTLEventObject) BaseObject() objectivec.Object {
	return o.Object
}



// MTLEventObjectFromID constructs a [MTLEventObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLEventObjectFromID(id objc.ID) MTLEventObject {
	return MTLEventObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// The device object that created the event.
//
// See: https://developer.apple.com/documentation/Metal/MTLEvent/device

func (o MTLEventObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}

// A string that identifies the event.
//
// See: https://developer.apple.com/documentation/Metal/MTLEvent/label

func (o MTLEventObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}






func (o MTLEventObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}






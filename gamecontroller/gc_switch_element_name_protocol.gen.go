// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The name for an element that represents a switch.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchElementName-c.protocol
type GCSwitchElementName interface {
	objectivec.IObject
	GCPhysicalInputElementName
}

// GCSwitchElementNameObject wraps an existing Objective-C object that conforms to the GCSwitchElementName protocol.
type GCSwitchElementNameObject struct {
	objectivec.Object
}

func (o GCSwitchElementNameObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCSwitchElementNameObjectFromID constructs a [GCSwitchElementNameObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCSwitchElementNameObjectFromID(id objc.ID) GCSwitchElementNameObject {
	return GCSwitchElementNameObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Code generated from Apple documentation for Security. DO NOT EDIT.

package security

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A `sec_protocol_metadata` instance conatins read-only properties of a connected and configured security protocol. Clients use this object to read information about a protocol instance. Properties include, for example, the negotiated TLS version, ciphersuite, and peer certificates.
//
// See: https://developer.apple.com/documentation/Security/OS_sec_protocol_metadata
type OS_sec_protocol_metadata interface {
	objectivec.IObject
}

// OS_sec_protocol_metadataObject wraps an existing Objective-C object that conforms to the OS_sec_protocol_metadata protocol.
type OS_sec_protocol_metadataObject struct {
	objectivec.Object
}

func (o OS_sec_protocol_metadataObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_sec_protocol_metadataObjectFromID constructs a [OS_sec_protocol_metadataObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_sec_protocol_metadataObjectFromID(id objc.ID) OS_sec_protocol_metadataObject {
	return OS_sec_protocol_metadataObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Code generated from Apple documentation for Security. DO NOT EDIT.

package security

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A `sec_protocol_options` instance is a container of options for security protocol instances, such as TLS. Protocol options are used to configure security protocols in the network stack. For example, clients may set the maximum and minimum allowed TLS versions through protocol options.
//
// See: https://developer.apple.com/documentation/Security/OS_sec_protocol_options
type OS_sec_protocol_options interface {
	objectivec.IObject
}

// OS_sec_protocol_optionsObject wraps an existing Objective-C object that conforms to the OS_sec_protocol_options protocol.
type OS_sec_protocol_optionsObject struct {
	objectivec.Object
}

func (o OS_sec_protocol_optionsObject) BaseObject() objectivec.Object {
	return o.Object
}

// OS_sec_protocol_optionsObjectFromID constructs a [OS_sec_protocol_optionsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OS_sec_protocol_optionsObjectFromID(id objc.ID) OS_sec_protocol_optionsObject {
	return OS_sec_protocol_optionsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol defining subclasses that represent entries made using a handle and a format string.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryWithPayload
type OSLogEntryWithPayload interface {
	objectivec.IObject

	// The payload’s category.
	//
	// See: https://developer.apple.com/documentation/OSLog/OSLogEntryWithPayload/category
	Category() string

	// The payload’s components.
	//
	// See: https://developer.apple.com/documentation/OSLog/OSLogEntryWithPayload/components
	Components() []OSLogMessageComponent

	// The payload’s format string.
	//
	// See: https://developer.apple.com/documentation/OSLog/OSLogEntryWithPayload/formatString
	FormatString() string

	// The payload’s subsystem.
	//
	// See: https://developer.apple.com/documentation/OSLog/OSLogEntryWithPayload/subsystem
	Subsystem() string
}

// OSLogEntryWithPayloadObject wraps an existing Objective-C object that conforms to the OSLogEntryWithPayload protocol.
type OSLogEntryWithPayloadObject struct {
	objectivec.Object
}

func (o OSLogEntryWithPayloadObject) BaseObject() objectivec.Object {
	return o.Object
}

// OSLogEntryWithPayloadObjectFromID constructs a [OSLogEntryWithPayloadObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OSLogEntryWithPayloadObjectFromID(id objc.ID) OSLogEntryWithPayloadObject {
	return OSLogEntryWithPayloadObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The payload’s category.
//
// # Discussion
//
// The category is derived from the `os_log_t` handle used.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryWithPayload/category
func (o OSLogEntryWithPayloadObject) Category() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("category"))
	return foundation.NSStringFromID(rv).String()
}

// The payload’s components.
//
// # Discussion
//
// This property contains an array of [OSLogMessageComponent] objects from the
// composed message.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryWithPayload/components
func (o OSLogEntryWithPayloadObject) Components() []OSLogMessageComponent {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("components"))
	result := make([]OSLogMessageComponent, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = OSLogMessageComponentFromID(id)
	}
	return result
}

// The payload’s format string.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryWithPayload/formatString
func (o OSLogEntryWithPayloadObject) FormatString() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("formatString"))
	return foundation.NSStringFromID(rv).String()
}

// The payload’s subsystem.
//
// # Discussion
//
// The category is derived from the `os_log_t` handle used.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryWithPayload/subsystem
func (o OSLogEntryWithPayloadObject) Subsystem() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("subsystem"))
	return foundation.NSStringFromID(rv).String()
}

// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/os"
)

// A protocol that defines subclasses containing metadata about a process.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryFromProcess
type OSLogEntryFromProcess interface {
	objectivec.IObject

	// The activity identifier associated with the entry.
	//
	// See: https://developer.apple.com/documentation/OSLog/OSLogEntryFromProcess/activityIdentifier
	ActivityIdentifier() os.OSActivityID

	// The name of the process that made the entry.
	//
	// See: https://developer.apple.com/documentation/OSLog/OSLogEntryFromProcess/process
	Process() string

	// The process identifier that made the entry.
	//
	// See: https://developer.apple.com/documentation/OSLog/OSLogEntryFromProcess/processIdentifier
	ProcessIdentifier() int32

	// The name of the binary image that made the entry.
	//
	// See: https://developer.apple.com/documentation/OSLog/OSLogEntryFromProcess/sender
	Sender() string

	// The identifier of the thread that made the entry.
	//
	// See: https://developer.apple.com/documentation/OSLog/OSLogEntryFromProcess/threadIdentifier
	ThreadIdentifier() uint64
}

// OSLogEntryFromProcessObject wraps an existing Objective-C object that conforms to the OSLogEntryFromProcess protocol.
type OSLogEntryFromProcessObject struct {
	objectivec.Object
}

func (o OSLogEntryFromProcessObject) BaseObject() objectivec.Object {
	return o.Object
}

// OSLogEntryFromProcessObjectFromID constructs a [OSLogEntryFromProcessObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OSLogEntryFromProcessObjectFromID(id objc.ID) OSLogEntryFromProcessObject {
	return OSLogEntryFromProcessObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The activity identifier associated with the entry.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryFromProcess/activityIdentifier
func (o OSLogEntryFromProcessObject) ActivityIdentifier() os.OSActivityID {
	rv := objc.Send[os.OSActivityID](o.ID, objc.Sel("activityIdentifier"))
	return os.OSActivityID(rv)
}

// The name of the process that made the entry.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryFromProcess/process
func (o OSLogEntryFromProcessObject) Process() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("process"))
	return foundation.NSStringFromID(rv).String()
}

// The process identifier that made the entry.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryFromProcess/processIdentifier
func (o OSLogEntryFromProcessObject) ProcessIdentifier() int32 {
	rv := objc.Send[int32](o.ID, objc.Sel("processIdentifier"))
	return int32(rv)
}

// The name of the binary image that made the entry.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryFromProcess/sender
func (o OSLogEntryFromProcessObject) Sender() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sender"))
	return foundation.NSStringFromID(rv).String()
}

// The identifier of the thread that made the entry.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryFromProcess/threadIdentifier
func (o OSLogEntryFromProcessObject) ThreadIdentifier() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("threadIdentifier"))
	return uint64(rv)
}

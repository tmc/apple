// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/os"
)

// The class instance for the [OSLogEntrySignpost] class.
var (
	_OSLogEntrySignpostClass     OSLogEntrySignpostClass
	_OSLogEntrySignpostClassOnce sync.Once
)

func getOSLogEntrySignpostClass() OSLogEntrySignpostClass {
	_OSLogEntrySignpostClassOnce.Do(func() {
		_OSLogEntrySignpostClass = OSLogEntrySignpostClass{class: objc.GetClass("OSLogEntrySignpost")}
	})
	return _OSLogEntrySignpostClass
}

// GetOSLogEntrySignpostClass returns the class object for OSLogEntrySignpost.
func GetOSLogEntrySignpostClass() OSLogEntrySignpostClass {
	return getOSLogEntrySignpostClass()
}

type OSLogEntrySignpostClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OSLogEntrySignpostClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OSLogEntrySignpostClass) Alloc() OSLogEntrySignpost {
	rv := objc.Send[OSLogEntrySignpost](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// An entry containing a signpost.
//
// # Overview
//
// These entries are created by the os_signpost API. To learn more about
// signposts and how to create a signpost entry, see
// [os_signpost(_:dso:log:name:signpostID:)] and
// [os_signpost(_:dso:log:name:signpostID:_:_:)].
//
// # Accessing Signpost Details
//
//   - [OSLogEntrySignpost.SignpostIdentifier]: The signpost’s identifier.
//   - [OSLogEntrySignpost.SignpostName]: The signpost’s name.
//
// # Accessing Signpost Types
//
//   - [OSLogEntrySignpost.SignpostType]: The signpost’s type.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntrySignpost
//
// [os_signpost(_:dso:log:name:signpostID:)]: https://developer.apple.com/documentation/os/os_signpost(_:dso:log:name:signpostID:)-2oz8u
// [os_signpost(_:dso:log:name:signpostID:_:_:)]: https://developer.apple.com/documentation/os/os_signpost(_:dso:log:name:signpostID:_:_:)-2om9b
type OSLogEntrySignpost struct {
	OSLogEntry
}

// OSLogEntrySignpostFromID constructs a [OSLogEntrySignpost] from an objc.ID.
//
// An entry containing a signpost.
func OSLogEntrySignpostFromID(id objc.ID) OSLogEntrySignpost {
	return OSLogEntrySignpost{OSLogEntry: OSLogEntryFromID(id)}
}

// NOTE: OSLogEntrySignpost adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OSLogEntrySignpost] class.
//
// # Accessing Signpost Details
//
//   - [IOSLogEntrySignpost.SignpostIdentifier]: The signpost’s identifier.
//   - [IOSLogEntrySignpost.SignpostName]: The signpost’s name.
//
// # Accessing Signpost Types
//
//   - [IOSLogEntrySignpost.SignpostType]: The signpost’s type.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntrySignpost
type IOSLogEntrySignpost interface {
	IOSLogEntry

	// Topic: Accessing Signpost Details

	// The signpost’s identifier.
	SignpostIdentifier() os.OSSignpostID
	// The signpost’s name.
	SignpostName() string

	// Topic: Accessing Signpost Types

	// The signpost’s type.
	SignpostType() OSLogEntrySignpostType

	// The activity identifier associated with the entry.
	ActivityIdentifier() os.OSActivityID
	// The payload’s category.
	Category() string
	// The payload’s components.
	Components() []OSLogMessageComponent
	// The payload’s format string.
	FormatString() string
	// The name of the process that made the entry.
	Process() string
	// The process identifier that made the entry.
	ProcessIdentifier() int32
	// The name of the binary image that made the entry.
	Sender() string
	// The payload’s subsystem.
	Subsystem() string
	// The identifier of the thread that made the entry.
	ThreadIdentifier() uint64
}

// Init initializes the instance.
func (o OSLogEntrySignpost) Init() OSLogEntrySignpost {
	rv := objc.Send[OSLogEntrySignpost](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OSLogEntrySignpost) Autorelease() OSLogEntrySignpost {
	rv := objc.Send[OSLogEntrySignpost](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogEntrySignpost creates a new OSLogEntrySignpost instance.
func NewOSLogEntrySignpost() OSLogEntrySignpost {
	class := getOSLogEntrySignpostClass()
	rv := objc.Send[OSLogEntrySignpost](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The activity identifier associated with the entry.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryFromProcess/activityIdentifier
func (o OSLogEntrySignpost) ActivityIdentifier() os.OSActivityID {
	rv := objc.Send[os.OSActivityID](o.ID, objc.Sel("activityIdentifier"))
	return os.OSActivityID(rv)
}

// The payload’s category.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryWithPayload/category
func (o OSLogEntrySignpost) Category() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("category"))
	return foundation.NSStringFromID(rv).String()
}

// The payload’s components.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryWithPayload/components
func (o OSLogEntrySignpost) Components() []OSLogMessageComponent {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("components"))
	return objc.ConvertSlice(rv, func(id objc.ID) OSLogMessageComponent {
		return OSLogMessageComponentFromID(id)
	})
}

// The payload’s format string.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryWithPayload/formatString
func (o OSLogEntrySignpost) FormatString() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("formatString"))
	return foundation.NSStringFromID(rv).String()
}

// The name of the process that made the entry.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryFromProcess/process
func (o OSLogEntrySignpost) Process() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("process"))
	return foundation.NSStringFromID(rv).String()
}

// The process identifier that made the entry.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryFromProcess/processIdentifier
func (o OSLogEntrySignpost) ProcessIdentifier() int32 {
	rv := objc.Send[int32](o.ID, objc.Sel("processIdentifier"))
	return rv
}

// The name of the binary image that made the entry.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryFromProcess/sender
func (o OSLogEntrySignpost) Sender() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sender"))
	return foundation.NSStringFromID(rv).String()
}

// The payload’s subsystem.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryWithPayload/subsystem
func (o OSLogEntrySignpost) Subsystem() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("subsystem"))
	return foundation.NSStringFromID(rv).String()
}

// The identifier of the thread that made the entry.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntryFromProcess/threadIdentifier
func (o OSLogEntrySignpost) ThreadIdentifier() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("threadIdentifier"))
	return rv
}

// The signpost’s identifier.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntrySignpost/signpostIdentifier
func (o OSLogEntrySignpost) SignpostIdentifier() os.OSSignpostID {
	rv := objc.Send[os.OSSignpostID](o.ID, objc.Sel("signpostIdentifier"))
	return os.OSSignpostID(rv)
}

// The signpost’s name.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntrySignpost/signpostName
func (o OSLogEntrySignpost) SignpostName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("signpostName"))
	return foundation.NSStringFromID(rv).String()
}

// The signpost’s type.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogEntrySignpost/signpostType-swift.property
func (o OSLogEntrySignpost) SignpostType() OSLogEntrySignpostType {
	rv := objc.Send[OSLogEntrySignpostType](o.ID, objc.Sel("signpostType"))
	return OSLogEntrySignpostType(rv)
}

// Protocol methods for OSLogEntryFromProcess

// Protocol methods for OSLogEntryWithPayload

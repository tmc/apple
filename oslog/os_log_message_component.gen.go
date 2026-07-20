// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [OSLogMessageComponent] class.
var (
	_OSLogMessageComponentClass     OSLogMessageComponentClass
	_OSLogMessageComponentClassOnce sync.Once
)

func getOSLogMessageComponentClass() OSLogMessageComponentClass {
	_OSLogMessageComponentClassOnce.Do(func() {
		_OSLogMessageComponentClass = OSLogMessageComponentClass{class: objc.GetClass("OSLogMessageComponent")}
	})
	return _OSLogMessageComponentClass
}

// GetOSLogMessageComponentClass returns the class object for OSLogMessageComponent.
func GetOSLogMessageComponentClass() OSLogMessageComponentClass {
	return getOSLogMessageComponentClass()
}

type OSLogMessageComponentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OSLogMessageComponentClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OSLogMessageComponentClass) Alloc() OSLogMessageComponent {
	rv := objc.Send[OSLogMessageComponent](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// The message arguments for a particular entry.
//
// # Overview
//
// There is one component for each placeholder in the formatString plus one
// component for any text after the last placeholder.
//
// # Reading the Argument
//
//   - [OSLogMessageComponent.Argument]: The argument passed into the message component.
//   - [OSLogMessageComponent.SetArgument]
//   - [OSLogMessageComponent.ArgumentCategory]: The type of argument that corresponds to the placeholder.
//
// # Reading the Message Component
//
//   - [OSLogMessageComponent.FormatSubstring]: The text immediately preceding a placeholder.
//   - [OSLogMessageComponent.Placeholder]: The placeholder text for the message component.
//
// # Accessing the Argument
//
//   - [OSLogMessageComponent.ArgumentDataValue]: The argument formatted as a sequence of bytes.
//   - [OSLogMessageComponent.ArgumentDoubleValue]: The argument formatted as a double.
//   - [OSLogMessageComponent.ArgumentInt64Value]: The argument formatted as a signed 64-bit integer.
//   - [OSLogMessageComponent.ArgumentNumberValue]: The argument formatted as a number.
//   - [OSLogMessageComponent.ArgumentStringValue]: The argument formatted as a string.
//   - [OSLogMessageComponent.ArgumentUInt64Value]: The argument formatted as an unsigned 64-bit integer.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent
type OSLogMessageComponent struct {
	objectivec.Object
}

// OSLogMessageComponentFromID constructs a [OSLogMessageComponent] from an objc.ID.
//
// The message arguments for a particular entry.
func OSLogMessageComponentFromID(id objc.ID) OSLogMessageComponent {
	return OSLogMessageComponent{objectivec.Object{ID: id}}
}

// NOTE: OSLogMessageComponent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OSLogMessageComponent] class.
//
// # Reading the Argument
//
//   - [IOSLogMessageComponent.Argument]: The argument passed into the message component.
//   - [IOSLogMessageComponent.SetArgument]
//   - [IOSLogMessageComponent.ArgumentCategory]: The type of argument that corresponds to the placeholder.
//
// # Reading the Message Component
//
//   - [IOSLogMessageComponent.FormatSubstring]: The text immediately preceding a placeholder.
//   - [IOSLogMessageComponent.Placeholder]: The placeholder text for the message component.
//
// # Accessing the Argument
//
//   - [IOSLogMessageComponent.ArgumentDataValue]: The argument formatted as a sequence of bytes.
//   - [IOSLogMessageComponent.ArgumentDoubleValue]: The argument formatted as a double.
//   - [IOSLogMessageComponent.ArgumentInt64Value]: The argument formatted as a signed 64-bit integer.
//   - [IOSLogMessageComponent.ArgumentNumberValue]: The argument formatted as a number.
//   - [IOSLogMessageComponent.ArgumentStringValue]: The argument formatted as a string.
//   - [IOSLogMessageComponent.ArgumentUInt64Value]: The argument formatted as an unsigned 64-bit integer.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent
type IOSLogMessageComponent interface {
	objectivec.IObject

	// Topic: Reading the Argument

	// The argument passed into the message component.
	Argument() unsafe.Pointer
	SetArgument(value unsafe.Pointer)
	// The type of argument that corresponds to the placeholder.
	ArgumentCategory() OSLogMessageComponentArgumentCategory

	// Topic: Reading the Message Component

	// The text immediately preceding a placeholder.
	FormatSubstring() string
	// The placeholder text for the message component.
	Placeholder() string

	// Topic: Accessing the Argument

	// The argument formatted as a sequence of bytes.
	ArgumentDataValue() foundation.NSData
	// The argument formatted as a double.
	ArgumentDoubleValue() float64
	// The argument formatted as a signed 64-bit integer.
	ArgumentInt64Value() int64
	// The argument formatted as a number.
	ArgumentNumberValue() foundation.NSNumber
	// The argument formatted as a string.
	ArgumentStringValue() string
	// The argument formatted as an unsigned 64-bit integer.
	ArgumentUInt64Value() uint64

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (o OSLogMessageComponent) Init() OSLogMessageComponent {
	rv := objc.Send[OSLogMessageComponent](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OSLogMessageComponent) Autorelease() OSLogMessageComponent {
	rv := objc.Send[OSLogMessageComponent](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogMessageComponent creates a new OSLogMessageComponent instance.
func NewOSLogMessageComponent() OSLogMessageComponent {
	class := getOSLogMessageComponentClass()
	rv := objc.Send[OSLogMessageComponent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (o OSLogMessageComponent) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The argument passed into the message component.
//
// See: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argument-swift.property
func (o OSLogMessageComponent) Argument() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("argument"))
	return rv
}
func (o OSLogMessageComponent) SetArgument(value unsafe.Pointer) {
	objc.Send[struct{}](o.ID, objc.Sel("setArgument:"), value)
}

// The type of argument that corresponds to the placeholder.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentCategory-swift.property
func (o OSLogMessageComponent) ArgumentCategory() OSLogMessageComponentArgumentCategory {
	rv := objc.Send[OSLogMessageComponentArgumentCategory](o.ID, objc.Sel("argumentCategory"))
	return OSLogMessageComponentArgumentCategory(rv)
}

// The text immediately preceding a placeholder.
//
// # Discussion
//
// The `formatSubstring` property can be an empty string if there is nothing
// between two placeholders, or if it is between the placeholder and the
// bounds of the string.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/formatSubstring
func (o OSLogMessageComponent) FormatSubstring() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("formatSubstring"))
	return foundation.NSStringFromID(rv).String()
}

// The placeholder text for the message component.
//
// # Discussion
//
// The `placeholder` property holds an empty value when it is the last
// component.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/placeholder
func (o OSLogMessageComponent) Placeholder() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("placeholder"))
	return foundation.NSStringFromID(rv).String()
}

// The argument formatted as a sequence of bytes.
//
// # Discussion
//
// The `argumentDataValue` property can be `nil` if the argument can’t be
// decoded. For example, redacted arguments and the last component can’t be
// decoded.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentDataValue
func (o OSLogMessageComponent) ArgumentDataValue() foundation.NSData {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("argumentDataValue"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// The argument formatted as a double.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentDoubleValue
func (o OSLogMessageComponent) ArgumentDoubleValue() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("argumentDoubleValue"))
	return rv
}

// The argument formatted as a signed 64-bit integer.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentInt64Value
func (o OSLogMessageComponent) ArgumentInt64Value() int64 {
	rv := objc.Send[int64](o.ID, objc.Sel("argumentInt64Value"))
	return rv
}

// The argument formatted as a number.
//
// # Discussion
//
// The `argumentNumberValue` property can be `nil` if the argument can’t be
// decoded. For example, redacted arguments and the last component can’t be
// decoded.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentNumberValue
func (o OSLogMessageComponent) ArgumentNumberValue() foundation.NSNumber {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("argumentNumberValue"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// The argument formatted as a string.
//
// # Discussion
//
// The `argumentStringValue` property can be `nil` if the argument can’t be
// decoded. For example, redacted arguments or the last component can’t be
// decoded.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentStringValue
func (o OSLogMessageComponent) ArgumentStringValue() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("argumentStringValue"))
	return foundation.NSStringFromID(rv).String()
}

// The argument formatted as an unsigned 64-bit integer.
//
// See: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentUInt64Value
func (o OSLogMessageComponent) ArgumentUInt64Value() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("argumentUInt64Value"))
	return rv
}

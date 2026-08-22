// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AUParameter] class.
var (
	_AUParameterClass     AUParameterClass
	_AUParameterClassOnce sync.Once
)

func getAUParameterClass() AUParameterClass {
	_AUParameterClassOnce.Do(func() {
		_AUParameterClass = AUParameterClass{class: objc.GetClass("AUParameter")}
	})
	return _AUParameterClass
}

// GetAUParameterClass returns the class object for AUParameter.
func GetAUParameterClass() AUParameterClass {
	return getAUParameterClass()
}

type AUParameterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AUParameterClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AUParameterClass) Alloc() AUParameter {
	rv := objc.Send[AUParameter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a single audio unit parameter.
//
// # Querying Parameter Properties
//
//   - [AUParameter.MinValue]: The parameter’s minimum value.
//   - [AUParameter.MaxValue]: The parameter’s maximum value.
//   - [AUParameter.Unit]: The parameter’s unit of measurement.
//   - [AUParameter.UnitName]: The parameter’s localized unit name.
//   - [AUParameter.Flags]: The parameter’s characteristic details.
//   - [AUParameter.Address]: The parameter’s address.
//   - [AUParameter.ValueStrings]: The parameter’s localized value strings.
//   - [AUParameter.DependentParameters]: Any other parameter’s whose values may change as a side effect of this parameter’s value changing.
//
// # Managing Parameter Values
//
//   - [AUParameter.Value]: The parameter’s current value.
//   - [AUParameter.SetValue]
//   - [AUParameter.SetValueOriginator]: Sets the parameter’s value, avoiding redundant notifications to the originator.
//   - [AUParameter.SetValueOriginatorAtHostTime]: Sets the parameter’s value, preserving the host time of the gesture that initiated the change.
//   - [AUParameter.SetValueOriginatorAtHostTimeEventType]
//   - [AUParameter.StringFromValue]: Gets the string representation of a parameter value.
//   - [AUParameter.ValueFromString]: Converts a string into a parameter value.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameter
type AUParameter struct {
	AUParameterNode
}

// AUParameterFromID constructs a [AUParameter] from an objc.ID.
//
// An object that represents a single audio unit parameter.
func AUParameterFromID(id objc.ID) AUParameter {
	return AUParameter{AUParameterNode: AUParameterNodeFromID(id)}
}

// NOTE: AUParameter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AUParameter] class.
//
// # Querying Parameter Properties
//
//   - [IAUParameter.MinValue]: The parameter’s minimum value.
//   - [IAUParameter.MaxValue]: The parameter’s maximum value.
//   - [IAUParameter.Unit]: The parameter’s unit of measurement.
//   - [IAUParameter.UnitName]: The parameter’s localized unit name.
//   - [IAUParameter.Flags]: The parameter’s characteristic details.
//   - [IAUParameter.Address]: The parameter’s address.
//   - [IAUParameter.ValueStrings]: The parameter’s localized value strings.
//   - [IAUParameter.DependentParameters]: Any other parameter’s whose values may change as a side effect of this parameter’s value changing.
//
// # Managing Parameter Values
//
//   - [IAUParameter.Value]: The parameter’s current value.
//   - [IAUParameter.SetValue]
//   - [IAUParameter.SetValueOriginator]: Sets the parameter’s value, avoiding redundant notifications to the originator.
//   - [IAUParameter.SetValueOriginatorAtHostTime]: Sets the parameter’s value, preserving the host time of the gesture that initiated the change.
//   - [IAUParameter.SetValueOriginatorAtHostTimeEventType]
//   - [IAUParameter.StringFromValue]: Gets the string representation of a parameter value.
//   - [IAUParameter.ValueFromString]: Converts a string into a parameter value.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameter
type IAUParameter interface {
	IAUParameterNode

	// Topic: Querying Parameter Properties

	// The parameter’s minimum value.
	MinValue() AUValue
	// The parameter’s maximum value.
	MaxValue() AUValue
	// The parameter’s unit of measurement.
	Unit() AudioUnitParameterUnit
	// The parameter’s localized unit name.
	UnitName() string
	// The parameter’s characteristic details.
	Flags() AudioUnitParameterOptions
	// The parameter’s address.
	Address() AUParameterAddress
	// The parameter’s localized value strings.
	ValueStrings() []string
	// Any other parameter’s whose values may change as a side effect of this parameter’s value changing.
	DependentParameters() []foundation.NSNumber

	// Topic: Managing Parameter Values

	// The parameter’s current value.
	Value() AUValue
	SetValue(value AUValue)
	// Sets the parameter’s value, avoiding redundant notifications to the originator.
	SetValueOriginator(value AUValue, originator AUParameterObserverToken)
	// Sets the parameter’s value, preserving the host time of the gesture that initiated the change.
	SetValueOriginatorAtHostTime(value AUValue, originator AUParameterObserverToken, hostTime uint64)
	SetValueOriginatorAtHostTimeEventType(value AUValue, originator AUParameterObserverToken, hostTime uint64, eventType AUParameterAutomationEventType)
	// Gets the string representation of a parameter value.
	StringFromValue(value *AUValue) string
	// Converts a string into a parameter value.
	ValueFromString(string_ string) AUValue
}

// Init initializes the instance.
func (p AUParameter) Init() AUParameter {
	rv := objc.Send[AUParameter](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AUParameter) Autorelease() AUParameter {
	rv := objc.Send[AUParameter](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAUParameter creates a new AUParameter instance.
func NewAUParameter() AUParameter {
	class := getAUParameterClass()
	rv := objc.Send[AUParameter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUParameter/init(coder:)
func NewParameterWithCoder(coder foundation.INSCoder) AUParameter {
	instance := getAUParameterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AUParameterFromID(rv)
}

// Sets the parameter’s value, avoiding redundant notifications to the
// originator.
//
// value: The parameter’s new value.
//
// originator: The originator of the change in value. This token allows for observer
// management to avoid notification callback loops.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameter/setValue(_:originator:)
func (p AUParameter) SetValueOriginator(value AUValue, originator AUParameterObserverToken) {
	objc.Send[objc.ID](p.ID, objc.Sel("setValue:originator:"), value, originator)
}

// Sets the parameter’s value, preserving the host time of the gesture that
// initiated the change.
//
// value: The parameter’s new value.
//
// originator: The originator of the change in value. This token allows for observer
// management to avoid notification callback loops.
//
// hostTime: The time at which to schedule the change in value. This parameter allows
// for synchronization with other events.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameter/setValue(_:originator:atHostTime:)
func (p AUParameter) SetValueOriginatorAtHostTime(value AUValue, originator AUParameterObserverToken, hostTime uint64) {
	objc.Send[objc.ID](p.ID, objc.Sel("setValue:originator:atHostTime:"), value, originator, hostTime)
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUParameter/setValue(_:originator:atHostTime:eventType:)
func (p AUParameter) SetValueOriginatorAtHostTimeEventType(value AUValue, originator AUParameterObserverToken, hostTime uint64, eventType AUParameterAutomationEventType) {
	objc.Send[objc.ID](p.ID, objc.Sel("setValue:originator:atHostTime:eventType:"), value, originator, hostTime, eventType)
}

// Gets the string representation of a parameter value.
//
// value: The parameter value to represent as a string.
//
// # Return Value
//
// The string representation of a parameter value.
//
// # Discussion
//
// Pass `nil` into the `value` parameter to use the current value.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameter/string(fromValue:)
func (p AUParameter) StringFromValue(value *AUValue) string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("stringFromValue:"), unsafe.Pointer(value))
	return foundation.NSStringFromID(rv).String()
}

// Converts a string into a parameter value.
//
// string: The string representation of a parameter value.
//
// # Return Value
//
// The parameter value obtained from the string.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameter/value(from:)
func (p AUParameter) ValueFromString(string_ string) AUValue {
	rv := objc.Send[AUValue](p.ID, objc.Sel("valueFromString:"), objc.String(string_))
	return AUValue(rv)
}

// The parameter’s minimum value.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameter/minValue
func (p AUParameter) MinValue() AUValue {
	rv := objc.Send[AUValue](p.ID, objc.Sel("minValue"))
	return AUValue(rv)
}

// The parameter’s maximum value.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameter/maxValue
func (p AUParameter) MaxValue() AUValue {
	rv := objc.Send[AUValue](p.ID, objc.Sel("maxValue"))
	return AUValue(rv)
}

// The parameter’s unit of measurement.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameter/unit
func (p AUParameter) Unit() AudioUnitParameterUnit {
	rv := objc.Send[AudioUnitParameterUnit](p.ID, objc.Sel("unit"))
	return AudioUnitParameterUnit(rv)
}

// The parameter’s localized unit name.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameter/unitName
func (p AUParameter) UnitName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("unitName"))
	return foundation.NSStringFromID(rv).String()
}

// The parameter’s characteristic details.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameter/flags
func (p AUParameter) Flags() AudioUnitParameterOptions {
	rv := objc.Send[AudioUnitParameterOptions](p.ID, objc.Sel("flags"))
	return AudioUnitParameterOptions(rv)
}

// The parameter’s address.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameter/address
func (p AUParameter) Address() AUParameterAddress {
	rv := objc.Send[AUParameterAddress](p.ID, objc.Sel("address"))
	return AUParameterAddress(rv)
}

// The parameter’s localized value strings.
//
// # Discussion
//
// This property allows you to specify an array of strings to be used for the
// values of an indexed parameter—for example, a band parameter could
// publish these values: “High”, “Medium”, “Low”.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameter/valueStrings
func (p AUParameter) ValueStrings() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("valueStrings"))
	return objc.ConvertSliceToStrings(rv)
}

// Any other parameter’s whose values may change as a side effect of this
// parameter’s value changing.
//
// # Discussion
//
// The array contains [NSNumber] objects representing [AUParameterAddress]
// values.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameter/dependentParameters
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (p AUParameter) DependentParameters() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("dependentParameters"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// The parameter’s current value.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameter/value
func (p AUParameter) Value() AUValue {
	rv := objc.Send[AUValue](p.ID, objc.Sel("value"))
	return AUValue(rv)
}
func (p AUParameter) SetValue(value AUValue) {
	objc.Send[struct{}](p.ID, objc.Sel("setValue:"), value)
}

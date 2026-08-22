// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKSmartCardUserInteractionForPINOperation] class.
var (
	_TKSmartCardUserInteractionForPINOperationClass     TKSmartCardUserInteractionForPINOperationClass
	_TKSmartCardUserInteractionForPINOperationClassOnce sync.Once
)

func getTKSmartCardUserInteractionForPINOperationClass() TKSmartCardUserInteractionForPINOperationClass {
	_TKSmartCardUserInteractionForPINOperationClassOnce.Do(func() {
		_TKSmartCardUserInteractionForPINOperationClass = TKSmartCardUserInteractionForPINOperationClass{class: objc.GetClass("TKSmartCardUserInteractionForPINOperation")}
	})
	return _TKSmartCardUserInteractionForPINOperationClass
}

// GetTKSmartCardUserInteractionForPINOperationClass returns the class object for TKSmartCardUserInteractionForPINOperation.
func GetTKSmartCardUserInteractionForPINOperationClass() TKSmartCardUserInteractionForPINOperationClass {
	return getTKSmartCardUserInteractionForPINOperationClass()
}

type TKSmartCardUserInteractionForPINOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKSmartCardUserInteractionForPINOperationClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKSmartCardUserInteractionForPINOperationClass) Alloc() TKSmartCardUserInteractionForPINOperation {
	rv := objc.Send[TKSmartCardUserInteractionForPINOperation](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// A representation of user interaction for secure PIN operations on a Smart
// Card reader.
//
// # Overview
//
// There are two types of user interactions: those for secure PIN change and
// those for secure PIN validation. These interactions are instances of the
// [TKSmartCardUserInteractionForSecurePINChange], or
// [TKSmartCardUserInteractionForSecurePINVerification] subclasses of
// [TKSmartCardUserInteractionForPINOperation], respectively.
//
// You interact with instances of one of the subclasses of
// [TKSmartCardUserInteractionForPINOperation] when calling the
// [TKSmartCard.UserInteractionForSecurePINChangeWithPINFormatAPDUCurrentPINByteOffsetNewPINByteOffset]
// and
// [TKSmartCard.UserInteractionForSecurePINVerificationWithPINFormatAPDUPINByteOffset]
// methods on an [TKSmartCard] object.
//
// The result of a user interaction is available once the interaction has
// completed.
//
// # Managing Pin Completion
//
//   - [TKSmartCardUserInteractionForPINOperation.PINCompletion]: The conditions under which PIN entry should be considered complete.
//   - [TKSmartCardUserInteractionForPINOperation.SetPINCompletion]
//
// # Configuring Messages
//
//   - [TKSmartCardUserInteractionForPINOperation.PINMessageIndices]: A list of message indices referring to a predefined message table, used to specify the type and number of messages displayed during the PIN operation. `nil` by default.
//   - [TKSmartCardUserInteractionForPINOperation.SetPINMessageIndices]
//   - [TKSmartCardUserInteractionForPINOperation.Locale]: The locale for the displayed messages. If `nil`, the user’s current locale is used. By default, this value is the current locale of the system.
//   - [TKSmartCardUserInteractionForPINOperation.SetLocale]
//
// # Accessing Response Data
//
//   - [TKSmartCardUserInteractionForPINOperation.ResultSW]: The SW1-SW2 status bytes.
//   - [TKSmartCardUserInteractionForPINOperation.SetResultSW]
//   - [TKSmartCardUserInteractionForPINOperation.ResultData]: The returned data without SW1-SW2 bytes, if any.
//   - [TKSmartCardUserInteractionForPINOperation.SetResultData]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation
type TKSmartCardUserInteractionForPINOperation struct {
	TKSmartCardUserInteraction
}

// TKSmartCardUserInteractionForPINOperationFromID constructs a [TKSmartCardUserInteractionForPINOperation] from an objc.ID.
//
// A representation of user interaction for secure PIN operations on a Smart
// Card reader.
func TKSmartCardUserInteractionForPINOperationFromID(id objc.ID) TKSmartCardUserInteractionForPINOperation {
	return TKSmartCardUserInteractionForPINOperation{TKSmartCardUserInteraction: TKSmartCardUserInteractionFromID(id)}
}

// NOTE: TKSmartCardUserInteractionForPINOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKSmartCardUserInteractionForPINOperation] class.
//
// # Managing Pin Completion
//
//   - [ITKSmartCardUserInteractionForPINOperation.PINCompletion]: The conditions under which PIN entry should be considered complete.
//   - [ITKSmartCardUserInteractionForPINOperation.SetPINCompletion]
//
// # Configuring Messages
//
//   - [ITKSmartCardUserInteractionForPINOperation.PINMessageIndices]: A list of message indices referring to a predefined message table, used to specify the type and number of messages displayed during the PIN operation. `nil` by default.
//   - [ITKSmartCardUserInteractionForPINOperation.SetPINMessageIndices]
//   - [ITKSmartCardUserInteractionForPINOperation.Locale]: The locale for the displayed messages. If `nil`, the user’s current locale is used. By default, this value is the current locale of the system.
//   - [ITKSmartCardUserInteractionForPINOperation.SetLocale]
//
// # Accessing Response Data
//
//   - [ITKSmartCardUserInteractionForPINOperation.ResultSW]: The SW1-SW2 status bytes.
//   - [ITKSmartCardUserInteractionForPINOperation.SetResultSW]
//   - [ITKSmartCardUserInteractionForPINOperation.ResultData]: The returned data without SW1-SW2 bytes, if any.
//   - [ITKSmartCardUserInteractionForPINOperation.SetResultData]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation
type ITKSmartCardUserInteractionForPINOperation interface {
	ITKSmartCardUserInteraction

	// Topic: Managing Pin Completion

	// The conditions under which PIN entry should be considered complete.
	PINCompletion() TKSmartCardPINCompletion
	SetPINCompletion(value TKSmartCardPINCompletion)

	// Topic: Configuring Messages

	// A list of message indices referring to a predefined message table, used to specify the type and number of messages displayed during the PIN operation. `nil` by default.
	PINMessageIndices() []foundation.NSNumber
	SetPINMessageIndices(value []foundation.NSNumber)
	// The locale for the displayed messages. If `nil`, the user’s current locale is used. By default, this value is the current locale of the system.
	Locale() foundation.NSLocale
	SetLocale(value foundation.NSLocale)

	// Topic: Accessing Response Data

	// The SW1-SW2 status bytes.
	ResultSW() uint16
	SetResultSW(value uint16)
	// The returned data without SW1-SW2 bytes, if any.
	ResultData() foundation.NSData
	SetResultData(value foundation.NSData)
}

// Init initializes the instance.
func (t TKSmartCardUserInteractionForPINOperation) Init() TKSmartCardUserInteractionForPINOperation {
	rv := objc.Send[TKSmartCardUserInteractionForPINOperation](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKSmartCardUserInteractionForPINOperation) Autorelease() TKSmartCardUserInteractionForPINOperation {
	rv := objc.Send[TKSmartCardUserInteractionForPINOperation](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardUserInteractionForPINOperation creates a new TKSmartCardUserInteractionForPINOperation instance.
func NewTKSmartCardUserInteractionForPINOperation() TKSmartCardUserInteractionForPINOperation {
	class := getTKSmartCardUserInteractionForPINOperationClass()
	rv := objc.Send[TKSmartCardUserInteractionForPINOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The conditions under which PIN entry should be considered complete.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/pinCompletion
func (t TKSmartCardUserInteractionForPINOperation) PINCompletion() TKSmartCardPINCompletion {
	rv := objc.Send[TKSmartCardPINCompletion](t.ID, objc.Sel("PINCompletion"))
	return TKSmartCardPINCompletion(rv)
}
func (t TKSmartCardUserInteractionForPINOperation) SetPINCompletion(value TKSmartCardPINCompletion) {
	objc.Send[struct{}](t.ID, objc.Sel("setPINCompletion:"), value)
}

// A list of message indices referring to a predefined message table, used to
// specify the type and number of messages displayed during the PIN operation.
// `nil` by default.
//
// # Discussion
//
// If `nil`, the reader does not display any message (reader specific).
// Typically, PIN verification takes 1 message; PIN modification takes 1 – 3
// messages.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/pinMessageIndices
func (t TKSmartCardUserInteractionForPINOperation) PINMessageIndices() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("PINMessageIndices"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (t TKSmartCardUserInteractionForPINOperation) SetPINMessageIndices(value []foundation.NSNumber) {
	objc.Send[struct{}](t.ID, objc.Sel("setPINMessageIndices:"), objectivec.IObjectSliceToNSArray(value))
}

// The locale for the displayed messages. If `nil`, the user’s current
// locale is used. By default, this value is the current locale of the system.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/locale
func (t TKSmartCardUserInteractionForPINOperation) Locale() foundation.NSLocale {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("locale"))
	return foundation.NSLocaleFromID(objc.ID(rv))
}
func (t TKSmartCardUserInteractionForPINOperation) SetLocale(value foundation.NSLocale) {
	objc.Send[struct{}](t.ID, objc.Sel("setLocale:"), value)
}

// The SW1-SW2 status bytes.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/resultSW
func (t TKSmartCardUserInteractionForPINOperation) ResultSW() uint16 {
	rv := objc.Send[uint16](t.ID, objc.Sel("resultSW"))
	return rv
}
func (t TKSmartCardUserInteractionForPINOperation) SetResultSW(value uint16) {
	objc.Send[struct{}](t.ID, objc.Sel("setResultSW:"), value)
}

// The returned data without SW1-SW2 bytes, if any.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForPINOperation/resultData
func (t TKSmartCardUserInteractionForPINOperation) ResultData() foundation.NSData {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("resultData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (t TKSmartCardUserInteractionForPINOperation) SetResultData(value foundation.NSData) {
	objc.Send[struct{}](t.ID, objc.Sel("setResultData:"), value)
}

// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [TKTokenSmartCardPINAuthOperation] class.
var (
	_TKTokenSmartCardPINAuthOperationClass     TKTokenSmartCardPINAuthOperationClass
	_TKTokenSmartCardPINAuthOperationClassOnce sync.Once
)

func getTKTokenSmartCardPINAuthOperationClass() TKTokenSmartCardPINAuthOperationClass {
	_TKTokenSmartCardPINAuthOperationClassOnce.Do(func() {
		_TKTokenSmartCardPINAuthOperationClass = TKTokenSmartCardPINAuthOperationClass{class: objc.GetClass("TKTokenSmartCardPINAuthOperation")}
	})
	return _TKTokenSmartCardPINAuthOperationClass
}

// GetTKTokenSmartCardPINAuthOperationClass returns the class object for TKTokenSmartCardPINAuthOperation.
func GetTKTokenSmartCardPINAuthOperationClass() TKTokenSmartCardPINAuthOperationClass {
	return getTKTokenSmartCardPINAuthOperationClass()
}

type TKTokenSmartCardPINAuthOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKTokenSmartCardPINAuthOperationClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKTokenSmartCardPINAuthOperationClass) Alloc() TKTokenSmartCardPINAuthOperation {
	rv := objc.Send[TKTokenSmartCardPINAuthOperation](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// A Smart Card PIN authentication operation.
//
// # Configuring the Operation
//
//   - [TKTokenSmartCardPINAuthOperation.PINFormat]: The PIN format.
//   - [TKTokenSmartCardPINAuthOperation.SetPINFormat]
//   - [TKTokenSmartCardPINAuthOperation.APDUTemplate]: The template into which the PIN is filled in. `nil` by default.
//   - [TKTokenSmartCardPINAuthOperation.SetAPDUTemplate]
//   - [TKTokenSmartCardPINAuthOperation.PINByteOffset]: The offset, in bytes, within the APDU template to mark the location for filling in the PIN.
//   - [TKTokenSmartCardPINAuthOperation.SetPINByteOffset]
//   - [TKTokenSmartCardPINAuthOperation.SmartCard]: A Smart Card to which the formatted APDU is sent in order to authenticate.
//   - [TKTokenSmartCardPINAuthOperation.SetSmartCard]
//
// # Accessing the PIN
//
//   - [TKTokenSmartCardPINAuthOperation.PIN]: The PIN value resulting from performing the operation.
//   - [TKTokenSmartCardPINAuthOperation.SetPIN]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSmartCardPINAuthOperation
type TKTokenSmartCardPINAuthOperation struct {
	TKTokenAuthOperation
}

// TKTokenSmartCardPINAuthOperationFromID constructs a [TKTokenSmartCardPINAuthOperation] from an objc.ID.
//
// A Smart Card PIN authentication operation.
func TKTokenSmartCardPINAuthOperationFromID(id objc.ID) TKTokenSmartCardPINAuthOperation {
	return TKTokenSmartCardPINAuthOperation{TKTokenAuthOperation: TKTokenAuthOperationFromID(id)}
}

// NOTE: TKTokenSmartCardPINAuthOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKTokenSmartCardPINAuthOperation] class.
//
// # Configuring the Operation
//
//   - [ITKTokenSmartCardPINAuthOperation.PINFormat]: The PIN format.
//   - [ITKTokenSmartCardPINAuthOperation.SetPINFormat]
//   - [ITKTokenSmartCardPINAuthOperation.APDUTemplate]: The template into which the PIN is filled in. `nil` by default.
//   - [ITKTokenSmartCardPINAuthOperation.SetAPDUTemplate]
//   - [ITKTokenSmartCardPINAuthOperation.PINByteOffset]: The offset, in bytes, within the APDU template to mark the location for filling in the PIN.
//   - [ITKTokenSmartCardPINAuthOperation.SetPINByteOffset]
//   - [ITKTokenSmartCardPINAuthOperation.SmartCard]: A Smart Card to which the formatted APDU is sent in order to authenticate.
//   - [ITKTokenSmartCardPINAuthOperation.SetSmartCard]
//
// # Accessing the PIN
//
//   - [ITKTokenSmartCardPINAuthOperation.PIN]: The PIN value resulting from performing the operation.
//   - [ITKTokenSmartCardPINAuthOperation.SetPIN]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSmartCardPINAuthOperation
type ITKTokenSmartCardPINAuthOperation interface {
	ITKTokenAuthOperation

	// Topic: Configuring the Operation

	// The PIN format.
	PINFormat() ITKSmartCardPINFormat
	SetPINFormat(value ITKSmartCardPINFormat)
	// The template into which the PIN is filled in. `nil` by default.
	APDUTemplate() foundation.NSData
	SetAPDUTemplate(value foundation.NSData)
	// The offset, in bytes, within the APDU template to mark the location for filling in the PIN.
	PINByteOffset() int
	SetPINByteOffset(value int)
	// A Smart Card to which the formatted APDU is sent in order to authenticate.
	SmartCard() ITKSmartCard
	SetSmartCard(value ITKSmartCard)

	// Topic: Accessing the PIN

	// The PIN value resulting from performing the operation.
	PIN() string
	SetPIN(value string)
}

// Init initializes the instance.
func (t TKTokenSmartCardPINAuthOperation) Init() TKTokenSmartCardPINAuthOperation {
	rv := objc.Send[TKTokenSmartCardPINAuthOperation](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKTokenSmartCardPINAuthOperation) Autorelease() TKTokenSmartCardPINAuthOperation {
	rv := objc.Send[TKTokenSmartCardPINAuthOperation](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenSmartCardPINAuthOperation creates a new TKTokenSmartCardPINAuthOperation instance.
func NewTKTokenSmartCardPINAuthOperation() TKTokenSmartCardPINAuthOperation {
	class := getTKTokenSmartCardPINAuthOperationClass()
	rv := objc.Send[TKTokenSmartCardPINAuthOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The PIN format.
//
// # Discussion
//
// By default, this property is set to a [TKSmartCardPINFormat] object
// initialized without any further configuration.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSmartCardPINAuthOperation/pinFormat
func (t TKTokenSmartCardPINAuthOperation) PINFormat() ITKSmartCardPINFormat {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("PINFormat"))
	return TKSmartCardPINFormatFromID(objc.ID(rv))
}
func (t TKTokenSmartCardPINAuthOperation) SetPINFormat(value ITKSmartCardPINFormat) {
	objc.Send[struct{}](t.ID, objc.Sel("setPINFormat:"), value)
}

// The template into which the PIN is filled in. `nil` by default.
//
// # Discussion
//
// If `nil`, the system will not attempt to authenticate by sending the
// formatted APDU to the Smart Card. Instead, the token itself is expected to
// perform the authentication. You are encouraged to provide an APDU template,
// if possible, as it allows the use of a hardware interface for secure PIN
// entry, provided one exists.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSmartCardPINAuthOperation/apduTemplate
func (t TKTokenSmartCardPINAuthOperation) APDUTemplate() foundation.NSData {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("APDUTemplate"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (t TKTokenSmartCardPINAuthOperation) SetAPDUTemplate(value foundation.NSData) {
	objc.Send[struct{}](t.ID, objc.Sel("setAPDUTemplate:"), value)
}

// The offset, in bytes, within the APDU template to mark the location for
// filling in the PIN.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSmartCardPINAuthOperation/pinByteOffset
func (t TKTokenSmartCardPINAuthOperation) PINByteOffset() int {
	rv := objc.Send[int](t.ID, objc.Sel("PINByteOffset"))
	return rv
}
func (t TKTokenSmartCardPINAuthOperation) SetPINByteOffset(value int) {
	objc.Send[struct{}](t.ID, objc.Sel("setPINByteOffset:"), value)
}

// A Smart Card to which the formatted APDU is sent in order to authenticate.
//
// # Discussion
//
// This property is only used if the
// [TKTokenSmartCardPINAuthOperation.APDUTemplate] property has a set value.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSmartCardPINAuthOperation/smartCard
func (t TKTokenSmartCardPINAuthOperation) SmartCard() ITKSmartCard {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("smartCard"))
	return TKSmartCardFromID(objc.ID(rv))
}
func (t TKTokenSmartCardPINAuthOperation) SetSmartCard(value ITKSmartCard) {
	objc.Send[struct{}](t.ID, objc.Sel("setSmartCard:"), value)
}

// The PIN value resulting from performing the operation.
//
// # Discussion
//
// This property is set to the result of the operation after “ is called.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSmartCardPINAuthOperation/pin
func (t TKTokenSmartCardPINAuthOperation) PIN() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("PIN"))
	return foundation.NSStringFromID(rv).String()
}
func (t TKTokenSmartCardPINAuthOperation) SetPIN(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setPIN:"), objc.String(value))
}

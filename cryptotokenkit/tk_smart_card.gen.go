// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKSmartCard] class.
var (
	_TKSmartCardClass     TKSmartCardClass
	_TKSmartCardClassOnce sync.Once
)

func getTKSmartCardClass() TKSmartCardClass {
	_TKSmartCardClassOnce.Do(func() {
		_TKSmartCardClass = TKSmartCardClass{class: objc.GetClass("TKSmartCard")}
	})
	return _TKSmartCardClass
}

// GetTKSmartCardClass returns the class object for TKSmartCard.
func GetTKSmartCardClass() TKSmartCardClass {
	return getTKSmartCardClass()
}

type TKSmartCardClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKSmartCardClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKSmartCardClass) Alloc() TKSmartCard {
	rv := objc.Send[TKSmartCard](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a smart card.
//
// # Overview
//
// This class provides an interface for managing sessions with a smart card,
// transmitting requests, and facilitating user interaction.
//
// You can create a [TKSmartCard] object when a smart card is inserted into a
// slot, by calling the [TKSmartCardSlot.SmartCard] method on the
// corresponding [TKSmartCardSlot] object. To start communicating with the
// smart card, call the [TKSmartCard.BeginSessionWithReply] method on the
// [TKSmartCard] object. Once an exclusive session has been established, you
// transmit data using the [TKSmartCard.TransmitRequestReply] method. After
// you’ve finished communicating with a smart card, you call the
// [TKSmartCard.EndSession] method.
//
// If the smart card is physically removed from its slot, the session object
// becomes invalid, and any further calls to
// [TKSmartCard.TransmitRequestReply] will return an error. You can use
// Key-Value Observing on the [TKSmartCard.Valid] property to be notified when
// a smart card is invalidated, due to being removed from the slot or another
// reason.
//
// # Configuring the Smart Card
//
//   - [TKSmartCard.Slot]: The slot in which the Smart Card is inserted.
//   - [TKSmartCard.Valid]: Whether the Smart Card is valid and accessible from its slot.
//   - [TKSmartCard.Sensitive]: Whether sessions established for the Smart Card should be considered sensitive. [false](<https://developer.apple.com/documentation/Swift/false>) by default.
//   - [TKSmartCard.SetSensitive]
//   - [TKSmartCard.Context]: User-specified information. This property is automatically set to `nil` if the Smart Card is removed or another [TKSmartCard] object begins a session.
//   - [TKSmartCard.SetContext]
//
// # Setting the Communication Protocol
//
//   - [TKSmartCard.AllowedProtocols]: The protocols allowed for communication with the Smart Card. [any](<https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardProtocol/any>) by default.
//   - [TKSmartCard.SetAllowedProtocols]
//   - [TKSmartCard.CurrentProtocol]: The protocol used for communication with the Smart Card. Returns [TKSmartCardProtocolNone](<https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardProtocol/TKSmartCardProtocolNone>) if no session is currently established.
//
// # Communicating with the Smart Card
//
//   - [TKSmartCard.BeginSessionWithReply]: Begins a session with the Smart Card.
//   - [TKSmartCard.TransmitRequestReply]: Transmits data in Application Protocol Data Unit (APDU) format to the Smart Card.
//   - [TKSmartCard.EndSession]: Completes any pending transmissions and ends the session to the Smart Card.
//
// # Managing User Interaction
//
//   - [TKSmartCard.UserInteractionForSecurePINVerificationWithPINFormatAPDUPINByteOffset]: Creates and returns a new user interaction object for secure PIN verification using the Smart Card reader facilities.
//   - [TKSmartCard.UserInteractionForSecurePINChangeWithPINFormatAPDUCurrentPINByteOffsetNewPINByteOffset]: Creates a new user interaction object for secure PIN change using the smart card reader facilities (typically a HW keypad).
//
// # Configuring APDU Behavior
//
//   - [TKSmartCard.Cla]: The CLA byte used for APDU transmission. `0x00` by default.
//   - [TKSmartCard.SetCla]
//   - [TKSmartCard.UseExtendedLength]: Whether to use extended length APDU.
//   - [TKSmartCard.SetUseExtendedLength]
//   - [TKSmartCard.UseCommandChaining]: Whether to use command chaining of APDU with a data field longer than 255 bytes.
//   - [TKSmartCard.SetUseCommandChaining]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard
type TKSmartCard struct {
	objectivec.Object
}

// TKSmartCardFromID constructs a [TKSmartCard] from an objc.ID.
//
// A representation of a smart card.
func TKSmartCardFromID(id objc.ID) TKSmartCard {
	return TKSmartCard{objectivec.Object{ID: id}}
}

// NOTE: TKSmartCard adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKSmartCard] class.
//
// # Configuring the Smart Card
//
//   - [ITKSmartCard.Slot]: The slot in which the Smart Card is inserted.
//   - [ITKSmartCard.Valid]: Whether the Smart Card is valid and accessible from its slot.
//   - [ITKSmartCard.Sensitive]: Whether sessions established for the Smart Card should be considered sensitive. [false](<https://developer.apple.com/documentation/Swift/false>) by default.
//   - [ITKSmartCard.SetSensitive]
//   - [ITKSmartCard.Context]: User-specified information. This property is automatically set to `nil` if the Smart Card is removed or another [TKSmartCard] object begins a session.
//   - [ITKSmartCard.SetContext]
//
// # Setting the Communication Protocol
//
//   - [ITKSmartCard.AllowedProtocols]: The protocols allowed for communication with the Smart Card. [any](<https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardProtocol/any>) by default.
//   - [ITKSmartCard.SetAllowedProtocols]
//   - [ITKSmartCard.CurrentProtocol]: The protocol used for communication with the Smart Card. Returns [TKSmartCardProtocolNone](<https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardProtocol/TKSmartCardProtocolNone>) if no session is currently established.
//
// # Communicating with the Smart Card
//
//   - [ITKSmartCard.BeginSessionWithReply]: Begins a session with the Smart Card.
//   - [ITKSmartCard.TransmitRequestReply]: Transmits data in Application Protocol Data Unit (APDU) format to the Smart Card.
//   - [ITKSmartCard.EndSession]: Completes any pending transmissions and ends the session to the Smart Card.
//
// # Managing User Interaction
//
//   - [ITKSmartCard.UserInteractionForSecurePINVerificationWithPINFormatAPDUPINByteOffset]: Creates and returns a new user interaction object for secure PIN verification using the Smart Card reader facilities.
//   - [ITKSmartCard.UserInteractionForSecurePINChangeWithPINFormatAPDUCurrentPINByteOffsetNewPINByteOffset]: Creates a new user interaction object for secure PIN change using the smart card reader facilities (typically a HW keypad).
//
// # Configuring APDU Behavior
//
//   - [ITKSmartCard.Cla]: The CLA byte used for APDU transmission. `0x00` by default.
//   - [ITKSmartCard.SetCla]
//   - [ITKSmartCard.UseExtendedLength]: Whether to use extended length APDU.
//   - [ITKSmartCard.SetUseExtendedLength]
//   - [ITKSmartCard.UseCommandChaining]: Whether to use command chaining of APDU with a data field longer than 255 bytes.
//   - [ITKSmartCard.SetUseCommandChaining]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard
type ITKSmartCard interface {
	objectivec.IObject

	// Topic: Configuring the Smart Card

	// The slot in which the Smart Card is inserted.
	Slot() ITKSmartCardSlot
	// Whether the Smart Card is valid and accessible from its slot.
	Valid() bool
	// Whether sessions established for the Smart Card should be considered sensitive. [false](<https://developer.apple.com/documentation/Swift/false>) by default.
	Sensitive() bool
	SetSensitive(value bool)
	// User-specified information. This property is automatically set to `nil` if the Smart Card is removed or another [TKSmartCard] object begins a session.
	Context() objectivec.IObject
	SetContext(value objectivec.IObject)

	// Topic: Setting the Communication Protocol

	// The protocols allowed for communication with the Smart Card. [any](<https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardProtocol/any>) by default.
	AllowedProtocols() TKSmartCardProtocol
	SetAllowedProtocols(value TKSmartCardProtocol)
	// The protocol used for communication with the Smart Card. Returns [TKSmartCardProtocolNone](<https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardProtocol/TKSmartCardProtocolNone>) if no session is currently established.
	CurrentProtocol() TKSmartCardProtocol

	// Topic: Communicating with the Smart Card

	// Begins a session with the Smart Card.
	BeginSessionWithReply(reply BoolErrorHandler)
	// Transmits data in Application Protocol Data Unit (APDU) format to the Smart Card.
	TransmitRequestReply(request foundation.NSData, reply DataErrorHandler)
	// Completes any pending transmissions and ends the session to the Smart Card.
	EndSession()

	// Topic: Managing User Interaction

	// Creates and returns a new user interaction object for secure PIN verification using the Smart Card reader facilities.
	UserInteractionForSecurePINVerificationWithPINFormatAPDUPINByteOffset(PINFormat ITKSmartCardPINFormat, APDU foundation.NSData, PINByteOffset int) ITKSmartCardUserInteractionForSecurePINVerification
	// Creates a new user interaction object for secure PIN change using the smart card reader facilities (typically a HW keypad).
	UserInteractionForSecurePINChangeWithPINFormatAPDUCurrentPINByteOffsetNewPINByteOffset(PINFormat ITKSmartCardPINFormat, APDU foundation.NSData, currentPINByteOffset int, newPINByteOffset int) ITKSmartCardUserInteractionForSecurePINChange

	// Topic: Configuring APDU Behavior

	// The CLA byte used for APDU transmission. `0x00` by default.
	Cla() uint8
	SetCla(value uint8)
	// Whether to use extended length APDU.
	UseExtendedLength() bool
	SetUseExtendedLength(value bool)
	// Whether to use command chaining of APDU with a data field longer than 255 bytes.
	UseCommandChaining() bool
	SetUseCommandChaining(value bool)
}

// Init initializes the instance.
func (t TKSmartCard) Init() TKSmartCard {
	rv := objc.Send[TKSmartCard](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKSmartCard) Autorelease() TKSmartCard {
	rv := objc.Send[TKSmartCard](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCard creates a new TKSmartCard instance.
func NewTKSmartCard() TKSmartCard {
	class := getTKSmartCardClass()
	rv := objc.Send[TKSmartCard](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Begins a session with the Smart Card.
//
// reply: success: Whether the session could be established successfully. error:
// Contains information about the error preventing the transaction from being
// established.
//
// The [NSError] object is created in the [TKErrorDomain] domain with a code
// in the [TKError.Code] enumeration.
//
// # Discussion
//
// This method will fail if there is already an existing session for the Smart
// Card.
//
// Calls to this method must be balanced with calls to
// [TKSmartCard.EndSession].
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/beginSession(reply:)
//
// [TKError.Code]: https://developer.apple.com/documentation/CryptoTokenKit/TKError/Code
// [TKErrorDomain]: https://developer.apple.com/documentation/CryptoTokenKit/TKErrorDomain
func (t TKSmartCard) BeginSessionWithReply(reply BoolErrorHandler) {
	_block0, _ := NewBoolErrorBlock(reply)
	objc.Send[objc.ID](t.ID, objc.Sel("beginSessionWithReply:"), _block0)
}

// Transmits data in Application Protocol Data Unit (APDU) format to the Smart
// Card.
//
// request: The APDU request data.
//
// reply: response: The APDU response data, or `nil` if communication with the Smart
// Card failed. error: Contains information about the the error preventing the
// transaction from being established.
//
// The [NSError] object is created in the [TKErrorDomain] domain with a code
// in the [TKError.Code] enumeration.
//
// # Discussion
//
// You should only call this method after a session to the Smart Card has been
// established using the [TKSmartCard.BeginSessionWithReply] method, and
// before the session is terminated using the [TKSmartCard.EndSession] method.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/transmit(_:reply:)
//
// [TKError.Code]: https://developer.apple.com/documentation/CryptoTokenKit/TKError/Code
// [TKErrorDomain]: https://developer.apple.com/documentation/CryptoTokenKit/TKErrorDomain
func (t TKSmartCard) TransmitRequestReply(request foundation.NSData, reply DataErrorHandler) {
	_block1, _ := NewDataErrorBlock(reply)
	objc.Send[objc.ID](t.ID, objc.Sel("transmitRequest:reply:"), request, _block1)
}

// Completes any pending transmissions and ends the session to the Smart Card.
//
// # Discussion
//
// Calls to this method should balance calls to
// [TKSmartCard.BeginSessionWithReply].
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/endSession()
func (t TKSmartCard) EndSession() {
	objc.Send[objc.ID](t.ID, objc.Sel("endSession"))
}

// Creates and returns a new user interaction object for secure PIN
// verification using the Smart Card reader facilities.
//
// PINFormat: The PIN format descriptor.
//
// APDU: The Application Protocol Data Unit (APDU) used by the Smart Card to fill in
// PIN data.
//
// PINByteOffset: The offset, in bytes, within the Application Protocol Data Unit (APDU)
// field to mark a location of a PIN block for filling in the entered PIN.
//
// # Return Value
//
// A new user interaction object for secure PIN verification, or `nil` if this
// feature is not supported by the Smart Card reader.
//
// # Discussion
//
// You should only call this method after a session to the Smart Card has been
// established using the [TKSmartCard.BeginSessionWithReply] method, and
// before the session is terminated using the [TKSmartCard.EndSession] method.
//
// Once the interaction has been successfully completed, the results are
// available via the [TKSmartCardUserInteractionForPINOperation.ResultData]
// and [TKSmartCardUserInteractionForPINOperation.ResultSW] properties of the
// returned [TKSmartCardUserInteractionForSecurePINVerification] object.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/userInteractionForSecurePINVerification(_:apdu:pinByteOffset:)
func (t TKSmartCard) UserInteractionForSecurePINVerificationWithPINFormatAPDUPINByteOffset(PINFormat ITKSmartCardPINFormat, APDU foundation.NSData, PINByteOffset int) ITKSmartCardUserInteractionForSecurePINVerification {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("userInteractionForSecurePINVerificationWithPINFormat:APDU:PINByteOffset:"), PINFormat, APDU, PINByteOffset)
	return TKSmartCardUserInteractionForSecurePINVerificationFromID(rv)
}

// Creates a new user interaction object for secure PIN change using the smart
// card reader facilities (typically a HW keypad).
//
// PINFormat: The PIN format descriptor.
//
// APDU: The Application Protocol Data Unit (APDU) used by the Smart Card to fill in
// PIN data.
//
// currentPINByteOffset: The offset, in bytes, within the Application Protocol Data Unit (APDU)
// field to mark a location of a PIN block for filling in the entered PIN.
//
// newPINByteOffset: The offset, in bytes, within the Application Protocol Data Unit (APDU)
// field to mark a location of a PIN block for filling in the new PIN.
//
// # Return Value
//
// A new user interaction object for secure PIN verification, or `nil` if this
// feature is not supported by the Smart Card reader.
//
// # Discussion
//
// You should only call this method after a session to the Smart Card has been
// established using the [TKSmartCard.BeginSessionWithReply] method, and
// before the session is terminated using the [TKSmartCard.EndSession] method.
//
// Once the interaction has been successfully completed, the results are
// available via the [TKSmartCardUserInteractionForPINOperation.ResultData]
// and [TKSmartCardUserInteractionForPINOperation.ResultSW] properties of the
// returned [TKSmartCardUserInteractionForSecurePINVerification] object.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/userInteractionForSecurePINChange(_:apdu:currentPINByteOffset:newPINByteOffset:)
func (t TKSmartCard) UserInteractionForSecurePINChangeWithPINFormatAPDUCurrentPINByteOffsetNewPINByteOffset(PINFormat ITKSmartCardPINFormat, APDU foundation.NSData, currentPINByteOffset int, newPINByteOffset int) ITKSmartCardUserInteractionForSecurePINChange {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("userInteractionForSecurePINChangeWithPINFormat:APDU:currentPINByteOffset:newPINByteOffset:"), PINFormat, APDU, currentPINByteOffset, newPINByteOffset)
	return TKSmartCardUserInteractionForSecurePINChangeFromID(rv)
}

// The slot in which the Smart Card is inserted.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/slot
func (t TKSmartCard) Slot() ITKSmartCardSlot {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("slot"))
	return TKSmartCardSlotFromID(objc.ID(rv))
}

// Whether the Smart Card is valid and accessible from its slot.
//
// # Discussion
//
// Use Key-Value-Observing to be notified for changes to accessibility, such
// as when a Smart Card is physically removed from its slot. For more
// information, see [Key-Value Observing Programming Guide].
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/isValid
//
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
func (t TKSmartCard) Valid() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("valid"))
	return rv
}

// Whether sessions established for the Smart Card should be considered
// sensitive. false by default.
//
// # Discussion
//
// When this property is set to true, any sessions established for the
// receiver will begin and end by sending a reset command to the Smart Card.
// This is recommended anytime potentially sensitive information is
// transferred.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/isSensitive
func (t TKSmartCard) Sensitive() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("sensitive"))
	return rv
}
func (t TKSmartCard) SetSensitive(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setSensitive:"), value)
}

// User-specified information. This property is automatically set to `nil` if
// the Smart Card is removed or another [TKSmartCard] object begins a session.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/context
func (t TKSmartCard) Context() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("context"))
	return objectivec.Object{ID: rv}
}
func (t TKSmartCard) SetContext(value objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setContext:"), value)
}

// The protocols allowed for communication with the Smart Card.
// [TKSmartCardProtocolAny] by default.
//
// # Discussion
//
// This property is consulted only when beginning a session to a Smart Card.
// Any changes to this property will not be reflected by the current session,
// if one is already established.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/allowedProtocols
func (t TKSmartCard) AllowedProtocols() TKSmartCardProtocol {
	rv := objc.Send[TKSmartCardProtocol](t.ID, objc.Sel("allowedProtocols"))
	return TKSmartCardProtocol(rv)
}
func (t TKSmartCard) SetAllowedProtocols(value TKSmartCardProtocol) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowedProtocols:"), value)
}

// The protocol used for communication with the Smart Card. Returns
// [TKSmartCardProtocolNone] if no session is currently established.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/currentProtocol
//
// [TKSmartCardProtocolNone]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardProtocol/TKSmartCardProtocolNone
func (t TKSmartCard) CurrentProtocol() TKSmartCardProtocol {
	rv := objc.Send[TKSmartCardProtocol](t.ID, objc.Sel("currentProtocol"))
	return TKSmartCardProtocol(rv)
}

// The CLA byte used for APDU transmission. `0x00` by default.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/cla
func (t TKSmartCard) Cla() uint8 {
	rv := objc.Send[uint8](t.ID, objc.Sel("cla"))
	return rv
}
func (t TKSmartCard) SetCla(value uint8) {
	objc.Send[struct{}](t.ID, objc.Sel("setCla:"), value)
}

// Whether to use extended length APDU.
//
// # Discussion
//
// By default, this property is set to true when the Smart Card slot supports
// transmitting extended length commands, and the ATR announces that extended
// length APDU is supported.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/useExtendedLength
func (t TKSmartCard) UseExtendedLength() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("useExtendedLength"))
	return rv
}
func (t TKSmartCard) SetUseExtendedLength(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setUseExtendedLength:"), value)
}

// Whether to use command chaining of APDU with a data field longer than 255
// bytes.
//
// # Discussion
//
// By default, this property is set to true when the Smart Card ATR announces
// that command chaining is supported.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard/useCommandChaining
func (t TKSmartCard) UseCommandChaining() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("useCommandChaining"))
	return rv
}
func (t TKSmartCard) SetUseCommandChaining(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setUseCommandChaining:"), value)
}

// BeginSessionWithReplySync is a synchronous wrapper around [TKSmartCard.BeginSessionWithReply].
// It blocks until the completion handler fires or the context is cancelled.
func (t TKSmartCard) BeginSessionWithReplySync(ctx context.Context) (bool, error) {
	type result struct {
		val bool
		err error
	}
	done := make(chan result, 1)
	t.BeginSessionWithReply(func(val bool, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// TransmitRequestReplySync is a synchronous wrapper around [TKSmartCard.TransmitRequestReply].
// It blocks until the completion handler fires or the context is cancelled.
func (t TKSmartCard) TransmitRequestReplySync(ctx context.Context, request foundation.NSData) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	t.TransmitRequestReply(request, func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

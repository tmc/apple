// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKSmartCardSlot] class.
var (
	_TKSmartCardSlotClass     TKSmartCardSlotClass
	_TKSmartCardSlotClassOnce sync.Once
)

func getTKSmartCardSlotClass() TKSmartCardSlotClass {
	_TKSmartCardSlotClassOnce.Do(func() {
		_TKSmartCardSlotClass = TKSmartCardSlotClass{class: objc.GetClass("TKSmartCardSlot")}
	})
	return _TKSmartCardSlotClass
}

// GetTKSmartCardSlotClass returns the class object for TKSmartCardSlot.
func GetTKSmartCardSlotClass() TKSmartCardSlotClass {
	return getTKSmartCardSlotClass()
}

type TKSmartCardSlotClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKSmartCardSlotClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKSmartCardSlotClass) Alloc() TKSmartCardSlot {
	rv := objc.Send[TKSmartCardSlot](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// A single smart card reader slot in the system.
//
// # Overview
//
// Use the [TKSmartCardSlotManager] class to manage all the smart card reader
// slots available to the system. You can retrieve the names of available
// smart card reader slots for a system using the
// [TKSmartCardSlotManager.SlotNames] property of a manager object, and access
// instances of [TKSmartCardSlot] using the
// [TKSmartCardSlotManager.GetSlotWithNameReply] method.
//
// # Instantiating Smart Cards
//
//   - [TKSmartCardSlot.MakeSmartCard]: Creates a new [TKSmartCard](<https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard>) object representing the currently inserted Smart Card.
//
// # Getting the Slot State
//
//   - [TKSmartCardSlot.State]: The current state of the Smart Card reader slot.
//
// # Getting the Slot Configuration
//
//   - [TKSmartCardSlot.Name]: The name of the Smart Card reader slot.
//   - [TKSmartCardSlot.MaxInputLength]: The maximum length of input APDU (Application Protocol Data Unit) that the Smart Card reader slot is able to transfer to the Smart Card.
//   - [TKSmartCardSlot.MaxOutputLength]: The maximum length of output APDU (Application Protocol Data Unit) that the Smart Card reader slot is able to transfer from the Smart Card.
//
// # Reading the Answer to Reset
//
//   - [TKSmartCardSlot.ATR]: The ATR (Answer to Reset) of the inserted Smart Card, or `nil` if no Smart Card is inserted or the inserted Smart Card is mute.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot
type TKSmartCardSlot struct {
	objectivec.Object
}

// TKSmartCardSlotFromID constructs a [TKSmartCardSlot] from an objc.ID.
//
// A single smart card reader slot in the system.
func TKSmartCardSlotFromID(id objc.ID) TKSmartCardSlot {
	return TKSmartCardSlot{objectivec.Object{ID: id}}
}

// NOTE: TKSmartCardSlot adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKSmartCardSlot] class.
//
// # Instantiating Smart Cards
//
//   - [ITKSmartCardSlot.MakeSmartCard]: Creates a new [TKSmartCard](<https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard>) object representing the currently inserted Smart Card.
//
// # Getting the Slot State
//
//   - [ITKSmartCardSlot.State]: The current state of the Smart Card reader slot.
//
// # Getting the Slot Configuration
//
//   - [ITKSmartCardSlot.Name]: The name of the Smart Card reader slot.
//   - [ITKSmartCardSlot.MaxInputLength]: The maximum length of input APDU (Application Protocol Data Unit) that the Smart Card reader slot is able to transfer to the Smart Card.
//   - [ITKSmartCardSlot.MaxOutputLength]: The maximum length of output APDU (Application Protocol Data Unit) that the Smart Card reader slot is able to transfer from the Smart Card.
//
// # Reading the Answer to Reset
//
//   - [ITKSmartCardSlot.ATR]: The ATR (Answer to Reset) of the inserted Smart Card, or `nil` if no Smart Card is inserted or the inserted Smart Card is mute.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot
type ITKSmartCardSlot interface {
	objectivec.IObject

	// Topic: Instantiating Smart Cards

	// Creates a new [TKSmartCard](<https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCard>) object representing the currently inserted Smart Card.
	MakeSmartCard() ITKSmartCard

	// Topic: Getting the Slot State

	// The current state of the Smart Card reader slot.
	State() TKSmartCardSlotState

	// Topic: Getting the Slot Configuration

	// The name of the Smart Card reader slot.
	Name() string
	// The maximum length of input APDU (Application Protocol Data Unit) that the Smart Card reader slot is able to transfer to the Smart Card.
	MaxInputLength() int
	// The maximum length of output APDU (Application Protocol Data Unit) that the Smart Card reader slot is able to transfer from the Smart Card.
	MaxOutputLength() int

	// Topic: Reading the Answer to Reset

	// The ATR (Answer to Reset) of the inserted Smart Card, or `nil` if no Smart Card is inserted or the inserted Smart Card is mute.
	ATR() ITKSmartCardATR
}

// Init initializes the instance.
func (t TKSmartCardSlot) Init() TKSmartCardSlot {
	rv := objc.Send[TKSmartCardSlot](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKSmartCardSlot) Autorelease() TKSmartCardSlot {
	rv := objc.Send[TKSmartCardSlot](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardSlot creates a new TKSmartCardSlot instance.
func NewTKSmartCardSlot() TKSmartCardSlot {
	class := getTKSmartCardSlotClass()
	rv := objc.Send[TKSmartCardSlot](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new [TKSmartCard] object representing the currently inserted
// Smart Card.
//
// # Return Value
//
// A new [TKSmartCard] object, or `nil` if no Smart Card is currently
// inserted.
//
// # Discussion
//
// You can create multiple instances of [TKSmartCard] that represent the same
// Smart Card. Exclusivity of data transfer is handled by sessions on the
// individual [TKSmartCard] objects.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/makeSmartCard()
func (t TKSmartCardSlot) MakeSmartCard() ITKSmartCard {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("makeSmartCard"))
	return TKSmartCardFromID(rv)
}

// The current state of the Smart Card reader slot.
//
// # Discussion
//
// Use Key-Value-Observing to be notified for changes to the state. For more
// information, see [Key-Value Observing Programming Guide].
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/state-swift.property
//
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
func (t TKSmartCardSlot) State() TKSmartCardSlotState {
	rv := objc.Send[TKSmartCardSlotState](t.ID, objc.Sel("state"))
	return TKSmartCardSlotState(rv)
}

// The name of the Smart Card reader slot.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/name
func (t TKSmartCardSlot) Name() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The maximum length of input APDU (Application Protocol Data Unit) that the
// Smart Card reader slot is able to transfer to the Smart Card.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/maxInputLength
func (t TKSmartCardSlot) MaxInputLength() int {
	rv := objc.Send[int](t.ID, objc.Sel("maxInputLength"))
	return rv
}

// The maximum length of output APDU (Application Protocol Data Unit) that the
// Smart Card reader slot is able to transfer from the Smart Card.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/maxOutputLength
func (t TKSmartCardSlot) MaxOutputLength() int {
	rv := objc.Send[int](t.ID, objc.Sel("maxOutputLength"))
	return rv
}

// The ATR (Answer to Reset) of the inserted Smart Card, or `nil` if no Smart
// Card is inserted or the inserted Smart Card is mute.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/atr
func (t TKSmartCardSlot) ATR() ITKSmartCardATR {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("ATR"))
	return TKSmartCardATRFromID(objc.ID(rv))
}

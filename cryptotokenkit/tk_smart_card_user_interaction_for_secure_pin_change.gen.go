// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [TKSmartCardUserInteractionForSecurePINChange] class.
var (
	_TKSmartCardUserInteractionForSecurePINChangeClass     TKSmartCardUserInteractionForSecurePINChangeClass
	_TKSmartCardUserInteractionForSecurePINChangeClassOnce sync.Once
)

func getTKSmartCardUserInteractionForSecurePINChangeClass() TKSmartCardUserInteractionForSecurePINChangeClass {
	_TKSmartCardUserInteractionForSecurePINChangeClassOnce.Do(func() {
		_TKSmartCardUserInteractionForSecurePINChangeClass = TKSmartCardUserInteractionForSecurePINChangeClass{class: objc.GetClass("TKSmartCardUserInteractionForSecurePINChange")}
	})
	return _TKSmartCardUserInteractionForSecurePINChangeClass
}

// GetTKSmartCardUserInteractionForSecurePINChangeClass returns the class object for TKSmartCardUserInteractionForSecurePINChange.
func GetTKSmartCardUserInteractionForSecurePINChangeClass() TKSmartCardUserInteractionForSecurePINChangeClass {
	return getTKSmartCardUserInteractionForSecurePINChangeClass()
}

type TKSmartCardUserInteractionForSecurePINChangeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKSmartCardUserInteractionForSecurePINChangeClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKSmartCardUserInteractionForSecurePINChangeClass) Alloc() TKSmartCardUserInteractionForSecurePINChange {
	rv := objc.Send[TKSmartCardUserInteractionForSecurePINChange](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// A representation of the user interaction for secure PIN change operations
// on a Smart Card reader.
//
// # Overview
//
// The result of a user interaction is available once the interaction has
// completed.
//
// # Configuring User Interaction
//
//   - [TKSmartCardUserInteractionForSecurePINChange.PINConfirmation]: The way PIN confirmation is requested. [TKSmartCardPINConfirmationNone] by default.
//   - [TKSmartCardUserInteractionForSecurePINChange.SetPINConfirmation]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForSecurePINChange
type TKSmartCardUserInteractionForSecurePINChange struct {
	TKSmartCardUserInteractionForPINOperation
}

// TKSmartCardUserInteractionForSecurePINChangeFromID constructs a [TKSmartCardUserInteractionForSecurePINChange] from an objc.ID.
//
// A representation of the user interaction for secure PIN change operations
// on a Smart Card reader.
func TKSmartCardUserInteractionForSecurePINChangeFromID(id objc.ID) TKSmartCardUserInteractionForSecurePINChange {
	return TKSmartCardUserInteractionForSecurePINChange{TKSmartCardUserInteractionForPINOperation: TKSmartCardUserInteractionForPINOperationFromID(id)}
}

// NOTE: TKSmartCardUserInteractionForSecurePINChange adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKSmartCardUserInteractionForSecurePINChange] class.
//
// # Configuring User Interaction
//
//   - [ITKSmartCardUserInteractionForSecurePINChange.PINConfirmation]: The way PIN confirmation is requested. [TKSmartCardPINConfirmationNone] by default.
//   - [ITKSmartCardUserInteractionForSecurePINChange.SetPINConfirmation]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForSecurePINChange
type ITKSmartCardUserInteractionForSecurePINChange interface {
	ITKSmartCardUserInteractionForPINOperation

	// Topic: Configuring User Interaction

	// The way PIN confirmation is requested. [TKSmartCardPINConfirmationNone] by default.
	PINConfirmation() TKSmartCardPINConfirmation
	SetPINConfirmation(value TKSmartCardPINConfirmation)
}

// Init initializes the instance.
func (t TKSmartCardUserInteractionForSecurePINChange) Init() TKSmartCardUserInteractionForSecurePINChange {
	rv := objc.Send[TKSmartCardUserInteractionForSecurePINChange](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKSmartCardUserInteractionForSecurePINChange) Autorelease() TKSmartCardUserInteractionForSecurePINChange {
	rv := objc.Send[TKSmartCardUserInteractionForSecurePINChange](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardUserInteractionForSecurePINChange creates a new TKSmartCardUserInteractionForSecurePINChange instance.
func NewTKSmartCardUserInteractionForSecurePINChange() TKSmartCardUserInteractionForSecurePINChange {
	class := getTKSmartCardUserInteractionForSecurePINChangeClass()
	rv := objc.Send[TKSmartCardUserInteractionForSecurePINChange](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The way PIN confirmation is requested. [TKSmartCardPINConfirmationNone] by
// default.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForSecurePINChange/pinConfirmation
func (t TKSmartCardUserInteractionForSecurePINChange) PINConfirmation() TKSmartCardPINConfirmation {
	rv := objc.Send[TKSmartCardPINConfirmation](t.ID, objc.Sel("PINConfirmation"))
	return TKSmartCardPINConfirmation(rv)
}
func (t TKSmartCardUserInteractionForSecurePINChange) SetPINConfirmation(value TKSmartCardPINConfirmation) {
	objc.Send[struct{}](t.ID, objc.Sel("setPINConfirmation:"), value)
}

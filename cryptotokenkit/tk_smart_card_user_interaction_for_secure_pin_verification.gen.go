// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [TKSmartCardUserInteractionForSecurePINVerification] class.
var (
	_TKSmartCardUserInteractionForSecurePINVerificationClass     TKSmartCardUserInteractionForSecurePINVerificationClass
	_TKSmartCardUserInteractionForSecurePINVerificationClassOnce sync.Once
)

func getTKSmartCardUserInteractionForSecurePINVerificationClass() TKSmartCardUserInteractionForSecurePINVerificationClass {
	_TKSmartCardUserInteractionForSecurePINVerificationClassOnce.Do(func() {
		_TKSmartCardUserInteractionForSecurePINVerificationClass = TKSmartCardUserInteractionForSecurePINVerificationClass{class: objc.GetClass("TKSmartCardUserInteractionForSecurePINVerification")}
	})
	return _TKSmartCardUserInteractionForSecurePINVerificationClass
}

// GetTKSmartCardUserInteractionForSecurePINVerificationClass returns the class object for TKSmartCardUserInteractionForSecurePINVerification.
func GetTKSmartCardUserInteractionForSecurePINVerificationClass() TKSmartCardUserInteractionForSecurePINVerificationClass {
	return getTKSmartCardUserInteractionForSecurePINVerificationClass()
}

type TKSmartCardUserInteractionForSecurePINVerificationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKSmartCardUserInteractionForSecurePINVerificationClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKSmartCardUserInteractionForSecurePINVerificationClass) Alloc() TKSmartCardUserInteractionForSecurePINVerification {
	rv := objc.Send[TKSmartCardUserInteractionForSecurePINVerification](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// A representation of the user interaction for secure PIN change verification
// on a Smart Card reader.
//
// # Overview
//
// The result of a user interaction is available once the interaction has
// completed.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForSecurePINVerification
type TKSmartCardUserInteractionForSecurePINVerification struct {
	TKSmartCardUserInteractionForPINOperation
}

// TKSmartCardUserInteractionForSecurePINVerificationFromID constructs a [TKSmartCardUserInteractionForSecurePINVerification] from an objc.ID.
//
// A representation of the user interaction for secure PIN change verification
// on a Smart Card reader.
func TKSmartCardUserInteractionForSecurePINVerificationFromID(id objc.ID) TKSmartCardUserInteractionForSecurePINVerification {
	return TKSmartCardUserInteractionForSecurePINVerification{TKSmartCardUserInteractionForPINOperation: TKSmartCardUserInteractionForPINOperationFromID(id)}
}

// NOTE: TKSmartCardUserInteractionForSecurePINVerification adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKSmartCardUserInteractionForSecurePINVerification] class.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionForSecurePINVerification
type ITKSmartCardUserInteractionForSecurePINVerification interface {
	ITKSmartCardUserInteractionForPINOperation
}

// Init initializes the instance.
func (t TKSmartCardUserInteractionForSecurePINVerification) Init() TKSmartCardUserInteractionForSecurePINVerification {
	rv := objc.Send[TKSmartCardUserInteractionForSecurePINVerification](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKSmartCardUserInteractionForSecurePINVerification) Autorelease() TKSmartCardUserInteractionForSecurePINVerification {
	rv := objc.Send[TKSmartCardUserInteractionForSecurePINVerification](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardUserInteractionForSecurePINVerification creates a new TKSmartCardUserInteractionForSecurePINVerification instance.
func NewTKSmartCardUserInteractionForSecurePINVerification() TKSmartCardUserInteractionForSecurePINVerification {
	class := getTKSmartCardUserInteractionForSecurePINVerificationClass()
	rv := objc.Send[TKSmartCardUserInteractionForSecurePINVerification](objc.ID(class.class), objc.Sel("new"))
	return rv
}

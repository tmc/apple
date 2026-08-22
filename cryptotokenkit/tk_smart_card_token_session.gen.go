// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [TKSmartCardTokenSession] class.
var (
	_TKSmartCardTokenSessionClass     TKSmartCardTokenSessionClass
	_TKSmartCardTokenSessionClassOnce sync.Once
)

func getTKSmartCardTokenSessionClass() TKSmartCardTokenSessionClass {
	_TKSmartCardTokenSessionClassOnce.Do(func() {
		_TKSmartCardTokenSessionClass = TKSmartCardTokenSessionClass{class: objc.GetClass("TKSmartCardTokenSession")}
	})
	return _TKSmartCardTokenSessionClass
}

// GetTKSmartCardTokenSessionClass returns the class object for TKSmartCardTokenSession.
func GetTKSmartCardTokenSessionClass() TKSmartCardTokenSessionClass {
	return getTKSmartCardTokenSessionClass()
}

type TKSmartCardTokenSessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKSmartCardTokenSessionClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKSmartCardTokenSessionClass) Alloc() TKSmartCardTokenSession {
	rv := objc.Send[TKSmartCardTokenSession](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// A token session that is based on a smart card token.
//
// # Overview
//
// You can use the [TKSmartCardTokenSession.SmartCard] property to access and
// send APDUs to the underlying smart card.
//
// # Accessing the Smart Card
//
//   - [TKSmartCardTokenSession.SmartCard]: The smart card for the active exclusive session and selected application.
//
// # Instance Methods
//
//   - [TKSmartCardTokenSession.GetSmartCardWithError]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenSession
type TKSmartCardTokenSession struct {
	TKTokenSession
}

// TKSmartCardTokenSessionFromID constructs a [TKSmartCardTokenSession] from an objc.ID.
//
// A token session that is based on a smart card token.
func TKSmartCardTokenSessionFromID(id objc.ID) TKSmartCardTokenSession {
	return TKSmartCardTokenSession{TKTokenSession: TKTokenSessionFromID(id)}
}

// NOTE: TKSmartCardTokenSession adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKSmartCardTokenSession] class.
//
// # Accessing the Smart Card
//
//   - [ITKSmartCardTokenSession.SmartCard]: The smart card for the active exclusive session and selected application.
//
// # Instance Methods
//
//   - [ITKSmartCardTokenSession.GetSmartCardWithError]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenSession
type ITKSmartCardTokenSession interface {
	ITKTokenSession

	// Topic: Accessing the Smart Card

	// The smart card for the active exclusive session and selected application.
	SmartCard() ITKSmartCard

	// Topic: Instance Methods

	GetSmartCardWithError() (ITKSmartCard, error)
}

// Init initializes the instance.
func (t TKSmartCardTokenSession) Init() TKSmartCardTokenSession {
	rv := objc.Send[TKSmartCardTokenSession](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKSmartCardTokenSession) Autorelease() TKSmartCardTokenSession {
	rv := objc.Send[TKSmartCardTokenSession](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardTokenSession creates a new TKSmartCardTokenSession instance.
func NewTKSmartCardTokenSession() TKSmartCardTokenSession {
	class := getTKSmartCardTokenSessionClass()
	rv := objc.Send[TKSmartCardTokenSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a token session with the specified token.
//
// token: The token to which the initialized session is bound.
//
// # Return Value
//
// A new token session created with the specified token.
//
// # Discussion
//
// This is the designated initializer.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSession/init(token:)
func NewTKSmartCardTokenSessionWithToken(token ITKToken) TKSmartCardTokenSession {
	instance := getTKSmartCardTokenSessionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithToken:"), token)
	return TKSmartCardTokenSessionFromID(rv)
}

// # Discussion
//
// Returns a TKSmartCard instance with an active exclusive session and the
// SmartCard application selected. Replaces the deprecated @c smartCard
// property.
//
// The TKSmartCard object is only accessible within the methods of the
// TKTokenSessionDelegate protocol. If the associated token has an AID set,
// the returned card will have an exclusive session already opened and the
// specified application selected. In this scenario: Do not call -[TKSmartCard
// beginSessionWithReply:]) on the returned SmartCard instance. The system
// manages the session lifecycle and will terminate it automatically when the
// current token request servicing is finished. Do not call -[TKSmartCard
// endSession]. You can use the `smartCard.Context()` property to store any
// context-specific state information related to the card. This property is
// automatically set to `nil` if the card is reset or accessed by a different
// TKSmartCard instance (potentially in another process). Before performing an
// operation, check the `TKSmartCard.Context()` property for a previously
// stored value. This can help you avoid potentially costly restoration of the
// SmartCard state if it’s already available.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenSession/getSmartCard()
func (t TKSmartCardTokenSession) GetSmartCardWithError() (ITKSmartCard, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("getSmartCardWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return TKSmartCard{}, foundation.NSErrorFrom(errorPtr)
	}
	return TKSmartCardFromID(rv), nil

}

// The smart card for the active exclusive session and selected application.
//
// # Discussion
//
// This property can only be accessed in the implementation of a
// [TKTokenSessionDelegate] protocol delegate method. If the associated token
// has a value set for the [TKSmartCardToken.AID] property, this property
// opens an exclusive session to the card, with the application already
// selected.
//
// You should not call [TKSmartCard.BeginSessionWithReply] or
// [TKSmartCard.EndSession] on the returned value. Instead, the system will
// take care of beginning the exclusive session and terminating it when the
// current token request servicing is finished.
//
// You can store any kind of information representing state of the card using
// the [TKSmartCard.Context] property. This property will be automatically set
// to `nil` if the card is reset or accessed by different [TKSmartCard]
// instance, such as by another process. You can check the
// [TKSmartCard.Context] property for any previously stored values as a way to
// avoid costly state restoration before performing an operation.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenSession/smartCard
func (t TKSmartCardTokenSession) SmartCard() ITKSmartCard {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("smartCard"))
	return TKSmartCardFromID(objc.ID(rv))
}

// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKTokenSession] class.
var (
	_TKTokenSessionClass     TKTokenSessionClass
	_TKTokenSessionClassOnce sync.Once
)

func getTKTokenSessionClass() TKTokenSessionClass {
	_TKTokenSessionClassOnce.Do(func() {
		_TKTokenSessionClass = TKTokenSessionClass{class: objc.GetClass("TKTokenSession")}
	})
	return _TKTokenSessionClass
}

// GetTKTokenSessionClass returns the class object for TKTokenSession.
func GetTKTokenSessionClass() TKTokenSessionClass {
	return getTKTokenSessionClass()
}

type TKTokenSessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKTokenSessionClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKTokenSessionClass) Alloc() TKTokenSession {
	rv := objc.Send[TKTokenSession](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// A token session that manages the authentication state of a token.
//
// # Overview
//
// A token session communicates with its delegate to perform operations with
// its token that are bound to the authentication state.
//
// A session is always instantiated by a [TKToken] instance through the
// token’s delegate when the framework detects access to the token from a
// new authentication session.
//
// # Creating Token Sessions
//
//   - [TKTokenSession.InitWithToken]: Initializes a token session with the specified token.
//
// # Responding to Authentication Events
//
//   - [TKTokenSession.Delegate]: The token session delegate.
//   - [TKTokenSession.SetDelegate]
//
// # Accessing the Token
//
//   - [TKTokenSession.Token]: The token to which the session is bound.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSession
type TKTokenSession struct {
	objectivec.Object
}

// TKTokenSessionFromID constructs a [TKTokenSession] from an objc.ID.
//
// A token session that manages the authentication state of a token.
func TKTokenSessionFromID(id objc.ID) TKTokenSession {
	return TKTokenSession{objectivec.Object{ID: id}}
}

// NOTE: TKTokenSession adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKTokenSession] class.
//
// # Creating Token Sessions
//
//   - [ITKTokenSession.InitWithToken]: Initializes a token session with the specified token.
//
// # Responding to Authentication Events
//
//   - [ITKTokenSession.Delegate]: The token session delegate.
//   - [ITKTokenSession.SetDelegate]
//
// # Accessing the Token
//
//   - [ITKTokenSession.Token]: The token to which the session is bound.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSession
type ITKTokenSession interface {
	objectivec.IObject

	// Topic: Creating Token Sessions

	// Initializes a token session with the specified token.
	InitWithToken(token ITKToken) TKTokenSession

	// Topic: Responding to Authentication Events

	// The token session delegate.
	Delegate() TKTokenSessionDelegate
	SetDelegate(value TKTokenSessionDelegate)

	// Topic: Accessing the Token

	// The token to which the session is bound.
	Token() ITKToken
}

// Init initializes the instance.
func (t TKTokenSession) Init() TKTokenSession {
	rv := objc.Send[TKTokenSession](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKTokenSession) Autorelease() TKTokenSession {
	rv := objc.Send[TKTokenSession](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenSession creates a new TKTokenSession instance.
func NewTKTokenSession() TKTokenSession {
	class := getTKTokenSessionClass()
	rv := objc.Send[TKTokenSession](objc.ID(class.class), objc.Sel("new"))
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
func NewTKTokenSessionWithToken(token ITKToken) TKTokenSession {
	instance := getTKTokenSessionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithToken:"), token)
	return TKTokenSessionFromID(rv)
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
func (t TKTokenSession) InitWithToken(token ITKToken) TKTokenSession {
	rv := objc.Send[TKTokenSession](t.ID, objc.Sel("initWithToken:"), token)
	return rv
}

// The token session delegate.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSession/delegate
func (t TKTokenSession) Delegate() TKTokenSessionDelegate {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("delegate"))
	return TKTokenSessionDelegateObjectFromID(rv)
}
func (t TKTokenSession) SetDelegate(value TKTokenSessionDelegate) {
	objc.Send[struct{}](t.ID, objc.Sel("setDelegate:"), value)
}

// The token to which the session is bound.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenSession/token
func (t TKTokenSession) Token() ITKToken {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("token"))
	return TKTokenFromID(objc.ID(rv))
}

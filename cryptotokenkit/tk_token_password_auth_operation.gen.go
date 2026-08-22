// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [TKTokenPasswordAuthOperation] class.
var (
	_TKTokenPasswordAuthOperationClass     TKTokenPasswordAuthOperationClass
	_TKTokenPasswordAuthOperationClassOnce sync.Once
)

func getTKTokenPasswordAuthOperationClass() TKTokenPasswordAuthOperationClass {
	_TKTokenPasswordAuthOperationClassOnce.Do(func() {
		_TKTokenPasswordAuthOperationClass = TKTokenPasswordAuthOperationClass{class: objc.GetClass("TKTokenPasswordAuthOperation")}
	})
	return _TKTokenPasswordAuthOperationClass
}

// GetTKTokenPasswordAuthOperationClass returns the class object for TKTokenPasswordAuthOperation.
func GetTKTokenPasswordAuthOperationClass() TKTokenPasswordAuthOperationClass {
	return getTKTokenPasswordAuthOperationClass()
}

type TKTokenPasswordAuthOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKTokenPasswordAuthOperationClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKTokenPasswordAuthOperationClass) Alloc() TKTokenPasswordAuthOperation {
	rv := objc.Send[TKTokenPasswordAuthOperation](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// A password-based authentication operation.
//
// # Managing the Password
//
//   - [TKTokenPasswordAuthOperation.Password]: The password to be filled in when the “ is called.
//   - [TKTokenPasswordAuthOperation.SetPassword]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenPasswordAuthOperation
type TKTokenPasswordAuthOperation struct {
	TKTokenAuthOperation
}

// TKTokenPasswordAuthOperationFromID constructs a [TKTokenPasswordAuthOperation] from an objc.ID.
//
// A password-based authentication operation.
func TKTokenPasswordAuthOperationFromID(id objc.ID) TKTokenPasswordAuthOperation {
	return TKTokenPasswordAuthOperation{TKTokenAuthOperation: TKTokenAuthOperationFromID(id)}
}

// NOTE: TKTokenPasswordAuthOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKTokenPasswordAuthOperation] class.
//
// # Managing the Password
//
//   - [ITKTokenPasswordAuthOperation.Password]: The password to be filled in when the “ is called.
//   - [ITKTokenPasswordAuthOperation.SetPassword]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenPasswordAuthOperation
type ITKTokenPasswordAuthOperation interface {
	ITKTokenAuthOperation

	// Topic: Managing the Password

	// The password to be filled in when the `` is called.
	Password() string
	SetPassword(value string)
}

// Init initializes the instance.
func (t TKTokenPasswordAuthOperation) Init() TKTokenPasswordAuthOperation {
	rv := objc.Send[TKTokenPasswordAuthOperation](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKTokenPasswordAuthOperation) Autorelease() TKTokenPasswordAuthOperation {
	rv := objc.Send[TKTokenPasswordAuthOperation](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenPasswordAuthOperation creates a new TKTokenPasswordAuthOperation instance.
func NewTKTokenPasswordAuthOperation() TKTokenPasswordAuthOperation {
	class := getTKTokenPasswordAuthOperationClass()
	rv := objc.Send[TKTokenPasswordAuthOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The password to be filled in when the “ is called.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenPasswordAuthOperation/password
func (t TKTokenPasswordAuthOperation) Password() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("password"))
	return foundation.NSStringFromID(rv).String()
}
func (t TKTokenPasswordAuthOperation) SetPassword(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setPassword:"), objc.String(value))
}

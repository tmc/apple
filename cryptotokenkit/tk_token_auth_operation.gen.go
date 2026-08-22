// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKTokenAuthOperation] class.
var (
	_TKTokenAuthOperationClass     TKTokenAuthOperationClass
	_TKTokenAuthOperationClassOnce sync.Once
)

func getTKTokenAuthOperationClass() TKTokenAuthOperationClass {
	_TKTokenAuthOperationClassOnce.Do(func() {
		_TKTokenAuthOperationClass = TKTokenAuthOperationClass{class: objc.GetClass("TKTokenAuthOperation")}
	})
	return _TKTokenAuthOperationClass
}

// GetTKTokenAuthOperationClass returns the class object for TKTokenAuthOperation.
func GetTKTokenAuthOperationClass() TKTokenAuthOperationClass {
	return getTKTokenAuthOperationClass()
}

type TKTokenAuthOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKTokenAuthOperationClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKTokenAuthOperationClass) Alloc() TKTokenAuthOperation {
	rv := objc.Send[TKTokenAuthOperation](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// An authentication operation for a cryptographic token.
//
// # Overview
//
// The CryptoTokenKit framework provides the following concrete subclasses:
// [TKTokenPasswordAuthOperation], for password-based authentication, and
// [TKTokenSmartCardPINAuthOperation] for Smart Card PIN-based authentication.
//
// # Finishing the Operation
//
//   - [TKTokenAuthOperation.FinishWithError]: Finishes the authentication operation.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenAuthOperation
type TKTokenAuthOperation struct {
	objectivec.Object
}

// TKTokenAuthOperationFromID constructs a [TKTokenAuthOperation] from an objc.ID.
//
// An authentication operation for a cryptographic token.
func TKTokenAuthOperationFromID(id objc.ID) TKTokenAuthOperation {
	return TKTokenAuthOperation{objectivec.Object{ID: id}}
}

// NOTE: TKTokenAuthOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKTokenAuthOperation] class.
//
// # Finishing the Operation
//
//   - [ITKTokenAuthOperation.FinishWithError]: Finishes the authentication operation.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenAuthOperation
type ITKTokenAuthOperation interface {
	objectivec.IObject

	// Topic: Finishing the Operation

	// Finishes the authentication operation.
	FinishWithError() (bool, error)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t TKTokenAuthOperation) Init() TKTokenAuthOperation {
	rv := objc.Send[TKTokenAuthOperation](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKTokenAuthOperation) Autorelease() TKTokenAuthOperation {
	rv := objc.Send[TKTokenAuthOperation](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenAuthOperation creates a new TKTokenAuthOperation instance.
func NewTKTokenAuthOperation() TKTokenAuthOperation {
	class := getTKTokenAuthOperationClass()
	rv := objc.Send[TKTokenAuthOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Finishes the authentication operation.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenAuthOperation/finish()
func (t TKTokenAuthOperation) FinishWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("finishWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("finishWithError: returned NO with nil NSError")
	}
	return rv, nil

}
func (t TKTokenAuthOperation) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

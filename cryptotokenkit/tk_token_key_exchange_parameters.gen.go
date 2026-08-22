// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKTokenKeyExchangeParameters] class.
var (
	_TKTokenKeyExchangeParametersClass     TKTokenKeyExchangeParametersClass
	_TKTokenKeyExchangeParametersClassOnce sync.Once
)

func getTKTokenKeyExchangeParametersClass() TKTokenKeyExchangeParametersClass {
	_TKTokenKeyExchangeParametersClassOnce.Do(func() {
		_TKTokenKeyExchangeParametersClass = TKTokenKeyExchangeParametersClass{class: objc.GetClass("TKTokenKeyExchangeParameters")}
	})
	return _TKTokenKeyExchangeParametersClass
}

// GetTKTokenKeyExchangeParametersClass returns the class object for TKTokenKeyExchangeParameters.
func GetTKTokenKeyExchangeParametersClass() TKTokenKeyExchangeParametersClass {
	return getTKTokenKeyExchangeParametersClass()
}

type TKTokenKeyExchangeParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKTokenKeyExchangeParametersClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKTokenKeyExchangeParametersClass) Alloc() TKTokenKeyExchangeParameters {
	rv := objc.Send[TKTokenKeyExchangeParameters](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// Parameters used to perform specific key exchange operations.
//
// # Accessing Parameters
//
//   - [TKTokenKeyExchangeParameters.RequestedSize]: Returns the requested output size, in bytes, of key exchange result.
//   - [TKTokenKeyExchangeParameters.SharedInfo]: Returns shared information typically used during the key derivation (KDF) step of a key exchange algorithm.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeyExchangeParameters
type TKTokenKeyExchangeParameters struct {
	objectivec.Object
}

// TKTokenKeyExchangeParametersFromID constructs a [TKTokenKeyExchangeParameters] from an objc.ID.
//
// Parameters used to perform specific key exchange operations.
func TKTokenKeyExchangeParametersFromID(id objc.ID) TKTokenKeyExchangeParameters {
	return TKTokenKeyExchangeParameters{objectivec.Object{ID: id}}
}

// NOTE: TKTokenKeyExchangeParameters adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKTokenKeyExchangeParameters] class.
//
// # Accessing Parameters
//
//   - [ITKTokenKeyExchangeParameters.RequestedSize]: Returns the requested output size, in bytes, of key exchange result.
//   - [ITKTokenKeyExchangeParameters.SharedInfo]: Returns shared information typically used during the key derivation (KDF) step of a key exchange algorithm.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeyExchangeParameters
type ITKTokenKeyExchangeParameters interface {
	objectivec.IObject

	// Topic: Accessing Parameters

	// Returns the requested output size, in bytes, of key exchange result.
	RequestedSize() int
	// Returns shared information typically used during the key derivation (KDF) step of a key exchange algorithm.
	SharedInfo() foundation.NSData
}

// Init initializes the instance.
func (t TKTokenKeyExchangeParameters) Init() TKTokenKeyExchangeParameters {
	rv := objc.Send[TKTokenKeyExchangeParameters](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKTokenKeyExchangeParameters) Autorelease() TKTokenKeyExchangeParameters {
	rv := objc.Send[TKTokenKeyExchangeParameters](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenKeyExchangeParameters creates a new TKTokenKeyExchangeParameters instance.
func NewTKTokenKeyExchangeParameters() TKTokenKeyExchangeParameters {
	class := getTKTokenKeyExchangeParametersClass()
	rv := objc.Send[TKTokenKeyExchangeParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the requested output size, in bytes, of key exchange result.
//
// # Discussion
//
// This property should be ignored if the output size is not configurable for
// the specified key exchange algorithm.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeyExchangeParameters/requestedSize
func (t TKTokenKeyExchangeParameters) RequestedSize() int {
	rv := objc.Send[int](t.ID, objc.Sel("requestedSize"))
	return rv
}

// Returns shared information typically used during the key derivation (KDF)
// step of a key exchange algorithm.
//
// # Discussion
//
// This property should be ignored if shared information isn’t used by the
// specified key exchange algorithm.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeyExchangeParameters/sharedInfo
func (t TKTokenKeyExchangeParameters) SharedInfo() foundation.NSData {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("sharedInfo"))
	return foundation.NSDataFromID(objc.ID(rv))
}

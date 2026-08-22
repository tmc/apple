// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKTokenKeyAlgorithm] class.
var (
	_TKTokenKeyAlgorithmClass     TKTokenKeyAlgorithmClass
	_TKTokenKeyAlgorithmClassOnce sync.Once
)

func getTKTokenKeyAlgorithmClass() TKTokenKeyAlgorithmClass {
	_TKTokenKeyAlgorithmClassOnce.Do(func() {
		_TKTokenKeyAlgorithmClass = TKTokenKeyAlgorithmClass{class: objc.GetClass("TKTokenKeyAlgorithm")}
	})
	return _TKTokenKeyAlgorithmClass
}

// GetTKTokenKeyAlgorithmClass returns the class object for TKTokenKeyAlgorithm.
func GetTKTokenKeyAlgorithmClass() TKTokenKeyAlgorithmClass {
	return getTKTokenKeyAlgorithmClass()
}

type TKTokenKeyAlgorithmClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKTokenKeyAlgorithmClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKTokenKeyAlgorithmClass) Alloc() TKTokenKeyAlgorithm {
	rv := objc.Send[TKTokenKeyAlgorithm](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// Cryptographic algorithms used by token keys.
//
// # Overview
//
// Typically, the supported algorithm for a token key can be represented by a
// value of the [SecKeyAlgorithm] enumeration. However, tokens such as Smart
// Cards require that input data for operations take the format of a more
// specific algorithm. For example, a token may accept raw data to generate a
// cryptographic signature, but require that raw data to be formatted
// according to PKCS1 padding rules. To express such a requirement, a
// [TKTokenKeyAlgorithm] object defines a target algorithm and a set of other
// algorithms that were used. In the previous example, the target algorithm is
// `kSecKeyAlgorithmRSASignatureRaw` and the
// `kSecKeyAlgorithmRSASignatureDigestPKCS1v15SHA1` algorithm is also reported
// as being used.
//
// # Determining Algorithm Usage
//
//   - [TKTokenKeyAlgorithm.IsAlgorithm]: Returns whether the specified algorithm is the target operation algorithm.
//   - [TKTokenKeyAlgorithm.SupportsAlgorithm]: Whether the specified algorithm is the target operation algorithm, or one of the other algorithms used.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeyAlgorithm
type TKTokenKeyAlgorithm struct {
	objectivec.Object
}

// TKTokenKeyAlgorithmFromID constructs a [TKTokenKeyAlgorithm] from an objc.ID.
//
// Cryptographic algorithms used by token keys.
func TKTokenKeyAlgorithmFromID(id objc.ID) TKTokenKeyAlgorithm {
	return TKTokenKeyAlgorithm{objectivec.Object{ID: id}}
}

// NOTE: TKTokenKeyAlgorithm adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKTokenKeyAlgorithm] class.
//
// # Determining Algorithm Usage
//
//   - [ITKTokenKeyAlgorithm.IsAlgorithm]: Returns whether the specified algorithm is the target operation algorithm.
//   - [ITKTokenKeyAlgorithm.SupportsAlgorithm]: Whether the specified algorithm is the target operation algorithm, or one of the other algorithms used.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeyAlgorithm
type ITKTokenKeyAlgorithm interface {
	objectivec.IObject

	// Topic: Determining Algorithm Usage

	// Returns whether the specified algorithm is the target operation algorithm.
	IsAlgorithm(algorithm corefoundation.CFStringRef) bool
	// Whether the specified algorithm is the target operation algorithm, or one of the other algorithms used.
	SupportsAlgorithm(algorithm corefoundation.CFStringRef) bool
}

// Init initializes the instance.
func (t TKTokenKeyAlgorithm) Init() TKTokenKeyAlgorithm {
	rv := objc.Send[TKTokenKeyAlgorithm](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKTokenKeyAlgorithm) Autorelease() TKTokenKeyAlgorithm {
	rv := objc.Send[TKTokenKeyAlgorithm](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenKeyAlgorithm creates a new TKTokenKeyAlgorithm instance.
func NewTKTokenKeyAlgorithm() TKTokenKeyAlgorithm {
	class := getTKTokenKeyAlgorithmClass()
	rv := objc.Send[TKTokenKeyAlgorithm](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns whether the specified algorithm is the target operation algorithm.
//
// algorithm: The algorithm to be checked.
//
// # Return Value
//
// true if `algorithm` is the target operation algorithm; otherwise, false.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeyAlgorithm/isAlgorithm(_:)
func (t TKTokenKeyAlgorithm) IsAlgorithm(algorithm corefoundation.CFStringRef) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isAlgorithm:"), algorithm)
	return rv
}

// Whether the specified algorithm is the target operation algorithm, or one
// of the other algorithms used.
//
// algorithm: The algorithm to be checked.
//
// # Return Value
//
// true if `algorithm` is the target operation algorithm or one of the other
// algorithms used; otherwise, false.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeyAlgorithm/supportsAlgorithm(_:)
func (t TKTokenKeyAlgorithm) SupportsAlgorithm(algorithm corefoundation.CFStringRef) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("supportsAlgorithm:"), algorithm)
	return rv
}

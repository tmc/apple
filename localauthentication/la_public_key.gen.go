// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"context"
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [LAPublicKey] class.
var (
	_LAPublicKeyClass     LAPublicKeyClass
	_LAPublicKeyClassOnce sync.Once
)

func getLAPublicKeyClass() LAPublicKeyClass {
	_LAPublicKeyClassOnce.Do(func() {
		_LAPublicKeyClass = LAPublicKeyClass{class: objc.GetClass("LAPublicKey")}
	})
	return _LAPublicKeyClass
}

// GetLAPublicKeyClass returns the class object for LAPublicKey.
func GetLAPublicKeyClass() LAPublicKeyClass {
	return getLAPublicKeyClass()
}

type LAPublicKeyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LAPublicKeyClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LAPublicKeyClass) Alloc() LAPublicKey {
	rv := objc.Send[LAPublicKey](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// The public portion of an asymmetric key pair.
//
// # Checking algorithm support
//
//   - [LAPublicKey.CanEncryptUsingSecKeyAlgorithm]: Checks whether the algorithm you supply is valid for encrypting data with the key.
//   - [LAPublicKey.CanVerifyUsingSecKeyAlgorithm]: Checks whether the algorithm you supply is valid for verifying signatures with the key.
//
// # Performing cryptographic operations
//
//   - [LAPublicKey.EncryptDataSecKeyAlgorithmCompletion]: Encrypts the data you supply with a given algorithm.
//   - [LAPublicKey.ExportBytesWithCompletion]: Exports the data that represents a public key.
//   - [LAPublicKey.VerifyDataSignatureSecKeyAlgorithmCompletion]: Verifies a digital signature for the data you supply.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPublicKey
type LAPublicKey struct {
	objectivec.Object
}

// LAPublicKeyFromID constructs a [LAPublicKey] from an objc.ID.
//
// The public portion of an asymmetric key pair.
func LAPublicKeyFromID(id objc.ID) LAPublicKey {
	return LAPublicKey{objectivec.Object{ID: id}}
}

// NOTE: LAPublicKey adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [LAPublicKey] class.
//
// # Checking algorithm support
//
//   - [ILAPublicKey.CanEncryptUsingSecKeyAlgorithm]: Checks whether the algorithm you supply is valid for encrypting data with the key.
//   - [ILAPublicKey.CanVerifyUsingSecKeyAlgorithm]: Checks whether the algorithm you supply is valid for verifying signatures with the key.
//
// # Performing cryptographic operations
//
//   - [ILAPublicKey.EncryptDataSecKeyAlgorithmCompletion]: Encrypts the data you supply with a given algorithm.
//   - [ILAPublicKey.ExportBytesWithCompletion]: Exports the data that represents a public key.
//   - [ILAPublicKey.VerifyDataSignatureSecKeyAlgorithmCompletion]: Verifies a digital signature for the data you supply.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPublicKey
type ILAPublicKey interface {
	objectivec.IObject

	// Topic: Checking algorithm support

	// Checks whether the algorithm you supply is valid for encrypting data with the key.
	CanEncryptUsingSecKeyAlgorithm(algorithm corefoundation.CFStringRef) bool
	// Checks whether the algorithm you supply is valid for verifying signatures with the key.
	CanVerifyUsingSecKeyAlgorithm(algorithm corefoundation.CFStringRef) bool

	// Topic: Performing cryptographic operations

	// Encrypts the data you supply with a given algorithm.
	EncryptDataSecKeyAlgorithmCompletion(data foundation.INSData, algorithm corefoundation.CFStringRef, handler DataErrorHandler)
	// Exports the data that represents a public key.
	ExportBytesWithCompletion(handler DataErrorHandler)
	// Verifies a digital signature for the data you supply.
	VerifyDataSignatureSecKeyAlgorithmCompletion(signedData foundation.INSData, signature foundation.INSData, algorithm corefoundation.CFStringRef, handler ErrorHandler)
}

// Init initializes the instance.
func (p LAPublicKey) Init() LAPublicKey {
	rv := objc.Send[LAPublicKey](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p LAPublicKey) Autorelease() LAPublicKey {
	rv := objc.Send[LAPublicKey](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewLAPublicKey creates a new LAPublicKey instance.
func NewLAPublicKey() LAPublicKey {
	class := getLAPublicKeyClass()
	rv := objc.Send[LAPublicKey](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Checks whether the algorithm you supply is valid for encrypting data with
// the key.
//
// algorithm: A cryptographic algorithm.
//
// # Return Value
//
// A Boolean value that indicates whether the algorithm you supply is valid
// for encrypting data with the key.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPublicKey/canEncrypt(using:)
func (p LAPublicKey) CanEncryptUsingSecKeyAlgorithm(algorithm corefoundation.CFStringRef) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canEncryptUsingSecKeyAlgorithm:"), algorithm)
	return rv
}

// Checks whether the algorithm you supply is valid for verifying signatures
// with the key.
//
// algorithm: A cryptographic algorithm.
//
// # Return Value
//
// A Boolean value that indicates whether the algorithm you supply is valid
// for verifying signatures with the key.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPublicKey/canVerify(using:)
func (p LAPublicKey) CanVerifyUsingSecKeyAlgorithm(algorithm corefoundation.CFStringRef) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canVerifyUsingSecKeyAlgorithm:"), algorithm)
	return rv
}

// Encrypts the data you supply with a given algorithm.
//
// data: The data to encrypt.
//
// algorithm: The algorithm to use to encrypt the data.
//
// handler: A completion handler to call when the encryption operation completes.
//
// `data`: The encrypted data. `error`: An error object that indicates why the
// encryption failed, or `nil` if it succeeded.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPublicKey/encrypt(_:algorithm:completion:)
func (p LAPublicKey) EncryptDataSecKeyAlgorithmCompletion(data foundation.INSData, algorithm corefoundation.CFStringRef, handler DataErrorHandler) {
	_block2, _ := NewDataErrorBlock(handler)
	objc.Send[objc.ID](p.ID, objc.Sel("encryptData:secKeyAlgorithm:completion:"), data, algorithm, _block2)
}

// Exports the data that represents a public key.
//
// handler: A completion handler to call when the export operation completes.
//
// `data`: The data that represents the public key. `error`: An error object
// that indicates why the export operation failed, or `nil` if it succeeded.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPublicKey/exportBytes(completion:)
func (p LAPublicKey) ExportBytesWithCompletion(handler DataErrorHandler) {
	_block0, _ := NewDataErrorBlock(handler)
	objc.Send[objc.ID](p.ID, objc.Sel("exportBytesWithCompletion:"), _block0)
}

// Verifies a digital signature for the data you supply.
//
// signedData: The signed data.
//
// signature: The signature of the data.
//
// algorithm: An algorithm suitable for verifying signatures with this public key.
//
// handler: A completion handler to call when the verification operation completes.
//
// `error`: An error object that indicates why the verification operation
// failed, or `nil` if it succeeded.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPublicKey/verify(_:signature:algorithm:completion:)
func (p LAPublicKey) VerifyDataSignatureSecKeyAlgorithmCompletion(signedData foundation.INSData, signature foundation.INSData, algorithm corefoundation.CFStringRef, handler ErrorHandler) {
	_block3, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](p.ID, objc.Sel("verifyData:signature:secKeyAlgorithm:completion:"), signedData, signature, algorithm, _block3)
}

// EncryptDataSecKeyAlgorithmCompletionSync is a synchronous wrapper around [LAPublicKey.EncryptDataSecKeyAlgorithmCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (p LAPublicKey) EncryptDataSecKeyAlgorithmCompletionSync(ctx context.Context, data foundation.INSData, algorithm corefoundation.CFStringRef) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	p.EncryptDataSecKeyAlgorithmCompletion(data, algorithm, func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// ExportBytes is a synchronous wrapper around [LAPublicKey.ExportBytesWithCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (p LAPublicKey) ExportBytes(ctx context.Context) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	p.ExportBytesWithCompletion(func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// VerifyDataSignatureSecKeyAlgorithmCompletionSync is a synchronous wrapper around [LAPublicKey.VerifyDataSignatureSecKeyAlgorithmCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (p LAPublicKey) VerifyDataSignatureSecKeyAlgorithmCompletionSync(ctx context.Context, signedData foundation.INSData, signature foundation.INSData, algorithm corefoundation.CFStringRef) error {
	done := make(chan error, 1)
	p.VerifyDataSignatureSecKeyAlgorithmCompletion(signedData, signature, algorithm, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

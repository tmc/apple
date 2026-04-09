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

// The class instance for the [LAPrivateKey] class.
var (
	_LAPrivateKeyClass     LAPrivateKeyClass
	_LAPrivateKeyClassOnce sync.Once
)

func getLAPrivateKeyClass() LAPrivateKeyClass {
	_LAPrivateKeyClassOnce.Do(func() {
		_LAPrivateKeyClass = LAPrivateKeyClass{class: objc.GetClass("LAPrivateKey")}
	})
	return _LAPrivateKeyClass
}

// GetLAPrivateKeyClass returns the class object for LAPrivateKey.
func GetLAPrivateKeyClass() LAPrivateKeyClass {
	return getLAPrivateKeyClass()
}

type LAPrivateKeyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LAPrivateKeyClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LAPrivateKeyClass) Alloc() LAPrivateKey {
	rv := objc.Send[LAPrivateKey](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// The private portion of an asymmetric key pair.
//
// # Accessing the associated public key
//
//   - [LAPrivateKey.PublicKey]: The public key that corresponds with the private key in a key pair.
//
// # Checking algorithm support
//
//   - [LAPrivateKey.CanDecryptUsingSecKeyAlgorithm]: Checks whether the algorithm you supply is valid for decrypting data with the key.
//   - [LAPrivateKey.CanExchangeKeysUsingSecKeyAlgorithm]: Checks whether the algorithm you supply is valid for performing key exchanges.
//   - [LAPrivateKey.CanSignUsingSecKeyAlgorithm]: Checks whether the algorithm you supply is valid for signing data with the key.
//
// # Performing cryptographic operations
//
//   - [LAPrivateKey.DecryptDataSecKeyAlgorithmCompletion]: Decrypts the data you supply with a given algorithm.
//   - [LAPrivateKey.ExchangeKeysWithPublicKeySecKeyAlgorithmSecKeyParametersCompletion]: Performs a Diffie-Hellman style key exchange operation.
//   - [LAPrivateKey.SignDataSecKeyAlgorithmCompletion]: Generates a digital signature for the data you supply.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPrivateKey
type LAPrivateKey struct {
	objectivec.Object
}

// LAPrivateKeyFromID constructs a [LAPrivateKey] from an objc.ID.
//
// The private portion of an asymmetric key pair.
func LAPrivateKeyFromID(id objc.ID) LAPrivateKey {
	return LAPrivateKey{objectivec.Object{ID: id}}
}

// NOTE: LAPrivateKey adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [LAPrivateKey] class.
//
// # Accessing the associated public key
//
//   - [ILAPrivateKey.PublicKey]: The public key that corresponds with the private key in a key pair.
//
// # Checking algorithm support
//
//   - [ILAPrivateKey.CanDecryptUsingSecKeyAlgorithm]: Checks whether the algorithm you supply is valid for decrypting data with the key.
//   - [ILAPrivateKey.CanExchangeKeysUsingSecKeyAlgorithm]: Checks whether the algorithm you supply is valid for performing key exchanges.
//   - [ILAPrivateKey.CanSignUsingSecKeyAlgorithm]: Checks whether the algorithm you supply is valid for signing data with the key.
//
// # Performing cryptographic operations
//
//   - [ILAPrivateKey.DecryptDataSecKeyAlgorithmCompletion]: Decrypts the data you supply with a given algorithm.
//   - [ILAPrivateKey.ExchangeKeysWithPublicKeySecKeyAlgorithmSecKeyParametersCompletion]: Performs a Diffie-Hellman style key exchange operation.
//   - [ILAPrivateKey.SignDataSecKeyAlgorithmCompletion]: Generates a digital signature for the data you supply.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPrivateKey
type ILAPrivateKey interface {
	objectivec.IObject

	// Topic: Accessing the associated public key

	// The public key that corresponds with the private key in a key pair.
	PublicKey() ILAPublicKey

	// Topic: Checking algorithm support

	// Checks whether the algorithm you supply is valid for decrypting data with the key.
	CanDecryptUsingSecKeyAlgorithm(algorithm corefoundation.CFStringRef) bool
	// Checks whether the algorithm you supply is valid for performing key exchanges.
	CanExchangeKeysUsingSecKeyAlgorithm(algorithm corefoundation.CFStringRef) bool
	// Checks whether the algorithm you supply is valid for signing data with the key.
	CanSignUsingSecKeyAlgorithm(algorithm corefoundation.CFStringRef) bool

	// Topic: Performing cryptographic operations

	// Decrypts the data you supply with a given algorithm.
	DecryptDataSecKeyAlgorithmCompletion(data foundation.INSData, algorithm corefoundation.CFStringRef, handler DataErrorHandler)
	// Performs a Diffie-Hellman style key exchange operation.
	ExchangeKeysWithPublicKeySecKeyAlgorithmSecKeyParametersCompletion(publicKey foundation.INSData, algorithm corefoundation.CFStringRef, parameters foundation.INSDictionary, handler DataErrorHandler)
	// Generates a digital signature for the data you supply.
	SignDataSecKeyAlgorithmCompletion(data foundation.INSData, algorithm corefoundation.CFStringRef, handler DataErrorHandler)
}

// Init initializes the instance.
func (p LAPrivateKey) Init() LAPrivateKey {
	rv := objc.Send[LAPrivateKey](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p LAPrivateKey) Autorelease() LAPrivateKey {
	rv := objc.Send[LAPrivateKey](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewLAPrivateKey creates a new LAPrivateKey instance.
func NewLAPrivateKey() LAPrivateKey {
	class := getLAPrivateKeyClass()
	rv := objc.Send[LAPrivateKey](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Checks whether the algorithm you supply is valid for decrypting data with
// the key.
//
// algorithm: A cryptographic algorithm.
//
// # Return Value
//
// A Boolean value that indicates whether the algorithm you supply is valid
// for decrypting data with the key.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPrivateKey/canDecrypt(using:)
func (p LAPrivateKey) CanDecryptUsingSecKeyAlgorithm(algorithm corefoundation.CFStringRef) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canDecryptUsingSecKeyAlgorithm:"), algorithm)
	return rv
}

// Checks whether the algorithm you supply is valid for performing key
// exchanges.
//
// algorithm: A cryptographic algorithm.
//
// # Return Value
//
// A Boolean value that indicates whether the algorithm you supply is valid
// for performing key exchanges.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPrivateKey/canExchangeKeys(using:)
func (p LAPrivateKey) CanExchangeKeysUsingSecKeyAlgorithm(algorithm corefoundation.CFStringRef) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canExchangeKeysUsingSecKeyAlgorithm:"), algorithm)
	return rv
}

// Checks whether the algorithm you supply is valid for signing data with the
// key.
//
// algorithm: A cryptographic algorithm.
//
// # Return Value
//
// A Boolean value that indicates whether the algorithm you supply is valid
// for signing data with the key.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPrivateKey/canSign(using:)
func (p LAPrivateKey) CanSignUsingSecKeyAlgorithm(algorithm corefoundation.CFStringRef) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canSignUsingSecKeyAlgorithm:"), algorithm)
	return rv
}

// Decrypts the data you supply with a given algorithm.
//
// data: The data to decrypt.
//
// algorithm: The algorithm to use to decrypt the data.
//
// handler: A completion handler to call when the decryption operation completes.
//
// `data`: The decrypted data. `error`: An error object that indicates why the
// decryption failed, or `nil` if it succeeded.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPrivateKey/decrypt(_:algorithm:completion:)
func (p LAPrivateKey) DecryptDataSecKeyAlgorithmCompletion(data foundation.INSData, algorithm corefoundation.CFStringRef, handler DataErrorHandler) {
	_block2, _ := NewDataErrorBlock(handler)
	objc.Send[objc.ID](p.ID, objc.Sel("decryptData:secKeyAlgorithm:completion:"), data, algorithm, _block2)
}

// Performs a Diffie-Hellman style key exchange operation.
//
// publicKey: The remote party’s public key.
//
// algorithm: An algorithm suitable for performing this key exchange. For example,
// `ecdhKeyExchangeCofactorX963SHA256`.
//
// parameters: A dictionary with parameters for this key exchange.
//
// handler: A completion handler to call when the key exchange operation completes.
//
// data: The result of the key exchange operation. error: An error object that
// indicates why the key exchange failed, or `nil` if the exchange succeeded.
//
// # Discussion
//
// The algorithm you use determines the parameters in the dictionary that are
// required or optional. For more information, see
// [SecKeyKeyExchangeParameter].
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPrivateKey/exchangeKeys(publicKey:algorithm:parameters:completion:)
//
// [SecKeyKeyExchangeParameter]: https://developer.apple.com/documentation/Security/SecKeyKeyExchangeParameter
func (p LAPrivateKey) ExchangeKeysWithPublicKeySecKeyAlgorithmSecKeyParametersCompletion(publicKey foundation.INSData, algorithm corefoundation.CFStringRef, parameters foundation.INSDictionary, handler DataErrorHandler) {
	_block3, _ := NewDataErrorBlock(handler)
	objc.Send[objc.ID](p.ID, objc.Sel("exchangeKeysWithPublicKey:secKeyAlgorithm:secKeyParameters:completion:"), publicKey, algorithm, parameters, _block3)
}

// Generates a digital signature for the data you supply.
//
// data: The data to sign. The data is usually the digest of applying a
// cryptographic hash function to some actual data.
//
// algorithm: An algorithm suitable for this data signing operation. For example,
// `ecdsaSignatureMessageX962SHA256`.
//
// handler: A completion handler to call when the signing operation completes.
//
// data: The signature of the data you supply. error: An error object that
// indicates why the signing operation failed, or `nil` if it succeeded.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPrivateKey/sign(_:algorithm:completion:)
func (p LAPrivateKey) SignDataSecKeyAlgorithmCompletion(data foundation.INSData, algorithm corefoundation.CFStringRef, handler DataErrorHandler) {
	_block2, _ := NewDataErrorBlock(handler)
	objc.Send[objc.ID](p.ID, objc.Sel("signData:secKeyAlgorithm:completion:"), data, algorithm, _block2)
}

// The public key that corresponds with the private key in a key pair.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPrivateKey/publicKey
func (p LAPrivateKey) PublicKey() ILAPublicKey {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("publicKey"))
	return LAPublicKeyFromID(objc.ID(rv))
}

// DecryptDataSecKeyAlgorithmCompletionSync is a synchronous wrapper around [LAPrivateKey.DecryptDataSecKeyAlgorithmCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (p LAPrivateKey) DecryptDataSecKeyAlgorithmCompletionSync(ctx context.Context, data foundation.INSData, algorithm corefoundation.CFStringRef) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	p.DecryptDataSecKeyAlgorithmCompletion(data, algorithm, func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// ExchangeKeysWithPublicKeySecKeyAlgorithmSecKeyParametersCompletionSync is a synchronous wrapper around [LAPrivateKey.ExchangeKeysWithPublicKeySecKeyAlgorithmSecKeyParametersCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (p LAPrivateKey) ExchangeKeysWithPublicKeySecKeyAlgorithmSecKeyParametersCompletionSync(ctx context.Context, publicKey foundation.INSData, algorithm corefoundation.CFStringRef, parameters foundation.INSDictionary) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	p.ExchangeKeysWithPublicKeySecKeyAlgorithmSecKeyParametersCompletion(publicKey, algorithm, parameters, func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// SignDataSecKeyAlgorithmCompletionSync is a synchronous wrapper around [LAPrivateKey.SignDataSecKeyAlgorithmCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (p LAPrivateKey) SignDataSecKeyAlgorithmCompletionSync(ctx context.Context, data foundation.INSData, algorithm corefoundation.CFStringRef) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	p.SignDataSecKeyAlgorithmCompletion(data, algorithm, func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

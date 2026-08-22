// Code generated from Apple documentation for DeviceCheck. DO NOT EDIT.

package devicecheck

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DCAppAttestService] class.
var (
	_DCAppAttestServiceClass     DCAppAttestServiceClass
	_DCAppAttestServiceClassOnce sync.Once
)

func getDCAppAttestServiceClass() DCAppAttestServiceClass {
	_DCAppAttestServiceClassOnce.Do(func() {
		_DCAppAttestServiceClass = DCAppAttestServiceClass{class: objc.GetClass("DCAppAttestService")}
	})
	return _DCAppAttestServiceClass
}

// GetDCAppAttestServiceClass returns the class object for DCAppAttestService.
func GetDCAppAttestServiceClass() DCAppAttestServiceClass {
	return getDCAppAttestServiceClass()
}

type DCAppAttestServiceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DCAppAttestServiceClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DCAppAttestServiceClass) Alloc() DCAppAttestService {
	rv := objc.Send[DCAppAttestService](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// A service that you use to validate the instance of your app running on a
// device.
//
// # Overview
//
// Use the [DCAppAttestServiceClass.SharedService] instance of the
// [DCAppAttestService] class to assert the legitimacy of a particular
// instance of your app to your server. After ensuring service availability by
// reading the [DCAppAttestService.Supported] property, you use the service
// to:
//
// - Create a cryptographic key in the Secure Enclave by calling the
// [DCAppAttestService.GenerateKeyWithCompletionHandler] method. - Ask Apple
// to certify the key by calling the
// [DCAppAttestService.AttestKeyClientDataHashCompletionHandler] method. -
// Prepare an assertion of your app’s integrity to accompany any or all
// server requests using the
// [DCAppAttestService.GenerateAssertionClientDataHashCompletionHandler]
// method.
//
// For more information about how to support App Attest in your app, see
// [Establishing your app’s integrity]. For information about the
// complementary procedures you implement on your server, see [Validating apps
// that connect to your server].
//
// # Accessing the service
//
//   - [DCAppAttestService.IsSupported]: A Boolean value that indicates whether a particular device provides the App Attest service.
//
// # Preparing a key
//
//   - [DCAppAttestService.GenerateKeyWithCompletionHandler]: Creates a new cryptographic key for use with the App Attest service.
//   - [DCAppAttestService.AttestKeyClientDataHashCompletionHandler]: Asks Apple to attest to the validity of a generated cryptographic key.
//
// # Validating the app instance
//
//   - [DCAppAttestService.GenerateAssertionClientDataHashCompletionHandler]: Creates a block of data that demonstrates the legitimacy of an instance of your app running on a device.
//
// See: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService
//
// [Establishing your app’s integrity]: https://developer.apple.com/documentation/DeviceCheck/establishing-your-app-s-integrity
// [Validating apps that connect to your server]: https://developer.apple.com/documentation/DeviceCheck/validating-apps-that-connect-to-your-server
type DCAppAttestService struct {
	objectivec.Object
}

// DCAppAttestServiceFromID constructs a [DCAppAttestService] from an objc.ID.
//
// A service that you use to validate the instance of your app running on a
// device.
func DCAppAttestServiceFromID(id objc.ID) DCAppAttestService {
	return DCAppAttestService{objectivec.Object{ID: id}}
}

// NOTE: DCAppAttestService adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [DCAppAttestService] class.
//
// # Accessing the service
//
//   - [IDCAppAttestService.IsSupported]: A Boolean value that indicates whether a particular device provides the App Attest service.
//
// # Preparing a key
//
//   - [IDCAppAttestService.GenerateKeyWithCompletionHandler]: Creates a new cryptographic key for use with the App Attest service.
//   - [IDCAppAttestService.AttestKeyClientDataHashCompletionHandler]: Asks Apple to attest to the validity of a generated cryptographic key.
//
// # Validating the app instance
//
//   - [IDCAppAttestService.GenerateAssertionClientDataHashCompletionHandler]: Creates a block of data that demonstrates the legitimacy of an instance of your app running on a device.
//
// See: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService
type IDCAppAttestService interface {
	objectivec.IObject

	// Topic: Accessing the service

	// A Boolean value that indicates whether a particular device provides the App Attest service.
	IsSupported() bool

	// Topic: Preparing a key

	// Creates a new cryptographic key for use with the App Attest service.
	GenerateKeyWithCompletionHandler(completionHandler StringErrorHandler)
	// Asks Apple to attest to the validity of a generated cryptographic key.
	AttestKeyClientDataHashCompletionHandler(keyId string, clientDataHash foundation.NSData, completionHandler DataErrorHandler)

	// Topic: Validating the app instance

	// Creates a block of data that demonstrates the legitimacy of an instance of your app running on a device.
	GenerateAssertionClientDataHashCompletionHandler(keyId string, clientDataHash foundation.NSData, completionHandler DataErrorHandler)
}

// Init initializes the instance.
func (d DCAppAttestService) Init() DCAppAttestService {
	rv := objc.Send[DCAppAttestService](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DCAppAttestService) Autorelease() DCAppAttestService {
	rv := objc.Send[DCAppAttestService](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDCAppAttestService creates a new DCAppAttestService instance.
func NewDCAppAttestService() DCAppAttestService {
	class := getDCAppAttestServiceClass()
	rv := objc.Send[DCAppAttestService](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new cryptographic key for use with the App Attest service.
//
// completionHandler: A closure that the method calls upon completion with the following
// parameters:
//
// - `keyId`: An identifier that you use to refer to the key. The framework
// securely stores the key in the Secure Enclave. - `error`: A [DCError]
// instance that indicates the reason for failure, or `nil` on success.
//
// # Discussion
//
// Call this method to request the creation of a secure, unattested key pair
// on a device for a specific user. On success, the method provides your app
// with an identifier that represents the key pair stored in the Secure
// Enclave. Because there’s no way to use or retrieve the key without the
// identifier, you’ll want to either record it in your app or on your server
// right away. If key generation fails, the closure provides a [DCError] that
// indicates the reason for the failure.
//
// Create a unique key for each user account on a device. Otherwise it’s
// hard to detect an attack that uses a single compromised device to serve
// multiple remote users running a compromised version of your app. For more
// information, see [Assessing fraud risk].
//
// After you get the identifier, you call the
// [DCAppAttestService.AttestKeyClientDataHashCompletionHandler] method with
// the key identifier to ask Apple to attest to the validity of the associated
// key. Later, you call the
// [DCAppAttestService.GenerateAssertionClientDataHashCompletionHandler]
// method with the key identifier to answer a challenge from your server, and
// establish the legitimacy of this instance of your app.
//
// See: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService/generateKey(completionHandler:)
//
// [DCError]: https://developer.apple.com/documentation/DeviceCheck/DCError-swift.struct
// [Assessing fraud risk]: https://developer.apple.com/documentation/DeviceCheck/assessing-fraud-risk
//
// [DCError]: https://developer.apple.com/documentation/DeviceCheck/DCError-swift.struct
func (d DCAppAttestService) GenerateKeyWithCompletionHandler(completionHandler StringErrorHandler) {
	_block0, _ := NewStringErrorBlock(completionHandler)
	objc.Send[objc.ID](d.ID, objc.Sel("generateKeyWithCompletionHandler:"), _block0)
}

// Asks Apple to attest to the validity of a generated cryptographic key.
//
// keyId: The identifier you received when generating a cryptographic key by calling
// the [DCAppAttestService.GenerateKeyWithCompletionHandler] method.
//
// clientDataHash: A SHA256 hash of a unique, single-use data block that embeds a challenge
// from your server. Should be at least 16 bytes in length.
//
// completionHandler: A closure that the method calls upon completion with the following
// parameters:
//
// - `attestationObject`: A statement from Apple about the validity of the key
// associated with `keyId`. Send this to your server for processing. -
// `error`: A [DCError] instance that indicates the reason for failure, or
// `nil` on success.
//
// # Discussion
//
// This method asks Apple to attest to the validity of a key that you
// previously generated with a call to the
// [DCAppAttestService.GenerateKeyWithCompletionHandler] method. Provide the
// method with both the key identifier and a computed hash of a data block
// that includes a one-time challenge from your server to prevent replay
// attacks. For example, you can use CryptoKit to create a [SHA256] hash of
// challenge data:
//
// The attest method calls its completion handler to return an attestation
// object to you, which you must send to your server for verification. A
// compromised version of your app could falsify the verification result, thus
// circumventing App Attest.
//
// If you successfully verify the attestation object on your server, as
// described in [Validating apps that connect to your server], then you can
// associate the key identifier with the user on the device for future
// reference. You’ll need the identifier to generate assertions with calls
// to [DCAppAttestService.GenerateAssertionClientDataHashCompletionHandler].
// If your server fails to verify the attestation object, discard the key
// identifier.
//
// If the method’s completion handler returns the [serverUnavailable] error
// — typically due to network connectivity issues — it means that the
// framework failed to reach the App Attest service to complete the
// attestation. In this case, retry attestation again using the same key and
// client data hash later to avoid unnecessarily generating new keys. Retrying
// with the same inputs helps to preserve the risk metric for a given device.
//
// See: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService/attestKey(_:clientDataHash:completionHandler:)
//
// [DCError]: https://developer.apple.com/documentation/DeviceCheck/DCError-swift.struct
// [SHA256]: https://developer.apple.com/documentation/CryptoKit/SHA256
// [Validating apps that connect to your server]: https://developer.apple.com/documentation/DeviceCheck/validating-apps-that-connect-to-your-server
// [serverUnavailable]: https://developer.apple.com/documentation/DeviceCheck/DCError-swift.struct/serverUnavailable
func (d DCAppAttestService) AttestKeyClientDataHashCompletionHandler(keyId string, clientDataHash foundation.NSData, completionHandler DataErrorHandler) {
	_block2, _ := NewDataErrorBlock(completionHandler)
	objc.Send[objc.ID](d.ID, objc.Sel("attestKey:clientDataHash:completionHandler:"), objc.String(keyId), clientDataHash, _block2)
}

// Creates a block of data that demonstrates the legitimacy of an instance of
// your app running on a device.
//
// keyId: The identifier you received when generating a cryptographic key by calling
// the [DCAppAttestService.GenerateKeyWithCompletionHandler] method.
//
// clientDataHash: A SHA256 hash of a unique, single-use data block that represents the client
// data to be signed with the attested private key. Should be at least 16
// bytes in length.
//
// completionHandler: A closure that the method calls upon completion with the following
// parameters:
//
// - `assertionObject`: A data structure that you send to your server for
// processing. - `error` : A [DCError] instance that indicates the reason for
// failure, or `nil` on success.
//
// # Discussion
//
// After generating a key with the
// [DCAppAttestService.GenerateKeyWithCompletionHandler] method and validating
// it with the [DCAppAttestService.AttestKeyClientDataHashCompletionHandler]
// method, you can use the key at critical moments in your app’s life cycle
// — like when a user tries to access premium content — to reaffirm the
// legitimacy of a given instance of your app. Do this by using the
// [DCAppAttestService.GenerateAssertionClientDataHashCompletionHandler]
// method to sign server requests with your attested key.
//
// You provide the key identifier and a hash of the request that includes a
// challenge from your server to prevent replay attacks, where an attacker
// reuses captured network traffic to pose as someone else. The method returns
// an assertion object in its completion handler that you send to your server
// for verification, as described in [Establishing your app’s integrity].
//
// See: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService/generateAssertion(_:clientDataHash:completionHandler:)
//
// [DCError]: https://developer.apple.com/documentation/DeviceCheck/DCError-swift.struct
// [Establishing your app’s integrity]: https://developer.apple.com/documentation/DeviceCheck/establishing-your-app-s-integrity
func (d DCAppAttestService) GenerateAssertionClientDataHashCompletionHandler(keyId string, clientDataHash foundation.NSData, completionHandler DataErrorHandler) {
	_block2, _ := NewDataErrorBlock(completionHandler)
	objc.Send[objc.ID](d.ID, objc.Sel("generateAssertion:clientDataHash:completionHandler:"), objc.String(keyId), clientDataHash, _block2)
}

// A Boolean value that indicates whether a particular device provides the App
// Attest service.
//
// # Discussion
//
// If you read [DCAppAttestService.Supported] from within an app extension,
// the value might be true or false, depending on the extension type. However,
// most extensions don’t support App Attest. The
// [DCAppAttestService.GenerateKeyWithCompletionHandler] method fails when you
// call it from an app extension, regardless of the value of
// [DCAppAttestService.Supported].
//
// The only app extensions that support App Attest are watchOS extensions in
// watchOS 9 or later. For these extensions, you can use the results from
// [DCAppAttestService.Supported] to indicate whether your WatchKit extension
// bypasses attestation.
//
// See: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService/isSupported
func (d DCAppAttestService) IsSupported() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isSupported"))
	return rv
}

// The shared App Attest service that you use to validate your app.
//
// # Discussion
//
// Use the shared instance of the service to generate and to certify a
// cryptographic key, and then to assert your app’s validity using that key.
//
// See: https://developer.apple.com/documentation/DeviceCheck/DCAppAttestService/shared
func (_DCAppAttestServiceClass DCAppAttestServiceClass) SharedService() DCAppAttestService {
	rv := objc.Send[objc.ID](objc.ID(_DCAppAttestServiceClass.class), objc.Sel("sharedService"))
	return DCAppAttestServiceFromID(objc.ID(rv))
}

// GenerateKey is a synchronous wrapper around [DCAppAttestService.GenerateKeyWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d DCAppAttestService) GenerateKey(ctx context.Context) (*string, error) {
	type result struct {
		val *string
		err error
	}
	done := make(chan result, 1)
	d.GenerateKeyWithCompletionHandler(func(val *string, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// AttestKeyClientDataHash is a synchronous wrapper around [DCAppAttestService.AttestKeyClientDataHashCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d DCAppAttestService) AttestKeyClientDataHash(ctx context.Context, keyId string, clientDataHash foundation.NSData) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	d.AttestKeyClientDataHashCompletionHandler(keyId, clientDataHash, func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// GenerateAssertionClientDataHash is a synchronous wrapper around [DCAppAttestService.GenerateAssertionClientDataHashCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d DCAppAttestService) GenerateAssertionClientDataHash(ctx context.Context, keyId string, clientDataHash foundation.NSData) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	d.GenerateAssertionClientDataHashCompletionHandler(keyId, clientDataHash, func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

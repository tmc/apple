// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVContentKeySession] class.
var (
	_AVContentKeySessionClass     AVContentKeySessionClass
	_AVContentKeySessionClassOnce sync.Once
)

func getAVContentKeySessionClass() AVContentKeySessionClass {
	_AVContentKeySessionClassOnce.Do(func() {
		_AVContentKeySessionClass = AVContentKeySessionClass{class: objc.GetClass("AVContentKeySession")}
	})
	return _AVContentKeySessionClass
}

// GetAVContentKeySessionClass returns the class object for AVContentKeySession.
func GetAVContentKeySessionClass() AVContentKeySessionClass {
	return getAVContentKeySessionClass()
}

type AVContentKeySessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVContentKeySessionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVContentKeySessionClass) Alloc() AVContentKeySession {
	rv := objc.Send[AVContentKeySession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that creates and tracks decryption keys for media data.
//
// # Inspecting the session
//
//   - [AVContentKeySession.KeySystem]: The type of key system used to retrieve keys.
//   - [AVContentKeySession.StorageURL]: A URL that points to a writable storage directory.
//
// # Managing the delegate object
//
//   - [AVContentKeySession.SetDelegateQueue]: Sets the session’s delegate object and the dispatch queue on which to call the delegate’s methods.
//   - [AVContentKeySession.Delegate]: The content key session’s delegate object.
//   - [AVContentKeySession.DelegateQueue]: The dispatch queue the session uses to invoke delegate callbacks.
//
// # Managing content key recipients
//
//   - [AVContentKeySession.ContentKeyRecipients]: An array of content key recipients.
//   - [AVContentKeySession.AddContentKeyRecipient]: Tells the delegate that the specified recipient should have access to the decryption keys loaded with the session.
//   - [AVContentKeySession.RemoveContentKeyRecipient]: Tells the delegate to remove the specified recipient.
//
// # Processing requests
//
//   - [AVContentKeySession.ProcessContentKeyRequestWithIdentifierInitializationDataOptions]: Tells the delegate to start loading the content decryption key with the specified identifier and initialization data.
//
// # Managing expiration
//
//   - [AVContentKeySession.Expire]: Tells the delegate that the session expired as the result of normal, intentional processes.
//   - [AVContentKeySession.MakeSecureTokenForExpirationDateOfPersistableContentKeyCompletionHandler]: Creates a secure server playback context that the client sends to the key server to get an expiration date for the given persistable content key data.
//   - [AVContentKeySession.RenewExpiringResponseDataForContentKeyRequest]: Tells the delegate that previously provided response data for a content key request is about to expire.
//   - [AVContentKeySession.ContentProtectionSessionIdentifier]: The identifier for the current content protection session.
//
// # Invalidating content keys
//
//   - [AVContentKeySession.InvalidatePersistableContentKeyOptionsCompletionHandler]: Invalidates the persistable content key and creates a secure server playback context (SPC) to verify the outcome of an invalidation request.
//   - [AVContentKeySession.InvalidateAllPersistableContentKeysForAppOptionsCompletionHandler]: Invalidates all of an app’s persistable content keys and creates a secure server playback context (SPC) to verify the outcome of an invalidation request.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession
type AVContentKeySession struct {
	objectivec.Object
}

// AVContentKeySessionFromID constructs a [AVContentKeySession] from an objc.ID.
//
// An object that creates and tracks decryption keys for media data.
func AVContentKeySessionFromID(id objc.ID) AVContentKeySession {
	return AVContentKeySession{objectivec.Object{ID: id}}
}

// NOTE: AVContentKeySession adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVContentKeySession] class.
//
// # Inspecting the session
//
//   - [IAVContentKeySession.KeySystem]: The type of key system used to retrieve keys.
//   - [IAVContentKeySession.StorageURL]: A URL that points to a writable storage directory.
//
// # Managing the delegate object
//
//   - [IAVContentKeySession.SetDelegateQueue]: Sets the session’s delegate object and the dispatch queue on which to call the delegate’s methods.
//   - [IAVContentKeySession.Delegate]: The content key session’s delegate object.
//   - [IAVContentKeySession.DelegateQueue]: The dispatch queue the session uses to invoke delegate callbacks.
//
// # Managing content key recipients
//
//   - [IAVContentKeySession.ContentKeyRecipients]: An array of content key recipients.
//   - [IAVContentKeySession.AddContentKeyRecipient]: Tells the delegate that the specified recipient should have access to the decryption keys loaded with the session.
//   - [IAVContentKeySession.RemoveContentKeyRecipient]: Tells the delegate to remove the specified recipient.
//
// # Processing requests
//
//   - [IAVContentKeySession.ProcessContentKeyRequestWithIdentifierInitializationDataOptions]: Tells the delegate to start loading the content decryption key with the specified identifier and initialization data.
//
// # Managing expiration
//
//   - [IAVContentKeySession.Expire]: Tells the delegate that the session expired as the result of normal, intentional processes.
//   - [IAVContentKeySession.MakeSecureTokenForExpirationDateOfPersistableContentKeyCompletionHandler]: Creates a secure server playback context that the client sends to the key server to get an expiration date for the given persistable content key data.
//   - [IAVContentKeySession.RenewExpiringResponseDataForContentKeyRequest]: Tells the delegate that previously provided response data for a content key request is about to expire.
//   - [IAVContentKeySession.ContentProtectionSessionIdentifier]: The identifier for the current content protection session.
//
// # Invalidating content keys
//
//   - [IAVContentKeySession.InvalidatePersistableContentKeyOptionsCompletionHandler]: Invalidates the persistable content key and creates a secure server playback context (SPC) to verify the outcome of an invalidation request.
//   - [IAVContentKeySession.InvalidateAllPersistableContentKeysForAppOptionsCompletionHandler]: Invalidates all of an app’s persistable content keys and creates a secure server playback context (SPC) to verify the outcome of an invalidation request.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession
type IAVContentKeySession interface {
	objectivec.IObject

	// Topic: Inspecting the session

	// The type of key system used to retrieve keys.
	KeySystem() AVContentKeySystem
	// A URL that points to a writable storage directory.
	StorageURL() foundation.INSURL

	// Topic: Managing the delegate object

	// Sets the session’s delegate object and the dispatch queue on which to call the delegate’s methods.
	SetDelegateQueue(delegate AVContentKeySessionDelegate, delegateQueue dispatch.Queue)
	// The content key session’s delegate object.
	Delegate() AVContentKeySessionDelegate
	// The dispatch queue the session uses to invoke delegate callbacks.
	DelegateQueue() dispatch.Queue

	// Topic: Managing content key recipients

	// An array of content key recipients.
	ContentKeyRecipients() []objectivec.IObject
	// Tells the delegate that the specified recipient should have access to the decryption keys loaded with the session.
	AddContentKeyRecipient(recipient AVContentKeyRecipient)
	// Tells the delegate to remove the specified recipient.
	RemoveContentKeyRecipient(recipient AVContentKeyRecipient)

	// Topic: Processing requests

	// Tells the delegate to start loading the content decryption key with the specified identifier and initialization data.
	ProcessContentKeyRequestWithIdentifierInitializationDataOptions(identifier objectivec.IObject, initializationData foundation.INSData, options foundation.INSDictionary)

	// Topic: Managing expiration

	// Tells the delegate that the session expired as the result of normal, intentional processes.
	Expire()
	// Creates a secure server playback context that the client sends to the key server to get an expiration date for the given persistable content key data.
	MakeSecureTokenForExpirationDateOfPersistableContentKeyCompletionHandler(persistableContentKeyData foundation.INSData, handler DataErrorHandler)
	// Tells the delegate that previously provided response data for a content key request is about to expire.
	RenewExpiringResponseDataForContentKeyRequest(contentKeyRequest IAVContentKeyRequest)
	// The identifier for the current content protection session.
	ContentProtectionSessionIdentifier() foundation.INSData

	// Topic: Invalidating content keys

	// Invalidates the persistable content key and creates a secure server playback context (SPC) to verify the outcome of an invalidation request.
	InvalidatePersistableContentKeyOptionsCompletionHandler(persistableContentKeyData foundation.INSData, options foundation.INSDictionary, handler DataErrorHandler)
	// Invalidates all of an app’s persistable content keys and creates a secure server playback context (SPC) to verify the outcome of an invalidation request.
	InvalidateAllPersistableContentKeysForAppOptionsCompletionHandler(appIdentifier foundation.INSData, options foundation.INSDictionary, handler DataErrorHandler)

	// Boolean indicating whether advisory keys are enabled on the client.
	SupportsAdvisoryKeys() bool
	SetSupportsAdvisoryKeys(value bool)
}

// Init initializes the instance.
func (c AVContentKeySession) Init() AVContentKeySession {
	rv := objc.Send[AVContentKeySession](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVContentKeySession) Autorelease() AVContentKeySession {
	rv := objc.Send[AVContentKeySession](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVContentKeySession creates a new AVContentKeySession instance.
func NewAVContentKeySession() AVContentKeySession {
	class := getAVContentKeySessionClass()
	rv := objc.Send[AVContentKeySession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a content key session to manage a collection of content decryption
// keys.
//
// keySystem: A valid key system used to retrieve keys.
//
// # Return Value
//
// Returns a new [AVContentKeySession] instance.
//
// # Discussion
//
// The [AVContentKeySession] instance returned is capable of managing a
// collection of content decryption keys that correspond to the input key
// system. An [InvalidArgumentException] is raised when the value of
// `keySystem` is unsupported.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/init(keySystem:)
func NewContentKeySessionWithKeySystem(keySystem AVContentKeySystem) AVContentKeySession {
	rv := objc.Send[objc.ID](objc.ID(getAVContentKeySessionClass().class), objc.Sel("contentKeySessionWithKeySystem:"), objc.String(string(keySystem)))
	return AVContentKeySessionFromID(rv)
}

// Creates a content key session to manage a collection of content decryption
// keys; points to a directory that stores abnormal session termination
// reports.
//
// keySystem: A valid key system used to retrieve keys.
//
// storageURL: A URL that points to a writable directory. The session uses the directory
// to facilitate expired session reports after an abnormal session
// termination.
//
// # Return Value
//
// Returns a new AVContentKeySession instance.
//
// # Discussion
//
// The [AVContentKeySession] instance returned is capable of managing a
// collection of content decryption keys that correspond to the input key
// system. An [InvalidArgumentException] is raised when the value of
// `keySystem` is unsupported.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/init(keySystem:storageDirectoryAt:)
func NewContentKeySessionWithKeySystemStorageDirectoryAtURL(keySystem AVContentKeySystem, storageURL foundation.INSURL) AVContentKeySession {
	rv := objc.Send[objc.ID](objc.ID(getAVContentKeySessionClass().class), objc.Sel("contentKeySessionWithKeySystem:storageDirectoryAtURL:"), objc.String(string(keySystem)), storageURL)
	return AVContentKeySessionFromID(rv)
}

// Sets the session’s delegate object and the dispatch queue on which to
// call the delegate’s methods.
//
// delegate: An object that conforms to the [AVContentKeySessionDelegate] protocol.
//
// delegateQueue: The dispatch queue on which the session calls the delegate object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/setDelegate(_:queue:)
func (c AVContentKeySession) SetDelegateQueue(delegate AVContentKeySessionDelegate, delegateQueue dispatch.Queue) {
	objc.Send[objc.ID](c.ID, objc.Sel("setDelegate:queue:"), delegate, uintptr(delegateQueue.Handle()))
}

// Tells the delegate that the specified recipient should have access to the
// decryption keys loaded with the session.
//
// recipient: The content key recipient to use for the session.
//
// # Discussion
//
// Don’t add a recipient to a session that has expired or had already begun
// to process media data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/addContentKeyRecipient(_:)
func (c AVContentKeySession) AddContentKeyRecipient(recipient AVContentKeyRecipient) {
	objc.Send[objc.ID](c.ID, objc.Sel("addContentKeyRecipient:"), recipient)
}

// Tells the delegate to remove the specified recipient.
//
// recipient: The content key recipient to remove.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/removeContentKeyRecipient(_:)
func (c AVContentKeySession) RemoveContentKeyRecipient(recipient AVContentKeyRecipient) {
	objc.Send[objc.ID](c.ID, objc.Sel("removeContentKeyRecipient:"), recipient)
}

// Tells the delegate to start loading the content decryption key with the
// specified identifier and initialization data.
//
// identifier: The container- and protocol-specific identifier used to obtain a key
// response.
//
// initializationData: The container- and protocol-specific data used to obtain a key response.
//
// options: No options are currently defined. Set this value to `nil`.
//
// # Discussion
//
// Either the `identifier` or `initializationData` parameters must be
// non-`nil`. If required by the protocol, both parameters can be non-`nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/processContentKeyRequest(withIdentifier:initializationData:options:)
func (c AVContentKeySession) ProcessContentKeyRequestWithIdentifierInitializationDataOptions(identifier objectivec.IObject, initializationData foundation.INSData, options foundation.INSDictionary) {
	objc.Send[objc.ID](c.ID, objc.Sel("processContentKeyRequestWithIdentifier:initializationData:options:"), identifier, initializationData, options)
}

// Tells the delegate that the session expired as the result of normal,
// intentional processes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/expire()
func (c AVContentKeySession) Expire() {
	objc.Send[objc.ID](c.ID, objc.Sel("expire"))
}

// Creates a secure server playback context that the client sends to the key
// server to get an expiration date for the given persistable content key
// data.
//
// persistableContentKeyData: The previously created persistable content key data.
//
// handler: A block called after the secure token is ready.
//
// secureTokenData: The new secure token. error: A parameter that holds the
// error object that explains the error. If no error occurred, the value of
// this parameter is `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/makeSecureTokenForExpirationDate(ofPersistableContentKey:completionHandler:)
func (c AVContentKeySession) MakeSecureTokenForExpirationDateOfPersistableContentKeyCompletionHandler(persistableContentKeyData foundation.INSData, handler DataErrorHandler) {
	_block1, _ := NewDataErrorBlock(handler)
	objc.Send[objc.ID](c.ID, objc.Sel("makeSecureTokenForExpirationDateOfPersistableContentKey:completionHandler:"), persistableContentKeyData, _block1)
}

// Tells the delegate that previously provided response data for a content key
// request is about to expire.
//
// contentKeyRequest: The content key request that’s about to expire.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/renewExpiringResponseData(for:)
func (c AVContentKeySession) RenewExpiringResponseDataForContentKeyRequest(contentKeyRequest IAVContentKeyRequest) {
	objc.Send[objc.ID](c.ID, objc.Sel("renewExpiringResponseDataForContentKeyRequest:"), contentKeyRequest)
}

// Invalidates the persistable content key and creates a secure server
// playback context (SPC) to verify the outcome of an invalidation request.
//
// persistableContentKeyData: The persistable content key data to invalidate.
//
// options: Additional options to use when generating the server playback context. Pass
// `nil` to indicate no additional options.
//
// handler: The completion handler callback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/invalidatePersistableContentKey(_:options:completionHandler:)
func (c AVContentKeySession) InvalidatePersistableContentKeyOptionsCompletionHandler(persistableContentKeyData foundation.INSData, options foundation.INSDictionary, handler DataErrorHandler) {
	_block2, _ := NewDataErrorBlock(handler)
	objc.Send[objc.ID](c.ID, objc.Sel("invalidatePersistableContentKey:options:completionHandler:"), persistableContentKeyData, options, _block2)
}

// Invalidates all of an app’s persistable content keys and creates a secure
// server playback context (SPC) to verify the outcome of an invalidation
// request.
//
// appIdentifier: An opaque identifier for the app.
//
// options: Additional data necessary to generate the server playback context. Pass
// `nil` to indicate no additional options.
//
// See [AVContentKeySessionServerPlaybackContextOption] for supported options.
//
// handler: The completion handler callback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/invalidateAllPersistableContentKeys(forApp:options:completionHandler:)
func (c AVContentKeySession) InvalidateAllPersistableContentKeysForAppOptionsCompletionHandler(appIdentifier foundation.INSData, options foundation.INSDictionary, handler DataErrorHandler) {
	_block2, _ := NewDataErrorBlock(handler)
	objc.Send[objc.ID](c.ID, objc.Sel("invalidateAllPersistableContentKeysForApp:options:completionHandler:"), appIdentifier, options, _block2)
}

// Returns the expired session reports for content key sessions created with
// the specified app identifier.
//
// appIdentifier: The opaque identifier for the app.
//
// storageURL: The URL that points to the directory containing expired session reports.
//
// # Return Value
//
// Returns an array of expired session reports.
//
// # Discussion
//
// The system only returns expired session reports. It doesn’t include
// reports for active sessions.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/pendingExpiredSessionReports(withAppIdentifier:storageDirectoryAt:)
func (_AVContentKeySessionClass AVContentKeySessionClass) PendingExpiredSessionReportsWithAppIdentifierStorageDirectoryAtURL(appIdentifier foundation.INSData, storageURL foundation.INSURL) []foundation.NSData {
	rv := objc.Send[[]objc.ID](objc.ID(_AVContentKeySessionClass.class), objc.Sel("pendingExpiredSessionReportsWithAppIdentifier:storageDirectoryAtURL:"), appIdentifier, storageURL)
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSData {
		return foundation.NSDataFromID(id)
	})
}

// Removes expired session reports from storage.
//
// expiredSessionReports: An array of expired session reports to delete.
//
// appIdentifier: The opaque identifier for the app.
//
// storageURL: The URL that points to the directory containing expired session reports.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/removePendingExpiredSessionReports(_:withAppIdentifier:storageDirectoryAt:)
func (_AVContentKeySessionClass AVContentKeySessionClass) RemovePendingExpiredSessionReportsWithAppIdentifierStorageDirectoryAtURL(expiredSessionReports []foundation.NSData, appIdentifier foundation.INSData, storageURL foundation.INSURL) {
	objc.Send[objc.ID](objc.ID(_AVContentKeySessionClass.class), objc.Sel("removePendingExpiredSessionReports:withAppIdentifier:storageDirectoryAtURL:"), objectivec.IObjectSliceToNSArray(expiredSessionReports), appIdentifier, storageURL)
}

// The type of key system used to retrieve keys.
//
// # Discussion
//
// Valid values for keySystem are [fairPlayStreaming] and [clearKey].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/keySystem
//
// [clearKey]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySystem/clearKey
// [fairPlayStreaming]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySystem/fairPlayStreaming
func (c AVContentKeySession) KeySystem() AVContentKeySystem {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("keySystem"))
	return AVContentKeySystem(foundation.NSStringFromID(rv).String())
}

// A URL that points to a writable storage directory.
//
// # Discussion
//
// The writable directory stores expired session reports.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/storageURL
func (c AVContentKeySession) StorageURL() foundation.INSURL {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("storageURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The content key session’s delegate object.
//
// # Discussion
//
// Set the session’s delegate using the [SetDelegateQueue] method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/delegate
func (c AVContentKeySession) Delegate() AVContentKeySessionDelegate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("delegate"))
	return AVContentKeySessionDelegateObjectFromID(rv)
}

// The dispatch queue the session uses to invoke delegate callbacks.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/delegateQueue
func (c AVContentKeySession) DelegateQueue() dispatch.Queue {
	rv := objc.Send[uintptr](c.ID, objc.Sel("delegateQueue"))
	return dispatch.QueueFromHandle(rv)
}

// An array of content key recipients.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/contentKeyRecipients
func (c AVContentKeySession) ContentKeyRecipients() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("contentKeyRecipients"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// The identifier for the current content protection session.
//
// # Discussion
//
// The content protection session identifier is a unique string the session
// generates.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/contentProtectionSessionIdentifier
func (c AVContentKeySession) ContentProtectionSessionIdentifier() foundation.INSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("contentProtectionSessionIdentifier"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// Boolean indicating whether advisory keys are enabled on the client.
//
// # Discussion
//
// Set to true to enable advisory key loading, false to disable. false by
// default.
//
// Advisory key loading allows applications to make use of content keys
// provided speculatively by the key server. When enabled, FairPlay may cache
// these keys and return them immediately on subsequent requests without
// requiring a round-trip to the key server.
//
// The delegate must be prepared to handle advisory key requests by checking
// the `canBeFulfilledWithAdvisoryKey` property on [AVContentKeyRequest]
// objects.
//
// When an advisory key is already cached by FairPlay,
// `makeStreamingContentKeyRequestData` will return nil for the SPC data, and
// `canBeFulfilledWithAdvisoryKey` will return true. In this case, no request
// to the key server is necessary.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySession/supportsAdvisoryKeys
func (c AVContentKeySession) SupportsAdvisoryKeys() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("supportsAdvisoryKeys"))
	return rv
}
func (c AVContentKeySession) SetSupportsAdvisoryKeys(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setSupportsAdvisoryKeys:"), value)
}

// MakeSecureTokenForExpirationDateOfPersistableContentKey is a synchronous wrapper around [AVContentKeySession.MakeSecureTokenForExpirationDateOfPersistableContentKeyCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c AVContentKeySession) MakeSecureTokenForExpirationDateOfPersistableContentKey(ctx context.Context, persistableContentKeyData foundation.INSData) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	c.MakeSecureTokenForExpirationDateOfPersistableContentKeyCompletionHandler(persistableContentKeyData, func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// InvalidatePersistableContentKeyOptions is a synchronous wrapper around [AVContentKeySession.InvalidatePersistableContentKeyOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c AVContentKeySession) InvalidatePersistableContentKeyOptions(ctx context.Context, persistableContentKeyData foundation.INSData, options foundation.INSDictionary) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	c.InvalidatePersistableContentKeyOptionsCompletionHandler(persistableContentKeyData, options, func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// InvalidateAllPersistableContentKeysForAppOptions is a synchronous wrapper around [AVContentKeySession.InvalidateAllPersistableContentKeysForAppOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c AVContentKeySession) InvalidateAllPersistableContentKeysForAppOptions(ctx context.Context, appIdentifier foundation.INSData, options foundation.INSDictionary) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	c.InvalidateAllPersistableContentKeysForAppOptionsCompletionHandler(appIdentifier, options, func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

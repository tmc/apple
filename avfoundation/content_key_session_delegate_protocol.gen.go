// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A protocol that handles content key requests.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySessionDelegate
type AVContentKeySessionDelegate interface {
	objectivec.IObject

	// Provides the receiver with a new content key request object.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySessionDelegate/contentKeySession(_:didProvide:)-3coq5
	ContentKeySessionDidProvideContentKeyRequest(session IAVContentKeySession, keyRequest IAVContentKeyRequest)
}

// AVContentKeySessionDelegateObject wraps an existing Objective-C object that conforms to the AVContentKeySessionDelegate protocol.
type AVContentKeySessionDelegateObject struct {
	objectivec.Object
}

func (o AVContentKeySessionDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVContentKeySessionDelegateObjectFromID constructs a [AVContentKeySessionDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVContentKeySessionDelegateObjectFromID(id objc.ID) AVContentKeySessionDelegateObject {
	return AVContentKeySessionDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Provides the receiver with a new content key request object.
//
// session: The content key session that is providing the new content key request.
//
// keyRequest: The request for a new content key.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySessionDelegate/contentKeySession(_:didProvide:)-3coq5
func (o AVContentKeySessionDelegateObject) ContentKeySessionDidProvideContentKeyRequest(session IAVContentKeySession, keyRequest IAVContentKeyRequest) {
	objc.Send[struct{}](o.ID, objc.Sel("contentKeySession:didProvideContentKeyRequest:"), session, keyRequest)
}

// Provides the receiver with a new content key request object for the renewal
// of an existing content key.
//
// session: The content key session that is providing the new content key request.
//
// keyRequest: The request for the renewal of a previous content key.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySessionDelegate/contentKeySession(_:didProvideRenewingContentKeyRequest:)
func (o AVContentKeySessionDelegateObject) ContentKeySessionDidProvideRenewingContentKeyRequest(session IAVContentKeySession, keyRequest IAVContentKeyRequest) {
	objc.Send[struct{}](o.ID, objc.Sel("contentKeySession:didProvideRenewingContentKeyRequest:"), session, keyRequest)
}

// Provides the receiver with a new content key request object to process a
// persistable content key.
//
// session: The content key session that is providing the new persistable content key
// request.
//
// keyRequest: The request for a new persistable content key.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySessionDelegate/contentKeySession(_:didProvide:)-2wdgz
func (o AVContentKeySessionDelegateObject) ContentKeySessionDidProvidePersistableContentKeyRequest(session IAVContentKeySession, keyRequest IAVPersistableContentKeyRequest) {
	objc.Send[struct{}](o.ID, objc.Sel("contentKeySession:didProvidePersistableContentKeyRequest:"), session, keyRequest)
}

// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySessionDelegate/contentKeySession(_:didProvide:forInitializationData:)
func (o AVContentKeySessionDelegateObject) ContentKeySessionDidProvideContentKeyRequestsForInitializationData(session IAVContentKeySession, keyRequests []AVContentKeyRequest, initializationData foundation.INSData) {
	objc.Send[struct{}](o.ID, objc.Sel("contentKeySession:didProvideContentKeyRequests:forInitializationData:"), session, objectivec.IObjectSliceToNSArray(keyRequests), initializationData)
}

// Tells the delegate when external protection state has changed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySessionDelegate/contentKeySession(_:externalProtectionStatusDidChangeFor:)
func (o AVContentKeySessionDelegateObject) ContentKeySessionExternalProtectionStatusDidChangeForContentKey(session IAVContentKeySession, contentKey IAVContentKey) {
	objc.Send[struct{}](o.ID, objc.Sel("contentKeySession:externalProtectionStatusDidChangeForContentKey:"), session, contentKey)
}

// Provides the receiver with an updated persistable content key for a
// specific key request.
//
// session: The content key session that is providing the updated persistable content
// key.
//
// persistableContentKey: The updated persistent content key data. This data can be stored offline
// and used to answer future content key requests with the matching key
// identifier.
//
// keyIdentifier: A container- and protocol-specific identifier for the updated persistent
// content key.
//
// # Discussion
//
// If the content key session provides updated persistable content key data,
// previous key data is no longer valid and cannot be used to answer future
// loading requests.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySessionDelegate/contentKeySession(_:didUpdatePersistableContentKey:forContentKeyIdentifier:)
func (o AVContentKeySessionDelegateObject) ContentKeySessionDidUpdatePersistableContentKeyForContentKeyIdentifier(session IAVContentKeySession, persistableContentKey foundation.INSData, keyIdentifier objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("contentKeySession:didUpdatePersistableContentKey:forContentKeyIdentifier:"), session, persistableContentKey, keyIdentifier)
}

// Provides the receiver with a content key request object to retry.
//
// session: The content key session that is providing the content key request.
//
// keyRequest: The key request to be retried.
//
// retryReason: The reason for the retry request.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySessionDelegate/contentKeySession(_:shouldRetry:reason:)
func (o AVContentKeySessionDelegateObject) ContentKeySessionShouldRetryContentKeyRequestReason(session IAVContentKeySession, keyRequest IAVContentKeyRequest, retryReason AVContentKeyRequestRetryReason) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("contentKeySession:shouldRetryContentKeyRequest:reason:"), session, keyRequest, objc.String(string(retryReason)))
	return rv
}

// Tells the receiver the content protection session identifier changed.
//
// session: The content key session to be notified.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySessionDelegate/contentKeySessionContentProtectionSessionIdentifierDidChange(_:)
func (o AVContentKeySessionDelegateObject) ContentKeySessionContentProtectionSessionIdentifierDidChange(session IAVContentKeySession) {
	objc.Send[struct{}](o.ID, objc.Sel("contentKeySessionContentProtectionSessionIdentifierDidChange:"), session)
}

// Tells the receiver that the content key request failed.
//
// session: The content key session that initiated the content key request.
//
// keyRequest: The content key request that failed.
//
// err: An instance of [NSError] that describes the failure that occurred.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySessionDelegate/contentKeySession(_:contentKeyRequest:didFailWithError:)
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
func (o AVContentKeySessionDelegateObject) ContentKeySessionContentKeyRequestDidFailWithError(session IAVContentKeySession, keyRequest IAVContentKeyRequest, err foundation.INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("contentKeySession:contentKeyRequest:didFailWithError:"), session, keyRequest, err)
}

// Tells the content key session that the response to a content key requeset
// was successfully processed.
//
// session: The [AVContentKeySession] instance that initiated the content key request.
//
// keyRequest: The [AVContentKeyRequest] instance whose response was successfully
// processed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySessionDelegate/contentKeySession(_:contentKeyRequestDidSucceed:)
func (o AVContentKeySessionDelegateObject) ContentKeySessionContentKeyRequestDidSucceed(session IAVContentKeySession, keyRequest IAVContentKeyRequest) {
	objc.Send[struct{}](o.ID, objc.Sel("contentKeySession:contentKeyRequestDidSucceed:"), session, keyRequest)
}

// Notifies the sender that an expired session report has been generated.
//
// session: The [AVContentKeySession] object that recorded the generation of an expired
// session report.
//
// # Discussion
//
// This method is invoked when an expired session report is added to the
// [StorageURL] property.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySessionDelegate/contentKeySessionDidGenerateExpiredSessionReport(_:)
func (o AVContentKeySessionDelegateObject) ContentKeySessionDidGenerateExpiredSessionReport(session IAVContentKeySession) {
	objc.Send[struct{}](o.ID, objc.Sel("contentKeySessionDidGenerateExpiredSessionReport:"), session)
}

// AVContentKeySessionDelegateConfig holds optional typed callbacks for [AVContentKeySessionDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate
type AVContentKeySessionDelegateConfig struct {

	// Providing new content key requests
	// ContentKeySessionDidProvideRenewingContentKeyRequest — Provides the receiver with a new content key request object for the renewal of an existing content key.
	ContentKeySessionDidProvideRenewingContentKeyRequest func(session AVContentKeySession, keyRequest AVContentKeyRequest)

	// Updating and retrying content key requests
	// ContentKeySessionContentProtectionSessionIdentifierDidChange — Tells the receiver the content protection session identifier changed.
	ContentKeySessionContentProtectionSessionIdentifierDidChange func(session AVContentKeySession)
	// ContentKeySessionContentKeyRequestDidFailWithError — Tells the receiver that the content key request failed.
	ContentKeySessionContentKeyRequestDidFailWithError func(session AVContentKeySession, keyRequest AVContentKeyRequest, err foundation.NSError)
	// ContentKeySessionContentKeyRequestDidSucceed — Tells the content key session that the response to a content key requeset was successfully processed.
	ContentKeySessionContentKeyRequestDidSucceed func(session AVContentKeySession, keyRequest AVContentKeyRequest)
	// ContentKeySessionDidGenerateExpiredSessionReport — Notifies the sender that an expired session report has been generated.
	ContentKeySessionDidGenerateExpiredSessionReport func(session AVContentKeySession)

	// Other Methods
	// ContentKeySessionDidProvideContentKeyRequest — Provides the receiver with a new content key request object.
	ContentKeySessionDidProvideContentKeyRequest func(session AVContentKeySession, keyRequest AVContentKeyRequest)
	// ContentKeySessionDidProvidePersistableContentKeyRequest — Provides the receiver with a new content key request object to process a persistable content key.
	ContentKeySessionDidProvidePersistableContentKeyRequest func(session AVContentKeySession, keyRequest AVPersistableContentKeyRequest)
	// ContentKeySessionExternalProtectionStatusDidChangeForContentKey — Tells the delegate when external protection state has changed.
	ContentKeySessionExternalProtectionStatusDidChangeForContentKey func(session AVContentKeySession, contentKey AVContentKey)
	// ContentKeySessionShouldRetryContentKeyRequestReason — Provides the receiver with a content key request object to retry.
	ContentKeySessionShouldRetryContentKeyRequestReason func(session AVContentKeySession, keyRequest AVContentKeyRequest, retryReason AVContentKeyRequestRetryReason) bool
}

// NewAVContentKeySessionDelegate creates an Objective-C object implementing the [AVContentKeySessionDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVContentKeySessionDelegateObject] satisfies the [AVContentKeySessionDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcontentkeysessiondelegate
func NewAVContentKeySessionDelegate(config AVContentKeySessionDelegateConfig) AVContentKeySessionDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVContentKeySessionDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ContentKeySessionDidProvideContentKeyRequest != nil {
		fn := config.ContentKeySessionDidProvideContentKeyRequest
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("contentKeySession:didProvideContentKeyRequest:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, keyRequestID objc.ID) {
				session := AVContentKeySessionFromID(sessionID)
				keyRequest := AVContentKeyRequestFromID(keyRequestID)
				fn(session, keyRequest)
			},
		})
	}

	if config.ContentKeySessionDidProvideRenewingContentKeyRequest != nil {
		fn := config.ContentKeySessionDidProvideRenewingContentKeyRequest
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("contentKeySession:didProvideRenewingContentKeyRequest:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, keyRequestID objc.ID) {
				session := AVContentKeySessionFromID(sessionID)
				keyRequest := AVContentKeyRequestFromID(keyRequestID)
				fn(session, keyRequest)
			},
		})
	}

	if config.ContentKeySessionDidProvidePersistableContentKeyRequest != nil {
		fn := config.ContentKeySessionDidProvidePersistableContentKeyRequest
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("contentKeySession:didProvidePersistableContentKeyRequest:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, keyRequestID objc.ID) {
				session := AVContentKeySessionFromID(sessionID)
				keyRequest := AVPersistableContentKeyRequestFromID(keyRequestID)
				fn(session, keyRequest)
			},
		})
	}

	if config.ContentKeySessionExternalProtectionStatusDidChangeForContentKey != nil {
		fn := config.ContentKeySessionExternalProtectionStatusDidChangeForContentKey
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("contentKeySession:externalProtectionStatusDidChangeForContentKey:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, contentKeyID objc.ID) {
				session := AVContentKeySessionFromID(sessionID)
				contentKey := AVContentKeyFromID(contentKeyID)
				fn(session, contentKey)
			},
		})
	}

	if config.ContentKeySessionShouldRetryContentKeyRequestReason != nil {
		fn := config.ContentKeySessionShouldRetryContentKeyRequestReason
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("contentKeySession:shouldRetryContentKeyRequest:reason:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, keyRequestID objc.ID, retryReason AVContentKeyRequestRetryReason) bool {
				session := AVContentKeySessionFromID(sessionID)
				keyRequest := AVContentKeyRequestFromID(keyRequestID)
				return fn(session, keyRequest, retryReason)
			},
		})
	}

	if config.ContentKeySessionContentProtectionSessionIdentifierDidChange != nil {
		fn := config.ContentKeySessionContentProtectionSessionIdentifierDidChange
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("contentKeySessionContentProtectionSessionIdentifierDidChange:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID) {
				session := AVContentKeySessionFromID(sessionID)
				fn(session)
			},
		})
	}

	if config.ContentKeySessionContentKeyRequestDidFailWithError != nil {
		fn := config.ContentKeySessionContentKeyRequestDidFailWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("contentKeySession:contentKeyRequest:didFailWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, keyRequestID objc.ID, errID objc.ID) {
				session := AVContentKeySessionFromID(sessionID)
				keyRequest := AVContentKeyRequestFromID(keyRequestID)
				err := foundation.NSErrorFromID(errID)
				fn(session, keyRequest, err)
			},
		})
	}

	if config.ContentKeySessionContentKeyRequestDidSucceed != nil {
		fn := config.ContentKeySessionContentKeyRequestDidSucceed
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("contentKeySession:contentKeyRequestDidSucceed:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, keyRequestID objc.ID) {
				session := AVContentKeySessionFromID(sessionID)
				keyRequest := AVContentKeyRequestFromID(keyRequestID)
				fn(session, keyRequest)
			},
		})
	}

	if config.ContentKeySessionDidGenerateExpiredSessionReport != nil {
		fn := config.ContentKeySessionDidGenerateExpiredSessionReport
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("contentKeySessionDidGenerateExpiredSessionReport:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID) {
				session := AVContentKeySessionFromID(sessionID)
				fn(session)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AVContentKeySessionDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVContentKeySessionDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVContentKeySessionDelegateObjectFromID(instance)
}

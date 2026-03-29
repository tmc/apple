// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVContentKeyRequest] class.
var (
	_AVContentKeyRequestClass     AVContentKeyRequestClass
	_AVContentKeyRequestClassOnce sync.Once
)

func getAVContentKeyRequestClass() AVContentKeyRequestClass {
	_AVContentKeyRequestClassOnce.Do(func() {
		_AVContentKeyRequestClass = AVContentKeyRequestClass{class: objc.GetClass("AVContentKeyRequest")}
	})
	return _AVContentKeyRequestClass
}

// GetAVContentKeyRequestClass returns the class object for AVContentKeyRequest.
func GetAVContentKeyRequestClass() AVContentKeyRequestClass {
	return getAVContentKeyRequestClass()
}

type AVContentKeyRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVContentKeyRequestClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVContentKeyRequestClass) Alloc() AVContentKeyRequest {
	rv := objc.Send[AVContentKeyRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that encapsulates information about a content decryption key
// request issued from a content key session object.
//
// # Getting content key request data
//
//   - [AVContentKeyRequest.MakeStreamingContentKeyRequestDataForAppContentIdentifierOptionsCompletionHandler]: Obtains encrypted key request data for a specific combination of app and content.
//   - [AVContentKeyRequest.AVContentKeyRequestProtocolVersionsKey]: A key that specifies the versions of the content protection protocol supported by the application.
//   - [AVContentKeyRequest.AVContentKeyRequestRequiresValidationDataInSecureTokenKey]: A key that requires the secure token to have extended validation data.
//   - [AVContentKeyRequest.AVContentKeyRequestRandomDeviceIdentifierSeedKey]: Value is an NSData containing a 16-byte seed to randomize the user’s deviceID contained in the SPC blob during FairPlay key exchange
//   - [AVContentKeyRequest.AVContentKeyRequestShouldRandomizeDeviceIdentifierKey]: Value is an Boolean indicating whether the user’s deviceID contained in the SPC blob during FairPlay key exchange should be randomized using a system generated seed
//
// # Responding to the content key request
//
//   - [AVContentKeyRequest.ProcessContentKeyResponse]: Sends the specified content key response to the receiver for processing.
//   - [AVContentKeyRequest.ProcessContentKeyResponseError]: Tells the receiver that the app was unable to obtain a content key response.
//
// # Getting content key request properties
//
//   - [AVContentKeyRequest.Identifier]: The identifier for the content key.
//   - [AVContentKeyRequest.OriginatingRecipient]: The AVContentKeyRecipient which initiated this request, if any.
//   - [AVContentKeyRequest.CanProvidePersistableContentKey]: The content key request used to create a persistable content key or respond to a previous request with a persistable content key.
//   - [AVContentKeyRequest.Error]: The error description for a failed key request.
//   - [AVContentKeyRequest.InitializationData]: The data used to obtain a key response.
//   - [AVContentKeyRequest.RenewsExpiringResponseData]: A Boolean value that indicates whether the content key request renews previously provided response data.
//   - [AVContentKeyRequest.Status]: The current state of the content key request.
//
// # Inspecting a request
//
//   - [AVContentKeyRequest.ContentKey]: The generated content key.
//   - [AVContentKeyRequest.ContentKeySpecifier]: The requested content key specifier.
//   - [AVContentKeyRequest.Options]: A dictionary of options used to initialize key loading.
//
// # Instance Methods
//
//   - [AVContentKeyRequest.RespondByRequestingPersistableContentKeyRequestAndReturnError]: Tells the receiver that the app requires a persistable content key request object for processing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest
type AVContentKeyRequest struct {
	objectivec.Object
}

// AVContentKeyRequestFromID constructs a [AVContentKeyRequest] from an objc.ID.
//
// An object that encapsulates information about a content decryption key
// request issued from a content key session object.
func AVContentKeyRequestFromID(id objc.ID) AVContentKeyRequest {
	return AVContentKeyRequest{objectivec.Object{ID: id}}
}
// NOTE: AVContentKeyRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVContentKeyRequest] class.
//
// # Getting content key request data
//
//   - [IAVContentKeyRequest.MakeStreamingContentKeyRequestDataForAppContentIdentifierOptionsCompletionHandler]: Obtains encrypted key request data for a specific combination of app and content.
//   - [IAVContentKeyRequest.AVContentKeyRequestProtocolVersionsKey]: A key that specifies the versions of the content protection protocol supported by the application.
//   - [IAVContentKeyRequest.AVContentKeyRequestRequiresValidationDataInSecureTokenKey]: A key that requires the secure token to have extended validation data.
//   - [IAVContentKeyRequest.AVContentKeyRequestRandomDeviceIdentifierSeedKey]: Value is an NSData containing a 16-byte seed to randomize the user’s deviceID contained in the SPC blob during FairPlay key exchange
//   - [IAVContentKeyRequest.AVContentKeyRequestShouldRandomizeDeviceIdentifierKey]: Value is an Boolean indicating whether the user’s deviceID contained in the SPC blob during FairPlay key exchange should be randomized using a system generated seed
//
// # Responding to the content key request
//
//   - [IAVContentKeyRequest.ProcessContentKeyResponse]: Sends the specified content key response to the receiver for processing.
//   - [IAVContentKeyRequest.ProcessContentKeyResponseError]: Tells the receiver that the app was unable to obtain a content key response.
//
// # Getting content key request properties
//
//   - [IAVContentKeyRequest.Identifier]: The identifier for the content key.
//   - [IAVContentKeyRequest.OriginatingRecipient]: The AVContentKeyRecipient which initiated this request, if any.
//   - [IAVContentKeyRequest.CanProvidePersistableContentKey]: The content key request used to create a persistable content key or respond to a previous request with a persistable content key.
//   - [IAVContentKeyRequest.Error]: The error description for a failed key request.
//   - [IAVContentKeyRequest.InitializationData]: The data used to obtain a key response.
//   - [IAVContentKeyRequest.RenewsExpiringResponseData]: A Boolean value that indicates whether the content key request renews previously provided response data.
//   - [IAVContentKeyRequest.Status]: The current state of the content key request.
//
// # Inspecting a request
//
//   - [IAVContentKeyRequest.ContentKey]: The generated content key.
//   - [IAVContentKeyRequest.ContentKeySpecifier]: The requested content key specifier.
//   - [IAVContentKeyRequest.Options]: A dictionary of options used to initialize key loading.
//
// # Instance Methods
//
//   - [IAVContentKeyRequest.RespondByRequestingPersistableContentKeyRequestAndReturnError]: Tells the receiver that the app requires a persistable content key request object for processing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest
type IAVContentKeyRequest interface {
	objectivec.IObject

	// Topic: Getting content key request data

	// Obtains encrypted key request data for a specific combination of app and content.
	MakeStreamingContentKeyRequestDataForAppContentIdentifierOptionsCompletionHandler(appIdentifier foundation.INSData, contentIdentifier foundation.INSData, options foundation.INSDictionary, handler DataErrorHandler)
	// A key that specifies the versions of the content protection protocol supported by the application.
	AVContentKeyRequestProtocolVersionsKey() string
	// A key that requires the secure token to have extended validation data.
	AVContentKeyRequestRequiresValidationDataInSecureTokenKey() string
	// Value is an NSData containing a 16-byte seed to randomize the user’s deviceID contained in the SPC blob during FairPlay key exchange
	AVContentKeyRequestRandomDeviceIdentifierSeedKey() string
	// Value is an Boolean indicating whether the user’s deviceID contained in the SPC blob during FairPlay key exchange should be randomized using a system generated seed
	AVContentKeyRequestShouldRandomizeDeviceIdentifierKey() string

	// Topic: Responding to the content key request

	// Sends the specified content key response to the receiver for processing.
	ProcessContentKeyResponse(keyResponse IAVContentKeyResponse)
	// Tells the receiver that the app was unable to obtain a content key response.
	ProcessContentKeyResponseError(error_ foundation.INSError)

	// Topic: Getting content key request properties

	// The identifier for the content key.
	Identifier() objectivec.IObject
	// The AVContentKeyRecipient which initiated this request, if any.
	OriginatingRecipient() AVContentKeyRecipient
	// The content key request used to create a persistable content key or respond to a previous request with a persistable content key.
	CanProvidePersistableContentKey() bool
	// The error description for a failed key request.
	Error() foundation.INSError
	// The data used to obtain a key response.
	InitializationData() foundation.INSData
	// A Boolean value that indicates whether the content key request renews previously provided response data.
	RenewsExpiringResponseData() bool
	// The current state of the content key request.
	Status() AVContentKeyRequestStatus

	// Topic: Inspecting a request

	// The generated content key.
	ContentKey() IAVContentKey
	// The requested content key specifier.
	ContentKeySpecifier() IAVContentKeySpecifier
	// A dictionary of options used to initialize key loading.
	Options() foundation.INSDictionary

	// Topic: Instance Methods

	// Tells the receiver that the app requires a persistable content key request object for processing.
	RespondByRequestingPersistableContentKeyRequestAndReturnError() (bool, error)
}

// Init initializes the instance.
func (c AVContentKeyRequest) Init() AVContentKeyRequest {
	rv := objc.Send[AVContentKeyRequest](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVContentKeyRequest) Autorelease() AVContentKeyRequest {
	rv := objc.Send[AVContentKeyRequest](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVContentKeyRequest creates a new AVContentKeyRequest instance.
func NewAVContentKeyRequest() AVContentKeyRequest {
	class := getAVContentKeyRequestClass()
	rv := objc.Send[AVContentKeyRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Obtains encrypted key request data for a specific combination of app and
// content.
//
// appIdentifier: An opaque identifier for the app.
//
// contentIdentifier: An opaque identifier for the content.
//
// options: A dictionary containing any additional information required to obtain the
// key. The value of this parameter is `nil` when no additional information is
// required.
//
// handler: A block called after the streaming content key request has been prepared.
// 
// contentKeyRequestData: The streaming content key request data. error: An
// object that describes the error, if one occurred; otherwise, the value is
// `nil`.
//
// # Discussion
// 
// If [AVContentKeyRequestProtocolVersionsKey] is not specified in the
// `options` parameter, the default protocol of `1` is used.
//
// [AVContentKeyRequestProtocolVersionsKey]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequestProtocolVersionsKey
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/makeStreamingContentKeyRequestData(forApp:contentIdentifier:options:completionHandler:)
func (c AVContentKeyRequest) MakeStreamingContentKeyRequestDataForAppContentIdentifierOptionsCompletionHandler(appIdentifier foundation.INSData, contentIdentifier foundation.INSData, options foundation.INSDictionary, handler DataErrorHandler) {
_block3, _ := NewDataErrorBlock(handler)
	objc.Send[objc.ID](c.ID, objc.Sel("makeStreamingContentKeyRequestDataForApp:contentIdentifier:options:completionHandler:"), appIdentifier, contentIdentifier, options, _block3)
}
// Sends the specified content key response to the receiver for processing.
//
// keyResponse: An [AVContentKeyResponse] object carrying a response to a content key
// request.
//
// # Discussion
// 
// After receiving a content key request and calling
// [StreamingContentKeyRequestDataForAppContentIdentifierOptionsCompletionHandler]
// on that request, you must obtain a response to the request in accordance
// with the protocol used by the entity that controls the use of the media
// data. Use this method to provide the content key response, to make
// protected content available for processing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/processContentKeyResponse(_:)
func (c AVContentKeyRequest) ProcessContentKeyResponse(keyResponse IAVContentKeyResponse) {
	objc.Send[objc.ID](c.ID, objc.Sel("processContentKeyResponse:"), keyResponse)
}
// Tells the receiver that the app was unable to obtain a content key
// response.
//
// error: An [NSError] that describes why the content key response failed.
// //
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/processContentKeyResponseError(_:)
func (c AVContentKeyRequest) ProcessContentKeyResponseError(error_ foundation.INSError) {
	objc.Send[objc.ID](c.ID, objc.Sel("processContentKeyResponseError:"), error_)
}
// Tells the receiver that the app requires a persistable content key request
// object for processing.
//
// # Discussion
// 
// To create a key that persists across multiple playback sessions, use this
// method to request an [AVPersistableContentKeyRequest] object. If the
// underlying protocol supports persistable content keys, the delegate
// receives a persistable content key request via the
// [ContentKeySessionDidProvidePersistableContentKeyRequest] method. An
// [internalInconsistencyException] is returned if your delegate does not
// respond to [ContentKeySessionDidProvidePersistableContentKeyRequest].
//
// [internalInconsistencyException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/internalInconsistencyException
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/respondByRequestingPersistableContentKeyRequest()-7i2pw
func (c AVContentKeyRequest) RespondByRequestingPersistableContentKeyRequestAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("respondByRequestingPersistableContentKeyRequestAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("respondByRequestingPersistableContentKeyRequestAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// A key that specifies the versions of the content protection protocol
// supported by the application.
//
// See: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequestprotocolversionskey
func (c AVContentKeyRequest) AVContentKeyRequestProtocolVersionsKey() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("AVContentKeyRequestProtocolVersionsKey"))
	return foundation.NSStringFromID(rv).String()
}
// A key that requires the secure token to have extended validation data.
//
// See: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequestrequiresvalidationdatainsecuretokenkey
func (c AVContentKeyRequest) AVContentKeyRequestRequiresValidationDataInSecureTokenKey() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("AVContentKeyRequestRequiresValidationDataInSecureTokenKey"))
	return foundation.NSStringFromID(rv).String()
}
// Value is an NSData containing a 16-byte seed to randomize the user’s
// deviceID contained in the SPC blob during FairPlay key exchange
//
// See: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequestrandomdeviceidentifierseedkey
func (c AVContentKeyRequest) AVContentKeyRequestRandomDeviceIdentifierSeedKey() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("AVContentKeyRequestRandomDeviceIdentifierSeedKey"))
	return foundation.NSStringFromID(rv).String()
}
// Value is an Boolean indicating whether the user’s deviceID contained in
// the SPC blob during FairPlay key exchange should be randomized using a
// system generated seed
//
// See: https://developer.apple.com/documentation/avfoundation/avcontentkeyrequestshouldrandomizedeviceidentifierkey
func (c AVContentKeyRequest) AVContentKeyRequestShouldRandomizeDeviceIdentifierKey() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("AVContentKeyRequestShouldRandomizeDeviceIdentifierKey"))
	return foundation.NSStringFromID(rv).String()
}
// The identifier for the content key.
//
// # Discussion
// 
// This property is specific to the container and the protocol. To use the key
// with an HTTP Live Streaming [AVURLAsset], the identifier must be an [NSURL]
// that matches a key [URI] in the media playlist.
//
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/identifier
func (c AVContentKeyRequest) Identifier() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("identifier"))
	return objectivec.Object{ID: rv}
}
// The AVContentKeyRecipient which initiated this request, if any.
//
// # Discussion
// 
// The originatingRecipient is an AVFoundation object responsible for
// initiating an AVContentKeyRequest. For example, an AVURLAsset used for
// playback can trigger an AVContentKeyRequest.
// 
// If an application triggers key loading directly, for example with
// -[AVContentKeySession
// processContentKeyRequestWithIdentifier:initializationData:options:], the
// value of originatingRecipient will be nil.
// 
// The originatingRecipient of key requests from HLS interstitials will always
// be the corresponding interstitial AVURLAsset. To receive key requests for
// DRM-protected interstitial content, applications must ensure their
// AVContentKeySession is attached to these interstitial AVURLAssets.
// 
// These interstitial AVURLAssets may be retrieved from the primary AVURLAsset
// via AVPlayerInterstitialEventMonitor.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/originatingRecipient
func (c AVContentKeyRequest) OriginatingRecipient() AVContentKeyRecipient {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("originatingRecipient"))
	return AVContentKeyRecipientObjectFromID(rv)
}
// The content key request used to create a persistable content key or respond
// to a previous request with a persistable content key.
//
// # Discussion
// 
// The value of this property is automatically set to [YES] when the receiver
// is provided to the content key session’s delegate via the
// [ContentKeySessionDidProvidePersistableContentKeyRequest] method. When this
// property is set to [YES], the
// [PersistableContentKeyFromKeyVendorResponseOptionsError] method can be used
// to create a persistable content key from the response.
// 
// When this property is set to [NO] and there is a request for a persistable
// content key, send the [RespondByRequestingPersistableContentKeyRequest]
// method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/canProvidePersistableContentKey
func (c AVContentKeyRequest) CanProvidePersistableContentKey() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("canProvidePersistableContentKey"))
	return rv
}
// The error description for a failed key request.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/error
func (c AVContentKeyRequest) Error() foundation.INSError {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
// The data used to obtain a key response.
//
// # Discussion
// 
// This property is specific to the container and the protocol.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/initializationData
func (c AVContentKeyRequest) InitializationData() foundation.INSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("initializationData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the content key request renews
// previously provided response data.
//
// # Discussion
// 
// The value of this property is [YES] if the request renews previously
// provided response data that is expiring or has already expired.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/renewsExpiringResponseData
func (c AVContentKeyRequest) RenewsExpiringResponseData() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("renewsExpiringResponseData"))
	return rv
}
// The current state of the content key request.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/status-swift.property
func (c AVContentKeyRequest) Status() AVContentKeyRequestStatus {
	rv := objc.Send[AVContentKeyRequestStatus](c.ID, objc.Sel("status"))
	return AVContentKeyRequestStatus(rv)
}
// The generated content key.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/contentKey
func (c AVContentKeyRequest) ContentKey() IAVContentKey {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("contentKey"))
	return AVContentKeyFromID(objc.ID(rv))
}
// The requested content key specifier.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/contentKeySpecifier
func (c AVContentKeyRequest) ContentKeySpecifier() IAVContentKeySpecifier {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("contentKeySpecifier"))
	return AVContentKeySpecifierFromID(objc.ID(rv))
}
// A dictionary of options used to initialize key loading.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/options
func (c AVContentKeyRequest) Options() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("options"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// MakeStreamingContentKeyRequestDataForAppContentIdentifierOptions is a synchronous wrapper around [AVContentKeyRequest.MakeStreamingContentKeyRequestDataForAppContentIdentifierOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c AVContentKeyRequest) MakeStreamingContentKeyRequestDataForAppContentIdentifierOptions(ctx context.Context, appIdentifier foundation.INSData, contentIdentifier foundation.INSData, options foundation.INSDictionary) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	c.MakeStreamingContentKeyRequestDataForAppContentIdentifierOptionsCompletionHandler(appIdentifier, contentIdentifier, options, func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}


// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVContentKeyResponse] class.
var (
	_AVContentKeyResponseClass     AVContentKeyResponseClass
	_AVContentKeyResponseClassOnce sync.Once
)

func getAVContentKeyResponseClass() AVContentKeyResponseClass {
	_AVContentKeyResponseClassOnce.Do(func() {
		_AVContentKeyResponseClass = AVContentKeyResponseClass{class: objc.GetClass("AVContentKeyResponse")}
	})
	return _AVContentKeyResponseClass
}

// GetAVContentKeyResponseClass returns the class object for AVContentKeyResponse.
func GetAVContentKeyResponseClass() AVContentKeyResponseClass {
	return getAVContentKeyResponseClass()
}

type AVContentKeyResponseClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVContentKeyResponseClass) Alloc() AVContentKeyResponse {
	rv := objc.Send[AVContentKeyResponse](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that encapsulates information about a response to a content
// decryption key request.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyResponse
type AVContentKeyResponse struct {
	objectivec.Object
}

// AVContentKeyResponseFromID constructs a [AVContentKeyResponse] from an objc.ID.
//
// An object that encapsulates information about a response to a content
// decryption key request.
func AVContentKeyResponseFromID(id objc.ID) AVContentKeyResponse {
	return AVContentKeyResponse{objectivec.Object{ID: id}}
}
// NOTE: AVContentKeyResponse adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVContentKeyResponse] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyResponse
type IAVContentKeyResponse interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c AVContentKeyResponse) Init() AVContentKeyResponse {
	rv := objc.Send[AVContentKeyResponse](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVContentKeyResponse) Autorelease() AVContentKeyResponse {
	rv := objc.Send[AVContentKeyResponse](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVContentKeyResponse creates a new AVContentKeyResponse instance.
func NewAVContentKeyResponse() AVContentKeyResponse {
	class := getAVContentKeyResponseClass()
	rv := objc.Send[AVContentKeyResponse](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a content key response with an authorization token.
//
// authorizationTokenData: A data value that contains the authorization token.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyResponse/init(authorizationTokenData:)
func NewContentKeyResponseWithAuthorizationTokenData(authorizationTokenData foundation.INSData) AVContentKeyResponse {
	rv := objc.Send[objc.ID](objc.ID(getAVContentKeyResponseClass().class), objc.Sel("contentKeyResponseWithAuthorizationTokenData:"), authorizationTokenData)
	return AVContentKeyResponseFromID(rv)
}

// Creates a new key response object for key data and initialization vector
// sent in the clear.
//
// keyData: The key used for decrypting content.
//
// initializationVector: The initialization vector used for decrypting content. This value is `nil`
// when the initialization vector is contained in the media to be decrypted.
//
// # Return Value
// 
// Returns a new [AVContentKeyResponse] object to decrypt content.
//
// # Discussion
// 
// Use the results of this initializer when the content key session creates a
// key request using the [clearKey] parameter. The results are then passed to
// the [ProcessContentKeyResponse] method to supply the decrypter with key
// data.
//
// [clearKey]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySystem/clearKey
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyResponse/init(clearKeyData:initializationVector:)
func NewContentKeyResponseWithClearKeyDataInitializationVector(keyData foundation.INSData, initializationVector foundation.INSData) AVContentKeyResponse {
	rv := objc.Send[objc.ID](objc.ID(getAVContentKeyResponseClass().class), objc.Sel("contentKeyResponseWithClearKeyData:initializationVector:"), keyData, initializationVector)
	return AVContentKeyResponseFromID(rv)
}

// Creates a content key response with an encrypted key response data blob
// when FairPlay Streaming is the key delivery method.
//
// keyResponseData: The key data from the FairPlay Streaming key server.
//
// # Return Value
// 
// Returns a new [AVContentKeyResponse] object to decrypt content.
//
// # Discussion
// 
// Use the results of this initializer when the content key session creates a
// key request using the [fairPlayStreaming] parameter. The results are then
// passed to the [ProcessContentKeyResponse] method to supply the decrypter
// with key data.
//
// [fairPlayStreaming]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySystem/fairPlayStreaming
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyResponse/init(fairPlayStreamingKeyResponseData:)
func NewContentKeyResponseWithFairPlayStreamingKeyResponseData(keyResponseData foundation.INSData) AVContentKeyResponse {
	rv := objc.Send[objc.ID](objc.ID(getAVContentKeyResponseClass().class), objc.Sel("contentKeyResponseWithFairPlayStreamingKeyResponseData:"), keyResponseData)
	return AVContentKeyResponseFromID(rv)
}


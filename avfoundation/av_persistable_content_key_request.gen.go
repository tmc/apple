// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVPersistableContentKeyRequest] class.
var (
	_AVPersistableContentKeyRequestClass     AVPersistableContentKeyRequestClass
	_AVPersistableContentKeyRequestClassOnce sync.Once
)

func getAVPersistableContentKeyRequestClass() AVPersistableContentKeyRequestClass {
	_AVPersistableContentKeyRequestClassOnce.Do(func() {
		_AVPersistableContentKeyRequestClass = AVPersistableContentKeyRequestClass{class: objc.GetClass("AVPersistableContentKeyRequest")}
	})
	return _AVPersistableContentKeyRequestClass
}

// GetAVPersistableContentKeyRequestClass returns the class object for AVPersistableContentKeyRequest.
func GetAVPersistableContentKeyRequestClass() AVPersistableContentKeyRequestClass {
	return getAVPersistableContentKeyRequestClass()
}

type AVPersistableContentKeyRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPersistableContentKeyRequestClass) Alloc() AVPersistableContentKeyRequest {
	rv := objc.Send[AVPersistableContentKeyRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that encapsulates information about a persistable content
// decryption key request issued from a content key session.
//
// # Overview
// 
// This class allows clients to create and use persistable content keys.
//
// # Requesting persistable content key data
//
//   - [AVPersistableContentKeyRequest.PersistableContentKeyFromKeyVendorResponseOptionsError]: Creates a persistable content key from the content key context data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPersistableContentKeyRequest
type AVPersistableContentKeyRequest struct {
	AVContentKeyRequest
}

// AVPersistableContentKeyRequestFromID constructs a [AVPersistableContentKeyRequest] from an objc.ID.
//
// An object that encapsulates information about a persistable content
// decryption key request issued from a content key session.
func AVPersistableContentKeyRequestFromID(id objc.ID) AVPersistableContentKeyRequest {
	return AVPersistableContentKeyRequest{AVContentKeyRequest: AVContentKeyRequestFromID(id)}
}
// NOTE: AVPersistableContentKeyRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPersistableContentKeyRequest] class.
//
// # Requesting persistable content key data
//
//   - [IAVPersistableContentKeyRequest.PersistableContentKeyFromKeyVendorResponseOptionsError]: Creates a persistable content key from the content key context data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPersistableContentKeyRequest
type IAVPersistableContentKeyRequest interface {
	IAVContentKeyRequest

	// Topic: Requesting persistable content key data

	// Creates a persistable content key from the content key context data.
	PersistableContentKeyFromKeyVendorResponseOptionsError(keyVendorResponse foundation.INSData, options foundation.INSDictionary) (foundation.INSData, error)
}

// Init initializes the instance.
func (p AVPersistableContentKeyRequest) Init() AVPersistableContentKeyRequest {
	rv := objc.Send[AVPersistableContentKeyRequest](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPersistableContentKeyRequest) Autorelease() AVPersistableContentKeyRequest {
	rv := objc.Send[AVPersistableContentKeyRequest](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPersistableContentKeyRequest creates a new AVPersistableContentKeyRequest instance.
func NewAVPersistableContentKeyRequest() AVPersistableContentKeyRequest {
	class := getAVPersistableContentKeyRequestClass()
	rv := objc.Send[AVPersistableContentKeyRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a persistable content key from the content key context data.
//
// keyVendorResponse: The response returned from the key vendor.
//
// options: Additional information required to obtain the persistable content key. The
// value of this parameter is `nil` when no additional information is
// required.
//
// # Return Value
// 
// Returns a data object that contains the persistable content key.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPersistableContentKeyRequest/persistableContentKey(fromKeyVendorResponse:options:)
func (p AVPersistableContentKeyRequest) PersistableContentKeyFromKeyVendorResponseOptionsError(keyVendorResponse foundation.INSData, options foundation.INSDictionary) (foundation.INSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](p.ID, objc.Sel("persistableContentKeyFromKeyVendorResponse:options:error:"), keyVendorResponse, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSData{}, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDataFromID(rv), nil

}


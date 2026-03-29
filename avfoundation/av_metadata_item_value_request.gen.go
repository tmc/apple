// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMetadataItemValueRequest] class.
var (
	_AVMetadataItemValueRequestClass     AVMetadataItemValueRequestClass
	_AVMetadataItemValueRequestClassOnce sync.Once
)

func getAVMetadataItemValueRequestClass() AVMetadataItemValueRequestClass {
	_AVMetadataItemValueRequestClassOnce.Do(func() {
		_AVMetadataItemValueRequestClass = AVMetadataItemValueRequestClass{class: objc.GetClass("AVMetadataItemValueRequest")}
	})
	return _AVMetadataItemValueRequestClass
}

// GetAVMetadataItemValueRequestClass returns the class object for AVMetadataItemValueRequest.
func GetAVMetadataItemValueRequestClass() AVMetadataItemValueRequestClass {
	return getAVMetadataItemValueRequestClass()
}

type AVMetadataItemValueRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetadataItemValueRequestClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetadataItemValueRequestClass) Alloc() AVMetadataItemValueRequest {
	rv := objc.Send[AVMetadataItemValueRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that responds to a request to load the value of a metadata item.
//
// # Handling the response
//
//   - [AVMetadataItemValueRequest.RespondWithValue]: Returns the metadata item’s value.
//   - [AVMetadataItemValueRequest.RespondWithError]: Returns an error when the system fails to load the value.
//   - [AVMetadataItemValueRequest.MetadataItem]: The metadata item to request a value for.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItemValueRequest
type AVMetadataItemValueRequest struct {
	objectivec.Object
}

// AVMetadataItemValueRequestFromID constructs a [AVMetadataItemValueRequest] from an objc.ID.
//
// An object that responds to a request to load the value of a metadata item.
func AVMetadataItemValueRequestFromID(id objc.ID) AVMetadataItemValueRequest {
	return AVMetadataItemValueRequest{objectivec.Object{ID: id}}
}
// NOTE: AVMetadataItemValueRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetadataItemValueRequest] class.
//
// # Handling the response
//
//   - [IAVMetadataItemValueRequest.RespondWithValue]: Returns the metadata item’s value.
//   - [IAVMetadataItemValueRequest.RespondWithError]: Returns an error when the system fails to load the value.
//   - [IAVMetadataItemValueRequest.MetadataItem]: The metadata item to request a value for.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItemValueRequest
type IAVMetadataItemValueRequest interface {
	objectivec.IObject

	// Topic: Handling the response

	// Returns the metadata item’s value.
	RespondWithValue(value objectivec.IObject)
	// Returns an error when the system fails to load the value.
	RespondWithError(error_ foundation.INSError)
	// The metadata item to request a value for.
	MetadataItem() IAVMetadataItem
}

// Init initializes the instance.
func (m AVMetadataItemValueRequest) Init() AVMetadataItemValueRequest {
	rv := objc.Send[AVMetadataItemValueRequest](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetadataItemValueRequest) Autorelease() AVMetadataItemValueRequest {
	rv := objc.Send[AVMetadataItemValueRequest](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetadataItemValueRequest creates a new AVMetadataItemValueRequest instance.
func NewAVMetadataItemValueRequest() AVMetadataItemValueRequest {
	class := getAVMetadataItemValueRequestClass()
	rv := objc.Send[AVMetadataItemValueRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the metadata item’s value.
//
// value: The value to return for the request.
//
// # Discussion
// 
// You call this method to return the metadata item’s value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItemValueRequest/respond(value:)
func (m AVMetadataItemValueRequest) RespondWithValue(value objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("respondWithValue:"), value)
}
// Returns an error when the system fails to load the value.
//
// error: The error to return for the request.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItemValueRequest/respond(error:)
func (m AVMetadataItemValueRequest) RespondWithError(error_ foundation.INSError) {
	objc.Send[objc.ID](m.ID, objc.Sel("respondWithError:"), error_)
}

// The metadata item to request a value for.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItemValueRequest/metadataItem
func (m AVMetadataItemValueRequest) MetadataItem() IAVMetadataItem {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("metadataItem"))
	return AVMetadataItemFromID(objc.ID(rv))
}


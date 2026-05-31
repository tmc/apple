// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SFSpeechURLRecognitionRequest] class.
var (
	_SFSpeechURLRecognitionRequestClass     SFSpeechURLRecognitionRequestClass
	_SFSpeechURLRecognitionRequestClassOnce sync.Once
)

func getSFSpeechURLRecognitionRequestClass() SFSpeechURLRecognitionRequestClass {
	_SFSpeechURLRecognitionRequestClassOnce.Do(func() {
		_SFSpeechURLRecognitionRequestClass = SFSpeechURLRecognitionRequestClass{class: objc.GetClass("SFSpeechURLRecognitionRequest")}
	})
	return _SFSpeechURLRecognitionRequestClass
}

// GetSFSpeechURLRecognitionRequestClass returns the class object for SFSpeechURLRecognitionRequest.
func GetSFSpeechURLRecognitionRequestClass() SFSpeechURLRecognitionRequestClass {
	return getSFSpeechURLRecognitionRequestClass()
}

type SFSpeechURLRecognitionRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SFSpeechURLRecognitionRequestClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SFSpeechURLRecognitionRequestClass) Alloc() SFSpeechURLRecognitionRequest {
	rv := objc.Send[SFSpeechURLRecognitionRequest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A request to recognize speech in a recorded audio file.
//
// # Overview
//
// Use this object to perform speech recognition on the contents of an audio
// file.
//
// The following example shows a method that performs recognition on an audio
// file based on the user’s default language and prints out the
// transcription.
//
// Listing 1. Getting a speech recognizer and making a recognition request
//
// # Creating a speech recognition request
//
//   - [SFSpeechURLRecognitionRequest.InitWithURL]: Creates a speech recognition request, initialized with the specified URL.
//
// # Accessing the audio file URL
//
//   - [SFSpeechURLRecognitionRequest.URL]: The URL of the audio file.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechURLRecognitionRequest
type SFSpeechURLRecognitionRequest struct {
	SFSpeechRecognitionRequest
}

// SFSpeechURLRecognitionRequestFromID constructs a [SFSpeechURLRecognitionRequest] from an objc.ID.
//
// A request to recognize speech in a recorded audio file.
func SFSpeechURLRecognitionRequestFromID(id objc.ID) SFSpeechURLRecognitionRequest {
	return SFSpeechURLRecognitionRequest{SFSpeechRecognitionRequest: SFSpeechRecognitionRequestFromID(id)}
}

// NOTE: SFSpeechURLRecognitionRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SFSpeechURLRecognitionRequest] class.
//
// # Creating a speech recognition request
//
//   - [ISFSpeechURLRecognitionRequest.InitWithURL]: Creates a speech recognition request, initialized with the specified URL.
//
// # Accessing the audio file URL
//
//   - [ISFSpeechURLRecognitionRequest.URL]: The URL of the audio file.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechURLRecognitionRequest
type ISFSpeechURLRecognitionRequest interface {
	ISFSpeechRecognitionRequest

	// Topic: Creating a speech recognition request

	// Creates a speech recognition request, initialized with the specified URL.
	InitWithURL(URL foundation.NSURL) SFSpeechURLRecognitionRequest

	// Topic: Accessing the audio file URL

	// The URL of the audio file.
	URL() foundation.NSURL
}

// Init initializes the instance.
func (s SFSpeechURLRecognitionRequest) Init() SFSpeechURLRecognitionRequest {
	rv := objc.Send[SFSpeechURLRecognitionRequest](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SFSpeechURLRecognitionRequest) Autorelease() SFSpeechURLRecognitionRequest {
	rv := objc.Send[SFSpeechURLRecognitionRequest](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSpeechURLRecognitionRequest creates a new SFSpeechURLRecognitionRequest instance.
func NewSFSpeechURLRecognitionRequest() SFSpeechURLRecognitionRequest {
	class := getSFSpeechURLRecognitionRequestClass()
	rv := objc.Send[SFSpeechURLRecognitionRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a speech recognition request, initialized with the specified URL.
//
// # Discussion
//
// Use this method to create a request to recognize speech in a recorded audio
// file that resides at the specified URL. Pass the request to the
// recognizer’s [SFSpeechRecognizer.RecognitionTaskWithRequestDelegate]
// method to start recognition.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechURLRecognitionRequest/init(url:)
func NewSpeechURLRecognitionRequestWithURL(URL foundation.NSURL) SFSpeechURLRecognitionRequest {
	instance := getSFSpeechURLRecognitionRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), URL)
	return SFSpeechURLRecognitionRequestFromID(rv)
}

// Creates a speech recognition request, initialized with the specified URL.
//
// # Discussion
//
// Use this method to create a request to recognize speech in a recorded audio
// file that resides at the specified URL. Pass the request to the
// recognizer’s [SFSpeechRecognizer.RecognitionTaskWithRequestDelegate]
// method to start recognition.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechURLRecognitionRequest/init(url:)
func (s SFSpeechURLRecognitionRequest) InitWithURL(URL foundation.NSURL) SFSpeechURLRecognitionRequest {
	rv := objc.Send[SFSpeechURLRecognitionRequest](s.ID, objc.Sel("initWithURL:"), URL)
	return rv
}

// The URL of the audio file.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechURLRecognitionRequest/url
func (s SFSpeechURLRecognitionRequest) URL() foundation.NSURL {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

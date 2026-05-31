// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SFSpeechRecognizer] class.
var (
	_SFSpeechRecognizerClass     SFSpeechRecognizerClass
	_SFSpeechRecognizerClassOnce sync.Once
)

func getSFSpeechRecognizerClass() SFSpeechRecognizerClass {
	_SFSpeechRecognizerClassOnce.Do(func() {
		_SFSpeechRecognizerClass = SFSpeechRecognizerClass{class: objc.GetClass("SFSpeechRecognizer")}
	})
	return _SFSpeechRecognizerClass
}

// GetSFSpeechRecognizerClass returns the class object for SFSpeechRecognizer.
func GetSFSpeechRecognizerClass() SFSpeechRecognizerClass {
	return getSFSpeechRecognizerClass()
}

type SFSpeechRecognizerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SFSpeechRecognizerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SFSpeechRecognizerClass) Alloc() SFSpeechRecognizer {
	rv := objc.Send[SFSpeechRecognizer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An object you use to check for the availability of the speech recognition
// service, and to initiate the speech recognition process.
//
// # Overview
//
// An [SFSpeechRecognizer] object is the central object for managing the
// speech recognizer process. Use this object to:
//
// - Request authorization to use speech recognition services. - Specify the
// language to use during the recognition process. - Initiate new speech
// recognition tasks.
//
// # Set up speech recognition
//
// Each speech recognizer supports only one language, which you specify at
// creation time. The successful creation of a speech recognizer does not
// guarantee that speech recognition services are available. For some
// languages, the recognizer might require an Internet connection. Use the
// [SFSpeechRecognizer.Available] property to find out if speech recognition
// services are available for the current language.
//
// To initiate the speech recognition process, do the following:
//
// - Request authorization to use speech recognition. See [Asking Permission
// to Use Speech Recognition]. - Create an [SFSpeechRecognizer] object. -
// Verify the availability of services using the
// [SFSpeechRecognizer.Available] property of your speech recognizer object. -
// Prepare your audio content. - Create a recognition request object—an
// object that descends from [SFSpeechRecognitionRequest]. - Call the
// [SFSpeechRecognizer.RecognitionTaskWithRequestDelegate] or
// [SFSpeechRecognizer.RecognitionTaskWithRequestResultHandler] method to
// begin the recognition process.
//
// The type of recognition request object you create depends on whether you
// are processing an existing audio file or an incoming stream of audio. For
// existing audio files, create a [SFSpeechURLRecognitionRequest] object. For
// audio streams, create a [SFSpeechAudioBufferRecognitionRequest] object.
//
// # Create a great user experience for speech recognition
//
// Here are some tips to consider when adding speech recognition support to
// your app.
//
// - Because speech recognition is a network-based service, limits are
// enforced so that the service can remain freely available to all apps.
// Individual devices may be limited in the number of recognitions that can be
// performed per day, and each app may be throttled globally based on the
// number of requests it makes per day. If a recognition request fails quickly
// (within a second or two of starting), check to see if the recognition
// service became unavailable. If it is, you may want to ask users to try
// again later. - Speech recognition places a relatively high burden on
// battery life and network usage. To minimize this burden, the framework
// stops speech recognition tasks that last longer than one minute. This limit
// is similar to the one for keyboard-related dictation. - For example,
// display a visual indicator and play sounds at the beginning and end of
// speech recognition to help users understand that they’re being actively
// recorded. You can also display speech as it is being recognized so that
// users understand what your app is doing and see any mistakes made during
// the recognition process. - Some speech is not appropriate for recognition.
// Don’t send passwords, health or financial data, and other sensitive
// speech for recognition.
//
// # Creating a speech recognizer
//
//   - [SFSpeechRecognizer.InitWithLocale]: Creates a speech recognizer associated with the specified locale.
//
// # Monitoring speech recognition availability
//
//   - [SFSpeechRecognizer.Delegate]: The delegate object that handles changes to the availability of speech recognition services.
//   - [SFSpeechRecognizer.SetDelegate]
//   - [SFSpeechRecognizer.IsAvailable]: A Boolean value that indicates whether the speech recognizer is currently available.
//   - [SFSpeechRecognizer.SupportsOnDeviceRecognition]: A Boolean value that indicates whether the speech recognizer can operate without network access.
//   - [SFSpeechRecognizer.SetSupportsOnDeviceRecognition]
//
// # Configuring the speech recognizer
//
//   - [SFSpeechRecognizer.DefaultTaskHint]: A hint that indicates the type of speech recognition being requested.
//   - [SFSpeechRecognizer.SetDefaultTaskHint]
//   - [SFSpeechRecognizer.Queue]: The queue on which to execute recognition task handlers and delegate methods.
//   - [SFSpeechRecognizer.SetQueue]
//
// # Performing speech recognition on audio
//
//   - [SFSpeechRecognizer.RecognitionTaskWithRequestResultHandler]: Executes the speech recognition request and delivers the results to the specified handler block.
//   - [SFSpeechRecognizer.RecognitionTaskWithRequestDelegate]: Recognizes speech from the audio source associated with the specified request, using the specified delegate to manage the results.
//
// # Getting the current language
//
//   - [SFSpeechRecognizer.Locale]: The locale of the speech recognizer.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer
//
// [Asking Permission to Use Speech Recognition]: https://developer.apple.com/documentation/Speech/asking-permission-to-use-speech-recognition
type SFSpeechRecognizer struct {
	objectivec.Object
}

// SFSpeechRecognizerFromID constructs a [SFSpeechRecognizer] from an objc.ID.
//
// An object you use to check for the availability of the speech recognition
// service, and to initiate the speech recognition process.
func SFSpeechRecognizerFromID(id objc.ID) SFSpeechRecognizer {
	return SFSpeechRecognizer{objectivec.Object{ID: id}}
}

// NOTE: SFSpeechRecognizer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SFSpeechRecognizer] class.
//
// # Creating a speech recognizer
//
//   - [ISFSpeechRecognizer.InitWithLocale]: Creates a speech recognizer associated with the specified locale.
//
// # Monitoring speech recognition availability
//
//   - [ISFSpeechRecognizer.Delegate]: The delegate object that handles changes to the availability of speech recognition services.
//   - [ISFSpeechRecognizer.SetDelegate]
//   - [ISFSpeechRecognizer.IsAvailable]: A Boolean value that indicates whether the speech recognizer is currently available.
//   - [ISFSpeechRecognizer.SupportsOnDeviceRecognition]: A Boolean value that indicates whether the speech recognizer can operate without network access.
//   - [ISFSpeechRecognizer.SetSupportsOnDeviceRecognition]
//
// # Configuring the speech recognizer
//
//   - [ISFSpeechRecognizer.DefaultTaskHint]: A hint that indicates the type of speech recognition being requested.
//   - [ISFSpeechRecognizer.SetDefaultTaskHint]
//   - [ISFSpeechRecognizer.Queue]: The queue on which to execute recognition task handlers and delegate methods.
//   - [ISFSpeechRecognizer.SetQueue]
//
// # Performing speech recognition on audio
//
//   - [ISFSpeechRecognizer.RecognitionTaskWithRequestResultHandler]: Executes the speech recognition request and delivers the results to the specified handler block.
//   - [ISFSpeechRecognizer.RecognitionTaskWithRequestDelegate]: Recognizes speech from the audio source associated with the specified request, using the specified delegate to manage the results.
//
// # Getting the current language
//
//   - [ISFSpeechRecognizer.Locale]: The locale of the speech recognizer.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer
type ISFSpeechRecognizer interface {
	objectivec.IObject

	// Topic: Creating a speech recognizer

	// Creates a speech recognizer associated with the specified locale.
	InitWithLocale(locale foundation.NSLocale) SFSpeechRecognizer

	// Topic: Monitoring speech recognition availability

	// The delegate object that handles changes to the availability of speech recognition services.
	Delegate() SFSpeechRecognizerDelegate
	SetDelegate(value SFSpeechRecognizerDelegate)
	// A Boolean value that indicates whether the speech recognizer is currently available.
	IsAvailable() bool
	// A Boolean value that indicates whether the speech recognizer can operate without network access.
	SupportsOnDeviceRecognition() bool
	SetSupportsOnDeviceRecognition(value bool)

	// Topic: Configuring the speech recognizer

	// A hint that indicates the type of speech recognition being requested.
	DefaultTaskHint() SFSpeechRecognitionTaskHint
	SetDefaultTaskHint(value SFSpeechRecognitionTaskHint)
	// The queue on which to execute recognition task handlers and delegate methods.
	Queue() foundation.OperationQueue
	SetQueue(value foundation.OperationQueue)

	// Topic: Performing speech recognition on audio

	// Executes the speech recognition request and delivers the results to the specified handler block.
	RecognitionTaskWithRequestResultHandler(request ISFSpeechRecognitionRequest, resultHandler SFSpeechRecognitionResultErrorHandler) ISFSpeechRecognitionTask
	// Recognizes speech from the audio source associated with the specified request, using the specified delegate to manage the results.
	RecognitionTaskWithRequestDelegate(request ISFSpeechRecognitionRequest, delegate SFSpeechRecognitionTaskDelegate) ISFSpeechRecognitionTask

	// Topic: Getting the current language

	// The locale of the speech recognizer.
	Locale() foundation.NSLocale
}

// Init initializes the instance.
func (s SFSpeechRecognizer) Init() SFSpeechRecognizer {
	rv := objc.Send[SFSpeechRecognizer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SFSpeechRecognizer) Autorelease() SFSpeechRecognizer {
	rv := objc.Send[SFSpeechRecognizer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSpeechRecognizer creates a new SFSpeechRecognizer instance.
func NewSFSpeechRecognizer() SFSpeechRecognizer {
	class := getSFSpeechRecognizerClass()
	rv := objc.Send[SFSpeechRecognizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a speech recognizer associated with the specified locale.
//
// locale: The locale object representing the language you want to use for speech
// recognition. For a list of languages supported by the speech recognizer,
// see [SFSpeechRecognizerClass.SupportedLocales].
//
// # Return Value
//
// An initialized speech recognizer object, or `nil` if the specified language
// was not supported.
//
// # Discussion
//
// If you specify a language that is not supported by the speech recognizer,
// this method attempts to fall back to the language used by the keyboard for
// dictation. If that fails, this method returns `nil`.
//
// Even if this method returns a valid speech recognizer object, the speech
// recognition services may be temporarily unavailable. To determine whether
// speech recognition services are available, check the
// [SFSpeechRecognizer.Available] property.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/init(locale:)
func NewSpeechRecognizerWithLocale(locale foundation.NSLocale) SFSpeechRecognizer {
	instance := getSFSpeechRecognizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLocale:"), locale)
	return SFSpeechRecognizerFromID(rv)
}

// Creates a speech recognizer associated with the specified locale.
//
// locale: The locale object representing the language you want to use for speech
// recognition. For a list of languages supported by the speech recognizer,
// see [SFSpeechRecognizerClass.SupportedLocales].
//
// # Return Value
//
// An initialized speech recognizer object, or `nil` if the specified language
// was not supported.
//
// # Discussion
//
// If you specify a language that is not supported by the speech recognizer,
// this method attempts to fall back to the language used by the keyboard for
// dictation. If that fails, this method returns `nil`.
//
// Even if this method returns a valid speech recognizer object, the speech
// recognition services may be temporarily unavailable. To determine whether
// speech recognition services are available, check the
// [SFSpeechRecognizer.Available] property.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/init(locale:)
func (s SFSpeechRecognizer) InitWithLocale(locale foundation.NSLocale) SFSpeechRecognizer {
	rv := objc.Send[SFSpeechRecognizer](s.ID, objc.Sel("initWithLocale:"), locale)
	return rv
}

// Executes the speech recognition request and delivers the results to the
// specified handler block.
//
// request: A request (in an [SFSpeechRecognitionRequest] object) to recognize speech
// from an audio source.
//
// resultHandler: The block to call when partial or final results are available, or when an
// error occurs. If the
// [SFSpeechRecognitionRequest.ShouldReportPartialResults] property is `true`,
// this block may be called multiple times to deliver the partial and final
// results. The block has no return value and takes the following parameters:
//
// result: A [SFSpeechRecognitionResult] containing the partial or final
// transcriptions of the audio content. error: An error object if a problem
// occurred. This parameter is `nil` if speech recognition was successful.
//
// # Return Value
//
// The task object you can use to manage an in-progress recognition request.
//
// # Discussion
//
// Use this method to initiate the speech recognition process on the audio
// contained in the request object. This method executes asynchronously and
// returns a [SFSpeechRecognitionTask] object that you can use to cancel or
// finalize the recognition process later. As results become available, the
// method calls the block in the `resultHandler` parameter.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/recognitionTask(with:resultHandler:)
func (s SFSpeechRecognizer) RecognitionTaskWithRequestResultHandler(request ISFSpeechRecognitionRequest, resultHandler SFSpeechRecognitionResultErrorHandler) ISFSpeechRecognitionTask {
	_block1, _ := NewSFSpeechRecognitionResultErrorBlock(resultHandler)
	rv := objc.Send[objc.ID](s.ID, objc.Sel("recognitionTaskWithRequest:resultHandler:"), request, _block1)
	return SFSpeechRecognitionTaskFromID(rv)
}

// Recognizes speech from the audio source associated with the specified
// request, using the specified delegate to manage the results.
//
// request: A request (encapsulated in an [SFSpeechRecognitionRequest] object) to
// recognize speech from an audio source.
//
// delegate: An object that can handle results from the speech recognition task. This
// object must conform to the [SFSpeechRecognitionTaskDelegate] protocol.
//
// # Return Value
//
// The task object you can use to manage an in-progress recognition request.
//
// # Discussion
//
// Use this method to initiate the speech recognition process on the audio
// contained in the request object. This method executes asynchronously and
// returns a [SFSpeechRecognitionTask] object that you can use to cancel or
// finalize the recognition process later. As results become available, the
// method calls the methods of the provided `delegate` object.
//
// Note that the [SFSpeechRecognitionTask] object returned by this method does
// not retain your delegate object. You must maintain a strong reference to
// your delegate while speech recognition is in progress.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/recognitionTask(with:delegate:)
func (s SFSpeechRecognizer) RecognitionTaskWithRequestDelegate(request ISFSpeechRecognitionRequest, delegate SFSpeechRecognitionTaskDelegate) ISFSpeechRecognitionTask {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("recognitionTaskWithRequest:delegate:"), request, delegate)
	return SFSpeechRecognitionTaskFromID(rv)
}

// Asks the user to allow your app to perform speech recognition.
//
// handler: The block to execute when your app’s authorization status is known. The
// status parameter of the block contains your app’s authorization status.
// The system does not guarantee the execution of this block on your app’s
// main dispatch queue.
//
// # Discussion
//
// Call this method before performing any other tasks associated with speech
// recognition. This method executes asynchronously, returning shortly after
// you call it. At some point later, the system calls the provided `handler`
// block with the results.
//
// When your app’s authorization status is
// [SFSpeechRecognizerAuthorizationStatus.notDetermined], this method causes
// the system to prompt the user to grant or deny permission for your app to
// use speech recognition. The prompt includes the custom message you specify
// in the [NSSpeechRecognitionUsageDescription] key of your app’s
// `Info.Plist()` file. The user’s response is saved so that future calls to
// this method do not prompt the user again.
//
// For more information about requesting authorization, see [Asking Permission
// to Use Speech Recognition].
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/requestAuthorization(_:)
//
// [Asking Permission to Use Speech Recognition]: https://developer.apple.com/documentation/Speech/asking-permission-to-use-speech-recognition
// [SFSpeechRecognizerAuthorizationStatus.notDetermined]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizerAuthorizationStatus/notDetermined
func (_SFSpeechRecognizerClass SFSpeechRecognizerClass) RequestAuthorization(handler SFSpeechRecognizerAuthorizationStatusHandler) {
	_block0, _ := NewSFSpeechRecognizerAuthorizationStatusBlock(handler)
	objc.Send[objc.ID](objc.ID(_SFSpeechRecognizerClass.class), objc.Sel("requestAuthorization:"), _block0)
}

// Returns your app’s current authorization to perform speech recognition.
//
// # Return Value
//
// The app’s current authorization status value. For a list of values, see
// [SFSpeechRecognizerAuthorizationStatus].
//
// # Discussion
//
// The user can reject your app’s request to perform speech recognition, but
// your request can also be denied if speech recognition is not supported on
// the device. The app can also change your app’s authorization status at
// any time from the Settings app.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/authorizationStatus()
//
// [SFSpeechRecognizerAuthorizationStatus]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizerAuthorizationStatus
func (_SFSpeechRecognizerClass SFSpeechRecognizerClass) AuthorizationStatus() SFSpeechRecognizerAuthorizationStatus {
	rv := objc.Send[SFSpeechRecognizerAuthorizationStatus](objc.ID(_SFSpeechRecognizerClass.class), objc.Sel("authorizationStatus"))
	return SFSpeechRecognizerAuthorizationStatus(rv)
}

// Returns the set of locales that are supported by the speech recognizer.
//
// # Return Value
//
// A set of locales that support speech recognition.
//
// # Discussion
//
// This method returns the locales for which speech recognition is supported.
// Support for a locale does not guarantee that speech recognition is
// currently possible for that locale. For some locales, the speech recognizer
// requires an active Internet connection to communicate with Apple’s
// servers. If the speech recognizer is currently unable to process requests,
// [SFSpeechRecognizer.Available] returns `false`.
//
// Speech recognition supports the same locales that are supported by the
// keyboard’s dictation feature. For a list of these locales, see [QuickType
// Keyboard: Dictation].
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/supportedLocales()
//
// [QuickType Keyboard: Dictation]: https://www.apple.com/ios/feature-availability/#quicktype-keyboard-dictation
func (_SFSpeechRecognizerClass SFSpeechRecognizerClass) SupportedLocales() foundation.INSSet {
	rv := objc.Send[objc.ID](objc.ID(_SFSpeechRecognizerClass.class), objc.Sel("supportedLocales"))
	return foundation.NSSetFromID(rv)
}

// The delegate object that handles changes to the availability of speech
// recognition services.
//
// # Discussion
//
// Provide a delegate object when you want to monitor changes to the
// availability of speech recognition services. Your delegate object must
// conform to the [SFSpeechRecognizerDelegate] protocol.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/delegate
func (s SFSpeechRecognizer) Delegate() SFSpeechRecognizerDelegate {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return SFSpeechRecognizerDelegateObjectFromID(rv)
}
func (s SFSpeechRecognizer) SetDelegate(value SFSpeechRecognizerDelegate) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the speech recognizer is currently
// available.
//
// # Discussion
//
// When the value of this property is `true`, you may create new speech
// recognition tasks. When value of this property is `false`, speech
// recognition services are not available.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/isAvailable
func (s SFSpeechRecognizer) IsAvailable() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isAvailable"))
	return rv
}

// A Boolean value that indicates whether the speech recognizer can operate
// without network access.
//
// # Discussion
//
// An [SFSpeechRecognitionRequest] can only honor its
// [SFSpeechRecognitionRequest.RequiresOnDeviceRecognition] property if
// [SFSpeechRecognizer.SupportsOnDeviceRecognition] is `true`. If
// [SFSpeechRecognizer.SupportsOnDeviceRecognition] is `false`, the
// [SFSpeechRecognizer] requires a network in order to recognize speech.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/supportsOnDeviceRecognition
func (s SFSpeechRecognizer) SupportsOnDeviceRecognition() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("supportsOnDeviceRecognition"))
	return rv
}
func (s SFSpeechRecognizer) SetSupportsOnDeviceRecognition(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setSupportsOnDeviceRecognition:"), value)
}

// A hint that indicates the type of speech recognition being requested.
//
// # Discussion
//
// By default, the value of this property overrides the
// [SFSpeechRecognitionTaskHint.unspecified] value for requests. For possible
// values, see [SFSpeechRecognitionTaskHint].
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/defaultTaskHint
//
// [SFSpeechRecognitionTaskHint.unspecified]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskHint/unspecified
// [SFSpeechRecognitionTaskHint]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskHint
func (s SFSpeechRecognizer) DefaultTaskHint() SFSpeechRecognitionTaskHint {
	rv := objc.Send[SFSpeechRecognitionTaskHint](s.ID, objc.Sel("defaultTaskHint"))
	return SFSpeechRecognitionTaskHint(rv)
}
func (s SFSpeechRecognizer) SetDefaultTaskHint(value SFSpeechRecognitionTaskHint) {
	objc.Send[struct{}](s.ID, objc.Sel("setDefaultTaskHint:"), value)
}

// The queue on which to execute recognition task handlers and delegate
// methods.
//
// # Discussion
//
// The default value of this property is the app’s main queue. Assign a
// different queue if you want delegate methods and handlers to be executed on
// a background queue.
//
// The handler you pass to the [SFSpeechRecognizerClass.RequestAuthorization]
// method does not use this queue.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/queue
func (s SFSpeechRecognizer) Queue() foundation.OperationQueue {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("queue"))
	return foundation.OperationQueueFromID(objc.ID(rv))
}
func (s SFSpeechRecognizer) SetQueue(value foundation.OperationQueue) {
	objc.Send[struct{}](s.ID, objc.Sel("setQueue:"), value)
}

// The locale of the speech recognizer.
//
// # Discussion
//
// The locale of the speech recognizer is an [NSLocale] object. The default
// value of this property is the system locale (that is, `+[NSLocale
// systemLocale]`).
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognizer/locale
func (s SFSpeechRecognizer) Locale() foundation.NSLocale {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("locale"))
	return foundation.NSLocaleFromID(objc.ID(rv))
}

// RequestAuthorizationSync is a synchronous wrapper around [SFSpeechRecognizer.RequestAuthorization].
// It blocks until the completion handler fires or the context is cancelled.
func (sc SFSpeechRecognizerClass) RequestAuthorizationSync(ctx context.Context) (SFSpeechRecognizerAuthorizationStatus, error) {
	done := make(chan SFSpeechRecognizerAuthorizationStatus, 1)
	sc.RequestAuthorization(func(val SFSpeechRecognizerAuthorizationStatus) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return *new(SFSpeechRecognizerAuthorizationStatus), ctx.Err()
	}
}

// RecognitionTaskWithRequestResultHandlerSync is a synchronous wrapper around [SFSpeechRecognizer.RecognitionTaskWithRequestResultHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s SFSpeechRecognizer) RecognitionTaskWithRequestResultHandlerSync(ctx context.Context, request ISFSpeechRecognitionRequest) (*SFSpeechRecognitionResult, error) {
	type result struct {
		val *SFSpeechRecognitionResult
		err error
	}
	done := make(chan result, 1)
	s.RecognitionTaskWithRequestResultHandler(request, func(val *SFSpeechRecognitionResult, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

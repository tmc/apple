// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SFSpeechRecognitionRequest] class.
var (
	_SFSpeechRecognitionRequestClass     SFSpeechRecognitionRequestClass
	_SFSpeechRecognitionRequestClassOnce sync.Once
)

func getSFSpeechRecognitionRequestClass() SFSpeechRecognitionRequestClass {
	_SFSpeechRecognitionRequestClassOnce.Do(func() {
		_SFSpeechRecognitionRequestClass = SFSpeechRecognitionRequestClass{class: objc.GetClass("SFSpeechRecognitionRequest")}
	})
	return _SFSpeechRecognitionRequestClass
}

// GetSFSpeechRecognitionRequestClass returns the class object for SFSpeechRecognitionRequest.
func GetSFSpeechRecognitionRequestClass() SFSpeechRecognitionRequestClass {
	return getSFSpeechRecognitionRequestClass()
}

type SFSpeechRecognitionRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SFSpeechRecognitionRequestClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SFSpeechRecognitionRequestClass) Alloc() SFSpeechRecognitionRequest {
	rv := objc.Send[SFSpeechRecognitionRequest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that represents a request to recognize speech from an
// audio source.
//
// # Overview
//
// Don’t create [SFSpeechRecognitionRequest] objects directly. Create an
// [SFSpeechURLRecognitionRequest] or [SFSpeechAudioBufferRecognitionRequest]
// object instead. Use the properties of this class to configure various
// aspects of your request object before you start the speech recognition
// process. For example, use the
// [SFSpeechRecognitionRequest.ShouldReportPartialResults] property to specify
// whether you want partial results or only the final result of speech
// recognition.
//
// # Configuring a recognition request
//
//   - [SFSpeechRecognitionRequest.RequiresOnDeviceRecognition]: A Boolean value that determines whether a request must keep its audio data on the device.
//   - [SFSpeechRecognitionRequest.SetRequiresOnDeviceRecognition]
//   - [SFSpeechRecognitionRequest.ShouldReportPartialResults]: A Boolean value that indicates whether you want intermediate results returned for each utterance.
//   - [SFSpeechRecognitionRequest.SetShouldReportPartialResults]
//   - [SFSpeechRecognitionRequest.ContextualStrings]: An array of phrases that should be recognized, even if they are not in the system vocabulary.
//   - [SFSpeechRecognitionRequest.SetContextualStrings]
//
// # Speech Type Classification
//
//   - [SFSpeechRecognitionRequest.TaskHint]: A value that indicates the type of speech recognition being performed.
//   - [SFSpeechRecognitionRequest.SetTaskHint]
//
// # Punctuation
//
//   - [SFSpeechRecognitionRequest.AddsPunctuation]: A Boolean value that indicates whether to add punctuation to speech recognition results.
//   - [SFSpeechRecognitionRequest.SetAddsPunctuation]
//
// # Instance Properties
//
//   - [SFSpeechRecognitionRequest.CustomizedLanguageModel]
//   - [SFSpeechRecognitionRequest.SetCustomizedLanguageModel]
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest
type SFSpeechRecognitionRequest struct {
	objectivec.Object
}

// SFSpeechRecognitionRequestFromID constructs a [SFSpeechRecognitionRequest] from an objc.ID.
//
// An abstract class that represents a request to recognize speech from an
// audio source.
func SFSpeechRecognitionRequestFromID(id objc.ID) SFSpeechRecognitionRequest {
	return SFSpeechRecognitionRequest{objectivec.Object{ID: id}}
}

// NOTE: SFSpeechRecognitionRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SFSpeechRecognitionRequest] class.
//
// # Configuring a recognition request
//
//   - [ISFSpeechRecognitionRequest.RequiresOnDeviceRecognition]: A Boolean value that determines whether a request must keep its audio data on the device.
//   - [ISFSpeechRecognitionRequest.SetRequiresOnDeviceRecognition]
//   - [ISFSpeechRecognitionRequest.ShouldReportPartialResults]: A Boolean value that indicates whether you want intermediate results returned for each utterance.
//   - [ISFSpeechRecognitionRequest.SetShouldReportPartialResults]
//   - [ISFSpeechRecognitionRequest.ContextualStrings]: An array of phrases that should be recognized, even if they are not in the system vocabulary.
//   - [ISFSpeechRecognitionRequest.SetContextualStrings]
//
// # Speech Type Classification
//
//   - [ISFSpeechRecognitionRequest.TaskHint]: A value that indicates the type of speech recognition being performed.
//   - [ISFSpeechRecognitionRequest.SetTaskHint]
//
// # Punctuation
//
//   - [ISFSpeechRecognitionRequest.AddsPunctuation]: A Boolean value that indicates whether to add punctuation to speech recognition results.
//   - [ISFSpeechRecognitionRequest.SetAddsPunctuation]
//
// # Instance Properties
//
//   - [ISFSpeechRecognitionRequest.CustomizedLanguageModel]
//   - [ISFSpeechRecognitionRequest.SetCustomizedLanguageModel]
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest
type ISFSpeechRecognitionRequest interface {
	objectivec.IObject

	// Topic: Configuring a recognition request

	// A Boolean value that determines whether a request must keep its audio data on the device.
	RequiresOnDeviceRecognition() bool
	SetRequiresOnDeviceRecognition(value bool)
	// A Boolean value that indicates whether you want intermediate results returned for each utterance.
	ShouldReportPartialResults() bool
	SetShouldReportPartialResults(value bool)
	// An array of phrases that should be recognized, even if they are not in the system vocabulary.
	ContextualStrings() []string
	SetContextualStrings(value []string)

	// Topic: Speech Type Classification

	// A value that indicates the type of speech recognition being performed.
	TaskHint() SFSpeechRecognitionTaskHint
	SetTaskHint(value SFSpeechRecognitionTaskHint)

	// Topic: Punctuation

	// A Boolean value that indicates whether to add punctuation to speech recognition results.
	AddsPunctuation() bool
	SetAddsPunctuation(value bool)

	// Topic: Instance Properties

	CustomizedLanguageModel() ISFSpeechLanguageModelConfiguration
	SetCustomizedLanguageModel(value ISFSpeechLanguageModelConfiguration)
}

// Init initializes the instance.
func (s SFSpeechRecognitionRequest) Init() SFSpeechRecognitionRequest {
	rv := objc.Send[SFSpeechRecognitionRequest](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SFSpeechRecognitionRequest) Autorelease() SFSpeechRecognitionRequest {
	rv := objc.Send[SFSpeechRecognitionRequest](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSpeechRecognitionRequest creates a new SFSpeechRecognitionRequest instance.
func NewSFSpeechRecognitionRequest() SFSpeechRecognitionRequest {
	class := getSFSpeechRecognitionRequestClass()
	rv := objc.Send[SFSpeechRecognitionRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that determines whether a request must keep its audio data
// on the device.
//
// # Discussion
//
// Set this property to `true` to prevent an [SFSpeechRecognitionRequest] from
// sending audio over the network. However, on-device requests won’t be as
// accurate.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/requiresOnDeviceRecognition
func (s SFSpeechRecognitionRequest) RequiresOnDeviceRecognition() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("requiresOnDeviceRecognition"))
	return rv
}
func (s SFSpeechRecognitionRequest) SetRequiresOnDeviceRecognition(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setRequiresOnDeviceRecognition:"), value)
}

// A Boolean value that indicates whether you want intermediate results
// returned for each utterance.
//
// # Discussion
//
// The default value of this property is `true`. If you want only final
// results (and you don’t care about intermediate results), set this
// property to `false` to prevent the system from doing extra work.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/shouldReportPartialResults
func (s SFSpeechRecognitionRequest) ShouldReportPartialResults() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("shouldReportPartialResults"))
	return rv
}
func (s SFSpeechRecognitionRequest) SetShouldReportPartialResults(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setShouldReportPartialResults:"), value)
}

// An array of phrases that should be recognized, even if they are not in the
// system vocabulary.
//
// # Discussion
//
// Use this property to specify short custom phrases that are unique to your
// app. You might include phrases with the names of characters, products, or
// places that are specific to your app. You might also include
// domain-specific terminology or unusual or made-up words. Assigning custom
// phrases to this property improves the likelihood of those phrases being
// recognized.
//
// Keep phrases relatively brief, limiting them to one or two words whenever
// possible. Lengthy phrases are less likely to be recognized. In addition,
// try to limit each phrase to something the user can say without pausing.
//
// Limit the total number of phrases to no more than 100.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/contextualStrings
func (s SFSpeechRecognitionRequest) ContextualStrings() []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("contextualStrings"))
	return objc.ConvertSliceToStrings(rv)
}
func (s SFSpeechRecognitionRequest) SetContextualStrings(value []string) {
	objc.Send[struct{}](s.ID, objc.Sel("setContextualStrings:"), objectivec.StringSliceToNSArray(value))
}

// A value that indicates the type of speech recognition being performed.
//
// # Discussion
//
// The default value of this property is
// [SFSpeechRecognitionTaskHint.unspecified]. For a valid list of values, see
// [SFSpeechRecognitionTaskHint].
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/taskHint
//
// [SFSpeechRecognitionTaskHint.unspecified]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskHint/unspecified
// [SFSpeechRecognitionTaskHint]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskHint
func (s SFSpeechRecognitionRequest) TaskHint() SFSpeechRecognitionTaskHint {
	rv := objc.Send[SFSpeechRecognitionTaskHint](s.ID, objc.Sel("taskHint"))
	return SFSpeechRecognitionTaskHint(rv)
}
func (s SFSpeechRecognitionRequest) SetTaskHint(value SFSpeechRecognitionTaskHint) {
	objc.Send[struct{}](s.ID, objc.Sel("setTaskHint:"), value)
}

// A Boolean value that indicates whether to add punctuation to speech
// recognition results.
//
// # Discussion
//
// Set this property to `true` for the speech framework to automatically
// include punctuation in the recognition results. Punctuation includes a
// period or question mark at the end of a sentence, and a comma within a
// sentence.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/addsPunctuation
func (s SFSpeechRecognitionRequest) AddsPunctuation() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("addsPunctuation"))
	return rv
}
func (s SFSpeechRecognitionRequest) SetAddsPunctuation(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setAddsPunctuation:"), value)
}

// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionRequest/customizedLanguageModel
func (s SFSpeechRecognitionRequest) CustomizedLanguageModel() ISFSpeechLanguageModelConfiguration {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("customizedLanguageModel"))
	return SFSpeechLanguageModelConfigurationFromID(objc.ID(rv))
}
func (s SFSpeechRecognitionRequest) SetCustomizedLanguageModel(value ISFSpeechLanguageModelConfiguration) {
	objc.Send[struct{}](s.ID, objc.Sel("setCustomizedLanguageModel:"), value)
}

// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNRecognizeTextRequest] class.
var (
	_VNRecognizeTextRequestClass     VNRecognizeTextRequestClass
	_VNRecognizeTextRequestClassOnce sync.Once
)

func getVNRecognizeTextRequestClass() VNRecognizeTextRequestClass {
	_VNRecognizeTextRequestClassOnce.Do(func() {
		_VNRecognizeTextRequestClass = VNRecognizeTextRequestClass{class: objc.GetClass("VNRecognizeTextRequest")}
	})
	return _VNRecognizeTextRequestClass
}

// GetVNRecognizeTextRequestClass returns the class object for VNRecognizeTextRequest.
func GetVNRecognizeTextRequestClass() VNRecognizeTextRequestClass {
	return getVNRecognizeTextRequestClass()
}

type VNRecognizeTextRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNRecognizeTextRequestClass) Alloc() VNRecognizeTextRequest {
	rv := objc.Send[VNRecognizeTextRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An image-analysis request that finds and recognizes text in an image.
//
// # Overview
// 
// By default, a text recognition request first locates all possible glyphs or
// characters in the input image, and then analyzes each string. To specify or
// limit the languages to find in the request, set the [VNRecognizeTextRequest.RecognitionLanguages]
// property to an array that contains the names of the languages of text you
// want to recognize. Vision returns the result of this request in a
// [VNRecognizedTextObservation] object.
//
// # Customizing Recognition Constraints
//
//   - [VNRecognizeTextRequest.MinimumTextHeight]: The minimum height, relative to the image height, of the text to recognize.
//   - [VNRecognizeTextRequest.SetMinimumTextHeight]
//   - [VNRecognizeTextRequest.RecognitionLevel]: A value that determines whether the request prioritizes accuracy or speed in text recognition.
//   - [VNRecognizeTextRequest.SetRecognitionLevel]
//
// # Specifying the Language
//
//   - [VNRecognizeTextRequest.AutomaticallyDetectsLanguage]: A Boolean value that indicates whether to attempt detecting the language to use the appropriate model for recognition and language correction.
//   - [VNRecognizeTextRequest.SetAutomaticallyDetectsLanguage]
//   - [VNRecognizeTextRequest.RecognitionLanguages]: An array of languages to detect, in priority order.
//   - [VNRecognizeTextRequest.SetRecognitionLanguages]
//   - [VNRecognizeTextRequest.UsesLanguageCorrection]: A Boolean value that indicates whether the request applies language correction during the recognition process.
//   - [VNRecognizeTextRequest.SetUsesLanguageCorrection]
//   - [VNRecognizeTextRequest.CustomWords]: An array of strings to supplement the recognized languages at the word-recognition stage.
//   - [VNRecognizeTextRequest.SetCustomWords]
//   - [VNRecognizeTextRequest.SupportedRecognitionLanguagesAndReturnError]: Returns the identifiers of the languages that the request supports.
//
// # Identifying Request Revisions
//
//   - [VNRecognizeTextRequest.VNRecognizeTextRequestRevision3]: A constant for specifying revision 3 of the text recognition request.
//   - [VNRecognizeTextRequest.VNRecognizeTextRequestRevision2]: A constant for specifying revision 2 of the text recognition request.
//   - [VNRecognizeTextRequest.VNRecognizeTextRequestRevision1]: A constant for specifying revision 1 of the text recognition request.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest
type VNRecognizeTextRequest struct {
	VNImageBasedRequest
}

// VNRecognizeTextRequestFromID constructs a [VNRecognizeTextRequest] from an objc.ID.
//
// An image-analysis request that finds and recognizes text in an image.
func VNRecognizeTextRequestFromID(id objc.ID) VNRecognizeTextRequest {
	return VNRecognizeTextRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}
// NOTE: VNRecognizeTextRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNRecognizeTextRequest] class.
//
// # Customizing Recognition Constraints
//
//   - [IVNRecognizeTextRequest.MinimumTextHeight]: The minimum height, relative to the image height, of the text to recognize.
//   - [IVNRecognizeTextRequest.SetMinimumTextHeight]
//   - [IVNRecognizeTextRequest.RecognitionLevel]: A value that determines whether the request prioritizes accuracy or speed in text recognition.
//   - [IVNRecognizeTextRequest.SetRecognitionLevel]
//
// # Specifying the Language
//
//   - [IVNRecognizeTextRequest.AutomaticallyDetectsLanguage]: A Boolean value that indicates whether to attempt detecting the language to use the appropriate model for recognition and language correction.
//   - [IVNRecognizeTextRequest.SetAutomaticallyDetectsLanguage]
//   - [IVNRecognizeTextRequest.RecognitionLanguages]: An array of languages to detect, in priority order.
//   - [IVNRecognizeTextRequest.SetRecognitionLanguages]
//   - [IVNRecognizeTextRequest.UsesLanguageCorrection]: A Boolean value that indicates whether the request applies language correction during the recognition process.
//   - [IVNRecognizeTextRequest.SetUsesLanguageCorrection]
//   - [IVNRecognizeTextRequest.CustomWords]: An array of strings to supplement the recognized languages at the word-recognition stage.
//   - [IVNRecognizeTextRequest.SetCustomWords]
//   - [IVNRecognizeTextRequest.SupportedRecognitionLanguagesAndReturnError]: Returns the identifiers of the languages that the request supports.
//
// # Identifying Request Revisions
//
//   - [IVNRecognizeTextRequest.VNRecognizeTextRequestRevision3]: A constant for specifying revision 3 of the text recognition request.
//   - [IVNRecognizeTextRequest.VNRecognizeTextRequestRevision2]: A constant for specifying revision 2 of the text recognition request.
//   - [IVNRecognizeTextRequest.VNRecognizeTextRequestRevision1]: A constant for specifying revision 1 of the text recognition request.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest
type IVNRecognizeTextRequest interface {
	IVNImageBasedRequest
	VNRequestProgressProviding

	// Topic: Customizing Recognition Constraints

	// The minimum height, relative to the image height, of the text to recognize.
	MinimumTextHeight() float32
	SetMinimumTextHeight(value float32)
	// A value that determines whether the request prioritizes accuracy or speed in text recognition.
	RecognitionLevel() VNRequestTextRecognitionLevel
	SetRecognitionLevel(value VNRequestTextRecognitionLevel)

	// Topic: Specifying the Language

	// A Boolean value that indicates whether to attempt detecting the language to use the appropriate model for recognition and language correction.
	AutomaticallyDetectsLanguage() bool
	SetAutomaticallyDetectsLanguage(value bool)
	// An array of languages to detect, in priority order.
	RecognitionLanguages() []string
	SetRecognitionLanguages(value []string)
	// A Boolean value that indicates whether the request applies language correction during the recognition process.
	UsesLanguageCorrection() bool
	SetUsesLanguageCorrection(value bool)
	// An array of strings to supplement the recognized languages at the word-recognition stage.
	CustomWords() []string
	SetCustomWords(value []string)
	// Returns the identifiers of the languages that the request supports.
	SupportedRecognitionLanguagesAndReturnError() ([]string, error)

	// Topic: Identifying Request Revisions

	// A constant for specifying revision 3 of the text recognition request.
	VNRecognizeTextRequestRevision3() int
	// A constant for specifying revision 2 of the text recognition request.
	VNRecognizeTextRequestRevision2() int
	// A constant for specifying revision 1 of the text recognition request.
	VNRecognizeTextRequestRevision1() int
}

// Init initializes the instance.
func (r VNRecognizeTextRequest) Init() VNRecognizeTextRequest {
	rv := objc.Send[VNRecognizeTextRequest](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r VNRecognizeTextRequest) Autorelease() VNRecognizeTextRequest {
	rv := objc.Send[VNRecognizeTextRequest](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNRecognizeTextRequest creates a new VNRecognizeTextRequest instance.
func NewVNRecognizeTextRequest() VNRecognizeTextRequest {
	class := getVNRecognizeTextRequestClass()
	rv := objc.Send[VNRecognizeTextRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new Vision request with an optional completion handler.
//
// completionHandler: The block to invoke after the request finishes processing.
//
// # Discussion
// 
// Vision executes the completion handler on the same queue that it executes
// the request; however, this queue differs from the one where you called
// [PerformRequestsError].
//
// See: https://developer.apple.com/documentation/Vision/VNRequest/init(completionHandler:)
func NewRecognizeTextRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNRecognizeTextRequest {
	instance := getVNRecognizeTextRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNRecognizeTextRequestFromID(rv)
}

// Returns the identifiers of the languages that the request supports.
//
// # Return Value
// 
// The language identifiers.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/supportedRecognitionLanguages()
func (r VNRecognizeTextRequest) SupportedRecognitionLanguagesAndReturnError() ([]string, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("supportedRecognitionLanguagesAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objc.ConvertSliceToStrings(rv), nil

}

// The minimum height, relative to the image height, of the text to recognize.
//
// # Discussion
// 
// Specify a floating-point number relative to the image height. For example,
// to limit recognition to text that’s half of the image height, use `0.5`.
// Increasing the size reduces memory consumption and expedites recognition
// with the tradeoff of ignoring text smaller than the minimum height. The
// default value is 1/32, or `0.03125`.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/minimumTextHeight
func (r VNRecognizeTextRequest) MinimumTextHeight() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("minimumTextHeight"))
	return rv
}
func (r VNRecognizeTextRequest) SetMinimumTextHeight(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setMinimumTextHeight:"), value)
}

// A value that determines whether the request prioritizes accuracy or speed
// in text recognition.
//
// # Discussion
// 
// The recognition level determines which techniques the request uses during
// the text recognition. Set this value to [RequestTextRecognitionLevelFast]
// to prioritize speed over accuracy, and to
// [RequestTextRecognitionLevelAccurate] for longer, more computationally
// intensive recognition.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/recognitionLevel
func (r VNRecognizeTextRequest) RecognitionLevel() VNRequestTextRecognitionLevel {
	rv := objc.Send[VNRequestTextRecognitionLevel](r.ID, objc.Sel("recognitionLevel"))
	return VNRequestTextRecognitionLevel(rv)
}
func (r VNRecognizeTextRequest) SetRecognitionLevel(value VNRequestTextRecognitionLevel) {
	objc.Send[struct{}](r.ID, objc.Sel("setRecognitionLevel:"), value)
}

// A Boolean value that indicates whether to attempt detecting the language to
// use the appropriate model for recognition and language correction.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/automaticallyDetectsLanguage
func (r VNRecognizeTextRequest) AutomaticallyDetectsLanguage() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("automaticallyDetectsLanguage"))
	return rv
}
func (r VNRecognizeTextRequest) SetAutomaticallyDetectsLanguage(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setAutomaticallyDetectsLanguage:"), value)
}

// An array of languages to detect, in priority order.
//
// # Discussion
// 
// The order of the languages in the array defines the order in which
// languages are used during language processing and text recognition.
// 
// Specify the languages as ISO language codes.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/recognitionLanguages
func (r VNRecognizeTextRequest) RecognitionLanguages() []string {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("recognitionLanguages"))
	return objc.ConvertSliceToStrings(rv)
}
func (r VNRecognizeTextRequest) SetRecognitionLanguages(value []string) {
	objc.Send[struct{}](r.ID, objc.Sel("setRecognitionLanguages:"), objectivec.StringSliceToNSArray(value))
}

// A Boolean value that indicates whether the request applies language
// correction during the recognition process.
//
// # Discussion
// 
// When this value is [true], Vision applies language correction during the
// recognition process. Disabling this property returns the raw recognition
// results, which provides performance benefits but less accurate results.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/usesLanguageCorrection
func (r VNRecognizeTextRequest) UsesLanguageCorrection() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("usesLanguageCorrection"))
	return rv
}
func (r VNRecognizeTextRequest) SetUsesLanguageCorrection(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setUsesLanguageCorrection:"), value)
}

// An array of strings to supplement the recognized languages at the
// word-recognition stage.
//
// # Discussion
// 
// Custom words take precedence over the standard lexicon. The request ignores
// this value if [UsesLanguageCorrection] is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizeTextRequest/customWords
func (r VNRecognizeTextRequest) CustomWords() []string {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("customWords"))
	return objc.ConvertSliceToStrings(rv)
}
func (r VNRecognizeTextRequest) SetCustomWords(value []string) {
	objc.Send[struct{}](r.ID, objc.Sel("setCustomWords:"), objectivec.StringSliceToNSArray(value))
}

// A constant for specifying revision 3 of the text recognition request.
//
// See: https://developer.apple.com/documentation/vision/vnrecognizetextrequestrevision3
func (r VNRecognizeTextRequest) VNRecognizeTextRequestRevision3() int {
	rv := objc.Send[int](r.ID, objc.Sel("VNRecognizeTextRequestRevision3"))
	return rv
}

// A constant for specifying revision 2 of the text recognition request.
//
// See: https://developer.apple.com/documentation/vision/vnrecognizetextrequestrevision2
func (r VNRecognizeTextRequest) VNRecognizeTextRequestRevision2() int {
	rv := objc.Send[int](r.ID, objc.Sel("VNRecognizeTextRequestRevision2"))
	return rv
}

// A constant for specifying revision 1 of the text recognition request.
//
// See: https://developer.apple.com/documentation/vision/vnrecognizetextrequestrevision1
func (r VNRecognizeTextRequest) VNRecognizeTextRequestRevision1() int {
	rv := objc.Send[int](r.ID, objc.Sel("VNRecognizeTextRequestRevision1"))
	return rv
}

// A Boolean set to true when a request can’t determine its progress in
// fractions completed.
//
// # Discussion
// 
// A value of [true] doesn’t mean that the request will run forever. Rather,
// it means that the nature of the request can’t be broken down into
// identifiable fractions to report. The [ProgressHandler] will still be
// called at suitable intervals.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Vision/VNRequestProgressProviding/indeterminate
func (r VNRecognizeTextRequest) Indeterminate() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("indeterminate"))
	return rv
}

// A block of code executed periodically during a Vision request to report
// progress on long-running tasks.
//
// # Discussion
// 
// The progress handler is an optional method that allows clients of the
// request to report progress to the user or to display partial results as
// they become available. The Vision framework may call this handler on a
// different dispatch queue from the thread on which you initiated the
// original request, so ensure that your handler can execute asynchronously,
// in a thread-safe manner.
//
// See: https://developer.apple.com/documentation/Vision/VNRequestProgressProviding/progressHandler
func (r VNRecognizeTextRequest) ProgressHandler() VNRequestProgressHandler {
	rv := objc.Send[VNRequestProgressHandler](r.ID, objc.Sel("progressHandler"))
	return VNRequestProgressHandler(rv)
}
func (r VNRecognizeTextRequest) SetProgressHandler(value VNRequestProgressHandler) {
	objc.Send[struct{}](r.ID, objc.Sel("setProgressHandler:"), value)
}

			// Protocol methods for VNRequestProgressProviding
			


// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SNClassifySoundRequest] class.
var (
	_SNClassifySoundRequestClass     SNClassifySoundRequestClass
	_SNClassifySoundRequestClassOnce sync.Once
)

func getSNClassifySoundRequestClass() SNClassifySoundRequestClass {
	_SNClassifySoundRequestClassOnce.Do(func() {
		_SNClassifySoundRequestClass = SNClassifySoundRequestClass{class: objc.GetClass("SNClassifySoundRequest")}
	})
	return _SNClassifySoundRequestClass
}

// GetSNClassifySoundRequestClass returns the class object for SNClassifySoundRequest.
func GetSNClassifySoundRequestClass() SNClassifySoundRequestClass {
	return getSNClassifySoundRequestClass()
}

type SNClassifySoundRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SNClassifySoundRequestClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SNClassifySoundRequestClass) Alloc() SNClassifySoundRequest {
	rv := objc.Send[SNClassifySoundRequest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A request that classifies sound using a Core ML model.
//
// # Overview
//
// An [SNClassifySoundRequest] represents a specific sound classification
// model. Analyze audio data with a sound classification model by:
//
// - Creating an [SNClassifySoundRequest], either with the Sound Analysis
// model, or by providing your custom Core ML model. - Adding the sound
// request to an [SNAudioFileAnalyzer] or [SNAudioStreamAnalyzer] to process
// an audio file or stream, respectively.
//
// For more information about creating and using classify sound requests, see:
//
// - [Classifying Sounds in an Audio File] - [Classifying Sounds in an Audio
// Stream]
//
// # Creating a Request
//
//   - [SNClassifySoundRequest.InitWithMLModelError]: Creates a request that uses a custom sound classification model.
//   - [SNClassifySoundRequest.InitWithClassifierIdentifierError]: Creates a request that uses the framework’s built-in sound classification model.
//
// # Configuring a Request
//
//   - [SNClassifySoundRequest.OverlapFactor]: The amount of overlap between successive analysis windows when the model operates on a fixed-size audio block.
//   - [SNClassifySoundRequest.SetOverlapFactor]
//   - [SNClassifySoundRequest.WindowDuration]: The duration of the audio buffer the request sends to the underlying sound classifier for each prediction.
//   - [SNClassifySoundRequest.SetWindowDuration]
//
// # Inspecting a Request
//
//   - [SNClassifySoundRequest.KnownClassifications]: A string array that contains every prediction label in the request’s underlying sound classifier model.
//
// # Instance Properties
//
//   - [SNClassifySoundRequest.WindowDurationConstraint]: A range or list of sound duration times the request’s underlying sound classifier supports.
//   - [SNClassifySoundRequest.SetWindowDurationConstraint]
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest
//
// [Classifying Sounds in an Audio File]: https://developer.apple.com/documentation/SoundAnalysis/classifying-sounds-in-an-audio-file
// [Classifying Sounds in an Audio Stream]: https://developer.apple.com/documentation/SoundAnalysis/classifying-sounds-in-an-audio-stream
type SNClassifySoundRequest struct {
	objectivec.Object
}

// SNClassifySoundRequestFromID constructs a [SNClassifySoundRequest] from an objc.ID.
//
// A request that classifies sound using a Core ML model.
func SNClassifySoundRequestFromID(id objc.ID) SNClassifySoundRequest {
	return SNClassifySoundRequest{objectivec.Object{ID: id}}
}

// NOTE: SNClassifySoundRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SNClassifySoundRequest] class.
//
// # Creating a Request
//
//   - [ISNClassifySoundRequest.InitWithMLModelError]: Creates a request that uses a custom sound classification model.
//   - [ISNClassifySoundRequest.InitWithClassifierIdentifierError]: Creates a request that uses the framework’s built-in sound classification model.
//
// # Configuring a Request
//
//   - [ISNClassifySoundRequest.OverlapFactor]: The amount of overlap between successive analysis windows when the model operates on a fixed-size audio block.
//   - [ISNClassifySoundRequest.SetOverlapFactor]
//   - [ISNClassifySoundRequest.WindowDuration]: The duration of the audio buffer the request sends to the underlying sound classifier for each prediction.
//   - [ISNClassifySoundRequest.SetWindowDuration]
//
// # Inspecting a Request
//
//   - [ISNClassifySoundRequest.KnownClassifications]: A string array that contains every prediction label in the request’s underlying sound classifier model.
//
// # Instance Properties
//
//   - [ISNClassifySoundRequest.WindowDurationConstraint]: A range or list of sound duration times the request’s underlying sound classifier supports.
//   - [ISNClassifySoundRequest.SetWindowDurationConstraint]
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest
type ISNClassifySoundRequest interface {
	objectivec.IObject
	SNRequest

	// Topic: Creating a Request

	// Creates a request that uses a custom sound classification model.
	InitWithMLModelError(mlModel *coreml.MLModel) (SNClassifySoundRequest, error)
	// Creates a request that uses the framework’s built-in sound classification model.
	InitWithClassifierIdentifierError(classifierIdentifier SNClassifierIdentifier) (SNClassifySoundRequest, error)

	// Topic: Configuring a Request

	// The amount of overlap between successive analysis windows when the model operates on a fixed-size audio block.
	OverlapFactor() float64
	SetOverlapFactor(value float64)
	// The duration of the audio buffer the request sends to the underlying sound classifier for each prediction.
	WindowDuration() coremedia.CMTime
	SetWindowDuration(value coremedia.CMTime)

	// Topic: Inspecting a Request

	// A string array that contains every prediction label in the request’s underlying sound classifier model.
	KnownClassifications() []string

	// Topic: Instance Properties

	// A range or list of sound duration times the request’s underlying sound classifier supports.
	WindowDurationConstraint() ISNTimeDurationConstraint
	SetWindowDurationConstraint(value ISNTimeDurationConstraint)
}

// Init initializes the instance.
func (c SNClassifySoundRequest) Init() SNClassifySoundRequest {
	rv := objc.Send[SNClassifySoundRequest](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c SNClassifySoundRequest) Autorelease() SNClassifySoundRequest {
	rv := objc.Send[SNClassifySoundRequest](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewSNClassifySoundRequest creates a new SNClassifySoundRequest instance.
func NewSNClassifySoundRequest() SNClassifySoundRequest {
	class := getSNClassifySoundRequestClass()
	rv := objc.Send[SNClassifySoundRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a request that uses the framework’s built-in sound classification
// model.
//
// classifierIdentifier: A sound classifier version identifier, such as [version1].
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/init(classifierIdentifier:)
//
// [version1]: https://developer.apple.com/documentation/SoundAnalysis/SNClassifierIdentifier/version1
func NewClassifySoundRequestWithClassifierIdentifierError(classifierIdentifier SNClassifierIdentifier) (SNClassifySoundRequest, error) {
	var errorPtr objc.ID
	instance := getSNClassifySoundRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithClassifierIdentifier:error:"), objc.String(string(classifierIdentifier)), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SNClassifySoundRequest{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SNClassifySoundRequest{}, objc.ErrInitFailed
	}
	return SNClassifySoundRequestFromID(rv), nil
}

// Creates a request that uses a custom sound classification model.
//
// mlModel: A Core ML sound classification model.
//
// # Discussion
//
// The model you provide must accept audio data as input and produce a
// classification dictionary output that contains the probability of each
// category. For example, you can generate a sound classifier model by
// creating an [MLSoundClassifier] and training it with your own audio files.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/init(mlModel:)
//
// [MLSoundClassifier]: https://developer.apple.com/documentation/CreateML/MLSoundClassifier
func NewClassifySoundRequestWithMLModelError(mlModel *coreml.MLModel) (SNClassifySoundRequest, error) {
	var errorPtr objc.ID
	instance := getSNClassifySoundRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMLModel:error:"), mlModel.ID, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SNClassifySoundRequest{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SNClassifySoundRequest{}, objc.ErrInitFailed
	}
	return SNClassifySoundRequestFromID(rv), nil
}

// Creates a request that uses a custom sound classification model.
//
// mlModel: A Core ML sound classification model.
//
// # Discussion
//
// The model you provide must accept audio data as input and produce a
// classification dictionary output that contains the probability of each
// category. For example, you can generate a sound classifier model by
// creating an [MLSoundClassifier] and training it with your own audio files.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/init(mlModel:)
//
// [MLSoundClassifier]: https://developer.apple.com/documentation/CreateML/MLSoundClassifier
func (c SNClassifySoundRequest) InitWithMLModelError(mlModel *coreml.MLModel) (SNClassifySoundRequest, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("initWithMLModel:error:"), mlModel.ID, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SNClassifySoundRequest{}, foundation.NSErrorFrom(errorPtr)
	}
	return SNClassifySoundRequestFromID(rv), nil

}

// Creates a request that uses the framework’s built-in sound classification
// model.
//
// classifierIdentifier: A sound classifier version identifier, such as [version1].
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/init(classifierIdentifier:)
//
// [version1]: https://developer.apple.com/documentation/SoundAnalysis/SNClassifierIdentifier/version1
func (c SNClassifySoundRequest) InitWithClassifierIdentifierError(classifierIdentifier SNClassifierIdentifier) (SNClassifySoundRequest, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("initWithClassifierIdentifier:error:"), objc.String(string(classifierIdentifier)), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SNClassifySoundRequest{}, foundation.NSErrorFrom(errorPtr)
	}
	return SNClassifySoundRequestFromID(rv), nil

}

// The amount of overlap between successive analysis windows when the model
// operates on a fixed-size audio block.
//
// # Discussion
//
// The property defaults to `0.5` (50%) and supports values in the range
// `[0.0, 1.0]`.
//
// Sound analyses that use a fixed-size audio block typically benefit with an
// overlap factor that’s greater than zero. An overlap factor of `0.0` may
// negatively affect your sound classifier’s accuracy because a sound may
// span two analysis windows. However, an overlap factor of `0.5` ensures each
// sound aligns near the center of at least one analysis window.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/overlapFactor
func (c SNClassifySoundRequest) OverlapFactor() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("overlapFactor"))
	return rv
}
func (c SNClassifySoundRequest) SetOverlapFactor(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setOverlapFactor:"), value)
}

// The duration of the audio buffer the request sends to the underlying sound
// classifier for each prediction.
//
// # Discussion
//
// Configure the window duration with a value that satisfies the request’s
// [SNClassifySoundRequest.WindowDurationConstraint].
//
// The request sends larger audio buffer windows less frequently, which can
// make the classifications more accurate, but less precise in indicating when
// they occur. Requests with smaller buffer window durations sharpen the time
// resolution of each prediction, but send smaller audio buffers to the sound
// classifier.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/windowDuration
func (c SNClassifySoundRequest) WindowDuration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](c.ID, objc.Sel("windowDuration"))
	return coremedia.CMTime(rv)
}
func (c SNClassifySoundRequest) SetWindowDuration(value coremedia.CMTime) {
	objc.Send[struct{}](c.ID, objc.Sel("setWindowDuration:"), value)
}

// A string array that contains every prediction label in the request’s
// underlying sound classifier model.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/knownClassifications
func (c SNClassifySoundRequest) KnownClassifications() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("knownClassifications"))
	return objc.ConvertSliceToStrings(rv)
}

// A range or list of sound duration times the request’s underlying sound
// classifier supports.
//
// See: https://developer.apple.com/documentation/soundanalysis/snclassifysoundrequest/windowdurationconstraint-5no60
func (c SNClassifySoundRequest) WindowDurationConstraint() ISNTimeDurationConstraint {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("windowDurationConstraint"))
	return SNTimeDurationConstraintFromID(objc.ID(rv))
}
func (c SNClassifySoundRequest) SetWindowDurationConstraint(value ISNTimeDurationConstraint) {
	objc.Send[struct{}](c.ID, objc.Sel("setWindowDurationConstraint:"), value)
}

// Protocol methods for SNRequest

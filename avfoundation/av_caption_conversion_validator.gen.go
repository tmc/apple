// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptionConversionValidator] class.
var (
	_AVCaptionConversionValidatorClass     AVCaptionConversionValidatorClass
	_AVCaptionConversionValidatorClassOnce sync.Once
)

func getAVCaptionConversionValidatorClass() AVCaptionConversionValidatorClass {
	_AVCaptionConversionValidatorClassOnce.Do(func() {
		_AVCaptionConversionValidatorClass = AVCaptionConversionValidatorClass{class: objc.GetClass("AVCaptionConversionValidator")}
	})
	return _AVCaptionConversionValidatorClass
}

// GetAVCaptionConversionValidatorClass returns the class object for AVCaptionConversionValidator.
func GetAVCaptionConversionValidatorClass() AVCaptionConversionValidatorClass {
	return getAVCaptionConversionValidatorClass()
}

type AVCaptionConversionValidatorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptionConversionValidatorClass) Alloc() AVCaptionConversionValidator {
	rv := objc.Send[AVCaptionConversionValidator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that validates captions for a conversion operation.
//
// # Creating a validator
//
//   - [AVCaptionConversionValidator.InitWithCaptionsTimeRangeConversionSettings]: Creates an object that validates captions for a conversion operation.
//
// # Inspecting the validator
//
//   - [AVCaptionConversionValidator.Captions]: The array of captions that the system validates.
//   - [AVCaptionConversionValidator.TimeRange]: The time range of the media timeline in which the captions must exist.
//
// # Validating captions
//
//   - [AVCaptionConversionValidator.ValidateCaptionConversionWithWarningHandler]: Validates the object’s captions.
//   - [AVCaptionConversionValidator.Warnings]: The collection of warnings the validator encountered.
//   - [AVCaptionConversionValidator.StopValidating]: Stops the active validation operation.
//
// # Checking the status
//
//   - [AVCaptionConversionValidator.Status]: A value that indicates the status of validation.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator
type AVCaptionConversionValidator struct {
	objectivec.Object
}

// AVCaptionConversionValidatorFromID constructs a [AVCaptionConversionValidator] from an objc.ID.
//
// An object that validates captions for a conversion operation.
func AVCaptionConversionValidatorFromID(id objc.ID) AVCaptionConversionValidator {
	return AVCaptionConversionValidator{objectivec.Object{ID: id}}
}
// NOTE: AVCaptionConversionValidator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptionConversionValidator] class.
//
// # Creating a validator
//
//   - [IAVCaptionConversionValidator.InitWithCaptionsTimeRangeConversionSettings]: Creates an object that validates captions for a conversion operation.
//
// # Inspecting the validator
//
//   - [IAVCaptionConversionValidator.Captions]: The array of captions that the system validates.
//   - [IAVCaptionConversionValidator.TimeRange]: The time range of the media timeline in which the captions must exist.
//
// # Validating captions
//
//   - [IAVCaptionConversionValidator.ValidateCaptionConversionWithWarningHandler]: Validates the object’s captions.
//   - [IAVCaptionConversionValidator.Warnings]: The collection of warnings the validator encountered.
//   - [IAVCaptionConversionValidator.StopValidating]: Stops the active validation operation.
//
// # Checking the status
//
//   - [IAVCaptionConversionValidator.Status]: A value that indicates the status of validation.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator
type IAVCaptionConversionValidator interface {
	objectivec.IObject

	// Topic: Creating a validator

	// Creates an object that validates captions for a conversion operation.
	InitWithCaptionsTimeRangeConversionSettings(captions []AVCaption, timeRange objectivec.IObject, conversionSettings foundation.INSDictionary) AVCaptionConversionValidator

	// Topic: Inspecting the validator

	// The array of captions that the system validates.
	Captions() []AVCaption
	// The time range of the media timeline in which the captions must exist.
	TimeRange() objectivec.IObject

	// Topic: Validating captions

	// Validates the object’s captions.
	ValidateCaptionConversionWithWarningHandler(handler AVCaptionConversionWarningHandler)
	// The collection of warnings the validator encountered.
	Warnings() []AVCaptionConversionWarning
	// Stops the active validation operation.
	StopValidating()

	// Topic: Checking the status

	// A value that indicates the status of validation.
	Status() AVCaptionConversionValidatorStatus
}

// Init initializes the instance.
func (c AVCaptionConversionValidator) Init() AVCaptionConversionValidator {
	rv := objc.Send[AVCaptionConversionValidator](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptionConversionValidator) Autorelease() AVCaptionConversionValidator {
	rv := objc.Send[AVCaptionConversionValidator](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptionConversionValidator creates a new AVCaptionConversionValidator instance.
func NewAVCaptionConversionValidator() AVCaptionConversionValidator {
	class := getAVCaptionConversionValidatorClass()
	rv := objc.Send[AVCaptionConversionValidator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an object that validates captions for a conversion operation.
//
// captions: The array of captions that the system validates.
//
// timeRange: The time range of the media timeline where the captions exist.
//
// conversionSettings: A dictionary that describes the conversion operation.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/init(captions:timeRange:conversionSettings:)
// timeRange is a [coremedia.CMTimeRange].
func NewCaptionConversionValidatorWithCaptionsTimeRangeConversionSettings(captions []AVCaption, timeRange objectivec.IObject, conversionSettings foundation.INSDictionary) AVCaptionConversionValidator {
	instance := getAVCaptionConversionValidatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCaptions:timeRange:conversionSettings:"), objectivec.IObjectSliceToNSArray(captions), timeRange, conversionSettings)
	return AVCaptionConversionValidatorFromID(rv)
}

// Creates an object that validates captions for a conversion operation.
//
// captions: The array of captions that the system validates.
//
// timeRange: The time range of the media timeline where the captions exist.
//
// conversionSettings: A dictionary that describes the conversion operation.
//
// timeRange is a [coremedia.CMTimeRange].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/init(captions:timeRange:conversionSettings:)
func (c AVCaptionConversionValidator) InitWithCaptionsTimeRangeConversionSettings(captions []AVCaption, timeRange objectivec.IObject, conversionSettings foundation.INSDictionary) AVCaptionConversionValidator {
	rv := objc.Send[AVCaptionConversionValidator](c.ID, objc.Sel("initWithCaptions:timeRange:conversionSettings:"), objectivec.IObjectSliceToNSArray(captions), timeRange, conversionSettings)
	return rv
}
// Validates the object’s captions.
//
// handler: The callback the system invokes when it finishes validation.
//
// # Discussion
// 
// When the object finishes validating and reports all warnings, it invokes
// the callback once with a value of `nil` for its warning parameter. When
// this occurs, the validator’s [Status] value changes to
// [CaptionConversionValidatorStatusCompleted].
// 
// Stop an in-progress validation operation by calling [StopValidating].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/validateCaptionConversion(warningHandler:)
func (c AVCaptionConversionValidator) ValidateCaptionConversionWithWarningHandler(handler AVCaptionConversionWarningHandler) {
_block0, _cleanup0 := NewAVCaptionConversionWarningBlock(handler)
	defer _cleanup0()
	objc.Send[objc.ID](c.ID, objc.Sel("validateCaptionConversionWithWarningHandler:"), _block0)
}
// Stops the active validation operation.
//
// # Discussion
// 
// You can call this method at any time, even within the validator’s
// callback to its handler.
// 
// Calling this method stops validation and changes the [Status] value to
// [CaptionConversionValidatorStatusStopped].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/stopValidating()
func (c AVCaptionConversionValidator) StopValidating() {
	objc.Send[objc.ID](c.ID, objc.Sel("stopValidating"))
}

// A convenience initializer to create an object that validates captions for a
// conversion operation.
//
// captions: The array of captions that the system validates.
//
// timeRange: The time range of the media timeline where the captions exist.
//
// conversionSettings: A dictionary that describes the conversion operation.
//
// timeRange is a [coremedia.CMTimeRange].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/captionConversionValidatorWithCaptions:timeRange:conversionSettings:
func (_AVCaptionConversionValidatorClass AVCaptionConversionValidatorClass) CaptionConversionValidatorWithCaptionsTimeRangeConversionSettings(captions []AVCaption, timeRange objectivec.IObject, conversionSettings foundation.INSDictionary) AVCaptionConversionValidator {
	rv := objc.Send[objc.ID](objc.ID(_AVCaptionConversionValidatorClass.class), objc.Sel("captionConversionValidatorWithCaptions:timeRange:conversionSettings:"), objectivec.IObjectSliceToNSArray(captions), timeRange, conversionSettings)
	return AVCaptionConversionValidatorFromID(rv)
}

// The array of captions that the system validates.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/captions
func (c AVCaptionConversionValidator) Captions() []AVCaption {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("captions"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaption {
		return AVCaptionFromID(id)
	})
}
// The time range of the media timeline in which the captions must exist.
//
// # Discussion
// 
// If captions need to appear only after the start of the associated media,
// the start time of this time range can be less than the start time of the
// first caption’s time range.
// 
// If the media duration is unknown, this time range can have a duration of
// [positiveInfinity]. However, to comprehensively validate the conversion of
// closed captions, set the duration of the time range to the duration of the
// associated media.
//
// [positiveInfinity]: https://developer.apple.com/documentation/CoreMedia/CMTime/positiveInfinity
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/timeRange
func (c AVCaptionConversionValidator) TimeRange() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("timeRange"))
	return objectivec.Object{ID: rv}
}
// The collection of warnings the validator encountered.
//
// # Discussion
// 
// This property value may change while the validator’s status is
// [CaptionConversionValidatorStatusValidating].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/warnings
func (c AVCaptionConversionValidator) Warnings() []AVCaptionConversionWarning {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("warnings"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptionConversionWarning {
		return AVCaptionConversionWarningFromID(id)
	})
}
// A value that indicates the status of validation.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/status-swift.property
func (c AVCaptionConversionValidator) Status() AVCaptionConversionValidatorStatus {
	rv := objc.Send[AVCaptionConversionValidatorStatus](c.ID, objc.Sel("status"))
	return AVCaptionConversionValidatorStatus(rv)
}

// ValidateCaptionConversionWithWarningHandlerSync is a synchronous wrapper around [AVCaptionConversionValidator.ValidateCaptionConversionWithWarningHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c AVCaptionConversionValidator) ValidateCaptionConversionWithWarningHandlerSync(ctx context.Context) (*AVCaptionConversionWarning, error) {
	done := make(chan *AVCaptionConversionWarning, 1)
	c.ValidateCaptionConversionWithWarningHandler(func(val *AVCaptionConversionWarning) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}


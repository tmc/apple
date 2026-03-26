// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptionConversionWarning] class.
var (
	_AVCaptionConversionWarningClass     AVCaptionConversionWarningClass
	_AVCaptionConversionWarningClassOnce sync.Once
)

func getAVCaptionConversionWarningClass() AVCaptionConversionWarningClass {
	_AVCaptionConversionWarningClassOnce.Do(func() {
		_AVCaptionConversionWarningClass = AVCaptionConversionWarningClass{class: objc.GetClass("AVCaptionConversionWarning")}
	})
	return _AVCaptionConversionWarningClass
}

// GetAVCaptionConversionWarningClass returns the class object for AVCaptionConversionWarning.
func GetAVCaptionConversionWarningClass() AVCaptionConversionWarningClass {
	return getAVCaptionConversionWarningClass()
}

type AVCaptionConversionWarningClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptionConversionWarningClass) Alloc() AVCaptionConversionWarning {
	rv := objc.Send[AVCaptionConversionWarning](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a conversion warning produced by a validator.
//
// # Inspecting the warning
//
//   - [AVCaptionConversionWarning.WarningType]: A type that indicates the nature of the validation warning.
//   - [AVCaptionConversionWarning.RangeOfCaptions]: The range of the captions for which the system issued a warning.
//   - [AVCaptionConversionWarning.Adjustment]: A correction the converter makes when it converts a caption to a specific format.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionWarning
type AVCaptionConversionWarning struct {
	objectivec.Object
}

// AVCaptionConversionWarningFromID constructs a [AVCaptionConversionWarning] from an objc.ID.
//
// An object that represents a conversion warning produced by a validator.
func AVCaptionConversionWarningFromID(id objc.ID) AVCaptionConversionWarning {
	return AVCaptionConversionWarning{objectivec.Object{ID: id}}
}
// NOTE: AVCaptionConversionWarning adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptionConversionWarning] class.
//
// # Inspecting the warning
//
//   - [IAVCaptionConversionWarning.WarningType]: A type that indicates the nature of the validation warning.
//   - [IAVCaptionConversionWarning.RangeOfCaptions]: The range of the captions for which the system issued a warning.
//   - [IAVCaptionConversionWarning.Adjustment]: A correction the converter makes when it converts a caption to a specific format.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionWarning
type IAVCaptionConversionWarning interface {
	objectivec.IObject

	// Topic: Inspecting the warning

	// A type that indicates the nature of the validation warning.
	WarningType() AVCaptionConversionWarningType
	// The range of the captions for which the system issued a warning.
	RangeOfCaptions() foundation.NSRange
	// A correction the converter makes when it converts a caption to a specific format.
	Adjustment() IAVCaptionConversionAdjustment

	// The collection of warnings the validator encountered.
	Warnings() IAVCaptionConversionWarning
	SetWarnings(value IAVCaptionConversionWarning)
}

// Init initializes the instance.
func (c AVCaptionConversionWarning) Init() AVCaptionConversionWarning {
	rv := objc.Send[AVCaptionConversionWarning](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptionConversionWarning) Autorelease() AVCaptionConversionWarning {
	rv := objc.Send[AVCaptionConversionWarning](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptionConversionWarning creates a new AVCaptionConversionWarning instance.
func NewAVCaptionConversionWarning() AVCaptionConversionWarning {
	class := getAVCaptionConversionWarningClass()
	rv := objc.Send[AVCaptionConversionWarning](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A type that indicates the nature of the validation warning.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionWarning/warningType-swift.property
func (c AVCaptionConversionWarning) WarningType() AVCaptionConversionWarningType {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("warningType"))
	return AVCaptionConversionWarningType(foundation.NSStringFromID(rv).String())
}
// The range of the captions for which the system issued a warning.
//
// # Discussion
// 
// This object only references captions with the same time range. If captions
// with different start times and durations have similar problems, or if
// individual captions have multiple problems, the validator generates
// separate instances of this class for each problem case.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionWarning/rangeOfCaptions
func (c AVCaptionConversionWarning) RangeOfCaptions() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](c.ID, objc.Sel("rangeOfCaptions"))
	return foundation.NSRange(rv)
}
// A correction the converter makes when it converts a caption to a specific
// format.
//
// # Discussion
// 
// If this value is `nil` and you perform the conversion without correcting
// the problem, the system doesn’t include captions that you indicate in the
// output media data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionWarning/adjustment
func (c AVCaptionConversionWarning) Adjustment() IAVCaptionConversionAdjustment {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("adjustment"))
	return AVCaptionConversionAdjustmentFromID(objc.ID(rv))
}
// The collection of warnings the validator encountered.
//
// See: https://developer.apple.com/documentation/avfoundation/avcaptionconversionvalidator/warnings
func (c AVCaptionConversionWarning) Warnings() IAVCaptionConversionWarning {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("warnings"))
	return AVCaptionConversionWarningFromID(objc.ID(rv))
}
func (c AVCaptionConversionWarning) SetWarnings(value IAVCaptionConversionWarning) {
	objc.Send[struct{}](c.ID, objc.Sel("setWarnings:"), value)
}


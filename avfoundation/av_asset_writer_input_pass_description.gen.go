// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetWriterInputPassDescription] class.
var (
	_AVAssetWriterInputPassDescriptionClass     AVAssetWriterInputPassDescriptionClass
	_AVAssetWriterInputPassDescriptionClassOnce sync.Once
)

func getAVAssetWriterInputPassDescriptionClass() AVAssetWriterInputPassDescriptionClass {
	_AVAssetWriterInputPassDescriptionClassOnce.Do(func() {
		_AVAssetWriterInputPassDescriptionClass = AVAssetWriterInputPassDescriptionClass{class: objc.GetClass("AVAssetWriterInputPassDescription")}
	})
	return _AVAssetWriterInputPassDescriptionClass
}

// GetAVAssetWriterInputPassDescriptionClass returns the class object for AVAssetWriterInputPassDescription.
func GetAVAssetWriterInputPassDescriptionClass() AVAssetWriterInputPassDescriptionClass {
	return getAVAssetWriterInputPassDescriptionClass()
}

type AVAssetWriterInputPassDescriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetWriterInputPassDescriptionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetWriterInputPassDescriptionClass) Alloc() AVAssetWriterInputPassDescription {
	rv := objc.Send[AVAssetWriterInputPassDescription](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that defines the interface to query for the requirements of the
// current pass.
//
// # Getting source time ranges
//
//   - [AVAssetWriterInputPassDescription.SourceTimeRanges]: An array of time ranges.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputPassDescription
type AVAssetWriterInputPassDescription struct {
	objectivec.Object
}

// AVAssetWriterInputPassDescriptionFromID constructs a [AVAssetWriterInputPassDescription] from an objc.ID.
//
// An object that defines the interface to query for the requirements of the
// current pass.
func AVAssetWriterInputPassDescriptionFromID(id objc.ID) AVAssetWriterInputPassDescription {
	return AVAssetWriterInputPassDescription{objectivec.Object{ID: id}}
}

// NOTE: AVAssetWriterInputPassDescription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetWriterInputPassDescription] class.
//
// # Getting source time ranges
//
//   - [IAVAssetWriterInputPassDescription.SourceTimeRanges]: An array of time ranges.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputPassDescription
type IAVAssetWriterInputPassDescription interface {
	objectivec.IObject

	// Topic: Getting source time ranges

	// An array of time ranges.
	SourceTimeRanges() []foundation.NSValue
}

// Init initializes the instance.
func (a AVAssetWriterInputPassDescription) Init() AVAssetWriterInputPassDescription {
	rv := objc.Send[AVAssetWriterInputPassDescription](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetWriterInputPassDescription) Autorelease() AVAssetWriterInputPassDescription {
	rv := objc.Send[AVAssetWriterInputPassDescription](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetWriterInputPassDescription creates a new AVAssetWriterInputPassDescription instance.
func NewAVAssetWriterInputPassDescription() AVAssetWriterInputPassDescription {
	class := getAVAssetWriterInputPassDescriptionClass()
	rv := objc.Send[AVAssetWriterInputPassDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An array of time ranges.
//
// # Discussion
//
// Each element in the array is an [NSValue] object that wraps a [CMTimeRange]
// structure that represents one source time range. The value of this property
// is suitable to pass to the [reset(forReadingTimeRanges:)] method of
// [AVAssetReaderOutput].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputPassDescription/sourceTimeRanges
//
// [CMTimeRange]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange
// [NSValue]: https://developer.apple.com/documentation/Foundation/NSValue
// [reset(forReadingTimeRanges:)]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutput/reset(forReadingTimeRanges:)
func (a AVAssetWriterInputPassDescription) SourceTimeRanges() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("sourceTimeRanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}

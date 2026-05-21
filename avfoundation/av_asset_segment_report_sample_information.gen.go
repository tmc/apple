// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetSegmentReportSampleInformation] class.
var (
	_AVAssetSegmentReportSampleInformationClass     AVAssetSegmentReportSampleInformationClass
	_AVAssetSegmentReportSampleInformationClassOnce sync.Once
)

func getAVAssetSegmentReportSampleInformationClass() AVAssetSegmentReportSampleInformationClass {
	_AVAssetSegmentReportSampleInformationClassOnce.Do(func() {
		_AVAssetSegmentReportSampleInformationClass = AVAssetSegmentReportSampleInformationClass{class: objc.GetClass("AVAssetSegmentReportSampleInformation")}
	})
	return _AVAssetSegmentReportSampleInformationClass
}

// GetAVAssetSegmentReportSampleInformationClass returns the class object for AVAssetSegmentReportSampleInformation.
func GetAVAssetSegmentReportSampleInformationClass() AVAssetSegmentReportSampleInformationClass {
	return getAVAssetSegmentReportSampleInformationClass()
}

type AVAssetSegmentReportSampleInformationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetSegmentReportSampleInformationClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetSegmentReportSampleInformationClass) Alloc() AVAssetSegmentReportSampleInformation {
	rv := objc.Send[AVAssetSegmentReportSampleInformation](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that provides information about sample data in a track.
//
// # Inspecting the information
//
//   - [AVAssetSegmentReportSampleInformation.PresentationTimeStamp]: The presentation timestamp (PTS) of a sample.
//   - [AVAssetSegmentReportSampleInformation.Offset]: The offset of a sample in the segment.
//   - [AVAssetSegmentReportSampleInformation.Length]: The length of the sample data.
//   - [AVAssetSegmentReportSampleInformation.IsSyncSample]: A Boolean value that indicates whether the sample is a key frame.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentReportSampleInformation
type AVAssetSegmentReportSampleInformation struct {
	objectivec.Object
}

// AVAssetSegmentReportSampleInformationFromID constructs a [AVAssetSegmentReportSampleInformation] from an objc.ID.
//
// An object that provides information about sample data in a track.
func AVAssetSegmentReportSampleInformationFromID(id objc.ID) AVAssetSegmentReportSampleInformation {
	return AVAssetSegmentReportSampleInformation{objectivec.Object{ID: id}}
}

// NOTE: AVAssetSegmentReportSampleInformation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetSegmentReportSampleInformation] class.
//
// # Inspecting the information
//
//   - [IAVAssetSegmentReportSampleInformation.PresentationTimeStamp]: The presentation timestamp (PTS) of a sample.
//   - [IAVAssetSegmentReportSampleInformation.Offset]: The offset of a sample in the segment.
//   - [IAVAssetSegmentReportSampleInformation.Length]: The length of the sample data.
//   - [IAVAssetSegmentReportSampleInformation.IsSyncSample]: A Boolean value that indicates whether the sample is a key frame.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentReportSampleInformation
type IAVAssetSegmentReportSampleInformation interface {
	objectivec.IObject

	// Topic: Inspecting the information

	// The presentation timestamp (PTS) of a sample.
	PresentationTimeStamp() coremedia.CMTime
	// The offset of a sample in the segment.
	Offset() int
	// The length of the sample data.
	Length() int
	// A Boolean value that indicates whether the sample is a key frame.
	IsSyncSample() bool
}

// Init initializes the instance.
func (a AVAssetSegmentReportSampleInformation) Init() AVAssetSegmentReportSampleInformation {
	rv := objc.Send[AVAssetSegmentReportSampleInformation](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetSegmentReportSampleInformation) Autorelease() AVAssetSegmentReportSampleInformation {
	rv := objc.Send[AVAssetSegmentReportSampleInformation](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetSegmentReportSampleInformation creates a new AVAssetSegmentReportSampleInformation instance.
func NewAVAssetSegmentReportSampleInformation() AVAssetSegmentReportSampleInformation {
	class := getAVAssetSegmentReportSampleInformationClass()
	rv := objc.Send[AVAssetSegmentReportSampleInformation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The presentation timestamp (PTS) of a sample.
//
// # Discussion
//
// This timestamp may be different from the
// [AVAssetSegmentTrackReport.EarliestPresentationTimeStamp] if the video’s
// author encodes it using frame reordering.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentReportSampleInformation/presentationTimeStamp
func (a AVAssetSegmentReportSampleInformation) PresentationTimeStamp() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](a.ID, objc.Sel("presentationTimeStamp"))
	return coremedia.CMTime(rv)
}

// The offset of a sample in the segment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentReportSampleInformation/offset
func (a AVAssetSegmentReportSampleInformation) Offset() int {
	rv := objc.Send[int](a.ID, objc.Sel("offset"))
	return rv
}

// The length of the sample data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentReportSampleInformation/length
func (a AVAssetSegmentReportSampleInformation) Length() int {
	rv := objc.Send[int](a.ID, objc.Sel("length"))
	return rv
}

// A Boolean value that indicates whether the sample is a key frame.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentReportSampleInformation/isSyncSample
func (a AVAssetSegmentReportSampleInformation) IsSyncSample() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isSyncSample"))
	return rv
}

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetSegmentTrackReport] class.
var (
	_AVAssetSegmentTrackReportClass     AVAssetSegmentTrackReportClass
	_AVAssetSegmentTrackReportClassOnce sync.Once
)

func getAVAssetSegmentTrackReportClass() AVAssetSegmentTrackReportClass {
	_AVAssetSegmentTrackReportClassOnce.Do(func() {
		_AVAssetSegmentTrackReportClass = AVAssetSegmentTrackReportClass{class: objc.GetClass("AVAssetSegmentTrackReport")}
	})
	return _AVAssetSegmentTrackReportClass
}

// GetAVAssetSegmentTrackReportClass returns the class object for AVAssetSegmentTrackReport.
func GetAVAssetSegmentTrackReportClass() AVAssetSegmentTrackReportClass {
	return getAVAssetSegmentTrackReportClass()
}

type AVAssetSegmentTrackReportClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetSegmentTrackReportClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetSegmentTrackReportClass) Alloc() AVAssetSegmentTrackReport {
	rv := objc.Send[AVAssetSegmentTrackReport](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that provides information on a track in segment data.
//
// # Inspecting a report
//
//   - [AVAssetSegmentTrackReport.TrackID]: A persistent unique identifier for a track.
//   - [AVAssetSegmentTrackReport.MediaType]: The type of media a track contains.
//   - [AVAssetSegmentTrackReport.Duration]: The duration of a track.
//   - [AVAssetSegmentTrackReport.EarliestPresentationTimeStamp]: The earliest presentation timestamp (PTS) for this track.
//   - [AVAssetSegmentTrackReport.FirstVideoSampleInformation]: Information about the first video sample in a track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentTrackReport
type AVAssetSegmentTrackReport struct {
	objectivec.Object
}

// AVAssetSegmentTrackReportFromID constructs a [AVAssetSegmentTrackReport] from an objc.ID.
//
// An object that provides information on a track in segment data.
func AVAssetSegmentTrackReportFromID(id objc.ID) AVAssetSegmentTrackReport {
	return AVAssetSegmentTrackReport{objectivec.Object{ID: id}}
}
// NOTE: AVAssetSegmentTrackReport adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetSegmentTrackReport] class.
//
// # Inspecting a report
//
//   - [IAVAssetSegmentTrackReport.TrackID]: A persistent unique identifier for a track.
//   - [IAVAssetSegmentTrackReport.MediaType]: The type of media a track contains.
//   - [IAVAssetSegmentTrackReport.Duration]: The duration of a track.
//   - [IAVAssetSegmentTrackReport.EarliestPresentationTimeStamp]: The earliest presentation timestamp (PTS) for this track.
//   - [IAVAssetSegmentTrackReport.FirstVideoSampleInformation]: Information about the first video sample in a track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentTrackReport
type IAVAssetSegmentTrackReport interface {
	objectivec.IObject

	// Topic: Inspecting a report

	// A persistent unique identifier for a track.
	TrackID() int32
	// The type of media a track contains.
	MediaType() AVMediaType
	// The duration of a track.
	Duration() coremedia.CMTime
	// The earliest presentation timestamp (PTS) for this track.
	EarliestPresentationTimeStamp() coremedia.CMTime
	// Information about the first video sample in a track.
	FirstVideoSampleInformation() IAVAssetSegmentReportSampleInformation

	// The type of segment data.
	SegmentType() AVAssetSegmentType
	SetSegmentType(value AVAssetSegmentType)
	// The reports for the segment’s track data.
	TrackReports() IAVAssetSegmentTrackReport
	SetTrackReports(value IAVAssetSegmentTrackReport)
}

// Init initializes the instance.
func (a AVAssetSegmentTrackReport) Init() AVAssetSegmentTrackReport {
	rv := objc.Send[AVAssetSegmentTrackReport](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetSegmentTrackReport) Autorelease() AVAssetSegmentTrackReport {
	rv := objc.Send[AVAssetSegmentTrackReport](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetSegmentTrackReport creates a new AVAssetSegmentTrackReport instance.
func NewAVAssetSegmentTrackReport() AVAssetSegmentTrackReport {
	class := getAVAssetSegmentTrackReportClass()
	rv := objc.Send[AVAssetSegmentTrackReport](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A persistent unique identifier for a track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentTrackReport/trackID
func (a AVAssetSegmentTrackReport) TrackID() int32 {
	rv := objc.Send[int32](a.ID, objc.Sel("trackID"))
	return rv
}
// The type of media a track contains.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentTrackReport/mediaType
func (a AVAssetSegmentTrackReport) MediaType() AVMediaType {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("mediaType"))
	return AVMediaType(foundation.NSStringFromID(rv).String())
}
// The duration of a track.
//
// # Discussion
// 
// The value is [invalid] if there’s no information available.
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentTrackReport/duration
func (a AVAssetSegmentTrackReport) Duration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](a.ID, objc.Sel("duration"))
	return coremedia.CMTime(rv)
}
// The earliest presentation timestamp (PTS) for this track.
//
// # Discussion
// 
// The value is [invalid] if there’s no information available.
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentTrackReport/earliestPresentationTimeStamp
func (a AVAssetSegmentTrackReport) EarliestPresentationTimeStamp() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](a.ID, objc.Sel("earliestPresentationTimeStamp"))
	return coremedia.CMTime(rv)
}
// Information about the first video sample in a track.
//
// # Discussion
// 
// The value is `nil` if this track isn’t a video track, or if sample
// information isn’t available.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentTrackReport/firstVideoSampleInformation
func (a AVAssetSegmentTrackReport) FirstVideoSampleInformation() IAVAssetSegmentReportSampleInformation {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("firstVideoSampleInformation"))
	return AVAssetSegmentReportSampleInformationFromID(objc.ID(rv))
}
// The type of segment data.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetsegmentreport/segmenttype
func (a AVAssetSegmentTrackReport) SegmentType() AVAssetSegmentType {
	rv := objc.Send[AVAssetSegmentType](a.ID, objc.Sel("segmentType"))
	return AVAssetSegmentType(rv)
}
func (a AVAssetSegmentTrackReport) SetSegmentType(value AVAssetSegmentType) {
	objc.Send[struct{}](a.ID, objc.Sel("setSegmentType:"), value)
}
// The reports for the segment’s track data.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetsegmentreport/trackreports
func (a AVAssetSegmentTrackReport) TrackReports() IAVAssetSegmentTrackReport {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("trackReports"))
	return AVAssetSegmentTrackReportFromID(objc.ID(rv))
}
func (a AVAssetSegmentTrackReport) SetTrackReports(value IAVAssetSegmentTrackReport) {
	objc.Send[struct{}](a.ID, objc.Sel("setTrackReports:"), value)
}


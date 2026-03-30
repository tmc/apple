// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetSegmentReport] class.
var (
	_AVAssetSegmentReportClass     AVAssetSegmentReportClass
	_AVAssetSegmentReportClassOnce sync.Once
)

func getAVAssetSegmentReportClass() AVAssetSegmentReportClass {
	_AVAssetSegmentReportClassOnce.Do(func() {
		_AVAssetSegmentReportClass = AVAssetSegmentReportClass{class: objc.GetClass("AVAssetSegmentReport")}
	})
	return _AVAssetSegmentReportClass
}

// GetAVAssetSegmentReportClass returns the class object for AVAssetSegmentReport.
func GetAVAssetSegmentReportClass() AVAssetSegmentReportClass {
	return getAVAssetSegmentReportClass()
}

type AVAssetSegmentReportClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetSegmentReportClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetSegmentReportClass) Alloc() AVAssetSegmentReport {
	rv := objc.Send[AVAssetSegmentReport](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that provides information about segment data.
//
// # Overview
//
// You receive a segment report through the
// [AssetWriterDidOutputSegmentDataSegmentTypeSegmentReport] delegate method.
//
// # Inspecting a report
//
//   - [AVAssetSegmentReport.SegmentType]: The type of segment data.
//   - [AVAssetSegmentReport.TrackReports]: The reports for the segment’s track data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentReport
type AVAssetSegmentReport struct {
	objectivec.Object
}

// AVAssetSegmentReportFromID constructs a [AVAssetSegmentReport] from an objc.ID.
//
// An object that provides information about segment data.
func AVAssetSegmentReportFromID(id objc.ID) AVAssetSegmentReport {
	return AVAssetSegmentReport{objectivec.Object{ID: id}}
}

// NOTE: AVAssetSegmentReport adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetSegmentReport] class.
//
// # Inspecting a report
//
//   - [IAVAssetSegmentReport.SegmentType]: The type of segment data.
//   - [IAVAssetSegmentReport.TrackReports]: The reports for the segment’s track data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentReport
type IAVAssetSegmentReport interface {
	objectivec.IObject

	// Topic: Inspecting a report

	// The type of segment data.
	SegmentType() AVAssetSegmentType
	// The reports for the segment’s track data.
	TrackReports() []AVAssetSegmentTrackReport
}

// Init initializes the instance.
func (a AVAssetSegmentReport) Init() AVAssetSegmentReport {
	rv := objc.Send[AVAssetSegmentReport](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetSegmentReport) Autorelease() AVAssetSegmentReport {
	rv := objc.Send[AVAssetSegmentReport](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetSegmentReport creates a new AVAssetSegmentReport instance.
func NewAVAssetSegmentReport() AVAssetSegmentReport {
	class := getAVAssetSegmentReportClass()
	rv := objc.Send[AVAssetSegmentReport](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The type of segment data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentReport/segmentType
func (a AVAssetSegmentReport) SegmentType() AVAssetSegmentType {
	rv := objc.Send[AVAssetSegmentType](a.ID, objc.Sel("segmentType"))
	return AVAssetSegmentType(rv)
}

// The reports for the segment’s track data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentReport/trackReports
func (a AVAssetSegmentReport) TrackReports() []AVAssetSegmentTrackReport {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("trackReports"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVAssetSegmentTrackReport {
		return AVAssetSegmentTrackReportFromID(id)
	})
}

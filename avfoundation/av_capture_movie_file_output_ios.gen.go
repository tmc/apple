// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
//go:build ios
// +build ios

package avfoundation

import (
	"github.com/tmc/apple/objc"
)

// Returns a list of supported keys to use in the output settings dictionary.
//
// connection: The connection that delivers the media to encode.
//
// # Return Value
//
// An array of keys that can be set in the
// [SetOutputSettingsForConnection]method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/supportedOutputSettingsKeys(for:)
func (c AVCaptureMovieFileOutput) SupportedOutputSettingsKeysForConnection(connection IAVCaptureConnection) []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("supportedOutputSettingsKeysForConnection:"), connection)
	return objc.ConvertSliceToStrings(rv)
}

// A Boolean value that indicates whether the movie file output records video
// orientation and mirroring information as a metadata track.
//
// connection: A connection delivering video media to the movie file output. This method
// throws an invalid argument exception if the value isn’t a video
// connection or if the connection doesn’t terminate at the movie file
// output.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/recordsVideoOrientationAndMirroringChangesAsMetadataTrack(for:)
func (c AVCaptureMovieFileOutput) RecordsVideoOrientationAndMirroringChangesAsMetadataTrackForConnection(connection IAVCaptureConnection) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("recordsVideoOrientationAndMirroringChangesAsMetadataTrackForConnection:"), connection)
	return rv
}

// Sets whether the movie file output creates a timed metadata track to
// capture changes to the connection’s video orientation and mirroring.
//
// doRecordChanges: A Boolean value that indicates whether to capture orientation and mirroring
// information. The system observes this value only at the start of recording.
// Setting a different value has no effect until you start a new recording.
//
// connection: A connection that delivers video media to the movie file output. This
// method throws an invalid argument exception if the value isn’t a video
// connection or if the connection doesn’t terminate at the movie file
// output.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/setRecordsVideoOrientationAndMirroringChangesAsMetadataTrack(_:for:)
func (c AVCaptureMovieFileOutput) SetRecordsVideoOrientationAndMirroringChangesAsMetadataTrackForConnection(doRecordChanges bool, connection IAVCaptureConnection) {
	objc.Send[objc.ID](c.ID, objc.Sel("setRecordsVideoOrientationAndMirroringChanges:asMetadataTrackForConnection:"), doRecordChanges, connection)
}

// The video codecs types the output supports for recording movie files.
//
// # Discussion
//
// The first codec in this list is the default for recording movie files. To
// record using a different codec, call the [SetOutputSettingsForConnection]
// method, passing a video settings dictionary with a value for
// [AVVideoCodecKey] that matches one of the other values in this list.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/availableVideoCodecTypes
//
// [AVVideoCodecKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecKey
func (c AVCaptureMovieFileOutput) AvailableVideoCodecTypes() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availableVideoCodecTypes"))
	return objc.ConvertSliceToStrings(rv)
}

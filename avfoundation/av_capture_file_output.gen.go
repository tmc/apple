// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVCaptureFileOutput] class.
var (
	_AVCaptureFileOutputClass     AVCaptureFileOutputClass
	_AVCaptureFileOutputClassOnce sync.Once
)

func getAVCaptureFileOutputClass() AVCaptureFileOutputClass {
	_AVCaptureFileOutputClassOnce.Do(func() {
		_AVCaptureFileOutputClass = AVCaptureFileOutputClass{class: objc.GetClass("AVCaptureFileOutput")}
	})
	return _AVCaptureFileOutputClass
}

// GetAVCaptureFileOutputClass returns the class object for AVCaptureFileOutput.
func GetAVCaptureFileOutputClass() AVCaptureFileOutputClass {
	return getAVCaptureFileOutputClass()
}

type AVCaptureFileOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureFileOutputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureFileOutputClass) Alloc() AVCaptureFileOutput {
	rv := objc.Send[AVCaptureFileOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// The abstract superclass for capture outputs that can record captured data
// to a file.
//
// # Setting file output properties
//
//   - [AVCaptureFileOutput.Delegate]: The delegate object for the capture file output.
//   - [AVCaptureFileOutput.SetDelegate]
//   - [AVCaptureFileOutput.MaxRecordedDuration]: The longest duration allowed for the recording.
//   - [AVCaptureFileOutput.SetMaxRecordedDuration]
//   - [AVCaptureFileOutput.MaxRecordedFileSize]: The maximum size, in bytes, of the data that should be recorded by the receiver.
//   - [AVCaptureFileOutput.SetMaxRecordedFileSize]
//   - [AVCaptureFileOutput.MinFreeDiskSpaceLimit]: The minimum amount of free space, in bytes, required for recording to continue on a given volume.
//   - [AVCaptureFileOutput.SetMinFreeDiskSpaceLimit]
//   - [AVCaptureFileOutput.OutputFileURL]: The URL to which output is directed.
//   - [AVCaptureFileOutput.RecordedDuration]: Indicates the duration of the media recorded to the current output file.
//   - [AVCaptureFileOutput.RecordedFileSize]: Indicates the size, in bytes, of the data recorded to the current output file.
//   - [AVCaptureFileOutput.Recording]: Indicates whether recording is in progress.
//   - [AVCaptureFileOutput.RecordingPaused]: Indicates whether recording to the current output file is paused.
//
// # Managing recording
//
//   - [AVCaptureFileOutput.StartRecordingToOutputFileURLRecordingDelegate]: Starts recording media to the specified output URL.
//   - [AVCaptureFileOutput.StopRecording]: Tells the receiver to stop recording to the current file.
//   - [AVCaptureFileOutput.PauseRecording]: Pauses recording to the current output file.
//   - [AVCaptureFileOutput.ResumeRecording]: Resumes recording to the current output file after it was previously paused using [pauseRecording()](<doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureFileOutput/pauseRecording()>).
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput
type AVCaptureFileOutput struct {
	AVCaptureOutput
}

// AVCaptureFileOutputFromID constructs a [AVCaptureFileOutput] from an objc.ID.
//
// The abstract superclass for capture outputs that can record captured data
// to a file.
func AVCaptureFileOutputFromID(id objc.ID) AVCaptureFileOutput {
	return AVCaptureFileOutput{AVCaptureOutput: AVCaptureOutputFromID(id)}
}
// NOTE: AVCaptureFileOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureFileOutput] class.
//
// # Setting file output properties
//
//   - [IAVCaptureFileOutput.Delegate]: The delegate object for the capture file output.
//   - [IAVCaptureFileOutput.SetDelegate]
//   - [IAVCaptureFileOutput.MaxRecordedDuration]: The longest duration allowed for the recording.
//   - [IAVCaptureFileOutput.SetMaxRecordedDuration]
//   - [IAVCaptureFileOutput.MaxRecordedFileSize]: The maximum size, in bytes, of the data that should be recorded by the receiver.
//   - [IAVCaptureFileOutput.SetMaxRecordedFileSize]
//   - [IAVCaptureFileOutput.MinFreeDiskSpaceLimit]: The minimum amount of free space, in bytes, required for recording to continue on a given volume.
//   - [IAVCaptureFileOutput.SetMinFreeDiskSpaceLimit]
//   - [IAVCaptureFileOutput.OutputFileURL]: The URL to which output is directed.
//   - [IAVCaptureFileOutput.RecordedDuration]: Indicates the duration of the media recorded to the current output file.
//   - [IAVCaptureFileOutput.RecordedFileSize]: Indicates the size, in bytes, of the data recorded to the current output file.
//   - [IAVCaptureFileOutput.Recording]: Indicates whether recording is in progress.
//   - [IAVCaptureFileOutput.RecordingPaused]: Indicates whether recording to the current output file is paused.
//
// # Managing recording
//
//   - [IAVCaptureFileOutput.StartRecordingToOutputFileURLRecordingDelegate]: Starts recording media to the specified output URL.
//   - [IAVCaptureFileOutput.StopRecording]: Tells the receiver to stop recording to the current file.
//   - [IAVCaptureFileOutput.PauseRecording]: Pauses recording to the current output file.
//   - [IAVCaptureFileOutput.ResumeRecording]: Resumes recording to the current output file after it was previously paused using [pauseRecording()](<doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureFileOutput/pauseRecording()>).
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput
type IAVCaptureFileOutput interface {
	IAVCaptureOutput

	// Topic: Setting file output properties

	// The delegate object for the capture file output.
	Delegate() AVCaptureFileOutputDelegate
	SetDelegate(value AVCaptureFileOutputDelegate)
	// The longest duration allowed for the recording.
	MaxRecordedDuration() coremedia.CMTime
	SetMaxRecordedDuration(value coremedia.CMTime)
	// The maximum size, in bytes, of the data that should be recorded by the receiver.
	MaxRecordedFileSize() int64
	SetMaxRecordedFileSize(value int64)
	// The minimum amount of free space, in bytes, required for recording to continue on a given volume.
	MinFreeDiskSpaceLimit() int64
	SetMinFreeDiskSpaceLimit(value int64)
	// The URL to which output is directed.
	OutputFileURL() foundation.INSURL
	// Indicates the duration of the media recorded to the current output file.
	RecordedDuration() coremedia.CMTime
	// Indicates the size, in bytes, of the data recorded to the current output file.
	RecordedFileSize() int64
	// Indicates whether recording is in progress.
	Recording() bool
	// Indicates whether recording to the current output file is paused.
	RecordingPaused() bool

	// Topic: Managing recording

	// Starts recording media to the specified output URL.
	StartRecordingToOutputFileURLRecordingDelegate(outputFileURL foundation.INSURL, delegate AVCaptureFileOutputRecordingDelegate)
	// Tells the receiver to stop recording to the current file.
	StopRecording()
	// Pauses recording to the current output file.
	PauseRecording()
	// Resumes recording to the current output file after it was previously paused using [pauseRecording()](<doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureFileOutput/pauseRecording()>).
	ResumeRecording()
}

// Init initializes the instance.
func (c AVCaptureFileOutput) Init() AVCaptureFileOutput {
	rv := objc.Send[AVCaptureFileOutput](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureFileOutput) Autorelease() AVCaptureFileOutput {
	rv := objc.Send[AVCaptureFileOutput](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureFileOutput creates a new AVCaptureFileOutput instance.
func NewAVCaptureFileOutput() AVCaptureFileOutput {
	class := getAVCaptureFileOutputClass()
	rv := objc.Send[AVCaptureFileOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Starts recording media to the specified output URL.
//
// outputFileURL: An object specifying the output file URL.
// 
// This method raises an [InvalidArgumentException] if the argument isn’t a
// valid file URL.
//
// delegate: A delegate object that’s notified of changes to the recording state.
//
// # Discussion
// 
// A failure occurs if you attempt to record to a URL where a file exists. To
// overwrite the content, delete the old file before calling this method.
// 
// In macOS, calling this method from within the
// [CaptureOutputDidOutputSampleBufferFromConnection] method guarantees that
// the first samples written to the new file are those passed to the delegate
// method.
// 
// When you stop recording by calling [StopRecording], by changing files using
// this method, or because of an error, the framework writes any remaining
// file data in the background. Therefore, for the system to notify you upon
// completion, you must adopt the
// [CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError]
// delegate method. The recording delegate can also optionally implement
// methods that inform it when the output object starts writing data, when it
// pauses or resumes recording, and when it’s about to finish recording.
// 
// In macOS, you don’t need to call [StopRecording] before calling this
// method while another recording is in progress. If you call this method
// while the output object is recording, the framework preserves media samples
// between the old file and the new file. In iOS, to avoid any errors, you
// must call [StopRecording] before calling this method again.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/startRecording(to:recordingDelegate:)
func (c AVCaptureFileOutput) StartRecordingToOutputFileURLRecordingDelegate(outputFileURL foundation.INSURL, delegate AVCaptureFileOutputRecordingDelegate) {
	objc.Send[objc.ID](c.ID, objc.Sel("startRecordingToOutputFileURL:recordingDelegate:"), outputFileURL, delegate)
}
// Tells the receiver to stop recording to the current file.
//
// # Discussion
// 
// You can call this method when they want to stop recording new samples to
// the current file, and do not want to continue recording to another file. If
// you want to switch from one file to another, you should not call this
// method. Instead you should simply call
// [StartRecordingToOutputFileURLRecordingDelegate] with the new file URL.
// 
// When recording is stopped either by calling this method, by changing files
// using [StartRecordingToOutputFileURLRecordingDelegate], or because of an
// error, the remaining data that needs to be included to the file will be
// written in the background. Therefore, before using the file, you must wait
// until the delegate that was specified in
// [StartRecordingToOutputFileURLRecordingDelegate] is notified when all data
// has been written to the file using the
// [CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError]
// method.
// 
// In macOS, if this method is called within the
// captureOutput:didOutputSampleBuffer:fromConnection: delegate method, the
// last samples written to the current file are guaranteed to be those that
// were output immediately before those in the sample buffer passed to that
// method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/stopRecording()
func (c AVCaptureFileOutput) StopRecording() {
	objc.Send[objc.ID](c.ID, objc.Sel("stopRecording"))
}
// Pauses recording to the current output file.
//
// # Discussion
// 
// This method causes the receiver to stop writing captured samples to the
// current output file returned by [OutputFileURL], but leaves the file open
// so that samples can be written to it in the future, if [ResumeRecording] is
// called. This allows you to record multiple media segments that are not
// contiguous in time to a single file.
// 
// In macOS, if this method is called within the
// captureOutput:didOutputSampleBuffer:fromConnection: delegate method, the
// last samples written to the current file are guaranteed to be those that
// were output immediately before those in the sample buffer passed to that
// method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/pauseRecording()
func (c AVCaptureFileOutput) PauseRecording() {
	objc.Send[objc.ID](c.ID, objc.Sel("pauseRecording"))
}
// Resumes recording to the current output file after it was previously paused
// using [PauseRecording].
//
// # Discussion
// 
// This method causes the receiver to resume writing captured samples to the
// current output file returned by [OutputFileURL], after recording was
// previously paused using [PauseRecording]. This allows you to record
// multiple media segments that are not contiguous in time to a single file.
// 
// In macOS, if this method is called within the
// captureOutput:didOutputSampleBuffer:fromConnection: delegate method, the
// first samples written to the current file are guaranteed to be those
// contained in the sample buffer passed to that method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/resumeRecording()
func (c AVCaptureFileOutput) ResumeRecording() {
	objc.Send[objc.ID](c.ID, objc.Sel("resumeRecording"))
}

// The delegate object for the capture file output.
//
// # Discussion
// 
// The delegate is an object conforming to the [AVCaptureFileOutputDelegate]
// protocol that will be able to monitor and control recording along exact
// sample boundaries.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/delegate
func (c AVCaptureFileOutput) Delegate() AVCaptureFileOutputDelegate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("delegate"))
	return AVCaptureFileOutputDelegateObjectFromID(rv)
}
func (c AVCaptureFileOutput) SetDelegate(value AVCaptureFileOutputDelegate) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelegate:"), value)
}
// The longest duration allowed for the recording.
//
// # Discussion
// 
// This property specifies a hard limit on the duration of recorded files.
// Recording is stopped when the limit is reached and the
// [CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError]
// delegate method is invoked with an appropriate error. The default value of
// this property is [invalid], which indicates no limit.
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/maxRecordedDuration
func (c AVCaptureFileOutput) MaxRecordedDuration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](c.ID, objc.Sel("maxRecordedDuration"))
	return coremedia.CMTime(rv)
}
func (c AVCaptureFileOutput) SetMaxRecordedDuration(value coremedia.CMTime) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaxRecordedDuration:"), value)
}
// The maximum size, in bytes, of the data that should be recorded by the
// receiver.
//
// # Discussion
// 
// This property specifies a hard limit on the data size of recorded files.
// Recording is stopped when the limit is reached and the
// [CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError]
// delegate method is invoked with an appropriate error. The default value of
// this property is `0`, which indicates no limit.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/maxRecordedFileSize
func (c AVCaptureFileOutput) MaxRecordedFileSize() int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("maxRecordedFileSize"))
	return rv
}
func (c AVCaptureFileOutput) SetMaxRecordedFileSize(value int64) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaxRecordedFileSize:"), value)
}
// The minimum amount of free space, in bytes, required for recording to
// continue on a given volume.
//
// # Discussion
// 
// This property specifies a hard lower limit on the amount of free space that
// must remain on a target volume for recording to continue. Recording is
// stopped when the limit is reached and the
// [CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError]
// delegate method is invoked with an appropriate error.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/minFreeDiskSpaceLimit
func (c AVCaptureFileOutput) MinFreeDiskSpaceLimit() int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("minFreeDiskSpaceLimit"))
	return rv
}
func (c AVCaptureFileOutput) SetMinFreeDiskSpaceLimit(value int64) {
	objc.Send[struct{}](c.ID, objc.Sel("setMinFreeDiskSpaceLimit:"), value)
}
// The URL to which output is directed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/outputFileURL
func (c AVCaptureFileOutput) OutputFileURL() foundation.INSURL {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("outputFileURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// Indicates the duration of the media recorded to the current output file.
//
// # Discussion
// 
// If recording is in progress, this property returns the total time recorded
// so far.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/recordedDuration
func (c AVCaptureFileOutput) RecordedDuration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](c.ID, objc.Sel("recordedDuration"))
	return coremedia.CMTime(rv)
}
// Indicates the size, in bytes, of the data recorded to the current output
// file.
//
// # Discussion
// 
// If a recording is in progress, this property returns the size in bytes of
// the data recorded so far.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/recordedFileSize
func (c AVCaptureFileOutput) RecordedFileSize() int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("recordedFileSize"))
	return rv
}
// Indicates whether recording is in progress.
//
// # Discussion
// 
// The value of this property is [true] when the file output currently has a
// file to which it is writing new samples, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/isRecording
func (c AVCaptureFileOutput) Recording() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isRecording"))
	return rv
}
// Indicates whether recording to the current output file is paused.
//
// # Discussion
// 
// This property indicates recording to the file returned by [OutputFileURL]
// has been previously paused using the [PauseRecording] method. When a
// recording is paused, captured samples are not written to the output file,
// but new samples can be written to the same file in the future by calling
// [ResumeRecording].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/isRecordingPaused
func (c AVCaptureFileOutput) RecordingPaused() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isRecordingPaused"))
	return rv
}


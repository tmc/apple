// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureAudioFileOutput] class.
var (
	_AVCaptureAudioFileOutputClass     AVCaptureAudioFileOutputClass
	_AVCaptureAudioFileOutputClassOnce sync.Once
)

func getAVCaptureAudioFileOutputClass() AVCaptureAudioFileOutputClass {
	_AVCaptureAudioFileOutputClassOnce.Do(func() {
		_AVCaptureAudioFileOutputClass = AVCaptureAudioFileOutputClass{class: objc.GetClass("AVCaptureAudioFileOutput")}
	})
	return _AVCaptureAudioFileOutputClass
}

// GetAVCaptureAudioFileOutputClass returns the class object for AVCaptureAudioFileOutput.
func GetAVCaptureAudioFileOutputClass() AVCaptureAudioFileOutputClass {
	return getAVCaptureAudioFileOutputClass()
}

type AVCaptureAudioFileOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureAudioFileOutputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureAudioFileOutputClass) Alloc() AVCaptureAudioFileOutput {
	rv := objc.Send[AVCaptureAudioFileOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A capture output that records audio and saves the recorded audio to a file.
//
// # Overview
// 
// [AVCaptureAudioFileOutput] implements the complete file recording interface
// declared by [AVCaptureFileOutput] for writing media data to audio files. In
// addition, you can configure options specific to the audio file formats,
// including writing metadata collections to each file and specifying audio
// encoding options. [AVCaptureAudioFileOutput] does not, however, support
// [StartRecordingToOutputFileURLRecordingDelegate]—use
// [AVCaptureAudioFileOutput.StartRecordingToOutputFileURLOutputFileTypeRecordingDelegate] instead.
//
// # Starting a recording
//
//   - [AVCaptureAudioFileOutput.StartRecordingToOutputFileURLOutputFileTypeRecordingDelegate]: Tells the receiver to start recording to a new file of the specified format, and specifies a delegate that will be notified when recording is finished.
//
// # Configuring output
//
//   - [AVCaptureAudioFileOutput.AudioSettings]: The settings used to decode or re-encode audio before it is output by the receiver.
//   - [AVCaptureAudioFileOutput.SetAudioSettings]
//   - [AVCaptureAudioFileOutput.Metadata]: A collection of metadata to be written to the receiver’s output files.
//   - [AVCaptureAudioFileOutput.SetMetadata]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioFileOutput
type AVCaptureAudioFileOutput struct {
	AVCaptureFileOutput
}

// AVCaptureAudioFileOutputFromID constructs a [AVCaptureAudioFileOutput] from an objc.ID.
//
// A capture output that records audio and saves the recorded audio to a file.
func AVCaptureAudioFileOutputFromID(id objc.ID) AVCaptureAudioFileOutput {
	return AVCaptureAudioFileOutput{AVCaptureFileOutput: AVCaptureFileOutputFromID(id)}
}
// NOTE: AVCaptureAudioFileOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureAudioFileOutput] class.
//
// # Starting a recording
//
//   - [IAVCaptureAudioFileOutput.StartRecordingToOutputFileURLOutputFileTypeRecordingDelegate]: Tells the receiver to start recording to a new file of the specified format, and specifies a delegate that will be notified when recording is finished.
//
// # Configuring output
//
//   - [IAVCaptureAudioFileOutput.AudioSettings]: The settings used to decode or re-encode audio before it is output by the receiver.
//   - [IAVCaptureAudioFileOutput.SetAudioSettings]
//   - [IAVCaptureAudioFileOutput.Metadata]: A collection of metadata to be written to the receiver’s output files.
//   - [IAVCaptureAudioFileOutput.SetMetadata]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioFileOutput
type IAVCaptureAudioFileOutput interface {
	IAVCaptureFileOutput

	// Topic: Starting a recording

	// Tells the receiver to start recording to a new file of the specified format, and specifies a delegate that will be notified when recording is finished.
	StartRecordingToOutputFileURLOutputFileTypeRecordingDelegate(outputFileURL foundation.INSURL, fileType AVFileType, delegate AVCaptureFileOutputRecordingDelegate)

	// Topic: Configuring output

	// The settings used to decode or re-encode audio before it is output by the receiver.
	AudioSettings() foundation.INSDictionary
	SetAudioSettings(value foundation.INSDictionary)
	// A collection of metadata to be written to the receiver’s output files.
	Metadata() []AVMetadataItem
	SetMetadata(value []AVMetadataItem)
}

// Init initializes the instance.
func (c AVCaptureAudioFileOutput) Init() AVCaptureAudioFileOutput {
	rv := objc.Send[AVCaptureAudioFileOutput](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureAudioFileOutput) Autorelease() AVCaptureAudioFileOutput {
	rv := objc.Send[AVCaptureAudioFileOutput](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureAudioFileOutput creates a new AVCaptureAudioFileOutput instance.
func NewAVCaptureAudioFileOutput() AVCaptureAudioFileOutput {
	class := getAVCaptureAudioFileOutputClass()
	rv := objc.Send[AVCaptureAudioFileOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Tells the receiver to start recording to a new file of the specified
// format, and specifies a delegate that will be notified when recording is
// finished.
//
// outputFileURL: The URL of the output file.
// 
// This method throws an [InvalidArgumentException] if the URL is not a valid
// file URL.
// 
// If a file at the given URL already exists when capturing starts, recording
// to the new file will fail.
//
// fileType: A UTI indicating the format of the file to be written.
// 
// UTIs for common audio file types are declared in `AVMediaFormat.H()`.
//
// delegate: An object conforming to the [AVCaptureFileOutputRecordingDelegate]
// protocol.
// 
// You must specify a delegate to be notified when recording is finished.
//
// # Discussion
// 
// You do not need not to call [StopRecording] before calling this method
// while another recording is in progress. If this method is invoked while an
// existing output file was already being recorded, no media samples will be
// discarded between the old file and the new file.
// 
// When recording is stopped—by calling `stopRecording`, by changing files
// using this method, or because of an error—the remaining data that needs
// to be included to the file will be written in the background. Therefore,
// you must specify a delegate that will be notified when all data has been
// written to the file using the
// [CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError]
// method. The recording delegate can also optionally implement methods that
// inform it when data starts being written, when recording is paused and
// resumed, and when recording is about to be finished.
// 
// In macOS, if this method is called within the
// [CaptureOutputDidOutputSampleBufferFromConnection] delegate method, the
// first samples written to the new file are guaranteed to be those contained
// in the sample buffer passed to that method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioFileOutput/startRecording(to:outputFileType:recordingDelegate:)
func (c AVCaptureAudioFileOutput) StartRecordingToOutputFileURLOutputFileTypeRecordingDelegate(outputFileURL foundation.INSURL, fileType AVFileType, delegate AVCaptureFileOutputRecordingDelegate) {
	objc.Send[objc.ID](c.ID, objc.Sel("startRecordingToOutputFileURL:outputFileType:recordingDelegate:"), outputFileURL, objc.String(string(fileType)), delegate)
}

// Returns an array containing UTIs identifying the file types
// [AVCaptureAudioFileOutput] can write.
//
// # Return Value
// 
// An array containing UTIs identifying the file types
// [AVCaptureAudioFileOutput] can write.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioFileOutput/availableOutputFileTypes()
func (_AVCaptureAudioFileOutputClass AVCaptureAudioFileOutputClass) AvailableOutputFileTypes() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_AVCaptureAudioFileOutputClass.class), objc.Sel("availableOutputFileTypes"))
	return objc.ConvertSliceToStrings(rv)
}

// The settings used to decode or re-encode audio before it is output by the
// receiver.
//
// # Discussion
// 
// The value of this property is a dictionary containing values for audio
// settings keys defined in `AVAudioSettings.H()`. If you set the value of
// this property to `nil`, the output vends samples in their device native
// format.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioFileOutput/audioSettings
func (c AVCaptureAudioFileOutput) AudioSettings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("audioSettings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (c AVCaptureAudioFileOutput) SetAudioSettings(value foundation.INSDictionary) {
	objc.Send[struct{}](c.ID, objc.Sel("setAudioSettings:"), value)
}
// A collection of metadata to be written to the receiver’s output files.
//
// # Discussion
// 
// The value of this property is an array of [AVMetadataItem] objects
// representing the collection of top-level metadata to be written in each
// output file. Only ID3 v2.2, v2.3, or v2.4 style metadata items are
// supported.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioFileOutput/metadata
func (c AVCaptureAudioFileOutput) Metadata() []AVMetadataItem {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("metadata"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}
func (c AVCaptureAudioFileOutput) SetMetadata(value []AVMetadataItem) {
	objc.Send[struct{}](c.ID, objc.Sel("setMetadata:"), objectivec.IObjectSliceToNSArray(value))
}


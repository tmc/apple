// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// Methods for responding to events that occur while recording captured media to a file.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutputRecordingDelegate
type AVCaptureFileOutputRecordingDelegate interface {
	objectivec.IObject

	// Informs the delegate when all pending data has been written to an output file.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutputRecordingDelegate/fileOutput(_:didFinishRecordingTo:from:error:)
	CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError(output IAVCaptureFileOutput, outputFileURL foundation.INSURL, connections []AVCaptureConnection, error_ foundation.INSError)
}

// AVCaptureFileOutputRecordingDelegateObject wraps an existing Objective-C object that conforms to the AVCaptureFileOutputRecordingDelegate protocol.
type AVCaptureFileOutputRecordingDelegateObject struct {
	objectivec.Object
}
func (o AVCaptureFileOutputRecordingDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVCaptureFileOutputRecordingDelegateObjectFromID constructs a [AVCaptureFileOutputRecordingDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVCaptureFileOutputRecordingDelegateObjectFromID(id objc.ID) AVCaptureFileOutputRecordingDelegateObject {
	return AVCaptureFileOutputRecordingDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Informs the delegate when all pending data has been written to an output
// file.
//
// output: The capture file output that has finished writing the file.
//
// outputFileURL: The file URL of the file that is being written.
//
// connections: An array of [AVCaptureConnection] objects attached to the file output that
// provided the data that is being written to the file.
//
// error: If the file was not written successfully, an error object that describes
// the problem; otherwise `nil`.
//
// # Discussion
// 
// This method is called whenever a file is finished. If the file was forced
// to be finished due to an error, the error is described in the error
// parameter—otherwise, the error parameter is `nil`.
// 
// This method is called when the file output has finished writing all data to
// a file whose recording was stopped, either because
// [StartRecordingToOutputFileURLRecordingDelegate] or [StopRecording] were
// called, or because an error (described by the error parameter) occurred (if
// no error occurred, the error parameter is `nil`).
// 
// This method is always called for each recording request, even if no data is
// successfully written to the file.
// 
// You should not assume that this method will be called on a specific thread.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutputRecordingDelegate/fileOutput(_:didFinishRecordingTo:from:error:)
func (o AVCaptureFileOutputRecordingDelegateObject) CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError(output IAVCaptureFileOutput, outputFileURL foundation.INSURL, connections []AVCaptureConnection, error_ foundation.INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:didFinishRecordingToOutputFileAtURL:fromConnections:error:"), output, outputFileURL, objectivec.IObjectSliceToNSArray(connections), error_)
	}
// Informs the delegate when the output has started writing to a file.
//
// output: The capture file output that started writing the file.
//
// fileURL: The file URL of the file that is being written.
//
// connections: An array of [AVCaptureConnection] objects attached to the file output that
// provided the data that is being written to the file.
//
// # Discussion
// 
// If an error condition prevents any data from being written, this method may
// not be called.
// [CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError] and
// [CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError] are
// always called, even if no data is written.
// 
// You should not assume that this method will be called on a specific thread,
// and should make this method as efficient as possible.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutputRecordingDelegate/fileOutput(_:didStartRecordingTo:from:)
func (o AVCaptureFileOutputRecordingDelegateObject) CaptureOutputDidStartRecordingToOutputFileAtURLFromConnections(output IAVCaptureFileOutput, fileURL foundation.INSURL, connections []AVCaptureConnection) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:didStartRecordingToOutputFileAtURL:fromConnections:"), output, fileURL, objectivec.IObjectSliceToNSArray(connections))
	}
//
// output: The capture file output that started writing the file.
//
// fileURL: The file URL of the file that is being written.
//
// startPTS: The timestamp of the first buffer written to the file, synced with
// AVCaptureSession.synchronizationClock
//
// connections: An array of AVCaptureConnection objects attached to the file output that
// provided the data that is being written to the file.
//
// # Discussion
// 
// Informs the delegate when the output has started writing to a file.
// 
// This method is called when the file output has started writing data to a
// file. If an error condition prevents any data from being written, this
// method may not be called.
// captureOutput:willFinishRecordingToOutputFileAtURL:fromConnections:error:
// and
// captureOutput:didFinishRecordingToOutputFileAtURL:fromConnections:error:
// will always be called, even if no data is written.
// 
// If this method is implemented, the alternative delegate callback
// -captureOutput:didStartRecordingToOutputFileAtURL:fromConnections will not
// be called.
// 
// Clients should not assume that this method will be called on a specific
// thread, and should also try to make this method as efficient as possible.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutputRecordingDelegate/fileOutput(_:didStartRecordingTo:startPTS:from:)
func (o AVCaptureFileOutputRecordingDelegateObject) CaptureOutputDidStartRecordingToOutputFileAtURLStartPTSFromConnections(output IAVCaptureFileOutput, fileURL foundation.INSURL, startPTS coremedia.CMTime, connections []AVCaptureConnection) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:didStartRecordingToOutputFileAtURL:startPTS:fromConnections:"), output, fileURL, startPTS, objectivec.IObjectSliceToNSArray(connections))
	}
// Informs the delegate when the output will stop writing new samples to a
// file.
//
// output: The capture file output that will finish writing the file.
//
// fileURL: The file URL of the file that is being written.
//
// connections: An array of [AVCaptureConnection] objects attached to the file output that
// provided the data that is being written to the file.
//
// error: An error describing what caused the file to stop recording, or `nil` if
// there was no error.
//
// # Discussion
// 
// This method is called when the file output will stop recording new samples
// to the file at [OutputFileURL], either because
// [StartRecordingToOutputFileURLRecordingDelegate] or [StopRecording] was
// called, or because an error (described by the error parameter) occurred (if
// no error occurred, the error parameter is `nil`).
// 
// This method is always called for each recording request, even if no data is
// successfully written to the file.
// 
// You should not assume that this method will be called on a specific thread,
// and should make this method as efficient as possible.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutputRecordingDelegate/fileOutput(_:willFinishRecordingTo:from:error:)
func (o AVCaptureFileOutputRecordingDelegateObject) CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError(output IAVCaptureFileOutput, fileURL foundation.INSURL, connections []AVCaptureConnection, error_ foundation.INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:willFinishRecordingToOutputFileAtURL:fromConnections:error:"), output, fileURL, objectivec.IObjectSliceToNSArray(connections), error_)
	}
// Called whenever the output is recording to a file and successfully pauses
// the recording at the request of a client.
//
// output: The capture file output that has paused its file recording.
//
// fileURL: The file URL of the file that is being written.
//
// connections: An array of [AVCaptureConnection] objects attached to the file output that
// provided the data that is being written to the file.
//
// # Discussion
// 
// This method is called whenever a request to pause recording is actually
// respected.
// 
// It is safe for delegates to change what the file output is currently doing
// (starting a new file, for example) from within this method. If recording to
// a file is stopped, either manually or due to an error, this method is not
// guaranteed to be called, even if a previous call to `pauseRecording` was
// made.
// 
// You should not assume that this method will be called on a specific thread,
// and should make this method as efficient as possible.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutputRecordingDelegate/fileOutput(_:didPauseRecordingTo:from:)
func (o AVCaptureFileOutputRecordingDelegateObject) CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections(output IAVCaptureFileOutput, fileURL foundation.INSURL, connections []AVCaptureConnection) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:didPauseRecordingToOutputFileAtURL:fromConnections:"), output, fileURL, objectivec.IObjectSliceToNSArray(connections))
	}
// Called whenever the output, at the request of the client, successfully
// resumes a file recording that was paused.
//
// output: The capture file output that has resumed its paused file recording.
//
// fileURL: The file URL of the file that is being written.
//
// connections: An array of [AVCaptureConnection] objects attached to the file output that
// provided the data that is being written to the file.
//
// # Discussion
// 
// Delegates can use this method to be informed when a request to resume
// recording is actually respected.
// 
// It is safe for delegates to change what the file output is currently doing
// (starting a new file, for example) from within this method. If recording to
// a file is stopped, either manually or due to an error, this method is not
// guaranteed to be called, even if a previous call to [ResumeRecording] was
// made.
// 
// You should not assume that this method will be called on a specific thread,
// and should make this method as efficient as possible.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutputRecordingDelegate/fileOutput(_:didResumeRecordingTo:from:)
func (o AVCaptureFileOutputRecordingDelegateObject) CaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections(output IAVCaptureFileOutput, fileURL foundation.INSURL, connections []AVCaptureConnection) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:didResumeRecordingToOutputFileAtURL:fromConnections:"), output, fileURL, objectivec.IObjectSliceToNSArray(connections))
	}


// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVCaptureAudioPreviewOutput] class.
var (
	_AVCaptureAudioPreviewOutputClass     AVCaptureAudioPreviewOutputClass
	_AVCaptureAudioPreviewOutputClassOnce sync.Once
)

func getAVCaptureAudioPreviewOutputClass() AVCaptureAudioPreviewOutputClass {
	_AVCaptureAudioPreviewOutputClassOnce.Do(func() {
		_AVCaptureAudioPreviewOutputClass = AVCaptureAudioPreviewOutputClass{class: objc.GetClass("AVCaptureAudioPreviewOutput")}
	})
	return _AVCaptureAudioPreviewOutputClass
}

// GetAVCaptureAudioPreviewOutputClass returns the class object for AVCaptureAudioPreviewOutput.
func GetAVCaptureAudioPreviewOutputClass() AVCaptureAudioPreviewOutputClass {
	return getAVCaptureAudioPreviewOutputClass()
}

type AVCaptureAudioPreviewOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureAudioPreviewOutputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureAudioPreviewOutputClass) Alloc() AVCaptureAudioPreviewOutput {
	rv := objc.Send[AVCaptureAudioPreviewOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A capture output that provides a preview of the captured audio.
//
// # Configuring the output
//
//   - [AVCaptureAudioPreviewOutput.Volume]: The output volume of the audio preview.
//   - [AVCaptureAudioPreviewOutput.SetVolume]
//   - [AVCaptureAudioPreviewOutput.OutputDeviceUniqueID]: The unique identifier of the Core Audio output device to use for audio preview.
//   - [AVCaptureAudioPreviewOutput.SetOutputDeviceUniqueID]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioPreviewOutput
type AVCaptureAudioPreviewOutput struct {
	AVCaptureOutput
}

// AVCaptureAudioPreviewOutputFromID constructs a [AVCaptureAudioPreviewOutput] from an objc.ID.
//
// A capture output that provides a preview of the captured audio.
func AVCaptureAudioPreviewOutputFromID(id objc.ID) AVCaptureAudioPreviewOutput {
	return AVCaptureAudioPreviewOutput{AVCaptureOutput: AVCaptureOutputFromID(id)}
}

// NOTE: AVCaptureAudioPreviewOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureAudioPreviewOutput] class.
//
// # Configuring the output
//
//   - [IAVCaptureAudioPreviewOutput.Volume]: The output volume of the audio preview.
//   - [IAVCaptureAudioPreviewOutput.SetVolume]
//   - [IAVCaptureAudioPreviewOutput.OutputDeviceUniqueID]: The unique identifier of the Core Audio output device to use for audio preview.
//   - [IAVCaptureAudioPreviewOutput.SetOutputDeviceUniqueID]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioPreviewOutput
type IAVCaptureAudioPreviewOutput interface {
	IAVCaptureOutput

	// Topic: Configuring the output

	// The output volume of the audio preview.
	Volume() float32
	SetVolume(value float32)
	// The unique identifier of the Core Audio output device to use for audio preview.
	OutputDeviceUniqueID() string
	SetOutputDeviceUniqueID(value string)
}

// Init initializes the instance.
func (c AVCaptureAudioPreviewOutput) Init() AVCaptureAudioPreviewOutput {
	rv := objc.Send[AVCaptureAudioPreviewOutput](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureAudioPreviewOutput) Autorelease() AVCaptureAudioPreviewOutput {
	rv := objc.Send[AVCaptureAudioPreviewOutput](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureAudioPreviewOutput creates a new AVCaptureAudioPreviewOutput instance.
func NewAVCaptureAudioPreviewOutput() AVCaptureAudioPreviewOutput {
	class := getAVCaptureAudioPreviewOutputClass()
	rv := objc.Send[AVCaptureAudioPreviewOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The output volume of the audio preview.
//
// # Discussion
//
// A value of `1.0` indicates maximum volume, and a value of `0.0` mutes the
// audio preview.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioPreviewOutput/volume
func (c AVCaptureAudioPreviewOutput) Volume() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("volume"))
	return rv
}
func (c AVCaptureAudioPreviewOutput) SetVolume(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setVolume:"), value)
}

// The unique identifier of the Core Audio output device to use for audio
// preview.
//
// # Discussion
//
// Set the value to the unique identifier of the audio output device, or `nil`
// to use default system output.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioPreviewOutput/outputDeviceUniqueID
func (c AVCaptureAudioPreviewOutput) OutputDeviceUniqueID() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("outputDeviceUniqueID"))
	return foundation.NSStringFromID(rv).String()
}
func (c AVCaptureAudioPreviewOutput) SetOutputDeviceUniqueID(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setOutputDeviceUniqueID:"), objc.String(value))
}

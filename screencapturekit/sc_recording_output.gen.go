// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SCRecordingOutput] class.
var (
	_SCRecordingOutputClass     SCRecordingOutputClass
	_SCRecordingOutputClassOnce sync.Once
)

func getSCRecordingOutputClass() SCRecordingOutputClass {
	_SCRecordingOutputClassOnce.Do(func() {
		_SCRecordingOutputClass = SCRecordingOutputClass{class: objc.GetClass("SCRecordingOutput")}
	})
	return _SCRecordingOutputClass
}

// GetSCRecordingOutputClass returns the class object for SCRecordingOutput.
func GetSCRecordingOutputClass() SCRecordingOutputClass {
	return getSCRecordingOutputClass()
}

type SCRecordingOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SCRecordingOutputClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SCRecordingOutputClass) Alloc() SCRecordingOutput {
	rv := objc.Send[SCRecordingOutput](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Creating a recording output
//
//   - [SCRecordingOutput.InitWithConfigurationDelegate]
//
// # Configuring the recording output
//
//   - [SCRecordingOutput.RecordedDuration]
//   - [SCRecordingOutput.RecordedFileSize]
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutput
type SCRecordingOutput struct {
	objectivec.Object
}

// SCRecordingOutputFromID constructs a [SCRecordingOutput] from an objc.ID.
func SCRecordingOutputFromID(id objc.ID) SCRecordingOutput {
	return SCRecordingOutput{objectivec.Object{ID: id}}
}

// NOTE: SCRecordingOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SCRecordingOutput] class.
//
// # Creating a recording output
//
//   - [ISCRecordingOutput.InitWithConfigurationDelegate]
//
// # Configuring the recording output
//
//   - [ISCRecordingOutput.RecordedDuration]
//   - [ISCRecordingOutput.RecordedFileSize]
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutput
type ISCRecordingOutput interface {
	objectivec.IObject

	// Topic: Creating a recording output

	InitWithConfigurationDelegate(recordingOutputConfiguration ISCRecordingOutputConfiguration, delegate SCRecordingOutputDelegate) SCRecordingOutput

	// Topic: Configuring the recording output

	RecordedDuration() coremedia.CMTime
	RecordedFileSize() int
}

// Init initializes the instance.
func (r SCRecordingOutput) Init() SCRecordingOutput {
	rv := objc.Send[SCRecordingOutput](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r SCRecordingOutput) Autorelease() SCRecordingOutput {
	rv := objc.Send[SCRecordingOutput](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCRecordingOutput creates a new SCRecordingOutput instance.
func NewSCRecordingOutput() SCRecordingOutput {
	class := getSCRecordingOutputClass()
	rv := objc.Send[SCRecordingOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutput/init(configuration:delegate:)
func NewRecordingOutputWithConfigurationDelegate(recordingOutputConfiguration ISCRecordingOutputConfiguration, delegate SCRecordingOutputDelegate) SCRecordingOutput {
	instance := getSCRecordingOutputClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:delegate:"), recordingOutputConfiguration, delegate)
	return SCRecordingOutputFromID(rv)
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutput/init(configuration:delegate:)
func (r SCRecordingOutput) InitWithConfigurationDelegate(recordingOutputConfiguration ISCRecordingOutputConfiguration, delegate SCRecordingOutputDelegate) SCRecordingOutput {
	rv := objc.Send[SCRecordingOutput](r.ID, objc.Sel("initWithConfiguration:delegate:"), recordingOutputConfiguration, delegate)
	return rv
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutput/recordedDuration
func (r SCRecordingOutput) RecordedDuration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](r.ID, objc.Sel("recordedDuration"))
	return coremedia.CMTime(rv)
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutput/recordedFileSize
func (r SCRecordingOutput) RecordedFileSize() int {
	rv := objc.Send[int](r.ID, objc.Sel("recordedFileSize"))
	return rv
}

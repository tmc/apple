// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SCRecordingOutputConfiguration] class.
var (
	_SCRecordingOutputConfigurationClass     SCRecordingOutputConfigurationClass
	_SCRecordingOutputConfigurationClassOnce sync.Once
)

func getSCRecordingOutputConfigurationClass() SCRecordingOutputConfigurationClass {
	_SCRecordingOutputConfigurationClassOnce.Do(func() {
		_SCRecordingOutputConfigurationClass = SCRecordingOutputConfigurationClass{class: objc.GetClass("SCRecordingOutputConfiguration")}
	})
	return _SCRecordingOutputConfigurationClass
}

// GetSCRecordingOutputConfigurationClass returns the class object for SCRecordingOutputConfiguration.
func GetSCRecordingOutputConfigurationClass() SCRecordingOutputConfigurationClass {
	return getSCRecordingOutputConfigurationClass()
}

type SCRecordingOutputConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (sc SCRecordingOutputConfigurationClass) Alloc() SCRecordingOutputConfiguration {
	rv := objc.Send[SCRecordingOutputConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

//
// # Instance Properties
//
//   - [SCRecordingOutputConfiguration.AvailableOutputFileTypes]
//   - [SCRecordingOutputConfiguration.AvailableVideoCodecTypes]
//   - [SCRecordingOutputConfiguration.OutputFileType]
//   - [SCRecordingOutputConfiguration.SetOutputFileType]
//   - [SCRecordingOutputConfiguration.OutputURL]
//   - [SCRecordingOutputConfiguration.SetOutputURL]
//   - [SCRecordingOutputConfiguration.VideoCodecType]
//   - [SCRecordingOutputConfiguration.SetVideoCodecType]
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputConfiguration
type SCRecordingOutputConfiguration struct {
	objectivec.Object
}

// SCRecordingOutputConfigurationFromID constructs a [SCRecordingOutputConfiguration] from an objc.ID.
func SCRecordingOutputConfigurationFromID(id objc.ID) SCRecordingOutputConfiguration {
	return SCRecordingOutputConfiguration{objectivec.Object{ID: id}}
}
// NOTE: SCRecordingOutputConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SCRecordingOutputConfiguration] class.
//
// # Instance Properties
//
//   - [ISCRecordingOutputConfiguration.AvailableOutputFileTypes]
//   - [ISCRecordingOutputConfiguration.AvailableVideoCodecTypes]
//   - [ISCRecordingOutputConfiguration.OutputFileType]
//   - [ISCRecordingOutputConfiguration.SetOutputFileType]
//   - [ISCRecordingOutputConfiguration.OutputURL]
//   - [ISCRecordingOutputConfiguration.SetOutputURL]
//   - [ISCRecordingOutputConfiguration.VideoCodecType]
//   - [ISCRecordingOutputConfiguration.SetVideoCodecType]
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputConfiguration
type ISCRecordingOutputConfiguration interface {
	objectivec.IObject

	// Topic: Instance Properties

	AvailableOutputFileTypes() []string
	AvailableVideoCodecTypes() []string
	OutputFileType() foundation.NSString
	SetOutputFileType(value foundation.NSString)
	OutputURL() foundation.INSURL
	SetOutputURL(value foundation.INSURL)
	VideoCodecType() foundation.NSString
	SetVideoCodecType(value foundation.NSString)
}

// Init initializes the instance.
func (r SCRecordingOutputConfiguration) Init() SCRecordingOutputConfiguration {
	rv := objc.Send[SCRecordingOutputConfiguration](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r SCRecordingOutputConfiguration) Autorelease() SCRecordingOutputConfiguration {
	rv := objc.Send[SCRecordingOutputConfiguration](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCRecordingOutputConfiguration creates a new SCRecordingOutputConfiguration instance.
func NewSCRecordingOutputConfiguration() SCRecordingOutputConfiguration {
	class := getSCRecordingOutputConfigurationClass()
	rv := objc.Send[SCRecordingOutputConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputConfiguration/availableOutputFileTypes
func (r SCRecordingOutputConfiguration) AvailableOutputFileTypes() []string {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("availableOutputFileTypes"))
	return objc.ConvertSliceToStrings(rv)
}
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputConfiguration/availableVideoCodecTypes
func (r SCRecordingOutputConfiguration) AvailableVideoCodecTypes() []string {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("availableVideoCodecTypes"))
	return objc.ConvertSliceToStrings(rv)
}
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputConfiguration/outputFileType
func (r SCRecordingOutputConfiguration) OutputFileType() foundation.NSString {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("outputFileType"))
	return foundation.NSStringFromID(objc.ID(rv))
}
func (r SCRecordingOutputConfiguration) SetOutputFileType(value foundation.NSString) {
	objc.Send[struct{}](r.ID, objc.Sel("setOutputFileType:"), value)
}
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputConfiguration/outputURL
func (r SCRecordingOutputConfiguration) OutputURL() foundation.INSURL {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("outputURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (r SCRecordingOutputConfiguration) SetOutputURL(value foundation.INSURL) {
	objc.Send[struct{}](r.ID, objc.Sel("setOutputURL:"), value)
}
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputConfiguration/videoCodecType
func (r SCRecordingOutputConfiguration) VideoCodecType() foundation.NSString {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("videoCodecType"))
	return foundation.NSStringFromID(objc.ID(rv))
}
func (r SCRecordingOutputConfiguration) SetVideoCodecType(value foundation.NSString) {
	objc.Send[struct{}](r.ID, objc.Sel("setVideoCodecType:"), value)
}


// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioRecorder] class.
var (
	_AVAudioRecorderClass     AVAudioRecorderClass
	_AVAudioRecorderClassOnce sync.Once
)

func getAVAudioRecorderClass() AVAudioRecorderClass {
	_AVAudioRecorderClassOnce.Do(func() {
		_AVAudioRecorderClass = AVAudioRecorderClass{class: objc.GetClass("AVAudioRecorder")}
	})
	return _AVAudioRecorderClass
}

// GetAVAudioRecorderClass returns the class object for AVAudioRecorder.
func GetAVAudioRecorderClass() AVAudioRecorderClass {
	return getAVAudioRecorderClass()
}

type AVAudioRecorderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioRecorderClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioRecorderClass) Alloc() AVAudioRecorder {
	rv := objc.Send[AVAudioRecorder](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioRecorder.BaseInit]
//   - [AVAudioRecorder.FinishedRecording]
//   - [AVAudioRecorder.InstantaneousMetering]
//   - [AVAudioRecorder.PrivCommonCleanup]
//   - [AVAudioRecorder.PrivRemoveSessionPropertyListeners]
//   - [AVAudioRecorder.PrivSetDelegate]
//   - [AVAudioRecorder.Record]
//   - [AVAudioRecorder.SetInstantaneousMetering]
//   - [AVAudioRecorder.Url]
//   - [AVAudioRecorder.MeteringEnabled]
//   - [AVAudioRecorder.SetMeteringEnabled]
//   - [AVAudioRecorder.Recording]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder
type AVAudioRecorder struct {
	objectivec.Object
}

// AVAudioRecorderFromID constructs a [AVAudioRecorder] from an objc.ID.
func AVAudioRecorderFromID(id objc.ID) AVAudioRecorder {
	return AVAudioRecorder{objectivec.Object{ID: id}}
}

// Ensure AVAudioRecorder implements IAVAudioRecorder.
var _ IAVAudioRecorder = AVAudioRecorder{}

// An interface definition for the [AVAudioRecorder] class.
//
// # Methods
//
//   - [IAVAudioRecorder.BaseInit]
//   - [IAVAudioRecorder.FinishedRecording]
//   - [IAVAudioRecorder.InstantaneousMetering]
//   - [IAVAudioRecorder.PrivCommonCleanup]
//   - [IAVAudioRecorder.PrivRemoveSessionPropertyListeners]
//   - [IAVAudioRecorder.PrivSetDelegate]
//   - [IAVAudioRecorder.Record]
//   - [IAVAudioRecorder.SetInstantaneousMetering]
//   - [IAVAudioRecorder.Url]
//   - [IAVAudioRecorder.MeteringEnabled]
//   - [IAVAudioRecorder.SetMeteringEnabled]
//   - [IAVAudioRecorder.Recording]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder
type IAVAudioRecorder interface {
	objectivec.IObject

	// Topic: Methods

	BaseInit() objectivec.IObject
	FinishedRecording(recording objectivec.IObject)
	InstantaneousMetering() bool
	PrivCommonCleanup()
	PrivRemoveSessionPropertyListeners()
	PrivSetDelegate(delegate objectivec.IObject)
	Record() bool
	SetInstantaneousMetering(metering bool)
	Url() foundation.INSURL
	MeteringEnabled() bool
	SetMeteringEnabled(value bool)
	Recording() bool
}

// Init initializes the instance.
func (a AVAudioRecorder) Init() AVAudioRecorder {
	rv := objc.Send[AVAudioRecorder](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioRecorder) Autorelease() AVAudioRecorder {
	rv := objc.Send[AVAudioRecorder](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioRecorder creates a new AVAudioRecorder instance.
func NewAVAudioRecorder() AVAudioRecorder {
	class := getAVAudioRecorderClass()
	rv := objc.Send[AVAudioRecorder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/baseInit
func (a AVAudioRecorder) BaseInit() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("baseInit"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/finishedRecording:
func (a AVAudioRecorder) FinishedRecording(recording objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("finishedRecording:"), recording)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/instantaneousMetering
func (a AVAudioRecorder) InstantaneousMetering() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("instantaneousMetering"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/privCommonCleanup
func (a AVAudioRecorder) PrivCommonCleanup() {
	objc.Send[objc.ID](a.ID, objc.Sel("privCommonCleanup"))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/privRemoveSessionPropertyListeners
func (a AVAudioRecorder) PrivRemoveSessionPropertyListeners() {
	objc.Send[objc.ID](a.ID, objc.Sel("privRemoveSessionPropertyListeners"))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/privSetDelegate:
func (a AVAudioRecorder) PrivSetDelegate(delegate objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("privSetDelegate:"), delegate)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/record
func (a AVAudioRecorder) Record() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("record"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/setInstantaneousMetering:
func (a AVAudioRecorder) SetInstantaneousMetering(metering bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setInstantaneousMetering:"), metering)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/meteringEnabled
func (a AVAudioRecorder) MeteringEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("meteringEnabled"))
	return rv
}
func (a AVAudioRecorder) SetMeteringEnabled(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setMeteringEnabled:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/recording
func (a AVAudioRecorder) Recording() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("recording"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/url
func (a AVAudioRecorder) Url() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

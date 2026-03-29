// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMusicTrack] class.
var (
	_AVMusicTrackClass     AVMusicTrackClass
	_AVMusicTrackClassOnce sync.Once
)

func getAVMusicTrackClass() AVMusicTrackClass {
	_AVMusicTrackClassOnce.Do(func() {
		_AVMusicTrackClass = AVMusicTrackClass{class: objc.GetClass("AVMusicTrack")}
	})
	return _AVMusicTrackClass
}

// GetAVMusicTrackClass returns the class object for AVMusicTrack.
func GetAVMusicTrackClass() AVMusicTrackClass {
	return getAVMusicTrackClass()
}

type AVMusicTrackClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMusicTrackClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMusicTrackClass) Alloc() AVMusicTrack {
	rv := objc.Send[AVMusicTrack](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVMusicTrack.CreateEventIterator]
//   - [AVMusicTrack.DoAddAUPresetEventAtBeat]
//   - [AVMusicTrack.DoAddExtendedNoteOnEventAtBeat]
//   - [AVMusicTrack.DoAddExtendedTempoEventAtBeat]
//   - [AVMusicTrack.DoAddMIDIChannelEventAtBeat]
//   - [AVMusicTrack.DoAddMIDIMetaEventAtBeat]
//   - [AVMusicTrack.DoAddMIDINoteEventAtBeat]
//   - [AVMusicTrack.DoAddMIDISysexEventAtBeat]
//   - [AVMusicTrack.DoAddParameterEventAtBeat]
//   - [AVMusicTrack.DoAddUserEventAtBeat]
//   - [AVMusicTrack.LoopingEnabled]
//   - [AVMusicTrack.SetLoopingEnabled]
//   - [AVMusicTrack.Muted]
//   - [AVMusicTrack.SetMuted]
//   - [AVMusicTrack.Soloed]
//   - [AVMusicTrack.SetSoloed]
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack
type AVMusicTrack struct {
	objectivec.Object
}

// AVMusicTrackFromID constructs a [AVMusicTrack] from an objc.ID.
func AVMusicTrackFromID(id objc.ID) AVMusicTrack {
	return AVMusicTrack{objectivec.Object{ID: id}}
}
// Ensure AVMusicTrack implements IAVMusicTrack.
var _ IAVMusicTrack = AVMusicTrack{}

// An interface definition for the [AVMusicTrack] class.
//
// # Methods
//
//   - [IAVMusicTrack.CreateEventIterator]
//   - [IAVMusicTrack.DoAddAUPresetEventAtBeat]
//   - [IAVMusicTrack.DoAddExtendedNoteOnEventAtBeat]
//   - [IAVMusicTrack.DoAddExtendedTempoEventAtBeat]
//   - [IAVMusicTrack.DoAddMIDIChannelEventAtBeat]
//   - [IAVMusicTrack.DoAddMIDIMetaEventAtBeat]
//   - [IAVMusicTrack.DoAddMIDINoteEventAtBeat]
//   - [IAVMusicTrack.DoAddMIDISysexEventAtBeat]
//   - [IAVMusicTrack.DoAddParameterEventAtBeat]
//   - [IAVMusicTrack.DoAddUserEventAtBeat]
//   - [IAVMusicTrack.LoopingEnabled]
//   - [IAVMusicTrack.SetLoopingEnabled]
//   - [IAVMusicTrack.Muted]
//   - [IAVMusicTrack.SetMuted]
//   - [IAVMusicTrack.Soloed]
//   - [IAVMusicTrack.SetSoloed]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack
type IAVMusicTrack interface {
	objectivec.IObject

	// Topic: Methods

	CreateEventIterator() objectivec.IObject
	DoAddAUPresetEventAtBeat(event objectivec.IObject, beat float64)
	DoAddExtendedNoteOnEventAtBeat(event objectivec.IObject, beat float64)
	DoAddExtendedTempoEventAtBeat(event objectivec.IObject, beat float64)
	DoAddMIDIChannelEventAtBeat(event objectivec.IObject, beat float64)
	DoAddMIDIMetaEventAtBeat(event objectivec.IObject, beat float64)
	DoAddMIDINoteEventAtBeat(event objectivec.IObject, beat float64)
	DoAddMIDISysexEventAtBeat(event objectivec.IObject, beat float64)
	DoAddParameterEventAtBeat(event objectivec.IObject, beat float64)
	DoAddUserEventAtBeat(event objectivec.IObject, beat float64)
	LoopingEnabled() bool
	SetLoopingEnabled(value bool)
	Muted() bool
	SetMuted(value bool)
	Soloed() bool
	SetSoloed(value bool)
}

// Init initializes the instance.
func (m AVMusicTrack) Init() AVMusicTrack {
	rv := objc.Send[AVMusicTrack](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMusicTrack) Autorelease() AVMusicTrack {
	rv := objc.Send[AVMusicTrack](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMusicTrack creates a new AVMusicTrack instance.
func NewAVMusicTrack() AVMusicTrack {
	class := getAVMusicTrackClass()
	rv := objc.Send[AVMusicTrack](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/createEventIterator
func (m AVMusicTrack) CreateEventIterator() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("createEventIterator"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/doAddAUPresetEvent:atBeat:
func (m AVMusicTrack) DoAddAUPresetEventAtBeat(event objectivec.IObject, beat float64) {
	objc.Send[objc.ID](m.ID, objc.Sel("doAddAUPresetEvent:atBeat:"), event, beat)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/doAddExtendedNoteOnEvent:atBeat:
func (m AVMusicTrack) DoAddExtendedNoteOnEventAtBeat(event objectivec.IObject, beat float64) {
	objc.Send[objc.ID](m.ID, objc.Sel("doAddExtendedNoteOnEvent:atBeat:"), event, beat)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/doAddExtendedTempoEvent:atBeat:
func (m AVMusicTrack) DoAddExtendedTempoEventAtBeat(event objectivec.IObject, beat float64) {
	objc.Send[objc.ID](m.ID, objc.Sel("doAddExtendedTempoEvent:atBeat:"), event, beat)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/doAddMIDIChannelEvent:atBeat:
func (m AVMusicTrack) DoAddMIDIChannelEventAtBeat(event objectivec.IObject, beat float64) {
	objc.Send[objc.ID](m.ID, objc.Sel("doAddMIDIChannelEvent:atBeat:"), event, beat)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/doAddMIDIMetaEvent:atBeat:
func (m AVMusicTrack) DoAddMIDIMetaEventAtBeat(event objectivec.IObject, beat float64) {
	objc.Send[objc.ID](m.ID, objc.Sel("doAddMIDIMetaEvent:atBeat:"), event, beat)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/doAddMIDINoteEvent:atBeat:
func (m AVMusicTrack) DoAddMIDINoteEventAtBeat(event objectivec.IObject, beat float64) {
	objc.Send[objc.ID](m.ID, objc.Sel("doAddMIDINoteEvent:atBeat:"), event, beat)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/doAddMIDISysexEvent:atBeat:
func (m AVMusicTrack) DoAddMIDISysexEventAtBeat(event objectivec.IObject, beat float64) {
	objc.Send[objc.ID](m.ID, objc.Sel("doAddMIDISysexEvent:atBeat:"), event, beat)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/doAddParameterEvent:atBeat:
func (m AVMusicTrack) DoAddParameterEventAtBeat(event objectivec.IObject, beat float64) {
	objc.Send[objc.ID](m.ID, objc.Sel("doAddParameterEvent:atBeat:"), event, beat)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/doAddUserEvent:atBeat:
func (m AVMusicTrack) DoAddUserEventAtBeat(event objectivec.IObject, beat float64) {
	objc.Send[objc.ID](m.ID, objc.Sel("doAddUserEvent:atBeat:"), event, beat)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/loopingEnabled
func (m AVMusicTrack) LoopingEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("loopingEnabled"))
	return rv
}
func (m AVMusicTrack) SetLoopingEnabled(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setLoopingEnabled:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/muted
func (m AVMusicTrack) Muted() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("muted"))
	return rv
}
func (m AVMusicTrack) SetMuted(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setMuted:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/soloed
func (m AVMusicTrack) Soloed() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("soloed"))
	return rv
}
func (m AVMusicTrack) SetSoloed(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setSoloed:"), value)
}


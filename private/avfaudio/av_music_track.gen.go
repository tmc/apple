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

// # Methods
//
//   - [AVMusicTrack.CreateEventIterator]
//   - [AVMusicTrack.DestinationMIDIEndpoint]
//   - [AVMusicTrack.SetDestinationMIDIEndpoint]
//   - [AVMusicTrack.DoAddAUPresetEventAtBeat]
//   - [AVMusicTrack.DoAddExtendedNoteOnEventAtBeat]
//   - [AVMusicTrack.DoAddExtendedTempoEventAtBeat]
//   - [AVMusicTrack.DoAddMIDIChannelEventAtBeat]
//   - [AVMusicTrack.DoAddMIDIMetaEventAtBeat]
//   - [AVMusicTrack.DoAddMIDINoteEventAtBeat]
//   - [AVMusicTrack.DoAddMIDISysexEventAtBeat]
//   - [AVMusicTrack.DoAddParameterEventAtBeat]
//   - [AVMusicTrack.DoAddUserEventAtBeat]
//   - [AVMusicTrack.Index]
//   - [AVMusicTrack.Track]
//   - [AVMusicTrack.LoopingEnabled]
//   - [AVMusicTrack.SetLoopingEnabled]
//   - [AVMusicTrack.Muted]
//   - [AVMusicTrack.SetMuted]
//   - [AVMusicTrack.Soloed]
//   - [AVMusicTrack.SetSoloed]
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
//   - [IAVMusicTrack.DestinationMIDIEndpoint]
//   - [IAVMusicTrack.SetDestinationMIDIEndpoint]
//   - [IAVMusicTrack.DoAddAUPresetEventAtBeat]
//   - [IAVMusicTrack.DoAddExtendedNoteOnEventAtBeat]
//   - [IAVMusicTrack.DoAddExtendedTempoEventAtBeat]
//   - [IAVMusicTrack.DoAddMIDIChannelEventAtBeat]
//   - [IAVMusicTrack.DoAddMIDIMetaEventAtBeat]
//   - [IAVMusicTrack.DoAddMIDINoteEventAtBeat]
//   - [IAVMusicTrack.DoAddMIDISysexEventAtBeat]
//   - [IAVMusicTrack.DoAddParameterEventAtBeat]
//   - [IAVMusicTrack.DoAddUserEventAtBeat]
//   - [IAVMusicTrack.Index]
//   - [IAVMusicTrack.Track]
//   - [IAVMusicTrack.LoopingEnabled]
//   - [IAVMusicTrack.SetLoopingEnabled]
//   - [IAVMusicTrack.Muted]
//   - [IAVMusicTrack.SetMuted]
//   - [IAVMusicTrack.Soloed]
//   - [IAVMusicTrack.SetSoloed]
type IAVMusicTrack interface {
	objectivec.IObject

	// Topic: Methods

	CreateEventIterator() objectivec.IObject
	DestinationMIDIEndpoint() uint32
	SetDestinationMIDIEndpoint(value uint32)
	DoAddAUPresetEventAtBeat(event objectivec.IObject, beat float64)
	DoAddExtendedNoteOnEventAtBeat(event objectivec.IObject, beat float64)
	DoAddExtendedTempoEventAtBeat(event objectivec.IObject, beat float64)
	DoAddMIDIChannelEventAtBeat(event objectivec.IObject, beat float64)
	DoAddMIDIMetaEventAtBeat(event objectivec.IObject, beat float64)
	DoAddMIDINoteEventAtBeat(event objectivec.IObject, beat float64)
	DoAddMIDISysexEventAtBeat(event objectivec.IObject, beat float64)
	DoAddParameterEventAtBeat(event objectivec.IObject, beat float64)
	DoAddUserEventAtBeat(event objectivec.IObject, beat float64)
	Index() uint64
	Track() OpaqueMusicTrackRef
	LoopingEnabled() bool
	SetLoopingEnabled(value bool)
	Muted() bool
	SetMuted(value bool)
	Soloed() bool
	SetSoloed(value bool)
}

// Init initializes the instance.
func (a AVMusicTrack) Init() AVMusicTrack {
	rv := objc.Send[AVMusicTrack](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVMusicTrack) Autorelease() AVMusicTrack {
	rv := objc.Send[AVMusicTrack](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMusicTrack creates a new AVMusicTrack instance.
func NewAVMusicTrack() AVMusicTrack {
	class := getAVMusicTrackClass()
	rv := objc.Send[AVMusicTrack](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (a AVMusicTrack) CreateEventIterator() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("createEventIterator"))
	return objectivec.Object{ID: rv}
}
func (a AVMusicTrack) DoAddAUPresetEventAtBeat(event objectivec.IObject, beat float64) {
	objc.Send[objc.ID](a.ID, objc.Sel("doAddAUPresetEvent:atBeat:"), event, beat)
}
func (a AVMusicTrack) DoAddExtendedNoteOnEventAtBeat(event objectivec.IObject, beat float64) {
	objc.Send[objc.ID](a.ID, objc.Sel("doAddExtendedNoteOnEvent:atBeat:"), event, beat)
}
func (a AVMusicTrack) DoAddExtendedTempoEventAtBeat(event objectivec.IObject, beat float64) {
	objc.Send[objc.ID](a.ID, objc.Sel("doAddExtendedTempoEvent:atBeat:"), event, beat)
}
func (a AVMusicTrack) DoAddMIDIChannelEventAtBeat(event objectivec.IObject, beat float64) {
	objc.Send[objc.ID](a.ID, objc.Sel("doAddMIDIChannelEvent:atBeat:"), event, beat)
}
func (a AVMusicTrack) DoAddMIDIMetaEventAtBeat(event objectivec.IObject, beat float64) {
	objc.Send[objc.ID](a.ID, objc.Sel("doAddMIDIMetaEvent:atBeat:"), event, beat)
}
func (a AVMusicTrack) DoAddMIDINoteEventAtBeat(event objectivec.IObject, beat float64) {
	objc.Send[objc.ID](a.ID, objc.Sel("doAddMIDINoteEvent:atBeat:"), event, beat)
}
func (a AVMusicTrack) DoAddMIDISysexEventAtBeat(event objectivec.IObject, beat float64) {
	objc.Send[objc.ID](a.ID, objc.Sel("doAddMIDISysexEvent:atBeat:"), event, beat)
}
func (a AVMusicTrack) DoAddParameterEventAtBeat(event objectivec.IObject, beat float64) {
	objc.Send[objc.ID](a.ID, objc.Sel("doAddParameterEvent:atBeat:"), event, beat)
}
func (a AVMusicTrack) DoAddUserEventAtBeat(event objectivec.IObject, beat float64) {
	objc.Send[objc.ID](a.ID, objc.Sel("doAddUserEvent:atBeat:"), event, beat)
}
func (a AVMusicTrack) Index() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("index"))
	return rv
}
func (a AVMusicTrack) Track() OpaqueMusicTrackRef {
	rv := objc.Send[OpaqueMusicTrackRef](a.ID, objc.Sel("track"))
	return OpaqueMusicTrackRef(rv)
}

func (a AVMusicTrack) DestinationMIDIEndpoint() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("destinationMIDIEndpoint"))
	return rv
}
func (a AVMusicTrack) SetDestinationMIDIEndpoint(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setDestinationMIDIEndpoint:"), value)
}
func (a AVMusicTrack) LoopingEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("loopingEnabled"))
	return rv
}
func (a AVMusicTrack) SetLoopingEnabled(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setLoopingEnabled:"), value)
}
func (a AVMusicTrack) Muted() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("muted"))
	return rv
}
func (a AVMusicTrack) SetMuted(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setMuted:"), value)
}
func (a AVMusicTrack) Soloed() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("soloed"))
	return rv
}
func (a AVMusicTrack) SetSoloed(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setSoloed:"), value)
}

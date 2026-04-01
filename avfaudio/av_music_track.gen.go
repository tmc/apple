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

// A collection of music events that you can offset, set to a muted state,
// modify independently from other track events, and send to a specified
// destination.
//
// # Configuring Music Track Properties
//
//   - [AVMusicTrack.Muted]: A Boolean value that indicates whether the track is in a muted state.
//   - [AVMusicTrack.SetMuted]
//   - [AVMusicTrack.Soloed]: A Boolean value that indicates whether the track is in a soloed state.
//   - [AVMusicTrack.SetSoloed]
//   - [AVMusicTrack.OffsetTime]: The offset of the track’s start time, in beats.
//   - [AVMusicTrack.SetOffsetTime]
//   - [AVMusicTrack.TimeResolution]: The time resolution value for the sequence, in ticks (pulses) per quarter note.
//   - [AVMusicTrack.UsesAutomatedParameters]: A Boolean value that indicates whether the track is an automation track.
//   - [AVMusicTrack.SetUsesAutomatedParameters]
//
// # Configuring the Track Duration
//
//   - [AVMusicTrack.LengthInBeats]: The total duration of the track, in beats.
//   - [AVMusicTrack.SetLengthInBeats]
//   - [AVMusicTrack.LengthInSeconds]: The total duration of the track, in seconds.
//   - [AVMusicTrack.SetLengthInSeconds]
//
// # Configuring the Track Destinations
//
//   - [AVMusicTrack.DestinationAudioUnit]: The audio unit that receives the track’s events.
//   - [AVMusicTrack.SetDestinationAudioUnit]
//
// # Configuring the Looping State
//
//   - [AVMusicTrack.LoopingEnabled]: A Boolean value that indicates whether the track is in a looping state.
//   - [AVMusicTrack.SetLoopingEnabled]
//   - [AVMusicTrack.LoopRange]: The timestamp range for the loop, in beats.
//   - [AVMusicTrack.SetLoopRange]
//   - [AVMusicTrack.NumberOfLoops]: The number of times the track’s loop repeats.
//   - [AVMusicTrack.SetNumberOfLoops]
//
// # Adding and Clearing Events
//
//   - [AVMusicTrack.AddEventAtBeat]: Adds a music event to a track at the time you specify.
//   - [AVMusicTrack.MoveEventsInRangeByAmount]: Moves the beat location of all events in the given beat range by the amount you specify.
//   - [AVMusicTrack.ClearEventsInRange]: Removes all events in the given beat range from the music track.
//
// # Cutting and Copying Events
//
//   - [AVMusicTrack.CutEventsInRange]: Splices all events in the beat range from the music track.
//   - [AVMusicTrack.CopyEventsInRangeFromTrackInsertAtBeat]: Copies the events from the source track and splices them into the current music track.
//   - [AVMusicTrack.CopyAndMergeEventsInRangeFromTrackMergeAtBeat]: Copies the events from the source track and merges them into the current music track.
//
// # Iterating Over Events
//
//   - [AVMusicTrack.EnumerateEventsInRangeUsingBlock]: Iterates through the music events within the track.
//
// # Getting the End of Track Timestamp
//
//   - [AVMusicTrack.AVMusicTimeStampEndOfTrack]: A timestamp you use to access all events in a music track through a beat range.
//   - [AVMusicTrack.SetAVMusicTimeStampEndOfTrack]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack
type AVMusicTrack struct {
	objectivec.Object
}

// AVMusicTrackFromID constructs a [AVMusicTrack] from an objc.ID.
//
// A collection of music events that you can offset, set to a muted state,
// modify independently from other track events, and send to a specified
// destination.
func AVMusicTrackFromID(id objc.ID) AVMusicTrack {
	return AVMusicTrack{objectivec.Object{ID: id}}
}

// NOTE: AVMusicTrack adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMusicTrack] class.
//
// # Configuring Music Track Properties
//
//   - [IAVMusicTrack.Muted]: A Boolean value that indicates whether the track is in a muted state.
//   - [IAVMusicTrack.SetMuted]
//   - [IAVMusicTrack.Soloed]: A Boolean value that indicates whether the track is in a soloed state.
//   - [IAVMusicTrack.SetSoloed]
//   - [IAVMusicTrack.OffsetTime]: The offset of the track’s start time, in beats.
//   - [IAVMusicTrack.SetOffsetTime]
//   - [IAVMusicTrack.TimeResolution]: The time resolution value for the sequence, in ticks (pulses) per quarter note.
//   - [IAVMusicTrack.UsesAutomatedParameters]: A Boolean value that indicates whether the track is an automation track.
//   - [IAVMusicTrack.SetUsesAutomatedParameters]
//
// # Configuring the Track Duration
//
//   - [IAVMusicTrack.LengthInBeats]: The total duration of the track, in beats.
//   - [IAVMusicTrack.SetLengthInBeats]
//   - [IAVMusicTrack.LengthInSeconds]: The total duration of the track, in seconds.
//   - [IAVMusicTrack.SetLengthInSeconds]
//
// # Configuring the Track Destinations
//
//   - [IAVMusicTrack.DestinationAudioUnit]: The audio unit that receives the track’s events.
//   - [IAVMusicTrack.SetDestinationAudioUnit]
//
// # Configuring the Looping State
//
//   - [IAVMusicTrack.LoopingEnabled]: A Boolean value that indicates whether the track is in a looping state.
//   - [IAVMusicTrack.SetLoopingEnabled]
//   - [IAVMusicTrack.LoopRange]: The timestamp range for the loop, in beats.
//   - [IAVMusicTrack.SetLoopRange]
//   - [IAVMusicTrack.NumberOfLoops]: The number of times the track’s loop repeats.
//   - [IAVMusicTrack.SetNumberOfLoops]
//
// # Adding and Clearing Events
//
//   - [IAVMusicTrack.AddEventAtBeat]: Adds a music event to a track at the time you specify.
//   - [IAVMusicTrack.MoveEventsInRangeByAmount]: Moves the beat location of all events in the given beat range by the amount you specify.
//   - [IAVMusicTrack.ClearEventsInRange]: Removes all events in the given beat range from the music track.
//
// # Cutting and Copying Events
//
//   - [IAVMusicTrack.CutEventsInRange]: Splices all events in the beat range from the music track.
//   - [IAVMusicTrack.CopyEventsInRangeFromTrackInsertAtBeat]: Copies the events from the source track and splices them into the current music track.
//   - [IAVMusicTrack.CopyAndMergeEventsInRangeFromTrackMergeAtBeat]: Copies the events from the source track and merges them into the current music track.
//
// # Iterating Over Events
//
//   - [IAVMusicTrack.EnumerateEventsInRangeUsingBlock]: Iterates through the music events within the track.
//
// # Getting the End of Track Timestamp
//
//   - [IAVMusicTrack.AVMusicTimeStampEndOfTrack]: A timestamp you use to access all events in a music track through a beat range.
//   - [IAVMusicTrack.SetAVMusicTimeStampEndOfTrack]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack
type IAVMusicTrack interface {
	objectivec.IObject

	// Topic: Configuring Music Track Properties

	// A Boolean value that indicates whether the track is in a muted state.
	Muted() bool
	SetMuted(value bool)
	// A Boolean value that indicates whether the track is in a soloed state.
	Soloed() bool
	SetSoloed(value bool)
	// The offset of the track’s start time, in beats.
	OffsetTime() AVMusicTimeStamp
	SetOffsetTime(value AVMusicTimeStamp)
	// The time resolution value for the sequence, in ticks (pulses) per quarter note.
	TimeResolution() uint
	// A Boolean value that indicates whether the track is an automation track.
	UsesAutomatedParameters() bool
	SetUsesAutomatedParameters(value bool)

	// Topic: Configuring the Track Duration

	// The total duration of the track, in beats.
	LengthInBeats() AVMusicTimeStamp
	SetLengthInBeats(value AVMusicTimeStamp)
	// The total duration of the track, in seconds.
	LengthInSeconds() float64
	SetLengthInSeconds(value float64)

	// Topic: Configuring the Track Destinations

	// The audio unit that receives the track’s events.
	DestinationAudioUnit() IAVAudioUnit
	SetDestinationAudioUnit(value IAVAudioUnit)

	// Topic: Configuring the Looping State

	// A Boolean value that indicates whether the track is in a looping state.
	LoopingEnabled() bool
	SetLoopingEnabled(value bool)
	// The timestamp range for the loop, in beats.
	LoopRange() AVBeatRange
	SetLoopRange(value AVBeatRange)
	// The number of times the track’s loop repeats.
	NumberOfLoops() int
	SetNumberOfLoops(value int)

	// Topic: Adding and Clearing Events

	// Adds a music event to a track at the time you specify.
	AddEventAtBeat(event IAVMusicEvent, beat AVMusicTimeStamp)
	// Moves the beat location of all events in the given beat range by the amount you specify.
	MoveEventsInRangeByAmount(range_ AVBeatRange, beatAmount AVMusicTimeStamp)
	// Removes all events in the given beat range from the music track.
	ClearEventsInRange(range_ AVBeatRange)

	// Topic: Cutting and Copying Events

	// Splices all events in the beat range from the music track.
	CutEventsInRange(range_ AVBeatRange)
	// Copies the events from the source track and splices them into the current music track.
	CopyEventsInRangeFromTrackInsertAtBeat(range_ AVBeatRange, sourceTrack IAVMusicTrack, insertStartBeat AVMusicTimeStamp)
	// Copies the events from the source track and merges them into the current music track.
	CopyAndMergeEventsInRangeFromTrackMergeAtBeat(range_ AVBeatRange, sourceTrack IAVMusicTrack, mergeStartBeat AVMusicTimeStamp)

	// Topic: Iterating Over Events

	// Iterates through the music events within the track.
	EnumerateEventsInRangeUsingBlock(range_ AVBeatRange, block AVMusicEventEnumerationBlock)

	// Topic: Getting the End of Track Timestamp

	// A timestamp you use to access all events in a music track through a beat range.
	AVMusicTimeStampEndOfTrack() float64
	SetAVMusicTimeStampEndOfTrack(value float64)
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

// Adds a music event to a track at the time you specify.
//
// event: The event to add.
//
// beat: The time to add the event at.
//
// # Discussion
//
// The system copies event contents into the track, so you can add the same
// event at different timestamps. You can’t add all [AVMusicEvent]
// subclasses to a track.
//
// - You can only add [AVExtendedTempoEvent] and [AVMIDIMetaEvent] with
// certain [AVMIDIMetaEvent.EventType] to a sequencer’s tempo track. - You
// can add [AVParameterEvent] to automation tracks. - You can’t add other
// event subclasses to tempo or automation tracks.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/addEvent(_:at:)
//
// [AVMIDIMetaEvent.EventType]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/EventType
func (m AVMusicTrack) AddEventAtBeat(event IAVMusicEvent, beat AVMusicTimeStamp) {
	objc.Send[objc.ID](m.ID, objc.Sel("addEvent:atBeat:"), event, beat)
}

// Moves the beat location of all events in the given beat range by the amount
// you specify.
//
// range: The range of beats.
//
// beatAmount: The amount of beats to shift each event.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/moveEvents(in:by:)
func (m AVMusicTrack) MoveEventsInRangeByAmount(range_ AVBeatRange, beatAmount AVMusicTimeStamp) {
	objc.Send[objc.ID](m.ID, objc.Sel("moveEventsInRange:byAmount:"), range_, beatAmount)
}

// Removes all events in the given beat range from the music track.
//
// range: The range of beats.
//
// # Discussion
//
// The system won’t modify the events outside of the range you specify.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/clearEvents(in:)
func (m AVMusicTrack) ClearEventsInRange(range_ AVBeatRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("clearEventsInRange:"), range_)
}

// Splices all events in the beat range from the music track.
//
// range: The range of beats.
//
// # Discussion
//
// All events past the end of the range you specify shift backward by the
// duration of the range.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/cutEvents(in:)
func (m AVMusicTrack) CutEventsInRange(range_ AVBeatRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("cutEventsInRange:"), range_)
}

// Copies the events from the source track and splices them into the current
// music track.
//
// range: The range of beats.
//
// sourceTrack: The music track to copy the events from.
//
// insertStartBeat: The start beat to splice the events into.
//
// # Discussion
//
// All events originally at or past the insertion beat shift forward by the
// duration of the copied-in range.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/copyEvents(in:from:insertAt:)
func (m AVMusicTrack) CopyEventsInRangeFromTrackInsertAtBeat(range_ AVBeatRange, sourceTrack IAVMusicTrack, insertStartBeat AVMusicTimeStamp) {
	objc.Send[objc.ID](m.ID, objc.Sel("copyEventsInRange:fromTrack:insertAtBeat:"), range_, sourceTrack, insertStartBeat)
}

// Copies the events from the source track and merges them into the current
// music track.
//
// range: The range of beats.
//
// sourceTrack: The music track to copy the events from.
//
// mergeStartBeat: The start beat where the copied events merge into.
//
// # Discussion
//
// The system won’t modify events originally at or past the start beat.
// Copying events from track to track follows the same type-exclusion rules as
// adding events.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/copyAndMergeEvents(in:from:mergeAt:)
func (m AVMusicTrack) CopyAndMergeEventsInRangeFromTrackMergeAtBeat(range_ AVBeatRange, sourceTrack IAVMusicTrack, mergeStartBeat AVMusicTimeStamp) {
	objc.Send[objc.ID](m.ID, objc.Sel("copyAndMergeEventsInRange:fromTrack:mergeAtBeat:"), range_, sourceTrack, mergeStartBeat)
}

// Iterates through the music events within the track.
//
// range: The range to iterate through.
//
// block: The block to call for each event.
//
// # Discussion
//
// Examine each event the block returns by using [isKind(of:)] to determine
// the subclass, and then cast and access it accordingly.
//
// The iteration may continue after removing an event.
//
// The event object returned through the block won’t be the same instances
// you add to the [AVMusicTrack], though the content is identical.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/enumerateEvents(in:using:)
//
// [isKind(of:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isKind(of:)
func (m AVMusicTrack) EnumerateEventsInRangeUsingBlock(range_ AVBeatRange, block AVMusicEventEnumerationBlock) {
	objc.Send[objc.ID](m.ID, objc.Sel("enumerateEventsInRange:usingBlock:"), range_, block)
}

// A Boolean value that indicates whether the track is in a muted state.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/isMuted
func (m AVMusicTrack) Muted() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isMuted"))
	return rv
}
func (m AVMusicTrack) SetMuted(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setMuted:"), value)
}

// A Boolean value that indicates whether the track is in a soloed state.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/isSoloed
func (m AVMusicTrack) Soloed() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isSoloed"))
	return rv
}
func (m AVMusicTrack) SetSoloed(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setSoloed:"), value)
}

// The offset of the track’s start time, in beats.
//
// # Discussion
//
// By default, this value is `0`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/offsetTime
func (m AVMusicTrack) OffsetTime() AVMusicTimeStamp {
	rv := objc.Send[AVMusicTimeStamp](m.ID, objc.Sel("offsetTime"))
	return AVMusicTimeStamp(rv)
}
func (m AVMusicTrack) SetOffsetTime(value AVMusicTimeStamp) {
	objc.Send[struct{}](m.ID, objc.Sel("setOffsetTime:"), value)
}

// The time resolution value for the sequence, in ticks (pulses) per quarter
// note.
//
// # Discussion
//
// If you use a MIDI file to construct the containing sequence, the resolution
// is the contents of the file. If you want to keep a time resolution when
// writing a new file, retrieve this value and then specify it when writing to
// an audio sequencer. It doesn’t affect the rendering or notion of time of
// the sequence — only it’s MIDI file representation.
//
// By default, the framework sets this value to `480` when creating the
// sequence manually, or to a value from a MIDI file if you use it to create
// the sequence.
//
// You can only retrieve this value from the tempo track.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/timeResolution
func (m AVMusicTrack) TimeResolution() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("timeResolution"))
	return rv
}

// A Boolean value that indicates whether the track is an automation track.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/usesAutomatedParameters
func (m AVMusicTrack) UsesAutomatedParameters() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("usesAutomatedParameters"))
	return rv
}
func (m AVMusicTrack) SetUsesAutomatedParameters(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setUsesAutomatedParameters:"), value)
}

// The total duration of the track, in beats.
//
// # Discussion
//
// This property returns the beat of the last event in the track, plus any
// additional time that’s necessary to fade out the ending notes, or to
// round a loop point to a musical bar.
//
// If the user doesn’t set this value, the track length always adjusts to
// the end of the last active event in a track, and adjusts dynamically as the
// user adds or removes events.
//
// This property returns the maximum of the user-set track length or the
// calculated length.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/lengthInBeats
func (m AVMusicTrack) LengthInBeats() AVMusicTimeStamp {
	rv := objc.Send[AVMusicTimeStamp](m.ID, objc.Sel("lengthInBeats"))
	return AVMusicTimeStamp(rv)
}
func (m AVMusicTrack) SetLengthInBeats(value AVMusicTimeStamp) {
	objc.Send[struct{}](m.ID, objc.Sel("setLengthInBeats:"), value)
}

// The total duration of the track, in seconds.
//
// # Discussion
//
// This property returns the time of the last event in the track, plus any
// additional time that’s necessary to fade out the ending notes, or to
// round a loop point to a musical bar.
//
// If the user doesn’t set this value, the track length always adjusts to
// the end of the last active event in a track, and adjusts dynamically as the
// user adds or removes events.
//
// This property returns the maximum of the user-set track length or the
// calculated length.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/lengthInSeconds
func (m AVMusicTrack) LengthInSeconds() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("lengthInSeconds"))
	return rv
}
func (m AVMusicTrack) SetLengthInSeconds(value float64) {
	objc.Send[struct{}](m.ID, objc.Sel("setLengthInSeconds:"), value)
}

// The audio unit that receives the track’s events.
//
// # Discussion
//
// This property and a [DestinationMIDIEndpoint] are mutually exclusive. You
// must attach the audio to an audio engine to receive events. The track must
// be part of the [AVAudioSequencer] you associate with the same engine. When
// playing, the track sends it’s events to that [AVAudioUnit]. You can’t
// change the destination audio unit while the track’s sequence is in a
// playing state.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/destinationAudioUnit
func (m AVMusicTrack) DestinationAudioUnit() IAVAudioUnit {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("destinationAudioUnit"))
	return AVAudioUnitFromID(objc.ID(rv))
}
func (m AVMusicTrack) SetDestinationAudioUnit(value IAVAudioUnit) {
	objc.Send[struct{}](m.ID, objc.Sel("setDestinationAudioUnit:"), value)
}

// A Boolean value that indicates whether the track is in a looping state.
//
// # Discussion
//
// If you don’t set [LoopRange], the framework loops the full track.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/isLoopingEnabled
func (m AVMusicTrack) LoopingEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isLoopingEnabled"))
	return rv
}
func (m AVMusicTrack) SetLoopingEnabled(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setLoopingEnabled:"), value)
}

// The timestamp range for the loop, in beats.
//
// # Discussion
//
// You set the loop by specifying its beat range.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/loopRange
func (m AVMusicTrack) LoopRange() AVBeatRange {
	rv := objc.Send[AVBeatRange](m.ID, objc.Sel("loopRange"))
	return AVBeatRange(rv)
}
func (m AVMusicTrack) SetLoopRange(value AVBeatRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setLoopRange:"), value)
}

// The number of times the track’s loop repeats.
//
// # Discussion
//
// Use the value [AVMusicTrackLoopCountForever] to loop the track forever.
// Otherwise, valid values start at `1`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrack/numberOfLoops
func (m AVMusicTrack) NumberOfLoops() int {
	rv := objc.Send[int](m.ID, objc.Sel("numberOfLoops"))
	return rv
}
func (m AVMusicTrack) SetNumberOfLoops(value int) {
	objc.Send[struct{}](m.ID, objc.Sel("setNumberOfLoops:"), value)
}

// A timestamp you use to access all events in a music track through a beat
// range.
//
// See: https://developer.apple.com/documentation/avfaudio/avmusictimestampendoftrack
func (m AVMusicTrack) AVMusicTimeStampEndOfTrack() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("AVMusicTimeStampEndOfTrack"))
	return rv
}
func (m AVMusicTrack) SetAVMusicTimeStampEndOfTrack(value float64) {
	objc.Send[struct{}](m.ID, objc.Sel("setAVMusicTimeStampEndOfTrack:"), value)
}

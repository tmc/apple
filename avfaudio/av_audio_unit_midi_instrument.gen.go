// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAudioUnitMIDIInstrument] class.
var (
	_AVAudioUnitMIDIInstrumentClass     AVAudioUnitMIDIInstrumentClass
	_AVAudioUnitMIDIInstrumentClassOnce sync.Once
)

func getAVAudioUnitMIDIInstrumentClass() AVAudioUnitMIDIInstrumentClass {
	_AVAudioUnitMIDIInstrumentClassOnce.Do(func() {
		_AVAudioUnitMIDIInstrumentClass = AVAudioUnitMIDIInstrumentClass{class: objc.GetClass("AVAudioUnitMIDIInstrument")}
	})
	return _AVAudioUnitMIDIInstrumentClass
}

// GetAVAudioUnitMIDIInstrumentClass returns the class object for AVAudioUnitMIDIInstrument.
func GetAVAudioUnitMIDIInstrumentClass() AVAudioUnitMIDIInstrumentClass {
	return getAVAudioUnitMIDIInstrumentClass()
}

type AVAudioUnitMIDIInstrumentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitMIDIInstrumentClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitMIDIInstrumentClass) Alloc() AVAudioUnitMIDIInstrument {
	rv := objc.Send[AVAudioUnitMIDIInstrument](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents music devices or remote instruments.
//
// # Overview
//
// Use an [AVAudioUnitMIDIInstrument] in a chain that processes real-time
// (live) input and has the general concept of music events; for example,
// notes.
//
// # Creating a MIDI instrument
//
//   - [AVAudioUnitMIDIInstrument.InitWithAudioComponentDescription]: Creates a MIDI instrument audio unit with the component description you specify.
//
// # Sending information to the MIDI instrument
//
//   - [AVAudioUnitMIDIInstrument.SendControllerWithValueOnChannel]: Sends a MIDI controller event to the instrument.
//   - [AVAudioUnitMIDIInstrument.SendMIDIEventData1]: Sends a MIDI event which contains one data byte to the instrument.
//   - [AVAudioUnitMIDIInstrument.SendMIDIEventData1Data2]: Sends a MIDI event which contains two data bytes to the instrument.
//   - [AVAudioUnitMIDIInstrument.SendMIDISysExEvent]: Sends a MIDI System Exclusive event to the instrument.
//   - [AVAudioUnitMIDIInstrument.SendPitchBendOnChannel]: Sends a MIDI Pitch Bend event to the instrument.
//   - [AVAudioUnitMIDIInstrument.SendPressureOnChannel]: Sends a MIDI channel pressure event to the instrument.
//   - [AVAudioUnitMIDIInstrument.SendPressureForKeyWithValueOnChannel]: Sends a MIDI Polyphonic key pressure event to the instrument.
//   - [AVAudioUnitMIDIInstrument.SendProgramChangeOnChannel]: Sends MIDI Program Change and Bank Select events to the instrument.
//   - [AVAudioUnitMIDIInstrument.SendProgramChangeBankMSBBankLSBOnChannel]: Sends MIDI Program Change and Bank Select events to the instrument.
//   - [AVAudioUnitMIDIInstrument.SendMIDIEventList]: Sends a MIDI event list to the instrument.
//
// # Starting and stopping play
//
//   - [AVAudioUnitMIDIInstrument.StartNoteWithVelocityOnChannel]: Sends a MIDI Note On event to the instrument.
//   - [AVAudioUnitMIDIInstrument.StopNoteOnChannel]: Sends a MIDI Note Off event to the instrument.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument
type AVAudioUnitMIDIInstrument struct {
	AVAudioUnit
}

// AVAudioUnitMIDIInstrumentFromID constructs a [AVAudioUnitMIDIInstrument] from an objc.ID.
//
// An object that represents music devices or remote instruments.
func AVAudioUnitMIDIInstrumentFromID(id objc.ID) AVAudioUnitMIDIInstrument {
	return AVAudioUnitMIDIInstrument{AVAudioUnit: AVAudioUnitFromID(id)}
}

// NOTE: AVAudioUnitMIDIInstrument adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioUnitMIDIInstrument] class.
//
// # Creating a MIDI instrument
//
//   - [IAVAudioUnitMIDIInstrument.InitWithAudioComponentDescription]: Creates a MIDI instrument audio unit with the component description you specify.
//
// # Sending information to the MIDI instrument
//
//   - [IAVAudioUnitMIDIInstrument.SendControllerWithValueOnChannel]: Sends a MIDI controller event to the instrument.
//   - [IAVAudioUnitMIDIInstrument.SendMIDIEventData1]: Sends a MIDI event which contains one data byte to the instrument.
//   - [IAVAudioUnitMIDIInstrument.SendMIDIEventData1Data2]: Sends a MIDI event which contains two data bytes to the instrument.
//   - [IAVAudioUnitMIDIInstrument.SendMIDISysExEvent]: Sends a MIDI System Exclusive event to the instrument.
//   - [IAVAudioUnitMIDIInstrument.SendPitchBendOnChannel]: Sends a MIDI Pitch Bend event to the instrument.
//   - [IAVAudioUnitMIDIInstrument.SendPressureOnChannel]: Sends a MIDI channel pressure event to the instrument.
//   - [IAVAudioUnitMIDIInstrument.SendPressureForKeyWithValueOnChannel]: Sends a MIDI Polyphonic key pressure event to the instrument.
//   - [IAVAudioUnitMIDIInstrument.SendProgramChangeOnChannel]: Sends MIDI Program Change and Bank Select events to the instrument.
//   - [IAVAudioUnitMIDIInstrument.SendProgramChangeBankMSBBankLSBOnChannel]: Sends MIDI Program Change and Bank Select events to the instrument.
//   - [IAVAudioUnitMIDIInstrument.SendMIDIEventList]: Sends a MIDI event list to the instrument.
//
// # Starting and stopping play
//
//   - [IAVAudioUnitMIDIInstrument.StartNoteWithVelocityOnChannel]: Sends a MIDI Note On event to the instrument.
//   - [IAVAudioUnitMIDIInstrument.StopNoteOnChannel]: Sends a MIDI Note Off event to the instrument.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument
type IAVAudioUnitMIDIInstrument interface {
	IAVAudioUnit
	AVAudio3DMixing
	AVAudioMixing
	AVAudioStereoMixing

	// Topic: Creating a MIDI instrument

	// Creates a MIDI instrument audio unit with the component description you specify.
	InitWithAudioComponentDescription(description unsafe.Pointer) AVAudioUnitMIDIInstrument

	// Topic: Sending information to the MIDI instrument

	// Sends a MIDI controller event to the instrument.
	SendControllerWithValueOnChannel(controller uint8, value uint8, channel uint8)
	// Sends a MIDI event which contains one data byte to the instrument.
	SendMIDIEventData1(midiStatus uint8, data1 uint8)
	// Sends a MIDI event which contains two data bytes to the instrument.
	SendMIDIEventData1Data2(midiStatus uint8, data1 uint8, data2 uint8)
	// Sends a MIDI System Exclusive event to the instrument.
	SendMIDISysExEvent(midiData foundation.INSData)
	// Sends a MIDI Pitch Bend event to the instrument.
	SendPitchBendOnChannel(pitchbend uint16, channel uint8)
	// Sends a MIDI channel pressure event to the instrument.
	SendPressureOnChannel(pressure uint8, channel uint8)
	// Sends a MIDI Polyphonic key pressure event to the instrument.
	SendPressureForKeyWithValueOnChannel(key uint8, value uint8, channel uint8)
	// Sends MIDI Program Change and Bank Select events to the instrument.
	SendProgramChangeOnChannel(program uint8, channel uint8)
	// Sends MIDI Program Change and Bank Select events to the instrument.
	SendProgramChangeBankMSBBankLSBOnChannel(program uint8, bankMSB uint8, bankLSB uint8, channel uint8)
	// Sends a MIDI event list to the instrument.
	SendMIDIEventList(eventList unsafe.Pointer)

	// Topic: Starting and stopping play

	// Sends a MIDI Note On event to the instrument.
	StartNoteWithVelocityOnChannel(note uint8, velocity uint8, channel uint8)
	// Sends a MIDI Note Off event to the instrument.
	StopNoteOnChannel(note uint8, channel uint8)
}

// Init initializes the instance.
func (a AVAudioUnitMIDIInstrument) Init() AVAudioUnitMIDIInstrument {
	rv := objc.Send[AVAudioUnitMIDIInstrument](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitMIDIInstrument) Autorelease() AVAudioUnitMIDIInstrument {
	rv := objc.Send[AVAudioUnitMIDIInstrument](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitMIDIInstrument creates a new AVAudioUnitMIDIInstrument instance.
func NewAVAudioUnitMIDIInstrument() AVAudioUnitMIDIInstrument {
	class := getAVAudioUnitMIDIInstrumentClass()
	rv := objc.Send[AVAudioUnitMIDIInstrument](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a MIDI instrument audio unit with the component description you
// specify.
//
// description: The description of the audio component.
//
// # Return Value
//
// A new [AVAudioUnitMIDIInstrument] instance.
//
// # Discussion
//
// The component type must be `kAudioUnitType_MusicDevice` or
// `kAudioUnitType_RemoteInstrument`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/init(audioComponentDescription:)
func NewAudioUnitMIDIInstrumentWithAudioComponentDescription(description unsafe.Pointer) AVAudioUnitMIDIInstrument {
	instance := getAVAudioUnitMIDIInstrumentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAudioComponentDescription:"), description)
	return AVAudioUnitMIDIInstrumentFromID(rv)
}

// Creates a MIDI instrument audio unit with the component description you
// specify.
//
// description: The description of the audio component.
//
// description is a [audiotoolbox.AudioComponentDescription].
//
// # Return Value
//
// A new [AVAudioUnitMIDIInstrument] instance.
//
// # Discussion
//
// The component type must be `kAudioUnitType_MusicDevice` or
// `kAudioUnitType_RemoteInstrument`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/init(audioComponentDescription:)
func (a AVAudioUnitMIDIInstrument) InitWithAudioComponentDescription(description unsafe.Pointer) AVAudioUnitMIDIInstrument {
	rv := objc.Send[AVAudioUnitMIDIInstrument](a.ID, objc.Sel("initWithAudioComponentDescription:"), description)
	return rv
}

// Sends a MIDI controller event to the instrument.
//
// controller: Specifies a standard MIDI controller number. The valid range is `0` to
// `127`.
//
// value: Value for the controller. The valid range is `0` to `127`.
//
// channel: The channel number to send the event to. The valid range is `0` to `15`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/sendController(_:withValue:onChannel:)
func (a AVAudioUnitMIDIInstrument) SendControllerWithValueOnChannel(controller uint8, value uint8, channel uint8) {
	objc.Send[objc.ID](a.ID, objc.Sel("sendController:withValue:onChannel:"), controller, value, channel)
}

// Sends a MIDI event which contains one data byte to the instrument.
//
// midiStatus: The status value of the MIDI event.
//
// data1: The data byte of the MIDI event.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/sendMIDIEvent(_:data1:)
func (a AVAudioUnitMIDIInstrument) SendMIDIEventData1(midiStatus uint8, data1 uint8) {
	objc.Send[objc.ID](a.ID, objc.Sel("sendMIDIEvent:data1:"), midiStatus, data1)
}

// Sends a MIDI event which contains two data bytes to the instrument.
//
// midiStatus: The status value of the MIDI event.
//
// data1: The first data byte of the MIDI event.
//
// data2: The first data byte of the MIDI event.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/sendMIDIEvent(_:data1:data2:)
func (a AVAudioUnitMIDIInstrument) SendMIDIEventData1Data2(midiStatus uint8, data1 uint8, data2 uint8) {
	objc.Send[objc.ID](a.ID, objc.Sel("sendMIDIEvent:data1:data2:"), midiStatus, data1, data2)
}

// Sends a MIDI System Exclusive event to the instrument.
//
// midiData: The system exclusive data you want to send to the instrument.
//
// # Discussion
//
// The `midiData` parameter should contain the complete [SysEx] data,
// including start (F0) and termination (F7) bytes.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/sendMIDISysExEvent(_:)
func (a AVAudioUnitMIDIInstrument) SendMIDISysExEvent(midiData foundation.INSData) {
	objc.Send[objc.ID](a.ID, objc.Sel("sendMIDISysExEvent:"), midiData)
}

// Sends a MIDI Pitch Bend event to the instrument.
//
// pitchbend: Value of the pitchbend. The valid range of values is `0` to `16383`.
//
// channel: The channel number to send the event to. The valid range of values is `0`
// to `15`.
//
// # Discussion
//
// If this method isn’t invoked, then the system uses the default pitch bend
// value of `8192` (no pitch).
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/sendPitchBend(_:onChannel:)
func (a AVAudioUnitMIDIInstrument) SendPitchBendOnChannel(pitchbend uint16, channel uint8) {
	objc.Send[objc.ID](a.ID, objc.Sel("sendPitchBend:onChannel:"), pitchbend, channel)
}

// Sends a MIDI channel pressure event to the instrument.
//
// pressure: The value of the pressure. The valid range is `0` to `127`.
//
// channel: The channel number to send the event to. The valid range is `0` to `15`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/sendPressure(_:onChannel:)
func (a AVAudioUnitMIDIInstrument) SendPressureOnChannel(pressure uint8, channel uint8) {
	objc.Send[objc.ID](a.ID, objc.Sel("sendPressure:onChannel:"), pressure, channel)
}

// Sends a MIDI Polyphonic key pressure event to the instrument.
//
// key: The key (note) number to which the pressure event applies. The valid range
// is `0` to `127`.
//
// value: The value of the pressure. The valid range is `0` to `127`.
//
// channel: The channel number to send the event to. The valid range is `0` to `15`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/sendPressure(forKey:withValue:onChannel:)
func (a AVAudioUnitMIDIInstrument) SendPressureForKeyWithValueOnChannel(key uint8, value uint8, channel uint8) {
	objc.Send[objc.ID](a.ID, objc.Sel("sendPressureForKey:withValue:onChannel:"), key, value, channel)
}

// Sends MIDI Program Change and Bank Select events to the instrument.
//
// program: The program (preset) number within the bank to load. The valid range is `0`
// to `127`.
//
// channel: The channel number to send the event to. The valid range is `0` to `15`.
//
// # Discussion
//
// The system loads the instrument from the bank that was previously set by
// the MIDI “Bank Select” controller messages (0 and 31). The system uses
// bank `0` if not previously set.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/sendProgramChange(_:onChannel:)
func (a AVAudioUnitMIDIInstrument) SendProgramChangeOnChannel(program uint8, channel uint8) {
	objc.Send[objc.ID](a.ID, objc.Sel("sendProgramChange:onChannel:"), program, channel)
}

// Sends MIDI Program Change and Bank Select events to the instrument.
//
// program: Specifies the program (preset) number within the bank to load. The valid
// range is `0` to `127`.
//
// bankMSB: Specifies the most significant byte value for the bank to select. The valid
// range is `0` to `127`.
//
// bankLSB: Specifies the least significant byte value for the bank to select. The
// valid range is `0` to `127`.
//
// channel: The channel number to send the event to. The valid range is `0` to `15`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/sendProgramChange(_:bankMSB:bankLSB:onChannel:)
func (a AVAudioUnitMIDIInstrument) SendProgramChangeBankMSBBankLSBOnChannel(program uint8, bankMSB uint8, bankLSB uint8, channel uint8) {
	objc.Send[objc.ID](a.ID, objc.Sel("sendProgramChange:bankMSB:bankLSB:onChannel:"), program, bankMSB, bankLSB, channel)
}

// Sends a MIDI event list to the instrument.
//
// eventList: The MIDI event list.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/send(_:)
func (a AVAudioUnitMIDIInstrument) SendMIDIEventList(eventList unsafe.Pointer) {
	objc.Send[objc.ID](a.ID, objc.Sel("sendMIDIEventList:"), eventList)
}

// Sends a MIDI Note On event to the instrument.
//
// note: The note number (key) to play. The valid range is `0` to `127`.
//
// velocity: Specifies the volume to play the note at. The valid range is `0` to `127`.
//
// channel: The channel number to send the event to. The valid range is `0` to `15`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/startNote(_:withVelocity:onChannel:)
func (a AVAudioUnitMIDIInstrument) StartNoteWithVelocityOnChannel(note uint8, velocity uint8, channel uint8) {
	objc.Send[objc.ID](a.ID, objc.Sel("startNote:withVelocity:onChannel:"), note, velocity, channel)
}

// Sends a MIDI Note Off event to the instrument.
//
// note: The note number (key) to stop. The valid range is `0` to `127`.
//
// channel: The channel number to send the event to. The valid range is `0` to `15`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/stopNote(_:onChannel:)
func (a AVAudioUnitMIDIInstrument) StopNoteOnChannel(note uint8, channel uint8) {
	objc.Send[objc.ID](a.ID, objc.Sel("stopNote:onChannel:"), note, channel)
}

// Gets the audio mixing destination object that corresponds to the specified
// mixer node and input bus.
//
// mixer: The mixer to get destination details for.
//
// bus: The input bus.
//
// # Return Value
//
// Returns `self` if the specified mixer or input bus matches its connection
// point. If the mixer or input bus doesn’t match its connection point, or
// if the source node isn’t in a connected state to the mixer or input bus,
// the method returns `nil.`
//
// # Discussion
//
// When you connect a source node to multiple mixers downstream, setting
// [AVAudioMixing] properties directly on the source node applies the change
// to all of them. Use this method to get the corresponding
// [AVAudioMixingDestination] for a specific mixer. Properties set on
// individual destination instances don’t reflect at the source node level.
//
// If there’s any disconnection between the source and mixer nodes, the
// return value can be invalid. Fetch the return value every time you want to
// set or get properties on a specific mixer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixing/destination(forMixer:bus:)
func (a AVAudioUnitMIDIInstrument) DestinationForMixerBus(mixer IAVAudioNode, bus AVAudioNodeBus) IAVAudioMixingDestination {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("destinationForMixer:bus:"), mixer, bus)
	return AVAudioMixingDestinationFromID(rv)
}

// A value that simulates filtering of the direct path of sound due to an
// obstacle.
//
// # Discussion
//
// The value of `obstruction` is in decibels. The system blocks only the
// direct path of sound between the source and listener.
//
// The default value is `0.0`, and the range of valid values is `-100` to `0`.
// Only the [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/obstruction
func (a AVAudioUnitMIDIInstrument) Obstruction() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("obstruction"))
	return rv
}
func (a AVAudioUnitMIDIInstrument) SetObstruction(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setObstruction:"), value)
}

// A value that simulates filtering of the direct and reverb paths of sound
// due to an obstacle.
//
// # Discussion
//
// The value of `obstruction` is in decibels. The system blocks the direct and
// reverb paths of sound between the source and listener.
//
// The default value is `0.0`, and the range of valid values is `-100` to `0`.
// Only the [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/occlusion
func (a AVAudioUnitMIDIInstrument) Occlusion() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("occlusion"))
	return rv
}
func (a AVAudioUnitMIDIInstrument) SetOcclusion(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setOcclusion:"), value)
}

// The bus’s stereo pan.
//
// # Discussion
//
// The default value is `0.0`, and the range of valid values is `-1.0` to
// `1.0`. Only the [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioStereoMixing/pan
func (a AVAudioUnitMIDIInstrument) Pan() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("pan"))
	return rv
}
func (a AVAudioUnitMIDIInstrument) SetPan(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setPan:"), value)
}

// The in-head mode for a point source.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/pointSourceInHeadMode
func (a AVAudioUnitMIDIInstrument) PointSourceInHeadMode() AVAudio3DMixingPointSourceInHeadMode {
	rv := objc.Send[AVAudio3DMixingPointSourceInHeadMode](a.ID, objc.Sel("pointSourceInHeadMode"))
	return AVAudio3DMixingPointSourceInHeadMode(rv)
}
func (a AVAudioUnitMIDIInstrument) SetPointSourceInHeadMode(value AVAudio3DMixingPointSourceInHeadMode) {
	objc.Send[struct{}](a.ID, objc.Sel("setPointSourceInHeadMode:"), value)
}

// The location of the source in the 3D environment.
//
// # Discussion
//
// The system specifies the coordinates in meters. Only the
// [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/position
func (a AVAudioUnitMIDIInstrument) Position() AVAudio3DPoint {
	rv := objc.Send[AVAudio3DPoint](a.ID, objc.Sel("position"))
	return AVAudio3DPoint(rv)
}
func (a AVAudioUnitMIDIInstrument) SetPosition(value AVAudio3DPoint) {
	objc.Send[struct{}](a.ID, objc.Sel("setPosition:"), value)
}

// A value that changes the playback rate of the input signal.
//
// # Discussion
//
// A value of `2.0` results in the output audio playing one octave higher. A
// value of `0.5` results in the output audio playing one octave lower.
//
// The default value is `1.0`, and the range of valid values is `0.5` to
// `2.0`. Only the [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/rate
func (a AVAudioUnitMIDIInstrument) Rate() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("rate"))
	return rv
}
func (a AVAudioUnitMIDIInstrument) SetRate(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setRate:"), value)
}

// The type of rendering algorithm the mixer uses.
//
// # Discussion
//
// Depending on the current output format of the [AVAudioEnvironmentNode]
// instance, the system may only support a subset of the rendering algorithms.
// You can retrieve an array of valid rendering algorithms by calling the
// [ApplicableRenderingAlgorithms] function of the [AVAudioEnvironmentNode]
// instance.
//
// The default rendering algorithm is
// [AVAudio3DMixingRenderingAlgorithmEqualPowerPanning]. Only the
// [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/renderingAlgorithm
func (a AVAudioUnitMIDIInstrument) RenderingAlgorithm() AVAudio3DMixingRenderingAlgorithm {
	rv := objc.Send[AVAudio3DMixingRenderingAlgorithm](a.ID, objc.Sel("renderingAlgorithm"))
	return AVAudio3DMixingRenderingAlgorithm(rv)
}
func (a AVAudioUnitMIDIInstrument) SetRenderingAlgorithm(value AVAudio3DMixingRenderingAlgorithm) {
	objc.Send[struct{}](a.ID, objc.Sel("setRenderingAlgorithm:"), value)
}

// A value that controls the blend of dry and reverb processed audio.
//
// # Discussion
//
// This property controls the amount of the source’s audio that the
// [AVAudioEnvironmentNode] instance processes. A value of `0.5` results in an
// equal blend of dry and processed (wet) audio.
//
// The default is `0.0`, and the range of valid values is `0.0` (completely
// dry) to `1.0` (completely wet). Only the [AVAudioEnvironmentNode] class
// implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/reverbBlend
func (a AVAudioUnitMIDIInstrument) ReverbBlend() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("reverbBlend"))
	return rv
}
func (a AVAudioUnitMIDIInstrument) SetReverbBlend(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setReverbBlend:"), value)
}

// The source mode for the input bus of the audio environment node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/sourceMode
func (a AVAudioUnitMIDIInstrument) SourceMode() AVAudio3DMixingSourceMode {
	rv := objc.Send[AVAudio3DMixingSourceMode](a.ID, objc.Sel("sourceMode"))
	return AVAudio3DMixingSourceMode(rv)
}
func (a AVAudioUnitMIDIInstrument) SetSourceMode(value AVAudio3DMixingSourceMode) {
	objc.Send[struct{}](a.ID, objc.Sel("setSourceMode:"), value)
}

// The bus’s input volume.
//
// # Discussion
//
// The default value is `1.0`, and the range of valid values is `0.0` to
// `1.0`. Only the [AVAudioEnvironmentNode] and the [AVAudioMixerNode]
// implement this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixing/volume
func (a AVAudioUnitMIDIInstrument) Volume() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("volume"))
	return rv
}
func (a AVAudioUnitMIDIInstrument) SetVolume(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setVolume:"), value)
}

// Protocol methods for AVAudioMixing
